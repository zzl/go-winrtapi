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
type AudioDecoderDegradation int32

const (
	AudioDecoderDegradation_None               AudioDecoderDegradation = 0
	AudioDecoderDegradation_DownmixTo2Channels AudioDecoderDegradation = 1
	AudioDecoderDegradation_DownmixTo6Channels AudioDecoderDegradation = 2
	AudioDecoderDegradation_DownmixTo8Channels AudioDecoderDegradation = 3
)

// enum
type AudioDecoderDegradationReason int32

const (
	AudioDecoderDegradationReason_None                     AudioDecoderDegradationReason = 0
	AudioDecoderDegradationReason_LicensingRequirement     AudioDecoderDegradationReason = 1
	AudioDecoderDegradationReason_SpatialAudioNotSupported AudioDecoderDegradationReason = 2
)

// enum
type CodecCategory int32

const (
	CodecCategory_Encoder CodecCategory = 0
	CodecCategory_Decoder CodecCategory = 1
)

// enum
type CodecKind int32

const (
	CodecKind_Audio CodecKind = 0
	CodecKind_Video CodecKind = 1
)

// enum
type FaceDetectionMode int32

const (
	FaceDetectionMode_HighPerformance FaceDetectionMode = 0
	FaceDetectionMode_Balanced        FaceDetectionMode = 1
	FaceDetectionMode_HighQuality     FaceDetectionMode = 2
)

// enum
type MediaDecoderStatus int32

const (
	MediaDecoderStatus_FullySupported               MediaDecoderStatus = 0
	MediaDecoderStatus_UnsupportedSubtype           MediaDecoderStatus = 1
	MediaDecoderStatus_UnsupportedEncoderProperties MediaDecoderStatus = 2
	MediaDecoderStatus_Degraded                     MediaDecoderStatus = 3
)

// enum
type MediaSourceState int32

const (
	MediaSourceState_Initial MediaSourceState = 0
	MediaSourceState_Opening MediaSourceState = 1
	MediaSourceState_Opened  MediaSourceState = 2
	MediaSourceState_Failed  MediaSourceState = 3
	MediaSourceState_Closed  MediaSourceState = 4
)

// enum
type MediaSourceStatus int32

const (
	MediaSourceStatus_FullySupported MediaSourceStatus = 0
	MediaSourceStatus_Unknown        MediaSourceStatus = 1
)

// enum
type MediaStreamSourceClosedReason int32

const (
	MediaStreamSourceClosedReason_Done                               MediaStreamSourceClosedReason = 0
	MediaStreamSourceClosedReason_UnknownError                       MediaStreamSourceClosedReason = 1
	MediaStreamSourceClosedReason_AppReportedError                   MediaStreamSourceClosedReason = 2
	MediaStreamSourceClosedReason_UnsupportedProtectionSystem        MediaStreamSourceClosedReason = 3
	MediaStreamSourceClosedReason_ProtectionSystemFailure            MediaStreamSourceClosedReason = 4
	MediaStreamSourceClosedReason_UnsupportedEncodingFormat          MediaStreamSourceClosedReason = 5
	MediaStreamSourceClosedReason_MissingSampleRequestedEventHandler MediaStreamSourceClosedReason = 6
)

// enum
type MediaStreamSourceErrorStatus int32

const (
	MediaStreamSourceErrorStatus_Other                   MediaStreamSourceErrorStatus = 0
	MediaStreamSourceErrorStatus_OutOfMemory             MediaStreamSourceErrorStatus = 1
	MediaStreamSourceErrorStatus_FailedToOpenFile        MediaStreamSourceErrorStatus = 2
	MediaStreamSourceErrorStatus_FailedToConnectToServer MediaStreamSourceErrorStatus = 3
	MediaStreamSourceErrorStatus_ConnectionToServerLost  MediaStreamSourceErrorStatus = 4
	MediaStreamSourceErrorStatus_UnspecifiedNetworkError MediaStreamSourceErrorStatus = 5
	MediaStreamSourceErrorStatus_DecodeError             MediaStreamSourceErrorStatus = 6
	MediaStreamSourceErrorStatus_UnsupportedMediaFormat  MediaStreamSourceErrorStatus = 7
)

// enum
type MediaTrackKind int32

const (
	MediaTrackKind_Audio         MediaTrackKind = 0
	MediaTrackKind_Video         MediaTrackKind = 1
	MediaTrackKind_TimedMetadata MediaTrackKind = 2
)

// enum
type MseAppendMode int32

const (
	MseAppendMode_Segments MseAppendMode = 0
	MseAppendMode_Sequence MseAppendMode = 1
)

// enum
type MseEndOfStreamStatus int32

const (
	MseEndOfStreamStatus_Success      MseEndOfStreamStatus = 0
	MseEndOfStreamStatus_NetworkError MseEndOfStreamStatus = 1
	MseEndOfStreamStatus_DecodeError  MseEndOfStreamStatus = 2
	MseEndOfStreamStatus_UnknownError MseEndOfStreamStatus = 3
)

// enum
type MseReadyState int32

const (
	MseReadyState_Closed MseReadyState = 0
	MseReadyState_Open   MseReadyState = 1
	MseReadyState_Ended  MseReadyState = 2
)

// enum
type SceneAnalysisRecommendation int32

const (
	SceneAnalysisRecommendation_Standard SceneAnalysisRecommendation = 0
	SceneAnalysisRecommendation_Hdr      SceneAnalysisRecommendation = 1
	SceneAnalysisRecommendation_LowLight SceneAnalysisRecommendation = 2
)

// enum
type TimedMetadataKind int32

const (
	TimedMetadataKind_Caption       TimedMetadataKind = 0
	TimedMetadataKind_Chapter       TimedMetadataKind = 1
	TimedMetadataKind_Custom        TimedMetadataKind = 2
	TimedMetadataKind_Data          TimedMetadataKind = 3
	TimedMetadataKind_Description   TimedMetadataKind = 4
	TimedMetadataKind_Subtitle      TimedMetadataKind = 5
	TimedMetadataKind_ImageSubtitle TimedMetadataKind = 6
	TimedMetadataKind_Speech        TimedMetadataKind = 7
)

// enum
type TimedMetadataTrackErrorCode int32

const (
	TimedMetadataTrackErrorCode_None            TimedMetadataTrackErrorCode = 0
	TimedMetadataTrackErrorCode_DataFormatError TimedMetadataTrackErrorCode = 1
	TimedMetadataTrackErrorCode_NetworkError    TimedMetadataTrackErrorCode = 2
	TimedMetadataTrackErrorCode_InternalError   TimedMetadataTrackErrorCode = 3
)

// enum
type TimedTextBoutenPosition int32

const (
	TimedTextBoutenPosition_Before  TimedTextBoutenPosition = 0
	TimedTextBoutenPosition_After   TimedTextBoutenPosition = 1
	TimedTextBoutenPosition_Outside TimedTextBoutenPosition = 2
)

// enum
type TimedTextBoutenType int32

const (
	TimedTextBoutenType_None         TimedTextBoutenType = 0
	TimedTextBoutenType_Auto         TimedTextBoutenType = 1
	TimedTextBoutenType_FilledCircle TimedTextBoutenType = 2
	TimedTextBoutenType_OpenCircle   TimedTextBoutenType = 3
	TimedTextBoutenType_FilledDot    TimedTextBoutenType = 4
	TimedTextBoutenType_OpenDot      TimedTextBoutenType = 5
	TimedTextBoutenType_FilledSesame TimedTextBoutenType = 6
	TimedTextBoutenType_OpenSesame   TimedTextBoutenType = 7
)

// enum
type TimedTextDisplayAlignment int32

const (
	TimedTextDisplayAlignment_Before TimedTextDisplayAlignment = 0
	TimedTextDisplayAlignment_After  TimedTextDisplayAlignment = 1
	TimedTextDisplayAlignment_Center TimedTextDisplayAlignment = 2
)

// enum
type TimedTextFlowDirection int32

const (
	TimedTextFlowDirection_LeftToRight TimedTextFlowDirection = 0
	TimedTextFlowDirection_RightToLeft TimedTextFlowDirection = 1
)

// enum
type TimedTextFontStyle int32

const (
	TimedTextFontStyle_Normal  TimedTextFontStyle = 0
	TimedTextFontStyle_Oblique TimedTextFontStyle = 1
	TimedTextFontStyle_Italic  TimedTextFontStyle = 2
)

// enum
type TimedTextLineAlignment int32

const (
	TimedTextLineAlignment_Start  TimedTextLineAlignment = 0
	TimedTextLineAlignment_End    TimedTextLineAlignment = 1
	TimedTextLineAlignment_Center TimedTextLineAlignment = 2
)

// enum
type TimedTextRubyAlign int32

const (
	TimedTextRubyAlign_Center       TimedTextRubyAlign = 0
	TimedTextRubyAlign_Start        TimedTextRubyAlign = 1
	TimedTextRubyAlign_End          TimedTextRubyAlign = 2
	TimedTextRubyAlign_SpaceAround  TimedTextRubyAlign = 3
	TimedTextRubyAlign_SpaceBetween TimedTextRubyAlign = 4
	TimedTextRubyAlign_WithBase     TimedTextRubyAlign = 5
)

// enum
type TimedTextRubyPosition int32

const (
	TimedTextRubyPosition_Before  TimedTextRubyPosition = 0
	TimedTextRubyPosition_After   TimedTextRubyPosition = 1
	TimedTextRubyPosition_Outside TimedTextRubyPosition = 2
)

// enum
type TimedTextRubyReserve int32

const (
	TimedTextRubyReserve_None    TimedTextRubyReserve = 0
	TimedTextRubyReserve_Before  TimedTextRubyReserve = 1
	TimedTextRubyReserve_After   TimedTextRubyReserve = 2
	TimedTextRubyReserve_Both    TimedTextRubyReserve = 3
	TimedTextRubyReserve_Outside TimedTextRubyReserve = 4
)

// enum
type TimedTextScrollMode int32

const (
	TimedTextScrollMode_Popon  TimedTextScrollMode = 0
	TimedTextScrollMode_Rollup TimedTextScrollMode = 1
)

// enum
type TimedTextUnit int32

const (
	TimedTextUnit_Pixels     TimedTextUnit = 0
	TimedTextUnit_Percentage TimedTextUnit = 1
)

// enum
type TimedTextWeight int32

const (
	TimedTextWeight_Normal TimedTextWeight = 400
	TimedTextWeight_Bold   TimedTextWeight = 700
)

// enum
type TimedTextWrapping int32

const (
	TimedTextWrapping_NoWrap TimedTextWrapping = 0
	TimedTextWrapping_Wrap   TimedTextWrapping = 1
)

// enum
type TimedTextWritingMode int32

const (
	TimedTextWritingMode_LeftRightTopBottom TimedTextWritingMode = 0
	TimedTextWritingMode_RightLeftTopBottom TimedTextWritingMode = 1
	TimedTextWritingMode_TopBottomRightLeft TimedTextWritingMode = 2
	TimedTextWritingMode_TopBottomLeftRight TimedTextWritingMode = 3
	TimedTextWritingMode_LeftRight          TimedTextWritingMode = 4
	TimedTextWritingMode_RightLeft          TimedTextWritingMode = 5
	TimedTextWritingMode_TopBottom          TimedTextWritingMode = 6
)

// enum
type VideoStabilizationEffectEnabledChangedReason int32

const (
	VideoStabilizationEffectEnabledChangedReason_Programmatic     VideoStabilizationEffectEnabledChangedReason = 0
	VideoStabilizationEffectEnabledChangedReason_PixelRateTooHigh VideoStabilizationEffectEnabledChangedReason = 1
	VideoStabilizationEffectEnabledChangedReason_RunningSlowly    VideoStabilizationEffectEnabledChangedReason = 2
)

// structs

type MseTimeRange struct {
	Start TimeSpan
	End   TimeSpan
}

type TimedTextDouble struct {
	Value float64
	Unit  TimedTextUnit
}

type TimedTextPadding struct {
	Before float64
	After  float64
	Start  float64
	End    float64
	Unit   TimedTextUnit
}

type TimedTextPoint struct {
	X    float64
	Y    float64
	Unit TimedTextUnit
}

type TimedTextSize struct {
	Height float64
	Width  float64
	Unit   TimedTextUnit
}

// interfaces

// 1E3692E4-4027-4847-A70B-DF1D9A2A7B04
var IID_IAudioStreamDescriptor = syscall.GUID{0x1E3692E4, 0x4027, 0x4847,
	[8]byte{0xA7, 0x0B, 0xDF, 0x1D, 0x9A, 0x2A, 0x7B, 0x04}}

type IAudioStreamDescriptorInterface interface {
	win32.IInspectableInterface
	Get_EncodingProperties() *IAudioEncodingProperties
}

type IAudioStreamDescriptorVtbl struct {
	win32.IInspectableVtbl
	Get_EncodingProperties uintptr
}

type IAudioStreamDescriptor struct {
	win32.IInspectable
}

