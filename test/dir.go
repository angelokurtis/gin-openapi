package test

import (
	"path"
	"runtime"
)

func Dir() string {
	_, b, _, _ := runtime.Caller(0)
	return path.Dir(b)
}
