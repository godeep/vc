package vc

import "testing"

func TestCompare(t *testing.T) {
	testcases := []struct {
		a, b Vector
		r    Ordering
	}{
		// Empty vectors are identical
		{Vector{}, Vector{}, Equal},
		{Vector{}, nil, Equal},
		{nil, Vector{}, Equal},

		// Zero is the implied value for a missing Clock
		{Vector{clock{42, 0}}, Vector{clock{77, 0}}, Equal},

		// Equal vectors are equal
		{Vector{clock{42, 33}}, Vector{clock{42, 33}}, Equal},
		{Vector{clock{42, 33}, clock{77, 24}}, Vector{clock{42, 33}, clock{77, 24}}, Equal},

		// These a-vectors are all greater than the b-vector
		{Vector{clock{42, 1}}, nil, Descendant},
		{Vector{clock{42, 1}}, Vector{}, Descendant},
		{
			Vector{clock{42, 1}},
			Vector{clock{42, 0}},
			Descendant,
		},
		{
			Vector{clock{42, 2}},
			Vector{clock{42, 1}},
			Descendant,
		},
		{
			Vector{clock{22, 22}, clock{42, 2}},
			Vector{clock{22, 22}, clock{42, 1}},
			Descendant,
		},
		{
			Vector{clock{42, 2}, clock{77, 3}},
			Vector{clock{42, 1}, clock{77, 3}},
			Descendant,
		},
		{
			Vector{clock{22, 22}, clock{42, 2}, clock{77, 3}},
			Vector{clock{22, 22}, clock{42, 1}, clock{77, 3}},
			Descendant,
		},
		{
			Vector{clock{22, 23}, clock{42, 2}, clock{77, 4}},
			Vector{clock{22, 22}, clock{42, 1}, clock{77, 3}},
			Descendant,
		},

		// These a-vectors are all lesser than the b-vector
		{nil, Vector{clock{42, 1}}, Ancestor},
		{Vector{}, Vector{clock{42, 1}}, Ancestor},
		{
			Vector{clock{42, 0}},
			Vector{clock{42, 1}},
			Ancestor,
		},
		{
			Vector{clock{42, 1}},
			Vector{clock{42, 2}},
			Ancestor,
		},
		{
			Vector{clock{22, 22}, clock{42, 1}},
			Vector{clock{22, 22}, clock{42, 2}},
			Ancestor,
		},
		{
			Vector{clock{42, 1}, clock{77, 3}},
			Vector{clock{42, 2}, clock{77, 3}},
			Ancestor,
		},
		{
			Vector{clock{22, 22}, clock{42, 1}, clock{77, 3}},
			Vector{clock{22, 22}, clock{42, 2}, clock{77, 3}},
			Ancestor,
		},
		{
			Vector{clock{22, 22}, clock{42, 1}, clock{77, 3}},
			Vector{clock{22, 23}, clock{42, 2}, clock{77, 4}},
			Ancestor,
		},

		// These are all in conflict
		{
			Vector{clock{42, 2}},
			Vector{clock{43, 1}},
			Concurrent,
		},
		{
			Vector{clock{43, 1}},
			Vector{clock{42, 2}},
			Concurrent,
		},
		{
			Vector{clock{22, 23}, clock{42, 1}},
			Vector{clock{22, 22}, clock{42, 2}},
			Concurrent,
		},
		{
			Vector{clock{22, 21}, clock{42, 2}},
			Vector{clock{22, 22}, clock{42, 1}},
			Concurrent,
		},
		{
			Vector{clock{22, 21}, clock{42, 2}, clock{43, 1}},
			Vector{clock{20, 1}, clock{22, 22}, clock{42, 1}},
			Concurrent,
		},
	}

	for i, tc := range testcases {
		if r := tc.a.Compare(tc.b); r != tc.r {
			t.Errorf("%d: %+v.Compare(%+v) == %v (expected %v)", i, tc.a, tc.b, r, tc.r)
		}
	}
}
