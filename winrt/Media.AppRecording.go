package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type AppRecordingSaveScreenshotOption int32

const (
	AppRecordingSaveScreenshotOption_None              AppRecordingSaveScreenshotOption = 0
	AppRecordingSaveScreenshotOption_HdrContentVisible AppRecordingSaveScreenshotOption = 1
)

// structs

type AppRecordingContract struct {
}

// interfaces

// E7E26076-A044-48E2-A512-3094D574C7CC
var IID_IAppRecordingManager = syscall.GUID{0xE7E26076, 0xA044, 0x48E2,
	[8]byte{0xA5, 0x12, 0x30, 0x94, 0xD5, 0x74, 0xC7, 0xCC}}

type IAppRecordingManagerInterface interface {
	win32.IInspectableInterface
	GetStatus() *IAppRecordingStatus
	StartRecordingToFileAsync(file *IStorageFile) *IAsyncOperation[*IAppRecordingResult]
	RecordTimeSpanToFileAsync(startTime DateTime, duration TimeSpan, file *IStorageFile) *IAsyncOperation[*IAppRecordingResult]
	Get_SupportedScreenshotMediaEncodingSubtypes() *IVectorView[string]
	SaveScreenshotToFilesAsync(folder *IStorageFolder, filenamePrefix string, option AppRecordingSaveScreenshotOption, requestedFormats *IIterable[string]) *IAsyncOperation[*IAppRecordingSaveScreenshotResult]
}

type IAppRecordingManagerVtbl struct {
	win32.IInspectableVtbl
	GetStatus                                    uintptr
	StartRecordingToFileAsync                    uintptr
	RecordTimeSpanToFileAsync                    uintptr
	Get_SupportedScreenshotMediaEncodingSubtypes uintptr
	SaveScreenshotToFilesAsync                   uintptr
}

type IAppRecordingManager struct {
	win32.IInspectable
}

