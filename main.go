package main

import (
	"github.com/sabaruto/simulator-test/internal/common"
	"github.com/sabaruto/simulator-test/internal/objects"
	windowmanager "github.com/sabaruto/simulator-test/internal/window-manager"
)

func main() {
	objs := &[]common.Object{
		objects.NewDot(
			common.Position{X: 500, Y: 500},
			50,
		),
		objects.NewDot(
			common.Position{X: 600, Y: 600},
			50,
		),
		objects.NewTower(
			common.Position{X: 400, Y: 400},
			0,
			100,
		),
	}

	wm := windowmanager.AnimatedWindowManager{
		Width:           1000,
		Height:          1000,
		BackgroundColor: "#242E24",
		Debug:           true,
	}

	wm.SetFrameRate(60)
	wm.Start(objs)
}
