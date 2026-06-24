// https://leetcode.com/problems/merge-sorted-array/description/?envType=problem-list-v2&envId=two-pointers
// 88 easy

package main


import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{0}
	nums2 := []int{1}

	m := 0
	//n := 3

	nums1 = append(nums1[:m], nums2...)
	sort.Ints(nums1)
	fmt.Println(nums1)
}
