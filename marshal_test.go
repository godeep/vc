package vc

import "testing"

func TestMarshalling(t *testing.T) {
	testcases := []Vector{
		{},
		{clock{1, 2}},
		{clock{1, 2}, clock{3, 4}, clock{5, 6}},
	}

	for _, v0 := range testcases {
		bs, err := v0.MarshalBinary()
		if err != nil {
			t.Fatal(err)
		}

		var v1 Vector
		err = v1.UnmarshalBinary(bs)
		if err != nil {
			t.Fatal(err)
		}

		if v0.Compare(v1) != Equal {
			t.Errorf("%+v != %+v", v0, v1)
		}
	}
}
