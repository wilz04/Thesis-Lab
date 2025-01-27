package ca

import (
	"math/rand"
	"time"
)

/* Maximo valor para cualquiera de los componentes del criterio de
* decision */
const MAX = 256

func NewAgent(profile int) (me *Agent) {
	me = &Agent{
		rSupplier: nil,
		gSupplier: nil,
		bSupplier: nil,

		source: NewColor(),
		target: NewColor(),

		history: NewHistory(),
	}

	// Inicializacion del generador de numeros pseudoaleatorios
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
		if (me.V0 > 0 && me.V1 > 0 && me.V2 > 0) &&
			(me.W0 > 0 && me.W1 > 0 && me.W2 > 0) &&
			(d0 > 0 && d1 > 0 && d2 > 0) && ((profile == 0) &&
			(d1-d0 > -MAX && d1-d0 <= 0) &&
			(d2-d1 > -MAX && d2-d1 <= 0) ||

			(profile == 1) &&
				(d1-d0 > -MAX && d1-d0 <= 0) &&
				(d2-d1 >= 0 && d2-d1 < MAX) ||

			(profile == 2) &&
				(d1-d0 >= 0 && d1-d0 < MAX) &&
				(d2-d1 > -MAX && d2-d1 <= 0) ||

			(profile == 3) &&
				(d1-d0 >= 0 && d1-d0 < MAX) &&
				(d2-d1 >= 0 && d2-d1 < MAX)) {

			return
		}
	}
}
