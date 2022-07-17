package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type AdcChannelMode int32

const (
	AdcChannelMode_SingleEnded  AdcChannelMode = 0
	AdcChannelMode_Differential AdcChannelMode = 1
)

// interfaces

// 040BF414-2588-4A56-ABEF-73A260ACC60A
var IID_IAdcChannel = syscall.GUID{0x040BF414, 0x2588, 0x4A56,
	[8]byte{0xAB, 0xEF, 0x73, 0xA2, 0x60, 0xAC, 0xC6, 0x0A}}

type IAdcChannelInterface interface {
	win32.IInspectableInterface
	Get_Controller() *IAdcController
	ReadValue() int32
	ReadRatio() float64
}

type IAdcChannelVtbl struct {
	win32.IInspectableVtbl
	Get_Controller uintptr
	ReadValue      uintptr
	ReadRatio      uintptr
}

type IAdcChannel struct {
	win32.IInspectable
}

func (this *IAdcChannel) Vtbl() *IAdcChannelVtbl {
	return (*IAdcChannelVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdcChannel) Get_Controller() *IAdcController {
	var _result *IAdcController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Controller, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdcChannel) ReadValue() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdcChannel) ReadRatio() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadRatio, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 2A76E4B0-A896-4219-86B6-EA8CDCE98F56
var IID_IAdcController = syscall.GUID{0x2A76E4B0, 0xA896, 0x4219,
	[8]byte{0x86, 0xB6, 0xEA, 0x8C, 0xDC, 0xE9, 0x8F, 0x56}}

type IAdcControllerInterface interface {
	win32.IInspectableInterface
	Get_ChannelCount() int32
	Get_ResolutionInBits() int32
	Get_MinValue() int32
	Get_MaxValue() int32
	Get_ChannelMode() AdcChannelMode
	Put_ChannelMode(value AdcChannelMode)
	IsChannelModeSupported(channelMode AdcChannelMode) bool
	OpenChannel(channelNumber int32) *IAdcChannel
}

type IAdcControllerVtbl struct {
	win32.IInspectableVtbl
	Get_ChannelCount       uintptr
	Get_ResolutionInBits   uintptr
	Get_MinValue           uintptr
	Get_MaxValue           uintptr
	Get_ChannelMode        uintptr
	Put_ChannelMode        uintptr
	IsChannelModeSupported uintptr
	OpenChannel            uintptr
}

type IAdcController struct {
	win32.IInspectable
}

func (this *IAdcController) Vtbl() *IAdcControllerVtbl {
	return (*IAdcControllerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdcController) Get_ChannelCount() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ChannelCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdcController) Get_ResolutionInBits() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResolutionInBits, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdcController) Get_MinValue() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdcController) Get_MaxValue() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdcController) Get_ChannelMode() AdcChannelMode {
	var _result AdcChannelMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ChannelMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdcController) Put_ChannelMode(value AdcChannelMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ChannelMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAdcController) IsChannelModeSupported(channelMode AdcChannelMode) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsChannelModeSupported, uintptr(unsafe.Pointer(this)), uintptr(channelMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdcController) OpenChannel(channelNumber int32) *IAdcChannel {
	var _result *IAdcChannel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenChannel, uintptr(unsafe.Pointer(this)), uintptr(channelNumber), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CCE98E0C-01F8-4891-BC3B-BE53EF279CA4
var IID_IAdcControllerStatics = syscall.GUID{0xCCE98E0C, 0x01F8, 0x4891,
	[8]byte{0xBC, 0x3B, 0xBE, 0x53, 0xEF, 0x27, 0x9C, 0xA4}}

type IAdcControllerStaticsInterface interface {
	win32.IInspectableInterface
	GetControllersAsync(provider unsafe.Pointer) *IAsyncOperation[*IVectorView[*IAdcController]]
}

type IAdcControllerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetControllersAsync uintptr
}

type IAdcControllerStatics struct {
	win32.IInspectable
}

func (this *IAdcControllerStatics) Vtbl() *IAdcControllerStaticsVtbl {
	return (*IAdcControllerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdcControllerStatics) GetControllersAsync(provider unsafe.Pointer) *IAsyncOperation[*IVectorView[*IAdcController]] {
	var _result *IAsyncOperation[*IVectorView[*IAdcController]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetControllersAsync, uintptr(unsafe.Pointer(this)), uintptr(provider), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A2B93B1D-977B-4F5A-A5FE-A6ABAFFE6484
var IID_IAdcControllerStatics2 = syscall.GUID{0xA2B93B1D, 0x977B, 0x4F5A,
	[8]byte{0xA5, 0xFE, 0xA6, 0xAB, 0xAF, 0xFE, 0x64, 0x84}}

type IAdcControllerStatics2Interface interface {
	win32.IInspectableInterface
	GetDefaultAsync() *IAsyncOperation[*IAdcController]
}

type IAdcControllerStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetDefaultAsync uintptr
}

type IAdcControllerStatics2 struct {
	win32.IInspectable
}

func (this *IAdcControllerStatics2) Vtbl() *IAdcControllerStatics2Vtbl {
	return (*IAdcControllerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdcControllerStatics2) GetDefaultAsync() *IAsyncOperation[*IAdcController] {
	var _result *IAsyncOperation[*IAdcController]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefaultAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type AdcChannel struct {
	RtClass
	*IAdcChannel
}

type AdcController struct {
	RtClass
	*IAdcController
}

func NewIAdcControllerStatics2() *IAdcControllerStatics2 {
	var p *IAdcControllerStatics2
	hs := NewHStr("Windows.Devices.Adc.AdcController")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAdcControllerStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIAdcControllerStatics() *IAdcControllerStatics {
	var p *IAdcControllerStatics
	hs := NewHStr("Windows.Devices.Adc.AdcController")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAdcControllerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
