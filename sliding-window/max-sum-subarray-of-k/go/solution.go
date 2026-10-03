package main

import (
	"fmt"
)

func maxSubarraySum(nums []int, k int) int {
	low, high, res, sum := 0, k-1, 0, 0
	for i := range k {
		sum += nums[i]
	}

	for high < len(nums) {
		res = max(sum, res)
		if high == len(nums)-1 {
			break
		}
		low++
		high++
		sum = sum - nums[low-1] + nums[high]
	}
	return res
}

func expectedError(res, expectedRes int) (int, string, error) {
	if res == expectedRes {
		return res, "output matched...with the expected result", nil
	} else {
		return -1, "", fmt.Errorf("error !!! output does not match with expected result")
	}
}

func main() {
	nums := []int{1, 4, 2, 10, 23, 3, 1, 0, 20}
	k := 4
	expectedRes := 39
	outputRes := maxSubarraySum(nums, k)
	res, message, err := expectedError(outputRes, expectedRes)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		fmt.Println(message)
	}
}
