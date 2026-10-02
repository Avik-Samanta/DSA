package main

import (
	"fmt"
	"slices"
)

func sortColors(nums []int) {
	low, mid, high := 0, 0, len(nums)-1
	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[mid], nums[low] = nums[low], nums[mid]
			low++
			mid++
		case 1:
			mid++
		case 2:
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
}

func checkError(nums, expectedNums []int) ([]int, string, error) {
	if slices.Equal(nums, expectedNums) {
		return nums, "output matched with expected result. Correct !!!", nil
	}
	return []int{}, "", fmt.Errorf("out put does not match....wrong ans")
}

func main() {
	nums := []int{2, 0, 1, 2, 0, 1}
	expectedNums := slices.Clone(nums)
	slices.Sort(expectedNums)
	sortColors(nums)
	res, message, err := checkError(nums, expectedNums)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		fmt.Println(message)
	}
}
