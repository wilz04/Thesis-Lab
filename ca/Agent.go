package ca

import (
	"math"
	"math/rand"
	"time"
)

// MAX is the max agent value
const MAX = 256

// REWARDENABLED enable or disable the reward
var REWARDENABLED = false

// Agent is the agent structure
type Agent struct {
	ID int

	V0 int // \upsilon_i
	V1 int // \upsilon'_i
	V2 int // \upsilon''_i
	W0 int // \psi_i
	W1 int // \psi'_i
	W2 int // \psi''_i

	R0 int
	R1 int
	R2 int

	source *Color
	target *Color

	xSupplier *Supplier
	ySupplier *Supplier
	zSupplier *Supplier

	grid    Grid
	history History
}

// LastV0 returns the last V0 in history
func (me *Agent) LastV0() int {
	var i = len(me.history) - 1
	if i >= 0 {
		return me.history[i].v0
	}

	return 0
}

// LastW0 returns the last W0 in history
func (me *Agent) LastW0() int {
	var i = len(me.history) - 1
	if i >= 0 {
		return me.history[i].w0
	}

	return 0
}

// LastV1 returns the last V1 in history
func (me *Agent) LastV1() int {
	var i = len(me.history) - 1
	if i >= 0 {
		return me.history[i].v1
	}

	return 0
}

// LastW1 returns the last W1 in history
func (me *Agent) LastW1() int {
	var i = len(me.history) - 1
	if i >= 0 {
		return me.history[i].w1
	}

	return 0
}

// LastV2 returns the last V2 in history
func (me *Agent) LastV2() int {
	var i = len(me.history) - 1
	if i >= 0 {
		return me.history[i].v2
	}

	return 0
}

// LastW2 returns the last W2 in history
func (me *Agent) LastW2() int {
	var i = len(me.history) - 1
	if i >= 0 {
		return me.history[i].w2
	}

	return 0
}

// LastC0 returns the last C0 in history
func (me *Agent) LastC0() int {
	var i = len(me.history) - 1
	if i >= 0 {
		return me.history[i].c0
	}

	return 0
}

// LastC1 returns the last C1 in history
func (me *Agent) LastC1() int {
	var i = len(me.history) - 1
	if i >= 0 {
		return me.history[i].c1
	}

	return 0
}

// LastC2 returns the last C2 in history
func (me *Agent) LastC2() int {
	var i = len(me.history) - 1
	if i >= 0 {
		return me.history[i].c2
	}

	return 0
}

// XSupplier returns the supplier for X
func (me *Agent) XSupplier() int {
	if me.xSupplier != nil {
		return me.xSupplier.Last()
	}

	return -1
}

// YSupplier returns the supplier for Y
func (me *Agent) YSupplier() int {
	if me.ySupplier != nil {
		return me.ySupplier.Last()
	}

	return -1
}

// ZSupplier returns the supplier for Z
func (me *Agent) ZSupplier() int {
	if me.zSupplier != nil {
		return me.zSupplier.Last()
	}

	return -1
}

