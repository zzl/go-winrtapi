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
type GpioChangePolarity int32

const (
	GpioChangePolarity_Falling GpioChangePolarity = 0
	GpioChangePolarity_Rising  GpioChangePolarity = 1
	GpioChangePolarity_Both    GpioChangePolarity = 2
)

// enum
type GpioOpenStatus int32

const (
	GpioOpenStatus_PinOpened        GpioOpenStatus = 0
	GpioOpenStatus_PinUnavailable   GpioOpenStatus = 1
	GpioOpenStatus_SharingViolation GpioOpenStatus = 2
	GpioOpenStatus_MuxingConflict   GpioOpenStatus = 3
	GpioOpenStatus_UnknownError     GpioOpenStatus = 4
)

// enum
type GpioPinDriveMode int32

const (
	GpioPinDriveMode_Input                    GpioPinDriveMode = 0
	GpioPinDriveMode_Output                   GpioPinDriveMode = 1
	GpioPinDriveMode_InputPullUp              GpioPinDriveMode = 2
	GpioPinDriveMode_InputPullDown            GpioPinDriveMode = 3
	GpioPinDriveMode_OutputOpenDrain          GpioPinDriveMode = 4
	GpioPinDriveMode_OutputOpenDrainPullUp    GpioPinDriveMode = 5
	GpioPinDriveMode_OutputOpenSource         GpioPinDriveMode = 6
	GpioPinDriveMode_OutputOpenSourcePullDown GpioPinDriveMode = 7
)

// enum
type GpioPinEdge int32

const (
	GpioPinEdge_FallingEdge GpioPinEdge = 0
	GpioPinEdge_RisingEdge  GpioPinEdge = 1
)

// enum
type GpioPinValue int32

const (
	GpioPinValue_Low  GpioPinValue = 0
	GpioPinValue_High GpioPinValue = 1
)

// enum
type GpioSharingMode int32

const (
	GpioSharingMode_Exclusive      GpioSharingMode = 0
	GpioSharingMode_SharedReadOnly GpioSharingMode = 1
)

// structs

type GpioChangeCount struct {
	Count        uint64
	RelativeTime TimeSpan
}

type GpioChangeRecord struct {
	RelativeTime TimeSpan
	Edge         GpioPinEdge
}

// interfaces

// CB5EC0DE-6801-43FF-803D-4576628A8B26
var IID_IGpioChangeCounter = syscall.GUID{0xCB5EC0DE, 0x6801, 0x43FF,
	[8]byte{0x80, 0x3D, 0x45, 0x76, 0x62, 0x8A, 0x8B, 0x26}}

type IGpioChangeCounterInterface interface {
	win32.IInspectableInterface
	Put_Polarity(value GpioChangePolarity)
	Get_Polarity() GpioChangePolarity
	Get_IsStarted() bool
	Start()
	Stop()
	Read() GpioChangeCount
	Reset() GpioChangeCount
}

type IGpioChangeCounterVtbl struct {
	win32.IInspectableVtbl
	Put_Polarity  uintptr
	Get_Polarity  uintptr
	Get_IsStarted uintptr
	Start         uintptr
	Stop          uintptr
	Read          uintptr
	Reset         uintptr
}

type IGpioChangeCounter struct {
	win32.IInspectable
}

