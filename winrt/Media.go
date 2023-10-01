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
type AudioBufferAccessMode int32

const (
	AudioBufferAccessMode_Read      AudioBufferAccessMode = 0
	AudioBufferAccessMode_ReadWrite AudioBufferAccessMode = 1
	AudioBufferAccessMode_Write     AudioBufferAccessMode = 2
)

// enum
type AudioProcessing int32

const (
	AudioProcessing_Default AudioProcessing = 0
	AudioProcessing_Raw     AudioProcessing = 1
)

// enum
type MediaPlaybackAutoRepeatMode int32

const (
	MediaPlaybackAutoRepeatMode_None  MediaPlaybackAutoRepeatMode = 0
	MediaPlaybackAutoRepeatMode_Track MediaPlaybackAutoRepeatMode = 1
	MediaPlaybackAutoRepeatMode_List  MediaPlaybackAutoRepeatMode = 2
)

// enum
type MediaPlaybackStatus int32

const (
	MediaPlaybackStatus_Closed   MediaPlaybackStatus = 0
	MediaPlaybackStatus_Changing MediaPlaybackStatus = 1
	MediaPlaybackStatus_Stopped  MediaPlaybackStatus = 2
	MediaPlaybackStatus_Playing  MediaPlaybackStatus = 3
	MediaPlaybackStatus_Paused   MediaPlaybackStatus = 4
)

// enum
type MediaPlaybackType int32

const (
	MediaPlaybackType_Unknown MediaPlaybackType = 0
	MediaPlaybackType_Music   MediaPlaybackType = 1
	MediaPlaybackType_Video   MediaPlaybackType = 2
	MediaPlaybackType_Image   MediaPlaybackType = 3
)

// enum
type MediaTimelineControllerState int32

const (
	MediaTimelineControllerState_Paused  MediaTimelineControllerState = 0
	MediaTimelineControllerState_Running MediaTimelineControllerState = 1
	MediaTimelineControllerState_Stalled MediaTimelineControllerState = 2
	MediaTimelineControllerState_Error   MediaTimelineControllerState = 3
)

// enum
type SoundLevel int32

const (
	SoundLevel_Muted SoundLevel = 0
	SoundLevel_Low   SoundLevel = 1
	SoundLevel_Full  SoundLevel = 2
)

// enum
type SystemMediaTransportControlsButton int32

const (
	SystemMediaTransportControlsButton_Play        SystemMediaTransportControlsButton = 0
	SystemMediaTransportControlsButton_Pause       SystemMediaTransportControlsButton = 1
	SystemMediaTransportControlsButton_Stop        SystemMediaTransportControlsButton = 2
	SystemMediaTransportControlsButton_Record      SystemMediaTransportControlsButton = 3
	SystemMediaTransportControlsButton_FastForward SystemMediaTransportControlsButton = 4
	SystemMediaTransportControlsButton_Rewind      SystemMediaTransportControlsButton = 5
	SystemMediaTransportControlsButton_Next        SystemMediaTransportControlsButton = 6
	SystemMediaTransportControlsButton_Previous    SystemMediaTransportControlsButton = 7
	SystemMediaTransportControlsButton_ChannelUp   SystemMediaTransportControlsButton = 8
	SystemMediaTransportControlsButton_ChannelDown SystemMediaTransportControlsButton = 9
)

// enum
type SystemMediaTransportControlsProperty int32

const (
	SystemMediaTransportControlsProperty_SoundLevel SystemMediaTransportControlsProperty = 0
)

// structs

type MediaControlContract struct {
}

type MediaTimeRange struct {
	Start TimeSpan
	End   TimeSpan
}

// interfaces

// 35175827-724B-4C6A-B130-F6537F9AE0D0
var IID_IAudioBuffer = syscall.GUID{0x35175827, 0x724B, 0x4C6A,
	[8]byte{0xB1, 0x30, 0xF6, 0x53, 0x7F, 0x9A, 0xE0, 0xD0}}

type IAudioBufferInterface interface {
	win32.IInspectableInterface
	Get_Capacity() uint32
	Get_Length() uint32
	Put_Length(value uint32)
}

type IAudioBufferVtbl struct {
	win32.IInspectableVtbl
	Get_Capacity uintptr
	Get_Length   uintptr
	Put_Length   uintptr
}

type IAudioBuffer struct {
	win32.IInspectable
}

