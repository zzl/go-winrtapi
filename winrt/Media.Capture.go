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
type AppBroadcastCameraCaptureState int32

const (
	AppBroadcastCameraCaptureState_Stopped AppBroadcastCameraCaptureState = 0
	AppBroadcastCameraCaptureState_Started AppBroadcastCameraCaptureState = 1
	AppBroadcastCameraCaptureState_Failed  AppBroadcastCameraCaptureState = 2
)

// enum
type AppBroadcastCameraOverlayLocation int32

const (
	AppBroadcastCameraOverlayLocation_TopLeft      AppBroadcastCameraOverlayLocation = 0
	AppBroadcastCameraOverlayLocation_TopCenter    AppBroadcastCameraOverlayLocation = 1
	AppBroadcastCameraOverlayLocation_TopRight     AppBroadcastCameraOverlayLocation = 2
	AppBroadcastCameraOverlayLocation_MiddleLeft   AppBroadcastCameraOverlayLocation = 3
	AppBroadcastCameraOverlayLocation_MiddleCenter AppBroadcastCameraOverlayLocation = 4
	AppBroadcastCameraOverlayLocation_MiddleRight  AppBroadcastCameraOverlayLocation = 5
	AppBroadcastCameraOverlayLocation_BottomLeft   AppBroadcastCameraOverlayLocation = 6
	AppBroadcastCameraOverlayLocation_BottomCenter AppBroadcastCameraOverlayLocation = 7
	AppBroadcastCameraOverlayLocation_BottomRight  AppBroadcastCameraOverlayLocation = 8
)

// enum
type AppBroadcastCameraOverlaySize int32

const (
	AppBroadcastCameraOverlaySize_Small  AppBroadcastCameraOverlaySize = 0
	AppBroadcastCameraOverlaySize_Medium AppBroadcastCameraOverlaySize = 1
	AppBroadcastCameraOverlaySize_Large  AppBroadcastCameraOverlaySize = 2
)

// enum
type AppBroadcastCaptureTargetType int32

const (
	AppBroadcastCaptureTargetType_AppView       AppBroadcastCaptureTargetType = 0
	AppBroadcastCaptureTargetType_EntireDisplay AppBroadcastCaptureTargetType = 1
)

// enum
type AppBroadcastExitBroadcastModeReason int32

const (
	AppBroadcastExitBroadcastModeReason_NormalExit             AppBroadcastExitBroadcastModeReason = 0
	AppBroadcastExitBroadcastModeReason_UserCanceled           AppBroadcastExitBroadcastModeReason = 1
	AppBroadcastExitBroadcastModeReason_AuthorizationFail      AppBroadcastExitBroadcastModeReason = 2
	AppBroadcastExitBroadcastModeReason_ForegroundAppActivated AppBroadcastExitBroadcastModeReason = 3
)

// enum
type AppBroadcastMicrophoneCaptureState int32

const (
	AppBroadcastMicrophoneCaptureState_Stopped AppBroadcastMicrophoneCaptureState = 0
	AppBroadcastMicrophoneCaptureState_Started AppBroadcastMicrophoneCaptureState = 1
	AppBroadcastMicrophoneCaptureState_Failed  AppBroadcastMicrophoneCaptureState = 2
)

// enum
type AppBroadcastPlugInState int32

const (
	AppBroadcastPlugInState_Unknown                 AppBroadcastPlugInState = 0
	AppBroadcastPlugInState_Initialized             AppBroadcastPlugInState = 1
	AppBroadcastPlugInState_MicrosoftSignInRequired AppBroadcastPlugInState = 2
	AppBroadcastPlugInState_OAuthSignInRequired     AppBroadcastPlugInState = 3
	AppBroadcastPlugInState_ProviderSignInRequired  AppBroadcastPlugInState = 4
	AppBroadcastPlugInState_InBandwidthTest         AppBroadcastPlugInState = 5
	AppBroadcastPlugInState_ReadyToBroadcast        AppBroadcastPlugInState = 6
)

// enum
type AppBroadcastPreviewState int32

const (
	AppBroadcastPreviewState_Started AppBroadcastPreviewState = 0
	AppBroadcastPreviewState_Stopped AppBroadcastPreviewState = 1
	AppBroadcastPreviewState_Failed  AppBroadcastPreviewState = 2
)

// enum
type AppBroadcastSignInResult int32

const (
	AppBroadcastSignInResult_Success              AppBroadcastSignInResult = 0
	AppBroadcastSignInResult_AuthenticationFailed AppBroadcastSignInResult = 1
	AppBroadcastSignInResult_Unauthorized         AppBroadcastSignInResult = 2
	AppBroadcastSignInResult_ServiceUnavailable   AppBroadcastSignInResult = 3
	AppBroadcastSignInResult_Unknown              AppBroadcastSignInResult = 4
)

// enum
type AppBroadcastSignInState int32

const (
	AppBroadcastSignInState_NotSignedIn               AppBroadcastSignInState = 0
	AppBroadcastSignInState_MicrosoftSignInInProgress AppBroadcastSignInState = 1
	AppBroadcastSignInState_MicrosoftSignInComplete   AppBroadcastSignInState = 2
	AppBroadcastSignInState_OAuthSignInInProgress     AppBroadcastSignInState = 3
	AppBroadcastSignInState_OAuthSignInComplete       AppBroadcastSignInState = 4
)

// enum
type AppBroadcastStreamState int32

const (
	AppBroadcastStreamState_Initializing AppBroadcastStreamState = 0
	AppBroadcastStreamState_StreamReady  AppBroadcastStreamState = 1
	AppBroadcastStreamState_Started      AppBroadcastStreamState = 2
	AppBroadcastStreamState_Paused       AppBroadcastStreamState = 3
	AppBroadcastStreamState_Terminated   AppBroadcastStreamState = 4
)

// enum
type AppBroadcastTerminationReason int32

const (
	AppBroadcastTerminationReason_NormalTermination          AppBroadcastTerminationReason = 0
	AppBroadcastTerminationReason_LostConnectionToService    AppBroadcastTerminationReason = 1
	AppBroadcastTerminationReason_NoNetworkConnectivity      AppBroadcastTerminationReason = 2
	AppBroadcastTerminationReason_ServiceAbort               AppBroadcastTerminationReason = 3
	AppBroadcastTerminationReason_ServiceError               AppBroadcastTerminationReason = 4
	AppBroadcastTerminationReason_ServiceUnavailable         AppBroadcastTerminationReason = 5
	AppBroadcastTerminationReason_InternalError              AppBroadcastTerminationReason = 6
	AppBroadcastTerminationReason_UnsupportedFormat          AppBroadcastTerminationReason = 7
	AppBroadcastTerminationReason_BackgroundTaskTerminated   AppBroadcastTerminationReason = 8
	AppBroadcastTerminationReason_BackgroundTaskUnresponsive AppBroadcastTerminationReason = 9
)

// enum
type AppBroadcastVideoEncodingBitrateMode int32

const (
	AppBroadcastVideoEncodingBitrateMode_Custom AppBroadcastVideoEncodingBitrateMode = 0
	AppBroadcastVideoEncodingBitrateMode_Auto   AppBroadcastVideoEncodingBitrateMode = 1
)

// enum
type AppBroadcastVideoEncodingResolutionMode int32

const (
	AppBroadcastVideoEncodingResolutionMode_Custom AppBroadcastVideoEncodingResolutionMode = 0
	AppBroadcastVideoEncodingResolutionMode_Auto   AppBroadcastVideoEncodingResolutionMode = 1
)

// enum
type AppCaptureHistoricalBufferLengthUnit int32

const (
	AppCaptureHistoricalBufferLengthUnit_Megabytes AppCaptureHistoricalBufferLengthUnit = 0
	AppCaptureHistoricalBufferLengthUnit_Seconds   AppCaptureHistoricalBufferLengthUnit = 1
)

// enum
type AppCaptureMetadataPriority int32

const (
	AppCaptureMetadataPriority_Informational AppCaptureMetadataPriority = 0
	AppCaptureMetadataPriority_Important     AppCaptureMetadataPriority = 1
)

// enum
type AppCaptureMicrophoneCaptureState int32

const (
	AppCaptureMicrophoneCaptureState_Stopped AppCaptureMicrophoneCaptureState = 0
	AppCaptureMicrophoneCaptureState_Started AppCaptureMicrophoneCaptureState = 1
	AppCaptureMicrophoneCaptureState_Failed  AppCaptureMicrophoneCaptureState = 2
)

// enum
type AppCaptureRecordingState int32

const (
	AppCaptureRecordingState_InProgress AppCaptureRecordingState = 0
	AppCaptureRecordingState_Completed  AppCaptureRecordingState = 1
	AppCaptureRecordingState_Failed     AppCaptureRecordingState = 2
)

// enum
type AppCaptureVideoEncodingBitrateMode int32

const (
	AppCaptureVideoEncodingBitrateMode_Custom   AppCaptureVideoEncodingBitrateMode = 0
	AppCaptureVideoEncodingBitrateMode_High     AppCaptureVideoEncodingBitrateMode = 1
	AppCaptureVideoEncodingBitrateMode_Standard AppCaptureVideoEncodingBitrateMode = 2
)

// enum
type AppCaptureVideoEncodingFrameRateMode int32

const (
	AppCaptureVideoEncodingFrameRateMode_Standard AppCaptureVideoEncodingFrameRateMode = 0
	AppCaptureVideoEncodingFrameRateMode_High     AppCaptureVideoEncodingFrameRateMode = 1
)

// enum
type AppCaptureVideoEncodingResolutionMode int32

const (
	AppCaptureVideoEncodingResolutionMode_Custom   AppCaptureVideoEncodingResolutionMode = 0
	AppCaptureVideoEncodingResolutionMode_High     AppCaptureVideoEncodingResolutionMode = 1
	AppCaptureVideoEncodingResolutionMode_Standard AppCaptureVideoEncodingResolutionMode = 2
)

// enum
type CameraCaptureUIMaxPhotoResolution int32

const (
	CameraCaptureUIMaxPhotoResolution_HighestAvailable CameraCaptureUIMaxPhotoResolution = 0
	CameraCaptureUIMaxPhotoResolution_VerySmallQvga    CameraCaptureUIMaxPhotoResolution = 1
	CameraCaptureUIMaxPhotoResolution_SmallVga         CameraCaptureUIMaxPhotoResolution = 2
	CameraCaptureUIMaxPhotoResolution_MediumXga        CameraCaptureUIMaxPhotoResolution = 3
	CameraCaptureUIMaxPhotoResolution_Large3M          CameraCaptureUIMaxPhotoResolution = 4
	CameraCaptureUIMaxPhotoResolution_VeryLarge5M      CameraCaptureUIMaxPhotoResolution = 5
)

// enum
type CameraCaptureUIMaxVideoResolution int32

const (
	CameraCaptureUIMaxVideoResolution_HighestAvailable   CameraCaptureUIMaxVideoResolution = 0
	CameraCaptureUIMaxVideoResolution_LowDefinition      CameraCaptureUIMaxVideoResolution = 1
	CameraCaptureUIMaxVideoResolution_StandardDefinition CameraCaptureUIMaxVideoResolution = 2
	CameraCaptureUIMaxVideoResolution_HighDefinition     CameraCaptureUIMaxVideoResolution = 3
)

// enum
type CameraCaptureUIMode int32

const (
	CameraCaptureUIMode_PhotoOrVideo CameraCaptureUIMode = 0
	CameraCaptureUIMode_Photo        CameraCaptureUIMode = 1
	CameraCaptureUIMode_Video        CameraCaptureUIMode = 2
)

// enum
type CameraCaptureUIPhotoFormat int32

const (
	CameraCaptureUIPhotoFormat_Jpeg   CameraCaptureUIPhotoFormat = 0
	CameraCaptureUIPhotoFormat_Png    CameraCaptureUIPhotoFormat = 1
	CameraCaptureUIPhotoFormat_JpegXR CameraCaptureUIPhotoFormat = 2
)

// enum
type CameraCaptureUIVideoFormat int32

const (
	CameraCaptureUIVideoFormat_Mp4 CameraCaptureUIVideoFormat = 0
	CameraCaptureUIVideoFormat_Wmv CameraCaptureUIVideoFormat = 1
)

// enum
type ForegroundActivationArgument int32

const (
	ForegroundActivationArgument_SignInRequired ForegroundActivationArgument = 0
	ForegroundActivationArgument_MoreSettings   ForegroundActivationArgument = 1
)

// enum
type GameBarCommand int32

const (
	GameBarCommand_OpenGameBar              GameBarCommand = 0
	GameBarCommand_RecordHistoricalBuffer   GameBarCommand = 1
	GameBarCommand_ToggleStartStopRecord    GameBarCommand = 2
	GameBarCommand_StartRecord              GameBarCommand = 3
	GameBarCommand_StopRecord               GameBarCommand = 4
	GameBarCommand_TakeScreenshot           GameBarCommand = 5
	GameBarCommand_StartBroadcast           GameBarCommand = 6
	GameBarCommand_StopBroadcast            GameBarCommand = 7
	GameBarCommand_PauseBroadcast           GameBarCommand = 8
	GameBarCommand_ResumeBroadcast          GameBarCommand = 9
	GameBarCommand_ToggleStartStopBroadcast GameBarCommand = 10
	GameBarCommand_ToggleMicrophoneCapture  GameBarCommand = 11
	GameBarCommand_ToggleCameraCapture      GameBarCommand = 12
	GameBarCommand_ToggleRecordingIndicator GameBarCommand = 13
)

// enum
type GameBarCommandOrigin int32

const (
	GameBarCommandOrigin_ShortcutKey GameBarCommandOrigin = 0
	GameBarCommandOrigin_Cortana     GameBarCommandOrigin = 1
	GameBarCommandOrigin_AppCommand  GameBarCommandOrigin = 2
)

// enum
type GameBarServicesDisplayMode int32

const (
	GameBarServicesDisplayMode_Windowed            GameBarServicesDisplayMode = 0
	GameBarServicesDisplayMode_FullScreenExclusive GameBarServicesDisplayMode = 1
)

// enum
type GameBarTargetCapturePolicy int32

const (
	GameBarTargetCapturePolicy_EnabledBySystem       GameBarTargetCapturePolicy = 0
	GameBarTargetCapturePolicy_EnabledByUser         GameBarTargetCapturePolicy = 1
	GameBarTargetCapturePolicy_NotEnabled            GameBarTargetCapturePolicy = 2
	GameBarTargetCapturePolicy_ProhibitedBySystem    GameBarTargetCapturePolicy = 3
	GameBarTargetCapturePolicy_ProhibitedByPublisher GameBarTargetCapturePolicy = 4
)

// enum
type KnownVideoProfile int32

const (
	KnownVideoProfile_VideoRecording        KnownVideoProfile = 0
	KnownVideoProfile_HighQualityPhoto      KnownVideoProfile = 1
	KnownVideoProfile_BalancedVideoAndPhoto KnownVideoProfile = 2
	KnownVideoProfile_VideoConferencing     KnownVideoProfile = 3
	KnownVideoProfile_PhotoSequence         KnownVideoProfile = 4
	KnownVideoProfile_HighFrameRate         KnownVideoProfile = 5
	KnownVideoProfile_VariablePhotoSequence KnownVideoProfile = 6
	KnownVideoProfile_HdrWithWcgVideo       KnownVideoProfile = 7
	KnownVideoProfile_HdrWithWcgPhoto       KnownVideoProfile = 8
	KnownVideoProfile_VideoHdr8             KnownVideoProfile = 9
	KnownVideoProfile_CompressedCamera      KnownVideoProfile = 10
)

// enum
type MediaCaptureDeviceExclusiveControlReleaseMode int32

const (
	MediaCaptureDeviceExclusiveControlReleaseMode_OnDispose           MediaCaptureDeviceExclusiveControlReleaseMode = 0
	MediaCaptureDeviceExclusiveControlReleaseMode_OnAllStreamsStopped MediaCaptureDeviceExclusiveControlReleaseMode = 1
)

// enum
type MediaCaptureDeviceExclusiveControlStatus int32

const (
	MediaCaptureDeviceExclusiveControlStatus_ExclusiveControlAvailable MediaCaptureDeviceExclusiveControlStatus = 0
	MediaCaptureDeviceExclusiveControlStatus_SharedReadOnlyAvailable   MediaCaptureDeviceExclusiveControlStatus = 1
)

// enum
type MediaCaptureMemoryPreference int32

const (
	MediaCaptureMemoryPreference_Auto MediaCaptureMemoryPreference = 0
	MediaCaptureMemoryPreference_Cpu  MediaCaptureMemoryPreference = 1
)

// enum
type MediaCaptureSharingMode int32

const (
	MediaCaptureSharingMode_ExclusiveControl MediaCaptureSharingMode = 0
	MediaCaptureSharingMode_SharedReadOnly   MediaCaptureSharingMode = 1
)

// enum
type MediaCaptureThermalStatus int32

const (
	MediaCaptureThermalStatus_Normal     MediaCaptureThermalStatus = 0
	MediaCaptureThermalStatus_Overheated MediaCaptureThermalStatus = 1
)

// enum
type MediaCategory int32

const (
	MediaCategory_Other          MediaCategory = 0
	MediaCategory_Communications MediaCategory = 1
	MediaCategory_Media          MediaCategory = 2
	MediaCategory_GameChat       MediaCategory = 3
	MediaCategory_Speech         MediaCategory = 4
	MediaCategory_FarFieldSpeech MediaCategory = 5
	MediaCategory_UniformSpeech  MediaCategory = 6
	MediaCategory_VoiceTyping    MediaCategory = 7
)

// enum
type MediaStreamType int32

const (
	MediaStreamType_VideoPreview MediaStreamType = 0
	MediaStreamType_VideoRecord  MediaStreamType = 1
	MediaStreamType_Audio        MediaStreamType = 2
	MediaStreamType_Photo        MediaStreamType = 3
	MediaStreamType_Metadata     MediaStreamType = 4
)

// enum
type PhotoCaptureSource int32

const (
	PhotoCaptureSource_Auto         PhotoCaptureSource = 0
	PhotoCaptureSource_VideoPreview PhotoCaptureSource = 1
	PhotoCaptureSource_Photo        PhotoCaptureSource = 2
)

// enum
type PowerlineFrequency int32

const (
	PowerlineFrequency_Disabled   PowerlineFrequency = 0
	PowerlineFrequency_FiftyHertz PowerlineFrequency = 1
	PowerlineFrequency_SixtyHertz PowerlineFrequency = 2
	PowerlineFrequency_Auto       PowerlineFrequency = 3
)

// enum
type StreamingCaptureMode int32

const (
	StreamingCaptureMode_AudioAndVideo StreamingCaptureMode = 0
	StreamingCaptureMode_Audio         StreamingCaptureMode = 1
	StreamingCaptureMode_Video         StreamingCaptureMode = 2
)

// enum
type VideoDeviceCharacteristic int32

const (
	VideoDeviceCharacteristic_AllStreamsIndependent         VideoDeviceCharacteristic = 0
	VideoDeviceCharacteristic_PreviewRecordStreamsIdentical VideoDeviceCharacteristic = 1
	VideoDeviceCharacteristic_PreviewPhotoStreamsIdentical  VideoDeviceCharacteristic = 2
	VideoDeviceCharacteristic_RecordPhotoStreamsIdentical   VideoDeviceCharacteristic = 3
	VideoDeviceCharacteristic_AllStreamsIdentical           VideoDeviceCharacteristic = 4
)

// enum
type VideoRotation int32

const (
	VideoRotation_None                VideoRotation = 0
	VideoRotation_Clockwise90Degrees  VideoRotation = 1
	VideoRotation_Clockwise180Degrees VideoRotation = 2
	VideoRotation_Clockwise270Degrees VideoRotation = 3
)

// structs

type AppBroadcastContract struct {
}

type AppCaptureContract struct {
}

type AppCaptureMetadataContract struct {
}

type CameraCaptureUIContract struct {
}

type GameBarContract struct {
}

type WhiteBalanceGain struct {
	R float64
	G float64
	B float64
}

// func types

//2014EFFB-5CD8-4F08-A314-0D360DA59F14
type MediaCaptureFailedEventHandler func(sender *IMediaCapture, errorEventArgs *IMediaCaptureFailedEventArgs) com.Error

//3FAE8F2E-4FE1-4FFD-AABA-E1F1337D4E53
type RecordLimitationExceededEventHandler func(sender *IMediaCapture) com.Error

// interfaces

// F072728B-B292-4491-9D41-99807A550BBF
var IID_IAdvancedCapturedPhoto = syscall.GUID{0xF072728B, 0xB292, 0x4491,
	[8]byte{0x9D, 0x41, 0x99, 0x80, 0x7A, 0x55, 0x0B, 0xBF}}

type IAdvancedCapturedPhotoInterface interface {
	win32.IInspectableInterface
	Get_Frame() *ICapturedFrame
	Get_Mode() AdvancedPhotoMode
	Get_Context() interface{}
}

type IAdvancedCapturedPhotoVtbl struct {
	win32.IInspectableVtbl
	Get_Frame   uintptr
	Get_Mode    uintptr
	Get_Context uintptr
}

type IAdvancedCapturedPhoto struct {
	win32.IInspectable
}

func (this *IAdvancedCapturedPhoto) Vtbl() *IAdvancedCapturedPhotoVtbl {
	return (*IAdvancedCapturedPhotoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdvancedCapturedPhoto) Get_Frame() *ICapturedFrame {
	var _result *ICapturedFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Frame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdvancedCapturedPhoto) Get_Mode() AdvancedPhotoMode {
	var _result AdvancedPhotoMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdvancedCapturedPhoto) Get_Context() interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Context, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 18CF6CD8-CFFE-42D8-8104-017BB318F4A1
var IID_IAdvancedCapturedPhoto2 = syscall.GUID{0x18CF6CD8, 0xCFFE, 0x42D8,
	[8]byte{0x81, 0x04, 0x01, 0x7B, 0xB3, 0x18, 0xF4, 0xA1}}

type IAdvancedCapturedPhoto2Interface interface {
	win32.IInspectableInterface
	Get_FrameBoundsRelativeToReferencePhoto() *IReference[Rect]
}

type IAdvancedCapturedPhoto2Vtbl struct {
	win32.IInspectableVtbl
	Get_FrameBoundsRelativeToReferencePhoto uintptr
}

type IAdvancedCapturedPhoto2 struct {
	win32.IInspectable
}

func (this *IAdvancedCapturedPhoto2) Vtbl() *IAdvancedCapturedPhoto2Vtbl {
	return (*IAdvancedCapturedPhoto2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdvancedCapturedPhoto2) Get_FrameBoundsRelativeToReferencePhoto() *IReference[Rect] {
	var _result *IReference[Rect]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameBoundsRelativeToReferencePhoto, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 83FFAAFA-6667-44DC-973C-A6BCE596AA0F
var IID_IAdvancedPhotoCapture = syscall.GUID{0x83FFAAFA, 0x6667, 0x44DC,
	[8]byte{0x97, 0x3C, 0xA6, 0xBC, 0xE5, 0x96, 0xAA, 0x0F}}

type IAdvancedPhotoCaptureInterface interface {
	win32.IInspectableInterface
	CaptureAsync() *IAsyncOperation[*IAdvancedCapturedPhoto]
	CaptureWithContextAsync(context interface{}) *IAsyncOperation[*IAdvancedCapturedPhoto]
	Add_OptionalReferencePhotoCaptured(handler TypedEventHandler[*IAdvancedPhotoCapture, *IOptionalReferencePhotoCapturedEventArgs]) EventRegistrationToken
	Remove_OptionalReferencePhotoCaptured(token EventRegistrationToken)
	Add_AllPhotosCaptured(handler TypedEventHandler[*IAdvancedPhotoCapture, interface{}]) EventRegistrationToken
	Remove_AllPhotosCaptured(token EventRegistrationToken)
	FinishAsync() *IAsyncAction
}

type IAdvancedPhotoCaptureVtbl struct {
	win32.IInspectableVtbl
	CaptureAsync                          uintptr
	CaptureWithContextAsync               uintptr
	Add_OptionalReferencePhotoCaptured    uintptr
	Remove_OptionalReferencePhotoCaptured uintptr
	Add_AllPhotosCaptured                 uintptr
	Remove_AllPhotosCaptured              uintptr
	FinishAsync                           uintptr
}

type IAdvancedPhotoCapture struct {
	win32.IInspectable
}

func (this *IAdvancedPhotoCapture) Vtbl() *IAdvancedPhotoCaptureVtbl {
	return (*IAdvancedPhotoCaptureVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdvancedPhotoCapture) CaptureAsync() *IAsyncOperation[*IAdvancedCapturedPhoto] {
	var _result *IAsyncOperation[*IAdvancedCapturedPhoto]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CaptureAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdvancedPhotoCapture) CaptureWithContextAsync(context interface{}) *IAsyncOperation[*IAdvancedCapturedPhoto] {
	var _result *IAsyncOperation[*IAdvancedCapturedPhoto]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CaptureWithContextAsync, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&context)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdvancedPhotoCapture) Add_OptionalReferencePhotoCaptured(handler TypedEventHandler[*IAdvancedPhotoCapture, *IOptionalReferencePhotoCapturedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_OptionalReferencePhotoCaptured, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdvancedPhotoCapture) Remove_OptionalReferencePhotoCaptured(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_OptionalReferencePhotoCaptured, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAdvancedPhotoCapture) Add_AllPhotosCaptured(handler TypedEventHandler[*IAdvancedPhotoCapture, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AllPhotosCaptured, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdvancedPhotoCapture) Remove_AllPhotosCaptured(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AllPhotosCaptured, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAdvancedPhotoCapture) FinishAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FinishAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BAD1E72A-FA94-46F9-95FC-D71511CDA70B
var IID_IAppBroadcastBackgroundService = syscall.GUID{0xBAD1E72A, 0xFA94, 0x46F9,
	[8]byte{0x95, 0xFC, 0xD7, 0x15, 0x11, 0xCD, 0xA7, 0x0B}}

type IAppBroadcastBackgroundServiceInterface interface {
	win32.IInspectableInterface
	Put_PlugInState(value AppBroadcastPlugInState)
	Get_PlugInState() AppBroadcastPlugInState
	Put_SignInInfo(value *IAppBroadcastBackgroundServiceSignInInfo)
	Get_SignInInfo() *IAppBroadcastBackgroundServiceSignInInfo
	Put_StreamInfo(value *IAppBroadcastBackgroundServiceStreamInfo)
	Get_StreamInfo() *IAppBroadcastBackgroundServiceStreamInfo
	Get_AppId() string
	Get_BroadcastTitle() string
	Put_ViewerCount(value uint32)
	Get_ViewerCount() uint32
	TerminateBroadcast(reason AppBroadcastTerminationReason, providerSpecificReason uint32)
	Add_HeartbeatRequested(handler TypedEventHandler[*IAppBroadcastBackgroundService, *IAppBroadcastHeartbeatRequestedEventArgs]) EventRegistrationToken
	Remove_HeartbeatRequested(token EventRegistrationToken)
	Get_TitleId() string
}

type IAppBroadcastBackgroundServiceVtbl struct {
	win32.IInspectableVtbl
	Put_PlugInState           uintptr
	Get_PlugInState           uintptr
	Put_SignInInfo            uintptr
	Get_SignInInfo            uintptr
	Put_StreamInfo            uintptr
	Get_StreamInfo            uintptr
	Get_AppId                 uintptr
	Get_BroadcastTitle        uintptr
	Put_ViewerCount           uintptr
	Get_ViewerCount           uintptr
	TerminateBroadcast        uintptr
	Add_HeartbeatRequested    uintptr
	Remove_HeartbeatRequested uintptr
	Get_TitleId               uintptr
}

type IAppBroadcastBackgroundService struct {
	win32.IInspectable
}

func (this *IAppBroadcastBackgroundService) Vtbl() *IAppBroadcastBackgroundServiceVtbl {
	return (*IAppBroadcastBackgroundServiceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastBackgroundService) Put_PlugInState(value AppBroadcastPlugInState) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PlugInState, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppBroadcastBackgroundService) Get_PlugInState() AppBroadcastPlugInState {
	var _result AppBroadcastPlugInState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlugInState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastBackgroundService) Put_SignInInfo(value *IAppBroadcastBackgroundServiceSignInInfo) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SignInInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAppBroadcastBackgroundService) Get_SignInInfo() *IAppBroadcastBackgroundServiceSignInInfo {
	var _result *IAppBroadcastBackgroundServiceSignInInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SignInInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppBroadcastBackgroundService) Put_StreamInfo(value *IAppBroadcastBackgroundServiceStreamInfo) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StreamInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAppBroadcastBackgroundService) Get_StreamInfo() *IAppBroadcastBackgroundServiceStreamInfo {
	var _result *IAppBroadcastBackgroundServiceStreamInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StreamInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppBroadcastBackgroundService) Get_AppId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppBroadcastBackgroundService) Get_BroadcastTitle() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BroadcastTitle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppBroadcastBackgroundService) Put_ViewerCount(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ViewerCount, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppBroadcastBackgroundService) Get_ViewerCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ViewerCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastBackgroundService) TerminateBroadcast(reason AppBroadcastTerminationReason, providerSpecificReason uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TerminateBroadcast, uintptr(unsafe.Pointer(this)), uintptr(reason), uintptr(providerSpecificReason))
	_ = _hr
}

func (this *IAppBroadcastBackgroundService) Add_HeartbeatRequested(handler TypedEventHandler[*IAppBroadcastBackgroundService, *IAppBroadcastHeartbeatRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_HeartbeatRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastBackgroundService) Remove_HeartbeatRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_HeartbeatRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppBroadcastBackgroundService) Get_TitleId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TitleId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// FC8CCBBF-5549-4B87-959F-23CA401FD473
var IID_IAppBroadcastBackgroundService2 = syscall.GUID{0xFC8CCBBF, 0x5549, 0x4B87,
	[8]byte{0x95, 0x9F, 0x23, 0xCA, 0x40, 0x1F, 0xD4, 0x73}}

type IAppBroadcastBackgroundService2Interface interface {
	win32.IInspectableInterface
	Put_BroadcastTitle(value string)
	Get_BroadcastLanguage() string
	Put_BroadcastLanguage(value string)
	Get_BroadcastChannel() string
	Put_BroadcastChannel(value string)
	Add_BroadcastTitleChanged(handler TypedEventHandler[*IAppBroadcastBackgroundService, interface{}]) EventRegistrationToken
	Remove_BroadcastTitleChanged(token EventRegistrationToken)
	Add_BroadcastLanguageChanged(handler TypedEventHandler[*IAppBroadcastBackgroundService, interface{}]) EventRegistrationToken
	Remove_BroadcastLanguageChanged(token EventRegistrationToken)
	Add_BroadcastChannelChanged(handler TypedEventHandler[*IAppBroadcastBackgroundService, interface{}]) EventRegistrationToken
	Remove_BroadcastChannelChanged(token EventRegistrationToken)
}

type IAppBroadcastBackgroundService2Vtbl struct {
	win32.IInspectableVtbl
	Put_BroadcastTitle              uintptr
	Get_BroadcastLanguage           uintptr
	Put_BroadcastLanguage           uintptr
	Get_BroadcastChannel            uintptr
	Put_BroadcastChannel            uintptr
	Add_BroadcastTitleChanged       uintptr
	Remove_BroadcastTitleChanged    uintptr
	Add_BroadcastLanguageChanged    uintptr
	Remove_BroadcastLanguageChanged uintptr
	Add_BroadcastChannelChanged     uintptr
	Remove_BroadcastChannelChanged  uintptr
}

