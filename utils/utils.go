package utils

func FindItemInList(items []string, searchItem string) bool {
	found := false
	for _, item := range items {
		if item == searchItem {
			found = true
		}
	}
	return found
}
