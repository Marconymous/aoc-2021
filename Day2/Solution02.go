package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("input")
	lines := strings.Split(string(data), "\r\n")

	aim := 0
	len := 0
	depth := 0
	for _, line := range lines {
		ops := strings.Split(line, " ")
		fmt.Printf("%s => %s\n", ops[0], ops[1])

		if opDist, err := strconv.Atoi(ops[1]); err == nil {
			switch ops[0] {
			case "forward":
				depth += opDist * aim
				len += opDist
			case "up":
				aim -= opDist
			case "down":
				aim += opDist
			}
		} else {
			fmt.Printf(err.Error())
		}
	}

	fmt.Printf("Position:\n  - Aim: %d\n  - Length: %d\n  - Depth: %d\n  - Distance %d\n", aim, len, depth, depth*len)
}
