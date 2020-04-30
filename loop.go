package main

var (
	// Twins is a map that store the opposite of each start/end loop token
	Twins map[int]int
)

// InitTwins initialize twins map, only if OptimizeLoop is set to true
func InitTwins() {
	Twins = make(map[int]int)
}

// StartLoop check if selected register is equals to 0. If it is, then go the the corresponding end loop token.
func StartLoop() {
	if OptimizeLoops {
		if Registers[Ptr] == 0 {
			goTo, alreadyFound := Twins[InstNumber] // try to see if the matching brakcet instruction number is already known

			/*
				If it is, then go to this instruction
				If it is not, then find it, put it in the map and then go to the position
			*/

			if !alreadyFound { // let's find it and put it in the map, as well as the end -> start direction
				goTo = InstNumber
				var loopStack int = 1
				for loopStack != 0 {
					goTo++
					if Instructions[goTo] == '[' {
						loopStack++
					} else if Instructions[goTo] == ']' {
						loopStack--
					}
				}
				Twins[InstNumber] = goTo
				Twins[goTo] = InstNumber
			}

			InstNumber = goTo
		}
	} else {
		if Registers[Ptr] == 0 {
			var goTo int = InstNumber // goTo will become the instruction number of the matching end loop symbol
			var loopStack int = 1     // loopStack is the level of nested loops
			for loopStack != 0 {      // search for the matching end loop symbol
				goTo++
				if Instructions[goTo] == '[' {
					loopStack++
				} else if Instructions[goTo] == ']' {
					loopStack--
				}
			}
			InstNumber = goTo
		}
	}
}

// EndLoop check if selected register is equals to 0. If it is, then continue. Else, go to corresponding start loop token.
func EndLoop() {
	if OptimizeLoops {
		if Registers[Ptr] != 0 {
			goTo, alreadyFound := Twins[InstNumber]

			if !alreadyFound {
				goTo = InstNumber
				var loopStack int = 1
				for loopStack != 0 {
					goTo--
					if Instructions[goTo] == ']' {
						loopStack++
					} else if Instructions[goTo] == '[' {
						loopStack--
					}
				}
				Twins[InstNumber] = goTo
				Twins[goTo] = InstNumber
			}

			InstNumber = goTo
		}
	} else {
		if Registers[Ptr] != 0 {
			var goTo int = InstNumber
			var loopStack int = 1
			for loopStack != 0 { // seach for the matching start loop symbol
				goTo--
				if Instructions[goTo] == ']' {
					loopStack++
				} else if Instructions[goTo] == '[' {
					loopStack--
				}
			}
			InstNumber = goTo
		}
	}

}
