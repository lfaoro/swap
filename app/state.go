package app

import "slices"

type State int

const (
	CoinTable State = iota
	InputAmount
	RateTable
	InputAddress
	InputRefundAddress
	TrxStatus
)

func (s State) String() string {
	names := map[State]string{
		CoinTable:          "CoinTable",
		InputAmount:        "InputAmount",
		InputAddress:       "InputAddress",
		InputRefundAddress: "InputRefundAddress",
		RateTable:          "RateTable",
		TrxStatus:          "TransactionStatus",
	}
	return names[s]
}

type AppState struct {
	states  []State
	current int
}

func (s *AppState) Init() *AppState {
	s.states = []State{
		CoinTable,
		InputAmount,
		RateTable,
		InputAddress,
		InputRefundAddress,
		TrxStatus,
	}
	s.current = 0
	return s
}

func (s *AppState) Next() {
	if s.current >= len(s.states)-1 {
		return
	}
	s.current++
}

func (s *AppState) Prev() {
	if s.current == 0 {
		return
	}
	s.current--
}

func (s *AppState) Current() State {
	return s.states[s.current]
}

func (s *AppState) CurrentString() string {
	return s.Current().String()
}

func (s *AppState) GoTo(state State) bool {
	for i, st := range s.states {
		if st == state {
			s.current = i
			return true
		}
	}
	return false
}

func (s *AppState) IsAt(state State) bool {
	return s.Current() == state
}

// IsOneOf checks if the current state is one of the provided states.
func (s *AppState) IsOneOf(states ...State) bool {
	current := s.Current()
	return slices.Contains(states, current)
}