func (this *IAudioStreamDescriptor) Vtbl() *IAudioStreamDescriptorVtbl {
	return (*IAudioStreamDescriptorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioStreamDescriptor) Get_EncodingProperties() *IAudioEncodingProperties {
	var _result *IAudioEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EncodingProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2E68F1F6-A448-497B-8840-85082665ACF9
var IID_IAudioStreamDescriptor2 = syscall.GUID{0x2E68F1F6, 0xA448, 0x497B,
	[8]byte{0x88, 0x40, 0x85, 0x08, 0x26, 0x65, 0xAC, 0xF9}}

type IAudioStreamDescriptor2Interface interface {
	win32.IInspectableInterface
	Put_LeadingEncoderPadding(value *IReference[uint32])
	Get_LeadingEncoderPadding() *IReference[uint32]
	Put_TrailingEncoderPadding(value *IReference[uint32])
	Get_TrailingEncoderPadding() *IReference[uint32]
}

type IAudioStreamDescriptor2Vtbl struct {
	win32.IInspectableVtbl
	Put_LeadingEncoderPadding  uintptr
	Get_LeadingEncoderPadding  uintptr
	Put_TrailingEncoderPadding uintptr
	Get_TrailingEncoderPadding uintptr
}

type IAudioStreamDescriptor2 struct {
	win32.IInspectable
}

func (this *IAudioStreamDescriptor2) Vtbl() *IAudioStreamDescriptor2Vtbl {
	return (*IAudioStreamDescriptor2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioStreamDescriptor2) Put_LeadingEncoderPadding(value *IReference[uint32]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LeadingEncoderPadding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAudioStreamDescriptor2) Get_LeadingEncoderPadding() *IReference[uint32] {
	var _result *IReference[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LeadingEncoderPadding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioStreamDescriptor2) Put_TrailingEncoderPadding(value *IReference[uint32]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TrailingEncoderPadding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAudioStreamDescriptor2) Get_TrailingEncoderPadding() *IReference[uint32] {
	var _result *IReference[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TrailingEncoderPadding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4D220DA1-8E83-44EF-8973-2F63E993F36B
var IID_IAudioStreamDescriptor3 = syscall.GUID{0x4D220DA1, 0x8E83, 0x44EF,
	[8]byte{0x89, 0x73, 0x2F, 0x63, 0xE9, 0x93, 0xF3, 0x6B}}

type IAudioStreamDescriptor3Interface interface {
	win32.IInspectableInterface
	Copy() *IAudioStreamDescriptor
}

type IAudioStreamDescriptor3Vtbl struct {
	win32.IInspectableVtbl
	Copy uintptr
}

type IAudioStreamDescriptor3 struct {
	win32.IInspectable
}

func (this *IAudioStreamDescriptor3) Vtbl() *IAudioStreamDescriptor3Vtbl {
	return (*IAudioStreamDescriptor3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioStreamDescriptor3) Copy() *IAudioStreamDescriptor {
	var _result *IAudioStreamDescriptor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Copy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4A86CE9E-4CB1-4380-8E0C-83504B7F5BF3
var IID_IAudioStreamDescriptorFactory = syscall.GUID{0x4A86CE9E, 0x4CB1, 0x4380,
	[8]byte{0x8E, 0x0C, 0x83, 0x50, 0x4B, 0x7F, 0x5B, 0xF3}}

type IAudioStreamDescriptorFactoryInterface interface {
	win32.IInspectableInterface
	Create(encodingProperties *IAudioEncodingProperties) *IAudioStreamDescriptor
}

type IAudioStreamDescriptorFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IAudioStreamDescriptorFactory struct {
	win32.IInspectable
}

func (this *IAudioStreamDescriptorFactory) Vtbl() *IAudioStreamDescriptorFactoryVtbl {
	return (*IAudioStreamDescriptorFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioStreamDescriptorFactory) Create(encodingProperties *IAudioEncodingProperties) *IAudioStreamDescriptor {
	var _result *IAudioStreamDescriptor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(encodingProperties)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F23B6E77-3EF7-40DE-B943-068B1321701D
var IID_IAudioTrack = syscall.GUID{0xF23B6E77, 0x3EF7, 0x40DE,
	[8]byte{0xB9, 0x43, 0x06, 0x8B, 0x13, 0x21, 0x70, 0x1D}}

type IAudioTrackInterface interface {
	win32.IInspectableInterface
	Add_OpenFailed(handler TypedEventHandler[*IMediaTrack, *IAudioTrackOpenFailedEventArgs]) EventRegistrationToken
	Remove_OpenFailed(token EventRegistrationToken)
	GetEncodingProperties() *IAudioEncodingProperties
	Get_PlaybackItem() *IMediaPlaybackItem
	Get_Name() string
	Get_SupportInfo() *IAudioTrackSupportInfo
}

type IAudioTrackVtbl struct {
	win32.IInspectableVtbl
	Add_OpenFailed        uintptr
	Remove_OpenFailed     uintptr
	GetEncodingProperties uintptr
	Get_PlaybackItem      uintptr
	Get_Name              uintptr
	Get_SupportInfo       uintptr
}

type IAudioTrack struct {
	win32.IInspectable
}

func (this *IAudioTrack) Vtbl() *IAudioTrackVtbl {
	return (*IAudioTrackVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioTrack) Add_OpenFailed(handler TypedEventHandler[*IMediaTrack, *IAudioTrackOpenFailedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_OpenFailed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioTrack) Remove_OpenFailed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_OpenFailed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAudioTrack) GetEncodingProperties() *IAudioEncodingProperties {
	var _result *IAudioEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetEncodingProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioTrack) Get_PlaybackItem() *IMediaPlaybackItem {
	var _result *IMediaPlaybackItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlaybackItem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioTrack) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAudioTrack) Get_SupportInfo() *IAudioTrackSupportInfo {
	var _result *IAudioTrackSupportInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EEDDB9B9-BB7C-4112-BF76-9384676F824B
var IID_IAudioTrackOpenFailedEventArgs = syscall.GUID{0xEEDDB9B9, 0xBB7C, 0x4112,
	[8]byte{0xBF, 0x76, 0x93, 0x84, 0x67, 0x6F, 0x82, 0x4B}}

type IAudioTrackOpenFailedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ExtendedError() HResult
}

type IAudioTrackOpenFailedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ExtendedError uintptr
}

type IAudioTrackOpenFailedEventArgs struct {
	win32.IInspectable
}

func (this *IAudioTrackOpenFailedEventArgs) Vtbl() *IAudioTrackOpenFailedEventArgsVtbl {
	return (*IAudioTrackOpenFailedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioTrackOpenFailedEventArgs) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 178BEFF7-CC39-44A6-B951-4A5653F073FA
var IID_IAudioTrackSupportInfo = syscall.GUID{0x178BEFF7, 0xCC39, 0x44A6,
	[8]byte{0xB9, 0x51, 0x4A, 0x56, 0x53, 0xF0, 0x73, 0xFA}}

type IAudioTrackSupportInfoInterface interface {
	win32.IInspectableInterface
	Get_DecoderStatus() MediaDecoderStatus
	Get_Degradation() AudioDecoderDegradation
	Get_DegradationReason() AudioDecoderDegradationReason
	Get_MediaSourceStatus() MediaSourceStatus
}

type IAudioTrackSupportInfoVtbl struct {
	win32.IInspectableVtbl
	Get_DecoderStatus     uintptr
	Get_Degradation       uintptr
	Get_DegradationReason uintptr
	Get_MediaSourceStatus uintptr
}

type IAudioTrackSupportInfo struct {
	win32.IInspectable
}

func (this *IAudioTrackSupportInfo) Vtbl() *IAudioTrackSupportInfoVtbl {
	return (*IAudioTrackSupportInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioTrackSupportInfo) Get_DecoderStatus() MediaDecoderStatus {
	var _result MediaDecoderStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DecoderStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioTrackSupportInfo) Get_Degradation() AudioDecoderDegradation {
	var _result AudioDecoderDegradation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Degradation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioTrackSupportInfo) Get_DegradationReason() AudioDecoderDegradationReason {
	var _result AudioDecoderDegradationReason
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DegradationReason, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioTrackSupportInfo) Get_MediaSourceStatus() MediaSourceStatus {
	var _result MediaSourceStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaSourceStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 72A98001-D38A-4C0A-8FA6-75CDDAF4664C
var IID_IChapterCue = syscall.GUID{0x72A98001, 0xD38A, 0x4C0A,
	[8]byte{0x8F, 0xA6, 0x75, 0xCD, 0xDA, 0xF4, 0x66, 0x4C}}

type IChapterCueInterface interface {
	win32.IInspectableInterface
	Put_Title(value string)
	Get_Title() string
}

type IChapterCueVtbl struct {
	win32.IInspectableVtbl
	Put_Title uintptr
	Get_Title uintptr
}

type IChapterCue struct {
	win32.IInspectable
}

func (this *IChapterCue) Vtbl() *IChapterCueVtbl {
	return (*IChapterCueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IChapterCue) Put_Title(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Title, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IChapterCue) Get_Title() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Title, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 51E89F85-EA97-499C-86AC-4CE5E73F3A42
var IID_ICodecInfo = syscall.GUID{0x51E89F85, 0xEA97, 0x499C,
	[8]byte{0x86, 0xAC, 0x4C, 0xE5, 0xE7, 0x3F, 0x3A, 0x42}}

type ICodecInfoInterface interface {
	win32.IInspectableInterface
	Get_Kind() CodecKind
	Get_Category() CodecCategory
	Get_Subtypes() *IVectorView[string]
	Get_DisplayName() string
	Get_IsTrusted() bool
}

type ICodecInfoVtbl struct {
	win32.IInspectableVtbl
	Get_Kind        uintptr
	Get_Category    uintptr
	Get_Subtypes    uintptr
	Get_DisplayName uintptr
	Get_IsTrusted   uintptr
}

type ICodecInfo struct {
	win32.IInspectable
}

func (this *ICodecInfo) Vtbl() *ICodecInfoVtbl {
	return (*ICodecInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICodecInfo) Get_Kind() CodecKind {
	var _result CodecKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Kind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICodecInfo) Get_Category() CodecCategory {
	var _result CodecCategory
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Category, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICodecInfo) Get_Subtypes() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Subtypes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICodecInfo) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecInfo) Get_IsTrusted() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsTrusted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 222A953A-AF61-4E04-808A-A4634E2F3AC4
var IID_ICodecQuery = syscall.GUID{0x222A953A, 0xAF61, 0x4E04,
	[8]byte{0x80, 0x8A, 0xA4, 0x63, 0x4E, 0x2F, 0x3A, 0xC4}}

type ICodecQueryInterface interface {
	win32.IInspectableInterface
	FindAllAsync(kind CodecKind, category CodecCategory, subType string) *IAsyncOperation[*IVectorView[*ICodecInfo]]
}

type ICodecQueryVtbl struct {
	win32.IInspectableVtbl
	FindAllAsync uintptr
}

type ICodecQuery struct {
	win32.IInspectable
}

func (this *ICodecQuery) Vtbl() *ICodecQueryVtbl {
	return (*ICodecQueryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICodecQuery) FindAllAsync(kind CodecKind, category CodecCategory, subType string) *IAsyncOperation[*IVectorView[*ICodecInfo]] {
	var _result *IAsyncOperation[*IVectorView[*ICodecInfo]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllAsync, uintptr(unsafe.Pointer(this)), uintptr(kind), uintptr(category), NewHStr(subType).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A66AC4F2-888B-4224-8CF6-2A8D4EB02382
var IID_ICodecSubtypesStatics = syscall.GUID{0xA66AC4F2, 0x888B, 0x4224,
	[8]byte{0x8C, 0xF6, 0x2A, 0x8D, 0x4E, 0xB0, 0x23, 0x82}}

type ICodecSubtypesStaticsInterface interface {
	win32.IInspectableInterface
	Get_VideoFormatDV25() string
	Get_VideoFormatDV50() string
	Get_VideoFormatDvc() string
	Get_VideoFormatDvh1() string
	Get_VideoFormatDvhD() string
	Get_VideoFormatDvsd() string
	Get_VideoFormatDvsl() string
	Get_VideoFormatH263() string
	Get_VideoFormatH264() string
	Get_VideoFormatH265() string
	Get_VideoFormatH264ES() string
	Get_VideoFormatHevc() string
	Get_VideoFormatHevcES() string
	Get_VideoFormatM4S2() string
	Get_VideoFormatMjpg() string
	Get_VideoFormatMP43() string
	Get_VideoFormatMP4S() string
	Get_VideoFormatMP4V() string
	Get_VideoFormatMpeg2() string
	Get_VideoFormatVP80() string
	Get_VideoFormatVP90() string
	Get_VideoFormatMpg1() string
	Get_VideoFormatMss1() string
	Get_VideoFormatMss2() string
	Get_VideoFormatWmv1() string
	Get_VideoFormatWmv2() string
	Get_VideoFormatWmv3() string
	Get_VideoFormatWvc1() string
	Get_VideoFormat420O() string
	Get_AudioFormatAac() string
	Get_AudioFormatAdts() string
	Get_AudioFormatAlac() string
	Get_AudioFormatAmrNB() string
	Get_AudioFormatAmrWB() string
	Get_AudioFormatAmrWP() string
	Get_AudioFormatDolbyAC3() string
	Get_AudioFormatDolbyAC3Spdif() string
	Get_AudioFormatDolbyDDPlus() string
	Get_AudioFormatDrm() string
	Get_AudioFormatDts() string
	Get_AudioFormatFlac() string
	Get_AudioFormatFloat() string
	Get_AudioFormatMP3() string
	Get_AudioFormatMPeg() string
	Get_AudioFormatMsp1() string
	Get_AudioFormatOpus() string
	Get_AudioFormatPcm() string
	Get_AudioFormatWmaSpdif() string
	Get_AudioFormatWMAudioLossless() string
	Get_AudioFormatWMAudioV8() string
	Get_AudioFormatWMAudioV9() string
}

type ICodecSubtypesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_VideoFormatDV25            uintptr
	Get_VideoFormatDV50            uintptr
	Get_VideoFormatDvc             uintptr
	Get_VideoFormatDvh1            uintptr
	Get_VideoFormatDvhD            uintptr
	Get_VideoFormatDvsd            uintptr
	Get_VideoFormatDvsl            uintptr
	Get_VideoFormatH263            uintptr
	Get_VideoFormatH264            uintptr
	Get_VideoFormatH265            uintptr
	Get_VideoFormatH264ES          uintptr
	Get_VideoFormatHevc            uintptr
	Get_VideoFormatHevcES          uintptr
	Get_VideoFormatM4S2            uintptr
	Get_VideoFormatMjpg            uintptr
	Get_VideoFormatMP43            uintptr
	Get_VideoFormatMP4S            uintptr
	Get_VideoFormatMP4V            uintptr
	Get_VideoFormatMpeg2           uintptr
	Get_VideoFormatVP80            uintptr
	Get_VideoFormatVP90            uintptr
	Get_VideoFormatMpg1            uintptr
	Get_VideoFormatMss1            uintptr
	Get_VideoFormatMss2            uintptr
	Get_VideoFormatWmv1            uintptr
	Get_VideoFormatWmv2            uintptr
	Get_VideoFormatWmv3            uintptr
	Get_VideoFormatWvc1            uintptr
	Get_VideoFormat420O            uintptr
	Get_AudioFormatAac             uintptr
	Get_AudioFormatAdts            uintptr
	Get_AudioFormatAlac            uintptr
	Get_AudioFormatAmrNB           uintptr
	Get_AudioFormatAmrWB           uintptr
	Get_AudioFormatAmrWP           uintptr
	Get_AudioFormatDolbyAC3        uintptr
	Get_AudioFormatDolbyAC3Spdif   uintptr
	Get_AudioFormatDolbyDDPlus     uintptr
	Get_AudioFormatDrm             uintptr
	Get_AudioFormatDts             uintptr
	Get_AudioFormatFlac            uintptr
	Get_AudioFormatFloat           uintptr
	Get_AudioFormatMP3             uintptr
	Get_AudioFormatMPeg            uintptr
	Get_AudioFormatMsp1            uintptr
	Get_AudioFormatOpus            uintptr
	Get_AudioFormatPcm             uintptr
	Get_AudioFormatWmaSpdif        uintptr
	Get_AudioFormatWMAudioLossless uintptr
	Get_AudioFormatWMAudioV8       uintptr
	Get_AudioFormatWMAudioV9       uintptr
}

type ICodecSubtypesStatics struct {
	win32.IInspectable
}

func (this *ICodecSubtypesStatics) Vtbl() *ICodecSubtypesStaticsVtbl {
	return (*ICodecSubtypesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICodecSubtypesStatics) Get_VideoFormatDV25() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormatDV25, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_VideoFormatDV50() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormatDV50, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_VideoFormatDvc() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormatDvc, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_VideoFormatDvh1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormatDvh1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_VideoFormatDvhD() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormatDvhD, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_VideoFormatDvsd() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormatDvsd, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_VideoFormatDvsl() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormatDvsl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_VideoFormatH263() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormatH263, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_VideoFormatH264() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormatH264, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_VideoFormatH265() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormatH265, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_VideoFormatH264ES() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormatH264ES, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_VideoFormatHevc() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormatHevc, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_VideoFormatHevcES() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormatHevcES, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_VideoFormatM4S2() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormatM4S2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_VideoFormatMjpg() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormatMjpg, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_VideoFormatMP43() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormatMP43, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_VideoFormatMP4S() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormatMP4S, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_VideoFormatMP4V() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormatMP4V, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_VideoFormatMpeg2() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormatMpeg2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_VideoFormatVP80() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormatVP80, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_VideoFormatVP90() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormatVP90, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_VideoFormatMpg1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormatMpg1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_VideoFormatMss1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormatMss1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_VideoFormatMss2() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormatMss2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_VideoFormatWmv1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormatWmv1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_VideoFormatWmv2() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormatWmv2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_VideoFormatWmv3() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormatWmv3, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_VideoFormatWvc1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormatWvc1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_VideoFormat420O() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormat420O, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_AudioFormatAac() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioFormatAac, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_AudioFormatAdts() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioFormatAdts, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_AudioFormatAlac() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioFormatAlac, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_AudioFormatAmrNB() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioFormatAmrNB, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_AudioFormatAmrWB() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioFormatAmrWB, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_AudioFormatAmrWP() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioFormatAmrWP, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_AudioFormatDolbyAC3() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioFormatDolbyAC3, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_AudioFormatDolbyAC3Spdif() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioFormatDolbyAC3Spdif, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_AudioFormatDolbyDDPlus() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioFormatDolbyDDPlus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_AudioFormatDrm() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioFormatDrm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_AudioFormatDts() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioFormatDts, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_AudioFormatFlac() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioFormatFlac, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_AudioFormatFloat() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioFormatFloat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_AudioFormatMP3() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioFormatMP3, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_AudioFormatMPeg() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioFormatMPeg, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_AudioFormatMsp1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioFormatMsp1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_AudioFormatOpus() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioFormatOpus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_AudioFormatPcm() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioFormatPcm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_AudioFormatWmaSpdif() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioFormatWmaSpdif, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_AudioFormatWMAudioLossless() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioFormatWMAudioLossless, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_AudioFormatWMAudioV8() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioFormatWMAudioV8, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICodecSubtypesStatics) Get_AudioFormatWMAudioV9() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioFormatWMAudioV9, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 7C7F676D-1FBC-4E2D-9A87-EE38BD1DC637
var IID_IDataCue = syscall.GUID{0x7C7F676D, 0x1FBC, 0x4E2D,
	[8]byte{0x9A, 0x87, 0xEE, 0x38, 0xBD, 0x1D, 0xC6, 0x37}}

type IDataCueInterface interface {
	win32.IInspectableInterface
	Put_Data(value *IBuffer)
	Get_Data() *IBuffer
}

type IDataCueVtbl struct {
	win32.IInspectableVtbl
	Put_Data uintptr
	Get_Data uintptr
}

type IDataCue struct {
	win32.IInspectable
}

func (this *IDataCue) Vtbl() *IDataCueVtbl {
	return (*IDataCueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataCue) Put_Data(value *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IDataCue) Get_Data() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BC561B15-95F2-49E8-96F1-8DD5DAC68D93
var IID_IDataCue2 = syscall.GUID{0xBC561B15, 0x95F2, 0x49E8,
	[8]byte{0x96, 0xF1, 0x8D, 0xD5, 0xDA, 0xC6, 0x8D, 0x93}}

type IDataCue2Interface interface {
	win32.IInspectableInterface
	Get_Properties() *IPropertySet
}

type IDataCue2Vtbl struct {
	win32.IInspectableVtbl
	Get_Properties uintptr
}

type IDataCue2 struct {
	win32.IInspectable
}

func (this *IDataCue2) Vtbl() *IDataCue2Vtbl {
	return (*IDataCue2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataCue2) Get_Properties() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 19918426-C65B-46BA-85F8-13880576C90A
var IID_IFaceDetectedEventArgs = syscall.GUID{0x19918426, 0xC65B, 0x46BA,
	[8]byte{0x85, 0xF8, 0x13, 0x88, 0x05, 0x76, 0xC9, 0x0A}}

type IFaceDetectedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ResultFrame() *IFaceDetectionEffectFrame
}

type IFaceDetectedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ResultFrame uintptr
}

type IFaceDetectedEventArgs struct {
	win32.IInspectable
}

func (this *IFaceDetectedEventArgs) Vtbl() *IFaceDetectedEventArgsVtbl {
	return (*IFaceDetectedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFaceDetectedEventArgs) Get_ResultFrame() *IFaceDetectionEffectFrame {
	var _result *IFaceDetectionEffectFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResultFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AE15EBD2-0542-42A9-BC90-F283A29F46C1
var IID_IFaceDetectionEffect = syscall.GUID{0xAE15EBD2, 0x0542, 0x42A9,
	[8]byte{0xBC, 0x90, 0xF2, 0x83, 0xA2, 0x9F, 0x46, 0xC1}}

type IFaceDetectionEffectInterface interface {
	win32.IInspectableInterface
	Put_Enabled(value bool)
	Get_Enabled() bool
	Put_DesiredDetectionInterval(value TimeSpan)
	Get_DesiredDetectionInterval() TimeSpan
	Add_FaceDetected(handler TypedEventHandler[*IFaceDetectionEffect, *IFaceDetectedEventArgs]) EventRegistrationToken
	Remove_FaceDetected(cookie EventRegistrationToken)
}

type IFaceDetectionEffectVtbl struct {
	win32.IInspectableVtbl
	Put_Enabled                  uintptr
	Get_Enabled                  uintptr
	Put_DesiredDetectionInterval uintptr
	Get_DesiredDetectionInterval uintptr
	Add_FaceDetected             uintptr
	Remove_FaceDetected          uintptr
}

type IFaceDetectionEffect struct {
	win32.IInspectable
}

func (this *IFaceDetectionEffect) Vtbl() *IFaceDetectionEffectVtbl {
	return (*IFaceDetectionEffectVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFaceDetectionEffect) Put_Enabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Enabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IFaceDetectionEffect) Get_Enabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Enabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFaceDetectionEffect) Put_DesiredDetectionInterval(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DesiredDetectionInterval, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IFaceDetectionEffect) Get_DesiredDetectionInterval() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesiredDetectionInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFaceDetectionEffect) Add_FaceDetected(handler TypedEventHandler[*IFaceDetectionEffect, *IFaceDetectedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_FaceDetected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFaceDetectionEffect) Remove_FaceDetected(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_FaceDetected, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

// 43DCA081-B848-4F33-B702-1FD2624FB016
var IID_IFaceDetectionEffectDefinition = syscall.GUID{0x43DCA081, 0xB848, 0x4F33,
	[8]byte{0xB7, 0x02, 0x1F, 0xD2, 0x62, 0x4F, 0xB0, 0x16}}

type IFaceDetectionEffectDefinitionInterface interface {
	win32.IInspectableInterface
	Put_DetectionMode(value FaceDetectionMode)
	Get_DetectionMode() FaceDetectionMode
	Put_SynchronousDetectionEnabled(value bool)
	Get_SynchronousDetectionEnabled() bool
}

type IFaceDetectionEffectDefinitionVtbl struct {
	win32.IInspectableVtbl
	Put_DetectionMode               uintptr
	Get_DetectionMode               uintptr
	Put_SynchronousDetectionEnabled uintptr
	Get_SynchronousDetectionEnabled uintptr
}

type IFaceDetectionEffectDefinition struct {
	win32.IInspectable
}

func (this *IFaceDetectionEffectDefinition) Vtbl() *IFaceDetectionEffectDefinitionVtbl {
	return (*IFaceDetectionEffectDefinitionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFaceDetectionEffectDefinition) Put_DetectionMode(value FaceDetectionMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DetectionMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IFaceDetectionEffectDefinition) Get_DetectionMode() FaceDetectionMode {
	var _result FaceDetectionMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DetectionMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFaceDetectionEffectDefinition) Put_SynchronousDetectionEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SynchronousDetectionEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IFaceDetectionEffectDefinition) Get_SynchronousDetectionEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SynchronousDetectionEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 8AB08993-5DC8-447B-A247-5270BD802ECE
var IID_IFaceDetectionEffectFrame = syscall.GUID{0x8AB08993, 0x5DC8, 0x447B,
	[8]byte{0xA2, 0x47, 0x52, 0x70, 0xBD, 0x80, 0x2E, 0xCE}}

type IFaceDetectionEffectFrameInterface interface {
	win32.IInspectableInterface
	Get_DetectedFaces() *IVectorView[unsafe.Pointer]
}

type IFaceDetectionEffectFrameVtbl struct {
	win32.IInspectableVtbl
	Get_DetectedFaces uintptr
}

type IFaceDetectionEffectFrame struct {
	win32.IInspectable
}

func (this *IFaceDetectionEffectFrame) Vtbl() *IFaceDetectionEffectFrameVtbl {
	return (*IFaceDetectionEffectFrameVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFaceDetectionEffectFrame) Get_DetectedFaces() *IVectorView[unsafe.Pointer] {
	var _result *IVectorView[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DetectedFaces, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 55F1A7AE-D957-4DC9-9D1C-8553A82A7D99
var IID_IHighDynamicRangeControl = syscall.GUID{0x55F1A7AE, 0xD957, 0x4DC9,
	[8]byte{0x9D, 0x1C, 0x85, 0x53, 0xA8, 0x2A, 0x7D, 0x99}}

type IHighDynamicRangeControlInterface interface {
	win32.IInspectableInterface
	Put_Enabled(value bool)
	Get_Enabled() bool
}

type IHighDynamicRangeControlVtbl struct {
	win32.IInspectableVtbl
	Put_Enabled uintptr
	Get_Enabled uintptr
}

type IHighDynamicRangeControl struct {
	win32.IInspectable
}

func (this *IHighDynamicRangeControl) Vtbl() *IHighDynamicRangeControlVtbl {
	return (*IHighDynamicRangeControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHighDynamicRangeControl) Put_Enabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Enabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IHighDynamicRangeControl) Get_Enabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Enabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 0F57806B-253B-4119-BB40-3A90E51384F7
var IID_IHighDynamicRangeOutput = syscall.GUID{0x0F57806B, 0x253B, 0x4119,
	[8]byte{0xBB, 0x40, 0x3A, 0x90, 0xE5, 0x13, 0x84, 0xF7}}

type IHighDynamicRangeOutputInterface interface {
	win32.IInspectableInterface
	Get_Certainty() float64
	Get_FrameControllers() *IVectorView[*IFrameController]
}

type IHighDynamicRangeOutputVtbl struct {
	win32.IInspectableVtbl
	Get_Certainty        uintptr
	Get_FrameControllers uintptr
}

type IHighDynamicRangeOutput struct {
	win32.IInspectable
}

func (this *IHighDynamicRangeOutput) Vtbl() *IHighDynamicRangeOutputVtbl {
	return (*IHighDynamicRangeOutputVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHighDynamicRangeOutput) Get_Certainty() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Certainty, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHighDynamicRangeOutput) Get_FrameControllers() *IVectorView[*IFrameController] {
	var _result *IVectorView[*IFrameController]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameControllers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 52828282-367B-440B-9116-3C84570DD270
var IID_IImageCue = syscall.GUID{0x52828282, 0x367B, 0x440B,
	[8]byte{0x91, 0x16, 0x3C, 0x84, 0x57, 0x0D, 0xD2, 0x70}}

type IImageCueInterface interface {
	win32.IInspectableInterface
	Get_Position() TimedTextPoint
	Put_Position(value TimedTextPoint)
	Get_Extent() TimedTextSize
	Put_Extent(value TimedTextSize)
	Put_SoftwareBitmap(value *ISoftwareBitmap)
	Get_SoftwareBitmap() *ISoftwareBitmap
}

type IImageCueVtbl struct {
	win32.IInspectableVtbl
	Get_Position       uintptr
	Put_Position       uintptr
	Get_Extent         uintptr
	Put_Extent         uintptr
	Put_SoftwareBitmap uintptr
	Get_SoftwareBitmap uintptr
}

type IImageCue struct {
	win32.IInspectable
}

func (this *IImageCue) Vtbl() *IImageCueVtbl {
	return (*IImageCueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IImageCue) Get_Position() TimedTextPoint {
	var _result TimedTextPoint
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageCue) Put_Position(value TimedTextPoint) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IImageCue) Get_Extent() TimedTextSize {
	var _result TimedTextSize
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Extent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageCue) Put_Extent(value TimedTextSize) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Extent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IImageCue) Put_SoftwareBitmap(value *ISoftwareBitmap) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SoftwareBitmap, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IImageCue) Get_SoftwareBitmap() *ISoftwareBitmap {
	var _result *ISoftwareBitmap
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SoftwareBitmap, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 25BC45E1-9B08-4C2E-A855-4542F1A75DEB
var IID_IInitializeMediaStreamSourceRequestedEventArgs = syscall.GUID{0x25BC45E1, 0x9B08, 0x4C2E,
	[8]byte{0xA8, 0x55, 0x45, 0x42, 0xF1, 0xA7, 0x5D, 0xEB}}

type IInitializeMediaStreamSourceRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Source() *IMediaStreamSource
	Get_RandomAccessStream() *IRandomAccessStream
	GetDeferral() *IDeferral
}

type IInitializeMediaStreamSourceRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Source             uintptr
	Get_RandomAccessStream uintptr
	GetDeferral            uintptr
}

type IInitializeMediaStreamSourceRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IInitializeMediaStreamSourceRequestedEventArgs) Vtbl() *IInitializeMediaStreamSourceRequestedEventArgsVtbl {
	return (*IInitializeMediaStreamSourceRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInitializeMediaStreamSourceRequestedEventArgs) Get_Source() *IMediaStreamSource {
	var _result *IMediaStreamSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Source, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInitializeMediaStreamSourceRequestedEventArgs) Get_RandomAccessStream() *IRandomAccessStream {
	var _result *IRandomAccessStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RandomAccessStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInitializeMediaStreamSourceRequestedEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 78EDBE35-27A0-42E0-9CD3-738D2089DE9C
var IID_ILowLightFusionResult = syscall.GUID{0x78EDBE35, 0x27A0, 0x42E0,
	[8]byte{0x9C, 0xD3, 0x73, 0x8D, 0x20, 0x89, 0xDE, 0x9C}}

type ILowLightFusionResultInterface interface {
	win32.IInspectableInterface
	Get_Frame() *ISoftwareBitmap
}

type ILowLightFusionResultVtbl struct {
	win32.IInspectableVtbl
	Get_Frame uintptr
}

type ILowLightFusionResult struct {
	win32.IInspectable
}

func (this *ILowLightFusionResult) Vtbl() *ILowLightFusionResultVtbl {
	return (*ILowLightFusionResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILowLightFusionResult) Get_Frame() *ISoftwareBitmap {
	var _result *ISoftwareBitmap
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Frame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5305016D-C29E-40E2-87A9-9E1FD2F192F5
var IID_ILowLightFusionStatics = syscall.GUID{0x5305016D, 0xC29E, 0x40E2,
	[8]byte{0x87, 0xA9, 0x9E, 0x1F, 0xD2, 0xF1, 0x92, 0xF5}}

type ILowLightFusionStaticsInterface interface {
	win32.IInspectableInterface
	Get_SupportedBitmapPixelFormats() *IVectorView[BitmapPixelFormat]
	Get_MaxSupportedFrameCount() int32
	FuseAsync(frameSet *IIterable[*ISoftwareBitmap]) *IAsyncOperationWithProgress[*ILowLightFusionResult, float64]
}

type ILowLightFusionStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_SupportedBitmapPixelFormats uintptr
	Get_MaxSupportedFrameCount      uintptr
	FuseAsync                       uintptr
}

type ILowLightFusionStatics struct {
	win32.IInspectable
}

func (this *ILowLightFusionStatics) Vtbl() *ILowLightFusionStaticsVtbl {
	return (*ILowLightFusionStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILowLightFusionStatics) Get_SupportedBitmapPixelFormats() *IVectorView[BitmapPixelFormat] {
	var _result *IVectorView[BitmapPixelFormat]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedBitmapPixelFormats, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILowLightFusionStatics) Get_MaxSupportedFrameCount() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxSupportedFrameCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILowLightFusionStatics) FuseAsync(frameSet *IIterable[*ISoftwareBitmap]) *IAsyncOperationWithProgress[*ILowLightFusionResult, float64] {
	var _result *IAsyncOperationWithProgress[*ILowLightFusionResult, float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FuseAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(frameSet)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2B7E40AA-DE07-424F-83F1-F1DE46C4FA2E
var IID_IMediaBinder = syscall.GUID{0x2B7E40AA, 0xDE07, 0x424F,
	[8]byte{0x83, 0xF1, 0xF1, 0xDE, 0x46, 0xC4, 0xFA, 0x2E}}

type IMediaBinderInterface interface {
	win32.IInspectableInterface
	Add_Binding(handler TypedEventHandler[*IMediaBinder, *IMediaBindingEventArgs]) EventRegistrationToken
	Remove_Binding(token EventRegistrationToken)
	Get_Token() string
	Put_Token(value string)
	Get_Source() *IMediaSource2
}

type IMediaBinderVtbl struct {
	win32.IInspectableVtbl
	Add_Binding    uintptr
	Remove_Binding uintptr
	Get_Token      uintptr
	Put_Token      uintptr
	Get_Source     uintptr
}

type IMediaBinder struct {
	win32.IInspectable
}

func (this *IMediaBinder) Vtbl() *IMediaBinderVtbl {
	return (*IMediaBinderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaBinder) Add_Binding(handler TypedEventHandler[*IMediaBinder, *IMediaBindingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Binding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaBinder) Remove_Binding(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Binding, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaBinder) Get_Token() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Token, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaBinder) Put_Token(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Token, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMediaBinder) Get_Source() *IMediaSource2 {
	var _result *IMediaSource2
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Source, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B61CB25A-1B6D-4630-A86D-2F0837F712E5
var IID_IMediaBindingEventArgs = syscall.GUID{0xB61CB25A, 0x1B6D, 0x4630,
	[8]byte{0xA8, 0x6D, 0x2F, 0x08, 0x37, 0xF7, 0x12, 0xE5}}

type IMediaBindingEventArgsInterface interface {
	win32.IInspectableInterface
	Add_Canceled(handler TypedEventHandler[*IMediaBindingEventArgs, interface{}]) EventRegistrationToken
	Remove_Canceled(token EventRegistrationToken)
	Get_MediaBinder() *IMediaBinder
	GetDeferral() *IDeferral
	SetUri(uri *IUriRuntimeClass)
	SetStream(stream *IRandomAccessStream, contentType string)
	SetStreamReference(stream *IRandomAccessStreamReference, contentType string)
}

type IMediaBindingEventArgsVtbl struct {
	win32.IInspectableVtbl
	Add_Canceled       uintptr
	Remove_Canceled    uintptr
	Get_MediaBinder    uintptr
	GetDeferral        uintptr
	SetUri             uintptr
	SetStream          uintptr
	SetStreamReference uintptr
}

type IMediaBindingEventArgs struct {
	win32.IInspectable
}

func (this *IMediaBindingEventArgs) Vtbl() *IMediaBindingEventArgsVtbl {
	return (*IMediaBindingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaBindingEventArgs) Add_Canceled(handler TypedEventHandler[*IMediaBindingEventArgs, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Canceled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaBindingEventArgs) Remove_Canceled(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Canceled, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaBindingEventArgs) Get_MediaBinder() *IMediaBinder {
	var _result *IMediaBinder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaBinder, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaBindingEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaBindingEventArgs) SetUri(uri *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)))
	_ = _hr
}

func (this *IMediaBindingEventArgs) SetStream(stream *IRandomAccessStream, contentType string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)), NewHStr(contentType).Ptr)
	_ = _hr
}

func (this *IMediaBindingEventArgs) SetStreamReference(stream *IRandomAccessStreamReference, contentType string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetStreamReference, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)), NewHStr(contentType).Ptr)
	_ = _hr
}

// 0464CCEB-BB5A-482F-B8BA-F0284C696567
var IID_IMediaBindingEventArgs2 = syscall.GUID{0x0464CCEB, 0xBB5A, 0x482F,
	[8]byte{0xB8, 0xBA, 0xF0, 0x28, 0x4C, 0x69, 0x65, 0x67}}

type IMediaBindingEventArgs2Interface interface {
	win32.IInspectableInterface
	SetAdaptiveMediaSource(mediaSource *IAdaptiveMediaSource)
	SetStorageFile(file *IStorageFile)
}

type IMediaBindingEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	SetAdaptiveMediaSource uintptr
	SetStorageFile         uintptr
}

type IMediaBindingEventArgs2 struct {
	win32.IInspectable
}

func (this *IMediaBindingEventArgs2) Vtbl() *IMediaBindingEventArgs2Vtbl {
	return (*IMediaBindingEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaBindingEventArgs2) SetAdaptiveMediaSource(mediaSource *IAdaptiveMediaSource) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetAdaptiveMediaSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(mediaSource)))
	_ = _hr
}

func (this *IMediaBindingEventArgs2) SetStorageFile(file *IStorageFile) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetStorageFile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)))
	_ = _hr
}

// F8EB475E-19BE-44FC-A5ED-7ABA315037F9
var IID_IMediaBindingEventArgs3 = syscall.GUID{0xF8EB475E, 0x19BE, 0x44FC,
	[8]byte{0xA5, 0xED, 0x7A, 0xBA, 0x31, 0x50, 0x37, 0xF9}}

type IMediaBindingEventArgs3Interface interface {
	win32.IInspectableInterface
	SetDownloadOperation(downloadOperation *IDownloadOperation)
}

type IMediaBindingEventArgs3Vtbl struct {
	win32.IInspectableVtbl
	SetDownloadOperation uintptr
}

type IMediaBindingEventArgs3 struct {
	win32.IInspectable
}

func (this *IMediaBindingEventArgs3) Vtbl() *IMediaBindingEventArgs3Vtbl {
	return (*IMediaBindingEventArgs3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaBindingEventArgs3) SetDownloadOperation(downloadOperation *IDownloadOperation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetDownloadOperation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(downloadOperation)))
	_ = _hr
}

// C7D15E5D-59DC-431F-A0EE-27744323B36D
var IID_IMediaCue = syscall.GUID{0xC7D15E5D, 0x59DC, 0x431F,
	[8]byte{0xA0, 0xEE, 0x27, 0x74, 0x43, 0x23, 0xB3, 0x6D}}

type IMediaCueInterface interface {
	win32.IInspectableInterface
	Put_StartTime(value TimeSpan)
	Get_StartTime() TimeSpan
	Put_Duration(value TimeSpan)
	Get_Duration() TimeSpan
	Put_Id(value string)
	Get_Id() string
}

type IMediaCueVtbl struct {
	win32.IInspectableVtbl
	Put_StartTime uintptr
	Get_StartTime uintptr
	Put_Duration  uintptr
	Get_Duration  uintptr
	Put_Id        uintptr
	Get_Id        uintptr
}

type IMediaCue struct {
	win32.IInspectable
}

func (this *IMediaCue) Vtbl() *IMediaCueVtbl {
	return (*IMediaCueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCue) Put_StartTime(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StartTime, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IMediaCue) Get_StartTime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StartTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCue) Put_Duration(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Duration, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IMediaCue) Get_Duration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Duration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCue) Put_Id(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Id, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMediaCue) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// D12F47F7-5FA4-4E68-9FE5-32160DCEE57E
var IID_IMediaCueEventArgs = syscall.GUID{0xD12F47F7, 0x5FA4, 0x4E68,
	[8]byte{0x9F, 0xE5, 0x32, 0x16, 0x0D, 0xCE, 0xE5, 0x7E}}

type IMediaCueEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Cue() *IMediaCue
}

type IMediaCueEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Cue uintptr
}

type IMediaCueEventArgs struct {
	win32.IInspectable
}

func (this *IMediaCueEventArgs) Vtbl() *IMediaCueEventArgsVtbl {
	return (*IMediaCueEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCueEventArgs) Get_Cue() *IMediaCue {
	var _result *IMediaCue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Cue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E7BFB599-A09D-4C21-BCDF-20AF4F86B3D9
var IID_IMediaSource = syscall.GUID{0xE7BFB599, 0xA09D, 0x4C21,
	[8]byte{0xBC, 0xDF, 0x20, 0xAF, 0x4F, 0x86, 0xB3, 0xD9}}

type IMediaSourceInterface interface {
	win32.IInspectableInterface
}

type IMediaSourceVtbl struct {
	win32.IInspectableVtbl
}

type IMediaSource struct {
	win32.IInspectable
}

func (this *IMediaSource) Vtbl() *IMediaSourceVtbl {
	return (*IMediaSourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 2EB61048-655F-4C37-B813-B4E45DFA0ABE
var IID_IMediaSource2 = syscall.GUID{0x2EB61048, 0x655F, 0x4C37,
	[8]byte{0xB8, 0x13, 0xB4, 0xE4, 0x5D, 0xFA, 0x0A, 0xBE}}

type IMediaSource2Interface interface {
	win32.IInspectableInterface
	Add_OpenOperationCompleted(handler TypedEventHandler[*IMediaSource2, *IMediaSourceOpenOperationCompletedEventArgs]) EventRegistrationToken
	Remove_OpenOperationCompleted(token EventRegistrationToken)
	Get_CustomProperties() *IPropertySet
	Get_Duration() *IReference[TimeSpan]
	Get_IsOpen() bool
	Get_ExternalTimedTextSources() *IObservableVector[*ITimedTextSource]
	Get_ExternalTimedMetadataTracks() *IObservableVector[*ITimedMetadataTrack]
}

type IMediaSource2Vtbl struct {
	win32.IInspectableVtbl
	Add_OpenOperationCompleted      uintptr
	Remove_OpenOperationCompleted   uintptr
	Get_CustomProperties            uintptr
	Get_Duration                    uintptr
	Get_IsOpen                      uintptr
	Get_ExternalTimedTextSources    uintptr
	Get_ExternalTimedMetadataTracks uintptr
}

type IMediaSource2 struct {
	win32.IInspectable
}

func (this *IMediaSource2) Vtbl() *IMediaSource2Vtbl {
	return (*IMediaSource2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaSource2) Add_OpenOperationCompleted(handler TypedEventHandler[*IMediaSource2, *IMediaSourceOpenOperationCompletedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_OpenOperationCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaSource2) Remove_OpenOperationCompleted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_OpenOperationCompleted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaSource2) Get_CustomProperties() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CustomProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaSource2) Get_Duration() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Duration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaSource2) Get_IsOpen() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsOpen, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaSource2) Get_ExternalTimedTextSources() *IObservableVector[*ITimedTextSource] {
	var _result *IObservableVector[*ITimedTextSource]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExternalTimedTextSources, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaSource2) Get_ExternalTimedMetadataTracks() *IObservableVector[*ITimedMetadataTrack] {
	var _result *IObservableVector[*ITimedMetadataTrack]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExternalTimedMetadataTracks, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B59F0D9B-4B6E-41ED-BBB4-7C7509A994AD
var IID_IMediaSource3 = syscall.GUID{0xB59F0D9B, 0x4B6E, 0x41ED,
	[8]byte{0xBB, 0xB4, 0x7C, 0x75, 0x09, 0xA9, 0x94, 0xAD}}

type IMediaSource3Interface interface {
	win32.IInspectableInterface
	Add_StateChanged(handler TypedEventHandler[*IMediaSource2, *IMediaSourceStateChangedEventArgs]) EventRegistrationToken
	Remove_StateChanged(token EventRegistrationToken)
	Get_State() MediaSourceState
	Reset()
}

type IMediaSource3Vtbl struct {
	win32.IInspectableVtbl
	Add_StateChanged    uintptr
	Remove_StateChanged uintptr
	Get_State           uintptr
	Reset               uintptr
}

type IMediaSource3 struct {
	win32.IInspectable
}

func (this *IMediaSource3) Vtbl() *IMediaSource3Vtbl {
	return (*IMediaSource3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaSource3) Add_StateChanged(handler TypedEventHandler[*IMediaSource2, *IMediaSourceStateChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaSource3) Remove_StateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaSource3) Get_State() MediaSourceState {
	var _result MediaSourceState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaSource3) Reset() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Reset, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// BDAFAD57-8EFF-4C63-85A6-84DE0AE3E4F2
var IID_IMediaSource4 = syscall.GUID{0xBDAFAD57, 0x8EFF, 0x4C63,
	[8]byte{0x85, 0xA6, 0x84, 0xDE, 0x0A, 0xE3, 0xE4, 0xF2}}

type IMediaSource4Interface interface {
	win32.IInspectableInterface
	Get_AdaptiveMediaSource() *IAdaptiveMediaSource
	Get_MediaStreamSource() *IMediaStreamSource
	Get_MseStreamSource() *IMseStreamSource
	Get_Uri() *IUriRuntimeClass
	OpenAsync() *IAsyncAction
}

type IMediaSource4Vtbl struct {
	win32.IInspectableVtbl
	Get_AdaptiveMediaSource uintptr
	Get_MediaStreamSource   uintptr
	Get_MseStreamSource     uintptr
	Get_Uri                 uintptr
	OpenAsync               uintptr
}

type IMediaSource4 struct {
	win32.IInspectable
}

func (this *IMediaSource4) Vtbl() *IMediaSource4Vtbl {
	return (*IMediaSource4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaSource4) Get_AdaptiveMediaSource() *IAdaptiveMediaSource {
	var _result *IAdaptiveMediaSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AdaptiveMediaSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaSource4) Get_MediaStreamSource() *IMediaStreamSource {
	var _result *IMediaStreamSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaStreamSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaSource4) Get_MseStreamSource() *IMseStreamSource {
	var _result *IMseStreamSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MseStreamSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaSource4) Get_Uri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Uri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaSource4) OpenAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 331A22AE-ED2E-4A22-94C8-B743A92B3022
var IID_IMediaSource5 = syscall.GUID{0x331A22AE, 0xED2E, 0x4A22,
	[8]byte{0x94, 0xC8, 0xB7, 0x43, 0xA9, 0x2B, 0x30, 0x22}}

type IMediaSource5Interface interface {
	win32.IInspectableInterface
	Get_DownloadOperation() *IDownloadOperation
}

type IMediaSource5Vtbl struct {
	win32.IInspectableVtbl
	Get_DownloadOperation uintptr
}

type IMediaSource5 struct {
	win32.IInspectable
}

func (this *IMediaSource5) Vtbl() *IMediaSource5Vtbl {
	return (*IMediaSource5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaSource5) Get_DownloadOperation() *IDownloadOperation {
	var _result *IDownloadOperation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DownloadOperation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 61E1EA97-1916-4810-B7F4-B642BE829596
var IID_IMediaSourceAppServiceConnection = syscall.GUID{0x61E1EA97, 0x1916, 0x4810,
	[8]byte{0xB7, 0xF4, 0xB6, 0x42, 0xBE, 0x82, 0x95, 0x96}}

type IMediaSourceAppServiceConnectionInterface interface {
	win32.IInspectableInterface
	Add_InitializeMediaStreamSourceRequested(handler TypedEventHandler[*IMediaSourceAppServiceConnection, *IInitializeMediaStreamSourceRequestedEventArgs]) EventRegistrationToken
	Remove_InitializeMediaStreamSourceRequested(token EventRegistrationToken)
	Start()
}

type IMediaSourceAppServiceConnectionVtbl struct {
	win32.IInspectableVtbl
	Add_InitializeMediaStreamSourceRequested    uintptr
	Remove_InitializeMediaStreamSourceRequested uintptr
	Start                                       uintptr
}

type IMediaSourceAppServiceConnection struct {
	win32.IInspectable
}

func (this *IMediaSourceAppServiceConnection) Vtbl() *IMediaSourceAppServiceConnectionVtbl {
	return (*IMediaSourceAppServiceConnectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaSourceAppServiceConnection) Add_InitializeMediaStreamSourceRequested(handler TypedEventHandler[*IMediaSourceAppServiceConnection, *IInitializeMediaStreamSourceRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_InitializeMediaStreamSourceRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaSourceAppServiceConnection) Remove_InitializeMediaStreamSourceRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_InitializeMediaStreamSourceRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaSourceAppServiceConnection) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 65B912EB-80B9-44F9-9C1E-E120F6D92838
var IID_IMediaSourceAppServiceConnectionFactory = syscall.GUID{0x65B912EB, 0x80B9, 0x44F9,
	[8]byte{0x9C, 0x1E, 0xE1, 0x20, 0xF6, 0xD9, 0x28, 0x38}}

type IMediaSourceAppServiceConnectionFactoryInterface interface {
	win32.IInspectableInterface
	Create(appServiceConnection *IAppServiceConnection) *IMediaSourceAppServiceConnection
}

type IMediaSourceAppServiceConnectionFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IMediaSourceAppServiceConnectionFactory struct {
	win32.IInspectable
}

func (this *IMediaSourceAppServiceConnectionFactory) Vtbl() *IMediaSourceAppServiceConnectionFactoryVtbl {
	return (*IMediaSourceAppServiceConnectionFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaSourceAppServiceConnectionFactory) Create(appServiceConnection *IAppServiceConnection) *IMediaSourceAppServiceConnection {
	var _result *IMediaSourceAppServiceConnection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(appServiceConnection)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5C0A8965-37C5-4E9D-8D21-1CDEE90CECC6
var IID_IMediaSourceError = syscall.GUID{0x5C0A8965, 0x37C5, 0x4E9D,
	[8]byte{0x8D, 0x21, 0x1C, 0xDE, 0xE9, 0x0C, 0xEC, 0xC6}}

type IMediaSourceErrorInterface interface {
	win32.IInspectableInterface
	Get_ExtendedError() HResult
}

type IMediaSourceErrorVtbl struct {
	win32.IInspectableVtbl
	Get_ExtendedError uintptr
}

type IMediaSourceError struct {
	win32.IInspectable
}

func (this *IMediaSourceError) Vtbl() *IMediaSourceErrorVtbl {
	return (*IMediaSourceErrorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaSourceError) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// FC682CEB-E281-477C-A8E0-1ACD654114C8
var IID_IMediaSourceOpenOperationCompletedEventArgs = syscall.GUID{0xFC682CEB, 0xE281, 0x477C,
	[8]byte{0xA8, 0xE0, 0x1A, 0xCD, 0x65, 0x41, 0x14, 0xC8}}

type IMediaSourceOpenOperationCompletedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Error() *IMediaSourceError
}

type IMediaSourceOpenOperationCompletedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Error uintptr
}

type IMediaSourceOpenOperationCompletedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaSourceOpenOperationCompletedEventArgs) Vtbl() *IMediaSourceOpenOperationCompletedEventArgsVtbl {
	return (*IMediaSourceOpenOperationCompletedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaSourceOpenOperationCompletedEventArgs) Get_Error() *IMediaSourceError {
	var _result *IMediaSourceError
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Error, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0A30AF82-9071-4BAC-BC39-CA2A93B717A9
var IID_IMediaSourceStateChangedEventArgs = syscall.GUID{0x0A30AF82, 0x9071, 0x4BAC,
	[8]byte{0xBC, 0x39, 0xCA, 0x2A, 0x93, 0xB7, 0x17, 0xA9}}

type IMediaSourceStateChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_OldState() MediaSourceState
	Get_NewState() MediaSourceState
}

type IMediaSourceStateChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_OldState uintptr
	Get_NewState uintptr
}

type IMediaSourceStateChangedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaSourceStateChangedEventArgs) Vtbl() *IMediaSourceStateChangedEventArgsVtbl {
	return (*IMediaSourceStateChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaSourceStateChangedEventArgs) Get_OldState() MediaSourceState {
	var _result MediaSourceState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OldState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaSourceStateChangedEventArgs) Get_NewState() MediaSourceState {
	var _result MediaSourceState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NewState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F77D6FA4-4652-410E-B1D8-E9A5E245A45C
var IID_IMediaSourceStatics = syscall.GUID{0xF77D6FA4, 0x4652, 0x410E,
	[8]byte{0xB1, 0xD8, 0xE9, 0xA5, 0xE2, 0x45, 0xA4, 0x5C}}

type IMediaSourceStaticsInterface interface {
	win32.IInspectableInterface
	CreateFromAdaptiveMediaSource(mediaSource *IAdaptiveMediaSource) *IMediaSource2
	CreateFromMediaStreamSource(mediaSource *IMediaStreamSource) *IMediaSource2
	CreateFromMseStreamSource(mediaSource *IMseStreamSource) *IMediaSource2
	CreateFromIMediaSource(mediaSource *IMediaSource) *IMediaSource2
	CreateFromStorageFile(file *IStorageFile) *IMediaSource2
	CreateFromStream(stream *IRandomAccessStream, contentType string) *IMediaSource2
	CreateFromStreamReference(stream *IRandomAccessStreamReference, contentType string) *IMediaSource2
	CreateFromUri(uri *IUriRuntimeClass) *IMediaSource2
}

type IMediaSourceStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateFromAdaptiveMediaSource uintptr
	CreateFromMediaStreamSource   uintptr
	CreateFromMseStreamSource     uintptr
	CreateFromIMediaSource        uintptr
	CreateFromStorageFile         uintptr
	CreateFromStream              uintptr
	CreateFromStreamReference     uintptr
	CreateFromUri                 uintptr
}

type IMediaSourceStatics struct {
	win32.IInspectable
}

func (this *IMediaSourceStatics) Vtbl() *IMediaSourceStaticsVtbl {
	return (*IMediaSourceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaSourceStatics) CreateFromAdaptiveMediaSource(mediaSource *IAdaptiveMediaSource) *IMediaSource2 {
	var _result *IMediaSource2
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromAdaptiveMediaSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(mediaSource)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaSourceStatics) CreateFromMediaStreamSource(mediaSource *IMediaStreamSource) *IMediaSource2 {
	var _result *IMediaSource2
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromMediaStreamSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(mediaSource)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaSourceStatics) CreateFromMseStreamSource(mediaSource *IMseStreamSource) *IMediaSource2 {
	var _result *IMediaSource2
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromMseStreamSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(mediaSource)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaSourceStatics) CreateFromIMediaSource(mediaSource *IMediaSource) *IMediaSource2 {
	var _result *IMediaSource2
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromIMediaSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(mediaSource)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaSourceStatics) CreateFromStorageFile(file *IStorageFile) *IMediaSource2 {
	var _result *IMediaSource2
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromStorageFile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaSourceStatics) CreateFromStream(stream *IRandomAccessStream, contentType string) *IMediaSource2 {
	var _result *IMediaSource2
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)), NewHStr(contentType).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaSourceStatics) CreateFromStreamReference(stream *IRandomAccessStreamReference, contentType string) *IMediaSource2 {
	var _result *IMediaSource2
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromStreamReference, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)), NewHStr(contentType).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaSourceStatics) CreateFromUri(uri *IUriRuntimeClass) *IMediaSource2 {
	var _result *IMediaSource2
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EEE161A4-7F13-4896-B8CB-DF0DE5BCB9F1
var IID_IMediaSourceStatics2 = syscall.GUID{0xEEE161A4, 0x7F13, 0x4896,
	[8]byte{0xB8, 0xCB, 0xDF, 0x0D, 0xE5, 0xBC, 0xB9, 0xF1}}

type IMediaSourceStatics2Interface interface {
	win32.IInspectableInterface
	CreateFromMediaBinder(binder *IMediaBinder) *IMediaSource2
}

type IMediaSourceStatics2Vtbl struct {
	win32.IInspectableVtbl
	CreateFromMediaBinder uintptr
}

type IMediaSourceStatics2 struct {
	win32.IInspectable
}

func (this *IMediaSourceStatics2) Vtbl() *IMediaSourceStatics2Vtbl {
	return (*IMediaSourceStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaSourceStatics2) CreateFromMediaBinder(binder *IMediaBinder) *IMediaSource2 {
	var _result *IMediaSource2
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromMediaBinder, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(binder)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 453A30D6-2BEA-4122-9F73-EACE04526E35
var IID_IMediaSourceStatics3 = syscall.GUID{0x453A30D6, 0x2BEA, 0x4122,
	[8]byte{0x9F, 0x73, 0xEA, 0xCE, 0x04, 0x52, 0x6E, 0x35}}

type IMediaSourceStatics3Interface interface {
	win32.IInspectableInterface
	CreateFromMediaFrameSource(frameSource *IMediaFrameSource) *IMediaSource2
}

type IMediaSourceStatics3Vtbl struct {
	win32.IInspectableVtbl
	CreateFromMediaFrameSource uintptr
}

type IMediaSourceStatics3 struct {
	win32.IInspectable
}

func (this *IMediaSourceStatics3) Vtbl() *IMediaSourceStatics3Vtbl {
	return (*IMediaSourceStatics3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaSourceStatics3) CreateFromMediaFrameSource(frameSource *IMediaFrameSource) *IMediaSource2 {
	var _result *IMediaSource2
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromMediaFrameSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(frameSource)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 281B3BFC-E50A-4428-A500-9C4ED918D3F0
var IID_IMediaSourceStatics4 = syscall.GUID{0x281B3BFC, 0xE50A, 0x4428,
	[8]byte{0xA5, 0x00, 0x9C, 0x4E, 0xD9, 0x18, 0xD3, 0xF0}}

type IMediaSourceStatics4Interface interface {
	win32.IInspectableInterface
	CreateFromDownloadOperation(downloadOperation *IDownloadOperation) *IMediaSource2
}

type IMediaSourceStatics4Vtbl struct {
	win32.IInspectableVtbl
	CreateFromDownloadOperation uintptr
}

type IMediaSourceStatics4 struct {
	win32.IInspectable
}

func (this *IMediaSourceStatics4) Vtbl() *IMediaSourceStatics4Vtbl {
	return (*IMediaSourceStatics4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaSourceStatics4) CreateFromDownloadOperation(downloadOperation *IDownloadOperation) *IMediaSource2 {
	var _result *IMediaSource2
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromDownloadOperation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(downloadOperation)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 80F16E6E-92F7-451E-97D2-AFD80742DA70
var IID_IMediaStreamDescriptor = syscall.GUID{0x80F16E6E, 0x92F7, 0x451E,
	[8]byte{0x97, 0xD2, 0xAF, 0xD8, 0x07, 0x42, 0xDA, 0x70}}

type IMediaStreamDescriptorInterface interface {
	win32.IInspectableInterface
	Get_IsSelected() bool
	Put_Name(value string)
	Get_Name() string
	Put_Language(value string)
	Get_Language() string
}

type IMediaStreamDescriptorVtbl struct {
	win32.IInspectableVtbl
	Get_IsSelected uintptr
	Put_Name       uintptr
	Get_Name       uintptr
	Put_Language   uintptr
	Get_Language   uintptr
}

type IMediaStreamDescriptor struct {
	win32.IInspectable
}

func (this *IMediaStreamDescriptor) Vtbl() *IMediaStreamDescriptorVtbl {
	return (*IMediaStreamDescriptorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaStreamDescriptor) Get_IsSelected() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSelected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaStreamDescriptor) Put_Name(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Name, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMediaStreamDescriptor) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaStreamDescriptor) Put_Language(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Language, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMediaStreamDescriptor) Get_Language() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Language, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 5073010F-E8B2-4071-B00B-EBF337A76B58
var IID_IMediaStreamDescriptor2 = syscall.GUID{0x5073010F, 0xE8B2, 0x4071,
	[8]byte{0xB0, 0x0B, 0xEB, 0xF3, 0x37, 0xA7, 0x6B, 0x58}}

type IMediaStreamDescriptor2Interface interface {
	win32.IInspectableInterface
	Put_Label(value string)
	Get_Label() string
}

type IMediaStreamDescriptor2Vtbl struct {
	win32.IInspectableVtbl
	Put_Label uintptr
	Get_Label uintptr
}

type IMediaStreamDescriptor2 struct {
	win32.IInspectable
}

func (this *IMediaStreamDescriptor2) Vtbl() *IMediaStreamDescriptor2Vtbl {
	return (*IMediaStreamDescriptor2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaStreamDescriptor2) Put_Label(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Label, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMediaStreamDescriptor2) Get_Label() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Label, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 5C8DB627-4B80-4361-9837-6CB7481AD9D6
var IID_IMediaStreamSample = syscall.GUID{0x5C8DB627, 0x4B80, 0x4361,
	[8]byte{0x98, 0x37, 0x6C, 0xB7, 0x48, 0x1A, 0xD9, 0xD6}}

type IMediaStreamSampleInterface interface {
	win32.IInspectableInterface
	Add_Processed(handler TypedEventHandler[*IMediaStreamSample, interface{}]) EventRegistrationToken
	Remove_Processed(token EventRegistrationToken)
	Get_Buffer() *IBuffer
	Get_Timestamp() TimeSpan
	Get_ExtendedProperties() *IMap[syscall.GUID, interface{}]
	Get_Protection() *IMediaStreamSampleProtectionProperties
	Put_DecodeTimestamp(value TimeSpan)
	Get_DecodeTimestamp() TimeSpan
	Put_Duration(value TimeSpan)
	Get_Duration() TimeSpan
	Put_KeyFrame(value bool)
	Get_KeyFrame() bool
	Put_Discontinuous(value bool)
	Get_Discontinuous() bool
}

type IMediaStreamSampleVtbl struct {
	win32.IInspectableVtbl
	Add_Processed          uintptr
	Remove_Processed       uintptr
	Get_Buffer             uintptr
	Get_Timestamp          uintptr
	Get_ExtendedProperties uintptr
	Get_Protection         uintptr
	Put_DecodeTimestamp    uintptr
	Get_DecodeTimestamp    uintptr
	Put_Duration           uintptr
	Get_Duration           uintptr
	Put_KeyFrame           uintptr
	Get_KeyFrame           uintptr
	Put_Discontinuous      uintptr
	Get_Discontinuous      uintptr
}

type IMediaStreamSample struct {
	win32.IInspectable
}

func (this *IMediaStreamSample) Vtbl() *IMediaStreamSampleVtbl {
	return (*IMediaStreamSampleVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaStreamSample) Add_Processed(handler TypedEventHandler[*IMediaStreamSample, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Processed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaStreamSample) Remove_Processed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Processed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaStreamSample) Get_Buffer() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Buffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaStreamSample) Get_Timestamp() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaStreamSample) Get_ExtendedProperties() *IMap[syscall.GUID, interface{}] {
	var _result *IMap[syscall.GUID, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaStreamSample) Get_Protection() *IMediaStreamSampleProtectionProperties {
	var _result *IMediaStreamSampleProtectionProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Protection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaStreamSample) Put_DecodeTimestamp(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DecodeTimestamp, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IMediaStreamSample) Get_DecodeTimestamp() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DecodeTimestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaStreamSample) Put_Duration(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Duration, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IMediaStreamSample) Get_Duration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Duration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaStreamSample) Put_KeyFrame(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_KeyFrame, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaStreamSample) Get_KeyFrame() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KeyFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaStreamSample) Put_Discontinuous(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Discontinuous, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaStreamSample) Get_Discontinuous() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Discontinuous, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 45078691-FCE8-4746-A1C8-10C25D3D7CD3
var IID_IMediaStreamSample2 = syscall.GUID{0x45078691, 0xFCE8, 0x4746,
	[8]byte{0xA1, 0xC8, 0x10, 0xC2, 0x5D, 0x3D, 0x7C, 0xD3}}

type IMediaStreamSample2Interface interface {
	win32.IInspectableInterface
	Get_Direct3D11Surface() unsafe.Pointer
}

type IMediaStreamSample2Vtbl struct {
	win32.IInspectableVtbl
	Get_Direct3D11Surface uintptr
}

type IMediaStreamSample2 struct {
	win32.IInspectable
}

func (this *IMediaStreamSample2) Vtbl() *IMediaStreamSample2Vtbl {
	return (*IMediaStreamSample2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaStreamSample2) Get_Direct3D11Surface() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Direct3D11Surface, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4EB88292-ECDF-493E-841D-DD4ADD7CACA2
var IID_IMediaStreamSampleProtectionProperties = syscall.GUID{0x4EB88292, 0xECDF, 0x493E,
	[8]byte{0x84, 0x1D, 0xDD, 0x4A, 0xDD, 0x7C, 0xAC, 0xA2}}

type IMediaStreamSampleProtectionPropertiesInterface interface {
	win32.IInspectableInterface
	SetKeyIdentifier(valueLength uint32, value *byte)
	GetKeyIdentifier(valueLength uint32, value *byte)
	SetInitializationVector(valueLength uint32, value *byte)
	GetInitializationVector(valueLength uint32, value *byte)
	SetSubSampleMapping(valueLength uint32, value *byte)
	GetSubSampleMapping(valueLength uint32, value *byte)
}

type IMediaStreamSampleProtectionPropertiesVtbl struct {
	win32.IInspectableVtbl
	SetKeyIdentifier        uintptr
	GetKeyIdentifier        uintptr
	SetInitializationVector uintptr
	GetInitializationVector uintptr
	SetSubSampleMapping     uintptr
	GetSubSampleMapping     uintptr
}

type IMediaStreamSampleProtectionProperties struct {
	win32.IInspectable
}

func (this *IMediaStreamSampleProtectionProperties) Vtbl() *IMediaStreamSampleProtectionPropertiesVtbl {
	return (*IMediaStreamSampleProtectionPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaStreamSampleProtectionProperties) SetKeyIdentifier(valueLength uint32, value *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetKeyIdentifier, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaStreamSampleProtectionProperties) GetKeyIdentifier(valueLength uint32, value *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetKeyIdentifier, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaStreamSampleProtectionProperties) SetInitializationVector(valueLength uint32, value *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetInitializationVector, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaStreamSampleProtectionProperties) GetInitializationVector(valueLength uint32, value *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetInitializationVector, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaStreamSampleProtectionProperties) SetSubSampleMapping(valueLength uint32, value *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetSubSampleMapping, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaStreamSampleProtectionProperties) GetSubSampleMapping(valueLength uint32, value *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSubSampleMapping, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// DFDF218F-A6CF-4579-BE41-73DD941AD972
var IID_IMediaStreamSampleStatics = syscall.GUID{0xDFDF218F, 0xA6CF, 0x4579,
	[8]byte{0xBE, 0x41, 0x73, 0xDD, 0x94, 0x1A, 0xD9, 0x72}}

type IMediaStreamSampleStaticsInterface interface {
	win32.IInspectableInterface
	CreateFromBuffer(buffer *IBuffer, timestamp TimeSpan) *IMediaStreamSample
	CreateFromStreamAsync(stream *IInputStream, count uint32, timestamp TimeSpan) *IAsyncOperation[*IMediaStreamSample]
}

type IMediaStreamSampleStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateFromBuffer      uintptr
	CreateFromStreamAsync uintptr
}

type IMediaStreamSampleStatics struct {
	win32.IInspectable
}

func (this *IMediaStreamSampleStatics) Vtbl() *IMediaStreamSampleStaticsVtbl {
	return (*IMediaStreamSampleStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaStreamSampleStatics) CreateFromBuffer(buffer *IBuffer, timestamp TimeSpan) *IMediaStreamSample {
	var _result *IMediaStreamSample
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromBuffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(buffer)), *(*uintptr)(unsafe.Pointer(&timestamp)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaStreamSampleStatics) CreateFromStreamAsync(stream *IInputStream, count uint32, timestamp TimeSpan) *IAsyncOperation[*IMediaStreamSample] {
	var _result *IAsyncOperation[*IMediaStreamSample]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromStreamAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)), uintptr(count), *(*uintptr)(unsafe.Pointer(&timestamp)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9EFE9521-6D46-494C-A2F8-D662922E2DD7
var IID_IMediaStreamSampleStatics2 = syscall.GUID{0x9EFE9521, 0x6D46, 0x494C,
	[8]byte{0xA2, 0xF8, 0xD6, 0x62, 0x92, 0x2E, 0x2D, 0xD7}}

type IMediaStreamSampleStatics2Interface interface {
	win32.IInspectableInterface
	CreateFromDirect3D11Surface(surface unsafe.Pointer, timestamp TimeSpan) *IMediaStreamSample
}

type IMediaStreamSampleStatics2Vtbl struct {
	win32.IInspectableVtbl
	CreateFromDirect3D11Surface uintptr
}

type IMediaStreamSampleStatics2 struct {
	win32.IInspectable
}

func (this *IMediaStreamSampleStatics2) Vtbl() *IMediaStreamSampleStatics2Vtbl {
	return (*IMediaStreamSampleStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaStreamSampleStatics2) CreateFromDirect3D11Surface(surface unsafe.Pointer, timestamp TimeSpan) *IMediaStreamSample {
	var _result *IMediaStreamSample
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromDirect3D11Surface, uintptr(unsafe.Pointer(this)), uintptr(surface), *(*uintptr)(unsafe.Pointer(&timestamp)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3712D543-45EB-4138-AA62-C01E26F3843F
var IID_IMediaStreamSource = syscall.GUID{0x3712D543, 0x45EB, 0x4138,
	[8]byte{0xAA, 0x62, 0xC0, 0x1E, 0x26, 0xF3, 0x84, 0x3F}}

type IMediaStreamSourceInterface interface {
	win32.IInspectableInterface
	Add_Closed(handler TypedEventHandler[*IMediaStreamSource, *IMediaStreamSourceClosedEventArgs]) EventRegistrationToken
	Remove_Closed(token EventRegistrationToken)
	Add_Starting(handler TypedEventHandler[*IMediaStreamSource, *IMediaStreamSourceStartingEventArgs]) EventRegistrationToken
	Remove_Starting(token EventRegistrationToken)
	Add_Paused(handler TypedEventHandler[*IMediaStreamSource, interface{}]) EventRegistrationToken
	Remove_Paused(token EventRegistrationToken)
	Add_SampleRequested(handler TypedEventHandler[*IMediaStreamSource, *IMediaStreamSourceSampleRequestedEventArgs]) EventRegistrationToken
	Remove_SampleRequested(token EventRegistrationToken)
	Add_SwitchStreamsRequested(handler TypedEventHandler[*IMediaStreamSource, *IMediaStreamSourceSwitchStreamsRequestedEventArgs]) EventRegistrationToken
	Remove_SwitchStreamsRequested(token EventRegistrationToken)
	NotifyError(errorStatus MediaStreamSourceErrorStatus)
	AddStreamDescriptor(descriptor *IMediaStreamDescriptor)
	Put_MediaProtectionManager(value *IMediaProtectionManager)
	Get_MediaProtectionManager() *IMediaProtectionManager
	Put_Duration(value TimeSpan)
	Get_Duration() TimeSpan
	Put_CanSeek(value bool)
	Get_CanSeek() bool
	Put_BufferTime(value TimeSpan)
	Get_BufferTime() TimeSpan
	SetBufferedRange(startOffset TimeSpan, endOffset TimeSpan)
	Get_MusicProperties() *IMusicProperties
	Get_VideoProperties() *IVideoProperties
	Put_Thumbnail(value *IRandomAccessStreamReference)
	Get_Thumbnail() *IRandomAccessStreamReference
	AddProtectionKey(streamDescriptor *IMediaStreamDescriptor, keyIdentifierLength uint32, keyIdentifier *byte, licenseDataLength uint32, licenseData *byte)
}

type IMediaStreamSourceVtbl struct {
	win32.IInspectableVtbl
	Add_Closed                    uintptr
	Remove_Closed                 uintptr
	Add_Starting                  uintptr
	Remove_Starting               uintptr
	Add_Paused                    uintptr
	Remove_Paused                 uintptr
	Add_SampleRequested           uintptr
	Remove_SampleRequested        uintptr
	Add_SwitchStreamsRequested    uintptr
	Remove_SwitchStreamsRequested uintptr
	NotifyError                   uintptr
	AddStreamDescriptor           uintptr
	Put_MediaProtectionManager    uintptr
	Get_MediaProtectionManager    uintptr
	Put_Duration                  uintptr
	Get_Duration                  uintptr
	Put_CanSeek                   uintptr
	Get_CanSeek                   uintptr
	Put_BufferTime                uintptr
	Get_BufferTime                uintptr
	SetBufferedRange              uintptr
	Get_MusicProperties           uintptr
	Get_VideoProperties           uintptr
	Put_Thumbnail                 uintptr
	Get_Thumbnail                 uintptr
	AddProtectionKey              uintptr
}

type IMediaStreamSource struct {
	win32.IInspectable
}

func (this *IMediaStreamSource) Vtbl() *IMediaStreamSourceVtbl {
	return (*IMediaStreamSourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaStreamSource) Add_Closed(handler TypedEventHandler[*IMediaStreamSource, *IMediaStreamSourceClosedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Closed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaStreamSource) Remove_Closed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Closed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaStreamSource) Add_Starting(handler TypedEventHandler[*IMediaStreamSource, *IMediaStreamSourceStartingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Starting, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaStreamSource) Remove_Starting(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Starting, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaStreamSource) Add_Paused(handler TypedEventHandler[*IMediaStreamSource, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Paused, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaStreamSource) Remove_Paused(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Paused, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaStreamSource) Add_SampleRequested(handler TypedEventHandler[*IMediaStreamSource, *IMediaStreamSourceSampleRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SampleRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaStreamSource) Remove_SampleRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SampleRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaStreamSource) Add_SwitchStreamsRequested(handler TypedEventHandler[*IMediaStreamSource, *IMediaStreamSourceSwitchStreamsRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SwitchStreamsRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaStreamSource) Remove_SwitchStreamsRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SwitchStreamsRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaStreamSource) NotifyError(errorStatus MediaStreamSourceErrorStatus) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().NotifyError, uintptr(unsafe.Pointer(this)), uintptr(errorStatus))
	_ = _hr
}

func (this *IMediaStreamSource) AddStreamDescriptor(descriptor *IMediaStreamDescriptor) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddStreamDescriptor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(descriptor)))
	_ = _hr
}

func (this *IMediaStreamSource) Put_MediaProtectionManager(value *IMediaProtectionManager) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MediaProtectionManager, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaStreamSource) Get_MediaProtectionManager() *IMediaProtectionManager {
	var _result *IMediaProtectionManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaProtectionManager, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaStreamSource) Put_Duration(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Duration, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IMediaStreamSource) Get_Duration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Duration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaStreamSource) Put_CanSeek(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CanSeek, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaStreamSource) Get_CanSeek() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanSeek, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaStreamSource) Put_BufferTime(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BufferTime, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IMediaStreamSource) Get_BufferTime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BufferTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaStreamSource) SetBufferedRange(startOffset TimeSpan, endOffset TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetBufferedRange, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&startOffset)), *(*uintptr)(unsafe.Pointer(&endOffset)))
	_ = _hr
}

func (this *IMediaStreamSource) Get_MusicProperties() *IMusicProperties {
	var _result *IMusicProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MusicProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaStreamSource) Get_VideoProperties() *IVideoProperties {
	var _result *IVideoProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaStreamSource) Put_Thumbnail(value *IRandomAccessStreamReference) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Thumbnail, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaStreamSource) Get_Thumbnail() *IRandomAccessStreamReference {
	var _result *IRandomAccessStreamReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Thumbnail, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaStreamSource) AddProtectionKey(streamDescriptor *IMediaStreamDescriptor, keyIdentifierLength uint32, keyIdentifier *byte, licenseDataLength uint32, licenseData *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddProtectionKey, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(streamDescriptor)), uintptr(keyIdentifierLength), uintptr(unsafe.Pointer(keyIdentifier)), uintptr(licenseDataLength), uintptr(unsafe.Pointer(licenseData)))
	_ = _hr
}

// EC55D0AD-2E6A-4F74-ADBB-B562D1533849
var IID_IMediaStreamSource2 = syscall.GUID{0xEC55D0AD, 0x2E6A, 0x4F74,
	[8]byte{0xAD, 0xBB, 0xB5, 0x62, 0xD1, 0x53, 0x38, 0x49}}

type IMediaStreamSource2Interface interface {
	win32.IInspectableInterface
	Add_SampleRendered(handler TypedEventHandler[*IMediaStreamSource, *IMediaStreamSourceSampleRenderedEventArgs]) EventRegistrationToken
	Remove_SampleRendered(token EventRegistrationToken)
}

type IMediaStreamSource2Vtbl struct {
	win32.IInspectableVtbl
	Add_SampleRendered    uintptr
	Remove_SampleRendered uintptr
}

type IMediaStreamSource2 struct {
	win32.IInspectable
}

func (this *IMediaStreamSource2) Vtbl() *IMediaStreamSource2Vtbl {
	return (*IMediaStreamSource2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaStreamSource2) Add_SampleRendered(handler TypedEventHandler[*IMediaStreamSource, *IMediaStreamSourceSampleRenderedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SampleRendered, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaStreamSource2) Remove_SampleRendered(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SampleRendered, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 6A2A2746-3DDD-4DDF-A121-94045ECF9440
var IID_IMediaStreamSource3 = syscall.GUID{0x6A2A2746, 0x3DDD, 0x4DDF,
	[8]byte{0xA1, 0x21, 0x94, 0x04, 0x5E, 0xCF, 0x94, 0x40}}

type IMediaStreamSource3Interface interface {
	win32.IInspectableInterface
	Put_MaxSupportedPlaybackRate(value *IReference[float64])
	Get_MaxSupportedPlaybackRate() *IReference[float64]
}

type IMediaStreamSource3Vtbl struct {
	win32.IInspectableVtbl
	Put_MaxSupportedPlaybackRate uintptr
	Get_MaxSupportedPlaybackRate uintptr
}

type IMediaStreamSource3 struct {
	win32.IInspectable
}

func (this *IMediaStreamSource3) Vtbl() *IMediaStreamSource3Vtbl {
	return (*IMediaStreamSource3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaStreamSource3) Put_MaxSupportedPlaybackRate(value *IReference[float64]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxSupportedPlaybackRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaStreamSource3) Get_MaxSupportedPlaybackRate() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxSupportedPlaybackRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1D0CFCAB-830D-417C-A3A9-2454FD6415C7
var IID_IMediaStreamSource4 = syscall.GUID{0x1D0CFCAB, 0x830D, 0x417C,
	[8]byte{0xA3, 0xA9, 0x24, 0x54, 0xFD, 0x64, 0x15, 0xC7}}

type IMediaStreamSource4Interface interface {
	win32.IInspectableInterface
	Put_IsLive(value bool)
	Get_IsLive() bool
}

type IMediaStreamSource4Vtbl struct {
	win32.IInspectableVtbl
	Put_IsLive uintptr
	Get_IsLive uintptr
}

type IMediaStreamSource4 struct {
	win32.IInspectable
}

func (this *IMediaStreamSource4) Vtbl() *IMediaStreamSource4Vtbl {
	return (*IMediaStreamSource4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaStreamSource4) Put_IsLive(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsLive, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaStreamSource4) Get_IsLive() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsLive, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// CD8C7EB2-4816-4E24-88F0-491EF7386406
var IID_IMediaStreamSourceClosedEventArgs = syscall.GUID{0xCD8C7EB2, 0x4816, 0x4E24,
	[8]byte{0x88, 0xF0, 0x49, 0x1E, 0xF7, 0x38, 0x64, 0x06}}

type IMediaStreamSourceClosedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Request() *IMediaStreamSourceClosedRequest
}

type IMediaStreamSourceClosedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Request uintptr
}

type IMediaStreamSourceClosedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaStreamSourceClosedEventArgs) Vtbl() *IMediaStreamSourceClosedEventArgsVtbl {
	return (*IMediaStreamSourceClosedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaStreamSourceClosedEventArgs) Get_Request() *IMediaStreamSourceClosedRequest {
	var _result *IMediaStreamSourceClosedRequest
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Request, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 907C00E9-18A3-4951-887A-2C1EEBD5C69E
var IID_IMediaStreamSourceClosedRequest = syscall.GUID{0x907C00E9, 0x18A3, 0x4951,
	[8]byte{0x88, 0x7A, 0x2C, 0x1E, 0xEB, 0xD5, 0xC6, 0x9E}}

type IMediaStreamSourceClosedRequestInterface interface {
	win32.IInspectableInterface
	Get_Reason() MediaStreamSourceClosedReason
}

type IMediaStreamSourceClosedRequestVtbl struct {
	win32.IInspectableVtbl
	Get_Reason uintptr
}

type IMediaStreamSourceClosedRequest struct {
	win32.IInspectable
}

func (this *IMediaStreamSourceClosedRequest) Vtbl() *IMediaStreamSourceClosedRequestVtbl {
	return (*IMediaStreamSourceClosedRequestVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaStreamSourceClosedRequest) Get_Reason() MediaStreamSourceClosedReason {
	var _result MediaStreamSourceClosedReason
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Reason, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// EF77E0D9-D158-4B7A-863F-203342FBFD41
var IID_IMediaStreamSourceFactory = syscall.GUID{0xEF77E0D9, 0xD158, 0x4B7A,
	[8]byte{0x86, 0x3F, 0x20, 0x33, 0x42, 0xFB, 0xFD, 0x41}}

type IMediaStreamSourceFactoryInterface interface {
	win32.IInspectableInterface
	CreateFromDescriptor(descriptor *IMediaStreamDescriptor) *IMediaStreamSource
	CreateFromDescriptors(descriptor *IMediaStreamDescriptor, descriptor2 *IMediaStreamDescriptor) *IMediaStreamSource
}

type IMediaStreamSourceFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateFromDescriptor  uintptr
	CreateFromDescriptors uintptr
}

type IMediaStreamSourceFactory struct {
	win32.IInspectable
}

func (this *IMediaStreamSourceFactory) Vtbl() *IMediaStreamSourceFactoryVtbl {
	return (*IMediaStreamSourceFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaStreamSourceFactory) CreateFromDescriptor(descriptor *IMediaStreamDescriptor) *IMediaStreamSource {
	var _result *IMediaStreamSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromDescriptor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(descriptor)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaStreamSourceFactory) CreateFromDescriptors(descriptor *IMediaStreamDescriptor, descriptor2 *IMediaStreamDescriptor) *IMediaStreamSource {
	var _result *IMediaStreamSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromDescriptors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(descriptor)), uintptr(unsafe.Pointer(descriptor2)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9D697B05-D4F2-4C7A-9DFE-8D6CD0B3EE84
var IID_IMediaStreamSourceSampleRenderedEventArgs = syscall.GUID{0x9D697B05, 0xD4F2, 0x4C7A,
	[8]byte{0x9D, 0xFE, 0x8D, 0x6C, 0xD0, 0xB3, 0xEE, 0x84}}

type IMediaStreamSourceSampleRenderedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_SampleLag() TimeSpan
}

type IMediaStreamSourceSampleRenderedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_SampleLag uintptr
}

type IMediaStreamSourceSampleRenderedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaStreamSourceSampleRenderedEventArgs) Vtbl() *IMediaStreamSourceSampleRenderedEventArgsVtbl {
	return (*IMediaStreamSourceSampleRenderedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaStreamSourceSampleRenderedEventArgs) Get_SampleLag() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SampleLag, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4DB341A9-3501-4D9B-83F9-8F235C822532
var IID_IMediaStreamSourceSampleRequest = syscall.GUID{0x4DB341A9, 0x3501, 0x4D9B,
	[8]byte{0x83, 0xF9, 0x8F, 0x23, 0x5C, 0x82, 0x25, 0x32}}

type IMediaStreamSourceSampleRequestInterface interface {
	win32.IInspectableInterface
	Get_StreamDescriptor() *IMediaStreamDescriptor
	GetDeferral() *IMediaStreamSourceSampleRequestDeferral
	Put_Sample(value *IMediaStreamSample)
	Get_Sample() *IMediaStreamSample
	ReportSampleProgress(progress uint32)
}

type IMediaStreamSourceSampleRequestVtbl struct {
	win32.IInspectableVtbl
	Get_StreamDescriptor uintptr
	GetDeferral          uintptr
	Put_Sample           uintptr
	Get_Sample           uintptr
	ReportSampleProgress uintptr
}

type IMediaStreamSourceSampleRequest struct {
	win32.IInspectable
}

func (this *IMediaStreamSourceSampleRequest) Vtbl() *IMediaStreamSourceSampleRequestVtbl {
	return (*IMediaStreamSourceSampleRequestVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaStreamSourceSampleRequest) Get_StreamDescriptor() *IMediaStreamDescriptor {
	var _result *IMediaStreamDescriptor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StreamDescriptor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaStreamSourceSampleRequest) GetDeferral() *IMediaStreamSourceSampleRequestDeferral {
	var _result *IMediaStreamSourceSampleRequestDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaStreamSourceSampleRequest) Put_Sample(value *IMediaStreamSample) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Sample, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaStreamSourceSampleRequest) Get_Sample() *IMediaStreamSample {
	var _result *IMediaStreamSample
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Sample, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaStreamSourceSampleRequest) ReportSampleProgress(progress uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReportSampleProgress, uintptr(unsafe.Pointer(this)), uintptr(progress))
	_ = _hr
}

// 7895CC02-F982-43C8-9D16-C62D999319BE
var IID_IMediaStreamSourceSampleRequestDeferral = syscall.GUID{0x7895CC02, 0xF982, 0x43C8,
	[8]byte{0x9D, 0x16, 0xC6, 0x2D, 0x99, 0x93, 0x19, 0xBE}}

type IMediaStreamSourceSampleRequestDeferralInterface interface {
	win32.IInspectableInterface
	Complete()
}

type IMediaStreamSourceSampleRequestDeferralVtbl struct {
	win32.IInspectableVtbl
	Complete uintptr
}

type IMediaStreamSourceSampleRequestDeferral struct {
	win32.IInspectable
}

func (this *IMediaStreamSourceSampleRequestDeferral) Vtbl() *IMediaStreamSourceSampleRequestDeferralVtbl {
	return (*IMediaStreamSourceSampleRequestDeferralVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaStreamSourceSampleRequestDeferral) Complete() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Complete, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 10F9BB9E-71C5-492F-847F-0DA1F35E81F8
var IID_IMediaStreamSourceSampleRequestedEventArgs = syscall.GUID{0x10F9BB9E, 0x71C5, 0x492F,
	[8]byte{0x84, 0x7F, 0x0D, 0xA1, 0xF3, 0x5E, 0x81, 0xF8}}

type IMediaStreamSourceSampleRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Request() *IMediaStreamSourceSampleRequest
}

type IMediaStreamSourceSampleRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Request uintptr
}

type IMediaStreamSourceSampleRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaStreamSourceSampleRequestedEventArgs) Vtbl() *IMediaStreamSourceSampleRequestedEventArgsVtbl {
	return (*IMediaStreamSourceSampleRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaStreamSourceSampleRequestedEventArgs) Get_Request() *IMediaStreamSourceSampleRequest {
	var _result *IMediaStreamSourceSampleRequest
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Request, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F41468F2-C274-4940-A5BB-28A572452FA7
var IID_IMediaStreamSourceStartingEventArgs = syscall.GUID{0xF41468F2, 0xC274, 0x4940,
	[8]byte{0xA5, 0xBB, 0x28, 0xA5, 0x72, 0x45, 0x2F, 0xA7}}

type IMediaStreamSourceStartingEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Request() *IMediaStreamSourceStartingRequest
}

type IMediaStreamSourceStartingEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Request uintptr
}

type IMediaStreamSourceStartingEventArgs struct {
	win32.IInspectable
}

func (this *IMediaStreamSourceStartingEventArgs) Vtbl() *IMediaStreamSourceStartingEventArgsVtbl {
	return (*IMediaStreamSourceStartingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaStreamSourceStartingEventArgs) Get_Request() *IMediaStreamSourceStartingRequest {
	var _result *IMediaStreamSourceStartingRequest
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Request, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2A9093E4-35C4-4B1B-A791-0D99DB56DD1D
var IID_IMediaStreamSourceStartingRequest = syscall.GUID{0x2A9093E4, 0x35C4, 0x4B1B,
	[8]byte{0xA7, 0x91, 0x0D, 0x99, 0xDB, 0x56, 0xDD, 0x1D}}

type IMediaStreamSourceStartingRequestInterface interface {
	win32.IInspectableInterface
	Get_StartPosition() *IReference[TimeSpan]
	GetDeferral() *IMediaStreamSourceStartingRequestDeferral
	SetActualStartPosition(position TimeSpan)
}

type IMediaStreamSourceStartingRequestVtbl struct {
	win32.IInspectableVtbl
	Get_StartPosition      uintptr
	GetDeferral            uintptr
	SetActualStartPosition uintptr
}

type IMediaStreamSourceStartingRequest struct {
	win32.IInspectable
}

func (this *IMediaStreamSourceStartingRequest) Vtbl() *IMediaStreamSourceStartingRequestVtbl {
	return (*IMediaStreamSourceStartingRequestVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaStreamSourceStartingRequest) Get_StartPosition() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StartPosition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaStreamSourceStartingRequest) GetDeferral() *IMediaStreamSourceStartingRequestDeferral {
	var _result *IMediaStreamSourceStartingRequestDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaStreamSourceStartingRequest) SetActualStartPosition(position TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetActualStartPosition, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&position)))
	_ = _hr
}

// 3F1356A5-6340-4DC4-9910-068ED9F598F8
var IID_IMediaStreamSourceStartingRequestDeferral = syscall.GUID{0x3F1356A5, 0x6340, 0x4DC4,
	[8]byte{0x99, 0x10, 0x06, 0x8E, 0xD9, 0xF5, 0x98, 0xF8}}

type IMediaStreamSourceStartingRequestDeferralInterface interface {
	win32.IInspectableInterface
	Complete()
}

type IMediaStreamSourceStartingRequestDeferralVtbl struct {
	win32.IInspectableVtbl
	Complete uintptr
}

type IMediaStreamSourceStartingRequestDeferral struct {
	win32.IInspectable
}

func (this *IMediaStreamSourceStartingRequestDeferral) Vtbl() *IMediaStreamSourceStartingRequestDeferralVtbl {
	return (*IMediaStreamSourceStartingRequestDeferralVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaStreamSourceStartingRequestDeferral) Complete() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Complete, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 41B8808E-38A9-4EC3-9BA0-B69B85501E90
var IID_IMediaStreamSourceSwitchStreamsRequest = syscall.GUID{0x41B8808E, 0x38A9, 0x4EC3,
	[8]byte{0x9B, 0xA0, 0xB6, 0x9B, 0x85, 0x50, 0x1E, 0x90}}

type IMediaStreamSourceSwitchStreamsRequestInterface interface {
	win32.IInspectableInterface
	Get_OldStreamDescriptor() *IMediaStreamDescriptor
	Get_NewStreamDescriptor() *IMediaStreamDescriptor
	GetDeferral() *IMediaStreamSourceSwitchStreamsRequestDeferral
}

type IMediaStreamSourceSwitchStreamsRequestVtbl struct {
	win32.IInspectableVtbl
	Get_OldStreamDescriptor uintptr
	Get_NewStreamDescriptor uintptr
	GetDeferral             uintptr
}

type IMediaStreamSourceSwitchStreamsRequest struct {
	win32.IInspectable
}

func (this *IMediaStreamSourceSwitchStreamsRequest) Vtbl() *IMediaStreamSourceSwitchStreamsRequestVtbl {
	return (*IMediaStreamSourceSwitchStreamsRequestVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaStreamSourceSwitchStreamsRequest) Get_OldStreamDescriptor() *IMediaStreamDescriptor {
	var _result *IMediaStreamDescriptor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OldStreamDescriptor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaStreamSourceSwitchStreamsRequest) Get_NewStreamDescriptor() *IMediaStreamDescriptor {
	var _result *IMediaStreamDescriptor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NewStreamDescriptor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaStreamSourceSwitchStreamsRequest) GetDeferral() *IMediaStreamSourceSwitchStreamsRequestDeferral {
	var _result *IMediaStreamSourceSwitchStreamsRequestDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BEE3D835-A505-4F9A-B943-2B8CB1B4BBD9
var IID_IMediaStreamSourceSwitchStreamsRequestDeferral = syscall.GUID{0xBEE3D835, 0xA505, 0x4F9A,
	[8]byte{0xB9, 0x43, 0x2B, 0x8C, 0xB1, 0xB4, 0xBB, 0xD9}}

type IMediaStreamSourceSwitchStreamsRequestDeferralInterface interface {
	win32.IInspectableInterface
	Complete()
}

type IMediaStreamSourceSwitchStreamsRequestDeferralVtbl struct {
	win32.IInspectableVtbl
	Complete uintptr
}

type IMediaStreamSourceSwitchStreamsRequestDeferral struct {
	win32.IInspectable
}

func (this *IMediaStreamSourceSwitchStreamsRequestDeferral) Vtbl() *IMediaStreamSourceSwitchStreamsRequestDeferralVtbl {
	return (*IMediaStreamSourceSwitchStreamsRequestDeferralVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaStreamSourceSwitchStreamsRequestDeferral) Complete() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Complete, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 42202B72-6EA1-4677-981E-350A0DA412AA
var IID_IMediaStreamSourceSwitchStreamsRequestedEventArgs = syscall.GUID{0x42202B72, 0x6EA1, 0x4677,
	[8]byte{0x98, 0x1E, 0x35, 0x0A, 0x0D, 0xA4, 0x12, 0xAA}}

type IMediaStreamSourceSwitchStreamsRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Request() *IMediaStreamSourceSwitchStreamsRequest
}

type IMediaStreamSourceSwitchStreamsRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Request uintptr
}

type IMediaStreamSourceSwitchStreamsRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaStreamSourceSwitchStreamsRequestedEventArgs) Vtbl() *IMediaStreamSourceSwitchStreamsRequestedEventArgsVtbl {
	return (*IMediaStreamSourceSwitchStreamsRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaStreamSourceSwitchStreamsRequestedEventArgs) Get_Request() *IMediaStreamSourceSwitchStreamsRequest {
	var _result *IMediaStreamSourceSwitchStreamsRequest
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Request, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 03E1FAFC-C931-491A-B46B-C10EE8C256B7
var IID_IMediaTrack = syscall.GUID{0x03E1FAFC, 0xC931, 0x491A,
	[8]byte{0xB4, 0x6B, 0xC1, 0x0E, 0xE8, 0xC2, 0x56, 0xB7}}

type IMediaTrackInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	Get_Language() string
	Get_TrackKind() MediaTrackKind
	Put_Label(value string)
	Get_Label() string
}

type IMediaTrackVtbl struct {
	win32.IInspectableVtbl
	Get_Id        uintptr
	Get_Language  uintptr
	Get_TrackKind uintptr
	Put_Label     uintptr
	Get_Label     uintptr
}

type IMediaTrack struct {
	win32.IInspectable
}

func (this *IMediaTrack) Vtbl() *IMediaTrackVtbl {
	return (*IMediaTrackVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaTrack) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaTrack) Get_Language() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Language, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaTrack) Get_TrackKind() MediaTrackKind {
	var _result MediaTrackKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TrackKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaTrack) Put_Label(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Label, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMediaTrack) Get_Label() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Label, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 0C1AA3E3-DF8D-4079-A3FE-6849184B4E2F
var IID_IMseSourceBuffer = syscall.GUID{0x0C1AA3E3, 0xDF8D, 0x4079,
	[8]byte{0xA3, 0xFE, 0x68, 0x49, 0x18, 0x4B, 0x4E, 0x2F}}

type IMseSourceBufferInterface interface {
	win32.IInspectableInterface
	Add_UpdateStarting(handler TypedEventHandler[*IMseSourceBuffer, interface{}]) EventRegistrationToken
	Remove_UpdateStarting(token EventRegistrationToken)
	Add_Updated(handler TypedEventHandler[*IMseSourceBuffer, interface{}]) EventRegistrationToken
	Remove_Updated(token EventRegistrationToken)
	Add_UpdateEnded(handler TypedEventHandler[*IMseSourceBuffer, interface{}]) EventRegistrationToken
	Remove_UpdateEnded(token EventRegistrationToken)
	Add_ErrorOccurred(handler TypedEventHandler[*IMseSourceBuffer, interface{}]) EventRegistrationToken
	Remove_ErrorOccurred(token EventRegistrationToken)
	Add_Aborted(handler TypedEventHandler[*IMseSourceBuffer, interface{}]) EventRegistrationToken
	Remove_Aborted(token EventRegistrationToken)
	Get_Mode() MseAppendMode
	Put_Mode(value MseAppendMode)
	Get_IsUpdating() bool
	Get_Buffered() *IVectorView[MseTimeRange]
	Get_TimestampOffset() TimeSpan
	Put_TimestampOffset(value TimeSpan)
	Get_AppendWindowStart() TimeSpan
	Put_AppendWindowStart(value TimeSpan)
	Get_AppendWindowEnd() *IReference[TimeSpan]
	Put_AppendWindowEnd(value *IReference[TimeSpan])
	AppendBuffer(buffer *IBuffer)
	AppendStream(stream *IInputStream)
	AppendStreamMaxSize(stream *IInputStream, maxSize uint64)
	Abort()
	Remove(start TimeSpan, end *IReference[TimeSpan])
}

type IMseSourceBufferVtbl struct {
	win32.IInspectableVtbl
	Add_UpdateStarting    uintptr
	Remove_UpdateStarting uintptr
	Add_Updated           uintptr
	Remove_Updated        uintptr
	Add_UpdateEnded       uintptr
	Remove_UpdateEnded    uintptr
	Add_ErrorOccurred     uintptr
	Remove_ErrorOccurred  uintptr
	Add_Aborted           uintptr
	Remove_Aborted        uintptr
	Get_Mode              uintptr
	Put_Mode              uintptr
	Get_IsUpdating        uintptr
	Get_Buffered          uintptr
	Get_TimestampOffset   uintptr
	Put_TimestampOffset   uintptr
	Get_AppendWindowStart uintptr
	Put_AppendWindowStart uintptr
	Get_AppendWindowEnd   uintptr
	Put_AppendWindowEnd   uintptr
	AppendBuffer          uintptr
	AppendStream          uintptr
	AppendStreamMaxSize   uintptr
	Abort                 uintptr
	Remove                uintptr
}

type IMseSourceBuffer struct {
	win32.IInspectable
}

func (this *IMseSourceBuffer) Vtbl() *IMseSourceBufferVtbl {
	return (*IMseSourceBufferVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMseSourceBuffer) Add_UpdateStarting(handler TypedEventHandler[*IMseSourceBuffer, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_UpdateStarting, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMseSourceBuffer) Remove_UpdateStarting(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_UpdateStarting, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMseSourceBuffer) Add_Updated(handler TypedEventHandler[*IMseSourceBuffer, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Updated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMseSourceBuffer) Remove_Updated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Updated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMseSourceBuffer) Add_UpdateEnded(handler TypedEventHandler[*IMseSourceBuffer, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_UpdateEnded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMseSourceBuffer) Remove_UpdateEnded(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_UpdateEnded, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMseSourceBuffer) Add_ErrorOccurred(handler TypedEventHandler[*IMseSourceBuffer, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ErrorOccurred, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMseSourceBuffer) Remove_ErrorOccurred(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ErrorOccurred, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMseSourceBuffer) Add_Aborted(handler TypedEventHandler[*IMseSourceBuffer, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Aborted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMseSourceBuffer) Remove_Aborted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Aborted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMseSourceBuffer) Get_Mode() MseAppendMode {
	var _result MseAppendMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMseSourceBuffer) Put_Mode(value MseAppendMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Mode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMseSourceBuffer) Get_IsUpdating() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsUpdating, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMseSourceBuffer) Get_Buffered() *IVectorView[MseTimeRange] {
	var _result *IVectorView[MseTimeRange]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Buffered, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMseSourceBuffer) Get_TimestampOffset() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TimestampOffset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMseSourceBuffer) Put_TimestampOffset(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TimestampOffset, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IMseSourceBuffer) Get_AppendWindowStart() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppendWindowStart, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMseSourceBuffer) Put_AppendWindowStart(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AppendWindowStart, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IMseSourceBuffer) Get_AppendWindowEnd() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppendWindowEnd, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMseSourceBuffer) Put_AppendWindowEnd(value *IReference[TimeSpan]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AppendWindowEnd, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMseSourceBuffer) AppendBuffer(buffer *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AppendBuffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(buffer)))
	_ = _hr
}

func (this *IMseSourceBuffer) AppendStream(stream *IInputStream) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AppendStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)))
	_ = _hr
}

func (this *IMseSourceBuffer) AppendStreamMaxSize(stream *IInputStream, maxSize uint64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AppendStreamMaxSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)), uintptr(maxSize))
	_ = _hr
}

