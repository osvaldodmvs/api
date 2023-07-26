package utils

import "strings"

func IsValidCategory(submittedCategory string) bool {
	validCategories := []string{
		"electronics",
		"books",
		"food",
		"clothing",
		"tools",
		"other",
	}

	for _, validCategory := range validCategories {
		if strings.EqualFold(string(validCategory), string(submittedCategory)) {
			return true
		}
	}

	return false
}
