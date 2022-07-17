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
type AudioDeviceNodeCreationStatus int32

const (
	AudioDeviceNodeCreationStatus_Success            AudioDeviceNodeCreationStatus = 0
	AudioDeviceNodeCreationStatus_DeviceNotAvailable AudioDeviceNodeCreationStatus = 1
	AudioDeviceNodeCreationStatus_FormatNotSupported AudioDeviceNodeCreationStatus = 2
	AudioDeviceNodeCreationStatus_UnknownFailure     AudioDeviceNodeCreationStatus = 3
	AudioDeviceNodeCreationStatus_AccessDenied       AudioDeviceNodeCreationStatus = 4
)

// enum
type AudioFileNodeCreationStatus int32

const (
	AudioFileNodeCreationStatus_Success            AudioFileNodeCreationStatus = 0
	AudioFileNodeCreationStatus_FileNotFound       AudioFileNodeCreationStatus = 1
	AudioFileNodeCreationStatus_InvalidFileType    AudioFileNodeCreationStatus = 2
	AudioFileNodeCreationStatus_FormatNotSupported AudioFileNodeCreationStatus = 3
	AudioFileNodeCreationStatus_UnknownFailure     AudioFileNodeCreationStatus = 4
)

// enum
type AudioGraphCreationStatus int32

const (
	AudioGraphCreationStatus_Success            AudioGraphCreationStatus = 0
	AudioGraphCreationStatus_DeviceNotAvailable AudioGraphCreationStatus = 1
	AudioGraphCreationStatus_FormatNotSupported AudioGraphCreationStatus = 2
	AudioGraphCreationStatus_UnknownFailure     AudioGraphCreationStatus = 3
)

// enum
type AudioGraphUnrecoverableError int32

const (
	AudioGraphUnrecoverableError_None                     AudioGraphUnrecoverableError = 0
	AudioGraphUnrecoverableError_AudioDeviceLost          AudioGraphUnrecoverableError = 1
	AudioGraphUnrecoverableError_AudioSessionDisconnected AudioGraphUnrecoverableError = 2
	AudioGraphUnrecoverableError_UnknownFailure           AudioGraphUnrecoverableError = 3
)

// enum
type AudioNodeEmitterDecayKind int32

const (
	AudioNodeEmitterDecayKind_Natural AudioNodeEmitterDecayKind = 0
	AudioNodeEmitterDecayKind_Custom  AudioNodeEmitterDecayKind = 1
)

// enum
// flags
type AudioNodeEmitterSettings uint32

const (
	AudioNodeEmitterSettings_None           AudioNodeEmitterSettings = 0
	AudioNodeEmitterSettings_DisableDoppler AudioNodeEmitterSettings = 1
)

// enum
type AudioNodeEmitterShapeKind int32

const (
	AudioNodeEmitterShapeKind_Omnidirectional AudioNodeEmitterShapeKind = 0
	AudioNodeEmitterShapeKind_Cone            AudioNodeEmitterShapeKind = 1
)

// enum
type AudioPlaybackConnectionOpenResultStatus int32

const (
	AudioPlaybackConnectionOpenResultStatus_Success         AudioPlaybackConnectionOpenResultStatus = 0
	AudioPlaybackConnectionOpenResultStatus_RequestTimedOut AudioPlaybackConnectionOpenResultStatus = 1
	AudioPlaybackConnectionOpenResultStatus_DeniedBySystem  AudioPlaybackConnectionOpenResultStatus = 2
	AudioPlaybackConnectionOpenResultStatus_UnknownFailure  AudioPlaybackConnectionOpenResultStatus = 3
)

// enum
type AudioPlaybackConnectionState int32

const (
	AudioPlaybackConnectionState_Closed AudioPlaybackConnectionState = 0
	AudioPlaybackConnectionState_Opened AudioPlaybackConnectionState = 1
)

// enum
type MediaSourceAudioInputNodeCreationStatus int32

const (
	MediaSourceAudioInputNodeCreationStatus_Success            MediaSourceAudioInputNodeCreationStatus = 0
	MediaSourceAudioInputNodeCreationStatus_FormatNotSupported MediaSourceAudioInputNodeCreationStatus = 1
	MediaSourceAudioInputNodeCreationStatus_NetworkError       MediaSourceAudioInputNodeCreationStatus = 2
	MediaSourceAudioInputNodeCreationStatus_UnknownFailure     MediaSourceAudioInputNodeCreationStatus = 3
)

// enum
type MixedRealitySpatialAudioFormatPolicy int32

const (
	MixedRealitySpatialAudioFormatPolicy_UseMixedRealityDefaultSpatialAudioFormat        MixedRealitySpatialAudioFormatPolicy = 0
	MixedRealitySpatialAudioFormatPolicy_UseDeviceConfigurationDefaultSpatialAudioFormat MixedRealitySpatialAudioFormatPolicy = 1
)

// enum
type QuantumSizeSelectionMode int32

const (
	QuantumSizeSelectionMode_SystemDefault    QuantumSizeSelectionMode = 0
	QuantumSizeSelectionMode_LowestLatency    QuantumSizeSelectionMode = 1
	QuantumSizeSelectionMode_ClosestToDesired QuantumSizeSelectionMode = 2
)

// enum
type SetDefaultSpatialAudioFormatStatus int32

const (
	SetDefaultSpatialAudioFormatStatus_Succeeded                       SetDefaultSpatialAudioFormatStatus = 0
	SetDefaultSpatialAudioFormatStatus_AccessDenied                    SetDefaultSpatialAudioFormatStatus = 1
	SetDefaultSpatialAudioFormatStatus_LicenseExpired                  SetDefaultSpatialAudioFormatStatus = 2
	SetDefaultSpatialAudioFormatStatus_LicenseNotValidForAudioEndpoint SetDefaultSpatialAudioFormatStatus = 3
	SetDefaultSpatialAudioFormatStatus_NotSupportedOnAudioEndpoint     SetDefaultSpatialAudioFormatStatus = 4
	SetDefaultSpatialAudioFormatStatus_UnknownError                    SetDefaultSpatialAudioFormatStatus = 5
)

// enum
type SpatialAudioModel int32

const (
	SpatialAudioModel_ObjectBased SpatialAudioModel = 0
	SpatialAudioModel_FoldDown    SpatialAudioModel = 1
)

// interfaces

// B01B6BE1-6F4E-49E2-AC01-559D62BEB3A9
var IID_IAudioDeviceInputNode = syscall.GUID{0xB01B6BE1, 0x6F4E, 0x49E2,
	[8]byte{0xAC, 0x01, 0x55, 0x9D, 0x62, 0xBE, 0xB3, 0xA9}}

type IAudioDeviceInputNodeInterface interface {
	win32.IInspectableInterface
	Get_Device() *IDeviceInformation
}

type IAudioDeviceInputNodeVtbl struct {
	win32.IInspectableVtbl
	Get_Device uintptr
}

type IAudioDeviceInputNode struct {
	win32.IInspectable
}

func (this *IAudioDeviceInputNode) Vtbl() *IAudioDeviceInputNodeVtbl {
	return (*IAudioDeviceInputNodeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioDeviceInputNode) Get_Device() *IDeviceInformation {
	var _result *IDeviceInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Device, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 362EDBFF-FF1C-4434-9E0F-BD2EF522AC82
var IID_IAudioDeviceOutputNode = syscall.GUID{0x362EDBFF, 0xFF1C, 0x4434,
	[8]byte{0x9E, 0x0F, 0xBD, 0x2E, 0xF5, 0x22, 0xAC, 0x82}}

type IAudioDeviceOutputNodeInterface interface {
	win32.IInspectableInterface
	Get_Device() *IDeviceInformation
}

type IAudioDeviceOutputNodeVtbl struct {
	win32.IInspectableVtbl
	Get_Device uintptr
}

type IAudioDeviceOutputNode struct {
	win32.IInspectable
}

func (this *IAudioDeviceOutputNode) Vtbl() *IAudioDeviceOutputNodeVtbl {
	return (*IAudioDeviceOutputNodeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioDeviceOutputNode) Get_Device() *IDeviceInformation {
	var _result *IDeviceInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Device, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 905B67C8-6F65-4CD4-8890-4694843C276D
var IID_IAudioFileInputNode = syscall.GUID{0x905B67C8, 0x6F65, 0x4CD4,
	[8]byte{0x88, 0x90, 0x46, 0x94, 0x84, 0x3C, 0x27, 0x6D}}

type IAudioFileInputNodeInterface interface {
	win32.IInspectableInterface
	Put_PlaybackSpeedFactor(value float64)
	Get_PlaybackSpeedFactor() float64
	Get_Position() TimeSpan
	Seek(position TimeSpan)
	Get_StartTime() *IReference[TimeSpan]
	Put_StartTime(value *IReference[TimeSpan])
	Get_EndTime() *IReference[TimeSpan]
	Put_EndTime(value *IReference[TimeSpan])
	Get_LoopCount() *IReference[int32]
	Put_LoopCount(value *IReference[int32])
	Get_Duration() TimeSpan
	Get_SourceFile() *IStorageFile
	Add_FileCompleted(handler TypedEventHandler[*IAudioFileInputNode, interface{}]) EventRegistrationToken
	Remove_FileCompleted(token EventRegistrationToken)
}

type IAudioFileInputNodeVtbl struct {
	win32.IInspectableVtbl
	Put_PlaybackSpeedFactor uintptr
	Get_PlaybackSpeedFactor uintptr
	Get_Position            uintptr
	Seek                    uintptr
	Get_StartTime           uintptr
	Put_StartTime           uintptr
	Get_EndTime             uintptr
	Put_EndTime             uintptr
	Get_LoopCount           uintptr
	Put_LoopCount           uintptr
	Get_Duration            uintptr
	Get_SourceFile          uintptr
	Add_FileCompleted       uintptr
	Remove_FileCompleted    uintptr
}

type IAudioFileInputNode struct {
	win32.IInspectable
}

func (this *IAudioFileInputNode) Vtbl() *IAudioFileInputNodeVtbl {
	return (*IAudioFileInputNodeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioFileInputNode) Put_PlaybackSpeedFactor(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PlaybackSpeedFactor, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAudioFileInputNode) Get_PlaybackSpeedFactor() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlaybackSpeedFactor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioFileInputNode) Get_Position() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioFileInputNode) Seek(position TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Seek, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&position)))
	_ = _hr
}

func (this *IAudioFileInputNode) Get_StartTime() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StartTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioFileInputNode) Put_StartTime(value *IReference[TimeSpan]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StartTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAudioFileInputNode) Get_EndTime() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EndTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioFileInputNode) Put_EndTime(value *IReference[TimeSpan]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_EndTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAudioFileInputNode) Get_LoopCount() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LoopCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioFileInputNode) Put_LoopCount(value *IReference[int32]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LoopCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAudioFileInputNode) Get_Duration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Duration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioFileInputNode) Get_SourceFile() *IStorageFile {
	var _result *IStorageFile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SourceFile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioFileInputNode) Add_FileCompleted(handler TypedEventHandler[*IAudioFileInputNode, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_FileCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioFileInputNode) Remove_FileCompleted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_FileCompleted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 50E01980-5166-4093-80F8-ADA00089E9CF
var IID_IAudioFileOutputNode = syscall.GUID{0x50E01980, 0x5166, 0x4093,
	[8]byte{0x80, 0xF8, 0xAD, 0xA0, 0x00, 0x89, 0xE9, 0xCF}}

type IAudioFileOutputNodeInterface interface {
	win32.IInspectableInterface
	Get_File() *IStorageFile
	Get_FileEncodingProfile() *IMediaEncodingProfile
	FinalizeAsync() *IAsyncOperation[TranscodeFailureReason]
}

type IAudioFileOutputNodeVtbl struct {
	win32.IInspectableVtbl
	Get_File                uintptr
	Get_FileEncodingProfile uintptr
	FinalizeAsync           uintptr
}

type IAudioFileOutputNode struct {
	win32.IInspectable
}

func (this *IAudioFileOutputNode) Vtbl() *IAudioFileOutputNodeVtbl {
	return (*IAudioFileOutputNodeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioFileOutputNode) Get_File() *IStorageFile {
	var _result *IStorageFile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_File, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioFileOutputNode) Get_FileEncodingProfile() *IMediaEncodingProfile {
	var _result *IMediaEncodingProfile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FileEncodingProfile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioFileOutputNode) FinalizeAsync() *IAsyncOperation[TranscodeFailureReason] {
	var _result *IAsyncOperation[TranscodeFailureReason]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FinalizeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DC7C829E-0208-4504-A5A8-F0F268920A65
var IID_IAudioFrameCompletedEventArgs = syscall.GUID{0xDC7C829E, 0x0208, 0x4504,
	[8]byte{0xA5, 0xA8, 0xF0, 0xF2, 0x68, 0x92, 0x0A, 0x65}}

type IAudioFrameCompletedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Frame() *IAudioFrame
}

type IAudioFrameCompletedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Frame uintptr
}

type IAudioFrameCompletedEventArgs struct {
	win32.IInspectable
}

func (this *IAudioFrameCompletedEventArgs) Vtbl() *IAudioFrameCompletedEventArgsVtbl {
	return (*IAudioFrameCompletedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioFrameCompletedEventArgs) Get_Frame() *IAudioFrame {
	var _result *IAudioFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Frame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 01B266C7-FD96-4FF5-A3C5-D27A9BF44237
var IID_IAudioFrameInputNode = syscall.GUID{0x01B266C7, 0xFD96, 0x4FF5,
	[8]byte{0xA3, 0xC5, 0xD2, 0x7A, 0x9B, 0xF4, 0x42, 0x37}}

type IAudioFrameInputNodeInterface interface {
	win32.IInspectableInterface
	Put_PlaybackSpeedFactor(value float64)
	Get_PlaybackSpeedFactor() float64
	AddFrame(frame *IAudioFrame)
	DiscardQueuedFrames()
	Get_QueuedSampleCount() uint64
	Add_AudioFrameCompleted(handler TypedEventHandler[*IAudioFrameInputNode, *IAudioFrameCompletedEventArgs]) EventRegistrationToken
	Remove_AudioFrameCompleted(token EventRegistrationToken)
	Add_QuantumStarted(handler TypedEventHandler[*IAudioFrameInputNode, *IFrameInputNodeQuantumStartedEventArgs]) EventRegistrationToken
	Remove_QuantumStarted(token EventRegistrationToken)
}

type IAudioFrameInputNodeVtbl struct {
	win32.IInspectableVtbl
	Put_PlaybackSpeedFactor    uintptr
	Get_PlaybackSpeedFactor    uintptr
	AddFrame                   uintptr
	DiscardQueuedFrames        uintptr
	Get_QueuedSampleCount      uintptr
	Add_AudioFrameCompleted    uintptr
	Remove_AudioFrameCompleted uintptr
	Add_QuantumStarted         uintptr
	Remove_QuantumStarted      uintptr
}

type IAudioFrameInputNode struct {
	win32.IInspectable
}

func (this *IAudioFrameInputNode) Vtbl() *IAudioFrameInputNodeVtbl {
	return (*IAudioFrameInputNodeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioFrameInputNode) Put_PlaybackSpeedFactor(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PlaybackSpeedFactor, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAudioFrameInputNode) Get_PlaybackSpeedFactor() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlaybackSpeedFactor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioFrameInputNode) AddFrame(frame *IAudioFrame) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(frame)))
	_ = _hr
}