func (this *IAudioBuffer) Vtbl() *IAudioBufferVtbl {
	return (*IAudioBufferVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioBuffer) Get_Capacity() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Capacity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioBuffer) Get_Length() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Length, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioBuffer) Put_Length(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Length, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// E36AC304-AAB2-4277-9ED0-43CEDF8E29C6
var IID_IAudioFrame = syscall.GUID{0xE36AC304, 0xAAB2, 0x4277,
	[8]byte{0x9E, 0xD0, 0x43, 0xCE, 0xDF, 0x8E, 0x29, 0xC6}}

type IAudioFrameInterface interface {
	win32.IInspectableInterface
	LockBuffer(mode AudioBufferAccessMode) *IAudioBuffer
}

type IAudioFrameVtbl struct {
	win32.IInspectableVtbl
	LockBuffer uintptr
}

type IAudioFrame struct {
	win32.IInspectable
}

func (this *IAudioFrame) Vtbl() *IAudioFrameVtbl {
	return (*IAudioFrameVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioFrame) LockBuffer(mode AudioBufferAccessMode) *IAudioBuffer {
	var _result *IAudioBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LockBuffer, uintptr(unsafe.Pointer(this)), uintptr(mode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 91A90ADE-2422-40A6-B9AD-30D02404317D
var IID_IAudioFrameFactory = syscall.GUID{0x91A90ADE, 0x2422, 0x40A6,
	[8]byte{0xB9, 0xAD, 0x30, 0xD0, 0x24, 0x04, 0x31, 0x7D}}

type IAudioFrameFactoryInterface interface {
	win32.IInspectableInterface
	Create(capacity uint32) *IAudioFrame
}

type IAudioFrameFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IAudioFrameFactory struct {
	win32.IInspectable
}

func (this *IAudioFrameFactory) Vtbl() *IAudioFrameFactoryVtbl {
	return (*IAudioFrameFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioFrameFactory) Create(capacity uint32) *IAudioFrame {
	var _result *IAudioFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(capacity), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EA137EFA-D852-438E-882B-C990109A78F4
var IID_IAutoRepeatModeChangeRequestedEventArgs = syscall.GUID{0xEA137EFA, 0xD852, 0x438E,
	[8]byte{0x88, 0x2B, 0xC9, 0x90, 0x10, 0x9A, 0x78, 0xF4}}

type IAutoRepeatModeChangeRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_RequestedAutoRepeatMode() MediaPlaybackAutoRepeatMode
}

type IAutoRepeatModeChangeRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_RequestedAutoRepeatMode uintptr
}

type IAutoRepeatModeChangeRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IAutoRepeatModeChangeRequestedEventArgs) Vtbl() *IAutoRepeatModeChangeRequestedEventArgsVtbl {
	return (*IAutoRepeatModeChangeRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAutoRepeatModeChangeRequestedEventArgs) Get_RequestedAutoRepeatMode() MediaPlaybackAutoRepeatMode {
	var _result MediaPlaybackAutoRepeatMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestedAutoRepeatMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// CD0BC7EF-54E7-411F-9933-F0E98B0A96D2
var IID_IImageDisplayProperties = syscall.GUID{0xCD0BC7EF, 0x54E7, 0x411F,
	[8]byte{0x99, 0x33, 0xF0, 0xE9, 0x8B, 0x0A, 0x96, 0xD2}}

type IImageDisplayPropertiesInterface interface {
	win32.IInspectableInterface
	Get_Title() string
	Put_Title(value string)
	Get_Subtitle() string
	Put_Subtitle(value string)
}

type IImageDisplayPropertiesVtbl struct {
	win32.IInspectableVtbl
	Get_Title    uintptr
	Put_Title    uintptr
	Get_Subtitle uintptr
	Put_Subtitle uintptr
}

type IImageDisplayProperties struct {
	win32.IInspectable
}

func (this *IImageDisplayProperties) Vtbl() *IImageDisplayPropertiesVtbl {
	return (*IImageDisplayPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IImageDisplayProperties) Get_Title() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Title, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IImageDisplayProperties) Put_Title(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Title, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IImageDisplayProperties) Get_Subtitle() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Subtitle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IImageDisplayProperties) Put_Subtitle(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Subtitle, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// 98F1FBE1-7A8D-42CB-B6FE-8FE698264F13
var IID_IMediaControl = syscall.GUID{0x98F1FBE1, 0x7A8D, 0x42CB,
	[8]byte{0xB6, 0xFE, 0x8F, 0xE6, 0x98, 0x26, 0x4F, 0x13}}

type IMediaControlInterface interface {
	win32.IInspectableInterface
	Add_SoundLevelChanged(handler EventHandler[interface{}]) EventRegistrationToken
	Remove_SoundLevelChanged(cookie EventRegistrationToken)
	Add_PlayPressed(handler EventHandler[interface{}]) EventRegistrationToken
	Remove_PlayPressed(cookie EventRegistrationToken)
	Add_PausePressed(handler EventHandler[interface{}]) EventRegistrationToken
	Remove_PausePressed(cookie EventRegistrationToken)
	Add_StopPressed(handler EventHandler[interface{}]) EventRegistrationToken
	Remove_StopPressed(cookie EventRegistrationToken)
	Add_PlayPauseTogglePressed(handler EventHandler[interface{}]) EventRegistrationToken
	Remove_PlayPauseTogglePressed(cookie EventRegistrationToken)
	Add_RecordPressed(handler EventHandler[interface{}]) EventRegistrationToken
	Remove_RecordPressed(cookie EventRegistrationToken)
	Add_NextTrackPressed(handler EventHandler[interface{}]) EventRegistrationToken
	Remove_NextTrackPressed(cookie EventRegistrationToken)
	Add_PreviousTrackPressed(handler EventHandler[interface{}]) EventRegistrationToken
	Remove_PreviousTrackPressed(cookie EventRegistrationToken)
	Add_FastForwardPressed(handler EventHandler[interface{}]) EventRegistrationToken
	Remove_FastForwardPressed(cookie EventRegistrationToken)
	Add_RewindPressed(handler EventHandler[interface{}]) EventRegistrationToken
	Remove_RewindPressed(cookie EventRegistrationToken)
	Add_ChannelUpPressed(handler EventHandler[interface{}]) EventRegistrationToken
	Remove_ChannelUpPressed(cookie EventRegistrationToken)
	Add_ChannelDownPressed(handler EventHandler[interface{}]) EventRegistrationToken
	Remove_ChannelDownPressed(cookie EventRegistrationToken)
	Get_SoundLevel() SoundLevel
	Put_TrackName(value string)
	Get_TrackName() string
	Put_ArtistName(value string)
	Get_ArtistName() string
	Put_IsPlaying(value bool)
	Get_IsPlaying() bool
	Put_AlbumArt(value *IUriRuntimeClass)
	Get_AlbumArt() *IUriRuntimeClass
}

type IMediaControlVtbl struct {
	win32.IInspectableVtbl
	Add_SoundLevelChanged         uintptr
	Remove_SoundLevelChanged      uintptr
	Add_PlayPressed               uintptr
	Remove_PlayPressed            uintptr
	Add_PausePressed              uintptr
	Remove_PausePressed           uintptr
	Add_StopPressed               uintptr
	Remove_StopPressed            uintptr
	Add_PlayPauseTogglePressed    uintptr
	Remove_PlayPauseTogglePressed uintptr
	Add_RecordPressed             uintptr
	Remove_RecordPressed          uintptr
	Add_NextTrackPressed          uintptr
	Remove_NextTrackPressed       uintptr
	Add_PreviousTrackPressed      uintptr
	Remove_PreviousTrackPressed   uintptr
	Add_FastForwardPressed        uintptr
	Remove_FastForwardPressed     uintptr
	Add_RewindPressed             uintptr
	Remove_RewindPressed          uintptr
	Add_ChannelUpPressed          uintptr
	Remove_ChannelUpPressed       uintptr
	Add_ChannelDownPressed        uintptr
	Remove_ChannelDownPressed     uintptr
	Get_SoundLevel                uintptr
	Put_TrackName                 uintptr
	Get_TrackName                 uintptr
	Put_ArtistName                uintptr
	Get_ArtistName                uintptr
	Put_IsPlaying                 uintptr
	Get_IsPlaying                 uintptr
	Put_AlbumArt                  uintptr
	Get_AlbumArt                  uintptr
}

type IMediaControl struct {
	win32.IInspectable
}

func (this *IMediaControl) Vtbl() *IMediaControlVtbl {
	return (*IMediaControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaControl) Add_SoundLevelChanged(handler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SoundLevelChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaControl) Remove_SoundLevelChanged(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SoundLevelChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IMediaControl) Add_PlayPressed(handler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PlayPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaControl) Remove_PlayPressed(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PlayPressed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IMediaControl) Add_PausePressed(handler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PausePressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaControl) Remove_PausePressed(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PausePressed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IMediaControl) Add_StopPressed(handler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StopPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaControl) Remove_StopPressed(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StopPressed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IMediaControl) Add_PlayPauseTogglePressed(handler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PlayPauseTogglePressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaControl) Remove_PlayPauseTogglePressed(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PlayPauseTogglePressed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IMediaControl) Add_RecordPressed(handler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_RecordPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaControl) Remove_RecordPressed(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_RecordPressed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IMediaControl) Add_NextTrackPressed(handler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_NextTrackPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaControl) Remove_NextTrackPressed(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_NextTrackPressed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IMediaControl) Add_PreviousTrackPressed(handler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PreviousTrackPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaControl) Remove_PreviousTrackPressed(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PreviousTrackPressed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IMediaControl) Add_FastForwardPressed(handler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_FastForwardPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaControl) Remove_FastForwardPressed(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_FastForwardPressed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IMediaControl) Add_RewindPressed(handler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_RewindPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaControl) Remove_RewindPressed(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_RewindPressed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IMediaControl) Add_ChannelUpPressed(handler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ChannelUpPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaControl) Remove_ChannelUpPressed(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ChannelUpPressed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IMediaControl) Add_ChannelDownPressed(handler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ChannelDownPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaControl) Remove_ChannelDownPressed(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ChannelDownPressed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IMediaControl) Get_SoundLevel() SoundLevel {
	var _result SoundLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SoundLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaControl) Put_TrackName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TrackName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMediaControl) Get_TrackName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TrackName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaControl) Put_ArtistName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ArtistName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMediaControl) Get_ArtistName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ArtistName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaControl) Put_IsPlaying(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsPlaying, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaControl) Get_IsPlaying() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPlaying, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaControl) Put_AlbumArt(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AlbumArt, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaControl) Get_AlbumArt() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlbumArt, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 07915118-45DF-442B-8A3F-F7826A6370AB
var IID_IMediaExtension = syscall.GUID{0x07915118, 0x45DF, 0x442B,
	[8]byte{0x8A, 0x3F, 0xF7, 0x82, 0x6A, 0x63, 0x70, 0xAB}}

type IMediaExtensionInterface interface {
	win32.IInspectableInterface
	SetProperties(configuration *IPropertySet)
}

type IMediaExtensionVtbl struct {
	win32.IInspectableVtbl
	SetProperties uintptr
}

type IMediaExtension struct {
	win32.IInspectable
}

func (this *IMediaExtension) Vtbl() *IMediaExtensionVtbl {
	return (*IMediaExtensionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaExtension) SetProperties(configuration *IPropertySet) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(configuration)))
	_ = _hr
}

// 4A25EAF5-242D-4DFB-97F4-69B7C42576FF
var IID_IMediaExtensionManager = syscall.GUID{0x4A25EAF5, 0x242D, 0x4DFB,
	[8]byte{0x97, 0xF4, 0x69, 0xB7, 0xC4, 0x25, 0x76, 0xFF}}

type IMediaExtensionManagerInterface interface {
	win32.IInspectableInterface
	RegisterSchemeHandler(activatableClassId string, scheme string)
	RegisterSchemeHandlerWithSettings(activatableClassId string, scheme string, configuration *IPropertySet)
	RegisterByteStreamHandler(activatableClassId string, fileExtension string, mimeType string)
	RegisterByteStreamHandlerWithSettings(activatableClassId string, fileExtension string, mimeType string, configuration *IPropertySet)
	RegisterAudioDecoder(activatableClassId string, inputSubtype syscall.GUID, outputSubtype syscall.GUID)
	RegisterAudioDecoderWithSettings(activatableClassId string, inputSubtype syscall.GUID, outputSubtype syscall.GUID, configuration *IPropertySet)
	RegisterAudioEncoder(activatableClassId string, inputSubtype syscall.GUID, outputSubtype syscall.GUID)
	RegisterAudioEncoderWithSettings(activatableClassId string, inputSubtype syscall.GUID, outputSubtype syscall.GUID, configuration *IPropertySet)
	RegisterVideoDecoder(activatableClassId string, inputSubtype syscall.GUID, outputSubtype syscall.GUID)
	RegisterVideoDecoderWithSettings(activatableClassId string, inputSubtype syscall.GUID, outputSubtype syscall.GUID, configuration *IPropertySet)
	RegisterVideoEncoder(activatableClassId string, inputSubtype syscall.GUID, outputSubtype syscall.GUID)
	RegisterVideoEncoderWithSettings(activatableClassId string, inputSubtype syscall.GUID, outputSubtype syscall.GUID, configuration *IPropertySet)
}

type IMediaExtensionManagerVtbl struct {
	win32.IInspectableVtbl
	RegisterSchemeHandler                 uintptr
	RegisterSchemeHandlerWithSettings     uintptr
	RegisterByteStreamHandler             uintptr
	RegisterByteStreamHandlerWithSettings uintptr
	RegisterAudioDecoder                  uintptr
	RegisterAudioDecoderWithSettings      uintptr
	RegisterAudioEncoder                  uintptr
	RegisterAudioEncoderWithSettings      uintptr
	RegisterVideoDecoder                  uintptr
	RegisterVideoDecoderWithSettings      uintptr
	RegisterVideoEncoder                  uintptr
	RegisterVideoEncoderWithSettings      uintptr
}

type IMediaExtensionManager struct {
	win32.IInspectable
}

func (this *IMediaExtensionManager) Vtbl() *IMediaExtensionManagerVtbl {
	return (*IMediaExtensionManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaExtensionManager) RegisterSchemeHandler(activatableClassId string, scheme string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterSchemeHandler, uintptr(unsafe.Pointer(this)), NewHStr(activatableClassId).Ptr, NewHStr(scheme).Ptr)
	_ = _hr
}

