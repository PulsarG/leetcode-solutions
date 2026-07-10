// https://leetcode.com/problems/transpose-matrix/description/?envType=problem-list-v2&envId=matrix
// 867

package main

import "fmt"

func main() {
	fmt.Println(Run())
}

func Run() [][]int {
	var m = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	r := make([][]int, len(m[0]))
	for i := range r {
		r[i] = make([]int, len(m))
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			r[j][i] = m[i][j]
		}
	}

	return r
}