type IAppBroadcastBackgroundService2 struct {
	win32.IInspectable
}

func (this *IAppBroadcastBackgroundService2) Vtbl() *IAppBroadcastBackgroundService2Vtbl {
	return (*IAppBroadcastBackgroundService2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastBackgroundService2) Put_BroadcastTitle(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BroadcastTitle, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IAppBroadcastBackgroundService2) Get_BroadcastLanguage() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BroadcastLanguage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppBroadcastBackgroundService2) Put_BroadcastLanguage(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BroadcastLanguage, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IAppBroadcastBackgroundService2) Get_BroadcastChannel() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BroadcastChannel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppBroadcastBackgroundService2) Put_BroadcastChannel(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BroadcastChannel, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IAppBroadcastBackgroundService2) Add_BroadcastTitleChanged(handler TypedEventHandler[*IAppBroadcastBackgroundService, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_BroadcastTitleChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastBackgroundService2) Remove_BroadcastTitleChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_BroadcastTitleChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppBroadcastBackgroundService2) Add_BroadcastLanguageChanged(handler TypedEventHandler[*IAppBroadcastBackgroundService, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_BroadcastLanguageChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastBackgroundService2) Remove_BroadcastLanguageChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_BroadcastLanguageChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppBroadcastBackgroundService2) Add_BroadcastChannelChanged(handler TypedEventHandler[*IAppBroadcastBackgroundService, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_BroadcastChannelChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastBackgroundService2) Remove_BroadcastChannelChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_BroadcastChannelChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 5E735275-88C8-4ECA-89BA-4825985DB880
var IID_IAppBroadcastBackgroundServiceSignInInfo = syscall.GUID{0x5E735275, 0x88C8, 0x4ECA,
	[8]byte{0x89, 0xBA, 0x48, 0x25, 0x98, 0x5D, 0xB8, 0x80}}

type IAppBroadcastBackgroundServiceSignInInfoInterface interface {
	win32.IInspectableInterface
	Get_SignInState() AppBroadcastSignInState
	Put_OAuthRequestUri(value *IUriRuntimeClass)
	Get_OAuthRequestUri() *IUriRuntimeClass
	Put_OAuthCallbackUri(value *IUriRuntimeClass)
	Get_OAuthCallbackUri() *IUriRuntimeClass
	Get_AuthenticationResult() unsafe.Pointer
	Put_UserName(value string)
	Get_UserName() string
	Add_SignInStateChanged(handler TypedEventHandler[*IAppBroadcastBackgroundServiceSignInInfo, *IAppBroadcastSignInStateChangedEventArgs]) EventRegistrationToken
	Remove_SignInStateChanged(token EventRegistrationToken)
}

type IAppBroadcastBackgroundServiceSignInInfoVtbl struct {
	win32.IInspectableVtbl
	Get_SignInState           uintptr
	Put_OAuthRequestUri       uintptr
	Get_OAuthRequestUri       uintptr
	Put_OAuthCallbackUri      uintptr
	Get_OAuthCallbackUri      uintptr
	Get_AuthenticationResult  uintptr
	Put_UserName              uintptr
	Get_UserName              uintptr
	Add_SignInStateChanged    uintptr
	Remove_SignInStateChanged uintptr
}

type IAppBroadcastBackgroundServiceSignInInfo struct {
	win32.IInspectable
}

func (this *IAppBroadcastBackgroundServiceSignInInfo) Vtbl() *IAppBroadcastBackgroundServiceSignInInfoVtbl {
	return (*IAppBroadcastBackgroundServiceSignInInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastBackgroundServiceSignInInfo) Get_SignInState() AppBroadcastSignInState {
	var _result AppBroadcastSignInState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SignInState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastBackgroundServiceSignInInfo) Put_OAuthRequestUri(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_OAuthRequestUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAppBroadcastBackgroundServiceSignInInfo) Get_OAuthRequestUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OAuthRequestUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppBroadcastBackgroundServiceSignInInfo) Put_OAuthCallbackUri(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_OAuthCallbackUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAppBroadcastBackgroundServiceSignInInfo) Get_OAuthCallbackUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OAuthCallbackUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppBroadcastBackgroundServiceSignInInfo) Get_AuthenticationResult() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AuthenticationResult, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastBackgroundServiceSignInInfo) Put_UserName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_UserName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IAppBroadcastBackgroundServiceSignInInfo) Get_UserName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UserName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppBroadcastBackgroundServiceSignInInfo) Add_SignInStateChanged(handler TypedEventHandler[*IAppBroadcastBackgroundServiceSignInInfo, *IAppBroadcastSignInStateChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SignInStateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastBackgroundServiceSignInInfo) Remove_SignInStateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SignInStateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 9104285C-62CF-4A3C-A7EE-AEB507404645
var IID_IAppBroadcastBackgroundServiceSignInInfo2 = syscall.GUID{0x9104285C, 0x62CF, 0x4A3C,
	[8]byte{0xA7, 0xEE, 0xAE, 0xB5, 0x07, 0x40, 0x46, 0x45}}

type IAppBroadcastBackgroundServiceSignInInfo2Interface interface {
	win32.IInspectableInterface
	Add_UserNameChanged(handler TypedEventHandler[*IAppBroadcastBackgroundServiceSignInInfo, interface{}]) EventRegistrationToken
	Remove_UserNameChanged(token EventRegistrationToken)
}

type IAppBroadcastBackgroundServiceSignInInfo2Vtbl struct {
	win32.IInspectableVtbl
	Add_UserNameChanged    uintptr
	Remove_UserNameChanged uintptr
}

type IAppBroadcastBackgroundServiceSignInInfo2 struct {
	win32.IInspectable
}

func (this *IAppBroadcastBackgroundServiceSignInInfo2) Vtbl() *IAppBroadcastBackgroundServiceSignInInfo2Vtbl {
	return (*IAppBroadcastBackgroundServiceSignInInfo2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastBackgroundServiceSignInInfo2) Add_UserNameChanged(handler TypedEventHandler[*IAppBroadcastBackgroundServiceSignInInfo, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_UserNameChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastBackgroundServiceSignInInfo2) Remove_UserNameChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_UserNameChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 31DC02BC-990A-4904-AA96-FE364381F136
var IID_IAppBroadcastBackgroundServiceStreamInfo = syscall.GUID{0x31DC02BC, 0x990A, 0x4904,
	[8]byte{0xAA, 0x96, 0xFE, 0x36, 0x43, 0x81, 0xF1, 0x36}}

type IAppBroadcastBackgroundServiceStreamInfoInterface interface {
	win32.IInspectableInterface
	Get_StreamState() AppBroadcastStreamState
	Put_DesiredVideoEncodingBitrate(value uint64)
	Get_DesiredVideoEncodingBitrate() uint64
	Put_BandwidthTestBitrate(value uint64)
	Get_BandwidthTestBitrate() uint64
	Put_AudioCodec(value string)
	Get_AudioCodec() string
	Get_BroadcastStreamReader() *IAppBroadcastStreamReader
	Add_StreamStateChanged(handler TypedEventHandler[*IAppBroadcastBackgroundServiceStreamInfo, *IAppBroadcastStreamStateChangedEventArgs]) EventRegistrationToken
	Remove_StreamStateChanged(token EventRegistrationToken)
	Add_VideoEncodingResolutionChanged(handler TypedEventHandler[*IAppBroadcastBackgroundServiceStreamInfo, interface{}]) EventRegistrationToken
	Remove_VideoEncodingResolutionChanged(token EventRegistrationToken)
	Add_VideoEncodingBitrateChanged(handler TypedEventHandler[*IAppBroadcastBackgroundServiceStreamInfo, interface{}]) EventRegistrationToken
	Remove_VideoEncodingBitrateChanged(token EventRegistrationToken)
}

type IAppBroadcastBackgroundServiceStreamInfoVtbl struct {
	win32.IInspectableVtbl
	Get_StreamState                       uintptr
	Put_DesiredVideoEncodingBitrate       uintptr
	Get_DesiredVideoEncodingBitrate       uintptr
	Put_BandwidthTestBitrate              uintptr
	Get_BandwidthTestBitrate              uintptr
	Put_AudioCodec                        uintptr
	Get_AudioCodec                        uintptr
	Get_BroadcastStreamReader             uintptr
	Add_StreamStateChanged                uintptr
	Remove_StreamStateChanged             uintptr
	Add_VideoEncodingResolutionChanged    uintptr
	Remove_VideoEncodingResolutionChanged uintptr
	Add_VideoEncodingBitrateChanged       uintptr
	Remove_VideoEncodingBitrateChanged    uintptr
}

type IAppBroadcastBackgroundServiceStreamInfo struct {
	win32.IInspectable
}

func (this *IAppBroadcastBackgroundServiceStreamInfo) Vtbl() *IAppBroadcastBackgroundServiceStreamInfoVtbl {
	return (*IAppBroadcastBackgroundServiceStreamInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastBackgroundServiceStreamInfo) Get_StreamState() AppBroadcastStreamState {
	var _result AppBroadcastStreamState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StreamState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastBackgroundServiceStreamInfo) Put_DesiredVideoEncodingBitrate(value uint64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DesiredVideoEncodingBitrate, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppBroadcastBackgroundServiceStreamInfo) Get_DesiredVideoEncodingBitrate() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesiredVideoEncodingBitrate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastBackgroundServiceStreamInfo) Put_BandwidthTestBitrate(value uint64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BandwidthTestBitrate, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppBroadcastBackgroundServiceStreamInfo) Get_BandwidthTestBitrate() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BandwidthTestBitrate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastBackgroundServiceStreamInfo) Put_AudioCodec(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AudioCodec, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IAppBroadcastBackgroundServiceStreamInfo) Get_AudioCodec() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioCodec, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppBroadcastBackgroundServiceStreamInfo) Get_BroadcastStreamReader() *IAppBroadcastStreamReader {
	var _result *IAppBroadcastStreamReader
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BroadcastStreamReader, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppBroadcastBackgroundServiceStreamInfo) Add_StreamStateChanged(handler TypedEventHandler[*IAppBroadcastBackgroundServiceStreamInfo, *IAppBroadcastStreamStateChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StreamStateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastBackgroundServiceStreamInfo) Remove_StreamStateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StreamStateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppBroadcastBackgroundServiceStreamInfo) Add_VideoEncodingResolutionChanged(handler TypedEventHandler[*IAppBroadcastBackgroundServiceStreamInfo, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_VideoEncodingResolutionChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastBackgroundServiceStreamInfo) Remove_VideoEncodingResolutionChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_VideoEncodingResolutionChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppBroadcastBackgroundServiceStreamInfo) Add_VideoEncodingBitrateChanged(handler TypedEventHandler[*IAppBroadcastBackgroundServiceStreamInfo, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_VideoEncodingBitrateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastBackgroundServiceStreamInfo) Remove_VideoEncodingBitrateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_VideoEncodingBitrateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// BD1E9F6D-94DC-4FCE-9541-A9F129596334
var IID_IAppBroadcastBackgroundServiceStreamInfo2 = syscall.GUID{0xBD1E9F6D, 0x94DC, 0x4FCE,
	[8]byte{0x95, 0x41, 0xA9, 0xF1, 0x29, 0x59, 0x63, 0x34}}

type IAppBroadcastBackgroundServiceStreamInfo2Interface interface {
	win32.IInspectableInterface
	ReportProblemWithStream()
}

type IAppBroadcastBackgroundServiceStreamInfo2Vtbl struct {
	win32.IInspectableVtbl
	ReportProblemWithStream uintptr
}

type IAppBroadcastBackgroundServiceStreamInfo2 struct {
	win32.IInspectable
}

func (this *IAppBroadcastBackgroundServiceStreamInfo2) Vtbl() *IAppBroadcastBackgroundServiceStreamInfo2Vtbl {
	return (*IAppBroadcastBackgroundServiceStreamInfo2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastBackgroundServiceStreamInfo2) ReportProblemWithStream() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReportProblemWithStream, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 1E334CD0-B882-4B88-8692-05999ACEB70F
var IID_IAppBroadcastCameraCaptureStateChangedEventArgs = syscall.GUID{0x1E334CD0, 0xB882, 0x4B88,
	[8]byte{0x86, 0x92, 0x05, 0x99, 0x9A, 0xCE, 0xB7, 0x0F}}

type IAppBroadcastCameraCaptureStateChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_State() AppBroadcastCameraCaptureState
	Get_ErrorCode() uint32
}

type IAppBroadcastCameraCaptureStateChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_State     uintptr
	Get_ErrorCode uintptr
}

type IAppBroadcastCameraCaptureStateChangedEventArgs struct {
	win32.IInspectable
}

func (this *IAppBroadcastCameraCaptureStateChangedEventArgs) Vtbl() *IAppBroadcastCameraCaptureStateChangedEventArgsVtbl {
	return (*IAppBroadcastCameraCaptureStateChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastCameraCaptureStateChangedEventArgs) Get_State() AppBroadcastCameraCaptureState {
	var _result AppBroadcastCameraCaptureState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastCameraCaptureStateChangedEventArgs) Get_ErrorCode() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ErrorCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// B2CB27A5-70FC-4E17-80BD-6BA0FD3FF3A0
var IID_IAppBroadcastGlobalSettings = syscall.GUID{0xB2CB27A5, 0x70FC, 0x4E17,
	[8]byte{0x80, 0xBD, 0x6B, 0xA0, 0xFD, 0x3F, 0xF3, 0xA0}}

type IAppBroadcastGlobalSettingsInterface interface {
	win32.IInspectableInterface
	Get_IsBroadcastEnabled() bool
	Get_IsDisabledByPolicy() bool
	Get_IsGpuConstrained() bool
	Get_HasHardwareEncoder() bool
	Put_IsAudioCaptureEnabled(value bool)
	Get_IsAudioCaptureEnabled() bool
	Put_IsMicrophoneCaptureEnabledByDefault(value bool)
	Get_IsMicrophoneCaptureEnabledByDefault() bool
	Put_IsEchoCancellationEnabled(value bool)
	Get_IsEchoCancellationEnabled() bool
	Put_SystemAudioGain(value float64)
	Get_SystemAudioGain() float64
	Put_MicrophoneGain(value float64)
	Get_MicrophoneGain() float64
	Put_IsCameraCaptureEnabledByDefault(value bool)
	Get_IsCameraCaptureEnabledByDefault() bool
	Put_SelectedCameraId(value string)
	Get_SelectedCameraId() string
	Put_CameraOverlayLocation(value AppBroadcastCameraOverlayLocation)
	Get_CameraOverlayLocation() AppBroadcastCameraOverlayLocation
	Put_CameraOverlaySize(value AppBroadcastCameraOverlaySize)
	Get_CameraOverlaySize() AppBroadcastCameraOverlaySize
	Put_IsCursorImageCaptureEnabled(value bool)
	Get_IsCursorImageCaptureEnabled() bool
}

type IAppBroadcastGlobalSettingsVtbl struct {
	win32.IInspectableVtbl
	Get_IsBroadcastEnabled                  uintptr
	Get_IsDisabledByPolicy                  uintptr
	Get_IsGpuConstrained                    uintptr
	Get_HasHardwareEncoder                  uintptr
	Put_IsAudioCaptureEnabled               uintptr
	Get_IsAudioCaptureEnabled               uintptr
	Put_IsMicrophoneCaptureEnabledByDefault uintptr
	Get_IsMicrophoneCaptureEnabledByDefault uintptr
	Put_IsEchoCancellationEnabled           uintptr
	Get_IsEchoCancellationEnabled           uintptr
	Put_SystemAudioGain                     uintptr
	Get_SystemAudioGain                     uintptr
	Put_MicrophoneGain                      uintptr
	Get_MicrophoneGain                      uintptr
	Put_IsCameraCaptureEnabledByDefault     uintptr
	Get_IsCameraCaptureEnabledByDefault     uintptr
	Put_SelectedCameraId                    uintptr
	Get_SelectedCameraId                    uintptr
	Put_CameraOverlayLocation               uintptr
	Get_CameraOverlayLocation               uintptr
	Put_CameraOverlaySize                   uintptr
	Get_CameraOverlaySize                   uintptr
	Put_IsCursorImageCaptureEnabled         uintptr
	Get_IsCursorImageCaptureEnabled         uintptr
}

type IAppBroadcastGlobalSettings struct {
	win32.IInspectable
}

func (this *IAppBroadcastGlobalSettings) Vtbl() *IAppBroadcastGlobalSettingsVtbl {
	return (*IAppBroadcastGlobalSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastGlobalSettings) Get_IsBroadcastEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsBroadcastEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastGlobalSettings) Get_IsDisabledByPolicy() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDisabledByPolicy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastGlobalSettings) Get_IsGpuConstrained() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsGpuConstrained, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastGlobalSettings) Get_HasHardwareEncoder() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasHardwareEncoder, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastGlobalSettings) Put_IsAudioCaptureEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsAudioCaptureEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAppBroadcastGlobalSettings) Get_IsAudioCaptureEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsAudioCaptureEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastGlobalSettings) Put_IsMicrophoneCaptureEnabledByDefault(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsMicrophoneCaptureEnabledByDefault, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAppBroadcastGlobalSettings) Get_IsMicrophoneCaptureEnabledByDefault() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsMicrophoneCaptureEnabledByDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastGlobalSettings) Put_IsEchoCancellationEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsEchoCancellationEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAppBroadcastGlobalSettings) Get_IsEchoCancellationEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEchoCancellationEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastGlobalSettings) Put_SystemAudioGain(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SystemAudioGain, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppBroadcastGlobalSettings) Get_SystemAudioGain() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SystemAudioGain, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastGlobalSettings) Put_MicrophoneGain(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MicrophoneGain, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppBroadcastGlobalSettings) Get_MicrophoneGain() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MicrophoneGain, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastGlobalSettings) Put_IsCameraCaptureEnabledByDefault(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsCameraCaptureEnabledByDefault, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAppBroadcastGlobalSettings) Get_IsCameraCaptureEnabledByDefault() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCameraCaptureEnabledByDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastGlobalSettings) Put_SelectedCameraId(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SelectedCameraId, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IAppBroadcastGlobalSettings) Get_SelectedCameraId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SelectedCameraId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppBroadcastGlobalSettings) Put_CameraOverlayLocation(value AppBroadcastCameraOverlayLocation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CameraOverlayLocation, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppBroadcastGlobalSettings) Get_CameraOverlayLocation() AppBroadcastCameraOverlayLocation {
	var _result AppBroadcastCameraOverlayLocation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CameraOverlayLocation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastGlobalSettings) Put_CameraOverlaySize(value AppBroadcastCameraOverlaySize) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CameraOverlaySize, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppBroadcastGlobalSettings) Get_CameraOverlaySize() AppBroadcastCameraOverlaySize {
	var _result AppBroadcastCameraOverlaySize
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CameraOverlaySize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastGlobalSettings) Put_IsCursorImageCaptureEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsCursorImageCaptureEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAppBroadcastGlobalSettings) Get_IsCursorImageCaptureEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCursorImageCaptureEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// CEA54283-EE51-4DBF-9472-79A9ED4E2165
var IID_IAppBroadcastHeartbeatRequestedEventArgs = syscall.GUID{0xCEA54283, 0xEE51, 0x4DBF,
	[8]byte{0x94, 0x72, 0x79, 0xA9, 0xED, 0x4E, 0x21, 0x65}}

type IAppBroadcastHeartbeatRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	Put_Handled(value bool)
	Get_Handled() bool
}

type IAppBroadcastHeartbeatRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Put_Handled uintptr
	Get_Handled uintptr
}

type IAppBroadcastHeartbeatRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IAppBroadcastHeartbeatRequestedEventArgs) Vtbl() *IAppBroadcastHeartbeatRequestedEventArgsVtbl {
	return (*IAppBroadcastHeartbeatRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastHeartbeatRequestedEventArgs) Put_Handled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Handled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAppBroadcastHeartbeatRequestedEventArgs) Get_Handled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Handled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 364E018B-1E4E-411F-AB3E-92959844C156
var IID_IAppBroadcastManagerStatics = syscall.GUID{0x364E018B, 0x1E4E, 0x411F,
	[8]byte{0xAB, 0x3E, 0x92, 0x95, 0x98, 0x44, 0xC1, 0x56}}

type IAppBroadcastManagerStaticsInterface interface {
	win32.IInspectableInterface
	GetGlobalSettings() *IAppBroadcastGlobalSettings
	ApplyGlobalSettings(value *IAppBroadcastGlobalSettings)
	GetProviderSettings() *IAppBroadcastProviderSettings
	ApplyProviderSettings(value *IAppBroadcastProviderSettings)
}

type IAppBroadcastManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetGlobalSettings     uintptr
	ApplyGlobalSettings   uintptr
	GetProviderSettings   uintptr
	ApplyProviderSettings uintptr
}

type IAppBroadcastManagerStatics struct {
	win32.IInspectable
}

func (this *IAppBroadcastManagerStatics) Vtbl() *IAppBroadcastManagerStaticsVtbl {
	return (*IAppBroadcastManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastManagerStatics) GetGlobalSettings() *IAppBroadcastGlobalSettings {
	var _result *IAppBroadcastGlobalSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetGlobalSettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppBroadcastManagerStatics) ApplyGlobalSettings(value *IAppBroadcastGlobalSettings) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ApplyGlobalSettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAppBroadcastManagerStatics) GetProviderSettings() *IAppBroadcastProviderSettings {
	var _result *IAppBroadcastProviderSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetProviderSettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppBroadcastManagerStatics) ApplyProviderSettings(value *IAppBroadcastProviderSettings) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ApplyProviderSettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// A86AD5E9-9440-4908-9D09-65B7E315D795
var IID_IAppBroadcastMicrophoneCaptureStateChangedEventArgs = syscall.GUID{0xA86AD5E9, 0x9440, 0x4908,
	[8]byte{0x9D, 0x09, 0x65, 0xB7, 0xE3, 0x15, 0xD7, 0x95}}

type IAppBroadcastMicrophoneCaptureStateChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_State() AppBroadcastMicrophoneCaptureState
	Get_ErrorCode() uint32
}

type IAppBroadcastMicrophoneCaptureStateChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_State     uintptr
	Get_ErrorCode uintptr
}

type IAppBroadcastMicrophoneCaptureStateChangedEventArgs struct {
	win32.IInspectable
}

func (this *IAppBroadcastMicrophoneCaptureStateChangedEventArgs) Vtbl() *IAppBroadcastMicrophoneCaptureStateChangedEventArgsVtbl {
	return (*IAppBroadcastMicrophoneCaptureStateChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastMicrophoneCaptureStateChangedEventArgs) Get_State() AppBroadcastMicrophoneCaptureState {
	var _result AppBroadcastMicrophoneCaptureState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastMicrophoneCaptureStateChangedEventArgs) Get_ErrorCode() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ErrorCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 520C1E66-6513-4574-AC54-23B79729615B
var IID_IAppBroadcastPlugIn = syscall.GUID{0x520C1E66, 0x6513, 0x4574,
	[8]byte{0xAC, 0x54, 0x23, 0xB7, 0x97, 0x29, 0x61, 0x5B}}

type IAppBroadcastPlugInInterface interface {
	win32.IInspectableInterface
	Get_AppId() string
	Get_ProviderSettings() *IAppBroadcastProviderSettings
	Get_Logo() *IRandomAccessStreamReference
	Get_DisplayName() string
}

type IAppBroadcastPlugInVtbl struct {
	win32.IInspectableVtbl
	Get_AppId            uintptr
	Get_ProviderSettings uintptr
	Get_Logo             uintptr
	Get_DisplayName      uintptr
}

type IAppBroadcastPlugIn struct {
	win32.IInspectable
}

func (this *IAppBroadcastPlugIn) Vtbl() *IAppBroadcastPlugInVtbl {
	return (*IAppBroadcastPlugInVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastPlugIn) Get_AppId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppBroadcastPlugIn) Get_ProviderSettings() *IAppBroadcastProviderSettings {
	var _result *IAppBroadcastProviderSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProviderSettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppBroadcastPlugIn) Get_Logo() *IRandomAccessStreamReference {
	var _result *IRandomAccessStreamReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Logo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppBroadcastPlugIn) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// E550D979-27A1-49A7-BBF4-D7A9E9D07668
var IID_IAppBroadcastPlugInManager = syscall.GUID{0xE550D979, 0x27A1, 0x49A7,
	[8]byte{0xBB, 0xF4, 0xD7, 0xA9, 0xE9, 0xD0, 0x76, 0x68}}

type IAppBroadcastPlugInManagerInterface interface {
	win32.IInspectableInterface
	Get_IsBroadcastProviderAvailable() bool
	Get_PlugInList() *IVectorView[*IAppBroadcastPlugIn]
	Get_DefaultPlugIn() *IAppBroadcastPlugIn
	Put_DefaultPlugIn(value *IAppBroadcastPlugIn)
}

type IAppBroadcastPlugInManagerVtbl struct {
	win32.IInspectableVtbl
	Get_IsBroadcastProviderAvailable uintptr
	Get_PlugInList                   uintptr
	Get_DefaultPlugIn                uintptr
	Put_DefaultPlugIn                uintptr
}

type IAppBroadcastPlugInManager struct {
	win32.IInspectable
}

func (this *IAppBroadcastPlugInManager) Vtbl() *IAppBroadcastPlugInManagerVtbl {
	return (*IAppBroadcastPlugInManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastPlugInManager) Get_IsBroadcastProviderAvailable() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsBroadcastProviderAvailable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastPlugInManager) Get_PlugInList() *IVectorView[*IAppBroadcastPlugIn] {
	var _result *IVectorView[*IAppBroadcastPlugIn]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlugInList, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppBroadcastPlugInManager) Get_DefaultPlugIn() *IAppBroadcastPlugIn {
	var _result *IAppBroadcastPlugIn
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DefaultPlugIn, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppBroadcastPlugInManager) Put_DefaultPlugIn(value *IAppBroadcastPlugIn) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DefaultPlugIn, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// F2645C20-5C76-4CDC-9364-82FE9EB6534D
var IID_IAppBroadcastPlugInManagerStatics = syscall.GUID{0xF2645C20, 0x5C76, 0x4CDC,
	[8]byte{0x93, 0x64, 0x82, 0xFE, 0x9E, 0xB6, 0x53, 0x4D}}

type IAppBroadcastPlugInManagerStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *IAppBroadcastPlugInManager
	GetForUser(user *IUser) *IAppBroadcastPlugInManager
}

type IAppBroadcastPlugInManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
	GetForUser uintptr
}

type IAppBroadcastPlugInManagerStatics struct {
	win32.IInspectable
}

func (this *IAppBroadcastPlugInManagerStatics) Vtbl() *IAppBroadcastPlugInManagerStaticsVtbl {
	return (*IAppBroadcastPlugInManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastPlugInManagerStatics) GetDefault() *IAppBroadcastPlugInManager {
	var _result *IAppBroadcastPlugInManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppBroadcastPlugInManagerStatics) GetForUser(user *IUser) *IAppBroadcastPlugInManager {
	var _result *IAppBroadcastPlugInManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4881D0F2-ABC5-4FC6-84B0-89370BB47212
var IID_IAppBroadcastPlugInStateChangedEventArgs = syscall.GUID{0x4881D0F2, 0xABC5, 0x4FC6,
	[8]byte{0x84, 0xB0, 0x89, 0x37, 0x0B, 0xB4, 0x72, 0x12}}

type IAppBroadcastPlugInStateChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_PlugInState() AppBroadcastPlugInState
}

type IAppBroadcastPlugInStateChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_PlugInState uintptr
}

type IAppBroadcastPlugInStateChangedEventArgs struct {
	win32.IInspectable
}

