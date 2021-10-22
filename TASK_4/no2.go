// Tulis metode yang mengambil `offset` integer dan string. Menghasilkan string baru, di mana setiap huruf digeser oleh `offset`. Anda mungkin berasumsi bahwa string hanya berisi huruf kecil dan
// Spasi. Saat menggeser "z" dengan tiga huruf, bungkus ke depan alfabet untuk menghasilkan huruf "c".
// Perhatikan bahwa huruf 'a' memiliki kode 97, 'b' memiliki kode 98, 'z' memiliki kode 122. Anda mungkin juga ingin menggunakan operasi modulo `%` untuk menangani pembungkusan "z" ke depan alfabet .

package main

import (
	"fmt"
	"strings"
)

func caesar(offset int, input string) string {
	transform := func(r rune) rune {
		if offset > 26 {
			offset = offset % 26
		}

		s := int(r) + offset
		if s > 'z' {
			return rune(s - 26)
		} else if s < 'a' {
			return rune(s + 26)
		}
		return rune(s)
	}

	result := strings.Map(transform, input)

	return result
}

func main() {

	fmt.Println(caesar(3, "abc")) // def

	fmt.Println(caesar(2, "alta")) // cnvc

	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi

	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))

	// bcdefghijklmnopqrstuvwxyza

	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))

	// mnopqrstuvwxyzabcdefghijkl

}
