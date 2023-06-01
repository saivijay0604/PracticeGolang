package main

import "fmt"

//Given an integer n, return true if it is a power of two. Otherwise, return false.

//An integer n is a power of two, if there exists an integer x such that n == 2x.

func main() {
	n := 16
	fmt.Println(isPowerOfTwo(n))

}
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	for n%2 == 0 {
		n = n / 2
	}
	if n == 1 {
		return true
	}
	return false
}
