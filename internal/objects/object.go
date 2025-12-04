package objects

import (
	"fmt"
	"math/rand"

	"github.com/quartercastle/vector"
	"github.com/sabaruto/simulator-test/internal/common"

	"github.com/tfriedel6/canvas"
)

type object struct {
	position      *vector.Vector
	objectManager common.ObjectManager
	debugGetter   common.DebugGetter
	id            *int64
}

func (o object) GetPosition() vector.Vector {
	return *o.position
}

func (o object) GetCanvas() *canvas.Canvas {
	return o.objectManager.GetCanvas()
}

func (o *object) SetObjectManager(objectManager common.ObjectManager) {
	o.objectManager = objectManager
}

func (o object) ReceiveRay(rayPosition vector.Vector, rayDirection vector.Vector) (*vector.Vector, *string) {
	return nil, nil
}

func (o *object) GetID() int64 {
	if o.id == nil {
		o.id = createID()
	}

	return *o.id
}

func createID() *int64 {
	newId := rand.Int63()
	return &newId
}

func (o object) Equal(other common.Object) bool {
	return o.GetID() == other.GetID()
}

func (o object) String() string {
	return fmt.Sprintf("Base(id: %d, position: %v)", o.GetID(), o.GetPosition())
}
