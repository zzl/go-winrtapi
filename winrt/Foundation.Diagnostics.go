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
type CausalityRelation int32

const (
	CausalityRelation_AssignDelegate CausalityRelation = 0
	CausalityRelation_Join           CausalityRelation = 1
	CausalityRelation_Choice         CausalityRelation = 2
	CausalityRelation_Cancel         CausalityRelation = 3
	CausalityRelation_Error          CausalityRelation = 4
)

// enum
type CausalitySource int32

const (
	CausalitySource_Application CausalitySource = 0
	CausalitySource_Library     CausalitySource = 1
	CausalitySource_System      CausalitySource = 2
)

// enum
type CausalitySynchronousWork int32

const (
	CausalitySynchronousWork_CompletionNotification CausalitySynchronousWork = 0
	CausalitySynchronousWork_ProgressNotification   CausalitySynchronousWork = 1
	CausalitySynchronousWork_Execution              CausalitySynchronousWork = 2
)

// enum
type CausalityTraceLevel int32

const (
	CausalityTraceLevel_Required  CausalityTraceLevel = 0
	CausalityTraceLevel_Important CausalityTraceLevel = 1
	CausalityTraceLevel_Verbose   CausalityTraceLevel = 2
)

// enum
// flags
type ErrorOptions uint32

const (
	ErrorOptions_None                 ErrorOptions = 0
	ErrorOptions_SuppressExceptions   ErrorOptions = 1
	ErrorOptions_ForceExceptions      ErrorOptions = 2
	ErrorOptions_UseSetErrorInfo      ErrorOptions = 4
	ErrorOptions_SuppressSetErrorInfo ErrorOptions = 8
)

// enum
type LoggingFieldFormat int32

const (
	LoggingFieldFormat_Default       LoggingFieldFormat = 0
	LoggingFieldFormat_Hidden        LoggingFieldFormat = 1
	LoggingFieldFormat_String        LoggingFieldFormat = 2
	LoggingFieldFormat_Boolean       LoggingFieldFormat = 3
	LoggingFieldFormat_Hexadecimal   LoggingFieldFormat = 4
	LoggingFieldFormat_ProcessId     LoggingFieldFormat = 5
	LoggingFieldFormat_ThreadId      LoggingFieldFormat = 6
	LoggingFieldFormat_Port          LoggingFieldFormat = 7
	LoggingFieldFormat_Ipv4Address   LoggingFieldFormat = 8
	LoggingFieldFormat_Ipv6Address   LoggingFieldFormat = 9
	LoggingFieldFormat_SocketAddress LoggingFieldFormat = 10
	LoggingFieldFormat_Xml           LoggingFieldFormat = 11
	LoggingFieldFormat_Json          LoggingFieldFormat = 12
	LoggingFieldFormat_Win32Error    LoggingFieldFormat = 13
	LoggingFieldFormat_NTStatus      LoggingFieldFormat = 14
	LoggingFieldFormat_HResult       LoggingFieldFormat = 15
	LoggingFieldFormat_FileTime      LoggingFieldFormat = 16
	LoggingFieldFormat_Signed        LoggingFieldFormat = 17
	LoggingFieldFormat_Unsigned      LoggingFieldFormat = 18
)

// enum
type LoggingLevel int32

const (
	LoggingLevel_Verbose     LoggingLevel = 0
	LoggingLevel_Information LoggingLevel = 1
	LoggingLevel_Warning     LoggingLevel = 2
	LoggingLevel_Error       LoggingLevel = 3
	LoggingLevel_Critical    LoggingLevel = 4
)

// enum
type LoggingOpcode int32

const (
	LoggingOpcode_Info    LoggingOpcode = 0
	LoggingOpcode_Start   LoggingOpcode = 1
	LoggingOpcode_Stop    LoggingOpcode = 2
	LoggingOpcode_Reply   LoggingOpcode = 6
	LoggingOpcode_Resume  LoggingOpcode = 7
	LoggingOpcode_Suspend LoggingOpcode = 8
	LoggingOpcode_Send    LoggingOpcode = 9
)

// interfaces

// 50850B26-267E-451B-A890-AB6A370245EE
var IID_IAsyncCausalityTracerStatics = syscall.GUID{0x50850B26, 0x267E, 0x451B,
	[8]byte{0xA8, 0x90, 0xAB, 0x6A, 0x37, 0x02, 0x45, 0xEE}}

type IAsyncCausalityTracerStaticsInterface interface {
	win32.IInspectableInterface
	TraceOperationCreation(traceLevel CausalityTraceLevel, source CausalitySource, platformId syscall.GUID, operationId uint64, operationName string, relatedContext uint64)
	TraceOperationCompletion(traceLevel CausalityTraceLevel, source CausalitySource, platformId syscall.GUID, operationId uint64, status AsyncStatus)
	TraceOperationRelation(traceLevel CausalityTraceLevel, source CausalitySource, platformId syscall.GUID, operationId uint64, relation CausalityRelation)
	TraceSynchronousWorkStart(traceLevel CausalityTraceLevel, source CausalitySource, platformId syscall.GUID, operationId uint64, work CausalitySynchronousWork)
	TraceSynchronousWorkCompletion(traceLevel CausalityTraceLevel, source CausalitySource, work CausalitySynchronousWork)
	Add_TracingStatusChanged(handler EventHandler[*ITracingStatusChangedEventArgs]) EventRegistrationToken
	Remove_TracingStatusChanged(cookie EventRegistrationToken)
}

type IAsyncCausalityTracerStaticsVtbl struct {
	win32.IInspectableVtbl
	TraceOperationCreation         uintptr
	TraceOperationCompletion       uintptr
	TraceOperationRelation         uintptr
	TraceSynchronousWorkStart      uintptr
	TraceSynchronousWorkCompletion uintptr
	Add_TracingStatusChanged       uintptr
	Remove_TracingStatusChanged    uintptr
}

type IAsyncCausalityTracerStatics struct {
	win32.IInspectable
}

func (this *IAsyncCausalityTracerStatics) Vtbl() *IAsyncCausalityTracerStaticsVtbl {
	return (*IAsyncCausalityTracerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAsyncCausalityTracerStatics) TraceOperationCreation(traceLevel CausalityTraceLevel, source CausalitySource, platformId syscall.GUID, operationId uint64, operationName string, relatedContext uint64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TraceOperationCreation, uintptr(unsafe.Pointer(this)), uintptr(traceLevel), uintptr(source), uintptr(unsafe.Pointer(&platformId)), uintptr(operationId), NewHStr(operationName).Ptr, uintptr(relatedContext))
	_ = _hr
}

func (this *IAsyncCausalityTracerStatics) TraceOperationCompletion(traceLevel CausalityTraceLevel, source CausalitySource, platformId syscall.GUID, operationId uint64, status AsyncStatus) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TraceOperationCompletion, uintptr(unsafe.Pointer(this)), uintptr(traceLevel), uintptr(source), uintptr(unsafe.Pointer(&platformId)), uintptr(operationId), uintptr(status))
	_ = _hr
}

func (this *IAsyncCausalityTracerStatics) TraceOperationRelation(traceLevel CausalityTraceLevel, source CausalitySource, platformId syscall.GUID, operationId uint64, relation CausalityRelation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TraceOperationRelation, uintptr(unsafe.Pointer(this)), uintptr(traceLevel), uintptr(source), uintptr(unsafe.Pointer(&platformId)), uintptr(operationId), uintptr(relation))
	_ = _hr
}

func (this *IAsyncCausalityTracerStatics) TraceSynchronousWorkStart(traceLevel CausalityTraceLevel, source CausalitySource, platformId syscall.GUID, operationId uint64, work CausalitySynchronousWork) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TraceSynchronousWorkStart, uintptr(unsafe.Pointer(this)), uintptr(traceLevel), uintptr(source), uintptr(unsafe.Pointer(&platformId)), uintptr(operationId), uintptr(work))
	_ = _hr
}

func (this *IAsyncCausalityTracerStatics) TraceSynchronousWorkCompletion(traceLevel CausalityTraceLevel, source CausalitySource, work CausalitySynchronousWork) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TraceSynchronousWorkCompletion, uintptr(unsafe.Pointer(this)), uintptr(traceLevel), uintptr(source), uintptr(work))
	_ = _hr
}

