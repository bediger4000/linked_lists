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

### Is a list palindromic?

Given a singly linked list of integers,
determine whether or not it's a palindrome.

My algorithm uses the "finding the middle element of a list" and "reversing a list in place" tasks.
Ruins the original list in order to find out if that list qualifies as a palindrome,
although the code could put the original list back together, if so required.

This might make a good interview question if you have a lot of time to check coding ability.
Any "check palindrome" algorithm will integrate other algorithms (in my case, reverse-in-place,
and find-the-middle-item) to accomplish a larger task.
This problem would let an interviewer know that a candidate could analyze a problem,
breaking it into smaller pieces,
then integrate the smaller pieces' results into a whole.
The problem demands design skills, as well as small algorithm knowledge.
The downsides are that there's a couple of gotchas:
finding the "middle" of a list could bog down the candidate down
in deciding which node to use as the middle of an even-number-sized list.
The case of a palindrome with a single unique element in the middle
(which will only occur in an odd-number-sized list) might ruin the candidate's efforts.

### XOR linked list

Daily Coding Problem: Problem #590 [Hard]

This problem was asked by Google.

An XOR linked list is a more memory efficient doubly linked list. Instead
of each node holding next and prev fields, it holds a field named both,
which is an XOR of the next node and the previous node. Implement an XOR
linked list; it has an add(element) which adds the element to the end, and
a get(index) which returns the node at index.

If using a language that has no pointers (such as Python), you can assume
you have access to get_pointer and dereference_pointer functions that
converts between nodes and memory addresses.

#### Analysis of XOR linked list

Some programming languages will be more conducive to getting this
problem correct than others.
The problem statement acknowledges that when it says to assume
some magical functions if you're doing it in Python.
I chose Go, which has pretty strict type safety,
and memory safety even though it has pointers.
The combination of the two safeties causes my code to look a bit odd.

The Go compiler allows the programmer to convert any pointer address
to a pointer of an arbitrary type
via the semi-magic function `unsafe.Pointer()`.
One can then convert the pointer of aribtrary type to a
numerical value of type `uintptr`
Go will do bitwise operations like Xor on `uintptr` values.
So my code converts pointers to `uintptr` numerical values and back
a lot.

I've done this problem (in C) as an exercize in the past,
because it just seemed so outlandish.
I think that someone who was just informed of this very hacky
idea (node carries prev XOR next as a single field)
might have a lot of trouble with it.
`node->next` and `node->prev` are one thing,
but combining them via XOR seems to add cognitive burden to the problem.

I think this problem doesn't merit the "hard" rating
for a candidate who has experimented with XOR operations
in the past, but to someone (entry-level or even mid-level)
who hasn't explored XOR, it might seem very difficult.
Keeping track of 3 pointers (current, previous, next) is error-prone
even in regular doubly linked list operations,
the XOR just adds to the difficulty,
as would the shenanigans working around Go's type- and memory-safety.
The programming language chosen should influence the interviewer's
final judgement.
The interviewer should expect some mild flailing from any candidate,
but should look for candidates that work it out carefully.

The candidate can't really do much to redeem themselves if they
don't get the nature of XOR.
One consequence of the XOR is that the `both` field of the head
of the list will contain the pointer to the 2nd element,
and the `both` field of the tail of the list will be the pointer
to the next-to-last-element.
This almost mandates special cases for inserting the first and second
elements of a list.
All intermediate list node have some unnatural numeric value that
doesn't look like a pointer.

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


