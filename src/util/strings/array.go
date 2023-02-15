package strings

// if your array is over heavy, you should use this function
func ComplexDepduplicationStringArray[T comparable](array [][]T) [][]T {
	if len(array) == 0 {
		return array
	}

	var btree = NewStringBTree[T]()
	for _, v := range array {
		btree.Insert(v)
	}

	var result [][]T
	btree.Walk(func(t []T) {
		result = append(result, t)
	})

	return result
}

// if your array is not over heavy, you should use this function
func DepduplicationStringArray[T comparable](array []T) []T {
	if len(array) == 0 {
		return array
	}

	var result []T
	search := func(t T) bool {
		for _, v := range result {
			if v == t {
				return true
			}
		}
		return false
	}

	for _, v := range array {
		if !search(v) {
			result = append(result, v)
		}
	}

	return result
}

func Filter[T any](array []T, filter func(T) bool) []T {
	var result []T
	for _, v := range array {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}
