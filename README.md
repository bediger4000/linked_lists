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

[Code](merge.go)

This is actually just a test for problems that require more than 1 list.

```
$ ./cp 1 3 0 12 -- 5 -- 6 7 8 9 -- 100 99 98

Found 4 lists
1 -> 3 -> 0 -> 12 -> 
5 -> 
6 -> 7 -> 8 -> 9 -> 
100 -> 99 -> 98 -> 
```

---

### Build a LIFO, a stack, from a linked list

[Code](stk.go)

```
$ ./stk 1 10 2 0
Empty: true
0 2 10 1 
```
---

### Daily Coding Problem: Problem #443 [Medium] 

This problem was asked by Apple.

Implement a queue using two stacks.
Recall that a queue is a FIFO (first-in, first-out)
data structure with the following methods:
enqueue, which inserts an element into the queue, and dequeue, which removes it.

```
$ ./queue 1 2 3 4 5
1 2 3 4 5
```

#### Analysis

[Code](list/queue.go)

This is Old School.
As far as I know, there's only one way to do this.
I'll refer to the two stacks as "enqueue stack" and "dequeue stack".
On a dequeue, if the dequeue stack is empty,
pop each element off the enqueue stack and push it on the dequeue stack,
leaving the enqueue stack empty.
On an enqueue, push the item on the enqueue stack.

If the candidate doesn't know this one, the problem is worthless.
Even if the candidate knows the answer, it's just some easy stack manipulation.
Pushing or popping an item from the head of a list is about as basic as things come.

---

### Merge two sorted lists

[Code](sort.go)

```
$ ./merge  10 6 2 0 -- 11 8 7 4 1
10 -> 6 -> 2 -> 0 -> 
11 -> 8 -> 7 -> 4 -> 1 -> 
11 -> 10 -> 8 -> 7 -> 6 -> 4 -> 2 -> 1 -> 0 -> 
```

The lists are in ascending order.
The obvious follow-on is merging N sorted lists.

---

### Find the middle item in a linked list

[Code](middle.go)

Set two `*Node` pointers to head of list.
One steps through list 2 at a time,
the other pointer steps one at a time.
When the 2x pointer gets to the end of the list,
the other pointer holds address of  middle item.
Except for even-number-length lists.
It's arguable what the "middle" of those lists is.

---

### Daily Coding Problem: Problem #559 [Medium] 

This problem was asked by Google.

Given k sorted singly linked lists, write a function to merge all the lists
into one sorted singly linked list.

I haven't tried this one. Apparently there's a way to take advantage of the
sorted lists so that it's not an O(N^2) algorithm.

---

### Reverse a linked list in place

This is: Daily Coding Problem: Problem #465 [Easy]

[Code](reverse.go)

This problem was asked by Google.

Given the head of a singly linked list, reverse it in-place.

This amounts to taking the head element off a linked list,
and setting that element's pointer to the reversed list.
Repeat until the linked list is empty, and all elements are on the reveresed list.

```
$ ./reverse 1 2 3 4
1 -> 2 -> 3 -> 4 -> 
4 -> 3 -> 2 -> 1 -> 
```

---

### Is a list palindromic?

[Code](palindrome.go)

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

---

### XOR linked list

[Code](xorlist.go)

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
The interviewer should expect some mild flailing from any candidate.
Look for candidates that can work this problem out carefully.

The candidate can't really do much to redeem themselves if they
don't get the nature of XOR.
If the interviewer notices this,
it might be worthwhile to help out on the XOR part,
unless you want to see a candidate's innate inspiration ability,
which doesn't exist in anybody.
One consequence of the XOR is that the `both` field of the head
of the list will contain the pointer to the 2nd element,
and the `both` field of the tail of the list will be the pointer
to the next-to-last-element.
This almost mandates special cases for inserting the first and second
elements of a list.
All intermediate list node `both` fields have some unnatural numeric value that
doesn't look like a pointer.

---

### Remove kth Last element

[Code](removekthlast.go)

This problem was asked by Google.

Given a singly linked list and an integer k,
remove the kth last element from the list.
k is guaranteed to be smaller than the length of the list.

The list is very long,
so making more than one pass is prohibitively expensive.

Do this in constant space and in one pass.

#### Analysis

This problem requires 2 pointers into the last,
one that leads, and a 2nd that's k+1 elements behind.
When the leader is null/nil,
the 2nd pointer has the node before the element to be
eliminated from the list.

Since "k is guaranteed to be smaller than the length of the list",
I'm not sure there are any corner cases worth talking about.
This is a very straightforward problem.
I guess if the interviewer is being picky about "in one pass",
[my solution](removekthlast.go) uses 2 for-loops,
one to get the lead pointer to the k-2'th place in the list,
and one to find the end of the list from there
and increment both lead and trailing pointers.

If you're interviewing for an entry-level position,
this might be worth asking.
Otherwise, I don't see it.

---

### Daily Coding Problem: Problem #699 [Easy]

This problem was asked by Airbnb.

Given a linked list and a positive integer k, rotate the list to the right by k places.

For example,
given the linked list 7 -> 7 -> 3 -> 5 and k = 2,
it should become 3 -> 5 -> 7 -> 7.

Given the linked list 1 -> 2 -> 3 -> 4 -> 5 and k = 3,
it should become 3 -> 4 -> 5 -> 1 -> 2.

