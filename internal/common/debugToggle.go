package common

type DebugToggle struct {
	window         bool
	rays           bool
	agentMetrics   bool
	verboseLogging bool
}

func (d DebugToggle) Window() bool {
	return d.window
}

func (d DebugToggle) Rays() bool {
	return d.rays
}

func (d DebugToggle) AgentMetrics() bool {
	return d.agentMetrics
}

func (d DebugToggle) VerboseLogging() bool {
	return d.verboseLogging
}

func (d *DebugToggle) SetWindow(debugOn bool) {
	d.window = debugOn
}
func (d *DebugToggle) SetRays(debugOn bool) {
	d.rays = debugOn
}
func (d *DebugToggle) SetAgentMetrics(debugOn bool) {
	d.agentMetrics = debugOn
}

func (d *DebugToggle) SetVerboseLogging(debugOn bool) {
	d.verboseLogging = debugOn
}
