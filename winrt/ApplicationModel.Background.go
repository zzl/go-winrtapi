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
type AlarmAccessStatus int32

const (
	AlarmAccessStatus_Unspecified                    AlarmAccessStatus = 0
	AlarmAccessStatus_AllowedWithWakeupCapability    AlarmAccessStatus = 1
	AlarmAccessStatus_AllowedWithoutWakeupCapability AlarmAccessStatus = 2
	AlarmAccessStatus_Denied                         AlarmAccessStatus = 3
)

// enum
type ApplicationTriggerResult int32

const (
	ApplicationTriggerResult_Allowed          ApplicationTriggerResult = 0
	ApplicationTriggerResult_CurrentlyRunning ApplicationTriggerResult = 1
	ApplicationTriggerResult_DisabledByPolicy ApplicationTriggerResult = 2
	ApplicationTriggerResult_UnknownError     ApplicationTriggerResult = 3
)

// enum
type BackgroundAccessRequestKind int32

const (
	BackgroundAccessRequestKind_AlwaysAllowed                BackgroundAccessRequestKind = 0
	BackgroundAccessRequestKind_AllowedSubjectToSystemPolicy BackgroundAccessRequestKind = 1
)

// enum
type BackgroundAccessStatus int32

const (
	BackgroundAccessStatus_Unspecified                             BackgroundAccessStatus = 0
	BackgroundAccessStatus_AllowedWithAlwaysOnRealTimeConnectivity BackgroundAccessStatus = 1
	BackgroundAccessStatus_AllowedMayUseActiveRealTimeConnectivity BackgroundAccessStatus = 2
	BackgroundAccessStatus_Denied                                  BackgroundAccessStatus = 3
	BackgroundAccessStatus_AlwaysAllowed                           BackgroundAccessStatus = 4
	BackgroundAccessStatus_AllowedSubjectToSystemPolicy            BackgroundAccessStatus = 5
	BackgroundAccessStatus_DeniedBySystemPolicy                    BackgroundAccessStatus = 6
	BackgroundAccessStatus_DeniedByUser                            BackgroundAccessStatus = 7
)

// enum
type BackgroundTaskCancellationReason int32

const (
	BackgroundTaskCancellationReason_Abort                 BackgroundTaskCancellationReason = 0
	BackgroundTaskCancellationReason_Terminating           BackgroundTaskCancellationReason = 1
	BackgroundTaskCancellationReason_LoggingOff            BackgroundTaskCancellationReason = 2
	BackgroundTaskCancellationReason_ServicingUpdate       BackgroundTaskCancellationReason = 3
	BackgroundTaskCancellationReason_IdleTask              BackgroundTaskCancellationReason = 4
	BackgroundTaskCancellationReason_Uninstall             BackgroundTaskCancellationReason = 5
	BackgroundTaskCancellationReason_ConditionLoss         BackgroundTaskCancellationReason = 6
	BackgroundTaskCancellationReason_SystemPolicy          BackgroundTaskCancellationReason = 7
	BackgroundTaskCancellationReason_QuietHoursEntered     BackgroundTaskCancellationReason = 8
	BackgroundTaskCancellationReason_ExecutionTimeExceeded BackgroundTaskCancellationReason = 9
	BackgroundTaskCancellationReason_ResourceRevocation    BackgroundTaskCancellationReason = 10
	BackgroundTaskCancellationReason_EnergySaver           BackgroundTaskCancellationReason = 11
)

// enum
type BackgroundTaskThrottleCounter int32

const (
	BackgroundTaskThrottleCounter_All     BackgroundTaskThrottleCounter = 0
	BackgroundTaskThrottleCounter_Cpu     BackgroundTaskThrottleCounter = 1
	BackgroundTaskThrottleCounter_Network BackgroundTaskThrottleCounter = 2
)

// enum
type BackgroundWorkCostValue int32

const (
	BackgroundWorkCostValue_Low    BackgroundWorkCostValue = 0
	BackgroundWorkCostValue_Medium BackgroundWorkCostValue = 1
	BackgroundWorkCostValue_High   BackgroundWorkCostValue = 2
)

// enum
type CustomSystemEventTriggerRecurrence int32

const (
	CustomSystemEventTriggerRecurrence_Once   CustomSystemEventTriggerRecurrence = 0
	CustomSystemEventTriggerRecurrence_Always CustomSystemEventTriggerRecurrence = 1
)

// enum
type DeviceTriggerResult int32

const (
	DeviceTriggerResult_Allowed        DeviceTriggerResult = 0
	DeviceTriggerResult_DeniedByUser   DeviceTriggerResult = 1
	DeviceTriggerResult_DeniedBySystem DeviceTriggerResult = 2
	DeviceTriggerResult_LowBattery     DeviceTriggerResult = 3
)

// enum
type LocationTriggerType int32

const (
	LocationTriggerType_Geofence LocationTriggerType = 0
)

// enum
type MediaProcessingTriggerResult int32

const (
	MediaProcessingTriggerResult_Allowed          MediaProcessingTriggerResult = 0
	MediaProcessingTriggerResult_CurrentlyRunning MediaProcessingTriggerResult = 1
	MediaProcessingTriggerResult_DisabledByPolicy MediaProcessingTriggerResult = 2
	MediaProcessingTriggerResult_UnknownError     MediaProcessingTriggerResult = 3
)

// enum
type SystemConditionType int32

const (
	SystemConditionType_Invalid                   SystemConditionType = 0
	SystemConditionType_UserPresent               SystemConditionType = 1
	SystemConditionType_UserNotPresent            SystemConditionType = 2
	SystemConditionType_InternetAvailable         SystemConditionType = 3
	SystemConditionType_InternetNotAvailable      SystemConditionType = 4
	SystemConditionType_SessionConnected          SystemConditionType = 5
	SystemConditionType_SessionDisconnected       SystemConditionType = 6
	SystemConditionType_FreeNetworkAvailable      SystemConditionType = 7
	SystemConditionType_BackgroundWorkCostNotHigh SystemConditionType = 8
)

// enum
type SystemTriggerType int32

const (
	SystemTriggerType_Invalid                      SystemTriggerType = 0
	SystemTriggerType_SmsReceived                  SystemTriggerType = 1
	SystemTriggerType_UserPresent                  SystemTriggerType = 2
	SystemTriggerType_UserAway                     SystemTriggerType = 3
	SystemTriggerType_NetworkStateChange           SystemTriggerType = 4
	SystemTriggerType_ControlChannelReset          SystemTriggerType = 5
	SystemTriggerType_InternetAvailable            SystemTriggerType = 6
	SystemTriggerType_SessionConnected             SystemTriggerType = 7
	SystemTriggerType_ServicingComplete            SystemTriggerType = 8
	SystemTriggerType_LockScreenApplicationAdded   SystemTriggerType = 9
	SystemTriggerType_LockScreenApplicationRemoved SystemTriggerType = 10
	SystemTriggerType_TimeZoneChange               SystemTriggerType = 11
	SystemTriggerType_OnlineIdConnectedStateChange SystemTriggerType = 12
	SystemTriggerType_BackgroundWorkCostChange     SystemTriggerType = 13
	SystemTriggerType_PowerStateChange             SystemTriggerType = 14
	SystemTriggerType_DefaultSignInAccountChange   SystemTriggerType = 15
)

// structs

type BackgroundAlarmApplicationContract struct {
}

// func types

//A6C4BAC0-51F8-4C57-AC3F-156DD1680C4F
type BackgroundTaskCanceledEventHandler func(sender *IBackgroundTaskInstance, reason BackgroundTaskCancellationReason) com.Error

//5B38E929-A086-46A7-A678-439135822BCF
type BackgroundTaskCompletedEventHandler func(sender *IBackgroundTaskRegistration, args *IBackgroundTaskCompletedEventArgs) com.Error

//46E0683C-8A88-4C99-804C-76897F6277A6
type BackgroundTaskProgressEventHandler func(sender *IBackgroundTaskRegistration, args *IBackgroundTaskProgressEventArgs) com.Error

// interfaces

// D0DD4342-E37B-4823-A5FE-6B31DFEFDEB0
var IID_IActivitySensorTrigger = syscall.GUID{0xD0DD4342, 0xE37B, 0x4823,
	[8]byte{0xA5, 0xFE, 0x6B, 0x31, 0xDF, 0xEF, 0xDE, 0xB0}}

type IActivitySensorTriggerInterface interface {
	win32.IInspectableInterface
	Get_SubscribedActivities() *IVector[ActivityType]
	Get_ReportInterval() uint32
	Get_SupportedActivities() *IVectorView[ActivityType]
	Get_MinimumReportInterval() uint32
}

type IActivitySensorTriggerVtbl struct {
	win32.IInspectableVtbl
	Get_SubscribedActivities  uintptr
	Get_ReportInterval        uintptr
	Get_SupportedActivities   uintptr
	Get_MinimumReportInterval uintptr
}

type IActivitySensorTrigger struct {
	win32.IInspectable
}

func (this *IActivitySensorTrigger) Vtbl() *IActivitySensorTriggerVtbl {
	return (*IActivitySensorTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IActivitySensorTrigger) Get_SubscribedActivities() *IVector[ActivityType] {
	var _result *IVector[ActivityType]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SubscribedActivities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IActivitySensorTrigger) Get_ReportInterval() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReportInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IActivitySensorTrigger) Get_SupportedActivities() *IVectorView[ActivityType] {
	var _result *IVectorView[ActivityType]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedActivities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IActivitySensorTrigger) Get_MinimumReportInterval() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinimumReportInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A72691C3-3837-44F7-831B-0132CC872BC3
var IID_IActivitySensorTriggerFactory = syscall.GUID{0xA72691C3, 0x3837, 0x44F7,
	[8]byte{0x83, 0x1B, 0x01, 0x32, 0xCC, 0x87, 0x2B, 0xC3}}

type IActivitySensorTriggerFactoryInterface interface {
	win32.IInspectableInterface
	Create(reportIntervalInMilliseconds uint32) *IActivitySensorTrigger
}

type IActivitySensorTriggerFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IActivitySensorTriggerFactory struct {
	win32.IInspectable
}

