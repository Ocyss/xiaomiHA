package agents

import (
	"context"
	"errors"
	"fmt"
	agentTools "github.com/Ocyss/xiaomiHA/agents/tools"
	"github.com/Ocyss/xiaomiHA/utils"
	"github.com/sashabaranov/go-openai"
)

var client = initClient()

func initClient() *openai.Client {
	openaiConfig := openai.DefaultConfig(utils.C.Gpt.ApiKey)
	openaiConfig.BaseURL = utils.C.Gpt.ApiBase
	return openai.NewClientWithConfig(openaiConfig)
}

type ChatOption struct {
	Model         string
	Messages      []openai.ChatCompletionMessage
	RespFormat    *openai.ChatCompletionResponseFormat
	ToolChoice    any
	toolCallCount int
}

func CreateChat(ctx context.Context, o *ChatOption) (string, *openai.ChatCompletionResponse, error) {
	if o.toolCallCount > 8 {
		return "", nil, errors.New("too many tool calls")
	}
	if o.Model == "" {
		o.Model = openai.GPT4o
	}
	if o.RespFormat == nil {
		o.RespFormat = &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONObject,
		}
	}
	completion, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:          o.Model,
		Messages:       o.Messages,
		ResponseFormat: o.RespFormat,
		Tools:          agentTools.Tools,
		ToolChoice:     o.ToolChoice,
	})
	if err != nil {
		return "", nil, fmt.Errorf("failed to CreateChatCompletion: %w", err)
	} else if len(completion.Choices) == 0 {
		return "", nil, errors.New("failed to choices length")
	}
	msg := completion.Choices[0].Message
	if len(msg.ToolCalls) > 0 {
		newMsg := append([]openai.ChatCompletionMessage{}, o.Messages...)
		for _, toolCall := range msg.ToolCalls {
			if t, ok := agentTools.ToolsCall[toolCall.Function.Name]; ok {
				out, err := t.Func(toolCall.Function.Arguments)
				if err != nil {
					return "", nil, err
				}
				newMsg = append(newMsg, msg, openai.ChatCompletionMessage{
					Role:       openai.ChatMessageRoleTool,
					Content:    out,
					Name:       toolCall.Function.Name,
					ToolCallID: toolCall.ID,
				})
			}
		}
		if len(newMsg) > len(o.Messages) {
			o.toolCallCount++
			o.Messages = newMsg
			o.ToolChoice = nil
			return CreateChat(ctx, o)
		} else {
			return "", nil, errors.New("the error toolsCalls")
		}
	}
	return completion.Choices[0].Message.Content, &completion, nil
}
