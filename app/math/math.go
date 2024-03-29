package math

func Sum(numbers []int) int {
	result := 0
	for _, num := range numbers {
		result += num
	}

	return result
}

func Sub(numbers []int) int {
	result := 0
	counter := 0

	for _, num := range numbers {
		if counter == 0 {
			result = num
		} else {
			result -= num
		}

		counter++
	}

	return result
}

func Mul(numbers []int) int {
	result := 1
	for _, num := range numbers {
		result *= num
	}

	return result
}

func Div(numbers []int) float64 {
	result := 0.0
	counter := 0

	for _, num := range numbers {
		if counter == 0 {
			result = float64(num)
		} else {
			result /= float64(num)
		}

		counter++
	}

	return result
}

func And(numbers []int) int {
	result := 0
	counter := 0

	for _, num := range numbers {
		if counter == 0 {
			result = num
		} else {
			result &= num
		}

		counter++
	}

	return result
}

func Or(numbers []int) int {
	result := 0
	counter := 0

	for _, num := range numbers {
		if counter == 0 {
			result = num
		} else {
			result |= num
		}

		counter++
	}

	return result
}

func Xor(numbers []int) int {
	result := 0
	counter := 0

	for _, num := range numbers {
		if counter == 0 {
			result = num
		} else {
			result ^= num
		}

		counter++
	}

	return result
}
