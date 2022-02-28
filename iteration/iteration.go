package iteration

import "strings"

func Repeat(v string, count int) string {
	result := strings.Repeat(v, count)
	return result
}
