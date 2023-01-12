package main

import "fmt"

// 买卖股票的最好时机
func maxProfit(prices []int) int {
	// write code here
	var length int = len(prices)
	minPrice := prices[0]
	var totalRevenu int
	for i := 1; i < length; i++ {
		if prices[i] > minPrice {
			totalRevenu += (prices[i] - prices[i-1])
			minPrice += prices[i]
		} else {
			minPrice += prices[i]
		}
	}
	return totalRevenu
}
func minValue(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func minPathSum(matrix [][]int) int {
	// write code here
	var dp [][]int = make([][]int, len(matrix))
	var width int = len(matrix)
	var length int = len(matrix[0])
	for k := 0; k < width; k++ {
		dp[k] = make([]int, len(matrix[0]))
	}
	for i := 0; i < width; i++ {
		for j := 0; j < length; j++ {
			if i == 0 && j != 0 {
				dp[i][j] = dp[0][j-1] + matrix[0][j]
			} else if i != 0 && j == 0 {
				dp[i][j] = dp[i-1][0] + matrix[i][0]
			} else if i == 0 && j == 0 {
				dp[i][j] = matrix[i][j]
			} else {
				dp[i][j] = minValue(dp[i-1][j], dp[i][j-1]) + matrix[i][j]
			}
		}
	}
	return dp[width-1][length-1]
}
func main() {
	var matrix [][]int = [][]int{{1, 2, 3}, {1, 2, 3}}
	fmt.Println(minPathSum(matrix))
}

//func spiralOrder(matrix [][]int) []int {
//	// write code here
//	var length int = len(matrix)
//	var res []int
//	if length == 0 {
//		return res
//	}
//	var width int = len(matrix[0])
//	var visited [][]int
//	total := length * width
//	for cur := 0; cur < total; {
//		for i := 0; i < length; i++ {
//		}
//		for j := 0; j < width; j++ {
//		}
//		for i := 0; i < length; i++ {
//		}
//		for i := 0; i < length; i++ {
//		}
//	}
//	return res
//}
//
