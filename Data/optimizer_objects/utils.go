package optimizerobjects

import "golang.org/x/exp/constraints"

func bool_to_int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func abs_val[L constraints.Integer | constraints.Float](l L) L {
	if l < 0 {
		return L(-l)
	}
	return L(l)
}
