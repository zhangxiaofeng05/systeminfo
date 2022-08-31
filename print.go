package systeminfo

import (
	"fmt"
)

func Print(name string, value interface{}) {
	fmt.Printf("%s: %v\n", name, value)
}

func Debugln(name string, value interface{}) {
	if IsDebugging() {
		fmt.Printf("%s: %v\n", name, value)
	}
}
