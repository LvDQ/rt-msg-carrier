package log

import (
	"os"
	"rt-msg-carrier/configs"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger
var AccessLogger *logrus.Logger

func init() {
	//日志实例化
	Logger = logrus.New()
	AccessLogger = logrus.New()
	//读取配置
	logcfg := configs.Get().LogConfig

	file := configs.ProjectBizLogFile

	// src, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	// if err != nil {
	// 	fmt.Println("err", err)
	// }
	// fmt.Println(src)
	//设置输出到文件
	// Logger.Out = src
	// 输出至标准输出
	Logger.Out = os.Stdout
	//设置日志级别
	// Logger.SetLevel(logrus.DebugLevel)
	//设置日志格式
	// Logger.SetFormatter(&logrus.JSONFormatter{})

	switch strings.ToLower(logcfg.Level) {
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "warn", "warning":
		logrus.SetLevel(logrus.WarnLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
	switch strings.ToLower(logcfg.Format) {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{})
	default:
		logrus.SetFormatter(&logrus.TextFormatter{})
	}

	//设置日志切割 rotatelogs
	logWriter1, _ := rotatelogs.New(
		file+".%Y%m%d_writer1.log",
		//生成软链 指向最新的日志文件
		rotatelogs.WithLinkName(file+".debug_log.latest"),
		//文件最大保存时间
		rotatelogs.WithMaxAge(7*24*time.Hour),
		//设置日志切割时间间隔(1天)(隔多久分割一次)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	logWriter2, _ := rotatelogs.New(
		file+".%Y%m%d_writer2.log",
		//生成软链 指向最新的日志文件
		rotatelogs.WithLinkName(file+".error_log.latest"),
		//文件最大保存时间
		rotatelogs.WithMaxAge(7*24*time.Hour),
		//设置日志切割时间间隔(1天)(隔多久分割一次)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	debugWriteMap := lfshook.WriterMap{
		logrus.DebugLevel: logWriter1,
		logrus.InfoLevel:  logWriter1,
		logrus.WarnLevel:  logWriter1,
		logrus.ErrorLevel: logWriter1,
		logrus.PanicLevel: logWriter1,
		logrus.FatalLevel: logWriter1,
	}
	lfDebugHook := lfshook.NewHook(debugWriteMap, &logrus.JSONFormatter{})
	Logger.AddHook(lfDebugHook)

	errorWriteMap := lfshook.WriterMap{
		logrus.ErrorLevel: logWriter2,
		logrus.PanicLevel: logWriter2,
		logrus.FatalLevel: logWriter2,
	}
	lfErrorHook := lfshook.NewHook(errorWriteMap, &logrus.JSONFormatter{})
	Logger.AddHook(lfErrorHook)

}

func NewLogger() *logrus.Logger {
	if Logger != nil {
		return Logger
	}
	return Logger
}

func NewAccessLogger() *logrus.Logger {
	if AccessLogger != nil {
		return AccessLogger
	}
	file := configs.ProjectAccessLogFile
	//设置日志切割 rotatelogs
	logWriter1, _ := rotatelogs.New(
		file+".%Y%m%d_access_log.log",
		//生成软链 指向最新的日志文件
		rotatelogs.WithLinkName(file+".access_log.latest"),
		//文件最大保存时间
		rotatelogs.WithMaxAge(7*24*time.Hour),
		//设置日志切割时间间隔(1天)(隔多久分割一次)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	WriteMap := lfshook.WriterMap{
		logrus.DebugLevel: logWriter1,
		logrus.InfoLevel:  logWriter1,
		logrus.WarnLevel:  logWriter1,
		logrus.ErrorLevel: logWriter1,
		logrus.PanicLevel: logWriter1,
		logrus.FatalLevel: logWriter1,
	}
	lfHook := lfshook.NewHook(WriteMap, &logrus.TextFormatter{})
	AccessLogger.AddHook(lfHook)

	return AccessLogger
}
