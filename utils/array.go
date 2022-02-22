package utils

func StrArrayContains(arr []string, str string) bool {
	for _, ele := range arr {
		if ele == str {
			return true
		}
	}
	return false
}
