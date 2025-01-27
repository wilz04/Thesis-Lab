package ca

// Interruptor que habilita o deshabilita la oferta de recompensas
var REWARDENABLED = false

type Agent struct {
	ID int

	V0, W0, V1, W1, V2, W2, R0, R1, R2 int

	source, target                  *Color
	rSupplier, gSupplier, bSupplier *Supplier

	grid    Grid
	history History
}
