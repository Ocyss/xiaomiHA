// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                  = new(Query)
	CalorieRecord      *calorieRecord
	HrRecord           *hrRecord
	SleepSegment       *sleepSegment
	Spo2Record         *spo2Record
	SportReport        *sportReport
	StepRecord         *stepRecord
	StressRecord       *stressRecord
	TrainingLoadRecord *trainingLoadRecord
	VitalityRecord     *vitalityRecord
	WeightItemRecord   *weightItemRecord
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	CalorieRecord = &Q.CalorieRecord
	HrRecord = &Q.HrRecord
	SleepSegment = &Q.SleepSegment
	Spo2Record = &Q.Spo2Record
	SportReport = &Q.SportReport
	StepRecord = &Q.StepRecord
	StressRecord = &Q.StressRecord
	TrainingLoadRecord = &Q.TrainingLoadRecord
	VitalityRecord = &Q.VitalityRecord
	WeightItemRecord = &Q.WeightItemRecord
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                 db,
		CalorieRecord:      newCalorieRecord(db, opts...),
		HrRecord:           newHrRecord(db, opts...),
		SleepSegment:       newSleepSegment(db, opts...),
		Spo2Record:         newSpo2Record(db, opts...),
		SportReport:        newSportReport(db, opts...),
		StepRecord:         newStepRecord(db, opts...),
		StressRecord:       newStressRecord(db, opts...),
		TrainingLoadRecord: newTrainingLoadRecord(db, opts...),
		VitalityRecord:     newVitalityRecord(db, opts...),
		WeightItemRecord:   newWeightItemRecord(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	CalorieRecord      calorieRecord
	HrRecord           hrRecord
	SleepSegment       sleepSegment
	Spo2Record         spo2Record
	SportReport        sportReport
	StepRecord         stepRecord
	StressRecord       stressRecord
	TrainingLoadRecord trainingLoadRecord
	VitalityRecord     vitalityRecord
	WeightItemRecord   weightItemRecord
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                 db,
		CalorieRecord:      q.CalorieRecord.clone(db),
		HrRecord:           q.HrRecord.clone(db),
		SleepSegment:       q.SleepSegment.clone(db),
		Spo2Record:         q.Spo2Record.clone(db),
		SportReport:        q.SportReport.clone(db),
		StepRecord:         q.StepRecord.clone(db),
		StressRecord:       q.StressRecord.clone(db),
		TrainingLoadRecord: q.TrainingLoadRecord.clone(db),
		VitalityRecord:     q.VitalityRecord.clone(db),
		WeightItemRecord:   q.WeightItemRecord.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                 db,
		CalorieRecord:      q.CalorieRecord.replaceDB(db),
		HrRecord:           q.HrRecord.replaceDB(db),
		SleepSegment:       q.SleepSegment.replaceDB(db),
		Spo2Record:         q.Spo2Record.replaceDB(db),
		SportReport:        q.SportReport.replaceDB(db),
		StepRecord:         q.StepRecord.replaceDB(db),
		StressRecord:       q.StressRecord.replaceDB(db),
		TrainingLoadRecord: q.TrainingLoadRecord.replaceDB(db),
		VitalityRecord:     q.VitalityRecord.replaceDB(db),
		WeightItemRecord:   q.WeightItemRecord.replaceDB(db),
	}
}

type queryCtx struct {
	CalorieRecord      ICalorieRecordDo
	HrRecord           IHrRecordDo
	SleepSegment       ISleepSegmentDo
	Spo2Record         ISpo2RecordDo
	SportReport        ISportReportDo
	StepRecord         IStepRecordDo
	StressRecord       IStressRecordDo
	TrainingLoadRecord ITrainingLoadRecordDo
	VitalityRecord     IVitalityRecordDo
	WeightItemRecord   IWeightItemRecordDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		CalorieRecord:      q.CalorieRecord.WithContext(ctx),
		HrRecord:           q.HrRecord.WithContext(ctx),
		SleepSegment:       q.SleepSegment.WithContext(ctx),
		Spo2Record:         q.Spo2Record.WithContext(ctx),
		SportReport:        q.SportReport.WithContext(ctx),
		StepRecord:         q.StepRecord.WithContext(ctx),
		StressRecord:       q.StressRecord.WithContext(ctx),
		TrainingLoadRecord: q.TrainingLoadRecord.WithContext(ctx),
		VitalityRecord:     q.VitalityRecord.WithContext(ctx),
		WeightItemRecord:   q.WeightItemRecord.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
