package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"log"
	"syscall"
	"unsafe"
)

// enums

// enum
type SyndicationErrorStatus int32

const (
	SyndicationErrorStatus_Unknown                  SyndicationErrorStatus = 0
	SyndicationErrorStatus_MissingRequiredElement   SyndicationErrorStatus = 1
	SyndicationErrorStatus_MissingRequiredAttribute SyndicationErrorStatus = 2
	SyndicationErrorStatus_InvalidXml               SyndicationErrorStatus = 3
	SyndicationErrorStatus_UnexpectedContent        SyndicationErrorStatus = 4
	SyndicationErrorStatus_UnsupportedFormat        SyndicationErrorStatus = 5
)

// enum
type SyndicationFormat int32

const (
	SyndicationFormat_Atom10 SyndicationFormat = 0
	SyndicationFormat_Rss20  SyndicationFormat = 1
	SyndicationFormat_Rss10  SyndicationFormat = 2
	SyndicationFormat_Rss092 SyndicationFormat = 3
	SyndicationFormat_Rss091 SyndicationFormat = 4
	SyndicationFormat_Atom03 SyndicationFormat = 5
)

// enum
type SyndicationTextType int32

const (
	SyndicationTextType_Text  SyndicationTextType = 0
	SyndicationTextType_Html  SyndicationTextType = 1
	SyndicationTextType_Xhtml SyndicationTextType = 2
)

// structs

type RetrievalProgress struct {
	BytesRetrieved       uint32
	TotalBytesToRetrieve uint32
}

type TransferProgress struct {
	BytesSent            uint32
	TotalBytesToSend     uint32
	BytesRetrieved       uint32
	TotalBytesToRetrieve uint32
}

// interfaces

// 71E8F969-526E-4001-9A91-E84F83161AB1
var IID_ISyndicationAttribute = syscall.GUID{0x71E8F969, 0x526E, 0x4001,
	[8]byte{0x9A, 0x91, 0xE8, 0x4F, 0x83, 0x16, 0x1A, 0xB1}}

type ISyndicationAttributeInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Put_Name(value string)
	Get_Namespace() string
	Put_Namespace(value string)
	Get_Value() string
	Put_Value(value string)
}

type ISyndicationAttributeVtbl struct {
	win32.IInspectableVtbl
	Get_Name      uintptr
	Put_Name      uintptr
	Get_Namespace uintptr
	Put_Namespace uintptr
	Get_Value     uintptr
	Put_Value     uintptr
}

type ISyndicationAttribute struct {
	win32.IInspectable
}

func (this *ISyndicationAttribute) Vtbl() *ISyndicationAttributeVtbl {
	return (*ISyndicationAttributeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISyndicationAttribute) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISyndicationAttribute) Put_Name(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Name, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISyndicationAttribute) Get_Namespace() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Namespace, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISyndicationAttribute) Put_Namespace(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Namespace, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISyndicationAttribute) Get_Value() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISyndicationAttribute) Put_Value(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Value, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// 624F1599-ED3E-420F-BE86-640414886E4B
var IID_ISyndicationAttributeFactory = syscall.GUID{0x624F1599, 0xED3E, 0x420F,
	[8]byte{0xBE, 0x86, 0x64, 0x04, 0x14, 0x88, 0x6E, 0x4B}}

type ISyndicationAttributeFactoryInterface interface {
	win32.IInspectableInterface
	CreateSyndicationAttribute(attributeName string, attributeNamespace string, attributeValue string) *ISyndicationAttribute
}

type ISyndicationAttributeFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateSyndicationAttribute uintptr
}

type ISyndicationAttributeFactory struct {
	win32.IInspectable
}

func (this *ISyndicationAttributeFactory) Vtbl() *ISyndicationAttributeFactoryVtbl {
	return (*ISyndicationAttributeFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISyndicationAttributeFactory) CreateSyndicationAttribute(attributeName string, attributeNamespace string, attributeValue string) *ISyndicationAttribute {
	var _result *ISyndicationAttribute
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSyndicationAttribute, uintptr(unsafe.Pointer(this)), NewHStr(attributeName).Ptr, NewHStr(attributeNamespace).Ptr, NewHStr(attributeValue).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8715626F-0CBA-4A7F-89FF-ECB5281423B6
var IID_ISyndicationCategory = syscall.GUID{0x8715626F, 0x0CBA, 0x4A7F,
	[8]byte{0x89, 0xFF, 0xEC, 0xB5, 0x28, 0x14, 0x23, 0xB6}}

type ISyndicationCategoryInterface interface {
	win32.IInspectableInterface
	Get_Label() string
	Put_Label(value string)
	Get_Scheme() string
	Put_Scheme(value string)
	Get_Term() string
	Put_Term(value string)
}

type ISyndicationCategoryVtbl struct {
	win32.IInspectableVtbl
	Get_Label  uintptr
	Put_Label  uintptr
	Get_Scheme uintptr
	Put_Scheme uintptr
	Get_Term   uintptr
	Put_Term   uintptr
}

type ISyndicationCategory struct {
	win32.IInspectable
}

func (this *ISyndicationCategory) Vtbl() *ISyndicationCategoryVtbl {
	return (*ISyndicationCategoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISyndicationCategory) Get_Label() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Label, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISyndicationCategory) Put_Label(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Label, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISyndicationCategory) Get_Scheme() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Scheme, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISyndicationCategory) Put_Scheme(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Scheme, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISyndicationCategory) Get_Term() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Term, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISyndicationCategory) Put_Term(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Term, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// AB42802F-49E0-4525-8AB2-AB45C02528FF
var IID_ISyndicationCategoryFactory = syscall.GUID{0xAB42802F, 0x49E0, 0x4525,
	[8]byte{0x8A, 0xB2, 0xAB, 0x45, 0xC0, 0x25, 0x28, 0xFF}}

type ISyndicationCategoryFactoryInterface interface {
	win32.IInspectableInterface
	CreateSyndicationCategory(term string) *ISyndicationCategory
	CreateSyndicationCategoryEx(term string, scheme string, label string) *ISyndicationCategory
}

type ISyndicationCategoryFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateSyndicationCategory   uintptr
	CreateSyndicationCategoryEx uintptr
}

type ISyndicationCategoryFactory struct {
	win32.IInspectable
}

