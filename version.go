package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

var (
	// reference:
	// 1. https://www.forkingbytes.com/blog/dynamic-versioning-your-go-application/
	// 2. https://goreleaser.com/customization/builds/
	// 3. https://github.com/goreleaser/goreleaser/blob/main/main.go
	version = "dev"
	commit  = ""
	date    = ""
	builtBy = ""
)

func getVersion() string {
	return buildVersion(version, commit, date, builtBy)
}

func buildVersion(version, commit, date, builtBy string) string {
	result := version
	if commit != "" {
		result = fmt.Sprintf("%s\ncommit: %s", result, commit)
	}
	if date != "" {
		result = fmt.Sprintf("%s\nbuilt at: %s", result, date)
	}
	if builtBy != "" {
		result = fmt.Sprintf("%s\nbuilt by: %s", result, builtBy)
	}
	result = fmt.Sprintf("%s\ngoos: %s\ngoarch: %s", result, runtime.GOOS, runtime.GOARCH)
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Sum != "" {
		result = fmt.Sprintf("%s\nmodule version: %s, checksum: %s", result, info.Main.Version, info.Main.Sum)
	}
	return result
}
