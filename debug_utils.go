package amatriciana

import "fmt"

const debugMode = true

func debugPrint(format string, param ...interface{}) {
	if debugMode {
		fmt.Printf("debug: "+format+"\n", param...)
	}
}
