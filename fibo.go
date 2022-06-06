package fibo

var m map[int]int

// for loop fibonacci
func fibo0(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}
	one := 1
	two := 0
	rst := 0
	for i := 2; i <= n; i++ {
		rst = one + two
		two = one
		one = rst
	}
	return rst

}

// for loop fibonacci  golang  type
func fibo1(n int) int {
	n1, n2 := 1, 1

	for i := 0; i < n-2; i++ {
		n1, n2 = n2, n1+n2
	}

	return n2
}

// recursive fibonacci algorithm
func fibo2(n int) int {
	if n < 0 {
		return 0

	}
	if n < 2 {
		return n
	}
	return fibo2(n-1) + fibo2(n-2)

}

// refactoring memoization algorithm
func fibo3(n int) int {
	if n < 0 {
		return 0

	}
	if n < 2 {
		return n
	}
	if m == nil {
		m = make(map[int]int)
	}
	if v, ok := m[n]; ok {
		return v
	}
	v := fibo2(n-1) + fibo2(n-2)
	m[n] = v
	return v

}
