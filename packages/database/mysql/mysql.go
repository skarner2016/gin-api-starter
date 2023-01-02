package mysql

import (
	"database/sql"
	"fmt"
	"skarner2016/gin-api-starter/packages/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type DBConf struct {
	Host            string
	Port            int64
	User            string
	Pass            string
	Database        string
	Charset         string
	Collation       string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
}

func Setup() *sql.DB {
	dbconf := new(DBConf)
	if err := config.APPConfig.UnmarshalKey("mysql.default", &dbconf); err != nil {
		panic("mysql setup: parse database config err:" + err.Error())
	}

	// dsn
	dsn := getDSN(dbconf)

	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("mysql setup: mysql open err:" + err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("mysql setup: mysql open err:" + err.Error())
	}

	// 数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(dbconf.MaxIdleConns)
	// 连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭
	sqlDB.SetMaxOpenConns(dbconf.MaxOpenConns)
	// 数据库连接课复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Duration(dbconf.ConnMaxLifetime) * time.Minute)

	mode := config.APPConfig.GetString("mode")
	if mode == "debug" {
		DB = db.Debug()
	} else {
		DB = db
	}

	return sqlDB
}

func getDSN(c *DBConf) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		c.User,
		c.Pass,
		c.Host,
		c.Port,
		c.Database,
		c.Charset,
	)
}

func GetDB() *gorm.DB {
	if DB == nil {
		Setup()
	}

	return DB
}
