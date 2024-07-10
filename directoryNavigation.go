package main;

func minOperations(logs []string) int {
    /** 
        its a stack

        if "../" pop from the stack
        if "c+/" push to the stack
        if "./" don't do anything

        return the length of the stack
    */
    stack := 0;
    for _, log := range logs {
        if log == "../" {
            if stack > 0 {
                stack--;
            }
        } else if log != "./" {
            stack++;
        }
    }

    return stack;
}