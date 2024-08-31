package config

import (
	"path/filepath"
	"runtime"
)

var (
	// Get current vfile full path from runtime
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	ProjectRootPath = filepath.Join(filepath.Dir(b), "../")
)