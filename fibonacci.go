package main

import "fmt"

func FibonacciClosure() func() int {
	var prev0, prev1 int = 0, 1
	var curr int = 0
	return func() int {
		var result = curr
		prev0 = prev1
		prev1 = curr
		curr = prev0 + prev1
		return result
	}
}

func Fibonacci() {
	f := FibonacciClosure()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
