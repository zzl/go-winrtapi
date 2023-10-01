package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type AccountPictureKind int32

const (
	AccountPictureKind_SmallImage AccountPictureKind = 0
	AccountPictureKind_LargeImage AccountPictureKind = 1
	AccountPictureKind_Video      AccountPictureKind = 2
)

// enum
type SetAccountPictureResult int32

const (
	SetAccountPictureResult_Success             SetAccountPictureResult = 0
	SetAccountPictureResult_ChangeDisabled      SetAccountPictureResult = 1
	SetAccountPictureResult_LargeOrDynamicError SetAccountPictureResult = 2
	SetAccountPictureResult_VideoFrameSizeError SetAccountPictureResult = 3
	SetAccountPictureResult_FileSizeError       SetAccountPictureResult = 4
	SetAccountPictureResult_Failure             SetAccountPictureResult = 5
)

// enum
type SetImageFeedResult int32

const (
	SetImageFeedResult_Success        SetImageFeedResult = 0
	SetImageFeedResult_ChangeDisabled SetImageFeedResult = 1
	SetImageFeedResult_UserCanceled   SetImageFeedResult = 2
)

// structs

type UserProfileContract struct {
}

type UserProfileLockScreenContract struct {
}

// interfaces

// 928BF3D0-CF7C-4AB0-A7DC-6DC5BCD44252
var IID_IAdvertisingManagerForUser = syscall.GUID{0x928BF3D0, 0xCF7C, 0x4AB0,
	[8]byte{0xA7, 0xDC, 0x6D, 0xC5, 0xBC, 0xD4, 0x42, 0x52}}

type IAdvertisingManagerForUserInterface interface {
	win32.IInspectableInterface
	Get_AdvertisingId() string
	Get_User() *IUser
}

type IAdvertisingManagerForUserVtbl struct {
	win32.IInspectableVtbl
	Get_AdvertisingId uintptr
	Get_User          uintptr
}

type IAdvertisingManagerForUser struct {
	win32.IInspectable
}

func (this *IAdvertisingManagerForUser) Vtbl() *IAdvertisingManagerForUserVtbl {
	return (*IAdvertisingManagerForUserVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdvertisingManagerForUser) Get_AdvertisingId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AdvertisingId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAdvertisingManagerForUser) Get_User() *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_User, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// ADD3468C-A273-48CB-B346-3544522D5581
var IID_IAdvertisingManagerStatics = syscall.GUID{0xADD3468C, 0xA273, 0x48CB,
	[8]byte{0xB3, 0x46, 0x35, 0x44, 0x52, 0x2D, 0x55, 0x81}}

type IAdvertisingManagerStaticsInterface interface {
	win32.IInspectableInterface
	Get_AdvertisingId() string
}

type IAdvertisingManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_AdvertisingId uintptr
}

type IAdvertisingManagerStatics struct {
	win32.IInspectable
}

func (this *IAdvertisingManagerStatics) Vtbl() *IAdvertisingManagerStaticsVtbl {
	return (*IAdvertisingManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdvertisingManagerStatics) Get_AdvertisingId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AdvertisingId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// DD0947AF-1A6D-46B0-95BC-F3F9D6BEB9FB
var IID_IAdvertisingManagerStatics2 = syscall.GUID{0xDD0947AF, 0x1A6D, 0x46B0,
	[8]byte{0x95, 0xBC, 0xF3, 0xF9, 0xD6, 0xBE, 0xB9, 0xFB}}

type IAdvertisingManagerStatics2Interface interface {
	win32.IInspectableInterface
	GetForUser(user *IUser) *IAdvertisingManagerForUser
}

type IAdvertisingManagerStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetForUser uintptr
}

type IAdvertisingManagerStatics2 struct {
	win32.IInspectable
}

