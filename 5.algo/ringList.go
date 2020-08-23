package main

import "fmt"

type Node struct {
	value    int
	nextNode *Node
}

// 使用map，时间复杂度、空间复杂度均为O(n)
func hasCycle2(head *Node) bool {
	if head == nil || head.nextNode == nil {
		return false
	}

	m := make(map[*Node]bool, 5)
	for head.nextNode != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = true
		head = head.nextNode
	}
	return false
}

// 快慢指针，最优解，时间复杂度O(n),空间复杂度O(1)
func hasCycle(head *Node) bool {
	if head == nil || head.nextNode == nil {
		return false
	}

	slow := head
	fast := head
	for fast != nil && fast.nextNode != nil {
		slow = slow.nextNode
		fast = fast.nextNode.nextNode
		if fast == slow {
			return true
		}
	}
	return false
}

func main() {
	node1 := new(Node)
	node2 := new(Node)
	node3 := new(Node)
	node4 := new(Node)
	node5 := new(Node)
	node1.value = 1
	node2.value = 2
	node3.value = 3
	node4.value = 4
	node5.value = 5
	node1.nextNode = node2
	node2.nextNode = node3
	node3.nextNode = node4
	node4.nextNode = node2
	//node4.nextNode = node5

	fmt.Println(hasCycle2(node1))
}
