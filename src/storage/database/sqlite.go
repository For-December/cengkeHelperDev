package database

import (
	"cengkeHelperDev/src/constant/config"
	"cengkeHelperDev/src/dbmodels"
	"cengkeHelperDev/src/utils/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var Client *gorm.DB

func init() {
	var err error
	var cfg gorm.Config
	cfg = gorm.Config{
		PrepareStmt: true,
		Logger:      gormLogger.Default.LogMode(gormLogger.Info),
		//NamingStrategy: schema.NamingStrategy{
		//	TablePrefix: "test",
		//},
		ConnPool: nil,
	}
	// 连接到SQLite数据库
	if Client, err = gorm.Open(sqlite.Open("resources/test.db"), &cfg); err != nil {
		panic(err)
	}

	TableAutoMigrate()
}

func TableAutoMigrate() {
	if !config.EnvCfg.AutoMigrate {
		logger.Info("未启用迁移数据库")
		return
	}
	if err := Client.AutoMigrate(&dbmodels.TimeInfo{}, &dbmodels.CourseInfo{},
		&dbmodels.PostRecord{}, &dbmodels.CommentRecord{}, &dbmodels.StarRecord{},
	); err != nil {
		panic(err)
		return
	}

}
