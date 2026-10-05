package main

import (
	"fmt"
	"math"
	"slices"
)

func charcterReplacement(s string, k int) int {
	low, high, res := 0, 0, math.MinInt
	charArr := make([]int, 255)
	for i := range s {
		charArr[int(s[i])]++

		maxChar := slices.Max(charArr)
		lenStr := high - low + 1
		diff := lenStr - maxChar

		for diff > k {
			charArr[int(s[low])]--
			low++
			maxChar = slices.Max(charArr)
			lenStr = high - low + 1
			diff = lenStr - maxChar
		}
		lenStr = i - low + 1
		res = max(res, lenStr)
		high++
	}
	return res
}

func errorCheck(expectedResult, result int) (int, string, error) {
	if expectedResult == result {
		return result, "output matched....right ans !!!", nil
	} else {
		return -1, "", fmt.Errorf("wrong ans...check the code logic")
	}
}
func main() {
	inputString := "AABABBA"
	k := 1

	expectedResult := 4
	result := charcterReplacement(inputString, k)

	res, msg, err := errorCheck(expectedResult, result)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		fmt.Println(msg)
	}
}
