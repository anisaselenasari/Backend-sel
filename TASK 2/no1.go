//1. Luas Permukaan tabung
package main

import "fmt"

func main() {
	var luas, r, t float32
	// float32 adalah angka 32-bit
	const pi = 3.14

	fmt.Print("Masukkan jari jari: ")
	fmt.Scanln(&r)
	//	&r melakukan input nilai jari-jari ke dalam variable r
	fmt.Print("Masukkan tinggi tabung: ")
	fmt.Scanln(&t)
	// &t melakukan input nilai tinggi ke dalam variable t

	luas = 2 * 3.14 * r * (r + t)

	fmt.Println("Luas Permukaan Tabung adalah", luas)
}
