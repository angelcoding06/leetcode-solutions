package solutions

import (
	"strconv"
)

func IsPalindrome(x int) bool {
	s1 := strconv.Itoa(x) //cadena
	var reversed string
	for i := len(s1) - 1; i >= 0; i-- {
		reversed += string(s1[i])
	}
	return s1 == reversed
}
