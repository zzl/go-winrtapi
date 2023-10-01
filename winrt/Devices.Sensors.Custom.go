package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// interfaces

// A136F9AD-4034-4B4D-99DD-531AAC649C09
var IID_ICustomSensor = syscall.GUID{0xA136F9AD, 0x4034, 0x4B4D,
	[8]byte{0x99, 0xDD, 0x53, 0x1A, 0xAC, 0x64, 0x9C, 0x09}}

type ICustomSensorInterface interface {
	win32.IInspectableInterface
	GetCurrentReading() *ICustomSensorReading
	Get_MinimumReportInterval() uint32
	Put_ReportInterval(value uint32)
	Get_ReportInterval() uint32
	Get_DeviceId() string
	Add_ReadingChanged(handler TypedEventHandler[*ICustomSensor, *ICustomSensorReadingChangedEventArgs]) EventRegistrationToken
	Remove_ReadingChanged(token EventRegistrationToken)
}

type ICustomSensorVtbl struct {
	win32.IInspectableVtbl
	GetCurrentReading         uintptr
	Get_MinimumReportInterval uintptr
	Put_ReportInterval        uintptr
	Get_ReportInterval        uintptr
	Get_DeviceId              uintptr
	Add_ReadingChanged        uintptr
	Remove_ReadingChanged     uintptr
}

type ICustomSensor struct {
	win32.IInspectable
}

func (this *ICustomSensor) Vtbl() *ICustomSensorVtbl {
	return (*ICustomSensorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICustomSensor) GetCurrentReading() *ICustomSensorReading {
	var _result *ICustomSensorReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentReading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICustomSensor) Get_MinimumReportInterval() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinimumReportInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICustomSensor) Put_ReportInterval(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReportInterval, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICustomSensor) Get_ReportInterval() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICustomSensor) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICustomSensor) Add_ReadingChanged(handler TypedEventHandler[*ICustomSensor, *ICustomSensorReadingChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ReadingChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICustomSensor) Remove_ReadingChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ReadingChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 20DB3111-EC58-4D9F-BFBD-E77825088510
var IID_ICustomSensor2 = syscall.GUID{0x20DB3111, 0xEC58, 0x4D9F,
	[8]byte{0xBF, 0xBD, 0xE7, 0x78, 0x25, 0x08, 0x85, 0x10}}

type ICustomSensor2Interface interface {
	win32.IInspectableInterface
	Put_ReportLatency(value uint32)
	Get_ReportLatency() uint32
	Get_MaxBatchSize() uint32
}

type ICustomSensor2Vtbl struct {
	win32.IInspectableVtbl
	Put_ReportLatency uintptr
	Get_ReportLatency uintptr
	Get_MaxBatchSize  uintptr
}

type ICustomSensor2 struct {
	win32.IInspectable
}

func (this *ICustomSensor2) Vtbl() *ICustomSensor2Vtbl {
	return (*ICustomSensor2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICustomSensor2) Put_ReportLatency(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReportLatency, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICustomSensor2) Get_ReportLatency() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportLatency, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICustomSensor2) Get_MaxBatchSize() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxBatchSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 64004F4D-446A-4366-A87A-5F963268EC53
var IID_ICustomSensorReading = syscall.GUID{0x64004F4D, 0x446A, 0x4366,
	[8]byte{0xA8, 0x7A, 0x5F, 0x96, 0x32, 0x68, 0xEC, 0x53}}

type ICustomSensorReadingInterface interface {
	win32.IInspectableInterface
	Get_Timestamp() DateTime
	Get_Properties() *IMapView[string, interface{}]
}

type ICustomSensorReadingVtbl struct {
	win32.IInspectableVtbl
	Get_Timestamp  uintptr
	Get_Properties uintptr
}

type ICustomSensorReading struct {
	win32.IInspectable
}

func (this *ICustomSensorReading) Vtbl() *ICustomSensorReadingVtbl {
	return (*ICustomSensorReadingVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICustomSensorReading) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICustomSensorReading) Get_Properties() *IMapView[string, interface{}] {
	var _result *IMapView[string, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 223C98EA-BF73-4992-9A48-D3C897594CCB
var IID_ICustomSensorReading2 = syscall.GUID{0x223C98EA, 0xBF73, 0x4992,
	[8]byte{0x9A, 0x48, 0xD3, 0xC8, 0x97, 0x59, 0x4C, 0xCB}}

type ICustomSensorReading2Interface interface {
	win32.IInspectableInterface
	Get_PerformanceCount() *IReference[TimeSpan]
}

type ICustomSensorReading2Vtbl struct {
	win32.IInspectableVtbl
	Get_PerformanceCount uintptr
}

type ICustomSensorReading2 struct {
	win32.IInspectable
}

func (this *ICustomSensorReading2) Vtbl() *ICustomSensorReading2Vtbl {
	return (*ICustomSensorReading2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICustomSensorReading2) Get_PerformanceCount() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PerformanceCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6B202023-CFFD-4CC1-8FF0-E21823D76FCC
var IID_ICustomSensorReadingChangedEventArgs = syscall.GUID{0x6B202023, 0xCFFD, 0x4CC1,
	[8]byte{0x8F, 0xF0, 0xE2, 0x18, 0x23, 0xD7, 0x6F, 0xCC}}

type ICustomSensorReadingChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Reading() *ICustomSensorReading
}

type ICustomSensorReadingChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Reading uintptr
}

type ICustomSensorReadingChangedEventArgs struct {
	win32.IInspectable
}

func (this *ICustomSensorReadingChangedEventArgs) Vtbl() *ICustomSensorReadingChangedEventArgsVtbl {
	return (*ICustomSensorReadingChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICustomSensorReadingChangedEventArgs) Get_Reading() *ICustomSensorReading {
	var _result *ICustomSensorReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Reading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 992052CF-F422-4C7D-836B-E7DC74A7124B
var IID_ICustomSensorStatics = syscall.GUID{0x992052CF, 0xF422, 0x4C7D,
	[8]byte{0x83, 0x6B, 0xE7, 0xDC, 0x74, 0xA7, 0x12, 0x4B}}

type ICustomSensorStaticsInterface interface {
	win32.IInspectableInterface
	GetDeviceSelector(interfaceId syscall.GUID) string
	FromIdAsync(sensorId string) *IAsyncOperation[*ICustomSensor]
}

type ICustomSensorStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelector uintptr
	FromIdAsync       uintptr
}

type ICustomSensorStatics struct {
	win32.IInspectable
}

func (this *ICustomSensorStatics) Vtbl() *ICustomSensorStaticsVtbl {
	return (*ICustomSensorStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICustomSensorStatics) GetDeviceSelector(interfaceId syscall.GUID) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&interfaceId)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICustomSensorStatics) FromIdAsync(sensorId string) *IAsyncOperation[*ICustomSensor] {
	var _result *IAsyncOperation[*ICustomSensor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(sensorId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type CustomSensor struct {
	RtClass
	*ICustomSensor
}

func NewICustomSensorStatics() *ICustomSensorStatics {
	var p *ICustomSensorStatics
	hs := NewHStr("Windows.Devices.Sensors.Custom.CustomSensor")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICustomSensorStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type CustomSensorReading struct {
	RtClass
	*ICustomSensorReading
}

type CustomSensorReadingChangedEventArgs struct {
	RtClass
	*ICustomSensorReadingChangedEventArgs
}
