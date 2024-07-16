package solutions


func IsValid(s string) bool {	
	var stack []rune
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else if v == ')' || v == ']' || v == '}' {
			if len(stack) == 0 {
				return false
			}
			if v == ')' && stack[len(stack)-1] != '(' {
				return false
			}
			if v == ']' && stack[len(stack)-1] != '[' {
				return false
			}
			if v == '}' && stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
