package arraysslices

func Sum(numbers []int) (sum int) {
	for _, num := range numbers {
		sum += num
	}
	return
}

func SumAll(numbersToSum ...[]int) (result []int) {
	for _, numbers := range numbersToSum {
		result = append(result, Sum(numbers)) 
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (tailResults []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			tailResults = append(tailResults, 0)
		} else {
			tailSum := Sum(numbers[1:])
			tailResults = append(tailResults, tailSum)
		}
	}
	return
}