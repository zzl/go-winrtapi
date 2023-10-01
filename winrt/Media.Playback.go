package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"log"
	"syscall"
	"unsafe"
)

// enums

// enum
type AutoLoadedDisplayPropertyKind int32

const (
	AutoLoadedDisplayPropertyKind_None         AutoLoadedDisplayPropertyKind = 0
	AutoLoadedDisplayPropertyKind_MusicOrVideo AutoLoadedDisplayPropertyKind = 1
	AutoLoadedDisplayPropertyKind_Music        AutoLoadedDisplayPropertyKind = 2
	AutoLoadedDisplayPropertyKind_Video        AutoLoadedDisplayPropertyKind = 3
)

// enum
type FailedMediaStreamKind int32

const (
	FailedMediaStreamKind_Unknown FailedMediaStreamKind = 0
	FailedMediaStreamKind_Audio   FailedMediaStreamKind = 1
	FailedMediaStreamKind_Video   FailedMediaStreamKind = 2
)

// enum
type MediaBreakInsertionMethod int32

const (
	MediaBreakInsertionMethod_Interrupt MediaBreakInsertionMethod = 0
	MediaBreakInsertionMethod_Replace   MediaBreakInsertionMethod = 1
)

// enum
type MediaCommandEnablingRule int32

const (
	MediaCommandEnablingRule_Auto   MediaCommandEnablingRule = 0
	MediaCommandEnablingRule_Always MediaCommandEnablingRule = 1
	MediaCommandEnablingRule_Never  MediaCommandEnablingRule = 2
)

// enum
type MediaPlaybackItemChangedReason int32

const (
	MediaPlaybackItemChangedReason_InitialItem  MediaPlaybackItemChangedReason = 0
	MediaPlaybackItemChangedReason_EndOfStream  MediaPlaybackItemChangedReason = 1
	MediaPlaybackItemChangedReason_Error        MediaPlaybackItemChangedReason = 2
	MediaPlaybackItemChangedReason_AppRequested MediaPlaybackItemChangedReason = 3
)

// enum
type MediaPlaybackItemErrorCode int32

const (
	MediaPlaybackItemErrorCode_None                    MediaPlaybackItemErrorCode = 0
	MediaPlaybackItemErrorCode_Aborted                 MediaPlaybackItemErrorCode = 1
	MediaPlaybackItemErrorCode_NetworkError            MediaPlaybackItemErrorCode = 2
	MediaPlaybackItemErrorCode_DecodeError             MediaPlaybackItemErrorCode = 3
	MediaPlaybackItemErrorCode_SourceNotSupportedError MediaPlaybackItemErrorCode = 4
	MediaPlaybackItemErrorCode_EncryptionError         MediaPlaybackItemErrorCode = 5
)

// enum
type MediaPlaybackSessionVideoConstrictionReason int32

const (
	MediaPlaybackSessionVideoConstrictionReason_None                      MediaPlaybackSessionVideoConstrictionReason = 0
	MediaPlaybackSessionVideoConstrictionReason_VirtualMachine            MediaPlaybackSessionVideoConstrictionReason = 1
	MediaPlaybackSessionVideoConstrictionReason_UnsupportedDisplayAdapter MediaPlaybackSessionVideoConstrictionReason = 2
	MediaPlaybackSessionVideoConstrictionReason_UnsignedDriver            MediaPlaybackSessionVideoConstrictionReason = 3
	MediaPlaybackSessionVideoConstrictionReason_FrameServerEnabled        MediaPlaybackSessionVideoConstrictionReason = 4
	MediaPlaybackSessionVideoConstrictionReason_OutputProtectionFailed    MediaPlaybackSessionVideoConstrictionReason = 5
	MediaPlaybackSessionVideoConstrictionReason_Unknown                   MediaPlaybackSessionVideoConstrictionReason = 6
)

// enum
type MediaPlaybackState int32

const (
	MediaPlaybackState_None      MediaPlaybackState = 0
	MediaPlaybackState_Opening   MediaPlaybackState = 1
	MediaPlaybackState_Buffering MediaPlaybackState = 2
	MediaPlaybackState_Playing   MediaPlaybackState = 3
	MediaPlaybackState_Paused    MediaPlaybackState = 4
)

// enum
type MediaPlayerAudioCategory int32

const (
	MediaPlayerAudioCategory_Other          MediaPlayerAudioCategory = 0
	MediaPlayerAudioCategory_Communications MediaPlayerAudioCategory = 3
	MediaPlayerAudioCategory_Alerts         MediaPlayerAudioCategory = 4
	MediaPlayerAudioCategory_SoundEffects   MediaPlayerAudioCategory = 5
	MediaPlayerAudioCategory_GameEffects    MediaPlayerAudioCategory = 6
	MediaPlayerAudioCategory_GameMedia      MediaPlayerAudioCategory = 7
	MediaPlayerAudioCategory_GameChat       MediaPlayerAudioCategory = 8
	MediaPlayerAudioCategory_Speech         MediaPlayerAudioCategory = 9
	MediaPlayerAudioCategory_Movie          MediaPlayerAudioCategory = 10
	MediaPlayerAudioCategory_Media          MediaPlayerAudioCategory = 11
)

// enum
type MediaPlayerAudioDeviceType int32

const (
	MediaPlayerAudioDeviceType_Console        MediaPlayerAudioDeviceType = 0
	MediaPlayerAudioDeviceType_Multimedia     MediaPlayerAudioDeviceType = 1
	MediaPlayerAudioDeviceType_Communications MediaPlayerAudioDeviceType = 2
)

// enum
type MediaPlayerError int32

const (
	MediaPlayerError_Unknown            MediaPlayerError = 0
	MediaPlayerError_Aborted            MediaPlayerError = 1
	MediaPlayerError_NetworkError       MediaPlayerError = 2
	MediaPlayerError_DecodingError      MediaPlayerError = 3
	MediaPlayerError_SourceNotSupported MediaPlayerError = 4
)

// enum
type MediaPlayerState int32

const (
	MediaPlayerState_Closed    MediaPlayerState = 0
	MediaPlayerState_Opening   MediaPlayerState = 1
	MediaPlayerState_Buffering MediaPlayerState = 2
	MediaPlayerState_Playing   MediaPlayerState = 3
	MediaPlayerState_Paused    MediaPlayerState = 4
	MediaPlayerState_Stopped   MediaPlayerState = 5
)

// enum
type SphericalVideoProjectionMode int32

const (
	SphericalVideoProjectionMode_Spherical SphericalVideoProjectionMode = 0
	SphericalVideoProjectionMode_Flat      SphericalVideoProjectionMode = 1
)

// enum
type StereoscopicVideoRenderMode int32

const (
	StereoscopicVideoRenderMode_Mono   StereoscopicVideoRenderMode = 0
	StereoscopicVideoRenderMode_Stereo StereoscopicVideoRenderMode = 1
)

// enum
type TimedMetadataTrackPresentationMode int32

const (
	TimedMetadataTrackPresentationMode_Disabled             TimedMetadataTrackPresentationMode = 0
	TimedMetadataTrackPresentationMode_Hidden               TimedMetadataTrackPresentationMode = 1
	TimedMetadataTrackPresentationMode_ApplicationPresented TimedMetadataTrackPresentationMode = 2
	TimedMetadataTrackPresentationMode_PlatformPresented    TimedMetadataTrackPresentationMode = 3
)

// interfaces

// 856DDBC1-55F7-471F-A0F2-68AC4C904592
var IID_IBackgroundMediaPlayerStatics = syscall.GUID{0x856DDBC1, 0x55F7, 0x471F,
	[8]byte{0xA0, 0xF2, 0x68, 0xAC, 0x4C, 0x90, 0x45, 0x92}}

type IBackgroundMediaPlayerStaticsInterface interface {
	win32.IInspectableInterface
	Get_Current() *IMediaPlayer
	Add_MessageReceivedFromBackground(value EventHandler[*IMediaPlayerDataReceivedEventArgs]) EventRegistrationToken
	Remove_MessageReceivedFromBackground(token EventRegistrationToken)
	Add_MessageReceivedFromForeground(value EventHandler[*IMediaPlayerDataReceivedEventArgs]) EventRegistrationToken
	Remove_MessageReceivedFromForeground(token EventRegistrationToken)
	SendMessageToBackground(value *IPropertySet)
	SendMessageToForeground(value *IPropertySet)
	IsMediaPlaying() bool
	Shutdown()
}

type IBackgroundMediaPlayerStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Current                          uintptr
	Add_MessageReceivedFromBackground    uintptr
	Remove_MessageReceivedFromBackground uintptr
	Add_MessageReceivedFromForeground    uintptr
	Remove_MessageReceivedFromForeground uintptr
	SendMessageToBackground              uintptr
	SendMessageToForeground              uintptr
	IsMediaPlaying                       uintptr
	Shutdown                             uintptr
}

type IBackgroundMediaPlayerStatics struct {
	win32.IInspectable
}

func (this *IBackgroundMediaPlayerStatics) Vtbl() *IBackgroundMediaPlayerStaticsVtbl {
	return (*IBackgroundMediaPlayerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundMediaPlayerStatics) Get_Current() *IMediaPlayer {
	var _result *IMediaPlayer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Current, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundMediaPlayerStatics) Add_MessageReceivedFromBackground(value EventHandler[*IMediaPlayerDataReceivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_MessageReceivedFromBackground, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundMediaPlayerStatics) Remove_MessageReceivedFromBackground(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_MessageReceivedFromBackground, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IBackgroundMediaPlayerStatics) Add_MessageReceivedFromForeground(value EventHandler[*IMediaPlayerDataReceivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_MessageReceivedFromForeground, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundMediaPlayerStatics) Remove_MessageReceivedFromForeground(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_MessageReceivedFromForeground, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IBackgroundMediaPlayerStatics) SendMessageToBackground(value *IPropertySet) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendMessageToBackground, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IBackgroundMediaPlayerStatics) SendMessageToForeground(value *IPropertySet) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendMessageToForeground, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IBackgroundMediaPlayerStatics) IsMediaPlaying() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsMediaPlaying, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundMediaPlayerStatics) Shutdown() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Shutdown, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 1743A892-5C43-4A15-967A-572D2D0F26C6
var IID_ICurrentMediaPlaybackItemChangedEventArgs = syscall.GUID{0x1743A892, 0x5C43, 0x4A15,
	[8]byte{0x96, 0x7A, 0x57, 0x2D, 0x2D, 0x0F, 0x26, 0xC6}}

type ICurrentMediaPlaybackItemChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_NewItem() *IMediaPlaybackItem
	Get_OldItem() *IMediaPlaybackItem
}

type ICurrentMediaPlaybackItemChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_NewItem uintptr
	Get_OldItem uintptr
}

type ICurrentMediaPlaybackItemChangedEventArgs struct {
	win32.IInspectable
}

func (this *ICurrentMediaPlaybackItemChangedEventArgs) Vtbl() *ICurrentMediaPlaybackItemChangedEventArgsVtbl {
	return (*ICurrentMediaPlaybackItemChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICurrentMediaPlaybackItemChangedEventArgs) Get_NewItem() *IMediaPlaybackItem {
	var _result *IMediaPlaybackItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NewItem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICurrentMediaPlaybackItemChangedEventArgs) Get_OldItem() *IMediaPlaybackItem {
	var _result *IMediaPlaybackItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OldItem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1D80A51E-996E-40A9-BE48-E66EC90B2B7D
var IID_ICurrentMediaPlaybackItemChangedEventArgs2 = syscall.GUID{0x1D80A51E, 0x996E, 0x40A9,
	[8]byte{0xBE, 0x48, 0xE6, 0x6E, 0xC9, 0x0B, 0x2B, 0x7D}}

type ICurrentMediaPlaybackItemChangedEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_Reason() MediaPlaybackItemChangedReason
}

type ICurrentMediaPlaybackItemChangedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_Reason uintptr
}

type ICurrentMediaPlaybackItemChangedEventArgs2 struct {
	win32.IInspectable
}

func (this *ICurrentMediaPlaybackItemChangedEventArgs2) Vtbl() *ICurrentMediaPlaybackItemChangedEventArgs2Vtbl {
	return (*ICurrentMediaPlaybackItemChangedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICurrentMediaPlaybackItemChangedEventArgs2) Get_Reason() MediaPlaybackItemChangedReason {
	var _result MediaPlaybackItemChangedReason
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Reason, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 714BE270-0DEF-4EBC-A489-6B34930E1558
var IID_IMediaBreak = syscall.GUID{0x714BE270, 0x0DEF, 0x4EBC,
	[8]byte{0xA4, 0x89, 0x6B, 0x34, 0x93, 0x0E, 0x15, 0x58}}

type IMediaBreakInterface interface {
	win32.IInspectableInterface
	Get_PlaybackList() *IMediaPlaybackList
	Get_PresentationPosition() *IReference[TimeSpan]
	Get_InsertionMethod() MediaBreakInsertionMethod
	Get_CustomProperties() *IPropertySet
	Get_CanStart() bool
	Put_CanStart(value bool)
}

type IMediaBreakVtbl struct {
	win32.IInspectableVtbl
	Get_PlaybackList         uintptr
	Get_PresentationPosition uintptr
	Get_InsertionMethod      uintptr
	Get_CustomProperties     uintptr
	Get_CanStart             uintptr
	Put_CanStart             uintptr
}

type IMediaBreak struct {
	win32.IInspectable
}

func (this *IMediaBreak) Vtbl() *IMediaBreakVtbl {
	return (*IMediaBreakVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaBreak) Get_PlaybackList() *IMediaPlaybackList {
	var _result *IMediaPlaybackList
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlaybackList, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaBreak) Get_PresentationPosition() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PresentationPosition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaBreak) Get_InsertionMethod() MediaBreakInsertionMethod {
	var _result MediaBreakInsertionMethod
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InsertionMethod, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaBreak) Get_CustomProperties() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CustomProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaBreak) Get_CanStart() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanStart, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaBreak) Put_CanStart(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CanStart, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 32B93276-1C5D-4FEE-8732-236DC3A88580
var IID_IMediaBreakEndedEventArgs = syscall.GUID{0x32B93276, 0x1C5D, 0x4FEE,
	[8]byte{0x87, 0x32, 0x23, 0x6D, 0xC3, 0xA8, 0x85, 0x80}}

type IMediaBreakEndedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_MediaBreak() *IMediaBreak
}

type IMediaBreakEndedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_MediaBreak uintptr
}

type IMediaBreakEndedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaBreakEndedEventArgs) Vtbl() *IMediaBreakEndedEventArgsVtbl {
	return (*IMediaBreakEndedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaBreakEndedEventArgs) Get_MediaBreak() *IMediaBreak {
	var _result *IMediaBreak
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaBreak, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4516E002-18E0-4079-8B5F-D33495C15D2E
var IID_IMediaBreakFactory = syscall.GUID{0x4516E002, 0x18E0, 0x4079,
	[8]byte{0x8B, 0x5F, 0xD3, 0x34, 0x95, 0xC1, 0x5D, 0x2E}}

type IMediaBreakFactoryInterface interface {
	win32.IInspectableInterface
	Create(insertionMethod MediaBreakInsertionMethod) *IMediaBreak
	CreateWithPresentationPosition(insertionMethod MediaBreakInsertionMethod, presentationPosition TimeSpan) *IMediaBreak
}

type IMediaBreakFactoryVtbl struct {
	win32.IInspectableVtbl
	Create                         uintptr
	CreateWithPresentationPosition uintptr
}

type IMediaBreakFactory struct {
	win32.IInspectable
}

func (this *IMediaBreakFactory) Vtbl() *IMediaBreakFactoryVtbl {
	return (*IMediaBreakFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaBreakFactory) Create(insertionMethod MediaBreakInsertionMethod) *IMediaBreak {
	var _result *IMediaBreak
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(insertionMethod), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaBreakFactory) CreateWithPresentationPosition(insertionMethod MediaBreakInsertionMethod, presentationPosition TimeSpan) *IMediaBreak {
	var _result *IMediaBreak
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithPresentationPosition, uintptr(unsafe.Pointer(this)), uintptr(insertionMethod), *(*uintptr)(unsafe.Pointer(&presentationPosition)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A854DDB1-FEB4-4D9B-9D97-0FDBE58E5E39
var IID_IMediaBreakManager = syscall.GUID{0xA854DDB1, 0xFEB4, 0x4D9B,
	[8]byte{0x9D, 0x97, 0x0F, 0xDB, 0xE5, 0x8E, 0x5E, 0x39}}

type IMediaBreakManagerInterface interface {
	win32.IInspectableInterface
	Add_BreaksSeekedOver(handler TypedEventHandler[*IMediaBreakManager, *IMediaBreakSeekedOverEventArgs]) EventRegistrationToken
	Remove_BreaksSeekedOver(token EventRegistrationToken)
	Add_BreakStarted(handler TypedEventHandler[*IMediaBreakManager, *IMediaBreakStartedEventArgs]) EventRegistrationToken
	Remove_BreakStarted(token EventRegistrationToken)
	Add_BreakEnded(handler TypedEventHandler[*IMediaBreakManager, *IMediaBreakEndedEventArgs]) EventRegistrationToken
	Remove_BreakEnded(token EventRegistrationToken)
	Add_BreakSkipped(handler TypedEventHandler[*IMediaBreakManager, *IMediaBreakSkippedEventArgs]) EventRegistrationToken
	Remove_BreakSkipped(token EventRegistrationToken)
	Get_CurrentBreak() *IMediaBreak
	Get_PlaybackSession() *IMediaPlaybackSession
	PlayBreak(value *IMediaBreak)
	SkipCurrentBreak()
}

type IMediaBreakManagerVtbl struct {
	win32.IInspectableVtbl
	Add_BreaksSeekedOver    uintptr
	Remove_BreaksSeekedOver uintptr
	Add_BreakStarted        uintptr
	Remove_BreakStarted     uintptr
	Add_BreakEnded          uintptr
	Remove_BreakEnded       uintptr
	Add_BreakSkipped        uintptr
	Remove_BreakSkipped     uintptr
	Get_CurrentBreak        uintptr
	Get_PlaybackSession     uintptr
	PlayBreak               uintptr
	SkipCurrentBreak        uintptr
}

type IMediaBreakManager struct {
	win32.IInspectable
}

func (this *IMediaBreakManager) Vtbl() *IMediaBreakManagerVtbl {
	return (*IMediaBreakManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaBreakManager) Add_BreaksSeekedOver(handler TypedEventHandler[*IMediaBreakManager, *IMediaBreakSeekedOverEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_BreaksSeekedOver, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaBreakManager) Remove_BreaksSeekedOver(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_BreaksSeekedOver, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaBreakManager) Add_BreakStarted(handler TypedEventHandler[*IMediaBreakManager, *IMediaBreakStartedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_BreakStarted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaBreakManager) Remove_BreakStarted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_BreakStarted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaBreakManager) Add_BreakEnded(handler TypedEventHandler[*IMediaBreakManager, *IMediaBreakEndedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_BreakEnded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaBreakManager) Remove_BreakEnded(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_BreakEnded, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaBreakManager) Add_BreakSkipped(handler TypedEventHandler[*IMediaBreakManager, *IMediaBreakSkippedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_BreakSkipped, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaBreakManager) Remove_BreakSkipped(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_BreakSkipped, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaBreakManager) Get_CurrentBreak() *IMediaBreak {
	var _result *IMediaBreak
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentBreak, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaBreakManager) Get_PlaybackSession() *IMediaPlaybackSession {
	var _result *IMediaPlaybackSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlaybackSession, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaBreakManager) PlayBreak(value *IMediaBreak) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PlayBreak, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaBreakManager) SkipCurrentBreak() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SkipCurrentBreak, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// A19A5813-98B6-41D8-83DA-F971D22B7BBA
var IID_IMediaBreakSchedule = syscall.GUID{0xA19A5813, 0x98B6, 0x41D8,
	[8]byte{0x83, 0xDA, 0xF9, 0x71, 0xD2, 0x2B, 0x7B, 0xBA}}

type IMediaBreakScheduleInterface interface {
	win32.IInspectableInterface
	Add_ScheduleChanged(handler TypedEventHandler[*IMediaBreakSchedule, interface{}]) EventRegistrationToken
	Remove_ScheduleChanged(token EventRegistrationToken)
	InsertMidrollBreak(mediaBreak *IMediaBreak)
	RemoveMidrollBreak(mediaBreak *IMediaBreak)
	Get_MidrollBreaks() *IVectorView[*IMediaBreak]
	Put_PrerollBreak(value *IMediaBreak)
	Get_PrerollBreak() *IMediaBreak
	Put_PostrollBreak(value *IMediaBreak)
	Get_PostrollBreak() *IMediaBreak
	Get_PlaybackItem() *IMediaPlaybackItem
}

type IMediaBreakScheduleVtbl struct {
	win32.IInspectableVtbl
	Add_ScheduleChanged    uintptr
	Remove_ScheduleChanged uintptr
	InsertMidrollBreak     uintptr
	RemoveMidrollBreak     uintptr
	Get_MidrollBreaks      uintptr
	Put_PrerollBreak       uintptr
	Get_PrerollBreak       uintptr
	Put_PostrollBreak      uintptr
	Get_PostrollBreak      uintptr
	Get_PlaybackItem       uintptr
}

type IMediaBreakSchedule struct {
	win32.IInspectable
}

func (this *IMediaBreakSchedule) Vtbl() *IMediaBreakScheduleVtbl {
	return (*IMediaBreakScheduleVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaBreakSchedule) Add_ScheduleChanged(handler TypedEventHandler[*IMediaBreakSchedule, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ScheduleChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaBreakSchedule) Remove_ScheduleChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ScheduleChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaBreakSchedule) InsertMidrollBreak(mediaBreak *IMediaBreak) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertMidrollBreak, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(mediaBreak)))
	_ = _hr
}

func (this *IMediaBreakSchedule) RemoveMidrollBreak(mediaBreak *IMediaBreak) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveMidrollBreak, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(mediaBreak)))
	_ = _hr
}

func (this *IMediaBreakSchedule) Get_MidrollBreaks() *IVectorView[*IMediaBreak] {
	var _result *IVectorView[*IMediaBreak]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MidrollBreaks, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaBreakSchedule) Put_PrerollBreak(value *IMediaBreak) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PrerollBreak, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaBreakSchedule) Get_PrerollBreak() *IMediaBreak {
	var _result *IMediaBreak
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PrerollBreak, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaBreakSchedule) Put_PostrollBreak(value *IMediaBreak) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PostrollBreak, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaBreakSchedule) Get_PostrollBreak() *IMediaBreak {
	var _result *IMediaBreak
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PostrollBreak, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaBreakSchedule) Get_PlaybackItem() *IMediaPlaybackItem {
	var _result *IMediaPlaybackItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlaybackItem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E5AA6746-0606-4492-B9D3-C3C8FDE0A4EA
var IID_IMediaBreakSeekedOverEventArgs = syscall.GUID{0xE5AA6746, 0x0606, 0x4492,
	[8]byte{0xB9, 0xD3, 0xC3, 0xC8, 0xFD, 0xE0, 0xA4, 0xEA}}

type IMediaBreakSeekedOverEventArgsInterface interface {
	win32.IInspectableInterface
	Get_SeekedOverBreaks() *IVectorView[*IMediaBreak]
	Get_OldPosition() TimeSpan
	Get_NewPosition() TimeSpan
}

type IMediaBreakSeekedOverEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_SeekedOverBreaks uintptr
	Get_OldPosition      uintptr
	Get_NewPosition      uintptr
}

type IMediaBreakSeekedOverEventArgs struct {
	win32.IInspectable
}

func (this *IMediaBreakSeekedOverEventArgs) Vtbl() *IMediaBreakSeekedOverEventArgsVtbl {
	return (*IMediaBreakSeekedOverEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaBreakSeekedOverEventArgs) Get_SeekedOverBreaks() *IVectorView[*IMediaBreak] {
	var _result *IVectorView[*IMediaBreak]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SeekedOverBreaks, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaBreakSeekedOverEventArgs) Get_OldPosition() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OldPosition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaBreakSeekedOverEventArgs) Get_NewPosition() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NewPosition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 6EE94C05-2F54-4A3E-A3AB-24C3B270B4A3
var IID_IMediaBreakSkippedEventArgs = syscall.GUID{0x6EE94C05, 0x2F54, 0x4A3E,
	[8]byte{0xA3, 0xAB, 0x24, 0xC3, 0xB2, 0x70, 0xB4, 0xA3}}

type IMediaBreakSkippedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_MediaBreak() *IMediaBreak
}

type IMediaBreakSkippedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_MediaBreak uintptr
}

type IMediaBreakSkippedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaBreakSkippedEventArgs) Vtbl() *IMediaBreakSkippedEventArgsVtbl {
	return (*IMediaBreakSkippedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaBreakSkippedEventArgs) Get_MediaBreak() *IMediaBreak {
	var _result *IMediaBreak
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaBreak, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A87EFE71-DFD4-454A-956E-0A4A648395F8
var IID_IMediaBreakStartedEventArgs = syscall.GUID{0xA87EFE71, 0xDFD4, 0x454A,
	[8]byte{0x95, 0x6E, 0x0A, 0x4A, 0x64, 0x83, 0x95, 0xF8}}

type IMediaBreakStartedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_MediaBreak() *IMediaBreak
}

type IMediaBreakStartedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_MediaBreak uintptr
}

type IMediaBreakStartedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaBreakStartedEventArgs) Vtbl() *IMediaBreakStartedEventArgsVtbl {
	return (*IMediaBreakStartedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaBreakStartedEventArgs) Get_MediaBreak() *IMediaBreak {
	var _result *IMediaBreak
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaBreak, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5C1D0BA7-3856-48B9-8DC6-244BF107BF8C
var IID_IMediaEnginePlaybackSource = syscall.GUID{0x5C1D0BA7, 0x3856, 0x48B9,
	[8]byte{0x8D, 0xC6, 0x24, 0x4B, 0xF1, 0x07, 0xBF, 0x8C}}

type IMediaEnginePlaybackSourceInterface interface {
	win32.IInspectableInterface
	Get_CurrentItem() *IMediaPlaybackItem
	SetPlaybackSource(source *IMediaPlaybackSource)
}

type IMediaEnginePlaybackSourceVtbl struct {
	win32.IInspectableVtbl
	Get_CurrentItem   uintptr
	SetPlaybackSource uintptr
}

type IMediaEnginePlaybackSource struct {
	win32.IInspectable
}

func (this *IMediaEnginePlaybackSource) Vtbl() *IMediaEnginePlaybackSourceVtbl {
	return (*IMediaEnginePlaybackSourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaEnginePlaybackSource) Get_CurrentItem() *IMediaPlaybackItem {
	var _result *IMediaPlaybackItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentItem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaEnginePlaybackSource) SetPlaybackSource(source *IMediaPlaybackSource) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetPlaybackSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(source)))
	_ = _hr
}

// 1E3C1B48-7097-4384-A217-C1291DFA8C16
var IID_IMediaItemDisplayProperties = syscall.GUID{0x1E3C1B48, 0x7097, 0x4384,
	[8]byte{0xA2, 0x17, 0xC1, 0x29, 0x1D, 0xFA, 0x8C, 0x16}}

type IMediaItemDisplayPropertiesInterface interface {
	win32.IInspectableInterface
	Get_Type() MediaPlaybackType
	Put_Type(value MediaPlaybackType)
	Get_MusicProperties() *IMusicDisplayProperties
	Get_VideoProperties() *IVideoDisplayProperties
	Get_Thumbnail() *IRandomAccessStreamReference
	Put_Thumbnail(value *IRandomAccessStreamReference)
	ClearAll()
}

type IMediaItemDisplayPropertiesVtbl struct {
	win32.IInspectableVtbl
	Get_Type            uintptr
	Put_Type            uintptr
	Get_MusicProperties uintptr
	Get_VideoProperties uintptr
	Get_Thumbnail       uintptr
	Put_Thumbnail       uintptr
	ClearAll            uintptr
}

type IMediaItemDisplayProperties struct {
	win32.IInspectable
}

func (this *IMediaItemDisplayProperties) Vtbl() *IMediaItemDisplayPropertiesVtbl {
	return (*IMediaItemDisplayPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaItemDisplayProperties) Get_Type() MediaPlaybackType {
	var _result MediaPlaybackType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Type, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaItemDisplayProperties) Put_Type(value MediaPlaybackType) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Type, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMediaItemDisplayProperties) Get_MusicProperties() *IMusicDisplayProperties {
	var _result *IMusicDisplayProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MusicProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaItemDisplayProperties) Get_VideoProperties() *IVideoDisplayProperties {
	var _result *IVideoDisplayProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaItemDisplayProperties) Get_Thumbnail() *IRandomAccessStreamReference {
	var _result *IRandomAccessStreamReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Thumbnail, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaItemDisplayProperties) Put_Thumbnail(value *IRandomAccessStreamReference) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Thumbnail, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaItemDisplayProperties) ClearAll() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearAll, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 5ACEE5A6-5CB6-4A5A-8521-CC86B1C1ED37
var IID_IMediaPlaybackCommandManager = syscall.GUID{0x5ACEE5A6, 0x5CB6, 0x4A5A,
	[8]byte{0x85, 0x21, 0xCC, 0x86, 0xB1, 0xC1, 0xED, 0x37}}

type IMediaPlaybackCommandManagerInterface interface {
	win32.IInspectableInterface
	Get_IsEnabled() bool
	Put_IsEnabled(value bool)
	Get_MediaPlayer() *IMediaPlayer
	Get_PlayBehavior() *IMediaPlaybackCommandManagerCommandBehavior
	Get_PauseBehavior() *IMediaPlaybackCommandManagerCommandBehavior
	Get_NextBehavior() *IMediaPlaybackCommandManagerCommandBehavior
	Get_PreviousBehavior() *IMediaPlaybackCommandManagerCommandBehavior
	Get_FastForwardBehavior() *IMediaPlaybackCommandManagerCommandBehavior
	Get_RewindBehavior() *IMediaPlaybackCommandManagerCommandBehavior
	Get_ShuffleBehavior() *IMediaPlaybackCommandManagerCommandBehavior
	Get_AutoRepeatModeBehavior() *IMediaPlaybackCommandManagerCommandBehavior
	Get_PositionBehavior() *IMediaPlaybackCommandManagerCommandBehavior
	Get_RateBehavior() *IMediaPlaybackCommandManagerCommandBehavior
	Add_PlayReceived(handler TypedEventHandler[*IMediaPlaybackCommandManager, *IMediaPlaybackCommandManagerPlayReceivedEventArgs]) EventRegistrationToken
	Remove_PlayReceived(token EventRegistrationToken)
	Add_PauseReceived(handler TypedEventHandler[*IMediaPlaybackCommandManager, *IMediaPlaybackCommandManagerPauseReceivedEventArgs]) EventRegistrationToken
	Remove_PauseReceived(token EventRegistrationToken)
	Add_NextReceived(handler TypedEventHandler[*IMediaPlaybackCommandManager, *IMediaPlaybackCommandManagerNextReceivedEventArgs]) EventRegistrationToken
	Remove_NextReceived(token EventRegistrationToken)
	Add_PreviousReceived(handler TypedEventHandler[*IMediaPlaybackCommandManager, *IMediaPlaybackCommandManagerPreviousReceivedEventArgs]) EventRegistrationToken
	Remove_PreviousReceived(token EventRegistrationToken)
	Add_FastForwardReceived(handler TypedEventHandler[*IMediaPlaybackCommandManager, *IMediaPlaybackCommandManagerFastForwardReceivedEventArgs]) EventRegistrationToken
	Remove_FastForwardReceived(token EventRegistrationToken)
	Add_RewindReceived(handler TypedEventHandler[*IMediaPlaybackCommandManager, *IMediaPlaybackCommandManagerRewindReceivedEventArgs]) EventRegistrationToken
	Remove_RewindReceived(token EventRegistrationToken)
	Add_ShuffleReceived(handler TypedEventHandler[*IMediaPlaybackCommandManager, *IMediaPlaybackCommandManagerShuffleReceivedEventArgs]) EventRegistrationToken
	Remove_ShuffleReceived(token EventRegistrationToken)
	Add_AutoRepeatModeReceived(handler TypedEventHandler[*IMediaPlaybackCommandManager, *IMediaPlaybackCommandManagerAutoRepeatModeReceivedEventArgs]) EventRegistrationToken
	Remove_AutoRepeatModeReceived(token EventRegistrationToken)
	Add_PositionReceived(handler TypedEventHandler[*IMediaPlaybackCommandManager, *IMediaPlaybackCommandManagerPositionReceivedEventArgs]) EventRegistrationToken
	Remove_PositionReceived(token EventRegistrationToken)
	Add_RateReceived(handler TypedEventHandler[*IMediaPlaybackCommandManager, *IMediaPlaybackCommandManagerRateReceivedEventArgs]) EventRegistrationToken
	Remove_RateReceived(token EventRegistrationToken)
}

type IMediaPlaybackCommandManagerVtbl struct {
	win32.IInspectableVtbl
	Get_IsEnabled                 uintptr
	Put_IsEnabled                 uintptr
	Get_MediaPlayer               uintptr
	Get_PlayBehavior              uintptr
	Get_PauseBehavior             uintptr
	Get_NextBehavior              uintptr
	Get_PreviousBehavior          uintptr
	Get_FastForwardBehavior       uintptr
	Get_RewindBehavior            uintptr
	Get_ShuffleBehavior           uintptr
	Get_AutoRepeatModeBehavior    uintptr
	Get_PositionBehavior          uintptr
	Get_RateBehavior              uintptr
	Add_PlayReceived              uintptr
	Remove_PlayReceived           uintptr
	Add_PauseReceived             uintptr
	Remove_PauseReceived          uintptr
	Add_NextReceived              uintptr
	Remove_NextReceived           uintptr
	Add_PreviousReceived          uintptr
	Remove_PreviousReceived       uintptr
	Add_FastForwardReceived       uintptr
	Remove_FastForwardReceived    uintptr
	Add_RewindReceived            uintptr
	Remove_RewindReceived         uintptr
	Add_ShuffleReceived           uintptr
	Remove_ShuffleReceived        uintptr
	Add_AutoRepeatModeReceived    uintptr
	Remove_AutoRepeatModeReceived uintptr
	Add_PositionReceived          uintptr
	Remove_PositionReceived       uintptr
	Add_RateReceived              uintptr
	Remove_RateReceived           uintptr
}

type IMediaPlaybackCommandManager struct {
	win32.IInspectable
}

