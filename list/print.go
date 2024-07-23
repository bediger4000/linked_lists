package list

import (
	"fmt"
	"io"
	"sort"
	"unsafe"
)

// Print runs a linked list and prints its values on stdout
func Print(list *Node) {
	for node := list; node != nil; node = node.Next {
		fmt.Printf("%d -> ", node.Data)
	}
}

// PrintN runs a linked list, printing first n values on stdout
func PrintN(list *Node, n int) {
	count := 0
	var node *Node
	for node = list; count < n && node != nil; node = node.Next {
		fmt.Printf("%d -> ", node.Data)
		count++
	}
	if node != nil {
		fmt.Print("...")
	}
}

type addressList []uint

func (l addressList) Len() int           { return len(l) }
func (l addressList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l addressList) Less(i, j int) bool { return l[i] < l[j] }

// Grint puts a dot-language GraphViz format representation
// of a linked list on fout. Nodes are in memory-address-order,
// low addresses to left, high to right, in a row. .Next contents
// represented by a bold arrow to that "next" node.
func Grint(fout io.Writer, list *Node) {

	fmt.Fprintf(fout, "digraph g {")
	fmt.Fprintf(fout, "rankdir=\"LR\";")
	fmt.Fprintf(fout, "node [wdith=0.1,height=0.1,pin=true];\n")

	var addresses addressList

	// gather addresses
	for node := list; node != nil; node = node.Next {
		addresses = append(addresses, uint(uintptr(unsafe.Pointer(node))))
	}

	// sort addresses low to high
	sort.Sort(addresses)

	// Output Node structs in sorted address order.
	// Go heap allocator seems to use 8192-byte chunks, but the chunks
	// may not be allocated in low-to-high fashion themselves.
	pos := 0
	for _, address := range addresses {
		node := (*Node)(unsafe.Pointer(uintptr(address)))
		fmt.Fprintf(fout, "n%p [label=\"%d\",pos=\"%d,0\"];\n", node, node.Data, pos)
		pos++
	}

	// invisible edges to keep nodes lined up vertically
	fmt.Fprintf(fout, "edge [weight=200,arrowhead=none,style=\"invis\"];\n")
	spacer := ""
	for _, addr := range addresses {
		fmt.Fprintf(fout, "%sn0x%x", spacer, addr)
		spacer = " ->\n"
	}
	fmt.Fprintf(fout, ";")

	// .Next pointers as bold arrowhead edges
	spacer = ""
	fmt.Fprintf(fout, "edge [weight=1,arrowhead=normal,style=\"bold\"];\n")
	for node := list; node != nil; node = node.Next {
		fmt.Fprintf(fout, "%sn%p", spacer, node)
		spacer = " ->\n"
	}
	fmt.Fprintf(fout, ";\n")

	fmt.Fprintf(fout, "}\n")
}
