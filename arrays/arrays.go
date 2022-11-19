package arrays

func SumArray(numbers []int)(sum int){
	for _, number := range(numbers){
		sum += number
	}
	return
}