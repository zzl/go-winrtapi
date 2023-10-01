package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type PhotoImportAccessMode int32

const (
	PhotoImportAccessMode_ReadWrite     PhotoImportAccessMode = 0
	PhotoImportAccessMode_ReadOnly      PhotoImportAccessMode = 1
	PhotoImportAccessMode_ReadAndDelete PhotoImportAccessMode = 2
)

// enum
type PhotoImportConnectionTransport int32

const (
	PhotoImportConnectionTransport_Unknown   PhotoImportConnectionTransport = 0
	PhotoImportConnectionTransport_Usb       PhotoImportConnectionTransport = 1
	PhotoImportConnectionTransport_IP        PhotoImportConnectionTransport = 2
	PhotoImportConnectionTransport_Bluetooth PhotoImportConnectionTransport = 3
)

// enum
type PhotoImportContentType int32

const (
	PhotoImportContentType_Unknown PhotoImportContentType = 0
	PhotoImportContentType_Image   PhotoImportContentType = 1
	PhotoImportContentType_Video   PhotoImportContentType = 2
)

// enum
type PhotoImportContentTypeFilter int32

const (
	PhotoImportContentTypeFilter_OnlyImages                    PhotoImportContentTypeFilter = 0
	PhotoImportContentTypeFilter_OnlyVideos                    PhotoImportContentTypeFilter = 1
	PhotoImportContentTypeFilter_ImagesAndVideos               PhotoImportContentTypeFilter = 2
	PhotoImportContentTypeFilter_ImagesAndVideosFromCameraRoll PhotoImportContentTypeFilter = 3
)

// enum
type PhotoImportImportMode int32

const (
	PhotoImportImportMode_ImportEverything          PhotoImportImportMode = 0
	PhotoImportImportMode_IgnoreSidecars            PhotoImportImportMode = 1
	PhotoImportImportMode_IgnoreSiblings            PhotoImportImportMode = 2
	PhotoImportImportMode_IgnoreSidecarsAndSiblings PhotoImportImportMode = 3
)

// enum
type PhotoImportItemSelectionMode int32

const (
	PhotoImportItemSelectionMode_SelectAll  PhotoImportItemSelectionMode = 0
	PhotoImportItemSelectionMode_SelectNone PhotoImportItemSelectionMode = 1
	PhotoImportItemSelectionMode_SelectNew  PhotoImportItemSelectionMode = 2
)

// enum
type PhotoImportPowerSource int32

const (
	PhotoImportPowerSource_Unknown  PhotoImportPowerSource = 0
	PhotoImportPowerSource_Battery  PhotoImportPowerSource = 1
	PhotoImportPowerSource_External PhotoImportPowerSource = 2
)

// enum
type PhotoImportSourceType int32

const (
	PhotoImportSourceType_Generic             PhotoImportSourceType = 0
	PhotoImportSourceType_Camera              PhotoImportSourceType = 1
	PhotoImportSourceType_MediaPlayer         PhotoImportSourceType = 2
	PhotoImportSourceType_Phone               PhotoImportSourceType = 3
	PhotoImportSourceType_Video               PhotoImportSourceType = 4
	PhotoImportSourceType_PersonalInfoManager PhotoImportSourceType = 5
	PhotoImportSourceType_AudioRecorder       PhotoImportSourceType = 6
)

// enum
type PhotoImportStage int32

const (
	PhotoImportStage_NotStarted                      PhotoImportStage = 0
	PhotoImportStage_FindingItems                    PhotoImportStage = 1
	PhotoImportStage_ImportingItems                  PhotoImportStage = 2
	PhotoImportStage_DeletingImportedItemsFromSource PhotoImportStage = 3
)

// enum
type PhotoImportStorageMediumType int32

const (
	PhotoImportStorageMediumType_Undefined PhotoImportStorageMediumType = 0
	PhotoImportStorageMediumType_Fixed     PhotoImportStorageMediumType = 1
	PhotoImportStorageMediumType_Removable PhotoImportStorageMediumType = 2
)

// enum
type PhotoImportSubfolderCreationMode int32

const (
	PhotoImportSubfolderCreationMode_DoNotCreateSubfolders        PhotoImportSubfolderCreationMode = 0
	PhotoImportSubfolderCreationMode_CreateSubfoldersFromFileDate PhotoImportSubfolderCreationMode = 1
	PhotoImportSubfolderCreationMode_CreateSubfoldersFromExifDate PhotoImportSubfolderCreationMode = 2
	PhotoImportSubfolderCreationMode_KeepOriginalFolderStructure  PhotoImportSubfolderCreationMode = 3
)

// enum
type PhotoImportSubfolderDateFormat int32

const (
	PhotoImportSubfolderDateFormat_Year         PhotoImportSubfolderDateFormat = 0
	PhotoImportSubfolderDateFormat_YearMonth    PhotoImportSubfolderDateFormat = 1
	PhotoImportSubfolderDateFormat_YearMonthDay PhotoImportSubfolderDateFormat = 2
)

// structs

type PhotoImportProgress struct {
	ItemsImported      uint32
	TotalItemsToImport uint32
	BytesImported      uint64
	TotalBytesToImport uint64
	ImportProgress     float64
}

// interfaces

// F4E112F8-843D-428A-A1A6-81510292B0AE
var IID_IPhotoImportDeleteImportedItemsFromSourceResult = syscall.GUID{0xF4E112F8, 0x843D, 0x428A,
	[8]byte{0xA1, 0xA6, 0x81, 0x51, 0x02, 0x92, 0xB0, 0xAE}}

