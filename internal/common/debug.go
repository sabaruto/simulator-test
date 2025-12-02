package common

type DebugGetter interface {
	GetWindowState() bool
	GetRaysState() bool
	GetAgentMetricsState() bool
}

type DebugManager interface {
	DebugGetter
	SetWindowState(bool)
	SetRaysState(bool)
	SetAgentMetricsState(bool)
}
