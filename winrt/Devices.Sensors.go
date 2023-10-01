package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type AccelerometerReadingType int32

const (
	AccelerometerReadingType_Standard AccelerometerReadingType = 0
	AccelerometerReadingType_Linear   AccelerometerReadingType = 1
	AccelerometerReadingType_Gravity  AccelerometerReadingType = 2
)

// enum
type ActivitySensorReadingConfidence int32

const (
	ActivitySensorReadingConfidence_High ActivitySensorReadingConfidence = 0
	ActivitySensorReadingConfidence_Low  ActivitySensorReadingConfidence = 1
)

// enum
type ActivityType int32

const (
	ActivityType_Unknown    ActivityType = 0
	ActivityType_Idle       ActivityType = 1
	ActivityType_Stationary ActivityType = 2
	ActivityType_Fidgeting  ActivityType = 3
	ActivityType_Walking    ActivityType = 4
	ActivityType_Running    ActivityType = 5
	ActivityType_InVehicle  ActivityType = 6
	ActivityType_Biking     ActivityType = 7
)

// enum
type MagnetometerAccuracy int32

const (
	MagnetometerAccuracy_Unknown     MagnetometerAccuracy = 0
	MagnetometerAccuracy_Unreliable  MagnetometerAccuracy = 1
	MagnetometerAccuracy_Approximate MagnetometerAccuracy = 2
	MagnetometerAccuracy_High        MagnetometerAccuracy = 3
)

// enum
type PedometerStepKind int32

const (
	PedometerStepKind_Unknown PedometerStepKind = 0
	PedometerStepKind_Walking PedometerStepKind = 1
	PedometerStepKind_Running PedometerStepKind = 2
)

// enum
type SensorOptimizationGoal int32

const (
	SensorOptimizationGoal_Precision       SensorOptimizationGoal = 0
	SensorOptimizationGoal_PowerEfficiency SensorOptimizationGoal = 1
)

// enum
type SensorReadingType int32

const (
	SensorReadingType_Absolute SensorReadingType = 0
	SensorReadingType_Relative SensorReadingType = 1
)

// enum
type SensorType int32

const (
	SensorType_Accelerometer             SensorType = 0
	SensorType_ActivitySensor            SensorType = 1
	SensorType_Barometer                 SensorType = 2
	SensorType_Compass                   SensorType = 3
	SensorType_CustomSensor              SensorType = 4
	SensorType_Gyroscope                 SensorType = 5
	SensorType_ProximitySensor           SensorType = 6
	SensorType_Inclinometer              SensorType = 7
	SensorType_LightSensor               SensorType = 8
	SensorType_OrientationSensor         SensorType = 9
	SensorType_Pedometer                 SensorType = 10
	SensorType_RelativeInclinometer      SensorType = 11
	SensorType_RelativeOrientationSensor SensorType = 12
	SensorType_SimpleOrientationSensor   SensorType = 13
)

// enum
type SimpleOrientation int32

const (
	SimpleOrientation_NotRotated                        SimpleOrientation = 0
	SimpleOrientation_Rotated90DegreesCounterclockwise  SimpleOrientation = 1
	SimpleOrientation_Rotated180DegreesCounterclockwise SimpleOrientation = 2
	SimpleOrientation_Rotated270DegreesCounterclockwise SimpleOrientation = 3
	SimpleOrientation_Faceup                            SimpleOrientation = 4
	SimpleOrientation_Facedown                          SimpleOrientation = 5
)

// interfaces

// DF184548-2711-4DA7-8098-4B82205D3C7D
var IID_IAccelerometer = syscall.GUID{0xDF184548, 0x2711, 0x4DA7,
	[8]byte{0x80, 0x98, 0x4B, 0x82, 0x20, 0x5D, 0x3C, 0x7D}}

type IAccelerometerInterface interface {
	win32.IInspectableInterface
	GetCurrentReading() *IAccelerometerReading
	Get_MinimumReportInterval() uint32
	Put_ReportInterval(value uint32)
	Get_ReportInterval() uint32
	Add_ReadingChanged(handler TypedEventHandler[*IAccelerometer, *IAccelerometerReadingChangedEventArgs]) EventRegistrationToken
	Remove_ReadingChanged(token EventRegistrationToken)
	Add_Shaken(handler TypedEventHandler[*IAccelerometer, *IAccelerometerShakenEventArgs]) EventRegistrationToken
	Remove_Shaken(token EventRegistrationToken)
}

type IAccelerometerVtbl struct {
	win32.IInspectableVtbl
	GetCurrentReading         uintptr
	Get_MinimumReportInterval uintptr
	Put_ReportInterval        uintptr
	Get_ReportInterval        uintptr
	Add_ReadingChanged        uintptr
	Remove_ReadingChanged     uintptr
	Add_Shaken                uintptr
	Remove_Shaken             uintptr
}

type IAccelerometer struct {
	win32.IInspectable
}

