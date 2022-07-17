package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type GlobalSystemMediaTransportControlsSessionPlaybackStatus int32

const (
	GlobalSystemMediaTransportControlsSessionPlaybackStatus_Closed   GlobalSystemMediaTransportControlsSessionPlaybackStatus = 0
	GlobalSystemMediaTransportControlsSessionPlaybackStatus_Opened   GlobalSystemMediaTransportControlsSessionPlaybackStatus = 1
	GlobalSystemMediaTransportControlsSessionPlaybackStatus_Changing GlobalSystemMediaTransportControlsSessionPlaybackStatus = 2
	GlobalSystemMediaTransportControlsSessionPlaybackStatus_Stopped  GlobalSystemMediaTransportControlsSessionPlaybackStatus = 3
	GlobalSystemMediaTransportControlsSessionPlaybackStatus_Playing  GlobalSystemMediaTransportControlsSessionPlaybackStatus = 4
	GlobalSystemMediaTransportControlsSessionPlaybackStatus_Paused   GlobalSystemMediaTransportControlsSessionPlaybackStatus = 5
)

// interfaces

// 6969CB39-0BFA-5FE0-8D73-09CC5E5408E1
var IID_ICurrentSessionChangedEventArgs = syscall.GUID{0x6969CB39, 0x0BFA, 0x5FE0,
	[8]byte{0x8D, 0x73, 0x09, 0xCC, 0x5E, 0x54, 0x08, 0xE1}}

type ICurrentSessionChangedEventArgsInterface interface {
	win32.IInspectableInterface
}

type ICurrentSessionChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
}

type ICurrentSessionChangedEventArgs struct {
	win32.IInspectable
}

