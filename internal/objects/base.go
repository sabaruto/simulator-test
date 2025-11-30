package objects

import (
	"fmt"
	"math/rand"

	"github.com/quartercastle/vector"
	"github.com/sabaruto/simulator-test/internal/common"

	"github.com/tfriedel6/canvas"
)

type BaseObject struct {
	position      vector.Vector
	objectManager *common.ObjectManager
	debugToggles  *common.DebugToggle
	id            *int64
}

func (b BaseObject) GetPosition() vector.Vector {
	return b.position
}

func (b BaseObject) GetCanvas() *canvas.Canvas {
	return b.objectManager.GetCanvas()
}

func (b *BaseObject) SetObjectManager(objectManager *common.ObjectManager) {
	b.objectManager = objectManager
}

func (b BaseObject) ReceiveRay(rayPosition vector.Vector, rayDirection vector.Vector) (intersectPoint *vector.Vector, colour *string) {
	return nil, nil
}

func (b *BaseObject) GetID() int64 {
	if b.id == nil {
		newId := rand.Int63()

		b.id = &newId
	}

	return *b.id
}

func (b BaseObject) Equal(other common.Object) bool {
	return b.GetID() == other.GetID()
}

func (b BaseObject) String() string {
	return fmt.Sprintf("Base(id: %d, position: %v)", b.GetID(), b.GetPosition())
}
