package feishu

import (
	"context"
	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
	"testing"
)

func TestSendMessage(t *testing.T) {
	type args struct {
		ctx context.Context
		opt MsgOpt
	}
	t2, _ := larkcard.NewMessageCard().
		Config(&larkcard.MessageCardConfig{}).
		Header(&larkcard.MessageCardHeader{}).
		Elements([]larkcard.MessageCardElement{
			larkcard.NewMessageCardMarkdown().Content("notation字号\n标准emoji 😁😢🌞💼🏆❌✅\n*斜体*\n**粗体**\n~~删除线~~\n[差异化跳转]($urlVal)\n<at id=all></at>"),
		}).
		JSON()
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"1", args{
			context.Background(), MsgOpt{
				Content: `{"config":{},"i18n_elements":{"zh_cn":[{"tag":"markdown","content":"开发者运营项目启动会","text_align":"left","text_size":"normal","icon":{"tag":"standard_icon","token":"submit-feedback_outlined","color":"grey"}},{"tag":"markdown","content":"2024年1月25日（周四） 14:30 - 15:00 (GMT+8)","text_align":"left","text_size":"normal","icon":{"tag":"standard_icon","token":"time_outlined","color":"grey"}},{"tag":"person_list","persons":[{"id":"39afa217"},{"id":"78af3bbf"},{"id":"29ff53b6"}],"size":"small","lines":1,"show_avatar":true,"show_name":true,"icon":{"tag":"standard_icon","token":"group_outlined","color":"grey"}},{"tag":"action","layout":"default","actions":[{"tag":"button","text":{"tag":"plain_text","content":"编辑日程信息"},"type":"default","complex_interaction":true,"width":"default","size":"medium"},{"tag":"button","text":{"tag":"plain_text","content":"确认创建日程"},"type":"primary","complex_interaction":true,"width":"default","size":"medium"}]}]},"i18n_header":{"zh_cn":{"title":{"tag":"plain_text","content":"为您智能创建了一个会议日程，请确认："},"subtitle":{"tag":"plain_text","content":""},"template":"orange","ud_icon":{"tag":"standard_icon","token":"calendar_colorful"}}}}`,
			},
		}, false},
		{"2", args{
			context.Background(), MsgOpt{
				Content: t2,
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendMessage(tt.args.ctx, tt.args.opt); (err != nil) != tt.wantErr {
				t.Errorf("SendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
