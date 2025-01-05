package tools

func RemoveElementFromSlice[T comparable](slice []T, item T) []T {
	for i, e := range slice {
		if e == item {
			return append(slice[:i], slice[i+1:]...)
		}
	}

	return slice
}
