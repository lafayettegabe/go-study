package main

func Sum(arr []int) int {
	total := 0
	for _, value := range arr {
		total += value
	}
	return total
}

func SumAll(arr [][]int) []int {
	sums := make([]int, len(arr))
	for index := range sums {
		sums[index] = Sum(arr[index])
	}
	return sums
}

func SumAllTails(arr [][]int) []int {
	sums := make([]int, len(arr))
	for index := range sums {
		if len(arr[index]) == 0 {
			sums[index] = 0
		} else {
			sums[index] = Sum(arr[index][1:])
		}
	}
	return sums
}
