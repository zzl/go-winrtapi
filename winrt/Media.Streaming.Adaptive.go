package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type AdaptiveMediaSourceCreationStatus int32

const (
	AdaptiveMediaSourceCreationStatus_Success                        AdaptiveMediaSourceCreationStatus = 0
	AdaptiveMediaSourceCreationStatus_ManifestDownloadFailure        AdaptiveMediaSourceCreationStatus = 1
	AdaptiveMediaSourceCreationStatus_ManifestParseFailure           AdaptiveMediaSourceCreationStatus = 2
	AdaptiveMediaSourceCreationStatus_UnsupportedManifestContentType AdaptiveMediaSourceCreationStatus = 3
	AdaptiveMediaSourceCreationStatus_UnsupportedManifestVersion     AdaptiveMediaSourceCreationStatus = 4
	AdaptiveMediaSourceCreationStatus_UnsupportedManifestProfile     AdaptiveMediaSourceCreationStatus = 5
	AdaptiveMediaSourceCreationStatus_UnknownFailure                 AdaptiveMediaSourceCreationStatus = 6
)

// enum
type AdaptiveMediaSourceDiagnosticType int32

const (
	AdaptiveMediaSourceDiagnosticType_ManifestUnchangedUponReload              AdaptiveMediaSourceDiagnosticType = 0
	AdaptiveMediaSourceDiagnosticType_ManifestMismatchUponReload               AdaptiveMediaSourceDiagnosticType = 1
	AdaptiveMediaSourceDiagnosticType_ManifestSignaledEndOfLiveEventUponReload AdaptiveMediaSourceDiagnosticType = 2
	AdaptiveMediaSourceDiagnosticType_MediaSegmentSkipped                      AdaptiveMediaSourceDiagnosticType = 3
	AdaptiveMediaSourceDiagnosticType_ResourceNotFound                         AdaptiveMediaSourceDiagnosticType = 4
	AdaptiveMediaSourceDiagnosticType_ResourceTimedOut                         AdaptiveMediaSourceDiagnosticType = 5
	AdaptiveMediaSourceDiagnosticType_ResourceParsingError                     AdaptiveMediaSourceDiagnosticType = 6
	AdaptiveMediaSourceDiagnosticType_BitrateDisabled                          AdaptiveMediaSourceDiagnosticType = 7
	AdaptiveMediaSourceDiagnosticType_FatalMediaSourceError                    AdaptiveMediaSourceDiagnosticType = 8
)

// enum
type AdaptiveMediaSourceDownloadBitrateChangedReason int32

const (
	AdaptiveMediaSourceDownloadBitrateChangedReason_SufficientInboundBitsPerSecond   AdaptiveMediaSourceDownloadBitrateChangedReason = 0
	AdaptiveMediaSourceDownloadBitrateChangedReason_InsufficientInboundBitsPerSecond AdaptiveMediaSourceDownloadBitrateChangedReason = 1
	AdaptiveMediaSourceDownloadBitrateChangedReason_LowBufferLevel                   AdaptiveMediaSourceDownloadBitrateChangedReason = 2
	AdaptiveMediaSourceDownloadBitrateChangedReason_PositionChanged                  AdaptiveMediaSourceDownloadBitrateChangedReason = 3
	AdaptiveMediaSourceDownloadBitrateChangedReason_TrackSelectionChanged            AdaptiveMediaSourceDownloadBitrateChangedReason = 4
	AdaptiveMediaSourceDownloadBitrateChangedReason_DesiredBitratesChanged           AdaptiveMediaSourceDownloadBitrateChangedReason = 5
	AdaptiveMediaSourceDownloadBitrateChangedReason_ErrorInPreviousBitrate           AdaptiveMediaSourceDownloadBitrateChangedReason = 6
)

// enum
type AdaptiveMediaSourceResourceType int32

const (
	AdaptiveMediaSourceResourceType_Manifest              AdaptiveMediaSourceResourceType = 0
	AdaptiveMediaSourceResourceType_InitializationSegment AdaptiveMediaSourceResourceType = 1
	AdaptiveMediaSourceResourceType_MediaSegment          AdaptiveMediaSourceResourceType = 2
	AdaptiveMediaSourceResourceType_Key                   AdaptiveMediaSourceResourceType = 3
	AdaptiveMediaSourceResourceType_InitializationVector  AdaptiveMediaSourceResourceType = 4
	AdaptiveMediaSourceResourceType_MediaSegmentIndex     AdaptiveMediaSourceResourceType = 5
)

// interfaces

// 4C7332EF-D39F-4396-B4D9-043957A7C964
var IID_IAdaptiveMediaSource = syscall.GUID{0x4C7332EF, 0xD39F, 0x4396,
	[8]byte{0xB4, 0xD9, 0x04, 0x39, 0x57, 0xA7, 0xC9, 0x64}}

type IAdaptiveMediaSourceInterface interface {
	win32.IInspectableInterface
	Get_IsLive() bool
	Get_DesiredLiveOffset() TimeSpan
	Put_DesiredLiveOffset(value TimeSpan)
	Get_InitialBitrate() uint32
	Put_InitialBitrate(value uint32)
	Get_CurrentDownloadBitrate() uint32
	Get_CurrentPlaybackBitrate() uint32
	Get_AvailableBitrates() *IVectorView[uint32]
	Get_DesiredMinBitrate() *IReference[uint32]
	Put_DesiredMinBitrate(value *IReference[uint32])
	Get_DesiredMaxBitrate() *IReference[uint32]
	Put_DesiredMaxBitrate(value *IReference[uint32])
	Get_AudioOnlyPlayback() bool
	Get_InboundBitsPerSecond() uint64
	Get_InboundBitsPerSecondWindow() TimeSpan
	Put_InboundBitsPerSecondWindow(value TimeSpan)
	Add_DownloadBitrateChanged(handler TypedEventHandler[*IAdaptiveMediaSource, *IAdaptiveMediaSourceDownloadBitrateChangedEventArgs]) EventRegistrationToken
	Remove_DownloadBitrateChanged(token EventRegistrationToken)
	Add_PlaybackBitrateChanged(handler TypedEventHandler[*IAdaptiveMediaSource, *IAdaptiveMediaSourcePlaybackBitrateChangedEventArgs]) EventRegistrationToken
	Remove_PlaybackBitrateChanged(token EventRegistrationToken)
	Add_DownloadRequested(handler TypedEventHandler[*IAdaptiveMediaSource, *IAdaptiveMediaSourceDownloadRequestedEventArgs]) EventRegistrationToken
	Remove_DownloadRequested(token EventRegistrationToken)
	Add_DownloadCompleted(handler TypedEventHandler[*IAdaptiveMediaSource, *IAdaptiveMediaSourceDownloadCompletedEventArgs]) EventRegistrationToken
	Remove_DownloadCompleted(token EventRegistrationToken)
	Add_DownloadFailed(handler TypedEventHandler[*IAdaptiveMediaSource, *IAdaptiveMediaSourceDownloadFailedEventArgs]) EventRegistrationToken
	Remove_DownloadFailed(token EventRegistrationToken)
}