func (this *ISyndicationCategoryFactory) Vtbl() *ISyndicationCategoryFactoryVtbl {
	return (*ISyndicationCategoryFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISyndicationCategoryFactory) CreateSyndicationCategory(term string) *ISyndicationCategory {
	var _result *ISyndicationCategory
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSyndicationCategory, uintptr(unsafe.Pointer(this)), NewHStr(term).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationCategoryFactory) CreateSyndicationCategoryEx(term string, scheme string, label string) *ISyndicationCategory {
	var _result *ISyndicationCategory
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSyndicationCategoryEx, uintptr(unsafe.Pointer(this)), NewHStr(term).Ptr, NewHStr(scheme).Ptr, NewHStr(label).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9E18A9B7-7249-4B45-B229-7DF895A5A1F5
var IID_ISyndicationClient = syscall.GUID{0x9E18A9B7, 0x7249, 0x4B45,
	[8]byte{0xB2, 0x29, 0x7D, 0xF8, 0x95, 0xA5, 0xA1, 0xF5}}

type ISyndicationClientInterface interface {
	win32.IInspectableInterface
	Get_ServerCredential() *IPasswordCredential
	Put_ServerCredential(value *IPasswordCredential)
	Get_ProxyCredential() *IPasswordCredential
	Put_ProxyCredential(value *IPasswordCredential)
	Get_MaxResponseBufferSize() uint32
	Put_MaxResponseBufferSize(value uint32)
	Get_Timeout() uint32
	Put_Timeout(value uint32)
	Get_BypassCacheOnRetrieve() bool
	Put_BypassCacheOnRetrieve(value bool)
	SetRequestHeader(name string, value string)
	RetrieveFeedAsync(uri *IUriRuntimeClass) *IAsyncOperationWithProgress[*ISyndicationFeed, RetrievalProgress]
}

type ISyndicationClientVtbl struct {
	win32.IInspectableVtbl
	Get_ServerCredential      uintptr
	Put_ServerCredential      uintptr
	Get_ProxyCredential       uintptr
	Put_ProxyCredential       uintptr
	Get_MaxResponseBufferSize uintptr
	Put_MaxResponseBufferSize uintptr
	Get_Timeout               uintptr
	Put_Timeout               uintptr
	Get_BypassCacheOnRetrieve uintptr
	Put_BypassCacheOnRetrieve uintptr
	SetRequestHeader          uintptr
	RetrieveFeedAsync         uintptr
}

type ISyndicationClient struct {
	win32.IInspectable
}

func (this *ISyndicationClient) Vtbl() *ISyndicationClientVtbl {
	return (*ISyndicationClientVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISyndicationClient) Get_ServerCredential() *IPasswordCredential {
	var _result *IPasswordCredential
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServerCredential, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationClient) Put_ServerCredential(value *IPasswordCredential) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ServerCredential, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ISyndicationClient) Get_ProxyCredential() *IPasswordCredential {
	var _result *IPasswordCredential
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProxyCredential, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationClient) Put_ProxyCredential(value *IPasswordCredential) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ProxyCredential, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ISyndicationClient) Get_MaxResponseBufferSize() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxResponseBufferSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISyndicationClient) Put_MaxResponseBufferSize(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxResponseBufferSize, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISyndicationClient) Get_Timeout() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timeout, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISyndicationClient) Put_Timeout(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Timeout, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISyndicationClient) Get_BypassCacheOnRetrieve() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BypassCacheOnRetrieve, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISyndicationClient) Put_BypassCacheOnRetrieve(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BypassCacheOnRetrieve, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ISyndicationClient) SetRequestHeader(name string, value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetRequestHeader, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISyndicationClient) RetrieveFeedAsync(uri *IUriRuntimeClass) *IAsyncOperationWithProgress[*ISyndicationFeed, RetrievalProgress] {
	var _result *IAsyncOperationWithProgress[*ISyndicationFeed, RetrievalProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RetrieveFeedAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2EC4B32C-A79B-4114-B29A-05DFFBAFB9A4
var IID_ISyndicationClientFactory = syscall.GUID{0x2EC4B32C, 0xA79B, 0x4114,
	[8]byte{0xB2, 0x9A, 0x05, 0xDF, 0xFB, 0xAF, 0xB9, 0xA4}}

type ISyndicationClientFactoryInterface interface {
	win32.IInspectableInterface
	CreateSyndicationClient(serverCredential *IPasswordCredential) *ISyndicationClient
}

type ISyndicationClientFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateSyndicationClient uintptr
}

type ISyndicationClientFactory struct {
	win32.IInspectable
}

func (this *ISyndicationClientFactory) Vtbl() *ISyndicationClientFactoryVtbl {
	return (*ISyndicationClientFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISyndicationClientFactory) CreateSyndicationClient(serverCredential *IPasswordCredential) *ISyndicationClient {
	var _result *ISyndicationClient
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSyndicationClient, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(serverCredential)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4641FEFE-0E55-40D0-B8D0-6A2CCBA9FC7C
var IID_ISyndicationContent = syscall.GUID{0x4641FEFE, 0x0E55, 0x40D0,
	[8]byte{0xB8, 0xD0, 0x6A, 0x2C, 0xCB, 0xA9, 0xFC, 0x7C}}

type ISyndicationContentInterface interface {
	win32.IInspectableInterface
	Get_SourceUri() *IUriRuntimeClass
	Put_SourceUri(value *IUriRuntimeClass)
}

type ISyndicationContentVtbl struct {
	win32.IInspectableVtbl
	Get_SourceUri uintptr
	Put_SourceUri uintptr
}

type ISyndicationContent struct {
	win32.IInspectable
}

func (this *ISyndicationContent) Vtbl() *ISyndicationContentVtbl {
	return (*ISyndicationContentVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISyndicationContent) Get_SourceUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SourceUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationContent) Put_SourceUri(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SourceUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 3D2FBB93-9520-4173-9388-7E2DF324A8A0
var IID_ISyndicationContentFactory = syscall.GUID{0x3D2FBB93, 0x9520, 0x4173,
	[8]byte{0x93, 0x88, 0x7E, 0x2D, 0xF3, 0x24, 0xA8, 0xA0}}

type ISyndicationContentFactoryInterface interface {
	win32.IInspectableInterface
	CreateSyndicationContent(text string, type_ SyndicationTextType) *ISyndicationContent
	CreateSyndicationContentWithSourceUri(sourceUri *IUriRuntimeClass) *ISyndicationContent
}

type ISyndicationContentFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateSyndicationContent              uintptr
	CreateSyndicationContentWithSourceUri uintptr
}

type ISyndicationContentFactory struct {
	win32.IInspectable
}

func (this *ISyndicationContentFactory) Vtbl() *ISyndicationContentFactoryVtbl {
	return (*ISyndicationContentFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISyndicationContentFactory) CreateSyndicationContent(text string, type_ SyndicationTextType) *ISyndicationContent {
	var _result *ISyndicationContent
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSyndicationContent, uintptr(unsafe.Pointer(this)), NewHStr(text).Ptr, uintptr(type_), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationContentFactory) CreateSyndicationContentWithSourceUri(sourceUri *IUriRuntimeClass) *ISyndicationContent {
	var _result *ISyndicationContent
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSyndicationContentWithSourceUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sourceUri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1FBB2361-45C7-4833-8AA0-BE5F3B58A7F4
var IID_ISyndicationErrorStatics = syscall.GUID{0x1FBB2361, 0x45C7, 0x4833,
	[8]byte{0x8A, 0xA0, 0xBE, 0x5F, 0x3B, 0x58, 0xA7, 0xF4}}

type ISyndicationErrorStaticsInterface interface {
	win32.IInspectableInterface
	GetStatus(hresult int32) SyndicationErrorStatus
}

type ISyndicationErrorStaticsVtbl struct {
	win32.IInspectableVtbl
	GetStatus uintptr
}

type ISyndicationErrorStatics struct {
	win32.IInspectable
}

func (this *ISyndicationErrorStatics) Vtbl() *ISyndicationErrorStaticsVtbl {
	return (*ISyndicationErrorStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISyndicationErrorStatics) GetStatus(hresult int32) SyndicationErrorStatus {
	var _result SyndicationErrorStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStatus, uintptr(unsafe.Pointer(this)), uintptr(hresult), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7FFE3CD2-5B66-4D62-8403-1BC10D910D6B
var IID_ISyndicationFeed = syscall.GUID{0x7FFE3CD2, 0x5B66, 0x4D62,
	[8]byte{0x84, 0x03, 0x1B, 0xC1, 0x0D, 0x91, 0x0D, 0x6B}}

type ISyndicationFeedInterface interface {
	win32.IInspectableInterface
	Get_Authors() *IVector[*ISyndicationPerson]
	Get_Categories() *IVector[*ISyndicationCategory]
	Get_Contributors() *IVector[*ISyndicationPerson]
	Get_Generator() *ISyndicationGenerator
	Put_Generator(value *ISyndicationGenerator)
	Get_IconUri() *IUriRuntimeClass
	Put_IconUri(value *IUriRuntimeClass)
	Get_Id() string
	Put_Id(value string)
	Get_Items() *IVector[*ISyndicationItem]
	Get_LastUpdatedTime() DateTime
	Put_LastUpdatedTime(value DateTime)
	Get_Links() *IVector[*ISyndicationLink]
	Get_ImageUri() *IUriRuntimeClass
	Put_ImageUri(value *IUriRuntimeClass)
	Get_Rights() *ISyndicationText
	Put_Rights(value *ISyndicationText)
	Get_Subtitle() *ISyndicationText
	Put_Subtitle(value *ISyndicationText)
	Get_Title() *ISyndicationText
	Put_Title(value *ISyndicationText)
	Get_FirstUri() *IUriRuntimeClass
	Get_LastUri() *IUriRuntimeClass
	Get_NextUri() *IUriRuntimeClass
	Get_PreviousUri() *IUriRuntimeClass
	Get_SourceFormat() SyndicationFormat
	Load(feed string)
	LoadFromXml(feedDocument unsafe.Pointer)
}

type ISyndicationFeedVtbl struct {
	win32.IInspectableVtbl
	Get_Authors         uintptr
	Get_Categories      uintptr
	Get_Contributors    uintptr
	Get_Generator       uintptr
	Put_Generator       uintptr
	Get_IconUri         uintptr
	Put_IconUri         uintptr
	Get_Id              uintptr
	Put_Id              uintptr
	Get_Items           uintptr
	Get_LastUpdatedTime uintptr
	Put_LastUpdatedTime uintptr
	Get_Links           uintptr
	Get_ImageUri        uintptr
	Put_ImageUri        uintptr
	Get_Rights          uintptr
	Put_Rights          uintptr
	Get_Subtitle        uintptr
	Put_Subtitle        uintptr
	Get_Title           uintptr
	Put_Title           uintptr
	Get_FirstUri        uintptr
	Get_LastUri         uintptr
	Get_NextUri         uintptr
	Get_PreviousUri     uintptr
	Get_SourceFormat    uintptr
	Load                uintptr
	LoadFromXml         uintptr
}

type ISyndicationFeed struct {
	win32.IInspectable
}

func (this *ISyndicationFeed) Vtbl() *ISyndicationFeedVtbl {
	return (*ISyndicationFeedVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISyndicationFeed) Get_Authors() *IVector[*ISyndicationPerson] {
	var _result *IVector[*ISyndicationPerson]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Authors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationFeed) Get_Categories() *IVector[*ISyndicationCategory] {
	var _result *IVector[*ISyndicationCategory]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Categories, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationFeed) Get_Contributors() *IVector[*ISyndicationPerson] {
	var _result *IVector[*ISyndicationPerson]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Contributors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationFeed) Get_Generator() *ISyndicationGenerator {
	var _result *ISyndicationGenerator
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Generator, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationFeed) Put_Generator(value *ISyndicationGenerator) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Generator, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ISyndicationFeed) Get_IconUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IconUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationFeed) Put_IconUri(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IconUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ISyndicationFeed) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISyndicationFeed) Put_Id(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Id, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISyndicationFeed) Get_Items() *IVector[*ISyndicationItem] {
	var _result *IVector[*ISyndicationItem]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Items, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationFeed) Get_LastUpdatedTime() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LastUpdatedTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISyndicationFeed) Put_LastUpdatedTime(value DateTime) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LastUpdatedTime, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ISyndicationFeed) Get_Links() *IVector[*ISyndicationLink] {
	var _result *IVector[*ISyndicationLink]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Links, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationFeed) Get_ImageUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ImageUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationFeed) Put_ImageUri(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ImageUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ISyndicationFeed) Get_Rights() *ISyndicationText {
	var _result *ISyndicationText
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Rights, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationFeed) Put_Rights(value *ISyndicationText) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Rights, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ISyndicationFeed) Get_Subtitle() *ISyndicationText {
	var _result *ISyndicationText
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Subtitle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationFeed) Put_Subtitle(value *ISyndicationText) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Subtitle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ISyndicationFeed) Get_Title() *ISyndicationText {
	var _result *ISyndicationText
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Title, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationFeed) Put_Title(value *ISyndicationText) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Title, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ISyndicationFeed) Get_FirstUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FirstUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationFeed) Get_LastUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LastUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationFeed) Get_NextUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NextUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationFeed) Get_PreviousUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PreviousUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationFeed) Get_SourceFormat() SyndicationFormat {
	var _result SyndicationFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SourceFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISyndicationFeed) Load(feed string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Load, uintptr(unsafe.Pointer(this)), NewHStr(feed).Ptr)
	_ = _hr
}

