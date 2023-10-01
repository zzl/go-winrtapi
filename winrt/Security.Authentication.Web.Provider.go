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
type WebAccountClientViewType int32

const (
	WebAccountClientViewType_IdOnly          WebAccountClientViewType = 0
	WebAccountClientViewType_IdAndProperties WebAccountClientViewType = 1
)

// enum
type WebAccountProviderOperationKind int32

const (
	WebAccountProviderOperationKind_RequestToken     WebAccountProviderOperationKind = 0
	WebAccountProviderOperationKind_GetTokenSilently WebAccountProviderOperationKind = 1
	WebAccountProviderOperationKind_AddAccount       WebAccountProviderOperationKind = 2
	WebAccountProviderOperationKind_ManageAccount    WebAccountProviderOperationKind = 3
	WebAccountProviderOperationKind_DeleteAccount    WebAccountProviderOperationKind = 4
	WebAccountProviderOperationKind_RetrieveCookies  WebAccountProviderOperationKind = 5
	WebAccountProviderOperationKind_SignOutAccount   WebAccountProviderOperationKind = 6
)

// enum
type WebAccountScope int32

const (
	WebAccountScope_PerUser        WebAccountScope = 0
	WebAccountScope_PerApplication WebAccountScope = 1
)

// enum
// flags
type WebAccountSelectionOptions uint32

const (
	WebAccountSelectionOptions_Default WebAccountSelectionOptions = 0
	WebAccountSelectionOptions_New     WebAccountSelectionOptions = 1
)

// interfaces

// E7BD66BA-0BC7-4C66-BFD4-65D3082CBCA8
var IID_IWebAccountClientView = syscall.GUID{0xE7BD66BA, 0x0BC7, 0x4C66,
	[8]byte{0xBF, 0xD4, 0x65, 0xD3, 0x08, 0x2C, 0xBC, 0xA8}}

type IWebAccountClientViewInterface interface {
	win32.IInspectableInterface
	Get_ApplicationCallbackUri() *IUriRuntimeClass
	Get_Type() WebAccountClientViewType
	Get_AccountPairwiseId() string
}

type IWebAccountClientViewVtbl struct {
	win32.IInspectableVtbl
	Get_ApplicationCallbackUri uintptr
	Get_Type                   uintptr
	Get_AccountPairwiseId      uintptr
}

type IWebAccountClientView struct {
	win32.IInspectable
}

func (this *IWebAccountClientView) Vtbl() *IWebAccountClientViewVtbl {
	return (*IWebAccountClientViewVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountClientView) Get_ApplicationCallbackUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ApplicationCallbackUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountClientView) Get_Type() WebAccountClientViewType {
	var _result WebAccountClientViewType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Type, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWebAccountClientView) Get_AccountPairwiseId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AccountPairwiseId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 616D16A4-DE22-4855-A326-06CEBF2A3F23
var IID_IWebAccountClientViewFactory = syscall.GUID{0x616D16A4, 0xDE22, 0x4855,
	[8]byte{0xA3, 0x26, 0x06, 0xCE, 0xBF, 0x2A, 0x3F, 0x23}}

type IWebAccountClientViewFactoryInterface interface {
	win32.IInspectableInterface
	Create(viewType WebAccountClientViewType, applicationCallbackUri *IUriRuntimeClass) *IWebAccountClientView
	CreateWithPairwiseId(viewType WebAccountClientViewType, applicationCallbackUri *IUriRuntimeClass, accountPairwiseId string) *IWebAccountClientView
}

type IWebAccountClientViewFactoryVtbl struct {
	win32.IInspectableVtbl
	Create               uintptr
	CreateWithPairwiseId uintptr
}

type IWebAccountClientViewFactory struct {
	win32.IInspectable
}

