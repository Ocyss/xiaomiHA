// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTrainingLoadRecord = "training_load_record"

// TrainingLoadRecord mapped from table <training_load_record>
type TrainingLoadRecord struct {
	Key             string  `gorm:"column:key;type:TEXT;primaryKey" json:"key"`
	Sid             string  `gorm:"column:sid;type:TEXT;not null" json:"sid"`
	Time            int32   `gorm:"column:time;type:INTEGER;not null" json:"time"`
	Value           *string `gorm:"column:value;type:TEXT" json:"value"`
	ZoneOffsetInSec int32   `gorm:"column:zoneOffsetInSec;type:INTEGER;not null" json:"zoneOffsetInSec"`
	ZoneName        *string `gorm:"column:zoneName;type:TEXT" json:"zoneName"`
	TimeIn0Tz       int32   `gorm:"column:timeIn0Tz;type:INTEGER;not null" json:"timeIn0Tz"`
	IsUpload        int32   `gorm:"column:isUpload;type:INTEGER;not null" json:"isUpload"`
	IsDeleted       int32   `gorm:"column:isDeleted;type:INTEGER;not null" json:"isDeleted"`
}

// TableName TrainingLoadRecord's table name
func (*TrainingLoadRecord) TableName() string {
	return TableNameTrainingLoadRecord
}