func (this *IMediaExtensionManager) RegisterSchemeHandlerWithSettings(activatableClassId string, scheme string, configuration *IPropertySet) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterSchemeHandlerWithSettings, uintptr(unsafe.Pointer(this)), NewHStr(activatableClassId).Ptr, NewHStr(scheme).Ptr, uintptr(unsafe.Pointer(configuration)))
	_ = _hr
}

func (this *IMediaExtensionManager) RegisterByteStreamHandler(activatableClassId string, fileExtension string, mimeType string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterByteStreamHandler, uintptr(unsafe.Pointer(this)), NewHStr(activatableClassId).Ptr, NewHStr(fileExtension).Ptr, NewHStr(mimeType).Ptr)
	_ = _hr
}

func (this *IMediaExtensionManager) RegisterByteStreamHandlerWithSettings(activatableClassId string, fileExtension string, mimeType string, configuration *IPropertySet) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterByteStreamHandlerWithSettings, uintptr(unsafe.Pointer(this)), NewHStr(activatableClassId).Ptr, NewHStr(fileExtension).Ptr, NewHStr(mimeType).Ptr, uintptr(unsafe.Pointer(configuration)))
	_ = _hr
}

func (this *IMediaExtensionManager) RegisterAudioDecoder(activatableClassId string, inputSubtype syscall.GUID, outputSubtype syscall.GUID) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterAudioDecoder, uintptr(unsafe.Pointer(this)), NewHStr(activatableClassId).Ptr, uintptr(unsafe.Pointer(&inputSubtype)), uintptr(unsafe.Pointer(&outputSubtype)))
	_ = _hr
}

func (this *IMediaExtensionManager) RegisterAudioDecoderWithSettings(activatableClassId string, inputSubtype syscall.GUID, outputSubtype syscall.GUID, configuration *IPropertySet) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterAudioDecoderWithSettings, uintptr(unsafe.Pointer(this)), NewHStr(activatableClassId).Ptr, uintptr(unsafe.Pointer(&inputSubtype)), uintptr(unsafe.Pointer(&outputSubtype)), uintptr(unsafe.Pointer(configuration)))
	_ = _hr
}

func (this *IMediaExtensionManager) RegisterAudioEncoder(activatableClassId string, inputSubtype syscall.GUID, outputSubtype syscall.GUID) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterAudioEncoder, uintptr(unsafe.Pointer(this)), NewHStr(activatableClassId).Ptr, uintptr(unsafe.Pointer(&inputSubtype)), uintptr(unsafe.Pointer(&outputSubtype)))
	_ = _hr
}

func (this *IMediaExtensionManager) RegisterAudioEncoderWithSettings(activatableClassId string, inputSubtype syscall.GUID, outputSubtype syscall.GUID, configuration *IPropertySet) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterAudioEncoderWithSettings, uintptr(unsafe.Pointer(this)), NewHStr(activatableClassId).Ptr, uintptr(unsafe.Pointer(&inputSubtype)), uintptr(unsafe.Pointer(&outputSubtype)), uintptr(unsafe.Pointer(configuration)))
	_ = _hr
}

func (this *IMediaExtensionManager) RegisterVideoDecoder(activatableClassId string, inputSubtype syscall.GUID, outputSubtype syscall.GUID) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterVideoDecoder, uintptr(unsafe.Pointer(this)), NewHStr(activatableClassId).Ptr, uintptr(unsafe.Pointer(&inputSubtype)), uintptr(unsafe.Pointer(&outputSubtype)))
	_ = _hr
}

func (this *IMediaExtensionManager) RegisterVideoDecoderWithSettings(activatableClassId string, inputSubtype syscall.GUID, outputSubtype syscall.GUID, configuration *IPropertySet) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterVideoDecoderWithSettings, uintptr(unsafe.Pointer(this)), NewHStr(activatableClassId).Ptr, uintptr(unsafe.Pointer(&inputSubtype)), uintptr(unsafe.Pointer(&outputSubtype)), uintptr(unsafe.Pointer(configuration)))
	_ = _hr
}

func (this *IMediaExtensionManager) RegisterVideoEncoder(activatableClassId string, inputSubtype syscall.GUID, outputSubtype syscall.GUID) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterVideoEncoder, uintptr(unsafe.Pointer(this)), NewHStr(activatableClassId).Ptr, uintptr(unsafe.Pointer(&inputSubtype)), uintptr(unsafe.Pointer(&outputSubtype)))
	_ = _hr
}

func (this *IMediaExtensionManager) RegisterVideoEncoderWithSettings(activatableClassId string, inputSubtype syscall.GUID, outputSubtype syscall.GUID, configuration *IPropertySet) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterVideoEncoderWithSettings, uintptr(unsafe.Pointer(this)), NewHStr(activatableClassId).Ptr, uintptr(unsafe.Pointer(&inputSubtype)), uintptr(unsafe.Pointer(&outputSubtype)), uintptr(unsafe.Pointer(configuration)))
	_ = _hr
}

// 5BCEBF47-4043-4FED-ACAF-54EC29DFB1F7
var IID_IMediaExtensionManager2 = syscall.GUID{0x5BCEBF47, 0x4043, 0x4FED,
	[8]byte{0xAC, 0xAF, 0x54, 0xEC, 0x29, 0xDF, 0xB1, 0xF7}}

type IMediaExtensionManager2Interface interface {
	win32.IInspectableInterface
	RegisterMediaExtensionForAppService(extension *IMediaExtension, connection *IAppServiceConnection)
}

type IMediaExtensionManager2Vtbl struct {
	win32.IInspectableVtbl
	RegisterMediaExtensionForAppService uintptr
}

type IMediaExtensionManager2 struct {
	win32.IInspectable
}

func (this *IMediaExtensionManager2) Vtbl() *IMediaExtensionManager2Vtbl {
	return (*IMediaExtensionManager2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaExtensionManager2) RegisterMediaExtensionForAppService(extension *IMediaExtension, connection *IAppServiceConnection) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterMediaExtensionForAppService, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(extension)), uintptr(unsafe.Pointer(connection)))
	_ = _hr
}

// BFB52F8C-5943-47D8-8E10-05308AA5FBD0
var IID_IMediaFrame = syscall.GUID{0xBFB52F8C, 0x5943, 0x47D8,
	[8]byte{0x8E, 0x10, 0x05, 0x30, 0x8A, 0xA5, 0xFB, 0xD0}}

type IMediaFrameInterface interface {
	win32.IInspectableInterface
	Get_Type() string
	Get_IsReadOnly() bool
	Put_RelativeTime(value *IReference[TimeSpan])
	Get_RelativeTime() *IReference[TimeSpan]
	Put_SystemRelativeTime(value *IReference[TimeSpan])
	Get_SystemRelativeTime() *IReference[TimeSpan]
	Put_Duration(value *IReference[TimeSpan])
	Get_Duration() *IReference[TimeSpan]
	Put_IsDiscontinuous(value bool)
	Get_IsDiscontinuous() bool
	Get_ExtendedProperties() *IPropertySet
}

type IMediaFrameVtbl struct {
	win32.IInspectableVtbl
	Get_Type               uintptr
	Get_IsReadOnly         uintptr
	Put_RelativeTime       uintptr
	Get_RelativeTime       uintptr
	Put_SystemRelativeTime uintptr
	Get_SystemRelativeTime uintptr
	Put_Duration           uintptr
	Get_Duration           uintptr
	Put_IsDiscontinuous    uintptr
	Get_IsDiscontinuous    uintptr
	Get_ExtendedProperties uintptr
}

type IMediaFrame struct {
	win32.IInspectable
}

func (this *IMediaFrame) Vtbl() *IMediaFrameVtbl {
	return (*IMediaFrameVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaFrame) Get_Type() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Type, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaFrame) Get_IsReadOnly() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsReadOnly, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaFrame) Put_RelativeTime(value *IReference[TimeSpan]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RelativeTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaFrame) Get_RelativeTime() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RelativeTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaFrame) Put_SystemRelativeTime(value *IReference[TimeSpan]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SystemRelativeTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaFrame) Get_SystemRelativeTime() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SystemRelativeTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaFrame) Put_Duration(value *IReference[TimeSpan]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Duration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaFrame) Get_Duration() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Duration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaFrame) Put_IsDiscontinuous(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsDiscontinuous, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaFrame) Get_IsDiscontinuous() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDiscontinuous, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaFrame) Get_ExtendedProperties() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1803DEF8-DCA5-4B6F-9C20-E3D3C0643625
var IID_IMediaMarker = syscall.GUID{0x1803DEF8, 0xDCA5, 0x4B6F,
	[8]byte{0x9C, 0x20, 0xE3, 0xD3, 0xC0, 0x64, 0x36, 0x25}}

type IMediaMarkerInterface interface {
	win32.IInspectableInterface
	Get_Time() TimeSpan
	Get_MediaMarkerType() string
	Get_Text() string
}

type IMediaMarkerVtbl struct {
	win32.IInspectableVtbl
	Get_Time            uintptr
	Get_MediaMarkerType uintptr
	Get_Text            uintptr
}

type IMediaMarker struct {
	win32.IInspectable
}

func (this *IMediaMarker) Vtbl() *IMediaMarkerVtbl {
	return (*IMediaMarkerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaMarker) Get_Time() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Time, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaMarker) Get_MediaMarkerType() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaMarkerType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaMarker) Get_Text() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Text, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// BB198040-482F-4743-8832-45853821ECE0
