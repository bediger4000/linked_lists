package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

/*
 * This question was asked by Snapchat.
 *
 * Given the head to a singly linked list,
 * where each node also has a "random" pointer
 * that points to anywhere in the linked list,
 * deep clone the list.
 */

type Node struct {
	Data  int
	Next  *Node
	Extra *Node
}

func main() {
	n := 1
	graphViz := false
	if os.Args[1] == "-g" {
		n = 2
		graphViz = true
	}
	head := Compose(os.Args[n:])

	deepCopy := Copy(head)

	if graphViz {
		fmt.Printf("digraph g {\nrankdir=\"LR\";\n")
		fmt.Printf("subgraph cluster_0 {\nlabel=\"original\";\n")
		DrawPrefixed(head, "a")
		fmt.Printf("\n}\n")
		fmt.Printf("subgraph cluster_1 {\nlabel=\"deep copy\";\n")
		DrawPrefixed(deepCopy, "b")
		fmt.Printf("\n}\n")
		fmt.Printf("\n}\n")

		return
	}

	TextOutput(head, deepCopy)
}

func TextOutput(head, deepCopy *Node) {
	originalNodes := make(map[*Node]bool)
	copyNodes := make(map[*Node]bool)
	fmt.Printf("    Original                          |    Deep Copy\n")
	var h, c *Node
	for h, c = head, deepCopy; h != nil && c != nil; h, c = h.Next, c.Next {
		originalNodes[h] = true
		copyNodes[c] = true
		if (h.Data != c.Data) || (h.Extra.Data != c.Extra.Data) {
			fmt.Printf("Something wrong with these nodes\n")
		}
		fmt.Printf("%p   %d   %p   %d   |   %p   %d   %p   %d\n",
			h, h.Data, h.Extra, h.Extra.Data,
			c, c.Data, c.Extra, c.Extra.Data,
		)
	}

	// Check if .Extra elements point off-list
	for h, c = head, deepCopy; h != nil && c != nil; h, c = h.Next, c.Next {
		if originalNodes[c] {
			fmt.Printf("Node at %p, value %d in both lists\n", c, c.Data)
		}
		if originalNodes[c.Extra] {
			fmt.Printf("Copy node at %p, value %d, Extra points off list, %p, %d\n", c, c.Data, c.Extra, c.Extra.Data)
		}
		if copyNodes[h] {
			fmt.Printf("Node at %p, value %d in both lists\n", h, h.Data)
		}
		if copyNodes[h.Extra] {
			fmt.Printf("Original node at %p, value %d, Extra points off list, %p, %d\n", h, h.Data, h.Extra, h.Extra.Data)
		}
	}

	if h != nil || c != nil {
		fmt.Printf("Problem: two lists aren't the same length\n")
		fmt.Printf("Original:\n")
		count := 0
		for n := head; n != nil; n = n.Next {
			fmt.Printf("%d -> ", n.Data)
		}
		fmt.Printf("\n%d nodes in original\n", count)
		fmt.Printf("Deep copy:\n")
		count = 0
		for n := deepCopy; n != nil; n = n.Next {
			fmt.Printf("%d -> ", n.Data)
		}
		fmt.Printf("\n%d nodes in deep copy\n", count)
	}
}

func Compose(stringValues []string) (head *Node) {
	if len(stringValues) == 0 {
		return
	}
	var nodes []*Node
	var tail *Node
	for _, str := range stringValues {
		num, err := strconv.Atoi(str)
		if err != nil {
			continue
		}
		if head == nil {
			head = &Node{Data: num}
			tail = head
			continue
		}
		nodes = append(nodes, tail)
		tail.Next = &Node{Data: num}
		tail = tail.Next
	}

	rand.Seed(time.Now().UnixNano() + int64(os.Getpid()))
	max := len(nodes)

	for n := head; n != nil; n = n.Next {
		idx := rand.Intn(max)
		n.Extra = nodes[idx]
	}

	return
}

func Copy(head *Node) *Node {
	var reversed *Node

	// use this to keep track of newly created list nodes
	// by their .Data value
	refs := make(map[int]*Node)

	// Create a "backwards" new linked list,
	// keeping track of which data value the
	// extra pointer goes to
	for p := head; p != nil; p = p.Next {
		c := &Node{
			Data:  p.Data,
			Next:  reversed,
			Extra: p.Extra, // new list nodes point into original list
		}
		refs[c.Data] = c
		reversed = c
	}

	// Reverse the new linked list in place, replacing the
	// new list's nodes' Extra pointers so they point into
	// the new list.
	var tmp, cpy *Node
	for reversed != nil {
		tmp, reversed = reversed, reversed.Next
		// replace pointer to original with copy that has the same Data value
		tmp.Extra = refs[tmp.Extra.Data]
		tmp.Next, cpy = cpy, tmp
	}

	return cpy
}

func Draw(head *Node) {
	fmt.Printf("digraph g {\n")
	fmt.Printf("rankdir=\"LR\";\n")
	DrawPrefixed(head, "N")
	fmt.Printf("}\n")
}

// DrawPrefixed output GraphViz "dot" format, customized
// for this list implementation - the .Next links are weighted
// heavily to keep them in a straight line, while the .Extra
// links get routed around. The .Extra edges should come out dashed.
func DrawPrefixed(head *Node, prefix string) {
	for n := head; n != nil; n = n.Next {
		fmt.Printf("%s%p [label=\"%d\"];\n", prefix, n, n.Data)
	}
	fmt.Printf("edge [weight=100];\n")
	for n := head; n != nil; n = n.Next {
		if n.Next != nil {
			fmt.Printf("%s%p -> %s%p\n", prefix, n, prefix, n.Next)
		}
	}
	fmt.Printf("edge [weight=1,style=\"dashed\"];\n")
	for n := head; n != nil; n = n.Next {
		if n.Extra != nil {
			fmt.Printf("%s%p -> %s%p\n", prefix, n, prefix, n.Extra)
		}
	}
}
