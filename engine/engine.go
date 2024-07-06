package engine

import (
	"fmt"
	"strings"

	"github.com/brecht-vde/prompter/engine/internal"
)

type Engine struct {
}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) Render(template string, variables map[string]interface{}) (string, error) {
	l := internal.NewLexer(template)
	p := internal.NewParser(l)
	t := p.Parse()

	errs := p.Errors()

	if len(errs) > 0 {
		return "", fmt.Errorf("errors occurred while parsing the template: \n%s", strings.Join(errs, "\n"))
	}

	rendition, err := internal.Eval(*t, variables)

	if err != nil {
		return "", err
	}

	return rendition, nil
}
