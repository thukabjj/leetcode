package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	printArray(nums)
}

func printArray(nums []int) {
	for _, num := range nums {
		fmt.Printf("%d \n", num)
	}
}

func moveZeroes(nums []int) {
	if len(nums) == 1 {
		return
	}

	left := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] != 0 {
			nums[left] = nums[right]
			left++
		}

	}
	for ; left < len(nums); left++ {
		nums[left] = 0
	}
}