func (this *IMseSourceBuffer) Abort() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Abort, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IMseSourceBuffer) Remove(start TimeSpan, end *IReference[TimeSpan]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&start)), uintptr(unsafe.Pointer(end)))
	_ = _hr
}

// 95FAE8E7-A8E7-4EBF-8927-145E940BA511
var IID_IMseSourceBufferList = syscall.GUID{0x95FAE8E7, 0xA8E7, 0x4EBF,
	[8]byte{0x89, 0x27, 0x14, 0x5E, 0x94, 0x0B, 0xA5, 0x11}}

type IMseSourceBufferListInterface interface {
	win32.IInspectableInterface
	Add_SourceBufferAdded(handler TypedEventHandler[*IMseSourceBufferList, interface{}]) EventRegistrationToken
	Remove_SourceBufferAdded(token EventRegistrationToken)
	Add_SourceBufferRemoved(handler TypedEventHandler[*IMseSourceBufferList, interface{}]) EventRegistrationToken
	Remove_SourceBufferRemoved(token EventRegistrationToken)
	Get_Buffers() *IVectorView[*IMseSourceBuffer]
}

type IMseSourceBufferListVtbl struct {
	win32.IInspectableVtbl
	Add_SourceBufferAdded      uintptr
	Remove_SourceBufferAdded   uintptr
	Add_SourceBufferRemoved    uintptr
	Remove_SourceBufferRemoved uintptr
	Get_Buffers                uintptr
}

