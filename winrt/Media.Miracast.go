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
type MiracastReceiverApplySettingsStatus int32

const (
	MiracastReceiverApplySettingsStatus_Success              MiracastReceiverApplySettingsStatus = 0
	MiracastReceiverApplySettingsStatus_UnknownFailure       MiracastReceiverApplySettingsStatus = 1
	MiracastReceiverApplySettingsStatus_MiracastNotSupported MiracastReceiverApplySettingsStatus = 2
	MiracastReceiverApplySettingsStatus_AccessDenied         MiracastReceiverApplySettingsStatus = 3
	MiracastReceiverApplySettingsStatus_FriendlyNameTooLong  MiracastReceiverApplySettingsStatus = 4
	MiracastReceiverApplySettingsStatus_ModelNameTooLong     MiracastReceiverApplySettingsStatus = 5
	MiracastReceiverApplySettingsStatus_ModelNumberTooLong   MiracastReceiverApplySettingsStatus = 6
	MiracastReceiverApplySettingsStatus_InvalidSettings      MiracastReceiverApplySettingsStatus = 7
)

// enum
type MiracastReceiverAuthorizationMethod int32

const (
	MiracastReceiverAuthorizationMethod_None                  MiracastReceiverAuthorizationMethod = 0
	MiracastReceiverAuthorizationMethod_ConfirmConnection     MiracastReceiverAuthorizationMethod = 1
	MiracastReceiverAuthorizationMethod_PinDisplayIfRequested MiracastReceiverAuthorizationMethod = 2
	MiracastReceiverAuthorizationMethod_PinDisplayRequired    MiracastReceiverAuthorizationMethod = 3
)

// enum
type MiracastReceiverDisconnectReason int32

const (
	MiracastReceiverDisconnectReason_Finished               MiracastReceiverDisconnectReason = 0
	MiracastReceiverDisconnectReason_AppSpecificError       MiracastReceiverDisconnectReason = 1
	MiracastReceiverDisconnectReason_ConnectionNotAccepted  MiracastReceiverDisconnectReason = 2
	MiracastReceiverDisconnectReason_DisconnectedByUser     MiracastReceiverDisconnectReason = 3
	MiracastReceiverDisconnectReason_FailedToStartStreaming MiracastReceiverDisconnectReason = 4
	MiracastReceiverDisconnectReason_MediaDecodingError     MiracastReceiverDisconnectReason = 5
	MiracastReceiverDisconnectReason_MediaStreamingError    MiracastReceiverDisconnectReason = 6
	MiracastReceiverDisconnectReason_MediaDecryptionError   MiracastReceiverDisconnectReason = 7
)

// enum
type MiracastReceiverGameControllerDeviceUsageMode int32

const (
	MiracastReceiverGameControllerDeviceUsageMode_AsGameController   MiracastReceiverGameControllerDeviceUsageMode = 0
	MiracastReceiverGameControllerDeviceUsageMode_AsMouseAndKeyboard MiracastReceiverGameControllerDeviceUsageMode = 1
)

// enum
type MiracastReceiverListeningStatus int32

const (
	MiracastReceiverListeningStatus_NotListening        MiracastReceiverListeningStatus = 0
	MiracastReceiverListeningStatus_Listening           MiracastReceiverListeningStatus = 1
	MiracastReceiverListeningStatus_ConnectionPending   MiracastReceiverListeningStatus = 2
	MiracastReceiverListeningStatus_Connected           MiracastReceiverListeningStatus = 3
	MiracastReceiverListeningStatus_DisabledByPolicy    MiracastReceiverListeningStatus = 4
	MiracastReceiverListeningStatus_TemporarilyDisabled MiracastReceiverListeningStatus = 5
)

// enum
type MiracastReceiverSessionStartStatus int32

const (
	MiracastReceiverSessionStartStatus_Success              MiracastReceiverSessionStartStatus = 0
	MiracastReceiverSessionStartStatus_UnknownFailure       MiracastReceiverSessionStartStatus = 1
	MiracastReceiverSessionStartStatus_MiracastNotSupported MiracastReceiverSessionStartStatus = 2
	MiracastReceiverSessionStartStatus_AccessDenied         MiracastReceiverSessionStartStatus = 3
)

// enum
type MiracastReceiverWiFiStatus int32

const (
	MiracastReceiverWiFiStatus_MiracastSupportUndetermined MiracastReceiverWiFiStatus = 0
	MiracastReceiverWiFiStatus_MiracastNotSupported        MiracastReceiverWiFiStatus = 1
	MiracastReceiverWiFiStatus_MiracastSupportNotOptimized MiracastReceiverWiFiStatus = 2
	MiracastReceiverWiFiStatus_MiracastSupported           MiracastReceiverWiFiStatus = 3
)

// enum
type MiracastTransmitterAuthorizationStatus int32

const (
	MiracastTransmitterAuthorizationStatus_Undecided    MiracastTransmitterAuthorizationStatus = 0
	MiracastTransmitterAuthorizationStatus_Allowed      MiracastTransmitterAuthorizationStatus = 1
	MiracastTransmitterAuthorizationStatus_AlwaysPrompt MiracastTransmitterAuthorizationStatus = 2
	MiracastTransmitterAuthorizationStatus_Blocked      MiracastTransmitterAuthorizationStatus = 3
)

// interfaces

// 7A315258-E444-51B4-AFF7-B88DAA1229E0
var IID_IMiracastReceiver = syscall.GUID{0x7A315258, 0xE444, 0x51B4,
	[8]byte{0xAF, 0xF7, 0xB8, 0x8D, 0xAA, 0x12, 0x29, 0xE0}}

