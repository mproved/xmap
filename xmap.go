package xmap

func Invert[K comparable, V comparable](m map[K]V) (output map[V]K) {
	result := make(map[V]K)

	for k, v := range m {
		result[v] = k
	}

	return result
}

func Merge[T comparable, U any](maps ...map[T]U) map[T]U {
	result := make(map[T]U)

	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}

	return result
}
