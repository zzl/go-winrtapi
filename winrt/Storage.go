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
type ApplicationDataCreateDisposition int32

const (
	ApplicationDataCreateDisposition_Always   ApplicationDataCreateDisposition = 0
	ApplicationDataCreateDisposition_Existing ApplicationDataCreateDisposition = 1
)

// enum
type ApplicationDataLocality int32

const (
	ApplicationDataLocality_Local       ApplicationDataLocality = 0
	ApplicationDataLocality_Roaming     ApplicationDataLocality = 1
	ApplicationDataLocality_Temporary   ApplicationDataLocality = 2
	ApplicationDataLocality_LocalCache  ApplicationDataLocality = 3
	ApplicationDataLocality_SharedLocal ApplicationDataLocality = 4
)

// enum
type CreationCollisionOption int32

const (
	CreationCollisionOption_GenerateUniqueName CreationCollisionOption = 0
	CreationCollisionOption_ReplaceExisting    CreationCollisionOption = 1
	CreationCollisionOption_FailIfExists       CreationCollisionOption = 2
	CreationCollisionOption_OpenIfExists       CreationCollisionOption = 3
)

// enum
type FileAccessMode int32

const (
	FileAccessMode_Read      FileAccessMode = 0
	FileAccessMode_ReadWrite FileAccessMode = 1
)

// enum
// flags
type FileAttributes uint32

const (
	FileAttributes_Normal            FileAttributes = 0
	FileAttributes_ReadOnly          FileAttributes = 1
	FileAttributes_Directory         FileAttributes = 16
	FileAttributes_Archive           FileAttributes = 32
	FileAttributes_Temporary         FileAttributes = 256
	FileAttributes_LocallyIncomplete FileAttributes = 512
)

// enum
type KnownFolderId int32

const (
	KnownFolderId_AppCaptures        KnownFolderId = 0
	KnownFolderId_CameraRoll         KnownFolderId = 1
	KnownFolderId_DocumentsLibrary   KnownFolderId = 2
	KnownFolderId_HomeGroup          KnownFolderId = 3
	KnownFolderId_MediaServerDevices KnownFolderId = 4
	KnownFolderId_MusicLibrary       KnownFolderId = 5
	KnownFolderId_Objects3D          KnownFolderId = 6
	KnownFolderId_PicturesLibrary    KnownFolderId = 7
	KnownFolderId_Playlists          KnownFolderId = 8
	KnownFolderId_RecordedCalls      KnownFolderId = 9
	KnownFolderId_RemovableDevices   KnownFolderId = 10
	KnownFolderId_SavedPictures      KnownFolderId = 11
	KnownFolderId_Screenshots        KnownFolderId = 12
	KnownFolderId_VideosLibrary      KnownFolderId = 13
	KnownFolderId_AllAppMods         KnownFolderId = 14
	KnownFolderId_CurrentAppMods     KnownFolderId = 15
	KnownFolderId_DownloadsFolder    KnownFolderId = 16
)

// enum
type KnownFoldersAccessStatus int32

const (
	KnownFoldersAccessStatus_DeniedBySystem      KnownFoldersAccessStatus = 0
	KnownFoldersAccessStatus_NotDeclaredByApp    KnownFoldersAccessStatus = 1
	KnownFoldersAccessStatus_DeniedByUser        KnownFoldersAccessStatus = 2
	KnownFoldersAccessStatus_UserPromptRequired  KnownFoldersAccessStatus = 3
	KnownFoldersAccessStatus_Allowed             KnownFoldersAccessStatus = 4
	KnownFoldersAccessStatus_AllowedPerAppFolder KnownFoldersAccessStatus = 5
)

// enum
type KnownLibraryId int32

const (
	KnownLibraryId_Music     KnownLibraryId = 0
	KnownLibraryId_Pictures  KnownLibraryId = 1
	KnownLibraryId_Videos    KnownLibraryId = 2
	KnownLibraryId_Documents KnownLibraryId = 3
)

// enum
type NameCollisionOption int32

const (
	NameCollisionOption_GenerateUniqueName NameCollisionOption = 0
	NameCollisionOption_ReplaceExisting    NameCollisionOption = 1
	NameCollisionOption_FailIfExists       NameCollisionOption = 2
)

// enum
type StorageDeleteOption int32

const (
	StorageDeleteOption_Default         StorageDeleteOption = 0
	StorageDeleteOption_PermanentDelete StorageDeleteOption = 1
)

// enum
// flags
type StorageItemTypes uint32

const (
	StorageItemTypes_None   StorageItemTypes = 0
	StorageItemTypes_File   StorageItemTypes = 1
	StorageItemTypes_Folder StorageItemTypes = 2
)

// enum
type StorageLibraryChangeType int32

const (
	StorageLibraryChangeType_Created               StorageLibraryChangeType = 0
	StorageLibraryChangeType_Deleted               StorageLibraryChangeType = 1
	StorageLibraryChangeType_MovedOrRenamed        StorageLibraryChangeType = 2
	StorageLibraryChangeType_ContentsChanged       StorageLibraryChangeType = 3
	StorageLibraryChangeType_MovedOutOfLibrary     StorageLibraryChangeType = 4
	StorageLibraryChangeType_MovedIntoLibrary      StorageLibraryChangeType = 5
	StorageLibraryChangeType_ContentsReplaced      StorageLibraryChangeType = 6
	StorageLibraryChangeType_IndexingStatusChanged StorageLibraryChangeType = 7
	StorageLibraryChangeType_EncryptionChanged     StorageLibraryChangeType = 8
	StorageLibraryChangeType_ChangeTrackingLost    StorageLibraryChangeType = 9
)

// enum
// flags
type StorageOpenOptions uint32

const (
	StorageOpenOptions_None                   StorageOpenOptions = 0
	StorageOpenOptions_AllowOnlyReaders       StorageOpenOptions = 1
	StorageOpenOptions_AllowReadersAndWriters StorageOpenOptions = 2
)

// enum
type StreamedFileFailureMode int32

const (
	StreamedFileFailureMode_Failed               StreamedFileFailureMode = 0
	StreamedFileFailureMode_CurrentlyUnavailable StreamedFileFailureMode = 1
	StreamedFileFailureMode_Incomplete           StreamedFileFailureMode = 2
)

// func types

// A05791E6-CC9F-4687-ACAB-A364FD785463
type ApplicationDataSetVersionHandler func(setVersionRequest *ISetVersionRequest) com.Error

// FEF6A824-2FE1-4D07-A35B-B77C50B5F4CC
type StreamedFileDataRequestedHandler func(stream *IOutputStream) com.Error

// interfaces

// 7301D60A-79A2-48C9-9EC0-3FDA092F79E1
var IID_IAppDataPaths = syscall.GUID{0x7301D60A, 0x79A2, 0x48C9,
	[8]byte{0x9E, 0xC0, 0x3F, 0xDA, 0x09, 0x2F, 0x79, 0xE1}}

type IAppDataPathsInterface interface {
	win32.IInspectableInterface
	Get_Cookies() string
	Get_Desktop() string
	Get_Documents() string
	Get_Favorites() string
	Get_History() string
	Get_InternetCache() string
	Get_LocalAppData() string
	Get_ProgramData() string
	Get_RoamingAppData() string
}

type IAppDataPathsVtbl struct {
	win32.IInspectableVtbl
	Get_Cookies        uintptr
	Get_Desktop        uintptr
	Get_Documents      uintptr
	Get_Favorites      uintptr
	Get_History        uintptr
	Get_InternetCache  uintptr
	Get_LocalAppData   uintptr
	Get_ProgramData    uintptr
	Get_RoamingAppData uintptr
}

type IAppDataPaths struct {
	win32.IInspectable
}

func (this *IAppDataPaths) Vtbl() *IAppDataPathsVtbl {
	return (*IAppDataPathsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppDataPaths) Get_Cookies() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Cookies, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppDataPaths) Get_Desktop() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Desktop, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppDataPaths) Get_Documents() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Documents, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppDataPaths) Get_Favorites() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Favorites, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppDataPaths) Get_History() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_History, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppDataPaths) Get_InternetCache() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InternetCache, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppDataPaths) Get_LocalAppData() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocalAppData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppDataPaths) Get_ProgramData() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProgramData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppDataPaths) Get_RoamingAppData() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RoamingAppData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// D8EB2AFE-A9D9-4B14-B999-E3921379D903
var IID_IAppDataPathsStatics = syscall.GUID{0xD8EB2AFE, 0xA9D9, 0x4B14,
	[8]byte{0xB9, 0x99, 0xE3, 0x92, 0x13, 0x79, 0xD9, 0x03}}

type IAppDataPathsStaticsInterface interface {
	win32.IInspectableInterface
	GetForUser(user *IUser) *IAppDataPaths
	GetDefault() *IAppDataPaths
}

type IAppDataPathsStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForUser uintptr
	GetDefault uintptr
}

type IAppDataPathsStatics struct {
	win32.IInspectable
}

func (this *IAppDataPathsStatics) Vtbl() *IAppDataPathsStaticsVtbl {
	return (*IAppDataPathsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppDataPathsStatics) GetForUser(user *IUser) *IAppDataPaths {
	var _result *IAppDataPaths
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppDataPathsStatics) GetDefault() *IAppDataPaths {
	var _result *IAppDataPaths
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C3DA6FB7-B744-4B45-B0B8-223A0938D0DC
var IID_IApplicationData = syscall.GUID{0xC3DA6FB7, 0xB744, 0x4B45,
	[8]byte{0xB0, 0xB8, 0x22, 0x3A, 0x09, 0x38, 0xD0, 0xDC}}

type IApplicationDataInterface interface {
	win32.IInspectableInterface
	Get_Version() uint32
	SetVersionAsync(desiredVersion uint32, handler ApplicationDataSetVersionHandler) *IAsyncAction
	ClearAllAsync() *IAsyncAction
	ClearAsync(locality ApplicationDataLocality) *IAsyncAction
	Get_LocalSettings() *IApplicationDataContainer
	Get_RoamingSettings() *IApplicationDataContainer
	Get_LocalFolder() *IStorageFolder
	Get_RoamingFolder() *IStorageFolder
	Get_TemporaryFolder() *IStorageFolder
	Add_DataChanged(handler TypedEventHandler[*IApplicationData, interface{}]) EventRegistrationToken
	Remove_DataChanged(token EventRegistrationToken)
	SignalDataChanged()
	Get_RoamingStorageQuota() uint64
}

type IApplicationDataVtbl struct {
	win32.IInspectableVtbl
	Get_Version             uintptr
	SetVersionAsync         uintptr
	ClearAllAsync           uintptr
	ClearAsync              uintptr
	Get_LocalSettings       uintptr
	Get_RoamingSettings     uintptr
	Get_LocalFolder         uintptr
	Get_RoamingFolder       uintptr
	Get_TemporaryFolder     uintptr
	Add_DataChanged         uintptr
	Remove_DataChanged      uintptr
	SignalDataChanged       uintptr
	Get_RoamingStorageQuota uintptr
}

type IApplicationData struct {
	win32.IInspectable
}

func (this *IApplicationData) Vtbl() *IApplicationDataVtbl {
	return (*IApplicationDataVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationData) Get_Version() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Version, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApplicationData) SetVersionAsync(desiredVersion uint32, handler ApplicationDataSetVersionHandler) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetVersionAsync, uintptr(unsafe.Pointer(this)), uintptr(desiredVersion), uintptr(unsafe.Pointer(NewOneArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationData) ClearAllAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearAllAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationData) ClearAsync(locality ApplicationDataLocality) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearAsync, uintptr(unsafe.Pointer(this)), uintptr(locality), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationData) Get_LocalSettings() *IApplicationDataContainer {
	var _result *IApplicationDataContainer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocalSettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationData) Get_RoamingSettings() *IApplicationDataContainer {
	var _result *IApplicationDataContainer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RoamingSettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationData) Get_LocalFolder() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocalFolder, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationData) Get_RoamingFolder() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RoamingFolder, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationData) Get_TemporaryFolder() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TemporaryFolder, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationData) Add_DataChanged(handler TypedEventHandler[*IApplicationData, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DataChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApplicationData) Remove_DataChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DataChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IApplicationData) SignalDataChanged() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SignalDataChanged, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IApplicationData) Get_RoamingStorageQuota() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RoamingStorageQuota, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 9E65CD69-0BA3-4E32-BE29-B02DE6607638
var IID_IApplicationData2 = syscall.GUID{0x9E65CD69, 0x0BA3, 0x4E32,
	[8]byte{0xBE, 0x29, 0xB0, 0x2D, 0xE6, 0x60, 0x76, 0x38}}

type IApplicationData2Interface interface {
	win32.IInspectableInterface
	Get_LocalCacheFolder() *IStorageFolder
}

type IApplicationData2Vtbl struct {
	win32.IInspectableVtbl
	Get_LocalCacheFolder uintptr
}

type IApplicationData2 struct {
	win32.IInspectable
}

