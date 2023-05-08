package ppcache

type ByteView []byte

func (b ByteView) Len() int {
	return len(b)
}

func (b ByteView) Slice() []byte {
	return b.Copy()
}

func (b ByteView) Copy() []byte {
	now := make([]byte, len(b))
	copy(now, b)
	return now
}

func (b ByteView) String() string {
	return string(b.Copy())
}
