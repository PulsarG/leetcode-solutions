// https://leetcode.com/problems/container-with-most-water/description/?envType=problem-list-v2&envId=two-pointers
// 11

package main

import "fmt"

func main() {
	fmt.Println(Run())
}

func Run() int {
	var h = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	l := 0
	r := len(h) - 1
	res := 0

	for l < r {
		tmp := 0
		x := r - l
		a, b := h[l], h[r]
		if a >= b {
			tmp = x * b
			if tmp > res {
				res = tmp
			}
			r--
		} else {
			tmp = x * a
			if tmp > res {
				res = tmp
			}
			l++
		}
	}

	return res
}
