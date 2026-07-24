// https://leetcode.com/problems/roman-to-integer/description/?envType=problem-list-v2&envId=hash-table
// 13

package main

import (
	"fmt"
)

func main() {
	fmt.Println(Run())
}

func Run() int {
	s := "MCMXCIV"

	var r [256]int
	r['I'] = 1
	r['V'] = 5
	r['X'] = 10
	r['L'] = 50
	r['C'] = 100
	r['D'] = 500
	r['M'] = 1000

	res := 0
	n := len(s)

	for i := 0; i < n; i++ {
		cur := r[s[i]]
		if i+1 < n && cur < r[s[i+1]] {
			res -= cur
		} else {
			res += cur
		}
	}

	return res
}
