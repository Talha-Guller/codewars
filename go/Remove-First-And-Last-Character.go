//https://www.codewars.com/kata/56bc28ad5bdaeb48760009b0/train/go
package kata

import (
	"strings"
)

func RemoveChar(word string) string {
	str := ""
	kelime := strings.Split(word, "")
	for i := 0; i < len(kelime); i++ {
		if i == 0 {
			continue
		} else if i == len(kelime)-1 {
			continue
		} else {
			str = str + kelime[i]
		}
	}
	return str
}
