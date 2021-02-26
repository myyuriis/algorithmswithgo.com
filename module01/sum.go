package module01

// Sum will sum up all of the numbers passed
// in and return the result

/*
Solved using normal way
func Sum(numbers []int) int {
	total := 0
	if len(numbers) == 0 {
		return 0
	}

	for _, number := range numbers {
		total = total + number
	}

	return total
}
*/

// Solved using recursive function
func Sum(numbers []int) int {
	// This is the base case
	if len(numbers) == 0 {
		return 0
	}

	return numbers[0] + Sum(numbers[1:])
}

/*
Explanation for the recursive function
Sum(3, 4, 5) =>
	3 + Sum(4, 5) =>
		4 + Sum(5) =>
			5 + Sum() =>
				0 (This is the base case)
*/