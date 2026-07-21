// https://leetcode.com/problems/spiral-matrix/description/?envType=problem-list-v2&envId=matrix
// 54

package main

import "fmt"

func main() {
	fmt.Println(Run())
}

func Run() [][]int {
	m := [][]int{{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16}}

	i, j, l := 0, 0, len(m[0])
	for i < l {
		c := l - i - 1
		for j < c {
			tmp := m[i][j]
			x, y := i, j
			for k := 0; k < 4; k++ {
				z := l - x - 1
				w := m[y][z]
				m[y][z] = tmp
				tmp = w
				x, y = y, z
			}
			j++
		}
		i++
		j = i
	}

	return m
}