type IAdaptiveMediaSourceVtbl struct {
	win32.IInspectableVtbl
	Get_IsLive                     uintptr
	Get_DesiredLiveOffset          uintptr
	Put_DesiredLiveOffset          uintptr
	Get_InitialBitrate             uintptr
	Put_InitialBitrate             uintptr
	Get_CurrentDownloadBitrate     uintptr
	Get_CurrentPlaybackBitrate     uintptr
	Get_AvailableBitrates          uintptr
	Get_DesiredMinBitrate          uintptr
	Put_DesiredMinBitrate          uintptr
	Get_DesiredMaxBitrate          uintptr
	Put_DesiredMaxBitrate          uintptr
	Get_AudioOnlyPlayback          uintptr
	Get_InboundBitsPerSecond       uintptr
	Get_InboundBitsPerSecondWindow uintptr
	Put_InboundBitsPerSecondWindow uintptr
	Add_DownloadBitrateChanged     uintptr
	Remove_DownloadBitrateChanged  uintptr
	Add_PlaybackBitrateChanged     uintptr
	Remove_PlaybackBitrateChanged  uintptr
	Add_DownloadRequested          uintptr
	Remove_DownloadRequested       uintptr
	Add_DownloadCompleted          uintptr
	Remove_DownloadCompleted       uintptr
	Add_DownloadFailed             uintptr
	Remove_DownloadFailed          uintptr
}

type IAdaptiveMediaSource struct {
	win32.IInspectable
}

func (this *IAdaptiveMediaSource) Vtbl() *IAdaptiveMediaSourceVtbl {
	return (*IAdaptiveMediaSourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveMediaSource) Get_IsLive() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsLive, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSource) Get_DesiredLiveOffset() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesiredLiveOffset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSource) Put_DesiredLiveOffset(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DesiredLiveOffset, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IAdaptiveMediaSource) Get_InitialBitrate() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InitialBitrate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSource) Put_InitialBitrate(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InitialBitrate, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAdaptiveMediaSource) Get_CurrentDownloadBitrate() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentDownloadBitrate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSource) Get_CurrentPlaybackBitrate() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentPlaybackBitrate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSource) Get_AvailableBitrates() *IVectorView[uint32] {
	var _result *IVectorView[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AvailableBitrates, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSource) Get_DesiredMinBitrate() *IReference[uint32] {
	var _result *IReference[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesiredMinBitrate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSource) Put_DesiredMinBitrate(value *IReference[uint32]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DesiredMinBitrate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAdaptiveMediaSource) Get_DesiredMaxBitrate() *IReference[uint32] {
	var _result *IReference[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesiredMaxBitrate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSource) Put_DesiredMaxBitrate(value *IReference[uint32]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DesiredMaxBitrate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAdaptiveMediaSource) Get_AudioOnlyPlayback() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioOnlyPlayback, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSource) Get_InboundBitsPerSecond() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InboundBitsPerSecond, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSource) Get_InboundBitsPerSecondWindow() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InboundBitsPerSecondWindow, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSource) Put_InboundBitsPerSecondWindow(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InboundBitsPerSecondWindow, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IAdaptiveMediaSource) Add_DownloadBitrateChanged(handler TypedEventHandler[*IAdaptiveMediaSource, *IAdaptiveMediaSourceDownloadBitrateChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DownloadBitrateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSource) Remove_DownloadBitrateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DownloadBitrateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAdaptiveMediaSource) Add_PlaybackBitrateChanged(handler TypedEventHandler[*IAdaptiveMediaSource, *IAdaptiveMediaSourcePlaybackBitrateChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PlaybackBitrateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSource) Remove_PlaybackBitrateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PlaybackBitrateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAdaptiveMediaSource) Add_DownloadRequested(handler TypedEventHandler[*IAdaptiveMediaSource, *IAdaptiveMediaSourceDownloadRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DownloadRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSource) Remove_DownloadRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DownloadRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAdaptiveMediaSource) Add_DownloadCompleted(handler TypedEventHandler[*IAdaptiveMediaSource, *IAdaptiveMediaSourceDownloadCompletedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DownloadCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSource) Remove_DownloadCompleted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DownloadCompleted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAdaptiveMediaSource) Add_DownloadFailed(handler TypedEventHandler[*IAdaptiveMediaSource, *IAdaptiveMediaSourceDownloadFailedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DownloadFailed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSource) Remove_DownloadFailed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DownloadFailed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 17890342-6760-4BB9-A58A-F7AA98B08C0E
var IID_IAdaptiveMediaSource2 = syscall.GUID{0x17890342, 0x6760, 0x4BB9,
	[8]byte{0xA5, 0x8A, 0xF7, 0xAA, 0x98, 0xB0, 0x8C, 0x0E}}

type IAdaptiveMediaSource2Interface interface {
	win32.IInspectableInterface
	Get_AdvancedSettings() *IAdaptiveMediaSourceAdvancedSettings
}

type IAdaptiveMediaSource2Vtbl struct {
	win32.IInspectableVtbl
	Get_AdvancedSettings uintptr
}

type IAdaptiveMediaSource2 struct {
	win32.IInspectable
}

func (this *IAdaptiveMediaSource2) Vtbl() *IAdaptiveMediaSource2Vtbl {
	return (*IAdaptiveMediaSource2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveMediaSource2) Get_AdvancedSettings() *IAdaptiveMediaSourceAdvancedSettings {
	var _result *IAdaptiveMediaSourceAdvancedSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AdvancedSettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BA7023FD-C334-461B-A36E-C99F54F7174A
var IID_IAdaptiveMediaSource3 = syscall.GUID{0xBA7023FD, 0xC334, 0x461B,
	[8]byte{0xA3, 0x6E, 0xC9, 0x9F, 0x54, 0xF7, 0x17, 0x4A}}

type IAdaptiveMediaSource3Interface interface {
	win32.IInspectableInterface
	Get_MinLiveOffset() *IReference[TimeSpan]
	Get_MaxSeekableWindowSize() *IReference[TimeSpan]
	Get_DesiredSeekableWindowSize() *IReference[TimeSpan]
	Put_DesiredSeekableWindowSize(value *IReference[TimeSpan])
	Get_Diagnostics() *IAdaptiveMediaSourceDiagnostics
	GetCorrelatedTimes() *IAdaptiveMediaSourceCorrelatedTimes
}

type IAdaptiveMediaSource3Vtbl struct {
	win32.IInspectableVtbl
	Get_MinLiveOffset             uintptr
	Get_MaxSeekableWindowSize     uintptr
	Get_DesiredSeekableWindowSize uintptr
	Put_DesiredSeekableWindowSize uintptr
	Get_Diagnostics               uintptr
	GetCorrelatedTimes            uintptr
}

type IAdaptiveMediaSource3 struct {
	win32.IInspectable
}

func (this *IAdaptiveMediaSource3) Vtbl() *IAdaptiveMediaSource3Vtbl {
	return (*IAdaptiveMediaSource3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveMediaSource3) Get_MinLiveOffset() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinLiveOffset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSource3) Get_MaxSeekableWindowSize() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxSeekableWindowSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSource3) Get_DesiredSeekableWindowSize() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesiredSeekableWindowSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSource3) Put_DesiredSeekableWindowSize(value *IReference[TimeSpan]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DesiredSeekableWindowSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAdaptiveMediaSource3) Get_Diagnostics() *IAdaptiveMediaSourceDiagnostics {
	var _result *IAdaptiveMediaSourceDiagnostics
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Diagnostics, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSource3) GetCorrelatedTimes() *IAdaptiveMediaSourceCorrelatedTimes {
	var _result *IAdaptiveMediaSourceCorrelatedTimes
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCorrelatedTimes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 55DB1680-1AEB-47DC-AA08-9A11610BA45A
var IID_IAdaptiveMediaSourceAdvancedSettings = syscall.GUID{0x55DB1680, 0x1AEB, 0x47DC,
	[8]byte{0xAA, 0x08, 0x9A, 0x11, 0x61, 0x0B, 0xA4, 0x5A}}

type IAdaptiveMediaSourceAdvancedSettingsInterface interface {
	win32.IInspectableInterface
	Get_AllSegmentsIndependent() bool
	Put_AllSegmentsIndependent(value bool)
	Get_DesiredBitrateHeadroomRatio() *IReference[float64]
	Put_DesiredBitrateHeadroomRatio(value *IReference[float64])
	Get_BitrateDowngradeTriggerRatio() *IReference[float64]
	Put_BitrateDowngradeTriggerRatio(value *IReference[float64])
}

type IAdaptiveMediaSourceAdvancedSettingsVtbl struct {
	win32.IInspectableVtbl
	Get_AllSegmentsIndependent       uintptr
	Put_AllSegmentsIndependent       uintptr
	Get_DesiredBitrateHeadroomRatio  uintptr
	Put_DesiredBitrateHeadroomRatio  uintptr
	Get_BitrateDowngradeTriggerRatio uintptr
	Put_BitrateDowngradeTriggerRatio uintptr
}

type IAdaptiveMediaSourceAdvancedSettings struct {
	win32.IInspectable
}

func (this *IAdaptiveMediaSourceAdvancedSettings) Vtbl() *IAdaptiveMediaSourceAdvancedSettingsVtbl {
	return (*IAdaptiveMediaSourceAdvancedSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveMediaSourceAdvancedSettings) Get_AllSegmentsIndependent() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllSegmentsIndependent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSourceAdvancedSettings) Put_AllSegmentsIndependent(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AllSegmentsIndependent, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAdaptiveMediaSourceAdvancedSettings) Get_DesiredBitrateHeadroomRatio() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesiredBitrateHeadroomRatio, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceAdvancedSettings) Put_DesiredBitrateHeadroomRatio(value *IReference[float64]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DesiredBitrateHeadroomRatio, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAdaptiveMediaSourceAdvancedSettings) Get_BitrateDowngradeTriggerRatio() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BitrateDowngradeTriggerRatio, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceAdvancedSettings) Put_BitrateDowngradeTriggerRatio(value *IReference[float64]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BitrateDowngradeTriggerRatio, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 05108787-E032-48E1-AB8D-002B0B3051DF
var IID_IAdaptiveMediaSourceCorrelatedTimes = syscall.GUID{0x05108787, 0xE032, 0x48E1,
	[8]byte{0xAB, 0x8D, 0x00, 0x2B, 0x0B, 0x30, 0x51, 0xDF}}

type IAdaptiveMediaSourceCorrelatedTimesInterface interface {
	win32.IInspectableInterface
	Get_Position() *IReference[TimeSpan]
	Get_PresentationTimeStamp() *IReference[TimeSpan]
	Get_ProgramDateTime() *IReference[DateTime]
}

type IAdaptiveMediaSourceCorrelatedTimesVtbl struct {
	win32.IInspectableVtbl
	Get_Position              uintptr
	Get_PresentationTimeStamp uintptr
	Get_ProgramDateTime       uintptr
}

type IAdaptiveMediaSourceCorrelatedTimes struct {
	win32.IInspectable
}

func (this *IAdaptiveMediaSourceCorrelatedTimes) Vtbl() *IAdaptiveMediaSourceCorrelatedTimesVtbl {
	return (*IAdaptiveMediaSourceCorrelatedTimesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveMediaSourceCorrelatedTimes) Get_Position() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceCorrelatedTimes) Get_PresentationTimeStamp() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PresentationTimeStamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceCorrelatedTimes) Get_ProgramDateTime() *IReference[DateTime] {
	var _result *IReference[DateTime]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProgramDateTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4686B6B2-800F-4E31-9093-76D4782013E7
var IID_IAdaptiveMediaSourceCreationResult = syscall.GUID{0x4686B6B2, 0x800F, 0x4E31,
	[8]byte{0x90, 0x93, 0x76, 0xD4, 0x78, 0x20, 0x13, 0xE7}}

type IAdaptiveMediaSourceCreationResultInterface interface {
	win32.IInspectableInterface
	Get_Status() AdaptiveMediaSourceCreationStatus
	Get_MediaSource() *IAdaptiveMediaSource
	Get_HttpResponseMessage() *IHttpResponseMessage
}

type IAdaptiveMediaSourceCreationResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status              uintptr
	Get_MediaSource         uintptr
	Get_HttpResponseMessage uintptr
}