func (this *IApplicationData2) Vtbl() *IApplicationData2Vtbl {
	return (*IApplicationData2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationData2) Get_LocalCacheFolder() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocalCacheFolder, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DC222CF4-2772-4C1D-AA2C-C9F743ADE8D1
var IID_IApplicationData3 = syscall.GUID{0xDC222CF4, 0x2772, 0x4C1D,
	[8]byte{0xAA, 0x2C, 0xC9, 0xF7, 0x43, 0xAD, 0xE8, 0xD1}}

type IApplicationData3Interface interface {
	win32.IInspectableInterface
	GetPublisherCacheFolder(folderName string) *IStorageFolder
	ClearPublisherCacheFolderAsync(folderName string) *IAsyncAction
	Get_SharedLocalFolder() *IStorageFolder
}

type IApplicationData3Vtbl struct {
	win32.IInspectableVtbl
	GetPublisherCacheFolder        uintptr
	ClearPublisherCacheFolderAsync uintptr
	Get_SharedLocalFolder          uintptr
}

type IApplicationData3 struct {
	win32.IInspectable
}

func (this *IApplicationData3) Vtbl() *IApplicationData3Vtbl {
	return (*IApplicationData3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationData3) GetPublisherCacheFolder(folderName string) *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPublisherCacheFolder, uintptr(unsafe.Pointer(this)), NewHStr(folderName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationData3) ClearPublisherCacheFolderAsync(folderName string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearPublisherCacheFolderAsync, uintptr(unsafe.Pointer(this)), NewHStr(folderName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationData3) Get_SharedLocalFolder() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SharedLocalFolder, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C5AEFD1E-F467-40BA-8566-AB640A441E1D
var IID_IApplicationDataContainer = syscall.GUID{0xC5AEFD1E, 0xF467, 0x40BA,
	[8]byte{0x85, 0x66, 0xAB, 0x64, 0x0A, 0x44, 0x1E, 0x1D}}

type IApplicationDataContainerInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Get_Locality() ApplicationDataLocality
	Get_Values() *IPropertySet
	Get_Containers() *IMapView[string, *IApplicationDataContainer]
	CreateContainer(name string, disposition ApplicationDataCreateDisposition) *IApplicationDataContainer
	DeleteContainer(name string)
}

type IApplicationDataContainerVtbl struct {
	win32.IInspectableVtbl
	Get_Name        uintptr
	Get_Locality    uintptr
	Get_Values      uintptr
	Get_Containers  uintptr
	CreateContainer uintptr
	DeleteContainer uintptr
}

type IApplicationDataContainer struct {
	win32.IInspectable
}

func (this *IApplicationDataContainer) Vtbl() *IApplicationDataContainerVtbl {
	return (*IApplicationDataContainerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationDataContainer) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IApplicationDataContainer) Get_Locality() ApplicationDataLocality {
	var _result ApplicationDataLocality
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Locality, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IApplicationDataContainer) Get_Values() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Values, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationDataContainer) Get_Containers() *IMapView[string, *IApplicationDataContainer] {
	var _result *IMapView[string, *IApplicationDataContainer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Containers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationDataContainer) CreateContainer(name string, disposition ApplicationDataCreateDisposition) *IApplicationDataContainer {
	var _result *IApplicationDataContainer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateContainer, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(disposition), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IApplicationDataContainer) DeleteContainer(name string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DeleteContainer, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr)
	_ = _hr
}

// 5612147B-E843-45E3-94D8-06169E3C8E17
var IID_IApplicationDataStatics = syscall.GUID{0x5612147B, 0xE843, 0x45E3,
	[8]byte{0x94, 0xD8, 0x06, 0x16, 0x9E, 0x3C, 0x8E, 0x17}}

type IApplicationDataStaticsInterface interface {
	win32.IInspectableInterface
	Get_Current() *IApplicationData
}

type IApplicationDataStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Current uintptr
}

type IApplicationDataStatics struct {
	win32.IInspectable
}

func (this *IApplicationDataStatics) Vtbl() *IApplicationDataStaticsVtbl {
	return (*IApplicationDataStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationDataStatics) Get_Current() *IApplicationData {
	var _result *IApplicationData
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Current, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CD606211-CF49-40A4-A47C-C7F0DBBA8107
var IID_IApplicationDataStatics2 = syscall.GUID{0xCD606211, 0xCF49, 0x40A4,
	[8]byte{0xA4, 0x7C, 0xC7, 0xF0, 0xDB, 0xBA, 0x81, 0x07}}

type IApplicationDataStatics2Interface interface {
	win32.IInspectableInterface
	GetForUserAsync(user *IUser) *IAsyncOperation[*IApplicationData]
}

type IApplicationDataStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetForUserAsync uintptr
}

type IApplicationDataStatics2 struct {
	win32.IInspectable
}

func (this *IApplicationDataStatics2) Vtbl() *IApplicationDataStatics2Vtbl {
	return (*IApplicationDataStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IApplicationDataStatics2) GetForUserAsync(user *IUser) *IAsyncOperation[*IApplicationData] {
	var _result *IAsyncOperation[*IApplicationData]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForUserAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8FFC224A-E782-495D-B614-654C4F0B2370
var IID_ICachedFileManagerStatics = syscall.GUID{0x8FFC224A, 0xE782, 0x495D,
	[8]byte{0xB6, 0x14, 0x65, 0x4C, 0x4F, 0x0B, 0x23, 0x70}}

type ICachedFileManagerStaticsInterface interface {
	win32.IInspectableInterface
	DeferUpdates(file *IStorageFile)
	CompleteUpdatesAsync(file *IStorageFile) *IAsyncOperation[FileUpdateStatus]
}

type ICachedFileManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	DeferUpdates         uintptr
	CompleteUpdatesAsync uintptr
}

type ICachedFileManagerStatics struct {
	win32.IInspectable
}

func (this *ICachedFileManagerStatics) Vtbl() *ICachedFileManagerStaticsVtbl {
	return (*ICachedFileManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICachedFileManagerStatics) DeferUpdates(file *IStorageFile) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DeferUpdates, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)))
	_ = _hr
}

func (this *ICachedFileManagerStatics) CompleteUpdatesAsync(file *IStorageFile) *IAsyncOperation[FileUpdateStatus] {
	var _result *IAsyncOperation[FileUpdateStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CompleteUpdatesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 27862ED0-404E-47DF-A1E2-E37308BE7B37
var IID_IDownloadsFolderStatics = syscall.GUID{0x27862ED0, 0x404E, 0x47DF,
	[8]byte{0xA1, 0xE2, 0xE3, 0x73, 0x08, 0xBE, 0x7B, 0x37}}

type IDownloadsFolderStaticsInterface interface {
	win32.IInspectableInterface
	CreateFileAsync(desiredName string) *IAsyncOperation[*IStorageFile]
	CreateFolderAsync(desiredName string) *IAsyncOperation[*IStorageFolder]
	CreateFileWithCollisionOptionAsync(desiredName string, option CreationCollisionOption) *IAsyncOperation[*IStorageFile]
	CreateFolderWithCollisionOptionAsync(desiredName string, option CreationCollisionOption) *IAsyncOperation[*IStorageFolder]
}

type IDownloadsFolderStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateFileAsync                      uintptr
	CreateFolderAsync                    uintptr
	CreateFileWithCollisionOptionAsync   uintptr
	CreateFolderWithCollisionOptionAsync uintptr
}

type IDownloadsFolderStatics struct {
	win32.IInspectable
}

func (this *IDownloadsFolderStatics) Vtbl() *IDownloadsFolderStaticsVtbl {
	return (*IDownloadsFolderStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDownloadsFolderStatics) CreateFileAsync(desiredName string) *IAsyncOperation[*IStorageFile] {
	var _result *IAsyncOperation[*IStorageFile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFileAsync, uintptr(unsafe.Pointer(this)), NewHStr(desiredName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDownloadsFolderStatics) CreateFolderAsync(desiredName string) *IAsyncOperation[*IStorageFolder] {
	var _result *IAsyncOperation[*IStorageFolder]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFolderAsync, uintptr(unsafe.Pointer(this)), NewHStr(desiredName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDownloadsFolderStatics) CreateFileWithCollisionOptionAsync(desiredName string, option CreationCollisionOption) *IAsyncOperation[*IStorageFile] {
	var _result *IAsyncOperation[*IStorageFile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFileWithCollisionOptionAsync, uintptr(unsafe.Pointer(this)), NewHStr(desiredName).Ptr, uintptr(option), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDownloadsFolderStatics) CreateFolderWithCollisionOptionAsync(desiredName string, option CreationCollisionOption) *IAsyncOperation[*IStorageFolder] {
	var _result *IAsyncOperation[*IStorageFolder]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFolderWithCollisionOptionAsync, uintptr(unsafe.Pointer(this)), NewHStr(desiredName).Ptr, uintptr(option), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E93045BD-8EF8-4F8E-8D15-AC0E265F390D
var IID_IDownloadsFolderStatics2 = syscall.GUID{0xE93045BD, 0x8EF8, 0x4F8E,
	[8]byte{0x8D, 0x15, 0xAC, 0x0E, 0x26, 0x5F, 0x39, 0x0D}}

type IDownloadsFolderStatics2Interface interface {
	win32.IInspectableInterface
	CreateFileForUserAsync(user *IUser, desiredName string) *IAsyncOperation[*IStorageFile]
	CreateFolderForUserAsync(user *IUser, desiredName string) *IAsyncOperation[*IStorageFolder]
	CreateFileForUserWithCollisionOptionAsync(user *IUser, desiredName string, option CreationCollisionOption) *IAsyncOperation[*IStorageFile]
	CreateFolderForUserWithCollisionOptionAsync(user *IUser, desiredName string, option CreationCollisionOption) *IAsyncOperation[*IStorageFolder]
}

type IDownloadsFolderStatics2Vtbl struct {
	win32.IInspectableVtbl
	CreateFileForUserAsync                      uintptr
	CreateFolderForUserAsync                    uintptr
	CreateFileForUserWithCollisionOptionAsync   uintptr
	CreateFolderForUserWithCollisionOptionAsync uintptr
}

type IDownloadsFolderStatics2 struct {
	win32.IInspectable
}

func (this *IDownloadsFolderStatics2) Vtbl() *IDownloadsFolderStatics2Vtbl {
	return (*IDownloadsFolderStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDownloadsFolderStatics2) CreateFileForUserAsync(user *IUser, desiredName string) *IAsyncOperation[*IStorageFile] {
	var _result *IAsyncOperation[*IStorageFile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFileForUserAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), NewHStr(desiredName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDownloadsFolderStatics2) CreateFolderForUserAsync(user *IUser, desiredName string) *IAsyncOperation[*IStorageFolder] {
	var _result *IAsyncOperation[*IStorageFolder]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFolderForUserAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), NewHStr(desiredName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDownloadsFolderStatics2) CreateFileForUserWithCollisionOptionAsync(user *IUser, desiredName string, option CreationCollisionOption) *IAsyncOperation[*IStorageFile] {
	var _result *IAsyncOperation[*IStorageFile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFileForUserWithCollisionOptionAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), NewHStr(desiredName).Ptr, uintptr(option), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDownloadsFolderStatics2) CreateFolderForUserWithCollisionOptionAsync(user *IUser, desiredName string, option CreationCollisionOption) *IAsyncOperation[*IStorageFolder] {
	var _result *IAsyncOperation[*IStorageFolder]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFolderForUserWithCollisionOptionAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), NewHStr(desiredName).Ptr, uintptr(option), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 887411EB-7F54-4732-A5F0-5E43E3B8C2F5
var IID_IFileIOStatics = syscall.GUID{0x887411EB, 0x7F54, 0x4732,
	[8]byte{0xA5, 0xF0, 0x5E, 0x43, 0xE3, 0xB8, 0xC2, 0xF5}}

type IFileIOStaticsInterface interface {
	win32.IInspectableInterface
	ReadTextAsync(file *IStorageFile) *IAsyncOperation[string]
	ReadTextWithEncodingAsync(file *IStorageFile, encoding UnicodeEncoding) *IAsyncOperation[string]
	WriteTextAsync(file *IStorageFile, contents string) *IAsyncAction
	WriteTextWithEncodingAsync(file *IStorageFile, contents string, encoding UnicodeEncoding) *IAsyncAction
	AppendTextAsync(file *IStorageFile, contents string) *IAsyncAction
	AppendTextWithEncodingAsync(file *IStorageFile, contents string, encoding UnicodeEncoding) *IAsyncAction
	ReadLinesAsync(file *IStorageFile) *IAsyncOperation[*IVector[string]]
	ReadLinesWithEncodingAsync(file *IStorageFile, encoding UnicodeEncoding) *IAsyncOperation[*IVector[string]]
	WriteLinesAsync(file *IStorageFile, lines *IIterable[string]) *IAsyncAction
	WriteLinesWithEncodingAsync(file *IStorageFile, lines *IIterable[string], encoding UnicodeEncoding) *IAsyncAction
	AppendLinesAsync(file *IStorageFile, lines *IIterable[string]) *IAsyncAction
	AppendLinesWithEncodingAsync(file *IStorageFile, lines *IIterable[string], encoding UnicodeEncoding) *IAsyncAction
	ReadBufferAsync(file *IStorageFile) *IAsyncOperation[*IBuffer]
	WriteBufferAsync(file *IStorageFile, buffer *IBuffer) *IAsyncAction
	WriteBytesAsync(file *IStorageFile, bufferLength uint32, buffer *byte) *IAsyncAction
}

type IFileIOStaticsVtbl struct {
	win32.IInspectableVtbl
	ReadTextAsync                uintptr
	ReadTextWithEncodingAsync    uintptr
	WriteTextAsync               uintptr
	WriteTextWithEncodingAsync   uintptr
	AppendTextAsync              uintptr
	AppendTextWithEncodingAsync  uintptr
	ReadLinesAsync               uintptr
	ReadLinesWithEncodingAsync   uintptr
	WriteLinesAsync              uintptr
	WriteLinesWithEncodingAsync  uintptr
	AppendLinesAsync             uintptr
	AppendLinesWithEncodingAsync uintptr
	ReadBufferAsync              uintptr
	WriteBufferAsync             uintptr
	WriteBytesAsync              uintptr
}

type IFileIOStatics struct {
	win32.IInspectable
}

func (this *IFileIOStatics) Vtbl() *IFileIOStaticsVtbl {
	return (*IFileIOStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFileIOStatics) ReadTextAsync(file *IStorageFile) *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadTextAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileIOStatics) ReadTextWithEncodingAsync(file *IStorageFile, encoding UnicodeEncoding) *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadTextWithEncodingAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(encoding), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileIOStatics) WriteTextAsync(file *IStorageFile, contents string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteTextAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), NewHStr(contents).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileIOStatics) WriteTextWithEncodingAsync(file *IStorageFile, contents string, encoding UnicodeEncoding) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteTextWithEncodingAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), NewHStr(contents).Ptr, uintptr(encoding), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileIOStatics) AppendTextAsync(file *IStorageFile, contents string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AppendTextAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), NewHStr(contents).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileIOStatics) AppendTextWithEncodingAsync(file *IStorageFile, contents string, encoding UnicodeEncoding) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AppendTextWithEncodingAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), NewHStr(contents).Ptr, uintptr(encoding), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileIOStatics) ReadLinesAsync(file *IStorageFile) *IAsyncOperation[*IVector[string]] {
	var _result *IAsyncOperation[*IVector[string]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadLinesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileIOStatics) ReadLinesWithEncodingAsync(file *IStorageFile, encoding UnicodeEncoding) *IAsyncOperation[*IVector[string]] {
	var _result *IAsyncOperation[*IVector[string]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadLinesWithEncodingAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(encoding), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileIOStatics) WriteLinesAsync(file *IStorageFile, lines *IIterable[string]) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteLinesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(lines)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileIOStatics) WriteLinesWithEncodingAsync(file *IStorageFile, lines *IIterable[string], encoding UnicodeEncoding) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteLinesWithEncodingAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(lines)), uintptr(encoding), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileIOStatics) AppendLinesAsync(file *IStorageFile, lines *IIterable[string]) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AppendLinesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(lines)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileIOStatics) AppendLinesWithEncodingAsync(file *IStorageFile, lines *IIterable[string], encoding UnicodeEncoding) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AppendLinesWithEncodingAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(lines)), uintptr(encoding), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileIOStatics) ReadBufferAsync(file *IStorageFile) *IAsyncOperation[*IBuffer] {
	var _result *IAsyncOperation[*IBuffer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadBufferAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileIOStatics) WriteBufferAsync(file *IStorageFile, buffer *IBuffer) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteBufferAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileIOStatics) WriteBytesAsync(file *IStorageFile, bufferLength uint32, buffer *byte) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteBytesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(bufferLength), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5D115E66-27E8-492F-B8E5-2F90896CD4CD
