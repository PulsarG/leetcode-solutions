// https://leetcode.com/problems/move-zeroes/description/?envType=problem-list-v2&envId=two-pointers
// 283

package main

import (
	"fmt"
)

func main() {
	fmt.Print(Run())
}

func Run() []int {
	nums := []int{0, 0, 1}
	z := 0
	for nZ := range nums {
		if nums[nZ] != 0 {
			nums[z] = nums[nZ]
			z++
		}
	}
	for ; z < len(nums); z++ {
		nums[z] = 0
	}
	return nums
}
