package agentData

import (
	"github.com/Ocyss/xiaomiHA/model"
	"github.com/Ocyss/xiaomiHA/query"
	"time"
)

type CsvData struct {
	Type  string  `csv:"type"`
	Time  CsvTime `csv:"time"`
	Value string  `csv:"value"`
}

type CsvTime struct {
	Time      *int64
	StartTime *int64
	EndTime   *int64
}

const layout = "06/01/02 15:04"

func (t *CsvTime) MarshalCSV() (string, error) {
	if t.StartTime != nil && t.EndTime != nil {
		return time.Unix(int64(*t.StartTime), 0).Format(layout) + "到" + time.Unix(int64(*t.EndTime), 0).Format(layout), nil
	}
	return time.Unix(int64(*t.Time), 0).Format(layout), nil
}

func (t *CsvTime) UnmarshalCSV(csv string) (err error) {
	return nil
}

func init() {
	query.SetDefault(model.Db)
}

type GetCsvRecord interface {
	GetCsvRecord() ([]*CsvData, error)
	GetCsvAverage() ([]*CsvData, error)
}

type Comm struct {
	Time    *CsvTime
	Limit   *int
	AvgTime *time.Duration
}

func NewLimitComm(n int) Comm {
	return Comm{Limit: &n}
}
func NewTimeComm(n *int64) Comm {
	return Comm{Time: &CsvTime{Time: n}}
}
func NewTime2Comm(n1, n2 *int64) Comm {
	return Comm{Time: &CsvTime{StartTime: n1, EndTime: n2}}
}
