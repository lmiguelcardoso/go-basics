package arraysandslices

func Sum(numbers []int) (sum int){
	for _,number := range numbers{
		sum += number	
	}
	return
}


func SumAll(numbersToSum ...[]int) []int {
	sums := []int{}

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}


func SumAllTails(numbersToSum ...[]int) []int{
	sums:= []int{}

	for _, numbers := range numbersToSum{
		tail := numbers[1:]
		sums = append(sums, tail...)
	}
	
	return sums
}