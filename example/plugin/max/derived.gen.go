// Code generated by goderive DO NOT EDIT.

package max

func deriveCompare(this, that boat) int {
	return deriveCompare_(&this, &that)
}

func deriveMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func deriveMaxs(list []boat, def boat) boat {
	if len(list) == 0 {
		return def
	}
	m := list[0]
	list = list[1:]
	for i, v := range list {
		if deriveCompare(v, m) > 0 {
			m = list[i]
		}
	}
	return m
}

func deriveCompare_(this, that *boat) int {
	if this == nil {
		if that == nil {
			return 0
		}
		return -1
	}
	if that == nil {
		return 1
	}
	if c := deriveCompare_i(this.length, that.length); c != 0 {
		return c
	}
	return 0
}

func deriveCompare_i(this, that int) int {
	if this != that {
		if this < that {
			return -1
		} else {
			return 1
		}
	}
	return 0
}
