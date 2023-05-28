package constprop

import (
	"fmt"
	"sync"

	"github.com/keen-cp/compiler-project-c/ir"
	"github.com/keen-cp/compiler-project-c/util"
)

// === Structs ===
type edge struct {
	src *ir.Block
	dst *ir.Block
}

func (e edge) String() string {
	srcLabel := "<nil>"
	if e.src != nil {
		srcLabel = e.src.Label()
	}

	dstLabel := "<nil>"
	if e.dst != nil {
		dstLabel = e.dst.Label()
	}

	return fmt.Sprintf("%v -> %v", srcLabel, dstLabel)
}

type propInfo struct {
	regs            map[*ir.Register]latticeType
	visitedBlocks   map[*ir.Block]bool
	workSet         map[*ir.Register]struct{}
	flowSet         map[edge]struct{}
	executableEdges map[edge]bool
	instrBlocks     map[ir.Instr]*ir.Block
}

type latticeType interface {
	isSameType(ol latticeType) bool
}

type latticeTop struct{}

func (l latticeTop) isSameType(ol latticeType) bool {
	_, ok := ol.(*latticeTop)
	return ok
}

func (l latticeTop) String() string {
	return "top"
}

type latticeConst struct {
	*ir.Literal
}

func (l latticeConst) isSameType(ol latticeType) bool {
	if v, ok := ol.(*latticeConst); !ok {
		return false
	} else {
		return l.IsEqual(v.Literal)
	}
}

func (l latticeConst) String() string {
	return fmt.Sprintf("const (%v)", l.Value)
}

type latticeBottom struct{}

func (l latticeBottom) isSameType(ol latticeType) bool {
	_, ok := ol.(*latticeBottom)
	return ok
}

func (l latticeBottom) String() string {
	return "bottom"
}

// === Constant propagation ===
func PropagateConstants(p *ir.ProgramIr) {
	var wg sync.WaitGroup

	for _, v := range p.Functions {
		wg.Add(1)
		go processCfg(&wg, v.Cfg, v.Registers)
	}
	wg.Wait()
}

func processCfg(wg *sync.WaitGroup, b *ir.Block, regs map[string]*ir.Register) {
	// Initialize info struct
	info := &propInfo{
		regs:            map[*ir.Register]latticeType{},
		visitedBlocks:   map[*ir.Block]bool{},
		workSet:         map[*ir.Register]struct{}{},
		flowSet:         map[edge]struct{}{},
		executableEdges: map[edge]bool{},
		instrBlocks:     map[ir.Instr]*ir.Block{},
	}

	// Set initial lattice type for each register in the function
	for _, v := range regs {
		if v.Def == nil {
			info.regs[v] = &latticeBottom{}
		} else {
			info.regs[v] = &latticeTop{}
		}
	}

	// Map instructions back to their owner blocks
	mappedBlocks := map[*ir.Block]bool{}
	mapInstrsToBlock(b, mappedBlocks, info)

	// Create edge into entry block to kickstart algorithm
	entryEdge := edge{
		src: nil,
		dst: b,
	}
	info.flowSet[entryEdge] = struct{}{}

	// Run sparse conditional constant propagation
	for len(info.flowSet) != 0 || len(info.workSet) != 0 {
		for edg := range info.flowSet {
			delete(info.flowSet, edg)

			if info.executableEdges[edg] {
				break
			}

			info.executableEdges[edg] = true

			if !info.visitedBlocks[edg.dst] {
				processBlock(edg.dst, info)
			} else {
				processPhis(edg.dst.Phis, info)
			}

			info.visitedBlocks[edg.dst] = true

			break
		}

		for reg := range info.workSet {
			delete(info.workSet, reg)

			for use := range reg.Uses {
				if phi, ok := use.(*ir.PhiInstr); ok {
					processPhiInstr(phi, info)

				} else if info.visitedBlocks[info.instrBlocks[use]] {
					processInstr(use, info)
				}
			}

			break
		}
	}

	cleanup(info)
	wg.Done()
}

func mapInstrsToBlock(b *ir.Block, visited map[*ir.Block]bool, info *propInfo) {
	visited[b] = true

	// Map each instruction to its owner block
	for _, phi := range b.Phis {
		info.instrBlocks[phi] = b
	}

	for _, alloc := range b.Allocs {
		info.instrBlocks[alloc] = b
	}

	for _, instr := range b.Instrs {
		info.instrBlocks[instr] = b
	}

	// Mark edges as initially non-executable
	if b.Next != nil {
		edg := edge{
			src: b,
			dst: b.Next,
		}
		info.executableEdges[edg] = false
	}

	if b.Els != nil {
		edg := edge{
			src: b,
			dst: b.Els,
		}
		info.executableEdges[edg] = false
	}

	// Process next block
	if b.Next != nil && !visited[b.Next] {
		mapInstrsToBlock(b.Next, visited, info)

	}

	// Process else block
	if b.Els != nil && !visited[b.Els] {
		mapInstrsToBlock(b.Els, visited, info)

		// Mark edge as initially non-executable
		edg := edge{
			src: b,
			dst: b.Els,
		}
		info.executableEdges[edg] = false
	}
}