type IPhotoImportDeleteImportedItemsFromSourceResultInterface interface {
	win32.IInspectableInterface
	Get_Session() *IPhotoImportSession
	Get_HasSucceeded() bool
	Get_DeletedItems() *IVectorView[*IPhotoImportItem]
	Get_PhotosCount() uint32
	Get_PhotosSizeInBytes() uint64
	Get_VideosCount() uint32
	Get_VideosSizeInBytes() uint64
	Get_SidecarsCount() uint32
	Get_SidecarsSizeInBytes() uint64
	Get_SiblingsCount() uint32
	Get_SiblingsSizeInBytes() uint64
	Get_TotalCount() uint32
	Get_TotalSizeInBytes() uint64
}

type IPhotoImportDeleteImportedItemsFromSourceResultVtbl struct {
	win32.IInspectableVtbl
	Get_Session             uintptr
	Get_HasSucceeded        uintptr
	Get_DeletedItems        uintptr
	Get_PhotosCount         uintptr
	Get_PhotosSizeInBytes   uintptr
	Get_VideosCount         uintptr
	Get_VideosSizeInBytes   uintptr
	Get_SidecarsCount       uintptr
	Get_SidecarsSizeInBytes uintptr
	Get_SiblingsCount       uintptr
	Get_SiblingsSizeInBytes uintptr
	Get_TotalCount          uintptr
	Get_TotalSizeInBytes    uintptr
}

type IPhotoImportDeleteImportedItemsFromSourceResult struct {
	win32.IInspectable
}

func (this *IPhotoImportDeleteImportedItemsFromSourceResult) Vtbl() *IPhotoImportDeleteImportedItemsFromSourceResultVtbl {
	return (*IPhotoImportDeleteImportedItemsFromSourceResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPhotoImportDeleteImportedItemsFromSourceResult) Get_Session() *IPhotoImportSession {
	var _result *IPhotoImportSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Session, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoImportDeleteImportedItemsFromSourceResult) Get_HasSucceeded() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasSucceeded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportDeleteImportedItemsFromSourceResult) Get_DeletedItems() *IVectorView[*IPhotoImportItem] {
	var _result *IVectorView[*IPhotoImportItem]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeletedItems, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoImportDeleteImportedItemsFromSourceResult) Get_PhotosCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhotosCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportDeleteImportedItemsFromSourceResult) Get_PhotosSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhotosSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportDeleteImportedItemsFromSourceResult) Get_VideosCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideosCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportDeleteImportedItemsFromSourceResult) Get_VideosSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideosSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportDeleteImportedItemsFromSourceResult) Get_SidecarsCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SidecarsCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportDeleteImportedItemsFromSourceResult) Get_SidecarsSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SidecarsSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportDeleteImportedItemsFromSourceResult) Get_SiblingsCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SiblingsCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportDeleteImportedItemsFromSourceResult) Get_SiblingsSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SiblingsSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportDeleteImportedItemsFromSourceResult) Get_TotalCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TotalCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportDeleteImportedItemsFromSourceResult) Get_TotalSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TotalSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 3915E647-6C78-492B-844E-8FE5E8F6BFB9
var IID_IPhotoImportFindItemsResult = syscall.GUID{0x3915E647, 0x6C78, 0x492B,
	[8]byte{0x84, 0x4E, 0x8F, 0xE5, 0xE8, 0xF6, 0xBF, 0xB9}}

type IPhotoImportFindItemsResultInterface interface {
	win32.IInspectableInterface
	Get_Session() *IPhotoImportSession
	Get_HasSucceeded() bool
	Get_FoundItems() *IVectorView[*IPhotoImportItem]
	Get_PhotosCount() uint32
	Get_PhotosSizeInBytes() uint64
	Get_VideosCount() uint32
	Get_VideosSizeInBytes() uint64
	Get_SidecarsCount() uint32
	Get_SidecarsSizeInBytes() uint64
	Get_SiblingsCount() uint32
	Get_SiblingsSizeInBytes() uint64
	Get_TotalCount() uint32
	Get_TotalSizeInBytes() uint64
	SelectAll()
	SelectNone()
	SelectNewAsync() *IAsyncAction
	SetImportMode(value PhotoImportImportMode)
	Get_ImportMode() PhotoImportImportMode
	Get_SelectedPhotosCount() uint32
	Get_SelectedPhotosSizeInBytes() uint64
	Get_SelectedVideosCount() uint32
	Get_SelectedVideosSizeInBytes() uint64
	Get_SelectedSidecarsCount() uint32
	Get_SelectedSidecarsSizeInBytes() uint64
	Get_SelectedSiblingsCount() uint32
	Get_SelectedSiblingsSizeInBytes() uint64
	Get_SelectedTotalCount() uint32
	Get_SelectedTotalSizeInBytes() uint64
	Add_SelectionChanged(value TypedEventHandler[*IPhotoImportFindItemsResult, *IPhotoImportSelectionChangedEventArgs]) EventRegistrationToken
	Remove_SelectionChanged(token EventRegistrationToken)
	ImportItemsAsync() *IAsyncOperationWithProgress[*IPhotoImportImportItemsResult, PhotoImportProgress]
	Add_ItemImported(value TypedEventHandler[*IPhotoImportFindItemsResult, *IPhotoImportItemImportedEventArgs]) EventRegistrationToken
	Remove_ItemImported(token EventRegistrationToken)
}

