package main

import (
	"fmt"
	"strings"
)

func validCh(ch byte) bool {
	if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
		return true
	}
	return false
}

func generateTag(caption string) string {
	result := "#"
	i, j, captionLen := 0, 0, len(caption)
	first := false
	for j < captionLen {
		if !(validCh(caption[j])) || j == captionLen-1 {
			//fmt.Printf("%c %d %d\n", caption[j], j, captionLen)
			if j == captionLen-1 && validCh(caption[j]) {
				j = captionLen
			}
			if !validCh(caption[i]) {
				i++
				j++
				continue
			}
			if first {
				result += strings.ToUpper(string(caption[i]))
			} else {
				result += strings.ToLower(string(caption[i]))
				first = true
			}
			if j > i+1 {
				result += strings.ToLower(caption[i+1 : j])
			}

			for j < captionLen {
				if validCh(caption[j]) {
					i = j
					break
				}
				j++
			}
			continue
		}
		j++
	}
	if len(result) > 100 {
		return result[0:100]
	}
	return result
}

func main() {
	fmt.Println(generateTag("Leetcode daily streak achieved "))
	fmt.Println(generateTag(" fPysaRtLQLiMKVvRhMkkDLNedQKffPnCjbITBTOVhoVjiKbfSawvpisDaNzXJctQkn"))
	fmt.Println(generateTag("hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh"))
}
