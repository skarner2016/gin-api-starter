package log

import (
	"fmt"
	"path"
	"runtime"
	"skarner2016/gin-api-starter/packages/config"
	"strings"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var InstanceMap map[Instance]*zap.SugaredLogger

func Setup() {
	confMap := make(map[Instance]*LogConf, 0)
	if err := config.APPConfig.UnmarshalKey("log", &confMap); err != nil {
		panic(fmt.Sprintf("parse log config error:%v", err))
	}

	InstanceMap = make(map[Instance]*zap.SugaredLogger, 0)
	for i, conf := range confMap {
		InstanceMap[i] = logInit(i, conf)
	}
}

func logInit(i Instance, conf *LogConf) *zap.SugaredLogger {
	logFile := getLogFile(conf)
	writeSyncer := getWriter(conf, logFile)
	encoder := getEncoder()
	level := getLogLevel(conf.Level)

	core := zapcore.NewCore(encoder, writeSyncer, level)
	logger := zap.New(core, zap.AddCaller())

	return logger.Sugar()
}

func getLogFile(conf *LogConf) string {
	logFile := ""
	if strings.HasPrefix(conf.Path, "/") {
		// 绝对路径
		logFile = fmt.Sprintf("%s/%s", conf.Path, conf.File)
	} else {
		// 相对路径 /../../storage/
		currentDir := getCurrentDir()
		logFile = fmt.Sprintf("%s/../../%s/%s", currentDir, conf.Path, conf.File)
	}

	fmt.Println("logfile: " + logFile)

	return logFile
}

func getCurrentDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	return path.Dir(filename)
}

func getWriter(conf *LogConf, logFile string) zapcore.WriteSyncer {

	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    conf.MaxSize,
		MaxAge:     conf.MaxAge,
		MaxBackups: conf.MaxBackup,
		LocalTime:  conf.LocalTime,
		Compress:   conf.Compress,
	}

	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogLevel(level string) zapcore.LevelEnabler {

	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	case "panic":
		return zapcore.PanicLevel
	case "info":
	default:
		return zapcore.InfoLevel
	}

	return zapcore.InfoLevel
}

func GetLogger(instance Instance) *zap.SugaredLogger {
	if _, ok := InstanceMap[instance]; !ok {
		Setup()
	}

	sugaredLogger, ok := InstanceMap[instance]
	if !ok {
		panic(fmt.Sprintf("get logger error:%s", instance))
	}

	return sugaredLogger
}