var IID_IMediaMarkerTypesStatics = syscall.GUID{0xBB198040, 0x482F, 0x4743,
	[8]byte{0x88, 0x32, 0x45, 0x85, 0x38, 0x21, 0xEC, 0xE0}}

type IMediaMarkerTypesStaticsInterface interface {
	win32.IInspectableInterface
	Get_Bookmark() string
}

type IMediaMarkerTypesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Bookmark uintptr
}

type IMediaMarkerTypesStatics struct {
	win32.IInspectable
}

func (this *IMediaMarkerTypesStatics) Vtbl() *IMediaMarkerTypesStaticsVtbl {
	return (*IMediaMarkerTypesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaMarkerTypesStatics) Get_Bookmark() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bookmark, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// AFEAB189-F8DD-466E-AA10-920B52353FDF
var IID_IMediaMarkers = syscall.GUID{0xAFEAB189, 0xF8DD, 0x466E,
	[8]byte{0xAA, 0x10, 0x92, 0x0B, 0x52, 0x35, 0x3F, 0xDF}}

type IMediaMarkersInterface interface {
	win32.IInspectableInterface
	Get_Markers() *IVectorView[*IMediaMarker]
}

type IMediaMarkersVtbl struct {
	win32.IInspectableVtbl
	Get_Markers uintptr
}

type IMediaMarkers struct {
	win32.IInspectable
}

func (this *IMediaMarkers) Vtbl() *IMediaMarkersVtbl {
	return (*IMediaMarkersVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaMarkers) Get_Markers() *IVectorView[*IMediaMarker] {
	var _result *IVectorView[*IMediaMarker]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Markers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EB8564AC-A351-4F4E-B4F0-9BF2408993DB
var IID_IMediaProcessingTriggerDetails = syscall.GUID{0xEB8564AC, 0xA351, 0x4F4E,
	[8]byte{0xB4, 0xF0, 0x9B, 0xF2, 0x40, 0x89, 0x93, 0xDB}}

type IMediaProcessingTriggerDetailsInterface interface {
	win32.IInspectableInterface
	Get_Arguments() *IPropertySet
}

type IMediaProcessingTriggerDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_Arguments uintptr
}

type IMediaProcessingTriggerDetails struct {
	win32.IInspectable
}

func (this *IMediaProcessingTriggerDetails) Vtbl() *IMediaProcessingTriggerDetailsVtbl {
	return (*IMediaProcessingTriggerDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaProcessingTriggerDetails) Get_Arguments() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Arguments, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8ED361F3-0B78-4360-BF71-0C841999EA1B
var IID_IMediaTimelineController = syscall.GUID{0x8ED361F3, 0x0B78, 0x4360,
	[8]byte{0xBF, 0x71, 0x0C, 0x84, 0x19, 0x99, 0xEA, 0x1B}}

type IMediaTimelineControllerInterface interface {
	win32.IInspectableInterface
	Start()
	Resume()
	Pause()
	Get_Position() TimeSpan
	Put_Position(value TimeSpan)
	Get_ClockRate() float64
	Put_ClockRate(value float64)
	Get_State() MediaTimelineControllerState
	Add_PositionChanged(positionChangedEventHandler TypedEventHandler[*IMediaTimelineController, interface{}]) EventRegistrationToken
	Remove_PositionChanged(eventCookie EventRegistrationToken)
	Add_StateChanged(stateChangedEventHandler TypedEventHandler[*IMediaTimelineController, interface{}]) EventRegistrationToken
	Remove_StateChanged(eventCookie EventRegistrationToken)
}

type IMediaTimelineControllerVtbl struct {
	win32.IInspectableVtbl
	Start                  uintptr
	Resume                 uintptr
	Pause                  uintptr
	Get_Position           uintptr
	Put_Position           uintptr
	Get_ClockRate          uintptr
	Put_ClockRate          uintptr
	Get_State              uintptr
	Add_PositionChanged    uintptr
	Remove_PositionChanged uintptr
	Add_StateChanged       uintptr
	Remove_StateChanged    uintptr
}

type IMediaTimelineController struct {
	win32.IInspectable
}

func (this *IMediaTimelineController) Vtbl() *IMediaTimelineControllerVtbl {
	return (*IMediaTimelineControllerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaTimelineController) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IMediaTimelineController) Resume() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Resume, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IMediaTimelineController) Pause() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Pause, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IMediaTimelineController) Get_Position() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaTimelineController) Put_Position(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Position, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IMediaTimelineController) Get_ClockRate() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ClockRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaTimelineController) Put_ClockRate(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ClockRate, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMediaTimelineController) Get_State() MediaTimelineControllerState {
	var _result MediaTimelineControllerState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaTimelineController) Add_PositionChanged(positionChangedEventHandler TypedEventHandler[*IMediaTimelineController, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PositionChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(positionChangedEventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaTimelineController) Remove_PositionChanged(eventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PositionChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&eventCookie)))
	_ = _hr
}

func (this *IMediaTimelineController) Add_StateChanged(stateChangedEventHandler TypedEventHandler[*IMediaTimelineController, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(stateChangedEventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaTimelineController) Remove_StateChanged(eventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&eventCookie)))
	_ = _hr
}

// EF74EA38-9E72-4DF9-8355-6E90C81BBADD
var IID_IMediaTimelineController2 = syscall.GUID{0xEF74EA38, 0x9E72, 0x4DF9,
	[8]byte{0x83, 0x55, 0x6E, 0x90, 0xC8, 0x1B, 0xBA, 0xDD}}

type IMediaTimelineController2Interface interface {
	win32.IInspectableInterface
	Get_Duration() *IReference[TimeSpan]
	Put_Duration(value *IReference[TimeSpan])
	Get_IsLoopingEnabled() bool
	Put_IsLoopingEnabled(value bool)
	Add_Failed(eventHandler TypedEventHandler[*IMediaTimelineController, *IMediaTimelineControllerFailedEventArgs]) EventRegistrationToken
	Remove_Failed(token EventRegistrationToken)
	Add_Ended(eventHandler TypedEventHandler[*IMediaTimelineController, interface{}]) EventRegistrationToken
	Remove_Ended(token EventRegistrationToken)
}

type IMediaTimelineController2Vtbl struct {
	win32.IInspectableVtbl
	Get_Duration         uintptr
	Put_Duration         uintptr
	Get_IsLoopingEnabled uintptr
	Put_IsLoopingEnabled uintptr
	Add_Failed           uintptr
	Remove_Failed        uintptr
	Add_Ended            uintptr
	Remove_Ended         uintptr
}

type IMediaTimelineController2 struct {
	win32.IInspectable
}

func (this *IMediaTimelineController2) Vtbl() *IMediaTimelineController2Vtbl {
	return (*IMediaTimelineController2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaTimelineController2) Get_Duration() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Duration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaTimelineController2) Put_Duration(value *IReference[TimeSpan]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Duration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaTimelineController2) Get_IsLoopingEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsLoopingEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaTimelineController2) Put_IsLoopingEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsLoopingEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaTimelineController2) Add_Failed(eventHandler TypedEventHandler[*IMediaTimelineController, *IMediaTimelineControllerFailedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Failed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(eventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaTimelineController2) Remove_Failed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Failed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaTimelineController2) Add_Ended(eventHandler TypedEventHandler[*IMediaTimelineController, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Ended, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(eventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaTimelineController2) Remove_Ended(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Ended, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 8821F81D-3E77-43FB-BE26-4FC87A044834
var IID_IMediaTimelineControllerFailedEventArgs = syscall.GUID{0x8821F81D, 0x3E77, 0x43FB,
	[8]byte{0xBE, 0x26, 0x4F, 0xC8, 0x7A, 0x04, 0x48, 0x34}}

type IMediaTimelineControllerFailedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ExtendedError() HResult
}

type IMediaTimelineControllerFailedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ExtendedError uintptr
}

type IMediaTimelineControllerFailedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaTimelineControllerFailedEventArgs) Vtbl() *IMediaTimelineControllerFailedEventArgsVtbl {
	return (*IMediaTimelineControllerFailedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaTimelineControllerFailedEventArgs) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 6BBF0C59-D0A0-4D26-92A0-F978E1D18E7B
var IID_IMusicDisplayProperties = syscall.GUID{0x6BBF0C59, 0xD0A0, 0x4D26,
	[8]byte{0x92, 0xA0, 0xF9, 0x78, 0xE1, 0xD1, 0x8E, 0x7B}}

type IMusicDisplayPropertiesInterface interface {
	win32.IInspectableInterface
	Get_Title() string
	Put_Title(value string)
	Get_AlbumArtist() string
	Put_AlbumArtist(value string)
	Get_Artist() string
	Put_Artist(value string)
}

type IMusicDisplayPropertiesVtbl struct {
	win32.IInspectableVtbl
	Get_Title       uintptr
	Put_Title       uintptr
	Get_AlbumArtist uintptr
	Put_AlbumArtist uintptr
	Get_Artist      uintptr
	Put_Artist      uintptr
}

type IMusicDisplayProperties struct {
	win32.IInspectable
}

func (this *IMusicDisplayProperties) Vtbl() *IMusicDisplayPropertiesVtbl {
	return (*IMusicDisplayPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMusicDisplayProperties) Get_Title() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Title, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMusicDisplayProperties) Put_Title(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Title, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMusicDisplayProperties) Get_AlbumArtist() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlbumArtist, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMusicDisplayProperties) Put_AlbumArtist(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AlbumArtist, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMusicDisplayProperties) Get_Artist() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Artist, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMusicDisplayProperties) Put_Artist(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Artist, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// 00368462-97D3-44B9-B00F-008AFCEFAF18
var IID_IMusicDisplayProperties2 = syscall.GUID{0x00368462, 0x97D3, 0x44B9,
	[8]byte{0xB0, 0x0F, 0x00, 0x8A, 0xFC, 0xEF, 0xAF, 0x18}}

type IMusicDisplayProperties2Interface interface {
	win32.IInspectableInterface
	Get_AlbumTitle() string
	Put_AlbumTitle(value string)
	Get_TrackNumber() uint32
	Put_TrackNumber(value uint32)
	Get_Genres() *IVector[string]
}

type IMusicDisplayProperties2Vtbl struct {
	win32.IInspectableVtbl
	Get_AlbumTitle  uintptr
	Put_AlbumTitle  uintptr
	Get_TrackNumber uintptr
	Put_TrackNumber uintptr
	Get_Genres      uintptr
}

type IMusicDisplayProperties2 struct {
	win32.IInspectable
}

func (this *IMusicDisplayProperties2) Vtbl() *IMusicDisplayProperties2Vtbl {
	return (*IMusicDisplayProperties2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMusicDisplayProperties2) Get_AlbumTitle() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlbumTitle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMusicDisplayProperties2) Put_AlbumTitle(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AlbumTitle, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMusicDisplayProperties2) Get_TrackNumber() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TrackNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMusicDisplayProperties2) Put_TrackNumber(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TrackNumber, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMusicDisplayProperties2) Get_Genres() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Genres, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4DB51AC1-0681-4E8C-9401-B8159D9EEFC7
