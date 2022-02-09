//https://www.codewars.com/kata/5a4d303f880385399b000001/train/go
package kata

import (
	"strconv"
	"strings"
)

func Strong(n int) string {
	var fak int
	var sonuc int
	dönüşüm := strconv.FormatInt(int64(n), 10)
	for i := 0; i < len(dönüşüm); i++ {
		deger := strings.Split(dönüşüm, "")
		sayi, _ := strconv.ParseInt(deger[i], 10, 64)
		k := 2
		fak = 1
		for k <= int(sayi) {
			fak = k * fak
			k++
		}
		sonuc = sonuc + fak
	}
	if sonuc == n {
		return "STRONG!!!!"
	}
	return "Not Strong !!"
}
