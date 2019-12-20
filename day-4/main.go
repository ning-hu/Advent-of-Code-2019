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

func isValid(x int) (bool, bool) {
	// Format the string to be padded with zeros in case there
	// are leading zeros that don't appear in int format.
	a := strings.Split(fmt.Sprintf("%06d", x), "")

	// Keep an array that stores the dupe sizes found in x.
	// I.e. 111122: [4, 2]
	count, dupes, inc := 1, []int{}, true
	for i := 1; i < len(a); i++ {
		if a[i-1] == a[i] {
			count += 1
		} else {
			dupes = append(dupes, count)
			count = 1
		}

		if a[i-1] > a[i] {
			inc = false
		}
	}

	dupes = append(dupes, count)
	dupe1, dupe2 := false, false
	for j := range dupes {
		if dupes[j] == 2 {
			dupe2 = true
		}

		if dupes[j] >= 2 {
			dupe1 = true
		}
	}

	return dupe1 && inc, dupe2 && inc
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	check(err)

	bounds := strings.Split(string(content), "-")
	curr, _ := strconv.Atoi(bounds[0])
	end, _ := strconv.Atoi(bounds[1])
	count1, count2 := 0, 0

	for ; curr <= end; curr++ {
		c1, c2 := isValid(curr)

		if c1 {
			count1 += 1
		}

		if c2 {
			count2 += 1
		}
	}

	fmt.Println("Number of valid passwords for part 1 is:", count1)
	fmt.Println("Number of valid passwords for part 2 is:", count2)
}
