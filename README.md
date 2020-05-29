# Linked List Interview Questions

A compendium of linked list developer interview questions in Go

## Building

```go
$ go build cp.go
$ go build merge.go
$ go build stk.go
```

## Running

### Merge two sorted lists

```
$ ./merge  10 6 2 0 -- 11 8 7 4 1
10 -> 6 -> 2 -> 0 -> 
11 -> 8 -> 7 -> 4 -> 1 -> 
11 -> 10 -> 8 -> 7 -> 6 -> 4 -> 2 -> 1 -> 0 -> 
```

The lists are in ascending order.
The obvious follow-on is merging N sorted lists.

### Read in multiple lists

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

The follow-on here is to build a LIFO, a queue, from 2 linked lists.
