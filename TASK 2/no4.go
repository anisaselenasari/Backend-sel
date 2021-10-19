package main

import "fmt"

func primeNumber(number int) bool {
	var isPrime bool
	if number <= 2 {
		if number == 2 {
			isPrime = true
		} else {
			isPrime = false
		}
	} else {
		i := 2
		for i <= number {
			if i >= number {
				isPrime = true
				break
			} else {
				if number%i == 0 {
					isPrime = false
					break
				}
			}
			i++
		}
	}
	return isPrime
}

func main() {
	fmt.Println(primeNumber(11)) // true

	fmt.Println(primeNumber(13)) // true

	fmt.Println(primeNumber(17)) // true

	fmt.Println(primeNumber(20)) // false

	fmt.Println(primeNumber(35)) // false

	fmt.Println("Masukkan Sebuah Bilangan: ")
	var number int
	fmt.Scan(&number)
	if primeNumber(number) == true {
		fmt.Println("Bilangan Prima")
	} else {
		fmt.Println("Bukan Bilangan Prima")
	}
}
