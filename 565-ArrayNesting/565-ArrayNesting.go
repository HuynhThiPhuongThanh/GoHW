// Last updated: 6/27/2025, 4:38:46 PM
package main

import "fmt"

func arrayNesting(nums []int) int {
	n := len(nums)
	visited := make([]bool, n)
	maxLen := 0

	for i := 0; i < n; i++ {
		if !visited[i] {
			count := 0
			current := i
			for !visited[current] {
				visited[current] = true
				current = nums[current]
				count++
			}
			if count > maxLen {
				maxLen = count
			}
		}
	}

	return maxLen
}

func pritn() {
	fmt.Println(arrayNesting([]int{5, 4, 0, 3, 1, 6, 2})) // Output: 4
	fmt.Println(arrayNesting([]int{0, 1, 2}))             // Output: 1
}
