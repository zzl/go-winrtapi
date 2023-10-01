package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"log"
	"syscall"
	"unsafe"
)

// structs

type AppBroadcastingContract struct {
}

// interfaces

// 00F95A68-8907-48A0-B8EF-24D208137542
var IID_IAppBroadcastingMonitor = syscall.GUID{0x00F95A68, 0x8907, 0x48A0,
	[8]byte{0xB8, 0xEF, 0x24, 0xD2, 0x08, 0x13, 0x75, 0x42}}

type IAppBroadcastingMonitorInterface interface {
	win32.IInspectableInterface
	Get_IsCurrentAppBroadcasting() bool
	Add_IsCurrentAppBroadcastingChanged(handler TypedEventHandler[*IAppBroadcastingMonitor, interface{}]) EventRegistrationToken
	Remove_IsCurrentAppBroadcastingChanged(token EventRegistrationToken)
}

type IAppBroadcastingMonitorVtbl struct {
	win32.IInspectableVtbl
	Get_IsCurrentAppBroadcasting           uintptr
	Add_IsCurrentAppBroadcastingChanged    uintptr
	Remove_IsCurrentAppBroadcastingChanged uintptr
}

type IAppBroadcastingMonitor struct {
	win32.IInspectable
}

func (this *IAppBroadcastingMonitor) Vtbl() *IAppBroadcastingMonitorVtbl {
	return (*IAppBroadcastingMonitorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastingMonitor) Get_IsCurrentAppBroadcasting() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCurrentAppBroadcasting, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastingMonitor) Add_IsCurrentAppBroadcastingChanged(handler TypedEventHandler[*IAppBroadcastingMonitor, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_IsCurrentAppBroadcastingChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastingMonitor) Remove_IsCurrentAppBroadcastingChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_IsCurrentAppBroadcastingChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 1225E4DF-03A1-42F8-8B80-C9228CD9CF2E
var IID_IAppBroadcastingStatus = syscall.GUID{0x1225E4DF, 0x03A1, 0x42F8,
	[8]byte{0x8B, 0x80, 0xC9, 0x22, 0x8C, 0xD9, 0xCF, 0x2E}}

type IAppBroadcastingStatusInterface interface {
	win32.IInspectableInterface
	Get_CanStartBroadcast() bool
	Get_Details() *IAppBroadcastingStatusDetails
}

type IAppBroadcastingStatusVtbl struct {
	win32.IInspectableVtbl
	Get_CanStartBroadcast uintptr
	Get_Details           uintptr
}

type IAppBroadcastingStatus struct {
	win32.IInspectable
}

func (this *IAppBroadcastingStatus) Vtbl() *IAppBroadcastingStatusVtbl {
	return (*IAppBroadcastingStatusVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastingStatus) Get_CanStartBroadcast() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanStartBroadcast, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastingStatus) Get_Details() *IAppBroadcastingStatusDetails {
	var _result *IAppBroadcastingStatusDetails
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Details, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 069DADA4-B573-4E3C-8E19-1BAFACD09713
var IID_IAppBroadcastingStatusDetails = syscall.GUID{0x069DADA4, 0xB573, 0x4E3C,
	[8]byte{0x8E, 0x19, 0x1B, 0xAF, 0xAC, 0xD0, 0x97, 0x13}}

type IAppBroadcastingStatusDetailsInterface interface {
	win32.IInspectableInterface
	Get_IsAnyAppBroadcasting() bool
	Get_IsCaptureResourceUnavailable() bool
	Get_IsGameStreamInProgress() bool
	Get_IsGpuConstrained() bool
	Get_IsAppInactive() bool
	Get_IsBlockedForApp() bool
	Get_IsDisabledByUser() bool
	Get_IsDisabledBySystem() bool
}

type IAppBroadcastingStatusDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_IsAnyAppBroadcasting         uintptr
	Get_IsCaptureResourceUnavailable uintptr
	Get_IsGameStreamInProgress       uintptr
	Get_IsGpuConstrained             uintptr
	Get_IsAppInactive                uintptr
	Get_IsBlockedForApp              uintptr
	Get_IsDisabledByUser             uintptr
	Get_IsDisabledBySystem           uintptr
}

type IAppBroadcastingStatusDetails struct {
	win32.IInspectable
}

func (this *IAppBroadcastingStatusDetails) Vtbl() *IAppBroadcastingStatusDetailsVtbl {
	return (*IAppBroadcastingStatusDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastingStatusDetails) Get_IsAnyAppBroadcasting() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsAnyAppBroadcasting, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastingStatusDetails) Get_IsCaptureResourceUnavailable() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCaptureResourceUnavailable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastingStatusDetails) Get_IsGameStreamInProgress() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsGameStreamInProgress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastingStatusDetails) Get_IsGpuConstrained() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsGpuConstrained, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastingStatusDetails) Get_IsAppInactive() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsAppInactive, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastingStatusDetails) Get_IsBlockedForApp() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsBlockedForApp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastingStatusDetails) Get_IsDisabledByUser() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDisabledByUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastingStatusDetails) Get_IsDisabledBySystem() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDisabledBySystem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// E56F9F8F-EE99-4DCA-A3C3-70AF3DB44F5F
