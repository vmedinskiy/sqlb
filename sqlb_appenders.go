package sqlb

type Int64A int64

func (a Int64A) String() string {
	res, _ := String(int64(a))
	return res
}

func (a Int64A) AppendTo(buf []byte) []byte {
	out := append(buf, []byte(a.String())...)
	return out
}

type StringQA string

func (sq StringQA) String() string {
	return `"` + Str(sq).String() + `"`
}

func (sq StringQA) AppendTo(buf []byte) []byte {
	return Str(sq.String()).AppendTo(buf)
}
