package windowmanager

import (
	"fmt"
	"time"

	"github.com/sabaruto/simulator-test/internal/common"
	"github.com/tfriedel6/canvas/sdlcanvas"
)

type AnimatedWindowManager struct {
	BackgroundColor string
	Debug           bool
	Height          int
	Width           int
	frameRate       int
	objectManager   *common.ObjectManager
	window          *sdlcanvas.Window
}

func NewAnimatedWindowManager(width int, height int) *AnimatedWindowManager {
	return &AnimatedWindowManager{
		Width:           width,
		Height:          height,
		BackgroundColor: "#000",
	}
}

func (s *AnimatedWindowManager) Start(objects *[]common.Object) {
	lastFrameTime := time.Now()
	lastDrawTime := time.Now()

	window, canvas, err := sdlcanvas.CreateWindow(s.Width, s.Height, "Animated Window Manager")
	if err != nil {
		panic(err)
	}

	s.window = window
	s.objectManager = common.NewObjectManager(canvas, objects)

	defer window.Destroy()

	window.MainLoop(func() {
		frameTimeDiff := time.Since(lastFrameTime)
		drawTimeDiff := time.Since(lastDrawTime)

		s.objectManager.MoveObjects(frameTimeDiff)
		s.drawFrame(drawTimeDiff, &lastDrawTime)

		lastFrameTime = time.Now()
	})
}

func (s *AnimatedWindowManager) drawFrame(drawTimeDiff time.Duration, lastDrawTime *time.Time) {
	if drawTimeDiff < s.drawDuration() {
		return
	}

	s.drawBackground()
	s.objectManager.DrawObjects()
	s.drawDebugMetrics()

	*lastDrawTime = time.Now()
}

func (s AnimatedWindowManager) drawDuration() time.Duration {
	return time.Duration(int(time.Second) / s.frameRate)
}

func (s AnimatedWindowManager) drawBackground() {
	canvas := s.objectManager.GetCanvas()

	canvas.SetFillStyle(s.BackgroundColor)
	canvas.FillRect(0, 0, float64(s.Width), float64(s.Height))
}

func (s AnimatedWindowManager) drawDebugMetrics() {
	if !s.Debug {
		return
	}

	canvas := s.objectManager.GetCanvas()
	canvas.SetFillStyle("#FFF")
	canvas.SetFont("Righteous-Regular.ttf", 20)
	canvas.FillText(fmt.Sprintf("FPS: %.0f", s.window.FPS()), 20, 30)
	canvas.FillText(fmt.Sprintf("Draw Rate: %d", s.frameRate), 20, 50)
}

func (s *AnimatedWindowManager) SetFrameRate(frameRate int) {
	s.frameRate = frameRate
}
