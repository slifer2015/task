package common

func StringInArray(s string, arr []string) bool {
	for i := range arr {
		if s == arr[i] {
			return true
		}
	}
	return false
}
