package main

import (
	"fmt"
	"math"
)

func totalFruit(fruits []int) int {
	low, basket, res := 0, 2, math.MinInt
	mp := make(map[int]int)

	for high := range fruits {
		mp[fruits[high]]++

		for len(mp) > basket {
			mp[fruits[low]]--
			if mp[fruits[low]] == 0 {
				delete(mp, fruits[low])
			}
			low++
		}
		fruitLen := high - low + 1
		res = max(res, fruitLen)
	}
	return res
}

func checkError(expectedRes, result int) (int, string, error) {
	if expectedRes == result {
		return result, "right output !!!", nil
	} else {
		return -1, "", fmt.Errorf("wrong ans !!! , check the code")
	}
}

func main() {
	fruits := []int{1, 2, 3, 2, 2}
	expectedRes := 4
	result := totalFruit(fruits)

	res, msg, err := checkError(expectedRes, result)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		fmt.Println(msg)
	}
}