func (this *ICurrentSessionChangedEventArgs) Vtbl() *ICurrentSessionChangedEventArgsVtbl {
	return (*ICurrentSessionChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 7148C835-9B14-5AE2-AB85-DC9B1C14E1A8
var IID_IGlobalSystemMediaTransportControlsSession = syscall.GUID{0x7148C835, 0x9B14, 0x5AE2,
	[8]byte{0xAB, 0x85, 0xDC, 0x9B, 0x1C, 0x14, 0xE1, 0xA8}}

type IGlobalSystemMediaTransportControlsSessionInterface interface {
	win32.IInspectableInterface
	Get_SourceAppUserModelId() string
	TryGetMediaPropertiesAsync() *IAsyncOperation[*IGlobalSystemMediaTransportControlsSessionMediaProperties]
	GetTimelineProperties() *IGlobalSystemMediaTransportControlsSessionTimelineProperties
	GetPlaybackInfo() *IGlobalSystemMediaTransportControlsSessionPlaybackInfo
	TryPlayAsync() *IAsyncOperation[bool]
	TryPauseAsync() *IAsyncOperation[bool]
	TryStopAsync() *IAsyncOperation[bool]
	TryRecordAsync() *IAsyncOperation[bool]
	TryFastForwardAsync() *IAsyncOperation[bool]
	TryRewindAsync() *IAsyncOperation[bool]
	TrySkipNextAsync() *IAsyncOperation[bool]
	TrySkipPreviousAsync() *IAsyncOperation[bool]
	TryChangeChannelUpAsync() *IAsyncOperation[bool]
	TryChangeChannelDownAsync() *IAsyncOperation[bool]
	TryTogglePlayPauseAsync() *IAsyncOperation[bool]
	TryChangeAutoRepeatModeAsync(requestedAutoRepeatMode MediaPlaybackAutoRepeatMode) *IAsyncOperation[bool]
	TryChangePlaybackRateAsync(requestedPlaybackRate float64) *IAsyncOperation[bool]
	TryChangeShuffleActiveAsync(requestedShuffleState bool) *IAsyncOperation[bool]
	TryChangePlaybackPositionAsync(requestedPlaybackPosition int64) *IAsyncOperation[bool]
	Add_TimelinePropertiesChanged(handler TypedEventHandler[*IGlobalSystemMediaTransportControlsSession, *ITimelinePropertiesChangedEventArgs]) EventRegistrationToken
	Remove_TimelinePropertiesChanged(token EventRegistrationToken)
	Add_PlaybackInfoChanged(handler TypedEventHandler[*IGlobalSystemMediaTransportControlsSession, *IPlaybackInfoChangedEventArgs]) EventRegistrationToken
	Remove_PlaybackInfoChanged(token EventRegistrationToken)
	Add_MediaPropertiesChanged(handler TypedEventHandler[*IGlobalSystemMediaTransportControlsSession, *IMediaPropertiesChangedEventArgs]) EventRegistrationToken
	Remove_MediaPropertiesChanged(token EventRegistrationToken)
}

type IGlobalSystemMediaTransportControlsSessionVtbl struct {
	win32.IInspectableVtbl
	Get_SourceAppUserModelId         uintptr
	TryGetMediaPropertiesAsync       uintptr
	GetTimelineProperties            uintptr
	GetPlaybackInfo                  uintptr
	TryPlayAsync                     uintptr
	TryPauseAsync                    uintptr
	TryStopAsync                     uintptr
	TryRecordAsync                   uintptr
	TryFastForwardAsync              uintptr
	TryRewindAsync                   uintptr
	TrySkipNextAsync                 uintptr
	TrySkipPreviousAsync             uintptr
	TryChangeChannelUpAsync          uintptr
	TryChangeChannelDownAsync        uintptr
	TryTogglePlayPauseAsync          uintptr
	TryChangeAutoRepeatModeAsync     uintptr
	TryChangePlaybackRateAsync       uintptr
	TryChangeShuffleActiveAsync      uintptr
	TryChangePlaybackPositionAsync   uintptr
	Add_TimelinePropertiesChanged    uintptr
	Remove_TimelinePropertiesChanged uintptr
	Add_PlaybackInfoChanged          uintptr
	Remove_PlaybackInfoChanged       uintptr
	Add_MediaPropertiesChanged       uintptr
	Remove_MediaPropertiesChanged    uintptr
}

type IGlobalSystemMediaTransportControlsSession struct {
	win32.IInspectable
}

func (this *IGlobalSystemMediaTransportControlsSession) Vtbl() *IGlobalSystemMediaTransportControlsSessionVtbl {
	return (*IGlobalSystemMediaTransportControlsSessionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGlobalSystemMediaTransportControlsSession) Get_SourceAppUserModelId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SourceAppUserModelId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGlobalSystemMediaTransportControlsSession) TryGetMediaPropertiesAsync() *IAsyncOperation[*IGlobalSystemMediaTransportControlsSessionMediaProperties] {
	var _result *IAsyncOperation[*IGlobalSystemMediaTransportControlsSessionMediaProperties]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetMediaPropertiesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSession) GetTimelineProperties() *IGlobalSystemMediaTransportControlsSessionTimelineProperties {
	var _result *IGlobalSystemMediaTransportControlsSessionTimelineProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetTimelineProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSession) GetPlaybackInfo() *IGlobalSystemMediaTransportControlsSessionPlaybackInfo {
	var _result *IGlobalSystemMediaTransportControlsSessionPlaybackInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPlaybackInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSession) TryPlayAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryPlayAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSession) TryPauseAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryPauseAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSession) TryStopAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryStopAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSession) TryRecordAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryRecordAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSession) TryFastForwardAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryFastForwardAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSession) TryRewindAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryRewindAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSession) TrySkipNextAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TrySkipNextAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSession) TrySkipPreviousAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TrySkipPreviousAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSession) TryChangeChannelUpAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryChangeChannelUpAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSession) TryChangeChannelDownAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryChangeChannelDownAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSession) TryTogglePlayPauseAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryTogglePlayPauseAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSession) TryChangeAutoRepeatModeAsync(requestedAutoRepeatMode MediaPlaybackAutoRepeatMode) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryChangeAutoRepeatModeAsync, uintptr(unsafe.Pointer(this)), uintptr(requestedAutoRepeatMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSession) TryChangePlaybackRateAsync(requestedPlaybackRate float64) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryChangePlaybackRateAsync, uintptr(unsafe.Pointer(this)), uintptr(requestedPlaybackRate), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSession) TryChangeShuffleActiveAsync(requestedShuffleState bool) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryChangeShuffleActiveAsync, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&requestedShuffleState))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSession) TryChangePlaybackPositionAsync(requestedPlaybackPosition int64) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryChangePlaybackPositionAsync, uintptr(unsafe.Pointer(this)), uintptr(requestedPlaybackPosition), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSession) Add_TimelinePropertiesChanged(handler TypedEventHandler[*IGlobalSystemMediaTransportControlsSession, *ITimelinePropertiesChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_TimelinePropertiesChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSession) Remove_TimelinePropertiesChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_TimelinePropertiesChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGlobalSystemMediaTransportControlsSession) Add_PlaybackInfoChanged(handler TypedEventHandler[*IGlobalSystemMediaTransportControlsSession, *IPlaybackInfoChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PlaybackInfoChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSession) Remove_PlaybackInfoChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PlaybackInfoChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGlobalSystemMediaTransportControlsSession) Add_MediaPropertiesChanged(handler TypedEventHandler[*IGlobalSystemMediaTransportControlsSession, *IMediaPropertiesChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_MediaPropertiesChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSession) Remove_MediaPropertiesChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_MediaPropertiesChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// CACE8EAC-E86E-504A-AB31-5FF8FF1BCE49
var IID_IGlobalSystemMediaTransportControlsSessionManager = syscall.GUID{0xCACE8EAC, 0xE86E, 0x504A,
	[8]byte{0xAB, 0x31, 0x5F, 0xF8, 0xFF, 0x1B, 0xCE, 0x49}}

type IGlobalSystemMediaTransportControlsSessionManagerInterface interface {
	win32.IInspectableInterface
	GetCurrentSession() *IGlobalSystemMediaTransportControlsSession
	GetSessions() *IVectorView[*IGlobalSystemMediaTransportControlsSession]
	Add_CurrentSessionChanged(handler TypedEventHandler[*IGlobalSystemMediaTransportControlsSessionManager, *ICurrentSessionChangedEventArgs]) EventRegistrationToken
	Remove_CurrentSessionChanged(token EventRegistrationToken)
	Add_SessionsChanged(handler TypedEventHandler[*IGlobalSystemMediaTransportControlsSessionManager, *ISessionsChangedEventArgs]) EventRegistrationToken
	Remove_SessionsChanged(token EventRegistrationToken)
}

type IGlobalSystemMediaTransportControlsSessionManagerVtbl struct {
	win32.IInspectableVtbl
	GetCurrentSession            uintptr
	GetSessions                  uintptr
	Add_CurrentSessionChanged    uintptr
	Remove_CurrentSessionChanged uintptr
	Add_SessionsChanged          uintptr
	Remove_SessionsChanged       uintptr
}

type IGlobalSystemMediaTransportControlsSessionManager struct {
	win32.IInspectable
}

func (this *IGlobalSystemMediaTransportControlsSessionManager) Vtbl() *IGlobalSystemMediaTransportControlsSessionManagerVtbl {
	return (*IGlobalSystemMediaTransportControlsSessionManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGlobalSystemMediaTransportControlsSessionManager) GetCurrentSession() *IGlobalSystemMediaTransportControlsSession {
	var _result *IGlobalSystemMediaTransportControlsSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentSession, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionManager) GetSessions() *IVectorView[*IGlobalSystemMediaTransportControlsSession] {
	var _result *IVectorView[*IGlobalSystemMediaTransportControlsSession]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSessions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionManager) Add_CurrentSessionChanged(handler TypedEventHandler[*IGlobalSystemMediaTransportControlsSessionManager, *ICurrentSessionChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_CurrentSessionChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionManager) Remove_CurrentSessionChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_CurrentSessionChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGlobalSystemMediaTransportControlsSessionManager) Add_SessionsChanged(handler TypedEventHandler[*IGlobalSystemMediaTransportControlsSessionManager, *ISessionsChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SessionsChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionManager) Remove_SessionsChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SessionsChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 2050C4EE-11A0-57DE-AED7-C97C70338245
var IID_IGlobalSystemMediaTransportControlsSessionManagerStatics = syscall.GUID{0x2050C4EE, 0x11A0, 0x57DE,
	[8]byte{0xAE, 0xD7, 0xC9, 0x7C, 0x70, 0x33, 0x82, 0x45}}

type IGlobalSystemMediaTransportControlsSessionManagerStaticsInterface interface {
	win32.IInspectableInterface
	RequestAsync() *IAsyncOperation[*IGlobalSystemMediaTransportControlsSessionManager]
}

type IGlobalSystemMediaTransportControlsSessionManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	RequestAsync uintptr
}

type IGlobalSystemMediaTransportControlsSessionManagerStatics struct {
	win32.IInspectable
}

func (this *IGlobalSystemMediaTransportControlsSessionManagerStatics) Vtbl() *IGlobalSystemMediaTransportControlsSessionManagerStaticsVtbl {
	return (*IGlobalSystemMediaTransportControlsSessionManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGlobalSystemMediaTransportControlsSessionManagerStatics) RequestAsync() *IAsyncOperation[*IGlobalSystemMediaTransportControlsSessionManager] {
	var _result *IAsyncOperation[*IGlobalSystemMediaTransportControlsSessionManager]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 68856CF6-ADB4-54B2-AC16-05837907ACB6
var IID_IGlobalSystemMediaTransportControlsSessionMediaProperties = syscall.GUID{0x68856CF6, 0xADB4, 0x54B2,
	[8]byte{0xAC, 0x16, 0x05, 0x83, 0x79, 0x07, 0xAC, 0xB6}}

type IGlobalSystemMediaTransportControlsSessionMediaPropertiesInterface interface {
	win32.IInspectableInterface
	Get_Title() string
	Get_Subtitle() string
	Get_AlbumArtist() string
	Get_Artist() string
	Get_AlbumTitle() string
	Get_TrackNumber() int32
	Get_Genres() *IVectorView[string]
	Get_AlbumTrackCount() int32
	Get_PlaybackType() *IReference[MediaPlaybackType]
	Get_Thumbnail() *IRandomAccessStreamReference
}

type IGlobalSystemMediaTransportControlsSessionMediaPropertiesVtbl struct {
	win32.IInspectableVtbl
	Get_Title           uintptr
	Get_Subtitle        uintptr
	Get_AlbumArtist     uintptr
	Get_Artist          uintptr
	Get_AlbumTitle      uintptr
	Get_TrackNumber     uintptr
	Get_Genres          uintptr
	Get_AlbumTrackCount uintptr
	Get_PlaybackType    uintptr
	Get_Thumbnail       uintptr
}

type IGlobalSystemMediaTransportControlsSessionMediaProperties struct {
	win32.IInspectable
}

func (this *IGlobalSystemMediaTransportControlsSessionMediaProperties) Vtbl() *IGlobalSystemMediaTransportControlsSessionMediaPropertiesVtbl {
	return (*IGlobalSystemMediaTransportControlsSessionMediaPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGlobalSystemMediaTransportControlsSessionMediaProperties) Get_Title() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Title, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGlobalSystemMediaTransportControlsSessionMediaProperties) Get_Subtitle() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Subtitle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGlobalSystemMediaTransportControlsSessionMediaProperties) Get_AlbumArtist() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlbumArtist, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGlobalSystemMediaTransportControlsSessionMediaProperties) Get_Artist() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Artist, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGlobalSystemMediaTransportControlsSessionMediaProperties) Get_AlbumTitle() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlbumTitle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGlobalSystemMediaTransportControlsSessionMediaProperties) Get_TrackNumber() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TrackNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionMediaProperties) Get_Genres() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Genres, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionMediaProperties) Get_AlbumTrackCount() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlbumTrackCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionMediaProperties) Get_PlaybackType() *IReference[MediaPlaybackType] {
	var _result *IReference[MediaPlaybackType]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlaybackType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionMediaProperties) Get_Thumbnail() *IRandomAccessStreamReference {
	var _result *IRandomAccessStreamReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Thumbnail, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6501A3E6-BC7A-503A-BB1B-68F158F3FB03
var IID_IGlobalSystemMediaTransportControlsSessionPlaybackControls = syscall.GUID{0x6501A3E6, 0xBC7A, 0x503A,
	[8]byte{0xBB, 0x1B, 0x68, 0xF1, 0x58, 0xF3, 0xFB, 0x03}}

type IGlobalSystemMediaTransportControlsSessionPlaybackControlsInterface interface {
	win32.IInspectableInterface
	Get_IsPlayEnabled() bool
	Get_IsPauseEnabled() bool
	Get_IsStopEnabled() bool
	Get_IsRecordEnabled() bool
	Get_IsFastForwardEnabled() bool
	Get_IsRewindEnabled() bool
	Get_IsNextEnabled() bool
	Get_IsPreviousEnabled() bool
	Get_IsChannelUpEnabled() bool
	Get_IsChannelDownEnabled() bool
	Get_IsPlayPauseToggleEnabled() bool
	Get_IsShuffleEnabled() bool
	Get_IsRepeatEnabled() bool
	Get_IsPlaybackRateEnabled() bool
	Get_IsPlaybackPositionEnabled() bool
}

type IGlobalSystemMediaTransportControlsSessionPlaybackControlsVtbl struct {
	win32.IInspectableVtbl
	Get_IsPlayEnabled             uintptr
	Get_IsPauseEnabled            uintptr
	Get_IsStopEnabled             uintptr
	Get_IsRecordEnabled           uintptr
	Get_IsFastForwardEnabled      uintptr
	Get_IsRewindEnabled           uintptr
	Get_IsNextEnabled             uintptr
	Get_IsPreviousEnabled         uintptr
	Get_IsChannelUpEnabled        uintptr
	Get_IsChannelDownEnabled      uintptr
	Get_IsPlayPauseToggleEnabled  uintptr
	Get_IsShuffleEnabled          uintptr
	Get_IsRepeatEnabled           uintptr
	Get_IsPlaybackRateEnabled     uintptr
	Get_IsPlaybackPositionEnabled uintptr
}

type IGlobalSystemMediaTransportControlsSessionPlaybackControls struct {
	win32.IInspectable
}

func (this *IGlobalSystemMediaTransportControlsSessionPlaybackControls) Vtbl() *IGlobalSystemMediaTransportControlsSessionPlaybackControlsVtbl {
	return (*IGlobalSystemMediaTransportControlsSessionPlaybackControlsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGlobalSystemMediaTransportControlsSessionPlaybackControls) Get_IsPlayEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPlayEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionPlaybackControls) Get_IsPauseEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPauseEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionPlaybackControls) Get_IsStopEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsStopEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionPlaybackControls) Get_IsRecordEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsRecordEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionPlaybackControls) Get_IsFastForwardEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsFastForwardEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionPlaybackControls) Get_IsRewindEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsRewindEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionPlaybackControls) Get_IsNextEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsNextEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionPlaybackControls) Get_IsPreviousEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPreviousEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionPlaybackControls) Get_IsChannelUpEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsChannelUpEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionPlaybackControls) Get_IsChannelDownEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsChannelDownEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionPlaybackControls) Get_IsPlayPauseToggleEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPlayPauseToggleEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionPlaybackControls) Get_IsShuffleEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsShuffleEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionPlaybackControls) Get_IsRepeatEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsRepeatEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionPlaybackControls) Get_IsPlaybackRateEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPlaybackRateEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionPlaybackControls) Get_IsPlaybackPositionEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPlaybackPositionEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 94B4B6CF-E8BA-51AD-87A7-C10ADE106127
