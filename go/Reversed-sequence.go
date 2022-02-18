//https://www.codewars.com/kata/5a00e05cc374cb34d100000d/train/go
package kata

func ReverseSeq(n int) []int {
	var sayilar []int
	for i := n; i > 0; i-- {
		sayilar = append(sayilar, i)
	}
	return sayilar
}
