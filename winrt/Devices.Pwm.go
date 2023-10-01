package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type PwmPulsePolarity int32

const (
	PwmPulsePolarity_ActiveHigh PwmPulsePolarity = 0
	PwmPulsePolarity_ActiveLow  PwmPulsePolarity = 1
)

// interfaces

// C45F5C85-D2E8-42CF-9BD6-CF5ED029E6A7
var IID_IPwmController = syscall.GUID{0xC45F5C85, 0xD2E8, 0x42CF,
	[8]byte{0x9B, 0xD6, 0xCF, 0x5E, 0xD0, 0x29, 0xE6, 0xA7}}

type IPwmControllerInterface interface {
	win32.IInspectableInterface
	Get_PinCount() int32
	Get_ActualFrequency() float64
	SetDesiredFrequency(desiredFrequency float64) float64
	Get_MinFrequency() float64
	Get_MaxFrequency() float64
	OpenPin(pinNumber int32) *IPwmPin
}

type IPwmControllerVtbl struct {
	win32.IInspectableVtbl
	Get_PinCount        uintptr
	Get_ActualFrequency uintptr
	SetDesiredFrequency uintptr
	Get_MinFrequency    uintptr
	Get_MaxFrequency    uintptr
	OpenPin             uintptr
}

type IPwmController struct {
	win32.IInspectable
}

func (this *IPwmController) Vtbl() *IPwmControllerVtbl {
	return (*IPwmControllerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPwmController) Get_PinCount() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PinCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPwmController) Get_ActualFrequency() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ActualFrequency, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPwmController) SetDesiredFrequency(desiredFrequency float64) float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetDesiredFrequency, uintptr(unsafe.Pointer(this)), uintptr(desiredFrequency), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPwmController) Get_MinFrequency() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinFrequency, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPwmController) Get_MaxFrequency() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxFrequency, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPwmController) OpenPin(pinNumber int32) *IPwmPin {
	var _result *IPwmPin
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenPin, uintptr(unsafe.Pointer(this)), uintptr(pinNumber), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4263BDA1-8946-4404-BD48-81DD124AF4D9
var IID_IPwmControllerStatics = syscall.GUID{0x4263BDA1, 0x8946, 0x4404,
	[8]byte{0xBD, 0x48, 0x81, 0xDD, 0x12, 0x4A, 0xF4, 0xD9}}

type IPwmControllerStaticsInterface interface {
	win32.IInspectableInterface
	GetControllersAsync(provider unsafe.Pointer) *IAsyncOperation[*IVectorView[*IPwmController]]
}

type IPwmControllerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetControllersAsync uintptr
}

type IPwmControllerStatics struct {
	win32.IInspectable
}

