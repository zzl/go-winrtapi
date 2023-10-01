package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type PlatformDiagnosticActionState int32

const (
	PlatformDiagnosticActionState_Success                 PlatformDiagnosticActionState = 0
	PlatformDiagnosticActionState_FreeNetworkNotAvailable PlatformDiagnosticActionState = 1
	PlatformDiagnosticActionState_ACPowerNotAvailable     PlatformDiagnosticActionState = 2
)

// enum
type PlatformDiagnosticEscalationType int32

const (
	PlatformDiagnosticEscalationType_OnCompletion PlatformDiagnosticEscalationType = 0
	PlatformDiagnosticEscalationType_OnFailure    PlatformDiagnosticEscalationType = 1
)

// enum
// flags
type PlatformDiagnosticEventBufferLatencies uint32

const (
	PlatformDiagnosticEventBufferLatencies_Normal       PlatformDiagnosticEventBufferLatencies = 1
	PlatformDiagnosticEventBufferLatencies_CostDeferred PlatformDiagnosticEventBufferLatencies = 2
	PlatformDiagnosticEventBufferLatencies_Realtime     PlatformDiagnosticEventBufferLatencies = 4
)

// enum
type PlatformDiagnosticTracePriority int32

const (
	PlatformDiagnosticTracePriority_Normal       PlatformDiagnosticTracePriority = 0
	PlatformDiagnosticTracePriority_UserElevated PlatformDiagnosticTracePriority = 1
)

// enum
type PlatformDiagnosticTraceSlotState int32

const (
	PlatformDiagnosticTraceSlotState_NotRunning PlatformDiagnosticTraceSlotState = 0
	PlatformDiagnosticTraceSlotState_Running    PlatformDiagnosticTraceSlotState = 1
	PlatformDiagnosticTraceSlotState_Throttled  PlatformDiagnosticTraceSlotState = 2
)

// enum
type PlatformDiagnosticTraceSlotType int32

const (
	PlatformDiagnosticTraceSlotType_Alternative PlatformDiagnosticTraceSlotType = 0
	PlatformDiagnosticTraceSlotType_AlwaysOn    PlatformDiagnosticTraceSlotType = 1
	PlatformDiagnosticTraceSlotType_Mini        PlatformDiagnosticTraceSlotType = 2
)

// interfaces

// C1145CFA-9292-4267-890A-9EA3ED072312
var IID_IPlatformDiagnosticActionsStatics = syscall.GUID{0xC1145CFA, 0x9292, 0x4267,
	[8]byte{0x89, 0x0A, 0x9E, 0xA3, 0xED, 0x07, 0x23, 0x12}}

type IPlatformDiagnosticActionsStaticsInterface interface {
	win32.IInspectableInterface
	IsScenarioEnabled(scenarioId syscall.GUID) bool
	TryEscalateScenario(scenarioId syscall.GUID, escalationType PlatformDiagnosticEscalationType, outputDirectory string, timestampOutputDirectory bool, forceEscalationUpload bool, triggers *IMapView[string, string]) bool
	DownloadLatestSettingsForNamespace(partner string, feature string, isScenarioNamespace bool, downloadOverCostedNetwork bool, downloadOverBattery bool) PlatformDiagnosticActionState
	GetActiveScenarioList() *IVectorView[syscall.GUID]
	ForceUpload(latency PlatformDiagnosticEventBufferLatencies, uploadOverCostedNetwork bool, uploadOverBattery bool) PlatformDiagnosticActionState
	IsTraceRunning(slotType PlatformDiagnosticTraceSlotType, scenarioId syscall.GUID, traceProfileHash uint64) PlatformDiagnosticTraceSlotState
	GetActiveTraceRuntime(slotType PlatformDiagnosticTraceSlotType) *IPlatformDiagnosticTraceRuntimeInfo
	GetKnownTraceList(slotType PlatformDiagnosticTraceSlotType) *IVectorView[*IPlatformDiagnosticTraceInfo]
}

