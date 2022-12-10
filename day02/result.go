package day02

type result string

const (
	WIN            result = "WIN"
	LOSS           result = "LOSS"
	DRAW           result = "DRAW"
	INVALID_RESULT result = "INVALID_RESULT"
)

func (r result) toScore() int {
	switch r {
	case WIN:
		return 6
	case DRAW:
		return 3
	case LOSS:
		return 0
	default:
		return 0
	}
}

func (r result) Validate() (result, bool) {
	switch r {
	case WIN:
		return WIN, true
	case DRAW:
		return DRAW, true
	case LOSS:
		return LOSS, true
	default:
		return INVALID_RESULT, false
	}
}
