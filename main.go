package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"lab3/ca"

	_ "github.com/go-sql-driver/mysql"
)

// N is the number of agents in the system
const N = 512

// T is the number of iterations
const T = 256

// OBS is the number of observations per treatment
const OBS = 16

// MAXPSC is the max number of records per statement
const MAXPSC = 4096

var dbo *sql.DB

// var stm *sql.Stmt
var url = "root:@/sim"
var psc = 0
var dbotmp = ""
var serie = "1"
var stmstr1 = `insert into sim.serie_%s (
_trid, _obid,
_t, _i,

_V0, _W0, _V1,
_W1, _V2, _W2,

_x, _y, _z,

_R0, _R1, _R2,

_xSup, _ySup, _zSup,

_lastV0, _lastV1, _lastV2,
_lastW0, _lastW1, _lastW2,
_lastC0, _lastC1, _lastC2
) values `

var stmstr2 = `(
	%d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, 
	%d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d
)`

func findTreatments() {
	var lev = []int{8, 24, 96, 384}
	var h, i, j, k int
	for h = 0; h < len(lev); h++ {
		for i = 0; i < len(lev); i++ {
			for j = 0; j < len(lev); j++ {
				for k = 0; k < len(lev); k++ {
					if true &&
						lev[h] != lev[i] &&
						lev[h] != lev[j] &&
						lev[h] != lev[k] &&
						lev[i] != lev[j] &&
						lev[i] != lev[k] &&
						lev[j] != lev[k] {
						fmt.Printf(
							"{%d, %d, %d, %d},\n",
							lev[h], lev[i], lev[j], lev[k],
						)
					}
				}
			}
		}
	}
}

func getTreatments() [][]int {
	return [][]int{
		{8, 24, 96, 384},
		{8, 24, 384, 96},
		{8, 96, 24, 384},
		{8, 96, 384, 24},
		{8, 384, 24, 96},
		{8, 384, 96, 24},
		{24, 8, 96, 384},
		{24, 8, 384, 96},
		{24, 96, 8, 384},
		{24, 96, 384, 8},
		{24, 384, 8, 96},
		{24, 384, 96, 8},
		{96, 8, 24, 384},
		{96, 8, 384, 24},
		{96, 24, 8, 384},
		{96, 24, 384, 8},
		{96, 384, 8, 24},
		{96, 384, 24, 8},
		{384, 8, 24, 96},
		{384, 8, 96, 24},
		{384, 24, 8, 96},
		{384, 24, 96, 8},
		{384, 96, 8, 24},
		{384, 96, 24, 8},
	}
}

func save(
	trid, obid,
	t, i,

	V0, W0, V1,
	W1, V2, W2,

	x, y, z,

	R0, R1, R2,

	xSup, ySup, zSup,

	lastV0, lastV1, lastV2,
	lastW0, lastW1, lastW2,
	lastC0, lastC1, lastC2 int,
) {
	dbotmp += fmt.Sprintf(
		stmstr2,
		trid, obid,
		t, i,

		V0, W0, V1,
		W1, V2, W2,

		x, y, z,

		R0, R1, R2,

		xSup, ySup, zSup,

		lastV0, lastV1, lastV2,
		lastW0, lastW1, lastW2,
		lastC0, lastC1, lastC2,
	)

	if psc++; psc < MAXPSC {
		return
	}

	var stm2, err = dbo.Query(fmt.Sprintf(stmstr1, serie) + strings.ReplaceAll(dbotmp, ")(", "), ("))
	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring queries if you are using transactions
	stm2.Close()

	dbotmp = ""
	psc = 0
}

func finish() {
	if dbotmp == "" {
		return
	}

	var stm2, err = dbo.Query(fmt.Sprintf(stmstr1, serie) + strings.ReplaceAll(dbotmp, ")(", "), ("))
	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring queries if you are using transactions
	stm2.Close()

	dbotmp = ""
	psc = 0
}

