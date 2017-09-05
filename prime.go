package main

import "math"

func isPrime(x int) int {
	if x < 2 {
		return -1
	}

	if x < 4 {
		return 1
	}

	if math.Mod(float64(x), 2) == 0 {
		return 0
	}

	for i := 3; float64(i) < math.Floor(math.Sqrt(float64(x))); i += 2 {
		if math.Mod(float64(x), float64(i)) == 0 {
			return 0
		}
	}

	return 1
}

func nextPrime(x int) int {
	for isPrime(x) != 1 {
		x++
	}

	return x
}
