package main

import (
	"fmt"
	"strconv"
)

func ispalindrome(n int) bool {
	str := strconv.Itoa(n)
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-i-1] {
			return false
		}

	}
	return true
}
func largestpalindromeproduct() (int, int, int) {
	largestpalindrome := 0
	var multiplicand1, multiplicand2 int
	for i := 99; i >= 10; i-- {
		for j := i; j >= 10; j-- {
			product := i * j
			if product < largestpalindrome {
				break
			}
			if ispalindrome(product) && product > largestpalindrome {
				largestpalindrome = product
				multiplicand1 = i
				multiplicand2 = j
			}
		}
	}
	return largestpalindrome, multiplicand1, multiplicand2
}
func main() {
	result, multiplicand1, multiplicand2 := largestpalindromeproduct()
	fmt.Printf("the largest palindrome product is:%d\n", result)
	fmt.Printf("the multiplicands are :%d and %d\n", multiplicand1, multiplicand2)
}