type IMiracastReceiverInterface interface {
	win32.IInspectableInterface
	GetDefaultSettings() *IMiracastReceiverSettings
	GetCurrentSettings() *IMiracastReceiverSettings
	GetCurrentSettingsAsync() *IAsyncOperation[*IMiracastReceiverSettings]
	DisconnectAllAndApplySettings(settings *IMiracastReceiverSettings) *IMiracastReceiverApplySettingsResult
	DisconnectAllAndApplySettingsAsync(settings *IMiracastReceiverSettings) *IAsyncOperation[*IMiracastReceiverApplySettingsResult]
	GetStatus() *IMiracastReceiverStatus
	GetStatusAsync() *IAsyncOperation[*IMiracastReceiverStatus]
	Add_StatusChanged(handler TypedEventHandler[*IMiracastReceiver, interface{}]) EventRegistrationToken
	Remove_StatusChanged(token EventRegistrationToken)
	CreateSession(view *ICoreApplicationView) *IMiracastReceiverSession
	CreateSessionAsync(view *ICoreApplicationView) *IAsyncOperation[*IMiracastReceiverSession]
	ClearKnownTransmitters()
	RemoveKnownTransmitter(transmitter *IMiracastTransmitter)
}

type IMiracastReceiverVtbl struct {
	win32.IInspectableVtbl
	GetDefaultSettings                 uintptr
	GetCurrentSettings                 uintptr
	GetCurrentSettingsAsync            uintptr
	DisconnectAllAndApplySettings      uintptr
	DisconnectAllAndApplySettingsAsync uintptr
	GetStatus                          uintptr
	GetStatusAsync                     uintptr
	Add_StatusChanged                  uintptr
	Remove_StatusChanged               uintptr
	CreateSession                      uintptr
	CreateSessionAsync                 uintptr
	ClearKnownTransmitters             uintptr
	RemoveKnownTransmitter             uintptr
}

type IMiracastReceiver struct {
	win32.IInspectable
}

func (this *IMiracastReceiver) Vtbl() *IMiracastReceiverVtbl {
	return (*IMiracastReceiverVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMiracastReceiver) GetDefaultSettings() *IMiracastReceiverSettings {
	var _result *IMiracastReceiverSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefaultSettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMiracastReceiver) GetCurrentSettings() *IMiracastReceiverSettings {
	var _result *IMiracastReceiverSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentSettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMiracastReceiver) GetCurrentSettingsAsync() *IAsyncOperation[*IMiracastReceiverSettings] {
	var _result *IAsyncOperation[*IMiracastReceiverSettings]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentSettingsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMiracastReceiver) DisconnectAllAndApplySettings(settings *IMiracastReceiverSettings) *IMiracastReceiverApplySettingsResult {
	var _result *IMiracastReceiverApplySettingsResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DisconnectAllAndApplySettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(settings)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMiracastReceiver) DisconnectAllAndApplySettingsAsync(settings *IMiracastReceiverSettings) *IAsyncOperation[*IMiracastReceiverApplySettingsResult] {
	var _result *IAsyncOperation[*IMiracastReceiverApplySettingsResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DisconnectAllAndApplySettingsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(settings)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMiracastReceiver) GetStatus() *IMiracastReceiverStatus {
	var _result *IMiracastReceiverStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMiracastReceiver) GetStatusAsync() *IAsyncOperation[*IMiracastReceiverStatus] {
	var _result *IAsyncOperation[*IMiracastReceiverStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStatusAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMiracastReceiver) Add_StatusChanged(handler TypedEventHandler[*IMiracastReceiver, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StatusChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiver) Remove_StatusChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StatusChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMiracastReceiver) CreateSession(view *ICoreApplicationView) *IMiracastReceiverSession {
	var _result *IMiracastReceiverSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSession, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(view)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMiracastReceiver) CreateSessionAsync(view *ICoreApplicationView) *IAsyncOperation[*IMiracastReceiverSession] {
	var _result *IAsyncOperation[*IMiracastReceiverSession]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSessionAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(view)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMiracastReceiver) ClearKnownTransmitters() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearKnownTransmitters, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IMiracastReceiver) RemoveKnownTransmitter(transmitter *IMiracastTransmitter) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveKnownTransmitter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(transmitter)))
	_ = _hr
}

// D0AA6272-09CD-58E1-A4F2-5D5143D312F9
var IID_IMiracastReceiverApplySettingsResult = syscall.GUID{0xD0AA6272, 0x09CD, 0x58E1,
	[8]byte{0xA4, 0xF2, 0x5D, 0x51, 0x43, 0xD3, 0x12, 0xF9}}

type IMiracastReceiverApplySettingsResultInterface interface {
	win32.IInspectableInterface
	Get_Status() MiracastReceiverApplySettingsStatus
	Get_ExtendedError() HResult
}

type IMiracastReceiverApplySettingsResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status        uintptr
	Get_ExtendedError uintptr
}

type IMiracastReceiverApplySettingsResult struct {
	win32.IInspectable
}

