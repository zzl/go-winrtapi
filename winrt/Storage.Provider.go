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
// flags
type CachedFileOptions uint32

const (
	CachedFileOptions_None                     CachedFileOptions = 0
	CachedFileOptions_RequireUpdateOnAccess    CachedFileOptions = 1
	CachedFileOptions_UseCachedFileWhenOffline CachedFileOptions = 2
	CachedFileOptions_DenyAccessWhenOffline    CachedFileOptions = 4
)

// enum
type CachedFileTarget int32

const (
	CachedFileTarget_Local  CachedFileTarget = 0
	CachedFileTarget_Remote CachedFileTarget = 1
)

// enum
type FileUpdateStatus int32

const (
	FileUpdateStatus_Incomplete           FileUpdateStatus = 0
	FileUpdateStatus_Complete             FileUpdateStatus = 1
	FileUpdateStatus_UserInputNeeded      FileUpdateStatus = 2
	FileUpdateStatus_CurrentlyUnavailable FileUpdateStatus = 3
	FileUpdateStatus_Failed               FileUpdateStatus = 4
	FileUpdateStatus_CompleteAndRenamed   FileUpdateStatus = 5
)

// enum
type ReadActivationMode int32

const (
	ReadActivationMode_NotNeeded    ReadActivationMode = 0
	ReadActivationMode_BeforeAccess ReadActivationMode = 1
)

// enum
// flags
type StorageProviderHardlinkPolicy uint32

const (
	StorageProviderHardlinkPolicy_None    StorageProviderHardlinkPolicy = 0
	StorageProviderHardlinkPolicy_Allowed StorageProviderHardlinkPolicy = 1
)

// enum
type StorageProviderHydrationPolicy int32

const (
	StorageProviderHydrationPolicy_Partial     StorageProviderHydrationPolicy = 0
	StorageProviderHydrationPolicy_Progressive StorageProviderHydrationPolicy = 1
	StorageProviderHydrationPolicy_Full        StorageProviderHydrationPolicy = 2
	StorageProviderHydrationPolicy_AlwaysFull  StorageProviderHydrationPolicy = 3
)

// enum
// flags
type StorageProviderHydrationPolicyModifier uint32

const (
	StorageProviderHydrationPolicyModifier_None                      StorageProviderHydrationPolicyModifier = 0
	StorageProviderHydrationPolicyModifier_ValidationRequired        StorageProviderHydrationPolicyModifier = 1
	StorageProviderHydrationPolicyModifier_StreamingAllowed          StorageProviderHydrationPolicyModifier = 2
	StorageProviderHydrationPolicyModifier_AutoDehydrationAllowed    StorageProviderHydrationPolicyModifier = 4
	StorageProviderHydrationPolicyModifier_AllowFullRestartHydration StorageProviderHydrationPolicyModifier = 8
)

// enum
// flags
type StorageProviderInSyncPolicy uint32

const (
	StorageProviderInSyncPolicy_Default                     StorageProviderInSyncPolicy = 0
	StorageProviderInSyncPolicy_FileCreationTime            StorageProviderInSyncPolicy = 1
	StorageProviderInSyncPolicy_FileReadOnlyAttribute       StorageProviderInSyncPolicy = 2
	StorageProviderInSyncPolicy_FileHiddenAttribute         StorageProviderInSyncPolicy = 4
	StorageProviderInSyncPolicy_FileSystemAttribute         StorageProviderInSyncPolicy = 8
	StorageProviderInSyncPolicy_DirectoryCreationTime       StorageProviderInSyncPolicy = 16
	StorageProviderInSyncPolicy_DirectoryReadOnlyAttribute  StorageProviderInSyncPolicy = 32
	StorageProviderInSyncPolicy_DirectoryHiddenAttribute    StorageProviderInSyncPolicy = 64
	StorageProviderInSyncPolicy_DirectorySystemAttribute    StorageProviderInSyncPolicy = 128
	StorageProviderInSyncPolicy_FileLastWriteTime           StorageProviderInSyncPolicy = 256
	StorageProviderInSyncPolicy_DirectoryLastWriteTime      StorageProviderInSyncPolicy = 512
	StorageProviderInSyncPolicy_PreserveInsyncForSyncEngine StorageProviderInSyncPolicy = 2147483648
)

// enum
type StorageProviderPopulationPolicy int32

const (
	StorageProviderPopulationPolicy_Full       StorageProviderPopulationPolicy = 1
	StorageProviderPopulationPolicy_AlwaysFull StorageProviderPopulationPolicy = 2
)

// enum
type StorageProviderProtectionMode int32

const (
	StorageProviderProtectionMode_Unknown  StorageProviderProtectionMode = 0
	StorageProviderProtectionMode_Personal StorageProviderProtectionMode = 1
)

// enum
type StorageProviderState int32

const (
	StorageProviderState_InSync  StorageProviderState = 0
	StorageProviderState_Syncing StorageProviderState = 1
	StorageProviderState_Paused  StorageProviderState = 2
	StorageProviderState_Error   StorageProviderState = 3
	StorageProviderState_Warning StorageProviderState = 4
	StorageProviderState_Offline StorageProviderState = 5
)

// enum
type StorageProviderUICommandState int32

const (
	StorageProviderUICommandState_Enabled  StorageProviderUICommandState = 0
	StorageProviderUICommandState_Disabled StorageProviderUICommandState = 1
	StorageProviderUICommandState_Hidden   StorageProviderUICommandState = 2
)

// enum
type StorageProviderUriSourceStatus int32

const (
	StorageProviderUriSourceStatus_Success      StorageProviderUriSourceStatus = 0
	StorageProviderUriSourceStatus_NoSyncRoot   StorageProviderUriSourceStatus = 1
	StorageProviderUriSourceStatus_FileNotFound StorageProviderUriSourceStatus = 2
)

// enum
type UIStatus int32

const (
	UIStatus_Unavailable UIStatus = 0
	UIStatus_Hidden      UIStatus = 1
	UIStatus_Visible     UIStatus = 2
	UIStatus_Complete    UIStatus = 3
)

// enum
type WriteActivationMode int32

const (
	WriteActivationMode_ReadOnly   WriteActivationMode = 0
	WriteActivationMode_NotNeeded  WriteActivationMode = 1
	WriteActivationMode_AfterWrite WriteActivationMode = 2
)

// structs

type CloudFilesContract struct {
}

// interfaces

// 9FC90920-7BCF-4888-A81E-102D7034D7CE
var IID_ICachedFileUpdaterStatics = syscall.GUID{0x9FC90920, 0x7BCF, 0x4888,
	[8]byte{0xA8, 0x1E, 0x10, 0x2D, 0x70, 0x34, 0xD7, 0xCE}}

type ICachedFileUpdaterStaticsInterface interface {
	win32.IInspectableInterface
	SetUpdateInformation(file *IStorageFile, contentId string, readMode ReadActivationMode, writeMode WriteActivationMode, options CachedFileOptions)
}

type ICachedFileUpdaterStaticsVtbl struct {
	win32.IInspectableVtbl
	SetUpdateInformation uintptr
}

type ICachedFileUpdaterStatics struct {
	win32.IInspectable
}

func (this *ICachedFileUpdaterStatics) Vtbl() *ICachedFileUpdaterStaticsVtbl {
	return (*ICachedFileUpdaterStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICachedFileUpdaterStatics) SetUpdateInformation(file *IStorageFile, contentId string, readMode ReadActivationMode, writeMode WriteActivationMode, options CachedFileOptions) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetUpdateInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), NewHStr(contentId).Ptr, uintptr(readMode), uintptr(writeMode), uintptr(options))
	_ = _hr
}

