package utils

func Sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func Last[T any](array []T, count int) []T {
	result := []T{}

	if count < 1 {
		return result
	}
	return array[len(array)-count:]
}
