//https://www.codewars.com/kata/54ff3102c1bad923760001f3/train/go

package kata

import "strings"

func GetCount(str string) (count int) {
	count = 0
	word := []string{"a", "e", "i", "o", "u"}
	for i := 0; i < len(word); i++ {
		sonuc := strings.Split(str, "")
		for j := 0; j < len(sonuc); j++ {
			varMi := strings.EqualFold(strings.ToLower(string(sonuc[j])), word[i])
			if varMi {
				count = count + 1
			}
		}

	}
	return count
}
