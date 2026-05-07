package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{2, 1, 4, 3, 10, 6, 9, 8, 7, 5}
	M := 15

	solutions := make(chan []bool)
	go subsetSum(numbers, M, solutions)

	counter := 1
	for solution := range solutions {
		fmt.Printf("S#%d: ", counter)
		for i, chosen := range solution {
			if chosen {
				fmt.Printf("%d ", numbers[i])
			}
		}
		counter++
		fmt.Println()
	}
}

func subsetSum(numbers []int, M int, solutions chan []bool) {
	// Preprocessing (O(nlogn))
	sort.Ints(numbers)

	n := len(numbers)

	// Preprocessing O(n)
	suffixSum := make([]int, n+1)
	suffixSum[n] = 0
	for i := n - 1; i >= 0; i-- {
		suffixSum[i] = suffixSum[i+1] + numbers[i]
	}

	include := make([]bool, n)
	for i := range include {
		include[i] = false
	}

	var backtrack func(weight, k int)
	backtrack = func(weight, k int) {
		if weight == M {
			solution := make([]bool, n)
			copy(solution, include)
			solutions <- solution
			return
		}

		if weight > M || k == n {
			return
		}

		if weight+suffixSum[k] < M {
			return
		}

		backtrack(weight, k+1)
		if weight+numbers[k] <= M {
			include[k] = true
			backtrack(weight+numbers[k], k+1)
			include[k] = false
		}
	}

	backtrack(0, 0)

	close(solutions)
}
