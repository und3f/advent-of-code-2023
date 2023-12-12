package state

type State interface {
	IsEnd() bool
	Next(rune) State
}

type StateSpring struct {
	NextDamaged, NextWorking State
}

type StateEnd struct{}

func New(damaged []uint64) State {
	var next State = &StateEnd{}

	for i := len(damaged) - 1; i >= 0; i-- {
		if !next.IsEnd() {
			next = &StateSpring{NextWorking: next}
		}

		next = NewDamaged(damaged[i], next)
	}

	return next
}

func NewDamaged(size uint64, following State) State {
	next := &StateSpring{NextDamaged: following}

	for i := uint64(1); i < size; i++ {
		next = &StateSpring{NextDamaged: next}
	}
	next.NextWorking = next

	return next
}

func (s *StateSpring) Next(r rune) State {
	switch r {
	case '.':
		return s.NextWorking
	case '#':
		return s.NextDamaged
	}
	return nil
}

func (s *StateSpring) IsEnd() bool {
	return false
}

func (s *StateEnd) Next(sym rune) State {
	if sym == '.' {
		return s
	}

	return nil
}

func (s *StateEnd) IsEnd() bool {
	return true
}
