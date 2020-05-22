package main

import (
	"fmt"
	"log"

	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	fmt.Println("hello")

	log.SetOutput(&lumberjack.Logger{
		Filename:   "demo.log",
		MaxSize:    1,    // megabytes
		MaxBackups: 3,    // files
		MaxAge:     28,   // days
		Compress:   true, // zip
	})

	for i := 1; i <= 100000; i++ {
		log.Println("This is the log message will be in the log file.")
	}
}
