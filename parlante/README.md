# Nick Parlante's Linked List Problems

http://cslibrary.stanford.edu/105/LinkedListProblems.pdf

Parlante's reasons for studying linked lists and doing problems

## 1 - Count()

Write a `Count()` function that counts the number of times a given
`int` occurs in a list.

[Code](count.go)

```
$ go build count.go
$ ./count 5 12 11 20 5 6 55 5
Found 2 nodes with value 5
```

That is, the first argument ("5") is a string representation
of the `int` to count in a list whose data values are the
remainder of the integer string representations on the command line.

## 2 - GetNth()

Write a `GetNth()` function that takes a linked list and integer index and returns
the data value stored at that index position.
`GetNth()` uses the C numbering convention that the first node is index 0,
the second nodes is index 1, ... and so on.
So for the list {42, 13, 666} `GetNth()` with index 1 should return 13.
The index should be in the range [0, length-1].
If it is not, `GetNth()` should assert() fail (or you could implement some other error case strategy).

[Code](getnth.go)

```
$ go build getnth.go
$ ./getnth 5 12 11 20 5 6 55 5
Found data value 55 at index 5
```

The first argument ("5") is a string representation
of the list index to find in a list whose data values are the
remainder of the integer string representations on the command line.
List node at index 0 has data value 12.

## 3 - DeleteList()

I'm doing this in Go, no such function can be implemented.

## 4 - Pop()

Write a `Pop()` function that's the opposite of `Push()`.
`Pop()` takes a non-empty list, deletes the head node, and returns the head node's data.
`Pop()` should assert() fail if there is not a node to pop.

[Code](pop.go)

```
$ go build pop.go
$ ./pop  3 2 1
3 -> 2 -> 1 -> 
Head data was 3
2 -> 1 -> 
```

## 5 - InsertNth()

A more difficult problem is to write a function `InsertNth()`
which can insert a new node at any index within a list.
`Push()` is similar,
but can only insert a node at the head end of the list (index 0).
The caller may specify any index in the range [0..length],
and the new node should be inserted so as to be at that index.

[Code](insertnth.go)

```
$ go build insertnth.go
$ ./insertnth 2 55 0 1 2 3 4 5
0 -> 1 -> 2 -> 3 -> 4 -> 5 -> 
0 -> 1 -> 55 -> 2 -> 3 -> 4 -> 5 -> 
```

The first command line argument ("2") is the index at which the
second command line argument ("55") should appear in a list which
gets composed from the remaining command line arguments in order.

## 6 - SortedInsert()

Write a `SortedInsert()` function
which given a list that is sorted in increasing order,
and a single node,
inserts the node into the correct sorted position in the list.
While `Push()` allocates a new node to add to the list,
`SortedInsert()` takes an existing node,
and just rearranges pointers to insert it into the list.
There are many possible solutions to this problem.

[Code](sortedinsert.go)

```
$ go build sortedinsert.go
$ ./sortedinsert 5 0 1 3 9 27
0 -> 1 -> 3 -> 9 -> 27 -> 
0 -> 1 -> 3 -> 5 -> 9 -> 27 -> 
```

The first command line argument ("5" in the example above) is a
string representation of a node's data value.
The rest of the command line arguments get converted into a list.
The string representations must appear in increasing order on the command line.

## 7 - InsertSort()

Write an `InsertSort()` function which given a list,
rearranges its nodes so they are sorted in increasing order.
It should use `SortedInsert()`.

[Code](insertsort.go)

```
$ go build insertsort.go
$ ./insertsort 1 99 0 2 11 4 3
1 -> 99 -> 0 -> 2 -> 11 -> 4 -> 3 -> 
0 -> 1 -> 2 -> 3 -> 4 -> 11 -> 99 -> 
```

## 8 - Append()

Write an `Append()` function that takes two lists,
'a' and 'b', appends 'b' onto the end of 'a',
and then sets 'b' to NULL (since it is now trailing off the end of 'a').

