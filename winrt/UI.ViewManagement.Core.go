package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type CoreInputViewKind int32

const (
	CoreInputViewKind_Default     CoreInputViewKind = 0
	CoreInputViewKind_Keyboard    CoreInputViewKind = 1
	CoreInputViewKind_Handwriting CoreInputViewKind = 2
	CoreInputViewKind_Emoji       CoreInputViewKind = 3
	CoreInputViewKind_Symbols     CoreInputViewKind = 4
	CoreInputViewKind_Clipboard   CoreInputViewKind = 5
	CoreInputViewKind_Dictation   CoreInputViewKind = 6
)

// enum
type CoreInputViewOcclusionKind int32

const (
	CoreInputViewOcclusionKind_Docked   CoreInputViewOcclusionKind = 0
	CoreInputViewOcclusionKind_Floating CoreInputViewOcclusionKind = 1
	CoreInputViewOcclusionKind_Overlay  CoreInputViewOcclusionKind = 2
)

// enum
type CoreInputViewXYFocusTransferDirection int32

const (
	CoreInputViewXYFocusTransferDirection_Up    CoreInputViewXYFocusTransferDirection = 0
	CoreInputViewXYFocusTransferDirection_Right CoreInputViewXYFocusTransferDirection = 1
	CoreInputViewXYFocusTransferDirection_Down  CoreInputViewXYFocusTransferDirection = 2
	CoreInputViewXYFocusTransferDirection_Left  CoreInputViewXYFocusTransferDirection = 3
)

// interfaces

// D77C94AE-46B8-5D4A-9489-8DDEC3D639A6
var IID_ICoreFrameworkInputView = syscall.GUID{0xD77C94AE, 0x46B8, 0x5D4A,
	[8]byte{0x94, 0x89, 0x8D, 0xDE, 0xC3, 0xD6, 0x39, 0xA6}}

type ICoreFrameworkInputViewInterface interface {
	win32.IInspectableInterface
	Add_PrimaryViewAnimationStarting(handler TypedEventHandler[*ICoreFrameworkInputView, *ICoreFrameworkInputViewAnimationStartingEventArgs]) EventRegistrationToken
	Remove_PrimaryViewAnimationStarting(token EventRegistrationToken)
	Add_OcclusionsChanged(handler TypedEventHandler[*ICoreFrameworkInputView, *ICoreFrameworkInputViewOcclusionsChangedEventArgs]) EventRegistrationToken
	Remove_OcclusionsChanged(token EventRegistrationToken)
}

type ICoreFrameworkInputViewVtbl struct {
	win32.IInspectableVtbl
	Add_PrimaryViewAnimationStarting    uintptr
	Remove_PrimaryViewAnimationStarting uintptr
	Add_OcclusionsChanged               uintptr
	Remove_OcclusionsChanged            uintptr
}

type ICoreFrameworkInputView struct {
	win32.IInspectable
}

