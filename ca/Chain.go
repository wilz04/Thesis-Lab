package ca

// Supplier is the supplychain.
type Supplier struct {
	chain []int

	value int
	axis  int
}

// NewChain Initialize the supplychain.
func NewChain(i int) (me *Supplier) {
	me = &Supplier{
		chain: []int{i},
		value: -MAX,
	}
	return
}

// IndexOf returns the position of i in chain.
func (me *Supplier) IndexOf(i int) int {
	var j, k = 0, 0
	for k, j = range me.chain {
		if j == i {
			return k
		}
	}

	return -1
}

// Push adds i at the end of a copy of chain and returns this copy.
func (me *Supplier) Push(i int) {
	var mycopy = make([]int, len(me.chain)+1)
	copy(mycopy, me.chain)
	copy(mycopy[len(me.chain):], []int{i})
	me.chain = mycopy
}

// Last returns the last agent ID.
func (me *Supplier) Last() int {
	return me.chain[len(me.chain)-1]
}

// LastNd returns the second last agent ID.
func (me *Supplier) LastNd() int {
	if len(me.chain) < 2 {
		return -1
	}

	return me.chain[len(me.chain)-2]
}
