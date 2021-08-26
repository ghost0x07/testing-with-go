package iteration

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(sliceOfNumbers ...[]int) []int {
	sum := []int{}

	for _, numbers := range sliceOfNumbers {

		sum = append(sum, Sum(numbers))
	}

	return sum
}
