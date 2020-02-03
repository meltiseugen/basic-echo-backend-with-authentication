package utils

func Contains(s []string, p string) bool {
	for _, n := range s {
		if p == n {
			return true
		}
	}
	return false
}