func processBlock(b *ir.Block, info *propInfo) {
	processPhis(b.Phis, info)

	for _, allocs := range b.Allocs {
		processInstr(allocs, info)
	}

	for _, instr := range b.Instrs {
		processInstr(instr, info)
	}
}

func processPhis(phis []*ir.PhiInstr, info *propInfo) {
	for _, v := range phis {
		processPhiInstr(v, info)
	}
}

func processPhiInstr(phi *ir.PhiInstr, info *propInfo) {
	var resultType latticeType
	var firstConst *latticeConst

	// Loop through all phi operands
	for _, phiVal := range phi.Values {
		edge := edge{
			src: phiVal.Block,
			dst: info.instrBlocks[phi],
		}

		// Get current lattice type for this operand (or top if the operand's corresponsing edge is
		// not executable)
		var val latticeType
		if info.executableEdges[edge] {
			val = getLatticeType(phiVal.Value, info)
		} else {
			val = &latticeTop{}
		}

		// Check the operand's lattice type: immediately set the phi target to bottom if the operand
		// is bottom or there are now two *different* const operands
		switch v := val.(type) {
		case *latticeConst:
			if firstConst == nil {
				firstConst = v
			}

			if !v.IsEqual(firstConst.Literal) {
				resultType = &latticeBottom{}
				break
			}

		case *latticeBottom:
			resultType = &latticeBottom{}
			break
		}
	}

	// If no bottoms or consts seen as operands: phi target is top
	if resultType == nil && firstConst == nil {
		resultType = &latticeTop{}

	} else if resultType == nil {
		// If no bottoms seen: phi target is const
		resultType = firstConst
	}

	// Update the phi target lattice type
	oldType := info.regs[phi.Target]
	info.regs[phi.Target] = resultType

	// If the lattice type changed, add phi target to the work set
	if !oldType.isSameType(resultType) {
		info.workSet[phi.Target] = struct{}{}
	}
}

func processInstr(instr ir.Instr, info *propInfo) {
	var resultType latticeType
	var target *ir.Register

	switch v := instr.(type) {
	case *ir.AllocInstr:
		target = v.Target
		resultType = &latticeBottom{}
	case *ir.LoadInstr:
		target = v.Reg
		resultType = &latticeBottom{}
	case *ir.GepInstr:
		target = v.Target
		resultType = &latticeBottom{}
	case *ir.CallInstr:
		if v.Target != nil {
			target = v.Target
			resultType = &latticeBottom{}
		}
	case *ir.CompInstr:
		target = v.Target
		resultType = processCompInstr(v, info)
	case *ir.BranchInstr:
		processBranchInstr(v, info)
	case *ir.BinaryInstr:
		target = v.Target
		resultType = processBinaryInstr(v, info)
	case *ir.ConvInstr:
		target = v.Target
		resultType = processConvInstr(v, info)
	}

	if resultType == nil {
		return
	}

	// Update the instruction target lattice type
	oldType := info.regs[target]
	info.regs[target] = resultType

	// If the lattice type changed, add target to the work set
	if !oldType.isSameType(resultType) {
		info.workSet[target] = struct{}{}
	}
}

func processCompInstr(comp *ir.CompInstr, info *propInfo) latticeType {
	// Get operand lattice types from register map
	leftOp := getLatticeType(comp.Op1, info)
	rightOp := getLatticeType(comp.Op2, info)

	leftConst, leftOk := leftOp.(*latticeConst)
	rightConst, rightOk := rightOp.(*latticeConst)

	if leftOk && rightOk {
		// Otherwise, if both operands are constant, evaluate
		operResult, err := leftConst.DoCondition(rightConst.Literal, comp.Condition, comp.IsGuard)

		// If evaluation fails, treat target as a bottom
		if err != nil {
			return &latticeBottom{}
		}
		return &latticeConst{operResult}

	} else {
		return &latticeBottom{}
	}
}

func processBranchInstr(br *ir.BranchInstr, info *propInfo) {
	// Create outgoing edges
	nextEdge := edge{
		src: info.instrBlocks[br],
		dst: br.Next,
	}

	elseEdge := edge{
		src: info.instrBlocks[br],
		dst: br.Els,
	}

	// For unconditional jumps, only add the next edge
	if br.Cond == nil {
		info.flowSet[nextEdge] = struct{}{}
		return
	}

	op := getLatticeType(br.Cond, info)

	// With a constant condition, add the relevant edge only
	if constOp, ok := op.(*latticeConst); ok {
		if constOp.ToBool() {
			info.flowSet[nextEdge] = struct{}{}
		} else {
			info.flowSet[elseEdge] = struct{}{}
		}

	} else {
		// Otherwise, add both edges
		info.flowSet[nextEdge] = struct{}{}
		info.flowSet[elseEdge] = struct{}{}
	}
}