func (this *IAppBroadcastPlugInStateChangedEventArgs) Vtbl() *IAppBroadcastPlugInStateChangedEventArgsVtbl {
	return (*IAppBroadcastPlugInStateChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastPlugInStateChangedEventArgs) Get_PlugInState() AppBroadcastPlugInState {
	var _result AppBroadcastPlugInState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlugInState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 14B60F5A-6E4A-4B80-A14F-67EE77D153E7
var IID_IAppBroadcastPreview = syscall.GUID{0x14B60F5A, 0x6E4A, 0x4B80,
	[8]byte{0xA1, 0x4F, 0x67, 0xEE, 0x77, 0xD1, 0x53, 0xE7}}

type IAppBroadcastPreviewInterface interface {
	win32.IInspectableInterface
	StopPreview()
	Get_PreviewState() AppBroadcastPreviewState
	Get_ErrorCode() *IReference[uint32]
	Add_PreviewStateChanged(value TypedEventHandler[*IAppBroadcastPreview, *IAppBroadcastPreviewStateChangedEventArgs]) EventRegistrationToken
	Remove_PreviewStateChanged(token EventRegistrationToken)
	Get_PreviewStreamReader() *IAppBroadcastPreviewStreamReader
}

type IAppBroadcastPreviewVtbl struct {
	win32.IInspectableVtbl
	StopPreview                uintptr
	Get_PreviewState           uintptr
	Get_ErrorCode              uintptr
	Add_PreviewStateChanged    uintptr
	Remove_PreviewStateChanged uintptr
	Get_PreviewStreamReader    uintptr
}

type IAppBroadcastPreview struct {
	win32.IInspectable
}

func (this *IAppBroadcastPreview) Vtbl() *IAppBroadcastPreviewVtbl {
	return (*IAppBroadcastPreviewVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastPreview) StopPreview() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopPreview, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IAppBroadcastPreview) Get_PreviewState() AppBroadcastPreviewState {
	var _result AppBroadcastPreviewState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PreviewState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastPreview) Get_ErrorCode() *IReference[uint32] {
	var _result *IReference[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ErrorCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppBroadcastPreview) Add_PreviewStateChanged(value TypedEventHandler[*IAppBroadcastPreview, *IAppBroadcastPreviewStateChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PreviewStateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastPreview) Remove_PreviewStateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PreviewStateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppBroadcastPreview) Get_PreviewStreamReader() *IAppBroadcastPreviewStreamReader {
	var _result *IAppBroadcastPreviewStreamReader
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PreviewStreamReader, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5A57F2DE-8DEA-4E86-90AD-03FC26B9653C
var IID_IAppBroadcastPreviewStateChangedEventArgs = syscall.GUID{0x5A57F2DE, 0x8DEA, 0x4E86,
	[8]byte{0x90, 0xAD, 0x03, 0xFC, 0x26, 0xB9, 0x65, 0x3C}}

type IAppBroadcastPreviewStateChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_PreviewState() AppBroadcastPreviewState
	Get_ErrorCode() uint32
}

type IAppBroadcastPreviewStateChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_PreviewState uintptr
	Get_ErrorCode    uintptr
}

type IAppBroadcastPreviewStateChangedEventArgs struct {
	win32.IInspectable
}

func (this *IAppBroadcastPreviewStateChangedEventArgs) Vtbl() *IAppBroadcastPreviewStateChangedEventArgsVtbl {
	return (*IAppBroadcastPreviewStateChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastPreviewStateChangedEventArgs) Get_PreviewState() AppBroadcastPreviewState {
	var _result AppBroadcastPreviewState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PreviewState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastPreviewStateChangedEventArgs) Get_ErrorCode() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ErrorCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 92228D50-DB3F-40A8-8CD4-F4E371DDAB37
var IID_IAppBroadcastPreviewStreamReader = syscall.GUID{0x92228D50, 0xDB3F, 0x40A8,
	[8]byte{0x8C, 0xD4, 0xF4, 0xE3, 0x71, 0xDD, 0xAB, 0x37}}

type IAppBroadcastPreviewStreamReaderInterface interface {
	win32.IInspectableInterface
	Get_VideoWidth() uint32
	Get_VideoHeight() uint32
	Get_VideoStride() uint32
	Get_VideoBitmapPixelFormat() BitmapPixelFormat
	Get_VideoBitmapAlphaMode() BitmapAlphaMode
	TryGetNextVideoFrame() *IAppBroadcastPreviewStreamVideoFrame
	Add_VideoFrameArrived(value TypedEventHandler[*IAppBroadcastPreviewStreamReader, interface{}]) EventRegistrationToken
	Remove_VideoFrameArrived(token EventRegistrationToken)
}

type IAppBroadcastPreviewStreamReaderVtbl struct {
	win32.IInspectableVtbl
	Get_VideoWidth             uintptr
	Get_VideoHeight            uintptr
	Get_VideoStride            uintptr
	Get_VideoBitmapPixelFormat uintptr
	Get_VideoBitmapAlphaMode   uintptr
	TryGetNextVideoFrame       uintptr
	Add_VideoFrameArrived      uintptr
	Remove_VideoFrameArrived   uintptr
}

type IAppBroadcastPreviewStreamReader struct {
	win32.IInspectable
}

func (this *IAppBroadcastPreviewStreamReader) Vtbl() *IAppBroadcastPreviewStreamReaderVtbl {
	return (*IAppBroadcastPreviewStreamReaderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastPreviewStreamReader) Get_VideoWidth() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoWidth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastPreviewStreamReader) Get_VideoHeight() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoHeight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastPreviewStreamReader) Get_VideoStride() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoStride, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastPreviewStreamReader) Get_VideoBitmapPixelFormat() BitmapPixelFormat {
	var _result BitmapPixelFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoBitmapPixelFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastPreviewStreamReader) Get_VideoBitmapAlphaMode() BitmapAlphaMode {
	var _result BitmapAlphaMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoBitmapAlphaMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastPreviewStreamReader) TryGetNextVideoFrame() *IAppBroadcastPreviewStreamVideoFrame {
	var _result *IAppBroadcastPreviewStreamVideoFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetNextVideoFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppBroadcastPreviewStreamReader) Add_VideoFrameArrived(value TypedEventHandler[*IAppBroadcastPreviewStreamReader, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_VideoFrameArrived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastPreviewStreamReader) Remove_VideoFrameArrived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_VideoFrameArrived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 010FBEA1-94FE-4499-B8C0-8D244279FB12
var IID_IAppBroadcastPreviewStreamVideoFrame = syscall.GUID{0x010FBEA1, 0x94FE, 0x4499,
	[8]byte{0xB8, 0xC0, 0x8D, 0x24, 0x42, 0x79, 0xFB, 0x12}}

type IAppBroadcastPreviewStreamVideoFrameInterface interface {
	win32.IInspectableInterface
	Get_VideoHeader() *IAppBroadcastPreviewStreamVideoHeader
	Get_VideoBuffer() *IBuffer
}

type IAppBroadcastPreviewStreamVideoFrameVtbl struct {
	win32.IInspectableVtbl
	Get_VideoHeader uintptr
	Get_VideoBuffer uintptr
}

type IAppBroadcastPreviewStreamVideoFrame struct {
	win32.IInspectable
}

func (this *IAppBroadcastPreviewStreamVideoFrame) Vtbl() *IAppBroadcastPreviewStreamVideoFrameVtbl {
	return (*IAppBroadcastPreviewStreamVideoFrameVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastPreviewStreamVideoFrame) Get_VideoHeader() *IAppBroadcastPreviewStreamVideoHeader {
	var _result *IAppBroadcastPreviewStreamVideoHeader
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoHeader, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppBroadcastPreviewStreamVideoFrame) Get_VideoBuffer() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoBuffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8BEF6113-DA84-4499-A7AB-87118CB4A157
var IID_IAppBroadcastPreviewStreamVideoHeader = syscall.GUID{0x8BEF6113, 0xDA84, 0x4499,
	[8]byte{0xA7, 0xAB, 0x87, 0x11, 0x8C, 0xB4, 0xA1, 0x57}}

type IAppBroadcastPreviewStreamVideoHeaderInterface interface {
	win32.IInspectableInterface
	Get_AbsoluteTimestamp() DateTime
	Get_RelativeTimestamp() TimeSpan
	Get_Duration() TimeSpan
	Get_FrameId() uint64
}

type IAppBroadcastPreviewStreamVideoHeaderVtbl struct {
	win32.IInspectableVtbl
	Get_AbsoluteTimestamp uintptr
	Get_RelativeTimestamp uintptr
	Get_Duration          uintptr
	Get_FrameId           uintptr
}

type IAppBroadcastPreviewStreamVideoHeader struct {
	win32.IInspectable
}

func (this *IAppBroadcastPreviewStreamVideoHeader) Vtbl() *IAppBroadcastPreviewStreamVideoHeaderVtbl {
	return (*IAppBroadcastPreviewStreamVideoHeaderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastPreviewStreamVideoHeader) Get_AbsoluteTimestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AbsoluteTimestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastPreviewStreamVideoHeader) Get_RelativeTimestamp() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RelativeTimestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastPreviewStreamVideoHeader) Get_Duration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Duration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastPreviewStreamVideoHeader) Get_FrameId() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C30BDF62-9948-458F-AD50-AA06EC03DA08
var IID_IAppBroadcastProviderSettings = syscall.GUID{0xC30BDF62, 0x9948, 0x458F,
	[8]byte{0xAD, 0x50, 0xAA, 0x06, 0xEC, 0x03, 0xDA, 0x08}}

type IAppBroadcastProviderSettingsInterface interface {
	win32.IInspectableInterface
	Put_DefaultBroadcastTitle(value string)
	Get_DefaultBroadcastTitle() string
	Put_AudioEncodingBitrate(value uint32)
	Get_AudioEncodingBitrate() uint32
	Put_CustomVideoEncodingBitrate(value uint32)
	Get_CustomVideoEncodingBitrate() uint32
	Put_CustomVideoEncodingHeight(value uint32)
	Get_CustomVideoEncodingHeight() uint32
	Put_CustomVideoEncodingWidth(value uint32)
	Get_CustomVideoEncodingWidth() uint32
	Put_VideoEncodingBitrateMode(value AppBroadcastVideoEncodingBitrateMode)
	Get_VideoEncodingBitrateMode() AppBroadcastVideoEncodingBitrateMode
	Put_VideoEncodingResolutionMode(value AppBroadcastVideoEncodingResolutionMode)
	Get_VideoEncodingResolutionMode() AppBroadcastVideoEncodingResolutionMode
}

type IAppBroadcastProviderSettingsVtbl struct {
	win32.IInspectableVtbl
	Put_DefaultBroadcastTitle       uintptr
	Get_DefaultBroadcastTitle       uintptr
	Put_AudioEncodingBitrate        uintptr
	Get_AudioEncodingBitrate        uintptr
	Put_CustomVideoEncodingBitrate  uintptr
	Get_CustomVideoEncodingBitrate  uintptr
	Put_CustomVideoEncodingHeight   uintptr
	Get_CustomVideoEncodingHeight   uintptr
	Put_CustomVideoEncodingWidth    uintptr
	Get_CustomVideoEncodingWidth    uintptr
	Put_VideoEncodingBitrateMode    uintptr
	Get_VideoEncodingBitrateMode    uintptr
	Put_VideoEncodingResolutionMode uintptr
	Get_VideoEncodingResolutionMode uintptr
}

type IAppBroadcastProviderSettings struct {
	win32.IInspectable
}

func (this *IAppBroadcastProviderSettings) Vtbl() *IAppBroadcastProviderSettingsVtbl {
	return (*IAppBroadcastProviderSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastProviderSettings) Put_DefaultBroadcastTitle(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DefaultBroadcastTitle, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IAppBroadcastProviderSettings) Get_DefaultBroadcastTitle() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DefaultBroadcastTitle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppBroadcastProviderSettings) Put_AudioEncodingBitrate(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AudioEncodingBitrate, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppBroadcastProviderSettings) Get_AudioEncodingBitrate() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioEncodingBitrate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastProviderSettings) Put_CustomVideoEncodingBitrate(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CustomVideoEncodingBitrate, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppBroadcastProviderSettings) Get_CustomVideoEncodingBitrate() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CustomVideoEncodingBitrate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastProviderSettings) Put_CustomVideoEncodingHeight(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CustomVideoEncodingHeight, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppBroadcastProviderSettings) Get_CustomVideoEncodingHeight() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CustomVideoEncodingHeight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastProviderSettings) Put_CustomVideoEncodingWidth(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CustomVideoEncodingWidth, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppBroadcastProviderSettings) Get_CustomVideoEncodingWidth() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CustomVideoEncodingWidth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastProviderSettings) Put_VideoEncodingBitrateMode(value AppBroadcastVideoEncodingBitrateMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_VideoEncodingBitrateMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppBroadcastProviderSettings) Get_VideoEncodingBitrateMode() AppBroadcastVideoEncodingBitrateMode {
	var _result AppBroadcastVideoEncodingBitrateMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoEncodingBitrateMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastProviderSettings) Put_VideoEncodingResolutionMode(value AppBroadcastVideoEncodingResolutionMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_VideoEncodingResolutionMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppBroadcastProviderSettings) Get_VideoEncodingResolutionMode() AppBroadcastVideoEncodingResolutionMode {
	var _result AppBroadcastVideoEncodingResolutionMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoEncodingResolutionMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 8660B4D6-969B-4E3C-AC3A-8B042EE4EE63
var IID_IAppBroadcastServices = syscall.GUID{0x8660B4D6, 0x969B, 0x4E3C,
	[8]byte{0xAC, 0x3A, 0x8B, 0x04, 0x2E, 0xE4, 0xEE, 0x63}}

type IAppBroadcastServicesInterface interface {
	win32.IInspectableInterface
	Get_CaptureTargetType() AppBroadcastCaptureTargetType
	Put_CaptureTargetType(value AppBroadcastCaptureTargetType)
	Get_BroadcastTitle() string
	Put_BroadcastTitle(value string)
	Get_BroadcastLanguage() string
	Put_BroadcastLanguage(value string)
	Get_UserName() string
	Get_CanCapture() bool
	EnterBroadcastModeAsync(plugIn *IAppBroadcastPlugIn) *IAsyncOperation[uint32]
	ExitBroadcastMode(reason AppBroadcastExitBroadcastModeReason)
	StartBroadcast()
	PauseBroadcast()
	ResumeBroadcast()
	StartPreview(desiredSize Size) *IAppBroadcastPreview
	Get_State() *IAppBroadcastState
}

type IAppBroadcastServicesVtbl struct {
	win32.IInspectableVtbl
	Get_CaptureTargetType   uintptr
	Put_CaptureTargetType   uintptr
	Get_BroadcastTitle      uintptr
	Put_BroadcastTitle      uintptr
	Get_BroadcastLanguage   uintptr
	Put_BroadcastLanguage   uintptr
	Get_UserName            uintptr
	Get_CanCapture          uintptr
	EnterBroadcastModeAsync uintptr
	ExitBroadcastMode       uintptr
	StartBroadcast          uintptr
	PauseBroadcast          uintptr
	ResumeBroadcast         uintptr
	StartPreview            uintptr
	Get_State               uintptr
}

type IAppBroadcastServices struct {
	win32.IInspectable
}

func (this *IAppBroadcastServices) Vtbl() *IAppBroadcastServicesVtbl {
	return (*IAppBroadcastServicesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastServices) Get_CaptureTargetType() AppBroadcastCaptureTargetType {
	var _result AppBroadcastCaptureTargetType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CaptureTargetType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastServices) Put_CaptureTargetType(value AppBroadcastCaptureTargetType) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CaptureTargetType, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppBroadcastServices) Get_BroadcastTitle() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BroadcastTitle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppBroadcastServices) Put_BroadcastTitle(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BroadcastTitle, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IAppBroadcastServices) Get_BroadcastLanguage() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BroadcastLanguage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppBroadcastServices) Put_BroadcastLanguage(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BroadcastLanguage, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IAppBroadcastServices) Get_UserName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UserName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppBroadcastServices) Get_CanCapture() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanCapture, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastServices) EnterBroadcastModeAsync(plugIn *IAppBroadcastPlugIn) *IAsyncOperation[uint32] {
	var _result *IAsyncOperation[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnterBroadcastModeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(plugIn)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppBroadcastServices) ExitBroadcastMode(reason AppBroadcastExitBroadcastModeReason) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ExitBroadcastMode, uintptr(unsafe.Pointer(this)), uintptr(reason))
	_ = _hr
}

func (this *IAppBroadcastServices) StartBroadcast() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartBroadcast, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IAppBroadcastServices) PauseBroadcast() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PauseBroadcast, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IAppBroadcastServices) ResumeBroadcast() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ResumeBroadcast, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IAppBroadcastServices) StartPreview(desiredSize Size) *IAppBroadcastPreview {
	var _result *IAppBroadcastPreview
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartPreview, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&desiredSize)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppBroadcastServices) Get_State() *IAppBroadcastState {
	var _result *IAppBroadcastState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 02B692A4-5919-4A9E-8D5E-C9BB0DD3377A
var IID_IAppBroadcastSignInStateChangedEventArgs = syscall.GUID{0x02B692A4, 0x5919, 0x4A9E,
	[8]byte{0x8D, 0x5E, 0xC9, 0xBB, 0x0D, 0xD3, 0x37, 0x7A}}

type IAppBroadcastSignInStateChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_SignInState() AppBroadcastSignInState
	Get_Result() AppBroadcastSignInResult
}

type IAppBroadcastSignInStateChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_SignInState uintptr
	Get_Result      uintptr
}

type IAppBroadcastSignInStateChangedEventArgs struct {
	win32.IInspectable
}

func (this *IAppBroadcastSignInStateChangedEventArgs) Vtbl() *IAppBroadcastSignInStateChangedEventArgsVtbl {
	return (*IAppBroadcastSignInStateChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastSignInStateChangedEventArgs) Get_SignInState() AppBroadcastSignInState {
	var _result AppBroadcastSignInState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SignInState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastSignInStateChangedEventArgs) Get_Result() AppBroadcastSignInResult {
	var _result AppBroadcastSignInResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Result, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// EE08056D-8099-4DDD-922E-C56DAC58ABFB
var IID_IAppBroadcastState = syscall.GUID{0xEE08056D, 0x8099, 0x4DDD,
	[8]byte{0x92, 0x2E, 0xC5, 0x6D, 0xAC, 0x58, 0xAB, 0xFB}}

type IAppBroadcastStateInterface interface {
	win32.IInspectableInterface
	Get_IsCaptureTargetRunning() bool
	Get_ViewerCount() uint32
	Get_ShouldCaptureMicrophone() bool
	Put_ShouldCaptureMicrophone(value bool)
	RestartMicrophoneCapture()
	Get_ShouldCaptureCamera() bool
	Put_ShouldCaptureCamera(value bool)
	RestartCameraCapture()
	Get_EncodedVideoSize() Size
	Get_MicrophoneCaptureState() AppBroadcastMicrophoneCaptureState
	Get_MicrophoneCaptureError() uint32
	Get_CameraCaptureState() AppBroadcastCameraCaptureState
	Get_CameraCaptureError() uint32
	Get_StreamState() AppBroadcastStreamState
	Get_PlugInState() AppBroadcastPlugInState
	Get_OAuthRequestUri() *IUriRuntimeClass
	Get_OAuthCallbackUri() *IUriRuntimeClass
	Get_AuthenticationResult() unsafe.Pointer
	Put_AuthenticationResult(value unsafe.Pointer)
	Put_SignInState(value AppBroadcastSignInState)
	Get_SignInState() AppBroadcastSignInState
	Get_TerminationReason() AppBroadcastTerminationReason
	Get_TerminationReasonPlugInSpecific() uint32
	Add_ViewerCountChanged(value TypedEventHandler[*IAppBroadcastState, *IAppBroadcastViewerCountChangedEventArgs]) EventRegistrationToken
	Remove_ViewerCountChanged(token EventRegistrationToken)
	Add_MicrophoneCaptureStateChanged(value TypedEventHandler[*IAppBroadcastState, *IAppBroadcastMicrophoneCaptureStateChangedEventArgs]) EventRegistrationToken
	Remove_MicrophoneCaptureStateChanged(token EventRegistrationToken)
	Add_CameraCaptureStateChanged(value TypedEventHandler[*IAppBroadcastState, *IAppBroadcastCameraCaptureStateChangedEventArgs]) EventRegistrationToken
	Remove_CameraCaptureStateChanged(token EventRegistrationToken)
	Add_PlugInStateChanged(handler TypedEventHandler[*IAppBroadcastState, *IAppBroadcastPlugInStateChangedEventArgs]) EventRegistrationToken
	Remove_PlugInStateChanged(token EventRegistrationToken)
	Add_StreamStateChanged(handler TypedEventHandler[*IAppBroadcastState, *IAppBroadcastStreamStateChangedEventArgs]) EventRegistrationToken
	Remove_StreamStateChanged(token EventRegistrationToken)
	Add_CaptureTargetClosed(value TypedEventHandler[*IAppBroadcastState, interface{}]) EventRegistrationToken
	Remove_CaptureTargetClosed(token EventRegistrationToken)
}

type IAppBroadcastStateVtbl struct {
	win32.IInspectableVtbl
	Get_IsCaptureTargetRunning           uintptr
	Get_ViewerCount                      uintptr
	Get_ShouldCaptureMicrophone          uintptr
	Put_ShouldCaptureMicrophone          uintptr
	RestartMicrophoneCapture             uintptr
	Get_ShouldCaptureCamera              uintptr
	Put_ShouldCaptureCamera              uintptr
	RestartCameraCapture                 uintptr
	Get_EncodedVideoSize                 uintptr
	Get_MicrophoneCaptureState           uintptr
	Get_MicrophoneCaptureError           uintptr
	Get_CameraCaptureState               uintptr
	Get_CameraCaptureError               uintptr
	Get_StreamState                      uintptr
	Get_PlugInState                      uintptr
	Get_OAuthRequestUri                  uintptr
	Get_OAuthCallbackUri                 uintptr
	Get_AuthenticationResult             uintptr
	Put_AuthenticationResult             uintptr
	Put_SignInState                      uintptr
	Get_SignInState                      uintptr
	Get_TerminationReason                uintptr
	Get_TerminationReasonPlugInSpecific  uintptr
	Add_ViewerCountChanged               uintptr
	Remove_ViewerCountChanged            uintptr
	Add_MicrophoneCaptureStateChanged    uintptr
	Remove_MicrophoneCaptureStateChanged uintptr
	Add_CameraCaptureStateChanged        uintptr
	Remove_CameraCaptureStateChanged     uintptr
	Add_PlugInStateChanged               uintptr
	Remove_PlugInStateChanged            uintptr
	Add_StreamStateChanged               uintptr
	Remove_StreamStateChanged            uintptr
	Add_CaptureTargetClosed              uintptr
	Remove_CaptureTargetClosed           uintptr
}

type IAppBroadcastState struct {
	win32.IInspectable
}

func (this *IAppBroadcastState) Vtbl() *IAppBroadcastStateVtbl {
	return (*IAppBroadcastStateVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastState) Get_IsCaptureTargetRunning() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCaptureTargetRunning, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastState) Get_ViewerCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ViewerCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastState) Get_ShouldCaptureMicrophone() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ShouldCaptureMicrophone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastState) Put_ShouldCaptureMicrophone(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ShouldCaptureMicrophone, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAppBroadcastState) RestartMicrophoneCapture() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RestartMicrophoneCapture, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IAppBroadcastState) Get_ShouldCaptureCamera() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ShouldCaptureCamera, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastState) Put_ShouldCaptureCamera(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ShouldCaptureCamera, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAppBroadcastState) RestartCameraCapture() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RestartCameraCapture, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IAppBroadcastState) Get_EncodedVideoSize() Size {
	var _result Size
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EncodedVideoSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastState) Get_MicrophoneCaptureState() AppBroadcastMicrophoneCaptureState {
	var _result AppBroadcastMicrophoneCaptureState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MicrophoneCaptureState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastState) Get_MicrophoneCaptureError() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MicrophoneCaptureError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastState) Get_CameraCaptureState() AppBroadcastCameraCaptureState {
	var _result AppBroadcastCameraCaptureState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CameraCaptureState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastState) Get_CameraCaptureError() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CameraCaptureError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastState) Get_StreamState() AppBroadcastStreamState {
	var _result AppBroadcastStreamState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StreamState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastState) Get_PlugInState() AppBroadcastPlugInState {
	var _result AppBroadcastPlugInState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlugInState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastState) Get_OAuthRequestUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OAuthRequestUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppBroadcastState) Get_OAuthCallbackUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OAuthCallbackUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppBroadcastState) Get_AuthenticationResult() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AuthenticationResult, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastState) Put_AuthenticationResult(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AuthenticationResult, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppBroadcastState) Put_SignInState(value AppBroadcastSignInState) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SignInState, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppBroadcastState) Get_SignInState() AppBroadcastSignInState {
	var _result AppBroadcastSignInState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SignInState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastState) Get_TerminationReason() AppBroadcastTerminationReason {
	var _result AppBroadcastTerminationReason
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TerminationReason, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastState) Get_TerminationReasonPlugInSpecific() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TerminationReasonPlugInSpecific, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastState) Add_ViewerCountChanged(value TypedEventHandler[*IAppBroadcastState, *IAppBroadcastViewerCountChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ViewerCountChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastState) Remove_ViewerCountChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ViewerCountChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppBroadcastState) Add_MicrophoneCaptureStateChanged(value TypedEventHandler[*IAppBroadcastState, *IAppBroadcastMicrophoneCaptureStateChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_MicrophoneCaptureStateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastState) Remove_MicrophoneCaptureStateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_MicrophoneCaptureStateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppBroadcastState) Add_CameraCaptureStateChanged(value TypedEventHandler[*IAppBroadcastState, *IAppBroadcastCameraCaptureStateChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_CameraCaptureStateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastState) Remove_CameraCaptureStateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_CameraCaptureStateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppBroadcastState) Add_PlugInStateChanged(handler TypedEventHandler[*IAppBroadcastState, *IAppBroadcastPlugInStateChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PlugInStateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastState) Remove_PlugInStateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PlugInStateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppBroadcastState) Add_StreamStateChanged(handler TypedEventHandler[*IAppBroadcastState, *IAppBroadcastStreamStateChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StreamStateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastState) Remove_StreamStateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StreamStateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppBroadcastState) Add_CaptureTargetClosed(value TypedEventHandler[*IAppBroadcastState, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_CaptureTargetClosed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastState) Remove_CaptureTargetClosed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_CaptureTargetClosed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// EFAB4AC8-21BA-453F-8BB7-5E938A2E9A74
var IID_IAppBroadcastStreamAudioFrame = syscall.GUID{0xEFAB4AC8, 0x21BA, 0x453F,
	[8]byte{0x8B, 0xB7, 0x5E, 0x93, 0x8A, 0x2E, 0x9A, 0x74}}

type IAppBroadcastStreamAudioFrameInterface interface {
	win32.IInspectableInterface
	Get_AudioHeader() *IAppBroadcastStreamAudioHeader
	Get_AudioBuffer() *IBuffer
}

type IAppBroadcastStreamAudioFrameVtbl struct {
	win32.IInspectableVtbl
	Get_AudioHeader uintptr
	Get_AudioBuffer uintptr
}

type IAppBroadcastStreamAudioFrame struct {
	win32.IInspectable
}

func (this *IAppBroadcastStreamAudioFrame) Vtbl() *IAppBroadcastStreamAudioFrameVtbl {
	return (*IAppBroadcastStreamAudioFrameVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastStreamAudioFrame) Get_AudioHeader() *IAppBroadcastStreamAudioHeader {
	var _result *IAppBroadcastStreamAudioHeader
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioHeader, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppBroadcastStreamAudioFrame) Get_AudioBuffer() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioBuffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BF21A570-6B78-4216-9F07-5AFF5256F1B7
var IID_IAppBroadcastStreamAudioHeader = syscall.GUID{0xBF21A570, 0x6B78, 0x4216,
	[8]byte{0x9F, 0x07, 0x5A, 0xFF, 0x52, 0x56, 0xF1, 0xB7}}

type IAppBroadcastStreamAudioHeaderInterface interface {
	win32.IInspectableInterface
	Get_AbsoluteTimestamp() DateTime
	Get_RelativeTimestamp() TimeSpan
	Get_Duration() TimeSpan
	Get_HasDiscontinuity() bool
	Get_FrameId() uint64
}

type IAppBroadcastStreamAudioHeaderVtbl struct {
	win32.IInspectableVtbl
	Get_AbsoluteTimestamp uintptr
	Get_RelativeTimestamp uintptr
	Get_Duration          uintptr
	Get_HasDiscontinuity  uintptr
	Get_FrameId           uintptr
}

type IAppBroadcastStreamAudioHeader struct {
	win32.IInspectable
}

func (this *IAppBroadcastStreamAudioHeader) Vtbl() *IAppBroadcastStreamAudioHeaderVtbl {
	return (*IAppBroadcastStreamAudioHeaderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastStreamAudioHeader) Get_AbsoluteTimestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AbsoluteTimestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastStreamAudioHeader) Get_RelativeTimestamp() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RelativeTimestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastStreamAudioHeader) Get_Duration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Duration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastStreamAudioHeader) Get_HasDiscontinuity() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasDiscontinuity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastStreamAudioHeader) Get_FrameId() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// B338BCF9-3364-4460-B5F1-3CC2796A8AA2
var IID_IAppBroadcastStreamReader = syscall.GUID{0xB338BCF9, 0x3364, 0x4460,
	[8]byte{0xB5, 0xF1, 0x3C, 0xC2, 0x79, 0x6A, 0x8A, 0xA2}}

type IAppBroadcastStreamReaderInterface interface {
	win32.IInspectableInterface
	Get_AudioChannels() uint32
	Get_AudioSampleRate() uint32
	Get_AudioAacSequence() *IBuffer
	Get_AudioBitrate() uint32
	TryGetNextAudioFrame() *IAppBroadcastStreamAudioFrame
	Get_VideoWidth() uint32
	Get_VideoHeight() uint32
	Get_VideoBitrate() uint32
	TryGetNextVideoFrame() *IAppBroadcastStreamVideoFrame
	Add_AudioFrameArrived(value TypedEventHandler[*IAppBroadcastStreamReader, interface{}]) EventRegistrationToken
	Remove_AudioFrameArrived(token EventRegistrationToken)
	Add_VideoFrameArrived(value TypedEventHandler[*IAppBroadcastStreamReader, interface{}]) EventRegistrationToken
	Remove_VideoFrameArrived(token EventRegistrationToken)
}

type IAppBroadcastStreamReaderVtbl struct {
	win32.IInspectableVtbl
	Get_AudioChannels        uintptr
	Get_AudioSampleRate      uintptr
	Get_AudioAacSequence     uintptr
	Get_AudioBitrate         uintptr
	TryGetNextAudioFrame     uintptr
	Get_VideoWidth           uintptr
	Get_VideoHeight          uintptr
	Get_VideoBitrate         uintptr
	TryGetNextVideoFrame     uintptr
	Add_AudioFrameArrived    uintptr
	Remove_AudioFrameArrived uintptr
	Add_VideoFrameArrived    uintptr
	Remove_VideoFrameArrived uintptr
}

type IAppBroadcastStreamReader struct {
	win32.IInspectable
}

func (this *IAppBroadcastStreamReader) Vtbl() *IAppBroadcastStreamReaderVtbl {
	return (*IAppBroadcastStreamReaderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastStreamReader) Get_AudioChannels() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioChannels, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastStreamReader) Get_AudioSampleRate() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioSampleRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastStreamReader) Get_AudioAacSequence() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioAacSequence, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppBroadcastStreamReader) Get_AudioBitrate() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioBitrate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastStreamReader) TryGetNextAudioFrame() *IAppBroadcastStreamAudioFrame {
	var _result *IAppBroadcastStreamAudioFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetNextAudioFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppBroadcastStreamReader) Get_VideoWidth() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoWidth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastStreamReader) Get_VideoHeight() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoHeight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastStreamReader) Get_VideoBitrate() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoBitrate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastStreamReader) TryGetNextVideoFrame() *IAppBroadcastStreamVideoFrame {
	var _result *IAppBroadcastStreamVideoFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetNextVideoFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppBroadcastStreamReader) Add_AudioFrameArrived(value TypedEventHandler[*IAppBroadcastStreamReader, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AudioFrameArrived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastStreamReader) Remove_AudioFrameArrived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AudioFrameArrived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppBroadcastStreamReader) Add_VideoFrameArrived(value TypedEventHandler[*IAppBroadcastStreamReader, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_VideoFrameArrived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastStreamReader) Remove_VideoFrameArrived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_VideoFrameArrived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 5108A733-D008-4A89-93BE-58AED961374E
var IID_IAppBroadcastStreamStateChangedEventArgs = syscall.GUID{0x5108A733, 0xD008, 0x4A89,
	[8]byte{0x93, 0xBE, 0x58, 0xAE, 0xD9, 0x61, 0x37, 0x4E}}

type IAppBroadcastStreamStateChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_StreamState() AppBroadcastStreamState
}

type IAppBroadcastStreamStateChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_StreamState uintptr
}

type IAppBroadcastStreamStateChangedEventArgs struct {
	win32.IInspectable
}

func (this *IAppBroadcastStreamStateChangedEventArgs) Vtbl() *IAppBroadcastStreamStateChangedEventArgsVtbl {
	return (*IAppBroadcastStreamStateChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastStreamStateChangedEventArgs) Get_StreamState() AppBroadcastStreamState {
	var _result AppBroadcastStreamState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StreamState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 0F97CF2B-C9E4-4E88-8194-D814CBD585D8
var IID_IAppBroadcastStreamVideoFrame = syscall.GUID{0x0F97CF2B, 0xC9E4, 0x4E88,
	[8]byte{0x81, 0x94, 0xD8, 0x14, 0xCB, 0xD5, 0x85, 0xD8}}

type IAppBroadcastStreamVideoFrameInterface interface {
	win32.IInspectableInterface
	Get_VideoHeader() *IAppBroadcastStreamVideoHeader
	Get_VideoBuffer() *IBuffer
}

type IAppBroadcastStreamVideoFrameVtbl struct {
	win32.IInspectableVtbl
	Get_VideoHeader uintptr
	Get_VideoBuffer uintptr
}

type IAppBroadcastStreamVideoFrame struct {
	win32.IInspectable
}

func (this *IAppBroadcastStreamVideoFrame) Vtbl() *IAppBroadcastStreamVideoFrameVtbl {
	return (*IAppBroadcastStreamVideoFrameVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastStreamVideoFrame) Get_VideoHeader() *IAppBroadcastStreamVideoHeader {
	var _result *IAppBroadcastStreamVideoHeader
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoHeader, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppBroadcastStreamVideoFrame) Get_VideoBuffer() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoBuffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0B9EBECE-7E32-432D-8CA2-36BF10B9F462
var IID_IAppBroadcastStreamVideoHeader = syscall.GUID{0x0B9EBECE, 0x7E32, 0x432D,
	[8]byte{0x8C, 0xA2, 0x36, 0xBF, 0x10, 0xB9, 0xF4, 0x62}}

type IAppBroadcastStreamVideoHeaderInterface interface {
	win32.IInspectableInterface
	Get_AbsoluteTimestamp() DateTime
	Get_RelativeTimestamp() TimeSpan
	Get_Duration() TimeSpan
	Get_IsKeyFrame() bool
	Get_HasDiscontinuity() bool
	Get_FrameId() uint64
}

type IAppBroadcastStreamVideoHeaderVtbl struct {
	win32.IInspectableVtbl
	Get_AbsoluteTimestamp uintptr
	Get_RelativeTimestamp uintptr
	Get_Duration          uintptr
	Get_IsKeyFrame        uintptr
	Get_HasDiscontinuity  uintptr
	Get_FrameId           uintptr
}

type IAppBroadcastStreamVideoHeader struct {
	win32.IInspectable
}

func (this *IAppBroadcastStreamVideoHeader) Vtbl() *IAppBroadcastStreamVideoHeaderVtbl {
	return (*IAppBroadcastStreamVideoHeaderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastStreamVideoHeader) Get_AbsoluteTimestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AbsoluteTimestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastStreamVideoHeader) Get_RelativeTimestamp() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RelativeTimestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastStreamVideoHeader) Get_Duration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Duration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastStreamVideoHeader) Get_IsKeyFrame() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsKeyFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastStreamVideoHeader) Get_HasDiscontinuity() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasDiscontinuity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppBroadcastStreamVideoHeader) Get_FrameId() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// DEEBAB35-EC5E-4D8F-B1C0-5DA6E8C75638
