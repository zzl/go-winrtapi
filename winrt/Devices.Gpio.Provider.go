package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"log"
	"syscall"
	"unsafe"
)

// enums

// enum
type ProviderGpioPinDriveMode int32

const (
	ProviderGpioPinDriveMode_Input                    ProviderGpioPinDriveMode = 0
	ProviderGpioPinDriveMode_Output                   ProviderGpioPinDriveMode = 1
	ProviderGpioPinDriveMode_InputPullUp              ProviderGpioPinDriveMode = 2
	ProviderGpioPinDriveMode_InputPullDown            ProviderGpioPinDriveMode = 3
	ProviderGpioPinDriveMode_OutputOpenDrain          ProviderGpioPinDriveMode = 4
	ProviderGpioPinDriveMode_OutputOpenDrainPullUp    ProviderGpioPinDriveMode = 5
	ProviderGpioPinDriveMode_OutputOpenSource         ProviderGpioPinDriveMode = 6
	ProviderGpioPinDriveMode_OutputOpenSourcePullDown ProviderGpioPinDriveMode = 7
)

// enum
type ProviderGpioPinEdge int32

const (
	ProviderGpioPinEdge_FallingEdge ProviderGpioPinEdge = 0
	ProviderGpioPinEdge_RisingEdge  ProviderGpioPinEdge = 1
)

// enum
type ProviderGpioPinValue int32

const (
	ProviderGpioPinValue_Low  ProviderGpioPinValue = 0
	ProviderGpioPinValue_High ProviderGpioPinValue = 1
)

// enum
type ProviderGpioSharingMode int32

const (
	ProviderGpioSharingMode_Exclusive      ProviderGpioSharingMode = 0
	ProviderGpioSharingMode_SharedReadOnly ProviderGpioSharingMode = 1
)

// interfaces

// AD11CEC7-19EA-4B21-874F-B91AED4A25DB
var IID_IGpioControllerProvider = syscall.GUID{0xAD11CEC7, 0x19EA, 0x4B21,
	[8]byte{0x87, 0x4F, 0xB9, 0x1A, 0xED, 0x4A, 0x25, 0xDB}}

type IGpioControllerProviderInterface interface {
	win32.IInspectableInterface
	Get_PinCount() int32
	OpenPinProvider(pin int32, sharingMode ProviderGpioSharingMode) *IGpioPinProvider
}

type IGpioControllerProviderVtbl struct {
	win32.IInspectableVtbl
	Get_PinCount    uintptr
	OpenPinProvider uintptr
}

type IGpioControllerProvider struct {
	win32.IInspectable
}

func (this *IGpioControllerProvider) Vtbl() *IGpioControllerProviderVtbl {
	return (*IGpioControllerProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGpioControllerProvider) Get_PinCount() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PinCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGpioControllerProvider) OpenPinProvider(pin int32, sharingMode ProviderGpioSharingMode) *IGpioPinProvider {
	var _result *IGpioPinProvider
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenPinProvider, uintptr(unsafe.Pointer(this)), uintptr(pin), uintptr(sharingMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 42344CB7-6ABC-40FF-9CE7-73B85301B900
var IID_IGpioPinProvider = syscall.GUID{0x42344CB7, 0x6ABC, 0x40FF,
	[8]byte{0x9C, 0xE7, 0x73, 0xB8, 0x53, 0x01, 0xB9, 0x00}}

type IGpioPinProviderInterface interface {
	win32.IInspectableInterface
	Add_ValueChanged(handler TypedEventHandler[*IGpioPinProvider, *IGpioPinProviderValueChangedEventArgs]) EventRegistrationToken
	Remove_ValueChanged(token EventRegistrationToken)
	Get_DebounceTimeout() TimeSpan
	Put_DebounceTimeout(value TimeSpan)
	Get_PinNumber() int32
	Get_SharingMode() ProviderGpioSharingMode
	IsDriveModeSupported(driveMode ProviderGpioPinDriveMode) bool
	GetDriveMode() ProviderGpioPinDriveMode
	SetDriveMode(value ProviderGpioPinDriveMode)
	Write(value ProviderGpioPinValue)
	Read() ProviderGpioPinValue
}

type IGpioPinProviderVtbl struct {
	win32.IInspectableVtbl
	Add_ValueChanged     uintptr
	Remove_ValueChanged  uintptr
	Get_DebounceTimeout  uintptr
	Put_DebounceTimeout  uintptr
	Get_PinNumber        uintptr
	Get_SharingMode      uintptr
	IsDriveModeSupported uintptr
	GetDriveMode         uintptr
	SetDriveMode         uintptr
	Write                uintptr
	Read                 uintptr
}

type IGpioPinProvider struct {
	win32.IInspectable
}

func (this *IGpioPinProvider) Vtbl() *IGpioPinProviderVtbl {
	return (*IGpioPinProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGpioPinProvider) Add_ValueChanged(handler TypedEventHandler[*IGpioPinProvider, *IGpioPinProviderValueChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ValueChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGpioPinProvider) Remove_ValueChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ValueChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGpioPinProvider) Get_DebounceTimeout() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DebounceTimeout, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGpioPinProvider) Put_DebounceTimeout(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DebounceTimeout, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IGpioPinProvider) Get_PinNumber() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PinNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGpioPinProvider) Get_SharingMode() ProviderGpioSharingMode {
	var _result ProviderGpioSharingMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SharingMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGpioPinProvider) IsDriveModeSupported(driveMode ProviderGpioPinDriveMode) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsDriveModeSupported, uintptr(unsafe.Pointer(this)), uintptr(driveMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGpioPinProvider) GetDriveMode() ProviderGpioPinDriveMode {
	var _result ProviderGpioPinDriveMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDriveMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGpioPinProvider) SetDriveMode(value ProviderGpioPinDriveMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetDriveMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGpioPinProvider) Write(value ProviderGpioPinValue) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Write, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGpioPinProvider) Read() ProviderGpioPinValue {
	var _result ProviderGpioPinValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Read, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 32A6D6F2-3D5B-44CD-8FBE-13A69F2EDB24
var IID_IGpioPinProviderValueChangedEventArgs = syscall.GUID{0x32A6D6F2, 0x3D5B, 0x44CD,
	[8]byte{0x8F, 0xBE, 0x13, 0xA6, 0x9F, 0x2E, 0xDB, 0x24}}

type IGpioPinProviderValueChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Edge() ProviderGpioPinEdge
}

type IGpioPinProviderValueChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Edge uintptr
}

type IGpioPinProviderValueChangedEventArgs struct {
	win32.IInspectable
}

func (this *IGpioPinProviderValueChangedEventArgs) Vtbl() *IGpioPinProviderValueChangedEventArgsVtbl {
	return (*IGpioPinProviderValueChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGpioPinProviderValueChangedEventArgs) Get_Edge() ProviderGpioPinEdge {
	var _result ProviderGpioPinEdge
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Edge, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 3ECB0B59-568C-4392-B24A-8A59A902B1F1
var IID_IGpioPinProviderValueChangedEventArgsFactory = syscall.GUID{0x3ECB0B59, 0x568C, 0x4392,
	[8]byte{0xB2, 0x4A, 0x8A, 0x59, 0xA9, 0x02, 0xB1, 0xF1}}

type IGpioPinProviderValueChangedEventArgsFactoryInterface interface {
	win32.IInspectableInterface
	Create(edge ProviderGpioPinEdge) *IGpioPinProviderValueChangedEventArgs
}

type IGpioPinProviderValueChangedEventArgsFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IGpioPinProviderValueChangedEventArgsFactory struct {
	win32.IInspectable
}

func (this *IGpioPinProviderValueChangedEventArgsFactory) Vtbl() *IGpioPinProviderValueChangedEventArgsFactoryVtbl {
	return (*IGpioPinProviderValueChangedEventArgsFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGpioPinProviderValueChangedEventArgsFactory) Create(edge ProviderGpioPinEdge) *IGpioPinProviderValueChangedEventArgs {
	var _result *IGpioPinProviderValueChangedEventArgs
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(edge), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 44E82707-08CA-434A-AFE0-D61580446F7E
var IID_IGpioProvider = syscall.GUID{0x44E82707, 0x08CA, 0x434A,
	[8]byte{0xAF, 0xE0, 0xD6, 0x15, 0x80, 0x44, 0x6F, 0x7E}}

type IGpioProviderInterface interface {
	win32.IInspectableInterface
	GetControllers() *IVectorView[*IGpioControllerProvider]
}

type IGpioProviderVtbl struct {
	win32.IInspectableVtbl
	GetControllers uintptr
}

type IGpioProvider struct {
	win32.IInspectable
}

func (this *IGpioProvider) Vtbl() *IGpioProviderVtbl {
	return (*IGpioProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGpioProvider) GetControllers() *IVectorView[*IGpioControllerProvider] {
	var _result *IVectorView[*IGpioControllerProvider]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetControllers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type GpioPinProviderValueChangedEventArgs struct {
	RtClass
	*IGpioPinProviderValueChangedEventArgs
}

func NewGpioPinProviderValueChangedEventArgs_Create(edge ProviderGpioPinEdge) *GpioPinProviderValueChangedEventArgs {
	hs := NewHStr("Windows.Devices.Gpio.Provider.GpioPinProviderValueChangedEventArgs")
	var pFac *IGpioPinProviderValueChangedEventArgsFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGpioPinProviderValueChangedEventArgsFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IGpioPinProviderValueChangedEventArgs
	p = pFac.Create(edge)
	result := &GpioPinProviderValueChangedEventArgs{
		RtClass:                               RtClass{PInspect: &p.IInspectable},
		IGpioPinProviderValueChangedEventArgs: p,
	}
	com.AddToScope(result)
	return result
}