type IPhotoImportFindItemsResultVtbl struct {
	win32.IInspectableVtbl
	Get_Session                     uintptr
	Get_HasSucceeded                uintptr
	Get_FoundItems                  uintptr
	Get_PhotosCount                 uintptr
	Get_PhotosSizeInBytes           uintptr
	Get_VideosCount                 uintptr
	Get_VideosSizeInBytes           uintptr
	Get_SidecarsCount               uintptr
	Get_SidecarsSizeInBytes         uintptr
	Get_SiblingsCount               uintptr
	Get_SiblingsSizeInBytes         uintptr
	Get_TotalCount                  uintptr
	Get_TotalSizeInBytes            uintptr
	SelectAll                       uintptr
	SelectNone                      uintptr
	SelectNewAsync                  uintptr
	SetImportMode                   uintptr
	Get_ImportMode                  uintptr
	Get_SelectedPhotosCount         uintptr
	Get_SelectedPhotosSizeInBytes   uintptr
	Get_SelectedVideosCount         uintptr
	Get_SelectedVideosSizeInBytes   uintptr
	Get_SelectedSidecarsCount       uintptr
	Get_SelectedSidecarsSizeInBytes uintptr
	Get_SelectedSiblingsCount       uintptr
	Get_SelectedSiblingsSizeInBytes uintptr
	Get_SelectedTotalCount          uintptr
	Get_SelectedTotalSizeInBytes    uintptr
	Add_SelectionChanged            uintptr
	Remove_SelectionChanged         uintptr
	ImportItemsAsync                uintptr
	Add_ItemImported                uintptr
	Remove_ItemImported             uintptr
}

type IPhotoImportFindItemsResult struct {
	win32.IInspectable
}

