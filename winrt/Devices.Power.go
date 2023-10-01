package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// interfaces

// BC894FC6-0072-47C8-8B5D-614AAA7A437E
var IID_IBattery = syscall.GUID{0xBC894FC6, 0x0072, 0x47C8,
	[8]byte{0x8B, 0x5D, 0x61, 0x4A, 0xAA, 0x7A, 0x43, 0x7E}}

type IBatteryInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	GetReport() *IBatteryReport
	Add_ReportUpdated(handler TypedEventHandler[*IBattery, interface{}]) EventRegistrationToken
	Remove_ReportUpdated(token EventRegistrationToken)
}

type IBatteryVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId         uintptr
	GetReport            uintptr
	Add_ReportUpdated    uintptr
	Remove_ReportUpdated uintptr
}

type IBattery struct {
	win32.IInspectable
}

func (this *IBattery) Vtbl() *IBatteryVtbl {
	return (*IBatteryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBattery) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBattery) GetReport() *IBatteryReport {
	var _result *IBatteryReport
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetReport, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBattery) Add_ReportUpdated(handler TypedEventHandler[*IBattery, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ReportUpdated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBattery) Remove_ReportUpdated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ReportUpdated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// C9858C3A-4E13-420A-A8D0-24F18F395401
var IID_IBatteryReport = syscall.GUID{0xC9858C3A, 0x4E13, 0x420A,
	[8]byte{0xA8, 0xD0, 0x24, 0xF1, 0x8F, 0x39, 0x54, 0x01}}

type IBatteryReportInterface interface {
	win32.IInspectableInterface
	Get_ChargeRateInMilliwatts() *IReference[int32]
	Get_DesignCapacityInMilliwattHours() *IReference[int32]
	Get_FullChargeCapacityInMilliwattHours() *IReference[int32]
	Get_RemainingCapacityInMilliwattHours() *IReference[int32]
	Get_Status() unsafe.Pointer
}

type IBatteryReportVtbl struct {
	win32.IInspectableVtbl
	Get_ChargeRateInMilliwatts             uintptr
	Get_DesignCapacityInMilliwattHours     uintptr
	Get_FullChargeCapacityInMilliwattHours uintptr
	Get_RemainingCapacityInMilliwattHours  uintptr
	Get_Status                             uintptr
}

type IBatteryReport struct {
	win32.IInspectable
}

func (this *IBatteryReport) Vtbl() *IBatteryReportVtbl {
	return (*IBatteryReportVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBatteryReport) Get_ChargeRateInMilliwatts() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ChargeRateInMilliwatts, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBatteryReport) Get_DesignCapacityInMilliwattHours() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesignCapacityInMilliwattHours, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBatteryReport) Get_FullChargeCapacityInMilliwattHours() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FullChargeCapacityInMilliwattHours, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBatteryReport) Get_RemainingCapacityInMilliwattHours() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RemainingCapacityInMilliwattHours, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBatteryReport) Get_Status() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 79CD72B6-9E5E-4452-BEA6-DFCD541E597F
var IID_IBatteryStatics = syscall.GUID{0x79CD72B6, 0x9E5E, 0x4452,
	[8]byte{0xBE, 0xA6, 0xDF, 0xCD, 0x54, 0x1E, 0x59, 0x7F}}

type IBatteryStaticsInterface interface {
	win32.IInspectableInterface
	Get_AggregateBattery() *IBattery
	FromIdAsync(deviceId string) *IAsyncOperation[*IBattery]
	GetDeviceSelector() string
}

type IBatteryStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_AggregateBattery uintptr
	FromIdAsync          uintptr
	GetDeviceSelector    uintptr
}

type IBatteryStatics struct {
	win32.IInspectable
}

func (this *IBatteryStatics) Vtbl() *IBatteryStaticsVtbl {
	return (*IBatteryStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBatteryStatics) Get_AggregateBattery() *IBattery {
	var _result *IBattery
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AggregateBattery, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBatteryStatics) FromIdAsync(deviceId string) *IAsyncOperation[*IBattery] {
	var _result *IAsyncOperation[*IBattery]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBatteryStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// classes

type Battery struct {
	RtClass
	*IBattery
}

func NewIBatteryStatics() *IBatteryStatics {
	var p *IBatteryStatics
	hs := NewHStr("Windows.Devices.Power.Battery")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IBatteryStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type BatteryReport struct {
	RtClass
	*IBatteryReport
}