// NewAgent initializes a new agent according to its class
func NewAgent(class int) (me *Agent) {
	me = &Agent{
		xSupplier: nil,
		ySupplier: nil,
		zSupplier: nil,

		source: NewColor(),
		target: NewColor(),

		history: NewHistory(),
	}

	// initialize local pseudorandom generator
	rand.Seed(time.Now().UnixNano())

	var d0, d1, d2 int
	for {
		me.V0 = rand.Intn(MAX)
		me.W0 = rand.Intn(MAX)
		me.V1 = rand.Intn(MAX)
		me.W1 = rand.Intn(MAX)
		me.V2 = rand.Intn(MAX)
		me.W2 = rand.Intn(MAX)

		d0 = me.GetDelta0()
		d1 = me.GetDelta1()
		d2 = me.GetDelta2()
		/*
			if
				(me.V0 > 0 && me.V1 > 0 && me.V2 > 0) &&
				(me.W0 > 0 && me.W1 > 0 && me.W2 > 0) &&
				(d0 > 0 && d1 > 0 && d2 > 0) && (

				(class == 0) &&
				(d1 - d0 > -MAX && d1 - d0 <= 0 ) &&
				(d2 - d1 > -MAX && d2 - d1 <= 0 ) ||

				(class == 1) &&
				(d1 - d0 > -MAX && d1 - d0 <=  0) &&
				(d2 - d1 >=   0 && d2 - d1 < MAX) ||

				(class == 2) &&
				(d1 - d0 >=   0 && d1 - d0 < MAX) &&
				(d2 - d1 > -MAX && d2 - d1 <=  0) ||

				(class == 3) &&
				(d1 - d0 >=   0 && d1 - d0 < MAX) &&
				(d2 - d1 >=   0 && d2 - d1 < MAX)) {

				return
			}
		*/
		if (me.V0 > 0 && me.V1 > 0 && me.V2 > 0) &&
			(me.W0 > 0 && me.W1 > 0 && me.W2 > 0) &&
			(d0 > 0 && d1 > 0 && d2 > 0) && ((class == 0) &&
			(d1-d0 > -MAX && d1-d0 <= 0) &&
			(d2-d1 > -MAX && d2-d1 <= 0) ||

			(class == 1) &&
				(d1-d0 > -MAX && d1-d0 <= 0) &&
				(d2-d1 >= 0 && d2-d1 < MAX) ||

			(class == 2) &&
				(d1-d0 >= 0 && d1-d0 < MAX) &&
				(d2-d1 > -MAX && d2-d1 <= 0) ||

			(class == 3) &&
				(d1-d0 >= 0 && d1-d0 < MAX) &&
				(d2-d1 >= 0 && d2-d1 < MAX)) {

			return
		}
	}
}

// GetChain0 returns the chain of \upsilon_i agents. (not certified)
func (me *Agent) GetChain0(x rune, chain *Supplier) *Supplier {
	chain.Push(me.ID)
	var oldChain, newChain *Supplier

	if me.Has(x) > 0 {
		return chain
	}

	var min = MAX
	var ups = 0
	var i = 0
	for _, i = range me.Neighborhood() {
		if i == -1 {
			continue
		}
		/*
			if chain.IndexOf(i) != -1 {
				continue
			}
		*/
		if !(true &&
			me.grid[chain.LastNd()].target.X-me.target.X == me.target.X-me.grid[i].target.X &&
			me.grid[chain.LastNd()].target.Y-me.target.Y == me.target.Y-me.grid[i].target.Y &&
			me.grid[chain.LastNd()].target.Z-me.target.Z == me.target.Z-me.grid[i].target.Z) {
			continue
		}

		// fmt.Printf("%d\n", i)
		newChain = me.grid[i].GetChain0(x, chain)
		if newChain == nil {
			ups = MAX
		} else {
			ups = me.grid[newChain.Last()].V0
		}

		if ups < min {
			oldChain = newChain
			min = ups
		}
	}
	return oldChain
}

// GetChain1 returns the chain of \upsilon'_i agents. (not certified)
func (me *Agent) GetChain1(x rune, chain *Supplier) *Supplier {
	chain.Push(me.ID)
	var oldChain, newChain *Supplier

	if me.Has(x) > 0 {
		return chain
	}

	var min = MAX
	var ups = 0
	var i = 0
	for _, i = range me.Neighborhood() {
		if i == -1 {
			continue
		}
		/*
			if chain.IndexOf(i) != -1 {
				continue
			}
		*/
		if !(true &&
			me.grid[chain.LastNd()].target.X-me.target.X == me.target.X-me.grid[i].target.X &&
			me.grid[chain.LastNd()].target.Y-me.target.Y == me.target.Y-me.grid[i].target.Y &&
			me.grid[chain.LastNd()].target.Z-me.target.Z == me.target.Z-me.grid[i].target.Z) {
			continue
		}

		newChain = me.grid[i].GetChain1(x, chain)
		if newChain == nil {
			ups = MAX
		} else {
			ups = me.grid[newChain.Last()].V1
		}

		if ups < min {
			oldChain = newChain
			min = ups
		}
	}
	return oldChain
}

