package collections

// CopyMap creates and return copy of the original map
func CopyMap[K comparable, V any](originalMap map[K]V) map[K]V {
	result := make(map[K]V, len(originalMap))
	for k, v := range originalMap {
		result[k] = v
	}
	return result
}
