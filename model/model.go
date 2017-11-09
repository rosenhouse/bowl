package model

type Throw int

const ThrowStrike = Throw(0xAAAA)
const ThrowSpare = Throw(0xBBBB)

type Frame []Throw

func (f Frame) PinsDownInFirstThrow() int {
	if len(f) < 1 {
		return -1
	}
	if f[0] == ThrowStrike {
		return 10
	}
	return int(f[0])
}

func (f Frame) PinsDownTotal() int {
	panic("not implemented")
}
