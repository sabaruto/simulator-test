package animations

import (
	"time"

	"github.com/sabaruto/simulator-test/internal/common"
	"github.com/sabaruto/simulator-test/internal/objects"
)

func CreateTowerAnimation() *[]common.Object {
	width := 60
	depth := 30
	startX := 250
	startY := 150
	rows := 15
	columns := 16
	minHeight := 10
	maxHeight := 100
	duration := 350 * time.Millisecond
	percentageIncrease := 2 / float64(columns)

	towerBuilder := objects.NewTowerBuilder().
		Width(float64(width)).
		Depth(float64(depth)).
		MinHeight(float64(minHeight)).
		MaxHeight(float64(maxHeight)).
		StartMovingUp().
		OscillationDuration(duration)

	objs := &[]common.Object{}

	startPercentage := float64(0)
	for y := startY; y < startY+depth*rows; y += depth {
		for x := startX; x < startX+width*columns; x += width {
			*objs = append(
				*objs,
				towerBuilder.
					Position(common.Position{X: float64(x), Y: float64(y)}).
					StartPercentage(startPercentage).
					Build(),
			)
			startPercentage += percentageIncrease
		}
		startPercentage += percentageIncrease
	}

	return objs
}