func (this *IMiracastReceiverApplySettingsResult) Vtbl() *IMiracastReceiverApplySettingsResultVtbl {
	return (*IMiracastReceiverApplySettingsResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMiracastReceiverApplySettingsResult) Get_Status() MiracastReceiverApplySettingsStatus {
	var _result MiracastReceiverApplySettingsStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverApplySettingsResult) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 704B2F36-D2E5-551F-A854-F822B7917D28
var IID_IMiracastReceiverConnection = syscall.GUID{0x704B2F36, 0xD2E5, 0x551F,
	[8]byte{0xA8, 0x54, 0xF8, 0x22, 0xB7, 0x91, 0x7D, 0x28}}

type IMiracastReceiverConnectionInterface interface {
	win32.IInspectableInterface
	Disconnect(reason MiracastReceiverDisconnectReason)
	DisconnectWithMessage(reason MiracastReceiverDisconnectReason, message string)
	Pause()
	PauseAsync() *IAsyncAction
	Resume()
	ResumeAsync() *IAsyncAction
	Get_Transmitter() *IMiracastTransmitter
	Get_InputDevices() *IMiracastReceiverInputDevices
	Get_CursorImageChannel() *IMiracastReceiverCursorImageChannel
	Get_StreamControl() *IMiracastReceiverStreamControl
}

type IMiracastReceiverConnectionVtbl struct {
	win32.IInspectableVtbl
	Disconnect             uintptr
	DisconnectWithMessage  uintptr
	Pause                  uintptr
	PauseAsync             uintptr
	Resume                 uintptr
	ResumeAsync            uintptr
	Get_Transmitter        uintptr
	Get_InputDevices       uintptr
	Get_CursorImageChannel uintptr
	Get_StreamControl      uintptr
}

type IMiracastReceiverConnection struct {
	win32.IInspectable
}

func (this *IMiracastReceiverConnection) Vtbl() *IMiracastReceiverConnectionVtbl {
	return (*IMiracastReceiverConnectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMiracastReceiverConnection) Disconnect(reason MiracastReceiverDisconnectReason) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Disconnect, uintptr(unsafe.Pointer(this)), uintptr(reason))
	_ = _hr
}

func (this *IMiracastReceiverConnection) DisconnectWithMessage(reason MiracastReceiverDisconnectReason, message string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DisconnectWithMessage, uintptr(unsafe.Pointer(this)), uintptr(reason), NewHStr(message).Ptr)
	_ = _hr
}

func (this *IMiracastReceiverConnection) Pause() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Pause, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IMiracastReceiverConnection) PauseAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PauseAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMiracastReceiverConnection) Resume() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Resume, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IMiracastReceiverConnection) ResumeAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ResumeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMiracastReceiverConnection) Get_Transmitter() *IMiracastTransmitter {
	var _result *IMiracastTransmitter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Transmitter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMiracastReceiverConnection) Get_InputDevices() *IMiracastReceiverInputDevices {
	var _result *IMiracastReceiverInputDevices
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InputDevices, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMiracastReceiverConnection) Get_CursorImageChannel() *IMiracastReceiverCursorImageChannel {
	var _result *IMiracastReceiverCursorImageChannel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CursorImageChannel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMiracastReceiverConnection) Get_StreamControl() *IMiracastReceiverStreamControl {
	var _result *IMiracastReceiverStreamControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StreamControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7D8DFA39-307A-5C0F-94BD-D0C69D169982
var IID_IMiracastReceiverConnectionCreatedEventArgs = syscall.GUID{0x7D8DFA39, 0x307A, 0x5C0F,
	[8]byte{0x94, 0xBD, 0xD0, 0xC6, 0x9D, 0x16, 0x99, 0x82}}

type IMiracastReceiverConnectionCreatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Connection() *IMiracastReceiverConnection
	Get_Pin() string
	GetDeferral() *IDeferral
}

type IMiracastReceiverConnectionCreatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Connection uintptr
	Get_Pin        uintptr
	GetDeferral    uintptr
}

type IMiracastReceiverConnectionCreatedEventArgs struct {
	win32.IInspectable
}