func (this *ISyndicationFeed) LoadFromXml(feedDocument unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LoadFromXml, uintptr(unsafe.Pointer(this)), uintptr(feedDocument))
	_ = _hr
}

// 23472232-8BE9-48B7-8934-6205131D9357
var IID_ISyndicationFeedFactory = syscall.GUID{0x23472232, 0x8BE9, 0x48B7,
	[8]byte{0x89, 0x34, 0x62, 0x05, 0x13, 0x1D, 0x93, 0x57}}

type ISyndicationFeedFactoryInterface interface {
	win32.IInspectableInterface
	CreateSyndicationFeed(title string, subtitle string, uri *IUriRuntimeClass) *ISyndicationFeed
}

type ISyndicationFeedFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateSyndicationFeed uintptr
}

type ISyndicationFeedFactory struct {
	win32.IInspectable
}

func (this *ISyndicationFeedFactory) Vtbl() *ISyndicationFeedFactoryVtbl {
	return (*ISyndicationFeedFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISyndicationFeedFactory) CreateSyndicationFeed(title string, subtitle string, uri *IUriRuntimeClass) *ISyndicationFeed {
	var _result *ISyndicationFeed
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSyndicationFeed, uintptr(unsafe.Pointer(this)), NewHStr(title).Ptr, NewHStr(subtitle).Ptr, uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9768B379-FB2B-4F6D-B41C-088A5868825C
var IID_ISyndicationGenerator = syscall.GUID{0x9768B379, 0xFB2B, 0x4F6D,
	[8]byte{0xB4, 0x1C, 0x08, 0x8A, 0x58, 0x68, 0x82, 0x5C}}

type ISyndicationGeneratorInterface interface {
	win32.IInspectableInterface
	Get_Text() string
	Put_Text(value string)
	Get_Uri() *IUriRuntimeClass
	Put_Uri(value *IUriRuntimeClass)
	Get_Version() string
	Put_Version(value string)
}

type ISyndicationGeneratorVtbl struct {
	win32.IInspectableVtbl
	Get_Text    uintptr
	Put_Text    uintptr
	Get_Uri     uintptr
	Put_Uri     uintptr
	Get_Version uintptr
	Put_Version uintptr
}

type ISyndicationGenerator struct {
	win32.IInspectable
}

func (this *ISyndicationGenerator) Vtbl() *ISyndicationGeneratorVtbl {
	return (*ISyndicationGeneratorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISyndicationGenerator) Get_Text() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Text, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISyndicationGenerator) Put_Text(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Text, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISyndicationGenerator) Get_Uri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Uri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationGenerator) Put_Uri(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Uri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ISyndicationGenerator) Get_Version() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Version, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISyndicationGenerator) Put_Version(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Version, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// A34083E3-1E26-4DBC-BA9D-1AB84BEFF97B
var IID_ISyndicationGeneratorFactory = syscall.GUID{0xA34083E3, 0x1E26, 0x4DBC,
	[8]byte{0xBA, 0x9D, 0x1A, 0xB8, 0x4B, 0xEF, 0xF9, 0x7B}}

type ISyndicationGeneratorFactoryInterface interface {
	win32.IInspectableInterface
	CreateSyndicationGenerator(text string) *ISyndicationGenerator
}

type ISyndicationGeneratorFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateSyndicationGenerator uintptr
}

