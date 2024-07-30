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

func newRainbow(db *gorm.DB, opts ...gen.DOOption) rainbow {
	_rainbow := rainbow{}

	_rainbow.rainbowDo.UseDB(db, opts...)
	_rainbow.rainbowDo.UseModel(&model.Rainbow{})

	tableName := _rainbow.rainbowDo.TableName()
	_rainbow.ALL = field.NewAsterisk(tableName)
	_rainbow.TimeInZero = field.NewInt32(tableName, "timeInZero")
	_rainbow.Value = field.NewString(tableName, "value")
	_rainbow.IsUpload = field.NewInt32(tableName, "isUpload")

	_rainbow.fillFieldMap()

	return _rainbow
}

type rainbow struct {
	rainbowDo

	ALL        field.Asterisk
	TimeInZero field.Int32
	Value      field.String
	IsUpload   field.Int32

	fieldMap map[string]field.Expr
}

func (r rainbow) Table(newTableName string) *rainbow {
	r.rainbowDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r rainbow) As(alias string) *rainbow {
	r.rainbowDo.DO = *(r.rainbowDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *rainbow) updateTableName(table string) *rainbow {
	r.ALL = field.NewAsterisk(table)
	r.TimeInZero = field.NewInt32(table, "timeInZero")
	r.Value = field.NewString(table, "value")
	r.IsUpload = field.NewInt32(table, "isUpload")

	r.fillFieldMap()

	return r
}

func (r *rainbow) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *rainbow) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 3)
	r.fieldMap["timeInZero"] = r.TimeInZero
	r.fieldMap["value"] = r.Value
	r.fieldMap["isUpload"] = r.IsUpload
}

func (r rainbow) clone(db *gorm.DB) rainbow {
	r.rainbowDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r rainbow) replaceDB(db *gorm.DB) rainbow {
	r.rainbowDo.ReplaceDB(db)
	return r
}

type rainbowDo struct{ gen.DO }

type IRainbowDo interface {
	gen.SubQuery
	Debug() IRainbowDo
	WithContext(ctx context.Context) IRainbowDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IRainbowDo
	WriteDB() IRainbowDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IRainbowDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRainbowDo
	Not(conds ...gen.Condition) IRainbowDo
	Or(conds ...gen.Condition) IRainbowDo
	Select(conds ...field.Expr) IRainbowDo
	Where(conds ...gen.Condition) IRainbowDo
	Order(conds ...field.Expr) IRainbowDo
	Distinct(cols ...field.Expr) IRainbowDo
	Omit(cols ...field.Expr) IRainbowDo
	Join(table schema.Tabler, on ...field.Expr) IRainbowDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRainbowDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRainbowDo
	Group(cols ...field.Expr) IRainbowDo
	Having(conds ...gen.Condition) IRainbowDo
	Limit(limit int) IRainbowDo
	Offset(offset int) IRainbowDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRainbowDo
	Unscoped() IRainbowDo
	Create(values ...*model.Rainbow) error
	CreateInBatches(values []*model.Rainbow, batchSize int) error
	Save(values ...*model.Rainbow) error
	First() (*model.Rainbow, error)
	Take() (*model.Rainbow, error)
	Last() (*model.Rainbow, error)
	Find() ([]*model.Rainbow, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Rainbow, err error)
	FindInBatches(result *[]*model.Rainbow, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Rainbow) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRainbowDo
	Assign(attrs ...field.AssignExpr) IRainbowDo
	Joins(fields ...field.RelationField) IRainbowDo
	Preload(fields ...field.RelationField) IRainbowDo
	FirstOrInit() (*model.Rainbow, error)
	FirstOrCreate() (*model.Rainbow, error)
	FindByPage(offset int, limit int) (result []*model.Rainbow, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRainbowDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r rainbowDo) Debug() IRainbowDo {
	return r.withDO(r.DO.Debug())
}

func (r rainbowDo) WithContext(ctx context.Context) IRainbowDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r rainbowDo) ReadDB() IRainbowDo {
	return r.Clauses(dbresolver.Read)
}

func (r rainbowDo) WriteDB() IRainbowDo {
	return r.Clauses(dbresolver.Write)
}

func (r rainbowDo) Session(config *gorm.Session) IRainbowDo {
	return r.withDO(r.DO.Session(config))
}

func (r rainbowDo) Clauses(conds ...clause.Expression) IRainbowDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r rainbowDo) Returning(value interface{}, columns ...string) IRainbowDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r rainbowDo) Not(conds ...gen.Condition) IRainbowDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r rainbowDo) Or(conds ...gen.Condition) IRainbowDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r rainbowDo) Select(conds ...field.Expr) IRainbowDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r rainbowDo) Where(conds ...gen.Condition) IRainbowDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r rainbowDo) Order(conds ...field.Expr) IRainbowDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r rainbowDo) Distinct(cols ...field.Expr) IRainbowDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r rainbowDo) Omit(cols ...field.Expr) IRainbowDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r rainbowDo) Join(table schema.Tabler, on ...field.Expr) IRainbowDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r rainbowDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRainbowDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r rainbowDo) RightJoin(table schema.Tabler, on ...field.Expr) IRainbowDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r rainbowDo) Group(cols ...field.Expr) IRainbowDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r rainbowDo) Having(conds ...gen.Condition) IRainbowDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r rainbowDo) Limit(limit int) IRainbowDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r rainbowDo) Offset(offset int) IRainbowDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r rainbowDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRainbowDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r rainbowDo) Unscoped() IRainbowDo {
	return r.withDO(r.DO.Unscoped())
}

func (r rainbowDo) Create(values ...*model.Rainbow) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r rainbowDo) CreateInBatches(values []*model.Rainbow, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r rainbowDo) Save(values ...*model.Rainbow) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r rainbowDo) First() (*model.Rainbow, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Rainbow), nil
	}
}

func (r rainbowDo) Take() (*model.Rainbow, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Rainbow), nil
	}
}

func (r rainbowDo) Last() (*model.Rainbow, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Rainbow), nil
	}
}

func (r rainbowDo) Find() ([]*model.Rainbow, error) {
	result, err := r.DO.Find()
	return result.([]*model.Rainbow), err
}

func (r rainbowDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Rainbow, err error) {
	buf := make([]*model.Rainbow, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r rainbowDo) FindInBatches(result *[]*model.Rainbow, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r rainbowDo) Attrs(attrs ...field.AssignExpr) IRainbowDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r rainbowDo) Assign(attrs ...field.AssignExpr) IRainbowDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r rainbowDo) Joins(fields ...field.RelationField) IRainbowDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r rainbowDo) Preload(fields ...field.RelationField) IRainbowDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r rainbowDo) FirstOrInit() (*model.Rainbow, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Rainbow), nil
	}
}

func (r rainbowDo) FirstOrCreate() (*model.Rainbow, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Rainbow), nil
	}
}

func (r rainbowDo) FindByPage(offset int, limit int) (result []*model.Rainbow, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r rainbowDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r rainbowDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r rainbowDo) Delete(models ...*model.Rainbow) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *rainbowDo) withDO(do gen.Dao) *rainbowDo {
	r.DO = *do.(*gen.DO)
	return r
}