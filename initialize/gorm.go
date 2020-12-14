package initialize

import (
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"os"
	"study-gorm/global"
	"study-gorm/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	return GormMysql()
}

func GormMysql() *gorm.DB {
	m := global.GvaConfig.Mysql
	dsn := m.Username + ":" + m.Password+"@tcp("+m.Path+")/"+m.Dbname+"?"+m.Config
	mysqlConfig := mysql.Config{
		DSN: dsn, // DSN data source name
		DefaultStringSize: 256, // string 类型字段的默认长度
		DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig(m.LogMode)); err != nil {
		global.GvaLog.Error("MySQL启动异常", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// 注册数据库表
func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.SysUser{},
		model.SysAuthority{},
		model.SysBaseMenu{},
		model.SysBaseMenuParameter{},
		model.JwtBlacklist{},
		model.SysWorkflow{},
		model.SysWorkflowStepInfo{},
		model.SysDictionary{},
		model.SysDictionaryDetail{},
		model.ExaFileUploadAndDownload{},
		model.ExaFile{},
		model.ExaFileChunk{},
		model.ExaSimpleUploader{},
		model.ExaCustomer{},
		model.SysOperationRecord{},
	)
	if err != nil {
		global.GvaLog.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.GvaLog.Info("register table success")
}

func gormConfig(mod bool) *gorm.Config {
	if mod {
		return &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	} else {
		return &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	}
}


