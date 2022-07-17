package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type CortanaPermission int32

const (
	CortanaPermission_BrowsingHistory      CortanaPermission = 0
	CortanaPermission_Calendar             CortanaPermission = 1
	CortanaPermission_CallHistory          CortanaPermission = 2
	CortanaPermission_Contacts             CortanaPermission = 3
	CortanaPermission_Email                CortanaPermission = 4
	CortanaPermission_InputPersonalization CortanaPermission = 5
	CortanaPermission_Location             CortanaPermission = 6
	CortanaPermission_Messaging            CortanaPermission = 7
	CortanaPermission_Microphone           CortanaPermission = 8
	CortanaPermission_Personalization      CortanaPermission = 9
	CortanaPermission_PhoneCall            CortanaPermission = 10
)

// enum
type CortanaPermissionsChangeResult int32

const (
	CortanaPermissionsChangeResult_Success          CortanaPermissionsChangeResult = 0
	CortanaPermissionsChangeResult_Unavailable      CortanaPermissionsChangeResult = 1
	CortanaPermissionsChangeResult_DisabledByPolicy CortanaPermissionsChangeResult = 2
)

// interfaces

// 951EC6B1-FC83-586D-8B84-2452C8981625
var IID_ICortanaActionableInsights = syscall.GUID{0x951EC6B1, 0xFC83, 0x586D,
	[8]byte{0x8B, 0x84, 0x24, 0x52, 0xC8, 0x98, 0x16, 0x25}}

type ICortanaActionableInsightsInterface interface {
	win32.IInspectableInterface
	Get_User() *IUser
	IsAvailableAsync() *IAsyncOperation[bool]
	ShowInsightsForImageAsync(imageStream *IRandomAccessStreamReference) *IAsyncAction
	ShowInsightsForImageWithOptionsAsync(imageStream *IRandomAccessStreamReference, options *ICortanaActionableInsightsOptions) *IAsyncAction
	ShowInsightsForTextAsync(text string) *IAsyncAction
	ShowInsightsForTextWithOptionsAsync(text string, options *ICortanaActionableInsightsOptions) *IAsyncAction
	ShowInsightsAsync(datapackage *IDataPackage) *IAsyncAction
	ShowInsightsWithOptionsAsync(datapackage *IDataPackage, options *ICortanaActionableInsightsOptions) *IAsyncAction
}

type ICortanaActionableInsightsVtbl struct {
	win32.IInspectableVtbl
	Get_User                             uintptr
	IsAvailableAsync                     uintptr
	ShowInsightsForImageAsync            uintptr
	ShowInsightsForImageWithOptionsAsync uintptr
	ShowInsightsForTextAsync             uintptr
	ShowInsightsForTextWithOptionsAsync  uintptr
	ShowInsightsAsync                    uintptr
	ShowInsightsWithOptionsAsync         uintptr
}

type ICortanaActionableInsights struct {
	win32.IInspectable
}

func (this *ICortanaActionableInsights) Vtbl() *ICortanaActionableInsightsVtbl {
	return (*ICortanaActionableInsightsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICortanaActionableInsights) Get_User() *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_User, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICortanaActionableInsights) IsAvailableAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsAvailableAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICortanaActionableInsights) ShowInsightsForImageAsync(imageStream *IRandomAccessStreamReference) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowInsightsForImageAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(imageStream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICortanaActionableInsights) ShowInsightsForImageWithOptionsAsync(imageStream *IRandomAccessStreamReference, options *ICortanaActionableInsightsOptions) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowInsightsForImageWithOptionsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(imageStream)), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICortanaActionableInsights) ShowInsightsForTextAsync(text string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowInsightsForTextAsync, uintptr(unsafe.Pointer(this)), NewHStr(text).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICortanaActionableInsights) ShowInsightsForTextWithOptionsAsync(text string, options *ICortanaActionableInsightsOptions) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowInsightsForTextWithOptionsAsync, uintptr(unsafe.Pointer(this)), NewHStr(text).Ptr, uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICortanaActionableInsights) ShowInsightsAsync(datapackage *IDataPackage) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowInsightsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(datapackage)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICortanaActionableInsights) ShowInsightsWithOptionsAsync(datapackage *IDataPackage, options *ICortanaActionableInsightsOptions) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowInsightsWithOptionsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(datapackage)), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AAC2BBCF-9782-5420-B81E-7AE56AF31815
