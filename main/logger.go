package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// error logger
var errorLogger *zap.SugaredLogger

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

func initLogger(logpath string, loglevel string) *zap.SugaredLogger {
	level := getLoggerLevel(loglevel)
	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logpath,
		MaxSize:    128,
		MaxBackups: 100,
		MaxAge:     30,
		LocalTime:  true,
		Compress:   true,
	})
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(level)),
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), zapcore.AddSync(os.Stdout), zap.NewAtomicLevelAt(level)),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	res := logger.Sugar()
	return res
}

//Debug ...
func Debug(args ...interface{}) {
	errorLogger.Debug(args...)
}

//Debugf ...
func Debugf(template string, args ...interface{}) {
	errorLogger.Debugf(template, args...)
}

//Info ...
func Info(args ...interface{}) {
	errorLogger.Info(args...)
}

//Infof ...
func Infof(template string, args ...interface{}) {
	errorLogger.Infof(template, args...)
}

//Warn ...
func Warn(args ...interface{}) {
	errorLogger.Warn(args...)
}

//Warnf ...
func Warnf(template string, args ...interface{}) {
	errorLogger.Warnf(template, args...)
}

//Error ...
func Error(args ...interface{}) {
	errorLogger.Error(args...)
}

//Errorf ...
func Errorf(template string, args ...interface{}) {
	errorLogger.Errorf(template, args...)
}

//DPanic ...
func DPanic(args ...interface{}) {
	errorLogger.DPanic(args...)
}

//DPanicf ...
func DPanicf(template string, args ...interface{}) {
	errorLogger.DPanicf(template, args...)
}

//Panic ..
func Panic(args ...interface{}) {
	errorLogger.Panic(args...)
}

//Panicf ...
func Panicf(template string, args ...interface{}) {
	errorLogger.Panicf(template, args...)
}

//Fatal ...
func Fatal(args ...interface{}) {
	errorLogger.Fatal(args...)
}

//Fatalf ...
func Fatalf(template string, args ...interface{}) {
	errorLogger.Fatalf(template, args...)
}