var IID_IGlobalSystemMediaTransportControlsSessionPlaybackInfo = syscall.GUID{0x94B4B6CF, 0xE8BA, 0x51AD,
	[8]byte{0x87, 0xA7, 0xC1, 0x0A, 0xDE, 0x10, 0x61, 0x27}}

type IGlobalSystemMediaTransportControlsSessionPlaybackInfoInterface interface {
	win32.IInspectableInterface
	Get_Controls() *IGlobalSystemMediaTransportControlsSessionPlaybackControls
	Get_PlaybackStatus() GlobalSystemMediaTransportControlsSessionPlaybackStatus
	Get_PlaybackType() *IReference[MediaPlaybackType]
	Get_AutoRepeatMode() *IReference[MediaPlaybackAutoRepeatMode]
	Get_PlaybackRate() *IReference[float64]
	Get_IsShuffleActive() *IReference[bool]
}

type IGlobalSystemMediaTransportControlsSessionPlaybackInfoVtbl struct {
	win32.IInspectableVtbl
	Get_Controls        uintptr
	Get_PlaybackStatus  uintptr
	Get_PlaybackType    uintptr
	Get_AutoRepeatMode  uintptr
	Get_PlaybackRate    uintptr
	Get_IsShuffleActive uintptr
}

type IGlobalSystemMediaTransportControlsSessionPlaybackInfo struct {
	win32.IInspectable
}

