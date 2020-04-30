package main

var (
	// Ptr is a pseudo pointer to the pointed register. It is an int representing the index of the register in the array
	Ptr int = 0

	// Registers is an array (slice) which contains all the cells
	Registers []byte
)

func init() {
	Registers = make([]byte, 1, 30000)
}

// IncreasePointer increase to the pointer.
// As the registers are dynamically allocated, if the pointer overflow the number of registers, new registers are append to the memory.
func IncreasePointer() {
	Ptr++
	for Ptr >= len(Registers) {
		Registers = append(Registers, byte(0))
	}
}

// DecreasePointer decrease the pointer
// If the pointer goes under 0, an execution error is raised
func DecreasePointer() {
	Ptr--
	if Ptr < 0 {
		ExitWithError(ExecutionError{"Memory Error : index out of bounds", -1})
	}
}

// IncreaseRegister add 1 to the selected register
func IncreaseRegister() {
	Registers[Ptr]++
}

// DecreaseRegister substrac 1 to the selected register
func DecreaseRegister() {
	Registers[Ptr]--
}