[Code](append.go)

```
$ go build append.go
$ ./append  1 2 3 -- 4 5 6
a: 4 -> 5 -> 6 -> 
b: 100 -> 101 -> 102 -> 

a: 4 -> 5 -> 6 -> 100 -> 101 -> 102 -> 
b: 
$ ./append -- 4 5 6
a: 
b: 4 -> 5 -> 6 -> 

a: 4 -> 5 -> 6 -> 
b: 
```


## 9 - FrontBackSplit()

Given a list, split it into two sublists - one for the front half,
and one for the back half.
If the number of elements is odd,
the extra element should go in the front list.
So `FrontBackSplit()` on the list `{2, 3, 5, 7, 11}`
should yield the two lists `{2, 3, 5}`
and `{7, 11}`.
Getting this right for all the cases is harder than it looks.
You should check your solution against a few cases
(length = 2, length = 3, length=4) to make sure that the list
gets split correctly near the short-list boundary conditions.
If it works right for length=4,
it probably works right for length=1000.
You will probably need special case code to deal with the (length <2) cases.

Hint. Probably the simplest strategy is to compute the length of the list,
then use a for loop to hop over the right number of nodes to find the last
node of the front half,
and then cut the list at that point.
There is a trick technique that uses two pointers to traverse the list.
A "slow" pointer advances one nodes at a time,
while the "fast" pointer goes two nodes at a time.
When the fast pointer reaches the end,
the slow pointer will be about half way.
For either strategy, care is required to split the list at the right point.

[Code](frontbacksplit.go)

```
$ go build frontbacksplit.go
$ ./frontbacksplit 2 3 5 7 11
source list: 2 -> 3 -> 5 -> 7 -> 11 -> 

source list: 
Front list: 2 -> 3 -> 5 -> 
Back list:  7 -> 11 -> 
```

## 10 - RemoveDuplicates()

Write a `RemoveDuplicates()` function
which takes a list sorted in increasing order and
deletes any duplicate nodes from the list.
Ideally, the list should only be traversed once.

[Code](removeduplicates.go)

```
$ go build removeduplicates.go
$ ./removeduplicates.go 1 1 2 2 3 3 4 4 5
source list: 1 -> 1 -> 2 -> 2 -> 3 -> 3 -> 4 -> 4 -> 5 -> 

deduped list: 1 -> 2 -> 3 -> 4 -> 5 -> 
```

## 11 - MoveNode()

This is a variant on Push().
Instead of creating a new node and pushing it onto the given list,
`MoveNode()` takes two lists,
removes the front node from the second list
and pushes it onto the front of the first.
This turns out to be a handy utility function to have
for several later problems.
Both `Push()` and `MoveNode()` are designed
around the feature that list operations work most naturally at the head of the list.

[Code](movenode.go)

```
$ go build movenode.go
$ ./movenode  1 2 3 -- 4 5 6
a: 1 -> 2 -> 3 -> 
b: 4 -> 5 -> 6 -> 
After MoveNode
a: 2 -> 3 -> 
b: 1 -> 4 -> 5 -> 6 -> 
```

## 12 - AlternatingSplit()

Write a function `AlternatingSplit()` that takes one list
and divides up its nodes to make two smaller lists.
The sublists should be made from alternating elements in the original list.
So if the original list is {a, b, a, b, a},
then one sublist should be {a, a, a} and the
other should be {b, b}.
You may want to use `MoveNode()` as a helper.
The elements in the new lists may be in any order
(for some implementations, it turns out to be convenient
if they are in the reverse order from the original list.)

[Code](alternatingsplit.go)

```
$ go build alternatingsplit.go
$ ./alternatingsplit 1 1 2 2 3 3 4 4 5 5
source list: 1 -> 1 -> 2 -> 2 -> 3 -> 3 -> 4 -> 4 -> 5 -> 5 -> 

a list: 
5 -> 4 -> 3 -> 2 -> 1 -> 
b list: 
5 -> 4 -> 3 -> 2 -> 1 -> 
```

