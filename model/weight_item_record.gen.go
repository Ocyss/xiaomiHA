// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWeightItemRecord = "weight_item_record"

// WeightItemRecord mapped from table <weight_item_record>
type WeightItemRecord struct {
	Key             string   `gorm:"column:key;type:TEXT;primaryKey" json:"key"`
	Sid             string   `gorm:"column:sid;type:TEXT;not null" json:"sid"`
	Time            int32    `gorm:"column:time;type:INTEGER;not null" json:"time"`
	ZeroTimeIn0Tz   int32    `gorm:"column:zeroTimeIn0Tz;type:INTEGER;not null" json:"zeroTimeIn0Tz"`
	Weight          float64  `gorm:"column:weight;type:REAL;not null" json:"weight"`
	Bmi             *float64 `gorm:"column:bmi;type:REAL" json:"bmi"`
	BodyFatRate     *float64 `gorm:"column:bodyFatRate;type:REAL" json:"bodyFatRate"`
	MoistureRate    *float64 `gorm:"column:moistureRate;type:REAL" json:"moistureRate"`
	BoneMass        *float64 `gorm:"column:boneMass;type:REAL" json:"boneMass"`
	BasalMetabolism *int32   `gorm:"column:basalMetabolism;type:INTEGER" json:"basalMetabolism"`
	MuscleMass      *float64 `gorm:"column:muscleMass;type:REAL" json:"muscleMass"`
	ProteinRate     *float64 `gorm:"column:proteinRate;type:REAL" json:"proteinRate"`
	VisceralFat     *float64 `gorm:"column:visceralFat;type:REAL" json:"visceralFat"`
	Value           *string  `gorm:"column:value;type:TEXT" json:"value"`
	ZoneOffsetInSec int32    `gorm:"column:zoneOffsetInSec;type:INTEGER;not null" json:"zoneOffsetInSec"`
	ZoneName        *string  `gorm:"column:zoneName;type:TEXT" json:"zoneName"`
	TimeIn0Tz       int32    `gorm:"column:timeIn0Tz;type:INTEGER;not null" json:"timeIn0Tz"`
	IsUpload        int32    `gorm:"column:isUpload;type:INTEGER;not null" json:"isUpload"`
	IsDeleted       int32    `gorm:"column:isDeleted;type:INTEGER;not null" json:"isDeleted"`
}

// TableName WeightItemRecord's table name
func (*WeightItemRecord) TableName() string {
	return TableNameWeightItemRecord
}