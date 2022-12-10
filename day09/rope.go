package day09

import "math"

type ropeHead struct {
	Point
}

type ropeTail struct {
	Point
	history map[Point]bool
}

type rope struct {
	head ropeHead
	tail ropeTail
}

func NewRope() rope {
	return rope{
		head: ropeHead{},
		tail: ropeTail{
			history: map[Point]bool{Point{}: true},
		},
	}
}

func (r *rope) Move(m move) {
	for i := 0; i < m.amount; i++ {
		r.move(m.direction)
	}
}

func (r *rope) move(d direction) {
	oldHeadPos := r.head.Point
	r.head.move(d)

	if math.Abs(float64(r.tail.x-r.head.x)) > 1 || math.Abs(float64(r.tail.y-r.head.y)) > 1 {
		r.tail.Point = oldHeadPos
		r.tail.history[oldHeadPos] = true
	}
}

func (r *rope) CountTailPositions() int {
	return len(r.tail.history)
}
