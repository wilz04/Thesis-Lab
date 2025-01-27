package ca

import (
	"math/rand"
	"time"
)

// Matriz de agentes
type Grid []*Agent

/* Metodo constructor, inicializa la matriz de agentes, con capacidad
* para n agentes */
func NewGrid(n int) (me Grid) {
	me = make(Grid, n)
	return
}

/* Metodo que agrega la referencia de un agente a la matriz de agentes,
* en una posicion determinada */
func (me Grid) Add(agent *Agent, i int) {
	me[i] = agent
	agent.grid = me
}

/* Metodo que baraja la matriz e inicializa la oferta y la demanda de
* cada agente */
func (me Grid) Shuffle() {
	// Inicializacion del generador de numeros pseudoaleatorios
	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(me), func(i, j int) {
		me[i], me[j] = me[j], me[i]

		me[i].SetSource(j)
		me[j].SetSource(i)

		me[i].SetTarget(i)
		me[j].SetTarget(j)
	})
}

/* Metodo que genera, baraja y retorna la secuencia de enteros entre 0
* y N - 1 */
func (me Grid) Shuffled() []int {
	// Inicializacion del generador de numeros pseudoaleatorios
	rand.Seed(time.Now().UnixNano())

	var seq = make([]int, len(me))
	var i = 0
	for i = range seq {
		seq[i] = i
	}

	rand.Shuffle(len(seq), func(i, j int) {
		seq[i], seq[j] = seq[j], seq[i]
	})

	return seq
}
