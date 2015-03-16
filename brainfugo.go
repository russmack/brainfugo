// Brainfugo is a Go implementation of a Brainfuck interpreter.
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// main calls the brainfuck source scanner, then the evaluator.
func main() {
	src := scanSrc(os.Args[1])
	eval(src)
}

// scanSrc reads the brainfuck source code file.
func scanSrc(filename string) []byte {
	src, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Unable to read file:", err)
		os.Exit(0)
	}
	return src
}

// eval examines each byte of the brainfuck source code and processes it accordingly.
func eval(b []byte) {
	// The interpretor's memory registers.
	c := [30000]byte{}
	// The register pointer, points to the current register.
	p := 0
	// Keep track of nested loops.
	nestedLoops := 0
	// Iterate over the source code byte instructions.
	// i is instruction pointer, not to be confused with register pointer.
	for i := 0; i < len(b); i++ {
		switch string(b[i]) {
		case ">":
			// Move pointer to next register.
			p++
		case "<":
			// Move pointer to previous register.
			p--
		case "+":
			// Increment value in current register.
			c[p]++
		case "-":
			// Decrement value in current register.
			c[p]--
		case ".":
			// Output value in current register.
			fmt.Print(string(c[p]))
		case ",":
			// Read input into current register.
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			c[p] = input[0]
		case "[":
			// Start of a loop.
			// If the current register value is zero,
			// move the instruction pointer to the instruction following the matching close.
			if c[p] != 0 {
				// Continue to next source byte.
			} else {
				i++
				loopend := "]"[0]
				for nestedLoops > 0 || b[i] != loopend {
					if b[i] == 91 {
						// Count start of a nested loop.
						nestedLoops++
					}
					if b[i] == 93 {
						// Count end of a nested loop.
						nestedLoops--
					}
					i++
				}
			}
		case "]":
			// End of a loop.
			// If the current register value is non-zero,
			// then move instruction pointer to instruction following start of loop.
			i--
			loopbegin := "["[0]
			for nestedLoops > 0 || b[i] != loopbegin {
				if b[i] == 91 {
					nestedLoops--
				}
				if b[i] == 93 {
					nestedLoops++
				}
				i--
			}
			i--
		}
	}
}