func (this *IMiracastReceiverConnectionCreatedEventArgs) Vtbl() *IMiracastReceiverConnectionCreatedEventArgsVtbl {
	return (*IMiracastReceiverConnectionCreatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMiracastReceiverConnectionCreatedEventArgs) Get_Connection() *IMiracastReceiverConnection {
	var _result *IMiracastReceiverConnection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Connection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMiracastReceiverConnectionCreatedEventArgs) Get_Pin() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Pin, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMiracastReceiverConnectionCreatedEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D9AC332D-723A-5A9D-B90A-81153EFA2A0F
var IID_IMiracastReceiverCursorImageChannel = syscall.GUID{0xD9AC332D, 0x723A, 0x5A9D,
	[8]byte{0xB9, 0x0A, 0x81, 0x15, 0x3E, 0xFA, 0x2A, 0x0F}}

type IMiracastReceiverCursorImageChannelInterface interface {
	win32.IInspectableInterface
	Get_IsEnabled() bool
	Get_MaxImageSize() unsafe.Pointer
	Get_Position() unsafe.Pointer
	Get_ImageStream() *IRandomAccessStreamWithContentType
	Add_ImageStreamChanged(handler TypedEventHandler[*IMiracastReceiverCursorImageChannel, interface{}]) EventRegistrationToken
	Remove_ImageStreamChanged(token EventRegistrationToken)
	Add_PositionChanged(handler TypedEventHandler[*IMiracastReceiverCursorImageChannel, interface{}]) EventRegistrationToken
	Remove_PositionChanged(token EventRegistrationToken)
}

type IMiracastReceiverCursorImageChannelVtbl struct {
	win32.IInspectableVtbl
	Get_IsEnabled             uintptr
	Get_MaxImageSize          uintptr
	Get_Position              uintptr
	Get_ImageStream           uintptr
	Add_ImageStreamChanged    uintptr
	Remove_ImageStreamChanged uintptr
	Add_PositionChanged       uintptr
	Remove_PositionChanged    uintptr
}

type IMiracastReceiverCursorImageChannel struct {
	win32.IInspectable
}

func (this *IMiracastReceiverCursorImageChannel) Vtbl() *IMiracastReceiverCursorImageChannelVtbl {
	return (*IMiracastReceiverCursorImageChannelVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMiracastReceiverCursorImageChannel) Get_IsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverCursorImageChannel) Get_MaxImageSize() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxImageSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverCursorImageChannel) Get_Position() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverCursorImageChannel) Get_ImageStream() *IRandomAccessStreamWithContentType {
	var _result *IRandomAccessStreamWithContentType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ImageStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMiracastReceiverCursorImageChannel) Add_ImageStreamChanged(handler TypedEventHandler[*IMiracastReceiverCursorImageChannel, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ImageStreamChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverCursorImageChannel) Remove_ImageStreamChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ImageStreamChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMiracastReceiverCursorImageChannel) Add_PositionChanged(handler TypedEventHandler[*IMiracastReceiverCursorImageChannel, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PositionChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverCursorImageChannel) Remove_PositionChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PositionChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// CCDBEDFF-BD00-5B9C-8E4C-00CACF86B634
var IID_IMiracastReceiverCursorImageChannelSettings = syscall.GUID{0xCCDBEDFF, 0xBD00, 0x5B9C,
	[8]byte{0x8E, 0x4C, 0x00, 0xCA, 0xCF, 0x86, 0xB6, 0x34}}

type IMiracastReceiverCursorImageChannelSettingsInterface interface {
	win32.IInspectableInterface
	Get_IsEnabled() bool
	Put_IsEnabled(value bool)
	Get_MaxImageSize() unsafe.Pointer
	Put_MaxImageSize(value unsafe.Pointer)
}

type IMiracastReceiverCursorImageChannelSettingsVtbl struct {
	win32.IInspectableVtbl
	Get_IsEnabled    uintptr
	Put_IsEnabled    uintptr
	Get_MaxImageSize uintptr
	Put_MaxImageSize uintptr
}

type IMiracastReceiverCursorImageChannelSettings struct {
	win32.IInspectable
}

func (this *IMiracastReceiverCursorImageChannelSettings) Vtbl() *IMiracastReceiverCursorImageChannelSettingsVtbl {
	return (*IMiracastReceiverCursorImageChannelSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMiracastReceiverCursorImageChannelSettings) Get_IsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverCursorImageChannelSettings) Put_IsEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMiracastReceiverCursorImageChannelSettings) Get_MaxImageSize() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxImageSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverCursorImageChannelSettings) Put_MaxImageSize(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxImageSize, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// D9A15E5E-5FEE-57E6-B4B0-04727DB93229
var IID_IMiracastReceiverDisconnectedEventArgs = syscall.GUID{0xD9A15E5E, 0x5FEE, 0x57E6,
	[8]byte{0xB4, 0xB0, 0x04, 0x72, 0x7D, 0xB9, 0x32, 0x29}}

type IMiracastReceiverDisconnectedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Connection() *IMiracastReceiverConnection
}

type IMiracastReceiverDisconnectedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Connection uintptr
}

type IMiracastReceiverDisconnectedEventArgs struct {
	win32.IInspectable
}

