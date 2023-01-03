package mysql_test

import (
	"fmt"
	"skarner2016/gin-api-starter/models"
	"skarner2016/gin-api-starter/packages/config"
	"skarner2016/gin-api-starter/packages/database/mysql"
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	ID   int64
	Name string
}

func TestConnectDB(t *testing.T) {
	config.Setup()

	db := mysql.GetDB(mysql.InstanceDefault)

	user := &User{}
	err := db.Limit(1).Find(&user).Error

	assert.Equal(t, err, nil)
	fmt.Println(user)

	fmt.Println("test")
}

func TestMigrat(t *testing.T) {
	config.Setup()

	db := mysql.GetDB(mysql.InstanceDefault)

	db.AutoMigrate(
		&models.User{},
	)
}
