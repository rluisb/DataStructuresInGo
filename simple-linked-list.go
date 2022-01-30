package simpleLinkedList

import (
	"fmt"
)

type Node struct {
  content int
  next *Node
}

type SimpleLinkedList struct {
  nodes []Node
}

func (s *SimpleLinkedList) AddNode(node Node)  {
    s.nodes = append(s.nodes, node)
}

func main() {
  thirdNode := Node{22, nil}
  secondNode := Node{1, &thirdNode}
  firstNode := Node{10, &secondNode}
	simpleLinkedList := SimpleLinkedList{}
  simpleLinkedList.AddNode(firstNode)
  simpleLinkedList.AddNode(secondNode)
  simpleLinkedList.AddNode(thirdNode)
  for i := 0; i < len(simpleLinkedList.nodes); i++ {
    fmt.Print(" -> ", simpleLinkedList.nodes[i].content, " -> ")
  }
}