func (this *IGpioChangeCounter) Vtbl() *IGpioChangeCounterVtbl {
	return (*IGpioChangeCounterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGpioChangeCounter) Put_Polarity(value GpioChangePolarity) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Polarity, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGpioChangeCounter) Get_Polarity() GpioChangePolarity {
	var _result GpioChangePolarity
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Polarity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGpioChangeCounter) Get_IsStarted() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsStarted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGpioChangeCounter) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IGpioChangeCounter) Stop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IGpioChangeCounter) Read() GpioChangeCount {
	var _result GpioChangeCount
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Read, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGpioChangeCounter) Reset() GpioChangeCount {
	var _result GpioChangeCount
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Reset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 147D94B6-0A9E-410C-B4FA-F89F4052084D
var IID_IGpioChangeCounterFactory = syscall.GUID{0x147D94B6, 0x0A9E, 0x410C,
	[8]byte{0xB4, 0xFA, 0xF8, 0x9F, 0x40, 0x52, 0x08, 0x4D}}

type IGpioChangeCounterFactoryInterface interface {
	win32.IInspectableInterface
	Create(pin *IGpioPin) *IGpioChangeCounter
}

type IGpioChangeCounterFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IGpioChangeCounterFactory struct {
	win32.IInspectable
}

func (this *IGpioChangeCounterFactory) Vtbl() *IGpioChangeCounterFactoryVtbl {
	return (*IGpioChangeCounterFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGpioChangeCounterFactory) Create(pin *IGpioPin) *IGpioChangeCounter {
	var _result *IGpioChangeCounter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pin)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0ABC885F-E031-48E8-8590-70DE78363C6D
var IID_IGpioChangeReader = syscall.GUID{0x0ABC885F, 0xE031, 0x48E8,
	[8]byte{0x85, 0x90, 0x70, 0xDE, 0x78, 0x36, 0x3C, 0x6D}}

type IGpioChangeReaderInterface interface {
	win32.IInspectableInterface
	Get_Capacity() int32
	Get_Length() int32
	Get_IsEmpty() bool
	Get_IsOverflowed() bool
	Put_Polarity(value GpioChangePolarity)
	Get_Polarity() GpioChangePolarity
	Get_IsStarted() bool
	Start()
	Stop()
	Clear()
	GetNextItem() GpioChangeRecord
	PeekNextItem() GpioChangeRecord
	GetAllItems() *IVector[GpioChangeRecord]
	WaitForItemsAsync(count int32) *IAsyncAction
}

type IGpioChangeReaderVtbl struct {
	win32.IInspectableVtbl
	Get_Capacity      uintptr
	Get_Length        uintptr
	Get_IsEmpty       uintptr
	Get_IsOverflowed  uintptr
	Put_Polarity      uintptr
	Get_Polarity      uintptr
	Get_IsStarted     uintptr
	Start             uintptr
	Stop              uintptr
	Clear             uintptr
	GetNextItem       uintptr
	PeekNextItem      uintptr
	GetAllItems       uintptr
	WaitForItemsAsync uintptr
}

type IGpioChangeReader struct {
	win32.IInspectable
}

func (this *IGpioChangeReader) Vtbl() *IGpioChangeReaderVtbl {
	return (*IGpioChangeReaderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGpioChangeReader) Get_Capacity() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Capacity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGpioChangeReader) Get_Length() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Length, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGpioChangeReader) Get_IsEmpty() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEmpty, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGpioChangeReader) Get_IsOverflowed() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsOverflowed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGpioChangeReader) Put_Polarity(value GpioChangePolarity) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Polarity, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGpioChangeReader) Get_Polarity() GpioChangePolarity {
	var _result GpioChangePolarity
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Polarity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGpioChangeReader) Get_IsStarted() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsStarted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGpioChangeReader) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IGpioChangeReader) Stop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IGpioChangeReader) Clear() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Clear, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IGpioChangeReader) GetNextItem() GpioChangeRecord {
	var _result GpioChangeRecord
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNextItem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGpioChangeReader) PeekNextItem() GpioChangeRecord {
	var _result GpioChangeRecord
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PeekNextItem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGpioChangeReader) GetAllItems() *IVector[GpioChangeRecord] {
	var _result *IVector[GpioChangeRecord]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAllItems, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGpioChangeReader) WaitForItemsAsync(count int32) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WaitForItemsAsync, uintptr(unsafe.Pointer(this)), uintptr(count), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A9598EF3-390E-441A-9D1C-E8DE0B2DF0DF
var IID_IGpioChangeReaderFactory = syscall.GUID{0xA9598EF3, 0x390E, 0x441A,
	[8]byte{0x9D, 0x1C, 0xE8, 0xDE, 0x0B, 0x2D, 0xF0, 0xDF}}

type IGpioChangeReaderFactoryInterface interface {
	win32.IInspectableInterface
	Create(pin *IGpioPin) *IGpioChangeReader
	CreateWithCapacity(pin *IGpioPin, minCapacity int32) *IGpioChangeReader
}

type IGpioChangeReaderFactoryVtbl struct {
	win32.IInspectableVtbl
	Create             uintptr
	CreateWithCapacity uintptr
}

type IGpioChangeReaderFactory struct {
	win32.IInspectable
}

func (this *IGpioChangeReaderFactory) Vtbl() *IGpioChangeReaderFactoryVtbl {
	return (*IGpioChangeReaderFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGpioChangeReaderFactory) Create(pin *IGpioPin) *IGpioChangeReader {
	var _result *IGpioChangeReader
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pin)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGpioChangeReaderFactory) CreateWithCapacity(pin *IGpioPin, minCapacity int32) *IGpioChangeReader {
	var _result *IGpioChangeReader
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithCapacity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pin)), uintptr(minCapacity), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 284012E3-7461-469C-A8BC-61D69D08A53C
var IID_IGpioController = syscall.GUID{0x284012E3, 0x7461, 0x469C,
	[8]byte{0xA8, 0xBC, 0x61, 0xD6, 0x9D, 0x08, 0xA5, 0x3C}}

