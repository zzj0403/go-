package recursion

import "fmt"

// Fibonacci 斐波那契数列
type Fibs struct {
	val map[int]int
}

func NewFibs(n int) *Fibs {
	return &Fibs{val: make(map[int]int, n)}
}

func (f *Fibs) Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	if _, ok := f.val[n]; !ok {
		f.val[n] = f.Fibonacci(n-1) + f.Fibonacci(n-2)
	}
	return f.val[n]
}

func (fib *Fibs) Print(n int) {
	fmt.Println(fib.val[n])
}
