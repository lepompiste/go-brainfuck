package main

import (
	"os"
)

// WriteRegister prints the content of the selected register as ASCII char
func WriteRegister() {
	os.Stdout.Write([]byte{registers[ptr]})
}

// ReadRegister reads the input into the selected register
func ReadRegister() {
	buffer := make([]byte, 1, 1)

	os.Stdin.Read(buffer)

	// TODO make it as option
	for buffer[0] == '\r' { // ignoring carriage return
		os.Stdin.Read(buffer) // Read char from buffer
	}

	registers[ptr] = buffer[0] // Set it in register

	// TODO make it as option
	for ; buffer[0] != '\n'; os.Stdin.Read(buffer) { // Empty buffer
	}
}
