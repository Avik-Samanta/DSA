# Triplets with Smaller Sum
**link:** https://www.geeksforgeeks.org/problems/count-triplets-with-sum-smaller-than-x5549/1  
**platform:** GeeksforGeeks  
**difficulty:** Medium


## Approach
Two Pointers

* Sort the array.
* Start a Loop nums[i], then set l = i + 1 and r = n - 1.
* Find three numbers such that nums[i] + nums[l] + nums[r] = sum.
  * If sum is smaller → sum < target -> then find all the possibilites like res += r - l and l++.
  * If larger or equal → sum >= target -> r--.
* Return the res which is all the posibilities of triplets with sum less than target.

## Complexity
Time: O(n^2 + nlogn) => O(n^2)  
Space: O(1)

