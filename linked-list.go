package main

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

func (l *LinkedList) append(node *Node)  {
    if l.head == nil {
      l.head = node
      return
    }

    current := l.head
    
    for {
      if current.next == nil {
        current.next = node
        return
      }
      current = current.next
    }
}


func main() {
  linkedList := LinkedList{}
  linkedList.prepend(&Node{content: 10})
  linkedList.prepend(&Node{content: 30})
  linkedList.prepend(&Node{content: 43})
  linkedList.prepend(&Node{content: 5})
  linkedList.prepend(&Node{content: 99})

  linkedList.append(&Node{content: 43})
  linkedList.append(&Node{content: 5342})
  linkedList.append(&Node{content: 11223})


  current := linkedList.head

  for {
    if current.next == nil {
      fmt.Print(current.content, " -> ", current.next)
      return
    }
    fmt.Print(current.content, " -> ")
    current = current.next
  }
}