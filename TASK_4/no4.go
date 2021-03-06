// Tulis program di Golang untuk menemukan nilai maksimum serta minimum di antara 6 angka inputan. Gunakan multiple return fungsi, pointer untuk referencing maupun deferencing!

package main

import (
	"fmt"
)

func getMinMax(numbers ...*int) (min int, max int) {
	min = *numbers[0]
	max = *numbers[0]
	for _, n := range numbers {
		if *n < min {
			min = *n
		}
		if *n > max {
			max = *n
		}
	}
	// input
	return
}

func main() {

	var a1, a2, a3, a4, a5, a6, min, max int

	fmt.Scan(&a1)

	fmt.Scan(&a2)

	fmt.Scan(&a3)

	fmt.Scan(&a4)

	fmt.Scan(&a5)

	fmt.Scan(&a6)

	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)

	fmt.Println("Nilai min ", min)

	fmt.Println("Nilai max ", max)

}
