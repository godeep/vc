package vc

import "testing"

func TestUpdate(t *testing.T) {
	var v Vector

	// Append

	v = v.Update(42)
	expected := Vector{clock{42, 1}}

	if v.Compare(expected) != Equal {
		t.Errorf("Update error, %+v != %+v", v, expected)
	}

	// Insert at front

	v = v.Update(36)
	expected = Vector{clock{36, 1}, clock{42, 1}}

	if v.Compare(expected) != Equal {
		t.Errorf("Update error, %+v != %+v", v, expected)
	}

	// Insert in moddle

	v = v.Update(37)
	expected = Vector{clock{36, 1}, clock{37, 1}, clock{42, 1}}

	if v.Compare(expected) != Equal {
		t.Errorf("Update error, %+v != %+v", v, expected)
	}

	// Update existing

	v = v.Update(37)
	expected = Vector{clock{36, 1}, clock{37, 2}, clock{42, 1}}

	if v.Compare(expected) != Equal {
		t.Errorf("Update error, %+v != %+v", v, expected)
	}
}

func TestCopy(t *testing.T) {
	v0 := Vector{clock{42, 1}}
	v1 := v0.Copy()
	v1.Update(42)
	if v0.Compare(v1) != Ancestor {
		t.Errorf("Copy error, %+v should be ancestor of %+v", v0, v1)
	}
}

func TestMerge(t *testing.T) {
	testcases := []struct {
		a, b, m Vector
	}{
		// No-ops
		{
			Vector{},
			Vector{},
			Vector{},
		},
		{
			Vector{clock{22, 1}, clock{42, 1}},
			Vector{clock{22, 1}, clock{42, 1}},
			Vector{clock{22, 1}, clock{42, 1}},
		},

		// Appends
		{
			Vector{},
			Vector{clock{22, 1}, clock{42, 1}},
			Vector{clock{22, 1}, clock{42, 1}},
		},
		{
			Vector{clock{22, 1}},
			Vector{clock{42, 1}},
			Vector{clock{22, 1}, clock{42, 1}},
		},
		{
			Vector{clock{22, 1}},
			Vector{clock{22, 1}, clock{42, 1}},
			Vector{clock{22, 1}, clock{42, 1}},
		},

		// Insert
		{
			Vector{clock{22, 1}, clock{42, 1}},
			Vector{clock{22, 1}, clock{23, 2}, clock{42, 1}},
			Vector{clock{22, 1}, clock{23, 2}, clock{42, 1}},
		},
		{
			Vector{clock{42, 1}},
			Vector{clock{22, 1}},
			Vector{clock{22, 1}, clock{42, 1}},
		},

		// Update
		{
			Vector{clock{22, 1}, clock{42, 2}},
			Vector{clock{22, 2}, clock{42, 1}},
			Vector{clock{22, 2}, clock{42, 2}},
		},

		// All of the above
		{
			Vector{clock{10, 1}, clock{20, 2}, clock{30, 1}},
			Vector{clock{5, 1}, clock{10, 2}, clock{15, 1}, clock{20, 1}, clock{25, 1}, clock{35, 1}},
			Vector{clock{5, 1}, clock{10, 2}, clock{15, 1}, clock{20, 2}, clock{25, 1}, clock{30, 1}, clock{35, 1}},
		},
	}

	for i, tc := range testcases {
		if m := tc.a.Merge(tc.b); m.Compare(tc.m) != Equal {
			t.Errorf("%d: %+v.Merge(%+v) == %+v (expected %+v)", i, tc.a, tc.b, m, tc.m)
		}
	}

}
