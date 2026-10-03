# Minimum Size Subarray Sum
**link:** https://leetcode.com/problems/minimum-size-subarray-sum/description/  
**platform:** Leetcode    
**difficulty:** Medium    


## Approach
Sliding Window

* Sort the array.
* Start a Loop nums[i], then set l = i + 1 and r = n - 1.
* Find three numbers such that nums[i] + nums[l] + nums[r] = sum.
  * If sum is smaller → sum < target -> then find all the possibilites like res += r - l and l++.
  * If larger or equal → sum >= target -> r--.
* Return the res which is all the posibilities of triplets with sum less than target.

## Complexity
Time: O(n)  
Space: O(1)
