package util

type Spinner struct {
	curr int
	size int
	max  int
}

func NewSpinner(start, size int) *Spinner {
	return &Spinner{
		curr: start,
		size: size,
		max:  size - 1,
	}
}

func (s *Spinner) Spin(delta int) int {
	var (
		clicks int
	)

	// remove the size of the spinner from the spin count
	revolutions := delta / s.size
	clicks += abs(revolutions)

	remainder := delta % s.size

	// move the spinner the last bit
	if remainder > 0 {
		s.curr += remainder
		if s.curr > s.max {
			clicks++
			s.curr -= s.size
		}
	} else if remainder < 0 {
		start := s.curr
		s.curr += remainder
		if s.curr <= 0 && start > 0 {
			clicks++
		}
		if s.curr < 0 {
			s.curr += s.size
		}
	}

	return clicks
}

func (s *Spinner) Position() int {
	return s.curr
}