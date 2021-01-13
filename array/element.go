package array

// ErrorNotFound error
const ErrorNotFound = -1

// Get return element at given index of given array. If fails, return default value.
func Get(arr []string, index int, defaultValue string) string {
	if index < 0 || index >= len(arr) {
		return defaultValue
	}
	return arr[index]
}

// IndexOf returns index of pin in haystack array. If not found, return -1.
func IndexOf(haystack []string, pin string) int {
	for index, hay := range haystack {
		if hay == pin {
			return index
		}
	}
	return -1
}
