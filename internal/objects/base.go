package objects

import (
	"github.com/sabaruto/simulator-test/internal/common"

	"github.com/tfriedel6/canvas"
)

type Base struct {
	position      common.Position
	objectManager *common.ObjectManager
}

func (b Base) GetPosition() common.Position {
	return b.position
}

func (b Base) GetCanvas() *canvas.Canvas {
	return b.objectManager.GetCanvas()
}

func (b *Base) SetObjectManager(objectManager *common.ObjectManager) {
	b.objectManager = objectManager
}
