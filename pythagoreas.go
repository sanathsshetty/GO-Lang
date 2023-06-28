package main

import "fmt"

func findspecialpythogorastriple(n int) (int, int, int) {
	for a := 1; a < n/3; a++ {
		for b := a; b <= (n-a)/2; b++ {
			c := n - a - b
			if a*a+b*b == c*c {
				return a, b, c
			}
		}
	}
	return -1, -1, -1
}
func main() {
	n := 150
	a, b, c := findspecialpythogorastriple(n)
	if a != -1 && b != -1 && c != -1 {
		fmt.Printf("the special pythagoras triple for sum %d is(%d,%d,%d)\n", n, a, b, c)

	} else {
		fmt.Printf("No special pythagorean triple found sum %d \n", n)
	}
}
