//https://www.codewars.com/kata/576bb71bbbcf0951d5000044/train/go
package kata

func CountPositivesSumNegatives(numbers []int) []int {
	var res []int
	sayac := 0
	total2 := 0
	for _, v := range numbers {
		if v > 0 {
			sayac = sayac + 1
		} else if v < 0 {
			total2 = total2 + v
		}
	}
	res = append(res, sayac)
	res = append(res, total2)

	return res // your code here
}
