package main

import (
	"fmt"
	"math"
)

func main() {
	var bilangan int
	fmt.Println("Masukkan Sebuah Bilangan Bulat")
	fmt.Scanf("%d", &bilangan)

	var sqri int = int(math.Sqrt(float64(bilangan)))

	fmt.Println("Faktor dari Bilangan Tersebut Adalah: ")
	var i int
	for i = 1; i <= sqri; i++ {
		if bilangan%i == 0 {
			if bilangan/i == i {
				fmt.Println(i)
			} else {
				fmt.Println(i)
				fmt.Println(bilangan / i)
			}
		}
	}
}