var IID_IMusicDisplayProperties3 = syscall.GUID{0x4DB51AC1, 0x0681, 0x4E8C,
	[8]byte{0x94, 0x01, 0xB8, 0x15, 0x9D, 0x9E, 0xEF, 0xC7}}

type IMusicDisplayProperties3Interface interface {
	win32.IInspectableInterface
	Get_AlbumTrackCount() uint32
	Put_AlbumTrackCount(value uint32)
}

type IMusicDisplayProperties3Vtbl struct {
	win32.IInspectableVtbl
	Get_AlbumTrackCount uintptr
	Put_AlbumTrackCount uintptr
}

type IMusicDisplayProperties3 struct {
	win32.IInspectable
}

func (this *IMusicDisplayProperties3) Vtbl() *IMusicDisplayProperties3Vtbl {
	return (*IMusicDisplayProperties3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMusicDisplayProperties3) Get_AlbumTrackCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlbumTrackCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMusicDisplayProperties3) Put_AlbumTrackCount(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AlbumTrackCount, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// B4493F88-EB28-4961-9C14-335E44F3E125
var IID_IPlaybackPositionChangeRequestedEventArgs = syscall.GUID{0xB4493F88, 0xEB28, 0x4961,
	[8]byte{0x9C, 0x14, 0x33, 0x5E, 0x44, 0xF3, 0xE1, 0x25}}

type IPlaybackPositionChangeRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_RequestedPlaybackPosition() TimeSpan
}

type IPlaybackPositionChangeRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_RequestedPlaybackPosition uintptr
}

type IPlaybackPositionChangeRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IPlaybackPositionChangeRequestedEventArgs) Vtbl() *IPlaybackPositionChangeRequestedEventArgsVtbl {
	return (*IPlaybackPositionChangeRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPlaybackPositionChangeRequestedEventArgs) Get_RequestedPlaybackPosition() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestedPlaybackPosition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 2CE2C41F-3CD6-4F77-9BA7-EB27C26A2140
var IID_IPlaybackRateChangeRequestedEventArgs = syscall.GUID{0x2CE2C41F, 0x3CD6, 0x4F77,
	[8]byte{0x9B, 0xA7, 0xEB, 0x27, 0xC2, 0x6A, 0x21, 0x40}}

type IPlaybackRateChangeRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_RequestedPlaybackRate() float64
}

type IPlaybackRateChangeRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_RequestedPlaybackRate uintptr
}

type IPlaybackRateChangeRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IPlaybackRateChangeRequestedEventArgs) Vtbl() *IPlaybackRateChangeRequestedEventArgsVtbl {
	return (*IPlaybackRateChangeRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPlaybackRateChangeRequestedEventArgs) Get_RequestedPlaybackRate() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestedPlaybackRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 49B593FE-4FD0-4666-A314-C0E01940D302
var IID_IShuffleEnabledChangeRequestedEventArgs = syscall.GUID{0x49B593FE, 0x4FD0, 0x4666,
	[8]byte{0xA3, 0x14, 0xC0, 0xE0, 0x19, 0x40, 0xD3, 0x02}}

type IShuffleEnabledChangeRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_RequestedShuffleEnabled() bool
}

type IShuffleEnabledChangeRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_RequestedShuffleEnabled uintptr
}

type IShuffleEnabledChangeRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IShuffleEnabledChangeRequestedEventArgs) Vtbl() *IShuffleEnabledChangeRequestedEventArgsVtbl {
	return (*IShuffleEnabledChangeRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IShuffleEnabledChangeRequestedEventArgs) Get_RequestedShuffleEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestedShuffleEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 99FA3FF4-1742-42A6-902E-087D41F965EC
var IID_ISystemMediaTransportControls = syscall.GUID{0x99FA3FF4, 0x1742, 0x42A6,
	[8]byte{0x90, 0x2E, 0x08, 0x7D, 0x41, 0xF9, 0x65, 0xEC}}

type ISystemMediaTransportControlsInterface interface {
	win32.IInspectableInterface
	Get_PlaybackStatus() MediaPlaybackStatus
	Put_PlaybackStatus(value MediaPlaybackStatus)
	Get_DisplayUpdater() *ISystemMediaTransportControlsDisplayUpdater
	Get_SoundLevel() SoundLevel
	Get_IsEnabled() bool
	Put_IsEnabled(value bool)
	Get_IsPlayEnabled() bool
	Put_IsPlayEnabled(value bool)
	Get_IsStopEnabled() bool
	Put_IsStopEnabled(value bool)
	Get_IsPauseEnabled() bool
	Put_IsPauseEnabled(value bool)
	Get_IsRecordEnabled() bool
	Put_IsRecordEnabled(value bool)
	Get_IsFastForwardEnabled() bool
	Put_IsFastForwardEnabled(value bool)
	Get_IsRewindEnabled() bool
	Put_IsRewindEnabled(value bool)
	Get_IsPreviousEnabled() bool
	Put_IsPreviousEnabled(value bool)
	Get_IsNextEnabled() bool
	Put_IsNextEnabled(value bool)
	Get_IsChannelUpEnabled() bool
	Put_IsChannelUpEnabled(value bool)
	Get_IsChannelDownEnabled() bool
	Put_IsChannelDownEnabled(value bool)
	Add_ButtonPressed(handler TypedEventHandler[*ISystemMediaTransportControls, *ISystemMediaTransportControlsButtonPressedEventArgs]) EventRegistrationToken
	Remove_ButtonPressed(token EventRegistrationToken)
	Add_PropertyChanged(handler TypedEventHandler[*ISystemMediaTransportControls, *ISystemMediaTransportControlsPropertyChangedEventArgs]) EventRegistrationToken
	Remove_PropertyChanged(token EventRegistrationToken)
}

type ISystemMediaTransportControlsVtbl struct {
	win32.IInspectableVtbl
	Get_PlaybackStatus       uintptr
	Put_PlaybackStatus       uintptr
	Get_DisplayUpdater       uintptr
	Get_SoundLevel           uintptr
	Get_IsEnabled            uintptr
	Put_IsEnabled            uintptr
	Get_IsPlayEnabled        uintptr
	Put_IsPlayEnabled        uintptr
	Get_IsStopEnabled        uintptr
	Put_IsStopEnabled        uintptr
	Get_IsPauseEnabled       uintptr
	Put_IsPauseEnabled       uintptr
	Get_IsRecordEnabled      uintptr
	Put_IsRecordEnabled      uintptr
	Get_IsFastForwardEnabled uintptr
	Put_IsFastForwardEnabled uintptr
	Get_IsRewindEnabled      uintptr
	Put_IsRewindEnabled      uintptr
	Get_IsPreviousEnabled    uintptr
	Put_IsPreviousEnabled    uintptr
	Get_IsNextEnabled        uintptr
	Put_IsNextEnabled        uintptr
	Get_IsChannelUpEnabled   uintptr
	Put_IsChannelUpEnabled   uintptr
	Get_IsChannelDownEnabled uintptr
	Put_IsChannelDownEnabled uintptr
	Add_ButtonPressed        uintptr
	Remove_ButtonPressed     uintptr
	Add_PropertyChanged      uintptr
	Remove_PropertyChanged   uintptr
}

type ISystemMediaTransportControls struct {
	win32.IInspectable
}

func (this *ISystemMediaTransportControls) Vtbl() *ISystemMediaTransportControlsVtbl {
	return (*ISystemMediaTransportControlsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemMediaTransportControls) Get_PlaybackStatus() MediaPlaybackStatus {
	var _result MediaPlaybackStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlaybackStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMediaTransportControls) Put_PlaybackStatus(value MediaPlaybackStatus) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PlaybackStatus, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISystemMediaTransportControls) Get_DisplayUpdater() *ISystemMediaTransportControlsDisplayUpdater {
	var _result *ISystemMediaTransportControlsDisplayUpdater
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayUpdater, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISystemMediaTransportControls) Get_SoundLevel() SoundLevel {
	var _result SoundLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SoundLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMediaTransportControls) Get_IsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMediaTransportControls) Put_IsEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ISystemMediaTransportControls) Get_IsPlayEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPlayEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMediaTransportControls) Put_IsPlayEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsPlayEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ISystemMediaTransportControls) Get_IsStopEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsStopEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMediaTransportControls) Put_IsStopEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsStopEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ISystemMediaTransportControls) Get_IsPauseEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPauseEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMediaTransportControls) Put_IsPauseEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsPauseEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ISystemMediaTransportControls) Get_IsRecordEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsRecordEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMediaTransportControls) Put_IsRecordEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsRecordEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ISystemMediaTransportControls) Get_IsFastForwardEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsFastForwardEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMediaTransportControls) Put_IsFastForwardEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsFastForwardEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ISystemMediaTransportControls) Get_IsRewindEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsRewindEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMediaTransportControls) Put_IsRewindEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsRewindEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ISystemMediaTransportControls) Get_IsPreviousEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPreviousEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMediaTransportControls) Put_IsPreviousEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsPreviousEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ISystemMediaTransportControls) Get_IsNextEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsNextEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMediaTransportControls) Put_IsNextEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsNextEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ISystemMediaTransportControls) Get_IsChannelUpEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsChannelUpEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMediaTransportControls) Put_IsChannelUpEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsChannelUpEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ISystemMediaTransportControls) Get_IsChannelDownEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsChannelDownEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMediaTransportControls) Put_IsChannelDownEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsChannelDownEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ISystemMediaTransportControls) Add_ButtonPressed(handler TypedEventHandler[*ISystemMediaTransportControls, *ISystemMediaTransportControlsButtonPressedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ButtonPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMediaTransportControls) Remove_ButtonPressed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ButtonPressed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISystemMediaTransportControls) Add_PropertyChanged(handler TypedEventHandler[*ISystemMediaTransportControls, *ISystemMediaTransportControlsPropertyChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PropertyChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMediaTransportControls) Remove_PropertyChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PropertyChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// EA98D2F6-7F3C-4AF2-A586-72889808EFB1
var IID_ISystemMediaTransportControls2 = syscall.GUID{0xEA98D2F6, 0x7F3C, 0x4AF2,
	[8]byte{0xA5, 0x86, 0x72, 0x88, 0x98, 0x08, 0xEF, 0xB1}}

type ISystemMediaTransportControls2Interface interface {
	win32.IInspectableInterface
	Get_AutoRepeatMode() MediaPlaybackAutoRepeatMode
	Put_AutoRepeatMode(value MediaPlaybackAutoRepeatMode)
	Get_ShuffleEnabled() bool
	Put_ShuffleEnabled(value bool)
	Get_PlaybackRate() float64
	Put_PlaybackRate(value float64)
	UpdateTimelineProperties(timelineProperties *ISystemMediaTransportControlsTimelineProperties)
	Add_PlaybackPositionChangeRequested(handler TypedEventHandler[*ISystemMediaTransportControls, *IPlaybackPositionChangeRequestedEventArgs]) EventRegistrationToken
	Remove_PlaybackPositionChangeRequested(token EventRegistrationToken)
	Add_PlaybackRateChangeRequested(handler TypedEventHandler[*ISystemMediaTransportControls, *IPlaybackRateChangeRequestedEventArgs]) EventRegistrationToken
	Remove_PlaybackRateChangeRequested(token EventRegistrationToken)
	Add_ShuffleEnabledChangeRequested(handler TypedEventHandler[*ISystemMediaTransportControls, *IShuffleEnabledChangeRequestedEventArgs]) EventRegistrationToken
	Remove_ShuffleEnabledChangeRequested(token EventRegistrationToken)
	Add_AutoRepeatModeChangeRequested(handler TypedEventHandler[*ISystemMediaTransportControls, *IAutoRepeatModeChangeRequestedEventArgs]) EventRegistrationToken
	Remove_AutoRepeatModeChangeRequested(token EventRegistrationToken)
}

type ISystemMediaTransportControls2Vtbl struct {
	win32.IInspectableVtbl
	Get_AutoRepeatMode                     uintptr
	Put_AutoRepeatMode                     uintptr
	Get_ShuffleEnabled                     uintptr
	Put_ShuffleEnabled                     uintptr
	Get_PlaybackRate                       uintptr
	Put_PlaybackRate                       uintptr
	UpdateTimelineProperties               uintptr
	Add_PlaybackPositionChangeRequested    uintptr
	Remove_PlaybackPositionChangeRequested uintptr
	Add_PlaybackRateChangeRequested        uintptr
	Remove_PlaybackRateChangeRequested     uintptr
	Add_ShuffleEnabledChangeRequested      uintptr
	Remove_ShuffleEnabledChangeRequested   uintptr
	Add_AutoRepeatModeChangeRequested      uintptr
	Remove_AutoRepeatModeChangeRequested   uintptr
}

type ISystemMediaTransportControls2 struct {
	win32.IInspectable
}

func (this *ISystemMediaTransportControls2) Vtbl() *ISystemMediaTransportControls2Vtbl {
	return (*ISystemMediaTransportControls2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemMediaTransportControls2) Get_AutoRepeatMode() MediaPlaybackAutoRepeatMode {
	var _result MediaPlaybackAutoRepeatMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AutoRepeatMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMediaTransportControls2) Put_AutoRepeatMode(value MediaPlaybackAutoRepeatMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AutoRepeatMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISystemMediaTransportControls2) Get_ShuffleEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ShuffleEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMediaTransportControls2) Put_ShuffleEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ShuffleEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ISystemMediaTransportControls2) Get_PlaybackRate() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlaybackRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMediaTransportControls2) Put_PlaybackRate(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PlaybackRate, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISystemMediaTransportControls2) UpdateTimelineProperties(timelineProperties *ISystemMediaTransportControlsTimelineProperties) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UpdateTimelineProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(timelineProperties)))
	_ = _hr
}

func (this *ISystemMediaTransportControls2) Add_PlaybackPositionChangeRequested(handler TypedEventHandler[*ISystemMediaTransportControls, *IPlaybackPositionChangeRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PlaybackPositionChangeRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMediaTransportControls2) Remove_PlaybackPositionChangeRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PlaybackPositionChangeRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISystemMediaTransportControls2) Add_PlaybackRateChangeRequested(handler TypedEventHandler[*ISystemMediaTransportControls, *IPlaybackRateChangeRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PlaybackRateChangeRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMediaTransportControls2) Remove_PlaybackRateChangeRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PlaybackRateChangeRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISystemMediaTransportControls2) Add_ShuffleEnabledChangeRequested(handler TypedEventHandler[*ISystemMediaTransportControls, *IShuffleEnabledChangeRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ShuffleEnabledChangeRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMediaTransportControls2) Remove_ShuffleEnabledChangeRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ShuffleEnabledChangeRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISystemMediaTransportControls2) Add_AutoRepeatModeChangeRequested(handler TypedEventHandler[*ISystemMediaTransportControls, *IAutoRepeatModeChangeRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AutoRepeatModeChangeRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMediaTransportControls2) Remove_AutoRepeatModeChangeRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AutoRepeatModeChangeRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// B7F47116-A56F-4DC8-9E11-92031F4A87C2
var IID_ISystemMediaTransportControlsButtonPressedEventArgs = syscall.GUID{0xB7F47116, 0xA56F, 0x4DC8,
	[8]byte{0x9E, 0x11, 0x92, 0x03, 0x1F, 0x4A, 0x87, 0xC2}}

type ISystemMediaTransportControlsButtonPressedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Button() SystemMediaTransportControlsButton
}

type ISystemMediaTransportControlsButtonPressedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Button uintptr
}

type ISystemMediaTransportControlsButtonPressedEventArgs struct {
	win32.IInspectable
}

