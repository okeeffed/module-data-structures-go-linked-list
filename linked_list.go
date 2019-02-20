/*
@author Dennis O'Keeffe

Methods:

1. insertFirst: Insert at the head of list
2. size: Fetch size
3. getFirst: Fetch the first node
4. getLast: Fetch the last node
5. clear: Remove all node
6. removeFirst: Remove head of list
7. removeLast: Remove last node
8. insertLast: Insert at end of list
9. getAt: Fetch node at index
10. removeAt: Remove node at index
11. insertAt: Insert at index
12. forEach: Iterate through list and run function on list
*/

package linkedlist

import (
	"sync"
)

// Node - Setup base node type for linked list
type Node struct {
	data int
	next *Node
}

// List - Basic singularly linked list
type List struct {
	head *Node
	size int
	lock sync.RWMutex
}

func (l *List) insertFirst(n *Node) {
	l.lock.Lock()
	if l.head == nil {
		l.head = n
	} else {
		last := l.head
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = n
	}
	l.size++
	l.lock.Unlock()
}

func (l *List) getFirst() *Node {
	return (*l).head
}
