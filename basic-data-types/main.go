package main

import (
	"fmt"
	"sort"
)

func add(x int, y int) int {
	return x + y
}

// if the params have the same type, you can just use one
func subtract(x, y int) int {
	return y - x
}

// return the numbers as in order
func order(x, y, z int) (int, int, int) {
	s := []int{x, y, z}
	sort.Ints(s)
	return s[0], s[1], s[2]
}

func main() {
	// fmt.Println(add(42, 13))
	// fmt.Println(subtract(13, 42))
	x, y, z := order(3, 2, 4)
	fmt.Println(x, y, z)
}
