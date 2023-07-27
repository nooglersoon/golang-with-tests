package main

func Sum(numbers []int) int {

	var sum int = 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

// Variadic function that can take a variable number of arguments
func SumAll(numbersToSum ...[]int) []int {

	lengthOfNumbers := len(numbersToSum)

	// Create allSum slices
	allSum := make([]int, lengthOfNumbers)

	for index, numbers := range numbersToSum {
		allSum[index] = Sum(numbers)
	}

	// Alternative, since slice can go through without worrying less about its capacity

	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {

	var sums []int

	for _, numbers := range numbersToSum {

		var tail []int

		// Slicing slice/array from given index to the end of array
		if len(numbers) > 0 {
			tail = numbers[1:]
		}
		sums = append(sums, Sum(tail))
	}

	return sums

}