var IID_IKnownFoldersCameraRollStatics = syscall.GUID{0x5D115E66, 0x27E8, 0x492F,
	[8]byte{0xB8, 0xE5, 0x2F, 0x90, 0x89, 0x6C, 0xD4, 0xCD}}

type IKnownFoldersCameraRollStaticsInterface interface {
	win32.IInspectableInterface
	Get_CameraRoll() *IStorageFolder
}

type IKnownFoldersCameraRollStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_CameraRoll uintptr
}

type IKnownFoldersCameraRollStatics struct {
	win32.IInspectable
}

func (this *IKnownFoldersCameraRollStatics) Vtbl() *IKnownFoldersCameraRollStaticsVtbl {
	return (*IKnownFoldersCameraRollStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKnownFoldersCameraRollStatics) Get_CameraRoll() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CameraRoll, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DAD5ECD6-306F-4D6A-B496-46BA8EB106CE
var IID_IKnownFoldersPlaylistsStatics = syscall.GUID{0xDAD5ECD6, 0x306F, 0x4D6A,
	[8]byte{0xB4, 0x96, 0x46, 0xBA, 0x8E, 0xB1, 0x06, 0xCE}}

type IKnownFoldersPlaylistsStaticsInterface interface {
	win32.IInspectableInterface
	Get_Playlists() *IStorageFolder
}

type IKnownFoldersPlaylistsStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Playlists uintptr
}

type IKnownFoldersPlaylistsStatics struct {
	win32.IInspectable
}

func (this *IKnownFoldersPlaylistsStatics) Vtbl() *IKnownFoldersPlaylistsStaticsVtbl {
	return (*IKnownFoldersPlaylistsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKnownFoldersPlaylistsStatics) Get_Playlists() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Playlists, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 055C93EA-253D-467C-B6CA-A97DA1E9A18D
var IID_IKnownFoldersSavedPicturesStatics = syscall.GUID{0x055C93EA, 0x253D, 0x467C,
	[8]byte{0xB6, 0xCA, 0xA9, 0x7D, 0xA1, 0xE9, 0xA1, 0x8D}}

type IKnownFoldersSavedPicturesStaticsInterface interface {
	win32.IInspectableInterface
	Get_SavedPictures() *IStorageFolder
}

type IKnownFoldersSavedPicturesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_SavedPictures uintptr
}

type IKnownFoldersSavedPicturesStatics struct {
	win32.IInspectable
}

func (this *IKnownFoldersSavedPicturesStatics) Vtbl() *IKnownFoldersSavedPicturesStaticsVtbl {
	return (*IKnownFoldersSavedPicturesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKnownFoldersSavedPicturesStatics) Get_SavedPictures() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SavedPictures, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5A2A7520-4802-452D-9AD9-4351ADA7EC35
var IID_IKnownFoldersStatics = syscall.GUID{0x5A2A7520, 0x4802, 0x452D,
	[8]byte{0x9A, 0xD9, 0x43, 0x51, 0xAD, 0xA7, 0xEC, 0x35}}

type IKnownFoldersStaticsInterface interface {
	win32.IInspectableInterface
	Get_MusicLibrary() *IStorageFolder
	Get_PicturesLibrary() *IStorageFolder
	Get_VideosLibrary() *IStorageFolder
	Get_DocumentsLibrary() *IStorageFolder
	Get_HomeGroup() *IStorageFolder
	Get_RemovableDevices() *IStorageFolder
	Get_MediaServerDevices() *IStorageFolder
}

type IKnownFoldersStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_MusicLibrary       uintptr
	Get_PicturesLibrary    uintptr
	Get_VideosLibrary      uintptr
	Get_DocumentsLibrary   uintptr
	Get_HomeGroup          uintptr
	Get_RemovableDevices   uintptr
	Get_MediaServerDevices uintptr
}

type IKnownFoldersStatics struct {
	win32.IInspectable
}

func (this *IKnownFoldersStatics) Vtbl() *IKnownFoldersStaticsVtbl {
	return (*IKnownFoldersStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKnownFoldersStatics) Get_MusicLibrary() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MusicLibrary, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKnownFoldersStatics) Get_PicturesLibrary() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PicturesLibrary, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKnownFoldersStatics) Get_VideosLibrary() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideosLibrary, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKnownFoldersStatics) Get_DocumentsLibrary() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DocumentsLibrary, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKnownFoldersStatics) Get_HomeGroup() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HomeGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKnownFoldersStatics) Get_RemovableDevices() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RemovableDevices, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKnownFoldersStatics) Get_MediaServerDevices() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaServerDevices, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 194BD0CD-CF6E-4D07-9D53-E9163A2536E9
var IID_IKnownFoldersStatics2 = syscall.GUID{0x194BD0CD, 0xCF6E, 0x4D07,
	[8]byte{0x9D, 0x53, 0xE9, 0x16, 0x3A, 0x25, 0x36, 0xE9}}

type IKnownFoldersStatics2Interface interface {
	win32.IInspectableInterface
	Get_Objects3D() *IStorageFolder
	Get_AppCaptures() *IStorageFolder
	Get_RecordedCalls() *IStorageFolder
}

type IKnownFoldersStatics2Vtbl struct {
	win32.IInspectableVtbl
	Get_Objects3D     uintptr
	Get_AppCaptures   uintptr
	Get_RecordedCalls uintptr
}

type IKnownFoldersStatics2 struct {
	win32.IInspectable
}

func (this *IKnownFoldersStatics2) Vtbl() *IKnownFoldersStatics2Vtbl {
	return (*IKnownFoldersStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKnownFoldersStatics2) Get_Objects3D() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Objects3D, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKnownFoldersStatics2) Get_AppCaptures() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppCaptures, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKnownFoldersStatics2) Get_RecordedCalls() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RecordedCalls, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C5194341-9742-4ED5-823D-FC1401148764
var IID_IKnownFoldersStatics3 = syscall.GUID{0xC5194341, 0x9742, 0x4ED5,
	[8]byte{0x82, 0x3D, 0xFC, 0x14, 0x01, 0x14, 0x87, 0x64}}

type IKnownFoldersStatics3Interface interface {
	win32.IInspectableInterface
	GetFolderForUserAsync(user *IUser, folderId KnownFolderId) *IAsyncOperation[*IStorageFolder]
}

type IKnownFoldersStatics3Vtbl struct {
	win32.IInspectableVtbl
	GetFolderForUserAsync uintptr
}

type IKnownFoldersStatics3 struct {
	win32.IInspectable
}

