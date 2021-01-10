package text

// StartWith return true if s starts with m
func StartWith(s string, m string) bool {
	if len(s) < len(m) {
		return false
	}
	if s[:len(m)] == m {
		return true
	}
	return false
}
