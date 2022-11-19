package arrays

func SumArray(numbers []int)(sum int){
	for _, number := range(numbers){
		sum += number
	}
	return
}

func SumAll(firstArray []int, secondArray []int)(arraySum [2]int){
	firstSum := 0
	secondSum := 0
	for _, number := range(firstArray){
		firstSum += number
	}

	for _, number := range(secondArray){
		secondSum += number
	}

	arraySum[0] = firstSum
	arraySum[1] = secondSum
	return
}