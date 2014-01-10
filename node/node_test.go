package node

import "testing"
import "fmt"

func TestRemove(t *testing.T) {
	nl := NewNodeList()

	n := NewNode(10.0, nil)
	nl.Insert(n)
	nl.Insert(NewNode(20.0, nil))
	n2 := NewNode(30.0, nil)
	nl.Insert(n2)
	fmt.Println(nl)
	nl.Remove(n)
	fmt.Println(nl)
	n2.AddChild(n)
	fmt.Println(nl)

	n3 := NewNode(40.0, nil)
	nl.Insert(n3)
	nr := nl.Remove(NewNode(20.0, nil))
	fmt.Println(nl)
	n3.AddChild(nr)
	fmt.Println(nl)
}
