package main

import (
	"fmt"
)

/*
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
示例 1：

输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
示例 2：

输入：nums = [0]
输出：[[],[0]]

链接：https://leetcode.cn/leetbook/read/top-interview-questions-medium/xv67o6/
*/

var result [][]int

func subsets(nums []int) [][]int {

	traceback(nums, 0)
	return result
}

func traceback(nums []int, index int) {

	if index == len(nums) {
		return
	}
	var next int
	for next < len(result) {
		result[next] = append(result[next], nums[index])
		next++
	}
	traceback(nums, index+1)
}

func main() {

	fmt.Println(len([][]int{}))
}
