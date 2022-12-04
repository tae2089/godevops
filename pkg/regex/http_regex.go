// Package regex  is regex function list package
package regex

import (
	"log"
	"regexp"
)

// IsHttpsUrl returns bool about https matched
func IsHttpsUrl(url string) bool {
	matched, err := regexp.MatchString("https:\\/\\/+[A-Za-z.]+", url)
	if err != nil {
		log.Println(err)
		return false
	}
	return matched
}
