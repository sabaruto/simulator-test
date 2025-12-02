package window

import (
	"fmt"
	"time"

	"github.com/sabaruto/simulator-test/internal/common"
	"github.com/sabaruto/simulator-test/internal/objects"
	"github.com/tfriedel6/canvas/sdlcanvas"
)

type AnimatedWindowManager struct {
	BackgroundColor string
	DebugManager    common.DebugManager
	Height          int
	Width           int
	frameRate       int
	objectManager   common.ObjectManager
	window          *sdlcanvas.Window
}

func NewAnimatedWindowManager(width int, height int) *AnimatedWindowManager {
	return &AnimatedWindowManager{
		Width:           width,
		Height:          height,
		BackgroundColor: "#000",
	}
}

func (s *AnimatedWindowManager) KeyUp(scancode int, rn rune, name string) {
	if name == "KeyD" {
		s.DebugManager.SetWindowState(!s.DebugManager.GetWindowState())
		fmt.Println("Debug Window value set:", s.DebugManager.GetRaysState())
	}

	if name == "KeyR" {
		s.DebugManager.SetRaysState(!s.DebugManager.GetRaysState())
		fmt.Println("Debug Rays value set:", s.DebugManager.GetRaysState())
	}

	if name == "KeyA" {
		s.DebugManager.SetAgentMetricsState(!s.DebugManager.GetAgentMetricsState())
		fmt.Println("Debug Agent Metrics value set:", s.DebugManager.GetAgentMetricsState())
	}
}

func (s *AnimatedWindowManager) Start(objectList *[]common.Object) {
	lastFrameTime := time.Now()
	lastDrawTime := time.Now()

	window, canvas, err := sdlcanvas.CreateWindow(s.Width, s.Height, "Animated Window Manager")
	if err != nil {
		panic(err)
	}

	window.KeyUp = s.KeyUp

	s.window = window
	s.objectManager = objects.NewObjectManager(canvas, objectList)

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
	cv := s.objectManager.GetCanvas()

	cv.SetFillStyle(s.BackgroundColor)
	cv.FillRect(0, 0, float64(s.Width), float64(s.Height))
}

func (s AnimatedWindowManager) drawDebugMetrics() {
	if !s.DebugManager.GetWindowState() {
		return
	}

	cv := s.objectManager.GetCanvas()
	cv.SetFillStyle("#FFF")
	cv.SetFont("assets/Righteous-Regular.ttf", 20)
	cv.FillText(fmt.Sprintf("FPS: %.0f", s.window.FPS()), 20, 30)
	cv.FillText(fmt.Sprintf("Draw Rate: %d", s.frameRate), 20, 50)
}

func (s *AnimatedWindowManager) SetFrameRate(frameRate int) {
	s.frameRate = frameRate
}