// 9E6F41E6-BAF2-4A97-B600-9333F5DF80FD
var IID_ICachedFileUpdaterUI = syscall.GUID{0x9E6F41E6, 0xBAF2, 0x4A97,
	[8]byte{0xB6, 0x00, 0x93, 0x33, 0xF5, 0xDF, 0x80, 0xFD}}

type ICachedFileUpdaterUIInterface interface {
	win32.IInspectableInterface
	Get_Title() string
	Put_Title(value string)
	Get_UpdateTarget() CachedFileTarget
	Add_FileUpdateRequested(handler TypedEventHandler[*ICachedFileUpdaterUI, *IFileUpdateRequestedEventArgs]) EventRegistrationToken
	Remove_FileUpdateRequested(token EventRegistrationToken)
	Add_UIRequested(handler TypedEventHandler[*ICachedFileUpdaterUI, interface{}]) EventRegistrationToken
	Remove_UIRequested(token EventRegistrationToken)
	Get_UIStatus() UIStatus
}

type ICachedFileUpdaterUIVtbl struct {
	win32.IInspectableVtbl
	Get_Title                  uintptr
	Put_Title                  uintptr
	Get_UpdateTarget           uintptr
	Add_FileUpdateRequested    uintptr
	Remove_FileUpdateRequested uintptr
	Add_UIRequested            uintptr
	Remove_UIRequested         uintptr
	Get_UIStatus               uintptr
}

type ICachedFileUpdaterUI struct {
	win32.IInspectable
}

