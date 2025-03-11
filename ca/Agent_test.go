package ca

import (
	"testing"
)

func TestSetSource(t *testing.T) {
	var agent = NewAgent(0)
	var dataset = [][]int{
		{0, 0, 0, 0},
		{7, 7, 0, 0},
		{8, 0, 1, 0},
		{9, 1, 1, 0},
		{63, 7, 7, 0},
		{64, 0, 0, 1},
		{65, 1, 0, 1},
		{71, 7, 0, 1},
		{72, 0, 1, 1},
		{73, 1, 1, 1},
		{511, 7, 7, 7},
	}

	var i = 0
	for i = range dataset {
		agent.SetSource(dataset[i][0])
		if agent.source.X != dataset[i][1] || agent.source.Y != dataset[i][2] || agent.source.Z != dataset[i][3] {
			t.Errorf("Error with source %d: X=%d; Y=%d; Z=%d\n", dataset[i][0], agent.source.X, agent.source.Y, agent.source.Z)
			return
		}
	}
}

func TestSetTarget(t *testing.T) {
	var agent = NewAgent(0)
	var dataset = [][]int{
		{0, 0, 0, 0},
		{7, 7, 0, 0},
		{8, 0, 1, 0},
		{9, 1, 1, 0},
		{63, 7, 7, 0},
		{64, 0, 0, 1},
		{65, 1, 0, 1},
		{71, 7, 0, 1},
		{72, 0, 1, 1},
		{73, 1, 1, 1},
		{511, 7, 7, 7},
	}

	var i = 0
	for i = range dataset {
		agent.SetTarget(dataset[i][0])
		if agent.target.X != dataset[i][1] || agent.target.Y != dataset[i][2] || agent.target.Z != dataset[i][3] {
			t.Errorf("Error with target %d: X=%d; Y=%d; Z=%d\n", dataset[i][0], agent.target.X, agent.target.Y, agent.target.Z)
			return
		}
	}
}

func TestGetReward(t *testing.T) {
	var agent = NewAgent(0)
	agent.V0 = 0
	agent.W0 = 0
	agent.V1 = 128
	agent.W1 = 64
	agent.V2 = 256

	agent.W2 = 127
	if agent.GetReward() != 64 { // 32
		t.Errorf("Error computing the reward with W2 = %d", agent.W2)
		return
	}

	agent.W2 = 128
	if agent.GetReward() != 63 { // 32
		t.Errorf("Error computing the reward with W2 = %d", agent.W2)
		return
	}
}

func TestPrepare(t *testing.T) {
	var n = 512
	var prod *Agent
	var cons = NewAgent(0)
	var grid = NewGrid(n)
	var dataset = [][]int{
		{6, 2, 256, 128 /**/, 8, 4, 256, 128 /**/, 8, 4, 256, 128},
		{6, 2, 128, 256 /**/, 8, 4, 128, 256 /**/, 8, 4, 128, 256},
	}

	cons.V0, cons.W0 = dataset[0][0], dataset[0][1]
	cons.V1, cons.W1 = dataset[0][4], dataset[0][5]
	cons.V2, cons.W2 = dataset[0][8], dataset[0][9]

	cons.SetSource(1)
	grid.Add(cons, 0)

	// initialize grid
	var i int
	for i = 1; i < n; i++ {
		prod = NewAgent(0)

		if i%50 == 0 {
			prod.V0, prod.W0 = dataset[0][2], dataset[0][3]
			prod.V1, prod.W1 = dataset[0][6], dataset[0][7]
			prod.V2, prod.W2 = dataset[0][10], dataset[0][11]
		} else {
			prod.V0, prod.W0 = dataset[1][2], dataset[1][3]
			prod.V1, prod.W1 = dataset[1][6], dataset[1][7]
			prod.V2, prod.W2 = dataset[1][10], dataset[1][11]
		}

		if i%100 == 0 {
			prod.SetTarget(0)
		} else {
			prod.SetTarget(1)
		}

		grid.Add(prod, i)
	}

	cons.Prepare()

	if cons.R0 != 20 && cons.R1 != 20 || cons.R2 != 20 {
		t.Errorf("Error: bad expected payout (R0 is %d, R1 is %d and R2 is %d)", cons.R0, cons.R1, cons.R2)
		return
	}

	testLearn(grid, t)
}

