package ui

import "fmt"

var (
	DebugOn   bool
	VerboseOn bool
)

func Error(err error, args ...interface{}) {
	fmt.Printf("Error: %s: %s\n", fmt.Sprint(args...), err.Error())
}

func Debug(args ...interface{}) {
	if DebugOn {
		fmt.Println(args...)
	}
}

func Verbose(args ...interface{}) {
	if VerboseOn {
		fmt.Println(args...)
	}
}
