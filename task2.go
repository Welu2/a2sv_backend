// package main

// import "fmt"
// // func main(){
// // 	for _,j := range "adffa"{
// // 		fmt.Printf("%c\n",j)
// // 	}
// // }

// func freq_count[T comparable](items []T) map[T]int {
	
// 	dictionary := make(map[T]int)
// 	for _, j := range items {
		
// 		dictionary[j]++
// 	}
// 	return dictionary
// }
// func main(){
// 	word := "asdfrs"
// 	charCounts := freq_count([]rune(word))

// 	// Print in the format "a:1 b:3"
// 	for char, count := range charCounts {
// 		fmt.Printf("%c:%d ", char, count)
// 	}
// }


package main

import (
	"fmt"
	"strings"
	"unicode"
)

// WordFrequency returns a map of word counts, ignoring case and punctuation.
func WordFrequency(input string) map[string]int {
	counts := make(map[string]int)

	// Function to decide where to split: split on anything that isn't a letter or number
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	// Split the string into fields based on the function above
	words := strings.FieldsFunc(input, f)

	for _, word := range words {
		// Normalize to lowercase
		word = strings.ToLower(word)
		counts[word]++
	}

	return counts
}

func main() {
	text := "Hello world! Hello, Go... is it working? Yes, it is."
	result := WordFrequency(text)

	for word, count := range result {
		fmt.Printf("%s:%d ", word, count)
	}
}
