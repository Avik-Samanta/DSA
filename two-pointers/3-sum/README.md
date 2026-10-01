# 3 Sum
**link:** https://leetcode.com/problems/3sum/description/
**platform:** Leetcode
**question no.:** 15
**difficulty:** Medium


## Approach
Two Pointers

* Sort the array.
* Fix nums[i], then set l = i + 1 and r = n - 1.
* Find two numbers such that nums[l] + nums[r] = -nums[i].
* If sum is smaller → l++; if larger → r--.
* If equal → store triplet, move both pointers, and skip duplicates.
* Skip duplicate i values to avoid duplicate triplets.

## Complexity
Time: O(n^2 + nlogn) => O(n^2)
Space: O(1)
