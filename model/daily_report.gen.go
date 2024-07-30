// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameDailyReport = "daily_report"

// DailyReport mapped from table <daily_report>
type DailyReport struct {
	Sid        string `gorm:"column:sid;type:TEXT;not null" json:"sid"`
	DataType   string `gorm:"column:dataType;type:TEXT;primaryKey" json:"dataType"`
	UpdateTime int32  `gorm:"column:updateTime;type:INTEGER;not null" json:"updateTime"`
	TimeInZero int32  `gorm:"column:timeInZero;type:INTEGER;not null" json:"timeInZero"`
	ViewTag    string `gorm:"column:viewTag;type:TEXT;not null" json:"viewTag"`
	Value      string `gorm:"column:value;type:TEXT;not null" json:"value"`
	IsUpload   int32  `gorm:"column:isUpload;type:INTEGER;not null" json:"isUpload"`
	IsDeleted  int32  `gorm:"column:isDeleted;type:INTEGER;not null" json:"isDeleted"`
}

// TableName DailyReport's table name
func (*DailyReport) TableName() string {
	return TableNameDailyReport
}