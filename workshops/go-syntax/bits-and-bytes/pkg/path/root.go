package path

import (
	"path/filepath"
	"runtime"
)

var (
	_, base, _, _ = runtime.Caller(0)

	// Root returns the relative root dir of this project during runtime
	Root = filepath.Join(filepath.Dir(base), "../..")
)