func (this *IPhotoImportFindItemsResult) Vtbl() *IPhotoImportFindItemsResultVtbl {
	return (*IPhotoImportFindItemsResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPhotoImportFindItemsResult) Get_Session() *IPhotoImportSession {
	var _result *IPhotoImportSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Session, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoImportFindItemsResult) Get_HasSucceeded() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasSucceeded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportFindItemsResult) Get_FoundItems() *IVectorView[*IPhotoImportItem] {
	var _result *IVectorView[*IPhotoImportItem]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FoundItems, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoImportFindItemsResult) Get_PhotosCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhotosCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportFindItemsResult) Get_PhotosSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhotosSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportFindItemsResult) Get_VideosCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideosCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportFindItemsResult) Get_VideosSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideosSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportFindItemsResult) Get_SidecarsCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SidecarsCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportFindItemsResult) Get_SidecarsSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SidecarsSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportFindItemsResult) Get_SiblingsCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SiblingsCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportFindItemsResult) Get_SiblingsSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SiblingsSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportFindItemsResult) Get_TotalCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TotalCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportFindItemsResult) Get_TotalSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TotalSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportFindItemsResult) SelectAll() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SelectAll, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IPhotoImportFindItemsResult) SelectNone() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SelectNone, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IPhotoImportFindItemsResult) SelectNewAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SelectNewAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoImportFindItemsResult) SetImportMode(value PhotoImportImportMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetImportMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPhotoImportFindItemsResult) Get_ImportMode() PhotoImportImportMode {
	var _result PhotoImportImportMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ImportMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportFindItemsResult) Get_SelectedPhotosCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SelectedPhotosCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportFindItemsResult) Get_SelectedPhotosSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SelectedPhotosSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportFindItemsResult) Get_SelectedVideosCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SelectedVideosCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportFindItemsResult) Get_SelectedVideosSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SelectedVideosSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportFindItemsResult) Get_SelectedSidecarsCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SelectedSidecarsCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportFindItemsResult) Get_SelectedSidecarsSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SelectedSidecarsSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportFindItemsResult) Get_SelectedSiblingsCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SelectedSiblingsCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportFindItemsResult) Get_SelectedSiblingsSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SelectedSiblingsSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportFindItemsResult) Get_SelectedTotalCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SelectedTotalCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportFindItemsResult) Get_SelectedTotalSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SelectedTotalSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportFindItemsResult) Add_SelectionChanged(value TypedEventHandler[*IPhotoImportFindItemsResult, *IPhotoImportSelectionChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SelectionChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportFindItemsResult) Remove_SelectionChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SelectionChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPhotoImportFindItemsResult) ImportItemsAsync() *IAsyncOperationWithProgress[*IPhotoImportImportItemsResult, PhotoImportProgress] {
	var _result *IAsyncOperationWithProgress[*IPhotoImportImportItemsResult, PhotoImportProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ImportItemsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoImportFindItemsResult) Add_ItemImported(value TypedEventHandler[*IPhotoImportFindItemsResult, *IPhotoImportItemImportedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ItemImported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportFindItemsResult) Remove_ItemImported(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ItemImported, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// FBDD6A3B-ECF9-406A-815E-5015625B0A88
var IID_IPhotoImportFindItemsResult2 = syscall.GUID{0xFBDD6A3B, 0xECF9, 0x406A,
	[8]byte{0x81, 0x5E, 0x50, 0x15, 0x62, 0x5B, 0x0A, 0x88}}

type IPhotoImportFindItemsResult2Interface interface {
	win32.IInspectableInterface
	AddItemsInDateRangeToSelection(rangeStart DateTime, rangeLength TimeSpan)
}

type IPhotoImportFindItemsResult2Vtbl struct {
	win32.IInspectableVtbl
	AddItemsInDateRangeToSelection uintptr
}

type IPhotoImportFindItemsResult2 struct {
	win32.IInspectable
}

func (this *IPhotoImportFindItemsResult2) Vtbl() *IPhotoImportFindItemsResult2Vtbl {
	return (*IPhotoImportFindItemsResult2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPhotoImportFindItemsResult2) AddItemsInDateRangeToSelection(rangeStart DateTime, rangeLength TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddItemsInDateRangeToSelection, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&rangeStart)), *(*uintptr)(unsafe.Pointer(&rangeLength)))
	_ = _hr
}

// E4D4F478-D419-4443-A84E-F06A850C0B00
var IID_IPhotoImportImportItemsResult = syscall.GUID{0xE4D4F478, 0xD419, 0x4443,
	[8]byte{0xA8, 0x4E, 0xF0, 0x6A, 0x85, 0x0C, 0x0B, 0x00}}

type IPhotoImportImportItemsResultInterface interface {
	win32.IInspectableInterface
	Get_Session() *IPhotoImportSession
	Get_HasSucceeded() bool
	Get_ImportedItems() *IVectorView[*IPhotoImportItem]
	Get_PhotosCount() uint32
	Get_PhotosSizeInBytes() uint64
	Get_VideosCount() uint32
	Get_VideosSizeInBytes() uint64
	Get_SidecarsCount() uint32
	Get_SidecarsSizeInBytes() uint64
	Get_SiblingsCount() uint32
	Get_SiblingsSizeInBytes() uint64
	Get_TotalCount() uint32
	Get_TotalSizeInBytes() uint64
	DeleteImportedItemsFromSourceAsync() *IAsyncOperationWithProgress[*IPhotoImportDeleteImportedItemsFromSourceResult, float64]
}

type IPhotoImportImportItemsResultVtbl struct {
	win32.IInspectableVtbl
	Get_Session                        uintptr
	Get_HasSucceeded                   uintptr
	Get_ImportedItems                  uintptr
	Get_PhotosCount                    uintptr
	Get_PhotosSizeInBytes              uintptr
	Get_VideosCount                    uintptr
	Get_VideosSizeInBytes              uintptr
	Get_SidecarsCount                  uintptr
	Get_SidecarsSizeInBytes            uintptr
	Get_SiblingsCount                  uintptr
	Get_SiblingsSizeInBytes            uintptr
	Get_TotalCount                     uintptr
	Get_TotalSizeInBytes               uintptr
	DeleteImportedItemsFromSourceAsync uintptr
}

type IPhotoImportImportItemsResult struct {
	win32.IInspectable
}

func (this *IPhotoImportImportItemsResult) Vtbl() *IPhotoImportImportItemsResultVtbl {
	return (*IPhotoImportImportItemsResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPhotoImportImportItemsResult) Get_Session() *IPhotoImportSession {
	var _result *IPhotoImportSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Session, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoImportImportItemsResult) Get_HasSucceeded() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasSucceeded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportImportItemsResult) Get_ImportedItems() *IVectorView[*IPhotoImportItem] {
	var _result *IVectorView[*IPhotoImportItem]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ImportedItems, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoImportImportItemsResult) Get_PhotosCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhotosCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportImportItemsResult) Get_PhotosSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhotosSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportImportItemsResult) Get_VideosCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideosCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportImportItemsResult) Get_VideosSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideosSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportImportItemsResult) Get_SidecarsCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SidecarsCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportImportItemsResult) Get_SidecarsSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SidecarsSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportImportItemsResult) Get_SiblingsCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SiblingsCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportImportItemsResult) Get_SiblingsSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SiblingsSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportImportItemsResult) Get_TotalCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TotalCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportImportItemsResult) Get_TotalSizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TotalSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportImportItemsResult) DeleteImportedItemsFromSourceAsync() *IAsyncOperationWithProgress[*IPhotoImportDeleteImportedItemsFromSourceResult, float64] {
	var _result *IAsyncOperationWithProgress[*IPhotoImportDeleteImportedItemsFromSourceResult, float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DeleteImportedItemsFromSourceAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A9D07E76-9BFC-43B8-B356-633B6A988C9E
var IID_IPhotoImportItem = syscall.GUID{0xA9D07E76, 0x9BFC, 0x43B8,
	[8]byte{0xB3, 0x56, 0x63, 0x3B, 0x6A, 0x98, 0x8C, 0x9E}}

type IPhotoImportItemInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Get_ItemKey() uint64
	Get_ContentType() PhotoImportContentType
	Get_SizeInBytes() uint64
	Get_Date() DateTime
	Get_Sibling() *IPhotoImportSidecar
	Get_Sidecars() *IVectorView[*IPhotoImportSidecar]
	Get_VideoSegments() *IVectorView[*IPhotoImportVideoSegment]
	Get_IsSelected() bool
	Put_IsSelected(value bool)
	Get_Thumbnail() *IRandomAccessStreamReference
	Get_ImportedFileNames() *IVectorView[string]
	Get_DeletedFileNames() *IVectorView[string]
}

type IPhotoImportItemVtbl struct {
	win32.IInspectableVtbl
	Get_Name              uintptr
	Get_ItemKey           uintptr
	Get_ContentType       uintptr
	Get_SizeInBytes       uintptr
	Get_Date              uintptr
	Get_Sibling           uintptr
	Get_Sidecars          uintptr
	Get_VideoSegments     uintptr
	Get_IsSelected        uintptr
	Put_IsSelected        uintptr
	Get_Thumbnail         uintptr
	Get_ImportedFileNames uintptr
	Get_DeletedFileNames  uintptr
}

type IPhotoImportItem struct {
	win32.IInspectable
}

func (this *IPhotoImportItem) Vtbl() *IPhotoImportItemVtbl {
	return (*IPhotoImportItemVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPhotoImportItem) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPhotoImportItem) Get_ItemKey() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ItemKey, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportItem) Get_ContentType() PhotoImportContentType {
	var _result PhotoImportContentType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportItem) Get_SizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportItem) Get_Date() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Date, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportItem) Get_Sibling() *IPhotoImportSidecar {
	var _result *IPhotoImportSidecar
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Sibling, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoImportItem) Get_Sidecars() *IVectorView[*IPhotoImportSidecar] {
	var _result *IVectorView[*IPhotoImportSidecar]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Sidecars, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoImportItem) Get_VideoSegments() *IVectorView[*IPhotoImportVideoSegment] {
	var _result *IVectorView[*IPhotoImportVideoSegment]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoSegments, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoImportItem) Get_IsSelected() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSelected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportItem) Put_IsSelected(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsSelected, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPhotoImportItem) Get_Thumbnail() *IRandomAccessStreamReference {
	var _result *IRandomAccessStreamReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Thumbnail, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoImportItem) Get_ImportedFileNames() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ImportedFileNames, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoImportItem) Get_DeletedFileNames() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeletedFileNames, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F1053505-F53B-46A3-9E30-3610791A9110
