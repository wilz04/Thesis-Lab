package ca

func (me *Agent) Request(rt rune) bool {
	if rt == R && me.rSupplier == nil ||
		rt == G && me.gSupplier == nil ||
		rt == B && me.bSupplier == nil {
		return false
	}

	var he *Agent
	var ok, value = false, 0
	switch rt {
	case R:
		he = me.grid[me.rSupplier.Id()]
		if ok, value = he.Response(rt, me.rSupplier.Action()); ok {
			me.history.logreq(rt, me.rSupplier.Action(), value)
			me.rSupplier = nil
			return true
		}
	case G:
		he = me.grid[me.gSupplier.Id()]
		if ok, value = he.Response(rt, me.gSupplier.Action()); ok {
			me.history.logreq(rt, me.gSupplier.Action(), value)
			me.gSupplier = nil
			return true
		}
	case B:
		he = me.grid[me.bSupplier.Id()]
		if ok, value = he.Response(rt, me.bSupplier.Action()); ok {
			me.history.logreq(rt, me.bSupplier.Action(), value)
			me.bSupplier = nil
			return true
		}
	}

	return false
}
