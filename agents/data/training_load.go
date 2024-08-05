package agentData

import (
	"github.com/Ocyss/xiaomiHA/model"
	"github.com/Ocyss/xiaomiHA/query"
	"github.com/tidwall/gjson"
	"gorm.io/gen"
)

type TrainingLoad struct {
	Comm
}

func (h *TrainingLoad) GetRecord() ([]*model.TrainingLoadRecord, error) {
	var conds []gen.Condition
	if h.Time != nil {
		if h.Time.Time != nil {
			conds = append(conds, query.TrainingLoadRecord.Time.Gte(*h.Time.Time))
		} else {
			conds = append(conds, query.TrainingLoadRecord.Time.Between(*h.Time.StartTime, *h.Time.EndTime))
		}
	}
	q := query.TrainingLoadRecord.Where(conds...).Order(query.TrainingLoadRecord.Time.Desc())
	if h.Limit != nil {
		q = q.Limit(*h.Limit)
	}
	records, err := q.Find()
	return records, err
}

func (h *TrainingLoad) GetCsvRecord() ([]*CsvData, error) {
	record, err := h.GetRecord()
	if err != nil {
		return nil, err
	}
	ans := make([]*CsvData, 0, len(record))
	for _, v := range record {
		if v.Value != nil {
			ans = append(ans, &CsvData{
				Type: "训练负荷",
				Time: CsvTime{
					Time: &v.Time,
				},
				Value: gjson.Get(*v.Value, "wtl_sum").String(),
			})
		}
	}
	return ans, nil
}
func (h *TrainingLoad) GetCsvAverage() ([]*CsvData, error) {
	return nil, nil
}

var _ GetCsvRecord = (*TrainingLoad)(nil)
