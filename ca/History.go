package ca

// JUMP is the default jump size.
const JUMP = 8

// Score structure
type Score struct {
	axis int
	jump int

	v0 int
	v1 int
	v2 int

	w0 int
	w1 int
	w2 int

	c0 int
	c1 int
	c2 int
}

// History structure
type History []*Score

// NewHistory initializes a new agent's transfer's history.
func NewHistory() (me History) {
	me = make(History, 1)
	me[0] = &Score{jump: 0, v0: 0, v1: 0, v2: 0, w0: 0, w1: 0, w2: 0, c0: 0, c1: 0, c2: 0}
	return
}

// archive adds a new history score.
func (me History) archive() History {
	var result = make(History, len(me)+1)
	copy(result, me)
	copy(result[len(me):], History{&Score{jump: JUMP, v0: 0, v1: 0, v2: 0, w0: 0, w1: 0, w2: 0, c0: 0, c1: 0, c2: 0}})
	return result
}

// logreq records a new delta-transaction request. (partially certified)
func (me History) logreq(x rune, axis, value int) {
	var i = len(me) - 1
	switch axis {
	case 0:
		if me[i].w0 < value {
			me[i].w0 = value
		}
	case 11:
		if me[i].w1 < value {
			me[i].w1 = value
		}
	case 12:
		if me[i].w1 < value {
			me[i].w1 = value
		}
	case 21:
		if me[i].w2 < value {
			me[i].w2 = value
		}
	case 22:
		if me[i].w2 < value {
			me[i].w2 = value
		}
	}
}

// logres records a new delta-transaction response. (partially certified)
func (me History) logres(x rune, axis, value int) {
	var i = len(me) - 1
	switch axis {
	case 0:
		me[i].v0 += value
		me[i].c0++
	case 11:
		me[i].v1 += value
		me[i].c1++
	case 21:
		me[i].v1 += value
		me[i].c1++
	case 12:
		me[i].v2 += value
		me[i].c2++
	case 22:
		me[i].v2 += value
		me[i].c2++
	}
}