func run(treat []int, trid, obid int) {
	var grid = ca.NewGrid(N)

	// initialize local pseudorandom generator
	rand.Seed(time.Now().UnixNano())

	// initialize grid
	var i, j, k, t int
	k = 0
	for j = range treat {
		for i = 0; i < treat[j]; i++ {
			grid.Add(ca.NewAgent(j), k)
			k++
		}
	}

	grid.Shuffle()

	//
	for i = 0; i < N; i++ {
		grid[i].Prepare()
	}

	// explore and transfer X
	var next = make([][]int, N)
	// the size of the neighborhood
	// var m = 27
	var x, y, z int
	// var x = -1
	// do it t times
	for t = 0; t < T; t++ {
		// look an agent j N times by every agent i
		for j = 0; j < N; j++ {
			// randomly
			for _, i = range grid.Shuffled() {
				// then, if "next" is null, initialize it
				if j == 0 {
					next[i] = grid.Shuffled() // [i].Neighborhood()
				}

				// grid[i].Search()
				if next[i][j] == -1 {
					continue
				}

				// Strategy 1: look at all, then transfer from the best.
				grid[i].Look(next[i][j], j == 0)
				if j == N-1 { // m-1
					// if /*(t == 0 || t == 127) &&*/ grid[i].Has(ca.X) > 0
					// x = -1
					/*
						fmt.Printf(
							"%d; %d; %d; %d; %d; %d; %d; %d; %d; %d; %d; %d; %d; %d; %d; %d; %d; %d; %d; %d; %d; %d; %d\n",
							t, i,
							grid[i].GetDelta1()-grid[i].GetDelta0(), grid[i].GetDelta2()-grid[i].GetDelta1(),
							grid[i].GetDelta0(), grid[i].GetDelta1(), grid[i].GetDelta2(),
							grid[i].V0, grid[i].W0,
							grid[i].V1, grid[i].W1,
							grid[i].V2, grid[i].W2,
							grid[i].Has(ca.X),
							grid[i].R0, grid[i].R1, grid[i].R2,
							grid[i].GetV0(), grid[i].GetV1(), grid[i].GetV2(),
							grid[i].GetW0(), grid[i].GetW1(), grid[i].GetW2(),
						)
					*/
					x, y, z = grid[i].XSupplier(), grid[i].YSupplier(), grid[i].ZSupplier()

					grid[i].Request(ca.X)
					grid[i].Request(ca.Y)
					grid[i].Request(ca.Z)

					// Audit 1 begin
					/*
						if x != -1 && grid[i].LastW0() == 0 && grid[i].LastW1() == 0 && grid[i].LastW2() == 0 {
							grid[i].Request(ca.X)
							panic("Error! The producer of X rejects the request.")
						}

						if y != -1 && grid[i].LastW0() == 0 && grid[i].LastW1() == 0 && grid[i].LastW2() == 0 {
							grid[i].Request(ca.Y)
							panic("Error! The producer of Y rejects the request.")
						}

						if z != -1 && grid[i].LastW0() == 0 && grid[i].LastW1() == 0 && grid[i].LastW2() == 0 {
							grid[i].Request(ca.Z)
							panic("Error! The producer of Z rejects the request.")
						}
					*/
					// end
					/*
						if _, err := stm.Exec(
							trid, obid,
							t, grid[i].ID,
							grid[i].V0, grid[i].W0,
							grid[i].V1, grid[i].W1,
							grid[i].V2, grid[i].W2,
							grid[i].Has(ca.X),
							grid[i].Has(ca.Y),
							grid[i].Has(ca.Z),
							grid[i].R0, grid[i].R1, grid[i].R2,
							x, y, z,
							grid[i].LastV0(), grid[i].LastV1(), grid[i].LastV2(),
							grid[i].LastW0(), grid[i].LastW1(), grid[i].LastW2(),
							grid[i].LastC0(), grid[i].LastC1(), grid[i].LastC2(),
						); err != nil {
							panic(err.Error())
						}
					*/

					save(
						trid, obid,
						t, grid[i].ID,
						grid[i].V0, grid[i].W0,
						grid[i].V1, grid[i].W1,
						grid[i].V2, grid[i].W2,
						grid[i].Has(ca.X),
						grid[i].Has(ca.Y),
						grid[i].Has(ca.Z),
						grid[i].R0, grid[i].R1, grid[i].R2,
						x, y, z,
						grid[i].LastV0(), grid[i].LastV1(), grid[i].LastV2(),
						grid[i].LastW0(), grid[i].LastW1(), grid[i].LastW2(),
						grid[i].LastC0(), grid[i].LastC1(), grid[i].LastC2(),
					)

				}
			}
		}

		// next the learning
		for i = 0; i < N; i++ {
			grid[i].Learn()
		}
	}
}

func main() {
	serie = os.Args[1]
	ca.REWARDENABLED = serie == "1"

	var err error
	dbo, err = sql.Open("mysql", url)
	if err != nil {
		panic(err.Error())
	}
	defer dbo.Close()
	/*
		stm, err = dbo.Prepare(stmstr)
		if err != nil {
			panic(err.Error())
		}
	*/
	var treats = getTreatments()
	var i, j int
	for i = 0; i < len(treats); i++ {
		for j = 0; j < OBS; j++ {
			run(treats[i], i, j)
			// return
		}
	}

	finish()
}
