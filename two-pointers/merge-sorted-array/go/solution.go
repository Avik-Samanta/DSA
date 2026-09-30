package main

import (
	"fmt"
)

func merge(nums1, nums2 []int, m, n int) {
	l, r, k := m-1, n-1, m+n-1
	for l >= 0 && r >= 0 {
		if nums1[l] >= nums2[r] {
			nums1[k] = nums1[l]
			l--
		} else {
			nums1[k] = nums2[r]
			r--
		}
		k--
	}
	for r >= 0 {
		nums1[k] = nums2[r]
		r--
		k--
	}
}

func checkError(nums, expectedNums []int, m, n int) (string, error) {
	for i := range m + n {
		if nums[i] != expectedNums[i] {
			return "", fmt.Errorf("does not match with the expected result")
		}
	}
	return "code run successfully", nil
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 3, 5}
	m, n := 3, 3
	expectedNums := []int{1, 2, 2, 3, 3, 5}
	merge(nums1, nums2, m, n)
	message, err := checkError(nums1, expectedNums, m, n)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(message)

}
