package calculator

import "golang.org/x/exp/constraints"

func CalNextValue[V constraints.Integer](in []V) (V, V) {
	diffs := [][]V{}

	for diffL := in; !isAllEqual(diffL, 0); diffL = diff(diffL) {
		diffs = append(diffs, diffL)
	}

	rVal := V(0)
	lVal := V(0)
	for i := len(diffs) - 1; i >= 0; i-- {
		diffL := diffs[i]

		rVal += diffL[len(diffL)-1]
		lVal = diffL[0] - lVal
	}

	return lVal, rVal
}

func diff[V constraints.Integer](in []V) []V {
	out := make([]V, len(in)-1)
	for i := 0; i < len(in)-1; i++ {
		out[i] = in[i+1] - in[i]
	}

	return out
}

func isAllEqual[V constraints.Integer](in []V, val V) bool {
	for _, v := range in {
		if v != val {
			return false
		}
	}
	return true
}
