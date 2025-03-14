package util

func ArrContain[T comparable](arr []T, key T) bool {
	for _, item := range arr {
		if item == key {
			return true
		}
	}
	return false
}

func SplitArr[T any](arr []T, num int) [][]T {
	length := len(arr)
	if length <= num {
		return [][]T{arr}
	}

	quantity := (length + num - 1) / num
	segments := make([][]T, 0, quantity)

	for i := 0; i < quantity; i++ {
		end := (i + 1) * num
		if end > length {
			end = length
		}

		segment := arr[i*num : end]
		segments = append(segments, segment)
	}

	return segments
}

func Subtract(slice1, slice2 []string) []string {
	m := make(map[string]bool)
	for _, v := range slice1 {
		m[v] = true
	}

	var result []string
	for _, v := range slice2 {
		if !m[v] {
			result = append(result, v)
		}
	}
	return result
}