var IID_ICortanaActionableInsightsOptions = syscall.GUID{0xAAC2BBCF, 0x9782, 0x5420,
	[8]byte{0xB8, 0x1E, 0x7A, 0xE5, 0x6A, 0xF3, 0x18, 0x15}}

type ICortanaActionableInsightsOptionsInterface interface {
	win32.IInspectableInterface
	Get_ContentSourceWebLink() *IUriRuntimeClass
	Put_ContentSourceWebLink(value *IUriRuntimeClass)
	Get_SurroundingText() string
	Put_SurroundingText(value string)
}

type ICortanaActionableInsightsOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_ContentSourceWebLink uintptr
	Put_ContentSourceWebLink uintptr
	Get_SurroundingText      uintptr
	Put_SurroundingText      uintptr
}

type ICortanaActionableInsightsOptions struct {
	win32.IInspectable
}

func (this *ICortanaActionableInsightsOptions) Vtbl() *ICortanaActionableInsightsOptionsVtbl {
	return (*ICortanaActionableInsightsOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICortanaActionableInsightsOptions) Get_ContentSourceWebLink() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentSourceWebLink, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICortanaActionableInsightsOptions) Put_ContentSourceWebLink(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ContentSourceWebLink, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ICortanaActionableInsightsOptions) Get_SurroundingText() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SurroundingText, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICortanaActionableInsightsOptions) Put_SurroundingText(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SurroundingText, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// B5DED412-9D2F-5CB5-9B05-356A0B836C10
var IID_ICortanaActionableInsightsStatics = syscall.GUID{0xB5DED412, 0x9D2F, 0x5CB5,
	[8]byte{0x9B, 0x05, 0x35, 0x6A, 0x0B, 0x83, 0x6C, 0x10}}

type ICortanaActionableInsightsStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *ICortanaActionableInsights
	GetForUser(user *IUser) *ICortanaActionableInsights
}

type ICortanaActionableInsightsStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
	GetForUser uintptr
}

type ICortanaActionableInsightsStatics struct {
	win32.IInspectable
}

