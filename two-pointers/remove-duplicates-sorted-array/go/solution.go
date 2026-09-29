package main

import (
	"fmt"
)

func removeDuplicate(nums []int) int {
	l, r, k := 0, 1, 1
	for r < len(nums) {
		if nums[r] == nums[r-1] {
			r++
		} else {
			nums[l+1] = nums[r]
			r++
			l++
			k++
		}
	}
	return k
}

func checkError(nums []int, expectedNums []int, res int) (int, string, error) {
	if res != len(expectedNums) {
		return 0, "", fmt.Errorf("error in k")
	}
	for i := range res {
		if nums[i] != expectedNums[i] {
			return 0, "", fmt.Errorf("does not match with the expected result")
		}
	}
	return res, "code run successfully", nil
}

func main() {
	nums := []int{0, 0, 0, 1, 1, 1, 1, 2, 2, 3, 3, 4}
	expectedNums := []int{0, 1, 2, 3, 4}
	res := removeDuplicate(nums)
	result, message, err := checkError(nums, expectedNums, res)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
	fmt.Println(message)
}
