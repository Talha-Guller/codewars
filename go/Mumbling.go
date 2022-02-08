//https://www.codewars.com/kata/5667e8f4e3f572a8f2000039/train/go
package kata

import (
	"strings"
)

func Accum(s string) string {

	str := make([]string, len(s))

	for i := 0; i < len(s); i++ {
		str[i] = strings.ToUpper(string(s[i])) + strings.ToLower(strings.Repeat(string(s[i]), i))
	}

	return strings.Join(str, "-")
}