#### Analysis

There's an error with the problem statement.
One of the two examples is incorrect.
First example gives k = 2, then gives the 3-valued node,
the third node in the list,
as the head of the rotated list.

The second example gives k = 3, then wants the 3rd element of the old
list as the head of the new list.

* k = 2, third node as head of rotated list.
* k = 3, 4th node as head of rotated list.

One or the other is incorrect.

I take the second example as incorrect.

I thought of two ways to do this.

* [First algorithm](rotate1.go)
  1. Find kth node of list, or return nil if list is too short.
  2. Find tail node of list. List is in original form at this point.
  3. Set tail.Next to head node. List is now circular.
  4. Find Node before kth node. This is modified list's tail node
  5. Set node-before-kth-node's Next element to nil
  6. Return kth node as head of rotated list
* [Second algorithm](rotate2.go)
  1. Find tail node of list.
  2. Set tail.Next to head of list, making a circular list.
  3. Move head and tail k items through list
  4. Set tail.Next to nil
  5. return head as head of rotated list

The second algorithm has the benefit that any list, even of length 1,
can be rotated any number of elements.
There's no problem with "too short" lists.

This isn't a bad problem for a whiteboard interview.
It has a data structure,
it's not something that candidates would have done a lot in the past.
It has pointers.
The candidate has to design an algorithm,
and there's corner cases like short list, zero length list,
rotate by more than list length,
rotate by zero
to take into account.

The mistaken problem statement could be a way to try out
a candidate's critical thinking skills.
Identifying bugs in requirements is necessary on the job.

If the candidate finds their attempts at clarification of the
problematic test case(s) rebuffed,
they should consider not interviewing further with that company.
That's a red flag about company culture and processes.

---

## Daily Coding Problem: Problem #714 [Easy]

This problem was asked by Google.

Given the head of a singly linked list,
swap every two nodes and return its head.

For example,
given `1 -> 2 -> 3 -> 4`, return `2 -> 1 -> 4 -> 3`.

### Analysis

[Code](swap.go)

This is somewhat harder than the usual linked list question.
You have to keep track of 3 pointers,
the two list nodes that switch positions,
and something that
gets set to the new-first-of-two-swapped-nodes.

This last "something" is the hard part.
Once the program has swapped Next pointers for two nodes,
the Next pointer of the node previous to those two
still points to what's now the second of the two swapped nodes.
I chose to use a pointer-to-a-pointer,
storing the address of the Next pointer that needs updating
once the next two nodes get swapped.

### Interview analysis

Because of needing to keep track of a 3rd pointer,
the Next pointer that needs to point to a swapped node,
I'd say this is more of a "medium" than an easy.
It's not a bad problem for a linked list question,
which tend to the basic,
and demand pointer swapping book-keeping more than anything else.

If I got this in an interview,
I'd point out the necessity of having this 3rd pointer around,
because I certainly wouldn't get it correct without some experimentation.
In fact, I tried to do this problem in a single sitting
to simulate a whiteboard interview.
I failed because I didn't immediately see the necessity of the 3rd pointer.
It was only after literally sleeping on the problem overnight
that I see the need for the 3rd pointer.

Even if the candidate hoses up the programming,
test cases like zero-length-list, single item list,
odd-number-of-nodes lists should get extra credit from the interviewer.

The interviewer should probably be ready to give a hint
about the 3rd pointer, although I'm not sure what hint I'd give
if I threw this problem out for discussion or whiteboarding.

---
## Daily Coding Problem: Problem #715 [Easy]

This problem was asked by Google.

Determine whether a doubly linked list is a palindrome.
What if it's singly linked?

For example,
1 -> 4 -> 3 -> 4 -> 1 returns True
while 1 -> 4 returns False.

### Analysis

Palindromic singly-linked-lists are [above](#is-a-list-palindromic).

[Code](dlpalindrome.go) for doubly-linked list palindrome check.

I chose to find the tail of the doubly-linked list,
then walk the list forward and backward comparing node data values.
This is conceptually easier than the singly-linked list method
of walking halfway while making a reversed-in-place linked list,
then comparing node data values from the middle of the list
to the ends of the list.

Both singly and doubly-linked list methods visit 2N nodes to do the check.
For the doubly-linked list method,
there are no ugly pointer manipulations to create a reversed-in-place linked list,
so no need to reverse the half of the list reversed-in-place
to get the list back to its original state.

As an interview question, this is OK.
It probably even merits an "easy" rating.
The candidate can do the exact same method as for a singly-linked list,
but this wouldn't demonstrate knowledge of doubly-linked lists.
Comparing algorithms for the two types of lists is probably where the
candidate can show knowledge and competency,
and where the interviewer could get a better feel for the candidate's ability.

---

## Cracking the Coding Interview

Linked list questions.

## Interview Question 2.3

[Code](removeitem.go)

[Another version](removeitem2.go)

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

---

### Interview Question 2.8

Given a circular linked list, implement an algorithm that returns
the node at the beginning of the loop.

[Code](cycletest.go)

#### Circular linked list

A (corrupt) linked list in which a node's next pointer points
to an earlier node, so as to make a loop in the linked list.

#### Example

* Input: a -> b -> c -> d -> e -> c
* Output: c

#### Analysis

This isn't the usual definition of a "circular linked list".
This is just a list with (mistakenly) a cycle in it.