type IMseSourceBufferList struct {
	win32.IInspectable
}

func (this *IMseSourceBufferList) Vtbl() *IMseSourceBufferListVtbl {
	return (*IMseSourceBufferListVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMseSourceBufferList) Add_SourceBufferAdded(handler TypedEventHandler[*IMseSourceBufferList, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SourceBufferAdded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMseSourceBufferList) Remove_SourceBufferAdded(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SourceBufferAdded, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMseSourceBufferList) Add_SourceBufferRemoved(handler TypedEventHandler[*IMseSourceBufferList, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SourceBufferRemoved, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMseSourceBufferList) Remove_SourceBufferRemoved(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SourceBufferRemoved, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMseSourceBufferList) Get_Buffers() *IVectorView[*IMseSourceBuffer] {
	var _result *IVectorView[*IMseSourceBuffer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Buffers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B0B4198D-02F4-4923-88DD-81BC3F360FFA
var IID_IMseStreamSource = syscall.GUID{0xB0B4198D, 0x02F4, 0x4923,
	[8]byte{0x88, 0xDD, 0x81, 0xBC, 0x3F, 0x36, 0x0F, 0xFA}}

type IMseStreamSourceInterface interface {
	win32.IInspectableInterface
	Add_Opened(handler TypedEventHandler[*IMseStreamSource, interface{}]) EventRegistrationToken
	Remove_Opened(token EventRegistrationToken)
	Add_Ended(handler TypedEventHandler[*IMseStreamSource, interface{}]) EventRegistrationToken
	Remove_Ended(token EventRegistrationToken)
	Add_Closed(handler TypedEventHandler[*IMseStreamSource, interface{}]) EventRegistrationToken
	Remove_Closed(token EventRegistrationToken)
	Get_SourceBuffers() *IMseSourceBufferList
	Get_ActiveSourceBuffers() *IMseSourceBufferList
	Get_ReadyState() MseReadyState
	Get_Duration() *IReference[TimeSpan]
	Put_Duration(value *IReference[TimeSpan])
	AddSourceBuffer(mimeType string) *IMseSourceBuffer
	RemoveSourceBuffer(buffer *IMseSourceBuffer)
	EndOfStream(status MseEndOfStreamStatus)
}

type IMseStreamSourceVtbl struct {
	win32.IInspectableVtbl
	Add_Opened              uintptr
	Remove_Opened           uintptr
	Add_Ended               uintptr
	Remove_Ended            uintptr
	Add_Closed              uintptr
	Remove_Closed           uintptr
	Get_SourceBuffers       uintptr
	Get_ActiveSourceBuffers uintptr
	Get_ReadyState          uintptr
	Get_Duration            uintptr
	Put_Duration            uintptr
	AddSourceBuffer         uintptr
	RemoveSourceBuffer      uintptr
	EndOfStream             uintptr
}

type IMseStreamSource struct {
	win32.IInspectable
}

func (this *IMseStreamSource) Vtbl() *IMseStreamSourceVtbl {
	return (*IMseStreamSourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMseStreamSource) Add_Opened(handler TypedEventHandler[*IMseStreamSource, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Opened, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMseStreamSource) Remove_Opened(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Opened, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMseStreamSource) Add_Ended(handler TypedEventHandler[*IMseStreamSource, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Ended, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMseStreamSource) Remove_Ended(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Ended, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMseStreamSource) Add_Closed(handler TypedEventHandler[*IMseStreamSource, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Closed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMseStreamSource) Remove_Closed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Closed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMseStreamSource) Get_SourceBuffers() *IMseSourceBufferList {
	var _result *IMseSourceBufferList
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SourceBuffers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMseStreamSource) Get_ActiveSourceBuffers() *IMseSourceBufferList {
	var _result *IMseSourceBufferList
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ActiveSourceBuffers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMseStreamSource) Get_ReadyState() MseReadyState {
	var _result MseReadyState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReadyState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMseStreamSource) Get_Duration() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Duration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMseStreamSource) Put_Duration(value *IReference[TimeSpan]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Duration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMseStreamSource) AddSourceBuffer(mimeType string) *IMseSourceBuffer {
	var _result *IMseSourceBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddSourceBuffer, uintptr(unsafe.Pointer(this)), NewHStr(mimeType).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMseStreamSource) RemoveSourceBuffer(buffer *IMseSourceBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveSourceBuffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(buffer)))
	_ = _hr
}

func (this *IMseStreamSource) EndOfStream(status MseEndOfStreamStatus) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EndOfStream, uintptr(unsafe.Pointer(this)), uintptr(status))
	_ = _hr
}

// 66F57D37-F9E7-418A-9CDE-A020E956552B
var IID_IMseStreamSource2 = syscall.GUID{0x66F57D37, 0xF9E7, 0x418A,
	[8]byte{0x9C, 0xDE, 0xA0, 0x20, 0xE9, 0x56, 0x55, 0x2B}}

type IMseStreamSource2Interface interface {
	win32.IInspectableInterface
	Get_LiveSeekableRange() *IReference[MseTimeRange]
	Put_LiveSeekableRange(value *IReference[MseTimeRange])
}

type IMseStreamSource2Vtbl struct {
	win32.IInspectableVtbl
	Get_LiveSeekableRange uintptr
	Put_LiveSeekableRange uintptr
}

type IMseStreamSource2 struct {
	win32.IInspectable
}

func (this *IMseStreamSource2) Vtbl() *IMseStreamSource2Vtbl {
	return (*IMseStreamSource2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMseStreamSource2) Get_LiveSeekableRange() *IReference[MseTimeRange] {
	var _result *IReference[MseTimeRange]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LiveSeekableRange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMseStreamSource2) Put_LiveSeekableRange(value *IReference[MseTimeRange]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LiveSeekableRange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 465C679D-D570-43CE-BA21-0BFF5F3FBD0A
var IID_IMseStreamSourceStatics = syscall.GUID{0x465C679D, 0xD570, 0x43CE,
	[8]byte{0xBA, 0x21, 0x0B, 0xFF, 0x5F, 0x3F, 0xBD, 0x0A}}

type IMseStreamSourceStaticsInterface interface {
	win32.IInspectableInterface
	IsContentTypeSupported(contentType string) bool
}

type IMseStreamSourceStaticsVtbl struct {
	win32.IInspectableVtbl
	IsContentTypeSupported uintptr
}

type IMseStreamSourceStatics struct {
	win32.IInspectable
}

func (this *IMseStreamSourceStatics) Vtbl() *IMseStreamSourceStaticsVtbl {
	return (*IMseStreamSourceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMseStreamSourceStatics) IsContentTypeSupported(contentType string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsContentTypeSupported, uintptr(unsafe.Pointer(this)), NewHStr(contentType).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C04BA319-CA41-4813-BFFD-7B08B0ED2557
var IID_ISceneAnalysisEffect = syscall.GUID{0xC04BA319, 0xCA41, 0x4813,
	[8]byte{0xBF, 0xFD, 0x7B, 0x08, 0xB0, 0xED, 0x25, 0x57}}

type ISceneAnalysisEffectInterface interface {
	win32.IInspectableInterface
	Get_HighDynamicRangeAnalyzer() *IHighDynamicRangeControl
	Put_DesiredAnalysisInterval(value TimeSpan)
	Get_DesiredAnalysisInterval() TimeSpan
	Add_SceneAnalyzed(handler TypedEventHandler[*ISceneAnalysisEffect, *ISceneAnalyzedEventArgs]) EventRegistrationToken
	Remove_SceneAnalyzed(cookie EventRegistrationToken)
}

type ISceneAnalysisEffectVtbl struct {
	win32.IInspectableVtbl
	Get_HighDynamicRangeAnalyzer uintptr
	Put_DesiredAnalysisInterval  uintptr
	Get_DesiredAnalysisInterval  uintptr
	Add_SceneAnalyzed            uintptr
	Remove_SceneAnalyzed         uintptr
}

type ISceneAnalysisEffect struct {
	win32.IInspectable
}

func (this *ISceneAnalysisEffect) Vtbl() *ISceneAnalysisEffectVtbl {
	return (*ISceneAnalysisEffectVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISceneAnalysisEffect) Get_HighDynamicRangeAnalyzer() *IHighDynamicRangeControl {
	var _result *IHighDynamicRangeControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HighDynamicRangeAnalyzer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISceneAnalysisEffect) Put_DesiredAnalysisInterval(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DesiredAnalysisInterval, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ISceneAnalysisEffect) Get_DesiredAnalysisInterval() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesiredAnalysisInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISceneAnalysisEffect) Add_SceneAnalyzed(handler TypedEventHandler[*ISceneAnalysisEffect, *ISceneAnalyzedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SceneAnalyzed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISceneAnalysisEffect) Remove_SceneAnalyzed(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SceneAnalyzed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

// D8B10E4C-7FD9-42E1-85EB-6572C297C987
var IID_ISceneAnalysisEffectFrame = syscall.GUID{0xD8B10E4C, 0x7FD9, 0x42E1,
	[8]byte{0x85, 0xEB, 0x65, 0x72, 0xC2, 0x97, 0xC9, 0x87}}

type ISceneAnalysisEffectFrameInterface interface {
	win32.IInspectableInterface
	Get_FrameControlValues() *ICapturedFrameControlValues
	Get_HighDynamicRange() *IHighDynamicRangeOutput
}

type ISceneAnalysisEffectFrameVtbl struct {
	win32.IInspectableVtbl
	Get_FrameControlValues uintptr
	Get_HighDynamicRange   uintptr
}

type ISceneAnalysisEffectFrame struct {
	win32.IInspectable
}

func (this *ISceneAnalysisEffectFrame) Vtbl() *ISceneAnalysisEffectFrameVtbl {
	return (*ISceneAnalysisEffectFrameVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISceneAnalysisEffectFrame) Get_FrameControlValues() *ICapturedFrameControlValues {
	var _result *ICapturedFrameControlValues
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameControlValues, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISceneAnalysisEffectFrame) Get_HighDynamicRange() *IHighDynamicRangeOutput {
	var _result *IHighDynamicRangeOutput
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HighDynamicRange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2D4E29BE-061F-47AE-9915-02524B5F9A5F
var IID_ISceneAnalysisEffectFrame2 = syscall.GUID{0x2D4E29BE, 0x061F, 0x47AE,
	[8]byte{0x99, 0x15, 0x02, 0x52, 0x4B, 0x5F, 0x9A, 0x5F}}

type ISceneAnalysisEffectFrame2Interface interface {
	win32.IInspectableInterface
	Get_AnalysisRecommendation() SceneAnalysisRecommendation
}

type ISceneAnalysisEffectFrame2Vtbl struct {
	win32.IInspectableVtbl
	Get_AnalysisRecommendation uintptr
}

type ISceneAnalysisEffectFrame2 struct {
	win32.IInspectable
}

func (this *ISceneAnalysisEffectFrame2) Vtbl() *ISceneAnalysisEffectFrame2Vtbl {
	return (*ISceneAnalysisEffectFrame2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISceneAnalysisEffectFrame2) Get_AnalysisRecommendation() SceneAnalysisRecommendation {
	var _result SceneAnalysisRecommendation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AnalysisRecommendation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 146B9588-2851-45E4-AD55-44CF8DF8DB4D
var IID_ISceneAnalyzedEventArgs = syscall.GUID{0x146B9588, 0x2851, 0x45E4,
	[8]byte{0xAD, 0x55, 0x44, 0xCF, 0x8D, 0xF8, 0xDB, 0x4D}}

type ISceneAnalyzedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ResultFrame() *ISceneAnalysisEffectFrame
}

type ISceneAnalyzedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ResultFrame uintptr
}

type ISceneAnalyzedEventArgs struct {
	win32.IInspectable
}

func (this *ISceneAnalyzedEventArgs) Vtbl() *ISceneAnalyzedEventArgsVtbl {
	return (*ISceneAnalyzedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISceneAnalyzedEventArgs) Get_ResultFrame() *ISceneAnalysisEffectFrame {
	var _result *ISceneAnalysisEffectFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResultFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 77206F1F-C34F-494F-8077-2BAD9FF4ECF1
var IID_ISingleSelectMediaTrackList = syscall.GUID{0x77206F1F, 0xC34F, 0x494F,
	[8]byte{0x80, 0x77, 0x2B, 0xAD, 0x9F, 0xF4, 0xEC, 0xF1}}

type ISingleSelectMediaTrackListInterface interface {
	win32.IInspectableInterface
	Add_SelectedIndexChanged(handler TypedEventHandler[*ISingleSelectMediaTrackList, interface{}]) EventRegistrationToken
	Remove_SelectedIndexChanged(token EventRegistrationToken)
	Put_SelectedIndex(value int32)
	Get_SelectedIndex() int32
}

type ISingleSelectMediaTrackListVtbl struct {
	win32.IInspectableVtbl
	Add_SelectedIndexChanged    uintptr
	Remove_SelectedIndexChanged uintptr
	Put_SelectedIndex           uintptr
	Get_SelectedIndex           uintptr
}

type ISingleSelectMediaTrackList struct {
	win32.IInspectable
}

func (this *ISingleSelectMediaTrackList) Vtbl() *ISingleSelectMediaTrackListVtbl {
	return (*ISingleSelectMediaTrackListVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISingleSelectMediaTrackList) Add_SelectedIndexChanged(handler TypedEventHandler[*ISingleSelectMediaTrackList, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SelectedIndexChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISingleSelectMediaTrackList) Remove_SelectedIndexChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SelectedIndexChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISingleSelectMediaTrackList) Put_SelectedIndex(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SelectedIndex, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISingleSelectMediaTrackList) Get_SelectedIndex() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SelectedIndex, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// AEE254DC-1725-4BAD-8043-A98499B017A2
var IID_ISpeechCue = syscall.GUID{0xAEE254DC, 0x1725, 0x4BAD,
	[8]byte{0x80, 0x43, 0xA9, 0x84, 0x99, 0xB0, 0x17, 0xA2}}

type ISpeechCueInterface interface {
	win32.IInspectableInterface
	Get_Text() string
	Put_Text(value string)
	Get_StartPositionInInput() *IReference[int32]
	Put_StartPositionInInput(value *IReference[int32])
	Get_EndPositionInInput() *IReference[int32]
	Put_EndPositionInInput(value *IReference[int32])
}

type ISpeechCueVtbl struct {
	win32.IInspectableVtbl
	Get_Text                 uintptr
	Put_Text                 uintptr
	Get_StartPositionInInput uintptr
	Put_StartPositionInInput uintptr
	Get_EndPositionInInput   uintptr
	Put_EndPositionInInput   uintptr
}

type ISpeechCue struct {
	win32.IInspectable
}

func (this *ISpeechCue) Vtbl() *ISpeechCueVtbl {
	return (*ISpeechCueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechCue) Get_Text() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Text, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISpeechCue) Put_Text(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Text, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISpeechCue) Get_StartPositionInInput() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StartPositionInInput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpeechCue) Put_StartPositionInInput(value *IReference[int32]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StartPositionInInput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ISpeechCue) Get_EndPositionInInput() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EndPositionInInput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpeechCue) Put_EndPositionInInput(value *IReference[int32]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_EndPositionInInput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 133336BF-296A-463E-9FF9-01CD25691408
var IID_ITimedMetadataStreamDescriptor = syscall.GUID{0x133336BF, 0x296A, 0x463E,
	[8]byte{0x9F, 0xF9, 0x01, 0xCD, 0x25, 0x69, 0x14, 0x08}}

type ITimedMetadataStreamDescriptorInterface interface {
	win32.IInspectableInterface
	Get_EncodingProperties() *IMediaEncodingProperties
	Copy() *IMediaStreamDescriptor
}

type ITimedMetadataStreamDescriptorVtbl struct {
	win32.IInspectableVtbl
	Get_EncodingProperties uintptr
	Copy                   uintptr
}

type ITimedMetadataStreamDescriptor struct {
	win32.IInspectable
}

func (this *ITimedMetadataStreamDescriptor) Vtbl() *ITimedMetadataStreamDescriptorVtbl {
	return (*ITimedMetadataStreamDescriptorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimedMetadataStreamDescriptor) Get_EncodingProperties() *IMediaEncodingProperties {
	var _result *IMediaEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EncodingProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITimedMetadataStreamDescriptor) Copy() *IMediaStreamDescriptor {
	var _result *IMediaStreamDescriptor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Copy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C027DE30-7362-4FF9-98B1-2DFD0B8D1CAE
var IID_ITimedMetadataStreamDescriptorFactory = syscall.GUID{0xC027DE30, 0x7362, 0x4FF9,
	[8]byte{0x98, 0xB1, 0x2D, 0xFD, 0x0B, 0x8D, 0x1C, 0xAE}}

type ITimedMetadataStreamDescriptorFactoryInterface interface {
	win32.IInspectableInterface
	Create(encodingProperties *IMediaEncodingProperties) *IMediaStreamDescriptor
}

type ITimedMetadataStreamDescriptorFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ITimedMetadataStreamDescriptorFactory struct {
	win32.IInspectable
}

func (this *ITimedMetadataStreamDescriptorFactory) Vtbl() *ITimedMetadataStreamDescriptorFactoryVtbl {
	return (*ITimedMetadataStreamDescriptorFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimedMetadataStreamDescriptorFactory) Create(encodingProperties *IMediaEncodingProperties) *IMediaStreamDescriptor {
	var _result *IMediaStreamDescriptor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(encodingProperties)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9E6AED9E-F67A-49A9-B330-CF03B0E9CF07
var IID_ITimedMetadataTrack = syscall.GUID{0x9E6AED9E, 0xF67A, 0x49A9,
	[8]byte{0xB3, 0x30, 0xCF, 0x03, 0xB0, 0xE9, 0xCF, 0x07}}

type ITimedMetadataTrackInterface interface {
	win32.IInspectableInterface
	Add_CueEntered(handler TypedEventHandler[*ITimedMetadataTrack, *IMediaCueEventArgs]) EventRegistrationToken
	Remove_CueEntered(token EventRegistrationToken)
	Add_CueExited(handler TypedEventHandler[*ITimedMetadataTrack, *IMediaCueEventArgs]) EventRegistrationToken
	Remove_CueExited(token EventRegistrationToken)
	Add_TrackFailed(handler TypedEventHandler[*ITimedMetadataTrack, *ITimedMetadataTrackFailedEventArgs]) EventRegistrationToken
	Remove_TrackFailed(token EventRegistrationToken)
	Get_Cues() *IVectorView[*IMediaCue]
	Get_ActiveCues() *IVectorView[*IMediaCue]
	Get_TimedMetadataKind() TimedMetadataKind
	Get_DispatchType() string
	AddCue(cue *IMediaCue)
	RemoveCue(cue *IMediaCue)
}

type ITimedMetadataTrackVtbl struct {
	win32.IInspectableVtbl
	Add_CueEntered        uintptr
	Remove_CueEntered     uintptr
	Add_CueExited         uintptr
	Remove_CueExited      uintptr
	Add_TrackFailed       uintptr
	Remove_TrackFailed    uintptr
	Get_Cues              uintptr
	Get_ActiveCues        uintptr
	Get_TimedMetadataKind uintptr
	Get_DispatchType      uintptr
	AddCue                uintptr
	RemoveCue             uintptr
}

type ITimedMetadataTrack struct {
	win32.IInspectable
}

func (this *ITimedMetadataTrack) Vtbl() *ITimedMetadataTrackVtbl {
	return (*ITimedMetadataTrackVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimedMetadataTrack) Add_CueEntered(handler TypedEventHandler[*ITimedMetadataTrack, *IMediaCueEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_CueEntered, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedMetadataTrack) Remove_CueEntered(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_CueEntered, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ITimedMetadataTrack) Add_CueExited(handler TypedEventHandler[*ITimedMetadataTrack, *IMediaCueEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_CueExited, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedMetadataTrack) Remove_CueExited(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_CueExited, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ITimedMetadataTrack) Add_TrackFailed(handler TypedEventHandler[*ITimedMetadataTrack, *ITimedMetadataTrackFailedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_TrackFailed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedMetadataTrack) Remove_TrackFailed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_TrackFailed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ITimedMetadataTrack) Get_Cues() *IVectorView[*IMediaCue] {
	var _result *IVectorView[*IMediaCue]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Cues, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITimedMetadataTrack) Get_ActiveCues() *IVectorView[*IMediaCue] {
	var _result *IVectorView[*IMediaCue]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ActiveCues, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITimedMetadataTrack) Get_TimedMetadataKind() TimedMetadataKind {
	var _result TimedMetadataKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TimedMetadataKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedMetadataTrack) Get_DispatchType() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DispatchType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ITimedMetadataTrack) AddCue(cue *IMediaCue) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddCue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(cue)))
	_ = _hr
}

func (this *ITimedMetadataTrack) RemoveCue(cue *IMediaCue) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveCue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(cue)))
	_ = _hr
}

// 21B4B648-9F9D-40BA-A8F3-1A92753AEF0B
var IID_ITimedMetadataTrack2 = syscall.GUID{0x21B4B648, 0x9F9D, 0x40BA,
	[8]byte{0xA8, 0xF3, 0x1A, 0x92, 0x75, 0x3A, 0xEF, 0x0B}}

type ITimedMetadataTrack2Interface interface {
	win32.IInspectableInterface
	Get_PlaybackItem() *IMediaPlaybackItem
	Get_Name() string
}

type ITimedMetadataTrack2Vtbl struct {
	win32.IInspectableVtbl
	Get_PlaybackItem uintptr
	Get_Name         uintptr
}

type ITimedMetadataTrack2 struct {
	win32.IInspectable
}

func (this *ITimedMetadataTrack2) Vtbl() *ITimedMetadataTrack2Vtbl {
	return (*ITimedMetadataTrack2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimedMetadataTrack2) Get_PlaybackItem() *IMediaPlaybackItem {
	var _result *IMediaPlaybackItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlaybackItem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITimedMetadataTrack2) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// B3767915-4114-4819-B9D9-DD76089E72F8
var IID_ITimedMetadataTrackError = syscall.GUID{0xB3767915, 0x4114, 0x4819,
	[8]byte{0xB9, 0xD9, 0xDD, 0x76, 0x08, 0x9E, 0x72, 0xF8}}

type ITimedMetadataTrackErrorInterface interface {
	win32.IInspectableInterface
	Get_ErrorCode() TimedMetadataTrackErrorCode
	Get_ExtendedError() HResult
}

type ITimedMetadataTrackErrorVtbl struct {
	win32.IInspectableVtbl
	Get_ErrorCode     uintptr
	Get_ExtendedError uintptr
}

type ITimedMetadataTrackError struct {
	win32.IInspectable
}

func (this *ITimedMetadataTrackError) Vtbl() *ITimedMetadataTrackErrorVtbl {
	return (*ITimedMetadataTrackErrorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimedMetadataTrackError) Get_ErrorCode() TimedMetadataTrackErrorCode {
	var _result TimedMetadataTrackErrorCode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ErrorCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedMetadataTrackError) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 8DD57611-97B3-4E1F-852C-0F482C81AD26
var IID_ITimedMetadataTrackFactory = syscall.GUID{0x8DD57611, 0x97B3, 0x4E1F,
	[8]byte{0x85, 0x2C, 0x0F, 0x48, 0x2C, 0x81, 0xAD, 0x26}}

type ITimedMetadataTrackFactoryInterface interface {
	win32.IInspectableInterface
	Create(id string, language string, kind TimedMetadataKind) *ITimedMetadataTrack
}

type ITimedMetadataTrackFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ITimedMetadataTrackFactory struct {
	win32.IInspectable
}

func (this *ITimedMetadataTrackFactory) Vtbl() *ITimedMetadataTrackFactoryVtbl {
	return (*ITimedMetadataTrackFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimedMetadataTrackFactory) Create(id string, language string, kind TimedMetadataKind) *ITimedMetadataTrack {
	var _result *ITimedMetadataTrack
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(id).Ptr, NewHStr(language).Ptr, uintptr(kind), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A57FC9D1-6789-4D4D-B07F-84B4F31ACB70
var IID_ITimedMetadataTrackFailedEventArgs = syscall.GUID{0xA57FC9D1, 0x6789, 0x4D4D,
	[8]byte{0xB0, 0x7F, 0x84, 0xB4, 0xF3, 0x1A, 0xCB, 0x70}}

type ITimedMetadataTrackFailedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Error() *ITimedMetadataTrackError
}

type ITimedMetadataTrackFailedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Error uintptr
}

type ITimedMetadataTrackFailedEventArgs struct {
	win32.IInspectable
}

func (this *ITimedMetadataTrackFailedEventArgs) Vtbl() *ITimedMetadataTrackFailedEventArgsVtbl {
	return (*ITimedMetadataTrackFailedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimedMetadataTrackFailedEventArgs) Get_Error() *ITimedMetadataTrackError {
	var _result *ITimedMetadataTrackError
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Error, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3B7F2024-F74E-4ADE-93C5-219DA05B6856
var IID_ITimedMetadataTrackProvider = syscall.GUID{0x3B7F2024, 0xF74E, 0x4ADE,
	[8]byte{0x93, 0xC5, 0x21, 0x9D, 0xA0, 0x5B, 0x68, 0x56}}

type ITimedMetadataTrackProviderInterface interface {
	win32.IInspectableInterface
	Get_TimedMetadataTracks() *IVectorView[*ITimedMetadataTrack]
}

type ITimedMetadataTrackProviderVtbl struct {
	win32.IInspectableVtbl
	Get_TimedMetadataTracks uintptr
}

type ITimedMetadataTrackProvider struct {
	win32.IInspectable
}

func (this *ITimedMetadataTrackProvider) Vtbl() *ITimedMetadataTrackProviderVtbl {
	return (*ITimedMetadataTrackProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimedMetadataTrackProvider) Get_TimedMetadataTracks() *IVectorView[*ITimedMetadataTrack] {
	var _result *IVectorView[*ITimedMetadataTrack]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TimedMetadataTracks, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D9062783-5597-5092-820C-8F738E0F774A
var IID_ITimedTextBouten = syscall.GUID{0xD9062783, 0x5597, 0x5092,
	[8]byte{0x82, 0x0C, 0x8F, 0x73, 0x8E, 0x0F, 0x77, 0x4A}}

type ITimedTextBoutenInterface interface {
	win32.IInspectableInterface
	Get_Type() TimedTextBoutenType
	Put_Type(value TimedTextBoutenType)
	Get_Color() unsafe.Pointer
	Put_Color(value unsafe.Pointer)
	Get_Position() TimedTextBoutenPosition
	Put_Position(value TimedTextBoutenPosition)
}

type ITimedTextBoutenVtbl struct {
	win32.IInspectableVtbl
	Get_Type     uintptr
	Put_Type     uintptr
	Get_Color    uintptr
	Put_Color    uintptr
	Get_Position uintptr
	Put_Position uintptr
}

type ITimedTextBouten struct {
	win32.IInspectable
}

func (this *ITimedTextBouten) Vtbl() *ITimedTextBoutenVtbl {
	return (*ITimedTextBoutenVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimedTextBouten) Get_Type() TimedTextBoutenType {
	var _result TimedTextBoutenType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Type, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextBouten) Put_Type(value TimedTextBoutenType) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Type, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ITimedTextBouten) Get_Color() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Color, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextBouten) Put_Color(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Color, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ITimedTextBouten) Get_Position() TimedTextBoutenPosition {
	var _result TimedTextBoutenPosition
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextBouten) Put_Position(value TimedTextBoutenPosition) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Position, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 51C79E51-3B86-494D-B359-BB2EA7ACA9A9
var IID_ITimedTextCue = syscall.GUID{0x51C79E51, 0x3B86, 0x494D,
	[8]byte{0xB3, 0x59, 0xBB, 0x2E, 0xA7, 0xAC, 0xA9, 0xA9}}

type ITimedTextCueInterface interface {
	win32.IInspectableInterface
	Get_CueRegion() *ITimedTextRegion
	Put_CueRegion(value *ITimedTextRegion)
	Get_CueStyle() *ITimedTextStyle
	Put_CueStyle(value *ITimedTextStyle)
	Get_Lines() *IVector[*ITimedTextLine]
}

type ITimedTextCueVtbl struct {
	win32.IInspectableVtbl
	Get_CueRegion uintptr
	Put_CueRegion uintptr
	Get_CueStyle  uintptr
	Put_CueStyle  uintptr
	Get_Lines     uintptr
}

type ITimedTextCue struct {
	win32.IInspectable
}

func (this *ITimedTextCue) Vtbl() *ITimedTextCueVtbl {
	return (*ITimedTextCueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimedTextCue) Get_CueRegion() *ITimedTextRegion {
	var _result *ITimedTextRegion
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CueRegion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITimedTextCue) Put_CueRegion(value *ITimedTextRegion) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CueRegion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ITimedTextCue) Get_CueStyle() *ITimedTextStyle {
	var _result *ITimedTextStyle
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CueStyle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITimedTextCue) Put_CueStyle(value *ITimedTextStyle) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CueStyle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ITimedTextCue) Get_Lines() *IVector[*ITimedTextLine] {
	var _result *IVector[*ITimedTextLine]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Lines, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 978D7CE2-7308-4C66-BE50-65777289F5DF
var IID_ITimedTextLine = syscall.GUID{0x978D7CE2, 0x7308, 0x4C66,
	[8]byte{0xBE, 0x50, 0x65, 0x77, 0x72, 0x89, 0xF5, 0xDF}}

type ITimedTextLineInterface interface {
	win32.IInspectableInterface
	Get_Text() string
	Put_Text(value string)
	Get_Subformats() *IVector[*ITimedTextSubformat]
}

type ITimedTextLineVtbl struct {
	win32.IInspectableVtbl
	Get_Text       uintptr
	Put_Text       uintptr
	Get_Subformats uintptr
}

type ITimedTextLine struct {
	win32.IInspectable
}

func (this *ITimedTextLine) Vtbl() *ITimedTextLineVtbl {
	return (*ITimedTextLineVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimedTextLine) Get_Text() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Text, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ITimedTextLine) Put_Text(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Text, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ITimedTextLine) Get_Subformats() *IVector[*ITimedTextSubformat] {
	var _result *IVector[*ITimedTextSubformat]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Subformats, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1ED0881F-8A06-4222-9F59-B21BF40124B4
var IID_ITimedTextRegion = syscall.GUID{0x1ED0881F, 0x8A06, 0x4222,
	[8]byte{0x9F, 0x59, 0xB2, 0x1B, 0xF4, 0x01, 0x24, 0xB4}}

type ITimedTextRegionInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Put_Name(value string)
	Get_Position() TimedTextPoint
	Put_Position(value TimedTextPoint)
	Get_Extent() TimedTextSize
	Put_Extent(value TimedTextSize)
	Get_Background() unsafe.Pointer
	Put_Background(value unsafe.Pointer)
	Get_WritingMode() TimedTextWritingMode
	Put_WritingMode(value TimedTextWritingMode)
	Get_DisplayAlignment() TimedTextDisplayAlignment
	Put_DisplayAlignment(value TimedTextDisplayAlignment)
	Get_LineHeight() TimedTextDouble
	Put_LineHeight(value TimedTextDouble)
	Get_IsOverflowClipped() bool
	Put_IsOverflowClipped(value bool)
	Get_Padding() TimedTextPadding
	Put_Padding(value TimedTextPadding)
	Get_TextWrapping() TimedTextWrapping
	Put_TextWrapping(value TimedTextWrapping)
	Get_ZIndex() int32
	Put_ZIndex(value int32)
	Get_ScrollMode() TimedTextScrollMode
	Put_ScrollMode(value TimedTextScrollMode)
}

type ITimedTextRegionVtbl struct {
	win32.IInspectableVtbl
	Get_Name              uintptr
	Put_Name              uintptr
	Get_Position          uintptr
	Put_Position          uintptr
	Get_Extent            uintptr
	Put_Extent            uintptr
	Get_Background        uintptr
	Put_Background        uintptr
	Get_WritingMode       uintptr
	Put_WritingMode       uintptr
	Get_DisplayAlignment  uintptr
	Put_DisplayAlignment  uintptr
	Get_LineHeight        uintptr
	Put_LineHeight        uintptr
	Get_IsOverflowClipped uintptr
	Put_IsOverflowClipped uintptr
	Get_Padding           uintptr
	Put_Padding           uintptr
	Get_TextWrapping      uintptr
	Put_TextWrapping      uintptr
	Get_ZIndex            uintptr
	Put_ZIndex            uintptr
	Get_ScrollMode        uintptr
	Put_ScrollMode        uintptr
}

type ITimedTextRegion struct {
	win32.IInspectable
}

func (this *ITimedTextRegion) Vtbl() *ITimedTextRegionVtbl {
	return (*ITimedTextRegionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimedTextRegion) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ITimedTextRegion) Put_Name(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Name, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ITimedTextRegion) Get_Position() TimedTextPoint {
	var _result TimedTextPoint
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextRegion) Put_Position(value TimedTextPoint) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ITimedTextRegion) Get_Extent() TimedTextSize {
	var _result TimedTextSize
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Extent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextRegion) Put_Extent(value TimedTextSize) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Extent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ITimedTextRegion) Get_Background() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Background, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextRegion) Put_Background(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Background, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ITimedTextRegion) Get_WritingMode() TimedTextWritingMode {
	var _result TimedTextWritingMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WritingMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextRegion) Put_WritingMode(value TimedTextWritingMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_WritingMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ITimedTextRegion) Get_DisplayAlignment() TimedTextDisplayAlignment {
	var _result TimedTextDisplayAlignment
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayAlignment, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextRegion) Put_DisplayAlignment(value TimedTextDisplayAlignment) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DisplayAlignment, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ITimedTextRegion) Get_LineHeight() TimedTextDouble {
	var _result TimedTextDouble
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LineHeight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextRegion) Put_LineHeight(value TimedTextDouble) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LineHeight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ITimedTextRegion) Get_IsOverflowClipped() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsOverflowClipped, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextRegion) Put_IsOverflowClipped(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsOverflowClipped, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ITimedTextRegion) Get_Padding() TimedTextPadding {
	var _result TimedTextPadding
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Padding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextRegion) Put_Padding(value TimedTextPadding) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Padding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ITimedTextRegion) Get_TextWrapping() TimedTextWrapping {
	var _result TimedTextWrapping
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TextWrapping, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextRegion) Put_TextWrapping(value TimedTextWrapping) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TextWrapping, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ITimedTextRegion) Get_ZIndex() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ZIndex, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextRegion) Put_ZIndex(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ZIndex, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ITimedTextRegion) Get_ScrollMode() TimedTextScrollMode {
	var _result TimedTextScrollMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScrollMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextRegion) Put_ScrollMode(value TimedTextScrollMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ScrollMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 10335C29-5B3C-5693-9959-D05A0BD24628
var IID_ITimedTextRuby = syscall.GUID{0x10335C29, 0x5B3C, 0x5693,
	[8]byte{0x99, 0x59, 0xD0, 0x5A, 0x0B, 0xD2, 0x46, 0x28}}

type ITimedTextRubyInterface interface {
	win32.IInspectableInterface
	Get_Text() string
	Put_Text(value string)
	Get_Position() TimedTextRubyPosition
	Put_Position(value TimedTextRubyPosition)
	Get_Align() TimedTextRubyAlign
	Put_Align(value TimedTextRubyAlign)
	Get_Reserve() TimedTextRubyReserve
	Put_Reserve(value TimedTextRubyReserve)
}

type ITimedTextRubyVtbl struct {
	win32.IInspectableVtbl
	Get_Text     uintptr
	Put_Text     uintptr
	Get_Position uintptr
	Put_Position uintptr
	Get_Align    uintptr
	Put_Align    uintptr
	Get_Reserve  uintptr
	Put_Reserve  uintptr
}

type ITimedTextRuby struct {
	win32.IInspectable
}

func (this *ITimedTextRuby) Vtbl() *ITimedTextRubyVtbl {
	return (*ITimedTextRubyVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimedTextRuby) Get_Text() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Text, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ITimedTextRuby) Put_Text(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Text, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ITimedTextRuby) Get_Position() TimedTextRubyPosition {
	var _result TimedTextRubyPosition
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextRuby) Put_Position(value TimedTextRubyPosition) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Position, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ITimedTextRuby) Get_Align() TimedTextRubyAlign {
	var _result TimedTextRubyAlign
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Align, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextRuby) Put_Align(value TimedTextRubyAlign) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Align, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ITimedTextRuby) Get_Reserve() TimedTextRubyReserve {
	var _result TimedTextRubyReserve
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Reserve, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextRuby) Put_Reserve(value TimedTextRubyReserve) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Reserve, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// C4ED9BA6-101F-404D-A949-82F33FCD93B7
var IID_ITimedTextSource = syscall.GUID{0xC4ED9BA6, 0x101F, 0x404D,
	[8]byte{0xA9, 0x49, 0x82, 0xF3, 0x3F, 0xCD, 0x93, 0xB7}}

type ITimedTextSourceInterface interface {
	win32.IInspectableInterface
	Add_Resolved(handler TypedEventHandler[*ITimedTextSource, *ITimedTextSourceResolveResultEventArgs]) EventRegistrationToken
	Remove_Resolved(token EventRegistrationToken)
}

type ITimedTextSourceVtbl struct {
	win32.IInspectableVtbl
	Add_Resolved    uintptr
	Remove_Resolved uintptr
}

type ITimedTextSource struct {
	win32.IInspectable
}

func (this *ITimedTextSource) Vtbl() *ITimedTextSourceVtbl {
	return (*ITimedTextSourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimedTextSource) Add_Resolved(handler TypedEventHandler[*ITimedTextSource, *ITimedTextSourceResolveResultEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Resolved, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextSource) Remove_Resolved(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Resolved, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 48907C9C-DCD8-4C33-9AD3-6CDCE7B1C566
var IID_ITimedTextSourceResolveResultEventArgs = syscall.GUID{0x48907C9C, 0xDCD8, 0x4C33,
	[8]byte{0x9A, 0xD3, 0x6C, 0xDC, 0xE7, 0xB1, 0xC5, 0x66}}

type ITimedTextSourceResolveResultEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Error() *ITimedMetadataTrackError
	Get_Tracks() *IVectorView[*ITimedMetadataTrack]
}

type ITimedTextSourceResolveResultEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Error  uintptr
	Get_Tracks uintptr
}

type ITimedTextSourceResolveResultEventArgs struct {
	win32.IInspectable
}

func (this *ITimedTextSourceResolveResultEventArgs) Vtbl() *ITimedTextSourceResolveResultEventArgsVtbl {
	return (*ITimedTextSourceResolveResultEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimedTextSourceResolveResultEventArgs) Get_Error() *ITimedMetadataTrackError {
	var _result *ITimedMetadataTrackError
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Error, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITimedTextSourceResolveResultEventArgs) Get_Tracks() *IVectorView[*ITimedMetadataTrack] {
	var _result *IVectorView[*ITimedMetadataTrack]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Tracks, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7E311853-9ABA-4AC4-BB98-2FB176C3BFDD
var IID_ITimedTextSourceStatics = syscall.GUID{0x7E311853, 0x9ABA, 0x4AC4,
	[8]byte{0xBB, 0x98, 0x2F, 0xB1, 0x76, 0xC3, 0xBF, 0xDD}}

type ITimedTextSourceStaticsInterface interface {
	win32.IInspectableInterface
	CreateFromStream(stream *IRandomAccessStream) *ITimedTextSource
	CreateFromUri(uri *IUriRuntimeClass) *ITimedTextSource
	CreateFromStreamWithLanguage(stream *IRandomAccessStream, defaultLanguage string) *ITimedTextSource
	CreateFromUriWithLanguage(uri *IUriRuntimeClass, defaultLanguage string) *ITimedTextSource
}

type ITimedTextSourceStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateFromStream             uintptr
	CreateFromUri                uintptr
	CreateFromStreamWithLanguage uintptr
	CreateFromUriWithLanguage    uintptr
}

type ITimedTextSourceStatics struct {
	win32.IInspectable
}

func (this *ITimedTextSourceStatics) Vtbl() *ITimedTextSourceStaticsVtbl {
	return (*ITimedTextSourceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimedTextSourceStatics) CreateFromStream(stream *IRandomAccessStream) *ITimedTextSource {
	var _result *ITimedTextSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITimedTextSourceStatics) CreateFromUri(uri *IUriRuntimeClass) *ITimedTextSource {
	var _result *ITimedTextSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITimedTextSourceStatics) CreateFromStreamWithLanguage(stream *IRandomAccessStream, defaultLanguage string) *ITimedTextSource {
	var _result *ITimedTextSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromStreamWithLanguage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)), NewHStr(defaultLanguage).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITimedTextSourceStatics) CreateFromUriWithLanguage(uri *IUriRuntimeClass, defaultLanguage string) *ITimedTextSource {
	var _result *ITimedTextSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromUriWithLanguage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), NewHStr(defaultLanguage).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B66B7602-923E-43FA-9633-587075812DB5
var IID_ITimedTextSourceStatics2 = syscall.GUID{0xB66B7602, 0x923E, 0x43FA,
	[8]byte{0x96, 0x33, 0x58, 0x70, 0x75, 0x81, 0x2D, 0xB5}}

type ITimedTextSourceStatics2Interface interface {
	win32.IInspectableInterface
	CreateFromStreamWithIndex(stream *IRandomAccessStream, indexStream *IRandomAccessStream) *ITimedTextSource
	CreateFromUriWithIndex(uri *IUriRuntimeClass, indexUri *IUriRuntimeClass) *ITimedTextSource
	CreateFromStreamWithIndexAndLanguage(stream *IRandomAccessStream, indexStream *IRandomAccessStream, defaultLanguage string) *ITimedTextSource
	CreateFromUriWithIndexAndLanguage(uri *IUriRuntimeClass, indexUri *IUriRuntimeClass, defaultLanguage string) *ITimedTextSource
}

type ITimedTextSourceStatics2Vtbl struct {
	win32.IInspectableVtbl
	CreateFromStreamWithIndex            uintptr
	CreateFromUriWithIndex               uintptr
	CreateFromStreamWithIndexAndLanguage uintptr
	CreateFromUriWithIndexAndLanguage    uintptr
}

type ITimedTextSourceStatics2 struct {
	win32.IInspectable
}

func (this *ITimedTextSourceStatics2) Vtbl() *ITimedTextSourceStatics2Vtbl {
	return (*ITimedTextSourceStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimedTextSourceStatics2) CreateFromStreamWithIndex(stream *IRandomAccessStream, indexStream *IRandomAccessStream) *ITimedTextSource {
	var _result *ITimedTextSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromStreamWithIndex, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(indexStream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITimedTextSourceStatics2) CreateFromUriWithIndex(uri *IUriRuntimeClass, indexUri *IUriRuntimeClass) *ITimedTextSource {
	var _result *ITimedTextSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromUriWithIndex, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(indexUri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITimedTextSourceStatics2) CreateFromStreamWithIndexAndLanguage(stream *IRandomAccessStream, indexStream *IRandomAccessStream, defaultLanguage string) *ITimedTextSource {
	var _result *ITimedTextSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromStreamWithIndexAndLanguage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(indexStream)), NewHStr(defaultLanguage).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITimedTextSourceStatics2) CreateFromUriWithIndexAndLanguage(uri *IUriRuntimeClass, indexUri *IUriRuntimeClass, defaultLanguage string) *ITimedTextSource {
	var _result *ITimedTextSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromUriWithIndexAndLanguage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(indexUri)), NewHStr(defaultLanguage).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1BB2384D-A825-40C2-A7F5-281EAEDF3B55
var IID_ITimedTextStyle = syscall.GUID{0x1BB2384D, 0xA825, 0x40C2,
	[8]byte{0xA7, 0xF5, 0x28, 0x1E, 0xAE, 0xDF, 0x3B, 0x55}}

type ITimedTextStyleInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Put_Name(value string)
	Get_FontFamily() string
	Put_FontFamily(value string)
	Get_FontSize() TimedTextDouble
	Put_FontSize(value TimedTextDouble)
	Get_FontWeight() TimedTextWeight
	Put_FontWeight(value TimedTextWeight)
	Get_Foreground() unsafe.Pointer
	Put_Foreground(value unsafe.Pointer)
	Get_Background() unsafe.Pointer
	Put_Background(value unsafe.Pointer)
	Get_IsBackgroundAlwaysShown() bool
	Put_IsBackgroundAlwaysShown(value bool)
	Get_FlowDirection() TimedTextFlowDirection
	Put_FlowDirection(value TimedTextFlowDirection)
	Get_LineAlignment() TimedTextLineAlignment
	Put_LineAlignment(value TimedTextLineAlignment)
	Get_OutlineColor() unsafe.Pointer
	Put_OutlineColor(value unsafe.Pointer)
	Get_OutlineThickness() TimedTextDouble
	Put_OutlineThickness(value TimedTextDouble)
	Get_OutlineRadius() TimedTextDouble
	Put_OutlineRadius(value TimedTextDouble)
}

type ITimedTextStyleVtbl struct {
	win32.IInspectableVtbl
	Get_Name                    uintptr
	Put_Name                    uintptr
	Get_FontFamily              uintptr
	Put_FontFamily              uintptr
	Get_FontSize                uintptr
	Put_FontSize                uintptr
	Get_FontWeight              uintptr
	Put_FontWeight              uintptr
	Get_Foreground              uintptr
	Put_Foreground              uintptr
	Get_Background              uintptr
	Put_Background              uintptr
	Get_IsBackgroundAlwaysShown uintptr
	Put_IsBackgroundAlwaysShown uintptr
	Get_FlowDirection           uintptr
	Put_FlowDirection           uintptr
	Get_LineAlignment           uintptr
	Put_LineAlignment           uintptr
	Get_OutlineColor            uintptr
	Put_OutlineColor            uintptr
	Get_OutlineThickness        uintptr
	Put_OutlineThickness        uintptr
	Get_OutlineRadius           uintptr
	Put_OutlineRadius           uintptr
}

type ITimedTextStyle struct {
	win32.IInspectable
}

func (this *ITimedTextStyle) Vtbl() *ITimedTextStyleVtbl {
	return (*ITimedTextStyleVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimedTextStyle) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ITimedTextStyle) Put_Name(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Name, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ITimedTextStyle) Get_FontFamily() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FontFamily, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ITimedTextStyle) Put_FontFamily(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FontFamily, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ITimedTextStyle) Get_FontSize() TimedTextDouble {
	var _result TimedTextDouble
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FontSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextStyle) Put_FontSize(value TimedTextDouble) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FontSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ITimedTextStyle) Get_FontWeight() TimedTextWeight {
	var _result TimedTextWeight
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FontWeight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextStyle) Put_FontWeight(value TimedTextWeight) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FontWeight, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ITimedTextStyle) Get_Foreground() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Foreground, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextStyle) Put_Foreground(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Foreground, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ITimedTextStyle) Get_Background() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Background, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextStyle) Put_Background(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Background, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ITimedTextStyle) Get_IsBackgroundAlwaysShown() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsBackgroundAlwaysShown, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextStyle) Put_IsBackgroundAlwaysShown(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsBackgroundAlwaysShown, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ITimedTextStyle) Get_FlowDirection() TimedTextFlowDirection {
	var _result TimedTextFlowDirection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FlowDirection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextStyle) Put_FlowDirection(value TimedTextFlowDirection) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FlowDirection, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ITimedTextStyle) Get_LineAlignment() TimedTextLineAlignment {
	var _result TimedTextLineAlignment
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LineAlignment, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextStyle) Put_LineAlignment(value TimedTextLineAlignment) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LineAlignment, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ITimedTextStyle) Get_OutlineColor() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutlineColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextStyle) Put_OutlineColor(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_OutlineColor, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ITimedTextStyle) Get_OutlineThickness() TimedTextDouble {
	var _result TimedTextDouble
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutlineThickness, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextStyle) Put_OutlineThickness(value TimedTextDouble) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_OutlineThickness, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ITimedTextStyle) Get_OutlineRadius() TimedTextDouble {
	var _result TimedTextDouble
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutlineRadius, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextStyle) Put_OutlineRadius(value TimedTextDouble) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_OutlineRadius, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

