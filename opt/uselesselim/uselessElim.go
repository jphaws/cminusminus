package uselesselim

import (
	"sync"

	"github.com/keen-cp/compiler-project-c/ir"
	"github.com/keen-cp/compiler-project-c/opt"
)

func EliminateUselessCode(p *ir.ProgramIr) {
	var wg sync.WaitGroup

	for _, v := range p.Functions {
		wg.Add(1)
		go processFunction(v, &wg)
	}

	wg.Wait()
}

func processFunction(fn *ir.Function, wg *sync.WaitGroup) {
	// Mark all instruction as initially useless
	markedInstrs := map[ir.Instr]bool{}

	for k := range fn.Instrs {
		markedInstrs[k] = false
	}

	// Mark critical instructions as useful
	markUseful(markedInstrs)

	// Remove useless instructions
	sweep(markedInstrs, fn.Instrs)
	wg.Done()
}

func markUseful(markedInstrs map[ir.Instr]bool) {
	workSet := map[ir.Instr]struct{}{}

	// Find initial critical instructions
	for v := range markedInstrs {
		switch v.(type) {
		case *ir.BranchInstr, *ir.StoreInstr, *ir.CallInstr, *ir.RetInstr:
			markInstr(v, markedInstrs, workSet)
		}
	}

	// Repeatedly mark dependencies of the critical instructions
	for len(workSet) != 0 {
		for v := range workSet {
			delete(workSet, v)
			markDependencies(v, markedInstrs, workSet)
		}
	}
}

func markDependencies(instr ir.Instr,
	markedInstrs map[ir.Instr]bool, workSet map[ir.Instr]struct{}) {

	switch v := instr.(type) {
	case *ir.LoadInstr:
		if reg, ok := v.Mem.(*ir.Register); ok {
			markInstr(reg.Def, markedInstrs, workSet)
		}

	case *ir.StoreInstr:
		if reg, ok := v.Reg.(*ir.Register); ok {
			markInstr(reg.Def, markedInstrs, workSet)
		}
		if reg, ok := v.Mem.(*ir.Register); ok {
			markInstr(reg.Def, markedInstrs, workSet)
		}

	case *ir.GepInstr:
		if reg, ok := v.Base.(*ir.Register); ok {
			markInstr(reg.Def, markedInstrs, workSet)
		}

	case *ir.CallInstr:
		for _, arg := range v.Arguments {
			if reg, ok := arg.(*ir.Register); ok {
				markInstr(reg.Def, markedInstrs, workSet)
			}
		}

	case *ir.RetInstr:
		if v.Src != nil {
			if reg, ok := v.Src.(*ir.Register); ok {
				markInstr(reg.Def, markedInstrs, workSet)
			}
		}

	case *ir.CompInstr:
		if reg, ok := v.Op1.(*ir.Register); ok {
			markInstr(reg.Def, markedInstrs, workSet)
		}
		if reg, ok := v.Op2.(*ir.Register); ok {
			markInstr(reg.Def, markedInstrs, workSet)
		}

	case *ir.BranchInstr:
		if v.Cond != nil {
			if reg, ok := v.Cond.(*ir.Register); ok {
				markInstr(reg.Def, markedInstrs, workSet)
			}
		}

	case *ir.BinaryInstr:
		if reg, ok := v.Op1.(*ir.Register); ok {
			markInstr(reg.Def, markedInstrs, workSet)
		}
		if reg, ok := v.Op2.(*ir.Register); ok {
			markInstr(reg.Def, markedInstrs, workSet)
		}

	case *ir.ConvInstr:
		if reg, ok := v.Src.(*ir.Register); ok {
			markInstr(reg.Def, markedInstrs, workSet)
		}

	case *ir.PhiInstr:
		for _, phiVal := range v.Values {
			if reg, ok := phiVal.Value.(*ir.Register); ok {
				markInstr(reg.Def, markedInstrs, workSet)
			}
		}
	}
}

func markInstr(instr ir.Instr, markedInstrs map[ir.Instr]bool, workSet map[ir.Instr]struct{}) {
	if instr != nil && !markedInstrs[instr] {
		markedInstrs[instr] = true
		workSet[instr] = struct{}{}
	}
}

func sweep(markedInstrs map[ir.Instr]bool, instrs map[ir.Instr]*ir.Block) {
	for inst, marked := range markedInstrs {
		if !marked {
			opt.DeleteInstr(inst, instrs)
		}
	}
}
