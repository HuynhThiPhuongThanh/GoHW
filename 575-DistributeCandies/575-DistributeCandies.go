// Last updated: 6/27/2025, 4:38:43 PM
package main

import (
	"fmt"
)

func distributeCandies(candyType []int) int {
	typeSet := make(map[int]bool)

	for _, candy := range candyType {
		typeSet[candy] = true
	}

	maxCanEat := len(candyType) / 2
	numTypes := len(typeSet)

	if numTypes < maxCanEat {
		return numTypes
	}
	return maxCanEat
}

func print() {
	fmt.Println(distributeCandies([]int{1, 1, 2, 2, 3, 3}))
	fmt.Println(distributeCandies([]int{1, 1, 2, 3}))       
	fmt.Println(distributeCandies([]int{6, 6, 6, 6}))       
}
