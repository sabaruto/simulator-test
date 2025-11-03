package objects

import (
	"time"

	"github.com/sabaruto/simulator-test/internal/common"
)

type Tower struct {
	Base
	minHeight           float64
	maxHeight           float64
	startTime           time.Time
	oscillationDuration time.Duration
}

func NewTower(position common.Position, minHeight float64, maxHeight float64) *Tower {
	return &Tower{
		Base: Base{
			position: position,
		},
		minHeight:           minHeight,
		maxHeight:           maxHeight,
		startTime:           time.Now(),
		oscillationDuration: 5 * time.Second,
	}
}

func (t Tower) Draw() {
	cv := t.GetCanvas()
	cv.SetFillStyle(128, 128, 128)
	cv.BeginPath()
	cv.Rect(t.position.X, t.position.Y, 50, 40)
	cv.Fill()
}

func (t Tower) Move(duration time.Duration) {
}