var IID_IPhotoImportItem2 = syscall.GUID{0xF1053505, 0xF53B, 0x46A3,
	[8]byte{0x9E, 0x30, 0x36, 0x10, 0x79, 0x1A, 0x91, 0x10}}

type IPhotoImportItem2Interface interface {
	win32.IInspectableInterface
	Get_Path() string
}

type IPhotoImportItem2Vtbl struct {
	win32.IInspectableVtbl
	Get_Path uintptr
}

type IPhotoImportItem2 struct {
	win32.IInspectable
}

func (this *IPhotoImportItem2) Vtbl() *IPhotoImportItem2Vtbl {
	return (*IPhotoImportItem2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPhotoImportItem2) Get_Path() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Path, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 42CB2FDD-7D68-47B5-BC7C-CEB73E0C77DC
var IID_IPhotoImportItemImportedEventArgs = syscall.GUID{0x42CB2FDD, 0x7D68, 0x47B5,
	[8]byte{0xBC, 0x7C, 0xCE, 0xB7, 0x3E, 0x0C, 0x77, 0xDC}}

type IPhotoImportItemImportedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ImportedItem() *IPhotoImportItem
}

type IPhotoImportItemImportedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ImportedItem uintptr
}

type IPhotoImportItemImportedEventArgs struct {
	win32.IInspectable
}

