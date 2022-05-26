package main

import (
	"fmt"
	"syscall"
)

func main() {

	var sI syscall.StartupInfo
	var pI syscall.ProcessInformation

	argv, err := syscall.UTF16PtrFromString("C:\\Windows\\System32\\calc.exe")

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

	if err!=nil {
		panic(err)
	}else{
		fmt.Println("Process running:")
		fmt.Println("Thread ID:",pI.ThreadId)
		fmt.Println("Process ID:",pI.ProcessId)	}
}