func (this *IAdvertisingManagerStatics2) Vtbl() *IAdvertisingManagerStatics2Vtbl {
	return (*IAdvertisingManagerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdvertisingManagerStatics2) GetForUser(user *IUser) *IAdvertisingManagerForUser {
	var _result *IAdvertisingManagerForUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1BC57F1C-E971-5757-B8E0-512F8B8C46D2
var IID_IAssignedAccessSettings = syscall.GUID{0x1BC57F1C, 0xE971, 0x5757,
	[8]byte{0xB8, 0xE0, 0x51, 0x2F, 0x8B, 0x8C, 0x46, 0xD2}}

type IAssignedAccessSettingsInterface interface {
	win32.IInspectableInterface
	Get_IsEnabled() bool
	Get_IsSingleAppKioskMode() bool
	Get_User() *IUser
}

type IAssignedAccessSettingsVtbl struct {
	win32.IInspectableVtbl
	Get_IsEnabled            uintptr
	Get_IsSingleAppKioskMode uintptr
	Get_User                 uintptr
}

type IAssignedAccessSettings struct {
	win32.IInspectable
}

func (this *IAssignedAccessSettings) Vtbl() *IAssignedAccessSettingsVtbl {
	return (*IAssignedAccessSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAssignedAccessSettings) Get_IsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAssignedAccessSettings) Get_IsSingleAppKioskMode() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSingleAppKioskMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAssignedAccessSettings) Get_User() *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_User, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 34A81D0D-8A29-5EF3-A7BE-618E6AC3BD01
var IID_IAssignedAccessSettingsStatics = syscall.GUID{0x34A81D0D, 0x8A29, 0x5EF3,
	[8]byte{0xA7, 0xBE, 0x61, 0x8E, 0x6A, 0xC3, 0xBD, 0x01}}

type IAssignedAccessSettingsStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *IAssignedAccessSettings
	GetForUser(user *IUser) *IAssignedAccessSettings
}

type IAssignedAccessSettingsStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
	GetForUser uintptr
}

type IAssignedAccessSettingsStatics struct {
	win32.IInspectable
}

func (this *IAssignedAccessSettingsStatics) Vtbl() *IAssignedAccessSettingsStaticsVtbl {
	return (*IAssignedAccessSettingsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAssignedAccessSettingsStatics) GetDefault() *IAssignedAccessSettings {
	var _result *IAssignedAccessSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAssignedAccessSettingsStatics) GetForUser(user *IUser) *IAssignedAccessSettings {
	var _result *IAssignedAccessSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E5E9ECCD-2711-44E0-973C-491D78048D24
var IID_IDiagnosticsSettings = syscall.GUID{0xE5E9ECCD, 0x2711, 0x44E0,
	[8]byte{0x97, 0x3C, 0x49, 0x1D, 0x78, 0x04, 0x8D, 0x24}}

type IDiagnosticsSettingsInterface interface {
	win32.IInspectableInterface
	Get_CanUseDiagnosticsToTailorExperiences() bool
	Get_User() *IUser
}

type IDiagnosticsSettingsVtbl struct {
	win32.IInspectableVtbl
	Get_CanUseDiagnosticsToTailorExperiences uintptr
	Get_User                                 uintptr
}

type IDiagnosticsSettings struct {
	win32.IInspectable
}

func (this *IDiagnosticsSettings) Vtbl() *IDiagnosticsSettingsVtbl {
	return (*IDiagnosticsSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDiagnosticsSettings) Get_CanUseDiagnosticsToTailorExperiences() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanUseDiagnosticsToTailorExperiences, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDiagnosticsSettings) Get_User() *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_User, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 72D2E80F-5390-4793-990B-3CCC7D6AC9C8
var IID_IDiagnosticsSettingsStatics = syscall.GUID{0x72D2E80F, 0x5390, 0x4793,
	[8]byte{0x99, 0x0B, 0x3C, 0xCC, 0x7D, 0x6A, 0xC9, 0xC8}}

type IDiagnosticsSettingsStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *IDiagnosticsSettings
	GetForUser(user *IUser) *IDiagnosticsSettings
}

type IDiagnosticsSettingsStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
	GetForUser uintptr
}

type IDiagnosticsSettingsStatics struct {
	win32.IInspectable
}