// 655F492D-6111-4787-89CC-686FECE57E14
var IID_ITimedTextStyle2 = syscall.GUID{0x655F492D, 0x6111, 0x4787,
	[8]byte{0x89, 0xCC, 0x68, 0x6F, 0xEC, 0xE5, 0x7E, 0x14}}

type ITimedTextStyle2Interface interface {
	win32.IInspectableInterface
	Get_FontStyle() TimedTextFontStyle
	Put_FontStyle(value TimedTextFontStyle)
	Get_IsUnderlineEnabled() bool
	Put_IsUnderlineEnabled(value bool)
	Get_IsLineThroughEnabled() bool
	Put_IsLineThroughEnabled(value bool)
	Get_IsOverlineEnabled() bool
	Put_IsOverlineEnabled(value bool)
}

type ITimedTextStyle2Vtbl struct {
	win32.IInspectableVtbl
	Get_FontStyle            uintptr
	Put_FontStyle            uintptr
	Get_IsUnderlineEnabled   uintptr
	Put_IsUnderlineEnabled   uintptr
	Get_IsLineThroughEnabled uintptr
	Put_IsLineThroughEnabled uintptr
	Get_IsOverlineEnabled    uintptr
	Put_IsOverlineEnabled    uintptr
}

type ITimedTextStyle2 struct {
	win32.IInspectable
}