func (this *IActivitySensorTriggerFactory) Vtbl() *IActivitySensorTriggerFactoryVtbl {
	return (*IActivitySensorTriggerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IActivitySensorTriggerFactory) Create(reportIntervalInMilliseconds uint32) *IActivitySensorTrigger {
	var _result *IActivitySensorTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(reportIntervalInMilliseconds), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CA03FA3B-CCE6-4DE2-B09B-9628BD33BBBE
var IID_IAlarmApplicationManagerStatics = syscall.GUID{0xCA03FA3B, 0xCCE6, 0x4DE2,
	[8]byte{0xB0, 0x9B, 0x96, 0x28, 0xBD, 0x33, 0xBB, 0xBE}}

type IAlarmApplicationManagerStaticsInterface interface {
	win32.IInspectableInterface
	RequestAccessAsync() *IAsyncOperation[AlarmAccessStatus]
	GetAccessStatus() AlarmAccessStatus
}

type IAlarmApplicationManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	RequestAccessAsync uintptr
	GetAccessStatus    uintptr
}

type IAlarmApplicationManagerStatics struct {
	win32.IInspectable
}

func (this *IAlarmApplicationManagerStatics) Vtbl() *IAlarmApplicationManagerStaticsVtbl {
	return (*IAlarmApplicationManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAlarmApplicationManagerStatics) RequestAccessAsync() *IAsyncOperation[AlarmAccessStatus] {
	var _result *IAsyncOperation[AlarmAccessStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAlarmApplicationManagerStatics) GetAccessStatus() AlarmAccessStatus {
	var _result AlarmAccessStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAccessStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 74D4F496-8D37-44EC-9481-2A0B9854EB48
var IID_IAppBroadcastTrigger = syscall.GUID{0x74D4F496, 0x8D37, 0x44EC,
	[8]byte{0x94, 0x81, 0x2A, 0x0B, 0x98, 0x54, 0xEB, 0x48}}

type IAppBroadcastTriggerInterface interface {
	win32.IInspectableInterface
	Put_ProviderInfo(value *IAppBroadcastTriggerProviderInfo)
	Get_ProviderInfo() *IAppBroadcastTriggerProviderInfo
}

type IAppBroadcastTriggerVtbl struct {
	win32.IInspectableVtbl
	Put_ProviderInfo uintptr
	Get_ProviderInfo uintptr
}

type IAppBroadcastTrigger struct {
	win32.IInspectable
}

func (this *IAppBroadcastTrigger) Vtbl() *IAppBroadcastTriggerVtbl {
	return (*IAppBroadcastTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastTrigger) Put_ProviderInfo(value *IAppBroadcastTriggerProviderInfo) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ProviderInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAppBroadcastTrigger) Get_ProviderInfo() *IAppBroadcastTriggerProviderInfo {
	var _result *IAppBroadcastTriggerProviderInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProviderInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 280B9F44-22F4-4618-A02E-E7E411EB7238
var IID_IAppBroadcastTriggerFactory = syscall.GUID{0x280B9F44, 0x22F4, 0x4618,
	[8]byte{0xA0, 0x2E, 0xE7, 0xE4, 0x11, 0xEB, 0x72, 0x38}}

type IAppBroadcastTriggerFactoryInterface interface {
	win32.IInspectableInterface
	CreateAppBroadcastTrigger(providerKey string) *IAppBroadcastTrigger
}

type IAppBroadcastTriggerFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateAppBroadcastTrigger uintptr
}

type IAppBroadcastTriggerFactory struct {
	win32.IInspectable
}

func (this *IAppBroadcastTriggerFactory) Vtbl() *IAppBroadcastTriggerFactoryVtbl {
	return (*IAppBroadcastTriggerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastTriggerFactory) CreateAppBroadcastTrigger(providerKey string) *IAppBroadcastTrigger {
	var _result *IAppBroadcastTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAppBroadcastTrigger, uintptr(unsafe.Pointer(this)), NewHStr(providerKey).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F219352D-9DE8-4420-9CE2-5EFF8F17376B
var IID_IAppBroadcastTriggerProviderInfo = syscall.GUID{0xF219352D, 0x9DE8, 0x4420,
	[8]byte{0x9C, 0xE2, 0x5E, 0xFF, 0x8F, 0x17, 0x37, 0x6B}}

type IAppBroadcastTriggerProviderInfoInterface interface {
	win32.IInspectableInterface
	Put_DisplayNameResource(value string)
	Get_DisplayNameResource() string
	Put_LogoResource(value string)
	Get_LogoResource() string
	Put_VideoKeyFrameInterval(value TimeSpan)
	Get_VideoKeyFrameInterval() TimeSpan
	Put_MaxVideoBitrate(value uint32)
	Get_MaxVideoBitrate() uint32
	Put_MaxVideoWidth(value uint32)
	Get_MaxVideoWidth() uint32
	Put_MaxVideoHeight(value uint32)
	Get_MaxVideoHeight() uint32
}

type IAppBroadcastTriggerProviderInfoVtbl struct {
	win32.IInspectableVtbl
	Put_DisplayNameResource   uintptr
	Get_DisplayNameResource   uintptr
	Put_LogoResource          uintptr
	Get_LogoResource          uintptr
	Put_VideoKeyFrameInterval uintptr
	Get_VideoKeyFrameInterval uintptr
	Put_MaxVideoBitrate       uintptr
	Get_MaxVideoBitrate       uintptr
	Put_MaxVideoWidth         uintptr
	Get_MaxVideoWidth         uintptr
	Put_MaxVideoHeight        uintptr
	Get_MaxVideoHeight        uintptr
}

type IAppBroadcastTriggerProviderInfo struct {
	win32.IInspectable
}

func (this *IAppBroadcastTriggerProviderInfo) Vtbl() *IAppBroadcastTriggerProviderInfoVtbl {
	return (*IAppBroadcastTriggerProviderInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastTriggerProviderInfo) Put_DisplayNameResource(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DisplayNameResource, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IAppBroadcastTriggerProviderInfo) Get_DisplayNameResource() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayNameResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppBroadcastTriggerProviderInfo) Put_LogoResource(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LogoResource, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IAppBroadcastTriggerProviderInfo) Get_LogoResource() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LogoResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppBroadcastTriggerProviderInfo) Put_VideoKeyFrameInterval(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_VideoKeyFrameInterval, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IAppBroadcastTriggerProviderInfo) Get_VideoKeyFrameInterval() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoKeyFrameInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastTriggerProviderInfo) Put_MaxVideoBitrate(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxVideoBitrate, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppBroadcastTriggerProviderInfo) Get_MaxVideoBitrate() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxVideoBitrate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastTriggerProviderInfo) Put_MaxVideoWidth(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxVideoWidth, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppBroadcastTriggerProviderInfo) Get_MaxVideoWidth() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxVideoWidth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastTriggerProviderInfo) Put_MaxVideoHeight(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxVideoHeight, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppBroadcastTriggerProviderInfo) Get_MaxVideoHeight() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxVideoHeight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 0B468630-9574-492C-9E93-1A3AE6335FE9
var IID_IApplicationTrigger = syscall.GUID{0x0B468630, 0x9574, 0x492C,
	[8]byte{0x9E, 0x93, 0x1A, 0x3A, 0xE6, 0x33, 0x5F, 0xE9}}

type IApplicationTriggerInterface interface {
	win32.IInspectableInterface
	RequestAsync() *IAsyncOperation[ApplicationTriggerResult]
	RequestAsyncWithArguments(arguments *IPropertySet) *IAsyncOperation[ApplicationTriggerResult]
}

type IApplicationTriggerVtbl struct {
	win32.IInspectableVtbl
	RequestAsync              uintptr
	RequestAsyncWithArguments uintptr
}

type IApplicationTrigger struct {
	win32.IInspectable
}

func (this *IApplicationTrigger) Vtbl() *IApplicationTriggerVtbl {
	return (*IApplicationTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationTrigger) RequestAsync() *IAsyncOperation[ApplicationTriggerResult] {
	var _result *IAsyncOperation[ApplicationTriggerResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationTrigger) RequestAsyncWithArguments(arguments *IPropertySet) *IAsyncOperation[ApplicationTriggerResult] {
	var _result *IAsyncOperation[ApplicationTriggerResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAsyncWithArguments, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(arguments)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 97DC6AB2-2219-4A9E-9C5E-41D047F76E82
var IID_IApplicationTriggerDetails = syscall.GUID{0x97DC6AB2, 0x2219, 0x4A9E,
	[8]byte{0x9C, 0x5E, 0x41, 0xD0, 0x47, 0xF7, 0x6E, 0x82}}

type IApplicationTriggerDetailsInterface interface {
	win32.IInspectableInterface
	Get_Arguments() *IPropertySet
}

type IApplicationTriggerDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_Arguments uintptr
}

type IApplicationTriggerDetails struct {
	win32.IInspectable
}

func (this *IApplicationTriggerDetails) Vtbl() *IApplicationTriggerDetailsVtbl {
	return (*IApplicationTriggerDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationTriggerDetails) Get_Arguments() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Arguments, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 64D4040C-C201-42AD-AA2A-E21BA3425B6D
var IID_IAppointmentStoreNotificationTrigger = syscall.GUID{0x64D4040C, 0xC201, 0x42AD,
	[8]byte{0xAA, 0x2A, 0xE2, 0x1B, 0xA3, 0x42, 0x5B, 0x6D}}

type IAppointmentStoreNotificationTriggerInterface interface {
	win32.IInspectableInterface
}

type IAppointmentStoreNotificationTriggerVtbl struct {
	win32.IInspectableVtbl
}

type IAppointmentStoreNotificationTrigger struct {
	win32.IInspectable
}

func (this *IAppointmentStoreNotificationTrigger) Vtbl() *IAppointmentStoreNotificationTriggerVtbl {
	return (*IAppointmentStoreNotificationTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// AE48A1EE-8951-400A-8302-9C9C9A2A3A3B
var IID_IBackgroundCondition = syscall.GUID{0xAE48A1EE, 0x8951, 0x400A,
	[8]byte{0x83, 0x02, 0x9C, 0x9C, 0x9A, 0x2A, 0x3A, 0x3B}}

type IBackgroundConditionInterface interface {
	win32.IInspectableInterface
}

type IBackgroundConditionVtbl struct {
	win32.IInspectableVtbl
}

type IBackgroundCondition struct {
	win32.IInspectable
}

func (this *IBackgroundCondition) Vtbl() *IBackgroundConditionVtbl {
	return (*IBackgroundConditionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// E826EA58-66A9-4D41-83D4-B4C18C87B846
var IID_IBackgroundExecutionManagerStatics = syscall.GUID{0xE826EA58, 0x66A9, 0x4D41,
	[8]byte{0x83, 0xD4, 0xB4, 0xC1, 0x8C, 0x87, 0xB8, 0x46}}

type IBackgroundExecutionManagerStaticsInterface interface {
	win32.IInspectableInterface
	RequestAccessAsync() *IAsyncOperation[BackgroundAccessStatus]
	RequestAccessForApplicationAsync(applicationId string) *IAsyncOperation[BackgroundAccessStatus]
	RemoveAccess()
	RemoveAccessForApplication(applicationId string)
	GetAccessStatus() BackgroundAccessStatus
	GetAccessStatusForApplication(applicationId string) BackgroundAccessStatus
}

type IBackgroundExecutionManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	RequestAccessAsync               uintptr
	RequestAccessForApplicationAsync uintptr
	RemoveAccess                     uintptr
	RemoveAccessForApplication       uintptr
	GetAccessStatus                  uintptr
	GetAccessStatusForApplication    uintptr
}

type IBackgroundExecutionManagerStatics struct {
	win32.IInspectable
}

func (this *IBackgroundExecutionManagerStatics) Vtbl() *IBackgroundExecutionManagerStaticsVtbl {
	return (*IBackgroundExecutionManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundExecutionManagerStatics) RequestAccessAsync() *IAsyncOperation[BackgroundAccessStatus] {
	var _result *IAsyncOperation[BackgroundAccessStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundExecutionManagerStatics) RequestAccessForApplicationAsync(applicationId string) *IAsyncOperation[BackgroundAccessStatus] {
	var _result *IAsyncOperation[BackgroundAccessStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessForApplicationAsync, uintptr(unsafe.Pointer(this)), NewHStr(applicationId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundExecutionManagerStatics) RemoveAccess() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveAccess, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IBackgroundExecutionManagerStatics) RemoveAccessForApplication(applicationId string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveAccessForApplication, uintptr(unsafe.Pointer(this)), NewHStr(applicationId).Ptr)
	_ = _hr
}

func (this *IBackgroundExecutionManagerStatics) GetAccessStatus() BackgroundAccessStatus {
	var _result BackgroundAccessStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAccessStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundExecutionManagerStatics) GetAccessStatusForApplication(applicationId string) BackgroundAccessStatus {
	var _result BackgroundAccessStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAccessStatusForApplication, uintptr(unsafe.Pointer(this)), NewHStr(applicationId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 469B24EF-9BBB-4E18-999A-FD6512931BE9
var IID_IBackgroundExecutionManagerStatics2 = syscall.GUID{0x469B24EF, 0x9BBB, 0x4E18,
	[8]byte{0x99, 0x9A, 0xFD, 0x65, 0x12, 0x93, 0x1B, 0xE9}}

type IBackgroundExecutionManagerStatics2Interface interface {
	win32.IInspectableInterface
	RequestAccessKindAsync(requestedAccess BackgroundAccessRequestKind, reason string) *IAsyncOperation[bool]
}

type IBackgroundExecutionManagerStatics2Vtbl struct {
	win32.IInspectableVtbl
	RequestAccessKindAsync uintptr
}

type IBackgroundExecutionManagerStatics2 struct {
	win32.IInspectable
}

func (this *IBackgroundExecutionManagerStatics2) Vtbl() *IBackgroundExecutionManagerStatics2Vtbl {
	return (*IBackgroundExecutionManagerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundExecutionManagerStatics2) RequestAccessKindAsync(requestedAccess BackgroundAccessRequestKind, reason string) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessKindAsync, uintptr(unsafe.Pointer(this)), uintptr(requestedAccess), NewHStr(reason).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 98A5D3F6-5A25-5B6C-9192-D77A43DFEDC4
var IID_IBackgroundExecutionManagerStatics3 = syscall.GUID{0x98A5D3F6, 0x5A25, 0x5B6C,
	[8]byte{0x91, 0x92, 0xD7, 0x7A, 0x43, 0xDF, 0xED, 0xC4}}

type IBackgroundExecutionManagerStatics3Interface interface {
	win32.IInspectableInterface
	RequestAccessKindForModernStandbyAsync(requestedAccess BackgroundAccessRequestKind, reason string) *IAsyncOperation[bool]
	GetAccessStatusForModernStandby() BackgroundAccessStatus
	GetAccessStatusForModernStandbyForApplication(applicationId string) BackgroundAccessStatus
}

type IBackgroundExecutionManagerStatics3Vtbl struct {
	win32.IInspectableVtbl
	RequestAccessKindForModernStandbyAsync        uintptr
	GetAccessStatusForModernStandby               uintptr
	GetAccessStatusForModernStandbyForApplication uintptr
}

type IBackgroundExecutionManagerStatics3 struct {
	win32.IInspectable
}

func (this *IBackgroundExecutionManagerStatics3) Vtbl() *IBackgroundExecutionManagerStatics3Vtbl {
	return (*IBackgroundExecutionManagerStatics3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundExecutionManagerStatics3) RequestAccessKindForModernStandbyAsync(requestedAccess BackgroundAccessRequestKind, reason string) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessKindForModernStandbyAsync, uintptr(unsafe.Pointer(this)), uintptr(requestedAccess), NewHStr(reason).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundExecutionManagerStatics3) GetAccessStatusForModernStandby() BackgroundAccessStatus {
	var _result BackgroundAccessStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAccessStatusForModernStandby, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundExecutionManagerStatics3) GetAccessStatusForModernStandbyForApplication(applicationId string) BackgroundAccessStatus {
	var _result BackgroundAccessStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAccessStatusForModernStandbyForApplication, uintptr(unsafe.Pointer(this)), NewHStr(applicationId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7D13D534-FD12-43CE-8C22-EA1FF13C06DF
var IID_IBackgroundTask = syscall.GUID{0x7D13D534, 0xFD12, 0x43CE,
	[8]byte{0x8C, 0x22, 0xEA, 0x1F, 0xF1, 0x3C, 0x06, 0xDF}}

type IBackgroundTaskInterface interface {
	win32.IInspectableInterface
	Run(taskInstance *IBackgroundTaskInstance)
}

type IBackgroundTaskVtbl struct {
	win32.IInspectableVtbl
	Run uintptr
}

type IBackgroundTask struct {
	win32.IInspectable
}

func (this *IBackgroundTask) Vtbl() *IBackgroundTaskVtbl {
	return (*IBackgroundTaskVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTask) Run(taskInstance *IBackgroundTaskInstance) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Run, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(taskInstance)))
	_ = _hr
}

// 0351550E-3E64-4572-A93A-84075A37C917
var IID_IBackgroundTaskBuilder = syscall.GUID{0x0351550E, 0x3E64, 0x4572,
	[8]byte{0xA9, 0x3A, 0x84, 0x07, 0x5A, 0x37, 0xC9, 0x17}}

type IBackgroundTaskBuilderInterface interface {
	win32.IInspectableInterface
	Put_TaskEntryPoint(value string)
	Get_TaskEntryPoint() string
	SetTrigger(trigger *IBackgroundTrigger)
	AddCondition(condition *IBackgroundCondition)
	Put_Name(value string)
	Get_Name() string
	Register() *IBackgroundTaskRegistration
}

type IBackgroundTaskBuilderVtbl struct {
	win32.IInspectableVtbl
	Put_TaskEntryPoint uintptr
	Get_TaskEntryPoint uintptr
	SetTrigger         uintptr
	AddCondition       uintptr
	Put_Name           uintptr
	Get_Name           uintptr
	Register           uintptr
}

type IBackgroundTaskBuilder struct {
	win32.IInspectable
}

func (this *IBackgroundTaskBuilder) Vtbl() *IBackgroundTaskBuilderVtbl {
	return (*IBackgroundTaskBuilderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTaskBuilder) Put_TaskEntryPoint(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TaskEntryPoint, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IBackgroundTaskBuilder) Get_TaskEntryPoint() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TaskEntryPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBackgroundTaskBuilder) SetTrigger(trigger *IBackgroundTrigger) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetTrigger, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(trigger)))
	_ = _hr
}

func (this *IBackgroundTaskBuilder) AddCondition(condition *IBackgroundCondition) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddCondition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(condition)))
	_ = _hr
}

func (this *IBackgroundTaskBuilder) Put_Name(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Name, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IBackgroundTaskBuilder) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBackgroundTaskBuilder) Register() *IBackgroundTaskRegistration {
	var _result *IBackgroundTaskRegistration
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Register, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6AE7CFB1-104F-406D-8DB6-844A570F42BB
var IID_IBackgroundTaskBuilder2 = syscall.GUID{0x6AE7CFB1, 0x104F, 0x406D,
	[8]byte{0x8D, 0xB6, 0x84, 0x4A, 0x57, 0x0F, 0x42, 0xBB}}

type IBackgroundTaskBuilder2Interface interface {
	win32.IInspectableInterface
	Put_CancelOnConditionLoss(value bool)
	Get_CancelOnConditionLoss() bool
}

type IBackgroundTaskBuilder2Vtbl struct {
	win32.IInspectableVtbl
	Put_CancelOnConditionLoss uintptr
	Get_CancelOnConditionLoss uintptr
}

type IBackgroundTaskBuilder2 struct {
	win32.IInspectable
}

func (this *IBackgroundTaskBuilder2) Vtbl() *IBackgroundTaskBuilder2Vtbl {
	return (*IBackgroundTaskBuilder2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTaskBuilder2) Put_CancelOnConditionLoss(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CancelOnConditionLoss, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IBackgroundTaskBuilder2) Get_CancelOnConditionLoss() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CancelOnConditionLoss, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 28C74F4A-8BA9-4C09-A24F-19683E2C924C
var IID_IBackgroundTaskBuilder3 = syscall.GUID{0x28C74F4A, 0x8BA9, 0x4C09,
	[8]byte{0xA2, 0x4F, 0x19, 0x68, 0x3E, 0x2C, 0x92, 0x4C}}

type IBackgroundTaskBuilder3Interface interface {
	win32.IInspectableInterface
	Put_IsNetworkRequested(value bool)
	Get_IsNetworkRequested() bool
}

type IBackgroundTaskBuilder3Vtbl struct {
	win32.IInspectableVtbl
	Put_IsNetworkRequested uintptr
	Get_IsNetworkRequested uintptr
}

type IBackgroundTaskBuilder3 struct {
	win32.IInspectable
}

func (this *IBackgroundTaskBuilder3) Vtbl() *IBackgroundTaskBuilder3Vtbl {
	return (*IBackgroundTaskBuilder3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTaskBuilder3) Put_IsNetworkRequested(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsNetworkRequested, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IBackgroundTaskBuilder3) Get_IsNetworkRequested() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsNetworkRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4755E522-CBA2-4E35-BD16-A6DA7F1C19AA
var IID_IBackgroundTaskBuilder4 = syscall.GUID{0x4755E522, 0xCBA2, 0x4E35,
	[8]byte{0xBD, 0x16, 0xA6, 0xDA, 0x7F, 0x1C, 0x19, 0xAA}}

type IBackgroundTaskBuilder4Interface interface {
	win32.IInspectableInterface
	Get_TaskGroup() *IBackgroundTaskRegistrationGroup
	Put_TaskGroup(value *IBackgroundTaskRegistrationGroup)
}

type IBackgroundTaskBuilder4Vtbl struct {
	win32.IInspectableVtbl
	Get_TaskGroup uintptr
	Put_TaskGroup uintptr
}

type IBackgroundTaskBuilder4 struct {
	win32.IInspectable
}

func (this *IBackgroundTaskBuilder4) Vtbl() *IBackgroundTaskBuilder4Vtbl {
	return (*IBackgroundTaskBuilder4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTaskBuilder4) Get_TaskGroup() *IBackgroundTaskRegistrationGroup {
	var _result *IBackgroundTaskRegistrationGroup
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TaskGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundTaskBuilder4) Put_TaskGroup(value *IBackgroundTaskRegistrationGroup) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TaskGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 077103F6-99F5-4AF4-BCAD-4731D0330D43
var IID_IBackgroundTaskBuilder5 = syscall.GUID{0x077103F6, 0x99F5, 0x4AF4,
	[8]byte{0xBC, 0xAD, 0x47, 0x31, 0xD0, 0x33, 0x0D, 0x43}}

type IBackgroundTaskBuilder5Interface interface {
	win32.IInspectableInterface
	SetTaskEntryPointClsid(TaskEntryPoint syscall.GUID)
}

type IBackgroundTaskBuilder5Vtbl struct {
	win32.IInspectableVtbl
	SetTaskEntryPointClsid uintptr
}

type IBackgroundTaskBuilder5 struct {
	win32.IInspectable
}

func (this *IBackgroundTaskBuilder5) Vtbl() *IBackgroundTaskBuilder5Vtbl {
	return (*IBackgroundTaskBuilder5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTaskBuilder5) SetTaskEntryPointClsid(TaskEntryPoint syscall.GUID) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetTaskEntryPointClsid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&TaskEntryPoint)))
	_ = _hr
}

// 565D25CF-F209-48F4-9967-2B184F7BFBF0
var IID_IBackgroundTaskCompletedEventArgs = syscall.GUID{0x565D25CF, 0xF209, 0x48F4,
	[8]byte{0x99, 0x67, 0x2B, 0x18, 0x4F, 0x7B, 0xFB, 0xF0}}

type IBackgroundTaskCompletedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_InstanceId() syscall.GUID
	CheckResult()
}

type IBackgroundTaskCompletedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_InstanceId uintptr
	CheckResult    uintptr
}

type IBackgroundTaskCompletedEventArgs struct {
	win32.IInspectable
}

func (this *IBackgroundTaskCompletedEventArgs) Vtbl() *IBackgroundTaskCompletedEventArgsVtbl {
	return (*IBackgroundTaskCompletedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTaskCompletedEventArgs) Get_InstanceId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InstanceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundTaskCompletedEventArgs) CheckResult() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CheckResult, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 93CC156D-AF27-4DD3-846E-24EE40CADD25
var IID_IBackgroundTaskDeferral = syscall.GUID{0x93CC156D, 0xAF27, 0x4DD3,
	[8]byte{0x84, 0x6E, 0x24, 0xEE, 0x40, 0xCA, 0xDD, 0x25}}

type IBackgroundTaskDeferralInterface interface {
	win32.IInspectableInterface
	Complete()
}

type IBackgroundTaskDeferralVtbl struct {
	win32.IInspectableVtbl
	Complete uintptr
}

type IBackgroundTaskDeferral struct {
	win32.IInspectable
}

func (this *IBackgroundTaskDeferral) Vtbl() *IBackgroundTaskDeferralVtbl {
	return (*IBackgroundTaskDeferralVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTaskDeferral) Complete() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Complete, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 865BDA7A-21D8-4573-8F32-928A1B0641F6
var IID_IBackgroundTaskInstance = syscall.GUID{0x865BDA7A, 0x21D8, 0x4573,
	[8]byte{0x8F, 0x32, 0x92, 0x8A, 0x1B, 0x06, 0x41, 0xF6}}

type IBackgroundTaskInstanceInterface interface {
	win32.IInspectableInterface
	Get_InstanceId() syscall.GUID
	Get_Task() *IBackgroundTaskRegistration
	Get_Progress() uint32
	Put_Progress(value uint32)
	Get_TriggerDetails() interface{}
	Add_Canceled(cancelHandler BackgroundTaskCanceledEventHandler) EventRegistrationToken
	Remove_Canceled(cookie EventRegistrationToken)
	Get_SuspendedCount() uint32
	GetDeferral() *IBackgroundTaskDeferral
}

type IBackgroundTaskInstanceVtbl struct {
	win32.IInspectableVtbl
	Get_InstanceId     uintptr
	Get_Task           uintptr
	Get_Progress       uintptr
	Put_Progress       uintptr
	Get_TriggerDetails uintptr
	Add_Canceled       uintptr
	Remove_Canceled    uintptr
	Get_SuspendedCount uintptr
	GetDeferral        uintptr
}

type IBackgroundTaskInstance struct {
	win32.IInspectable
}

func (this *IBackgroundTaskInstance) Vtbl() *IBackgroundTaskInstanceVtbl {
	return (*IBackgroundTaskInstanceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTaskInstance) Get_InstanceId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InstanceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundTaskInstance) Get_Task() *IBackgroundTaskRegistration {
	var _result *IBackgroundTaskRegistration
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Task, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundTaskInstance) Get_Progress() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Progress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundTaskInstance) Put_Progress(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Progress, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IBackgroundTaskInstance) Get_TriggerDetails() interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TriggerDetails, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundTaskInstance) Add_Canceled(cancelHandler BackgroundTaskCanceledEventHandler) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Canceled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(cancelHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundTaskInstance) Remove_Canceled(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Canceled, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IBackgroundTaskInstance) Get_SuspendedCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SuspendedCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundTaskInstance) GetDeferral() *IBackgroundTaskDeferral {
	var _result *IBackgroundTaskDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4F7D0176-0C76-4FB4-896D-5DE1864122F6
var IID_IBackgroundTaskInstance2 = syscall.GUID{0x4F7D0176, 0x0C76, 0x4FB4,
	[8]byte{0x89, 0x6D, 0x5D, 0xE1, 0x86, 0x41, 0x22, 0xF6}}

type IBackgroundTaskInstance2Interface interface {
	win32.IInspectableInterface
	GetThrottleCount(counter BackgroundTaskThrottleCounter) uint32
}

type IBackgroundTaskInstance2Vtbl struct {
	win32.IInspectableVtbl
	GetThrottleCount uintptr
}

type IBackgroundTaskInstance2 struct {
	win32.IInspectable
}

func (this *IBackgroundTaskInstance2) Vtbl() *IBackgroundTaskInstance2Vtbl {
	return (*IBackgroundTaskInstance2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTaskInstance2) GetThrottleCount(counter BackgroundTaskThrottleCounter) uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetThrottleCount, uintptr(unsafe.Pointer(this)), uintptr(counter), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7F29F23C-AA04-4B08-97B0-06D874CDABF5
var IID_IBackgroundTaskInstance4 = syscall.GUID{0x7F29F23C, 0xAA04, 0x4B08,
	[8]byte{0x97, 0xB0, 0x06, 0xD8, 0x74, 0xCD, 0xAB, 0xF5}}

type IBackgroundTaskInstance4Interface interface {
	win32.IInspectableInterface
	Get_User() *IUser
}

type IBackgroundTaskInstance4Vtbl struct {
	win32.IInspectableVtbl
	Get_User uintptr
}

type IBackgroundTaskInstance4 struct {
	win32.IInspectable
}

func (this *IBackgroundTaskInstance4) Vtbl() *IBackgroundTaskInstance4Vtbl {
	return (*IBackgroundTaskInstance4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTaskInstance4) Get_User() *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_User, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FB1468AC-8332-4D0A-9532-03EAE684DA31
var IID_IBackgroundTaskProgressEventArgs = syscall.GUID{0xFB1468AC, 0x8332, 0x4D0A,
	[8]byte{0x95, 0x32, 0x03, 0xEA, 0xE6, 0x84, 0xDA, 0x31}}

type IBackgroundTaskProgressEventArgsInterface interface {
	win32.IInspectableInterface
	Get_InstanceId() syscall.GUID
	Get_Progress() uint32
}

type IBackgroundTaskProgressEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_InstanceId uintptr
	Get_Progress   uintptr
}

type IBackgroundTaskProgressEventArgs struct {
	win32.IInspectable
}

func (this *IBackgroundTaskProgressEventArgs) Vtbl() *IBackgroundTaskProgressEventArgsVtbl {
	return (*IBackgroundTaskProgressEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTaskProgressEventArgs) Get_InstanceId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InstanceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundTaskProgressEventArgs) Get_Progress() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Progress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 10654CC2-A26E-43BF-8C12-1FB40DBFBFA0
var IID_IBackgroundTaskRegistration = syscall.GUID{0x10654CC2, 0xA26E, 0x43BF,
	[8]byte{0x8C, 0x12, 0x1F, 0xB4, 0x0D, 0xBF, 0xBF, 0xA0}}

type IBackgroundTaskRegistrationInterface interface {
	win32.IInspectableInterface
	Get_TaskId() syscall.GUID
	Get_Name() string
	Add_Progress(handler BackgroundTaskProgressEventHandler) EventRegistrationToken
	Remove_Progress(cookie EventRegistrationToken)
	Add_Completed(handler BackgroundTaskCompletedEventHandler) EventRegistrationToken
	Remove_Completed(cookie EventRegistrationToken)
	Unregister(cancelTask bool)
}

type IBackgroundTaskRegistrationVtbl struct {
	win32.IInspectableVtbl
	Get_TaskId       uintptr
	Get_Name         uintptr
	Add_Progress     uintptr
	Remove_Progress  uintptr
	Add_Completed    uintptr
	Remove_Completed uintptr
	Unregister       uintptr
}

type IBackgroundTaskRegistration struct {
	win32.IInspectable
}

func (this *IBackgroundTaskRegistration) Vtbl() *IBackgroundTaskRegistrationVtbl {
	return (*IBackgroundTaskRegistrationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTaskRegistration) Get_TaskId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TaskId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundTaskRegistration) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBackgroundTaskRegistration) Add_Progress(handler BackgroundTaskProgressEventHandler) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Progress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundTaskRegistration) Remove_Progress(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Progress, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IBackgroundTaskRegistration) Add_Completed(handler BackgroundTaskCompletedEventHandler) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Completed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundTaskRegistration) Remove_Completed(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Completed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IBackgroundTaskRegistration) Unregister(cancelTask bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Unregister, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&cancelTask))))
	_ = _hr
}