func (this *ICachedFileUpdaterUI) Vtbl() *ICachedFileUpdaterUIVtbl {
	return (*ICachedFileUpdaterUIVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICachedFileUpdaterUI) Get_Title() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Title, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICachedFileUpdaterUI) Put_Title(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Title, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICachedFileUpdaterUI) Get_UpdateTarget() CachedFileTarget {
	var _result CachedFileTarget
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UpdateTarget, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICachedFileUpdaterUI) Add_FileUpdateRequested(handler TypedEventHandler[*ICachedFileUpdaterUI, *IFileUpdateRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_FileUpdateRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICachedFileUpdaterUI) Remove_FileUpdateRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_FileUpdateRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ICachedFileUpdaterUI) Add_UIRequested(handler TypedEventHandler[*ICachedFileUpdaterUI, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_UIRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICachedFileUpdaterUI) Remove_UIRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_UIRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ICachedFileUpdaterUI) Get_UIStatus() UIStatus {
	var _result UIStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UIStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 8856A21C-8699-4340-9F49-F7CAD7FE8991
var IID_ICachedFileUpdaterUI2 = syscall.GUID{0x8856A21C, 0x8699, 0x4340,
	[8]byte{0x9F, 0x49, 0xF7, 0xCA, 0xD7, 0xFE, 0x89, 0x91}}

type ICachedFileUpdaterUI2Interface interface {
	win32.IInspectableInterface
	Get_UpdateRequest() *IFileUpdateRequest
	GetDeferral() *IFileUpdateRequestDeferral
}

type ICachedFileUpdaterUI2Vtbl struct {
	win32.IInspectableVtbl
	Get_UpdateRequest uintptr
	GetDeferral       uintptr
}

type ICachedFileUpdaterUI2 struct {
	win32.IInspectable
}

func (this *ICachedFileUpdaterUI2) Vtbl() *ICachedFileUpdaterUI2Vtbl {
	return (*ICachedFileUpdaterUI2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICachedFileUpdaterUI2) Get_UpdateRequest() *IFileUpdateRequest {
	var _result *IFileUpdateRequest
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UpdateRequest, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICachedFileUpdaterUI2) GetDeferral() *IFileUpdateRequestDeferral {
	var _result *IFileUpdateRequestDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 40C82536-C1FE-4D93-A792-1E736BC70837
var IID_IFileUpdateRequest = syscall.GUID{0x40C82536, 0xC1FE, 0x4D93,
	[8]byte{0xA7, 0x92, 0x1E, 0x73, 0x6B, 0xC7, 0x08, 0x37}}

type IFileUpdateRequestInterface interface {
	win32.IInspectableInterface
	Get_ContentId() string
	Get_File() *IStorageFile
	Get_Status() FileUpdateStatus
	Put_Status(value FileUpdateStatus)
	GetDeferral() *IFileUpdateRequestDeferral
	UpdateLocalFile(value *IStorageFile)
}

type IFileUpdateRequestVtbl struct {
	win32.IInspectableVtbl
	Get_ContentId   uintptr
	Get_File        uintptr
	Get_Status      uintptr
	Put_Status      uintptr
	GetDeferral     uintptr
	UpdateLocalFile uintptr
}

type IFileUpdateRequest struct {
	win32.IInspectable
}

func (this *IFileUpdateRequest) Vtbl() *IFileUpdateRequestVtbl {
	return (*IFileUpdateRequestVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFileUpdateRequest) Get_ContentId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IFileUpdateRequest) Get_File() *IStorageFile {
	var _result *IStorageFile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_File, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileUpdateRequest) Get_Status() FileUpdateStatus {
	var _result FileUpdateStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFileUpdateRequest) Put_Status(value FileUpdateStatus) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Status, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IFileUpdateRequest) GetDeferral() *IFileUpdateRequestDeferral {
	var _result *IFileUpdateRequestDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileUpdateRequest) UpdateLocalFile(value *IStorageFile) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UpdateLocalFile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 82484648-BDBE-447B-A2EE-7AFE6A032A94
var IID_IFileUpdateRequest2 = syscall.GUID{0x82484648, 0xBDBE, 0x447B,
	[8]byte{0xA2, 0xEE, 0x7A, 0xFE, 0x6A, 0x03, 0x2A, 0x94}}

type IFileUpdateRequest2Interface interface {
	win32.IInspectableInterface
	Get_UserInputNeededMessage() string
	Put_UserInputNeededMessage(value string)
}

type IFileUpdateRequest2Vtbl struct {
	win32.IInspectableVtbl
	Get_UserInputNeededMessage uintptr
	Put_UserInputNeededMessage uintptr
}

type IFileUpdateRequest2 struct {
	win32.IInspectable
}

func (this *IFileUpdateRequest2) Vtbl() *IFileUpdateRequest2Vtbl {
	return (*IFileUpdateRequest2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFileUpdateRequest2) Get_UserInputNeededMessage() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UserInputNeededMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IFileUpdateRequest2) Put_UserInputNeededMessage(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_UserInputNeededMessage, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// FFCEDB2B-8ADE-44A5-BB00-164C4E72F13A
var IID_IFileUpdateRequestDeferral = syscall.GUID{0xFFCEDB2B, 0x8ADE, 0x44A5,
	[8]byte{0xBB, 0x00, 0x16, 0x4C, 0x4E, 0x72, 0xF1, 0x3A}}

type IFileUpdateRequestDeferralInterface interface {
	win32.IInspectableInterface
	Complete()
}

type IFileUpdateRequestDeferralVtbl struct {
	win32.IInspectableVtbl
	Complete uintptr
}

type IFileUpdateRequestDeferral struct {
	win32.IInspectable
}

func (this *IFileUpdateRequestDeferral) Vtbl() *IFileUpdateRequestDeferralVtbl {
	return (*IFileUpdateRequestDeferralVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFileUpdateRequestDeferral) Complete() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Complete, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 7B0A9342-3905-438D-AAEF-78AE265F8DD2
var IID_IFileUpdateRequestedEventArgs = syscall.GUID{0x7B0A9342, 0x3905, 0x438D,
	[8]byte{0xAA, 0xEF, 0x78, 0xAE, 0x26, 0x5F, 0x8D, 0xD2}}

type IFileUpdateRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Request() *IFileUpdateRequest
}

type IFileUpdateRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Request uintptr
}

type IFileUpdateRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IFileUpdateRequestedEventArgs) Vtbl() *IFileUpdateRequestedEventArgsVtbl {
	return (*IFileUpdateRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFileUpdateRequestedEventArgs) Get_Request() *IFileUpdateRequest {
	var _result *IFileUpdateRequest
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Request, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1955B9C1-0184-5A88-87DF-4544F464365D
var IID_IStorageProviderFileTypeInfo = syscall.GUID{0x1955B9C1, 0x0184, 0x5A88,
	[8]byte{0x87, 0xDF, 0x45, 0x44, 0xF4, 0x64, 0x36, 0x5D}}

type IStorageProviderFileTypeInfoInterface interface {
	win32.IInspectableInterface
	Get_FileExtension() string
	Get_IconResource() string
}

type IStorageProviderFileTypeInfoVtbl struct {
	win32.IInspectableVtbl
	Get_FileExtension uintptr
	Get_IconResource  uintptr
}

type IStorageProviderFileTypeInfo struct {
	win32.IInspectable
}

func (this *IStorageProviderFileTypeInfo) Vtbl() *IStorageProviderFileTypeInfoVtbl {
	return (*IStorageProviderFileTypeInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageProviderFileTypeInfo) Get_FileExtension() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FileExtension, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStorageProviderFileTypeInfo) Get_IconResource() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IconResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 3FA12C6F-CCE6-5D5D-80B1-389E7CF92DBF
var IID_IStorageProviderFileTypeInfoFactory = syscall.GUID{0x3FA12C6F, 0xCCE6, 0x5D5D,
	[8]byte{0x80, 0xB1, 0x38, 0x9E, 0x7C, 0xF9, 0x2D, 0xBF}}

type IStorageProviderFileTypeInfoFactoryInterface interface {
	win32.IInspectableInterface
	CreateInstance(fileExtension string, iconResource string) *IStorageProviderFileTypeInfo
}

type IStorageProviderFileTypeInfoFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateInstance uintptr
}

type IStorageProviderFileTypeInfoFactory struct {
	win32.IInspectable
}

func (this *IStorageProviderFileTypeInfoFactory) Vtbl() *IStorageProviderFileTypeInfoFactoryVtbl {
	return (*IStorageProviderFileTypeInfoFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageProviderFileTypeInfoFactory) CreateInstance(fileExtension string, iconResource string) *IStorageProviderFileTypeInfo {
	var _result *IStorageProviderFileTypeInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateInstance, uintptr(unsafe.Pointer(this)), NewHStr(fileExtension).Ptr, NewHStr(iconResource).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2564711D-AA89-4D12-82E3-F72A92E33966
var IID_IStorageProviderGetContentInfoForPathResult = syscall.GUID{0x2564711D, 0xAA89, 0x4D12,
	[8]byte{0x82, 0xE3, 0xF7, 0x2A, 0x92, 0xE3, 0x39, 0x66}}

type IStorageProviderGetContentInfoForPathResultInterface interface {
	win32.IInspectableInterface
	Get_Status() StorageProviderUriSourceStatus
	Put_Status(value StorageProviderUriSourceStatus)
	Get_ContentUri() string
	Put_ContentUri(value string)
	Get_ContentId() string
	Put_ContentId(value string)
}

type IStorageProviderGetContentInfoForPathResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status     uintptr
	Put_Status     uintptr
	Get_ContentUri uintptr
	Put_ContentUri uintptr
	Get_ContentId  uintptr
	Put_ContentId  uintptr
}

type IStorageProviderGetContentInfoForPathResult struct {
	win32.IInspectable
}

func (this *IStorageProviderGetContentInfoForPathResult) Vtbl() *IStorageProviderGetContentInfoForPathResultVtbl {
	return (*IStorageProviderGetContentInfoForPathResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageProviderGetContentInfoForPathResult) Get_Status() StorageProviderUriSourceStatus {
	var _result StorageProviderUriSourceStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStorageProviderGetContentInfoForPathResult) Put_Status(value StorageProviderUriSourceStatus) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Status, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IStorageProviderGetContentInfoForPathResult) Get_ContentUri() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStorageProviderGetContentInfoForPathResult) Put_ContentUri(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ContentUri, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IStorageProviderGetContentInfoForPathResult) Get_ContentId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStorageProviderGetContentInfoForPathResult) Put_ContentId(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ContentId, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// 63711A9D-4118-45A6-ACB6-22C49D019F40
var IID_IStorageProviderGetPathForContentUriResult = syscall.GUID{0x63711A9D, 0x4118, 0x45A6,
	[8]byte{0xAC, 0xB6, 0x22, 0xC4, 0x9D, 0x01, 0x9F, 0x40}}

type IStorageProviderGetPathForContentUriResultInterface interface {
	win32.IInspectableInterface
	Get_Status() StorageProviderUriSourceStatus
	Put_Status(value StorageProviderUriSourceStatus)
	Get_Path() string
	Put_Path(value string)
}

type IStorageProviderGetPathForContentUriResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
	Put_Status uintptr
	Get_Path   uintptr
	Put_Path   uintptr
}

type IStorageProviderGetPathForContentUriResult struct {
	win32.IInspectable
}

func (this *IStorageProviderGetPathForContentUriResult) Vtbl() *IStorageProviderGetPathForContentUriResultVtbl {
	return (*IStorageProviderGetPathForContentUriResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageProviderGetPathForContentUriResult) Get_Status() StorageProviderUriSourceStatus {
	var _result StorageProviderUriSourceStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStorageProviderGetPathForContentUriResult) Put_Status(value StorageProviderUriSourceStatus) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Status, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IStorageProviderGetPathForContentUriResult) Get_Path() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Path, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStorageProviderGetPathForContentUriResult) Put_Path(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Path, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// 2D2C1C97-2704-4729-8FA9-7E6B8E158C2F
var IID_IStorageProviderItemPropertiesStatics = syscall.GUID{0x2D2C1C97, 0x2704, 0x4729,
	[8]byte{0x8F, 0xA9, 0x7E, 0x6B, 0x8E, 0x15, 0x8C, 0x2F}}

type IStorageProviderItemPropertiesStaticsInterface interface {
	win32.IInspectableInterface
	SetAsync(item *IStorageItem, itemProperties *IIterable[*IStorageProviderItemProperty]) *IAsyncAction
}

type IStorageProviderItemPropertiesStaticsVtbl struct {
	win32.IInspectableVtbl
	SetAsync uintptr
}

type IStorageProviderItemPropertiesStatics struct {
	win32.IInspectable
}

func (this *IStorageProviderItemPropertiesStatics) Vtbl() *IStorageProviderItemPropertiesStaticsVtbl {
	return (*IStorageProviderItemPropertiesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageProviderItemPropertiesStatics) SetAsync(item *IStorageItem, itemProperties *IIterable[*IStorageProviderItemProperty]) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(item)), uintptr(unsafe.Pointer(itemProperties)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 476CB558-730B-4188-B7B5-63B716ED476D
var IID_IStorageProviderItemProperty = syscall.GUID{0x476CB558, 0x730B, 0x4188,
	[8]byte{0xB7, 0xB5, 0x63, 0xB7, 0x16, 0xED, 0x47, 0x6D}}

type IStorageProviderItemPropertyInterface interface {
	win32.IInspectableInterface
	Put_Id(value int32)
	Get_Id() int32
	Put_Value(value string)
	Get_Value() string
	Put_IconResource(value string)
	Get_IconResource() string
}

type IStorageProviderItemPropertyVtbl struct {
	win32.IInspectableVtbl
	Put_Id           uintptr
	Get_Id           uintptr
	Put_Value        uintptr
	Get_Value        uintptr
	Put_IconResource uintptr
	Get_IconResource uintptr
}

type IStorageProviderItemProperty struct {
	win32.IInspectable
}

func (this *IStorageProviderItemProperty) Vtbl() *IStorageProviderItemPropertyVtbl {
	return (*IStorageProviderItemPropertyVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageProviderItemProperty) Put_Id(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Id, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IStorageProviderItemProperty) Get_Id() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStorageProviderItemProperty) Put_Value(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Value, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IStorageProviderItemProperty) Get_Value() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStorageProviderItemProperty) Put_IconResource(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IconResource, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IStorageProviderItemProperty) Get_IconResource() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IconResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// C5B383BB-FF1F-4298-831E-FF1C08089690
var IID_IStorageProviderItemPropertyDefinition = syscall.GUID{0xC5B383BB, 0xFF1F, 0x4298,
	[8]byte{0x83, 0x1E, 0xFF, 0x1C, 0x08, 0x08, 0x96, 0x90}}

type IStorageProviderItemPropertyDefinitionInterface interface {
	win32.IInspectableInterface
	Get_Id() int32
	Put_Id(value int32)
	Get_DisplayNameResource() string
	Put_DisplayNameResource(value string)
}

type IStorageProviderItemPropertyDefinitionVtbl struct {
	win32.IInspectableVtbl
	Get_Id                  uintptr
	Put_Id                  uintptr
	Get_DisplayNameResource uintptr
	Put_DisplayNameResource uintptr
}

type IStorageProviderItemPropertyDefinition struct {
	win32.IInspectable
}

func (this *IStorageProviderItemPropertyDefinition) Vtbl() *IStorageProviderItemPropertyDefinitionVtbl {
	return (*IStorageProviderItemPropertyDefinitionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageProviderItemPropertyDefinition) Get_Id() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStorageProviderItemPropertyDefinition) Put_Id(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Id, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IStorageProviderItemPropertyDefinition) Get_DisplayNameResource() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayNameResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStorageProviderItemPropertyDefinition) Put_DisplayNameResource(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DisplayNameResource, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// 8F6F9C3E-F632-4A9B-8D99-D2D7A11DF56A
var IID_IStorageProviderItemPropertySource = syscall.GUID{0x8F6F9C3E, 0xF632, 0x4A9B,
	[8]byte{0x8D, 0x99, 0xD2, 0xD7, 0xA1, 0x1D, 0xF5, 0x6A}}

type IStorageProviderItemPropertySourceInterface interface {
	win32.IInspectableInterface
	GetItemProperties(itemPath string) *IIterable[*IStorageProviderItemProperty]
}

type IStorageProviderItemPropertySourceVtbl struct {
	win32.IInspectableVtbl
	GetItemProperties uintptr
}

type IStorageProviderItemPropertySource struct {
	win32.IInspectable
}

func (this *IStorageProviderItemPropertySource) Vtbl() *IStorageProviderItemPropertySourceVtbl {
	return (*IStorageProviderItemPropertySourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageProviderItemPropertySource) GetItemProperties(itemPath string) *IIterable[*IStorageProviderItemProperty] {
	var _result *IIterable[*IStorageProviderItemProperty]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetItemProperties, uintptr(unsafe.Pointer(this)), NewHStr(itemPath).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EF38E591-A7CB-5E7D-9B5E-22749842697C
var IID_IStorageProviderMoreInfoUI = syscall.GUID{0xEF38E591, 0xA7CB, 0x5E7D,
	[8]byte{0x9B, 0x5E, 0x22, 0x74, 0x98, 0x42, 0x69, 0x7C}}

type IStorageProviderMoreInfoUIInterface interface {
	win32.IInspectableInterface
	Get_Message() string
	Put_Message(value string)
	Get_Command() *IStorageProviderUICommand
	Put_Command(value *IStorageProviderUICommand)
}

type IStorageProviderMoreInfoUIVtbl struct {
	win32.IInspectableVtbl
	Get_Message uintptr
	Put_Message uintptr
	Get_Command uintptr
	Put_Command uintptr
}

type IStorageProviderMoreInfoUI struct {
	win32.IInspectable
}

func (this *IStorageProviderMoreInfoUI) Vtbl() *IStorageProviderMoreInfoUIVtbl {
	return (*IStorageProviderMoreInfoUIVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageProviderMoreInfoUI) Get_Message() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Message, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStorageProviderMoreInfoUI) Put_Message(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Message, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IStorageProviderMoreInfoUI) Get_Command() *IStorageProviderUICommand {
	var _result *IStorageProviderUICommand
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Command, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageProviderMoreInfoUI) Put_Command(value *IStorageProviderUICommand) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Command, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 658D2F0E-63B7-4567-ACF9-51ABE301DDA5
var IID_IStorageProviderPropertyCapabilities = syscall.GUID{0x658D2F0E, 0x63B7, 0x4567,
	[8]byte{0xAC, 0xF9, 0x51, 0xAB, 0xE3, 0x01, 0xDD, 0xA5}}

type IStorageProviderPropertyCapabilitiesInterface interface {
	win32.IInspectableInterface
	IsPropertySupported(propertyCanonicalName string) bool
}

type IStorageProviderPropertyCapabilitiesVtbl struct {
	win32.IInspectableVtbl
	IsPropertySupported uintptr
}

type IStorageProviderPropertyCapabilities struct {
	win32.IInspectable
}

func (this *IStorageProviderPropertyCapabilities) Vtbl() *IStorageProviderPropertyCapabilitiesVtbl {
	return (*IStorageProviderPropertyCapabilitiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageProviderPropertyCapabilities) IsPropertySupported(propertyCanonicalName string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsPropertySupported, uintptr(unsafe.Pointer(this)), NewHStr(propertyCanonicalName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// BA6295C3-312E-544F-9FD5-1F81B21F3649
var IID_IStorageProviderQuotaUI = syscall.GUID{0xBA6295C3, 0x312E, 0x544F,
	[8]byte{0x9F, 0xD5, 0x1F, 0x81, 0xB2, 0x1F, 0x36, 0x49}}

type IStorageProviderQuotaUIInterface interface {
	win32.IInspectableInterface
	Get_QuotaTotalInBytes() uint64
	Put_QuotaTotalInBytes(value uint64)
	Get_QuotaUsedInBytes() uint64
	Put_QuotaUsedInBytes(value uint64)
	Get_QuotaUsedLabel() string
	Put_QuotaUsedLabel(value string)
	Get_QuotaUsedColor() *IReference[unsafe.Pointer]
	Put_QuotaUsedColor(value *IReference[unsafe.Pointer])
}

type IStorageProviderQuotaUIVtbl struct {
	win32.IInspectableVtbl
	Get_QuotaTotalInBytes uintptr
	Put_QuotaTotalInBytes uintptr
	Get_QuotaUsedInBytes  uintptr
	Put_QuotaUsedInBytes  uintptr
	Get_QuotaUsedLabel    uintptr
	Put_QuotaUsedLabel    uintptr
	Get_QuotaUsedColor    uintptr
	Put_QuotaUsedColor    uintptr
}

type IStorageProviderQuotaUI struct {
	win32.IInspectable
}

func (this *IStorageProviderQuotaUI) Vtbl() *IStorageProviderQuotaUIVtbl {
	return (*IStorageProviderQuotaUIVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageProviderQuotaUI) Get_QuotaTotalInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_QuotaTotalInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStorageProviderQuotaUI) Put_QuotaTotalInBytes(value uint64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_QuotaTotalInBytes, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IStorageProviderQuotaUI) Get_QuotaUsedInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_QuotaUsedInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStorageProviderQuotaUI) Put_QuotaUsedInBytes(value uint64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_QuotaUsedInBytes, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IStorageProviderQuotaUI) Get_QuotaUsedLabel() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_QuotaUsedLabel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStorageProviderQuotaUI) Put_QuotaUsedLabel(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_QuotaUsedLabel, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IStorageProviderQuotaUI) Get_QuotaUsedColor() *IReference[unsafe.Pointer] {
	var _result *IReference[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_QuotaUsedColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageProviderQuotaUI) Put_QuotaUsedColor(value *IReference[unsafe.Pointer]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_QuotaUsedColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// D6B6A758-198D-5B80-977F-5FF73DA33118
var IID_IStorageProviderStatusUI = syscall.GUID{0xD6B6A758, 0x198D, 0x5B80,
	[8]byte{0x97, 0x7F, 0x5F, 0xF7, 0x3D, 0xA3, 0x31, 0x18}}

type IStorageProviderStatusUIInterface interface {
	win32.IInspectableInterface
	Get_ProviderState() StorageProviderState
	Put_ProviderState(value StorageProviderState)
	Get_ProviderStateLabel() string
	Put_ProviderStateLabel(value string)
	Get_ProviderStateIcon() *IUriRuntimeClass
	Put_ProviderStateIcon(value *IUriRuntimeClass)
	Get_SyncStatusCommand() *IStorageProviderUICommand
	Put_SyncStatusCommand(value *IStorageProviderUICommand)
	Get_QuotaUI() *IStorageProviderQuotaUI
	Put_QuotaUI(value *IStorageProviderQuotaUI)
	Get_MoreInfoUI() *IStorageProviderMoreInfoUI
	Put_MoreInfoUI(value *IStorageProviderMoreInfoUI)
	Get_ProviderPrimaryCommand() *IStorageProviderUICommand
	Put_ProviderPrimaryCommand(value *IStorageProviderUICommand)
	Get_ProviderSecondaryCommands() *IVector[*IStorageProviderUICommand]
	Put_ProviderSecondaryCommands(value *IVector[*IStorageProviderUICommand])
}

type IStorageProviderStatusUIVtbl struct {
	win32.IInspectableVtbl
	Get_ProviderState             uintptr
	Put_ProviderState             uintptr
	Get_ProviderStateLabel        uintptr
	Put_ProviderStateLabel        uintptr
	Get_ProviderStateIcon         uintptr
	Put_ProviderStateIcon         uintptr
	Get_SyncStatusCommand         uintptr
	Put_SyncStatusCommand         uintptr
	Get_QuotaUI                   uintptr
	Put_QuotaUI                   uintptr
	Get_MoreInfoUI                uintptr
	Put_MoreInfoUI                uintptr
	Get_ProviderPrimaryCommand    uintptr
	Put_ProviderPrimaryCommand    uintptr
	Get_ProviderSecondaryCommands uintptr
	Put_ProviderSecondaryCommands uintptr
}

type IStorageProviderStatusUI struct {
	win32.IInspectable
}

func (this *IStorageProviderStatusUI) Vtbl() *IStorageProviderStatusUIVtbl {
	return (*IStorageProviderStatusUIVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageProviderStatusUI) Get_ProviderState() StorageProviderState {
	var _result StorageProviderState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProviderState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStorageProviderStatusUI) Put_ProviderState(value StorageProviderState) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ProviderState, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IStorageProviderStatusUI) Get_ProviderStateLabel() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProviderStateLabel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStorageProviderStatusUI) Put_ProviderStateLabel(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ProviderStateLabel, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IStorageProviderStatusUI) Get_ProviderStateIcon() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProviderStateIcon, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageProviderStatusUI) Put_ProviderStateIcon(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ProviderStateIcon, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IStorageProviderStatusUI) Get_SyncStatusCommand() *IStorageProviderUICommand {
	var _result *IStorageProviderUICommand
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SyncStatusCommand, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageProviderStatusUI) Put_SyncStatusCommand(value *IStorageProviderUICommand) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SyncStatusCommand, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IStorageProviderStatusUI) Get_QuotaUI() *IStorageProviderQuotaUI {
	var _result *IStorageProviderQuotaUI
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_QuotaUI, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageProviderStatusUI) Put_QuotaUI(value *IStorageProviderQuotaUI) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_QuotaUI, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IStorageProviderStatusUI) Get_MoreInfoUI() *IStorageProviderMoreInfoUI {
	var _result *IStorageProviderMoreInfoUI
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MoreInfoUI, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageProviderStatusUI) Put_MoreInfoUI(value *IStorageProviderMoreInfoUI) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MoreInfoUI, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IStorageProviderStatusUI) Get_ProviderPrimaryCommand() *IStorageProviderUICommand {
	var _result *IStorageProviderUICommand
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProviderPrimaryCommand, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageProviderStatusUI) Put_ProviderPrimaryCommand(value *IStorageProviderUICommand) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ProviderPrimaryCommand, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IStorageProviderStatusUI) Get_ProviderSecondaryCommands() *IVector[*IStorageProviderUICommand] {
	var _result *IVector[*IStorageProviderUICommand]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProviderSecondaryCommands, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageProviderStatusUI) Put_ProviderSecondaryCommands(value *IVector[*IStorageProviderUICommand]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ProviderSecondaryCommands, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// A306C249-3D66-5E70-9007-E43DF96051FF
var IID_IStorageProviderStatusUISource = syscall.GUID{0xA306C249, 0x3D66, 0x5E70,
	[8]byte{0x90, 0x07, 0xE4, 0x3D, 0xF9, 0x60, 0x51, 0xFF}}

type IStorageProviderStatusUISourceInterface interface {
	win32.IInspectableInterface
	GetStatusUI() *IStorageProviderStatusUI
	Add_StatusUIChanged(handler TypedEventHandler[*IStorageProviderStatusUISource, interface{}]) EventRegistrationToken
	Remove_StatusUIChanged(token EventRegistrationToken)
}

type IStorageProviderStatusUISourceVtbl struct {
	win32.IInspectableVtbl
	GetStatusUI            uintptr
	Add_StatusUIChanged    uintptr
	Remove_StatusUIChanged uintptr
}

type IStorageProviderStatusUISource struct {
	win32.IInspectable
}

func (this *IStorageProviderStatusUISource) Vtbl() *IStorageProviderStatusUISourceVtbl {
	return (*IStorageProviderStatusUISourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageProviderStatusUISource) GetStatusUI() *IStorageProviderStatusUI {
	var _result *IStorageProviderStatusUI
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStatusUI, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageProviderStatusUISource) Add_StatusUIChanged(handler TypedEventHandler[*IStorageProviderStatusUISource, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StatusUIChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStorageProviderStatusUISource) Remove_StatusUIChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StatusUIChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 12E46B74-4E5A-58D1-A62F-0376E8EE7DD8
var IID_IStorageProviderStatusUISourceFactory = syscall.GUID{0x12E46B74, 0x4E5A, 0x58D1,
	[8]byte{0xA6, 0x2F, 0x03, 0x76, 0xE8, 0xEE, 0x7D, 0xD8}}

type IStorageProviderStatusUISourceFactoryInterface interface {
	win32.IInspectableInterface
	GetStatusUISource(syncRootId string) *IStorageProviderStatusUISource
}

type IStorageProviderStatusUISourceFactoryVtbl struct {
	win32.IInspectableVtbl
	GetStatusUISource uintptr
}

type IStorageProviderStatusUISourceFactory struct {
	win32.IInspectable
}

func (this *IStorageProviderStatusUISourceFactory) Vtbl() *IStorageProviderStatusUISourceFactoryVtbl {
	return (*IStorageProviderStatusUISourceFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageProviderStatusUISourceFactory) GetStatusUISource(syncRootId string) *IStorageProviderStatusUISource {
	var _result *IStorageProviderStatusUISource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStatusUISource, uintptr(unsafe.Pointer(this)), NewHStr(syncRootId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7C1305C4-99F9-41AC-8904-AB055D654926
var IID_IStorageProviderSyncRootInfo = syscall.GUID{0x7C1305C4, 0x99F9, 0x41AC,
	[8]byte{0x89, 0x04, 0xAB, 0x05, 0x5D, 0x65, 0x49, 0x26}}

type IStorageProviderSyncRootInfoInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	Put_Id(value string)
	Get_Context() *IBuffer
	Put_Context(value *IBuffer)
	Get_Path() *IStorageFolder
	Put_Path(value *IStorageFolder)
	Get_DisplayNameResource() string
	Put_DisplayNameResource(value string)
	Get_IconResource() string
	Put_IconResource(value string)
	Get_HydrationPolicy() StorageProviderHydrationPolicy
	Put_HydrationPolicy(value StorageProviderHydrationPolicy)
	Get_HydrationPolicyModifier() StorageProviderHydrationPolicyModifier
	Put_HydrationPolicyModifier(value StorageProviderHydrationPolicyModifier)
	Get_PopulationPolicy() StorageProviderPopulationPolicy
	Put_PopulationPolicy(value StorageProviderPopulationPolicy)
	Get_InSyncPolicy() StorageProviderInSyncPolicy
	Put_InSyncPolicy(value StorageProviderInSyncPolicy)
	Get_HardlinkPolicy() StorageProviderHardlinkPolicy
	Put_HardlinkPolicy(value StorageProviderHardlinkPolicy)
	Get_ShowSiblingsAsGroup() bool
	Put_ShowSiblingsAsGroup(value bool)
	Get_Version() string
	Put_Version(value string)
	Get_ProtectionMode() StorageProviderProtectionMode
	Put_ProtectionMode(value StorageProviderProtectionMode)
	Get_AllowPinning() bool
	Put_AllowPinning(value bool)
	Get_StorageProviderItemPropertyDefinitions() *IVector[*IStorageProviderItemPropertyDefinition]
	Get_RecycleBinUri() *IUriRuntimeClass
	Put_RecycleBinUri(value *IUriRuntimeClass)
}

type IStorageProviderSyncRootInfoVtbl struct {
	win32.IInspectableVtbl
	Get_Id                                     uintptr
	Put_Id                                     uintptr
	Get_Context                                uintptr
	Put_Context                                uintptr
	Get_Path                                   uintptr
	Put_Path                                   uintptr
	Get_DisplayNameResource                    uintptr
	Put_DisplayNameResource                    uintptr
	Get_IconResource                           uintptr
	Put_IconResource                           uintptr
	Get_HydrationPolicy                        uintptr
	Put_HydrationPolicy                        uintptr
	Get_HydrationPolicyModifier                uintptr
	Put_HydrationPolicyModifier                uintptr
	Get_PopulationPolicy                       uintptr
	Put_PopulationPolicy                       uintptr
	Get_InSyncPolicy                           uintptr
	Put_InSyncPolicy                           uintptr
	Get_HardlinkPolicy                         uintptr
	Put_HardlinkPolicy                         uintptr
	Get_ShowSiblingsAsGroup                    uintptr
	Put_ShowSiblingsAsGroup                    uintptr
	Get_Version                                uintptr
	Put_Version                                uintptr
	Get_ProtectionMode                         uintptr
	Put_ProtectionMode                         uintptr
	Get_AllowPinning                           uintptr
	Put_AllowPinning                           uintptr
	Get_StorageProviderItemPropertyDefinitions uintptr
	Get_RecycleBinUri                          uintptr
	Put_RecycleBinUri                          uintptr
}

type IStorageProviderSyncRootInfo struct {
	win32.IInspectable
}

func (this *IStorageProviderSyncRootInfo) Vtbl() *IStorageProviderSyncRootInfoVtbl {
	return (*IStorageProviderSyncRootInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageProviderSyncRootInfo) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStorageProviderSyncRootInfo) Put_Id(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Id, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IStorageProviderSyncRootInfo) Get_Context() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Context, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageProviderSyncRootInfo) Put_Context(value *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Context, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IStorageProviderSyncRootInfo) Get_Path() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Path, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageProviderSyncRootInfo) Put_Path(value *IStorageFolder) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Path, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IStorageProviderSyncRootInfo) Get_DisplayNameResource() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayNameResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStorageProviderSyncRootInfo) Put_DisplayNameResource(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DisplayNameResource, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IStorageProviderSyncRootInfo) Get_IconResource() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IconResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStorageProviderSyncRootInfo) Put_IconResource(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IconResource, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IStorageProviderSyncRootInfo) Get_HydrationPolicy() StorageProviderHydrationPolicy {
	var _result StorageProviderHydrationPolicy
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HydrationPolicy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStorageProviderSyncRootInfo) Put_HydrationPolicy(value StorageProviderHydrationPolicy) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_HydrationPolicy, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IStorageProviderSyncRootInfo) Get_HydrationPolicyModifier() StorageProviderHydrationPolicyModifier {
	var _result StorageProviderHydrationPolicyModifier
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HydrationPolicyModifier, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStorageProviderSyncRootInfo) Put_HydrationPolicyModifier(value StorageProviderHydrationPolicyModifier) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_HydrationPolicyModifier, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IStorageProviderSyncRootInfo) Get_PopulationPolicy() StorageProviderPopulationPolicy {
	var _result StorageProviderPopulationPolicy
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PopulationPolicy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStorageProviderSyncRootInfo) Put_PopulationPolicy(value StorageProviderPopulationPolicy) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PopulationPolicy, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IStorageProviderSyncRootInfo) Get_InSyncPolicy() StorageProviderInSyncPolicy {
	var _result StorageProviderInSyncPolicy
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InSyncPolicy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStorageProviderSyncRootInfo) Put_InSyncPolicy(value StorageProviderInSyncPolicy) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InSyncPolicy, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IStorageProviderSyncRootInfo) Get_HardlinkPolicy() StorageProviderHardlinkPolicy {
	var _result StorageProviderHardlinkPolicy
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HardlinkPolicy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStorageProviderSyncRootInfo) Put_HardlinkPolicy(value StorageProviderHardlinkPolicy) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_HardlinkPolicy, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IStorageProviderSyncRootInfo) Get_ShowSiblingsAsGroup() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ShowSiblingsAsGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStorageProviderSyncRootInfo) Put_ShowSiblingsAsGroup(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ShowSiblingsAsGroup, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IStorageProviderSyncRootInfo) Get_Version() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Version, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStorageProviderSyncRootInfo) Put_Version(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Version, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IStorageProviderSyncRootInfo) Get_ProtectionMode() StorageProviderProtectionMode {
	var _result StorageProviderProtectionMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProtectionMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStorageProviderSyncRootInfo) Put_ProtectionMode(value StorageProviderProtectionMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ProtectionMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IStorageProviderSyncRootInfo) Get_AllowPinning() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllowPinning, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStorageProviderSyncRootInfo) Put_AllowPinning(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AllowPinning, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IStorageProviderSyncRootInfo) Get_StorageProviderItemPropertyDefinitions() *IVector[*IStorageProviderItemPropertyDefinition] {
	var _result *IVector[*IStorageProviderItemPropertyDefinition]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StorageProviderItemPropertyDefinitions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageProviderSyncRootInfo) Get_RecycleBinUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RecycleBinUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageProviderSyncRootInfo) Put_RecycleBinUri(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RecycleBinUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// CF51B023-7CF1-5166-BDBA-EFD95F529E31
var IID_IStorageProviderSyncRootInfo2 = syscall.GUID{0xCF51B023, 0x7CF1, 0x5166,
	[8]byte{0xBD, 0xBA, 0xEF, 0xD9, 0x5F, 0x52, 0x9E, 0x31}}

type IStorageProviderSyncRootInfo2Interface interface {
	win32.IInspectableInterface
	Get_ProviderId() syscall.GUID
	Put_ProviderId(value syscall.GUID)
}

type IStorageProviderSyncRootInfo2Vtbl struct {
	win32.IInspectableVtbl
	Get_ProviderId uintptr
	Put_ProviderId uintptr
}

type IStorageProviderSyncRootInfo2 struct {
	win32.IInspectable
}

func (this *IStorageProviderSyncRootInfo2) Vtbl() *IStorageProviderSyncRootInfo2Vtbl {
	return (*IStorageProviderSyncRootInfo2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageProviderSyncRootInfo2) Get_ProviderId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProviderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStorageProviderSyncRootInfo2) Put_ProviderId(value syscall.GUID) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ProviderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

// 507A6617-BEF6-56FD-855E-75ACE2E45CF5
var IID_IStorageProviderSyncRootInfo3 = syscall.GUID{0x507A6617, 0xBEF6, 0x56FD,
	[8]byte{0x85, 0x5E, 0x75, 0xAC, 0xE2, 0xE4, 0x5C, 0xF5}}

type IStorageProviderSyncRootInfo3Interface interface {
	win32.IInspectableInterface
	Get_FallbackFileTypeInfo() *IVector[*IStorageProviderFileTypeInfo]
}

type IStorageProviderSyncRootInfo3Vtbl struct {
	win32.IInspectableVtbl
	Get_FallbackFileTypeInfo uintptr
}

type IStorageProviderSyncRootInfo3 struct {
	win32.IInspectable
}

func (this *IStorageProviderSyncRootInfo3) Vtbl() *IStorageProviderSyncRootInfo3Vtbl {
	return (*IStorageProviderSyncRootInfo3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageProviderSyncRootInfo3) Get_FallbackFileTypeInfo() *IVector[*IStorageProviderFileTypeInfo] {
	var _result *IVector[*IStorageProviderFileTypeInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FallbackFileTypeInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3E99FBBF-8FE3-4B40-ABC7-F6FC3D74C98E
var IID_IStorageProviderSyncRootManagerStatics = syscall.GUID{0x3E99FBBF, 0x8FE3, 0x4B40,
	[8]byte{0xAB, 0xC7, 0xF6, 0xFC, 0x3D, 0x74, 0xC9, 0x8E}}

type IStorageProviderSyncRootManagerStaticsInterface interface {
	win32.IInspectableInterface
	Register(syncRootInformation *IStorageProviderSyncRootInfo)
	Unregister(id string)
	GetSyncRootInformationForFolder(folder *IStorageFolder) *IStorageProviderSyncRootInfo
	GetSyncRootInformationForId(id string) *IStorageProviderSyncRootInfo
	GetCurrentSyncRoots() *IVectorView[*IStorageProviderSyncRootInfo]
}

type IStorageProviderSyncRootManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	Register                        uintptr
	Unregister                      uintptr
	GetSyncRootInformationForFolder uintptr
	GetSyncRootInformationForId     uintptr
	GetCurrentSyncRoots             uintptr
}

type IStorageProviderSyncRootManagerStatics struct {
	win32.IInspectable
}

func (this *IStorageProviderSyncRootManagerStatics) Vtbl() *IStorageProviderSyncRootManagerStaticsVtbl {
	return (*IStorageProviderSyncRootManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageProviderSyncRootManagerStatics) Register(syncRootInformation *IStorageProviderSyncRootInfo) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Register, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(syncRootInformation)))
	_ = _hr
}

func (this *IStorageProviderSyncRootManagerStatics) Unregister(id string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Unregister, uintptr(unsafe.Pointer(this)), NewHStr(id).Ptr)
	_ = _hr
}

