package agentData

import (
	"github.com/Ocyss/xiaomiHA/model"
	"github.com/Ocyss/xiaomiHA/query"
	"github.com/Ocyss/xiaomiHA/utils"
	"github.com/tidwall/gjson"
	"gorm.io/gen"
	"strconv"
)

type Stress struct {
	Comm
}

func (h *Stress) GetRecord() ([]*model.StressRecord, error) {
	var conds []gen.Condition
	if h.Time != nil {
		if h.Time.Time != nil {
			conds = append(conds, query.StressRecord.Time.Gte(*h.Time.Time))
		} else {
			conds = append(conds, query.StressRecord.Time.Between(*h.Time.StartTime, *h.Time.EndTime))
		}
	}
	q := query.StressRecord.Where(conds...).Order(query.StressRecord.Time.Desc())
	if h.Limit != nil {
		q = q.Limit(*h.Limit)
	}
	records, err := q.Find()
	return records, err
}

func (h *Stress) GetCsvRecord() ([]*CsvData, error) {
	record, err := h.GetRecord()
	if err != nil {
		return nil, err
	}
	ans := make([]*CsvData, 0, len(record))
	for _, v := range record {
		if v.Value != nil {
			ans = append(ans, &CsvData{
				Type: "压力",
				Time: CsvTime{
					Time: &v.Time,
				},
				Value: gjson.Get(*v.Value, "stress").String(),
			})
		}
	}
	return ans, nil
}
func (h *Stress) GetCsvAverage() ([]*CsvData, error) {
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
			val += gjson.Get(*v.Value, "stress").Int()
		}
		if n < 2 {
			continue
		}
		ans = append(ans, &CsvData{
			Type: "压力avg",
			Time: CsvTime{
				StartTime: &group[0].Time,
				EndTime:   &group[len(group)-1].Time,
			},
			Value: strconv.FormatInt(val/int64(n), 10),
		})
	}
	return ans, nil
}

var _ GetCsvRecord = (*Stress)(nil)
