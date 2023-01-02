package models_test

import (
	"fmt"
	"skarner2016/gin-api-starter/models"
	"skarner2016/gin-api-starter/packages/config"
	"skarner2016/gin-api-starter/packages/database/mysql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetByID(t *testing.T) {
	config.Setup()
	user := &models.User{}

	err := mysql.GetDB().Where("id = ?", 1).Find(&user).Error
	assert.Equal(t, err, nil)

	fmt.Println(user)
}

func TestInserUser(t *testing.T) {
	config.Setup()
	user := &models.User{
		Name:      "lisi",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := mysql.GetDB().Create(&user).Error
	assert.Equal(t, err, nil)

	fmt.Println(user)
}