var IID_IAppBroadcastingUI = syscall.GUID{0xE56F9F8F, 0xEE99, 0x4DCA,
	[8]byte{0xA3, 0xC3, 0x70, 0xAF, 0x3D, 0xB4, 0x4F, 0x5F}}

type IAppBroadcastingUIInterface interface {
	win32.IInspectableInterface
	GetStatus() *IAppBroadcastingStatus
	ShowBroadcastUI()
}

type IAppBroadcastingUIVtbl struct {
	win32.IInspectableVtbl
	GetStatus       uintptr
	ShowBroadcastUI uintptr
}

type IAppBroadcastingUI struct {
	win32.IInspectable
}

func (this *IAppBroadcastingUI) Vtbl() *IAppBroadcastingUIVtbl {
	return (*IAppBroadcastingUIVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastingUI) GetStatus() *IAppBroadcastingStatus {
	var _result *IAppBroadcastingStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppBroadcastingUI) ShowBroadcastUI() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowBroadcastUI, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 55A8A79D-23CB-4579-9C34-886FE02C045A
var IID_IAppBroadcastingUIStatics = syscall.GUID{0x55A8A79D, 0x23CB, 0x4579,
	[8]byte{0x9C, 0x34, 0x88, 0x6F, 0xE0, 0x2C, 0x04, 0x5A}}

type IAppBroadcastingUIStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *IAppBroadcastingUI
	GetForUser(user *IUser) *IAppBroadcastingUI
}

type IAppBroadcastingUIStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
	GetForUser uintptr
}

type IAppBroadcastingUIStatics struct {
	win32.IInspectable
}

func (this *IAppBroadcastingUIStatics) Vtbl() *IAppBroadcastingUIStaticsVtbl {
	return (*IAppBroadcastingUIStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastingUIStatics) GetDefault() *IAppBroadcastingUI {
	var _result *IAppBroadcastingUI
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppBroadcastingUIStatics) GetForUser(user *IUser) *IAppBroadcastingUI {
	var _result *IAppBroadcastingUI
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type AppBroadcastingMonitor struct {
	RtClass
	*IAppBroadcastingMonitor
}

func NewAppBroadcastingMonitor() *AppBroadcastingMonitor {
	hs := NewHStr("Windows.Media.AppBroadcasting.AppBroadcastingMonitor")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &AppBroadcastingMonitor{
		RtClass:                 RtClass{PInspect: p},
		IAppBroadcastingMonitor: (*IAppBroadcastingMonitor)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type AppBroadcastingStatus struct {
	RtClass
	*IAppBroadcastingStatus
}

type AppBroadcastingStatusDetails struct {
	RtClass
	*IAppBroadcastingStatusDetails
}

type AppBroadcastingUI struct {
	RtClass
	*IAppBroadcastingUI
}

func NewIAppBroadcastingUIStatics() *IAppBroadcastingUIStatics {
	var p *IAppBroadcastingUIStatics
	hs := NewHStr("Windows.Media.AppBroadcasting.AppBroadcastingUI")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAppBroadcastingUIStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
