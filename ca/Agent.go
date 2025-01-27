package ca

import (
	"math"
	"math/rand"
	"time"
)

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

// RSupplier returns the supplier for X
func (me *Agent) RSupplier() int {
	if me.rSupplier != nil {
		return me.rSupplier.Last()
	}

	return -1
}

// GSupplier returns the supplier for Y
func (me *Agent) GSupplier() int {
	if me.gSupplier != nil {
		return me.gSupplier.Last()
	}

	return -1
}

// BSupplier returns the supplier for Z
func (me *Agent) BSupplier() int {
	if me.bSupplier != nil {
		return me.bSupplier.Last()
	}

	return -1
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
			me.grid[chain.LastNd()].target.R-me.target.R == me.target.R-me.grid[i].target.R &&
			me.grid[chain.LastNd()].target.G-me.target.G == me.target.G-me.grid[i].target.G &&
			me.grid[chain.LastNd()].target.B-me.target.B == me.target.B-me.grid[i].target.B) {
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
			me.grid[chain.LastNd()].target.R-me.target.R == me.target.R-me.grid[i].target.R &&
			me.grid[chain.LastNd()].target.G-me.target.G == me.target.G-me.grid[i].target.G &&
			me.grid[chain.LastNd()].target.B-me.target.B == me.target.B-me.grid[i].target.B) {
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
			me.grid[chain.LastNd()].target.R-me.target.R == me.target.R-me.grid[i].target.R &&
			me.grid[chain.LastNd()].target.G-me.target.G == me.target.G-me.grid[i].target.G &&
			me.grid[chain.LastNd()].target.B-me.target.B == me.target.B-me.grid[i].target.B) {
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
			if me.Has(R) > 0 && he.Has(R) < 0 {
				u2.R++
			}

			if me.Has(G) > 0 && he.Has(G) < 0 {
				u2.G++
			}

			if me.Has(B) > 0 && he.Has(B) < 0 {
				u2.B++
			}
		}

		if /*v1 &&*/ w >= me.V1 {
			if me.Has(R) > 0 && he.Has(R) < 0 {
				u1.R++
			}

			if me.Has(G) > 0 && he.Has(G) < 0 {
				u1.G++
			}

			if me.Has(B) > 0 && he.Has(B) < 0 {
				u1.B++
			}
		}

		if /*v0 &&*/ w >= me.V0 {
			if me.Has(R) > 0 && he.Has(R) < 0 {
				u0.R++
			}

			if me.Has(G) > 0 && he.Has(G) < 0 {
				u0.G++
			}

			if me.Has(B) > 0 && he.Has(B) < 0 {
				u0.B++
			}
		}
	}

	// compute the expected payout
	me.R2 = (u2.R + u2.G + u2.B) // * (me.V2 - me.W2)
	me.R1 = (u1.R + u1.G + u1.B) // * (me.V1 - me.W1)
	me.R0 = (u0.R + u0.G + u0.B) // * (me.V0 - me.W0)
}

// (not certified)
func (me *Agent) distance(he *Agent) int {
	var x = math.Abs(float64(me.target.R - he.target.R))
	var y = math.Abs(float64(me.target.G - he.target.G))
	var z = math.Abs(float64(me.target.B - he.target.B))
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

		col.R = me.target.R + x // 26%2 - 1
		col.G = me.target.G + y // ((26 - 26%3)/3)%3 - 1
		col.B = me.target.B + z // ((((26 - 26%3)/3) - ((26 - 26%3)/3)%3)/3)%3 - 1

		if false ||
			col.R < 0 || col.G < 0 || col.B < 0 ||
			col.R > 7 || col.G > 7 || col.B > 7 ||

			x == 0 && y == 0 && z == 0 ||

			x != 0 && y != 0 || x != 0 && z != 0 || y != 0 && z != 0 {
			seq[i] = -1
		} else {
			seq[i] = col.R*64 + col.G*8 + col.B
		}
	}

	// sort the 27 neighbors by time of transaction
	rand.Shuffle(len(seq), func(i, j int) {
		seq[i], seq[j] = seq[j], seq[i]
	})

	return seq
}

// Like returns the best way to buy to he. In case of draw, Like prefers d00 over the another ways (partially certified)
func (me *Agent) free(he *Agent) (int, int) { // x rune
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
			return chain.value, chain.axis
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
			return chain.value, chain.axis
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
			return chain.value, chain.axis
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
			return chain.value, chain.axis
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
			return chain.value, chain.axis
		}
	}

	return 0, -1
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
