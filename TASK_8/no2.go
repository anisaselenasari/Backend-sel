package main

import "fmt"

func fibo(n int) int {

	// your code here
	fiboCollect := []int{0, 1, 1}
	i := 3
	for i <= n {
		fiboCollect = append(fiboCollect, fiboCollect[i-1]+fiboCollect[i-2])
		i++
	}

	return fiboCollect[n]
}

func main() {

	fmt.Println(fibo(0)) // 0

	fmt.Println(fibo(1)) // 1

	fmt.Println(fibo(2)) // 1

	fmt.Println(fibo(3)) // 2

	fmt.Println(fibo(5)) // 5

	fmt.Println(fibo(6)) // 8

	fmt.Println(fibo(7)) // 13

	fmt.Println(fibo(9)) // 34

	fmt.Println(fibo(10)) // 55

}
