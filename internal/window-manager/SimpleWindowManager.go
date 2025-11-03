package windowmanager

import (
	"github.com/sabaruto/simulator-test/internal/common"
	"github.com/tfriedel6/canvas/sdlcanvas"
)

type SimpleWindowManager struct {
	window          *sdlcanvas.Window
	objectManager   *common.ObjectManager
	Width           int
	Height          int
	BackgroundColor string
}

func NewSimpleWindowManager(width int, height int) *SimpleWindowManager {
	return &SimpleWindowManager{
		Width:           width,
		Height:          height,
		BackgroundColor: "#000",
	}
}

func (s *SimpleWindowManager) Start(objects *[]common.Object) {
	window, canvas, err := sdlcanvas.CreateWindow(s.Width, s.Height, "Simple Window Manager")
	if err != nil {
		panic(err)
	}

	s.window = window
	s.objectManager = common.NewObjectManager(canvas, objects)

	defer window.Destroy()

	window.MainLoop(func() {
		canvas.SetFillStyle(s.BackgroundColor)
		canvas.FillRect(0, 0, float64(s.Width), float64(s.Height))

		s.objectManager.DrawObjects()

	})
}

func (s SimpleWindowManager) SetFrameRate(frameRate int) {
}
