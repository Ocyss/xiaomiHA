// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/Ocyss/xiaomiHA/model"
)

func newWeightItemRecord(db *gorm.DB, opts ...gen.DOOption) weightItemRecord {
	_weightItemRecord := weightItemRecord{}

	_weightItemRecord.weightItemRecordDo.UseDB(db, opts...)
	_weightItemRecord.weightItemRecordDo.UseModel(&model.WeightItemRecord{})

	tableName := _weightItemRecord.weightItemRecordDo.TableName()
	_weightItemRecord.ALL = field.NewAsterisk(tableName)
	_weightItemRecord.Key = field.NewString(tableName, "key")
	_weightItemRecord.Sid = field.NewString(tableName, "sid")
	_weightItemRecord.Time = field.NewInt32(tableName, "time")
	_weightItemRecord.ZeroTimeIn0Tz = field.NewInt32(tableName, "zeroTimeIn0Tz")
	_weightItemRecord.Weight = field.NewFloat64(tableName, "weight")
	_weightItemRecord.Bmi = field.NewFloat64(tableName, "bmi")
	_weightItemRecord.BodyFatRate = field.NewFloat64(tableName, "bodyFatRate")
	_weightItemRecord.MoistureRate = field.NewFloat64(tableName, "moistureRate")
	_weightItemRecord.BoneMass = field.NewFloat64(tableName, "boneMass")
	_weightItemRecord.BasalMetabolism = field.NewInt32(tableName, "basalMetabolism")
	_weightItemRecord.MuscleMass = field.NewFloat64(tableName, "muscleMass")
	_weightItemRecord.ProteinRate = field.NewFloat64(tableName, "proteinRate")
	_weightItemRecord.VisceralFat = field.NewFloat64(tableName, "visceralFat")
	_weightItemRecord.Value = field.NewString(tableName, "value")
	_weightItemRecord.ZoneOffsetInSec = field.NewInt32(tableName, "zoneOffsetInSec")
	_weightItemRecord.ZoneName = field.NewString(tableName, "zoneName")
	_weightItemRecord.TimeIn0Tz = field.NewInt32(tableName, "timeIn0Tz")
	_weightItemRecord.IsUpload = field.NewInt32(tableName, "isUpload")
	_weightItemRecord.IsDeleted = field.NewInt32(tableName, "isDeleted")

	_weightItemRecord.fillFieldMap()

	return _weightItemRecord
}

type weightItemRecord struct {
	weightItemRecordDo

	ALL             field.Asterisk
	Key             field.String
	Sid             field.String
	Time            field.Int32
	ZeroTimeIn0Tz   field.Int32
	Weight          field.Float64
	Bmi             field.Float64
	BodyFatRate     field.Float64
	MoistureRate    field.Float64
	BoneMass        field.Float64
	BasalMetabolism field.Int32
	MuscleMass      field.Float64
	ProteinRate     field.Float64
	VisceralFat     field.Float64
	Value           field.String
	ZoneOffsetInSec field.Int32
	ZoneName        field.String
	TimeIn0Tz       field.Int32
	IsUpload        field.Int32
	IsDeleted       field.Int32

	fieldMap map[string]field.Expr
}

