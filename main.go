// package main

// import (
// 	"fmt"
// )

// func main() {
//   fmt.Println("Hello Wolrd")
// }

package main

import (
	"fmt"
)

type Node struct {
  previous *Node
  content int
  next *Node
}

type DoubleLinkedList struct {
  nodes []Node
}

func (d *DoubleLinkedList) AddNode(node Node)  {
    d.nodes = append(d.nodes, node)
}

func main() {
  var firstNode Node
  var secondNode Node
  var thirdNode Node

  firstNode = Node{content: 10, next: &secondNode}
  secondNode = Node{previous: &firstNode, content: 1, next: &thirdNode}
  thirdNode = Node{previous: &secondNode, content: 22}
  

	doubleLinkedList := DoubleLinkedList{}
  doubleLinkedList.AddNode(firstNode)
  doubleLinkedList.AddNode(secondNode)
  doubleLinkedList.AddNode(thirdNode)
  for i := 0; i <= len(doubleLinkedList.nodes); i++ {
    if doubleLinkedList.nodes[i].next == nil {
        fmt.Print(doubleLinkedList.nodes[i].content, " <-> nil")
        return
    }
    if doubleLinkedList.nodes[i].previous == nil{
        fmt.Print("nil <-> ")
    }
    fmt.Print(doubleLinkedList.nodes[i].content, " <-> ")
  }
}