func (this *IPwmControllerStatics) Vtbl() *IPwmControllerStaticsVtbl {
	return (*IPwmControllerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPwmControllerStatics) GetControllersAsync(provider unsafe.Pointer) *IAsyncOperation[*IVectorView[*IPwmController]] {
	var _result *IAsyncOperation[*IVectorView[*IPwmController]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetControllersAsync, uintptr(unsafe.Pointer(this)), uintptr(provider), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 44FC5B1F-F119-4BDD-97AD-F76EF986736D
var IID_IPwmControllerStatics2 = syscall.GUID{0x44FC5B1F, 0xF119, 0x4BDD,
	[8]byte{0x97, 0xAD, 0xF7, 0x6E, 0xF9, 0x86, 0x73, 0x6D}}

type IPwmControllerStatics2Interface interface {
	win32.IInspectableInterface
	GetDefaultAsync() *IAsyncOperation[*IPwmController]
}

type IPwmControllerStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetDefaultAsync uintptr
}

type IPwmControllerStatics2 struct {
	win32.IInspectable
}

func (this *IPwmControllerStatics2) Vtbl() *IPwmControllerStatics2Vtbl {
	return (*IPwmControllerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPwmControllerStatics2) GetDefaultAsync() *IAsyncOperation[*IPwmController] {
	var _result *IAsyncOperation[*IPwmController]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefaultAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B2581871-0229-4344-AE3F-9B7CD0E66B94
var IID_IPwmControllerStatics3 = syscall.GUID{0xB2581871, 0x0229, 0x4344,
	[8]byte{0xAE, 0x3F, 0x9B, 0x7C, 0xD0, 0xE6, 0x6B, 0x94}}

type IPwmControllerStatics3Interface interface {
	win32.IInspectableInterface
	GetDeviceSelector() string
	GetDeviceSelectorFromFriendlyName(friendlyName string) string
	FromIdAsync(deviceId string) *IAsyncOperation[*IPwmController]
}

type IPwmControllerStatics3Vtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelector                 uintptr
	GetDeviceSelectorFromFriendlyName uintptr
	FromIdAsync                       uintptr
}

type IPwmControllerStatics3 struct {
	win32.IInspectable
}

func (this *IPwmControllerStatics3) Vtbl() *IPwmControllerStatics3Vtbl {
	return (*IPwmControllerStatics3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPwmControllerStatics3) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPwmControllerStatics3) GetDeviceSelectorFromFriendlyName(friendlyName string) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorFromFriendlyName, uintptr(unsafe.Pointer(this)), NewHStr(friendlyName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPwmControllerStatics3) FromIdAsync(deviceId string) *IAsyncOperation[*IPwmController] {
	var _result *IAsyncOperation[*IPwmController]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 22972DC8-C6CF-4821-B7F9-C6454FB6AF79
var IID_IPwmPin = syscall.GUID{0x22972DC8, 0xC6CF, 0x4821,
	[8]byte{0xB7, 0xF9, 0xC6, 0x45, 0x4F, 0xB6, 0xAF, 0x79}}

type IPwmPinInterface interface {
	win32.IInspectableInterface
	Get_Controller() *IPwmController
	GetActiveDutyCyclePercentage() float64
	SetActiveDutyCyclePercentage(dutyCyclePercentage float64)
	Get_Polarity() PwmPulsePolarity
	Put_Polarity(value PwmPulsePolarity)
	Start()
	Stop()
	Get_IsStarted() bool
}

type IPwmPinVtbl struct {
	win32.IInspectableVtbl
	Get_Controller               uintptr
	GetActiveDutyCyclePercentage uintptr
	SetActiveDutyCyclePercentage uintptr
	Get_Polarity                 uintptr
	Put_Polarity                 uintptr
	Start                        uintptr
	Stop                         uintptr
	Get_IsStarted                uintptr
}

type IPwmPin struct {
	win32.IInspectable
}

func (this *IPwmPin) Vtbl() *IPwmPinVtbl {
	return (*IPwmPinVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPwmPin) Get_Controller() *IPwmController {
	var _result *IPwmController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Controller, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPwmPin) GetActiveDutyCyclePercentage() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetActiveDutyCyclePercentage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPwmPin) SetActiveDutyCyclePercentage(dutyCyclePercentage float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetActiveDutyCyclePercentage, uintptr(unsafe.Pointer(this)), uintptr(dutyCyclePercentage))
	_ = _hr
}

func (this *IPwmPin) Get_Polarity() PwmPulsePolarity {
	var _result PwmPulsePolarity
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Polarity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPwmPin) Put_Polarity(value PwmPulsePolarity) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Polarity, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPwmPin) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IPwmPin) Stop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IPwmPin) Get_IsStarted() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsStarted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type PwmController struct {
	RtClass
	*IPwmController
}

func NewIPwmControllerStatics2() *IPwmControllerStatics2 {
	var p *IPwmControllerStatics2
	hs := NewHStr("Windows.Devices.Pwm.PwmController")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPwmControllerStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIPwmControllerStatics() *IPwmControllerStatics {
	var p *IPwmControllerStatics
	hs := NewHStr("Windows.Devices.Pwm.PwmController")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPwmControllerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIPwmControllerStatics3() *IPwmControllerStatics3 {
	var p *IPwmControllerStatics3
	hs := NewHStr("Windows.Devices.Pwm.PwmController")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPwmControllerStatics3, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type PwmPin struct {
	RtClass
	*IPwmPin
}
