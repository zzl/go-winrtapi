package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"log"
	"syscall"
	"unsafe"
)

// enums

// enum
type PlatformTelemetryRegistrationStatus int32

const (
	PlatformTelemetryRegistrationStatus_Success            PlatformTelemetryRegistrationStatus = 0
	PlatformTelemetryRegistrationStatus_SettingsOutOfRange PlatformTelemetryRegistrationStatus = 1
	PlatformTelemetryRegistrationStatus_UnknownFailure     PlatformTelemetryRegistrationStatus = 2
)

// interfaces

// 9BF3F25D-D5C3-4EEA-8DBE-9C8DBB0D9D8F
var IID_IPlatformTelemetryClientStatics = syscall.GUID{0x9BF3F25D, 0xD5C3, 0x4EEA,
	[8]byte{0x8D, 0xBE, 0x9C, 0x8D, 0xBB, 0x0D, 0x9D, 0x8F}}

type IPlatformTelemetryClientStaticsInterface interface {
	win32.IInspectableInterface
	Register(id string) *IPlatformTelemetryRegistrationResult
	RegisterWithSettings(id string, settings *IPlatformTelemetryRegistrationSettings) *IPlatformTelemetryRegistrationResult
}

type IPlatformTelemetryClientStaticsVtbl struct {
	win32.IInspectableVtbl
	Register             uintptr
	RegisterWithSettings uintptr
}

type IPlatformTelemetryClientStatics struct {
	win32.IInspectable
}

func (this *IPlatformTelemetryClientStatics) Vtbl() *IPlatformTelemetryClientStaticsVtbl {
	return (*IPlatformTelemetryClientStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPlatformTelemetryClientStatics) Register(id string) *IPlatformTelemetryRegistrationResult {
	var _result *IPlatformTelemetryRegistrationResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Register, uintptr(unsafe.Pointer(this)), NewHStr(id).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPlatformTelemetryClientStatics) RegisterWithSettings(id string, settings *IPlatformTelemetryRegistrationSettings) *IPlatformTelemetryRegistrationResult {
	var _result *IPlatformTelemetryRegistrationResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterWithSettings, uintptr(unsafe.Pointer(this)), NewHStr(id).Ptr, uintptr(unsafe.Pointer(settings)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4D8518AB-2292-49BD-A15A-3D71D2145112
var IID_IPlatformTelemetryRegistrationResult = syscall.GUID{0x4D8518AB, 0x2292, 0x49BD,
	[8]byte{0xA1, 0x5A, 0x3D, 0x71, 0xD2, 0x14, 0x51, 0x12}}

type IPlatformTelemetryRegistrationResultInterface interface {
	win32.IInspectableInterface
	Get_Status() PlatformTelemetryRegistrationStatus
}

type IPlatformTelemetryRegistrationResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
}

type IPlatformTelemetryRegistrationResult struct {
	win32.IInspectable
}

func (this *IPlatformTelemetryRegistrationResult) Vtbl() *IPlatformTelemetryRegistrationResultVtbl {
	return (*IPlatformTelemetryRegistrationResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPlatformTelemetryRegistrationResult) Get_Status() PlatformTelemetryRegistrationStatus {
	var _result PlatformTelemetryRegistrationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 819A8582-CA19-415E-BB79-9C224BFA3A73
var IID_IPlatformTelemetryRegistrationSettings = syscall.GUID{0x819A8582, 0xCA19, 0x415E,
	[8]byte{0xBB, 0x79, 0x9C, 0x22, 0x4B, 0xFA, 0x3A, 0x73}}

type IPlatformTelemetryRegistrationSettingsInterface interface {
	win32.IInspectableInterface
	Get_StorageSize() uint32
	Put_StorageSize(value uint32)
	Get_UploadQuotaSize() uint32
	Put_UploadQuotaSize(value uint32)
}

type IPlatformTelemetryRegistrationSettingsVtbl struct {
	win32.IInspectableVtbl
	Get_StorageSize     uintptr
	Put_StorageSize     uintptr
	Get_UploadQuotaSize uintptr
	Put_UploadQuotaSize uintptr
}

type IPlatformTelemetryRegistrationSettings struct {
	win32.IInspectable
}

func (this *IPlatformTelemetryRegistrationSettings) Vtbl() *IPlatformTelemetryRegistrationSettingsVtbl {
	return (*IPlatformTelemetryRegistrationSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPlatformTelemetryRegistrationSettings) Get_StorageSize() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StorageSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPlatformTelemetryRegistrationSettings) Put_StorageSize(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StorageSize, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPlatformTelemetryRegistrationSettings) Get_UploadQuotaSize() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UploadQuotaSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPlatformTelemetryRegistrationSettings) Put_UploadQuotaSize(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_UploadQuotaSize, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// classes

type PlatformTelemetryClient struct {
	RtClass
}

func NewIPlatformTelemetryClientStatics() *IPlatformTelemetryClientStatics {
	var p *IPlatformTelemetryClientStatics
	hs := NewHStr("Windows.System.Diagnostics.Telemetry.PlatformTelemetryClient")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPlatformTelemetryClientStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type PlatformTelemetryRegistrationResult struct {
	RtClass
	*IPlatformTelemetryRegistrationResult
}

type PlatformTelemetryRegistrationSettings struct {
	RtClass
	*IPlatformTelemetryRegistrationSettings
}

func NewPlatformTelemetryRegistrationSettings() *PlatformTelemetryRegistrationSettings {
	hs := NewHStr("Windows.System.Diagnostics.Telemetry.PlatformTelemetryRegistrationSettings")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &PlatformTelemetryRegistrationSettings{
		RtClass:                                RtClass{PInspect: p},
		IPlatformTelemetryRegistrationSettings: (*IPlatformTelemetryRegistrationSettings)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}