func (this *IAccelerometer) Vtbl() *IAccelerometerVtbl {
	return (*IAccelerometerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccelerometer) GetCurrentReading() *IAccelerometerReading {
	var _result *IAccelerometerReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentReading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAccelerometer) Get_MinimumReportInterval() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinimumReportInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAccelerometer) Put_ReportInterval(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReportInterval, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAccelerometer) Get_ReportInterval() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAccelerometer) Add_ReadingChanged(handler TypedEventHandler[*IAccelerometer, *IAccelerometerReadingChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ReadingChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAccelerometer) Remove_ReadingChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ReadingChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAccelerometer) Add_Shaken(handler TypedEventHandler[*IAccelerometer, *IAccelerometerShakenEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Shaken, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAccelerometer) Remove_Shaken(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Shaken, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// E8F092EE-4964-401A-B602-220D7153C60A
var IID_IAccelerometer2 = syscall.GUID{0xE8F092EE, 0x4964, 0x401A,
	[8]byte{0xB6, 0x02, 0x22, 0x0D, 0x71, 0x53, 0xC6, 0x0A}}

type IAccelerometer2Interface interface {
	win32.IInspectableInterface
	Put_ReadingTransform(value unsafe.Pointer)
	Get_ReadingTransform() unsafe.Pointer
}

type IAccelerometer2Vtbl struct {
	win32.IInspectableVtbl
	Put_ReadingTransform uintptr
	Get_ReadingTransform uintptr
}

type IAccelerometer2 struct {
	win32.IInspectable
}

func (this *IAccelerometer2) Vtbl() *IAccelerometer2Vtbl {
	return (*IAccelerometer2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccelerometer2) Put_ReadingTransform(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReadingTransform, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAccelerometer2) Get_ReadingTransform() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReadingTransform, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 87E0022A-ED80-49EB-BF8A-A4EA31E5CD84
var IID_IAccelerometer3 = syscall.GUID{0x87E0022A, 0xED80, 0x49EB,
	[8]byte{0xBF, 0x8A, 0xA4, 0xEA, 0x31, 0xE5, 0xCD, 0x84}}

type IAccelerometer3Interface interface {
	win32.IInspectableInterface
	Put_ReportLatency(value uint32)
	Get_ReportLatency() uint32
	Get_MaxBatchSize() uint32
}

type IAccelerometer3Vtbl struct {
	win32.IInspectableVtbl
	Put_ReportLatency uintptr
	Get_ReportLatency uintptr
	Get_MaxBatchSize  uintptr
}

type IAccelerometer3 struct {
	win32.IInspectable
}

func (this *IAccelerometer3) Vtbl() *IAccelerometer3Vtbl {
	return (*IAccelerometer3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccelerometer3) Put_ReportLatency(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReportLatency, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAccelerometer3) Get_ReportLatency() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportLatency, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAccelerometer3) Get_MaxBatchSize() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxBatchSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 1D373C4F-42D3-45B2-8144-AB7FB665EB59
var IID_IAccelerometer4 = syscall.GUID{0x1D373C4F, 0x42D3, 0x45B2,
	[8]byte{0x81, 0x44, 0xAB, 0x7F, 0xB6, 0x65, 0xEB, 0x59}}

type IAccelerometer4Interface interface {
	win32.IInspectableInterface
	Get_ReadingType() AccelerometerReadingType
}

type IAccelerometer4Vtbl struct {
	win32.IInspectableVtbl
	Get_ReadingType uintptr
}

type IAccelerometer4 struct {
	win32.IInspectable
}

func (this *IAccelerometer4) Vtbl() *IAccelerometer4Vtbl {
	return (*IAccelerometer4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccelerometer4) Get_ReadingType() AccelerometerReadingType {
	var _result AccelerometerReadingType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReadingType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7E7E7021-DEF4-53A6-AF43-806FD538EDF6
var IID_IAccelerometer5 = syscall.GUID{0x7E7E7021, 0xDEF4, 0x53A6,
	[8]byte{0xAF, 0x43, 0x80, 0x6F, 0xD5, 0x38, 0xED, 0xF6}}

type IAccelerometer5Interface interface {
	win32.IInspectableInterface
	Get_ReportThreshold() *IAccelerometerDataThreshold
}

type IAccelerometer5Vtbl struct {
	win32.IInspectableVtbl
	Get_ReportThreshold uintptr
}

type IAccelerometer5 struct {
	win32.IInspectable
}

func (this *IAccelerometer5) Vtbl() *IAccelerometer5Vtbl {
	return (*IAccelerometer5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccelerometer5) Get_ReportThreshold() *IAccelerometerDataThreshold {
	var _result *IAccelerometerDataThreshold
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportThreshold, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F92C1B68-6320-5577-879E-9942621C3DD9
var IID_IAccelerometerDataThreshold = syscall.GUID{0xF92C1B68, 0x6320, 0x5577,
	[8]byte{0x87, 0x9E, 0x99, 0x42, 0x62, 0x1C, 0x3D, 0xD9}}

type IAccelerometerDataThresholdInterface interface {
	win32.IInspectableInterface
	Get_XAxisInGForce() float64
	Put_XAxisInGForce(value float64)
	Get_YAxisInGForce() float64
	Put_YAxisInGForce(value float64)
	Get_ZAxisInGForce() float64
	Put_ZAxisInGForce(value float64)
}

type IAccelerometerDataThresholdVtbl struct {
	win32.IInspectableVtbl
	Get_XAxisInGForce uintptr
	Put_XAxisInGForce uintptr
	Get_YAxisInGForce uintptr
	Put_YAxisInGForce uintptr
	Get_ZAxisInGForce uintptr
	Put_ZAxisInGForce uintptr
}

type IAccelerometerDataThreshold struct {
	win32.IInspectable
}

func (this *IAccelerometerDataThreshold) Vtbl() *IAccelerometerDataThresholdVtbl {
	return (*IAccelerometerDataThresholdVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccelerometerDataThreshold) Get_XAxisInGForce() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_XAxisInGForce, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAccelerometerDataThreshold) Put_XAxisInGForce(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_XAxisInGForce, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAccelerometerDataThreshold) Get_YAxisInGForce() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_YAxisInGForce, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAccelerometerDataThreshold) Put_YAxisInGForce(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_YAxisInGForce, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAccelerometerDataThreshold) Get_ZAxisInGForce() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ZAxisInGForce, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAccelerometerDataThreshold) Put_ZAxisInGForce(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ZAxisInGForce, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 7EAC64A9-97D5-446D-AB5A-917DF9B96A2C
var IID_IAccelerometerDeviceId = syscall.GUID{0x7EAC64A9, 0x97D5, 0x446D,
	[8]byte{0xAB, 0x5A, 0x91, 0x7D, 0xF9, 0xB9, 0x6A, 0x2C}}

type IAccelerometerDeviceIdInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
}

type IAccelerometerDeviceIdVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId uintptr
}

type IAccelerometerDeviceId struct {
	win32.IInspectable
}

func (this *IAccelerometerDeviceId) Vtbl() *IAccelerometerDeviceIdVtbl {
	return (*IAccelerometerDeviceIdVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccelerometerDeviceId) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// B9FE7ACB-D351-40AF-8BB6-7AA9AE641FB7
var IID_IAccelerometerReading = syscall.GUID{0xB9FE7ACB, 0xD351, 0x40AF,
	[8]byte{0x8B, 0xB6, 0x7A, 0xA9, 0xAE, 0x64, 0x1F, 0xB7}}

type IAccelerometerReadingInterface interface {
	win32.IInspectableInterface
	Get_Timestamp() DateTime
	Get_AccelerationX() float64
	Get_AccelerationY() float64
	Get_AccelerationZ() float64
}

type IAccelerometerReadingVtbl struct {
	win32.IInspectableVtbl
	Get_Timestamp     uintptr
	Get_AccelerationX uintptr
	Get_AccelerationY uintptr
	Get_AccelerationZ uintptr
}

type IAccelerometerReading struct {
	win32.IInspectable
}

func (this *IAccelerometerReading) Vtbl() *IAccelerometerReadingVtbl {
	return (*IAccelerometerReadingVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccelerometerReading) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAccelerometerReading) Get_AccelerationX() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AccelerationX, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAccelerometerReading) Get_AccelerationY() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AccelerationY, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAccelerometerReading) Get_AccelerationZ() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AccelerationZ, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 0A864AA2-15AE-4A40-BE55-DB58D7DE7389
var IID_IAccelerometerReading2 = syscall.GUID{0x0A864AA2, 0x15AE, 0x4A40,
	[8]byte{0xBE, 0x55, 0xDB, 0x58, 0xD7, 0xDE, 0x73, 0x89}}

type IAccelerometerReading2Interface interface {
	win32.IInspectableInterface
	Get_PerformanceCount() *IReference[TimeSpan]
	Get_Properties() *IMapView[string, interface{}]
}

type IAccelerometerReading2Vtbl struct {
	win32.IInspectableVtbl
	Get_PerformanceCount uintptr
	Get_Properties       uintptr
}

type IAccelerometerReading2 struct {
	win32.IInspectable
}

func (this *IAccelerometerReading2) Vtbl() *IAccelerometerReading2Vtbl {
	return (*IAccelerometerReading2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccelerometerReading2) Get_PerformanceCount() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PerformanceCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAccelerometerReading2) Get_Properties() *IMapView[string, interface{}] {
	var _result *IMapView[string, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0095C65B-B6AC-475A-9F44-8B32D35A3F25
var IID_IAccelerometerReadingChangedEventArgs = syscall.GUID{0x0095C65B, 0xB6AC, 0x475A,
	[8]byte{0x9F, 0x44, 0x8B, 0x32, 0xD3, 0x5A, 0x3F, 0x25}}

type IAccelerometerReadingChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Reading() *IAccelerometerReading
}

type IAccelerometerReadingChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Reading uintptr
}

type IAccelerometerReadingChangedEventArgs struct {
	win32.IInspectable
}

func (this *IAccelerometerReadingChangedEventArgs) Vtbl() *IAccelerometerReadingChangedEventArgsVtbl {
	return (*IAccelerometerReadingChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccelerometerReadingChangedEventArgs) Get_Reading() *IAccelerometerReading {
	var _result *IAccelerometerReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Reading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 95FF01D1-4A28-4F35-98E8-8178AAE4084A
var IID_IAccelerometerShakenEventArgs = syscall.GUID{0x95FF01D1, 0x4A28, 0x4F35,
	[8]byte{0x98, 0xE8, 0x81, 0x78, 0xAA, 0xE4, 0x08, 0x4A}}

type IAccelerometerShakenEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Timestamp() DateTime
}

type IAccelerometerShakenEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Timestamp uintptr
}

type IAccelerometerShakenEventArgs struct {
	win32.IInspectable
}

func (this *IAccelerometerShakenEventArgs) Vtbl() *IAccelerometerShakenEventArgsVtbl {
	return (*IAccelerometerShakenEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccelerometerShakenEventArgs) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A5E28B74-5A87-4A2D-BECC-0F906EA061DD
var IID_IAccelerometerStatics = syscall.GUID{0xA5E28B74, 0x5A87, 0x4A2D,
	[8]byte{0xBE, 0xCC, 0x0F, 0x90, 0x6E, 0xA0, 0x61, 0xDD}}

type IAccelerometerStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *IAccelerometer
}

type IAccelerometerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
}

type IAccelerometerStatics struct {
	win32.IInspectable
}

func (this *IAccelerometerStatics) Vtbl() *IAccelerometerStaticsVtbl {
	return (*IAccelerometerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccelerometerStatics) GetDefault() *IAccelerometer {
	var _result *IAccelerometer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C4C4842F-D86B-4685-B2D7-3396F798D57B
var IID_IAccelerometerStatics2 = syscall.GUID{0xC4C4842F, 0xD86B, 0x4685,
	[8]byte{0xB2, 0xD7, 0x33, 0x96, 0xF7, 0x98, 0xD5, 0x7B}}

type IAccelerometerStatics2Interface interface {
	win32.IInspectableInterface
	GetDefaultWithAccelerometerReadingType(readingType AccelerometerReadingType) *IAccelerometer
}

type IAccelerometerStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetDefaultWithAccelerometerReadingType uintptr
}

type IAccelerometerStatics2 struct {
	win32.IInspectable
}

func (this *IAccelerometerStatics2) Vtbl() *IAccelerometerStatics2Vtbl {
	return (*IAccelerometerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccelerometerStatics2) GetDefaultWithAccelerometerReadingType(readingType AccelerometerReadingType) *IAccelerometer {
	var _result *IAccelerometer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefaultWithAccelerometerReadingType, uintptr(unsafe.Pointer(this)), uintptr(readingType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9DE218CF-455D-4CF3-8200-70E1410340F8
var IID_IAccelerometerStatics3 = syscall.GUID{0x9DE218CF, 0x455D, 0x4CF3,
	[8]byte{0x82, 0x00, 0x70, 0xE1, 0x41, 0x03, 0x40, 0xF8}}

type IAccelerometerStatics3Interface interface {
	win32.IInspectableInterface
	FromIdAsync(deviceId string) *IAsyncOperation[*IAccelerometer]
	GetDeviceSelector(readingType AccelerometerReadingType) string
}

type IAccelerometerStatics3Vtbl struct {
	win32.IInspectableVtbl
	FromIdAsync       uintptr
	GetDeviceSelector uintptr
}

type IAccelerometerStatics3 struct {
	win32.IInspectable
}

func (this *IAccelerometerStatics3) Vtbl() *IAccelerometerStatics3Vtbl {
	return (*IAccelerometerStatics3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccelerometerStatics3) FromIdAsync(deviceId string) *IAsyncOperation[*IAccelerometer] {
	var _result *IAsyncOperation[*IAccelerometer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAccelerometerStatics3) GetDeviceSelector(readingType AccelerometerReadingType) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(readingType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// CD7A630C-FB5F-48EB-B09B-A2708D1C61EF
var IID_IActivitySensor = syscall.GUID{0xCD7A630C, 0xFB5F, 0x48EB,
	[8]byte{0xB0, 0x9B, 0xA2, 0x70, 0x8D, 0x1C, 0x61, 0xEF}}

type IActivitySensorInterface interface {
	win32.IInspectableInterface
	GetCurrentReadingAsync() *IAsyncOperation[*IActivitySensorReading]
	Get_SubscribedActivities() *IVector[ActivityType]
	Get_PowerInMilliwatts() float64
	Get_DeviceId() string
	Get_SupportedActivities() *IVectorView[ActivityType]
	Get_MinimumReportInterval() uint32
	Add_ReadingChanged(handler TypedEventHandler[*IActivitySensor, *IActivitySensorReadingChangedEventArgs]) EventRegistrationToken
	Remove_ReadingChanged(token EventRegistrationToken)
}

type IActivitySensorVtbl struct {
	win32.IInspectableVtbl
	GetCurrentReadingAsync    uintptr
	Get_SubscribedActivities  uintptr
	Get_PowerInMilliwatts     uintptr
	Get_DeviceId              uintptr
	Get_SupportedActivities   uintptr
	Get_MinimumReportInterval uintptr
	Add_ReadingChanged        uintptr
	Remove_ReadingChanged     uintptr
}

type IActivitySensor struct {
	win32.IInspectable
}

func (this *IActivitySensor) Vtbl() *IActivitySensorVtbl {
	return (*IActivitySensorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IActivitySensor) GetCurrentReadingAsync() *IAsyncOperation[*IActivitySensorReading] {
	var _result *IAsyncOperation[*IActivitySensorReading]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentReadingAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IActivitySensor) Get_SubscribedActivities() *IVector[ActivityType] {
	var _result *IVector[ActivityType]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SubscribedActivities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IActivitySensor) Get_PowerInMilliwatts() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PowerInMilliwatts, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IActivitySensor) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IActivitySensor) Get_SupportedActivities() *IVectorView[ActivityType] {
	var _result *IVectorView[ActivityType]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedActivities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IActivitySensor) Get_MinimumReportInterval() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinimumReportInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IActivitySensor) Add_ReadingChanged(handler TypedEventHandler[*IActivitySensor, *IActivitySensorReadingChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ReadingChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IActivitySensor) Remove_ReadingChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ReadingChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 85125A96-1472-40A2-B2AE-E1EF29226C78
var IID_IActivitySensorReading = syscall.GUID{0x85125A96, 0x1472, 0x40A2,
	[8]byte{0xB2, 0xAE, 0xE1, 0xEF, 0x29, 0x22, 0x6C, 0x78}}

type IActivitySensorReadingInterface interface {
	win32.IInspectableInterface
	Get_Timestamp() DateTime
	Get_Activity() ActivityType
	Get_Confidence() ActivitySensorReadingConfidence
}

type IActivitySensorReadingVtbl struct {
	win32.IInspectableVtbl
	Get_Timestamp  uintptr
	Get_Activity   uintptr
	Get_Confidence uintptr
}

type IActivitySensorReading struct {
	win32.IInspectable
}

func (this *IActivitySensorReading) Vtbl() *IActivitySensorReadingVtbl {
	return (*IActivitySensorReadingVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IActivitySensorReading) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IActivitySensorReading) Get_Activity() ActivityType {
	var _result ActivityType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Activity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IActivitySensorReading) Get_Confidence() ActivitySensorReadingConfidence {
	var _result ActivitySensorReadingConfidence
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Confidence, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4F3C2915-D93B-47BD-960A-F20FB2F322B9
var IID_IActivitySensorReadingChangeReport = syscall.GUID{0x4F3C2915, 0xD93B, 0x47BD,
	[8]byte{0x96, 0x0A, 0xF2, 0x0F, 0xB2, 0xF3, 0x22, 0xB9}}

type IActivitySensorReadingChangeReportInterface interface {
	win32.IInspectableInterface
	Get_Reading() *IActivitySensorReading
}

type IActivitySensorReadingChangeReportVtbl struct {
	win32.IInspectableVtbl
	Get_Reading uintptr
}

type IActivitySensorReadingChangeReport struct {
	win32.IInspectable
}

func (this *IActivitySensorReadingChangeReport) Vtbl() *IActivitySensorReadingChangeReportVtbl {
	return (*IActivitySensorReadingChangeReportVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IActivitySensorReadingChangeReport) Get_Reading() *IActivitySensorReading {
	var _result *IActivitySensorReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Reading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DE386717-AEB6-4EC7-946A-D9CC19B951EC
var IID_IActivitySensorReadingChangedEventArgs = syscall.GUID{0xDE386717, 0xAEB6, 0x4EC7,
	[8]byte{0x94, 0x6A, 0xD9, 0xCC, 0x19, 0xB9, 0x51, 0xEC}}

type IActivitySensorReadingChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Reading() *IActivitySensorReading
}

type IActivitySensorReadingChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Reading uintptr
}

type IActivitySensorReadingChangedEventArgs struct {
	win32.IInspectable
}

func (this *IActivitySensorReadingChangedEventArgs) Vtbl() *IActivitySensorReadingChangedEventArgsVtbl {
	return (*IActivitySensorReadingChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IActivitySensorReadingChangedEventArgs) Get_Reading() *IActivitySensorReading {
	var _result *IActivitySensorReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Reading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A71E0E9D-EE8B-45D1-B25B-08CC0DF92AB6
var IID_IActivitySensorStatics = syscall.GUID{0xA71E0E9D, 0xEE8B, 0x45D1,
	[8]byte{0xB2, 0x5B, 0x08, 0xCC, 0x0D, 0xF9, 0x2A, 0xB6}}

type IActivitySensorStaticsInterface interface {
	win32.IInspectableInterface
	GetDefaultAsync() *IAsyncOperation[*IActivitySensor]
	GetDeviceSelector() string
	FromIdAsync(deviceId string) *IAsyncOperation[*IActivitySensor]
	GetSystemHistoryAsync(fromTime DateTime) *IAsyncOperation[*IVectorView[*IActivitySensorReading]]
	GetSystemHistoryWithDurationAsync(fromTime DateTime, duration TimeSpan) *IAsyncOperation[*IVectorView[*IActivitySensorReading]]
}

type IActivitySensorStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefaultAsync                   uintptr
	GetDeviceSelector                 uintptr
	FromIdAsync                       uintptr
	GetSystemHistoryAsync             uintptr
	GetSystemHistoryWithDurationAsync uintptr
}

type IActivitySensorStatics struct {
	win32.IInspectable
}

func (this *IActivitySensorStatics) Vtbl() *IActivitySensorStaticsVtbl {
	return (*IActivitySensorStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IActivitySensorStatics) GetDefaultAsync() *IAsyncOperation[*IActivitySensor] {
	var _result *IAsyncOperation[*IActivitySensor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefaultAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IActivitySensorStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IActivitySensorStatics) FromIdAsync(deviceId string) *IAsyncOperation[*IActivitySensor] {
	var _result *IAsyncOperation[*IActivitySensor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IActivitySensorStatics) GetSystemHistoryAsync(fromTime DateTime) *IAsyncOperation[*IVectorView[*IActivitySensorReading]] {
	var _result *IAsyncOperation[*IVectorView[*IActivitySensorReading]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSystemHistoryAsync, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&fromTime)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IActivitySensorStatics) GetSystemHistoryWithDurationAsync(fromTime DateTime, duration TimeSpan) *IAsyncOperation[*IVectorView[*IActivitySensorReading]] {
	var _result *IAsyncOperation[*IVectorView[*IActivitySensorReading]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSystemHistoryWithDurationAsync, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&fromTime)), *(*uintptr)(unsafe.Pointer(&duration)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2C9E6612-B9CA-4677-B263-243297F79D3A
var IID_IActivitySensorTriggerDetails = syscall.GUID{0x2C9E6612, 0xB9CA, 0x4677,
	[8]byte{0xB2, 0x63, 0x24, 0x32, 0x97, 0xF7, 0x9D, 0x3A}}

type IActivitySensorTriggerDetailsInterface interface {
	win32.IInspectableInterface
	ReadReports() *IVectorView[*IActivitySensorReadingChangeReport]
}

type IActivitySensorTriggerDetailsVtbl struct {
	win32.IInspectableVtbl
	ReadReports uintptr
}

type IActivitySensorTriggerDetails struct {
	win32.IInspectable
}

func (this *IActivitySensorTriggerDetails) Vtbl() *IActivitySensorTriggerDetailsVtbl {
	return (*IActivitySensorTriggerDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IActivitySensorTriggerDetails) ReadReports() *IVectorView[*IActivitySensorReadingChangeReport] {
	var _result *IVectorView[*IActivitySensorReadingChangeReport]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadReports, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 72F057FD-8F04-49F1-B4A7-F4E363B701A2
var IID_IAltimeter = syscall.GUID{0x72F057FD, 0x8F04, 0x49F1,
	[8]byte{0xB4, 0xA7, 0xF4, 0xE3, 0x63, 0xB7, 0x01, 0xA2}}

type IAltimeterInterface interface {
	win32.IInspectableInterface
	GetCurrentReading() *IAltimeterReading
	Get_DeviceId() string
	Get_MinimumReportInterval() uint32
	Put_ReportInterval(value uint32)
	Get_ReportInterval() uint32
	Add_ReadingChanged(handler TypedEventHandler[*IAltimeter, *IAltimeterReadingChangedEventArgs]) EventRegistrationToken
	Remove_ReadingChanged(token EventRegistrationToken)
}

type IAltimeterVtbl struct {
	win32.IInspectableVtbl
	GetCurrentReading         uintptr
	Get_DeviceId              uintptr
	Get_MinimumReportInterval uintptr
	Put_ReportInterval        uintptr
	Get_ReportInterval        uintptr
	Add_ReadingChanged        uintptr
	Remove_ReadingChanged     uintptr
}

type IAltimeter struct {
	win32.IInspectable
}

func (this *IAltimeter) Vtbl() *IAltimeterVtbl {
	return (*IAltimeterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAltimeter) GetCurrentReading() *IAltimeterReading {
	var _result *IAltimeterReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentReading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAltimeter) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAltimeter) Get_MinimumReportInterval() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinimumReportInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAltimeter) Put_ReportInterval(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReportInterval, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAltimeter) Get_ReportInterval() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAltimeter) Add_ReadingChanged(handler TypedEventHandler[*IAltimeter, *IAltimeterReadingChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ReadingChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAltimeter) Remove_ReadingChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ReadingChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// C9471BF9-2ADD-48F5-9F08-3D0C7660D938
var IID_IAltimeter2 = syscall.GUID{0xC9471BF9, 0x2ADD, 0x48F5,
	[8]byte{0x9F, 0x08, 0x3D, 0x0C, 0x76, 0x60, 0xD9, 0x38}}

type IAltimeter2Interface interface {
	win32.IInspectableInterface
	Put_ReportLatency(value uint32)
	Get_ReportLatency() uint32
	Get_MaxBatchSize() uint32
}

type IAltimeter2Vtbl struct {
	win32.IInspectableVtbl
	Put_ReportLatency uintptr
	Get_ReportLatency uintptr
	Get_MaxBatchSize  uintptr
}

type IAltimeter2 struct {
	win32.IInspectable
}

func (this *IAltimeter2) Vtbl() *IAltimeter2Vtbl {
	return (*IAltimeter2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAltimeter2) Put_ReportLatency(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReportLatency, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAltimeter2) Get_ReportLatency() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportLatency, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAltimeter2) Get_MaxBatchSize() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxBatchSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// FBE8EF73-7F5E-48C8-AA1A-F1F3BEFC1144
var IID_IAltimeterReading = syscall.GUID{0xFBE8EF73, 0x7F5E, 0x48C8,
	[8]byte{0xAA, 0x1A, 0xF1, 0xF3, 0xBE, 0xFC, 0x11, 0x44}}

type IAltimeterReadingInterface interface {
	win32.IInspectableInterface
	Get_Timestamp() DateTime
	Get_AltitudeChangeInMeters() float64
}

type IAltimeterReadingVtbl struct {
	win32.IInspectableVtbl
	Get_Timestamp              uintptr
	Get_AltitudeChangeInMeters uintptr
}

type IAltimeterReading struct {
	win32.IInspectable
}

func (this *IAltimeterReading) Vtbl() *IAltimeterReadingVtbl {
	return (*IAltimeterReadingVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAltimeterReading) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAltimeterReading) Get_AltitudeChangeInMeters() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AltitudeChangeInMeters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 543A1BD9-6D0B-42B2-BD69-BC8FAE0F782C
var IID_IAltimeterReading2 = syscall.GUID{0x543A1BD9, 0x6D0B, 0x42B2,
	[8]byte{0xBD, 0x69, 0xBC, 0x8F, 0xAE, 0x0F, 0x78, 0x2C}}

type IAltimeterReading2Interface interface {
	win32.IInspectableInterface
	Get_PerformanceCount() *IReference[TimeSpan]
	Get_Properties() *IMapView[string, interface{}]
}

type IAltimeterReading2Vtbl struct {
	win32.IInspectableVtbl
	Get_PerformanceCount uintptr
	Get_Properties       uintptr
}

type IAltimeterReading2 struct {
	win32.IInspectable
}

func (this *IAltimeterReading2) Vtbl() *IAltimeterReading2Vtbl {
	return (*IAltimeterReading2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAltimeterReading2) Get_PerformanceCount() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PerformanceCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAltimeterReading2) Get_Properties() *IMapView[string, interface{}] {
	var _result *IMapView[string, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7069D077-446D-47F7-998C-EBC23B45E4A2
var IID_IAltimeterReadingChangedEventArgs = syscall.GUID{0x7069D077, 0x446D, 0x47F7,
	[8]byte{0x99, 0x8C, 0xEB, 0xC2, 0x3B, 0x45, 0xE4, 0xA2}}

type IAltimeterReadingChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Reading() *IAltimeterReading
}

type IAltimeterReadingChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Reading uintptr
}

type IAltimeterReadingChangedEventArgs struct {
	win32.IInspectable
}

func (this *IAltimeterReadingChangedEventArgs) Vtbl() *IAltimeterReadingChangedEventArgsVtbl {
	return (*IAltimeterReadingChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAltimeterReadingChangedEventArgs) Get_Reading() *IAltimeterReading {
	var _result *IAltimeterReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Reading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9EB4D7C3-E5AC-47CE-8EEF-D3718168C01F
var IID_IAltimeterStatics = syscall.GUID{0x9EB4D7C3, 0xE5AC, 0x47CE,
	[8]byte{0x8E, 0xEF, 0xD3, 0x71, 0x81, 0x68, 0xC0, 0x1F}}

type IAltimeterStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *IAltimeter
}

type IAltimeterStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
}

type IAltimeterStatics struct {
	win32.IInspectable
}

func (this *IAltimeterStatics) Vtbl() *IAltimeterStaticsVtbl {
	return (*IAltimeterStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAltimeterStatics) GetDefault() *IAltimeter {
	var _result *IAltimeter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 934475A8-78BF-452F-B017-F0209CE6DAB4
var IID_IBarometer = syscall.GUID{0x934475A8, 0x78BF, 0x452F,
	[8]byte{0xB0, 0x17, 0xF0, 0x20, 0x9C, 0xE6, 0xDA, 0xB4}}

type IBarometerInterface interface {
	win32.IInspectableInterface
	GetCurrentReading() *IBarometerReading
	Get_DeviceId() string
	Get_MinimumReportInterval() uint32
	Put_ReportInterval(value uint32)
	Get_ReportInterval() uint32
	Add_ReadingChanged(handler TypedEventHandler[*IBarometer, *IBarometerReadingChangedEventArgs]) EventRegistrationToken
	Remove_ReadingChanged(token EventRegistrationToken)
}

type IBarometerVtbl struct {
	win32.IInspectableVtbl
	GetCurrentReading         uintptr
	Get_DeviceId              uintptr
	Get_MinimumReportInterval uintptr
	Put_ReportInterval        uintptr
	Get_ReportInterval        uintptr
	Add_ReadingChanged        uintptr
	Remove_ReadingChanged     uintptr
}

type IBarometer struct {
	win32.IInspectable
}

func (this *IBarometer) Vtbl() *IBarometerVtbl {
	return (*IBarometerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBarometer) GetCurrentReading() *IBarometerReading {
	var _result *IBarometerReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentReading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBarometer) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBarometer) Get_MinimumReportInterval() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinimumReportInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarometer) Put_ReportInterval(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReportInterval, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IBarometer) Get_ReportInterval() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarometer) Add_ReadingChanged(handler TypedEventHandler[*IBarometer, *IBarometerReadingChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ReadingChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarometer) Remove_ReadingChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ReadingChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 32BCC418-3EEB-4D04-9574-7633A8781F9F
var IID_IBarometer2 = syscall.GUID{0x32BCC418, 0x3EEB, 0x4D04,
	[8]byte{0x95, 0x74, 0x76, 0x33, 0xA8, 0x78, 0x1F, 0x9F}}

type IBarometer2Interface interface {
	win32.IInspectableInterface
	Put_ReportLatency(value uint32)
	Get_ReportLatency() uint32
	Get_MaxBatchSize() uint32
}

type IBarometer2Vtbl struct {
	win32.IInspectableVtbl
	Put_ReportLatency uintptr
	Get_ReportLatency uintptr
	Get_MaxBatchSize  uintptr
}

type IBarometer2 struct {
	win32.IInspectable
}

func (this *IBarometer2) Vtbl() *IBarometer2Vtbl {
	return (*IBarometer2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBarometer2) Put_ReportLatency(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReportLatency, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IBarometer2) Get_ReportLatency() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportLatency, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarometer2) Get_MaxBatchSize() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxBatchSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 0E35F0EA-02B5-5A04-B03D-822084863A54
var IID_IBarometer3 = syscall.GUID{0x0E35F0EA, 0x02B5, 0x5A04,
	[8]byte{0xB0, 0x3D, 0x82, 0x20, 0x84, 0x86, 0x3A, 0x54}}

type IBarometer3Interface interface {
	win32.IInspectableInterface
	Get_ReportThreshold() *IBarometerDataThreshold
}

type IBarometer3Vtbl struct {
	win32.IInspectableVtbl
	Get_ReportThreshold uintptr
}

type IBarometer3 struct {
	win32.IInspectable
}

func (this *IBarometer3) Vtbl() *IBarometer3Vtbl {
	return (*IBarometer3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBarometer3) Get_ReportThreshold() *IBarometerDataThreshold {
	var _result *IBarometerDataThreshold
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportThreshold, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 076B952C-CB62-5A90-A0D1-F85E4A936394
var IID_IBarometerDataThreshold = syscall.GUID{0x076B952C, 0xCB62, 0x5A90,
	[8]byte{0xA0, 0xD1, 0xF8, 0x5E, 0x4A, 0x93, 0x63, 0x94}}

type IBarometerDataThresholdInterface interface {
	win32.IInspectableInterface
	Get_Hectopascals() float64
	Put_Hectopascals(value float64)
}

type IBarometerDataThresholdVtbl struct {
	win32.IInspectableVtbl
	Get_Hectopascals uintptr
	Put_Hectopascals uintptr
}

type IBarometerDataThreshold struct {
	win32.IInspectable
}

func (this *IBarometerDataThreshold) Vtbl() *IBarometerDataThresholdVtbl {
	return (*IBarometerDataThresholdVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBarometerDataThreshold) Get_Hectopascals() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Hectopascals, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarometerDataThreshold) Put_Hectopascals(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Hectopascals, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// F5B9D2E6-1DF6-4A1A-A7AD-321D4F5DB247
var IID_IBarometerReading = syscall.GUID{0xF5B9D2E6, 0x1DF6, 0x4A1A,
	[8]byte{0xA7, 0xAD, 0x32, 0x1D, 0x4F, 0x5D, 0xB2, 0x47}}

type IBarometerReadingInterface interface {
	win32.IInspectableInterface
	Get_Timestamp() DateTime
	Get_StationPressureInHectopascals() float64
}

type IBarometerReadingVtbl struct {
	win32.IInspectableVtbl
	Get_Timestamp                     uintptr
	Get_StationPressureInHectopascals uintptr
}

type IBarometerReading struct {
	win32.IInspectable
}

func (this *IBarometerReading) Vtbl() *IBarometerReadingVtbl {
	return (*IBarometerReadingVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBarometerReading) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBarometerReading) Get_StationPressureInHectopascals() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StationPressureInHectopascals, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 85A244EB-90C5-4875-891C-3865B4C357E7
var IID_IBarometerReading2 = syscall.GUID{0x85A244EB, 0x90C5, 0x4875,
	[8]byte{0x89, 0x1C, 0x38, 0x65, 0xB4, 0xC3, 0x57, 0xE7}}

type IBarometerReading2Interface interface {
	win32.IInspectableInterface
	Get_PerformanceCount() *IReference[TimeSpan]
	Get_Properties() *IMapView[string, interface{}]
}

type IBarometerReading2Vtbl struct {
	win32.IInspectableVtbl
	Get_PerformanceCount uintptr
	Get_Properties       uintptr
}

type IBarometerReading2 struct {
	win32.IInspectable
}

func (this *IBarometerReading2) Vtbl() *IBarometerReading2Vtbl {
	return (*IBarometerReading2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBarometerReading2) Get_PerformanceCount() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PerformanceCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBarometerReading2) Get_Properties() *IMapView[string, interface{}] {
	var _result *IMapView[string, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3D84945F-037B-404F-9BBB-6232D69543C3
var IID_IBarometerReadingChangedEventArgs = syscall.GUID{0x3D84945F, 0x037B, 0x404F,
	[8]byte{0x9B, 0xBB, 0x62, 0x32, 0xD6, 0x95, 0x43, 0xC3}}

type IBarometerReadingChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Reading() *IBarometerReading
}

type IBarometerReadingChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Reading uintptr
}

type IBarometerReadingChangedEventArgs struct {
	win32.IInspectable
}

func (this *IBarometerReadingChangedEventArgs) Vtbl() *IBarometerReadingChangedEventArgsVtbl {
	return (*IBarometerReadingChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBarometerReadingChangedEventArgs) Get_Reading() *IBarometerReading {
	var _result *IBarometerReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Reading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 286B270A-02E3-4F86-84FC-FDD892B5940F
var IID_IBarometerStatics = syscall.GUID{0x286B270A, 0x02E3, 0x4F86,
	[8]byte{0x84, 0xFC, 0xFD, 0xD8, 0x92, 0xB5, 0x94, 0x0F}}

type IBarometerStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *IBarometer
}

type IBarometerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
}

type IBarometerStatics struct {
	win32.IInspectable
}

func (this *IBarometerStatics) Vtbl() *IBarometerStaticsVtbl {
	return (*IBarometerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBarometerStatics) GetDefault() *IBarometer {
	var _result *IBarometer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8FC6B1E7-95FF-44AC-878E-D65C8308C34C
var IID_IBarometerStatics2 = syscall.GUID{0x8FC6B1E7, 0x95FF, 0x44AC,
	[8]byte{0x87, 0x8E, 0xD6, 0x5C, 0x83, 0x08, 0xC3, 0x4C}}

type IBarometerStatics2Interface interface {
	win32.IInspectableInterface
	FromIdAsync(deviceId string) *IAsyncOperation[*IBarometer]
	GetDeviceSelector() string
}

type IBarometerStatics2Vtbl struct {
	win32.IInspectableVtbl
	FromIdAsync       uintptr
	GetDeviceSelector uintptr
}

type IBarometerStatics2 struct {
	win32.IInspectable
}

func (this *IBarometerStatics2) Vtbl() *IBarometerStatics2Vtbl {
	return (*IBarometerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBarometerStatics2) FromIdAsync(deviceId string) *IAsyncOperation[*IBarometer] {
	var _result *IAsyncOperation[*IBarometer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBarometerStatics2) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 292FFA94-1B45-403C-BA06-B106DBA69A64
var IID_ICompass = syscall.GUID{0x292FFA94, 0x1B45, 0x403C,
	[8]byte{0xBA, 0x06, 0xB1, 0x06, 0xDB, 0xA6, 0x9A, 0x64}}

type ICompassInterface interface {
	win32.IInspectableInterface
	GetCurrentReading() *ICompassReading
	Get_MinimumReportInterval() uint32
	Put_ReportInterval(value uint32)
	Get_ReportInterval() uint32
	Add_ReadingChanged(handler TypedEventHandler[*ICompass, *ICompassReadingChangedEventArgs]) EventRegistrationToken
	Remove_ReadingChanged(token EventRegistrationToken)
}

type ICompassVtbl struct {
	win32.IInspectableVtbl
	GetCurrentReading         uintptr
	Get_MinimumReportInterval uintptr
	Put_ReportInterval        uintptr
	Get_ReportInterval        uintptr
	Add_ReadingChanged        uintptr
	Remove_ReadingChanged     uintptr
}

type ICompass struct {
	win32.IInspectable
}

func (this *ICompass) Vtbl() *ICompassVtbl {
	return (*ICompassVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompass) GetCurrentReading() *ICompassReading {
	var _result *ICompassReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentReading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompass) Get_MinimumReportInterval() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinimumReportInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompass) Put_ReportInterval(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReportInterval, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompass) Get_ReportInterval() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompass) Add_ReadingChanged(handler TypedEventHandler[*ICompass, *ICompassReadingChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ReadingChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompass) Remove_ReadingChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ReadingChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 36F26D09-C7D7-434F-B461-979DDFC2322F
var IID_ICompass2 = syscall.GUID{0x36F26D09, 0xC7D7, 0x434F,
	[8]byte{0xB4, 0x61, 0x97, 0x9D, 0xDF, 0xC2, 0x32, 0x2F}}

type ICompass2Interface interface {
	win32.IInspectableInterface
	Put_ReadingTransform(value unsafe.Pointer)
	Get_ReadingTransform() unsafe.Pointer
}

type ICompass2Vtbl struct {
	win32.IInspectableVtbl
	Put_ReadingTransform uintptr
	Get_ReadingTransform uintptr
}

type ICompass2 struct {
	win32.IInspectable
}

func (this *ICompass2) Vtbl() *ICompass2Vtbl {
	return (*ICompass2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompass2) Put_ReadingTransform(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReadingTransform, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompass2) Get_ReadingTransform() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReadingTransform, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A424801B-C5EA-4D45-A0EC-4B791F041A89
var IID_ICompass3 = syscall.GUID{0xA424801B, 0xC5EA, 0x4D45,
	[8]byte{0xA0, 0xEC, 0x4B, 0x79, 0x1F, 0x04, 0x1A, 0x89}}

type ICompass3Interface interface {
	win32.IInspectableInterface
	Put_ReportLatency(value uint32)
	Get_ReportLatency() uint32
	Get_MaxBatchSize() uint32
}

type ICompass3Vtbl struct {
	win32.IInspectableVtbl
	Put_ReportLatency uintptr
	Get_ReportLatency uintptr
	Get_MaxBatchSize  uintptr
}

type ICompass3 struct {
	win32.IInspectable
}

func (this *ICompass3) Vtbl() *ICompass3Vtbl {
	return (*ICompass3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompass3) Put_ReportLatency(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReportLatency, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompass3) Get_ReportLatency() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportLatency, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompass3) Get_MaxBatchSize() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxBatchSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 291E7F11-EC32-5DCC-BFCB-0BB39EBA5774
var IID_ICompass4 = syscall.GUID{0x291E7F11, 0xEC32, 0x5DCC,
	[8]byte{0xBF, 0xCB, 0x0B, 0xB3, 0x9E, 0xBA, 0x57, 0x74}}

type ICompass4Interface interface {
	win32.IInspectableInterface
	Get_ReportThreshold() *ICompassDataThreshold
}

type ICompass4Vtbl struct {
	win32.IInspectableVtbl
	Get_ReportThreshold uintptr
}

type ICompass4 struct {
	win32.IInspectable
}

func (this *ICompass4) Vtbl() *ICompass4Vtbl {
	return (*ICompass4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompass4) Get_ReportThreshold() *ICompassDataThreshold {
	var _result *ICompassDataThreshold
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportThreshold, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D15B52B3-D39D-5EC8-B2E4-F193E6AB34ED
var IID_ICompassDataThreshold = syscall.GUID{0xD15B52B3, 0xD39D, 0x5EC8,
	[8]byte{0xB2, 0xE4, 0xF1, 0x93, 0xE6, 0xAB, 0x34, 0xED}}

type ICompassDataThresholdInterface interface {
	win32.IInspectableInterface
	Get_Degrees() float64
	Put_Degrees(value float64)
}

type ICompassDataThresholdVtbl struct {
	win32.IInspectableVtbl
	Get_Degrees uintptr
	Put_Degrees uintptr
}

type ICompassDataThreshold struct {
	win32.IInspectable
}

func (this *ICompassDataThreshold) Vtbl() *ICompassDataThresholdVtbl {
	return (*ICompassDataThresholdVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompassDataThreshold) Get_Degrees() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Degrees, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompassDataThreshold) Put_Degrees(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Degrees, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// D181CA29-B085-4B1D-870A-4FF57BA74FD4
var IID_ICompassDeviceId = syscall.GUID{0xD181CA29, 0xB085, 0x4B1D,
	[8]byte{0x87, 0x0A, 0x4F, 0xF5, 0x7B, 0xA7, 0x4F, 0xD4}}

type ICompassDeviceIdInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
}

type ICompassDeviceIdVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId uintptr
}

type ICompassDeviceId struct {
	win32.IInspectable
}

func (this *ICompassDeviceId) Vtbl() *ICompassDeviceIdVtbl {
	return (*ICompassDeviceIdVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompassDeviceId) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 82911128-513D-4DC9-B781-5EEDFBF02D0C
var IID_ICompassReading = syscall.GUID{0x82911128, 0x513D, 0x4DC9,
	[8]byte{0xB7, 0x81, 0x5E, 0xED, 0xFB, 0xF0, 0x2D, 0x0C}}

type ICompassReadingInterface interface {
	win32.IInspectableInterface
	Get_Timestamp() DateTime
	Get_HeadingMagneticNorth() float64
	Get_HeadingTrueNorth() *IReference[float64]
}

type ICompassReadingVtbl struct {
	win32.IInspectableVtbl
	Get_Timestamp            uintptr
	Get_HeadingMagneticNorth uintptr
	Get_HeadingTrueNorth     uintptr
}

type ICompassReading struct {
	win32.IInspectable
}

func (this *ICompassReading) Vtbl() *ICompassReadingVtbl {
	return (*ICompassReadingVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompassReading) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompassReading) Get_HeadingMagneticNorth() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HeadingMagneticNorth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompassReading) Get_HeadingTrueNorth() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HeadingTrueNorth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B13A661E-51BB-4A12-BEDD-AD47FF87D2E8
var IID_ICompassReading2 = syscall.GUID{0xB13A661E, 0x51BB, 0x4A12,
	[8]byte{0xBE, 0xDD, 0xAD, 0x47, 0xFF, 0x87, 0xD2, 0xE8}}

type ICompassReading2Interface interface {
	win32.IInspectableInterface
	Get_PerformanceCount() *IReference[TimeSpan]
	Get_Properties() *IMapView[string, interface{}]
}

type ICompassReading2Vtbl struct {
	win32.IInspectableVtbl
	Get_PerformanceCount uintptr
	Get_Properties       uintptr
}

type ICompassReading2 struct {
	win32.IInspectable
}

func (this *ICompassReading2) Vtbl() *ICompassReading2Vtbl {
	return (*ICompassReading2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompassReading2) Get_PerformanceCount() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PerformanceCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompassReading2) Get_Properties() *IMapView[string, interface{}] {
	var _result *IMapView[string, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8F1549B0-E8BC-4C7E-B009-4E41DF137072
var IID_ICompassReadingChangedEventArgs = syscall.GUID{0x8F1549B0, 0xE8BC, 0x4C7E,
	[8]byte{0xB0, 0x09, 0x4E, 0x41, 0xDF, 0x13, 0x70, 0x72}}

type ICompassReadingChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Reading() *ICompassReading
}

type ICompassReadingChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Reading uintptr
}

type ICompassReadingChangedEventArgs struct {
	win32.IInspectable
}

func (this *ICompassReadingChangedEventArgs) Vtbl() *ICompassReadingChangedEventArgsVtbl {
	return (*ICompassReadingChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompassReadingChangedEventArgs) Get_Reading() *ICompassReading {
	var _result *ICompassReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Reading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E761354E-8911-40F7-9E16-6ECC7DAEC5DE
var IID_ICompassReadingHeadingAccuracy = syscall.GUID{0xE761354E, 0x8911, 0x40F7,
	[8]byte{0x9E, 0x16, 0x6E, 0xCC, 0x7D, 0xAE, 0xC5, 0xDE}}

type ICompassReadingHeadingAccuracyInterface interface {
	win32.IInspectableInterface
	Get_HeadingAccuracy() MagnetometerAccuracy
}

type ICompassReadingHeadingAccuracyVtbl struct {
	win32.IInspectableVtbl
	Get_HeadingAccuracy uintptr
}

type ICompassReadingHeadingAccuracy struct {
	win32.IInspectable
}

func (this *ICompassReadingHeadingAccuracy) Vtbl() *ICompassReadingHeadingAccuracyVtbl {
	return (*ICompassReadingHeadingAccuracyVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompassReadingHeadingAccuracy) Get_HeadingAccuracy() MagnetometerAccuracy {
	var _result MagnetometerAccuracy
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HeadingAccuracy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 9ABC97DF-56EC-4C25-B54D-40A68BB5B269
var IID_ICompassStatics = syscall.GUID{0x9ABC97DF, 0x56EC, 0x4C25,
	[8]byte{0xB5, 0x4D, 0x40, 0xA6, 0x8B, 0xB5, 0xB2, 0x69}}

type ICompassStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *ICompass
}

type ICompassStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
}

type ICompassStatics struct {
	win32.IInspectable
}

func (this *ICompassStatics) Vtbl() *ICompassStaticsVtbl {
	return (*ICompassStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompassStatics) GetDefault() *ICompass {
	var _result *ICompass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0ACE0EAD-3BAA-4990-9CE4-BE0913754ED2
var IID_ICompassStatics2 = syscall.GUID{0x0ACE0EAD, 0x3BAA, 0x4990,
	[8]byte{0x9C, 0xE4, 0xBE, 0x09, 0x13, 0x75, 0x4E, 0xD2}}

type ICompassStatics2Interface interface {
	win32.IInspectableInterface
	GetDeviceSelector() string
	FromIdAsync(deviceId string) *IAsyncOperation[*ICompass]
}

type ICompassStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelector uintptr
	FromIdAsync       uintptr
}

type ICompassStatics2 struct {
	win32.IInspectable
}

func (this *ICompassStatics2) Vtbl() *ICompassStatics2Vtbl {
	return (*ICompassStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompassStatics2) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICompassStatics2) FromIdAsync(deviceId string) *IAsyncOperation[*ICompass] {
	var _result *IAsyncOperation[*ICompass]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FDB9A9C4-84B1-4CA2-9763-9B589506C70C
var IID_IGyrometer = syscall.GUID{0xFDB9A9C4, 0x84B1, 0x4CA2,
	[8]byte{0x97, 0x63, 0x9B, 0x58, 0x95, 0x06, 0xC7, 0x0C}}

type IGyrometerInterface interface {
	win32.IInspectableInterface
	GetCurrentReading() *IGyrometerReading
	Get_MinimumReportInterval() uint32
	Put_ReportInterval(value uint32)
	Get_ReportInterval() uint32
	Add_ReadingChanged(handler TypedEventHandler[*IGyrometer, *IGyrometerReadingChangedEventArgs]) EventRegistrationToken
	Remove_ReadingChanged(token EventRegistrationToken)
}

type IGyrometerVtbl struct {
	win32.IInspectableVtbl
	GetCurrentReading         uintptr
	Get_MinimumReportInterval uintptr
	Put_ReportInterval        uintptr
	Get_ReportInterval        uintptr
	Add_ReadingChanged        uintptr
	Remove_ReadingChanged     uintptr
}

type IGyrometer struct {
	win32.IInspectable
}

func (this *IGyrometer) Vtbl() *IGyrometerVtbl {
	return (*IGyrometerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGyrometer) GetCurrentReading() *IGyrometerReading {
	var _result *IGyrometerReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentReading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGyrometer) Get_MinimumReportInterval() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinimumReportInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGyrometer) Put_ReportInterval(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReportInterval, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGyrometer) Get_ReportInterval() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGyrometer) Add_ReadingChanged(handler TypedEventHandler[*IGyrometer, *IGyrometerReadingChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ReadingChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGyrometer) Remove_ReadingChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ReadingChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 63DF2443-8CE8-41C3-AC44-8698810B557F
var IID_IGyrometer2 = syscall.GUID{0x63DF2443, 0x8CE8, 0x41C3,
	[8]byte{0xAC, 0x44, 0x86, 0x98, 0x81, 0x0B, 0x55, 0x7F}}

type IGyrometer2Interface interface {
	win32.IInspectableInterface
	Put_ReadingTransform(value unsafe.Pointer)
	Get_ReadingTransform() unsafe.Pointer
}

type IGyrometer2Vtbl struct {
	win32.IInspectableVtbl
	Put_ReadingTransform uintptr
	Get_ReadingTransform uintptr
}

type IGyrometer2 struct {
	win32.IInspectable
}

func (this *IGyrometer2) Vtbl() *IGyrometer2Vtbl {
	return (*IGyrometer2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGyrometer2) Put_ReadingTransform(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReadingTransform, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGyrometer2) Get_ReadingTransform() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReadingTransform, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5D6F88D5-8FBC-4484-914B-528ADFD947B1
var IID_IGyrometer3 = syscall.GUID{0x5D6F88D5, 0x8FBC, 0x4484,
	[8]byte{0x91, 0x4B, 0x52, 0x8A, 0xDF, 0xD9, 0x47, 0xB1}}

type IGyrometer3Interface interface {
	win32.IInspectableInterface
	Put_ReportLatency(value uint32)
	Get_ReportLatency() uint32
	Get_MaxBatchSize() uint32
}

type IGyrometer3Vtbl struct {
	win32.IInspectableVtbl
	Put_ReportLatency uintptr
	Get_ReportLatency uintptr
	Get_MaxBatchSize  uintptr
}

type IGyrometer3 struct {
	win32.IInspectable
}

func (this *IGyrometer3) Vtbl() *IGyrometer3Vtbl {
	return (*IGyrometer3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGyrometer3) Put_ReportLatency(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReportLatency, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGyrometer3) Get_ReportLatency() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportLatency, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGyrometer3) Get_MaxBatchSize() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxBatchSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 0628A60C-4C4B-5096-94E6-C356DF68BEF7
var IID_IGyrometer4 = syscall.GUID{0x0628A60C, 0x4C4B, 0x5096,
	[8]byte{0x94, 0xE6, 0xC3, 0x56, 0xDF, 0x68, 0xBE, 0xF7}}

type IGyrometer4Interface interface {
	win32.IInspectableInterface
	Get_ReportThreshold() *IGyrometerDataThreshold
}

type IGyrometer4Vtbl struct {
	win32.IInspectableVtbl
	Get_ReportThreshold uintptr
}

type IGyrometer4 struct {
	win32.IInspectable
}

func (this *IGyrometer4) Vtbl() *IGyrometer4Vtbl {
	return (*IGyrometer4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGyrometer4) Get_ReportThreshold() *IGyrometerDataThreshold {
	var _result *IGyrometerDataThreshold
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportThreshold, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8648B31E-6E52-5259-BBAD-242A69DC38C8
var IID_IGyrometerDataThreshold = syscall.GUID{0x8648B31E, 0x6E52, 0x5259,
	[8]byte{0xBB, 0xAD, 0x24, 0x2A, 0x69, 0xDC, 0x38, 0xC8}}

type IGyrometerDataThresholdInterface interface {
	win32.IInspectableInterface
	Get_XAxisInDegreesPerSecond() float64
	Put_XAxisInDegreesPerSecond(value float64)
	Get_YAxisInDegreesPerSecond() float64
	Put_YAxisInDegreesPerSecond(value float64)
	Get_ZAxisInDegreesPerSecond() float64
	Put_ZAxisInDegreesPerSecond(value float64)
}

type IGyrometerDataThresholdVtbl struct {
	win32.IInspectableVtbl
	Get_XAxisInDegreesPerSecond uintptr
	Put_XAxisInDegreesPerSecond uintptr
	Get_YAxisInDegreesPerSecond uintptr
	Put_YAxisInDegreesPerSecond uintptr
	Get_ZAxisInDegreesPerSecond uintptr
	Put_ZAxisInDegreesPerSecond uintptr
}

type IGyrometerDataThreshold struct {
	win32.IInspectable
}

func (this *IGyrometerDataThreshold) Vtbl() *IGyrometerDataThresholdVtbl {
	return (*IGyrometerDataThresholdVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGyrometerDataThreshold) Get_XAxisInDegreesPerSecond() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_XAxisInDegreesPerSecond, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGyrometerDataThreshold) Put_XAxisInDegreesPerSecond(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_XAxisInDegreesPerSecond, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGyrometerDataThreshold) Get_YAxisInDegreesPerSecond() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_YAxisInDegreesPerSecond, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGyrometerDataThreshold) Put_YAxisInDegreesPerSecond(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_YAxisInDegreesPerSecond, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGyrometerDataThreshold) Get_ZAxisInDegreesPerSecond() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ZAxisInDegreesPerSecond, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGyrometerDataThreshold) Put_ZAxisInDegreesPerSecond(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ZAxisInDegreesPerSecond, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 1EE5E978-89A2-4275-9E95-7126F4708760
var IID_IGyrometerDeviceId = syscall.GUID{0x1EE5E978, 0x89A2, 0x4275,
	[8]byte{0x9E, 0x95, 0x71, 0x26, 0xF4, 0x70, 0x87, 0x60}}

type IGyrometerDeviceIdInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
}

type IGyrometerDeviceIdVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId uintptr
}

type IGyrometerDeviceId struct {
	win32.IInspectable
}

func (this *IGyrometerDeviceId) Vtbl() *IGyrometerDeviceIdVtbl {
	return (*IGyrometerDeviceIdVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGyrometerDeviceId) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// B3D6DE5C-1EE4-456F-9DE7-E2493B5C8E03
var IID_IGyrometerReading = syscall.GUID{0xB3D6DE5C, 0x1EE4, 0x456F,
	[8]byte{0x9D, 0xE7, 0xE2, 0x49, 0x3B, 0x5C, 0x8E, 0x03}}

type IGyrometerReadingInterface interface {
	win32.IInspectableInterface
	Get_Timestamp() DateTime
	Get_AngularVelocityX() float64
	Get_AngularVelocityY() float64
	Get_AngularVelocityZ() float64
}

type IGyrometerReadingVtbl struct {
	win32.IInspectableVtbl
	Get_Timestamp        uintptr
	Get_AngularVelocityX uintptr
	Get_AngularVelocityY uintptr
	Get_AngularVelocityZ uintptr
}

type IGyrometerReading struct {
	win32.IInspectable
}

func (this *IGyrometerReading) Vtbl() *IGyrometerReadingVtbl {
	return (*IGyrometerReadingVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGyrometerReading) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGyrometerReading) Get_AngularVelocityX() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AngularVelocityX, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGyrometerReading) Get_AngularVelocityY() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AngularVelocityY, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGyrometerReading) Get_AngularVelocityZ() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AngularVelocityZ, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 16AFE13C-2B89-44BB-822B-D1E1556FF09B
var IID_IGyrometerReading2 = syscall.GUID{0x16AFE13C, 0x2B89, 0x44BB,
	[8]byte{0x82, 0x2B, 0xD1, 0xE1, 0x55, 0x6F, 0xF0, 0x9B}}

type IGyrometerReading2Interface interface {
	win32.IInspectableInterface
	Get_PerformanceCount() *IReference[TimeSpan]
	Get_Properties() *IMapView[string, interface{}]
}

type IGyrometerReading2Vtbl struct {
	win32.IInspectableVtbl
	Get_PerformanceCount uintptr
	Get_Properties       uintptr
}

type IGyrometerReading2 struct {
	win32.IInspectable
}

func (this *IGyrometerReading2) Vtbl() *IGyrometerReading2Vtbl {
	return (*IGyrometerReading2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGyrometerReading2) Get_PerformanceCount() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PerformanceCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGyrometerReading2) Get_Properties() *IMapView[string, interface{}] {
	var _result *IMapView[string, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0FDF1895-6F9E-42CE-8D58-388C0AB8356D
var IID_IGyrometerReadingChangedEventArgs = syscall.GUID{0x0FDF1895, 0x6F9E, 0x42CE,
	[8]byte{0x8D, 0x58, 0x38, 0x8C, 0x0A, 0xB8, 0x35, 0x6D}}

type IGyrometerReadingChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Reading() *IGyrometerReading
}

type IGyrometerReadingChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Reading uintptr
}

type IGyrometerReadingChangedEventArgs struct {
	win32.IInspectable
}

func (this *IGyrometerReadingChangedEventArgs) Vtbl() *IGyrometerReadingChangedEventArgsVtbl {
	return (*IGyrometerReadingChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGyrometerReadingChangedEventArgs) Get_Reading() *IGyrometerReading {
	var _result *IGyrometerReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Reading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 83B6E7C9-E49D-4B39-86E6-CD554BE4C5C1
var IID_IGyrometerStatics = syscall.GUID{0x83B6E7C9, 0xE49D, 0x4B39,
	[8]byte{0x86, 0xE6, 0xCD, 0x55, 0x4B, 0xE4, 0xC5, 0xC1}}

type IGyrometerStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *IGyrometer
}

type IGyrometerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
}

type IGyrometerStatics struct {
	win32.IInspectable
}

func (this *IGyrometerStatics) Vtbl() *IGyrometerStaticsVtbl {
	return (*IGyrometerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGyrometerStatics) GetDefault() *IGyrometer {
	var _result *IGyrometer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EF83F7A1-D700-4204-9613-79C6B161DF4E
var IID_IGyrometerStatics2 = syscall.GUID{0xEF83F7A1, 0xD700, 0x4204,
	[8]byte{0x96, 0x13, 0x79, 0xC6, 0xB1, 0x61, 0xDF, 0x4E}}

type IGyrometerStatics2Interface interface {
	win32.IInspectableInterface
	GetDeviceSelector() string
	FromIdAsync(deviceId string) *IAsyncOperation[*IGyrometer]
}

type IGyrometerStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelector uintptr
	FromIdAsync       uintptr
}

type IGyrometerStatics2 struct {
	win32.IInspectable
}

func (this *IGyrometerStatics2) Vtbl() *IGyrometerStatics2Vtbl {
	return (*IGyrometerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGyrometerStatics2) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGyrometerStatics2) FromIdAsync(deviceId string) *IAsyncOperation[*IGyrometer] {
	var _result *IAsyncOperation[*IGyrometer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A3CD45B9-1BF1-4F65-A704-E2DA04F182C0
var IID_IHingeAngleReading = syscall.GUID{0xA3CD45B9, 0x1BF1, 0x4F65,
	[8]byte{0xA7, 0x04, 0xE2, 0xDA, 0x04, 0xF1, 0x82, 0xC0}}

type IHingeAngleReadingInterface interface {
	win32.IInspectableInterface
	Get_Timestamp() DateTime
	Get_AngleInDegrees() float64
	Get_Properties() *IMapView[string, interface{}]
}

type IHingeAngleReadingVtbl struct {
	win32.IInspectableVtbl
	Get_Timestamp      uintptr
	Get_AngleInDegrees uintptr
	Get_Properties     uintptr
}

type IHingeAngleReading struct {
	win32.IInspectable
}

func (this *IHingeAngleReading) Vtbl() *IHingeAngleReadingVtbl {
	return (*IHingeAngleReadingVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHingeAngleReading) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHingeAngleReading) Get_AngleInDegrees() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AngleInDegrees, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHingeAngleReading) Get_Properties() *IMapView[string, interface{}] {
	var _result *IMapView[string, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E9D3BE02-BFDF-437F-8C29-88C77393D309
var IID_IHingeAngleSensor = syscall.GUID{0xE9D3BE02, 0xBFDF, 0x437F,
	[8]byte{0x8C, 0x29, 0x88, 0xC7, 0x73, 0x93, 0xD3, 0x09}}

type IHingeAngleSensorInterface interface {
	win32.IInspectableInterface
	GetCurrentReadingAsync() *IAsyncOperation[*IHingeAngleReading]
	Get_DeviceId() string
	Get_MinReportThresholdInDegrees() float64
	Get_ReportThresholdInDegrees() float64
	Put_ReportThresholdInDegrees(value float64)
	Add_ReadingChanged(handler TypedEventHandler[*IHingeAngleSensor, *IHingeAngleSensorReadingChangedEventArgs]) EventRegistrationToken
	Remove_ReadingChanged(token EventRegistrationToken)
}

type IHingeAngleSensorVtbl struct {
	win32.IInspectableVtbl
	GetCurrentReadingAsync          uintptr
	Get_DeviceId                    uintptr
	Get_MinReportThresholdInDegrees uintptr
	Get_ReportThresholdInDegrees    uintptr
	Put_ReportThresholdInDegrees    uintptr
	Add_ReadingChanged              uintptr
	Remove_ReadingChanged           uintptr
}

type IHingeAngleSensor struct {
	win32.IInspectable
}

func (this *IHingeAngleSensor) Vtbl() *IHingeAngleSensorVtbl {
	return (*IHingeAngleSensorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHingeAngleSensor) GetCurrentReadingAsync() *IAsyncOperation[*IHingeAngleReading] {
	var _result *IAsyncOperation[*IHingeAngleReading]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentReadingAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHingeAngleSensor) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHingeAngleSensor) Get_MinReportThresholdInDegrees() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinReportThresholdInDegrees, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHingeAngleSensor) Get_ReportThresholdInDegrees() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportThresholdInDegrees, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHingeAngleSensor) Put_ReportThresholdInDegrees(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReportThresholdInDegrees, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IHingeAngleSensor) Add_ReadingChanged(handler TypedEventHandler[*IHingeAngleSensor, *IHingeAngleSensorReadingChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ReadingChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHingeAngleSensor) Remove_ReadingChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ReadingChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 24D9558B-FAD0-42B8-A854-78923049A1BA
var IID_IHingeAngleSensorReadingChangedEventArgs = syscall.GUID{0x24D9558B, 0xFAD0, 0x42B8,
	[8]byte{0xA8, 0x54, 0x78, 0x92, 0x30, 0x49, 0xA1, 0xBA}}

type IHingeAngleSensorReadingChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Reading() *IHingeAngleReading
}

type IHingeAngleSensorReadingChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Reading uintptr
}

type IHingeAngleSensorReadingChangedEventArgs struct {
	win32.IInspectable
}

func (this *IHingeAngleSensorReadingChangedEventArgs) Vtbl() *IHingeAngleSensorReadingChangedEventArgsVtbl {
	return (*IHingeAngleSensorReadingChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHingeAngleSensorReadingChangedEventArgs) Get_Reading() *IHingeAngleReading {
	var _result *IHingeAngleReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Reading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B7B63910-FBB1-4123-89CE-4EA34EB0DFCA
var IID_IHingeAngleSensorStatics = syscall.GUID{0xB7B63910, 0xFBB1, 0x4123,
	[8]byte{0x89, 0xCE, 0x4E, 0xA3, 0x4E, 0xB0, 0xDF, 0xCA}}

type IHingeAngleSensorStaticsInterface interface {
	win32.IInspectableInterface
	GetDeviceSelector() string
	GetDefaultAsync() *IAsyncOperation[*IHingeAngleSensor]
	GetRelatedToAdjacentPanelsAsync(firstPanelId string, secondPanelId string) *IAsyncOperation[*IHingeAngleSensor]
	FromIdAsync(deviceId string) *IAsyncOperation[*IHingeAngleSensor]
}

type IHingeAngleSensorStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelector               uintptr
	GetDefaultAsync                 uintptr
	GetRelatedToAdjacentPanelsAsync uintptr
	FromIdAsync                     uintptr
}

type IHingeAngleSensorStatics struct {
	win32.IInspectable
}

func (this *IHingeAngleSensorStatics) Vtbl() *IHingeAngleSensorStaticsVtbl {
	return (*IHingeAngleSensorStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHingeAngleSensorStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHingeAngleSensorStatics) GetDefaultAsync() *IAsyncOperation[*IHingeAngleSensor] {
	var _result *IAsyncOperation[*IHingeAngleSensor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefaultAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHingeAngleSensorStatics) GetRelatedToAdjacentPanelsAsync(firstPanelId string, secondPanelId string) *IAsyncOperation[*IHingeAngleSensor] {
	var _result *IAsyncOperation[*IHingeAngleSensor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetRelatedToAdjacentPanelsAsync, uintptr(unsafe.Pointer(this)), NewHStr(firstPanelId).Ptr, NewHStr(secondPanelId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHingeAngleSensorStatics) FromIdAsync(deviceId string) *IAsyncOperation[*IHingeAngleSensor] {
	var _result *IAsyncOperation[*IHingeAngleSensor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2648CA6F-2286-406F-9161-F0C4BD806EBF
var IID_IInclinometer = syscall.GUID{0x2648CA6F, 0x2286, 0x406F,
	[8]byte{0x91, 0x61, 0xF0, 0xC4, 0xBD, 0x80, 0x6E, 0xBF}}

type IInclinometerInterface interface {
	win32.IInspectableInterface
	GetCurrentReading() *IInclinometerReading
	Get_MinimumReportInterval() uint32
	Put_ReportInterval(value uint32)
	Get_ReportInterval() uint32
	Add_ReadingChanged(handler TypedEventHandler[*IInclinometer, *IInclinometerReadingChangedEventArgs]) EventRegistrationToken
	Remove_ReadingChanged(token EventRegistrationToken)
}

type IInclinometerVtbl struct {
	win32.IInspectableVtbl
	GetCurrentReading         uintptr
	Get_MinimumReportInterval uintptr
	Put_ReportInterval        uintptr
	Get_ReportInterval        uintptr
	Add_ReadingChanged        uintptr
	Remove_ReadingChanged     uintptr
}

type IInclinometer struct {
	win32.IInspectable
}

func (this *IInclinometer) Vtbl() *IInclinometerVtbl {
	return (*IInclinometerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInclinometer) GetCurrentReading() *IInclinometerReading {
	var _result *IInclinometerReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentReading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInclinometer) Get_MinimumReportInterval() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinimumReportInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInclinometer) Put_ReportInterval(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReportInterval, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IInclinometer) Get_ReportInterval() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInclinometer) Add_ReadingChanged(handler TypedEventHandler[*IInclinometer, *IInclinometerReadingChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ReadingChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInclinometer) Remove_ReadingChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ReadingChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 029F3393-28B2-45F8-BB16-61E86A7FAE6E
var IID_IInclinometer2 = syscall.GUID{0x029F3393, 0x28B2, 0x45F8,
	[8]byte{0xBB, 0x16, 0x61, 0xE8, 0x6A, 0x7F, 0xAE, 0x6E}}

type IInclinometer2Interface interface {
	win32.IInspectableInterface
	Put_ReadingTransform(value unsafe.Pointer)
	Get_ReadingTransform() unsafe.Pointer
	Get_ReadingType() SensorReadingType
}

type IInclinometer2Vtbl struct {
	win32.IInspectableVtbl
	Put_ReadingTransform uintptr
	Get_ReadingTransform uintptr
	Get_ReadingType      uintptr
}

type IInclinometer2 struct {
	win32.IInspectable
}

func (this *IInclinometer2) Vtbl() *IInclinometer2Vtbl {
	return (*IInclinometer2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInclinometer2) Put_ReadingTransform(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReadingTransform, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IInclinometer2) Get_ReadingTransform() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReadingTransform, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInclinometer2) Get_ReadingType() SensorReadingType {
	var _result SensorReadingType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReadingType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 3A095004-D765-4384-A3D7-0283F3ABE6AE
var IID_IInclinometer3 = syscall.GUID{0x3A095004, 0xD765, 0x4384,
	[8]byte{0xA3, 0xD7, 0x02, 0x83, 0xF3, 0xAB, 0xE6, 0xAE}}

type IInclinometer3Interface interface {
	win32.IInspectableInterface
	Put_ReportLatency(value uint32)
	Get_ReportLatency() uint32
	Get_MaxBatchSize() uint32
}

type IInclinometer3Vtbl struct {
	win32.IInspectableVtbl
	Put_ReportLatency uintptr
	Get_ReportLatency uintptr
	Get_MaxBatchSize  uintptr
}

type IInclinometer3 struct {
	win32.IInspectable
}

func (this *IInclinometer3) Vtbl() *IInclinometer3Vtbl {
	return (*IInclinometer3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInclinometer3) Put_ReportLatency(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReportLatency, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IInclinometer3) Get_ReportLatency() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportLatency, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInclinometer3) Get_MaxBatchSize() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxBatchSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 43852618-8FCA-548E-BBF5-5C50412B6AA4
var IID_IInclinometer4 = syscall.GUID{0x43852618, 0x8FCA, 0x548E,
	[8]byte{0xBB, 0xF5, 0x5C, 0x50, 0x41, 0x2B, 0x6A, 0xA4}}

type IInclinometer4Interface interface {
	win32.IInspectableInterface
	Get_ReportThreshold() *IInclinometerDataThreshold
}

type IInclinometer4Vtbl struct {
	win32.IInspectableVtbl
	Get_ReportThreshold uintptr
}

type IInclinometer4 struct {
	win32.IInspectable
}

func (this *IInclinometer4) Vtbl() *IInclinometer4Vtbl {
	return (*IInclinometer4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInclinometer4) Get_ReportThreshold() *IInclinometerDataThreshold {
	var _result *IInclinometerDataThreshold
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportThreshold, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F80A4783-7BFE-545E-BB60-A0EBC47BD2FB
var IID_IInclinometerDataThreshold = syscall.GUID{0xF80A4783, 0x7BFE, 0x545E,
	[8]byte{0xBB, 0x60, 0xA0, 0xEB, 0xC4, 0x7B, 0xD2, 0xFB}}

type IInclinometerDataThresholdInterface interface {
	win32.IInspectableInterface
	Get_PitchInDegrees() float32
	Put_PitchInDegrees(value float32)
	Get_RollInDegrees() float32
	Put_RollInDegrees(value float32)
	Get_YawInDegrees() float32
	Put_YawInDegrees(value float32)
}

type IInclinometerDataThresholdVtbl struct {
	win32.IInspectableVtbl
	Get_PitchInDegrees uintptr
	Put_PitchInDegrees uintptr
	Get_RollInDegrees  uintptr
	Put_RollInDegrees  uintptr
	Get_YawInDegrees   uintptr
	Put_YawInDegrees   uintptr
}

type IInclinometerDataThreshold struct {
	win32.IInspectable
}

func (this *IInclinometerDataThreshold) Vtbl() *IInclinometerDataThresholdVtbl {
	return (*IInclinometerDataThresholdVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInclinometerDataThreshold) Get_PitchInDegrees() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PitchInDegrees, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInclinometerDataThreshold) Put_PitchInDegrees(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PitchInDegrees, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IInclinometerDataThreshold) Get_RollInDegrees() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RollInDegrees, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInclinometerDataThreshold) Put_RollInDegrees(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RollInDegrees, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IInclinometerDataThreshold) Get_YawInDegrees() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_YawInDegrees, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInclinometerDataThreshold) Put_YawInDegrees(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_YawInDegrees, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 01E91982-41FF-4406-AE83-62210FF16FE3
var IID_IInclinometerDeviceId = syscall.GUID{0x01E91982, 0x41FF, 0x4406,
	[8]byte{0xAE, 0x83, 0x62, 0x21, 0x0F, 0xF1, 0x6F, 0xE3}}

type IInclinometerDeviceIdInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
}

type IInclinometerDeviceIdVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId uintptr
}

type IInclinometerDeviceId struct {
	win32.IInspectable
}

func (this *IInclinometerDeviceId) Vtbl() *IInclinometerDeviceIdVtbl {
	return (*IInclinometerDeviceIdVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInclinometerDeviceId) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 9F44F055-B6F6-497F-B127-1A775E501458
var IID_IInclinometerReading = syscall.GUID{0x9F44F055, 0xB6F6, 0x497F,
	[8]byte{0xB1, 0x27, 0x1A, 0x77, 0x5E, 0x50, 0x14, 0x58}}

type IInclinometerReadingInterface interface {
	win32.IInspectableInterface
	Get_Timestamp() DateTime
	Get_PitchDegrees() float32
	Get_RollDegrees() float32
	Get_YawDegrees() float32
}

type IInclinometerReadingVtbl struct {
	win32.IInspectableVtbl
	Get_Timestamp    uintptr
	Get_PitchDegrees uintptr
	Get_RollDegrees  uintptr
	Get_YawDegrees   uintptr
}

type IInclinometerReading struct {
	win32.IInspectable
}

func (this *IInclinometerReading) Vtbl() *IInclinometerReadingVtbl {
	return (*IInclinometerReadingVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInclinometerReading) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInclinometerReading) Get_PitchDegrees() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PitchDegrees, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInclinometerReading) Get_RollDegrees() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RollDegrees, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInclinometerReading) Get_YawDegrees() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_YawDegrees, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4F164781-E90B-4658-8915-0103E08A805A
var IID_IInclinometerReading2 = syscall.GUID{0x4F164781, 0xE90B, 0x4658,
	[8]byte{0x89, 0x15, 0x01, 0x03, 0xE0, 0x8A, 0x80, 0x5A}}

type IInclinometerReading2Interface interface {
	win32.IInspectableInterface
	Get_PerformanceCount() *IReference[TimeSpan]
	Get_Properties() *IMapView[string, interface{}]
}

type IInclinometerReading2Vtbl struct {
	win32.IInspectableVtbl
	Get_PerformanceCount uintptr
	Get_Properties       uintptr
}

type IInclinometerReading2 struct {
	win32.IInspectable
}

func (this *IInclinometerReading2) Vtbl() *IInclinometerReading2Vtbl {
	return (*IInclinometerReading2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInclinometerReading2) Get_PerformanceCount() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PerformanceCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInclinometerReading2) Get_Properties() *IMapView[string, interface{}] {
	var _result *IMapView[string, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4AE91DC1-E7EB-4938-8511-AE0D6B440438
var IID_IInclinometerReadingChangedEventArgs = syscall.GUID{0x4AE91DC1, 0xE7EB, 0x4938,
	[8]byte{0x85, 0x11, 0xAE, 0x0D, 0x6B, 0x44, 0x04, 0x38}}

type IInclinometerReadingChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Reading() *IInclinometerReading
}

type IInclinometerReadingChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Reading uintptr
}

type IInclinometerReadingChangedEventArgs struct {
	win32.IInspectable
}

func (this *IInclinometerReadingChangedEventArgs) Vtbl() *IInclinometerReadingChangedEventArgsVtbl {
	return (*IInclinometerReadingChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInclinometerReadingChangedEventArgs) Get_Reading() *IInclinometerReading {
	var _result *IInclinometerReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Reading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B453E880-1FE3-4986-A257-E6ECE2723949
var IID_IInclinometerReadingYawAccuracy = syscall.GUID{0xB453E880, 0x1FE3, 0x4986,
	[8]byte{0xA2, 0x57, 0xE6, 0xEC, 0xE2, 0x72, 0x39, 0x49}}

type IInclinometerReadingYawAccuracyInterface interface {
	win32.IInspectableInterface
	Get_YawAccuracy() MagnetometerAccuracy
}

type IInclinometerReadingYawAccuracyVtbl struct {
	win32.IInspectableVtbl
	Get_YawAccuracy uintptr
}

type IInclinometerReadingYawAccuracy struct {
	win32.IInspectable
}

func (this *IInclinometerReadingYawAccuracy) Vtbl() *IInclinometerReadingYawAccuracyVtbl {
	return (*IInclinometerReadingYawAccuracyVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInclinometerReadingYawAccuracy) Get_YawAccuracy() MagnetometerAccuracy {
	var _result MagnetometerAccuracy
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_YawAccuracy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F22EC551-9C30-453A-8B49-3C3EEB33CB61
var IID_IInclinometerStatics = syscall.GUID{0xF22EC551, 0x9C30, 0x453A,
	[8]byte{0x8B, 0x49, 0x3C, 0x3E, 0xEB, 0x33, 0xCB, 0x61}}

type IInclinometerStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *IInclinometer
}

type IInclinometerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
}

type IInclinometerStatics struct {
	win32.IInspectable
}

func (this *IInclinometerStatics) Vtbl() *IInclinometerStaticsVtbl {
	return (*IInclinometerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInclinometerStatics) GetDefault() *IInclinometer {
	var _result *IInclinometer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 043F9775-6A1E-499C-86E0-638C1A864B00
var IID_IInclinometerStatics2 = syscall.GUID{0x043F9775, 0x6A1E, 0x499C,
	[8]byte{0x86, 0xE0, 0x63, 0x8C, 0x1A, 0x86, 0x4B, 0x00}}

type IInclinometerStatics2Interface interface {
	win32.IInspectableInterface
	GetDefaultForRelativeReadings() *IInclinometer
}

type IInclinometerStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetDefaultForRelativeReadings uintptr
}

type IInclinometerStatics2 struct {
	win32.IInspectable
}

func (this *IInclinometerStatics2) Vtbl() *IInclinometerStatics2Vtbl {
	return (*IInclinometerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInclinometerStatics2) GetDefaultForRelativeReadings() *IInclinometer {
	var _result *IInclinometer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefaultForRelativeReadings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BD9A4280-B91A-4829-9392-ABC0B6BDF2B4
var IID_IInclinometerStatics3 = syscall.GUID{0xBD9A4280, 0xB91A, 0x4829,
	[8]byte{0x93, 0x92, 0xAB, 0xC0, 0xB6, 0xBD, 0xF2, 0xB4}}

type IInclinometerStatics3Interface interface {
	win32.IInspectableInterface
	GetDefaultWithSensorReadingType(sensorReadingtype SensorReadingType) *IInclinometer
}

type IInclinometerStatics3Vtbl struct {
	win32.IInspectableVtbl
	GetDefaultWithSensorReadingType uintptr
}

type IInclinometerStatics3 struct {
	win32.IInspectable
}

func (this *IInclinometerStatics3) Vtbl() *IInclinometerStatics3Vtbl {
	return (*IInclinometerStatics3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInclinometerStatics3) GetDefaultWithSensorReadingType(sensorReadingtype SensorReadingType) *IInclinometer {
	var _result *IInclinometer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefaultWithSensorReadingType, uintptr(unsafe.Pointer(this)), uintptr(sensorReadingtype), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E8BA96F9-6E85-4A83-AED0-D7CDCC9856C8
var IID_IInclinometerStatics4 = syscall.GUID{0xE8BA96F9, 0x6E85, 0x4A83,
	[8]byte{0xAE, 0xD0, 0xD7, 0xCD, 0xCC, 0x98, 0x56, 0xC8}}

type IInclinometerStatics4Interface interface {
	win32.IInspectableInterface
	GetDeviceSelector(readingType SensorReadingType) string
	FromIdAsync(deviceId string) *IAsyncOperation[*IInclinometer]
}

type IInclinometerStatics4Vtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelector uintptr
	FromIdAsync       uintptr
}

type IInclinometerStatics4 struct {
	win32.IInspectable
}

func (this *IInclinometerStatics4) Vtbl() *IInclinometerStatics4Vtbl {
	return (*IInclinometerStatics4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInclinometerStatics4) GetDeviceSelector(readingType SensorReadingType) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(readingType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IInclinometerStatics4) FromIdAsync(deviceId string) *IAsyncOperation[*IInclinometer] {
	var _result *IAsyncOperation[*IInclinometer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F84C0718-0C54-47AE-922E-789F57FB03A0
var IID_ILightSensor = syscall.GUID{0xF84C0718, 0x0C54, 0x47AE,
	[8]byte{0x92, 0x2E, 0x78, 0x9F, 0x57, 0xFB, 0x03, 0xA0}}

type ILightSensorInterface interface {
	win32.IInspectableInterface
	GetCurrentReading() *ILightSensorReading
	Get_MinimumReportInterval() uint32
	Put_ReportInterval(value uint32)
	Get_ReportInterval() uint32
	Add_ReadingChanged(handler TypedEventHandler[*ILightSensor, *ILightSensorReadingChangedEventArgs]) EventRegistrationToken
	Remove_ReadingChanged(token EventRegistrationToken)
}

type ILightSensorVtbl struct {
	win32.IInspectableVtbl
	GetCurrentReading         uintptr
	Get_MinimumReportInterval uintptr
	Put_ReportInterval        uintptr
	Get_ReportInterval        uintptr
	Add_ReadingChanged        uintptr
	Remove_ReadingChanged     uintptr
}

type ILightSensor struct {
	win32.IInspectable
}

func (this *ILightSensor) Vtbl() *ILightSensorVtbl {
	return (*ILightSensorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILightSensor) GetCurrentReading() *ILightSensorReading {
	var _result *ILightSensorReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentReading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILightSensor) Get_MinimumReportInterval() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinimumReportInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILightSensor) Put_ReportInterval(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReportInterval, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ILightSensor) Get_ReportInterval() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILightSensor) Add_ReadingChanged(handler TypedEventHandler[*ILightSensor, *ILightSensorReadingChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ReadingChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILightSensor) Remove_ReadingChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ReadingChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 486B24E8-A94C-4090-8F48-09F782A9F7D5
var IID_ILightSensor2 = syscall.GUID{0x486B24E8, 0xA94C, 0x4090,
	[8]byte{0x8F, 0x48, 0x09, 0xF7, 0x82, 0xA9, 0xF7, 0xD5}}

type ILightSensor2Interface interface {
	win32.IInspectableInterface
	Put_ReportLatency(value uint32)
	Get_ReportLatency() uint32
	Get_MaxBatchSize() uint32
}

type ILightSensor2Vtbl struct {
	win32.IInspectableVtbl
	Put_ReportLatency uintptr
	Get_ReportLatency uintptr
	Get_MaxBatchSize  uintptr
}

type ILightSensor2 struct {
	win32.IInspectable
}

func (this *ILightSensor2) Vtbl() *ILightSensor2Vtbl {
	return (*ILightSensor2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILightSensor2) Put_ReportLatency(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReportLatency, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ILightSensor2) Get_ReportLatency() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportLatency, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILightSensor2) Get_MaxBatchSize() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxBatchSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4876D0FF-9F4C-5F72-ADBD-A3471B063C00
var IID_ILightSensor3 = syscall.GUID{0x4876D0FF, 0x9F4C, 0x5F72,
	[8]byte{0xAD, 0xBD, 0xA3, 0x47, 0x1B, 0x06, 0x3C, 0x00}}

type ILightSensor3Interface interface {
	win32.IInspectableInterface
	Get_ReportThreshold() *ILightSensorDataThreshold
}

type ILightSensor3Vtbl struct {
	win32.IInspectableVtbl
	Get_ReportThreshold uintptr
}

type ILightSensor3 struct {
	win32.IInspectable
}

func (this *ILightSensor3) Vtbl() *ILightSensor3Vtbl {
	return (*ILightSensor3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILightSensor3) Get_ReportThreshold() *ILightSensorDataThreshold {
	var _result *ILightSensorDataThreshold
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportThreshold, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B160AFD1-878F-5492-9F2C-33DC3AE584A3
var IID_ILightSensorDataThreshold = syscall.GUID{0xB160AFD1, 0x878F, 0x5492,
	[8]byte{0x9F, 0x2C, 0x33, 0xDC, 0x3A, 0xE5, 0x84, 0xA3}}

type ILightSensorDataThresholdInterface interface {
	win32.IInspectableInterface
	Get_LuxPercentage() float32
	Put_LuxPercentage(value float32)
	Get_AbsoluteLux() float32
	Put_AbsoluteLux(value float32)
}

type ILightSensorDataThresholdVtbl struct {
	win32.IInspectableVtbl
	Get_LuxPercentage uintptr
	Put_LuxPercentage uintptr
	Get_AbsoluteLux   uintptr
	Put_AbsoluteLux   uintptr
}

type ILightSensorDataThreshold struct {
	win32.IInspectable
}

func (this *ILightSensorDataThreshold) Vtbl() *ILightSensorDataThresholdVtbl {
	return (*ILightSensorDataThresholdVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILightSensorDataThreshold) Get_LuxPercentage() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LuxPercentage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILightSensorDataThreshold) Put_LuxPercentage(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LuxPercentage, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ILightSensorDataThreshold) Get_AbsoluteLux() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AbsoluteLux, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILightSensorDataThreshold) Put_AbsoluteLux(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AbsoluteLux, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 7FEE49F8-0AFB-4F51-87F0-6C26375CE94F
var IID_ILightSensorDeviceId = syscall.GUID{0x7FEE49F8, 0x0AFB, 0x4F51,
	[8]byte{0x87, 0xF0, 0x6C, 0x26, 0x37, 0x5C, 0xE9, 0x4F}}

type ILightSensorDeviceIdInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
}

type ILightSensorDeviceIdVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId uintptr
}

type ILightSensorDeviceId struct {
	win32.IInspectable
}

func (this *ILightSensorDeviceId) Vtbl() *ILightSensorDeviceIdVtbl {
	return (*ILightSensorDeviceIdVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILightSensorDeviceId) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// FFDF6300-227C-4D2B-B302-FC0142485C68
var IID_ILightSensorReading = syscall.GUID{0xFFDF6300, 0x227C, 0x4D2B,
	[8]byte{0xB3, 0x02, 0xFC, 0x01, 0x42, 0x48, 0x5C, 0x68}}

type ILightSensorReadingInterface interface {
	win32.IInspectableInterface
	Get_Timestamp() DateTime
	Get_IlluminanceInLux() float32
}

type ILightSensorReadingVtbl struct {
	win32.IInspectableVtbl
	Get_Timestamp        uintptr
	Get_IlluminanceInLux uintptr
}

type ILightSensorReading struct {
	win32.IInspectable
}

func (this *ILightSensorReading) Vtbl() *ILightSensorReadingVtbl {
	return (*ILightSensorReadingVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILightSensorReading) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILightSensorReading) Get_IlluminanceInLux() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IlluminanceInLux, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// B7512185-44A3-44C9-8190-9EF6DE0A8A74
var IID_ILightSensorReading2 = syscall.GUID{0xB7512185, 0x44A3, 0x44C9,
	[8]byte{0x81, 0x90, 0x9E, 0xF6, 0xDE, 0x0A, 0x8A, 0x74}}

type ILightSensorReading2Interface interface {
	win32.IInspectableInterface
	Get_PerformanceCount() *IReference[TimeSpan]
	Get_Properties() *IMapView[string, interface{}]
}

type ILightSensorReading2Vtbl struct {
	win32.IInspectableVtbl
	Get_PerformanceCount uintptr
	Get_Properties       uintptr
}

type ILightSensorReading2 struct {
	win32.IInspectable
}

func (this *ILightSensorReading2) Vtbl() *ILightSensorReading2Vtbl {
	return (*ILightSensorReading2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILightSensorReading2) Get_PerformanceCount() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PerformanceCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILightSensorReading2) Get_Properties() *IMapView[string, interface{}] {
	var _result *IMapView[string, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A3A2F4CF-258B-420C-B8AB-8EDD601ECF50
var IID_ILightSensorReadingChangedEventArgs = syscall.GUID{0xA3A2F4CF, 0x258B, 0x420C,
	[8]byte{0xB8, 0xAB, 0x8E, 0xDD, 0x60, 0x1E, 0xCF, 0x50}}

type ILightSensorReadingChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Reading() *ILightSensorReading
}

type ILightSensorReadingChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Reading uintptr
}

type ILightSensorReadingChangedEventArgs struct {
	win32.IInspectable
}

func (this *ILightSensorReadingChangedEventArgs) Vtbl() *ILightSensorReadingChangedEventArgsVtbl {
	return (*ILightSensorReadingChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILightSensorReadingChangedEventArgs) Get_Reading() *ILightSensorReading {
	var _result *ILightSensorReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Reading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 45DB8C84-C3A8-471E-9A53-6457FAD87C0E
var IID_ILightSensorStatics = syscall.GUID{0x45DB8C84, 0xC3A8, 0x471E,
	[8]byte{0x9A, 0x53, 0x64, 0x57, 0xFA, 0xD8, 0x7C, 0x0E}}

type ILightSensorStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *ILightSensor
}

type ILightSensorStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
}

type ILightSensorStatics struct {
	win32.IInspectable
}

func (this *ILightSensorStatics) Vtbl() *ILightSensorStaticsVtbl {
	return (*ILightSensorStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILightSensorStatics) GetDefault() *ILightSensor {
	var _result *ILightSensor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0EC0A650-DDC6-40AB-ACE3-EC3359D42C51
var IID_ILightSensorStatics2 = syscall.GUID{0x0EC0A650, 0xDDC6, 0x40AB,
	[8]byte{0xAC, 0xE3, 0xEC, 0x33, 0x59, 0xD4, 0x2C, 0x51}}

type ILightSensorStatics2Interface interface {
	win32.IInspectableInterface
	GetDeviceSelector() string
	FromIdAsync(deviceId string) *IAsyncOperation[*ILightSensor]
}

type ILightSensorStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelector uintptr
	FromIdAsync       uintptr
}

type ILightSensorStatics2 struct {
	win32.IInspectable
}

func (this *ILightSensorStatics2) Vtbl() *ILightSensorStatics2Vtbl {
	return (*ILightSensorStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILightSensorStatics2) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILightSensorStatics2) FromIdAsync(deviceId string) *IAsyncOperation[*ILightSensor] {
	var _result *IAsyncOperation[*ILightSensor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 484F626E-D3C9-4111-B3F6-2CF1FAA418D5
var IID_IMagnetometer = syscall.GUID{0x484F626E, 0xD3C9, 0x4111,
	[8]byte{0xB3, 0xF6, 0x2C, 0xF1, 0xFA, 0xA4, 0x18, 0xD5}}

type IMagnetometerInterface interface {
	win32.IInspectableInterface
	GetCurrentReading() *IMagnetometerReading
	Get_MinimumReportInterval() uint32
	Put_ReportInterval(value uint32)
	Get_ReportInterval() uint32
	Add_ReadingChanged(handler TypedEventHandler[*IMagnetometer, *IMagnetometerReadingChangedEventArgs]) EventRegistrationToken
	Remove_ReadingChanged(token EventRegistrationToken)
}

type IMagnetometerVtbl struct {
	win32.IInspectableVtbl
	GetCurrentReading         uintptr
	Get_MinimumReportInterval uintptr
	Put_ReportInterval        uintptr
	Get_ReportInterval        uintptr
	Add_ReadingChanged        uintptr
	Remove_ReadingChanged     uintptr
}

type IMagnetometer struct {
	win32.IInspectable
}

func (this *IMagnetometer) Vtbl() *IMagnetometerVtbl {
	return (*IMagnetometerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMagnetometer) GetCurrentReading() *IMagnetometerReading {
	var _result *IMagnetometerReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentReading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMagnetometer) Get_MinimumReportInterval() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinimumReportInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagnetometer) Put_ReportInterval(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReportInterval, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMagnetometer) Get_ReportInterval() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagnetometer) Add_ReadingChanged(handler TypedEventHandler[*IMagnetometer, *IMagnetometerReadingChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ReadingChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagnetometer) Remove_ReadingChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ReadingChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// B4656C85-26F6-444B-A9E2-A23F966CD368
var IID_IMagnetometer2 = syscall.GUID{0xB4656C85, 0x26F6, 0x444B,
	[8]byte{0xA9, 0xE2, 0xA2, 0x3F, 0x96, 0x6C, 0xD3, 0x68}}

type IMagnetometer2Interface interface {
	win32.IInspectableInterface
	Put_ReadingTransform(value unsafe.Pointer)
	Get_ReadingTransform() unsafe.Pointer
}

type IMagnetometer2Vtbl struct {
	win32.IInspectableVtbl
	Put_ReadingTransform uintptr
	Get_ReadingTransform uintptr
}

type IMagnetometer2 struct {
	win32.IInspectable
}

func (this *IMagnetometer2) Vtbl() *IMagnetometer2Vtbl {
	return (*IMagnetometer2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMagnetometer2) Put_ReadingTransform(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReadingTransform, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMagnetometer2) Get_ReadingTransform() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReadingTransform, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// BE93DB7C-A625-48EF-ACF7-FAC104832671
var IID_IMagnetometer3 = syscall.GUID{0xBE93DB7C, 0xA625, 0x48EF,
	[8]byte{0xAC, 0xF7, 0xFA, 0xC1, 0x04, 0x83, 0x26, 0x71}}

type IMagnetometer3Interface interface {
	win32.IInspectableInterface
	Put_ReportLatency(value uint32)
	Get_ReportLatency() uint32
	Get_MaxBatchSize() uint32
}

type IMagnetometer3Vtbl struct {
	win32.IInspectableVtbl
	Put_ReportLatency uintptr
	Get_ReportLatency uintptr
	Get_MaxBatchSize  uintptr
}

type IMagnetometer3 struct {
	win32.IInspectable
}

func (this *IMagnetometer3) Vtbl() *IMagnetometer3Vtbl {
	return (*IMagnetometer3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMagnetometer3) Put_ReportLatency(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReportLatency, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMagnetometer3) Get_ReportLatency() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportLatency, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagnetometer3) Get_MaxBatchSize() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxBatchSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// DFB17901-3E0F-508F-B24B-F2BB75015F40
var IID_IMagnetometer4 = syscall.GUID{0xDFB17901, 0x3E0F, 0x508F,
	[8]byte{0xB2, 0x4B, 0xF2, 0xBB, 0x75, 0x01, 0x5F, 0x40}}

type IMagnetometer4Interface interface {
	win32.IInspectableInterface
	Get_ReportThreshold() *IMagnetometerDataThreshold
}

type IMagnetometer4Vtbl struct {
	win32.IInspectableVtbl
	Get_ReportThreshold uintptr
}

type IMagnetometer4 struct {
	win32.IInspectable
}

func (this *IMagnetometer4) Vtbl() *IMagnetometer4Vtbl {
	return (*IMagnetometer4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMagnetometer4) Get_ReportThreshold() *IMagnetometerDataThreshold {
	var _result *IMagnetometerDataThreshold
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportThreshold, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D177CB01-9063-5FA5-B596-B445E9DC3401
var IID_IMagnetometerDataThreshold = syscall.GUID{0xD177CB01, 0x9063, 0x5FA5,
	[8]byte{0xB5, 0x96, 0xB4, 0x45, 0xE9, 0xDC, 0x34, 0x01}}

type IMagnetometerDataThresholdInterface interface {
	win32.IInspectableInterface
	Get_XAxisMicroteslas() float32
	Put_XAxisMicroteslas(value float32)
	Get_YAxisMicroteslas() float32
	Put_YAxisMicroteslas(value float32)
	Get_ZAxisMicroteslas() float32
	Put_ZAxisMicroteslas(value float32)
}

type IMagnetometerDataThresholdVtbl struct {
	win32.IInspectableVtbl
	Get_XAxisMicroteslas uintptr
	Put_XAxisMicroteslas uintptr
	Get_YAxisMicroteslas uintptr
	Put_YAxisMicroteslas uintptr
	Get_ZAxisMicroteslas uintptr
	Put_ZAxisMicroteslas uintptr
}

type IMagnetometerDataThreshold struct {
	win32.IInspectable
}

func (this *IMagnetometerDataThreshold) Vtbl() *IMagnetometerDataThresholdVtbl {
	return (*IMagnetometerDataThresholdVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMagnetometerDataThreshold) Get_XAxisMicroteslas() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_XAxisMicroteslas, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagnetometerDataThreshold) Put_XAxisMicroteslas(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_XAxisMicroteslas, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMagnetometerDataThreshold) Get_YAxisMicroteslas() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_YAxisMicroteslas, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagnetometerDataThreshold) Put_YAxisMicroteslas(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_YAxisMicroteslas, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMagnetometerDataThreshold) Get_ZAxisMicroteslas() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ZAxisMicroteslas, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagnetometerDataThreshold) Put_ZAxisMicroteslas(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ZAxisMicroteslas, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 58B498C2-7E4B-404C-9FC5-5DE8B40EBAE3
var IID_IMagnetometerDeviceId = syscall.GUID{0x58B498C2, 0x7E4B, 0x404C,
	[8]byte{0x9F, 0xC5, 0x5D, 0xE8, 0xB4, 0x0E, 0xBA, 0xE3}}

type IMagnetometerDeviceIdInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
}

type IMagnetometerDeviceIdVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId uintptr
}

type IMagnetometerDeviceId struct {
	win32.IInspectable
}

func (this *IMagnetometerDeviceId) Vtbl() *IMagnetometerDeviceIdVtbl {
	return (*IMagnetometerDeviceIdVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMagnetometerDeviceId) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 0C2CC40D-EBFD-4E5C-BB11-AFC29B3CAE61
var IID_IMagnetometerReading = syscall.GUID{0x0C2CC40D, 0xEBFD, 0x4E5C,
	[8]byte{0xBB, 0x11, 0xAF, 0xC2, 0x9B, 0x3C, 0xAE, 0x61}}

type IMagnetometerReadingInterface interface {
	win32.IInspectableInterface
	Get_Timestamp() DateTime
	Get_MagneticFieldX() float32
	Get_MagneticFieldY() float32
	Get_MagneticFieldZ() float32
	Get_DirectionalAccuracy() MagnetometerAccuracy
}

type IMagnetometerReadingVtbl struct {
	win32.IInspectableVtbl
	Get_Timestamp           uintptr
	Get_MagneticFieldX      uintptr
	Get_MagneticFieldY      uintptr
	Get_MagneticFieldZ      uintptr
	Get_DirectionalAccuracy uintptr
}

type IMagnetometerReading struct {
	win32.IInspectable
}

func (this *IMagnetometerReading) Vtbl() *IMagnetometerReadingVtbl {
	return (*IMagnetometerReadingVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMagnetometerReading) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagnetometerReading) Get_MagneticFieldX() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MagneticFieldX, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagnetometerReading) Get_MagneticFieldY() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MagneticFieldY, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagnetometerReading) Get_MagneticFieldZ() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MagneticFieldZ, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMagnetometerReading) Get_DirectionalAccuracy() MagnetometerAccuracy {
	var _result MagnetometerAccuracy
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DirectionalAccuracy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D4C95C61-61D9-404B-A328-066F177A1409
var IID_IMagnetometerReading2 = syscall.GUID{0xD4C95C61, 0x61D9, 0x404B,
	[8]byte{0xA3, 0x28, 0x06, 0x6F, 0x17, 0x7A, 0x14, 0x09}}

type IMagnetometerReading2Interface interface {
	win32.IInspectableInterface
	Get_PerformanceCount() *IReference[TimeSpan]
	Get_Properties() *IMapView[string, interface{}]
}

type IMagnetometerReading2Vtbl struct {
	win32.IInspectableVtbl
	Get_PerformanceCount uintptr
	Get_Properties       uintptr
}

type IMagnetometerReading2 struct {
	win32.IInspectable
}

func (this *IMagnetometerReading2) Vtbl() *IMagnetometerReading2Vtbl {
	return (*IMagnetometerReading2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMagnetometerReading2) Get_PerformanceCount() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PerformanceCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMagnetometerReading2) Get_Properties() *IMapView[string, interface{}] {
	var _result *IMapView[string, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 17EAE872-2EB9-4EE7-8AD0-3127537D949B
var IID_IMagnetometerReadingChangedEventArgs = syscall.GUID{0x17EAE872, 0x2EB9, 0x4EE7,
	[8]byte{0x8A, 0xD0, 0x31, 0x27, 0x53, 0x7D, 0x94, 0x9B}}

type IMagnetometerReadingChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Reading() *IMagnetometerReading
}

type IMagnetometerReadingChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Reading uintptr
}

type IMagnetometerReadingChangedEventArgs struct {
	win32.IInspectable
}

func (this *IMagnetometerReadingChangedEventArgs) Vtbl() *IMagnetometerReadingChangedEventArgsVtbl {
	return (*IMagnetometerReadingChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMagnetometerReadingChangedEventArgs) Get_Reading() *IMagnetometerReading {
	var _result *IMagnetometerReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Reading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 853C64CC-0698-4DDA-A6DF-9CB9CC4AB40A
var IID_IMagnetometerStatics = syscall.GUID{0x853C64CC, 0x0698, 0x4DDA,
	[8]byte{0xA6, 0xDF, 0x9C, 0xB9, 0xCC, 0x4A, 0xB4, 0x0A}}

type IMagnetometerStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *IMagnetometer
}

type IMagnetometerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
}

type IMagnetometerStatics struct {
	win32.IInspectable
}

func (this *IMagnetometerStatics) Vtbl() *IMagnetometerStaticsVtbl {
	return (*IMagnetometerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMagnetometerStatics) GetDefault() *IMagnetometer {
	var _result *IMagnetometer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2C0819F0-FFC6-4F89-A06F-18FA10792933
var IID_IMagnetometerStatics2 = syscall.GUID{0x2C0819F0, 0xFFC6, 0x4F89,
	[8]byte{0xA0, 0x6F, 0x18, 0xFA, 0x10, 0x79, 0x29, 0x33}}

type IMagnetometerStatics2Interface interface {
	win32.IInspectableInterface
	GetDeviceSelector() string
	FromIdAsync(deviceId string) *IAsyncOperation[*IMagnetometer]
}

type IMagnetometerStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelector uintptr
	FromIdAsync       uintptr
}

type IMagnetometerStatics2 struct {
	win32.IInspectable
}

func (this *IMagnetometerStatics2) Vtbl() *IMagnetometerStatics2Vtbl {
	return (*IMagnetometerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMagnetometerStatics2) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMagnetometerStatics2) FromIdAsync(deviceId string) *IAsyncOperation[*IMagnetometer] {
	var _result *IAsyncOperation[*IMagnetometer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5E354635-CF6B-4C63-ABD8-10252B0BF6EC
var IID_IOrientationSensor = syscall.GUID{0x5E354635, 0xCF6B, 0x4C63,
	[8]byte{0xAB, 0xD8, 0x10, 0x25, 0x2B, 0x0B, 0xF6, 0xEC}}

type IOrientationSensorInterface interface {
	win32.IInspectableInterface
	GetCurrentReading() *IOrientationSensorReading
	Get_MinimumReportInterval() uint32
	Put_ReportInterval(value uint32)
	Get_ReportInterval() uint32
	Add_ReadingChanged(handler TypedEventHandler[*IOrientationSensor, *IOrientationSensorReadingChangedEventArgs]) EventRegistrationToken
	Remove_ReadingChanged(token EventRegistrationToken)
}

type IOrientationSensorVtbl struct {
	win32.IInspectableVtbl
	GetCurrentReading         uintptr
	Get_MinimumReportInterval uintptr
	Put_ReportInterval        uintptr
	Get_ReportInterval        uintptr
	Add_ReadingChanged        uintptr
	Remove_ReadingChanged     uintptr
}

type IOrientationSensor struct {
	win32.IInspectable
}

func (this *IOrientationSensor) Vtbl() *IOrientationSensorVtbl {
	return (*IOrientationSensorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IOrientationSensor) GetCurrentReading() *IOrientationSensorReading {
	var _result *IOrientationSensorReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentReading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IOrientationSensor) Get_MinimumReportInterval() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinimumReportInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IOrientationSensor) Put_ReportInterval(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReportInterval, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IOrientationSensor) Get_ReportInterval() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IOrientationSensor) Add_ReadingChanged(handler TypedEventHandler[*IOrientationSensor, *IOrientationSensorReadingChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ReadingChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IOrientationSensor) Remove_ReadingChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ReadingChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 0D924CF9-2F1F-49C9-8042-4A1813D67760
var IID_IOrientationSensor2 = syscall.GUID{0x0D924CF9, 0x2F1F, 0x49C9,
	[8]byte{0x80, 0x42, 0x4A, 0x18, 0x13, 0xD6, 0x77, 0x60}}

type IOrientationSensor2Interface interface {
	win32.IInspectableInterface
	Put_ReadingTransform(value unsafe.Pointer)
	Get_ReadingTransform() unsafe.Pointer
	Get_ReadingType() SensorReadingType
}

type IOrientationSensor2Vtbl struct {
	win32.IInspectableVtbl
	Put_ReadingTransform uintptr
	Get_ReadingTransform uintptr
	Get_ReadingType      uintptr
}

type IOrientationSensor2 struct {
	win32.IInspectable
}

func (this *IOrientationSensor2) Vtbl() *IOrientationSensor2Vtbl {
	return (*IOrientationSensor2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IOrientationSensor2) Put_ReadingTransform(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReadingTransform, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IOrientationSensor2) Get_ReadingTransform() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReadingTransform, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IOrientationSensor2) Get_ReadingType() SensorReadingType {
	var _result SensorReadingType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReadingType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 2CCE578D-646B-48C5-B7EE-44FDC4C6AAFD
var IID_IOrientationSensor3 = syscall.GUID{0x2CCE578D, 0x646B, 0x48C5,
	[8]byte{0xB7, 0xEE, 0x44, 0xFD, 0xC4, 0xC6, 0xAA, 0xFD}}

type IOrientationSensor3Interface interface {
	win32.IInspectableInterface
	Put_ReportLatency(value uint32)
	Get_ReportLatency() uint32
	Get_MaxBatchSize() uint32
}

type IOrientationSensor3Vtbl struct {
	win32.IInspectableVtbl
	Put_ReportLatency uintptr
	Get_ReportLatency uintptr
	Get_MaxBatchSize  uintptr
}

type IOrientationSensor3 struct {
	win32.IInspectable
}

func (this *IOrientationSensor3) Vtbl() *IOrientationSensor3Vtbl {
	return (*IOrientationSensor3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IOrientationSensor3) Put_ReportLatency(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReportLatency, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IOrientationSensor3) Get_ReportLatency() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportLatency, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IOrientationSensor3) Get_MaxBatchSize() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxBatchSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5A69B648-4C29-49EC-B28F-EA1D117B66F0
var IID_IOrientationSensorDeviceId = syscall.GUID{0x5A69B648, 0x4C29, 0x49EC,
	[8]byte{0xB2, 0x8F, 0xEA, 0x1D, 0x11, 0x7B, 0x66, 0xF0}}

type IOrientationSensorDeviceIdInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
}

type IOrientationSensorDeviceIdVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId uintptr
}

type IOrientationSensorDeviceId struct {
	win32.IInspectable
}

func (this *IOrientationSensorDeviceId) Vtbl() *IOrientationSensorDeviceIdVtbl {
	return (*IOrientationSensorDeviceIdVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IOrientationSensorDeviceId) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 4756C993-6595-4897-BCC6-D537EE757564
var IID_IOrientationSensorReading = syscall.GUID{0x4756C993, 0x6595, 0x4897,
	[8]byte{0xBC, 0xC6, 0xD5, 0x37, 0xEE, 0x75, 0x75, 0x64}}

type IOrientationSensorReadingInterface interface {
	win32.IInspectableInterface
	Get_Timestamp() DateTime
	Get_RotationMatrix() *ISensorRotationMatrix
	Get_Quaternion() *ISensorQuaternion
}

type IOrientationSensorReadingVtbl struct {
	win32.IInspectableVtbl
	Get_Timestamp      uintptr
	Get_RotationMatrix uintptr
	Get_Quaternion     uintptr
}

type IOrientationSensorReading struct {
	win32.IInspectable
}

func (this *IOrientationSensorReading) Vtbl() *IOrientationSensorReadingVtbl {
	return (*IOrientationSensorReadingVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IOrientationSensorReading) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IOrientationSensorReading) Get_RotationMatrix() *ISensorRotationMatrix {
	var _result *ISensorRotationMatrix
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RotationMatrix, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IOrientationSensorReading) Get_Quaternion() *ISensorQuaternion {
	var _result *ISensorQuaternion
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Quaternion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 00576E5F-49F8-4C05-9E07-24FAC79408C3
var IID_IOrientationSensorReading2 = syscall.GUID{0x00576E5F, 0x49F8, 0x4C05,
	[8]byte{0x9E, 0x07, 0x24, 0xFA, 0xC7, 0x94, 0x08, 0xC3}}

type IOrientationSensorReading2Interface interface {
	win32.IInspectableInterface
	Get_PerformanceCount() *IReference[TimeSpan]
	Get_Properties() *IMapView[string, interface{}]
}

type IOrientationSensorReading2Vtbl struct {
	win32.IInspectableVtbl
	Get_PerformanceCount uintptr
	Get_Properties       uintptr
}

type IOrientationSensorReading2 struct {
	win32.IInspectable
}

func (this *IOrientationSensorReading2) Vtbl() *IOrientationSensorReading2Vtbl {
	return (*IOrientationSensorReading2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IOrientationSensorReading2) Get_PerformanceCount() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PerformanceCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IOrientationSensorReading2) Get_Properties() *IMapView[string, interface{}] {
	var _result *IMapView[string, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 012C1186-C3BA-46BC-AE65-7A98996CBFB8
var IID_IOrientationSensorReadingChangedEventArgs = syscall.GUID{0x012C1186, 0xC3BA, 0x46BC,
	[8]byte{0xAE, 0x65, 0x7A, 0x98, 0x99, 0x6C, 0xBF, 0xB8}}

type IOrientationSensorReadingChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Reading() *IOrientationSensorReading
}

type IOrientationSensorReadingChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Reading uintptr
}

type IOrientationSensorReadingChangedEventArgs struct {
	win32.IInspectable
}

func (this *IOrientationSensorReadingChangedEventArgs) Vtbl() *IOrientationSensorReadingChangedEventArgsVtbl {
	return (*IOrientationSensorReadingChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IOrientationSensorReadingChangedEventArgs) Get_Reading() *IOrientationSensorReading {
	var _result *IOrientationSensorReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Reading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D1AC9824-3F5A-49A2-BC7B-1180BC38CD2B
var IID_IOrientationSensorReadingYawAccuracy = syscall.GUID{0xD1AC9824, 0x3F5A, 0x49A2,
	[8]byte{0xBC, 0x7B, 0x11, 0x80, 0xBC, 0x38, 0xCD, 0x2B}}

type IOrientationSensorReadingYawAccuracyInterface interface {
	win32.IInspectableInterface
	Get_YawAccuracy() MagnetometerAccuracy
}

type IOrientationSensorReadingYawAccuracyVtbl struct {
	win32.IInspectableVtbl
	Get_YawAccuracy uintptr
}

type IOrientationSensorReadingYawAccuracy struct {
	win32.IInspectable
}

func (this *IOrientationSensorReadingYawAccuracy) Vtbl() *IOrientationSensorReadingYawAccuracyVtbl {
	return (*IOrientationSensorReadingYawAccuracyVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IOrientationSensorReadingYawAccuracy) Get_YawAccuracy() MagnetometerAccuracy {
	var _result MagnetometerAccuracy
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_YawAccuracy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 10EF8712-FB4C-428A-898B-2765E409E669
var IID_IOrientationSensorStatics = syscall.GUID{0x10EF8712, 0xFB4C, 0x428A,
	[8]byte{0x89, 0x8B, 0x27, 0x65, 0xE4, 0x09, 0xE6, 0x69}}

type IOrientationSensorStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *IOrientationSensor
}

type IOrientationSensorStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
}

type IOrientationSensorStatics struct {
	win32.IInspectable
}

func (this *IOrientationSensorStatics) Vtbl() *IOrientationSensorStaticsVtbl {
	return (*IOrientationSensorStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IOrientationSensorStatics) GetDefault() *IOrientationSensor {
	var _result *IOrientationSensor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 59DA0D0B-D40A-4C71-9276-8A272A0A6619
var IID_IOrientationSensorStatics2 = syscall.GUID{0x59DA0D0B, 0xD40A, 0x4C71,
	[8]byte{0x92, 0x76, 0x8A, 0x27, 0x2A, 0x0A, 0x66, 0x19}}

type IOrientationSensorStatics2Interface interface {
	win32.IInspectableInterface
	GetDefaultForRelativeReadings() *IOrientationSensor
}

type IOrientationSensorStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetDefaultForRelativeReadings uintptr
}

type IOrientationSensorStatics2 struct {
	win32.IInspectable
}

func (this *IOrientationSensorStatics2) Vtbl() *IOrientationSensorStatics2Vtbl {
	return (*IOrientationSensorStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IOrientationSensorStatics2) GetDefaultForRelativeReadings() *IOrientationSensor {
	var _result *IOrientationSensor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefaultForRelativeReadings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D82CE920-2777-40FF-9F59-D654B085F12F
var IID_IOrientationSensorStatics3 = syscall.GUID{0xD82CE920, 0x2777, 0x40FF,
	[8]byte{0x9F, 0x59, 0xD6, 0x54, 0xB0, 0x85, 0xF1, 0x2F}}

type IOrientationSensorStatics3Interface interface {
	win32.IInspectableInterface
	GetDefaultWithSensorReadingType(sensorReadingtype SensorReadingType) *IOrientationSensor
	GetDefaultWithSensorReadingTypeAndSensorOptimizationGoal(sensorReadingType SensorReadingType, optimizationGoal SensorOptimizationGoal) *IOrientationSensor
}

type IOrientationSensorStatics3Vtbl struct {
	win32.IInspectableVtbl
	GetDefaultWithSensorReadingType                          uintptr
	GetDefaultWithSensorReadingTypeAndSensorOptimizationGoal uintptr
}

type IOrientationSensorStatics3 struct {
	win32.IInspectable
}

func (this *IOrientationSensorStatics3) Vtbl() *IOrientationSensorStatics3Vtbl {
	return (*IOrientationSensorStatics3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IOrientationSensorStatics3) GetDefaultWithSensorReadingType(sensorReadingtype SensorReadingType) *IOrientationSensor {
	var _result *IOrientationSensor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefaultWithSensorReadingType, uintptr(unsafe.Pointer(this)), uintptr(sensorReadingtype), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IOrientationSensorStatics3) GetDefaultWithSensorReadingTypeAndSensorOptimizationGoal(sensorReadingType SensorReadingType, optimizationGoal SensorOptimizationGoal) *IOrientationSensor {
	var _result *IOrientationSensor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefaultWithSensorReadingTypeAndSensorOptimizationGoal, uintptr(unsafe.Pointer(this)), uintptr(sensorReadingType), uintptr(optimizationGoal), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A67FEB55-2C85-4B28-A0FE-58C4B20495F5
var IID_IOrientationSensorStatics4 = syscall.GUID{0xA67FEB55, 0x2C85, 0x4B28,
	[8]byte{0xA0, 0xFE, 0x58, 0xC4, 0xB2, 0x04, 0x95, 0xF5}}

type IOrientationSensorStatics4Interface interface {
	win32.IInspectableInterface
	GetDeviceSelector(readingType SensorReadingType) string
	GetDeviceSelectorWithSensorReadingTypeAndSensorOptimizationGoal(readingType SensorReadingType, optimizationGoal SensorOptimizationGoal) string
	FromIdAsync(deviceId string) *IAsyncOperation[*IOrientationSensor]
}

type IOrientationSensorStatics4Vtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelector                                               uintptr
	GetDeviceSelectorWithSensorReadingTypeAndSensorOptimizationGoal uintptr
	FromIdAsync                                                     uintptr
}

type IOrientationSensorStatics4 struct {
	win32.IInspectable
}

func (this *IOrientationSensorStatics4) Vtbl() *IOrientationSensorStatics4Vtbl {
	return (*IOrientationSensorStatics4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IOrientationSensorStatics4) GetDeviceSelector(readingType SensorReadingType) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(readingType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IOrientationSensorStatics4) GetDeviceSelectorWithSensorReadingTypeAndSensorOptimizationGoal(readingType SensorReadingType, optimizationGoal SensorOptimizationGoal) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelectorWithSensorReadingTypeAndSensorOptimizationGoal, uintptr(unsafe.Pointer(this)), uintptr(readingType), uintptr(optimizationGoal), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IOrientationSensorStatics4) FromIdAsync(deviceId string) *IAsyncOperation[*IOrientationSensor] {
	var _result *IAsyncOperation[*IOrientationSensor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9A1E013D-3D98-45F8-8920-8E4ECACA5F97
var IID_IPedometer = syscall.GUID{0x9A1E013D, 0x3D98, 0x45F8,
	[8]byte{0x89, 0x20, 0x8E, 0x4E, 0xCA, 0xCA, 0x5F, 0x97}}

type IPedometerInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_PowerInMilliwatts() float64
	Get_MinimumReportInterval() uint32
	Put_ReportInterval(value uint32)
	Get_ReportInterval() uint32
	Add_ReadingChanged(handler TypedEventHandler[*IPedometer, *IPedometerReadingChangedEventArgs]) EventRegistrationToken
	Remove_ReadingChanged(token EventRegistrationToken)
}

type IPedometerVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId              uintptr
	Get_PowerInMilliwatts     uintptr
	Get_MinimumReportInterval uintptr
	Put_ReportInterval        uintptr
	Get_ReportInterval        uintptr
	Add_ReadingChanged        uintptr
	Remove_ReadingChanged     uintptr
}

type IPedometer struct {
	win32.IInspectable
}

func (this *IPedometer) Vtbl() *IPedometerVtbl {
	return (*IPedometerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPedometer) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPedometer) Get_PowerInMilliwatts() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PowerInMilliwatts, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPedometer) Get_MinimumReportInterval() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinimumReportInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPedometer) Put_ReportInterval(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReportInterval, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPedometer) Get_ReportInterval() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPedometer) Add_ReadingChanged(handler TypedEventHandler[*IPedometer, *IPedometerReadingChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ReadingChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPedometer) Remove_ReadingChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ReadingChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// E5A406DF-2B81-4ADD-B2FF-77AB6C98BA19
var IID_IPedometer2 = syscall.GUID{0xE5A406DF, 0x2B81, 0x4ADD,
	[8]byte{0xB2, 0xFF, 0x77, 0xAB, 0x6C, 0x98, 0xBA, 0x19}}

type IPedometer2Interface interface {
	win32.IInspectableInterface
	GetCurrentReadings() *IMapView[PedometerStepKind, *IPedometerReading]
}

type IPedometer2Vtbl struct {
	win32.IInspectableVtbl
	GetCurrentReadings uintptr
}

type IPedometer2 struct {
	win32.IInspectable
}

func (this *IPedometer2) Vtbl() *IPedometer2Vtbl {
	return (*IPedometer2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPedometer2) GetCurrentReadings() *IMapView[PedometerStepKind, *IPedometerReading] {
	var _result *IMapView[PedometerStepKind, *IPedometerReading]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentReadings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CBAD8F50-7A54-466B-9010-77A162FCA5D7
var IID_IPedometerDataThresholdFactory = syscall.GUID{0xCBAD8F50, 0x7A54, 0x466B,
	[8]byte{0x90, 0x10, 0x77, 0xA1, 0x62, 0xFC, 0xA5, 0xD7}}

type IPedometerDataThresholdFactoryInterface interface {
	win32.IInspectableInterface
	Create(sensor *IPedometer, stepGoal int32) *ISensorDataThreshold
}

type IPedometerDataThresholdFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IPedometerDataThresholdFactory struct {
	win32.IInspectable
}

func (this *IPedometerDataThresholdFactory) Vtbl() *IPedometerDataThresholdFactoryVtbl {
	return (*IPedometerDataThresholdFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPedometerDataThresholdFactory) Create(sensor *IPedometer, stepGoal int32) *ISensorDataThreshold {
	var _result *ISensorDataThreshold
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sensor)), uintptr(stepGoal), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2245DCF4-A8E1-432F-896A-BE0DD9B02D24
var IID_IPedometerReading = syscall.GUID{0x2245DCF4, 0xA8E1, 0x432F,
	[8]byte{0x89, 0x6A, 0xBE, 0x0D, 0xD9, 0xB0, 0x2D, 0x24}}

type IPedometerReadingInterface interface {
	win32.IInspectableInterface
	Get_StepKind() PedometerStepKind
	Get_CumulativeSteps() int32
	Get_Timestamp() DateTime
	Get_CumulativeStepsDuration() TimeSpan
}

type IPedometerReadingVtbl struct {
	win32.IInspectableVtbl
	Get_StepKind                uintptr
	Get_CumulativeSteps         uintptr
	Get_Timestamp               uintptr
	Get_CumulativeStepsDuration uintptr
}

type IPedometerReading struct {
	win32.IInspectable
}

func (this *IPedometerReading) Vtbl() *IPedometerReadingVtbl {
	return (*IPedometerReadingVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPedometerReading) Get_StepKind() PedometerStepKind {
	var _result PedometerStepKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StepKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPedometerReading) Get_CumulativeSteps() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CumulativeSteps, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPedometerReading) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPedometerReading) Get_CumulativeStepsDuration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CumulativeStepsDuration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F855E47E-ABBC-4456-86A8-25CF2B333742
var IID_IPedometerReadingChangedEventArgs = syscall.GUID{0xF855E47E, 0xABBC, 0x4456,
	[8]byte{0x86, 0xA8, 0x25, 0xCF, 0x2B, 0x33, 0x37, 0x42}}

type IPedometerReadingChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Reading() *IPedometerReading
}

type IPedometerReadingChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Reading uintptr
}

type IPedometerReadingChangedEventArgs struct {
	win32.IInspectable
}

func (this *IPedometerReadingChangedEventArgs) Vtbl() *IPedometerReadingChangedEventArgsVtbl {
	return (*IPedometerReadingChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPedometerReadingChangedEventArgs) Get_Reading() *IPedometerReading {
	var _result *IPedometerReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Reading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 82980A2F-4083-4DFB-B411-938EA0F4B946
var IID_IPedometerStatics = syscall.GUID{0x82980A2F, 0x4083, 0x4DFB,
	[8]byte{0xB4, 0x11, 0x93, 0x8E, 0xA0, 0xF4, 0xB9, 0x46}}

type IPedometerStaticsInterface interface {
	win32.IInspectableInterface
	FromIdAsync(deviceId string) *IAsyncOperation[*IPedometer]
	GetDefaultAsync() *IAsyncOperation[*IPedometer]
	GetDeviceSelector() string
	GetSystemHistoryAsync(fromTime DateTime) *IAsyncOperation[*IVectorView[*IPedometerReading]]
	GetSystemHistoryWithDurationAsync(fromTime DateTime, duration TimeSpan) *IAsyncOperation[*IVectorView[*IPedometerReading]]
}

type IPedometerStaticsVtbl struct {
	win32.IInspectableVtbl
	FromIdAsync                       uintptr
	GetDefaultAsync                   uintptr
	GetDeviceSelector                 uintptr
	GetSystemHistoryAsync             uintptr
	GetSystemHistoryWithDurationAsync uintptr
}

type IPedometerStatics struct {
	win32.IInspectable
}

func (this *IPedometerStatics) Vtbl() *IPedometerStaticsVtbl {
	return (*IPedometerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPedometerStatics) FromIdAsync(deviceId string) *IAsyncOperation[*IPedometer] {
	var _result *IAsyncOperation[*IPedometer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPedometerStatics) GetDefaultAsync() *IAsyncOperation[*IPedometer] {
	var _result *IAsyncOperation[*IPedometer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefaultAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPedometerStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPedometerStatics) GetSystemHistoryAsync(fromTime DateTime) *IAsyncOperation[*IVectorView[*IPedometerReading]] {
	var _result *IAsyncOperation[*IVectorView[*IPedometerReading]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSystemHistoryAsync, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&fromTime)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPedometerStatics) GetSystemHistoryWithDurationAsync(fromTime DateTime, duration TimeSpan) *IAsyncOperation[*IVectorView[*IPedometerReading]] {
	var _result *IAsyncOperation[*IVectorView[*IPedometerReading]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSystemHistoryWithDurationAsync, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&fromTime)), *(*uintptr)(unsafe.Pointer(&duration)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 79F5C6BB-CE0E-4133-B47E-8627EA72F677
var IID_IPedometerStatics2 = syscall.GUID{0x79F5C6BB, 0xCE0E, 0x4133,
	[8]byte{0xB4, 0x7E, 0x86, 0x27, 0xEA, 0x72, 0xF6, 0x77}}

type IPedometerStatics2Interface interface {
	win32.IInspectableInterface
	GetReadingsFromTriggerDetails(triggerDetails *ISensorDataThresholdTriggerDetails) *IVectorView[*IPedometerReading]
}

type IPedometerStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetReadingsFromTriggerDetails uintptr
}

type IPedometerStatics2 struct {
	win32.IInspectable
}

func (this *IPedometerStatics2) Vtbl() *IPedometerStatics2Vtbl {
	return (*IPedometerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPedometerStatics2) GetReadingsFromTriggerDetails(triggerDetails *ISensorDataThresholdTriggerDetails) *IVectorView[*IPedometerReading] {
	var _result *IVectorView[*IPedometerReading]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetReadingsFromTriggerDetails, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(triggerDetails)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 54C076B8-ECFB-4944-B928-74FC504D47EE
var IID_IProximitySensor = syscall.GUID{0x54C076B8, 0xECFB, 0x4944,
	[8]byte{0xB9, 0x28, 0x74, 0xFC, 0x50, 0x4D, 0x47, 0xEE}}

type IProximitySensorInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_MaxDistanceInMillimeters() *IReference[uint32]
	Get_MinDistanceInMillimeters() *IReference[uint32]
	GetCurrentReading() *IProximitySensorReading
	Add_ReadingChanged(handler TypedEventHandler[*IProximitySensor, *IProximitySensorReadingChangedEventArgs]) EventRegistrationToken
	Remove_ReadingChanged(token EventRegistrationToken)
	CreateDisplayOnOffController() *IClosable
}

type IProximitySensorVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId                 uintptr
	Get_MaxDistanceInMillimeters uintptr
	Get_MinDistanceInMillimeters uintptr
	GetCurrentReading            uintptr
	Add_ReadingChanged           uintptr
	Remove_ReadingChanged        uintptr
	CreateDisplayOnOffController uintptr
}

type IProximitySensor struct {
	win32.IInspectable
}

func (this *IProximitySensor) Vtbl() *IProximitySensorVtbl {
	return (*IProximitySensorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProximitySensor) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IProximitySensor) Get_MaxDistanceInMillimeters() *IReference[uint32] {
	var _result *IReference[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxDistanceInMillimeters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProximitySensor) Get_MinDistanceInMillimeters() *IReference[uint32] {
	var _result *IReference[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinDistanceInMillimeters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProximitySensor) GetCurrentReading() *IProximitySensorReading {
	var _result *IProximitySensorReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentReading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProximitySensor) Add_ReadingChanged(handler TypedEventHandler[*IProximitySensor, *IProximitySensorReadingChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ReadingChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProximitySensor) Remove_ReadingChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ReadingChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IProximitySensor) CreateDisplayOnOffController() *IClosable {
	var _result *IClosable
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDisplayOnOffController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 905AC121-6D27-4AD3-9DB5-6467F2A5AD9D
var IID_IProximitySensorDataThresholdFactory = syscall.GUID{0x905AC121, 0x6D27, 0x4AD3,
	[8]byte{0x9D, 0xB5, 0x64, 0x67, 0xF2, 0xA5, 0xAD, 0x9D}}

type IProximitySensorDataThresholdFactoryInterface interface {
	win32.IInspectableInterface
	Create(sensor *IProximitySensor) *ISensorDataThreshold
}

type IProximitySensorDataThresholdFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IProximitySensorDataThresholdFactory struct {
	win32.IInspectable
}

func (this *IProximitySensorDataThresholdFactory) Vtbl() *IProximitySensorDataThresholdFactoryVtbl {
	return (*IProximitySensorDataThresholdFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProximitySensorDataThresholdFactory) Create(sensor *IProximitySensor) *ISensorDataThreshold {
	var _result *ISensorDataThreshold
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sensor)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 71228D59-132D-4D5F-8FF9-2F0DB8751CED
var IID_IProximitySensorReading = syscall.GUID{0x71228D59, 0x132D, 0x4D5F,
	[8]byte{0x8F, 0xF9, 0x2F, 0x0D, 0xB8, 0x75, 0x1C, 0xED}}

type IProximitySensorReadingInterface interface {
	win32.IInspectableInterface
	Get_Timestamp() DateTime
	Get_IsDetected() bool
	Get_DistanceInMillimeters() *IReference[uint32]
}

type IProximitySensorReadingVtbl struct {
	win32.IInspectableVtbl
	Get_Timestamp             uintptr
	Get_IsDetected            uintptr
	Get_DistanceInMillimeters uintptr
}

type IProximitySensorReading struct {
	win32.IInspectable
}

func (this *IProximitySensorReading) Vtbl() *IProximitySensorReadingVtbl {
	return (*IProximitySensorReadingVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProximitySensorReading) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProximitySensorReading) Get_IsDetected() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDetected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProximitySensorReading) Get_DistanceInMillimeters() *IReference[uint32] {
	var _result *IReference[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DistanceInMillimeters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CFC2F366-C3E8-40FD-8CC3-67E289004938
var IID_IProximitySensorReadingChangedEventArgs = syscall.GUID{0xCFC2F366, 0xC3E8, 0x40FD,
	[8]byte{0x8C, 0xC3, 0x67, 0xE2, 0x89, 0x00, 0x49, 0x38}}

type IProximitySensorReadingChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Reading() *IProximitySensorReading
}

type IProximitySensorReadingChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Reading uintptr
}

type IProximitySensorReadingChangedEventArgs struct {
	win32.IInspectable
}

func (this *IProximitySensorReadingChangedEventArgs) Vtbl() *IProximitySensorReadingChangedEventArgsVtbl {
	return (*IProximitySensorReadingChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProximitySensorReadingChangedEventArgs) Get_Reading() *IProximitySensorReading {
	var _result *IProximitySensorReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Reading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 29186649-6269-4E57-A5AD-82BE80813392
var IID_IProximitySensorStatics = syscall.GUID{0x29186649, 0x6269, 0x4E57,
	[8]byte{0xA5, 0xAD, 0x82, 0xBE, 0x80, 0x81, 0x33, 0x92}}

type IProximitySensorStaticsInterface interface {
	win32.IInspectableInterface
	GetDeviceSelector() string
	FromId(sensorId string) *IProximitySensor
}

type IProximitySensorStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelector uintptr
	FromId            uintptr
}

type IProximitySensorStatics struct {
	win32.IInspectable
}

func (this *IProximitySensorStatics) Vtbl() *IProximitySensorStaticsVtbl {
	return (*IProximitySensorStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProximitySensorStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IProximitySensorStatics) FromId(sensorId string) *IProximitySensor {
	var _result *IProximitySensor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromId, uintptr(unsafe.Pointer(this)), NewHStr(sensorId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CBF473AE-E9CA-422F-AD67-4C3D25DF350C
var IID_IProximitySensorStatics2 = syscall.GUID{0xCBF473AE, 0xE9CA, 0x422F,
	[8]byte{0xAD, 0x67, 0x4C, 0x3D, 0x25, 0xDF, 0x35, 0x0C}}

type IProximitySensorStatics2Interface interface {
	win32.IInspectableInterface
	GetReadingsFromTriggerDetails(triggerDetails *ISensorDataThresholdTriggerDetails) *IVectorView[*IProximitySensorReading]
}

type IProximitySensorStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetReadingsFromTriggerDetails uintptr
}

type IProximitySensorStatics2 struct {
	win32.IInspectable
}

func (this *IProximitySensorStatics2) Vtbl() *IProximitySensorStatics2Vtbl {
	return (*IProximitySensorStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProximitySensorStatics2) GetReadingsFromTriggerDetails(triggerDetails *ISensorDataThresholdTriggerDetails) *IVectorView[*IProximitySensorReading] {
	var _result *IVectorView[*IProximitySensorReading]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetReadingsFromTriggerDetails, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(triggerDetails)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 54DAEC61-FE4B-4E07-B260-3A4CDFBE396E
var IID_ISensorDataThreshold = syscall.GUID{0x54DAEC61, 0xFE4B, 0x4E07,
	[8]byte{0xB2, 0x60, 0x3A, 0x4C, 0xDF, 0xBE, 0x39, 0x6E}}

type ISensorDataThresholdInterface interface {
	win32.IInspectableInterface
}

type ISensorDataThresholdVtbl struct {
	win32.IInspectableVtbl
}

type ISensorDataThreshold struct {
	win32.IInspectable
}

func (this *ISensorDataThreshold) Vtbl() *ISensorDataThresholdVtbl {
	return (*ISensorDataThresholdVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 9106F1B7-E88D-48B1-BC90-619C7B349391
var IID_ISensorDataThresholdTriggerDetails = syscall.GUID{0x9106F1B7, 0xE88D, 0x48B1,
	[8]byte{0xBC, 0x90, 0x61, 0x9C, 0x7B, 0x34, 0x93, 0x91}}

type ISensorDataThresholdTriggerDetailsInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_SensorType() SensorType
}

type ISensorDataThresholdTriggerDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId   uintptr
	Get_SensorType uintptr
}

type ISensorDataThresholdTriggerDetails struct {
	win32.IInspectable
}

func (this *ISensorDataThresholdTriggerDetails) Vtbl() *ISensorDataThresholdTriggerDetailsVtbl {
	return (*ISensorDataThresholdTriggerDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISensorDataThresholdTriggerDetails) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISensorDataThresholdTriggerDetails) Get_SensorType() SensorType {
	var _result SensorType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SensorType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C9C5C827-C71C-46E7-9DA3-36A193B232BC
var IID_ISensorQuaternion = syscall.GUID{0xC9C5C827, 0xC71C, 0x46E7,
	[8]byte{0x9D, 0xA3, 0x36, 0xA1, 0x93, 0xB2, 0x32, 0xBC}}

type ISensorQuaternionInterface interface {
	win32.IInspectableInterface
	Get_W() float32
	Get_X() float32
	Get_Y() float32
	Get_Z() float32
}

type ISensorQuaternionVtbl struct {
	win32.IInspectableVtbl
	Get_W uintptr
	Get_X uintptr
	Get_Y uintptr
	Get_Z uintptr
}

type ISensorQuaternion struct {
	win32.IInspectable
}

func (this *ISensorQuaternion) Vtbl() *ISensorQuaternionVtbl {
	return (*ISensorQuaternionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISensorQuaternion) Get_W() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_W, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISensorQuaternion) Get_X() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_X, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISensorQuaternion) Get_Y() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Y, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISensorQuaternion) Get_Z() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Z, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 0A3D5A67-22F4-4392-9538-65D0BD064AA6
var IID_ISensorRotationMatrix = syscall.GUID{0x0A3D5A67, 0x22F4, 0x4392,
	[8]byte{0x95, 0x38, 0x65, 0xD0, 0xBD, 0x06, 0x4A, 0xA6}}

type ISensorRotationMatrixInterface interface {
	win32.IInspectableInterface
	Get_M11() float32
	Get_M12() float32
	Get_M13() float32
	Get_M21() float32
	Get_M22() float32
	Get_M23() float32
	Get_M31() float32
	Get_M32() float32
	Get_M33() float32
}

type ISensorRotationMatrixVtbl struct {
	win32.IInspectableVtbl
	Get_M11 uintptr
	Get_M12 uintptr
	Get_M13 uintptr
	Get_M21 uintptr
	Get_M22 uintptr
	Get_M23 uintptr
	Get_M31 uintptr
	Get_M32 uintptr
	Get_M33 uintptr
}

type ISensorRotationMatrix struct {
	win32.IInspectable
}

func (this *ISensorRotationMatrix) Vtbl() *ISensorRotationMatrixVtbl {
	return (*ISensorRotationMatrixVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISensorRotationMatrix) Get_M11() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_M11, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISensorRotationMatrix) Get_M12() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_M12, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISensorRotationMatrix) Get_M13() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_M13, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISensorRotationMatrix) Get_M21() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_M21, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISensorRotationMatrix) Get_M22() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_M22, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISensorRotationMatrix) Get_M23() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_M23, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISensorRotationMatrix) Get_M31() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_M31, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISensorRotationMatrix) Get_M32() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_M32, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISensorRotationMatrix) Get_M33() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_M33, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5FF53856-214A-4DEE-A3F9-616F1AB06FFD
var IID_ISimpleOrientationSensor = syscall.GUID{0x5FF53856, 0x214A, 0x4DEE,
	[8]byte{0xA3, 0xF9, 0x61, 0x6F, 0x1A, 0xB0, 0x6F, 0xFD}}

type ISimpleOrientationSensorInterface interface {
	win32.IInspectableInterface
	GetCurrentOrientation() SimpleOrientation
	Add_OrientationChanged(handler TypedEventHandler[*ISimpleOrientationSensor, *ISimpleOrientationSensorOrientationChangedEventArgs]) EventRegistrationToken
	Remove_OrientationChanged(token EventRegistrationToken)
}

type ISimpleOrientationSensorVtbl struct {
	win32.IInspectableVtbl
	GetCurrentOrientation     uintptr
	Add_OrientationChanged    uintptr
	Remove_OrientationChanged uintptr
}

type ISimpleOrientationSensor struct {
	win32.IInspectable
}

func (this *ISimpleOrientationSensor) Vtbl() *ISimpleOrientationSensorVtbl {
	return (*ISimpleOrientationSensorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISimpleOrientationSensor) GetCurrentOrientation() SimpleOrientation {
	var _result SimpleOrientation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentOrientation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISimpleOrientationSensor) Add_OrientationChanged(handler TypedEventHandler[*ISimpleOrientationSensor, *ISimpleOrientationSensorOrientationChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_OrientationChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISimpleOrientationSensor) Remove_OrientationChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_OrientationChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// A277A798-8870-453E-8BD6-B8F5D8D7941B
var IID_ISimpleOrientationSensor2 = syscall.GUID{0xA277A798, 0x8870, 0x453E,
	[8]byte{0x8B, 0xD6, 0xB8, 0xF5, 0xD8, 0xD7, 0x94, 0x1B}}

type ISimpleOrientationSensor2Interface interface {
	win32.IInspectableInterface
	Put_ReadingTransform(value unsafe.Pointer)
	Get_ReadingTransform() unsafe.Pointer
}

type ISimpleOrientationSensor2Vtbl struct {
	win32.IInspectableVtbl
	Put_ReadingTransform uintptr
	Get_ReadingTransform uintptr
}

type ISimpleOrientationSensor2 struct {
	win32.IInspectable
}

func (this *ISimpleOrientationSensor2) Vtbl() *ISimpleOrientationSensor2Vtbl {
	return (*ISimpleOrientationSensor2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISimpleOrientationSensor2) Put_ReadingTransform(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReadingTransform, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISimpleOrientationSensor2) Get_ReadingTransform() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReadingTransform, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// FBC00ACB-3B76-41F6-8091-30EFE646D3CF
var IID_ISimpleOrientationSensorDeviceId = syscall.GUID{0xFBC00ACB, 0x3B76, 0x41F6,
	[8]byte{0x80, 0x91, 0x30, 0xEF, 0xE6, 0x46, 0xD3, 0xCF}}

type ISimpleOrientationSensorDeviceIdInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
}

type ISimpleOrientationSensorDeviceIdVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId uintptr
}

type ISimpleOrientationSensorDeviceId struct {
	win32.IInspectable
}

func (this *ISimpleOrientationSensorDeviceId) Vtbl() *ISimpleOrientationSensorDeviceIdVtbl {
	return (*ISimpleOrientationSensorDeviceIdVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISimpleOrientationSensorDeviceId) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// BCD5C660-23D4-4B4C-A22E-BA81ADE0C601
var IID_ISimpleOrientationSensorOrientationChangedEventArgs = syscall.GUID{0xBCD5C660, 0x23D4, 0x4B4C,
	[8]byte{0xA2, 0x2E, 0xBA, 0x81, 0xAD, 0xE0, 0xC6, 0x01}}

type ISimpleOrientationSensorOrientationChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Timestamp() DateTime
	Get_Orientation() SimpleOrientation
}

type ISimpleOrientationSensorOrientationChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Timestamp   uintptr
	Get_Orientation uintptr
}

type ISimpleOrientationSensorOrientationChangedEventArgs struct {
	win32.IInspectable
}

func (this *ISimpleOrientationSensorOrientationChangedEventArgs) Vtbl() *ISimpleOrientationSensorOrientationChangedEventArgsVtbl {
	return (*ISimpleOrientationSensorOrientationChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISimpleOrientationSensorOrientationChangedEventArgs) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISimpleOrientationSensorOrientationChangedEventArgs) Get_Orientation() SimpleOrientation {
	var _result SimpleOrientation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Orientation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 72ED066F-70AA-40C6-9B1B-3433F7459B4E
var IID_ISimpleOrientationSensorStatics = syscall.GUID{0x72ED066F, 0x70AA, 0x40C6,
	[8]byte{0x9B, 0x1B, 0x34, 0x33, 0xF7, 0x45, 0x9B, 0x4E}}

type ISimpleOrientationSensorStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *ISimpleOrientationSensor
}

type ISimpleOrientationSensorStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
}

type ISimpleOrientationSensorStatics struct {
	win32.IInspectable
}

func (this *ISimpleOrientationSensorStatics) Vtbl() *ISimpleOrientationSensorStaticsVtbl {
	return (*ISimpleOrientationSensorStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISimpleOrientationSensorStatics) GetDefault() *ISimpleOrientationSensor {
	var _result *ISimpleOrientationSensor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 848F9C7F-B138-4E11-8910-A2A2A3B56D83
var IID_ISimpleOrientationSensorStatics2 = syscall.GUID{0x848F9C7F, 0xB138, 0x4E11,
	[8]byte{0x89, 0x10, 0xA2, 0xA2, 0xA3, 0xB5, 0x6D, 0x83}}

type ISimpleOrientationSensorStatics2Interface interface {
	win32.IInspectableInterface
	GetDeviceSelector() string
	FromIdAsync(deviceId string) *IAsyncOperation[*ISimpleOrientationSensor]
}

type ISimpleOrientationSensorStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelector uintptr
	FromIdAsync       uintptr
}

type ISimpleOrientationSensorStatics2 struct {
	win32.IInspectable
}

func (this *ISimpleOrientationSensorStatics2) Vtbl() *ISimpleOrientationSensorStatics2Vtbl {
	return (*ISimpleOrientationSensorStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISimpleOrientationSensorStatics2) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISimpleOrientationSensorStatics2) FromIdAsync(deviceId string) *IAsyncOperation[*ISimpleOrientationSensor] {
	var _result *IAsyncOperation[*ISimpleOrientationSensor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type Accelerometer struct {
	RtClass
	*IAccelerometer
}

func NewIAccelerometerStatics3() *IAccelerometerStatics3 {
	var p *IAccelerometerStatics3
	hs := NewHStr("Windows.Devices.Sensors.Accelerometer")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAccelerometerStatics3, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIAccelerometerStatics2() *IAccelerometerStatics2 {
	var p *IAccelerometerStatics2
	hs := NewHStr("Windows.Devices.Sensors.Accelerometer")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAccelerometerStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIAccelerometerStatics() *IAccelerometerStatics {
	var p *IAccelerometerStatics
	hs := NewHStr("Windows.Devices.Sensors.Accelerometer")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAccelerometerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type AccelerometerReading struct {
	RtClass
	*IAccelerometerReading
}

type AccelerometerReadingChangedEventArgs struct {
	RtClass
	*IAccelerometerReadingChangedEventArgs
}

type AccelerometerShakenEventArgs struct {
	RtClass
	*IAccelerometerShakenEventArgs
}

type ActivitySensor struct {
	RtClass
	*IActivitySensor
}

func NewIActivitySensorStatics() *IActivitySensorStatics {
	var p *IActivitySensorStatics
	hs := NewHStr("Windows.Devices.Sensors.ActivitySensor")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IActivitySensorStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type ActivitySensorReading struct {
	RtClass
	*IActivitySensorReading
}

type ActivitySensorReadingChangeReport struct {
	RtClass
	*IActivitySensorReadingChangeReport
}

type ActivitySensorReadingChangedEventArgs struct {
	RtClass
	*IActivitySensorReadingChangedEventArgs
}

type Altimeter struct {
	RtClass
	*IAltimeter
}

func NewIAltimeterStatics() *IAltimeterStatics {
	var p *IAltimeterStatics
	hs := NewHStr("Windows.Devices.Sensors.Altimeter")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAltimeterStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type AltimeterReading struct {
	RtClass
	*IAltimeterReading
}

type AltimeterReadingChangedEventArgs struct {
	RtClass
	*IAltimeterReadingChangedEventArgs
}

type Barometer struct {
	RtClass
	*IBarometer
}

func NewIBarometerStatics() *IBarometerStatics {
	var p *IBarometerStatics
	hs := NewHStr("Windows.Devices.Sensors.Barometer")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IBarometerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIBarometerStatics2() *IBarometerStatics2 {
	var p *IBarometerStatics2
	hs := NewHStr("Windows.Devices.Sensors.Barometer")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IBarometerStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type BarometerReading struct {
	RtClass
	*IBarometerReading
}

type BarometerReadingChangedEventArgs struct {
	RtClass
	*IBarometerReadingChangedEventArgs
}

type Compass struct {
	RtClass
	*ICompass
}

func NewICompassStatics2() *ICompassStatics2 {
	var p *ICompassStatics2
	hs := NewHStr("Windows.Devices.Sensors.Compass")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICompassStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewICompassStatics() *ICompassStatics {
	var p *ICompassStatics
	hs := NewHStr("Windows.Devices.Sensors.Compass")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICompassStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type CompassReading struct {
	RtClass
	*ICompassReading
}

type CompassReadingChangedEventArgs struct {
	RtClass
	*ICompassReadingChangedEventArgs
}

type Gyrometer struct {
	RtClass
	*IGyrometer
}

func NewIGyrometerStatics() *IGyrometerStatics {
	var p *IGyrometerStatics
	hs := NewHStr("Windows.Devices.Sensors.Gyrometer")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGyrometerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIGyrometerStatics2() *IGyrometerStatics2 {
	var p *IGyrometerStatics2
	hs := NewHStr("Windows.Devices.Sensors.Gyrometer")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGyrometerStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type GyrometerReading struct {
	RtClass
	*IGyrometerReading
}

type GyrometerReadingChangedEventArgs struct {
	RtClass
	*IGyrometerReadingChangedEventArgs
}

type Inclinometer struct {
	RtClass
	*IInclinometer
}

func NewIInclinometerStatics() *IInclinometerStatics {
	var p *IInclinometerStatics
	hs := NewHStr("Windows.Devices.Sensors.Inclinometer")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IInclinometerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIInclinometerStatics4() *IInclinometerStatics4 {
	var p *IInclinometerStatics4
	hs := NewHStr("Windows.Devices.Sensors.Inclinometer")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IInclinometerStatics4, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIInclinometerStatics3() *IInclinometerStatics3 {
	var p *IInclinometerStatics3
	hs := NewHStr("Windows.Devices.Sensors.Inclinometer")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IInclinometerStatics3, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIInclinometerStatics2() *IInclinometerStatics2 {
	var p *IInclinometerStatics2
	hs := NewHStr("Windows.Devices.Sensors.Inclinometer")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IInclinometerStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type InclinometerReading struct {
	RtClass
	*IInclinometerReading
}

type InclinometerReadingChangedEventArgs struct {
	RtClass
	*IInclinometerReadingChangedEventArgs
}

type LightSensor struct {
	RtClass
	*ILightSensor
}

func NewILightSensorStatics2() *ILightSensorStatics2 {
	var p *ILightSensorStatics2
	hs := NewHStr("Windows.Devices.Sensors.LightSensor")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ILightSensorStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewILightSensorStatics() *ILightSensorStatics {
	var p *ILightSensorStatics
	hs := NewHStr("Windows.Devices.Sensors.LightSensor")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ILightSensorStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type LightSensorReading struct {
	RtClass
	*ILightSensorReading
}

type LightSensorReadingChangedEventArgs struct {
	RtClass
	*ILightSensorReadingChangedEventArgs
}

type Magnetometer struct {
	RtClass
	*IMagnetometer
}

func NewIMagnetometerStatics() *IMagnetometerStatics {
	var p *IMagnetometerStatics
	hs := NewHStr("Windows.Devices.Sensors.Magnetometer")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMagnetometerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIMagnetometerStatics2() *IMagnetometerStatics2 {
	var p *IMagnetometerStatics2
	hs := NewHStr("Windows.Devices.Sensors.Magnetometer")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMagnetometerStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type MagnetometerReading struct {
	RtClass
	*IMagnetometerReading
}

type MagnetometerReadingChangedEventArgs struct {
	RtClass
	*IMagnetometerReadingChangedEventArgs
}

type OrientationSensor struct {
	RtClass
	*IOrientationSensor
}

func NewIOrientationSensorStatics3() *IOrientationSensorStatics3 {
	var p *IOrientationSensorStatics3
	hs := NewHStr("Windows.Devices.Sensors.OrientationSensor")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IOrientationSensorStatics3, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIOrientationSensorStatics() *IOrientationSensorStatics {
	var p *IOrientationSensorStatics
	hs := NewHStr("Windows.Devices.Sensors.OrientationSensor")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IOrientationSensorStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIOrientationSensorStatics4() *IOrientationSensorStatics4 {
	var p *IOrientationSensorStatics4
	hs := NewHStr("Windows.Devices.Sensors.OrientationSensor")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IOrientationSensorStatics4, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIOrientationSensorStatics2() *IOrientationSensorStatics2 {
	var p *IOrientationSensorStatics2
	hs := NewHStr("Windows.Devices.Sensors.OrientationSensor")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IOrientationSensorStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type OrientationSensorReading struct {
	RtClass
	*IOrientationSensorReading
}

type OrientationSensorReadingChangedEventArgs struct {
	RtClass
	*IOrientationSensorReadingChangedEventArgs
}

type Pedometer struct {
	RtClass
	*IPedometer
}

func NewIPedometerStatics2() *IPedometerStatics2 {
	var p *IPedometerStatics2
	hs := NewHStr("Windows.Devices.Sensors.Pedometer")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPedometerStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIPedometerStatics() *IPedometerStatics {
	var p *IPedometerStatics
	hs := NewHStr("Windows.Devices.Sensors.Pedometer")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPedometerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type PedometerReading struct {
	RtClass
	*IPedometerReading
}

type PedometerReadingChangedEventArgs struct {
	RtClass
	*IPedometerReadingChangedEventArgs
}

type ProximitySensor struct {
	RtClass
	*IProximitySensor
}

func NewIProximitySensorStatics2() *IProximitySensorStatics2 {
	var p *IProximitySensorStatics2
	hs := NewHStr("Windows.Devices.Sensors.ProximitySensor")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IProximitySensorStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIProximitySensorStatics() *IProximitySensorStatics {
	var p *IProximitySensorStatics
	hs := NewHStr("Windows.Devices.Sensors.ProximitySensor")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IProximitySensorStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type ProximitySensorDisplayOnOffController struct {
	RtClass
	*IClosable
}

type ProximitySensorReading struct {
	RtClass
	*IProximitySensorReading
}

type ProximitySensorReadingChangedEventArgs struct {
	RtClass
	*IProximitySensorReadingChangedEventArgs
}

type SensorQuaternion struct {
	RtClass
	*ISensorQuaternion
}

type SensorRotationMatrix struct {
	RtClass
	*ISensorRotationMatrix
}

type SimpleOrientationSensor struct {
	RtClass
	*ISimpleOrientationSensor
}

func NewISimpleOrientationSensorStatics() *ISimpleOrientationSensorStatics {
	var p *ISimpleOrientationSensorStatics
	hs := NewHStr("Windows.Devices.Sensors.SimpleOrientationSensor")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISimpleOrientationSensorStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewISimpleOrientationSensorStatics2() *ISimpleOrientationSensorStatics2 {
	var p *ISimpleOrientationSensorStatics2
	hs := NewHStr("Windows.Devices.Sensors.SimpleOrientationSensor")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISimpleOrientationSensorStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type SimpleOrientationSensorOrientationChangedEventArgs struct {
	RtClass
	*ISimpleOrientationSensorOrientationChangedEventArgs
}
