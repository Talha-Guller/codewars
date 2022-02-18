//https://www.codewars.com/kata/52597aa56021e91c93000cb0/train/go
package kata

func MoveZeros(arr []int) []int {
	var arr2 []int
	var dizi []int

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			arr2 = append(arr2, arr[i])
		} else {
			dizi = append(dizi, arr[i])
		}

	}
	dizi = append(dizi, arr2...)

	return dizi
}
