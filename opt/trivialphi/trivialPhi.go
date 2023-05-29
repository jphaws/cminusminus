package trivialphi // Mini

import (
	"sync"

	"github.com/keen-cp/compiler-project-c/ir"
	"github.com/keen-cp/compiler-project-c/opt"
)

func RemoveTrivialPhis(p *ir.ProgramIr) {
	var wg sync.WaitGroup

	for _, v := range p.Functions {
		wg.Add(1)
		go processFunction(v, &wg)
	}

	wg.Wait()
}

func processFunction(fn *ir.Function, wg *sync.WaitGroup) {
	// Collect all phis from the CFG
	nonTriv := []*ir.PhiInstr{}
	for inst := range fn.Instrs {
		if phi, ok := inst.(*ir.PhiInstr); ok {
			nonTriv = append(nonTriv, phi)
		}
	}

	// Get initial trivial functions
	triv, nonTriv := getTrivials(nonTriv)

	// Run until no more trivial phis are found
	for len(triv) != 0 {
		// Remove all currently known trivial phis
		removeTrivials(triv, fn.Instrs)

		// Find any new trivial phis (uncovered as a result of removing the previous ones)
		triv, nonTriv = getTrivials(nonTriv)
	}

	wg.Done()
}

func getTrivials(allPhis []*ir.PhiInstr) (triv map[*ir.PhiInstr]ir.Value, nonTriv []*ir.PhiInstr) {
	triv = map[*ir.PhiInstr]ir.Value{}

	for _, phi := range allPhis {
		trivial := true
		var firstOther ir.Value

		for _, v := range phi.Values {
			// Skip phi values which equal the target (generally the result of a backedge)
			if v.Value == phi.Target {
				continue
			}

			// Mark the phi as non-trivial if this value does not equal the first seen value
			if firstOther == nil {
				firstOther = v.Value
			}

			if !v.Value.IsEqual(firstOther) {
				trivial = false
				break
			}
		}

		if trivial {
			triv[phi] = firstOther
		} else {
			nonTriv = append(nonTriv, phi)
		}
	}

	return
}

func removeTrivials(triv map[*ir.PhiInstr]ir.Value, allInstrs map[ir.Instr]*ir.Block) {
	nonCanonMap := map[*ir.Register]ir.Value{}

	// Create a direct map from phi targets to their replacement values
	for phi, val := range triv {
		nonCanonMap[phi.Target] = val
	}

	// Canonize the map (some targets may point to values which themselves will be rewritten)
	// ex: During a rewrite, maybe r21 -> r10, but during this same rewrite, r10 -> r5
	//	   The map should be updated so r21 -> r5
	canonMap := canonizeMap(nonCanonMap)

	// Remove all trivial phi instructions and rewrite their target uses
	for phi := range triv {
		val := canonMap[phi.Target]

		// Bail out if the trivial phi does not need to have its uses updated
		if val == nil {
			continue
		}

		// Rewrite all uses of this phi target
		opt.RewriteUses(phi.Target, val)
		opt.DeleteInstr(phi, allInstrs)
	}
}

func canonizeMap(nonCanonMap map[*ir.Register]ir.Value) map[*ir.Register]ir.Value {
	canonMap := map[*ir.Register]ir.Value{}

	// Populate the canonical map with a map for each register in the non-canonical map
	for target, val := range nonCanonMap {
		for {
			// Check if the mapped replacement is a register; if not, stop
			var reg *ir.Register
			var ok bool
			if reg, ok = val.(*ir.Register); !ok {
				break
			}

			// Check if the mapped replacement is itself in the map; if not stop
			var newVal ir.Value
			var present bool
			if newVal, present = nonCanonMap[reg]; !present {
				break
			}

			// Check if the mapped replacement would be replaced by itself; if so, stop
			if newVal == nil {
				break
			}
			val = newVal
		}

		canonMap[target] = val
	}

	return canonMap
}
