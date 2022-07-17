package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type GameChatMessageOrigin int32

const (
	GameChatMessageOrigin_Voice GameChatMessageOrigin = 0
	GameChatMessageOrigin_Text  GameChatMessageOrigin = 1
)

// enum
type GameChatOverlayPosition int32

const (
	GameChatOverlayPosition_BottomCenter GameChatOverlayPosition = 0
	GameChatOverlayPosition_BottomLeft   GameChatOverlayPosition = 1
	GameChatOverlayPosition_BottomRight  GameChatOverlayPosition = 2
	GameChatOverlayPosition_MiddleRight  GameChatOverlayPosition = 3
	GameChatOverlayPosition_MiddleLeft   GameChatOverlayPosition = 4
	GameChatOverlayPosition_TopCenter    GameChatOverlayPosition = 5
	GameChatOverlayPosition_TopLeft      GameChatOverlayPosition = 6
	GameChatOverlayPosition_TopRight     GameChatOverlayPosition = 7
)

// structs

type GameChatOverlayContract struct {
}

type GamingUIProviderContract struct {
}

// interfaces

// 1DB9A292-CC78-4173-BE45-B61E67283EA7
var IID_IGameBarStatics = syscall.GUID{0x1DB9A292, 0xCC78, 0x4173,
	[8]byte{0xBE, 0x45, 0xB6, 0x1E, 0x67, 0x28, 0x3E, 0xA7}}

type IGameBarStaticsInterface interface {
	win32.IInspectableInterface
	Add_VisibilityChanged(handler EventHandler[interface{}]) EventRegistrationToken
	Remove_VisibilityChanged(token EventRegistrationToken)
	Add_IsInputRedirectedChanged(handler EventHandler[interface{}]) EventRegistrationToken
	Remove_IsInputRedirectedChanged(token EventRegistrationToken)
	Get_Visible() bool
	Get_IsInputRedirected() bool
}

type IGameBarStaticsVtbl struct {
	win32.IInspectableVtbl
	Add_VisibilityChanged           uintptr
	Remove_VisibilityChanged        uintptr
	Add_IsInputRedirectedChanged    uintptr
	Remove_IsInputRedirectedChanged uintptr
	Get_Visible                     uintptr
	Get_IsInputRedirected           uintptr
}

type IGameBarStatics struct {
	win32.IInspectable
}

func (this *IGameBarStatics) Vtbl() *IGameBarStaticsVtbl {
	return (*IGameBarStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameBarStatics) Add_VisibilityChanged(handler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_VisibilityChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGameBarStatics) Remove_VisibilityChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_VisibilityChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGameBarStatics) Add_IsInputRedirectedChanged(handler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_IsInputRedirectedChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGameBarStatics) Remove_IsInputRedirectedChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_IsInputRedirectedChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGameBarStatics) Get_Visible() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Visible, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGameBarStatics) Get_IsInputRedirected() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsInputRedirected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A28201F1-3FB9-4E42-A403-7AFCE2023B1E
var IID_IGameChatMessageReceivedEventArgs = syscall.GUID{0xA28201F1, 0x3FB9, 0x4E42,
	[8]byte{0xA4, 0x03, 0x7A, 0xFC, 0xE2, 0x02, 0x3B, 0x1E}}

type IGameChatMessageReceivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_AppId() string
	Get_AppDisplayName() string
	Get_SenderName() string
	Get_Message() string
	Get_Origin() GameChatMessageOrigin
}

type IGameChatMessageReceivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_AppId          uintptr
	Get_AppDisplayName uintptr
	Get_SenderName     uintptr
	Get_Message        uintptr
	Get_Origin         uintptr
}

type IGameChatMessageReceivedEventArgs struct {
	win32.IInspectable
}