var IID_IAppBroadcastTriggerDetails = syscall.GUID{0xDEEBAB35, 0xEC5E, 0x4D8F,
	[8]byte{0xB1, 0xC0, 0x5D, 0xA6, 0xE8, 0xC7, 0x56, 0x38}}

type IAppBroadcastTriggerDetailsInterface interface {
	win32.IInspectableInterface
	Get_BackgroundService() *IAppBroadcastBackgroundService
}

type IAppBroadcastTriggerDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_BackgroundService uintptr
}

type IAppBroadcastTriggerDetails struct {
	win32.IInspectable
}

func (this *IAppBroadcastTriggerDetails) Vtbl() *IAppBroadcastTriggerDetailsVtbl {
	return (*IAppBroadcastTriggerDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastTriggerDetails) Get_BackgroundService() *IAppBroadcastBackgroundService {
	var _result *IAppBroadcastBackgroundService
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BackgroundService, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E6E11825-5401-4ADE-8BD2-C14ECEE6807D
var IID_IAppBroadcastViewerCountChangedEventArgs = syscall.GUID{0xE6E11825, 0x5401, 0x4ADE,
	[8]byte{0x8B, 0xD2, 0xC1, 0x4E, 0xCE, 0xE6, 0x80, 0x7D}}

type IAppBroadcastViewerCountChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ViewerCount() uint32
}

type IAppBroadcastViewerCountChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ViewerCount uintptr
}

type IAppBroadcastViewerCountChangedEventArgs struct {
	win32.IInspectable
}

func (this *IAppBroadcastViewerCountChangedEventArgs) Vtbl() *IAppBroadcastViewerCountChangedEventArgsVtbl {
	return (*IAppBroadcastViewerCountChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppBroadcastViewerCountChangedEventArgs) Get_ViewerCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ViewerCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 9749D453-A29A-45ED-8F29-22D09942CFF7
var IID_IAppCapture = syscall.GUID{0x9749D453, 0xA29A, 0x45ED,
	[8]byte{0x8F, 0x29, 0x22, 0xD0, 0x99, 0x42, 0xCF, 0xF7}}

type IAppCaptureInterface interface {
	win32.IInspectableInterface
	Get_IsCapturingAudio() bool
	Get_IsCapturingVideo() bool
	Add_CapturingChanged(handler TypedEventHandler[*IAppCapture, interface{}]) EventRegistrationToken
	Remove_CapturingChanged(token EventRegistrationToken)
}

type IAppCaptureVtbl struct {
	win32.IInspectableVtbl
	Get_IsCapturingAudio    uintptr
	Get_IsCapturingVideo    uintptr
	Add_CapturingChanged    uintptr
	Remove_CapturingChanged uintptr
}

type IAppCapture struct {
	win32.IInspectable
}

func (this *IAppCapture) Vtbl() *IAppCaptureVtbl {
	return (*IAppCaptureVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppCapture) Get_IsCapturingAudio() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCapturingAudio, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCapture) Get_IsCapturingVideo() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCapturingVideo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCapture) Add_CapturingChanged(handler TypedEventHandler[*IAppCapture, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_CapturingChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCapture) Remove_CapturingChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_CapturingChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 19E8E0EF-236C-40F9-B38F-9B7DD65D1CCC
var IID_IAppCaptureAlternateShortcutKeys = syscall.GUID{0x19E8E0EF, 0x236C, 0x40F9,
	[8]byte{0xB3, 0x8F, 0x9B, 0x7D, 0xD6, 0x5D, 0x1C, 0xCC}}

type IAppCaptureAlternateShortcutKeysInterface interface {
	win32.IInspectableInterface
	Put_ToggleGameBarKey(value VirtualKey)
	Get_ToggleGameBarKey() VirtualKey
	Put_ToggleGameBarKeyModifiers(value VirtualKeyModifiers)
	Get_ToggleGameBarKeyModifiers() VirtualKeyModifiers
	Put_SaveHistoricalVideoKey(value VirtualKey)
	Get_SaveHistoricalVideoKey() VirtualKey
	Put_SaveHistoricalVideoKeyModifiers(value VirtualKeyModifiers)
	Get_SaveHistoricalVideoKeyModifiers() VirtualKeyModifiers
	Put_ToggleRecordingKey(value VirtualKey)
	Get_ToggleRecordingKey() VirtualKey
	Put_ToggleRecordingKeyModifiers(value VirtualKeyModifiers)
	Get_ToggleRecordingKeyModifiers() VirtualKeyModifiers
	Put_TakeScreenshotKey(value VirtualKey)
	Get_TakeScreenshotKey() VirtualKey
	Put_TakeScreenshotKeyModifiers(value VirtualKeyModifiers)
	Get_TakeScreenshotKeyModifiers() VirtualKeyModifiers
	Put_ToggleRecordingIndicatorKey(value VirtualKey)
	Get_ToggleRecordingIndicatorKey() VirtualKey
	Put_ToggleRecordingIndicatorKeyModifiers(value VirtualKeyModifiers)
	Get_ToggleRecordingIndicatorKeyModifiers() VirtualKeyModifiers
}

type IAppCaptureAlternateShortcutKeysVtbl struct {
	win32.IInspectableVtbl
	Put_ToggleGameBarKey                     uintptr
	Get_ToggleGameBarKey                     uintptr
	Put_ToggleGameBarKeyModifiers            uintptr
	Get_ToggleGameBarKeyModifiers            uintptr
	Put_SaveHistoricalVideoKey               uintptr
	Get_SaveHistoricalVideoKey               uintptr
	Put_SaveHistoricalVideoKeyModifiers      uintptr
	Get_SaveHistoricalVideoKeyModifiers      uintptr
	Put_ToggleRecordingKey                   uintptr
	Get_ToggleRecordingKey                   uintptr
	Put_ToggleRecordingKeyModifiers          uintptr
	Get_ToggleRecordingKeyModifiers          uintptr
	Put_TakeScreenshotKey                    uintptr
	Get_TakeScreenshotKey                    uintptr
	Put_TakeScreenshotKeyModifiers           uintptr
	Get_TakeScreenshotKeyModifiers           uintptr
	Put_ToggleRecordingIndicatorKey          uintptr
	Get_ToggleRecordingIndicatorKey          uintptr
	Put_ToggleRecordingIndicatorKeyModifiers uintptr
	Get_ToggleRecordingIndicatorKeyModifiers uintptr
}

type IAppCaptureAlternateShortcutKeys struct {
	win32.IInspectable
}

func (this *IAppCaptureAlternateShortcutKeys) Vtbl() *IAppCaptureAlternateShortcutKeysVtbl {
	return (*IAppCaptureAlternateShortcutKeysVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppCaptureAlternateShortcutKeys) Put_ToggleGameBarKey(value VirtualKey) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ToggleGameBarKey, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppCaptureAlternateShortcutKeys) Get_ToggleGameBarKey() VirtualKey {
	var _result VirtualKey
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ToggleGameBarKey, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureAlternateShortcutKeys) Put_ToggleGameBarKeyModifiers(value VirtualKeyModifiers) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ToggleGameBarKeyModifiers, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppCaptureAlternateShortcutKeys) Get_ToggleGameBarKeyModifiers() VirtualKeyModifiers {
	var _result VirtualKeyModifiers
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ToggleGameBarKeyModifiers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureAlternateShortcutKeys) Put_SaveHistoricalVideoKey(value VirtualKey) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SaveHistoricalVideoKey, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppCaptureAlternateShortcutKeys) Get_SaveHistoricalVideoKey() VirtualKey {
	var _result VirtualKey
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SaveHistoricalVideoKey, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureAlternateShortcutKeys) Put_SaveHistoricalVideoKeyModifiers(value VirtualKeyModifiers) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SaveHistoricalVideoKeyModifiers, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppCaptureAlternateShortcutKeys) Get_SaveHistoricalVideoKeyModifiers() VirtualKeyModifiers {
	var _result VirtualKeyModifiers
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SaveHistoricalVideoKeyModifiers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureAlternateShortcutKeys) Put_ToggleRecordingKey(value VirtualKey) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ToggleRecordingKey, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppCaptureAlternateShortcutKeys) Get_ToggleRecordingKey() VirtualKey {
	var _result VirtualKey
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ToggleRecordingKey, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureAlternateShortcutKeys) Put_ToggleRecordingKeyModifiers(value VirtualKeyModifiers) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ToggleRecordingKeyModifiers, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppCaptureAlternateShortcutKeys) Get_ToggleRecordingKeyModifiers() VirtualKeyModifiers {
	var _result VirtualKeyModifiers
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ToggleRecordingKeyModifiers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureAlternateShortcutKeys) Put_TakeScreenshotKey(value VirtualKey) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TakeScreenshotKey, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppCaptureAlternateShortcutKeys) Get_TakeScreenshotKey() VirtualKey {
	var _result VirtualKey
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TakeScreenshotKey, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureAlternateShortcutKeys) Put_TakeScreenshotKeyModifiers(value VirtualKeyModifiers) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TakeScreenshotKeyModifiers, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppCaptureAlternateShortcutKeys) Get_TakeScreenshotKeyModifiers() VirtualKeyModifiers {
	var _result VirtualKeyModifiers
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TakeScreenshotKeyModifiers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureAlternateShortcutKeys) Put_ToggleRecordingIndicatorKey(value VirtualKey) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ToggleRecordingIndicatorKey, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppCaptureAlternateShortcutKeys) Get_ToggleRecordingIndicatorKey() VirtualKey {
	var _result VirtualKey
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ToggleRecordingIndicatorKey, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureAlternateShortcutKeys) Put_ToggleRecordingIndicatorKeyModifiers(value VirtualKeyModifiers) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ToggleRecordingIndicatorKeyModifiers, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppCaptureAlternateShortcutKeys) Get_ToggleRecordingIndicatorKeyModifiers() VirtualKeyModifiers {
	var _result VirtualKeyModifiers
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ToggleRecordingIndicatorKeyModifiers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C3669090-DD17-47F0-95E5-CE42286CF338
var IID_IAppCaptureAlternateShortcutKeys2 = syscall.GUID{0xC3669090, 0xDD17, 0x47F0,
	[8]byte{0x95, 0xE5, 0xCE, 0x42, 0x28, 0x6C, 0xF3, 0x38}}

type IAppCaptureAlternateShortcutKeys2Interface interface {
	win32.IInspectableInterface
	Put_ToggleMicrophoneCaptureKey(value VirtualKey)
	Get_ToggleMicrophoneCaptureKey() VirtualKey
	Put_ToggleMicrophoneCaptureKeyModifiers(value VirtualKeyModifiers)
	Get_ToggleMicrophoneCaptureKeyModifiers() VirtualKeyModifiers
}

type IAppCaptureAlternateShortcutKeys2Vtbl struct {
	win32.IInspectableVtbl
	Put_ToggleMicrophoneCaptureKey          uintptr
	Get_ToggleMicrophoneCaptureKey          uintptr
	Put_ToggleMicrophoneCaptureKeyModifiers uintptr
	Get_ToggleMicrophoneCaptureKeyModifiers uintptr
}

type IAppCaptureAlternateShortcutKeys2 struct {
	win32.IInspectable
}

func (this *IAppCaptureAlternateShortcutKeys2) Vtbl() *IAppCaptureAlternateShortcutKeys2Vtbl {
	return (*IAppCaptureAlternateShortcutKeys2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppCaptureAlternateShortcutKeys2) Put_ToggleMicrophoneCaptureKey(value VirtualKey) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ToggleMicrophoneCaptureKey, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppCaptureAlternateShortcutKeys2) Get_ToggleMicrophoneCaptureKey() VirtualKey {
	var _result VirtualKey
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ToggleMicrophoneCaptureKey, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureAlternateShortcutKeys2) Put_ToggleMicrophoneCaptureKeyModifiers(value VirtualKeyModifiers) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ToggleMicrophoneCaptureKeyModifiers, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppCaptureAlternateShortcutKeys2) Get_ToggleMicrophoneCaptureKeyModifiers() VirtualKeyModifiers {
	var _result VirtualKeyModifiers
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ToggleMicrophoneCaptureKeyModifiers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7B81448C-418E-469C-A49A-45B597C826B6
var IID_IAppCaptureAlternateShortcutKeys3 = syscall.GUID{0x7B81448C, 0x418E, 0x469C,
	[8]byte{0xA4, 0x9A, 0x45, 0xB5, 0x97, 0xC8, 0x26, 0xB6}}

type IAppCaptureAlternateShortcutKeys3Interface interface {
	win32.IInspectableInterface
	Put_ToggleCameraCaptureKey(value VirtualKey)
	Get_ToggleCameraCaptureKey() VirtualKey
	Put_ToggleCameraCaptureKeyModifiers(value VirtualKeyModifiers)
	Get_ToggleCameraCaptureKeyModifiers() VirtualKeyModifiers
	Put_ToggleBroadcastKey(value VirtualKey)
	Get_ToggleBroadcastKey() VirtualKey
	Put_ToggleBroadcastKeyModifiers(value VirtualKeyModifiers)
	Get_ToggleBroadcastKeyModifiers() VirtualKeyModifiers
}

type IAppCaptureAlternateShortcutKeys3Vtbl struct {
	win32.IInspectableVtbl
	Put_ToggleCameraCaptureKey          uintptr
	Get_ToggleCameraCaptureKey          uintptr
	Put_ToggleCameraCaptureKeyModifiers uintptr
	Get_ToggleCameraCaptureKeyModifiers uintptr
	Put_ToggleBroadcastKey              uintptr
	Get_ToggleBroadcastKey              uintptr
	Put_ToggleBroadcastKeyModifiers     uintptr
	Get_ToggleBroadcastKeyModifiers     uintptr
}

type IAppCaptureAlternateShortcutKeys3 struct {
	win32.IInspectable
}

func (this *IAppCaptureAlternateShortcutKeys3) Vtbl() *IAppCaptureAlternateShortcutKeys3Vtbl {
	return (*IAppCaptureAlternateShortcutKeys3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppCaptureAlternateShortcutKeys3) Put_ToggleCameraCaptureKey(value VirtualKey) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ToggleCameraCaptureKey, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppCaptureAlternateShortcutKeys3) Get_ToggleCameraCaptureKey() VirtualKey {
	var _result VirtualKey
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ToggleCameraCaptureKey, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureAlternateShortcutKeys3) Put_ToggleCameraCaptureKeyModifiers(value VirtualKeyModifiers) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ToggleCameraCaptureKeyModifiers, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppCaptureAlternateShortcutKeys3) Get_ToggleCameraCaptureKeyModifiers() VirtualKeyModifiers {
	var _result VirtualKeyModifiers
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ToggleCameraCaptureKeyModifiers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureAlternateShortcutKeys3) Put_ToggleBroadcastKey(value VirtualKey) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ToggleBroadcastKey, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppCaptureAlternateShortcutKeys3) Get_ToggleBroadcastKey() VirtualKey {
	var _result VirtualKey
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ToggleBroadcastKey, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureAlternateShortcutKeys3) Put_ToggleBroadcastKeyModifiers(value VirtualKeyModifiers) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ToggleBroadcastKeyModifiers, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppCaptureAlternateShortcutKeys3) Get_ToggleBroadcastKeyModifiers() VirtualKeyModifiers {
	var _result VirtualKeyModifiers
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ToggleBroadcastKeyModifiers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C1F5563B-FFA1-44C9-975F-27FBEB553B35
var IID_IAppCaptureDurationGeneratedEventArgs = syscall.GUID{0xC1F5563B, 0xFFA1, 0x44C9,
	[8]byte{0x97, 0x5F, 0x27, 0xFB, 0xEB, 0x55, 0x3B, 0x35}}

type IAppCaptureDurationGeneratedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Duration() TimeSpan
}

type IAppCaptureDurationGeneratedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Duration uintptr
}

type IAppCaptureDurationGeneratedEventArgs struct {
	win32.IInspectable
}

func (this *IAppCaptureDurationGeneratedEventArgs) Vtbl() *IAppCaptureDurationGeneratedEventArgsVtbl {
	return (*IAppCaptureDurationGeneratedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppCaptureDurationGeneratedEventArgs) Get_Duration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Duration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4189FBF4-465E-45BF-907F-165B3FB23758
var IID_IAppCaptureFileGeneratedEventArgs = syscall.GUID{0x4189FBF4, 0x465E, 0x45BF,
	[8]byte{0x90, 0x7F, 0x16, 0x5B, 0x3F, 0xB2, 0x37, 0x58}}

type IAppCaptureFileGeneratedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_File() *IStorageFile
}

type IAppCaptureFileGeneratedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_File uintptr
}

type IAppCaptureFileGeneratedEventArgs struct {
	win32.IInspectable
}

func (this *IAppCaptureFileGeneratedEventArgs) Vtbl() *IAppCaptureFileGeneratedEventArgsVtbl {
	return (*IAppCaptureFileGeneratedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppCaptureFileGeneratedEventArgs) Get_File() *IStorageFile {
	var _result *IStorageFile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_File, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7D9E3EA7-6282-4735-8D4E-AA45F90F6723
var IID_IAppCaptureManagerStatics = syscall.GUID{0x7D9E3EA7, 0x6282, 0x4735,
	[8]byte{0x8D, 0x4E, 0xAA, 0x45, 0xF9, 0x0F, 0x67, 0x23}}

type IAppCaptureManagerStaticsInterface interface {
	win32.IInspectableInterface
	GetCurrentSettings() *IAppCaptureSettings
	ApplySettings(appCaptureSettings *IAppCaptureSettings)
}

type IAppCaptureManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetCurrentSettings uintptr
	ApplySettings      uintptr
}

type IAppCaptureManagerStatics struct {
	win32.IInspectable
}

func (this *IAppCaptureManagerStatics) Vtbl() *IAppCaptureManagerStaticsVtbl {
	return (*IAppCaptureManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppCaptureManagerStatics) GetCurrentSettings() *IAppCaptureSettings {
	var _result *IAppCaptureSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentSettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppCaptureManagerStatics) ApplySettings(appCaptureSettings *IAppCaptureSettings) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ApplySettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(appCaptureSettings)))
	_ = _hr
}

// E0CE4877-9AAF-46B4-AD31-6A60B441C780
var IID_IAppCaptureMetadataWriter = syscall.GUID{0xE0CE4877, 0x9AAF, 0x46B4,
	[8]byte{0xAD, 0x31, 0x6A, 0x60, 0xB4, 0x41, 0xC7, 0x80}}

type IAppCaptureMetadataWriterInterface interface {
	win32.IInspectableInterface
	AddStringEvent(name string, value string, priority AppCaptureMetadataPriority)
	AddInt32Event(name string, value int32, priority AppCaptureMetadataPriority)
	AddDoubleEvent(name string, value float64, priority AppCaptureMetadataPriority)
	StartStringState(name string, value string, priority AppCaptureMetadataPriority)
	StartInt32State(name string, value int32, priority AppCaptureMetadataPriority)
	StartDoubleState(name string, value float64, priority AppCaptureMetadataPriority)
	StopState(name string)
	StopAllStates()
	Get_RemainingStorageBytesAvailable() uint64
	Add_MetadataPurged(handler TypedEventHandler[*IAppCaptureMetadataWriter, interface{}]) EventRegistrationToken
	Remove_MetadataPurged(token EventRegistrationToken)
}

type IAppCaptureMetadataWriterVtbl struct {
	win32.IInspectableVtbl
	AddStringEvent                     uintptr
	AddInt32Event                      uintptr
	AddDoubleEvent                     uintptr
	StartStringState                   uintptr
	StartInt32State                    uintptr
	StartDoubleState                   uintptr
	StopState                          uintptr
	StopAllStates                      uintptr
	Get_RemainingStorageBytesAvailable uintptr
	Add_MetadataPurged                 uintptr
	Remove_MetadataPurged              uintptr
}

type IAppCaptureMetadataWriter struct {
	win32.IInspectable
}

func (this *IAppCaptureMetadataWriter) Vtbl() *IAppCaptureMetadataWriterVtbl {
	return (*IAppCaptureMetadataWriterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppCaptureMetadataWriter) AddStringEvent(name string, value string, priority AppCaptureMetadataPriority) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddStringEvent, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, NewHStr(value).Ptr, uintptr(priority))
	_ = _hr
}

func (this *IAppCaptureMetadataWriter) AddInt32Event(name string, value int32, priority AppCaptureMetadataPriority) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddInt32Event, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value), uintptr(priority))
	_ = _hr
}

func (this *IAppCaptureMetadataWriter) AddDoubleEvent(name string, value float64, priority AppCaptureMetadataPriority) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddDoubleEvent, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value), uintptr(priority))
	_ = _hr
}

func (this *IAppCaptureMetadataWriter) StartStringState(name string, value string, priority AppCaptureMetadataPriority) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartStringState, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, NewHStr(value).Ptr, uintptr(priority))
	_ = _hr
}

func (this *IAppCaptureMetadataWriter) StartInt32State(name string, value int32, priority AppCaptureMetadataPriority) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartInt32State, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value), uintptr(priority))
	_ = _hr
}

func (this *IAppCaptureMetadataWriter) StartDoubleState(name string, value float64, priority AppCaptureMetadataPriority) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartDoubleState, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(value), uintptr(priority))
	_ = _hr
}

func (this *IAppCaptureMetadataWriter) StopState(name string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopState, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr)
	_ = _hr
}

func (this *IAppCaptureMetadataWriter) StopAllStates() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopAllStates, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IAppCaptureMetadataWriter) Get_RemainingStorageBytesAvailable() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RemainingStorageBytesAvailable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureMetadataWriter) Add_MetadataPurged(handler TypedEventHandler[*IAppCaptureMetadataWriter, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_MetadataPurged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureMetadataWriter) Remove_MetadataPurged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_MetadataPurged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 324D249E-45BC-4C35-BC35-E469FC7A69E0
var IID_IAppCaptureMicrophoneCaptureStateChangedEventArgs = syscall.GUID{0x324D249E, 0x45BC, 0x4C35,
	[8]byte{0xBC, 0x35, 0xE4, 0x69, 0xFC, 0x7A, 0x69, 0xE0}}

type IAppCaptureMicrophoneCaptureStateChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_State() AppCaptureMicrophoneCaptureState
	Get_ErrorCode() uint32
}

type IAppCaptureMicrophoneCaptureStateChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_State     uintptr
	Get_ErrorCode uintptr
}

type IAppCaptureMicrophoneCaptureStateChangedEventArgs struct {
	win32.IInspectable
}

func (this *IAppCaptureMicrophoneCaptureStateChangedEventArgs) Vtbl() *IAppCaptureMicrophoneCaptureStateChangedEventArgsVtbl {
	return (*IAppCaptureMicrophoneCaptureStateChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppCaptureMicrophoneCaptureStateChangedEventArgs) Get_State() AppCaptureMicrophoneCaptureState {
	var _result AppCaptureMicrophoneCaptureState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureMicrophoneCaptureStateChangedEventArgs) Get_ErrorCode() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ErrorCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C66020A9-1538-495C-9BBB-2BA870EC5861
var IID_IAppCaptureRecordOperation = syscall.GUID{0xC66020A9, 0x1538, 0x495C,
	[8]byte{0x9B, 0xBB, 0x2B, 0xA8, 0x70, 0xEC, 0x58, 0x61}}

type IAppCaptureRecordOperationInterface interface {
	win32.IInspectableInterface
	StopRecording()
	Get_State() AppCaptureRecordingState
	Get_ErrorCode() *IReference[uint32]
	Get_Duration() *IReference[TimeSpan]
	Get_File() *IStorageFile
	Get_IsFileTruncated() *IReference[bool]
	Add_StateChanged(value TypedEventHandler[*IAppCaptureRecordOperation, *IAppCaptureRecordingStateChangedEventArgs]) EventRegistrationToken
	Remove_StateChanged(token EventRegistrationToken)
	Add_DurationGenerated(value TypedEventHandler[*IAppCaptureRecordOperation, *IAppCaptureDurationGeneratedEventArgs]) EventRegistrationToken
	Remove_DurationGenerated(token EventRegistrationToken)
	Add_FileGenerated(value TypedEventHandler[*IAppCaptureRecordOperation, *IAppCaptureFileGeneratedEventArgs]) EventRegistrationToken
	Remove_FileGenerated(token EventRegistrationToken)
}

type IAppCaptureRecordOperationVtbl struct {
	win32.IInspectableVtbl
	StopRecording            uintptr
	Get_State                uintptr
	Get_ErrorCode            uintptr
	Get_Duration             uintptr
	Get_File                 uintptr
	Get_IsFileTruncated      uintptr
	Add_StateChanged         uintptr
	Remove_StateChanged      uintptr
	Add_DurationGenerated    uintptr
	Remove_DurationGenerated uintptr
	Add_FileGenerated        uintptr
	Remove_FileGenerated     uintptr
}

type IAppCaptureRecordOperation struct {
	win32.IInspectable
}

func (this *IAppCaptureRecordOperation) Vtbl() *IAppCaptureRecordOperationVtbl {
	return (*IAppCaptureRecordOperationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppCaptureRecordOperation) StopRecording() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopRecording, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IAppCaptureRecordOperation) Get_State() AppCaptureRecordingState {
	var _result AppCaptureRecordingState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureRecordOperation) Get_ErrorCode() *IReference[uint32] {
	var _result *IReference[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ErrorCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppCaptureRecordOperation) Get_Duration() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Duration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppCaptureRecordOperation) Get_File() *IStorageFile {
	var _result *IStorageFile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_File, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppCaptureRecordOperation) Get_IsFileTruncated() *IReference[bool] {
	var _result *IReference[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsFileTruncated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppCaptureRecordOperation) Add_StateChanged(value TypedEventHandler[*IAppCaptureRecordOperation, *IAppCaptureRecordingStateChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureRecordOperation) Remove_StateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppCaptureRecordOperation) Add_DurationGenerated(value TypedEventHandler[*IAppCaptureRecordOperation, *IAppCaptureDurationGeneratedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DurationGenerated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureRecordOperation) Remove_DurationGenerated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DurationGenerated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppCaptureRecordOperation) Add_FileGenerated(value TypedEventHandler[*IAppCaptureRecordOperation, *IAppCaptureFileGeneratedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_FileGenerated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureRecordOperation) Remove_FileGenerated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_FileGenerated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 24FC8712-E305-490D-B415-6B1C9049736B
var IID_IAppCaptureRecordingStateChangedEventArgs = syscall.GUID{0x24FC8712, 0xE305, 0x490D,
	[8]byte{0xB4, 0x15, 0x6B, 0x1C, 0x90, 0x49, 0x73, 0x6B}}

type IAppCaptureRecordingStateChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_State() AppCaptureRecordingState
	Get_ErrorCode() uint32
}

type IAppCaptureRecordingStateChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_State     uintptr
	Get_ErrorCode uintptr
}

type IAppCaptureRecordingStateChangedEventArgs struct {
	win32.IInspectable
}

func (this *IAppCaptureRecordingStateChangedEventArgs) Vtbl() *IAppCaptureRecordingStateChangedEventArgsVtbl {
	return (*IAppCaptureRecordingStateChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppCaptureRecordingStateChangedEventArgs) Get_State() AppCaptureRecordingState {
	var _result AppCaptureRecordingState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureRecordingStateChangedEventArgs) Get_ErrorCode() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ErrorCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 44FEC0B5-34F5-4F18-AE8C-B9123ABBFC0D
var IID_IAppCaptureServices = syscall.GUID{0x44FEC0B5, 0x34F5, 0x4F18,
	[8]byte{0xAE, 0x8C, 0xB9, 0x12, 0x3A, 0xBB, 0xFC, 0x0D}}

type IAppCaptureServicesInterface interface {
	win32.IInspectableInterface
	Record() *IAppCaptureRecordOperation
	RecordTimeSpan(startTime DateTime, duration TimeSpan) *IAppCaptureRecordOperation
	Get_CanCapture() bool
	Get_State() *IAppCaptureState
}

type IAppCaptureServicesVtbl struct {
	win32.IInspectableVtbl
	Record         uintptr
	RecordTimeSpan uintptr
	Get_CanCapture uintptr
	Get_State      uintptr
}

type IAppCaptureServices struct {
	win32.IInspectable
}

func (this *IAppCaptureServices) Vtbl() *IAppCaptureServicesVtbl {
	return (*IAppCaptureServicesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppCaptureServices) Record() *IAppCaptureRecordOperation {
	var _result *IAppCaptureRecordOperation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Record, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppCaptureServices) RecordTimeSpan(startTime DateTime, duration TimeSpan) *IAppCaptureRecordOperation {
	var _result *IAppCaptureRecordOperation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RecordTimeSpan, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&startTime)), *(*uintptr)(unsafe.Pointer(&duration)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppCaptureServices) Get_CanCapture() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanCapture, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureServices) Get_State() *IAppCaptureState {
	var _result *IAppCaptureState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 14683A86-8807-48D3-883A-970EE4532A39
var IID_IAppCaptureSettings = syscall.GUID{0x14683A86, 0x8807, 0x48D3,
	[8]byte{0x88, 0x3A, 0x97, 0x0E, 0xE4, 0x53, 0x2A, 0x39}}

type IAppCaptureSettingsInterface interface {
	win32.IInspectableInterface
	Put_AppCaptureDestinationFolder(value *IStorageFolder)
	Get_AppCaptureDestinationFolder() *IStorageFolder
	Put_AudioEncodingBitrate(value uint32)
	Get_AudioEncodingBitrate() uint32
	Put_IsAudioCaptureEnabled(value bool)
	Get_IsAudioCaptureEnabled() bool
	Put_CustomVideoEncodingBitrate(value uint32)
	Get_CustomVideoEncodingBitrate() uint32
	Put_CustomVideoEncodingHeight(value uint32)
	Get_CustomVideoEncodingHeight() uint32
	Put_CustomVideoEncodingWidth(value uint32)
	Get_CustomVideoEncodingWidth() uint32
	Put_HistoricalBufferLength(value uint32)
	Get_HistoricalBufferLength() uint32
	Put_HistoricalBufferLengthUnit(value AppCaptureHistoricalBufferLengthUnit)
	Get_HistoricalBufferLengthUnit() AppCaptureHistoricalBufferLengthUnit
	Put_IsHistoricalCaptureEnabled(value bool)
	Get_IsHistoricalCaptureEnabled() bool
	Put_IsHistoricalCaptureOnBatteryAllowed(value bool)
	Get_IsHistoricalCaptureOnBatteryAllowed() bool
	Put_IsHistoricalCaptureOnWirelessDisplayAllowed(value bool)
	Get_IsHistoricalCaptureOnWirelessDisplayAllowed() bool
	Put_MaximumRecordLength(value TimeSpan)
	Get_MaximumRecordLength() TimeSpan
	Put_ScreenshotDestinationFolder(value *IStorageFolder)
	Get_ScreenshotDestinationFolder() *IStorageFolder
	Put_VideoEncodingBitrateMode(value AppCaptureVideoEncodingBitrateMode)
	Get_VideoEncodingBitrateMode() AppCaptureVideoEncodingBitrateMode
	Put_VideoEncodingResolutionMode(value AppCaptureVideoEncodingResolutionMode)
	Get_VideoEncodingResolutionMode() AppCaptureVideoEncodingResolutionMode
	Put_IsAppCaptureEnabled(value bool)
	Get_IsAppCaptureEnabled() bool
	Get_IsCpuConstrained() bool
	Get_IsDisabledByPolicy() bool
	Get_IsMemoryConstrained() bool
	Get_HasHardwareEncoder() bool
}

type IAppCaptureSettingsVtbl struct {
	win32.IInspectableVtbl
	Put_AppCaptureDestinationFolder                 uintptr
	Get_AppCaptureDestinationFolder                 uintptr
	Put_AudioEncodingBitrate                        uintptr
	Get_AudioEncodingBitrate                        uintptr
	Put_IsAudioCaptureEnabled                       uintptr
	Get_IsAudioCaptureEnabled                       uintptr
	Put_CustomVideoEncodingBitrate                  uintptr
	Get_CustomVideoEncodingBitrate                  uintptr
	Put_CustomVideoEncodingHeight                   uintptr
	Get_CustomVideoEncodingHeight                   uintptr
	Put_CustomVideoEncodingWidth                    uintptr
	Get_CustomVideoEncodingWidth                    uintptr
	Put_HistoricalBufferLength                      uintptr
	Get_HistoricalBufferLength                      uintptr
	Put_HistoricalBufferLengthUnit                  uintptr
	Get_HistoricalBufferLengthUnit                  uintptr
	Put_IsHistoricalCaptureEnabled                  uintptr
	Get_IsHistoricalCaptureEnabled                  uintptr
	Put_IsHistoricalCaptureOnBatteryAllowed         uintptr
	Get_IsHistoricalCaptureOnBatteryAllowed         uintptr
	Put_IsHistoricalCaptureOnWirelessDisplayAllowed uintptr
	Get_IsHistoricalCaptureOnWirelessDisplayAllowed uintptr
	Put_MaximumRecordLength                         uintptr
	Get_MaximumRecordLength                         uintptr
	Put_ScreenshotDestinationFolder                 uintptr
	Get_ScreenshotDestinationFolder                 uintptr
	Put_VideoEncodingBitrateMode                    uintptr
	Get_VideoEncodingBitrateMode                    uintptr
	Put_VideoEncodingResolutionMode                 uintptr
	Get_VideoEncodingResolutionMode                 uintptr
	Put_IsAppCaptureEnabled                         uintptr
	Get_IsAppCaptureEnabled                         uintptr
	Get_IsCpuConstrained                            uintptr
	Get_IsDisabledByPolicy                          uintptr
	Get_IsMemoryConstrained                         uintptr
	Get_HasHardwareEncoder                          uintptr
}

type IAppCaptureSettings struct {
	win32.IInspectable
}

func (this *IAppCaptureSettings) Vtbl() *IAppCaptureSettingsVtbl {
	return (*IAppCaptureSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppCaptureSettings) Put_AppCaptureDestinationFolder(value *IStorageFolder) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AppCaptureDestinationFolder, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAppCaptureSettings) Get_AppCaptureDestinationFolder() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppCaptureDestinationFolder, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppCaptureSettings) Put_AudioEncodingBitrate(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AudioEncodingBitrate, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppCaptureSettings) Get_AudioEncodingBitrate() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioEncodingBitrate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureSettings) Put_IsAudioCaptureEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsAudioCaptureEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAppCaptureSettings) Get_IsAudioCaptureEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsAudioCaptureEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureSettings) Put_CustomVideoEncodingBitrate(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CustomVideoEncodingBitrate, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppCaptureSettings) Get_CustomVideoEncodingBitrate() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CustomVideoEncodingBitrate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureSettings) Put_CustomVideoEncodingHeight(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CustomVideoEncodingHeight, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppCaptureSettings) Get_CustomVideoEncodingHeight() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CustomVideoEncodingHeight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureSettings) Put_CustomVideoEncodingWidth(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CustomVideoEncodingWidth, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppCaptureSettings) Get_CustomVideoEncodingWidth() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CustomVideoEncodingWidth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureSettings) Put_HistoricalBufferLength(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_HistoricalBufferLength, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppCaptureSettings) Get_HistoricalBufferLength() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HistoricalBufferLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureSettings) Put_HistoricalBufferLengthUnit(value AppCaptureHistoricalBufferLengthUnit) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_HistoricalBufferLengthUnit, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppCaptureSettings) Get_HistoricalBufferLengthUnit() AppCaptureHistoricalBufferLengthUnit {
	var _result AppCaptureHistoricalBufferLengthUnit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HistoricalBufferLengthUnit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureSettings) Put_IsHistoricalCaptureEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsHistoricalCaptureEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAppCaptureSettings) Get_IsHistoricalCaptureEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsHistoricalCaptureEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureSettings) Put_IsHistoricalCaptureOnBatteryAllowed(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsHistoricalCaptureOnBatteryAllowed, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAppCaptureSettings) Get_IsHistoricalCaptureOnBatteryAllowed() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsHistoricalCaptureOnBatteryAllowed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureSettings) Put_IsHistoricalCaptureOnWirelessDisplayAllowed(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsHistoricalCaptureOnWirelessDisplayAllowed, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAppCaptureSettings) Get_IsHistoricalCaptureOnWirelessDisplayAllowed() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsHistoricalCaptureOnWirelessDisplayAllowed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureSettings) Put_MaximumRecordLength(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaximumRecordLength, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IAppCaptureSettings) Get_MaximumRecordLength() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaximumRecordLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureSettings) Put_ScreenshotDestinationFolder(value *IStorageFolder) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ScreenshotDestinationFolder, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAppCaptureSettings) Get_ScreenshotDestinationFolder() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScreenshotDestinationFolder, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppCaptureSettings) Put_VideoEncodingBitrateMode(value AppCaptureVideoEncodingBitrateMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_VideoEncodingBitrateMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppCaptureSettings) Get_VideoEncodingBitrateMode() AppCaptureVideoEncodingBitrateMode {
	var _result AppCaptureVideoEncodingBitrateMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoEncodingBitrateMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureSettings) Put_VideoEncodingResolutionMode(value AppCaptureVideoEncodingResolutionMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_VideoEncodingResolutionMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppCaptureSettings) Get_VideoEncodingResolutionMode() AppCaptureVideoEncodingResolutionMode {
	var _result AppCaptureVideoEncodingResolutionMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoEncodingResolutionMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureSettings) Put_IsAppCaptureEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsAppCaptureEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAppCaptureSettings) Get_IsAppCaptureEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsAppCaptureEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureSettings) Get_IsCpuConstrained() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCpuConstrained, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureSettings) Get_IsDisabledByPolicy() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDisabledByPolicy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureSettings) Get_IsMemoryConstrained() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsMemoryConstrained, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureSettings) Get_HasHardwareEncoder() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasHardwareEncoder, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// FCB8CEE7-E26B-476F-9B1A-EC342D2A8FDE
