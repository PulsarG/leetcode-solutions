// https://leetcode.com/problems/two-sum/description/?envType=problem-list-v2&envId=hash-table
// 1

package main

import "fmt"

func main() {
	fmt.Println(Run())
}

func Run() []int {
	n := []int{2, 7, 11, 15}
	t := 9
	r := make([]int, 0)

	m := make(map[int]int, 0)

	for i, v := range n {
		_, keyIsInMap := m[v]
		if !keyIsInMap {
			//if v <= t {
			m[t-v] = i
			//	}
		} else {
			r = append(r, i)
			r = append(r, m[v])
		}
	}

	return r
}