func (this *ICortanaActionableInsightsStatics) Vtbl() *ICortanaActionableInsightsStaticsVtbl {
	return (*ICortanaActionableInsightsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICortanaActionableInsightsStatics) GetDefault() *ICortanaActionableInsights {
	var _result *ICortanaActionableInsights
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICortanaActionableInsightsStatics) GetForUser(user *IUser) *ICortanaActionableInsights {
	var _result *ICortanaActionableInsights
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 191330E0-8695-438A-9545-3DA4E822DDB4
var IID_ICortanaPermissionsManager = syscall.GUID{0x191330E0, 0x8695, 0x438A,
	[8]byte{0x95, 0x45, 0x3D, 0xA4, 0xE8, 0x22, 0xDD, 0xB4}}

type ICortanaPermissionsManagerInterface interface {
	win32.IInspectableInterface
	IsSupported() bool
	ArePermissionsGrantedAsync(permissions *IIterable[CortanaPermission]) *IAsyncOperation[bool]
	GrantPermissionsAsync(permissions *IIterable[CortanaPermission]) *IAsyncOperation[CortanaPermissionsChangeResult]
	RevokePermissionsAsync(permissions *IIterable[CortanaPermission]) *IAsyncOperation[CortanaPermissionsChangeResult]
}

type ICortanaPermissionsManagerVtbl struct {
	win32.IInspectableVtbl
	IsSupported                uintptr
	ArePermissionsGrantedAsync uintptr
	GrantPermissionsAsync      uintptr
	RevokePermissionsAsync     uintptr
}

type ICortanaPermissionsManager struct {
	win32.IInspectable
}

func (this *ICortanaPermissionsManager) Vtbl() *ICortanaPermissionsManagerVtbl {
	return (*ICortanaPermissionsManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICortanaPermissionsManager) IsSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICortanaPermissionsManager) ArePermissionsGrantedAsync(permissions *IIterable[CortanaPermission]) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ArePermissionsGrantedAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(permissions)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICortanaPermissionsManager) GrantPermissionsAsync(permissions *IIterable[CortanaPermission]) *IAsyncOperation[CortanaPermissionsChangeResult] {
	var _result *IAsyncOperation[CortanaPermissionsChangeResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GrantPermissionsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(permissions)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICortanaPermissionsManager) RevokePermissionsAsync(permissions *IIterable[CortanaPermission]) *IAsyncOperation[CortanaPermissionsChangeResult] {
	var _result *IAsyncOperation[CortanaPermissionsChangeResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RevokePermissionsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(permissions)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 76B1E67A-B045-4414-9D6D-2AD3A5FE3A7E
var IID_ICortanaPermissionsManagerStatics = syscall.GUID{0x76B1E67A, 0xB045, 0x4414,
	[8]byte{0x9D, 0x6D, 0x2A, 0xD3, 0xA5, 0xFE, 0x3A, 0x7E}}

type ICortanaPermissionsManagerStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *ICortanaPermissionsManager
}

type ICortanaPermissionsManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
}

type ICortanaPermissionsManagerStatics struct {
	win32.IInspectable
}

func (this *ICortanaPermissionsManagerStatics) Vtbl() *ICortanaPermissionsManagerStaticsVtbl {
	return (*ICortanaPermissionsManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICortanaPermissionsManagerStatics) GetDefault() *ICortanaPermissionsManager {
	var _result *ICortanaPermissionsManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 54D571A7-8062-40F4-ABE7-DEDFD697B019
var IID_ICortanaSettings = syscall.GUID{0x54D571A7, 0x8062, 0x40F4,
	[8]byte{0xAB, 0xE7, 0xDE, 0xDF, 0xD6, 0x97, 0xB0, 0x19}}

type ICortanaSettingsInterface interface {
	win32.IInspectableInterface
	Get_HasUserConsentToVoiceActivation() bool
	Get_IsVoiceActivationEnabled() bool
	Put_IsVoiceActivationEnabled(value bool)
}

type ICortanaSettingsVtbl struct {
	win32.IInspectableVtbl
	Get_HasUserConsentToVoiceActivation uintptr
	Get_IsVoiceActivationEnabled        uintptr
	Put_IsVoiceActivationEnabled        uintptr
}

type ICortanaSettings struct {
	win32.IInspectable
}

func (this *ICortanaSettings) Vtbl() *ICortanaSettingsVtbl {
	return (*ICortanaSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICortanaSettings) Get_HasUserConsentToVoiceActivation() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasUserConsentToVoiceActivation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICortanaSettings) Get_IsVoiceActivationEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsVoiceActivationEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICortanaSettings) Put_IsVoiceActivationEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsVoiceActivationEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 8B2CCD7E-2EC0-446D-9285-33F07CE8AC04
var IID_ICortanaSettingsStatics = syscall.GUID{0x8B2CCD7E, 0x2EC0, 0x446D,
	[8]byte{0x92, 0x85, 0x33, 0xF0, 0x7C, 0xE8, 0xAC, 0x04}}

type ICortanaSettingsStaticsInterface interface {
	win32.IInspectableInterface
	IsSupported() bool
	GetDefault() *ICortanaSettings
}

type ICortanaSettingsStaticsVtbl struct {
	win32.IInspectableVtbl
	IsSupported uintptr
	GetDefault  uintptr
}

type ICortanaSettingsStatics struct {
	win32.IInspectable
}

func (this *ICortanaSettingsStatics) Vtbl() *ICortanaSettingsStaticsVtbl {
	return (*ICortanaSettingsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICortanaSettingsStatics) IsSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICortanaSettingsStatics) GetDefault() *ICortanaSettings {
	var _result *ICortanaSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type CortanaSettings struct {
	RtClass
	*ICortanaSettings
}

func NewICortanaSettingsStatics() *ICortanaSettingsStatics {
	var p *ICortanaSettingsStatics
	hs := NewHStr("Windows.Services.Cortana.CortanaSettings")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICortanaSettingsStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
