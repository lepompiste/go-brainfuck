package main

import (
	"os"
)

// WriteRegister prints the content of the selected register as ASCII char
func WriteRegister() {
	os.Stdout.Write([]byte{Registers[Ptr]})
}

// ReadRegister reads the input into the selected register
func ReadRegister() {
	buffer := make([]byte, 1, 1)

	os.Stdin.Read(buffer) // Read char from buffer

	// ignoring carriage return depending on the flag
	if IgnoreCR {
		for buffer[0] == '\r' {
			os.Stdin.Read(buffer)
		}
	}

	Registers[Ptr] = buffer[0] // Set it in register

	// Flush stdin, maybe as an option, DISABLED NOW
	/*
		for ; buffer[0] != '\n'; os.Stdin.Read(buffer) {
		}
	*/
}
