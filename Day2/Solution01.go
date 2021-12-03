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

	depth := 0
	len := 0
	for _, line := range lines {
		ops := strings.Split(line, " ")
		fmt.Printf("%s => %s\n", ops[0], ops[1])

		if opDist, err := strconv.Atoi(ops[1]); err == nil {
			switch ops[0] {
			case "forward":
				len += opDist
			case "up":
				depth -= opDist
			case "down":
				depth += opDist
			}
		} else {
			fmt.Printf(err.Error())
		}
	}

	fmt.Printf("Position:\n  - Depth: %d\n  - Length: %d\n  - Distance %d\n", depth, len, depth*len)
}
