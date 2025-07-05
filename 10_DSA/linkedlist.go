package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func (l *LinkedList) InsertAtHead(val int) {
	newNode := &ListNode{Val: val, Next: l.Head}
	l.Head = newNode
}

func (l *LinkedList) InsertAtTail(val int) {
	newNode := &ListNode{Val: val}

	if l.Head == nil {
		l.Head = newNode
		return
	}
	cur := l.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newNode
}

func (l *LinkedList) DeleteVal(val int) {
	// if the list is empty
	if l.Head == nil {
		return
	}
	// if the deleted element is first element
	if l.Head.Val == val {
		l.Head = l.Head.Next
		return
	}

	prev := l.Head
	cur := l.Head.Next

	for cur != nil {
		if cur.Val == val {
			prev.Next = cur.Next
			break
		}
		prev = cur
		cur = cur.Next
	}
}

func (l *LinkedList) searchByVal(val int) bool {
	cur := l.Head

	for cur != nil {
		if cur.Val == val {
			return true
		}
		cur = cur.Next
	}
	return false
}

func (l *LinkedList) reverserList() {
	var prev *ListNode = nil
	temp := l.Head

	for temp != nil {
		front := temp.Next
		temp.Next = prev
		prev = temp
		temp = front
	}
	l.Head = prev
}

func (l *LinkedList) Print() {
	current := l.Head
	for current != nil {
		fmt.Print(current.Val, " → ")
		current = current.Next
	}
	fmt.Println("nil")
}

//insert at tail
// insert at head
// search value
// delete a value

func main() {
	ll := &LinkedList{}

	ll.InsertAtHead(3)
	ll.InsertAtHead(2)
	ll.InsertAtHead(1)
	ll.Print() // 1 → 2 → 3 → nil

	ll.InsertAtTail(4)
	ll.InsertAtTail(5)
	ll.Print() // 1 → 2 → 3 → 4 → 5 → nil

	ll.DeleteVal(3)
	ll.Print() // 1 → 2 → 4 → 5 → nil

	fmt.Println("Search 4:", ll.searchByVal(4))     // true
	fmt.Println("Search 100:", ll.searchByVal(100)) // false
}