func (this *IAudioFrameInputNode) DiscardQueuedFrames() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DiscardQueuedFrames, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IAudioFrameInputNode) Get_QueuedSampleCount() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_QueuedSampleCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioFrameInputNode) Add_AudioFrameCompleted(handler TypedEventHandler[*IAudioFrameInputNode, *IAudioFrameCompletedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AudioFrameCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioFrameInputNode) Remove_AudioFrameCompleted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AudioFrameCompleted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAudioFrameInputNode) Add_QuantumStarted(handler TypedEventHandler[*IAudioFrameInputNode, *IFrameInputNodeQuantumStartedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_QuantumStarted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioFrameInputNode) Remove_QuantumStarted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_QuantumStarted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// B847371B-3299-45F5-88B3-C9D12A3F1CC8
var IID_IAudioFrameOutputNode = syscall.GUID{0xB847371B, 0x3299, 0x45F5,
	[8]byte{0x88, 0xB3, 0xC9, 0xD1, 0x2A, 0x3F, 0x1C, 0xC8}}

type IAudioFrameOutputNodeInterface interface {
	win32.IInspectableInterface
	GetFrame() *IAudioFrame
}

type IAudioFrameOutputNodeVtbl struct {
	win32.IInspectableVtbl
	GetFrame uintptr
}

type IAudioFrameOutputNode struct {
	win32.IInspectable
}

func (this *IAudioFrameOutputNode) Vtbl() *IAudioFrameOutputNodeVtbl {
	return (*IAudioFrameOutputNodeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioFrameOutputNode) GetFrame() *IAudioFrame {
	var _result *IAudioFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1AD46EED-E48C-4E14-9660-2C4F83E9CDD8
var IID_IAudioGraph = syscall.GUID{0x1AD46EED, 0xE48C, 0x4E14,
	[8]byte{0x96, 0x60, 0x2C, 0x4F, 0x83, 0xE9, 0xCD, 0xD8}}

type IAudioGraphInterface interface {
	win32.IInspectableInterface
	CreateFrameInputNode() *IAudioFrameInputNode
	CreateFrameInputNodeWithFormat(encodingProperties *IAudioEncodingProperties) *IAudioFrameInputNode
	CreateDeviceInputNodeAsync(category MediaCategory) *IAsyncOperation[*ICreateAudioDeviceInputNodeResult]
	CreateDeviceInputNodeWithFormatAsync(category MediaCategory, encodingProperties *IAudioEncodingProperties) *IAsyncOperation[*ICreateAudioDeviceInputNodeResult]
	CreateDeviceInputNodeWithFormatOnDeviceAsync(category MediaCategory, encodingProperties *IAudioEncodingProperties, device *IDeviceInformation) *IAsyncOperation[*ICreateAudioDeviceInputNodeResult]
	CreateFrameOutputNode() *IAudioFrameOutputNode
	CreateFrameOutputNodeWithFormat(encodingProperties *IAudioEncodingProperties) *IAudioFrameOutputNode
	CreateDeviceOutputNodeAsync() *IAsyncOperation[*ICreateAudioDeviceOutputNodeResult]
	CreateFileInputNodeAsync(file *IStorageFile) *IAsyncOperation[*ICreateAudioFileInputNodeResult]
	CreateFileOutputNodeAsync(file *IStorageFile) *IAsyncOperation[*ICreateAudioFileOutputNodeResult]
	CreateFileOutputNodeWithFileProfileAsync(file *IStorageFile, fileEncodingProfile *IMediaEncodingProfile) *IAsyncOperation[*ICreateAudioFileOutputNodeResult]
	CreateSubmixNode() *IAudioInputNode
	CreateSubmixNodeWithFormat(encodingProperties *IAudioEncodingProperties) *IAudioInputNode
	Start()
	Stop()
	ResetAllNodes()
	Add_QuantumStarted(handler TypedEventHandler[*IAudioGraph, interface{}]) EventRegistrationToken
	Remove_QuantumStarted(token EventRegistrationToken)
	Add_QuantumProcessed(handler TypedEventHandler[*IAudioGraph, interface{}]) EventRegistrationToken
	Remove_QuantumProcessed(token EventRegistrationToken)
	Add_UnrecoverableErrorOccurred(handler TypedEventHandler[*IAudioGraph, *IAudioGraphUnrecoverableErrorOccurredEventArgs]) EventRegistrationToken
	Remove_UnrecoverableErrorOccurred(token EventRegistrationToken)
	Get_CompletedQuantumCount() uint64
	Get_EncodingProperties() *IAudioEncodingProperties
	Get_LatencyInSamples() int32
	Get_PrimaryRenderDevice() *IDeviceInformation
	Get_RenderDeviceAudioProcessing() AudioProcessing
	Get_SamplesPerQuantum() int32
}

type IAudioGraphVtbl struct {
	win32.IInspectableVtbl
	CreateFrameInputNode                         uintptr
	CreateFrameInputNodeWithFormat               uintptr
	CreateDeviceInputNodeAsync                   uintptr
	CreateDeviceInputNodeWithFormatAsync         uintptr
	CreateDeviceInputNodeWithFormatOnDeviceAsync uintptr
	CreateFrameOutputNode                        uintptr
	CreateFrameOutputNodeWithFormat              uintptr
	CreateDeviceOutputNodeAsync                  uintptr
	CreateFileInputNodeAsync                     uintptr
	CreateFileOutputNodeAsync                    uintptr
	CreateFileOutputNodeWithFileProfileAsync     uintptr
	CreateSubmixNode                             uintptr
	CreateSubmixNodeWithFormat                   uintptr
	Start                                        uintptr
	Stop                                         uintptr
	ResetAllNodes                                uintptr
	Add_QuantumStarted                           uintptr
	Remove_QuantumStarted                        uintptr
	Add_QuantumProcessed                         uintptr
	Remove_QuantumProcessed                      uintptr
	Add_UnrecoverableErrorOccurred               uintptr
	Remove_UnrecoverableErrorOccurred            uintptr
	Get_CompletedQuantumCount                    uintptr
	Get_EncodingProperties                       uintptr
	Get_LatencyInSamples                         uintptr
	Get_PrimaryRenderDevice                      uintptr
	Get_RenderDeviceAudioProcessing              uintptr
	Get_SamplesPerQuantum                        uintptr
}

type IAudioGraph struct {
	win32.IInspectable
}

func (this *IAudioGraph) Vtbl() *IAudioGraphVtbl {
	return (*IAudioGraphVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioGraph) CreateFrameInputNode() *IAudioFrameInputNode {
	var _result *IAudioFrameInputNode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFrameInputNode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioGraph) CreateFrameInputNodeWithFormat(encodingProperties *IAudioEncodingProperties) *IAudioFrameInputNode {
	var _result *IAudioFrameInputNode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFrameInputNodeWithFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(encodingProperties)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioGraph) CreateDeviceInputNodeAsync(category MediaCategory) *IAsyncOperation[*ICreateAudioDeviceInputNodeResult] {
	var _result *IAsyncOperation[*ICreateAudioDeviceInputNodeResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDeviceInputNodeAsync, uintptr(unsafe.Pointer(this)), uintptr(category), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioGraph) CreateDeviceInputNodeWithFormatAsync(category MediaCategory, encodingProperties *IAudioEncodingProperties) *IAsyncOperation[*ICreateAudioDeviceInputNodeResult] {
	var _result *IAsyncOperation[*ICreateAudioDeviceInputNodeResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDeviceInputNodeWithFormatAsync, uintptr(unsafe.Pointer(this)), uintptr(category), uintptr(unsafe.Pointer(encodingProperties)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioGraph) CreateDeviceInputNodeWithFormatOnDeviceAsync(category MediaCategory, encodingProperties *IAudioEncodingProperties, device *IDeviceInformation) *IAsyncOperation[*ICreateAudioDeviceInputNodeResult] {
	var _result *IAsyncOperation[*ICreateAudioDeviceInputNodeResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDeviceInputNodeWithFormatOnDeviceAsync, uintptr(unsafe.Pointer(this)), uintptr(category), uintptr(unsafe.Pointer(encodingProperties)), uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioGraph) CreateFrameOutputNode() *IAudioFrameOutputNode {
	var _result *IAudioFrameOutputNode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFrameOutputNode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioGraph) CreateFrameOutputNodeWithFormat(encodingProperties *IAudioEncodingProperties) *IAudioFrameOutputNode {
	var _result *IAudioFrameOutputNode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFrameOutputNodeWithFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(encodingProperties)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioGraph) CreateDeviceOutputNodeAsync() *IAsyncOperation[*ICreateAudioDeviceOutputNodeResult] {
	var _result *IAsyncOperation[*ICreateAudioDeviceOutputNodeResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDeviceOutputNodeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioGraph) CreateFileInputNodeAsync(file *IStorageFile) *IAsyncOperation[*ICreateAudioFileInputNodeResult] {
	var _result *IAsyncOperation[*ICreateAudioFileInputNodeResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFileInputNodeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioGraph) CreateFileOutputNodeAsync(file *IStorageFile) *IAsyncOperation[*ICreateAudioFileOutputNodeResult] {
	var _result *IAsyncOperation[*ICreateAudioFileOutputNodeResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFileOutputNodeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioGraph) CreateFileOutputNodeWithFileProfileAsync(file *IStorageFile, fileEncodingProfile *IMediaEncodingProfile) *IAsyncOperation[*ICreateAudioFileOutputNodeResult] {
	var _result *IAsyncOperation[*ICreateAudioFileOutputNodeResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFileOutputNodeWithFileProfileAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(fileEncodingProfile)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioGraph) CreateSubmixNode() *IAudioInputNode {
	var _result *IAudioInputNode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSubmixNode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioGraph) CreateSubmixNodeWithFormat(encodingProperties *IAudioEncodingProperties) *IAudioInputNode {
	var _result *IAudioInputNode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSubmixNodeWithFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(encodingProperties)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioGraph) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IAudioGraph) Stop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IAudioGraph) ResetAllNodes() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ResetAllNodes, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IAudioGraph) Add_QuantumStarted(handler TypedEventHandler[*IAudioGraph, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_QuantumStarted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioGraph) Remove_QuantumStarted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_QuantumStarted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAudioGraph) Add_QuantumProcessed(handler TypedEventHandler[*IAudioGraph, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_QuantumProcessed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioGraph) Remove_QuantumProcessed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_QuantumProcessed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAudioGraph) Add_UnrecoverableErrorOccurred(handler TypedEventHandler[*IAudioGraph, *IAudioGraphUnrecoverableErrorOccurredEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_UnrecoverableErrorOccurred, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioGraph) Remove_UnrecoverableErrorOccurred(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_UnrecoverableErrorOccurred, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAudioGraph) Get_CompletedQuantumCount() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CompletedQuantumCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioGraph) Get_EncodingProperties() *IAudioEncodingProperties {
	var _result *IAudioEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EncodingProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioGraph) Get_LatencyInSamples() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LatencyInSamples, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioGraph) Get_PrimaryRenderDevice() *IDeviceInformation {
	var _result *IDeviceInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PrimaryRenderDevice, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioGraph) Get_RenderDeviceAudioProcessing() AudioProcessing {
	var _result AudioProcessing
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RenderDeviceAudioProcessing, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioGraph) Get_SamplesPerQuantum() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SamplesPerQuantum, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4E4C3BD5-4FC1-45F6-A947-3CD38F4FD839
var IID_IAudioGraph2 = syscall.GUID{0x4E4C3BD5, 0x4FC1, 0x45F6,
	[8]byte{0xA9, 0x47, 0x3C, 0xD3, 0x8F, 0x4F, 0xD8, 0x39}}

type IAudioGraph2Interface interface {
	win32.IInspectableInterface
	CreateFrameInputNodeWithFormatAndEmitter(encodingProperties *IAudioEncodingProperties, emitter *IAudioNodeEmitter) *IAudioFrameInputNode
	CreateDeviceInputNodeWithFormatAndEmitterOnDeviceAsync(category MediaCategory, encodingProperties *IAudioEncodingProperties, device *IDeviceInformation, emitter *IAudioNodeEmitter) *IAsyncOperation[*ICreateAudioDeviceInputNodeResult]
	CreateFileInputNodeWithEmitterAsync(file *IStorageFile, emitter *IAudioNodeEmitter) *IAsyncOperation[*ICreateAudioFileInputNodeResult]
	CreateSubmixNodeWithFormatAndEmitter(encodingProperties *IAudioEncodingProperties, emitter *IAudioNodeEmitter) *IAudioInputNode
	CreateBatchUpdater() *IClosable
}

type IAudioGraph2Vtbl struct {
	win32.IInspectableVtbl
	CreateFrameInputNodeWithFormatAndEmitter               uintptr
	CreateDeviceInputNodeWithFormatAndEmitterOnDeviceAsync uintptr
	CreateFileInputNodeWithEmitterAsync                    uintptr
	CreateSubmixNodeWithFormatAndEmitter                   uintptr
	CreateBatchUpdater                                     uintptr
}

type IAudioGraph2 struct {
	win32.IInspectable
}

func (this *IAudioGraph2) Vtbl() *IAudioGraph2Vtbl {
	return (*IAudioGraph2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioGraph2) CreateFrameInputNodeWithFormatAndEmitter(encodingProperties *IAudioEncodingProperties, emitter *IAudioNodeEmitter) *IAudioFrameInputNode {
	var _result *IAudioFrameInputNode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFrameInputNodeWithFormatAndEmitter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(encodingProperties)), uintptr(unsafe.Pointer(emitter)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioGraph2) CreateDeviceInputNodeWithFormatAndEmitterOnDeviceAsync(category MediaCategory, encodingProperties *IAudioEncodingProperties, device *IDeviceInformation, emitter *IAudioNodeEmitter) *IAsyncOperation[*ICreateAudioDeviceInputNodeResult] {
	var _result *IAsyncOperation[*ICreateAudioDeviceInputNodeResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDeviceInputNodeWithFormatAndEmitterOnDeviceAsync, uintptr(unsafe.Pointer(this)), uintptr(category), uintptr(unsafe.Pointer(encodingProperties)), uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(emitter)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioGraph2) CreateFileInputNodeWithEmitterAsync(file *IStorageFile, emitter *IAudioNodeEmitter) *IAsyncOperation[*ICreateAudioFileInputNodeResult] {
	var _result *IAsyncOperation[*ICreateAudioFileInputNodeResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFileInputNodeWithEmitterAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(emitter)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioGraph2) CreateSubmixNodeWithFormatAndEmitter(encodingProperties *IAudioEncodingProperties, emitter *IAudioNodeEmitter) *IAudioInputNode {
	var _result *IAudioInputNode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSubmixNodeWithFormatAndEmitter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(encodingProperties)), uintptr(unsafe.Pointer(emitter)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioGraph2) CreateBatchUpdater() *IClosable {
	var _result *IClosable
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateBatchUpdater, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DDCD25AE-1185-42A7-831D-6A9B0FC86820
var IID_IAudioGraph3 = syscall.GUID{0xDDCD25AE, 0x1185, 0x42A7,
	[8]byte{0x83, 0x1D, 0x6A, 0x9B, 0x0F, 0xC8, 0x68, 0x20}}

type IAudioGraph3Interface interface {
	win32.IInspectableInterface
	CreateMediaSourceAudioInputNodeAsync(mediaSource *IMediaSource2) *IAsyncOperation[*ICreateMediaSourceAudioInputNodeResult]
	CreateMediaSourceAudioInputNodeWithEmitterAsync(mediaSource *IMediaSource2, emitter *IAudioNodeEmitter) *IAsyncOperation[*ICreateMediaSourceAudioInputNodeResult]
}

type IAudioGraph3Vtbl struct {
	win32.IInspectableVtbl
	CreateMediaSourceAudioInputNodeAsync            uintptr
	CreateMediaSourceAudioInputNodeWithEmitterAsync uintptr
}

type IAudioGraph3 struct {
	win32.IInspectable
}

func (this *IAudioGraph3) Vtbl() *IAudioGraph3Vtbl {
	return (*IAudioGraph3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioGraph3) CreateMediaSourceAudioInputNodeAsync(mediaSource *IMediaSource2) *IAsyncOperation[*ICreateMediaSourceAudioInputNodeResult] {
	var _result *IAsyncOperation[*ICreateMediaSourceAudioInputNodeResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateMediaSourceAudioInputNodeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(mediaSource)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioGraph3) CreateMediaSourceAudioInputNodeWithEmitterAsync(mediaSource *IMediaSource2, emitter *IAudioNodeEmitter) *IAsyncOperation[*ICreateMediaSourceAudioInputNodeResult] {
	var _result *IAsyncOperation[*ICreateMediaSourceAudioInputNodeResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateMediaSourceAudioInputNodeWithEmitterAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(mediaSource)), uintptr(unsafe.Pointer(emitter)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 763070ED-D04E-4FAC-B233-600B42EDD469
var IID_IAudioGraphConnection = syscall.GUID{0x763070ED, 0xD04E, 0x4FAC,
	[8]byte{0xB2, 0x33, 0x60, 0x0B, 0x42, 0xED, 0xD4, 0x69}}

type IAudioGraphConnectionInterface interface {
	win32.IInspectableInterface
	Get_Destination() *IAudioNode
	Put_Gain(value float64)
	Get_Gain() float64
}

type IAudioGraphConnectionVtbl struct {
	win32.IInspectableVtbl
	Get_Destination uintptr
	Put_Gain        uintptr
	Get_Gain        uintptr
}

type IAudioGraphConnection struct {
	win32.IInspectable
}

func (this *IAudioGraphConnection) Vtbl() *IAudioGraphConnectionVtbl {
	return (*IAudioGraphConnectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioGraphConnection) Get_Destination() *IAudioNode {
	var _result *IAudioNode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Destination, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioGraphConnection) Put_Gain(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Gain, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAudioGraphConnection) Get_Gain() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Gain, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 1D59647F-E6FE-4628-84F8-9D8BDBA25785
var IID_IAudioGraphSettings = syscall.GUID{0x1D59647F, 0xE6FE, 0x4628,
	[8]byte{0x84, 0xF8, 0x9D, 0x8B, 0xDB, 0xA2, 0x57, 0x85}}

type IAudioGraphSettingsInterface interface {
	win32.IInspectableInterface
	Get_EncodingProperties() *IAudioEncodingProperties
	Put_EncodingProperties(value *IAudioEncodingProperties)
	Get_PrimaryRenderDevice() *IDeviceInformation
	Put_PrimaryRenderDevice(value *IDeviceInformation)
	Get_QuantumSizeSelectionMode() QuantumSizeSelectionMode
	Put_QuantumSizeSelectionMode(value QuantumSizeSelectionMode)
	Get_DesiredSamplesPerQuantum() int32
	Put_DesiredSamplesPerQuantum(value int32)
	Get_AudioRenderCategory() unsafe.Pointer
	Put_AudioRenderCategory(value unsafe.Pointer)
	Get_DesiredRenderDeviceAudioProcessing() AudioProcessing
	Put_DesiredRenderDeviceAudioProcessing(value AudioProcessing)
}

type IAudioGraphSettingsVtbl struct {
	win32.IInspectableVtbl
	Get_EncodingProperties                 uintptr
	Put_EncodingProperties                 uintptr
	Get_PrimaryRenderDevice                uintptr
	Put_PrimaryRenderDevice                uintptr
	Get_QuantumSizeSelectionMode           uintptr
	Put_QuantumSizeSelectionMode           uintptr
	Get_DesiredSamplesPerQuantum           uintptr
	Put_DesiredSamplesPerQuantum           uintptr
	Get_AudioRenderCategory                uintptr
	Put_AudioRenderCategory                uintptr
	Get_DesiredRenderDeviceAudioProcessing uintptr
	Put_DesiredRenderDeviceAudioProcessing uintptr
}

type IAudioGraphSettings struct {
	win32.IInspectable
}

func (this *IAudioGraphSettings) Vtbl() *IAudioGraphSettingsVtbl {
	return (*IAudioGraphSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioGraphSettings) Get_EncodingProperties() *IAudioEncodingProperties {
	var _result *IAudioEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EncodingProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioGraphSettings) Put_EncodingProperties(value *IAudioEncodingProperties) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_EncodingProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAudioGraphSettings) Get_PrimaryRenderDevice() *IDeviceInformation {
	var _result *IDeviceInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PrimaryRenderDevice, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioGraphSettings) Put_PrimaryRenderDevice(value *IDeviceInformation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PrimaryRenderDevice, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAudioGraphSettings) Get_QuantumSizeSelectionMode() QuantumSizeSelectionMode {
	var _result QuantumSizeSelectionMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_QuantumSizeSelectionMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioGraphSettings) Put_QuantumSizeSelectionMode(value QuantumSizeSelectionMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_QuantumSizeSelectionMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAudioGraphSettings) Get_DesiredSamplesPerQuantum() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesiredSamplesPerQuantum, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioGraphSettings) Put_DesiredSamplesPerQuantum(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DesiredSamplesPerQuantum, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAudioGraphSettings) Get_AudioRenderCategory() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioRenderCategory, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioGraphSettings) Put_AudioRenderCategory(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AudioRenderCategory, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAudioGraphSettings) Get_DesiredRenderDeviceAudioProcessing() AudioProcessing {
	var _result AudioProcessing
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesiredRenderDeviceAudioProcessing, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioGraphSettings) Put_DesiredRenderDeviceAudioProcessing(value AudioProcessing) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DesiredRenderDeviceAudioProcessing, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 72919787-4DAB-46E3-B4C9-D8E1A2636062
var IID_IAudioGraphSettings2 = syscall.GUID{0x72919787, 0x4DAB, 0x46E3,
	[8]byte{0xB4, 0xC9, 0xD8, 0xE1, 0xA2, 0x63, 0x60, 0x62}}

type IAudioGraphSettings2Interface interface {
	win32.IInspectableInterface
	Put_MaxPlaybackSpeedFactor(value float64)
	Get_MaxPlaybackSpeedFactor() float64
}

type IAudioGraphSettings2Vtbl struct {
	win32.IInspectableVtbl
	Put_MaxPlaybackSpeedFactor uintptr
	Get_MaxPlaybackSpeedFactor uintptr
}

type IAudioGraphSettings2 struct {
	win32.IInspectable
}

func (this *IAudioGraphSettings2) Vtbl() *IAudioGraphSettings2Vtbl {
	return (*IAudioGraphSettings2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioGraphSettings2) Put_MaxPlaybackSpeedFactor(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxPlaybackSpeedFactor, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAudioGraphSettings2) Get_MaxPlaybackSpeedFactor() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxPlaybackSpeedFactor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A5D91CC6-C2EB-4A61-A214-1D66D75F83DA
var IID_IAudioGraphSettingsFactory = syscall.GUID{0xA5D91CC6, 0xC2EB, 0x4A61,
	[8]byte{0xA2, 0x14, 0x1D, 0x66, 0xD7, 0x5F, 0x83, 0xDA}}

type IAudioGraphSettingsFactoryInterface interface {
	win32.IInspectableInterface
	Create(audioRenderCategory unsafe.Pointer) *IAudioGraphSettings
}

type IAudioGraphSettingsFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IAudioGraphSettingsFactory struct {
	win32.IInspectable
}

func (this *IAudioGraphSettingsFactory) Vtbl() *IAudioGraphSettingsFactoryVtbl {
	return (*IAudioGraphSettingsFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioGraphSettingsFactory) Create(audioRenderCategory unsafe.Pointer) *IAudioGraphSettings {
	var _result *IAudioGraphSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(audioRenderCategory), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 76EC3132-E159-4AB7-A82A-17BEB4B31E94
var IID_IAudioGraphStatics = syscall.GUID{0x76EC3132, 0xE159, 0x4AB7,
	[8]byte{0xA8, 0x2A, 0x17, 0xBE, 0xB4, 0xB3, 0x1E, 0x94}}

type IAudioGraphStaticsInterface interface {
	win32.IInspectableInterface
	CreateAsync(settings *IAudioGraphSettings) *IAsyncOperation[*ICreateAudioGraphResult]
}

type IAudioGraphStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateAsync uintptr
}

type IAudioGraphStatics struct {
	win32.IInspectable
}

func (this *IAudioGraphStatics) Vtbl() *IAudioGraphStaticsVtbl {
	return (*IAudioGraphStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioGraphStatics) CreateAsync(settings *IAudioGraphSettings) *IAsyncOperation[*ICreateAudioGraphResult] {
	var _result *IAsyncOperation[*ICreateAudioGraphResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(settings)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C3D9CBE0-3FF6-4FB3-B262-50D435C55423
var IID_IAudioGraphUnrecoverableErrorOccurredEventArgs = syscall.GUID{0xC3D9CBE0, 0x3FF6, 0x4FB3,
	[8]byte{0xB2, 0x62, 0x50, 0xD4, 0x35, 0xC5, 0x54, 0x23}}

type IAudioGraphUnrecoverableErrorOccurredEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Error() AudioGraphUnrecoverableError
}

type IAudioGraphUnrecoverableErrorOccurredEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Error uintptr
}

type IAudioGraphUnrecoverableErrorOccurredEventArgs struct {
	win32.IInspectable
}

func (this *IAudioGraphUnrecoverableErrorOccurredEventArgs) Vtbl() *IAudioGraphUnrecoverableErrorOccurredEventArgsVtbl {
	return (*IAudioGraphUnrecoverableErrorOccurredEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioGraphUnrecoverableErrorOccurredEventArgs) Get_Error() AudioGraphUnrecoverableError {
	var _result AudioGraphUnrecoverableError
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Error, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D148005C-8428-4784-B7FD-A99D468C5D20
var IID_IAudioInputNode = syscall.GUID{0xD148005C, 0x8428, 0x4784,
	[8]byte{0xB7, 0xFD, 0xA9, 0x9D, 0x46, 0x8C, 0x5D, 0x20}}

type IAudioInputNodeInterface interface {
	win32.IInspectableInterface
	Get_OutgoingConnections() *IVectorView[*IAudioGraphConnection]
	AddOutgoingConnection(destination *IAudioNode)
	AddOutgoingConnectionWithGain(destination *IAudioNode, gain float64)
	RemoveOutgoingConnection(destination *IAudioNode)
}

type IAudioInputNodeVtbl struct {
	win32.IInspectableVtbl
	Get_OutgoingConnections       uintptr
	AddOutgoingConnection         uintptr
	AddOutgoingConnectionWithGain uintptr
	RemoveOutgoingConnection      uintptr
}

type IAudioInputNode struct {
	win32.IInspectable
}

func (this *IAudioInputNode) Vtbl() *IAudioInputNodeVtbl {
	return (*IAudioInputNodeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioInputNode) Get_OutgoingConnections() *IVectorView[*IAudioGraphConnection] {
	var _result *IVectorView[*IAudioGraphConnection]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutgoingConnections, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioInputNode) AddOutgoingConnection(destination *IAudioNode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddOutgoingConnection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(destination)))
	_ = _hr
}

func (this *IAudioInputNode) AddOutgoingConnectionWithGain(destination *IAudioNode, gain float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddOutgoingConnectionWithGain, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(destination)), uintptr(gain))
	_ = _hr
}

func (this *IAudioInputNode) RemoveOutgoingConnection(destination *IAudioNode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveOutgoingConnection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(destination)))
	_ = _hr
}

// 905156B7-CA68-4C6D-A8BC-E3EE17FE3FD2
var IID_IAudioInputNode2 = syscall.GUID{0x905156B7, 0xCA68, 0x4C6D,
	[8]byte{0xA8, 0xBC, 0xE3, 0xEE, 0x17, 0xFE, 0x3F, 0xD2}}

type IAudioInputNode2Interface interface {
	win32.IInspectableInterface
	Get_Emitter() *IAudioNodeEmitter
}

type IAudioInputNode2Vtbl struct {
	win32.IInspectableVtbl
	Get_Emitter uintptr
}

type IAudioInputNode2 struct {
	win32.IInspectable
}

func (this *IAudioInputNode2) Vtbl() *IAudioInputNode2Vtbl {
	return (*IAudioInputNode2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioInputNode2) Get_Emitter() *IAudioNodeEmitter {
	var _result *IAudioNodeEmitter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Emitter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 15389D7F-DBD8-4819-BF03-668E9357CD6D
var IID_IAudioNode = syscall.GUID{0x15389D7F, 0xDBD8, 0x4819,
	[8]byte{0xBF, 0x03, 0x66, 0x8E, 0x93, 0x57, 0xCD, 0x6D}}

type IAudioNodeInterface interface {
	win32.IInspectableInterface
	Get_EffectDefinitions() *IVector[*IAudioEffectDefinition]
	Put_OutgoingGain(value float64)
	Get_OutgoingGain() float64
	Get_EncodingProperties() *IAudioEncodingProperties
	Get_ConsumeInput() bool
	Put_ConsumeInput(value bool)
	Start()
	Stop()
	Reset()
	DisableEffectsByDefinition(definition *IAudioEffectDefinition)
	EnableEffectsByDefinition(definition *IAudioEffectDefinition)
}

type IAudioNodeVtbl struct {
	win32.IInspectableVtbl
	Get_EffectDefinitions      uintptr
	Put_OutgoingGain           uintptr
	Get_OutgoingGain           uintptr
	Get_EncodingProperties     uintptr
	Get_ConsumeInput           uintptr
	Put_ConsumeInput           uintptr
	Start                      uintptr
	Stop                       uintptr
	Reset                      uintptr
	DisableEffectsByDefinition uintptr
	EnableEffectsByDefinition  uintptr
}

type IAudioNode struct {
	win32.IInspectable
}

func (this *IAudioNode) Vtbl() *IAudioNodeVtbl {
	return (*IAudioNodeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioNode) Get_EffectDefinitions() *IVector[*IAudioEffectDefinition] {
	var _result *IVector[*IAudioEffectDefinition]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EffectDefinitions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioNode) Put_OutgoingGain(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_OutgoingGain, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAudioNode) Get_OutgoingGain() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutgoingGain, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioNode) Get_EncodingProperties() *IAudioEncodingProperties {
	var _result *IAudioEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EncodingProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioNode) Get_ConsumeInput() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConsumeInput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioNode) Put_ConsumeInput(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ConsumeInput, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAudioNode) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IAudioNode) Stop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IAudioNode) Reset() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Reset, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IAudioNode) DisableEffectsByDefinition(definition *IAudioEffectDefinition) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DisableEffectsByDefinition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(definition)))
	_ = _hr
}

func (this *IAudioNode) EnableEffectsByDefinition(definition *IAudioEffectDefinition) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnableEffectsByDefinition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(definition)))
	_ = _hr
}

// 3676971D-880A-47B8-ADF7-1323A9D965BE
var IID_IAudioNodeEmitter = syscall.GUID{0x3676971D, 0x880A, 0x47B8,
	[8]byte{0xAD, 0xF7, 0x13, 0x23, 0xA9, 0xD9, 0x65, 0xBE}}

type IAudioNodeEmitterInterface interface {
	win32.IInspectableInterface
	Get_Position() unsafe.Pointer
	Put_Position(value unsafe.Pointer)
	Get_Direction() unsafe.Pointer
	Put_Direction(value unsafe.Pointer)
	Get_Shape() *IAudioNodeEmitterShape
	Get_DecayModel() *IAudioNodeEmitterDecayModel
	Get_Gain() float64
	Put_Gain(value float64)
	Get_DistanceScale() float64
	Put_DistanceScale(value float64)
	Get_DopplerScale() float64
	Put_DopplerScale(value float64)
	Get_DopplerVelocity() unsafe.Pointer
	Put_DopplerVelocity(value unsafe.Pointer)
	Get_IsDopplerDisabled() bool
}

type IAudioNodeEmitterVtbl struct {
	win32.IInspectableVtbl
	Get_Position          uintptr
	Put_Position          uintptr
	Get_Direction         uintptr
	Put_Direction         uintptr
	Get_Shape             uintptr
	Get_DecayModel        uintptr
	Get_Gain              uintptr
	Put_Gain              uintptr
	Get_DistanceScale     uintptr
	Put_DistanceScale     uintptr
	Get_DopplerScale      uintptr
	Put_DopplerScale      uintptr
	Get_DopplerVelocity   uintptr
	Put_DopplerVelocity   uintptr
	Get_IsDopplerDisabled uintptr
}

type IAudioNodeEmitter struct {
	win32.IInspectable
}

func (this *IAudioNodeEmitter) Vtbl() *IAudioNodeEmitterVtbl {
	return (*IAudioNodeEmitterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioNodeEmitter) Get_Position() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioNodeEmitter) Put_Position(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Position, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAudioNodeEmitter) Get_Direction() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Direction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioNodeEmitter) Put_Direction(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Direction, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAudioNodeEmitter) Get_Shape() *IAudioNodeEmitterShape {
	var _result *IAudioNodeEmitterShape
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Shape, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioNodeEmitter) Get_DecayModel() *IAudioNodeEmitterDecayModel {
	var _result *IAudioNodeEmitterDecayModel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DecayModel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioNodeEmitter) Get_Gain() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Gain, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioNodeEmitter) Put_Gain(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Gain, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAudioNodeEmitter) Get_DistanceScale() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DistanceScale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioNodeEmitter) Put_DistanceScale(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DistanceScale, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAudioNodeEmitter) Get_DopplerScale() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DopplerScale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioNodeEmitter) Put_DopplerScale(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DopplerScale, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAudioNodeEmitter) Get_DopplerVelocity() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DopplerVelocity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioNodeEmitter) Put_DopplerVelocity(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DopplerVelocity, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAudioNodeEmitter) Get_IsDopplerDisabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDopplerDisabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4AB6EECB-EC29-47F8-818C-B6B660A5AEB1
var IID_IAudioNodeEmitter2 = syscall.GUID{0x4AB6EECB, 0xEC29, 0x47F8,
	[8]byte{0x81, 0x8C, 0xB6, 0xB6, 0x60, 0xA5, 0xAE, 0xB1}}

type IAudioNodeEmitter2Interface interface {
	win32.IInspectableInterface
	Get_SpatialAudioModel() SpatialAudioModel
	Put_SpatialAudioModel(value SpatialAudioModel)
}

type IAudioNodeEmitter2Vtbl struct {
	win32.IInspectableVtbl
	Get_SpatialAudioModel uintptr
	Put_SpatialAudioModel uintptr
}

type IAudioNodeEmitter2 struct {
	win32.IInspectable
}

func (this *IAudioNodeEmitter2) Vtbl() *IAudioNodeEmitter2Vtbl {
	return (*IAudioNodeEmitter2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioNodeEmitter2) Get_SpatialAudioModel() SpatialAudioModel {
	var _result SpatialAudioModel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SpatialAudioModel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioNodeEmitter2) Put_SpatialAudioModel(value SpatialAudioModel) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SpatialAudioModel, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// E99B2CEE-02CA-4375-9326-0C6AE4BCDFB5
var IID_IAudioNodeEmitterConeProperties = syscall.GUID{0xE99B2CEE, 0x02CA, 0x4375,
	[8]byte{0x93, 0x26, 0x0C, 0x6A, 0xE4, 0xBC, 0xDF, 0xB5}}

type IAudioNodeEmitterConePropertiesInterface interface {
	win32.IInspectableInterface
	Get_InnerAngle() float64
	Get_OuterAngle() float64
	Get_OuterAngleGain() float64
}

type IAudioNodeEmitterConePropertiesVtbl struct {
	win32.IInspectableVtbl
	Get_InnerAngle     uintptr
	Get_OuterAngle     uintptr
	Get_OuterAngleGain uintptr
}

type IAudioNodeEmitterConeProperties struct {
	win32.IInspectable
}

func (this *IAudioNodeEmitterConeProperties) Vtbl() *IAudioNodeEmitterConePropertiesVtbl {
	return (*IAudioNodeEmitterConePropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioNodeEmitterConeProperties) Get_InnerAngle() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InnerAngle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioNodeEmitterConeProperties) Get_OuterAngle() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OuterAngle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioNodeEmitterConeProperties) Get_OuterAngleGain() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OuterAngleGain, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 1D1D5AF7-0D53-4FA9-BD84-D5816A86F3FF
var IID_IAudioNodeEmitterDecayModel = syscall.GUID{0x1D1D5AF7, 0x0D53, 0x4FA9,
	[8]byte{0xBD, 0x84, 0xD5, 0x81, 0x6A, 0x86, 0xF3, 0xFF}}

type IAudioNodeEmitterDecayModelInterface interface {
	win32.IInspectableInterface
	Get_Kind() AudioNodeEmitterDecayKind
	Get_MinGain() float64
	Get_MaxGain() float64
	Get_NaturalProperties() *IAudioNodeEmitterNaturalDecayModelProperties
}

type IAudioNodeEmitterDecayModelVtbl struct {
	win32.IInspectableVtbl
	Get_Kind              uintptr
	Get_MinGain           uintptr
	Get_MaxGain           uintptr
	Get_NaturalProperties uintptr
}

type IAudioNodeEmitterDecayModel struct {
	win32.IInspectable
}

func (this *IAudioNodeEmitterDecayModel) Vtbl() *IAudioNodeEmitterDecayModelVtbl {
	return (*IAudioNodeEmitterDecayModelVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioNodeEmitterDecayModel) Get_Kind() AudioNodeEmitterDecayKind {
	var _result AudioNodeEmitterDecayKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Kind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioNodeEmitterDecayModel) Get_MinGain() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinGain, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioNodeEmitterDecayModel) Get_MaxGain() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxGain, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioNodeEmitterDecayModel) Get_NaturalProperties() *IAudioNodeEmitterNaturalDecayModelProperties {
	var _result *IAudioNodeEmitterNaturalDecayModelProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NaturalProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C7787CA8-F178-462F-BC81-8DD5CBE5DAE8
var IID_IAudioNodeEmitterDecayModelStatics = syscall.GUID{0xC7787CA8, 0xF178, 0x462F,
	[8]byte{0xBC, 0x81, 0x8D, 0xD5, 0xCB, 0xE5, 0xDA, 0xE8}}

type IAudioNodeEmitterDecayModelStaticsInterface interface {
	win32.IInspectableInterface
	CreateNatural(minGain float64, maxGain float64, unityGainDistance float64, cutoffDistance float64) *IAudioNodeEmitterDecayModel
	CreateCustom(minGain float64, maxGain float64) *IAudioNodeEmitterDecayModel
}

type IAudioNodeEmitterDecayModelStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateNatural uintptr
	CreateCustom  uintptr
}

type IAudioNodeEmitterDecayModelStatics struct {
	win32.IInspectable
}

func (this *IAudioNodeEmitterDecayModelStatics) Vtbl() *IAudioNodeEmitterDecayModelStaticsVtbl {
	return (*IAudioNodeEmitterDecayModelStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioNodeEmitterDecayModelStatics) CreateNatural(minGain float64, maxGain float64, unityGainDistance float64, cutoffDistance float64) *IAudioNodeEmitterDecayModel {
	var _result *IAudioNodeEmitterDecayModel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateNatural, uintptr(unsafe.Pointer(this)), uintptr(minGain), uintptr(maxGain), uintptr(unityGainDistance), uintptr(cutoffDistance), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioNodeEmitterDecayModelStatics) CreateCustom(minGain float64, maxGain float64) *IAudioNodeEmitterDecayModel {
	var _result *IAudioNodeEmitterDecayModel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateCustom, uintptr(unsafe.Pointer(this)), uintptr(minGain), uintptr(maxGain), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FDC8489A-6AD6-4CE4-B7F7-A99370DF7EE9
var IID_IAudioNodeEmitterFactory = syscall.GUID{0xFDC8489A, 0x6AD6, 0x4CE4,
	[8]byte{0xB7, 0xF7, 0xA9, 0x93, 0x70, 0xDF, 0x7E, 0xE9}}

type IAudioNodeEmitterFactoryInterface interface {
	win32.IInspectableInterface
	CreateAudioNodeEmitter(shape *IAudioNodeEmitterShape, decayModel *IAudioNodeEmitterDecayModel, settings AudioNodeEmitterSettings) *IAudioNodeEmitter
}

type IAudioNodeEmitterFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateAudioNodeEmitter uintptr
}

type IAudioNodeEmitterFactory struct {
	win32.IInspectable
}

func (this *IAudioNodeEmitterFactory) Vtbl() *IAudioNodeEmitterFactoryVtbl {
	return (*IAudioNodeEmitterFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioNodeEmitterFactory) CreateAudioNodeEmitter(shape *IAudioNodeEmitterShape, decayModel *IAudioNodeEmitterDecayModel, settings AudioNodeEmitterSettings) *IAudioNodeEmitter {
	var _result *IAudioNodeEmitter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAudioNodeEmitter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(decayModel)), uintptr(settings), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 48934BCF-CF2C-4EFC-9331-75BD22DF1F0C
var IID_IAudioNodeEmitterNaturalDecayModelProperties = syscall.GUID{0x48934BCF, 0xCF2C, 0x4EFC,
	[8]byte{0x93, 0x31, 0x75, 0xBD, 0x22, 0xDF, 0x1F, 0x0C}}

type IAudioNodeEmitterNaturalDecayModelPropertiesInterface interface {
	win32.IInspectableInterface
	Get_UnityGainDistance() float64
	Get_CutoffDistance() float64
}

type IAudioNodeEmitterNaturalDecayModelPropertiesVtbl struct {
	win32.IInspectableVtbl
	Get_UnityGainDistance uintptr
	Get_CutoffDistance    uintptr
}

type IAudioNodeEmitterNaturalDecayModelProperties struct {
	win32.IInspectable
}

func (this *IAudioNodeEmitterNaturalDecayModelProperties) Vtbl() *IAudioNodeEmitterNaturalDecayModelPropertiesVtbl {
	return (*IAudioNodeEmitterNaturalDecayModelPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioNodeEmitterNaturalDecayModelProperties) Get_UnityGainDistance() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UnityGainDistance, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioNodeEmitterNaturalDecayModelProperties) Get_CutoffDistance() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CutoffDistance, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// EA0311C5-E73D-44BC-859C-45553BBC4828
var IID_IAudioNodeEmitterShape = syscall.GUID{0xEA0311C5, 0xE73D, 0x44BC,
	[8]byte{0x85, 0x9C, 0x45, 0x55, 0x3B, 0xBC, 0x48, 0x28}}

type IAudioNodeEmitterShapeInterface interface {
	win32.IInspectableInterface
	Get_Kind() AudioNodeEmitterShapeKind
	Get_ConeProperties() *IAudioNodeEmitterConeProperties
}

type IAudioNodeEmitterShapeVtbl struct {
	win32.IInspectableVtbl
	Get_Kind           uintptr
	Get_ConeProperties uintptr
}

type IAudioNodeEmitterShape struct {
	win32.IInspectable
}

func (this *IAudioNodeEmitterShape) Vtbl() *IAudioNodeEmitterShapeVtbl {
	return (*IAudioNodeEmitterShapeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioNodeEmitterShape) Get_Kind() AudioNodeEmitterShapeKind {
	var _result AudioNodeEmitterShapeKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Kind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioNodeEmitterShape) Get_ConeProperties() *IAudioNodeEmitterConeProperties {
	var _result *IAudioNodeEmitterConeProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConeProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 57BB2771-FFA5-4B86-A779-E264AEB9145F
var IID_IAudioNodeEmitterShapeStatics = syscall.GUID{0x57BB2771, 0xFFA5, 0x4B86,
	[8]byte{0xA7, 0x79, 0xE2, 0x64, 0xAE, 0xB9, 0x14, 0x5F}}

type IAudioNodeEmitterShapeStaticsInterface interface {
	win32.IInspectableInterface
	CreateCone(innerAngle float64, outerAngle float64, outerAngleGain float64) *IAudioNodeEmitterShape
	CreateOmnidirectional() *IAudioNodeEmitterShape
}

type IAudioNodeEmitterShapeStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateCone            uintptr
	CreateOmnidirectional uintptr
}

type IAudioNodeEmitterShapeStatics struct {
	win32.IInspectable
}

func (this *IAudioNodeEmitterShapeStatics) Vtbl() *IAudioNodeEmitterShapeStaticsVtbl {
	return (*IAudioNodeEmitterShapeStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioNodeEmitterShapeStatics) CreateCone(innerAngle float64, outerAngle float64, outerAngleGain float64) *IAudioNodeEmitterShape {
	var _result *IAudioNodeEmitterShape
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateCone, uintptr(unsafe.Pointer(this)), uintptr(innerAngle), uintptr(outerAngle), uintptr(outerAngleGain), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioNodeEmitterShapeStatics) CreateOmnidirectional() *IAudioNodeEmitterShape {
	var _result *IAudioNodeEmitterShape
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateOmnidirectional, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D9722E16-0C0A-41DA-B755-6C77835FB1EB
var IID_IAudioNodeListener = syscall.GUID{0xD9722E16, 0x0C0A, 0x41DA,
	[8]byte{0xB7, 0x55, 0x6C, 0x77, 0x83, 0x5F, 0xB1, 0xEB}}

type IAudioNodeListenerInterface interface {
	win32.IInspectableInterface
	Get_Position() unsafe.Pointer
	Put_Position(value unsafe.Pointer)
	Get_Orientation() unsafe.Pointer
	Put_Orientation(value unsafe.Pointer)
	Get_SpeedOfSound() float64
	Put_SpeedOfSound(value float64)
	Get_DopplerVelocity() unsafe.Pointer
	Put_DopplerVelocity(value unsafe.Pointer)
}

type IAudioNodeListenerVtbl struct {
	win32.IInspectableVtbl
	Get_Position        uintptr
	Put_Position        uintptr
	Get_Orientation     uintptr
	Put_Orientation     uintptr
	Get_SpeedOfSound    uintptr
	Put_SpeedOfSound    uintptr
	Get_DopplerVelocity uintptr
	Put_DopplerVelocity uintptr
}

type IAudioNodeListener struct {
	win32.IInspectable
}

func (this *IAudioNodeListener) Vtbl() *IAudioNodeListenerVtbl {
	return (*IAudioNodeListenerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioNodeListener) Get_Position() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioNodeListener) Put_Position(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Position, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAudioNodeListener) Get_Orientation() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Orientation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioNodeListener) Put_Orientation(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Orientation, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAudioNodeListener) Get_SpeedOfSound() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SpeedOfSound, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioNodeListener) Put_SpeedOfSound(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SpeedOfSound, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAudioNodeListener) Get_DopplerVelocity() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DopplerVelocity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioNodeListener) Put_DopplerVelocity(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DopplerVelocity, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 0E0F907C-79FF-4544-9EEB-01257B15105A
var IID_IAudioNodeWithListener = syscall.GUID{0x0E0F907C, 0x79FF, 0x4544,
	[8]byte{0x9E, 0xEB, 0x01, 0x25, 0x7B, 0x15, 0x10, 0x5A}}

type IAudioNodeWithListenerInterface interface {
	win32.IInspectableInterface
	Put_Listener(value *IAudioNodeListener)
	Get_Listener() *IAudioNodeListener
}

type IAudioNodeWithListenerVtbl struct {
	win32.IInspectableVtbl
	Put_Listener uintptr
	Get_Listener uintptr
}

type IAudioNodeWithListener struct {
	win32.IInspectable
}

func (this *IAudioNodeWithListener) Vtbl() *IAudioNodeWithListenerVtbl {
	return (*IAudioNodeWithListenerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioNodeWithListener) Put_Listener(value *IAudioNodeListener) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Listener, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAudioNodeWithListener) Get_Listener() *IAudioNodeListener {
	var _result *IAudioNodeListener
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Listener, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1A4C1DEA-CAFC-50E7-8718-EA3F81CBFA51
var IID_IAudioPlaybackConnection = syscall.GUID{0x1A4C1DEA, 0xCAFC, 0x50E7,
	[8]byte{0x87, 0x18, 0xEA, 0x3F, 0x81, 0xCB, 0xFA, 0x51}}

type IAudioPlaybackConnectionInterface interface {
	win32.IInspectableInterface
	Start()
	StartAsync() *IAsyncAction
	Get_DeviceId() string
	Get_State() AudioPlaybackConnectionState
	Open() *IAudioPlaybackConnectionOpenResult
	OpenAsync() *IAsyncOperation[*IAudioPlaybackConnectionOpenResult]
	Add_StateChanged(handler TypedEventHandler[*IAudioPlaybackConnection, interface{}]) EventRegistrationToken
	Remove_StateChanged(token EventRegistrationToken)
}

type IAudioPlaybackConnectionVtbl struct {
	win32.IInspectableVtbl
	Start               uintptr
	StartAsync          uintptr
	Get_DeviceId        uintptr
	Get_State           uintptr
	Open                uintptr
	OpenAsync           uintptr
	Add_StateChanged    uintptr
	Remove_StateChanged uintptr
}

type IAudioPlaybackConnection struct {
	win32.IInspectable
}

func (this *IAudioPlaybackConnection) Vtbl() *IAudioPlaybackConnectionVtbl {
	return (*IAudioPlaybackConnectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioPlaybackConnection) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IAudioPlaybackConnection) StartAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioPlaybackConnection) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAudioPlaybackConnection) Get_State() AudioPlaybackConnectionState {
	var _result AudioPlaybackConnectionState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioPlaybackConnection) Open() *IAudioPlaybackConnectionOpenResult {
	var _result *IAudioPlaybackConnectionOpenResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Open, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioPlaybackConnection) OpenAsync() *IAsyncOperation[*IAudioPlaybackConnectionOpenResult] {
	var _result *IAsyncOperation[*IAudioPlaybackConnectionOpenResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioPlaybackConnection) Add_StateChanged(handler TypedEventHandler[*IAudioPlaybackConnection, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioPlaybackConnection) Remove_StateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 4E656AEF-39F9-5FC9-A519-A5BBFD9FE921
var IID_IAudioPlaybackConnectionOpenResult = syscall.GUID{0x4E656AEF, 0x39F9, 0x5FC9,
	[8]byte{0xA5, 0x19, 0xA5, 0xBB, 0xFD, 0x9F, 0xE9, 0x21}}

type IAudioPlaybackConnectionOpenResultInterface interface {
	win32.IInspectableInterface
	Get_Status() AudioPlaybackConnectionOpenResultStatus
	Get_ExtendedError() HResult
}

type IAudioPlaybackConnectionOpenResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status        uintptr
	Get_ExtendedError uintptr
}

type IAudioPlaybackConnectionOpenResult struct {
	win32.IInspectable
}

func (this *IAudioPlaybackConnectionOpenResult) Vtbl() *IAudioPlaybackConnectionOpenResultVtbl {
	return (*IAudioPlaybackConnectionOpenResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioPlaybackConnectionOpenResult) Get_Status() AudioPlaybackConnectionOpenResultStatus {
	var _result AudioPlaybackConnectionOpenResultStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioPlaybackConnectionOpenResult) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// E60963A2-69E6-5FFC-9E13-824A85213DAF
var IID_IAudioPlaybackConnectionStatics = syscall.GUID{0xE60963A2, 0x69E6, 0x5FFC,
	[8]byte{0x9E, 0x13, 0x82, 0x4A, 0x85, 0x21, 0x3D, 0xAF}}

type IAudioPlaybackConnectionStaticsInterface interface {
	win32.IInspectableInterface
	GetDeviceSelector() string
	TryCreateFromId(id string) *IAudioPlaybackConnection
}

type IAudioPlaybackConnectionStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelector uintptr
	TryCreateFromId   uintptr
}

type IAudioPlaybackConnectionStatics struct {
	win32.IInspectable
}

func (this *IAudioPlaybackConnectionStatics) Vtbl() *IAudioPlaybackConnectionStaticsVtbl {
	return (*IAudioPlaybackConnectionStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioPlaybackConnectionStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAudioPlaybackConnectionStatics) TryCreateFromId(id string) *IAudioPlaybackConnection {
	var _result *IAudioPlaybackConnection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryCreateFromId, uintptr(unsafe.Pointer(this)), NewHStr(id).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1D13D136-0199-4CDC-B84E-E72C2B581ECE
var IID_IAudioStateMonitor = syscall.GUID{0x1D13D136, 0x0199, 0x4CDC,
	[8]byte{0xB8, 0x4E, 0xE7, 0x2C, 0x2B, 0x58, 0x1E, 0xCE}}

type IAudioStateMonitorInterface interface {
	win32.IInspectableInterface
	Add_SoundLevelChanged(handler TypedEventHandler[*IAudioStateMonitor, interface{}]) EventRegistrationToken
	Remove_SoundLevelChanged(token EventRegistrationToken)
	Get_SoundLevel() SoundLevel
}

type IAudioStateMonitorVtbl struct {
	win32.IInspectableVtbl
	Add_SoundLevelChanged    uintptr
	Remove_SoundLevelChanged uintptr
	Get_SoundLevel           uintptr
}

type IAudioStateMonitor struct {
	win32.IInspectable
}

func (this *IAudioStateMonitor) Vtbl() *IAudioStateMonitorVtbl {
	return (*IAudioStateMonitorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioStateMonitor) Add_SoundLevelChanged(handler TypedEventHandler[*IAudioStateMonitor, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SoundLevelChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioStateMonitor) Remove_SoundLevelChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SoundLevelChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAudioStateMonitor) Get_SoundLevel() SoundLevel {
	var _result SoundLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SoundLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 6374EA4C-1B3B-4001-94D9-DD225330FA40
var IID_IAudioStateMonitorStatics = syscall.GUID{0x6374EA4C, 0x1B3B, 0x4001,
	[8]byte{0x94, 0xD9, 0xDD, 0x22, 0x53, 0x30, 0xFA, 0x40}}

type IAudioStateMonitorStaticsInterface interface {
	win32.IInspectableInterface
	CreateForRenderMonitoring() *IAudioStateMonitor
	CreateForRenderMonitoringWithCategory(category unsafe.Pointer) *IAudioStateMonitor
	CreateForRenderMonitoringWithCategoryAndDeviceRole(category unsafe.Pointer, role AudioDeviceRole) *IAudioStateMonitor
	CreateForRenderMonitoringWithCategoryAndDeviceId(category unsafe.Pointer, deviceId string) *IAudioStateMonitor
	CreateForCaptureMonitoring() *IAudioStateMonitor
	CreateForCaptureMonitoringWithCategory(category MediaCategory) *IAudioStateMonitor
	CreateForCaptureMonitoringWithCategoryAndDeviceRole(category MediaCategory, role AudioDeviceRole) *IAudioStateMonitor
	CreateForCaptureMonitoringWithCategoryAndDeviceId(category MediaCategory, deviceId string) *IAudioStateMonitor
}

type IAudioStateMonitorStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateForRenderMonitoring                           uintptr
	CreateForRenderMonitoringWithCategory               uintptr
	CreateForRenderMonitoringWithCategoryAndDeviceRole  uintptr
	CreateForRenderMonitoringWithCategoryAndDeviceId    uintptr
	CreateForCaptureMonitoring                          uintptr
	CreateForCaptureMonitoringWithCategory              uintptr
	CreateForCaptureMonitoringWithCategoryAndDeviceRole uintptr
	CreateForCaptureMonitoringWithCategoryAndDeviceId   uintptr
}

type IAudioStateMonitorStatics struct {
	win32.IInspectable
}

func (this *IAudioStateMonitorStatics) Vtbl() *IAudioStateMonitorStaticsVtbl {
	return (*IAudioStateMonitorStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioStateMonitorStatics) CreateForRenderMonitoring() *IAudioStateMonitor {
	var _result *IAudioStateMonitor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateForRenderMonitoring, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioStateMonitorStatics) CreateForRenderMonitoringWithCategory(category unsafe.Pointer) *IAudioStateMonitor {
	var _result *IAudioStateMonitor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateForRenderMonitoringWithCategory, uintptr(unsafe.Pointer(this)), uintptr(category), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioStateMonitorStatics) CreateForRenderMonitoringWithCategoryAndDeviceRole(category unsafe.Pointer, role AudioDeviceRole) *IAudioStateMonitor {
	var _result *IAudioStateMonitor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateForRenderMonitoringWithCategoryAndDeviceRole, uintptr(unsafe.Pointer(this)), uintptr(category), uintptr(role), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioStateMonitorStatics) CreateForRenderMonitoringWithCategoryAndDeviceId(category unsafe.Pointer, deviceId string) *IAudioStateMonitor {
	var _result *IAudioStateMonitor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateForRenderMonitoringWithCategoryAndDeviceId, uintptr(unsafe.Pointer(this)), uintptr(category), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioStateMonitorStatics) CreateForCaptureMonitoring() *IAudioStateMonitor {
	var _result *IAudioStateMonitor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateForCaptureMonitoring, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioStateMonitorStatics) CreateForCaptureMonitoringWithCategory(category MediaCategory) *IAudioStateMonitor {
	var _result *IAudioStateMonitor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateForCaptureMonitoringWithCategory, uintptr(unsafe.Pointer(this)), uintptr(category), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioStateMonitorStatics) CreateForCaptureMonitoringWithCategoryAndDeviceRole(category MediaCategory, role AudioDeviceRole) *IAudioStateMonitor {
	var _result *IAudioStateMonitor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateForCaptureMonitoringWithCategoryAndDeviceRole, uintptr(unsafe.Pointer(this)), uintptr(category), uintptr(role), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioStateMonitorStatics) CreateForCaptureMonitoringWithCategoryAndDeviceId(category MediaCategory, deviceId string) *IAudioStateMonitor {
	var _result *IAudioStateMonitor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateForCaptureMonitoringWithCategoryAndDeviceId, uintptr(unsafe.Pointer(this)), uintptr(category), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 16EEC7A8-1CA7-40EF-91A4-D346E0AA1BBA
var IID_ICreateAudioDeviceInputNodeResult = syscall.GUID{0x16EEC7A8, 0x1CA7, 0x40EF,
	[8]byte{0x91, 0xA4, 0xD3, 0x46, 0xE0, 0xAA, 0x1B, 0xBA}}

type ICreateAudioDeviceInputNodeResultInterface interface {
	win32.IInspectableInterface
	Get_Status() AudioDeviceNodeCreationStatus
	Get_DeviceInputNode() *IAudioDeviceInputNode
}

type ICreateAudioDeviceInputNodeResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status          uintptr
	Get_DeviceInputNode uintptr
}

type ICreateAudioDeviceInputNodeResult struct {
	win32.IInspectable
}

func (this *ICreateAudioDeviceInputNodeResult) Vtbl() *ICreateAudioDeviceInputNodeResultVtbl {
	return (*ICreateAudioDeviceInputNodeResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICreateAudioDeviceInputNodeResult) Get_Status() AudioDeviceNodeCreationStatus {
	var _result AudioDeviceNodeCreationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICreateAudioDeviceInputNodeResult) Get_DeviceInputNode() *IAudioDeviceInputNode {
	var _result *IAudioDeviceInputNode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceInputNode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 921C69CE-3F35-41C7-9622-79F608BAEDC2
var IID_ICreateAudioDeviceInputNodeResult2 = syscall.GUID{0x921C69CE, 0x3F35, 0x41C7,
	[8]byte{0x96, 0x22, 0x79, 0xF6, 0x08, 0xBA, 0xED, 0xC2}}

type ICreateAudioDeviceInputNodeResult2Interface interface {
	win32.IInspectableInterface
	Get_ExtendedError() HResult
}

type ICreateAudioDeviceInputNodeResult2Vtbl struct {
	win32.IInspectableVtbl
	Get_ExtendedError uintptr
}

type ICreateAudioDeviceInputNodeResult2 struct {
	win32.IInspectable
}

func (this *ICreateAudioDeviceInputNodeResult2) Vtbl() *ICreateAudioDeviceInputNodeResult2Vtbl {
	return (*ICreateAudioDeviceInputNodeResult2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICreateAudioDeviceInputNodeResult2) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F7776D27-1D9A-47F7-9CD4-2859CC1B7BFF
var IID_ICreateAudioDeviceOutputNodeResult = syscall.GUID{0xF7776D27, 0x1D9A, 0x47F7,
	[8]byte{0x9C, 0xD4, 0x28, 0x59, 0xCC, 0x1B, 0x7B, 0xFF}}

type ICreateAudioDeviceOutputNodeResultInterface interface {
	win32.IInspectableInterface
	Get_Status() AudioDeviceNodeCreationStatus
	Get_DeviceOutputNode() *IAudioDeviceOutputNode
}

type ICreateAudioDeviceOutputNodeResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status           uintptr
	Get_DeviceOutputNode uintptr
}

type ICreateAudioDeviceOutputNodeResult struct {
	win32.IInspectable
}

func (this *ICreateAudioDeviceOutputNodeResult) Vtbl() *ICreateAudioDeviceOutputNodeResultVtbl {
	return (*ICreateAudioDeviceOutputNodeResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICreateAudioDeviceOutputNodeResult) Get_Status() AudioDeviceNodeCreationStatus {
	var _result AudioDeviceNodeCreationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICreateAudioDeviceOutputNodeResult) Get_DeviceOutputNode() *IAudioDeviceOutputNode {
	var _result *IAudioDeviceOutputNode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceOutputNode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4864269F-BDCE-4AB1-BD38-FBAE93AEDACA
var IID_ICreateAudioDeviceOutputNodeResult2 = syscall.GUID{0x4864269F, 0xBDCE, 0x4AB1,
	[8]byte{0xBD, 0x38, 0xFB, 0xAE, 0x93, 0xAE, 0xDA, 0xCA}}

type ICreateAudioDeviceOutputNodeResult2Interface interface {
	win32.IInspectableInterface
	Get_ExtendedError() HResult
}

type ICreateAudioDeviceOutputNodeResult2Vtbl struct {
	win32.IInspectableVtbl
	Get_ExtendedError uintptr
}

type ICreateAudioDeviceOutputNodeResult2 struct {
	win32.IInspectable
}

func (this *ICreateAudioDeviceOutputNodeResult2) Vtbl() *ICreateAudioDeviceOutputNodeResult2Vtbl {
	return (*ICreateAudioDeviceOutputNodeResult2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICreateAudioDeviceOutputNodeResult2) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// CE83D61C-E297-4C50-9CE7-1C7A69D6BD09
var IID_ICreateAudioFileInputNodeResult = syscall.GUID{0xCE83D61C, 0xE297, 0x4C50,
	[8]byte{0x9C, 0xE7, 0x1C, 0x7A, 0x69, 0xD6, 0xBD, 0x09}}

type ICreateAudioFileInputNodeResultInterface interface {
	win32.IInspectableInterface
	Get_Status() AudioFileNodeCreationStatus
	Get_FileInputNode() *IAudioFileInputNode
}

type ICreateAudioFileInputNodeResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status        uintptr
	Get_FileInputNode uintptr
}

type ICreateAudioFileInputNodeResult struct {
	win32.IInspectable
}

func (this *ICreateAudioFileInputNodeResult) Vtbl() *ICreateAudioFileInputNodeResultVtbl {
	return (*ICreateAudioFileInputNodeResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICreateAudioFileInputNodeResult) Get_Status() AudioFileNodeCreationStatus {
	var _result AudioFileNodeCreationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICreateAudioFileInputNodeResult) Get_FileInputNode() *IAudioFileInputNode {
	var _result *IAudioFileInputNode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FileInputNode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F9082020-3D80-4FE0-81C1-768FEA7CA7E0
var IID_ICreateAudioFileInputNodeResult2 = syscall.GUID{0xF9082020, 0x3D80, 0x4FE0,
	[8]byte{0x81, 0xC1, 0x76, 0x8F, 0xEA, 0x7C, 0xA7, 0xE0}}

type ICreateAudioFileInputNodeResult2Interface interface {
	win32.IInspectableInterface
	Get_ExtendedError() HResult
}

type ICreateAudioFileInputNodeResult2Vtbl struct {
	win32.IInspectableVtbl
	Get_ExtendedError uintptr
}

type ICreateAudioFileInputNodeResult2 struct {
	win32.IInspectable
}

func (this *ICreateAudioFileInputNodeResult2) Vtbl() *ICreateAudioFileInputNodeResult2Vtbl {
	return (*ICreateAudioFileInputNodeResult2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICreateAudioFileInputNodeResult2) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 47D6BA7B-E909-453F-866E-5540CDA734FF
var IID_ICreateAudioFileOutputNodeResult = syscall.GUID{0x47D6BA7B, 0xE909, 0x453F,
	[8]byte{0x86, 0x6E, 0x55, 0x40, 0xCD, 0xA7, 0x34, 0xFF}}

type ICreateAudioFileOutputNodeResultInterface interface {
	win32.IInspectableInterface
	Get_Status() AudioFileNodeCreationStatus
	Get_FileOutputNode() *IAudioFileOutputNode
}

type ICreateAudioFileOutputNodeResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status         uintptr
	Get_FileOutputNode uintptr
}

type ICreateAudioFileOutputNodeResult struct {
	win32.IInspectable
}

func (this *ICreateAudioFileOutputNodeResult) Vtbl() *ICreateAudioFileOutputNodeResultVtbl {
	return (*ICreateAudioFileOutputNodeResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICreateAudioFileOutputNodeResult) Get_Status() AudioFileNodeCreationStatus {
	var _result AudioFileNodeCreationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICreateAudioFileOutputNodeResult) Get_FileOutputNode() *IAudioFileOutputNode {
	var _result *IAudioFileOutputNode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FileOutputNode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9F01B50D-3318-47B3-A60A-1B492BE7FC0D
var IID_ICreateAudioFileOutputNodeResult2 = syscall.GUID{0x9F01B50D, 0x3318, 0x47B3,
	[8]byte{0xA6, 0x0A, 0x1B, 0x49, 0x2B, 0xE7, 0xFC, 0x0D}}

type ICreateAudioFileOutputNodeResult2Interface interface {
	win32.IInspectableInterface
	Get_ExtendedError() HResult
}

type ICreateAudioFileOutputNodeResult2Vtbl struct {
	win32.IInspectableVtbl
	Get_ExtendedError uintptr
}

type ICreateAudioFileOutputNodeResult2 struct {
	win32.IInspectable
}

func (this *ICreateAudioFileOutputNodeResult2) Vtbl() *ICreateAudioFileOutputNodeResult2Vtbl {
	return (*ICreateAudioFileOutputNodeResult2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICreateAudioFileOutputNodeResult2) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5453EF7E-7BDE-4B76-BB5D-48F79CFC8C0B
var IID_ICreateAudioGraphResult = syscall.GUID{0x5453EF7E, 0x7BDE, 0x4B76,
	[8]byte{0xBB, 0x5D, 0x48, 0xF7, 0x9C, 0xFC, 0x8C, 0x0B}}

type ICreateAudioGraphResultInterface interface {
	win32.IInspectableInterface
	Get_Status() AudioGraphCreationStatus
	Get_Graph() *IAudioGraph
}

type ICreateAudioGraphResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
	Get_Graph  uintptr
}

type ICreateAudioGraphResult struct {
	win32.IInspectable
}

func (this *ICreateAudioGraphResult) Vtbl() *ICreateAudioGraphResultVtbl {
	return (*ICreateAudioGraphResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICreateAudioGraphResult) Get_Status() AudioGraphCreationStatus {
	var _result AudioGraphCreationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICreateAudioGraphResult) Get_Graph() *IAudioGraph {
	var _result *IAudioGraph
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Graph, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6D738DFC-88C6-4FCB-A534-85CEDD4050A1
var IID_ICreateAudioGraphResult2 = syscall.GUID{0x6D738DFC, 0x88C6, 0x4FCB,
	[8]byte{0xA5, 0x34, 0x85, 0xCE, 0xDD, 0x40, 0x50, 0xA1}}

type ICreateAudioGraphResult2Interface interface {
	win32.IInspectableInterface
	Get_ExtendedError() HResult
}

type ICreateAudioGraphResult2Vtbl struct {
	win32.IInspectableVtbl
	Get_ExtendedError uintptr
}

type ICreateAudioGraphResult2 struct {
	win32.IInspectable
}

func (this *ICreateAudioGraphResult2) Vtbl() *ICreateAudioGraphResult2Vtbl {
	return (*ICreateAudioGraphResult2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICreateAudioGraphResult2) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 46A658A3-53C0-4D59-9E51-CC1D1044A4C4
var IID_ICreateMediaSourceAudioInputNodeResult = syscall.GUID{0x46A658A3, 0x53C0, 0x4D59,
	[8]byte{0x9E, 0x51, 0xCC, 0x1D, 0x10, 0x44, 0xA4, 0xC4}}

type ICreateMediaSourceAudioInputNodeResultInterface interface {
	win32.IInspectableInterface
	Get_Status() MediaSourceAudioInputNodeCreationStatus
	Get_Node() *IMediaSourceAudioInputNode
}

type ICreateMediaSourceAudioInputNodeResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
	Get_Node   uintptr
}

type ICreateMediaSourceAudioInputNodeResult struct {
	win32.IInspectable
}

func (this *ICreateMediaSourceAudioInputNodeResult) Vtbl() *ICreateMediaSourceAudioInputNodeResultVtbl {
	return (*ICreateMediaSourceAudioInputNodeResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICreateMediaSourceAudioInputNodeResult) Get_Status() MediaSourceAudioInputNodeCreationStatus {
	var _result MediaSourceAudioInputNodeCreationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICreateMediaSourceAudioInputNodeResult) Get_Node() *IMediaSourceAudioInputNode {
	var _result *IMediaSourceAudioInputNode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Node, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 63514CE8-6A1A-49E3-97EC-28FD5BE114E5
var IID_ICreateMediaSourceAudioInputNodeResult2 = syscall.GUID{0x63514CE8, 0x6A1A, 0x49E3,
	[8]byte{0x97, 0xEC, 0x28, 0xFD, 0x5B, 0xE1, 0x14, 0xE5}}

type ICreateMediaSourceAudioInputNodeResult2Interface interface {
	win32.IInspectableInterface
	Get_ExtendedError() HResult
}

type ICreateMediaSourceAudioInputNodeResult2Vtbl struct {
	win32.IInspectableVtbl
	Get_ExtendedError uintptr
}

type ICreateMediaSourceAudioInputNodeResult2 struct {
	win32.IInspectable
}

func (this *ICreateMediaSourceAudioInputNodeResult2) Vtbl() *ICreateMediaSourceAudioInputNodeResult2Vtbl {
	return (*ICreateMediaSourceAudioInputNodeResult2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICreateMediaSourceAudioInputNodeResult2) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 0E4D3FAA-36B8-4C91-B9DA-11F44A8A6610
var IID_IEchoEffectDefinition = syscall.GUID{0x0E4D3FAA, 0x36B8, 0x4C91,
	[8]byte{0xB9, 0xDA, 0x11, 0xF4, 0x4A, 0x8A, 0x66, 0x10}}

type IEchoEffectDefinitionInterface interface {
	win32.IInspectableInterface
	Put_WetDryMix(value float64)
	Get_WetDryMix() float64
	Put_Feedback(value float64)
	Get_Feedback() float64
	Put_Delay(value float64)
	Get_Delay() float64
}

type IEchoEffectDefinitionVtbl struct {
	win32.IInspectableVtbl
	Put_WetDryMix uintptr
	Get_WetDryMix uintptr
	Put_Feedback  uintptr
	Get_Feedback  uintptr
	Put_Delay     uintptr
	Get_Delay     uintptr
}

type IEchoEffectDefinition struct {
	win32.IInspectable
}

func (this *IEchoEffectDefinition) Vtbl() *IEchoEffectDefinitionVtbl {
	return (*IEchoEffectDefinitionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IEchoEffectDefinition) Put_WetDryMix(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_WetDryMix, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IEchoEffectDefinition) Get_WetDryMix() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WetDryMix, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IEchoEffectDefinition) Put_Feedback(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Feedback, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IEchoEffectDefinition) Get_Feedback() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Feedback, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IEchoEffectDefinition) Put_Delay(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Delay, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IEchoEffectDefinition) Get_Delay() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Delay, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 0D4E2257-AAF2-4E86-A54C-FB79DB8F6C12
var IID_IEchoEffectDefinitionFactory = syscall.GUID{0x0D4E2257, 0xAAF2, 0x4E86,
	[8]byte{0xA5, 0x4C, 0xFB, 0x79, 0xDB, 0x8F, 0x6C, 0x12}}

type IEchoEffectDefinitionFactoryInterface interface {
	win32.IInspectableInterface
	Create(audioGraph *IAudioGraph) *IEchoEffectDefinition
}

type IEchoEffectDefinitionFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IEchoEffectDefinitionFactory struct {
	win32.IInspectable
}

func (this *IEchoEffectDefinitionFactory) Vtbl() *IEchoEffectDefinitionFactoryVtbl {
	return (*IEchoEffectDefinitionFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IEchoEffectDefinitionFactory) Create(audioGraph *IAudioGraph) *IEchoEffectDefinition {
	var _result *IEchoEffectDefinition
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(audioGraph)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C00A5A6A-262D-4B85-9BB7-43280B62ED0C
var IID_IEqualizerBand = syscall.GUID{0xC00A5A6A, 0x262D, 0x4B85,
	[8]byte{0x9B, 0xB7, 0x43, 0x28, 0x0B, 0x62, 0xED, 0x0C}}

type IEqualizerBandInterface interface {
	win32.IInspectableInterface
	Get_Bandwidth() float64
	Put_Bandwidth(value float64)
	Get_FrequencyCenter() float64
	Put_FrequencyCenter(value float64)
	Get_Gain() float64
	Put_Gain(value float64)
}

type IEqualizerBandVtbl struct {
	win32.IInspectableVtbl
	Get_Bandwidth       uintptr
	Put_Bandwidth       uintptr
	Get_FrequencyCenter uintptr
	Put_FrequencyCenter uintptr
	Get_Gain            uintptr
	Put_Gain            uintptr
}

type IEqualizerBand struct {
	win32.IInspectable
}

func (this *IEqualizerBand) Vtbl() *IEqualizerBandVtbl {
	return (*IEqualizerBandVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IEqualizerBand) Get_Bandwidth() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bandwidth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IEqualizerBand) Put_Bandwidth(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Bandwidth, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IEqualizerBand) Get_FrequencyCenter() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrequencyCenter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IEqualizerBand) Put_FrequencyCenter(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FrequencyCenter, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IEqualizerBand) Get_Gain() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Gain, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IEqualizerBand) Put_Gain(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Gain, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 023F6F1F-83FE-449A-A822-C696442D16B0
var IID_IEqualizerEffectDefinition = syscall.GUID{0x023F6F1F, 0x83FE, 0x449A,
	[8]byte{0xA8, 0x22, 0xC6, 0x96, 0x44, 0x2D, 0x16, 0xB0}}

type IEqualizerEffectDefinitionInterface interface {
	win32.IInspectableInterface
	Get_Bands() *IVectorView[*IEqualizerBand]
}

type IEqualizerEffectDefinitionVtbl struct {
	win32.IInspectableVtbl
	Get_Bands uintptr
}

type IEqualizerEffectDefinition struct {
	win32.IInspectable
}

func (this *IEqualizerEffectDefinition) Vtbl() *IEqualizerEffectDefinitionVtbl {
	return (*IEqualizerEffectDefinitionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IEqualizerEffectDefinition) Get_Bands() *IVectorView[*IEqualizerBand] {
	var _result *IVectorView[*IEqualizerBand]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bands, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D2876FC4-D410-4EB5-9E69-C9AA1277EAF0
var IID_IEqualizerEffectDefinitionFactory = syscall.GUID{0xD2876FC4, 0xD410, 0x4EB5,
	[8]byte{0x9E, 0x69, 0xC9, 0xAA, 0x12, 0x77, 0xEA, 0xF0}}

type IEqualizerEffectDefinitionFactoryInterface interface {
	win32.IInspectableInterface
	Create(audioGraph *IAudioGraph) *IEqualizerEffectDefinition
}

type IEqualizerEffectDefinitionFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IEqualizerEffectDefinitionFactory struct {
	win32.IInspectable
}

func (this *IEqualizerEffectDefinitionFactory) Vtbl() *IEqualizerEffectDefinitionFactoryVtbl {
	return (*IEqualizerEffectDefinitionFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IEqualizerEffectDefinitionFactory) Create(audioGraph *IAudioGraph) *IEqualizerEffectDefinition {
	var _result *IEqualizerEffectDefinition
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(audioGraph)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3D9BD498-A306-4F06-BD9F-E9EFC8226304
var IID_IFrameInputNodeQuantumStartedEventArgs = syscall.GUID{0x3D9BD498, 0xA306, 0x4F06,
	[8]byte{0xBD, 0x9F, 0xE9, 0xEF, 0xC8, 0x22, 0x63, 0x04}}

type IFrameInputNodeQuantumStartedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_RequiredSamples() int32
}

type IFrameInputNodeQuantumStartedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_RequiredSamples uintptr
}

type IFrameInputNodeQuantumStartedEventArgs struct {
	win32.IInspectable
}

func (this *IFrameInputNodeQuantumStartedEventArgs) Vtbl() *IFrameInputNodeQuantumStartedEventArgsVtbl {
	return (*IFrameInputNodeQuantumStartedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFrameInputNodeQuantumStartedEventArgs) Get_RequiredSamples() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequiredSamples, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 6B755D19-2603-47BA-BDEB-39055E3486DC
var IID_ILimiterEffectDefinition = syscall.GUID{0x6B755D19, 0x2603, 0x47BA,
	[8]byte{0xBD, 0xEB, 0x39, 0x05, 0x5E, 0x34, 0x86, 0xDC}}

type ILimiterEffectDefinitionInterface interface {
	win32.IInspectableInterface
	Put_Release(value uint32)
	Get_Release() uint32
	Put_Loudness(value uint32)
	Get_Loudness() uint32
}

type ILimiterEffectDefinitionVtbl struct {
	win32.IInspectableVtbl
	Put_Release  uintptr
	Get_Release  uintptr
	Put_Loudness uintptr
	Get_Loudness uintptr
}

type ILimiterEffectDefinition struct {
	win32.IInspectable
}

func (this *ILimiterEffectDefinition) Vtbl() *ILimiterEffectDefinitionVtbl {
	return (*ILimiterEffectDefinitionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILimiterEffectDefinition) Put_Release(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Release, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ILimiterEffectDefinition) Get_Release() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Release, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILimiterEffectDefinition) Put_Loudness(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Loudness, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ILimiterEffectDefinition) Get_Loudness() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Loudness, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// ECBAE6F1-61FF-45EF-B8F5-48659A57C72D
var IID_ILimiterEffectDefinitionFactory = syscall.GUID{0xECBAE6F1, 0x61FF, 0x45EF,
	[8]byte{0xB8, 0xF5, 0x48, 0x65, 0x9A, 0x57, 0xC7, 0x2D}}

type ILimiterEffectDefinitionFactoryInterface interface {
	win32.IInspectableInterface
	Create(audioGraph *IAudioGraph) *ILimiterEffectDefinition
}

type ILimiterEffectDefinitionFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ILimiterEffectDefinitionFactory struct {
	win32.IInspectable
}

func (this *ILimiterEffectDefinitionFactory) Vtbl() *ILimiterEffectDefinitionFactoryVtbl {
	return (*ILimiterEffectDefinitionFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILimiterEffectDefinitionFactory) Create(audioGraph *IAudioGraph) *ILimiterEffectDefinition {
	var _result *ILimiterEffectDefinition
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(audioGraph)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 99D8983B-A88A-4041-8E4F-DDBAC0C91FD3
var IID_IMediaSourceAudioInputNode = syscall.GUID{0x99D8983B, 0xA88A, 0x4041,
	[8]byte{0x8E, 0x4F, 0xDD, 0xBA, 0xC0, 0xC9, 0x1F, 0xD3}}

type IMediaSourceAudioInputNodeInterface interface {
	win32.IInspectableInterface
	Put_PlaybackSpeedFactor(value float64)
	Get_PlaybackSpeedFactor() float64
	Get_Position() TimeSpan
	Seek(position TimeSpan)
	Get_StartTime() *IReference[TimeSpan]
	Put_StartTime(value *IReference[TimeSpan])
	Get_EndTime() *IReference[TimeSpan]
	Put_EndTime(value *IReference[TimeSpan])
	Get_LoopCount() *IReference[int32]
	Put_LoopCount(value *IReference[int32])
	Get_Duration() TimeSpan
	Get_MediaSource() *IMediaSource2
	Add_MediaSourceCompleted(handler TypedEventHandler[*IMediaSourceAudioInputNode, interface{}]) EventRegistrationToken
	Remove_MediaSourceCompleted(token EventRegistrationToken)
}

type IMediaSourceAudioInputNodeVtbl struct {
	win32.IInspectableVtbl
	Put_PlaybackSpeedFactor     uintptr
	Get_PlaybackSpeedFactor     uintptr
	Get_Position                uintptr
	Seek                        uintptr
	Get_StartTime               uintptr
	Put_StartTime               uintptr
	Get_EndTime                 uintptr
	Put_EndTime                 uintptr
	Get_LoopCount               uintptr
	Put_LoopCount               uintptr
	Get_Duration                uintptr
	Get_MediaSource             uintptr
	Add_MediaSourceCompleted    uintptr
	Remove_MediaSourceCompleted uintptr
}

type IMediaSourceAudioInputNode struct {
	win32.IInspectable
}

func (this *IMediaSourceAudioInputNode) Vtbl() *IMediaSourceAudioInputNodeVtbl {
	return (*IMediaSourceAudioInputNodeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaSourceAudioInputNode) Put_PlaybackSpeedFactor(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PlaybackSpeedFactor, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMediaSourceAudioInputNode) Get_PlaybackSpeedFactor() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlaybackSpeedFactor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaSourceAudioInputNode) Get_Position() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaSourceAudioInputNode) Seek(position TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Seek, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&position)))
	_ = _hr
}

func (this *IMediaSourceAudioInputNode) Get_StartTime() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StartTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaSourceAudioInputNode) Put_StartTime(value *IReference[TimeSpan]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StartTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaSourceAudioInputNode) Get_EndTime() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EndTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaSourceAudioInputNode) Put_EndTime(value *IReference[TimeSpan]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_EndTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaSourceAudioInputNode) Get_LoopCount() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LoopCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaSourceAudioInputNode) Put_LoopCount(value *IReference[int32]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LoopCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaSourceAudioInputNode) Get_Duration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Duration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaSourceAudioInputNode) Get_MediaSource() *IMediaSource2 {
	var _result *IMediaSource2
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaSourceAudioInputNode) Add_MediaSourceCompleted(handler TypedEventHandler[*IMediaSourceAudioInputNode, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_MediaSourceCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaSourceAudioInputNode) Remove_MediaSourceCompleted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_MediaSourceCompleted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 4606AA89-F563-4D0A-8F6E-F0CDDFF35D84
var IID_IReverbEffectDefinition = syscall.GUID{0x4606AA89, 0xF563, 0x4D0A,
	[8]byte{0x8F, 0x6E, 0xF0, 0xCD, 0xDF, 0xF3, 0x5D, 0x84}}

type IReverbEffectDefinitionInterface interface {
	win32.IInspectableInterface
	Put_WetDryMix(value float64)
	Get_WetDryMix() float64
	Put_ReflectionsDelay(value uint32)
	Get_ReflectionsDelay() uint32
	Put_ReverbDelay(value byte)
	Get_ReverbDelay() byte
	Put_RearDelay(value byte)
	Get_RearDelay() byte
	Put_PositionLeft(value byte)
	Get_PositionLeft() byte
	Put_PositionRight(value byte)
	Get_PositionRight() byte
	Put_PositionMatrixLeft(value byte)
	Get_PositionMatrixLeft() byte
	Put_PositionMatrixRight(value byte)
	Get_PositionMatrixRight() byte
	Put_EarlyDiffusion(value byte)
	Get_EarlyDiffusion() byte
	Put_LateDiffusion(value byte)
	Get_LateDiffusion() byte
	Put_LowEQGain(value byte)
	Get_LowEQGain() byte
	Put_LowEQCutoff(value byte)
	Get_LowEQCutoff() byte
	Put_HighEQGain(value byte)
	Get_HighEQGain() byte
	Put_HighEQCutoff(value byte)
	Get_HighEQCutoff() byte
	Put_RoomFilterFreq(value float64)
	Get_RoomFilterFreq() float64
	Put_RoomFilterMain(value float64)
	Get_RoomFilterMain() float64
	Put_RoomFilterHF(value float64)
	Get_RoomFilterHF() float64
	Put_ReflectionsGain(value float64)
	Get_ReflectionsGain() float64
	Put_ReverbGain(value float64)
	Get_ReverbGain() float64
	Put_DecayTime(value float64)
	Get_DecayTime() float64
	Put_Density(value float64)
	Get_Density() float64
	Put_RoomSize(value float64)
	Get_RoomSize() float64
	Put_DisableLateField(value bool)
	Get_DisableLateField() bool
}

type IReverbEffectDefinitionVtbl struct {
	win32.IInspectableVtbl
	Put_WetDryMix           uintptr
	Get_WetDryMix           uintptr
	Put_ReflectionsDelay    uintptr
	Get_ReflectionsDelay    uintptr
	Put_ReverbDelay         uintptr
	Get_ReverbDelay         uintptr
	Put_RearDelay           uintptr
	Get_RearDelay           uintptr
	Put_PositionLeft        uintptr
	Get_PositionLeft        uintptr
	Put_PositionRight       uintptr
	Get_PositionRight       uintptr
	Put_PositionMatrixLeft  uintptr
	Get_PositionMatrixLeft  uintptr
	Put_PositionMatrixRight uintptr
	Get_PositionMatrixRight uintptr
	Put_EarlyDiffusion      uintptr
	Get_EarlyDiffusion      uintptr
	Put_LateDiffusion       uintptr
	Get_LateDiffusion       uintptr
	Put_LowEQGain           uintptr
	Get_LowEQGain           uintptr
	Put_LowEQCutoff         uintptr
	Get_LowEQCutoff         uintptr
	Put_HighEQGain          uintptr
	Get_HighEQGain          uintptr
	Put_HighEQCutoff        uintptr
	Get_HighEQCutoff        uintptr
	Put_RoomFilterFreq      uintptr
	Get_RoomFilterFreq      uintptr
	Put_RoomFilterMain      uintptr
	Get_RoomFilterMain      uintptr
	Put_RoomFilterHF        uintptr
	Get_RoomFilterHF        uintptr
	Put_ReflectionsGain     uintptr
	Get_ReflectionsGain     uintptr
	Put_ReverbGain          uintptr
	Get_ReverbGain          uintptr
	Put_DecayTime           uintptr
	Get_DecayTime           uintptr
	Put_Density             uintptr
	Get_Density             uintptr
	Put_RoomSize            uintptr
	Get_RoomSize            uintptr
	Put_DisableLateField    uintptr
	Get_DisableLateField    uintptr
}

type IReverbEffectDefinition struct {
	win32.IInspectable
}

func (this *IReverbEffectDefinition) Vtbl() *IReverbEffectDefinitionVtbl {
	return (*IReverbEffectDefinitionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IReverbEffectDefinition) Put_WetDryMix(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_WetDryMix, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IReverbEffectDefinition) Get_WetDryMix() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WetDryMix, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReverbEffectDefinition) Put_ReflectionsDelay(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReflectionsDelay, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IReverbEffectDefinition) Get_ReflectionsDelay() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReflectionsDelay, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReverbEffectDefinition) Put_ReverbDelay(value byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReverbDelay, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IReverbEffectDefinition) Get_ReverbDelay() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReverbDelay, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReverbEffectDefinition) Put_RearDelay(value byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RearDelay, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IReverbEffectDefinition) Get_RearDelay() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RearDelay, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReverbEffectDefinition) Put_PositionLeft(value byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PositionLeft, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IReverbEffectDefinition) Get_PositionLeft() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PositionLeft, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReverbEffectDefinition) Put_PositionRight(value byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PositionRight, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IReverbEffectDefinition) Get_PositionRight() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PositionRight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReverbEffectDefinition) Put_PositionMatrixLeft(value byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PositionMatrixLeft, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IReverbEffectDefinition) Get_PositionMatrixLeft() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PositionMatrixLeft, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReverbEffectDefinition) Put_PositionMatrixRight(value byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PositionMatrixRight, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IReverbEffectDefinition) Get_PositionMatrixRight() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PositionMatrixRight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReverbEffectDefinition) Put_EarlyDiffusion(value byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_EarlyDiffusion, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IReverbEffectDefinition) Get_EarlyDiffusion() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EarlyDiffusion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReverbEffectDefinition) Put_LateDiffusion(value byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LateDiffusion, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IReverbEffectDefinition) Get_LateDiffusion() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LateDiffusion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReverbEffectDefinition) Put_LowEQGain(value byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LowEQGain, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IReverbEffectDefinition) Get_LowEQGain() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LowEQGain, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReverbEffectDefinition) Put_LowEQCutoff(value byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LowEQCutoff, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IReverbEffectDefinition) Get_LowEQCutoff() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LowEQCutoff, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReverbEffectDefinition) Put_HighEQGain(value byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_HighEQGain, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IReverbEffectDefinition) Get_HighEQGain() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HighEQGain, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReverbEffectDefinition) Put_HighEQCutoff(value byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_HighEQCutoff, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IReverbEffectDefinition) Get_HighEQCutoff() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HighEQCutoff, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReverbEffectDefinition) Put_RoomFilterFreq(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RoomFilterFreq, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IReverbEffectDefinition) Get_RoomFilterFreq() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RoomFilterFreq, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReverbEffectDefinition) Put_RoomFilterMain(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RoomFilterMain, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IReverbEffectDefinition) Get_RoomFilterMain() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RoomFilterMain, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReverbEffectDefinition) Put_RoomFilterHF(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RoomFilterHF, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IReverbEffectDefinition) Get_RoomFilterHF() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RoomFilterHF, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReverbEffectDefinition) Put_ReflectionsGain(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReflectionsGain, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IReverbEffectDefinition) Get_ReflectionsGain() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReflectionsGain, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReverbEffectDefinition) Put_ReverbGain(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReverbGain, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IReverbEffectDefinition) Get_ReverbGain() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReverbGain, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReverbEffectDefinition) Put_DecayTime(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DecayTime, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IReverbEffectDefinition) Get_DecayTime() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DecayTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReverbEffectDefinition) Put_Density(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Density, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IReverbEffectDefinition) Get_Density() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Density, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReverbEffectDefinition) Put_RoomSize(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RoomSize, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IReverbEffectDefinition) Get_RoomSize() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RoomSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IReverbEffectDefinition) Put_DisableLateField(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DisableLateField, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IReverbEffectDefinition) Get_DisableLateField() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisableLateField, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A7D5CBFE-100B-4FF0-9DA6-DC4E05A759F0
