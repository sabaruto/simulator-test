package common

import (
	"github.com/quartercastle/vector"
	"time"
)

type WindowManager interface {
	Start([]*Object)
	SetFrameRate(frameRate int)
}

type Object interface {
	GetPosition() vector.Vector
	Move(diffTime time.Duration)
	Draw()
	SetObjectManager(om *ObjectManager)
}

func Distance(a, b vector.Vector) float64 {
	return b.Sub(a).Magnitude()
}
