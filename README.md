# Linked List Interview Questions

A compendium of linked list developer interview questions in Go

## Building

```go
$ go build cp.go
$ go build merge.go
$ go build stk.go
$ go build queue.go
$ go build reverse.go
```

## Problems and puzzles

### Read in multiple lists

This is actually just a test for problems that require more than 1 list.

```
$ ./cp 1 3 0 12 -- 5 -- 6 7 8 9 -- 100 99 98

Found 4 lists
1 -> 3 -> 0 -> 12 -> 
5 -> 
6 -> 7 -> 8 -> 9 -> 
100 -> 99 -> 98 -> 
```

### Build a FIFO, a stack, from a linked list

```
$ ./stk 1 10 2 0
Empty: true
0 2 10 1 
```

### Build a LIFO, a queue from 2 linked lists

Build a LIFO, a queue, from 2 linked lists.

```
$ ./queue 1 2 3 4 5
1 2 3 4 5
```

### Merge two sorted lists

```
$ ./merge  10 6 2 0 -- 11 8 7 4 1
10 -> 6 -> 2 -> 0 -> 
11 -> 8 -> 7 -> 4 -> 1 -> 
11 -> 10 -> 8 -> 7 -> 6 -> 4 -> 2 -> 1 -> 0 -> 
```

The lists are in ascending order.
The obvious follow-on is merging N sorted lists.

### Find the middle item in a linked list

Set two `*Node` pointers to head of list.
One steps through list 2 at a time,
the other pointer steps one at a time.
When the 2x pointer gets to the end of the list,
the other pointer holds address of  middle item.
Except for even-number-length lists.
It's arguable what the "middle" of those lists is.

### Daily Coding Problem: Problem #559 [Medium] 

This problem was asked by Google.

Given k sorted singly linked lists, write a function to merge all the lists
into one sorted singly linked list.

I haven't tried this one. Apparently there's a way to take advantage of the
sorted lists so that it's not an O(N^2) algorithm.

### Reverse a linked list in place

This amounts to taking the head element off a linked list,
and setting that element's pointer to the reversed list.
Repeat until the linked list is empty, and all elements are on the reveresed list.

```
$ ./reverse 1 2 3 4
1 -> 2 -> 3 -> 4 -> 
4 -> 3 -> 2 -> 1 -> 
```
## Cracking the Coding Interview

Linked list questions.

## Interview Question 2.3

Implement an algorithm to delete a node in the middle
(i.e., any node but the first and last node, not necessarily
the exact middle) of a singly linked list, given only access
to that node

### Example

Input: remove the node c from the linked list

    a->b->c->d->e->f

Result: nothing is returned, but the new linked list looks like

    a->b->d->e->f

### Analysis

This is a strange question.
"Given only access to that node" implies a function that takes a `*Node`,
pointing to the node you want to remove from the list.

You don't have a pointer to the previous node,
which is what you'd need to surgically remove a node from a list.
The best you can do is copy the remaining list items' data "one back",
and trim off the final node.
This is going to leave dangling pointers in languages without garbage collection
unless you're very careful.

### Interview Question 2.8

Given a circular linked list, implement an algorithm that returns
the node at the beginning of the loop.

#### Circular linked list

A (corrupt) linked list in which a node's next pointer points
to an earlier node, so as to make a loop in the linked list.

#### Example

* Input: a -> b -> c -> d -> e -> c
* Output: c

#### Analysis

This isn't the usual definition of a "circular linked list".
This is just a list with (mistakenly) a cycle in it.