func (this *IKnownFoldersStatics3) Vtbl() *IKnownFoldersStatics3Vtbl {
	return (*IKnownFoldersStatics3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKnownFoldersStatics3) GetFolderForUserAsync(user *IUser, folderId KnownFolderId) *IAsyncOperation[*IStorageFolder] {
	var _result *IAsyncOperation[*IStorageFolder]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFolderForUserAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(folderId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1722E6BF-9FF9-4B21-BED5-90ECB13A192E
var IID_IKnownFoldersStatics4 = syscall.GUID{0x1722E6BF, 0x9FF9, 0x4B21,
	[8]byte{0xBE, 0xD5, 0x90, 0xEC, 0xB1, 0x3A, 0x19, 0x2E}}

type IKnownFoldersStatics4Interface interface {
	win32.IInspectableInterface
	RequestAccessAsync(folderId KnownFolderId) *IAsyncOperation[KnownFoldersAccessStatus]
	RequestAccessForUserAsync(user *IUser, folderId KnownFolderId) *IAsyncOperation[KnownFoldersAccessStatus]
	GetFolderAsync(folderId KnownFolderId) *IAsyncOperation[*IStorageFolder]
}

type IKnownFoldersStatics4Vtbl struct {
	win32.IInspectableVtbl
	RequestAccessAsync        uintptr
	RequestAccessForUserAsync uintptr
	GetFolderAsync            uintptr
}

type IKnownFoldersStatics4 struct {
	win32.IInspectable
}

func (this *IKnownFoldersStatics4) Vtbl() *IKnownFoldersStatics4Vtbl {
	return (*IKnownFoldersStatics4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKnownFoldersStatics4) RequestAccessAsync(folderId KnownFolderId) *IAsyncOperation[KnownFoldersAccessStatus] {
	var _result *IAsyncOperation[KnownFoldersAccessStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessAsync, uintptr(unsafe.Pointer(this)), uintptr(folderId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKnownFoldersStatics4) RequestAccessForUserAsync(user *IUser, folderId KnownFolderId) *IAsyncOperation[KnownFoldersAccessStatus] {
	var _result *IAsyncOperation[KnownFoldersAccessStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessForUserAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(folderId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKnownFoldersStatics4) GetFolderAsync(folderId KnownFolderId) *IAsyncOperation[*IStorageFolder] {
	var _result *IAsyncOperation[*IStorageFolder]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFolderAsync, uintptr(unsafe.Pointer(this)), uintptr(folderId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0F2F3758-8EC7-4381-922B-8F6C07D288F3
var IID_IPathIOStatics = syscall.GUID{0x0F2F3758, 0x8EC7, 0x4381,
	[8]byte{0x92, 0x2B, 0x8F, 0x6C, 0x07, 0xD2, 0x88, 0xF3}}

type IPathIOStaticsInterface interface {
	win32.IInspectableInterface
	ReadTextAsync(absolutePath string) *IAsyncOperation[string]
	ReadTextWithEncodingAsync(absolutePath string, encoding UnicodeEncoding) *IAsyncOperation[string]
	WriteTextAsync(absolutePath string, contents string) *IAsyncAction
	WriteTextWithEncodingAsync(absolutePath string, contents string, encoding UnicodeEncoding) *IAsyncAction
	AppendTextAsync(absolutePath string, contents string) *IAsyncAction
	AppendTextWithEncodingAsync(absolutePath string, contents string, encoding UnicodeEncoding) *IAsyncAction
	ReadLinesAsync(absolutePath string) *IAsyncOperation[*IVector[string]]
	ReadLinesWithEncodingAsync(absolutePath string, encoding UnicodeEncoding) *IAsyncOperation[*IVector[string]]
	WriteLinesAsync(absolutePath string, lines *IIterable[string]) *IAsyncAction
	WriteLinesWithEncodingAsync(absolutePath string, lines *IIterable[string], encoding UnicodeEncoding) *IAsyncAction
	AppendLinesAsync(absolutePath string, lines *IIterable[string]) *IAsyncAction
	AppendLinesWithEncodingAsync(absolutePath string, lines *IIterable[string], encoding UnicodeEncoding) *IAsyncAction
	ReadBufferAsync(absolutePath string) *IAsyncOperation[*IBuffer]
	WriteBufferAsync(absolutePath string, buffer *IBuffer) *IAsyncAction
	WriteBytesAsync(absolutePath string, bufferLength uint32, buffer *byte) *IAsyncAction
}

type IPathIOStaticsVtbl struct {
	win32.IInspectableVtbl
	ReadTextAsync                uintptr
	ReadTextWithEncodingAsync    uintptr
	WriteTextAsync               uintptr
	WriteTextWithEncodingAsync   uintptr
	AppendTextAsync              uintptr
	AppendTextWithEncodingAsync  uintptr
	ReadLinesAsync               uintptr
	ReadLinesWithEncodingAsync   uintptr
	WriteLinesAsync              uintptr
	WriteLinesWithEncodingAsync  uintptr
	AppendLinesAsync             uintptr
	AppendLinesWithEncodingAsync uintptr
	ReadBufferAsync              uintptr
	WriteBufferAsync             uintptr
	WriteBytesAsync              uintptr
}

type IPathIOStatics struct {
	win32.IInspectable
}

func (this *IPathIOStatics) Vtbl() *IPathIOStaticsVtbl {
	return (*IPathIOStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPathIOStatics) ReadTextAsync(absolutePath string) *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadTextAsync, uintptr(unsafe.Pointer(this)), NewHStr(absolutePath).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPathIOStatics) ReadTextWithEncodingAsync(absolutePath string, encoding UnicodeEncoding) *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadTextWithEncodingAsync, uintptr(unsafe.Pointer(this)), NewHStr(absolutePath).Ptr, uintptr(encoding), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPathIOStatics) WriteTextAsync(absolutePath string, contents string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteTextAsync, uintptr(unsafe.Pointer(this)), NewHStr(absolutePath).Ptr, NewHStr(contents).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPathIOStatics) WriteTextWithEncodingAsync(absolutePath string, contents string, encoding UnicodeEncoding) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteTextWithEncodingAsync, uintptr(unsafe.Pointer(this)), NewHStr(absolutePath).Ptr, NewHStr(contents).Ptr, uintptr(encoding), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPathIOStatics) AppendTextAsync(absolutePath string, contents string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AppendTextAsync, uintptr(unsafe.Pointer(this)), NewHStr(absolutePath).Ptr, NewHStr(contents).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPathIOStatics) AppendTextWithEncodingAsync(absolutePath string, contents string, encoding UnicodeEncoding) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AppendTextWithEncodingAsync, uintptr(unsafe.Pointer(this)), NewHStr(absolutePath).Ptr, NewHStr(contents).Ptr, uintptr(encoding), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPathIOStatics) ReadLinesAsync(absolutePath string) *IAsyncOperation[*IVector[string]] {
	var _result *IAsyncOperation[*IVector[string]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadLinesAsync, uintptr(unsafe.Pointer(this)), NewHStr(absolutePath).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPathIOStatics) ReadLinesWithEncodingAsync(absolutePath string, encoding UnicodeEncoding) *IAsyncOperation[*IVector[string]] {
	var _result *IAsyncOperation[*IVector[string]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadLinesWithEncodingAsync, uintptr(unsafe.Pointer(this)), NewHStr(absolutePath).Ptr, uintptr(encoding), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPathIOStatics) WriteLinesAsync(absolutePath string, lines *IIterable[string]) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteLinesAsync, uintptr(unsafe.Pointer(this)), NewHStr(absolutePath).Ptr, uintptr(unsafe.Pointer(lines)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPathIOStatics) WriteLinesWithEncodingAsync(absolutePath string, lines *IIterable[string], encoding UnicodeEncoding) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteLinesWithEncodingAsync, uintptr(unsafe.Pointer(this)), NewHStr(absolutePath).Ptr, uintptr(unsafe.Pointer(lines)), uintptr(encoding), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPathIOStatics) AppendLinesAsync(absolutePath string, lines *IIterable[string]) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AppendLinesAsync, uintptr(unsafe.Pointer(this)), NewHStr(absolutePath).Ptr, uintptr(unsafe.Pointer(lines)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPathIOStatics) AppendLinesWithEncodingAsync(absolutePath string, lines *IIterable[string], encoding UnicodeEncoding) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AppendLinesWithEncodingAsync, uintptr(unsafe.Pointer(this)), NewHStr(absolutePath).Ptr, uintptr(unsafe.Pointer(lines)), uintptr(encoding), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPathIOStatics) ReadBufferAsync(absolutePath string) *IAsyncOperation[*IBuffer] {
	var _result *IAsyncOperation[*IBuffer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadBufferAsync, uintptr(unsafe.Pointer(this)), NewHStr(absolutePath).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPathIOStatics) WriteBufferAsync(absolutePath string, buffer *IBuffer) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteBufferAsync, uintptr(unsafe.Pointer(this)), NewHStr(absolutePath).Ptr, uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPathIOStatics) WriteBytesAsync(absolutePath string, bufferLength uint32, buffer *byte) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteBytesAsync, uintptr(unsafe.Pointer(this)), NewHStr(absolutePath).Ptr, uintptr(bufferLength), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 033508A2-781A-437A-B078-3F32BADCFE47
var IID_ISetVersionDeferral = syscall.GUID{0x033508A2, 0x781A, 0x437A,
	[8]byte{0xB0, 0x78, 0x3F, 0x32, 0xBA, 0xDC, 0xFE, 0x47}}

type ISetVersionDeferralInterface interface {
	win32.IInspectableInterface
	Complete()
}

type ISetVersionDeferralVtbl struct {
	win32.IInspectableVtbl
	Complete uintptr
}

type ISetVersionDeferral struct {
	win32.IInspectable
}

func (this *ISetVersionDeferral) Vtbl() *ISetVersionDeferralVtbl {
	return (*ISetVersionDeferralVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISetVersionDeferral) Complete() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Complete, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// B9C76B9B-1056-4E69-8330-162619956F9B
var IID_ISetVersionRequest = syscall.GUID{0xB9C76B9B, 0x1056, 0x4E69,
	[8]byte{0x83, 0x30, 0x16, 0x26, 0x19, 0x95, 0x6F, 0x9B}}

type ISetVersionRequestInterface interface {
	win32.IInspectableInterface
	Get_CurrentVersion() uint32
	Get_DesiredVersion() uint32
	GetDeferral() *ISetVersionDeferral
}

type ISetVersionRequestVtbl struct {
	win32.IInspectableVtbl
	Get_CurrentVersion uintptr
	Get_DesiredVersion uintptr
	GetDeferral        uintptr
}

type ISetVersionRequest struct {
	win32.IInspectable
}

func (this *ISetVersionRequest) Vtbl() *ISetVersionRequestVtbl {
	return (*ISetVersionRequestVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISetVersionRequest) Get_CurrentVersion() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentVersion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISetVersionRequest) Get_DesiredVersion() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesiredVersion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISetVersionRequest) GetDeferral() *ISetVersionDeferral {
	var _result *ISetVersionDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FA3F6186-4214-428C-A64C-14C9AC7315EA
var IID_IStorageFile = syscall.GUID{0xFA3F6186, 0x4214, 0x428C,
	[8]byte{0xA6, 0x4C, 0x14, 0xC9, 0xAC, 0x73, 0x15, 0xEA}}

type IStorageFileInterface interface {
	win32.IInspectableInterface
	Get_FileType() string
	Get_ContentType() string
	OpenAsync(accessMode FileAccessMode) *IAsyncOperation[*IRandomAccessStream]
	OpenTransactedWriteAsync() *IAsyncOperation[*IStorageStreamTransaction]
	CopyOverloadDefaultNameAndOptions(destinationFolder *IStorageFolder) *IAsyncOperation[*IStorageFile]
	CopyOverloadDefaultOptions(destinationFolder *IStorageFolder, desiredNewName string) *IAsyncOperation[*IStorageFile]
	CopyOverload(destinationFolder *IStorageFolder, desiredNewName string, option NameCollisionOption) *IAsyncOperation[*IStorageFile]
	CopyAndReplaceAsync(fileToReplace *IStorageFile) *IAsyncAction
	MoveOverloadDefaultNameAndOptions(destinationFolder *IStorageFolder) *IAsyncAction
	MoveOverloadDefaultOptions(destinationFolder *IStorageFolder, desiredNewName string) *IAsyncAction
	MoveOverload(destinationFolder *IStorageFolder, desiredNewName string, option NameCollisionOption) *IAsyncAction
	MoveAndReplaceAsync(fileToReplace *IStorageFile) *IAsyncAction
}

type IStorageFileVtbl struct {
	win32.IInspectableVtbl
	Get_FileType                      uintptr
	Get_ContentType                   uintptr
	OpenAsync                         uintptr
	OpenTransactedWriteAsync          uintptr
	CopyOverloadDefaultNameAndOptions uintptr
	CopyOverloadDefaultOptions        uintptr
	CopyOverload                      uintptr
	CopyAndReplaceAsync               uintptr
	MoveOverloadDefaultNameAndOptions uintptr
	MoveOverloadDefaultOptions        uintptr
	MoveOverload                      uintptr
	MoveAndReplaceAsync               uintptr
}

type IStorageFile struct {
	win32.IInspectable
}

func (this *IStorageFile) Vtbl() *IStorageFileVtbl {
	return (*IStorageFileVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageFile) Get_FileType() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FileType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStorageFile) Get_ContentType() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStorageFile) OpenAsync(accessMode FileAccessMode) *IAsyncOperation[*IRandomAccessStream] {
	var _result *IAsyncOperation[*IRandomAccessStream]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenAsync, uintptr(unsafe.Pointer(this)), uintptr(accessMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageFile) OpenTransactedWriteAsync() *IAsyncOperation[*IStorageStreamTransaction] {
	var _result *IAsyncOperation[*IStorageStreamTransaction]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenTransactedWriteAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageFile) CopyOverloadDefaultNameAndOptions(destinationFolder *IStorageFolder) *IAsyncOperation[*IStorageFile] {
	var _result *IAsyncOperation[*IStorageFile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CopyOverloadDefaultNameAndOptions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(destinationFolder)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageFile) CopyOverloadDefaultOptions(destinationFolder *IStorageFolder, desiredNewName string) *IAsyncOperation[*IStorageFile] {
	var _result *IAsyncOperation[*IStorageFile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CopyOverloadDefaultOptions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(destinationFolder)), NewHStr(desiredNewName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageFile) CopyOverload(destinationFolder *IStorageFolder, desiredNewName string, option NameCollisionOption) *IAsyncOperation[*IStorageFile] {
	var _result *IAsyncOperation[*IStorageFile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CopyOverload, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(destinationFolder)), NewHStr(desiredNewName).Ptr, uintptr(option), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageFile) CopyAndReplaceAsync(fileToReplace *IStorageFile) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CopyAndReplaceAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fileToReplace)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageFile) MoveOverloadDefaultNameAndOptions(destinationFolder *IStorageFolder) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MoveOverloadDefaultNameAndOptions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(destinationFolder)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageFile) MoveOverloadDefaultOptions(destinationFolder *IStorageFolder, desiredNewName string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MoveOverloadDefaultOptions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(destinationFolder)), NewHStr(desiredNewName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageFile) MoveOverload(destinationFolder *IStorageFolder, desiredNewName string, option NameCollisionOption) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MoveOverload, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(destinationFolder)), NewHStr(desiredNewName).Ptr, uintptr(option), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageFile) MoveAndReplaceAsync(fileToReplace *IStorageFile) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MoveAndReplaceAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fileToReplace)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 954E4BCF-0A77-42FB-B777-C2ED58A52E44
var IID_IStorageFile2 = syscall.GUID{0x954E4BCF, 0x0A77, 0x42FB,
	[8]byte{0xB7, 0x77, 0xC2, 0xED, 0x58, 0xA5, 0x2E, 0x44}}

type IStorageFile2Interface interface {
	win32.IInspectableInterface
	OpenWithOptionsAsync(accessMode FileAccessMode, options StorageOpenOptions) *IAsyncOperation[*IRandomAccessStream]
	OpenTransactedWriteWithOptionsAsync(options StorageOpenOptions) *IAsyncOperation[*IStorageStreamTransaction]
}

type IStorageFile2Vtbl struct {
	win32.IInspectableVtbl
	OpenWithOptionsAsync                uintptr
	OpenTransactedWriteWithOptionsAsync uintptr
}

type IStorageFile2 struct {
	win32.IInspectable
}

func (this *IStorageFile2) Vtbl() *IStorageFile2Vtbl {
	return (*IStorageFile2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageFile2) OpenWithOptionsAsync(accessMode FileAccessMode, options StorageOpenOptions) *IAsyncOperation[*IRandomAccessStream] {
	var _result *IAsyncOperation[*IRandomAccessStream]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenWithOptionsAsync, uintptr(unsafe.Pointer(this)), uintptr(accessMode), uintptr(options), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageFile2) OpenTransactedWriteWithOptionsAsync(options StorageOpenOptions) *IAsyncOperation[*IStorageStreamTransaction] {
	var _result *IAsyncOperation[*IStorageStreamTransaction]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenTransactedWriteWithOptionsAsync, uintptr(unsafe.Pointer(this)), uintptr(options), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AFCBBE9B-582B-4133-9648-E44CA46EE491
var IID_IStorageFilePropertiesWithAvailability = syscall.GUID{0xAFCBBE9B, 0x582B, 0x4133,
	[8]byte{0x96, 0x48, 0xE4, 0x4C, 0xA4, 0x6E, 0xE4, 0x91}}

type IStorageFilePropertiesWithAvailabilityInterface interface {
	win32.IInspectableInterface
	Get_IsAvailable() bool
}

type IStorageFilePropertiesWithAvailabilityVtbl struct {
	win32.IInspectableVtbl
	Get_IsAvailable uintptr
}

type IStorageFilePropertiesWithAvailability struct {
	win32.IInspectable
}

func (this *IStorageFilePropertiesWithAvailability) Vtbl() *IStorageFilePropertiesWithAvailabilityVtbl {
	return (*IStorageFilePropertiesWithAvailabilityVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageFilePropertiesWithAvailability) Get_IsAvailable() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsAvailable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5984C710-DAF2-43C8-8BB4-A4D3EACFD03F
var IID_IStorageFileStatics = syscall.GUID{0x5984C710, 0xDAF2, 0x43C8,
	[8]byte{0x8B, 0xB4, 0xA4, 0xD3, 0xEA, 0xCF, 0xD0, 0x3F}}

type IStorageFileStaticsInterface interface {
	win32.IInspectableInterface
	GetFileFromPathAsync(path string) *IAsyncOperation[*IStorageFile]
	GetFileFromApplicationUriAsync(uri *IUriRuntimeClass) *IAsyncOperation[*IStorageFile]
	CreateStreamedFileAsync(displayNameWithExtension string, dataRequested StreamedFileDataRequestedHandler, thumbnail *IRandomAccessStreamReference) *IAsyncOperation[*IStorageFile]
	ReplaceWithStreamedFileAsync(fileToReplace *IStorageFile, dataRequested StreamedFileDataRequestedHandler, thumbnail *IRandomAccessStreamReference) *IAsyncOperation[*IStorageFile]
	CreateStreamedFileFromUriAsync(displayNameWithExtension string, uri *IUriRuntimeClass, thumbnail *IRandomAccessStreamReference) *IAsyncOperation[*IStorageFile]
	ReplaceWithStreamedFileFromUriAsync(fileToReplace *IStorageFile, uri *IUriRuntimeClass, thumbnail *IRandomAccessStreamReference) *IAsyncOperation[*IStorageFile]
}

type IStorageFileStaticsVtbl struct {
	win32.IInspectableVtbl
	GetFileFromPathAsync                uintptr
	GetFileFromApplicationUriAsync      uintptr
	CreateStreamedFileAsync             uintptr
	ReplaceWithStreamedFileAsync        uintptr
	CreateStreamedFileFromUriAsync      uintptr
	ReplaceWithStreamedFileFromUriAsync uintptr
}

type IStorageFileStatics struct {
	win32.IInspectable
}

func (this *IStorageFileStatics) Vtbl() *IStorageFileStaticsVtbl {
	return (*IStorageFileStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageFileStatics) GetFileFromPathAsync(path string) *IAsyncOperation[*IStorageFile] {
	var _result *IAsyncOperation[*IStorageFile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFileFromPathAsync, uintptr(unsafe.Pointer(this)), NewHStr(path).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageFileStatics) GetFileFromApplicationUriAsync(uri *IUriRuntimeClass) *IAsyncOperation[*IStorageFile] {
	var _result *IAsyncOperation[*IStorageFile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFileFromApplicationUriAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageFileStatics) CreateStreamedFileAsync(displayNameWithExtension string, dataRequested StreamedFileDataRequestedHandler, thumbnail *IRandomAccessStreamReference) *IAsyncOperation[*IStorageFile] {
	var _result *IAsyncOperation[*IStorageFile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateStreamedFileAsync, uintptr(unsafe.Pointer(this)), NewHStr(displayNameWithExtension).Ptr, uintptr(unsafe.Pointer(NewOneArgFuncDelegate(dataRequested))), uintptr(unsafe.Pointer(thumbnail)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageFileStatics) ReplaceWithStreamedFileAsync(fileToReplace *IStorageFile, dataRequested StreamedFileDataRequestedHandler, thumbnail *IRandomAccessStreamReference) *IAsyncOperation[*IStorageFile] {
	var _result *IAsyncOperation[*IStorageFile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReplaceWithStreamedFileAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fileToReplace)), uintptr(unsafe.Pointer(NewOneArgFuncDelegate(dataRequested))), uintptr(unsafe.Pointer(thumbnail)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageFileStatics) CreateStreamedFileFromUriAsync(displayNameWithExtension string, uri *IUriRuntimeClass, thumbnail *IRandomAccessStreamReference) *IAsyncOperation[*IStorageFile] {
	var _result *IAsyncOperation[*IStorageFile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateStreamedFileFromUriAsync, uintptr(unsafe.Pointer(this)), NewHStr(displayNameWithExtension).Ptr, uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(thumbnail)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageFileStatics) ReplaceWithStreamedFileFromUriAsync(fileToReplace *IStorageFile, uri *IUriRuntimeClass, thumbnail *IRandomAccessStreamReference) *IAsyncOperation[*IStorageFile] {
	var _result *IAsyncOperation[*IStorageFile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReplaceWithStreamedFileFromUriAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fileToReplace)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(thumbnail)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5C76A781-212E-4AF9-8F04-740CAE108974
var IID_IStorageFileStatics2 = syscall.GUID{0x5C76A781, 0x212E, 0x4AF9,
	[8]byte{0x8F, 0x04, 0x74, 0x0C, 0xAE, 0x10, 0x89, 0x74}}

type IStorageFileStatics2Interface interface {
	win32.IInspectableInterface
	GetFileFromPathForUserAsync(user *IUser, path string) *IAsyncOperation[*IStorageFile]
}

type IStorageFileStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetFileFromPathForUserAsync uintptr
}

type IStorageFileStatics2 struct {
	win32.IInspectable
}

func (this *IStorageFileStatics2) Vtbl() *IStorageFileStatics2Vtbl {
	return (*IStorageFileStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageFileStatics2) GetFileFromPathForUserAsync(user *IUser, path string) *IAsyncOperation[*IStorageFile] {
	var _result *IAsyncOperation[*IStorageFile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFileFromPathForUserAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), NewHStr(path).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 72D1CB78-B3EF-4F75-A80B-6FD9DAE2944B
var IID_IStorageFolder = syscall.GUID{0x72D1CB78, 0xB3EF, 0x4F75,
	[8]byte{0xA8, 0x0B, 0x6F, 0xD9, 0xDA, 0xE2, 0x94, 0x4B}}

type IStorageFolderInterface interface {
	win32.IInspectableInterface
	CreateFileAsyncOverloadDefaultOptions(desiredName string) *IAsyncOperation[*IStorageFile]
	CreateFileAsync(desiredName string, options CreationCollisionOption) *IAsyncOperation[*IStorageFile]
	CreateFolderAsyncOverloadDefaultOptions(desiredName string) *IAsyncOperation[*IStorageFolder]
	CreateFolderAsync(desiredName string, options CreationCollisionOption) *IAsyncOperation[*IStorageFolder]
	GetFileAsync(name string) *IAsyncOperation[*IStorageFile]
	GetFolderAsync(name string) *IAsyncOperation[*IStorageFolder]
	GetItemAsync(name string) *IAsyncOperation[*IStorageItem]
	GetFilesAsyncOverloadDefaultOptionsStartAndCount() *IAsyncOperation[*IVectorView[*IStorageFile]]
	GetFoldersAsyncOverloadDefaultOptionsStartAndCount() *IAsyncOperation[*IVectorView[*IStorageFolder]]
	GetItemsAsyncOverloadDefaultStartAndCount() *IAsyncOperation[*IVectorView[*IStorageItem]]
}

type IStorageFolderVtbl struct {
	win32.IInspectableVtbl
	CreateFileAsyncOverloadDefaultOptions              uintptr
	CreateFileAsync                                    uintptr
	CreateFolderAsyncOverloadDefaultOptions            uintptr
	CreateFolderAsync                                  uintptr
	GetFileAsync                                       uintptr
	GetFolderAsync                                     uintptr
	GetItemAsync                                       uintptr
	GetFilesAsyncOverloadDefaultOptionsStartAndCount   uintptr
	GetFoldersAsyncOverloadDefaultOptionsStartAndCount uintptr
	GetItemsAsyncOverloadDefaultStartAndCount          uintptr
}

type IStorageFolder struct {
	win32.IInspectable
}

func (this *IStorageFolder) Vtbl() *IStorageFolderVtbl {
	return (*IStorageFolderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageFolder) CreateFileAsyncOverloadDefaultOptions(desiredName string) *IAsyncOperation[*IStorageFile] {
	var _result *IAsyncOperation[*IStorageFile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFileAsyncOverloadDefaultOptions, uintptr(unsafe.Pointer(this)), NewHStr(desiredName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageFolder) CreateFileAsync(desiredName string, options CreationCollisionOption) *IAsyncOperation[*IStorageFile] {
	var _result *IAsyncOperation[*IStorageFile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFileAsync, uintptr(unsafe.Pointer(this)), NewHStr(desiredName).Ptr, uintptr(options), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageFolder) CreateFolderAsyncOverloadDefaultOptions(desiredName string) *IAsyncOperation[*IStorageFolder] {
	var _result *IAsyncOperation[*IStorageFolder]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFolderAsyncOverloadDefaultOptions, uintptr(unsafe.Pointer(this)), NewHStr(desiredName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageFolder) CreateFolderAsync(desiredName string, options CreationCollisionOption) *IAsyncOperation[*IStorageFolder] {
	var _result *IAsyncOperation[*IStorageFolder]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFolderAsync, uintptr(unsafe.Pointer(this)), NewHStr(desiredName).Ptr, uintptr(options), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageFolder) GetFileAsync(name string) *IAsyncOperation[*IStorageFile] {
	var _result *IAsyncOperation[*IStorageFile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFileAsync, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageFolder) GetFolderAsync(name string) *IAsyncOperation[*IStorageFolder] {
	var _result *IAsyncOperation[*IStorageFolder]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFolderAsync, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageFolder) GetItemAsync(name string) *IAsyncOperation[*IStorageItem] {
	var _result *IAsyncOperation[*IStorageItem]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetItemAsync, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageFolder) GetFilesAsyncOverloadDefaultOptionsStartAndCount() *IAsyncOperation[*IVectorView[*IStorageFile]] {
	var _result *IAsyncOperation[*IVectorView[*IStorageFile]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFilesAsyncOverloadDefaultOptionsStartAndCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageFolder) GetFoldersAsyncOverloadDefaultOptionsStartAndCount() *IAsyncOperation[*IVectorView[*IStorageFolder]] {
	var _result *IAsyncOperation[*IVectorView[*IStorageFolder]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFoldersAsyncOverloadDefaultOptionsStartAndCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageFolder) GetItemsAsyncOverloadDefaultStartAndCount() *IAsyncOperation[*IVectorView[*IStorageItem]] {
	var _result *IAsyncOperation[*IVectorView[*IStorageItem]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetItemsAsyncOverloadDefaultStartAndCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E827E8B9-08D9-4A8E-A0AC-FE5ED3CBBBD3
var IID_IStorageFolder2 = syscall.GUID{0xE827E8B9, 0x08D9, 0x4A8E,
	[8]byte{0xA0, 0xAC, 0xFE, 0x5E, 0xD3, 0xCB, 0xBB, 0xD3}}

type IStorageFolder2Interface interface {
	win32.IInspectableInterface
	TryGetItemAsync(name string) *IAsyncOperation[*IStorageItem]
}

type IStorageFolder2Vtbl struct {
	win32.IInspectableVtbl
	TryGetItemAsync uintptr
}

type IStorageFolder2 struct {
	win32.IInspectable
}

func (this *IStorageFolder2) Vtbl() *IStorageFolder2Vtbl {
	return (*IStorageFolder2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageFolder2) TryGetItemAsync(name string) *IAsyncOperation[*IStorageItem] {
	var _result *IAsyncOperation[*IStorageItem]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetItemAsync, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9F617899-BDE1-4124-AEB3-B06AD96F98D4
var IID_IStorageFolder3 = syscall.GUID{0x9F617899, 0xBDE1, 0x4124,
	[8]byte{0xAE, 0xB3, 0xB0, 0x6A, 0xD9, 0x6F, 0x98, 0xD4}}

type IStorageFolder3Interface interface {
	win32.IInspectableInterface
	TryGetChangeTracker() *IStorageLibraryChangeTracker
}

type IStorageFolder3Vtbl struct {
	win32.IInspectableVtbl
	TryGetChangeTracker uintptr
}

type IStorageFolder3 struct {
	win32.IInspectable
}

func (this *IStorageFolder3) Vtbl() *IStorageFolder3Vtbl {
	return (*IStorageFolder3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageFolder3) TryGetChangeTracker() *IStorageLibraryChangeTracker {
	var _result *IStorageLibraryChangeTracker
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetChangeTracker, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 08F327FF-85D5-48B9-AEE9-28511E339F9F
var IID_IStorageFolderStatics = syscall.GUID{0x08F327FF, 0x85D5, 0x48B9,
	[8]byte{0xAE, 0xE9, 0x28, 0x51, 0x1E, 0x33, 0x9F, 0x9F}}

type IStorageFolderStaticsInterface interface {
	win32.IInspectableInterface
	GetFolderFromPathAsync(path string) *IAsyncOperation[*IStorageFolder]
}

type IStorageFolderStaticsVtbl struct {
	win32.IInspectableVtbl
	GetFolderFromPathAsync uintptr
}

type IStorageFolderStatics struct {
	win32.IInspectable
}

func (this *IStorageFolderStatics) Vtbl() *IStorageFolderStaticsVtbl {
	return (*IStorageFolderStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageFolderStatics) GetFolderFromPathAsync(path string) *IAsyncOperation[*IStorageFolder] {
	var _result *IAsyncOperation[*IStorageFolder]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFolderFromPathAsync, uintptr(unsafe.Pointer(this)), NewHStr(path).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B4656DC3-71D2-467D-8B29-371F0F62BF6F
var IID_IStorageFolderStatics2 = syscall.GUID{0xB4656DC3, 0x71D2, 0x467D,
	[8]byte{0x8B, 0x29, 0x37, 0x1F, 0x0F, 0x62, 0xBF, 0x6F}}

type IStorageFolderStatics2Interface interface {
	win32.IInspectableInterface
	GetFolderFromPathForUserAsync(user *IUser, path string) *IAsyncOperation[*IStorageFolder]
}

type IStorageFolderStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetFolderFromPathForUserAsync uintptr
}

type IStorageFolderStatics2 struct {
	win32.IInspectable
}

func (this *IStorageFolderStatics2) Vtbl() *IStorageFolderStatics2Vtbl {
	return (*IStorageFolderStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageFolderStatics2) GetFolderFromPathForUserAsync(user *IUser, path string) *IAsyncOperation[*IStorageFolder] {
	var _result *IAsyncOperation[*IStorageFolder]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFolderFromPathForUserAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), NewHStr(path).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4207A996-CA2F-42F7-BDE8-8B10457A7F30
var IID_IStorageItem = syscall.GUID{0x4207A996, 0xCA2F, 0x42F7,
	[8]byte{0xBD, 0xE8, 0x8B, 0x10, 0x45, 0x7A, 0x7F, 0x30}}

type IStorageItemInterface interface {
	win32.IInspectableInterface
	RenameAsyncOverloadDefaultOptions(desiredName string) *IAsyncAction
	RenameAsync(desiredName string, option NameCollisionOption) *IAsyncAction
	DeleteAsyncOverloadDefaultOptions() *IAsyncAction
	DeleteAsync(option StorageDeleteOption) *IAsyncAction
	GetBasicPropertiesAsync() *IAsyncOperation[*IBasicProperties]
	Get_Name() string
	Get_Path() string
	Get_Attributes() FileAttributes
	Get_DateCreated() DateTime
	IsOfType(type_ StorageItemTypes) bool
}

type IStorageItemVtbl struct {
	win32.IInspectableVtbl
	RenameAsyncOverloadDefaultOptions uintptr
	RenameAsync                       uintptr
	DeleteAsyncOverloadDefaultOptions uintptr
	DeleteAsync                       uintptr
	GetBasicPropertiesAsync           uintptr
	Get_Name                          uintptr
	Get_Path                          uintptr
	Get_Attributes                    uintptr
	Get_DateCreated                   uintptr
	IsOfType                          uintptr
}

type IStorageItem struct {
	win32.IInspectable
}

func (this *IStorageItem) Vtbl() *IStorageItemVtbl {
	return (*IStorageItemVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageItem) RenameAsyncOverloadDefaultOptions(desiredName string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RenameAsyncOverloadDefaultOptions, uintptr(unsafe.Pointer(this)), NewHStr(desiredName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageItem) RenameAsync(desiredName string, option NameCollisionOption) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RenameAsync, uintptr(unsafe.Pointer(this)), NewHStr(desiredName).Ptr, uintptr(option), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageItem) DeleteAsyncOverloadDefaultOptions() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DeleteAsyncOverloadDefaultOptions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageItem) DeleteAsync(option StorageDeleteOption) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DeleteAsync, uintptr(unsafe.Pointer(this)), uintptr(option), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageItem) GetBasicPropertiesAsync() *IAsyncOperation[*IBasicProperties] {
	var _result *IAsyncOperation[*IBasicProperties]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetBasicPropertiesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageItem) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStorageItem) Get_Path() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Path, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStorageItem) Get_Attributes() FileAttributes {
	var _result FileAttributes
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Attributes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStorageItem) Get_DateCreated() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DateCreated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStorageItem) IsOfType(type_ StorageItemTypes) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsOfType, uintptr(unsafe.Pointer(this)), uintptr(type_), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 53F926D2-083C-4283-B45B-81C007237E44
var IID_IStorageItem2 = syscall.GUID{0x53F926D2, 0x083C, 0x4283,
	[8]byte{0xB4, 0x5B, 0x81, 0xC0, 0x07, 0x23, 0x7E, 0x44}}

type IStorageItem2Interface interface {
	win32.IInspectableInterface
	GetParentAsync() *IAsyncOperation[*IStorageFolder]
	IsEqual(item *IStorageItem) bool
}

type IStorageItem2Vtbl struct {
	win32.IInspectableVtbl
	GetParentAsync uintptr
	IsEqual        uintptr
}

type IStorageItem2 struct {
	win32.IInspectable
}

func (this *IStorageItem2) Vtbl() *IStorageItem2Vtbl {
	return (*IStorageItem2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageItem2) GetParentAsync() *IAsyncOperation[*IStorageFolder] {
	var _result *IAsyncOperation[*IStorageFolder]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetParentAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageItem2) IsEqual(item *IStorageItem) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsEqual, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(item)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 86664478-8029-46FE-A789-1C2F3E2FFB5C
var IID_IStorageItemProperties = syscall.GUID{0x86664478, 0x8029, 0x46FE,
	[8]byte{0xA7, 0x89, 0x1C, 0x2F, 0x3E, 0x2F, 0xFB, 0x5C}}

type IStorageItemPropertiesInterface interface {
	win32.IInspectableInterface
	GetThumbnailAsyncOverloadDefaultSizeDefaultOptions(mode ThumbnailMode) *IAsyncOperation[*IRandomAccessStreamWithContentType]
	GetThumbnailAsyncOverloadDefaultOptions(mode ThumbnailMode, requestedSize uint32) *IAsyncOperation[*IRandomAccessStreamWithContentType]
	GetThumbnailAsync(mode ThumbnailMode, requestedSize uint32, options ThumbnailOptions) *IAsyncOperation[*IRandomAccessStreamWithContentType]
	Get_DisplayName() string
	Get_DisplayType() string
	Get_FolderRelativeId() string
	Get_Properties() *IStorageItemContentProperties
}

type IStorageItemPropertiesVtbl struct {
	win32.IInspectableVtbl
	GetThumbnailAsyncOverloadDefaultSizeDefaultOptions uintptr
	GetThumbnailAsyncOverloadDefaultOptions            uintptr
	GetThumbnailAsync                                  uintptr
	Get_DisplayName                                    uintptr
	Get_DisplayType                                    uintptr
	Get_FolderRelativeId                               uintptr
	Get_Properties                                     uintptr
}

type IStorageItemProperties struct {
	win32.IInspectable
}

func (this *IStorageItemProperties) Vtbl() *IStorageItemPropertiesVtbl {
	return (*IStorageItemPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageItemProperties) GetThumbnailAsyncOverloadDefaultSizeDefaultOptions(mode ThumbnailMode) *IAsyncOperation[*IRandomAccessStreamWithContentType] {
	var _result *IAsyncOperation[*IRandomAccessStreamWithContentType]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetThumbnailAsyncOverloadDefaultSizeDefaultOptions, uintptr(unsafe.Pointer(this)), uintptr(mode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageItemProperties) GetThumbnailAsyncOverloadDefaultOptions(mode ThumbnailMode, requestedSize uint32) *IAsyncOperation[*IRandomAccessStreamWithContentType] {
	var _result *IAsyncOperation[*IRandomAccessStreamWithContentType]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetThumbnailAsyncOverloadDefaultOptions, uintptr(unsafe.Pointer(this)), uintptr(mode), uintptr(requestedSize), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageItemProperties) GetThumbnailAsync(mode ThumbnailMode, requestedSize uint32, options ThumbnailOptions) *IAsyncOperation[*IRandomAccessStreamWithContentType] {
	var _result *IAsyncOperation[*IRandomAccessStreamWithContentType]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetThumbnailAsync, uintptr(unsafe.Pointer(this)), uintptr(mode), uintptr(requestedSize), uintptr(options), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageItemProperties) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStorageItemProperties) Get_DisplayType() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStorageItemProperties) Get_FolderRelativeId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FolderRelativeId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStorageItemProperties) Get_Properties() *IStorageItemContentProperties {
	var _result *IStorageItemContentProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8E86A951-04B9-4BD2-929D-FEF3F71621D0
var IID_IStorageItemProperties2 = syscall.GUID{0x8E86A951, 0x04B9, 0x4BD2,
	[8]byte{0x92, 0x9D, 0xFE, 0xF3, 0xF7, 0x16, 0x21, 0xD0}}

type IStorageItemProperties2Interface interface {
	win32.IInspectableInterface
	GetScaledImageAsThumbnailAsyncOverloadDefaultSizeDefaultOptions(mode ThumbnailMode) *IAsyncOperation[*IRandomAccessStreamWithContentType]
	GetScaledImageAsThumbnailAsyncOverloadDefaultOptions(mode ThumbnailMode, requestedSize uint32) *IAsyncOperation[*IRandomAccessStreamWithContentType]
	GetScaledImageAsThumbnailAsync(mode ThumbnailMode, requestedSize uint32, options ThumbnailOptions) *IAsyncOperation[*IRandomAccessStreamWithContentType]
}

type IStorageItemProperties2Vtbl struct {
	win32.IInspectableVtbl
	GetScaledImageAsThumbnailAsyncOverloadDefaultSizeDefaultOptions uintptr
	GetScaledImageAsThumbnailAsyncOverloadDefaultOptions            uintptr
	GetScaledImageAsThumbnailAsync                                  uintptr
}

type IStorageItemProperties2 struct {
	win32.IInspectable
}

func (this *IStorageItemProperties2) Vtbl() *IStorageItemProperties2Vtbl {
	return (*IStorageItemProperties2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageItemProperties2) GetScaledImageAsThumbnailAsyncOverloadDefaultSizeDefaultOptions(mode ThumbnailMode) *IAsyncOperation[*IRandomAccessStreamWithContentType] {
	var _result *IAsyncOperation[*IRandomAccessStreamWithContentType]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetScaledImageAsThumbnailAsyncOverloadDefaultSizeDefaultOptions, uintptr(unsafe.Pointer(this)), uintptr(mode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageItemProperties2) GetScaledImageAsThumbnailAsyncOverloadDefaultOptions(mode ThumbnailMode, requestedSize uint32) *IAsyncOperation[*IRandomAccessStreamWithContentType] {
	var _result *IAsyncOperation[*IRandomAccessStreamWithContentType]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetScaledImageAsThumbnailAsyncOverloadDefaultOptions, uintptr(unsafe.Pointer(this)), uintptr(mode), uintptr(requestedSize), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageItemProperties2) GetScaledImageAsThumbnailAsync(mode ThumbnailMode, requestedSize uint32, options ThumbnailOptions) *IAsyncOperation[*IRandomAccessStreamWithContentType] {
	var _result *IAsyncOperation[*IRandomAccessStreamWithContentType]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetScaledImageAsThumbnailAsync, uintptr(unsafe.Pointer(this)), uintptr(mode), uintptr(requestedSize), uintptr(options), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 861BF39B-6368-4DEE-B40E-74684A5CE714
var IID_IStorageItemPropertiesWithProvider = syscall.GUID{0x861BF39B, 0x6368, 0x4DEE,
	[8]byte{0xB4, 0x0E, 0x74, 0x68, 0x4A, 0x5C, 0xE7, 0x14}}

type IStorageItemPropertiesWithProviderInterface interface {
	win32.IInspectableInterface
	Get_Provider() *IStorageProvider
}

type IStorageItemPropertiesWithProviderVtbl struct {
	win32.IInspectableVtbl
	Get_Provider uintptr
}

type IStorageItemPropertiesWithProvider struct {
	win32.IInspectable
}

func (this *IStorageItemPropertiesWithProvider) Vtbl() *IStorageItemPropertiesWithProviderVtbl {
	return (*IStorageItemPropertiesWithProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageItemPropertiesWithProvider) Get_Provider() *IStorageProvider {
	var _result *IStorageProvider
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Provider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1EDD7103-0E5E-4D6C-B5E8-9318983D6A03
var IID_IStorageLibrary = syscall.GUID{0x1EDD7103, 0x0E5E, 0x4D6C,
	[8]byte{0xB5, 0xE8, 0x93, 0x18, 0x98, 0x3D, 0x6A, 0x03}}

type IStorageLibraryInterface interface {
	win32.IInspectableInterface
	RequestAddFolderAsync() *IAsyncOperation[*IStorageFolder]
	RequestRemoveFolderAsync(folder *IStorageFolder) *IAsyncOperation[bool]
	Get_Folders() *IObservableVector[*IStorageFolder]
	Get_SaveFolder() *IStorageFolder
	Add_DefinitionChanged(handler TypedEventHandler[*IStorageLibrary, interface{}]) EventRegistrationToken
	Remove_DefinitionChanged(eventCookie EventRegistrationToken)
}

type IStorageLibraryVtbl struct {
	win32.IInspectableVtbl
	RequestAddFolderAsync    uintptr
	RequestRemoveFolderAsync uintptr
	Get_Folders              uintptr
	Get_SaveFolder           uintptr
	Add_DefinitionChanged    uintptr
	Remove_DefinitionChanged uintptr
}

type IStorageLibrary struct {
	win32.IInspectable
}

func (this *IStorageLibrary) Vtbl() *IStorageLibraryVtbl {
	return (*IStorageLibraryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageLibrary) RequestAddFolderAsync() *IAsyncOperation[*IStorageFolder] {
	var _result *IAsyncOperation[*IStorageFolder]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAddFolderAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageLibrary) RequestRemoveFolderAsync(folder *IStorageFolder) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestRemoveFolderAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(folder)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageLibrary) Get_Folders() *IObservableVector[*IStorageFolder] {
	var _result *IObservableVector[*IStorageFolder]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Folders, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageLibrary) Get_SaveFolder() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SaveFolder, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageLibrary) Add_DefinitionChanged(handler TypedEventHandler[*IStorageLibrary, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DefinitionChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStorageLibrary) Remove_DefinitionChanged(eventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DefinitionChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&eventCookie)))
	_ = _hr
}

// 5B0CE348-FCB3-4031-AFB0-A68D7BD44534
var IID_IStorageLibrary2 = syscall.GUID{0x5B0CE348, 0xFCB3, 0x4031,
	[8]byte{0xAF, 0xB0, 0xA6, 0x8D, 0x7B, 0xD4, 0x45, 0x34}}

type IStorageLibrary2Interface interface {
	win32.IInspectableInterface
	Get_ChangeTracker() *IStorageLibraryChangeTracker
}

type IStorageLibrary2Vtbl struct {
	win32.IInspectableVtbl
	Get_ChangeTracker uintptr
}

type IStorageLibrary2 struct {
	win32.IInspectable
}

func (this *IStorageLibrary2) Vtbl() *IStorageLibrary2Vtbl {
	return (*IStorageLibrary2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageLibrary2) Get_ChangeTracker() *IStorageLibraryChangeTracker {
	var _result *IStorageLibraryChangeTracker
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ChangeTracker, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8A281291-2154-4201-8113-D2C05CE1AD23
var IID_IStorageLibrary3 = syscall.GUID{0x8A281291, 0x2154, 0x4201,
	[8]byte{0x81, 0x13, 0xD2, 0xC0, 0x5C, 0xE1, 0xAD, 0x23}}

type IStorageLibrary3Interface interface {
	win32.IInspectableInterface
	AreFolderSuggestionsAvailableAsync() *IAsyncOperation[bool]
}

type IStorageLibrary3Vtbl struct {
	win32.IInspectableVtbl
	AreFolderSuggestionsAvailableAsync uintptr
}

type IStorageLibrary3 struct {
	win32.IInspectable
}

func (this *IStorageLibrary3) Vtbl() *IStorageLibrary3Vtbl {
	return (*IStorageLibrary3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageLibrary3) AreFolderSuggestionsAvailableAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AreFolderSuggestionsAvailableAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 00980B23-2BE2-4909-AA48-159F5203A51E
var IID_IStorageLibraryChange = syscall.GUID{0x00980B23, 0x2BE2, 0x4909,
	[8]byte{0xAA, 0x48, 0x15, 0x9F, 0x52, 0x03, 0xA5, 0x1E}}

type IStorageLibraryChangeInterface interface {
	win32.IInspectableInterface
	Get_ChangeType() StorageLibraryChangeType
	Get_Path() string
	Get_PreviousPath() string
	IsOfType(type_ StorageItemTypes) bool
	GetStorageItemAsync() *IAsyncOperation[*IStorageItem]
}

type IStorageLibraryChangeVtbl struct {
	win32.IInspectableVtbl
	Get_ChangeType      uintptr
	Get_Path            uintptr
	Get_PreviousPath    uintptr
	IsOfType            uintptr
	GetStorageItemAsync uintptr
}

type IStorageLibraryChange struct {
	win32.IInspectable
}

func (this *IStorageLibraryChange) Vtbl() *IStorageLibraryChangeVtbl {
	return (*IStorageLibraryChangeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageLibraryChange) Get_ChangeType() StorageLibraryChangeType {
	var _result StorageLibraryChangeType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ChangeType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStorageLibraryChange) Get_Path() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Path, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStorageLibraryChange) Get_PreviousPath() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PreviousPath, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStorageLibraryChange) IsOfType(type_ StorageItemTypes) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsOfType, uintptr(unsafe.Pointer(this)), uintptr(type_), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStorageLibraryChange) GetStorageItemAsync() *IAsyncOperation[*IStorageItem] {
	var _result *IAsyncOperation[*IStorageItem]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStorageItemAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F205BC83-FCA2-41F9-8954-EE2E991EB96F
var IID_IStorageLibraryChangeReader = syscall.GUID{0xF205BC83, 0xFCA2, 0x41F9,
	[8]byte{0x89, 0x54, 0xEE, 0x2E, 0x99, 0x1E, 0xB9, 0x6F}}

type IStorageLibraryChangeReaderInterface interface {
	win32.IInspectableInterface
	ReadBatchAsync() *IAsyncOperation[*IVectorView[*IStorageLibraryChange]]
	AcceptChangesAsync() *IAsyncAction
}

type IStorageLibraryChangeReaderVtbl struct {
	win32.IInspectableVtbl
	ReadBatchAsync     uintptr
	AcceptChangesAsync uintptr
}

type IStorageLibraryChangeReader struct {
	win32.IInspectable
}

func (this *IStorageLibraryChangeReader) Vtbl() *IStorageLibraryChangeReaderVtbl {
	return (*IStorageLibraryChangeReaderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageLibraryChangeReader) ReadBatchAsync() *IAsyncOperation[*IVectorView[*IStorageLibraryChange]] {
	var _result *IAsyncOperation[*IVectorView[*IStorageLibraryChange]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadBatchAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageLibraryChangeReader) AcceptChangesAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AcceptChangesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// ABF4868B-FBCC-4A4F-999E-E7AB7C646DBE
var IID_IStorageLibraryChangeReader2 = syscall.GUID{0xABF4868B, 0xFBCC, 0x4A4F,
	[8]byte{0x99, 0x9E, 0xE7, 0xAB, 0x7C, 0x64, 0x6D, 0xBE}}

type IStorageLibraryChangeReader2Interface interface {
	win32.IInspectableInterface
	GetLastChangeId() uint64
}

type IStorageLibraryChangeReader2Vtbl struct {
	win32.IInspectableVtbl
	GetLastChangeId uintptr
}

type IStorageLibraryChangeReader2 struct {
	win32.IInspectable
}

func (this *IStorageLibraryChangeReader2) Vtbl() *IStorageLibraryChangeReader2Vtbl {
	return (*IStorageLibraryChangeReader2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageLibraryChangeReader2) GetLastChangeId() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetLastChangeId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 9E157316-6073-44F6-9681-7492D1286C90
var IID_IStorageLibraryChangeTracker = syscall.GUID{0x9E157316, 0x6073, 0x44F6,
	[8]byte{0x96, 0x81, 0x74, 0x92, 0xD1, 0x28, 0x6C, 0x90}}

type IStorageLibraryChangeTrackerInterface interface {
	win32.IInspectableInterface
	GetChangeReader() *IStorageLibraryChangeReader
	Enable()
	Reset()
}

type IStorageLibraryChangeTrackerVtbl struct {
	win32.IInspectableVtbl
	GetChangeReader uintptr
	Enable          uintptr
	Reset           uintptr
}

type IStorageLibraryChangeTracker struct {
	win32.IInspectable
}

func (this *IStorageLibraryChangeTracker) Vtbl() *IStorageLibraryChangeTrackerVtbl {
	return (*IStorageLibraryChangeTrackerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageLibraryChangeTracker) GetChangeReader() *IStorageLibraryChangeReader {
	var _result *IStorageLibraryChangeReader
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetChangeReader, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageLibraryChangeTracker) Enable() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Enable, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IStorageLibraryChangeTracker) Reset() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Reset, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// CD051C3B-0F9F-42F9-8FB3-158D82E13821
var IID_IStorageLibraryChangeTracker2 = syscall.GUID{0xCD051C3B, 0x0F9F, 0x42F9,
	[8]byte{0x8F, 0xB3, 0x15, 0x8D, 0x82, 0xE1, 0x38, 0x21}}

type IStorageLibraryChangeTracker2Interface interface {
	win32.IInspectableInterface
	EnableWithOptions(options *IStorageLibraryChangeTrackerOptions)
	Disable()
}

type IStorageLibraryChangeTracker2Vtbl struct {
	win32.IInspectableVtbl
	EnableWithOptions uintptr
	Disable           uintptr
}

type IStorageLibraryChangeTracker2 struct {
	win32.IInspectable
}

func (this *IStorageLibraryChangeTracker2) Vtbl() *IStorageLibraryChangeTracker2Vtbl {
	return (*IStorageLibraryChangeTracker2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageLibraryChangeTracker2) EnableWithOptions(options *IStorageLibraryChangeTrackerOptions) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnableWithOptions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(options)))
	_ = _hr
}

func (this *IStorageLibraryChangeTracker2) Disable() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Disable, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// BB52BCD4-1A6D-59C0-AD2A-823A20532483
var IID_IStorageLibraryChangeTrackerOptions = syscall.GUID{0xBB52BCD4, 0x1A6D, 0x59C0,
	[8]byte{0xAD, 0x2A, 0x82, 0x3A, 0x20, 0x53, 0x24, 0x83}}

type IStorageLibraryChangeTrackerOptionsInterface interface {
	win32.IInspectableInterface
	Get_TrackChangeDetails() bool
	Put_TrackChangeDetails(value bool)
}

type IStorageLibraryChangeTrackerOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_TrackChangeDetails uintptr
	Put_TrackChangeDetails uintptr
}

type IStorageLibraryChangeTrackerOptions struct {
	win32.IInspectable
}

func (this *IStorageLibraryChangeTrackerOptions) Vtbl() *IStorageLibraryChangeTrackerOptionsVtbl {
	return (*IStorageLibraryChangeTrackerOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageLibraryChangeTrackerOptions) Get_TrackChangeDetails() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TrackChangeDetails, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStorageLibraryChangeTrackerOptions) Put_TrackChangeDetails(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TrackChangeDetails, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 5281826A-BBE1-53BC-82CA-81CC7F039329
var IID_IStorageLibraryLastChangeId = syscall.GUID{0x5281826A, 0xBBE1, 0x53BC,
	[8]byte{0x82, 0xCA, 0x81, 0xCC, 0x7F, 0x03, 0x93, 0x29}}

type IStorageLibraryLastChangeIdInterface interface {
	win32.IInspectableInterface
}

type IStorageLibraryLastChangeIdVtbl struct {
	win32.IInspectableVtbl
}

type IStorageLibraryLastChangeId struct {
	win32.IInspectable
}

func (this *IStorageLibraryLastChangeId) Vtbl() *IStorageLibraryLastChangeIdVtbl {
	return (*IStorageLibraryLastChangeIdVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 81A49128-2CA3-5309-B0D1-CF0788E40762
var IID_IStorageLibraryLastChangeIdStatics = syscall.GUID{0x81A49128, 0x2CA3, 0x5309,
	[8]byte{0xB0, 0xD1, 0xCF, 0x07, 0x88, 0xE4, 0x07, 0x62}}

type IStorageLibraryLastChangeIdStaticsInterface interface {
	win32.IInspectableInterface
	Get_Unknown() uint64
}

type IStorageLibraryLastChangeIdStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Unknown uintptr
}

type IStorageLibraryLastChangeIdStatics struct {
	win32.IInspectable
}

func (this *IStorageLibraryLastChangeIdStatics) Vtbl() *IStorageLibraryLastChangeIdStaticsVtbl {
	return (*IStorageLibraryLastChangeIdStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageLibraryLastChangeIdStatics) Get_Unknown() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Unknown, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4208A6DB-684A-49C6-9E59-90121EE050D6
var IID_IStorageLibraryStatics = syscall.GUID{0x4208A6DB, 0x684A, 0x49C6,
	[8]byte{0x9E, 0x59, 0x90, 0x12, 0x1E, 0xE0, 0x50, 0xD6}}

type IStorageLibraryStaticsInterface interface {
	win32.IInspectableInterface
	GetLibraryAsync(libraryId KnownLibraryId) *IAsyncOperation[*IStorageLibrary]
}

type IStorageLibraryStaticsVtbl struct {
	win32.IInspectableVtbl
	GetLibraryAsync uintptr
}

type IStorageLibraryStatics struct {
	win32.IInspectable
}

func (this *IStorageLibraryStatics) Vtbl() *IStorageLibraryStaticsVtbl {
	return (*IStorageLibraryStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageLibraryStatics) GetLibraryAsync(libraryId KnownLibraryId) *IAsyncOperation[*IStorageLibrary] {
	var _result *IAsyncOperation[*IStorageLibrary]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetLibraryAsync, uintptr(unsafe.Pointer(this)), uintptr(libraryId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FFB08DDC-FA75-4695-B9D1-7F81F97832E3
var IID_IStorageLibraryStatics2 = syscall.GUID{0xFFB08DDC, 0xFA75, 0x4695,
	[8]byte{0xB9, 0xD1, 0x7F, 0x81, 0xF9, 0x78, 0x32, 0xE3}}

type IStorageLibraryStatics2Interface interface {
	win32.IInspectableInterface
	GetLibraryForUserAsync(user *IUser, libraryId KnownLibraryId) *IAsyncOperation[*IStorageLibrary]
}

type IStorageLibraryStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetLibraryForUserAsync uintptr
}

type IStorageLibraryStatics2 struct {
	win32.IInspectable
}

func (this *IStorageLibraryStatics2) Vtbl() *IStorageLibraryStatics2Vtbl {
	return (*IStorageLibraryStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageLibraryStatics2) GetLibraryForUserAsync(user *IUser, libraryId KnownLibraryId) *IAsyncOperation[*IStorageLibrary] {
	var _result *IAsyncOperation[*IStorageLibrary]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetLibraryForUserAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(libraryId), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E705EED4-D478-47D6-BA46-1A8EBE114A20
var IID_IStorageProvider = syscall.GUID{0xE705EED4, 0xD478, 0x47D6,
	[8]byte{0xBA, 0x46, 0x1A, 0x8E, 0xBE, 0x11, 0x4A, 0x20}}

type IStorageProviderInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	Get_DisplayName() string
}

type IStorageProviderVtbl struct {
	win32.IInspectableVtbl
	Get_Id          uintptr
	Get_DisplayName uintptr
}

type IStorageProvider struct {
	win32.IInspectable
}

func (this *IStorageProvider) Vtbl() *IStorageProviderVtbl {
	return (*IStorageProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageProvider) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStorageProvider) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 010D1917-3404-414B-9FD7-CD44472EAA39
var IID_IStorageProvider2 = syscall.GUID{0x010D1917, 0x3404, 0x414B,
	[8]byte{0x9F, 0xD7, 0xCD, 0x44, 0x47, 0x2E, 0xAA, 0x39}}

type IStorageProvider2Interface interface {
	win32.IInspectableInterface
	IsPropertySupportedForPartialFileAsync(propertyCanonicalName string) *IAsyncOperation[bool]
}

type IStorageProvider2Vtbl struct {
	win32.IInspectableVtbl
	IsPropertySupportedForPartialFileAsync uintptr
}

type IStorageProvider2 struct {
	win32.IInspectable
}

func (this *IStorageProvider2) Vtbl() *IStorageProvider2Vtbl {
	return (*IStorageProvider2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageProvider2) IsPropertySupportedForPartialFileAsync(propertyCanonicalName string) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsPropertySupportedForPartialFileAsync, uintptr(unsafe.Pointer(this)), NewHStr(propertyCanonicalName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F67CF363-A53D-4D94-AE2C-67232D93ACDD
var IID_IStorageStreamTransaction = syscall.GUID{0xF67CF363, 0xA53D, 0x4D94,
	[8]byte{0xAE, 0x2C, 0x67, 0x23, 0x2D, 0x93, 0xAC, 0xDD}}

type IStorageStreamTransactionInterface interface {
	win32.IInspectableInterface
	Get_Stream() *IRandomAccessStream
	CommitAsync() *IAsyncAction
}

type IStorageStreamTransactionVtbl struct {
	win32.IInspectableVtbl
	Get_Stream  uintptr
	CommitAsync uintptr
}

type IStorageStreamTransaction struct {
	win32.IInspectable
}

func (this *IStorageStreamTransaction) Vtbl() *IStorageStreamTransactionVtbl {
	return (*IStorageStreamTransactionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageStreamTransaction) Get_Stream() *IRandomAccessStream {
	var _result *IRandomAccessStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Stream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageStreamTransaction) CommitAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CommitAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1673FCCE-DABD-4D50-BEEE-180B8A8191B6
var IID_IStreamedFileDataRequest = syscall.GUID{0x1673FCCE, 0xDABD, 0x4D50,
	[8]byte{0xBE, 0xEE, 0x18, 0x0B, 0x8A, 0x81, 0x91, 0xB6}}

type IStreamedFileDataRequestInterface interface {
	win32.IInspectableInterface
	FailAndClose(failureMode StreamedFileFailureMode)
}

type IStreamedFileDataRequestVtbl struct {
	win32.IInspectableVtbl
	FailAndClose uintptr
}

type IStreamedFileDataRequest struct {
	win32.IInspectable
}

func (this *IStreamedFileDataRequest) Vtbl() *IStreamedFileDataRequestVtbl {
	return (*IStreamedFileDataRequestVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStreamedFileDataRequest) FailAndClose(failureMode StreamedFileFailureMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FailAndClose, uintptr(unsafe.Pointer(this)), uintptr(failureMode))
	_ = _hr
}

// 3F8F38B7-308C-47E1-924D-8645348E5DB7
var IID_ISystemAudioProperties = syscall.GUID{0x3F8F38B7, 0x308C, 0x47E1,
	[8]byte{0x92, 0x4D, 0x86, 0x45, 0x34, 0x8E, 0x5D, 0xB7}}

type ISystemAudioPropertiesInterface interface {
	win32.IInspectableInterface
	Get_EncodingBitrate() string
}

type ISystemAudioPropertiesVtbl struct {
	win32.IInspectableVtbl
	Get_EncodingBitrate uintptr
}

type ISystemAudioProperties struct {
	win32.IInspectable
}

func (this *ISystemAudioProperties) Vtbl() *ISystemAudioPropertiesVtbl {
	return (*ISystemAudioPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemAudioProperties) Get_EncodingBitrate() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EncodingBitrate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// E32ABF70-D8FA-45EC-A942-D2E26FB60BA5
var IID_ISystemDataPaths = syscall.GUID{0xE32ABF70, 0xD8FA, 0x45EC,
	[8]byte{0xA9, 0x42, 0xD2, 0xE2, 0x6F, 0xB6, 0x0B, 0xA5}}

type ISystemDataPathsInterface interface {
	win32.IInspectableInterface
	Get_Fonts() string
	Get_ProgramData() string
	Get_Public() string
	Get_PublicDesktop() string
	Get_PublicDocuments() string
	Get_PublicDownloads() string
	Get_PublicMusic() string
	Get_PublicPictures() string
	Get_PublicVideos() string
	Get_System() string
	Get_SystemHost() string
	Get_SystemX86() string
	Get_SystemX64() string
	Get_SystemArm() string
	Get_UserProfiles() string
	Get_Windows() string
}

type ISystemDataPathsVtbl struct {
	win32.IInspectableVtbl
	Get_Fonts           uintptr
	Get_ProgramData     uintptr
	Get_Public          uintptr
	Get_PublicDesktop   uintptr
	Get_PublicDocuments uintptr
	Get_PublicDownloads uintptr
	Get_PublicMusic     uintptr
	Get_PublicPictures  uintptr
	Get_PublicVideos    uintptr
	Get_System          uintptr
	Get_SystemHost      uintptr
	Get_SystemX86       uintptr
	Get_SystemX64       uintptr
	Get_SystemArm       uintptr
	Get_UserProfiles    uintptr
	Get_Windows         uintptr
}

type ISystemDataPaths struct {
	win32.IInspectable
}

func (this *ISystemDataPaths) Vtbl() *ISystemDataPathsVtbl {
	return (*ISystemDataPathsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemDataPaths) Get_Fonts() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Fonts, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemDataPaths) Get_ProgramData() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProgramData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemDataPaths) Get_Public() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Public, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemDataPaths) Get_PublicDesktop() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PublicDesktop, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemDataPaths) Get_PublicDocuments() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PublicDocuments, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemDataPaths) Get_PublicDownloads() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PublicDownloads, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemDataPaths) Get_PublicMusic() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PublicMusic, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemDataPaths) Get_PublicPictures() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PublicPictures, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemDataPaths) Get_PublicVideos() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PublicVideos, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemDataPaths) Get_System() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_System, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemDataPaths) Get_SystemHost() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SystemHost, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemDataPaths) Get_SystemX86() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SystemX86, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemDataPaths) Get_SystemX64() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SystemX64, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemDataPaths) Get_SystemArm() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SystemArm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemDataPaths) Get_UserProfiles() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UserProfiles, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemDataPaths) Get_Windows() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Windows, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// E0F96FD0-9920-4BCA-B379-F96FDF7CAAD8
var IID_ISystemDataPathsStatics = syscall.GUID{0xE0F96FD0, 0x9920, 0x4BCA,
	[8]byte{0xB3, 0x79, 0xF9, 0x6F, 0xDF, 0x7C, 0xAA, 0xD8}}

type ISystemDataPathsStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *ISystemDataPaths
}

type ISystemDataPathsStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
}

type ISystemDataPathsStatics struct {
	win32.IInspectable
}

func (this *ISystemDataPathsStatics) Vtbl() *ISystemDataPathsStaticsVtbl {
	return (*ISystemDataPathsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemDataPathsStatics) GetDefault() *ISystemDataPaths {
	var _result *ISystemDataPaths
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C0F46EB4-C174-481A-BC25-921986F6A6F3
var IID_ISystemGPSProperties = syscall.GUID{0xC0F46EB4, 0xC174, 0x481A,
	[8]byte{0xBC, 0x25, 0x92, 0x19, 0x86, 0xF6, 0xA6, 0xF3}}

type ISystemGPSPropertiesInterface interface {
	win32.IInspectableInterface
	Get_LatitudeDecimal() string
	Get_LongitudeDecimal() string
}

type ISystemGPSPropertiesVtbl struct {
	win32.IInspectableVtbl
	Get_LatitudeDecimal  uintptr
	Get_LongitudeDecimal uintptr
}

type ISystemGPSProperties struct {
	win32.IInspectable
}

func (this *ISystemGPSProperties) Vtbl() *ISystemGPSPropertiesVtbl {
	return (*ISystemGPSPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemGPSProperties) Get_LatitudeDecimal() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LatitudeDecimal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemGPSProperties) Get_LongitudeDecimal() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LongitudeDecimal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 011B2E30-8B39-4308-BEA1-E8AA61E47826
var IID_ISystemImageProperties = syscall.GUID{0x011B2E30, 0x8B39, 0x4308,
	[8]byte{0xBE, 0xA1, 0xE8, 0xAA, 0x61, 0xE4, 0x78, 0x26}}

type ISystemImagePropertiesInterface interface {
	win32.IInspectableInterface
	Get_HorizontalSize() string
	Get_VerticalSize() string
}

type ISystemImagePropertiesVtbl struct {
	win32.IInspectableVtbl
	Get_HorizontalSize uintptr
	Get_VerticalSize   uintptr
}

type ISystemImageProperties struct {
	win32.IInspectable
}

func (this *ISystemImageProperties) Vtbl() *ISystemImagePropertiesVtbl {
	return (*ISystemImagePropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemImageProperties) Get_HorizontalSize() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HorizontalSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemImageProperties) Get_VerticalSize() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VerticalSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// A42B3316-8415-40DC-8C44-98361D235430
var IID_ISystemMediaProperties = syscall.GUID{0xA42B3316, 0x8415, 0x40DC,
	[8]byte{0x8C, 0x44, 0x98, 0x36, 0x1D, 0x23, 0x54, 0x30}}

type ISystemMediaPropertiesInterface interface {
	win32.IInspectableInterface
	Get_Duration() string
	Get_Producer() string
	Get_Publisher() string
	Get_SubTitle() string
	Get_Writer() string
	Get_Year() string
}

type ISystemMediaPropertiesVtbl struct {
	win32.IInspectableVtbl
	Get_Duration  uintptr
	Get_Producer  uintptr
	Get_Publisher uintptr
	Get_SubTitle  uintptr
	Get_Writer    uintptr
	Get_Year      uintptr
}

type ISystemMediaProperties struct {
	win32.IInspectable
}

func (this *ISystemMediaProperties) Vtbl() *ISystemMediaPropertiesVtbl {
	return (*ISystemMediaPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemMediaProperties) Get_Duration() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Duration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemMediaProperties) Get_Producer() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Producer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemMediaProperties) Get_Publisher() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Publisher, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemMediaProperties) Get_SubTitle() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SubTitle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemMediaProperties) Get_Writer() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Writer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemMediaProperties) Get_Year() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Year, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// B47988D5-67AF-4BC3-8D39-5B89022026A1
var IID_ISystemMusicProperties = syscall.GUID{0xB47988D5, 0x67AF, 0x4BC3,
	[8]byte{0x8D, 0x39, 0x5B, 0x89, 0x02, 0x20, 0x26, 0xA1}}

type ISystemMusicPropertiesInterface interface {
	win32.IInspectableInterface
	Get_AlbumArtist() string
	Get_AlbumTitle() string
	Get_Artist() string
	Get_Composer() string
	Get_Conductor() string
	Get_DisplayArtist() string
	Get_Genre() string
	Get_TrackNumber() string
}

type ISystemMusicPropertiesVtbl struct {
	win32.IInspectableVtbl
	Get_AlbumArtist   uintptr
	Get_AlbumTitle    uintptr
	Get_Artist        uintptr
	Get_Composer      uintptr
	Get_Conductor     uintptr
	Get_DisplayArtist uintptr
	Get_Genre         uintptr
	Get_TrackNumber   uintptr
}

type ISystemMusicProperties struct {
	win32.IInspectable
}

func (this *ISystemMusicProperties) Vtbl() *ISystemMusicPropertiesVtbl {
	return (*ISystemMusicPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemMusicProperties) Get_AlbumArtist() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlbumArtist, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemMusicProperties) Get_AlbumTitle() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlbumTitle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemMusicProperties) Get_Artist() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Artist, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemMusicProperties) Get_Composer() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Composer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemMusicProperties) Get_Conductor() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Conductor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemMusicProperties) Get_DisplayArtist() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayArtist, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemMusicProperties) Get_Genre() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Genre, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemMusicProperties) Get_TrackNumber() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TrackNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 4734FC3D-AB21-4424-B735-F4353A56C8FC
var IID_ISystemPhotoProperties = syscall.GUID{0x4734FC3D, 0xAB21, 0x4424,
	[8]byte{0xB7, 0x35, 0xF4, 0x35, 0x3A, 0x56, 0xC8, 0xFC}}

type ISystemPhotoPropertiesInterface interface {
	win32.IInspectableInterface
	Get_CameraManufacturer() string
	Get_CameraModel() string
	Get_DateTaken() string
	Get_Orientation() string
	Get_PeopleNames() string
}

type ISystemPhotoPropertiesVtbl struct {
	win32.IInspectableVtbl
	Get_CameraManufacturer uintptr
	Get_CameraModel        uintptr
	Get_DateTaken          uintptr
	Get_Orientation        uintptr
	Get_PeopleNames        uintptr
}

type ISystemPhotoProperties struct {
	win32.IInspectable
}

func (this *ISystemPhotoProperties) Vtbl() *ISystemPhotoPropertiesVtbl {
	return (*ISystemPhotoPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemPhotoProperties) Get_CameraManufacturer() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CameraManufacturer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemPhotoProperties) Get_CameraModel() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CameraModel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemPhotoProperties) Get_DateTaken() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DateTaken, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemPhotoProperties) Get_Orientation() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Orientation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemPhotoProperties) Get_PeopleNames() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PeopleNames, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 917A71C1-85F3-4DD1-B001-A50BFD21C8D2
var IID_ISystemProperties = syscall.GUID{0x917A71C1, 0x85F3, 0x4DD1,
	[8]byte{0xB0, 0x01, 0xA5, 0x0B, 0xFD, 0x21, 0xC8, 0xD2}}

type ISystemPropertiesInterface interface {
	win32.IInspectableInterface
	Get_Author() string
	Get_Comment() string
	Get_ItemNameDisplay() string
	Get_Keywords() string
	Get_Rating() string
	Get_Title() string
	Get_Audio() *ISystemAudioProperties
	Get_GPS() *ISystemGPSProperties
	Get_Media() *ISystemMediaProperties
	Get_Music() *ISystemMusicProperties
	Get_Photo() *ISystemPhotoProperties
	Get_Video() *ISystemVideoProperties
	Get_Image() *ISystemImageProperties
}

type ISystemPropertiesVtbl struct {
	win32.IInspectableVtbl
	Get_Author          uintptr
	Get_Comment         uintptr
	Get_ItemNameDisplay uintptr
	Get_Keywords        uintptr
	Get_Rating          uintptr
	Get_Title           uintptr
	Get_Audio           uintptr
	Get_GPS             uintptr
	Get_Media           uintptr
	Get_Music           uintptr
	Get_Photo           uintptr
	Get_Video           uintptr
	Get_Image           uintptr
}

type ISystemProperties struct {
	win32.IInspectable
}

func (this *ISystemProperties) Vtbl() *ISystemPropertiesVtbl {
	return (*ISystemPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemProperties) Get_Author() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Author, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemProperties) Get_Comment() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Comment, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemProperties) Get_ItemNameDisplay() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ItemNameDisplay, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemProperties) Get_Keywords() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Keywords, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemProperties) Get_Rating() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Rating, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemProperties) Get_Title() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Title, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemProperties) Get_Audio() *ISystemAudioProperties {
	var _result *ISystemAudioProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Audio, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISystemProperties) Get_GPS() *ISystemGPSProperties {
	var _result *ISystemGPSProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GPS, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISystemProperties) Get_Media() *ISystemMediaProperties {
	var _result *ISystemMediaProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Media, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISystemProperties) Get_Music() *ISystemMusicProperties {
	var _result *ISystemMusicProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Music, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISystemProperties) Get_Photo() *ISystemPhotoProperties {
	var _result *ISystemPhotoProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Photo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISystemProperties) Get_Video() *ISystemVideoProperties {
	var _result *ISystemVideoProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Video, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISystemProperties) Get_Image() *ISystemImageProperties {
	var _result *ISystemImageProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Image, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2040F715-67F8-4322-9B80-4FA9FEFB83E8
var IID_ISystemVideoProperties = syscall.GUID{0x2040F715, 0x67F8, 0x4322,
	[8]byte{0x9B, 0x80, 0x4F, 0xA9, 0xFE, 0xFB, 0x83, 0xE8}}

type ISystemVideoPropertiesInterface interface {
	win32.IInspectableInterface
	Get_Director() string
	Get_FrameHeight() string
	Get_FrameWidth() string
	Get_Orientation() string
	Get_TotalBitrate() string
}

type ISystemVideoPropertiesVtbl struct {
	win32.IInspectableVtbl
	Get_Director     uintptr
	Get_FrameHeight  uintptr
	Get_FrameWidth   uintptr
	Get_Orientation  uintptr
	Get_TotalBitrate uintptr
}

type ISystemVideoProperties struct {
	win32.IInspectable
}

func (this *ISystemVideoProperties) Vtbl() *ISystemVideoPropertiesVtbl {
	return (*ISystemVideoPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemVideoProperties) Get_Director() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Director, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemVideoProperties) Get_FrameHeight() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameHeight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemVideoProperties) Get_FrameWidth() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameWidth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemVideoProperties) Get_Orientation() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Orientation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISystemVideoProperties) Get_TotalBitrate() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TotalBitrate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// F9C53912-ABC4-46FF-8A2B-DC9D7FA6E52F
var IID_IUserDataPaths = syscall.GUID{0xF9C53912, 0xABC4, 0x46FF,
	[8]byte{0x8A, 0x2B, 0xDC, 0x9D, 0x7F, 0xA6, 0xE5, 0x2F}}

type IUserDataPathsInterface interface {
	win32.IInspectableInterface
	Get_CameraRoll() string
	Get_Cookies() string
	Get_Desktop() string
	Get_Documents() string
	Get_Downloads() string
	Get_Favorites() string
	Get_History() string
	Get_InternetCache() string
	Get_LocalAppData() string
	Get_LocalAppDataLow() string
	Get_Music() string
	Get_Pictures() string
	Get_Profile() string
	Get_Recent() string
	Get_RoamingAppData() string
	Get_SavedPictures() string
	Get_Screenshots() string
	Get_Templates() string
	Get_Videos() string
}

type IUserDataPathsVtbl struct {
	win32.IInspectableVtbl
	Get_CameraRoll      uintptr
	Get_Cookies         uintptr
	Get_Desktop         uintptr
	Get_Documents       uintptr
	Get_Downloads       uintptr
	Get_Favorites       uintptr
	Get_History         uintptr
	Get_InternetCache   uintptr
	Get_LocalAppData    uintptr
	Get_LocalAppDataLow uintptr
	Get_Music           uintptr
	Get_Pictures        uintptr
	Get_Profile         uintptr
	Get_Recent          uintptr
	Get_RoamingAppData  uintptr
	Get_SavedPictures   uintptr
	Get_Screenshots     uintptr
	Get_Templates       uintptr
	Get_Videos          uintptr
}

type IUserDataPaths struct {
	win32.IInspectable
}

func (this *IUserDataPaths) Vtbl() *IUserDataPathsVtbl {
	return (*IUserDataPathsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUserDataPaths) Get_CameraRoll() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CameraRoll, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUserDataPaths) Get_Cookies() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Cookies, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUserDataPaths) Get_Desktop() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Desktop, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUserDataPaths) Get_Documents() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Documents, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUserDataPaths) Get_Downloads() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Downloads, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUserDataPaths) Get_Favorites() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Favorites, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUserDataPaths) Get_History() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_History, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUserDataPaths) Get_InternetCache() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InternetCache, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUserDataPaths) Get_LocalAppData() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocalAppData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUserDataPaths) Get_LocalAppDataLow() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocalAppDataLow, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUserDataPaths) Get_Music() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Music, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUserDataPaths) Get_Pictures() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Pictures, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUserDataPaths) Get_Profile() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Profile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUserDataPaths) Get_Recent() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Recent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUserDataPaths) Get_RoamingAppData() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RoamingAppData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUserDataPaths) Get_SavedPictures() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SavedPictures, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUserDataPaths) Get_Screenshots() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Screenshots, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUserDataPaths) Get_Templates() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Templates, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUserDataPaths) Get_Videos() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Videos, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 01B29DEF-E062-48A1-8B0C-F2C7A9CA56C0
