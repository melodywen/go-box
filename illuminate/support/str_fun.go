package support

import (
	"regexp"
	"strings"
)

type StrFun struct {
}

// Is Determine if a given string matches a given pattern.
func (_ StrFun) Is(pattern string, value string) bool {
	if pattern == "" {
		return false
	}
	if pattern == value {
		return true
	}
	// Asterisks are translated into zero-or-more regular expression wildcards
	// to make it convenient to check if the strings starts with the given
	// pattern such as "library/*", making any string check convenient.
	pattern = strings.ReplaceAll(pattern, "*", ".*")
	match, _ := regexp.MatchString("^"+pattern+"$", value)
	return match
}
