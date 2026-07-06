// https://leetcode.com/problems/3sum/description/?envType=problem-list-v2&envId=two-pointers
// 15

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(Run())
}

func Run() [][]int {
	var nums = []int{-1, 0, 1, 2, -1, -4}
	var res = make([][]int, 0)
	sort.Ints(nums)

	for i, v := range nums {
		if v > 0 {
			break
		}
		if i > 0 && v == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			summ := v + nums[l] + nums[r]
			if summ > 0 {
				r -= 1
			} else if summ < 0 {
				l += 1
			} else {
				res = append(res, []int{v, nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}

	return res
}