func (this *IGameChatMessageReceivedEventArgs) Vtbl() *IGameChatMessageReceivedEventArgsVtbl {
	return (*IGameChatMessageReceivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameChatMessageReceivedEventArgs) Get_AppId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGameChatMessageReceivedEventArgs) Get_AppDisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppDisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGameChatMessageReceivedEventArgs) Get_SenderName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SenderName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGameChatMessageReceivedEventArgs) Get_Message() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Message, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGameChatMessageReceivedEventArgs) Get_Origin() GameChatMessageOrigin {
	var _result GameChatMessageOrigin
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Origin, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// FBC64865-F6FC-4A48-AE07-03AC6ED43704
var IID_IGameChatOverlay = syscall.GUID{0xFBC64865, 0xF6FC, 0x4A48,
	[8]byte{0xAE, 0x07, 0x03, 0xAC, 0x6E, 0xD4, 0x37, 0x04}}

type IGameChatOverlayInterface interface {
	win32.IInspectableInterface
	Get_DesiredPosition() GameChatOverlayPosition
	Put_DesiredPosition(value GameChatOverlayPosition)
	AddMessage(sender string, message string, origin GameChatMessageOrigin)
}

type IGameChatOverlayVtbl struct {
	win32.IInspectableVtbl
	Get_DesiredPosition uintptr
	Put_DesiredPosition uintptr
	AddMessage          uintptr
}

type IGameChatOverlay struct {
	win32.IInspectable
}

func (this *IGameChatOverlay) Vtbl() *IGameChatOverlayVtbl {
	return (*IGameChatOverlayVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameChatOverlay) Get_DesiredPosition() GameChatOverlayPosition {
	var _result GameChatOverlayPosition
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesiredPosition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGameChatOverlay) Put_DesiredPosition(value GameChatOverlayPosition) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DesiredPosition, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IGameChatOverlay) AddMessage(sender string, message string, origin GameChatMessageOrigin) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddMessage, uintptr(unsafe.Pointer(this)), NewHStr(sender).Ptr, NewHStr(message).Ptr, uintptr(origin))
	_ = _hr
}

// 1E177397-59FB-4F4F-8E9A-80ACF817743C
var IID_IGameChatOverlayMessageSource = syscall.GUID{0x1E177397, 0x59FB, 0x4F4F,
	[8]byte{0x8E, 0x9A, 0x80, 0xAC, 0xF8, 0x17, 0x74, 0x3C}}

type IGameChatOverlayMessageSourceInterface interface {
	win32.IInspectableInterface
	Add_MessageReceived(handler TypedEventHandler[*IGameChatOverlayMessageSource, *IGameChatMessageReceivedEventArgs]) EventRegistrationToken
	Remove_MessageReceived(token EventRegistrationToken)
	SetDelayBeforeClosingAfterMessageReceived(value TimeSpan)
}

type IGameChatOverlayMessageSourceVtbl struct {
	win32.IInspectableVtbl
	Add_MessageReceived                       uintptr
	Remove_MessageReceived                    uintptr
	SetDelayBeforeClosingAfterMessageReceived uintptr
}

type IGameChatOverlayMessageSource struct {
	win32.IInspectable
}

func (this *IGameChatOverlayMessageSource) Vtbl() *IGameChatOverlayMessageSourceVtbl {
	return (*IGameChatOverlayMessageSourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameChatOverlayMessageSource) Add_MessageReceived(handler TypedEventHandler[*IGameChatOverlayMessageSource, *IGameChatMessageReceivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_MessageReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGameChatOverlayMessageSource) Remove_MessageReceived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_MessageReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGameChatOverlayMessageSource) SetDelayBeforeClosingAfterMessageReceived(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetDelayBeforeClosingAfterMessageReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

// 89ACF614-7867-49F7-9687-25D9DBF444D1
var IID_IGameChatOverlayStatics = syscall.GUID{0x89ACF614, 0x7867, 0x49F7,
	[8]byte{0x96, 0x87, 0x25, 0xD9, 0xDB, 0xF4, 0x44, 0xD1}}

type IGameChatOverlayStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *IGameChatOverlay
}

type IGameChatOverlayStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
}

type IGameChatOverlayStatics struct {
	win32.IInspectable
}

func (this *IGameChatOverlayStatics) Vtbl() *IGameChatOverlayStaticsVtbl {
	return (*IGameChatOverlayStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameChatOverlayStatics) GetDefault() *IGameChatOverlay {
	var _result *IGameChatOverlay
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A7B3203E-CAF7-4DED-BBD2-47DE43BB6DD5
var IID_IGameUIProviderActivatedEventArgs = syscall.GUID{0xA7B3203E, 0xCAF7, 0x4DED,
	[8]byte{0xBB, 0xD2, 0x47, 0xDE, 0x43, 0xBB, 0x6D, 0xD5}}

type IGameUIProviderActivatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_GameUIArgs() *IPropertySet
	ReportCompleted(results *IPropertySet)
}

type IGameUIProviderActivatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_GameUIArgs  uintptr
	ReportCompleted uintptr
}

type IGameUIProviderActivatedEventArgs struct {
	win32.IInspectable
}

func (this *IGameUIProviderActivatedEventArgs) Vtbl() *IGameUIProviderActivatedEventArgsVtbl {
	return (*IGameUIProviderActivatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameUIProviderActivatedEventArgs) Get_GameUIArgs() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GameUIArgs, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGameUIProviderActivatedEventArgs) ReportCompleted(results *IPropertySet) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReportCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(results)))
	_ = _hr
}

// classes

type GameBar struct {
	RtClass
}

func NewIGameBarStatics() *IGameBarStatics {
	var p *IGameBarStatics
	hs := NewHStr("Windows.Gaming.UI.GameBar")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGameBarStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
