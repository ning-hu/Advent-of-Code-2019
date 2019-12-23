package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const NUMMODES = 3

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Extracts the modes from an instruction.
func getModes(instn int) []int {
	instn, modes, i := instn/100, make([]int, NUMMODES), 0
	for instn != 0 {
		modes[i] = instn % 10
		instn = instn / 10
		i += 1
	}
	return modes
}

func intcode(ints []int) {
	i := 0
	for ints[i] != 99 {
		// To account for different modes, get the modes for the
		// current instruction. Set the values in the params array
		// so they are the index to the desired value in the ints
		// array.
		instn := ints[i]
		modes, params := getModes(instn), []int{}
		for j, mode := range modes {
			switch mode {
			case 1: // treat the value at the param as the desired value
				params = append(params, i+j+1)
			default: // treat the value at the param as an index to the desired value
				params = append(params, ints[i+j+1])
			}
		}

		// mod 100 since mode is the tens and ones digit.
		switch instn % 100 {
		case 1:
			ints[params[2]] = ints[params[0]] + ints[params[1]]
			i += 4
		case 2:
			ints[params[2]] = ints[params[0]] * ints[params[1]]
			i += 4
		case 3: // No immediate
			var input string
			fmt.Println("Opcode 3: enter an integer to be inserted in position", ints[i+1])
			_, err := fmt.Scanln(&input)
			check(err)
			num, err := strconv.ParseInt(input, 10, 64)
			check(err)

			ints[params[0]] = int(num)
			i += 2
		case 4:
			fmt.Println("Opcode 4: the value at position", ints[i+1], "is", ints[params[0]])
			i += 2
		case 5:
			// jump-if-nonzero
			if ints[params[0]] != 0 {
				i = ints[params[1]]
			} else {
				i += 3
			}
		case 6:
			// jump-if-zero
			if ints[params[0]] == 0 {
				i = ints[params[1]]
			} else {
				i += 3
			}
		case 7:
			if ints[params[0]] < ints[params[1]] {
				ints[params[2]] = 1
			} else {
				ints[params[2]] = 0
			}
			i += 4
		case 8:
			if ints[params[0]] == ints[params[1]] {
				ints[params[2]] = 1
			} else {
				ints[params[2]] = 0
			}
			i += 4
		default:
			check(fmt.Errorf("Invalid Intcode instruction %d\n", instn))
		}
	}
}

func main() {
	buf, err := ioutil.ReadFile("input.txt")
	check(err)

	strs := strings.Split(string(buf), ",")
	ints := make([]int, len(strs))
	for i := range ints {
		ints[i], err = strconv.Atoi(strs[i])
		check(err)
	}

	intcode(ints)
}
