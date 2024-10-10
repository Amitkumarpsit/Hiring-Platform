// backend/utils/helpers.go
package utils

// Example utility function
func StringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
