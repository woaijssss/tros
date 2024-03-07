package trlogger

import (
	"context"
	"fmt"
	"gitee.com/idigpower/tros/conf"
	"gitee.com/idigpower/tros/constants"
	context3 "gitee.com/idigpower/tros/context"
	"gitee.com/idigpower/tros/sys"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"sort"
	"time"
)

const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel = "panic"
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel = "fatal"
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel = "error"
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel = "warn"
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel = "info"
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel = "debug"
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel = "trace"
)

const (
	DefaultTimestampFormat = "2006-01-02 15:04:05.999999"
	logEntryKey            = "logEntry"
	logFormat              = "%s.log.%s"
	//logExtraFormat         = "%Y%m%d%H%M%S"
	logExtraFormat = "%Y%m%d%H"
)

type contextKey int8

var (
	key contextKey
)

type TrLogger struct {
	log *logrus.Logger
}

func DefaultTrLogger() *TrLogger {
	level := DebugLevel
	if conf.IsProd() {
		level = InfoLevel
	}
	return NewTrLogger(level, DefaultTimestampFormat, os.Stdout)
}

func NewTrLogger(level string, timestampFormat string, out io.Writer) *TrLogger {
	log := &logrus.Logger{
		Out: out,
		Formatter: &logrus.TextFormatter{
			DisableQuote:           true,
			FullTimestamp:          true,
			TimestampFormat:        timestampFormat,
			DisableSorting:         true,
			DisableLevelTruncation: true,
			PadLevelText:           false,
			SortingFunc: func(strings []string) {
				sort.Slice(strings, func(i, j int) bool {
					if strings[i] == "level" {
						return true
					}
					return false
				})
			},
		},
		Level: parseLevel(level),
	}

	xl := &TrLogger{log: log}
	setupLogger(xl)

	return xl
}

func (dl *TrLogger) Debugf(ctx context.Context, format string, args ...any) {
	dl.withFields(ctx).Debugf(format, args...)
}

func (dl *TrLogger) Infof(ctx context.Context, format string, args ...any) {
	dl.withFields(ctx).Infof(format, args...)
}

func (dl *TrLogger) Warnf(ctx context.Context, format string, args ...any) {
	dl.withFields(ctx).Warnf(format, args...)
}

func (dl *TrLogger) Errorf(ctx context.Context, format string, args ...any) {
	dl.withFields(ctx).Errorf(format, args...)
}

func (dl *TrLogger) Fatalf(ctx context.Context, format string, args ...any) {
	dl.withFields(ctx).Fatalf(format, args...)
}

func (dl *TrLogger) Panicf(ctx context.Context, format string, args ...any) {
	dl.withFields(ctx).Panicf(format, args...)
}

func (dl *TrLogger) Debug(ctx context.Context, args ...any) {
	dl.withFields(ctx).Debug(args...)
}

func (dl *TrLogger) Info(ctx context.Context, args ...any) {
	dl.withFields(ctx).Info(args...)
}

func (dl *TrLogger) Warn(ctx context.Context, args ...any) {
	dl.withFields(ctx).Warn(args...)
}

func (dl *TrLogger) Error(ctx context.Context, args ...any) {
	dl.withFields(ctx).Error(args...)
}

func (dl *TrLogger) Fatal(ctx context.Context, args ...any) {
	dl.withFields(ctx).Fatal(args...)
}

func (dl *TrLogger) Panic(ctx context.Context, args ...any) {
	dl.withFields(ctx).Panic(args...)
}

func (dl *TrLogger) Debugln(ctx context.Context, args ...any) {
	dl.withFields(ctx).Debugln(args...)
}

func (dl *TrLogger) Infoln(ctx context.Context, args ...any) {
	dl.withFields(ctx).Infoln(args...)
}

func (dl *TrLogger) Warnln(ctx context.Context, args ...any) {
	dl.withFields(ctx).Warnln(args...)
}

func (dl *TrLogger) Errorln(ctx context.Context, args ...any) {
	dl.withFields(ctx).Errorln(args...)
}

func (dl *TrLogger) Fatalln(ctx context.Context, args ...any) {
	dl.withFields(ctx).Fatalln(args...)
}

func (dl *TrLogger) Panicln(ctx context.Context, args ...any) {
	dl.withFields(ctx).Panicln(args...)
}

func (dl *TrLogger) SetLevel(level string) {
	dl.log.SetLevel(parseLevel(level))
}

func (dl *TrLogger) withFields(ctx context.Context) *logrus.Entry {
	//fmt.Println("--->: ", reflect.ValueOf(ctx.Value(interceptor.TraceKey)))
	traceId, ok := ctx.Value(constants.TraceId).(string)
	if !ok {
		traceId = context3.GenTraceID(ctx)
	}
	//if ctx.GetExtraValue(logEntryKey) != nil {
	//	return ctx.GetExtraValue(logEntryKey).(*logrus.Entry)
	//}
	//
	//if ctx.GoId == 0 {
	//	ctx.GoId = sys.QuickGetGoRoutineId()
	//}

	entry := dl.log.WithFields(logrus.Fields{
		constants.GoId:    sys.QuickGetGoRoutineId(),
		constants.TraceId: traceId,
	})

	//ctx.SetExtraKeyValue(logEntryKey, entry)

	return entry
}

func parseLevel(level string) logrus.Level {
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logLevel = logrus.DebugLevel
	}

	return logLevel
}

func setupLogger(xl *TrLogger) {
	logFile := conf.GetLogPath()
	appName := conf.GetAppName()
	if logFile != "" {
		// 配置文件切割
		writer, _ := rotatelogs.New(getLogFormat(logFile+appName),
			//rotatelogs.WithLinkName("./log/logfile.log"),
			rotatelogs.WithMaxAge(7*24*time.Hour),    // 文件最大保存时间
			rotatelogs.WithRotationTime(time.Hour/2), // 日志切割时间间隔
			//rotatelogs.WithRotationTime(time.Second), // 日志切割时间间隔
		)
		xl.log.SetOutput(writer)
	}
}

func getLogFormat(file string) string {
	return fmt.Sprintf(logFormat, file, logExtraFormat)
}
