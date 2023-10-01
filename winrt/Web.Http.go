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
type HttpCompletionOption int32

const (
	HttpCompletionOption_ResponseContentRead HttpCompletionOption = 0
	HttpCompletionOption_ResponseHeadersRead HttpCompletionOption = 1
)

// enum
type HttpProgressStage int32

const (
	HttpProgressStage_None               HttpProgressStage = 0
	HttpProgressStage_DetectingProxy     HttpProgressStage = 10
	HttpProgressStage_ResolvingName      HttpProgressStage = 20
	HttpProgressStage_ConnectingToServer HttpProgressStage = 30
	HttpProgressStage_NegotiatingSsl     HttpProgressStage = 40
	HttpProgressStage_SendingHeaders     HttpProgressStage = 50
	HttpProgressStage_SendingContent     HttpProgressStage = 60
	HttpProgressStage_WaitingForResponse HttpProgressStage = 70
	HttpProgressStage_ReceivingHeaders   HttpProgressStage = 80
	HttpProgressStage_ReceivingContent   HttpProgressStage = 90
)

// enum
type HttpResponseMessageSource int32

const (
	HttpResponseMessageSource_None    HttpResponseMessageSource = 0
	HttpResponseMessageSource_Cache   HttpResponseMessageSource = 1
	HttpResponseMessageSource_Network HttpResponseMessageSource = 2
)

// enum
type HttpStatusCode int32

const (
	HttpStatusCode_None                          HttpStatusCode = 0
	HttpStatusCode_Continue                      HttpStatusCode = 100
	HttpStatusCode_SwitchingProtocols            HttpStatusCode = 101
	HttpStatusCode_Processing                    HttpStatusCode = 102
	HttpStatusCode_Ok                            HttpStatusCode = 200
	HttpStatusCode_Created                       HttpStatusCode = 201
	HttpStatusCode_Accepted                      HttpStatusCode = 202
	HttpStatusCode_NonAuthoritativeInformation   HttpStatusCode = 203
	HttpStatusCode_NoContent                     HttpStatusCode = 204
	HttpStatusCode_ResetContent                  HttpStatusCode = 205
	HttpStatusCode_PartialContent                HttpStatusCode = 206
	HttpStatusCode_MultiStatus                   HttpStatusCode = 207
	HttpStatusCode_AlreadyReported               HttpStatusCode = 208
	HttpStatusCode_IMUsed                        HttpStatusCode = 226
	HttpStatusCode_MultipleChoices               HttpStatusCode = 300
	HttpStatusCode_MovedPermanently              HttpStatusCode = 301
	HttpStatusCode_Found                         HttpStatusCode = 302
	HttpStatusCode_SeeOther                      HttpStatusCode = 303
	HttpStatusCode_NotModified                   HttpStatusCode = 304
	HttpStatusCode_UseProxy                      HttpStatusCode = 305
	HttpStatusCode_TemporaryRedirect             HttpStatusCode = 307
	HttpStatusCode_PermanentRedirect             HttpStatusCode = 308
	HttpStatusCode_BadRequest                    HttpStatusCode = 400
	HttpStatusCode_Unauthorized                  HttpStatusCode = 401
	HttpStatusCode_PaymentRequired               HttpStatusCode = 402
	HttpStatusCode_Forbidden                     HttpStatusCode = 403
	HttpStatusCode_NotFound                      HttpStatusCode = 404
	HttpStatusCode_MethodNotAllowed              HttpStatusCode = 405
	HttpStatusCode_NotAcceptable                 HttpStatusCode = 406
	HttpStatusCode_ProxyAuthenticationRequired   HttpStatusCode = 407
	HttpStatusCode_RequestTimeout                HttpStatusCode = 408
	HttpStatusCode_Conflict                      HttpStatusCode = 409
	HttpStatusCode_Gone                          HttpStatusCode = 410
	HttpStatusCode_LengthRequired                HttpStatusCode = 411
	HttpStatusCode_PreconditionFailed            HttpStatusCode = 412
	HttpStatusCode_RequestEntityTooLarge         HttpStatusCode = 413
	HttpStatusCode_RequestUriTooLong             HttpStatusCode = 414
	HttpStatusCode_UnsupportedMediaType          HttpStatusCode = 415
	HttpStatusCode_RequestedRangeNotSatisfiable  HttpStatusCode = 416
	HttpStatusCode_ExpectationFailed             HttpStatusCode = 417
	HttpStatusCode_UnprocessableEntity           HttpStatusCode = 422
	HttpStatusCode_Locked                        HttpStatusCode = 423
	HttpStatusCode_FailedDependency              HttpStatusCode = 424
	HttpStatusCode_UpgradeRequired               HttpStatusCode = 426
	HttpStatusCode_PreconditionRequired          HttpStatusCode = 428
	HttpStatusCode_TooManyRequests               HttpStatusCode = 429
	HttpStatusCode_RequestHeaderFieldsTooLarge   HttpStatusCode = 431
	HttpStatusCode_InternalServerError           HttpStatusCode = 500
	HttpStatusCode_NotImplemented                HttpStatusCode = 501
	HttpStatusCode_BadGateway                    HttpStatusCode = 502
	HttpStatusCode_ServiceUnavailable            HttpStatusCode = 503
	HttpStatusCode_GatewayTimeout                HttpStatusCode = 504
	HttpStatusCode_HttpVersionNotSupported       HttpStatusCode = 505
	HttpStatusCode_VariantAlsoNegotiates         HttpStatusCode = 506
	HttpStatusCode_InsufficientStorage           HttpStatusCode = 507
	HttpStatusCode_LoopDetected                  HttpStatusCode = 508
	HttpStatusCode_NotExtended                   HttpStatusCode = 510
	HttpStatusCode_NetworkAuthenticationRequired HttpStatusCode = 511
)

// enum
type HttpVersion int32

const (
	HttpVersion_None   HttpVersion = 0
	HttpVersion_Http10 HttpVersion = 1
	HttpVersion_Http11 HttpVersion = 2
	HttpVersion_Http20 HttpVersion = 3
)

// structs

type HttpProgress struct {
	Stage               HttpProgressStage
	BytesSent           uint64
	TotalBytesToSend    *IReference[uint64]
	BytesReceived       uint64
	TotalBytesToReceive *IReference[uint64]
	Retries             uint32
}

// interfaces

// BC20C193-C41F-4FF7-9123-6435736EADC2
var IID_IHttpBufferContentFactory = syscall.GUID{0xBC20C193, 0xC41F, 0x4FF7,
	[8]byte{0x91, 0x23, 0x64, 0x35, 0x73, 0x6E, 0xAD, 0xC2}}

type IHttpBufferContentFactoryInterface interface {
	win32.IInspectableInterface
	CreateFromBuffer(content *IBuffer) *IHttpContent
	CreateFromBufferWithOffset(content *IBuffer, offset uint32, count uint32) *IHttpContent
}

type IHttpBufferContentFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateFromBuffer           uintptr
	CreateFromBufferWithOffset uintptr
}

type IHttpBufferContentFactory struct {
	win32.IInspectable
}