type ISyndicationGeneratorFactory struct {
	win32.IInspectable
}

func (this *ISyndicationGeneratorFactory) Vtbl() *ISyndicationGeneratorFactoryVtbl {
	return (*ISyndicationGeneratorFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISyndicationGeneratorFactory) CreateSyndicationGenerator(text string) *ISyndicationGenerator {
	var _result *ISyndicationGenerator
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSyndicationGenerator, uintptr(unsafe.Pointer(this)), NewHStr(text).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 548DB883-C384-45C1-8AE8-A378C4EC486C
var IID_ISyndicationItem = syscall.GUID{0x548DB883, 0xC384, 0x45C1,
	[8]byte{0x8A, 0xE8, 0xA3, 0x78, 0xC4, 0xEC, 0x48, 0x6C}}

type ISyndicationItemInterface interface {
	win32.IInspectableInterface
	Get_Authors() *IVector[*ISyndicationPerson]
	Get_Categories() *IVector[*ISyndicationCategory]
	Get_Contributors() *IVector[*ISyndicationPerson]
	Get_Content() *ISyndicationContent
	Put_Content(value *ISyndicationContent)
	Get_Id() string
	Put_Id(value string)
	Get_LastUpdatedTime() DateTime
	Put_LastUpdatedTime(value DateTime)
	Get_Links() *IVector[*ISyndicationLink]
	Get_PublishedDate() DateTime
	Put_PublishedDate(value DateTime)
	Get_Rights() *ISyndicationText
	Put_Rights(value *ISyndicationText)
	Get_Source() *ISyndicationFeed
	Put_Source(value *ISyndicationFeed)
	Get_Summary() *ISyndicationText
	Put_Summary(value *ISyndicationText)
	Get_Title() *ISyndicationText
	Put_Title(value *ISyndicationText)
	Get_CommentsUri() *IUriRuntimeClass
	Put_CommentsUri(value *IUriRuntimeClass)
	Get_EditUri() *IUriRuntimeClass
	Get_EditMediaUri() *IUriRuntimeClass
	Get_ETag() string
	Get_ItemUri() *IUriRuntimeClass
	Load(item string)
	LoadFromXml(itemDocument unsafe.Pointer)
}

type ISyndicationItemVtbl struct {
	win32.IInspectableVtbl
	Get_Authors         uintptr
	Get_Categories      uintptr
	Get_Contributors    uintptr
	Get_Content         uintptr
	Put_Content         uintptr
	Get_Id              uintptr
	Put_Id              uintptr
	Get_LastUpdatedTime uintptr
	Put_LastUpdatedTime uintptr
	Get_Links           uintptr
	Get_PublishedDate   uintptr
	Put_PublishedDate   uintptr
	Get_Rights          uintptr
	Put_Rights          uintptr
	Get_Source          uintptr
	Put_Source          uintptr
	Get_Summary         uintptr
	Put_Summary         uintptr
	Get_Title           uintptr
	Put_Title           uintptr
	Get_CommentsUri     uintptr
	Put_CommentsUri     uintptr
	Get_EditUri         uintptr
	Get_EditMediaUri    uintptr
	Get_ETag            uintptr
	Get_ItemUri         uintptr
	Load                uintptr
	LoadFromXml         uintptr
}

type ISyndicationItem struct {
	win32.IInspectable
}

func (this *ISyndicationItem) Vtbl() *ISyndicationItemVtbl {
	return (*ISyndicationItemVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISyndicationItem) Get_Authors() *IVector[*ISyndicationPerson] {
	var _result *IVector[*ISyndicationPerson]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Authors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationItem) Get_Categories() *IVector[*ISyndicationCategory] {
	var _result *IVector[*ISyndicationCategory]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Categories, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationItem) Get_Contributors() *IVector[*ISyndicationPerson] {
	var _result *IVector[*ISyndicationPerson]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Contributors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationItem) Get_Content() *ISyndicationContent {
	var _result *ISyndicationContent
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Content, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationItem) Put_Content(value *ISyndicationContent) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Content, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ISyndicationItem) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISyndicationItem) Put_Id(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Id, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISyndicationItem) Get_LastUpdatedTime() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LastUpdatedTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISyndicationItem) Put_LastUpdatedTime(value DateTime) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LastUpdatedTime, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ISyndicationItem) Get_Links() *IVector[*ISyndicationLink] {
	var _result *IVector[*ISyndicationLink]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Links, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationItem) Get_PublishedDate() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PublishedDate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISyndicationItem) Put_PublishedDate(value DateTime) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PublishedDate, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ISyndicationItem) Get_Rights() *ISyndicationText {
	var _result *ISyndicationText
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Rights, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationItem) Put_Rights(value *ISyndicationText) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Rights, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ISyndicationItem) Get_Source() *ISyndicationFeed {
	var _result *ISyndicationFeed
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Source, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationItem) Put_Source(value *ISyndicationFeed) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Source, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ISyndicationItem) Get_Summary() *ISyndicationText {
	var _result *ISyndicationText
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Summary, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationItem) Put_Summary(value *ISyndicationText) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Summary, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ISyndicationItem) Get_Title() *ISyndicationText {
	var _result *ISyndicationText
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Title, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationItem) Put_Title(value *ISyndicationText) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Title, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ISyndicationItem) Get_CommentsUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CommentsUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationItem) Put_CommentsUri(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CommentsUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ISyndicationItem) Get_EditUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EditUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationItem) Get_EditMediaUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EditMediaUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationItem) Get_ETag() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ETag, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISyndicationItem) Get_ItemUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ItemUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationItem) Load(item string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Load, uintptr(unsafe.Pointer(this)), NewHStr(item).Ptr)
	_ = _hr
}