func (this *IGlobalSystemMediaTransportControlsSessionPlaybackInfo) Vtbl() *IGlobalSystemMediaTransportControlsSessionPlaybackInfoVtbl {
	return (*IGlobalSystemMediaTransportControlsSessionPlaybackInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGlobalSystemMediaTransportControlsSessionPlaybackInfo) Get_Controls() *IGlobalSystemMediaTransportControlsSessionPlaybackControls {
	var _result *IGlobalSystemMediaTransportControlsSessionPlaybackControls
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Controls, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionPlaybackInfo) Get_PlaybackStatus() GlobalSystemMediaTransportControlsSessionPlaybackStatus {
	var _result GlobalSystemMediaTransportControlsSessionPlaybackStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlaybackStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionPlaybackInfo) Get_PlaybackType() *IReference[MediaPlaybackType] {
	var _result *IReference[MediaPlaybackType]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlaybackType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionPlaybackInfo) Get_AutoRepeatMode() *IReference[MediaPlaybackAutoRepeatMode] {
	var _result *IReference[MediaPlaybackAutoRepeatMode]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AutoRepeatMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionPlaybackInfo) Get_PlaybackRate() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlaybackRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionPlaybackInfo) Get_IsShuffleActive() *IReference[bool] {
	var _result *IReference[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsShuffleActive, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EDE34136-6F25-588D-8ECF-EA5B6735AAA5
var IID_IGlobalSystemMediaTransportControlsSessionTimelineProperties = syscall.GUID{0xEDE34136, 0x6F25, 0x588D,
	[8]byte{0x8E, 0xCF, 0xEA, 0x5B, 0x67, 0x35, 0xAA, 0xA5}}

type IGlobalSystemMediaTransportControlsSessionTimelinePropertiesInterface interface {
	win32.IInspectableInterface
	Get_StartTime() TimeSpan
	Get_EndTime() TimeSpan
	Get_MinSeekTime() TimeSpan
	Get_MaxSeekTime() TimeSpan
	Get_Position() TimeSpan
	Get_LastUpdatedTime() DateTime
}

type IGlobalSystemMediaTransportControlsSessionTimelinePropertiesVtbl struct {
	win32.IInspectableVtbl
	Get_StartTime       uintptr
	Get_EndTime         uintptr
	Get_MinSeekTime     uintptr
	Get_MaxSeekTime     uintptr
	Get_Position        uintptr
	Get_LastUpdatedTime uintptr
}

type IGlobalSystemMediaTransportControlsSessionTimelineProperties struct {
	win32.IInspectable
}

func (this *IGlobalSystemMediaTransportControlsSessionTimelineProperties) Vtbl() *IGlobalSystemMediaTransportControlsSessionTimelinePropertiesVtbl {
	return (*IGlobalSystemMediaTransportControlsSessionTimelinePropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGlobalSystemMediaTransportControlsSessionTimelineProperties) Get_StartTime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StartTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionTimelineProperties) Get_EndTime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EndTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionTimelineProperties) Get_MinSeekTime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinSeekTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionTimelineProperties) Get_MaxSeekTime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxSeekTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionTimelineProperties) Get_Position() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGlobalSystemMediaTransportControlsSessionTimelineProperties) Get_LastUpdatedTime() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LastUpdatedTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7D3741CB-ADF0-5CEF-91BA-CFABCDD77678
