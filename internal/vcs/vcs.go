// Package vcs manages the build version
package vcs

import (
	"runtime/debug"
)

func Version() string {
	bi, ok := debug.ReadBuildInfo()
	if ok {
		return bi.Main.Version
	}
	return ""
}