var IID_IReverbEffectDefinitionFactory = syscall.GUID{0xA7D5CBFE, 0x100B, 0x4FF0,
	[8]byte{0x9D, 0xA6, 0xDC, 0x4E, 0x05, 0xA7, 0x59, 0xF0}}

type IReverbEffectDefinitionFactoryInterface interface {
	win32.IInspectableInterface
	Create(audioGraph *IAudioGraph) *IReverbEffectDefinition
}

type IReverbEffectDefinitionFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IReverbEffectDefinitionFactory struct {
	win32.IInspectable
}

func (this *IReverbEffectDefinitionFactory) Vtbl() *IReverbEffectDefinitionFactoryVtbl {
	return (*IReverbEffectDefinitionFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IReverbEffectDefinitionFactory) Create(audioGraph *IAudioGraph) *IReverbEffectDefinition {
	var _result *IReverbEffectDefinition
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(audioGraph)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1C2AA511-1400-5E70-9EA9-AE151241E8EA
var IID_ISetDefaultSpatialAudioFormatResult = syscall.GUID{0x1C2AA511, 0x1400, 0x5E70,
	[8]byte{0x9E, 0xA9, 0xAE, 0x15, 0x12, 0x41, 0xE8, 0xEA}}

type ISetDefaultSpatialAudioFormatResultInterface interface {
	win32.IInspectableInterface
	Get_Status() SetDefaultSpatialAudioFormatStatus
}

type ISetDefaultSpatialAudioFormatResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
}

type ISetDefaultSpatialAudioFormatResult struct {
	win32.IInspectable
}

func (this *ISetDefaultSpatialAudioFormatResult) Vtbl() *ISetDefaultSpatialAudioFormatResultVtbl {
	return (*ISetDefaultSpatialAudioFormatResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISetDefaultSpatialAudioFormatResult) Get_Status() SetDefaultSpatialAudioFormatStatus {
	var _result SetDefaultSpatialAudioFormatStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// EE830034-61CF-5749-9DA4-10F0FE028199
var IID_ISpatialAudioDeviceConfiguration = syscall.GUID{0xEE830034, 0x61CF, 0x5749,
	[8]byte{0x9D, 0xA4, 0x10, 0xF0, 0xFE, 0x02, 0x81, 0x99}}

type ISpatialAudioDeviceConfigurationInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_IsSpatialAudioSupported() bool
	IsSpatialAudioFormatSupported(subtype string) bool
	Get_ActiveSpatialAudioFormat() string
	Get_DefaultSpatialAudioFormat() string
	SetDefaultSpatialAudioFormatAsync(subtype string) *IAsyncOperation[*ISetDefaultSpatialAudioFormatResult]
	Add_ConfigurationChanged(handler TypedEventHandler[*ISpatialAudioDeviceConfiguration, interface{}]) EventRegistrationToken
	Remove_ConfigurationChanged(token EventRegistrationToken)
}

type ISpatialAudioDeviceConfigurationVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId                      uintptr
	Get_IsSpatialAudioSupported       uintptr
	IsSpatialAudioFormatSupported     uintptr
	Get_ActiveSpatialAudioFormat      uintptr
	Get_DefaultSpatialAudioFormat     uintptr
	SetDefaultSpatialAudioFormatAsync uintptr
	Add_ConfigurationChanged          uintptr
	Remove_ConfigurationChanged       uintptr
}

type ISpatialAudioDeviceConfiguration struct {
	win32.IInspectable
}

func (this *ISpatialAudioDeviceConfiguration) Vtbl() *ISpatialAudioDeviceConfigurationVtbl {
	return (*ISpatialAudioDeviceConfigurationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialAudioDeviceConfiguration) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISpatialAudioDeviceConfiguration) Get_IsSpatialAudioSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSpatialAudioSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialAudioDeviceConfiguration) IsSpatialAudioFormatSupported(subtype string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsSpatialAudioFormatSupported, uintptr(unsafe.Pointer(this)), NewHStr(subtype).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialAudioDeviceConfiguration) Get_ActiveSpatialAudioFormat() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ActiveSpatialAudioFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISpatialAudioDeviceConfiguration) Get_DefaultSpatialAudioFormat() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DefaultSpatialAudioFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISpatialAudioDeviceConfiguration) SetDefaultSpatialAudioFormatAsync(subtype string) *IAsyncOperation[*ISetDefaultSpatialAudioFormatResult] {
	var _result *IAsyncOperation[*ISetDefaultSpatialAudioFormatResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetDefaultSpatialAudioFormatAsync, uintptr(unsafe.Pointer(this)), NewHStr(subtype).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialAudioDeviceConfiguration) Add_ConfigurationChanged(handler TypedEventHandler[*ISpatialAudioDeviceConfiguration, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ConfigurationChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialAudioDeviceConfiguration) Remove_ConfigurationChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ConfigurationChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 3EC37F7B-936D-4E04-9728-2827D9F758C4
var IID_ISpatialAudioDeviceConfigurationStatics = syscall.GUID{0x3EC37F7B, 0x936D, 0x4E04,
	[8]byte{0x97, 0x28, 0x28, 0x27, 0xD9, 0xF7, 0x58, 0xC4}}

type ISpatialAudioDeviceConfigurationStaticsInterface interface {
	win32.IInspectableInterface
	GetForDeviceId(deviceId string) *ISpatialAudioDeviceConfiguration
}

type ISpatialAudioDeviceConfigurationStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForDeviceId uintptr
}

type ISpatialAudioDeviceConfigurationStatics struct {
	win32.IInspectable
}

func (this *ISpatialAudioDeviceConfigurationStatics) Vtbl() *ISpatialAudioDeviceConfigurationStaticsVtbl {
	return (*ISpatialAudioDeviceConfigurationStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialAudioDeviceConfigurationStatics) GetForDeviceId(deviceId string) *ISpatialAudioDeviceConfiguration {
	var _result *ISpatialAudioDeviceConfiguration
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForDeviceId, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 32DF09A8-50F0-5395-9923-7D44CA71ED6D
var IID_ISpatialAudioFormatConfiguration = syscall.GUID{0x32DF09A8, 0x50F0, 0x5395,
	[8]byte{0x99, 0x23, 0x7D, 0x44, 0xCA, 0x71, 0xED, 0x6D}}

type ISpatialAudioFormatConfigurationInterface interface {
	win32.IInspectableInterface
	ReportLicenseChangedAsync(subtype string) *IAsyncAction
	ReportConfigurationChangedAsync(subtype string) *IAsyncAction
	Get_MixedRealityExclusiveModePolicy() MixedRealitySpatialAudioFormatPolicy
	Put_MixedRealityExclusiveModePolicy(value MixedRealitySpatialAudioFormatPolicy)
}

type ISpatialAudioFormatConfigurationVtbl struct {
	win32.IInspectableVtbl
	ReportLicenseChangedAsync           uintptr
	ReportConfigurationChangedAsync     uintptr
	Get_MixedRealityExclusiveModePolicy uintptr
	Put_MixedRealityExclusiveModePolicy uintptr
}

type ISpatialAudioFormatConfiguration struct {
	win32.IInspectable
}

func (this *ISpatialAudioFormatConfiguration) Vtbl() *ISpatialAudioFormatConfigurationVtbl {
	return (*ISpatialAudioFormatConfigurationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialAudioFormatConfiguration) ReportLicenseChangedAsync(subtype string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReportLicenseChangedAsync, uintptr(unsafe.Pointer(this)), NewHStr(subtype).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialAudioFormatConfiguration) ReportConfigurationChangedAsync(subtype string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReportConfigurationChangedAsync, uintptr(unsafe.Pointer(this)), NewHStr(subtype).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialAudioFormatConfiguration) Get_MixedRealityExclusiveModePolicy() MixedRealitySpatialAudioFormatPolicy {
	var _result MixedRealitySpatialAudioFormatPolicy
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MixedRealityExclusiveModePolicy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialAudioFormatConfiguration) Put_MixedRealityExclusiveModePolicy(value MixedRealitySpatialAudioFormatPolicy) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MixedRealityExclusiveModePolicy, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 2B5FEF71-67C9-4E5F-A35B-41680711F8C7
var IID_ISpatialAudioFormatConfigurationStatics = syscall.GUID{0x2B5FEF71, 0x67C9, 0x4E5F,
	[8]byte{0xA3, 0x5B, 0x41, 0x68, 0x07, 0x11, 0xF8, 0xC7}}

type ISpatialAudioFormatConfigurationStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *ISpatialAudioFormatConfiguration
}

type ISpatialAudioFormatConfigurationStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
}

type ISpatialAudioFormatConfigurationStatics struct {
	win32.IInspectable
}

func (this *ISpatialAudioFormatConfigurationStatics) Vtbl() *ISpatialAudioFormatConfigurationStaticsVtbl {
	return (*ISpatialAudioFormatConfigurationStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialAudioFormatConfigurationStatics) GetDefault() *ISpatialAudioFormatConfiguration {
	var _result *ISpatialAudioFormatConfiguration
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B3DE8A47-83EE-4266-A945-BEDF507AFEED
var IID_ISpatialAudioFormatSubtypeStatics = syscall.GUID{0xB3DE8A47, 0x83EE, 0x4266,
	[8]byte{0xA9, 0x45, 0xBE, 0xDF, 0x50, 0x7A, 0xFE, 0xED}}

type ISpatialAudioFormatSubtypeStaticsInterface interface {
	win32.IInspectableInterface
	Get_WindowsSonic() string
	Get_DolbyAtmosForHeadphones() string
	Get_DolbyAtmosForHomeTheater() string
	Get_DolbyAtmosForSpeakers() string
	Get_DTSHeadphoneX() string
	Get_DTSXUltra() string
}

type ISpatialAudioFormatSubtypeStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_WindowsSonic             uintptr
	Get_DolbyAtmosForHeadphones  uintptr
	Get_DolbyAtmosForHomeTheater uintptr
	Get_DolbyAtmosForSpeakers    uintptr
	Get_DTSHeadphoneX            uintptr
	Get_DTSXUltra                uintptr
}

type ISpatialAudioFormatSubtypeStatics struct {
	win32.IInspectable
}

func (this *ISpatialAudioFormatSubtypeStatics) Vtbl() *ISpatialAudioFormatSubtypeStaticsVtbl {
	return (*ISpatialAudioFormatSubtypeStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialAudioFormatSubtypeStatics) Get_WindowsSonic() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WindowsSonic, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISpatialAudioFormatSubtypeStatics) Get_DolbyAtmosForHeadphones() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DolbyAtmosForHeadphones, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISpatialAudioFormatSubtypeStatics) Get_DolbyAtmosForHomeTheater() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DolbyAtmosForHomeTheater, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISpatialAudioFormatSubtypeStatics) Get_DolbyAtmosForSpeakers() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DolbyAtmosForSpeakers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISpatialAudioFormatSubtypeStatics) Get_DTSHeadphoneX() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DTSHeadphoneX, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISpatialAudioFormatSubtypeStatics) Get_DTSXUltra() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DTSXUltra, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 4565E6CB-D95B-5621-B6AF-0E8849C57C80