func (this *ISystemMediaTransportControlsButtonPressedEventArgs) Vtbl() *ISystemMediaTransportControlsButtonPressedEventArgsVtbl {
	return (*ISystemMediaTransportControlsButtonPressedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemMediaTransportControlsButtonPressedEventArgs) Get_Button() SystemMediaTransportControlsButton {
	var _result SystemMediaTransportControlsButton
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Button, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 8ABBC53E-FA55-4ECF-AD8E-C984E5DD1550
var IID_ISystemMediaTransportControlsDisplayUpdater = syscall.GUID{0x8ABBC53E, 0xFA55, 0x4ECF,
	[8]byte{0xAD, 0x8E, 0xC9, 0x84, 0xE5, 0xDD, 0x15, 0x50}}

type ISystemMediaTransportControlsDisplayUpdaterInterface interface {
	win32.IInspectableInterface
	Get_Type() MediaPlaybackType
	Put_Type(value MediaPlaybackType)
	Get_AppMediaId() string
	Put_AppMediaId(value string)
	Get_Thumbnail() *IRandomAccessStreamReference
	Put_Thumbnail(value *IRandomAccessStreamReference)
	Get_MusicProperties() *IMusicDisplayProperties
	Get_VideoProperties() *IVideoDisplayProperties
	Get_ImageProperties() *IImageDisplayProperties
	CopyFromFileAsync(type_ MediaPlaybackType, source *IStorageFile) *IAsyncOperation[bool]
	ClearAll()
	Update()
}

type ISystemMediaTransportControlsDisplayUpdaterVtbl struct {
	win32.IInspectableVtbl
	Get_Type            uintptr
	Put_Type            uintptr
	Get_AppMediaId      uintptr
	Put_AppMediaId      uintptr
	Get_Thumbnail       uintptr
	Put_Thumbnail       uintptr
	Get_MusicProperties uintptr
	Get_VideoProperties uintptr
	Get_ImageProperties uintptr
	CopyFromFileAsync   uintptr
	ClearAll            uintptr
	Update              uintptr
}

type ISystemMediaTransportControlsDisplayUpdater struct {
	win32.IInspectable
}

func (this *ISystemMediaTransportControlsDisplayUpdater) Vtbl() *ISystemMediaTransportControlsDisplayUpdaterVtbl {
	return (*ISystemMediaTransportControlsDisplayUpdaterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemMediaTransportControlsDisplayUpdater) Get_Type() MediaPlaybackType {
	var _result MediaPlaybackType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Type, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMediaTransportControlsDisplayUpdater) Put_Type(value MediaPlaybackType) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Type, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISystemMediaTransportControlsDisplayUpdater) Get_AppMediaId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppMediaId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemMediaTransportControlsDisplayUpdater) Put_AppMediaId(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AppMediaId, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISystemMediaTransportControlsDisplayUpdater) Get_Thumbnail() *IRandomAccessStreamReference {
	var _result *IRandomAccessStreamReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Thumbnail, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISystemMediaTransportControlsDisplayUpdater) Put_Thumbnail(value *IRandomAccessStreamReference) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Thumbnail, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ISystemMediaTransportControlsDisplayUpdater) Get_MusicProperties() *IMusicDisplayProperties {
	var _result *IMusicDisplayProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MusicProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISystemMediaTransportControlsDisplayUpdater) Get_VideoProperties() *IVideoDisplayProperties {
	var _result *IVideoDisplayProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISystemMediaTransportControlsDisplayUpdater) Get_ImageProperties() *IImageDisplayProperties {
	var _result *IImageDisplayProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ImageProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISystemMediaTransportControlsDisplayUpdater) CopyFromFileAsync(type_ MediaPlaybackType, source *IStorageFile) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CopyFromFileAsync, uintptr(unsafe.Pointer(this)), uintptr(type_), uintptr(unsafe.Pointer(source)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISystemMediaTransportControlsDisplayUpdater) ClearAll() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearAll, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *ISystemMediaTransportControlsDisplayUpdater) Update() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Update, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// D0CA0936-339B-4CB3-8EEB-737607F56E08
var IID_ISystemMediaTransportControlsPropertyChangedEventArgs = syscall.GUID{0xD0CA0936, 0x339B, 0x4CB3,
	[8]byte{0x8E, 0xEB, 0x73, 0x76, 0x07, 0xF5, 0x6E, 0x08}}

type ISystemMediaTransportControlsPropertyChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Property() SystemMediaTransportControlsProperty
}

type ISystemMediaTransportControlsPropertyChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Property uintptr
}

type ISystemMediaTransportControlsPropertyChangedEventArgs struct {
	win32.IInspectable
}

