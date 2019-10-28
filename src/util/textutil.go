package util

// EmptyPassword read no password
func EmptyPassword(string) (string, error) {
	return "", nil
}
