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

func newSleepSegment(db *gorm.DB, opts ...gen.DOOption) sleepSegment {
	_sleepSegment := sleepSegment{}

	_sleepSegment.sleepSegmentDo.UseDB(db, opts...)
	_sleepSegment.sleepSegmentDo.UseModel(&model.SleepSegment{})

	tableName := _sleepSegment.sleepSegmentDo.TableName()
	_sleepSegment.ALL = field.NewAsterisk(tableName)
	_sleepSegment.Key = field.NewString(tableName, "key")
	_sleepSegment.Sid = field.NewString(tableName, "sid")
	_sleepSegment.Time = field.NewInt32(tableName, "time")
	_sleepSegment.IsComplete = field.NewInt32(tableName, "isComplete")
	_sleepSegment.Value = field.NewString(tableName, "value")
	_sleepSegment.ZoneOffsetInSec = field.NewInt32(tableName, "zoneOffsetInSec")
	_sleepSegment.ZoneName = field.NewString(tableName, "zoneName")
	_sleepSegment.TimeIn0Tz = field.NewInt32(tableName, "timeIn0Tz")
	_sleepSegment.IsUpload = field.NewInt32(tableName, "isUpload")
	_sleepSegment.IsDeleted = field.NewInt32(tableName, "isDeleted")

	_sleepSegment.fillFieldMap()

	return _sleepSegment
}

type sleepSegment struct {
	sleepSegmentDo

	ALL             field.Asterisk
	Key             field.String
	Sid             field.String
	Time            field.Int32
	IsComplete      field.Int32
	Value           field.String
	ZoneOffsetInSec field.Int32
	ZoneName        field.String
	TimeIn0Tz       field.Int32
	IsUpload        field.Int32
	IsDeleted       field.Int32

	fieldMap map[string]field.Expr
}

func (s sleepSegment) Table(newTableName string) *sleepSegment {
	s.sleepSegmentDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sleepSegment) As(alias string) *sleepSegment {
	s.sleepSegmentDo.DO = *(s.sleepSegmentDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sleepSegment) updateTableName(table string) *sleepSegment {
	s.ALL = field.NewAsterisk(table)
	s.Key = field.NewString(table, "key")
	s.Sid = field.NewString(table, "sid")
	s.Time = field.NewInt32(table, "time")
	s.IsComplete = field.NewInt32(table, "isComplete")
	s.Value = field.NewString(table, "value")
	s.ZoneOffsetInSec = field.NewInt32(table, "zoneOffsetInSec")
	s.ZoneName = field.NewString(table, "zoneName")
	s.TimeIn0Tz = field.NewInt32(table, "timeIn0Tz")
	s.IsUpload = field.NewInt32(table, "isUpload")
	s.IsDeleted = field.NewInt32(table, "isDeleted")

	s.fillFieldMap()

	return s
}

func (s *sleepSegment) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sleepSegment) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 10)
	s.fieldMap["key"] = s.Key
	s.fieldMap["sid"] = s.Sid
	s.fieldMap["time"] = s.Time
	s.fieldMap["isComplete"] = s.IsComplete
	s.fieldMap["value"] = s.Value
	s.fieldMap["zoneOffsetInSec"] = s.ZoneOffsetInSec
	s.fieldMap["zoneName"] = s.ZoneName
	s.fieldMap["timeIn0Tz"] = s.TimeIn0Tz
	s.fieldMap["isUpload"] = s.IsUpload
	s.fieldMap["isDeleted"] = s.IsDeleted
}

func (s sleepSegment) clone(db *gorm.DB) sleepSegment {
	s.sleepSegmentDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sleepSegment) replaceDB(db *gorm.DB) sleepSegment {
	s.sleepSegmentDo.ReplaceDB(db)
	return s
}

type sleepSegmentDo struct{ gen.DO }

