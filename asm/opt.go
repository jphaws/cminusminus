package asm

func removeTrivialMovs(blocks []*Block) {
	for _, b := range blocks {
		b.Instrs = trivialMov(b.Instrs)
		b.PhiOuts = trivialMov(b.PhiOuts)
	}
}

func trivialMov(instrs []Instr) []Instr {
	replace := instrs[:0]

	// Filter instruction slice values
	for _, inst := range instrs {
		// Ignore trivial move instructions (moves with the same source and destination)
		if mov, movOk := inst.(*MovInstr); movOk {
			if src, srcOk := mov.Src.(*Register); srcOk {
				if mov.Dst == src {
					continue
				}
			}
		}

		replace = append(replace, inst)
	}

	// Free extra instructions for potential garbage collection
	for i := len(replace); i < len(instrs); i++ {
		instrs[i] = nil
	}

	return replace
}
