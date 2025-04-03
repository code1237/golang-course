package lesson02

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}

	dividers := 1

	for i := 2; i <= n; i++ {
		if n%i == 0 {
			if dividers++; dividers > 2 {
				return false
			}
		}
	}

	return true
}

func FibonacciIterative(n int) int {
	if n < 1 {
		panic("n must be greater than zero")
	}

	if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	}

	prev, current := 0, 1

	for i := 2; i < n; i++ {
		current = prev + current
		prev = current - prev
	}

	return current
}
