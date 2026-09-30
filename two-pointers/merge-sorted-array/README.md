# Merge Sorted Array 
**link:** https://leetcode.com/problems/merge-sorted-array/description/
**platform:** Leetcode
**question no.:** 88
**difficulty:** Easy


## Approach
Two Pointers

First while:
Compare both arrays and place the larger element.

Second while:
If nums2 still has elements → copy them.

If nums1 still has elements → do nothing,
because they're already correctly placed.

## Complexity
Time: O(n)
Space: O(1)
