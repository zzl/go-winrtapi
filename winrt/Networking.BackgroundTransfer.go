package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type BackgroundTransferBehavior int32

const (
	BackgroundTransferBehavior_Parallel   BackgroundTransferBehavior = 0
	BackgroundTransferBehavior_Serialized BackgroundTransferBehavior = 1
)

// enum
type BackgroundTransferCostPolicy int32

const (
	BackgroundTransferCostPolicy_Default          BackgroundTransferCostPolicy = 0
	BackgroundTransferCostPolicy_UnrestrictedOnly BackgroundTransferCostPolicy = 1
	BackgroundTransferCostPolicy_Always           BackgroundTransferCostPolicy = 2
)

// enum
type BackgroundTransferPriority int32

const (
	BackgroundTransferPriority_Default BackgroundTransferPriority = 0
	BackgroundTransferPriority_High    BackgroundTransferPriority = 1
	BackgroundTransferPriority_Low     BackgroundTransferPriority = 2
)

// enum
type BackgroundTransferStatus int32

const (
	BackgroundTransferStatus_Idle                            BackgroundTransferStatus = 0
	BackgroundTransferStatus_Running                         BackgroundTransferStatus = 1
	BackgroundTransferStatus_PausedByApplication             BackgroundTransferStatus = 2
	BackgroundTransferStatus_PausedCostedNetwork             BackgroundTransferStatus = 3
	BackgroundTransferStatus_PausedNoNetwork                 BackgroundTransferStatus = 4
	BackgroundTransferStatus_Completed                       BackgroundTransferStatus = 5
	BackgroundTransferStatus_Canceled                        BackgroundTransferStatus = 6
	BackgroundTransferStatus_Error                           BackgroundTransferStatus = 7
	BackgroundTransferStatus_PausedRecoverableWebErrorStatus BackgroundTransferStatus = 8
	BackgroundTransferStatus_PausedSystemPolicy              BackgroundTransferStatus = 32
)

// structs

type BackgroundDownloadProgress struct {
	BytesReceived       uint64
	TotalBytesToReceive uint64
	Status              BackgroundTransferStatus
	HasResponseChanged  bool
	HasRestarted        bool
}

type BackgroundTransferFileRange struct {
	Offset uint64
	Length uint64
}

type BackgroundUploadProgress struct {
	BytesReceived       uint64
	BytesSent           uint64
	TotalBytesToReceive uint64
	TotalBytesToSend    uint64
	Status              BackgroundTransferStatus
	HasResponseChanged  bool
	HasRestarted        bool
}

// interfaces

// C1C79333-6649-4B1D-A826-A4B3DD234D0B
var IID_IBackgroundDownloader = syscall.GUID{0xC1C79333, 0x6649, 0x4B1D,
	[8]byte{0xA8, 0x26, 0xA4, 0xB3, 0xDD, 0x23, 0x4D, 0x0B}}

type IBackgroundDownloaderInterface interface {
	win32.IInspectableInterface
	CreateDownload(uri *IUriRuntimeClass, resultFile *IStorageFile) *IDownloadOperation
	CreateDownloadFromFile(uri *IUriRuntimeClass, resultFile *IStorageFile, requestBodyFile *IStorageFile) *IDownloadOperation
	CreateDownloadAsync(uri *IUriRuntimeClass, resultFile *IStorageFile, requestBodyStream *IInputStream) *IAsyncOperation[*IDownloadOperation]
}

type IBackgroundDownloaderVtbl struct {
	win32.IInspectableVtbl
	CreateDownload         uintptr
	CreateDownloadFromFile uintptr
	CreateDownloadAsync    uintptr
}

type IBackgroundDownloader struct {
	win32.IInspectable
}

