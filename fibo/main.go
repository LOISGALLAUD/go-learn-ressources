package main

import (
	"fmt"
)

func recursiveFibonacci(n int) (result int) {
	if n <= 1 {
		return n
	}
	result = recursiveFibonacci(n-1) + recursiveFibonacci(n-2)
	fmt.Println(result)
	return
}

func main() {
	var (
		previous int = 0
		current int = 1
		next int
		count int = 10
	)

	fmt.Print("Enter a integer: ")
	fmt.Scan(&count)
	fmt.Printf("Fibonacci Series up to %d terms\n", count)

	for i := 0; i < count; i++ {
		fmt.Println(previous)
		next = previous + current
		previous = current
		current = next
	}

	fmt.Println("Recursive Fibonacci")
	recursiveFibonacci(count)
}