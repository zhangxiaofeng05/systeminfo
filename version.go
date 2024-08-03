package main

import (
	"encoding/json"
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
	branch  = ""
	tagDate = ""
	builtBy = ""
)

type versionInfo struct {
	Version     string `json:"version"`
	Commit      string `json:"commit"`
	Branch      string `json:"branch"`
	TagDate     string `json:"tagDate"`
	BuiltBy     string `json:"builtBy"`
	Goos        string `json:"goos"`
	Goarch      string `json:"goarch"`
	MainVersion string `json:"mainVersion"`
	MainSum     string `json:"mainSum"`
}

func getVersion() *versionInfo {
	v := &versionInfo{
		Version: version,
		Commit:  commit,
		Branch:  branch,
		TagDate: tagDate,
		BuiltBy: builtBy,
		Goos:    runtime.GOOS,
		Goarch:  runtime.GOARCH,
	}
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Sum != "" {
		v.MainVersion = info.Main.Version
		v.MainSum = info.Main.Sum
	}
	return v
}

func (v *versionInfo) String() string {
	res, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return ""
	}
	return string(res)
}