func (this *IMediaPlaybackCommandManager) Vtbl() *IMediaPlaybackCommandManagerVtbl {
	return (*IMediaPlaybackCommandManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackCommandManager) Get_IsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackCommandManager) Put_IsEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaPlaybackCommandManager) Get_MediaPlayer() *IMediaPlayer {
	var _result *IMediaPlayer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaPlayer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackCommandManager) Get_PlayBehavior() *IMediaPlaybackCommandManagerCommandBehavior {
	var _result *IMediaPlaybackCommandManagerCommandBehavior
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlayBehavior, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackCommandManager) Get_PauseBehavior() *IMediaPlaybackCommandManagerCommandBehavior {
	var _result *IMediaPlaybackCommandManagerCommandBehavior
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PauseBehavior, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackCommandManager) Get_NextBehavior() *IMediaPlaybackCommandManagerCommandBehavior {
	var _result *IMediaPlaybackCommandManagerCommandBehavior
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NextBehavior, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackCommandManager) Get_PreviousBehavior() *IMediaPlaybackCommandManagerCommandBehavior {
	var _result *IMediaPlaybackCommandManagerCommandBehavior
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PreviousBehavior, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackCommandManager) Get_FastForwardBehavior() *IMediaPlaybackCommandManagerCommandBehavior {
	var _result *IMediaPlaybackCommandManagerCommandBehavior
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FastForwardBehavior, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackCommandManager) Get_RewindBehavior() *IMediaPlaybackCommandManagerCommandBehavior {
	var _result *IMediaPlaybackCommandManagerCommandBehavior
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RewindBehavior, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackCommandManager) Get_ShuffleBehavior() *IMediaPlaybackCommandManagerCommandBehavior {
	var _result *IMediaPlaybackCommandManagerCommandBehavior
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ShuffleBehavior, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackCommandManager) Get_AutoRepeatModeBehavior() *IMediaPlaybackCommandManagerCommandBehavior {
	var _result *IMediaPlaybackCommandManagerCommandBehavior
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AutoRepeatModeBehavior, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackCommandManager) Get_PositionBehavior() *IMediaPlaybackCommandManagerCommandBehavior {
	var _result *IMediaPlaybackCommandManagerCommandBehavior
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PositionBehavior, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackCommandManager) Get_RateBehavior() *IMediaPlaybackCommandManagerCommandBehavior {
	var _result *IMediaPlaybackCommandManagerCommandBehavior
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RateBehavior, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackCommandManager) Add_PlayReceived(handler TypedEventHandler[*IMediaPlaybackCommandManager, *IMediaPlaybackCommandManagerPlayReceivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PlayReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackCommandManager) Remove_PlayReceived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PlayReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackCommandManager) Add_PauseReceived(handler TypedEventHandler[*IMediaPlaybackCommandManager, *IMediaPlaybackCommandManagerPauseReceivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PauseReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackCommandManager) Remove_PauseReceived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PauseReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackCommandManager) Add_NextReceived(handler TypedEventHandler[*IMediaPlaybackCommandManager, *IMediaPlaybackCommandManagerNextReceivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_NextReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackCommandManager) Remove_NextReceived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_NextReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackCommandManager) Add_PreviousReceived(handler TypedEventHandler[*IMediaPlaybackCommandManager, *IMediaPlaybackCommandManagerPreviousReceivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PreviousReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackCommandManager) Remove_PreviousReceived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PreviousReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackCommandManager) Add_FastForwardReceived(handler TypedEventHandler[*IMediaPlaybackCommandManager, *IMediaPlaybackCommandManagerFastForwardReceivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_FastForwardReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackCommandManager) Remove_FastForwardReceived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_FastForwardReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackCommandManager) Add_RewindReceived(handler TypedEventHandler[*IMediaPlaybackCommandManager, *IMediaPlaybackCommandManagerRewindReceivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_RewindReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackCommandManager) Remove_RewindReceived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_RewindReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackCommandManager) Add_ShuffleReceived(handler TypedEventHandler[*IMediaPlaybackCommandManager, *IMediaPlaybackCommandManagerShuffleReceivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ShuffleReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackCommandManager) Remove_ShuffleReceived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ShuffleReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackCommandManager) Add_AutoRepeatModeReceived(handler TypedEventHandler[*IMediaPlaybackCommandManager, *IMediaPlaybackCommandManagerAutoRepeatModeReceivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AutoRepeatModeReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackCommandManager) Remove_AutoRepeatModeReceived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AutoRepeatModeReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackCommandManager) Add_PositionReceived(handler TypedEventHandler[*IMediaPlaybackCommandManager, *IMediaPlaybackCommandManagerPositionReceivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PositionReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackCommandManager) Remove_PositionReceived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PositionReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackCommandManager) Add_RateReceived(handler TypedEventHandler[*IMediaPlaybackCommandManager, *IMediaPlaybackCommandManagerRateReceivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_RateReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackCommandManager) Remove_RateReceived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_RateReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 3D6F4F23-5230-4411-A0E9-BAD94C2A045C
var IID_IMediaPlaybackCommandManagerAutoRepeatModeReceivedEventArgs = syscall.GUID{0x3D6F4F23, 0x5230, 0x4411,
	[8]byte{0xA0, 0xE9, 0xBA, 0xD9, 0x4C, 0x2A, 0x04, 0x5C}}

type IMediaPlaybackCommandManagerAutoRepeatModeReceivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Handled() bool
	Put_Handled(value bool)
	Get_AutoRepeatMode() MediaPlaybackAutoRepeatMode
	GetDeferral() *IDeferral
}

type IMediaPlaybackCommandManagerAutoRepeatModeReceivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Handled        uintptr
	Put_Handled        uintptr
	Get_AutoRepeatMode uintptr
	GetDeferral        uintptr
}

type IMediaPlaybackCommandManagerAutoRepeatModeReceivedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaPlaybackCommandManagerAutoRepeatModeReceivedEventArgs) Vtbl() *IMediaPlaybackCommandManagerAutoRepeatModeReceivedEventArgsVtbl {
	return (*IMediaPlaybackCommandManagerAutoRepeatModeReceivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackCommandManagerAutoRepeatModeReceivedEventArgs) Get_Handled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Handled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackCommandManagerAutoRepeatModeReceivedEventArgs) Put_Handled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Handled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaPlaybackCommandManagerAutoRepeatModeReceivedEventArgs) Get_AutoRepeatMode() MediaPlaybackAutoRepeatMode {
	var _result MediaPlaybackAutoRepeatMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AutoRepeatMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackCommandManagerAutoRepeatModeReceivedEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 786C1E78-CE78-4A10-AFD6-843FCBB90C2E
var IID_IMediaPlaybackCommandManagerCommandBehavior = syscall.GUID{0x786C1E78, 0xCE78, 0x4A10,
	[8]byte{0xAF, 0xD6, 0x84, 0x3F, 0xCB, 0xB9, 0x0C, 0x2E}}

type IMediaPlaybackCommandManagerCommandBehaviorInterface interface {
	win32.IInspectableInterface
	Get_CommandManager() *IMediaPlaybackCommandManager
	Get_IsEnabled() bool
	Get_EnablingRule() MediaCommandEnablingRule
	Put_EnablingRule(value MediaCommandEnablingRule)
	Add_IsEnabledChanged(handler TypedEventHandler[*IMediaPlaybackCommandManagerCommandBehavior, interface{}]) EventRegistrationToken
	Remove_IsEnabledChanged(token EventRegistrationToken)
}

type IMediaPlaybackCommandManagerCommandBehaviorVtbl struct {
	win32.IInspectableVtbl
	Get_CommandManager      uintptr
	Get_IsEnabled           uintptr
	Get_EnablingRule        uintptr
	Put_EnablingRule        uintptr
	Add_IsEnabledChanged    uintptr
	Remove_IsEnabledChanged uintptr
}

type IMediaPlaybackCommandManagerCommandBehavior struct {
	win32.IInspectable
}

func (this *IMediaPlaybackCommandManagerCommandBehavior) Vtbl() *IMediaPlaybackCommandManagerCommandBehaviorVtbl {
	return (*IMediaPlaybackCommandManagerCommandBehaviorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackCommandManagerCommandBehavior) Get_CommandManager() *IMediaPlaybackCommandManager {
	var _result *IMediaPlaybackCommandManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CommandManager, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackCommandManagerCommandBehavior) Get_IsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackCommandManagerCommandBehavior) Get_EnablingRule() MediaCommandEnablingRule {
	var _result MediaCommandEnablingRule
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EnablingRule, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackCommandManagerCommandBehavior) Put_EnablingRule(value MediaCommandEnablingRule) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_EnablingRule, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMediaPlaybackCommandManagerCommandBehavior) Add_IsEnabledChanged(handler TypedEventHandler[*IMediaPlaybackCommandManagerCommandBehavior, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_IsEnabledChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackCommandManagerCommandBehavior) Remove_IsEnabledChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_IsEnabledChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 30F064D9-B491-4D0A-BC21-3098BD1332E9
var IID_IMediaPlaybackCommandManagerFastForwardReceivedEventArgs = syscall.GUID{0x30F064D9, 0xB491, 0x4D0A,
	[8]byte{0xBC, 0x21, 0x30, 0x98, 0xBD, 0x13, 0x32, 0xE9}}

type IMediaPlaybackCommandManagerFastForwardReceivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Handled() bool
	Put_Handled(value bool)
	GetDeferral() *IDeferral
}

type IMediaPlaybackCommandManagerFastForwardReceivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Handled uintptr
	Put_Handled uintptr
	GetDeferral uintptr
}

type IMediaPlaybackCommandManagerFastForwardReceivedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaPlaybackCommandManagerFastForwardReceivedEventArgs) Vtbl() *IMediaPlaybackCommandManagerFastForwardReceivedEventArgsVtbl {
	return (*IMediaPlaybackCommandManagerFastForwardReceivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackCommandManagerFastForwardReceivedEventArgs) Get_Handled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Handled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackCommandManagerFastForwardReceivedEventArgs) Put_Handled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Handled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaPlaybackCommandManagerFastForwardReceivedEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E1504433-A2B0-45D4-B9DE-5F42AC14A839
var IID_IMediaPlaybackCommandManagerNextReceivedEventArgs = syscall.GUID{0xE1504433, 0xA2B0, 0x45D4,
	[8]byte{0xB9, 0xDE, 0x5F, 0x42, 0xAC, 0x14, 0xA8, 0x39}}

type IMediaPlaybackCommandManagerNextReceivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Handled() bool
	Put_Handled(value bool)
	GetDeferral() *IDeferral
}

type IMediaPlaybackCommandManagerNextReceivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Handled uintptr
	Put_Handled uintptr
	GetDeferral uintptr
}

type IMediaPlaybackCommandManagerNextReceivedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaPlaybackCommandManagerNextReceivedEventArgs) Vtbl() *IMediaPlaybackCommandManagerNextReceivedEventArgsVtbl {
	return (*IMediaPlaybackCommandManagerNextReceivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackCommandManagerNextReceivedEventArgs) Get_Handled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Handled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackCommandManagerNextReceivedEventArgs) Put_Handled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Handled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaPlaybackCommandManagerNextReceivedEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5CECCD1C-C25C-4221-B16C-C3C98CE012D6
var IID_IMediaPlaybackCommandManagerPauseReceivedEventArgs = syscall.GUID{0x5CECCD1C, 0xC25C, 0x4221,
	[8]byte{0xB1, 0x6C, 0xC3, 0xC9, 0x8C, 0xE0, 0x12, 0xD6}}

type IMediaPlaybackCommandManagerPauseReceivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Handled() bool
	Put_Handled(value bool)
	GetDeferral() *IDeferral
}

type IMediaPlaybackCommandManagerPauseReceivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Handled uintptr
	Put_Handled uintptr
	GetDeferral uintptr
}

type IMediaPlaybackCommandManagerPauseReceivedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaPlaybackCommandManagerPauseReceivedEventArgs) Vtbl() *IMediaPlaybackCommandManagerPauseReceivedEventArgsVtbl {
	return (*IMediaPlaybackCommandManagerPauseReceivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackCommandManagerPauseReceivedEventArgs) Get_Handled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Handled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackCommandManagerPauseReceivedEventArgs) Put_Handled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Handled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaPlaybackCommandManagerPauseReceivedEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9AF0004E-578B-4C56-A006-16159D888A48
var IID_IMediaPlaybackCommandManagerPlayReceivedEventArgs = syscall.GUID{0x9AF0004E, 0x578B, 0x4C56,
	[8]byte{0xA0, 0x06, 0x16, 0x15, 0x9D, 0x88, 0x8A, 0x48}}

type IMediaPlaybackCommandManagerPlayReceivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Handled() bool
	Put_Handled(value bool)
	GetDeferral() *IDeferral
}

type IMediaPlaybackCommandManagerPlayReceivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Handled uintptr
	Put_Handled uintptr
	GetDeferral uintptr
}

type IMediaPlaybackCommandManagerPlayReceivedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaPlaybackCommandManagerPlayReceivedEventArgs) Vtbl() *IMediaPlaybackCommandManagerPlayReceivedEventArgsVtbl {
	return (*IMediaPlaybackCommandManagerPlayReceivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackCommandManagerPlayReceivedEventArgs) Get_Handled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Handled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackCommandManagerPlayReceivedEventArgs) Put_Handled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Handled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaPlaybackCommandManagerPlayReceivedEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5591A754-D627-4BDD-A90D-86A015B24902
var IID_IMediaPlaybackCommandManagerPositionReceivedEventArgs = syscall.GUID{0x5591A754, 0xD627, 0x4BDD,
	[8]byte{0xA9, 0x0D, 0x86, 0xA0, 0x15, 0xB2, 0x49, 0x02}}

type IMediaPlaybackCommandManagerPositionReceivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Handled() bool
	Put_Handled(value bool)
	Get_Position() TimeSpan
	GetDeferral() *IDeferral
}

type IMediaPlaybackCommandManagerPositionReceivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Handled  uintptr
	Put_Handled  uintptr
	Get_Position uintptr
	GetDeferral  uintptr
}

type IMediaPlaybackCommandManagerPositionReceivedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaPlaybackCommandManagerPositionReceivedEventArgs) Vtbl() *IMediaPlaybackCommandManagerPositionReceivedEventArgsVtbl {
	return (*IMediaPlaybackCommandManagerPositionReceivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackCommandManagerPositionReceivedEventArgs) Get_Handled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Handled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackCommandManagerPositionReceivedEventArgs) Put_Handled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Handled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaPlaybackCommandManagerPositionReceivedEventArgs) Get_Position() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackCommandManagerPositionReceivedEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 525E3081-4632-4F76-99B1-D771623F6287
var IID_IMediaPlaybackCommandManagerPreviousReceivedEventArgs = syscall.GUID{0x525E3081, 0x4632, 0x4F76,
	[8]byte{0x99, 0xB1, 0xD7, 0x71, 0x62, 0x3F, 0x62, 0x87}}

type IMediaPlaybackCommandManagerPreviousReceivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Handled() bool
	Put_Handled(value bool)
	GetDeferral() *IDeferral
}

type IMediaPlaybackCommandManagerPreviousReceivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Handled uintptr
	Put_Handled uintptr
	GetDeferral uintptr
}

type IMediaPlaybackCommandManagerPreviousReceivedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaPlaybackCommandManagerPreviousReceivedEventArgs) Vtbl() *IMediaPlaybackCommandManagerPreviousReceivedEventArgsVtbl {
	return (*IMediaPlaybackCommandManagerPreviousReceivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackCommandManagerPreviousReceivedEventArgs) Get_Handled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Handled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackCommandManagerPreviousReceivedEventArgs) Put_Handled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Handled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaPlaybackCommandManagerPreviousReceivedEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 18EA3939-4A16-4169-8B05-3EB9F5FF78EB
var IID_IMediaPlaybackCommandManagerRateReceivedEventArgs = syscall.GUID{0x18EA3939, 0x4A16, 0x4169,
	[8]byte{0x8B, 0x05, 0x3E, 0xB9, 0xF5, 0xFF, 0x78, 0xEB}}

type IMediaPlaybackCommandManagerRateReceivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Handled() bool
	Put_Handled(value bool)
	Get_PlaybackRate() float64
	GetDeferral() *IDeferral
}

type IMediaPlaybackCommandManagerRateReceivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Handled      uintptr
	Put_Handled      uintptr
	Get_PlaybackRate uintptr
	GetDeferral      uintptr
}

type IMediaPlaybackCommandManagerRateReceivedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaPlaybackCommandManagerRateReceivedEventArgs) Vtbl() *IMediaPlaybackCommandManagerRateReceivedEventArgsVtbl {
	return (*IMediaPlaybackCommandManagerRateReceivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackCommandManagerRateReceivedEventArgs) Get_Handled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Handled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackCommandManagerRateReceivedEventArgs) Put_Handled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Handled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaPlaybackCommandManagerRateReceivedEventArgs) Get_PlaybackRate() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlaybackRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackCommandManagerRateReceivedEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9F085947-A3C0-425D-AAEF-97BA7898B141
var IID_IMediaPlaybackCommandManagerRewindReceivedEventArgs = syscall.GUID{0x9F085947, 0xA3C0, 0x425D,
	[8]byte{0xAA, 0xEF, 0x97, 0xBA, 0x78, 0x98, 0xB1, 0x41}}

type IMediaPlaybackCommandManagerRewindReceivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Handled() bool
	Put_Handled(value bool)
	GetDeferral() *IDeferral
}

type IMediaPlaybackCommandManagerRewindReceivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Handled uintptr
	Put_Handled uintptr
	GetDeferral uintptr
}

type IMediaPlaybackCommandManagerRewindReceivedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaPlaybackCommandManagerRewindReceivedEventArgs) Vtbl() *IMediaPlaybackCommandManagerRewindReceivedEventArgsVtbl {
	return (*IMediaPlaybackCommandManagerRewindReceivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackCommandManagerRewindReceivedEventArgs) Get_Handled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Handled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackCommandManagerRewindReceivedEventArgs) Put_Handled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Handled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaPlaybackCommandManagerRewindReceivedEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 50A05CEF-63EE-4A96-B7B5-FEE08B9FF90C
var IID_IMediaPlaybackCommandManagerShuffleReceivedEventArgs = syscall.GUID{0x50A05CEF, 0x63EE, 0x4A96,
	[8]byte{0xB7, 0xB5, 0xFE, 0xE0, 0x8B, 0x9F, 0xF9, 0x0C}}

type IMediaPlaybackCommandManagerShuffleReceivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Handled() bool
	Put_Handled(value bool)
	Get_IsShuffleRequested() bool
	GetDeferral() *IDeferral
}

type IMediaPlaybackCommandManagerShuffleReceivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Handled            uintptr
	Put_Handled            uintptr
	Get_IsShuffleRequested uintptr
	GetDeferral            uintptr
}

