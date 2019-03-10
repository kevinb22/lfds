# lfds
Implementations of simple lock free data structures in golang and c, because sanity and safety are overrated.

### Setup workspace for go
```bash
$ export GOPATH=`pwd`
$ export GOBIN=$GOPATH/bin
$ PATH=$PATH:$GOPATH:$GOBIN
$ export PATH
```

### Inspiration and References
[Writing Lock-Free Code: A Corrected Queue](http://www.drdobbs.com/parallel/writing-lock-free-code-a-corrected-queue/210604448?pgno=1)
[Fear and Loathing in Lock-Free Programming](https://medium.com/@tylerneely/fear-and-loathing-in-lock-free-programming-7158b1cdd50c)