var IID_IMediaPropertiesChangedEventArgs = syscall.GUID{0x7D3741CB, 0xADF0, 0x5CEF,
	[8]byte{0x91, 0xBA, 0xCF, 0xAB, 0xCD, 0xD7, 0x76, 0x78}}

type IMediaPropertiesChangedEventArgsInterface interface {
	win32.IInspectableInterface
}

type IMediaPropertiesChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
}

type IMediaPropertiesChangedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaPropertiesChangedEventArgs) Vtbl() *IMediaPropertiesChangedEventArgsVtbl {
	return (*IMediaPropertiesChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 786756C2-BC0D-50A5-8807-054291FEF139
var IID_IPlaybackInfoChangedEventArgs = syscall.GUID{0x786756C2, 0xBC0D, 0x50A5,
	[8]byte{0x88, 0x07, 0x05, 0x42, 0x91, 0xFE, 0xF1, 0x39}}

type IPlaybackInfoChangedEventArgsInterface interface {
	win32.IInspectableInterface
}

type IPlaybackInfoChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
}

type IPlaybackInfoChangedEventArgs struct {
	win32.IInspectable
}

func (this *IPlaybackInfoChangedEventArgs) Vtbl() *IPlaybackInfoChangedEventArgsVtbl {
	return (*IPlaybackInfoChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// BBF0CD32-42C4-5A58-B317-F34BBFBD26E0
var IID_ISessionsChangedEventArgs = syscall.GUID{0xBBF0CD32, 0x42C4, 0x5A58,
	[8]byte{0xB3, 0x17, 0xF3, 0x4B, 0xBF, 0xBD, 0x26, 0xE0}}

type ISessionsChangedEventArgsInterface interface {
	win32.IInspectableInterface
}

type ISessionsChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
}

type ISessionsChangedEventArgs struct {
	win32.IInspectable
}

func (this *ISessionsChangedEventArgs) Vtbl() *ISessionsChangedEventArgsVtbl {
	return (*ISessionsChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 29033A2F-C923-5A77-BCAF-055FF415AD32
var IID_ITimelinePropertiesChangedEventArgs = syscall.GUID{0x29033A2F, 0xC923, 0x5A77,
	[8]byte{0xBC, 0xAF, 0x05, 0x5F, 0xF4, 0x15, 0xAD, 0x32}}

type ITimelinePropertiesChangedEventArgsInterface interface {
	win32.IInspectableInterface
}

type ITimelinePropertiesChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
}

type ITimelinePropertiesChangedEventArgs struct {
	win32.IInspectable
}

func (this *ITimelinePropertiesChangedEventArgs) Vtbl() *ITimelinePropertiesChangedEventArgsVtbl {
	return (*ITimelinePropertiesChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// classes

type CurrentSessionChangedEventArgs struct {
	RtClass
	*ICurrentSessionChangedEventArgs
}

type GlobalSystemMediaTransportControlsSession struct {
	RtClass
	*IGlobalSystemMediaTransportControlsSession
}

type GlobalSystemMediaTransportControlsSessionManager struct {
	RtClass
	*IGlobalSystemMediaTransportControlsSessionManager
}

func NewIGlobalSystemMediaTransportControlsSessionManagerStatics() *IGlobalSystemMediaTransportControlsSessionManagerStatics {
	var p *IGlobalSystemMediaTransportControlsSessionManagerStatics
	hs := NewHStr("Windows.Media.Control.GlobalSystemMediaTransportControlsSessionManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGlobalSystemMediaTransportControlsSessionManagerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type GlobalSystemMediaTransportControlsSessionMediaProperties struct {
	RtClass
	*IGlobalSystemMediaTransportControlsSessionMediaProperties
}

type GlobalSystemMediaTransportControlsSessionPlaybackControls struct {
	RtClass
	*IGlobalSystemMediaTransportControlsSessionPlaybackControls
}

type GlobalSystemMediaTransportControlsSessionPlaybackInfo struct {
	RtClass
	*IGlobalSystemMediaTransportControlsSessionPlaybackInfo
}

type GlobalSystemMediaTransportControlsSessionTimelineProperties struct {
	RtClass
	*IGlobalSystemMediaTransportControlsSessionTimelineProperties
}

type MediaPropertiesChangedEventArgs struct {
	RtClass
	*IMediaPropertiesChangedEventArgs
}

type PlaybackInfoChangedEventArgs struct {
	RtClass
	*IPlaybackInfoChangedEventArgs
}

type SessionsChangedEventArgs struct {
	RtClass
	*ISessionsChangedEventArgs
}

type TimelinePropertiesChangedEventArgs struct {
	RtClass
	*ITimelinePropertiesChangedEventArgs
}
