package main

import (
	"flag"
	"fmt"
	"log"
	"syscall"
)

var path = flag.String("path", "C:\\Windows\\System32\\calc.exe", "the process path you want to run.")

func main() {
	flag.Parse()
	log.SetFlags(0)
	var sI syscall.StartupInfo
	var pI syscall.ProcessInformation

	argv, err := syscall.UTF16PtrFromString(*path)

	err = syscall.CreateProcess(
		nil,
		argv,
		nil,
		nil,
		true,
		0,
		nil,
		nil,
		&sI,
		&pI)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Process running:")
		fmt.Println("Thread ID:", pI.ThreadId)
		fmt.Println("Process ID:", pI.ProcessId)
	}
}