// GetChain2 returns the chain of \upsilon''_i agents. (not certified)
func (me *Agent) GetChain2(x rune, chain *Supplier) *Supplier {
	chain.Push(me.ID)
	var oldChain, newChain *Supplier

	if me.Has(x) > 0 {
		return chain
	}

	var min = MAX
	var ups = 0
	var i = 0
	for _, i = range me.Neighborhood() {
		if i == -1 {
			continue
		}
		/*
			if chain.IndexOf(i) != -1 {
				continue
			}
		*/
		if !(true &&
			me.grid[chain.LastNd()].target.X-me.target.X == me.target.X-me.grid[i].target.X &&
			me.grid[chain.LastNd()].target.Y-me.target.Y == me.target.Y-me.grid[i].target.Y &&
			me.grid[chain.LastNd()].target.Z-me.target.Z == me.target.Z-me.grid[i].target.Z) {
			continue
		}

		newChain = me.grid[i].GetChain2(x, chain)
		if newChain == nil {
			ups = MAX
		} else {
			ups = me.grid[newChain.Last()].V2
		}

		if ups < min {
			oldChain = newChain
			min = ups
		}
	}
	return oldChain
}

// GetDelta0 returns the exchange value when the agent doesn't use the system
func (me *Agent) GetDelta0() int {
	return me.V0 - me.W0
}

// GetDelta1 returns the exchange value when the agent uses the system but doesn't act circularly
func (me *Agent) GetDelta1() int {
	return me.V1 - me.W1
}

// GetDelta2 returns the exchange value when the agent uses the system and acts circularly
func (me *Agent) GetDelta2() int {
	return me.V2 - me.W2
}

// GetReward returns the reward, according to system 5.3
func (me *Agent) GetReward() int {
	if !REWARDENABLED {
		return 0
	}

	var d0 = me.GetDelta0()
	var d1 = me.GetDelta1()
	var d2 = me.GetDelta2()

	if d0 >= d2 || d1 >= d2 {
		return 0
	}

	var diff = d2 - d1
	if diff > 1 {
		// return (diff - diff%2) / 2
		return diff - 1
	}

	return 0
}

// SetSource initializes the source of an agent.
func (me *Agent) SetSource(xyz int) {
	var xy = xyz % 64
	me.source.Z = (xyz - xy) / 64
	me.source.X = xy % 8
	me.source.Y = (xy - me.source.X) / 8
	/*
		me.source.Z *= 32
		me.source.X *= 32
		me.source.Y *= 32
	*/
}

// SetTarget initializes the target of an agent.
func (me *Agent) SetTarget(xyz int) {
	me.ID = xyz

	var xy = xyz % 64
	me.target.Z = (xyz - xy) / 64
	me.target.X = xy % 8
	me.target.Y = (xy - me.target.X) / 8
	/*
		me.target.Z *= 32
		me.target.X *= 32
		me.target.Y *= 32
	*/
}

