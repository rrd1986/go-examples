package algo

import (
	"fmt"
)

func Fibonacci(n int) int {

	m := make(map[int]int)
	m[1] = 0
	m[2] = 1

	if n <= 1 {
		return m[1]
	} else if n == 2 {
		return m[2]
	} else if value, ok := m[n]; ok {
		return value
	}
	m[n] = Fibonacci(n-2) + Fibonacci(n-1)
	return m[n]

}

func FibonacciSeries(n int) (s []int) {

	for i := 1; i < n; i++ {
		s = append(s, Fibonacci(i))
	}
	return
}

func AlternateFibonacci(n int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
		fmt.Println(y)
	}

}
