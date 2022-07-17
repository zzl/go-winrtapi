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
// flags
type SpatialGestureSettings uint32

const (
	SpatialGestureSettings_None                  SpatialGestureSettings = 0
	SpatialGestureSettings_Tap                   SpatialGestureSettings = 1
	SpatialGestureSettings_DoubleTap             SpatialGestureSettings = 2
	SpatialGestureSettings_Hold                  SpatialGestureSettings = 4
	SpatialGestureSettings_ManipulationTranslate SpatialGestureSettings = 8
	SpatialGestureSettings_NavigationX           SpatialGestureSettings = 16
	SpatialGestureSettings_NavigationY           SpatialGestureSettings = 32
	SpatialGestureSettings_NavigationZ           SpatialGestureSettings = 64
	SpatialGestureSettings_NavigationRailsX      SpatialGestureSettings = 128
	SpatialGestureSettings_NavigationRailsY      SpatialGestureSettings = 256
	SpatialGestureSettings_NavigationRailsZ      SpatialGestureSettings = 512
)

// enum
type SpatialInteractionPressKind int32

const (
	SpatialInteractionPressKind_None       SpatialInteractionPressKind = 0
	SpatialInteractionPressKind_Select     SpatialInteractionPressKind = 1
	SpatialInteractionPressKind_Menu       SpatialInteractionPressKind = 2
	SpatialInteractionPressKind_Grasp      SpatialInteractionPressKind = 3
	SpatialInteractionPressKind_Touchpad   SpatialInteractionPressKind = 4
	SpatialInteractionPressKind_Thumbstick SpatialInteractionPressKind = 5
)

// enum
type SpatialInteractionSourceHandedness int32

const (
	SpatialInteractionSourceHandedness_Unspecified SpatialInteractionSourceHandedness = 0
	SpatialInteractionSourceHandedness_Left        SpatialInteractionSourceHandedness = 1
	SpatialInteractionSourceHandedness_Right       SpatialInteractionSourceHandedness = 2
)

// enum
type SpatialInteractionSourceKind int32

const (
	SpatialInteractionSourceKind_Other      SpatialInteractionSourceKind = 0
	SpatialInteractionSourceKind_Hand       SpatialInteractionSourceKind = 1
	SpatialInteractionSourceKind_Voice      SpatialInteractionSourceKind = 2
	SpatialInteractionSourceKind_Controller SpatialInteractionSourceKind = 3
)

// enum
type SpatialInteractionSourcePositionAccuracy int32

const (
	SpatialInteractionSourcePositionAccuracy_High        SpatialInteractionSourcePositionAccuracy = 0
	SpatialInteractionSourcePositionAccuracy_Approximate SpatialInteractionSourcePositionAccuracy = 1
)

// interfaces

// 71605BCC-0C35-4673-ADBD-CC04CAA6EF45
var IID_ISpatialGestureRecognizer = syscall.GUID{0x71605BCC, 0x0C35, 0x4673,
	[8]byte{0xAD, 0xBD, 0xCC, 0x04, 0xCA, 0xA6, 0xEF, 0x45}}

type ISpatialGestureRecognizerInterface interface {
	win32.IInspectableInterface
	Add_RecognitionStarted(handler TypedEventHandler[*ISpatialGestureRecognizer, *ISpatialRecognitionStartedEventArgs]) EventRegistrationToken
	Remove_RecognitionStarted(token EventRegistrationToken)
	Add_RecognitionEnded(handler TypedEventHandler[*ISpatialGestureRecognizer, *ISpatialRecognitionEndedEventArgs]) EventRegistrationToken
	Remove_RecognitionEnded(token EventRegistrationToken)
	Add_Tapped(handler TypedEventHandler[*ISpatialGestureRecognizer, *ISpatialTappedEventArgs]) EventRegistrationToken
	Remove_Tapped(token EventRegistrationToken)
	Add_HoldStarted(handler TypedEventHandler[*ISpatialGestureRecognizer, *ISpatialHoldStartedEventArgs]) EventRegistrationToken
	Remove_HoldStarted(token EventRegistrationToken)
	Add_HoldCompleted(handler TypedEventHandler[*ISpatialGestureRecognizer, *ISpatialHoldCompletedEventArgs]) EventRegistrationToken
	Remove_HoldCompleted(token EventRegistrationToken)
	Add_HoldCanceled(handler TypedEventHandler[*ISpatialGestureRecognizer, *ISpatialHoldCanceledEventArgs]) EventRegistrationToken
	Remove_HoldCanceled(token EventRegistrationToken)
	Add_ManipulationStarted(handler TypedEventHandler[*ISpatialGestureRecognizer, *ISpatialManipulationStartedEventArgs]) EventRegistrationToken
	Remove_ManipulationStarted(token EventRegistrationToken)
	Add_ManipulationUpdated(handler TypedEventHandler[*ISpatialGestureRecognizer, *ISpatialManipulationUpdatedEventArgs]) EventRegistrationToken
	Remove_ManipulationUpdated(token EventRegistrationToken)
	Add_ManipulationCompleted(handler TypedEventHandler[*ISpatialGestureRecognizer, *ISpatialManipulationCompletedEventArgs]) EventRegistrationToken
	Remove_ManipulationCompleted(token EventRegistrationToken)
	Add_ManipulationCanceled(handler TypedEventHandler[*ISpatialGestureRecognizer, *ISpatialManipulationCanceledEventArgs]) EventRegistrationToken
	Remove_ManipulationCanceled(token EventRegistrationToken)
	Add_NavigationStarted(handler TypedEventHandler[*ISpatialGestureRecognizer, *ISpatialNavigationStartedEventArgs]) EventRegistrationToken
	Remove_NavigationStarted(token EventRegistrationToken)
	Add_NavigationUpdated(handler TypedEventHandler[*ISpatialGestureRecognizer, *ISpatialNavigationUpdatedEventArgs]) EventRegistrationToken
	Remove_NavigationUpdated(token EventRegistrationToken)
	Add_NavigationCompleted(handler TypedEventHandler[*ISpatialGestureRecognizer, *ISpatialNavigationCompletedEventArgs]) EventRegistrationToken
	Remove_NavigationCompleted(token EventRegistrationToken)
	Add_NavigationCanceled(handler TypedEventHandler[*ISpatialGestureRecognizer, *ISpatialNavigationCanceledEventArgs]) EventRegistrationToken
	Remove_NavigationCanceled(token EventRegistrationToken)
	CaptureInteraction(interaction *ISpatialInteraction)
	CancelPendingGestures()
	TrySetGestureSettings(settings SpatialGestureSettings) bool
	Get_GestureSettings() SpatialGestureSettings
}

type ISpatialGestureRecognizerVtbl struct {
	win32.IInspectableVtbl
	Add_RecognitionStarted       uintptr
	Remove_RecognitionStarted    uintptr
	Add_RecognitionEnded         uintptr
	Remove_RecognitionEnded      uintptr
	Add_Tapped                   uintptr
	Remove_Tapped                uintptr
	Add_HoldStarted              uintptr
	Remove_HoldStarted           uintptr
	Add_HoldCompleted            uintptr
	Remove_HoldCompleted         uintptr
	Add_HoldCanceled             uintptr
	Remove_HoldCanceled          uintptr
	Add_ManipulationStarted      uintptr
	Remove_ManipulationStarted   uintptr
	Add_ManipulationUpdated      uintptr
	Remove_ManipulationUpdated   uintptr
	Add_ManipulationCompleted    uintptr
	Remove_ManipulationCompleted uintptr
	Add_ManipulationCanceled     uintptr
	Remove_ManipulationCanceled  uintptr
	Add_NavigationStarted        uintptr
	Remove_NavigationStarted     uintptr
	Add_NavigationUpdated        uintptr
	Remove_NavigationUpdated     uintptr
	Add_NavigationCompleted      uintptr
	Remove_NavigationCompleted   uintptr
	Add_NavigationCanceled       uintptr
	Remove_NavigationCanceled    uintptr
	CaptureInteraction           uintptr
	CancelPendingGestures        uintptr
	TrySetGestureSettings        uintptr
	Get_GestureSettings          uintptr
}

type ISpatialGestureRecognizer struct {
	win32.IInspectable
}

