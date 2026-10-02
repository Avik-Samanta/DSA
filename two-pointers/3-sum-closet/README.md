# 3 Sum Closet
**link:** https://leetcode.com/problems/3sum-closest/description/  
**platform:** Leetcode  
**question no.:** 16  
**difficulty:** Medium  


## Approach
Two Pointers

* Sort the array.
* Start a Loop nums[i], then set l = i + 1 and r = n - 1.
* Find three numbers such that nums[i] + nums[l] + nums[r] = sum.
* If sum is equal to target than return sum.
* If sum is smaller → l++ and find the diff of sum and target -> 
* For the lowest diff store the value of sum.
* if larger → r-- and find the diff of sum and target ->
* For the lowest diff store the value of sum.
* return the sum of the lowest diff between sum and target.

## Complexity
Time: O(n^2 + nlogn) => O(n^2)  
Space: O(1)
