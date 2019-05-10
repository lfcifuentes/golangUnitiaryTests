package main

import "fmt"

func Sum(x int, y int) int {
	return x + y
}

func main() {
	res := Sum(5, 5)
	fmt.Printf("Response: %d",res)
}