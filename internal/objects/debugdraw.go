package objects

import (
	"time"

	"github.com/quartercastle/vector"
)

type DebugDraw struct {
	BaseObject
}

func (d DebugDraw) Draw() {
	d.drawDebugRays()
}

func (d DebugDraw) drawDebugRays() {
	cv := d.GetCanvas()
	cv.BeginPath()
	cv.SetStrokeStyle("#FF00FF")
	for y := 10.0; y <= 990; y += 10 {
		rayStopPosition, _ := d.objectManager.SendRay(vector.Vector{10, y}, vector.Vector{1, 0}, nil)
		cv.MoveTo(10, y)

		if rayStopPosition != nil {
			cv.LineTo(rayStopPosition.X(), rayStopPosition.Y())
		} else {
			cv.LineTo(990, y)
		}
	}
	cv.Stroke()
}

func (d DebugDraw) Move(diffTime time.Duration) {
}