// 6138C703-BB86-4112-AFC3-7F939B166E3B
var IID_IBackgroundTaskRegistration2 = syscall.GUID{0x6138C703, 0xBB86, 0x4112,
	[8]byte{0xAF, 0xC3, 0x7F, 0x93, 0x9B, 0x16, 0x6E, 0x3B}}

type IBackgroundTaskRegistration2Interface interface {
	win32.IInspectableInterface
	Get_Trigger() *IBackgroundTrigger
}

type IBackgroundTaskRegistration2Vtbl struct {
	win32.IInspectableVtbl
	Get_Trigger uintptr
}

type IBackgroundTaskRegistration2 struct {
	win32.IInspectable
}

func (this *IBackgroundTaskRegistration2) Vtbl() *IBackgroundTaskRegistration2Vtbl {
	return (*IBackgroundTaskRegistration2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTaskRegistration2) Get_Trigger() *IBackgroundTrigger {
	var _result *IBackgroundTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Trigger, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FE338195-9423-4D8B-830D-B1DD2C7BADD5
var IID_IBackgroundTaskRegistration3 = syscall.GUID{0xFE338195, 0x9423, 0x4D8B,
	[8]byte{0x83, 0x0D, 0xB1, 0xDD, 0x2C, 0x7B, 0xAD, 0xD5}}

type IBackgroundTaskRegistration3Interface interface {
	win32.IInspectableInterface
	Get_TaskGroup() *IBackgroundTaskRegistrationGroup
}

type IBackgroundTaskRegistration3Vtbl struct {
	win32.IInspectableVtbl
	Get_TaskGroup uintptr
}

type IBackgroundTaskRegistration3 struct {
	win32.IInspectable
}

func (this *IBackgroundTaskRegistration3) Vtbl() *IBackgroundTaskRegistration3Vtbl {
	return (*IBackgroundTaskRegistration3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTaskRegistration3) Get_TaskGroup() *IBackgroundTaskRegistrationGroup {
	var _result *IBackgroundTaskRegistrationGroup
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TaskGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2AB1919A-871B-4167-8A76-055CD67B5B23
var IID_IBackgroundTaskRegistrationGroup = syscall.GUID{0x2AB1919A, 0x871B, 0x4167,
	[8]byte{0x8A, 0x76, 0x05, 0x5C, 0xD6, 0x7B, 0x5B, 0x23}}

type IBackgroundTaskRegistrationGroupInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	Get_Name() string
	Add_BackgroundActivated(handler TypedEventHandler[*IBackgroundTaskRegistrationGroup, *IBackgroundActivatedEventArgs]) EventRegistrationToken
	Remove_BackgroundActivated(token EventRegistrationToken)
	Get_AllTasks() *IMapView[syscall.GUID, *IBackgroundTaskRegistration]
}

type IBackgroundTaskRegistrationGroupVtbl struct {
	win32.IInspectableVtbl
	Get_Id                     uintptr
	Get_Name                   uintptr
	Add_BackgroundActivated    uintptr
	Remove_BackgroundActivated uintptr
	Get_AllTasks               uintptr
}

type IBackgroundTaskRegistrationGroup struct {
	win32.IInspectable
}

func (this *IBackgroundTaskRegistrationGroup) Vtbl() *IBackgroundTaskRegistrationGroupVtbl {
	return (*IBackgroundTaskRegistrationGroupVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTaskRegistrationGroup) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBackgroundTaskRegistrationGroup) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBackgroundTaskRegistrationGroup) Add_BackgroundActivated(handler TypedEventHandler[*IBackgroundTaskRegistrationGroup, *IBackgroundActivatedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_BackgroundActivated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundTaskRegistrationGroup) Remove_BackgroundActivated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_BackgroundActivated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IBackgroundTaskRegistrationGroup) Get_AllTasks() *IMapView[syscall.GUID, *IBackgroundTaskRegistration] {
	var _result *IMapView[syscall.GUID, *IBackgroundTaskRegistration]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllTasks, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 83D92B69-44CF-4631-9740-03C7D8741BC5
var IID_IBackgroundTaskRegistrationGroupFactory = syscall.GUID{0x83D92B69, 0x44CF, 0x4631,
	[8]byte{0x97, 0x40, 0x03, 0xC7, 0xD8, 0x74, 0x1B, 0xC5}}

type IBackgroundTaskRegistrationGroupFactoryInterface interface {
	win32.IInspectableInterface
	Create(id string) *IBackgroundTaskRegistrationGroup
	CreateWithName(id string, name string) *IBackgroundTaskRegistrationGroup
}

type IBackgroundTaskRegistrationGroupFactoryVtbl struct {
	win32.IInspectableVtbl
	Create         uintptr
	CreateWithName uintptr
}

type IBackgroundTaskRegistrationGroupFactory struct {
	win32.IInspectable
}

func (this *IBackgroundTaskRegistrationGroupFactory) Vtbl() *IBackgroundTaskRegistrationGroupFactoryVtbl {
	return (*IBackgroundTaskRegistrationGroupFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTaskRegistrationGroupFactory) Create(id string) *IBackgroundTaskRegistrationGroup {
	var _result *IBackgroundTaskRegistrationGroup
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(id).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundTaskRegistrationGroupFactory) CreateWithName(id string, name string) *IBackgroundTaskRegistrationGroup {
	var _result *IBackgroundTaskRegistrationGroup
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithName, uintptr(unsafe.Pointer(this)), NewHStr(id).Ptr, NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4C542F69-B000-42BA-A093-6A563C65E3F8
var IID_IBackgroundTaskRegistrationStatics = syscall.GUID{0x4C542F69, 0xB000, 0x42BA,
	[8]byte{0xA0, 0x93, 0x6A, 0x56, 0x3C, 0x65, 0xE3, 0xF8}}

type IBackgroundTaskRegistrationStaticsInterface interface {
	win32.IInspectableInterface
	Get_AllTasks() *IMapView[syscall.GUID, *IBackgroundTaskRegistration]
}

type IBackgroundTaskRegistrationStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_AllTasks uintptr
}

type IBackgroundTaskRegistrationStatics struct {
	win32.IInspectable
}

func (this *IBackgroundTaskRegistrationStatics) Vtbl() *IBackgroundTaskRegistrationStaticsVtbl {
	return (*IBackgroundTaskRegistrationStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTaskRegistrationStatics) Get_AllTasks() *IMapView[syscall.GUID, *IBackgroundTaskRegistration] {
	var _result *IMapView[syscall.GUID, *IBackgroundTaskRegistration]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllTasks, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 174B671E-B20D-4FA9-AD9A-E93AD6C71E01
var IID_IBackgroundTaskRegistrationStatics2 = syscall.GUID{0x174B671E, 0xB20D, 0x4FA9,
	[8]byte{0xAD, 0x9A, 0xE9, 0x3A, 0xD6, 0xC7, 0x1E, 0x01}}

type IBackgroundTaskRegistrationStatics2Interface interface {
	win32.IInspectableInterface
	Get_AllTaskGroups() *IMapView[string, *IBackgroundTaskRegistrationGroup]
	GetTaskGroup(groupId string) *IBackgroundTaskRegistrationGroup
}

type IBackgroundTaskRegistrationStatics2Vtbl struct {
	win32.IInspectableVtbl
	Get_AllTaskGroups uintptr
	GetTaskGroup      uintptr
}

type IBackgroundTaskRegistrationStatics2 struct {
	win32.IInspectable
}

func (this *IBackgroundTaskRegistrationStatics2) Vtbl() *IBackgroundTaskRegistrationStatics2Vtbl {
	return (*IBackgroundTaskRegistrationStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTaskRegistrationStatics2) Get_AllTaskGroups() *IMapView[string, *IBackgroundTaskRegistrationGroup] {
	var _result *IMapView[string, *IBackgroundTaskRegistrationGroup]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllTaskGroups, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundTaskRegistrationStatics2) GetTaskGroup(groupId string) *IBackgroundTaskRegistrationGroup {
	var _result *IBackgroundTaskRegistrationGroup
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetTaskGroup, uintptr(unsafe.Pointer(this)), NewHStr(groupId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 84B3A058-6027-4B87-9790-BDF3F757DBD7
var IID_IBackgroundTrigger = syscall.GUID{0x84B3A058, 0x6027, 0x4B87,
	[8]byte{0x97, 0x90, 0xBD, 0xF3, 0xF7, 0x57, 0xDB, 0xD7}}

type IBackgroundTriggerInterface interface {
	win32.IInspectableInterface
}

type IBackgroundTriggerVtbl struct {
	win32.IInspectableVtbl
}

type IBackgroundTrigger struct {
	win32.IInspectable
}

func (this *IBackgroundTrigger) Vtbl() *IBackgroundTriggerVtbl {
	return (*IBackgroundTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// C740A662-C310-4B82-B3E3-3BCFB9E4C77D
var IID_IBackgroundWorkCostStatics = syscall.GUID{0xC740A662, 0xC310, 0x4B82,
	[8]byte{0xB3, 0xE3, 0x3B, 0xCF, 0xB9, 0xE4, 0xC7, 0x7D}}

type IBackgroundWorkCostStaticsInterface interface {
	win32.IInspectableInterface
	Get_CurrentBackgroundWorkCost() BackgroundWorkCostValue
}

type IBackgroundWorkCostStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_CurrentBackgroundWorkCost uintptr
}

type IBackgroundWorkCostStatics struct {
	win32.IInspectable
}

func (this *IBackgroundWorkCostStatics) Vtbl() *IBackgroundWorkCostStaticsVtbl {
	return (*IBackgroundWorkCostStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundWorkCostStatics) Get_CurrentBackgroundWorkCost() BackgroundWorkCostValue {
	var _result BackgroundWorkCostValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentBackgroundWorkCost, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// AB3E2612-25D3-48AE-8724-D81877AE6129
var IID_IBluetoothLEAdvertisementPublisherTrigger = syscall.GUID{0xAB3E2612, 0x25D3, 0x48AE,
	[8]byte{0x87, 0x24, 0xD8, 0x18, 0x77, 0xAE, 0x61, 0x29}}

type IBluetoothLEAdvertisementPublisherTriggerInterface interface {
	win32.IInspectableInterface
	Get_Advertisement() *IBluetoothLEAdvertisement
}

type IBluetoothLEAdvertisementPublisherTriggerVtbl struct {
	win32.IInspectableVtbl
	Get_Advertisement uintptr
}

type IBluetoothLEAdvertisementPublisherTrigger struct {
	win32.IInspectable
}

func (this *IBluetoothLEAdvertisementPublisherTrigger) Vtbl() *IBluetoothLEAdvertisementPublisherTriggerVtbl {
	return (*IBluetoothLEAdvertisementPublisherTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAdvertisementPublisherTrigger) Get_Advertisement() *IBluetoothLEAdvertisement {
	var _result *IBluetoothLEAdvertisement
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Advertisement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AA28D064-38F4-597D-B597-4E55588C6503
var IID_IBluetoothLEAdvertisementPublisherTrigger2 = syscall.GUID{0xAA28D064, 0x38F4, 0x597D,
	[8]byte{0xB5, 0x97, 0x4E, 0x55, 0x58, 0x8C, 0x65, 0x03}}

type IBluetoothLEAdvertisementPublisherTrigger2Interface interface {
	win32.IInspectableInterface
	Get_PreferredTransmitPowerLevelInDBm() *IReference[int16]
	Put_PreferredTransmitPowerLevelInDBm(value *IReference[int16])
	Get_UseExtendedFormat() bool
	Put_UseExtendedFormat(value bool)
	Get_IsAnonymous() bool
	Put_IsAnonymous(value bool)
	Get_IncludeTransmitPowerLevel() bool
	Put_IncludeTransmitPowerLevel(value bool)
}

type IBluetoothLEAdvertisementPublisherTrigger2Vtbl struct {
	win32.IInspectableVtbl
	Get_PreferredTransmitPowerLevelInDBm uintptr
	Put_PreferredTransmitPowerLevelInDBm uintptr
	Get_UseExtendedFormat                uintptr
	Put_UseExtendedFormat                uintptr
	Get_IsAnonymous                      uintptr
	Put_IsAnonymous                      uintptr
	Get_IncludeTransmitPowerLevel        uintptr
	Put_IncludeTransmitPowerLevel        uintptr
}

type IBluetoothLEAdvertisementPublisherTrigger2 struct {
	win32.IInspectable
}

func (this *IBluetoothLEAdvertisementPublisherTrigger2) Vtbl() *IBluetoothLEAdvertisementPublisherTrigger2Vtbl {
	return (*IBluetoothLEAdvertisementPublisherTrigger2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAdvertisementPublisherTrigger2) Get_PreferredTransmitPowerLevelInDBm() *IReference[int16] {
	var _result *IReference[int16]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PreferredTransmitPowerLevelInDBm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEAdvertisementPublisherTrigger2) Put_PreferredTransmitPowerLevelInDBm(value *IReference[int16]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PreferredTransmitPowerLevelInDBm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IBluetoothLEAdvertisementPublisherTrigger2) Get_UseExtendedFormat() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UseExtendedFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementPublisherTrigger2) Put_UseExtendedFormat(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_UseExtendedFormat, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IBluetoothLEAdvertisementPublisherTrigger2) Get_IsAnonymous() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsAnonymous, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementPublisherTrigger2) Put_IsAnonymous(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsAnonymous, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IBluetoothLEAdvertisementPublisherTrigger2) Get_IncludeTransmitPowerLevel() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IncludeTransmitPowerLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementPublisherTrigger2) Put_IncludeTransmitPowerLevel(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IncludeTransmitPowerLevel, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 1AAB1819-BCE1-48EB-A827-59FB7CEE52A6
var IID_IBluetoothLEAdvertisementWatcherTrigger = syscall.GUID{0x1AAB1819, 0xBCE1, 0x48EB,
	[8]byte{0xA8, 0x27, 0x59, 0xFB, 0x7C, 0xEE, 0x52, 0xA6}}

type IBluetoothLEAdvertisementWatcherTriggerInterface interface {
	win32.IInspectableInterface
	Get_MinSamplingInterval() TimeSpan
	Get_MaxSamplingInterval() TimeSpan
	Get_MinOutOfRangeTimeout() TimeSpan
	Get_MaxOutOfRangeTimeout() TimeSpan
	Get_SignalStrengthFilter() *IBluetoothSignalStrengthFilter
	Put_SignalStrengthFilter(value *IBluetoothSignalStrengthFilter)
	Get_AdvertisementFilter() *IBluetoothLEAdvertisementFilter
	Put_AdvertisementFilter(value *IBluetoothLEAdvertisementFilter)
}

type IBluetoothLEAdvertisementWatcherTriggerVtbl struct {
	win32.IInspectableVtbl
	Get_MinSamplingInterval  uintptr
	Get_MaxSamplingInterval  uintptr
	Get_MinOutOfRangeTimeout uintptr
	Get_MaxOutOfRangeTimeout uintptr
	Get_SignalStrengthFilter uintptr
	Put_SignalStrengthFilter uintptr
	Get_AdvertisementFilter  uintptr
	Put_AdvertisementFilter  uintptr
}

type IBluetoothLEAdvertisementWatcherTrigger struct {
	win32.IInspectable
}

func (this *IBluetoothLEAdvertisementWatcherTrigger) Vtbl() *IBluetoothLEAdvertisementWatcherTriggerVtbl {
	return (*IBluetoothLEAdvertisementWatcherTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAdvertisementWatcherTrigger) Get_MinSamplingInterval() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinSamplingInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementWatcherTrigger) Get_MaxSamplingInterval() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxSamplingInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementWatcherTrigger) Get_MinOutOfRangeTimeout() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinOutOfRangeTimeout, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementWatcherTrigger) Get_MaxOutOfRangeTimeout() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxOutOfRangeTimeout, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementWatcherTrigger) Get_SignalStrengthFilter() *IBluetoothSignalStrengthFilter {
	var _result *IBluetoothSignalStrengthFilter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SignalStrengthFilter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEAdvertisementWatcherTrigger) Put_SignalStrengthFilter(value *IBluetoothSignalStrengthFilter) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SignalStrengthFilter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IBluetoothLEAdvertisementWatcherTrigger) Get_AdvertisementFilter() *IBluetoothLEAdvertisementFilter {
	var _result *IBluetoothLEAdvertisementFilter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AdvertisementFilter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBluetoothLEAdvertisementWatcherTrigger) Put_AdvertisementFilter(value *IBluetoothLEAdvertisementFilter) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AdvertisementFilter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 39B56799-EB39-5AB6-9932-AA9E4549604D
var IID_IBluetoothLEAdvertisementWatcherTrigger2 = syscall.GUID{0x39B56799, 0xEB39, 0x5AB6,
	[8]byte{0x99, 0x32, 0xAA, 0x9E, 0x45, 0x49, 0x60, 0x4D}}

type IBluetoothLEAdvertisementWatcherTrigger2Interface interface {
	win32.IInspectableInterface
	Get_AllowExtendedAdvertisements() bool
	Put_AllowExtendedAdvertisements(value bool)
}

type IBluetoothLEAdvertisementWatcherTrigger2Vtbl struct {
	win32.IInspectableVtbl
	Get_AllowExtendedAdvertisements uintptr
	Put_AllowExtendedAdvertisements uintptr
}

type IBluetoothLEAdvertisementWatcherTrigger2 struct {
	win32.IInspectable
}

func (this *IBluetoothLEAdvertisementWatcherTrigger2) Vtbl() *IBluetoothLEAdvertisementWatcherTrigger2Vtbl {
	return (*IBluetoothLEAdvertisementWatcherTrigger2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBluetoothLEAdvertisementWatcherTrigger2) Get_AllowExtendedAdvertisements() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllowExtendedAdvertisements, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBluetoothLEAdvertisementWatcherTrigger2) Put_AllowExtendedAdvertisements(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AllowExtendedAdvertisements, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// E21CAEEB-32F2-4D31-B553-B9E01BDE37E0
var IID_ICachedFileUpdaterTrigger = syscall.GUID{0xE21CAEEB, 0x32F2, 0x4D31,
	[8]byte{0xB5, 0x53, 0xB9, 0xE0, 0x1B, 0xDE, 0x37, 0xE0}}

type ICachedFileUpdaterTriggerInterface interface {
	win32.IInspectableInterface
}

type ICachedFileUpdaterTriggerVtbl struct {
	win32.IInspectableVtbl
}

type ICachedFileUpdaterTrigger struct {
	win32.IInspectable
}

func (this *ICachedFileUpdaterTrigger) Vtbl() *ICachedFileUpdaterTriggerVtbl {
	return (*ICachedFileUpdaterTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 71838C13-1314-47B4-9597-DC7E248C17CC
var IID_ICachedFileUpdaterTriggerDetails = syscall.GUID{0x71838C13, 0x1314, 0x47B4,
	[8]byte{0x95, 0x97, 0xDC, 0x7E, 0x24, 0x8C, 0x17, 0xCC}}

type ICachedFileUpdaterTriggerDetailsInterface interface {
	win32.IInspectableInterface
	Get_UpdateTarget() CachedFileTarget
	Get_UpdateRequest() *IFileUpdateRequest
	Get_CanRequestUserInput() bool
}

type ICachedFileUpdaterTriggerDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_UpdateTarget        uintptr
	Get_UpdateRequest       uintptr
	Get_CanRequestUserInput uintptr
}

type ICachedFileUpdaterTriggerDetails struct {
	win32.IInspectable
}

func (this *ICachedFileUpdaterTriggerDetails) Vtbl() *ICachedFileUpdaterTriggerDetailsVtbl {
	return (*ICachedFileUpdaterTriggerDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICachedFileUpdaterTriggerDetails) Get_UpdateTarget() CachedFileTarget {
	var _result CachedFileTarget
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UpdateTarget, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICachedFileUpdaterTriggerDetails) Get_UpdateRequest() *IFileUpdateRequest {
	var _result *IFileUpdateRequest
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UpdateRequest, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICachedFileUpdaterTriggerDetails) Get_CanRequestUserInput() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanRequestUserInput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 513B43BF-1D40-5C5D-78F5-C923FEE3739E
var IID_IChatMessageNotificationTrigger = syscall.GUID{0x513B43BF, 0x1D40, 0x5C5D,
	[8]byte{0x78, 0xF5, 0xC9, 0x23, 0xFE, 0xE3, 0x73, 0x9E}}

type IChatMessageNotificationTriggerInterface interface {
	win32.IInspectableInterface
}

type IChatMessageNotificationTriggerVtbl struct {
	win32.IInspectableVtbl
}

type IChatMessageNotificationTrigger struct {
	win32.IInspectable
}

func (this *IChatMessageNotificationTrigger) Vtbl() *IChatMessageNotificationTriggerVtbl {
	return (*IChatMessageNotificationTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 3EA3760E-BAF5-4077-88E9-060CF6F0C6D5
var IID_IChatMessageReceivedNotificationTrigger = syscall.GUID{0x3EA3760E, 0xBAF5, 0x4077,
	[8]byte{0x88, 0xE9, 0x06, 0x0C, 0xF6, 0xF0, 0xC6, 0xD5}}

type IChatMessageReceivedNotificationTriggerInterface interface {
	win32.IInspectableInterface
}

type IChatMessageReceivedNotificationTriggerVtbl struct {
	win32.IInspectableVtbl
}

type IChatMessageReceivedNotificationTrigger struct {
	win32.IInspectable
}

func (this *IChatMessageReceivedNotificationTrigger) Vtbl() *IChatMessageReceivedNotificationTriggerVtbl {
	return (*IChatMessageReceivedNotificationTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// FB91F28A-16A5-486D-974C-7835A8477BE2
var IID_ICommunicationBlockingAppSetAsActiveTrigger = syscall.GUID{0xFB91F28A, 0x16A5, 0x486D,
	[8]byte{0x97, 0x4C, 0x78, 0x35, 0xA8, 0x47, 0x7B, 0xE2}}

type ICommunicationBlockingAppSetAsActiveTriggerInterface interface {
	win32.IInspectableInterface
}

type ICommunicationBlockingAppSetAsActiveTriggerVtbl struct {
	win32.IInspectableVtbl
}

type ICommunicationBlockingAppSetAsActiveTrigger struct {
	win32.IInspectable
}

func (this *ICommunicationBlockingAppSetAsActiveTrigger) Vtbl() *ICommunicationBlockingAppSetAsActiveTriggerVtbl {
	return (*ICommunicationBlockingAppSetAsActiveTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// C833419B-4705-4571-9A16-06B997BF9C96
var IID_IContactStoreNotificationTrigger = syscall.GUID{0xC833419B, 0x4705, 0x4571,
	[8]byte{0x9A, 0x16, 0x06, 0xB9, 0x97, 0xBF, 0x9C, 0x96}}

type IContactStoreNotificationTriggerInterface interface {
	win32.IInspectableInterface
}

type IContactStoreNotificationTriggerVtbl struct {
	win32.IInspectableVtbl
}

type IContactStoreNotificationTrigger struct {
	win32.IInspectable
}

func (this *IContactStoreNotificationTrigger) Vtbl() *IContactStoreNotificationTriggerVtbl {
	return (*IContactStoreNotificationTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 710627EE-04FA-440B-80C0-173202199E5D
var IID_IContentPrefetchTrigger = syscall.GUID{0x710627EE, 0x04FA, 0x440B,
	[8]byte{0x80, 0xC0, 0x17, 0x32, 0x02, 0x19, 0x9E, 0x5D}}

type IContentPrefetchTriggerInterface interface {
	win32.IInspectableInterface
	Get_WaitInterval() TimeSpan
}

type IContentPrefetchTriggerVtbl struct {
	win32.IInspectableVtbl
	Get_WaitInterval uintptr
}

type IContentPrefetchTrigger struct {
	win32.IInspectable
}

func (this *IContentPrefetchTrigger) Vtbl() *IContentPrefetchTriggerVtbl {
	return (*IContentPrefetchTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IContentPrefetchTrigger) Get_WaitInterval() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WaitInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C2643EDA-8A03-409E-B8C4-88814C28CCB6
var IID_IContentPrefetchTriggerFactory = syscall.GUID{0xC2643EDA, 0x8A03, 0x409E,
	[8]byte{0xB8, 0xC4, 0x88, 0x81, 0x4C, 0x28, 0xCC, 0xB6}}

type IContentPrefetchTriggerFactoryInterface interface {
	win32.IInspectableInterface
	Create(waitInterval TimeSpan) *IContentPrefetchTrigger
}

type IContentPrefetchTriggerFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IContentPrefetchTriggerFactory struct {
	win32.IInspectable
}

func (this *IContentPrefetchTriggerFactory) Vtbl() *IContentPrefetchTriggerFactoryVtbl {
	return (*IContentPrefetchTriggerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IContentPrefetchTriggerFactory) Create(waitInterval TimeSpan) *IContentPrefetchTrigger {
	var _result *IContentPrefetchTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&waitInterval)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F3596798-CF6B-4EF4-A0CA-29CF4A278C87
var IID_ICustomSystemEventTrigger = syscall.GUID{0xF3596798, 0xCF6B, 0x4EF4,
	[8]byte{0xA0, 0xCA, 0x29, 0xCF, 0x4A, 0x27, 0x8C, 0x87}}

type ICustomSystemEventTriggerInterface interface {
	win32.IInspectableInterface
	Get_TriggerId() string
	Get_Recurrence() CustomSystemEventTriggerRecurrence
}

type ICustomSystemEventTriggerVtbl struct {
	win32.IInspectableVtbl
	Get_TriggerId  uintptr
	Get_Recurrence uintptr
}

type ICustomSystemEventTrigger struct {
	win32.IInspectable
}

func (this *ICustomSystemEventTrigger) Vtbl() *ICustomSystemEventTriggerVtbl {
	return (*ICustomSystemEventTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICustomSystemEventTrigger) Get_TriggerId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TriggerId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICustomSystemEventTrigger) Get_Recurrence() CustomSystemEventTriggerRecurrence {
	var _result CustomSystemEventTriggerRecurrence
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Recurrence, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 6BCB16C5-F2DC-41B2-9EFD-B96BDCD13CED
var IID_ICustomSystemEventTriggerFactory = syscall.GUID{0x6BCB16C5, 0xF2DC, 0x41B2,
	[8]byte{0x9E, 0xFD, 0xB9, 0x6B, 0xDC, 0xD1, 0x3C, 0xED}}

type ICustomSystemEventTriggerFactoryInterface interface {
	win32.IInspectableInterface
	Create(triggerId string, recurrence CustomSystemEventTriggerRecurrence) *ICustomSystemEventTrigger
}

type ICustomSystemEventTriggerFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ICustomSystemEventTriggerFactory struct {
	win32.IInspectable
}

func (this *ICustomSystemEventTriggerFactory) Vtbl() *ICustomSystemEventTriggerFactoryVtbl {
	return (*ICustomSystemEventTriggerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICustomSystemEventTriggerFactory) Create(triggerId string, recurrence CustomSystemEventTriggerRecurrence) *ICustomSystemEventTrigger {
	var _result *ICustomSystemEventTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(triggerId).Ptr, uintptr(recurrence), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 90875E64-3CDD-4EFB-AB1C-5B3B6A60CE34
var IID_IDeviceConnectionChangeTrigger = syscall.GUID{0x90875E64, 0x3CDD, 0x4EFB,
	[8]byte{0xAB, 0x1C, 0x5B, 0x3B, 0x6A, 0x60, 0xCE, 0x34}}

type IDeviceConnectionChangeTriggerInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_CanMaintainConnection() bool
	Get_MaintainConnection() bool
	Put_MaintainConnection(value bool)
}

type IDeviceConnectionChangeTriggerVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId              uintptr
	Get_CanMaintainConnection uintptr
	Get_MaintainConnection    uintptr
	Put_MaintainConnection    uintptr
}

type IDeviceConnectionChangeTrigger struct {
	win32.IInspectable
}

func (this *IDeviceConnectionChangeTrigger) Vtbl() *IDeviceConnectionChangeTriggerVtbl {
	return (*IDeviceConnectionChangeTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceConnectionChangeTrigger) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDeviceConnectionChangeTrigger) Get_CanMaintainConnection() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanMaintainConnection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDeviceConnectionChangeTrigger) Get_MaintainConnection() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaintainConnection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDeviceConnectionChangeTrigger) Put_MaintainConnection(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaintainConnection, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// C3EA246A-4EFD-4498-AA60-A4E4E3B17AB9
var IID_IDeviceConnectionChangeTriggerStatics = syscall.GUID{0xC3EA246A, 0x4EFD, 0x4498,
	[8]byte{0xAA, 0x60, 0xA4, 0xE4, 0xE3, 0xB1, 0x7A, 0xB9}}

type IDeviceConnectionChangeTriggerStaticsInterface interface {
	win32.IInspectableInterface
	FromIdAsync(deviceId string) *IAsyncOperation[*IDeviceConnectionChangeTrigger]
}

type IDeviceConnectionChangeTriggerStaticsVtbl struct {
	win32.IInspectableVtbl
	FromIdAsync uintptr
}

type IDeviceConnectionChangeTriggerStatics struct {
	win32.IInspectable
}

func (this *IDeviceConnectionChangeTriggerStatics) Vtbl() *IDeviceConnectionChangeTriggerStaticsVtbl {
	return (*IDeviceConnectionChangeTriggerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceConnectionChangeTriggerStatics) FromIdAsync(deviceId string) *IAsyncOperation[*IDeviceConnectionChangeTrigger] {
	var _result *IAsyncOperation[*IDeviceConnectionChangeTrigger]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 81278AB5-41AB-16DA-86C2-7F7BF0912F5B
var IID_IDeviceManufacturerNotificationTrigger = syscall.GUID{0x81278AB5, 0x41AB, 0x16DA,
	[8]byte{0x86, 0xC2, 0x7F, 0x7B, 0xF0, 0x91, 0x2F, 0x5B}}

type IDeviceManufacturerNotificationTriggerInterface interface {
	win32.IInspectableInterface
	Get_TriggerQualifier() string
	Get_OneShot() bool
}

type IDeviceManufacturerNotificationTriggerVtbl struct {
	win32.IInspectableVtbl
	Get_TriggerQualifier uintptr
	Get_OneShot          uintptr
}

type IDeviceManufacturerNotificationTrigger struct {
	win32.IInspectable
}

func (this *IDeviceManufacturerNotificationTrigger) Vtbl() *IDeviceManufacturerNotificationTriggerVtbl {
	return (*IDeviceManufacturerNotificationTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceManufacturerNotificationTrigger) Get_TriggerQualifier() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TriggerQualifier, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDeviceManufacturerNotificationTrigger) Get_OneShot() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OneShot, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7955DE75-25BB-4153-A1A2-3029FCABB652
var IID_IDeviceManufacturerNotificationTriggerFactory = syscall.GUID{0x7955DE75, 0x25BB, 0x4153,
	[8]byte{0xA1, 0xA2, 0x30, 0x29, 0xFC, 0xAB, 0xB6, 0x52}}

type IDeviceManufacturerNotificationTriggerFactoryInterface interface {
	win32.IInspectableInterface
	Create(triggerQualifier string, oneShot bool) *IDeviceManufacturerNotificationTrigger
}

type IDeviceManufacturerNotificationTriggerFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IDeviceManufacturerNotificationTriggerFactory struct {
	win32.IInspectable
}

func (this *IDeviceManufacturerNotificationTriggerFactory) Vtbl() *IDeviceManufacturerNotificationTriggerFactoryVtbl {
	return (*IDeviceManufacturerNotificationTriggerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceManufacturerNotificationTriggerFactory) Create(triggerQualifier string, oneShot bool) *IDeviceManufacturerNotificationTrigger {
	var _result *IDeviceManufacturerNotificationTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(triggerQualifier).Ptr, uintptr(*(*byte)(unsafe.Pointer(&oneShot))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1AB217AD-6E34-49D3-9E6F-17F1B6DFA881
var IID_IDeviceServicingTrigger = syscall.GUID{0x1AB217AD, 0x6E34, 0x49D3,
	[8]byte{0x9E, 0x6F, 0x17, 0xF1, 0xB6, 0xDF, 0xA8, 0x81}}

type IDeviceServicingTriggerInterface interface {
	win32.IInspectableInterface
	RequestAsyncSimple(deviceId string, expectedDuration TimeSpan) *IAsyncOperation[DeviceTriggerResult]
	RequestAsyncWithArguments(deviceId string, expectedDuration TimeSpan, arguments string) *IAsyncOperation[DeviceTriggerResult]
}

type IDeviceServicingTriggerVtbl struct {
	win32.IInspectableVtbl
	RequestAsyncSimple        uintptr
	RequestAsyncWithArguments uintptr
}

type IDeviceServicingTrigger struct {
	win32.IInspectable
}

func (this *IDeviceServicingTrigger) Vtbl() *IDeviceServicingTriggerVtbl {
	return (*IDeviceServicingTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceServicingTrigger) RequestAsyncSimple(deviceId string, expectedDuration TimeSpan) *IAsyncOperation[DeviceTriggerResult] {
	var _result *IAsyncOperation[DeviceTriggerResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAsyncSimple, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, *(*uintptr)(unsafe.Pointer(&expectedDuration)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDeviceServicingTrigger) RequestAsyncWithArguments(deviceId string, expectedDuration TimeSpan, arguments string) *IAsyncOperation[DeviceTriggerResult] {
	var _result *IAsyncOperation[DeviceTriggerResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAsyncWithArguments, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, *(*uintptr)(unsafe.Pointer(&expectedDuration)), NewHStr(arguments).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0DA68011-334F-4D57-B6EC-6DCA64B412E4
var IID_IDeviceUseTrigger = syscall.GUID{0x0DA68011, 0x334F, 0x4D57,
	[8]byte{0xB6, 0xEC, 0x6D, 0xCA, 0x64, 0xB4, 0x12, 0xE4}}

type IDeviceUseTriggerInterface interface {
	win32.IInspectableInterface
	RequestAsyncSimple(deviceId string) *IAsyncOperation[DeviceTriggerResult]
	RequestAsyncWithArguments(deviceId string, arguments string) *IAsyncOperation[DeviceTriggerResult]
}

type IDeviceUseTriggerVtbl struct {
	win32.IInspectableVtbl
	RequestAsyncSimple        uintptr
	RequestAsyncWithArguments uintptr
}

type IDeviceUseTrigger struct {
	win32.IInspectable
}

func (this *IDeviceUseTrigger) Vtbl() *IDeviceUseTriggerVtbl {
	return (*IDeviceUseTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceUseTrigger) RequestAsyncSimple(deviceId string) *IAsyncOperation[DeviceTriggerResult] {
	var _result *IAsyncOperation[DeviceTriggerResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAsyncSimple, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDeviceUseTrigger) RequestAsyncWithArguments(deviceId string, arguments string) *IAsyncOperation[DeviceTriggerResult] {
	var _result *IAsyncOperation[DeviceTriggerResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAsyncWithArguments, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, NewHStr(arguments).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A4617FDD-8573-4260-BEFC-5BEC89CB693D
var IID_IDeviceWatcherTrigger = syscall.GUID{0xA4617FDD, 0x8573, 0x4260,
	[8]byte{0xBE, 0xFC, 0x5B, 0xEC, 0x89, 0xCB, 0x69, 0x3D}}

type IDeviceWatcherTriggerInterface interface {
	win32.IInspectableInterface
}

type IDeviceWatcherTriggerVtbl struct {
	win32.IInspectableVtbl
}

type IDeviceWatcherTrigger struct {
	win32.IInspectable
}

func (this *IDeviceWatcherTrigger) Vtbl() *IDeviceWatcherTriggerVtbl {
	return (*IDeviceWatcherTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 986D06DA-47EB-4268-A4F2-F3F77188388A
var IID_IEmailStoreNotificationTrigger = syscall.GUID{0x986D06DA, 0x47EB, 0x4268,
	[8]byte{0xA4, 0xF2, 0xF3, 0xF7, 0x71, 0x88, 0x38, 0x8A}}

type IEmailStoreNotificationTriggerInterface interface {
	win32.IInspectableInterface
}

type IEmailStoreNotificationTriggerVtbl struct {
	win32.IInspectableVtbl
}

type IEmailStoreNotificationTrigger struct {
	win32.IInspectable
}

func (this *IEmailStoreNotificationTrigger) Vtbl() *IEmailStoreNotificationTriggerVtbl {
	return (*IEmailStoreNotificationTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// E25F8FC8-0696-474F-A732-F292B0CEBC5D
var IID_IGattCharacteristicNotificationTrigger = syscall.GUID{0xE25F8FC8, 0x0696, 0x474F,
	[8]byte{0xA7, 0x32, 0xF2, 0x92, 0xB0, 0xCE, 0xBC, 0x5D}}

type IGattCharacteristicNotificationTriggerInterface interface {
	win32.IInspectableInterface
	Get_Characteristic() *IGattCharacteristic
}

type IGattCharacteristicNotificationTriggerVtbl struct {
	win32.IInspectableVtbl
	Get_Characteristic uintptr
}

type IGattCharacteristicNotificationTrigger struct {
	win32.IInspectable
}

func (this *IGattCharacteristicNotificationTrigger) Vtbl() *IGattCharacteristicNotificationTriggerVtbl {
	return (*IGattCharacteristicNotificationTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattCharacteristicNotificationTrigger) Get_Characteristic() *IGattCharacteristic {
	var _result *IGattCharacteristic
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Characteristic, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9322A2C4-AE0E-42F2-B28C-F51372E69245
var IID_IGattCharacteristicNotificationTrigger2 = syscall.GUID{0x9322A2C4, 0xAE0E, 0x42F2,
	[8]byte{0xB2, 0x8C, 0xF5, 0x13, 0x72, 0xE6, 0x92, 0x45}}

type IGattCharacteristicNotificationTrigger2Interface interface {
	win32.IInspectableInterface
	Get_EventTriggeringMode() BluetoothEventTriggeringMode
}

type IGattCharacteristicNotificationTrigger2Vtbl struct {
	win32.IInspectableVtbl
	Get_EventTriggeringMode uintptr
}

type IGattCharacteristicNotificationTrigger2 struct {
	win32.IInspectable
}

func (this *IGattCharacteristicNotificationTrigger2) Vtbl() *IGattCharacteristicNotificationTrigger2Vtbl {
	return (*IGattCharacteristicNotificationTrigger2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattCharacteristicNotificationTrigger2) Get_EventTriggeringMode() BluetoothEventTriggeringMode {
	var _result BluetoothEventTriggeringMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EventTriggeringMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 57BA1995-B143-4575-9F6B-FD59D93ACE1A
var IID_IGattCharacteristicNotificationTriggerFactory = syscall.GUID{0x57BA1995, 0xB143, 0x4575,
	[8]byte{0x9F, 0x6B, 0xFD, 0x59, 0xD9, 0x3A, 0xCE, 0x1A}}

type IGattCharacteristicNotificationTriggerFactoryInterface interface {
	win32.IInspectableInterface
	Create(characteristic *IGattCharacteristic) *IGattCharacteristicNotificationTrigger
}

type IGattCharacteristicNotificationTriggerFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IGattCharacteristicNotificationTriggerFactory struct {
	win32.IInspectable
}

func (this *IGattCharacteristicNotificationTriggerFactory) Vtbl() *IGattCharacteristicNotificationTriggerFactoryVtbl {
	return (*IGattCharacteristicNotificationTriggerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattCharacteristicNotificationTriggerFactory) Create(characteristic *IGattCharacteristic) *IGattCharacteristicNotificationTrigger {
	var _result *IGattCharacteristicNotificationTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(characteristic)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5998E91F-8A53-4E9F-A32C-23CD33664CEE
var IID_IGattCharacteristicNotificationTriggerFactory2 = syscall.GUID{0x5998E91F, 0x8A53, 0x4E9F,
	[8]byte{0xA3, 0x2C, 0x23, 0xCD, 0x33, 0x66, 0x4C, 0xEE}}

type IGattCharacteristicNotificationTriggerFactory2Interface interface {
	win32.IInspectableInterface
	CreateWithEventTriggeringMode(characteristic *IGattCharacteristic, eventTriggeringMode BluetoothEventTriggeringMode) *IGattCharacteristicNotificationTrigger
}

type IGattCharacteristicNotificationTriggerFactory2Vtbl struct {
	win32.IInspectableVtbl
	CreateWithEventTriggeringMode uintptr
}

type IGattCharacteristicNotificationTriggerFactory2 struct {
	win32.IInspectable
}

func (this *IGattCharacteristicNotificationTriggerFactory2) Vtbl() *IGattCharacteristicNotificationTriggerFactory2Vtbl {
	return (*IGattCharacteristicNotificationTriggerFactory2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattCharacteristicNotificationTriggerFactory2) CreateWithEventTriggeringMode(characteristic *IGattCharacteristic, eventTriggeringMode BluetoothEventTriggeringMode) *IGattCharacteristicNotificationTrigger {
	var _result *IGattCharacteristicNotificationTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithEventTriggeringMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(characteristic)), uintptr(eventTriggeringMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DDC6A3E9-1557-4BD8-8542-468AA0C696F6
var IID_IGattServiceProviderTrigger = syscall.GUID{0xDDC6A3E9, 0x1557, 0x4BD8,
	[8]byte{0x85, 0x42, 0x46, 0x8A, 0xA0, 0xC6, 0x96, 0xF6}}

type IGattServiceProviderTriggerInterface interface {
	win32.IInspectableInterface
	Get_TriggerId() string
	Get_Service() *IGattLocalService
	Put_AdvertisingParameters(value *IGattServiceProviderAdvertisingParameters)
	Get_AdvertisingParameters() *IGattServiceProviderAdvertisingParameters
}

type IGattServiceProviderTriggerVtbl struct {
	win32.IInspectableVtbl
	Get_TriggerId             uintptr
	Get_Service               uintptr
	Put_AdvertisingParameters uintptr
	Get_AdvertisingParameters uintptr
}

type IGattServiceProviderTrigger struct {
	win32.IInspectable
}

func (this *IGattServiceProviderTrigger) Vtbl() *IGattServiceProviderTriggerVtbl {
	return (*IGattServiceProviderTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattServiceProviderTrigger) Get_TriggerId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TriggerId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGattServiceProviderTrigger) Get_Service() *IGattLocalService {
	var _result *IGattLocalService
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Service, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattServiceProviderTrigger) Put_AdvertisingParameters(value *IGattServiceProviderAdvertisingParameters) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AdvertisingParameters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IGattServiceProviderTrigger) Get_AdvertisingParameters() *IGattServiceProviderAdvertisingParameters {
	var _result *IGattServiceProviderAdvertisingParameters
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AdvertisingParameters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3C4691B1-B198-4E84-BAD4-CF4AD299ED3A
var IID_IGattServiceProviderTriggerResult = syscall.GUID{0x3C4691B1, 0xB198, 0x4E84,
	[8]byte{0xBA, 0xD4, 0xCF, 0x4A, 0xD2, 0x99, 0xED, 0x3A}}

type IGattServiceProviderTriggerResultInterface interface {
	win32.IInspectableInterface
	Get_Trigger() *IGattServiceProviderTrigger
	Get_Error() BluetoothError
}

type IGattServiceProviderTriggerResultVtbl struct {
	win32.IInspectableVtbl
	Get_Trigger uintptr
	Get_Error   uintptr
}

type IGattServiceProviderTriggerResult struct {
	win32.IInspectable
}

func (this *IGattServiceProviderTriggerResult) Vtbl() *IGattServiceProviderTriggerResultVtbl {
	return (*IGattServiceProviderTriggerResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattServiceProviderTriggerResult) Get_Trigger() *IGattServiceProviderTrigger {
	var _result *IGattServiceProviderTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Trigger, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGattServiceProviderTriggerResult) Get_Error() BluetoothError {
	var _result BluetoothError
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Error, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// B413A36A-E294-4591-A5A6-64891A828153
var IID_IGattServiceProviderTriggerStatics = syscall.GUID{0xB413A36A, 0xE294, 0x4591,
	[8]byte{0xA5, 0xA6, 0x64, 0x89, 0x1A, 0x82, 0x81, 0x53}}

type IGattServiceProviderTriggerStaticsInterface interface {
	win32.IInspectableInterface
	CreateAsync(triggerId string, serviceUuid syscall.GUID) *IAsyncOperation[*IGattServiceProviderTriggerResult]
}

type IGattServiceProviderTriggerStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateAsync uintptr
}

type IGattServiceProviderTriggerStatics struct {
	win32.IInspectable
}

func (this *IGattServiceProviderTriggerStatics) Vtbl() *IGattServiceProviderTriggerStaticsVtbl {
	return (*IGattServiceProviderTriggerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGattServiceProviderTriggerStatics) CreateAsync(triggerId string, serviceUuid syscall.GUID) *IAsyncOperation[*IGattServiceProviderTriggerResult] {
	var _result *IAsyncOperation[*IGattServiceProviderTriggerResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAsync, uintptr(unsafe.Pointer(this)), NewHStr(triggerId).Ptr, uintptr(unsafe.Pointer(&serviceUuid)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4818EDAA-04E1-4127-9A4C-19351B8A80A4
var IID_IGeovisitTrigger = syscall.GUID{0x4818EDAA, 0x04E1, 0x4127,
	[8]byte{0x9A, 0x4C, 0x19, 0x35, 0x1B, 0x8A, 0x80, 0xA4}}

type IGeovisitTriggerInterface interface {
	win32.IInspectableInterface
	Get_MonitoringScope() VisitMonitoringScope
	Put_MonitoringScope(value VisitMonitoringScope)
}

type IGeovisitTriggerVtbl struct {
	win32.IInspectableVtbl
	Get_MonitoringScope uintptr
	Put_MonitoringScope uintptr
}

type IGeovisitTrigger struct {
	win32.IInspectable
}

func (this *IGeovisitTrigger) Vtbl() *IGeovisitTriggerVtbl {
	return (*IGeovisitTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeovisitTrigger) Get_MonitoringScope() VisitMonitoringScope {
	var _result VisitMonitoringScope
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MonitoringScope, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGeovisitTrigger) Put_MonitoringScope(value VisitMonitoringScope) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MonitoringScope, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 47666A1C-6877-481E-8026-FF7E14A811A0
var IID_ILocationTrigger = syscall.GUID{0x47666A1C, 0x6877, 0x481E,
	[8]byte{0x80, 0x26, 0xFF, 0x7E, 0x14, 0xA8, 0x11, 0xA0}}

type ILocationTriggerInterface interface {
	win32.IInspectableInterface
	Get_TriggerType() LocationTriggerType
}

type ILocationTriggerVtbl struct {
	win32.IInspectableVtbl
	Get_TriggerType uintptr
}

type ILocationTrigger struct {
	win32.IInspectable
}

func (this *ILocationTrigger) Vtbl() *ILocationTriggerVtbl {
	return (*ILocationTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILocationTrigger) Get_TriggerType() LocationTriggerType {
	var _result LocationTriggerType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TriggerType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 1106BB07-FF69-4E09-AA8B-1384EA475E98
var IID_ILocationTriggerFactory = syscall.GUID{0x1106BB07, 0xFF69, 0x4E09,
	[8]byte{0xAA, 0x8B, 0x13, 0x84, 0xEA, 0x47, 0x5E, 0x98}}

type ILocationTriggerFactoryInterface interface {
	win32.IInspectableInterface
	Create(triggerType LocationTriggerType) *ILocationTrigger
}

type ILocationTriggerFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ILocationTriggerFactory struct {
	win32.IInspectable
}

func (this *ILocationTriggerFactory) Vtbl() *ILocationTriggerFactoryVtbl {
	return (*ILocationTriggerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILocationTriggerFactory) Create(triggerType LocationTriggerType) *ILocationTrigger {
	var _result *ILocationTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(triggerType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 68184C83-FC22-4CE5-841A-7239A9810047
var IID_IMaintenanceTrigger = syscall.GUID{0x68184C83, 0xFC22, 0x4CE5,
	[8]byte{0x84, 0x1A, 0x72, 0x39, 0xA9, 0x81, 0x00, 0x47}}

type IMaintenanceTriggerInterface interface {
	win32.IInspectableInterface
	Get_FreshnessTime() uint32
	Get_OneShot() bool
}

type IMaintenanceTriggerVtbl struct {
	win32.IInspectableVtbl
	Get_FreshnessTime uintptr
	Get_OneShot       uintptr
}

type IMaintenanceTrigger struct {
	win32.IInspectable
}

func (this *IMaintenanceTrigger) Vtbl() *IMaintenanceTriggerVtbl {
	return (*IMaintenanceTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMaintenanceTrigger) Get_FreshnessTime() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FreshnessTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMaintenanceTrigger) Get_OneShot() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OneShot, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4B3DDB2E-97DD-4629-88B0-B06CF9482AE5
var IID_IMaintenanceTriggerFactory = syscall.GUID{0x4B3DDB2E, 0x97DD, 0x4629,
	[8]byte{0x88, 0xB0, 0xB0, 0x6C, 0xF9, 0x48, 0x2A, 0xE5}}

type IMaintenanceTriggerFactoryInterface interface {
	win32.IInspectableInterface
	Create(freshnessTime uint32, oneShot bool) *IMaintenanceTrigger
}

type IMaintenanceTriggerFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IMaintenanceTriggerFactory struct {
	win32.IInspectable
}

func (this *IMaintenanceTriggerFactory) Vtbl() *IMaintenanceTriggerFactoryVtbl {
	return (*IMaintenanceTriggerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMaintenanceTriggerFactory) Create(freshnessTime uint32, oneShot bool) *IMaintenanceTrigger {
	var _result *IMaintenanceTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(freshnessTime), uintptr(*(*byte)(unsafe.Pointer(&oneShot))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9A95BE65-8A52-4B30-9011-CF38040EA8B0
var IID_IMediaProcessingTrigger = syscall.GUID{0x9A95BE65, 0x8A52, 0x4B30,
	[8]byte{0x90, 0x11, 0xCF, 0x38, 0x04, 0x0E, 0xA8, 0xB0}}

type IMediaProcessingTriggerInterface interface {
	win32.IInspectableInterface
	RequestAsync() *IAsyncOperation[MediaProcessingTriggerResult]
	RequestAsyncWithArguments(arguments *IPropertySet) *IAsyncOperation[MediaProcessingTriggerResult]
}

type IMediaProcessingTriggerVtbl struct {
	win32.IInspectableVtbl
	RequestAsync              uintptr
	RequestAsyncWithArguments uintptr
}

type IMediaProcessingTrigger struct {
	win32.IInspectable
}

func (this *IMediaProcessingTrigger) Vtbl() *IMediaProcessingTriggerVtbl {
	return (*IMediaProcessingTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaProcessingTrigger) RequestAsync() *IAsyncOperation[MediaProcessingTriggerResult] {
	var _result *IAsyncOperation[MediaProcessingTriggerResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaProcessingTrigger) RequestAsyncWithArguments(arguments *IPropertySet) *IAsyncOperation[MediaProcessingTriggerResult] {
	var _result *IAsyncOperation[MediaProcessingTriggerResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAsyncWithArguments, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(arguments)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E756C791-3001-4DE5-83C7-DE61D88831D0
var IID_INetworkOperatorHotspotAuthenticationTrigger = syscall.GUID{0xE756C791, 0x3001, 0x4DE5,
	[8]byte{0x83, 0xC7, 0xDE, 0x61, 0xD8, 0x88, 0x31, 0xD0}}

type INetworkOperatorHotspotAuthenticationTriggerInterface interface {
	win32.IInspectableInterface
}

type INetworkOperatorHotspotAuthenticationTriggerVtbl struct {
	win32.IInspectableVtbl
}

type INetworkOperatorHotspotAuthenticationTrigger struct {
	win32.IInspectable
}

func (this *INetworkOperatorHotspotAuthenticationTrigger) Vtbl() *INetworkOperatorHotspotAuthenticationTriggerVtbl {
	return (*INetworkOperatorHotspotAuthenticationTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 90089CC6-63CD-480C-95D1-6E6AEF801E4A
var IID_INetworkOperatorNotificationTrigger = syscall.GUID{0x90089CC6, 0x63CD, 0x480C,
	[8]byte{0x95, 0xD1, 0x6E, 0x6A, 0xEF, 0x80, 0x1E, 0x4A}}

type INetworkOperatorNotificationTriggerInterface interface {
	win32.IInspectableInterface
	Get_NetworkAccountId() string
}

type INetworkOperatorNotificationTriggerVtbl struct {
	win32.IInspectableVtbl
	Get_NetworkAccountId uintptr
}

type INetworkOperatorNotificationTrigger struct {
	win32.IInspectable
}

func (this *INetworkOperatorNotificationTrigger) Vtbl() *INetworkOperatorNotificationTriggerVtbl {
	return (*INetworkOperatorNotificationTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INetworkOperatorNotificationTrigger) Get_NetworkAccountId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkAccountId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 0A223E00-27D7-4353-ADB9-9265AAEA579D
var IID_INetworkOperatorNotificationTriggerFactory = syscall.GUID{0x0A223E00, 0x27D7, 0x4353,
	[8]byte{0xAD, 0xB9, 0x92, 0x65, 0xAA, 0xEA, 0x57, 0x9D}}

type INetworkOperatorNotificationTriggerFactoryInterface interface {
	win32.IInspectableInterface
	Create(networkAccountId string) *INetworkOperatorNotificationTrigger
}

type INetworkOperatorNotificationTriggerFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type INetworkOperatorNotificationTriggerFactory struct {
	win32.IInspectable
}

func (this *INetworkOperatorNotificationTriggerFactory) Vtbl() *INetworkOperatorNotificationTriggerFactoryVtbl {
	return (*INetworkOperatorNotificationTriggerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INetworkOperatorNotificationTriggerFactory) Create(networkAccountId string) *INetworkOperatorNotificationTrigger {
	var _result *INetworkOperatorNotificationTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(networkAccountId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8DCFE99B-D4C5-49F1-B7D3-82E87A0E9DDE
var IID_IPhoneTrigger = syscall.GUID{0x8DCFE99B, 0xD4C5, 0x49F1,
	[8]byte{0xB7, 0xD3, 0x82, 0xE8, 0x7A, 0x0E, 0x9D, 0xDE}}

type IPhoneTriggerInterface interface {
	win32.IInspectableInterface
	Get_OneShot() bool
	Get_TriggerType() unsafe.Pointer
}

type IPhoneTriggerVtbl struct {
	win32.IInspectableVtbl
	Get_OneShot     uintptr
	Get_TriggerType uintptr
}

type IPhoneTrigger struct {
	win32.IInspectable
}

func (this *IPhoneTrigger) Vtbl() *IPhoneTriggerVtbl {
	return (*IPhoneTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPhoneTrigger) Get_OneShot() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OneShot, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhoneTrigger) Get_TriggerType() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TriggerType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A0D93CDA-5FC1-48FB-A546-32262040157B
var IID_IPhoneTriggerFactory = syscall.GUID{0xA0D93CDA, 0x5FC1, 0x48FB,
	[8]byte{0xA5, 0x46, 0x32, 0x26, 0x20, 0x40, 0x15, 0x7B}}

type IPhoneTriggerFactoryInterface interface {
	win32.IInspectableInterface
	Create(type_ unsafe.Pointer, oneShot bool) *IPhoneTrigger
}

type IPhoneTriggerFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IPhoneTriggerFactory struct {
	win32.IInspectable
}

func (this *IPhoneTriggerFactory) Vtbl() *IPhoneTriggerFactoryVtbl {
	return (*IPhoneTriggerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPhoneTriggerFactory) Create(type_ unsafe.Pointer, oneShot bool) *IPhoneTrigger {
	var _result *IPhoneTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(type_), uintptr(*(*byte)(unsafe.Pointer(&oneShot))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6DD8ED1B-458E-4FC2-BC2E-D5664F77ED19
var IID_IPushNotificationTriggerFactory = syscall.GUID{0x6DD8ED1B, 0x458E, 0x4FC2,
	[8]byte{0xBC, 0x2E, 0xD5, 0x66, 0x4F, 0x77, 0xED, 0x19}}

type IPushNotificationTriggerFactoryInterface interface {
	win32.IInspectableInterface
	Create(applicationId string) *IBackgroundTrigger
}

type IPushNotificationTriggerFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IPushNotificationTriggerFactory struct {
	win32.IInspectable
}

func (this *IPushNotificationTriggerFactory) Vtbl() *IPushNotificationTriggerFactoryVtbl {
	return (*IPushNotificationTriggerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPushNotificationTriggerFactory) Create(applicationId string) *IBackgroundTrigger {
	var _result *IBackgroundTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(applicationId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 986D0D6A-B2F6-467F-A978-A44091C11A66
var IID_IRcsEndUserMessageAvailableTrigger = syscall.GUID{0x986D0D6A, 0xB2F6, 0x467F,
	[8]byte{0xA9, 0x78, 0xA4, 0x40, 0x91, 0xC1, 0x1A, 0x66}}

type IRcsEndUserMessageAvailableTriggerInterface interface {
	win32.IInspectableInterface
}

type IRcsEndUserMessageAvailableTriggerVtbl struct {
	win32.IInspectableVtbl
}

type IRcsEndUserMessageAvailableTrigger struct {
	win32.IInspectable
}

func (this *IRcsEndUserMessageAvailableTrigger) Vtbl() *IRcsEndUserMessageAvailableTriggerVtbl {
	return (*IRcsEndUserMessageAvailableTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// E8C4CAE2-0B53-4464-9394-FD875654DE64
var IID_IRfcommConnectionTrigger = syscall.GUID{0xE8C4CAE2, 0x0B53, 0x4464,
	[8]byte{0x93, 0x94, 0xFD, 0x87, 0x56, 0x54, 0xDE, 0x64}}

type IRfcommConnectionTriggerInterface interface {
	win32.IInspectableInterface
	Get_InboundConnection() *IRfcommInboundConnectionInformation
	Get_OutboundConnection() *IRfcommOutboundConnectionInformation
	Get_AllowMultipleConnections() bool
	Put_AllowMultipleConnections(value bool)
	Get_ProtectionLevel() SocketProtectionLevel
	Put_ProtectionLevel(value SocketProtectionLevel)
	Get_RemoteHostName() *IHostName
	Put_RemoteHostName(value *IHostName)
}

type IRfcommConnectionTriggerVtbl struct {
	win32.IInspectableVtbl
	Get_InboundConnection        uintptr
	Get_OutboundConnection       uintptr
	Get_AllowMultipleConnections uintptr
	Put_AllowMultipleConnections uintptr
	Get_ProtectionLevel          uintptr
	Put_ProtectionLevel          uintptr
	Get_RemoteHostName           uintptr
	Put_RemoteHostName           uintptr
}

type IRfcommConnectionTrigger struct {
	win32.IInspectable
}

func (this *IRfcommConnectionTrigger) Vtbl() *IRfcommConnectionTriggerVtbl {
	return (*IRfcommConnectionTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRfcommConnectionTrigger) Get_InboundConnection() *IRfcommInboundConnectionInformation {
	var _result *IRfcommInboundConnectionInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InboundConnection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRfcommConnectionTrigger) Get_OutboundConnection() *IRfcommOutboundConnectionInformation {
	var _result *IRfcommOutboundConnectionInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutboundConnection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRfcommConnectionTrigger) Get_AllowMultipleConnections() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllowMultipleConnections, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRfcommConnectionTrigger) Put_AllowMultipleConnections(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AllowMultipleConnections, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IRfcommConnectionTrigger) Get_ProtectionLevel() SocketProtectionLevel {
	var _result SocketProtectionLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRfcommConnectionTrigger) Put_ProtectionLevel(value SocketProtectionLevel) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IRfcommConnectionTrigger) Get_RemoteHostName() *IHostName {
	var _result *IHostName
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RemoteHostName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRfcommConnectionTrigger) Put_RemoteHostName(value *IHostName) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RemoteHostName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// F237F327-5181-4F24-96A7-700A4E5FAC62
var IID_ISecondaryAuthenticationFactorAuthenticationTrigger = syscall.GUID{0xF237F327, 0x5181, 0x4F24,
	[8]byte{0x96, 0xA7, 0x70, 0x0A, 0x4E, 0x5F, 0xAC, 0x62}}

type ISecondaryAuthenticationFactorAuthenticationTriggerInterface interface {
	win32.IInspectableInterface
}

type ISecondaryAuthenticationFactorAuthenticationTriggerVtbl struct {
	win32.IInspectableVtbl
}

type ISecondaryAuthenticationFactorAuthenticationTrigger struct {
	win32.IInspectable
}

func (this *ISecondaryAuthenticationFactorAuthenticationTrigger) Vtbl() *ISecondaryAuthenticationFactorAuthenticationTriggerVtbl {
	return (*ISecondaryAuthenticationFactorAuthenticationTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 5BC0F372-D48B-4B7F-ABEC-15F9BACC12E2
var IID_ISensorDataThresholdTrigger = syscall.GUID{0x5BC0F372, 0xD48B, 0x4B7F,
	[8]byte{0xAB, 0xEC, 0x15, 0xF9, 0xBA, 0xCC, 0x12, 0xE2}}

type ISensorDataThresholdTriggerInterface interface {
	win32.IInspectableInterface
}

type ISensorDataThresholdTriggerVtbl struct {
	win32.IInspectableVtbl
}

type ISensorDataThresholdTrigger struct {
	win32.IInspectable
}

func (this *ISensorDataThresholdTrigger) Vtbl() *ISensorDataThresholdTriggerVtbl {
	return (*ISensorDataThresholdTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 921FE675-7DF0-4DA3-97B3-E544EE857FE6
var IID_ISensorDataThresholdTriggerFactory = syscall.GUID{0x921FE675, 0x7DF0, 0x4DA3,
	[8]byte{0x97, 0xB3, 0xE5, 0x44, 0xEE, 0x85, 0x7F, 0xE6}}

type ISensorDataThresholdTriggerFactoryInterface interface {
	win32.IInspectableInterface
	Create(threshold *ISensorDataThreshold) *ISensorDataThresholdTrigger
}

type ISensorDataThresholdTriggerFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ISensorDataThresholdTriggerFactory struct {
	win32.IInspectable
}

func (this *ISensorDataThresholdTriggerFactory) Vtbl() *ISensorDataThresholdTriggerFactoryVtbl {
	return (*ISensorDataThresholdTriggerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISensorDataThresholdTriggerFactory) Create(threshold *ISensorDataThreshold) *ISensorDataThresholdTrigger {
	var _result *ISensorDataThresholdTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(threshold)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F53BC5AC-84CA-4972-8CE9-E58F97B37A50
var IID_ISmartCardTrigger = syscall.GUID{0xF53BC5AC, 0x84CA, 0x4972,
	[8]byte{0x8C, 0xE9, 0xE5, 0x8F, 0x97, 0xB3, 0x7A, 0x50}}

type ISmartCardTriggerInterface interface {
	win32.IInspectableInterface
	Get_TriggerType() unsafe.Pointer
}

type ISmartCardTriggerVtbl struct {
	win32.IInspectableVtbl
	Get_TriggerType uintptr
}

type ISmartCardTrigger struct {
	win32.IInspectable
}

func (this *ISmartCardTrigger) Vtbl() *ISmartCardTriggerVtbl {
	return (*ISmartCardTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmartCardTrigger) Get_TriggerType() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TriggerType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 63BF54C3-89C1-4E00-A9D3-97C629269DAD
var IID_ISmartCardTriggerFactory = syscall.GUID{0x63BF54C3, 0x89C1, 0x4E00,
	[8]byte{0xA9, 0xD3, 0x97, 0xC6, 0x29, 0x26, 0x9D, 0xAD}}

type ISmartCardTriggerFactoryInterface interface {
	win32.IInspectableInterface
	Create(triggerType unsafe.Pointer) *ISmartCardTrigger
}

type ISmartCardTriggerFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ISmartCardTriggerFactory struct {
	win32.IInspectable
}

func (this *ISmartCardTriggerFactory) Vtbl() *ISmartCardTriggerFactoryVtbl {
	return (*ISmartCardTriggerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmartCardTriggerFactory) Create(triggerType unsafe.Pointer) *ISmartCardTrigger {
	var _result *ISmartCardTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(triggerType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EA3AD8C8-6BA4-4AB2-8D21-BC6B09C77564
var IID_ISmsMessageReceivedTriggerFactory = syscall.GUID{0xEA3AD8C8, 0x6BA4, 0x4AB2,
	[8]byte{0x8D, 0x21, 0xBC, 0x6B, 0x09, 0xC7, 0x75, 0x64}}

type ISmsMessageReceivedTriggerFactoryInterface interface {
	win32.IInspectableInterface
	Create(filterRules *ISmsFilterRules) *IBackgroundTrigger
}

type ISmsMessageReceivedTriggerFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ISmsMessageReceivedTriggerFactory struct {
	win32.IInspectable
}

func (this *ISmsMessageReceivedTriggerFactory) Vtbl() *ISmsMessageReceivedTriggerFactoryVtbl {
	return (*ISmsMessageReceivedTriggerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISmsMessageReceivedTriggerFactory) Create(filterRules *ISmsFilterRules) *IBackgroundTrigger {
	var _result *IBackgroundTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(filterRules)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A9BBF810-9DDE-4F8A-83E3-B0E0E7A50D70
var IID_ISocketActivityTrigger = syscall.GUID{0xA9BBF810, 0x9DDE, 0x4F8A,
	[8]byte{0x83, 0xE3, 0xB0, 0xE0, 0xE7, 0xA5, 0x0D, 0x70}}

type ISocketActivityTriggerInterface interface {
	win32.IInspectableInterface
	Get_IsWakeFromLowPowerSupported() bool
}

type ISocketActivityTriggerVtbl struct {
	win32.IInspectableVtbl
	Get_IsWakeFromLowPowerSupported uintptr
}

type ISocketActivityTrigger struct {
	win32.IInspectable
}

func (this *ISocketActivityTrigger) Vtbl() *ISocketActivityTriggerVtbl {
	return (*ISocketActivityTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISocketActivityTrigger) Get_IsWakeFromLowPowerSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsWakeFromLowPowerSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 1EB0FFD0-5A85-499E-A888-824607124F50
var IID_IStorageLibraryChangeTrackerTriggerFactory = syscall.GUID{0x1EB0FFD0, 0x5A85, 0x499E,
	[8]byte{0xA8, 0x88, 0x82, 0x46, 0x07, 0x12, 0x4F, 0x50}}

type IStorageLibraryChangeTrackerTriggerFactoryInterface interface {
	win32.IInspectableInterface
	Create(tracker *IStorageLibraryChangeTracker) *IBackgroundTrigger
}

type IStorageLibraryChangeTrackerTriggerFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IStorageLibraryChangeTrackerTriggerFactory struct {
	win32.IInspectable
}

func (this *IStorageLibraryChangeTrackerTriggerFactory) Vtbl() *IStorageLibraryChangeTrackerTriggerFactoryVtbl {
	return (*IStorageLibraryChangeTrackerTriggerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageLibraryChangeTrackerTriggerFactory) Create(tracker *IStorageLibraryChangeTracker) *IBackgroundTrigger {
	var _result *IBackgroundTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(tracker)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1637E0A7-829C-45BC-929B-A1E7EA78D89B
var IID_IStorageLibraryContentChangedTrigger = syscall.GUID{0x1637E0A7, 0x829C, 0x45BC,
	[8]byte{0x92, 0x9B, 0xA1, 0xE7, 0xEA, 0x78, 0xD8, 0x9B}}

type IStorageLibraryContentChangedTriggerInterface interface {
	win32.IInspectableInterface
}

type IStorageLibraryContentChangedTriggerVtbl struct {
	win32.IInspectableVtbl
}

type IStorageLibraryContentChangedTrigger struct {
	win32.IInspectable
}

func (this *IStorageLibraryContentChangedTrigger) Vtbl() *IStorageLibraryContentChangedTriggerVtbl {
	return (*IStorageLibraryContentChangedTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 7F9F1B39-5F90-4E12-914E-A7D8E0BBFB18
var IID_IStorageLibraryContentChangedTriggerStatics = syscall.GUID{0x7F9F1B39, 0x5F90, 0x4E12,
	[8]byte{0x91, 0x4E, 0xA7, 0xD8, 0xE0, 0xBB, 0xFB, 0x18}}

type IStorageLibraryContentChangedTriggerStaticsInterface interface {
	win32.IInspectableInterface
	Create(storageLibrary *IStorageLibrary) *IStorageLibraryContentChangedTrigger
	CreateFromLibraries(storageLibraries *IIterable[*IStorageLibrary]) *IStorageLibraryContentChangedTrigger
}

type IStorageLibraryContentChangedTriggerStaticsVtbl struct {
	win32.IInspectableVtbl
	Create              uintptr
	CreateFromLibraries uintptr
}

type IStorageLibraryContentChangedTriggerStatics struct {
	win32.IInspectable
}

func (this *IStorageLibraryContentChangedTriggerStatics) Vtbl() *IStorageLibraryContentChangedTriggerStaticsVtbl {
	return (*IStorageLibraryContentChangedTriggerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageLibraryContentChangedTriggerStatics) Create(storageLibrary *IStorageLibrary) *IStorageLibraryContentChangedTrigger {
	var _result *IStorageLibraryContentChangedTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(storageLibrary)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageLibraryContentChangedTriggerStatics) CreateFromLibraries(storageLibraries *IIterable[*IStorageLibrary]) *IStorageLibraryContentChangedTrigger {
	var _result *IStorageLibraryContentChangedTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromLibraries, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(storageLibraries)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C15FB476-89C5-420B-ABD3-FB3030472128
var IID_ISystemCondition = syscall.GUID{0xC15FB476, 0x89C5, 0x420B,
	[8]byte{0xAB, 0xD3, 0xFB, 0x30, 0x30, 0x47, 0x21, 0x28}}

type ISystemConditionInterface interface {
	win32.IInspectableInterface
	Get_ConditionType() SystemConditionType
}

type ISystemConditionVtbl struct {
	win32.IInspectableVtbl
	Get_ConditionType uintptr
}

type ISystemCondition struct {
	win32.IInspectable
}

func (this *ISystemCondition) Vtbl() *ISystemConditionVtbl {
	return (*ISystemConditionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemCondition) Get_ConditionType() SystemConditionType {
	var _result SystemConditionType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConditionType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D269D1F1-05A7-49AE-87D7-16B2B8B9A553
var IID_ISystemConditionFactory = syscall.GUID{0xD269D1F1, 0x05A7, 0x49AE,
	[8]byte{0x87, 0xD7, 0x16, 0xB2, 0xB8, 0xB9, 0xA5, 0x53}}

type ISystemConditionFactoryInterface interface {
	win32.IInspectableInterface
	Create(conditionType SystemConditionType) *ISystemCondition
}

type ISystemConditionFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ISystemConditionFactory struct {
	win32.IInspectable
}

func (this *ISystemConditionFactory) Vtbl() *ISystemConditionFactoryVtbl {
	return (*ISystemConditionFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemConditionFactory) Create(conditionType SystemConditionType) *ISystemCondition {
	var _result *ISystemCondition
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(conditionType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1D80C776-3748-4463-8D7E-276DC139AC1C
var IID_ISystemTrigger = syscall.GUID{0x1D80C776, 0x3748, 0x4463,
	[8]byte{0x8D, 0x7E, 0x27, 0x6D, 0xC1, 0x39, 0xAC, 0x1C}}

type ISystemTriggerInterface interface {
	win32.IInspectableInterface
	Get_OneShot() bool
	Get_TriggerType() SystemTriggerType
}

type ISystemTriggerVtbl struct {
	win32.IInspectableVtbl
	Get_OneShot     uintptr
	Get_TriggerType uintptr
}

type ISystemTrigger struct {
	win32.IInspectable
}

func (this *ISystemTrigger) Vtbl() *ISystemTriggerVtbl {
	return (*ISystemTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemTrigger) Get_OneShot() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OneShot, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemTrigger) Get_TriggerType() SystemTriggerType {
	var _result SystemTriggerType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TriggerType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// E80423D4-8791-4579-8126-87EC8AAA407A
var IID_ISystemTriggerFactory = syscall.GUID{0xE80423D4, 0x8791, 0x4579,
	[8]byte{0x81, 0x26, 0x87, 0xEC, 0x8A, 0xAA, 0x40, 0x7A}}

type ISystemTriggerFactoryInterface interface {
	win32.IInspectableInterface
	Create(triggerType SystemTriggerType, oneShot bool) *ISystemTrigger
}

type ISystemTriggerFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ISystemTriggerFactory struct {
	win32.IInspectable
}

func (this *ISystemTriggerFactory) Vtbl() *ISystemTriggerFactoryVtbl {
	return (*ISystemTriggerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemTriggerFactory) Create(triggerType SystemTriggerType, oneShot bool) *ISystemTrigger {
	var _result *ISystemTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(triggerType), uintptr(*(*byte)(unsafe.Pointer(&oneShot))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 656E5556-0B2A-4377-BA70-3B45A935547F
var IID_ITimeTrigger = syscall.GUID{0x656E5556, 0x0B2A, 0x4377,
	[8]byte{0xBA, 0x70, 0x3B, 0x45, 0xA9, 0x35, 0x54, 0x7F}}

type ITimeTriggerInterface interface {
	win32.IInspectableInterface
	Get_FreshnessTime() uint32
	Get_OneShot() bool
}

type ITimeTriggerVtbl struct {
	win32.IInspectableVtbl
	Get_FreshnessTime uintptr
	Get_OneShot       uintptr
}

type ITimeTrigger struct {
	win32.IInspectable
}

func (this *ITimeTrigger) Vtbl() *ITimeTriggerVtbl {
	return (*ITimeTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimeTrigger) Get_FreshnessTime() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FreshnessTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimeTrigger) Get_OneShot() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OneShot, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 38C682FE-9B54-45E6-B2F3-269B87A6F734
var IID_ITimeTriggerFactory = syscall.GUID{0x38C682FE, 0x9B54, 0x45E6,
	[8]byte{0xB2, 0xF3, 0x26, 0x9B, 0x87, 0xA6, 0xF7, 0x34}}

type ITimeTriggerFactoryInterface interface {
	win32.IInspectableInterface
	Create(freshnessTime uint32, oneShot bool) *ITimeTrigger
}

type ITimeTriggerFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ITimeTriggerFactory struct {
	win32.IInspectable
}

func (this *ITimeTriggerFactory) Vtbl() *ITimeTriggerFactoryVtbl {
	return (*ITimeTriggerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimeTriggerFactory) Create(freshnessTime uint32, oneShot bool) *ITimeTrigger {
	var _result *ITimeTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(freshnessTime), uintptr(*(*byte)(unsafe.Pointer(&oneShot))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B09DFC27-6480-4349-8125-97B3EFAA0A3A
var IID_IToastNotificationActionTriggerFactory = syscall.GUID{0xB09DFC27, 0x6480, 0x4349,
	[8]byte{0x81, 0x25, 0x97, 0xB3, 0xEF, 0xAA, 0x0A, 0x3A}}

type IToastNotificationActionTriggerFactoryInterface interface {
	win32.IInspectableInterface
	Create(applicationId string) *IBackgroundTrigger
}

type IToastNotificationActionTriggerFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IToastNotificationActionTriggerFactory struct {
	win32.IInspectable
}

func (this *IToastNotificationActionTriggerFactory) Vtbl() *IToastNotificationActionTriggerFactoryVtbl {
	return (*IToastNotificationActionTriggerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastNotificationActionTriggerFactory) Create(applicationId string) *IBackgroundTrigger {
	var _result *IBackgroundTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(applicationId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 81C6FAAD-8797-4785-81B4-B0CCCB73D1D9
var IID_IToastNotificationHistoryChangedTriggerFactory = syscall.GUID{0x81C6FAAD, 0x8797, 0x4785,
	[8]byte{0x81, 0xB4, 0xB0, 0xCC, 0xCB, 0x73, 0xD1, 0xD9}}

type IToastNotificationHistoryChangedTriggerFactoryInterface interface {
	win32.IInspectableInterface
	Create(applicationId string) *IBackgroundTrigger
}

type IToastNotificationHistoryChangedTriggerFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IToastNotificationHistoryChangedTriggerFactory struct {
	win32.IInspectable
}

func (this *IToastNotificationHistoryChangedTriggerFactory) Vtbl() *IToastNotificationHistoryChangedTriggerFactoryVtbl {
	return (*IToastNotificationHistoryChangedTriggerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToastNotificationHistoryChangedTriggerFactory) Create(applicationId string) *IBackgroundTrigger {
	var _result *IBackgroundTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(applicationId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CAD4436C-69AB-4E18-A48A-5ED2AC435957
var IID_IUserNotificationChangedTriggerFactory = syscall.GUID{0xCAD4436C, 0x69AB, 0x4E18,
	[8]byte{0xA4, 0x8A, 0x5E, 0xD2, 0xAC, 0x43, 0x59, 0x57}}

type IUserNotificationChangedTriggerFactoryInterface interface {
	win32.IInspectableInterface
	Create(notificationKinds NotificationKinds) *IBackgroundTrigger
}

type IUserNotificationChangedTriggerFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IUserNotificationChangedTriggerFactory struct {
	win32.IInspectable
}

func (this *IUserNotificationChangedTriggerFactory) Vtbl() *IUserNotificationChangedTriggerFactoryVtbl {
	return (*IUserNotificationChangedTriggerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUserNotificationChangedTriggerFactory) Create(notificationKinds NotificationKinds) *IBackgroundTrigger {
	var _result *IBackgroundTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(notificationKinds), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type BackgroundTaskBuilder struct {
	RtClass
	*IBackgroundTaskBuilder
}

func NewBackgroundTaskBuilder() *BackgroundTaskBuilder {
	hs := NewHStr("Windows.ApplicationModel.Background.BackgroundTaskBuilder")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &BackgroundTaskBuilder{
		RtClass:                RtClass{PInspect: p},
		IBackgroundTaskBuilder: (*IBackgroundTaskBuilder)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type BackgroundWorkCost struct {
	RtClass
}

func NewIBackgroundWorkCostStatics() *IBackgroundWorkCostStatics {
	var p *IBackgroundWorkCostStatics
	hs := NewHStr("Windows.ApplicationModel.Background.BackgroundWorkCost")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IBackgroundWorkCostStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type ConversationalAgentTrigger struct {
	RtClass
	*IBackgroundTrigger
}

func NewConversationalAgentTrigger() *ConversationalAgentTrigger {
	hs := NewHStr("Windows.ApplicationModel.Background.ConversationalAgentTrigger")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &ConversationalAgentTrigger{
		RtClass:            RtClass{PInspect: p},
		IBackgroundTrigger: (*IBackgroundTrigger)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}
