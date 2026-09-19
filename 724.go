// https://leetcode.com/problems/find-pivot-index/description/

package main

import "fmt"

func main() {
	fmt.Println(Run())
}

func Run() int {
	nums := []int{-1, -1, -1, -1, 1, 1}
	pref := make([]int, len(nums))
	def := make([]int, len(nums))
	pref[0] = 0
	def[0] = 0
	for i := 0; i < len(nums)-1; i++ {
		pref[i+1] = pref[i] + nums[i]
	}
	idx := 0
	for i := len(nums) - 1; i > 0; i-- {
		def[idx+1] = def[idx] + nums[i]
		idx++
	}
	for i := 0; i < len(pref); i++ {
		if pref[i] == def[len(def)-i-1] {
			return i
		}
	}
	return -1
}
