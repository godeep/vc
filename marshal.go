package vc

import "encoding/binary"

// MarshalBinary returns a binary (XDR compliant) encoding of the Vector
func (v Vector) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 4+16*len(v))
	binary.BigEndian.PutUint32(bs, uint32(len(v)))
	for i := range v {
		binary.BigEndian.PutUint64(bs[4+i*16:], uint64(v[i].id))
		binary.BigEndian.PutUint64(bs[4+i*16+8:], uint64(v[i].value))
	}
	return bs, nil
}

// UnmarshalBinary unmarshals the binary representation of a Vector into itself.
func (v *Vector) UnmarshalBinary(bs []byte) error {
	l := int(binary.BigEndian.Uint32(bs))
	n := make([]clock, l)
	for i := range n {
		n[i].id = int64(binary.BigEndian.Uint64(bs[4+i*16:]))
		n[i].value = int64(binary.BigEndian.Uint64(bs[4+i*16+8:]))
	}
	*v = n
	return nil
}
