package agentData

import (
	"fmt"
	"github.com/Ocyss/xiaomiHA/model"
	"github.com/Ocyss/xiaomiHA/query"
	"github.com/Ocyss/xiaomiHA/utils"
	"github.com/tidwall/gjson"
	"gorm.io/gen"
)

type Step struct {
	Comm
}

func (h *Step) GetRecord() ([]*model.StepRecord, error) {
	var conds []gen.Condition
	if h.Time != nil {
		if h.Time.Time != nil {
			conds = append(conds, query.StepRecord.Time.Gte(*h.Time.Time))
		} else {
			conds = append(conds, query.StepRecord.Time.Between(*h.Time.StartTime, *h.Time.EndTime))
		}
	}
	q := query.StepRecord.Where(conds...).Order(query.StepRecord.Time.Desc())
	if h.Limit != nil {
		q = q.Limit(*h.Limit)
	}
	records, err := q.Find()
	return records, err
}

func (h *Step) GetCsvRecord() ([]*CsvData, error) {
	record, err := h.GetRecord()
	if err != nil {
		return nil, err
	}
	ans := make([]*CsvData, 0, len(record))
	for _, v := range record {
		if v.Value != nil {
			data := gjson.Parse(*v.Value)
			ans = append(ans, &CsvData{
				Type: "步数/距离",
				Time: CsvTime{
					Time: &v.Time,
				},
				Value: fmt.Sprintf("%s/%s", data.Get("steps"), data.Get("distance")),
			})
		}
	}
	return ans, nil
}
func (h *Step) GetCsvAverage() ([]*CsvData, error) {
	record, err := h.GetRecord()
	if err != nil {
		return nil, err
	}
	groups := utils.SliceGroup(record, h.AvgTime, func(i, j int, t int64) bool {
		return record[i].Time-t <= record[j].Time
	})
	ans := make([]*CsvData, 0, 10)
	for _, group := range groups {
		var val1, val2 int64 = 0, 0
		n := len(group)
		for _, v := range group {
			if v.Value == nil {
				n--
				continue
			}
			val1 += gjson.Get(*v.Value, "steps").Int()
			val2 += gjson.Get(*v.Value, "distance").Int()
		}
		if n < 2 {
			continue
		}
		ans = append(ans, &CsvData{
			Type: "步数/距离 avg",
			Time: CsvTime{
				StartTime: &group[0].Time,
				EndTime:   &group[len(group)-1].Time,
			},
			Value: fmt.Sprintf("%d/%d", val1/int64(n), val2/int64(n)),
		})
	}
	return ans, nil
}

var _ GetCsvRecord = (*Step)(nil)
