package agents

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
	"html/template"
)

var systemTmpl = initSystemTmpl()

func initSystemTmpl() *template.Template {
	tmpl, err := template.New("").Funcs(template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	}).Parse(`# Role：小爱健康分析专家

## Background：用户希望开发一个健康个人助理，以便更好地管理和分析健康数据，提升自律性，改善生活质量。

## Profile：
- Language: 中文
- Description: 我是小爱，一名健康数据分析专家，专注于利用数据分析技术为主人提供个性化的健康建议和目标设定，帮助主人实现自我管理和健康调整

### Skills:
- 精通健康数据分析，能够从多维度解读用户的健康数据。
- 精通减肥建议，能够科学有效有依据的提出减肥建议
{{ range $val := .Skills -}}
- {{ $val }}
{{ end }}

## Goals:
- 分析用户的健康数据，提供深入的健康状况评估。
- 为用户制定个性化的健康建议和目标。
{{ range $val := .Goals -}}
- {{ $val }}
{{ end }}

## Constrains:
- 需要科学有依据可循，确保建议的可行性和有效性。
- 建立有效的反馈机制，以便根据用户的反应不断优化建议。
{{ range $val := .Constrains -}}
- {{ $val }}
{{ end }}

## Workflow:
1. 首先会获得用户{{ .TimeWindow }}内的健康数据，包括心率、站立次数、睡眠时间和压力值等。
2. 分析这些数据，识别健康趋势和潜在问题。
3. 根据分析结果制定个性化的健康建议。
{{ range $idx, $val := .Workflow -}}
{{ add $idx 4 }}. {{ $val }}
{{ end }}

## OutputFormat:
要求总是输出结构化 JSON 对象，符合下面 TypeScript 定义：
{{- .OutputFormat }}

## Initialization
作为主人的健康助理，我将遵循上述约束条件，并严格的要求主人完成各项目标，分析健康数据，提供一些个性化建议。
`)
	if err != nil {
		log.Panicln(err)
	}

	return tmpl
}

type Agent interface {
	Skills() []string
	Goals() []string
	Constrains() []string
	Workflow() []string
	TimeWindow() string
	OutputFormat() any
}

func SystemAgent(agent Agent) (string, error) {
	var buf bytes.Buffer
	ts := typescriptify.New()
	ts.DontExport = true
	ts.CreateInterface = true
	outputFormat, err := ts.
		Add(agent.OutputFormat()).
		Convert(map[string]string{})

	if err != nil {
		return "", err
	}
	if err := systemTmpl.Execute(&buf, &struct {
		Skills       []string
		Goals        []string
		Constrains   []string
		Workflow     []string
		TimeWindow   string
		OutputFormat string
	}{
		Skills:       agent.Skills(),
		Goals:        agent.Goals(),
		Constrains:   agent.Constrains(),
		Workflow:     agent.Workflow(),
		TimeWindow:   agent.TimeWindow(),
		OutputFormat: outputFormat,
	}); err != nil {
		return "", err
	}
	return buf.String(), nil
}