func (this *ICoreFrameworkInputView) Vtbl() *ICoreFrameworkInputViewVtbl {
	return (*ICoreFrameworkInputViewVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreFrameworkInputView) Add_PrimaryViewAnimationStarting(handler TypedEventHandler[*ICoreFrameworkInputView, *ICoreFrameworkInputViewAnimationStartingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PrimaryViewAnimationStarting, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreFrameworkInputView) Remove_PrimaryViewAnimationStarting(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PrimaryViewAnimationStarting, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ICoreFrameworkInputView) Add_OcclusionsChanged(handler TypedEventHandler[*ICoreFrameworkInputView, *ICoreFrameworkInputViewOcclusionsChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_OcclusionsChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreFrameworkInputView) Remove_OcclusionsChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_OcclusionsChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// C0EC901C-BBA4-501B-AE8B-65C9E756A719
var IID_ICoreFrameworkInputViewAnimationStartingEventArgs = syscall.GUID{0xC0EC901C, 0xBBA4, 0x501B,
	[8]byte{0xAE, 0x8B, 0x65, 0xC9, 0xE7, 0x56, 0xA7, 0x19}}

type ICoreFrameworkInputViewAnimationStartingEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Occlusions() *IVectorView[*ICoreInputViewOcclusion]
	Get_FrameworkAnimationRecommended() bool
	Get_AnimationDuration() TimeSpan
}

type ICoreFrameworkInputViewAnimationStartingEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Occlusions                    uintptr
	Get_FrameworkAnimationRecommended uintptr
	Get_AnimationDuration             uintptr
}

type ICoreFrameworkInputViewAnimationStartingEventArgs struct {
	win32.IInspectable
}

func (this *ICoreFrameworkInputViewAnimationStartingEventArgs) Vtbl() *ICoreFrameworkInputViewAnimationStartingEventArgsVtbl {
	return (*ICoreFrameworkInputViewAnimationStartingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreFrameworkInputViewAnimationStartingEventArgs) Get_Occlusions() *IVectorView[*ICoreInputViewOcclusion] {
	var _result *IVectorView[*ICoreInputViewOcclusion]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Occlusions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICoreFrameworkInputViewAnimationStartingEventArgs) Get_FrameworkAnimationRecommended() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameworkAnimationRecommended, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreFrameworkInputViewAnimationStartingEventArgs) Get_AnimationDuration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AnimationDuration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F36F4949-C82C-53D1-A75D-2B2BAF0D9B0D
var IID_ICoreFrameworkInputViewOcclusionsChangedEventArgs = syscall.GUID{0xF36F4949, 0xC82C, 0x53D1,
	[8]byte{0xA7, 0x5D, 0x2B, 0x2B, 0xAF, 0x0D, 0x9B, 0x0D}}

type ICoreFrameworkInputViewOcclusionsChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Occlusions() *IVectorView[*ICoreInputViewOcclusion]
	Get_Handled() bool
}

type ICoreFrameworkInputViewOcclusionsChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Occlusions uintptr
	Get_Handled    uintptr
}

type ICoreFrameworkInputViewOcclusionsChangedEventArgs struct {
	win32.IInspectable
}

func (this *ICoreFrameworkInputViewOcclusionsChangedEventArgs) Vtbl() *ICoreFrameworkInputViewOcclusionsChangedEventArgsVtbl {
	return (*ICoreFrameworkInputViewOcclusionsChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreFrameworkInputViewOcclusionsChangedEventArgs) Get_Occlusions() *IVectorView[*ICoreInputViewOcclusion] {
	var _result *IVectorView[*ICoreInputViewOcclusion]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Occlusions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICoreFrameworkInputViewOcclusionsChangedEventArgs) Get_Handled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Handled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 6EEBD9B6-EAC2-5F8B-975F-772EE3E42EEB
var IID_ICoreFrameworkInputViewStatics = syscall.GUID{0x6EEBD9B6, 0xEAC2, 0x5F8B,
	[8]byte{0x97, 0x5F, 0x77, 0x2E, 0xE3, 0xE4, 0x2E, 0xEB}}

type ICoreFrameworkInputViewStaticsInterface interface {
	win32.IInspectableInterface
	GetForUIContext(context unsafe.Pointer) *ICoreFrameworkInputView
	GetForCurrentView() *ICoreFrameworkInputView
}

type ICoreFrameworkInputViewStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForUIContext   uintptr
	GetForCurrentView uintptr
}

type ICoreFrameworkInputViewStatics struct {
	win32.IInspectable
}

