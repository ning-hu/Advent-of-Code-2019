package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Gets the Manhatten distance given a string in the form "x,y".
func getDist(s string) int {
	vals := strings.Split(s, ",")
	x, err := strconv.Atoi(vals[0])
	check(err)
	y, err := strconv.Atoi(vals[1])
	check(err)

	return absInt(x) + absInt(y)
}

// Calculate the new coordinate after moving in the direction
// and steps indicated by the current node in the path.
func getNewCoord(s string, x, y int) (int, int, int, int) {
	dir, xnew, ynew, xsign, ysign := string(s[0]), x, y, 1, 1
	dist, err := strconv.Atoi(s[1:])
	check(err)

	switch dir {
	case "R":
		xnew, ynew = x+dist, y
	case "L":
		xnew, ynew, xsign = x-dist, y, -1
	case "U":
		xnew, ynew = x, y+dist
	case "D":
		xnew, ynew, ysign = x, y-dist, -1
	default:
		check(fmt.Errorf("Direction %s not recognized", dir))
	}

	return xnew, ynew, xsign, ysign
}

// Go through a path and construct a dictionary of xy coordinates
// which represents all positions along the path.
// "x,y": 1 if only one wire has passed through
// "x,y": 2 if two wires have passed through
func createDict(a []string, d map[string]int) {
	currx, curry, steps := 0, 0, 0
	for _, s := range a {
		x, y, xsign, ysign := getNewCoord(s, currx, curry)
		istart, jstart := currx, curry
		if x == currx {
			jstart = curry + ysign
		} else {
			istart = currx + xsign
		}

		for i := istart; i != x+xsign; i += xsign {
			for j := jstart; j != y+ysign; j += ysign {
				coord := strconv.Itoa(i) + "," + strconv.Itoa(j)
				steps += 1
				if _, ok := d[coord]; !ok {
					d[coord] = steps
				}
			}
		}

		currx, curry = x, y
	}
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	check(err)

	paths := strings.Split(string(content), "\n")
	path1 := strings.Split(paths[0], ",")
	path2 := strings.Split(paths[1], ",")

	d1, d2, m1, m2 := make(map[string]int), make(map[string]int), []string{}, []int{}

	createDict(path1, d1)
	createDict(path2, d2)

	for k, v1 := range d1 {
		if v2, ok := d2[k]; ok {
			m1 = append(m1, k)
			m2 = append(m2, v1+v2)
		}
	}

	// Sort the array of coordinates by Manhatten distance.
	sort.SliceStable(m1, func(i, j int) bool {
		return getDist(m1[i]) < getDist(m1[j])
	})

	// Sort the array of coordinates by total steps to intersection.
	sort.Ints(m2)

	fmt.Println("Manhatten distance for part 1:", getDist(m1[0]))
	fmt.Println("Total steps for part 2:", m2[0])
}
