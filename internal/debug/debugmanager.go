package debug

import "github.com/sabaruto/simulator-test/internal/common"

func NewDebugManager() common.DebugManager {
	return &debugManager{}
}

type debugManager struct {
	windowState       bool
	raysState         bool
	agentMetricsState bool
}

func (d debugManager) GetWindowState() bool {
	return d.windowState
}
func (d debugManager) GetRaysState() bool {
	return d.raysState
}
func (d debugManager) GetAgentMetricsState() bool {
	return d.agentMetricsState
}

func (d *debugManager) SetWindowState(windowState bool) {
	d.windowState = windowState
}

func (d *debugManager) SetRaysState(raysState bool) {
	d.raysState = raysState
}

func (d *debugManager) SetAgentMetricsState(agentMetricsState bool) {
	d.agentMetricsState = agentMetricsState
}
