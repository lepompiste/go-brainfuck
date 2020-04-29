package main

// StartLoop check if selected register is equals to 0. If it is, then go the the corresponding end loop token.
func StartLoop() {
	if registers[ptr] == 0 {
		var goTo int = instNumber // goTo will become the instruction number of the matching end loop symbol
		var loopStack int = 1     // loopStack is the level of nested loops
		for loopStack != 0 {      // search for the matching end loop symbol
			goTo++
			if instructions[goTo] == '[' {
				loopStack++
			} else if instructions[goTo] == ']' {
				loopStack--
			}
		}
		instNumber = goTo
	}
}

// EndLoop check if selected register is equals to 0. If it is, then continue. Else, go to corresponding start loop token.
func EndLoop() {
	if registers[ptr] != 0 {
		var goTo int = instNumber
		var loopStack int = 1
		for loopStack != 0 { // seach for the matching start loop symbol
			goTo--
			if instructions[goTo] == ']' {
				loopStack++
			} else if instructions[goTo] == '[' {
				loopStack--
			}
		}
		instNumber = goTo
	}
}
