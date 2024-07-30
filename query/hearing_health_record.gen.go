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

func newHearingHealthRecord(db *gorm.DB, opts ...gen.DOOption) hearingHealthRecord {
	_hearingHealthRecord := hearingHealthRecord{}

	_hearingHealthRecord.hearingHealthRecordDo.UseDB(db, opts...)
	_hearingHealthRecord.hearingHealthRecordDo.UseModel(&model.HearingHealthRecord{})

	tableName := _hearingHealthRecord.hearingHealthRecordDo.TableName()
	_hearingHealthRecord.ALL = field.NewAsterisk(tableName)
	_hearingHealthRecord.Key = field.NewString(tableName, "key")
	_hearingHealthRecord.Sid = field.NewString(tableName, "sid")
	_hearingHealthRecord.Time = field.NewInt32(tableName, "time")
	_hearingHealthRecord.Value = field.NewString(tableName, "value")
	_hearingHealthRecord.ZoneOffsetInSec = field.NewInt32(tableName, "zoneOffsetInSec")
	_hearingHealthRecord.ZoneName = field.NewString(tableName, "zoneName")
	_hearingHealthRecord.TimeIn0Tz = field.NewInt32(tableName, "timeIn0Tz")
	_hearingHealthRecord.IsUpload = field.NewInt32(tableName, "isUpload")
	_hearingHealthRecord.IsDeleted = field.NewInt32(tableName, "isDeleted")

	_hearingHealthRecord.fillFieldMap()

	return _hearingHealthRecord
}

type hearingHealthRecord struct {
	hearingHealthRecordDo

	ALL             field.Asterisk
	Key             field.String
	Sid             field.String
	Time            field.Int32
	Value           field.String
	ZoneOffsetInSec field.Int32
	ZoneName        field.String
	TimeIn0Tz       field.Int32
	IsUpload        field.Int32
	IsDeleted       field.Int32

	fieldMap map[string]field.Expr
}

func (h hearingHealthRecord) Table(newTableName string) *hearingHealthRecord {
	h.hearingHealthRecordDo.UseTable(newTableName)
	return h.updateTableName(newTableName)
}

func (h hearingHealthRecord) As(alias string) *hearingHealthRecord {
	h.hearingHealthRecordDo.DO = *(h.hearingHealthRecordDo.As(alias).(*gen.DO))
	return h.updateTableName(alias)
}

func (h *hearingHealthRecord) updateTableName(table string) *hearingHealthRecord {
	h.ALL = field.NewAsterisk(table)
	h.Key = field.NewString(table, "key")
	h.Sid = field.NewString(table, "sid")
	h.Time = field.NewInt32(table, "time")
	h.Value = field.NewString(table, "value")
	h.ZoneOffsetInSec = field.NewInt32(table, "zoneOffsetInSec")
	h.ZoneName = field.NewString(table, "zoneName")
	h.TimeIn0Tz = field.NewInt32(table, "timeIn0Tz")
	h.IsUpload = field.NewInt32(table, "isUpload")
	h.IsDeleted = field.NewInt32(table, "isDeleted")

	h.fillFieldMap()

	return h
}

