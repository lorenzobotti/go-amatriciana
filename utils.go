package amatriciana

func stringContains(rawInput string, contains rune) bool {
	input := []rune(rawInput)
	for _, char := range input {
		if char == contains {
			return true
		}
	}

	return false
}

func maxFloat(inputs ...float64) float64 {
	max := inputs[0]

	for _, num := range inputs {
		if num > max {
			max = num
		}
	}

	return max
}

func minFloat(inputs ...float64) float64 {
	min := inputs[0]

	for _, num := range inputs {
		if num < min {
			min = num
		}
	}

	return min
}