// Warning! cons prefers V2 if V2 = V1 = V0?
func testLearn(grid Grid, t *testing.T) {
	var m = 1
	var n = 512
	var h = 0
	var i = 0
	var j = 0
	var k = 2

	// test the producer with the expected consumers
	for i = 0; i < k; i++ {
		for h = 0; h < m; h++ {
			for j = 0; j < n; j++ {
				grid[50].Look(j, false)
				grid[150].Look(j, false)
				grid[250].Look(j, false)
				grid[350].Look(j, false)
				grid[450].Look(j, false)
			}
		}

		grid[50].Request(X)
		grid[150].Request(X)
		grid[250].Request(X)
		grid[350].Request(X)
		grid[450].Request(X)

		grid[0].Learn()
		grid[50].Learn()
	}

	if false ||
		grid[0].W0 != 2 || grid[0].W1 != 4 || grid[0].W2 != 4 ||
		grid[0].V0 != 6 || grid[0].V1 != 8 || grid[0].V2 != 8 {
		t.Errorf(
			"Error on the producer with the expected consumers, the records are %d, %d, %d, %d, %d, %d",
			grid[0].V0,
			grid[0].W0,
			grid[0].V1,
			grid[0].W1,
			grid[0].V2,
			grid[0].W2,
		)
		return
	}

	// test one consumer
	if false ||
		grid[50].W0 != 128 || grid[50].W1 != 128 || grid[50].W2 != 128 ||
		grid[50].V0 != 256 || grid[50].V1 != 256 || grid[50].V2 != 256 {
		t.Errorf(
			"Error on the consumer, the records are %d, %d, %d, %d, %d, %d",
			grid[50].V0,
			grid[50].W0,
			grid[50].V1,
			grid[50].W1,
			grid[50].V2,
			grid[50].W2,
		)
		return
	}

	// test the producer with less consumers than expected
	m = 2
	for i = 0; i < k; i++ {
		for h = 0; h < m; h++ {
			for j = 0; j < n; j++ {
				grid[50].Look(j, false)
			}

			grid[50].Request(X)
		}

		grid[0].Learn()
	}

	if false ||
		grid[0].W0 != 2 || grid[0].W1 != 4 || grid[0].W2 != 4 ||
		grid[0].V0 != 3 || grid[0].V1 != 8 || grid[0].V2 != 8 {
		t.Errorf(
			"Error on the producer with less consumers than expected, the records are. %d, %d, %d, %d, %d, %d",
			grid[0].V0,
			grid[0].W0,
			grid[0].V1,
			grid[0].W1,
			grid[0].V2,
			grid[0].W2,
		)
		return
	}

	// test the producer with more consumers than expected
	m = 6
	for i = 0; i < k; i++ {
		for h = 0; h < m; h++ {
			for j = 0; j < n; j++ {
				grid[50].Look(j, false)
				grid[150].Look(j, false)
				grid[250].Look(j, false)
				grid[350].Look(j, false)
				grid[450].Look(j, false)
			}

			grid[50].Request(X)
			grid[150].Request(X)
			grid[250].Request(X)
			grid[350].Request(X)
			grid[450].Request(X)
		}

		grid[0].Learn()
	}

	if false ||
		grid[0].W0 != 2 || grid[0].W1 != 4 || grid[0].W2 != 4 ||
		grid[0].V0 != 23 || grid[0].V1 != 8 || grid[0].V2 != 8 {
		t.Errorf(
			"Error on the producer with more consumers than expected, the records are %d, %d, %d, %d, %d, %d",
			grid[0].V0,
			grid[0].W0,
			grid[0].V1,
			grid[0].W1,
			grid[0].V2,
			grid[0].W2,
		)
		return
	}
}

func TestLike(t *testing.T) {
	var cons = NewAgent(0)
	var prod = NewAgent(0)
	var supp *Supplier
	//  A,  B,  C,  D,  F,  G,  H,  I,  J,  K,  L,  M,  N,  O,  P,  Q,
	// X1, X2, X3, X4, X5, X6, X7, X8, X9, X0
	var dataset = [][]int{
		{240 /**/, 4, 2, 1 /**/, 32, 16, 15 /**/, 256, 128, 127}, // 21
		{240 /**/, 4, 2, 1 /**/, 256, 128, 127 /**/, 32, 16, 15}, // 12
		{128 /**/, 256, 128, 127 /**/, 4, 2, 1 /**/, 32, 16, 15}, // 00

		{240 /**/, 4, 2, 1 /**/, 32, 16, 15 /**/, 256, 128, 129}, // 22
		{240 /**/, 4, 2, 1 /**/, 256, 128, 129 /**/, 32, 16, 15}, // 11
		{30 /**/, 256, 128, 129 /**/, 4, 2, 1 /**/, 32, 16, 15},  // 22
	}

	var i = 0
	for i = range dataset {
		cons.V0, prod.V0, prod.W0 = dataset[i][1], dataset[i][2], dataset[i][3]
		cons.V1, prod.V1, prod.W1 = dataset[i][4], dataset[i][5], dataset[i][6]
		cons.V2, prod.V2, prod.W2 = dataset[i][7], dataset[i][8], dataset[i][9]

		if prod.GetReward() != 0 {
			t.Errorf("Error %d.0: the reward is %d", i, prod.GetReward())
			return
		}

		supp = cons.like(prod)
		if supp.value != dataset[i][0] {
			t.Errorf("Error %d.1: consumer don't likes producer (val is %d instead of %d)", i, supp.value, dataset[i][0])
			return
		}
	}
}