func processBinaryInstr(bin *ir.BinaryInstr, info *propInfo) latticeType {
	// Get operand lattice types from register map
	leftOp := getLatticeType(bin.Op1, info)
	rightOp := getLatticeType(bin.Op2, info)

	leftConst, leftOk := leftOp.(*latticeConst)
	rightConst, rightOk := rightOp.(*latticeConst)

	// When multiplying by a constant 0, propagate 0
	// Or when and'ing with false, propagate false
	if (bin.Operator == ir.MulOperator || bin.Operator == ir.AndOperator) &&
		((leftOk && leftConst.Value == "0") || (rightOk && rightConst.Value == "0")) {

		return &latticeConst{&ir.Literal{
			Value: "0",
			Type:  leftConst.GetType(),
		}}

	} else if bin.Operator == ir.OrOperator &&
		// When or'ing with true, propagate true
		((leftOk && leftConst.ToBool()) || (rightOk && rightConst.ToBool())) {

		return &latticeConst{&ir.Literal{
			Value: "1",
			Type:  leftConst.GetType(),
		}}

	} else if leftOk && rightOk {
		// Otherwise, if both operands are constant, evaluate
		operResult, err := leftConst.DoOperation(rightConst.Literal, bin.Operator)

		// If evaluation fails, treat target as a bottom
		if err != nil {
			return &latticeBottom{}
		}
		return &latticeConst{operResult}

	} else {
		return &latticeBottom{}
	}
}

func processConvInstr(conv *ir.ConvInstr, info *propInfo) latticeType {
	op := getLatticeType(conv.Src, info)

	if constOp, ok := op.(*latticeConst); ok {
		return &latticeConst{&ir.Literal{
			Value: constOp.Value,
			Type:  conv.Target.GetType(),
		}}
	} else {
		return &latticeBottom{}
	}
}

func getLatticeType(val ir.Value, info *propInfo) latticeType {
	switch v := val.(type) {
	case *ir.Register:
		return info.regs[v]
	case *ir.Literal:
		return &latticeConst{v}
	}

	return nil
}

// === Cleanup/removal ===
func cleanup(info *propInfo) {
	for reg, lat := range info.regs {
		if c, ok := lat.(*latticeConst); ok {
			rewriteUses(reg, c.Literal, info)
		}
	}

	// Executable Edges
	for edg, exec := range info.executableEdges {
		if !exec {
			removeEdge(edg.src, edg.dst)
		}
	}
}

func rewriteUses(reg *ir.Register, lit *ir.Literal, info *propInfo) {
	for instr := range reg.Uses {
		switch v := instr.(type) {
		case *ir.LoadInstr:
			if v.Mem == reg {
				v.Mem = lit
				reg.DeleteUse(v)
			}

		case *ir.StoreInstr:
			if v.Mem == reg {
				v.Mem = lit
				reg.DeleteUse(v)
			}
			if v.Reg == reg {
				v.Reg = lit
				reg.DeleteUse(v)
			}

		case *ir.GepInstr:
			v.Base = lit
			reg.DeleteUse(v)

		case *ir.CallInstr:
			for i, arg := range v.Arguments {
				if arg == reg {
					v.Arguments[i] = lit
					reg.DeleteUse(v)
				}
			}

		case *ir.RetInstr:
			v.Src = lit
			reg.DeleteUse(v)

		case *ir.CompInstr:
			if v.Op1 == reg {
				v.Op1 = lit
				reg.DeleteUse(v)
			}
			if v.Op2 == reg {
				v.Op2 = lit
				reg.DeleteUse(v)
			}

		case *ir.BranchInstr:
			v.Cond = lit
			reg.DeleteUse(v)

		case *ir.BinaryInstr:
			if v.Op1 == reg {
				v.Op1 = lit
				reg.DeleteUse(v)
			}
			if v.Op2 == reg {
				v.Op2 = lit
				reg.DeleteUse(v)
			}

		case *ir.ConvInstr:
			v.Src = lit
			reg.DeleteUse(v)

		case *ir.PhiInstr:
			for _, phiVal := range v.Values {
				if phiVal.Value == reg {
					phiVal.Value = lit
					reg.DeleteUse(v)
				}
			}
		}
	}
}

func removeEdge(src *ir.Block, dst *ir.Block) {
	br := &ir.BranchInstr{}

	// Update/remove the source next and else block pointers
	if src.Next == dst {
		br.Next = src.Els
		src.Next = src.Els

	} else {
		br.Next = src.Next
		src.Els = nil
	}

	// Replace the final branch instruction
	// The source instructions list is always guaranteed to have at least one instruction (the branch)
	src.Instrs[len(src.Instrs)-1] = br

	// Remove src from list of previous blocks in dst
	for i, prev := range dst.Prev {
		if prev == src {
			dst.Prev = util.RemovePointerFromSlice(dst.Prev, i)
			break
		}
	}

	// Remove src from all phis (where it is present) in dst
	for _, phi := range dst.Phis {
		for i, p := range phi.Values {
			if p.Block == src {
				phi.Values = util.RemovePointerFromSlice(phi.Values, i)
				break
			}
		}
	}
}
