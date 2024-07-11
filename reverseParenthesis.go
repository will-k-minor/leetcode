package main

func reverseParentheses(s string) string {
    /**
        Is this just a bidirectional stack?

        YESS DO THIS ONE! NOT RECURSION~
        Oh no! its a series of arrays!
            every time we hit a (, thats a new array!
            when we hit a ), go up to the prev array

        RECURSION!!

        if we hit a (,
            reverse(s[x:n])
            append result to current array
        if we hit a )
            actually reverse strings in current array
            return
        else
            append to the array
    */
    return recursiveReversive(s);
}

func recursiveReversive(s string) string {
    if len(s) == 0 {
        return ""
    }
    var result string;
    i := 0
    j := len(s) - 1
    for i <= j {
        if s[i] == '(' {
            result += recursiveReversive(s[i+1:j])
        } else if s[i] == ')' {
            return result;
        } else {
            result += string(s[i])
        }
        i++

    }
    return result
}
