/*
* https://leetcode.com/problems/two-sum/
*
* Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
* You may assume that each input would have exactly one solution, and you may not use the same element twice.
* You can return the answer in any order.
*
* Constraints:
*  2 <= nums.length <= 104
*  -109 <= nums[i] <= 109
*  -109 <= target <= 109
*  Only one valid answer exists.
*
 */

package main

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i1, x := range nums {
		diff := target - x

		if i2, ok := hashMap[diff]; ok {
			return []int{i1, i2}
		}

		hashMap[x] = i1
	}
	return []int{}
}