type IPlatformDiagnosticActionsStaticsVtbl struct {
	win32.IInspectableVtbl
	IsScenarioEnabled                  uintptr
	TryEscalateScenario                uintptr
	DownloadLatestSettingsForNamespace uintptr
	GetActiveScenarioList              uintptr
	ForceUpload                        uintptr
	IsTraceRunning                     uintptr
	GetActiveTraceRuntime              uintptr
	GetKnownTraceList                  uintptr
}

type IPlatformDiagnosticActionsStatics struct {
	win32.IInspectable
}

func (this *IPlatformDiagnosticActionsStatics) Vtbl() *IPlatformDiagnosticActionsStaticsVtbl {
	return (*IPlatformDiagnosticActionsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPlatformDiagnosticActionsStatics) IsScenarioEnabled(scenarioId syscall.GUID) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsScenarioEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&scenarioId)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPlatformDiagnosticActionsStatics) TryEscalateScenario(scenarioId syscall.GUID, escalationType PlatformDiagnosticEscalationType, outputDirectory string, timestampOutputDirectory bool, forceEscalationUpload bool, triggers *IMapView[string, string]) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryEscalateScenario, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&scenarioId)), uintptr(escalationType), NewHStr(outputDirectory).Ptr, uintptr(*(*byte)(unsafe.Pointer(&timestampOutputDirectory))), uintptr(*(*byte)(unsafe.Pointer(&forceEscalationUpload))), uintptr(unsafe.Pointer(triggers)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPlatformDiagnosticActionsStatics) DownloadLatestSettingsForNamespace(partner string, feature string, isScenarioNamespace bool, downloadOverCostedNetwork bool, downloadOverBattery bool) PlatformDiagnosticActionState {
	var _result PlatformDiagnosticActionState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DownloadLatestSettingsForNamespace, uintptr(unsafe.Pointer(this)), NewHStr(partner).Ptr, NewHStr(feature).Ptr, uintptr(*(*byte)(unsafe.Pointer(&isScenarioNamespace))), uintptr(*(*byte)(unsafe.Pointer(&downloadOverCostedNetwork))), uintptr(*(*byte)(unsafe.Pointer(&downloadOverBattery))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPlatformDiagnosticActionsStatics) GetActiveScenarioList() *IVectorView[syscall.GUID] {
	var _result *IVectorView[syscall.GUID]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetActiveScenarioList, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPlatformDiagnosticActionsStatics) ForceUpload(latency PlatformDiagnosticEventBufferLatencies, uploadOverCostedNetwork bool, uploadOverBattery bool) PlatformDiagnosticActionState {
	var _result PlatformDiagnosticActionState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ForceUpload, uintptr(unsafe.Pointer(this)), uintptr(latency), uintptr(*(*byte)(unsafe.Pointer(&uploadOverCostedNetwork))), uintptr(*(*byte)(unsafe.Pointer(&uploadOverBattery))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPlatformDiagnosticActionsStatics) IsTraceRunning(slotType PlatformDiagnosticTraceSlotType, scenarioId syscall.GUID, traceProfileHash uint64) PlatformDiagnosticTraceSlotState {
	var _result PlatformDiagnosticTraceSlotState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsTraceRunning, uintptr(unsafe.Pointer(this)), uintptr(slotType), uintptr(unsafe.Pointer(&scenarioId)), uintptr(traceProfileHash), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPlatformDiagnosticActionsStatics) GetActiveTraceRuntime(slotType PlatformDiagnosticTraceSlotType) *IPlatformDiagnosticTraceRuntimeInfo {
	var _result *IPlatformDiagnosticTraceRuntimeInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetActiveTraceRuntime, uintptr(unsafe.Pointer(this)), uintptr(slotType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPlatformDiagnosticActionsStatics) GetKnownTraceList(slotType PlatformDiagnosticTraceSlotType) *IVectorView[*IPlatformDiagnosticTraceInfo] {
	var _result *IVectorView[*IPlatformDiagnosticTraceInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetKnownTraceList, uintptr(unsafe.Pointer(this)), uintptr(slotType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F870ED97-D597-4BF7-88DC-CF5C7DC2A1D2
var IID_IPlatformDiagnosticTraceInfo = syscall.GUID{0xF870ED97, 0xD597, 0x4BF7,
	[8]byte{0x88, 0xDC, 0xCF, 0x5C, 0x7D, 0xC2, 0xA1, 0xD2}}

type IPlatformDiagnosticTraceInfoInterface interface {
	win32.IInspectableInterface
	Get_ScenarioId() syscall.GUID
	Get_ProfileHash() uint64
	Get_IsExclusive() bool
	Get_IsAutoLogger() bool
	Get_MaxTraceDurationFileTime() int64
	Get_Priority() PlatformDiagnosticTracePriority
}

type IPlatformDiagnosticTraceInfoVtbl struct {
	win32.IInspectableVtbl
	Get_ScenarioId               uintptr
	Get_ProfileHash              uintptr
	Get_IsExclusive              uintptr
	Get_IsAutoLogger             uintptr
	Get_MaxTraceDurationFileTime uintptr
	Get_Priority                 uintptr
}

type IPlatformDiagnosticTraceInfo struct {
	win32.IInspectable
}

func (this *IPlatformDiagnosticTraceInfo) Vtbl() *IPlatformDiagnosticTraceInfoVtbl {
	return (*IPlatformDiagnosticTraceInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPlatformDiagnosticTraceInfo) Get_ScenarioId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScenarioId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPlatformDiagnosticTraceInfo) Get_ProfileHash() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProfileHash, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPlatformDiagnosticTraceInfo) Get_IsExclusive() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsExclusive, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPlatformDiagnosticTraceInfo) Get_IsAutoLogger() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsAutoLogger, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPlatformDiagnosticTraceInfo) Get_MaxTraceDurationFileTime() int64 {
	var _result int64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxTraceDurationFileTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPlatformDiagnosticTraceInfo) Get_Priority() PlatformDiagnosticTracePriority {
	var _result PlatformDiagnosticTracePriority
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Priority, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 3D4D5E2D-01D8-4768-8554-1EB1CA610986
var IID_IPlatformDiagnosticTraceRuntimeInfo = syscall.GUID{0x3D4D5E2D, 0x01D8, 0x4768,
	[8]byte{0x85, 0x54, 0x1E, 0xB1, 0xCA, 0x61, 0x09, 0x86}}

type IPlatformDiagnosticTraceRuntimeInfoInterface interface {
	win32.IInspectableInterface
	Get_RuntimeFileTime() int64
	Get_EtwRuntimeFileTime() int64
}

type IPlatformDiagnosticTraceRuntimeInfoVtbl struct {
	win32.IInspectableVtbl
	Get_RuntimeFileTime    uintptr
	Get_EtwRuntimeFileTime uintptr
}

type IPlatformDiagnosticTraceRuntimeInfo struct {
	win32.IInspectable
}

func (this *IPlatformDiagnosticTraceRuntimeInfo) Vtbl() *IPlatformDiagnosticTraceRuntimeInfoVtbl {
	return (*IPlatformDiagnosticTraceRuntimeInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPlatformDiagnosticTraceRuntimeInfo) Get_RuntimeFileTime() int64 {
	var _result int64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RuntimeFileTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPlatformDiagnosticTraceRuntimeInfo) Get_EtwRuntimeFileTime() int64 {
	var _result int64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EtwRuntimeFileTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type PlatformDiagnosticActions struct {
	RtClass
}

func NewIPlatformDiagnosticActionsStatics() *IPlatformDiagnosticActionsStatics {
	var p *IPlatformDiagnosticActionsStatics
	hs := NewHStr("Windows.System.Diagnostics.TraceReporting.PlatformDiagnosticActions")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPlatformDiagnosticActionsStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type PlatformDiagnosticTraceInfo struct {
	RtClass
	*IPlatformDiagnosticTraceInfo
}

type PlatformDiagnosticTraceRuntimeInfo struct {
	RtClass
	*IPlatformDiagnosticTraceRuntimeInfo
}