func (this *IMiracastReceiverDisconnectedEventArgs) Vtbl() *IMiracastReceiverDisconnectedEventArgsVtbl {
	return (*IMiracastReceiverDisconnectedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMiracastReceiverDisconnectedEventArgs) Get_Connection() *IMiracastReceiverConnection {
	var _result *IMiracastReceiverConnection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Connection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2D7171E8-BED4-5118-A058-E2477EB5888D
var IID_IMiracastReceiverGameControllerDevice = syscall.GUID{0x2D7171E8, 0xBED4, 0x5118,
	[8]byte{0xA0, 0x58, 0xE2, 0x47, 0x7E, 0xB5, 0x88, 0x8D}}

type IMiracastReceiverGameControllerDeviceInterface interface {
	win32.IInspectableInterface
	Get_TransmitInput() bool
	Put_TransmitInput(value bool)
	Get_IsRequestedByTransmitter() bool
	Get_IsTransmittingInput() bool
	Get_Mode() MiracastReceiverGameControllerDeviceUsageMode
	Put_Mode(value MiracastReceiverGameControllerDeviceUsageMode)
	Add_Changed(handler TypedEventHandler[*IMiracastReceiverGameControllerDevice, interface{}]) EventRegistrationToken
	Remove_Changed(token EventRegistrationToken)
}

type IMiracastReceiverGameControllerDeviceVtbl struct {
	win32.IInspectableVtbl
	Get_TransmitInput            uintptr
	Put_TransmitInput            uintptr
	Get_IsRequestedByTransmitter uintptr
	Get_IsTransmittingInput      uintptr
	Get_Mode                     uintptr
	Put_Mode                     uintptr
	Add_Changed                  uintptr
	Remove_Changed               uintptr
}

type IMiracastReceiverGameControllerDevice struct {
	win32.IInspectable
}

func (this *IMiracastReceiverGameControllerDevice) Vtbl() *IMiracastReceiverGameControllerDeviceVtbl {
	return (*IMiracastReceiverGameControllerDeviceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMiracastReceiverGameControllerDevice) Get_TransmitInput() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TransmitInput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverGameControllerDevice) Put_TransmitInput(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TransmitInput, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMiracastReceiverGameControllerDevice) Get_IsRequestedByTransmitter() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsRequestedByTransmitter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverGameControllerDevice) Get_IsTransmittingInput() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsTransmittingInput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverGameControllerDevice) Get_Mode() MiracastReceiverGameControllerDeviceUsageMode {
	var _result MiracastReceiverGameControllerDeviceUsageMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverGameControllerDevice) Put_Mode(value MiracastReceiverGameControllerDeviceUsageMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Mode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMiracastReceiverGameControllerDevice) Add_Changed(handler TypedEventHandler[*IMiracastReceiverGameControllerDevice, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Changed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverGameControllerDevice) Remove_Changed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Changed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// DA35BB02-28AA-5EE8-96F5-A42901C66F00
var IID_IMiracastReceiverInputDevices = syscall.GUID{0xDA35BB02, 0x28AA, 0x5EE8,
	[8]byte{0x96, 0xF5, 0xA4, 0x29, 0x01, 0xC6, 0x6F, 0x00}}

type IMiracastReceiverInputDevicesInterface interface {
	win32.IInspectableInterface
	Get_Keyboard() *IMiracastReceiverKeyboardDevice
	Get_GameController() *IMiracastReceiverGameControllerDevice
}

type IMiracastReceiverInputDevicesVtbl struct {
	win32.IInspectableVtbl
	Get_Keyboard       uintptr
	Get_GameController uintptr
}

type IMiracastReceiverInputDevices struct {
	win32.IInspectable
}

func (this *IMiracastReceiverInputDevices) Vtbl() *IMiracastReceiverInputDevicesVtbl {
	return (*IMiracastReceiverInputDevicesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMiracastReceiverInputDevices) Get_Keyboard() *IMiracastReceiverKeyboardDevice {
	var _result *IMiracastReceiverKeyboardDevice
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Keyboard, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMiracastReceiverInputDevices) Get_GameController() *IMiracastReceiverGameControllerDevice {
	var _result *IMiracastReceiverGameControllerDevice
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GameController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BEB67272-06C0-54FF-AC96-217464FF2501
var IID_IMiracastReceiverKeyboardDevice = syscall.GUID{0xBEB67272, 0x06C0, 0x54FF,
	[8]byte{0xAC, 0x96, 0x21, 0x74, 0x64, 0xFF, 0x25, 0x01}}

type IMiracastReceiverKeyboardDeviceInterface interface {
	win32.IInspectableInterface
	Get_TransmitInput() bool
	Put_TransmitInput(value bool)
	Get_IsRequestedByTransmitter() bool
	Get_IsTransmittingInput() bool
	Add_Changed(handler TypedEventHandler[*IMiracastReceiverKeyboardDevice, interface{}]) EventRegistrationToken
	Remove_Changed(token EventRegistrationToken)
}

type IMiracastReceiverKeyboardDeviceVtbl struct {
	win32.IInspectableVtbl
	Get_TransmitInput            uintptr
	Put_TransmitInput            uintptr
	Get_IsRequestedByTransmitter uintptr
	Get_IsTransmittingInput      uintptr
	Add_Changed                  uintptr
	Remove_Changed               uintptr
}

type IMiracastReceiverKeyboardDevice struct {
	win32.IInspectable
}

func (this *IMiracastReceiverKeyboardDevice) Vtbl() *IMiracastReceiverKeyboardDeviceVtbl {
	return (*IMiracastReceiverKeyboardDeviceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMiracastReceiverKeyboardDevice) Get_TransmitInput() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TransmitInput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverKeyboardDevice) Put_TransmitInput(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TransmitInput, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMiracastReceiverKeyboardDevice) Get_IsRequestedByTransmitter() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsRequestedByTransmitter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverKeyboardDevice) Get_IsTransmittingInput() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsTransmittingInput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverKeyboardDevice) Add_Changed(handler TypedEventHandler[*IMiracastReceiverKeyboardDevice, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Changed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverKeyboardDevice) Remove_Changed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Changed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 17CF519E-1246-531D-945A-6B158E39C3AA
var IID_IMiracastReceiverMediaSourceCreatedEventArgs = syscall.GUID{0x17CF519E, 0x1246, 0x531D,
	[8]byte{0x94, 0x5A, 0x6B, 0x15, 0x8E, 0x39, 0xC3, 0xAA}}

type IMiracastReceiverMediaSourceCreatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Connection() *IMiracastReceiverConnection
	Get_MediaSource() *IMediaSource2
	Get_CursorImageChannelSettings() *IMiracastReceiverCursorImageChannelSettings
	GetDeferral() *IDeferral
}

type IMiracastReceiverMediaSourceCreatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Connection                 uintptr
	Get_MediaSource                uintptr
	Get_CursorImageChannelSettings uintptr
	GetDeferral                    uintptr
}

type IMiracastReceiverMediaSourceCreatedEventArgs struct {
	win32.IInspectable
}

