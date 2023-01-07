package validation

func IsListEmpty(list []string) bool {
	if len(list) == 0 {
		return true
	}
	return false
}
