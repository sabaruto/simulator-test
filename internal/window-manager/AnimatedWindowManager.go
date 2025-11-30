package windowmanager

import (
	"fmt"
	"time"

	"github.com/sabaruto/simulator-test/internal/common"
	"github.com/tfriedel6/canvas/sdlcanvas"
)

type AnimatedWindowManager struct {
	BackgroundColor string
	DebugToggles    *common.DebugToggle
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

func (s *AnimatedWindowManager) KeyUp(scancode int, rn rune, name string) {
	if s.DebugToggles.VerboseLogging() {
		fmt.Println("Key Up event")
	}

	if name == "KeyD" {
		s.DebugToggles.SetWindow(!s.DebugToggles.Window())
		fmt.Println("Debug Window value set:", s.DebugToggles.Window())
	}

	if name == "KeyR" {
		s.DebugToggles.SetRays(!s.DebugToggles.Rays())
		fmt.Println("Debug Rays value set:", s.DebugToggles.Rays())
	}

	if name == "KeyA" {
		s.DebugToggles.SetAgentMetrics(!s.DebugToggles.AgentMetrics())
		fmt.Println("Debug Agent Metrics value set:", s.DebugToggles.AgentMetrics())
	}

	if name == "KeyV" {
		s.DebugToggles.SetVerboseLogging(!s.DebugToggles.VerboseLogging())
		fmt.Println("Debug Verbose Logging value set:", s.DebugToggles.VerboseLogging())
	}
}

func (s *AnimatedWindowManager) Start(objects *[]common.Object) {
	lastFrameTime := time.Now()
	lastDrawTime := time.Now()

	window, canvas, err := sdlcanvas.CreateWindow(s.Width, s.Height, "Animated Window Manager")
	if err != nil {
		panic(err)
	}

	window.KeyUp = s.KeyUp

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
	cv := s.objectManager.GetCanvas()

	cv.SetFillStyle(s.BackgroundColor)
	cv.FillRect(0, 0, float64(s.Width), float64(s.Height))
}

func (s AnimatedWindowManager) drawDebugMetrics() {
	if !s.DebugToggles.Window() {
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
