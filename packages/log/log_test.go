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
}

func TestWriteLogs(t *testing.T) {
	config.Setup()

	log.Setup()

	log.GetLogger(log.InstanceApp).Debug(time.Now().String())

	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go WriteLog(wg, i)
	}

	wg.Wait()
	fmt.Println("end")
}

func WriteLog(wg *sync.WaitGroup, i int) {
	for j := 0; j < 500; j++ {
		log.GetLogger(log.InstanceApp).Debugf("%d-%d-%s", i, j, time.Now().Local().String())
		time.Sleep(50 * time.Millisecond)
	}
	wg.Done()
}