func (this *IDiagnosticsSettingsStatics) Vtbl() *IDiagnosticsSettingsStaticsVtbl {
	return (*IDiagnosticsSettingsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDiagnosticsSettingsStatics) GetDefault() *IDiagnosticsSettings {
	var _result *IDiagnosticsSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDiagnosticsSettingsStatics) GetForUser(user *IUser) *IDiagnosticsSettings {
	var _result *IDiagnosticsSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3E945153-3A5E-452E-A601-F5BAAD2A4870
var IID_IFirstSignInSettings = syscall.GUID{0x3E945153, 0x3A5E, 0x452E,
	[8]byte{0xA6, 0x01, 0xF5, 0xBA, 0xAD, 0x2A, 0x48, 0x70}}

type IFirstSignInSettingsInterface interface {
	win32.IInspectableInterface
}

type IFirstSignInSettingsVtbl struct {
	win32.IInspectableVtbl
}

type IFirstSignInSettings struct {
	win32.IInspectable
}

func (this *IFirstSignInSettings) Vtbl() *IFirstSignInSettingsVtbl {
	return (*IFirstSignInSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 1CE18F0F-1C41-4EA0-B7A2-6F0C1C7E8438
var IID_IFirstSignInSettingsStatics = syscall.GUID{0x1CE18F0F, 0x1C41, 0x4EA0,
	[8]byte{0xB7, 0xA2, 0x6F, 0x0C, 0x1C, 0x7E, 0x84, 0x38}}

type IFirstSignInSettingsStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *IFirstSignInSettings
}

type IFirstSignInSettingsStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
}

type IFirstSignInSettingsStatics struct {
	win32.IInspectable
}

func (this *IFirstSignInSettingsStatics) Vtbl() *IFirstSignInSettingsStaticsVtbl {
	return (*IFirstSignInSettingsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFirstSignInSettingsStatics) GetDefault() *IFirstSignInSettings {
	var _result *IFirstSignInSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 150F0795-4F6E-40BA-A010-E27D81BDA7F5
var IID_IGlobalizationPreferencesForUser = syscall.GUID{0x150F0795, 0x4F6E, 0x40BA,
	[8]byte{0xA0, 0x10, 0xE2, 0x7D, 0x81, 0xBD, 0xA7, 0xF5}}

type IGlobalizationPreferencesForUserInterface interface {
	win32.IInspectableInterface
	Get_User() *IUser
	Get_Calendars() *IVectorView[string]
	Get_Clocks() *IVectorView[string]
	Get_Currencies() *IVectorView[string]
	Get_Languages() *IVectorView[string]
	Get_HomeGeographicRegion() string
	Get_WeekStartsOn() DayOfWeek
}

type IGlobalizationPreferencesForUserVtbl struct {
	win32.IInspectableVtbl
	Get_User                 uintptr
	Get_Calendars            uintptr
	Get_Clocks               uintptr
	Get_Currencies           uintptr
	Get_Languages            uintptr
	Get_HomeGeographicRegion uintptr
	Get_WeekStartsOn         uintptr
}

type IGlobalizationPreferencesForUser struct {
	win32.IInspectable
}

func (this *IGlobalizationPreferencesForUser) Vtbl() *IGlobalizationPreferencesForUserVtbl {
	return (*IGlobalizationPreferencesForUserVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGlobalizationPreferencesForUser) Get_User() *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_User, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalizationPreferencesForUser) Get_Calendars() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Calendars, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalizationPreferencesForUser) Get_Clocks() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Clocks, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalizationPreferencesForUser) Get_Currencies() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Currencies, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalizationPreferencesForUser) Get_Languages() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Languages, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalizationPreferencesForUser) Get_HomeGeographicRegion() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HomeGeographicRegion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGlobalizationPreferencesForUser) Get_WeekStartsOn() DayOfWeek {
	var _result DayOfWeek
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WeekStartsOn, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 01BF4326-ED37-4E96-B0E9-C1340D1EA158
var IID_IGlobalizationPreferencesStatics = syscall.GUID{0x01BF4326, 0xED37, 0x4E96,
	[8]byte{0xB0, 0xE9, 0xC1, 0x34, 0x0D, 0x1E, 0xA1, 0x58}}

type IGlobalizationPreferencesStaticsInterface interface {
	win32.IInspectableInterface
	Get_Calendars() *IVectorView[string]
	Get_Clocks() *IVectorView[string]
	Get_Currencies() *IVectorView[string]
	Get_Languages() *IVectorView[string]
	Get_HomeGeographicRegion() string
	Get_WeekStartsOn() DayOfWeek
}

type IGlobalizationPreferencesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Calendars            uintptr
	Get_Clocks               uintptr
	Get_Currencies           uintptr
	Get_Languages            uintptr
	Get_HomeGeographicRegion uintptr
	Get_WeekStartsOn         uintptr
}

