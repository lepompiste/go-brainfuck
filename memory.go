package main

var (
	ptr       int = 0
	registers []byte
)

func init() {
	registers = make([]byte, 1, 30000)
}

// IncreasePointer increase to the pointer.
// As the registers are dynamically allocated, if the pointer overflow the number of registers, new registers are append to the memory.
func IncreasePointer() {
	ptr++
	for ptr >= len(registers) {
		registers = append(registers, byte(0))
	}
}

// DecreasePointer decrease the pointer
// If the pointer goes under 0, an execution error is raised
func DecreasePointer() {
	ptr--
	if ptr < 0 {
		ExitWithError(ExecutionError{"Memory Error : index out of bounds", -1})
	}
}

// IncreaseRegister add 1 to the selected register
func IncreaseRegister() {
	registers[ptr]++
}

// DecreaseRegister substrac 1 to the selected register
func DecreaseRegister() {
	registers[ptr]--
}
