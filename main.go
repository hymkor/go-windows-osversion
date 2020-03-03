package osversion

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

// http://yamatyuu.net/computer/program/sdk/base/RtlGetVersion/index.html

var ntdll = windows.NewLazySystemDLL("ntdll.dll")
var procRtlGetVersion = ntdll.NewProc("RtlGetVersion")

type _OSVERSIONINFOEXW struct {
	dwOSVersionInfoSize uint32
	dwMajorVersion      uint32
	dwMinorVersion      uint32
	dwBuildNumber       uint32
	dwPlatformId        uint32
	szCSDVersion        [128]uint16
	wServicePackMajor   uint16
	wServicePackMinor   uint16
	wSuiteMask          uint16
	wProductType        uint8
	wReserved           uint8
}

type Version struct {
	Major            uint32
	Minor            uint32
	Build            uint32
	PlatformId       uint32
	ServicePackMajor uint16
	ServicePackMinor uint16
}

func Query() *Version {
	var buffer _OSVERSIONINFOEXW

	buffer.dwOSVersionInfoSize = uint32(unsafe.Sizeof(buffer))
	procRtlGetVersion.Call(uintptr(unsafe.Pointer(&buffer)))

	// _OSVERSIONINFOEXW is too large to return.
	// So, copy some member to smaller struct.
	return &Version{
		Major:            buffer.dwMajorVersion,
		Minor:            buffer.dwMinorVersion,
		Build:            buffer.dwBuildNumber,
		PlatformId:       buffer.dwPlatformId,
		ServicePackMajor: buffer.wServicePackMajor,
		ServicePackMinor: buffer.wServicePackMinor,
	}
}