func (this *ISyndicationItem) LoadFromXml(itemDocument unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LoadFromXml, uintptr(unsafe.Pointer(this)), uintptr(itemDocument))
	_ = _hr
}

// 251D434F-7DB8-487A-85E4-10D191E66EBB
var IID_ISyndicationItemFactory = syscall.GUID{0x251D434F, 0x7DB8, 0x487A,
	[8]byte{0x85, 0xE4, 0x10, 0xD1, 0x91, 0xE6, 0x6E, 0xBB}}

type ISyndicationItemFactoryInterface interface {
	win32.IInspectableInterface
	CreateSyndicationItem(title string, content *ISyndicationContent, uri *IUriRuntimeClass) *ISyndicationItem
}

type ISyndicationItemFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateSyndicationItem uintptr
}

type ISyndicationItemFactory struct {
	win32.IInspectable
}

func (this *ISyndicationItemFactory) Vtbl() *ISyndicationItemFactoryVtbl {
	return (*ISyndicationItemFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISyndicationItemFactory) CreateSyndicationItem(title string, content *ISyndicationContent, uri *IUriRuntimeClass) *ISyndicationItem {
	var _result *ISyndicationItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSyndicationItem, uintptr(unsafe.Pointer(this)), NewHStr(title).Ptr, uintptr(unsafe.Pointer(content)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 27553ABD-A10E-41B5-86BD-9759086EB0C5
var IID_ISyndicationLink = syscall.GUID{0x27553ABD, 0xA10E, 0x41B5,
	[8]byte{0x86, 0xBD, 0x97, 0x59, 0x08, 0x6E, 0xB0, 0xC5}}

type ISyndicationLinkInterface interface {
	win32.IInspectableInterface
	Get_Length() uint32
	Put_Length(value uint32)
	Get_MediaType() string
	Put_MediaType(value string)
	Get_Relationship() string
	Put_Relationship(value string)
	Get_Title() string
	Put_Title(value string)
	Get_Uri() *IUriRuntimeClass
	Put_Uri(value *IUriRuntimeClass)
	Get_ResourceLanguage() string
	Put_ResourceLanguage(value string)
}

type ISyndicationLinkVtbl struct {
	win32.IInspectableVtbl
	Get_Length           uintptr
	Put_Length           uintptr
	Get_MediaType        uintptr
	Put_MediaType        uintptr
	Get_Relationship     uintptr
	Put_Relationship     uintptr
	Get_Title            uintptr
	Put_Title            uintptr
	Get_Uri              uintptr
	Put_Uri              uintptr
	Get_ResourceLanguage uintptr
	Put_ResourceLanguage uintptr
}

type ISyndicationLink struct {
	win32.IInspectable
}

func (this *ISyndicationLink) Vtbl() *ISyndicationLinkVtbl {
	return (*ISyndicationLinkVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISyndicationLink) Get_Length() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Length, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISyndicationLink) Put_Length(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Length, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISyndicationLink) Get_MediaType() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISyndicationLink) Put_MediaType(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MediaType, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISyndicationLink) Get_Relationship() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Relationship, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISyndicationLink) Put_Relationship(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Relationship, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISyndicationLink) Get_Title() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Title, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISyndicationLink) Put_Title(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Title, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISyndicationLink) Get_Uri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Uri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationLink) Put_Uri(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Uri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ISyndicationLink) Get_ResourceLanguage() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceLanguage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISyndicationLink) Put_ResourceLanguage(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ResourceLanguage, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// 5ED863D4-5535-48AC-98D4-C190995080B3
var IID_ISyndicationLinkFactory = syscall.GUID{0x5ED863D4, 0x5535, 0x48AC,
	[8]byte{0x98, 0xD4, 0xC1, 0x90, 0x99, 0x50, 0x80, 0xB3}}

type ISyndicationLinkFactoryInterface interface {
	win32.IInspectableInterface
	CreateSyndicationLink(uri *IUriRuntimeClass) *ISyndicationLink
	CreateSyndicationLinkEx(uri *IUriRuntimeClass, relationship string, title string, mediaType string, length uint32) *ISyndicationLink
}

type ISyndicationLinkFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateSyndicationLink   uintptr
	CreateSyndicationLinkEx uintptr
}

type ISyndicationLinkFactory struct {
	win32.IInspectable
}

func (this *ISyndicationLinkFactory) Vtbl() *ISyndicationLinkFactoryVtbl {
	return (*ISyndicationLinkFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISyndicationLinkFactory) CreateSyndicationLink(uri *IUriRuntimeClass) *ISyndicationLink {
	var _result *ISyndicationLink
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSyndicationLink, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationLinkFactory) CreateSyndicationLinkEx(uri *IUriRuntimeClass, relationship string, title string, mediaType string, length uint32) *ISyndicationLink {
	var _result *ISyndicationLink
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSyndicationLinkEx, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), NewHStr(relationship).Ptr, NewHStr(title).Ptr, NewHStr(mediaType).Ptr, uintptr(length), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 753CEF78-51F8-45C0-A9F5-F1719DEC3FB2
var IID_ISyndicationNode = syscall.GUID{0x753CEF78, 0x51F8, 0x45C0,
	[8]byte{0xA9, 0xF5, 0xF1, 0x71, 0x9D, 0xEC, 0x3F, 0xB2}}

type ISyndicationNodeInterface interface {
	win32.IInspectableInterface
	Get_NodeName() string
	Put_NodeName(value string)
	Get_NodeNamespace() string
	Put_NodeNamespace(value string)
	Get_NodeValue() string
	Put_NodeValue(value string)
	Get_Language() string
	Put_Language(value string)
	Get_BaseUri() *IUriRuntimeClass
	Put_BaseUri(value *IUriRuntimeClass)
	Get_AttributeExtensions() *IVector[*ISyndicationAttribute]
	Get_ElementExtensions() *IVector[*ISyndicationNode]
	GetXmlDocument(format SyndicationFormat) unsafe.Pointer
}

type ISyndicationNodeVtbl struct {
	win32.IInspectableVtbl
	Get_NodeName            uintptr
	Put_NodeName            uintptr
	Get_NodeNamespace       uintptr
	Put_NodeNamespace       uintptr
	Get_NodeValue           uintptr
	Put_NodeValue           uintptr
	Get_Language            uintptr
	Put_Language            uintptr
	Get_BaseUri             uintptr
	Put_BaseUri             uintptr
	Get_AttributeExtensions uintptr
	Get_ElementExtensions   uintptr
	GetXmlDocument          uintptr
}

type ISyndicationNode struct {
	win32.IInspectable
}

func (this *ISyndicationNode) Vtbl() *ISyndicationNodeVtbl {
	return (*ISyndicationNodeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISyndicationNode) Get_NodeName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NodeName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISyndicationNode) Put_NodeName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_NodeName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISyndicationNode) Get_NodeNamespace() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NodeNamespace, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISyndicationNode) Put_NodeNamespace(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_NodeNamespace, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISyndicationNode) Get_NodeValue() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NodeValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISyndicationNode) Put_NodeValue(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_NodeValue, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISyndicationNode) Get_Language() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Language, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISyndicationNode) Put_Language(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Language, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISyndicationNode) Get_BaseUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BaseUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationNode) Put_BaseUri(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BaseUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ISyndicationNode) Get_AttributeExtensions() *IVector[*ISyndicationAttribute] {
	var _result *IVector[*ISyndicationAttribute]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AttributeExtensions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationNode) Get_ElementExtensions() *IVector[*ISyndicationNode] {
	var _result *IVector[*ISyndicationNode]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ElementExtensions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationNode) GetXmlDocument(format SyndicationFormat) unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetXmlDocument, uintptr(unsafe.Pointer(this)), uintptr(format), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 12902188-4ACB-49A8-B777-A5EB92E18A79