func (this *IPhotoImportItemImportedEventArgs) Vtbl() *IPhotoImportItemImportedEventArgsVtbl {
	return (*IPhotoImportItemImportedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPhotoImportItemImportedEventArgs) Get_ImportedItem() *IPhotoImportItem {
	var _result *IPhotoImportItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ImportedItem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2771903D-A046-4F06-9B9C-BFD662E83287
var IID_IPhotoImportManagerStatics = syscall.GUID{0x2771903D, 0xA046, 0x4F06,
	[8]byte{0x9B, 0x9C, 0xBF, 0xD6, 0x62, 0xE8, 0x32, 0x87}}

type IPhotoImportManagerStaticsInterface interface {
	win32.IInspectableInterface
	IsSupportedAsync() *IAsyncOperation[bool]
	FindAllSourcesAsync() *IAsyncOperation[*IVectorView[*IPhotoImportSource]]
	GetPendingOperations() *IVectorView[*IPhotoImportOperation]
}

type IPhotoImportManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	IsSupportedAsync     uintptr
	FindAllSourcesAsync  uintptr
	GetPendingOperations uintptr
}

type IPhotoImportManagerStatics struct {
	win32.IInspectable
}

func (this *IPhotoImportManagerStatics) Vtbl() *IPhotoImportManagerStaticsVtbl {
	return (*IPhotoImportManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPhotoImportManagerStatics) IsSupportedAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsSupportedAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoImportManagerStatics) FindAllSourcesAsync() *IAsyncOperation[*IVectorView[*IPhotoImportSource]] {
	var _result *IAsyncOperation[*IVectorView[*IPhotoImportSource]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllSourcesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoImportManagerStatics) GetPendingOperations() *IVectorView[*IPhotoImportOperation] {
	var _result *IVectorView[*IPhotoImportOperation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPendingOperations, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D9F797E4-A09A-4EE4-A4B1-20940277A5BE
var IID_IPhotoImportOperation = syscall.GUID{0xD9F797E4, 0xA09A, 0x4EE4,
	[8]byte{0xA4, 0xB1, 0x20, 0x94, 0x02, 0x77, 0xA5, 0xBE}}

type IPhotoImportOperationInterface interface {
	win32.IInspectableInterface
	Get_Stage() PhotoImportStage
	Get_Session() *IPhotoImportSession
	Get_ContinueFindingItemsAsync() *IAsyncOperationWithProgress[*IPhotoImportFindItemsResult, uint32]
	Get_ContinueImportingItemsAsync() *IAsyncOperationWithProgress[*IPhotoImportImportItemsResult, PhotoImportProgress]
	Get_ContinueDeletingImportedItemsFromSourceAsync() *IAsyncOperationWithProgress[*IPhotoImportDeleteImportedItemsFromSourceResult, float64]
}

type IPhotoImportOperationVtbl struct {
	win32.IInspectableVtbl
	Get_Stage                                        uintptr
	Get_Session                                      uintptr
	Get_ContinueFindingItemsAsync                    uintptr
	Get_ContinueImportingItemsAsync                  uintptr
	Get_ContinueDeletingImportedItemsFromSourceAsync uintptr
}

type IPhotoImportOperation struct {
	win32.IInspectable
}

func (this *IPhotoImportOperation) Vtbl() *IPhotoImportOperationVtbl {
	return (*IPhotoImportOperationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPhotoImportOperation) Get_Stage() PhotoImportStage {
	var _result PhotoImportStage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Stage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportOperation) Get_Session() *IPhotoImportSession {
	var _result *IPhotoImportSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Session, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoImportOperation) Get_ContinueFindingItemsAsync() *IAsyncOperationWithProgress[*IPhotoImportFindItemsResult, uint32] {
	var _result *IAsyncOperationWithProgress[*IPhotoImportFindItemsResult, uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContinueFindingItemsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoImportOperation) Get_ContinueImportingItemsAsync() *IAsyncOperationWithProgress[*IPhotoImportImportItemsResult, PhotoImportProgress] {
	var _result *IAsyncOperationWithProgress[*IPhotoImportImportItemsResult, PhotoImportProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContinueImportingItemsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoImportOperation) Get_ContinueDeletingImportedItemsFromSourceAsync() *IAsyncOperationWithProgress[*IPhotoImportDeleteImportedItemsFromSourceResult, float64] {
	var _result *IAsyncOperationWithProgress[*IPhotoImportDeleteImportedItemsFromSourceResult, float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContinueDeletingImportedItemsFromSourceAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 10461782-FA9D-4C30-8BC9-4D64911572D5
var IID_IPhotoImportSelectionChangedEventArgs = syscall.GUID{0x10461782, 0xFA9D, 0x4C30,
	[8]byte{0x8B, 0xC9, 0x4D, 0x64, 0x91, 0x15, 0x72, 0xD5}}

type IPhotoImportSelectionChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_IsSelectionEmpty() bool
}

type IPhotoImportSelectionChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_IsSelectionEmpty uintptr
}

type IPhotoImportSelectionChangedEventArgs struct {
	win32.IInspectable
}

func (this *IPhotoImportSelectionChangedEventArgs) Vtbl() *IPhotoImportSelectionChangedEventArgsVtbl {
	return (*IPhotoImportSelectionChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPhotoImportSelectionChangedEventArgs) Get_IsSelectionEmpty() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSelectionEmpty, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// AA63916E-ECDB-4EFE-94C6-5F5CAFE34CFB
var IID_IPhotoImportSession = syscall.GUID{0xAA63916E, 0xECDB, 0x4EFE,
	[8]byte{0x94, 0xC6, 0x5F, 0x5C, 0xAF, 0xE3, 0x4C, 0xFB}}

type IPhotoImportSessionInterface interface {
	win32.IInspectableInterface
	Get_Source() *IPhotoImportSource
	Get_SessionId() syscall.GUID
	Put_DestinationFolder(value *IStorageFolder)
	Get_DestinationFolder() *IStorageFolder
	Put_AppendSessionDateToDestinationFolder(value bool)
	Get_AppendSessionDateToDestinationFolder() bool
	Put_SubfolderCreationMode(value PhotoImportSubfolderCreationMode)
	Get_SubfolderCreationMode() PhotoImportSubfolderCreationMode
	Put_DestinationFileNamePrefix(value string)
	Get_DestinationFileNamePrefix() string
	FindItemsAsync(contentTypeFilter PhotoImportContentTypeFilter, itemSelectionMode PhotoImportItemSelectionMode) *IAsyncOperationWithProgress[*IPhotoImportFindItemsResult, uint32]
}

type IPhotoImportSessionVtbl struct {
	win32.IInspectableVtbl
	Get_Source                               uintptr
	Get_SessionId                            uintptr
	Put_DestinationFolder                    uintptr
	Get_DestinationFolder                    uintptr
	Put_AppendSessionDateToDestinationFolder uintptr
	Get_AppendSessionDateToDestinationFolder uintptr
	Put_SubfolderCreationMode                uintptr
	Get_SubfolderCreationMode                uintptr
	Put_DestinationFileNamePrefix            uintptr
	Get_DestinationFileNamePrefix            uintptr
	FindItemsAsync                           uintptr
}

type IPhotoImportSession struct {
	win32.IInspectable
}

func (this *IPhotoImportSession) Vtbl() *IPhotoImportSessionVtbl {
	return (*IPhotoImportSessionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPhotoImportSession) Get_Source() *IPhotoImportSource {
	var _result *IPhotoImportSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Source, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoImportSession) Get_SessionId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SessionId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportSession) Put_DestinationFolder(value *IStorageFolder) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DestinationFolder, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPhotoImportSession) Get_DestinationFolder() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DestinationFolder, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoImportSession) Put_AppendSessionDateToDestinationFolder(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AppendSessionDateToDestinationFolder, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPhotoImportSession) Get_AppendSessionDateToDestinationFolder() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppendSessionDateToDestinationFolder, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportSession) Put_SubfolderCreationMode(value PhotoImportSubfolderCreationMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SubfolderCreationMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPhotoImportSession) Get_SubfolderCreationMode() PhotoImportSubfolderCreationMode {
	var _result PhotoImportSubfolderCreationMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SubfolderCreationMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportSession) Put_DestinationFileNamePrefix(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DestinationFileNamePrefix, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPhotoImportSession) Get_DestinationFileNamePrefix() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DestinationFileNamePrefix, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPhotoImportSession) FindItemsAsync(contentTypeFilter PhotoImportContentTypeFilter, itemSelectionMode PhotoImportItemSelectionMode) *IAsyncOperationWithProgress[*IPhotoImportFindItemsResult, uint32] {
	var _result *IAsyncOperationWithProgress[*IPhotoImportFindItemsResult, uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindItemsAsync, uintptr(unsafe.Pointer(this)), uintptr(contentTypeFilter), uintptr(itemSelectionMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2A526710-3EC6-469D-A375-2B9F4785391E
var IID_IPhotoImportSession2 = syscall.GUID{0x2A526710, 0x3EC6, 0x469D,
	[8]byte{0xA3, 0x75, 0x2B, 0x9F, 0x47, 0x85, 0x39, 0x1E}}

type IPhotoImportSession2Interface interface {
	win32.IInspectableInterface
	Put_SubfolderDateFormat(value PhotoImportSubfolderDateFormat)
	Get_SubfolderDateFormat() PhotoImportSubfolderDateFormat
	Put_RememberDeselectedItems(value bool)
	Get_RememberDeselectedItems() bool
}

type IPhotoImportSession2Vtbl struct {
	win32.IInspectableVtbl
	Put_SubfolderDateFormat     uintptr
	Get_SubfolderDateFormat     uintptr
	Put_RememberDeselectedItems uintptr
	Get_RememberDeselectedItems uintptr
}

type IPhotoImportSession2 struct {
	win32.IInspectable
}

func (this *IPhotoImportSession2) Vtbl() *IPhotoImportSession2Vtbl {
	return (*IPhotoImportSession2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPhotoImportSession2) Put_SubfolderDateFormat(value PhotoImportSubfolderDateFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SubfolderDateFormat, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPhotoImportSession2) Get_SubfolderDateFormat() PhotoImportSubfolderDateFormat {
	var _result PhotoImportSubfolderDateFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SubfolderDateFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportSession2) Put_RememberDeselectedItems(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RememberDeselectedItems, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPhotoImportSession2) Get_RememberDeselectedItems() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RememberDeselectedItems, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 46D7D757-F802-44C7-9C98-7A71F4BC1486
var IID_IPhotoImportSidecar = syscall.GUID{0x46D7D757, 0xF802, 0x44C7,
	[8]byte{0x9C, 0x98, 0x7A, 0x71, 0xF4, 0xBC, 0x14, 0x86}}

type IPhotoImportSidecarInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Get_SizeInBytes() uint64
	Get_Date() DateTime
}

type IPhotoImportSidecarVtbl struct {
	win32.IInspectableVtbl
	Get_Name        uintptr
	Get_SizeInBytes uintptr
	Get_Date        uintptr
}

type IPhotoImportSidecar struct {
	win32.IInspectable
}

func (this *IPhotoImportSidecar) Vtbl() *IPhotoImportSidecarVtbl {
	return (*IPhotoImportSidecarVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPhotoImportSidecar) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPhotoImportSidecar) Get_SizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportSidecar) Get_Date() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Date, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 1F8EA35E-145B-4CD6-87F1-54965A982FEF
var IID_IPhotoImportSource = syscall.GUID{0x1F8EA35E, 0x145B, 0x4CD6,
	[8]byte{0x87, 0xF1, 0x54, 0x96, 0x5A, 0x98, 0x2F, 0xEF}}

type IPhotoImportSourceInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	Get_DisplayName() string
	Get_Description() string
	Get_Manufacturer() string
	Get_Model() string
	Get_SerialNumber() string
	Get_ConnectionProtocol() string
	Get_ConnectionTransport() PhotoImportConnectionTransport
	Get_Type() PhotoImportSourceType
	Get_PowerSource() PhotoImportPowerSource
	Get_BatteryLevelPercent() *IReference[uint32]
	Get_DateTime() *IReference[DateTime]
	Get_StorageMedia() *IVectorView[*IPhotoImportStorageMedium]
	Get_IsLocked() *IReference[bool]
	Get_IsMassStorage() bool
	Get_Thumbnail() *IRandomAccessStreamReference
	CreateImportSession() *IPhotoImportSession
}

type IPhotoImportSourceVtbl struct {
	win32.IInspectableVtbl
	Get_Id                  uintptr
	Get_DisplayName         uintptr
	Get_Description         uintptr
	Get_Manufacturer        uintptr
	Get_Model               uintptr
	Get_SerialNumber        uintptr
	Get_ConnectionProtocol  uintptr
	Get_ConnectionTransport uintptr
	Get_Type                uintptr
	Get_PowerSource         uintptr
	Get_BatteryLevelPercent uintptr
	Get_DateTime            uintptr
	Get_StorageMedia        uintptr
	Get_IsLocked            uintptr
	Get_IsMassStorage       uintptr
	Get_Thumbnail           uintptr
	CreateImportSession     uintptr
}

type IPhotoImportSource struct {
	win32.IInspectable
}

func (this *IPhotoImportSource) Vtbl() *IPhotoImportSourceVtbl {
	return (*IPhotoImportSourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPhotoImportSource) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPhotoImportSource) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPhotoImportSource) Get_Description() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Description, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPhotoImportSource) Get_Manufacturer() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Manufacturer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPhotoImportSource) Get_Model() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Model, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPhotoImportSource) Get_SerialNumber() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SerialNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPhotoImportSource) Get_ConnectionProtocol() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConnectionProtocol, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPhotoImportSource) Get_ConnectionTransport() PhotoImportConnectionTransport {
	var _result PhotoImportConnectionTransport
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConnectionTransport, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportSource) Get_Type() PhotoImportSourceType {
	var _result PhotoImportSourceType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Type, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportSource) Get_PowerSource() PhotoImportPowerSource {
	var _result PhotoImportPowerSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PowerSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportSource) Get_BatteryLevelPercent() *IReference[uint32] {
	var _result *IReference[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BatteryLevelPercent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoImportSource) Get_DateTime() *IReference[DateTime] {
	var _result *IReference[DateTime]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DateTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoImportSource) Get_StorageMedia() *IVectorView[*IPhotoImportStorageMedium] {
	var _result *IVectorView[*IPhotoImportStorageMedium]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StorageMedia, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoImportSource) Get_IsLocked() *IReference[bool] {
	var _result *IReference[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsLocked, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoImportSource) Get_IsMassStorage() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsMassStorage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportSource) Get_Thumbnail() *IRandomAccessStreamReference {
	var _result *IRandomAccessStreamReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Thumbnail, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoImportSource) CreateImportSession() *IPhotoImportSession {
	var _result *IPhotoImportSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateImportSession, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0528E586-32D8-467C-8CEE-23A1B2F43E85
var IID_IPhotoImportSourceStatics = syscall.GUID{0x0528E586, 0x32D8, 0x467C,
	[8]byte{0x8C, 0xEE, 0x23, 0xA1, 0xB2, 0xF4, 0x3E, 0x85}}

type IPhotoImportSourceStaticsInterface interface {
	win32.IInspectableInterface
	FromIdAsync(sourceId string) *IAsyncOperation[*IPhotoImportSource]
	FromFolderAsync(sourceRootFolder *IStorageFolder) *IAsyncOperation[*IPhotoImportSource]
}

type IPhotoImportSourceStaticsVtbl struct {
	win32.IInspectableVtbl
	FromIdAsync     uintptr
	FromFolderAsync uintptr
}

type IPhotoImportSourceStatics struct {
	win32.IInspectable
}

func (this *IPhotoImportSourceStatics) Vtbl() *IPhotoImportSourceStaticsVtbl {
	return (*IPhotoImportSourceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPhotoImportSourceStatics) FromIdAsync(sourceId string) *IAsyncOperation[*IPhotoImportSource] {
	var _result *IAsyncOperation[*IPhotoImportSource]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(sourceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoImportSourceStatics) FromFolderAsync(sourceRootFolder *IStorageFolder) *IAsyncOperation[*IPhotoImportSource] {
	var _result *IAsyncOperation[*IPhotoImportSource]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromFolderAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sourceRootFolder)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F2B9B093-FC85-487F-87C2-58D675D05B07
var IID_IPhotoImportStorageMedium = syscall.GUID{0xF2B9B093, 0xFC85, 0x487F,
	[8]byte{0x87, 0xC2, 0x58, 0xD6, 0x75, 0xD0, 0x5B, 0x07}}

type IPhotoImportStorageMediumInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Get_Description() string
	Get_SerialNumber() string
	Get_StorageMediumType() PhotoImportStorageMediumType
	Get_SupportedAccessMode() PhotoImportAccessMode
	Get_CapacityInBytes() uint64
	Get_AvailableSpaceInBytes() uint64
	Refresh()
}

type IPhotoImportStorageMediumVtbl struct {
	win32.IInspectableVtbl
	Get_Name                  uintptr
	Get_Description           uintptr
	Get_SerialNumber          uintptr
	Get_StorageMediumType     uintptr
	Get_SupportedAccessMode   uintptr
	Get_CapacityInBytes       uintptr
	Get_AvailableSpaceInBytes uintptr
	Refresh                   uintptr
}

type IPhotoImportStorageMedium struct {
	win32.IInspectable
}

func (this *IPhotoImportStorageMedium) Vtbl() *IPhotoImportStorageMediumVtbl {
	return (*IPhotoImportStorageMediumVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPhotoImportStorageMedium) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPhotoImportStorageMedium) Get_Description() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Description, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPhotoImportStorageMedium) Get_SerialNumber() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SerialNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPhotoImportStorageMedium) Get_StorageMediumType() PhotoImportStorageMediumType {
	var _result PhotoImportStorageMediumType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StorageMediumType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportStorageMedium) Get_SupportedAccessMode() PhotoImportAccessMode {
	var _result PhotoImportAccessMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedAccessMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportStorageMedium) Get_CapacityInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CapacityInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportStorageMedium) Get_AvailableSpaceInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AvailableSpaceInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportStorageMedium) Refresh() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Refresh, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 623C0289-321A-41D8-9166-8C62A333276C
