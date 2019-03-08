package lfstructures

// Node struct that builds up TrieberStack
type Node struct {
	Value Container
	Next  *Node
}

// Container is a generic container, accepting anything.
type Container []interface{}

// Put adds an element to the container.
func (c *Container) Put(elem interface{}) {
	*c = append(*c, elem)
}

// Get gets an element from the container.
func (c *Container) Get() interface{} {
	elem := (*c)[0]
	*c = (*c)[1:]
	return elem
}
