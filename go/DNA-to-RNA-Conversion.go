//https://www.codewars.com/kata/5556282156230d0e5e000089/train/go
package kata

import (
	"strings"
)

func DNAtoRNA(dna string) string {
	dna = strings.ToUpper(dna)
	str := strings.Split(dna, "")
	var dizi []string

	for _, v := range str {
		str2 := kontrol(v)
		dizi = append(dizi, str2)
	}
	dna2 := strings.Join(dizi, "")

	return dna2
}
func kontrol(str string) string {
	if str == "G" {
		str = "G"
	} else if str == "C" {
		str = "C"
	} else if str == "A" {
		str = "A"
	} else if str == "T" {
		str = "U"
	}
	return str
}
