package model

type Throw int

const ThrowStrike = Throw(0xAAAA)
const ThrowSpare = Throw(0xBBBB)

type Frame []Throw

func (f Frame) PinsDownInFirstThrow() int {
	panic("not implemented")
}

func (f Frame) PinsDownTotal() int {
	panic("not implemented")
}
