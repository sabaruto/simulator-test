package objects

import (
	"math"
	"time"

	"github.com/quartercastle/vector"
	"github.com/sabaruto/simulator-test/internal/common"
)

type Tower struct {
	Base
	width                      float64
	depth                      float64
	minHeight                  float64
	maxHeight                  float64
	currentOscillationDuration time.Duration
	oscillationDuration        time.Duration
	isMovingUp                 bool
	currentHeight              float64
}

func (t Tower) Draw() {
	cv := t.GetCanvas()

	common.NewRectBuilder().
		Canvas(cv).
		X(t.position.X()).
		Y(t.position.Y() - t.currentHeight).
		Width(t.width).
		Height(t.depth + t.currentHeight).
		FillColour("#363636ff").
		StrokeColour("#FFF").
		Build()

	common.NewRectBuilder().
		Canvas(cv).
		X(t.position.X()).
		Y(t.position.Y() - t.currentHeight).
		Width(t.width).
		Height(t.depth).
		FillColour("#575757ff").
		StrokeColour("#FFF").
		Build()
}

func (t *Tower) Move(duration time.Duration) {
	t.currentOscillationDuration += duration

	if t.currentOscillationDuration >= t.oscillationDuration {
		t.currentOscillationDuration -= t.oscillationDuration
		t.isMovingUp = !t.isMovingUp
	}

	heightDiff := t.maxHeight - t.minHeight
	oscilationPercentage := float64(t.currentOscillationDuration) / float64(t.oscillationDuration)

	newHeight := heightDiff/(1+math.Exp(-((oscilationPercentage-0.5)*7))) + t.minHeight

	// newHeight := heightDiff*float64(t.currentOscillationDuration)/float64(t.oscillationDuration) + t.minHeight

	if t.isMovingUp {
		t.currentHeight = newHeight
	} else {
		t.currentHeight = t.maxHeight - newHeight + t.minHeight
	}
}

type TowerBuilder struct {
	position            vector.Vector
	width               float64
	depth               float64
	minHeight           float64
	maxHeight           float64
	oscillationDuration time.Duration
	isMovingUp          bool
	startPercentage     float64
}

func NewTowerBuilder() *TowerBuilder {
	return &TowerBuilder{}
}

func (tb *TowerBuilder) Position(position vector.Vector) *TowerBuilder {
	tb.position = position
	return tb
}

func (tb *TowerBuilder) Depth(depth float64) *TowerBuilder {
	tb.depth = depth
	return tb
}

func (tb *TowerBuilder) Width(width float64) *TowerBuilder {
	tb.width = width
	return tb
}

func (tb *TowerBuilder) MinHeight(minHeight float64) *TowerBuilder {
	tb.minHeight = minHeight
	return tb
}
func (tb *TowerBuilder) MaxHeight(maxHeight float64) *TowerBuilder {
	tb.maxHeight = maxHeight
	return tb
}

func (tb *TowerBuilder) OscillationDuration(oscillationDuration time.Duration) *TowerBuilder {
	tb.oscillationDuration = oscillationDuration
	return tb
}

func (tb *TowerBuilder) StartPercentage(startPercentage float64) *TowerBuilder {
	if startPercentage > 1 {
		startPercentageInt := int(startPercentage)
		percentLeftOver := startPercentage - math.Floor(startPercentage)

		if startPercentageInt%2 == 0 {
			startPercentage = percentLeftOver
		} else {
			startPercentage = 1 - percentLeftOver
		}
	}

	tb.startPercentage = startPercentage
	return tb
}

func (tb *TowerBuilder) StartMovingDown() *TowerBuilder {
	tb.isMovingUp = false
	return tb
}

func (tb *TowerBuilder) StartMovingUp() *TowerBuilder {
	tb.isMovingUp = true
	return tb
}

func (tb *TowerBuilder) Build() *Tower {
	return &Tower{
		Base: Base{
			position: tb.position,
		},
		width:                      tb.width,
		depth:                      tb.depth,
		minHeight:                  tb.minHeight,
		maxHeight:                  tb.maxHeight,
		oscillationDuration:        tb.oscillationDuration,
		currentOscillationDuration: time.Duration(float64(tb.oscillationDuration) * tb.startPercentage),
		currentHeight:              (tb.maxHeight-tb.minHeight)*tb.startPercentage + tb.minHeight,
		isMovingUp:                 tb.isMovingUp,
	}
}
