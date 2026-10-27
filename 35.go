// https://leetcode.com/problems/search-insert-position/description/

package main

import "fmt"

func main() {
	a := []int{1, 3, 5, 6}
	fmt.Println(Run(a, 7))
}

func Run(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	st, end := 0, len(nums)-1
	for st <= end {
		x := (end + st) / 2
		if nums[x] == target {
			return x
		} else if nums[x] < target {
			st = x + 1
		} else {
			end = x - 1
		}
		if st > end {
			return st
		}
	}
	return -1
}