func (this *IAsyncCausalityTracerStatics) Add_TracingStatusChanged(handler EventHandler[*ITracingStatusChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_TracingStatusChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAsyncCausalityTracerStatics) Remove_TracingStatusChanged(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_TracingStatusChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

// 378CBB01-2CC9-428F-8C55-2C990D463E8F
var IID_IErrorDetails = syscall.GUID{0x378CBB01, 0x2CC9, 0x428F,
	[8]byte{0x8C, 0x55, 0x2C, 0x99, 0x0D, 0x46, 0x3E, 0x8F}}

type IErrorDetailsInterface interface {
	win32.IInspectableInterface
	Get_Description() string
	Get_LongDescription() string
	Get_HelpUri() *IUriRuntimeClass
}

type IErrorDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_Description     uintptr
	Get_LongDescription uintptr
	Get_HelpUri         uintptr
}

type IErrorDetails struct {
	win32.IInspectable
}

func (this *IErrorDetails) Vtbl() *IErrorDetailsVtbl {
	return (*IErrorDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IErrorDetails) Get_Description() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Description, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IErrorDetails) Get_LongDescription() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LongDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IErrorDetails) Get_HelpUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HelpUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B7703750-0B1D-46C8-AA0E-4B8178E4FCE9
var IID_IErrorDetailsStatics = syscall.GUID{0xB7703750, 0x0B1D, 0x46C8,
	[8]byte{0xAA, 0x0E, 0x4B, 0x81, 0x78, 0xE4, 0xFC, 0xE9}}

type IErrorDetailsStaticsInterface interface {
	win32.IInspectableInterface
	CreateFromHResultAsync(errorCode int32) *IAsyncOperation[*IErrorDetails]
}

type IErrorDetailsStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateFromHResultAsync uintptr
}

type IErrorDetailsStatics struct {
	win32.IInspectable
}

func (this *IErrorDetailsStatics) Vtbl() *IErrorDetailsStaticsVtbl {
	return (*IErrorDetailsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IErrorDetailsStatics) CreateFromHResultAsync(errorCode int32) *IAsyncOperation[*IErrorDetails] {
	var _result *IAsyncOperation[*IErrorDetails]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromHResultAsync, uintptr(unsafe.Pointer(this)), uintptr(errorCode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 16369792-B03E-4BA1-8BB8-D28F4AB4D2C0
var IID_IErrorReportingSettings = syscall.GUID{0x16369792, 0xB03E, 0x4BA1,
	[8]byte{0x8B, 0xB8, 0xD2, 0x8F, 0x4A, 0xB4, 0xD2, 0xC0}}

type IErrorReportingSettingsInterface interface {
	win32.IInspectableInterface
	SetErrorOptions(value ErrorOptions)
	GetErrorOptions() ErrorOptions
}

type IErrorReportingSettingsVtbl struct {
	win32.IInspectableVtbl
	SetErrorOptions uintptr
	GetErrorOptions uintptr
}

type IErrorReportingSettings struct {
	win32.IInspectable
}

func (this *IErrorReportingSettings) Vtbl() *IErrorReportingSettingsVtbl {
	return (*IErrorReportingSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IErrorReportingSettings) SetErrorOptions(value ErrorOptions) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetErrorOptions, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IErrorReportingSettings) GetErrorOptions() ErrorOptions {
	var _result ErrorOptions
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetErrorOptions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 24C74216-FED2-404C-895F-1F9699CB02F7
var IID_IFileLoggingSession = syscall.GUID{0x24C74216, 0xFED2, 0x404C,
	[8]byte{0x89, 0x5F, 0x1F, 0x96, 0x99, 0xCB, 0x02, 0xF7}}

type IFileLoggingSessionInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	AddLoggingChannel(loggingChannel *ILoggingChannel)
	AddLoggingChannelWithLevel(loggingChannel *ILoggingChannel, maxLevel LoggingLevel)
	RemoveLoggingChannel(loggingChannel *ILoggingChannel)
	CloseAndSaveToFileAsync() *IAsyncOperation[*IStorageFile]
	Add_LogFileGenerated(handler TypedEventHandler[*IFileLoggingSession, *ILogFileGeneratedEventArgs]) EventRegistrationToken
	Remove_LogFileGenerated(token EventRegistrationToken)
}

type IFileLoggingSessionVtbl struct {
	win32.IInspectableVtbl
	Get_Name                   uintptr
	AddLoggingChannel          uintptr
	AddLoggingChannelWithLevel uintptr
	RemoveLoggingChannel       uintptr
	CloseAndSaveToFileAsync    uintptr
	Add_LogFileGenerated       uintptr
	Remove_LogFileGenerated    uintptr
}

type IFileLoggingSession struct {
	win32.IInspectable
}

func (this *IFileLoggingSession) Vtbl() *IFileLoggingSessionVtbl {
	return (*IFileLoggingSessionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFileLoggingSession) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IFileLoggingSession) AddLoggingChannel(loggingChannel *ILoggingChannel) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddLoggingChannel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(loggingChannel)))
	_ = _hr
}

func (this *IFileLoggingSession) AddLoggingChannelWithLevel(loggingChannel *ILoggingChannel, maxLevel LoggingLevel) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddLoggingChannelWithLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(loggingChannel)), uintptr(maxLevel))
	_ = _hr
}

func (this *IFileLoggingSession) RemoveLoggingChannel(loggingChannel *ILoggingChannel) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveLoggingChannel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(loggingChannel)))
	_ = _hr
}

func (this *IFileLoggingSession) CloseAndSaveToFileAsync() *IAsyncOperation[*IStorageFile] {
	var _result *IAsyncOperation[*IStorageFile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CloseAndSaveToFileAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileLoggingSession) Add_LogFileGenerated(handler TypedEventHandler[*IFileLoggingSession, *ILogFileGeneratedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_LogFileGenerated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFileLoggingSession) Remove_LogFileGenerated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_LogFileGenerated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// EEA08DCE-8447-4DAA-9133-12EB46F697D4
var IID_IFileLoggingSessionFactory = syscall.GUID{0xEEA08DCE, 0x8447, 0x4DAA,
	[8]byte{0x91, 0x33, 0x12, 0xEB, 0x46, 0xF6, 0x97, 0xD4}}

type IFileLoggingSessionFactoryInterface interface {
	win32.IInspectableInterface
	Create(name string) *IFileLoggingSession
}

type IFileLoggingSessionFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IFileLoggingSessionFactory struct {
	win32.IInspectable
}

func (this *IFileLoggingSessionFactory) Vtbl() *IFileLoggingSessionFactoryVtbl {
	return (*IFileLoggingSessionFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFileLoggingSessionFactory) Create(name string) *IFileLoggingSession {
	var _result *IFileLoggingSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 269E976F-0D38-4C1A-B53F-B395D881DF84
var IID_ILogFileGeneratedEventArgs = syscall.GUID{0x269E976F, 0x0D38, 0x4C1A,
	[8]byte{0xB5, 0x3F, 0xB3, 0x95, 0xD8, 0x81, 0xDF, 0x84}}

type ILogFileGeneratedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_File() *IStorageFile
}

type ILogFileGeneratedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_File uintptr
}

type ILogFileGeneratedEventArgs struct {
	win32.IInspectable
}

func (this *ILogFileGeneratedEventArgs) Vtbl() *ILogFileGeneratedEventArgsVtbl {
	return (*ILogFileGeneratedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILogFileGeneratedEventArgs) Get_File() *IStorageFile {
	var _result *IStorageFile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_File, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BC032941-B766-4CB5-9848-97AC6BA6D60C
var IID_ILoggingActivity = syscall.GUID{0xBC032941, 0xB766, 0x4CB5,
	[8]byte{0x98, 0x48, 0x97, 0xAC, 0x6B, 0xA6, 0xD6, 0x0C}}

type ILoggingActivityInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Get_Id() syscall.GUID
}

type ILoggingActivityVtbl struct {
	win32.IInspectableVtbl
	Get_Name uintptr
	Get_Id   uintptr
}

type ILoggingActivity struct {
	win32.IInspectable
}

func (this *ILoggingActivity) Vtbl() *ILoggingActivityVtbl {
	return (*ILoggingActivityVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILoggingActivity) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILoggingActivity) Get_Id() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 26C29808-6322-456A-AF82-80C8642F178B
var IID_ILoggingActivity2 = syscall.GUID{0x26C29808, 0x6322, 0x456A,
	[8]byte{0xAF, 0x82, 0x80, 0xC8, 0x64, 0x2F, 0x17, 0x8B}}

type ILoggingActivity2Interface interface {
	win32.IInspectableInterface
	Get_Channel() *ILoggingChannel
	StopActivity(stopEventName string)
	StopActivityWithFields(stopEventName string, fields *ILoggingFields)
	StopActivityWithFieldsAndOptions(stopEventName string, fields *ILoggingFields, options *ILoggingOptions)
}

type ILoggingActivity2Vtbl struct {
	win32.IInspectableVtbl
	Get_Channel                      uintptr
	StopActivity                     uintptr
	StopActivityWithFields           uintptr
	StopActivityWithFieldsAndOptions uintptr
}

type ILoggingActivity2 struct {
	win32.IInspectable
}

func (this *ILoggingActivity2) Vtbl() *ILoggingActivity2Vtbl {
	return (*ILoggingActivity2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILoggingActivity2) Get_Channel() *ILoggingChannel {
	var _result *ILoggingChannel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Channel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILoggingActivity2) StopActivity(stopEventName string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopActivity, uintptr(unsafe.Pointer(this)), NewHStr(stopEventName).Ptr)
	_ = _hr
}

func (this *ILoggingActivity2) StopActivityWithFields(stopEventName string, fields *ILoggingFields) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopActivityWithFields, uintptr(unsafe.Pointer(this)), NewHStr(stopEventName).Ptr, uintptr(unsafe.Pointer(fields)))
	_ = _hr
}

func (this *ILoggingActivity2) StopActivityWithFieldsAndOptions(stopEventName string, fields *ILoggingFields, options *ILoggingOptions) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopActivityWithFieldsAndOptions, uintptr(unsafe.Pointer(this)), NewHStr(stopEventName).Ptr, uintptr(unsafe.Pointer(fields)), uintptr(unsafe.Pointer(options)))
	_ = _hr
}

