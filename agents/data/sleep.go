package agentData

import (
	"fmt"
	"github.com/Ocyss/xiaomiHA/model"
	"github.com/Ocyss/xiaomiHA/query"
	"github.com/tidwall/gjson"
	"gorm.io/gen"
	"time"
)

type SleepSegment struct {
	Comm
}

func (h *SleepSegment) GetRecord() ([]*model.SleepSegment, error) {
	var conds []gen.Condition
	if h.Time != nil {
		if h.Time.Time != nil {
			conds = append(conds, query.SleepSegment.Time.Gte(*h.Time.Time))
		} else {
			conds = append(conds, query.SleepSegment.Time.Between(*h.Time.StartTime, *h.Time.EndTime))
		}
	}
	q := query.SleepSegment.Where(conds...).Order(query.SleepSegment.Time.Desc())
	if h.Limit != nil {
		q = q.Limit(*h.Limit)
	}
	records, err := q.Find()
	return records, err
}

func (h *SleepSegment) GetCsvRecord() ([]*CsvData, error) {
	record, err := h.GetRecord()
	if err != nil {
		return nil, err
	}
	ans := make([]*CsvData, 0, len(record))

	for _, v := range record {
		if v.Value != nil {
			data := gjson.Parse(*v.Value)
			ans = append(ans, &CsvData{
				Type: "睡眠",
				Time: CsvTime{
					Time: &v.Time,
				},
				Value: fmt.Sprintf(`就寝时间:%s	起床时间:%s	睡眠总时长:%s	||	深睡时长:%s	浅睡时长:%s	REM睡眠时长:%s	||	清醒次数:%s	清醒时长:%s`,
					time.Unix(data.Get("bedtime").Int(), 0).Format(layout),
					time.Unix(data.Get("wake_up_time").Int(), 0).Format(layout),
					data.Get("duration"),
					data.Get("sleep_deep_duration"),
					data.Get("sleep_light_duration"),
					data.Get("sleep_rem_duration"),
					data.Get("awake_count"),
					data.Get("sleep_awake_duration")),
			})
		}
	}
	return ans, nil
}

func (h *SleepSegment) GetCsvAverage() ([]*CsvData, error) {
	return nil, nil
}

var _ GetCsvRecord = (*SleepSegment)(nil)
