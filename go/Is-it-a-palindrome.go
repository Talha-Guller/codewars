//https://www.codewars.com/kata/57a1fd2ce298a731b20006a4/train/go
package kata

import (
	"strings"
)

func IsPalindrome(str string) bool {
	birlestir := ""
	str2 := strings.ToLower(str)
	bak := strings.Split(str2, "")
	for i := len(bak) - 1; i >= 0; i-- {
		birlestir = birlestir + bak[i]
	}
	if str2 == birlestir {
		return true
	}

	return false
}
