package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// structs

type SystemManufacturersContract struct {
}

// interfaces

// 8D2EAE55-87EF-4266-86D0-C4AFBEB29BB9
var IID_IOemSupportInfo = syscall.GUID{0x8D2EAE55, 0x87EF, 0x4266,
	[8]byte{0x86, 0xD0, 0xC4, 0xAF, 0xBE, 0xB2, 0x9B, 0xB9}}

type IOemSupportInfoInterface interface {
	win32.IInspectableInterface
	Get_SupportLink() *IUriRuntimeClass
	Get_SupportAppLink() *IUriRuntimeClass
	Get_SupportProvider() string
}

type IOemSupportInfoVtbl struct {
	win32.IInspectableVtbl
	Get_SupportLink     uintptr
	Get_SupportAppLink  uintptr
	Get_SupportProvider uintptr
}

type IOemSupportInfo struct {
	win32.IInspectable
}

func (this *IOemSupportInfo) Vtbl() *IOemSupportInfoVtbl {
	return (*IOemSupportInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IOemSupportInfo) Get_SupportLink() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportLink, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IOemSupportInfo) Get_SupportAppLink() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportAppLink, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IOemSupportInfo) Get_SupportProvider() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 080CCA7C-637C-48C4-B728-F9273812DB8E
var IID_ISmbiosInformationStatics = syscall.GUID{0x080CCA7C, 0x637C, 0x48C4,
	[8]byte{0xB7, 0x28, 0xF9, 0x27, 0x38, 0x12, 0xDB, 0x8E}}

type ISmbiosInformationStaticsInterface interface {
	win32.IInspectableInterface
	Get_SerialNumber() string
}

type ISmbiosInformationStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_SerialNumber uintptr
}

type ISmbiosInformationStatics struct {
	win32.IInspectable
}

func (this *ISmbiosInformationStatics) Vtbl() *ISmbiosInformationStaticsVtbl {
	return (*ISmbiosInformationStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmbiosInformationStatics) Get_SerialNumber() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SerialNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 05880B99-8247-441B-A996-A1784BAB79A8
var IID_ISystemSupportDeviceInfo = syscall.GUID{0x05880B99, 0x8247, 0x441B,
	[8]byte{0xA9, 0x96, 0xA1, 0x78, 0x4B, 0xAB, 0x79, 0xA8}}

type ISystemSupportDeviceInfoInterface interface {
	win32.IInspectableInterface
	Get_OperatingSystem() string
	Get_FriendlyName() string
	Get_SystemManufacturer() string
	Get_SystemProductName() string
	Get_SystemSku() string
	Get_SystemHardwareVersion() string
	Get_SystemFirmwareVersion() string
}

type ISystemSupportDeviceInfoVtbl struct {
	win32.IInspectableVtbl
	Get_OperatingSystem       uintptr
	Get_FriendlyName          uintptr
	Get_SystemManufacturer    uintptr
	Get_SystemProductName     uintptr
	Get_SystemSku             uintptr
	Get_SystemHardwareVersion uintptr
	Get_SystemFirmwareVersion uintptr
}

type ISystemSupportDeviceInfo struct {
	win32.IInspectable
}

func (this *ISystemSupportDeviceInfo) Vtbl() *ISystemSupportDeviceInfoVtbl {
	return (*ISystemSupportDeviceInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemSupportDeviceInfo) Get_OperatingSystem() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OperatingSystem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemSupportDeviceInfo) Get_FriendlyName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FriendlyName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemSupportDeviceInfo) Get_SystemManufacturer() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SystemManufacturer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemSupportDeviceInfo) Get_SystemProductName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SystemProductName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemSupportDeviceInfo) Get_SystemSku() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SystemSku, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemSupportDeviceInfo) Get_SystemHardwareVersion() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SystemHardwareVersion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemSupportDeviceInfo) Get_SystemFirmwareVersion() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SystemFirmwareVersion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// EF750974-C422-45D7-A44D-5C1C0043A2B3
var IID_ISystemSupportInfoStatics = syscall.GUID{0xEF750974, 0xC422, 0x45D7,
	[8]byte{0xA4, 0x4D, 0x5C, 0x1C, 0x00, 0x43, 0xA2, 0xB3}}

type ISystemSupportInfoStaticsInterface interface {
	win32.IInspectableInterface
	Get_LocalSystemEdition() string
	Get_OemSupportInfo() *IOemSupportInfo
}

type ISystemSupportInfoStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_LocalSystemEdition uintptr
	Get_OemSupportInfo     uintptr
}

type ISystemSupportInfoStatics struct {
	win32.IInspectable
}

func (this *ISystemSupportInfoStatics) Vtbl() *ISystemSupportInfoStaticsVtbl {
	return (*ISystemSupportInfoStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemSupportInfoStatics) Get_LocalSystemEdition() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocalSystemEdition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemSupportInfoStatics) Get_OemSupportInfo() *IOemSupportInfo {
	var _result *IOemSupportInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OemSupportInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 33F349A4-3FA1-4986-AA4B-057420455E6D
var IID_ISystemSupportInfoStatics2 = syscall.GUID{0x33F349A4, 0x3FA1, 0x4986,
	[8]byte{0xAA, 0x4B, 0x05, 0x74, 0x20, 0x45, 0x5E, 0x6D}}

type ISystemSupportInfoStatics2Interface interface {
	win32.IInspectableInterface
	Get_LocalDeviceInfo() *ISystemSupportDeviceInfo
}

type ISystemSupportInfoStatics2Vtbl struct {
	win32.IInspectableVtbl
	Get_LocalDeviceInfo uintptr
}

type ISystemSupportInfoStatics2 struct {
	win32.IInspectable
}

func (this *ISystemSupportInfoStatics2) Vtbl() *ISystemSupportInfoStatics2Vtbl {
	return (*ISystemSupportInfoStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemSupportInfoStatics2) Get_LocalDeviceInfo() *ISystemSupportDeviceInfo {
	var _result *ISystemSupportDeviceInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocalDeviceInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type SystemSupportInfo struct {
	RtClass
}

func NewISystemSupportInfoStatics() *ISystemSupportInfoStatics {
	var p *ISystemSupportInfoStatics
	hs := NewHStr("Windows.System.Profile.SystemManufacturers.SystemSupportInfo")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISystemSupportInfoStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewISystemSupportInfoStatics2() *ISystemSupportInfoStatics2 {
	var p *ISystemSupportInfoStatics2
	hs := NewHStr("Windows.System.Profile.SystemManufacturers.SystemSupportInfo")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISystemSupportInfoStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
