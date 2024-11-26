package algorithms

func InList(arr []string, key string) bool {
	for _, s := range arr {
		if s == key {
			return true
		}
	}
	return false
}
