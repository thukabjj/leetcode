package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 2, 3, 4}
	fmt.Println(singleNumberSorted(nums))
	fmt.Println(singleNumberXOR(nums))

}

func singleNumberSorted(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i += 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}
func singleNumberXOR(nums []int) int {
	var result int
	for _, num := range nums {
		result ^= num
	}
	return result
}
