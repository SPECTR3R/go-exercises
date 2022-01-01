package intSlices

func GetItem(slice []int, index int) (int, bool) {
	if index >= 10 {
		return 0, false
	}
	return slice[index], true
}

func SetItem(slice []int, index, value int) []int {
	if index < 0 || index >= len(slice) {
		slice = append(slice, value)

		return slice
	}
	slice[index] = value
	return slice
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	slice := make([]int, length)

	if length < 0 {
		return []int{}
	}

	for i := 0; i < length; i++ {
		slice[i] = value
	}
	return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}
