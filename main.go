package main

import (
	"fmt"
	"strings"
)

// reduce the previous time by half. Did not use any slice.
// case-insensitive
func main() {
	givenword := "level"
	word := strings.ToLower(givenword)
	var check bool

	for i := 0; i <= (len(word)-1)/2; i++ {
		if word[i] == word[len(word)-i-1] {
			fmt.Println(word[i], word[len(word)-i-1])
			check = true
			continue
		}
		check = false
		break
	}

	if check == true {
		fmt.Println(word, "is a palindrome.")
		return
	}
	fmt.Println(word, "is NOT a palindrome.")
}
