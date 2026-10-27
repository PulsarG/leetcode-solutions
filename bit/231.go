// https://leetcode.com/problems/range-sum-query-2d-immutable/description/

package main

import "fmt"

func main() {
	fmt.Println(Run())
}

func Run() bool {
	n := 16
	x := n
	count := 0
	for x > 1 {
		x >>= 1
		count++
	}
	n2 := 1 << count
	fmt.Println(count)
	fmt.Println(x, n2)

	return (n == n2)
}
