package main

import "fmt"

func main() {
	result := Pelindrone(121)
	fmt.Println(result)

}

func Pelindrone(x int) bool {
	//var num int
	// fmt.Scan(&num)
	// fmt.Println("Pelindrone nu:", num)
	if x < 0 {
		return false
	}
	num := x
	rev := 0
	for x != 0 {
		rem := x % 10
		x = x / 10
		rev = rev*10 + rem
	}
	if rev == num {
		return true
	}
	return false

}
