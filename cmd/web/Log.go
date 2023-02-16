// package main

// import (
// 	"log"
// 	"runtime"
// )

// func CreateLogger() {
// 	// create file
// 	log.SetOutput()
// }

// func LogError(msg string) {
// 	_, file, line, ok := runtime.Caller(1)
// 	if !ok{
// 		//handle
// 	}
// 	log.Printf("[ERROR]: File: %s Line: %d: %s", file, line, msg)
// }

// func LogInfo(msg string) {
// 	_, file, line, ok := runtime.Caller(1)
// 	if !ok{
// 		//handle
// 	}
// 	log.Printf("[INFO]: File: %s Line: %d: %s", file, line, msg)
// }