// Prepare computes the expected payout.
func (me *Agent) Prepare() {
	// memory to store the units U to sell
	var u2 = NewColor()
	var u1 = NewColor()
	var u0 = NewColor()

	var he *Agent
	var d0, d1, d2, i, w = 0, 0, 0, 0, 0
	/*
		var v2 = (me.V2 <= me.V1) && (me.V2 <= me.V0)
		var v1 = (me.V1 <= me.V2) && (me.V1 <= me.V0)
		var v0 = (me.V0 <= me.V1) && (me.V0 <= me.V2)
	*/
	for i = 0; i < len(me.grid); i++ {
		he = me.grid[i]
		if me.ID == he.ID {
			continue
		}

		d0 = he.GetDelta0()
		d1 = he.GetDelta1()
		d2 = he.GetDelta2()

		w = -1
		// he would pay me w
		if (d2 >= 0) && (d2 >= d1) && (d2 >= d0) {
			w = he.W2
		}

		if (d1 >= 0) && (d1 >= d2) && (d1 >= d0) {
			w = he.W1
		}

		if (d0 >= 0) && (d0 >= d1) && (d0 >= d2) {
			w = he.W0
		}

		if /*v2 &&*/ w >= me.V2 {
			if me.Has(X) > 0 && he.Has(X) < 0 {
				u2.X++
			}

			if me.Has(Y) > 0 && he.Has(Y) < 0 {
				u2.Y++
			}

			if me.Has(Z) > 0 && he.Has(Z) < 0 {
				u2.Z++
			}
		}

		if /*v1 &&*/ w >= me.V1 {
			if me.Has(X) > 0 && he.Has(X) < 0 {
				u1.X++
			}

			if me.Has(Y) > 0 && he.Has(Y) < 0 {
				u1.Y++
			}

			if me.Has(Z) > 0 && he.Has(Z) < 0 {
				u1.Z++
			}
		}

		if /*v0 &&*/ w >= me.V0 {
			if me.Has(X) > 0 && he.Has(X) < 0 {
				u0.X++
			}

			if me.Has(Y) > 0 && he.Has(Y) < 0 {
				u0.Y++
			}

			if me.Has(Z) > 0 && he.Has(Z) < 0 {
				u0.Z++
			}
		}
	}

	// compute the expected payout
	me.R2 = (u2.X + u2.Y + u2.Z) // * (me.V2 - me.W2)
	me.R1 = (u1.X + u1.Y + u1.Z) // * (me.V1 - me.W1)
	me.R0 = (u0.X + u0.Y + u0.Z) // * (me.V0 - me.W0)
}

// Has returns the supply of x that the agent has.
func (me *Agent) Has(x rune) int {
	switch x {
	case X:
		return me.source.X - me.target.X
	case Y:
		return me.source.Y - me.target.Y
	default:
		return me.source.Z - me.target.Z
	}
}

// (not certified)
func (me *Agent) distance(he *Agent) int {
	var x = math.Abs(float64(me.target.X - he.target.X))
	var y = math.Abs(float64(me.target.Y - he.target.Y))
	var z = math.Abs(float64(me.target.Z - he.target.Z))
	return int(x+y+z) * 16
}

// Neighborhood returns the neighborhood indexes randomly sorted. (not certified)
func (me *Agent) Neighborhood() []int {
	// initialize local pseudorandom generator
	rand.Seed(time.Now().UnixNano())

	var seq = make([]int, 27)
	var col = NewColor()
	var i = 0
	var x, y, z = 0, 0, 0
	for i = range seq {
		x = i%3 - 1
		y = ((i-i%3)/3)%3 - 1
		z = ((((i-i%3)/3)-((i-i%3)/3)%3)/3)%3 - 1

		col.X = me.target.X + x // 26%2 - 1
		col.Y = me.target.Y + y // ((26 - 26%3)/3)%3 - 1
		col.Z = me.target.Z + z // ((((26 - 26%3)/3) - ((26 - 26%3)/3)%3)/3)%3 - 1

		if false ||
			col.X < 0 || col.Y < 0 || col.Z < 0 ||
			col.X > 7 || col.Y > 7 || col.Z > 7 ||

			x == 0 && y == 0 && z == 0 ||

			x != 0 && y != 0 || x != 0 && z != 0 || y != 0 && z != 0 {
			seq[i] = -1
		} else {
			seq[i] = col.X*64 + col.Y*8 + col.Z
		}
	}

	// sort the 27 neighbors by time of transaction
	rand.Shuffle(len(seq), func(i, j int) {
		seq[i], seq[j] = seq[j], seq[i]
	})

	return seq
}