func (this *IMiracastReceiverMediaSourceCreatedEventArgs) Vtbl() *IMiracastReceiverMediaSourceCreatedEventArgsVtbl {
	return (*IMiracastReceiverMediaSourceCreatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMiracastReceiverMediaSourceCreatedEventArgs) Get_Connection() *IMiracastReceiverConnection {
	var _result *IMiracastReceiverConnection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Connection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMiracastReceiverMediaSourceCreatedEventArgs) Get_MediaSource() *IMediaSource2 {
	var _result *IMediaSource2
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMiracastReceiverMediaSourceCreatedEventArgs) Get_CursorImageChannelSettings() *IMiracastReceiverCursorImageChannelSettings {
	var _result *IMiracastReceiverCursorImageChannelSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CursorImageChannelSettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMiracastReceiverMediaSourceCreatedEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1D2BCDB4-EF8B-5209-BFC9-C32116504803
var IID_IMiracastReceiverSession = syscall.GUID{0x1D2BCDB4, 0xEF8B, 0x5209,
	[8]byte{0xBF, 0xC9, 0xC3, 0x21, 0x16, 0x50, 0x48, 0x03}}

type IMiracastReceiverSessionInterface interface {
	win32.IInspectableInterface
	Add_ConnectionCreated(handler TypedEventHandler[*IMiracastReceiverSession, *IMiracastReceiverConnectionCreatedEventArgs]) EventRegistrationToken
	Remove_ConnectionCreated(token EventRegistrationToken)
	Add_MediaSourceCreated(handler TypedEventHandler[*IMiracastReceiverSession, *IMiracastReceiverMediaSourceCreatedEventArgs]) EventRegistrationToken
	Remove_MediaSourceCreated(token EventRegistrationToken)
	Add_Disconnected(handler TypedEventHandler[*IMiracastReceiverSession, *IMiracastReceiverDisconnectedEventArgs]) EventRegistrationToken
	Remove_Disconnected(token EventRegistrationToken)
	Get_AllowConnectionTakeover() bool
	Put_AllowConnectionTakeover(value bool)
	Get_MaxSimultaneousConnections() int32
	Put_MaxSimultaneousConnections(value int32)
	Start() *IMiracastReceiverSessionStartResult
	StartAsync() *IAsyncOperation[*IMiracastReceiverSessionStartResult]
}

type IMiracastReceiverSessionVtbl struct {
	win32.IInspectableVtbl
	Add_ConnectionCreated          uintptr
	Remove_ConnectionCreated       uintptr
	Add_MediaSourceCreated         uintptr
	Remove_MediaSourceCreated      uintptr
	Add_Disconnected               uintptr
	Remove_Disconnected            uintptr
	Get_AllowConnectionTakeover    uintptr
	Put_AllowConnectionTakeover    uintptr
	Get_MaxSimultaneousConnections uintptr
	Put_MaxSimultaneousConnections uintptr
	Start                          uintptr
	StartAsync                     uintptr
}

type IMiracastReceiverSession struct {
	win32.IInspectable
}

func (this *IMiracastReceiverSession) Vtbl() *IMiracastReceiverSessionVtbl {
	return (*IMiracastReceiverSessionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMiracastReceiverSession) Add_ConnectionCreated(handler TypedEventHandler[*IMiracastReceiverSession, *IMiracastReceiverConnectionCreatedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ConnectionCreated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverSession) Remove_ConnectionCreated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ConnectionCreated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMiracastReceiverSession) Add_MediaSourceCreated(handler TypedEventHandler[*IMiracastReceiverSession, *IMiracastReceiverMediaSourceCreatedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_MediaSourceCreated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverSession) Remove_MediaSourceCreated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_MediaSourceCreated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMiracastReceiverSession) Add_Disconnected(handler TypedEventHandler[*IMiracastReceiverSession, *IMiracastReceiverDisconnectedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Disconnected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverSession) Remove_Disconnected(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Disconnected, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMiracastReceiverSession) Get_AllowConnectionTakeover() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllowConnectionTakeover, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverSession) Put_AllowConnectionTakeover(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AllowConnectionTakeover, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMiracastReceiverSession) Get_MaxSimultaneousConnections() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxSimultaneousConnections, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverSession) Put_MaxSimultaneousConnections(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxSimultaneousConnections, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMiracastReceiverSession) Start() *IMiracastReceiverSessionStartResult {
	var _result *IMiracastReceiverSessionStartResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMiracastReceiverSession) StartAsync() *IAsyncOperation[*IMiracastReceiverSessionStartResult] {
	var _result *IAsyncOperation[*IMiracastReceiverSessionStartResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B7C573EE-40CA-51FF-95F2-C9DE34F2E90E
var IID_IMiracastReceiverSessionStartResult = syscall.GUID{0xB7C573EE, 0x40CA, 0x51FF,
	[8]byte{0x95, 0xF2, 0xC9, 0xDE, 0x34, 0xF2, 0xE9, 0x0E}}

type IMiracastReceiverSessionStartResultInterface interface {
	win32.IInspectableInterface
	Get_Status() MiracastReceiverSessionStartStatus
	Get_ExtendedError() HResult
}

type IMiracastReceiverSessionStartResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status        uintptr
	Get_ExtendedError uintptr
}

type IMiracastReceiverSessionStartResult struct {
	win32.IInspectable
}

func (this *IMiracastReceiverSessionStartResult) Vtbl() *IMiracastReceiverSessionStartResultVtbl {
	return (*IMiracastReceiverSessionStartResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMiracastReceiverSessionStartResult) Get_Status() MiracastReceiverSessionStartStatus {
	var _result MiracastReceiverSessionStartStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverSessionStartResult) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 57CD2F24-C55A-5FBE-9464-EB05307705DD
var IID_IMiracastReceiverSettings = syscall.GUID{0x57CD2F24, 0xC55A, 0x5FBE,
	[8]byte{0x94, 0x64, 0xEB, 0x05, 0x30, 0x77, 0x05, 0xDD}}

type IMiracastReceiverSettingsInterface interface {
	win32.IInspectableInterface
	Get_FriendlyName() string
	Put_FriendlyName(value string)
	Get_ModelName() string
	Put_ModelName(value string)
	Get_ModelNumber() string
	Put_ModelNumber(value string)
	Get_AuthorizationMethod() MiracastReceiverAuthorizationMethod
	Put_AuthorizationMethod(value MiracastReceiverAuthorizationMethod)
	Get_RequireAuthorizationFromKnownTransmitters() bool
	Put_RequireAuthorizationFromKnownTransmitters(value bool)
}

type IMiracastReceiverSettingsVtbl struct {
	win32.IInspectableVtbl
	Get_FriendlyName                              uintptr
	Put_FriendlyName                              uintptr
	Get_ModelName                                 uintptr
	Put_ModelName                                 uintptr
	Get_ModelNumber                               uintptr
	Put_ModelNumber                               uintptr
	Get_AuthorizationMethod                       uintptr
	Put_AuthorizationMethod                       uintptr
	Get_RequireAuthorizationFromKnownTransmitters uintptr
	Put_RequireAuthorizationFromKnownTransmitters uintptr
}

type IMiracastReceiverSettings struct {
	win32.IInspectable
}

func (this *IMiracastReceiverSettings) Vtbl() *IMiracastReceiverSettingsVtbl {
	return (*IMiracastReceiverSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMiracastReceiverSettings) Get_FriendlyName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FriendlyName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMiracastReceiverSettings) Put_FriendlyName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FriendlyName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMiracastReceiverSettings) Get_ModelName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ModelName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMiracastReceiverSettings) Put_ModelName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ModelName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMiracastReceiverSettings) Get_ModelNumber() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ModelNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMiracastReceiverSettings) Put_ModelNumber(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ModelNumber, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMiracastReceiverSettings) Get_AuthorizationMethod() MiracastReceiverAuthorizationMethod {
	var _result MiracastReceiverAuthorizationMethod
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AuthorizationMethod, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverSettings) Put_AuthorizationMethod(value MiracastReceiverAuthorizationMethod) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AuthorizationMethod, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMiracastReceiverSettings) Get_RequireAuthorizationFromKnownTransmitters() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequireAuthorizationFromKnownTransmitters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverSettings) Put_RequireAuthorizationFromKnownTransmitters(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RequireAuthorizationFromKnownTransmitters, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// C28A5591-23AB-519E-AD09-90BFF6DCC87E
var IID_IMiracastReceiverStatus = syscall.GUID{0xC28A5591, 0x23AB, 0x519E,
	[8]byte{0xAD, 0x09, 0x90, 0xBF, 0xF6, 0xDC, 0xC8, 0x7E}}

type IMiracastReceiverStatusInterface interface {
	win32.IInspectableInterface
	Get_ListeningStatus() MiracastReceiverListeningStatus
	Get_WiFiStatus() MiracastReceiverWiFiStatus
	Get_IsConnectionTakeoverSupported() bool
	Get_MaxSimultaneousConnections() int32
	Get_KnownTransmitters() *IVectorView[*IMiracastTransmitter]
}

type IMiracastReceiverStatusVtbl struct {
	win32.IInspectableVtbl
	Get_ListeningStatus               uintptr
	Get_WiFiStatus                    uintptr
	Get_IsConnectionTakeoverSupported uintptr
	Get_MaxSimultaneousConnections    uintptr
	Get_KnownTransmitters             uintptr
}

type IMiracastReceiverStatus struct {
	win32.IInspectable
}

func (this *IMiracastReceiverStatus) Vtbl() *IMiracastReceiverStatusVtbl {
	return (*IMiracastReceiverStatusVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMiracastReceiverStatus) Get_ListeningStatus() MiracastReceiverListeningStatus {
	var _result MiracastReceiverListeningStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ListeningStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverStatus) Get_WiFiStatus() MiracastReceiverWiFiStatus {
	var _result MiracastReceiverWiFiStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WiFiStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverStatus) Get_IsConnectionTakeoverSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsConnectionTakeoverSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverStatus) Get_MaxSimultaneousConnections() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxSimultaneousConnections, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverStatus) Get_KnownTransmitters() *IVectorView[*IMiracastTransmitter] {
	var _result *IVectorView[*IMiracastTransmitter]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KnownTransmitters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 38EA2D8B-2769-5AD7-8A8A-254B9DF7BA82
var IID_IMiracastReceiverStreamControl = syscall.GUID{0x38EA2D8B, 0x2769, 0x5AD7,
	[8]byte{0x8A, 0x8A, 0x25, 0x4B, 0x9D, 0xF7, 0xBA, 0x82}}

type IMiracastReceiverStreamControlInterface interface {
	win32.IInspectableInterface
	GetVideoStreamSettings() *IMiracastReceiverVideoStreamSettings
	GetVideoStreamSettingsAsync() *IAsyncOperation[*IMiracastReceiverVideoStreamSettings]
	SuggestVideoStreamSettings(settings *IMiracastReceiverVideoStreamSettings)
	SuggestVideoStreamSettingsAsync(settings *IMiracastReceiverVideoStreamSettings) *IAsyncAction
	Get_MuteAudio() bool
	Put_MuteAudio(value bool)
}

type IMiracastReceiverStreamControlVtbl struct {
	win32.IInspectableVtbl
	GetVideoStreamSettings          uintptr
	GetVideoStreamSettingsAsync     uintptr
	SuggestVideoStreamSettings      uintptr
	SuggestVideoStreamSettingsAsync uintptr
	Get_MuteAudio                   uintptr
	Put_MuteAudio                   uintptr
}

type IMiracastReceiverStreamControl struct {
	win32.IInspectable
}

func (this *IMiracastReceiverStreamControl) Vtbl() *IMiracastReceiverStreamControlVtbl {
	return (*IMiracastReceiverStreamControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMiracastReceiverStreamControl) GetVideoStreamSettings() *IMiracastReceiverVideoStreamSettings {
	var _result *IMiracastReceiverVideoStreamSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetVideoStreamSettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMiracastReceiverStreamControl) GetVideoStreamSettingsAsync() *IAsyncOperation[*IMiracastReceiverVideoStreamSettings] {
	var _result *IAsyncOperation[*IMiracastReceiverVideoStreamSettings]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetVideoStreamSettingsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMiracastReceiverStreamControl) SuggestVideoStreamSettings(settings *IMiracastReceiverVideoStreamSettings) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SuggestVideoStreamSettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(settings)))
	_ = _hr
}

