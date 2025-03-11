package ca

import (
	"math/rand"
	"time"
)

// Grid that organizes all agents.
type Grid []*Agent

// NewGrid Initialize the grid of size n.
func NewGrid(n int) (me Grid) {
	me = make(Grid, n)
	return
}

// Add adds a reference to an agent at position i on the grid.
func (me Grid) Add(agent *Agent, i int) {
	me[i] = agent
	agent.grid = me
}

// Shuffle shuffles the grid, and initialize sources and targets.
func (me Grid) Shuffle() {
	// initialize local pseudorandom generator
	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(me), func(i, j int) {
		me[i], me[j] = me[j], me[i]

		me[i].SetSource(j)
		me[j].SetSource(i)

		me[i].SetTarget(i)
		me[j].SetTarget(j)
	})
}

// Shuffled returns the grid indexes randomly sorted.
func (me Grid) Shuffled() []int {
	// initialize local pseudorandom generator
	rand.Seed(time.Now().UnixNano())

	var seq = make([]int, len(me))
	var i = 0
	for i = range seq {
		seq[i] = i
	}

	// sort agents by time of transaction
	rand.Shuffle(len(seq), func(i, j int) {
		seq[i], seq[j] = seq[j], seq[i]
	})

	return seq
}