func (this *IAppRecordingManager) Vtbl() *IAppRecordingManagerVtbl {
	return (*IAppRecordingManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppRecordingManager) GetStatus() *IAppRecordingStatus {
	var _result *IAppRecordingStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppRecordingManager) StartRecordingToFileAsync(file *IStorageFile) *IAsyncOperation[*IAppRecordingResult] {
	var _result *IAsyncOperation[*IAppRecordingResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartRecordingToFileAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppRecordingManager) RecordTimeSpanToFileAsync(startTime DateTime, duration TimeSpan, file *IStorageFile) *IAsyncOperation[*IAppRecordingResult] {
	var _result *IAsyncOperation[*IAppRecordingResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RecordTimeSpanToFileAsync, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&startTime)), *(*uintptr)(unsafe.Pointer(&duration)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppRecordingManager) Get_SupportedScreenshotMediaEncodingSubtypes() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedScreenshotMediaEncodingSubtypes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppRecordingManager) SaveScreenshotToFilesAsync(folder *IStorageFolder, filenamePrefix string, option AppRecordingSaveScreenshotOption, requestedFormats *IIterable[string]) *IAsyncOperation[*IAppRecordingSaveScreenshotResult] {
	var _result *IAsyncOperation[*IAppRecordingSaveScreenshotResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SaveScreenshotToFilesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(folder)), NewHStr(filenamePrefix).Ptr, uintptr(option), uintptr(unsafe.Pointer(requestedFormats)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 50E709F7-38CE-4BD3-9DB2-E72BBE9DE11D
var IID_IAppRecordingManagerStatics = syscall.GUID{0x50E709F7, 0x38CE, 0x4BD3,
	[8]byte{0x9D, 0xB2, 0xE7, 0x2B, 0xBE, 0x9D, 0xE1, 0x1D}}

type IAppRecordingManagerStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *IAppRecordingManager
}

type IAppRecordingManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
}

type IAppRecordingManagerStatics struct {
	win32.IInspectable
}

func (this *IAppRecordingManagerStatics) Vtbl() *IAppRecordingManagerStaticsVtbl {
	return (*IAppRecordingManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppRecordingManagerStatics) GetDefault() *IAppRecordingManager {
	var _result *IAppRecordingManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3A900864-C66D-46F9-B2D9-5BC2DAD070D7
var IID_IAppRecordingResult = syscall.GUID{0x3A900864, 0xC66D, 0x46F9,
	[8]byte{0xB2, 0xD9, 0x5B, 0xC2, 0xDA, 0xD0, 0x70, 0xD7}}

type IAppRecordingResultInterface interface {
	win32.IInspectableInterface
	Get_Succeeded() bool
	Get_ExtendedError() HResult
	Get_Duration() TimeSpan
	Get_IsFileTruncated() bool
}

type IAppRecordingResultVtbl struct {
	win32.IInspectableVtbl
	Get_Succeeded       uintptr
	Get_ExtendedError   uintptr
	Get_Duration        uintptr
	Get_IsFileTruncated uintptr
}

type IAppRecordingResult struct {
	win32.IInspectable
}

func (this *IAppRecordingResult) Vtbl() *IAppRecordingResultVtbl {
	return (*IAppRecordingResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppRecordingResult) Get_Succeeded() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Succeeded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppRecordingResult) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppRecordingResult) Get_Duration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Duration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppRecordingResult) Get_IsFileTruncated() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsFileTruncated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 9C5B8D0A-0ABB-4457-AAEE-24F9C12EC778
var IID_IAppRecordingSaveScreenshotResult = syscall.GUID{0x9C5B8D0A, 0x0ABB, 0x4457,
	[8]byte{0xAA, 0xEE, 0x24, 0xF9, 0xC1, 0x2E, 0xC7, 0x78}}

type IAppRecordingSaveScreenshotResultInterface interface {
	win32.IInspectableInterface
	Get_Succeeded() bool
	Get_ExtendedError() HResult
	Get_SavedScreenshotInfos() *IVectorView[*IAppRecordingSavedScreenshotInfo]
}

type IAppRecordingSaveScreenshotResultVtbl struct {
	win32.IInspectableVtbl
	Get_Succeeded            uintptr
	Get_ExtendedError        uintptr
	Get_SavedScreenshotInfos uintptr
}

type IAppRecordingSaveScreenshotResult struct {
	win32.IInspectable
}

func (this *IAppRecordingSaveScreenshotResult) Vtbl() *IAppRecordingSaveScreenshotResultVtbl {
	return (*IAppRecordingSaveScreenshotResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppRecordingSaveScreenshotResult) Get_Succeeded() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Succeeded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppRecordingSaveScreenshotResult) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppRecordingSaveScreenshotResult) Get_SavedScreenshotInfos() *IVectorView[*IAppRecordingSavedScreenshotInfo] {
	var _result *IVectorView[*IAppRecordingSavedScreenshotInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SavedScreenshotInfos, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9B642D0A-189A-4D00-BF25-E1BB1249D594
var IID_IAppRecordingSavedScreenshotInfo = syscall.GUID{0x9B642D0A, 0x189A, 0x4D00,
	[8]byte{0xBF, 0x25, 0xE1, 0xBB, 0x12, 0x49, 0xD5, 0x94}}

type IAppRecordingSavedScreenshotInfoInterface interface {
	win32.IInspectableInterface
	Get_File() *IStorageFile
	Get_MediaEncodingSubtype() string
}

type IAppRecordingSavedScreenshotInfoVtbl struct {
	win32.IInspectableVtbl
	Get_File                 uintptr
	Get_MediaEncodingSubtype uintptr
}

type IAppRecordingSavedScreenshotInfo struct {
	win32.IInspectable
}

func (this *IAppRecordingSavedScreenshotInfo) Vtbl() *IAppRecordingSavedScreenshotInfoVtbl {
	return (*IAppRecordingSavedScreenshotInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppRecordingSavedScreenshotInfo) Get_File() *IStorageFile {
	var _result *IStorageFile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_File, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppRecordingSavedScreenshotInfo) Get_MediaEncodingSubtype() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaEncodingSubtype, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 1D0CC82C-BC18-4B8A-A6EF-127EFAB3B5D9
var IID_IAppRecordingStatus = syscall.GUID{0x1D0CC82C, 0xBC18, 0x4B8A,
	[8]byte{0xA6, 0xEF, 0x12, 0x7E, 0xFA, 0xB3, 0xB5, 0xD9}}

type IAppRecordingStatusInterface interface {
	win32.IInspectableInterface
	Get_CanRecord() bool
	Get_CanRecordTimeSpan() bool
	Get_HistoricalBufferDuration() TimeSpan
	Get_Details() *IAppRecordingStatusDetails
}

type IAppRecordingStatusVtbl struct {
	win32.IInspectableVtbl
	Get_CanRecord                uintptr
	Get_CanRecordTimeSpan        uintptr
	Get_HistoricalBufferDuration uintptr
	Get_Details                  uintptr
}

type IAppRecordingStatus struct {
	win32.IInspectable
}

func (this *IAppRecordingStatus) Vtbl() *IAppRecordingStatusVtbl {
	return (*IAppRecordingStatusVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppRecordingStatus) Get_CanRecord() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanRecord, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppRecordingStatus) Get_CanRecordTimeSpan() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanRecordTimeSpan, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppRecordingStatus) Get_HistoricalBufferDuration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HistoricalBufferDuration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppRecordingStatus) Get_Details() *IAppRecordingStatusDetails {
	var _result *IAppRecordingStatusDetails
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Details, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B538A9B0-14ED-4412-AC45-6D672C9C9949
var IID_IAppRecordingStatusDetails = syscall.GUID{0xB538A9B0, 0x14ED, 0x4412,
	[8]byte{0xAC, 0x45, 0x6D, 0x67, 0x2C, 0x9C, 0x99, 0x49}}

type IAppRecordingStatusDetailsInterface interface {
	win32.IInspectableInterface
	Get_IsAnyAppBroadcasting() bool
	Get_IsCaptureResourceUnavailable() bool
	Get_IsGameStreamInProgress() bool
	Get_IsTimeSpanRecordingDisabled() bool
	Get_IsGpuConstrained() bool
	Get_IsAppInactive() bool
	Get_IsBlockedForApp() bool
	Get_IsDisabledByUser() bool
	Get_IsDisabledBySystem() bool
}

type IAppRecordingStatusDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_IsAnyAppBroadcasting         uintptr
	Get_IsCaptureResourceUnavailable uintptr
	Get_IsGameStreamInProgress       uintptr
	Get_IsTimeSpanRecordingDisabled  uintptr
	Get_IsGpuConstrained             uintptr
	Get_IsAppInactive                uintptr
	Get_IsBlockedForApp              uintptr
	Get_IsDisabledByUser             uintptr
	Get_IsDisabledBySystem           uintptr
}

type IAppRecordingStatusDetails struct {
	win32.IInspectable
}

func (this *IAppRecordingStatusDetails) Vtbl() *IAppRecordingStatusDetailsVtbl {
	return (*IAppRecordingStatusDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppRecordingStatusDetails) Get_IsAnyAppBroadcasting() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsAnyAppBroadcasting, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppRecordingStatusDetails) Get_IsCaptureResourceUnavailable() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCaptureResourceUnavailable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppRecordingStatusDetails) Get_IsGameStreamInProgress() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsGameStreamInProgress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppRecordingStatusDetails) Get_IsTimeSpanRecordingDisabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsTimeSpanRecordingDisabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppRecordingStatusDetails) Get_IsGpuConstrained() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsGpuConstrained, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppRecordingStatusDetails) Get_IsAppInactive() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsAppInactive, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppRecordingStatusDetails) Get_IsBlockedForApp() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsBlockedForApp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppRecordingStatusDetails) Get_IsDisabledByUser() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDisabledByUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppRecordingStatusDetails) Get_IsDisabledBySystem() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDisabledBySystem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type AppRecordingManager struct {
	RtClass
	*IAppRecordingManager
}

func NewIAppRecordingManagerStatics() *IAppRecordingManagerStatics {
	var p *IAppRecordingManagerStatics
	hs := NewHStr("Windows.Media.AppRecording.AppRecordingManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAppRecordingManagerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type AppRecordingResult struct {
	RtClass
	*IAppRecordingResult
}

type AppRecordingSaveScreenshotResult struct {
	RtClass
	*IAppRecordingSaveScreenshotResult
}

type AppRecordingSavedScreenshotInfo struct {
	RtClass
	*IAppRecordingSavedScreenshotInfo
}

type AppRecordingStatus struct {
	RtClass
	*IAppRecordingStatus
}

type AppRecordingStatusDetails struct {
	RtClass
	*IAppRecordingStatusDetails
}