type ISleepSegmentDo interface {
	gen.SubQuery
	Debug() ISleepSegmentDo
	WithContext(ctx context.Context) ISleepSegmentDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISleepSegmentDo
	WriteDB() ISleepSegmentDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISleepSegmentDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISleepSegmentDo
	Not(conds ...gen.Condition) ISleepSegmentDo
	Or(conds ...gen.Condition) ISleepSegmentDo
	Select(conds ...field.Expr) ISleepSegmentDo
	Where(conds ...gen.Condition) ISleepSegmentDo
	Order(conds ...field.Expr) ISleepSegmentDo
	Distinct(cols ...field.Expr) ISleepSegmentDo
	Omit(cols ...field.Expr) ISleepSegmentDo
	Join(table schema.Tabler, on ...field.Expr) ISleepSegmentDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISleepSegmentDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISleepSegmentDo
	Group(cols ...field.Expr) ISleepSegmentDo
	Having(conds ...gen.Condition) ISleepSegmentDo
	Limit(limit int) ISleepSegmentDo
	Offset(offset int) ISleepSegmentDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISleepSegmentDo
	Unscoped() ISleepSegmentDo
	Create(values ...*model.SleepSegment) error
	CreateInBatches(values []*model.SleepSegment, batchSize int) error
	Save(values ...*model.SleepSegment) error
	First() (*model.SleepSegment, error)
	Take() (*model.SleepSegment, error)
	Last() (*model.SleepSegment, error)
	Find() ([]*model.SleepSegment, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SleepSegment, err error)
	FindInBatches(result *[]*model.SleepSegment, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SleepSegment) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISleepSegmentDo
	Assign(attrs ...field.AssignExpr) ISleepSegmentDo
	Joins(fields ...field.RelationField) ISleepSegmentDo
	Preload(fields ...field.RelationField) ISleepSegmentDo
	FirstOrInit() (*model.SleepSegment, error)
	FirstOrCreate() (*model.SleepSegment, error)
	FindByPage(offset int, limit int) (result []*model.SleepSegment, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISleepSegmentDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sleepSegmentDo) Debug() ISleepSegmentDo {
	return s.withDO(s.DO.Debug())
}

func (s sleepSegmentDo) WithContext(ctx context.Context) ISleepSegmentDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sleepSegmentDo) ReadDB() ISleepSegmentDo {
	return s.Clauses(dbresolver.Read)
}

func (s sleepSegmentDo) WriteDB() ISleepSegmentDo {
	return s.Clauses(dbresolver.Write)
}

func (s sleepSegmentDo) Session(config *gorm.Session) ISleepSegmentDo {
	return s.withDO(s.DO.Session(config))
}

func (s sleepSegmentDo) Clauses(conds ...clause.Expression) ISleepSegmentDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sleepSegmentDo) Returning(value interface{}, columns ...string) ISleepSegmentDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sleepSegmentDo) Not(conds ...gen.Condition) ISleepSegmentDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sleepSegmentDo) Or(conds ...gen.Condition) ISleepSegmentDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sleepSegmentDo) Select(conds ...field.Expr) ISleepSegmentDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sleepSegmentDo) Where(conds ...gen.Condition) ISleepSegmentDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sleepSegmentDo) Order(conds ...field.Expr) ISleepSegmentDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sleepSegmentDo) Distinct(cols ...field.Expr) ISleepSegmentDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sleepSegmentDo) Omit(cols ...field.Expr) ISleepSegmentDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sleepSegmentDo) Join(table schema.Tabler, on ...field.Expr) ISleepSegmentDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sleepSegmentDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISleepSegmentDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sleepSegmentDo) RightJoin(table schema.Tabler, on ...field.Expr) ISleepSegmentDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sleepSegmentDo) Group(cols ...field.Expr) ISleepSegmentDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sleepSegmentDo) Having(conds ...gen.Condition) ISleepSegmentDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sleepSegmentDo) Limit(limit int) ISleepSegmentDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sleepSegmentDo) Offset(offset int) ISleepSegmentDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sleepSegmentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISleepSegmentDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sleepSegmentDo) Unscoped() ISleepSegmentDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sleepSegmentDo) Create(values ...*model.SleepSegment) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sleepSegmentDo) CreateInBatches(values []*model.SleepSegment, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sleepSegmentDo) Save(values ...*model.SleepSegment) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sleepSegmentDo) First() (*model.SleepSegment, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SleepSegment), nil
	}
}

func (s sleepSegmentDo) Take() (*model.SleepSegment, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SleepSegment), nil
	}
}

func (s sleepSegmentDo) Last() (*model.SleepSegment, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SleepSegment), nil
	}
}

func (s sleepSegmentDo) Find() ([]*model.SleepSegment, error) {
	result, err := s.DO.Find()
	return result.([]*model.SleepSegment), err
}

func (s sleepSegmentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SleepSegment, err error) {
	buf := make([]*model.SleepSegment, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sleepSegmentDo) FindInBatches(result *[]*model.SleepSegment, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sleepSegmentDo) Attrs(attrs ...field.AssignExpr) ISleepSegmentDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sleepSegmentDo) Assign(attrs ...field.AssignExpr) ISleepSegmentDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sleepSegmentDo) Joins(fields ...field.RelationField) ISleepSegmentDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sleepSegmentDo) Preload(fields ...field.RelationField) ISleepSegmentDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sleepSegmentDo) FirstOrInit() (*model.SleepSegment, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SleepSegment), nil
	}
}

func (s sleepSegmentDo) FirstOrCreate() (*model.SleepSegment, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SleepSegment), nil
	}
}

func (s sleepSegmentDo) FindByPage(offset int, limit int) (result []*model.SleepSegment, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sleepSegmentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sleepSegmentDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sleepSegmentDo) Delete(models ...*model.SleepSegment) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sleepSegmentDo) withDO(do gen.Dao) *sleepSegmentDo {
	s.DO = *do.(*gen.DO)
	return s
}