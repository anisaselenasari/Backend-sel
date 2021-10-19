package main

import "fmt"

func pangkat(base, pangkat int) int {
	var hasil int = 1
	for i := 0; i < pangkat; i++ {
		hasil *= base
	}
	return hasil
}

func main() {

	fmt.Println(pangkat(2, 3)) // 8

	fmt.Println(pangkat(5, 3)) // 125

	fmt.Println(pangkat(10, 2)) // 100

	fmt.Println(pangkat(2, 5)) // 32

	fmt.Println(pangkat(7, 3)) // 343

	fmt.Println("Program ini Mengambil Input x dan n lalu mengembalikan nilai x^n")
	fmt.Println("Masukkan Nilai x dan n (pisahkan dengan spasi):")
	var x, n int
	fmt.Scanf("%v", &x)
	fmt.Scanf("%v", &n)
	fmt.Println("Hasil x^n =", pangkat(x, n))
}
