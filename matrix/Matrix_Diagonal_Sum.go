// https://leetcode.com/problems/matrix-diagonal-sum/description/
// 1572

package main

import "fmt"

func main() {
	fmt.Println(run())
}

func run() int {
	var m = [][]int{{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1}}

	l, r, res := 0, len(m)-1, 0

	for i := 0; i < len(m); i++ {
		if l == r {
			res += m[i][l]
		} else {
			res += (m[i][l] + m[i][r])
		}
		l++
		r--
	}
	return res

	/* n := len(m)
	res := 0

	for i := 0; i < n; i++ {
		res += m[i][i]
		if i != n-1-i {
			res += m[i][n-1-i]
		}
	}
	return res */
}
