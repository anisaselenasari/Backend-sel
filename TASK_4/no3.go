// Fungsi swap number adalah fungsi yang bertugas untuk menukar 2 nilai dalam sebuah variabel. Contoh variabel a memiliki nilai 10,
// variabel b memiliki nilai 20. Setelah ditukar, a memiliki nilai 20 dan b memiliki nilai 10. Buatlah sebuah fungsi tersebut dengan menggunakan pointer!
// jadi tuh kaya misal si a mau nyimpen pointer nah itu tuh jadi ke pointer b
// trus pointer b itu = temp . sedangkan si temp itu di pointer a
// jadi kek si a= 20, trus b=10
package main

import "fmt"

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {

	a := 10

	b := 20

	swap(&a, &b)

	fmt.Println(a, b)

}