func (w weightItemRecord) Table(newTableName string) *weightItemRecord {
	w.weightItemRecordDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w weightItemRecord) As(alias string) *weightItemRecord {
	w.weightItemRecordDo.DO = *(w.weightItemRecordDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *weightItemRecord) updateTableName(table string) *weightItemRecord {
	w.ALL = field.NewAsterisk(table)
	w.Key = field.NewString(table, "key")
	w.Sid = field.NewString(table, "sid")
	w.Time = field.NewInt32(table, "time")
	w.ZeroTimeIn0Tz = field.NewInt32(table, "zeroTimeIn0Tz")
	w.Weight = field.NewFloat64(table, "weight")
	w.Bmi = field.NewFloat64(table, "bmi")
	w.BodyFatRate = field.NewFloat64(table, "bodyFatRate")
	w.MoistureRate = field.NewFloat64(table, "moistureRate")
	w.BoneMass = field.NewFloat64(table, "boneMass")
	w.BasalMetabolism = field.NewInt32(table, "basalMetabolism")
	w.MuscleMass = field.NewFloat64(table, "muscleMass")
	w.ProteinRate = field.NewFloat64(table, "proteinRate")
	w.VisceralFat = field.NewFloat64(table, "visceralFat")
	w.Value = field.NewString(table, "value")
	w.ZoneOffsetInSec = field.NewInt32(table, "zoneOffsetInSec")
	w.ZoneName = field.NewString(table, "zoneName")
	w.TimeIn0Tz = field.NewInt32(table, "timeIn0Tz")
	w.IsUpload = field.NewInt32(table, "isUpload")
	w.IsDeleted = field.NewInt32(table, "isDeleted")

	w.fillFieldMap()

	return w
}

func (w *weightItemRecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *weightItemRecord) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 19)
	w.fieldMap["key"] = w.Key
	w.fieldMap["sid"] = w.Sid
	w.fieldMap["time"] = w.Time
	w.fieldMap["zeroTimeIn0Tz"] = w.ZeroTimeIn0Tz
	w.fieldMap["weight"] = w.Weight
	w.fieldMap["bmi"] = w.Bmi
	w.fieldMap["bodyFatRate"] = w.BodyFatRate
	w.fieldMap["moistureRate"] = w.MoistureRate
	w.fieldMap["boneMass"] = w.BoneMass
	w.fieldMap["basalMetabolism"] = w.BasalMetabolism
	w.fieldMap["muscleMass"] = w.MuscleMass
	w.fieldMap["proteinRate"] = w.ProteinRate
	w.fieldMap["visceralFat"] = w.VisceralFat
	w.fieldMap["value"] = w.Value
	w.fieldMap["zoneOffsetInSec"] = w.ZoneOffsetInSec
	w.fieldMap["zoneName"] = w.ZoneName
	w.fieldMap["timeIn0Tz"] = w.TimeIn0Tz
	w.fieldMap["isUpload"] = w.IsUpload
	w.fieldMap["isDeleted"] = w.IsDeleted
}

func (w weightItemRecord) clone(db *gorm.DB) weightItemRecord {
	w.weightItemRecordDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w weightItemRecord) replaceDB(db *gorm.DB) weightItemRecord {
	w.weightItemRecordDo.ReplaceDB(db)
	return w
}

type weightItemRecordDo struct{ gen.DO }

type IWeightItemRecordDo interface {
	gen.SubQuery
	Debug() IWeightItemRecordDo
	WithContext(ctx context.Context) IWeightItemRecordDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWeightItemRecordDo
	WriteDB() IWeightItemRecordDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWeightItemRecordDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWeightItemRecordDo
	Not(conds ...gen.Condition) IWeightItemRecordDo
	Or(conds ...gen.Condition) IWeightItemRecordDo
	Select(conds ...field.Expr) IWeightItemRecordDo
	Where(conds ...gen.Condition) IWeightItemRecordDo
	Order(conds ...field.Expr) IWeightItemRecordDo
	Distinct(cols ...field.Expr) IWeightItemRecordDo
	Omit(cols ...field.Expr) IWeightItemRecordDo
	Join(table schema.Tabler, on ...field.Expr) IWeightItemRecordDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWeightItemRecordDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWeightItemRecordDo
	Group(cols ...field.Expr) IWeightItemRecordDo
	Having(conds ...gen.Condition) IWeightItemRecordDo
	Limit(limit int) IWeightItemRecordDo
	Offset(offset int) IWeightItemRecordDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWeightItemRecordDo
	Unscoped() IWeightItemRecordDo
	Create(values ...*model.WeightItemRecord) error
	CreateInBatches(values []*model.WeightItemRecord, batchSize int) error
	Save(values ...*model.WeightItemRecord) error
	First() (*model.WeightItemRecord, error)
	Take() (*model.WeightItemRecord, error)
	Last() (*model.WeightItemRecord, error)
	Find() ([]*model.WeightItemRecord, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WeightItemRecord, err error)
	FindInBatches(result *[]*model.WeightItemRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WeightItemRecord) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWeightItemRecordDo
	Assign(attrs ...field.AssignExpr) IWeightItemRecordDo
	Joins(fields ...field.RelationField) IWeightItemRecordDo
	Preload(fields ...field.RelationField) IWeightItemRecordDo
	FirstOrInit() (*model.WeightItemRecord, error)
	FirstOrCreate() (*model.WeightItemRecord, error)
	FindByPage(offset int, limit int) (result []*model.WeightItemRecord, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWeightItemRecordDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w weightItemRecordDo) Debug() IWeightItemRecordDo {
	return w.withDO(w.DO.Debug())
}

func (w weightItemRecordDo) WithContext(ctx context.Context) IWeightItemRecordDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w weightItemRecordDo) ReadDB() IWeightItemRecordDo {
	return w.Clauses(dbresolver.Read)
}