func (this *IWebAccountClientViewFactory) Vtbl() *IWebAccountClientViewFactoryVtbl {
	return (*IWebAccountClientViewFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountClientViewFactory) Create(viewType WebAccountClientViewType, applicationCallbackUri *IUriRuntimeClass) *IWebAccountClientView {
	var _result *IWebAccountClientView
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(viewType), uintptr(unsafe.Pointer(applicationCallbackUri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountClientViewFactory) CreateWithPairwiseId(viewType WebAccountClientViewType, applicationCallbackUri *IUriRuntimeClass, accountPairwiseId string) *IWebAccountClientView {
	var _result *IWebAccountClientView
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithPairwiseId, uintptr(unsafe.Pointer(this)), uintptr(viewType), uintptr(unsafe.Pointer(applicationCallbackUri)), NewHStr(accountPairwiseId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B2E8E1A6-D49A-4032-84BF-1A2847747BF1
var IID_IWebAccountManagerStatics = syscall.GUID{0xB2E8E1A6, 0xD49A, 0x4032,
	[8]byte{0x84, 0xBF, 0x1A, 0x28, 0x47, 0x74, 0x7B, 0xF1}}

type IWebAccountManagerStaticsInterface interface {
	win32.IInspectableInterface
	UpdateWebAccountPropertiesAsync(webAccount *IWebAccount, webAccountUserName string, additionalProperties *IMapView[string, string]) *IAsyncAction
	AddWebAccountAsync(webAccountId string, webAccountUserName string, props *IMapView[string, string]) *IAsyncOperation[*IWebAccount]
	DeleteWebAccountAsync(webAccount *IWebAccount) *IAsyncAction
	FindAllProviderWebAccountsAsync() *IAsyncOperation[*IVectorView[*IWebAccount]]
	PushCookiesAsync(uri *IUriRuntimeClass, cookies *IVectorView[*IHttpCookie]) *IAsyncAction
	SetViewAsync(webAccount *IWebAccount, view *IWebAccountClientView) *IAsyncAction
	ClearViewAsync(webAccount *IWebAccount, applicationCallbackUri *IUriRuntimeClass) *IAsyncAction
	GetViewsAsync(webAccount *IWebAccount) *IAsyncOperation[*IVectorView[*IWebAccountClientView]]
	SetWebAccountPictureAsync(webAccount *IWebAccount, webAccountPicture *IRandomAccessStream) *IAsyncAction
	ClearWebAccountPictureAsync(webAccount *IWebAccount) *IAsyncAction
}

type IWebAccountManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	UpdateWebAccountPropertiesAsync uintptr
	AddWebAccountAsync              uintptr
	DeleteWebAccountAsync           uintptr
	FindAllProviderWebAccountsAsync uintptr
	PushCookiesAsync                uintptr
	SetViewAsync                    uintptr
	ClearViewAsync                  uintptr
	GetViewsAsync                   uintptr
	SetWebAccountPictureAsync       uintptr
	ClearWebAccountPictureAsync     uintptr
}

type IWebAccountManagerStatics struct {
	win32.IInspectable
}

func (this *IWebAccountManagerStatics) Vtbl() *IWebAccountManagerStaticsVtbl {
	return (*IWebAccountManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountManagerStatics) UpdateWebAccountPropertiesAsync(webAccount *IWebAccount, webAccountUserName string, additionalProperties *IMapView[string, string]) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UpdateWebAccountPropertiesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(webAccount)), NewHStr(webAccountUserName).Ptr, uintptr(unsafe.Pointer(additionalProperties)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountManagerStatics) AddWebAccountAsync(webAccountId string, webAccountUserName string, props *IMapView[string, string]) *IAsyncOperation[*IWebAccount] {
	var _result *IAsyncOperation[*IWebAccount]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddWebAccountAsync, uintptr(unsafe.Pointer(this)), NewHStr(webAccountId).Ptr, NewHStr(webAccountUserName).Ptr, uintptr(unsafe.Pointer(props)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountManagerStatics) DeleteWebAccountAsync(webAccount *IWebAccount) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DeleteWebAccountAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(webAccount)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountManagerStatics) FindAllProviderWebAccountsAsync() *IAsyncOperation[*IVectorView[*IWebAccount]] {
	var _result *IAsyncOperation[*IVectorView[*IWebAccount]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllProviderWebAccountsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountManagerStatics) PushCookiesAsync(uri *IUriRuntimeClass, cookies *IVectorView[*IHttpCookie]) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PushCookiesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(cookies)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountManagerStatics) SetViewAsync(webAccount *IWebAccount, view *IWebAccountClientView) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetViewAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(webAccount)), uintptr(unsafe.Pointer(view)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountManagerStatics) ClearViewAsync(webAccount *IWebAccount, applicationCallbackUri *IUriRuntimeClass) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearViewAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(webAccount)), uintptr(unsafe.Pointer(applicationCallbackUri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountManagerStatics) GetViewsAsync(webAccount *IWebAccount) *IAsyncOperation[*IVectorView[*IWebAccountClientView]] {
	var _result *IAsyncOperation[*IVectorView[*IWebAccountClientView]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetViewsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(webAccount)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountManagerStatics) SetWebAccountPictureAsync(webAccount *IWebAccount, webAccountPicture *IRandomAccessStream) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetWebAccountPictureAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(webAccount)), uintptr(unsafe.Pointer(webAccountPicture)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountManagerStatics) ClearWebAccountPictureAsync(webAccount *IWebAccount) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearWebAccountPictureAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(webAccount)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 68A7A829-2D5F-4653-8BB0-BD2FA6BD2D87
var IID_IWebAccountManagerStatics2 = syscall.GUID{0x68A7A829, 0x2D5F, 0x4653,
	[8]byte{0x8B, 0xB0, 0xBD, 0x2F, 0xA6, 0xBD, 0x2D, 0x87}}

type IWebAccountManagerStatics2Interface interface {
	win32.IInspectableInterface
	PullCookiesAsync(uriString string, callerPFN string) *IAsyncAction
}

type IWebAccountManagerStatics2Vtbl struct {
	win32.IInspectableVtbl
	PullCookiesAsync uintptr
}

type IWebAccountManagerStatics2 struct {
	win32.IInspectable
}

func (this *IWebAccountManagerStatics2) Vtbl() *IWebAccountManagerStatics2Vtbl {
	return (*IWebAccountManagerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountManagerStatics2) PullCookiesAsync(uriString string, callerPFN string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PullCookiesAsync, uintptr(unsafe.Pointer(this)), NewHStr(uriString).Ptr, NewHStr(callerPFN).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DD4523A6-8A4F-4AA2-B15E-03F550AF1359
var IID_IWebAccountManagerStatics3 = syscall.GUID{0xDD4523A6, 0x8A4F, 0x4AA2,
	[8]byte{0xB1, 0x5E, 0x03, 0xF5, 0x50, 0xAF, 0x13, 0x59}}

type IWebAccountManagerStatics3Interface interface {
	win32.IInspectableInterface
	FindAllProviderWebAccountsForUserAsync(user *IUser) *IAsyncOperation[*IVectorView[*IWebAccount]]
	AddWebAccountForUserAsync(user *IUser, webAccountId string, webAccountUserName string, props *IMapView[string, string]) *IAsyncOperation[*IWebAccount]
	AddWebAccountWithScopeForUserAsync(user *IUser, webAccountId string, webAccountUserName string, props *IMapView[string, string], scope WebAccountScope) *IAsyncOperation[*IWebAccount]
	AddWebAccountWithScopeAndMapForUserAsync(user *IUser, webAccountId string, webAccountUserName string, props *IMapView[string, string], scope WebAccountScope, perUserWebAccountId string) *IAsyncOperation[*IWebAccount]
}

type IWebAccountManagerStatics3Vtbl struct {
	win32.IInspectableVtbl
	FindAllProviderWebAccountsForUserAsync   uintptr
	AddWebAccountForUserAsync                uintptr
	AddWebAccountWithScopeForUserAsync       uintptr
	AddWebAccountWithScopeAndMapForUserAsync uintptr
}

type IWebAccountManagerStatics3 struct {
	win32.IInspectable
}

func (this *IWebAccountManagerStatics3) Vtbl() *IWebAccountManagerStatics3Vtbl {
	return (*IWebAccountManagerStatics3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountManagerStatics3) FindAllProviderWebAccountsForUserAsync(user *IUser) *IAsyncOperation[*IVectorView[*IWebAccount]] {
	var _result *IAsyncOperation[*IVectorView[*IWebAccount]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllProviderWebAccountsForUserAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountManagerStatics3) AddWebAccountForUserAsync(user *IUser, webAccountId string, webAccountUserName string, props *IMapView[string, string]) *IAsyncOperation[*IWebAccount] {
	var _result *IAsyncOperation[*IWebAccount]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddWebAccountForUserAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), NewHStr(webAccountId).Ptr, NewHStr(webAccountUserName).Ptr, uintptr(unsafe.Pointer(props)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountManagerStatics3) AddWebAccountWithScopeForUserAsync(user *IUser, webAccountId string, webAccountUserName string, props *IMapView[string, string], scope WebAccountScope) *IAsyncOperation[*IWebAccount] {
	var _result *IAsyncOperation[*IWebAccount]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddWebAccountWithScopeForUserAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), NewHStr(webAccountId).Ptr, NewHStr(webAccountUserName).Ptr, uintptr(unsafe.Pointer(props)), uintptr(scope), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountManagerStatics3) AddWebAccountWithScopeAndMapForUserAsync(user *IUser, webAccountId string, webAccountUserName string, props *IMapView[string, string], scope WebAccountScope, perUserWebAccountId string) *IAsyncOperation[*IWebAccount] {
	var _result *IAsyncOperation[*IWebAccount]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddWebAccountWithScopeAndMapForUserAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), NewHStr(webAccountId).Ptr, NewHStr(webAccountUserName).Ptr, uintptr(unsafe.Pointer(props)), uintptr(scope), NewHStr(perUserWebAccountId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 59EBC2D2-F7DB-412F-BC3F-F2FEA04430B4
var IID_IWebAccountManagerStatics4 = syscall.GUID{0x59EBC2D2, 0xF7DB, 0x412F,
	[8]byte{0xBC, 0x3F, 0xF2, 0xFE, 0xA0, 0x44, 0x30, 0xB4}}

type IWebAccountManagerStatics4Interface interface {
	win32.IInspectableInterface
	InvalidateAppCacheForAllAccountsAsync() *IAsyncAction
	InvalidateAppCacheForAccountAsync(webAccount *IWebAccount) *IAsyncAction
}

type IWebAccountManagerStatics4Vtbl struct {
	win32.IInspectableVtbl
	InvalidateAppCacheForAllAccountsAsync uintptr
	InvalidateAppCacheForAccountAsync     uintptr
}

type IWebAccountManagerStatics4 struct {
	win32.IInspectable
}

func (this *IWebAccountManagerStatics4) Vtbl() *IWebAccountManagerStatics4Vtbl {
	return (*IWebAccountManagerStatics4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountManagerStatics4) InvalidateAppCacheForAllAccountsAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InvalidateAppCacheForAllAccountsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountManagerStatics4) InvalidateAppCacheForAccountAsync(webAccount *IWebAccount) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InvalidateAppCacheForAccountAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(webAccount)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E8FA446F-3A1B-48A4-8E90-1E59CA6F54DB
var IID_IWebAccountMapManagerStatics = syscall.GUID{0xE8FA446F, 0x3A1B, 0x48A4,
	[8]byte{0x8E, 0x90, 0x1E, 0x59, 0xCA, 0x6F, 0x54, 0xDB}}

type IWebAccountMapManagerStaticsInterface interface {
	win32.IInspectableInterface
	AddWebAccountWithScopeAndMapAsync(webAccountId string, webAccountUserName string, props *IMapView[string, string], scope WebAccountScope, perUserWebAccountId string) *IAsyncOperation[*IWebAccount]
	SetPerAppToPerUserAccountAsync(perAppAccount *IWebAccount, perUserWebAccountId string) *IAsyncAction
	GetPerUserFromPerAppAccountAsync(perAppAccount *IWebAccount) *IAsyncOperation[*IWebAccount]
	ClearPerUserFromPerAppAccountAsync(perAppAccount *IWebAccount) *IAsyncAction
}

type IWebAccountMapManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	AddWebAccountWithScopeAndMapAsync  uintptr
	SetPerAppToPerUserAccountAsync     uintptr
	GetPerUserFromPerAppAccountAsync   uintptr
	ClearPerUserFromPerAppAccountAsync uintptr
}

type IWebAccountMapManagerStatics struct {
	win32.IInspectable
}

func (this *IWebAccountMapManagerStatics) Vtbl() *IWebAccountMapManagerStaticsVtbl {
	return (*IWebAccountMapManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountMapManagerStatics) AddWebAccountWithScopeAndMapAsync(webAccountId string, webAccountUserName string, props *IMapView[string, string], scope WebAccountScope, perUserWebAccountId string) *IAsyncOperation[*IWebAccount] {
	var _result *IAsyncOperation[*IWebAccount]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddWebAccountWithScopeAndMapAsync, uintptr(unsafe.Pointer(this)), NewHStr(webAccountId).Ptr, NewHStr(webAccountUserName).Ptr, uintptr(unsafe.Pointer(props)), uintptr(scope), NewHStr(perUserWebAccountId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountMapManagerStatics) SetPerAppToPerUserAccountAsync(perAppAccount *IWebAccount, perUserWebAccountId string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetPerAppToPerUserAccountAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(perAppAccount)), NewHStr(perUserWebAccountId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountMapManagerStatics) GetPerUserFromPerAppAccountAsync(perAppAccount *IWebAccount) *IAsyncOperation[*IWebAccount] {
	var _result *IAsyncOperation[*IWebAccount]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPerUserFromPerAppAccountAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(perAppAccount)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountMapManagerStatics) ClearPerUserFromPerAppAccountAsync(perAppAccount *IWebAccount) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearPerUserFromPerAppAccountAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(perAppAccount)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 73EBDCCF-4378-4C79-9335-A5D7AB81594E
var IID_IWebAccountProviderAddAccountOperation = syscall.GUID{0x73EBDCCF, 0x4378, 0x4C79,
	[8]byte{0x93, 0x35, 0xA5, 0xD7, 0xAB, 0x81, 0x59, 0x4E}}

type IWebAccountProviderAddAccountOperationInterface interface {
	win32.IInspectableInterface
	ReportCompleted()
}

type IWebAccountProviderAddAccountOperationVtbl struct {
	win32.IInspectableVtbl
	ReportCompleted uintptr
}

type IWebAccountProviderAddAccountOperation struct {
	win32.IInspectable
}

func (this *IWebAccountProviderAddAccountOperation) Vtbl() *IWebAccountProviderAddAccountOperationVtbl {
	return (*IWebAccountProviderAddAccountOperationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountProviderAddAccountOperation) ReportCompleted() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReportCompleted, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// BBA4ACBB-993B-4D57-BBE4-1421E3668B4C
var IID_IWebAccountProviderBaseReportOperation = syscall.GUID{0xBBA4ACBB, 0x993B, 0x4D57,
	[8]byte{0xBB, 0xE4, 0x14, 0x21, 0xE3, 0x66, 0x8B, 0x4C}}

type IWebAccountProviderBaseReportOperationInterface interface {
	win32.IInspectableInterface
	ReportCompleted()
	ReportError(value *IWebProviderError)
}

type IWebAccountProviderBaseReportOperationVtbl struct {
	win32.IInspectableVtbl
	ReportCompleted uintptr
	ReportError     uintptr
}

type IWebAccountProviderBaseReportOperation struct {
	win32.IInspectable
}

func (this *IWebAccountProviderBaseReportOperation) Vtbl() *IWebAccountProviderBaseReportOperationVtbl {
	return (*IWebAccountProviderBaseReportOperationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountProviderBaseReportOperation) ReportCompleted() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReportCompleted, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IWebAccountProviderBaseReportOperation) ReportError(value *IWebProviderError) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReportError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 0ABB48B8-9E01-49C9-A355-7D48CAF7D6CA
var IID_IWebAccountProviderDeleteAccountOperation = syscall.GUID{0x0ABB48B8, 0x9E01, 0x49C9,
	[8]byte{0xA3, 0x55, 0x7D, 0x48, 0xCA, 0xF7, 0xD6, 0xCA}}

type IWebAccountProviderDeleteAccountOperationInterface interface {
	win32.IInspectableInterface
	Get_WebAccount() *IWebAccount
}

type IWebAccountProviderDeleteAccountOperationVtbl struct {
	win32.IInspectableVtbl
	Get_WebAccount uintptr
}

type IWebAccountProviderDeleteAccountOperation struct {
	win32.IInspectable
}

func (this *IWebAccountProviderDeleteAccountOperation) Vtbl() *IWebAccountProviderDeleteAccountOperationVtbl {
	return (*IWebAccountProviderDeleteAccountOperationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountProviderDeleteAccountOperation) Get_WebAccount() *IWebAccount {
	var _result *IWebAccount
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WebAccount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// ED20DC5C-D21B-463E-A9B7-C1FD0EDAE978
var IID_IWebAccountProviderManageAccountOperation = syscall.GUID{0xED20DC5C, 0xD21B, 0x463E,
	[8]byte{0xA9, 0xB7, 0xC1, 0xFD, 0x0E, 0xDA, 0xE9, 0x78}}

type IWebAccountProviderManageAccountOperationInterface interface {
	win32.IInspectableInterface
	Get_WebAccount() *IWebAccount
	ReportCompleted()
}

type IWebAccountProviderManageAccountOperationVtbl struct {
	win32.IInspectableVtbl
	Get_WebAccount  uintptr
	ReportCompleted uintptr
}

type IWebAccountProviderManageAccountOperation struct {
	win32.IInspectable
}

func (this *IWebAccountProviderManageAccountOperation) Vtbl() *IWebAccountProviderManageAccountOperationVtbl {
	return (*IWebAccountProviderManageAccountOperationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountProviderManageAccountOperation) Get_WebAccount() *IWebAccount {
	var _result *IWebAccount
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WebAccount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountProviderManageAccountOperation) ReportCompleted() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReportCompleted, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 6D5D2426-10B1-419A-A44E-F9C5161574E6
var IID_IWebAccountProviderOperation = syscall.GUID{0x6D5D2426, 0x10B1, 0x419A,
	[8]byte{0xA4, 0x4E, 0xF9, 0xC5, 0x16, 0x15, 0x74, 0xE6}}

type IWebAccountProviderOperationInterface interface {
	win32.IInspectableInterface
	Get_Kind() WebAccountProviderOperationKind
}

type IWebAccountProviderOperationVtbl struct {
	win32.IInspectableVtbl
	Get_Kind uintptr
}

type IWebAccountProviderOperation struct {
	win32.IInspectable
}

func (this *IWebAccountProviderOperation) Vtbl() *IWebAccountProviderOperationVtbl {
	return (*IWebAccountProviderOperationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountProviderOperation) Get_Kind() WebAccountProviderOperationKind {
	var _result WebAccountProviderOperationKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Kind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5A040441-0FA3-4AB1-A01C-20B110358594
var IID_IWebAccountProviderRetrieveCookiesOperation = syscall.GUID{0x5A040441, 0x0FA3, 0x4AB1,
	[8]byte{0xA0, 0x1C, 0x20, 0xB1, 0x10, 0x35, 0x85, 0x94}}

type IWebAccountProviderRetrieveCookiesOperationInterface interface {
	win32.IInspectableInterface
	Get_Context() *IUriRuntimeClass
	Get_Cookies() *IVector[*IHttpCookie]
	Put_Uri(uri *IUriRuntimeClass)
	Get_Uri() *IUriRuntimeClass
	Get_ApplicationCallbackUri() *IUriRuntimeClass
}

type IWebAccountProviderRetrieveCookiesOperationVtbl struct {
	win32.IInspectableVtbl
	Get_Context                uintptr
	Get_Cookies                uintptr
	Put_Uri                    uintptr
	Get_Uri                    uintptr
	Get_ApplicationCallbackUri uintptr
}

type IWebAccountProviderRetrieveCookiesOperation struct {
	win32.IInspectable
}

func (this *IWebAccountProviderRetrieveCookiesOperation) Vtbl() *IWebAccountProviderRetrieveCookiesOperationVtbl {
	return (*IWebAccountProviderRetrieveCookiesOperationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountProviderRetrieveCookiesOperation) Get_Context() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Context, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountProviderRetrieveCookiesOperation) Get_Cookies() *IVector[*IHttpCookie] {
	var _result *IVector[*IHttpCookie]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Cookies, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountProviderRetrieveCookiesOperation) Put_Uri(uri *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Uri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)))
	_ = _hr
}

func (this *IWebAccountProviderRetrieveCookiesOperation) Get_Uri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Uri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountProviderRetrieveCookiesOperation) Get_ApplicationCallbackUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ApplicationCallbackUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B890E21D-0C55-47BC-8C72-04A6FC7CAC07
var IID_IWebAccountProviderSignOutAccountOperation = syscall.GUID{0xB890E21D, 0x0C55, 0x47BC,
	[8]byte{0x8C, 0x72, 0x04, 0xA6, 0xFC, 0x7C, 0xAC, 0x07}}

type IWebAccountProviderSignOutAccountOperationInterface interface {
	win32.IInspectableInterface
	Get_WebAccount() *IWebAccount
	Get_ApplicationCallbackUri() *IUriRuntimeClass
	Get_ClientId() string
}

type IWebAccountProviderSignOutAccountOperationVtbl struct {
	win32.IInspectableVtbl
	Get_WebAccount             uintptr
	Get_ApplicationCallbackUri uintptr
	Get_ClientId               uintptr
}

type IWebAccountProviderSignOutAccountOperation struct {
	win32.IInspectable
}

func (this *IWebAccountProviderSignOutAccountOperation) Vtbl() *IWebAccountProviderSignOutAccountOperationVtbl {
	return (*IWebAccountProviderSignOutAccountOperationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountProviderSignOutAccountOperation) Get_WebAccount() *IWebAccount {
	var _result *IWebAccount
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WebAccount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountProviderSignOutAccountOperation) Get_ApplicationCallbackUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ApplicationCallbackUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountProviderSignOutAccountOperation) Get_ClientId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ClientId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// E0B545F8-3B0F-44DA-924C-7B18BAAA62A9
var IID_IWebAccountProviderSilentReportOperation = syscall.GUID{0xE0B545F8, 0x3B0F, 0x44DA,
	[8]byte{0x92, 0x4C, 0x7B, 0x18, 0xBA, 0xAA, 0x62, 0xA9}}

type IWebAccountProviderSilentReportOperationInterface interface {
	win32.IInspectableInterface
	ReportUserInteractionRequired()
	ReportUserInteractionRequiredWithError(value *IWebProviderError)
}

type IWebAccountProviderSilentReportOperationVtbl struct {
	win32.IInspectableVtbl
	ReportUserInteractionRequired          uintptr
	ReportUserInteractionRequiredWithError uintptr
}

type IWebAccountProviderSilentReportOperation struct {
	win32.IInspectable
}

func (this *IWebAccountProviderSilentReportOperation) Vtbl() *IWebAccountProviderSilentReportOperationVtbl {
	return (*IWebAccountProviderSilentReportOperationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountProviderSilentReportOperation) ReportUserInteractionRequired() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReportUserInteractionRequired, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IWebAccountProviderSilentReportOperation) ReportUserInteractionRequiredWithError(value *IWebProviderError) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReportUserInteractionRequiredWithError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 408F284B-1328-42DB-89A4-0BCE7A717D8E
var IID_IWebAccountProviderTokenObjects = syscall.GUID{0x408F284B, 0x1328, 0x42DB,
	[8]byte{0x89, 0xA4, 0x0B, 0xCE, 0x7A, 0x71, 0x7D, 0x8E}}

type IWebAccountProviderTokenObjectsInterface interface {
	win32.IInspectableInterface
	Get_Operation() *IWebAccountProviderOperation
}

type IWebAccountProviderTokenObjectsVtbl struct {
	win32.IInspectableVtbl
	Get_Operation uintptr
}

type IWebAccountProviderTokenObjects struct {
	win32.IInspectable
}

func (this *IWebAccountProviderTokenObjects) Vtbl() *IWebAccountProviderTokenObjectsVtbl {
	return (*IWebAccountProviderTokenObjectsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountProviderTokenObjects) Get_Operation() *IWebAccountProviderOperation {
	var _result *IWebAccountProviderOperation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Operation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1020B893-5CA5-4FFF-95FB-B820273FC395
var IID_IWebAccountProviderTokenObjects2 = syscall.GUID{0x1020B893, 0x5CA5, 0x4FFF,
	[8]byte{0x95, 0xFB, 0xB8, 0x20, 0x27, 0x3F, 0xC3, 0x95}}

type IWebAccountProviderTokenObjects2Interface interface {
	win32.IInspectableInterface
	Get_User() *IUser
}

type IWebAccountProviderTokenObjects2Vtbl struct {
	win32.IInspectableVtbl
	Get_User uintptr
}

type IWebAccountProviderTokenObjects2 struct {
	win32.IInspectable
}

func (this *IWebAccountProviderTokenObjects2) Vtbl() *IWebAccountProviderTokenObjects2Vtbl {
	return (*IWebAccountProviderTokenObjects2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountProviderTokenObjects2) Get_User() *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_User, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 95C613BE-2034-4C38-9434-D26C14B2B4B2
var IID_IWebAccountProviderTokenOperation = syscall.GUID{0x95C613BE, 0x2034, 0x4C38,
	[8]byte{0x94, 0x34, 0xD2, 0x6C, 0x14, 0xB2, 0xB4, 0xB2}}

type IWebAccountProviderTokenOperationInterface interface {
	win32.IInspectableInterface
	Get_ProviderRequest() *IWebProviderTokenRequest
	Get_ProviderResponses() *IVector[*IWebProviderTokenResponse]
	Put_CacheExpirationTime(value DateTime)
	Get_CacheExpirationTime() DateTime
}

type IWebAccountProviderTokenOperationVtbl struct {
	win32.IInspectableVtbl
	Get_ProviderRequest     uintptr
	Get_ProviderResponses   uintptr
	Put_CacheExpirationTime uintptr
	Get_CacheExpirationTime uintptr
}

type IWebAccountProviderTokenOperation struct {
	win32.IInspectable
}

func (this *IWebAccountProviderTokenOperation) Vtbl() *IWebAccountProviderTokenOperationVtbl {
	return (*IWebAccountProviderTokenOperationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountProviderTokenOperation) Get_ProviderRequest() *IWebProviderTokenRequest {
	var _result *IWebProviderTokenRequest
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProviderRequest, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountProviderTokenOperation) Get_ProviderResponses() *IVector[*IWebProviderTokenResponse] {
	var _result *IVector[*IWebProviderTokenResponse]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProviderResponses, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountProviderTokenOperation) Put_CacheExpirationTime(value DateTime) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CacheExpirationTime, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IWebAccountProviderTokenOperation) Get_CacheExpirationTime() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CacheExpirationTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 28FF92D3-8F80-42FB-944F-B2107BBD42E6
var IID_IWebAccountProviderUIReportOperation = syscall.GUID{0x28FF92D3, 0x8F80, 0x42FB,
	[8]byte{0x94, 0x4F, 0xB2, 0x10, 0x7B, 0xBD, 0x42, 0xE6}}

type IWebAccountProviderUIReportOperationInterface interface {
	win32.IInspectableInterface
	ReportUserCanceled()
}

type IWebAccountProviderUIReportOperationVtbl struct {
	win32.IInspectableVtbl
	ReportUserCanceled uintptr
}

type IWebAccountProviderUIReportOperation struct {
	win32.IInspectable
}

func (this *IWebAccountProviderUIReportOperation) Vtbl() *IWebAccountProviderUIReportOperationVtbl {
	return (*IWebAccountProviderUIReportOperationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountProviderUIReportOperation) ReportUserCanceled() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReportUserCanceled, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 5C6CE37C-12B2-423A-BF3D-85B8D7E53656
var IID_IWebAccountScopeManagerStatics = syscall.GUID{0x5C6CE37C, 0x12B2, 0x423A,
	[8]byte{0xBF, 0x3D, 0x85, 0xB8, 0xD7, 0xE5, 0x36, 0x56}}

type IWebAccountScopeManagerStaticsInterface interface {
	win32.IInspectableInterface
	AddWebAccountWithScopeAsync(webAccountId string, webAccountUserName string, props *IMapView[string, string], scope WebAccountScope) *IAsyncOperation[*IWebAccount]
	SetScopeAsync(webAccount *IWebAccount, scope WebAccountScope) *IAsyncAction
	GetScope(webAccount *IWebAccount) WebAccountScope
}

type IWebAccountScopeManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	AddWebAccountWithScopeAsync uintptr
	SetScopeAsync               uintptr
	GetScope                    uintptr
}

type IWebAccountScopeManagerStatics struct {
	win32.IInspectable
}

func (this *IWebAccountScopeManagerStatics) Vtbl() *IWebAccountScopeManagerStaticsVtbl {
	return (*IWebAccountScopeManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountScopeManagerStatics) AddWebAccountWithScopeAsync(webAccountId string, webAccountUserName string, props *IMapView[string, string], scope WebAccountScope) *IAsyncOperation[*IWebAccount] {
	var _result *IAsyncOperation[*IWebAccount]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddWebAccountWithScopeAsync, uintptr(unsafe.Pointer(this)), NewHStr(webAccountId).Ptr, NewHStr(webAccountUserName).Ptr, uintptr(unsafe.Pointer(props)), uintptr(scope), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountScopeManagerStatics) SetScopeAsync(webAccount *IWebAccount, scope WebAccountScope) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetScopeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(webAccount)), uintptr(scope), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccountScopeManagerStatics) GetScope(webAccount *IWebAccount) WebAccountScope {
	var _result WebAccountScope
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetScope, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(webAccount)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 1E18778B-8805-454B-9F11-468D2AF1095A
var IID_IWebProviderTokenRequest = syscall.GUID{0x1E18778B, 0x8805, 0x454B,
	[8]byte{0x9F, 0x11, 0x46, 0x8D, 0x2A, 0xF1, 0x09, 0x5A}}

type IWebProviderTokenRequestInterface interface {
	win32.IInspectableInterface
	Get_ClientRequest() *IWebTokenRequest
	Get_WebAccounts() *IVectorView[*IWebAccount]
	Get_WebAccountSelectionOptions() WebAccountSelectionOptions
	Get_ApplicationCallbackUri() *IUriRuntimeClass
	GetApplicationTokenBindingKeyAsync(keyType unsafe.Pointer, target *IUriRuntimeClass) *IAsyncOperation[*ICryptographicKey]
}

type IWebProviderTokenRequestVtbl struct {
	win32.IInspectableVtbl
	Get_ClientRequest                  uintptr
	Get_WebAccounts                    uintptr
	Get_WebAccountSelectionOptions     uintptr
	Get_ApplicationCallbackUri         uintptr
	GetApplicationTokenBindingKeyAsync uintptr
}

type IWebProviderTokenRequest struct {
	win32.IInspectable
}

func (this *IWebProviderTokenRequest) Vtbl() *IWebProviderTokenRequestVtbl {
	return (*IWebProviderTokenRequestVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebProviderTokenRequest) Get_ClientRequest() *IWebTokenRequest {
	var _result *IWebTokenRequest
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ClientRequest, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebProviderTokenRequest) Get_WebAccounts() *IVectorView[*IWebAccount] {
	var _result *IVectorView[*IWebAccount]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WebAccounts, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebProviderTokenRequest) Get_WebAccountSelectionOptions() WebAccountSelectionOptions {
	var _result WebAccountSelectionOptions
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WebAccountSelectionOptions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWebProviderTokenRequest) Get_ApplicationCallbackUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ApplicationCallbackUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebProviderTokenRequest) GetApplicationTokenBindingKeyAsync(keyType unsafe.Pointer, target *IUriRuntimeClass) *IAsyncOperation[*ICryptographicKey] {
	var _result *IAsyncOperation[*ICryptographicKey]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetApplicationTokenBindingKeyAsync, uintptr(unsafe.Pointer(this)), uintptr(keyType), uintptr(unsafe.Pointer(target)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B5D72E4C-10B1-4AA6-88B1-0B6C9E0C1E46
var IID_IWebProviderTokenRequest2 = syscall.GUID{0xB5D72E4C, 0x10B1, 0x4AA6,
	[8]byte{0x88, 0xB1, 0x0B, 0x6C, 0x9E, 0x0C, 0x1E, 0x46}}

type IWebProviderTokenRequest2Interface interface {
	win32.IInspectableInterface
	GetApplicationTokenBindingKeyIdAsync(keyType unsafe.Pointer, target *IUriRuntimeClass) *IAsyncOperation[*IBuffer]
}

type IWebProviderTokenRequest2Vtbl struct {
	win32.IInspectableVtbl
	GetApplicationTokenBindingKeyIdAsync uintptr
}

type IWebProviderTokenRequest2 struct {
	win32.IInspectable
}

func (this *IWebProviderTokenRequest2) Vtbl() *IWebProviderTokenRequest2Vtbl {
	return (*IWebProviderTokenRequest2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebProviderTokenRequest2) GetApplicationTokenBindingKeyIdAsync(keyType unsafe.Pointer, target *IUriRuntimeClass) *IAsyncOperation[*IBuffer] {
	var _result *IAsyncOperation[*IBuffer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetApplicationTokenBindingKeyIdAsync, uintptr(unsafe.Pointer(this)), uintptr(keyType), uintptr(unsafe.Pointer(target)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1B2716AA-4289-446E-9256-DAFB6F66A51E
var IID_IWebProviderTokenRequest3 = syscall.GUID{0x1B2716AA, 0x4289, 0x446E,
	[8]byte{0x92, 0x56, 0xDA, 0xFB, 0x6F, 0x66, 0xA5, 0x1E}}

type IWebProviderTokenRequest3Interface interface {
	win32.IInspectableInterface
	Get_ApplicationPackageFamilyName() string
	Get_ApplicationProcessName() string
	CheckApplicationForCapabilityAsync(capabilityName string) *IAsyncOperation[bool]
}

type IWebProviderTokenRequest3Vtbl struct {
	win32.IInspectableVtbl
	Get_ApplicationPackageFamilyName   uintptr
	Get_ApplicationProcessName         uintptr
	CheckApplicationForCapabilityAsync uintptr
}

type IWebProviderTokenRequest3 struct {
	win32.IInspectable
}

func (this *IWebProviderTokenRequest3) Vtbl() *IWebProviderTokenRequest3Vtbl {
	return (*IWebProviderTokenRequest3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebProviderTokenRequest3) Get_ApplicationPackageFamilyName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ApplicationPackageFamilyName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWebProviderTokenRequest3) Get_ApplicationProcessName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ApplicationProcessName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWebProviderTokenRequest3) CheckApplicationForCapabilityAsync(capabilityName string) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CheckApplicationForCapabilityAsync, uintptr(unsafe.Pointer(this)), NewHStr(capabilityName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EF213793-EF55-4186-B7CE-8CB2E7F9849E
var IID_IWebProviderTokenResponse = syscall.GUID{0xEF213793, 0xEF55, 0x4186,
	[8]byte{0xB7, 0xCE, 0x8C, 0xB2, 0xE7, 0xF9, 0x84, 0x9E}}

type IWebProviderTokenResponseInterface interface {
	win32.IInspectableInterface
	Get_ClientResponse() *IWebTokenResponse
}

type IWebProviderTokenResponseVtbl struct {
	win32.IInspectableVtbl
	Get_ClientResponse uintptr
}

type IWebProviderTokenResponse struct {
	win32.IInspectable
}

func (this *IWebProviderTokenResponse) Vtbl() *IWebProviderTokenResponseVtbl {
	return (*IWebProviderTokenResponseVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebProviderTokenResponse) Get_ClientResponse() *IWebTokenResponse {
	var _result *IWebTokenResponse
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ClientResponse, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FA49D99A-25BA-4077-9CFA-9DB4DEA7B71A
var IID_IWebProviderTokenResponseFactory = syscall.GUID{0xFA49D99A, 0x25BA, 0x4077,
	[8]byte{0x9C, 0xFA, 0x9D, 0xB4, 0xDE, 0xA7, 0xB7, 0x1A}}

type IWebProviderTokenResponseFactoryInterface interface {
	win32.IInspectableInterface
	Create(webTokenResponse *IWebTokenResponse) *IWebProviderTokenResponse
}

type IWebProviderTokenResponseFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IWebProviderTokenResponseFactory struct {
	win32.IInspectable
}

func (this *IWebProviderTokenResponseFactory) Vtbl() *IWebProviderTokenResponseFactoryVtbl {
	return (*IWebProviderTokenResponseFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebProviderTokenResponseFactory) Create(webTokenResponse *IWebTokenResponse) *IWebProviderTokenResponse {
	var _result *IWebProviderTokenResponse
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(webTokenResponse)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type WebAccountClientView struct {
	RtClass
	*IWebAccountClientView
}

func NewWebAccountClientView_Create(viewType WebAccountClientViewType, applicationCallbackUri *IUriRuntimeClass) *WebAccountClientView {
	hs := NewHStr("Windows.Security.Authentication.Web.Provider.WebAccountClientView")
	var pFac *IWebAccountClientViewFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWebAccountClientViewFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IWebAccountClientView
	p = pFac.Create(viewType, applicationCallbackUri)
	result := &WebAccountClientView{
		RtClass:               RtClass{PInspect: &p.IInspectable},
		IWebAccountClientView: p,
	}
	com.AddToScope(result)
	return result
}

func NewWebAccountClientView_CreateWithPairwiseId(viewType WebAccountClientViewType, applicationCallbackUri *IUriRuntimeClass, accountPairwiseId string) *WebAccountClientView {
	hs := NewHStr("Windows.Security.Authentication.Web.Provider.WebAccountClientView")
	var pFac *IWebAccountClientViewFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWebAccountClientViewFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IWebAccountClientView
	p = pFac.CreateWithPairwiseId(viewType, applicationCallbackUri, accountPairwiseId)
	result := &WebAccountClientView{
		RtClass:               RtClass{PInspect: &p.IInspectable},
		IWebAccountClientView: p,
	}
	com.AddToScope(result)
	return result
}

type WebAccountManager struct {
	RtClass
}

func NewIWebAccountManagerStatics() *IWebAccountManagerStatics {
	var p *IWebAccountManagerStatics
	hs := NewHStr("Windows.Security.Authentication.Web.Provider.WebAccountManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWebAccountManagerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIWebAccountManagerStatics2() *IWebAccountManagerStatics2 {
	var p *IWebAccountManagerStatics2
	hs := NewHStr("Windows.Security.Authentication.Web.Provider.WebAccountManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWebAccountManagerStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIWebAccountManagerStatics3() *IWebAccountManagerStatics3 {
	var p *IWebAccountManagerStatics3
	hs := NewHStr("Windows.Security.Authentication.Web.Provider.WebAccountManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWebAccountManagerStatics3, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIWebAccountScopeManagerStatics() *IWebAccountScopeManagerStatics {
	var p *IWebAccountScopeManagerStatics
	hs := NewHStr("Windows.Security.Authentication.Web.Provider.WebAccountManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWebAccountScopeManagerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIWebAccountManagerStatics4() *IWebAccountManagerStatics4 {
	var p *IWebAccountManagerStatics4
	hs := NewHStr("Windows.Security.Authentication.Web.Provider.WebAccountManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWebAccountManagerStatics4, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIWebAccountMapManagerStatics() *IWebAccountMapManagerStatics {
	var p *IWebAccountMapManagerStatics
	hs := NewHStr("Windows.Security.Authentication.Web.Provider.WebAccountManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWebAccountMapManagerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type WebAccountProviderAddAccountOperation struct {
	RtClass
	*IWebAccountProviderAddAccountOperation
}

type WebAccountProviderDeleteAccountOperation struct {
	RtClass
	*IWebAccountProviderDeleteAccountOperation
}

type WebAccountProviderGetTokenSilentOperation struct {
	RtClass
	*IWebAccountProviderTokenOperation
}

type WebAccountProviderManageAccountOperation struct {
	RtClass
	*IWebAccountProviderManageAccountOperation
}

type WebAccountProviderRequestTokenOperation struct {
	RtClass
	*IWebAccountProviderTokenOperation
}

type WebAccountProviderRetrieveCookiesOperation struct {
	RtClass
	*IWebAccountProviderRetrieveCookiesOperation
}

type WebAccountProviderSignOutAccountOperation struct {
	RtClass
	*IWebAccountProviderSignOutAccountOperation
}

type WebAccountProviderTriggerDetails struct {
	RtClass
	*IWebAccountProviderTokenObjects
}

type WebProviderTokenRequest struct {
	RtClass
	*IWebProviderTokenRequest
}

type WebProviderTokenResponse struct {
	RtClass
	*IWebProviderTokenResponse
}

func NewWebProviderTokenResponse_Create(webTokenResponse *IWebTokenResponse) *WebProviderTokenResponse {
	hs := NewHStr("Windows.Security.Authentication.Web.Provider.WebProviderTokenResponse")
	var pFac *IWebProviderTokenResponseFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWebProviderTokenResponseFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IWebProviderTokenResponse
	p = pFac.Create(webTokenResponse)
	result := &WebProviderTokenResponse{
		RtClass:                   RtClass{PInspect: &p.IInspectable},
		IWebProviderTokenResponse: p,
	}
	com.AddToScope(result)
	return result
}