type IMediaPlaybackCommandManagerShuffleReceivedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaPlaybackCommandManagerShuffleReceivedEventArgs) Vtbl() *IMediaPlaybackCommandManagerShuffleReceivedEventArgsVtbl {
	return (*IMediaPlaybackCommandManagerShuffleReceivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackCommandManagerShuffleReceivedEventArgs) Get_Handled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Handled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackCommandManagerShuffleReceivedEventArgs) Put_Handled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Handled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaPlaybackCommandManagerShuffleReceivedEventArgs) Get_IsShuffleRequested() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsShuffleRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackCommandManagerShuffleReceivedEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 047097D2-E4AF-48AB-B283-6929E674ECE2
var IID_IMediaPlaybackItem = syscall.GUID{0x047097D2, 0xE4AF, 0x48AB,
	[8]byte{0xB2, 0x83, 0x69, 0x29, 0xE6, 0x74, 0xEC, 0xE2}}

type IMediaPlaybackItemInterface interface {
	win32.IInspectableInterface
	Add_AudioTracksChanged(handler TypedEventHandler[*IMediaPlaybackItem, *IVectorChangedEventArgs]) EventRegistrationToken
	Remove_AudioTracksChanged(token EventRegistrationToken)
	Add_VideoTracksChanged(handler TypedEventHandler[*IMediaPlaybackItem, *IVectorChangedEventArgs]) EventRegistrationToken
	Remove_VideoTracksChanged(token EventRegistrationToken)
	Add_TimedMetadataTracksChanged(handler TypedEventHandler[*IMediaPlaybackItem, *IVectorChangedEventArgs]) EventRegistrationToken
	Remove_TimedMetadataTracksChanged(token EventRegistrationToken)
	Get_Source() *IMediaSource2
	Get_AudioTracks() *IVectorView[*IMediaTrack]
	Get_VideoTracks() *IVectorView[*IMediaTrack]
	Get_TimedMetadataTracks() *IVectorView[*ITimedMetadataTrack]
}

type IMediaPlaybackItemVtbl struct {
	win32.IInspectableVtbl
	Add_AudioTracksChanged            uintptr
	Remove_AudioTracksChanged         uintptr
	Add_VideoTracksChanged            uintptr
	Remove_VideoTracksChanged         uintptr
	Add_TimedMetadataTracksChanged    uintptr
	Remove_TimedMetadataTracksChanged uintptr
	Get_Source                        uintptr
	Get_AudioTracks                   uintptr
	Get_VideoTracks                   uintptr
	Get_TimedMetadataTracks           uintptr
}

type IMediaPlaybackItem struct {
	win32.IInspectable
}

func (this *IMediaPlaybackItem) Vtbl() *IMediaPlaybackItemVtbl {
	return (*IMediaPlaybackItemVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackItem) Add_AudioTracksChanged(handler TypedEventHandler[*IMediaPlaybackItem, *IVectorChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AudioTracksChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackItem) Remove_AudioTracksChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AudioTracksChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackItem) Add_VideoTracksChanged(handler TypedEventHandler[*IMediaPlaybackItem, *IVectorChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_VideoTracksChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackItem) Remove_VideoTracksChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_VideoTracksChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackItem) Add_TimedMetadataTracksChanged(handler TypedEventHandler[*IMediaPlaybackItem, *IVectorChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_TimedMetadataTracksChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackItem) Remove_TimedMetadataTracksChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_TimedMetadataTracksChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackItem) Get_Source() *IMediaSource2 {
	var _result *IMediaSource2
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Source, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackItem) Get_AudioTracks() *IVectorView[*IMediaTrack] {
	var _result *IVectorView[*IMediaTrack]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioTracks, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackItem) Get_VideoTracks() *IVectorView[*IMediaTrack] {
	var _result *IVectorView[*IMediaTrack]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoTracks, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackItem) Get_TimedMetadataTracks() *IVectorView[*ITimedMetadataTrack] {
	var _result *IVectorView[*ITimedMetadataTrack]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TimedMetadataTracks, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D859D171-D7EF-4B81-AC1F-F40493CBB091
var IID_IMediaPlaybackItem2 = syscall.GUID{0xD859D171, 0xD7EF, 0x4B81,
	[8]byte{0xAC, 0x1F, 0xF4, 0x04, 0x93, 0xCB, 0xB0, 0x91}}

type IMediaPlaybackItem2Interface interface {
	win32.IInspectableInterface
	Get_BreakSchedule() *IMediaBreakSchedule
	Get_StartTime() TimeSpan
	Get_DurationLimit() *IReference[TimeSpan]
	Get_CanSkip() bool
	Put_CanSkip(value bool)
	GetDisplayProperties() *IMediaItemDisplayProperties
	ApplyDisplayProperties(value *IMediaItemDisplayProperties)
}

type IMediaPlaybackItem2Vtbl struct {
	win32.IInspectableVtbl
	Get_BreakSchedule      uintptr
	Get_StartTime          uintptr
	Get_DurationLimit      uintptr
	Get_CanSkip            uintptr
	Put_CanSkip            uintptr
	GetDisplayProperties   uintptr
	ApplyDisplayProperties uintptr
}

type IMediaPlaybackItem2 struct {
	win32.IInspectable
}

func (this *IMediaPlaybackItem2) Vtbl() *IMediaPlaybackItem2Vtbl {
	return (*IMediaPlaybackItem2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackItem2) Get_BreakSchedule() *IMediaBreakSchedule {
	var _result *IMediaBreakSchedule
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BreakSchedule, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackItem2) Get_StartTime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StartTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackItem2) Get_DurationLimit() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DurationLimit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackItem2) Get_CanSkip() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanSkip, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackItem2) Put_CanSkip(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CanSkip, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaPlaybackItem2) GetDisplayProperties() *IMediaItemDisplayProperties {
	var _result *IMediaItemDisplayProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDisplayProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackItem2) ApplyDisplayProperties(value *IMediaItemDisplayProperties) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ApplyDisplayProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 0D328220-B80A-4D09-9FF8-F87094A1C831
var IID_IMediaPlaybackItem3 = syscall.GUID{0x0D328220, 0xB80A, 0x4D09,
	[8]byte{0x9F, 0xF8, 0xF8, 0x70, 0x94, 0xA1, 0xC8, 0x31}}

type IMediaPlaybackItem3Interface interface {
	win32.IInspectableInterface
	Get_IsDisabledInPlaybackList() bool
	Put_IsDisabledInPlaybackList(value bool)
	Get_TotalDownloadProgress() float64
	Get_AutoLoadedDisplayProperties() AutoLoadedDisplayPropertyKind
	Put_AutoLoadedDisplayProperties(value AutoLoadedDisplayPropertyKind)
}

type IMediaPlaybackItem3Vtbl struct {
	win32.IInspectableVtbl
	Get_IsDisabledInPlaybackList    uintptr
	Put_IsDisabledInPlaybackList    uintptr
	Get_TotalDownloadProgress       uintptr
	Get_AutoLoadedDisplayProperties uintptr
	Put_AutoLoadedDisplayProperties uintptr
}

type IMediaPlaybackItem3 struct {
	win32.IInspectable
}

func (this *IMediaPlaybackItem3) Vtbl() *IMediaPlaybackItem3Vtbl {
	return (*IMediaPlaybackItem3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackItem3) Get_IsDisabledInPlaybackList() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDisabledInPlaybackList, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackItem3) Put_IsDisabledInPlaybackList(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsDisabledInPlaybackList, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaPlaybackItem3) Get_TotalDownloadProgress() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TotalDownloadProgress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackItem3) Get_AutoLoadedDisplayProperties() AutoLoadedDisplayPropertyKind {
	var _result AutoLoadedDisplayPropertyKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AutoLoadedDisplayProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackItem3) Put_AutoLoadedDisplayProperties(value AutoLoadedDisplayPropertyKind) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AutoLoadedDisplayProperties, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 69FBEF2B-DCD6-4DF9-A450-DBF4C6F1C2C2
var IID_IMediaPlaybackItemError = syscall.GUID{0x69FBEF2B, 0xDCD6, 0x4DF9,
	[8]byte{0xA4, 0x50, 0xDB, 0xF4, 0xC6, 0xF1, 0xC2, 0xC2}}

type IMediaPlaybackItemErrorInterface interface {
	win32.IInspectableInterface
	Get_ErrorCode() MediaPlaybackItemErrorCode
	Get_ExtendedError() HResult
}

type IMediaPlaybackItemErrorVtbl struct {
	win32.IInspectableVtbl
	Get_ErrorCode     uintptr
	Get_ExtendedError uintptr
}

type IMediaPlaybackItemError struct {
	win32.IInspectable
}

