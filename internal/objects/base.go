package objects

import (
	"fmt"
	"math/rand"

	"github.com/quartercastle/vector"
	"github.com/sabaruto/simulator-test/internal/common"

	"github.com/tfriedel6/canvas"
)

type Base struct {
	position      vector.Vector
	objectManager *common.ObjectManager
	debug         bool
	id            *int64
}

func (b Base) GetPosition() vector.Vector {
	return b.position
}

func (b Base) GetCanvas() *canvas.Canvas {
	return b.objectManager.GetCanvas()
}

func (b *Base) SetObjectManager(objectManager *common.ObjectManager) {
	b.objectManager = objectManager
}

func (b Base) ReceiveRay(rayPosition vector.Vector, rayDirection vector.Vector) (intersectPoint *vector.Vector, colour *string) {
	return nil, nil
}

func (b *Base) GetID() int64 {
	if b.id == nil {
		newId := rand.Int63()

		b.id = &newId
	}

	return *b.id
}

func (b Base) Equal(other common.Object) bool {
	return b.GetID() == other.GetID()
}

func (b Base) String() string {
	return fmt.Sprintf("Base(id: %d, position: %v)", b.GetID(), b.GetPosition())
}
