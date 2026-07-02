// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/?envType=problem-list-v2&envId=two-pointers
// 26

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(Run())
}

func Run() int {
	nums := []int{1, 1, 2}
	count := len(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] == nums[i] {
			nums[i] = 101
			count--
		}
	}
	sort.Ints(nums)
	return count
}
