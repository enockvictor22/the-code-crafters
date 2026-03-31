package main

import (
	"fmt"
	// "strings"
)

func reverse(text string) string {
	// word := strings.Fields(text)
	// for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
	// 	word[i], word[j] = word[j], word[i]
	// }
	// return strings.Join(word, " ")
	word := []rune(text)
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		word[i], word[j] = word[j], word[i]
	}
	return string(word)
}

func main() {
	// var option int
	// fmt.Println("Welcome Codecrafter. \nPlease select your desired operation")
	// fmt.Scanln(&option)
	// fmt.Println("[1] reverse")

	fmt.Println(reverse("this is dominion's idea"))
	fmt.Println(reverse("Lagos Nigeria"))
	fmt.Println(reverse("Go is fun"))
}
