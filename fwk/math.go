package fwk

import "golang.org/x/exp/constraints"

func Abs[V constraints.Integer](a V) V {
	if a < 0 {
		return -a
	}
	return a
}

func Max[V constraints.Integer](a, b V) V {
	if a > b {
		return a
	}

	return b
}

func Min[V constraints.Integer](a, b V) V {
	if a < b {
		return a
	}

	return b
}
