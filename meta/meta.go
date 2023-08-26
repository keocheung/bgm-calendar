// Package meta contains meta data for bgm-calendar
package meta

var (
	// Version is the compile-time set version of bgm-calendar
	Version = "v0.0.0"
	// UserAgent is the User-Agent generated from Version
	UserAgent string
)

func init() {
	UserAgent = "keo/bgm-calendar/" + Version
}