type IAdaptiveMediaSourceCreationResult struct {
	win32.IInspectable
}

func (this *IAdaptiveMediaSourceCreationResult) Vtbl() *IAdaptiveMediaSourceCreationResultVtbl {
	return (*IAdaptiveMediaSourceCreationResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveMediaSourceCreationResult) Get_Status() AdaptiveMediaSourceCreationStatus {
	var _result AdaptiveMediaSourceCreationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSourceCreationResult) Get_MediaSource() *IAdaptiveMediaSource {
	var _result *IAdaptiveMediaSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceCreationResult) Get_HttpResponseMessage() *IHttpResponseMessage {
	var _result *IHttpResponseMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HttpResponseMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1C3243BF-1C44-404B-A201-DF45AC7898E8
var IID_IAdaptiveMediaSourceCreationResult2 = syscall.GUID{0x1C3243BF, 0x1C44, 0x404B,
	[8]byte{0xA2, 0x01, 0xDF, 0x45, 0xAC, 0x78, 0x98, 0xE8}}

type IAdaptiveMediaSourceCreationResult2Interface interface {
	win32.IInspectableInterface
	Get_ExtendedError() HResult
}

type IAdaptiveMediaSourceCreationResult2Vtbl struct {
	win32.IInspectableVtbl
	Get_ExtendedError uintptr
}

type IAdaptiveMediaSourceCreationResult2 struct {
	win32.IInspectable
}

func (this *IAdaptiveMediaSourceCreationResult2) Vtbl() *IAdaptiveMediaSourceCreationResult2Vtbl {
	return (*IAdaptiveMediaSourceCreationResult2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveMediaSourceCreationResult2) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 3AF64F06-6D9C-494A-B7A9-B3A5DEE6AD68
var IID_IAdaptiveMediaSourceDiagnosticAvailableEventArgs = syscall.GUID{0x3AF64F06, 0x6D9C, 0x494A,
	[8]byte{0xB7, 0xA9, 0xB3, 0xA5, 0xDE, 0xE6, 0xAD, 0x68}}

type IAdaptiveMediaSourceDiagnosticAvailableEventArgsInterface interface {
	win32.IInspectableInterface
	Get_DiagnosticType() AdaptiveMediaSourceDiagnosticType
	Get_RequestId() *IReference[int32]
	Get_Position() *IReference[TimeSpan]
	Get_SegmentId() *IReference[uint64]
	Get_ResourceType() *IReference[AdaptiveMediaSourceResourceType]
	Get_ResourceUri() *IUriRuntimeClass
	Get_ResourceByteRangeOffset() *IReference[uint64]
	Get_ResourceByteRangeLength() *IReference[uint64]
	Get_Bitrate() *IReference[uint32]
}

type IAdaptiveMediaSourceDiagnosticAvailableEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_DiagnosticType          uintptr
	Get_RequestId               uintptr
	Get_Position                uintptr
	Get_SegmentId               uintptr
	Get_ResourceType            uintptr
	Get_ResourceUri             uintptr
	Get_ResourceByteRangeOffset uintptr
	Get_ResourceByteRangeLength uintptr
	Get_Bitrate                 uintptr
}

type IAdaptiveMediaSourceDiagnosticAvailableEventArgs struct {
	win32.IInspectable
}

func (this *IAdaptiveMediaSourceDiagnosticAvailableEventArgs) Vtbl() *IAdaptiveMediaSourceDiagnosticAvailableEventArgsVtbl {
	return (*IAdaptiveMediaSourceDiagnosticAvailableEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveMediaSourceDiagnosticAvailableEventArgs) Get_DiagnosticType() AdaptiveMediaSourceDiagnosticType {
	var _result AdaptiveMediaSourceDiagnosticType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DiagnosticType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSourceDiagnosticAvailableEventArgs) Get_RequestId() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDiagnosticAvailableEventArgs) Get_Position() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDiagnosticAvailableEventArgs) Get_SegmentId() *IReference[uint64] {
	var _result *IReference[uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SegmentId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDiagnosticAvailableEventArgs) Get_ResourceType() *IReference[AdaptiveMediaSourceResourceType] {
	var _result *IReference[AdaptiveMediaSourceResourceType]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDiagnosticAvailableEventArgs) Get_ResourceUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDiagnosticAvailableEventArgs) Get_ResourceByteRangeOffset() *IReference[uint64] {
	var _result *IReference[uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceByteRangeOffset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDiagnosticAvailableEventArgs) Get_ResourceByteRangeLength() *IReference[uint64] {
	var _result *IReference[uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceByteRangeLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDiagnosticAvailableEventArgs) Get_Bitrate() *IReference[uint32] {
	var _result *IReference[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bitrate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8C6DD857-16A5-4D9F-810E-00BD901B3EF9
var IID_IAdaptiveMediaSourceDiagnosticAvailableEventArgs2 = syscall.GUID{0x8C6DD857, 0x16A5, 0x4D9F,
	[8]byte{0x81, 0x0E, 0x00, 0xBD, 0x90, 0x1B, 0x3E, 0xF9}}

type IAdaptiveMediaSourceDiagnosticAvailableEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_ExtendedError() HResult
}

type IAdaptiveMediaSourceDiagnosticAvailableEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_ExtendedError uintptr
}

type IAdaptiveMediaSourceDiagnosticAvailableEventArgs2 struct {
	win32.IInspectable
}

func (this *IAdaptiveMediaSourceDiagnosticAvailableEventArgs2) Vtbl() *IAdaptiveMediaSourceDiagnosticAvailableEventArgs2Vtbl {
	return (*IAdaptiveMediaSourceDiagnosticAvailableEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveMediaSourceDiagnosticAvailableEventArgs2) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C3650CD5-DAEB-4103-84DA-68769AD513FF
var IID_IAdaptiveMediaSourceDiagnosticAvailableEventArgs3 = syscall.GUID{0xC3650CD5, 0xDAEB, 0x4103,
	[8]byte{0x84, 0xDA, 0x68, 0x76, 0x9A, 0xD5, 0x13, 0xFF}}

type IAdaptiveMediaSourceDiagnosticAvailableEventArgs3Interface interface {
	win32.IInspectableInterface
	Get_ResourceDuration() *IReference[TimeSpan]
	Get_ResourceContentType() string
}

type IAdaptiveMediaSourceDiagnosticAvailableEventArgs3Vtbl struct {
	win32.IInspectableVtbl
	Get_ResourceDuration    uintptr
	Get_ResourceContentType uintptr
}

type IAdaptiveMediaSourceDiagnosticAvailableEventArgs3 struct {
	win32.IInspectable
}

func (this *IAdaptiveMediaSourceDiagnosticAvailableEventArgs3) Vtbl() *IAdaptiveMediaSourceDiagnosticAvailableEventArgs3Vtbl {
	return (*IAdaptiveMediaSourceDiagnosticAvailableEventArgs3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveMediaSourceDiagnosticAvailableEventArgs3) Get_ResourceDuration() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceDuration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDiagnosticAvailableEventArgs3) Get_ResourceContentType() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceContentType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 9B24EE68-962E-448C-AEBF-B29B56098E23
var IID_IAdaptiveMediaSourceDiagnostics = syscall.GUID{0x9B24EE68, 0x962E, 0x448C,
	[8]byte{0xAE, 0xBF, 0xB2, 0x9B, 0x56, 0x09, 0x8E, 0x23}}

type IAdaptiveMediaSourceDiagnosticsInterface interface {
	win32.IInspectableInterface
	Add_DiagnosticAvailable(handler TypedEventHandler[*IAdaptiveMediaSourceDiagnostics, *IAdaptiveMediaSourceDiagnosticAvailableEventArgs]) EventRegistrationToken
	Remove_DiagnosticAvailable(token EventRegistrationToken)
}

type IAdaptiveMediaSourceDiagnosticsVtbl struct {
	win32.IInspectableVtbl
	Add_DiagnosticAvailable    uintptr
	Remove_DiagnosticAvailable uintptr
}

type IAdaptiveMediaSourceDiagnostics struct {
	win32.IInspectable
}

func (this *IAdaptiveMediaSourceDiagnostics) Vtbl() *IAdaptiveMediaSourceDiagnosticsVtbl {
	return (*IAdaptiveMediaSourceDiagnosticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveMediaSourceDiagnostics) Add_DiagnosticAvailable(handler TypedEventHandler[*IAdaptiveMediaSourceDiagnostics, *IAdaptiveMediaSourceDiagnosticAvailableEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DiagnosticAvailable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSourceDiagnostics) Remove_DiagnosticAvailable(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DiagnosticAvailable, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 670C0A44-E04E-4EFF-816A-17399F78F4BA
var IID_IAdaptiveMediaSourceDownloadBitrateChangedEventArgs = syscall.GUID{0x670C0A44, 0xE04E, 0x4EFF,
	[8]byte{0x81, 0x6A, 0x17, 0x39, 0x9F, 0x78, 0xF4, 0xBA}}

type IAdaptiveMediaSourceDownloadBitrateChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_OldValue() uint32
	Get_NewValue() uint32
}

type IAdaptiveMediaSourceDownloadBitrateChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_OldValue uintptr
	Get_NewValue uintptr
}

type IAdaptiveMediaSourceDownloadBitrateChangedEventArgs struct {
	win32.IInspectable
}

func (this *IAdaptiveMediaSourceDownloadBitrateChangedEventArgs) Vtbl() *IAdaptiveMediaSourceDownloadBitrateChangedEventArgsVtbl {
	return (*IAdaptiveMediaSourceDownloadBitrateChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveMediaSourceDownloadBitrateChangedEventArgs) Get_OldValue() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OldValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSourceDownloadBitrateChangedEventArgs) Get_NewValue() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NewValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F3F1F444-96AE-4DE0-B540-2B3246E6968C
var IID_IAdaptiveMediaSourceDownloadBitrateChangedEventArgs2 = syscall.GUID{0xF3F1F444, 0x96AE, 0x4DE0,
	[8]byte{0xB5, 0x40, 0x2B, 0x32, 0x46, 0xE6, 0x96, 0x8C}}

type IAdaptiveMediaSourceDownloadBitrateChangedEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_Reason() AdaptiveMediaSourceDownloadBitrateChangedReason
}

type IAdaptiveMediaSourceDownloadBitrateChangedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_Reason uintptr
}

type IAdaptiveMediaSourceDownloadBitrateChangedEventArgs2 struct {
	win32.IInspectable
}

func (this *IAdaptiveMediaSourceDownloadBitrateChangedEventArgs2) Vtbl() *IAdaptiveMediaSourceDownloadBitrateChangedEventArgs2Vtbl {
	return (*IAdaptiveMediaSourceDownloadBitrateChangedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveMediaSourceDownloadBitrateChangedEventArgs2) Get_Reason() AdaptiveMediaSourceDownloadBitrateChangedReason {
	var _result AdaptiveMediaSourceDownloadBitrateChangedReason
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Reason, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 19240DC3-5B37-4A1A-8970-D621CB6CA83B
var IID_IAdaptiveMediaSourceDownloadCompletedEventArgs = syscall.GUID{0x19240DC3, 0x5B37, 0x4A1A,
	[8]byte{0x89, 0x70, 0xD6, 0x21, 0xCB, 0x6C, 0xA8, 0x3B}}

type IAdaptiveMediaSourceDownloadCompletedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ResourceType() AdaptiveMediaSourceResourceType
	Get_ResourceUri() *IUriRuntimeClass
	Get_ResourceByteRangeOffset() *IReference[uint64]
	Get_ResourceByteRangeLength() *IReference[uint64]
	Get_HttpResponseMessage() *IHttpResponseMessage
}

type IAdaptiveMediaSourceDownloadCompletedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ResourceType            uintptr
	Get_ResourceUri             uintptr
	Get_ResourceByteRangeOffset uintptr
	Get_ResourceByteRangeLength uintptr
	Get_HttpResponseMessage     uintptr
}

type IAdaptiveMediaSourceDownloadCompletedEventArgs struct {
	win32.IInspectable
}

func (this *IAdaptiveMediaSourceDownloadCompletedEventArgs) Vtbl() *IAdaptiveMediaSourceDownloadCompletedEventArgsVtbl {
	return (*IAdaptiveMediaSourceDownloadCompletedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveMediaSourceDownloadCompletedEventArgs) Get_ResourceType() AdaptiveMediaSourceResourceType {
	var _result AdaptiveMediaSourceResourceType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSourceDownloadCompletedEventArgs) Get_ResourceUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDownloadCompletedEventArgs) Get_ResourceByteRangeOffset() *IReference[uint64] {
	var _result *IReference[uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceByteRangeOffset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDownloadCompletedEventArgs) Get_ResourceByteRangeLength() *IReference[uint64] {
	var _result *IReference[uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceByteRangeLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDownloadCompletedEventArgs) Get_HttpResponseMessage() *IHttpResponseMessage {
	var _result *IHttpResponseMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HttpResponseMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 704744C4-964A-40E4-AF95-9177DD6DFA00
var IID_IAdaptiveMediaSourceDownloadCompletedEventArgs2 = syscall.GUID{0x704744C4, 0x964A, 0x40E4,
	[8]byte{0xAF, 0x95, 0x91, 0x77, 0xDD, 0x6D, 0xFA, 0x00}}

type IAdaptiveMediaSourceDownloadCompletedEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_RequestId() int32
	Get_Statistics() *IAdaptiveMediaSourceDownloadStatistics
	Get_Position() *IReference[TimeSpan]
}

type IAdaptiveMediaSourceDownloadCompletedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_RequestId  uintptr
	Get_Statistics uintptr
	Get_Position   uintptr
}

type IAdaptiveMediaSourceDownloadCompletedEventArgs2 struct {
	win32.IInspectable
}

func (this *IAdaptiveMediaSourceDownloadCompletedEventArgs2) Vtbl() *IAdaptiveMediaSourceDownloadCompletedEventArgs2Vtbl {
	return (*IAdaptiveMediaSourceDownloadCompletedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveMediaSourceDownloadCompletedEventArgs2) Get_RequestId() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSourceDownloadCompletedEventArgs2) Get_Statistics() *IAdaptiveMediaSourceDownloadStatistics {
	var _result *IAdaptiveMediaSourceDownloadStatistics
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Statistics, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDownloadCompletedEventArgs2) Get_Position() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0F8A8BD1-93B2-47C6-BADC-8BE2C8F7F6E8
var IID_IAdaptiveMediaSourceDownloadCompletedEventArgs3 = syscall.GUID{0x0F8A8BD1, 0x93B2, 0x47C6,
	[8]byte{0xBA, 0xDC, 0x8B, 0xE2, 0xC8, 0xF7, 0xF6, 0xE8}}

type IAdaptiveMediaSourceDownloadCompletedEventArgs3Interface interface {
	win32.IInspectableInterface
	Get_ResourceDuration() *IReference[TimeSpan]
	Get_ResourceContentType() string
}

type IAdaptiveMediaSourceDownloadCompletedEventArgs3Vtbl struct {
	win32.IInspectableVtbl
	Get_ResourceDuration    uintptr
	Get_ResourceContentType uintptr
}

type IAdaptiveMediaSourceDownloadCompletedEventArgs3 struct {
	win32.IInspectable
}

func (this *IAdaptiveMediaSourceDownloadCompletedEventArgs3) Vtbl() *IAdaptiveMediaSourceDownloadCompletedEventArgs3Vtbl {
	return (*IAdaptiveMediaSourceDownloadCompletedEventArgs3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveMediaSourceDownloadCompletedEventArgs3) Get_ResourceDuration() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceDuration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDownloadCompletedEventArgs3) Get_ResourceContentType() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceContentType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 37739048-F4AB-40A4-B135-C6DFD8BD7FF1
var IID_IAdaptiveMediaSourceDownloadFailedEventArgs = syscall.GUID{0x37739048, 0xF4AB, 0x40A4,
	[8]byte{0xB1, 0x35, 0xC6, 0xDF, 0xD8, 0xBD, 0x7F, 0xF1}}

type IAdaptiveMediaSourceDownloadFailedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ResourceType() AdaptiveMediaSourceResourceType
	Get_ResourceUri() *IUriRuntimeClass
	Get_ResourceByteRangeOffset() *IReference[uint64]
	Get_ResourceByteRangeLength() *IReference[uint64]
	Get_HttpResponseMessage() *IHttpResponseMessage
}

type IAdaptiveMediaSourceDownloadFailedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ResourceType            uintptr
	Get_ResourceUri             uintptr
	Get_ResourceByteRangeOffset uintptr
	Get_ResourceByteRangeLength uintptr
	Get_HttpResponseMessage     uintptr
}

type IAdaptiveMediaSourceDownloadFailedEventArgs struct {
	win32.IInspectable
}

func (this *IAdaptiveMediaSourceDownloadFailedEventArgs) Vtbl() *IAdaptiveMediaSourceDownloadFailedEventArgsVtbl {
	return (*IAdaptiveMediaSourceDownloadFailedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveMediaSourceDownloadFailedEventArgs) Get_ResourceType() AdaptiveMediaSourceResourceType {
	var _result AdaptiveMediaSourceResourceType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSourceDownloadFailedEventArgs) Get_ResourceUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDownloadFailedEventArgs) Get_ResourceByteRangeOffset() *IReference[uint64] {
	var _result *IReference[uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceByteRangeOffset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDownloadFailedEventArgs) Get_ResourceByteRangeLength() *IReference[uint64] {
	var _result *IReference[uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceByteRangeLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDownloadFailedEventArgs) Get_HttpResponseMessage() *IHttpResponseMessage {
	var _result *IHttpResponseMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HttpResponseMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 70919568-967C-4986-90C5-C6FC4B31E2D8
var IID_IAdaptiveMediaSourceDownloadFailedEventArgs2 = syscall.GUID{0x70919568, 0x967C, 0x4986,
	[8]byte{0x90, 0xC5, 0xC6, 0xFC, 0x4B, 0x31, 0xE2, 0xD8}}

type IAdaptiveMediaSourceDownloadFailedEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_RequestId() int32
	Get_ExtendedError() HResult
	Get_Statistics() *IAdaptiveMediaSourceDownloadStatistics
	Get_Position() *IReference[TimeSpan]
}

type IAdaptiveMediaSourceDownloadFailedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_RequestId     uintptr
	Get_ExtendedError uintptr
	Get_Statistics    uintptr
	Get_Position      uintptr
}

type IAdaptiveMediaSourceDownloadFailedEventArgs2 struct {
	win32.IInspectable
}

func (this *IAdaptiveMediaSourceDownloadFailedEventArgs2) Vtbl() *IAdaptiveMediaSourceDownloadFailedEventArgs2Vtbl {
	return (*IAdaptiveMediaSourceDownloadFailedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveMediaSourceDownloadFailedEventArgs2) Get_RequestId() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSourceDownloadFailedEventArgs2) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSourceDownloadFailedEventArgs2) Get_Statistics() *IAdaptiveMediaSourceDownloadStatistics {
	var _result *IAdaptiveMediaSourceDownloadStatistics
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Statistics, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDownloadFailedEventArgs2) Get_Position() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D0354549-1132-4A10-915A-C2211B5B9409
var IID_IAdaptiveMediaSourceDownloadFailedEventArgs3 = syscall.GUID{0xD0354549, 0x1132, 0x4A10,
	[8]byte{0x91, 0x5A, 0xC2, 0x21, 0x1B, 0x5B, 0x94, 0x09}}

type IAdaptiveMediaSourceDownloadFailedEventArgs3Interface interface {
	win32.IInspectableInterface
	Get_ResourceDuration() *IReference[TimeSpan]
	Get_ResourceContentType() string
}

type IAdaptiveMediaSourceDownloadFailedEventArgs3Vtbl struct {
	win32.IInspectableVtbl
	Get_ResourceDuration    uintptr
	Get_ResourceContentType uintptr
}

type IAdaptiveMediaSourceDownloadFailedEventArgs3 struct {
	win32.IInspectable
}

func (this *IAdaptiveMediaSourceDownloadFailedEventArgs3) Vtbl() *IAdaptiveMediaSourceDownloadFailedEventArgs3Vtbl {
	return (*IAdaptiveMediaSourceDownloadFailedEventArgs3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveMediaSourceDownloadFailedEventArgs3) Get_ResourceDuration() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceDuration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDownloadFailedEventArgs3) Get_ResourceContentType() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceContentType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 05C68F64-FA20-4DBD-9821-4BF4C9BF77AB
var IID_IAdaptiveMediaSourceDownloadRequestedDeferral = syscall.GUID{0x05C68F64, 0xFA20, 0x4DBD,
	[8]byte{0x98, 0x21, 0x4B, 0xF4, 0xC9, 0xBF, 0x77, 0xAB}}

type IAdaptiveMediaSourceDownloadRequestedDeferralInterface interface {
	win32.IInspectableInterface
	Complete()
}

type IAdaptiveMediaSourceDownloadRequestedDeferralVtbl struct {
	win32.IInspectableVtbl
	Complete uintptr
}

type IAdaptiveMediaSourceDownloadRequestedDeferral struct {
	win32.IInspectable
}

func (this *IAdaptiveMediaSourceDownloadRequestedDeferral) Vtbl() *IAdaptiveMediaSourceDownloadRequestedDeferralVtbl {
	return (*IAdaptiveMediaSourceDownloadRequestedDeferralVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveMediaSourceDownloadRequestedDeferral) Complete() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Complete, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// C83FDFFD-44A9-47A2-BF96-03398B4BFAAF
var IID_IAdaptiveMediaSourceDownloadRequestedEventArgs = syscall.GUID{0xC83FDFFD, 0x44A9, 0x47A2,
	[8]byte{0xBF, 0x96, 0x03, 0x39, 0x8B, 0x4B, 0xFA, 0xAF}}

type IAdaptiveMediaSourceDownloadRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ResourceType() AdaptiveMediaSourceResourceType
	Get_ResourceUri() *IUriRuntimeClass
	Get_ResourceByteRangeOffset() *IReference[uint64]
	Get_ResourceByteRangeLength() *IReference[uint64]
	Get_Result() *IAdaptiveMediaSourceDownloadResult
	GetDeferral() *IAdaptiveMediaSourceDownloadRequestedDeferral
}

type IAdaptiveMediaSourceDownloadRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ResourceType            uintptr
	Get_ResourceUri             uintptr
	Get_ResourceByteRangeOffset uintptr
	Get_ResourceByteRangeLength uintptr
	Get_Result                  uintptr
	GetDeferral                 uintptr
}

type IAdaptiveMediaSourceDownloadRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IAdaptiveMediaSourceDownloadRequestedEventArgs) Vtbl() *IAdaptiveMediaSourceDownloadRequestedEventArgsVtbl {
	return (*IAdaptiveMediaSourceDownloadRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveMediaSourceDownloadRequestedEventArgs) Get_ResourceType() AdaptiveMediaSourceResourceType {
	var _result AdaptiveMediaSourceResourceType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSourceDownloadRequestedEventArgs) Get_ResourceUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDownloadRequestedEventArgs) Get_ResourceByteRangeOffset() *IReference[uint64] {
	var _result *IReference[uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceByteRangeOffset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDownloadRequestedEventArgs) Get_ResourceByteRangeLength() *IReference[uint64] {
	var _result *IReference[uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceByteRangeLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDownloadRequestedEventArgs) Get_Result() *IAdaptiveMediaSourceDownloadResult {
	var _result *IAdaptiveMediaSourceDownloadResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Result, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDownloadRequestedEventArgs) GetDeferral() *IAdaptiveMediaSourceDownloadRequestedDeferral {
	var _result *IAdaptiveMediaSourceDownloadRequestedDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B37D8BFE-AA44-4D82-825B-611DE3BCFECB
var IID_IAdaptiveMediaSourceDownloadRequestedEventArgs2 = syscall.GUID{0xB37D8BFE, 0xAA44, 0x4D82,
	[8]byte{0x82, 0x5B, 0x61, 0x1D, 0xE3, 0xBC, 0xFE, 0xCB}}

type IAdaptiveMediaSourceDownloadRequestedEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_RequestId() int32
	Get_Position() *IReference[TimeSpan]
}

type IAdaptiveMediaSourceDownloadRequestedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_RequestId uintptr
	Get_Position  uintptr
}

type IAdaptiveMediaSourceDownloadRequestedEventArgs2 struct {
	win32.IInspectable
}

func (this *IAdaptiveMediaSourceDownloadRequestedEventArgs2) Vtbl() *IAdaptiveMediaSourceDownloadRequestedEventArgs2Vtbl {
	return (*IAdaptiveMediaSourceDownloadRequestedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveMediaSourceDownloadRequestedEventArgs2) Get_RequestId() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSourceDownloadRequestedEventArgs2) Get_Position() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 333C50FD-4F62-4481-AB44-1E47B0574225
var IID_IAdaptiveMediaSourceDownloadRequestedEventArgs3 = syscall.GUID{0x333C50FD, 0x4F62, 0x4481,
	[8]byte{0xAB, 0x44, 0x1E, 0x47, 0xB0, 0x57, 0x42, 0x25}}

type IAdaptiveMediaSourceDownloadRequestedEventArgs3Interface interface {
	win32.IInspectableInterface
	Get_ResourceDuration() *IReference[TimeSpan]
	Get_ResourceContentType() string
}

type IAdaptiveMediaSourceDownloadRequestedEventArgs3Vtbl struct {
	win32.IInspectableVtbl
	Get_ResourceDuration    uintptr
	Get_ResourceContentType uintptr
}

type IAdaptiveMediaSourceDownloadRequestedEventArgs3 struct {
	win32.IInspectable
}

func (this *IAdaptiveMediaSourceDownloadRequestedEventArgs3) Vtbl() *IAdaptiveMediaSourceDownloadRequestedEventArgs3Vtbl {
	return (*IAdaptiveMediaSourceDownloadRequestedEventArgs3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveMediaSourceDownloadRequestedEventArgs3) Get_ResourceDuration() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceDuration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDownloadRequestedEventArgs3) Get_ResourceContentType() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceContentType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// F4AFDC73-BCEE-4A6A-9F0A-FEC41E2339B0
var IID_IAdaptiveMediaSourceDownloadResult = syscall.GUID{0xF4AFDC73, 0xBCEE, 0x4A6A,
	[8]byte{0x9F, 0x0A, 0xFE, 0xC4, 0x1E, 0x23, 0x39, 0xB0}}

type IAdaptiveMediaSourceDownloadResultInterface interface {
	win32.IInspectableInterface
	Get_ResourceUri() *IUriRuntimeClass
	Put_ResourceUri(value *IUriRuntimeClass)
	Get_InputStream() *IInputStream
	Put_InputStream(value *IInputStream)
	Get_Buffer() *IBuffer
	Put_Buffer(value *IBuffer)
	Get_ContentType() string
	Put_ContentType(value string)
	Get_ExtendedStatus() uint32
	Put_ExtendedStatus(value uint32)
}

type IAdaptiveMediaSourceDownloadResultVtbl struct {
	win32.IInspectableVtbl
	Get_ResourceUri    uintptr
	Put_ResourceUri    uintptr
	Get_InputStream    uintptr
	Put_InputStream    uintptr
	Get_Buffer         uintptr
	Put_Buffer         uintptr
	Get_ContentType    uintptr
	Put_ContentType    uintptr
	Get_ExtendedStatus uintptr
	Put_ExtendedStatus uintptr
}

type IAdaptiveMediaSourceDownloadResult struct {
	win32.IInspectable
}

func (this *IAdaptiveMediaSourceDownloadResult) Vtbl() *IAdaptiveMediaSourceDownloadResultVtbl {
	return (*IAdaptiveMediaSourceDownloadResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveMediaSourceDownloadResult) Get_ResourceUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDownloadResult) Put_ResourceUri(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ResourceUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAdaptiveMediaSourceDownloadResult) Get_InputStream() *IInputStream {
	var _result *IInputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InputStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDownloadResult) Put_InputStream(value *IInputStream) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InputStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAdaptiveMediaSourceDownloadResult) Get_Buffer() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Buffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDownloadResult) Put_Buffer(value *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Buffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAdaptiveMediaSourceDownloadResult) Get_ContentType() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAdaptiveMediaSourceDownloadResult) Put_ContentType(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ContentType, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IAdaptiveMediaSourceDownloadResult) Get_ExtendedStatus() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSourceDownloadResult) Put_ExtendedStatus(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ExtendedStatus, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 15552CB7-7B80-4AC4-8660-A4B97F7C70F0
var IID_IAdaptiveMediaSourceDownloadResult2 = syscall.GUID{0x15552CB7, 0x7B80, 0x4AC4,
	[8]byte{0x86, 0x60, 0xA4, 0xB9, 0x7F, 0x7C, 0x70, 0xF0}}

type IAdaptiveMediaSourceDownloadResult2Interface interface {
	win32.IInspectableInterface
	Get_ResourceByteRangeOffset() *IReference[uint64]
	Put_ResourceByteRangeOffset(value *IReference[uint64])
	Get_ResourceByteRangeLength() *IReference[uint64]
	Put_ResourceByteRangeLength(value *IReference[uint64])
}

type IAdaptiveMediaSourceDownloadResult2Vtbl struct {
	win32.IInspectableVtbl
	Get_ResourceByteRangeOffset uintptr
	Put_ResourceByteRangeOffset uintptr
	Get_ResourceByteRangeLength uintptr
	Put_ResourceByteRangeLength uintptr
}

type IAdaptiveMediaSourceDownloadResult2 struct {
	win32.IInspectable
}

func (this *IAdaptiveMediaSourceDownloadResult2) Vtbl() *IAdaptiveMediaSourceDownloadResult2Vtbl {
	return (*IAdaptiveMediaSourceDownloadResult2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveMediaSourceDownloadResult2) Get_ResourceByteRangeOffset() *IReference[uint64] {
	var _result *IReference[uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceByteRangeOffset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDownloadResult2) Put_ResourceByteRangeOffset(value *IReference[uint64]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ResourceByteRangeOffset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAdaptiveMediaSourceDownloadResult2) Get_ResourceByteRangeLength() *IReference[uint64] {
	var _result *IReference[uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceByteRangeLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDownloadResult2) Put_ResourceByteRangeLength(value *IReference[uint64]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ResourceByteRangeLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// A306CEFB-E96A-4DFF-A9B8-1AE08C01AE98
var IID_IAdaptiveMediaSourceDownloadStatistics = syscall.GUID{0xA306CEFB, 0xE96A, 0x4DFF,
	[8]byte{0xA9, 0xB8, 0x1A, 0xE0, 0x8C, 0x01, 0xAE, 0x98}}

type IAdaptiveMediaSourceDownloadStatisticsInterface interface {
	win32.IInspectableInterface
	Get_ContentBytesReceivedCount() uint64
	Get_TimeToHeadersReceived() *IReference[TimeSpan]
	Get_TimeToFirstByteReceived() *IReference[TimeSpan]
	Get_TimeToLastByteReceived() *IReference[TimeSpan]
}

type IAdaptiveMediaSourceDownloadStatisticsVtbl struct {
	win32.IInspectableVtbl
	Get_ContentBytesReceivedCount uintptr
	Get_TimeToHeadersReceived     uintptr
	Get_TimeToFirstByteReceived   uintptr
	Get_TimeToLastByteReceived    uintptr
}

type IAdaptiveMediaSourceDownloadStatistics struct {
	win32.IInspectable
}

func (this *IAdaptiveMediaSourceDownloadStatistics) Vtbl() *IAdaptiveMediaSourceDownloadStatisticsVtbl {
	return (*IAdaptiveMediaSourceDownloadStatisticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveMediaSourceDownloadStatistics) Get_ContentBytesReceivedCount() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentBytesReceivedCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSourceDownloadStatistics) Get_TimeToHeadersReceived() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TimeToHeadersReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDownloadStatistics) Get_TimeToFirstByteReceived() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TimeToFirstByteReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceDownloadStatistics) Get_TimeToLastByteReceived() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TimeToLastByteReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 23A29F6D-7DDA-4A51-87A9-6FA8C5B292BE
var IID_IAdaptiveMediaSourcePlaybackBitrateChangedEventArgs = syscall.GUID{0x23A29F6D, 0x7DDA, 0x4A51,
	[8]byte{0x87, 0xA9, 0x6F, 0xA8, 0xC5, 0xB2, 0x92, 0xBE}}

type IAdaptiveMediaSourcePlaybackBitrateChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_OldValue() uint32
	Get_NewValue() uint32
	Get_AudioOnly() bool
}

type IAdaptiveMediaSourcePlaybackBitrateChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_OldValue  uintptr
	Get_NewValue  uintptr
	Get_AudioOnly uintptr
}

type IAdaptiveMediaSourcePlaybackBitrateChangedEventArgs struct {
	win32.IInspectable
}

func (this *IAdaptiveMediaSourcePlaybackBitrateChangedEventArgs) Vtbl() *IAdaptiveMediaSourcePlaybackBitrateChangedEventArgsVtbl {
	return (*IAdaptiveMediaSourcePlaybackBitrateChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveMediaSourcePlaybackBitrateChangedEventArgs) Get_OldValue() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OldValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSourcePlaybackBitrateChangedEventArgs) Get_NewValue() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NewValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSourcePlaybackBitrateChangedEventArgs) Get_AudioOnly() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioOnly, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 50A6BD5D-66EF-4CD3-9579-9E660507DC3F