func (this *IMediaPlaybackItemError) Vtbl() *IMediaPlaybackItemErrorVtbl {
	return (*IMediaPlaybackItemErrorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackItemError) Get_ErrorCode() MediaPlaybackItemErrorCode {
	var _result MediaPlaybackItemErrorCode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ErrorCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackItemError) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7133FCE1-1769-4FF9-A7C1-38D2C4D42360
var IID_IMediaPlaybackItemFactory = syscall.GUID{0x7133FCE1, 0x1769, 0x4FF9,
	[8]byte{0xA7, 0xC1, 0x38, 0xD2, 0xC4, 0xD4, 0x23, 0x60}}

type IMediaPlaybackItemFactoryInterface interface {
	win32.IInspectableInterface
	Create(source *IMediaSource2) *IMediaPlaybackItem
}

type IMediaPlaybackItemFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IMediaPlaybackItemFactory struct {
	win32.IInspectable
}

func (this *IMediaPlaybackItemFactory) Vtbl() *IMediaPlaybackItemFactoryVtbl {
	return (*IMediaPlaybackItemFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackItemFactory) Create(source *IMediaSource2) *IMediaPlaybackItem {
	var _result *IMediaPlaybackItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(source)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D77CDF3A-B947-4972-B35D-ADFB931A71E6
var IID_IMediaPlaybackItemFactory2 = syscall.GUID{0xD77CDF3A, 0xB947, 0x4972,
	[8]byte{0xB3, 0x5D, 0xAD, 0xFB, 0x93, 0x1A, 0x71, 0xE6}}

type IMediaPlaybackItemFactory2Interface interface {
	win32.IInspectableInterface
	CreateWithStartTime(source *IMediaSource2, startTime TimeSpan) *IMediaPlaybackItem
	CreateWithStartTimeAndDurationLimit(source *IMediaSource2, startTime TimeSpan, durationLimit TimeSpan) *IMediaPlaybackItem
}

type IMediaPlaybackItemFactory2Vtbl struct {
	win32.IInspectableVtbl
	CreateWithStartTime                 uintptr
	CreateWithStartTimeAndDurationLimit uintptr
}

type IMediaPlaybackItemFactory2 struct {
	win32.IInspectable
}

func (this *IMediaPlaybackItemFactory2) Vtbl() *IMediaPlaybackItemFactory2Vtbl {
	return (*IMediaPlaybackItemFactory2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackItemFactory2) CreateWithStartTime(source *IMediaSource2, startTime TimeSpan) *IMediaPlaybackItem {
	var _result *IMediaPlaybackItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithStartTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(source)), *(*uintptr)(unsafe.Pointer(&startTime)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackItemFactory2) CreateWithStartTimeAndDurationLimit(source *IMediaSource2, startTime TimeSpan, durationLimit TimeSpan) *IMediaPlaybackItem {
	var _result *IMediaPlaybackItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithStartTimeAndDurationLimit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(source)), *(*uintptr)(unsafe.Pointer(&startTime)), *(*uintptr)(unsafe.Pointer(&durationLimit)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7703134A-E9A7-47C3-862C-C656D30683D4
var IID_IMediaPlaybackItemFailedEventArgs = syscall.GUID{0x7703134A, 0xE9A7, 0x47C3,
	[8]byte{0x86, 0x2C, 0xC6, 0x56, 0xD3, 0x06, 0x83, 0xD4}}

type IMediaPlaybackItemFailedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Item() *IMediaPlaybackItem
	Get_Error() *IMediaPlaybackItemError
}

type IMediaPlaybackItemFailedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Item  uintptr
	Get_Error uintptr
}

type IMediaPlaybackItemFailedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaPlaybackItemFailedEventArgs) Vtbl() *IMediaPlaybackItemFailedEventArgsVtbl {
	return (*IMediaPlaybackItemFailedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackItemFailedEventArgs) Get_Item() *IMediaPlaybackItem {
	var _result *IMediaPlaybackItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Item, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackItemFailedEventArgs) Get_Error() *IMediaPlaybackItemError {
	var _result *IMediaPlaybackItemError
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Error, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CBD9BD82-3037-4FBE-AE8F-39FC39EDF4EF
var IID_IMediaPlaybackItemOpenedEventArgs = syscall.GUID{0xCBD9BD82, 0x3037, 0x4FBE,
	[8]byte{0xAE, 0x8F, 0x39, 0xFC, 0x39, 0xED, 0xF4, 0xEF}}

type IMediaPlaybackItemOpenedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Item() *IMediaPlaybackItem
}

type IMediaPlaybackItemOpenedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Item uintptr
}

type IMediaPlaybackItemOpenedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaPlaybackItemOpenedEventArgs) Vtbl() *IMediaPlaybackItemOpenedEventArgsVtbl {
	return (*IMediaPlaybackItemOpenedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackItemOpenedEventArgs) Get_Item() *IMediaPlaybackItem {
	var _result *IMediaPlaybackItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Item, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4B1BE7F4-4345-403C-8A67-F5DE91DF4C86
var IID_IMediaPlaybackItemStatics = syscall.GUID{0x4B1BE7F4, 0x4345, 0x403C,
	[8]byte{0x8A, 0x67, 0xF5, 0xDE, 0x91, 0xDF, 0x4C, 0x86}}

type IMediaPlaybackItemStaticsInterface interface {
	win32.IInspectableInterface
	FindFromMediaSource(source *IMediaSource2) *IMediaPlaybackItem
}

type IMediaPlaybackItemStaticsVtbl struct {
	win32.IInspectableVtbl
	FindFromMediaSource uintptr
}

type IMediaPlaybackItemStatics struct {
	win32.IInspectable
}

func (this *IMediaPlaybackItemStatics) Vtbl() *IMediaPlaybackItemStaticsVtbl {
	return (*IMediaPlaybackItemStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackItemStatics) FindFromMediaSource(source *IMediaSource2) *IMediaPlaybackItem {
	var _result *IMediaPlaybackItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindFromMediaSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(source)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7F77EE9C-DC42-4E26-A98D-7850DF8EC925
var IID_IMediaPlaybackList = syscall.GUID{0x7F77EE9C, 0xDC42, 0x4E26,
	[8]byte{0xA9, 0x8D, 0x78, 0x50, 0xDF, 0x8E, 0xC9, 0x25}}

type IMediaPlaybackListInterface interface {
	win32.IInspectableInterface
	Add_ItemFailed(handler TypedEventHandler[*IMediaPlaybackList, *IMediaPlaybackItemFailedEventArgs]) EventRegistrationToken
	Remove_ItemFailed(token EventRegistrationToken)
	Add_CurrentItemChanged(handler TypedEventHandler[*IMediaPlaybackList, *ICurrentMediaPlaybackItemChangedEventArgs]) EventRegistrationToken
	Remove_CurrentItemChanged(token EventRegistrationToken)
	Add_ItemOpened(handler TypedEventHandler[*IMediaPlaybackList, *IMediaPlaybackItemOpenedEventArgs]) EventRegistrationToken
	Remove_ItemOpened(token EventRegistrationToken)
	Get_Items() *IObservableVector[*IMediaPlaybackItem]
	Get_AutoRepeatEnabled() bool
	Put_AutoRepeatEnabled(value bool)
	Get_ShuffleEnabled() bool
	Put_ShuffleEnabled(value bool)
	Get_CurrentItem() *IMediaPlaybackItem
	Get_CurrentItemIndex() uint32
	MoveNext() *IMediaPlaybackItem
	MovePrevious() *IMediaPlaybackItem
	MoveTo(itemIndex uint32) *IMediaPlaybackItem
}

type IMediaPlaybackListVtbl struct {
	win32.IInspectableVtbl
	Add_ItemFailed            uintptr
	Remove_ItemFailed         uintptr
	Add_CurrentItemChanged    uintptr
	Remove_CurrentItemChanged uintptr
	Add_ItemOpened            uintptr
	Remove_ItemOpened         uintptr
	Get_Items                 uintptr
	Get_AutoRepeatEnabled     uintptr
	Put_AutoRepeatEnabled     uintptr
	Get_ShuffleEnabled        uintptr
	Put_ShuffleEnabled        uintptr
	Get_CurrentItem           uintptr
	Get_CurrentItemIndex      uintptr
	MoveNext                  uintptr
	MovePrevious              uintptr
	MoveTo                    uintptr
}

type IMediaPlaybackList struct {
	win32.IInspectable
}

func (this *IMediaPlaybackList) Vtbl() *IMediaPlaybackListVtbl {
	return (*IMediaPlaybackListVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackList) Add_ItemFailed(handler TypedEventHandler[*IMediaPlaybackList, *IMediaPlaybackItemFailedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ItemFailed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackList) Remove_ItemFailed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ItemFailed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackList) Add_CurrentItemChanged(handler TypedEventHandler[*IMediaPlaybackList, *ICurrentMediaPlaybackItemChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_CurrentItemChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackList) Remove_CurrentItemChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_CurrentItemChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackList) Add_ItemOpened(handler TypedEventHandler[*IMediaPlaybackList, *IMediaPlaybackItemOpenedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ItemOpened, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackList) Remove_ItemOpened(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ItemOpened, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackList) Get_Items() *IObservableVector[*IMediaPlaybackItem] {
	var _result *IObservableVector[*IMediaPlaybackItem]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Items, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackList) Get_AutoRepeatEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AutoRepeatEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackList) Put_AutoRepeatEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AutoRepeatEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaPlaybackList) Get_ShuffleEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ShuffleEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackList) Put_ShuffleEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ShuffleEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaPlaybackList) Get_CurrentItem() *IMediaPlaybackItem {
	var _result *IMediaPlaybackItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentItem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackList) Get_CurrentItemIndex() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentItemIndex, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackList) MoveNext() *IMediaPlaybackItem {
	var _result *IMediaPlaybackItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MoveNext, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackList) MovePrevious() *IMediaPlaybackItem {
	var _result *IMediaPlaybackItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MovePrevious, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackList) MoveTo(itemIndex uint32) *IMediaPlaybackItem {
	var _result *IMediaPlaybackItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MoveTo, uintptr(unsafe.Pointer(this)), uintptr(itemIndex), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0E09B478-600A-4274-A14B-0B6723D0F48B
var IID_IMediaPlaybackList2 = syscall.GUID{0x0E09B478, 0x600A, 0x4274,
	[8]byte{0xA1, 0x4B, 0x0B, 0x67, 0x23, 0xD0, 0xF4, 0x8B}}

type IMediaPlaybackList2Interface interface {
	win32.IInspectableInterface
	Get_MaxPrefetchTime() *IReference[TimeSpan]
	Put_MaxPrefetchTime(value *IReference[TimeSpan])
	Get_StartingItem() *IMediaPlaybackItem
	Put_StartingItem(value *IMediaPlaybackItem)
	Get_ShuffledItems() *IVectorView[*IMediaPlaybackItem]
	SetShuffledItems(value *IIterable[*IMediaPlaybackItem])
}

type IMediaPlaybackList2Vtbl struct {
	win32.IInspectableVtbl
	Get_MaxPrefetchTime uintptr
	Put_MaxPrefetchTime uintptr
	Get_StartingItem    uintptr
	Put_StartingItem    uintptr
	Get_ShuffledItems   uintptr
	SetShuffledItems    uintptr
}

type IMediaPlaybackList2 struct {
	win32.IInspectable
}

func (this *IMediaPlaybackList2) Vtbl() *IMediaPlaybackList2Vtbl {
	return (*IMediaPlaybackList2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackList2) Get_MaxPrefetchTime() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxPrefetchTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackList2) Put_MaxPrefetchTime(value *IReference[TimeSpan]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxPrefetchTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaPlaybackList2) Get_StartingItem() *IMediaPlaybackItem {
	var _result *IMediaPlaybackItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StartingItem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackList2) Put_StartingItem(value *IMediaPlaybackItem) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StartingItem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaPlaybackList2) Get_ShuffledItems() *IVectorView[*IMediaPlaybackItem] {
	var _result *IVectorView[*IMediaPlaybackItem]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ShuffledItems, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackList2) SetShuffledItems(value *IIterable[*IMediaPlaybackItem]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetShuffledItems, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// DD24BBA9-BC47-4463-AA90-C18B7E5FFDE1
var IID_IMediaPlaybackList3 = syscall.GUID{0xDD24BBA9, 0xBC47, 0x4463,
	[8]byte{0xAA, 0x90, 0xC1, 0x8B, 0x7E, 0x5F, 0xFD, 0xE1}}

type IMediaPlaybackList3Interface interface {
	win32.IInspectableInterface
	Get_MaxPlayedItemsToKeepOpen() *IReference[uint32]
	Put_MaxPlayedItemsToKeepOpen(value *IReference[uint32])
}

type IMediaPlaybackList3Vtbl struct {
	win32.IInspectableVtbl
	Get_MaxPlayedItemsToKeepOpen uintptr
	Put_MaxPlayedItemsToKeepOpen uintptr
}

type IMediaPlaybackList3 struct {
	win32.IInspectable
}

func (this *IMediaPlaybackList3) Vtbl() *IMediaPlaybackList3Vtbl {
	return (*IMediaPlaybackList3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackList3) Get_MaxPlayedItemsToKeepOpen() *IReference[uint32] {
	var _result *IReference[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxPlayedItemsToKeepOpen, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackList3) Put_MaxPlayedItemsToKeepOpen(value *IReference[uint32]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxPlayedItemsToKeepOpen, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// C32B683D-0407-41BA-8946-8B345A5A5435
var IID_IMediaPlaybackSession = syscall.GUID{0xC32B683D, 0x0407, 0x41BA,
	[8]byte{0x89, 0x46, 0x8B, 0x34, 0x5A, 0x5A, 0x54, 0x35}}

type IMediaPlaybackSessionInterface interface {
	win32.IInspectableInterface
	Add_PlaybackStateChanged(value TypedEventHandler[*IMediaPlaybackSession, interface{}]) EventRegistrationToken
	Remove_PlaybackStateChanged(token EventRegistrationToken)
	Add_PlaybackRateChanged(value TypedEventHandler[*IMediaPlaybackSession, interface{}]) EventRegistrationToken
	Remove_PlaybackRateChanged(token EventRegistrationToken)
	Add_SeekCompleted(value TypedEventHandler[*IMediaPlaybackSession, interface{}]) EventRegistrationToken
	Remove_SeekCompleted(token EventRegistrationToken)
	Add_BufferingStarted(value TypedEventHandler[*IMediaPlaybackSession, interface{}]) EventRegistrationToken
	Remove_BufferingStarted(token EventRegistrationToken)
	Add_BufferingEnded(value TypedEventHandler[*IMediaPlaybackSession, interface{}]) EventRegistrationToken
	Remove_BufferingEnded(token EventRegistrationToken)
	Add_BufferingProgressChanged(value TypedEventHandler[*IMediaPlaybackSession, interface{}]) EventRegistrationToken
	Remove_BufferingProgressChanged(token EventRegistrationToken)
	Add_DownloadProgressChanged(value TypedEventHandler[*IMediaPlaybackSession, interface{}]) EventRegistrationToken
	Remove_DownloadProgressChanged(token EventRegistrationToken)
	Add_NaturalDurationChanged(value TypedEventHandler[*IMediaPlaybackSession, interface{}]) EventRegistrationToken
	Remove_NaturalDurationChanged(token EventRegistrationToken)
	Add_PositionChanged(value TypedEventHandler[*IMediaPlaybackSession, interface{}]) EventRegistrationToken
	Remove_PositionChanged(token EventRegistrationToken)
	Add_NaturalVideoSizeChanged(value TypedEventHandler[*IMediaPlaybackSession, interface{}]) EventRegistrationToken
	Remove_NaturalVideoSizeChanged(token EventRegistrationToken)
	Get_MediaPlayer() *IMediaPlayer
	Get_NaturalDuration() TimeSpan
	Get_Position() TimeSpan
	Put_Position(value TimeSpan)
	Get_PlaybackState() MediaPlaybackState
	Get_CanSeek() bool
	Get_CanPause() bool
	Get_IsProtected() bool
	Get_PlaybackRate() float64
	Put_PlaybackRate(value float64)
	Get_BufferingProgress() float64
	Get_DownloadProgress() float64
	Get_NaturalVideoHeight() uint32
	Get_NaturalVideoWidth() uint32
	Get_NormalizedSourceRect() Rect
	Put_NormalizedSourceRect(value Rect)
	Get_StereoscopicVideoPackingMode() StereoscopicVideoPackingMode
	Put_StereoscopicVideoPackingMode(value StereoscopicVideoPackingMode)
}

type IMediaPlaybackSessionVtbl struct {
	win32.IInspectableVtbl
	Add_PlaybackStateChanged         uintptr
	Remove_PlaybackStateChanged      uintptr
	Add_PlaybackRateChanged          uintptr
	Remove_PlaybackRateChanged       uintptr
	Add_SeekCompleted                uintptr
	Remove_SeekCompleted             uintptr
	Add_BufferingStarted             uintptr
	Remove_BufferingStarted          uintptr
	Add_BufferingEnded               uintptr
	Remove_BufferingEnded            uintptr
	Add_BufferingProgressChanged     uintptr
	Remove_BufferingProgressChanged  uintptr
	Add_DownloadProgressChanged      uintptr
	Remove_DownloadProgressChanged   uintptr
	Add_NaturalDurationChanged       uintptr
	Remove_NaturalDurationChanged    uintptr
	Add_PositionChanged              uintptr
	Remove_PositionChanged           uintptr
	Add_NaturalVideoSizeChanged      uintptr
	Remove_NaturalVideoSizeChanged   uintptr
	Get_MediaPlayer                  uintptr
	Get_NaturalDuration              uintptr
	Get_Position                     uintptr
	Put_Position                     uintptr
	Get_PlaybackState                uintptr
	Get_CanSeek                      uintptr
	Get_CanPause                     uintptr
	Get_IsProtected                  uintptr
	Get_PlaybackRate                 uintptr
	Put_PlaybackRate                 uintptr
	Get_BufferingProgress            uintptr
	Get_DownloadProgress             uintptr
	Get_NaturalVideoHeight           uintptr
	Get_NaturalVideoWidth            uintptr
	Get_NormalizedSourceRect         uintptr
	Put_NormalizedSourceRect         uintptr
	Get_StereoscopicVideoPackingMode uintptr
	Put_StereoscopicVideoPackingMode uintptr
}

type IMediaPlaybackSession struct {
	win32.IInspectable
}

func (this *IMediaPlaybackSession) Vtbl() *IMediaPlaybackSessionVtbl {
	return (*IMediaPlaybackSessionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackSession) Add_PlaybackStateChanged(value TypedEventHandler[*IMediaPlaybackSession, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PlaybackStateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession) Remove_PlaybackStateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PlaybackStateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackSession) Add_PlaybackRateChanged(value TypedEventHandler[*IMediaPlaybackSession, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PlaybackRateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession) Remove_PlaybackRateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PlaybackRateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackSession) Add_SeekCompleted(value TypedEventHandler[*IMediaPlaybackSession, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SeekCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession) Remove_SeekCompleted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SeekCompleted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackSession) Add_BufferingStarted(value TypedEventHandler[*IMediaPlaybackSession, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_BufferingStarted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession) Remove_BufferingStarted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_BufferingStarted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackSession) Add_BufferingEnded(value TypedEventHandler[*IMediaPlaybackSession, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_BufferingEnded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession) Remove_BufferingEnded(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_BufferingEnded, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackSession) Add_BufferingProgressChanged(value TypedEventHandler[*IMediaPlaybackSession, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_BufferingProgressChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession) Remove_BufferingProgressChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_BufferingProgressChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackSession) Add_DownloadProgressChanged(value TypedEventHandler[*IMediaPlaybackSession, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DownloadProgressChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession) Remove_DownloadProgressChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DownloadProgressChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackSession) Add_NaturalDurationChanged(value TypedEventHandler[*IMediaPlaybackSession, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_NaturalDurationChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession) Remove_NaturalDurationChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_NaturalDurationChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackSession) Add_PositionChanged(value TypedEventHandler[*IMediaPlaybackSession, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PositionChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession) Remove_PositionChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PositionChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackSession) Add_NaturalVideoSizeChanged(value TypedEventHandler[*IMediaPlaybackSession, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_NaturalVideoSizeChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession) Remove_NaturalVideoSizeChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_NaturalVideoSizeChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackSession) Get_MediaPlayer() *IMediaPlayer {
	var _result *IMediaPlayer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaPlayer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackSession) Get_NaturalDuration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NaturalDuration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession) Get_Position() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession) Put_Position(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Position, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IMediaPlaybackSession) Get_PlaybackState() MediaPlaybackState {
	var _result MediaPlaybackState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlaybackState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession) Get_CanSeek() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanSeek, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession) Get_CanPause() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanPause, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession) Get_IsProtected() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsProtected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession) Get_PlaybackRate() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlaybackRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession) Put_PlaybackRate(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PlaybackRate, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMediaPlaybackSession) Get_BufferingProgress() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BufferingProgress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession) Get_DownloadProgress() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DownloadProgress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession) Get_NaturalVideoHeight() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NaturalVideoHeight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession) Get_NaturalVideoWidth() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NaturalVideoWidth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession) Get_NormalizedSourceRect() Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NormalizedSourceRect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession) Put_NormalizedSourceRect(value Rect) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_NormalizedSourceRect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IMediaPlaybackSession) Get_StereoscopicVideoPackingMode() StereoscopicVideoPackingMode {
	var _result StereoscopicVideoPackingMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StereoscopicVideoPackingMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession) Put_StereoscopicVideoPackingMode(value StereoscopicVideoPackingMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StereoscopicVideoPackingMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// F8BA7C79-1FC8-4097-AD70-C0FA18CC0050
var IID_IMediaPlaybackSession2 = syscall.GUID{0xF8BA7C79, 0x1FC8, 0x4097,
	[8]byte{0xAD, 0x70, 0xC0, 0xFA, 0x18, 0xCC, 0x00, 0x50}}

type IMediaPlaybackSession2Interface interface {
	win32.IInspectableInterface
	Add_BufferedRangesChanged(value TypedEventHandler[*IMediaPlaybackSession, interface{}]) EventRegistrationToken
	Remove_BufferedRangesChanged(token EventRegistrationToken)
	Add_PlayedRangesChanged(value TypedEventHandler[*IMediaPlaybackSession, interface{}]) EventRegistrationToken
	Remove_PlayedRangesChanged(token EventRegistrationToken)
	Add_SeekableRangesChanged(value TypedEventHandler[*IMediaPlaybackSession, interface{}]) EventRegistrationToken
	Remove_SeekableRangesChanged(token EventRegistrationToken)
	Add_SupportedPlaybackRatesChanged(value TypedEventHandler[*IMediaPlaybackSession, interface{}]) EventRegistrationToken
	Remove_SupportedPlaybackRatesChanged(token EventRegistrationToken)
	Get_SphericalVideoProjection() *IMediaPlaybackSphericalVideoProjection
	Get_IsMirroring() bool
	Put_IsMirroring(value bool)
	GetBufferedRanges() *IVectorView[MediaTimeRange]
	GetPlayedRanges() *IVectorView[MediaTimeRange]
	GetSeekableRanges() *IVectorView[MediaTimeRange]
	IsSupportedPlaybackRateRange(rate1 float64, rate2 float64) bool
}

type IMediaPlaybackSession2Vtbl struct {
	win32.IInspectableVtbl
	Add_BufferedRangesChanged            uintptr
	Remove_BufferedRangesChanged         uintptr
	Add_PlayedRangesChanged              uintptr
	Remove_PlayedRangesChanged           uintptr
	Add_SeekableRangesChanged            uintptr
	Remove_SeekableRangesChanged         uintptr
	Add_SupportedPlaybackRatesChanged    uintptr
	Remove_SupportedPlaybackRatesChanged uintptr
	Get_SphericalVideoProjection         uintptr
	Get_IsMirroring                      uintptr
	Put_IsMirroring                      uintptr
	GetBufferedRanges                    uintptr
	GetPlayedRanges                      uintptr
	GetSeekableRanges                    uintptr
	IsSupportedPlaybackRateRange         uintptr
}

type IMediaPlaybackSession2 struct {
	win32.IInspectable
}

func (this *IMediaPlaybackSession2) Vtbl() *IMediaPlaybackSession2Vtbl {
	return (*IMediaPlaybackSession2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackSession2) Add_BufferedRangesChanged(value TypedEventHandler[*IMediaPlaybackSession, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_BufferedRangesChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession2) Remove_BufferedRangesChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_BufferedRangesChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackSession2) Add_PlayedRangesChanged(value TypedEventHandler[*IMediaPlaybackSession, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PlayedRangesChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession2) Remove_PlayedRangesChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PlayedRangesChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackSession2) Add_SeekableRangesChanged(value TypedEventHandler[*IMediaPlaybackSession, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SeekableRangesChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession2) Remove_SeekableRangesChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SeekableRangesChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackSession2) Add_SupportedPlaybackRatesChanged(value TypedEventHandler[*IMediaPlaybackSession, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SupportedPlaybackRatesChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession2) Remove_SupportedPlaybackRatesChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SupportedPlaybackRatesChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackSession2) Get_SphericalVideoProjection() *IMediaPlaybackSphericalVideoProjection {
	var _result *IMediaPlaybackSphericalVideoProjection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SphericalVideoProjection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackSession2) Get_IsMirroring() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsMirroring, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession2) Put_IsMirroring(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsMirroring, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaPlaybackSession2) GetBufferedRanges() *IVectorView[MediaTimeRange] {
	var _result *IVectorView[MediaTimeRange]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetBufferedRanges, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackSession2) GetPlayedRanges() *IVectorView[MediaTimeRange] {
	var _result *IVectorView[MediaTimeRange]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPlayedRanges, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackSession2) GetSeekableRanges() *IVectorView[MediaTimeRange] {
	var _result *IVectorView[MediaTimeRange]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSeekableRanges, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlaybackSession2) IsSupportedPlaybackRateRange(rate1 float64, rate2 float64) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsSupportedPlaybackRateRange, uintptr(unsafe.Pointer(this)), uintptr(rate1), uintptr(rate2), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7BA2B41A-A3E2-405F-B77B-A4812C238B66
var IID_IMediaPlaybackSession3 = syscall.GUID{0x7BA2B41A, 0xA3E2, 0x405F,
	[8]byte{0xB7, 0x7B, 0xA4, 0x81, 0x2C, 0x23, 0x8B, 0x66}}

type IMediaPlaybackSession3Interface interface {
	win32.IInspectableInterface
	Get_PlaybackRotation() MediaRotation
	Put_PlaybackRotation(value MediaRotation)
	GetOutputDegradationPolicyState() *IMediaPlaybackSessionOutputDegradationPolicyState
}

type IMediaPlaybackSession3Vtbl struct {
	win32.IInspectableVtbl
	Get_PlaybackRotation            uintptr
	Put_PlaybackRotation            uintptr
	GetOutputDegradationPolicyState uintptr
}

type IMediaPlaybackSession3 struct {
	win32.IInspectable
}

func (this *IMediaPlaybackSession3) Vtbl() *IMediaPlaybackSession3Vtbl {
	return (*IMediaPlaybackSession3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackSession3) Get_PlaybackRotation() MediaRotation {
	var _result MediaRotation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlaybackRotation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSession3) Put_PlaybackRotation(value MediaRotation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PlaybackRotation, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMediaPlaybackSession3) GetOutputDegradationPolicyState() *IMediaPlaybackSessionOutputDegradationPolicyState {
	var _result *IMediaPlaybackSessionOutputDegradationPolicyState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetOutputDegradationPolicyState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CD6AAFED-74E2-43B5-B115-76236C33791A
var IID_IMediaPlaybackSessionBufferingStartedEventArgs = syscall.GUID{0xCD6AAFED, 0x74E2, 0x43B5,
	[8]byte{0xB1, 0x15, 0x76, 0x23, 0x6C, 0x33, 0x79, 0x1A}}

type IMediaPlaybackSessionBufferingStartedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_IsPlaybackInterruption() bool
}

type IMediaPlaybackSessionBufferingStartedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_IsPlaybackInterruption uintptr
}

type IMediaPlaybackSessionBufferingStartedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaPlaybackSessionBufferingStartedEventArgs) Vtbl() *IMediaPlaybackSessionBufferingStartedEventArgsVtbl {
	return (*IMediaPlaybackSessionBufferingStartedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackSessionBufferingStartedEventArgs) Get_IsPlaybackInterruption() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPlaybackInterruption, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 558E727D-F633-49F9-965A-ABAA1DB709BE
var IID_IMediaPlaybackSessionOutputDegradationPolicyState = syscall.GUID{0x558E727D, 0xF633, 0x49F9,
	[8]byte{0x96, 0x5A, 0xAB, 0xAA, 0x1D, 0xB7, 0x09, 0xBE}}

type IMediaPlaybackSessionOutputDegradationPolicyStateInterface interface {
	win32.IInspectableInterface
	Get_VideoConstrictionReason() MediaPlaybackSessionVideoConstrictionReason
}

type IMediaPlaybackSessionOutputDegradationPolicyStateVtbl struct {
	win32.IInspectableVtbl
	Get_VideoConstrictionReason uintptr
}

type IMediaPlaybackSessionOutputDegradationPolicyState struct {
	win32.IInspectable
}

func (this *IMediaPlaybackSessionOutputDegradationPolicyState) Vtbl() *IMediaPlaybackSessionOutputDegradationPolicyStateVtbl {
	return (*IMediaPlaybackSessionOutputDegradationPolicyStateVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackSessionOutputDegradationPolicyState) Get_VideoConstrictionReason() MediaPlaybackSessionVideoConstrictionReason {
	var _result MediaPlaybackSessionVideoConstrictionReason
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoConstrictionReason, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// EF9DC2BC-9317-4696-B051-2BAD643177B5
var IID_IMediaPlaybackSource = syscall.GUID{0xEF9DC2BC, 0x9317, 0x4696,
	[8]byte{0xB0, 0x51, 0x2B, 0xAD, 0x64, 0x31, 0x77, 0xB5}}

type IMediaPlaybackSourceInterface interface {
	win32.IInspectableInterface
}

type IMediaPlaybackSourceVtbl struct {
	win32.IInspectableVtbl
}

type IMediaPlaybackSource struct {
	win32.IInspectable
}

func (this *IMediaPlaybackSource) Vtbl() *IMediaPlaybackSourceVtbl {
	return (*IMediaPlaybackSourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// D405B37C-6F0E-4661-B8EE-D487BA9752D5
var IID_IMediaPlaybackSphericalVideoProjection = syscall.GUID{0xD405B37C, 0x6F0E, 0x4661,
	[8]byte{0xB8, 0xEE, 0xD4, 0x87, 0xBA, 0x97, 0x52, 0xD5}}

type IMediaPlaybackSphericalVideoProjectionInterface interface {
	win32.IInspectableInterface
	Get_IsEnabled() bool
	Put_IsEnabled(value bool)
	Get_FrameFormat() SphericalVideoFrameFormat
	Put_FrameFormat(value SphericalVideoFrameFormat)
	Get_HorizontalFieldOfViewInDegrees() float64
	Put_HorizontalFieldOfViewInDegrees(value float64)
	Get_ViewOrientation() Quaternion
	Put_ViewOrientation(value Quaternion)
	Get_ProjectionMode() SphericalVideoProjectionMode
	Put_ProjectionMode(value SphericalVideoProjectionMode)
}

type IMediaPlaybackSphericalVideoProjectionVtbl struct {
	win32.IInspectableVtbl
	Get_IsEnabled                      uintptr
	Put_IsEnabled                      uintptr
	Get_FrameFormat                    uintptr
	Put_FrameFormat                    uintptr
	Get_HorizontalFieldOfViewInDegrees uintptr
	Put_HorizontalFieldOfViewInDegrees uintptr
	Get_ViewOrientation                uintptr
	Put_ViewOrientation                uintptr
	Get_ProjectionMode                 uintptr
	Put_ProjectionMode                 uintptr
}

type IMediaPlaybackSphericalVideoProjection struct {
	win32.IInspectable
}

func (this *IMediaPlaybackSphericalVideoProjection) Vtbl() *IMediaPlaybackSphericalVideoProjectionVtbl {
	return (*IMediaPlaybackSphericalVideoProjectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackSphericalVideoProjection) Get_IsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSphericalVideoProjection) Put_IsEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaPlaybackSphericalVideoProjection) Get_FrameFormat() SphericalVideoFrameFormat {
	var _result SphericalVideoFrameFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSphericalVideoProjection) Put_FrameFormat(value SphericalVideoFrameFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FrameFormat, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMediaPlaybackSphericalVideoProjection) Get_HorizontalFieldOfViewInDegrees() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HorizontalFieldOfViewInDegrees, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSphericalVideoProjection) Put_HorizontalFieldOfViewInDegrees(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_HorizontalFieldOfViewInDegrees, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMediaPlaybackSphericalVideoProjection) Get_ViewOrientation() Quaternion {
	var _result Quaternion
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ViewOrientation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSphericalVideoProjection) Put_ViewOrientation(value Quaternion) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ViewOrientation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IMediaPlaybackSphericalVideoProjection) Get_ProjectionMode() SphericalVideoProjectionMode {
	var _result SphericalVideoProjectionMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProjectionMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackSphericalVideoProjection) Put_ProjectionMode(value SphericalVideoProjectionMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ProjectionMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 72B41319-BBFB-46A3-9372-9C9C744B9438
var IID_IMediaPlaybackTimedMetadataTrackList = syscall.GUID{0x72B41319, 0xBBFB, 0x46A3,
	[8]byte{0x93, 0x72, 0x9C, 0x9C, 0x74, 0x4B, 0x94, 0x38}}

type IMediaPlaybackTimedMetadataTrackListInterface interface {
	win32.IInspectableInterface
	Add_PresentationModeChanged(handler TypedEventHandler[*IVectorView[*ITimedMetadataTrack], *ITimedMetadataPresentationModeChangedEventArgs]) EventRegistrationToken
	Remove_PresentationModeChanged(token EventRegistrationToken)
	GetPresentationMode(index uint32) TimedMetadataTrackPresentationMode
	SetPresentationMode(index uint32, value TimedMetadataTrackPresentationMode)
}

type IMediaPlaybackTimedMetadataTrackListVtbl struct {
	win32.IInspectableVtbl
	Add_PresentationModeChanged    uintptr
	Remove_PresentationModeChanged uintptr
	GetPresentationMode            uintptr
	SetPresentationMode            uintptr
}

type IMediaPlaybackTimedMetadataTrackList struct {
	win32.IInspectable
}

func (this *IMediaPlaybackTimedMetadataTrackList) Vtbl() *IMediaPlaybackTimedMetadataTrackListVtbl {
	return (*IMediaPlaybackTimedMetadataTrackListVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlaybackTimedMetadataTrackList) Add_PresentationModeChanged(handler TypedEventHandler[*IVectorView[*ITimedMetadataTrack], *ITimedMetadataPresentationModeChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PresentationModeChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackTimedMetadataTrackList) Remove_PresentationModeChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PresentationModeChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlaybackTimedMetadataTrackList) GetPresentationMode(index uint32) TimedMetadataTrackPresentationMode {
	var _result TimedMetadataTrackPresentationMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPresentationMode, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlaybackTimedMetadataTrackList) SetPresentationMode(index uint32, value TimedMetadataTrackPresentationMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetPresentationMode, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(value))
	_ = _hr
}

// 381A83CB-6FFF-499B-8D64-2885DFC1249E
var IID_IMediaPlayer = syscall.GUID{0x381A83CB, 0x6FFF, 0x499B,
	[8]byte{0x8D, 0x64, 0x28, 0x85, 0xDF, 0xC1, 0x24, 0x9E}}

type IMediaPlayerInterface interface {
	win32.IInspectableInterface
	Get_AutoPlay() bool
	Put_AutoPlay(value bool)
	Get_NaturalDuration() TimeSpan
	Get_Position() TimeSpan
	Put_Position(value TimeSpan)
	Get_BufferingProgress() float64
	Get_CurrentState() MediaPlayerState
	Get_CanSeek() bool
	Get_CanPause() bool
	Get_IsLoopingEnabled() bool
	Put_IsLoopingEnabled(value bool)
	Get_IsProtected() bool
	Get_IsMuted() bool
	Put_IsMuted(value bool)
	Get_PlaybackRate() float64
	Put_PlaybackRate(value float64)
	Get_Volume() float64
	Put_Volume(value float64)
	Get_PlaybackMediaMarkers() *IPlaybackMediaMarkerSequence
	Add_MediaOpened(value TypedEventHandler[*IMediaPlayer, interface{}]) EventRegistrationToken
	Remove_MediaOpened(token EventRegistrationToken)
	Add_MediaEnded(value TypedEventHandler[*IMediaPlayer, interface{}]) EventRegistrationToken
	Remove_MediaEnded(token EventRegistrationToken)
	Add_MediaFailed(value TypedEventHandler[*IMediaPlayer, *IMediaPlayerFailedEventArgs]) EventRegistrationToken
	Remove_MediaFailed(token EventRegistrationToken)
	Add_CurrentStateChanged(value TypedEventHandler[*IMediaPlayer, interface{}]) EventRegistrationToken
	Remove_CurrentStateChanged(token EventRegistrationToken)
	Add_PlaybackMediaMarkerReached(value TypedEventHandler[*IMediaPlayer, *IPlaybackMediaMarkerReachedEventArgs]) EventRegistrationToken
	Remove_PlaybackMediaMarkerReached(token EventRegistrationToken)
	Add_MediaPlayerRateChanged(value TypedEventHandler[*IMediaPlayer, *IMediaPlayerRateChangedEventArgs]) EventRegistrationToken
	Remove_MediaPlayerRateChanged(token EventRegistrationToken)
	Add_VolumeChanged(value TypedEventHandler[*IMediaPlayer, interface{}]) EventRegistrationToken
	Remove_VolumeChanged(token EventRegistrationToken)
	Add_SeekCompleted(value TypedEventHandler[*IMediaPlayer, interface{}]) EventRegistrationToken
	Remove_SeekCompleted(token EventRegistrationToken)
	Add_BufferingStarted(value TypedEventHandler[*IMediaPlayer, interface{}]) EventRegistrationToken
	Remove_BufferingStarted(token EventRegistrationToken)
	Add_BufferingEnded(value TypedEventHandler[*IMediaPlayer, interface{}]) EventRegistrationToken
	Remove_BufferingEnded(token EventRegistrationToken)
	Play()
	Pause()
	SetUriSource(value *IUriRuntimeClass)
}

type IMediaPlayerVtbl struct {
	win32.IInspectableVtbl
	Get_AutoPlay                      uintptr
	Put_AutoPlay                      uintptr
	Get_NaturalDuration               uintptr
	Get_Position                      uintptr
	Put_Position                      uintptr
	Get_BufferingProgress             uintptr
	Get_CurrentState                  uintptr
	Get_CanSeek                       uintptr
	Get_CanPause                      uintptr
	Get_IsLoopingEnabled              uintptr
	Put_IsLoopingEnabled              uintptr
	Get_IsProtected                   uintptr
	Get_IsMuted                       uintptr
	Put_IsMuted                       uintptr
	Get_PlaybackRate                  uintptr
	Put_PlaybackRate                  uintptr
	Get_Volume                        uintptr
	Put_Volume                        uintptr
	Get_PlaybackMediaMarkers          uintptr
	Add_MediaOpened                   uintptr
	Remove_MediaOpened                uintptr
	Add_MediaEnded                    uintptr
	Remove_MediaEnded                 uintptr
	Add_MediaFailed                   uintptr
	Remove_MediaFailed                uintptr
	Add_CurrentStateChanged           uintptr
	Remove_CurrentStateChanged        uintptr
	Add_PlaybackMediaMarkerReached    uintptr
	Remove_PlaybackMediaMarkerReached uintptr
	Add_MediaPlayerRateChanged        uintptr
	Remove_MediaPlayerRateChanged     uintptr
	Add_VolumeChanged                 uintptr
	Remove_VolumeChanged              uintptr
	Add_SeekCompleted                 uintptr
	Remove_SeekCompleted              uintptr
	Add_BufferingStarted              uintptr
	Remove_BufferingStarted           uintptr
	Add_BufferingEnded                uintptr
	Remove_BufferingEnded             uintptr
	Play                              uintptr
	Pause                             uintptr
	SetUriSource                      uintptr
}

type IMediaPlayer struct {
	win32.IInspectable
}

func (this *IMediaPlayer) Vtbl() *IMediaPlayerVtbl {
	return (*IMediaPlayerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlayer) Get_AutoPlay() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AutoPlay, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer) Put_AutoPlay(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AutoPlay, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaPlayer) Get_NaturalDuration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NaturalDuration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer) Get_Position() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer) Put_Position(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Position, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IMediaPlayer) Get_BufferingProgress() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BufferingProgress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer) Get_CurrentState() MediaPlayerState {
	var _result MediaPlayerState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer) Get_CanSeek() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanSeek, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer) Get_CanPause() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanPause, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer) Get_IsLoopingEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsLoopingEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer) Put_IsLoopingEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsLoopingEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaPlayer) Get_IsProtected() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsProtected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer) Get_IsMuted() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsMuted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer) Put_IsMuted(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsMuted, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaPlayer) Get_PlaybackRate() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlaybackRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer) Put_PlaybackRate(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PlaybackRate, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMediaPlayer) Get_Volume() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Volume, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer) Put_Volume(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Volume, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMediaPlayer) Get_PlaybackMediaMarkers() *IPlaybackMediaMarkerSequence {
	var _result *IPlaybackMediaMarkerSequence
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlaybackMediaMarkers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlayer) Add_MediaOpened(value TypedEventHandler[*IMediaPlayer, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_MediaOpened, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer) Remove_MediaOpened(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_MediaOpened, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlayer) Add_MediaEnded(value TypedEventHandler[*IMediaPlayer, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_MediaEnded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer) Remove_MediaEnded(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_MediaEnded, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlayer) Add_MediaFailed(value TypedEventHandler[*IMediaPlayer, *IMediaPlayerFailedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_MediaFailed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer) Remove_MediaFailed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_MediaFailed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlayer) Add_CurrentStateChanged(value TypedEventHandler[*IMediaPlayer, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_CurrentStateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer) Remove_CurrentStateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_CurrentStateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlayer) Add_PlaybackMediaMarkerReached(value TypedEventHandler[*IMediaPlayer, *IPlaybackMediaMarkerReachedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PlaybackMediaMarkerReached, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer) Remove_PlaybackMediaMarkerReached(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PlaybackMediaMarkerReached, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlayer) Add_MediaPlayerRateChanged(value TypedEventHandler[*IMediaPlayer, *IMediaPlayerRateChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_MediaPlayerRateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer) Remove_MediaPlayerRateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_MediaPlayerRateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlayer) Add_VolumeChanged(value TypedEventHandler[*IMediaPlayer, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_VolumeChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer) Remove_VolumeChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_VolumeChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlayer) Add_SeekCompleted(value TypedEventHandler[*IMediaPlayer, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SeekCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer) Remove_SeekCompleted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SeekCompleted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlayer) Add_BufferingStarted(value TypedEventHandler[*IMediaPlayer, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_BufferingStarted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer) Remove_BufferingStarted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_BufferingStarted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlayer) Add_BufferingEnded(value TypedEventHandler[*IMediaPlayer, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_BufferingEnded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer) Remove_BufferingEnded(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_BufferingEnded, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlayer) Play() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Play, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IMediaPlayer) Pause() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Pause, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IMediaPlayer) SetUriSource(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetUriSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 3C841218-2123-4FC5-9082-2F883F77BDF5
var IID_IMediaPlayer2 = syscall.GUID{0x3C841218, 0x2123, 0x4FC5,
	[8]byte{0x90, 0x82, 0x2F, 0x88, 0x3F, 0x77, 0xBD, 0xF5}}

type IMediaPlayer2Interface interface {
	win32.IInspectableInterface
	Get_SystemMediaTransportControls() *ISystemMediaTransportControls
	Get_AudioCategory() MediaPlayerAudioCategory
	Put_AudioCategory(value MediaPlayerAudioCategory)
	Get_AudioDeviceType() MediaPlayerAudioDeviceType
	Put_AudioDeviceType(value MediaPlayerAudioDeviceType)
}

type IMediaPlayer2Vtbl struct {
	win32.IInspectableVtbl
	Get_SystemMediaTransportControls uintptr
	Get_AudioCategory                uintptr
	Put_AudioCategory                uintptr
	Get_AudioDeviceType              uintptr
	Put_AudioDeviceType              uintptr
}

type IMediaPlayer2 struct {
	win32.IInspectable
}

func (this *IMediaPlayer2) Vtbl() *IMediaPlayer2Vtbl {
	return (*IMediaPlayer2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlayer2) Get_SystemMediaTransportControls() *ISystemMediaTransportControls {
	var _result *ISystemMediaTransportControls
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SystemMediaTransportControls, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlayer2) Get_AudioCategory() MediaPlayerAudioCategory {
	var _result MediaPlayerAudioCategory
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioCategory, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer2) Put_AudioCategory(value MediaPlayerAudioCategory) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AudioCategory, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMediaPlayer2) Get_AudioDeviceType() MediaPlayerAudioDeviceType {
	var _result MediaPlayerAudioDeviceType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioDeviceType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer2) Put_AudioDeviceType(value MediaPlayerAudioDeviceType) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AudioDeviceType, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// EE0660DA-031B-4FEB-BD9B-92E0A0A8D299
var IID_IMediaPlayer3 = syscall.GUID{0xEE0660DA, 0x031B, 0x4FEB,
	[8]byte{0xBD, 0x9B, 0x92, 0xE0, 0xA0, 0xA8, 0xD2, 0x99}}

type IMediaPlayer3Interface interface {
	win32.IInspectableInterface
	Add_IsMutedChanged(value TypedEventHandler[*IMediaPlayer, interface{}]) EventRegistrationToken
	Remove_IsMutedChanged(token EventRegistrationToken)
	Add_SourceChanged(value TypedEventHandler[*IMediaPlayer, interface{}]) EventRegistrationToken
	Remove_SourceChanged(token EventRegistrationToken)
	Get_AudioBalance() float64
	Put_AudioBalance(value float64)
	Get_RealTimePlayback() bool
	Put_RealTimePlayback(value bool)
	Get_StereoscopicVideoRenderMode() StereoscopicVideoRenderMode
	Put_StereoscopicVideoRenderMode(value StereoscopicVideoRenderMode)
	Get_BreakManager() *IMediaBreakManager
	Get_CommandManager() *IMediaPlaybackCommandManager
	Get_AudioDevice() *IDeviceInformation
	Put_AudioDevice(value *IDeviceInformation)
	Get_TimelineController() *IMediaTimelineController
	Put_TimelineController(value *IMediaTimelineController)
	Get_TimelineControllerPositionOffset() TimeSpan
	Put_TimelineControllerPositionOffset(value TimeSpan)
	Get_PlaybackSession() *IMediaPlaybackSession
	StepForwardOneFrame()
	StepBackwardOneFrame()
	GetAsCastingSource() *ICastingSource
}

type IMediaPlayer3Vtbl struct {
	win32.IInspectableVtbl
	Add_IsMutedChanged                   uintptr
	Remove_IsMutedChanged                uintptr
	Add_SourceChanged                    uintptr
	Remove_SourceChanged                 uintptr
	Get_AudioBalance                     uintptr
	Put_AudioBalance                     uintptr
	Get_RealTimePlayback                 uintptr
	Put_RealTimePlayback                 uintptr
	Get_StereoscopicVideoRenderMode      uintptr
	Put_StereoscopicVideoRenderMode      uintptr
	Get_BreakManager                     uintptr
	Get_CommandManager                   uintptr
	Get_AudioDevice                      uintptr
	Put_AudioDevice                      uintptr
	Get_TimelineController               uintptr
	Put_TimelineController               uintptr
	Get_TimelineControllerPositionOffset uintptr
	Put_TimelineControllerPositionOffset uintptr
	Get_PlaybackSession                  uintptr
	StepForwardOneFrame                  uintptr
	StepBackwardOneFrame                 uintptr
	GetAsCastingSource                   uintptr
}

type IMediaPlayer3 struct {
	win32.IInspectable
}

func (this *IMediaPlayer3) Vtbl() *IMediaPlayer3Vtbl {
	return (*IMediaPlayer3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlayer3) Add_IsMutedChanged(value TypedEventHandler[*IMediaPlayer, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_IsMutedChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer3) Remove_IsMutedChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_IsMutedChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlayer3) Add_SourceChanged(value TypedEventHandler[*IMediaPlayer, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SourceChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer3) Remove_SourceChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SourceChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlayer3) Get_AudioBalance() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioBalance, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer3) Put_AudioBalance(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AudioBalance, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMediaPlayer3) Get_RealTimePlayback() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RealTimePlayback, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer3) Put_RealTimePlayback(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RealTimePlayback, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaPlayer3) Get_StereoscopicVideoRenderMode() StereoscopicVideoRenderMode {
	var _result StereoscopicVideoRenderMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StereoscopicVideoRenderMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer3) Put_StereoscopicVideoRenderMode(value StereoscopicVideoRenderMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StereoscopicVideoRenderMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMediaPlayer3) Get_BreakManager() *IMediaBreakManager {
	var _result *IMediaBreakManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BreakManager, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlayer3) Get_CommandManager() *IMediaPlaybackCommandManager {
	var _result *IMediaPlaybackCommandManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CommandManager, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlayer3) Get_AudioDevice() *IDeviceInformation {
	var _result *IDeviceInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioDevice, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlayer3) Put_AudioDevice(value *IDeviceInformation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AudioDevice, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaPlayer3) Get_TimelineController() *IMediaTimelineController {
	var _result *IMediaTimelineController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TimelineController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlayer3) Put_TimelineController(value *IMediaTimelineController) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TimelineController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaPlayer3) Get_TimelineControllerPositionOffset() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TimelineControllerPositionOffset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer3) Put_TimelineControllerPositionOffset(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TimelineControllerPositionOffset, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IMediaPlayer3) Get_PlaybackSession() *IMediaPlaybackSession {
	var _result *IMediaPlaybackSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlaybackSession, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlayer3) StepForwardOneFrame() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StepForwardOneFrame, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IMediaPlayer3) StepBackwardOneFrame() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StepBackwardOneFrame, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IMediaPlayer3) GetAsCastingSource() *ICastingSource {
	var _result *ICastingSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAsCastingSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 80035DB0-7448-4770-AFCF-2A57450914C5
var IID_IMediaPlayer4 = syscall.GUID{0x80035DB0, 0x7448, 0x4770,
	[8]byte{0xAF, 0xCF, 0x2A, 0x57, 0x45, 0x09, 0x14, 0xC5}}

type IMediaPlayer4Interface interface {
	win32.IInspectableInterface
	SetSurfaceSize(size Size)
	GetSurface(compositor *ICompositor) *IMediaPlayerSurface
}

type IMediaPlayer4Vtbl struct {
	win32.IInspectableVtbl
	SetSurfaceSize uintptr
	GetSurface     uintptr
}

type IMediaPlayer4 struct {
	win32.IInspectable
}

func (this *IMediaPlayer4) Vtbl() *IMediaPlayer4Vtbl {
	return (*IMediaPlayer4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlayer4) SetSurfaceSize(size Size) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetSurfaceSize, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&size)))
	_ = _hr
}