// Like returns the best way to buy to he. In case of draw, Like prefers d00 over the another ways (partially certified)
func (me *Agent) like(he *Agent) *Supplier { // x rune
	/*
		var chain0 = he.GetChain0(x, NewChain(me.ID))
		var chain1 = he.GetChain1(x, NewChain(me.ID))
		var chain2 = he.GetChain2(x, NewChain(me.ID))

		var d0, d1, d2 = -MAX, -MAX, -MAX
		if chain0 != nil {
			d0 = me.V0 - me.grid[chain0.Last()].V0
		}
		if chain1 != nil {
			d1 = me.V1 - me.grid[chain1.Last()].V1
		}
		if chain2 != nil {
			d2 = me.V2 - me.grid[chain2.Last()].V2
			d2 += me.grid[chain2.Last()].GetReward()
		}
	*/
	var chain = NewChain(me.ID)
	chain.Push(he.ID)

	var d00 = me.V0 - he.V0
	var d11 = me.V1 - he.V1
	var d12 = me.V1 - he.V2 + he.GetReward()
	var d22 = me.V2 - he.V2 + he.GetReward()
	var d21 = me.V2 - he.V1

	// (I || B) && (J || G) && (X5 || G) && (X6 || B) && Q
	if true &&
		(d00 >= d11 || he.GetDelta1() <= 0) &&
		(d00 >= d22 || he.GetDelta2() <= 0) &&
		(d00 >= d12 || he.GetDelta2() <= 0) &&
		(d00 >= d21 || he.GetDelta1() <= 0) {
		if he.GetDelta0() > 0 {
			chain.value = d00
			chain.axis = 0
			return chain
		}
	}
	// (F || G) && (H || D) && (X3 || G) && (X4 || B) && P
	if true &&
		(d11 >= d22 || he.GetDelta2() <= 0) &&
		(d11 >= d00 || he.GetDelta0() <= 0) &&
		(d11 >= d12 || he.GetDelta2() <= 0) &&
		(d11 >= d21 || he.GetDelta1() <= 0) {
		if he.GetDelta1() > 0 {
			chain.value = d11
			chain.axis = 11
			return chain
		}
	}
	// (K || B) && (L || G) && (X7 || D) && (X8 || B) && O
	if true &&
		(d12 >= d11 || he.GetDelta1() <= 0) &&
		(d12 >= d22 || he.GetDelta2() <= 0) &&
		(d12 >= d00 || he.GetDelta0() <= 0) &&
		(d12 >= d21 || he.GetDelta1() <= 0) {
		if he.GetDelta2() > 0 {
			chain.value = d12
			chain.axis = 12
			return chain
		}
	}
	// (M || B) && (N || G) && (X9 || G) && (X0 || D) && P
	if true &&
		(d21 >= d11 || he.GetDelta1() <= 0) &&
		(d21 >= d22 || he.GetDelta2() <= 0) &&
		(d21 >= d12 || he.GetDelta2() <= 0) &&
		(d21 >= d00 || he.GetDelta0() <= 0) {
		if he.GetDelta1() > 0 {
			chain.value = d21
			chain.axis = 21
			return chain
		}
	}
	// (A || B) && (C || D) && (X1 || G) && (X2 || B) && O
	if true &&
		(d22 >= d11 || he.GetDelta1() <= 0) &&
		(d22 >= d00 || he.GetDelta0() <= 0) &&
		(d22 >= d12 || he.GetDelta2() <= 0) &&
		(d22 >= d21 || he.GetDelta1() <= 0) {
		if he.GetDelta2() > 0 {
			chain.value = d22
			chain.axis = 22
			return chain
		}
	}

	return nil
}

// Search looks for another agent to exchange. (not certified)
func (me *Agent) Search() {
	var he *Agent
	var diff = -MAX
	var like int
	var chain *Supplier
	var j int
	for _, j = range me.grid.Shuffled() {
		he = me.grid[j]

		chain = me.like(he) // X, Y, Z?
		if like > diff {
			diff = like

			if me.Has(X) < 0 && he.Has(X) > 0 {
				me.xSupplier = chain
			}

			if me.Has(Y) < 0 && he.Has(Y) > 0 {
				me.ySupplier = chain
			}

			if me.Has(Z) < 0 && he.Has(Z) > 0 {
				me.zSupplier = chain
			}
		}
	}
}

