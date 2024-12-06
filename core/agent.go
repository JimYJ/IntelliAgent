package core

var (
	_ Workers = &Worker{}
	_ Tools   = &Tool{}
)

type Workers interface {
	Exec() error
}

type Tools interface {
	Use() (string, error)
	GetParams() []Param
}

type Worker struct {
	Name             string    // 任务名称
	Goal             string    // 任务目标
	Role             string    // 任务角色
	RoleDescription  string    // 任务角色描述
	Description      string    // 任务描述
	Prompt           string    // 任务提示词
	ResponseTemplate string    // 响应模版
	APIBody          OpenAIAPI //
	Tool             []Tools   // 任务工具
}

func (worker *Worker) Exec() error {
	if worker.Tool == nil {

	}
	return nil
}

type OpenAIAPI struct {
	Model  string
	Prompt string
}

type Tool struct {
	Name              string                              // 工具名称
	Description       string                              // 工具描述
	Run               func(params ...any) (string, error) // 工具执行函数
	Params            []any                               // 工具参数
	ParamsDescription []Param                             // 工具参数说明
}

type Param struct {
	Name        string // 参数名称
	ParamType   string // 参数类型
	Description string // 参数描述
	Required    bool   // 是否必填
}

func (tool *Tool) Use() (string, error) {
	return tool.Run(tool.Params...)
}

func (tool *Tool) GetParams() []Param {
	return tool.ParamsDescription
}
