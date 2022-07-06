package main

import "fmt"

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func findMax(nums []int, k int) int {

	// Checking array length
	if len(nums) < k {
		fmt.Println("Length of array is smaller than k")
		return -1
	}

	sum := 0
	// Initiated maximum sum of K size
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	ans := sum
	// Checking sum after kth index
	for i := k; i < len(nums); i++ {
		ans += nums[i] - nums[i-k]
		sum = Max(sum, ans)
	}

	return sum
}

func main() {
	arr := []int{2, 1, 3, 4, 5, 6}
	k := 3

	fmt.Println(findMax(arr, k))
}
