package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

var dest = os.Stdout
var Log = log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
var LogError = log.New(os.Stdout, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)

//TODO: Support setting an output device  for all  and for individual mode

func SetVsLogDevice(newDest *os.File) {
	dest = newDest
}
func VsLog(mode, msg string) {
	_, path, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Printf("Error getting Caller")
	}
	file := ExtractFileFromPath(path)
	t := time.Now()
	tstr := fmt.Sprintf(t.Format("2006-01-02 15:04:05"))
	fmt.Printf("%s %s %s:%d - %s\n", tstr, mode, file, line, msg)
	//fmt.Printf(VsLogMsg(mode, msg))
	mode = mode + " "
	switch mode {
	case "INFO":
		log.New(dest, mode, log.Ldate|log.Ltime|log.Lshortfile)
	case "ERROR":
		log.New(dest, mode, log.Ldate|log.Ltime|log.Lshortfile)
	default:
		log.New(dest, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
	}
}

func VsLogErr(msg string) string {
	mode := "ERROR"
	_, path, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Printf("Error getting Caller")
	}
	file := ExtractFileFromPath(path)
	t := time.Now()
	tstr := fmt.Sprintf(t.Format("2006-01-02 15:04:05"))
	return fmt.Sprintf("%s %s %s:%d - %s\n", tstr, mode, file, line, msg)
}

func VsLogMsg(mode, msg string) string {
	_, path, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Printf("Error getting Caller")
	}
	file := ExtractFileFromPath(path)
	t := time.Now()
	tstr := fmt.Sprintf(t.Format("2006-01-02 15:04:05"))
	return fmt.Sprintf("%s %s %s:%d - %s\n", tstr, mode, file, line, msg)
	// mode = mode + " "
	// switch mode {
	// case "INFO":
	// 	log.New(dest, mode, log.Ldate|log.Ltime|log.Lshortfile)
	// case "ERROR":
	// 	log.New(dest, mode, log.Ldate|log.Ltime|log.Lshortfile)
	// default:
	// 	log.New(dest, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
	// }
}

func VLog(mode, msg string) {
	_, path, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Printf("Error getting Caller")
	}
	file := ExtractFileFromPath(path)
	t := time.Now()
	tstr := fmt.Sprintf(t.Format("2006-01-02 15:04:05"))
	fmt.Printf("%s %s %s:%d - %s\n", tstr, mode, file, line, msg)
	//fmt.Printf(VsLogMsg(mode, msg))
	mode = mode + " "
	switch mode {
	case "INFO":
		log.New(dest, mode, log.Ldate|log.Ltime|log.Lshortfile)
	case "ERROR":
		log.New(dest, mode, log.Ldate|log.Ltime|log.Lshortfile)
	default:
		log.New(dest, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
	}
}

func VLogMsg(mode, msg string) string {
	_, path, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Printf("Error getting Caller")
	}
	file := ExtractFileFromPath(path)
	t := time.Now()
	tstr := fmt.Sprintf(t.Format("2006-01-02 15:04:05"))
	return fmt.Sprintf("%s %s %s:%d - %s\n", tstr, mode, file, line, msg)
	// mode = mode + " "
	// switch mode {
	// case "INFO":
	// 	log.New(dest, mode, log.Ldate|log.Ltime|log.Lshortfile)
	// case "ERROR":
	// 	log.New(dest, mode, log.Ldate|log.Ltime|log.Lshortfile)
	// default:
	// 	log.New(dest, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
	// }
}

func VLogErr(msg string) string {
	mode := "ERROR"
	_, path, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Printf("Error getting Caller")
	}
	file := ExtractFileFromPath(path)
	t := time.Now()
	tstr := fmt.Sprintf(t.Format("2006-01-02 15:04:05"))
	return fmt.Sprintf("%s %s %s:%d - %s\n", tstr, mode, file, line, msg)
}

func ExtractFileFromPath(path string) string {
	_, file := path, path
	for i := len(path) - 1; i > 0; i-- {
		if path[i] == '/' {
			file = path[i+1:]
			break
		}
	}
	return file
}
