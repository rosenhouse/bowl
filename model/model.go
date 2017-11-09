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

type NonFinalFrame []Throw

func (f NonFinalFrame) IsSpare() bool {
	return len(f) == 2 && f[1] == ThrowSpare
}

func (f NonFinalFrame) IsStrike() bool {
	return len(f) == 1 && f[0] == ThrowStrike
}

func (f NonFinalFrame) PinsKnockedDown() []int {
	p := []int{f[0].PinsDown()}
	if f[0] == ThrowStrike {
		return p
	}
	if f[1] == ThrowSpare {
		p = append(p, 10-p[0])
	} else {
		p = append(p, f[1].PinsDown())
	}
	return p
}

type FinalFrame []Throw

func (f FinalFrame) PinsKnockedDown() []int {
	p := []int{f[0].PinsDown()}
	if f[1] == ThrowSpare {
		p = append(p, 10-p[0])
	} else {
		p = append(p, f[1].PinsDown())
	}
	if p[0]+p[1] >= 10 {
		p = append(p, f[2].PinsDown())
	}
	return p
}

type Frame interface {
	PinsKnockedDown() []int
}

func ComputePinsKnockedDown(frames []Frame) []int {
	var p []int
	for _, frame := range frames {
		p = append(p, frame.PinsKnockedDown()...)
	}
	return p
}

func sum(xs []int) int {
	var total = 0
	for _, x := range xs {
		total += x
	}
	return total
}

func ScoreFrames(s []Frame) int {
	pinsKnockedDown := ComputePinsKnockedDown(s)
	switch f := s[0].(type) {
	case NonFinalFrame:
		if f.IsStrike() {
			return 10 + pinsKnockedDown[1] + pinsKnockedDown[2]
		}
		if f.IsSpare() {
			return 10 + pinsKnockedDown[2]
		}
		return sum(f.PinsKnockedDown())
	case FinalFrame:
		return sum(f.PinsKnockedDown())
	default:
		panic("unrecognized frame type")
	}
}