func (this *IStorageProviderSyncRootManagerStatics) GetSyncRootInformationForFolder(folder *IStorageFolder) *IStorageProviderSyncRootInfo {
	var _result *IStorageProviderSyncRootInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSyncRootInformationForFolder, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(folder)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageProviderSyncRootManagerStatics) GetSyncRootInformationForId(id string) *IStorageProviderSyncRootInfo {
	var _result *IStorageProviderSyncRootInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSyncRootInformationForId, uintptr(unsafe.Pointer(this)), NewHStr(id).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageProviderSyncRootManagerStatics) GetCurrentSyncRoots() *IVectorView[*IStorageProviderSyncRootInfo] {
	var _result *IVectorView[*IStorageProviderSyncRootInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentSyncRoots, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EFB6CFEE-1374-544E-9DF1-5598D2E9CFDD
var IID_IStorageProviderSyncRootManagerStatics2 = syscall.GUID{0xEFB6CFEE, 0x1374, 0x544E,
	[8]byte{0x9D, 0xF1, 0x55, 0x98, 0xD2, 0xE9, 0xCF, 0xDD}}

type IStorageProviderSyncRootManagerStatics2Interface interface {
	win32.IInspectableInterface
	IsSupported() bool
}

type IStorageProviderSyncRootManagerStatics2Vtbl struct {
	win32.IInspectableVtbl
	IsSupported uintptr
}

type IStorageProviderSyncRootManagerStatics2 struct {
	win32.IInspectable
}

func (this *IStorageProviderSyncRootManagerStatics2) Vtbl() *IStorageProviderSyncRootManagerStatics2Vtbl {
	return (*IStorageProviderSyncRootManagerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageProviderSyncRootManagerStatics2) IsSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 0C3E0760-D846-568F-9484-105CC57B502B
var IID_IStorageProviderUICommand = syscall.GUID{0x0C3E0760, 0xD846, 0x568F,
	[8]byte{0x94, 0x84, 0x10, 0x5C, 0xC5, 0x7B, 0x50, 0x2B}}

type IStorageProviderUICommandInterface interface {
	win32.IInspectableInterface
	Get_Label() string
	Get_Description() string
	Get_Icon() *IUriRuntimeClass
	Get_State() StorageProviderUICommandState
	Invoke()
}

type IStorageProviderUICommandVtbl struct {
	win32.IInspectableVtbl
	Get_Label       uintptr
	Get_Description uintptr
	Get_Icon        uintptr
	Get_State       uintptr
	Invoke          uintptr
}

type IStorageProviderUICommand struct {
	win32.IInspectable
}

func (this *IStorageProviderUICommand) Vtbl() *IStorageProviderUICommandVtbl {
	return (*IStorageProviderUICommandVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageProviderUICommand) Get_Label() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Label, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStorageProviderUICommand) Get_Description() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Description, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStorageProviderUICommand) Get_Icon() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Icon, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageProviderUICommand) Get_State() StorageProviderUICommandState {
	var _result StorageProviderUICommandState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStorageProviderUICommand) Invoke() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Invoke, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// B29806D1-8BE0-4962-8BB6-0D4C2E14D47A
var IID_IStorageProviderUriSource = syscall.GUID{0xB29806D1, 0x8BE0, 0x4962,
	[8]byte{0x8B, 0xB6, 0x0D, 0x4C, 0x2E, 0x14, 0xD4, 0x7A}}

type IStorageProviderUriSourceInterface interface {
	win32.IInspectableInterface
	GetPathForContentUri(contentUri string, result *IStorageProviderGetPathForContentUriResult)
	GetContentInfoForPath(path string, result *IStorageProviderGetContentInfoForPathResult)
}

type IStorageProviderUriSourceVtbl struct {
	win32.IInspectableVtbl
	GetPathForContentUri  uintptr
	GetContentInfoForPath uintptr
}

type IStorageProviderUriSource struct {
	win32.IInspectable
}

func (this *IStorageProviderUriSource) Vtbl() *IStorageProviderUriSourceVtbl {
	return (*IStorageProviderUriSourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageProviderUriSource) GetPathForContentUri(contentUri string, result *IStorageProviderGetPathForContentUriResult) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPathForContentUri, uintptr(unsafe.Pointer(this)), NewHStr(contentUri).Ptr, uintptr(unsafe.Pointer(result)))
	_ = _hr
}

