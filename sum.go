package main

// Use func Sum(numbers [size]int) int {} for array

// Sum ... func to calculate sum of slice
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll ... Func to sum 2 slices
func SumAll(NumbersToSum ...[]int) []int {
	// lengthOfNumbers := len(NumbersToSum)
	// fmt.Println(NumbersToSum)
	// sums := make([]int, lengthOfNumbers)
	var sums []int
	for _, numbers := range NumbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

// SumAllTails ...
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tails := numbers[1:]
			sums = append(sums, Sum(tails))
		}
	}
	return sums
}
