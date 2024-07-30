package agents

import "testing"

func TestSystemAgent(t *testing.T) {
	tests := []struct {
		name    string
		args    Agent
		want    string
		wantErr bool
	}{
		{name: "1", args: &TestAgent{}, want: "", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SystemAgent(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("SystemAgent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
			//if got != tt.want {
			//	t.Errorf("SystemAgent() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

type TestAgent struct {
	Time int64  `json:"time" ts_doc:"需要一个Time时间"`
	Name string `json:"name" ts_doc:"返回一个Name"`
	Age  int64  `json:"age" ts_doc:"用户的Age"`
	Sex  string `json:"sex"`
}

func (t *TestAgent) TimeWindow() string {
	return "20分钟内"
}

func (t *TestAgent) Skills() []string {
	return []string{"skill1", "skill2", "skill3"}
}

func (t *TestAgent) Goals() []string {
	return []string{"goal1", "goal2"}
}

func (t *TestAgent) Constrains() []string {
	return nil
}

func (t *TestAgent) Workflow() []string {
	return []string{"work1", "work2"}
}

func (t *TestAgent) OutputFormat() any {
	return t
}
