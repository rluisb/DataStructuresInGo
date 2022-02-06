package doubleLinkedList

import (
	"fmt"
)

type Node struct {
  previous *Node
  content int
  next *Node
}

type DoublyLinkedList struct {
  head *Node
  length int
}

func (d *DoublyLinkedList) prepend(node *Node)  {
    existingNode := d.head
    d.head = node
    d.head.next = existingNode
    d.length++
}

func main() {
  var firstNode Node
  var secondNode Node
  var thirdNode Node

  firstNode = Node{content: 10, next: &secondNode}
  secondNode = Node{previous: &firstNode, content: 1, next: &thirdNode}
  thirdNode = Node{previous: &secondNode, content: 22}
  

	doublyLinkedList := DoublyLinkedList{}
  doublyLinkedList.AddNode(firstNode)
  doublyLinkedList.AddNode(secondNode)
  doublyLinkedList.AddNode(thirdNode)
  for i := 0; i < len(doublyLinkedList.nodes); i++ {
    fmt.Print(doublyLinkedList.nodes[i].content, " <-> ")
  }
}


