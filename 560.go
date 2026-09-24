// https://leetcode.com/problems/range-sum-query-2d-immutable/description/

package main

import "fmt"

func main() {
	fmt.Println(Run())
}

func Run() int {
	nums := []int{1, 2, 1, 2, 1}
	k := 3
	/* if len(nums) == 1 {
		if nums[0] == k {
			return 1
		} else {
			return 0
		}
	} else if len(nums) == 0 {
		return 0
	}

	var c int
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				c++
			}
		}
	} */

	m := make(map[int]int, 0)
	var c, sum int
	sum = 0
	for i := range len(nums) {
		sum += nums[i]
		if sum == k {
			c++
		}
		if v, ok := m[sum-k]; ok {
			c += v
		}
		if _, ok := m[sum]; ok {
			m[sum] += 1
		} else {
			m[sum] = 1
		}
	}

	return c
}
