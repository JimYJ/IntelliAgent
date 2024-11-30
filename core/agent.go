package core

type IAgent interface {
	Exec() error
}

type Agent struct {
	Name        string
	Role        string
	Description string
	Params      map[string]any
	Prompt      string
}

type AgentParams struct {
	Name        string
	ParamType   string
	Description string
}
