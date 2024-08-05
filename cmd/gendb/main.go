package main

import (
	"fmt"
	"github.com/Ocyss/xiaomiHA/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// https://gorm.io/zh_CN/gen/database_to_structs.html#Quick-Start
//type CommonMethod struct {
//	Time int32
//}
//
//func (m *CommonMethod) GetCsvRecords() ([]*CommonMethod, error) {
//	var conds []gen.Condition
//	if m.Time.Time != nil {
//		conds = append(conds, query.TrainingLoadRecord.Time.Gte(*h.Time.Time))
//	} else {
//		conds = append(conds, query.TrainingLoadRecord.Time.Between(*h.Time.StartTime, *h.Time.EndTime))
//	}
//	records, err := query.TrainingLoadRecord.Order(query.TrainingLoadRecord.Time.Desc()).Where(conds...).Find()
//	return records, err
//}

func main() {
	genModels(model.SummaryTableNames, model.SDb)
	// TODO: 后面会覆盖掉前面的 gen 函数
	genModels(model.DataTableNames, model.Db)
}

func genModels(tableNames []string, db *gorm.DB) {
	g := gen.NewGenerator(gen.Config{
		OutPath: "query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		// 如果你希望为可为null的字段生成属性为指针类型, 设置 FieldNullable 为 true
		FieldNullable: true,
		// 如果你希望在 `Create` API 中为字段分配默认值, 设置 FieldCoverable 为 true, 参考: https://gorm.io/docs/create.html#Default-Values
		FieldCoverable: true,
		// 如果你希望生成无符号整数类型字段, 设置 FieldSignable 为 true
		FieldSignable: true,
		// 如果你希望从数据库生成索引标记, 设置 FieldWithIndexTag 为 true
		FieldWithIndexTag: true,
		// 如果你希望从数据库生成类型标记, 设置 FieldWithTypeTag 为 true
		FieldWithTypeTag: true,
		// 如果你需要对查询代码进行单元测试, 设置 WithUnitTest 为 true
		WithUnitTest: false,
	})

	g.UseDB(db) // reuse your gorm db
	g.WithOpts(gen.FieldModify(func(field gen.Field) gen.Field {
		if field.Type == "int32" {
			field.Type = "int64"
		}
		return field
	}))
	// Generate basic type-safe DAO API for struct `model.User` following conventions
	table := make([]any, 0, len(tableNames))
	for _, v := range tableNames {
		table = append(table, g.GenerateModel(v))
	}
	g.ApplyBasic(
		// Generate structs from all tables of current database
		table...,
	)

	// Generate the code
	g.Execute()
}

func getData(tableNames []string, db *gorm.DB) {
	for _, v := range tableNames {
		data := make(map[string]any)
		err := db.Table(v).Limit(3).Find(&data).Error
		if err != nil {
			log.Errorln(err)
		} else {
			fmt.Println(v, data)
		}
	}
}

func getTable(db *gorm.DB) {
	tables, err := db.Migrator().GetTables()
	if err != nil {
		panic(err)
	}
	for _, table := range tables {
		fmt.Println(table)
	}
}