func (this *IMiracastReceiverStreamControl) SuggestVideoStreamSettingsAsync(settings *IMiracastReceiverVideoStreamSettings) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SuggestVideoStreamSettingsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(settings)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMiracastReceiverStreamControl) Get_MuteAudio() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MuteAudio, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverStreamControl) Put_MuteAudio(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MuteAudio, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 169B5E1B-149D-52D0-B126-6F89744E4F50
var IID_IMiracastReceiverVideoStreamSettings = syscall.GUID{0x169B5E1B, 0x149D, 0x52D0,
	[8]byte{0xB1, 0x26, 0x6F, 0x89, 0x74, 0x4E, 0x4F, 0x50}}

type IMiracastReceiverVideoStreamSettingsInterface interface {
	win32.IInspectableInterface
	Get_Size() unsafe.Pointer
	Put_Size(value unsafe.Pointer)
	Get_Bitrate() int32
	Put_Bitrate(value int32)
}

type IMiracastReceiverVideoStreamSettingsVtbl struct {
	win32.IInspectableVtbl
	Get_Size    uintptr
	Put_Size    uintptr
	Get_Bitrate uintptr
	Put_Bitrate uintptr
}

type IMiracastReceiverVideoStreamSettings struct {
	win32.IInspectable
}

func (this *IMiracastReceiverVideoStreamSettings) Vtbl() *IMiracastReceiverVideoStreamSettingsVtbl {
	return (*IMiracastReceiverVideoStreamSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMiracastReceiverVideoStreamSettings) Get_Size() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Size, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverVideoStreamSettings) Put_Size(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Size, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMiracastReceiverVideoStreamSettings) Get_Bitrate() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bitrate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastReceiverVideoStreamSettings) Put_Bitrate(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Bitrate, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 342D79FD-2E64-5508-8A30-833D1EAC70D0
var IID_IMiracastTransmitter = syscall.GUID{0x342D79FD, 0x2E64, 0x5508,
	[8]byte{0x8A, 0x30, 0x83, 0x3D, 0x1E, 0xAC, 0x70, 0xD0}}

type IMiracastTransmitterInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Put_Name(value string)
	Get_AuthorizationStatus() MiracastTransmitterAuthorizationStatus
	Put_AuthorizationStatus(value MiracastTransmitterAuthorizationStatus)
	GetConnections() *IVectorView[*IMiracastReceiverConnection]
	Get_MacAddress() string
	Get_LastConnectionTime() DateTime
}

