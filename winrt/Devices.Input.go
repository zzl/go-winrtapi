package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type PointerDeviceType int32

const (
	PointerDeviceType_Touch PointerDeviceType = 0
	PointerDeviceType_Pen   PointerDeviceType = 1
	PointerDeviceType_Mouse PointerDeviceType = 2
)

// structs

type MouseDelta struct {
	X int32
	Y int32
}

type PointerDeviceUsage struct {
	UsagePage          uint32
	Usage              uint32
	MinLogical         int32
	MaxLogical         int32
	MinPhysical        int32
	MaxPhysical        int32
	Unit               uint32
	PhysicalMultiplier float32
}

// interfaces

// 3A3F9B56-6798-4BBC-833E-0F34B17C65FF
var IID_IKeyboardCapabilities = syscall.GUID{0x3A3F9B56, 0x6798, 0x4BBC,
	[8]byte{0x83, 0x3E, 0x0F, 0x34, 0xB1, 0x7C, 0x65, 0xFF}}

type IKeyboardCapabilitiesInterface interface {
	win32.IInspectableInterface
	Get_KeyboardPresent() int32
}

type IKeyboardCapabilitiesVtbl struct {
	win32.IInspectableVtbl
	Get_KeyboardPresent uintptr
}

type IKeyboardCapabilities struct {
	win32.IInspectable
}

func (this *IKeyboardCapabilities) Vtbl() *IKeyboardCapabilitiesVtbl {
	return (*IKeyboardCapabilitiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKeyboardCapabilities) Get_KeyboardPresent() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KeyboardPresent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// BCA5E023-7DD9-4B6B-9A92-55D43CB38F73
var IID_IMouseCapabilities = syscall.GUID{0xBCA5E023, 0x7DD9, 0x4B6B,
	[8]byte{0x9A, 0x92, 0x55, 0xD4, 0x3C, 0xB3, 0x8F, 0x73}}

type IMouseCapabilitiesInterface interface {
	win32.IInspectableInterface
	Get_MousePresent() int32
	Get_VerticalWheelPresent() int32
	Get_HorizontalWheelPresent() int32
	Get_SwapButtons() int32
	Get_NumberOfButtons() uint32
}

type IMouseCapabilitiesVtbl struct {
	win32.IInspectableVtbl
	Get_MousePresent           uintptr
	Get_VerticalWheelPresent   uintptr
	Get_HorizontalWheelPresent uintptr
	Get_SwapButtons            uintptr
	Get_NumberOfButtons        uintptr
}

type IMouseCapabilities struct {
	win32.IInspectable
}

func (this *IMouseCapabilities) Vtbl() *IMouseCapabilitiesVtbl {
	return (*IMouseCapabilitiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMouseCapabilities) Get_MousePresent() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MousePresent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMouseCapabilities) Get_VerticalWheelPresent() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VerticalWheelPresent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMouseCapabilities) Get_HorizontalWheelPresent() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HorizontalWheelPresent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMouseCapabilities) Get_SwapButtons() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SwapButtons, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMouseCapabilities) Get_NumberOfButtons() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NumberOfButtons, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 88EDF458-F2C8-49F4-BE1F-C256B388BC11
var IID_IMouseDevice = syscall.GUID{0x88EDF458, 0xF2C8, 0x49F4,
	[8]byte{0xBE, 0x1F, 0xC2, 0x56, 0xB3, 0x88, 0xBC, 0x11}}

type IMouseDeviceInterface interface {
	win32.IInspectableInterface
	Add_MouseMoved(handler TypedEventHandler[*IMouseDevice, *IMouseEventArgs]) EventRegistrationToken
	Remove_MouseMoved(cookie EventRegistrationToken)
}

type IMouseDeviceVtbl struct {
	win32.IInspectableVtbl
	Add_MouseMoved    uintptr
	Remove_MouseMoved uintptr
}

type IMouseDevice struct {
	win32.IInspectable
}

