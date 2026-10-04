package main

import (
	"fmt"
	"math"
)

func lenghtOfLongestSubstring(s string) int {
	low, res := 0, math.MinInt
	mp := make(map[byte]int)

	for high := range s {
		mp[s[high]]++
		lenSubstr := high - low + 1

		for len(mp) < lenSubstr {
			mp[s[low]]--

			if mp[s[low]] == 0 {
				delete(mp, s[low])
			}
			low++
			lenSubstr = high - low + 1
		}
		res = max(res, lenSubstr)
	}
	return res
}

func checkError(expectedRes, result int) (int, string, error) {
	if expectedRes == result {
		return result, "output matched !!!", nil
	} else {
		return -1, "", fmt.Errorf("wrong ans, output does not match")
	}
}

func main() {
	inputString := "abcabcbb"
	expectedRes := 3
	result := lenghtOfLongestSubstring(inputString)

	res, msg, err := checkError(expectedRes, result)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		fmt.Println(msg)
	}
}
