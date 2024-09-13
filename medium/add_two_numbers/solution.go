package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Function to add two numbers represented by linked lists
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{} // Create a dummy node to simplify the result list
	current := dummy     // Pointer to the current node in the result list
	carry := 0           // Initialize carry to 0

	// Loop until both lists are empty
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry // Start with carry

		// Add value from l1 if available
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		// Add value from l2 if available
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		// Update carry for the next iteration
		carry = sum / 10
		// Create a new node with the sum's last digit
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next // Move to the next node
	}

	return dummy.Next // Return the next of dummy to skip the initial node
}

// Helper function to create a linked list from a slice
func createLinkedList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummy.Next
}

// Helper function to print the linked list
func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	l1 := createLinkedList([]int{2, 4, 3})
	l2 := createLinkedList([]int{5, 6, 4})

	result := addTwoNumbers(l1, l2)
	printLinkedList(result) // Output: 7 -> 0 -> 8
}
