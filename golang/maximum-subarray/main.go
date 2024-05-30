package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, -2, 3, 10, -4, 20, 1, -3, 4}
	fmt.Println(maxSubArray(nums))

}

func maxSubArray(nums []int) int {
	return maxSubArrayHelper(nums, 0, len(nums)-1)
}

func maxSubArrayHelper(nums []int, left, right int) int {

	if left == right { // only contains one number
		return nums[left]
	}

	mid := left + (right-left)/2

	maxSumLeft := maxSubArrayHelper(nums, left, mid)
	maxSumRight := maxSubArrayHelper(nums, mid+1, right)

	maxSumCrossing := maxSumCrossing(nums, left, mid, right)

	return max(maxSumLeft, maxSumRight, maxSumCrossing)

}

func maxSumCrossing(nums []int, left, mid, right int) int {

	leftSum := math.MinInt64
	sum := 0

	for i := mid; i >= left; i-- {
		sum += nums[i]

		if sum > leftSum {
			leftSum = sum
		}
	}

	rightSum := math.MinInt64
	sum = 0

	for i := mid + 1; i <= right; i++ {

		sum += nums[i]

		if sum > rightSum {
			rightSum = sum
		}
	}

	return leftSum + rightSum
}

func max(nums ...int) int {

	maximum := nums[0]

	for _, num := range nums[1:] {
		if num > maximum {
			maximum = num
		}
	}

	return maximum
}