var IID_ISyndicationNodeFactory = syscall.GUID{0x12902188, 0x4ACB, 0x49A8,
	[8]byte{0xB7, 0x77, 0xA5, 0xEB, 0x92, 0xE1, 0x8A, 0x79}}

type ISyndicationNodeFactoryInterface interface {
	win32.IInspectableInterface
	CreateSyndicationNode(nodeName string, nodeNamespace string, nodeValue string) *ISyndicationNode
}

type ISyndicationNodeFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateSyndicationNode uintptr
}

type ISyndicationNodeFactory struct {
	win32.IInspectable
}

func (this *ISyndicationNodeFactory) Vtbl() *ISyndicationNodeFactoryVtbl {
	return (*ISyndicationNodeFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISyndicationNodeFactory) CreateSyndicationNode(nodeName string, nodeNamespace string, nodeValue string) *ISyndicationNode {
	var _result *ISyndicationNode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSyndicationNode, uintptr(unsafe.Pointer(this)), NewHStr(nodeName).Ptr, NewHStr(nodeNamespace).Ptr, NewHStr(nodeValue).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FA1EE5DA-A7C6-4517-A096-0143FAF29327
var IID_ISyndicationPerson = syscall.GUID{0xFA1EE5DA, 0xA7C6, 0x4517,
	[8]byte{0xA0, 0x96, 0x01, 0x43, 0xFA, 0xF2, 0x93, 0x27}}

type ISyndicationPersonInterface interface {
	win32.IInspectableInterface
	Get_Email() string
	Put_Email(value string)
	Get_Name() string
	Put_Name(value string)
	Get_Uri() *IUriRuntimeClass
	Put_Uri(value *IUriRuntimeClass)
}

type ISyndicationPersonVtbl struct {
	win32.IInspectableVtbl
	Get_Email uintptr
	Put_Email uintptr
	Get_Name  uintptr
	Put_Name  uintptr
	Get_Uri   uintptr
	Put_Uri   uintptr
}

type ISyndicationPerson struct {
	win32.IInspectable
}

func (this *ISyndicationPerson) Vtbl() *ISyndicationPersonVtbl {
	return (*ISyndicationPersonVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISyndicationPerson) Get_Email() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Email, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISyndicationPerson) Put_Email(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Email, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISyndicationPerson) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISyndicationPerson) Put_Name(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Name, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISyndicationPerson) Get_Uri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Uri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationPerson) Put_Uri(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Uri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// DCF4886D-229D-4B58-A49B-F3D2F0F5C99F
var IID_ISyndicationPersonFactory = syscall.GUID{0xDCF4886D, 0x229D, 0x4B58,
	[8]byte{0xA4, 0x9B, 0xF3, 0xD2, 0xF0, 0xF5, 0xC9, 0x9F}}

type ISyndicationPersonFactoryInterface interface {
	win32.IInspectableInterface
	CreateSyndicationPerson(name string) *ISyndicationPerson
	CreateSyndicationPersonEx(name string, email string, uri *IUriRuntimeClass) *ISyndicationPerson
}

type ISyndicationPersonFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateSyndicationPerson   uintptr
	CreateSyndicationPersonEx uintptr
}

type ISyndicationPersonFactory struct {
	win32.IInspectable
}

func (this *ISyndicationPersonFactory) Vtbl() *ISyndicationPersonFactoryVtbl {
	return (*ISyndicationPersonFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISyndicationPersonFactory) CreateSyndicationPerson(name string) *ISyndicationPerson {
	var _result *ISyndicationPerson
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSyndicationPerson, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationPersonFactory) CreateSyndicationPersonEx(name string, email string, uri *IUriRuntimeClass) *ISyndicationPerson {
	var _result *ISyndicationPerson
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSyndicationPersonEx, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, NewHStr(email).Ptr, uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B9CC5E80-313A-4091-A2A6-243E0EE923F9
var IID_ISyndicationText = syscall.GUID{0xB9CC5E80, 0x313A, 0x4091,
	[8]byte{0xA2, 0xA6, 0x24, 0x3E, 0x0E, 0xE9, 0x23, 0xF9}}

type ISyndicationTextInterface interface {
	win32.IInspectableInterface
	Get_Text() string
	Put_Text(value string)
	Get_Type() string
	Put_Type(value string)
	Get_Xml() unsafe.Pointer
	Put_Xml(value unsafe.Pointer)
}

type ISyndicationTextVtbl struct {
	win32.IInspectableVtbl
	Get_Text uintptr
	Put_Text uintptr
	Get_Type uintptr
	Put_Type uintptr
	Get_Xml  uintptr
	Put_Xml  uintptr
}

type ISyndicationText struct {
	win32.IInspectable
}

func (this *ISyndicationText) Vtbl() *ISyndicationTextVtbl {
	return (*ISyndicationTextVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISyndicationText) Get_Text() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Text, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISyndicationText) Put_Text(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Text, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISyndicationText) Get_Type() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Type, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISyndicationText) Put_Type(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Type, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISyndicationText) Get_Xml() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Xml, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISyndicationText) Put_Xml(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Xml, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// EE7342F7-11C6-4B25-AB62-E596BD162946
var IID_ISyndicationTextFactory = syscall.GUID{0xEE7342F7, 0x11C6, 0x4B25,
	[8]byte{0xAB, 0x62, 0xE5, 0x96, 0xBD, 0x16, 0x29, 0x46}}

type ISyndicationTextFactoryInterface interface {
	win32.IInspectableInterface
	CreateSyndicationText(text string) *ISyndicationText
	CreateSyndicationTextEx(text string, type_ SyndicationTextType) *ISyndicationText
}

type ISyndicationTextFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateSyndicationText   uintptr
	CreateSyndicationTextEx uintptr
}