var IID_IAdaptiveMediaSourceStatics = syscall.GUID{0x50A6BD5D, 0x66EF, 0x4CD3,
	[8]byte{0x95, 0x79, 0x9E, 0x66, 0x05, 0x07, 0xDC, 0x3F}}

type IAdaptiveMediaSourceStaticsInterface interface {
	win32.IInspectableInterface
	IsContentTypeSupported(contentType string) bool
	CreateFromUriAsync(uri *IUriRuntimeClass) *IAsyncOperation[*IAdaptiveMediaSourceCreationResult]
	CreateFromUriWithDownloaderAsync(uri *IUriRuntimeClass, httpClient *IHttpClient) *IAsyncOperation[*IAdaptiveMediaSourceCreationResult]
	CreateFromStreamAsync(stream *IInputStream, uri *IUriRuntimeClass, contentType string) *IAsyncOperation[*IAdaptiveMediaSourceCreationResult]
	CreateFromStreamWithDownloaderAsync(stream *IInputStream, uri *IUriRuntimeClass, contentType string, httpClient *IHttpClient) *IAsyncOperation[*IAdaptiveMediaSourceCreationResult]
}

type IAdaptiveMediaSourceStaticsVtbl struct {
	win32.IInspectableVtbl
	IsContentTypeSupported              uintptr
	CreateFromUriAsync                  uintptr
	CreateFromUriWithDownloaderAsync    uintptr
	CreateFromStreamAsync               uintptr
	CreateFromStreamWithDownloaderAsync uintptr
}

