package agentData

import (
	"fmt"
	"github.com/Ocyss/xiaomiHA/model"
	"github.com/Ocyss/xiaomiHA/query"
	"github.com/tidwall/gjson"
	"gorm.io/gen"
)

type SportReport struct {
	Comm
}

func (h *SportReport) GetRecord() ([]*model.SportReport, error) {
	var conds []gen.Condition
	if h.Time != nil {
		if h.Time.Time != nil {
			conds = append(conds, query.SportReport.Time.Gte(*h.Time.Time))
		} else {
			conds = append(conds, query.SportReport.Time.Between(*h.Time.StartTime, *h.Time.EndTime))
		}
	}
	q := query.SportReport.Where(conds...).Order(query.SportReport.Time.Desc())
	if h.Limit != nil {
		q = q.Limit(*h.Limit)
	}
	records, err := q.Find()
	return records, err
}

func (h *SportReport) GetCsvRecord() ([]*CsvData, error) {
	record, err := h.GetRecord()
	if err != nil {
		return nil, err
	}
	ans := make([]*CsvData, 0, len(record))
	for _, v := range record {
		if v.Value != nil {
			data := gjson.Parse(*v.Value)
			ans = append(ans, &CsvData{
				Type: "运动报告",
				Time: CsvTime{
					Time: &v.Time,
				},
				Value: fmt.Sprintf(`消耗卡路里:%s	运动时长:%s
平均心率:%s	最高心率:%s	最低心率:%s`,
					data.Get("calories"),
					data.Get("duration"),
					data.Get("avg_hrm"),
					data.Get("max_hrm"),
					data.Get("min_hrm"),
				),
			})
		}
	}
	return ans, nil
}
func (h *SportReport) GetCsvAverage() ([]*CsvData, error) {
	return nil, nil
}
