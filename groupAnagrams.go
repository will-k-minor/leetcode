package main

import (
	"fmt"
	"sort"
)

func sortString(s string) string {
	runes := []rune(s) // Convert string to slice of runes
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes) // Convert sorted slice of runes back to string
}

func groupAnagrams(strs []string) [][]string {
	/**
	  NOT an array of maps!!
	  just sort each word from a-z
	  if they match, boom! add it to the group of anagrams

	  [{"ant": ["tan", "nat"]}]
	*/
	anagramMap := map[string][]string{}
	result := [][]string{}

	for _, str := range strs {
		sortedString := sortString(str)
		arr, ok := anagramMap[sortedString]
		if ok == false {
			anagramMap[sortedString] = []string{str}
		} else {
			anagramMap[sortedString] = append(arr, str)
		}
	}

	for _, arr := range anagramMap {
		result = append(result, arr)
	}

	fmt.Println(anagramMap)
	return result
}