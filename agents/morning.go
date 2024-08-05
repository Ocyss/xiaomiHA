package agents

import (
	"context"
	"fmt"
	agentData "github.com/Ocyss/xiaomiHA/agents/data"
	agentTools "github.com/Ocyss/xiaomiHA/agents/tools"
	"github.com/Ocyss/xiaomiHA/internal/feishu"
	"github.com/Ocyss/xiaomiHA/internal/task"
	"github.com/Ocyss/xiaomiHA/utils"
	"github.com/hibiken/asynq"
	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
	"github.com/sashabaranov/go-openai"
	"time"
)

type MorningAgent struct {
	Title   string `json:"table" ts_doc:"新的一天，一个简短的标题, 可以是一句诗,歌词,英语,名言等"`
	Weather string `json:"weather" ts_doc:"关于天气的注意事项, 要和日程出行挂钩,支持md语法"`
	Sleep   string `json:"sleep" ts_doc:"睡眠情况评价, 需要详细，并提出改善意见, 因为是早安助手，需要重点关注晚上的睡眠，还应该给出午休时长等建议"`
}

func (d *MorningAgent) SystemAgentOptions() *SystemAgentOptions {
	return &SystemAgentOptions{
		Name: "早安助手",
		Desc: "早晨能获取一天中需要的信息和主人的健康情况，生成对应的关心话术，贴心提示等",
		Workflow: []string{
			"理解输入的睡眠数据，对主人进行提醒",
			"调用wather tool获取近24小时的天气",
			"根据主人的日程安排来进行一些天气注意事项, 需要重点关注上下班时间",
		},
	}
}

func (d *MorningAgent) ProcessTask(ctx context.Context, t *asynq.Task) error {
	prompt, csvFile, err := getPromptAndCsvFile(d)
	if err != nil {
		return err
	}
	messages, err := createChatMessages[MorningAgent](ctx, &ChatOption{
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleSystem, Content: prompt,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: fmt.Sprintf("%s用户信息: %s\n%s\n%s", SystemInfo(), utils.C.User.Base, utils.C.User.Other["morning"], csvFile),
			},
		},
		ToolChoice: openai.ToolChoice{
			Type:     openai.ToolTypeFunction,
			Function: openai.ToolFunction{Name: agentTools.WatherToolDesc.Name},
		},
	})
	if err != nil {
		return err
	}

	err = feishu.SendMessage(ctx, feishu.MsgOpt{
		Card: larkcard.NewMessageCard().
			Header(larkcard.NewMessageCardHeader().Title(larkcard.NewMessageCardPlainText().Content(messages.Title))).
			Elements([]larkcard.MessageCardElement{
				larkcard.NewMessageCardMarkdown().Content(messages.Weather),
				larkcard.NewMessageCardHr(),
				larkcard.NewMessageCardMarkdown().Content(messages.Sleep),
				larkcard.NewMessageCardNote().Elements([]larkcard.MessageCardNoteElement{
					larkcard.NewMessageCardPlainText().Content(time.Now().Format("2006-01-02 15:04:05")),
				}),
			}).
			Build(),
	})
	return err
}

func (d *MorningAgent) GenCsvFile() ([]*agentData.CsvData, error) {
	limit7 := agentData.NewLimitComm(7)
	t := utils.CreateTime()
	time3 := agentData.NewTimeComm(t.Day(2))
	return genCsvFile(
		[]agentData.GetCsvRecord{
			&agentData.SleepSegment{Comm: agentData.NewLimitComm(2)},
			&agentData.Vitality{Comm: agentData.NewLimitComm(3)},
			&agentData.HeartRate{Comm: limit7},
			&agentData.Spo2{Comm: limit7},
			&agentData.Stress{Comm: limit7},
		},
		[]agentData.GetCsvRecord{
			&agentData.HeartRate{Comm: time3},
			&agentData.Spo2{Comm: time3},
			&agentData.Stress{Comm: time3},
		},
	)
}

func init() {
	task.T.AddHandler(&asynq.PeriodicTaskConfig{
		Cronspec: "20 8 * * *",
		Task:     asynq.NewTask("morning", nil),
	}, &MorningAgent{})
}
