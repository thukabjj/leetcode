package main

import "fmt"

func main() {
	printClimbingStairs(2)
	printClimbingStairs(3)
	printClimbingStairs(4)
	printClimbingStairs(5)
	printClimbingStairs(6)
}

func printClimbingStairs(n int) {
	fmt.Println(climbStairs(n))
}

func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	prev := 1
	current := 1
	for i := 2; i <= n; i++ {
		temp := current
		current += prev
		prev = temp
	}
	return current
}
