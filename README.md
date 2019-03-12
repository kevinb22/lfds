# lfds
Implementations of simple lock free data structures in golang and c, because sanity and safety are overrated.

## Members
Kevin Bi (kevinb22@uw.edu) and Anirban Biswas (anirban@uw.edu)

## Lock free Hash Map
To test the code, from the `c/src` directory,
```bash
$ make
$ ./lf_map_test
```

This is an implementation of a lock free hash map inspired by the generation number approach of lock free datastructures in the Linux Scalability paper. When adding entires to the hash map, the version number is checked before modifying the hash table bucket to ensure no changes were made while elements in the linked list were being compared. If the list was modified in that duration, the version number would have been updated, and the insert would fail and be restarted.

`lf_map_get()` checks the version of the bucket it is searching in every time it inspects a node. If the version number has changed from when it first started searching the list, it starts looking from the head of the list again since the list was modified. The guarantees provided by `lf_map_get()` are weak. It is possible, that after finding the node in the linked list that has the key we are looking for, the value of the key gets updated. 

### Lock Free Stack and Queue
#### Setup workspace for go
```bash
$ export GOPATH=`pwd`
$ export GOBIN=$GOPATH/bin
$ PATH=$PATH:$GOPATH:$GOBIN
$ export PATH
```
#### Run go tests
To test the go code navigate to the lfstructures directory and run the following command:
```bash
$ go test
```
To test run an individual test run the following command:
```bash
$ go test -run NameOfTest
```
These are golang implementations of a Trieber Stack and a One Producer/One Consumer Queue inspired by various online articles. Go was selected because it was a (relatively) high level language that was not verbose and had atomic functions. After implementing the data structures it is evident why Go prefers the use of channels. For these data structures we relied on the previously mentioned atomic functions as well as unsafe.Pointer types which function similarly to C pointers.

* Both data structures use a common Node struct and Container interface. The Node struct contains two fields a value, Container interface to store data and next, a pointer to another Node. Nodes are intended to be chained together to make a linked list. The Container interface is used to give the Node the ability to store multiple data types (essentially replicating generica in Java since golang does not have generics). The container has a a Get() and Put() method to add to or remove the data it wraps.

* The Trieber Stack is implemented as a linked list under the hood and makes use of go's atomic.CompareAndSwapPointer, atomic.LoadPointer, and atomic.StorePointer functions to push and pop nodes to the head of the linked list which is a node field called Top. `Push()` function atomically gets a pointer to the head of the list and remembers this pointer. It allocates a new node and sets the nodes next field to the head. It then uses attempts to use atomic.CompareAndSwapPointer to update the head of the list to the new node. If this action fails it re-attempts with the new head of the list. `Pop()` works similarly except it attempts to grab the next node of the current head and reset the head of the list to the next node.

* The One Producer/One Consumer Queue allows two threads to work concurrently on the datastructure, which is also implemented as a linked list under the hood. It has three fields that are nodes, First, Divider, and Last. One thread, the producer, is allowed to add nodes to the Queue via `Produce()`, the other, the consumer, will remove nodes from the Queue via `Consume()`. The Divider field ensures that the producer and consumer threads never overstep their boundaries. The correctness of these operations is also guaranteed by use of go's atomic.CompareAndSwapPointer, atomic.LoadPointer, and atomic.StorePointer functions to append nodes to the tail of the linked list, trim the head of the linked list, and increment the divider node. Having more than one producer or consumer will break the model and cause race conditions and other errosr.

### Inspiration and References
* [An Analysis of Linux Scalability to Many Cores](https://courses.cs.washington.edu/courses/cse551/19wi/readings/mosbench-osdi10.pdf)
* [Writing Lock-Free Code: A Corrected Queue](http://www.drdobbs.com/parallel/writing-lock-free-code-a-corrected-queue/210604448?pgno=1)
* [Fear and Loathing in Lock-Free Programming](https://medium.com/@tylerneely/fear-and-loathing-in-lock-free-programming-7158b1cdd50c)
