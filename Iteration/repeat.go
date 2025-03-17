package iteration

import "strings"

func Repeat(character string, iterations int) string {
	var repeated strings.Builder
	for range iterations {
		repeated.WriteString(character)
	}
	return repeated.String()
}

func Fib(n int) int {
	if n < 2 {
		return n
	}

	a, b := 0, 1
	for i := 2; i < n; i++ {
		a, b = b, a + b
	}
	return a + b
}