var IID_IUserDataPathsStatics = syscall.GUID{0x01B29DEF, 0xE062, 0x48A1,
	[8]byte{0x8B, 0x0C, 0xF2, 0xC7, 0xA9, 0xCA, 0x56, 0xC0}}

type IUserDataPathsStaticsInterface interface {
	win32.IInspectableInterface
	GetForUser(user *IUser) *IUserDataPaths
	GetDefault() *IUserDataPaths
}

type IUserDataPathsStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForUser uintptr
	GetDefault uintptr
}

type IUserDataPathsStatics struct {
	win32.IInspectable
}

func (this *IUserDataPathsStatics) Vtbl() *IUserDataPathsStaticsVtbl {
	return (*IUserDataPathsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUserDataPathsStatics) GetForUser(user *IUser) *IUserDataPaths {
	var _result *IUserDataPaths
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserDataPathsStatics) GetDefault() *IUserDataPaths {
	var _result *IUserDataPaths
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type ApplicationData struct {
	RtClass
	*IApplicationData
}

func NewIApplicationDataStatics() *IApplicationDataStatics {
	var p *IApplicationDataStatics
	hs := NewHStr("Windows.Storage.ApplicationData")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IApplicationDataStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIApplicationDataStatics2() *IApplicationDataStatics2 {
	var p *IApplicationDataStatics2
	hs := NewHStr("Windows.Storage.ApplicationData")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IApplicationDataStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type ApplicationDataCompositeValue struct {
	RtClass
	*IPropertySet
}

func NewApplicationDataCompositeValue() *ApplicationDataCompositeValue {
	hs := NewHStr("Windows.Storage.ApplicationDataCompositeValue")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &ApplicationDataCompositeValue{
		RtClass:      RtClass{PInspect: p},
		IPropertySet: (*IPropertySet)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type ApplicationDataContainer struct {
	RtClass
	*IApplicationDataContainer
}

type ApplicationDataContainerSettings struct {
	RtClass
	*IPropertySet
}

type SetVersionDeferral struct {
	RtClass
	*ISetVersionDeferral
}

type SetVersionRequest struct {
	RtClass
	*ISetVersionRequest
}
