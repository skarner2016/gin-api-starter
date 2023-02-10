package log

import (
	"fmt"
	"skarner2016/gin-api-starter/packages/config"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var InstanceMap map[Instance]*zap.SugaredLogger

func Setup() {
	appPath := config.APPConfig.GetString("appPath")

	confMap := make(map[Instance]*LogConf, 0)
	if err := config.APPConfig.UnmarshalKey("log", &confMap); err != nil {
		panic(fmt.Sprintf("parse log config error:%v", err))
	}

	InstanceMap = make(map[Instance]*zap.SugaredLogger, 0)
	for i, conf := range confMap {
		InstanceMap[i] = logInit(i, conf, appPath)
	}
}

func logInit(i Instance, conf *LogConf, appPath string) *zap.SugaredLogger {
	logFile := fmt.Sprintf("%s/%s/%s", appPath, conf.Path, conf.File)
	writeSyncer := getWriter(conf, logFile)
	encoder := getEncoder()
	level := getLogLevel(conf.Level)

	core := zapcore.NewCore(encoder, writeSyncer, level)
	logger := zap.New(core, zap.AddCaller())

	return logger.Sugar()
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