func (this *ITimedTextStyle2) Vtbl() *ITimedTextStyle2Vtbl {
	return (*ITimedTextStyle2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimedTextStyle2) Get_FontStyle() TimedTextFontStyle {
	var _result TimedTextFontStyle
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FontStyle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextStyle2) Put_FontStyle(value TimedTextFontStyle) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FontStyle, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ITimedTextStyle2) Get_IsUnderlineEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsUnderlineEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextStyle2) Put_IsUnderlineEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsUnderlineEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ITimedTextStyle2) Get_IsLineThroughEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsLineThroughEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextStyle2) Put_IsLineThroughEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsLineThroughEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ITimedTextStyle2) Get_IsOverlineEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsOverlineEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextStyle2) Put_IsOverlineEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsOverlineEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// F803F93B-3E99-595E-BBB7-78A2FA13C270
var IID_ITimedTextStyle3 = syscall.GUID{0xF803F93B, 0x3E99, 0x595E,
	[8]byte{0xBB, 0xB7, 0x78, 0xA2, 0xFA, 0x13, 0xC2, 0x70}}

type ITimedTextStyle3Interface interface {
	win32.IInspectableInterface
	Get_Ruby() *ITimedTextRuby
	Get_Bouten() *ITimedTextBouten
	Get_IsTextCombined() bool
	Put_IsTextCombined(value bool)
	Get_FontAngleInDegrees() float64
	Put_FontAngleInDegrees(value float64)
}

