//https://www.codewars.com/kata/585a1a227cb58d8d740001c3/train/go

package kata

import "strings"

func Repeater(s string, n int) string {
	sonuc := strings.Repeat(s, n)
	return sonuc
}
