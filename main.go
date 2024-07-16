package main

import (
	"fmt"

	s "github.com/angelcoding06/leetcode-solutions/solutions"
)

func main() {

	node3b := s.ListNode{Val: 4, Next: nil}
	node2b := s.ListNode{Val: 3, Next: &node3b}
	node1b := s.ListNode{Val: 1, Next: &node2b}

	node3a := s.ListNode{Val: 4, Next: nil}
	node2a := s.ListNode{Val: 2, Next: &node3a}
	node1a := s.ListNode{Val: 1, Next: &node2a}

	fmt.Println(s.MergeTwoLists(&node1a, &node1b))
	
}
