package solutions

import (
	"fmt"
)

func RomanToInt(s string) int {
	Romanos := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	result := 0
	for i := 0; i < len(s); i++ {
		result += Romanos[string(s[i])]
		if i > 0 && Romanos[string(s[i])] > Romanos[string(s[i-1])] {
			result -= 2 * Romanos[string(s[i-1])]
			fmt.Println(Romanos[string(s[i])])
		}
	}
	return result
}
