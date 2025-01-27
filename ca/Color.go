package ca

// R is a type of resource
const R = 'X'

// G is a type of resource
const G = 'Y'

// B is a type of resource
const B = 'Z'

// Color structure
type Color struct {
	R int
	G int
	B int
}

// NewColor initializes a new color.
func NewColor() (me *Color) {
	me = &Color{R: 0, G: 0, B: 0}
	return
}
