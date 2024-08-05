package agentTools

import (
	"context"
	"errors"
	"fmt"
	"github.com/Ocyss/xiaomiHA/utils"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcalendar "github.com/larksuite/oapi-sdk-go/v3/service/calendar/v4"
	"github.com/sashabaranov/go-openai/jsonschema"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"strconv"
	"time"
)

var (
	fsConf   = utils.C.FeiShu
	fsClient = lark.NewClient(fsConf.AppId, fsConf.AppSecret)
)

func FsScheduleTool(args string) (string, error) {
	log.Debugln("FsScheduleTool", args)
	datas := gjson.Get(args, "summarys").Array()
	for _, data := range datas {
		summary := data.Get("summary").String()
		description := data.Get("description").String()
		_startTime := data.Get("start_time").String()
		startTime, err1 := time.ParseInLocation("2006/01/02 15:04:05", _startTime, time.Local)
		_endTime := data.Get("end_time").String()
		endTime, err2 := time.ParseInLocation("2006/01/02 15:04:05", _endTime, time.Local)
		if err := errors.Join(err1, err2); err != nil {
			return "", fmt.Errorf("时间格式错误: %v", err)
		}
		reminders := data.Get("reminders").Int()
		req := larkcalendar.NewCreateCalendarEventReqBuilder().
			CalendarId(fsConf.CalendarId).
			CalendarEvent(larkcalendar.NewCalendarEventBuilder().
				Summary(summary).
				Description(description).
				NeedNotification(false).
				StartTime(larkcalendar.NewTimeInfoBuilder().
					Timestamp(strconv.FormatInt(startTime.Unix(), 10)).
					Build()).
				EndTime(larkcalendar.NewTimeInfoBuilder().
					Timestamp(strconv.FormatInt(endTime.Unix(), 10)).
					Build()).
				Visibility(`private`).
				Reminders([]*larkcalendar.Reminder{
					larkcalendar.NewReminderBuilder().
						Minutes(int(reminders)).
						Build(),
				}).
				Build()).
			Build()

		// 发起请求
		resp, err := fsClient.Calendar.CalendarEvent.Create(context.Background(), req)

		// 处理错误
		if err != nil {

			return "", fmt.Errorf("发起请求失败: %v", err)
		}

		// 服务端错误处理
		if !resp.Success() {
			return "", fmt.Errorf("服务端错误: [%d] %s", resp.Code, resp.Msg)
		}
	}

	return "日程创建成功", nil
}

var FsScheduleToolDesc = &T{
	Name:        "FsSchedule",
	Description: "需要进行活动/日程等安排的时候，将在日历中创建日程,可以创建多个(不应该重叠)，一般情况下不设定特别久远的日程最好3天内",
	Parameters: jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"summarys": {
				Description: "日程数据",
				Type:        jsonschema.Array,
				Items: &jsonschema.Definition{
					Type: jsonschema.Object,
					Properties: map[string]jsonschema.Definition{
						"summary": {
							Type:        jsonschema.String,
							Description: "日程标题",
						},
						"description": {
							Type:        jsonschema.String,
							Description: "日程描述。支持解析Html标签。",
						},
						"start_time": {
							Type:        jsonschema.String,
							Description: "开始时间，如：2024/12/10 20:00:00",
						},
						"end_time": {
							Type:        jsonschema.String,
							Description: "结束时间，如：2024/12/10 21:00:00",
						},
						"reminders": {
							Type:        jsonschema.Integer,
							Description: "日程提醒时间，分钟单位，正数开始前，负数开始后，如 5 开始前5分钟提醒，-5 开始后5分钟提醒",
						},
					},
					Required: []string{"summary", "start_time", "end_time"},
				},
			},
		},
	},
	Func: FsScheduleTool,
}
