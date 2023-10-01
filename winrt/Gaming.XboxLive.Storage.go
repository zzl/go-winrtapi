package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type GameSaveErrorStatus int32

const (
	GameSaveErrorStatus_Ok                     GameSaveErrorStatus = 0
	GameSaveErrorStatus_Abort                  GameSaveErrorStatus = -2147467260
	GameSaveErrorStatus_InvalidContainerName   GameSaveErrorStatus = -2138898431
	GameSaveErrorStatus_NoAccess               GameSaveErrorStatus = -2138898430
	GameSaveErrorStatus_OutOfLocalStorage      GameSaveErrorStatus = -2138898429
	GameSaveErrorStatus_UserCanceled           GameSaveErrorStatus = -2138898428
	GameSaveErrorStatus_UpdateTooBig           GameSaveErrorStatus = -2138898427
	GameSaveErrorStatus_QuotaExceeded          GameSaveErrorStatus = -2138898426
	GameSaveErrorStatus_ProvidedBufferTooSmall GameSaveErrorStatus = -2138898425
	GameSaveErrorStatus_BlobNotFound           GameSaveErrorStatus = -2138898424
	GameSaveErrorStatus_NoXboxLiveInfo         GameSaveErrorStatus = -2138898423
	GameSaveErrorStatus_ContainerNotInSync     GameSaveErrorStatus = -2138898422
	GameSaveErrorStatus_ContainerSyncFailed    GameSaveErrorStatus = -2138898421
	GameSaveErrorStatus_UserHasNoXboxLiveInfo  GameSaveErrorStatus = -2138898420
	GameSaveErrorStatus_ObjectExpired          GameSaveErrorStatus = -2138898419
)

// interfaces

// 917281E0-7201-4953-AA2C-4008F03AEF45
var IID_IGameSaveBlobGetResult = syscall.GUID{0x917281E0, 0x7201, 0x4953,
	[8]byte{0xAA, 0x2C, 0x40, 0x08, 0xF0, 0x3A, 0xEF, 0x45}}

type IGameSaveBlobGetResultInterface interface {
	win32.IInspectableInterface
	Get_Status() GameSaveErrorStatus
	Get_Value() *IMapView[string, *IBuffer]
}

type IGameSaveBlobGetResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
	Get_Value  uintptr
}

type IGameSaveBlobGetResult struct {
	win32.IInspectable
}

