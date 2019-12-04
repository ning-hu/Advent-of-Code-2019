package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	sum1 := 0
	sum2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		check(err)

		// For part 1
		mass1 := (mass / 3) - 2
		sum1 += mass1

		// For part 2
		mass2 := mass1
		for mass2 > 0 {
			sum2 += mass2
			mass2 = (mass2 / 3) - 2
		}
	}

	err = scanner.Err()
	check(err)

	fmt.Println("Sum for part 1: ", sum1)
	fmt.Println("Sum for part 2: ", sum2)
}
