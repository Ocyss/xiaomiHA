package model

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DataTableNames = []string{
		"sport_report",          // 运动报告
		"sleep_segment",         // 睡眠段
		"hr_record",             // 心率记录
		"step_record",           // 步数记录
		"calorie_record",        // 卡路里记录
		"stand_record",          // 站立记录
		"mh_strength_record",    // 肌肉力量记录
		"stress_record",         // 压力记录
		"spo2_record",           // 血氧记录
		"hearing_health_record", // 听力健康记录
		"weight_item_record",    // 体重记录
		"pai_record",            // PAI记录
		"training_load_record",  // 训练负荷记录
		"vitality_record",       // 活力记录

		// "menstruation_record",           // 月经记录
		// "energy_record",                 // 能量记录（为空）
		// "v02max_record",                 // 最大摄氧量记录（为空）
		// "blood_pressure_record",         // 血压记录（为空）
		// "temperature_record",            // 温度记录（为空）
		// "running_ability_record",        // 跑步能力记录（为空）
		// "ecg_medical_record",            // 心电图医疗记录（为空）
		// "blood_pressure_medical_record", // 血压医疗记录（为空）
		// "blood_sugar_record",            // 血糖记录（为空）
		// "physical_fitness_status",       // 身体健康状况（为空）
		// "lactate_threshold",             // 乳酸阈值（为空）
	}

	SummaryTableNames = []string{
		"rainbow",       // 活力指标 （彩虹目标条）
		"daily_report",  // 每日活动
		"sport_summary", // 运动总结
	}

	Db, SDb = InitDb()
)

func InitDb() (dataDb *gorm.DB, summaryDb *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("temp/fitness_data"), &gorm.Config{})

	if err != nil {
		log.Fatalln("err init fitness_data db")
	}

	db2, err := gorm.Open(sqlite.Open("temp/fitness_summary"), &gorm.Config{})
	if err != nil {
		log.Fatalln("init fitness_summary db")
	}

	return db, db2
}
