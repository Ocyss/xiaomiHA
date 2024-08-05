package feishu

import (
	"context"
	"fmt"
	"github.com/Ocyss/xiaomiHA/utils"
	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

type MsgOpt struct {
	ReceiveId string
	UUID      string

	Card    *larkcard.MessageCard
	Content string
}

func SendMessage(ctx context.Context, opt MsgOpt) (err error) {
	if opt.ReceiveId == "" {
		opt.ReceiveId = utils.C.FeiShu.OpenId
	}
	content := opt.Content
	if opt.Card != nil {
		content, err = opt.Card.JSON()
		if err != nil {
			return fmt.Errorf("error to convert card to json: %w", err)
		}
	}
	resp, err := client.Im.Message.Create(ctx, larkim.NewCreateMessageReqBuilder().
		ReceiveIdType(`open_id`).
		Body(larkim.NewCreateMessageReqBodyBuilder().
			ReceiveId(opt.ReceiveId).
			MsgType("interactive").
			Content(content).
			Uuid(opt.UUID).
			Build()).
		Build())

	if err != nil {
		return fmt.Errorf("error to send message: %w", err)
	}

	if !resp.Success() {
		return fmt.Errorf("failed to send message:[%d] %s", resp.Code, resp.Msg)
	}

	return nil
}
