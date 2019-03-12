# lfds
Implementations of simple lock free data structures in golang and c, because sanity and safety are overrated.

### Setup workspace for go
```bash
$ export GOPATH=`pwd`
$ export GOBIN=$GOPATH/bin
$ PATH=$PATH:$GOPATH:$GOBIN
$ export PATH
```

## Lock free Hash Map
To test the code, from the `c/src` directory,
```bash
$ make
$ ./lf_map_test
```

This is an implementation of a lock free hash map inspired by the generation number approach of lock free datastructures in the Linux Scalability paper. When adding entires to the hash map, the version number is checked before modifying the hash table bucket to ensure no changes were made while elements in the linked list were being compared. If the list was modified in that duration, the version number would have been updated, and the insert would fail and be restarted.

`lf_map_get()` checks the version of the bucket it is searching in every time it inspects a node. If the version number has changed from when it first started searching the list, it starts looking from the head of the list again since the list was modified. The guarantees provided by `lf_map_get()` are weak. It is possible, that after finding the node in the linked list that has the key we are looking for, the value of the key gets updated. 


### Inspiration and References
* [Writing Lock-Free Code: A Corrected Queue](http://www.drdobbs.com/parallel/writing-lock-free-code-a-corrected-queue/210604448?pgno=1)
* [Fear and Loathing in Lock-Free Programming](https://medium.com/@tylerneely/fear-and-loathing-in-lock-free-programming-7158b1cdd50c)
