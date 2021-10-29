package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {

	// your code here
	start := jumps[0]
	last := len(jumps) - 1
	i := 1
	c := 0
	var res float64

	for i < len(jumps) {
		if len(jumps)-1 == i {
			res += math.Abs(float64(c - jumps[i]))
			// fmt.Println(jumps[i])
			return int(res)
		}
		if len(jumps)-2 == i {
			res += math.Abs(float64(c - jumps[i+1]))
			// fmt.Println(jumps[i+1])
			return int(res)
		}
		temp1 := math.Abs(float64(jumps[i] - jumps[last]))
		temp2 := math.Abs(float64(jumps[i+1] - jumps[last]))
		if i == 1 {
			if temp1 < temp2 {
				res = math.Abs(float64(start - jumps[i]))
				c = jumps[i]
				// fmt.Println(res)
				i += 1
			} else {
				res = math.Abs(float64(start - jumps[i+1]))
				c = jumps[i+1]
				// fmt.Println(res)
				i += 2
			}
		}
		if temp1 < temp2 {
			res += math.Abs(float64(c - jumps[i]))
			c = jumps[i]
			// fmt.Println(res)
			i += 1
		} else {
			res += math.Abs(float64(c - jumps[i+1]))
			c = jumps[i+1]
			// fmt.Println(res)
			i += 2
		}

	}
	return int(res)
}

func main() {

	fmt.Println(Frog([]int{10, 30, 40, 20})) // 30

	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40

}
