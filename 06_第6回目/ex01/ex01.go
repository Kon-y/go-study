package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func main() {
	answer := plus(1, 2)
	fmt.Println(answer)

	var a = 1
	// 足し算
	a += 1
	fmt.Println(a) // 2

	//　引き算
	a -= 1
	fmt.Println(a) // 1

	//　掛け算
	a *= 2
	fmt.Println(a) // 2

	//　割り算
	a /= 2
	fmt.Println(a) // 1

}
