package golang

import (
	"runtime"
)

// Version returns the golang runtime version
func Version() string {
	return runtime.Version()
}
