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
type FindAllWebAccountsStatus int32

const (
	FindAllWebAccountsStatus_Success                FindAllWebAccountsStatus = 0
	FindAllWebAccountsStatus_NotAllowedByProvider   FindAllWebAccountsStatus = 1
	FindAllWebAccountsStatus_NotSupportedByProvider FindAllWebAccountsStatus = 2
	FindAllWebAccountsStatus_ProviderError          FindAllWebAccountsStatus = 3
)

// enum
type WebTokenRequestPromptType int32

const (
	WebTokenRequestPromptType_Default             WebTokenRequestPromptType = 0
	WebTokenRequestPromptType_ForceAuthentication WebTokenRequestPromptType = 1
)

// enum
type WebTokenRequestStatus int32

const (
	WebTokenRequestStatus_Success                     WebTokenRequestStatus = 0
	WebTokenRequestStatus_UserCancel                  WebTokenRequestStatus = 1
	WebTokenRequestStatus_AccountSwitch               WebTokenRequestStatus = 2
	WebTokenRequestStatus_UserInteractionRequired     WebTokenRequestStatus = 3
	WebTokenRequestStatus_AccountProviderNotAvailable WebTokenRequestStatus = 4
	WebTokenRequestStatus_ProviderError               WebTokenRequestStatus = 5
)

// interfaces

// A5812B5D-B72E-420C-86AB-AAC0D7B7261F
var IID_IFindAllAccountsResult = syscall.GUID{0xA5812B5D, 0xB72E, 0x420C,
	[8]byte{0x86, 0xAB, 0xAA, 0xC0, 0xD7, 0xB7, 0x26, 0x1F}}

type IFindAllAccountsResultInterface interface {
	win32.IInspectableInterface
	Get_Accounts() *IVectorView[*IWebAccount]
	Get_Status() FindAllWebAccountsStatus
	Get_ProviderError() *IWebProviderError
}

type IFindAllAccountsResultVtbl struct {
	win32.IInspectableVtbl
	Get_Accounts      uintptr
	Get_Status        uintptr
	Get_ProviderError uintptr
}

type IFindAllAccountsResult struct {
	win32.IInspectable
}

func (this *IFindAllAccountsResult) Vtbl() *IFindAllAccountsResultVtbl {
	return (*IFindAllAccountsResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFindAllAccountsResult) Get_Accounts() *IVectorView[*IWebAccount] {
	var _result *IVectorView[*IWebAccount]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Accounts, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFindAllAccountsResult) Get_Status() FindAllWebAccountsStatus {
	var _result FindAllWebAccountsStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFindAllAccountsResult) Get_ProviderError() *IWebProviderError {
	var _result *IWebProviderError
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProviderError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6FB7037D-424E-44EC-977C-EF2415462A5A
var IID_IWebAccountEventArgs = syscall.GUID{0x6FB7037D, 0x424E, 0x44EC,
	[8]byte{0x97, 0x7C, 0xEF, 0x24, 0x15, 0x46, 0x2A, 0x5A}}

type IWebAccountEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Account() *IWebAccount
}

type IWebAccountEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Account uintptr
}

type IWebAccountEventArgs struct {
	win32.IInspectable
}

