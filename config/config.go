package config

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed version
var version string

//go:embed name
var name string

type LogLevel string

const (
	Debug LogLevel = "debug"
	Info  LogLevel = "info"
	Warn  LogLevel = "warn"
	Error LogLevel = "error"
)

func GetVersion() string {
	return strings.TrimSpace(version)
}

func GetName() string {
	return strings.TrimSpace(name)
}

func GetLogLevel() LogLevel {
	if IsDebug() {
		return Debug
	}
	logLevel := os.Getenv("XUI_LOG_LEVEL")
	if logLevel == "" {
		return Info
	}
	return LogLevel(logLevel)
}

func IsDebug() bool {
	return os.Getenv("XUI_DEBUG") == "true"
}


func GetDBAddressPath() string {
	return fmt.Sprintf("/etc/%s/%s.txt", GetName(), GetName())
}

func GetDBPath() string {
	var dat string
	dat, err := os.ReadFile(GetDBAddressPath())
	if err != nil {
		return err
    	}
    	return fmt.Print(dat)
}
