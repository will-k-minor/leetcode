package main

import (
    "fmt"
    "math"
    "strconv"
)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    num1 :=float64(0);
    num2 := float64(0);
    var exp = math.Pow(10, 2);

    for i := float64(0); l1 != nil; i++ {
        num1 = num1 + float64(l1.Val)*math.Pow(10, i) / exp
        // fmt.Printf("num1 %f", num1)
        l1 = l1.Next
    }

    for i := float64(0); l2 != nil; i++ {
        // fmt.Printf("%v", l2.Val)
        num2 = num2 + float64(l2.Val)*math.Pow(10, i) / exp
        l2 = l2.Next
    }
    sumString := strconv.Itoa(int((num1 + num2) * exp));
    result := &ListNode{}
    head := result;
    fmt.Printf("%v", sumString)
    for i := len(sumString) - 1; i >=0 ; i-- {
        val, _ := strconv.Atoi(string(sumString[i]))
        result.Next = &ListNode{Val: val}
        result = result.Next;
    }

    //fmt.Printf("%v", sumString);

    return head.Next;
}