func (this *IMediaPlayer4) GetSurface(compositor *ICompositor) *IMediaPlayerSurface {
	var _result *IMediaPlayerSurface
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSurface, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(compositor)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CFE537FD-F86A-4446-BF4D-C8E792B7B4B3
var IID_IMediaPlayer5 = syscall.GUID{0xCFE537FD, 0xF86A, 0x4446,
	[8]byte{0xBF, 0x4D, 0xC8, 0xE7, 0x92, 0xB7, 0xB4, 0xB3}}

type IMediaPlayer5Interface interface {
	win32.IInspectableInterface
	Add_VideoFrameAvailable(value TypedEventHandler[*IMediaPlayer, interface{}]) EventRegistrationToken
	Remove_VideoFrameAvailable(token EventRegistrationToken)
	Get_IsVideoFrameServerEnabled() bool
	Put_IsVideoFrameServerEnabled(value bool)
	CopyFrameToVideoSurface(destination unsafe.Pointer)
	CopyFrameToVideoSurfaceWithTargetRectangle(destination unsafe.Pointer, targetRectangle Rect)
	CopyFrameToStereoscopicVideoSurfaces(destinationLeftEye unsafe.Pointer, destinationRightEye unsafe.Pointer)
}

type IMediaPlayer5Vtbl struct {
	win32.IInspectableVtbl
	Add_VideoFrameAvailable                    uintptr
	Remove_VideoFrameAvailable                 uintptr
	Get_IsVideoFrameServerEnabled              uintptr
	Put_IsVideoFrameServerEnabled              uintptr
	CopyFrameToVideoSurface                    uintptr
	CopyFrameToVideoSurfaceWithTargetRectangle uintptr
	CopyFrameToStereoscopicVideoSurfaces       uintptr
}

type IMediaPlayer5 struct {
	win32.IInspectable
}

func (this *IMediaPlayer5) Vtbl() *IMediaPlayer5Vtbl {
	return (*IMediaPlayer5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlayer5) Add_VideoFrameAvailable(value TypedEventHandler[*IMediaPlayer, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_VideoFrameAvailable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer5) Remove_VideoFrameAvailable(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_VideoFrameAvailable, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlayer5) Get_IsVideoFrameServerEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsVideoFrameServerEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer5) Put_IsVideoFrameServerEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsVideoFrameServerEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaPlayer5) CopyFrameToVideoSurface(destination unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CopyFrameToVideoSurface, uintptr(unsafe.Pointer(this)), uintptr(destination))
	_ = _hr
}

