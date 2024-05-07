package recursion

// 迭代实现阶乘
type Fac struct {
	value map[int]int
}

func NewFactorial(n int) *Fac {
	return &Fac{
		make(map[int]int, n),
	}
}

func (f *Fac) Factorial(n int) int {
	if f.value[n] != 0 {
		return f.value[n]
	}
	if n <= 1 {
		f.value[n] = 1
		return 1
	} else {
		res := n * f.Factorial(n-1)
		f.value[n] = res
		return res
	}
}

func (f *Fac) Print(n int) {
	println(f.value[n])
}