func (this *ISpatialGestureRecognizer) Vtbl() *ISpatialGestureRecognizerVtbl {
	return (*ISpatialGestureRecognizerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialGestureRecognizer) Add_RecognitionStarted(handler TypedEventHandler[*ISpatialGestureRecognizer, *ISpatialRecognitionStartedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_RecognitionStarted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialGestureRecognizer) Remove_RecognitionStarted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_RecognitionStarted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISpatialGestureRecognizer) Add_RecognitionEnded(handler TypedEventHandler[*ISpatialGestureRecognizer, *ISpatialRecognitionEndedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_RecognitionEnded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialGestureRecognizer) Remove_RecognitionEnded(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_RecognitionEnded, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISpatialGestureRecognizer) Add_Tapped(handler TypedEventHandler[*ISpatialGestureRecognizer, *ISpatialTappedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Tapped, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialGestureRecognizer) Remove_Tapped(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Tapped, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISpatialGestureRecognizer) Add_HoldStarted(handler TypedEventHandler[*ISpatialGestureRecognizer, *ISpatialHoldStartedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_HoldStarted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialGestureRecognizer) Remove_HoldStarted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_HoldStarted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISpatialGestureRecognizer) Add_HoldCompleted(handler TypedEventHandler[*ISpatialGestureRecognizer, *ISpatialHoldCompletedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_HoldCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialGestureRecognizer) Remove_HoldCompleted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_HoldCompleted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISpatialGestureRecognizer) Add_HoldCanceled(handler TypedEventHandler[*ISpatialGestureRecognizer, *ISpatialHoldCanceledEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_HoldCanceled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialGestureRecognizer) Remove_HoldCanceled(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_HoldCanceled, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISpatialGestureRecognizer) Add_ManipulationStarted(handler TypedEventHandler[*ISpatialGestureRecognizer, *ISpatialManipulationStartedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ManipulationStarted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialGestureRecognizer) Remove_ManipulationStarted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ManipulationStarted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISpatialGestureRecognizer) Add_ManipulationUpdated(handler TypedEventHandler[*ISpatialGestureRecognizer, *ISpatialManipulationUpdatedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ManipulationUpdated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialGestureRecognizer) Remove_ManipulationUpdated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ManipulationUpdated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISpatialGestureRecognizer) Add_ManipulationCompleted(handler TypedEventHandler[*ISpatialGestureRecognizer, *ISpatialManipulationCompletedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ManipulationCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialGestureRecognizer) Remove_ManipulationCompleted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ManipulationCompleted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISpatialGestureRecognizer) Add_ManipulationCanceled(handler TypedEventHandler[*ISpatialGestureRecognizer, *ISpatialManipulationCanceledEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ManipulationCanceled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialGestureRecognizer) Remove_ManipulationCanceled(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ManipulationCanceled, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISpatialGestureRecognizer) Add_NavigationStarted(handler TypedEventHandler[*ISpatialGestureRecognizer, *ISpatialNavigationStartedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_NavigationStarted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialGestureRecognizer) Remove_NavigationStarted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_NavigationStarted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISpatialGestureRecognizer) Add_NavigationUpdated(handler TypedEventHandler[*ISpatialGestureRecognizer, *ISpatialNavigationUpdatedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_NavigationUpdated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialGestureRecognizer) Remove_NavigationUpdated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_NavigationUpdated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISpatialGestureRecognizer) Add_NavigationCompleted(handler TypedEventHandler[*ISpatialGestureRecognizer, *ISpatialNavigationCompletedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_NavigationCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialGestureRecognizer) Remove_NavigationCompleted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_NavigationCompleted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISpatialGestureRecognizer) Add_NavigationCanceled(handler TypedEventHandler[*ISpatialGestureRecognizer, *ISpatialNavigationCanceledEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_NavigationCanceled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialGestureRecognizer) Remove_NavigationCanceled(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_NavigationCanceled, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISpatialGestureRecognizer) CaptureInteraction(interaction *ISpatialInteraction) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CaptureInteraction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(interaction)))
	_ = _hr
}

func (this *ISpatialGestureRecognizer) CancelPendingGestures() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CancelPendingGestures, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *ISpatialGestureRecognizer) TrySetGestureSettings(settings SpatialGestureSettings) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TrySetGestureSettings, uintptr(unsafe.Pointer(this)), uintptr(settings), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialGestureRecognizer) Get_GestureSettings() SpatialGestureSettings {
	var _result SpatialGestureSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GestureSettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 77214186-57B9-3150-8382-698B24E264D0
var IID_ISpatialGestureRecognizerFactory = syscall.GUID{0x77214186, 0x57B9, 0x3150,
	[8]byte{0x83, 0x82, 0x69, 0x8B, 0x24, 0xE2, 0x64, 0xD0}}

type ISpatialGestureRecognizerFactoryInterface interface {
	win32.IInspectableInterface
	Create(settings SpatialGestureSettings) *ISpatialGestureRecognizer
}

type ISpatialGestureRecognizerFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ISpatialGestureRecognizerFactory struct {
	win32.IInspectable
}

func (this *ISpatialGestureRecognizerFactory) Vtbl() *ISpatialGestureRecognizerFactoryVtbl {
	return (*ISpatialGestureRecognizerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialGestureRecognizerFactory) Create(settings SpatialGestureSettings) *ISpatialGestureRecognizer {
	var _result *ISpatialGestureRecognizer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(settings), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5DFCB667-4CAA-4093-8C35-B601A839F31B
var IID_ISpatialHoldCanceledEventArgs = syscall.GUID{0x5DFCB667, 0x4CAA, 0x4093,
	[8]byte{0x8C, 0x35, 0xB6, 0x01, 0xA8, 0x39, 0xF3, 0x1B}}

type ISpatialHoldCanceledEventArgsInterface interface {
	win32.IInspectableInterface
	Get_InteractionSourceKind() SpatialInteractionSourceKind
}

type ISpatialHoldCanceledEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_InteractionSourceKind uintptr
}

type ISpatialHoldCanceledEventArgs struct {
	win32.IInspectable
}

func (this *ISpatialHoldCanceledEventArgs) Vtbl() *ISpatialHoldCanceledEventArgsVtbl {
	return (*ISpatialHoldCanceledEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialHoldCanceledEventArgs) Get_InteractionSourceKind() SpatialInteractionSourceKind {
	var _result SpatialInteractionSourceKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InteractionSourceKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 3F64470B-4CFD-43DA-8DC4-E64552173971
var IID_ISpatialHoldCompletedEventArgs = syscall.GUID{0x3F64470B, 0x4CFD, 0x43DA,
	[8]byte{0x8D, 0xC4, 0xE6, 0x45, 0x52, 0x17, 0x39, 0x71}}

type ISpatialHoldCompletedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_InteractionSourceKind() SpatialInteractionSourceKind
}

type ISpatialHoldCompletedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_InteractionSourceKind uintptr
}

type ISpatialHoldCompletedEventArgs struct {
	win32.IInspectable
}

func (this *ISpatialHoldCompletedEventArgs) Vtbl() *ISpatialHoldCompletedEventArgsVtbl {
	return (*ISpatialHoldCompletedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialHoldCompletedEventArgs) Get_InteractionSourceKind() SpatialInteractionSourceKind {
	var _result SpatialInteractionSourceKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InteractionSourceKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 8E343D79-ACB6-4144-8615-2CFBA8A3CB3F
var IID_ISpatialHoldStartedEventArgs = syscall.GUID{0x8E343D79, 0xACB6, 0x4144,
	[8]byte{0x86, 0x15, 0x2C, 0xFB, 0xA8, 0xA3, 0xCB, 0x3F}}

type ISpatialHoldStartedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_InteractionSourceKind() SpatialInteractionSourceKind
	TryGetPointerPose(coordinateSystem *ISpatialCoordinateSystem) *ISpatialPointerPose
}

type ISpatialHoldStartedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_InteractionSourceKind uintptr
	TryGetPointerPose         uintptr
}

type ISpatialHoldStartedEventArgs struct {
	win32.IInspectable
}

func (this *ISpatialHoldStartedEventArgs) Vtbl() *ISpatialHoldStartedEventArgsVtbl {
	return (*ISpatialHoldStartedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialHoldStartedEventArgs) Get_InteractionSourceKind() SpatialInteractionSourceKind {
	var _result SpatialInteractionSourceKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InteractionSourceKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialHoldStartedEventArgs) TryGetPointerPose(coordinateSystem *ISpatialCoordinateSystem) *ISpatialPointerPose {
	var _result *ISpatialPointerPose
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetPointerPose, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(coordinateSystem)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FC967639-88E6-4646-9112-4344AAEC9DFA
var IID_ISpatialInteraction = syscall.GUID{0xFC967639, 0x88E6, 0x4646,
	[8]byte{0x91, 0x12, 0x43, 0x44, 0xAA, 0xEC, 0x9D, 0xFA}}

type ISpatialInteractionInterface interface {
	win32.IInspectableInterface
	Get_SourceState() *ISpatialInteractionSourceState
}

type ISpatialInteractionVtbl struct {
	win32.IInspectableVtbl
	Get_SourceState uintptr
}

type ISpatialInteraction struct {
	win32.IInspectable
}

func (this *ISpatialInteraction) Vtbl() *ISpatialInteractionVtbl {
	return (*ISpatialInteractionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialInteraction) Get_SourceState() *ISpatialInteractionSourceState {
	var _result *ISpatialInteractionSourceState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SourceState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5F0E5BA3-0954-4E97-86C5-E7F30B114DFD
var IID_ISpatialInteractionController = syscall.GUID{0x5F0E5BA3, 0x0954, 0x4E97,
	[8]byte{0x86, 0xC5, 0xE7, 0xF3, 0x0B, 0x11, 0x4D, 0xFD}}

type ISpatialInteractionControllerInterface interface {
	win32.IInspectableInterface
	Get_HasTouchpad() bool
	Get_HasThumbstick() bool
	Get_SimpleHapticsController() *ISimpleHapticsController
	Get_VendorId() uint16
	Get_ProductId() uint16
	Get_Version() uint16
}

type ISpatialInteractionControllerVtbl struct {
	win32.IInspectableVtbl
	Get_HasTouchpad             uintptr
	Get_HasThumbstick           uintptr
	Get_SimpleHapticsController uintptr
	Get_VendorId                uintptr
	Get_ProductId               uintptr
	Get_Version                 uintptr
}

type ISpatialInteractionController struct {
	win32.IInspectable
}

func (this *ISpatialInteractionController) Vtbl() *ISpatialInteractionControllerVtbl {
	return (*ISpatialInteractionControllerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialInteractionController) Get_HasTouchpad() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasTouchpad, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionController) Get_HasThumbstick() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasThumbstick, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionController) Get_SimpleHapticsController() *ISimpleHapticsController {
	var _result *ISimpleHapticsController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SimpleHapticsController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialInteractionController) Get_VendorId() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VendorId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionController) Get_ProductId() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProductId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionController) Get_Version() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Version, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 35B6D924-C7A2-49B7-B72E-5436B2FB8F9C
var IID_ISpatialInteractionController2 = syscall.GUID{0x35B6D924, 0xC7A2, 0x49B7,
	[8]byte{0xB7, 0x2E, 0x54, 0x36, 0xB2, 0xFB, 0x8F, 0x9C}}

type ISpatialInteractionController2Interface interface {
	win32.IInspectableInterface
	TryGetRenderableModelAsync() *IAsyncOperation[*IRandomAccessStreamWithContentType]
}

type ISpatialInteractionController2Vtbl struct {
	win32.IInspectableVtbl
	TryGetRenderableModelAsync uintptr
}

type ISpatialInteractionController2 struct {
	win32.IInspectable
}

func (this *ISpatialInteractionController2) Vtbl() *ISpatialInteractionController2Vtbl {
	return (*ISpatialInteractionController2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialInteractionController2) TryGetRenderableModelAsync() *IAsyncOperation[*IRandomAccessStreamWithContentType] {
	var _result *IAsyncOperation[*IRandomAccessStreamWithContentType]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetRenderableModelAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 628466A0-9D91-4A0B-888D-165E670A8CD5
var IID_ISpatialInteractionController3 = syscall.GUID{0x628466A0, 0x9D91, 0x4A0B,
	[8]byte{0x88, 0x8D, 0x16, 0x5E, 0x67, 0x0A, 0x8C, 0xD5}}

type ISpatialInteractionController3Interface interface {
	win32.IInspectableInterface
	TryGetBatteryReport() *IBatteryReport
}

type ISpatialInteractionController3Vtbl struct {
	win32.IInspectableVtbl
	TryGetBatteryReport uintptr
}

type ISpatialInteractionController3 struct {
	win32.IInspectable
}

func (this *ISpatialInteractionController3) Vtbl() *ISpatialInteractionController3Vtbl {
	return (*ISpatialInteractionController3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialInteractionController3) TryGetBatteryReport() *IBatteryReport {
	var _result *IBatteryReport
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetBatteryReport, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 61056FB1-7BA9-4E35-B93F-9272CBA9B28B
var IID_ISpatialInteractionControllerProperties = syscall.GUID{0x61056FB1, 0x7BA9, 0x4E35,
	[8]byte{0xB9, 0x3F, 0x92, 0x72, 0xCB, 0xA9, 0xB2, 0x8B}}

type ISpatialInteractionControllerPropertiesInterface interface {
	win32.IInspectableInterface
	Get_IsTouchpadTouched() bool
	Get_IsTouchpadPressed() bool
	Get_IsThumbstickPressed() bool
	Get_ThumbstickX() float64
	Get_ThumbstickY() float64
	Get_TouchpadX() float64
	Get_TouchpadY() float64
}

type ISpatialInteractionControllerPropertiesVtbl struct {
	win32.IInspectableVtbl
	Get_IsTouchpadTouched   uintptr
	Get_IsTouchpadPressed   uintptr
	Get_IsThumbstickPressed uintptr
	Get_ThumbstickX         uintptr
	Get_ThumbstickY         uintptr
	Get_TouchpadX           uintptr
	Get_TouchpadY           uintptr
}

type ISpatialInteractionControllerProperties struct {
	win32.IInspectable
}

func (this *ISpatialInteractionControllerProperties) Vtbl() *ISpatialInteractionControllerPropertiesVtbl {
	return (*ISpatialInteractionControllerPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialInteractionControllerProperties) Get_IsTouchpadTouched() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsTouchpadTouched, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionControllerProperties) Get_IsTouchpadPressed() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsTouchpadPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionControllerProperties) Get_IsThumbstickPressed() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsThumbstickPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionControllerProperties) Get_ThumbstickX() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ThumbstickX, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionControllerProperties) Get_ThumbstickY() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ThumbstickY, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionControllerProperties) Get_TouchpadX() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TouchpadX, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionControllerProperties) Get_TouchpadY() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TouchpadY, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 075878E4-5961-3B41-9DFB-CEA5D89CC38A
var IID_ISpatialInteractionDetectedEventArgs = syscall.GUID{0x075878E4, 0x5961, 0x3B41,
	[8]byte{0x9D, 0xFB, 0xCE, 0xA5, 0xD8, 0x9C, 0xC3, 0x8A}}

type ISpatialInteractionDetectedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_InteractionSourceKind() SpatialInteractionSourceKind
	TryGetPointerPose(coordinateSystem *ISpatialCoordinateSystem) *ISpatialPointerPose
	Get_Interaction() *ISpatialInteraction
}

type ISpatialInteractionDetectedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_InteractionSourceKind uintptr
	TryGetPointerPose         uintptr
	Get_Interaction           uintptr
}

type ISpatialInteractionDetectedEventArgs struct {
	win32.IInspectable
}

func (this *ISpatialInteractionDetectedEventArgs) Vtbl() *ISpatialInteractionDetectedEventArgsVtbl {
	return (*ISpatialInteractionDetectedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialInteractionDetectedEventArgs) Get_InteractionSourceKind() SpatialInteractionSourceKind {
	var _result SpatialInteractionSourceKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InteractionSourceKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionDetectedEventArgs) TryGetPointerPose(coordinateSystem *ISpatialCoordinateSystem) *ISpatialPointerPose {
	var _result *ISpatialPointerPose
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetPointerPose, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(coordinateSystem)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialInteractionDetectedEventArgs) Get_Interaction() *ISpatialInteraction {
	var _result *ISpatialInteraction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Interaction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7B263E93-5F13-419C-97D5-834678266AA6
var IID_ISpatialInteractionDetectedEventArgs2 = syscall.GUID{0x7B263E93, 0x5F13, 0x419C,
	[8]byte{0x97, 0xD5, 0x83, 0x46, 0x78, 0x26, 0x6A, 0xA6}}

type ISpatialInteractionDetectedEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_InteractionSource() *ISpatialInteractionSource
}

type ISpatialInteractionDetectedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_InteractionSource uintptr
}

type ISpatialInteractionDetectedEventArgs2 struct {
	win32.IInspectable
}

func (this *ISpatialInteractionDetectedEventArgs2) Vtbl() *ISpatialInteractionDetectedEventArgs2Vtbl {
	return (*ISpatialInteractionDetectedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialInteractionDetectedEventArgs2) Get_InteractionSource() *ISpatialInteractionSource {
	var _result *ISpatialInteractionSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InteractionSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 32A64EA8-A15A-3995-B8BD-80513CB5ADEF
var IID_ISpatialInteractionManager = syscall.GUID{0x32A64EA8, 0xA15A, 0x3995,
	[8]byte{0xB8, 0xBD, 0x80, 0x51, 0x3C, 0xB5, 0xAD, 0xEF}}

type ISpatialInteractionManagerInterface interface {
	win32.IInspectableInterface
	Add_SourceDetected(handler TypedEventHandler[*ISpatialInteractionManager, *ISpatialInteractionSourceEventArgs]) EventRegistrationToken
	Remove_SourceDetected(token EventRegistrationToken)
	Add_SourceLost(handler TypedEventHandler[*ISpatialInteractionManager, *ISpatialInteractionSourceEventArgs]) EventRegistrationToken
	Remove_SourceLost(token EventRegistrationToken)
	Add_SourceUpdated(handler TypedEventHandler[*ISpatialInteractionManager, *ISpatialInteractionSourceEventArgs]) EventRegistrationToken
	Remove_SourceUpdated(token EventRegistrationToken)
	Add_SourcePressed(handler TypedEventHandler[*ISpatialInteractionManager, *ISpatialInteractionSourceEventArgs]) EventRegistrationToken
	Remove_SourcePressed(token EventRegistrationToken)
	Add_SourceReleased(handler TypedEventHandler[*ISpatialInteractionManager, *ISpatialInteractionSourceEventArgs]) EventRegistrationToken
	Remove_SourceReleased(token EventRegistrationToken)
	Add_InteractionDetected(handler TypedEventHandler[*ISpatialInteractionManager, *ISpatialInteractionDetectedEventArgs]) EventRegistrationToken
	Remove_InteractionDetected(token EventRegistrationToken)
	GetDetectedSourcesAtTimestamp(timeStamp *IPerceptionTimestamp) *IVectorView[*ISpatialInteractionSourceState]
}

type ISpatialInteractionManagerVtbl struct {
	win32.IInspectableVtbl
	Add_SourceDetected            uintptr
	Remove_SourceDetected         uintptr
	Add_SourceLost                uintptr
	Remove_SourceLost             uintptr
	Add_SourceUpdated             uintptr
	Remove_SourceUpdated          uintptr
	Add_SourcePressed             uintptr
	Remove_SourcePressed          uintptr
	Add_SourceReleased            uintptr
	Remove_SourceReleased         uintptr
	Add_InteractionDetected       uintptr
	Remove_InteractionDetected    uintptr
	GetDetectedSourcesAtTimestamp uintptr
}

type ISpatialInteractionManager struct {
	win32.IInspectable
}

func (this *ISpatialInteractionManager) Vtbl() *ISpatialInteractionManagerVtbl {
	return (*ISpatialInteractionManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialInteractionManager) Add_SourceDetected(handler TypedEventHandler[*ISpatialInteractionManager, *ISpatialInteractionSourceEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SourceDetected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionManager) Remove_SourceDetected(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SourceDetected, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISpatialInteractionManager) Add_SourceLost(handler TypedEventHandler[*ISpatialInteractionManager, *ISpatialInteractionSourceEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SourceLost, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionManager) Remove_SourceLost(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SourceLost, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISpatialInteractionManager) Add_SourceUpdated(handler TypedEventHandler[*ISpatialInteractionManager, *ISpatialInteractionSourceEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SourceUpdated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionManager) Remove_SourceUpdated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SourceUpdated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISpatialInteractionManager) Add_SourcePressed(handler TypedEventHandler[*ISpatialInteractionManager, *ISpatialInteractionSourceEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SourcePressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionManager) Remove_SourcePressed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SourcePressed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISpatialInteractionManager) Add_SourceReleased(handler TypedEventHandler[*ISpatialInteractionManager, *ISpatialInteractionSourceEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SourceReleased, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionManager) Remove_SourceReleased(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SourceReleased, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISpatialInteractionManager) Add_InteractionDetected(handler TypedEventHandler[*ISpatialInteractionManager, *ISpatialInteractionDetectedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_InteractionDetected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionManager) Remove_InteractionDetected(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_InteractionDetected, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISpatialInteractionManager) GetDetectedSourcesAtTimestamp(timeStamp *IPerceptionTimestamp) *IVectorView[*ISpatialInteractionSourceState] {
	var _result *IVectorView[*ISpatialInteractionSourceState]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDetectedSourcesAtTimestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(timeStamp)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 00E31FA6-8CA2-30BF-91FE-D9CB4A008990
var IID_ISpatialInteractionManagerStatics = syscall.GUID{0x00E31FA6, 0x8CA2, 0x30BF,
	[8]byte{0x91, 0xFE, 0xD9, 0xCB, 0x4A, 0x00, 0x89, 0x90}}

type ISpatialInteractionManagerStaticsInterface interface {
	win32.IInspectableInterface
	GetForCurrentView() *ISpatialInteractionManager
}

type ISpatialInteractionManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForCurrentView uintptr
}

type ISpatialInteractionManagerStatics struct {
	win32.IInspectable
}

func (this *ISpatialInteractionManagerStatics) Vtbl() *ISpatialInteractionManagerStaticsVtbl {
	return (*ISpatialInteractionManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialInteractionManagerStatics) GetForCurrentView() *ISpatialInteractionManager {
	var _result *ISpatialInteractionManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 93F16C52-B88A-5929-8D7C-48CB948B081C
var IID_ISpatialInteractionManagerStatics2 = syscall.GUID{0x93F16C52, 0xB88A, 0x5929,
	[8]byte{0x8D, 0x7C, 0x48, 0xCB, 0x94, 0x8B, 0x08, 0x1C}}

type ISpatialInteractionManagerStatics2Interface interface {
	win32.IInspectableInterface
	IsSourceKindSupported(kind SpatialInteractionSourceKind) bool
}

type ISpatialInteractionManagerStatics2Vtbl struct {
	win32.IInspectableVtbl
	IsSourceKindSupported uintptr
}

type ISpatialInteractionManagerStatics2 struct {
	win32.IInspectable
}

func (this *ISpatialInteractionManagerStatics2) Vtbl() *ISpatialInteractionManagerStatics2Vtbl {
	return (*ISpatialInteractionManagerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialInteractionManagerStatics2) IsSourceKindSupported(kind SpatialInteractionSourceKind) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsSourceKindSupported, uintptr(unsafe.Pointer(this)), uintptr(kind), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// FB5433BA-B0B3-3148-9F3B-E9F5DE568F5D
var IID_ISpatialInteractionSource = syscall.GUID{0xFB5433BA, 0xB0B3, 0x3148,
	[8]byte{0x9F, 0x3B, 0xE9, 0xF5, 0xDE, 0x56, 0x8F, 0x5D}}

type ISpatialInteractionSourceInterface interface {
	win32.IInspectableInterface
	Get_Id() uint32
	Get_Kind() SpatialInteractionSourceKind
}

type ISpatialInteractionSourceVtbl struct {
	win32.IInspectableVtbl
	Get_Id   uintptr
	Get_Kind uintptr
}

type ISpatialInteractionSource struct {
	win32.IInspectable
}

func (this *ISpatialInteractionSource) Vtbl() *ISpatialInteractionSourceVtbl {
	return (*ISpatialInteractionSourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialInteractionSource) Get_Id() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionSource) Get_Kind() SpatialInteractionSourceKind {
	var _result SpatialInteractionSourceKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Kind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// E4C5B70C-0470-4028-88C0-A0EB44D34EFE
var IID_ISpatialInteractionSource2 = syscall.GUID{0xE4C5B70C, 0x0470, 0x4028,
	[8]byte{0x88, 0xC0, 0xA0, 0xEB, 0x44, 0xD3, 0x4E, 0xFE}}

type ISpatialInteractionSource2Interface interface {
	win32.IInspectableInterface
	Get_IsPointingSupported() bool
	Get_IsMenuSupported() bool
	Get_IsGraspSupported() bool
	Get_Controller() *ISpatialInteractionController
	TryGetStateAtTimestamp(timestamp *IPerceptionTimestamp) *ISpatialInteractionSourceState
}

type ISpatialInteractionSource2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsPointingSupported uintptr
	Get_IsMenuSupported     uintptr
	Get_IsGraspSupported    uintptr
	Get_Controller          uintptr
	TryGetStateAtTimestamp  uintptr
}

type ISpatialInteractionSource2 struct {
	win32.IInspectable
}

func (this *ISpatialInteractionSource2) Vtbl() *ISpatialInteractionSource2Vtbl {
	return (*ISpatialInteractionSource2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialInteractionSource2) Get_IsPointingSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPointingSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionSource2) Get_IsMenuSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsMenuSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionSource2) Get_IsGraspSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsGraspSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionSource2) Get_Controller() *ISpatialInteractionController {
	var _result *ISpatialInteractionController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Controller, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialInteractionSource2) TryGetStateAtTimestamp(timestamp *IPerceptionTimestamp) *ISpatialInteractionSourceState {
	var _result *ISpatialInteractionSourceState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetStateAtTimestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(timestamp)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0406D9F9-9AFD-44F9-85DC-700023A962E3
var IID_ISpatialInteractionSource3 = syscall.GUID{0x0406D9F9, 0x9AFD, 0x44F9,
	[8]byte{0x85, 0xDC, 0x70, 0x00, 0x23, 0xA9, 0x62, 0xE3}}

type ISpatialInteractionSource3Interface interface {
	win32.IInspectableInterface
	Get_Handedness() SpatialInteractionSourceHandedness
}

type ISpatialInteractionSource3Vtbl struct {
	win32.IInspectableVtbl
	Get_Handedness uintptr
}

type ISpatialInteractionSource3 struct {
	win32.IInspectable
}

func (this *ISpatialInteractionSource3) Vtbl() *ISpatialInteractionSource3Vtbl {
	return (*ISpatialInteractionSource3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialInteractionSource3) Get_Handedness() SpatialInteractionSourceHandedness {
	var _result SpatialInteractionSourceHandedness
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Handedness, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 0073BC4D-DF66-5A91-A2BA-CEA3E5C58A19
var IID_ISpatialInteractionSource4 = syscall.GUID{0x0073BC4D, 0xDF66, 0x5A91,
	[8]byte{0xA2, 0xBA, 0xCE, 0xA3, 0xE5, 0xC5, 0x8A, 0x19}}

type ISpatialInteractionSource4Interface interface {
	win32.IInspectableInterface
	TryCreateHandMeshObserver() unsafe.Pointer
	TryCreateHandMeshObserverAsync() *IAsyncOperation[unsafe.Pointer]
}

type ISpatialInteractionSource4Vtbl struct {
	win32.IInspectableVtbl
	TryCreateHandMeshObserver      uintptr
	TryCreateHandMeshObserverAsync uintptr
}

type ISpatialInteractionSource4 struct {
	win32.IInspectable
}

func (this *ISpatialInteractionSource4) Vtbl() *ISpatialInteractionSource4Vtbl {
	return (*ISpatialInteractionSource4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialInteractionSource4) TryCreateHandMeshObserver() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryCreateHandMeshObserver, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionSource4) TryCreateHandMeshObserverAsync() *IAsyncOperation[unsafe.Pointer] {
	var _result *IAsyncOperation[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryCreateHandMeshObserverAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 23B786CF-EC23-3979-B27C-EB0E12FEB7C7
var IID_ISpatialInteractionSourceEventArgs = syscall.GUID{0x23B786CF, 0xEC23, 0x3979,
	[8]byte{0xB2, 0x7C, 0xEB, 0x0E, 0x12, 0xFE, 0xB7, 0xC7}}

type ISpatialInteractionSourceEventArgsInterface interface {
	win32.IInspectableInterface
	Get_State() *ISpatialInteractionSourceState
}

type ISpatialInteractionSourceEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_State uintptr
}

type ISpatialInteractionSourceEventArgs struct {
	win32.IInspectable
}

func (this *ISpatialInteractionSourceEventArgs) Vtbl() *ISpatialInteractionSourceEventArgsVtbl {
	return (*ISpatialInteractionSourceEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialInteractionSourceEventArgs) Get_State() *ISpatialInteractionSourceState {
	var _result *ISpatialInteractionSourceState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D8B4B467-E648-4D52-AB49-E0D227199F63
var IID_ISpatialInteractionSourceEventArgs2 = syscall.GUID{0xD8B4B467, 0xE648, 0x4D52,
	[8]byte{0xAB, 0x49, 0xE0, 0xD2, 0x27, 0x19, 0x9F, 0x63}}

type ISpatialInteractionSourceEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_PressKind() SpatialInteractionPressKind
}

type ISpatialInteractionSourceEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_PressKind uintptr
}

type ISpatialInteractionSourceEventArgs2 struct {
	win32.IInspectable
}

func (this *ISpatialInteractionSourceEventArgs2) Vtbl() *ISpatialInteractionSourceEventArgs2Vtbl {
	return (*ISpatialInteractionSourceEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialInteractionSourceEventArgs2) Get_PressKind() SpatialInteractionPressKind {
	var _result SpatialInteractionPressKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PressKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// EA4696C4-7E8B-30CA-BCC5-C77189CEA30A
var IID_ISpatialInteractionSourceLocation = syscall.GUID{0xEA4696C4, 0x7E8B, 0x30CA,
	[8]byte{0xBC, 0xC5, 0xC7, 0x71, 0x89, 0xCE, 0xA3, 0x0A}}

type ISpatialInteractionSourceLocationInterface interface {
	win32.IInspectableInterface
	Get_Position() *IReference[unsafe.Pointer]
	Get_Velocity() *IReference[unsafe.Pointer]
}

type ISpatialInteractionSourceLocationVtbl struct {
	win32.IInspectableVtbl
	Get_Position uintptr
	Get_Velocity uintptr
}

type ISpatialInteractionSourceLocation struct {
	win32.IInspectable
}

func (this *ISpatialInteractionSourceLocation) Vtbl() *ISpatialInteractionSourceLocationVtbl {
	return (*ISpatialInteractionSourceLocationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialInteractionSourceLocation) Get_Position() *IReference[unsafe.Pointer] {
	var _result *IReference[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialInteractionSourceLocation) Get_Velocity() *IReference[unsafe.Pointer] {
	var _result *IReference[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Velocity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4C671045-3917-40FC-A9AC-31C9CF5FF91B
var IID_ISpatialInteractionSourceLocation2 = syscall.GUID{0x4C671045, 0x3917, 0x40FC,
	[8]byte{0xA9, 0xAC, 0x31, 0xC9, 0xCF, 0x5F, 0xF9, 0x1B}}

type ISpatialInteractionSourceLocation2Interface interface {
	win32.IInspectableInterface
	Get_Orientation() *IReference[unsafe.Pointer]
}

type ISpatialInteractionSourceLocation2Vtbl struct {
	win32.IInspectableVtbl
	Get_Orientation uintptr
}

type ISpatialInteractionSourceLocation2 struct {
	win32.IInspectable
}

func (this *ISpatialInteractionSourceLocation2) Vtbl() *ISpatialInteractionSourceLocation2Vtbl {
	return (*ISpatialInteractionSourceLocation2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialInteractionSourceLocation2) Get_Orientation() *IReference[unsafe.Pointer] {
	var _result *IReference[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Orientation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6702E65E-E915-4CFB-9C1B-0538EFC86687
var IID_ISpatialInteractionSourceLocation3 = syscall.GUID{0x6702E65E, 0xE915, 0x4CFB,
	[8]byte{0x9C, 0x1B, 0x05, 0x38, 0xEF, 0xC8, 0x66, 0x87}}

type ISpatialInteractionSourceLocation3Interface interface {
	win32.IInspectableInterface
	Get_PositionAccuracy() SpatialInteractionSourcePositionAccuracy
	Get_AngularVelocity() *IReference[unsafe.Pointer]
	Get_SourcePointerPose() *ISpatialPointerInteractionSourcePose
}

type ISpatialInteractionSourceLocation3Vtbl struct {
	win32.IInspectableVtbl
	Get_PositionAccuracy  uintptr
	Get_AngularVelocity   uintptr
	Get_SourcePointerPose uintptr
}

type ISpatialInteractionSourceLocation3 struct {
	win32.IInspectable
}

func (this *ISpatialInteractionSourceLocation3) Vtbl() *ISpatialInteractionSourceLocation3Vtbl {
	return (*ISpatialInteractionSourceLocation3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialInteractionSourceLocation3) Get_PositionAccuracy() SpatialInteractionSourcePositionAccuracy {
	var _result SpatialInteractionSourcePositionAccuracy
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PositionAccuracy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionSourceLocation3) Get_AngularVelocity() *IReference[unsafe.Pointer] {
	var _result *IReference[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AngularVelocity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialInteractionSourceLocation3) Get_SourcePointerPose() *ISpatialPointerInteractionSourcePose {
	var _result *ISpatialPointerInteractionSourcePose
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SourcePointerPose, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 05604542-3EF7-3222-9F53-63C9CB7E3BC7
var IID_ISpatialInteractionSourceProperties = syscall.GUID{0x05604542, 0x3EF7, 0x3222,
	[8]byte{0x9F, 0x53, 0x63, 0xC9, 0xCB, 0x7E, 0x3B, 0xC7}}

type ISpatialInteractionSourcePropertiesInterface interface {
	win32.IInspectableInterface
	TryGetSourceLossMitigationDirection(coordinateSystem *ISpatialCoordinateSystem) *IReference[unsafe.Pointer]
	Get_SourceLossRisk() float64
	TryGetLocation(coordinateSystem *ISpatialCoordinateSystem) *ISpatialInteractionSourceLocation
}

type ISpatialInteractionSourcePropertiesVtbl struct {
	win32.IInspectableVtbl
	TryGetSourceLossMitigationDirection uintptr
	Get_SourceLossRisk                  uintptr
	TryGetLocation                      uintptr
}

type ISpatialInteractionSourceProperties struct {
	win32.IInspectable
}

func (this *ISpatialInteractionSourceProperties) Vtbl() *ISpatialInteractionSourcePropertiesVtbl {
	return (*ISpatialInteractionSourcePropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialInteractionSourceProperties) TryGetSourceLossMitigationDirection(coordinateSystem *ISpatialCoordinateSystem) *IReference[unsafe.Pointer] {
	var _result *IReference[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetSourceLossMitigationDirection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(coordinateSystem)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialInteractionSourceProperties) Get_SourceLossRisk() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SourceLossRisk, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionSourceProperties) TryGetLocation(coordinateSystem *ISpatialCoordinateSystem) *ISpatialInteractionSourceLocation {
	var _result *ISpatialInteractionSourceLocation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetLocation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(coordinateSystem)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D5C475EF-4B63-37EC-98B9-9FC652B9D2F2
var IID_ISpatialInteractionSourceState = syscall.GUID{0xD5C475EF, 0x4B63, 0x37EC,
	[8]byte{0x98, 0xB9, 0x9F, 0xC6, 0x52, 0xB9, 0xD2, 0xF2}}

type ISpatialInteractionSourceStateInterface interface {
	win32.IInspectableInterface
	Get_Source() *ISpatialInteractionSource
	Get_Properties() *ISpatialInteractionSourceProperties
	Get_IsPressed() bool
	Get_Timestamp() *IPerceptionTimestamp
	TryGetPointerPose(coordinateSystem *ISpatialCoordinateSystem) *ISpatialPointerPose
}

type ISpatialInteractionSourceStateVtbl struct {
	win32.IInspectableVtbl
	Get_Source        uintptr
	Get_Properties    uintptr
	Get_IsPressed     uintptr
	Get_Timestamp     uintptr
	TryGetPointerPose uintptr
}

type ISpatialInteractionSourceState struct {
	win32.IInspectable
}

func (this *ISpatialInteractionSourceState) Vtbl() *ISpatialInteractionSourceStateVtbl {
	return (*ISpatialInteractionSourceStateVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialInteractionSourceState) Get_Source() *ISpatialInteractionSource {
	var _result *ISpatialInteractionSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Source, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialInteractionSourceState) Get_Properties() *ISpatialInteractionSourceProperties {
	var _result *ISpatialInteractionSourceProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialInteractionSourceState) Get_IsPressed() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionSourceState) Get_Timestamp() *IPerceptionTimestamp {
	var _result *IPerceptionTimestamp
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialInteractionSourceState) TryGetPointerPose(coordinateSystem *ISpatialCoordinateSystem) *ISpatialPointerPose {
	var _result *ISpatialPointerPose
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetPointerPose, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(coordinateSystem)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 45F6D0BD-1773-492E-9BA3-8AC1CBE77C08
var IID_ISpatialInteractionSourceState2 = syscall.GUID{0x45F6D0BD, 0x1773, 0x492E,
	[8]byte{0x9B, 0xA3, 0x8A, 0xC1, 0xCB, 0xE7, 0x7C, 0x08}}

type ISpatialInteractionSourceState2Interface interface {
	win32.IInspectableInterface
	Get_IsSelectPressed() bool
	Get_IsMenuPressed() bool
	Get_IsGrasped() bool
	Get_SelectPressedValue() float64
	Get_ControllerProperties() *ISpatialInteractionControllerProperties
}

type ISpatialInteractionSourceState2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsSelectPressed      uintptr
	Get_IsMenuPressed        uintptr
	Get_IsGrasped            uintptr
	Get_SelectPressedValue   uintptr
	Get_ControllerProperties uintptr
}

type ISpatialInteractionSourceState2 struct {
	win32.IInspectable
}

func (this *ISpatialInteractionSourceState2) Vtbl() *ISpatialInteractionSourceState2Vtbl {
	return (*ISpatialInteractionSourceState2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialInteractionSourceState2) Get_IsSelectPressed() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSelectPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionSourceState2) Get_IsMenuPressed() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsMenuPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionSourceState2) Get_IsGrasped() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsGrasped, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionSourceState2) Get_SelectPressedValue() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SelectPressedValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialInteractionSourceState2) Get_ControllerProperties() *ISpatialInteractionControllerProperties {
	var _result *ISpatialInteractionControllerProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ControllerProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F2F00BC2-BD2B-4A01-A8FB-323E0158527C
var IID_ISpatialInteractionSourceState3 = syscall.GUID{0xF2F00BC2, 0xBD2B, 0x4A01,
	[8]byte{0xA8, 0xFB, 0x32, 0x3E, 0x01, 0x58, 0x52, 0x7C}}

type ISpatialInteractionSourceState3Interface interface {
	win32.IInspectableInterface
	TryGetHandPose() unsafe.Pointer
}

type ISpatialInteractionSourceState3Vtbl struct {
	win32.IInspectableVtbl
	TryGetHandPose uintptr
}

type ISpatialInteractionSourceState3 struct {
	win32.IInspectable
}

func (this *ISpatialInteractionSourceState3) Vtbl() *ISpatialInteractionSourceState3Vtbl {
	return (*ISpatialInteractionSourceState3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialInteractionSourceState3) TryGetHandPose() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetHandPose, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 2D40D1CB-E7DA-4220-B0BF-819301674780
var IID_ISpatialManipulationCanceledEventArgs = syscall.GUID{0x2D40D1CB, 0xE7DA, 0x4220,
	[8]byte{0xB0, 0xBF, 0x81, 0x93, 0x01, 0x67, 0x47, 0x80}}

type ISpatialManipulationCanceledEventArgsInterface interface {
	win32.IInspectableInterface
	Get_InteractionSourceKind() SpatialInteractionSourceKind
}

type ISpatialManipulationCanceledEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_InteractionSourceKind uintptr
}

type ISpatialManipulationCanceledEventArgs struct {
	win32.IInspectable
}

func (this *ISpatialManipulationCanceledEventArgs) Vtbl() *ISpatialManipulationCanceledEventArgsVtbl {
	return (*ISpatialManipulationCanceledEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialManipulationCanceledEventArgs) Get_InteractionSourceKind() SpatialInteractionSourceKind {
	var _result SpatialInteractionSourceKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InteractionSourceKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 05086802-F301-4343-9250-2FBAA5F87A37
var IID_ISpatialManipulationCompletedEventArgs = syscall.GUID{0x05086802, 0xF301, 0x4343,
	[8]byte{0x92, 0x50, 0x2F, 0xBA, 0xA5, 0xF8, 0x7A, 0x37}}

type ISpatialManipulationCompletedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_InteractionSourceKind() SpatialInteractionSourceKind
	TryGetCumulativeDelta(coordinateSystem *ISpatialCoordinateSystem) *ISpatialManipulationDelta
}

type ISpatialManipulationCompletedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_InteractionSourceKind uintptr
	TryGetCumulativeDelta     uintptr
}

type ISpatialManipulationCompletedEventArgs struct {
	win32.IInspectable
}

func (this *ISpatialManipulationCompletedEventArgs) Vtbl() *ISpatialManipulationCompletedEventArgsVtbl {
	return (*ISpatialManipulationCompletedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialManipulationCompletedEventArgs) Get_InteractionSourceKind() SpatialInteractionSourceKind {
	var _result SpatialInteractionSourceKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InteractionSourceKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialManipulationCompletedEventArgs) TryGetCumulativeDelta(coordinateSystem *ISpatialCoordinateSystem) *ISpatialManipulationDelta {
	var _result *ISpatialManipulationDelta
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetCumulativeDelta, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(coordinateSystem)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A7EC967A-D123-3A81-A15B-992923DCBE91
var IID_ISpatialManipulationDelta = syscall.GUID{0xA7EC967A, 0xD123, 0x3A81,
	[8]byte{0xA1, 0x5B, 0x99, 0x29, 0x23, 0xDC, 0xBE, 0x91}}

type ISpatialManipulationDeltaInterface interface {
	win32.IInspectableInterface
	Get_Translation() unsafe.Pointer
}

type ISpatialManipulationDeltaVtbl struct {
	win32.IInspectableVtbl
	Get_Translation uintptr
}

type ISpatialManipulationDelta struct {
	win32.IInspectable
}

func (this *ISpatialManipulationDelta) Vtbl() *ISpatialManipulationDeltaVtbl {
	return (*ISpatialManipulationDeltaVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialManipulationDelta) Get_Translation() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Translation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A1D6BBCE-42A5-377B-ADA6-D28E3D384737
var IID_ISpatialManipulationStartedEventArgs = syscall.GUID{0xA1D6BBCE, 0x42A5, 0x377B,
	[8]byte{0xAD, 0xA6, 0xD2, 0x8E, 0x3D, 0x38, 0x47, 0x37}}

type ISpatialManipulationStartedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_InteractionSourceKind() SpatialInteractionSourceKind
	TryGetPointerPose(coordinateSystem *ISpatialCoordinateSystem) *ISpatialPointerPose
}

type ISpatialManipulationStartedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_InteractionSourceKind uintptr
	TryGetPointerPose         uintptr
}

type ISpatialManipulationStartedEventArgs struct {
	win32.IInspectable
}

func (this *ISpatialManipulationStartedEventArgs) Vtbl() *ISpatialManipulationStartedEventArgsVtbl {
	return (*ISpatialManipulationStartedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialManipulationStartedEventArgs) Get_InteractionSourceKind() SpatialInteractionSourceKind {
	var _result SpatialInteractionSourceKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InteractionSourceKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialManipulationStartedEventArgs) TryGetPointerPose(coordinateSystem *ISpatialCoordinateSystem) *ISpatialPointerPose {
	var _result *ISpatialPointerPose
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetPointerPose, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(coordinateSystem)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5F230B9B-60C6-4DC6-BDC9-9F4A6F15FE49
var IID_ISpatialManipulationUpdatedEventArgs = syscall.GUID{0x5F230B9B, 0x60C6, 0x4DC6,
	[8]byte{0xBD, 0xC9, 0x9F, 0x4A, 0x6F, 0x15, 0xFE, 0x49}}

type ISpatialManipulationUpdatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_InteractionSourceKind() SpatialInteractionSourceKind
	TryGetCumulativeDelta(coordinateSystem *ISpatialCoordinateSystem) *ISpatialManipulationDelta
}

type ISpatialManipulationUpdatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_InteractionSourceKind uintptr
	TryGetCumulativeDelta     uintptr
}

type ISpatialManipulationUpdatedEventArgs struct {
	win32.IInspectable
}

func (this *ISpatialManipulationUpdatedEventArgs) Vtbl() *ISpatialManipulationUpdatedEventArgsVtbl {
	return (*ISpatialManipulationUpdatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialManipulationUpdatedEventArgs) Get_InteractionSourceKind() SpatialInteractionSourceKind {
	var _result SpatialInteractionSourceKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InteractionSourceKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialManipulationUpdatedEventArgs) TryGetCumulativeDelta(coordinateSystem *ISpatialCoordinateSystem) *ISpatialManipulationDelta {
	var _result *ISpatialManipulationDelta
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetCumulativeDelta, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(coordinateSystem)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CE503EDC-E8A5-46F0-92D4-3C122B35112A
var IID_ISpatialNavigationCanceledEventArgs = syscall.GUID{0xCE503EDC, 0xE8A5, 0x46F0,
	[8]byte{0x92, 0xD4, 0x3C, 0x12, 0x2B, 0x35, 0x11, 0x2A}}

type ISpatialNavigationCanceledEventArgsInterface interface {
	win32.IInspectableInterface
	Get_InteractionSourceKind() SpatialInteractionSourceKind
}

type ISpatialNavigationCanceledEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_InteractionSourceKind uintptr
}

type ISpatialNavigationCanceledEventArgs struct {
	win32.IInspectable
}

func (this *ISpatialNavigationCanceledEventArgs) Vtbl() *ISpatialNavigationCanceledEventArgsVtbl {
	return (*ISpatialNavigationCanceledEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialNavigationCanceledEventArgs) Get_InteractionSourceKind() SpatialInteractionSourceKind {
	var _result SpatialInteractionSourceKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InteractionSourceKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 012E80B7-AF3B-42C2-9E41-BAAA0E721F3A
var IID_ISpatialNavigationCompletedEventArgs = syscall.GUID{0x012E80B7, 0xAF3B, 0x42C2,
	[8]byte{0x9E, 0x41, 0xBA, 0xAA, 0x0E, 0x72, 0x1F, 0x3A}}

type ISpatialNavigationCompletedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_InteractionSourceKind() SpatialInteractionSourceKind
	Get_NormalizedOffset() unsafe.Pointer
}

type ISpatialNavigationCompletedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_InteractionSourceKind uintptr
	Get_NormalizedOffset      uintptr
}

type ISpatialNavigationCompletedEventArgs struct {
	win32.IInspectable
}

func (this *ISpatialNavigationCompletedEventArgs) Vtbl() *ISpatialNavigationCompletedEventArgsVtbl {
	return (*ISpatialNavigationCompletedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialNavigationCompletedEventArgs) Get_InteractionSourceKind() SpatialInteractionSourceKind {
	var _result SpatialInteractionSourceKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InteractionSourceKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialNavigationCompletedEventArgs) Get_NormalizedOffset() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NormalizedOffset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 754A348A-FB64-4656-8EBD-9DEECAAFE475
var IID_ISpatialNavigationStartedEventArgs = syscall.GUID{0x754A348A, 0xFB64, 0x4656,
	[8]byte{0x8E, 0xBD, 0x9D, 0xEE, 0xCA, 0xAF, 0xE4, 0x75}}

type ISpatialNavigationStartedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_InteractionSourceKind() SpatialInteractionSourceKind
	TryGetPointerPose(coordinateSystem *ISpatialCoordinateSystem) *ISpatialPointerPose
	Get_IsNavigatingX() bool
	Get_IsNavigatingY() bool
	Get_IsNavigatingZ() bool
}

type ISpatialNavigationStartedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_InteractionSourceKind uintptr
	TryGetPointerPose         uintptr
	Get_IsNavigatingX         uintptr
	Get_IsNavigatingY         uintptr
	Get_IsNavigatingZ         uintptr
}

type ISpatialNavigationStartedEventArgs struct {
	win32.IInspectable
}

func (this *ISpatialNavigationStartedEventArgs) Vtbl() *ISpatialNavigationStartedEventArgsVtbl {
	return (*ISpatialNavigationStartedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialNavigationStartedEventArgs) Get_InteractionSourceKind() SpatialInteractionSourceKind {
	var _result SpatialInteractionSourceKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InteractionSourceKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialNavigationStartedEventArgs) TryGetPointerPose(coordinateSystem *ISpatialCoordinateSystem) *ISpatialPointerPose {
	var _result *ISpatialPointerPose
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetPointerPose, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(coordinateSystem)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialNavigationStartedEventArgs) Get_IsNavigatingX() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsNavigatingX, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialNavigationStartedEventArgs) Get_IsNavigatingY() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsNavigatingY, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialNavigationStartedEventArgs) Get_IsNavigatingZ() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsNavigatingZ, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 9B713FD7-839D-4A74-8732-45466FC044B5
var IID_ISpatialNavigationUpdatedEventArgs = syscall.GUID{0x9B713FD7, 0x839D, 0x4A74,
	[8]byte{0x87, 0x32, 0x45, 0x46, 0x6F, 0xC0, 0x44, 0xB5}}

type ISpatialNavigationUpdatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_InteractionSourceKind() SpatialInteractionSourceKind
	Get_NormalizedOffset() unsafe.Pointer
}

type ISpatialNavigationUpdatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_InteractionSourceKind uintptr
	Get_NormalizedOffset      uintptr
}

type ISpatialNavigationUpdatedEventArgs struct {
	win32.IInspectable
}

func (this *ISpatialNavigationUpdatedEventArgs) Vtbl() *ISpatialNavigationUpdatedEventArgsVtbl {
	return (*ISpatialNavigationUpdatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialNavigationUpdatedEventArgs) Get_InteractionSourceKind() SpatialInteractionSourceKind {
	var _result SpatialInteractionSourceKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InteractionSourceKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialNavigationUpdatedEventArgs) Get_NormalizedOffset() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NormalizedOffset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A7104307-2C2B-4D3A-92A7-80CED7C4A0D0
var IID_ISpatialPointerInteractionSourcePose = syscall.GUID{0xA7104307, 0x2C2B, 0x4D3A,
	[8]byte{0x92, 0xA7, 0x80, 0xCE, 0xD7, 0xC4, 0xA0, 0xD0}}

type ISpatialPointerInteractionSourcePoseInterface interface {
	win32.IInspectableInterface
	Get_Position() unsafe.Pointer
	Get_ForwardDirection() unsafe.Pointer
	Get_UpDirection() unsafe.Pointer
}

type ISpatialPointerInteractionSourcePoseVtbl struct {
	win32.IInspectableVtbl
	Get_Position         uintptr
	Get_ForwardDirection uintptr
	Get_UpDirection      uintptr
}

type ISpatialPointerInteractionSourcePose struct {
	win32.IInspectable
}

func (this *ISpatialPointerInteractionSourcePose) Vtbl() *ISpatialPointerInteractionSourcePoseVtbl {
	return (*ISpatialPointerInteractionSourcePoseVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialPointerInteractionSourcePose) Get_Position() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialPointerInteractionSourcePose) Get_ForwardDirection() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ForwardDirection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialPointerInteractionSourcePose) Get_UpDirection() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UpDirection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// ECCD86B8-52DB-469F-9E3F-80C47F74BCE9
var IID_ISpatialPointerInteractionSourcePose2 = syscall.GUID{0xECCD86B8, 0x52DB, 0x469F,
	[8]byte{0x9E, 0x3F, 0x80, 0xC4, 0x7F, 0x74, 0xBC, 0xE9}}

type ISpatialPointerInteractionSourcePose2Interface interface {
	win32.IInspectableInterface
	Get_Orientation() unsafe.Pointer
	Get_PositionAccuracy() SpatialInteractionSourcePositionAccuracy
}

type ISpatialPointerInteractionSourcePose2Vtbl struct {
	win32.IInspectableVtbl
	Get_Orientation      uintptr
	Get_PositionAccuracy uintptr
}

type ISpatialPointerInteractionSourcePose2 struct {
	win32.IInspectable
}

func (this *ISpatialPointerInteractionSourcePose2) Vtbl() *ISpatialPointerInteractionSourcePose2Vtbl {
	return (*ISpatialPointerInteractionSourcePose2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialPointerInteractionSourcePose2) Get_Orientation() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Orientation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialPointerInteractionSourcePose2) Get_PositionAccuracy() SpatialInteractionSourcePositionAccuracy {
	var _result SpatialInteractionSourcePositionAccuracy
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PositionAccuracy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 6953A42E-C17E-357D-97A1-7269D0ED2D10
var IID_ISpatialPointerPose = syscall.GUID{0x6953A42E, 0xC17E, 0x357D,
	[8]byte{0x97, 0xA1, 0x72, 0x69, 0xD0, 0xED, 0x2D, 0x10}}

type ISpatialPointerPoseInterface interface {
	win32.IInspectableInterface
	Get_Timestamp() *IPerceptionTimestamp
	Get_Head() unsafe.Pointer
}

type ISpatialPointerPoseVtbl struct {
	win32.IInspectableVtbl
	Get_Timestamp uintptr
	Get_Head      uintptr
}

type ISpatialPointerPose struct {
	win32.IInspectable
}

func (this *ISpatialPointerPose) Vtbl() *ISpatialPointerPoseVtbl {
	return (*ISpatialPointerPoseVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialPointerPose) Get_Timestamp() *IPerceptionTimestamp {
	var _result *IPerceptionTimestamp
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialPointerPose) Get_Head() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Head, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 9D202B17-954E-4E0C-96D1-B6790B6FC2FD
var IID_ISpatialPointerPose2 = syscall.GUID{0x9D202B17, 0x954E, 0x4E0C,
	[8]byte{0x96, 0xD1, 0xB6, 0x79, 0x0B, 0x6F, 0xC2, 0xFD}}

type ISpatialPointerPose2Interface interface {
	win32.IInspectableInterface
	TryGetInteractionSourcePose(source *ISpatialInteractionSource) *ISpatialPointerInteractionSourcePose
}

type ISpatialPointerPose2Vtbl struct {
	win32.IInspectableVtbl
	TryGetInteractionSourcePose uintptr
}

type ISpatialPointerPose2 struct {
	win32.IInspectable
}

func (this *ISpatialPointerPose2) Vtbl() *ISpatialPointerPose2Vtbl {
	return (*ISpatialPointerPose2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialPointerPose2) TryGetInteractionSourcePose(source *ISpatialInteractionSource) *ISpatialPointerInteractionSourcePose {
	var _result *ISpatialPointerInteractionSourcePose
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetInteractionSourcePose, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(source)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6342F3F0-EC49-5B4B-B8D1-D16CBB16BE84
var IID_ISpatialPointerPose3 = syscall.GUID{0x6342F3F0, 0xEC49, 0x5B4B,
	[8]byte{0xB8, 0xD1, 0xD1, 0x6C, 0xBB, 0x16, 0xBE, 0x84}}

type ISpatialPointerPose3Interface interface {
	win32.IInspectableInterface
	Get_Eyes() unsafe.Pointer
	Get_IsHeadCapturedBySystem() bool
}

type ISpatialPointerPose3Vtbl struct {
	win32.IInspectableVtbl
	Get_Eyes                   uintptr
	Get_IsHeadCapturedBySystem uintptr
}

type ISpatialPointerPose3 struct {
	win32.IInspectable
}

func (this *ISpatialPointerPose3) Vtbl() *ISpatialPointerPose3Vtbl {
	return (*ISpatialPointerPose3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialPointerPose3) Get_Eyes() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Eyes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialPointerPose3) Get_IsHeadCapturedBySystem() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsHeadCapturedBySystem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A25591A9-ACA1-3EE0-9816-785CFB2E3FB8
var IID_ISpatialPointerPoseStatics = syscall.GUID{0xA25591A9, 0xACA1, 0x3EE0,
	[8]byte{0x98, 0x16, 0x78, 0x5C, 0xFB, 0x2E, 0x3F, 0xB8}}

type ISpatialPointerPoseStaticsInterface interface {
	win32.IInspectableInterface
	TryGetAtTimestamp(coordinateSystem *ISpatialCoordinateSystem, timestamp *IPerceptionTimestamp) *ISpatialPointerPose
}

type ISpatialPointerPoseStaticsVtbl struct {
	win32.IInspectableVtbl
	TryGetAtTimestamp uintptr
}

type ISpatialPointerPoseStatics struct {
	win32.IInspectable
}

func (this *ISpatialPointerPoseStatics) Vtbl() *ISpatialPointerPoseStaticsVtbl {
	return (*ISpatialPointerPoseStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialPointerPoseStatics) TryGetAtTimestamp(coordinateSystem *ISpatialCoordinateSystem, timestamp *IPerceptionTimestamp) *ISpatialPointerPose {
	var _result *ISpatialPointerPose
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetAtTimestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(coordinateSystem)), uintptr(unsafe.Pointer(timestamp)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0E35F5CB-3F75-43F3-AC81-D1DC2DF9B1FB
var IID_ISpatialRecognitionEndedEventArgs = syscall.GUID{0x0E35F5CB, 0x3F75, 0x43F3,
	[8]byte{0xAC, 0x81, 0xD1, 0xDC, 0x2D, 0xF9, 0xB1, 0xFB}}

type ISpatialRecognitionEndedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_InteractionSourceKind() SpatialInteractionSourceKind
}

type ISpatialRecognitionEndedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_InteractionSourceKind uintptr
}

type ISpatialRecognitionEndedEventArgs struct {
	win32.IInspectable
}

func (this *ISpatialRecognitionEndedEventArgs) Vtbl() *ISpatialRecognitionEndedEventArgsVtbl {
	return (*ISpatialRecognitionEndedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialRecognitionEndedEventArgs) Get_InteractionSourceKind() SpatialInteractionSourceKind {
	var _result SpatialInteractionSourceKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InteractionSourceKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 24DA128F-0008-4A6D-AA50-2A76F9CFB264
var IID_ISpatialRecognitionStartedEventArgs = syscall.GUID{0x24DA128F, 0x0008, 0x4A6D,
	[8]byte{0xAA, 0x50, 0x2A, 0x76, 0xF9, 0xCF, 0xB2, 0x64}}

type ISpatialRecognitionStartedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_InteractionSourceKind() SpatialInteractionSourceKind
	TryGetPointerPose(coordinateSystem *ISpatialCoordinateSystem) *ISpatialPointerPose
	IsGesturePossible(gesture SpatialGestureSettings) bool
}

type ISpatialRecognitionStartedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_InteractionSourceKind uintptr
	TryGetPointerPose         uintptr
	IsGesturePossible         uintptr
}

type ISpatialRecognitionStartedEventArgs struct {
	win32.IInspectable
}

func (this *ISpatialRecognitionStartedEventArgs) Vtbl() *ISpatialRecognitionStartedEventArgsVtbl {
	return (*ISpatialRecognitionStartedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialRecognitionStartedEventArgs) Get_InteractionSourceKind() SpatialInteractionSourceKind {
	var _result SpatialInteractionSourceKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InteractionSourceKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialRecognitionStartedEventArgs) TryGetPointerPose(coordinateSystem *ISpatialCoordinateSystem) *ISpatialPointerPose {
	var _result *ISpatialPointerPose
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetPointerPose, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(coordinateSystem)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialRecognitionStartedEventArgs) IsGesturePossible(gesture SpatialGestureSettings) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsGesturePossible, uintptr(unsafe.Pointer(this)), uintptr(gesture), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 296D83DE-F444-4AA1-B2BF-9DC88D567DA6
var IID_ISpatialTappedEventArgs = syscall.GUID{0x296D83DE, 0xF444, 0x4AA1,
	[8]byte{0xB2, 0xBF, 0x9D, 0xC8, 0x8D, 0x56, 0x7D, 0xA6}}

type ISpatialTappedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_InteractionSourceKind() SpatialInteractionSourceKind
	TryGetPointerPose(coordinateSystem *ISpatialCoordinateSystem) *ISpatialPointerPose
	Get_TapCount() uint32
}

type ISpatialTappedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_InteractionSourceKind uintptr
	TryGetPointerPose         uintptr
	Get_TapCount              uintptr
}

type ISpatialTappedEventArgs struct {
	win32.IInspectable
}

func (this *ISpatialTappedEventArgs) Vtbl() *ISpatialTappedEventArgsVtbl {
	return (*ISpatialTappedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialTappedEventArgs) Get_InteractionSourceKind() SpatialInteractionSourceKind {
	var _result SpatialInteractionSourceKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InteractionSourceKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialTappedEventArgs) TryGetPointerPose(coordinateSystem *ISpatialCoordinateSystem) *ISpatialPointerPose {
	var _result *ISpatialPointerPose
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetPointerPose, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(coordinateSystem)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialTappedEventArgs) Get_TapCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TapCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type SpatialGestureRecognizer struct {
	RtClass
	*ISpatialGestureRecognizer
}

func NewSpatialGestureRecognizer_Create(settings SpatialGestureSettings) *SpatialGestureRecognizer {
	hs := NewHStr("Windows.UI.Input.Spatial.SpatialGestureRecognizer")
	var pFac *ISpatialGestureRecognizerFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISpatialGestureRecognizerFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISpatialGestureRecognizer
	p = pFac.Create(settings)
	result := &SpatialGestureRecognizer{
		RtClass:                   RtClass{PInspect: &p.IInspectable},
		ISpatialGestureRecognizer: p,
	}
	com.AddToScope(result)
	return result
}
