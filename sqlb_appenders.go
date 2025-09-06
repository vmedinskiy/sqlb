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

type StringDQA string // double quoted string

func (sq StringDQA) String() string {
	return `"` + Str(sq).String() + `"`
}

func (sq StringDQA) AppendTo(buf []byte) []byte {
	return Str(sq.String()).AppendTo(buf)
}

type StringSQA string // single quoted string

func (sq StringSQA) String() string {
	return "'" + Str(sq).String() + "'"
}

func (sq StringSQA) AppendTo(buf []byte) []byte {
	return Str(sq.String()).AppendTo(buf)
}