func (this *IBackgroundDownloader) Vtbl() *IBackgroundDownloaderVtbl {
	return (*IBackgroundDownloaderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundDownloader) CreateDownload(uri *IUriRuntimeClass, resultFile *IStorageFile) *IDownloadOperation {
	var _result *IDownloadOperation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDownload, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(resultFile)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundDownloader) CreateDownloadFromFile(uri *IUriRuntimeClass, resultFile *IStorageFile, requestBodyFile *IStorageFile) *IDownloadOperation {
	var _result *IDownloadOperation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDownloadFromFile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(resultFile)), uintptr(unsafe.Pointer(requestBodyFile)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundDownloader) CreateDownloadAsync(uri *IUriRuntimeClass, resultFile *IStorageFile, requestBodyStream *IInputStream) *IAsyncOperation[*IDownloadOperation] {
	var _result *IAsyncOperation[*IDownloadOperation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDownloadAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(resultFile)), uintptr(unsafe.Pointer(requestBodyStream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A94A5847-348D-4A35-890E-8A1EF3798479
var IID_IBackgroundDownloader2 = syscall.GUID{0xA94A5847, 0x348D, 0x4A35,
	[8]byte{0x89, 0x0E, 0x8A, 0x1E, 0xF3, 0x79, 0x84, 0x79}}

type IBackgroundDownloader2Interface interface {
	win32.IInspectableInterface
	Get_TransferGroup() *IBackgroundTransferGroup
	Put_TransferGroup(value *IBackgroundTransferGroup)
	Get_SuccessToastNotification() *IToastNotification
	Put_SuccessToastNotification(value *IToastNotification)
	Get_FailureToastNotification() *IToastNotification
	Put_FailureToastNotification(value *IToastNotification)
	Get_SuccessTileNotification() *ITileNotification
	Put_SuccessTileNotification(value *ITileNotification)
	Get_FailureTileNotification() *ITileNotification
	Put_FailureTileNotification(value *ITileNotification)
}

type IBackgroundDownloader2Vtbl struct {
	win32.IInspectableVtbl
	Get_TransferGroup            uintptr
	Put_TransferGroup            uintptr
	Get_SuccessToastNotification uintptr
	Put_SuccessToastNotification uintptr
	Get_FailureToastNotification uintptr
	Put_FailureToastNotification uintptr
	Get_SuccessTileNotification  uintptr
	Put_SuccessTileNotification  uintptr
	Get_FailureTileNotification  uintptr
	Put_FailureTileNotification  uintptr
}

type IBackgroundDownloader2 struct {
	win32.IInspectable
}

func (this *IBackgroundDownloader2) Vtbl() *IBackgroundDownloader2Vtbl {
	return (*IBackgroundDownloader2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundDownloader2) Get_TransferGroup() *IBackgroundTransferGroup {
	var _result *IBackgroundTransferGroup
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TransferGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundDownloader2) Put_TransferGroup(value *IBackgroundTransferGroup) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TransferGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IBackgroundDownloader2) Get_SuccessToastNotification() *IToastNotification {
	var _result *IToastNotification
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SuccessToastNotification, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundDownloader2) Put_SuccessToastNotification(value *IToastNotification) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SuccessToastNotification, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IBackgroundDownloader2) Get_FailureToastNotification() *IToastNotification {
	var _result *IToastNotification
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FailureToastNotification, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundDownloader2) Put_FailureToastNotification(value *IToastNotification) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FailureToastNotification, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IBackgroundDownloader2) Get_SuccessTileNotification() *ITileNotification {
	var _result *ITileNotification
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SuccessTileNotification, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundDownloader2) Put_SuccessTileNotification(value *ITileNotification) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SuccessTileNotification, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IBackgroundDownloader2) Get_FailureTileNotification() *ITileNotification {
	var _result *ITileNotification
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FailureTileNotification, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundDownloader2) Put_FailureTileNotification(value *ITileNotification) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FailureTileNotification, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// D11A8C48-86E8-48E2-B615-6976AABF861D
var IID_IBackgroundDownloader3 = syscall.GUID{0xD11A8C48, 0x86E8, 0x48E2,
	[8]byte{0xB6, 0x15, 0x69, 0x76, 0xAA, 0xBF, 0x86, 0x1D}}

type IBackgroundDownloader3Interface interface {
	win32.IInspectableInterface
	Get_CompletionGroup() *IBackgroundTransferCompletionGroup
}

type IBackgroundDownloader3Vtbl struct {
	win32.IInspectableVtbl
	Get_CompletionGroup uintptr
}

type IBackgroundDownloader3 struct {
	win32.IInspectable
}

func (this *IBackgroundDownloader3) Vtbl() *IBackgroundDownloader3Vtbl {
	return (*IBackgroundDownloader3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundDownloader3) Get_CompletionGroup() *IBackgroundTransferCompletionGroup {
	var _result *IBackgroundTransferCompletionGroup
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CompletionGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 26836C24-D89E-46F4-A29A-4F4D4F144155
var IID_IBackgroundDownloaderFactory = syscall.GUID{0x26836C24, 0xD89E, 0x46F4,
	[8]byte{0xA2, 0x9A, 0x4F, 0x4D, 0x4F, 0x14, 0x41, 0x55}}

type IBackgroundDownloaderFactoryInterface interface {
	win32.IInspectableInterface
	CreateWithCompletionGroup(completionGroup *IBackgroundTransferCompletionGroup) *IBackgroundDownloader
}

type IBackgroundDownloaderFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateWithCompletionGroup uintptr
}

type IBackgroundDownloaderFactory struct {
	win32.IInspectable
}

func (this *IBackgroundDownloaderFactory) Vtbl() *IBackgroundDownloaderFactoryVtbl {
	return (*IBackgroundDownloaderFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundDownloaderFactory) CreateWithCompletionGroup(completionGroup *IBackgroundTransferCompletionGroup) *IBackgroundDownloader {
	var _result *IBackgroundDownloader
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithCompletionGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(completionGroup)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 52A65A35-C64E-426C-9919-540D0D21A650
var IID_IBackgroundDownloaderStaticMethods = syscall.GUID{0x52A65A35, 0xC64E, 0x426C,
	[8]byte{0x99, 0x19, 0x54, 0x0D, 0x0D, 0x21, 0xA6, 0x50}}

type IBackgroundDownloaderStaticMethodsInterface interface {
	win32.IInspectableInterface
	GetCurrentDownloadsAsync() *IAsyncOperation[*IVectorView[*IDownloadOperation]]
	GetCurrentDownloadsForGroupAsync(group string) *IAsyncOperation[*IVectorView[*IDownloadOperation]]
}

type IBackgroundDownloaderStaticMethodsVtbl struct {
	win32.IInspectableVtbl
	GetCurrentDownloadsAsync         uintptr
	GetCurrentDownloadsForGroupAsync uintptr
}

type IBackgroundDownloaderStaticMethods struct {
	win32.IInspectable
}

func (this *IBackgroundDownloaderStaticMethods) Vtbl() *IBackgroundDownloaderStaticMethodsVtbl {
	return (*IBackgroundDownloaderStaticMethodsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundDownloaderStaticMethods) GetCurrentDownloadsAsync() *IAsyncOperation[*IVectorView[*IDownloadOperation]] {
	var _result *IAsyncOperation[*IVectorView[*IDownloadOperation]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentDownloadsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundDownloaderStaticMethods) GetCurrentDownloadsForGroupAsync(group string) *IAsyncOperation[*IVectorView[*IDownloadOperation]] {
	var _result *IAsyncOperation[*IVectorView[*IDownloadOperation]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentDownloadsForGroupAsync, uintptr(unsafe.Pointer(this)), NewHStr(group).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2FAA1327-1AD4-4CA5-B2CD-08DBF0746AFE
var IID_IBackgroundDownloaderStaticMethods2 = syscall.GUID{0x2FAA1327, 0x1AD4, 0x4CA5,
	[8]byte{0xB2, 0xCD, 0x08, 0xDB, 0xF0, 0x74, 0x6A, 0xFE}}

type IBackgroundDownloaderStaticMethods2Interface interface {
	win32.IInspectableInterface
	GetCurrentDownloadsForTransferGroupAsync(group *IBackgroundTransferGroup) *IAsyncOperation[*IVectorView[*IDownloadOperation]]
}

type IBackgroundDownloaderStaticMethods2Vtbl struct {
	win32.IInspectableVtbl
	GetCurrentDownloadsForTransferGroupAsync uintptr
}

type IBackgroundDownloaderStaticMethods2 struct {
	win32.IInspectable
}

func (this *IBackgroundDownloaderStaticMethods2) Vtbl() *IBackgroundDownloaderStaticMethods2Vtbl {
	return (*IBackgroundDownloaderStaticMethods2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundDownloaderStaticMethods2) GetCurrentDownloadsForTransferGroupAsync(group *IBackgroundTransferGroup) *IAsyncOperation[*IVectorView[*IDownloadOperation]] {
	var _result *IAsyncOperation[*IVectorView[*IDownloadOperation]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentDownloadsForTransferGroupAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(group)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5D14E906-9266-4808-BD71-5925F2A3130A
var IID_IBackgroundDownloaderUserConsent = syscall.GUID{0x5D14E906, 0x9266, 0x4808,
	[8]byte{0xBD, 0x71, 0x59, 0x25, 0xF2, 0xA3, 0x13, 0x0A}}

type IBackgroundDownloaderUserConsentInterface interface {
	win32.IInspectableInterface
	RequestUnconstrainedDownloadsAsync(operations *IIterable[*IDownloadOperation]) *IAsyncOperation[*IUnconstrainedTransferRequestResult]
}

type IBackgroundDownloaderUserConsentVtbl struct {
	win32.IInspectableVtbl
	RequestUnconstrainedDownloadsAsync uintptr
}

type IBackgroundDownloaderUserConsent struct {
	win32.IInspectable
}

func (this *IBackgroundDownloaderUserConsent) Vtbl() *IBackgroundDownloaderUserConsentVtbl {
	return (*IBackgroundDownloaderUserConsentVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundDownloaderUserConsent) RequestUnconstrainedDownloadsAsync(operations *IIterable[*IDownloadOperation]) *IAsyncOperation[*IUnconstrainedTransferRequestResult] {
	var _result *IAsyncOperation[*IUnconstrainedTransferRequestResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestUnconstrainedDownloadsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(operations)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2A9DA250-C769-458C-AFE8-FEB8D4D3B2EF
var IID_IBackgroundTransferBase = syscall.GUID{0x2A9DA250, 0xC769, 0x458C,
	[8]byte{0xAF, 0xE8, 0xFE, 0xB8, 0xD4, 0xD3, 0xB2, 0xEF}}

type IBackgroundTransferBaseInterface interface {
	win32.IInspectableInterface
	SetRequestHeader(headerName string, headerValue string)
	Get_ServerCredential() *IPasswordCredential
	Put_ServerCredential(credential *IPasswordCredential)
	Get_ProxyCredential() *IPasswordCredential
	Put_ProxyCredential(credential *IPasswordCredential)
	Get_Method() string
	Put_Method(value string)
	Get_Group() string
	Put_Group(value string)
	Get_CostPolicy() BackgroundTransferCostPolicy
	Put_CostPolicy(value BackgroundTransferCostPolicy)
}

type IBackgroundTransferBaseVtbl struct {
	win32.IInspectableVtbl
	SetRequestHeader     uintptr
	Get_ServerCredential uintptr
	Put_ServerCredential uintptr
	Get_ProxyCredential  uintptr
	Put_ProxyCredential  uintptr
	Get_Method           uintptr
	Put_Method           uintptr
	Get_Group            uintptr
	Put_Group            uintptr
	Get_CostPolicy       uintptr
	Put_CostPolicy       uintptr
}

type IBackgroundTransferBase struct {
	win32.IInspectable
}

func (this *IBackgroundTransferBase) Vtbl() *IBackgroundTransferBaseVtbl {
	return (*IBackgroundTransferBaseVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTransferBase) SetRequestHeader(headerName string, headerValue string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetRequestHeader, uintptr(unsafe.Pointer(this)), NewHStr(headerName).Ptr, NewHStr(headerValue).Ptr)
	_ = _hr
}

func (this *IBackgroundTransferBase) Get_ServerCredential() *IPasswordCredential {
	var _result *IPasswordCredential
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServerCredential, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundTransferBase) Put_ServerCredential(credential *IPasswordCredential) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ServerCredential, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(credential)))
	_ = _hr
}

func (this *IBackgroundTransferBase) Get_ProxyCredential() *IPasswordCredential {
	var _result *IPasswordCredential
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProxyCredential, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundTransferBase) Put_ProxyCredential(credential *IPasswordCredential) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ProxyCredential, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(credential)))
	_ = _hr
}

func (this *IBackgroundTransferBase) Get_Method() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Method, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBackgroundTransferBase) Put_Method(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Method, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IBackgroundTransferBase) Get_Group() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Group, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBackgroundTransferBase) Put_Group(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Group, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IBackgroundTransferBase) Get_CostPolicy() BackgroundTransferCostPolicy {
	var _result BackgroundTransferCostPolicy
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CostPolicy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundTransferBase) Put_CostPolicy(value BackgroundTransferCostPolicy) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CostPolicy, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 2D930225-986B-574D-7950-0ADD47F5D706
var IID_IBackgroundTransferCompletionGroup = syscall.GUID{0x2D930225, 0x986B, 0x574D,
	[8]byte{0x79, 0x50, 0x0A, 0xDD, 0x47, 0xF5, 0xD7, 0x06}}

type IBackgroundTransferCompletionGroupInterface interface {
	win32.IInspectableInterface
	Get_Trigger() *IBackgroundTrigger
	Get_IsEnabled() bool
	Enable()
}

type IBackgroundTransferCompletionGroupVtbl struct {
	win32.IInspectableVtbl
	Get_Trigger   uintptr
	Get_IsEnabled uintptr
	Enable        uintptr
}

type IBackgroundTransferCompletionGroup struct {
	win32.IInspectable
}

func (this *IBackgroundTransferCompletionGroup) Vtbl() *IBackgroundTransferCompletionGroupVtbl {
	return (*IBackgroundTransferCompletionGroupVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTransferCompletionGroup) Get_Trigger() *IBackgroundTrigger {
	var _result *IBackgroundTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Trigger, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundTransferCompletionGroup) Get_IsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundTransferCompletionGroup) Enable() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Enable, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 7B6BE286-6E47-5136-7FCB-FA4389F46F5B
var IID_IBackgroundTransferCompletionGroupTriggerDetails = syscall.GUID{0x7B6BE286, 0x6E47, 0x5136,
	[8]byte{0x7F, 0xCB, 0xFA, 0x43, 0x89, 0xF4, 0x6F, 0x5B}}

type IBackgroundTransferCompletionGroupTriggerDetailsInterface interface {
	win32.IInspectableInterface
	Get_Downloads() *IVectorView[*IDownloadOperation]
	Get_Uploads() *IVectorView[*IUploadOperation]
}

type IBackgroundTransferCompletionGroupTriggerDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_Downloads uintptr
	Get_Uploads   uintptr
}

type IBackgroundTransferCompletionGroupTriggerDetails struct {
	win32.IInspectable
}

func (this *IBackgroundTransferCompletionGroupTriggerDetails) Vtbl() *IBackgroundTransferCompletionGroupTriggerDetailsVtbl {
	return (*IBackgroundTransferCompletionGroupTriggerDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTransferCompletionGroupTriggerDetails) Get_Downloads() *IVectorView[*IDownloadOperation] {
	var _result *IVectorView[*IDownloadOperation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Downloads, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundTransferCompletionGroupTriggerDetails) Get_Uploads() *IVectorView[*IUploadOperation] {
	var _result *IVectorView[*IUploadOperation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Uploads, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E8E15657-D7D1-4ED8-838E-674AC217ACE6
var IID_IBackgroundTransferContentPart = syscall.GUID{0xE8E15657, 0xD7D1, 0x4ED8,
	[8]byte{0x83, 0x8E, 0x67, 0x4A, 0xC2, 0x17, 0xAC, 0xE6}}

type IBackgroundTransferContentPartInterface interface {
	win32.IInspectableInterface
	SetHeader(headerName string, headerValue string)
	SetText(value string)
	SetFile(value *IStorageFile)
}

type IBackgroundTransferContentPartVtbl struct {
	win32.IInspectableVtbl
	SetHeader uintptr
	SetText   uintptr
	SetFile   uintptr
}

type IBackgroundTransferContentPart struct {
	win32.IInspectable
}

func (this *IBackgroundTransferContentPart) Vtbl() *IBackgroundTransferContentPartVtbl {
	return (*IBackgroundTransferContentPartVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTransferContentPart) SetHeader(headerName string, headerValue string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetHeader, uintptr(unsafe.Pointer(this)), NewHStr(headerName).Ptr, NewHStr(headerValue).Ptr)
	_ = _hr
}

func (this *IBackgroundTransferContentPart) SetText(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetText, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IBackgroundTransferContentPart) SetFile(value *IStorageFile) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetFile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 90EF98A9-7A01-4A0B-9F80-A0B0BB370F8D
var IID_IBackgroundTransferContentPartFactory = syscall.GUID{0x90EF98A9, 0x7A01, 0x4A0B,
	[8]byte{0x9F, 0x80, 0xA0, 0xB0, 0xBB, 0x37, 0x0F, 0x8D}}

type IBackgroundTransferContentPartFactoryInterface interface {
	win32.IInspectableInterface
	CreateWithName(name string) *IBackgroundTransferContentPart
	CreateWithNameAndFileName(name string, fileName string) *IBackgroundTransferContentPart
}

type IBackgroundTransferContentPartFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateWithName            uintptr
	CreateWithNameAndFileName uintptr
}

type IBackgroundTransferContentPartFactory struct {
	win32.IInspectable
}

func (this *IBackgroundTransferContentPartFactory) Vtbl() *IBackgroundTransferContentPartFactoryVtbl {
	return (*IBackgroundTransferContentPartFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTransferContentPartFactory) CreateWithName(name string) *IBackgroundTransferContentPart {
	var _result *IBackgroundTransferContentPart
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithName, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundTransferContentPartFactory) CreateWithNameAndFileName(name string, fileName string) *IBackgroundTransferContentPart {
	var _result *IBackgroundTransferContentPart
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithNameAndFileName, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, NewHStr(fileName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AAD33B04-1192-4BF4-8B68-39C5ADD244E2
var IID_IBackgroundTransferErrorStaticMethods = syscall.GUID{0xAAD33B04, 0x1192, 0x4BF4,
	[8]byte{0x8B, 0x68, 0x39, 0xC5, 0xAD, 0xD2, 0x44, 0xE2}}

type IBackgroundTransferErrorStaticMethodsInterface interface {
	win32.IInspectableInterface
	GetStatus(hresult int32) WebErrorStatus
}

type IBackgroundTransferErrorStaticMethodsVtbl struct {
	win32.IInspectableVtbl
	GetStatus uintptr
}

type IBackgroundTransferErrorStaticMethods struct {
	win32.IInspectable
}

func (this *IBackgroundTransferErrorStaticMethods) Vtbl() *IBackgroundTransferErrorStaticMethodsVtbl {
	return (*IBackgroundTransferErrorStaticMethodsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTransferErrorStaticMethods) GetStatus(hresult int32) WebErrorStatus {
	var _result WebErrorStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStatus, uintptr(unsafe.Pointer(this)), uintptr(hresult), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D8C3E3E4-6459-4540-85EB-AAA1C8903677
var IID_IBackgroundTransferGroup = syscall.GUID{0xD8C3E3E4, 0x6459, 0x4540,
	[8]byte{0x85, 0xEB, 0xAA, 0xA1, 0xC8, 0x90, 0x36, 0x77}}

type IBackgroundTransferGroupInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Get_TransferBehavior() BackgroundTransferBehavior
	Put_TransferBehavior(value BackgroundTransferBehavior)
}

type IBackgroundTransferGroupVtbl struct {
	win32.IInspectableVtbl
	Get_Name             uintptr
	Get_TransferBehavior uintptr
	Put_TransferBehavior uintptr
}

type IBackgroundTransferGroup struct {
	win32.IInspectable
}

func (this *IBackgroundTransferGroup) Vtbl() *IBackgroundTransferGroupVtbl {
	return (*IBackgroundTransferGroupVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTransferGroup) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBackgroundTransferGroup) Get_TransferBehavior() BackgroundTransferBehavior {
	var _result BackgroundTransferBehavior
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TransferBehavior, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundTransferGroup) Put_TransferBehavior(value BackgroundTransferBehavior) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TransferBehavior, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 02EC50B2-7D18-495B-AA22-32A97D45D3E2
var IID_IBackgroundTransferGroupStatics = syscall.GUID{0x02EC50B2, 0x7D18, 0x495B,
	[8]byte{0xAA, 0x22, 0x32, 0xA9, 0x7D, 0x45, 0xD3, 0xE2}}

type IBackgroundTransferGroupStaticsInterface interface {
	win32.IInspectableInterface
	CreateGroup(name string) *IBackgroundTransferGroup
}

type IBackgroundTransferGroupStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateGroup uintptr
}

type IBackgroundTransferGroupStatics struct {
	win32.IInspectable
}

func (this *IBackgroundTransferGroupStatics) Vtbl() *IBackgroundTransferGroupStaticsVtbl {
	return (*IBackgroundTransferGroupStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTransferGroupStatics) CreateGroup(name string) *IBackgroundTransferGroup {
	var _result *IBackgroundTransferGroup
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateGroup, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DED06846-90CA-44FB-8FB1-124154C0D539
var IID_IBackgroundTransferOperation = syscall.GUID{0xDED06846, 0x90CA, 0x44FB,
	[8]byte{0x8F, 0xB1, 0x12, 0x41, 0x54, 0xC0, 0xD5, 0x39}}

type IBackgroundTransferOperationInterface interface {
	win32.IInspectableInterface
	Get_Guid() syscall.GUID
	Get_RequestedUri() *IUriRuntimeClass
	Get_Method() string
	Get_Group() string
	Get_CostPolicy() BackgroundTransferCostPolicy
	Put_CostPolicy(value BackgroundTransferCostPolicy)
	GetResultStreamAt(position uint64) *IInputStream
	GetResponseInformation() *IResponseInformation
}

type IBackgroundTransferOperationVtbl struct {
	win32.IInspectableVtbl
	Get_Guid               uintptr
	Get_RequestedUri       uintptr
	Get_Method             uintptr
	Get_Group              uintptr
	Get_CostPolicy         uintptr
	Put_CostPolicy         uintptr
	GetResultStreamAt      uintptr
	GetResponseInformation uintptr
}

type IBackgroundTransferOperation struct {
	win32.IInspectable
}

func (this *IBackgroundTransferOperation) Vtbl() *IBackgroundTransferOperationVtbl {
	return (*IBackgroundTransferOperationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTransferOperation) Get_Guid() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Guid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundTransferOperation) Get_RequestedUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestedUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundTransferOperation) Get_Method() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Method, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBackgroundTransferOperation) Get_Group() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Group, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBackgroundTransferOperation) Get_CostPolicy() BackgroundTransferCostPolicy {
	var _result BackgroundTransferCostPolicy
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CostPolicy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundTransferOperation) Put_CostPolicy(value BackgroundTransferCostPolicy) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CostPolicy, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IBackgroundTransferOperation) GetResultStreamAt(position uint64) *IInputStream {
	var _result *IInputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetResultStreamAt, uintptr(unsafe.Pointer(this)), uintptr(position), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundTransferOperation) GetResponseInformation() *IResponseInformation {
	var _result *IResponseInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetResponseInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 04854327-5254-4B3A-915E-0AA49275C0F9
var IID_IBackgroundTransferOperationPriority = syscall.GUID{0x04854327, 0x5254, 0x4B3A,
	[8]byte{0x91, 0x5E, 0x0A, 0xA4, 0x92, 0x75, 0xC0, 0xF9}}

type IBackgroundTransferOperationPriorityInterface interface {
	win32.IInspectableInterface
	Get_Priority() BackgroundTransferPriority
	Put_Priority(value BackgroundTransferPriority)
}

type IBackgroundTransferOperationPriorityVtbl struct {
	win32.IInspectableVtbl
	Get_Priority uintptr
	Put_Priority uintptr
}

type IBackgroundTransferOperationPriority struct {
	win32.IInspectable
}

func (this *IBackgroundTransferOperationPriority) Vtbl() *IBackgroundTransferOperationPriorityVtbl {
	return (*IBackgroundTransferOperationPriorityVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTransferOperationPriority) Get_Priority() BackgroundTransferPriority {
	var _result BackgroundTransferPriority
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Priority, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundTransferOperationPriority) Put_Priority(value BackgroundTransferPriority) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Priority, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 3EBC7453-BF48-4A88-9248-B0C165184F5C
var IID_IBackgroundTransferRangesDownloadedEventArgs = syscall.GUID{0x3EBC7453, 0xBF48, 0x4A88,
	[8]byte{0x92, 0x48, 0xB0, 0xC1, 0x65, 0x18, 0x4F, 0x5C}}

type IBackgroundTransferRangesDownloadedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_WasDownloadRestarted() bool
	Get_AddedRanges() *IVector[BackgroundTransferFileRange]
	GetDeferral() *IDeferral
}

type IBackgroundTransferRangesDownloadedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_WasDownloadRestarted uintptr
	Get_AddedRanges          uintptr
	GetDeferral              uintptr
}

type IBackgroundTransferRangesDownloadedEventArgs struct {
	win32.IInspectable
}

func (this *IBackgroundTransferRangesDownloadedEventArgs) Vtbl() *IBackgroundTransferRangesDownloadedEventArgsVtbl {
	return (*IBackgroundTransferRangesDownloadedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundTransferRangesDownloadedEventArgs) Get_WasDownloadRestarted() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WasDownloadRestarted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundTransferRangesDownloadedEventArgs) Get_AddedRanges() *IVector[BackgroundTransferFileRange] {
	var _result *IVector[BackgroundTransferFileRange]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AddedRanges, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundTransferRangesDownloadedEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C595C9AE-CEAD-465B-8801-C55AC90A01CE
var IID_IBackgroundUploader = syscall.GUID{0xC595C9AE, 0xCEAD, 0x465B,
	[8]byte{0x88, 0x01, 0xC5, 0x5A, 0xC9, 0x0A, 0x01, 0xCE}}

type IBackgroundUploaderInterface interface {
	win32.IInspectableInterface
	CreateUpload(uri *IUriRuntimeClass, sourceFile *IStorageFile) *IUploadOperation
	CreateUploadFromStreamAsync(uri *IUriRuntimeClass, sourceStream *IInputStream) *IAsyncOperation[*IUploadOperation]
	CreateUploadWithFormDataAndAutoBoundaryAsync(uri *IUriRuntimeClass, parts *IIterable[*IBackgroundTransferContentPart]) *IAsyncOperation[*IUploadOperation]
	CreateUploadWithSubTypeAsync(uri *IUriRuntimeClass, parts *IIterable[*IBackgroundTransferContentPart], subType string) *IAsyncOperation[*IUploadOperation]
	CreateUploadWithSubTypeAndBoundaryAsync(uri *IUriRuntimeClass, parts *IIterable[*IBackgroundTransferContentPart], subType string, boundary string) *IAsyncOperation[*IUploadOperation]
}

type IBackgroundUploaderVtbl struct {
	win32.IInspectableVtbl
	CreateUpload                                 uintptr
	CreateUploadFromStreamAsync                  uintptr
	CreateUploadWithFormDataAndAutoBoundaryAsync uintptr
	CreateUploadWithSubTypeAsync                 uintptr
	CreateUploadWithSubTypeAndBoundaryAsync      uintptr
}

type IBackgroundUploader struct {
	win32.IInspectable
}

func (this *IBackgroundUploader) Vtbl() *IBackgroundUploaderVtbl {
	return (*IBackgroundUploaderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundUploader) CreateUpload(uri *IUriRuntimeClass, sourceFile *IStorageFile) *IUploadOperation {
	var _result *IUploadOperation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateUpload, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(sourceFile)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundUploader) CreateUploadFromStreamAsync(uri *IUriRuntimeClass, sourceStream *IInputStream) *IAsyncOperation[*IUploadOperation] {
	var _result *IAsyncOperation[*IUploadOperation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateUploadFromStreamAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(sourceStream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundUploader) CreateUploadWithFormDataAndAutoBoundaryAsync(uri *IUriRuntimeClass, parts *IIterable[*IBackgroundTransferContentPart]) *IAsyncOperation[*IUploadOperation] {
	var _result *IAsyncOperation[*IUploadOperation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateUploadWithFormDataAndAutoBoundaryAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(parts)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundUploader) CreateUploadWithSubTypeAsync(uri *IUriRuntimeClass, parts *IIterable[*IBackgroundTransferContentPart], subType string) *IAsyncOperation[*IUploadOperation] {
	var _result *IAsyncOperation[*IUploadOperation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateUploadWithSubTypeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(parts)), NewHStr(subType).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundUploader) CreateUploadWithSubTypeAndBoundaryAsync(uri *IUriRuntimeClass, parts *IIterable[*IBackgroundTransferContentPart], subType string, boundary string) *IAsyncOperation[*IUploadOperation] {
	var _result *IAsyncOperation[*IUploadOperation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateUploadWithSubTypeAndBoundaryAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(parts)), NewHStr(subType).Ptr, NewHStr(boundary).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8E0612CE-0C34-4463-807F-198A1B8BD4AD
var IID_IBackgroundUploader2 = syscall.GUID{0x8E0612CE, 0x0C34, 0x4463,
	[8]byte{0x80, 0x7F, 0x19, 0x8A, 0x1B, 0x8B, 0xD4, 0xAD}}

type IBackgroundUploader2Interface interface {
	win32.IInspectableInterface
	Get_TransferGroup() *IBackgroundTransferGroup
	Put_TransferGroup(value *IBackgroundTransferGroup)
	Get_SuccessToastNotification() *IToastNotification
	Put_SuccessToastNotification(value *IToastNotification)
	Get_FailureToastNotification() *IToastNotification
	Put_FailureToastNotification(value *IToastNotification)
	Get_SuccessTileNotification() *ITileNotification
	Put_SuccessTileNotification(value *ITileNotification)
	Get_FailureTileNotification() *ITileNotification
	Put_FailureTileNotification(value *ITileNotification)
}

type IBackgroundUploader2Vtbl struct {
	win32.IInspectableVtbl
	Get_TransferGroup            uintptr
	Put_TransferGroup            uintptr
	Get_SuccessToastNotification uintptr
	Put_SuccessToastNotification uintptr
	Get_FailureToastNotification uintptr
	Put_FailureToastNotification uintptr
	Get_SuccessTileNotification  uintptr
	Put_SuccessTileNotification  uintptr
	Get_FailureTileNotification  uintptr
	Put_FailureTileNotification  uintptr
}

type IBackgroundUploader2 struct {
	win32.IInspectable
}

func (this *IBackgroundUploader2) Vtbl() *IBackgroundUploader2Vtbl {
	return (*IBackgroundUploader2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundUploader2) Get_TransferGroup() *IBackgroundTransferGroup {
	var _result *IBackgroundTransferGroup
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TransferGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundUploader2) Put_TransferGroup(value *IBackgroundTransferGroup) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TransferGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IBackgroundUploader2) Get_SuccessToastNotification() *IToastNotification {
	var _result *IToastNotification
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SuccessToastNotification, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundUploader2) Put_SuccessToastNotification(value *IToastNotification) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SuccessToastNotification, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IBackgroundUploader2) Get_FailureToastNotification() *IToastNotification {
	var _result *IToastNotification
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FailureToastNotification, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundUploader2) Put_FailureToastNotification(value *IToastNotification) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FailureToastNotification, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IBackgroundUploader2) Get_SuccessTileNotification() *ITileNotification {
	var _result *ITileNotification
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SuccessTileNotification, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundUploader2) Put_SuccessTileNotification(value *ITileNotification) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SuccessTileNotification, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IBackgroundUploader2) Get_FailureTileNotification() *ITileNotification {
	var _result *ITileNotification
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FailureTileNotification, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundUploader2) Put_FailureTileNotification(value *ITileNotification) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FailureTileNotification, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// B95E9439-5BF0-4B3A-8C47-2C6199A854B9
var IID_IBackgroundUploader3 = syscall.GUID{0xB95E9439, 0x5BF0, 0x4B3A,
	[8]byte{0x8C, 0x47, 0x2C, 0x61, 0x99, 0xA8, 0x54, 0xB9}}

type IBackgroundUploader3Interface interface {
	win32.IInspectableInterface
	Get_CompletionGroup() *IBackgroundTransferCompletionGroup
}

type IBackgroundUploader3Vtbl struct {
	win32.IInspectableVtbl
	Get_CompletionGroup uintptr
}

type IBackgroundUploader3 struct {
	win32.IInspectable
}

func (this *IBackgroundUploader3) Vtbl() *IBackgroundUploader3Vtbl {
	return (*IBackgroundUploader3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundUploader3) Get_CompletionGroup() *IBackgroundTransferCompletionGroup {
	var _result *IBackgroundTransferCompletionGroup
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CompletionGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 736203C7-10E7-48A0-AC3C-1AC71095EC57
var IID_IBackgroundUploaderFactory = syscall.GUID{0x736203C7, 0x10E7, 0x48A0,
	[8]byte{0xAC, 0x3C, 0x1A, 0xC7, 0x10, 0x95, 0xEC, 0x57}}

type IBackgroundUploaderFactoryInterface interface {
	win32.IInspectableInterface
	CreateWithCompletionGroup(completionGroup *IBackgroundTransferCompletionGroup) *IBackgroundUploader
}

type IBackgroundUploaderFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateWithCompletionGroup uintptr
}

type IBackgroundUploaderFactory struct {
	win32.IInspectable
}

func (this *IBackgroundUploaderFactory) Vtbl() *IBackgroundUploaderFactoryVtbl {
	return (*IBackgroundUploaderFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundUploaderFactory) CreateWithCompletionGroup(completionGroup *IBackgroundTransferCompletionGroup) *IBackgroundUploader {
	var _result *IBackgroundUploader
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithCompletionGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(completionGroup)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F2875CFB-9B05-4741-9121-740A83E247DF
var IID_IBackgroundUploaderStaticMethods = syscall.GUID{0xF2875CFB, 0x9B05, 0x4741,
	[8]byte{0x91, 0x21, 0x74, 0x0A, 0x83, 0xE2, 0x47, 0xDF}}

type IBackgroundUploaderStaticMethodsInterface interface {
	win32.IInspectableInterface
	GetCurrentUploadsAsync() *IAsyncOperation[*IVectorView[*IUploadOperation]]
	GetCurrentUploadsForGroupAsync(group string) *IAsyncOperation[*IVectorView[*IUploadOperation]]
}

type IBackgroundUploaderStaticMethodsVtbl struct {
	win32.IInspectableVtbl
	GetCurrentUploadsAsync         uintptr
	GetCurrentUploadsForGroupAsync uintptr
}

type IBackgroundUploaderStaticMethods struct {
	win32.IInspectable
}

func (this *IBackgroundUploaderStaticMethods) Vtbl() *IBackgroundUploaderStaticMethodsVtbl {
	return (*IBackgroundUploaderStaticMethodsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundUploaderStaticMethods) GetCurrentUploadsAsync() *IAsyncOperation[*IVectorView[*IUploadOperation]] {
	var _result *IAsyncOperation[*IVectorView[*IUploadOperation]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentUploadsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundUploaderStaticMethods) GetCurrentUploadsForGroupAsync(group string) *IAsyncOperation[*IVectorView[*IUploadOperation]] {
	var _result *IAsyncOperation[*IVectorView[*IUploadOperation]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentUploadsForGroupAsync, uintptr(unsafe.Pointer(this)), NewHStr(group).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E919AC62-EA08-42F0-A2AC-07E467549080
var IID_IBackgroundUploaderStaticMethods2 = syscall.GUID{0xE919AC62, 0xEA08, 0x42F0,
	[8]byte{0xA2, 0xAC, 0x07, 0xE4, 0x67, 0x54, 0x90, 0x80}}

type IBackgroundUploaderStaticMethods2Interface interface {
	win32.IInspectableInterface
	GetCurrentUploadsForTransferGroupAsync(group *IBackgroundTransferGroup) *IAsyncOperation[*IVectorView[*IUploadOperation]]
}

type IBackgroundUploaderStaticMethods2Vtbl struct {
	win32.IInspectableVtbl
	GetCurrentUploadsForTransferGroupAsync uintptr
}

type IBackgroundUploaderStaticMethods2 struct {
	win32.IInspectable
}

func (this *IBackgroundUploaderStaticMethods2) Vtbl() *IBackgroundUploaderStaticMethods2Vtbl {
	return (*IBackgroundUploaderStaticMethods2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundUploaderStaticMethods2) GetCurrentUploadsForTransferGroupAsync(group *IBackgroundTransferGroup) *IAsyncOperation[*IVectorView[*IUploadOperation]] {
	var _result *IAsyncOperation[*IVectorView[*IUploadOperation]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentUploadsForTransferGroupAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(group)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3BB384CB-0760-461D-907F-5138F84D44C1
var IID_IBackgroundUploaderUserConsent = syscall.GUID{0x3BB384CB, 0x0760, 0x461D,
	[8]byte{0x90, 0x7F, 0x51, 0x38, 0xF8, 0x4D, 0x44, 0xC1}}

type IBackgroundUploaderUserConsentInterface interface {
	win32.IInspectableInterface
	RequestUnconstrainedUploadsAsync(operations *IIterable[*IUploadOperation]) *IAsyncOperation[*IUnconstrainedTransferRequestResult]
}

type IBackgroundUploaderUserConsentVtbl struct {
	win32.IInspectableVtbl
	RequestUnconstrainedUploadsAsync uintptr
}

type IBackgroundUploaderUserConsent struct {
	win32.IInspectable
}

func (this *IBackgroundUploaderUserConsent) Vtbl() *IBackgroundUploaderUserConsentVtbl {
	return (*IBackgroundUploaderUserConsentVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundUploaderUserConsent) RequestUnconstrainedUploadsAsync(operations *IIterable[*IUploadOperation]) *IAsyncOperation[*IUnconstrainedTransferRequestResult] {
	var _result *IAsyncOperation[*IUnconstrainedTransferRequestResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestUnconstrainedUploadsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(operations)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A8D6F754-7DC1-4CD9-8810-2A6AA9417E11
var IID_IContentPrefetcher = syscall.GUID{0xA8D6F754, 0x7DC1, 0x4CD9,
	[8]byte{0x88, 0x10, 0x2A, 0x6A, 0xA9, 0x41, 0x7E, 0x11}}

type IContentPrefetcherInterface interface {
	win32.IInspectableInterface
	Get_ContentUris() *IVector[*IUriRuntimeClass]
	Put_IndirectContentUri(value *IUriRuntimeClass)
	Get_IndirectContentUri() *IUriRuntimeClass
}

type IContentPrefetcherVtbl struct {
	win32.IInspectableVtbl
	Get_ContentUris        uintptr
	Put_IndirectContentUri uintptr
	Get_IndirectContentUri uintptr
}

type IContentPrefetcher struct {
	win32.IInspectable
}

func (this *IContentPrefetcher) Vtbl() *IContentPrefetcherVtbl {
	return (*IContentPrefetcherVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IContentPrefetcher) Get_ContentUris() *IVector[*IUriRuntimeClass] {
	var _result *IVector[*IUriRuntimeClass]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentUris, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IContentPrefetcher) Put_IndirectContentUri(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IndirectContentUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IContentPrefetcher) Get_IndirectContentUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IndirectContentUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E361FD08-132A-4FDE-A7CC-FCB0E66523AF
var IID_IContentPrefetcherTime = syscall.GUID{0xE361FD08, 0x132A, 0x4FDE,
	[8]byte{0xA7, 0xCC, 0xFC, 0xB0, 0xE6, 0x65, 0x23, 0xAF}}

type IContentPrefetcherTimeInterface interface {
	win32.IInspectableInterface
	Get_LastSuccessfulPrefetchTime() *IReference[DateTime]
}

type IContentPrefetcherTimeVtbl struct {
	win32.IInspectableVtbl
	Get_LastSuccessfulPrefetchTime uintptr
}

type IContentPrefetcherTime struct {
	win32.IInspectable
}

func (this *IContentPrefetcherTime) Vtbl() *IContentPrefetcherTimeVtbl {
	return (*IContentPrefetcherTimeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IContentPrefetcherTime) Get_LastSuccessfulPrefetchTime() *IReference[DateTime] {
	var _result *IReference[DateTime]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LastSuccessfulPrefetchTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BD87EBB0-5714-4E09-BA68-BEF73903B0D7
var IID_IDownloadOperation = syscall.GUID{0xBD87EBB0, 0x5714, 0x4E09,
	[8]byte{0xBA, 0x68, 0xBE, 0xF7, 0x39, 0x03, 0xB0, 0xD7}}

type IDownloadOperationInterface interface {
	win32.IInspectableInterface
	Get_ResultFile() *IStorageFile
	Get_Progress() BackgroundDownloadProgress
	StartAsync() *IAsyncOperationWithProgress[*IDownloadOperation, *IDownloadOperation]
	AttachAsync() *IAsyncOperationWithProgress[*IDownloadOperation, *IDownloadOperation]
	Pause()
	Resume()
}

type IDownloadOperationVtbl struct {
	win32.IInspectableVtbl
	Get_ResultFile uintptr
	Get_Progress   uintptr
	StartAsync     uintptr
	AttachAsync    uintptr
	Pause          uintptr
	Resume         uintptr
}

type IDownloadOperation struct {
	win32.IInspectable
}

func (this *IDownloadOperation) Vtbl() *IDownloadOperationVtbl {
	return (*IDownloadOperationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDownloadOperation) Get_ResultFile() *IStorageFile {
	var _result *IStorageFile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResultFile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDownloadOperation) Get_Progress() BackgroundDownloadProgress {
	var _result BackgroundDownloadProgress
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Progress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDownloadOperation) StartAsync() *IAsyncOperationWithProgress[*IDownloadOperation, *IDownloadOperation] {
	var _result *IAsyncOperationWithProgress[*IDownloadOperation, *IDownloadOperation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDownloadOperation) AttachAsync() *IAsyncOperationWithProgress[*IDownloadOperation, *IDownloadOperation] {
	var _result *IAsyncOperationWithProgress[*IDownloadOperation, *IDownloadOperation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AttachAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDownloadOperation) Pause() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Pause, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IDownloadOperation) Resume() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Resume, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// A3CCED40-8F9C-4353-9CD4-290DEE387C38
var IID_IDownloadOperation2 = syscall.GUID{0xA3CCED40, 0x8F9C, 0x4353,
	[8]byte{0x9C, 0xD4, 0x29, 0x0D, 0xEE, 0x38, 0x7C, 0x38}}

type IDownloadOperation2Interface interface {
	win32.IInspectableInterface
	Get_TransferGroup() *IBackgroundTransferGroup
}

type IDownloadOperation2Vtbl struct {
	win32.IInspectableVtbl
	Get_TransferGroup uintptr
}

type IDownloadOperation2 struct {
	win32.IInspectable
}

func (this *IDownloadOperation2) Vtbl() *IDownloadOperation2Vtbl {
	return (*IDownloadOperation2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDownloadOperation2) Get_TransferGroup() *IBackgroundTransferGroup {
	var _result *IBackgroundTransferGroup
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TransferGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5027351C-7D5E-4ADC-B8D3-DF5C6031B9CC
var IID_IDownloadOperation3 = syscall.GUID{0x5027351C, 0x7D5E, 0x4ADC,
	[8]byte{0xB8, 0xD3, 0xDF, 0x5C, 0x60, 0x31, 0xB9, 0xCC}}

type IDownloadOperation3Interface interface {
	win32.IInspectableInterface
	Get_IsRandomAccessRequired() bool
	Put_IsRandomAccessRequired(value bool)
	GetResultRandomAccessStreamReference() *IRandomAccessStreamReference
	GetDownloadedRanges() *IVector[BackgroundTransferFileRange]
	Add_RangesDownloaded(eventHandler TypedEventHandler[*IDownloadOperation, *IBackgroundTransferRangesDownloadedEventArgs]) EventRegistrationToken
	Remove_RangesDownloaded(eventCookie EventRegistrationToken)
	Put_RequestedUri(value *IUriRuntimeClass)
	Get_RecoverableWebErrorStatuses() *IVector[WebErrorStatus]
	Get_CurrentWebErrorStatus() *IReference[WebErrorStatus]
}

type IDownloadOperation3Vtbl struct {
	win32.IInspectableVtbl
	Get_IsRandomAccessRequired           uintptr
	Put_IsRandomAccessRequired           uintptr
	GetResultRandomAccessStreamReference uintptr
	GetDownloadedRanges                  uintptr
	Add_RangesDownloaded                 uintptr
	Remove_RangesDownloaded              uintptr
	Put_RequestedUri                     uintptr
	Get_RecoverableWebErrorStatuses      uintptr
	Get_CurrentWebErrorStatus            uintptr
}

type IDownloadOperation3 struct {
	win32.IInspectable
}

func (this *IDownloadOperation3) Vtbl() *IDownloadOperation3Vtbl {
	return (*IDownloadOperation3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDownloadOperation3) Get_IsRandomAccessRequired() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsRandomAccessRequired, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDownloadOperation3) Put_IsRandomAccessRequired(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsRandomAccessRequired, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IDownloadOperation3) GetResultRandomAccessStreamReference() *IRandomAccessStreamReference {
	var _result *IRandomAccessStreamReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetResultRandomAccessStreamReference, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDownloadOperation3) GetDownloadedRanges() *IVector[BackgroundTransferFileRange] {
	var _result *IVector[BackgroundTransferFileRange]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDownloadedRanges, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDownloadOperation3) Add_RangesDownloaded(eventHandler TypedEventHandler[*IDownloadOperation, *IBackgroundTransferRangesDownloadedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_RangesDownloaded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(eventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDownloadOperation3) Remove_RangesDownloaded(eventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_RangesDownloaded, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&eventCookie)))
	_ = _hr
}

func (this *IDownloadOperation3) Put_RequestedUri(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RequestedUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IDownloadOperation3) Get_RecoverableWebErrorStatuses() *IVector[WebErrorStatus] {
	var _result *IVector[WebErrorStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RecoverableWebErrorStatuses, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDownloadOperation3) Get_CurrentWebErrorStatus() *IReference[WebErrorStatus] {
	var _result *IReference[WebErrorStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentWebErrorStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0CDAAEF4-8CEF-404A-966D-F058400BED80
var IID_IDownloadOperation4 = syscall.GUID{0x0CDAAEF4, 0x8CEF, 0x404A,
	[8]byte{0x96, 0x6D, 0xF0, 0x58, 0x40, 0x0B, 0xED, 0x80}}

type IDownloadOperation4Interface interface {
	win32.IInspectableInterface
	MakeCurrentInTransferGroup()
}

type IDownloadOperation4Vtbl struct {
	win32.IInspectableVtbl
	MakeCurrentInTransferGroup uintptr
}

type IDownloadOperation4 struct {
	win32.IInspectable
}

func (this *IDownloadOperation4) Vtbl() *IDownloadOperation4Vtbl {
	return (*IDownloadOperation4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDownloadOperation4) MakeCurrentInTransferGroup() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MakeCurrentInTransferGroup, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// A699A86F-5590-463A-B8D6-1E491A2760A5
var IID_IDownloadOperation5 = syscall.GUID{0xA699A86F, 0x5590, 0x463A,
	[8]byte{0xB8, 0xD6, 0x1E, 0x49, 0x1A, 0x27, 0x60, 0xA5}}

type IDownloadOperation5Interface interface {
	win32.IInspectableInterface
	SetRequestHeader(headerName string, headerValue string)
	RemoveRequestHeader(headerName string)
}

type IDownloadOperation5Vtbl struct {
	win32.IInspectableVtbl
	SetRequestHeader    uintptr
	RemoveRequestHeader uintptr
}

type IDownloadOperation5 struct {
	win32.IInspectable
}

func (this *IDownloadOperation5) Vtbl() *IDownloadOperation5Vtbl {
	return (*IDownloadOperation5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDownloadOperation5) SetRequestHeader(headerName string, headerValue string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetRequestHeader, uintptr(unsafe.Pointer(this)), NewHStr(headerName).Ptr, NewHStr(headerValue).Ptr)
	_ = _hr
}

func (this *IDownloadOperation5) RemoveRequestHeader(headerName string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveRequestHeader, uintptr(unsafe.Pointer(this)), NewHStr(headerName).Ptr)
	_ = _hr
}

// F8BB9A12-F713-4792-8B68-D9D297F91D2E
var IID_IResponseInformation = syscall.GUID{0xF8BB9A12, 0xF713, 0x4792,
	[8]byte{0x8B, 0x68, 0xD9, 0xD2, 0x97, 0xF9, 0x1D, 0x2E}}

type IResponseInformationInterface interface {
	win32.IInspectableInterface
	Get_IsResumable() bool
	Get_ActualUri() *IUriRuntimeClass
	Get_StatusCode() uint32
	Get_Headers() *IMapView[string, string]
}

type IResponseInformationVtbl struct {
	win32.IInspectableVtbl
	Get_IsResumable uintptr
	Get_ActualUri   uintptr
	Get_StatusCode  uintptr
	Get_Headers     uintptr
}

type IResponseInformation struct {
	win32.IInspectable
}

func (this *IResponseInformation) Vtbl() *IResponseInformationVtbl {
	return (*IResponseInformationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IResponseInformation) Get_IsResumable() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsResumable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IResponseInformation) Get_ActualUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ActualUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IResponseInformation) Get_StatusCode() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StatusCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IResponseInformation) Get_Headers() *IMapView[string, string] {
	var _result *IMapView[string, string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Headers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4C24B81F-D944-4112-A98E-6A69522B7EBB
var IID_IUnconstrainedTransferRequestResult = syscall.GUID{0x4C24B81F, 0xD944, 0x4112,
	[8]byte{0xA9, 0x8E, 0x6A, 0x69, 0x52, 0x2B, 0x7E, 0xBB}}

type IUnconstrainedTransferRequestResultInterface interface {
	win32.IInspectableInterface
	Get_IsUnconstrained() bool
}

type IUnconstrainedTransferRequestResultVtbl struct {
	win32.IInspectableVtbl
	Get_IsUnconstrained uintptr
}

type IUnconstrainedTransferRequestResult struct {
	win32.IInspectable
}

func (this *IUnconstrainedTransferRequestResult) Vtbl() *IUnconstrainedTransferRequestResultVtbl {
	return (*IUnconstrainedTransferRequestResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUnconstrainedTransferRequestResult) Get_IsUnconstrained() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsUnconstrained, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 3E5624E0-7389-434C-8B35-427FD36BBDAE
var IID_IUploadOperation = syscall.GUID{0x3E5624E0, 0x7389, 0x434C,
	[8]byte{0x8B, 0x35, 0x42, 0x7F, 0xD3, 0x6B, 0xBD, 0xAE}}

type IUploadOperationInterface interface {
	win32.IInspectableInterface
	Get_SourceFile() *IStorageFile
	Get_Progress() BackgroundUploadProgress
	StartAsync() *IAsyncOperationWithProgress[*IUploadOperation, *IUploadOperation]
	AttachAsync() *IAsyncOperationWithProgress[*IUploadOperation, *IUploadOperation]
}

type IUploadOperationVtbl struct {
	win32.IInspectableVtbl
	Get_SourceFile uintptr
	Get_Progress   uintptr
	StartAsync     uintptr
	AttachAsync    uintptr
}

type IUploadOperation struct {
	win32.IInspectable
}

func (this *IUploadOperation) Vtbl() *IUploadOperationVtbl {
	return (*IUploadOperationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUploadOperation) Get_SourceFile() *IStorageFile {
	var _result *IStorageFile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SourceFile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUploadOperation) Get_Progress() BackgroundUploadProgress {
	var _result BackgroundUploadProgress
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Progress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUploadOperation) StartAsync() *IAsyncOperationWithProgress[*IUploadOperation, *IUploadOperation] {
	var _result *IAsyncOperationWithProgress[*IUploadOperation, *IUploadOperation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUploadOperation) AttachAsync() *IAsyncOperationWithProgress[*IUploadOperation, *IUploadOperation] {
	var _result *IAsyncOperationWithProgress[*IUploadOperation, *IUploadOperation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AttachAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 556189F2-2774-4DF6-9FA5-209F2BFB12F7
var IID_IUploadOperation2 = syscall.GUID{0x556189F2, 0x2774, 0x4DF6,
	[8]byte{0x9F, 0xA5, 0x20, 0x9F, 0x2B, 0xFB, 0x12, 0xF7}}

type IUploadOperation2Interface interface {
	win32.IInspectableInterface
	Get_TransferGroup() *IBackgroundTransferGroup
}

type IUploadOperation2Vtbl struct {
	win32.IInspectableVtbl
	Get_TransferGroup uintptr
}

type IUploadOperation2 struct {
	win32.IInspectable
}

func (this *IUploadOperation2) Vtbl() *IUploadOperation2Vtbl {
	return (*IUploadOperation2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUploadOperation2) Get_TransferGroup() *IBackgroundTransferGroup {
	var _result *IBackgroundTransferGroup
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TransferGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 42C92CA3-DE39-4546-BC62-3774B4294DE3
var IID_IUploadOperation3 = syscall.GUID{0x42C92CA3, 0xDE39, 0x4546,
	[8]byte{0xBC, 0x62, 0x37, 0x74, 0xB4, 0x29, 0x4D, 0xE3}}

type IUploadOperation3Interface interface {
	win32.IInspectableInterface
	MakeCurrentInTransferGroup()
}

type IUploadOperation3Vtbl struct {
	win32.IInspectableVtbl
	MakeCurrentInTransferGroup uintptr
}

type IUploadOperation3 struct {
	win32.IInspectable
}

func (this *IUploadOperation3) Vtbl() *IUploadOperation3Vtbl {
	return (*IUploadOperation3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUploadOperation3) MakeCurrentInTransferGroup() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MakeCurrentInTransferGroup, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 50EDEF31-FAC5-41EE-B030-DC77CAEE9FAA
var IID_IUploadOperation4 = syscall.GUID{0x50EDEF31, 0xFAC5, 0x41EE,
	[8]byte{0xB0, 0x30, 0xDC, 0x77, 0xCA, 0xEE, 0x9F, 0xAA}}

type IUploadOperation4Interface interface {
	win32.IInspectableInterface
	SetRequestHeader(headerName string, headerValue string)
	RemoveRequestHeader(headerName string)
}

type IUploadOperation4Vtbl struct {
	win32.IInspectableVtbl
	SetRequestHeader    uintptr
	RemoveRequestHeader uintptr
}

type IUploadOperation4 struct {
	win32.IInspectable
}

func (this *IUploadOperation4) Vtbl() *IUploadOperation4Vtbl {
	return (*IUploadOperation4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUploadOperation4) SetRequestHeader(headerName string, headerValue string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetRequestHeader, uintptr(unsafe.Pointer(this)), NewHStr(headerName).Ptr, NewHStr(headerValue).Ptr)
	_ = _hr
}

func (this *IUploadOperation4) RemoveRequestHeader(headerName string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveRequestHeader, uintptr(unsafe.Pointer(this)), NewHStr(headerName).Ptr)
	_ = _hr
}

// classes

type BackgroundTransferRangesDownloadedEventArgs struct {
	RtClass
	*IBackgroundTransferRangesDownloadedEventArgs
}