func (h *hearingHealthRecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := h.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (h *hearingHealthRecord) fillFieldMap() {
	h.fieldMap = make(map[string]field.Expr, 9)
	h.fieldMap["key"] = h.Key
	h.fieldMap["sid"] = h.Sid
	h.fieldMap["time"] = h.Time
	h.fieldMap["value"] = h.Value
	h.fieldMap["zoneOffsetInSec"] = h.ZoneOffsetInSec
	h.fieldMap["zoneName"] = h.ZoneName
	h.fieldMap["timeIn0Tz"] = h.TimeIn0Tz
	h.fieldMap["isUpload"] = h.IsUpload
	h.fieldMap["isDeleted"] = h.IsDeleted
}

func (h hearingHealthRecord) clone(db *gorm.DB) hearingHealthRecord {
	h.hearingHealthRecordDo.ReplaceConnPool(db.Statement.ConnPool)
	return h
}

func (h hearingHealthRecord) replaceDB(db *gorm.DB) hearingHealthRecord {
	h.hearingHealthRecordDo.ReplaceDB(db)
	return h
}

type hearingHealthRecordDo struct{ gen.DO }

type IHearingHealthRecordDo interface {
	gen.SubQuery
	Debug() IHearingHealthRecordDo
	WithContext(ctx context.Context) IHearingHealthRecordDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IHearingHealthRecordDo
	WriteDB() IHearingHealthRecordDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IHearingHealthRecordDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IHearingHealthRecordDo
	Not(conds ...gen.Condition) IHearingHealthRecordDo
	Or(conds ...gen.Condition) IHearingHealthRecordDo
	Select(conds ...field.Expr) IHearingHealthRecordDo
	Where(conds ...gen.Condition) IHearingHealthRecordDo
	Order(conds ...field.Expr) IHearingHealthRecordDo
	Distinct(cols ...field.Expr) IHearingHealthRecordDo
	Omit(cols ...field.Expr) IHearingHealthRecordDo
	Join(table schema.Tabler, on ...field.Expr) IHearingHealthRecordDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IHearingHealthRecordDo
	RightJoin(table schema.Tabler, on ...field.Expr) IHearingHealthRecordDo
	Group(cols ...field.Expr) IHearingHealthRecordDo
	Having(conds ...gen.Condition) IHearingHealthRecordDo
	Limit(limit int) IHearingHealthRecordDo
	Offset(offset int) IHearingHealthRecordDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IHearingHealthRecordDo
	Unscoped() IHearingHealthRecordDo
	Create(values ...*model.HearingHealthRecord) error
	CreateInBatches(values []*model.HearingHealthRecord, batchSize int) error
	Save(values ...*model.HearingHealthRecord) error
	First() (*model.HearingHealthRecord, error)
	Take() (*model.HearingHealthRecord, error)
	Last() (*model.HearingHealthRecord, error)
	Find() ([]*model.HearingHealthRecord, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HearingHealthRecord, err error)
	FindInBatches(result *[]*model.HearingHealthRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.HearingHealthRecord) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IHearingHealthRecordDo
	Assign(attrs ...field.AssignExpr) IHearingHealthRecordDo
	Joins(fields ...field.RelationField) IHearingHealthRecordDo
	Preload(fields ...field.RelationField) IHearingHealthRecordDo
	FirstOrInit() (*model.HearingHealthRecord, error)
	FirstOrCreate() (*model.HearingHealthRecord, error)
	FindByPage(offset int, limit int) (result []*model.HearingHealthRecord, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IHearingHealthRecordDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (h hearingHealthRecordDo) Debug() IHearingHealthRecordDo {
	return h.withDO(h.DO.Debug())
}

func (h hearingHealthRecordDo) WithContext(ctx context.Context) IHearingHealthRecordDo {
	return h.withDO(h.DO.WithContext(ctx))
}

func (h hearingHealthRecordDo) ReadDB() IHearingHealthRecordDo {
	return h.Clauses(dbresolver.Read)
}

func (h hearingHealthRecordDo) WriteDB() IHearingHealthRecordDo {
	return h.Clauses(dbresolver.Write)
}

func (h hearingHealthRecordDo) Session(config *gorm.Session) IHearingHealthRecordDo {
	return h.withDO(h.DO.Session(config))
}

func (h hearingHealthRecordDo) Clauses(conds ...clause.Expression) IHearingHealthRecordDo {
	return h.withDO(h.DO.Clauses(conds...))
}

func (h hearingHealthRecordDo) Returning(value interface{}, columns ...string) IHearingHealthRecordDo {
	return h.withDO(h.DO.Returning(value, columns...))
}

func (h hearingHealthRecordDo) Not(conds ...gen.Condition) IHearingHealthRecordDo {
	return h.withDO(h.DO.Not(conds...))
}

func (h hearingHealthRecordDo) Or(conds ...gen.Condition) IHearingHealthRecordDo {
	return h.withDO(h.DO.Or(conds...))
}

func (h hearingHealthRecordDo) Select(conds ...field.Expr) IHearingHealthRecordDo {
	return h.withDO(h.DO.Select(conds...))
}

func (h hearingHealthRecordDo) Where(conds ...gen.Condition) IHearingHealthRecordDo {
	return h.withDO(h.DO.Where(conds...))
}

func (h hearingHealthRecordDo) Order(conds ...field.Expr) IHearingHealthRecordDo {
	return h.withDO(h.DO.Order(conds...))
}

func (h hearingHealthRecordDo) Distinct(cols ...field.Expr) IHearingHealthRecordDo {
	return h.withDO(h.DO.Distinct(cols...))
}

func (h hearingHealthRecordDo) Omit(cols ...field.Expr) IHearingHealthRecordDo {
	return h.withDO(h.DO.Omit(cols...))
}

func (h hearingHealthRecordDo) Join(table schema.Tabler, on ...field.Expr) IHearingHealthRecordDo {
	return h.withDO(h.DO.Join(table, on...))
}

func (h hearingHealthRecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) IHearingHealthRecordDo {
	return h.withDO(h.DO.LeftJoin(table, on...))
}

func (h hearingHealthRecordDo) RightJoin(table schema.Tabler, on ...field.Expr) IHearingHealthRecordDo {
	return h.withDO(h.DO.RightJoin(table, on...))
}

func (h hearingHealthRecordDo) Group(cols ...field.Expr) IHearingHealthRecordDo {
	return h.withDO(h.DO.Group(cols...))
}

func (h hearingHealthRecordDo) Having(conds ...gen.Condition) IHearingHealthRecordDo {
	return h.withDO(h.DO.Having(conds...))
}

func (h hearingHealthRecordDo) Limit(limit int) IHearingHealthRecordDo {
	return h.withDO(h.DO.Limit(limit))
}

func (h hearingHealthRecordDo) Offset(offset int) IHearingHealthRecordDo {
	return h.withDO(h.DO.Offset(offset))
}

func (h hearingHealthRecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IHearingHealthRecordDo {
	return h.withDO(h.DO.Scopes(funcs...))
}

func (h hearingHealthRecordDo) Unscoped() IHearingHealthRecordDo {
	return h.withDO(h.DO.Unscoped())
}

func (h hearingHealthRecordDo) Create(values ...*model.HearingHealthRecord) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Create(values)
}

func (h hearingHealthRecordDo) CreateInBatches(values []*model.HearingHealthRecord, batchSize int) error {
	return h.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (h hearingHealthRecordDo) Save(values ...*model.HearingHealthRecord) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Save(values)
}

func (h hearingHealthRecordDo) First() (*model.HearingHealthRecord, error) {
	if result, err := h.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.HearingHealthRecord), nil
	}
}

func (h hearingHealthRecordDo) Take() (*model.HearingHealthRecord, error) {
	if result, err := h.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.HearingHealthRecord), nil
	}
}

