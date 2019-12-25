package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	// A)B => orbits[B] = A
	orbits := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		objs := strings.Split(scanner.Text(), ")")
		orbits[objs[1]] = objs[0]
	}

	count := 0
	for k, _ := range orbits {
		curr := k
		for curr != "COM" {
			count += 1
			curr = orbits[curr]
		}
	}

	steps1, steps2 := 0, 0
	curr1, curr2 := orbits["YOU"], orbits["SAN"]
	orb1, orb2 := make(map[string]int), make(map[string]int) // obj: steps to get to obj
	for curr1 != "COM" || curr2 != "COM" {
		if curr1 != "COM" {
			orb1[curr1] = steps1
			steps1 += 1
			curr1 = orbits[curr1]
		}

		if curr2 != "COM" {
			orb2[curr2] = steps2
			steps2 += 1
			curr2 = orbits[curr2]
		}
	}

	shortest := count
	for k, v := range orb1 {
		if val, ok := orb2[k]; ok && v+val < shortest {
			shortest = v + val
		}
	}

	fmt.Println("Number of orbits for part 1:", count)
	fmt.Println("Number of orbits between santa and me for part 2:", shortest)
}