var IID_IAppCaptureSettings2 = syscall.GUID{0xFCB8CEE7, 0xE26B, 0x476F,
	[8]byte{0x9B, 0x1A, 0xEC, 0x34, 0x2D, 0x2A, 0x8F, 0xDE}}

type IAppCaptureSettings2Interface interface {
	win32.IInspectableInterface
	Get_IsGpuConstrained() bool
	Get_AlternateShortcutKeys() *IAppCaptureAlternateShortcutKeys
}

type IAppCaptureSettings2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsGpuConstrained      uintptr
	Get_AlternateShortcutKeys uintptr
}

type IAppCaptureSettings2 struct {
	win32.IInspectable
}

func (this *IAppCaptureSettings2) Vtbl() *IAppCaptureSettings2Vtbl {
	return (*IAppCaptureSettings2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppCaptureSettings2) Get_IsGpuConstrained() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsGpuConstrained, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureSettings2) Get_AlternateShortcutKeys() *IAppCaptureAlternateShortcutKeys {
	var _result *IAppCaptureAlternateShortcutKeys
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlternateShortcutKeys, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A93502FE-88C2-42D6-AAAA-40FEFFD75AEC
var IID_IAppCaptureSettings3 = syscall.GUID{0xA93502FE, 0x88C2, 0x42D6,
	[8]byte{0xAA, 0xAA, 0x40, 0xFE, 0xFF, 0xD7, 0x5A, 0xEC}}

type IAppCaptureSettings3Interface interface {
	win32.IInspectableInterface
	Put_IsMicrophoneCaptureEnabled(value bool)
	Get_IsMicrophoneCaptureEnabled() bool
}

type IAppCaptureSettings3Vtbl struct {
	win32.IInspectableVtbl
	Put_IsMicrophoneCaptureEnabled uintptr
	Get_IsMicrophoneCaptureEnabled uintptr
}

type IAppCaptureSettings3 struct {
	win32.IInspectable
}

func (this *IAppCaptureSettings3) Vtbl() *IAppCaptureSettings3Vtbl {
	return (*IAppCaptureSettings3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppCaptureSettings3) Put_IsMicrophoneCaptureEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsMicrophoneCaptureEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAppCaptureSettings3) Get_IsMicrophoneCaptureEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsMicrophoneCaptureEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 07C2774C-1A81-482F-A244-049D95F25B0B
var IID_IAppCaptureSettings4 = syscall.GUID{0x07C2774C, 0x1A81, 0x482F,
	[8]byte{0xA2, 0x44, 0x04, 0x9D, 0x95, 0xF2, 0x5B, 0x0B}}

type IAppCaptureSettings4Interface interface {
	win32.IInspectableInterface
	Put_IsMicrophoneCaptureEnabledByDefault(value bool)
	Get_IsMicrophoneCaptureEnabledByDefault() bool
	Put_SystemAudioGain(value float64)
	Get_SystemAudioGain() float64
	Put_MicrophoneGain(value float64)
	Get_MicrophoneGain() float64
	Put_VideoEncodingFrameRateMode(value AppCaptureVideoEncodingFrameRateMode)
	Get_VideoEncodingFrameRateMode() AppCaptureVideoEncodingFrameRateMode
}

type IAppCaptureSettings4Vtbl struct {
	win32.IInspectableVtbl
	Put_IsMicrophoneCaptureEnabledByDefault uintptr
	Get_IsMicrophoneCaptureEnabledByDefault uintptr
	Put_SystemAudioGain                     uintptr
	Get_SystemAudioGain                     uintptr
	Put_MicrophoneGain                      uintptr
	Get_MicrophoneGain                      uintptr
	Put_VideoEncodingFrameRateMode          uintptr
	Get_VideoEncodingFrameRateMode          uintptr
}

type IAppCaptureSettings4 struct {
	win32.IInspectable
}

func (this *IAppCaptureSettings4) Vtbl() *IAppCaptureSettings4Vtbl {
	return (*IAppCaptureSettings4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppCaptureSettings4) Put_IsMicrophoneCaptureEnabledByDefault(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsMicrophoneCaptureEnabledByDefault, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAppCaptureSettings4) Get_IsMicrophoneCaptureEnabledByDefault() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsMicrophoneCaptureEnabledByDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureSettings4) Put_SystemAudioGain(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SystemAudioGain, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppCaptureSettings4) Get_SystemAudioGain() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SystemAudioGain, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureSettings4) Put_MicrophoneGain(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MicrophoneGain, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppCaptureSettings4) Get_MicrophoneGain() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MicrophoneGain, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureSettings4) Put_VideoEncodingFrameRateMode(value AppCaptureVideoEncodingFrameRateMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_VideoEncodingFrameRateMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAppCaptureSettings4) Get_VideoEncodingFrameRateMode() AppCaptureVideoEncodingFrameRateMode {
	var _result AppCaptureVideoEncodingFrameRateMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoEncodingFrameRateMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 18894522-B0E8-4BA0-8F13-3EAA5FA4013B
var IID_IAppCaptureSettings5 = syscall.GUID{0x18894522, 0xB0E8, 0x4BA0,
	[8]byte{0x8F, 0x13, 0x3E, 0xAA, 0x5F, 0xA4, 0x01, 0x3B}}

type IAppCaptureSettings5Interface interface {
	win32.IInspectableInterface
	Put_IsEchoCancellationEnabled(value bool)
	Get_IsEchoCancellationEnabled() bool
	Put_IsCursorImageCaptureEnabled(value bool)
	Get_IsCursorImageCaptureEnabled() bool
}

type IAppCaptureSettings5Vtbl struct {
	win32.IInspectableVtbl
	Put_IsEchoCancellationEnabled   uintptr
	Get_IsEchoCancellationEnabled   uintptr
	Put_IsCursorImageCaptureEnabled uintptr
	Get_IsCursorImageCaptureEnabled uintptr
}

type IAppCaptureSettings5 struct {
	win32.IInspectable
}

func (this *IAppCaptureSettings5) Vtbl() *IAppCaptureSettings5Vtbl {
	return (*IAppCaptureSettings5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppCaptureSettings5) Put_IsEchoCancellationEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsEchoCancellationEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAppCaptureSettings5) Get_IsEchoCancellationEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEchoCancellationEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureSettings5) Put_IsCursorImageCaptureEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsCursorImageCaptureEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAppCaptureSettings5) Get_IsCursorImageCaptureEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCursorImageCaptureEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 73134372-D4EB-44CE-9538-465F506AC4EA
var IID_IAppCaptureState = syscall.GUID{0x73134372, 0xD4EB, 0x44CE,
	[8]byte{0x95, 0x38, 0x46, 0x5F, 0x50, 0x6A, 0xC4, 0xEA}}

type IAppCaptureStateInterface interface {
	win32.IInspectableInterface
	Get_IsTargetRunning() bool
	Get_IsHistoricalCaptureEnabled() bool
	Get_ShouldCaptureMicrophone() bool
	Put_ShouldCaptureMicrophone(value bool)
	RestartMicrophoneCapture()
	Get_MicrophoneCaptureState() AppCaptureMicrophoneCaptureState
	Get_MicrophoneCaptureError() uint32
	Add_MicrophoneCaptureStateChanged(value TypedEventHandler[*IAppCaptureState, *IAppCaptureMicrophoneCaptureStateChangedEventArgs]) EventRegistrationToken
	Remove_MicrophoneCaptureStateChanged(token EventRegistrationToken)
	Add_CaptureTargetClosed(value TypedEventHandler[*IAppCaptureState, interface{}]) EventRegistrationToken
	Remove_CaptureTargetClosed(token EventRegistrationToken)
}

type IAppCaptureStateVtbl struct {
	win32.IInspectableVtbl
	Get_IsTargetRunning                  uintptr
	Get_IsHistoricalCaptureEnabled       uintptr
	Get_ShouldCaptureMicrophone          uintptr
	Put_ShouldCaptureMicrophone          uintptr
	RestartMicrophoneCapture             uintptr
	Get_MicrophoneCaptureState           uintptr
	Get_MicrophoneCaptureError           uintptr
	Add_MicrophoneCaptureStateChanged    uintptr
	Remove_MicrophoneCaptureStateChanged uintptr
	Add_CaptureTargetClosed              uintptr
	Remove_CaptureTargetClosed           uintptr
}

type IAppCaptureState struct {
	win32.IInspectable
}

func (this *IAppCaptureState) Vtbl() *IAppCaptureStateVtbl {
	return (*IAppCaptureStateVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppCaptureState) Get_IsTargetRunning() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsTargetRunning, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureState) Get_IsHistoricalCaptureEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsHistoricalCaptureEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureState) Get_ShouldCaptureMicrophone() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ShouldCaptureMicrophone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureState) Put_ShouldCaptureMicrophone(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ShouldCaptureMicrophone, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAppCaptureState) RestartMicrophoneCapture() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RestartMicrophoneCapture, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IAppCaptureState) Get_MicrophoneCaptureState() AppCaptureMicrophoneCaptureState {
	var _result AppCaptureMicrophoneCaptureState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MicrophoneCaptureState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureState) Get_MicrophoneCaptureError() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MicrophoneCaptureError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureState) Add_MicrophoneCaptureStateChanged(value TypedEventHandler[*IAppCaptureState, *IAppCaptureMicrophoneCaptureStateChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_MicrophoneCaptureStateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureState) Remove_MicrophoneCaptureStateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_MicrophoneCaptureStateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppCaptureState) Add_CaptureTargetClosed(value TypedEventHandler[*IAppCaptureState, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_CaptureTargetClosed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppCaptureState) Remove_CaptureTargetClosed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_CaptureTargetClosed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// F922DD6C-0A7E-4E74-8B20-9C1F902D08A1
var IID_IAppCaptureStatics = syscall.GUID{0xF922DD6C, 0x0A7E, 0x4E74,
	[8]byte{0x8B, 0x20, 0x9C, 0x1F, 0x90, 0x2D, 0x08, 0xA1}}

type IAppCaptureStaticsInterface interface {
	win32.IInspectableInterface
	GetForCurrentView() *IAppCapture
}

type IAppCaptureStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForCurrentView uintptr
}

type IAppCaptureStatics struct {
	win32.IInspectable
}

func (this *IAppCaptureStatics) Vtbl() *IAppCaptureStaticsVtbl {
	return (*IAppCaptureStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppCaptureStatics) GetForCurrentView() *IAppCapture {
	var _result *IAppCapture
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B2D881D4-836C-4DA4-AFD7-FACC041E1CF3
var IID_IAppCaptureStatics2 = syscall.GUID{0xB2D881D4, 0x836C, 0x4DA4,
	[8]byte{0xAF, 0xD7, 0xFA, 0xCC, 0x04, 0x1E, 0x1C, 0xF3}}

type IAppCaptureStatics2Interface interface {
	win32.IInspectableInterface
	SetAllowedAsync(allowed bool) *IAsyncAction
}

type IAppCaptureStatics2Vtbl struct {
	win32.IInspectableVtbl
	SetAllowedAsync uintptr
}

type IAppCaptureStatics2 struct {
	win32.IInspectable
}

func (this *IAppCaptureStatics2) Vtbl() *IAppCaptureStatics2Vtbl {
	return (*IAppCaptureStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppCaptureStatics2) SetAllowedAsync(allowed bool) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetAllowedAsync, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&allowed))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 48587540-6F93-4BB4-B8F3-E89E48948C91
var IID_ICameraCaptureUI = syscall.GUID{0x48587540, 0x6F93, 0x4BB4,
	[8]byte{0xB8, 0xF3, 0xE8, 0x9E, 0x48, 0x94, 0x8C, 0x91}}

type ICameraCaptureUIInterface interface {
	win32.IInspectableInterface
	Get_PhotoSettings() *ICameraCaptureUIPhotoCaptureSettings
	Get_VideoSettings() *ICameraCaptureUIVideoCaptureSettings
	CaptureFileAsync(mode CameraCaptureUIMode) *IAsyncOperation[*IStorageFile]
}

type ICameraCaptureUIVtbl struct {
	win32.IInspectableVtbl
	Get_PhotoSettings uintptr
	Get_VideoSettings uintptr
	CaptureFileAsync  uintptr
}

type ICameraCaptureUI struct {
	win32.IInspectable
}

func (this *ICameraCaptureUI) Vtbl() *ICameraCaptureUIVtbl {
	return (*ICameraCaptureUIVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICameraCaptureUI) Get_PhotoSettings() *ICameraCaptureUIPhotoCaptureSettings {
	var _result *ICameraCaptureUIPhotoCaptureSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhotoSettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICameraCaptureUI) Get_VideoSettings() *ICameraCaptureUIVideoCaptureSettings {
	var _result *ICameraCaptureUIVideoCaptureSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoSettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICameraCaptureUI) CaptureFileAsync(mode CameraCaptureUIMode) *IAsyncOperation[*IStorageFile] {
	var _result *IAsyncOperation[*IStorageFile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CaptureFileAsync, uintptr(unsafe.Pointer(this)), uintptr(mode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B9F5BE97-3472-46A8-8A9E-04CE42CCC97D
var IID_ICameraCaptureUIPhotoCaptureSettings = syscall.GUID{0xB9F5BE97, 0x3472, 0x46A8,
	[8]byte{0x8A, 0x9E, 0x04, 0xCE, 0x42, 0xCC, 0xC9, 0x7D}}

type ICameraCaptureUIPhotoCaptureSettingsInterface interface {
	win32.IInspectableInterface
	Get_Format() CameraCaptureUIPhotoFormat
	Put_Format(value CameraCaptureUIPhotoFormat)
	Get_MaxResolution() CameraCaptureUIMaxPhotoResolution
	Put_MaxResolution(value CameraCaptureUIMaxPhotoResolution)
	Get_CroppedSizeInPixels() Size
	Put_CroppedSizeInPixels(value Size)
	Get_CroppedAspectRatio() Size
	Put_CroppedAspectRatio(value Size)
	Get_AllowCropping() bool
	Put_AllowCropping(value bool)
}

type ICameraCaptureUIPhotoCaptureSettingsVtbl struct {
	win32.IInspectableVtbl
	Get_Format              uintptr
	Put_Format              uintptr
	Get_MaxResolution       uintptr
	Put_MaxResolution       uintptr
	Get_CroppedSizeInPixels uintptr
	Put_CroppedSizeInPixels uintptr
	Get_CroppedAspectRatio  uintptr
	Put_CroppedAspectRatio  uintptr
	Get_AllowCropping       uintptr
	Put_AllowCropping       uintptr
}

type ICameraCaptureUIPhotoCaptureSettings struct {
	win32.IInspectable
}

func (this *ICameraCaptureUIPhotoCaptureSettings) Vtbl() *ICameraCaptureUIPhotoCaptureSettingsVtbl {
	return (*ICameraCaptureUIPhotoCaptureSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICameraCaptureUIPhotoCaptureSettings) Get_Format() CameraCaptureUIPhotoFormat {
	var _result CameraCaptureUIPhotoFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Format, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICameraCaptureUIPhotoCaptureSettings) Put_Format(value CameraCaptureUIPhotoFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Format, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICameraCaptureUIPhotoCaptureSettings) Get_MaxResolution() CameraCaptureUIMaxPhotoResolution {
	var _result CameraCaptureUIMaxPhotoResolution
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxResolution, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICameraCaptureUIPhotoCaptureSettings) Put_MaxResolution(value CameraCaptureUIMaxPhotoResolution) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxResolution, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICameraCaptureUIPhotoCaptureSettings) Get_CroppedSizeInPixels() Size {
	var _result Size
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CroppedSizeInPixels, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICameraCaptureUIPhotoCaptureSettings) Put_CroppedSizeInPixels(value Size) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CroppedSizeInPixels, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ICameraCaptureUIPhotoCaptureSettings) Get_CroppedAspectRatio() Size {
	var _result Size
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CroppedAspectRatio, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICameraCaptureUIPhotoCaptureSettings) Put_CroppedAspectRatio(value Size) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CroppedAspectRatio, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ICameraCaptureUIPhotoCaptureSettings) Get_AllowCropping() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllowCropping, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICameraCaptureUIPhotoCaptureSettings) Put_AllowCropping(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AllowCropping, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 64E92D1F-A28D-425A-B84F-E568335FF24E
var IID_ICameraCaptureUIVideoCaptureSettings = syscall.GUID{0x64E92D1F, 0xA28D, 0x425A,
	[8]byte{0xB8, 0x4F, 0xE5, 0x68, 0x33, 0x5F, 0xF2, 0x4E}}

type ICameraCaptureUIVideoCaptureSettingsInterface interface {
	win32.IInspectableInterface
	Get_Format() CameraCaptureUIVideoFormat
	Put_Format(value CameraCaptureUIVideoFormat)
	Get_MaxResolution() CameraCaptureUIMaxVideoResolution
	Put_MaxResolution(value CameraCaptureUIMaxVideoResolution)
	Get_MaxDurationInSeconds() float32
	Put_MaxDurationInSeconds(value float32)
	Get_AllowTrimming() bool
	Put_AllowTrimming(value bool)
}

type ICameraCaptureUIVideoCaptureSettingsVtbl struct {
	win32.IInspectableVtbl
	Get_Format               uintptr
	Put_Format               uintptr
	Get_MaxResolution        uintptr
	Put_MaxResolution        uintptr
	Get_MaxDurationInSeconds uintptr
	Put_MaxDurationInSeconds uintptr
	Get_AllowTrimming        uintptr
	Put_AllowTrimming        uintptr
}

type ICameraCaptureUIVideoCaptureSettings struct {
	win32.IInspectable
}

func (this *ICameraCaptureUIVideoCaptureSettings) Vtbl() *ICameraCaptureUIVideoCaptureSettingsVtbl {
	return (*ICameraCaptureUIVideoCaptureSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICameraCaptureUIVideoCaptureSettings) Get_Format() CameraCaptureUIVideoFormat {
	var _result CameraCaptureUIVideoFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Format, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICameraCaptureUIVideoCaptureSettings) Put_Format(value CameraCaptureUIVideoFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Format, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICameraCaptureUIVideoCaptureSettings) Get_MaxResolution() CameraCaptureUIMaxVideoResolution {
	var _result CameraCaptureUIMaxVideoResolution
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxResolution, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICameraCaptureUIVideoCaptureSettings) Put_MaxResolution(value CameraCaptureUIMaxVideoResolution) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxResolution, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICameraCaptureUIVideoCaptureSettings) Get_MaxDurationInSeconds() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxDurationInSeconds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICameraCaptureUIVideoCaptureSettings) Put_MaxDurationInSeconds(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxDurationInSeconds, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICameraCaptureUIVideoCaptureSettings) Get_AllowTrimming() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllowTrimming, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICameraCaptureUIVideoCaptureSettings) Put_AllowTrimming(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AllowTrimming, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 3B0D5E34-3906-4B7D-946C-7BDE844499AE
var IID_ICameraOptionsUIStatics = syscall.GUID{0x3B0D5E34, 0x3906, 0x4B7D,
	[8]byte{0x94, 0x6C, 0x7B, 0xDE, 0x84, 0x44, 0x99, 0xAE}}

type ICameraOptionsUIStaticsInterface interface {
	win32.IInspectableInterface
	Show(mediaCapture *IMediaCapture)
}

type ICameraOptionsUIStaticsVtbl struct {
	win32.IInspectableVtbl
	Show uintptr
}

type ICameraOptionsUIStatics struct {
	win32.IInspectable
}

func (this *ICameraOptionsUIStatics) Vtbl() *ICameraOptionsUIStaticsVtbl {
	return (*ICameraOptionsUIStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICameraOptionsUIStatics) Show(mediaCapture *IMediaCapture) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Show, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(mediaCapture)))
	_ = _hr
}

// 1DD2DE1F-571B-44D8-8E80-A08A1578766E
var IID_ICapturedFrame = syscall.GUID{0x1DD2DE1F, 0x571B, 0x44D8,
	[8]byte{0x8E, 0x80, 0xA0, 0x8A, 0x15, 0x78, 0x76, 0x6E}}

type ICapturedFrameInterface interface {
	win32.IInspectableInterface
	Get_Width() uint32
	Get_Height() uint32
}

type ICapturedFrameVtbl struct {
	win32.IInspectableVtbl
	Get_Width  uintptr
	Get_Height uintptr
}

type ICapturedFrame struct {
	win32.IInspectable
}

func (this *ICapturedFrame) Vtbl() *ICapturedFrameVtbl {
	return (*ICapturedFrameVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICapturedFrame) Get_Width() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Width, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICapturedFrame) Get_Height() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Height, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 543FA6D1-BD78-4866-ADDA-24314BC65DEA
var IID_ICapturedFrame2 = syscall.GUID{0x543FA6D1, 0xBD78, 0x4866,
	[8]byte{0xAD, 0xDA, 0x24, 0x31, 0x4B, 0xC6, 0x5D, 0xEA}}

type ICapturedFrame2Interface interface {
	win32.IInspectableInterface
	Get_ControlValues() *ICapturedFrameControlValues
	Get_BitmapProperties() *IMap[string, *IBitmapTypedValue]
}

type ICapturedFrame2Vtbl struct {
	win32.IInspectableVtbl
	Get_ControlValues    uintptr
	Get_BitmapProperties uintptr
}

type ICapturedFrame2 struct {
	win32.IInspectable
}

