package main

import "fmt"

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
	fmt.Println(isValid("{[]{}}"))
	fmt.Println(isValid("{[)]}"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
	fmt.Println(isValid("{()}"))
	fmt.Println(isValid("{[()]}"))
	fmt.Println(isValid("{{[]}}"))
	fmt.Println(isValid("{{{[]}}}"))
}

func isValid(s string) bool {
	/**
	* The solution must use the Stack strategy adding the first part of the statement "(",  "[",  "{"
	* Otherwise if is not the first part of this statement validate if my stack does not contain any value it means that the input is invalid
	* e.g. ")[]" so we should return false.
	* if my stack is not empty we will start the validation process and remove the elements from stack that we already validated.
	 */

	stack := []rune{}

	for _, symbol := range s {
		if symbol == '(' || symbol == '[' || symbol == '{' {
			stack = append(stack, symbol) // adding to stack
		} else {
			if len(stack) == 0 { // does not contain any element then it should be an invalid input
				return false
			}

			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // Simulate the pop function
			comb := string(last) + string(symbol)
			fmt.Printf("Combination %v ", comb)
			if !isValidCombination(comb) {
				return false
			}
		}
	}

	return len(stack) <= 0
}

func isValidCombination(comb string) bool {
	if comb == "()" || comb == "[]" || comb == "{}" {
		return true
	}
	return false
}
