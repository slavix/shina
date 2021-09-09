package logger

import (
	"context"
	"fmt"
	formatter "github.com/fabienm/go-logrus-formatters"
	graylog "github.com/gemnasium/logrus-graylog-hook/v3"
	"github.com/sirupsen/logrus"
	"github.com/ztrue/tracerr"
	"runtime"
	"shina/pkg/config"
)

var logger *logrus.Logger

func Init(serviceName string, level logrus.Level) *logrus.Logger {
	logrusLogger := logrus.New()
	logrusLogger.SetLevel(level)

	gelfFmt := formatter.NewGelf(serviceName)
	logrusLogger.SetFormatter(gelfFmt)

	hook := graylog.NewGraylogHook(":"+config.GetString("graylog.port"), map[string]interface{}{})
	logrusLogger.AddHook(hook)

	logger = logrusLogger

	return logrusLogger
}

func Panic(err error, msg string) {
	logger.WithFields(logrus.Fields{
		"function": funcName(),
		"error":    err.Error(),
		"trace":    tracerr.Sprint(err),
	}).Panic(msg)
}

func Error(ctx context.Context, err error, msg string) {
	logger.WithFields(logrus.Fields{
		"requestID": GetRequestID(ctx),
		"function":  funcName(),
		"error":     err.Error(),
		"trace":     tracerr.Sprint(err),
	}).Error(msg)
}

func Warn(ctx context.Context, err error, msg string) {
	logger.WithFields(logrus.Fields{
		"requestID": GetRequestID(ctx),
		"function":  funcName(),
		"error":     err.Error(),
		"trace":     tracerr.Sprint(err),
	}).Warn(msg)
}

func Info(ctx context.Context, msg string) {
	logger.WithFields(logrus.Fields{
		"requestID": GetRequestID(ctx),
		"function":  funcName(),
		"message":   msg,
	}).Info(msg)
}

func Debug(ctx context.Context, msg string) {
	logger.WithFields(logrus.Fields{
		"requestID": GetRequestID(ctx),
		"function":  funcName(),
		"message":   msg,
	}).Debug(msg)
}

func Trace(ctx context.Context, msg string) {
	logger.WithFields(logrus.Fields{
		"requestID": GetRequestID(ctx),
		"function":  funcName(),
		"message":   msg,
	}).Trace(msg)
}

func LogHandler(message string, a ...interface{}) {
	logger.WithFields(logrus.Fields{
		"function": funcName(),
		"info":     fmt.Sprintf(message, a...),
	}).Info("server tcp log")
}

func GetRequestID(ctx context.Context) string {
	requestId, ok := ctx.Value("requestID").(string)

	if ok && requestId != "" {
		return requestId
	}

	logger.WithFields(logrus.Fields{
		"function": funcName(),
	}).Error("request doesn't have requestID")

	return ""
}

func funcName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}