// Look looks the agent j to exchange.
func (me *Agent) Look(j int, reset bool) {
	var he = me.grid[j]
	var chain = me.like(he)
	/*
		var xChain, xHerVal = me.like(he, X)
		var yChain, yHerVal = me.like(he, Y)
		var zChain, zHerVal = me.like(he, Z)
	*/
	if reset {
		me.xSupplier, me.ySupplier, me.zSupplier = nil, nil, nil
	}

	if chain == nil {
		return
	}

	if me.Has(X) < 0 && he.Has(X) > 0 {
		if me.xSupplier != nil {
			if chain.value > me.xSupplier.value {
				me.xSupplier = chain
			}
		} else {
			me.xSupplier = chain
		}
	}

	if me.Has(Y) < 0 && he.Has(Y) > 0 {
		if me.ySupplier != nil {
			if chain.value > me.ySupplier.value {
				me.ySupplier = chain
			}
		} else {
			me.ySupplier = chain
		}
	}

	if me.Has(Z) < 0 && he.Has(Z) > 0 {
		if me.zSupplier != nil {
			if chain.value > me.zSupplier.value {
				me.zSupplier = chain
			}
		} else {
			me.zSupplier = chain
		}
	}
}

// Request starts a transaction for x.
func (me *Agent) Request(x rune) bool {
	if false ||
		x == X && me.xSupplier == nil ||
		x == Y && me.ySupplier == nil ||
		x == Z && me.zSupplier == nil {
		return false
	}

	var he *Agent
	var ok, v = false, 0
	switch x {
	case X:
		he = me.grid[me.xSupplier.Last()]
		/*
			if he.Has(X) <= 0 {
				return false
			}
		*/
		if ok, v = he.Response(x, me.xSupplier.axis); ok {
			// me.source.X++
			me.history.logreq(x, me.xSupplier.axis, v)
			me.xSupplier = nil
			return true
		}
	case Y:
		he = me.grid[me.ySupplier.Last()]
		if ok, v = he.Response(x, me.ySupplier.axis); ok {
			me.history.logreq(x, me.ySupplier.axis, v)
			me.ySupplier = nil
			return true
		}
	case Z:
		he = me.grid[me.zSupplier.Last()]
		if ok, v = he.Response(x, me.zSupplier.axis); ok {
			me.history.logreq(x, me.zSupplier.axis, v)
			me.zSupplier = nil
			return true
		}
	}

	// log.Fatal("Error on transfer!")
	return false
}

// Response ends a transaction for x. (partially certified)
func (me *Agent) Response(x rune, axis int) (bool, int) {
	var d, v = 0, 0
	switch axis {
	case 12:
		d = me.GetDelta2() - me.GetReward()
		v = me.V2
	case 22:
		d = me.GetDelta2() - me.GetReward()
		v = me.V2
	case 11:
		d = me.GetDelta1()
		v = me.V1
	case 21:
		d = me.GetDelta1()
		v = me.V1
	case 0:
		d = me.GetDelta0()
		v = me.V0
	}

	if d > 0 {
		// me.source.X--
		me.history.logres(x, axis, d)
		return true, v
	}

	return false, v
}

// Learn adjusts agent values in order to increase revenue.
func (me *Agent) Learn() {
	me.history = me.history.archive()
	var last = me.history[len(me.history)-2]
	var prev = me.history[0]

	// be careful with the historical ws!
	if len(me.history)-3 >= 0 {
		prev = me.history[len(me.history)-3]
		// use the last w2
		if last.w2 == 0 {
			last.w2 = prev.w2
		}
		// use the last w1
		if last.w1 == 0 {
			last.w1 = prev.w1
		}
		// use the last w0
		if last.w0 == 0 {
			last.w0 = prev.w0
		}
	}

	// me.Prepare()

	// adjust V2 and W2
	if 0 < last.w2 && last.w2 <= me.W2 {
		me.W2 = last.w2
	}

	if last.c2 < me.R2 && me.V2-1 > me.W2 {
		me.V2 = me.V2 - 1
	}

	// adjust V1 and W1
	if 0 < last.w1 && last.w1 <= me.W1 {
		me.W1 = last.w1
	}

	if last.c1 < me.R1 && me.V1-1 > me.W1 {
		me.V1 = me.V1 - 1
	}

	// adjust V0 and W0
	if 0 < last.w0 && last.w0 <= me.W0 {
		me.W0 = last.w0
	}

	if last.c0 < me.R0 && me.V0-1 > me.W0 {
		me.V0 = me.V0 - 1
	}
}