type ISyndicationTextFactory struct {
	win32.IInspectable
}

func (this *ISyndicationTextFactory) Vtbl() *ISyndicationTextFactoryVtbl {
	return (*ISyndicationTextFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISyndicationTextFactory) CreateSyndicationText(text string) *ISyndicationText {
	var _result *ISyndicationText
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSyndicationText, uintptr(unsafe.Pointer(this)), NewHStr(text).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISyndicationTextFactory) CreateSyndicationTextEx(text string, type_ SyndicationTextType) *ISyndicationText {
	var _result *ISyndicationText
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSyndicationTextEx, uintptr(unsafe.Pointer(this)), NewHStr(text).Ptr, uintptr(type_), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type SyndicationAttribute struct {
	RtClass
	*ISyndicationAttribute
}

func NewSyndicationAttribute() *SyndicationAttribute {
	hs := NewHStr("Windows.Web.Syndication.SyndicationAttribute")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &SyndicationAttribute{
		RtClass:               RtClass{PInspect: p},
		ISyndicationAttribute: (*ISyndicationAttribute)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewSyndicationAttribute_CreateSyndicationAttribute(attributeName string, attributeNamespace string, attributeValue string) *SyndicationAttribute {
	hs := NewHStr("Windows.Web.Syndication.SyndicationAttribute")
	var pFac *ISyndicationAttributeFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISyndicationAttributeFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISyndicationAttribute
	p = pFac.CreateSyndicationAttribute(attributeName, attributeNamespace, attributeValue)
	result := &SyndicationAttribute{
		RtClass:               RtClass{PInspect: &p.IInspectable},
		ISyndicationAttribute: p,
	}
	com.AddToScope(result)
	return result
}

type SyndicationCategory struct {
	RtClass
	*ISyndicationCategory
}

func NewSyndicationCategory() *SyndicationCategory {
	hs := NewHStr("Windows.Web.Syndication.SyndicationCategory")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &SyndicationCategory{
		RtClass:              RtClass{PInspect: p},
		ISyndicationCategory: (*ISyndicationCategory)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewSyndicationCategory_CreateSyndicationCategory(term string) *SyndicationCategory {
	hs := NewHStr("Windows.Web.Syndication.SyndicationCategory")
	var pFac *ISyndicationCategoryFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISyndicationCategoryFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISyndicationCategory
	p = pFac.CreateSyndicationCategory(term)
	result := &SyndicationCategory{
		RtClass:              RtClass{PInspect: &p.IInspectable},
		ISyndicationCategory: p,
	}
	com.AddToScope(result)
	return result
}

func NewSyndicationCategory_CreateSyndicationCategoryEx(term string, scheme string, label string) *SyndicationCategory {
	hs := NewHStr("Windows.Web.Syndication.SyndicationCategory")
	var pFac *ISyndicationCategoryFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISyndicationCategoryFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISyndicationCategory
	p = pFac.CreateSyndicationCategoryEx(term, scheme, label)
	result := &SyndicationCategory{
		RtClass:              RtClass{PInspect: &p.IInspectable},
		ISyndicationCategory: p,
	}
	com.AddToScope(result)
	return result
}

type SyndicationClient struct {
	RtClass
	*ISyndicationClient
}

func NewSyndicationClient() *SyndicationClient {
	hs := NewHStr("Windows.Web.Syndication.SyndicationClient")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &SyndicationClient{
		RtClass:            RtClass{PInspect: p},
		ISyndicationClient: (*ISyndicationClient)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewSyndicationClient_CreateSyndicationClient(serverCredential *IPasswordCredential) *SyndicationClient {
	hs := NewHStr("Windows.Web.Syndication.SyndicationClient")
	var pFac *ISyndicationClientFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISyndicationClientFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISyndicationClient
	p = pFac.CreateSyndicationClient(serverCredential)
	result := &SyndicationClient{
		RtClass:            RtClass{PInspect: &p.IInspectable},
		ISyndicationClient: p,
	}
	com.AddToScope(result)
	return result
}

type SyndicationContent struct {
	RtClass
	*ISyndicationContent
}

func NewSyndicationContent() *SyndicationContent {
	hs := NewHStr("Windows.Web.Syndication.SyndicationContent")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &SyndicationContent{
		RtClass:             RtClass{PInspect: p},
		ISyndicationContent: (*ISyndicationContent)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewSyndicationContent_CreateSyndicationContent(text string, type_ SyndicationTextType) *SyndicationContent {
	hs := NewHStr("Windows.Web.Syndication.SyndicationContent")
	var pFac *ISyndicationContentFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISyndicationContentFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISyndicationContent
	p = pFac.CreateSyndicationContent(text, type_)
	result := &SyndicationContent{
		RtClass:             RtClass{PInspect: &p.IInspectable},
		ISyndicationContent: p,
	}
	com.AddToScope(result)
	return result
}

func NewSyndicationContent_CreateSyndicationContentWithSourceUri(sourceUri *IUriRuntimeClass) *SyndicationContent {
	hs := NewHStr("Windows.Web.Syndication.SyndicationContent")
	var pFac *ISyndicationContentFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISyndicationContentFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISyndicationContent
	p = pFac.CreateSyndicationContentWithSourceUri(sourceUri)
	result := &SyndicationContent{
		RtClass:             RtClass{PInspect: &p.IInspectable},
		ISyndicationContent: p,
	}
	com.AddToScope(result)
	return result
}

type SyndicationError struct {
	RtClass
}

func NewISyndicationErrorStatics() *ISyndicationErrorStatics {
	var p *ISyndicationErrorStatics
	hs := NewHStr("Windows.Web.Syndication.SyndicationError")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISyndicationErrorStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type SyndicationFeed struct {
	RtClass
	*ISyndicationFeed
}

func NewSyndicationFeed() *SyndicationFeed {
	hs := NewHStr("Windows.Web.Syndication.SyndicationFeed")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &SyndicationFeed{
		RtClass:          RtClass{PInspect: p},
		ISyndicationFeed: (*ISyndicationFeed)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewSyndicationFeed_CreateSyndicationFeed(title string, subtitle string, uri *IUriRuntimeClass) *SyndicationFeed {
	hs := NewHStr("Windows.Web.Syndication.SyndicationFeed")
	var pFac *ISyndicationFeedFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISyndicationFeedFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISyndicationFeed
	p = pFac.CreateSyndicationFeed(title, subtitle, uri)
	result := &SyndicationFeed{
		RtClass:          RtClass{PInspect: &p.IInspectable},
		ISyndicationFeed: p,
	}
	com.AddToScope(result)
	return result
}

type SyndicationGenerator struct {
	RtClass
	*ISyndicationGenerator
}

func NewSyndicationGenerator() *SyndicationGenerator {
	hs := NewHStr("Windows.Web.Syndication.SyndicationGenerator")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &SyndicationGenerator{
		RtClass:               RtClass{PInspect: p},
		ISyndicationGenerator: (*ISyndicationGenerator)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewSyndicationGenerator_CreateSyndicationGenerator(text string) *SyndicationGenerator {
	hs := NewHStr("Windows.Web.Syndication.SyndicationGenerator")
	var pFac *ISyndicationGeneratorFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISyndicationGeneratorFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISyndicationGenerator
	p = pFac.CreateSyndicationGenerator(text)
	result := &SyndicationGenerator{
		RtClass:               RtClass{PInspect: &p.IInspectable},
		ISyndicationGenerator: p,
	}
	com.AddToScope(result)
	return result
}

type SyndicationItem struct {
	RtClass
	*ISyndicationItem
}

func NewSyndicationItem() *SyndicationItem {
	hs := NewHStr("Windows.Web.Syndication.SyndicationItem")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &SyndicationItem{
		RtClass:          RtClass{PInspect: p},
		ISyndicationItem: (*ISyndicationItem)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewSyndicationItem_CreateSyndicationItem(title string, content *ISyndicationContent, uri *IUriRuntimeClass) *SyndicationItem {
	hs := NewHStr("Windows.Web.Syndication.SyndicationItem")
	var pFac *ISyndicationItemFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISyndicationItemFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISyndicationItem
	p = pFac.CreateSyndicationItem(title, content, uri)
	result := &SyndicationItem{
		RtClass:          RtClass{PInspect: &p.IInspectable},
		ISyndicationItem: p,
	}
	com.AddToScope(result)
	return result
}

type SyndicationLink struct {
	RtClass
	*ISyndicationLink
}

func NewSyndicationLink() *SyndicationLink {
	hs := NewHStr("Windows.Web.Syndication.SyndicationLink")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &SyndicationLink{
		RtClass:          RtClass{PInspect: p},
		ISyndicationLink: (*ISyndicationLink)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewSyndicationLink_CreateSyndicationLink(uri *IUriRuntimeClass) *SyndicationLink {
	hs := NewHStr("Windows.Web.Syndication.SyndicationLink")
	var pFac *ISyndicationLinkFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISyndicationLinkFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISyndicationLink
	p = pFac.CreateSyndicationLink(uri)
	result := &SyndicationLink{
		RtClass:          RtClass{PInspect: &p.IInspectable},
		ISyndicationLink: p,
	}
	com.AddToScope(result)
	return result
}

func NewSyndicationLink_CreateSyndicationLinkEx(uri *IUriRuntimeClass, relationship string, title string, mediaType string, length uint32) *SyndicationLink {
	hs := NewHStr("Windows.Web.Syndication.SyndicationLink")
	var pFac *ISyndicationLinkFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISyndicationLinkFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISyndicationLink
	p = pFac.CreateSyndicationLinkEx(uri, relationship, title, mediaType, length)
	result := &SyndicationLink{
		RtClass:          RtClass{PInspect: &p.IInspectable},
		ISyndicationLink: p,
	}
	com.AddToScope(result)
	return result
}

type SyndicationNode struct {
	RtClass
	*ISyndicationNode
}

func NewSyndicationNode() *SyndicationNode {
	hs := NewHStr("Windows.Web.Syndication.SyndicationNode")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &SyndicationNode{
		RtClass:          RtClass{PInspect: p},
		ISyndicationNode: (*ISyndicationNode)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewSyndicationNode_CreateSyndicationNode(nodeName string, nodeNamespace string, nodeValue string) *SyndicationNode {
	hs := NewHStr("Windows.Web.Syndication.SyndicationNode")
	var pFac *ISyndicationNodeFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISyndicationNodeFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISyndicationNode
	p = pFac.CreateSyndicationNode(nodeName, nodeNamespace, nodeValue)
	result := &SyndicationNode{
		RtClass:          RtClass{PInspect: &p.IInspectable},
		ISyndicationNode: p,
	}
	com.AddToScope(result)
	return result
}

type SyndicationPerson struct {
	RtClass
	*ISyndicationPerson
}

func NewSyndicationPerson() *SyndicationPerson {
	hs := NewHStr("Windows.Web.Syndication.SyndicationPerson")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &SyndicationPerson{
		RtClass:            RtClass{PInspect: p},
		ISyndicationPerson: (*ISyndicationPerson)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewSyndicationPerson_CreateSyndicationPerson(name string) *SyndicationPerson {
	hs := NewHStr("Windows.Web.Syndication.SyndicationPerson")
	var pFac *ISyndicationPersonFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISyndicationPersonFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISyndicationPerson
	p = pFac.CreateSyndicationPerson(name)
	result := &SyndicationPerson{
		RtClass:            RtClass{PInspect: &p.IInspectable},
		ISyndicationPerson: p,
	}
	com.AddToScope(result)
	return result
}

func NewSyndicationPerson_CreateSyndicationPersonEx(name string, email string, uri *IUriRuntimeClass) *SyndicationPerson {
	hs := NewHStr("Windows.Web.Syndication.SyndicationPerson")
	var pFac *ISyndicationPersonFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISyndicationPersonFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISyndicationPerson
	p = pFac.CreateSyndicationPersonEx(name, email, uri)
	result := &SyndicationPerson{
		RtClass:            RtClass{PInspect: &p.IInspectable},
		ISyndicationPerson: p,
	}
	com.AddToScope(result)
	return result
}

type SyndicationText struct {
	RtClass
	*ISyndicationText
}

func NewSyndicationText() *SyndicationText {
	hs := NewHStr("Windows.Web.Syndication.SyndicationText")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &SyndicationText{
		RtClass:          RtClass{PInspect: p},
		ISyndicationText: (*ISyndicationText)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewSyndicationText_CreateSyndicationText(text string) *SyndicationText {
	hs := NewHStr("Windows.Web.Syndication.SyndicationText")
	var pFac *ISyndicationTextFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISyndicationTextFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISyndicationText
	p = pFac.CreateSyndicationText(text)
	result := &SyndicationText{
		RtClass:          RtClass{PInspect: &p.IInspectable},
		ISyndicationText: p,
	}
	com.AddToScope(result)
	return result
}

func NewSyndicationText_CreateSyndicationTextEx(text string, type_ SyndicationTextType) *SyndicationText {
	hs := NewHStr("Windows.Web.Syndication.SyndicationText")
	var pFac *ISyndicationTextFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISyndicationTextFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISyndicationText
	p = pFac.CreateSyndicationTextEx(text, type_)
	result := &SyndicationText{
		RtClass:          RtClass{PInspect: &p.IInspectable},
		ISyndicationText: p,
	}
	com.AddToScope(result)
	return result
}
