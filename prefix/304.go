// https://leetcode.com/problems/range-sum-query-2d-immutable/description/

package main

import "fmt"

func main() {
	fmt.Println(Run())
}

func Run() int {
	mx := [][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}}
	m := Constructor(mx)
	res := m.SumRegion(2, 1, 4, 3)

	return res
}

type NumMatrix struct {
	M [][]int
	P [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	return NumMatrix{
		M: matrix,
		P: calcPref(matrix),
	}
}

func (m *NumMatrix) SumRegion(row1, col1, row2, col2 int) int {
	res := 0
	for i := row1; i <= row2; i++ {
		res += m.P[i][col2+1] - m.P[i][col1]
	}
	return res
}

func calcPref(m [][]int) [][]int {
	res := make([][]int, len(m)+1)
	for i := 0; i < len(res); i++ {
		res[i] = append(res[i], 0)
	}
	for j := 0; j < len(m); j++ {
		for i := 0; i < len(m[j]); i++ {
			res[j] = append(res[j], res[j][i]+m[j][i])
		}
	}
	return res
}