func (this *ICapturedFrame2) Vtbl() *ICapturedFrame2Vtbl {
	return (*ICapturedFrame2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICapturedFrame2) Get_ControlValues() *ICapturedFrameControlValues {
	var _result *ICapturedFrameControlValues
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ControlValues, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICapturedFrame2) Get_BitmapProperties() *IMap[string, *IBitmapTypedValue] {
	var _result *IMap[string, *IBitmapTypedValue]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BitmapProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 90C65B7F-4E0D-4CA4-882D-7A144FED0A90
var IID_ICapturedFrameControlValues = syscall.GUID{0x90C65B7F, 0x4E0D, 0x4CA4,
	[8]byte{0x88, 0x2D, 0x7A, 0x14, 0x4F, 0xED, 0x0A, 0x90}}

type ICapturedFrameControlValuesInterface interface {
	win32.IInspectableInterface
	Get_Exposure() *IReference[TimeSpan]
	Get_ExposureCompensation() *IReference[float32]
	Get_IsoSpeed() *IReference[uint32]
	Get_Focus() *IReference[uint32]
	Get_SceneMode() *IReference[CaptureSceneMode]
	Get_Flashed() *IReference[bool]
	Get_FlashPowerPercent() *IReference[float32]
	Get_WhiteBalance() *IReference[uint32]
	Get_ZoomFactor() *IReference[float32]
}

type ICapturedFrameControlValuesVtbl struct {
	win32.IInspectableVtbl
	Get_Exposure             uintptr
	Get_ExposureCompensation uintptr
	Get_IsoSpeed             uintptr
	Get_Focus                uintptr
	Get_SceneMode            uintptr
	Get_Flashed              uintptr
	Get_FlashPowerPercent    uintptr
	Get_WhiteBalance         uintptr
	Get_ZoomFactor           uintptr
}

type ICapturedFrameControlValues struct {
	win32.IInspectable
}

func (this *ICapturedFrameControlValues) Vtbl() *ICapturedFrameControlValuesVtbl {
	return (*ICapturedFrameControlValuesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICapturedFrameControlValues) Get_Exposure() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Exposure, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICapturedFrameControlValues) Get_ExposureCompensation() *IReference[float32] {
	var _result *IReference[float32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExposureCompensation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICapturedFrameControlValues) Get_IsoSpeed() *IReference[uint32] {
	var _result *IReference[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsoSpeed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICapturedFrameControlValues) Get_Focus() *IReference[uint32] {
	var _result *IReference[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Focus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICapturedFrameControlValues) Get_SceneMode() *IReference[CaptureSceneMode] {
	var _result *IReference[CaptureSceneMode]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SceneMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICapturedFrameControlValues) Get_Flashed() *IReference[bool] {
	var _result *IReference[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Flashed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICapturedFrameControlValues) Get_FlashPowerPercent() *IReference[float32] {
	var _result *IReference[float32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FlashPowerPercent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICapturedFrameControlValues) Get_WhiteBalance() *IReference[uint32] {
	var _result *IReference[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WhiteBalance, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICapturedFrameControlValues) Get_ZoomFactor() *IReference[float32] {
	var _result *IReference[float32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ZoomFactor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 500B2B88-06D2-4AA7-A7DB-D37AF73321D8
var IID_ICapturedFrameControlValues2 = syscall.GUID{0x500B2B88, 0x06D2, 0x4AA7,
	[8]byte{0xA7, 0xDB, 0xD3, 0x7A, 0xF7, 0x33, 0x21, 0xD8}}

type ICapturedFrameControlValues2Interface interface {
	win32.IInspectableInterface
	Get_FocusState() *IReference[MediaCaptureFocusState]
	Get_IsoDigitalGain() *IReference[float64]
	Get_IsoAnalogGain() *IReference[float64]
	Get_SensorFrameRate() *IMediaRatio
	Get_WhiteBalanceGain() *IReference[WhiteBalanceGain]
}

type ICapturedFrameControlValues2Vtbl struct {
	win32.IInspectableVtbl
	Get_FocusState       uintptr
	Get_IsoDigitalGain   uintptr
	Get_IsoAnalogGain    uintptr
	Get_SensorFrameRate  uintptr
	Get_WhiteBalanceGain uintptr
}

type ICapturedFrameControlValues2 struct {
	win32.IInspectable
}

func (this *ICapturedFrameControlValues2) Vtbl() *ICapturedFrameControlValues2Vtbl {
	return (*ICapturedFrameControlValues2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICapturedFrameControlValues2) Get_FocusState() *IReference[MediaCaptureFocusState] {
	var _result *IReference[MediaCaptureFocusState]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FocusState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICapturedFrameControlValues2) Get_IsoDigitalGain() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsoDigitalGain, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICapturedFrameControlValues2) Get_IsoAnalogGain() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsoAnalogGain, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICapturedFrameControlValues2) Get_SensorFrameRate() *IMediaRatio {
	var _result *IMediaRatio
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SensorFrameRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICapturedFrameControlValues2) Get_WhiteBalanceGain() *IReference[WhiteBalanceGain] {
	var _result *IReference[WhiteBalanceGain]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WhiteBalanceGain, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B58E8B6E-8503-49B5-9E86-897D26A3FF3D
var IID_ICapturedFrameWithSoftwareBitmap = syscall.GUID{0xB58E8B6E, 0x8503, 0x49B5,
	[8]byte{0x9E, 0x86, 0x89, 0x7D, 0x26, 0xA3, 0xFF, 0x3D}}

type ICapturedFrameWithSoftwareBitmapInterface interface {
	win32.IInspectableInterface
	Get_SoftwareBitmap() *ISoftwareBitmap
}

type ICapturedFrameWithSoftwareBitmapVtbl struct {
	win32.IInspectableVtbl
	Get_SoftwareBitmap uintptr
}

type ICapturedFrameWithSoftwareBitmap struct {
	win32.IInspectable
}

func (this *ICapturedFrameWithSoftwareBitmap) Vtbl() *ICapturedFrameWithSoftwareBitmapVtbl {
	return (*ICapturedFrameWithSoftwareBitmapVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICapturedFrameWithSoftwareBitmap) Get_SoftwareBitmap() *ISoftwareBitmap {
	var _result *ISoftwareBitmap
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SoftwareBitmap, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B0CE7E5A-CFCC-4D6C-8AD1-0869208ACA16
var IID_ICapturedPhoto = syscall.GUID{0xB0CE7E5A, 0xCFCC, 0x4D6C,
	[8]byte{0x8A, 0xD1, 0x08, 0x69, 0x20, 0x8A, 0xCA, 0x16}}

type ICapturedPhotoInterface interface {
	win32.IInspectableInterface
	Get_Frame() *ICapturedFrame
	Get_Thumbnail() *ICapturedFrame
}

type ICapturedPhotoVtbl struct {
	win32.IInspectableVtbl
	Get_Frame     uintptr
	Get_Thumbnail uintptr
}

type ICapturedPhoto struct {
	win32.IInspectable
}

func (this *ICapturedPhoto) Vtbl() *ICapturedPhotoVtbl {
	return (*ICapturedPhotoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICapturedPhoto) Get_Frame() *ICapturedFrame {
	var _result *ICapturedFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Frame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICapturedPhoto) Get_Thumbnail() *ICapturedFrame {
	var _result *ICapturedFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Thumbnail, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2DBEAD57-50A6-499E-8C6C-D330A7311796
var IID_IGameBarServices = syscall.GUID{0x2DBEAD57, 0x50A6, 0x499E,
	[8]byte{0x8C, 0x6C, 0xD3, 0x30, 0xA7, 0x31, 0x17, 0x96}}

type IGameBarServicesInterface interface {
	win32.IInspectableInterface
	Get_TargetCapturePolicy() GameBarTargetCapturePolicy
	EnableCapture()
	DisableCapture()
	Get_TargetInfo() *IGameBarServicesTargetInfo
	Get_SessionId() string
	Get_AppBroadcastServices() *IAppBroadcastServices
	Get_AppCaptureServices() *IAppCaptureServices
	Add_CommandReceived(value TypedEventHandler[*IGameBarServices, *IGameBarServicesCommandEventArgs]) EventRegistrationToken
	Remove_CommandReceived(token EventRegistrationToken)
}

type IGameBarServicesVtbl struct {
	win32.IInspectableVtbl
	Get_TargetCapturePolicy  uintptr
	EnableCapture            uintptr
	DisableCapture           uintptr
	Get_TargetInfo           uintptr
	Get_SessionId            uintptr
	Get_AppBroadcastServices uintptr
	Get_AppCaptureServices   uintptr
	Add_CommandReceived      uintptr
	Remove_CommandReceived   uintptr
}

type IGameBarServices struct {
	win32.IInspectable
}

func (this *IGameBarServices) Vtbl() *IGameBarServicesVtbl {
	return (*IGameBarServicesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameBarServices) Get_TargetCapturePolicy() GameBarTargetCapturePolicy {
	var _result GameBarTargetCapturePolicy
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TargetCapturePolicy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGameBarServices) EnableCapture() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnableCapture, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IGameBarServices) DisableCapture() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DisableCapture, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IGameBarServices) Get_TargetInfo() *IGameBarServicesTargetInfo {
	var _result *IGameBarServicesTargetInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TargetInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGameBarServices) Get_SessionId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SessionId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGameBarServices) Get_AppBroadcastServices() *IAppBroadcastServices {
	var _result *IAppBroadcastServices
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppBroadcastServices, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGameBarServices) Get_AppCaptureServices() *IAppCaptureServices {
	var _result *IAppCaptureServices
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppCaptureServices, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGameBarServices) Add_CommandReceived(value TypedEventHandler[*IGameBarServices, *IGameBarServicesCommandEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_CommandReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGameBarServices) Remove_CommandReceived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_CommandReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// A74226B2-F176-4FCF-8FBB-CF698B2EB8E0
var IID_IGameBarServicesCommandEventArgs = syscall.GUID{0xA74226B2, 0xF176, 0x4FCF,
	[8]byte{0x8F, 0xBB, 0xCF, 0x69, 0x8B, 0x2E, 0xB8, 0xE0}}

type IGameBarServicesCommandEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Command() GameBarCommand
	Get_Origin() GameBarCommandOrigin
}

type IGameBarServicesCommandEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Command uintptr
	Get_Origin  uintptr
}

type IGameBarServicesCommandEventArgs struct {
	win32.IInspectable
}

func (this *IGameBarServicesCommandEventArgs) Vtbl() *IGameBarServicesCommandEventArgsVtbl {
	return (*IGameBarServicesCommandEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameBarServicesCommandEventArgs) Get_Command() GameBarCommand {
	var _result GameBarCommand
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Command, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGameBarServicesCommandEventArgs) Get_Origin() GameBarCommandOrigin {
	var _result GameBarCommandOrigin
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Origin, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 3A4B9CFA-7F8B-4C60-9DBB-0BCD262DFFC6
var IID_IGameBarServicesManager = syscall.GUID{0x3A4B9CFA, 0x7F8B, 0x4C60,
	[8]byte{0x9D, 0xBB, 0x0B, 0xCD, 0x26, 0x2D, 0xFF, 0xC6}}

type IGameBarServicesManagerInterface interface {
	win32.IInspectableInterface
	Add_GameBarServicesCreated(value TypedEventHandler[*IGameBarServicesManager, *IGameBarServicesManagerGameBarServicesCreatedEventArgs]) EventRegistrationToken
	Remove_GameBarServicesCreated(token EventRegistrationToken)
}

type IGameBarServicesManagerVtbl struct {
	win32.IInspectableVtbl
	Add_GameBarServicesCreated    uintptr
	Remove_GameBarServicesCreated uintptr
}

type IGameBarServicesManager struct {
	win32.IInspectable
}

func (this *IGameBarServicesManager) Vtbl() *IGameBarServicesManagerVtbl {
	return (*IGameBarServicesManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameBarServicesManager) Add_GameBarServicesCreated(value TypedEventHandler[*IGameBarServicesManager, *IGameBarServicesManagerGameBarServicesCreatedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_GameBarServicesCreated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGameBarServicesManager) Remove_GameBarServicesCreated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_GameBarServicesCreated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// EDEDBD9C-143E-49A3-A5EA-0B1995C8D46E
var IID_IGameBarServicesManagerGameBarServicesCreatedEventArgs = syscall.GUID{0xEDEDBD9C, 0x143E, 0x49A3,
	[8]byte{0xA5, 0xEA, 0x0B, 0x19, 0x95, 0xC8, 0xD4, 0x6E}}

type IGameBarServicesManagerGameBarServicesCreatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_GameBarServices() *IGameBarServices
}

type IGameBarServicesManagerGameBarServicesCreatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_GameBarServices uintptr
}

type IGameBarServicesManagerGameBarServicesCreatedEventArgs struct {
	win32.IInspectable
}

func (this *IGameBarServicesManagerGameBarServicesCreatedEventArgs) Vtbl() *IGameBarServicesManagerGameBarServicesCreatedEventArgsVtbl {
	return (*IGameBarServicesManagerGameBarServicesCreatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameBarServicesManagerGameBarServicesCreatedEventArgs) Get_GameBarServices() *IGameBarServices {
	var _result *IGameBarServices
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GameBarServices, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 34C1B616-FF25-4792-98F2-D3753F15AC13
var IID_IGameBarServicesManagerStatics = syscall.GUID{0x34C1B616, 0xFF25, 0x4792,
	[8]byte{0x98, 0xF2, 0xD3, 0x75, 0x3F, 0x15, 0xAC, 0x13}}

type IGameBarServicesManagerStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *IGameBarServicesManager
}

type IGameBarServicesManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
}

type IGameBarServicesManagerStatics struct {
	win32.IInspectable
}

func (this *IGameBarServicesManagerStatics) Vtbl() *IGameBarServicesManagerStaticsVtbl {
	return (*IGameBarServicesManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameBarServicesManagerStatics) GetDefault() *IGameBarServicesManager {
	var _result *IGameBarServicesManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B4202F92-1611-4E05-B6EF-DFD737AE33B0
var IID_IGameBarServicesTargetInfo = syscall.GUID{0xB4202F92, 0x1611, 0x4E05,
	[8]byte{0xB6, 0xEF, 0xDF, 0xD7, 0x37, 0xAE, 0x33, 0xB0}}

type IGameBarServicesTargetInfoInterface interface {
	win32.IInspectableInterface
	Get_DisplayName() string
	Get_AppId() string
	Get_TitleId() string
	Get_DisplayMode() GameBarServicesDisplayMode
}

type IGameBarServicesTargetInfoVtbl struct {
	win32.IInspectableVtbl
	Get_DisplayName uintptr
	Get_AppId       uintptr
	Get_TitleId     uintptr
	Get_DisplayMode uintptr
}

type IGameBarServicesTargetInfo struct {
	win32.IInspectable
}

func (this *IGameBarServicesTargetInfo) Vtbl() *IGameBarServicesTargetInfoVtbl {
	return (*IGameBarServicesTargetInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameBarServicesTargetInfo) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGameBarServicesTargetInfo) Get_AppId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGameBarServicesTargetInfo) Get_TitleId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TitleId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGameBarServicesTargetInfo) Get_DisplayMode() GameBarServicesDisplayMode {
	var _result GameBarServicesDisplayMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 41C8BAF7-FF3F-49F0-A477-F195E3CE5108
var IID_ILowLagMediaRecording = syscall.GUID{0x41C8BAF7, 0xFF3F, 0x49F0,
	[8]byte{0xA4, 0x77, 0xF1, 0x95, 0xE3, 0xCE, 0x51, 0x08}}

type ILowLagMediaRecordingInterface interface {
	win32.IInspectableInterface
	StartAsync() *IAsyncAction
	StopAsync() *IAsyncAction
	FinishAsync() *IAsyncAction
}

type ILowLagMediaRecordingVtbl struct {
	win32.IInspectableVtbl
	StartAsync  uintptr
	StopAsync   uintptr
	FinishAsync uintptr
}

type ILowLagMediaRecording struct {
	win32.IInspectable
}

func (this *ILowLagMediaRecording) Vtbl() *ILowLagMediaRecordingVtbl {
	return (*ILowLagMediaRecordingVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILowLagMediaRecording) StartAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILowLagMediaRecording) StopAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILowLagMediaRecording) FinishAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FinishAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6369C758-5644-41E2-97AF-8EF56A25E225
var IID_ILowLagMediaRecording2 = syscall.GUID{0x6369C758, 0x5644, 0x41E2,
	[8]byte{0x97, 0xAF, 0x8E, 0xF5, 0x6A, 0x25, 0xE2, 0x25}}

type ILowLagMediaRecording2Interface interface {
	win32.IInspectableInterface
	PauseAsync(behavior MediaCapturePauseBehavior) *IAsyncAction
	ResumeAsync() *IAsyncAction
}

type ILowLagMediaRecording2Vtbl struct {
	win32.IInspectableVtbl
	PauseAsync  uintptr
	ResumeAsync uintptr
}

type ILowLagMediaRecording2 struct {
	win32.IInspectable
}

func (this *ILowLagMediaRecording2) Vtbl() *ILowLagMediaRecording2Vtbl {
	return (*ILowLagMediaRecording2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILowLagMediaRecording2) PauseAsync(behavior MediaCapturePauseBehavior) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PauseAsync, uintptr(unsafe.Pointer(this)), uintptr(behavior), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILowLagMediaRecording2) ResumeAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ResumeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5C33AB12-48F7-47DA-B41E-90880A5FE0EC
var IID_ILowLagMediaRecording3 = syscall.GUID{0x5C33AB12, 0x48F7, 0x47DA,
	[8]byte{0xB4, 0x1E, 0x90, 0x88, 0x0A, 0x5F, 0xE0, 0xEC}}

type ILowLagMediaRecording3Interface interface {
	win32.IInspectableInterface
	PauseWithResultAsync(behavior MediaCapturePauseBehavior) *IAsyncOperation[*IMediaCapturePauseResult]
	StopWithResultAsync() *IAsyncOperation[*IMediaCaptureStopResult]
}

type ILowLagMediaRecording3Vtbl struct {
	win32.IInspectableVtbl
	PauseWithResultAsync uintptr
	StopWithResultAsync  uintptr
}

type ILowLagMediaRecording3 struct {
	win32.IInspectable
}

func (this *ILowLagMediaRecording3) Vtbl() *ILowLagMediaRecording3Vtbl {
	return (*ILowLagMediaRecording3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILowLagMediaRecording3) PauseWithResultAsync(behavior MediaCapturePauseBehavior) *IAsyncOperation[*IMediaCapturePauseResult] {
	var _result *IAsyncOperation[*IMediaCapturePauseResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PauseWithResultAsync, uintptr(unsafe.Pointer(this)), uintptr(behavior), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILowLagMediaRecording3) StopWithResultAsync() *IAsyncOperation[*IMediaCaptureStopResult] {
	var _result *IAsyncOperation[*IMediaCaptureStopResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopWithResultAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A37251B7-6B44-473D-8F24-F703D6C0EC44
var IID_ILowLagPhotoCapture = syscall.GUID{0xA37251B7, 0x6B44, 0x473D,
	[8]byte{0x8F, 0x24, 0xF7, 0x03, 0xD6, 0xC0, 0xEC, 0x44}}

type ILowLagPhotoCaptureInterface interface {
	win32.IInspectableInterface
	CaptureAsync() *IAsyncOperation[*ICapturedPhoto]
	FinishAsync() *IAsyncAction
}

type ILowLagPhotoCaptureVtbl struct {
	win32.IInspectableVtbl
	CaptureAsync uintptr
	FinishAsync  uintptr
}

type ILowLagPhotoCapture struct {
	win32.IInspectable
}

func (this *ILowLagPhotoCapture) Vtbl() *ILowLagPhotoCaptureVtbl {
	return (*ILowLagPhotoCaptureVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILowLagPhotoCapture) CaptureAsync() *IAsyncOperation[*ICapturedPhoto] {
	var _result *IAsyncOperation[*ICapturedPhoto]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CaptureAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILowLagPhotoCapture) FinishAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FinishAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7CC346BB-B9A9-4C91-8FFA-287E9C668669
var IID_ILowLagPhotoSequenceCapture = syscall.GUID{0x7CC346BB, 0xB9A9, 0x4C91,
	[8]byte{0x8F, 0xFA, 0x28, 0x7E, 0x9C, 0x66, 0x86, 0x69}}

type ILowLagPhotoSequenceCaptureInterface interface {
	win32.IInspectableInterface
	StartAsync() *IAsyncAction
	StopAsync() *IAsyncAction
	FinishAsync() *IAsyncAction
	Add_PhotoCaptured(handler TypedEventHandler[*ILowLagPhotoSequenceCapture, *IPhotoCapturedEventArgs]) EventRegistrationToken
	Remove_PhotoCaptured(token EventRegistrationToken)
}

type ILowLagPhotoSequenceCaptureVtbl struct {
	win32.IInspectableVtbl
	StartAsync           uintptr
	StopAsync            uintptr
	FinishAsync          uintptr
	Add_PhotoCaptured    uintptr
	Remove_PhotoCaptured uintptr
}

type ILowLagPhotoSequenceCapture struct {
	win32.IInspectable
}

func (this *ILowLagPhotoSequenceCapture) Vtbl() *ILowLagPhotoSequenceCaptureVtbl {
	return (*ILowLagPhotoSequenceCaptureVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILowLagPhotoSequenceCapture) StartAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILowLagPhotoSequenceCapture) StopAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILowLagPhotoSequenceCapture) FinishAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FinishAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILowLagPhotoSequenceCapture) Add_PhotoCaptured(handler TypedEventHandler[*ILowLagPhotoSequenceCapture, *IPhotoCapturedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PhotoCaptured, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILowLagPhotoSequenceCapture) Remove_PhotoCaptured(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PhotoCaptured, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// C61AFBB4-FB10-4A34-AC18-CA80D9C8E7EE
var IID_IMediaCapture = syscall.GUID{0xC61AFBB4, 0xFB10, 0x4A34,
	[8]byte{0xAC, 0x18, 0xCA, 0x80, 0xD9, 0xC8, 0xE7, 0xEE}}

type IMediaCaptureInterface interface {
	win32.IInspectableInterface
	InitializeAsync() *IAsyncAction
	InitializeWithSettingsAsync(mediaCaptureInitializationSettings *IMediaCaptureInitializationSettings) *IAsyncAction
	StartRecordToStorageFileAsync(encodingProfile *IMediaEncodingProfile, file *IStorageFile) *IAsyncAction
	StartRecordToStreamAsync(encodingProfile *IMediaEncodingProfile, stream *IRandomAccessStream) *IAsyncAction
	StartRecordToCustomSinkAsync(encodingProfile *IMediaEncodingProfile, customMediaSink *IMediaExtension) *IAsyncAction
	StartRecordToCustomSinkIdAsync(encodingProfile *IMediaEncodingProfile, customSinkActivationId string, customSinkSettings *IPropertySet) *IAsyncAction
	StopRecordAsync() *IAsyncAction
	CapturePhotoToStorageFileAsync(type_ *IImageEncodingProperties, file *IStorageFile) *IAsyncAction
	CapturePhotoToStreamAsync(type_ *IImageEncodingProperties, stream *IRandomAccessStream) *IAsyncAction
	AddEffectAsync(mediaStreamType MediaStreamType, effectActivationID string, effectSettings *IPropertySet) *IAsyncAction
	ClearEffectsAsync(mediaStreamType MediaStreamType) *IAsyncAction
	SetEncoderProperty(mediaStreamType MediaStreamType, propertyId syscall.GUID, propertyValue interface{})
	GetEncoderProperty(mediaStreamType MediaStreamType, propertyId syscall.GUID) interface{}
	Add_Failed(errorEventHandler MediaCaptureFailedEventHandler) EventRegistrationToken
	Remove_Failed(eventCookie EventRegistrationToken)
	Add_RecordLimitationExceeded(recordLimitationExceededEventHandler RecordLimitationExceededEventHandler) EventRegistrationToken
	Remove_RecordLimitationExceeded(eventCookie EventRegistrationToken)
	Get_MediaCaptureSettings() *IMediaCaptureSettings
	Get_AudioDeviceController() *IAudioDeviceController
	Get_VideoDeviceController() *IVideoDeviceController
	SetPreviewMirroring(value bool)
	GetPreviewMirroring() bool
	SetPreviewRotation(value VideoRotation)
	GetPreviewRotation() VideoRotation
	SetRecordRotation(value VideoRotation)
	GetRecordRotation() VideoRotation
}

type IMediaCaptureVtbl struct {
	win32.IInspectableVtbl
	InitializeAsync                 uintptr
	InitializeWithSettingsAsync     uintptr
	StartRecordToStorageFileAsync   uintptr
	StartRecordToStreamAsync        uintptr
	StartRecordToCustomSinkAsync    uintptr
	StartRecordToCustomSinkIdAsync  uintptr
	StopRecordAsync                 uintptr
	CapturePhotoToStorageFileAsync  uintptr
	CapturePhotoToStreamAsync       uintptr
	AddEffectAsync                  uintptr
	ClearEffectsAsync               uintptr
	SetEncoderProperty              uintptr
	GetEncoderProperty              uintptr
	Add_Failed                      uintptr
	Remove_Failed                   uintptr
	Add_RecordLimitationExceeded    uintptr
	Remove_RecordLimitationExceeded uintptr
	Get_MediaCaptureSettings        uintptr
	Get_AudioDeviceController       uintptr
	Get_VideoDeviceController       uintptr
	SetPreviewMirroring             uintptr
	GetPreviewMirroring             uintptr
	SetPreviewRotation              uintptr
	GetPreviewRotation              uintptr
	SetRecordRotation               uintptr
	GetRecordRotation               uintptr
}

type IMediaCapture struct {
	win32.IInspectable
}

func (this *IMediaCapture) Vtbl() *IMediaCaptureVtbl {
	return (*IMediaCaptureVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCapture) InitializeAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InitializeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture) InitializeWithSettingsAsync(mediaCaptureInitializationSettings *IMediaCaptureInitializationSettings) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InitializeWithSettingsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(mediaCaptureInitializationSettings)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture) StartRecordToStorageFileAsync(encodingProfile *IMediaEncodingProfile, file *IStorageFile) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartRecordToStorageFileAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(encodingProfile)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture) StartRecordToStreamAsync(encodingProfile *IMediaEncodingProfile, stream *IRandomAccessStream) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartRecordToStreamAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(encodingProfile)), uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture) StartRecordToCustomSinkAsync(encodingProfile *IMediaEncodingProfile, customMediaSink *IMediaExtension) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartRecordToCustomSinkAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(encodingProfile)), uintptr(unsafe.Pointer(customMediaSink)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture) StartRecordToCustomSinkIdAsync(encodingProfile *IMediaEncodingProfile, customSinkActivationId string, customSinkSettings *IPropertySet) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartRecordToCustomSinkIdAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(encodingProfile)), NewHStr(customSinkActivationId).Ptr, uintptr(unsafe.Pointer(customSinkSettings)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture) StopRecordAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopRecordAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture) CapturePhotoToStorageFileAsync(type_ *IImageEncodingProperties, file *IStorageFile) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CapturePhotoToStorageFileAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(type_)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture) CapturePhotoToStreamAsync(type_ *IImageEncodingProperties, stream *IRandomAccessStream) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CapturePhotoToStreamAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(type_)), uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture) AddEffectAsync(mediaStreamType MediaStreamType, effectActivationID string, effectSettings *IPropertySet) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddEffectAsync, uintptr(unsafe.Pointer(this)), uintptr(mediaStreamType), NewHStr(effectActivationID).Ptr, uintptr(unsafe.Pointer(effectSettings)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture) ClearEffectsAsync(mediaStreamType MediaStreamType) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearEffectsAsync, uintptr(unsafe.Pointer(this)), uintptr(mediaStreamType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture) SetEncoderProperty(mediaStreamType MediaStreamType, propertyId syscall.GUID, propertyValue interface{}) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetEncoderProperty, uintptr(unsafe.Pointer(this)), uintptr(mediaStreamType), uintptr(unsafe.Pointer(&propertyId)), *(*uintptr)(unsafe.Pointer(&propertyValue)))
	_ = _hr
}

func (this *IMediaCapture) GetEncoderProperty(mediaStreamType MediaStreamType, propertyId syscall.GUID) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetEncoderProperty, uintptr(unsafe.Pointer(this)), uintptr(mediaStreamType), uintptr(unsafe.Pointer(&propertyId)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCapture) Add_Failed(errorEventHandler MediaCaptureFailedEventHandler) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Failed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(errorEventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCapture) Remove_Failed(eventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Failed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&eventCookie)))
	_ = _hr
}

func (this *IMediaCapture) Add_RecordLimitationExceeded(recordLimitationExceededEventHandler RecordLimitationExceededEventHandler) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_RecordLimitationExceeded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewOneArgFuncDelegate(recordLimitationExceededEventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCapture) Remove_RecordLimitationExceeded(eventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_RecordLimitationExceeded, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&eventCookie)))
	_ = _hr
}

