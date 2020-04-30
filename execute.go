package main

var (
	// InstNumber is the address (index in slice) of the current instruction
	InstNumber int = 0

	// TotalInst is the address + 1 of the last instruction. Used to know when the program end is reached
	TotalInst int

	// Instructions is a n array (slice) containing all the instructions. If the -O flag isn't used, it may contain also comments but they're ignored by the executer (ExecuteInstruction function).
	Instructions []byte
)

func init() {
	Instructions = make([]byte, 0, 100)
}

// ExecuteInstruction execute the given instruction
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
	TotalInst = len(Instructions)

	for InstNumber < TotalInst {
		ExecuteInstruction(Instructions[InstNumber])
		InstNumber++
	}
}
