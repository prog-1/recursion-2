package main

func gcd(a, b int) int {
	if a == b {
		return a
	} else if a > b {
		return gcd(b, a-b)
	} else /*if b > a*/ {
		return gcd(a, b-a)
	}
}

// func main() {
// 	var a, b int
// 	fmt.Scanln(&a, &b)
// 	fmt.Println(gcd(a, b))
// }
