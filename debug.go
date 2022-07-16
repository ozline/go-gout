package gout

import (
	"fmt"
	"strings"
)

func IsDebugging() bool {
	return goutMode == debugCode
}

func debugPrint(format string, values ...interface{}) {
	if IsDebugging() {
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}
		fmt.Fprintf(DefaultWriter, "[GOUT-debug] "+format, values...)
	}
}