var IID_ISpatialAudioFormatSubtypeStatics2 = syscall.GUID{0x4565E6CB, 0xD95B, 0x5621,
	[8]byte{0xB6, 0xAF, 0x0E, 0x88, 0x49, 0xC5, 0x7C, 0x80}}

type ISpatialAudioFormatSubtypeStatics2Interface interface {
	win32.IInspectableInterface
	Get_DTSXForHomeTheater() string
}

type ISpatialAudioFormatSubtypeStatics2Vtbl struct {
	win32.IInspectableVtbl
	Get_DTSXForHomeTheater uintptr
}

type ISpatialAudioFormatSubtypeStatics2 struct {
	win32.IInspectable
}

func (this *ISpatialAudioFormatSubtypeStatics2) Vtbl() *ISpatialAudioFormatSubtypeStatics2Vtbl {
	return (*ISpatialAudioFormatSubtypeStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialAudioFormatSubtypeStatics2) Get_DTSXForHomeTheater() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DTSXForHomeTheater, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// classes

type AudioDeviceInputNode struct {
	RtClass
	*IAudioDeviceInputNode
}

type AudioDeviceOutputNode struct {
	RtClass
	*IAudioDeviceOutputNode
}

type AudioFileInputNode struct {
	RtClass
	*IAudioFileInputNode
}

type AudioFileOutputNode struct {
	RtClass
	*IAudioFileOutputNode
}

type AudioFrameCompletedEventArgs struct {
	RtClass
	*IAudioFrameCompletedEventArgs
}

type AudioFrameInputNode struct {
	RtClass
	*IAudioFrameInputNode
}

type AudioFrameOutputNode struct {
	RtClass
	*IAudioFrameOutputNode
}

type AudioGraph struct {
	RtClass
	*IAudioGraph
}

func NewIAudioGraphStatics() *IAudioGraphStatics {
	var p *IAudioGraphStatics
	hs := NewHStr("Windows.Media.Audio.AudioGraph")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAudioGraphStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type AudioGraphBatchUpdater struct {
	RtClass
	*IClosable
}

type AudioGraphConnection struct {
	RtClass
	*IAudioGraphConnection
}

type AudioGraphSettings struct {
	RtClass
	*IAudioGraphSettings
}

func NewAudioGraphSettings_Create(audioRenderCategory unsafe.Pointer) *AudioGraphSettings {
	hs := NewHStr("Windows.Media.Audio.AudioGraphSettings")
	var pFac *IAudioGraphSettingsFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAudioGraphSettingsFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IAudioGraphSettings
	p = pFac.Create(audioRenderCategory)
	result := &AudioGraphSettings{
		RtClass:             RtClass{PInspect: &p.IInspectable},
		IAudioGraphSettings: p,
	}
	com.AddToScope(result)
	return result
}

type AudioGraphUnrecoverableErrorOccurredEventArgs struct {
	RtClass
	*IAudioGraphUnrecoverableErrorOccurredEventArgs
}

type AudioNodeEmitter struct {
	RtClass
	*IAudioNodeEmitter
}

func NewAudioNodeEmitter() *AudioNodeEmitter {
	hs := NewHStr("Windows.Media.Audio.AudioNodeEmitter")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &AudioNodeEmitter{
		RtClass:           RtClass{PInspect: p},
		IAudioNodeEmitter: (*IAudioNodeEmitter)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewAudioNodeEmitter_CreateAudioNodeEmitter(shape *IAudioNodeEmitterShape, decayModel *IAudioNodeEmitterDecayModel, settings AudioNodeEmitterSettings) *AudioNodeEmitter {
	hs := NewHStr("Windows.Media.Audio.AudioNodeEmitter")
	var pFac *IAudioNodeEmitterFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAudioNodeEmitterFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IAudioNodeEmitter
	p = pFac.CreateAudioNodeEmitter(shape, decayModel, settings)
	result := &AudioNodeEmitter{
		RtClass:           RtClass{PInspect: &p.IInspectable},
		IAudioNodeEmitter: p,
	}
	com.AddToScope(result)
	return result
}

type AudioNodeEmitterConeProperties struct {
	RtClass
	*IAudioNodeEmitterConeProperties
}

type AudioNodeEmitterDecayModel struct {
	RtClass
	*IAudioNodeEmitterDecayModel
}

func NewIAudioNodeEmitterDecayModelStatics() *IAudioNodeEmitterDecayModelStatics {
	var p *IAudioNodeEmitterDecayModelStatics
	hs := NewHStr("Windows.Media.Audio.AudioNodeEmitterDecayModel")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAudioNodeEmitterDecayModelStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type AudioNodeEmitterNaturalDecayModelProperties struct {
	RtClass
	*IAudioNodeEmitterNaturalDecayModelProperties
}

type AudioNodeEmitterShape struct {
	RtClass
	*IAudioNodeEmitterShape
}

func NewIAudioNodeEmitterShapeStatics() *IAudioNodeEmitterShapeStatics {
	var p *IAudioNodeEmitterShapeStatics
	hs := NewHStr("Windows.Media.Audio.AudioNodeEmitterShape")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAudioNodeEmitterShapeStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type AudioNodeListener struct {
	RtClass
	*IAudioNodeListener
}

func NewAudioNodeListener() *AudioNodeListener {
	hs := NewHStr("Windows.Media.Audio.AudioNodeListener")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &AudioNodeListener{
		RtClass:            RtClass{PInspect: p},
		IAudioNodeListener: (*IAudioNodeListener)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type AudioPlaybackConnection struct {
	RtClass
	*IAudioPlaybackConnection
}

func NewIAudioPlaybackConnectionStatics() *IAudioPlaybackConnectionStatics {
	var p *IAudioPlaybackConnectionStatics
	hs := NewHStr("Windows.Media.Audio.AudioPlaybackConnection")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAudioPlaybackConnectionStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type AudioPlaybackConnectionOpenResult struct {
	RtClass
	*IAudioPlaybackConnectionOpenResult
}

type AudioSubmixNode struct {
	RtClass
	*IAudioInputNode
}

type CreateAudioDeviceInputNodeResult struct {
	RtClass
	*ICreateAudioDeviceInputNodeResult
}

type CreateAudioDeviceOutputNodeResult struct {
	RtClass
	*ICreateAudioDeviceOutputNodeResult
}

type CreateAudioFileInputNodeResult struct {
	RtClass
	*ICreateAudioFileInputNodeResult
}

type CreateAudioFileOutputNodeResult struct {
	RtClass
	*ICreateAudioFileOutputNodeResult
}

type CreateAudioGraphResult struct {
	RtClass
	*ICreateAudioGraphResult
}

type CreateMediaSourceAudioInputNodeResult struct {
	RtClass
	*ICreateMediaSourceAudioInputNodeResult
}

type EchoEffectDefinition struct {
	RtClass
	*IEchoEffectDefinition
}

func NewEchoEffectDefinition_Create(audioGraph *IAudioGraph) *EchoEffectDefinition {
	hs := NewHStr("Windows.Media.Audio.EchoEffectDefinition")
	var pFac *IEchoEffectDefinitionFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IEchoEffectDefinitionFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IEchoEffectDefinition
	p = pFac.Create(audioGraph)
	result := &EchoEffectDefinition{
		RtClass:               RtClass{PInspect: &p.IInspectable},
		IEchoEffectDefinition: p,
	}
	com.AddToScope(result)
	return result
}

type EqualizerBand struct {
	RtClass
	*IEqualizerBand
}

type EqualizerEffectDefinition struct {
	RtClass
	*IEqualizerEffectDefinition
}

func NewEqualizerEffectDefinition_Create(audioGraph *IAudioGraph) *EqualizerEffectDefinition {
	hs := NewHStr("Windows.Media.Audio.EqualizerEffectDefinition")
	var pFac *IEqualizerEffectDefinitionFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IEqualizerEffectDefinitionFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IEqualizerEffectDefinition
	p = pFac.Create(audioGraph)
	result := &EqualizerEffectDefinition{
		RtClass:                    RtClass{PInspect: &p.IInspectable},
		IEqualizerEffectDefinition: p,
	}
	com.AddToScope(result)
	return result
}

type FrameInputNodeQuantumStartedEventArgs struct {
	RtClass
	*IFrameInputNodeQuantumStartedEventArgs
}

type LimiterEffectDefinition struct {
	RtClass
	*ILimiterEffectDefinition
}

func NewLimiterEffectDefinition_Create(audioGraph *IAudioGraph) *LimiterEffectDefinition {
	hs := NewHStr("Windows.Media.Audio.LimiterEffectDefinition")
	var pFac *ILimiterEffectDefinitionFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ILimiterEffectDefinitionFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ILimiterEffectDefinition
	p = pFac.Create(audioGraph)
	result := &LimiterEffectDefinition{
		RtClass:                  RtClass{PInspect: &p.IInspectable},
		ILimiterEffectDefinition: p,
	}
	com.AddToScope(result)
	return result
}

type MediaSourceAudioInputNode struct {
	RtClass
	*IMediaSourceAudioInputNode
}

type ReverbEffectDefinition struct {
	RtClass
	*IReverbEffectDefinition
}

func NewReverbEffectDefinition_Create(audioGraph *IAudioGraph) *ReverbEffectDefinition {
	hs := NewHStr("Windows.Media.Audio.ReverbEffectDefinition")
	var pFac *IReverbEffectDefinitionFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IReverbEffectDefinitionFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IReverbEffectDefinition
	p = pFac.Create(audioGraph)
	result := &ReverbEffectDefinition{
		RtClass:                 RtClass{PInspect: &p.IInspectable},
		IReverbEffectDefinition: p,
	}
	com.AddToScope(result)
	return result
}

type SetDefaultSpatialAudioFormatResult struct {
	RtClass
	*ISetDefaultSpatialAudioFormatResult
}

type SpatialAudioDeviceConfiguration struct {
	RtClass
	*ISpatialAudioDeviceConfiguration
}

func NewISpatialAudioDeviceConfigurationStatics() *ISpatialAudioDeviceConfigurationStatics {
	var p *ISpatialAudioDeviceConfigurationStatics
	hs := NewHStr("Windows.Media.Audio.SpatialAudioDeviceConfiguration")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISpatialAudioDeviceConfigurationStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type SpatialAudioFormatConfiguration struct {
	RtClass
	*ISpatialAudioFormatConfiguration
}

func NewISpatialAudioFormatConfigurationStatics() *ISpatialAudioFormatConfigurationStatics {
	var p *ISpatialAudioFormatConfigurationStatics
	hs := NewHStr("Windows.Media.Audio.SpatialAudioFormatConfiguration")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISpatialAudioFormatConfigurationStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type SpatialAudioFormatSubtype struct {
	RtClass
}

func NewISpatialAudioFormatSubtypeStatics2() *ISpatialAudioFormatSubtypeStatics2 {
	var p *ISpatialAudioFormatSubtypeStatics2
	hs := NewHStr("Windows.Media.Audio.SpatialAudioFormatSubtype")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISpatialAudioFormatSubtypeStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewISpatialAudioFormatSubtypeStatics() *ISpatialAudioFormatSubtypeStatics {
	var p *ISpatialAudioFormatSubtypeStatics
	hs := NewHStr("Windows.Media.Audio.SpatialAudioFormatSubtype")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISpatialAudioFormatSubtypeStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