// 6B33B483-E10A-4C58-97D5-10FB451074FB
var IID_ILoggingActivityFactory = syscall.GUID{0x6B33B483, 0xE10A, 0x4C58,
	[8]byte{0x97, 0xD5, 0x10, 0xFB, 0x45, 0x10, 0x74, 0xFB}}

type ILoggingActivityFactoryInterface interface {
	win32.IInspectableInterface
	CreateLoggingActivity(activityName string, loggingChannel *ILoggingChannel) *ILoggingActivity
	CreateLoggingActivityWithLevel(activityName string, loggingChannel *ILoggingChannel, level LoggingLevel) *ILoggingActivity
}

type ILoggingActivityFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateLoggingActivity          uintptr
	CreateLoggingActivityWithLevel uintptr
}

type ILoggingActivityFactory struct {
	win32.IInspectable
}

func (this *ILoggingActivityFactory) Vtbl() *ILoggingActivityFactoryVtbl {
	return (*ILoggingActivityFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILoggingActivityFactory) CreateLoggingActivity(activityName string, loggingChannel *ILoggingChannel) *ILoggingActivity {
	var _result *ILoggingActivity
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateLoggingActivity, uintptr(unsafe.Pointer(this)), NewHStr(activityName).Ptr, uintptr(unsafe.Pointer(loggingChannel)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILoggingActivityFactory) CreateLoggingActivityWithLevel(activityName string, loggingChannel *ILoggingChannel, level LoggingLevel) *ILoggingActivity {
	var _result *ILoggingActivity
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateLoggingActivityWithLevel, uintptr(unsafe.Pointer(this)), NewHStr(activityName).Ptr, uintptr(unsafe.Pointer(loggingChannel)), uintptr(level), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E9A50343-11D7-4F01-B5CA-CF495278C0A8
var IID_ILoggingChannel = syscall.GUID{0xE9A50343, 0x11D7, 0x4F01,
	[8]byte{0xB5, 0xCA, 0xCF, 0x49, 0x52, 0x78, 0xC0, 0xA8}}

type ILoggingChannelInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Get_Enabled() bool
	Get_Level() LoggingLevel
	LogMessage(eventString string)
	LogMessageWithLevel(eventString string, level LoggingLevel)
	LogValuePair(value1 string, value2 int32)
	LogValuePairWithLevel(value1 string, value2 int32, level LoggingLevel)
	Add_LoggingEnabled(handler TypedEventHandler[*ILoggingChannel, interface{}]) EventRegistrationToken
	Remove_LoggingEnabled(token EventRegistrationToken)
}

type ILoggingChannelVtbl struct {
	win32.IInspectableVtbl
	Get_Name              uintptr
	Get_Enabled           uintptr
	Get_Level             uintptr
	LogMessage            uintptr
	LogMessageWithLevel   uintptr
	LogValuePair          uintptr
	LogValuePairWithLevel uintptr
	Add_LoggingEnabled    uintptr
	Remove_LoggingEnabled uintptr
}

type ILoggingChannel struct {
	win32.IInspectable
}

func (this *ILoggingChannel) Vtbl() *ILoggingChannelVtbl {
	return (*ILoggingChannelVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILoggingChannel) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILoggingChannel) Get_Enabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Enabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILoggingChannel) Get_Level() LoggingLevel {
	var _result LoggingLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Level, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILoggingChannel) LogMessage(eventString string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LogMessage, uintptr(unsafe.Pointer(this)), NewHStr(eventString).Ptr)
	_ = _hr
}

func (this *ILoggingChannel) LogMessageWithLevel(eventString string, level LoggingLevel) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LogMessageWithLevel, uintptr(unsafe.Pointer(this)), NewHStr(eventString).Ptr, uintptr(level))
	_ = _hr
}

func (this *ILoggingChannel) LogValuePair(value1 string, value2 int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LogValuePair, uintptr(unsafe.Pointer(this)), NewHStr(value1).Ptr, uintptr(value2))
	_ = _hr
}

func (this *ILoggingChannel) LogValuePairWithLevel(value1 string, value2 int32, level LoggingLevel) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LogValuePairWithLevel, uintptr(unsafe.Pointer(this)), NewHStr(value1).Ptr, uintptr(value2), uintptr(level))
	_ = _hr
}

