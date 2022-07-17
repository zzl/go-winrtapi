package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type DisplayMonitorConnectionKind int32

const (
	DisplayMonitorConnectionKind_Internal DisplayMonitorConnectionKind = 0
	DisplayMonitorConnectionKind_Wired    DisplayMonitorConnectionKind = 1
	DisplayMonitorConnectionKind_Wireless DisplayMonitorConnectionKind = 2
	DisplayMonitorConnectionKind_Virtual  DisplayMonitorConnectionKind = 3
)

// enum
type DisplayMonitorDescriptorKind int32

const (
	DisplayMonitorDescriptorKind_Edid      DisplayMonitorDescriptorKind = 0
	DisplayMonitorDescriptorKind_DisplayId DisplayMonitorDescriptorKind = 1
)

// enum
type DisplayMonitorPhysicalConnectorKind int32

const (
	DisplayMonitorPhysicalConnectorKind_Unknown     DisplayMonitorPhysicalConnectorKind = 0
	DisplayMonitorPhysicalConnectorKind_HD15        DisplayMonitorPhysicalConnectorKind = 1
	DisplayMonitorPhysicalConnectorKind_AnalogTV    DisplayMonitorPhysicalConnectorKind = 2
	DisplayMonitorPhysicalConnectorKind_Dvi         DisplayMonitorPhysicalConnectorKind = 3
	DisplayMonitorPhysicalConnectorKind_Hdmi        DisplayMonitorPhysicalConnectorKind = 4
	DisplayMonitorPhysicalConnectorKind_Lvds        DisplayMonitorPhysicalConnectorKind = 5
	DisplayMonitorPhysicalConnectorKind_Sdi         DisplayMonitorPhysicalConnectorKind = 6
	DisplayMonitorPhysicalConnectorKind_DisplayPort DisplayMonitorPhysicalConnectorKind = 7
)

// enum
type DisplayMonitorUsageKind int32

const (
	DisplayMonitorUsageKind_Standard       DisplayMonitorUsageKind = 0
	DisplayMonitorUsageKind_HeadMounted    DisplayMonitorUsageKind = 1
	DisplayMonitorUsageKind_SpecialPurpose DisplayMonitorUsageKind = 2
)

// interfaces

// 1F6B15D4-1D01-4C51-87E2-6F954A772B59
var IID_IDisplayMonitor = syscall.GUID{0x1F6B15D4, 0x1D01, 0x4C51,
	[8]byte{0x87, 0xE2, 0x6F, 0x95, 0x4A, 0x77, 0x2B, 0x59}}

type IDisplayMonitorInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_DisplayName() string
	Get_ConnectionKind() DisplayMonitorConnectionKind
	Get_PhysicalConnector() DisplayMonitorPhysicalConnectorKind
	Get_DisplayAdapterDeviceId() string
	Get_DisplayAdapterId() unsafe.Pointer
	Get_DisplayAdapterTargetId() uint32
	Get_UsageKind() DisplayMonitorUsageKind
	Get_NativeResolutionInRawPixels() unsafe.Pointer
	Get_PhysicalSizeInInches() *IReference[Size]
	Get_RawDpiX() float32
	Get_RawDpiY() float32
	Get_RedPrimary() Point
	Get_GreenPrimary() Point
	Get_BluePrimary() Point
	Get_WhitePoint() Point
	Get_MaxLuminanceInNits() float32
	Get_MinLuminanceInNits() float32
	Get_MaxAverageFullFrameLuminanceInNits() float32
	GetDescriptor(descriptorKind DisplayMonitorDescriptorKind) []byte
}

type IDisplayMonitorVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId                           uintptr
	Get_DisplayName                        uintptr
	Get_ConnectionKind                     uintptr
	Get_PhysicalConnector                  uintptr
	Get_DisplayAdapterDeviceId             uintptr
	Get_DisplayAdapterId                   uintptr
	Get_DisplayAdapterTargetId             uintptr
	Get_UsageKind                          uintptr
	Get_NativeResolutionInRawPixels        uintptr
	Get_PhysicalSizeInInches               uintptr
	Get_RawDpiX                            uintptr
	Get_RawDpiY                            uintptr
	Get_RedPrimary                         uintptr
	Get_GreenPrimary                       uintptr
	Get_BluePrimary                        uintptr
	Get_WhitePoint                         uintptr
	Get_MaxLuminanceInNits                 uintptr
	Get_MinLuminanceInNits                 uintptr
	Get_MaxAverageFullFrameLuminanceInNits uintptr
	GetDescriptor                          uintptr
}

type IDisplayMonitor struct {
	win32.IInspectable
}

