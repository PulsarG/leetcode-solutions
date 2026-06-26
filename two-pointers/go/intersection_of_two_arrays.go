// https://leetcode.com/problems/intersection-of-two-arrays/description/
// 349 easy

package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Print(Run())
}

func Run() []int {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	sl := make([]int, 0)

	for _, v := range nums1 {
		if slices.Contains(nums2, v) {
			if !slices.Contains(sl, v) {
				sl = append(sl, v)
			}
		}
	}

	return sl

}
