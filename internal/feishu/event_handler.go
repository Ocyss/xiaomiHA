package feishu

import (
	"context"
	"fmt"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkcalendar "github.com/larksuite/oapi-sdk-go/v3/service/calendar/v4"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

// 接收消息v2.0
func p2MessageReceiveV1(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
	fmt.Printf("[ OnP2MessageReceiveV1 access ], data: %s\n", larkcore.Prettify(event))
	return nil
}

// 用户进入与机器人的会话v2.0
func p2ChatAccessEventBotP2pChatEnteredV1(ctx context.Context, event *larkim.P2ChatAccessEventBotP2pChatEnteredV1) error {
	return nil
}

// 消息已读v2.0
func p2MessageReadV1(ctx context.Context, event *larkim.P2MessageReadV1) error { return nil }

// 消息被reactionv2.0
func p2MessageReactionCreatedV1(ctx context.Context, event *larkim.P2MessageReactionCreatedV1) error {
	return nil
}

// 消息被取消reactionv2.0
func p2MessageReactionDeletedV1(ctx context.Context, event *larkim.P2MessageReactionDeletedV1) error {
	return nil
}

// 消息撤回v2.0
func p2MessageRecalledV1(ctx context.Context, event *larkim.P2MessageRecalledV1) error { return nil }

// 用户和机器人的会话首次被创建v1.0
func p1P2PChatCreatedV1(ctx context.Context, event *larkim.P1P2PChatCreatedV1) error { return nil }

// ACL创建v2.0
func p2CalendarAclCreatedV4(ctx context.Context, event *larkcalendar.P2CalendarAclCreatedV4) error {
	return nil
}

// ACL删除v2.0
func p2CalendarAclDeletedV4(ctx context.Context, event *larkcalendar.P2CalendarAclDeletedV4) error {
	return nil
}

// 日历变更v2.0
func p2CalendarChangedV4(ctx context.Context, event *larkcalendar.P2CalendarChangedV4) error {
	return nil
}

// 日程变更v2.0
func p2CalendarEventChangedV4(ctx context.Context, event *larkcalendar.P2CalendarEventChangedV4) error {
	return nil
}