type IGlobalizationPreferencesStatics struct {
	win32.IInspectable
}

func (this *IGlobalizationPreferencesStatics) Vtbl() *IGlobalizationPreferencesStaticsVtbl {
	return (*IGlobalizationPreferencesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGlobalizationPreferencesStatics) Get_Calendars() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Calendars, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalizationPreferencesStatics) Get_Clocks() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Clocks, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalizationPreferencesStatics) Get_Currencies() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Currencies, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalizationPreferencesStatics) Get_Languages() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Languages, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGlobalizationPreferencesStatics) Get_HomeGeographicRegion() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HomeGeographicRegion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGlobalizationPreferencesStatics) Get_WeekStartsOn() DayOfWeek {
	var _result DayOfWeek
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WeekStartsOn, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// FCCE85F1-4300-4CD0-9CAC-1A8E7B7E18F4
var IID_IGlobalizationPreferencesStatics2 = syscall.GUID{0xFCCE85F1, 0x4300, 0x4CD0,
	[8]byte{0x9C, 0xAC, 0x1A, 0x8E, 0x7B, 0x7E, 0x18, 0xF4}}

type IGlobalizationPreferencesStatics2Interface interface {
	win32.IInspectableInterface
	TrySetHomeGeographicRegion(region string) bool
	TrySetLanguages(languageTags *IIterable[string]) bool
}

type IGlobalizationPreferencesStatics2Vtbl struct {
	win32.IInspectableVtbl
	TrySetHomeGeographicRegion uintptr
	TrySetLanguages            uintptr
}

type IGlobalizationPreferencesStatics2 struct {
	win32.IInspectable
}

func (this *IGlobalizationPreferencesStatics2) Vtbl() *IGlobalizationPreferencesStatics2Vtbl {
	return (*IGlobalizationPreferencesStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGlobalizationPreferencesStatics2) TrySetHomeGeographicRegion(region string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TrySetHomeGeographicRegion, uintptr(unsafe.Pointer(this)), NewHStr(region).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGlobalizationPreferencesStatics2) TrySetLanguages(languageTags *IIterable[string]) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TrySetLanguages, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(languageTags)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 1E059733-35F5-40D8-B9E8-AEF3EF856FCE
var IID_IGlobalizationPreferencesStatics3 = syscall.GUID{0x1E059733, 0x35F5, 0x40D8,
	[8]byte{0xB9, 0xE8, 0xAE, 0xF3, 0xEF, 0x85, 0x6F, 0xCE}}

type IGlobalizationPreferencesStatics3Interface interface {
	win32.IInspectableInterface
	GetForUser(user *IUser) *IGlobalizationPreferencesForUser
}

type IGlobalizationPreferencesStatics3Vtbl struct {
	win32.IInspectableVtbl
	GetForUser uintptr
}

type IGlobalizationPreferencesStatics3 struct {
	win32.IInspectable
}

func (this *IGlobalizationPreferencesStatics3) Vtbl() *IGlobalizationPreferencesStatics3Vtbl {
	return (*IGlobalizationPreferencesStatics3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGlobalizationPreferencesStatics3) GetForUser(user *IUser) *IGlobalizationPreferencesForUser {
	var _result *IGlobalizationPreferencesForUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2C0D73F6-03A9-41A6-9B01-495251FF51D5
var IID_ILockScreenImageFeedStatics = syscall.GUID{0x2C0D73F6, 0x03A9, 0x41A6,
	[8]byte{0x9B, 0x01, 0x49, 0x52, 0x51, 0xFF, 0x51, 0xD5}}

type ILockScreenImageFeedStaticsInterface interface {
	win32.IInspectableInterface
	RequestSetImageFeedAsync(syndicationFeedUri *IUriRuntimeClass) *IAsyncOperation[SetImageFeedResult]
	TryRemoveImageFeed() bool
}

type ILockScreenImageFeedStaticsVtbl struct {
	win32.IInspectableVtbl
	RequestSetImageFeedAsync uintptr
	TryRemoveImageFeed       uintptr
}

type ILockScreenImageFeedStatics struct {
	win32.IInspectable
}

func (this *ILockScreenImageFeedStatics) Vtbl() *ILockScreenImageFeedStaticsVtbl {
	return (*ILockScreenImageFeedStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILockScreenImageFeedStatics) RequestSetImageFeedAsync(syndicationFeedUri *IUriRuntimeClass) *IAsyncOperation[SetImageFeedResult] {
	var _result *IAsyncOperation[SetImageFeedResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestSetImageFeedAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(syndicationFeedUri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILockScreenImageFeedStatics) TryRemoveImageFeed() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryRemoveImageFeed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 3EE9D3AD-B607-40AE-B426-7631D9821269
var IID_ILockScreenStatics = syscall.GUID{0x3EE9D3AD, 0xB607, 0x40AE,
	[8]byte{0xB4, 0x26, 0x76, 0x31, 0xD9, 0x82, 0x12, 0x69}}

type ILockScreenStaticsInterface interface {
	win32.IInspectableInterface
	Get_OriginalImageFile() *IUriRuntimeClass
	GetImageStream() *IRandomAccessStream
	SetImageFileAsync(value *IStorageFile) *IAsyncAction
	SetImageStreamAsync(value *IRandomAccessStream) *IAsyncAction
}

type ILockScreenStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_OriginalImageFile uintptr
	GetImageStream        uintptr
	SetImageFileAsync     uintptr
	SetImageStreamAsync   uintptr
}

type ILockScreenStatics struct {
	win32.IInspectable
}

func (this *ILockScreenStatics) Vtbl() *ILockScreenStaticsVtbl {
	return (*ILockScreenStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILockScreenStatics) Get_OriginalImageFile() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OriginalImageFile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILockScreenStatics) GetImageStream() *IRandomAccessStream {
	var _result *IRandomAccessStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetImageStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILockScreenStatics) SetImageFileAsync(value *IStorageFile) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetImageFileAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILockScreenStatics) SetImageStreamAsync(value *IRandomAccessStream) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetImageStreamAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 77F3A910-48FA-489C-934E-2AE85BA8F772
var IID_IUserInformationStatics = syscall.GUID{0x77F3A910, 0x48FA, 0x489C,
	[8]byte{0x93, 0x4E, 0x2A, 0xE8, 0x5B, 0xA8, 0xF7, 0x72}}

type IUserInformationStaticsInterface interface {
	win32.IInspectableInterface
	Get_AccountPictureChangeEnabled() bool
	Get_NameAccessAllowed() bool
	GetAccountPicture(kind AccountPictureKind) *IStorageFile
	SetAccountPictureAsync(image *IStorageFile) *IAsyncOperation[SetAccountPictureResult]
	SetAccountPicturesAsync(smallImage *IStorageFile, largeImage *IStorageFile, video *IStorageFile) *IAsyncOperation[SetAccountPictureResult]
	SetAccountPictureFromStreamAsync(image *IRandomAccessStream) *IAsyncOperation[SetAccountPictureResult]
	SetAccountPicturesFromStreamsAsync(smallImage *IRandomAccessStream, largeImage *IRandomAccessStream, video *IRandomAccessStream) *IAsyncOperation[SetAccountPictureResult]
	Add_AccountPictureChanged(changeHandler EventHandler[interface{}]) EventRegistrationToken
	Remove_AccountPictureChanged(token EventRegistrationToken)
	GetDisplayNameAsync() *IAsyncOperation[string]
	GetFirstNameAsync() *IAsyncOperation[string]
	GetLastNameAsync() *IAsyncOperation[string]
	GetPrincipalNameAsync() *IAsyncOperation[string]
	GetSessionInitiationProtocolUriAsync() *IAsyncOperation[*IUriRuntimeClass]
	GetDomainNameAsync() *IAsyncOperation[string]
}

type IUserInformationStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_AccountPictureChangeEnabled      uintptr
	Get_NameAccessAllowed                uintptr
	GetAccountPicture                    uintptr
	SetAccountPictureAsync               uintptr
	SetAccountPicturesAsync              uintptr
	SetAccountPictureFromStreamAsync     uintptr
	SetAccountPicturesFromStreamsAsync   uintptr
	Add_AccountPictureChanged            uintptr
	Remove_AccountPictureChanged         uintptr
	GetDisplayNameAsync                  uintptr
	GetFirstNameAsync                    uintptr
	GetLastNameAsync                     uintptr
	GetPrincipalNameAsync                uintptr
	GetSessionInitiationProtocolUriAsync uintptr
	GetDomainNameAsync                   uintptr
}

type IUserInformationStatics struct {
	win32.IInspectable
}

func (this *IUserInformationStatics) Vtbl() *IUserInformationStaticsVtbl {
	return (*IUserInformationStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUserInformationStatics) Get_AccountPictureChangeEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AccountPictureChangeEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUserInformationStatics) Get_NameAccessAllowed() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NameAccessAllowed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUserInformationStatics) GetAccountPicture(kind AccountPictureKind) *IStorageFile {
	var _result *IStorageFile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAccountPicture, uintptr(unsafe.Pointer(this)), uintptr(kind), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserInformationStatics) SetAccountPictureAsync(image *IStorageFile) *IAsyncOperation[SetAccountPictureResult] {
	var _result *IAsyncOperation[SetAccountPictureResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetAccountPictureAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(image)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserInformationStatics) SetAccountPicturesAsync(smallImage *IStorageFile, largeImage *IStorageFile, video *IStorageFile) *IAsyncOperation[SetAccountPictureResult] {
	var _result *IAsyncOperation[SetAccountPictureResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetAccountPicturesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(smallImage)), uintptr(unsafe.Pointer(largeImage)), uintptr(unsafe.Pointer(video)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserInformationStatics) SetAccountPictureFromStreamAsync(image *IRandomAccessStream) *IAsyncOperation[SetAccountPictureResult] {
	var _result *IAsyncOperation[SetAccountPictureResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetAccountPictureFromStreamAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(image)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserInformationStatics) SetAccountPicturesFromStreamsAsync(smallImage *IRandomAccessStream, largeImage *IRandomAccessStream, video *IRandomAccessStream) *IAsyncOperation[SetAccountPictureResult] {
	var _result *IAsyncOperation[SetAccountPictureResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetAccountPicturesFromStreamsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(smallImage)), uintptr(unsafe.Pointer(largeImage)), uintptr(unsafe.Pointer(video)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserInformationStatics) Add_AccountPictureChanged(changeHandler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AccountPictureChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(changeHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUserInformationStatics) Remove_AccountPictureChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AccountPictureChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IUserInformationStatics) GetDisplayNameAsync() *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDisplayNameAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserInformationStatics) GetFirstNameAsync() *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFirstNameAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserInformationStatics) GetLastNameAsync() *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetLastNameAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserInformationStatics) GetPrincipalNameAsync() *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPrincipalNameAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserInformationStatics) GetSessionInitiationProtocolUriAsync() *IAsyncOperation[*IUriRuntimeClass] {
	var _result *IAsyncOperation[*IUriRuntimeClass]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSessionInitiationProtocolUriAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserInformationStatics) GetDomainNameAsync() *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDomainNameAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8CEDDAB4-7998-46D5-8DD3-184F1C5F9AB9
var IID_IUserProfilePersonalizationSettings = syscall.GUID{0x8CEDDAB4, 0x7998, 0x46D5,
	[8]byte{0x8D, 0xD3, 0x18, 0x4F, 0x1C, 0x5F, 0x9A, 0xB9}}

type IUserProfilePersonalizationSettingsInterface interface {
	win32.IInspectableInterface
	TrySetLockScreenImageAsync(imageFile *IStorageFile) *IAsyncOperation[bool]
	TrySetWallpaperImageAsync(imageFile *IStorageFile) *IAsyncOperation[bool]
}

type IUserProfilePersonalizationSettingsVtbl struct {
	win32.IInspectableVtbl
	TrySetLockScreenImageAsync uintptr
	TrySetWallpaperImageAsync  uintptr
}

type IUserProfilePersonalizationSettings struct {
	win32.IInspectable
}

func (this *IUserProfilePersonalizationSettings) Vtbl() *IUserProfilePersonalizationSettingsVtbl {
	return (*IUserProfilePersonalizationSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUserProfilePersonalizationSettings) TrySetLockScreenImageAsync(imageFile *IStorageFile) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TrySetLockScreenImageAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(imageFile)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserProfilePersonalizationSettings) TrySetWallpaperImageAsync(imageFile *IStorageFile) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TrySetWallpaperImageAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(imageFile)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 91ACB841-5037-454B-9883-BB772D08DD16
var IID_IUserProfilePersonalizationSettingsStatics = syscall.GUID{0x91ACB841, 0x5037, 0x454B,
	[8]byte{0x98, 0x83, 0xBB, 0x77, 0x2D, 0x08, 0xDD, 0x16}}

type IUserProfilePersonalizationSettingsStaticsInterface interface {
	win32.IInspectableInterface
	Get_Current() *IUserProfilePersonalizationSettings
	IsSupported() bool
}

type IUserProfilePersonalizationSettingsStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Current uintptr
	IsSupported uintptr
}

type IUserProfilePersonalizationSettingsStatics struct {
	win32.IInspectable
}

func (this *IUserProfilePersonalizationSettingsStatics) Vtbl() *IUserProfilePersonalizationSettingsStaticsVtbl {
	return (*IUserProfilePersonalizationSettingsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUserProfilePersonalizationSettingsStatics) Get_Current() *IUserProfilePersonalizationSettings {
	var _result *IUserProfilePersonalizationSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Current, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserProfilePersonalizationSettingsStatics) IsSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type AdvertisingManagerForUser struct {
	RtClass
	*IAdvertisingManagerForUser
}

type DiagnosticsSettings struct {
	RtClass
	*IDiagnosticsSettings
}

func NewIDiagnosticsSettingsStatics() *IDiagnosticsSettingsStatics {
	var p *IDiagnosticsSettingsStatics
	hs := NewHStr("Windows.System.UserProfile.DiagnosticsSettings")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IDiagnosticsSettingsStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type FirstSignInSettings struct {
	RtClass
	*IFirstSignInSettings
}

func NewIFirstSignInSettingsStatics() *IFirstSignInSettingsStatics {
	var p *IFirstSignInSettingsStatics
	hs := NewHStr("Windows.System.UserProfile.FirstSignInSettings")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IFirstSignInSettingsStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type GlobalizationPreferences struct {
	RtClass
}

func NewIGlobalizationPreferencesStatics3() *IGlobalizationPreferencesStatics3 {
	var p *IGlobalizationPreferencesStatics3
	hs := NewHStr("Windows.System.UserProfile.GlobalizationPreferences")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGlobalizationPreferencesStatics3, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIGlobalizationPreferencesStatics() *IGlobalizationPreferencesStatics {
	var p *IGlobalizationPreferencesStatics
	hs := NewHStr("Windows.System.UserProfile.GlobalizationPreferences")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGlobalizationPreferencesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIGlobalizationPreferencesStatics2() *IGlobalizationPreferencesStatics2 {
	var p *IGlobalizationPreferencesStatics2
	hs := NewHStr("Windows.System.UserProfile.GlobalizationPreferences")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGlobalizationPreferencesStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type GlobalizationPreferencesForUser struct {
	RtClass
	*IGlobalizationPreferencesForUser
}

type LockScreen struct {
	RtClass
}

func NewILockScreenImageFeedStatics() *ILockScreenImageFeedStatics {
	var p *ILockScreenImageFeedStatics
	hs := NewHStr("Windows.System.UserProfile.LockScreen")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ILockScreenImageFeedStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewILockScreenStatics() *ILockScreenStatics {
	var p *ILockScreenStatics
	hs := NewHStr("Windows.System.UserProfile.LockScreen")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ILockScreenStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type UserInformation struct {
	RtClass
}

func NewIUserInformationStatics() *IUserInformationStatics {
	var p *IUserInformationStatics
	hs := NewHStr("Windows.System.UserProfile.UserInformation")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IUserInformationStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type UserProfilePersonalizationSettings struct {
	RtClass
	*IUserProfilePersonalizationSettings
}

func NewIUserProfilePersonalizationSettingsStatics() *IUserProfilePersonalizationSettingsStatics {
	var p *IUserProfilePersonalizationSettingsStatics
	hs := NewHStr("Windows.System.UserProfile.UserProfilePersonalizationSettings")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IUserProfilePersonalizationSettingsStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
