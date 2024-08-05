package agentData

import (
	"fmt"
	"github.com/Ocyss/xiaomiHA/model"
	"github.com/Ocyss/xiaomiHA/query"
	"github.com/Ocyss/xiaomiHA/utils"
	"github.com/tidwall/gjson"
	"gorm.io/gen"
)

type Calorie struct {
	Comm
}

func (h *Calorie) GetRecord() ([]*model.CalorieRecord, error) {
	var conds []gen.Condition
	if h.Time != nil {
		if h.Time.Time != nil {
			conds = append(conds, query.CalorieRecord.Time.Gte(*h.Time.Time))
		} else {
			conds = append(conds, query.CalorieRecord.Time.Between(*h.Time.StartTime, *h.Time.EndTime))
		}
	}
	q := query.CalorieRecord.Where(conds...).Order(query.CalorieRecord.Time.Desc())
	if h.Limit != nil {
		q = q.Limit(*h.Limit)
	}
	records, err := q.Find()
	return records, err
}

func (h *Calorie) GetCsvRecord() ([]*CsvData, error) {
	record, err := h.GetRecord()
	if err != nil {
		return nil, err
	}
	ans := make([]*CsvData, 0, len(record))
	for _, v := range record {
		if v.Value != nil {
			ans = append(ans, &CsvData{
				Type: "卡路里",
				Time: CsvTime{
					Time: &v.Time,
				},
				Value: gjson.Get(*v.Value, "calories").String(),
			})
		}
	}
	return ans, nil
}

func (h *Calorie) GetCsvAverage() ([]*CsvData, error) {
	record, err := h.GetRecord()
	if err != nil {
		return nil, err
	}
	groups := utils.SliceGroup(record, h.AvgTime, func(i, j int, t int64) bool {
		return record[i].Time-t <= record[j].Time
	})
	ans := make([]*CsvData, 0, 10)
	for _, group := range groups {
		var val float64 = 0
		n := len(group)
		for _, v := range group {
			if v.Value == nil {
				n--
				continue
			}
			val += gjson.Get(*v.Value, "calories").Float()
		}
		if n < 2 {
			continue
		}
		ans = append(ans, &CsvData{
			Type: "卡路里avg",
			Time: CsvTime{
				StartTime: &group[0].Time,
				EndTime:   &group[len(group)-1].Time,
			},
			Value: fmt.Sprintf("%.2f", val/float64(n)),
		})
	}
	return ans, nil
}

var _ GetCsvRecord = (*Calorie)(nil)
