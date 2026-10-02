# Dutch National Flag(Sort Colors)
**link:** https://leetcode.com/problems/sort-colors/description/
**platform:** Leetcode  
**question no.:** 75  
**difficulty:** Medium  


## Approach
Two Pointers

* Set low = 0, mid = 0, high = n - 1
* While mid <= high:
    * If nums[mid] == 0:
        * Swap nums[mid] and nums[low]
        * low++, mid++
    * If nums[mid] == 1:
        * mid++
    * If nums[mid] == 2:
        * Swap nums[mid] and nums[high]
        * high--
        * Don’t increment mid
* Array is sorted when loop ends.

## Complexity
Time: O(n)  
Space: O(1)
