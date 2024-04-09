package main

import "fmt"

func main() {
	sum := calculateSum(12, 13)
	fmt.Print("Total sum=", sum)
}

func calculateSum(first int, second int) int {
	return first + second
}

type Order struct {
	ID       string
	Date     int64
	City     string
	Amount   float64
	District string
}