func (this *IMediaPlayer5) CopyFrameToVideoSurfaceWithTargetRectangle(destination unsafe.Pointer, targetRectangle Rect) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CopyFrameToVideoSurfaceWithTargetRectangle, uintptr(unsafe.Pointer(this)), uintptr(destination), uintptr(unsafe.Pointer(&targetRectangle)))
	_ = _hr
}

func (this *IMediaPlayer5) CopyFrameToStereoscopicVideoSurfaces(destinationLeftEye unsafe.Pointer, destinationRightEye unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CopyFrameToStereoscopicVideoSurfaces, uintptr(unsafe.Pointer(this)), uintptr(destinationLeftEye), uintptr(destinationRightEye))
	_ = _hr
}

// E0CAA086-AE65-414C-B010-8BC55F00E692
var IID_IMediaPlayer6 = syscall.GUID{0xE0CAA086, 0xAE65, 0x414C,
	[8]byte{0xB0, 0x10, 0x8B, 0xC5, 0x5F, 0x00, 0xE6, 0x92}}

type IMediaPlayer6Interface interface {
	win32.IInspectableInterface
	Add_SubtitleFrameChanged(handler TypedEventHandler[*IMediaPlayer, interface{}]) EventRegistrationToken
	Remove_SubtitleFrameChanged(token EventRegistrationToken)
	RenderSubtitlesToSurface(destination unsafe.Pointer) bool
	RenderSubtitlesToSurfaceWithTargetRectangle(destination unsafe.Pointer, targetRectangle Rect) bool
}

type IMediaPlayer6Vtbl struct {
	win32.IInspectableVtbl
	Add_SubtitleFrameChanged                    uintptr
	Remove_SubtitleFrameChanged                 uintptr
	RenderSubtitlesToSurface                    uintptr
	RenderSubtitlesToSurfaceWithTargetRectangle uintptr
}

type IMediaPlayer6 struct {
	win32.IInspectable
}

func (this *IMediaPlayer6) Vtbl() *IMediaPlayer6Vtbl {
	return (*IMediaPlayer6Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlayer6) Add_SubtitleFrameChanged(handler TypedEventHandler[*IMediaPlayer, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SubtitleFrameChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer6) Remove_SubtitleFrameChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SubtitleFrameChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaPlayer6) RenderSubtitlesToSurface(destination unsafe.Pointer) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RenderSubtitlesToSurface, uintptr(unsafe.Pointer(this)), uintptr(destination), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayer6) RenderSubtitlesToSurfaceWithTargetRectangle(destination unsafe.Pointer, targetRectangle Rect) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RenderSubtitlesToSurfaceWithTargetRectangle, uintptr(unsafe.Pointer(this)), uintptr(destination), uintptr(unsafe.Pointer(&targetRectangle)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5D1DC478-4500-4531-B3F4-777A71491F7F
var IID_IMediaPlayer7 = syscall.GUID{0x5D1DC478, 0x4500, 0x4531,
	[8]byte{0xB3, 0xF4, 0x77, 0x7A, 0x71, 0x49, 0x1F, 0x7F}}

type IMediaPlayer7Interface interface {
	win32.IInspectableInterface
	Get_AudioStateMonitor() *IAudioStateMonitor
}

type IMediaPlayer7Vtbl struct {
	win32.IInspectableVtbl
	Get_AudioStateMonitor uintptr
}

type IMediaPlayer7 struct {
	win32.IInspectable
}

func (this *IMediaPlayer7) Vtbl() *IMediaPlayer7Vtbl {
	return (*IMediaPlayer7Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlayer7) Get_AudioStateMonitor() *IAudioStateMonitor {
	var _result *IAudioStateMonitor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioStateMonitor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C75A9405-C801-412A-835B-83FC0E622A8E
var IID_IMediaPlayerDataReceivedEventArgs = syscall.GUID{0xC75A9405, 0xC801, 0x412A,
	[8]byte{0x83, 0x5B, 0x83, 0xFC, 0x0E, 0x62, 0x2A, 0x8E}}

type IMediaPlayerDataReceivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Data() *IPropertySet
}

type IMediaPlayerDataReceivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Data uintptr
}

type IMediaPlayerDataReceivedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaPlayerDataReceivedEventArgs) Vtbl() *IMediaPlayerDataReceivedEventArgsVtbl {
	return (*IMediaPlayerDataReceivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlayerDataReceivedEventArgs) Get_Data() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 85A1DEDA-CAB6-4CC0-8BE3-6035F4DE2591
var IID_IMediaPlayerEffects = syscall.GUID{0x85A1DEDA, 0xCAB6, 0x4CC0,
	[8]byte{0x8B, 0xE3, 0x60, 0x35, 0xF4, 0xDE, 0x25, 0x91}}

type IMediaPlayerEffectsInterface interface {
	win32.IInspectableInterface
	AddAudioEffect(activatableClassId string, effectOptional bool, configuration *IPropertySet)
	RemoveAllEffects()
}

type IMediaPlayerEffectsVtbl struct {
	win32.IInspectableVtbl
	AddAudioEffect   uintptr
	RemoveAllEffects uintptr
}

type IMediaPlayerEffects struct {
	win32.IInspectable
}

func (this *IMediaPlayerEffects) Vtbl() *IMediaPlayerEffectsVtbl {
	return (*IMediaPlayerEffectsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlayerEffects) AddAudioEffect(activatableClassId string, effectOptional bool, configuration *IPropertySet) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddAudioEffect, uintptr(unsafe.Pointer(this)), NewHStr(activatableClassId).Ptr, uintptr(*(*byte)(unsafe.Pointer(&effectOptional))), uintptr(unsafe.Pointer(configuration)))
	_ = _hr
}

func (this *IMediaPlayerEffects) RemoveAllEffects() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveAllEffects, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// FA419A79-1BBE-46C5-AE1F-8EE69FB3C2C7
var IID_IMediaPlayerEffects2 = syscall.GUID{0xFA419A79, 0x1BBE, 0x46C5,
	[8]byte{0xAE, 0x1F, 0x8E, 0xE6, 0x9F, 0xB3, 0xC2, 0xC7}}

type IMediaPlayerEffects2Interface interface {
	win32.IInspectableInterface
	AddVideoEffect(activatableClassId string, effectOptional bool, effectConfiguration *IPropertySet)
}

type IMediaPlayerEffects2Vtbl struct {
	win32.IInspectableVtbl
	AddVideoEffect uintptr
}

type IMediaPlayerEffects2 struct {
	win32.IInspectable
}

func (this *IMediaPlayerEffects2) Vtbl() *IMediaPlayerEffects2Vtbl {
	return (*IMediaPlayerEffects2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlayerEffects2) AddVideoEffect(activatableClassId string, effectOptional bool, effectConfiguration *IPropertySet) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddVideoEffect, uintptr(unsafe.Pointer(this)), NewHStr(activatableClassId).Ptr, uintptr(*(*byte)(unsafe.Pointer(&effectOptional))), uintptr(unsafe.Pointer(effectConfiguration)))
	_ = _hr
}

// 2744E9B9-A7E3-4F16-BAC4-7914EBC08301
var IID_IMediaPlayerFailedEventArgs = syscall.GUID{0x2744E9B9, 0xA7E3, 0x4F16,
	[8]byte{0xBA, 0xC4, 0x79, 0x14, 0xEB, 0xC0, 0x83, 0x01}}

type IMediaPlayerFailedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Error() MediaPlayerError
	Get_ExtendedErrorCode() HResult
	Get_ErrorMessage() string
}

type IMediaPlayerFailedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Error             uintptr
	Get_ExtendedErrorCode uintptr
	Get_ErrorMessage      uintptr
}

type IMediaPlayerFailedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaPlayerFailedEventArgs) Vtbl() *IMediaPlayerFailedEventArgsVtbl {
	return (*IMediaPlayerFailedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlayerFailedEventArgs) Get_Error() MediaPlayerError {
	var _result MediaPlayerError
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Error, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayerFailedEventArgs) Get_ExtendedErrorCode() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedErrorCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaPlayerFailedEventArgs) Get_ErrorMessage() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ErrorMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 40600D58-3B61-4BB2-989F-FC65608B6CAB
var IID_IMediaPlayerRateChangedEventArgs = syscall.GUID{0x40600D58, 0x3B61, 0x4BB2,
	[8]byte{0x98, 0x9F, 0xFC, 0x65, 0x60, 0x8B, 0x6C, 0xAB}}

type IMediaPlayerRateChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_NewRate() float64
}

type IMediaPlayerRateChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_NewRate uintptr
}

type IMediaPlayerRateChangedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaPlayerRateChangedEventArgs) Vtbl() *IMediaPlayerRateChangedEventArgsVtbl {
	return (*IMediaPlayerRateChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlayerRateChangedEventArgs) Get_NewRate() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NewRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// BD4F8897-1423-4C3E-82C5-0FB1AF94F715
var IID_IMediaPlayerSource = syscall.GUID{0xBD4F8897, 0x1423, 0x4C3E,
	[8]byte{0x82, 0xC5, 0x0F, 0xB1, 0xAF, 0x94, 0xF7, 0x15}}

type IMediaPlayerSourceInterface interface {
	win32.IInspectableInterface
	Get_ProtectionManager() *IMediaProtectionManager
	Put_ProtectionManager(value *IMediaProtectionManager)
	SetFileSource(file *IStorageFile)
	SetStreamSource(stream *IRandomAccessStream)
	SetMediaSource(source *IMediaSource)
}

type IMediaPlayerSourceVtbl struct {
	win32.IInspectableVtbl
	Get_ProtectionManager uintptr
	Put_ProtectionManager uintptr
	SetFileSource         uintptr
	SetStreamSource       uintptr
	SetMediaSource        uintptr
}

type IMediaPlayerSource struct {
	win32.IInspectable
}

func (this *IMediaPlayerSource) Vtbl() *IMediaPlayerSourceVtbl {
	return (*IMediaPlayerSourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlayerSource) Get_ProtectionManager() *IMediaProtectionManager {
	var _result *IMediaProtectionManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProtectionManager, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlayerSource) Put_ProtectionManager(value *IMediaProtectionManager) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ProtectionManager, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaPlayerSource) SetFileSource(file *IStorageFile) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetFileSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)))
	_ = _hr
}

func (this *IMediaPlayerSource) SetStreamSource(stream *IRandomAccessStream) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetStreamSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)))
	_ = _hr
}

func (this *IMediaPlayerSource) SetMediaSource(source *IMediaSource) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetMediaSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(source)))
	_ = _hr
}

// 82449B9F-7322-4C0B-B03B-3E69A48260C5
var IID_IMediaPlayerSource2 = syscall.GUID{0x82449B9F, 0x7322, 0x4C0B,
	[8]byte{0xB0, 0x3B, 0x3E, 0x69, 0xA4, 0x82, 0x60, 0xC5}}

type IMediaPlayerSource2Interface interface {
	win32.IInspectableInterface
	Get_Source() *IMediaPlaybackSource
	Put_Source(value *IMediaPlaybackSource)
}

type IMediaPlayerSource2Vtbl struct {
	win32.IInspectableVtbl
	Get_Source uintptr
	Put_Source uintptr
}

type IMediaPlayerSource2 struct {
	win32.IInspectable
}

func (this *IMediaPlayerSource2) Vtbl() *IMediaPlayerSource2Vtbl {
	return (*IMediaPlayerSource2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlayerSource2) Get_Source() *IMediaPlaybackSource {
	var _result *IMediaPlaybackSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Source, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlayerSource2) Put_Source(value *IMediaPlaybackSource) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Source, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 0ED653BC-B736-49C3-830B-764A3845313A
var IID_IMediaPlayerSurface = syscall.GUID{0x0ED653BC, 0xB736, 0x49C3,
	[8]byte{0x83, 0x0B, 0x76, 0x4A, 0x38, 0x45, 0x31, 0x3A}}

type IMediaPlayerSurfaceInterface interface {
	win32.IInspectableInterface
	Get_CompositionSurface() *ICompositionSurface
	Get_Compositor() *ICompositor
	Get_MediaPlayer() *IMediaPlayer
}

type IMediaPlayerSurfaceVtbl struct {
	win32.IInspectableVtbl
	Get_CompositionSurface uintptr
	Get_Compositor         uintptr
	Get_MediaPlayer        uintptr
}

type IMediaPlayerSurface struct {
	win32.IInspectable
}

func (this *IMediaPlayerSurface) Vtbl() *IMediaPlayerSurfaceVtbl {
	return (*IMediaPlayerSurfaceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaPlayerSurface) Get_CompositionSurface() *ICompositionSurface {
	var _result *ICompositionSurface
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CompositionSurface, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlayerSurface) Get_Compositor() *ICompositor {
	var _result *ICompositor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Compositor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaPlayerSurface) Get_MediaPlayer() *IMediaPlayer {
	var _result *IMediaPlayer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaPlayer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C4D22F5C-3C1C-4444-B6B9-778B0422D41A
var IID_IPlaybackMediaMarker = syscall.GUID{0xC4D22F5C, 0x3C1C, 0x4444,
	[8]byte{0xB6, 0xB9, 0x77, 0x8B, 0x04, 0x22, 0xD4, 0x1A}}

type IPlaybackMediaMarkerInterface interface {
	win32.IInspectableInterface
	Get_Time() TimeSpan
	Get_MediaMarkerType() string
	Get_Text() string
}

type IPlaybackMediaMarkerVtbl struct {
	win32.IInspectableVtbl
	Get_Time            uintptr
	Get_MediaMarkerType uintptr
	Get_Text            uintptr
}

type IPlaybackMediaMarker struct {
	win32.IInspectable
}

func (this *IPlaybackMediaMarker) Vtbl() *IPlaybackMediaMarkerVtbl {
	return (*IPlaybackMediaMarkerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPlaybackMediaMarker) Get_Time() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Time, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPlaybackMediaMarker) Get_MediaMarkerType() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaMarkerType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPlaybackMediaMarker) Get_Text() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Text, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 8C530A78-E0AE-4E1A-A8C8-E23F982A937B
var IID_IPlaybackMediaMarkerFactory = syscall.GUID{0x8C530A78, 0xE0AE, 0x4E1A,
	[8]byte{0xA8, 0xC8, 0xE2, 0x3F, 0x98, 0x2A, 0x93, 0x7B}}

type IPlaybackMediaMarkerFactoryInterface interface {
	win32.IInspectableInterface
	CreateFromTime(value TimeSpan) *IPlaybackMediaMarker
	Create(value TimeSpan, mediaMarketType string, text string) *IPlaybackMediaMarker
}

type IPlaybackMediaMarkerFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateFromTime uintptr
	Create         uintptr
}

type IPlaybackMediaMarkerFactory struct {
	win32.IInspectable
}

func (this *IPlaybackMediaMarkerFactory) Vtbl() *IPlaybackMediaMarkerFactoryVtbl {
	return (*IPlaybackMediaMarkerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPlaybackMediaMarkerFactory) CreateFromTime(value TimeSpan) *IPlaybackMediaMarker {
	var _result *IPlaybackMediaMarker
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromTime, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPlaybackMediaMarkerFactory) Create(value TimeSpan, mediaMarketType string, text string) *IPlaybackMediaMarker {
	var _result *IPlaybackMediaMarker
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)), NewHStr(mediaMarketType).Ptr, NewHStr(text).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 578CD1B9-90E2-4E60-ABC4-8740B01F6196
var IID_IPlaybackMediaMarkerReachedEventArgs = syscall.GUID{0x578CD1B9, 0x90E2, 0x4E60,
	[8]byte{0xAB, 0xC4, 0x87, 0x40, 0xB0, 0x1F, 0x61, 0x96}}

type IPlaybackMediaMarkerReachedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_PlaybackMediaMarker() *IPlaybackMediaMarker
}

type IPlaybackMediaMarkerReachedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_PlaybackMediaMarker uintptr
}

type IPlaybackMediaMarkerReachedEventArgs struct {
	win32.IInspectable
}

func (this *IPlaybackMediaMarkerReachedEventArgs) Vtbl() *IPlaybackMediaMarkerReachedEventArgsVtbl {
	return (*IPlaybackMediaMarkerReachedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPlaybackMediaMarkerReachedEventArgs) Get_PlaybackMediaMarker() *IPlaybackMediaMarker {
	var _result *IPlaybackMediaMarker
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlaybackMediaMarker, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F2810CEE-638B-46CF-8817-1D111FE9D8C4
var IID_IPlaybackMediaMarkerSequence = syscall.GUID{0xF2810CEE, 0x638B, 0x46CF,
	[8]byte{0x88, 0x17, 0x1D, 0x11, 0x1F, 0xE9, 0xD8, 0xC4}}

type IPlaybackMediaMarkerSequenceInterface interface {
	win32.IInspectableInterface
	Get_Size() uint32
	Insert(value *IPlaybackMediaMarker)
	Clear()
}

type IPlaybackMediaMarkerSequenceVtbl struct {
	win32.IInspectableVtbl
	Get_Size uintptr
	Insert   uintptr
	Clear    uintptr
}

type IPlaybackMediaMarkerSequence struct {
	win32.IInspectable
}

func (this *IPlaybackMediaMarkerSequence) Vtbl() *IPlaybackMediaMarkerSequenceVtbl {
	return (*IPlaybackMediaMarkerSequenceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPlaybackMediaMarkerSequence) Get_Size() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Size, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPlaybackMediaMarkerSequence) Insert(value *IPlaybackMediaMarker) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Insert, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPlaybackMediaMarkerSequence) Clear() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Clear, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// D1636099-65DF-45AE-8CEF-DC0B53FDC2BB
var IID_ITimedMetadataPresentationModeChangedEventArgs = syscall.GUID{0xD1636099, 0x65DF, 0x45AE,
	[8]byte{0x8C, 0xEF, 0xDC, 0x0B, 0x53, 0xFD, 0xC2, 0xBB}}

type ITimedMetadataPresentationModeChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Track() *ITimedMetadataTrack
	Get_OldPresentationMode() TimedMetadataTrackPresentationMode
	Get_NewPresentationMode() TimedMetadataTrackPresentationMode
}

type ITimedMetadataPresentationModeChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Track               uintptr
	Get_OldPresentationMode uintptr
	Get_NewPresentationMode uintptr
}

type ITimedMetadataPresentationModeChangedEventArgs struct {
	win32.IInspectable
}

func (this *ITimedMetadataPresentationModeChangedEventArgs) Vtbl() *ITimedMetadataPresentationModeChangedEventArgsVtbl {
	return (*ITimedMetadataPresentationModeChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimedMetadataPresentationModeChangedEventArgs) Get_Track() *ITimedMetadataTrack {
	var _result *ITimedMetadataTrack
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Track, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITimedMetadataPresentationModeChangedEventArgs) Get_OldPresentationMode() TimedMetadataTrackPresentationMode {
	var _result TimedMetadataTrackPresentationMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OldPresentationMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITimedMetadataPresentationModeChangedEventArgs) Get_NewPresentationMode() TimedMetadataTrackPresentationMode {
	var _result TimedMetadataTrackPresentationMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NewPresentationMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type CurrentMediaPlaybackItemChangedEventArgs struct {
	RtClass
	*ICurrentMediaPlaybackItemChangedEventArgs
}

type MediaBreak struct {
	RtClass
	*IMediaBreak
}

func NewMediaBreak_Create(insertionMethod MediaBreakInsertionMethod) *MediaBreak {
	hs := NewHStr("Windows.Media.Playback.MediaBreak")
	var pFac *IMediaBreakFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMediaBreakFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IMediaBreak
	p = pFac.Create(insertionMethod)
	result := &MediaBreak{
		RtClass:     RtClass{PInspect: &p.IInspectable},
		IMediaBreak: p,
	}
	com.AddToScope(result)
	return result
}

func NewMediaBreak_CreateWithPresentationPosition(insertionMethod MediaBreakInsertionMethod, presentationPosition TimeSpan) *MediaBreak {
	hs := NewHStr("Windows.Media.Playback.MediaBreak")
	var pFac *IMediaBreakFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMediaBreakFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IMediaBreak
	p = pFac.CreateWithPresentationPosition(insertionMethod, presentationPosition)
	result := &MediaBreak{
		RtClass:     RtClass{PInspect: &p.IInspectable},
		IMediaBreak: p,
	}
	com.AddToScope(result)
	return result
}

type MediaBreakEndedEventArgs struct {
	RtClass
	*IMediaBreakEndedEventArgs
}

type MediaBreakManager struct {
	RtClass
	*IMediaBreakManager
}

type MediaBreakSchedule struct {
	RtClass
	*IMediaBreakSchedule
}

type MediaBreakSeekedOverEventArgs struct {
	RtClass
	*IMediaBreakSeekedOverEventArgs
}

type MediaBreakSkippedEventArgs struct {
	RtClass
	*IMediaBreakSkippedEventArgs
}

type MediaBreakStartedEventArgs struct {
	RtClass
	*IMediaBreakStartedEventArgs
}

type MediaPlaybackAudioTrackList struct {
	RtClass
	*IVectorView[*IMediaTrack]
}

type MediaPlaybackCommandManager struct {
	RtClass
	*IMediaPlaybackCommandManager
}

type MediaPlaybackCommandManagerAutoRepeatModeReceivedEventArgs struct {
	RtClass
	*IMediaPlaybackCommandManagerAutoRepeatModeReceivedEventArgs
}

type MediaPlaybackCommandManagerCommandBehavior struct {
	RtClass
	*IMediaPlaybackCommandManagerCommandBehavior
}

type MediaPlaybackCommandManagerFastForwardReceivedEventArgs struct {
	RtClass
	*IMediaPlaybackCommandManagerFastForwardReceivedEventArgs
}

type MediaPlaybackCommandManagerNextReceivedEventArgs struct {
	RtClass
	*IMediaPlaybackCommandManagerNextReceivedEventArgs
}

type MediaPlaybackCommandManagerPauseReceivedEventArgs struct {
	RtClass
	*IMediaPlaybackCommandManagerPauseReceivedEventArgs
}

type MediaPlaybackCommandManagerPlayReceivedEventArgs struct {
	RtClass
	*IMediaPlaybackCommandManagerPlayReceivedEventArgs
}

type MediaPlaybackCommandManagerPositionReceivedEventArgs struct {
	RtClass
	*IMediaPlaybackCommandManagerPositionReceivedEventArgs
}

type MediaPlaybackCommandManagerPreviousReceivedEventArgs struct {
	RtClass
	*IMediaPlaybackCommandManagerPreviousReceivedEventArgs
}

type MediaPlaybackCommandManagerRateReceivedEventArgs struct {
	RtClass
	*IMediaPlaybackCommandManagerRateReceivedEventArgs
}

type MediaPlaybackCommandManagerRewindReceivedEventArgs struct {
	RtClass
	*IMediaPlaybackCommandManagerRewindReceivedEventArgs
}

type MediaPlaybackCommandManagerShuffleReceivedEventArgs struct {
	RtClass
	*IMediaPlaybackCommandManagerShuffleReceivedEventArgs
}

type MediaPlaybackItem struct {
	RtClass
	*IMediaPlaybackItem
}

func NewMediaPlaybackItem_CreateWithStartTime(source *IMediaSource2, startTime TimeSpan) *MediaPlaybackItem {
	hs := NewHStr("Windows.Media.Playback.MediaPlaybackItem")
	var pFac *IMediaPlaybackItemFactory2
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMediaPlaybackItemFactory2, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IMediaPlaybackItem
	p = pFac.CreateWithStartTime(source, startTime)
	result := &MediaPlaybackItem{
		RtClass:            RtClass{PInspect: &p.IInspectable},
		IMediaPlaybackItem: p,
	}
	com.AddToScope(result)
	return result
}

func NewMediaPlaybackItem_CreateWithStartTimeAndDurationLimit(source *IMediaSource2, startTime TimeSpan, durationLimit TimeSpan) *MediaPlaybackItem {
	hs := NewHStr("Windows.Media.Playback.MediaPlaybackItem")
	var pFac *IMediaPlaybackItemFactory2
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMediaPlaybackItemFactory2, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IMediaPlaybackItem
	p = pFac.CreateWithStartTimeAndDurationLimit(source, startTime, durationLimit)
	result := &MediaPlaybackItem{
		RtClass:            RtClass{PInspect: &p.IInspectable},
		IMediaPlaybackItem: p,
	}
	com.AddToScope(result)
	return result
}

func NewIMediaPlaybackItemStatics() *IMediaPlaybackItemStatics {
	var p *IMediaPlaybackItemStatics
	hs := NewHStr("Windows.Media.Playback.MediaPlaybackItem")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMediaPlaybackItemStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type MediaPlaybackItemError struct {
	RtClass
	*IMediaPlaybackItemError
}

type MediaPlaybackItemFailedEventArgs struct {
	RtClass
	*IMediaPlaybackItemFailedEventArgs
}

type MediaPlaybackItemOpenedEventArgs struct {
	RtClass
	*IMediaPlaybackItemOpenedEventArgs
}

type MediaPlaybackList struct {
	RtClass
	*IMediaPlaybackList
}

func NewMediaPlaybackList() *MediaPlaybackList {
	hs := NewHStr("Windows.Media.Playback.MediaPlaybackList")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &MediaPlaybackList{
		RtClass:            RtClass{PInspect: p},
		IMediaPlaybackList: (*IMediaPlaybackList)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type MediaPlaybackSession struct {
	RtClass
	*IMediaPlaybackSession
}

type MediaPlaybackSessionBufferingStartedEventArgs struct {
	RtClass
	*IMediaPlaybackSessionBufferingStartedEventArgs
}

type MediaPlaybackSessionOutputDegradationPolicyState struct {
	RtClass
	*IMediaPlaybackSessionOutputDegradationPolicyState
}

type MediaPlaybackSphericalVideoProjection struct {
	RtClass
	*IMediaPlaybackSphericalVideoProjection
}

type MediaPlaybackTimedMetadataTrackList struct {
	RtClass
	*IVectorView[*ITimedMetadataTrack]
}

type MediaPlaybackVideoTrackList struct {
	RtClass
	*IVectorView[*IMediaTrack]
}

type MediaPlayer struct {
	RtClass
	*IMediaPlayer
}

func NewMediaPlayer() *MediaPlayer {
	hs := NewHStr("Windows.Media.Playback.MediaPlayer")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &MediaPlayer{
		RtClass:      RtClass{PInspect: p},
		IMediaPlayer: (*IMediaPlayer)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type MediaPlayerFailedEventArgs struct {
	RtClass
	*IMediaPlayerFailedEventArgs
}

type MediaPlayerRateChangedEventArgs struct {
	RtClass
	*IMediaPlayerRateChangedEventArgs
}

type MediaPlayerSurface struct {
	RtClass
	*IMediaPlayerSurface
}

type TimedMetadataPresentationModeChangedEventArgs struct {
	RtClass
	*ITimedMetadataPresentationModeChangedEventArgs
}
