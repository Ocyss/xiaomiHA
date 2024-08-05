package agentData

import (
	"fmt"
	"github.com/Ocyss/xiaomiHA/model"
	"github.com/Ocyss/xiaomiHA/query"
	"github.com/tidwall/gjson"
	"gorm.io/gen"
)

type Vitality struct {
	Comm
}

func (h *Vitality) GetRecord() ([]*model.VitalityRecord, error) {
	var conds []gen.Condition
	if h.Time != nil {
		if h.Time.Time != nil {
			conds = append(conds, query.VitalityRecord.Time.Gte(*h.Time.Time))
		} else {
			conds = append(conds, query.VitalityRecord.Time.Between(*h.Time.StartTime, *h.Time.EndTime))
		}
	}
	q := query.VitalityRecord.Where(conds...).Order(query.VitalityRecord.Time.Desc())
	if h.Limit != nil {
		q = q.Limit(*h.Limit)
	}
	records, err := q.Find()
	return records, err
}

func (h *Vitality) GetCsvRecord() ([]*CsvData, error) {
	record, err := h.GetRecord()
	if err != nil {
		return nil, err
	}
	ans := make([]*CsvData, 0, len(record))
	for _, v := range record {
		if v.Value != nil {
			data := gjson.Parse(*v.Value)
			ans = append(ans, &CsvData{
				Type: "活力值",
				Time: CsvTime{
					Time: &v.Time,
				},
				Value: fmt.Sprintf(`每日高中低强度活力：%d,%d,%d,建议持续时间/类型: %d,%d`,
					data.Get("daily_high_intensity_vitality").Int(),
					data.Get("daily_medium_intensity_vitality").Int(),
					data.Get("daily_low_intensity_vitality").Int(),
					data.Get("suggested_activity_duration").Int(),
					data.Get("suggested_activity_type").Int()),
			})
		}
	}
	return ans, nil
}
func (h *Vitality) GetCsvAverage() ([]*CsvData, error) {
	return nil, nil
}

var _ GetCsvRecord = (*Vitality)(nil)
