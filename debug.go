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

// func debugPrintRoute(httpMethod, absolutePath string, handlers HandlersChain) {
// 	if IsDebugging() {
// 		nuHandlers := len(handlers)
// 		handlerName := nameOfFunction(handlers.Last())
// 		if DebugPrintRouteFunc == nil {
// 			debugPrint("%-6s %-25s --> %s (%d handlers)\n", httpMethod, absolutePath, handlerName, nuHandlers)
// 		} else {
// 			DebugPrintRouteFunc(httpMethod, absolutePath, handlerName, nuHandlers)
// 		}
// 	}
// }
