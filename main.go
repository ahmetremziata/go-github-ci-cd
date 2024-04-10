package main

import "fmt"

func main() {
	sum := calculateSum(12, 13)
	fmt.Print("Total sum=", sum)
}

func calculateSum(first int, second int) int {
	return first + second
}
