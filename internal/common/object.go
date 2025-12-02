package common

import (
	"time"

	"github.com/quartercastle/vector"
	"github.com/tfriedel6/canvas"
)

type Object interface {
	GetID() int64
	Equal(other Object) bool
	GetPosition() vector.Vector
	Move(diffTime time.Duration)
	Draw()
	SetObjectManager(om ObjectManager)
	ReceiveRay(rayPosition vector.Vector, rayDirection vector.Vector) (intersectPoint *vector.Vector, colour *string)
}

type ObjectManager interface {
	GetCanvas() *canvas.Canvas
	DrawObjects()
	MoveObjects(diffTime time.Duration)
	SendRay(rayPosition vector.Vector, rayDirection vector.Vector, ignoreIDs []int64) (*vector.Vector, *string)
}
