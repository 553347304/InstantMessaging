package conv

func SliceMap[T comparable](s []T) map[T]bool {
	m := make(map[T]bool)
	for _, v := range s {
		m[v] = true
	}
	return m
}
