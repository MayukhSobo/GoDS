package main

import "fmt"

// Node A node for storing data into the link
type Node struct {
	Data int
	Next *Node
}

// LL Singly link list
type LL struct {
	Head *Node
	Curr *Node
	Size int
}

// Push push back each element in the link list
func Push(ll *LL, data int) {
	if ll.Head == nil {
		var head Node
		head.Data = data
		head.Next = nil
		ll.Head = &head
		ll.Curr = &head
	} else {
		curr := ll.Head
		for {
			if curr.Next == nil {
				break
			}
			curr = curr.Next
		}
		var temp Node
		temp.Data = data
		temp.Next = nil
		curr.Next = &temp
		ll.Curr = curr.Next
	}
	ll.Size++
}

// PushBulk Push all the elements of the list
func PushBulk(ll *LL, data []int) {
	for _, each := range data {
		Push(ll, each)
	}
}

// PushAt Push an element at an index
func PushAt(ll *LL, data int, index int) {
	if index < ll.Size {
		count := 0
		temp := ll.Head
		var temp2 *Node
		for count != index {
			temp2 = temp
			temp = temp.Next
			count++
		}
		if temp == ll.Head {
			var newHead Node
			newHead.Data = data
			newHead.Next = ll.Head
			ll.Head = &newHead
		} else {
			var newNode Node
			newNode.Data = data
			newNode.Next = temp
			temp2.Next = &newNode
		}

		ll.Size++
	}
}

// Pop Removes an element from the back of the linklist
func Pop(ll *LL) {
	temp := ll.Head
	var temp2 *Node
	for temp.Next != nil {
		temp2 = temp
		temp = temp.Next
	}
	temp2.Next = nil
	ll.Size--
}

func print(currNode *Node, reverse bool) {
	if currNode == nil {
		return
	}
	if !reverse {
		fmt.Println(currNode.Data)
	}
	print(currNode.Next, reverse)
	if reverse {
		fmt.Println(currNode.Data)
	}
}

// Print Print the link list
func Print(ll *LL, reverse bool) {
	print(ll.Head, reverse)
	fmt.Println("======================")
}

// Remove Remove an element from an index
func Remove(ll *LL, index int) {
	if index < ll.Size {
		count := 0
		temp := ll.Head
		var temp2 *Node
		for count != index {
			temp2 = temp
			temp = temp.Next
			count++
		}
		if temp == ll.Head {
			ll.Head = ll.Head.Next
			temp.Next = nil
		} else {
			next := temp.Next
			temp2.Next = next
			temp.Next = nil
		}
		ll.Size--
	}
}

// FindAt Find an element at an index in link list
func FindAt(ll *LL, index int) *Node {
	count := 0
	ret := ll.Head
	for count != index {
		ret = ret.Next
		count++
	}
	return ret
}

// Find find if an element is present in link list or not
func Find(ll *LL, elem int) *Node {
	ret := ll.Head
	for ret.Next != nil {
		if ret.Data == elem {
			return ret
		}
		ret = ret.Next
	}
	return nil
}

func main() {
	ll := LL{nil, nil, 0}
	Push(&ll, 10)
	Push(&ll, 20)
	Push(&ll, 30)
	Push(&ll, 40)
	Push(&ll, 50)
	Push(&ll, 60)
	Push(&ll, 70)
	Push(&ll, 80)
	PushBulk(&ll, []int{90, 100, 110, 120})
	PushAt(&ll, 5, 0)
	Print(&ll, false)
	Remove(&ll, 1)
	Print(&ll, false)
	fmt.Println(FindAt(&ll, 3))
	fmt.Println(Find(&ll, 13))
	fmt.Println(Find(&ll, 5))
}
