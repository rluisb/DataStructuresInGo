package simpleLinkedList

import (
	"fmt"
)

type Node struct {
  content int
  next *Node
}

type LinkedList struct {
  head *Node
  length int
}

func (l *LinkedList) prepend(node *Node)  {
    existingNode := l.head
    l.head = node
    l.head.next = existingNode
    l.length++
}

func main() {
  linkedList := LinkedList{}
  linkedList.prepend(&Node{content: 10})
  linkedList.prepend(&Node{content: 30})
  linkedList.prepend(&Node{content: 43})
  linkedList.prepend(&Node{content: 5})
  linkedList.prepend(&Node{content: 99})
  
  current := linkedList.head

  for {
    fmt.Print(current.content, " -> ")
    if current.next == nil {
      fmt.Print(current.content, " -> nil")
      return
    }
    current = current.next
  }
}


