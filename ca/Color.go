package ca

// X is a type of resource
const X = 'X'

// Y is a type of resource
const Y = 'Y'

// Z is a type of resource
const Z = 'Z'

// Color structure
type Color struct {
	X int
	Y int
	Z int
}

// NewColor initializes a new color.
func NewColor() (me *Color) {
	me = &Color{X: 0, Y: 0, Z: 0}
	return
}