func (this *IWebAccountEventArgs) Vtbl() *IWebAccountEventArgsVtbl {
	return (*IWebAccountEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountEventArgs) Get_Account() *IWebAccount {
	var _result *IWebAccount
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Account, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7445F5FD-AA9D-4619-8D5D-C138A4EDE3E5
var IID_IWebAccountMonitor = syscall.GUID{0x7445F5FD, 0xAA9D, 0x4619,
	[8]byte{0x8D, 0x5D, 0xC1, 0x38, 0xA4, 0xED, 0xE3, 0xE5}}

type IWebAccountMonitorInterface interface {
	win32.IInspectableInterface
	Add_Updated(handler TypedEventHandler[*IWebAccountMonitor, *IWebAccountEventArgs]) EventRegistrationToken
	Remove_Updated(token EventRegistrationToken)
	Add_Removed(handler TypedEventHandler[*IWebAccountMonitor, *IWebAccountEventArgs]) EventRegistrationToken
	Remove_Removed(token EventRegistrationToken)
	Add_DefaultSignInAccountChanged(handler TypedEventHandler[*IWebAccountMonitor, interface{}]) EventRegistrationToken
	Remove_DefaultSignInAccountChanged(token EventRegistrationToken)
}

type IWebAccountMonitorVtbl struct {
	win32.IInspectableVtbl
	Add_Updated                        uintptr
	Remove_Updated                     uintptr
	Add_Removed                        uintptr
	Remove_Removed                     uintptr
	Add_DefaultSignInAccountChanged    uintptr
	Remove_DefaultSignInAccountChanged uintptr
}

type IWebAccountMonitor struct {
	win32.IInspectable
}

func (this *IWebAccountMonitor) Vtbl() *IWebAccountMonitorVtbl {
	return (*IWebAccountMonitorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountMonitor) Add_Updated(handler TypedEventHandler[*IWebAccountMonitor, *IWebAccountEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Updated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWebAccountMonitor) Remove_Updated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Updated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IWebAccountMonitor) Add_Removed(handler TypedEventHandler[*IWebAccountMonitor, *IWebAccountEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Removed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWebAccountMonitor) Remove_Removed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Removed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IWebAccountMonitor) Add_DefaultSignInAccountChanged(handler TypedEventHandler[*IWebAccountMonitor, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DefaultSignInAccountChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWebAccountMonitor) Remove_DefaultSignInAccountChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DefaultSignInAccountChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// A7ADC1F8-24B8-4F01-9AE5-24545E71233A
var IID_IWebAccountMonitor2 = syscall.GUID{0xA7ADC1F8, 0x24B8, 0x4F01,
	[8]byte{0x9A, 0xE5, 0x24, 0x54, 0x5E, 0x71, 0x23, 0x3A}}

type IWebAccountMonitor2Interface interface {
	win32.IInspectableInterface
	Add_AccountPictureUpdated(handler TypedEventHandler[*IWebAccountMonitor, *IWebAccountEventArgs]) EventRegistrationToken
	Remove_AccountPictureUpdated(token EventRegistrationToken)
}

type IWebAccountMonitor2Vtbl struct {
	win32.IInspectableVtbl
	Add_AccountPictureUpdated    uintptr
	Remove_AccountPictureUpdated uintptr
}

type IWebAccountMonitor2 struct {
	win32.IInspectable
}

func (this *IWebAccountMonitor2) Vtbl() *IWebAccountMonitor2Vtbl {
	return (*IWebAccountMonitor2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountMonitor2) Add_AccountPictureUpdated(handler TypedEventHandler[*IWebAccountMonitor, *IWebAccountEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AccountPictureUpdated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWebAccountMonitor2) Remove_AccountPictureUpdated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AccountPictureUpdated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 6ACA7C92-A581-4479-9C10-752EFF44FD34
var IID_IWebAuthenticationCoreManagerStatics = syscall.GUID{0x6ACA7C92, 0xA581, 0x4479,
	[8]byte{0x9C, 0x10, 0x75, 0x2E, 0xFF, 0x44, 0xFD, 0x34}}

type IWebAuthenticationCoreManagerStaticsInterface interface {
	win32.IInspectableInterface
	GetTokenSilentlyAsync(request *IWebTokenRequest) *IAsyncOperation[*IWebTokenRequestResult]
	GetTokenSilentlyWithWebAccountAsync(request *IWebTokenRequest, webAccount *IWebAccount) *IAsyncOperation[*IWebTokenRequestResult]
	RequestTokenAsync(request *IWebTokenRequest) *IAsyncOperation[*IWebTokenRequestResult]
	RequestTokenWithWebAccountAsync(request *IWebTokenRequest, webAccount *IWebAccount) *IAsyncOperation[*IWebTokenRequestResult]
	FindAccountAsync(provider *IWebAccountProvider, webAccountId string) *IAsyncOperation[*IWebAccount]
	FindAccountProviderAsync(webAccountProviderId string) *IAsyncOperation[*IWebAccountProvider]
	FindAccountProviderWithAuthorityAsync(webAccountProviderId string, authority string) *IAsyncOperation[*IWebAccountProvider]
}

type IWebAuthenticationCoreManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetTokenSilentlyAsync                 uintptr
	GetTokenSilentlyWithWebAccountAsync   uintptr
	RequestTokenAsync                     uintptr
	RequestTokenWithWebAccountAsync       uintptr
	FindAccountAsync                      uintptr
	FindAccountProviderAsync              uintptr
	FindAccountProviderWithAuthorityAsync uintptr
}

type IWebAuthenticationCoreManagerStatics struct {
	win32.IInspectable
}

func (this *IWebAuthenticationCoreManagerStatics) Vtbl() *IWebAuthenticationCoreManagerStaticsVtbl {
	return (*IWebAuthenticationCoreManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAuthenticationCoreManagerStatics) GetTokenSilentlyAsync(request *IWebTokenRequest) *IAsyncOperation[*IWebTokenRequestResult] {
	var _result *IAsyncOperation[*IWebTokenRequestResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetTokenSilentlyAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(request)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAuthenticationCoreManagerStatics) GetTokenSilentlyWithWebAccountAsync(request *IWebTokenRequest, webAccount *IWebAccount) *IAsyncOperation[*IWebTokenRequestResult] {
	var _result *IAsyncOperation[*IWebTokenRequestResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetTokenSilentlyWithWebAccountAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(request)), uintptr(unsafe.Pointer(webAccount)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAuthenticationCoreManagerStatics) RequestTokenAsync(request *IWebTokenRequest) *IAsyncOperation[*IWebTokenRequestResult] {
	var _result *IAsyncOperation[*IWebTokenRequestResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestTokenAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(request)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAuthenticationCoreManagerStatics) RequestTokenWithWebAccountAsync(request *IWebTokenRequest, webAccount *IWebAccount) *IAsyncOperation[*IWebTokenRequestResult] {
	var _result *IAsyncOperation[*IWebTokenRequestResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestTokenWithWebAccountAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(request)), uintptr(unsafe.Pointer(webAccount)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAuthenticationCoreManagerStatics) FindAccountAsync(provider *IWebAccountProvider, webAccountId string) *IAsyncOperation[*IWebAccount] {
	var _result *IAsyncOperation[*IWebAccount]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAccountAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(provider)), NewHStr(webAccountId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAuthenticationCoreManagerStatics) FindAccountProviderAsync(webAccountProviderId string) *IAsyncOperation[*IWebAccountProvider] {
	var _result *IAsyncOperation[*IWebAccountProvider]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAccountProviderAsync, uintptr(unsafe.Pointer(this)), NewHStr(webAccountProviderId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAuthenticationCoreManagerStatics) FindAccountProviderWithAuthorityAsync(webAccountProviderId string, authority string) *IAsyncOperation[*IWebAccountProvider] {
	var _result *IAsyncOperation[*IWebAccountProvider]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAccountProviderWithAuthorityAsync, uintptr(unsafe.Pointer(this)), NewHStr(webAccountProviderId).Ptr, NewHStr(authority).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F584184A-8B57-4820-B6A4-70A5B6FCF44A
var IID_IWebAuthenticationCoreManagerStatics2 = syscall.GUID{0xF584184A, 0x8B57, 0x4820,
	[8]byte{0xB6, 0xA4, 0x70, 0xA5, 0xB6, 0xFC, 0xF4, 0x4A}}

type IWebAuthenticationCoreManagerStatics2Interface interface {
	win32.IInspectableInterface
	FindAccountProviderWithAuthorityForUserAsync(webAccountProviderId string, authority string, user *IUser) *IAsyncOperation[*IWebAccountProvider]
}

type IWebAuthenticationCoreManagerStatics2Vtbl struct {
	win32.IInspectableVtbl
	FindAccountProviderWithAuthorityForUserAsync uintptr
}

type IWebAuthenticationCoreManagerStatics2 struct {
	win32.IInspectable
}

func (this *IWebAuthenticationCoreManagerStatics2) Vtbl() *IWebAuthenticationCoreManagerStatics2Vtbl {
	return (*IWebAuthenticationCoreManagerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAuthenticationCoreManagerStatics2) FindAccountProviderWithAuthorityForUserAsync(webAccountProviderId string, authority string, user *IUser) *IAsyncOperation[*IWebAccountProvider] {
	var _result *IAsyncOperation[*IWebAccountProvider]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAccountProviderWithAuthorityForUserAsync, uintptr(unsafe.Pointer(this)), NewHStr(webAccountProviderId).Ptr, NewHStr(authority).Ptr, uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2404EEB2-8924-4D93-AB3A-99688B419D56
var IID_IWebAuthenticationCoreManagerStatics3 = syscall.GUID{0x2404EEB2, 0x8924, 0x4D93,
	[8]byte{0xAB, 0x3A, 0x99, 0x68, 0x8B, 0x41, 0x9D, 0x56}}

type IWebAuthenticationCoreManagerStatics3Interface interface {
	win32.IInspectableInterface
	CreateWebAccountMonitor(webAccounts *IIterable[*IWebAccount]) *IWebAccountMonitor
}

type IWebAuthenticationCoreManagerStatics3Vtbl struct {
	win32.IInspectableVtbl
	CreateWebAccountMonitor uintptr
}

type IWebAuthenticationCoreManagerStatics3 struct {
	win32.IInspectable
}

func (this *IWebAuthenticationCoreManagerStatics3) Vtbl() *IWebAuthenticationCoreManagerStatics3Vtbl {
	return (*IWebAuthenticationCoreManagerStatics3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAuthenticationCoreManagerStatics3) CreateWebAccountMonitor(webAccounts *IIterable[*IWebAccount]) *IWebAccountMonitor {
	var _result *IWebAccountMonitor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWebAccountMonitor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(webAccounts)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 54E633FE-96E0-41E8-9832-1298897C2AAF
var IID_IWebAuthenticationCoreManagerStatics4 = syscall.GUID{0x54E633FE, 0x96E0, 0x41E8,
	[8]byte{0x98, 0x32, 0x12, 0x98, 0x89, 0x7C, 0x2A, 0xAF}}

type IWebAuthenticationCoreManagerStatics4Interface interface {
	win32.IInspectableInterface
	FindAllAccountsAsync(provider *IWebAccountProvider) *IAsyncOperation[*IFindAllAccountsResult]
	FindAllAccountsWithClientIdAsync(provider *IWebAccountProvider, clientId string) *IAsyncOperation[*IFindAllAccountsResult]
	FindSystemAccountProviderAsync(webAccountProviderId string) *IAsyncOperation[*IWebAccountProvider]
	FindSystemAccountProviderWithAuthorityAsync(webAccountProviderId string, authority string) *IAsyncOperation[*IWebAccountProvider]
	FindSystemAccountProviderWithAuthorityForUserAsync(webAccountProviderId string, authority string, user *IUser) *IAsyncOperation[*IWebAccountProvider]
}

type IWebAuthenticationCoreManagerStatics4Vtbl struct {
	win32.IInspectableVtbl
	FindAllAccountsAsync                               uintptr
	FindAllAccountsWithClientIdAsync                   uintptr
	FindSystemAccountProviderAsync                     uintptr
	FindSystemAccountProviderWithAuthorityAsync        uintptr
	FindSystemAccountProviderWithAuthorityForUserAsync uintptr
}

type IWebAuthenticationCoreManagerStatics4 struct {
	win32.IInspectable
}

func (this *IWebAuthenticationCoreManagerStatics4) Vtbl() *IWebAuthenticationCoreManagerStatics4Vtbl {
	return (*IWebAuthenticationCoreManagerStatics4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAuthenticationCoreManagerStatics4) FindAllAccountsAsync(provider *IWebAccountProvider) *IAsyncOperation[*IFindAllAccountsResult] {
	var _result *IAsyncOperation[*IFindAllAccountsResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllAccountsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(provider)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAuthenticationCoreManagerStatics4) FindAllAccountsWithClientIdAsync(provider *IWebAccountProvider, clientId string) *IAsyncOperation[*IFindAllAccountsResult] {
	var _result *IAsyncOperation[*IFindAllAccountsResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllAccountsWithClientIdAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(provider)), NewHStr(clientId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAuthenticationCoreManagerStatics4) FindSystemAccountProviderAsync(webAccountProviderId string) *IAsyncOperation[*IWebAccountProvider] {
	var _result *IAsyncOperation[*IWebAccountProvider]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindSystemAccountProviderAsync, uintptr(unsafe.Pointer(this)), NewHStr(webAccountProviderId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAuthenticationCoreManagerStatics4) FindSystemAccountProviderWithAuthorityAsync(webAccountProviderId string, authority string) *IAsyncOperation[*IWebAccountProvider] {
	var _result *IAsyncOperation[*IWebAccountProvider]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindSystemAccountProviderWithAuthorityAsync, uintptr(unsafe.Pointer(this)), NewHStr(webAccountProviderId).Ptr, NewHStr(authority).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAuthenticationCoreManagerStatics4) FindSystemAccountProviderWithAuthorityForUserAsync(webAccountProviderId string, authority string, user *IUser) *IAsyncOperation[*IWebAccountProvider] {
	var _result *IAsyncOperation[*IWebAccountProvider]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindSystemAccountProviderWithAuthorityForUserAsync, uintptr(unsafe.Pointer(this)), NewHStr(webAccountProviderId).Ptr, NewHStr(authority).Ptr, uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DB191BB1-50C5-4809-8DCA-09C99410245C
var IID_IWebProviderError = syscall.GUID{0xDB191BB1, 0x50C5, 0x4809,
	[8]byte{0x8D, 0xCA, 0x09, 0xC9, 0x94, 0x10, 0x24, 0x5C}}

type IWebProviderErrorInterface interface {
	win32.IInspectableInterface
	Get_ErrorCode() uint32
	Get_ErrorMessage() string
	Get_Properties() *IMap[string, string]
}

type IWebProviderErrorVtbl struct {
	win32.IInspectableVtbl
	Get_ErrorCode    uintptr
	Get_ErrorMessage uintptr
	Get_Properties   uintptr
}

type IWebProviderError struct {
	win32.IInspectable
}

func (this *IWebProviderError) Vtbl() *IWebProviderErrorVtbl {
	return (*IWebProviderErrorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebProviderError) Get_ErrorCode() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ErrorCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWebProviderError) Get_ErrorMessage() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ErrorMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWebProviderError) Get_Properties() *IMap[string, string] {
	var _result *IMap[string, string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E3C40A2D-89EF-4E37-847F-A8B9D5A32910
var IID_IWebProviderErrorFactory = syscall.GUID{0xE3C40A2D, 0x89EF, 0x4E37,
	[8]byte{0x84, 0x7F, 0xA8, 0xB9, 0xD5, 0xA3, 0x29, 0x10}}

type IWebProviderErrorFactoryInterface interface {
	win32.IInspectableInterface
	Create(errorCode uint32, errorMessage string) *IWebProviderError
}

type IWebProviderErrorFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IWebProviderErrorFactory struct {
	win32.IInspectable
}

func (this *IWebProviderErrorFactory) Vtbl() *IWebProviderErrorFactoryVtbl {
	return (*IWebProviderErrorFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebProviderErrorFactory) Create(errorCode uint32, errorMessage string) *IWebProviderError {
	var _result *IWebProviderError
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(errorCode), NewHStr(errorMessage).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B77B4D68-ADCB-4673-B364-0CF7B35CAF97
var IID_IWebTokenRequest = syscall.GUID{0xB77B4D68, 0xADCB, 0x4673,
	[8]byte{0xB3, 0x64, 0x0C, 0xF7, 0xB3, 0x5C, 0xAF, 0x97}}

type IWebTokenRequestInterface interface {
	win32.IInspectableInterface
	Get_WebAccountProvider() *IWebAccountProvider
	Get_Scope() string
	Get_ClientId() string
	Get_PromptType() WebTokenRequestPromptType
	Get_Properties() *IMap[string, string]
}

type IWebTokenRequestVtbl struct {
	win32.IInspectableVtbl
	Get_WebAccountProvider uintptr
	Get_Scope              uintptr
	Get_ClientId           uintptr
	Get_PromptType         uintptr
	Get_Properties         uintptr
}

type IWebTokenRequest struct {
	win32.IInspectable
}

func (this *IWebTokenRequest) Vtbl() *IWebTokenRequestVtbl {
	return (*IWebTokenRequestVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebTokenRequest) Get_WebAccountProvider() *IWebAccountProvider {
	var _result *IWebAccountProvider
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WebAccountProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebTokenRequest) Get_Scope() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Scope, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWebTokenRequest) Get_ClientId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ClientId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWebTokenRequest) Get_PromptType() WebTokenRequestPromptType {
	var _result WebTokenRequestPromptType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PromptType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWebTokenRequest) Get_Properties() *IMap[string, string] {
	var _result *IMap[string, string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D700C079-30C8-4397-9654-961C3BE8B855
var IID_IWebTokenRequest2 = syscall.GUID{0xD700C079, 0x30C8, 0x4397,
	[8]byte{0x96, 0x54, 0x96, 0x1C, 0x3B, 0xE8, 0xB8, 0x55}}

type IWebTokenRequest2Interface interface {
	win32.IInspectableInterface
	Get_AppProperties() *IMap[string, string]
}

type IWebTokenRequest2Vtbl struct {
	win32.IInspectableVtbl
	Get_AppProperties uintptr
}

type IWebTokenRequest2 struct {
	win32.IInspectable
}

func (this *IWebTokenRequest2) Vtbl() *IWebTokenRequest2Vtbl {
	return (*IWebTokenRequest2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebTokenRequest2) Get_AppProperties() *IMap[string, string] {
	var _result *IMap[string, string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5A755B51-3BB1-41A5-A63D-90BC32C7DB9A
var IID_IWebTokenRequest3 = syscall.GUID{0x5A755B51, 0x3BB1, 0x41A5,
	[8]byte{0xA6, 0x3D, 0x90, 0xBC, 0x32, 0xC7, 0xDB, 0x9A}}

type IWebTokenRequest3Interface interface {
	win32.IInspectableInterface
	Get_CorrelationId() string
	Put_CorrelationId(value string)
}

type IWebTokenRequest3Vtbl struct {
	win32.IInspectableVtbl
	Get_CorrelationId uintptr
	Put_CorrelationId uintptr
}

type IWebTokenRequest3 struct {
	win32.IInspectable
}

func (this *IWebTokenRequest3) Vtbl() *IWebTokenRequest3Vtbl {
	return (*IWebTokenRequest3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebTokenRequest3) Get_CorrelationId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CorrelationId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWebTokenRequest3) Put_CorrelationId(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CorrelationId, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// 6CF2141C-0FF0-4C67-B84F-99DDBE4A72C9
var IID_IWebTokenRequestFactory = syscall.GUID{0x6CF2141C, 0x0FF0, 0x4C67,
	[8]byte{0xB8, 0x4F, 0x99, 0xDD, 0xBE, 0x4A, 0x72, 0xC9}}

type IWebTokenRequestFactoryInterface interface {
	win32.IInspectableInterface
	Create(provider *IWebAccountProvider, scope string, clientId string) *IWebTokenRequest
	CreateWithPromptType(provider *IWebAccountProvider, scope string, clientId string, promptType WebTokenRequestPromptType) *IWebTokenRequest
	CreateWithProvider(provider *IWebAccountProvider) *IWebTokenRequest
	CreateWithScope(provider *IWebAccountProvider, scope string) *IWebTokenRequest
}

type IWebTokenRequestFactoryVtbl struct {
	win32.IInspectableVtbl
	Create               uintptr
	CreateWithPromptType uintptr
	CreateWithProvider   uintptr
	CreateWithScope      uintptr
}

type IWebTokenRequestFactory struct {
	win32.IInspectable
}

func (this *IWebTokenRequestFactory) Vtbl() *IWebTokenRequestFactoryVtbl {
	return (*IWebTokenRequestFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebTokenRequestFactory) Create(provider *IWebAccountProvider, scope string, clientId string) *IWebTokenRequest {
	var _result *IWebTokenRequest
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(provider)), NewHStr(scope).Ptr, NewHStr(clientId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebTokenRequestFactory) CreateWithPromptType(provider *IWebAccountProvider, scope string, clientId string, promptType WebTokenRequestPromptType) *IWebTokenRequest {
	var _result *IWebTokenRequest
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithPromptType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(provider)), NewHStr(scope).Ptr, NewHStr(clientId).Ptr, uintptr(promptType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebTokenRequestFactory) CreateWithProvider(provider *IWebAccountProvider) *IWebTokenRequest {
	var _result *IWebTokenRequest
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(provider)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebTokenRequestFactory) CreateWithScope(provider *IWebAccountProvider, scope string) *IWebTokenRequest {
	var _result *IWebTokenRequest
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithScope, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(provider)), NewHStr(scope).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C12A8305-D1F8-4483-8D54-38FE292784FF
var IID_IWebTokenRequestResult = syscall.GUID{0xC12A8305, 0xD1F8, 0x4483,
	[8]byte{0x8D, 0x54, 0x38, 0xFE, 0x29, 0x27, 0x84, 0xFF}}

type IWebTokenRequestResultInterface interface {
	win32.IInspectableInterface
	Get_ResponseData() *IVectorView[*IWebTokenResponse]
	Get_ResponseStatus() WebTokenRequestStatus
	Get_ResponseError() *IWebProviderError
	InvalidateCacheAsync() *IAsyncAction
}

type IWebTokenRequestResultVtbl struct {
	win32.IInspectableVtbl
	Get_ResponseData     uintptr
	Get_ResponseStatus   uintptr
	Get_ResponseError    uintptr
	InvalidateCacheAsync uintptr
}

type IWebTokenRequestResult struct {
	win32.IInspectable
}

func (this *IWebTokenRequestResult) Vtbl() *IWebTokenRequestResultVtbl {
	return (*IWebTokenRequestResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebTokenRequestResult) Get_ResponseData() *IVectorView[*IWebTokenResponse] {
	var _result *IVectorView[*IWebTokenResponse]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResponseData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebTokenRequestResult) Get_ResponseStatus() WebTokenRequestStatus {
	var _result WebTokenRequestStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResponseStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWebTokenRequestResult) Get_ResponseError() *IWebProviderError {
	var _result *IWebProviderError
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResponseError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebTokenRequestResult) InvalidateCacheAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InvalidateCacheAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 67A7C5CA-83F6-44C6-A3B1-0EB69E41FA8A
var IID_IWebTokenResponse = syscall.GUID{0x67A7C5CA, 0x83F6, 0x44C6,
	[8]byte{0xA3, 0xB1, 0x0E, 0xB6, 0x9E, 0x41, 0xFA, 0x8A}}

type IWebTokenResponseInterface interface {
	win32.IInspectableInterface
	Get_Token() string
	Get_ProviderError() *IWebProviderError
	Get_WebAccount() *IWebAccount
	Get_Properties() *IMap[string, string]
}

type IWebTokenResponseVtbl struct {
	win32.IInspectableVtbl
	Get_Token         uintptr
	Get_ProviderError uintptr
	Get_WebAccount    uintptr
	Get_Properties    uintptr
}

type IWebTokenResponse struct {
	win32.IInspectable
}

func (this *IWebTokenResponse) Vtbl() *IWebTokenResponseVtbl {
	return (*IWebTokenResponseVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebTokenResponse) Get_Token() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Token, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWebTokenResponse) Get_ProviderError() *IWebProviderError {
	var _result *IWebProviderError
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProviderError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebTokenResponse) Get_WebAccount() *IWebAccount {
	var _result *IWebAccount
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WebAccount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebTokenResponse) Get_Properties() *IMap[string, string] {
	var _result *IMap[string, string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AB6BF7F8-5450-4EF6-97F7-052B0431C0F0
var IID_IWebTokenResponseFactory = syscall.GUID{0xAB6BF7F8, 0x5450, 0x4EF6,
	[8]byte{0x97, 0xF7, 0x05, 0x2B, 0x04, 0x31, 0xC0, 0xF0}}

type IWebTokenResponseFactoryInterface interface {
	win32.IInspectableInterface
	CreateWithToken(token string) *IWebTokenResponse
	CreateWithTokenAndAccount(token string, webAccount *IWebAccount) *IWebTokenResponse
	CreateWithTokenAccountAndError(token string, webAccount *IWebAccount, error *IWebProviderError) *IWebTokenResponse
}

type IWebTokenResponseFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateWithToken                uintptr
	CreateWithTokenAndAccount      uintptr
	CreateWithTokenAccountAndError uintptr
}

type IWebTokenResponseFactory struct {
	win32.IInspectable
}

func (this *IWebTokenResponseFactory) Vtbl() *IWebTokenResponseFactoryVtbl {
	return (*IWebTokenResponseFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebTokenResponseFactory) CreateWithToken(token string) *IWebTokenResponse {
	var _result *IWebTokenResponse
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithToken, uintptr(unsafe.Pointer(this)), NewHStr(token).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebTokenResponseFactory) CreateWithTokenAndAccount(token string, webAccount *IWebAccount) *IWebTokenResponse {
	var _result *IWebTokenResponse
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithTokenAndAccount, uintptr(unsafe.Pointer(this)), NewHStr(token).Ptr, uintptr(unsafe.Pointer(webAccount)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebTokenResponseFactory) CreateWithTokenAccountAndError(token string, webAccount *IWebAccount, error *IWebProviderError) *IWebTokenResponse {
	var _result *IWebTokenResponse
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithTokenAccountAndError, uintptr(unsafe.Pointer(this)), NewHStr(token).Ptr, uintptr(unsafe.Pointer(webAccount)), uintptr(unsafe.Pointer(error)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type WebAccountEventArgs struct {
	RtClass
	*IWebAccountEventArgs
}

type WebAccountMonitor struct {
	RtClass
	*IWebAccountMonitor
}

type WebAuthenticationCoreManager struct {
	RtClass
}

func NewIWebAuthenticationCoreManagerStatics2() *IWebAuthenticationCoreManagerStatics2 {
	var p *IWebAuthenticationCoreManagerStatics2
	hs := NewHStr("Windows.Security.Authentication.Web.Core.WebAuthenticationCoreManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWebAuthenticationCoreManagerStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIWebAuthenticationCoreManagerStatics4() *IWebAuthenticationCoreManagerStatics4 {
	var p *IWebAuthenticationCoreManagerStatics4
	hs := NewHStr("Windows.Security.Authentication.Web.Core.WebAuthenticationCoreManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWebAuthenticationCoreManagerStatics4, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIWebAuthenticationCoreManagerStatics() *IWebAuthenticationCoreManagerStatics {
	var p *IWebAuthenticationCoreManagerStatics
	hs := NewHStr("Windows.Security.Authentication.Web.Core.WebAuthenticationCoreManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWebAuthenticationCoreManagerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIWebAuthenticationCoreManagerStatics3() *IWebAuthenticationCoreManagerStatics3 {
	var p *IWebAuthenticationCoreManagerStatics3
	hs := NewHStr("Windows.Security.Authentication.Web.Core.WebAuthenticationCoreManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWebAuthenticationCoreManagerStatics3, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type WebProviderError struct {
	RtClass
	*IWebProviderError
}

func NewWebProviderError_Create(errorCode uint32, errorMessage string) *WebProviderError {
	hs := NewHStr("Windows.Security.Authentication.Web.Core.WebProviderError")
	var pFac *IWebProviderErrorFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWebProviderErrorFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IWebProviderError
	p = pFac.Create(errorCode, errorMessage)
	result := &WebProviderError{
		RtClass:           RtClass{PInspect: &p.IInspectable},
		IWebProviderError: p,
	}
	com.AddToScope(result)
	return result
}

type WebTokenRequest struct {
	RtClass
	*IWebTokenRequest
}

func NewWebTokenRequest_Create(provider *IWebAccountProvider, scope string, clientId string) *WebTokenRequest {
	hs := NewHStr("Windows.Security.Authentication.Web.Core.WebTokenRequest")
	var pFac *IWebTokenRequestFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWebTokenRequestFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IWebTokenRequest
	p = pFac.Create(provider, scope, clientId)
	result := &WebTokenRequest{
		RtClass:          RtClass{PInspect: &p.IInspectable},
		IWebTokenRequest: p,
	}
	com.AddToScope(result)
	return result
}

func NewWebTokenRequest_CreateWithPromptType(provider *IWebAccountProvider, scope string, clientId string, promptType WebTokenRequestPromptType) *WebTokenRequest {
	hs := NewHStr("Windows.Security.Authentication.Web.Core.WebTokenRequest")
	var pFac *IWebTokenRequestFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWebTokenRequestFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IWebTokenRequest
	p = pFac.CreateWithPromptType(provider, scope, clientId, promptType)
	result := &WebTokenRequest{
		RtClass:          RtClass{PInspect: &p.IInspectable},
		IWebTokenRequest: p,
	}
	com.AddToScope(result)
	return result
}

func NewWebTokenRequest_CreateWithProvider(provider *IWebAccountProvider) *WebTokenRequest {
	hs := NewHStr("Windows.Security.Authentication.Web.Core.WebTokenRequest")
	var pFac *IWebTokenRequestFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWebTokenRequestFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IWebTokenRequest
	p = pFac.CreateWithProvider(provider)
	result := &WebTokenRequest{
		RtClass:          RtClass{PInspect: &p.IInspectable},
		IWebTokenRequest: p,
	}
	com.AddToScope(result)
	return result
}

func NewWebTokenRequest_CreateWithScope(provider *IWebAccountProvider, scope string) *WebTokenRequest {
	hs := NewHStr("Windows.Security.Authentication.Web.Core.WebTokenRequest")
	var pFac *IWebTokenRequestFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWebTokenRequestFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IWebTokenRequest
	p = pFac.CreateWithScope(provider, scope)
	result := &WebTokenRequest{
		RtClass:          RtClass{PInspect: &p.IInspectable},
		IWebTokenRequest: p,
	}
	com.AddToScope(result)
	return result
}

type WebTokenRequestResult struct {
	RtClass
	*IWebTokenRequestResult
}

type WebTokenResponse struct {
	RtClass
	*IWebTokenResponse
}

func NewWebTokenResponse() *WebTokenResponse {
	hs := NewHStr("Windows.Security.Authentication.Web.Core.WebTokenResponse")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &WebTokenResponse{
		RtClass:           RtClass{PInspect: p},
		IWebTokenResponse: (*IWebTokenResponse)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewWebTokenResponse_CreateWithToken(token string) *WebTokenResponse {
	hs := NewHStr("Windows.Security.Authentication.Web.Core.WebTokenResponse")
	var pFac *IWebTokenResponseFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWebTokenResponseFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IWebTokenResponse
	p = pFac.CreateWithToken(token)
	result := &WebTokenResponse{
		RtClass:           RtClass{PInspect: &p.IInspectable},
		IWebTokenResponse: p,
	}
	com.AddToScope(result)
	return result
}

func NewWebTokenResponse_CreateWithTokenAndAccount(token string, webAccount *IWebAccount) *WebTokenResponse {
	hs := NewHStr("Windows.Security.Authentication.Web.Core.WebTokenResponse")
	var pFac *IWebTokenResponseFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWebTokenResponseFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IWebTokenResponse
	p = pFac.CreateWithTokenAndAccount(token, webAccount)
	result := &WebTokenResponse{
		RtClass:           RtClass{PInspect: &p.IInspectable},
		IWebTokenResponse: p,
	}
	com.AddToScope(result)
	return result
}

func NewWebTokenResponse_CreateWithTokenAccountAndError(token string, webAccount *IWebAccount, error *IWebProviderError) *WebTokenResponse {
	hs := NewHStr("Windows.Security.Authentication.Web.Core.WebTokenResponse")
	var pFac *IWebTokenResponseFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWebTokenResponseFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IWebTokenResponse
	p = pFac.CreateWithTokenAccountAndError(token, webAccount, error)
	result := &WebTokenResponse{
		RtClass:           RtClass{PInspect: &p.IInspectable},
		IWebTokenResponse: p,
	}
	com.AddToScope(result)
	return result
}
