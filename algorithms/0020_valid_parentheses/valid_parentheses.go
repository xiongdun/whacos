package main

import "fmt"

func main() {

	fmt.Println(valid_parentheses("(((((({[]}))))))()(){[()]}{()}{[]}"))
	fmt.Println(valid_parentheses("){"))

	nums := []int{1, 3, 3, 4, 5, 5, 6}

	fmt.Println(nums)

	fmt.Println(nums[:])
}

func valid_parentheses(s string) bool {
	if len(s) == 0 {
		return true
	}

	if len(s)%2 != 0 {
		return false
	}

	var arr []byte
	var m = map[byte]byte{')': '(', ']': '[', '}': '{'}

	for i := 0; i < len(s); i++ {
		if s[i] == 40 || s[i] == 91 || s[i] == 123 {
			arr = append(arr, byte(s[i]))
		} else {
			if len(s) > 0 {
				flag := arr[len(arr)-1]
				if m[byte(s[i])] != flag {
					return false
				}
				arr = arr[:len(arr)-1]
			} else {
				return false
			}
		}
	}

	if len(arr) <= 0 {
		return true
	}

	return false

}