func (this *ICoreFrameworkInputViewStatics) Vtbl() *ICoreFrameworkInputViewStaticsVtbl {
	return (*ICoreFrameworkInputViewStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreFrameworkInputViewStatics) GetForUIContext(context unsafe.Pointer) *ICoreFrameworkInputView {
	var _result *ICoreFrameworkInputView
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForUIContext, uintptr(unsafe.Pointer(this)), uintptr(context), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICoreFrameworkInputViewStatics) GetForCurrentView() *ICoreFrameworkInputView {
	var _result *ICoreFrameworkInputView
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C770CD7A-7001-4C32-BF94-25C1F554CBF1
var IID_ICoreInputView = syscall.GUID{0xC770CD7A, 0x7001, 0x4C32,
	[8]byte{0xBF, 0x94, 0x25, 0xC1, 0xF5, 0x54, 0xCB, 0xF1}}

type ICoreInputViewInterface interface {
	win32.IInspectableInterface
	Add_OcclusionsChanged(handler TypedEventHandler[*ICoreInputView, *ICoreInputViewOcclusionsChangedEventArgs]) EventRegistrationToken
	Remove_OcclusionsChanged(token EventRegistrationToken)
	GetCoreInputViewOcclusions() *IVectorView[*ICoreInputViewOcclusion]
	TryShowPrimaryView() bool
	TryHidePrimaryView() bool
}

type ICoreInputViewVtbl struct {
	win32.IInspectableVtbl
	Add_OcclusionsChanged      uintptr
	Remove_OcclusionsChanged   uintptr
	GetCoreInputViewOcclusions uintptr
	TryShowPrimaryView         uintptr
	TryHidePrimaryView         uintptr
}

type ICoreInputView struct {
	win32.IInspectable
}

func (this *ICoreInputView) Vtbl() *ICoreInputViewVtbl {
	return (*ICoreInputViewVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreInputView) Add_OcclusionsChanged(handler TypedEventHandler[*ICoreInputView, *ICoreInputViewOcclusionsChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_OcclusionsChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreInputView) Remove_OcclusionsChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_OcclusionsChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ICoreInputView) GetCoreInputViewOcclusions() *IVectorView[*ICoreInputViewOcclusion] {
	var _result *IVectorView[*ICoreInputViewOcclusion]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCoreInputViewOcclusions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICoreInputView) TryShowPrimaryView() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryShowPrimaryView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreInputView) TryHidePrimaryView() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryHidePrimaryView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 0ED726C1-E09A-4AE8-AEDF-DFA4857D1A01
var IID_ICoreInputView2 = syscall.GUID{0x0ED726C1, 0xE09A, 0x4AE8,
	[8]byte{0xAE, 0xDF, 0xDF, 0xA4, 0x85, 0x7D, 0x1A, 0x01}}

type ICoreInputView2Interface interface {
	win32.IInspectableInterface
	Add_XYFocusTransferringFromPrimaryView(handler TypedEventHandler[*ICoreInputView, *ICoreInputViewTransferringXYFocusEventArgs]) EventRegistrationToken
	Remove_XYFocusTransferringFromPrimaryView(token EventRegistrationToken)
	Add_XYFocusTransferredToPrimaryView(handler TypedEventHandler[*ICoreInputView, interface{}]) EventRegistrationToken
	Remove_XYFocusTransferredToPrimaryView(token EventRegistrationToken)
	TryTransferXYFocusToPrimaryView(origin Rect, direction CoreInputViewXYFocusTransferDirection) bool
}

type ICoreInputView2Vtbl struct {
	win32.IInspectableVtbl
	Add_XYFocusTransferringFromPrimaryView    uintptr
	Remove_XYFocusTransferringFromPrimaryView uintptr
	Add_XYFocusTransferredToPrimaryView       uintptr
	Remove_XYFocusTransferredToPrimaryView    uintptr
	TryTransferXYFocusToPrimaryView           uintptr
}

type ICoreInputView2 struct {
	win32.IInspectable
}

func (this *ICoreInputView2) Vtbl() *ICoreInputView2Vtbl {
	return (*ICoreInputView2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreInputView2) Add_XYFocusTransferringFromPrimaryView(handler TypedEventHandler[*ICoreInputView, *ICoreInputViewTransferringXYFocusEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_XYFocusTransferringFromPrimaryView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreInputView2) Remove_XYFocusTransferringFromPrimaryView(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_XYFocusTransferringFromPrimaryView, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ICoreInputView2) Add_XYFocusTransferredToPrimaryView(handler TypedEventHandler[*ICoreInputView, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_XYFocusTransferredToPrimaryView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreInputView2) Remove_XYFocusTransferredToPrimaryView(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_XYFocusTransferredToPrimaryView, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ICoreInputView2) TryTransferXYFocusToPrimaryView(origin Rect, direction CoreInputViewXYFocusTransferDirection) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryTransferXYFocusToPrimaryView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&origin)), uintptr(direction), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// BC941653-3AB9-4849-8F58-46E7F0353CFC
var IID_ICoreInputView3 = syscall.GUID{0xBC941653, 0x3AB9, 0x4849,
	[8]byte{0x8F, 0x58, 0x46, 0xE7, 0xF0, 0x35, 0x3C, 0xFC}}

type ICoreInputView3Interface interface {
	win32.IInspectableInterface
	TryShow() bool
	TryShowWithKind(type_ CoreInputViewKind) bool
	TryHide() bool
}

type ICoreInputView3Vtbl struct {
	win32.IInspectableVtbl
	TryShow         uintptr
	TryShowWithKind uintptr
	TryHide         uintptr
}

type ICoreInputView3 struct {
	win32.IInspectable
}

func (this *ICoreInputView3) Vtbl() *ICoreInputView3Vtbl {
	return (*ICoreInputView3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreInputView3) TryShow() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryShow, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreInputView3) TryShowWithKind(type_ CoreInputViewKind) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryShowWithKind, uintptr(unsafe.Pointer(this)), uintptr(type_), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreInputView3) TryHide() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryHide, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 002863D6-D9EF-57EB-8CEF-77F6CE1B7EE7
var IID_ICoreInputView4 = syscall.GUID{0x002863D6, 0xD9EF, 0x57EB,
	[8]byte{0x8C, 0xEF, 0x77, 0xF6, 0xCE, 0x1B, 0x7E, 0xE7}}

type ICoreInputView4Interface interface {
	win32.IInspectableInterface
	Add_PrimaryViewShowing(handler TypedEventHandler[*ICoreInputView, *ICoreInputViewShowingEventArgs]) EventRegistrationToken
	Remove_PrimaryViewShowing(token EventRegistrationToken)
	Add_PrimaryViewHiding(handler TypedEventHandler[*ICoreInputView, *ICoreInputViewHidingEventArgs]) EventRegistrationToken
	Remove_PrimaryViewHiding(token EventRegistrationToken)
}

type ICoreInputView4Vtbl struct {
	win32.IInspectableVtbl
	Add_PrimaryViewShowing    uintptr
	Remove_PrimaryViewShowing uintptr
	Add_PrimaryViewHiding     uintptr
	Remove_PrimaryViewHiding  uintptr
}

type ICoreInputView4 struct {
	win32.IInspectable
}

func (this *ICoreInputView4) Vtbl() *ICoreInputView4Vtbl {
	return (*ICoreInputView4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreInputView4) Add_PrimaryViewShowing(handler TypedEventHandler[*ICoreInputView, *ICoreInputViewShowingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PrimaryViewShowing, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreInputView4) Remove_PrimaryViewShowing(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PrimaryViewShowing, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ICoreInputView4) Add_PrimaryViewHiding(handler TypedEventHandler[*ICoreInputView, *ICoreInputViewHidingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PrimaryViewHiding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreInputView4) Remove_PrimaryViewHiding(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PrimaryViewHiding, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 136316E0-C6D5-5C57-811E-1AD8A99BA6AB
var IID_ICoreInputView5 = syscall.GUID{0x136316E0, 0xC6D5, 0x5C57,
	[8]byte{0x81, 0x1E, 0x1A, 0xD8, 0xA9, 0x9B, 0xA6, 0xAB}}

type ICoreInputView5Interface interface {
	win32.IInspectableInterface
	IsKindSupported(type_ CoreInputViewKind) bool
	Add_SupportedKindsChanged(handler TypedEventHandler[*ICoreInputView, interface{}]) EventRegistrationToken
	Remove_SupportedKindsChanged(token EventRegistrationToken)
	Add_PrimaryViewAnimationStarting(handler TypedEventHandler[*ICoreInputView, *ICoreInputViewAnimationStartingEventArgs]) EventRegistrationToken
	Remove_PrimaryViewAnimationStarting(token EventRegistrationToken)
}

type ICoreInputView5Vtbl struct {
	win32.IInspectableVtbl
	IsKindSupported                     uintptr
	Add_SupportedKindsChanged           uintptr
	Remove_SupportedKindsChanged        uintptr
	Add_PrimaryViewAnimationStarting    uintptr
	Remove_PrimaryViewAnimationStarting uintptr
}

type ICoreInputView5 struct {
	win32.IInspectable
}

func (this *ICoreInputView5) Vtbl() *ICoreInputView5Vtbl {
	return (*ICoreInputView5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreInputView5) IsKindSupported(type_ CoreInputViewKind) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsKindSupported, uintptr(unsafe.Pointer(this)), uintptr(type_), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreInputView5) Add_SupportedKindsChanged(handler TypedEventHandler[*ICoreInputView, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SupportedKindsChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreInputView5) Remove_SupportedKindsChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SupportedKindsChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ICoreInputView5) Add_PrimaryViewAnimationStarting(handler TypedEventHandler[*ICoreInputView, *ICoreInputViewAnimationStartingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PrimaryViewAnimationStarting, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreInputView5) Remove_PrimaryViewAnimationStarting(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PrimaryViewAnimationStarting, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// A9144AF2-B55C-5EA1-B8AB-5340F3E94897
var IID_ICoreInputViewAnimationStartingEventArgs = syscall.GUID{0xA9144AF2, 0xB55C, 0x5EA1,
	[8]byte{0xB8, 0xAB, 0x53, 0x40, 0xF3, 0xE9, 0x48, 0x97}}

type ICoreInputViewAnimationStartingEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Occlusions() *IVectorView[*ICoreInputViewOcclusion]
	Get_Handled() bool
	Put_Handled(value bool)
	Get_AnimationDuration() TimeSpan
}

type ICoreInputViewAnimationStartingEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Occlusions        uintptr
	Get_Handled           uintptr
	Put_Handled           uintptr
	Get_AnimationDuration uintptr
}

type ICoreInputViewAnimationStartingEventArgs struct {
	win32.IInspectable
}

func (this *ICoreInputViewAnimationStartingEventArgs) Vtbl() *ICoreInputViewAnimationStartingEventArgsVtbl {
	return (*ICoreInputViewAnimationStartingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreInputViewAnimationStartingEventArgs) Get_Occlusions() *IVectorView[*ICoreInputViewOcclusion] {
	var _result *IVectorView[*ICoreInputViewOcclusion]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Occlusions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICoreInputViewAnimationStartingEventArgs) Get_Handled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Handled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreInputViewAnimationStartingEventArgs) Put_Handled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Handled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICoreInputViewAnimationStartingEventArgs) Get_AnimationDuration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AnimationDuration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// EADA47BD-BAC5-5336-848D-41083584DAAD
var IID_ICoreInputViewHidingEventArgs = syscall.GUID{0xEADA47BD, 0xBAC5, 0x5336,
	[8]byte{0x84, 0x8D, 0x41, 0x08, 0x35, 0x84, 0xDA, 0xAD}}

type ICoreInputViewHidingEventArgsInterface interface {
	win32.IInspectableInterface
	TryCancel() bool
}

type ICoreInputViewHidingEventArgsVtbl struct {
	win32.IInspectableVtbl
	TryCancel uintptr
}

type ICoreInputViewHidingEventArgs struct {
	win32.IInspectable
}

func (this *ICoreInputViewHidingEventArgs) Vtbl() *ICoreInputViewHidingEventArgsVtbl {
	return (*ICoreInputViewHidingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreInputViewHidingEventArgs) TryCancel() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryCancel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// CC36CE06-3865-4177-B5F5-8B65E0B9CE84
var IID_ICoreInputViewOcclusion = syscall.GUID{0xCC36CE06, 0x3865, 0x4177,
	[8]byte{0xB5, 0xF5, 0x8B, 0x65, 0xE0, 0xB9, 0xCE, 0x84}}

type ICoreInputViewOcclusionInterface interface {
	win32.IInspectableInterface
	Get_OccludingRect() Rect
	Get_OcclusionKind() CoreInputViewOcclusionKind
}

type ICoreInputViewOcclusionVtbl struct {
	win32.IInspectableVtbl
	Get_OccludingRect uintptr
	Get_OcclusionKind uintptr
}

type ICoreInputViewOcclusion struct {
	win32.IInspectable
}

func (this *ICoreInputViewOcclusion) Vtbl() *ICoreInputViewOcclusionVtbl {
	return (*ICoreInputViewOcclusionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreInputViewOcclusion) Get_OccludingRect() Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OccludingRect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreInputViewOcclusion) Get_OcclusionKind() CoreInputViewOcclusionKind {
	var _result CoreInputViewOcclusionKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OcclusionKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// BE1027E8-B3EE-4DF7-9554-89CDC66082C2
var IID_ICoreInputViewOcclusionsChangedEventArgs = syscall.GUID{0xBE1027E8, 0xB3EE, 0x4DF7,
	[8]byte{0x95, 0x54, 0x89, 0xCD, 0xC6, 0x60, 0x82, 0xC2}}

type ICoreInputViewOcclusionsChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Occlusions() *IVectorView[*ICoreInputViewOcclusion]
	Get_Handled() bool
	Put_Handled(value bool)
}

type ICoreInputViewOcclusionsChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Occlusions uintptr
	Get_Handled    uintptr
	Put_Handled    uintptr
}

type ICoreInputViewOcclusionsChangedEventArgs struct {
	win32.IInspectable
}

func (this *ICoreInputViewOcclusionsChangedEventArgs) Vtbl() *ICoreInputViewOcclusionsChangedEventArgsVtbl {
	return (*ICoreInputViewOcclusionsChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreInputViewOcclusionsChangedEventArgs) Get_Occlusions() *IVectorView[*ICoreInputViewOcclusion] {
	var _result *IVectorView[*ICoreInputViewOcclusion]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Occlusions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICoreInputViewOcclusionsChangedEventArgs) Get_Handled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Handled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreInputViewOcclusionsChangedEventArgs) Put_Handled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Handled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// CA52261B-FB9E-5DAF-A98C-262B8B76AF50
var IID_ICoreInputViewShowingEventArgs = syscall.GUID{0xCA52261B, 0xFB9E, 0x5DAF,
	[8]byte{0xA9, 0x8C, 0x26, 0x2B, 0x8B, 0x76, 0xAF, 0x50}}

type ICoreInputViewShowingEventArgsInterface interface {
	win32.IInspectableInterface
	TryCancel() bool
}

type ICoreInputViewShowingEventArgsVtbl struct {
	win32.IInspectableVtbl
	TryCancel uintptr
}

type ICoreInputViewShowingEventArgs struct {
	win32.IInspectable
}

func (this *ICoreInputViewShowingEventArgs) Vtbl() *ICoreInputViewShowingEventArgsVtbl {
	return (*ICoreInputViewShowingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreInputViewShowingEventArgs) TryCancel() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryCancel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7D9B97CD-EDBE-49CF-A54F-337DE052907F
var IID_ICoreInputViewStatics = syscall.GUID{0x7D9B97CD, 0xEDBE, 0x49CF,
	[8]byte{0xA5, 0x4F, 0x33, 0x7D, 0xE0, 0x52, 0x90, 0x7F}}

type ICoreInputViewStaticsInterface interface {
	win32.IInspectableInterface
	GetForCurrentView() *ICoreInputView
}

type ICoreInputViewStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForCurrentView uintptr
}

type ICoreInputViewStatics struct {
	win32.IInspectable
}

func (this *ICoreInputViewStatics) Vtbl() *ICoreInputViewStaticsVtbl {
	return (*ICoreInputViewStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreInputViewStatics) GetForCurrentView() *ICoreInputView {
	var _result *ICoreInputView
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7EBC0862-D049-4E52-87B0-1E90E98C49ED
var IID_ICoreInputViewStatics2 = syscall.GUID{0x7EBC0862, 0xD049, 0x4E52,
	[8]byte{0x87, 0xB0, 0x1E, 0x90, 0xE9, 0x8C, 0x49, 0xED}}

type ICoreInputViewStatics2Interface interface {
	win32.IInspectableInterface
	GetForUIContext(context unsafe.Pointer) *ICoreInputView
}

type ICoreInputViewStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetForUIContext uintptr
}

type ICoreInputViewStatics2 struct {
	win32.IInspectable
}

func (this *ICoreInputViewStatics2) Vtbl() *ICoreInputViewStatics2Vtbl {
	return (*ICoreInputViewStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreInputViewStatics2) GetForUIContext(context unsafe.Pointer) *ICoreInputView {
	var _result *ICoreInputView
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForUIContext, uintptr(unsafe.Pointer(this)), uintptr(context), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 04DE169F-BA02-4850-8B55-D82D03BA6D7F
var IID_ICoreInputViewTransferringXYFocusEventArgs = syscall.GUID{0x04DE169F, 0xBA02, 0x4850,
	[8]byte{0x8B, 0x55, 0xD8, 0x2D, 0x03, 0xBA, 0x6D, 0x7F}}

type ICoreInputViewTransferringXYFocusEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Origin() Rect
	Get_Direction() CoreInputViewXYFocusTransferDirection
	Put_TransferHandled(value bool)
	Get_TransferHandled() bool
	Put_KeepPrimaryViewVisible(value bool)
	Get_KeepPrimaryViewVisible() bool
}

type ICoreInputViewTransferringXYFocusEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Origin                 uintptr
	Get_Direction              uintptr
	Put_TransferHandled        uintptr
	Get_TransferHandled        uintptr
	Put_KeepPrimaryViewVisible uintptr
	Get_KeepPrimaryViewVisible uintptr
}

type ICoreInputViewTransferringXYFocusEventArgs struct {
	win32.IInspectable
}

func (this *ICoreInputViewTransferringXYFocusEventArgs) Vtbl() *ICoreInputViewTransferringXYFocusEventArgsVtbl {
	return (*ICoreInputViewTransferringXYFocusEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreInputViewTransferringXYFocusEventArgs) Get_Origin() Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Origin, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreInputViewTransferringXYFocusEventArgs) Get_Direction() CoreInputViewXYFocusTransferDirection {
	var _result CoreInputViewXYFocusTransferDirection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Direction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreInputViewTransferringXYFocusEventArgs) Put_TransferHandled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TransferHandled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICoreInputViewTransferringXYFocusEventArgs) Get_TransferHandled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TransferHandled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreInputViewTransferringXYFocusEventArgs) Put_KeepPrimaryViewVisible(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_KeepPrimaryViewVisible, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICoreInputViewTransferringXYFocusEventArgs) Get_KeepPrimaryViewVisible() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KeepPrimaryViewVisible, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 78A51AC4-15C0-5A1B-A75B-ACBF9CB8BB9E
var IID_IUISettingsController = syscall.GUID{0x78A51AC4, 0x15C0, 0x5A1B,
	[8]byte{0xA7, 0x5B, 0xAC, 0xBF, 0x9C, 0xB8, 0xBB, 0x9E}}

type IUISettingsControllerInterface interface {
	win32.IInspectableInterface
	SetAdvancedEffectsEnabled(value bool)
	SetAnimationsEnabled(value bool)
	SetAutoHideScrollBars(value bool)
	SetMessageDuration(value uint32)
	SetTextScaleFactor(value float64)
}

type IUISettingsControllerVtbl struct {
	win32.IInspectableVtbl
	SetAdvancedEffectsEnabled uintptr
	SetAnimationsEnabled      uintptr
	SetAutoHideScrollBars     uintptr
	SetMessageDuration        uintptr
	SetTextScaleFactor        uintptr
}

type IUISettingsController struct {
	win32.IInspectable
}

func (this *IUISettingsController) Vtbl() *IUISettingsControllerVtbl {
	return (*IUISettingsControllerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUISettingsController) SetAdvancedEffectsEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetAdvancedEffectsEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IUISettingsController) SetAnimationsEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetAnimationsEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IUISettingsController) SetAutoHideScrollBars(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetAutoHideScrollBars, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IUISettingsController) SetMessageDuration(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetMessageDuration, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IUISettingsController) SetTextScaleFactor(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetTextScaleFactor, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// EB3C68CC-C220-578C-8119-7DB324ED26A6
var IID_IUISettingsControllerStatics = syscall.GUID{0xEB3C68CC, 0xC220, 0x578C,
	[8]byte{0x81, 0x19, 0x7D, 0xB3, 0x24, 0xED, 0x26, 0xA6}}

type IUISettingsControllerStaticsInterface interface {
	win32.IInspectableInterface
	RequestDefaultAsync() *IAsyncOperation[*IUISettingsController]
}

type IUISettingsControllerStaticsVtbl struct {
	win32.IInspectableVtbl
	RequestDefaultAsync uintptr
}

type IUISettingsControllerStatics struct {
	win32.IInspectable
}

func (this *IUISettingsControllerStatics) Vtbl() *IUISettingsControllerStaticsVtbl {
	return (*IUISettingsControllerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUISettingsControllerStatics) RequestDefaultAsync() *IAsyncOperation[*IUISettingsController] {
	var _result *IAsyncOperation[*IUISettingsController]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestDefaultAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type CoreFrameworkInputView struct {
	RtClass
	*ICoreFrameworkInputView
}

func NewICoreFrameworkInputViewStatics() *ICoreFrameworkInputViewStatics {
	var p *ICoreFrameworkInputViewStatics
	hs := NewHStr("Windows.UI.ViewManagement.Core.CoreFrameworkInputView")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICoreFrameworkInputViewStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type CoreFrameworkInputViewAnimationStartingEventArgs struct {
	RtClass
	*ICoreFrameworkInputViewAnimationStartingEventArgs
}

type CoreFrameworkInputViewOcclusionsChangedEventArgs struct {
	RtClass
	*ICoreFrameworkInputViewOcclusionsChangedEventArgs
}

type CoreInputView struct {
	RtClass
	*ICoreInputView
}

func NewICoreInputViewStatics2() *ICoreInputViewStatics2 {
	var p *ICoreInputViewStatics2
	hs := NewHStr("Windows.UI.ViewManagement.Core.CoreInputView")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICoreInputViewStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewICoreInputViewStatics() *ICoreInputViewStatics {
	var p *ICoreInputViewStatics
	hs := NewHStr("Windows.UI.ViewManagement.Core.CoreInputView")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICoreInputViewStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type CoreInputViewAnimationStartingEventArgs struct {
	RtClass
	*ICoreInputViewAnimationStartingEventArgs
}

type CoreInputViewHidingEventArgs struct {
	RtClass
	*ICoreInputViewHidingEventArgs
}

type CoreInputViewOcclusion struct {
	RtClass
	*ICoreInputViewOcclusion
}

type CoreInputViewOcclusionsChangedEventArgs struct {
	RtClass
	*ICoreInputViewOcclusionsChangedEventArgs
}

type CoreInputViewShowingEventArgs struct {
	RtClass
	*ICoreInputViewShowingEventArgs
}

type CoreInputViewTransferringXYFocusEventArgs struct {
	RtClass
	*ICoreInputViewTransferringXYFocusEventArgs
}
