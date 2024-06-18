package engine

type Engine struct {
}

func New() *Engine {
	return &Engine{}
}

func (e *Engine) Render(template string, variables map[string]interface{}) (string, error) {
	return "", nil
}
