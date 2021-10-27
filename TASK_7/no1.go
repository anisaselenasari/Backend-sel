// SIMPLE EQS
// We have three different integers, x, y and z, which satisfy the following three relations:
// x + y + z = A
// xyz = B
// x^2 + y^2 + z^2 = C
// You are asked to write a program that solves for x, y and z for given values of A, B and C. (1 ≤ A, B, C ≤ 10000).

package main

import (
	"fmt"
)

func SimpleEquations(a, b, c int) {
	// your code here
	for x := 1; x <= 10000; x++ {
		for y := 1; y <= 10000-x; y++ {
			for z := 1; z <= 10000/(x*y); z++ {
				if x+y+z == a {
					if x*y*z == b {
						if x*x+y*y+z*z == c {
							fmt.Println(x, y, z)
							return
						}
					}
				}
			}
		}
	}

	fmt.Println("no solution")
	return
}
func main() {

	SimpleEquations(1, 2, 3) // no solution

	SimpleEquations(6, 6, 14) // 1 2 3

	SimpleEquations(16, 50, 126) // no solution

}
