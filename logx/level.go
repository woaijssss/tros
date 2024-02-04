package trlogger

import (
	"gitee.com/idigpower/tros/constants"
	"github.com/spf13/viper"
)

// Level type of log level
type Level int8

const (
	// LevelDebug debug level
	LevelDebug = iota
	// LevelInfo info level
	LevelInfo
	// LevelWarn warn level
	LevelWarn
	// LevelError error level
	LevelError
	// LevelFatal fatal level
	LevelFatal
)

// LevelFromString convert string to Level
// if convert failed then return LevelInfo
func LevelFromString(str string) Level {
	switch str {
	case "debug", "DEBUG":
		return LevelDebug
	case "info", "INFO":
		return LevelInfo
	case "warn", "WARN":
		return LevelWarn
	case "error", "ERROR":
		return LevelError
	case "fatal", "FATAL":
		return LevelFatal
	}
	return LevelInfo
}

// GetConfiguredLevel get configured Level
func GetConfiguredLevel() Level {
	str := viper.GetString(constants.LogLevel)
	return LevelFromString(str)
}

// IsDebugLevel check if Level is LevelDebug
func IsDebugLevel() bool {
	return GetConfiguredLevel() == LevelDebug
}
