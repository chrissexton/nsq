package util

import (
	"fmt"
	"runtime"
)

const BINARY_VERSION = "0.3.0-alpha"

func Version(app string) string {
	return fmt.Sprintf("%s v%s (built w/%s)", app, BINARY_VERSION, runtime.Version())
}
