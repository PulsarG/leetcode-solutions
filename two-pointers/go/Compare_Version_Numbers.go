// https://leetcode.com/problems/compare-version-numbers/description/?envType=problem-list-v2&envId=two-pointers
// 165

package main

import (
	"fmt"
	"strconv"
	//"strconv"
	"strings"
)

func main() {
	fmt.Println(Run())
}

func Run() int {
	v1, v2 := "1.0.1", "1"
	r := 0
	s1 := strings.Split(v1, ".")
	s2 := strings.Split(v2, ".")

	for i := 0; ; i++ {
		a, _ := strconv.Atoi(s1[i])
		b, _ := strconv.Atoi(s2[i])
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		if i+1 == len(s1) && i+1 != len(s2) {
			s1 = append(s1, "0")
		}
		if i+1 == len(s2) && i+1 != len(s1) {
			s2 = append(s2, "0")
		}
		if i+1 == len(s1) && i+1 == len(s2) {
			break
		}
	}

	return r
}
