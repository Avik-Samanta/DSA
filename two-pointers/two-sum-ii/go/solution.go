package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		total := numbers[l] + numbers[r]
		if total == target {
			return []int{l + 1, r + 1}
		} else if total < target {
			l++
		} else {
			r--
		}
	}
	return []int{}
}

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(numbers, target)
	fmt.Println(result)
}
