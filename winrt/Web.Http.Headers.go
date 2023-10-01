package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"log"
	"syscall"
	"unsafe"
)

// interfaces

// 9A586B89-D5D0-4FBE-BD9D-B5B3636811B4
var IID_IHttpCacheDirectiveHeaderValueCollection = syscall.GUID{0x9A586B89, 0xD5D0, 0x4FBE,
	[8]byte{0xBD, 0x9D, 0xB5, 0xB3, 0x63, 0x68, 0x11, 0xB4}}

type IHttpCacheDirectiveHeaderValueCollectionInterface interface {
	win32.IInspectableInterface
	Get_MaxAge() *IReference[TimeSpan]
	Put_MaxAge(value *IReference[TimeSpan])
	Get_MaxStale() *IReference[TimeSpan]
	Put_MaxStale(value *IReference[TimeSpan])
	Get_MinFresh() *IReference[TimeSpan]
	Put_MinFresh(value *IReference[TimeSpan])
	Get_SharedMaxAge() *IReference[TimeSpan]
	Put_SharedMaxAge(value *IReference[TimeSpan])
	ParseAdd(input string)
	TryParseAdd(input string) bool
}

type IHttpCacheDirectiveHeaderValueCollectionVtbl struct {
	win32.IInspectableVtbl
	Get_MaxAge       uintptr
	Put_MaxAge       uintptr
	Get_MaxStale     uintptr
	Put_MaxStale     uintptr
	Get_MinFresh     uintptr
	Put_MinFresh     uintptr
	Get_SharedMaxAge uintptr
	Put_SharedMaxAge uintptr
	ParseAdd         uintptr
	TryParseAdd      uintptr
}

type IHttpCacheDirectiveHeaderValueCollection struct {
	win32.IInspectable
}

func (this *IHttpCacheDirectiveHeaderValueCollection) Vtbl() *IHttpCacheDirectiveHeaderValueCollectionVtbl {
	return (*IHttpCacheDirectiveHeaderValueCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpCacheDirectiveHeaderValueCollection) Get_MaxAge() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxAge, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpCacheDirectiveHeaderValueCollection) Put_MaxAge(value *IReference[TimeSpan]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxAge, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpCacheDirectiveHeaderValueCollection) Get_MaxStale() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxStale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpCacheDirectiveHeaderValueCollection) Put_MaxStale(value *IReference[TimeSpan]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxStale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpCacheDirectiveHeaderValueCollection) Get_MinFresh() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinFresh, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpCacheDirectiveHeaderValueCollection) Put_MinFresh(value *IReference[TimeSpan]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MinFresh, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpCacheDirectiveHeaderValueCollection) Get_SharedMaxAge() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SharedMaxAge, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpCacheDirectiveHeaderValueCollection) Put_SharedMaxAge(value *IReference[TimeSpan]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SharedMaxAge, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpCacheDirectiveHeaderValueCollection) ParseAdd(input string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ParseAdd, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr)
	_ = _hr
}

