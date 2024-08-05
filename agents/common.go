package agents

import (
	"bytes"
	"fmt"
	agentData "github.com/Ocyss/xiaomiHA/agents/data"
	"github.com/Ocyss/xiaomiHA/utils"
	log "github.com/sirupsen/logrus"
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
	"html/template"
	"time"
)

var (
	systemTmpl = initSystemTmpl()
)

func initSystemTmpl() *template.Template {
	tmpl, err := template.New("").Funcs(template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	}).Parse(fmt.Sprintf(`# Role：小爱健康分析专家-{{ .Name }}

## Profile：
- Language: 中文
- Description: %s{{ .Desc }}

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
1. 首先会获得用户{{ .TimeWindow }}内的健康数据，包括心率,睡眠,压力等和一些附加数据。
2. 分析这些数据，识别健康趋势和潜在问题，优先专注于当前职责。
{{ range $idx, $val := .Workflow -}}
{{ add $idx 3 }}. {{ $val }}
{{ end }}

## OutputFormat:
要求总是输出结构化 JSON 对象，符合下面 TypeScript 定义：
{{- .OutputFormat }}

## Initialization
作为主人的健康助理，我将遵循上述约束条件，并严格的要求主人完成各项目标，分析健康数据，提供一些个性化建议。
`, utils.C.Character))
	if err != nil {
		log.Panicln(err)
	}

	return tmpl
}

type SystemAgentOptions struct {
	Name         string
	Desc         string
	Skills       []string
	Goals        []string
	Constrains   []string
	Workflow     []string
	TimeWindow   string
	OutputFormat string
}

type Agent interface {
	SystemAgentOptions() *SystemAgentOptions
}
type AgentCsv interface {
	Agent
	GenCsvFile() ([]*agentData.CsvData, error)
}

func SystemAgent(agent Agent) (string, error) {
	var buf bytes.Buffer
	opt := agent.SystemAgentOptions()
	if opt.OutputFormat == "" {
		ts := typescriptify.New()
		ts.DontExport = true
		ts.CreateInterface = true
		outputFormat, err := ts.
			Add(agent).
			Convert(map[string]string{})
		if err != nil {
			return "", err
		}
		opt.OutputFormat = outputFormat
	}

	if err := systemTmpl.Execute(&buf, opt); err != nil {
		return "", err
	}
	return buf.String(), nil
}
func SystemInfo() string {
	now := time.Now()
	return fmt.Sprintf("系统信息:\n当前时间:%s\n星期:%s\n", now.Format("2006-01-02 15:04:05"), now.Weekday().String())
}
