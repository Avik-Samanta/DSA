package main

import (
	"fmt"
)

func sortedSquare(nums []int) []int {
	res := []int{}
	l, r, k := -1, 0, 0
	for _, value := range nums {
		if value < 0 {
			l++
		}
	}
	r = l + 1
	for l >= 0 && r < len(nums) {
		if nums[l]*nums[l] >= nums[r]*nums[r] {
			res = append(res, nums[r]*nums[r])
			r++
		} else {
			res = append(res, nums[l]*nums[l])
			l--
		}
		k++
	}
	for l >= 0 {
		res = append(res, nums[l]*nums[l])
		k++
		l--
	}
	for r < len(nums) {
		res = append(res, nums[r]*nums[r])
		r++
		k++
	}
	return res
}

func checkError(res, expectedNums []int) ([]int, string, error) {
	if len(res) != len(expectedNums) {
		return []int{}, "", fmt.Errorf("output length does not match with the ans lenght")
	}
	for i := range expectedNums {
		if res[i] != expectedNums[i] {
			return []int{}, "", fmt.Errorf("does not match with the expected result")
		}
	}
	return res, "code run successfully", nil
}

func main() {
	nums := []int{-4, -2, 0, 3, 10}
	expectedNums := []int{0, 4, 9, 16, 100}
	res := sortedSquare(nums)
	result, message, err := checkError(res, expectedNums)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
	fmt.Println(message)
}
