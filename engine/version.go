//
// Warning:
// This file is generated by go compiler, don't change it!!!
// Generated Time: "2021-09-29 14:47:48"
//
package engine

type ver struct {
	Version   string
	ReleaseTime string
}

var defaultVer = ver{
	Version:   `V0.0.2-9ac4c01bc579ffaf2102400490af111b240fbbab`,
	ReleaseTime: "2021-09-29 14:47:48",
}
func Version() ver{
	return defaultVer
}