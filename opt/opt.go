package opt

import (
	"github.com/jphaws/cminusminus/ir"
	"github.com/jphaws/cminusminus/util"
)

func DeleteInstr(instr ir.Instr, instrBlocks map[ir.Instr]*ir.Block) {
	// Remove defs/uses for all registers in the instruction
	switch v := instr.(type) {
	case *ir.LoadInstr:
		v.Reg.Def = nil

		if reg, ok := v.Mem.(*ir.Register); ok {
			reg.DeleteUse(v)
		}

	case *ir.StoreInstr:
		if reg, ok := v.Reg.(*ir.Register); ok {
			reg.DeleteUse(v)
		}
		if reg, ok := v.Mem.(*ir.Register); ok {
			reg.DeleteUse(v)
		}

	case *ir.GepInstr:
		v.Target.Def = nil

		if reg, ok := v.Base.(*ir.Register); ok {
			reg.DeleteUse(v)
		}

	case *ir.CallInstr:
		if v.Target != nil {
			v.Target.Def = nil
		}

		for _, arg := range v.Arguments {
			if reg, ok := arg.(*ir.Register); ok {
				reg.DeleteUse(v)
			}
		}

	case *ir.RetInstr:
		if v.Src != nil {
			if reg, ok := v.Src.(*ir.Register); ok {
				reg.DeleteUse(v)
			}
		}

	case *ir.CompInstr:
		v.Target.Def = nil

		if reg, ok := v.Op1.(*ir.Register); ok {
			reg.DeleteUse(v)
		}
		if reg, ok := v.Op2.(*ir.Register); ok {
			reg.DeleteUse(v)
		}

	case *ir.BranchInstr:
		if v.Cond != nil {
			if reg, ok := v.Cond.(*ir.Register); ok {
				reg.DeleteUse(v)
			}
		}

	case *ir.BinaryInstr:
		v.Target.Def = nil

		if reg, ok := v.Op1.(*ir.Register); ok {
			reg.DeleteUse(v)
		}
		if reg, ok := v.Op2.(*ir.Register); ok {
			reg.DeleteUse(v)
		}

	case *ir.ConvInstr:
		v.Target.Def = nil

		if reg, ok := v.Src.(*ir.Register); ok {
			reg.DeleteUse(v)
		}

	case *ir.PhiInstr:
		v.Target.Def = nil

		for _, phiVal := range v.Values {
			if reg, ok := phiVal.Value.(*ir.Register); ok {
				reg.DeleteUse(v)
			}
		}
	}

	// Remove instruction from relevant block instruction slice
	defBlock := instrBlocks[instr]
	delete(instrBlocks, instr)

	switch instr.(type) {
	case *ir.PhiInstr:
		for i, v := range defBlock.Phis {
			if v == instr {
				defBlock.Phis = util.OrderedRemovePointerFromSlice(defBlock.Phis, i)
				return
			}
		}

	case *ir.AllocInstr:
		for i, v := range defBlock.Allocs {
			if v == instr {
				defBlock.Allocs = util.OrderedRemovePointerFromSlice(defBlock.Allocs, i)
				return
			}
		}

	default:
		for i, v := range defBlock.Instrs {
			if v == instr {
				defBlock.Instrs = util.OrderedRemoveFromSlice(defBlock.Instrs, i)
				v = nil // For GC efficiency
				return
			}
		}
	}
}

func RewriteUses(oldReg *ir.Register, newVal ir.Value) {
	for instr := range oldReg.Uses {
		numRemoved := 0

		switch v := instr.(type) {
		case *ir.LoadInstr:
			v.Mem = newVal
			numRemoved++

		case *ir.StoreInstr:
			if v.Mem == oldReg {
				v.Mem = newVal
				numRemoved++
			}
			if v.Reg == oldReg {
				v.Reg = newVal
				numRemoved++
			}

		case *ir.GepInstr:
			v.Base = newVal
			numRemoved++

		case *ir.CallInstr:
			for i, arg := range v.Arguments {
				if arg == oldReg {
					v.Arguments[i] = newVal
					numRemoved++
				}
			}

		case *ir.RetInstr:
			v.Src = newVal
			oldReg.DeleteUse(v)
			numRemoved++

		case *ir.CompInstr:
			if v.Op1 == oldReg {
				v.Op1 = newVal
				numRemoved++
			}
			if v.Op2 == oldReg {
				v.Op2 = newVal
				numRemoved++
			}

		case *ir.BranchInstr:
			v.Cond = newVal
			oldReg.DeleteUse(v)
			numRemoved++

		case *ir.BinaryInstr:
			if v.Op1 == oldReg {
				v.Op1 = newVal
				numRemoved++
			}
			if v.Op2 == oldReg {
				v.Op2 = newVal
				numRemoved++
			}

		case *ir.ConvInstr:
			v.Src = newVal
			oldReg.DeleteUse(v)
			numRemoved++

		case *ir.PhiInstr:
			for _, phiVal := range v.Values {
				if phiVal.Value == oldReg {
					phiVal.Value = newVal
					numRemoved++
				}
			}
		}

		for i := 0; i < numRemoved; i++ {
			oldReg.DeleteUse(instr)
			addUseIfRegister(newVal, instr)
		}
	}
}

func addUseIfRegister(val ir.Value, instr ir.Instr) {
	if reg, ok := val.(*ir.Register); ok {
		reg.AddUse(instr)
	}
}