func (this *IMediaCapture) Get_MediaCaptureSettings() *IMediaCaptureSettings {
	var _result *IMediaCaptureSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaCaptureSettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture) Get_AudioDeviceController() *IAudioDeviceController {
	var _result *IAudioDeviceController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioDeviceController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture) Get_VideoDeviceController() *IVideoDeviceController {
	var _result *IVideoDeviceController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoDeviceController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture) SetPreviewMirroring(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetPreviewMirroring, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaCapture) GetPreviewMirroring() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPreviewMirroring, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCapture) SetPreviewRotation(value VideoRotation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetPreviewRotation, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMediaCapture) GetPreviewRotation() VideoRotation {
	var _result VideoRotation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPreviewRotation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCapture) SetRecordRotation(value VideoRotation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetRecordRotation, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMediaCapture) GetRecordRotation() VideoRotation {
	var _result VideoRotation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetRecordRotation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 9CC68260-7DA1-4043-B652-21B8878DAFF9
var IID_IMediaCapture2 = syscall.GUID{0x9CC68260, 0x7DA1, 0x4043,
	[8]byte{0xB6, 0x52, 0x21, 0xB8, 0x87, 0x8D, 0xAF, 0xF9}}

type IMediaCapture2Interface interface {
	win32.IInspectableInterface
	PrepareLowLagRecordToStorageFileAsync(encodingProfile *IMediaEncodingProfile, file *IStorageFile) *IAsyncOperation[*ILowLagMediaRecording]
	PrepareLowLagRecordToStreamAsync(encodingProfile *IMediaEncodingProfile, stream *IRandomAccessStream) *IAsyncOperation[*ILowLagMediaRecording]
	PrepareLowLagRecordToCustomSinkAsync(encodingProfile *IMediaEncodingProfile, customMediaSink *IMediaExtension) *IAsyncOperation[*ILowLagMediaRecording]
	PrepareLowLagRecordToCustomSinkIdAsync(encodingProfile *IMediaEncodingProfile, customSinkActivationId string, customSinkSettings *IPropertySet) *IAsyncOperation[*ILowLagMediaRecording]
	PrepareLowLagPhotoCaptureAsync(type_ *IImageEncodingProperties) *IAsyncOperation[*ILowLagPhotoCapture]
	PrepareLowLagPhotoSequenceCaptureAsync(type_ *IImageEncodingProperties) *IAsyncOperation[*ILowLagPhotoSequenceCapture]
	SetEncodingPropertiesAsync(mediaStreamType MediaStreamType, mediaEncodingProperties *IMediaEncodingProperties, encoderProperties *IMap[syscall.GUID, interface{}]) *IAsyncAction
}

type IMediaCapture2Vtbl struct {
	win32.IInspectableVtbl
	PrepareLowLagRecordToStorageFileAsync  uintptr
	PrepareLowLagRecordToStreamAsync       uintptr
	PrepareLowLagRecordToCustomSinkAsync   uintptr
	PrepareLowLagRecordToCustomSinkIdAsync uintptr
	PrepareLowLagPhotoCaptureAsync         uintptr
	PrepareLowLagPhotoSequenceCaptureAsync uintptr
	SetEncodingPropertiesAsync             uintptr
}

type IMediaCapture2 struct {
	win32.IInspectable
}

func (this *IMediaCapture2) Vtbl() *IMediaCapture2Vtbl {
	return (*IMediaCapture2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCapture2) PrepareLowLagRecordToStorageFileAsync(encodingProfile *IMediaEncodingProfile, file *IStorageFile) *IAsyncOperation[*ILowLagMediaRecording] {
	var _result *IAsyncOperation[*ILowLagMediaRecording]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PrepareLowLagRecordToStorageFileAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(encodingProfile)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture2) PrepareLowLagRecordToStreamAsync(encodingProfile *IMediaEncodingProfile, stream *IRandomAccessStream) *IAsyncOperation[*ILowLagMediaRecording] {
	var _result *IAsyncOperation[*ILowLagMediaRecording]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PrepareLowLagRecordToStreamAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(encodingProfile)), uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture2) PrepareLowLagRecordToCustomSinkAsync(encodingProfile *IMediaEncodingProfile, customMediaSink *IMediaExtension) *IAsyncOperation[*ILowLagMediaRecording] {
	var _result *IAsyncOperation[*ILowLagMediaRecording]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PrepareLowLagRecordToCustomSinkAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(encodingProfile)), uintptr(unsafe.Pointer(customMediaSink)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture2) PrepareLowLagRecordToCustomSinkIdAsync(encodingProfile *IMediaEncodingProfile, customSinkActivationId string, customSinkSettings *IPropertySet) *IAsyncOperation[*ILowLagMediaRecording] {
	var _result *IAsyncOperation[*ILowLagMediaRecording]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PrepareLowLagRecordToCustomSinkIdAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(encodingProfile)), NewHStr(customSinkActivationId).Ptr, uintptr(unsafe.Pointer(customSinkSettings)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture2) PrepareLowLagPhotoCaptureAsync(type_ *IImageEncodingProperties) *IAsyncOperation[*ILowLagPhotoCapture] {
	var _result *IAsyncOperation[*ILowLagPhotoCapture]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PrepareLowLagPhotoCaptureAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(type_)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture2) PrepareLowLagPhotoSequenceCaptureAsync(type_ *IImageEncodingProperties) *IAsyncOperation[*ILowLagPhotoSequenceCapture] {
	var _result *IAsyncOperation[*ILowLagPhotoSequenceCapture]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PrepareLowLagPhotoSequenceCaptureAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(type_)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture2) SetEncodingPropertiesAsync(mediaStreamType MediaStreamType, mediaEncodingProperties *IMediaEncodingProperties, encoderProperties *IMap[syscall.GUID, interface{}]) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetEncodingPropertiesAsync, uintptr(unsafe.Pointer(this)), uintptr(mediaStreamType), uintptr(unsafe.Pointer(mediaEncodingProperties)), uintptr(unsafe.Pointer(encoderProperties)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D4136F30-1564-466E-BC0A-AF94E02AB016
var IID_IMediaCapture3 = syscall.GUID{0xD4136F30, 0x1564, 0x466E,
	[8]byte{0xBC, 0x0A, 0xAF, 0x94, 0xE0, 0x2A, 0xB0, 0x16}}

type IMediaCapture3Interface interface {
	win32.IInspectableInterface
	PrepareVariablePhotoSequenceCaptureAsync(type_ *IImageEncodingProperties) *IAsyncOperation[*IVariablePhotoSequenceCapture]
	Add_FocusChanged(handler TypedEventHandler[*IMediaCapture, *IMediaCaptureFocusChangedEventArgs]) EventRegistrationToken
	Remove_FocusChanged(token EventRegistrationToken)
	Add_PhotoConfirmationCaptured(handler TypedEventHandler[*IMediaCapture, *IPhotoConfirmationCapturedEventArgs]) EventRegistrationToken
	Remove_PhotoConfirmationCaptured(token EventRegistrationToken)
}

type IMediaCapture3Vtbl struct {
	win32.IInspectableVtbl
	PrepareVariablePhotoSequenceCaptureAsync uintptr
	Add_FocusChanged                         uintptr
	Remove_FocusChanged                      uintptr
	Add_PhotoConfirmationCaptured            uintptr
	Remove_PhotoConfirmationCaptured         uintptr
}

type IMediaCapture3 struct {
	win32.IInspectable
}

func (this *IMediaCapture3) Vtbl() *IMediaCapture3Vtbl {
	return (*IMediaCapture3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCapture3) PrepareVariablePhotoSequenceCaptureAsync(type_ *IImageEncodingProperties) *IAsyncOperation[*IVariablePhotoSequenceCapture] {
	var _result *IAsyncOperation[*IVariablePhotoSequenceCapture]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PrepareVariablePhotoSequenceCaptureAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(type_)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture3) Add_FocusChanged(handler TypedEventHandler[*IMediaCapture, *IMediaCaptureFocusChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_FocusChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCapture3) Remove_FocusChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_FocusChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaCapture3) Add_PhotoConfirmationCaptured(handler TypedEventHandler[*IMediaCapture, *IPhotoConfirmationCapturedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PhotoConfirmationCaptured, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCapture3) Remove_PhotoConfirmationCaptured(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PhotoConfirmationCaptured, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// BACD6FD6-FB08-4947-AEA2-CE14EFF0CE13
var IID_IMediaCapture4 = syscall.GUID{0xBACD6FD6, 0xFB08, 0x4947,
	[8]byte{0xAE, 0xA2, 0xCE, 0x14, 0xEF, 0xF0, 0xCE, 0x13}}

type IMediaCapture4Interface interface {
	win32.IInspectableInterface
	AddAudioEffectAsync(definition *IAudioEffectDefinition) *IAsyncOperation[*IMediaExtension]
	AddVideoEffectAsync(definition *IVideoEffectDefinition, mediaStreamType MediaStreamType) *IAsyncOperation[*IMediaExtension]
	PauseRecordAsync(behavior MediaCapturePauseBehavior) *IAsyncAction
	ResumeRecordAsync() *IAsyncAction
	Add_CameraStreamStateChanged(handler TypedEventHandler[*IMediaCapture, interface{}]) EventRegistrationToken
	Remove_CameraStreamStateChanged(token EventRegistrationToken)
	Get_CameraStreamState() CameraStreamState
	GetPreviewFrameAsync() *IAsyncOperation[*IVideoFrame]
	GetPreviewFrameCopyAsync(destination *IVideoFrame) *IAsyncOperation[*IVideoFrame]
	Add_ThermalStatusChanged(handler TypedEventHandler[*IMediaCapture, interface{}]) EventRegistrationToken
	Remove_ThermalStatusChanged(token EventRegistrationToken)
	Get_ThermalStatus() MediaCaptureThermalStatus
	PrepareAdvancedPhotoCaptureAsync(encodingProperties *IImageEncodingProperties) *IAsyncOperation[*IAdvancedPhotoCapture]
}

type IMediaCapture4Vtbl struct {
	win32.IInspectableVtbl
	AddAudioEffectAsync              uintptr
	AddVideoEffectAsync              uintptr
	PauseRecordAsync                 uintptr
	ResumeRecordAsync                uintptr
	Add_CameraStreamStateChanged     uintptr
	Remove_CameraStreamStateChanged  uintptr
	Get_CameraStreamState            uintptr
	GetPreviewFrameAsync             uintptr
	GetPreviewFrameCopyAsync         uintptr
	Add_ThermalStatusChanged         uintptr
	Remove_ThermalStatusChanged      uintptr
	Get_ThermalStatus                uintptr
	PrepareAdvancedPhotoCaptureAsync uintptr
}

type IMediaCapture4 struct {
	win32.IInspectable
}

func (this *IMediaCapture4) Vtbl() *IMediaCapture4Vtbl {
	return (*IMediaCapture4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCapture4) AddAudioEffectAsync(definition *IAudioEffectDefinition) *IAsyncOperation[*IMediaExtension] {
	var _result *IAsyncOperation[*IMediaExtension]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddAudioEffectAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(definition)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture4) AddVideoEffectAsync(definition *IVideoEffectDefinition, mediaStreamType MediaStreamType) *IAsyncOperation[*IMediaExtension] {
	var _result *IAsyncOperation[*IMediaExtension]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddVideoEffectAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(definition)), uintptr(mediaStreamType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture4) PauseRecordAsync(behavior MediaCapturePauseBehavior) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PauseRecordAsync, uintptr(unsafe.Pointer(this)), uintptr(behavior), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture4) ResumeRecordAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ResumeRecordAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture4) Add_CameraStreamStateChanged(handler TypedEventHandler[*IMediaCapture, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_CameraStreamStateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCapture4) Remove_CameraStreamStateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_CameraStreamStateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaCapture4) Get_CameraStreamState() CameraStreamState {
	var _result CameraStreamState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CameraStreamState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCapture4) GetPreviewFrameAsync() *IAsyncOperation[*IVideoFrame] {
	var _result *IAsyncOperation[*IVideoFrame]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPreviewFrameAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture4) GetPreviewFrameCopyAsync(destination *IVideoFrame) *IAsyncOperation[*IVideoFrame] {
	var _result *IAsyncOperation[*IVideoFrame]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPreviewFrameCopyAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(destination)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture4) Add_ThermalStatusChanged(handler TypedEventHandler[*IMediaCapture, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ThermalStatusChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCapture4) Remove_ThermalStatusChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ThermalStatusChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaCapture4) Get_ThermalStatus() MediaCaptureThermalStatus {
	var _result MediaCaptureThermalStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ThermalStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCapture4) PrepareAdvancedPhotoCaptureAsync(encodingProperties *IImageEncodingProperties) *IAsyncOperation[*IAdvancedPhotoCapture] {
	var _result *IAsyncOperation[*IAdvancedPhotoCapture]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PrepareAdvancedPhotoCaptureAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(encodingProperties)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DA787C22-3A9B-4720-A71E-97900A316E5A
var IID_IMediaCapture5 = syscall.GUID{0xDA787C22, 0x3A9B, 0x4720,
	[8]byte{0xA7, 0x1E, 0x97, 0x90, 0x0A, 0x31, 0x6E, 0x5A}}

type IMediaCapture5Interface interface {
	win32.IInspectableInterface
	RemoveEffectAsync(effect *IMediaExtension) *IAsyncAction
	PauseRecordWithResultAsync(behavior MediaCapturePauseBehavior) *IAsyncOperation[*IMediaCapturePauseResult]
	StopRecordWithResultAsync() *IAsyncOperation[*IMediaCaptureStopResult]
	Get_FrameSources() *IMapView[string, *IMediaFrameSource]
	CreateFrameReaderAsync(inputSource *IMediaFrameSource) *IAsyncOperation[*IMediaFrameReader]
	CreateFrameReaderWithSubtypeAsync(inputSource *IMediaFrameSource, outputSubtype string) *IAsyncOperation[*IMediaFrameReader]
	CreateFrameReaderWithSubtypeAndSizeAsync(inputSource *IMediaFrameSource, outputSubtype string, outputSize BitmapSize) *IAsyncOperation[*IMediaFrameReader]
}

type IMediaCapture5Vtbl struct {
	win32.IInspectableVtbl
	RemoveEffectAsync                        uintptr
	PauseRecordWithResultAsync               uintptr
	StopRecordWithResultAsync                uintptr
	Get_FrameSources                         uintptr
	CreateFrameReaderAsync                   uintptr
	CreateFrameReaderWithSubtypeAsync        uintptr
	CreateFrameReaderWithSubtypeAndSizeAsync uintptr
}

type IMediaCapture5 struct {
	win32.IInspectable
}

func (this *IMediaCapture5) Vtbl() *IMediaCapture5Vtbl {
	return (*IMediaCapture5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCapture5) RemoveEffectAsync(effect *IMediaExtension) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveEffectAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(effect)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture5) PauseRecordWithResultAsync(behavior MediaCapturePauseBehavior) *IAsyncOperation[*IMediaCapturePauseResult] {
	var _result *IAsyncOperation[*IMediaCapturePauseResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PauseRecordWithResultAsync, uintptr(unsafe.Pointer(this)), uintptr(behavior), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture5) StopRecordWithResultAsync() *IAsyncOperation[*IMediaCaptureStopResult] {
	var _result *IAsyncOperation[*IMediaCaptureStopResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopRecordWithResultAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture5) Get_FrameSources() *IMapView[string, *IMediaFrameSource] {
	var _result *IMapView[string, *IMediaFrameSource]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameSources, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture5) CreateFrameReaderAsync(inputSource *IMediaFrameSource) *IAsyncOperation[*IMediaFrameReader] {
	var _result *IAsyncOperation[*IMediaFrameReader]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFrameReaderAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(inputSource)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture5) CreateFrameReaderWithSubtypeAsync(inputSource *IMediaFrameSource, outputSubtype string) *IAsyncOperation[*IMediaFrameReader] {
	var _result *IAsyncOperation[*IMediaFrameReader]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFrameReaderWithSubtypeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(inputSource)), NewHStr(outputSubtype).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapture5) CreateFrameReaderWithSubtypeAndSizeAsync(inputSource *IMediaFrameSource, outputSubtype string, outputSize BitmapSize) *IAsyncOperation[*IMediaFrameReader] {
	var _result *IAsyncOperation[*IMediaFrameReader]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFrameReaderWithSubtypeAndSizeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(inputSource)), NewHStr(outputSubtype).Ptr, *(*uintptr)(unsafe.Pointer(&outputSize)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 228948BD-4B20-4BB1-9FD6-A583212A1012
var IID_IMediaCapture6 = syscall.GUID{0x228948BD, 0x4B20, 0x4BB1,
	[8]byte{0x9F, 0xD6, 0xA5, 0x83, 0x21, 0x2A, 0x10, 0x12}}

type IMediaCapture6Interface interface {
	win32.IInspectableInterface
	Add_CaptureDeviceExclusiveControlStatusChanged(handler TypedEventHandler[*IMediaCapture, *IMediaCaptureDeviceExclusiveControlStatusChangedEventArgs]) EventRegistrationToken
	Remove_CaptureDeviceExclusiveControlStatusChanged(token EventRegistrationToken)
	CreateMultiSourceFrameReaderAsync(inputSources *IIterable[*IMediaFrameSource]) *IAsyncOperation[*IMultiSourceMediaFrameReader]
}

type IMediaCapture6Vtbl struct {
	win32.IInspectableVtbl
	Add_CaptureDeviceExclusiveControlStatusChanged    uintptr
	Remove_CaptureDeviceExclusiveControlStatusChanged uintptr
	CreateMultiSourceFrameReaderAsync                 uintptr
}

type IMediaCapture6 struct {
	win32.IInspectable
}

func (this *IMediaCapture6) Vtbl() *IMediaCapture6Vtbl {
	return (*IMediaCapture6Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCapture6) Add_CaptureDeviceExclusiveControlStatusChanged(handler TypedEventHandler[*IMediaCapture, *IMediaCaptureDeviceExclusiveControlStatusChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_CaptureDeviceExclusiveControlStatusChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCapture6) Remove_CaptureDeviceExclusiveControlStatusChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_CaptureDeviceExclusiveControlStatusChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaCapture6) CreateMultiSourceFrameReaderAsync(inputSources *IIterable[*IMediaFrameSource]) *IAsyncOperation[*IMultiSourceMediaFrameReader] {
	var _result *IAsyncOperation[*IMultiSourceMediaFrameReader]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateMultiSourceFrameReaderAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(inputSources)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9169F102-8888-541A-95BC-24E4D462542A
var IID_IMediaCapture7 = syscall.GUID{0x9169F102, 0x8888, 0x541A,
	[8]byte{0x95, 0xBC, 0x24, 0xE4, 0xD4, 0x62, 0x54, 0x2A}}

type IMediaCapture7Interface interface {
	win32.IInspectableInterface
	CreateRelativePanelWatcher(captureMode StreamingCaptureMode, displayRegion unsafe.Pointer) *IMediaCaptureRelativePanelWatcher
}

type IMediaCapture7Vtbl struct {
	win32.IInspectableVtbl
	CreateRelativePanelWatcher uintptr
}

type IMediaCapture7 struct {
	win32.IInspectable
}

func (this *IMediaCapture7) Vtbl() *IMediaCapture7Vtbl {
	return (*IMediaCapture7Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCapture7) CreateRelativePanelWatcher(captureMode StreamingCaptureMode, displayRegion unsafe.Pointer) *IMediaCaptureRelativePanelWatcher {
	var _result *IMediaCaptureRelativePanelWatcher
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateRelativePanelWatcher, uintptr(unsafe.Pointer(this)), uintptr(captureMode), uintptr(displayRegion), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9D2F920D-A588-43C6-89D6-5AD322AF006A
var IID_IMediaCaptureDeviceExclusiveControlStatusChangedEventArgs = syscall.GUID{0x9D2F920D, 0xA588, 0x43C6,
	[8]byte{0x89, 0xD6, 0x5A, 0xD3, 0x22, 0xAF, 0x00, 0x6A}}

type IMediaCaptureDeviceExclusiveControlStatusChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_Status() MediaCaptureDeviceExclusiveControlStatus
}

type IMediaCaptureDeviceExclusiveControlStatusChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId uintptr
	Get_Status   uintptr
}

type IMediaCaptureDeviceExclusiveControlStatusChangedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaCaptureDeviceExclusiveControlStatusChangedEventArgs) Vtbl() *IMediaCaptureDeviceExclusiveControlStatusChangedEventArgsVtbl {
	return (*IMediaCaptureDeviceExclusiveControlStatusChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCaptureDeviceExclusiveControlStatusChangedEventArgs) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaCaptureDeviceExclusiveControlStatusChangedEventArgs) Get_Status() MediaCaptureDeviceExclusiveControlStatus {
	var _result MediaCaptureDeviceExclusiveControlStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 80FDE3F4-54C4-42C0-8D19-CEA1A87CA18B
var IID_IMediaCaptureFailedEventArgs = syscall.GUID{0x80FDE3F4, 0x54C4, 0x42C0,
	[8]byte{0x8D, 0x19, 0xCE, 0xA1, 0xA8, 0x7C, 0xA1, 0x8B}}

type IMediaCaptureFailedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Message() string
	Get_Code() uint32
}

type IMediaCaptureFailedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Message uintptr
	Get_Code    uintptr
}

type IMediaCaptureFailedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaCaptureFailedEventArgs) Vtbl() *IMediaCaptureFailedEventArgsVtbl {
	return (*IMediaCaptureFailedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCaptureFailedEventArgs) Get_Message() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Message, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaCaptureFailedEventArgs) Get_Code() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Code, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 81E1BC7F-2277-493E-ABEE-D3F44FF98C04
var IID_IMediaCaptureFocusChangedEventArgs = syscall.GUID{0x81E1BC7F, 0x2277, 0x493E,
	[8]byte{0xAB, 0xEE, 0xD3, 0xF4, 0x4F, 0xF9, 0x8C, 0x04}}

type IMediaCaptureFocusChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_FocusState() MediaCaptureFocusState
}

type IMediaCaptureFocusChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_FocusState uintptr
}

type IMediaCaptureFocusChangedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaCaptureFocusChangedEventArgs) Vtbl() *IMediaCaptureFocusChangedEventArgsVtbl {
	return (*IMediaCaptureFocusChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCaptureFocusChangedEventArgs) Get_FocusState() MediaCaptureFocusState {
	var _result MediaCaptureFocusState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FocusState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 9782BA70-EA65-4900-9356-8CA887726884
var IID_IMediaCaptureInitializationSettings = syscall.GUID{0x9782BA70, 0xEA65, 0x4900,
	[8]byte{0x93, 0x56, 0x8C, 0xA8, 0x87, 0x72, 0x68, 0x84}}

type IMediaCaptureInitializationSettingsInterface interface {
	win32.IInspectableInterface
	Put_AudioDeviceId(value string)
	Get_AudioDeviceId() string
	Put_VideoDeviceId(value string)
	Get_VideoDeviceId() string
	Put_StreamingCaptureMode(value StreamingCaptureMode)
	Get_StreamingCaptureMode() StreamingCaptureMode
	Put_PhotoCaptureSource(value PhotoCaptureSource)
	Get_PhotoCaptureSource() PhotoCaptureSource
}

type IMediaCaptureInitializationSettingsVtbl struct {
	win32.IInspectableVtbl
	Put_AudioDeviceId        uintptr
	Get_AudioDeviceId        uintptr
	Put_VideoDeviceId        uintptr
	Get_VideoDeviceId        uintptr
	Put_StreamingCaptureMode uintptr
	Get_StreamingCaptureMode uintptr
	Put_PhotoCaptureSource   uintptr
	Get_PhotoCaptureSource   uintptr
}

type IMediaCaptureInitializationSettings struct {
	win32.IInspectable
}

func (this *IMediaCaptureInitializationSettings) Vtbl() *IMediaCaptureInitializationSettingsVtbl {
	return (*IMediaCaptureInitializationSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCaptureInitializationSettings) Put_AudioDeviceId(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AudioDeviceId, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMediaCaptureInitializationSettings) Get_AudioDeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioDeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaCaptureInitializationSettings) Put_VideoDeviceId(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_VideoDeviceId, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMediaCaptureInitializationSettings) Get_VideoDeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoDeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaCaptureInitializationSettings) Put_StreamingCaptureMode(value StreamingCaptureMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StreamingCaptureMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMediaCaptureInitializationSettings) Get_StreamingCaptureMode() StreamingCaptureMode {
	var _result StreamingCaptureMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StreamingCaptureMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCaptureInitializationSettings) Put_PhotoCaptureSource(value PhotoCaptureSource) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PhotoCaptureSource, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMediaCaptureInitializationSettings) Get_PhotoCaptureSource() PhotoCaptureSource {
	var _result PhotoCaptureSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhotoCaptureSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 404E0626-C9DC-43E9-AEE4-E6BF1B57B44C
var IID_IMediaCaptureInitializationSettings2 = syscall.GUID{0x404E0626, 0xC9DC, 0x43E9,
	[8]byte{0xAE, 0xE4, 0xE6, 0xBF, 0x1B, 0x57, 0xB4, 0x4C}}

type IMediaCaptureInitializationSettings2Interface interface {
	win32.IInspectableInterface
	Put_MediaCategory(value MediaCategory)
	Get_MediaCategory() MediaCategory
	Put_AudioProcessing(value AudioProcessing)
	Get_AudioProcessing() AudioProcessing
}

type IMediaCaptureInitializationSettings2Vtbl struct {
	win32.IInspectableVtbl
	Put_MediaCategory   uintptr
	Get_MediaCategory   uintptr
	Put_AudioProcessing uintptr
	Get_AudioProcessing uintptr
}

type IMediaCaptureInitializationSettings2 struct {
	win32.IInspectable
}

func (this *IMediaCaptureInitializationSettings2) Vtbl() *IMediaCaptureInitializationSettings2Vtbl {
	return (*IMediaCaptureInitializationSettings2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCaptureInitializationSettings2) Put_MediaCategory(value MediaCategory) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MediaCategory, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMediaCaptureInitializationSettings2) Get_MediaCategory() MediaCategory {
	var _result MediaCategory
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaCategory, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCaptureInitializationSettings2) Put_AudioProcessing(value AudioProcessing) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AudioProcessing, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMediaCaptureInitializationSettings2) Get_AudioProcessing() AudioProcessing {
	var _result AudioProcessing
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioProcessing, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4160519D-BE48-4730-8104-0CF6E9E97948
var IID_IMediaCaptureInitializationSettings3 = syscall.GUID{0x4160519D, 0xBE48, 0x4730,
	[8]byte{0x81, 0x04, 0x0C, 0xF6, 0xE9, 0xE9, 0x79, 0x48}}

type IMediaCaptureInitializationSettings3Interface interface {
	win32.IInspectableInterface
	Put_AudioSource(value *IMediaSource)
	Get_AudioSource() *IMediaSource
	Put_VideoSource(value *IMediaSource)
	Get_VideoSource() *IMediaSource
}

type IMediaCaptureInitializationSettings3Vtbl struct {
	win32.IInspectableVtbl
	Put_AudioSource uintptr
	Get_AudioSource uintptr
	Put_VideoSource uintptr
	Get_VideoSource uintptr
}

type IMediaCaptureInitializationSettings3 struct {
	win32.IInspectable
}

func (this *IMediaCaptureInitializationSettings3) Vtbl() *IMediaCaptureInitializationSettings3Vtbl {
	return (*IMediaCaptureInitializationSettings3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCaptureInitializationSettings3) Put_AudioSource(value *IMediaSource) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AudioSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaCaptureInitializationSettings3) Get_AudioSource() *IMediaSource {
	var _result *IMediaSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCaptureInitializationSettings3) Put_VideoSource(value *IMediaSource) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_VideoSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaCaptureInitializationSettings3) Get_VideoSource() *IMediaSource {
	var _result *IMediaSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F502A537-4CB7-4D28-95ED-4F9F012E0518
var IID_IMediaCaptureInitializationSettings4 = syscall.GUID{0xF502A537, 0x4CB7, 0x4D28,
	[8]byte{0x95, 0xED, 0x4F, 0x9F, 0x01, 0x2E, 0x05, 0x18}}

type IMediaCaptureInitializationSettings4Interface interface {
	win32.IInspectableInterface
	Get_VideoProfile() *IMediaCaptureVideoProfile
	Put_VideoProfile(value *IMediaCaptureVideoProfile)
	Get_PreviewMediaDescription() *IMediaCaptureVideoProfileMediaDescription
	Put_PreviewMediaDescription(value *IMediaCaptureVideoProfileMediaDescription)
	Get_RecordMediaDescription() *IMediaCaptureVideoProfileMediaDescription
	Put_RecordMediaDescription(value *IMediaCaptureVideoProfileMediaDescription)
	Get_PhotoMediaDescription() *IMediaCaptureVideoProfileMediaDescription
	Put_PhotoMediaDescription(value *IMediaCaptureVideoProfileMediaDescription)
}

type IMediaCaptureInitializationSettings4Vtbl struct {
	win32.IInspectableVtbl
	Get_VideoProfile            uintptr
	Put_VideoProfile            uintptr
	Get_PreviewMediaDescription uintptr
	Put_PreviewMediaDescription uintptr
	Get_RecordMediaDescription  uintptr
	Put_RecordMediaDescription  uintptr
	Get_PhotoMediaDescription   uintptr
	Put_PhotoMediaDescription   uintptr
}

type IMediaCaptureInitializationSettings4 struct {
	win32.IInspectable
}

func (this *IMediaCaptureInitializationSettings4) Vtbl() *IMediaCaptureInitializationSettings4Vtbl {
	return (*IMediaCaptureInitializationSettings4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCaptureInitializationSettings4) Get_VideoProfile() *IMediaCaptureVideoProfile {
	var _result *IMediaCaptureVideoProfile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoProfile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCaptureInitializationSettings4) Put_VideoProfile(value *IMediaCaptureVideoProfile) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_VideoProfile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaCaptureInitializationSettings4) Get_PreviewMediaDescription() *IMediaCaptureVideoProfileMediaDescription {
	var _result *IMediaCaptureVideoProfileMediaDescription
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PreviewMediaDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCaptureInitializationSettings4) Put_PreviewMediaDescription(value *IMediaCaptureVideoProfileMediaDescription) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PreviewMediaDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaCaptureInitializationSettings4) Get_RecordMediaDescription() *IMediaCaptureVideoProfileMediaDescription {
	var _result *IMediaCaptureVideoProfileMediaDescription
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RecordMediaDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCaptureInitializationSettings4) Put_RecordMediaDescription(value *IMediaCaptureVideoProfileMediaDescription) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RecordMediaDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaCaptureInitializationSettings4) Get_PhotoMediaDescription() *IMediaCaptureVideoProfileMediaDescription {
	var _result *IMediaCaptureVideoProfileMediaDescription
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhotoMediaDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCaptureInitializationSettings4) Put_PhotoMediaDescription(value *IMediaCaptureVideoProfileMediaDescription) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PhotoMediaDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// D5A2E3B8-2626-4E94-B7B3-5308A0F64B1A
var IID_IMediaCaptureInitializationSettings5 = syscall.GUID{0xD5A2E3B8, 0x2626, 0x4E94,
	[8]byte{0xB7, 0xB3, 0x53, 0x08, 0xA0, 0xF6, 0x4B, 0x1A}}

type IMediaCaptureInitializationSettings5Interface interface {
	win32.IInspectableInterface
	Get_SourceGroup() *IMediaFrameSourceGroup
	Put_SourceGroup(value *IMediaFrameSourceGroup)
	Get_SharingMode() MediaCaptureSharingMode
	Put_SharingMode(value MediaCaptureSharingMode)
	Get_MemoryPreference() MediaCaptureMemoryPreference
	Put_MemoryPreference(value MediaCaptureMemoryPreference)
}

type IMediaCaptureInitializationSettings5Vtbl struct {
	win32.IInspectableVtbl
	Get_SourceGroup      uintptr
	Put_SourceGroup      uintptr
	Get_SharingMode      uintptr
	Put_SharingMode      uintptr
	Get_MemoryPreference uintptr
	Put_MemoryPreference uintptr
}

type IMediaCaptureInitializationSettings5 struct {
	win32.IInspectable
}

func (this *IMediaCaptureInitializationSettings5) Vtbl() *IMediaCaptureInitializationSettings5Vtbl {
	return (*IMediaCaptureInitializationSettings5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCaptureInitializationSettings5) Get_SourceGroup() *IMediaFrameSourceGroup {
	var _result *IMediaFrameSourceGroup
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SourceGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCaptureInitializationSettings5) Put_SourceGroup(value *IMediaFrameSourceGroup) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SourceGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaCaptureInitializationSettings5) Get_SharingMode() MediaCaptureSharingMode {
	var _result MediaCaptureSharingMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SharingMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCaptureInitializationSettings5) Put_SharingMode(value MediaCaptureSharingMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SharingMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMediaCaptureInitializationSettings5) Get_MemoryPreference() MediaCaptureMemoryPreference {
	var _result MediaCaptureMemoryPreference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MemoryPreference, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCaptureInitializationSettings5) Put_MemoryPreference(value MediaCaptureMemoryPreference) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MemoryPreference, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// B2E26B47-3DB1-4D33-AB63-0FFA09056585
var IID_IMediaCaptureInitializationSettings6 = syscall.GUID{0xB2E26B47, 0x3DB1, 0x4D33,
	[8]byte{0xAB, 0x63, 0x0F, 0xFA, 0x09, 0x05, 0x65, 0x85}}

type IMediaCaptureInitializationSettings6Interface interface {
	win32.IInspectableInterface
	Get_AlwaysPlaySystemShutterSound() bool
	Put_AlwaysPlaySystemShutterSound(value bool)
}

type IMediaCaptureInitializationSettings6Vtbl struct {
	win32.IInspectableVtbl
	Get_AlwaysPlaySystemShutterSound uintptr
	Put_AlwaysPlaySystemShutterSound uintptr
}

type IMediaCaptureInitializationSettings6 struct {
	win32.IInspectable
}

func (this *IMediaCaptureInitializationSettings6) Vtbl() *IMediaCaptureInitializationSettings6Vtbl {
	return (*IMediaCaptureInitializationSettings6Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCaptureInitializationSettings6) Get_AlwaysPlaySystemShutterSound() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlwaysPlaySystemShutterSound, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCaptureInitializationSettings6) Put_AlwaysPlaySystemShutterSound(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AlwaysPlaySystemShutterSound, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 41546967-F58A-5D82-9EF4-ED572FB5E34E
var IID_IMediaCaptureInitializationSettings7 = syscall.GUID{0x41546967, 0xF58A, 0x5D82,
	[8]byte{0x9E, 0xF4, 0xED, 0x57, 0x2F, 0xB5, 0xE3, 0x4E}}

type IMediaCaptureInitializationSettings7Interface interface {
	win32.IInspectableInterface
	Get_DeviceUriPasswordCredential() *IPasswordCredential
	Put_DeviceUriPasswordCredential(value *IPasswordCredential)
	Get_DeviceUri() *IUriRuntimeClass
	Put_DeviceUri(value *IUriRuntimeClass)
}

type IMediaCaptureInitializationSettings7Vtbl struct {
	win32.IInspectableVtbl
	Get_DeviceUriPasswordCredential uintptr
	Put_DeviceUriPasswordCredential uintptr
	Get_DeviceUri                   uintptr
	Put_DeviceUri                   uintptr
}

type IMediaCaptureInitializationSettings7 struct {
	win32.IInspectable
}

func (this *IMediaCaptureInitializationSettings7) Vtbl() *IMediaCaptureInitializationSettings7Vtbl {
	return (*IMediaCaptureInitializationSettings7Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCaptureInitializationSettings7) Get_DeviceUriPasswordCredential() *IPasswordCredential {
	var _result *IPasswordCredential
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceUriPasswordCredential, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCaptureInitializationSettings7) Put_DeviceUriPasswordCredential(value *IPasswordCredential) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DeviceUriPasswordCredential, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaCaptureInitializationSettings7) Get_DeviceUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCaptureInitializationSettings7) Put_DeviceUri(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DeviceUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// AEC47CA3-4477-4B04-A06F-2C1C5182FE9D
var IID_IMediaCapturePauseResult = syscall.GUID{0xAEC47CA3, 0x4477, 0x4B04,
	[8]byte{0xA0, 0x6F, 0x2C, 0x1C, 0x51, 0x82, 0xFE, 0x9D}}

type IMediaCapturePauseResultInterface interface {
	win32.IInspectableInterface
	Get_LastFrame() *IVideoFrame
	Get_RecordDuration() TimeSpan
}

type IMediaCapturePauseResultVtbl struct {
	win32.IInspectableVtbl
	Get_LastFrame      uintptr
	Get_RecordDuration uintptr
}

type IMediaCapturePauseResult struct {
	win32.IInspectable
}

func (this *IMediaCapturePauseResult) Vtbl() *IMediaCapturePauseResultVtbl {
	return (*IMediaCapturePauseResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCapturePauseResult) Get_LastFrame() *IVideoFrame {
	var _result *IVideoFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LastFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCapturePauseResult) Get_RecordDuration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RecordDuration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7D896566-04BE-5B89-B30E-BD34A9F12DB0
var IID_IMediaCaptureRelativePanelWatcher = syscall.GUID{0x7D896566, 0x04BE, 0x5B89,
	[8]byte{0xB3, 0x0E, 0xBD, 0x34, 0xA9, 0xF1, 0x2D, 0xB0}}

type IMediaCaptureRelativePanelWatcherInterface interface {
	win32.IInspectableInterface
	Get_RelativePanel() Panel
	Add_Changed(handler TypedEventHandler[*IMediaCaptureRelativePanelWatcher, interface{}]) EventRegistrationToken
	Remove_Changed(token EventRegistrationToken)
	Start()
	Stop()
}

type IMediaCaptureRelativePanelWatcherVtbl struct {
	win32.IInspectableVtbl
	Get_RelativePanel uintptr
	Add_Changed       uintptr
	Remove_Changed    uintptr
	Start             uintptr
	Stop              uintptr
}

type IMediaCaptureRelativePanelWatcher struct {
	win32.IInspectable
}

func (this *IMediaCaptureRelativePanelWatcher) Vtbl() *IMediaCaptureRelativePanelWatcherVtbl {
	return (*IMediaCaptureRelativePanelWatcherVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCaptureRelativePanelWatcher) Get_RelativePanel() Panel {
	var _result Panel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RelativePanel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCaptureRelativePanelWatcher) Add_Changed(handler TypedEventHandler[*IMediaCaptureRelativePanelWatcher, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Changed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCaptureRelativePanelWatcher) Remove_Changed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Changed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaCaptureRelativePanelWatcher) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IMediaCaptureRelativePanelWatcher) Stop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 1D83AAFE-6D45-4477-8DC4-AC5BC01C4091
var IID_IMediaCaptureSettings = syscall.GUID{0x1D83AAFE, 0x6D45, 0x4477,
	[8]byte{0x8D, 0xC4, 0xAC, 0x5B, 0xC0, 0x1C, 0x40, 0x91}}

type IMediaCaptureSettingsInterface interface {
	win32.IInspectableInterface
	Get_AudioDeviceId() string
	Get_VideoDeviceId() string
	Get_StreamingCaptureMode() StreamingCaptureMode
	Get_PhotoCaptureSource() PhotoCaptureSource
	Get_VideoDeviceCharacteristic() VideoDeviceCharacteristic
}

type IMediaCaptureSettingsVtbl struct {
	win32.IInspectableVtbl
	Get_AudioDeviceId             uintptr
	Get_VideoDeviceId             uintptr
	Get_StreamingCaptureMode      uintptr
	Get_PhotoCaptureSource        uintptr
	Get_VideoDeviceCharacteristic uintptr
}

type IMediaCaptureSettings struct {
	win32.IInspectable
}

func (this *IMediaCaptureSettings) Vtbl() *IMediaCaptureSettingsVtbl {
	return (*IMediaCaptureSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCaptureSettings) Get_AudioDeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioDeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaCaptureSettings) Get_VideoDeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoDeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaCaptureSettings) Get_StreamingCaptureMode() StreamingCaptureMode {
	var _result StreamingCaptureMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StreamingCaptureMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCaptureSettings) Get_PhotoCaptureSource() PhotoCaptureSource {
	var _result PhotoCaptureSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhotoCaptureSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCaptureSettings) Get_VideoDeviceCharacteristic() VideoDeviceCharacteristic {
	var _result VideoDeviceCharacteristic
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoDeviceCharacteristic, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 6F9E7CFB-FA9F-4B13-9CBE-5AB94F1F3493
var IID_IMediaCaptureSettings2 = syscall.GUID{0x6F9E7CFB, 0xFA9F, 0x4B13,
	[8]byte{0x9C, 0xBE, 0x5A, 0xB9, 0x4F, 0x1F, 0x34, 0x93}}

type IMediaCaptureSettings2Interface interface {
	win32.IInspectableInterface
	Get_ConcurrentRecordAndPhotoSupported() bool
	Get_ConcurrentRecordAndPhotoSequenceSupported() bool
	Get_CameraSoundRequiredForRegion() bool
	Get_Horizontal35mmEquivalentFocalLength() *IReference[uint32]
	Get_PitchOffsetDegrees() *IReference[int32]
	Get_Vertical35mmEquivalentFocalLength() *IReference[uint32]
	Get_MediaCategory() MediaCategory
	Get_AudioProcessing() AudioProcessing
}

type IMediaCaptureSettings2Vtbl struct {
	win32.IInspectableVtbl
	Get_ConcurrentRecordAndPhotoSupported         uintptr
	Get_ConcurrentRecordAndPhotoSequenceSupported uintptr
	Get_CameraSoundRequiredForRegion              uintptr
	Get_Horizontal35mmEquivalentFocalLength       uintptr
	Get_PitchOffsetDegrees                        uintptr
	Get_Vertical35mmEquivalentFocalLength         uintptr
	Get_MediaCategory                             uintptr
	Get_AudioProcessing                           uintptr
}

type IMediaCaptureSettings2 struct {
	win32.IInspectable
}

func (this *IMediaCaptureSettings2) Vtbl() *IMediaCaptureSettings2Vtbl {
	return (*IMediaCaptureSettings2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCaptureSettings2) Get_ConcurrentRecordAndPhotoSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConcurrentRecordAndPhotoSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCaptureSettings2) Get_ConcurrentRecordAndPhotoSequenceSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConcurrentRecordAndPhotoSequenceSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCaptureSettings2) Get_CameraSoundRequiredForRegion() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CameraSoundRequiredForRegion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCaptureSettings2) Get_Horizontal35mmEquivalentFocalLength() *IReference[uint32] {
	var _result *IReference[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Horizontal35mmEquivalentFocalLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCaptureSettings2) Get_PitchOffsetDegrees() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PitchOffsetDegrees, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCaptureSettings2) Get_Vertical35mmEquivalentFocalLength() *IReference[uint32] {
	var _result *IReference[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Vertical35mmEquivalentFocalLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCaptureSettings2) Get_MediaCategory() MediaCategory {
	var _result MediaCategory
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaCategory, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCaptureSettings2) Get_AudioProcessing() AudioProcessing {
	var _result AudioProcessing
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioProcessing, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 303C67C2-8058-4B1B-B877-8C2EF3528440
var IID_IMediaCaptureSettings3 = syscall.GUID{0x303C67C2, 0x8058, 0x4B1B,
	[8]byte{0xB8, 0x77, 0x8C, 0x2E, 0xF3, 0x52, 0x84, 0x40}}

type IMediaCaptureSettings3Interface interface {
	win32.IInspectableInterface
	Get_Direct3D11Device() unsafe.Pointer
}

type IMediaCaptureSettings3Vtbl struct {
	win32.IInspectableVtbl
	Get_Direct3D11Device uintptr
}

type IMediaCaptureSettings3 struct {
	win32.IInspectable
}

func (this *IMediaCaptureSettings3) Vtbl() *IMediaCaptureSettings3Vtbl {
	return (*IMediaCaptureSettings3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCaptureSettings3) Get_Direct3D11Device() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Direct3D11Device, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// ACEF81FF-99ED-4645-965E-1925CFC63834
var IID_IMediaCaptureStatics = syscall.GUID{0xACEF81FF, 0x99ED, 0x4645,
	[8]byte{0x96, 0x5E, 0x19, 0x25, 0xCF, 0xC6, 0x38, 0x34}}

type IMediaCaptureStaticsInterface interface {
	win32.IInspectableInterface
	IsVideoProfileSupported(videoDeviceId string) bool
	FindAllVideoProfiles(videoDeviceId string) *IVectorView[*IMediaCaptureVideoProfile]
	FindConcurrentProfiles(videoDeviceId string) *IVectorView[*IMediaCaptureVideoProfile]
	FindKnownVideoProfiles(videoDeviceId string, name KnownVideoProfile) *IVectorView[*IMediaCaptureVideoProfile]
}

type IMediaCaptureStaticsVtbl struct {
	win32.IInspectableVtbl
	IsVideoProfileSupported uintptr
	FindAllVideoProfiles    uintptr
	FindConcurrentProfiles  uintptr
	FindKnownVideoProfiles  uintptr
}

type IMediaCaptureStatics struct {
	win32.IInspectable
}

func (this *IMediaCaptureStatics) Vtbl() *IMediaCaptureStaticsVtbl {
	return (*IMediaCaptureStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCaptureStatics) IsVideoProfileSupported(videoDeviceId string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsVideoProfileSupported, uintptr(unsafe.Pointer(this)), NewHStr(videoDeviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCaptureStatics) FindAllVideoProfiles(videoDeviceId string) *IVectorView[*IMediaCaptureVideoProfile] {
	var _result *IVectorView[*IMediaCaptureVideoProfile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllVideoProfiles, uintptr(unsafe.Pointer(this)), NewHStr(videoDeviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCaptureStatics) FindConcurrentProfiles(videoDeviceId string) *IVectorView[*IMediaCaptureVideoProfile] {
	var _result *IVectorView[*IMediaCaptureVideoProfile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindConcurrentProfiles, uintptr(unsafe.Pointer(this)), NewHStr(videoDeviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCaptureStatics) FindKnownVideoProfiles(videoDeviceId string, name KnownVideoProfile) *IVectorView[*IMediaCaptureVideoProfile] {
	var _result *IVectorView[*IMediaCaptureVideoProfile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindKnownVideoProfiles, uintptr(unsafe.Pointer(this)), NewHStr(videoDeviceId).Ptr, uintptr(name), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F9DB6A2A-A092-4AD1-97D4-F201F9D082DB
var IID_IMediaCaptureStopResult = syscall.GUID{0xF9DB6A2A, 0xA092, 0x4AD1,
	[8]byte{0x97, 0xD4, 0xF2, 0x01, 0xF9, 0xD0, 0x82, 0xDB}}

type IMediaCaptureStopResultInterface interface {
	win32.IInspectableInterface
	Get_LastFrame() *IVideoFrame
	Get_RecordDuration() TimeSpan
}

type IMediaCaptureStopResultVtbl struct {
	win32.IInspectableVtbl
	Get_LastFrame      uintptr
	Get_RecordDuration uintptr
}

type IMediaCaptureStopResult struct {
	win32.IInspectable
}

func (this *IMediaCaptureStopResult) Vtbl() *IMediaCaptureStopResultVtbl {
	return (*IMediaCaptureStopResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCaptureStopResult) Get_LastFrame() *IVideoFrame {
	var _result *IVideoFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LastFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCaptureStopResult) Get_RecordDuration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RecordDuration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 27727073-549E-447F-A20A-4F03C479D8C0
var IID_IMediaCaptureVideoPreview = syscall.GUID{0x27727073, 0x549E, 0x447F,
	[8]byte{0xA2, 0x0A, 0x4F, 0x03, 0xC4, 0x79, 0xD8, 0xC0}}

type IMediaCaptureVideoPreviewInterface interface {
	win32.IInspectableInterface
	StartPreviewAsync() *IAsyncAction
	StartPreviewToCustomSinkAsync(encodingProfile *IMediaEncodingProfile, customMediaSink *IMediaExtension) *IAsyncAction
	StartPreviewToCustomSinkIdAsync(encodingProfile *IMediaEncodingProfile, customSinkActivationId string, customSinkSettings *IPropertySet) *IAsyncAction
	StopPreviewAsync() *IAsyncAction
}

type IMediaCaptureVideoPreviewVtbl struct {
	win32.IInspectableVtbl
	StartPreviewAsync               uintptr
	StartPreviewToCustomSinkAsync   uintptr
	StartPreviewToCustomSinkIdAsync uintptr
	StopPreviewAsync                uintptr
}

type IMediaCaptureVideoPreview struct {
	win32.IInspectable
}

func (this *IMediaCaptureVideoPreview) Vtbl() *IMediaCaptureVideoPreviewVtbl {
	return (*IMediaCaptureVideoPreviewVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCaptureVideoPreview) StartPreviewAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartPreviewAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCaptureVideoPreview) StartPreviewToCustomSinkAsync(encodingProfile *IMediaEncodingProfile, customMediaSink *IMediaExtension) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartPreviewToCustomSinkAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(encodingProfile)), uintptr(unsafe.Pointer(customMediaSink)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCaptureVideoPreview) StartPreviewToCustomSinkIdAsync(encodingProfile *IMediaEncodingProfile, customSinkActivationId string, customSinkSettings *IPropertySet) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartPreviewToCustomSinkIdAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(encodingProfile)), NewHStr(customSinkActivationId).Ptr, uintptr(unsafe.Pointer(customSinkSettings)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCaptureVideoPreview) StopPreviewAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopPreviewAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 21A073BF-A3EE-4ECF-9EF6-50B0BC4E1305
var IID_IMediaCaptureVideoProfile = syscall.GUID{0x21A073BF, 0xA3EE, 0x4ECF,
	[8]byte{0x9E, 0xF6, 0x50, 0xB0, 0xBC, 0x4E, 0x13, 0x05}}

type IMediaCaptureVideoProfileInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	Get_VideoDeviceId() string
	Get_SupportedPreviewMediaDescription() *IVectorView[*IMediaCaptureVideoProfileMediaDescription]
	Get_SupportedRecordMediaDescription() *IVectorView[*IMediaCaptureVideoProfileMediaDescription]
	Get_SupportedPhotoMediaDescription() *IVectorView[*IMediaCaptureVideoProfileMediaDescription]
	GetConcurrency() *IVectorView[*IMediaCaptureVideoProfile]
}

type IMediaCaptureVideoProfileVtbl struct {
	win32.IInspectableVtbl
	Get_Id                               uintptr
	Get_VideoDeviceId                    uintptr
	Get_SupportedPreviewMediaDescription uintptr
	Get_SupportedRecordMediaDescription  uintptr
	Get_SupportedPhotoMediaDescription   uintptr
	GetConcurrency                       uintptr
}

type IMediaCaptureVideoProfile struct {
	win32.IInspectable
}

func (this *IMediaCaptureVideoProfile) Vtbl() *IMediaCaptureVideoProfileVtbl {
	return (*IMediaCaptureVideoProfileVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCaptureVideoProfile) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaCaptureVideoProfile) Get_VideoDeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoDeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaCaptureVideoProfile) Get_SupportedPreviewMediaDescription() *IVectorView[*IMediaCaptureVideoProfileMediaDescription] {
	var _result *IVectorView[*IMediaCaptureVideoProfileMediaDescription]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedPreviewMediaDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCaptureVideoProfile) Get_SupportedRecordMediaDescription() *IVectorView[*IMediaCaptureVideoProfileMediaDescription] {
	var _result *IVectorView[*IMediaCaptureVideoProfileMediaDescription]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedRecordMediaDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCaptureVideoProfile) Get_SupportedPhotoMediaDescription() *IVectorView[*IMediaCaptureVideoProfileMediaDescription] {
	var _result *IVectorView[*IMediaCaptureVideoProfileMediaDescription]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedPhotoMediaDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCaptureVideoProfile) GetConcurrency() *IVectorView[*IMediaCaptureVideoProfile] {
	var _result *IVectorView[*IMediaCaptureVideoProfile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetConcurrency, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 97DDC95F-94CE-468F-9316-FC5BC2638F6B
var IID_IMediaCaptureVideoProfile2 = syscall.GUID{0x97DDC95F, 0x94CE, 0x468F,
	[8]byte{0x93, 0x16, 0xFC, 0x5B, 0xC2, 0x63, 0x8F, 0x6B}}

type IMediaCaptureVideoProfile2Interface interface {
	win32.IInspectableInterface
	Get_FrameSourceInfos() *IVectorView[*IMediaFrameSourceInfo]
	Get_Properties() *IMapView[syscall.GUID, interface{}]
}

type IMediaCaptureVideoProfile2Vtbl struct {
	win32.IInspectableVtbl
	Get_FrameSourceInfos uintptr
	Get_Properties       uintptr
}

type IMediaCaptureVideoProfile2 struct {
	win32.IInspectable
}

func (this *IMediaCaptureVideoProfile2) Vtbl() *IMediaCaptureVideoProfile2Vtbl {
	return (*IMediaCaptureVideoProfile2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCaptureVideoProfile2) Get_FrameSourceInfos() *IVectorView[*IMediaFrameSourceInfo] {
	var _result *IVectorView[*IMediaFrameSourceInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameSourceInfos, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaCaptureVideoProfile2) Get_Properties() *IMapView[syscall.GUID, interface{}] {
	var _result *IMapView[syscall.GUID, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8012AFEF-B691-49FF-83F2-C1E76EAAEA1B
var IID_IMediaCaptureVideoProfileMediaDescription = syscall.GUID{0x8012AFEF, 0xB691, 0x49FF,
	[8]byte{0x83, 0xF2, 0xC1, 0xE7, 0x6E, 0xAA, 0xEA, 0x1B}}

type IMediaCaptureVideoProfileMediaDescriptionInterface interface {
	win32.IInspectableInterface
	Get_Width() uint32
	Get_Height() uint32
	Get_FrameRate() float64
	Get_IsVariablePhotoSequenceSupported() bool
	Get_IsHdrVideoSupported() bool
}

type IMediaCaptureVideoProfileMediaDescriptionVtbl struct {
	win32.IInspectableVtbl
	Get_Width                            uintptr
	Get_Height                           uintptr
	Get_FrameRate                        uintptr
	Get_IsVariablePhotoSequenceSupported uintptr
	Get_IsHdrVideoSupported              uintptr
}

type IMediaCaptureVideoProfileMediaDescription struct {
	win32.IInspectable
}

func (this *IMediaCaptureVideoProfileMediaDescription) Vtbl() *IMediaCaptureVideoProfileMediaDescriptionVtbl {
	return (*IMediaCaptureVideoProfileMediaDescriptionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCaptureVideoProfileMediaDescription) Get_Width() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Width, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCaptureVideoProfileMediaDescription) Get_Height() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Height, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCaptureVideoProfileMediaDescription) Get_FrameRate() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCaptureVideoProfileMediaDescription) Get_IsVariablePhotoSequenceSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsVariablePhotoSequenceSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaCaptureVideoProfileMediaDescription) Get_IsHdrVideoSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsHdrVideoSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C6A6EF13-322D-413A-B85A-68A88E02F4E9
var IID_IMediaCaptureVideoProfileMediaDescription2 = syscall.GUID{0xC6A6EF13, 0x322D, 0x413A,
	[8]byte{0xB8, 0x5A, 0x68, 0xA8, 0x8E, 0x02, 0xF4, 0xE9}}

type IMediaCaptureVideoProfileMediaDescription2Interface interface {
	win32.IInspectableInterface
	Get_Subtype() string
	Get_Properties() *IMapView[syscall.GUID, interface{}]
}

type IMediaCaptureVideoProfileMediaDescription2Vtbl struct {
	win32.IInspectableVtbl
	Get_Subtype    uintptr
	Get_Properties uintptr
}

type IMediaCaptureVideoProfileMediaDescription2 struct {
	win32.IInspectable
}

func (this *IMediaCaptureVideoProfileMediaDescription2) Vtbl() *IMediaCaptureVideoProfileMediaDescription2Vtbl {
	return (*IMediaCaptureVideoProfileMediaDescription2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCaptureVideoProfileMediaDescription2) Get_Subtype() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Subtype, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaCaptureVideoProfileMediaDescription2) Get_Properties() *IMapView[syscall.GUID, interface{}] {
	var _result *IMapView[syscall.GUID, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 470F88B3-1E6D-4051-9C8B-F1D85AF047B7
var IID_IOptionalReferencePhotoCapturedEventArgs = syscall.GUID{0x470F88B3, 0x1E6D, 0x4051,
	[8]byte{0x9C, 0x8B, 0xF1, 0xD8, 0x5A, 0xF0, 0x47, 0xB7}}

type IOptionalReferencePhotoCapturedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Frame() *ICapturedFrame
	Get_Context() interface{}
}

type IOptionalReferencePhotoCapturedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Frame   uintptr
	Get_Context uintptr
}

type IOptionalReferencePhotoCapturedEventArgs struct {
	win32.IInspectable
}

func (this *IOptionalReferencePhotoCapturedEventArgs) Vtbl() *IOptionalReferencePhotoCapturedEventArgsVtbl {
	return (*IOptionalReferencePhotoCapturedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IOptionalReferencePhotoCapturedEventArgs) Get_Frame() *ICapturedFrame {
	var _result *ICapturedFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Frame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IOptionalReferencePhotoCapturedEventArgs) Get_Context() interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Context, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 373BFBC1-984E-4FF0-BF85-1C00AABC5A45
var IID_IPhotoCapturedEventArgs = syscall.GUID{0x373BFBC1, 0x984E, 0x4FF0,
	[8]byte{0xBF, 0x85, 0x1C, 0x00, 0xAA, 0xBC, 0x5A, 0x45}}

type IPhotoCapturedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Frame() *ICapturedFrame
	Get_Thumbnail() *ICapturedFrame
	Get_CaptureTimeOffset() TimeSpan
}

type IPhotoCapturedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Frame             uintptr
	Get_Thumbnail         uintptr
	Get_CaptureTimeOffset uintptr
}

type IPhotoCapturedEventArgs struct {
	win32.IInspectable
}

func (this *IPhotoCapturedEventArgs) Vtbl() *IPhotoCapturedEventArgsVtbl {
	return (*IPhotoCapturedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPhotoCapturedEventArgs) Get_Frame() *ICapturedFrame {
	var _result *ICapturedFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Frame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoCapturedEventArgs) Get_Thumbnail() *ICapturedFrame {
	var _result *ICapturedFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Thumbnail, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoCapturedEventArgs) Get_CaptureTimeOffset() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CaptureTimeOffset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// AB473672-C28A-4827-8F8D-3636D3BEB51E
var IID_IPhotoConfirmationCapturedEventArgs = syscall.GUID{0xAB473672, 0xC28A, 0x4827,
	[8]byte{0x8F, 0x8D, 0x36, 0x36, 0xD3, 0xBE, 0xB5, 0x1E}}

type IPhotoConfirmationCapturedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Frame() *ICapturedFrame
	Get_CaptureTimeOffset() TimeSpan
}

type IPhotoConfirmationCapturedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Frame             uintptr
	Get_CaptureTimeOffset uintptr
}

type IPhotoConfirmationCapturedEventArgs struct {
	win32.IInspectable
}

func (this *IPhotoConfirmationCapturedEventArgs) Vtbl() *IPhotoConfirmationCapturedEventArgsVtbl {
	return (*IPhotoConfirmationCapturedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPhotoConfirmationCapturedEventArgs) Get_Frame() *ICapturedFrame {
	var _result *ICapturedFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Frame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoConfirmationCapturedEventArgs) Get_CaptureTimeOffset() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CaptureTimeOffset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 89179EF7-CD12-4E0E-A6D4-5B3DE98B2E9B
var IID_IScreenCapture = syscall.GUID{0x89179EF7, 0xCD12, 0x4E0E,
	[8]byte{0xA6, 0xD4, 0x5B, 0x3D, 0xE9, 0x8B, 0x2E, 0x9B}}

type IScreenCaptureInterface interface {
	win32.IInspectableInterface
	Get_AudioSource() *IMediaSource
	Get_VideoSource() *IMediaSource
	Get_IsAudioSuspended() bool
	Get_IsVideoSuspended() bool
	Add_SourceSuspensionChanged(handler TypedEventHandler[*IScreenCapture, *ISourceSuspensionChangedEventArgs]) EventRegistrationToken
	Remove_SourceSuspensionChanged(token EventRegistrationToken)
}

type IScreenCaptureVtbl struct {
	win32.IInspectableVtbl
	Get_AudioSource                uintptr
	Get_VideoSource                uintptr
	Get_IsAudioSuspended           uintptr
	Get_IsVideoSuspended           uintptr
	Add_SourceSuspensionChanged    uintptr
	Remove_SourceSuspensionChanged uintptr
}

type IScreenCapture struct {
	win32.IInspectable
}

func (this *IScreenCapture) Vtbl() *IScreenCaptureVtbl {
	return (*IScreenCaptureVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IScreenCapture) Get_AudioSource() *IMediaSource {
	var _result *IMediaSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IScreenCapture) Get_VideoSource() *IMediaSource {
	var _result *IMediaSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IScreenCapture) Get_IsAudioSuspended() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsAudioSuspended, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IScreenCapture) Get_IsVideoSuspended() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsVideoSuspended, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IScreenCapture) Add_SourceSuspensionChanged(handler TypedEventHandler[*IScreenCapture, *ISourceSuspensionChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SourceSuspensionChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IScreenCapture) Remove_SourceSuspensionChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SourceSuspensionChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// C898C3B0-C8A5-11E2-8B8B-0800200C9A66
var IID_IScreenCaptureStatics = syscall.GUID{0xC898C3B0, 0xC8A5, 0x11E2,
	[8]byte{0x8B, 0x8B, 0x08, 0x00, 0x20, 0x0C, 0x9A, 0x66}}

type IScreenCaptureStaticsInterface interface {
	win32.IInspectableInterface
	GetForCurrentView() *IScreenCapture
}

type IScreenCaptureStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForCurrentView uintptr
}

type IScreenCaptureStatics struct {
	win32.IInspectable
}

func (this *IScreenCaptureStatics) Vtbl() *IScreenCaptureStaticsVtbl {
	return (*IScreenCaptureStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IScreenCaptureStatics) GetForCurrentView() *IScreenCapture {
	var _result *IScreenCapture
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2ECE7B5E-D49B-4394-BC32-F97D6CEDEC1C
var IID_ISourceSuspensionChangedEventArgs = syscall.GUID{0x2ECE7B5E, 0xD49B, 0x4394,
	[8]byte{0xBC, 0x32, 0xF9, 0x7D, 0x6C, 0xED, 0xEC, 0x1C}}

type ISourceSuspensionChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_IsAudioSuspended() bool
	Get_IsVideoSuspended() bool
}

type ISourceSuspensionChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_IsAudioSuspended uintptr
	Get_IsVideoSuspended uintptr
}

type ISourceSuspensionChangedEventArgs struct {
	win32.IInspectable
}

func (this *ISourceSuspensionChangedEventArgs) Vtbl() *ISourceSuspensionChangedEventArgsVtbl {
	return (*ISourceSuspensionChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISourceSuspensionChangedEventArgs) Get_IsAudioSuspended() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsAudioSuspended, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISourceSuspensionChangedEventArgs) Get_IsVideoSuspended() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsVideoSuspended, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D8770A6F-4390-4B5E-AD3E-0F8AF0963490
var IID_IVideoStreamConfiguration = syscall.GUID{0xD8770A6F, 0x4390, 0x4B5E,
	[8]byte{0xAD, 0x3E, 0x0F, 0x8A, 0xF0, 0x96, 0x34, 0x90}}

type IVideoStreamConfigurationInterface interface {
	win32.IInspectableInterface
	Get_InputProperties() *IVideoEncodingProperties
	Get_OutputProperties() *IVideoEncodingProperties
}

type IVideoStreamConfigurationVtbl struct {
	win32.IInspectableVtbl
	Get_InputProperties  uintptr
	Get_OutputProperties uintptr
}

type IVideoStreamConfiguration struct {
	win32.IInspectable
}

func (this *IVideoStreamConfiguration) Vtbl() *IVideoStreamConfigurationVtbl {
	return (*IVideoStreamConfigurationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoStreamConfiguration) Get_InputProperties() *IVideoEncodingProperties {
	var _result *IVideoEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InputProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoStreamConfiguration) Get_OutputProperties() *IVideoEncodingProperties {
	var _result *IVideoEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutputProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type AdvancedCapturedPhoto struct {
	RtClass
	*IAdvancedCapturedPhoto
}

type AdvancedPhotoCapture struct {
	RtClass
	*IAdvancedPhotoCapture
}

type AppCaptureMetadataWriter struct {
	RtClass
	*IAppCaptureMetadataWriter
}

func NewAppCaptureMetadataWriter() *AppCaptureMetadataWriter {
	hs := NewHStr("Windows.Media.Capture.AppCaptureMetadataWriter")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &AppCaptureMetadataWriter{
		RtClass:                   RtClass{PInspect: p},
		IAppCaptureMetadataWriter: (*IAppCaptureMetadataWriter)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type CapturedFrame struct {
	RtClass
	*ICapturedFrame
}

type CapturedFrameControlValues struct {
	RtClass
	*ICapturedFrameControlValues
}

type CapturedPhoto struct {
	RtClass
	*ICapturedPhoto
}

type LowLagMediaRecording struct {
	RtClass
	*ILowLagMediaRecording
}

type LowLagPhotoCapture struct {
	RtClass
	*ILowLagPhotoCapture
}

type LowLagPhotoSequenceCapture struct {
	RtClass
	*ILowLagPhotoSequenceCapture
}

type MediaCapture struct {
	RtClass
	*IMediaCapture
}

func NewMediaCapture() *MediaCapture {
	hs := NewHStr("Windows.Media.Capture.MediaCapture")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &MediaCapture{
		RtClass:       RtClass{PInspect: p},
		IMediaCapture: (*IMediaCapture)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewIMediaCaptureStatics() *IMediaCaptureStatics {
	var p *IMediaCaptureStatics
	hs := NewHStr("Windows.Media.Capture.MediaCapture")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMediaCaptureStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type MediaCaptureDeviceExclusiveControlStatusChangedEventArgs struct {
	RtClass
	*IMediaCaptureDeviceExclusiveControlStatusChangedEventArgs
}

type MediaCaptureFocusChangedEventArgs struct {
	RtClass
	*IMediaCaptureFocusChangedEventArgs
}

type MediaCaptureInitializationSettings struct {
	RtClass
	*IMediaCaptureInitializationSettings
}

func NewMediaCaptureInitializationSettings() *MediaCaptureInitializationSettings {
	hs := NewHStr("Windows.Media.Capture.MediaCaptureInitializationSettings")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &MediaCaptureInitializationSettings{
		RtClass:                             RtClass{PInspect: p},
		IMediaCaptureInitializationSettings: (*IMediaCaptureInitializationSettings)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type MediaCapturePauseResult struct {
	RtClass
	*IMediaCapturePauseResult
}

type MediaCaptureRelativePanelWatcher struct {
	RtClass
	*IMediaCaptureRelativePanelWatcher
}

type MediaCaptureSettings struct {
	RtClass
	*IMediaCaptureSettings
}

type MediaCaptureStopResult struct {
	RtClass
	*IMediaCaptureStopResult
}

type MediaCaptureVideoProfile struct {
	RtClass
	*IMediaCaptureVideoProfile
}

type MediaCaptureVideoProfileMediaDescription struct {
	RtClass
	*IMediaCaptureVideoProfileMediaDescription
}

type OptionalReferencePhotoCapturedEventArgs struct {
	RtClass
	*IOptionalReferencePhotoCapturedEventArgs
}

type PhotoCapturedEventArgs struct {
	RtClass
	*IPhotoCapturedEventArgs
}

type PhotoConfirmationCapturedEventArgs struct {
	RtClass
	*IPhotoConfirmationCapturedEventArgs
}

type VideoStreamConfiguration struct {
	RtClass
	*IVideoStreamConfiguration
}
