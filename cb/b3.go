package cb

import (
	"lab3/ca"
	"math/rand"
	"time"
)

// Cantidad de agentes en la simulacion
const N = 512

// Cantidad de iteraciones
const T = 256

func save(
	trid, obid,
	t int, grid ca.Grid, i,

	x, y, z int,
) {
}

func run(treat []int, trid, obid int) {
	var grid = ca.NewGrid(N)
	// Inicializacion del generador de numeros pseudoaleatorios
	rand.Seed(time.Now().UnixNano())
	// Inicializacion de la matriz de agentes
	var i, j, k, t int
	k = 0
	for j = range treat {
		for i = 0; i < treat[j]; i++ {
			grid.Add(ca.NewAgent(j), k)
			k++
		}
	}

	grid.Shuffle()
	for i = 0; i < N; i++ {
		grid[i].Prepare()
	}
	// Busqueda y transferencia de recursos
	var next = make([][]int, N)
	var r, g, b int
	// En cada iteracion
	for t = 0; t < T; t++ {
		// Cada agente i analiza a cada agente k
		for j = 0; j < N; j++ {
			// El agente i es elegido aleatoriamente
			for _, i = range grid.Shuffled() {
				// El arreglo "next" es inicializado si es nulo
				if j == 0 {
					next[i] = grid.Shuffled()
				}
				// El agente i analiza al agente k = next[i][j]
				grid[i].Match(next[i][j], j == 0)
				if j == N-1 {
					r = grid[i].RSupplier()
					g = grid[i].GSupplier()
					b = grid[i].BSupplier()

					grid[i].Request(ca.R)
					grid[i].Request(ca.G)
					grid[i].Request(ca.B)

					save(trid, obid, t, grid, i, r, g, b)
				}
			}
		}
		// Cada agente i ajusta su criterio de decision
		for i = 0; i < N; i++ {
			grid[i].Learn()
		}
	}
}