type IGpioControllerInterface interface {
	win32.IInspectableInterface
	Get_PinCount() int32
	OpenPin(pinNumber int32) *IGpioPin
	OpenPinWithSharingMode(pinNumber int32, sharingMode GpioSharingMode) *IGpioPin
	TryOpenPin(pinNumber int32, sharingMode GpioSharingMode, pin *IGpioPin, openStatus GpioOpenStatus) bool
}

type IGpioControllerVtbl struct {
	win32.IInspectableVtbl
	Get_PinCount           uintptr
	OpenPin                uintptr
	OpenPinWithSharingMode uintptr
	TryOpenPin             uintptr
}

type IGpioController struct {
	win32.IInspectable
}

func (this *IGpioController) Vtbl() *IGpioControllerVtbl {
	return (*IGpioControllerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGpioController) Get_PinCount() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PinCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGpioController) OpenPin(pinNumber int32) *IGpioPin {
	var _result *IGpioPin
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenPin, uintptr(unsafe.Pointer(this)), uintptr(pinNumber), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGpioController) OpenPinWithSharingMode(pinNumber int32, sharingMode GpioSharingMode) *IGpioPin {
	var _result *IGpioPin
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenPinWithSharingMode, uintptr(unsafe.Pointer(this)), uintptr(pinNumber), uintptr(sharingMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGpioController) TryOpenPin(pinNumber int32, sharingMode GpioSharingMode, pin *IGpioPin, openStatus GpioOpenStatus) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryOpenPin, uintptr(unsafe.Pointer(this)), uintptr(pinNumber), uintptr(sharingMode), uintptr(unsafe.Pointer(pin)), uintptr(openStatus), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 2ED6F42E-7AF7-4116-9533-C43D99A1FB64
var IID_IGpioControllerStatics = syscall.GUID{0x2ED6F42E, 0x7AF7, 0x4116,
	[8]byte{0x95, 0x33, 0xC4, 0x3D, 0x99, 0xA1, 0xFB, 0x64}}

type IGpioControllerStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *IGpioController
}

type IGpioControllerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
}

type IGpioControllerStatics struct {
	win32.IInspectable
}

