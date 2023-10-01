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
type HttpCacheReadBehavior int32

const (
	HttpCacheReadBehavior_Default       HttpCacheReadBehavior = 0
	HttpCacheReadBehavior_MostRecent    HttpCacheReadBehavior = 1
	HttpCacheReadBehavior_OnlyFromCache HttpCacheReadBehavior = 2
	HttpCacheReadBehavior_NoCache       HttpCacheReadBehavior = 3
)

// enum
type HttpCacheWriteBehavior int32

const (
	HttpCacheWriteBehavior_Default HttpCacheWriteBehavior = 0
	HttpCacheWriteBehavior_NoCache HttpCacheWriteBehavior = 1
)

// enum
type HttpCookieUsageBehavior int32

const (
	HttpCookieUsageBehavior_Default   HttpCookieUsageBehavior = 0
	HttpCookieUsageBehavior_NoCookies HttpCookieUsageBehavior = 1
)

// interfaces

// 71C89B09-E131-4B54-A53C-EB43FF37E9BB
var IID_IHttpBaseProtocolFilter = syscall.GUID{0x71C89B09, 0xE131, 0x4B54,
	[8]byte{0xA5, 0x3C, 0xEB, 0x43, 0xFF, 0x37, 0xE9, 0xBB}}

type IHttpBaseProtocolFilterInterface interface {
	win32.IInspectableInterface
	Get_AllowAutoRedirect() bool
	Put_AllowAutoRedirect(value bool)
	Get_AllowUI() bool
	Put_AllowUI(value bool)
	Get_AutomaticDecompression() bool
	Put_AutomaticDecompression(value bool)
	Get_CacheControl() *IHttpCacheControl
	Get_CookieManager() *IHttpCookieManager
	Get_ClientCertificate() *ICertificate
	Put_ClientCertificate(value *ICertificate)
	Get_IgnorableServerCertificateErrors() *IVector[ChainValidationResult]
	Get_MaxConnectionsPerServer() uint32
	Put_MaxConnectionsPerServer(value uint32)
	Get_ProxyCredential() *IPasswordCredential
	Put_ProxyCredential(value *IPasswordCredential)
	Get_ServerCredential() *IPasswordCredential
	Put_ServerCredential(value *IPasswordCredential)
	Get_UseProxy() bool
	Put_UseProxy(value bool)
}

type IHttpBaseProtocolFilterVtbl struct {
	win32.IInspectableVtbl
	Get_AllowAutoRedirect                uintptr
	Put_AllowAutoRedirect                uintptr
	Get_AllowUI                          uintptr
	Put_AllowUI                          uintptr
	Get_AutomaticDecompression           uintptr
	Put_AutomaticDecompression           uintptr
	Get_CacheControl                     uintptr
	Get_CookieManager                    uintptr
	Get_ClientCertificate                uintptr
	Put_ClientCertificate                uintptr
	Get_IgnorableServerCertificateErrors uintptr
	Get_MaxConnectionsPerServer          uintptr
	Put_MaxConnectionsPerServer          uintptr
	Get_ProxyCredential                  uintptr
	Put_ProxyCredential                  uintptr
	Get_ServerCredential                 uintptr
	Put_ServerCredential                 uintptr
	Get_UseProxy                         uintptr
	Put_UseProxy                         uintptr
}

type IHttpBaseProtocolFilter struct {
	win32.IInspectable
}

