package main

import (
	"flag"

	"github.com/sabaruto/simulator-test/internal/common"
	"github.com/sabaruto/simulator-test/internal/debug"
	"github.com/sabaruto/simulator-test/internal/objects"
	"github.com/sabaruto/simulator-test/internal/window"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	flag.Parse()

	debug.ManageMemProfiler(*cpuprofile)
	debug.ManageCpuProfiler(*memprofile)

	debugManager := debug.NewDebugManager()

	// objs := animations.CreateTowerAnimation()
	objs := &[]common.Object{
		objects.NewDotAgentBuilder().
			Position(200, 400).
			Radius(30).
			Colour("#808080").
			DebugClient(debugManager).
			Build(),
		objects.NewDotAgentBuilder().
			Position(600, 400).
			Radius(30).
			Colour("#808080").
			DebugClient(debugManager).
			Build(),
		objects.NewDotAgentBuilder().
			Position(400, 400).
			Radius(50).
			Colour("#008080").
			DebugClient(debugManager).
			Build(),
	}

	wm := window.AnimatedWindowManager{
		Width:           1000,
		Height:          1000,
		BackgroundColor: "#242E24",
		DebugManager:    debugManager,
	}

	wm.SetFrameRate(120)
	wm.Start(objs)
}