func (this *IDisplayMonitor) Vtbl() *IDisplayMonitorVtbl {
	return (*IDisplayMonitorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDisplayMonitor) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDisplayMonitor) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDisplayMonitor) Get_ConnectionKind() DisplayMonitorConnectionKind {
	var _result DisplayMonitorConnectionKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConnectionKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDisplayMonitor) Get_PhysicalConnector() DisplayMonitorPhysicalConnectorKind {
	var _result DisplayMonitorPhysicalConnectorKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhysicalConnector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDisplayMonitor) Get_DisplayAdapterDeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayAdapterDeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDisplayMonitor) Get_DisplayAdapterId() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayAdapterId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDisplayMonitor) Get_DisplayAdapterTargetId() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayAdapterTargetId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDisplayMonitor) Get_UsageKind() DisplayMonitorUsageKind {
	var _result DisplayMonitorUsageKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UsageKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDisplayMonitor) Get_NativeResolutionInRawPixels() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NativeResolutionInRawPixels, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDisplayMonitor) Get_PhysicalSizeInInches() *IReference[Size] {
	var _result *IReference[Size]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhysicalSizeInInches, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDisplayMonitor) Get_RawDpiX() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RawDpiX, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDisplayMonitor) Get_RawDpiY() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RawDpiY, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDisplayMonitor) Get_RedPrimary() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RedPrimary, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDisplayMonitor) Get_GreenPrimary() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GreenPrimary, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDisplayMonitor) Get_BluePrimary() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BluePrimary, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDisplayMonitor) Get_WhitePoint() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WhitePoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDisplayMonitor) Get_MaxLuminanceInNits() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxLuminanceInNits, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDisplayMonitor) Get_MinLuminanceInNits() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinLuminanceInNits, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDisplayMonitor) Get_MaxAverageFullFrameLuminanceInNits() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxAverageFullFrameLuminanceInNits, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDisplayMonitor) GetDescriptor(descriptorKind DisplayMonitorDescriptorKind) []byte {
	var _result []byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDescriptor, uintptr(unsafe.Pointer(this)), uintptr(descriptorKind), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 023018E6-CB23-5830-96DF-A7BF6E602577
var IID_IDisplayMonitor2 = syscall.GUID{0x023018E6, 0xCB23, 0x5830,
	[8]byte{0x96, 0xDF, 0xA7, 0xBF, 0x6E, 0x60, 0x25, 0x77}}

type IDisplayMonitor2Interface interface {
	win32.IInspectableInterface
	Get_IsDolbyVisionSupportedInHdrMode() bool
}

type IDisplayMonitor2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsDolbyVisionSupportedInHdrMode uintptr
}

type IDisplayMonitor2 struct {
	win32.IInspectable
}

func (this *IDisplayMonitor2) Vtbl() *IDisplayMonitor2Vtbl {
	return (*IDisplayMonitor2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDisplayMonitor2) Get_IsDolbyVisionSupportedInHdrMode() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDolbyVisionSupportedInHdrMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 6EAE698F-A228-4C05-821D-B695D667DE8E
var IID_IDisplayMonitorStatics = syscall.GUID{0x6EAE698F, 0xA228, 0x4C05,
	[8]byte{0x82, 0x1D, 0xB6, 0x95, 0xD6, 0x67, 0xDE, 0x8E}}

type IDisplayMonitorStaticsInterface interface {
	win32.IInspectableInterface
	GetDeviceSelector() string
	FromIdAsync(deviceId string) *IAsyncOperation[*IDisplayMonitor]
	FromInterfaceIdAsync(deviceInterfaceId string) *IAsyncOperation[*IDisplayMonitor]
}

type IDisplayMonitorStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelector    uintptr
	FromIdAsync          uintptr
	FromInterfaceIdAsync uintptr
}

type IDisplayMonitorStatics struct {
	win32.IInspectable
}

func (this *IDisplayMonitorStatics) Vtbl() *IDisplayMonitorStaticsVtbl {
	return (*IDisplayMonitorStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDisplayMonitorStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDisplayMonitorStatics) FromIdAsync(deviceId string) *IAsyncOperation[*IDisplayMonitor] {
	var _result *IAsyncOperation[*IDisplayMonitor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDisplayMonitorStatics) FromInterfaceIdAsync(deviceInterfaceId string) *IAsyncOperation[*IDisplayMonitor] {
	var _result *IAsyncOperation[*IDisplayMonitor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromInterfaceIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceInterfaceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type DisplayMonitor struct {
	RtClass
	*IDisplayMonitor
}

func NewIDisplayMonitorStatics() *IDisplayMonitorStatics {
	var p *IDisplayMonitorStatics
	hs := NewHStr("Windows.Devices.Display.DisplayMonitor")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IDisplayMonitorStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
