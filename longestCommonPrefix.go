package main

import "sort"

func longestCommonPrefix(strs []string) string {
	// find the shortest word, as that is the longest possible prefix
	// each word must contain the prefix
	// the length of the shortest word is the max

	// loop through each word, and see if it matches the prefix
	// if it does not match, slice the prefix and start over
	// if the slice reaches 0, there is no common prefix
	// if we reach the end of the loop, there is a common prefix

	sort.Strings(strs) //sorts strings in place :)
	currentPrefix := strs[0]

	for k := 1; k < len(strs); k++ {
		for len(currentPrefix) > 0 &&
			(len(currentPrefix) > len(strs[k]) ||
				strs[k][:len(currentPrefix)] != currentPrefix) {
			currentPrefix = currentPrefix[:len(currentPrefix)-1]
		}
	}

	return currentPrefix
}
