// BINARY SEARCH ARGO
// Dalam pencarian biner kami menggunakan informasi bahwa semua elemen diurutkan. Mari kita coba menyelesaikan tugas di mana kita meminta posisi nilai x dalam array yang diurutkan.
// Mari kita lihat bagaimana jumlah kandidat berkurang, misalnya untuk nilai x = 31.
// Untuk setiap langkah algoritma, kita harus mengingat awal dan akhir dari potongan array yang tersisa (masing-masing, variabel memohon dan akhir). Elemen tengah dari
// slice dapat dengan mudah dihitung sebagai mid = [(beg+end)/2].

package main

import "fmt"

func BinarySearch(array []int, x int) {

	// your code here
	l, r := 0, len(array)-1

	if x > array[len(array)-1] || x < array[0] {
		fmt.Println(-1)
		return
	} else if array[l] == x {
		fmt.Println(l)
		return
	} else if array[r] == x {
		fmt.Println(r)
		return
	} else {
		for l+1 != r {
			if array[(l+r)/2] == x {
				fmt.Println((l + r) / 2)
				return
			} else if array[(l+r)/2] < x {
				l = (l + r) / 2
			} else {
				r = (l + r) / 2
			}
		}
	}

	fmt.Println(-1)
	return
}

func main() {

	BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3) // 2

	BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5) // 3

	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53) // 6

	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100) // -1

}
