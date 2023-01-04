package log_test

import (
	"fmt"
	"skarner2016/gin-api-starter/packages/config"
	"skarner2016/gin-api-starter/packages/log"
	"sync"
	"testing"
	"time"
)

func TestSetUp(t *testing.T) {
	config.Setup()

	log.Setup()

	// logger := log.GetLogger(log.InstanceDefault)
	// fmt.Println(fmt.Sprintf("%+v", logger))
	log.GetLogger(log.InstanceDefault).Debug(time.Now().String())

	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go WriteLog(wg, i)
	}

	wg.Wait()
	fmt.Println("end")
}

func WriteLog(wg *sync.WaitGroup, i int) {
	for j := 0; j < 1000; j++ {
		log.GetLogger(log.InstanceDefault).Debugf("%d-%d-%s", i, j, time.Now().Local().String())
		time.Sleep(50 * time.Millisecond)
	}
	wg.Done()
}
