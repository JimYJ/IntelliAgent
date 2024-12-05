package core

type WorkerInterface interface {
	Exec() error
}

type Worker struct {
	Name            string
	Goal            string
	Role            string
	RoleDescription string
	Description     string
	Params          map[string]any
	Prompt          string
}

type WorkerParams struct {
	Name        string
	ParamType   string
	Description string
}