func (this *IHttpCacheDirectiveHeaderValueCollection) TryParseAdd(input string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParseAdd, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 393361AF-0F7D-4820-9FDD-A2B956EEAEAB
var IID_IHttpChallengeHeaderValue = syscall.GUID{0x393361AF, 0x0F7D, 0x4820,
	[8]byte{0x9F, 0xDD, 0xA2, 0xB9, 0x56, 0xEE, 0xAE, 0xAB}}

type IHttpChallengeHeaderValueInterface interface {
	win32.IInspectableInterface
	Get_Parameters() *IVector[*IHttpNameValueHeaderValue]
	Get_Scheme() string
	Get_Token() string
}

type IHttpChallengeHeaderValueVtbl struct {
	win32.IInspectableVtbl
	Get_Parameters uintptr
	Get_Scheme     uintptr
	Get_Token      uintptr
}

type IHttpChallengeHeaderValue struct {
	win32.IInspectable
}

func (this *IHttpChallengeHeaderValue) Vtbl() *IHttpChallengeHeaderValueVtbl {
	return (*IHttpChallengeHeaderValueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpChallengeHeaderValue) Get_Parameters() *IVector[*IHttpNameValueHeaderValue] {
	var _result *IVector[*IHttpNameValueHeaderValue]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Parameters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpChallengeHeaderValue) Get_Scheme() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Scheme, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHttpChallengeHeaderValue) Get_Token() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Token, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// CA9E5F81-AEE0-4353-A10B-E625BABD64C2
var IID_IHttpChallengeHeaderValueCollection = syscall.GUID{0xCA9E5F81, 0xAEE0, 0x4353,
	[8]byte{0xA1, 0x0B, 0xE6, 0x25, 0xBA, 0xBD, 0x64, 0xC2}}

type IHttpChallengeHeaderValueCollectionInterface interface {
	win32.IInspectableInterface
	ParseAdd(input string)
	TryParseAdd(input string) bool
}

type IHttpChallengeHeaderValueCollectionVtbl struct {
	win32.IInspectableVtbl
	ParseAdd    uintptr
	TryParseAdd uintptr
}

type IHttpChallengeHeaderValueCollection struct {
	win32.IInspectable
}

func (this *IHttpChallengeHeaderValueCollection) Vtbl() *IHttpChallengeHeaderValueCollectionVtbl {
	return (*IHttpChallengeHeaderValueCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpChallengeHeaderValueCollection) ParseAdd(input string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ParseAdd, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr)
	_ = _hr
}

func (this *IHttpChallengeHeaderValueCollection) TryParseAdd(input string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParseAdd, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C452C451-D99C-40AA-9399-90EEB98FC613
var IID_IHttpChallengeHeaderValueFactory = syscall.GUID{0xC452C451, 0xD99C, 0x40AA,
	[8]byte{0x93, 0x99, 0x90, 0xEE, 0xB9, 0x8F, 0xC6, 0x13}}

type IHttpChallengeHeaderValueFactoryInterface interface {
	win32.IInspectableInterface
	CreateFromScheme(scheme string) *IHttpChallengeHeaderValue
	CreateFromSchemeWithToken(scheme string, token string) *IHttpChallengeHeaderValue
}

type IHttpChallengeHeaderValueFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateFromScheme          uintptr
	CreateFromSchemeWithToken uintptr
}

type IHttpChallengeHeaderValueFactory struct {
	win32.IInspectable
}

func (this *IHttpChallengeHeaderValueFactory) Vtbl() *IHttpChallengeHeaderValueFactoryVtbl {
	return (*IHttpChallengeHeaderValueFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpChallengeHeaderValueFactory) CreateFromScheme(scheme string) *IHttpChallengeHeaderValue {
	var _result *IHttpChallengeHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromScheme, uintptr(unsafe.Pointer(this)), NewHStr(scheme).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpChallengeHeaderValueFactory) CreateFromSchemeWithToken(scheme string, token string) *IHttpChallengeHeaderValue {
	var _result *IHttpChallengeHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromSchemeWithToken, uintptr(unsafe.Pointer(this)), NewHStr(scheme).Ptr, NewHStr(token).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F3D38A72-FC01-4D01-A008-FCB7C459D635
var IID_IHttpChallengeHeaderValueStatics = syscall.GUID{0xF3D38A72, 0xFC01, 0x4D01,
	[8]byte{0xA0, 0x08, 0xFC, 0xB7, 0xC4, 0x59, 0xD6, 0x35}}

type IHttpChallengeHeaderValueStaticsInterface interface {
	win32.IInspectableInterface
	Parse(input string) *IHttpChallengeHeaderValue
	TryParse(input string, challengeHeaderValue *IHttpChallengeHeaderValue) bool
}

type IHttpChallengeHeaderValueStaticsVtbl struct {
	win32.IInspectableVtbl
	Parse    uintptr
	TryParse uintptr
}

type IHttpChallengeHeaderValueStatics struct {
	win32.IInspectable
}

func (this *IHttpChallengeHeaderValueStatics) Vtbl() *IHttpChallengeHeaderValueStaticsVtbl {
	return (*IHttpChallengeHeaderValueStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpChallengeHeaderValueStatics) Parse(input string) *IHttpChallengeHeaderValue {
	var _result *IHttpChallengeHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Parse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpChallengeHeaderValueStatics) TryParse(input string, challengeHeaderValue *IHttpChallengeHeaderValue) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(challengeHeaderValue)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// CB4AF27A-4E90-45EB-8DCD-FD1408F4C44F
var IID_IHttpConnectionOptionHeaderValue = syscall.GUID{0xCB4AF27A, 0x4E90, 0x45EB,
	[8]byte{0x8D, 0xCD, 0xFD, 0x14, 0x08, 0xF4, 0xC4, 0x4F}}

type IHttpConnectionOptionHeaderValueInterface interface {
	win32.IInspectableInterface
	Get_Token() string
}

type IHttpConnectionOptionHeaderValueVtbl struct {
	win32.IInspectableVtbl
	Get_Token uintptr
}

type IHttpConnectionOptionHeaderValue struct {
	win32.IInspectable
}

func (this *IHttpConnectionOptionHeaderValue) Vtbl() *IHttpConnectionOptionHeaderValueVtbl {
	return (*IHttpConnectionOptionHeaderValueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpConnectionOptionHeaderValue) Get_Token() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Token, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// E4F56C1D-5142-4E00-8E0F-019509337629
var IID_IHttpConnectionOptionHeaderValueCollection = syscall.GUID{0xE4F56C1D, 0x5142, 0x4E00,
	[8]byte{0x8E, 0x0F, 0x01, 0x95, 0x09, 0x33, 0x76, 0x29}}

type IHttpConnectionOptionHeaderValueCollectionInterface interface {
	win32.IInspectableInterface
	ParseAdd(input string)
	TryParseAdd(input string) bool
}

type IHttpConnectionOptionHeaderValueCollectionVtbl struct {
	win32.IInspectableVtbl
	ParseAdd    uintptr
	TryParseAdd uintptr
}

type IHttpConnectionOptionHeaderValueCollection struct {
	win32.IInspectable
}

func (this *IHttpConnectionOptionHeaderValueCollection) Vtbl() *IHttpConnectionOptionHeaderValueCollectionVtbl {
	return (*IHttpConnectionOptionHeaderValueCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpConnectionOptionHeaderValueCollection) ParseAdd(input string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ParseAdd, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr)
	_ = _hr
}

func (this *IHttpConnectionOptionHeaderValueCollection) TryParseAdd(input string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParseAdd, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D93CCC1E-0B7D-4C3F-A58D-A2A1BDEABC0A
var IID_IHttpConnectionOptionHeaderValueFactory = syscall.GUID{0xD93CCC1E, 0x0B7D, 0x4C3F,
	[8]byte{0xA5, 0x8D, 0xA2, 0xA1, 0xBD, 0xEA, 0xBC, 0x0A}}

type IHttpConnectionOptionHeaderValueFactoryInterface interface {
	win32.IInspectableInterface
	Create(token string) *IHttpConnectionOptionHeaderValue
}

type IHttpConnectionOptionHeaderValueFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IHttpConnectionOptionHeaderValueFactory struct {
	win32.IInspectable
}

func (this *IHttpConnectionOptionHeaderValueFactory) Vtbl() *IHttpConnectionOptionHeaderValueFactoryVtbl {
	return (*IHttpConnectionOptionHeaderValueFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpConnectionOptionHeaderValueFactory) Create(token string) *IHttpConnectionOptionHeaderValue {
	var _result *IHttpConnectionOptionHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(token).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AAA75D37-A946-4B1F-85AF-48B68B3C50BD
var IID_IHttpConnectionOptionHeaderValueStatics = syscall.GUID{0xAAA75D37, 0xA946, 0x4B1F,
	[8]byte{0x85, 0xAF, 0x48, 0xB6, 0x8B, 0x3C, 0x50, 0xBD}}

type IHttpConnectionOptionHeaderValueStaticsInterface interface {
	win32.IInspectableInterface
	Parse(input string) *IHttpConnectionOptionHeaderValue
	TryParse(input string, connectionOptionHeaderValue *IHttpConnectionOptionHeaderValue) bool
}

type IHttpConnectionOptionHeaderValueStaticsVtbl struct {
	win32.IInspectableVtbl
	Parse    uintptr
	TryParse uintptr
}

type IHttpConnectionOptionHeaderValueStatics struct {
	win32.IInspectable
}

func (this *IHttpConnectionOptionHeaderValueStatics) Vtbl() *IHttpConnectionOptionHeaderValueStaticsVtbl {
	return (*IHttpConnectionOptionHeaderValueStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpConnectionOptionHeaderValueStatics) Parse(input string) *IHttpConnectionOptionHeaderValue {
	var _result *IHttpConnectionOptionHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Parse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpConnectionOptionHeaderValueStatics) TryParse(input string, connectionOptionHeaderValue *IHttpConnectionOptionHeaderValue) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(connectionOptionHeaderValue)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// BCF7F92A-9376-4D85-BCCC-9F4F9ACAB434
var IID_IHttpContentCodingHeaderValue = syscall.GUID{0xBCF7F92A, 0x9376, 0x4D85,
	[8]byte{0xBC, 0xCC, 0x9F, 0x4F, 0x9A, 0xCA, 0xB4, 0x34}}

type IHttpContentCodingHeaderValueInterface interface {
	win32.IInspectableInterface
	Get_ContentCoding() string
}

type IHttpContentCodingHeaderValueVtbl struct {
	win32.IInspectableVtbl
	Get_ContentCoding uintptr
}

type IHttpContentCodingHeaderValue struct {
	win32.IInspectable
}

func (this *IHttpContentCodingHeaderValue) Vtbl() *IHttpContentCodingHeaderValueVtbl {
	return (*IHttpContentCodingHeaderValueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpContentCodingHeaderValue) Get_ContentCoding() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentCoding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 7D221721-A6DB-436E-8E83-91596192819C
var IID_IHttpContentCodingHeaderValueCollection = syscall.GUID{0x7D221721, 0xA6DB, 0x436E,
	[8]byte{0x8E, 0x83, 0x91, 0x59, 0x61, 0x92, 0x81, 0x9C}}

type IHttpContentCodingHeaderValueCollectionInterface interface {
	win32.IInspectableInterface
	ParseAdd(input string)
	TryParseAdd(input string) bool
}

type IHttpContentCodingHeaderValueCollectionVtbl struct {
	win32.IInspectableVtbl
	ParseAdd    uintptr
	TryParseAdd uintptr
}

type IHttpContentCodingHeaderValueCollection struct {
	win32.IInspectable
}

func (this *IHttpContentCodingHeaderValueCollection) Vtbl() *IHttpContentCodingHeaderValueCollectionVtbl {
	return (*IHttpContentCodingHeaderValueCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpContentCodingHeaderValueCollection) ParseAdd(input string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ParseAdd, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr)
	_ = _hr
}

func (this *IHttpContentCodingHeaderValueCollection) TryParseAdd(input string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParseAdd, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C53D2BD7-332B-4350-8510-2E67A2289A5A
var IID_IHttpContentCodingHeaderValueFactory = syscall.GUID{0xC53D2BD7, 0x332B, 0x4350,
	[8]byte{0x85, 0x10, 0x2E, 0x67, 0xA2, 0x28, 0x9A, 0x5A}}

type IHttpContentCodingHeaderValueFactoryInterface interface {
	win32.IInspectableInterface
	Create(contentCoding string) *IHttpContentCodingHeaderValue
}

type IHttpContentCodingHeaderValueFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IHttpContentCodingHeaderValueFactory struct {
	win32.IInspectable
}

func (this *IHttpContentCodingHeaderValueFactory) Vtbl() *IHttpContentCodingHeaderValueFactoryVtbl {
	return (*IHttpContentCodingHeaderValueFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpContentCodingHeaderValueFactory) Create(contentCoding string) *IHttpContentCodingHeaderValue {
	var _result *IHttpContentCodingHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(contentCoding).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 94D8602E-F9BF-42F7-AA46-ED272A41E212
var IID_IHttpContentCodingHeaderValueStatics = syscall.GUID{0x94D8602E, 0xF9BF, 0x42F7,
	[8]byte{0xAA, 0x46, 0xED, 0x27, 0x2A, 0x41, 0xE2, 0x12}}

type IHttpContentCodingHeaderValueStaticsInterface interface {
	win32.IInspectableInterface
	Parse(input string) *IHttpContentCodingHeaderValue
	TryParse(input string, contentCodingHeaderValue *IHttpContentCodingHeaderValue) bool
}

type IHttpContentCodingHeaderValueStaticsVtbl struct {
	win32.IInspectableVtbl
	Parse    uintptr
	TryParse uintptr
}

type IHttpContentCodingHeaderValueStatics struct {
	win32.IInspectable
}

func (this *IHttpContentCodingHeaderValueStatics) Vtbl() *IHttpContentCodingHeaderValueStaticsVtbl {
	return (*IHttpContentCodingHeaderValueStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpContentCodingHeaderValueStatics) Parse(input string) *IHttpContentCodingHeaderValue {
	var _result *IHttpContentCodingHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Parse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpContentCodingHeaderValueStatics) TryParse(input string, contentCodingHeaderValue *IHttpContentCodingHeaderValue) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(contentCodingHeaderValue)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 94531CD5-8B13-4D73-8651-F76B38F88495
var IID_IHttpContentCodingWithQualityHeaderValue = syscall.GUID{0x94531CD5, 0x8B13, 0x4D73,
	[8]byte{0x86, 0x51, 0xF7, 0x6B, 0x38, 0xF8, 0x84, 0x95}}

type IHttpContentCodingWithQualityHeaderValueInterface interface {
	win32.IInspectableInterface
	Get_ContentCoding() string
	Get_Quality() *IReference[float64]
}

type IHttpContentCodingWithQualityHeaderValueVtbl struct {
	win32.IInspectableVtbl
	Get_ContentCoding uintptr
	Get_Quality       uintptr
}

type IHttpContentCodingWithQualityHeaderValue struct {
	win32.IInspectable
}

func (this *IHttpContentCodingWithQualityHeaderValue) Vtbl() *IHttpContentCodingWithQualityHeaderValueVtbl {
	return (*IHttpContentCodingWithQualityHeaderValueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpContentCodingWithQualityHeaderValue) Get_ContentCoding() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentCoding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHttpContentCodingWithQualityHeaderValue) Get_Quality() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Quality, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7C0D753E-E899-4378-B5C8-412D820711CC
var IID_IHttpContentCodingWithQualityHeaderValueCollection = syscall.GUID{0x7C0D753E, 0xE899, 0x4378,
	[8]byte{0xB5, 0xC8, 0x41, 0x2D, 0x82, 0x07, 0x11, 0xCC}}

type IHttpContentCodingWithQualityHeaderValueCollectionInterface interface {
	win32.IInspectableInterface
	ParseAdd(input string)
	TryParseAdd(input string) bool
}

type IHttpContentCodingWithQualityHeaderValueCollectionVtbl struct {
	win32.IInspectableVtbl
	ParseAdd    uintptr
	TryParseAdd uintptr
}

type IHttpContentCodingWithQualityHeaderValueCollection struct {
	win32.IInspectable
}

func (this *IHttpContentCodingWithQualityHeaderValueCollection) Vtbl() *IHttpContentCodingWithQualityHeaderValueCollectionVtbl {
	return (*IHttpContentCodingWithQualityHeaderValueCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpContentCodingWithQualityHeaderValueCollection) ParseAdd(input string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ParseAdd, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr)
	_ = _hr
}

func (this *IHttpContentCodingWithQualityHeaderValueCollection) TryParseAdd(input string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParseAdd, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C45EEE1A-C553-46FC-ADE2-D75C1D53DF7B
var IID_IHttpContentCodingWithQualityHeaderValueFactory = syscall.GUID{0xC45EEE1A, 0xC553, 0x46FC,
	[8]byte{0xAD, 0xE2, 0xD7, 0x5C, 0x1D, 0x53, 0xDF, 0x7B}}

type IHttpContentCodingWithQualityHeaderValueFactoryInterface interface {
	win32.IInspectableInterface
	CreateFromValue(contentCoding string) *IHttpContentCodingWithQualityHeaderValue
	CreateFromValueWithQuality(contentCoding string, quality float64) *IHttpContentCodingWithQualityHeaderValue
}

type IHttpContentCodingWithQualityHeaderValueFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateFromValue            uintptr
	CreateFromValueWithQuality uintptr
}

type IHttpContentCodingWithQualityHeaderValueFactory struct {
	win32.IInspectable
}

func (this *IHttpContentCodingWithQualityHeaderValueFactory) Vtbl() *IHttpContentCodingWithQualityHeaderValueFactoryVtbl {
	return (*IHttpContentCodingWithQualityHeaderValueFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpContentCodingWithQualityHeaderValueFactory) CreateFromValue(contentCoding string) *IHttpContentCodingWithQualityHeaderValue {
	var _result *IHttpContentCodingWithQualityHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromValue, uintptr(unsafe.Pointer(this)), NewHStr(contentCoding).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpContentCodingWithQualityHeaderValueFactory) CreateFromValueWithQuality(contentCoding string, quality float64) *IHttpContentCodingWithQualityHeaderValue {
	var _result *IHttpContentCodingWithQualityHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromValueWithQuality, uintptr(unsafe.Pointer(this)), NewHStr(contentCoding).Ptr, uintptr(quality), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E8C9357C-8F89-4801-8E75-4C9ABFC3DE71
var IID_IHttpContentCodingWithQualityHeaderValueStatics = syscall.GUID{0xE8C9357C, 0x8F89, 0x4801,
	[8]byte{0x8E, 0x75, 0x4C, 0x9A, 0xBF, 0xC3, 0xDE, 0x71}}

type IHttpContentCodingWithQualityHeaderValueStaticsInterface interface {
	win32.IInspectableInterface
	Parse(input string) *IHttpContentCodingWithQualityHeaderValue
	TryParse(input string, contentCodingWithQualityHeaderValue *IHttpContentCodingWithQualityHeaderValue) bool
}

type IHttpContentCodingWithQualityHeaderValueStaticsVtbl struct {
	win32.IInspectableVtbl
	Parse    uintptr
	TryParse uintptr
}

type IHttpContentCodingWithQualityHeaderValueStatics struct {
	win32.IInspectable
}

func (this *IHttpContentCodingWithQualityHeaderValueStatics) Vtbl() *IHttpContentCodingWithQualityHeaderValueStaticsVtbl {
	return (*IHttpContentCodingWithQualityHeaderValueStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpContentCodingWithQualityHeaderValueStatics) Parse(input string) *IHttpContentCodingWithQualityHeaderValue {
	var _result *IHttpContentCodingWithQualityHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Parse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpContentCodingWithQualityHeaderValueStatics) TryParse(input string, contentCodingWithQualityHeaderValue *IHttpContentCodingWithQualityHeaderValue) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(contentCodingWithQualityHeaderValue)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F2A2EEDC-2629-4B49-9908-96A168E9365E
var IID_IHttpContentDispositionHeaderValue = syscall.GUID{0xF2A2EEDC, 0x2629, 0x4B49,
	[8]byte{0x99, 0x08, 0x96, 0xA1, 0x68, 0xE9, 0x36, 0x5E}}

type IHttpContentDispositionHeaderValueInterface interface {
	win32.IInspectableInterface
	Get_DispositionType() string
	Put_DispositionType(value string)
	Get_FileName() string
	Put_FileName(value string)
	Get_FileNameStar() string
	Put_FileNameStar(value string)
	Get_Name() string
	Put_Name(value string)
	Get_Parameters() *IVector[*IHttpNameValueHeaderValue]
	Get_Size() *IReference[uint64]
	Put_Size(value *IReference[uint64])
}

type IHttpContentDispositionHeaderValueVtbl struct {
	win32.IInspectableVtbl
	Get_DispositionType uintptr
	Put_DispositionType uintptr
	Get_FileName        uintptr
	Put_FileName        uintptr
	Get_FileNameStar    uintptr
	Put_FileNameStar    uintptr
	Get_Name            uintptr
	Put_Name            uintptr
	Get_Parameters      uintptr
	Get_Size            uintptr
	Put_Size            uintptr
}

type IHttpContentDispositionHeaderValue struct {
	win32.IInspectable
}

func (this *IHttpContentDispositionHeaderValue) Vtbl() *IHttpContentDispositionHeaderValueVtbl {
	return (*IHttpContentDispositionHeaderValueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpContentDispositionHeaderValue) Get_DispositionType() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DispositionType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHttpContentDispositionHeaderValue) Put_DispositionType(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DispositionType, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IHttpContentDispositionHeaderValue) Get_FileName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FileName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHttpContentDispositionHeaderValue) Put_FileName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FileName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IHttpContentDispositionHeaderValue) Get_FileNameStar() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FileNameStar, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHttpContentDispositionHeaderValue) Put_FileNameStar(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FileNameStar, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IHttpContentDispositionHeaderValue) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHttpContentDispositionHeaderValue) Put_Name(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Name, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IHttpContentDispositionHeaderValue) Get_Parameters() *IVector[*IHttpNameValueHeaderValue] {
	var _result *IVector[*IHttpNameValueHeaderValue]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Parameters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpContentDispositionHeaderValue) Get_Size() *IReference[uint64] {
	var _result *IReference[uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Size, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpContentDispositionHeaderValue) Put_Size(value *IReference[uint64]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Size, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 9915BBC4-456C-4E81-8295-B2AB3CBCF545
var IID_IHttpContentDispositionHeaderValueFactory = syscall.GUID{0x9915BBC4, 0x456C, 0x4E81,
	[8]byte{0x82, 0x95, 0xB2, 0xAB, 0x3C, 0xBC, 0xF5, 0x45}}

type IHttpContentDispositionHeaderValueFactoryInterface interface {
	win32.IInspectableInterface
	Create(dispositionType string) *IHttpContentDispositionHeaderValue
}

type IHttpContentDispositionHeaderValueFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IHttpContentDispositionHeaderValueFactory struct {
	win32.IInspectable
}

func (this *IHttpContentDispositionHeaderValueFactory) Vtbl() *IHttpContentDispositionHeaderValueFactoryVtbl {
	return (*IHttpContentDispositionHeaderValueFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpContentDispositionHeaderValueFactory) Create(dispositionType string) *IHttpContentDispositionHeaderValue {
	var _result *IHttpContentDispositionHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(dispositionType).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 29C56067-5A37-46E4-B074-C5177D69CA66
var IID_IHttpContentDispositionHeaderValueStatics = syscall.GUID{0x29C56067, 0x5A37, 0x46E4,
	[8]byte{0xB0, 0x74, 0xC5, 0x17, 0x7D, 0x69, 0xCA, 0x66}}

type IHttpContentDispositionHeaderValueStaticsInterface interface {
	win32.IInspectableInterface
	Parse(input string) *IHttpContentDispositionHeaderValue
	TryParse(input string, contentDispositionHeaderValue *IHttpContentDispositionHeaderValue) bool
}

type IHttpContentDispositionHeaderValueStaticsVtbl struct {
	win32.IInspectableVtbl
	Parse    uintptr
	TryParse uintptr
}

type IHttpContentDispositionHeaderValueStatics struct {
	win32.IInspectable
}

func (this *IHttpContentDispositionHeaderValueStatics) Vtbl() *IHttpContentDispositionHeaderValueStaticsVtbl {
	return (*IHttpContentDispositionHeaderValueStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpContentDispositionHeaderValueStatics) Parse(input string) *IHttpContentDispositionHeaderValue {
	var _result *IHttpContentDispositionHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Parse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpContentDispositionHeaderValueStatics) TryParse(input string, contentDispositionHeaderValue *IHttpContentDispositionHeaderValue) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(contentDispositionHeaderValue)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 40612A44-47AE-4B7E-9124-69628B64AA18
var IID_IHttpContentHeaderCollection = syscall.GUID{0x40612A44, 0x47AE, 0x4B7E,
	[8]byte{0x91, 0x24, 0x69, 0x62, 0x8B, 0x64, 0xAA, 0x18}}

type IHttpContentHeaderCollectionInterface interface {
	win32.IInspectableInterface
	Get_ContentDisposition() *IHttpContentDispositionHeaderValue
	Put_ContentDisposition(value *IHttpContentDispositionHeaderValue)
	Get_ContentEncoding() *IHttpContentCodingHeaderValueCollection
	Get_ContentLanguage() *IHttpLanguageHeaderValueCollection
	Get_ContentLength() *IReference[uint64]
	Put_ContentLength(value *IReference[uint64])
	Get_ContentLocation() *IUriRuntimeClass
	Put_ContentLocation(value *IUriRuntimeClass)
	Get_ContentMD5() *IBuffer
	Put_ContentMD5(value *IBuffer)
	Get_ContentRange() *IHttpContentRangeHeaderValue
	Put_ContentRange(value *IHttpContentRangeHeaderValue)
	Get_ContentType() *IHttpMediaTypeHeaderValue
	Put_ContentType(value *IHttpMediaTypeHeaderValue)
	Get_Expires() *IReference[DateTime]
	Put_Expires(value *IReference[DateTime])
	Get_LastModified() *IReference[DateTime]
	Put_LastModified(value *IReference[DateTime])
	Append(name string, value string)
	TryAppendWithoutValidation(name string, value string) bool
}

type IHttpContentHeaderCollectionVtbl struct {
	win32.IInspectableVtbl
	Get_ContentDisposition     uintptr
	Put_ContentDisposition     uintptr
	Get_ContentEncoding        uintptr
	Get_ContentLanguage        uintptr
	Get_ContentLength          uintptr
	Put_ContentLength          uintptr
	Get_ContentLocation        uintptr
	Put_ContentLocation        uintptr
	Get_ContentMD5             uintptr
	Put_ContentMD5             uintptr
	Get_ContentRange           uintptr
	Put_ContentRange           uintptr
	Get_ContentType            uintptr
	Put_ContentType            uintptr
	Get_Expires                uintptr
	Put_Expires                uintptr
	Get_LastModified           uintptr
	Put_LastModified           uintptr
	Append                     uintptr
	TryAppendWithoutValidation uintptr
}

type IHttpContentHeaderCollection struct {
	win32.IInspectable
}

func (this *IHttpContentHeaderCollection) Vtbl() *IHttpContentHeaderCollectionVtbl {
	return (*IHttpContentHeaderCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpContentHeaderCollection) Get_ContentDisposition() *IHttpContentDispositionHeaderValue {
	var _result *IHttpContentDispositionHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentDisposition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpContentHeaderCollection) Put_ContentDisposition(value *IHttpContentDispositionHeaderValue) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ContentDisposition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpContentHeaderCollection) Get_ContentEncoding() *IHttpContentCodingHeaderValueCollection {
	var _result *IHttpContentCodingHeaderValueCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentEncoding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpContentHeaderCollection) Get_ContentLanguage() *IHttpLanguageHeaderValueCollection {
	var _result *IHttpLanguageHeaderValueCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentLanguage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpContentHeaderCollection) Get_ContentLength() *IReference[uint64] {
	var _result *IReference[uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpContentHeaderCollection) Put_ContentLength(value *IReference[uint64]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ContentLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpContentHeaderCollection) Get_ContentLocation() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentLocation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpContentHeaderCollection) Put_ContentLocation(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ContentLocation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpContentHeaderCollection) Get_ContentMD5() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentMD5, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpContentHeaderCollection) Put_ContentMD5(value *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ContentMD5, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpContentHeaderCollection) Get_ContentRange() *IHttpContentRangeHeaderValue {
	var _result *IHttpContentRangeHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentRange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpContentHeaderCollection) Put_ContentRange(value *IHttpContentRangeHeaderValue) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ContentRange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpContentHeaderCollection) Get_ContentType() *IHttpMediaTypeHeaderValue {
	var _result *IHttpMediaTypeHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpContentHeaderCollection) Put_ContentType(value *IHttpMediaTypeHeaderValue) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ContentType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpContentHeaderCollection) Get_Expires() *IReference[DateTime] {
	var _result *IReference[DateTime]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Expires, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpContentHeaderCollection) Put_Expires(value *IReference[DateTime]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Expires, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpContentHeaderCollection) Get_LastModified() *IReference[DateTime] {
	var _result *IReference[DateTime]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LastModified, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpContentHeaderCollection) Put_LastModified(value *IReference[DateTime]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LastModified, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpContentHeaderCollection) Append(name string, value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Append, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, NewHStr(value).Ptr)
	_ = _hr
}

func (this *IHttpContentHeaderCollection) TryAppendWithoutValidation(name string, value string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryAppendWithoutValidation, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, NewHStr(value).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 04D967D3-A4F6-495C-9530-8579FCBA8AA9
var IID_IHttpContentRangeHeaderValue = syscall.GUID{0x04D967D3, 0xA4F6, 0x495C,
	[8]byte{0x95, 0x30, 0x85, 0x79, 0xFC, 0xBA, 0x8A, 0xA9}}

type IHttpContentRangeHeaderValueInterface interface {
	win32.IInspectableInterface
	Get_FirstBytePosition() *IReference[uint64]
	Get_LastBytePosition() *IReference[uint64]
	Get_Length() *IReference[uint64]
	Get_Unit() string
	Put_Unit(value string)
}

type IHttpContentRangeHeaderValueVtbl struct {
	win32.IInspectableVtbl
	Get_FirstBytePosition uintptr
	Get_LastBytePosition  uintptr
	Get_Length            uintptr
	Get_Unit              uintptr
	Put_Unit              uintptr
}

type IHttpContentRangeHeaderValue struct {
	win32.IInspectable
}

func (this *IHttpContentRangeHeaderValue) Vtbl() *IHttpContentRangeHeaderValueVtbl {
	return (*IHttpContentRangeHeaderValueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpContentRangeHeaderValue) Get_FirstBytePosition() *IReference[uint64] {
	var _result *IReference[uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FirstBytePosition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpContentRangeHeaderValue) Get_LastBytePosition() *IReference[uint64] {
	var _result *IReference[uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LastBytePosition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpContentRangeHeaderValue) Get_Length() *IReference[uint64] {
	var _result *IReference[uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Length, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpContentRangeHeaderValue) Get_Unit() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Unit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHttpContentRangeHeaderValue) Put_Unit(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Unit, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// 3F5BD691-A03C-4456-9A6F-EF27ECD03CAE
var IID_IHttpContentRangeHeaderValueFactory = syscall.GUID{0x3F5BD691, 0xA03C, 0x4456,
	[8]byte{0x9A, 0x6F, 0xEF, 0x27, 0xEC, 0xD0, 0x3C, 0xAE}}

type IHttpContentRangeHeaderValueFactoryInterface interface {
	win32.IInspectableInterface
	CreateFromLength(length uint64) *IHttpContentRangeHeaderValue
	CreateFromRange(from uint64, to uint64) *IHttpContentRangeHeaderValue
	CreateFromRangeWithLength(from uint64, to uint64, length uint64) *IHttpContentRangeHeaderValue
}

type IHttpContentRangeHeaderValueFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateFromLength          uintptr
	CreateFromRange           uintptr
	CreateFromRangeWithLength uintptr
}

type IHttpContentRangeHeaderValueFactory struct {
	win32.IInspectable
}

func (this *IHttpContentRangeHeaderValueFactory) Vtbl() *IHttpContentRangeHeaderValueFactoryVtbl {
	return (*IHttpContentRangeHeaderValueFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpContentRangeHeaderValueFactory) CreateFromLength(length uint64) *IHttpContentRangeHeaderValue {
	var _result *IHttpContentRangeHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromLength, uintptr(unsafe.Pointer(this)), uintptr(length), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpContentRangeHeaderValueFactory) CreateFromRange(from uint64, to uint64) *IHttpContentRangeHeaderValue {
	var _result *IHttpContentRangeHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromRange, uintptr(unsafe.Pointer(this)), uintptr(from), uintptr(to), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpContentRangeHeaderValueFactory) CreateFromRangeWithLength(from uint64, to uint64, length uint64) *IHttpContentRangeHeaderValue {
	var _result *IHttpContentRangeHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromRangeWithLength, uintptr(unsafe.Pointer(this)), uintptr(from), uintptr(to), uintptr(length), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 80A346CA-174C-4FAE-821C-134CD294AA38
var IID_IHttpContentRangeHeaderValueStatics = syscall.GUID{0x80A346CA, 0x174C, 0x4FAE,
	[8]byte{0x82, 0x1C, 0x13, 0x4C, 0xD2, 0x94, 0xAA, 0x38}}

type IHttpContentRangeHeaderValueStaticsInterface interface {
	win32.IInspectableInterface
	Parse(input string) *IHttpContentRangeHeaderValue
	TryParse(input string, contentRangeHeaderValue *IHttpContentRangeHeaderValue) bool
}

type IHttpContentRangeHeaderValueStaticsVtbl struct {
	win32.IInspectableVtbl
	Parse    uintptr
	TryParse uintptr
}

type IHttpContentRangeHeaderValueStatics struct {
	win32.IInspectable
}

func (this *IHttpContentRangeHeaderValueStatics) Vtbl() *IHttpContentRangeHeaderValueStaticsVtbl {
	return (*IHttpContentRangeHeaderValueStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpContentRangeHeaderValueStatics) Parse(input string) *IHttpContentRangeHeaderValue {
	var _result *IHttpContentRangeHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Parse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpContentRangeHeaderValueStatics) TryParse(input string, contentRangeHeaderValue *IHttpContentRangeHeaderValue) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(contentRangeHeaderValue)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// CBD46217-4B29-412B-BD90-B3D814AB8E1B
var IID_IHttpCookiePairHeaderValue = syscall.GUID{0xCBD46217, 0x4B29, 0x412B,
	[8]byte{0xBD, 0x90, 0xB3, 0xD8, 0x14, 0xAB, 0x8E, 0x1B}}

type IHttpCookiePairHeaderValueInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Get_Value() string
	Put_Value(value string)
}

type IHttpCookiePairHeaderValueVtbl struct {
	win32.IInspectableVtbl
	Get_Name  uintptr
	Get_Value uintptr
	Put_Value uintptr
}

type IHttpCookiePairHeaderValue struct {
	win32.IInspectable
}

func (this *IHttpCookiePairHeaderValue) Vtbl() *IHttpCookiePairHeaderValueVtbl {
	return (*IHttpCookiePairHeaderValueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpCookiePairHeaderValue) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHttpCookiePairHeaderValue) Get_Value() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHttpCookiePairHeaderValue) Put_Value(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Value, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// F3F44350-581E-4ECC-9F59-E507D04F06E6
var IID_IHttpCookiePairHeaderValueCollection = syscall.GUID{0xF3F44350, 0x581E, 0x4ECC,
	[8]byte{0x9F, 0x59, 0xE5, 0x07, 0xD0, 0x4F, 0x06, 0xE6}}

type IHttpCookiePairHeaderValueCollectionInterface interface {
	win32.IInspectableInterface
	ParseAdd(input string)
	TryParseAdd(input string) bool
}

type IHttpCookiePairHeaderValueCollectionVtbl struct {
	win32.IInspectableVtbl
	ParseAdd    uintptr
	TryParseAdd uintptr
}

type IHttpCookiePairHeaderValueCollection struct {
	win32.IInspectable
}

func (this *IHttpCookiePairHeaderValueCollection) Vtbl() *IHttpCookiePairHeaderValueCollectionVtbl {
	return (*IHttpCookiePairHeaderValueCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpCookiePairHeaderValueCollection) ParseAdd(input string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ParseAdd, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr)
	_ = _hr
}

func (this *IHttpCookiePairHeaderValueCollection) TryParseAdd(input string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParseAdd, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 635E326F-146F-4F56-AA21-2CB7D6D58B1E
var IID_IHttpCookiePairHeaderValueFactory = syscall.GUID{0x635E326F, 0x146F, 0x4F56,
	[8]byte{0xAA, 0x21, 0x2C, 0xB7, 0xD6, 0xD5, 0x8B, 0x1E}}

type IHttpCookiePairHeaderValueFactoryInterface interface {
	win32.IInspectableInterface
	CreateFromName(name string) *IHttpCookiePairHeaderValue
	CreateFromNameWithValue(name string, value string) *IHttpCookiePairHeaderValue
}

type IHttpCookiePairHeaderValueFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateFromName          uintptr
	CreateFromNameWithValue uintptr
}

type IHttpCookiePairHeaderValueFactory struct {
	win32.IInspectable
}

func (this *IHttpCookiePairHeaderValueFactory) Vtbl() *IHttpCookiePairHeaderValueFactoryVtbl {
	return (*IHttpCookiePairHeaderValueFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpCookiePairHeaderValueFactory) CreateFromName(name string) *IHttpCookiePairHeaderValue {
	var _result *IHttpCookiePairHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromName, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpCookiePairHeaderValueFactory) CreateFromNameWithValue(name string, value string) *IHttpCookiePairHeaderValue {
	var _result *IHttpCookiePairHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromNameWithValue, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, NewHStr(value).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6E866D48-06AF-4462-8158-99388D5DCA81
var IID_IHttpCookiePairHeaderValueStatics = syscall.GUID{0x6E866D48, 0x06AF, 0x4462,
	[8]byte{0x81, 0x58, 0x99, 0x38, 0x8D, 0x5D, 0xCA, 0x81}}

type IHttpCookiePairHeaderValueStaticsInterface interface {
	win32.IInspectableInterface
	Parse(input string) *IHttpCookiePairHeaderValue
	TryParse(input string, cookiePairHeaderValue *IHttpCookiePairHeaderValue) bool
}

type IHttpCookiePairHeaderValueStaticsVtbl struct {
	win32.IInspectableVtbl
	Parse    uintptr
	TryParse uintptr
}

type IHttpCookiePairHeaderValueStatics struct {
	win32.IInspectable
}

func (this *IHttpCookiePairHeaderValueStatics) Vtbl() *IHttpCookiePairHeaderValueStaticsVtbl {
	return (*IHttpCookiePairHeaderValueStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpCookiePairHeaderValueStatics) Parse(input string) *IHttpCookiePairHeaderValue {
	var _result *IHttpCookiePairHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Parse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpCookiePairHeaderValueStatics) TryParse(input string, cookiePairHeaderValue *IHttpCookiePairHeaderValue) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(cookiePairHeaderValue)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C34CC3CB-542E-4177-A6C7-B674CE193FBF
var IID_IHttpCredentialsHeaderValue = syscall.GUID{0xC34CC3CB, 0x542E, 0x4177,
	[8]byte{0xA6, 0xC7, 0xB6, 0x74, 0xCE, 0x19, 0x3F, 0xBF}}

type IHttpCredentialsHeaderValueInterface interface {
	win32.IInspectableInterface
	Get_Parameters() *IVector[*IHttpNameValueHeaderValue]
	Get_Scheme() string
	Get_Token() string
}

type IHttpCredentialsHeaderValueVtbl struct {
	win32.IInspectableVtbl
	Get_Parameters uintptr
	Get_Scheme     uintptr
	Get_Token      uintptr
}

type IHttpCredentialsHeaderValue struct {
	win32.IInspectable
}

func (this *IHttpCredentialsHeaderValue) Vtbl() *IHttpCredentialsHeaderValueVtbl {
	return (*IHttpCredentialsHeaderValueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpCredentialsHeaderValue) Get_Parameters() *IVector[*IHttpNameValueHeaderValue] {
	var _result *IVector[*IHttpNameValueHeaderValue]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Parameters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpCredentialsHeaderValue) Get_Scheme() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Scheme, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHttpCredentialsHeaderValue) Get_Token() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Token, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// F21D9E91-4D1C-4182-BFD1-34470A62F950
var IID_IHttpCredentialsHeaderValueFactory = syscall.GUID{0xF21D9E91, 0x4D1C, 0x4182,
	[8]byte{0xBF, 0xD1, 0x34, 0x47, 0x0A, 0x62, 0xF9, 0x50}}

type IHttpCredentialsHeaderValueFactoryInterface interface {
	win32.IInspectableInterface
	CreateFromScheme(scheme string) *IHttpCredentialsHeaderValue
	CreateFromSchemeWithToken(scheme string, token string) *IHttpCredentialsHeaderValue
}

type IHttpCredentialsHeaderValueFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateFromScheme          uintptr
	CreateFromSchemeWithToken uintptr
}

type IHttpCredentialsHeaderValueFactory struct {
	win32.IInspectable
}

func (this *IHttpCredentialsHeaderValueFactory) Vtbl() *IHttpCredentialsHeaderValueFactoryVtbl {
	return (*IHttpCredentialsHeaderValueFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpCredentialsHeaderValueFactory) CreateFromScheme(scheme string) *IHttpCredentialsHeaderValue {
	var _result *IHttpCredentialsHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromScheme, uintptr(unsafe.Pointer(this)), NewHStr(scheme).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpCredentialsHeaderValueFactory) CreateFromSchemeWithToken(scheme string, token string) *IHttpCredentialsHeaderValue {
	var _result *IHttpCredentialsHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromSchemeWithToken, uintptr(unsafe.Pointer(this)), NewHStr(scheme).Ptr, NewHStr(token).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A69B2BE6-CE8C-4443-A35A-1B727B131036
var IID_IHttpCredentialsHeaderValueStatics = syscall.GUID{0xA69B2BE6, 0xCE8C, 0x4443,
	[8]byte{0xA3, 0x5A, 0x1B, 0x72, 0x7B, 0x13, 0x10, 0x36}}

type IHttpCredentialsHeaderValueStaticsInterface interface {
	win32.IInspectableInterface
	Parse(input string) *IHttpCredentialsHeaderValue
	TryParse(input string, credentialsHeaderValue *IHttpCredentialsHeaderValue) bool
}

type IHttpCredentialsHeaderValueStaticsVtbl struct {
	win32.IInspectableVtbl
	Parse    uintptr
	TryParse uintptr
}

type IHttpCredentialsHeaderValueStatics struct {
	win32.IInspectable
}

func (this *IHttpCredentialsHeaderValueStatics) Vtbl() *IHttpCredentialsHeaderValueStaticsVtbl {
	return (*IHttpCredentialsHeaderValueStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpCredentialsHeaderValueStatics) Parse(input string) *IHttpCredentialsHeaderValue {
	var _result *IHttpCredentialsHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Parse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpCredentialsHeaderValueStatics) TryParse(input string, credentialsHeaderValue *IHttpCredentialsHeaderValue) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(credentialsHeaderValue)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// EAFCAA6A-C4DC-49E2-A27D-043ADF5867A3
var IID_IHttpDateOrDeltaHeaderValue = syscall.GUID{0xEAFCAA6A, 0xC4DC, 0x49E2,
	[8]byte{0xA2, 0x7D, 0x04, 0x3A, 0xDF, 0x58, 0x67, 0xA3}}

type IHttpDateOrDeltaHeaderValueInterface interface {
	win32.IInspectableInterface
	Get_Date() *IReference[DateTime]
	Get_Delta() *IReference[TimeSpan]
}

type IHttpDateOrDeltaHeaderValueVtbl struct {
	win32.IInspectableVtbl
	Get_Date  uintptr
	Get_Delta uintptr
}

type IHttpDateOrDeltaHeaderValue struct {
	win32.IInspectable
}

func (this *IHttpDateOrDeltaHeaderValue) Vtbl() *IHttpDateOrDeltaHeaderValueVtbl {
	return (*IHttpDateOrDeltaHeaderValueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpDateOrDeltaHeaderValue) Get_Date() *IReference[DateTime] {
	var _result *IReference[DateTime]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Date, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpDateOrDeltaHeaderValue) Get_Delta() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Delta, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7C2659A8-6672-4E90-9A9A-F39766F7F576
var IID_IHttpDateOrDeltaHeaderValueStatics = syscall.GUID{0x7C2659A8, 0x6672, 0x4E90,
	[8]byte{0x9A, 0x9A, 0xF3, 0x97, 0x66, 0xF7, 0xF5, 0x76}}

type IHttpDateOrDeltaHeaderValueStaticsInterface interface {
	win32.IInspectableInterface
	Parse(input string) *IHttpDateOrDeltaHeaderValue
	TryParse(input string, dateOrDeltaHeaderValue *IHttpDateOrDeltaHeaderValue) bool
}

type IHttpDateOrDeltaHeaderValueStaticsVtbl struct {
	win32.IInspectableVtbl
	Parse    uintptr
	TryParse uintptr
}

type IHttpDateOrDeltaHeaderValueStatics struct {
	win32.IInspectable
}

func (this *IHttpDateOrDeltaHeaderValueStatics) Vtbl() *IHttpDateOrDeltaHeaderValueStaticsVtbl {
	return (*IHttpDateOrDeltaHeaderValueStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpDateOrDeltaHeaderValueStatics) Parse(input string) *IHttpDateOrDeltaHeaderValue {
	var _result *IHttpDateOrDeltaHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Parse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpDateOrDeltaHeaderValueStatics) TryParse(input string, dateOrDeltaHeaderValue *IHttpDateOrDeltaHeaderValue) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(dateOrDeltaHeaderValue)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4CE585CD-3A99-43AF-A2E6-EC232FEA9658
var IID_IHttpExpectationHeaderValue = syscall.GUID{0x4CE585CD, 0x3A99, 0x43AF,
	[8]byte{0xA2, 0xE6, 0xEC, 0x23, 0x2F, 0xEA, 0x96, 0x58}}

type IHttpExpectationHeaderValueInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Get_Value() string
	Put_Value(value string)
	Get_Parameters() *IVector[*IHttpNameValueHeaderValue]
}

type IHttpExpectationHeaderValueVtbl struct {
	win32.IInspectableVtbl
	Get_Name       uintptr
	Get_Value      uintptr
	Put_Value      uintptr
	Get_Parameters uintptr
}

type IHttpExpectationHeaderValue struct {
	win32.IInspectable
}

func (this *IHttpExpectationHeaderValue) Vtbl() *IHttpExpectationHeaderValueVtbl {
	return (*IHttpExpectationHeaderValueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpExpectationHeaderValue) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHttpExpectationHeaderValue) Get_Value() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHttpExpectationHeaderValue) Put_Value(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Value, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IHttpExpectationHeaderValue) Get_Parameters() *IVector[*IHttpNameValueHeaderValue] {
	var _result *IVector[*IHttpNameValueHeaderValue]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Parameters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E78521B3-A0E2-4AC4-9E66-79706CB9FD58
var IID_IHttpExpectationHeaderValueCollection = syscall.GUID{0xE78521B3, 0xA0E2, 0x4AC4,
	[8]byte{0x9E, 0x66, 0x79, 0x70, 0x6C, 0xB9, 0xFD, 0x58}}

type IHttpExpectationHeaderValueCollectionInterface interface {
	win32.IInspectableInterface
	ParseAdd(input string)
	TryParseAdd(input string) bool
}

type IHttpExpectationHeaderValueCollectionVtbl struct {
	win32.IInspectableVtbl
	ParseAdd    uintptr
	TryParseAdd uintptr
}

type IHttpExpectationHeaderValueCollection struct {
	win32.IInspectable
}

func (this *IHttpExpectationHeaderValueCollection) Vtbl() *IHttpExpectationHeaderValueCollectionVtbl {
	return (*IHttpExpectationHeaderValueCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpExpectationHeaderValueCollection) ParseAdd(input string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ParseAdd, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr)
	_ = _hr
}

func (this *IHttpExpectationHeaderValueCollection) TryParseAdd(input string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParseAdd, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4EA275CB-D53E-4868-8856-1E21A5030DC0
var IID_IHttpExpectationHeaderValueFactory = syscall.GUID{0x4EA275CB, 0xD53E, 0x4868,
	[8]byte{0x88, 0x56, 0x1E, 0x21, 0xA5, 0x03, 0x0D, 0xC0}}

type IHttpExpectationHeaderValueFactoryInterface interface {
	win32.IInspectableInterface
	CreateFromName(name string) *IHttpExpectationHeaderValue
	CreateFromNameWithValue(name string, value string) *IHttpExpectationHeaderValue
}

type IHttpExpectationHeaderValueFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateFromName          uintptr
	CreateFromNameWithValue uintptr
}

type IHttpExpectationHeaderValueFactory struct {
	win32.IInspectable
}

func (this *IHttpExpectationHeaderValueFactory) Vtbl() *IHttpExpectationHeaderValueFactoryVtbl {
	return (*IHttpExpectationHeaderValueFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpExpectationHeaderValueFactory) CreateFromName(name string) *IHttpExpectationHeaderValue {
	var _result *IHttpExpectationHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromName, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpExpectationHeaderValueFactory) CreateFromNameWithValue(name string, value string) *IHttpExpectationHeaderValue {
	var _result *IHttpExpectationHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromNameWithValue, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, NewHStr(value).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3019ABE2-CFE5-473B-A57F-FBA5B14EB257
var IID_IHttpExpectationHeaderValueStatics = syscall.GUID{0x3019ABE2, 0xCFE5, 0x473B,
	[8]byte{0xA5, 0x7F, 0xFB, 0xA5, 0xB1, 0x4E, 0xB2, 0x57}}

type IHttpExpectationHeaderValueStaticsInterface interface {
	win32.IInspectableInterface
	Parse(input string) *IHttpExpectationHeaderValue
	TryParse(input string, expectationHeaderValue *IHttpExpectationHeaderValue) bool
}

type IHttpExpectationHeaderValueStaticsVtbl struct {
	win32.IInspectableVtbl
	Parse    uintptr
	TryParse uintptr
}

type IHttpExpectationHeaderValueStatics struct {
	win32.IInspectable
}

func (this *IHttpExpectationHeaderValueStatics) Vtbl() *IHttpExpectationHeaderValueStaticsVtbl {
	return (*IHttpExpectationHeaderValueStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpExpectationHeaderValueStatics) Parse(input string) *IHttpExpectationHeaderValue {
	var _result *IHttpExpectationHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Parse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpExpectationHeaderValueStatics) TryParse(input string, expectationHeaderValue *IHttpExpectationHeaderValue) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(expectationHeaderValue)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 9EBD7CA3-8219-44F6-9902-8C56DFD3340C
var IID_IHttpLanguageHeaderValueCollection = syscall.GUID{0x9EBD7CA3, 0x8219, 0x44F6,
	[8]byte{0x99, 0x02, 0x8C, 0x56, 0xDF, 0xD3, 0x34, 0x0C}}

type IHttpLanguageHeaderValueCollectionInterface interface {
	win32.IInspectableInterface
	ParseAdd(input string)
	TryParseAdd(input string) bool
}

type IHttpLanguageHeaderValueCollectionVtbl struct {
	win32.IInspectableVtbl
	ParseAdd    uintptr
	TryParseAdd uintptr
}

type IHttpLanguageHeaderValueCollection struct {
	win32.IInspectable
}

func (this *IHttpLanguageHeaderValueCollection) Vtbl() *IHttpLanguageHeaderValueCollectionVtbl {
	return (*IHttpLanguageHeaderValueCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpLanguageHeaderValueCollection) ParseAdd(input string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ParseAdd, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr)
	_ = _hr
}

func (this *IHttpLanguageHeaderValueCollection) TryParseAdd(input string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParseAdd, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7256E102-0080-4DB4-A083-7DE7B2E5BA4C
var IID_IHttpLanguageRangeWithQualityHeaderValue = syscall.GUID{0x7256E102, 0x0080, 0x4DB4,
	[8]byte{0xA0, 0x83, 0x7D, 0xE7, 0xB2, 0xE5, 0xBA, 0x4C}}

type IHttpLanguageRangeWithQualityHeaderValueInterface interface {
	win32.IInspectableInterface
	Get_LanguageRange() string
	Get_Quality() *IReference[float64]
}

type IHttpLanguageRangeWithQualityHeaderValueVtbl struct {
	win32.IInspectableVtbl
	Get_LanguageRange uintptr
	Get_Quality       uintptr
}

type IHttpLanguageRangeWithQualityHeaderValue struct {
	win32.IInspectable
}

func (this *IHttpLanguageRangeWithQualityHeaderValue) Vtbl() *IHttpLanguageRangeWithQualityHeaderValueVtbl {
	return (*IHttpLanguageRangeWithQualityHeaderValueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpLanguageRangeWithQualityHeaderValue) Get_LanguageRange() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LanguageRange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHttpLanguageRangeWithQualityHeaderValue) Get_Quality() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Quality, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 885D5ABD-4B4F-480A-89CE-8AEDCEE6E3A0
var IID_IHttpLanguageRangeWithQualityHeaderValueCollection = syscall.GUID{0x885D5ABD, 0x4B4F, 0x480A,
	[8]byte{0x89, 0xCE, 0x8A, 0xED, 0xCE, 0xE6, 0xE3, 0xA0}}

type IHttpLanguageRangeWithQualityHeaderValueCollectionInterface interface {
	win32.IInspectableInterface
	ParseAdd(input string)
	TryParseAdd(input string) bool
}

type IHttpLanguageRangeWithQualityHeaderValueCollectionVtbl struct {
	win32.IInspectableVtbl
	ParseAdd    uintptr
	TryParseAdd uintptr
}

type IHttpLanguageRangeWithQualityHeaderValueCollection struct {
	win32.IInspectable
}

func (this *IHttpLanguageRangeWithQualityHeaderValueCollection) Vtbl() *IHttpLanguageRangeWithQualityHeaderValueCollectionVtbl {
	return (*IHttpLanguageRangeWithQualityHeaderValueCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpLanguageRangeWithQualityHeaderValueCollection) ParseAdd(input string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ParseAdd, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr)
	_ = _hr
}

func (this *IHttpLanguageRangeWithQualityHeaderValueCollection) TryParseAdd(input string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParseAdd, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7BB83970-780F-4C83-9FE4-DC3087F6BD55
var IID_IHttpLanguageRangeWithQualityHeaderValueFactory = syscall.GUID{0x7BB83970, 0x780F, 0x4C83,
	[8]byte{0x9F, 0xE4, 0xDC, 0x30, 0x87, 0xF6, 0xBD, 0x55}}

type IHttpLanguageRangeWithQualityHeaderValueFactoryInterface interface {
	win32.IInspectableInterface
	CreateFromLanguageRange(languageRange string) *IHttpLanguageRangeWithQualityHeaderValue
	CreateFromLanguageRangeWithQuality(languageRange string, quality float64) *IHttpLanguageRangeWithQualityHeaderValue
}

type IHttpLanguageRangeWithQualityHeaderValueFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateFromLanguageRange            uintptr
	CreateFromLanguageRangeWithQuality uintptr
}

type IHttpLanguageRangeWithQualityHeaderValueFactory struct {
	win32.IInspectable
}

func (this *IHttpLanguageRangeWithQualityHeaderValueFactory) Vtbl() *IHttpLanguageRangeWithQualityHeaderValueFactoryVtbl {
	return (*IHttpLanguageRangeWithQualityHeaderValueFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpLanguageRangeWithQualityHeaderValueFactory) CreateFromLanguageRange(languageRange string) *IHttpLanguageRangeWithQualityHeaderValue {
	var _result *IHttpLanguageRangeWithQualityHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromLanguageRange, uintptr(unsafe.Pointer(this)), NewHStr(languageRange).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpLanguageRangeWithQualityHeaderValueFactory) CreateFromLanguageRangeWithQuality(languageRange string, quality float64) *IHttpLanguageRangeWithQualityHeaderValue {
	var _result *IHttpLanguageRangeWithQualityHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromLanguageRangeWithQuality, uintptr(unsafe.Pointer(this)), NewHStr(languageRange).Ptr, uintptr(quality), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2541E146-F308-46F5-B695-42F54024EC68
var IID_IHttpLanguageRangeWithQualityHeaderValueStatics = syscall.GUID{0x2541E146, 0xF308, 0x46F5,
	[8]byte{0xB6, 0x95, 0x42, 0xF5, 0x40, 0x24, 0xEC, 0x68}}

type IHttpLanguageRangeWithQualityHeaderValueStaticsInterface interface {
	win32.IInspectableInterface
	Parse(input string) *IHttpLanguageRangeWithQualityHeaderValue
	TryParse(input string, languageRangeWithQualityHeaderValue *IHttpLanguageRangeWithQualityHeaderValue) bool
}

type IHttpLanguageRangeWithQualityHeaderValueStaticsVtbl struct {
	win32.IInspectableVtbl
	Parse    uintptr
	TryParse uintptr
}

type IHttpLanguageRangeWithQualityHeaderValueStatics struct {
	win32.IInspectable
}

func (this *IHttpLanguageRangeWithQualityHeaderValueStatics) Vtbl() *IHttpLanguageRangeWithQualityHeaderValueStaticsVtbl {
	return (*IHttpLanguageRangeWithQualityHeaderValueStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpLanguageRangeWithQualityHeaderValueStatics) Parse(input string) *IHttpLanguageRangeWithQualityHeaderValue {
	var _result *IHttpLanguageRangeWithQualityHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Parse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpLanguageRangeWithQualityHeaderValueStatics) TryParse(input string, languageRangeWithQualityHeaderValue *IHttpLanguageRangeWithQualityHeaderValue) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(languageRangeWithQualityHeaderValue)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 16B28533-E728-4FCB-BDB0-08A431A14844
var IID_IHttpMediaTypeHeaderValue = syscall.GUID{0x16B28533, 0xE728, 0x4FCB,
	[8]byte{0xBD, 0xB0, 0x08, 0xA4, 0x31, 0xA1, 0x48, 0x44}}

type IHttpMediaTypeHeaderValueInterface interface {
	win32.IInspectableInterface
	Get_CharSet() string
	Put_CharSet(value string)
	Get_MediaType() string
	Put_MediaType(value string)
	Get_Parameters() *IVector[*IHttpNameValueHeaderValue]
}

type IHttpMediaTypeHeaderValueVtbl struct {
	win32.IInspectableVtbl
	Get_CharSet    uintptr
	Put_CharSet    uintptr
	Get_MediaType  uintptr
	Put_MediaType  uintptr
	Get_Parameters uintptr
}

type IHttpMediaTypeHeaderValue struct {
	win32.IInspectable
}

func (this *IHttpMediaTypeHeaderValue) Vtbl() *IHttpMediaTypeHeaderValueVtbl {
	return (*IHttpMediaTypeHeaderValueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpMediaTypeHeaderValue) Get_CharSet() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CharSet, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHttpMediaTypeHeaderValue) Put_CharSet(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CharSet, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IHttpMediaTypeHeaderValue) Get_MediaType() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHttpMediaTypeHeaderValue) Put_MediaType(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MediaType, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IHttpMediaTypeHeaderValue) Get_Parameters() *IVector[*IHttpNameValueHeaderValue] {
	var _result *IVector[*IHttpNameValueHeaderValue]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Parameters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BED747A8-CD17-42DD-9367-AB9C5B56DD7D
var IID_IHttpMediaTypeHeaderValueFactory = syscall.GUID{0xBED747A8, 0xCD17, 0x42DD,
	[8]byte{0x93, 0x67, 0xAB, 0x9C, 0x5B, 0x56, 0xDD, 0x7D}}

type IHttpMediaTypeHeaderValueFactoryInterface interface {
	win32.IInspectableInterface
	Create(mediaType string) *IHttpMediaTypeHeaderValue
}

type IHttpMediaTypeHeaderValueFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IHttpMediaTypeHeaderValueFactory struct {
	win32.IInspectable
}

func (this *IHttpMediaTypeHeaderValueFactory) Vtbl() *IHttpMediaTypeHeaderValueFactoryVtbl {
	return (*IHttpMediaTypeHeaderValueFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpMediaTypeHeaderValueFactory) Create(mediaType string) *IHttpMediaTypeHeaderValue {
	var _result *IHttpMediaTypeHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(mediaType).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E04D83DF-1D41-4D8C-A2DE-6FD2ED87399B
var IID_IHttpMediaTypeHeaderValueStatics = syscall.GUID{0xE04D83DF, 0x1D41, 0x4D8C,
	[8]byte{0xA2, 0xDE, 0x6F, 0xD2, 0xED, 0x87, 0x39, 0x9B}}

type IHttpMediaTypeHeaderValueStaticsInterface interface {
	win32.IInspectableInterface
	Parse(input string) *IHttpMediaTypeHeaderValue
	TryParse(input string, mediaTypeHeaderValue *IHttpMediaTypeHeaderValue) bool
}

type IHttpMediaTypeHeaderValueStaticsVtbl struct {
	win32.IInspectableVtbl
	Parse    uintptr
	TryParse uintptr
}

type IHttpMediaTypeHeaderValueStatics struct {
	win32.IInspectable
}

func (this *IHttpMediaTypeHeaderValueStatics) Vtbl() *IHttpMediaTypeHeaderValueStaticsVtbl {
	return (*IHttpMediaTypeHeaderValueStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpMediaTypeHeaderValueStatics) Parse(input string) *IHttpMediaTypeHeaderValue {
	var _result *IHttpMediaTypeHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Parse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpMediaTypeHeaderValueStatics) TryParse(input string, mediaTypeHeaderValue *IHttpMediaTypeHeaderValue) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(mediaTypeHeaderValue)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 188D5E32-76BE-44A0-B1CD-2074BDED2DDE
var IID_IHttpMediaTypeWithQualityHeaderValue = syscall.GUID{0x188D5E32, 0x76BE, 0x44A0,
	[8]byte{0xB1, 0xCD, 0x20, 0x74, 0xBD, 0xED, 0x2D, 0xDE}}

type IHttpMediaTypeWithQualityHeaderValueInterface interface {
	win32.IInspectableInterface
	Get_CharSet() string
	Put_CharSet(value string)
	Get_MediaType() string
	Put_MediaType(value string)
	Get_Parameters() *IVector[*IHttpNameValueHeaderValue]
	Get_Quality() *IReference[float64]
	Put_Quality(value *IReference[float64])
}

type IHttpMediaTypeWithQualityHeaderValueVtbl struct {
	win32.IInspectableVtbl
	Get_CharSet    uintptr
	Put_CharSet    uintptr
	Get_MediaType  uintptr
	Put_MediaType  uintptr
	Get_Parameters uintptr
	Get_Quality    uintptr
	Put_Quality    uintptr
}

type IHttpMediaTypeWithQualityHeaderValue struct {
	win32.IInspectable
}

func (this *IHttpMediaTypeWithQualityHeaderValue) Vtbl() *IHttpMediaTypeWithQualityHeaderValueVtbl {
	return (*IHttpMediaTypeWithQualityHeaderValueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpMediaTypeWithQualityHeaderValue) Get_CharSet() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CharSet, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHttpMediaTypeWithQualityHeaderValue) Put_CharSet(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CharSet, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IHttpMediaTypeWithQualityHeaderValue) Get_MediaType() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHttpMediaTypeWithQualityHeaderValue) Put_MediaType(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MediaType, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IHttpMediaTypeWithQualityHeaderValue) Get_Parameters() *IVector[*IHttpNameValueHeaderValue] {
	var _result *IVector[*IHttpNameValueHeaderValue]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Parameters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpMediaTypeWithQualityHeaderValue) Get_Quality() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Quality, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpMediaTypeWithQualityHeaderValue) Put_Quality(value *IReference[float64]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Quality, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 3C0C6B73-1342-4587-A056-18D02FF67165
var IID_IHttpMediaTypeWithQualityHeaderValueCollection = syscall.GUID{0x3C0C6B73, 0x1342, 0x4587,
	[8]byte{0xA0, 0x56, 0x18, 0xD0, 0x2F, 0xF6, 0x71, 0x65}}

type IHttpMediaTypeWithQualityHeaderValueCollectionInterface interface {
	win32.IInspectableInterface
	ParseAdd(input string)
	TryParseAdd(input string) bool
}

type IHttpMediaTypeWithQualityHeaderValueCollectionVtbl struct {
	win32.IInspectableVtbl
	ParseAdd    uintptr
	TryParseAdd uintptr
}

type IHttpMediaTypeWithQualityHeaderValueCollection struct {
	win32.IInspectable
}

func (this *IHttpMediaTypeWithQualityHeaderValueCollection) Vtbl() *IHttpMediaTypeWithQualityHeaderValueCollectionVtbl {
	return (*IHttpMediaTypeWithQualityHeaderValueCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpMediaTypeWithQualityHeaderValueCollection) ParseAdd(input string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ParseAdd, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr)
	_ = _hr
}

func (this *IHttpMediaTypeWithQualityHeaderValueCollection) TryParseAdd(input string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParseAdd, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4C6D20F4-9457-44E6-A323-D122B958780B
var IID_IHttpMediaTypeWithQualityHeaderValueFactory = syscall.GUID{0x4C6D20F4, 0x9457, 0x44E6,
	[8]byte{0xA3, 0x23, 0xD1, 0x22, 0xB9, 0x58, 0x78, 0x0B}}

type IHttpMediaTypeWithQualityHeaderValueFactoryInterface interface {
	win32.IInspectableInterface
	CreateFromMediaType(mediaType string) *IHttpMediaTypeWithQualityHeaderValue
	CreateFromMediaTypeWithQuality(mediaType string, quality float64) *IHttpMediaTypeWithQualityHeaderValue
}

type IHttpMediaTypeWithQualityHeaderValueFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateFromMediaType            uintptr
	CreateFromMediaTypeWithQuality uintptr
}

type IHttpMediaTypeWithQualityHeaderValueFactory struct {
	win32.IInspectable
}

func (this *IHttpMediaTypeWithQualityHeaderValueFactory) Vtbl() *IHttpMediaTypeWithQualityHeaderValueFactoryVtbl {
	return (*IHttpMediaTypeWithQualityHeaderValueFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpMediaTypeWithQualityHeaderValueFactory) CreateFromMediaType(mediaType string) *IHttpMediaTypeWithQualityHeaderValue {
	var _result *IHttpMediaTypeWithQualityHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromMediaType, uintptr(unsafe.Pointer(this)), NewHStr(mediaType).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpMediaTypeWithQualityHeaderValueFactory) CreateFromMediaTypeWithQuality(mediaType string, quality float64) *IHttpMediaTypeWithQualityHeaderValue {
	var _result *IHttpMediaTypeWithQualityHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromMediaTypeWithQuality, uintptr(unsafe.Pointer(this)), NewHStr(mediaType).Ptr, uintptr(quality), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5B070CD9-B560-4FC8-9835-7E6C0A657B24
var IID_IHttpMediaTypeWithQualityHeaderValueStatics = syscall.GUID{0x5B070CD9, 0xB560, 0x4FC8,
	[8]byte{0x98, 0x35, 0x7E, 0x6C, 0x0A, 0x65, 0x7B, 0x24}}

type IHttpMediaTypeWithQualityHeaderValueStaticsInterface interface {
	win32.IInspectableInterface
	Parse(input string) *IHttpMediaTypeWithQualityHeaderValue
	TryParse(input string, mediaTypeWithQualityHeaderValue *IHttpMediaTypeWithQualityHeaderValue) bool
}

type IHttpMediaTypeWithQualityHeaderValueStaticsVtbl struct {
	win32.IInspectableVtbl
	Parse    uintptr
	TryParse uintptr
}

type IHttpMediaTypeWithQualityHeaderValueStatics struct {
	win32.IInspectable
}

func (this *IHttpMediaTypeWithQualityHeaderValueStatics) Vtbl() *IHttpMediaTypeWithQualityHeaderValueStaticsVtbl {
	return (*IHttpMediaTypeWithQualityHeaderValueStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpMediaTypeWithQualityHeaderValueStatics) Parse(input string) *IHttpMediaTypeWithQualityHeaderValue {
	var _result *IHttpMediaTypeWithQualityHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Parse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpMediaTypeWithQualityHeaderValueStatics) TryParse(input string, mediaTypeWithQualityHeaderValue *IHttpMediaTypeWithQualityHeaderValue) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(mediaTypeWithQualityHeaderValue)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 43BC3FF4-6119-4ADF-938C-34BFFFCF92ED
var IID_IHttpMethodHeaderValueCollection = syscall.GUID{0x43BC3FF4, 0x6119, 0x4ADF,
	[8]byte{0x93, 0x8C, 0x34, 0xBF, 0xFF, 0xCF, 0x92, 0xED}}

type IHttpMethodHeaderValueCollectionInterface interface {
	win32.IInspectableInterface
	ParseAdd(input string)
	TryParseAdd(input string) bool
}

type IHttpMethodHeaderValueCollectionVtbl struct {
	win32.IInspectableVtbl
	ParseAdd    uintptr
	TryParseAdd uintptr
}

type IHttpMethodHeaderValueCollection struct {
	win32.IInspectable
}

func (this *IHttpMethodHeaderValueCollection) Vtbl() *IHttpMethodHeaderValueCollectionVtbl {
	return (*IHttpMethodHeaderValueCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpMethodHeaderValueCollection) ParseAdd(input string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ParseAdd, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr)
	_ = _hr
}

func (this *IHttpMethodHeaderValueCollection) TryParseAdd(input string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParseAdd, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D8BA7463-5B9A-4D1B-93F9-AA5B44ECFDDF
var IID_IHttpNameValueHeaderValue = syscall.GUID{0xD8BA7463, 0x5B9A, 0x4D1B,
	[8]byte{0x93, 0xF9, 0xAA, 0x5B, 0x44, 0xEC, 0xFD, 0xDF}}

type IHttpNameValueHeaderValueInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Get_Value() string
	Put_Value(value string)
}

type IHttpNameValueHeaderValueVtbl struct {
	win32.IInspectableVtbl
	Get_Name  uintptr
	Get_Value uintptr
	Put_Value uintptr
}

type IHttpNameValueHeaderValue struct {
	win32.IInspectable
}

func (this *IHttpNameValueHeaderValue) Vtbl() *IHttpNameValueHeaderValueVtbl {
	return (*IHttpNameValueHeaderValueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpNameValueHeaderValue) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHttpNameValueHeaderValue) Get_Value() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHttpNameValueHeaderValue) Put_Value(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Value, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// 770E2267-CBF8-4736-A925-93FBE10C7CA8
var IID_IHttpNameValueHeaderValueFactory = syscall.GUID{0x770E2267, 0xCBF8, 0x4736,
	[8]byte{0xA9, 0x25, 0x93, 0xFB, 0xE1, 0x0C, 0x7C, 0xA8}}

type IHttpNameValueHeaderValueFactoryInterface interface {
	win32.IInspectableInterface
	CreateFromName(name string) *IHttpNameValueHeaderValue
	CreateFromNameWithValue(name string, value string) *IHttpNameValueHeaderValue
}

type IHttpNameValueHeaderValueFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateFromName          uintptr
	CreateFromNameWithValue uintptr
}

type IHttpNameValueHeaderValueFactory struct {
	win32.IInspectable
}

func (this *IHttpNameValueHeaderValueFactory) Vtbl() *IHttpNameValueHeaderValueFactoryVtbl {
	return (*IHttpNameValueHeaderValueFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpNameValueHeaderValueFactory) CreateFromName(name string) *IHttpNameValueHeaderValue {
	var _result *IHttpNameValueHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromName, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpNameValueHeaderValueFactory) CreateFromNameWithValue(name string, value string) *IHttpNameValueHeaderValue {
	var _result *IHttpNameValueHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromNameWithValue, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, NewHStr(value).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FFD4030F-1130-4152-8659-256909A9D115
var IID_IHttpNameValueHeaderValueStatics = syscall.GUID{0xFFD4030F, 0x1130, 0x4152,
	[8]byte{0x86, 0x59, 0x25, 0x69, 0x09, 0xA9, 0xD1, 0x15}}

type IHttpNameValueHeaderValueStaticsInterface interface {
	win32.IInspectableInterface
	Parse(input string) *IHttpNameValueHeaderValue
	TryParse(input string, nameValueHeaderValue *IHttpNameValueHeaderValue) bool
}

type IHttpNameValueHeaderValueStaticsVtbl struct {
	win32.IInspectableVtbl
	Parse    uintptr
	TryParse uintptr
}

type IHttpNameValueHeaderValueStatics struct {
	win32.IInspectable
}

func (this *IHttpNameValueHeaderValueStatics) Vtbl() *IHttpNameValueHeaderValueStaticsVtbl {
	return (*IHttpNameValueHeaderValueStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpNameValueHeaderValueStatics) Parse(input string) *IHttpNameValueHeaderValue {
	var _result *IHttpNameValueHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Parse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpNameValueHeaderValueStatics) TryParse(input string, nameValueHeaderValue *IHttpNameValueHeaderValue) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(nameValueHeaderValue)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F4FEEE03-EBD4-4160-B9FF-807C5183B6E6
var IID_IHttpProductHeaderValue = syscall.GUID{0xF4FEEE03, 0xEBD4, 0x4160,
	[8]byte{0xB9, 0xFF, 0x80, 0x7C, 0x51, 0x83, 0xB6, 0xE6}}

type IHttpProductHeaderValueInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Get_Version() string
}

type IHttpProductHeaderValueVtbl struct {
	win32.IInspectableVtbl
	Get_Name    uintptr
	Get_Version uintptr
}

type IHttpProductHeaderValue struct {
	win32.IInspectable
}

func (this *IHttpProductHeaderValue) Vtbl() *IHttpProductHeaderValueVtbl {
	return (*IHttpProductHeaderValueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpProductHeaderValue) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHttpProductHeaderValue) Get_Version() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Version, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 611AA4F5-82BC-42FB-977B-DC00536E5E86
var IID_IHttpProductHeaderValueFactory = syscall.GUID{0x611AA4F5, 0x82BC, 0x42FB,
	[8]byte{0x97, 0x7B, 0xDC, 0x00, 0x53, 0x6E, 0x5E, 0x86}}

type IHttpProductHeaderValueFactoryInterface interface {
	win32.IInspectableInterface
	CreateFromName(productName string) *IHttpProductHeaderValue
	CreateFromNameWithVersion(productName string, productVersion string) *IHttpProductHeaderValue
}

type IHttpProductHeaderValueFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateFromName            uintptr
	CreateFromNameWithVersion uintptr
}

type IHttpProductHeaderValueFactory struct {
	win32.IInspectable
}

func (this *IHttpProductHeaderValueFactory) Vtbl() *IHttpProductHeaderValueFactoryVtbl {
	return (*IHttpProductHeaderValueFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpProductHeaderValueFactory) CreateFromName(productName string) *IHttpProductHeaderValue {
	var _result *IHttpProductHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromName, uintptr(unsafe.Pointer(this)), NewHStr(productName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpProductHeaderValueFactory) CreateFromNameWithVersion(productName string, productVersion string) *IHttpProductHeaderValue {
	var _result *IHttpProductHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromNameWithVersion, uintptr(unsafe.Pointer(this)), NewHStr(productName).Ptr, NewHStr(productVersion).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 90C33E29-BEFC-4337-BE62-49F097975F53
var IID_IHttpProductHeaderValueStatics = syscall.GUID{0x90C33E29, 0xBEFC, 0x4337,
	[8]byte{0xBE, 0x62, 0x49, 0xF0, 0x97, 0x97, 0x5F, 0x53}}

type IHttpProductHeaderValueStaticsInterface interface {
	win32.IInspectableInterface
	Parse(input string) *IHttpProductHeaderValue
	TryParse(input string, productHeaderValue *IHttpProductHeaderValue) bool
}

type IHttpProductHeaderValueStaticsVtbl struct {
	win32.IInspectableVtbl
	Parse    uintptr
	TryParse uintptr
}

type IHttpProductHeaderValueStatics struct {
	win32.IInspectable
}

func (this *IHttpProductHeaderValueStatics) Vtbl() *IHttpProductHeaderValueStaticsVtbl {
	return (*IHttpProductHeaderValueStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpProductHeaderValueStatics) Parse(input string) *IHttpProductHeaderValue {
	var _result *IHttpProductHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Parse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpProductHeaderValueStatics) TryParse(input string, productHeaderValue *IHttpProductHeaderValue) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(productHeaderValue)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 1B1A8732-4C35-486A-966F-646489198E4D
var IID_IHttpProductInfoHeaderValue = syscall.GUID{0x1B1A8732, 0x4C35, 0x486A,
	[8]byte{0x96, 0x6F, 0x64, 0x64, 0x89, 0x19, 0x8E, 0x4D}}

type IHttpProductInfoHeaderValueInterface interface {
	win32.IInspectableInterface
	Get_Product() *IHttpProductHeaderValue
	Get_Comment() string
}

type IHttpProductInfoHeaderValueVtbl struct {
	win32.IInspectableVtbl
	Get_Product uintptr
	Get_Comment uintptr
}

type IHttpProductInfoHeaderValue struct {
	win32.IInspectable
}

func (this *IHttpProductInfoHeaderValue) Vtbl() *IHttpProductInfoHeaderValueVtbl {
	return (*IHttpProductInfoHeaderValueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpProductInfoHeaderValue) Get_Product() *IHttpProductHeaderValue {
	var _result *IHttpProductHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Product, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpProductInfoHeaderValue) Get_Comment() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Comment, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 877DF74A-D69B-44F8-AD4F-453AF9C42ED0
var IID_IHttpProductInfoHeaderValueCollection = syscall.GUID{0x877DF74A, 0xD69B, 0x44F8,
	[8]byte{0xAD, 0x4F, 0x45, 0x3A, 0xF9, 0xC4, 0x2E, 0xD0}}

type IHttpProductInfoHeaderValueCollectionInterface interface {
	win32.IInspectableInterface
	ParseAdd(input string)
	TryParseAdd(input string) bool
}

type IHttpProductInfoHeaderValueCollectionVtbl struct {
	win32.IInspectableVtbl
	ParseAdd    uintptr
	TryParseAdd uintptr
}

type IHttpProductInfoHeaderValueCollection struct {
	win32.IInspectable
}

func (this *IHttpProductInfoHeaderValueCollection) Vtbl() *IHttpProductInfoHeaderValueCollectionVtbl {
	return (*IHttpProductInfoHeaderValueCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpProductInfoHeaderValueCollection) ParseAdd(input string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ParseAdd, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr)
	_ = _hr
}

func (this *IHttpProductInfoHeaderValueCollection) TryParseAdd(input string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParseAdd, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 24220FBE-EABE-4464-B460-EC010B7C41E2
var IID_IHttpProductInfoHeaderValueFactory = syscall.GUID{0x24220FBE, 0xEABE, 0x4464,
	[8]byte{0xB4, 0x60, 0xEC, 0x01, 0x0B, 0x7C, 0x41, 0xE2}}

type IHttpProductInfoHeaderValueFactoryInterface interface {
	win32.IInspectableInterface
	CreateFromComment(productComment string) *IHttpProductInfoHeaderValue
	CreateFromNameWithVersion(productName string, productVersion string) *IHttpProductInfoHeaderValue
}

type IHttpProductInfoHeaderValueFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateFromComment         uintptr
	CreateFromNameWithVersion uintptr
}

type IHttpProductInfoHeaderValueFactory struct {
	win32.IInspectable
}

func (this *IHttpProductInfoHeaderValueFactory) Vtbl() *IHttpProductInfoHeaderValueFactoryVtbl {
	return (*IHttpProductInfoHeaderValueFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpProductInfoHeaderValueFactory) CreateFromComment(productComment string) *IHttpProductInfoHeaderValue {
	var _result *IHttpProductInfoHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromComment, uintptr(unsafe.Pointer(this)), NewHStr(productComment).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpProductInfoHeaderValueFactory) CreateFromNameWithVersion(productName string, productVersion string) *IHttpProductInfoHeaderValue {
	var _result *IHttpProductInfoHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromNameWithVersion, uintptr(unsafe.Pointer(this)), NewHStr(productName).Ptr, NewHStr(productVersion).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DB7FD857-327A-4E73-81E5-7059A302B042
var IID_IHttpProductInfoHeaderValueStatics = syscall.GUID{0xDB7FD857, 0x327A, 0x4E73,
	[8]byte{0x81, 0xE5, 0x70, 0x59, 0xA3, 0x02, 0xB0, 0x42}}

type IHttpProductInfoHeaderValueStaticsInterface interface {
	win32.IInspectableInterface
	Parse(input string) *IHttpProductInfoHeaderValue
	TryParse(input string, productInfoHeaderValue *IHttpProductInfoHeaderValue) bool
}

type IHttpProductInfoHeaderValueStaticsVtbl struct {
	win32.IInspectableVtbl
	Parse    uintptr
	TryParse uintptr
}

type IHttpProductInfoHeaderValueStatics struct {
	win32.IInspectable
}

func (this *IHttpProductInfoHeaderValueStatics) Vtbl() *IHttpProductInfoHeaderValueStaticsVtbl {
	return (*IHttpProductInfoHeaderValueStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpProductInfoHeaderValueStatics) Parse(input string) *IHttpProductInfoHeaderValue {
	var _result *IHttpProductInfoHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Parse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpProductInfoHeaderValueStatics) TryParse(input string, productInfoHeaderValue *IHttpProductInfoHeaderValue) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(productInfoHeaderValue)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// AF40329B-B544-469B-86B9-AC3D466FEA36
var IID_IHttpRequestHeaderCollection = syscall.GUID{0xAF40329B, 0xB544, 0x469B,
	[8]byte{0x86, 0xB9, 0xAC, 0x3D, 0x46, 0x6F, 0xEA, 0x36}}

type IHttpRequestHeaderCollectionInterface interface {
	win32.IInspectableInterface
	Get_Accept() *IHttpMediaTypeWithQualityHeaderValueCollection
	Get_AcceptEncoding() *IHttpContentCodingWithQualityHeaderValueCollection
	Get_AcceptLanguage() *IHttpLanguageRangeWithQualityHeaderValueCollection
	Get_Authorization() *IHttpCredentialsHeaderValue
	Put_Authorization(value *IHttpCredentialsHeaderValue)
	Get_CacheControl() *IHttpCacheDirectiveHeaderValueCollection
	Get_Connection() *IHttpConnectionOptionHeaderValueCollection
	Get_Cookie() *IHttpCookiePairHeaderValueCollection
	Get_Date() *IReference[DateTime]
	Put_Date(value *IReference[DateTime])
	Get_Expect() *IHttpExpectationHeaderValueCollection
	Get_From() string
	Put_From(value string)
	Get_Host() *IHostName
	Put_Host(value *IHostName)
	Get_IfModifiedSince() *IReference[DateTime]
	Put_IfModifiedSince(value *IReference[DateTime])
	Get_IfUnmodifiedSince() *IReference[DateTime]
	Put_IfUnmodifiedSince(value *IReference[DateTime])
	Get_MaxForwards() *IReference[uint32]
	Put_MaxForwards(value *IReference[uint32])
	Get_ProxyAuthorization() *IHttpCredentialsHeaderValue
	Put_ProxyAuthorization(value *IHttpCredentialsHeaderValue)
	Get_Referer() *IUriRuntimeClass
	Put_Referer(value *IUriRuntimeClass)
	Get_TransferEncoding() *IHttpTransferCodingHeaderValueCollection
	Get_UserAgent() *IHttpProductInfoHeaderValueCollection
	Append(name string, value string)
	TryAppendWithoutValidation(name string, value string) bool
}

type IHttpRequestHeaderCollectionVtbl struct {
	win32.IInspectableVtbl
	Get_Accept                 uintptr
	Get_AcceptEncoding         uintptr
	Get_AcceptLanguage         uintptr
	Get_Authorization          uintptr
	Put_Authorization          uintptr
	Get_CacheControl           uintptr
	Get_Connection             uintptr
	Get_Cookie                 uintptr
	Get_Date                   uintptr
	Put_Date                   uintptr
	Get_Expect                 uintptr
	Get_From                   uintptr
	Put_From                   uintptr
	Get_Host                   uintptr
	Put_Host                   uintptr
	Get_IfModifiedSince        uintptr
	Put_IfModifiedSince        uintptr
	Get_IfUnmodifiedSince      uintptr
	Put_IfUnmodifiedSince      uintptr
	Get_MaxForwards            uintptr
	Put_MaxForwards            uintptr
	Get_ProxyAuthorization     uintptr
	Put_ProxyAuthorization     uintptr
	Get_Referer                uintptr
	Put_Referer                uintptr
	Get_TransferEncoding       uintptr
	Get_UserAgent              uintptr
	Append                     uintptr
	TryAppendWithoutValidation uintptr
}

type IHttpRequestHeaderCollection struct {
	win32.IInspectable
}

func (this *IHttpRequestHeaderCollection) Vtbl() *IHttpRequestHeaderCollectionVtbl {
	return (*IHttpRequestHeaderCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpRequestHeaderCollection) Get_Accept() *IHttpMediaTypeWithQualityHeaderValueCollection {
	var _result *IHttpMediaTypeWithQualityHeaderValueCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Accept, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpRequestHeaderCollection) Get_AcceptEncoding() *IHttpContentCodingWithQualityHeaderValueCollection {
	var _result *IHttpContentCodingWithQualityHeaderValueCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AcceptEncoding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpRequestHeaderCollection) Get_AcceptLanguage() *IHttpLanguageRangeWithQualityHeaderValueCollection {
	var _result *IHttpLanguageRangeWithQualityHeaderValueCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AcceptLanguage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpRequestHeaderCollection) Get_Authorization() *IHttpCredentialsHeaderValue {
	var _result *IHttpCredentialsHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Authorization, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpRequestHeaderCollection) Put_Authorization(value *IHttpCredentialsHeaderValue) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Authorization, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpRequestHeaderCollection) Get_CacheControl() *IHttpCacheDirectiveHeaderValueCollection {
	var _result *IHttpCacheDirectiveHeaderValueCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CacheControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpRequestHeaderCollection) Get_Connection() *IHttpConnectionOptionHeaderValueCollection {
	var _result *IHttpConnectionOptionHeaderValueCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Connection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpRequestHeaderCollection) Get_Cookie() *IHttpCookiePairHeaderValueCollection {
	var _result *IHttpCookiePairHeaderValueCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Cookie, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpRequestHeaderCollection) Get_Date() *IReference[DateTime] {
	var _result *IReference[DateTime]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Date, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpRequestHeaderCollection) Put_Date(value *IReference[DateTime]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Date, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpRequestHeaderCollection) Get_Expect() *IHttpExpectationHeaderValueCollection {
	var _result *IHttpExpectationHeaderValueCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Expect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpRequestHeaderCollection) Get_From() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_From, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHttpRequestHeaderCollection) Put_From(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_From, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IHttpRequestHeaderCollection) Get_Host() *IHostName {
	var _result *IHostName
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Host, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpRequestHeaderCollection) Put_Host(value *IHostName) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Host, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpRequestHeaderCollection) Get_IfModifiedSince() *IReference[DateTime] {
	var _result *IReference[DateTime]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IfModifiedSince, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpRequestHeaderCollection) Put_IfModifiedSince(value *IReference[DateTime]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IfModifiedSince, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpRequestHeaderCollection) Get_IfUnmodifiedSince() *IReference[DateTime] {
	var _result *IReference[DateTime]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IfUnmodifiedSince, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpRequestHeaderCollection) Put_IfUnmodifiedSince(value *IReference[DateTime]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IfUnmodifiedSince, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpRequestHeaderCollection) Get_MaxForwards() *IReference[uint32] {
	var _result *IReference[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxForwards, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpRequestHeaderCollection) Put_MaxForwards(value *IReference[uint32]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxForwards, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpRequestHeaderCollection) Get_ProxyAuthorization() *IHttpCredentialsHeaderValue {
	var _result *IHttpCredentialsHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProxyAuthorization, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpRequestHeaderCollection) Put_ProxyAuthorization(value *IHttpCredentialsHeaderValue) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ProxyAuthorization, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpRequestHeaderCollection) Get_Referer() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Referer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpRequestHeaderCollection) Put_Referer(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Referer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpRequestHeaderCollection) Get_TransferEncoding() *IHttpTransferCodingHeaderValueCollection {
	var _result *IHttpTransferCodingHeaderValueCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TransferEncoding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpRequestHeaderCollection) Get_UserAgent() *IHttpProductInfoHeaderValueCollection {
	var _result *IHttpProductInfoHeaderValueCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UserAgent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpRequestHeaderCollection) Append(name string, value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Append, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, NewHStr(value).Ptr)
	_ = _hr
}

func (this *IHttpRequestHeaderCollection) TryAppendWithoutValidation(name string, value string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryAppendWithoutValidation, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, NewHStr(value).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7A990969-FA3F-41ED-AAC6-BF957975C16B
var IID_IHttpResponseHeaderCollection = syscall.GUID{0x7A990969, 0xFA3F, 0x41ED,
	[8]byte{0xAA, 0xC6, 0xBF, 0x95, 0x79, 0x75, 0xC1, 0x6B}}

type IHttpResponseHeaderCollectionInterface interface {
	win32.IInspectableInterface
	Get_Age() *IReference[TimeSpan]
	Put_Age(value *IReference[TimeSpan])
	Get_Allow() *IHttpMethodHeaderValueCollection
	Get_CacheControl() *IHttpCacheDirectiveHeaderValueCollection
	Get_Connection() *IHttpConnectionOptionHeaderValueCollection
	Get_Date() *IReference[DateTime]
	Put_Date(value *IReference[DateTime])
	Get_Location() *IUriRuntimeClass
	Put_Location(value *IUriRuntimeClass)
	Get_ProxyAuthenticate() *IHttpChallengeHeaderValueCollection
	Get_RetryAfter() *IHttpDateOrDeltaHeaderValue
	Put_RetryAfter(value *IHttpDateOrDeltaHeaderValue)
	Get_TransferEncoding() *IHttpTransferCodingHeaderValueCollection
	Get_WwwAuthenticate() *IHttpChallengeHeaderValueCollection
	Append(name string, value string)
	TryAppendWithoutValidation(name string, value string) bool
}

type IHttpResponseHeaderCollectionVtbl struct {
	win32.IInspectableVtbl
	Get_Age                    uintptr
	Put_Age                    uintptr
	Get_Allow                  uintptr
	Get_CacheControl           uintptr
	Get_Connection             uintptr
	Get_Date                   uintptr
	Put_Date                   uintptr
	Get_Location               uintptr
	Put_Location               uintptr
	Get_ProxyAuthenticate      uintptr
	Get_RetryAfter             uintptr
	Put_RetryAfter             uintptr
	Get_TransferEncoding       uintptr
	Get_WwwAuthenticate        uintptr
	Append                     uintptr
	TryAppendWithoutValidation uintptr
}

type IHttpResponseHeaderCollection struct {
	win32.IInspectable
}

func (this *IHttpResponseHeaderCollection) Vtbl() *IHttpResponseHeaderCollectionVtbl {
	return (*IHttpResponseHeaderCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpResponseHeaderCollection) Get_Age() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Age, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpResponseHeaderCollection) Put_Age(value *IReference[TimeSpan]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Age, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpResponseHeaderCollection) Get_Allow() *IHttpMethodHeaderValueCollection {
	var _result *IHttpMethodHeaderValueCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Allow, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpResponseHeaderCollection) Get_CacheControl() *IHttpCacheDirectiveHeaderValueCollection {
	var _result *IHttpCacheDirectiveHeaderValueCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CacheControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpResponseHeaderCollection) Get_Connection() *IHttpConnectionOptionHeaderValueCollection {
	var _result *IHttpConnectionOptionHeaderValueCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Connection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpResponseHeaderCollection) Get_Date() *IReference[DateTime] {
	var _result *IReference[DateTime]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Date, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpResponseHeaderCollection) Put_Date(value *IReference[DateTime]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Date, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpResponseHeaderCollection) Get_Location() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Location, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpResponseHeaderCollection) Put_Location(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Location, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpResponseHeaderCollection) Get_ProxyAuthenticate() *IHttpChallengeHeaderValueCollection {
	var _result *IHttpChallengeHeaderValueCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProxyAuthenticate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpResponseHeaderCollection) Get_RetryAfter() *IHttpDateOrDeltaHeaderValue {
	var _result *IHttpDateOrDeltaHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RetryAfter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpResponseHeaderCollection) Put_RetryAfter(value *IHttpDateOrDeltaHeaderValue) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RetryAfter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IHttpResponseHeaderCollection) Get_TransferEncoding() *IHttpTransferCodingHeaderValueCollection {
	var _result *IHttpTransferCodingHeaderValueCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TransferEncoding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpResponseHeaderCollection) Get_WwwAuthenticate() *IHttpChallengeHeaderValueCollection {
	var _result *IHttpChallengeHeaderValueCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WwwAuthenticate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpResponseHeaderCollection) Append(name string, value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Append, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, NewHStr(value).Ptr)
	_ = _hr
}

func (this *IHttpResponseHeaderCollection) TryAppendWithoutValidation(name string, value string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryAppendWithoutValidation, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, NewHStr(value).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 436F32F9-3DED-42BD-B38A-5496A2511CE6
var IID_IHttpTransferCodingHeaderValue = syscall.GUID{0x436F32F9, 0x3DED, 0x42BD,
	[8]byte{0xB3, 0x8A, 0x54, 0x96, 0xA2, 0x51, 0x1C, 0xE6}}

type IHttpTransferCodingHeaderValueInterface interface {
	win32.IInspectableInterface
	Get_Parameters() *IVector[*IHttpNameValueHeaderValue]
	Get_Value() string
}

type IHttpTransferCodingHeaderValueVtbl struct {
	win32.IInspectableVtbl
	Get_Parameters uintptr
	Get_Value      uintptr
}

type IHttpTransferCodingHeaderValue struct {
	win32.IInspectable
}

func (this *IHttpTransferCodingHeaderValue) Vtbl() *IHttpTransferCodingHeaderValueVtbl {
	return (*IHttpTransferCodingHeaderValueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpTransferCodingHeaderValue) Get_Parameters() *IVector[*IHttpNameValueHeaderValue] {
	var _result *IVector[*IHttpNameValueHeaderValue]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Parameters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpTransferCodingHeaderValue) Get_Value() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 202C8C34-2C03-49B8-9665-73E27CB2FC79
var IID_IHttpTransferCodingHeaderValueCollection = syscall.GUID{0x202C8C34, 0x2C03, 0x49B8,
	[8]byte{0x96, 0x65, 0x73, 0xE2, 0x7C, 0xB2, 0xFC, 0x79}}

type IHttpTransferCodingHeaderValueCollectionInterface interface {
	win32.IInspectableInterface
	ParseAdd(input string)
	TryParseAdd(input string) bool
}

type IHttpTransferCodingHeaderValueCollectionVtbl struct {
	win32.IInspectableVtbl
	ParseAdd    uintptr
	TryParseAdd uintptr
}

type IHttpTransferCodingHeaderValueCollection struct {
	win32.IInspectable
}

func (this *IHttpTransferCodingHeaderValueCollection) Vtbl() *IHttpTransferCodingHeaderValueCollectionVtbl {
	return (*IHttpTransferCodingHeaderValueCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpTransferCodingHeaderValueCollection) ParseAdd(input string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ParseAdd, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr)
	_ = _hr
}

func (this *IHttpTransferCodingHeaderValueCollection) TryParseAdd(input string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParseAdd, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// BB62DFFC-E361-4F08-8E4F-C9E723DE703B
var IID_IHttpTransferCodingHeaderValueFactory = syscall.GUID{0xBB62DFFC, 0xE361, 0x4F08,
	[8]byte{0x8E, 0x4F, 0xC9, 0xE7, 0x23, 0xDE, 0x70, 0x3B}}

type IHttpTransferCodingHeaderValueFactoryInterface interface {
	win32.IInspectableInterface
	Create(input string) *IHttpTransferCodingHeaderValue
}

type IHttpTransferCodingHeaderValueFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IHttpTransferCodingHeaderValueFactory struct {
	win32.IInspectable
}

func (this *IHttpTransferCodingHeaderValueFactory) Vtbl() *IHttpTransferCodingHeaderValueFactoryVtbl {
	return (*IHttpTransferCodingHeaderValueFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpTransferCodingHeaderValueFactory) Create(input string) *IHttpTransferCodingHeaderValue {
	var _result *IHttpTransferCodingHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6AB8892A-1A98-4D32-A906-7470A9875CE5
var IID_IHttpTransferCodingHeaderValueStatics = syscall.GUID{0x6AB8892A, 0x1A98, 0x4D32,
	[8]byte{0xA9, 0x06, 0x74, 0x70, 0xA9, 0x87, 0x5C, 0xE5}}

type IHttpTransferCodingHeaderValueStaticsInterface interface {
	win32.IInspectableInterface
	Parse(input string) *IHttpTransferCodingHeaderValue
	TryParse(input string, transferCodingHeaderValue *IHttpTransferCodingHeaderValue) bool
}

type IHttpTransferCodingHeaderValueStaticsVtbl struct {
	win32.IInspectableVtbl
	Parse    uintptr
	TryParse uintptr
}

type IHttpTransferCodingHeaderValueStatics struct {
	win32.IInspectable
}

func (this *IHttpTransferCodingHeaderValueStatics) Vtbl() *IHttpTransferCodingHeaderValueStaticsVtbl {
	return (*IHttpTransferCodingHeaderValueStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHttpTransferCodingHeaderValueStatics) Parse(input string) *IHttpTransferCodingHeaderValue {
	var _result *IHttpTransferCodingHeaderValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Parse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHttpTransferCodingHeaderValueStatics) TryParse(input string, transferCodingHeaderValue *IHttpTransferCodingHeaderValue) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryParse, uintptr(unsafe.Pointer(this)), NewHStr(input).Ptr, uintptr(unsafe.Pointer(transferCodingHeaderValue)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type HttpCacheDirectiveHeaderValueCollection struct {
	RtClass
	*IHttpCacheDirectiveHeaderValueCollection
}

type HttpChallengeHeaderValue struct {
	RtClass
	*IHttpChallengeHeaderValue
}

func NewHttpChallengeHeaderValue_CreateFromScheme(scheme string) *HttpChallengeHeaderValue {
	hs := NewHStr("Windows.Web.Http.Headers.HttpChallengeHeaderValue")
	var pFac *IHttpChallengeHeaderValueFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpChallengeHeaderValueFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpChallengeHeaderValue
	p = pFac.CreateFromScheme(scheme)
	result := &HttpChallengeHeaderValue{
		RtClass:                   RtClass{PInspect: &p.IInspectable},
		IHttpChallengeHeaderValue: p,
	}
	com.AddToScope(result)
	return result
}

func NewHttpChallengeHeaderValue_CreateFromSchemeWithToken(scheme string, token string) *HttpChallengeHeaderValue {
	hs := NewHStr("Windows.Web.Http.Headers.HttpChallengeHeaderValue")
	var pFac *IHttpChallengeHeaderValueFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpChallengeHeaderValueFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpChallengeHeaderValue
	p = pFac.CreateFromSchemeWithToken(scheme, token)
	result := &HttpChallengeHeaderValue{
		RtClass:                   RtClass{PInspect: &p.IInspectable},
		IHttpChallengeHeaderValue: p,
	}
	com.AddToScope(result)
	return result
}

func NewIHttpChallengeHeaderValueStatics() *IHttpChallengeHeaderValueStatics {
	var p *IHttpChallengeHeaderValueStatics
	hs := NewHStr("Windows.Web.Http.Headers.HttpChallengeHeaderValue")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpChallengeHeaderValueStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type HttpChallengeHeaderValueCollection struct {
	RtClass
	*IHttpChallengeHeaderValueCollection
}

type HttpConnectionOptionHeaderValue struct {
	RtClass
	*IHttpConnectionOptionHeaderValue
}

func NewHttpConnectionOptionHeaderValue_Create(token string) *HttpConnectionOptionHeaderValue {
	hs := NewHStr("Windows.Web.Http.Headers.HttpConnectionOptionHeaderValue")
	var pFac *IHttpConnectionOptionHeaderValueFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpConnectionOptionHeaderValueFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpConnectionOptionHeaderValue
	p = pFac.Create(token)
	result := &HttpConnectionOptionHeaderValue{
		RtClass:                          RtClass{PInspect: &p.IInspectable},
		IHttpConnectionOptionHeaderValue: p,
	}
	com.AddToScope(result)
	return result
}

func NewIHttpConnectionOptionHeaderValueStatics() *IHttpConnectionOptionHeaderValueStatics {
	var p *IHttpConnectionOptionHeaderValueStatics
	hs := NewHStr("Windows.Web.Http.Headers.HttpConnectionOptionHeaderValue")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpConnectionOptionHeaderValueStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type HttpConnectionOptionHeaderValueCollection struct {
	RtClass
	*IHttpConnectionOptionHeaderValueCollection
}

type HttpContentCodingHeaderValue struct {
	RtClass
	*IHttpContentCodingHeaderValue
}

func NewHttpContentCodingHeaderValue_Create(contentCoding string) *HttpContentCodingHeaderValue {
	hs := NewHStr("Windows.Web.Http.Headers.HttpContentCodingHeaderValue")
	var pFac *IHttpContentCodingHeaderValueFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpContentCodingHeaderValueFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpContentCodingHeaderValue
	p = pFac.Create(contentCoding)
	result := &HttpContentCodingHeaderValue{
		RtClass:                       RtClass{PInspect: &p.IInspectable},
		IHttpContentCodingHeaderValue: p,
	}
	com.AddToScope(result)
	return result
}

func NewIHttpContentCodingHeaderValueStatics() *IHttpContentCodingHeaderValueStatics {
	var p *IHttpContentCodingHeaderValueStatics
	hs := NewHStr("Windows.Web.Http.Headers.HttpContentCodingHeaderValue")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpContentCodingHeaderValueStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type HttpContentCodingHeaderValueCollection struct {
	RtClass
	*IHttpContentCodingHeaderValueCollection
}

type HttpContentCodingWithQualityHeaderValue struct {
	RtClass
	*IHttpContentCodingWithQualityHeaderValue
}

func NewHttpContentCodingWithQualityHeaderValue_CreateFromValue(contentCoding string) *HttpContentCodingWithQualityHeaderValue {
	hs := NewHStr("Windows.Web.Http.Headers.HttpContentCodingWithQualityHeaderValue")
	var pFac *IHttpContentCodingWithQualityHeaderValueFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpContentCodingWithQualityHeaderValueFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpContentCodingWithQualityHeaderValue
	p = pFac.CreateFromValue(contentCoding)
	result := &HttpContentCodingWithQualityHeaderValue{
		RtClass:                                  RtClass{PInspect: &p.IInspectable},
		IHttpContentCodingWithQualityHeaderValue: p,
	}
	com.AddToScope(result)
	return result
}

func NewHttpContentCodingWithQualityHeaderValue_CreateFromValueWithQuality(contentCoding string, quality float64) *HttpContentCodingWithQualityHeaderValue {
	hs := NewHStr("Windows.Web.Http.Headers.HttpContentCodingWithQualityHeaderValue")
	var pFac *IHttpContentCodingWithQualityHeaderValueFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpContentCodingWithQualityHeaderValueFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpContentCodingWithQualityHeaderValue
	p = pFac.CreateFromValueWithQuality(contentCoding, quality)
	result := &HttpContentCodingWithQualityHeaderValue{
		RtClass:                                  RtClass{PInspect: &p.IInspectable},
		IHttpContentCodingWithQualityHeaderValue: p,
	}
	com.AddToScope(result)
	return result
}

func NewIHttpContentCodingWithQualityHeaderValueStatics() *IHttpContentCodingWithQualityHeaderValueStatics {
	var p *IHttpContentCodingWithQualityHeaderValueStatics
	hs := NewHStr("Windows.Web.Http.Headers.HttpContentCodingWithQualityHeaderValue")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpContentCodingWithQualityHeaderValueStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type HttpContentCodingWithQualityHeaderValueCollection struct {
	RtClass
	*IHttpContentCodingWithQualityHeaderValueCollection
}

type HttpContentDispositionHeaderValue struct {
	RtClass
	*IHttpContentDispositionHeaderValue
}

func NewHttpContentDispositionHeaderValue_Create(dispositionType string) *HttpContentDispositionHeaderValue {
	hs := NewHStr("Windows.Web.Http.Headers.HttpContentDispositionHeaderValue")
	var pFac *IHttpContentDispositionHeaderValueFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpContentDispositionHeaderValueFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpContentDispositionHeaderValue
	p = pFac.Create(dispositionType)
	result := &HttpContentDispositionHeaderValue{
		RtClass:                            RtClass{PInspect: &p.IInspectable},
		IHttpContentDispositionHeaderValue: p,
	}
	com.AddToScope(result)
	return result
}

func NewIHttpContentDispositionHeaderValueStatics() *IHttpContentDispositionHeaderValueStatics {
	var p *IHttpContentDispositionHeaderValueStatics
	hs := NewHStr("Windows.Web.Http.Headers.HttpContentDispositionHeaderValue")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpContentDispositionHeaderValueStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type HttpContentHeaderCollection struct {
	RtClass
	*IHttpContentHeaderCollection
}

func NewHttpContentHeaderCollection() *HttpContentHeaderCollection {
	hs := NewHStr("Windows.Web.Http.Headers.HttpContentHeaderCollection")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &HttpContentHeaderCollection{
		RtClass:                      RtClass{PInspect: p},
		IHttpContentHeaderCollection: (*IHttpContentHeaderCollection)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type HttpContentRangeHeaderValue struct {
	RtClass
	*IHttpContentRangeHeaderValue
}

func NewHttpContentRangeHeaderValue_CreateFromLength(length uint64) *HttpContentRangeHeaderValue {
	hs := NewHStr("Windows.Web.Http.Headers.HttpContentRangeHeaderValue")
	var pFac *IHttpContentRangeHeaderValueFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpContentRangeHeaderValueFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpContentRangeHeaderValue
	p = pFac.CreateFromLength(length)
	result := &HttpContentRangeHeaderValue{
		RtClass:                      RtClass{PInspect: &p.IInspectable},
		IHttpContentRangeHeaderValue: p,
	}
	com.AddToScope(result)
	return result
}

func NewHttpContentRangeHeaderValue_CreateFromRange(from uint64, to uint64) *HttpContentRangeHeaderValue {
	hs := NewHStr("Windows.Web.Http.Headers.HttpContentRangeHeaderValue")
	var pFac *IHttpContentRangeHeaderValueFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpContentRangeHeaderValueFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpContentRangeHeaderValue
	p = pFac.CreateFromRange(from, to)
	result := &HttpContentRangeHeaderValue{
		RtClass:                      RtClass{PInspect: &p.IInspectable},
		IHttpContentRangeHeaderValue: p,
	}
	com.AddToScope(result)
	return result
}

func NewHttpContentRangeHeaderValue_CreateFromRangeWithLength(from uint64, to uint64, length uint64) *HttpContentRangeHeaderValue {
	hs := NewHStr("Windows.Web.Http.Headers.HttpContentRangeHeaderValue")
	var pFac *IHttpContentRangeHeaderValueFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpContentRangeHeaderValueFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpContentRangeHeaderValue
	p = pFac.CreateFromRangeWithLength(from, to, length)
	result := &HttpContentRangeHeaderValue{
		RtClass:                      RtClass{PInspect: &p.IInspectable},
		IHttpContentRangeHeaderValue: p,
	}
	com.AddToScope(result)
	return result
}

func NewIHttpContentRangeHeaderValueStatics() *IHttpContentRangeHeaderValueStatics {
	var p *IHttpContentRangeHeaderValueStatics
	hs := NewHStr("Windows.Web.Http.Headers.HttpContentRangeHeaderValue")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpContentRangeHeaderValueStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type HttpCookiePairHeaderValue struct {
	RtClass
	*IHttpCookiePairHeaderValue
}

func NewHttpCookiePairHeaderValue_CreateFromName(name string) *HttpCookiePairHeaderValue {
	hs := NewHStr("Windows.Web.Http.Headers.HttpCookiePairHeaderValue")
	var pFac *IHttpCookiePairHeaderValueFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpCookiePairHeaderValueFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpCookiePairHeaderValue
	p = pFac.CreateFromName(name)
	result := &HttpCookiePairHeaderValue{
		RtClass:                    RtClass{PInspect: &p.IInspectable},
		IHttpCookiePairHeaderValue: p,
	}
	com.AddToScope(result)
	return result
}

func NewHttpCookiePairHeaderValue_CreateFromNameWithValue(name string, value string) *HttpCookiePairHeaderValue {
	hs := NewHStr("Windows.Web.Http.Headers.HttpCookiePairHeaderValue")
	var pFac *IHttpCookiePairHeaderValueFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpCookiePairHeaderValueFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpCookiePairHeaderValue
	p = pFac.CreateFromNameWithValue(name, value)
	result := &HttpCookiePairHeaderValue{
		RtClass:                    RtClass{PInspect: &p.IInspectable},
		IHttpCookiePairHeaderValue: p,
	}
	com.AddToScope(result)
	return result
}

func NewIHttpCookiePairHeaderValueStatics() *IHttpCookiePairHeaderValueStatics {
	var p *IHttpCookiePairHeaderValueStatics
	hs := NewHStr("Windows.Web.Http.Headers.HttpCookiePairHeaderValue")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpCookiePairHeaderValueStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type HttpCookiePairHeaderValueCollection struct {
	RtClass
	*IHttpCookiePairHeaderValueCollection
}

type HttpCredentialsHeaderValue struct {
	RtClass
	*IHttpCredentialsHeaderValue
}

func NewHttpCredentialsHeaderValue_CreateFromScheme(scheme string) *HttpCredentialsHeaderValue {
	hs := NewHStr("Windows.Web.Http.Headers.HttpCredentialsHeaderValue")
	var pFac *IHttpCredentialsHeaderValueFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpCredentialsHeaderValueFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpCredentialsHeaderValue
	p = pFac.CreateFromScheme(scheme)
	result := &HttpCredentialsHeaderValue{
		RtClass:                     RtClass{PInspect: &p.IInspectable},
		IHttpCredentialsHeaderValue: p,
	}
	com.AddToScope(result)
	return result
}

func NewHttpCredentialsHeaderValue_CreateFromSchemeWithToken(scheme string, token string) *HttpCredentialsHeaderValue {
	hs := NewHStr("Windows.Web.Http.Headers.HttpCredentialsHeaderValue")
	var pFac *IHttpCredentialsHeaderValueFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpCredentialsHeaderValueFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpCredentialsHeaderValue
	p = pFac.CreateFromSchemeWithToken(scheme, token)
	result := &HttpCredentialsHeaderValue{
		RtClass:                     RtClass{PInspect: &p.IInspectable},
		IHttpCredentialsHeaderValue: p,
	}
	com.AddToScope(result)
	return result
}

func NewIHttpCredentialsHeaderValueStatics() *IHttpCredentialsHeaderValueStatics {
	var p *IHttpCredentialsHeaderValueStatics
	hs := NewHStr("Windows.Web.Http.Headers.HttpCredentialsHeaderValue")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpCredentialsHeaderValueStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type HttpDateOrDeltaHeaderValue struct {
	RtClass
	*IHttpDateOrDeltaHeaderValue
}

func NewIHttpDateOrDeltaHeaderValueStatics() *IHttpDateOrDeltaHeaderValueStatics {
	var p *IHttpDateOrDeltaHeaderValueStatics
	hs := NewHStr("Windows.Web.Http.Headers.HttpDateOrDeltaHeaderValue")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpDateOrDeltaHeaderValueStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type HttpExpectationHeaderValue struct {
	RtClass
	*IHttpExpectationHeaderValue
}

func NewHttpExpectationHeaderValue_CreateFromName(name string) *HttpExpectationHeaderValue {
	hs := NewHStr("Windows.Web.Http.Headers.HttpExpectationHeaderValue")
	var pFac *IHttpExpectationHeaderValueFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpExpectationHeaderValueFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpExpectationHeaderValue
	p = pFac.CreateFromName(name)
	result := &HttpExpectationHeaderValue{
		RtClass:                     RtClass{PInspect: &p.IInspectable},
		IHttpExpectationHeaderValue: p,
	}
	com.AddToScope(result)
	return result
}

func NewHttpExpectationHeaderValue_CreateFromNameWithValue(name string, value string) *HttpExpectationHeaderValue {
	hs := NewHStr("Windows.Web.Http.Headers.HttpExpectationHeaderValue")
	var pFac *IHttpExpectationHeaderValueFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpExpectationHeaderValueFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpExpectationHeaderValue
	p = pFac.CreateFromNameWithValue(name, value)
	result := &HttpExpectationHeaderValue{
		RtClass:                     RtClass{PInspect: &p.IInspectable},
		IHttpExpectationHeaderValue: p,
	}
	com.AddToScope(result)
	return result
}

func NewIHttpExpectationHeaderValueStatics() *IHttpExpectationHeaderValueStatics {
	var p *IHttpExpectationHeaderValueStatics
	hs := NewHStr("Windows.Web.Http.Headers.HttpExpectationHeaderValue")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpExpectationHeaderValueStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type HttpExpectationHeaderValueCollection struct {
	RtClass
	*IHttpExpectationHeaderValueCollection
}

type HttpLanguageHeaderValueCollection struct {
	RtClass
	*IHttpLanguageHeaderValueCollection
}

type HttpLanguageRangeWithQualityHeaderValue struct {
	RtClass
	*IHttpLanguageRangeWithQualityHeaderValue
}

func NewHttpLanguageRangeWithQualityHeaderValue_CreateFromLanguageRange(languageRange string) *HttpLanguageRangeWithQualityHeaderValue {
	hs := NewHStr("Windows.Web.Http.Headers.HttpLanguageRangeWithQualityHeaderValue")
	var pFac *IHttpLanguageRangeWithQualityHeaderValueFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpLanguageRangeWithQualityHeaderValueFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpLanguageRangeWithQualityHeaderValue
	p = pFac.CreateFromLanguageRange(languageRange)
	result := &HttpLanguageRangeWithQualityHeaderValue{
		RtClass:                                  RtClass{PInspect: &p.IInspectable},
		IHttpLanguageRangeWithQualityHeaderValue: p,
	}
	com.AddToScope(result)
	return result
}

func NewHttpLanguageRangeWithQualityHeaderValue_CreateFromLanguageRangeWithQuality(languageRange string, quality float64) *HttpLanguageRangeWithQualityHeaderValue {
	hs := NewHStr("Windows.Web.Http.Headers.HttpLanguageRangeWithQualityHeaderValue")
	var pFac *IHttpLanguageRangeWithQualityHeaderValueFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpLanguageRangeWithQualityHeaderValueFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpLanguageRangeWithQualityHeaderValue
	p = pFac.CreateFromLanguageRangeWithQuality(languageRange, quality)
	result := &HttpLanguageRangeWithQualityHeaderValue{
		RtClass:                                  RtClass{PInspect: &p.IInspectable},
		IHttpLanguageRangeWithQualityHeaderValue: p,
	}
	com.AddToScope(result)
	return result
}

func NewIHttpLanguageRangeWithQualityHeaderValueStatics() *IHttpLanguageRangeWithQualityHeaderValueStatics {
	var p *IHttpLanguageRangeWithQualityHeaderValueStatics
	hs := NewHStr("Windows.Web.Http.Headers.HttpLanguageRangeWithQualityHeaderValue")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpLanguageRangeWithQualityHeaderValueStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type HttpLanguageRangeWithQualityHeaderValueCollection struct {
	RtClass
	*IHttpLanguageRangeWithQualityHeaderValueCollection
}

type HttpMediaTypeHeaderValue struct {
	RtClass
	*IHttpMediaTypeHeaderValue
}

func NewHttpMediaTypeHeaderValue_Create(mediaType string) *HttpMediaTypeHeaderValue {
	hs := NewHStr("Windows.Web.Http.Headers.HttpMediaTypeHeaderValue")
	var pFac *IHttpMediaTypeHeaderValueFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpMediaTypeHeaderValueFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpMediaTypeHeaderValue
	p = pFac.Create(mediaType)
	result := &HttpMediaTypeHeaderValue{
		RtClass:                   RtClass{PInspect: &p.IInspectable},
		IHttpMediaTypeHeaderValue: p,
	}
	com.AddToScope(result)
	return result
}

func NewIHttpMediaTypeHeaderValueStatics() *IHttpMediaTypeHeaderValueStatics {
	var p *IHttpMediaTypeHeaderValueStatics
	hs := NewHStr("Windows.Web.Http.Headers.HttpMediaTypeHeaderValue")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpMediaTypeHeaderValueStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type HttpMediaTypeWithQualityHeaderValue struct {
	RtClass
	*IHttpMediaTypeWithQualityHeaderValue
}

func NewHttpMediaTypeWithQualityHeaderValue_CreateFromMediaType(mediaType string) *HttpMediaTypeWithQualityHeaderValue {
	hs := NewHStr("Windows.Web.Http.Headers.HttpMediaTypeWithQualityHeaderValue")
	var pFac *IHttpMediaTypeWithQualityHeaderValueFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpMediaTypeWithQualityHeaderValueFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpMediaTypeWithQualityHeaderValue
	p = pFac.CreateFromMediaType(mediaType)
	result := &HttpMediaTypeWithQualityHeaderValue{
		RtClass:                              RtClass{PInspect: &p.IInspectable},
		IHttpMediaTypeWithQualityHeaderValue: p,
	}
	com.AddToScope(result)
	return result
}

func NewHttpMediaTypeWithQualityHeaderValue_CreateFromMediaTypeWithQuality(mediaType string, quality float64) *HttpMediaTypeWithQualityHeaderValue {
	hs := NewHStr("Windows.Web.Http.Headers.HttpMediaTypeWithQualityHeaderValue")
	var pFac *IHttpMediaTypeWithQualityHeaderValueFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpMediaTypeWithQualityHeaderValueFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpMediaTypeWithQualityHeaderValue
	p = pFac.CreateFromMediaTypeWithQuality(mediaType, quality)
	result := &HttpMediaTypeWithQualityHeaderValue{
		RtClass:                              RtClass{PInspect: &p.IInspectable},
		IHttpMediaTypeWithQualityHeaderValue: p,
	}
	com.AddToScope(result)
	return result
}

func NewIHttpMediaTypeWithQualityHeaderValueStatics() *IHttpMediaTypeWithQualityHeaderValueStatics {
	var p *IHttpMediaTypeWithQualityHeaderValueStatics
	hs := NewHStr("Windows.Web.Http.Headers.HttpMediaTypeWithQualityHeaderValue")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpMediaTypeWithQualityHeaderValueStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type HttpMediaTypeWithQualityHeaderValueCollection struct {
	RtClass
	*IHttpMediaTypeWithQualityHeaderValueCollection
}

type HttpMethodHeaderValueCollection struct {
	RtClass
	*IHttpMethodHeaderValueCollection
}

type HttpNameValueHeaderValue struct {
	RtClass
	*IHttpNameValueHeaderValue
}

func NewHttpNameValueHeaderValue_CreateFromName(name string) *HttpNameValueHeaderValue {
	hs := NewHStr("Windows.Web.Http.Headers.HttpNameValueHeaderValue")
	var pFac *IHttpNameValueHeaderValueFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpNameValueHeaderValueFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpNameValueHeaderValue
	p = pFac.CreateFromName(name)
	result := &HttpNameValueHeaderValue{
		RtClass:                   RtClass{PInspect: &p.IInspectable},
		IHttpNameValueHeaderValue: p,
	}
	com.AddToScope(result)
	return result
}

func NewHttpNameValueHeaderValue_CreateFromNameWithValue(name string, value string) *HttpNameValueHeaderValue {
	hs := NewHStr("Windows.Web.Http.Headers.HttpNameValueHeaderValue")
	var pFac *IHttpNameValueHeaderValueFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpNameValueHeaderValueFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpNameValueHeaderValue
	p = pFac.CreateFromNameWithValue(name, value)
	result := &HttpNameValueHeaderValue{
		RtClass:                   RtClass{PInspect: &p.IInspectable},
		IHttpNameValueHeaderValue: p,
	}
	com.AddToScope(result)
	return result
}

func NewIHttpNameValueHeaderValueStatics() *IHttpNameValueHeaderValueStatics {
	var p *IHttpNameValueHeaderValueStatics
	hs := NewHStr("Windows.Web.Http.Headers.HttpNameValueHeaderValue")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpNameValueHeaderValueStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type HttpProductHeaderValue struct {
	RtClass
	*IHttpProductHeaderValue
}

func NewHttpProductHeaderValue_CreateFromName(productName string) *HttpProductHeaderValue {
	hs := NewHStr("Windows.Web.Http.Headers.HttpProductHeaderValue")
	var pFac *IHttpProductHeaderValueFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpProductHeaderValueFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpProductHeaderValue
	p = pFac.CreateFromName(productName)
	result := &HttpProductHeaderValue{
		RtClass:                 RtClass{PInspect: &p.IInspectable},
		IHttpProductHeaderValue: p,
	}
	com.AddToScope(result)
	return result
}

func NewHttpProductHeaderValue_CreateFromNameWithVersion(productName string, productVersion string) *HttpProductHeaderValue {
	hs := NewHStr("Windows.Web.Http.Headers.HttpProductHeaderValue")
	var pFac *IHttpProductHeaderValueFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpProductHeaderValueFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpProductHeaderValue
	p = pFac.CreateFromNameWithVersion(productName, productVersion)
	result := &HttpProductHeaderValue{
		RtClass:                 RtClass{PInspect: &p.IInspectable},
		IHttpProductHeaderValue: p,
	}
	com.AddToScope(result)
	return result
}

func NewIHttpProductHeaderValueStatics() *IHttpProductHeaderValueStatics {
	var p *IHttpProductHeaderValueStatics
	hs := NewHStr("Windows.Web.Http.Headers.HttpProductHeaderValue")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpProductHeaderValueStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type HttpProductInfoHeaderValue struct {
	RtClass
	*IHttpProductInfoHeaderValue
}

func NewHttpProductInfoHeaderValue_CreateFromComment(productComment string) *HttpProductInfoHeaderValue {
	hs := NewHStr("Windows.Web.Http.Headers.HttpProductInfoHeaderValue")
	var pFac *IHttpProductInfoHeaderValueFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpProductInfoHeaderValueFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpProductInfoHeaderValue
	p = pFac.CreateFromComment(productComment)
	result := &HttpProductInfoHeaderValue{
		RtClass:                     RtClass{PInspect: &p.IInspectable},
		IHttpProductInfoHeaderValue: p,
	}
	com.AddToScope(result)
	return result
}

func NewHttpProductInfoHeaderValue_CreateFromNameWithVersion(productName string, productVersion string) *HttpProductInfoHeaderValue {
	hs := NewHStr("Windows.Web.Http.Headers.HttpProductInfoHeaderValue")
	var pFac *IHttpProductInfoHeaderValueFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpProductInfoHeaderValueFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpProductInfoHeaderValue
	p = pFac.CreateFromNameWithVersion(productName, productVersion)
	result := &HttpProductInfoHeaderValue{
		RtClass:                     RtClass{PInspect: &p.IInspectable},
		IHttpProductInfoHeaderValue: p,
	}
	com.AddToScope(result)
	return result
}

func NewIHttpProductInfoHeaderValueStatics() *IHttpProductInfoHeaderValueStatics {
	var p *IHttpProductInfoHeaderValueStatics
	hs := NewHStr("Windows.Web.Http.Headers.HttpProductInfoHeaderValue")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpProductInfoHeaderValueStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type HttpProductInfoHeaderValueCollection struct {
	RtClass
	*IHttpProductInfoHeaderValueCollection
}

type HttpRequestHeaderCollection struct {
	RtClass
	*IHttpRequestHeaderCollection
}

type HttpResponseHeaderCollection struct {
	RtClass
	*IHttpResponseHeaderCollection
}

type HttpTransferCodingHeaderValue struct {
	RtClass
	*IHttpTransferCodingHeaderValue
}

func NewHttpTransferCodingHeaderValue_Create(input string) *HttpTransferCodingHeaderValue {
	hs := NewHStr("Windows.Web.Http.Headers.HttpTransferCodingHeaderValue")
	var pFac *IHttpTransferCodingHeaderValueFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpTransferCodingHeaderValueFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IHttpTransferCodingHeaderValue
	p = pFac.Create(input)
	result := &HttpTransferCodingHeaderValue{
		RtClass:                        RtClass{PInspect: &p.IInspectable},
		IHttpTransferCodingHeaderValue: p,
	}
	com.AddToScope(result)
	return result
}

func NewIHttpTransferCodingHeaderValueStatics() *IHttpTransferCodingHeaderValueStatics {
	var p *IHttpTransferCodingHeaderValueStatics
	hs := NewHStr("Windows.Web.Http.Headers.HttpTransferCodingHeaderValue")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHttpTransferCodingHeaderValueStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type HttpTransferCodingHeaderValueCollection struct {
	RtClass
	*IHttpTransferCodingHeaderValueCollection
}