func (this *ILoggingChannel) Add_LoggingEnabled(handler TypedEventHandler[*ILoggingChannel, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_LoggingEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILoggingChannel) Remove_LoggingEnabled(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_LoggingEnabled, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 9F4C3CF3-0BAC-45A5-9E33-BAF3F3A246A5
var IID_ILoggingChannel2 = syscall.GUID{0x9F4C3CF3, 0x0BAC, 0x45A5,
	[8]byte{0x9E, 0x33, 0xBA, 0xF3, 0xF3, 0xA2, 0x46, 0xA5}}

type ILoggingChannel2Interface interface {
	win32.IInspectableInterface
	Get_Id() syscall.GUID
}

type ILoggingChannel2Vtbl struct {
	win32.IInspectableVtbl
	Get_Id uintptr
}

type ILoggingChannel2 struct {
	win32.IInspectable
}

func (this *ILoggingChannel2) Vtbl() *ILoggingChannel2Vtbl {
	return (*ILoggingChannel2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILoggingChannel2) Get_Id() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4EDC5B9C-AF80-4A9B-B0DC-398F9AE5207B
var IID_ILoggingChannelFactory = syscall.GUID{0x4EDC5B9C, 0xAF80, 0x4A9B,
	[8]byte{0xB0, 0xDC, 0x39, 0x8F, 0x9A, 0xE5, 0x20, 0x7B}}

type ILoggingChannelFactoryInterface interface {
	win32.IInspectableInterface
	Create(name string) *ILoggingChannel
}

type ILoggingChannelFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ILoggingChannelFactory struct {
	win32.IInspectable
}

func (this *ILoggingChannelFactory) Vtbl() *ILoggingChannelFactoryVtbl {
	return (*ILoggingChannelFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILoggingChannelFactory) Create(name string) *ILoggingChannel {
	var _result *ILoggingChannel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4C6EF5DD-3B27-4DC9-99F0-299C6E4603A1
var IID_ILoggingChannelFactory2 = syscall.GUID{0x4C6EF5DD, 0x3B27, 0x4DC9,
	[8]byte{0x99, 0xF0, 0x29, 0x9C, 0x6E, 0x46, 0x03, 0xA1}}

type ILoggingChannelFactory2Interface interface {
	win32.IInspectableInterface
	CreateWithOptions(name string, options *ILoggingChannelOptions) *ILoggingChannel
	CreateWithOptionsAndId(name string, options *ILoggingChannelOptions, id syscall.GUID) *ILoggingChannel
}

type ILoggingChannelFactory2Vtbl struct {
	win32.IInspectableVtbl
	CreateWithOptions      uintptr
	CreateWithOptionsAndId uintptr
}

type ILoggingChannelFactory2 struct {
	win32.IInspectable
}

func (this *ILoggingChannelFactory2) Vtbl() *ILoggingChannelFactory2Vtbl {
	return (*ILoggingChannelFactory2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILoggingChannelFactory2) CreateWithOptions(name string, options *ILoggingChannelOptions) *ILoggingChannel {
	var _result *ILoggingChannel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithOptions, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILoggingChannelFactory2) CreateWithOptionsAndId(name string, options *ILoggingChannelOptions, id syscall.GUID) *ILoggingChannel {
	var _result *ILoggingChannel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithOptionsAndId, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&id)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C3E847FF-0EBB-4A53-8C54-DEC24926CB2C
var IID_ILoggingChannelOptions = syscall.GUID{0xC3E847FF, 0x0EBB, 0x4A53,
	[8]byte{0x8C, 0x54, 0xDE, 0xC2, 0x49, 0x26, 0xCB, 0x2C}}

type ILoggingChannelOptionsInterface interface {
	win32.IInspectableInterface
	Get_Group() syscall.GUID
	Put_Group(value syscall.GUID)
}

type ILoggingChannelOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_Group uintptr
	Put_Group uintptr
}

type ILoggingChannelOptions struct {
	win32.IInspectable
}

func (this *ILoggingChannelOptions) Vtbl() *ILoggingChannelOptionsVtbl {
	return (*ILoggingChannelOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILoggingChannelOptions) Get_Group() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Group, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILoggingChannelOptions) Put_Group(value syscall.GUID) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Group, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

// A93151DA-7FAF-4191-8755-5E86DC65D896
var IID_ILoggingChannelOptionsFactory = syscall.GUID{0xA93151DA, 0x7FAF, 0x4191,
	[8]byte{0x87, 0x55, 0x5E, 0x86, 0xDC, 0x65, 0xD8, 0x96}}

type ILoggingChannelOptionsFactoryInterface interface {
	win32.IInspectableInterface
	Create(group syscall.GUID) *ILoggingChannelOptions
}

type ILoggingChannelOptionsFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ILoggingChannelOptionsFactory struct {
	win32.IInspectable
}

func (this *ILoggingChannelOptionsFactory) Vtbl() *ILoggingChannelOptionsFactoryVtbl {
	return (*ILoggingChannelOptionsFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILoggingChannelOptionsFactory) Create(group syscall.GUID) *ILoggingChannelOptions {
	var _result *ILoggingChannelOptions
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&group)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D7F6B7AF-762D-4579-83BD-52C23BC333BC
var IID_ILoggingFields = syscall.GUID{0xD7F6B7AF, 0x762D, 0x4579,
	[8]byte{0x83, 0xBD, 0x52, 0xC2, 0x3B, 0xC3, 0x33, 0xBC}}

type ILoggingFieldsInterface interface {
	win32.IInspectableInterface
	Clear()
	BeginStruct(name string)
	BeginStructWithTags(name string, tags int32)
	EndStruct()
	AddEmpty(name string)
	AddEmptyWithFormat(name string, format LoggingFieldFormat)
	AddEmptyWithFormatAndTags(name string, format LoggingFieldFormat, tags int32)
	AddUInt8(name string, value byte)
	AddUInt8WithFormat(name string, value byte, format LoggingFieldFormat)
	AddUInt8WithFormatAndTags(name string, value byte, format LoggingFieldFormat, tags int32)
	AddUInt8Array(name string, valueLength uint32, value *byte)
	AddUInt8ArrayWithFormat(name string, valueLength uint32, value *byte, format LoggingFieldFormat)
	AddUInt8ArrayWithFormatAndTags(name string, valueLength uint32, value *byte, format LoggingFieldFormat, tags int32)
	AddInt16(name string, value int16)
	AddInt16WithFormat(name string, value int16, format LoggingFieldFormat)
	AddInt16WithFormatAndTags(name string, value int16, format LoggingFieldFormat, tags int32)
	AddInt16Array(name string, valueLength uint32, value *int16)
	AddInt16ArrayWithFormat(name string, valueLength uint32, value *int16, format LoggingFieldFormat)
	AddInt16ArrayWithFormatAndTags(name string, valueLength uint32, value *int16, format LoggingFieldFormat, tags int32)
	AddUInt16(name string, value uint16)
	AddUInt16WithFormat(name string, value uint16, format LoggingFieldFormat)
	AddUInt16WithFormatAndTags(name string, value uint16, format LoggingFieldFormat, tags int32)
	AddUInt16Array(name string, valueLength uint32, value *uint16)
	AddUInt16ArrayWithFormat(name string, valueLength uint32, value *uint16, format LoggingFieldFormat)
	AddUInt16ArrayWithFormatAndTags(name string, valueLength uint32, value *uint16, format LoggingFieldFormat, tags int32)
	AddInt32(name string, value int32)
	AddInt32WithFormat(name string, value int32, format LoggingFieldFormat)
	AddInt32WithFormatAndTags(name string, value int32, format LoggingFieldFormat, tags int32)
	AddInt32Array(name string, valueLength uint32, value *int32)
	AddInt32ArrayWithFormat(name string, valueLength uint32, value *int32, format LoggingFieldFormat)
	AddInt32ArrayWithFormatAndTags(name string, valueLength uint32, value *int32, format LoggingFieldFormat, tags int32)
	AddUInt32(name string, value uint32)
	AddUInt32WithFormat(name string, value uint32, format LoggingFieldFormat)
	AddUInt32WithFormatAndTags(name string, value uint32, format LoggingFieldFormat, tags int32)
	AddUInt32Array(name string, valueLength uint32, value *uint32)
	AddUInt32ArrayWithFormat(name string, valueLength uint32, value *uint32, format LoggingFieldFormat)
	AddUInt32ArrayWithFormatAndTags(name string, valueLength uint32, value *uint32, format LoggingFieldFormat, tags int32)
	AddInt64(name string, value int64)
	AddInt64WithFormat(name string, value int64, format LoggingFieldFormat)
	AddInt64WithFormatAndTags(name string, value int64, format LoggingFieldFormat, tags int32)
	AddInt64Array(name string, valueLength uint32, value *int64)
	AddInt64ArrayWithFormat(name string, valueLength uint32, value *int64, format LoggingFieldFormat)
	AddInt64ArrayWithFormatAndTags(name string, valueLength uint32, value *int64, format LoggingFieldFormat, tags int32)
	AddUInt64(name string, value uint64)
	AddUInt64WithFormat(name string, value uint64, format LoggingFieldFormat)
	AddUInt64WithFormatAndTags(name string, value uint64, format LoggingFieldFormat, tags int32)
	AddUInt64Array(name string, valueLength uint32, value *uint64)
	AddUInt64ArrayWithFormat(name string, valueLength uint32, value *uint64, format LoggingFieldFormat)
	AddUInt64ArrayWithFormatAndTags(name string, valueLength uint32, value *uint64, format LoggingFieldFormat, tags int32)
	AddSingle(name string, value float32)
	AddSingleWithFormat(name string, value float32, format LoggingFieldFormat)
	AddSingleWithFormatAndTags(name string, value float32, format LoggingFieldFormat, tags int32)
	AddSingleArray(name string, valueLength uint32, value *float32)
	AddSingleArrayWithFormat(name string, valueLength uint32, value *float32, format LoggingFieldFormat)
	AddSingleArrayWithFormatAndTags(name string, valueLength uint32, value *float32, format LoggingFieldFormat, tags int32)
	AddDouble(name string, value float64)
	AddDoubleWithFormat(name string, value float64, format LoggingFieldFormat)
	AddDoubleWithFormatAndTags(name string, value float64, format LoggingFieldFormat, tags int32)
	AddDoubleArray(name string, valueLength uint32, value *float64)
	AddDoubleArrayWithFormat(name string, valueLength uint32, value *float64, format LoggingFieldFormat)
	AddDoubleArrayWithFormatAndTags(name string, valueLength uint32, value *float64, format LoggingFieldFormat, tags int32)
	AddChar16(name string, value uint16)
	AddChar16WithFormat(name string, value uint16, format LoggingFieldFormat)
	AddChar16WithFormatAndTags(name string, value uint16, format LoggingFieldFormat, tags int32)
	AddChar16Array(name string, valueLength uint32, value *uint16)
	AddChar16ArrayWithFormat(name string, valueLength uint32, value *uint16, format LoggingFieldFormat)
	AddChar16ArrayWithFormatAndTags(name string, valueLength uint32, value *uint16, format LoggingFieldFormat, tags int32)
	AddBoolean(name string, value bool)
	AddBooleanWithFormat(name string, value bool, format LoggingFieldFormat)
	AddBooleanWithFormatAndTags(name string, value bool, format LoggingFieldFormat, tags int32)
	AddBooleanArray(name string, valueLength uint32, value *bool)
	AddBooleanArrayWithFormat(name string, valueLength uint32, value *bool, format LoggingFieldFormat)
	AddBooleanArrayWithFormatAndTags(name string, valueLength uint32, value *bool, format LoggingFieldFormat, tags int32)
	AddString(name string, value string)
	AddStringWithFormat(name string, value string, format LoggingFieldFormat)
	AddStringWithFormatAndTags(name string, value string, format LoggingFieldFormat, tags int32)
	AddStringArray(name string, valueLength uint32, value *string)
	AddStringArrayWithFormat(name string, valueLength uint32, value *string, format LoggingFieldFormat)
	AddStringArrayWithFormatAndTags(name string, valueLength uint32, value *string, format LoggingFieldFormat, tags int32)
	AddGuid(name string, value syscall.GUID)
	AddGuidWithFormat(name string, value syscall.GUID, format LoggingFieldFormat)
	AddGuidWithFormatAndTags(name string, value syscall.GUID, format LoggingFieldFormat, tags int32)
	AddGuidArray(name string, valueLength uint32, value *syscall.GUID)
	AddGuidArrayWithFormat(name string, valueLength uint32, value *syscall.GUID, format LoggingFieldFormat)
	AddGuidArrayWithFormatAndTags(name string, valueLength uint32, value *syscall.GUID, format LoggingFieldFormat, tags int32)
	AddDateTime(name string, value DateTime)
	AddDateTimeWithFormat(name string, value DateTime, format LoggingFieldFormat)
	AddDateTimeWithFormatAndTags(name string, value DateTime, format LoggingFieldFormat, tags int32)
	AddDateTimeArray(name string, valueLength uint32, value *DateTime)
	AddDateTimeArrayWithFormat(name string, valueLength uint32, value *DateTime, format LoggingFieldFormat)
	AddDateTimeArrayWithFormatAndTags(name string, valueLength uint32, value *DateTime, format LoggingFieldFormat, tags int32)
	AddTimeSpan(name string, value TimeSpan)
	AddTimeSpanWithFormat(name string, value TimeSpan, format LoggingFieldFormat)
	AddTimeSpanWithFormatAndTags(name string, value TimeSpan, format LoggingFieldFormat, tags int32)
	AddTimeSpanArray(name string, valueLength uint32, value *TimeSpan)
	AddTimeSpanArrayWithFormat(name string, valueLength uint32, value *TimeSpan, format LoggingFieldFormat)
	AddTimeSpanArrayWithFormatAndTags(name string, valueLength uint32, value *TimeSpan, format LoggingFieldFormat, tags int32)
	AddPoint(name string, value Point)
	AddPointWithFormat(name string, value Point, format LoggingFieldFormat)
	AddPointWithFormatAndTags(name string, value Point, format LoggingFieldFormat, tags int32)
	AddPointArray(name string, valueLength uint32, value *Point)
	AddPointArrayWithFormat(name string, valueLength uint32, value *Point, format LoggingFieldFormat)
	AddPointArrayWithFormatAndTags(name string, valueLength uint32, value *Point, format LoggingFieldFormat, tags int32)
	AddSize(name string, value Size)
	AddSizeWithFormat(name string, value Size, format LoggingFieldFormat)
	AddSizeWithFormatAndTags(name string, value Size, format LoggingFieldFormat, tags int32)
	AddSizeArray(name string, valueLength uint32, value *Size)
	AddSizeArrayWithFormat(name string, valueLength uint32, value *Size, format LoggingFieldFormat)
	AddSizeArrayWithFormatAndTags(name string, valueLength uint32, value *Size, format LoggingFieldFormat, tags int32)
	AddRect(name string, value Rect)
	AddRectWithFormat(name string, value Rect, format LoggingFieldFormat)
	AddRectWithFormatAndTags(name string, value Rect, format LoggingFieldFormat, tags int32)
	AddRectArray(name string, valueLength uint32, value *Rect)
	AddRectArrayWithFormat(name string, valueLength uint32, value *Rect, format LoggingFieldFormat)
	AddRectArrayWithFormatAndTags(name string, valueLength uint32, value *Rect, format LoggingFieldFormat, tags int32)
}

type ILoggingFieldsVtbl struct {
	win32.IInspectableVtbl
	Clear                             uintptr
	BeginStruct                       uintptr
	BeginStructWithTags               uintptr
	EndStruct                         uintptr
	AddEmpty                          uintptr
	AddEmptyWithFormat                uintptr
	AddEmptyWithFormatAndTags         uintptr
	AddUInt8                          uintptr
	AddUInt8WithFormat                uintptr
	AddUInt8WithFormatAndTags         uintptr
	AddUInt8Array                     uintptr
	AddUInt8ArrayWithFormat           uintptr
	AddUInt8ArrayWithFormatAndTags    uintptr
	AddInt16                          uintptr
	AddInt16WithFormat                uintptr
	AddInt16WithFormatAndTags         uintptr
	AddInt16Array                     uintptr
	AddInt16ArrayWithFormat           uintptr
	AddInt16ArrayWithFormatAndTags    uintptr
	AddUInt16                         uintptr
	AddUInt16WithFormat               uintptr
	AddUInt16WithFormatAndTags        uintptr
	AddUInt16Array                    uintptr
	AddUInt16ArrayWithFormat          uintptr
	AddUInt16ArrayWithFormatAndTags   uintptr
	AddInt32                          uintptr
	AddInt32WithFormat                uintptr
	AddInt32WithFormatAndTags         uintptr
	AddInt32Array                     uintptr
	AddInt32ArrayWithFormat           uintptr
	AddInt32ArrayWithFormatAndTags    uintptr
	AddUInt32                         uintptr
	AddUInt32WithFormat               uintptr
	AddUInt32WithFormatAndTags        uintptr
	AddUInt32Array                    uintptr
	AddUInt32ArrayWithFormat          uintptr
	AddUInt32ArrayWithFormatAndTags   uintptr
	AddInt64                          uintptr
	AddInt64WithFormat                uintptr
	AddInt64WithFormatAndTags         uintptr
	AddInt64Array                     uintptr
	AddInt64ArrayWithFormat           uintptr
	AddInt64ArrayWithFormatAndTags    uintptr
	AddUInt64                         uintptr
	AddUInt64WithFormat               uintptr
	AddUInt64WithFormatAndTags        uintptr
	AddUInt64Array                    uintptr
	AddUInt64ArrayWithFormat          uintptr
	AddUInt64ArrayWithFormatAndTags   uintptr
	AddSingle                         uintptr
	AddSingleWithFormat               uintptr
	AddSingleWithFormatAndTags        uintptr
	AddSingleArray                    uintptr
	AddSingleArrayWithFormat          uintptr
	AddSingleArrayWithFormatAndTags   uintptr
	AddDouble                         uintptr
	AddDoubleWithFormat               uintptr
	AddDoubleWithFormatAndTags        uintptr
	AddDoubleArray                    uintptr
	AddDoubleArrayWithFormat          uintptr
	AddDoubleArrayWithFormatAndTags   uintptr
	AddChar16                         uintptr
	AddChar16WithFormat               uintptr
	AddChar16WithFormatAndTags        uintptr
	AddChar16Array                    uintptr
	AddChar16ArrayWithFormat          uintptr
	AddChar16ArrayWithFormatAndTags   uintptr
	AddBoolean                        uintptr
	AddBooleanWithFormat              uintptr
	AddBooleanWithFormatAndTags       uintptr
	AddBooleanArray                   uintptr
	AddBooleanArrayWithFormat         uintptr
	AddBooleanArrayWithFormatAndTags  uintptr
	AddString                         uintptr
	AddStringWithFormat               uintptr
	AddStringWithFormatAndTags        uintptr
	AddStringArray                    uintptr
	AddStringArrayWithFormat          uintptr
	AddStringArrayWithFormatAndTags   uintptr
	AddGuid                           uintptr
	AddGuidWithFormat                 uintptr
	AddGuidWithFormatAndTags          uintptr
	AddGuidArray                      uintptr
	AddGuidArrayWithFormat            uintptr
	AddGuidArrayWithFormatAndTags     uintptr
	AddDateTime                       uintptr
	AddDateTimeWithFormat             uintptr
	AddDateTimeWithFormatAndTags      uintptr
	AddDateTimeArray                  uintptr
	AddDateTimeArrayWithFormat        uintptr
	AddDateTimeArrayWithFormatAndTags uintptr
	AddTimeSpan                       uintptr
	AddTimeSpanWithFormat             uintptr
	AddTimeSpanWithFormatAndTags      uintptr
	AddTimeSpanArray                  uintptr
	AddTimeSpanArrayWithFormat        uintptr
	AddTimeSpanArrayWithFormatAndTags uintptr
	AddPoint                          uintptr
	AddPointWithFormat                uintptr
	AddPointWithFormatAndTags         uintptr
	AddPointArray                     uintptr
	AddPointArrayWithFormat           uintptr
	AddPointArrayWithFormatAndTags    uintptr
	AddSize                           uintptr
	AddSizeWithFormat                 uintptr
	AddSizeWithFormatAndTags          uintptr
	AddSizeArray                      uintptr
	AddSizeArrayWithFormat            uintptr
	AddSizeArrayWithFormatAndTags     uintptr
	AddRect                           uintptr
	AddRectWithFormat                 uintptr
	AddRectWithFormatAndTags          uintptr
	AddRectArray                      uintptr
	AddRectArrayWithFormat            uintptr
	AddRectArrayWithFormatAndTags     uintptr
}

type ILoggingFields struct {
	win32.IInspectable
}

func (this *ILoggingFields) Vtbl() *ILoggingFieldsVtbl {
	return (*ILoggingFieldsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILoggingFields) Clear() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Clear, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *ILoggingFields) BeginStruct(name string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().BeginStruct, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr)
	_ = _hr
}

func (this *ILoggingFields) BeginStructWithTags(name string, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().BeginStructWithTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) EndStruct() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EndStruct, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *ILoggingFields) AddEmpty(name string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddEmpty, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr)
	_ = _hr
}

func (this *ILoggingFields) AddEmptyWithFormat(name string, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddEmptyWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddEmptyWithFormatAndTags(name string, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddEmptyWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddUInt8(name string, value byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddUInt8, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value))
	_ = _hr
}

func (this *ILoggingFields) AddUInt8WithFormat(name string, value byte, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddUInt8WithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddUInt8WithFormatAndTags(name string, value byte, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddUInt8WithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddUInt8Array(name string, valueLength uint32, value *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddUInt8Array, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ILoggingFields) AddUInt8ArrayWithFormat(name string, valueLength uint32, value *byte, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddUInt8ArrayWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddUInt8ArrayWithFormatAndTags(name string, valueLength uint32, value *byte, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddUInt8ArrayWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddInt16(name string, value int16) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddInt16, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value))
	_ = _hr
}

func (this *ILoggingFields) AddInt16WithFormat(name string, value int16, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddInt16WithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddInt16WithFormatAndTags(name string, value int16, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddInt16WithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddInt16Array(name string, valueLength uint32, value *int16) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddInt16Array, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ILoggingFields) AddInt16ArrayWithFormat(name string, valueLength uint32, value *int16, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddInt16ArrayWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddInt16ArrayWithFormatAndTags(name string, valueLength uint32, value *int16, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddInt16ArrayWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddUInt16(name string, value uint16) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddUInt16, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value))
	_ = _hr
}

func (this *ILoggingFields) AddUInt16WithFormat(name string, value uint16, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddUInt16WithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddUInt16WithFormatAndTags(name string, value uint16, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddUInt16WithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddUInt16Array(name string, valueLength uint32, value *uint16) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddUInt16Array, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ILoggingFields) AddUInt16ArrayWithFormat(name string, valueLength uint32, value *uint16, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddUInt16ArrayWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddUInt16ArrayWithFormatAndTags(name string, valueLength uint32, value *uint16, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddUInt16ArrayWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddInt32(name string, value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddInt32, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value))
	_ = _hr
}

func (this *ILoggingFields) AddInt32WithFormat(name string, value int32, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddInt32WithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddInt32WithFormatAndTags(name string, value int32, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddInt32WithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddInt32Array(name string, valueLength uint32, value *int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddInt32Array, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ILoggingFields) AddInt32ArrayWithFormat(name string, valueLength uint32, value *int32, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddInt32ArrayWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddInt32ArrayWithFormatAndTags(name string, valueLength uint32, value *int32, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddInt32ArrayWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddUInt32(name string, value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddUInt32, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value))
	_ = _hr
}

func (this *ILoggingFields) AddUInt32WithFormat(name string, value uint32, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddUInt32WithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddUInt32WithFormatAndTags(name string, value uint32, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddUInt32WithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddUInt32Array(name string, valueLength uint32, value *uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddUInt32Array, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ILoggingFields) AddUInt32ArrayWithFormat(name string, valueLength uint32, value *uint32, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddUInt32ArrayWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddUInt32ArrayWithFormatAndTags(name string, valueLength uint32, value *uint32, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddUInt32ArrayWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddInt64(name string, value int64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddInt64, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value))
	_ = _hr
}

func (this *ILoggingFields) AddInt64WithFormat(name string, value int64, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddInt64WithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddInt64WithFormatAndTags(name string, value int64, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddInt64WithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddInt64Array(name string, valueLength uint32, value *int64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddInt64Array, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ILoggingFields) AddInt64ArrayWithFormat(name string, valueLength uint32, value *int64, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddInt64ArrayWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddInt64ArrayWithFormatAndTags(name string, valueLength uint32, value *int64, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddInt64ArrayWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddUInt64(name string, value uint64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddUInt64, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value))
	_ = _hr
}

func (this *ILoggingFields) AddUInt64WithFormat(name string, value uint64, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddUInt64WithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddUInt64WithFormatAndTags(name string, value uint64, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddUInt64WithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddUInt64Array(name string, valueLength uint32, value *uint64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddUInt64Array, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ILoggingFields) AddUInt64ArrayWithFormat(name string, valueLength uint32, value *uint64, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddUInt64ArrayWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddUInt64ArrayWithFormatAndTags(name string, valueLength uint32, value *uint64, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddUInt64ArrayWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddSingle(name string, value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddSingle, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value))
	_ = _hr
}

func (this *ILoggingFields) AddSingleWithFormat(name string, value float32, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddSingleWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddSingleWithFormatAndTags(name string, value float32, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddSingleWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddSingleArray(name string, valueLength uint32, value *float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddSingleArray, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ILoggingFields) AddSingleArrayWithFormat(name string, valueLength uint32, value *float32, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddSingleArrayWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddSingleArrayWithFormatAndTags(name string, valueLength uint32, value *float32, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddSingleArrayWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddDouble(name string, value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddDouble, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value))
	_ = _hr
}

func (this *ILoggingFields) AddDoubleWithFormat(name string, value float64, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddDoubleWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddDoubleWithFormatAndTags(name string, value float64, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddDoubleWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddDoubleArray(name string, valueLength uint32, value *float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddDoubleArray, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ILoggingFields) AddDoubleArrayWithFormat(name string, valueLength uint32, value *float64, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddDoubleArrayWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddDoubleArrayWithFormatAndTags(name string, valueLength uint32, value *float64, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddDoubleArrayWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddChar16(name string, value uint16) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddChar16, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value))
	_ = _hr
}

func (this *ILoggingFields) AddChar16WithFormat(name string, value uint16, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddChar16WithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddChar16WithFormatAndTags(name string, value uint16, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddChar16WithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddChar16Array(name string, valueLength uint32, value *uint16) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddChar16Array, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ILoggingFields) AddChar16ArrayWithFormat(name string, valueLength uint32, value *uint16, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddChar16ArrayWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddChar16ArrayWithFormatAndTags(name string, valueLength uint32, value *uint16, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddChar16ArrayWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddBoolean(name string, value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddBoolean, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ILoggingFields) AddBooleanWithFormat(name string, value bool, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddBooleanWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(*(*byte)(unsafe.Pointer(&value))), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddBooleanWithFormatAndTags(name string, value bool, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddBooleanWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(*(*byte)(unsafe.Pointer(&value))), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddBooleanArray(name string, valueLength uint32, value *bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddBooleanArray, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ILoggingFields) AddBooleanArrayWithFormat(name string, valueLength uint32, value *bool, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddBooleanArrayWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddBooleanArrayWithFormatAndTags(name string, valueLength uint32, value *bool, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddBooleanArrayWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddString(name string, value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddString, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, NewHStr(value).Ptr)
	_ = _hr
}

func (this *ILoggingFields) AddStringWithFormat(name string, value string, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddStringWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, NewHStr(value).Ptr, uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddStringWithFormatAndTags(name string, value string, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddStringWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, NewHStr(value).Ptr, uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddStringArray(name string, valueLength uint32, value *string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddStringArray, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ILoggingFields) AddStringArrayWithFormat(name string, valueLength uint32, value *string, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddStringArrayWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddStringArrayWithFormatAndTags(name string, valueLength uint32, value *string, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddStringArrayWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddGuid(name string, value syscall.GUID) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddGuid, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ILoggingFields) AddGuidWithFormat(name string, value syscall.GUID, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddGuidWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&value)), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddGuidWithFormatAndTags(name string, value syscall.GUID, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddGuidWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&value)), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddGuidArray(name string, valueLength uint32, value *syscall.GUID) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddGuidArray, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ILoggingFields) AddGuidArrayWithFormat(name string, valueLength uint32, value *syscall.GUID, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddGuidArrayWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddGuidArrayWithFormatAndTags(name string, valueLength uint32, value *syscall.GUID, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddGuidArrayWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddDateTime(name string, value DateTime) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddDateTime, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ILoggingFields) AddDateTimeWithFormat(name string, value DateTime, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddDateTimeWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, *(*uintptr)(unsafe.Pointer(&value)), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddDateTimeWithFormatAndTags(name string, value DateTime, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddDateTimeWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, *(*uintptr)(unsafe.Pointer(&value)), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddDateTimeArray(name string, valueLength uint32, value *DateTime) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddDateTimeArray, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ILoggingFields) AddDateTimeArrayWithFormat(name string, valueLength uint32, value *DateTime, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddDateTimeArrayWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddDateTimeArrayWithFormatAndTags(name string, valueLength uint32, value *DateTime, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddDateTimeArrayWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddTimeSpan(name string, value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddTimeSpan, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ILoggingFields) AddTimeSpanWithFormat(name string, value TimeSpan, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddTimeSpanWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, *(*uintptr)(unsafe.Pointer(&value)), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddTimeSpanWithFormatAndTags(name string, value TimeSpan, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddTimeSpanWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, *(*uintptr)(unsafe.Pointer(&value)), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddTimeSpanArray(name string, valueLength uint32, value *TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddTimeSpanArray, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ILoggingFields) AddTimeSpanArrayWithFormat(name string, valueLength uint32, value *TimeSpan, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddTimeSpanArrayWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddTimeSpanArrayWithFormatAndTags(name string, valueLength uint32, value *TimeSpan, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddTimeSpanArrayWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddPoint(name string, value Point) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddPoint, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ILoggingFields) AddPointWithFormat(name string, value Point, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddPointWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, *(*uintptr)(unsafe.Pointer(&value)), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddPointWithFormatAndTags(name string, value Point, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddPointWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, *(*uintptr)(unsafe.Pointer(&value)), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddPointArray(name string, valueLength uint32, value *Point) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddPointArray, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ILoggingFields) AddPointArrayWithFormat(name string, valueLength uint32, value *Point, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddPointArrayWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddPointArrayWithFormatAndTags(name string, valueLength uint32, value *Point, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddPointArrayWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddSize(name string, value Size) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddSize, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ILoggingFields) AddSizeWithFormat(name string, value Size, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddSizeWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, *(*uintptr)(unsafe.Pointer(&value)), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddSizeWithFormatAndTags(name string, value Size, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddSizeWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, *(*uintptr)(unsafe.Pointer(&value)), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddSizeArray(name string, valueLength uint32, value *Size) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddSizeArray, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ILoggingFields) AddSizeArrayWithFormat(name string, valueLength uint32, value *Size, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddSizeArrayWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddSizeArrayWithFormatAndTags(name string, valueLength uint32, value *Size, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddSizeArrayWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddRect(name string, value Rect) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddRect, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ILoggingFields) AddRectWithFormat(name string, value Rect, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddRectWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&value)), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddRectWithFormatAndTags(name string, value Rect, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddRectWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&value)), uintptr(format), uintptr(tags))
	_ = _hr
}

func (this *ILoggingFields) AddRectArray(name string, valueLength uint32, value *Rect) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddRectArray, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ILoggingFields) AddRectArrayWithFormat(name string, valueLength uint32, value *Rect, format LoggingFieldFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddRectArrayWithFormat, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format))
	_ = _hr
}

func (this *ILoggingFields) AddRectArrayWithFormatAndTags(name string, valueLength uint32, value *Rect, format LoggingFieldFormat, tags int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddRectArrayWithFormatAndTags, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(format), uintptr(tags))
	_ = _hr
}

// 90BC7850-0192-4F5D-AC26-006ADACA12D8
var IID_ILoggingOptions = syscall.GUID{0x90BC7850, 0x0192, 0x4F5D,
	[8]byte{0xAC, 0x26, 0x00, 0x6A, 0xDA, 0xCA, 0x12, 0xD8}}

type ILoggingOptionsInterface interface {
	win32.IInspectableInterface
	Get_Keywords() int64
	Put_Keywords(value int64)
	Get_Tags() int32
	Put_Tags(value int32)
	Get_Task() int16
	Put_Task(value int16)
	Get_Opcode() LoggingOpcode
	Put_Opcode(value LoggingOpcode)
	Get_ActivityId() syscall.GUID
	Put_ActivityId(value syscall.GUID)
	Get_RelatedActivityId() syscall.GUID
	Put_RelatedActivityId(value syscall.GUID)
}

type ILoggingOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_Keywords          uintptr
	Put_Keywords          uintptr
	Get_Tags              uintptr
	Put_Tags              uintptr
	Get_Task              uintptr
	Put_Task              uintptr
	Get_Opcode            uintptr
	Put_Opcode            uintptr
	Get_ActivityId        uintptr
	Put_ActivityId        uintptr
	Get_RelatedActivityId uintptr
	Put_RelatedActivityId uintptr
}

type ILoggingOptions struct {
	win32.IInspectable
}

func (this *ILoggingOptions) Vtbl() *ILoggingOptionsVtbl {
	return (*ILoggingOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILoggingOptions) Get_Keywords() int64 {
	var _result int64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Keywords, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILoggingOptions) Put_Keywords(value int64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Keywords, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ILoggingOptions) Get_Tags() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Tags, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILoggingOptions) Put_Tags(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Tags, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ILoggingOptions) Get_Task() int16 {
	var _result int16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Task, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILoggingOptions) Put_Task(value int16) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Task, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ILoggingOptions) Get_Opcode() LoggingOpcode {
	var _result LoggingOpcode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Opcode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILoggingOptions) Put_Opcode(value LoggingOpcode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Opcode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ILoggingOptions) Get_ActivityId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ActivityId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILoggingOptions) Put_ActivityId(value syscall.GUID) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ActivityId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ILoggingOptions) Get_RelatedActivityId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RelatedActivityId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILoggingOptions) Put_RelatedActivityId(value syscall.GUID) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RelatedActivityId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

// D713C6CB-98AB-464B-9F22-A3268478368A
var IID_ILoggingOptionsFactory = syscall.GUID{0xD713C6CB, 0x98AB, 0x464B,
	[8]byte{0x9F, 0x22, 0xA3, 0x26, 0x84, 0x78, 0x36, 0x8A}}

type ILoggingOptionsFactoryInterface interface {
	win32.IInspectableInterface
	CreateWithKeywords(keywords int64) *ILoggingOptions
}

type ILoggingOptionsFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateWithKeywords uintptr
}

type ILoggingOptionsFactory struct {
	win32.IInspectable
}

func (this *ILoggingOptionsFactory) Vtbl() *ILoggingOptionsFactoryVtbl {
	return (*ILoggingOptionsFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILoggingOptionsFactory) CreateWithKeywords(keywords int64) *ILoggingOptions {
	var _result *ILoggingOptions
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithKeywords, uintptr(unsafe.Pointer(this)), uintptr(keywords), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6221F306-9380-4AD7-BAF5-41EA9310D768
var IID_ILoggingSession = syscall.GUID{0x6221F306, 0x9380, 0x4AD7,
	[8]byte{0xBA, 0xF5, 0x41, 0xEA, 0x93, 0x10, 0xD7, 0x68}}

type ILoggingSessionInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	SaveToFileAsync(folder *IStorageFolder, fileName string) *IAsyncOperation[*IStorageFile]
	AddLoggingChannel(loggingChannel *ILoggingChannel)
	AddLoggingChannelWithLevel(loggingChannel *ILoggingChannel, maxLevel LoggingLevel)
	RemoveLoggingChannel(loggingChannel *ILoggingChannel)
}

type ILoggingSessionVtbl struct {
	win32.IInspectableVtbl
	Get_Name                   uintptr
	SaveToFileAsync            uintptr
	AddLoggingChannel          uintptr
	AddLoggingChannelWithLevel uintptr
	RemoveLoggingChannel       uintptr
}

type ILoggingSession struct {
	win32.IInspectable
}

func (this *ILoggingSession) Vtbl() *ILoggingSessionVtbl {
	return (*ILoggingSessionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILoggingSession) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILoggingSession) SaveToFileAsync(folder *IStorageFolder, fileName string) *IAsyncOperation[*IStorageFile] {
	var _result *IAsyncOperation[*IStorageFile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SaveToFileAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(folder)), NewHStr(fileName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILoggingSession) AddLoggingChannel(loggingChannel *ILoggingChannel) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddLoggingChannel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(loggingChannel)))
	_ = _hr
}

func (this *ILoggingSession) AddLoggingChannelWithLevel(loggingChannel *ILoggingChannel, maxLevel LoggingLevel) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddLoggingChannelWithLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(loggingChannel)), uintptr(maxLevel))
	_ = _hr
}

func (this *ILoggingSession) RemoveLoggingChannel(loggingChannel *ILoggingChannel) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveLoggingChannel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(loggingChannel)))
	_ = _hr
}

// 4E937EE5-58FD-45E0-8C2F-A132EFF95C1E
var IID_ILoggingSessionFactory = syscall.GUID{0x4E937EE5, 0x58FD, 0x45E0,
	[8]byte{0x8C, 0x2F, 0xA1, 0x32, 0xEF, 0xF9, 0x5C, 0x1E}}

type ILoggingSessionFactoryInterface interface {
	win32.IInspectableInterface
	Create(name string) *ILoggingSession
}

type ILoggingSessionFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ILoggingSessionFactory struct {
	win32.IInspectable
}

func (this *ILoggingSessionFactory) Vtbl() *ILoggingSessionFactoryVtbl {
	return (*ILoggingSessionFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILoggingSessionFactory) Create(name string) *ILoggingSession {
	var _result *ILoggingSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 65F16C35-E388-4E26-B17A-F51CD3A83916
var IID_ILoggingTarget = syscall.GUID{0x65F16C35, 0xE388, 0x4E26,
	[8]byte{0xB1, 0x7A, 0xF5, 0x1C, 0xD3, 0xA8, 0x39, 0x16}}

type ILoggingTargetInterface interface {
	win32.IInspectableInterface
	IsEnabled() bool
	IsEnabledWithLevel(level LoggingLevel) bool
	IsEnabledWithLevelAndKeywords(level LoggingLevel, keywords int64) bool
	LogEvent(eventName string)
	LogEventWithFields(eventName string, fields *ILoggingFields)
	LogEventWithFieldsAndLevel(eventName string, fields *ILoggingFields, level LoggingLevel)
	LogEventWithFieldsAndOptions(eventName string, fields *ILoggingFields, level LoggingLevel, options *ILoggingOptions)
	StartActivity(startEventName string) *ILoggingActivity
	StartActivityWithFields(startEventName string, fields *ILoggingFields) *ILoggingActivity
	StartActivityWithFieldsAndLevel(startEventName string, fields *ILoggingFields, level LoggingLevel) *ILoggingActivity
	StartActivityWithFieldsAndOptions(startEventName string, fields *ILoggingFields, level LoggingLevel, options *ILoggingOptions) *ILoggingActivity
}

type ILoggingTargetVtbl struct {
	win32.IInspectableVtbl
	IsEnabled                         uintptr
	IsEnabledWithLevel                uintptr
	IsEnabledWithLevelAndKeywords     uintptr
	LogEvent                          uintptr
	LogEventWithFields                uintptr
	LogEventWithFieldsAndLevel        uintptr
	LogEventWithFieldsAndOptions      uintptr
	StartActivity                     uintptr
	StartActivityWithFields           uintptr
	StartActivityWithFieldsAndLevel   uintptr
	StartActivityWithFieldsAndOptions uintptr
}

type ILoggingTarget struct {
	win32.IInspectable
}

func (this *ILoggingTarget) Vtbl() *ILoggingTargetVtbl {
	return (*ILoggingTargetVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILoggingTarget) IsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILoggingTarget) IsEnabledWithLevel(level LoggingLevel) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsEnabledWithLevel, uintptr(unsafe.Pointer(this)), uintptr(level), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILoggingTarget) IsEnabledWithLevelAndKeywords(level LoggingLevel, keywords int64) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsEnabledWithLevelAndKeywords, uintptr(unsafe.Pointer(this)), uintptr(level), uintptr(keywords), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILoggingTarget) LogEvent(eventName string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LogEvent, uintptr(unsafe.Pointer(this)), NewHStr(eventName).Ptr)
	_ = _hr
}

func (this *ILoggingTarget) LogEventWithFields(eventName string, fields *ILoggingFields) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LogEventWithFields, uintptr(unsafe.Pointer(this)), NewHStr(eventName).Ptr, uintptr(unsafe.Pointer(fields)))
	_ = _hr
}

func (this *ILoggingTarget) LogEventWithFieldsAndLevel(eventName string, fields *ILoggingFields, level LoggingLevel) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LogEventWithFieldsAndLevel, uintptr(unsafe.Pointer(this)), NewHStr(eventName).Ptr, uintptr(unsafe.Pointer(fields)), uintptr(level))
	_ = _hr
}

func (this *ILoggingTarget) LogEventWithFieldsAndOptions(eventName string, fields *ILoggingFields, level LoggingLevel, options *ILoggingOptions) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LogEventWithFieldsAndOptions, uintptr(unsafe.Pointer(this)), NewHStr(eventName).Ptr, uintptr(unsafe.Pointer(fields)), uintptr(level), uintptr(unsafe.Pointer(options)))
	_ = _hr
}

func (this *ILoggingTarget) StartActivity(startEventName string) *ILoggingActivity {
	var _result *ILoggingActivity
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartActivity, uintptr(unsafe.Pointer(this)), NewHStr(startEventName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILoggingTarget) StartActivityWithFields(startEventName string, fields *ILoggingFields) *ILoggingActivity {
	var _result *ILoggingActivity
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartActivityWithFields, uintptr(unsafe.Pointer(this)), NewHStr(startEventName).Ptr, uintptr(unsafe.Pointer(fields)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILoggingTarget) StartActivityWithFieldsAndLevel(startEventName string, fields *ILoggingFields, level LoggingLevel) *ILoggingActivity {
	var _result *ILoggingActivity
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartActivityWithFieldsAndLevel, uintptr(unsafe.Pointer(this)), NewHStr(startEventName).Ptr, uintptr(unsafe.Pointer(fields)), uintptr(level), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILoggingTarget) StartActivityWithFieldsAndOptions(startEventName string, fields *ILoggingFields, level LoggingLevel, options *ILoggingOptions) *ILoggingActivity {
	var _result *ILoggingActivity
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartActivityWithFieldsAndOptions, uintptr(unsafe.Pointer(this)), NewHStr(startEventName).Ptr, uintptr(unsafe.Pointer(fields)), uintptr(level), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 410B7711-FF3B-477F-9C9A-D2EFDA302DC3
var IID_ITracingStatusChangedEventArgs = syscall.GUID{0x410B7711, 0xFF3B, 0x477F,
	[8]byte{0x9C, 0x9A, 0xD2, 0xEF, 0xDA, 0x30, 0x2D, 0xC3}}

type ITracingStatusChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Enabled() bool
	Get_TraceLevel() CausalityTraceLevel
}

type ITracingStatusChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Enabled    uintptr
	Get_TraceLevel uintptr
}

type ITracingStatusChangedEventArgs struct {
	win32.IInspectable
}

func (this *ITracingStatusChangedEventArgs) Vtbl() *ITracingStatusChangedEventArgsVtbl {
	return (*ITracingStatusChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITracingStatusChangedEventArgs) Get_Enabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Enabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITracingStatusChangedEventArgs) Get_TraceLevel() CausalityTraceLevel {
	var _result CausalityTraceLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TraceLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type AsyncCausalityTracer struct {
	RtClass
}

func NewIAsyncCausalityTracerStatics() *IAsyncCausalityTracerStatics {
	var p *IAsyncCausalityTracerStatics
	hs := NewHStr("Windows.Foundation.Diagnostics.AsyncCausalityTracer")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAsyncCausalityTracerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type RuntimeBrokerErrorSettings struct {
	RtClass
	*IErrorReportingSettings
}

func NewRuntimeBrokerErrorSettings() *RuntimeBrokerErrorSettings {
	hs := NewHStr("Windows.Foundation.Diagnostics.RuntimeBrokerErrorSettings")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &RuntimeBrokerErrorSettings{
		RtClass:                 RtClass{PInspect: p},
		IErrorReportingSettings: (*IErrorReportingSettings)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}
