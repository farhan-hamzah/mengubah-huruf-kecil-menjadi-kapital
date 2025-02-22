package main

import (
	"fmt"
)
func huruf() {
	var char rune
	fmt.Scanf("%c", &char) 
	if char >= 'a' && char <= 'z' {
		upper := char - ('a' - 'A')
		fmt.Printf("%c\n", upper) 
	} else {
		fmt.Printf("%c\n", char) 
	}
}
func main() {
	huruf()
}
