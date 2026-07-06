// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/?envType=problem-list-v2&envId=two-pointers
// 167

package main

import "fmt"

func main() {
	fmt.Println(Run())
}

func Run() []int {
	nums := []int{2, 7, 11, 15}
	target := 9

	a, b := 0, len(nums)-1
	for a < b {
		sum := nums[a] + nums[b]
		if sum == target {
			return []int{a + 1, b + 1}
		}
		if sum > target {
			b--
		} else {
			a++
		}
	}
	return nilпеш
}
