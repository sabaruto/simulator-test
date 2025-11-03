package common

import (
	"math"
	"time"
)

type WindowManager interface {
	Start([]*Object)
	SetFrameRate(frameRate int)
}

type Object interface {
	GetPosition() Position
	Move(diffTime time.Duration)
	Draw()
	SetObjectManager(om *ObjectManager)
}

type Position struct {
	X float64
	Y float64
}

func Distance(pos1 Position, pos2 Position) float64 {
	xDiff := pos2.X - pos1.X
	yDiff := pos2.Y - pos1.Y

	sqrdDistance := xDiff + yDiff

	return math.Sqrt(sqrdDistance)
}

func (p Position) Distance(otherPos Position) float64 {
	return Distance(p, otherPos)
}
