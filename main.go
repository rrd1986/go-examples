package main

import (
	"fmt"
	"go-examples/src/algo"
)

//go env -w  GO111MODULE=off

func main() {
	unsortedList := []int{120, 56, 3, 5, 7, 8, 1, 6, 10}
	fmt.Println(unsortedList)
	sl := algo.Quick_sort(unsortedList)
	fmt.Println(sl)

	fibonacciList := algo.FibonacciSeries(20)
	fmt.Println(fibonacciList)

	algo.AlternateFibonacci(20)
}
