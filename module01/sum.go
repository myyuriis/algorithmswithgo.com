package module01

// Sum will sum up all of the numbers passed
// in and return the result
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
