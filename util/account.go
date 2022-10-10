package util

func IsAdmin(roleNames []string) bool {
	if roleNames == nil || len(roleNames) == 0 {
		return false
	}
	for _, item := range roleNames {
		if item == "admin" {
			return true
		}
	}
	return false
}
