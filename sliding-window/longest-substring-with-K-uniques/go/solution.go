package main

import (
	"fmt"
	"math"
)

func longestKSubstr(s string, k int) int {
	low, high, res := 0, 0, math.MinInt
	mp := make(map[byte]int)

	for high < len(s) {
		mp[s[high]]++

		for len(mp) > k {
			mp[s[low]]--

			if mp[s[low]] == 0 {
				delete(mp, s[low])
			}
			low++
		}

		if len(mp) == k {
			len := high - low + 1
			res = max(res, len)
		}
		high++
	}
	if res == math.MinInt {
		return -1
	} else {
		return res
	}
}

func checkResult(expectedRes, res int) (int, string, error) {
	if expectedRes == res {
		return res, "output matched !!!! right ans", nil
	} else {
		return -1, "", fmt.Errorf("out put does not match, error ")
	}
}
func main() {
	nums := "aabacbebebe"
	k := 3
	expectedRes := 7
	res := longestKSubstr(nums, k)

	result, msg, err := checkResult(expectedRes, res)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
		fmt.Println(msg)
	}
}
