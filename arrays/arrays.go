package arrays

func SumArray(numbers []int)(sum int){
	for _, number := range(numbers){
		sum += number
	}
	return
}

func SumAll(numberToSum...[]int) []int{
	// lengthOfNumbers := len(numberToSum)
	// sums := make([]int, lengthOfNumbers) // make a slice of integer type of length.
	var sums []int
	for _, arrayItem := range(numberToSum){
		sums = append(sums, SumArray(arrayItem))
	}
	return sums
}

func SumAllTails(arraysToSum... []int) []int{
	var sums []int
	for _, array:= range(arraysToSum){
		if len(array) != 0{
			tailArray := array[1:] // slice the slice from index 1 to last. Skipping head.
			sums = append(sums, SumArray(tailArray))
		} else {
			sums = append(sums, 0)
		}
	}

	return sums
}