// https://leetcode.com/problems/rotate-image/description/?envType=problem-list-v2&envId=matrix
// 48

package main

import "fmt"

func main() {
	fmt.Println(Run())
}

func Run() []int {
	m := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	r := make([]int, 0)

	lS := len(m[0])
	lC := len(m)

	vL := 0      // x
	vP := lS - 1 // y
	nL := lC - 1 // z
	nP := lS - 1 // w

	count := lS * lC

	for i := 0; i < lC; i++ {
		for i := vL; i <= vP; i++ {
			if count != 0 {
				r = append(r, m[vL][i])
				count--
			} else {
				break
			}
		}
		vL++
		for i := vL; i <= nL; i++ {
			if count != 0 {
				r = append(r, m[i][vP])
				count--
			} else {
				break
			}
		}
		nP--
		for i := nP; i >= vL-1; i-- {
			if count != 0 {
				r = append(r, m[nL][i])
				count--
			} else {
				break
			}
		}
		vP--
		for i := nL - 1; i > vL-1; i-- {
			if count != 0 {
				r = append(r, m[i][vL-1])
				count--
			} else {
				break
			}
		}
		nL--
	}

	return r
}
