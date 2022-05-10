package main

import (
	"fmt"
	"strings"
)

func uniqueSymbolSearch(str string) bool {
	set := make(map[string]struct{})
	str = strings.ToLower(str)
	

	arr := strings.Split(str, "")
	for _, v := range arr {
		if _, ok := set[v]; ok {
			return false
		}
		set[v] = struct{}{}
	}

	return true
}

func main() {
	fmt.Println("abcd: ", uniqueSymbolSearch("abcd"))
	fmt.Println("abCdefA: ", uniqueSymbolSearch("abCdefA"))
	fmt.Println("aabcd: ", uniqueSymbolSearch("aabcd"))
}