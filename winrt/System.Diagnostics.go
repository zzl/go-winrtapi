package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type DiagnosticActionState int32

const (
	DiagnosticActionState_Initializing        DiagnosticActionState = 0
	DiagnosticActionState_Downloading         DiagnosticActionState = 1
	DiagnosticActionState_VerifyingTrust      DiagnosticActionState = 2
	DiagnosticActionState_Detecting           DiagnosticActionState = 3
	DiagnosticActionState_Resolving           DiagnosticActionState = 4
	DiagnosticActionState_VerifyingResolution DiagnosticActionState = 5
	DiagnosticActionState_Executing           DiagnosticActionState = 6
)

// interfaces

// C265A296-E73B-4097-B28F-3442F03DD831
var IID_IDiagnosticActionResult = syscall.GUID{0xC265A296, 0xE73B, 0x4097,
	[8]byte{0xB2, 0x8F, 0x34, 0x42, 0xF0, 0x3D, 0xD8, 0x31}}

type IDiagnosticActionResultInterface interface {
	win32.IInspectableInterface
	Get_ExtendedError() HResult
	Get_Results() *IPropertySet
}

type IDiagnosticActionResultVtbl struct {
	win32.IInspectableVtbl
	Get_ExtendedError uintptr
	Get_Results       uintptr
}

type IDiagnosticActionResult struct {
	win32.IInspectable
}

func (this *IDiagnosticActionResult) Vtbl() *IDiagnosticActionResultVtbl {
	return (*IDiagnosticActionResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDiagnosticActionResult) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDiagnosticActionResult) Get_Results() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Results, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 187B270A-02E3-4F86-84FC-FDD892B5940F
var IID_IDiagnosticInvoker = syscall.GUID{0x187B270A, 0x02E3, 0x4F86,
	[8]byte{0x84, 0xFC, 0xFD, 0xD8, 0x92, 0xB5, 0x94, 0x0F}}

type IDiagnosticInvokerInterface interface {
	win32.IInspectableInterface
	RunDiagnosticActionAsync(context *IJsonObject) *IAsyncOperationWithProgress[*IDiagnosticActionResult, DiagnosticActionState]
}

type IDiagnosticInvokerVtbl struct {
	win32.IInspectableVtbl
	RunDiagnosticActionAsync uintptr
}

type IDiagnosticInvoker struct {
	win32.IInspectable
}