func (w weightItemRecordDo) WriteDB() IWeightItemRecordDo {
	return w.Clauses(dbresolver.Write)
}

func (w weightItemRecordDo) Session(config *gorm.Session) IWeightItemRecordDo {
	return w.withDO(w.DO.Session(config))
}

func (w weightItemRecordDo) Clauses(conds ...clause.Expression) IWeightItemRecordDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w weightItemRecordDo) Returning(value interface{}, columns ...string) IWeightItemRecordDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w weightItemRecordDo) Not(conds ...gen.Condition) IWeightItemRecordDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w weightItemRecordDo) Or(conds ...gen.Condition) IWeightItemRecordDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w weightItemRecordDo) Select(conds ...field.Expr) IWeightItemRecordDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w weightItemRecordDo) Where(conds ...gen.Condition) IWeightItemRecordDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w weightItemRecordDo) Order(conds ...field.Expr) IWeightItemRecordDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w weightItemRecordDo) Distinct(cols ...field.Expr) IWeightItemRecordDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w weightItemRecordDo) Omit(cols ...field.Expr) IWeightItemRecordDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w weightItemRecordDo) Join(table schema.Tabler, on ...field.Expr) IWeightItemRecordDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w weightItemRecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWeightItemRecordDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w weightItemRecordDo) RightJoin(table schema.Tabler, on ...field.Expr) IWeightItemRecordDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w weightItemRecordDo) Group(cols ...field.Expr) IWeightItemRecordDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w weightItemRecordDo) Having(conds ...gen.Condition) IWeightItemRecordDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w weightItemRecordDo) Limit(limit int) IWeightItemRecordDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w weightItemRecordDo) Offset(offset int) IWeightItemRecordDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w weightItemRecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWeightItemRecordDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w weightItemRecordDo) Unscoped() IWeightItemRecordDo {
	return w.withDO(w.DO.Unscoped())
}

func (w weightItemRecordDo) Create(values ...*model.WeightItemRecord) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w weightItemRecordDo) CreateInBatches(values []*model.WeightItemRecord, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w weightItemRecordDo) Save(values ...*model.WeightItemRecord) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w weightItemRecordDo) First() (*model.WeightItemRecord, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WeightItemRecord), nil
	}
}

func (w weightItemRecordDo) Take() (*model.WeightItemRecord, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WeightItemRecord), nil
	}
}

func (w weightItemRecordDo) Last() (*model.WeightItemRecord, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WeightItemRecord), nil
	}
}

func (w weightItemRecordDo) Find() ([]*model.WeightItemRecord, error) {
	result, err := w.DO.Find()
	return result.([]*model.WeightItemRecord), err
}

func (w weightItemRecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WeightItemRecord, err error) {
	buf := make([]*model.WeightItemRecord, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w weightItemRecordDo) FindInBatches(result *[]*model.WeightItemRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w weightItemRecordDo) Attrs(attrs ...field.AssignExpr) IWeightItemRecordDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w weightItemRecordDo) Assign(attrs ...field.AssignExpr) IWeightItemRecordDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w weightItemRecordDo) Joins(fields ...field.RelationField) IWeightItemRecordDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w weightItemRecordDo) Preload(fields ...field.RelationField) IWeightItemRecordDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w weightItemRecordDo) FirstOrInit() (*model.WeightItemRecord, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WeightItemRecord), nil
	}
}

func (w weightItemRecordDo) FirstOrCreate() (*model.WeightItemRecord, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WeightItemRecord), nil
	}
}

func (w weightItemRecordDo) FindByPage(offset int, limit int) (result []*model.WeightItemRecord, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w weightItemRecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w weightItemRecordDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w weightItemRecordDo) Delete(models ...*model.WeightItemRecord) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *weightItemRecordDo) withDO(do gen.Dao) *weightItemRecordDo {
	w.DO = *do.(*gen.DO)
	return w
}
