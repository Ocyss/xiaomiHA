package agentTools

import (
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

var (
	ToolsCall = map[string]*T{
		WatherToolDesc.Name:     WatherToolDesc,
		FsScheduleToolDesc.Name: FsScheduleToolDesc,
	}
	Tools = []openai.Tool{WatherToolDesc.Tool(), FsScheduleToolDesc.Tool()}
)

type T struct {
	Name        string
	Description string
	Parameters  jsonschema.Definition
	Func        func(string) (string, error)
}

func (t *T) Tool() openai.Tool {
	return openai.Tool{
		Type: openai.ToolTypeFunction,
		Function: &openai.FunctionDefinition{
			Name:        t.Name,
			Description: t.Description,
			Parameters:  t.Parameters,
		},
	}
}