type IMiracastTransmitterVtbl struct {
	win32.IInspectableVtbl
	Get_Name                uintptr
	Put_Name                uintptr
	Get_AuthorizationStatus uintptr
	Put_AuthorizationStatus uintptr
	GetConnections          uintptr
	Get_MacAddress          uintptr
	Get_LastConnectionTime  uintptr
}

type IMiracastTransmitter struct {
	win32.IInspectable
}

func (this *IMiracastTransmitter) Vtbl() *IMiracastTransmitterVtbl {
	return (*IMiracastTransmitterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMiracastTransmitter) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMiracastTransmitter) Put_Name(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Name, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMiracastTransmitter) Get_AuthorizationStatus() MiracastTransmitterAuthorizationStatus {
	var _result MiracastTransmitterAuthorizationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AuthorizationStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMiracastTransmitter) Put_AuthorizationStatus(value MiracastTransmitterAuthorizationStatus) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AuthorizationStatus, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMiracastTransmitter) GetConnections() *IVectorView[*IMiracastReceiverConnection] {
	var _result *IVectorView[*IMiracastReceiverConnection]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetConnections, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMiracastTransmitter) Get_MacAddress() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MacAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMiracastTransmitter) Get_LastConnectionTime() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LastConnectionTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type MiracastReceiver struct {
	RtClass
	*IMiracastReceiver
}

func NewMiracastReceiver() *MiracastReceiver {
	hs := NewHStr("Windows.Media.Miracast.MiracastReceiver")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &MiracastReceiver{
		RtClass:           RtClass{PInspect: p},
		IMiracastReceiver: (*IMiracastReceiver)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type MiracastReceiverApplySettingsResult struct {
	RtClass
	*IMiracastReceiverApplySettingsResult
}

type MiracastReceiverConnection struct {
	RtClass
	*IMiracastReceiverConnection
}

type MiracastReceiverConnectionCreatedEventArgs struct {
	RtClass
	*IMiracastReceiverConnectionCreatedEventArgs
}

type MiracastReceiverCursorImageChannel struct {
	RtClass
	*IMiracastReceiverCursorImageChannel
}

type MiracastReceiverCursorImageChannelSettings struct {
	RtClass
	*IMiracastReceiverCursorImageChannelSettings
}

type MiracastReceiverDisconnectedEventArgs struct {
	RtClass
	*IMiracastReceiverDisconnectedEventArgs
}

type MiracastReceiverGameControllerDevice struct {
	RtClass
	*IMiracastReceiverGameControllerDevice
}

type MiracastReceiverInputDevices struct {
	RtClass
	*IMiracastReceiverInputDevices
}

type MiracastReceiverKeyboardDevice struct {
	RtClass
	*IMiracastReceiverKeyboardDevice
}

type MiracastReceiverMediaSourceCreatedEventArgs struct {
	RtClass
	*IMiracastReceiverMediaSourceCreatedEventArgs
}

type MiracastReceiverSession struct {
	RtClass
	*IMiracastReceiverSession
}

type MiracastReceiverSessionStartResult struct {
	RtClass
	*IMiracastReceiverSessionStartResult
}

type MiracastReceiverSettings struct {
	RtClass
	*IMiracastReceiverSettings
}

type MiracastReceiverStatus struct {
	RtClass
	*IMiracastReceiverStatus
}

type MiracastReceiverStreamControl struct {
	RtClass
	*IMiracastReceiverStreamControl
}

type MiracastReceiverVideoStreamSettings struct {
	RtClass
	*IMiracastReceiverVideoStreamSettings
}

type MiracastTransmitter struct {
	RtClass
	*IMiracastTransmitter
}
