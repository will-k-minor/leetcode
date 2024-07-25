package main

import "strings";

func lengthOfLongestSubstring(s string) int {
    /**
        loop through the string
            keep track of current substring
            every time we run into a repeating char inside the currentArray, 
                slice the currentArray starting at 1st index + 1
            append char to array
            do comparison against length of currentArray to max
    */

    byteArray := []byte(s);
    currentArray := []byte{}
    maxLength := 0;

    for _, char := range byteArray {
        idx := strings.Index(string(currentArray[:]), string(char))
        if idx != -1 {
            currentArray = currentArray[idx + 1:]
        }
        currentArray = append(currentArray, char);
        if maxLength < len(currentArray) {
            maxLength = len(currentArray);
        }
    }

    return maxLength;
}