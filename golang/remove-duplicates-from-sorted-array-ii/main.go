package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))

}

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 0 {
		return 0
	}

	fast := 1
	for slow := 1; slow < len(nums); slow++ {
		if fast == 1 || nums[slow] != nums[fast-2] {
			nums[fast] = nums[slow]
			fast++
		}
	}
	fmt.Printf("Array Final Order %v\n", nums)
	return fast
}
