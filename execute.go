package main

var (
	instNumber   int = 0
	totalInst    int
	instructions []byte
)

// ExecuteInstruction execute the inst instruction
func ExecuteInstruction(inst byte) {
	switch inst {
	case '>':
		IncreasePointer()
	case '<':
		DecreasePointer()
	case '+':
		IncreaseRegister()
	case '-':
		DecreaseRegister()
	case '.':
		WriteRegister()
	case ',':
		ReadRegister()
	case '[':
		StartLoop()
	case ']':
		EndLoop()
	}
}

// Execute called to start the execution
// Control the end of the execution
func Execute() {
	totalInst = len(instructions)

	for instNumber < totalInst {
		ExecuteInstruction(instructions[instNumber])
		instNumber++
	}
}
