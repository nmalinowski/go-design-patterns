package main

import "fmt"

func main() {
	//log := getLoggerInstance()
	//
	//log.SetLogLevel(1)
	//log.Log("Error log level 1")
	//
	//log = getLoggerInstance()
	//log.SetLogLevel(2)
	//log.Log("Error log level 2")
	//
	//log = getLoggerInstance()
	//log.SetLogLevel(3)
	//log.Log("Error log level 3")

	// TODO: create several goroutines that try to get the
	// logger instance concurrently
	for i := 1; i < 10; i++ {
		go getLoggerInstance()
	}
	fmt.Scanln()
}