func (this *ISystemMediaTransportControlsPropertyChangedEventArgs) Vtbl() *ISystemMediaTransportControlsPropertyChangedEventArgsVtbl {
	return (*ISystemMediaTransportControlsPropertyChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemMediaTransportControlsPropertyChangedEventArgs) Get_Property() SystemMediaTransportControlsProperty {
	var _result SystemMediaTransportControlsProperty
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Property, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 43BA380A-ECA4-4832-91AB-D415FAE484C6
var IID_ISystemMediaTransportControlsStatics = syscall.GUID{0x43BA380A, 0xECA4, 0x4832,
	[8]byte{0x91, 0xAB, 0xD4, 0x15, 0xFA, 0xE4, 0x84, 0xC6}}

type ISystemMediaTransportControlsStaticsInterface interface {
	win32.IInspectableInterface
	GetForCurrentView() *ISystemMediaTransportControls
}

type ISystemMediaTransportControlsStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForCurrentView uintptr
}

type ISystemMediaTransportControlsStatics struct {
	win32.IInspectable
}

func (this *ISystemMediaTransportControlsStatics) Vtbl() *ISystemMediaTransportControlsStaticsVtbl {
	return (*ISystemMediaTransportControlsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemMediaTransportControlsStatics) GetForCurrentView() *ISystemMediaTransportControls {
	var _result *ISystemMediaTransportControls
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5125316A-C3A2-475B-8507-93534DC88F15
var IID_ISystemMediaTransportControlsTimelineProperties = syscall.GUID{0x5125316A, 0xC3A2, 0x475B,
	[8]byte{0x85, 0x07, 0x93, 0x53, 0x4D, 0xC8, 0x8F, 0x15}}

type ISystemMediaTransportControlsTimelinePropertiesInterface interface {
	win32.IInspectableInterface
	Get_StartTime() TimeSpan
	Put_StartTime(value TimeSpan)
	Get_EndTime() TimeSpan
	Put_EndTime(value TimeSpan)
	Get_MinSeekTime() TimeSpan
	Put_MinSeekTime(value TimeSpan)
	Get_MaxSeekTime() TimeSpan
	Put_MaxSeekTime(value TimeSpan)
	Get_Position() TimeSpan
	Put_Position(value TimeSpan)
}

type ISystemMediaTransportControlsTimelinePropertiesVtbl struct {
	win32.IInspectableVtbl
	Get_StartTime   uintptr
	Put_StartTime   uintptr
	Get_EndTime     uintptr
	Put_EndTime     uintptr
	Get_MinSeekTime uintptr
	Put_MinSeekTime uintptr
	Get_MaxSeekTime uintptr
	Put_MaxSeekTime uintptr
	Get_Position    uintptr
	Put_Position    uintptr
}

type ISystemMediaTransportControlsTimelineProperties struct {
	win32.IInspectable
}

func (this *ISystemMediaTransportControlsTimelineProperties) Vtbl() *ISystemMediaTransportControlsTimelinePropertiesVtbl {
	return (*ISystemMediaTransportControlsTimelinePropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemMediaTransportControlsTimelineProperties) Get_StartTime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StartTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMediaTransportControlsTimelineProperties) Put_StartTime(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StartTime, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ISystemMediaTransportControlsTimelineProperties) Get_EndTime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EndTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMediaTransportControlsTimelineProperties) Put_EndTime(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_EndTime, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ISystemMediaTransportControlsTimelineProperties) Get_MinSeekTime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinSeekTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMediaTransportControlsTimelineProperties) Put_MinSeekTime(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MinSeekTime, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ISystemMediaTransportControlsTimelineProperties) Get_MaxSeekTime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxSeekTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMediaTransportControlsTimelineProperties) Put_MaxSeekTime(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxSeekTime, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ISystemMediaTransportControlsTimelineProperties) Get_Position() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemMediaTransportControlsTimelineProperties) Put_Position(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Position, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

// 5609FDB1-5D2D-4872-8170-45DEE5BC2F5C
var IID_IVideoDisplayProperties = syscall.GUID{0x5609FDB1, 0x5D2D, 0x4872,
	[8]byte{0x81, 0x70, 0x45, 0xDE, 0xE5, 0xBC, 0x2F, 0x5C}}

type IVideoDisplayPropertiesInterface interface {
	win32.IInspectableInterface
	Get_Title() string
	Put_Title(value string)
	Get_Subtitle() string
	Put_Subtitle(value string)
}

type IVideoDisplayPropertiesVtbl struct {
	win32.IInspectableVtbl
	Get_Title    uintptr
	Put_Title    uintptr
	Get_Subtitle uintptr
	Put_Subtitle uintptr
}

type IVideoDisplayProperties struct {
	win32.IInspectable
}

func (this *IVideoDisplayProperties) Vtbl() *IVideoDisplayPropertiesVtbl {
	return (*IVideoDisplayPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoDisplayProperties) Get_Title() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Title, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IVideoDisplayProperties) Put_Title(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Title, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IVideoDisplayProperties) Get_Subtitle() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Subtitle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IVideoDisplayProperties) Put_Subtitle(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Subtitle, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// B410E1CE-AB52-41AB-A486-CC10FAB152F9
var IID_IVideoDisplayProperties2 = syscall.GUID{0xB410E1CE, 0xAB52, 0x41AB,
	[8]byte{0xA4, 0x86, 0xCC, 0x10, 0xFA, 0xB1, 0x52, 0xF9}}

type IVideoDisplayProperties2Interface interface {
	win32.IInspectableInterface
	Get_Genres() *IVector[string]
}

type IVideoDisplayProperties2Vtbl struct {
	win32.IInspectableVtbl
	Get_Genres uintptr
}

type IVideoDisplayProperties2 struct {
	win32.IInspectable
}

func (this *IVideoDisplayProperties2) Vtbl() *IVideoDisplayProperties2Vtbl {
	return (*IVideoDisplayProperties2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoDisplayProperties2) Get_Genres() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Genres, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1FCDA5E8-BAF1-4521-980C-3BCEBB44CF38
var IID_IVideoEffectsStatics = syscall.GUID{0x1FCDA5E8, 0xBAF1, 0x4521,
	[8]byte{0x98, 0x0C, 0x3B, 0xCE, 0xBB, 0x44, 0xCF, 0x38}}

type IVideoEffectsStaticsInterface interface {
	win32.IInspectableInterface
	Get_VideoStabilization() string
}

type IVideoEffectsStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_VideoStabilization uintptr
}

type IVideoEffectsStatics struct {
	win32.IInspectable
}

func (this *IVideoEffectsStatics) Vtbl() *IVideoEffectsStaticsVtbl {
	return (*IVideoEffectsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoEffectsStatics) Get_VideoStabilization() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoStabilization, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 0CC06625-90FC-4C92-BD95-7DED21819D1C
var IID_IVideoFrame = syscall.GUID{0x0CC06625, 0x90FC, 0x4C92,
	[8]byte{0xBD, 0x95, 0x7D, 0xED, 0x21, 0x81, 0x9D, 0x1C}}

type IVideoFrameInterface interface {
	win32.IInspectableInterface
	Get_SoftwareBitmap() *ISoftwareBitmap
	CopyToAsync(frame *IVideoFrame) *IAsyncAction
	Get_Direct3DSurface() unsafe.Pointer
}

type IVideoFrameVtbl struct {
	win32.IInspectableVtbl
	Get_SoftwareBitmap  uintptr
	CopyToAsync         uintptr
	Get_Direct3DSurface uintptr
}

type IVideoFrame struct {
	win32.IInspectable
}

func (this *IVideoFrame) Vtbl() *IVideoFrameVtbl {
	return (*IVideoFrameVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoFrame) Get_SoftwareBitmap() *ISoftwareBitmap {
	var _result *ISoftwareBitmap
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SoftwareBitmap, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoFrame) CopyToAsync(frame *IVideoFrame) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CopyToAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(frame)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoFrame) Get_Direct3DSurface() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Direct3DSurface, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 3837840D-336C-4366-8D46-060798736C5D
var IID_IVideoFrame2 = syscall.GUID{0x3837840D, 0x336C, 0x4366,
	[8]byte{0x8D, 0x46, 0x06, 0x07, 0x98, 0x73, 0x6C, 0x5D}}

type IVideoFrame2Interface interface {
	win32.IInspectableInterface
	CopyToWithBoundsAsync(frame *IVideoFrame, sourceBounds *IReference[BitmapBounds], destinationBounds *IReference[BitmapBounds]) *IAsyncAction
}

type IVideoFrame2Vtbl struct {
	win32.IInspectableVtbl
	CopyToWithBoundsAsync uintptr
}

type IVideoFrame2 struct {
	win32.IInspectable
}

func (this *IVideoFrame2) Vtbl() *IVideoFrame2Vtbl {
	return (*IVideoFrame2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoFrame2) CopyToWithBoundsAsync(frame *IVideoFrame, sourceBounds *IReference[BitmapBounds], destinationBounds *IReference[BitmapBounds]) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CopyToWithBoundsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(frame)), uintptr(unsafe.Pointer(sourceBounds)), uintptr(unsafe.Pointer(destinationBounds)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 014B6D69-2228-4C92-92FF-50C380D3E776
var IID_IVideoFrameFactory = syscall.GUID{0x014B6D69, 0x2228, 0x4C92,
	[8]byte{0x92, 0xFF, 0x50, 0xC3, 0x80, 0xD3, 0xE7, 0x76}}

type IVideoFrameFactoryInterface interface {
	win32.IInspectableInterface
	Create(format BitmapPixelFormat, width int32, height int32) *IVideoFrame
	CreateWithAlpha(format BitmapPixelFormat, width int32, height int32, alpha BitmapAlphaMode) *IVideoFrame
}

type IVideoFrameFactoryVtbl struct {
	win32.IInspectableVtbl
	Create          uintptr
	CreateWithAlpha uintptr
}

type IVideoFrameFactory struct {
	win32.IInspectable
}

func (this *IVideoFrameFactory) Vtbl() *IVideoFrameFactoryVtbl {
	return (*IVideoFrameFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoFrameFactory) Create(format BitmapPixelFormat, width int32, height int32) *IVideoFrame {
	var _result *IVideoFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(format), uintptr(width), uintptr(height), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoFrameFactory) CreateWithAlpha(format BitmapPixelFormat, width int32, height int32, alpha BitmapAlphaMode) *IVideoFrame {
	var _result *IVideoFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithAlpha, uintptr(unsafe.Pointer(this)), uintptr(format), uintptr(width), uintptr(height), uintptr(alpha), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AB2A556F-6111-4B33-8EC3-2B209A02E17A
var IID_IVideoFrameStatics = syscall.GUID{0xAB2A556F, 0x6111, 0x4B33,
	[8]byte{0x8E, 0xC3, 0x2B, 0x20, 0x9A, 0x02, 0xE1, 0x7A}}

type IVideoFrameStaticsInterface interface {
	win32.IInspectableInterface
	CreateAsDirect3D11SurfaceBacked(format unsafe.Pointer, width int32, height int32) *IVideoFrame
	CreateAsDirect3D11SurfaceBackedWithDevice(format unsafe.Pointer, width int32, height int32, device unsafe.Pointer) *IVideoFrame
	CreateWithSoftwareBitmap(bitmap *ISoftwareBitmap) *IVideoFrame
	CreateWithDirect3D11Surface(surface unsafe.Pointer) *IVideoFrame
}

type IVideoFrameStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateAsDirect3D11SurfaceBacked           uintptr
	CreateAsDirect3D11SurfaceBackedWithDevice uintptr
	CreateWithSoftwareBitmap                  uintptr
	CreateWithDirect3D11Surface               uintptr
}

type IVideoFrameStatics struct {
	win32.IInspectable
}

func (this *IVideoFrameStatics) Vtbl() *IVideoFrameStaticsVtbl {
	return (*IVideoFrameStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoFrameStatics) CreateAsDirect3D11SurfaceBacked(format unsafe.Pointer, width int32, height int32) *IVideoFrame {
	var _result *IVideoFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAsDirect3D11SurfaceBacked, uintptr(unsafe.Pointer(this)), uintptr(format), uintptr(width), uintptr(height), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoFrameStatics) CreateAsDirect3D11SurfaceBackedWithDevice(format unsafe.Pointer, width int32, height int32, device unsafe.Pointer) *IVideoFrame {
	var _result *IVideoFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAsDirect3D11SurfaceBackedWithDevice, uintptr(unsafe.Pointer(this)), uintptr(format), uintptr(width), uintptr(height), uintptr(device), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoFrameStatics) CreateWithSoftwareBitmap(bitmap *ISoftwareBitmap) *IVideoFrame {
	var _result *IVideoFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithSoftwareBitmap, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bitmap)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoFrameStatics) CreateWithDirect3D11Surface(surface unsafe.Pointer) *IVideoFrame {
	var _result *IVideoFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithDirect3D11Surface, uintptr(unsafe.Pointer(this)), uintptr(surface), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type AudioBuffer struct {
	RtClass
	*IAudioBuffer
}

type AudioFrame struct {
	RtClass
	*IAudioFrame
}

func NewAudioFrame_Create(capacity uint32) *AudioFrame {
	hs := NewHStr("Windows.Media.AudioFrame")
	var pFac *IAudioFrameFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAudioFrameFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IAudioFrame
	p = pFac.Create(capacity)
	result := &AudioFrame{
		RtClass:     RtClass{PInspect: &p.IInspectable},
		IAudioFrame: p,
	}
	com.AddToScope(result)
	return result
}

type MediaMarkerTypes struct {
	RtClass
}

func NewIMediaMarkerTypesStatics() *IMediaMarkerTypesStatics {
	var p *IMediaMarkerTypesStatics
	hs := NewHStr("Windows.Media.MediaMarkerTypes")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMediaMarkerTypesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type MediaTimelineController struct {
	RtClass
	*IMediaTimelineController
}

func NewMediaTimelineController() *MediaTimelineController {
	hs := NewHStr("Windows.Media.MediaTimelineController")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &MediaTimelineController{
		RtClass:                  RtClass{PInspect: p},
		IMediaTimelineController: (*IMediaTimelineController)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type MediaTimelineControllerFailedEventArgs struct {
	RtClass
	*IMediaTimelineControllerFailedEventArgs
}

type VideoEffects struct {
	RtClass
}

func NewIVideoEffectsStatics() *IVideoEffectsStatics {
	var p *IVideoEffectsStatics
	hs := NewHStr("Windows.Media.VideoEffects")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IVideoEffectsStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type VideoFrame struct {
	RtClass
	*IVideoFrame
}

func NewVideoFrame_Create(format BitmapPixelFormat, width int32, height int32) *VideoFrame {
	hs := NewHStr("Windows.Media.VideoFrame")
	var pFac *IVideoFrameFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IVideoFrameFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IVideoFrame
	p = pFac.Create(format, width, height)
	result := &VideoFrame{
		RtClass:     RtClass{PInspect: &p.IInspectable},
		IVideoFrame: p,
	}
	com.AddToScope(result)
	return result
}

func NewVideoFrame_CreateWithAlpha(format BitmapPixelFormat, width int32, height int32, alpha BitmapAlphaMode) *VideoFrame {
	hs := NewHStr("Windows.Media.VideoFrame")
	var pFac *IVideoFrameFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IVideoFrameFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IVideoFrame
	p = pFac.CreateWithAlpha(format, width, height, alpha)
	result := &VideoFrame{
		RtClass:     RtClass{PInspect: &p.IInspectable},
		IVideoFrame: p,
	}
	com.AddToScope(result)
	return result
}

func NewIVideoFrameStatics() *IVideoFrameStatics {
	var p *IVideoFrameStatics
	hs := NewHStr("Windows.Media.VideoFrame")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IVideoFrameStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
