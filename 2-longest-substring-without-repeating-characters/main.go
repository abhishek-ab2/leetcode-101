package main

import "fmt"


func longestSubStr(str string) int {
	char_counts := make(map[rune]int)
	runes := []rune(str)

	max_len := 0
	start := 0
	for end, r := range runes {
		char_counts[r]++

		for char_counts[r] > 1 {
			char_counts[runes[start]]--
			start++
		}

		if l := end-start + 1; l > max_len{
			max_len = l
		}
	}
	return max_len
}

func main() {
	str := "abcabcabcabcabc"
	fmt.Println(longestSubStr(str))
}
