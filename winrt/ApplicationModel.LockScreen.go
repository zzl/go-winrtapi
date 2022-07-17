package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// interfaces

// 38EE31AD-D94F-4E7C-81FA-4F4436506281
var IID_ILockApplicationHost = syscall.GUID{0x38EE31AD, 0xD94F, 0x4E7C,
	[8]byte{0x81, 0xFA, 0x4F, 0x44, 0x36, 0x50, 0x62, 0x81}}

type ILockApplicationHostInterface interface {
	win32.IInspectableInterface
	RequestUnlock()
	Add_Unlocking(handler TypedEventHandler[*ILockApplicationHost, *ILockScreenUnlockingEventArgs]) EventRegistrationToken
	Remove_Unlocking(token EventRegistrationToken)
}

type ILockApplicationHostVtbl struct {
	win32.IInspectableVtbl
	RequestUnlock    uintptr
	Add_Unlocking    uintptr
	Remove_Unlocking uintptr
}

type ILockApplicationHost struct {
	win32.IInspectable
}

func (this *ILockApplicationHost) Vtbl() *ILockApplicationHostVtbl {
	return (*ILockApplicationHostVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILockApplicationHost) RequestUnlock() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestUnlock, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *ILockApplicationHost) Add_Unlocking(handler TypedEventHandler[*ILockApplicationHost, *ILockScreenUnlockingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Unlocking, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILockApplicationHost) Remove_Unlocking(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Unlocking, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// F48FAB8E-23D7-4E63-96A1-666FF52D3B2C
var IID_ILockApplicationHostStatics = syscall.GUID{0xF48FAB8E, 0x23D7, 0x4E63,
	[8]byte{0x96, 0xA1, 0x66, 0x6F, 0xF5, 0x2D, 0x3B, 0x2C}}

type ILockApplicationHostStaticsInterface interface {
	win32.IInspectableInterface
	GetForCurrentView() *ILockApplicationHost
}

type ILockApplicationHostStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForCurrentView uintptr
}

type ILockApplicationHostStatics struct {
	win32.IInspectable
}

func (this *ILockApplicationHostStatics) Vtbl() *ILockApplicationHostStaticsVtbl {
	return (*ILockApplicationHostStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILockApplicationHostStatics) GetForCurrentView() *ILockApplicationHost {
	var _result *ILockApplicationHost
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E95105D9-2BFF-4DB0-9B4F-3824778B9C9A
var IID_ILockScreenBadge = syscall.GUID{0xE95105D9, 0x2BFF, 0x4DB0,
	[8]byte{0x9B, 0x4F, 0x38, 0x24, 0x77, 0x8B, 0x9C, 0x9A}}

type ILockScreenBadgeInterface interface {
	win32.IInspectableInterface
	Get_Logo() *IRandomAccessStream
	Get_Glyph() *IRandomAccessStream
	Get_Number() *IReference[uint32]
	Get_AutomationName() string
	LaunchApp()
}

type ILockScreenBadgeVtbl struct {
	win32.IInspectableVtbl
	Get_Logo           uintptr
	Get_Glyph          uintptr
	Get_Number         uintptr
	Get_AutomationName uintptr
	LaunchApp          uintptr
}

type ILockScreenBadge struct {
	win32.IInspectable
}

func (this *ILockScreenBadge) Vtbl() *ILockScreenBadgeVtbl {
	return (*ILockScreenBadgeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILockScreenBadge) Get_Logo() *IRandomAccessStream {
	var _result *IRandomAccessStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Logo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILockScreenBadge) Get_Glyph() *IRandomAccessStream {
	var _result *IRandomAccessStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Glyph, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILockScreenBadge) Get_Number() *IReference[uint32] {
	var _result *IReference[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Number, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILockScreenBadge) Get_AutomationName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AutomationName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILockScreenBadge) LaunchApp() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchApp, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// F59AA65C-9711-4DC9-A630-95B6CB8CDAD0
var IID_ILockScreenInfo = syscall.GUID{0xF59AA65C, 0x9711, 0x4DC9,
	[8]byte{0xA6, 0x30, 0x95, 0xB6, 0xCB, 0x8C, 0xDA, 0xD0}}

type ILockScreenInfoInterface interface {
	win32.IInspectableInterface
	Add_LockScreenImageChanged(handler TypedEventHandler[*ILockScreenInfo, interface{}]) EventRegistrationToken
	Remove_LockScreenImageChanged(token EventRegistrationToken)
	Get_LockScreenImage() *IRandomAccessStream
	Add_BadgesChanged(handler TypedEventHandler[*ILockScreenInfo, interface{}]) EventRegistrationToken
	Remove_BadgesChanged(token EventRegistrationToken)
	Get_Badges() *IVectorView[*ILockScreenBadge]
	Add_DetailTextChanged(handler TypedEventHandler[*ILockScreenInfo, interface{}]) EventRegistrationToken
	Remove_DetailTextChanged(token EventRegistrationToken)
	Get_DetailText() *IVectorView[string]
	Add_AlarmIconChanged(handler TypedEventHandler[*ILockScreenInfo, interface{}]) EventRegistrationToken
	Remove_AlarmIconChanged(token EventRegistrationToken)
	Get_AlarmIcon() *IRandomAccessStream
}

type ILockScreenInfoVtbl struct {
	win32.IInspectableVtbl
	Add_LockScreenImageChanged    uintptr
	Remove_LockScreenImageChanged uintptr
	Get_LockScreenImage           uintptr
	Add_BadgesChanged             uintptr
	Remove_BadgesChanged          uintptr
	Get_Badges                    uintptr
	Add_DetailTextChanged         uintptr
	Remove_DetailTextChanged      uintptr
	Get_DetailText                uintptr
	Add_AlarmIconChanged          uintptr
	Remove_AlarmIconChanged       uintptr
	Get_AlarmIcon                 uintptr
}

type ILockScreenInfo struct {
	win32.IInspectable
}

func (this *ILockScreenInfo) Vtbl() *ILockScreenInfoVtbl {
	return (*ILockScreenInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILockScreenInfo) Add_LockScreenImageChanged(handler TypedEventHandler[*ILockScreenInfo, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_LockScreenImageChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILockScreenInfo) Remove_LockScreenImageChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_LockScreenImageChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ILockScreenInfo) Get_LockScreenImage() *IRandomAccessStream {
	var _result *IRandomAccessStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LockScreenImage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILockScreenInfo) Add_BadgesChanged(handler TypedEventHandler[*ILockScreenInfo, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_BadgesChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILockScreenInfo) Remove_BadgesChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_BadgesChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ILockScreenInfo) Get_Badges() *IVectorView[*ILockScreenBadge] {
	var _result *IVectorView[*ILockScreenBadge]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Badges, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILockScreenInfo) Add_DetailTextChanged(handler TypedEventHandler[*ILockScreenInfo, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DetailTextChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILockScreenInfo) Remove_DetailTextChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DetailTextChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ILockScreenInfo) Get_DetailText() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DetailText, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILockScreenInfo) Add_AlarmIconChanged(handler TypedEventHandler[*ILockScreenInfo, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AlarmIconChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILockScreenInfo) Remove_AlarmIconChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AlarmIconChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ILockScreenInfo) Get_AlarmIcon() *IRandomAccessStream {
	var _result *IRandomAccessStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlarmIcon, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7E7D1AD6-5203-43E7-9BD6-7C3947D1E3FE
var IID_ILockScreenUnlockingDeferral = syscall.GUID{0x7E7D1AD6, 0x5203, 0x43E7,
	[8]byte{0x9B, 0xD6, 0x7C, 0x39, 0x47, 0xD1, 0xE3, 0xFE}}

type ILockScreenUnlockingDeferralInterface interface {
	win32.IInspectableInterface
	Complete()
}

type ILockScreenUnlockingDeferralVtbl struct {
	win32.IInspectableVtbl
	Complete uintptr
}

type ILockScreenUnlockingDeferral struct {
	win32.IInspectable
}

func (this *ILockScreenUnlockingDeferral) Vtbl() *ILockScreenUnlockingDeferralVtbl {
	return (*ILockScreenUnlockingDeferralVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILockScreenUnlockingDeferral) Complete() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Complete, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 44E6C007-75FB-4ABB-9F8B-824748900C71
var IID_ILockScreenUnlockingEventArgs = syscall.GUID{0x44E6C007, 0x75FB, 0x4ABB,
	[8]byte{0x9F, 0x8B, 0x82, 0x47, 0x48, 0x90, 0x0C, 0x71}}

type ILockScreenUnlockingEventArgsInterface interface {
	win32.IInspectableInterface
	GetDeferral() *ILockScreenUnlockingDeferral
	Get_Deadline() DateTime
}

type ILockScreenUnlockingEventArgsVtbl struct {
	win32.IInspectableVtbl
	GetDeferral  uintptr
	Get_Deadline uintptr
}

type ILockScreenUnlockingEventArgs struct {
	win32.IInspectable
}

func (this *ILockScreenUnlockingEventArgs) Vtbl() *ILockScreenUnlockingEventArgsVtbl {
	return (*ILockScreenUnlockingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILockScreenUnlockingEventArgs) GetDeferral() *ILockScreenUnlockingDeferral {
	var _result *ILockScreenUnlockingDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILockScreenUnlockingEventArgs) Get_Deadline() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Deadline, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type LockApplicationHost struct {
	RtClass
	*ILockApplicationHost
}

func NewILockApplicationHostStatics() *ILockApplicationHostStatics {
	var p *ILockApplicationHostStatics
	hs := NewHStr("Windows.ApplicationModel.LockScreen.LockApplicationHost")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ILockApplicationHostStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type LockScreenBadge struct {
	RtClass
	*ILockScreenBadge
}

type LockScreenInfo struct {
	RtClass
	*ILockScreenInfo
}

type LockScreenUnlockingDeferral struct {
	RtClass
	*ILockScreenUnlockingDeferral
}

type LockScreenUnlockingEventArgs struct {
	RtClass
	*ILockScreenUnlockingEventArgs
}
