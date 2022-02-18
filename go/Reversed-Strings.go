//https://www.codewars.com/kata/5168bb5dfe9a00b126000018/train/go
package kata

import (
	"strings"
)

func Solution(word string) string {
	var dizi []string
	str := strings.Split(word, "")

	for i := len(str) - 1; i >= 0; i-- {
		dizi = append(dizi, str[i])
	}
	kelime := strings.Join(dizi, "")

	return kelime
}