type ITimedTextStyle3Vtbl struct {
	win32.IInspectableVtbl
	Get_Ruby               uintptr
	Get_Bouten             uintptr
	Get_IsTextCombined     uintptr
	Put_IsTextCombined     uintptr
	Get_FontAngleInDegrees uintptr
	Put_FontAngleInDegrees uintptr
}

type ITimedTextStyle3 struct {
	win32.IInspectable
}

func (this *ITimedTextStyle3) Vtbl() *ITimedTextStyle3Vtbl {
	return (*ITimedTextStyle3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimedTextStyle3) Get_Ruby() *ITimedTextRuby {
	var _result *ITimedTextRuby
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ruby, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITimedTextStyle3) Get_Bouten() *ITimedTextBouten {
	var _result *ITimedTextBouten
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bouten, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITimedTextStyle3) Get_IsTextCombined() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsTextCombined, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextStyle3) Put_IsTextCombined(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsTextCombined, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ITimedTextStyle3) Get_FontAngleInDegrees() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FontAngleInDegrees, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextStyle3) Put_FontAngleInDegrees(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FontAngleInDegrees, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// D713502F-3261-4722-A0C2-B937B2390F14
var IID_ITimedTextSubformat = syscall.GUID{0xD713502F, 0x3261, 0x4722,
	[8]byte{0xA0, 0xC2, 0xB9, 0x37, 0xB2, 0x39, 0x0F, 0x14}}

type ITimedTextSubformatInterface interface {
	win32.IInspectableInterface
	Get_StartIndex() int32
	Put_StartIndex(value int32)
	Get_Length() int32
	Put_Length(value int32)
	Get_SubformatStyle() *ITimedTextStyle
	Put_SubformatStyle(value *ITimedTextStyle)
}

type ITimedTextSubformatVtbl struct {
	win32.IInspectableVtbl
	Get_StartIndex     uintptr
	Put_StartIndex     uintptr
	Get_Length         uintptr
	Put_Length         uintptr
	Get_SubformatStyle uintptr
	Put_SubformatStyle uintptr
}

type ITimedTextSubformat struct {
	win32.IInspectable
}

func (this *ITimedTextSubformat) Vtbl() *ITimedTextSubformatVtbl {
	return (*ITimedTextSubformatVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimedTextSubformat) Get_StartIndex() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StartIndex, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextSubformat) Put_StartIndex(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StartIndex, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ITimedTextSubformat) Get_Length() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Length, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedTextSubformat) Put_Length(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Length, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ITimedTextSubformat) Get_SubformatStyle() *ITimedTextStyle {
	var _result *ITimedTextStyle
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SubformatStyle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITimedTextSubformat) Put_SubformatStyle(value *ITimedTextStyle) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SubformatStyle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 0808A650-9698-4E57-877B-BD7CB2EE0F8A
var IID_IVideoStabilizationEffect = syscall.GUID{0x0808A650, 0x9698, 0x4E57,
	[8]byte{0x87, 0x7B, 0xBD, 0x7C, 0xB2, 0xEE, 0x0F, 0x8A}}

type IVideoStabilizationEffectInterface interface {
	win32.IInspectableInterface
	Put_Enabled(value bool)
	Get_Enabled() bool
	Add_EnabledChanged(handler TypedEventHandler[*IVideoStabilizationEffect, *IVideoStabilizationEffectEnabledChangedEventArgs]) EventRegistrationToken
	Remove_EnabledChanged(cookie EventRegistrationToken)
	GetRecommendedStreamConfiguration(controller *IVideoDeviceController, desiredProperties *IVideoEncodingProperties) *IVideoStreamConfiguration
}

