package helpers

func ContainsValue(m map[string]string, target string) bool {
	for _, v := range m {
		if v == target {
			return true
		}
	}
	return false
}
