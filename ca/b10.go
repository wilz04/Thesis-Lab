package ca

func (me *Agent) Response(rt rune, action int) (bool, int) {
	var delta, value = 0, 0
	switch action {
	case 7:
		delta = me.GetDelta2() - me.GetReward()
		value = me.V2
	case 5:
		delta = me.GetDelta2() - me.GetReward()
		value = me.V2
	case 3:
		delta = me.GetDelta1()
		value = me.V1
	case 1:
		delta = me.GetDelta1()
		value = me.V1
	case 0:
		delta = me.GetDelta0()
		value = me.V0
	}

	if delta > 0 {
		me.history.logres(rt, action, delta)
		return true, value
	}

	return false, value
}
