package kitchenSink

/**
Takes in a list of numbers and returns the sum
*/
func Sum(numbers []int) int {

	sum := 0
	for _, n := range numbers {
		sum += n
	}

	return sum
}
