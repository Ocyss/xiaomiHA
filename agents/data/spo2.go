package agentData

import (
	"github.com/Ocyss/xiaomiHA/model"
	"github.com/Ocyss/xiaomiHA/query"
	"github.com/Ocyss/xiaomiHA/utils"
	"github.com/tidwall/gjson"
	"gorm.io/gen"
	"strconv"
)

type Spo2 struct {
	Comm
}

func (h *Spo2) GetRecord() ([]*model.Spo2Record, error) {
	var conds []gen.Condition
	if h.Time != nil {
		if h.Time.Time != nil {
			conds = append(conds, query.Spo2Record.Time.Gte(*h.Time.Time))
		} else {
			conds = append(conds, query.Spo2Record.Time.Between(*h.Time.StartTime, *h.Time.EndTime))
		}
	}
	q := query.Spo2Record.Where(conds...).Order(query.Spo2Record.Time.Desc())
	if h.Limit != nil {
		q = q.Limit(*h.Limit)
	}
	records, err := q.Find()
	return records, err
}

func (h *Spo2) GetCsvRecord() ([]*CsvData, error) {
	record, err := h.GetRecord()
	if err != nil {
		return nil, err
	}
	ans := make([]*CsvData, 0, len(record))
	for _, v := range record {
		if v.Value != nil {
			ans = append(ans, &CsvData{
				Type: "血氧",
				Time: CsvTime{
					Time: &v.Time,
				},
				Value: gjson.Get(*v.Value, "spo2").String(),
			})
		}
	}
	return ans, nil
}
func (h *Spo2) GetCsvAverage() ([]*CsvData, error) {
	record, err := h.GetRecord()
	if err != nil {
		return nil, err
	}
	groups := utils.SliceGroup(record, h.AvgTime, func(i, j int, t int64) bool {
		return record[i].Time-t <= record[j].Time
	})
	ans := make([]*CsvData, 0, 10)
	for _, group := range groups {
		var val int64 = 0
		n := len(group)
		for _, v := range group {
			if v.Value == nil {
				n--
				continue
			}
			val += gjson.Get(*v.Value, "spo2").Int()
		}
		if n < 2 {
			continue
		}
		ans = append(ans, &CsvData{
			Type: "血氧avg",
			Time: CsvTime{
				StartTime: &group[0].Time,
				EndTime:   &group[len(group)-1].Time,
			},
			Value: strconv.FormatInt(val/int64(n), 10),
		})
	}
	return ans, nil
}

var _ GetCsvRecord = (*Spo2)(nil)