type IAdaptiveMediaSourceStatics struct {
	win32.IInspectable
}

func (this *IAdaptiveMediaSourceStatics) Vtbl() *IAdaptiveMediaSourceStaticsVtbl {
	return (*IAdaptiveMediaSourceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdaptiveMediaSourceStatics) IsContentTypeSupported(contentType string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsContentTypeSupported, uintptr(unsafe.Pointer(this)), NewHStr(contentType).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdaptiveMediaSourceStatics) CreateFromUriAsync(uri *IUriRuntimeClass) *IAsyncOperation[*IAdaptiveMediaSourceCreationResult] {
	var _result *IAsyncOperation[*IAdaptiveMediaSourceCreationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromUriAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceStatics) CreateFromUriWithDownloaderAsync(uri *IUriRuntimeClass, httpClient *IHttpClient) *IAsyncOperation[*IAdaptiveMediaSourceCreationResult] {
	var _result *IAsyncOperation[*IAdaptiveMediaSourceCreationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromUriWithDownloaderAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(httpClient)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceStatics) CreateFromStreamAsync(stream *IInputStream, uri *IUriRuntimeClass, contentType string) *IAsyncOperation[*IAdaptiveMediaSourceCreationResult] {
	var _result *IAsyncOperation[*IAdaptiveMediaSourceCreationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromStreamAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(uri)), NewHStr(contentType).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdaptiveMediaSourceStatics) CreateFromStreamWithDownloaderAsync(stream *IInputStream, uri *IUriRuntimeClass, contentType string, httpClient *IHttpClient) *IAsyncOperation[*IAdaptiveMediaSourceCreationResult] {
	var _result *IAsyncOperation[*IAdaptiveMediaSourceCreationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromStreamWithDownloaderAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(uri)), NewHStr(contentType).Ptr, uintptr(unsafe.Pointer(httpClient)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type AdaptiveMediaSource struct {
	RtClass
	*IAdaptiveMediaSource
}

func NewIAdaptiveMediaSourceStatics() *IAdaptiveMediaSourceStatics {
	var p *IAdaptiveMediaSourceStatics
	hs := NewHStr("Windows.Media.Streaming.Adaptive.AdaptiveMediaSource")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAdaptiveMediaSourceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type AdaptiveMediaSourceAdvancedSettings struct {
	RtClass
	*IAdaptiveMediaSourceAdvancedSettings
}

type AdaptiveMediaSourceCorrelatedTimes struct {
	RtClass
	*IAdaptiveMediaSourceCorrelatedTimes
}

type AdaptiveMediaSourceCreationResult struct {
	RtClass
	*IAdaptiveMediaSourceCreationResult
}

type AdaptiveMediaSourceDiagnosticAvailableEventArgs struct {
	RtClass
	*IAdaptiveMediaSourceDiagnosticAvailableEventArgs
}

type AdaptiveMediaSourceDiagnostics struct {
	RtClass
	*IAdaptiveMediaSourceDiagnostics
}

type AdaptiveMediaSourceDownloadBitrateChangedEventArgs struct {
	RtClass
	*IAdaptiveMediaSourceDownloadBitrateChangedEventArgs
}

type AdaptiveMediaSourceDownloadCompletedEventArgs struct {
	RtClass
	*IAdaptiveMediaSourceDownloadCompletedEventArgs
}

type AdaptiveMediaSourceDownloadFailedEventArgs struct {
	RtClass
	*IAdaptiveMediaSourceDownloadFailedEventArgs
}

type AdaptiveMediaSourceDownloadRequestedDeferral struct {
	RtClass
	*IAdaptiveMediaSourceDownloadRequestedDeferral
}

type AdaptiveMediaSourceDownloadRequestedEventArgs struct {
	RtClass
	*IAdaptiveMediaSourceDownloadRequestedEventArgs
}

type AdaptiveMediaSourceDownloadResult struct {
	RtClass
	*IAdaptiveMediaSourceDownloadResult
}

type AdaptiveMediaSourceDownloadStatistics struct {
	RtClass
	*IAdaptiveMediaSourceDownloadStatistics
}

type AdaptiveMediaSourcePlaybackBitrateChangedEventArgs struct {
	RtClass
	*IAdaptiveMediaSourcePlaybackBitrateChangedEventArgs
}