type IVideoStabilizationEffectVtbl struct {
	win32.IInspectableVtbl
	Put_Enabled                       uintptr
	Get_Enabled                       uintptr
	Add_EnabledChanged                uintptr
	Remove_EnabledChanged             uintptr
	GetRecommendedStreamConfiguration uintptr
}

type IVideoStabilizationEffect struct {
	win32.IInspectable
}

func (this *IVideoStabilizationEffect) Vtbl() *IVideoStabilizationEffectVtbl {
	return (*IVideoStabilizationEffectVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoStabilizationEffect) Put_Enabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Enabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IVideoStabilizationEffect) Get_Enabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Enabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoStabilizationEffect) Add_EnabledChanged(handler TypedEventHandler[*IVideoStabilizationEffect, *IVideoStabilizationEffectEnabledChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_EnabledChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoStabilizationEffect) Remove_EnabledChanged(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_EnabledChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IVideoStabilizationEffect) GetRecommendedStreamConfiguration(controller *IVideoDeviceController, desiredProperties *IVideoEncodingProperties) *IVideoStreamConfiguration {
	var _result *IVideoStreamConfiguration
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetRecommendedStreamConfiguration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(controller)), uintptr(unsafe.Pointer(desiredProperties)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 187EFF28-67BB-4713-B900-4168DA164529
var IID_IVideoStabilizationEffectEnabledChangedEventArgs = syscall.GUID{0x187EFF28, 0x67BB, 0x4713,
	[8]byte{0xB9, 0x00, 0x41, 0x68, 0xDA, 0x16, 0x45, 0x29}}

type IVideoStabilizationEffectEnabledChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Reason() VideoStabilizationEffectEnabledChangedReason
}

type IVideoStabilizationEffectEnabledChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Reason uintptr
}

type IVideoStabilizationEffectEnabledChangedEventArgs struct {
	win32.IInspectable
}

func (this *IVideoStabilizationEffectEnabledChangedEventArgs) Vtbl() *IVideoStabilizationEffectEnabledChangedEventArgsVtbl {
	return (*IVideoStabilizationEffectEnabledChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoStabilizationEffectEnabledChangedEventArgs) Get_Reason() VideoStabilizationEffectEnabledChangedReason {
	var _result VideoStabilizationEffectEnabledChangedReason
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Reason, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 12EE0D55-9C2B-4440-8057-2C7A90F0CBEC
var IID_IVideoStreamDescriptor = syscall.GUID{0x12EE0D55, 0x9C2B, 0x4440,
	[8]byte{0x80, 0x57, 0x2C, 0x7A, 0x90, 0xF0, 0xCB, 0xEC}}

type IVideoStreamDescriptorInterface interface {
	win32.IInspectableInterface
	Get_EncodingProperties() *IVideoEncodingProperties
}

type IVideoStreamDescriptorVtbl struct {
	win32.IInspectableVtbl
	Get_EncodingProperties uintptr
}

type IVideoStreamDescriptor struct {
	win32.IInspectable
}

func (this *IVideoStreamDescriptor) Vtbl() *IVideoStreamDescriptorVtbl {
	return (*IVideoStreamDescriptorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoStreamDescriptor) Get_EncodingProperties() *IVideoEncodingProperties {
	var _result *IVideoEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EncodingProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8B306E10-453E-4088-832D-C36FA4F94AF3
var IID_IVideoStreamDescriptor2 = syscall.GUID{0x8B306E10, 0x453E, 0x4088,
	[8]byte{0x83, 0x2D, 0xC3, 0x6F, 0xA4, 0xF9, 0x4A, 0xF3}}

type IVideoStreamDescriptor2Interface interface {
	win32.IInspectableInterface
	Copy() *IVideoStreamDescriptor
}

type IVideoStreamDescriptor2Vtbl struct {
	win32.IInspectableVtbl
	Copy uintptr
}

type IVideoStreamDescriptor2 struct {
	win32.IInspectable
}

func (this *IVideoStreamDescriptor2) Vtbl() *IVideoStreamDescriptor2Vtbl {
	return (*IVideoStreamDescriptor2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoStreamDescriptor2) Copy() *IVideoStreamDescriptor {
	var _result *IVideoStreamDescriptor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Copy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 494EF6D1-BB75-43D2-9E5E-7B79A3AFCED4
var IID_IVideoStreamDescriptorFactory = syscall.GUID{0x494EF6D1, 0xBB75, 0x43D2,
	[8]byte{0x9E, 0x5E, 0x7B, 0x79, 0xA3, 0xAF, 0xCE, 0xD4}}

type IVideoStreamDescriptorFactoryInterface interface {
	win32.IInspectableInterface
	Create(encodingProperties *IVideoEncodingProperties) *IVideoStreamDescriptor
}

type IVideoStreamDescriptorFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IVideoStreamDescriptorFactory struct {
	win32.IInspectable
}

func (this *IVideoStreamDescriptorFactory) Vtbl() *IVideoStreamDescriptorFactoryVtbl {
	return (*IVideoStreamDescriptorFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoStreamDescriptorFactory) Create(encodingProperties *IVideoEncodingProperties) *IVideoStreamDescriptor {
	var _result *IVideoStreamDescriptor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(encodingProperties)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 99F3B7F3-E298-4396-BB6A-A51BE6A2A20A
var IID_IVideoTrack = syscall.GUID{0x99F3B7F3, 0xE298, 0x4396,
	[8]byte{0xBB, 0x6A, 0xA5, 0x1B, 0xE6, 0xA2, 0xA2, 0x0A}}

type IVideoTrackInterface interface {
	win32.IInspectableInterface
	Add_OpenFailed(handler TypedEventHandler[*IMediaTrack, *IVideoTrackOpenFailedEventArgs]) EventRegistrationToken
	Remove_OpenFailed(token EventRegistrationToken)
	GetEncodingProperties() *IVideoEncodingProperties
	Get_PlaybackItem() *IMediaPlaybackItem
	Get_Name() string
	Get_SupportInfo() *IVideoTrackSupportInfo
}

type IVideoTrackVtbl struct {
	win32.IInspectableVtbl
	Add_OpenFailed        uintptr
	Remove_OpenFailed     uintptr
	GetEncodingProperties uintptr
	Get_PlaybackItem      uintptr
	Get_Name              uintptr
	Get_SupportInfo       uintptr
}

type IVideoTrack struct {
	win32.IInspectable
}

func (this *IVideoTrack) Vtbl() *IVideoTrackVtbl {
	return (*IVideoTrackVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoTrack) Add_OpenFailed(handler TypedEventHandler[*IMediaTrack, *IVideoTrackOpenFailedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_OpenFailed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoTrack) Remove_OpenFailed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_OpenFailed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IVideoTrack) GetEncodingProperties() *IVideoEncodingProperties {
	var _result *IVideoEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetEncodingProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoTrack) Get_PlaybackItem() *IMediaPlaybackItem {
	var _result *IMediaPlaybackItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlaybackItem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoTrack) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IVideoTrack) Get_SupportInfo() *IVideoTrackSupportInfo {
	var _result *IVideoTrackSupportInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7679E231-04F9-4C82-A4EE-8602C8BB4754
var IID_IVideoTrackOpenFailedEventArgs = syscall.GUID{0x7679E231, 0x04F9, 0x4C82,
	[8]byte{0xA4, 0xEE, 0x86, 0x02, 0xC8, 0xBB, 0x47, 0x54}}

type IVideoTrackOpenFailedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ExtendedError() HResult
}

type IVideoTrackOpenFailedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ExtendedError uintptr
}

type IVideoTrackOpenFailedEventArgs struct {
	win32.IInspectable
}

func (this *IVideoTrackOpenFailedEventArgs) Vtbl() *IVideoTrackOpenFailedEventArgsVtbl {
	return (*IVideoTrackOpenFailedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoTrackOpenFailedEventArgs) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4BB534A0-FC5F-450D-8FF0-778D590486DE
var IID_IVideoTrackSupportInfo = syscall.GUID{0x4BB534A0, 0xFC5F, 0x450D,
	[8]byte{0x8F, 0xF0, 0x77, 0x8D, 0x59, 0x04, 0x86, 0xDE}}

type IVideoTrackSupportInfoInterface interface {
	win32.IInspectableInterface
	Get_DecoderStatus() MediaDecoderStatus
	Get_MediaSourceStatus() MediaSourceStatus
}

type IVideoTrackSupportInfoVtbl struct {
	win32.IInspectableVtbl
	Get_DecoderStatus     uintptr
	Get_MediaSourceStatus uintptr
}

type IVideoTrackSupportInfo struct {
	win32.IInspectable
}

func (this *IVideoTrackSupportInfo) Vtbl() *IVideoTrackSupportInfoVtbl {
	return (*IVideoTrackSupportInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoTrackSupportInfo) Get_DecoderStatus() MediaDecoderStatus {
	var _result MediaDecoderStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DecoderStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoTrackSupportInfo) Get_MediaSourceStatus() MediaSourceStatus {
	var _result MediaSourceStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaSourceStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type AudioStreamDescriptor struct {
	RtClass
	*IAudioStreamDescriptor
}

func NewAudioStreamDescriptor_Create(encodingProperties *IAudioEncodingProperties) *AudioStreamDescriptor {
	hs := NewHStr("Windows.Media.Core.AudioStreamDescriptor")
	var pFac *IAudioStreamDescriptorFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAudioStreamDescriptorFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IAudioStreamDescriptor
	p = pFac.Create(encodingProperties)
	result := &AudioStreamDescriptor{
		RtClass:                RtClass{PInspect: &p.IInspectable},
		IAudioStreamDescriptor: p,
	}
	com.AddToScope(result)
	return result
}

type AudioTrack struct {
	RtClass
	*IMediaTrack
}

type AudioTrackOpenFailedEventArgs struct {
	RtClass
	*IAudioTrackOpenFailedEventArgs
}

type AudioTrackSupportInfo struct {
	RtClass
	*IAudioTrackSupportInfo
}

type ChapterCue struct {
	RtClass
	*IChapterCue
}

func NewChapterCue() *ChapterCue {
	hs := NewHStr("Windows.Media.Core.ChapterCue")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &ChapterCue{
		RtClass:     RtClass{PInspect: p},
		IChapterCue: (*IChapterCue)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type CodecInfo struct {
	RtClass
	*ICodecInfo
}

type CodecQuery struct {
	RtClass
	*ICodecQuery
}

func NewCodecQuery() *CodecQuery {
	hs := NewHStr("Windows.Media.Core.CodecQuery")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &CodecQuery{
		RtClass:     RtClass{PInspect: p},
		ICodecQuery: (*ICodecQuery)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type DataCue struct {
	RtClass
	*IDataCue
}

func NewDataCue() *DataCue {
	hs := NewHStr("Windows.Media.Core.DataCue")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &DataCue{
		RtClass:  RtClass{PInspect: p},
		IDataCue: (*IDataCue)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type FaceDetectedEventArgs struct {
	RtClass
	*IFaceDetectedEventArgs
}

type FaceDetectionEffect struct {
	RtClass
	*IFaceDetectionEffect
}

type FaceDetectionEffectDefinition struct {
	RtClass
	*IVideoEffectDefinition
}

func NewFaceDetectionEffectDefinition() *FaceDetectionEffectDefinition {
	hs := NewHStr("Windows.Media.Core.FaceDetectionEffectDefinition")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &FaceDetectionEffectDefinition{
		RtClass:                RtClass{PInspect: p},
		IVideoEffectDefinition: (*IVideoEffectDefinition)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type FaceDetectionEffectFrame struct {
	RtClass
	*IFaceDetectionEffectFrame
}

type HighDynamicRangeControl struct {
	RtClass
	*IHighDynamicRangeControl
}

type HighDynamicRangeOutput struct {
	RtClass
	*IHighDynamicRangeOutput
}

type ImageCue struct {
	RtClass
	*IImageCue
}

func NewImageCue() *ImageCue {
	hs := NewHStr("Windows.Media.Core.ImageCue")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &ImageCue{
		RtClass:   RtClass{PInspect: p},
		IImageCue: (*IImageCue)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type LowLightFusion struct {
	RtClass
}

func NewILowLightFusionStatics() *ILowLightFusionStatics {
	var p *ILowLightFusionStatics
	hs := NewHStr("Windows.Media.Core.LowLightFusion")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ILowLightFusionStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type LowLightFusionResult struct {
	RtClass
	*ILowLightFusionResult
}

type MediaBinder struct {
	RtClass
	*IMediaBinder
}

func NewMediaBinder() *MediaBinder {
	hs := NewHStr("Windows.Media.Core.MediaBinder")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &MediaBinder{
		RtClass:      RtClass{PInspect: p},
		IMediaBinder: (*IMediaBinder)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type MediaBindingEventArgs struct {
	RtClass
	*IMediaBindingEventArgs
}

type MediaCueEventArgs struct {
	RtClass
	*IMediaCueEventArgs
}

type MediaSource struct {
	RtClass
	*IMediaSource2
}

func NewIMediaSourceStatics3() *IMediaSourceStatics3 {
	var p *IMediaSourceStatics3
	hs := NewHStr("Windows.Media.Core.MediaSource")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMediaSourceStatics3, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIMediaSourceStatics2() *IMediaSourceStatics2 {
	var p *IMediaSourceStatics2
	hs := NewHStr("Windows.Media.Core.MediaSource")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMediaSourceStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIMediaSourceStatics() *IMediaSourceStatics {
	var p *IMediaSourceStatics
	hs := NewHStr("Windows.Media.Core.MediaSource")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMediaSourceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIMediaSourceStatics4() *IMediaSourceStatics4 {
	var p *IMediaSourceStatics4
	hs := NewHStr("Windows.Media.Core.MediaSource")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMediaSourceStatics4, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type MediaSourceError struct {
	RtClass
	*IMediaSourceError
}

type MediaSourceOpenOperationCompletedEventArgs struct {
	RtClass
	*IMediaSourceOpenOperationCompletedEventArgs
}

type MediaSourceStateChangedEventArgs struct {
	RtClass
	*IMediaSourceStateChangedEventArgs
}

type MediaStreamSample struct {
	RtClass
	*IMediaStreamSample
}

func NewIMediaStreamSampleStatics2() *IMediaStreamSampleStatics2 {
	var p *IMediaStreamSampleStatics2
	hs := NewHStr("Windows.Media.Core.MediaStreamSample")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMediaStreamSampleStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIMediaStreamSampleStatics() *IMediaStreamSampleStatics {
	var p *IMediaStreamSampleStatics
	hs := NewHStr("Windows.Media.Core.MediaStreamSample")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMediaStreamSampleStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type MediaStreamSamplePropertySet struct {
	RtClass
	*IMap[syscall.GUID, interface{}]
}

type MediaStreamSampleProtectionProperties struct {
	RtClass
	*IMediaStreamSampleProtectionProperties
}

type MediaStreamSource struct {
	RtClass
	*IMediaStreamSource
}

func NewMediaStreamSource_CreateFromDescriptor(descriptor *IMediaStreamDescriptor) *MediaStreamSource {
	hs := NewHStr("Windows.Media.Core.MediaStreamSource")
	var pFac *IMediaStreamSourceFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMediaStreamSourceFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IMediaStreamSource
	p = pFac.CreateFromDescriptor(descriptor)
	result := &MediaStreamSource{
		RtClass:            RtClass{PInspect: &p.IInspectable},
		IMediaStreamSource: p,
	}
	com.AddToScope(result)
	return result
}

func NewMediaStreamSource_CreateFromDescriptors(descriptor *IMediaStreamDescriptor, descriptor2 *IMediaStreamDescriptor) *MediaStreamSource {
	hs := NewHStr("Windows.Media.Core.MediaStreamSource")
	var pFac *IMediaStreamSourceFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMediaStreamSourceFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IMediaStreamSource
	p = pFac.CreateFromDescriptors(descriptor, descriptor2)
	result := &MediaStreamSource{
		RtClass:            RtClass{PInspect: &p.IInspectable},
		IMediaStreamSource: p,
	}
	com.AddToScope(result)
	return result
}

type MediaStreamSourceClosedEventArgs struct {
	RtClass
	*IMediaStreamSourceClosedEventArgs
}

type MediaStreamSourceClosedRequest struct {
	RtClass
	*IMediaStreamSourceClosedRequest
}

type MediaStreamSourceSampleRenderedEventArgs struct {
	RtClass
	*IMediaStreamSourceSampleRenderedEventArgs
}

type MediaStreamSourceSampleRequest struct {
	RtClass
	*IMediaStreamSourceSampleRequest
}

type MediaStreamSourceSampleRequestDeferral struct {
	RtClass
	*IMediaStreamSourceSampleRequestDeferral
}

type MediaStreamSourceSampleRequestedEventArgs struct {
	RtClass
	*IMediaStreamSourceSampleRequestedEventArgs
}

type MediaStreamSourceStartingEventArgs struct {
	RtClass
	*IMediaStreamSourceStartingEventArgs
}

type MediaStreamSourceStartingRequest struct {
	RtClass
	*IMediaStreamSourceStartingRequest
}

type MediaStreamSourceStartingRequestDeferral struct {
	RtClass
	*IMediaStreamSourceStartingRequestDeferral
}

type MediaStreamSourceSwitchStreamsRequest struct {
	RtClass
	*IMediaStreamSourceSwitchStreamsRequest
}

type MediaStreamSourceSwitchStreamsRequestDeferral struct {
	RtClass
	*IMediaStreamSourceSwitchStreamsRequestDeferral
}

type MediaStreamSourceSwitchStreamsRequestedEventArgs struct {
	RtClass
	*IMediaStreamSourceSwitchStreamsRequestedEventArgs
}

type MseSourceBuffer struct {
	RtClass
	*IMseSourceBuffer
}

type MseSourceBufferList struct {
	RtClass
	*IMseSourceBufferList
}

type MseStreamSource struct {
	RtClass
	*IMseStreamSource
}

func NewMseStreamSource() *MseStreamSource {
	hs := NewHStr("Windows.Media.Core.MseStreamSource")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &MseStreamSource{
		RtClass:          RtClass{PInspect: p},
		IMseStreamSource: (*IMseStreamSource)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewIMseStreamSourceStatics() *IMseStreamSourceStatics {
	var p *IMseStreamSourceStatics
	hs := NewHStr("Windows.Media.Core.MseStreamSource")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMseStreamSourceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type SceneAnalysisEffect struct {
	RtClass
	*ISceneAnalysisEffect
}

type SceneAnalysisEffectDefinition struct {
	RtClass
	*IVideoEffectDefinition
}

func NewSceneAnalysisEffectDefinition() *SceneAnalysisEffectDefinition {
	hs := NewHStr("Windows.Media.Core.SceneAnalysisEffectDefinition")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &SceneAnalysisEffectDefinition{
		RtClass:                RtClass{PInspect: p},
		IVideoEffectDefinition: (*IVideoEffectDefinition)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type SceneAnalysisEffectFrame struct {
	RtClass
	*ISceneAnalysisEffectFrame
}

type SceneAnalyzedEventArgs struct {
	RtClass
	*ISceneAnalyzedEventArgs
}

type SpeechCue struct {
	RtClass
	*ISpeechCue
}

func NewSpeechCue() *SpeechCue {
	hs := NewHStr("Windows.Media.Core.SpeechCue")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &SpeechCue{
		RtClass:    RtClass{PInspect: p},
		ISpeechCue: (*ISpeechCue)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type TimedMetadataTrack struct {
	RtClass
	*ITimedMetadataTrack
}

func NewTimedMetadataTrack_Create(id string, language string, kind TimedMetadataKind) *TimedMetadataTrack {
	hs := NewHStr("Windows.Media.Core.TimedMetadataTrack")
	var pFac *ITimedMetadataTrackFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITimedMetadataTrackFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ITimedMetadataTrack
	p = pFac.Create(id, language, kind)
	result := &TimedMetadataTrack{
		RtClass:             RtClass{PInspect: &p.IInspectable},
		ITimedMetadataTrack: p,
	}
	com.AddToScope(result)
	return result
}

type TimedMetadataTrackError struct {
	RtClass
	*ITimedMetadataTrackError
}

type TimedMetadataTrackFailedEventArgs struct {
	RtClass
	*ITimedMetadataTrackFailedEventArgs
}

type TimedTextBouten struct {
	RtClass
	*ITimedTextBouten
}

type TimedTextCue struct {
	RtClass
	*ITimedTextCue
}

func NewTimedTextCue() *TimedTextCue {
	hs := NewHStr("Windows.Media.Core.TimedTextCue")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &TimedTextCue{
		RtClass:       RtClass{PInspect: p},
		ITimedTextCue: (*ITimedTextCue)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type TimedTextLine struct {
	RtClass
	*ITimedTextLine
}

func NewTimedTextLine() *TimedTextLine {
	hs := NewHStr("Windows.Media.Core.TimedTextLine")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &TimedTextLine{
		RtClass:        RtClass{PInspect: p},
		ITimedTextLine: (*ITimedTextLine)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type TimedTextRegion struct {
	RtClass
	*ITimedTextRegion
}

func NewTimedTextRegion() *TimedTextRegion {
	hs := NewHStr("Windows.Media.Core.TimedTextRegion")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &TimedTextRegion{
		RtClass:          RtClass{PInspect: p},
		ITimedTextRegion: (*ITimedTextRegion)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type TimedTextRuby struct {
	RtClass
	*ITimedTextRuby
}

type TimedTextSource struct {
	RtClass
	*ITimedTextSource
}

func NewITimedTextSourceStatics2() *ITimedTextSourceStatics2 {
	var p *ITimedTextSourceStatics2
	hs := NewHStr("Windows.Media.Core.TimedTextSource")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITimedTextSourceStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewITimedTextSourceStatics() *ITimedTextSourceStatics {
	var p *ITimedTextSourceStatics
	hs := NewHStr("Windows.Media.Core.TimedTextSource")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITimedTextSourceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type TimedTextSourceResolveResultEventArgs struct {
	RtClass
	*ITimedTextSourceResolveResultEventArgs
}

type TimedTextStyle struct {
	RtClass
	*ITimedTextStyle
}

func NewTimedTextStyle() *TimedTextStyle {
	hs := NewHStr("Windows.Media.Core.TimedTextStyle")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &TimedTextStyle{
		RtClass:         RtClass{PInspect: p},
		ITimedTextStyle: (*ITimedTextStyle)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type TimedTextSubformat struct {
	RtClass
	*ITimedTextSubformat
}

func NewTimedTextSubformat() *TimedTextSubformat {
	hs := NewHStr("Windows.Media.Core.TimedTextSubformat")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &TimedTextSubformat{
		RtClass:             RtClass{PInspect: p},
		ITimedTextSubformat: (*ITimedTextSubformat)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type VideoStabilizationEffect struct {
	RtClass
	*IVideoStabilizationEffect
}

type VideoStabilizationEffectDefinition struct {
	RtClass
	*IVideoEffectDefinition
}

func NewVideoStabilizationEffectDefinition() *VideoStabilizationEffectDefinition {
	hs := NewHStr("Windows.Media.Core.VideoStabilizationEffectDefinition")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &VideoStabilizationEffectDefinition{
		RtClass:                RtClass{PInspect: p},
		IVideoEffectDefinition: (*IVideoEffectDefinition)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type VideoStabilizationEffectEnabledChangedEventArgs struct {
	RtClass
	*IVideoStabilizationEffectEnabledChangedEventArgs
}

type VideoStreamDescriptor struct {
	RtClass
	*IVideoStreamDescriptor
}

func NewVideoStreamDescriptor_Create(encodingProperties *IVideoEncodingProperties) *VideoStreamDescriptor {
	hs := NewHStr("Windows.Media.Core.VideoStreamDescriptor")
	var pFac *IVideoStreamDescriptorFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IVideoStreamDescriptorFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IVideoStreamDescriptor
	p = pFac.Create(encodingProperties)
	result := &VideoStreamDescriptor{
		RtClass:                RtClass{PInspect: &p.IInspectable},
		IVideoStreamDescriptor: p,
	}
	com.AddToScope(result)
	return result
}

type VideoTrack struct {
	RtClass
	*IMediaTrack
}

type VideoTrackOpenFailedEventArgs struct {
	RtClass
	*IVideoTrackOpenFailedEventArgs
}

type VideoTrackSupportInfo struct {
	RtClass
	*IVideoTrackSupportInfo
}
