package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	src := scanSrc(os.Args[1])
	eval(src)
}

func scanSrc(filename string) []byte {
	src, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Unable to read file:", err)
		os.Exit(0)
	}
	return src
}

func eval(b []byte) {
	c := [30000]byte{}
	p := 0
	nestedLoops := 0
	for i := 0; i < len(b); i++ {
		switch string(b[i]) {
		case ">":
			p++
		case "<":
			p--
		case "+":
			c[p]++
		case "-":
			c[p]--
		case ".":
			fmt.Print(string(c[p]))
		case ",":
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			c[p] = input[0]
		case "[":
			if c[p] != 0 {
			} else {
				i++
				lend := "]"[0]
				for nestedLoops > 0 || b[i] != lend {
					if b[i] == 91 {
						nestedLoops++
					}
					if b[i] == 93 {
						nestedLoops--
					}
					i++
				}
			}
		case "]":
			i--
			lbegin := "["[0]
			for nestedLoops > 0 || b[i] != lbegin {
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
