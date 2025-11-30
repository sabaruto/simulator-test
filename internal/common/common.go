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
	GetID() int64
	Equal(other Object) bool
	GetPosition() vector.Vector
	Move(diffTime time.Duration)
	Draw()
	SetObjectManager(om *ObjectManager)
	ReceiveRay(rayPosition vector.Vector, rayDirection vector.Vector) (intersectPoint *vector.Vector, colour *string)
}

func Distance(a, b vector.Vector) float64 {
	return b.Sub(a).Magnitude()
}
