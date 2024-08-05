package feishu

import (
	"context"
	"github.com/Ocyss/xiaomiHA/utils"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
)

var (
	conf   = utils.C.FeiShu
	client = lark.NewClient(conf.AppId, conf.AppSecret)
)

func StartWsClient() {
	eventHandler := dispatcher.NewEventDispatcher(conf.VerificationToken, conf.EncryptKey).
		OnP2MessageReceiveV1(p2MessageReceiveV1).                                     // 接收消息v2.0
		OnP2ChatAccessEventBotP2pChatEnteredV1(p2ChatAccessEventBotP2pChatEnteredV1). // 用户进入与机器人的会话v2.0
		OnP2MessageReadV1(p2MessageReadV1).                                           // 消息已读v2.0
		OnP2MessageReactionCreatedV1(p2MessageReactionCreatedV1).                     // 消息被reactionv2.0
		OnP2MessageReactionDeletedV1(p2MessageReactionDeletedV1).                     // 消息被取消reactionv2.0
		OnP2MessageRecalledV1(p2MessageRecalledV1).                                   // 消息撤回v2.0
		OnP1P2PChatCreatedV1(p1P2PChatCreatedV1).                                     // 用户和机器人的会话首次被创建v1.0
		OnP2CalendarAclCreatedV4(p2CalendarAclCreatedV4).                             // ACL创建v2.0
		OnP2CalendarAclDeletedV4(p2CalendarAclDeletedV4).                             // ACL删除v2.0
		OnP2CalendarChangedV4(p2CalendarChangedV4).                                   // 日历变更v2.0
		OnP2CalendarEventChangedV4(p2CalendarEventChangedV4)                          // 日程变更v2.0

	cli := larkws.NewClient(conf.AppId, conf.AppSecret,
		larkws.WithEventHandler(eventHandler),
		larkws.WithLogLevel(larkcore.LogLevelDebug),
	)

	err := cli.Start(context.Background())
	if err != nil {
		panic(err)
	}
}
