package sieve

import "sort"

func Sieve(limit int) []int {
	var primes []int
	mapping := make(map[int]int)

	// 1. discard a limit value of 1
	if limit < 2 {
		return primes
	}

	// 2. store all numbers up to limit in a map and assign value of "1"
	for n := 2; n <= limit; n++ {
		mapping[n] = 1
	}

	// 3. double loop to form calculations, then set non-primes to value of "0"
	for k := 2; k < limit; k++ {
		for v := 2; v < limit; v++ {
			check := k * v
			if check <= limit {
				mapping[check] = 0
			}
		}
	}

	// 4. filter out the primes from mapping (value of "1") and store those in the primes array
	for k, v := range mapping {
		if v == 1 {
			primes = append(primes, k)
		}
	}
	// 5. sort the order of primes
	sort.Ints(primes)
	return primes
}
