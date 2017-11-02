package main

import (
	"syscall"
	"unsafe"
)

func main() {
	k32, _ := syscall.LoadDLL("kernel32.dll")
	defer k32.Release()
	kproc, _ := k32.FindProc("SetThreadExecutionState")
	kproc.Call(0x80000003)

	u32, _ := syscall.LoadDLL("user32.dll")
	defer u32.Release()
	proc, _ := u32.FindProc("MessageBoxW")
	proc.Call(
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("exit ?"))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("AlwaysOn"))),
		0)
}