func (this *IDiagnosticInvoker) Vtbl() *IDiagnosticInvokerVtbl {
	return (*IDiagnosticInvokerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDiagnosticInvoker) RunDiagnosticActionAsync(context *IJsonObject) *IAsyncOperationWithProgress[*IDiagnosticActionResult, DiagnosticActionState] {
	var _result *IAsyncOperationWithProgress[*IDiagnosticActionResult, DiagnosticActionState]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RunDiagnosticActionAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E3BF945C-155A-4B52-A8EC-070C44F95000
var IID_IDiagnosticInvoker2 = syscall.GUID{0xE3BF945C, 0x155A, 0x4B52,
	[8]byte{0xA8, 0xEC, 0x07, 0x0C, 0x44, 0xF9, 0x50, 0x00}}

type IDiagnosticInvoker2Interface interface {
	win32.IInspectableInterface
	RunDiagnosticActionFromStringAsync(context string) *IAsyncOperationWithProgress[*IDiagnosticActionResult, DiagnosticActionState]
}

type IDiagnosticInvoker2Vtbl struct {
	win32.IInspectableVtbl
	RunDiagnosticActionFromStringAsync uintptr
}

type IDiagnosticInvoker2 struct {
	win32.IInspectable
}

func (this *IDiagnosticInvoker2) Vtbl() *IDiagnosticInvoker2Vtbl {
	return (*IDiagnosticInvoker2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDiagnosticInvoker2) RunDiagnosticActionFromStringAsync(context string) *IAsyncOperationWithProgress[*IDiagnosticActionResult, DiagnosticActionState] {
	var _result *IAsyncOperationWithProgress[*IDiagnosticActionResult, DiagnosticActionState]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RunDiagnosticActionFromStringAsync, uintptr(unsafe.Pointer(this)), NewHStr(context).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5CFAD8DE-F15C-4554-A813-C113C3881B09
var IID_IDiagnosticInvokerStatics = syscall.GUID{0x5CFAD8DE, 0xF15C, 0x4554,
	[8]byte{0xA8, 0x13, 0xC1, 0x13, 0xC3, 0x88, 0x1B, 0x09}}

type IDiagnosticInvokerStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *IDiagnosticInvoker
	GetForUser(user *IUser) *IDiagnosticInvoker
	Get_IsSupported() bool
}

type IDiagnosticInvokerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault      uintptr
	GetForUser      uintptr
	Get_IsSupported uintptr
}

type IDiagnosticInvokerStatics struct {
	win32.IInspectable
}

func (this *IDiagnosticInvokerStatics) Vtbl() *IDiagnosticInvokerStaticsVtbl {
	return (*IDiagnosticInvokerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDiagnosticInvokerStatics) GetDefault() *IDiagnosticInvoker {
	var _result *IDiagnosticInvoker
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDiagnosticInvokerStatics) GetForUser(user *IUser) *IDiagnosticInvoker {
	var _result *IDiagnosticInvoker
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDiagnosticInvokerStatics) Get_IsSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 0BBB2472-C8BF-423A-A810-B559AE4354E2
var IID_IProcessCpuUsage = syscall.GUID{0x0BBB2472, 0xC8BF, 0x423A,
	[8]byte{0xA8, 0x10, 0xB5, 0x59, 0xAE, 0x43, 0x54, 0xE2}}

type IProcessCpuUsageInterface interface {
	win32.IInspectableInterface
	GetReport() *IProcessCpuUsageReport
}

type IProcessCpuUsageVtbl struct {
	win32.IInspectableVtbl
	GetReport uintptr
}

type IProcessCpuUsage struct {
	win32.IInspectable
}

func (this *IProcessCpuUsage) Vtbl() *IProcessCpuUsageVtbl {
	return (*IProcessCpuUsageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProcessCpuUsage) GetReport() *IProcessCpuUsageReport {
	var _result *IProcessCpuUsageReport
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetReport, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8A6D9CAC-3987-4E2F-A119-6B5FA214F1B4
var IID_IProcessCpuUsageReport = syscall.GUID{0x8A6D9CAC, 0x3987, 0x4E2F,
	[8]byte{0xA1, 0x19, 0x6B, 0x5F, 0xA2, 0x14, 0xF1, 0xB4}}

type IProcessCpuUsageReportInterface interface {
	win32.IInspectableInterface
	Get_KernelTime() TimeSpan
	Get_UserTime() TimeSpan
}

type IProcessCpuUsageReportVtbl struct {
	win32.IInspectableVtbl
	Get_KernelTime uintptr
	Get_UserTime   uintptr
}

type IProcessCpuUsageReport struct {
	win32.IInspectable
}

func (this *IProcessCpuUsageReport) Vtbl() *IProcessCpuUsageReportVtbl {
	return (*IProcessCpuUsageReportVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProcessCpuUsageReport) Get_KernelTime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KernelTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProcessCpuUsageReport) Get_UserTime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UserTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// E830B04B-300E-4EE6-A0AB-5B5F5231B434
var IID_IProcessDiagnosticInfo = syscall.GUID{0xE830B04B, 0x300E, 0x4EE6,
	[8]byte{0xA0, 0xAB, 0x5B, 0x5F, 0x52, 0x31, 0xB4, 0x34}}

type IProcessDiagnosticInfoInterface interface {
	win32.IInspectableInterface
	Get_ProcessId() uint32
	Get_ExecutableFileName() string
	Get_Parent() *IProcessDiagnosticInfo
	Get_ProcessStartTime() DateTime
	Get_DiskUsage() *IProcessDiskUsage
	Get_MemoryUsage() *IProcessMemoryUsage
	Get_CpuUsage() *IProcessCpuUsage
}

type IProcessDiagnosticInfoVtbl struct {
	win32.IInspectableVtbl
	Get_ProcessId          uintptr
	Get_ExecutableFileName uintptr
	Get_Parent             uintptr
	Get_ProcessStartTime   uintptr
	Get_DiskUsage          uintptr
	Get_MemoryUsage        uintptr
	Get_CpuUsage           uintptr
}

type IProcessDiagnosticInfo struct {
	win32.IInspectable
}

func (this *IProcessDiagnosticInfo) Vtbl() *IProcessDiagnosticInfoVtbl {
	return (*IProcessDiagnosticInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProcessDiagnosticInfo) Get_ProcessId() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProcessId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProcessDiagnosticInfo) Get_ExecutableFileName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExecutableFileName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IProcessDiagnosticInfo) Get_Parent() *IProcessDiagnosticInfo {
	var _result *IProcessDiagnosticInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Parent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProcessDiagnosticInfo) Get_ProcessStartTime() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProcessStartTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProcessDiagnosticInfo) Get_DiskUsage() *IProcessDiskUsage {
	var _result *IProcessDiskUsage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DiskUsage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProcessDiagnosticInfo) Get_MemoryUsage() *IProcessMemoryUsage {
	var _result *IProcessMemoryUsage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MemoryUsage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProcessDiagnosticInfo) Get_CpuUsage() *IProcessCpuUsage {
	var _result *IProcessCpuUsage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CpuUsage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9558CB1A-3D0B-49EC-AB70-4F7A112805DE
var IID_IProcessDiagnosticInfo2 = syscall.GUID{0x9558CB1A, 0x3D0B, 0x49EC,
	[8]byte{0xAB, 0x70, 0x4F, 0x7A, 0x11, 0x28, 0x05, 0xDE}}

type IProcessDiagnosticInfo2Interface interface {
	win32.IInspectableInterface
	GetAppDiagnosticInfos() *IVector[*IAppDiagnosticInfo]
	Get_IsPackaged() bool
}

type IProcessDiagnosticInfo2Vtbl struct {
	win32.IInspectableVtbl
	GetAppDiagnosticInfos uintptr
	Get_IsPackaged        uintptr
}

type IProcessDiagnosticInfo2 struct {
	win32.IInspectable
}

func (this *IProcessDiagnosticInfo2) Vtbl() *IProcessDiagnosticInfo2Vtbl {
	return (*IProcessDiagnosticInfo2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProcessDiagnosticInfo2) GetAppDiagnosticInfos() *IVector[*IAppDiagnosticInfo] {
	var _result *IVector[*IAppDiagnosticInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAppDiagnosticInfos, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProcessDiagnosticInfo2) Get_IsPackaged() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPackaged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 2F41B260-B49F-428C-AA0E-84744F49CA95
var IID_IProcessDiagnosticInfoStatics = syscall.GUID{0x2F41B260, 0xB49F, 0x428C,
	[8]byte{0xAA, 0x0E, 0x84, 0x74, 0x4F, 0x49, 0xCA, 0x95}}

type IProcessDiagnosticInfoStaticsInterface interface {
	win32.IInspectableInterface
	GetForProcesses() *IVectorView[*IProcessDiagnosticInfo]
	GetForCurrentProcess() *IProcessDiagnosticInfo
}

type IProcessDiagnosticInfoStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForProcesses      uintptr
	GetForCurrentProcess uintptr
}

type IProcessDiagnosticInfoStatics struct {
	win32.IInspectable
}

func (this *IProcessDiagnosticInfoStatics) Vtbl() *IProcessDiagnosticInfoStaticsVtbl {
	return (*IProcessDiagnosticInfoStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProcessDiagnosticInfoStatics) GetForProcesses() *IVectorView[*IProcessDiagnosticInfo] {
	var _result *IVectorView[*IProcessDiagnosticInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForProcesses, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProcessDiagnosticInfoStatics) GetForCurrentProcess() *IProcessDiagnosticInfo {
	var _result *IProcessDiagnosticInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentProcess, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4A869897-9899-4A44-A29B-091663BE09B6
var IID_IProcessDiagnosticInfoStatics2 = syscall.GUID{0x4A869897, 0x9899, 0x4A44,
	[8]byte{0xA2, 0x9B, 0x09, 0x16, 0x63, 0xBE, 0x09, 0xB6}}

type IProcessDiagnosticInfoStatics2Interface interface {
	win32.IInspectableInterface
	TryGetForProcessId(processId uint32) *IProcessDiagnosticInfo
}

type IProcessDiagnosticInfoStatics2Vtbl struct {
	win32.IInspectableVtbl
	TryGetForProcessId uintptr
}

type IProcessDiagnosticInfoStatics2 struct {
	win32.IInspectable
}

func (this *IProcessDiagnosticInfoStatics2) Vtbl() *IProcessDiagnosticInfoStatics2Vtbl {
	return (*IProcessDiagnosticInfoStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProcessDiagnosticInfoStatics2) TryGetForProcessId(processId uint32) *IProcessDiagnosticInfo {
	var _result *IProcessDiagnosticInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetForProcessId, uintptr(unsafe.Pointer(this)), uintptr(processId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5AD78BFD-7E51-4E53-BFAA-5A6EE1AABBF8
var IID_IProcessDiskUsage = syscall.GUID{0x5AD78BFD, 0x7E51, 0x4E53,
	[8]byte{0xBF, 0xAA, 0x5A, 0x6E, 0xE1, 0xAA, 0xBB, 0xF8}}

type IProcessDiskUsageInterface interface {
	win32.IInspectableInterface
	GetReport() *IProcessDiskUsageReport
}

type IProcessDiskUsageVtbl struct {
	win32.IInspectableVtbl
	GetReport uintptr
}

type IProcessDiskUsage struct {
	win32.IInspectable
}

func (this *IProcessDiskUsage) Vtbl() *IProcessDiskUsageVtbl {
	return (*IProcessDiskUsageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProcessDiskUsage) GetReport() *IProcessDiskUsageReport {
	var _result *IProcessDiskUsageReport
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetReport, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 401627FD-535D-4C1F-81B8-DA54E1BE635E
var IID_IProcessDiskUsageReport = syscall.GUID{0x401627FD, 0x535D, 0x4C1F,
	[8]byte{0x81, 0xB8, 0xDA, 0x54, 0xE1, 0xBE, 0x63, 0x5E}}

type IProcessDiskUsageReportInterface interface {
	win32.IInspectableInterface
	Get_ReadOperationCount() int64
	Get_WriteOperationCount() int64
	Get_OtherOperationCount() int64
	Get_BytesReadCount() int64
	Get_BytesWrittenCount() int64
	Get_OtherBytesCount() int64
}

type IProcessDiskUsageReportVtbl struct {
	win32.IInspectableVtbl
	Get_ReadOperationCount  uintptr
	Get_WriteOperationCount uintptr
	Get_OtherOperationCount uintptr
	Get_BytesReadCount      uintptr
	Get_BytesWrittenCount   uintptr
	Get_OtherBytesCount     uintptr
}

type IProcessDiskUsageReport struct {
	win32.IInspectable
}

func (this *IProcessDiskUsageReport) Vtbl() *IProcessDiskUsageReportVtbl {
	return (*IProcessDiskUsageReportVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProcessDiskUsageReport) Get_ReadOperationCount() int64 {
	var _result int64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReadOperationCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProcessDiskUsageReport) Get_WriteOperationCount() int64 {
	var _result int64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WriteOperationCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProcessDiskUsageReport) Get_OtherOperationCount() int64 {
	var _result int64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OtherOperationCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProcessDiskUsageReport) Get_BytesReadCount() int64 {
	var _result int64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BytesReadCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProcessDiskUsageReport) Get_BytesWrittenCount() int64 {
	var _result int64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BytesWrittenCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProcessDiskUsageReport) Get_OtherBytesCount() int64 {
	var _result int64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OtherBytesCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F50B229B-827C-42B7-B07C-0E32627E6B3E
var IID_IProcessMemoryUsage = syscall.GUID{0xF50B229B, 0x827C, 0x42B7,
	[8]byte{0xB0, 0x7C, 0x0E, 0x32, 0x62, 0x7E, 0x6B, 0x3E}}

type IProcessMemoryUsageInterface interface {
	win32.IInspectableInterface
	GetReport() *IProcessMemoryUsageReport
}

type IProcessMemoryUsageVtbl struct {
	win32.IInspectableVtbl
	GetReport uintptr
}

type IProcessMemoryUsage struct {
	win32.IInspectable
}

func (this *IProcessMemoryUsage) Vtbl() *IProcessMemoryUsageVtbl {
	return (*IProcessMemoryUsageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProcessMemoryUsage) GetReport() *IProcessMemoryUsageReport {
	var _result *IProcessMemoryUsageReport
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetReport, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C2C77CBA-1951-4685-8532-7E749ECF8EEB
var IID_IProcessMemoryUsageReport = syscall.GUID{0xC2C77CBA, 0x1951, 0x4685,
	[8]byte{0x85, 0x32, 0x7E, 0x74, 0x9E, 0xCF, 0x8E, 0xEB}}

type IProcessMemoryUsageReportInterface interface {
	win32.IInspectableInterface
	Get_NonPagedPoolSizeInBytes() uint64
	Get_PageFaultCount() uint32
	Get_PageFileSizeInBytes() uint64
	Get_PagedPoolSizeInBytes() uint64
	Get_PeakNonPagedPoolSizeInBytes() uint64
	Get_PeakPageFileSizeInBytes() uint64
	Get_PeakPagedPoolSizeInBytes() uint64
	Get_PeakVirtualMemorySizeInBytes() uint64
	Get_PeakWorkingSetSizeInBytes() uint64
	Get_PrivatePageCount() uint64
	Get_VirtualMemorySizeInBytes() uint64
	Get_WorkingSetSizeInBytes() uint64
}

type IProcessMemoryUsageReportVtbl struct {
	win32.IInspectableVtbl
	Get_NonPagedPoolSizeInBytes      uintptr
	Get_PageFaultCount               uintptr
	Get_PageFileSizeInBytes          uintptr
	Get_PagedPoolSizeInBytes         uintptr
	Get_PeakNonPagedPoolSizeInBytes  uintptr
	Get_PeakPageFileSizeInBytes      uintptr
	Get_PeakPagedPoolSizeInBytes     uintptr
	Get_PeakVirtualMemorySizeInBytes uintptr
	Get_PeakWorkingSetSizeInBytes    uintptr
	Get_PrivatePageCount             uintptr
	Get_VirtualMemorySizeInBytes     uintptr
	Get_WorkingSetSizeInBytes        uintptr
}

type IProcessMemoryUsageReport struct {
	win32.IInspectable
}

func (this *IProcessMemoryUsageReport) Vtbl() *IProcessMemoryUsageReportVtbl {
	return (*IProcessMemoryUsageReportVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProcessMemoryUsageReport) Get_NonPagedPoolSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NonPagedPoolSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProcessMemoryUsageReport) Get_PageFaultCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PageFaultCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProcessMemoryUsageReport) Get_PageFileSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PageFileSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProcessMemoryUsageReport) Get_PagedPoolSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PagedPoolSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProcessMemoryUsageReport) Get_PeakNonPagedPoolSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PeakNonPagedPoolSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProcessMemoryUsageReport) Get_PeakPageFileSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PeakPageFileSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProcessMemoryUsageReport) Get_PeakPagedPoolSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PeakPagedPoolSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProcessMemoryUsageReport) Get_PeakVirtualMemorySizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PeakVirtualMemorySizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProcessMemoryUsageReport) Get_PeakWorkingSetSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PeakWorkingSetSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProcessMemoryUsageReport) Get_PrivatePageCount() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PrivatePageCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProcessMemoryUsageReport) Get_VirtualMemorySizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VirtualMemorySizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProcessMemoryUsageReport) Get_WorkingSetSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WorkingSetSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 6037B3AC-02D6-4234-8362-7FE3ADC81F5F
var IID_ISystemCpuUsage = syscall.GUID{0x6037B3AC, 0x02D6, 0x4234,
	[8]byte{0x83, 0x62, 0x7F, 0xE3, 0xAD, 0xC8, 0x1F, 0x5F}}

type ISystemCpuUsageInterface interface {
	win32.IInspectableInterface
	GetReport() *ISystemCpuUsageReport
}

type ISystemCpuUsageVtbl struct {
	win32.IInspectableVtbl
	GetReport uintptr
}

type ISystemCpuUsage struct {
	win32.IInspectable
}

func (this *ISystemCpuUsage) Vtbl() *ISystemCpuUsageVtbl {
	return (*ISystemCpuUsageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemCpuUsage) GetReport() *ISystemCpuUsageReport {
	var _result *ISystemCpuUsageReport
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetReport, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2C26D0B2-9483-4F62-AB57-82B29D9719B8
var IID_ISystemCpuUsageReport = syscall.GUID{0x2C26D0B2, 0x9483, 0x4F62,
	[8]byte{0xAB, 0x57, 0x82, 0xB2, 0x9D, 0x97, 0x19, 0xB8}}

type ISystemCpuUsageReportInterface interface {
	win32.IInspectableInterface
	Get_KernelTime() TimeSpan
	Get_UserTime() TimeSpan
	Get_IdleTime() TimeSpan
}

type ISystemCpuUsageReportVtbl struct {
	win32.IInspectableVtbl
	Get_KernelTime uintptr
	Get_UserTime   uintptr
	Get_IdleTime   uintptr
}

type ISystemCpuUsageReport struct {
	win32.IInspectable
}

func (this *ISystemCpuUsageReport) Vtbl() *ISystemCpuUsageReportVtbl {
	return (*ISystemCpuUsageReportVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemCpuUsageReport) Get_KernelTime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KernelTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemCpuUsageReport) Get_UserTime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UserTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemCpuUsageReport) Get_IdleTime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IdleTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A290FE05-DFF3-407F-9A1B-0B2B317CA800
var IID_ISystemDiagnosticInfo = syscall.GUID{0xA290FE05, 0xDFF3, 0x407F,
	[8]byte{0x9A, 0x1B, 0x0B, 0x2B, 0x31, 0x7C, 0xA8, 0x00}}

type ISystemDiagnosticInfoInterface interface {
	win32.IInspectableInterface
	Get_MemoryUsage() *ISystemMemoryUsage
	Get_CpuUsage() *ISystemCpuUsage
}

type ISystemDiagnosticInfoVtbl struct {
	win32.IInspectableVtbl
	Get_MemoryUsage uintptr
	Get_CpuUsage    uintptr
}

type ISystemDiagnosticInfo struct {
	win32.IInspectable
}

func (this *ISystemDiagnosticInfo) Vtbl() *ISystemDiagnosticInfoVtbl {
	return (*ISystemDiagnosticInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemDiagnosticInfo) Get_MemoryUsage() *ISystemMemoryUsage {
	var _result *ISystemMemoryUsage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MemoryUsage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISystemDiagnosticInfo) Get_CpuUsage() *ISystemCpuUsage {
	var _result *ISystemCpuUsage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CpuUsage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D404AC21-FC7D-45F0-9A3F-39203AED9F7E
var IID_ISystemDiagnosticInfoStatics = syscall.GUID{0xD404AC21, 0xFC7D, 0x45F0,
	[8]byte{0x9A, 0x3F, 0x39, 0x20, 0x3A, 0xED, 0x9F, 0x7E}}

type ISystemDiagnosticInfoStaticsInterface interface {
	win32.IInspectableInterface
	GetForCurrentSystem() *ISystemDiagnosticInfo
}

type ISystemDiagnosticInfoStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForCurrentSystem uintptr
}

type ISystemDiagnosticInfoStatics struct {
	win32.IInspectable
}

func (this *ISystemDiagnosticInfoStatics) Vtbl() *ISystemDiagnosticInfoStaticsVtbl {
	return (*ISystemDiagnosticInfoStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemDiagnosticInfoStatics) GetForCurrentSystem() *ISystemDiagnosticInfo {
	var _result *ISystemDiagnosticInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentSystem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 79DED189-6AF9-4DA9-A422-15F73255B3EB
var IID_ISystemDiagnosticInfoStatics2 = syscall.GUID{0x79DED189, 0x6AF9, 0x4DA9,
	[8]byte{0xA4, 0x22, 0x15, 0xF7, 0x32, 0x55, 0xB3, 0xEB}}

type ISystemDiagnosticInfoStatics2Interface interface {
	win32.IInspectableInterface
	IsArchitectureSupported(type_ ProcessorArchitecture) bool
	Get_PreferredArchitecture() ProcessorArchitecture
}

type ISystemDiagnosticInfoStatics2Vtbl struct {
	win32.IInspectableVtbl
	IsArchitectureSupported   uintptr
	Get_PreferredArchitecture uintptr
}

type ISystemDiagnosticInfoStatics2 struct {
	win32.IInspectable
}

func (this *ISystemDiagnosticInfoStatics2) Vtbl() *ISystemDiagnosticInfoStatics2Vtbl {
	return (*ISystemDiagnosticInfoStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemDiagnosticInfoStatics2) IsArchitectureSupported(type_ ProcessorArchitecture) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsArchitectureSupported, uintptr(unsafe.Pointer(this)), uintptr(type_), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemDiagnosticInfoStatics2) Get_PreferredArchitecture() ProcessorArchitecture {
	var _result ProcessorArchitecture
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PreferredArchitecture, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 17FFC595-1702-49CF-AA27-2F0A32591404
var IID_ISystemMemoryUsage = syscall.GUID{0x17FFC595, 0x1702, 0x49CF,
	[8]byte{0xAA, 0x27, 0x2F, 0x0A, 0x32, 0x59, 0x14, 0x04}}

type ISystemMemoryUsageInterface interface {
	win32.IInspectableInterface
	GetReport() *ISystemMemoryUsageReport
}

type ISystemMemoryUsageVtbl struct {
	win32.IInspectableVtbl
	GetReport uintptr
}

type ISystemMemoryUsage struct {
	win32.IInspectable
}

func (this *ISystemMemoryUsage) Vtbl() *ISystemMemoryUsageVtbl {
	return (*ISystemMemoryUsageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemMemoryUsage) GetReport() *ISystemMemoryUsageReport {
	var _result *ISystemMemoryUsageReport
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetReport, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 38663C87-2A9F-403A-BD19-2CF3E8169500
var IID_ISystemMemoryUsageReport = syscall.GUID{0x38663C87, 0x2A9F, 0x403A,
	[8]byte{0xBD, 0x19, 0x2C, 0xF3, 0xE8, 0x16, 0x95, 0x00}}

type ISystemMemoryUsageReportInterface interface {
	win32.IInspectableInterface
	Get_TotalPhysicalSizeInBytes() uint64
	Get_AvailableSizeInBytes() uint64
	Get_CommittedSizeInBytes() uint64
}

type ISystemMemoryUsageReportVtbl struct {
	win32.IInspectableVtbl
	Get_TotalPhysicalSizeInBytes uintptr
	Get_AvailableSizeInBytes     uintptr
	Get_CommittedSizeInBytes     uintptr
}

type ISystemMemoryUsageReport struct {
	win32.IInspectable
}

func (this *ISystemMemoryUsageReport) Vtbl() *ISystemMemoryUsageReportVtbl {
	return (*ISystemMemoryUsageReportVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemMemoryUsageReport) Get_TotalPhysicalSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TotalPhysicalSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMemoryUsageReport) Get_AvailableSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AvailableSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMemoryUsageReport) Get_CommittedSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CommittedSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type ProcessCpuUsage struct {
	RtClass
	*IProcessCpuUsage
}

type ProcessCpuUsageReport struct {
	RtClass
	*IProcessCpuUsageReport
}

type ProcessDiagnosticInfo struct {
	RtClass
	*IProcessDiagnosticInfo
}

func NewIProcessDiagnosticInfoStatics2() *IProcessDiagnosticInfoStatics2 {
	var p *IProcessDiagnosticInfoStatics2
	hs := NewHStr("Windows.System.Diagnostics.ProcessDiagnosticInfo")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IProcessDiagnosticInfoStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIProcessDiagnosticInfoStatics() *IProcessDiagnosticInfoStatics {
	var p *IProcessDiagnosticInfoStatics
	hs := NewHStr("Windows.System.Diagnostics.ProcessDiagnosticInfo")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IProcessDiagnosticInfoStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type ProcessDiskUsage struct {
	RtClass
	*IProcessDiskUsage
}

type ProcessDiskUsageReport struct {
	RtClass
	*IProcessDiskUsageReport
}

type ProcessMemoryUsage struct {
	RtClass
	*IProcessMemoryUsage
}

type ProcessMemoryUsageReport struct {
	RtClass
	*IProcessMemoryUsageReport
}

type SystemCpuUsage struct {
	RtClass
	*ISystemCpuUsage
}

type SystemCpuUsageReport struct {
	RtClass
	*ISystemCpuUsageReport
}

type SystemDiagnosticInfo struct {
	RtClass
	*ISystemDiagnosticInfo
}

func NewISystemDiagnosticInfoStatics2() *ISystemDiagnosticInfoStatics2 {
	var p *ISystemDiagnosticInfoStatics2
	hs := NewHStr("Windows.System.Diagnostics.SystemDiagnosticInfo")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISystemDiagnosticInfoStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewISystemDiagnosticInfoStatics() *ISystemDiagnosticInfoStatics {
	var p *ISystemDiagnosticInfoStatics
	hs := NewHStr("Windows.System.Diagnostics.SystemDiagnosticInfo")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISystemDiagnosticInfoStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type SystemMemoryUsage struct {
	RtClass
	*ISystemMemoryUsage
}

type SystemMemoryUsageReport struct {
	RtClass
	*ISystemMemoryUsageReport
}
