package agentData

import (
	"encoding/json"
	"fmt"
	"github.com/Ocyss/xiaomiHA/model"
	"github.com/Ocyss/xiaomiHA/query"
	"gorm.io/gen"
)

type WeightItem struct {
	Comm
}
type WeightItemJson struct {
	Bpm                int     `json:"bpm"`
	BodyAge            int     `json:"body_age"`
	BodyMoistureMass   float64 `json:"body_moisture_mass"`
	BodyScore          int     `json:"body_score"`
	BoneRate           float64 `json:"bone_rate"`
	FatControl         float64 `json:"fat_control"`
	FatFreeBody        float64 `json:"fat_free_body"`
	FatMass            float64 `json:"fat_mass"`
	MuscleControl      float64 `json:"muscle_control"`
	MuscleMass         float64 `json:"muscle_mass"`
	ProteinMass        float64 `json:"protein_mass"`
	ScoreStandardType  int     `json:"score_standard_type"`
	SkeletalMuscleMass float64 `json:"skeletal_muscle_mass"`
	Somatotype         int     `json:"somatotype"`
	StandardWeightV2   float64 `json:"standard_weight_v2"`
	Time               int     `json:"time"`
	Weight             float64 `json:"weight"`
	WeightControl      float64 `json:"weight_control"`
	Whr                float64 `json:"whr"`
}

func (h *WeightItem) GetRecord() ([]*model.WeightItemRecord, error) {
	var conds []gen.Condition
	if h.Time != nil {
		if h.Time.Time != nil {
			conds = append(conds, query.WeightItemRecord.Time.Gte(*h.Time.Time))
		} else {
			conds = append(conds, query.WeightItemRecord.Time.Between(*h.Time.StartTime, *h.Time.EndTime))
		}
	}
	q := query.WeightItemRecord.Where(conds...).Order(query.WeightItemRecord.Time.Desc())
	if h.Limit != nil {
		q = q.Limit(*h.Limit)
	}
	records, err := q.Find()
	return records, err
}

func (h *WeightItem) GetCsvRecord() ([]*CsvData, error) {
	record, err := h.GetRecord()
	if err != nil {
		return nil, err
	}
	ans := make([]*CsvData, 0, len(record))
	for _, v := range record {
		if v.Value != nil {
			var data WeightItemJson
			if err := json.Unmarshal([]byte(*v.Value), &data); err != nil {
				return nil, err
			}
			ans = append(ans, &CsvData{
				Type: "体重",
				Time: CsvTime{
					Time: &v.Time,
				},
				Value: fmt.Sprintf(`体重%.2fKg	体质评分%d	||	成分: 脂肪%.2f%%	水分%.2f%%	无机盐%.2f%%	蛋白质%.2f%%	||	去脂体重%.2fKg	肌肉量%.2fKg	骨骼肌量%.2fKg	骨盐量%.2fKg`,
					data.Weight,
					data.BodyScore,
					data.FatMass/data.Weight,
					data.BodyMoistureMass/data.Weight,
					data.BoneRate,
					data.ProteinMass/data.Weight,
					data.StandardWeightV2,
					data.MuscleMass,
					data.SkeletalMuscleMass,
					data.BoneRate*data.Weight,
				),
			})
		}
	}
	return ans, nil
}
func (h *WeightItem) GetCsvAverage() ([]*CsvData, error) {
	return nil, nil
}

var _ GetCsvRecord = (*WeightItem)(nil)
