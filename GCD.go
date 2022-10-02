package main

import "fmt"

func mainGCD() {
	fmt.Println(GCD1(1000, 665))
	fmt.Println(GCD2(665, 1000))
}

//greatest Common Divisor
func GCD1(a, b int) int {
	if a == b {
		return a
	}
	if a > b {
		return GCD1(a-b, b)
	}
	return GCD1(b-a, a)
}

func GCD2(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD2(b, a%b)
}
