package systeminfo

import "os"

const EnvSystemInfoMode = "SYSTEM_INFO_MODE"

const (
	DebugMode = "debug"
	InfoMode  = "info"
)

const (
	debugCode = iota
	InfoCode
)

var (
	systemInfoMode = debugCode
	modeName       = DebugMode
)

func init() {
	mode := os.Getenv(EnvSystemInfoMode)
	SetMode(mode)
}

func SetMode(value string) {
	if value == "" {
		value = DebugMode
	}

	switch value {
	case DebugMode:
		systemInfoMode = debugCode
	case InfoMode:
		systemInfoMode = InfoCode
	default:
		panic("system info mode unknown: " + value + " (available mode: debug info)")
	}

	modeName = value
}

func IsDebugging() bool {
	return systemInfoMode == debugCode
}
