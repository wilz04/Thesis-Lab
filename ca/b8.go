package ca

func (me *Agent) SetSource(rgb int) {
	var rg = rgb % 64
	me.source.B = (rgb - rg) / 64
	me.source.R = rg % 8
	me.source.G = (rg - me.source.R) / 8
}

func (me *Agent) SetTarget(rgb int) {
	me.ID = rgb
	var rg = rgb % 64
	me.target.B = (rgb - rg) / 64
	me.target.R = rg % 8
	me.target.G = (rg - me.target.R) / 8
}
