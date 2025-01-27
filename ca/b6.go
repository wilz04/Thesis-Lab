package ca

func (me *Agent) Match(j int, reset bool) {
	var he = me.grid[j]
	var value, action = me.free(he)
	if reset {
		me.rSupplier, me.gSupplier, me.bSupplier = nil, nil, nil
	}

	if action == -1 {
		return
	}

	var supplier = NewSupplier(he, value, action)
	if me.Has(R) < 0 && he.Has(R) > 0 {
		if me.rSupplier != nil {
			if value > me.rSupplier.value {
				me.rSupplier = supplier
			}
		} else {
			me.rSupplier = supplier
		}
	}

	if me.Has(G) < 0 && he.Has(G) > 0 {
		if me.gSupplier != nil {
			if value > me.gSupplier.value {
				me.gSupplier = supplier
			}
		} else {
			me.gSupplier = supplier
		}
	}

	if me.Has(B) < 0 && he.Has(B) > 0 {
		if me.bSupplier != nil {
			if value > me.bSupplier.value {
				me.bSupplier = supplier
			}
		} else {
			me.bSupplier = supplier
		}
	}
}
