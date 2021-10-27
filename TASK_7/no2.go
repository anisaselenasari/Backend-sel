// MONEY COINS
// Terdapat sebuah counter untuk melakukan penukaran uang menjadi pecahan yang lebih kecil.
// Counter ini memiliki coin uang dengan pecahan sbb : 1, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000 dan masing-masing pecahan jumlahnya saat ini tidak terbatas.
// Buatlah sebuah program untuk menghitung pertukaran uang. Hasil penukaran harus memenuhi optimal minimum jumlah koin hasil penukaran!

package main

import "fmt"

func moneyCoins(money int) []int {

	pecahanMoney := []int{1, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000}
	changes := []int{}
	// your code here
	for i := len(pecahanMoney) - 1; i >= 0; i-- {
		for pecahanMoney[i] <= money {
			changes = append(changes, pecahanMoney[i])
			money -= pecahanMoney[i]
			if money == 0 {
				return changes
			}
		}
	}
	return changes
}
func main() {

	fmt.Println(moneyCoins(123)) // [100 20 1 1 1]

	fmt.Println(moneyCoins(432)) // [200 200 20 10 1 1]

	fmt.Println(moneyCoins(543)) // [500, 20, 20, 1, 1, 1]

	fmt.Println(moneyCoins(7752)) // [5000, 2000, 500, 200, 50, 1, 1]

	fmt.Println(moneyCoins(15321)) // [10000 5000 200 100 20 1]

}
