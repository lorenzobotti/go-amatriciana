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
