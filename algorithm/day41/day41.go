package main

import (
	"fmt"
	"strconv"
)

//积跬步，至千里。
/*
输入一个非负整数数组numbers，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。
例如输入数组[3，32，321]，则打印出这三个数字能排成的最小数字为321323。
1.输出结果可能非常大，所以你需要返回一个字符串而不是整数
2.拼接起来的数字可能会有前导 0，最后结果不需要去掉前导 0
*/
//JZ45 把数组排成最小的数

func getPos1(nums []int, first, last int) int {

	var key = nums[first]
	for first < last {
		for first < last && verify(key, nums[last]) {
			last--
		}
		nums[first], nums[last] = nums[last], nums[first]
		for first < last && verify(nums[first], key) {
			first++
		}
		nums[first], nums[last] = nums[last], nums[first]
	}
	return first
}
func verify(a, b int) bool {
	var as = strconv.Itoa(a)
	var bs = strconv.Itoa(b)
	if as+bs <= bs+as {
		return true
	}
	return false
}
func getPos(nums []int, first, last int) int {

	var key = nums[first]
	for first < last {
		for first < last && key <= nums[last] {
			last--
		}
		nums[first], nums[last] = nums[last], nums[first]
		for first < last && key > nums[first] {
			first++
		}
		nums[first], nums[last] = nums[last], nums[first]
	}
	return first
}
func QuickSort(nums []int, first, last int) {
	if first >= last {
		return
	}
	var pos = getPos1(nums, first, last)
	QuickSort(nums, first, pos-1)
	QuickSort(nums, pos+1, last)
}

func main() {
	var nums = []int{3, 30, 34, 5, 9}
	//输出: "3033459"
	QuickSort(nums, 0, len(nums)-1)
	PrintMinNumber(nums)
	var nums1 = []int{10, 2}
	QuickSort(nums1, 0, 1)
	PrintMinNumber(nums1)
}

func PrintMinNumber(numbers []int) string {
	// write code here
	var ans string
	for _, v := range numbers {
		ans += strconv.Itoa(v)
	}
	fmt.Println(ans)
	return ans
}
