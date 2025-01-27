package ca

func (me *Agent) Has(rt rune) int {
	switch rt {
	case R:
		return me.source.R - me.target.R
	case G:
		return me.source.G - me.target.G
	default:
		return me.source.B - me.target.B
	}
}
