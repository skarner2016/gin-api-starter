package mysql

import (
	"fmt"
	"skarner2016/gin-api-starter/packages/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var InstanceMap map[Instance]*gorm.DB

func Setup() {
	confMap := make(map[Instance]*DBConf, 0)
	if err := config.APPConfig.UnmarshalKey("mysql", &confMap); err != nil {
		panic("mysql setup: parse config err:" + err.Error())
	}

	mode := config.APPConfig.GetString("mode")

	InstanceMap = make(map[Instance]*gorm.DB, 0)
	for i, conf := range confMap {
		// dsn
		dsn := getDSN(conf)

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
		sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
		// 连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭
		sqlDB.SetMaxOpenConns(conf.MaxOpenConns)
		// 数据库连接课复用的最大时间
		sqlDB.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifetime) * time.Minute)

		if mode == "debug" {
			InstanceMap[i] = db.Debug()
		} else {
			// DB = db
			InstanceMap[i] = db
		}
	}
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

func GetDB(instance Instance) *gorm.DB {
	if _, ok := InstanceMap[instance]; !ok {
		Setup()
	}

	db, _ := InstanceMap[instance]

	return db
}