func (h hearingHealthRecordDo) Last() (*model.HearingHealthRecord, error) {
	if result, err := h.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.HearingHealthRecord), nil
	}
}

func (h hearingHealthRecordDo) Find() ([]*model.HearingHealthRecord, error) {
	result, err := h.DO.Find()
	return result.([]*model.HearingHealthRecord), err
}

func (h hearingHealthRecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HearingHealthRecord, err error) {
	buf := make([]*model.HearingHealthRecord, 0, batchSize)
	err = h.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (h hearingHealthRecordDo) FindInBatches(result *[]*model.HearingHealthRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return h.DO.FindInBatches(result, batchSize, fc)
}

func (h hearingHealthRecordDo) Attrs(attrs ...field.AssignExpr) IHearingHealthRecordDo {
	return h.withDO(h.DO.Attrs(attrs...))
}

func (h hearingHealthRecordDo) Assign(attrs ...field.AssignExpr) IHearingHealthRecordDo {
	return h.withDO(h.DO.Assign(attrs...))
}

func (h hearingHealthRecordDo) Joins(fields ...field.RelationField) IHearingHealthRecordDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Joins(_f))
	}
	return &h
}

func (h hearingHealthRecordDo) Preload(fields ...field.RelationField) IHearingHealthRecordDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Preload(_f))
	}
	return &h
}

func (h hearingHealthRecordDo) FirstOrInit() (*model.HearingHealthRecord, error) {
	if result, err := h.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.HearingHealthRecord), nil
	}
}

func (h hearingHealthRecordDo) FirstOrCreate() (*model.HearingHealthRecord, error) {
	if result, err := h.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.HearingHealthRecord), nil
	}
}

func (h hearingHealthRecordDo) FindByPage(offset int, limit int) (result []*model.HearingHealthRecord, count int64, err error) {
	result, err = h.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = h.Offset(-1).Limit(-1).Count()
	return
}

func (h hearingHealthRecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = h.Count()
	if err != nil {
		return
	}

	err = h.Offset(offset).Limit(limit).Scan(result)
	return
}

func (h hearingHealthRecordDo) Scan(result interface{}) (err error) {
	return h.DO.Scan(result)
}

func (h hearingHealthRecordDo) Delete(models ...*model.HearingHealthRecord) (result gen.ResultInfo, err error) {
	return h.DO.Delete(models)
}

func (h *hearingHealthRecordDo) withDO(do gen.Dao) *hearingHealthRecordDo {
	h.DO = *do.(*gen.DO)
	return h
}