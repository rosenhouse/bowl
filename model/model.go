package model

type Throw int

const ThrowStrike = Throw(0xAAAA)
const ThrowSpare = Throw(0xBBBB)

func (t Throw) PinsDown() int {
	if t == ThrowSpare {
		panic("who knows!")
	}
	if t == ThrowStrike {
		return 10
	}
	return int(t)
}

type Frame []Throw

func (f Frame) PinsDownInFirstThrow() int {
	if len(f) < 1 {
		return -1
	}
	return f[0].PinsDown()
}

func (f Frame) PinsDownTotal() int {
	switch len(f) {
	case 0:
		return -1
	case 1:
		if f[0] == ThrowStrike {
			return 10
		}
		return -1
	case 2:
		if f[1] == ThrowSpare {
			return 10
		}
		return int(f[0]) + int(f[1])
	case 3: // 10th frame
		if f[1] == ThrowSpare {
			return 10 + f[2].PinsDown()
		}
		return 20 + f[2].PinsDown()
	default:
		panic("invalid frame")
	}
}
