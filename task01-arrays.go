package homework

func average(input [15]float32) (result float32) {
	var sum float32
	sum = 0
	for i := 0; i < len(input); i++ {
		sum += input[i]
	}
	result = sum / float32(len(input))
	return result

}
