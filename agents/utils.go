package agents

import (
	"context"
	"encoding/json"
	"fmt"
	agentData "github.com/Ocyss/xiaomiHA/agents/data"
	"github.com/gocarina/gocsv"
)

func createChatMessages[T any](ctx context.Context, option *ChatOption) (*T, error) {
	content, _, err := CreateChat(ctx, option)
	if err != nil {
		return nil, fmt.Errorf("failed to CreateChat: %w", err)
	}
	var ans T
	if err := json.Unmarshal([]byte(content), &ans); err != nil {
		return nil, fmt.Errorf("failed to json Unmarshal: %w", err)
	}
	return &ans, nil
}
func getPromptAndCsvFile(d AgentCsv) (string, string, error) {
	prompt, err := SystemAgent(d)
	if err != nil {
		return "", "", fmt.Errorf("failed to SystemAgent: %w", err)
	}
	csv, err := d.GenCsvFile()
	if err != nil {
		return "", "", fmt.Errorf("failed to GenCsvFile: %w", err)
	}
	csvFile, err := gocsv.MarshalString(csv)
	if err != nil {
		return "", "", fmt.Errorf("failed to csv MarshalBytes: %w", err)
	}
	return prompt, csvFile, nil
}

func genCsvFile(data, dataAvg []agentData.GetCsvRecord) ([]*agentData.CsvData, error) {
	ans := make([]*agentData.CsvData, 0, (len(data)+len(dataAvg))*3)
	for _, v := range data {
		if csv, err := v.GetCsvRecord(); err != nil {
			return nil, err
		} else {
			ans = append(ans, csv...)
		}
	}
	for _, v := range dataAvg {
		if csv, err := v.GetCsvAverage(); err != nil {
			return nil, err
		} else {
			ans = append(ans, csv...)
		}
	}
	return ans, nil
}
