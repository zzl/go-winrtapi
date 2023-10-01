package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type WebErrorStatus int32

const (
	WebErrorStatus_Unknown                          WebErrorStatus = 0
	WebErrorStatus_CertificateCommonNameIsIncorrect WebErrorStatus = 1
	WebErrorStatus_CertificateExpired               WebErrorStatus = 2
	WebErrorStatus_CertificateContainsErrors        WebErrorStatus = 3
	WebErrorStatus_CertificateRevoked               WebErrorStatus = 4
	WebErrorStatus_CertificateIsInvalid             WebErrorStatus = 5
	WebErrorStatus_ServerUnreachable                WebErrorStatus = 6
	WebErrorStatus_Timeout                          WebErrorStatus = 7
	WebErrorStatus_ErrorHttpInvalidServerResponse   WebErrorStatus = 8
	WebErrorStatus_ConnectionAborted                WebErrorStatus = 9
	WebErrorStatus_ConnectionReset                  WebErrorStatus = 10
	WebErrorStatus_Disconnected                     WebErrorStatus = 11
	WebErrorStatus_HttpToHttpsOnRedirection         WebErrorStatus = 12
	WebErrorStatus_HttpsToHttpOnRedirection         WebErrorStatus = 13
	WebErrorStatus_CannotConnect                    WebErrorStatus = 14
	WebErrorStatus_HostNameNotResolved              WebErrorStatus = 15
	WebErrorStatus_OperationCanceled                WebErrorStatus = 16
	WebErrorStatus_RedirectFailed                   WebErrorStatus = 17
	WebErrorStatus_UnexpectedStatusCode             WebErrorStatus = 18
	WebErrorStatus_UnexpectedRedirection            WebErrorStatus = 19
	WebErrorStatus_UnexpectedClientError            WebErrorStatus = 20
	WebErrorStatus_UnexpectedServerError            WebErrorStatus = 21
	WebErrorStatus_InsufficientRangeSupport         WebErrorStatus = 22
	WebErrorStatus_MissingContentLengthSupport      WebErrorStatus = 23
	WebErrorStatus_MultipleChoices                  WebErrorStatus = 300
	WebErrorStatus_MovedPermanently                 WebErrorStatus = 301
	WebErrorStatus_Found                            WebErrorStatus = 302
	WebErrorStatus_SeeOther                         WebErrorStatus = 303
	WebErrorStatus_NotModified                      WebErrorStatus = 304
	WebErrorStatus_UseProxy                         WebErrorStatus = 305
	WebErrorStatus_TemporaryRedirect                WebErrorStatus = 307
	WebErrorStatus_BadRequest                       WebErrorStatus = 400
	WebErrorStatus_Unauthorized                     WebErrorStatus = 401
	WebErrorStatus_PaymentRequired                  WebErrorStatus = 402
	WebErrorStatus_Forbidden                        WebErrorStatus = 403
	WebErrorStatus_NotFound                         WebErrorStatus = 404
	WebErrorStatus_MethodNotAllowed                 WebErrorStatus = 405
	WebErrorStatus_NotAcceptable                    WebErrorStatus = 406
	WebErrorStatus_ProxyAuthenticationRequired      WebErrorStatus = 407
	WebErrorStatus_RequestTimeout                   WebErrorStatus = 408
	WebErrorStatus_Conflict                         WebErrorStatus = 409
	WebErrorStatus_Gone                             WebErrorStatus = 410
	WebErrorStatus_LengthRequired                   WebErrorStatus = 411
	WebErrorStatus_PreconditionFailed               WebErrorStatus = 412
	WebErrorStatus_RequestEntityTooLarge            WebErrorStatus = 413
	WebErrorStatus_RequestUriTooLong                WebErrorStatus = 414
	WebErrorStatus_UnsupportedMediaType             WebErrorStatus = 415
	WebErrorStatus_RequestedRangeNotSatisfiable     WebErrorStatus = 416
	WebErrorStatus_ExpectationFailed                WebErrorStatus = 417
	WebErrorStatus_InternalServerError              WebErrorStatus = 500
	WebErrorStatus_NotImplemented                   WebErrorStatus = 501
	WebErrorStatus_BadGateway                       WebErrorStatus = 502
	WebErrorStatus_ServiceUnavailable               WebErrorStatus = 503
	WebErrorStatus_GatewayTimeout                   WebErrorStatus = 504
	WebErrorStatus_HttpVersionNotSupported          WebErrorStatus = 505
)

// interfaces

// B0ABA86A-9AEB-4D3A-9590-003E3CA7E290
var IID_IUriToStreamResolver = syscall.GUID{0xB0ABA86A, 0x9AEB, 0x4D3A,
	[8]byte{0x95, 0x90, 0x00, 0x3E, 0x3C, 0xA7, 0xE2, 0x90}}

type IUriToStreamResolverInterface interface {
	win32.IInspectableInterface
	UriToStreamAsync(uri *IUriRuntimeClass) *IAsyncOperation[*IInputStream]
}

type IUriToStreamResolverVtbl struct {
	win32.IInspectableVtbl
	UriToStreamAsync uintptr
}

type IUriToStreamResolver struct {
	win32.IInspectable
}

func (this *IUriToStreamResolver) Vtbl() *IUriToStreamResolverVtbl {
	return (*IUriToStreamResolverVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUriToStreamResolver) UriToStreamAsync(uri *IUriRuntimeClass) *IAsyncOperation[*IInputStream] {
	var _result *IAsyncOperation[*IInputStream]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UriToStreamAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FE616766-BF27-4064-87B7-6563BB11CE2E
var IID_IWebErrorStatics = syscall.GUID{0xFE616766, 0xBF27, 0x4064,
	[8]byte{0x87, 0xB7, 0x65, 0x63, 0xBB, 0x11, 0xCE, 0x2E}}

type IWebErrorStaticsInterface interface {
	win32.IInspectableInterface
	GetStatus(hresult int32) WebErrorStatus
}

type IWebErrorStaticsVtbl struct {
	win32.IInspectableVtbl
	GetStatus uintptr
}

type IWebErrorStatics struct {
	win32.IInspectable
}

func (this *IWebErrorStatics) Vtbl() *IWebErrorStaticsVtbl {
	return (*IWebErrorStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebErrorStatics) GetStatus(hresult int32) WebErrorStatus {
	var _result WebErrorStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStatus, uintptr(unsafe.Pointer(this)), uintptr(hresult), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type WebError struct {
	RtClass
}

func NewIWebErrorStatics() *IWebErrorStatics {
	var p *IWebErrorStatics
	hs := NewHStr("Windows.Web.WebError")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWebErrorStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
