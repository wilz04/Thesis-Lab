package main

import (
	"database/sql"
	"fmt"
	"lab3/cb"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

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

func main() {
	cb.Main()
}
