package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Frog(jumps []int) int {
	// your code here
	alljump1, alljump2 := []int{}, []int{}
	total1, total2 := 0, 0
	for i := 0; i+2 < len(jumps); i++ {
		if abs(jumps[i]-jumps[i+1]) < abs(jumps[i]-jumps[i+2]) {
			alljump1 = append(alljump1, abs(jumps[i]-jumps[i+1]))
		} else {
			alljump1 = append(alljump1, abs(jumps[i]-jumps[i+2]))
			i++
		}
	}
	for i := 0; i+1 < len(jumps); i++ {
		if i+2 == len(jumps) {
			alljump2 = append(alljump2, abs(jumps[i]-jumps[i+1]))
		} else {
			alljump2 = append(alljump2, abs(jumps[i]-jumps[i+2]))
			i++
		}
	}
	for _, val := range alljump1 {
		total1 += val
	}
	for _, val := range alljump2 {
		total2 += val
	}
	if total1 <= total2 {
		return total1
		//return alljump1
	}
	return total2
	//return alljump2
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20})) // 30

	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40

	fmt.Println(Frog([]int{10, 10})) // 0

	fmt.Println(Frog([]int{10, 30, 40, 20, 40})) // 30

	fmt.Println(Frog([]int{10, 30, 40, 20})) // 30

	fmt.Println(Frog([]int{10, 30, 40, 20, 40, 50, 40})) // 30

	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50, 60})) // 30

	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50, 20})) //60

	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50, 20})) //60
}