func (this *IGameSaveBlobGetResult) Vtbl() *IGameSaveBlobGetResultVtbl {
	return (*IGameSaveBlobGetResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameSaveBlobGetResult) Get_Status() GameSaveErrorStatus {
	var _result GameSaveErrorStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGameSaveBlobGetResult) Get_Value() *IMapView[string, *IBuffer] {
	var _result *IMapView[string, *IBuffer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// ADD38034-BAF0-4645-B6D0-46EDAFFB3C2B
var IID_IGameSaveBlobInfo = syscall.GUID{0xADD38034, 0xBAF0, 0x4645,
	[8]byte{0xB6, 0xD0, 0x46, 0xED, 0xAF, 0xFB, 0x3C, 0x2B}}

type IGameSaveBlobInfoInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Get_Size() uint32
}

type IGameSaveBlobInfoVtbl struct {
	win32.IInspectableVtbl
	Get_Name uintptr
	Get_Size uintptr
}

type IGameSaveBlobInfo struct {
	win32.IInspectable
}

func (this *IGameSaveBlobInfo) Vtbl() *IGameSaveBlobInfoVtbl {
	return (*IGameSaveBlobInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameSaveBlobInfo) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGameSaveBlobInfo) Get_Size() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Size, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C7578582-3697-42BF-989C-665D923B5231
var IID_IGameSaveBlobInfoGetResult = syscall.GUID{0xC7578582, 0x3697, 0x42BF,
	[8]byte{0x98, 0x9C, 0x66, 0x5D, 0x92, 0x3B, 0x52, 0x31}}

type IGameSaveBlobInfoGetResultInterface interface {
	win32.IInspectableInterface
	Get_Status() GameSaveErrorStatus
	Get_Value() *IVectorView[*IGameSaveBlobInfo]
}

type IGameSaveBlobInfoGetResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
	Get_Value  uintptr
}

type IGameSaveBlobInfoGetResult struct {
	win32.IInspectable
}

func (this *IGameSaveBlobInfoGetResult) Vtbl() *IGameSaveBlobInfoGetResultVtbl {
	return (*IGameSaveBlobInfoGetResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameSaveBlobInfoGetResult) Get_Status() GameSaveErrorStatus {
	var _result GameSaveErrorStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGameSaveBlobInfoGetResult) Get_Value() *IVectorView[*IGameSaveBlobInfo] {
	var _result *IVectorView[*IGameSaveBlobInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9FDD74B2-EEEE-447B-A9D2-7F96C0F83208
var IID_IGameSaveBlobInfoQuery = syscall.GUID{0x9FDD74B2, 0xEEEE, 0x447B,
	[8]byte{0xA9, 0xD2, 0x7F, 0x96, 0xC0, 0xF8, 0x32, 0x08}}

type IGameSaveBlobInfoQueryInterface interface {
	win32.IInspectableInterface
	GetBlobInfoAsync() *IAsyncOperation[*IGameSaveBlobInfoGetResult]
	GetBlobInfoWithIndexAndMaxAsync(startIndex uint32, maxNumberOfItems uint32) *IAsyncOperation[*IGameSaveBlobInfoGetResult]
	GetItemCountAsync() *IAsyncOperation[uint32]
}

type IGameSaveBlobInfoQueryVtbl struct {
	win32.IInspectableVtbl
	GetBlobInfoAsync                uintptr
	GetBlobInfoWithIndexAndMaxAsync uintptr
	GetItemCountAsync               uintptr
}

type IGameSaveBlobInfoQuery struct {
	win32.IInspectable
}

func (this *IGameSaveBlobInfoQuery) Vtbl() *IGameSaveBlobInfoQueryVtbl {
	return (*IGameSaveBlobInfoQueryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameSaveBlobInfoQuery) GetBlobInfoAsync() *IAsyncOperation[*IGameSaveBlobInfoGetResult] {
	var _result *IAsyncOperation[*IGameSaveBlobInfoGetResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetBlobInfoAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGameSaveBlobInfoQuery) GetBlobInfoWithIndexAndMaxAsync(startIndex uint32, maxNumberOfItems uint32) *IAsyncOperation[*IGameSaveBlobInfoGetResult] {
	var _result *IAsyncOperation[*IGameSaveBlobInfoGetResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetBlobInfoWithIndexAndMaxAsync, uintptr(unsafe.Pointer(this)), uintptr(startIndex), uintptr(maxNumberOfItems), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGameSaveBlobInfoQuery) GetItemCountAsync() *IAsyncOperation[uint32] {
	var _result *IAsyncOperation[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetItemCountAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C3C08F89-563F-4ECD-9C6F-33FD0E323D10
var IID_IGameSaveContainer = syscall.GUID{0xC3C08F89, 0x563F, 0x4ECD,
	[8]byte{0x9C, 0x6F, 0x33, 0xFD, 0x0E, 0x32, 0x3D, 0x10}}

type IGameSaveContainerInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Get_Provider() *IGameSaveProvider
	SubmitUpdatesAsync(blobsToWrite *IMapView[string, *IBuffer], blobsToDelete *IIterable[string], displayName string) *IAsyncOperation[*IGameSaveOperationResult]
	ReadAsync(blobsToRead *IMapView[string, *IBuffer]) *IAsyncOperation[*IGameSaveOperationResult]
	GetAsync(blobsToRead *IIterable[string]) *IAsyncOperation[*IGameSaveBlobGetResult]
	SubmitPropertySetUpdatesAsync(blobsToWrite *IPropertySet, blobsToDelete *IIterable[string], displayName string) *IAsyncOperation[*IGameSaveOperationResult]
	CreateBlobInfoQuery(blobNamePrefix string) *IGameSaveBlobInfoQuery
}

type IGameSaveContainerVtbl struct {
	win32.IInspectableVtbl
	Get_Name                      uintptr
	Get_Provider                  uintptr
	SubmitUpdatesAsync            uintptr
	ReadAsync                     uintptr
	GetAsync                      uintptr
	SubmitPropertySetUpdatesAsync uintptr
	CreateBlobInfoQuery           uintptr
}

type IGameSaveContainer struct {
	win32.IInspectable
}

func (this *IGameSaveContainer) Vtbl() *IGameSaveContainerVtbl {
	return (*IGameSaveContainerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameSaveContainer) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGameSaveContainer) Get_Provider() *IGameSaveProvider {
	var _result *IGameSaveProvider
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Provider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGameSaveContainer) SubmitUpdatesAsync(blobsToWrite *IMapView[string, *IBuffer], blobsToDelete *IIterable[string], displayName string) *IAsyncOperation[*IGameSaveOperationResult] {
	var _result *IAsyncOperation[*IGameSaveOperationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SubmitUpdatesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(blobsToWrite)), uintptr(unsafe.Pointer(blobsToDelete)), NewHStr(displayName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGameSaveContainer) ReadAsync(blobsToRead *IMapView[string, *IBuffer]) *IAsyncOperation[*IGameSaveOperationResult] {
	var _result *IAsyncOperation[*IGameSaveOperationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(blobsToRead)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGameSaveContainer) GetAsync(blobsToRead *IIterable[string]) *IAsyncOperation[*IGameSaveBlobGetResult] {
	var _result *IAsyncOperation[*IGameSaveBlobGetResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(blobsToRead)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGameSaveContainer) SubmitPropertySetUpdatesAsync(blobsToWrite *IPropertySet, blobsToDelete *IIterable[string], displayName string) *IAsyncOperation[*IGameSaveOperationResult] {
	var _result *IAsyncOperation[*IGameSaveOperationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SubmitPropertySetUpdatesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(blobsToWrite)), uintptr(unsafe.Pointer(blobsToDelete)), NewHStr(displayName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGameSaveContainer) CreateBlobInfoQuery(blobNamePrefix string) *IGameSaveBlobInfoQuery {
	var _result *IGameSaveBlobInfoQuery
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateBlobInfoQuery, uintptr(unsafe.Pointer(this)), NewHStr(blobNamePrefix).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B7E27300-155D-4BB4-B2BA-930306F391B5
var IID_IGameSaveContainerInfo = syscall.GUID{0xB7E27300, 0x155D, 0x4BB4,
	[8]byte{0xB2, 0xBA, 0x93, 0x03, 0x06, 0xF3, 0x91, 0xB5}}

type IGameSaveContainerInfoInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Get_TotalSize() uint64
	Get_DisplayName() string
	Get_LastModifiedTime() DateTime
	Get_NeedsSync() bool
}

type IGameSaveContainerInfoVtbl struct {
	win32.IInspectableVtbl
	Get_Name             uintptr
	Get_TotalSize        uintptr
	Get_DisplayName      uintptr
	Get_LastModifiedTime uintptr
	Get_NeedsSync        uintptr
}

type IGameSaveContainerInfo struct {
	win32.IInspectable
}

func (this *IGameSaveContainerInfo) Vtbl() *IGameSaveContainerInfoVtbl {
	return (*IGameSaveContainerInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameSaveContainerInfo) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGameSaveContainerInfo) Get_TotalSize() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TotalSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGameSaveContainerInfo) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGameSaveContainerInfo) Get_LastModifiedTime() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LastModifiedTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGameSaveContainerInfo) Get_NeedsSync() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NeedsSync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// FFC50D74-C581-4F9D-9E39-30A10C1E4C50
var IID_IGameSaveContainerInfoGetResult = syscall.GUID{0xFFC50D74, 0xC581, 0x4F9D,
	[8]byte{0x9E, 0x39, 0x30, 0xA1, 0x0C, 0x1E, 0x4C, 0x50}}

type IGameSaveContainerInfoGetResultInterface interface {
	win32.IInspectableInterface
	Get_Status() GameSaveErrorStatus
	Get_Value() *IVectorView[*IGameSaveContainerInfo]
}

type IGameSaveContainerInfoGetResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
	Get_Value  uintptr
}

type IGameSaveContainerInfoGetResult struct {
	win32.IInspectable
}

func (this *IGameSaveContainerInfoGetResult) Vtbl() *IGameSaveContainerInfoGetResultVtbl {
	return (*IGameSaveContainerInfoGetResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameSaveContainerInfoGetResult) Get_Status() GameSaveErrorStatus {
	var _result GameSaveErrorStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGameSaveContainerInfoGetResult) Get_Value() *IVectorView[*IGameSaveContainerInfo] {
	var _result *IVectorView[*IGameSaveContainerInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3C94E863-6F80-4327-9327-FFC11AFD42B3
var IID_IGameSaveContainerInfoQuery = syscall.GUID{0x3C94E863, 0x6F80, 0x4327,
	[8]byte{0x93, 0x27, 0xFF, 0xC1, 0x1A, 0xFD, 0x42, 0xB3}}

type IGameSaveContainerInfoQueryInterface interface {
	win32.IInspectableInterface
	GetContainerInfoAsync() *IAsyncOperation[*IGameSaveContainerInfoGetResult]
	GetContainerInfoWithIndexAndMaxAsync(startIndex uint32, maxNumberOfItems uint32) *IAsyncOperation[*IGameSaveContainerInfoGetResult]
	GetItemCountAsync() *IAsyncOperation[uint32]
}

type IGameSaveContainerInfoQueryVtbl struct {
	win32.IInspectableVtbl
	GetContainerInfoAsync                uintptr
	GetContainerInfoWithIndexAndMaxAsync uintptr
	GetItemCountAsync                    uintptr
}

type IGameSaveContainerInfoQuery struct {
	win32.IInspectable
}

func (this *IGameSaveContainerInfoQuery) Vtbl() *IGameSaveContainerInfoQueryVtbl {
	return (*IGameSaveContainerInfoQueryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameSaveContainerInfoQuery) GetContainerInfoAsync() *IAsyncOperation[*IGameSaveContainerInfoGetResult] {
	var _result *IAsyncOperation[*IGameSaveContainerInfoGetResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetContainerInfoAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGameSaveContainerInfoQuery) GetContainerInfoWithIndexAndMaxAsync(startIndex uint32, maxNumberOfItems uint32) *IAsyncOperation[*IGameSaveContainerInfoGetResult] {
	var _result *IAsyncOperation[*IGameSaveContainerInfoGetResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetContainerInfoWithIndexAndMaxAsync, uintptr(unsafe.Pointer(this)), uintptr(startIndex), uintptr(maxNumberOfItems), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGameSaveContainerInfoQuery) GetItemCountAsync() *IAsyncOperation[uint32] {
	var _result *IAsyncOperation[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetItemCountAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CF0F1A05-24A0-4582-9A55-B1BBBB9388D8
var IID_IGameSaveOperationResult = syscall.GUID{0xCF0F1A05, 0x24A0, 0x4582,
	[8]byte{0x9A, 0x55, 0xB1, 0xBB, 0xBB, 0x93, 0x88, 0xD8}}

type IGameSaveOperationResultInterface interface {
	win32.IInspectableInterface
	Get_Status() GameSaveErrorStatus
}

type IGameSaveOperationResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
}

type IGameSaveOperationResult struct {
	win32.IInspectable
}

func (this *IGameSaveOperationResult) Vtbl() *IGameSaveOperationResultVtbl {
	return (*IGameSaveOperationResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameSaveOperationResult) Get_Status() GameSaveErrorStatus {
	var _result GameSaveErrorStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 90A60394-80FE-4211-97F8-A5DE14DD95D2
var IID_IGameSaveProvider = syscall.GUID{0x90A60394, 0x80FE, 0x4211,
	[8]byte{0x97, 0xF8, 0xA5, 0xDE, 0x14, 0xDD, 0x95, 0xD2}}

type IGameSaveProviderInterface interface {
	win32.IInspectableInterface
	Get_User() *IUser
	CreateContainer(name string) *IGameSaveContainer
	DeleteContainerAsync(name string) *IAsyncOperation[*IGameSaveOperationResult]
	CreateContainerInfoQuery() *IGameSaveContainerInfoQuery
	CreateContainerInfoQueryWithName(containerNamePrefix string) *IGameSaveContainerInfoQuery
	GetRemainingBytesInQuotaAsync() *IAsyncOperation[int64]
	Get_ContainersChangedSinceLastSync() *IVectorView[string]
}

type IGameSaveProviderVtbl struct {
	win32.IInspectableVtbl
	Get_User                           uintptr
	CreateContainer                    uintptr
	DeleteContainerAsync               uintptr
	CreateContainerInfoQuery           uintptr
	CreateContainerInfoQueryWithName   uintptr
	GetRemainingBytesInQuotaAsync      uintptr
	Get_ContainersChangedSinceLastSync uintptr
}

type IGameSaveProvider struct {
	win32.IInspectable
}

func (this *IGameSaveProvider) Vtbl() *IGameSaveProviderVtbl {
	return (*IGameSaveProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameSaveProvider) Get_User() *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_User, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGameSaveProvider) CreateContainer(name string) *IGameSaveContainer {
	var _result *IGameSaveContainer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateContainer, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGameSaveProvider) DeleteContainerAsync(name string) *IAsyncOperation[*IGameSaveOperationResult] {
	var _result *IAsyncOperation[*IGameSaveOperationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DeleteContainerAsync, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGameSaveProvider) CreateContainerInfoQuery() *IGameSaveContainerInfoQuery {
	var _result *IGameSaveContainerInfoQuery
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateContainerInfoQuery, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGameSaveProvider) CreateContainerInfoQueryWithName(containerNamePrefix string) *IGameSaveContainerInfoQuery {
	var _result *IGameSaveContainerInfoQuery
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateContainerInfoQueryWithName, uintptr(unsafe.Pointer(this)), NewHStr(containerNamePrefix).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGameSaveProvider) GetRemainingBytesInQuotaAsync() *IAsyncOperation[int64] {
	var _result *IAsyncOperation[int64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetRemainingBytesInQuotaAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGameSaveProvider) Get_ContainersChangedSinceLastSync() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContainersChangedSinceLastSync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3AB90816-D393-4D65-AC16-41C3E67AB945
var IID_IGameSaveProviderGetResult = syscall.GUID{0x3AB90816, 0xD393, 0x4D65,
	[8]byte{0xAC, 0x16, 0x41, 0xC3, 0xE6, 0x7A, 0xB9, 0x45}}

type IGameSaveProviderGetResultInterface interface {
	win32.IInspectableInterface
	Get_Status() GameSaveErrorStatus
	Get_Value() *IGameSaveProvider
}

type IGameSaveProviderGetResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
	Get_Value  uintptr
}

type IGameSaveProviderGetResult struct {
	win32.IInspectable
}

func (this *IGameSaveProviderGetResult) Vtbl() *IGameSaveProviderGetResultVtbl {
	return (*IGameSaveProviderGetResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameSaveProviderGetResult) Get_Status() GameSaveErrorStatus {
	var _result GameSaveErrorStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGameSaveProviderGetResult) Get_Value() *IGameSaveProvider {
	var _result *IGameSaveProvider
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D01D3ED0-7B03-449D-8CBD-3402842A1048
var IID_IGameSaveProviderStatics = syscall.GUID{0xD01D3ED0, 0x7B03, 0x449D,
	[8]byte{0x8C, 0xBD, 0x34, 0x02, 0x84, 0x2A, 0x10, 0x48}}

type IGameSaveProviderStaticsInterface interface {
	win32.IInspectableInterface
	GetForUserAsync(user *IUser, serviceConfigId string) *IAsyncOperation[*IGameSaveProviderGetResult]
	GetSyncOnDemandForUserAsync(user *IUser, serviceConfigId string) *IAsyncOperation[*IGameSaveProviderGetResult]
}

type IGameSaveProviderStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForUserAsync             uintptr
	GetSyncOnDemandForUserAsync uintptr
}

type IGameSaveProviderStatics struct {
	win32.IInspectable
}

func (this *IGameSaveProviderStatics) Vtbl() *IGameSaveProviderStaticsVtbl {
	return (*IGameSaveProviderStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameSaveProviderStatics) GetForUserAsync(user *IUser, serviceConfigId string) *IAsyncOperation[*IGameSaveProviderGetResult] {
	var _result *IAsyncOperation[*IGameSaveProviderGetResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForUserAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), NewHStr(serviceConfigId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGameSaveProviderStatics) GetSyncOnDemandForUserAsync(user *IUser, serviceConfigId string) *IAsyncOperation[*IGameSaveProviderGetResult] {
	var _result *IAsyncOperation[*IGameSaveProviderGetResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSyncOnDemandForUserAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), NewHStr(serviceConfigId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type GameSaveBlobGetResult struct {
	RtClass
	*IGameSaveBlobGetResult
}

type GameSaveBlobInfo struct {
	RtClass
	*IGameSaveBlobInfo
}

type GameSaveBlobInfoGetResult struct {
	RtClass
	*IGameSaveBlobInfoGetResult
}

type GameSaveBlobInfoQuery struct {
	RtClass
	*IGameSaveBlobInfoQuery
}

type GameSaveContainer struct {
	RtClass
	*IGameSaveContainer
}

type GameSaveContainerInfo struct {
	RtClass
	*IGameSaveContainerInfo
}

type GameSaveContainerInfoGetResult struct {
	RtClass
	*IGameSaveContainerInfoGetResult
}

type GameSaveContainerInfoQuery struct {
	RtClass
	*IGameSaveContainerInfoQuery
}

type GameSaveOperationResult struct {
	RtClass
	*IGameSaveOperationResult
}

type GameSaveProvider struct {
	RtClass
	*IGameSaveProvider
}

func NewIGameSaveProviderStatics() *IGameSaveProviderStatics {
	var p *IGameSaveProviderStatics
	hs := NewHStr("Windows.Gaming.XboxLive.Storage.GameSaveProvider")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGameSaveProviderStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type GameSaveProviderGetResult struct {
	RtClass
	*IGameSaveProviderGetResult
}
