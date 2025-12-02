package objects

import (
	"math"
	"sort"
	"time"

	"github.com/quartercastle/vector"
	"github.com/sabaruto/simulator-test/internal/common"
	"github.com/tfriedel6/canvas"
)

type objectManager struct {
	canvas  *canvas.Canvas
	objects *[]common.Object
}

func NewObjectManager(canvas *canvas.Canvas, objects *[]common.Object) common.ObjectManager {
	// TODO: Update ordering w/ z values
	sort.Slice(*objects, func(i, j int) bool {
		return (*objects)[i].GetPosition().Y() < (*objects)[j].GetPosition().Y()
	})

	objectManager := &objectManager{
		canvas:  canvas,
		objects: objects,
	}

	for _, obj := range *objects {
		obj.SetObjectManager(objectManager)
	}

	return objectManager
}

func (o objectManager) GetObjects() *[]common.Object {
	return o.objects
}

func (o objectManager) MoveObjects(diffTime time.Duration) {
	for _, object := range *o.objects {
		object.Move(diffTime)
	}
}

func (o objectManager) DrawObjects() {
	for _, object := range *o.objects {
		object.Draw()
	}
}

func (o objectManager) GetObjectsInArea(position vector.Vector, radius float64) []*common.Object {
	var returnSlice []*common.Object

	for _, object := range *o.objects {
		if common.Distance(object.GetPosition(), position) < radius {
			returnSlice = append(returnSlice, &object)
		}
	}
	return returnSlice
}

// TODO: Find a method to find a closest object via a ray of some kind
func (o objectManager) SendRay(rayPosition vector.Vector, rayDirection vector.Vector, ignoreIDs []int64) (*vector.Vector, *string) {
	smallestDistance := math.Inf(1)
	var intersectPoint *vector.Vector
	var intersectColour *string

	for _, object := range *o.objects {
		skipObject := false
		for _, id := range ignoreIDs {
			if object.GetID() == id {
				skipObject = true
				continue
			}
		}

		if skipObject {
			continue
		}

		currIntersectPoint, currColour := object.ReceiveRay(rayPosition, rayDirection)

		if currIntersectPoint == nil {
			continue
		}

		intersectDistance := currIntersectPoint.Sub(rayPosition).Magnitude()
		if intersectDistance < smallestDistance {
			intersectPoint = currIntersectPoint
			intersectColour = currColour
			smallestDistance = intersectDistance
		}
	}

	if intersectPoint == nil {
		return nil, nil
	}

	return intersectPoint, intersectColour
}

func (o objectManager) GetCanvas() *canvas.Canvas {
	return o.canvas
}
