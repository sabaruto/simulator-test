package common

import (
	"sort"
	"time"

	"github.com/quartercastle/vector"
	"github.com/tfriedel6/canvas"
)

type ObjectManager struct {
	canvas  *canvas.Canvas
	objects *[]Object
}

func NewObjectManager(canvas *canvas.Canvas, objects *[]Object) *ObjectManager {
	// TODO: Update ordering w/ z values
	sort.Slice(*objects, func(i, j int) bool {
		return (*objects)[i].GetPosition().Y() < (*objects)[j].GetPosition().Y()
	})

	objectManager := &ObjectManager{
		canvas:  canvas,
		objects: objects,
	}

	for _, obj := range *objects {
		obj.SetObjectManager(objectManager)
	}

	return objectManager
}

func (o ObjectManager) GetObjects() *[]Object {
	return o.objects
}

func (o ObjectManager) MoveObjects(diffTime time.Duration) {
	for _, object := range *o.objects {
		object.Move(diffTime)
	}
}

func (o ObjectManager) DrawObjects() {
	for _, object := range *o.objects {
		object.Draw()
	}
}

func (o ObjectManager) GetObjectsInArea(position vector.Vector, radius float64) []*Object {
	var returnSlice []*Object

	for _, object := range *o.objects {
		if Distance(object.GetPosition(), position) < radius {
			returnSlice = append(returnSlice, &object)
		}
	}
	return returnSlice
}

func (o ObjectManager) GetCanvas() *canvas.Canvas {
	return o.canvas
}
