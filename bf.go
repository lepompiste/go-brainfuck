package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

var (
	// IgnoreCR is used to ignore carriage return in input
	IgnoreCR bool

	// OptimizeFile read only brainfuck's useful chars
	OptimizeFile bool

	// OptimizeLoops search only one time for the matching start/end token
	OptimizeLoops bool

	// Optimize set all Optimization vars to true if set to true
	Optimize bool

	// DurationBenchmark enable or disable time benchmarking
	DurationBenchmark bool
)

var (
	start      time.Time
	endparsing time.Time
)

// ExecutionError is error of execution datatype structure
type ExecutionError struct {
	Message string
	ID      int
}

func check(e error) {
	if e != nil {
		fmt.Println("Une erreur est survenue...")
		os.Exit(-1)
	}
}

// ExitWithError quits the program printing message
func ExitWithError(error ExecutionError) {
	os.Exit(error.ID)
	fmt.Println(error.Message)
}

// IsBFChar take a token and return wether it is a brainfuck char or not
func IsBFChar(t byte) bool {
	switch t {
	case '<', '>', '+', '-', '.', ',', '[', ']':
		return true
	default:
		return false
	}
}

func main() {
	// Parsing flags and args
	flag.BoolVar(&IgnoreCR, "ignore-cr", true, "Ignore the carriage return char in input.")
	flag.BoolVar(&OptimizeFile, "OF", false, "Optimize the file, reading only useful chars (startup may be longer)")
	flag.BoolVar(&OptimizeLoops, "OL", false, "Optimize the loops")
	flag.BoolVar(&Optimize, "O", false, "Optimize all can be optimized (currently loops and file)")
	flag.BoolVar(&DurationBenchmark, "duration", false, "At the end of execution, print in ms how long did the parsing and execution take")
	flag.Parse()

	if Optimize {
		OptimizeFile, OptimizeLoops = true, true
	}

	if OptimizeLoops {
		InitTwins()
		// As twins are optional, it may be manually initialized.
		// Other global structures are automatically initialized via init() function
	}

	if DurationBenchmark {
		start = time.Now()
	}

	if len(flag.Args()) == 1 { // checking usage

		// Opening and reading the source file
		file, err := os.Open(flag.Args()[0])
		check(err)

		chars := make([]byte, 100, 100)
		n, err := file.Read(chars) // Reading first 100 chars

		for {
			if err == nil {
				if !OptimizeFile {
					Instructions = append(Instructions, chars[:n]...)
				} else {
					for i := 0; i < n; i++ {
						if IsBFChar(chars[i]) { //If OptimizeFile is true, then only brainfuck chars are added to the intructions queue
							Instructions = append(Instructions, chars[i])
						}
					}
				}

				n, err = file.Read(chars) // Reading next 100 chars, if the end of file is reached, the next loop iteration will go to else and evaluate the error as io.EOF
			} else {
				if errors.Is(err, io.EOF) {
					goto reading_end // If the error is EOF, then stop reading the file and start execution
				} else {
					fmt.Println("Error reading the file : " + err.Error())
					os.Exit(1)
				}
			}
		}
	reading_end:

		if DurationBenchmark {
			endparsing = time.Now()
		}

		Execute() // Start the execution of the instructions queue

		if DurationBenchmark {
			end := time.Now()
			fmt.Println("\nParsing duration :", endparsing.Sub(start))
			fmt.Println("Execution duration :", end.Sub(endparsing))
			fmt.Println("Total duration :", end.Sub(start))
		}
	} else {
		fmt.Println("Usage : go-bf [options] <file>")
	}
}