func TestLike2(t *testing.T) {
	// A - B > C - D || C - E <= 0: With
	//(B - F) - (C - E) = 1, B - F > 0 And
	// R = A - B:
	// 1. E > C Or
	// 2. A > C - D + B Or
	// 3. B < A - C + D Or
	// 4. C < A - B + D Or
	// 5. D > A + B + C Or

	// With B = 65, F = 32, C = 64, E = 32:
	// 2.
	// 3.
	// 4.
	// 5.
}

func TestHas(t *testing.T) {
	var agent = NewAgent(0)
	var dataset = [][]int{
		{1, 2, 1 /**/, 2, 4, 2 /**/, 4, 8, 4},
		{0, 2, 2 /**/, 0, 4, 4 /**/, 0, 8, 8},
		{-1, 1, 2 /**/, -2, 2, 4 /**/, -4, 4, 8},
	}

	var i = 0
	for i = range dataset {
		agent.source.X, agent.target.X = dataset[i][1], dataset[i][2]
		agent.source.Y, agent.target.Y = dataset[i][4], dataset[i][5]
		agent.source.Z, agent.target.Z = dataset[i][7], dataset[i][8]

		if agent.Has(X) != dataset[i][0] || agent.Has(Y) != dataset[i][3] || agent.Has(Z) != dataset[i][6] {
			t.Errorf("Error: the agent has %d, %d, %d", agent.Has(X), agent.Has(Y), agent.Has(Z))
			return
		}
	}
}

func TestLook(t *testing.T) {
	var n = 512
	var grid = NewGrid(n)
	var treat = []int{128, 128, 128, 128}

	// initialize grid
	var i, j, k int
	k = 0
	for j = range treat {
		for i = 0; i < treat[j]; i++ {
			grid.Add(NewAgent(j), k)
			k++
		}
	}

	grid.Shuffle()

	grid[0].source.X, grid[0].source.Y, grid[0].source.Z = 0, 0, 0
	grid[0].target.X, grid[0].target.Y, grid[0].target.Z = 7, 7, 7

	grid[0].V0, grid[0].V1, grid[0].V2 = MAX*2, MAX*2, MAX*2
	grid[0].W0, grid[0].W1, grid[0].W2 = MAX*1, MAX*1, MAX*1

	grid[64].source.X, grid[128].source.X, grid[256].source.X = 7, 0, 0
	grid[64].target.X, grid[128].target.X, grid[256].target.X = 0, 0, 0
	grid[64].source.Y, grid[128].source.Y, grid[256].source.Y = 0, 7, 0
	grid[64].target.Y, grid[128].target.Y, grid[256].target.Y = 0, 0, 0
	grid[64].source.Z, grid[128].source.Z, grid[256].source.Z = 0, 0, 7
	grid[64].target.Z, grid[128].target.Z, grid[256].target.Z = 0, 0, 0

	grid[64].V0, grid[64].V1, grid[64].V2 = 32, 8, 16
	grid[64].W0, grid[64].W1, grid[64].W2 = 30, 4, 14

	grid[128].V0, grid[128].V1, grid[128].V2 = 8, 16, 4
	grid[128].W0, grid[128].W1, grid[128].W2 = 6, 14, 0

	grid[256].V0, grid[256].V1, grid[256].V2 = 2, 4, 8
	grid[256].W0, grid[256].W1, grid[256].W2 = 0, 3, 7

	for j = 0; j <= n/2; j++ {
		grid[0].Look(j, false)
	}

	grid[0].Look(n/2+1, true)

	for j = n/2 + 2; j < n; j++ {
		grid[0].Look(j, false)
	}

	if false ||
		grid[0].xSupplier != nil && grid[0].xSupplier.chain[1] == 64 ||
		grid[0].ySupplier != nil && grid[0].ySupplier.chain[1] == 128 ||
		grid[0].zSupplier != nil && grid[0].zSupplier.chain[1] == 256 {
		t.Errorf(
			"Error: search not restarted, the suppliers are %d, %d, %d",
			grid[0].xSupplier.chain[1],
			grid[0].ySupplier.chain[1],
			grid[0].zSupplier.chain[1],
		)
		return
	}

	for j = 0; j < n; j++ {
		grid[0].Look(j, false)
	}

	if grid[0].xSupplier == nil {
		t.Errorf("Error: the supplier for X must be 64")
		return
	}

	if grid[0].ySupplier == nil {
		t.Errorf("Error: the supplier for Y must be 128")
		return
	}

	if grid[0].zSupplier == nil {
		t.Errorf("Error: the supplier for Z must be 256")
		return
	}

	if grid[0].xSupplier.chain[1] != 64 || grid[0].ySupplier.chain[1] != 128 || grid[0].zSupplier.chain[1] != 256 {
		t.Errorf(
			"Warning: the suppliers are %d, %d, %d, please try again",
			grid[0].xSupplier.chain[1],
			grid[0].ySupplier.chain[1],
			grid[0].zSupplier.chain[1],
		)
		return
	}

	testRequest(grid, t)
}

