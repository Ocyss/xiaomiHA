package agentTools

import (
	"fmt"
	"github.com/Ocyss/xiaomiHA/utils"
	"github.com/go-resty/resty/v2"
	"github.com/sashabaranov/go-openai/jsonschema"
	"github.com/tidwall/gjson"
)

func WeatherTool(args string) (string, error) {
	c := utils.C.Gpt.Tools.Weather
	res, err := resty.New().R().
		SetQueryParam("location", c.Location).
		SetQueryParam("key", c.Key).
		Get("https://devapi.qweather.com/v7/weather/24h")
	if err != nil {
		return "", fmt.Errorf("获取天气失败: %v", err)
	}
	return gjson.GetBytes(res.Body(), "hourly").String(), nil
}

var WatherToolDesc = &T{
	Name:        "Weather",
	Description: "当需要安排户外活动，或者用户询问等需要使用实时天气时将获取24小时天气",
	Parameters: jsonschema.Definition{
		Type:       jsonschema.Object,
		Properties: map[string]jsonschema.Definition{},
	},
	Func: WeatherTool,
}
