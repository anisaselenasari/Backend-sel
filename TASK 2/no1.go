//1. Luas Permukaan tabung
package main

import "fmt"

func main() {
	var luas, r, t float32
	const pi = 3.14

	fmt.Print("Masukkan jari jari: ")
	fmt.Scanln(&r)
	fmt.Print("Masukkan tinggi tabung: ")
	fmt.Scanln(&t)

	luas = 2 * 3.14 * r * (r + t)

	fmt.Println("Luas Permukaan Tabung adalah", luas)
}
