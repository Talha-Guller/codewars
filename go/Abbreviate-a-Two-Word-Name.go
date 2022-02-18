//https://www.codewars.com/kata/57eadb7ecd143f4c9c0000a3/train/go
package kata

import "strings"

func AbbrevName(name string) string {
	var kelimeler []string
	str2 := ""
	str := strings.Split(name, " ")
	if len(str) == 2 {
		for _, v := range str {
			array := strings.Split(v, "")
			kelimeler = append(kelimeler, strings.ToUpper(array[0]))
		}
		str2 = strings.Join(kelimeler, ".")
	}
	return str2
}