func testRequest(grid Grid, t *testing.T) {
	var he = NewAgent(0)
	if he.Request(X) {
		t.Errorf("Error: request successful")
		return
	}

	if grid[0].Request('W') {
		t.Errorf("Error: request successful")
		return
	}

	if !grid[0].Request(X) || !grid[0].Request(Y) || !grid[0].Request(Z) {
		t.Errorf("Error: request failed")
		return
	}

	testResponse(grid, t)
}

func testResponse(grid Grid, t *testing.T) {
	var rec64 = grid[64].history[len(grid[64].history)-1]
	var rec128 = grid[128].history[len(grid[128].history)-1]
	var rec256 = grid[256].history[len(grid[256].history)-1]

	if false ||
		rec64.v0 != 0 || rec128.v0 != 0 || rec256.v0 != 2 ||
		rec64.v1 != 4 || rec128.v1 != 0 || rec256.v1 != 0 ||
		rec64.v2 != 0 || rec128.v2 != 3 || rec256.v2 != 0 ||
		rec64.c0 != 0 || rec128.c0 != 0 || rec256.c0 != 1 ||
		rec64.c1 != 1 || rec128.c1 != 0 || rec256.c1 != 0 ||
		rec64.c2 != 0 || rec128.c2 != 1 || rec256.c2 != 0 {
		t.Errorf("Error: the records are %d, %d, %d", rec64.v1, rec128.v2, rec256.v0)
		return
	}

	testLogres(grid, t)
}

func testLogres(grid Grid, t *testing.T) {
	var n = 512
	var j = 0
	for j = 0; j < n; j++ {
		grid[0].Look(j, false)
	}

	grid[0].Request(X)
	grid[0].Request(Y)
	grid[0].Request(Z)

	var rec64 = grid[64].history[len(grid[64].history)-1]
	var rec128 = grid[128].history[len(grid[128].history)-1]
	var rec256 = grid[256].history[len(grid[256].history)-1]

	if rec64.v1 != 8 || rec128.v2 != 6 || rec256.v0 != 4 {
		t.Errorf("Error: the records are %d, %d, %d", rec64.v1, rec128.v2, rec256.v0)
		return
	}

	testLogreq(grid, t)
	testArchive(grid, t)
}

func testLogreq(grid Grid, t *testing.T) {
	var rec = grid[0].history[len(grid[0].history)-1]
	if rec.w1 != 8 || rec.w2 != 0 || rec.w0 != 2 { // rec.w2 != 4
		t.Errorf("Error: the records are %d, %d, %d", rec.w1, rec.w2, rec.w0)
		return
	}
}

func testArchive(grid Grid, t *testing.T) {
	grid[64].history = grid[64].history.archive()
	grid[128].history = grid[128].history.archive()
	grid[256].history = grid[256].history.archive()

	var n = 512
	var j = 0
	for j = 0; j < n; j++ {
		grid[0].Look(j, false)
	}

	grid[0].Request(X)
	grid[0].Request(Y)
	grid[0].Request(Z)

	var his64 = grid[64].history
	var his128 = grid[128].history
	var his256 = grid[256].history

	if false ||
		his64[len(his64)-2].v1 != 8 || his128[len(his128)-2].v2 != 6 || his256[len(his256)-2].v0 != 4 ||
		his64[len(his64)-1].v1 != 4 || his128[len(his128)-1].v2 != 3 || his256[len(his256)-1].v0 != 2 {
		t.Errorf(
			"Error: the records are %d, %d, %d, %d, %d, %d",
			his64[len(his64)-2].v1,
			his128[len(his128)-2].v2,
			his256[len(his256)-2].v0,
			his64[len(his64)-1].v1,
			his128[len(his128)-1].v2,
			his256[len(his256)-1].v0,
		)
		return
	}
}
