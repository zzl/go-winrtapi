package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type RadioAccessStatus int32

const (
	RadioAccessStatus_Unspecified    RadioAccessStatus = 0
	RadioAccessStatus_Allowed        RadioAccessStatus = 1
	RadioAccessStatus_DeniedByUser   RadioAccessStatus = 2
	RadioAccessStatus_DeniedBySystem RadioAccessStatus = 3
)

// enum
type RadioKind int32

const (
	RadioKind_Other           RadioKind = 0
	RadioKind_WiFi            RadioKind = 1
	RadioKind_MobileBroadband RadioKind = 2
	RadioKind_Bluetooth       RadioKind = 3
	RadioKind_FM              RadioKind = 4
)

// enum
type RadioState int32

const (
	RadioState_Unknown  RadioState = 0
	RadioState_On       RadioState = 1
	RadioState_Off      RadioState = 2
	RadioState_Disabled RadioState = 3
)

// interfaces

// 252118DF-B33E-416A-875F-1CF38AE2D83E
var IID_IRadio = syscall.GUID{0x252118DF, 0xB33E, 0x416A,
	[8]byte{0x87, 0x5F, 0x1C, 0xF3, 0x8A, 0xE2, 0xD8, 0x3E}}

type IRadioInterface interface {
	win32.IInspectableInterface
	SetStateAsync(value RadioState) *IAsyncOperation[RadioAccessStatus]
	Add_StateChanged(handler TypedEventHandler[*IRadio, interface{}]) EventRegistrationToken
	Remove_StateChanged(eventCookie EventRegistrationToken)
	Get_State() RadioState
	Get_Name() string
	Get_Kind() RadioKind
}

type IRadioVtbl struct {
	win32.IInspectableVtbl
	SetStateAsync       uintptr
	Add_StateChanged    uintptr
	Remove_StateChanged uintptr
	Get_State           uintptr
	Get_Name            uintptr
	Get_Kind            uintptr
}

type IRadio struct {
	win32.IInspectable
}

func (this *IRadio) Vtbl() *IRadioVtbl {
	return (*IRadioVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadio) SetStateAsync(value RadioState) *IAsyncOperation[RadioAccessStatus] {
	var _result *IAsyncOperation[RadioAccessStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetStateAsync, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRadio) Add_StateChanged(handler TypedEventHandler[*IRadio, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRadio) Remove_StateChanged(eventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&eventCookie)))
	_ = _hr
}

func (this *IRadio) Get_State() RadioState {
	var _result RadioState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRadio) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IRadio) Get_Kind() RadioKind {
	var _result RadioKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Kind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5FB6A12E-67CB-46AE-AAE9-65919F86EFF4
var IID_IRadioStatics = syscall.GUID{0x5FB6A12E, 0x67CB, 0x46AE,
	[8]byte{0xAA, 0xE9, 0x65, 0x91, 0x9F, 0x86, 0xEF, 0xF4}}

type IRadioStaticsInterface interface {
	win32.IInspectableInterface
	GetRadiosAsync() *IAsyncOperation[*IVectorView[*IRadio]]
	GetDeviceSelector() string
	FromIdAsync(deviceId string) *IAsyncOperation[*IRadio]
	RequestAccessAsync() *IAsyncOperation[RadioAccessStatus]
}

type IRadioStaticsVtbl struct {
	win32.IInspectableVtbl
	GetRadiosAsync     uintptr
	GetDeviceSelector  uintptr
	FromIdAsync        uintptr
	RequestAccessAsync uintptr
}

type IRadioStatics struct {
	win32.IInspectable
}

func (this *IRadioStatics) Vtbl() *IRadioStaticsVtbl {
	return (*IRadioStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadioStatics) GetRadiosAsync() *IAsyncOperation[*IVectorView[*IRadio]] {
	var _result *IAsyncOperation[*IVectorView[*IRadio]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetRadiosAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRadioStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IRadioStatics) FromIdAsync(deviceId string) *IAsyncOperation[*IRadio] {
	var _result *IAsyncOperation[*IRadio]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRadioStatics) RequestAccessAsync() *IAsyncOperation[RadioAccessStatus] {
	var _result *IAsyncOperation[RadioAccessStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type Radio struct {
	RtClass
	*IRadio
}

func NewIRadioStatics() *IRadioStatics {
	var p *IRadioStatics
	hs := NewHStr("Windows.Devices.Radios.Radio")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IRadioStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
