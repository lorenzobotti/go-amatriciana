package amatriciana

import "fmt"

const debugMode = true

func debugPrintf(format string, param ...interface{}) {
	if debugMode {
		fmt.Printf("debug: "+format+"\n", param...)
	}
}
