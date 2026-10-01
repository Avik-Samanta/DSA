package main

import (
	"fmt"
	"math"
	"slices"
)

func threeSumCloset(nums []int, target int) int {
	max_diff := math.MaxInt
	var res int
	slices.Sort(nums)
	for i := range nums {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				return sum
			} else if sum < target {
				l++
				diff := target - sum
				if max_diff > diff {
					max_diff = diff
					res = sum
				}
			} else {
				r--
				diff := sum - target
				if max_diff > diff {
					max_diff = diff
					res = sum
				}
			}
		}
	}
	return res
}

func checkError(res, exprectedRes int) (int, string, error) {
	if res == exprectedRes {
		return res, "output matched with the expected result", nil
	} else {
		return -1, "", fmt.Errorf("output does not match, plz check the code logic")
	}
}

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	expectedResult := 2
	result := threeSumCloset(nums, target)
	res, message, err := checkError(result, expectedResult)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	fmt.Println(message)
}
