// https://leetcode.com/problems/isomorphic-strings/description/?envType=problem-list-v2&envId=hash-table
// 205

package main

import "fmt"

func main() {
	fmt.Println(Run())
}

func Run() bool {
	s := "badc"
	t := "baba"

	if len(s) != len(t) {
		return false
	}

	if len(s) == 1 {
		return true
	}

	m := make(map[string]int)
	m2 := make(map[string]int)

	for _, v := range s {
		c, ok := m[string(v)]
		if !ok {
			m[string(v)] = 1
		} else {
			m[string(v)] = c + 1
		}
	}

	for _, v := range t {
		c, ok := m2[string(v)]
		if !ok {
			m2[string(v)] = 1
		} else {
			m2[string(v)] = c + 1
		}
	}

	if len(m) != len(m2) {
		return false
	}

	m3 := make(map[rune]rune)

	for i, v := range s {
		fmt.Println("s ", i, v)
		fmt.Println("t ", i, t[i])
		_, ok := m3[v]
		if !ok {
			m3[v] = rune(t[i])
		} else {
			if m3[v] != rune(t[i]) {
				return false
			}
		}
	}

	return true
}
