package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	buf, err := ioutil.ReadFile("input.txt")
	check(err)

	strs := strings.Split(string(buf), ",")
	instructions := make([]int, len(strs))
	for i := range instructions {
		instructions[i], err = strconv.Atoi(strs[i])
		check(err)
	}

	for x := 0; x <= 99; x++ {
		for y := 0; y <= 99; y++ {
			ints := make([]int, len(instructions))
			copy(ints, instructions)
			ints[1], ints[2] = x, y

			i := 0
			for ints[i] != 99 {
				pos1, pos2, pos3 := ints[i+1], ints[i+2], ints[i+3]
				input1, input2 := ints[pos1], ints[pos2]

				sum, invalid := 0, false
				switch ints[i] {
				case 1:
					sum = input1 + input2
				case 2:
					sum = input1 * input2
				default:
					invalid = true
				}

				if invalid {
					break
				}

				ints[pos3] = sum
				i += 4
			}

			if x == 12 && y == 2 {
				fmt.Println("Value at position 0 for part 1: ", ints[0])
			}

			if ints[0] == 19690720 {
				fmt.Println("Noun and verb for part 2 are: ", x, y)
				fmt.Println("100 * noun + verb: ", 100*x+y)
				return
			}
		}
	}
}
