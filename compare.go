package vc

// Ordering represents the relationship between two Vectors.
type Ordering int

const (
	Equal Ordering = iota
	Descendant
	Ancestor
	Concurrent
)

const idMissing = -1

// Compare returns the Ordering that describes a's relation to b.
func (a Vector) Compare(b Vector) Ordering {
	var ai, bi int   // index into a and b
	var av, bv clock // value at current index

	result := Equal

	for ai < len(a) || bi < len(b) {
		if ai < len(a) {
			av = a[ai]
		} else {
			av = clock{id: idMissing}
		}

		if bi < len(b) {
			bv = b[bi]
		} else {
			bv = clock{id: idMissing}
		}

		switch {
		case av.id == bv.id:
			// We have a clock value for each side
			if av.value > bv.value {
				if result == Ancestor {
					return Concurrent
				}
				result = Descendant
			} else if av.value < bv.value {
				if result == Descendant {
					return Concurrent
				}
				result = Ancestor
			}

		case av.id > 0 && av.id < bv.id || bv.id < 0:
			// Value is missing on the b side
			if av.value > 0 {
				if result == Ancestor {
					return Concurrent
				}
				result = Descendant
			}

		case bv.id > 0 && bv.id < av.id || av.id < 0:
			// Value is missing on the a side
			if bv.value > 0 {
				if result == Descendant {
					return Concurrent
				}
				result = Ancestor
			}
		}

		if ai < len(a) && (av.id <= bv.id || bv.id < 0) {
			ai++
		}
		if bi < len(b) && (bv.id <= av.id || av.id < 0) {
			bi++
		}
	}

	return result
}