func (this *IHttpBufferContentFactory) Vtbl() *IHttpBufferContentFactoryVtbl {
	return (*IHttpBufferContentFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpBufferContentFactory) CreateFromBuffer(content *IBuffer) *IHttpContent {
	var _result *IHttpContent
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromBuffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(content)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpBufferContentFactory) CreateFromBufferWithOffset(content *IBuffer, offset uint32, count uint32) *IHttpContent {
	var _result *IHttpContent
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromBufferWithOffset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(content)), uintptr(offset), uintptr(count), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7FDA1151-3574-4880-A8BA-E6B1E0061F3D
var IID_IHttpClient = syscall.GUID{0x7FDA1151, 0x3574, 0x4880,
	[8]byte{0xA8, 0xBA, 0xE6, 0xB1, 0xE0, 0x06, 0x1F, 0x3D}}

type IHttpClientInterface interface {
	win32.IInspectableInterface
	DeleteAsync(uri *IUriRuntimeClass) *IAsyncOperationWithProgress[*IHttpResponseMessage, HttpProgress]
	GetAsync(uri *IUriRuntimeClass) *IAsyncOperationWithProgress[*IHttpResponseMessage, HttpProgress]
	GetWithOptionAsync(uri *IUriRuntimeClass, completionOption HttpCompletionOption) *IAsyncOperationWithProgress[*IHttpResponseMessage, HttpProgress]
	GetBufferAsync(uri *IUriRuntimeClass) *IAsyncOperationWithProgress[*IBuffer, HttpProgress]
	GetInputStreamAsync(uri *IUriRuntimeClass) *IAsyncOperationWithProgress[*IInputStream, HttpProgress]
	GetStringAsync(uri *IUriRuntimeClass) *IAsyncOperationWithProgress[string, HttpProgress]
	PostAsync(uri *IUriRuntimeClass, content *IHttpContent) *IAsyncOperationWithProgress[*IHttpResponseMessage, HttpProgress]
	PutAsync(uri *IUriRuntimeClass, content *IHttpContent) *IAsyncOperationWithProgress[*IHttpResponseMessage, HttpProgress]
	SendRequestAsync(request *IHttpRequestMessage) *IAsyncOperationWithProgress[*IHttpResponseMessage, HttpProgress]
	SendRequestWithOptionAsync(request *IHttpRequestMessage, completionOption HttpCompletionOption) *IAsyncOperationWithProgress[*IHttpResponseMessage, HttpProgress]
	Get_DefaultRequestHeaders() *IHttpRequestHeaderCollection
}

type IHttpClientVtbl struct {
	win32.IInspectableVtbl
	DeleteAsync                uintptr
	GetAsync                   uintptr
	GetWithOptionAsync         uintptr
	GetBufferAsync             uintptr
	GetInputStreamAsync        uintptr
	GetStringAsync             uintptr
	PostAsync                  uintptr
	PutAsync                   uintptr
	SendRequestAsync           uintptr
	SendRequestWithOptionAsync uintptr
	Get_DefaultRequestHeaders  uintptr
}

type IHttpClient struct {
	win32.IInspectable
}

func (this *IHttpClient) Vtbl() *IHttpClientVtbl {
	return (*IHttpClientVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpClient) DeleteAsync(uri *IUriRuntimeClass) *IAsyncOperationWithProgress[*IHttpResponseMessage, HttpProgress] {
	var _result *IAsyncOperationWithProgress[*IHttpResponseMessage, HttpProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DeleteAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpClient) GetAsync(uri *IUriRuntimeClass) *IAsyncOperationWithProgress[*IHttpResponseMessage, HttpProgress] {
	var _result *IAsyncOperationWithProgress[*IHttpResponseMessage, HttpProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpClient) GetWithOptionAsync(uri *IUriRuntimeClass, completionOption HttpCompletionOption) *IAsyncOperationWithProgress[*IHttpResponseMessage, HttpProgress] {
	var _result *IAsyncOperationWithProgress[*IHttpResponseMessage, HttpProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetWithOptionAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(completionOption), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpClient) GetBufferAsync(uri *IUriRuntimeClass) *IAsyncOperationWithProgress[*IBuffer, HttpProgress] {
	var _result *IAsyncOperationWithProgress[*IBuffer, HttpProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetBufferAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpClient) GetInputStreamAsync(uri *IUriRuntimeClass) *IAsyncOperationWithProgress[*IInputStream, HttpProgress] {
	var _result *IAsyncOperationWithProgress[*IInputStream, HttpProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetInputStreamAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpClient) GetStringAsync(uri *IUriRuntimeClass) *IAsyncOperationWithProgress[string, HttpProgress] {
	var _result *IAsyncOperationWithProgress[string, HttpProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStringAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpClient) PostAsync(uri *IUriRuntimeClass, content *IHttpContent) *IAsyncOperationWithProgress[*IHttpResponseMessage, HttpProgress] {
	var _result *IAsyncOperationWithProgress[*IHttpResponseMessage, HttpProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PostAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(content)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpClient) PutAsync(uri *IUriRuntimeClass, content *IHttpContent) *IAsyncOperationWithProgress[*IHttpResponseMessage, HttpProgress] {
	var _result *IAsyncOperationWithProgress[*IHttpResponseMessage, HttpProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PutAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(content)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpClient) SendRequestAsync(request *IHttpRequestMessage) *IAsyncOperationWithProgress[*IHttpResponseMessage, HttpProgress] {
	var _result *IAsyncOperationWithProgress[*IHttpResponseMessage, HttpProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendRequestAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(request)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpClient) SendRequestWithOptionAsync(request *IHttpRequestMessage, completionOption HttpCompletionOption) *IAsyncOperationWithProgress[*IHttpResponseMessage, HttpProgress] {
	var _result *IAsyncOperationWithProgress[*IHttpResponseMessage, HttpProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendRequestWithOptionAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(request)), uintptr(completionOption), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpClient) Get_DefaultRequestHeaders() *IHttpRequestHeaderCollection {
	var _result *IHttpRequestHeaderCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DefaultRequestHeaders, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CDD83348-E8B7-4CEC-B1B0-DC455FE72C92
var IID_IHttpClient2 = syscall.GUID{0xCDD83348, 0xE8B7, 0x4CEC,
	[8]byte{0xB1, 0xB0, 0xDC, 0x45, 0x5F, 0xE7, 0x2C, 0x92}}

type IHttpClient2Interface interface {
	win32.IInspectableInterface
	TryDeleteAsync(uri *IUriRuntimeClass) *IAsyncOperationWithProgress[*IHttpRequestResult, HttpProgress]
	TryGetAsync(uri *IUriRuntimeClass) *IAsyncOperationWithProgress[*IHttpRequestResult, HttpProgress]
	TryGetAsync2(uri *IUriRuntimeClass, completionOption HttpCompletionOption) *IAsyncOperationWithProgress[*IHttpRequestResult, HttpProgress]
	TryGetBufferAsync(uri *IUriRuntimeClass) *IAsyncOperationWithProgress[*IHttpGetBufferResult, HttpProgress]
	TryGetInputStreamAsync(uri *IUriRuntimeClass) *IAsyncOperationWithProgress[*IHttpGetInputStreamResult, HttpProgress]
	TryGetStringAsync(uri *IUriRuntimeClass) *IAsyncOperationWithProgress[*IHttpGetStringResult, HttpProgress]
	TryPostAsync(uri *IUriRuntimeClass, content *IHttpContent) *IAsyncOperationWithProgress[*IHttpRequestResult, HttpProgress]
	TryPutAsync(uri *IUriRuntimeClass, content *IHttpContent) *IAsyncOperationWithProgress[*IHttpRequestResult, HttpProgress]
	TrySendRequestAsync(request *IHttpRequestMessage) *IAsyncOperationWithProgress[*IHttpRequestResult, HttpProgress]
	TrySendRequestAsync2(request *IHttpRequestMessage, completionOption HttpCompletionOption) *IAsyncOperationWithProgress[*IHttpRequestResult, HttpProgress]
}

type IHttpClient2Vtbl struct {
	win32.IInspectableVtbl
	TryDeleteAsync         uintptr
	TryGetAsync            uintptr
	TryGetAsync2           uintptr
	TryGetBufferAsync      uintptr
	TryGetInputStreamAsync uintptr
	TryGetStringAsync      uintptr
	TryPostAsync           uintptr
	TryPutAsync            uintptr
	TrySendRequestAsync    uintptr
	TrySendRequestAsync2   uintptr
}

type IHttpClient2 struct {
	win32.IInspectable
}

func (this *IHttpClient2) Vtbl() *IHttpClient2Vtbl {
	return (*IHttpClient2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpClient2) TryDeleteAsync(uri *IUriRuntimeClass) *IAsyncOperationWithProgress[*IHttpRequestResult, HttpProgress] {
	var _result *IAsyncOperationWithProgress[*IHttpRequestResult, HttpProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryDeleteAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpClient2) TryGetAsync(uri *IUriRuntimeClass) *IAsyncOperationWithProgress[*IHttpRequestResult, HttpProgress] {
	var _result *IAsyncOperationWithProgress[*IHttpRequestResult, HttpProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpClient2) TryGetAsync2(uri *IUriRuntimeClass, completionOption HttpCompletionOption) *IAsyncOperationWithProgress[*IHttpRequestResult, HttpProgress] {
	var _result *IAsyncOperationWithProgress[*IHttpRequestResult, HttpProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetAsync2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(completionOption), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpClient2) TryGetBufferAsync(uri *IUriRuntimeClass) *IAsyncOperationWithProgress[*IHttpGetBufferResult, HttpProgress] {
	var _result *IAsyncOperationWithProgress[*IHttpGetBufferResult, HttpProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetBufferAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpClient2) TryGetInputStreamAsync(uri *IUriRuntimeClass) *IAsyncOperationWithProgress[*IHttpGetInputStreamResult, HttpProgress] {
	var _result *IAsyncOperationWithProgress[*IHttpGetInputStreamResult, HttpProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetInputStreamAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpClient2) TryGetStringAsync(uri *IUriRuntimeClass) *IAsyncOperationWithProgress[*IHttpGetStringResult, HttpProgress] {
	var _result *IAsyncOperationWithProgress[*IHttpGetStringResult, HttpProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetStringAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpClient2) TryPostAsync(uri *IUriRuntimeClass, content *IHttpContent) *IAsyncOperationWithProgress[*IHttpRequestResult, HttpProgress] {
	var _result *IAsyncOperationWithProgress[*IHttpRequestResult, HttpProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryPostAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(content)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpClient2) TryPutAsync(uri *IUriRuntimeClass, content *IHttpContent) *IAsyncOperationWithProgress[*IHttpRequestResult, HttpProgress] {
	var _result *IAsyncOperationWithProgress[*IHttpRequestResult, HttpProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryPutAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(content)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpClient2) TrySendRequestAsync(request *IHttpRequestMessage) *IAsyncOperationWithProgress[*IHttpRequestResult, HttpProgress] {
	var _result *IAsyncOperationWithProgress[*IHttpRequestResult, HttpProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TrySendRequestAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(request)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpClient2) TrySendRequestAsync2(request *IHttpRequestMessage, completionOption HttpCompletionOption) *IAsyncOperationWithProgress[*IHttpRequestResult, HttpProgress] {
	var _result *IAsyncOperationWithProgress[*IHttpRequestResult, HttpProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TrySendRequestAsync2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(request)), uintptr(completionOption), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C30C4ECA-E3FA-4F99-AFB4-63CC65009462
var IID_IHttpClientFactory = syscall.GUID{0xC30C4ECA, 0xE3FA, 0x4F99,
	[8]byte{0xAF, 0xB4, 0x63, 0xCC, 0x65, 0x00, 0x94, 0x62}}

type IHttpClientFactoryInterface interface {
	win32.IInspectableInterface
	Create(filter *IHttpFilter) *IHttpClient
}

type IHttpClientFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IHttpClientFactory struct {
	win32.IInspectable
}

func (this *IHttpClientFactory) Vtbl() *IHttpClientFactoryVtbl {
	return (*IHttpClientFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpClientFactory) Create(filter *IHttpFilter) *IHttpClient {
	var _result *IHttpClient
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(filter)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6B14A441-FBA7-4BD2-AF0A-839DE7C295DA
var IID_IHttpContent = syscall.GUID{0x6B14A441, 0xFBA7, 0x4BD2,
	[8]byte{0xAF, 0x0A, 0x83, 0x9D, 0xE7, 0xC2, 0x95, 0xDA}}

type IHttpContentInterface interface {
	win32.IInspectableInterface
	Get_Headers() *IHttpContentHeaderCollection
	BufferAllAsync() *IAsyncOperationWithProgress[uint64, uint64]
	ReadAsBufferAsync() *IAsyncOperationWithProgress[*IBuffer, uint64]
	ReadAsInputStreamAsync() *IAsyncOperationWithProgress[*IInputStream, uint64]
	ReadAsStringAsync() *IAsyncOperationWithProgress[string, uint64]
	TryComputeLength(length uint64) bool
	WriteToStreamAsync(outputStream *IOutputStream) *IAsyncOperationWithProgress[uint64, uint64]
}

type IHttpContentVtbl struct {
	win32.IInspectableVtbl
	Get_Headers            uintptr
	BufferAllAsync         uintptr
	ReadAsBufferAsync      uintptr
	ReadAsInputStreamAsync uintptr
	ReadAsStringAsync      uintptr
	TryComputeLength       uintptr
	WriteToStreamAsync     uintptr
}

type IHttpContent struct {
	win32.IInspectable
}

func (this *IHttpContent) Vtbl() *IHttpContentVtbl {
	return (*IHttpContentVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpContent) Get_Headers() *IHttpContentHeaderCollection {
	var _result *IHttpContentHeaderCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Headers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpContent) BufferAllAsync() *IAsyncOperationWithProgress[uint64, uint64] {
	var _result *IAsyncOperationWithProgress[uint64, uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().BufferAllAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpContent) ReadAsBufferAsync() *IAsyncOperationWithProgress[*IBuffer, uint64] {
	var _result *IAsyncOperationWithProgress[*IBuffer, uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadAsBufferAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpContent) ReadAsInputStreamAsync() *IAsyncOperationWithProgress[*IInputStream, uint64] {
	var _result *IAsyncOperationWithProgress[*IInputStream, uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadAsInputStreamAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpContent) ReadAsStringAsync() *IAsyncOperationWithProgress[string, uint64] {
	var _result *IAsyncOperationWithProgress[string, uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadAsStringAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpContent) TryComputeLength(length uint64) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryComputeLength, uintptr(unsafe.Pointer(this)), uintptr(length), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHttpContent) WriteToStreamAsync(outputStream *IOutputStream) *IAsyncOperationWithProgress[uint64, uint64] {
	var _result *IAsyncOperationWithProgress[uint64, uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteToStreamAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(outputStream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1F5488E2-CC2D-4779-86A7-88F10687D249
var IID_IHttpCookie = syscall.GUID{0x1F5488E2, 0xCC2D, 0x4779,
	[8]byte{0x86, 0xA7, 0x88, 0xF1, 0x06, 0x87, 0xD2, 0x49}}

type IHttpCookieInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Get_Domain() string
	Get_Path() string
	Get_Expires() *IReference[DateTime]
	Put_Expires(value *IReference[DateTime])
	Get_HttpOnly() bool
	Put_HttpOnly(value bool)
	Get_Secure() bool
	Put_Secure(value bool)
	Get_Value() string
	Put_Value(value string)
}

type IHttpCookieVtbl struct {
	win32.IInspectableVtbl
	Get_Name     uintptr
	Get_Domain   uintptr
	Get_Path     uintptr
	Get_Expires  uintptr
	Put_Expires  uintptr
	Get_HttpOnly uintptr
	Put_HttpOnly uintptr
	Get_Secure   uintptr
	Put_Secure   uintptr
	Get_Value    uintptr
	Put_Value    uintptr
}

type IHttpCookie struct {
	win32.IInspectable
}

func (this *IHttpCookie) Vtbl() *IHttpCookieVtbl {
	return (*IHttpCookieVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpCookie) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHttpCookie) Get_Domain() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Domain, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHttpCookie) Get_Path() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Path, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHttpCookie) Get_Expires() *IReference[DateTime] {
	var _result *IReference[DateTime]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Expires, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpCookie) Put_Expires(value *IReference[DateTime]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Expires, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpCookie) Get_HttpOnly() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HttpOnly, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHttpCookie) Put_HttpOnly(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_HttpOnly, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IHttpCookie) Get_Secure() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Secure, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHttpCookie) Put_Secure(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Secure, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IHttpCookie) Get_Value() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHttpCookie) Put_Value(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Value, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// 6A0585A9-931C-4CD1-A96D-C21701785C5F
var IID_IHttpCookieFactory = syscall.GUID{0x6A0585A9, 0x931C, 0x4CD1,
	[8]byte{0xA9, 0x6D, 0xC2, 0x17, 0x01, 0x78, 0x5C, 0x5F}}

type IHttpCookieFactoryInterface interface {
	win32.IInspectableInterface
	Create(name string, domain string, path string) *IHttpCookie
}

type IHttpCookieFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IHttpCookieFactory struct {
	win32.IInspectable
}

func (this *IHttpCookieFactory) Vtbl() *IHttpCookieFactoryVtbl {
	return (*IHttpCookieFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpCookieFactory) Create(name string, domain string, path string) *IHttpCookie {
	var _result *IHttpCookie
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, NewHStr(domain).Ptr, NewHStr(path).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7A431780-CD4F-4E57-A84A-5B0A53D6BB96
var IID_IHttpCookieManager = syscall.GUID{0x7A431780, 0xCD4F, 0x4E57,
	[8]byte{0xA8, 0x4A, 0x5B, 0x0A, 0x53, 0xD6, 0xBB, 0x96}}

type IHttpCookieManagerInterface interface {
	win32.IInspectableInterface
	SetCookie(cookie *IHttpCookie) bool
	SetCookieWithThirdParty(cookie *IHttpCookie, thirdParty bool) bool
	DeleteCookie(cookie *IHttpCookie)
	GetCookies(uri *IUriRuntimeClass) *IVectorView[*IHttpCookie]
}

type IHttpCookieManagerVtbl struct {
	win32.IInspectableVtbl
	SetCookie               uintptr
	SetCookieWithThirdParty uintptr
	DeleteCookie            uintptr
	GetCookies              uintptr
}

type IHttpCookieManager struct {
	win32.IInspectable
}

func (this *IHttpCookieManager) Vtbl() *IHttpCookieManagerVtbl {
	return (*IHttpCookieManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpCookieManager) SetCookie(cookie *IHttpCookie) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetCookie, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(cookie)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHttpCookieManager) SetCookieWithThirdParty(cookie *IHttpCookie, thirdParty bool) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetCookieWithThirdParty, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(cookie)), uintptr(*(*byte)(unsafe.Pointer(&thirdParty))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHttpCookieManager) DeleteCookie(cookie *IHttpCookie) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DeleteCookie, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(cookie)))
	_ = _hr
}

func (this *IHttpCookieManager) GetCookies(uri *IUriRuntimeClass) *IVectorView[*IHttpCookie] {
	var _result *IVectorView[*IHttpCookie]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCookies, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 43F0138C-2F73-4302-B5F3-EAE9238A5E01
var IID_IHttpFormUrlEncodedContentFactory = syscall.GUID{0x43F0138C, 0x2F73, 0x4302,
	[8]byte{0xB5, 0xF3, 0xEA, 0xE9, 0x23, 0x8A, 0x5E, 0x01}}

type IHttpFormUrlEncodedContentFactoryInterface interface {
	win32.IInspectableInterface
	Create(content *IIterable[*IKeyValuePair[string, string]]) *IHttpContent
}

type IHttpFormUrlEncodedContentFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IHttpFormUrlEncodedContentFactory struct {
	win32.IInspectable
}

func (this *IHttpFormUrlEncodedContentFactory) Vtbl() *IHttpFormUrlEncodedContentFactoryVtbl {
	return (*IHttpFormUrlEncodedContentFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpFormUrlEncodedContentFactory) Create(content *IIterable[*IKeyValuePair[string, string]]) *IHttpContent {
	var _result *IHttpContent
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(content)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 53D08E7C-E209-404E-9A49-742D8236FD3A
var IID_IHttpGetBufferResult = syscall.GUID{0x53D08E7C, 0xE209, 0x404E,
	[8]byte{0x9A, 0x49, 0x74, 0x2D, 0x82, 0x36, 0xFD, 0x3A}}

type IHttpGetBufferResultInterface interface {
	win32.IInspectableInterface
	Get_ExtendedError() HResult
	Get_RequestMessage() *IHttpRequestMessage
	Get_ResponseMessage() *IHttpResponseMessage
	Get_Succeeded() bool
	Get_Value() *IBuffer
}

type IHttpGetBufferResultVtbl struct {
	win32.IInspectableVtbl
	Get_ExtendedError   uintptr
	Get_RequestMessage  uintptr
	Get_ResponseMessage uintptr
	Get_Succeeded       uintptr
	Get_Value           uintptr
}

type IHttpGetBufferResult struct {
	win32.IInspectable
}

func (this *IHttpGetBufferResult) Vtbl() *IHttpGetBufferResultVtbl {
	return (*IHttpGetBufferResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpGetBufferResult) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHttpGetBufferResult) Get_RequestMessage() *IHttpRequestMessage {
	var _result *IHttpRequestMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpGetBufferResult) Get_ResponseMessage() *IHttpResponseMessage {
	var _result *IHttpResponseMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResponseMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpGetBufferResult) Get_Succeeded() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Succeeded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHttpGetBufferResult) Get_Value() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D5D63463-13AA-4EE0-BE95-A0C39FE91203
var IID_IHttpGetInputStreamResult = syscall.GUID{0xD5D63463, 0x13AA, 0x4EE0,
	[8]byte{0xBE, 0x95, 0xA0, 0xC3, 0x9F, 0xE9, 0x12, 0x03}}

type IHttpGetInputStreamResultInterface interface {
	win32.IInspectableInterface
	Get_ExtendedError() HResult
	Get_RequestMessage() *IHttpRequestMessage
	Get_ResponseMessage() *IHttpResponseMessage
	Get_Succeeded() bool
	Get_Value() *IInputStream
}

type IHttpGetInputStreamResultVtbl struct {
	win32.IInspectableVtbl
	Get_ExtendedError   uintptr
	Get_RequestMessage  uintptr
	Get_ResponseMessage uintptr
	Get_Succeeded       uintptr
	Get_Value           uintptr
}

type IHttpGetInputStreamResult struct {
	win32.IInspectable
}

func (this *IHttpGetInputStreamResult) Vtbl() *IHttpGetInputStreamResultVtbl {
	return (*IHttpGetInputStreamResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpGetInputStreamResult) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHttpGetInputStreamResult) Get_RequestMessage() *IHttpRequestMessage {
	var _result *IHttpRequestMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpGetInputStreamResult) Get_ResponseMessage() *IHttpResponseMessage {
	var _result *IHttpResponseMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResponseMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpGetInputStreamResult) Get_Succeeded() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Succeeded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHttpGetInputStreamResult) Get_Value() *IInputStream {
	var _result *IInputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9BAC466D-8509-4775-B16D-8953F47A7F5F
var IID_IHttpGetStringResult = syscall.GUID{0x9BAC466D, 0x8509, 0x4775,
	[8]byte{0xB1, 0x6D, 0x89, 0x53, 0xF4, 0x7A, 0x7F, 0x5F}}

type IHttpGetStringResultInterface interface {
	win32.IInspectableInterface
	Get_ExtendedError() HResult
	Get_RequestMessage() *IHttpRequestMessage
	Get_ResponseMessage() *IHttpResponseMessage
	Get_Succeeded() bool
	Get_Value() string
}

type IHttpGetStringResultVtbl struct {
	win32.IInspectableVtbl
	Get_ExtendedError   uintptr
	Get_RequestMessage  uintptr
	Get_ResponseMessage uintptr
	Get_Succeeded       uintptr
	Get_Value           uintptr
}

type IHttpGetStringResult struct {
	win32.IInspectable
}

func (this *IHttpGetStringResult) Vtbl() *IHttpGetStringResultVtbl {
	return (*IHttpGetStringResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpGetStringResult) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHttpGetStringResult) Get_RequestMessage() *IHttpRequestMessage {
	var _result *IHttpRequestMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpGetStringResult) Get_ResponseMessage() *IHttpResponseMessage {
	var _result *IHttpResponseMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResponseMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpGetStringResult) Get_Succeeded() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Succeeded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHttpGetStringResult) Get_Value() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 728D4022-700D-4FE0-AFA5-40299C58DBFD
var IID_IHttpMethod = syscall.GUID{0x728D4022, 0x700D, 0x4FE0,
	[8]byte{0xAF, 0xA5, 0x40, 0x29, 0x9C, 0x58, 0xDB, 0xFD}}

type IHttpMethodInterface interface {
	win32.IInspectableInterface
	Get_Method() string
}

type IHttpMethodVtbl struct {
	win32.IInspectableVtbl
	Get_Method uintptr
}

type IHttpMethod struct {
	win32.IInspectable
}

func (this *IHttpMethod) Vtbl() *IHttpMethodVtbl {
	return (*IHttpMethodVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpMethod) Get_Method() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Method, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 3C51D10D-36D7-40F8-A86D-E759CAF2F83F
var IID_IHttpMethodFactory = syscall.GUID{0x3C51D10D, 0x36D7, 0x40F8,
	[8]byte{0xA8, 0x6D, 0xE7, 0x59, 0xCA, 0xF2, 0xF8, 0x3F}}

type IHttpMethodFactoryInterface interface {
	win32.IInspectableInterface
	Create(method string) *IHttpMethod
}

type IHttpMethodFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IHttpMethodFactory struct {
	win32.IInspectable
}

func (this *IHttpMethodFactory) Vtbl() *IHttpMethodFactoryVtbl {
	return (*IHttpMethodFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpMethodFactory) Create(method string) *IHttpMethod {
	var _result *IHttpMethod
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(method).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 64D171F0-D99A-4153-8DC6-D68CC4CCE317
var IID_IHttpMethodStatics = syscall.GUID{0x64D171F0, 0xD99A, 0x4153,
	[8]byte{0x8D, 0xC6, 0xD6, 0x8C, 0xC4, 0xCC, 0xE3, 0x17}}

type IHttpMethodStaticsInterface interface {
	win32.IInspectableInterface
	Get_Delete() *IHttpMethod
	Get_Get() *IHttpMethod
	Get_Head() *IHttpMethod
	Get_Options() *IHttpMethod
	Get_Patch() *IHttpMethod
	Get_Post() *IHttpMethod
	Get_Put() *IHttpMethod
}

type IHttpMethodStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Delete  uintptr
	Get_Get     uintptr
	Get_Head    uintptr
	Get_Options uintptr
	Get_Patch   uintptr
	Get_Post    uintptr
	Get_Put     uintptr
}

type IHttpMethodStatics struct {
	win32.IInspectable
}

func (this *IHttpMethodStatics) Vtbl() *IHttpMethodStaticsVtbl {
	return (*IHttpMethodStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpMethodStatics) Get_Delete() *IHttpMethod {
	var _result *IHttpMethod
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Delete, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpMethodStatics) Get_Get() *IHttpMethod {
	var _result *IHttpMethod
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Get, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpMethodStatics) Get_Head() *IHttpMethod {
	var _result *IHttpMethod
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Head, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpMethodStatics) Get_Options() *IHttpMethod {
	var _result *IHttpMethod
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Options, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpMethodStatics) Get_Patch() *IHttpMethod {
	var _result *IHttpMethod
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Patch, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpMethodStatics) Get_Post() *IHttpMethod {
	var _result *IHttpMethod
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Post, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpMethodStatics) Get_Put() *IHttpMethod {
	var _result *IHttpMethod
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Put, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DF916AFF-9926-4AC9-AAF1-E0D04EF09BB9
var IID_IHttpMultipartContent = syscall.GUID{0xDF916AFF, 0x9926, 0x4AC9,
	[8]byte{0xAA, 0xF1, 0xE0, 0xD0, 0x4E, 0xF0, 0x9B, 0xB9}}

type IHttpMultipartContentInterface interface {
	win32.IInspectableInterface
	Add(content *IHttpContent)
}

type IHttpMultipartContentVtbl struct {
	win32.IInspectableVtbl
	Add uintptr
}

type IHttpMultipartContent struct {
	win32.IInspectable
}

func (this *IHttpMultipartContent) Vtbl() *IHttpMultipartContentVtbl {
	return (*IHttpMultipartContentVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpMultipartContent) Add(content *IHttpContent) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(content)))
	_ = _hr
}

// 7EB42E62-0222-4F20-B372-47D5DB5D33B4
var IID_IHttpMultipartContentFactory = syscall.GUID{0x7EB42E62, 0x0222, 0x4F20,
	[8]byte{0xB3, 0x72, 0x47, 0xD5, 0xDB, 0x5D, 0x33, 0xB4}}

type IHttpMultipartContentFactoryInterface interface {
	win32.IInspectableInterface
	CreateWithSubtype(subtype string) *IHttpContent
	CreateWithSubtypeAndBoundary(subtype string, boundary string) *IHttpContent
}

type IHttpMultipartContentFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateWithSubtype            uintptr
	CreateWithSubtypeAndBoundary uintptr
}

type IHttpMultipartContentFactory struct {
	win32.IInspectable
}

func (this *IHttpMultipartContentFactory) Vtbl() *IHttpMultipartContentFactoryVtbl {
	return (*IHttpMultipartContentFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpMultipartContentFactory) CreateWithSubtype(subtype string) *IHttpContent {
	var _result *IHttpContent
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithSubtype, uintptr(unsafe.Pointer(this)), NewHStr(subtype).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpMultipartContentFactory) CreateWithSubtypeAndBoundary(subtype string, boundary string) *IHttpContent {
	var _result *IHttpContent
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithSubtypeAndBoundary, uintptr(unsafe.Pointer(this)), NewHStr(subtype).Ptr, NewHStr(boundary).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 64D337E2-E967-4624-B6D1-CF74604A4A42
var IID_IHttpMultipartFormDataContent = syscall.GUID{0x64D337E2, 0xE967, 0x4624,
	[8]byte{0xB6, 0xD1, 0xCF, 0x74, 0x60, 0x4A, 0x4A, 0x42}}

type IHttpMultipartFormDataContentInterface interface {
	win32.IInspectableInterface
	Add(content *IHttpContent)
	AddWithName(content *IHttpContent, name string)
	AddWithNameAndFileName(content *IHttpContent, name string, fileName string)
}

type IHttpMultipartFormDataContentVtbl struct {
	win32.IInspectableVtbl
	Add                    uintptr
	AddWithName            uintptr
	AddWithNameAndFileName uintptr
}

type IHttpMultipartFormDataContent struct {
	win32.IInspectable
}

func (this *IHttpMultipartFormDataContent) Vtbl() *IHttpMultipartFormDataContentVtbl {
	return (*IHttpMultipartFormDataContentVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpMultipartFormDataContent) Add(content *IHttpContent) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(content)))
	_ = _hr
}

func (this *IHttpMultipartFormDataContent) AddWithName(content *IHttpContent, name string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddWithName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(content)), NewHStr(name).Ptr)
	_ = _hr
}

func (this *IHttpMultipartFormDataContent) AddWithNameAndFileName(content *IHttpContent, name string, fileName string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddWithNameAndFileName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(content)), NewHStr(name).Ptr, NewHStr(fileName).Ptr)
	_ = _hr
}

// A04D7311-5017-4622-93A8-49B24A4FCBFC
var IID_IHttpMultipartFormDataContentFactory = syscall.GUID{0xA04D7311, 0x5017, 0x4622,
	[8]byte{0x93, 0xA8, 0x49, 0xB2, 0x4A, 0x4F, 0xCB, 0xFC}}

type IHttpMultipartFormDataContentFactoryInterface interface {
	win32.IInspectableInterface
	CreateWithBoundary(boundary string) *IHttpContent
}

type IHttpMultipartFormDataContentFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateWithBoundary uintptr
}

type IHttpMultipartFormDataContentFactory struct {
	win32.IInspectable
}

func (this *IHttpMultipartFormDataContentFactory) Vtbl() *IHttpMultipartFormDataContentFactoryVtbl {
	return (*IHttpMultipartFormDataContentFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpMultipartFormDataContentFactory) CreateWithBoundary(boundary string) *IHttpContent {
	var _result *IHttpContent
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithBoundary, uintptr(unsafe.Pointer(this)), NewHStr(boundary).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F5762B3C-74D4-4811-B5DC-9F8B4E2F9ABF
var IID_IHttpRequestMessage = syscall.GUID{0xF5762B3C, 0x74D4, 0x4811,
	[8]byte{0xB5, 0xDC, 0x9F, 0x8B, 0x4E, 0x2F, 0x9A, 0xBF}}

type IHttpRequestMessageInterface interface {
	win32.IInspectableInterface
	Get_Content() *IHttpContent
	Put_Content(value *IHttpContent)
	Get_Headers() *IHttpRequestHeaderCollection
	Get_Method() *IHttpMethod
	Put_Method(value *IHttpMethod)
	Get_Properties() *IMap[string, interface{}]
	Get_RequestUri() *IUriRuntimeClass
	Put_RequestUri(value *IUriRuntimeClass)
	Get_TransportInformation() *IHttpTransportInformation
}

type IHttpRequestMessageVtbl struct {
	win32.IInspectableVtbl
	Get_Content              uintptr
	Put_Content              uintptr
	Get_Headers              uintptr
	Get_Method               uintptr
	Put_Method               uintptr
	Get_Properties           uintptr
	Get_RequestUri           uintptr
	Put_RequestUri           uintptr
	Get_TransportInformation uintptr
}

type IHttpRequestMessage struct {
	win32.IInspectable
}

func (this *IHttpRequestMessage) Vtbl() *IHttpRequestMessageVtbl {
	return (*IHttpRequestMessageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpRequestMessage) Get_Content() *IHttpContent {
	var _result *IHttpContent
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Content, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpRequestMessage) Put_Content(value *IHttpContent) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Content, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpRequestMessage) Get_Headers() *IHttpRequestHeaderCollection {
	var _result *IHttpRequestHeaderCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Headers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpRequestMessage) Get_Method() *IHttpMethod {
	var _result *IHttpMethod
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Method, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpRequestMessage) Put_Method(value *IHttpMethod) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Method, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpRequestMessage) Get_Properties() *IMap[string, interface{}] {
	var _result *IMap[string, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpRequestMessage) Get_RequestUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpRequestMessage) Put_RequestUri(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RequestUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpRequestMessage) Get_TransportInformation() *IHttpTransportInformation {
	var _result *IHttpTransportInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TransportInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5BAC994E-3886-412E-AEC3-52EC7F25616F
var IID_IHttpRequestMessageFactory = syscall.GUID{0x5BAC994E, 0x3886, 0x412E,
	[8]byte{0xAE, 0xC3, 0x52, 0xEC, 0x7F, 0x25, 0x61, 0x6F}}

type IHttpRequestMessageFactoryInterface interface {
	win32.IInspectableInterface
	Create(method *IHttpMethod, uri *IUriRuntimeClass) *IHttpRequestMessage
}

type IHttpRequestMessageFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IHttpRequestMessageFactory struct {
	win32.IInspectable
}

func (this *IHttpRequestMessageFactory) Vtbl() *IHttpRequestMessageFactoryVtbl {
	return (*IHttpRequestMessageFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpRequestMessageFactory) Create(method *IHttpMethod, uri *IUriRuntimeClass) *IHttpRequestMessage {
	var _result *IHttpRequestMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(method)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6ACF4DA8-B5EB-4A35-A902-4217FBE820C5
var IID_IHttpRequestResult = syscall.GUID{0x6ACF4DA8, 0xB5EB, 0x4A35,
	[8]byte{0xA9, 0x02, 0x42, 0x17, 0xFB, 0xE8, 0x20, 0xC5}}

type IHttpRequestResultInterface interface {
	win32.IInspectableInterface
	Get_ExtendedError() HResult
	Get_RequestMessage() *IHttpRequestMessage
	Get_ResponseMessage() *IHttpResponseMessage
	Get_Succeeded() bool
}

type IHttpRequestResultVtbl struct {
	win32.IInspectableVtbl
	Get_ExtendedError   uintptr
	Get_RequestMessage  uintptr
	Get_ResponseMessage uintptr
	Get_Succeeded       uintptr
}

type IHttpRequestResult struct {
	win32.IInspectable
}

func (this *IHttpRequestResult) Vtbl() *IHttpRequestResultVtbl {
	return (*IHttpRequestResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpRequestResult) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHttpRequestResult) Get_RequestMessage() *IHttpRequestMessage {
	var _result *IHttpRequestMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpRequestResult) Get_ResponseMessage() *IHttpResponseMessage {
	var _result *IHttpResponseMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResponseMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpRequestResult) Get_Succeeded() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Succeeded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// FEE200FB-8664-44E0-95D9-42696199BFFC
var IID_IHttpResponseMessage = syscall.GUID{0xFEE200FB, 0x8664, 0x44E0,
	[8]byte{0x95, 0xD9, 0x42, 0x69, 0x61, 0x99, 0xBF, 0xFC}}

type IHttpResponseMessageInterface interface {
	win32.IInspectableInterface
	Get_Content() *IHttpContent
	Put_Content(value *IHttpContent)
	Get_Headers() *IHttpResponseHeaderCollection
	Get_IsSuccessStatusCode() bool
	Get_ReasonPhrase() string
	Put_ReasonPhrase(value string)
	Get_RequestMessage() *IHttpRequestMessage
	Put_RequestMessage(value *IHttpRequestMessage)
	Get_Source() HttpResponseMessageSource
	Put_Source(value HttpResponseMessageSource)
	Get_StatusCode() HttpStatusCode
	Put_StatusCode(value HttpStatusCode)
	Get_Version() HttpVersion
	Put_Version(value HttpVersion)
	EnsureSuccessStatusCode() *IHttpResponseMessage
}

type IHttpResponseMessageVtbl struct {
	win32.IInspectableVtbl
	Get_Content             uintptr
	Put_Content             uintptr
	Get_Headers             uintptr
	Get_IsSuccessStatusCode uintptr
	Get_ReasonPhrase        uintptr
	Put_ReasonPhrase        uintptr
	Get_RequestMessage      uintptr
	Put_RequestMessage      uintptr
	Get_Source              uintptr
	Put_Source              uintptr
	Get_StatusCode          uintptr
	Put_StatusCode          uintptr
	Get_Version             uintptr
	Put_Version             uintptr
	EnsureSuccessStatusCode uintptr
}

type IHttpResponseMessage struct {
	win32.IInspectable
}

func (this *IHttpResponseMessage) Vtbl() *IHttpResponseMessageVtbl {
	return (*IHttpResponseMessageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpResponseMessage) Get_Content() *IHttpContent {
	var _result *IHttpContent
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Content, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpResponseMessage) Put_Content(value *IHttpContent) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Content, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpResponseMessage) Get_Headers() *IHttpResponseHeaderCollection {
	var _result *IHttpResponseHeaderCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Headers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpResponseMessage) Get_IsSuccessStatusCode() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSuccessStatusCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHttpResponseMessage) Get_ReasonPhrase() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReasonPhrase, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHttpResponseMessage) Put_ReasonPhrase(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReasonPhrase, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IHttpResponseMessage) Get_RequestMessage() *IHttpRequestMessage {
	var _result *IHttpRequestMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpResponseMessage) Put_RequestMessage(value *IHttpRequestMessage) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RequestMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpResponseMessage) Get_Source() HttpResponseMessageSource {
	var _result HttpResponseMessageSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Source, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHttpResponseMessage) Put_Source(value HttpResponseMessageSource) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Source, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IHttpResponseMessage) Get_StatusCode() HttpStatusCode {
	var _result HttpStatusCode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StatusCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHttpResponseMessage) Put_StatusCode(value HttpStatusCode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StatusCode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IHttpResponseMessage) Get_Version() HttpVersion {
	var _result HttpVersion
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Version, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHttpResponseMessage) Put_Version(value HttpVersion) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Version, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IHttpResponseMessage) EnsureSuccessStatusCode() *IHttpResponseMessage {
	var _result *IHttpResponseMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnsureSuccessStatusCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 52A8AF99-F095-43DA-B60F-7CFC2BC7EA2F
var IID_IHttpResponseMessageFactory = syscall.GUID{0x52A8AF99, 0xF095, 0x43DA,
	[8]byte{0xB6, 0x0F, 0x7C, 0xFC, 0x2B, 0xC7, 0xEA, 0x2F}}

type IHttpResponseMessageFactoryInterface interface {
	win32.IInspectableInterface
	Create(statusCode HttpStatusCode) *IHttpResponseMessage
}

type IHttpResponseMessageFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IHttpResponseMessageFactory struct {
	win32.IInspectable
}

func (this *IHttpResponseMessageFactory) Vtbl() *IHttpResponseMessageFactoryVtbl {
	return (*IHttpResponseMessageFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpResponseMessageFactory) Create(statusCode HttpStatusCode) *IHttpResponseMessage {
	var _result *IHttpResponseMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(statusCode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F3E64D9D-F725-407E-942F-0EDA189809F4
var IID_IHttpStreamContentFactory = syscall.GUID{0xF3E64D9D, 0xF725, 0x407E,
	[8]byte{0x94, 0x2F, 0x0E, 0xDA, 0x18, 0x98, 0x09, 0xF4}}

type IHttpStreamContentFactoryInterface interface {
	win32.IInspectableInterface
	CreateFromInputStream(content *IInputStream) *IHttpContent
}

type IHttpStreamContentFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateFromInputStream uintptr
}

type IHttpStreamContentFactory struct {
	win32.IInspectable
}

func (this *IHttpStreamContentFactory) Vtbl() *IHttpStreamContentFactoryVtbl {
	return (*IHttpStreamContentFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpStreamContentFactory) CreateFromInputStream(content *IInputStream) *IHttpContent {
	var _result *IHttpContent
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromInputStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(content)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 46649D5B-2E93-48EB-8E61-19677878E57F
var IID_IHttpStringContentFactory = syscall.GUID{0x46649D5B, 0x2E93, 0x48EB,
	[8]byte{0x8E, 0x61, 0x19, 0x67, 0x78, 0x78, 0xE5, 0x7F}}

type IHttpStringContentFactoryInterface interface {
	win32.IInspectableInterface
	CreateFromString(content string) *IHttpContent
	CreateFromStringWithEncoding(content string, encoding UnicodeEncoding) *IHttpContent
	CreateFromStringWithEncodingAndMediaType(content string, encoding UnicodeEncoding, mediaType string) *IHttpContent
}

type IHttpStringContentFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateFromString                         uintptr
	CreateFromStringWithEncoding             uintptr
	CreateFromStringWithEncodingAndMediaType uintptr
}

type IHttpStringContentFactory struct {
	win32.IInspectable
}

func (this *IHttpStringContentFactory) Vtbl() *IHttpStringContentFactoryVtbl {
	return (*IHttpStringContentFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpStringContentFactory) CreateFromString(content string) *IHttpContent {
	var _result *IHttpContent
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromString, uintptr(unsafe.Pointer(this)), NewHStr(content).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpStringContentFactory) CreateFromStringWithEncoding(content string, encoding UnicodeEncoding) *IHttpContent {
	var _result *IHttpContent
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromStringWithEncoding, uintptr(unsafe.Pointer(this)), NewHStr(content).Ptr, uintptr(encoding), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpStringContentFactory) CreateFromStringWithEncodingAndMediaType(content string, encoding UnicodeEncoding, mediaType string) *IHttpContent {
	var _result *IHttpContent
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromStringWithEncodingAndMediaType, uintptr(unsafe.Pointer(this)), NewHStr(content).Ptr, uintptr(encoding), NewHStr(mediaType).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 70127198-C6A7-4ED0-833A-83FD8B8F178D
var IID_IHttpTransportInformation = syscall.GUID{0x70127198, 0xC6A7, 0x4ED0,
	[8]byte{0x83, 0x3A, 0x83, 0xFD, 0x8B, 0x8F, 0x17, 0x8D}}

type IHttpTransportInformationInterface interface {
	win32.IInspectableInterface
	Get_ServerCertificate() *ICertificate
	Get_ServerCertificateErrorSeverity() SocketSslErrorSeverity
	Get_ServerCertificateErrors() *IVectorView[ChainValidationResult]
	Get_ServerIntermediateCertificates() *IVectorView[*ICertificate]
}

type IHttpTransportInformationVtbl struct {
	win32.IInspectableVtbl
	Get_ServerCertificate              uintptr
	Get_ServerCertificateErrorSeverity uintptr
	Get_ServerCertificateErrors        uintptr
	Get_ServerIntermediateCertificates uintptr
}

type IHttpTransportInformation struct {
	win32.IInspectable
}

func (this *IHttpTransportInformation) Vtbl() *IHttpTransportInformationVtbl {
	return (*IHttpTransportInformationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpTransportInformation) Get_ServerCertificate() *ICertificate {
	var _result *ICertificate
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServerCertificate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpTransportInformation) Get_ServerCertificateErrorSeverity() SocketSslErrorSeverity {
	var _result SocketSslErrorSeverity
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServerCertificateErrorSeverity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHttpTransportInformation) Get_ServerCertificateErrors() *IVectorView[ChainValidationResult] {
	var _result *IVectorView[ChainValidationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServerCertificateErrors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpTransportInformation) Get_ServerIntermediateCertificates() *IVectorView[*ICertificate] {
	var _result *IVectorView[*ICertificate]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServerIntermediateCertificates, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type HttpBufferContent struct {
	RtClass
	*IHttpContent
}

func NewHttpBufferContent_CreateFromBuffer(content *IBuffer) *HttpBufferContent {
	hs := NewHStr("Windows.Web.Http.HttpBufferContent")
	var pFac *IHttpBufferContentFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpBufferContentFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpContent
	p = pFac.CreateFromBuffer(content)
	result := &HttpBufferContent{
		RtClass:      RtClass{PInspect: &p.IInspectable},
		IHttpContent: p,
	}
	com.AddToScope(result)
	return result
}

func NewHttpBufferContent_CreateFromBufferWithOffset(content *IBuffer, offset uint32, count uint32) *HttpBufferContent {
	hs := NewHStr("Windows.Web.Http.HttpBufferContent")
	var pFac *IHttpBufferContentFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpBufferContentFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpContent
	p = pFac.CreateFromBufferWithOffset(content, offset, count)
	result := &HttpBufferContent{
		RtClass:      RtClass{PInspect: &p.IInspectable},
		IHttpContent: p,
	}
	com.AddToScope(result)
	return result
}

type HttpClient struct {
	RtClass
	*IHttpClient
}

func NewHttpClient() *HttpClient {
	hs := NewHStr("Windows.Web.Http.HttpClient")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &HttpClient{
		RtClass:     RtClass{PInspect: p},
		IHttpClient: (*IHttpClient)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewHttpClient_Create(filter *IHttpFilter) *HttpClient {
	hs := NewHStr("Windows.Web.Http.HttpClient")
	var pFac *IHttpClientFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpClientFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpClient
	p = pFac.Create(filter)
	result := &HttpClient{
		RtClass:     RtClass{PInspect: &p.IInspectable},
		IHttpClient: p,
	}
	com.AddToScope(result)
	return result
}

type HttpCookie struct {
	RtClass
	*IHttpCookie
}

func NewHttpCookie_Create(name string, domain string, path string) *HttpCookie {
	hs := NewHStr("Windows.Web.Http.HttpCookie")
	var pFac *IHttpCookieFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpCookieFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpCookie
	p = pFac.Create(name, domain, path)
	result := &HttpCookie{
		RtClass:     RtClass{PInspect: &p.IInspectable},
		IHttpCookie: p,
	}
	com.AddToScope(result)
	return result
}

type HttpCookieCollection struct {
	RtClass
	*IVectorView[*IHttpCookie]
}

type HttpCookieManager struct {
	RtClass
	*IHttpCookieManager
}

type HttpFormUrlEncodedContent struct {
	RtClass
	*IHttpContent
}

func NewHttpFormUrlEncodedContent_Create(content *IIterable[*IKeyValuePair[string, string]]) *HttpFormUrlEncodedContent {
	hs := NewHStr("Windows.Web.Http.HttpFormUrlEncodedContent")
	var pFac *IHttpFormUrlEncodedContentFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpFormUrlEncodedContentFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpContent
	p = pFac.Create(content)
	result := &HttpFormUrlEncodedContent{
		RtClass:      RtClass{PInspect: &p.IInspectable},
		IHttpContent: p,
	}
	com.AddToScope(result)
	return result
}

type HttpGetBufferResult struct {
	RtClass
	*IHttpGetBufferResult
}

type HttpGetInputStreamResult struct {
	RtClass
	*IHttpGetInputStreamResult
}

type HttpGetStringResult struct {
	RtClass
	*IHttpGetStringResult
}

type HttpMethod struct {
	RtClass
	*IHttpMethod
}

func NewHttpMethod_Create(method string) *HttpMethod {
	hs := NewHStr("Windows.Web.Http.HttpMethod")
	var pFac *IHttpMethodFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpMethodFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpMethod
	p = pFac.Create(method)
	result := &HttpMethod{
		RtClass:     RtClass{PInspect: &p.IInspectable},
		IHttpMethod: p,
	}
	com.AddToScope(result)
	return result
}

func NewIHttpMethodStatics() *IHttpMethodStatics {
	var p *IHttpMethodStatics
	hs := NewHStr("Windows.Web.Http.HttpMethod")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpMethodStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type HttpMultipartContent struct {
	RtClass
	*IHttpContent
}

func NewHttpMultipartContent() *HttpMultipartContent {
	hs := NewHStr("Windows.Web.Http.HttpMultipartContent")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &HttpMultipartContent{
		RtClass:      RtClass{PInspect: p},
		IHttpContent: (*IHttpContent)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewHttpMultipartContent_CreateWithSubtype(subtype string) *HttpMultipartContent {
	hs := NewHStr("Windows.Web.Http.HttpMultipartContent")
	var pFac *IHttpMultipartContentFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpMultipartContentFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpContent
	p = pFac.CreateWithSubtype(subtype)
	result := &HttpMultipartContent{
		RtClass:      RtClass{PInspect: &p.IInspectable},
		IHttpContent: p,
	}
	com.AddToScope(result)
	return result
}

func NewHttpMultipartContent_CreateWithSubtypeAndBoundary(subtype string, boundary string) *HttpMultipartContent {
	hs := NewHStr("Windows.Web.Http.HttpMultipartContent")
	var pFac *IHttpMultipartContentFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpMultipartContentFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpContent
	p = pFac.CreateWithSubtypeAndBoundary(subtype, boundary)
	result := &HttpMultipartContent{
		RtClass:      RtClass{PInspect: &p.IInspectable},
		IHttpContent: p,
	}
	com.AddToScope(result)
	return result
}

type HttpMultipartFormDataContent struct {
	RtClass
	*IHttpContent
}

func NewHttpMultipartFormDataContent() *HttpMultipartFormDataContent {
	hs := NewHStr("Windows.Web.Http.HttpMultipartFormDataContent")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &HttpMultipartFormDataContent{
		RtClass:      RtClass{PInspect: p},
		IHttpContent: (*IHttpContent)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewHttpMultipartFormDataContent_CreateWithBoundary(boundary string) *HttpMultipartFormDataContent {
	hs := NewHStr("Windows.Web.Http.HttpMultipartFormDataContent")
	var pFac *IHttpMultipartFormDataContentFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpMultipartFormDataContentFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpContent
	p = pFac.CreateWithBoundary(boundary)
	result := &HttpMultipartFormDataContent{
		RtClass:      RtClass{PInspect: &p.IInspectable},
		IHttpContent: p,
	}
	com.AddToScope(result)
	return result
}

type HttpRequestMessage struct {
	RtClass
	*IHttpRequestMessage
}

func NewHttpRequestMessage() *HttpRequestMessage {
	hs := NewHStr("Windows.Web.Http.HttpRequestMessage")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &HttpRequestMessage{
		RtClass:             RtClass{PInspect: p},
		IHttpRequestMessage: (*IHttpRequestMessage)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewHttpRequestMessage_Create(method *IHttpMethod, uri *IUriRuntimeClass) *HttpRequestMessage {
	hs := NewHStr("Windows.Web.Http.HttpRequestMessage")
	var pFac *IHttpRequestMessageFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpRequestMessageFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpRequestMessage
	p = pFac.Create(method, uri)
	result := &HttpRequestMessage{
		RtClass:             RtClass{PInspect: &p.IInspectable},
		IHttpRequestMessage: p,
	}
	com.AddToScope(result)
	return result
}

type HttpRequestResult struct {
	RtClass
	*IHttpRequestResult
}

type HttpResponseMessage struct {
	RtClass
	*IHttpResponseMessage
}

func NewHttpResponseMessage() *HttpResponseMessage {
	hs := NewHStr("Windows.Web.Http.HttpResponseMessage")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &HttpResponseMessage{
		RtClass:              RtClass{PInspect: p},
		IHttpResponseMessage: (*IHttpResponseMessage)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewHttpResponseMessage_Create(statusCode HttpStatusCode) *HttpResponseMessage {
	hs := NewHStr("Windows.Web.Http.HttpResponseMessage")
	var pFac *IHttpResponseMessageFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpResponseMessageFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpResponseMessage
	p = pFac.Create(statusCode)
	result := &HttpResponseMessage{
		RtClass:              RtClass{PInspect: &p.IInspectable},
		IHttpResponseMessage: p,
	}
	com.AddToScope(result)
	return result
}

type HttpStreamContent struct {
	RtClass
	*IHttpContent
}

func NewHttpStreamContent_CreateFromInputStream(content *IInputStream) *HttpStreamContent {
	hs := NewHStr("Windows.Web.Http.HttpStreamContent")
	var pFac *IHttpStreamContentFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpStreamContentFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpContent
	p = pFac.CreateFromInputStream(content)
	result := &HttpStreamContent{
		RtClass:      RtClass{PInspect: &p.IInspectable},
		IHttpContent: p,
	}
	com.AddToScope(result)
	return result
}

type HttpStringContent struct {
	RtClass
	*IHttpContent
}

func NewHttpStringContent_CreateFromString(content string) *HttpStringContent {
	hs := NewHStr("Windows.Web.Http.HttpStringContent")
	var pFac *IHttpStringContentFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpStringContentFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpContent
	p = pFac.CreateFromString(content)
	result := &HttpStringContent{
		RtClass:      RtClass{PInspect: &p.IInspectable},
		IHttpContent: p,
	}
	com.AddToScope(result)
	return result
}

func NewHttpStringContent_CreateFromStringWithEncoding(content string, encoding UnicodeEncoding) *HttpStringContent {
	hs := NewHStr("Windows.Web.Http.HttpStringContent")
	var pFac *IHttpStringContentFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpStringContentFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpContent
	p = pFac.CreateFromStringWithEncoding(content, encoding)
	result := &HttpStringContent{
		RtClass:      RtClass{PInspect: &p.IInspectable},
		IHttpContent: p,
	}
	com.AddToScope(result)
	return result
}

func NewHttpStringContent_CreateFromStringWithEncodingAndMediaType(content string, encoding UnicodeEncoding, mediaType string) *HttpStringContent {
	hs := NewHStr("Windows.Web.Http.HttpStringContent")
	var pFac *IHttpStringContentFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpStringContentFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpContent
	p = pFac.CreateFromStringWithEncodingAndMediaType(content, encoding, mediaType)
	result := &HttpStringContent{
		RtClass:      RtClass{PInspect: &p.IInspectable},
		IHttpContent: p,
	}
	com.AddToScope(result)
	return result
}

type HttpTransportInformation struct {
	RtClass
	*IHttpTransportInformation
}