func (this *IStorageProviderUriSource) GetContentInfoForPath(path string, result *IStorageProviderGetContentInfoForPathResult) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetContentInfoForPath, uintptr(unsafe.Pointer(this)), NewHStr(path).Ptr, uintptr(unsafe.Pointer(result)))
	_ = _hr
}

// classes

type StorageProviderFileTypeInfo struct {
	RtClass
	*IStorageProviderFileTypeInfo
}

func NewStorageProviderFileTypeInfo_CreateInstance(fileExtension string, iconResource string) *StorageProviderFileTypeInfo {
	hs := NewHStr("Windows.Storage.Provider.StorageProviderFileTypeInfo")
	var pFac *IStorageProviderFileTypeInfoFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IStorageProviderFileTypeInfoFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IStorageProviderFileTypeInfo
	p = pFac.CreateInstance(fileExtension, iconResource)
	result := &StorageProviderFileTypeInfo{
		RtClass:                      RtClass{PInspect: &p.IInspectable},
		IStorageProviderFileTypeInfo: p,
	}
	com.AddToScope(result)
	return result
}

type StorageProviderGetContentInfoForPathResult struct {
	RtClass
	*IStorageProviderGetContentInfoForPathResult
}

func NewStorageProviderGetContentInfoForPathResult() *StorageProviderGetContentInfoForPathResult {
	hs := NewHStr("Windows.Storage.Provider.StorageProviderGetContentInfoForPathResult")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &StorageProviderGetContentInfoForPathResult{
		RtClass: RtClass{PInspect: p},
		IStorageProviderGetContentInfoForPathResult: (*IStorageProviderGetContentInfoForPathResult)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type StorageProviderGetPathForContentUriResult struct {
	RtClass
	*IStorageProviderGetPathForContentUriResult
}

func NewStorageProviderGetPathForContentUriResult() *StorageProviderGetPathForContentUriResult {
	hs := NewHStr("Windows.Storage.Provider.StorageProviderGetPathForContentUriResult")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &StorageProviderGetPathForContentUriResult{
		RtClass: RtClass{PInspect: p},
		IStorageProviderGetPathForContentUriResult: (*IStorageProviderGetPathForContentUriResult)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type StorageProviderItemProperties struct {
	RtClass
}

func NewIStorageProviderItemPropertiesStatics() *IStorageProviderItemPropertiesStatics {
	var p *IStorageProviderItemPropertiesStatics
	hs := NewHStr("Windows.Storage.Provider.StorageProviderItemProperties")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IStorageProviderItemPropertiesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type StorageProviderItemProperty struct {
	RtClass
	*IStorageProviderItemProperty
}

func NewStorageProviderItemProperty() *StorageProviderItemProperty {
	hs := NewHStr("Windows.Storage.Provider.StorageProviderItemProperty")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &StorageProviderItemProperty{
		RtClass:                      RtClass{PInspect: p},
		IStorageProviderItemProperty: (*IStorageProviderItemProperty)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type StorageProviderItemPropertyDefinition struct {
	RtClass
	*IStorageProviderItemPropertyDefinition
}

func NewStorageProviderItemPropertyDefinition() *StorageProviderItemPropertyDefinition {
	hs := NewHStr("Windows.Storage.Provider.StorageProviderItemPropertyDefinition")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &StorageProviderItemPropertyDefinition{
		RtClass:                                RtClass{PInspect: p},
		IStorageProviderItemPropertyDefinition: (*IStorageProviderItemPropertyDefinition)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type StorageProviderSyncRootInfo struct {
	RtClass
	*IStorageProviderSyncRootInfo
}

func NewStorageProviderSyncRootInfo() *StorageProviderSyncRootInfo {
	hs := NewHStr("Windows.Storage.Provider.StorageProviderSyncRootInfo")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &StorageProviderSyncRootInfo{
		RtClass:                      RtClass{PInspect: p},
		IStorageProviderSyncRootInfo: (*IStorageProviderSyncRootInfo)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type StorageProviderSyncRootManager struct {
	RtClass
}

func NewIStorageProviderSyncRootManagerStatics2() *IStorageProviderSyncRootManagerStatics2 {
	var p *IStorageProviderSyncRootManagerStatics2
	hs := NewHStr("Windows.Storage.Provider.StorageProviderSyncRootManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IStorageProviderSyncRootManagerStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIStorageProviderSyncRootManagerStatics() *IStorageProviderSyncRootManagerStatics {
	var p *IStorageProviderSyncRootManagerStatics
	hs := NewHStr("Windows.Storage.Provider.StorageProviderSyncRootManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IStorageProviderSyncRootManagerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
