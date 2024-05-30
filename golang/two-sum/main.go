package main

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	println(result)
}

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, num := range nums {
		numberToFind := target - num

		if hashMapIndex, ok := hashMap[numberToFind]; ok {
			return []int{hashMapIndex, i}
		}
		hashMap[num] = i
	}
	return []int{}
}