func (this *IHttpBaseProtocolFilter) Vtbl() *IHttpBaseProtocolFilterVtbl {
	return (*IHttpBaseProtocolFilterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpBaseProtocolFilter) Get_AllowAutoRedirect() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllowAutoRedirect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHttpBaseProtocolFilter) Put_AllowAutoRedirect(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AllowAutoRedirect, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IHttpBaseProtocolFilter) Get_AllowUI() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllowUI, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHttpBaseProtocolFilter) Put_AllowUI(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AllowUI, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IHttpBaseProtocolFilter) Get_AutomaticDecompression() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AutomaticDecompression, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHttpBaseProtocolFilter) Put_AutomaticDecompression(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AutomaticDecompression, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IHttpBaseProtocolFilter) Get_CacheControl() *IHttpCacheControl {
	var _result *IHttpCacheControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CacheControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpBaseProtocolFilter) Get_CookieManager() *IHttpCookieManager {
	var _result *IHttpCookieManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CookieManager, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpBaseProtocolFilter) Get_ClientCertificate() *ICertificate {
	var _result *ICertificate
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ClientCertificate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpBaseProtocolFilter) Put_ClientCertificate(value *ICertificate) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ClientCertificate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpBaseProtocolFilter) Get_IgnorableServerCertificateErrors() *IVector[ChainValidationResult] {
	var _result *IVector[ChainValidationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IgnorableServerCertificateErrors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpBaseProtocolFilter) Get_MaxConnectionsPerServer() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxConnectionsPerServer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHttpBaseProtocolFilter) Put_MaxConnectionsPerServer(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxConnectionsPerServer, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IHttpBaseProtocolFilter) Get_ProxyCredential() *IPasswordCredential {
	var _result *IPasswordCredential
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProxyCredential, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpBaseProtocolFilter) Put_ProxyCredential(value *IPasswordCredential) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ProxyCredential, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpBaseProtocolFilter) Get_ServerCredential() *IPasswordCredential {
	var _result *IPasswordCredential
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServerCredential, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpBaseProtocolFilter) Put_ServerCredential(value *IPasswordCredential) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ServerCredential, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpBaseProtocolFilter) Get_UseProxy() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UseProxy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHttpBaseProtocolFilter) Put_UseProxy(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_UseProxy, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 2EC30013-9427-4900-A017-FA7DA3B5C9AE
var IID_IHttpBaseProtocolFilter2 = syscall.GUID{0x2EC30013, 0x9427, 0x4900,
	[8]byte{0xA0, 0x17, 0xFA, 0x7D, 0xA3, 0xB5, 0xC9, 0xAE}}

type IHttpBaseProtocolFilter2Interface interface {
	win32.IInspectableInterface
	Get_MaxVersion() HttpVersion
	Put_MaxVersion(value HttpVersion)
}

type IHttpBaseProtocolFilter2Vtbl struct {
	win32.IInspectableVtbl
	Get_MaxVersion uintptr
	Put_MaxVersion uintptr
}

type IHttpBaseProtocolFilter2 struct {
	win32.IInspectable
}

func (this *IHttpBaseProtocolFilter2) Vtbl() *IHttpBaseProtocolFilter2Vtbl {
	return (*IHttpBaseProtocolFilter2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpBaseProtocolFilter2) Get_MaxVersion() HttpVersion {
	var _result HttpVersion
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxVersion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHttpBaseProtocolFilter2) Put_MaxVersion(value HttpVersion) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxVersion, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// D43F4D4C-BD42-43AE-8717-AD2C8F4B2937
var IID_IHttpBaseProtocolFilter3 = syscall.GUID{0xD43F4D4C, 0xBD42, 0x43AE,
	[8]byte{0x87, 0x17, 0xAD, 0x2C, 0x8F, 0x4B, 0x29, 0x37}}

type IHttpBaseProtocolFilter3Interface interface {
	win32.IInspectableInterface
	Get_CookieUsageBehavior() HttpCookieUsageBehavior
	Put_CookieUsageBehavior(value HttpCookieUsageBehavior)
}

type IHttpBaseProtocolFilter3Vtbl struct {
	win32.IInspectableVtbl
	Get_CookieUsageBehavior uintptr
	Put_CookieUsageBehavior uintptr
}

type IHttpBaseProtocolFilter3 struct {
	win32.IInspectable
}

func (this *IHttpBaseProtocolFilter3) Vtbl() *IHttpBaseProtocolFilter3Vtbl {
	return (*IHttpBaseProtocolFilter3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpBaseProtocolFilter3) Get_CookieUsageBehavior() HttpCookieUsageBehavior {
	var _result HttpCookieUsageBehavior
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CookieUsageBehavior, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHttpBaseProtocolFilter3) Put_CookieUsageBehavior(value HttpCookieUsageBehavior) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CookieUsageBehavior, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 9FE36CCF-2983-4893-941F-EB518CA8CEF9
var IID_IHttpBaseProtocolFilter4 = syscall.GUID{0x9FE36CCF, 0x2983, 0x4893,
	[8]byte{0x94, 0x1F, 0xEB, 0x51, 0x8C, 0xA8, 0xCE, 0xF9}}

type IHttpBaseProtocolFilter4Interface interface {
	win32.IInspectableInterface
	Add_ServerCustomValidationRequested(handler TypedEventHandler[*IHttpBaseProtocolFilter, *IHttpServerCustomValidationRequestedEventArgs]) EventRegistrationToken
	Remove_ServerCustomValidationRequested(token EventRegistrationToken)
	ClearAuthenticationCache()
}

type IHttpBaseProtocolFilter4Vtbl struct {
	win32.IInspectableVtbl
	Add_ServerCustomValidationRequested    uintptr
	Remove_ServerCustomValidationRequested uintptr
	ClearAuthenticationCache               uintptr
}

type IHttpBaseProtocolFilter4 struct {
	win32.IInspectable
}

func (this *IHttpBaseProtocolFilter4) Vtbl() *IHttpBaseProtocolFilter4Vtbl {
	return (*IHttpBaseProtocolFilter4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpBaseProtocolFilter4) Add_ServerCustomValidationRequested(handler TypedEventHandler[*IHttpBaseProtocolFilter, *IHttpServerCustomValidationRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ServerCustomValidationRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHttpBaseProtocolFilter4) Remove_ServerCustomValidationRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ServerCustomValidationRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IHttpBaseProtocolFilter4) ClearAuthenticationCache() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearAuthenticationCache, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 416E4993-31E3-4816-BF09-E018EE8DC1F5
var IID_IHttpBaseProtocolFilter5 = syscall.GUID{0x416E4993, 0x31E3, 0x4816,
	[8]byte{0xBF, 0x09, 0xE0, 0x18, 0xEE, 0x8D, 0xC1, 0xF5}}

type IHttpBaseProtocolFilter5Interface interface {
	win32.IInspectableInterface
	Get_User() *IUser
}

type IHttpBaseProtocolFilter5Vtbl struct {
	win32.IInspectableVtbl
	Get_User uintptr
}

type IHttpBaseProtocolFilter5 struct {
	win32.IInspectable
}

func (this *IHttpBaseProtocolFilter5) Vtbl() *IHttpBaseProtocolFilter5Vtbl {
	return (*IHttpBaseProtocolFilter5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpBaseProtocolFilter5) Get_User() *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_User, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6D4DEE0C-E908-494E-B5A3-1263C9B8242A
var IID_IHttpBaseProtocolFilterStatics = syscall.GUID{0x6D4DEE0C, 0xE908, 0x494E,
	[8]byte{0xB5, 0xA3, 0x12, 0x63, 0xC9, 0xB8, 0x24, 0x2A}}

type IHttpBaseProtocolFilterStaticsInterface interface {
	win32.IInspectableInterface
	CreateForUser(user *IUser) *IHttpBaseProtocolFilter
}

type IHttpBaseProtocolFilterStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateForUser uintptr
}

type IHttpBaseProtocolFilterStatics struct {
	win32.IInspectable
}

func (this *IHttpBaseProtocolFilterStatics) Vtbl() *IHttpBaseProtocolFilterStaticsVtbl {
	return (*IHttpBaseProtocolFilterStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpBaseProtocolFilterStatics) CreateForUser(user *IUser) *IHttpBaseProtocolFilter {
	var _result *IHttpBaseProtocolFilter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateForUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C77E1CB4-3CEA-4EB5-AC85-04E186E63AB7
var IID_IHttpCacheControl = syscall.GUID{0xC77E1CB4, 0x3CEA, 0x4EB5,
	[8]byte{0xAC, 0x85, 0x04, 0xE1, 0x86, 0xE6, 0x3A, 0xB7}}

type IHttpCacheControlInterface interface {
	win32.IInspectableInterface
	Get_ReadBehavior() HttpCacheReadBehavior
	Put_ReadBehavior(value HttpCacheReadBehavior)
	Get_WriteBehavior() HttpCacheWriteBehavior
	Put_WriteBehavior(value HttpCacheWriteBehavior)
}

type IHttpCacheControlVtbl struct {
	win32.IInspectableVtbl
	Get_ReadBehavior  uintptr
	Put_ReadBehavior  uintptr
	Get_WriteBehavior uintptr
	Put_WriteBehavior uintptr
}

type IHttpCacheControl struct {
	win32.IInspectable
}

func (this *IHttpCacheControl) Vtbl() *IHttpCacheControlVtbl {
	return (*IHttpCacheControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpCacheControl) Get_ReadBehavior() HttpCacheReadBehavior {
	var _result HttpCacheReadBehavior
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReadBehavior, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHttpCacheControl) Put_ReadBehavior(value HttpCacheReadBehavior) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReadBehavior, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IHttpCacheControl) Get_WriteBehavior() HttpCacheWriteBehavior {
	var _result HttpCacheWriteBehavior
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WriteBehavior, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHttpCacheControl) Put_WriteBehavior(value HttpCacheWriteBehavior) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_WriteBehavior, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// A4CB6DD5-0902-439E-BFD7-E12552B165CE
var IID_IHttpFilter = syscall.GUID{0xA4CB6DD5, 0x0902, 0x439E,
	[8]byte{0xBF, 0xD7, 0xE1, 0x25, 0x52, 0xB1, 0x65, 0xCE}}

type IHttpFilterInterface interface {
	win32.IInspectableInterface
	SendRequestAsync(request *IHttpRequestMessage) *IAsyncOperationWithProgress[*IHttpResponseMessage, HttpProgress]
}

type IHttpFilterVtbl struct {
	win32.IInspectableVtbl
	SendRequestAsync uintptr
}

type IHttpFilter struct {
	win32.IInspectable
}

func (this *IHttpFilter) Vtbl() *IHttpFilterVtbl {
	return (*IHttpFilterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpFilter) SendRequestAsync(request *IHttpRequestMessage) *IAsyncOperationWithProgress[*IHttpResponseMessage, HttpProgress] {
	var _result *IAsyncOperationWithProgress[*IHttpResponseMessage, HttpProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendRequestAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(request)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3165FE32-E7DD-48B7-A361-939C750E63CC
var IID_IHttpServerCustomValidationRequestedEventArgs = syscall.GUID{0x3165FE32, 0xE7DD, 0x48B7,
	[8]byte{0xA3, 0x61, 0x93, 0x9C, 0x75, 0x0E, 0x63, 0xCC}}

type IHttpServerCustomValidationRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_RequestMessage() *IHttpRequestMessage
	Get_ServerCertificate() *ICertificate
	Get_ServerCertificateErrorSeverity() SocketSslErrorSeverity
	Get_ServerCertificateErrors() *IVectorView[ChainValidationResult]
	Get_ServerIntermediateCertificates() *IVectorView[*ICertificate]
	Reject()
	GetDeferral() *IDeferral
}

type IHttpServerCustomValidationRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_RequestMessage                 uintptr
	Get_ServerCertificate              uintptr
	Get_ServerCertificateErrorSeverity uintptr
	Get_ServerCertificateErrors        uintptr
	Get_ServerIntermediateCertificates uintptr
	Reject                             uintptr
	GetDeferral                        uintptr
}

type IHttpServerCustomValidationRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IHttpServerCustomValidationRequestedEventArgs) Vtbl() *IHttpServerCustomValidationRequestedEventArgsVtbl {
	return (*IHttpServerCustomValidationRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpServerCustomValidationRequestedEventArgs) Get_RequestMessage() *IHttpRequestMessage {
	var _result *IHttpRequestMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpServerCustomValidationRequestedEventArgs) Get_ServerCertificate() *ICertificate {
	var _result *ICertificate
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServerCertificate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpServerCustomValidationRequestedEventArgs) Get_ServerCertificateErrorSeverity() SocketSslErrorSeverity {
	var _result SocketSslErrorSeverity
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServerCertificateErrorSeverity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHttpServerCustomValidationRequestedEventArgs) Get_ServerCertificateErrors() *IVectorView[ChainValidationResult] {
	var _result *IVectorView[ChainValidationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServerCertificateErrors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpServerCustomValidationRequestedEventArgs) Get_ServerIntermediateCertificates() *IVectorView[*ICertificate] {
	var _result *IVectorView[*ICertificate]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServerIntermediateCertificates, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpServerCustomValidationRequestedEventArgs) Reject() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Reject, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IHttpServerCustomValidationRequestedEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type HttpBaseProtocolFilter struct {
	RtClass
	*IHttpBaseProtocolFilter
}

func NewHttpBaseProtocolFilter() *HttpBaseProtocolFilter {
	hs := NewHStr("Windows.Web.Http.Filters.HttpBaseProtocolFilter")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &HttpBaseProtocolFilter{
		RtClass:                 RtClass{PInspect: p},
		IHttpBaseProtocolFilter: (*IHttpBaseProtocolFilter)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewIHttpBaseProtocolFilterStatics() *IHttpBaseProtocolFilterStatics {
	var p *IHttpBaseProtocolFilterStatics
	hs := NewHStr("Windows.Web.Http.Filters.HttpBaseProtocolFilter")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpBaseProtocolFilterStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type HttpCacheControl struct {
	RtClass
	*IHttpCacheControl
}

type HttpServerCustomValidationRequestedEventArgs struct {
	RtClass
	*IHttpServerCustomValidationRequestedEventArgs
}
