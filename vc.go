package vc

// The Vector type represents a vector clock. The zero value is a usable
// vector clock. The individual indexes consist of ID-to-value mappings, where
// both the ID and value are non-negative 64 bit integers. This results in an
// effective 63 bit ID and value space.
type Vector []clock

// Clock represents one index in the vector. ID is a non-negative integer.
// Value is a non-negative integer, zero being the initial clock state.
type clock struct {
	id    int64
	value int64
}

// Update returns a new Vector with the index for the specific ID incremented
// by one. The returned vector may be shared with v where possible or may be a
// new allocation.
func (v Vector) Update(id int64) Vector {
	for i := range v {
		if v[i].id == id {
			// Update an existing index
			v[i].value++
			return v
		} else if v[i].id > id {
			// Insert a new index
			nv := make(Vector, len(v)+1)
			copy(nv, v[:i])
			nv[i].id = id
			nv[i].value = 1
			copy(nv[i+1:], v[i:])
			return nv
		}
	}
	// Append a new new index
	return append(v, clock{id, 1})
}

// Merge returns the vector containing the maximum indexes from a and b. The
// returned vector may be shared with a where possible or may be a new
// allocation.
func (a Vector) Merge(b Vector) Vector {
	var ai, bi int
	for bi < len(b) {
		if ai == len(a) {
			// We've reach the end of a, all that remains are appends
			return append(a, b[bi:]...)
		}

		if a[ai].id > b[bi].id {
			// The index from b should be inserted here
			n := make(Vector, len(a)+1)
			copy(n, a[:ai])
			n[ai] = b[bi]
			copy(n[ai+1:], a[ai:])
			a = n
		}

		if a[ai].id == b[bi].id {
			if v := b[bi].value; v > a[ai].value {
				a[ai].value = v
			}
		}

		if bi < len(b) && a[ai].id == b[bi].id {
			bi++
		}
		ai++
	}

	return a
}

// Copy returns an identical vector that is not shared with v.
func (v Vector) Copy() Vector {
	nv := make(Vector, len(v))
	copy(nv, v)
	return nv
}
