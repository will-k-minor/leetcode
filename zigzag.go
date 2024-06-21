package main

import "strings"

func zigzag(s string, numRows int) string {
    if numRows == 1 {
        return s;
    }
    
    arr := make([]string, numRows);
    counter := 0;
    dir := -1;

    for _, v := range s {
        arr[counter] += string(v);
        if counter == 0 || counter == numRows - 1 {
            dir = -dir;
        }
        counter += dir;
    }

    return strings.Join(arr, "");
}
