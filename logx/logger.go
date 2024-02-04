package trlogger

import (
	"context"
)

var GlobalTrLogger = DefaultTrLogger()

func Debugf(ctx context.Context, format string, args ...any) {
	GlobalTrLogger.Debugf(ctx, format, args...)
}

func Infof(ctx context.Context, format string, args ...any) {
	GlobalTrLogger.Infof(ctx, format, args...)
}

func Warnf(ctx context.Context, format string, args ...any) {
	GlobalTrLogger.Warnf(ctx, format, args...)
}

func Errorf(ctx context.Context, format string, args ...any) {
	GlobalTrLogger.Errorf(ctx, format, args...)
}

func Fatalf(ctx context.Context, format string, args ...any) {
	GlobalTrLogger.Fatalf(ctx, format, args...)
}

func Panicf(ctx context.Context, format string, args ...any) {
	GlobalTrLogger.Panicf(ctx, format, args...)
}

func Debug(ctx context.Context, args ...any) {
	GlobalTrLogger.Debug(ctx, args...)
}

func Info(ctx context.Context, args ...any) {
	GlobalTrLogger.Info(ctx, args...)
}

func Warn(ctx context.Context, args ...any) {
	GlobalTrLogger.Warn(ctx, args...)
}

func Error(ctx context.Context, args ...any) {
	GlobalTrLogger.Error(ctx, args...)
}

func Fatal(ctx context.Context, args ...any) {
	GlobalTrLogger.Fatal(ctx, args...)
}

func Panic(ctx context.Context, args ...any) {
	GlobalTrLogger.Panic(ctx, args...)
}

func Debugln(ctx context.Context, args ...any) {
	GlobalTrLogger.Debugln(ctx, args...)
}

func Infoln(ctx context.Context, args ...any) {
	GlobalTrLogger.Infoln(ctx, args...)
}

func Warnln(ctx context.Context, args ...any) {
	GlobalTrLogger.Warnln(ctx, args...)
}

func Errorln(ctx context.Context, args ...any) {
	GlobalTrLogger.Errorln(ctx, args...)
}

func Fatalln(ctx context.Context, args ...any) {
	GlobalTrLogger.Fatalln(ctx, args...)
}

func Panicln(ctx context.Context, args ...any) {
	GlobalTrLogger.Panicln(ctx, args...)
}