## 13 - ShuffleMerge()

Given two lists,
merge their nodes together to make one list,
taking nodes alternately between the two lists.
So `ShuffleMerge()` with {1, 2, 3} and {7, 13, 1}
should yield {1, 7, 2, 13, 3, 1}.
If either list runs out, all the nodes should be taken from the other list.

[Code](shufflemerge.go)

```
$ go build shufflemerge.go
$ ./shufflemerge 1 2 3 7 8 9 -- 4 5 6
a: 1 -> 2 -> 3 -> 7 -> 8 -> 9 -> 
b: 4 -> 5 -> 6 -> 
merged: 1 -> 4 -> 2 -> 5 -> 3 -> 6 -> 7 -> 8 -> 9 ->
```

## 14 - SortedMerge()

Write a `SortedMerge()` function that takes two lists,
each of which is sorted in increasing order,
and merges the two together into one list which is in increasing order.
`SortedMerge()` should return the new list.
The new list should be made by splicing
together the nodes of the first two lists (use `MoveNode()`).
Ideally, `Merge()` should only make one pass through each list.
`Merge()` is tricky to get right -
it may be solved iteratively or recursively.
There are many cases to deal with: either 'a' or 'b' may be empty,
during processing either 'a' or 'b' may run out first,
and finally there's the problem of starting the result list empty,
and building it up while going through 'a' and 'b'.

```
$ go build sortedmerge.go
$ ./sortedmerge 1 3 5 -- 0 4 7
a: 1 -> 3 -> 5 -> 
b: 0 -> 4 -> 7 -> 
merged: 0 -> 1 -> 3 -> 4 -> 5 -> 7 -> 
```

## 15 - MergeSort()

(This problem requires recursion) Given `FrontBackSplit()` and `SortedMerge()`,
it's pretty easy to write a classic recursive `MergeSort()`:
split the list into two smaller lists,
recursively sort those lists,
and finally merge the two sorted lists together into a single sorted list.
Ironically, this problem is easier than either `FrontBackSplit()`
or `SortedMerge()`.

I spent the summer of 2024 [investigating mergesort](/mergesort)

## 16 - SortedIntersect()

Given two lists sorted in increasing order,
create and return a new list representing the intersection of the two lists.
The new list should be made with its own memory -
the original lists should not be changed.
In other words,
this should be `Push()` list building, not `MoveNode()`.
Ideally, each list should only be traversed once.
This problem, along with `Union()` and `Difference()`
form a family of clever algorithms that exploit the
constraint that the lists are sorted to find common nodes efficiently.

[Code](sortedintersect.go]

```
$ go build sortedinterset.go
$ ./sortedinterset 1 3 5 -- 0 4 7
```

## 17 - Reverse()

Write an iterative `Reverse()` function
that reverses a list
by rearranging all the `.next` pointers and the head pointer.
Ideally, `Reverse()` should only need to make one pass of the list.
The iterative solution is moderately complex.

[Code](reverse.go]

```
$ go build reverse.go
$ ./reverse  1 2 3 4
source list: 1 -> 2 -> 3 -> 4 -> 

Reversed:    4 -> 3 -> 2 -> 1 -> 

```

## 18 - RecursiveReverse()

(This problem is difficult and is only possible if you are familiar with recursion.)
There is a short and efficient recursive solution to this problem.
As before, the code should only make a single pass over the list.
Doing it with multiple passes is easier but very slow,
so here we insist on doing it in one pass.
Solving this problem requires a real understanding of pointer code and recursion.

```
$ go build recursivereverse.go
$ ./recursivereverse 1 2 3 4
source list: 1 -> 2 -> 3 -> 4 -> 

Reversed:    4 -> 3 -> 2 -> 1 -> 
```
