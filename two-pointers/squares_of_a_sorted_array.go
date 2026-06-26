// https://leetcode.com/problems/squares-of-a-sorted-array/description/?envType=problem-list-v2&envId=two-pointers
// 977

package main

import (
	"fmt"
	//"math"
	"sort"
)

func main() {
	fmt.Print(Run())
}

func Run() []int {
	nums := []int{-4, -1, 0, 3, 10}
	for i, v := range nums {
		nums[i] = v * v
		//nums[i] = int(math.Pow(float64(v), 2))
	}
	sort.Ints(nums)
	return nums
}
