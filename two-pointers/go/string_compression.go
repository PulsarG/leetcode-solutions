// https://leetcode.com/problems/string-compression/description/?envType=problem-list-v2&envId=two-pointers
// 443

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Run())
}

func Run() int {
	var chars = []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	s, f, c := 0, 1, 1
	for f < len(chars) {
		if chars[s] == chars[f] && f != len(chars)-1 {
			f++
			c++
		} else if chars[s] != chars[f] && c == 1 {
			s = f
			f++
		} else if chars[s] != chars[f] && c > 1 {
			cc := strconv.Itoa(c)
			copy(chars[s+1:], cc)
			chars = append(chars[:s+len(cc)+1], chars[f:]...)
			c = 1
			s = s + 1 + len(cc)
			f = s + 1
		} else if chars[s] == chars[f] && f == len(chars)-1 {
			cc := strconv.Itoa(c + 1)
			copy(chars[s+1:], cc)
			chars = append(chars[:s+len(cc)+1], chars[f+1:]...)
			break
		}
	}
	return len(chars)
}
