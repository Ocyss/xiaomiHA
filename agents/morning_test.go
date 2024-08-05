package agents

import (
	"context"
	"github.com/gocarina/gocsv"
	"os"
	"testing"
)

func TestMorningAgent_GenCsvFile(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"1", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MorningAgent{}
			got, err := d.GenCsvFile()
			if (err != nil) != tt.wantErr {
				t.Errorf("GenCsvFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//t.Log(got)
			temp, err := os.OpenFile("xiaomiHA_test.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
			defer temp.Close()
			if err != nil {
				t.Errorf("create temp file error: %v", err)
				return
			}
			err = gocsv.MarshalFile(&got, temp)
			if err != nil {
				t.Errorf("unmarshal file error: %v", err)
				return
			}

		})
	}
}

func TestMorningAgent_ProcessTask(t *testing.T) {
	tests := []struct {
		name    string
		ctx     context.Context
		wantErr bool
	}{
		{"1", context.Background(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MorningAgent{}
			err := d.ProcessTask(tt.ctx, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("Messages() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMorningAgent_SystemAgentOptions(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"1", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := SystemAgent(&MorningAgent{})
			if (err != nil) != tt.wantErr {
				t.Errorf("Messages() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Log(d)
		})
	}
}
