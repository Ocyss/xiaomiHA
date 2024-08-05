package agentTools

import (
	"testing"
)

func TestFsScheduleTool(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		wantErr bool
	}{
		{"1", `{
							"summary":"测试日程",
							"description":"测试<br />用例1<br /><a href=\"https://baidu.com\">baidu<a/>",
							"start_time":"2024/08/04 23:50:00",
							"end_time":"2024/08/05 00:00:00",
							"reminders":5
						  }`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FsScheduleTool(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("FsScheduleTool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}

func TestWeatherTool(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		wantErr bool
	}{
		{"1", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WeatherTool(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("WeatherTool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}