func (this *IMouseDevice) Vtbl() *IMouseDeviceVtbl {
	return (*IMouseDeviceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMouseDevice) Add_MouseMoved(handler TypedEventHandler[*IMouseDevice, *IMouseEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_MouseMoved, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMouseDevice) Remove_MouseMoved(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_MouseMoved, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

// 484A9045-6D70-49DB-8E68-46FFBD17D38D
var IID_IMouseDeviceStatics = syscall.GUID{0x484A9045, 0x6D70, 0x49DB,
	[8]byte{0x8E, 0x68, 0x46, 0xFF, 0xBD, 0x17, 0xD3, 0x8D}}

type IMouseDeviceStaticsInterface interface {
	win32.IInspectableInterface
	GetForCurrentView() *IMouseDevice
}

type IMouseDeviceStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForCurrentView uintptr
}

type IMouseDeviceStatics struct {
	win32.IInspectable
}

func (this *IMouseDeviceStatics) Vtbl() *IMouseDeviceStaticsVtbl {
	return (*IMouseDeviceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMouseDeviceStatics) GetForCurrentView() *IMouseDevice {
	var _result *IMouseDevice
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F625AA5D-2354-4CC7-9230-96941C969FDE
var IID_IMouseEventArgs = syscall.GUID{0xF625AA5D, 0x2354, 0x4CC7,
	[8]byte{0x92, 0x30, 0x96, 0x94, 0x1C, 0x96, 0x9F, 0xDE}}

type IMouseEventArgsInterface interface {
	win32.IInspectableInterface
	Get_MouseDelta() MouseDelta
}

type IMouseEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_MouseDelta uintptr
}

type IMouseEventArgs struct {
	win32.IInspectable
}

func (this *IMouseEventArgs) Vtbl() *IMouseEventArgsVtbl {
	return (*IMouseEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMouseEventArgs) Get_MouseDelta() MouseDelta {
	var _result MouseDelta
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MouseDelta, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 8245C376-1EE3-53F7-B1F7-8334A16F2815
var IID_IPenButtonListener = syscall.GUID{0x8245C376, 0x1EE3, 0x53F7,
	[8]byte{0xB1, 0xF7, 0x83, 0x34, 0xA1, 0x6F, 0x28, 0x15}}

type IPenButtonListenerInterface interface {
	win32.IInspectableInterface
	IsSupported() bool
	Add_IsSupportedChanged(handler TypedEventHandler[*IPenButtonListener, interface{}]) EventRegistrationToken
	Remove_IsSupportedChanged(token EventRegistrationToken)
	Add_TailButtonClicked(handler TypedEventHandler[*IPenButtonListener, *IPenTailButtonClickedEventArgs]) EventRegistrationToken
	Remove_TailButtonClicked(token EventRegistrationToken)
	Add_TailButtonDoubleClicked(handler TypedEventHandler[*IPenButtonListener, *IPenTailButtonDoubleClickedEventArgs]) EventRegistrationToken
	Remove_TailButtonDoubleClicked(token EventRegistrationToken)
	Add_TailButtonLongPressed(handler TypedEventHandler[*IPenButtonListener, *IPenTailButtonLongPressedEventArgs]) EventRegistrationToken
	Remove_TailButtonLongPressed(token EventRegistrationToken)
}

type IPenButtonListenerVtbl struct {
	win32.IInspectableVtbl
	IsSupported                    uintptr
	Add_IsSupportedChanged         uintptr
	Remove_IsSupportedChanged      uintptr
	Add_TailButtonClicked          uintptr
	Remove_TailButtonClicked       uintptr
	Add_TailButtonDoubleClicked    uintptr
	Remove_TailButtonDoubleClicked uintptr
	Add_TailButtonLongPressed      uintptr
	Remove_TailButtonLongPressed   uintptr
}

type IPenButtonListener struct {
	win32.IInspectable
}

func (this *IPenButtonListener) Vtbl() *IPenButtonListenerVtbl {
	return (*IPenButtonListenerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPenButtonListener) IsSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPenButtonListener) Add_IsSupportedChanged(handler TypedEventHandler[*IPenButtonListener, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_IsSupportedChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPenButtonListener) Remove_IsSupportedChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_IsSupportedChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPenButtonListener) Add_TailButtonClicked(handler TypedEventHandler[*IPenButtonListener, *IPenTailButtonClickedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_TailButtonClicked, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPenButtonListener) Remove_TailButtonClicked(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_TailButtonClicked, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPenButtonListener) Add_TailButtonDoubleClicked(handler TypedEventHandler[*IPenButtonListener, *IPenTailButtonDoubleClickedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_TailButtonDoubleClicked, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPenButtonListener) Remove_TailButtonDoubleClicked(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_TailButtonDoubleClicked, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPenButtonListener) Add_TailButtonLongPressed(handler TypedEventHandler[*IPenButtonListener, *IPenTailButtonLongPressedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_TailButtonLongPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPenButtonListener) Remove_TailButtonLongPressed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_TailButtonLongPressed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 19A8A584-862F-5F69-BFEA-05F6584F133F
var IID_IPenButtonListenerStatics = syscall.GUID{0x19A8A584, 0x862F, 0x5F69,
	[8]byte{0xBF, 0xEA, 0x05, 0xF6, 0x58, 0x4F, 0x13, 0x3F}}

type IPenButtonListenerStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *IPenButtonListener
}

type IPenButtonListenerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
}

type IPenButtonListenerStatics struct {
	win32.IInspectable
}

func (this *IPenButtonListenerStatics) Vtbl() *IPenButtonListenerStaticsVtbl {
	return (*IPenButtonListenerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPenButtonListenerStatics) GetDefault() *IPenButtonListener {
	var _result *IPenButtonListener
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 31856EBA-A738-5A8C-B8F6-F97EF68D18EF
var IID_IPenDevice = syscall.GUID{0x31856EBA, 0xA738, 0x5A8C,
	[8]byte{0xB8, 0xF6, 0xF9, 0x7E, 0xF6, 0x8D, 0x18, 0xEF}}

type IPenDeviceInterface interface {
	win32.IInspectableInterface
	Get_PenId() syscall.GUID
}

type IPenDeviceVtbl struct {
	win32.IInspectableVtbl
	Get_PenId uintptr
}

type IPenDevice struct {
	win32.IInspectable
}

func (this *IPenDevice) Vtbl() *IPenDeviceVtbl {
	return (*IPenDeviceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPenDevice) Get_PenId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PenId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 9DFBBE01-0966-5180-BCB4-B85060E39479
var IID_IPenDeviceStatics = syscall.GUID{0x9DFBBE01, 0x0966, 0x5180,
	[8]byte{0xBC, 0xB4, 0xB8, 0x50, 0x60, 0xE3, 0x94, 0x79}}

type IPenDeviceStaticsInterface interface {
	win32.IInspectableInterface
	GetFromPointerId(pointerId uint32) *IPenDevice
}

type IPenDeviceStaticsVtbl struct {
	win32.IInspectableVtbl
	GetFromPointerId uintptr
}

type IPenDeviceStatics struct {
	win32.IInspectable
}

func (this *IPenDeviceStatics) Vtbl() *IPenDeviceStaticsVtbl {
	return (*IPenDeviceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPenDeviceStatics) GetFromPointerId(pointerId uint32) *IPenDevice {
	var _result *IPenDevice
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFromPointerId, uintptr(unsafe.Pointer(this)), uintptr(pointerId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 759F4D90-1DC0-55CB-AD18-B9101456F592
var IID_IPenDockListener = syscall.GUID{0x759F4D90, 0x1DC0, 0x55CB,
	[8]byte{0xAD, 0x18, 0xB9, 0x10, 0x14, 0x56, 0xF5, 0x92}}

type IPenDockListenerInterface interface {
	win32.IInspectableInterface
	IsSupported() bool
	Add_IsSupportedChanged(handler TypedEventHandler[*IPenDockListener, interface{}]) EventRegistrationToken
	Remove_IsSupportedChanged(token EventRegistrationToken)
	Add_Docked(handler TypedEventHandler[*IPenDockListener, *IPenDockedEventArgs]) EventRegistrationToken
	Remove_Docked(token EventRegistrationToken)
	Add_Undocked(handler TypedEventHandler[*IPenDockListener, *IPenUndockedEventArgs]) EventRegistrationToken
	Remove_Undocked(token EventRegistrationToken)
}

type IPenDockListenerVtbl struct {
	win32.IInspectableVtbl
	IsSupported               uintptr
	Add_IsSupportedChanged    uintptr
	Remove_IsSupportedChanged uintptr
	Add_Docked                uintptr
	Remove_Docked             uintptr
	Add_Undocked              uintptr
	Remove_Undocked           uintptr
}

type IPenDockListener struct {
	win32.IInspectable
}

func (this *IPenDockListener) Vtbl() *IPenDockListenerVtbl {
	return (*IPenDockListenerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPenDockListener) IsSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPenDockListener) Add_IsSupportedChanged(handler TypedEventHandler[*IPenDockListener, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_IsSupportedChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPenDockListener) Remove_IsSupportedChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_IsSupportedChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPenDockListener) Add_Docked(handler TypedEventHandler[*IPenDockListener, *IPenDockedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Docked, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPenDockListener) Remove_Docked(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Docked, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPenDockListener) Add_Undocked(handler TypedEventHandler[*IPenDockListener, *IPenUndockedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Undocked, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPenDockListener) Remove_Undocked(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Undocked, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// CAB75E9A-0016-5C72-969E-A97E11992A93
var IID_IPenDockListenerStatics = syscall.GUID{0xCAB75E9A, 0x0016, 0x5C72,
	[8]byte{0x96, 0x9E, 0xA9, 0x7E, 0x11, 0x99, 0x2A, 0x93}}

type IPenDockListenerStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *IPenDockListener
}

type IPenDockListenerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
}

type IPenDockListenerStatics struct {
	win32.IInspectable
}

func (this *IPenDockListenerStatics) Vtbl() *IPenDockListenerStaticsVtbl {
	return (*IPenDockListenerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPenDockListenerStatics) GetDefault() *IPenDockListener {
	var _result *IPenDockListener
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FD4277C6-CA63-5D4E-9ED3-A28A54521C8C
var IID_IPenDockedEventArgs = syscall.GUID{0xFD4277C6, 0xCA63, 0x5D4E,
	[8]byte{0x9E, 0xD3, 0xA2, 0x8A, 0x54, 0x52, 0x1C, 0x8C}}

type IPenDockedEventArgsInterface interface {
	win32.IInspectableInterface
}

type IPenDockedEventArgsVtbl struct {
	win32.IInspectableVtbl
}

type IPenDockedEventArgs struct {
	win32.IInspectable
}

func (this *IPenDockedEventArgs) Vtbl() *IPenDockedEventArgsVtbl {
	return (*IPenDockedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 5D2FB7B6-6AD3-5D3E-AB29-05EA2410E390
var IID_IPenTailButtonClickedEventArgs = syscall.GUID{0x5D2FB7B6, 0x6AD3, 0x5D3E,
	[8]byte{0xAB, 0x29, 0x05, 0xEA, 0x24, 0x10, 0xE3, 0x90}}

type IPenTailButtonClickedEventArgsInterface interface {
	win32.IInspectableInterface
}

type IPenTailButtonClickedEventArgsVtbl struct {
	win32.IInspectableVtbl
}

type IPenTailButtonClickedEventArgs struct {
	win32.IInspectable
}

func (this *IPenTailButtonClickedEventArgs) Vtbl() *IPenTailButtonClickedEventArgsVtbl {
	return (*IPenTailButtonClickedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 846321A2-618A-5478-B04C-B358231DA4A7
var IID_IPenTailButtonDoubleClickedEventArgs = syscall.GUID{0x846321A2, 0x618A, 0x5478,
	[8]byte{0xB0, 0x4C, 0xB3, 0x58, 0x23, 0x1D, 0xA4, 0xA7}}

type IPenTailButtonDoubleClickedEventArgsInterface interface {
	win32.IInspectableInterface
}

type IPenTailButtonDoubleClickedEventArgsVtbl struct {
	win32.IInspectableVtbl
}

type IPenTailButtonDoubleClickedEventArgs struct {
	win32.IInspectable
}

func (this *IPenTailButtonDoubleClickedEventArgs) Vtbl() *IPenTailButtonDoubleClickedEventArgsVtbl {
	return (*IPenTailButtonDoubleClickedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// F37C606E-C60A-5F42-B818-A53112406C13
var IID_IPenTailButtonLongPressedEventArgs = syscall.GUID{0xF37C606E, 0xC60A, 0x5F42,
	[8]byte{0xB8, 0x18, 0xA5, 0x31, 0x12, 0x40, 0x6C, 0x13}}

type IPenTailButtonLongPressedEventArgsInterface interface {
	win32.IInspectableInterface
}

type IPenTailButtonLongPressedEventArgsVtbl struct {
	win32.IInspectableVtbl
}

type IPenTailButtonLongPressedEventArgs struct {
	win32.IInspectable
}

func (this *IPenTailButtonLongPressedEventArgs) Vtbl() *IPenTailButtonLongPressedEventArgsVtbl {
	return (*IPenTailButtonLongPressedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// CCD09150-261B-59E6-A5D4-C1964CD03FEB
var IID_IPenUndockedEventArgs = syscall.GUID{0xCCD09150, 0x261B, 0x59E6,
	[8]byte{0xA5, 0xD4, 0xC1, 0x96, 0x4C, 0xD0, 0x3F, 0xEB}}

type IPenUndockedEventArgsInterface interface {
	win32.IInspectableInterface
}

type IPenUndockedEventArgsVtbl struct {
	win32.IInspectableVtbl
}

type IPenUndockedEventArgs struct {
	win32.IInspectable
}

func (this *IPenUndockedEventArgs) Vtbl() *IPenUndockedEventArgsVtbl {
	return (*IPenUndockedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 93C9BAFC-EBCB-467E-82C6-276FEAE36B5A
var IID_IPointerDevice = syscall.GUID{0x93C9BAFC, 0xEBCB, 0x467E,
	[8]byte{0x82, 0xC6, 0x27, 0x6F, 0xEA, 0xE3, 0x6B, 0x5A}}

type IPointerDeviceInterface interface {
	win32.IInspectableInterface
	Get_PointerDeviceType() PointerDeviceType
	Get_IsIntegrated() bool
	Get_MaxContacts() uint32
	Get_PhysicalDeviceRect() Rect
	Get_ScreenRect() Rect
	Get_SupportedUsages() *IVectorView[PointerDeviceUsage]
}

type IPointerDeviceVtbl struct {
	win32.IInspectableVtbl
	Get_PointerDeviceType  uintptr
	Get_IsIntegrated       uintptr
	Get_MaxContacts        uintptr
	Get_PhysicalDeviceRect uintptr
	Get_ScreenRect         uintptr
	Get_SupportedUsages    uintptr
}

type IPointerDevice struct {
	win32.IInspectable
}

func (this *IPointerDevice) Vtbl() *IPointerDeviceVtbl {
	return (*IPointerDeviceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPointerDevice) Get_PointerDeviceType() PointerDeviceType {
	var _result PointerDeviceType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PointerDeviceType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerDevice) Get_IsIntegrated() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsIntegrated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerDevice) Get_MaxContacts() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxContacts, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerDevice) Get_PhysicalDeviceRect() Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhysicalDeviceRect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerDevice) Get_ScreenRect() Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScreenRect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointerDevice) Get_SupportedUsages() *IVectorView[PointerDeviceUsage] {
	var _result *IVectorView[PointerDeviceUsage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedUsages, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F8A6D2A0-C484-489F-AE3E-30D2EE1FFD3E
var IID_IPointerDevice2 = syscall.GUID{0xF8A6D2A0, 0xC484, 0x489F,
	[8]byte{0xAE, 0x3E, 0x30, 0xD2, 0xEE, 0x1F, 0xFD, 0x3E}}

type IPointerDevice2Interface interface {
	win32.IInspectableInterface
	Get_MaxPointersWithZDistance() uint32
}

type IPointerDevice2Vtbl struct {
	win32.IInspectableVtbl
	Get_MaxPointersWithZDistance uintptr
}

type IPointerDevice2 struct {
	win32.IInspectable
}

func (this *IPointerDevice2) Vtbl() *IPointerDevice2Vtbl {
	return (*IPointerDevice2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPointerDevice2) Get_MaxPointersWithZDistance() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxPointersWithZDistance, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D8B89AA1-D1C6-416E-BD8D-5790914DC563
var IID_IPointerDeviceStatics = syscall.GUID{0xD8B89AA1, 0xD1C6, 0x416E,
	[8]byte{0xBD, 0x8D, 0x57, 0x90, 0x91, 0x4D, 0xC5, 0x63}}

type IPointerDeviceStaticsInterface interface {
	win32.IInspectableInterface
	GetPointerDevice(pointerId uint32) *IPointerDevice
	GetPointerDevices() *IVectorView[*IPointerDevice]
}

type IPointerDeviceStaticsVtbl struct {
	win32.IInspectableVtbl
	GetPointerDevice  uintptr
	GetPointerDevices uintptr
}

type IPointerDeviceStatics struct {
	win32.IInspectable
}

func (this *IPointerDeviceStatics) Vtbl() *IPointerDeviceStaticsVtbl {
	return (*IPointerDeviceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPointerDeviceStatics) GetPointerDevice(pointerId uint32) *IPointerDevice {
	var _result *IPointerDevice
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPointerDevice, uintptr(unsafe.Pointer(this)), uintptr(pointerId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPointerDeviceStatics) GetPointerDevices() *IVectorView[*IPointerDevice] {
	var _result *IVectorView[*IPointerDevice]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPointerDevices, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 20DD55F9-13F1-46C8-9285-2C05FA3EDA6F
var IID_ITouchCapabilities = syscall.GUID{0x20DD55F9, 0x13F1, 0x46C8,
	[8]byte{0x92, 0x85, 0x2C, 0x05, 0xFA, 0x3E, 0xDA, 0x6F}}

type ITouchCapabilitiesInterface interface {
	win32.IInspectableInterface
	Get_TouchPresent() int32
	Get_Contacts() uint32
}

type ITouchCapabilitiesVtbl struct {
	win32.IInspectableVtbl
	Get_TouchPresent uintptr
	Get_Contacts     uintptr
}

type ITouchCapabilities struct {
	win32.IInspectable
}

func (this *ITouchCapabilities) Vtbl() *ITouchCapabilitiesVtbl {
	return (*ITouchCapabilitiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITouchCapabilities) Get_TouchPresent() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TouchPresent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITouchCapabilities) Get_Contacts() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Contacts, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type PenButtonListener struct {
	RtClass
	*IPenButtonListener
}

func NewIPenButtonListenerStatics() *IPenButtonListenerStatics {
	var p *IPenButtonListenerStatics
	hs := NewHStr("Windows.Devices.Input.PenButtonListener")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPenButtonListenerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type PenDevice struct {
	RtClass
	*IPenDevice
}

func NewIPenDeviceStatics() *IPenDeviceStatics {
	var p *IPenDeviceStatics
	hs := NewHStr("Windows.Devices.Input.PenDevice")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPenDeviceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type PenDockListener struct {
	RtClass
	*IPenDockListener
}

func NewIPenDockListenerStatics() *IPenDockListenerStatics {
	var p *IPenDockListenerStatics
	hs := NewHStr("Windows.Devices.Input.PenDockListener")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPenDockListenerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type PenDockedEventArgs struct {
	RtClass
	*IPenDockedEventArgs
}

type PenTailButtonClickedEventArgs struct {
	RtClass
	*IPenTailButtonClickedEventArgs
}

type PenTailButtonDoubleClickedEventArgs struct {
	RtClass
	*IPenTailButtonDoubleClickedEventArgs
}

type PenTailButtonLongPressedEventArgs struct {
	RtClass
	*IPenTailButtonLongPressedEventArgs
}

type PenUndockedEventArgs struct {
	RtClass
	*IPenUndockedEventArgs
}
