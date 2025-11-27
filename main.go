package main

import (
	"flag"

	"github.com/sabaruto/simulator-test/internal/common"
	"github.com/sabaruto/simulator-test/internal/objects"
	windowmanager "github.com/sabaruto/simulator-test/internal/window-manager"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	flag.Parse()

	manageCpuProfiler(*cpuprofile)
	manageMemProfiler(*memprofile)

	// objs := animations.CreateTowerAnimation()
	objs := &[]common.Object{
		objects.NewDotAgentBuilder().
			Position(400, 400).
			Radius(30).
			Colour("#808080").
			Build(),
	}

	wm := windowmanager.AnimatedWindowManager{
		Width:           1000,
		Height:          1000,
		BackgroundColor: "#242E24",
		Debug:           true,
	}

	wm.SetFrameRate(120)
	wm.Start(objs)
}
