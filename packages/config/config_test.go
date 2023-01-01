package config_test

import (
	"fmt"
	"skarner2016/gin-api-starter/packages/config"
	"testing"
)

func TestGetAddr(t *testing.T) {

	config.SetupConfig()

	addr := config.APPConfig.Get("addr")

	fmt.Println(addr)
}