var IID_IPhotoImportVideoSegment = syscall.GUID{0x623C0289, 0x321A, 0x41D8,
	[8]byte{0x91, 0x66, 0x8C, 0x62, 0xA3, 0x33, 0x27, 0x6C}}

type IPhotoImportVideoSegmentInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Get_SizeInBytes() uint64
	Get_Date() DateTime
	Get_Sibling() *IPhotoImportSidecar
	Get_Sidecars() *IVectorView[*IPhotoImportSidecar]
}

type IPhotoImportVideoSegmentVtbl struct {
	win32.IInspectableVtbl
	Get_Name        uintptr
	Get_SizeInBytes uintptr
	Get_Date        uintptr
	Get_Sibling     uintptr
	Get_Sidecars    uintptr
}

type IPhotoImportVideoSegment struct {
	win32.IInspectable
}

func (this *IPhotoImportVideoSegment) Vtbl() *IPhotoImportVideoSegmentVtbl {
	return (*IPhotoImportVideoSegmentVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPhotoImportVideoSegment) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPhotoImportVideoSegment) Get_SizeInBytes() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportVideoSegment) Get_Date() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Date, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoImportVideoSegment) Get_Sibling() *IPhotoImportSidecar {
	var _result *IPhotoImportSidecar
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Sibling, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPhotoImportVideoSegment) Get_Sidecars() *IVectorView[*IPhotoImportSidecar] {
	var _result *IVectorView[*IPhotoImportSidecar]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Sidecars, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type PhotoImportDeleteImportedItemsFromSourceResult struct {
	RtClass
	*IPhotoImportDeleteImportedItemsFromSourceResult
}

type PhotoImportFindItemsResult struct {
	RtClass
	*IPhotoImportFindItemsResult
}

type PhotoImportImportItemsResult struct {
	RtClass
	*IPhotoImportImportItemsResult
}

type PhotoImportItem struct {
	RtClass
	*IPhotoImportItem
}

type PhotoImportItemImportedEventArgs struct {
	RtClass
	*IPhotoImportItemImportedEventArgs
}

type PhotoImportManager struct {
	RtClass
}

func NewIPhotoImportManagerStatics() *IPhotoImportManagerStatics {
	var p *IPhotoImportManagerStatics
	hs := NewHStr("Windows.Media.Import.PhotoImportManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPhotoImportManagerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type PhotoImportOperation struct {
	RtClass
	*IPhotoImportOperation
}

type PhotoImportSelectionChangedEventArgs struct {
	RtClass
	*IPhotoImportSelectionChangedEventArgs
}

type PhotoImportSession struct {
	RtClass
	*IPhotoImportSession
}

type PhotoImportSidecar struct {
	RtClass
	*IPhotoImportSidecar
}

type PhotoImportSource struct {
	RtClass
	*IPhotoImportSource
}

func NewIPhotoImportSourceStatics() *IPhotoImportSourceStatics {
	var p *IPhotoImportSourceStatics
	hs := NewHStr("Windows.Media.Import.PhotoImportSource")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPhotoImportSourceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type PhotoImportStorageMedium struct {
	RtClass
	*IPhotoImportStorageMedium
}

type PhotoImportVideoSegment struct {
	RtClass
	*IPhotoImportVideoSegment
}