func (this *IGpioControllerStatics) Vtbl() *IGpioControllerStaticsVtbl {
	return (*IGpioControllerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGpioControllerStatics) GetDefault() *IGpioController {
	var _result *IGpioController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 912B7D20-6CA4-4106-A373-FFFD346B0E5B
var IID_IGpioControllerStatics2 = syscall.GUID{0x912B7D20, 0x6CA4, 0x4106,
	[8]byte{0xA3, 0x73, 0xFF, 0xFD, 0x34, 0x6B, 0x0E, 0x5B}}

type IGpioControllerStatics2Interface interface {
	win32.IInspectableInterface
	GetControllersAsync(provider *IGpioProvider) *IAsyncOperation[*IVectorView[*IGpioController]]
	GetDefaultAsync() *IAsyncOperation[*IGpioController]
}

type IGpioControllerStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetControllersAsync uintptr
	GetDefaultAsync     uintptr
}

type IGpioControllerStatics2 struct {
	win32.IInspectable
}

func (this *IGpioControllerStatics2) Vtbl() *IGpioControllerStatics2Vtbl {
	return (*IGpioControllerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGpioControllerStatics2) GetControllersAsync(provider *IGpioProvider) *IAsyncOperation[*IVectorView[*IGpioController]] {
	var _result *IAsyncOperation[*IVectorView[*IGpioController]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetControllersAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(provider)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGpioControllerStatics2) GetDefaultAsync() *IAsyncOperation[*IGpioController] {
	var _result *IAsyncOperation[*IGpioController]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefaultAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 11D9B087-AFAE-4790-9EE9-E0EAC942D201
var IID_IGpioPin = syscall.GUID{0x11D9B087, 0xAFAE, 0x4790,
	[8]byte{0x9E, 0xE9, 0xE0, 0xEA, 0xC9, 0x42, 0xD2, 0x01}}

type IGpioPinInterface interface {
	win32.IInspectableInterface
	Add_ValueChanged(handler TypedEventHandler[*IGpioPin, *IGpioPinValueChangedEventArgs]) EventRegistrationToken
	Remove_ValueChanged(token EventRegistrationToken)
	Get_DebounceTimeout() TimeSpan
	Put_DebounceTimeout(value TimeSpan)
	Get_PinNumber() int32
	Get_SharingMode() GpioSharingMode
	IsDriveModeSupported(driveMode GpioPinDriveMode) bool
	GetDriveMode() GpioPinDriveMode
	SetDriveMode(value GpioPinDriveMode)
	Write(value GpioPinValue)
	Read() GpioPinValue
}

type IGpioPinVtbl struct {
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

type IGpioPin struct {
	win32.IInspectable
}

func (this *IGpioPin) Vtbl() *IGpioPinVtbl {
	return (*IGpioPinVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGpioPin) Add_ValueChanged(handler TypedEventHandler[*IGpioPin, *IGpioPinValueChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ValueChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGpioPin) Remove_ValueChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ValueChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGpioPin) Get_DebounceTimeout() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DebounceTimeout, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGpioPin) Put_DebounceTimeout(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DebounceTimeout, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IGpioPin) Get_PinNumber() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PinNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGpioPin) Get_SharingMode() GpioSharingMode {
	var _result GpioSharingMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SharingMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGpioPin) IsDriveModeSupported(driveMode GpioPinDriveMode) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsDriveModeSupported, uintptr(unsafe.Pointer(this)), uintptr(driveMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGpioPin) GetDriveMode() GpioPinDriveMode {
	var _result GpioPinDriveMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDriveMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGpioPin) SetDriveMode(value GpioPinDriveMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetDriveMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGpioPin) Write(value GpioPinValue) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Write, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGpioPin) Read() GpioPinValue {
	var _result GpioPinValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Read, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 3137AAE1-703D-4059-BD24-B5B25DFFB84E
var IID_IGpioPinValueChangedEventArgs = syscall.GUID{0x3137AAE1, 0x703D, 0x4059,
	[8]byte{0xBD, 0x24, 0xB5, 0xB2, 0x5D, 0xFF, 0xB8, 0x4E}}

type IGpioPinValueChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Edge() GpioPinEdge
}

type IGpioPinValueChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Edge uintptr
}

type IGpioPinValueChangedEventArgs struct {
	win32.IInspectable
}

func (this *IGpioPinValueChangedEventArgs) Vtbl() *IGpioPinValueChangedEventArgsVtbl {
	return (*IGpioPinValueChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGpioPinValueChangedEventArgs) Get_Edge() GpioPinEdge {
	var _result GpioPinEdge
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Edge, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type GpioChangeCounter struct {
	RtClass
	*IGpioChangeCounter
}

func NewGpioChangeCounter_Create(pin *IGpioPin) *GpioChangeCounter {
	hs := NewHStr("Windows.Devices.Gpio.GpioChangeCounter")
	var pFac *IGpioChangeCounterFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGpioChangeCounterFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IGpioChangeCounter
	p = pFac.Create(pin)
	result := &GpioChangeCounter{
		RtClass:            RtClass{PInspect: &p.IInspectable},
		IGpioChangeCounter: p,
	}
	com.AddToScope(result)
	return result
}

type GpioChangeReader struct {
	RtClass
	*IGpioChangeReader
}

func NewGpioChangeReader_Create(pin *IGpioPin) *GpioChangeReader {
	hs := NewHStr("Windows.Devices.Gpio.GpioChangeReader")
	var pFac *IGpioChangeReaderFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGpioChangeReaderFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IGpioChangeReader
	p = pFac.Create(pin)
	result := &GpioChangeReader{
		RtClass:           RtClass{PInspect: &p.IInspectable},
		IGpioChangeReader: p,
	}
	com.AddToScope(result)
	return result
}

func NewGpioChangeReader_CreateWithCapacity(pin *IGpioPin, minCapacity int32) *GpioChangeReader {
	hs := NewHStr("Windows.Devices.Gpio.GpioChangeReader")
	var pFac *IGpioChangeReaderFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGpioChangeReaderFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IGpioChangeReader
	p = pFac.CreateWithCapacity(pin, minCapacity)
	result := &GpioChangeReader{
		RtClass:           RtClass{PInspect: &p.IInspectable},
		IGpioChangeReader: p,
	}
	com.AddToScope(result)
	return result
}

type GpioController struct {
	RtClass
	*IGpioController
}

func NewIGpioControllerStatics() *IGpioControllerStatics {
	var p *IGpioControllerStatics
	hs := NewHStr("Windows.Devices.Gpio.GpioController")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGpioControllerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIGpioControllerStatics2() *IGpioControllerStatics2 {
	var p *IGpioControllerStatics2
	hs := NewHStr("Windows.Devices.Gpio.GpioController")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGpioControllerStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type GpioPin struct {
	RtClass
	*IGpioPin
}

type GpioPinValueChangedEventArgs struct {
	RtClass
	*IGpioPinValueChangedEventArgs
}
