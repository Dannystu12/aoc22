package day9

import (
	"fmt"
	"math"
)

type longRope struct {
	head  ropeHead
	knots []Point
	tail  ropeTail
}

func NewLongRope(initial Point, numKnots uint) (*longRope, error) {
	if numKnots < 1 {
		return nil, fmt.Errorf("long rope must have at least one knot")
	}

	lr := longRope{
		head:  ropeHead{initial},
		knots: make([]Point, numKnots),
		tail:  ropeTail{initial, map[Point]bool{initial: true}},
	}

	for i := 0; i < int(numKnots); i++ {
		lr.knots[i] = initial
	}

	return &lr, nil
}

func (r *longRope) Move(m move) {
	for i := 0; i < m.amount; i++ {
		r.move(m.direction)
	}
}

func (r *longRope) move(dir direction) {
	r.head.move(dir)
	previousNodeNew := r.head.Point

	for i := 0; i < len(r.knots); i++ {
		r.knots[i] = getNextPos(r.knots[i], previousNodeNew)
		previousNodeNew = r.knots[i]
	}

	r.tail.Point = getNextPos(r.tail.Point, previousNodeNew)
	r.tail.history[r.tail.Point] = true
}

func getNextPos(thisNode, parentNode Point) Point {
	xDif := parentNode.x - thisNode.x
	yDif := parentNode.y - thisNode.y
	absYDif := int(math.Abs(float64(yDif)))
	absXDif := int(math.Abs(float64(xDif)))

	yMove := 0
	if absYDif > 0 {
		yMove = yDif / absYDif
	}

	xMove := 0
	if absXDif > 0 {
		xMove = xDif / absXDif
	}

	diagonalMove := absYDif > 1 && absXDif > 0 || absYDif > 0 && absXDif > 1

	if diagonalMove {
		return Point{thisNode.x + xMove, thisNode.y + yMove}
	} else if absXDif > 1 {
		return Point{thisNode.x + xMove, thisNode.y}
	} else if absYDif > 1 {
		return Point{thisNode.x, thisNode.y + yMove}
	} else {
		return thisNode
	}

}

func (r *longRope) CountTailPositions() int {
	return len(r.tail.history)
}
