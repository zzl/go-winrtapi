package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// interfaces

// 8450902C-15ED-4FAF-93EA-2237BBF8CBD6
var IID_IAppExtension = syscall.GUID{0x8450902C, 0x15ED, 0x4FAF,
	[8]byte{0x93, 0xEA, 0x22, 0x37, 0xBB, 0xF8, 0xCB, 0xD6}}

type IAppExtensionInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	Get_DisplayName() string
	Get_Description() string
	Get_Package() *IPackage
	Get_AppInfo() *IAppInfo
	GetExtensionPropertiesAsync() *IAsyncOperation[*IPropertySet]
	GetPublicFolderAsync() *IAsyncOperation[*IStorageFolder]
}

type IAppExtensionVtbl struct {
	win32.IInspectableVtbl
	Get_Id                      uintptr
	Get_DisplayName             uintptr
	Get_Description             uintptr
	Get_Package                 uintptr
	Get_AppInfo                 uintptr
	GetExtensionPropertiesAsync uintptr
	GetPublicFolderAsync        uintptr
}

type IAppExtension struct {
	win32.IInspectable
}

func (this *IAppExtension) Vtbl() *IAppExtensionVtbl {
	return (*IAppExtensionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppExtension) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppExtension) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppExtension) Get_Description() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Description, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppExtension) Get_Package() *IPackage {
	var _result *IPackage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Package, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppExtension) Get_AppInfo() *IAppInfo {
	var _result *IAppInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppExtension) GetExtensionPropertiesAsync() *IAsyncOperation[*IPropertySet] {
	var _result *IAsyncOperation[*IPropertySet]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetExtensionPropertiesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppExtension) GetPublicFolderAsync() *IAsyncOperation[*IStorageFolder] {
	var _result *IAsyncOperation[*IStorageFolder]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPublicFolderAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 97872032-8426-4AD1-9084-92E88C2DA200
var IID_IAppExtensionCatalog = syscall.GUID{0x97872032, 0x8426, 0x4AD1,
	[8]byte{0x90, 0x84, 0x92, 0xE8, 0x8C, 0x2D, 0xA2, 0x00}}

type IAppExtensionCatalogInterface interface {
	win32.IInspectableInterface
	FindAllAsync() *IAsyncOperation[*IVectorView[*IAppExtension]]
	RequestRemovePackageAsync(packageFullName string) *IAsyncOperation[bool]
	Add_PackageInstalled(handler TypedEventHandler[*IAppExtensionCatalog, *IAppExtensionPackageInstalledEventArgs]) EventRegistrationToken
	Remove_PackageInstalled(token EventRegistrationToken)
	Add_PackageUpdating(handler TypedEventHandler[*IAppExtensionCatalog, *IAppExtensionPackageUpdatingEventArgs]) EventRegistrationToken
	Remove_PackageUpdating(token EventRegistrationToken)
	Add_PackageUpdated(handler TypedEventHandler[*IAppExtensionCatalog, *IAppExtensionPackageUpdatedEventArgs]) EventRegistrationToken
	Remove_PackageUpdated(token EventRegistrationToken)
	Add_PackageUninstalling(handler TypedEventHandler[*IAppExtensionCatalog, *IAppExtensionPackageUninstallingEventArgs]) EventRegistrationToken
	Remove_PackageUninstalling(token EventRegistrationToken)
	Add_PackageStatusChanged(handler TypedEventHandler[*IAppExtensionCatalog, *IAppExtensionPackageStatusChangedEventArgs]) EventRegistrationToken
	Remove_PackageStatusChanged(token EventRegistrationToken)
}

type IAppExtensionCatalogVtbl struct {
	win32.IInspectableVtbl
	FindAllAsync                uintptr
	RequestRemovePackageAsync   uintptr
	Add_PackageInstalled        uintptr
	Remove_PackageInstalled     uintptr
	Add_PackageUpdating         uintptr
	Remove_PackageUpdating      uintptr
	Add_PackageUpdated          uintptr
	Remove_PackageUpdated       uintptr
	Add_PackageUninstalling     uintptr
	Remove_PackageUninstalling  uintptr
	Add_PackageStatusChanged    uintptr
	Remove_PackageStatusChanged uintptr
}

type IAppExtensionCatalog struct {
	win32.IInspectable
}

func (this *IAppExtensionCatalog) Vtbl() *IAppExtensionCatalogVtbl {
	return (*IAppExtensionCatalogVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppExtensionCatalog) FindAllAsync() *IAsyncOperation[*IVectorView[*IAppExtension]] {
	var _result *IAsyncOperation[*IVectorView[*IAppExtension]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppExtensionCatalog) RequestRemovePackageAsync(packageFullName string) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestRemovePackageAsync, uintptr(unsafe.Pointer(this)), NewHStr(packageFullName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppExtensionCatalog) Add_PackageInstalled(handler TypedEventHandler[*IAppExtensionCatalog, *IAppExtensionPackageInstalledEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PackageInstalled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppExtensionCatalog) Remove_PackageInstalled(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PackageInstalled, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppExtensionCatalog) Add_PackageUpdating(handler TypedEventHandler[*IAppExtensionCatalog, *IAppExtensionPackageUpdatingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PackageUpdating, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppExtensionCatalog) Remove_PackageUpdating(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PackageUpdating, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppExtensionCatalog) Add_PackageUpdated(handler TypedEventHandler[*IAppExtensionCatalog, *IAppExtensionPackageUpdatedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PackageUpdated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppExtensionCatalog) Remove_PackageUpdated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PackageUpdated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppExtensionCatalog) Add_PackageUninstalling(handler TypedEventHandler[*IAppExtensionCatalog, *IAppExtensionPackageUninstallingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PackageUninstalling, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppExtensionCatalog) Remove_PackageUninstalling(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PackageUninstalling, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppExtensionCatalog) Add_PackageStatusChanged(handler TypedEventHandler[*IAppExtensionCatalog, *IAppExtensionPackageStatusChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PackageStatusChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppExtensionCatalog) Remove_PackageStatusChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PackageStatusChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 3C36668A-5F18-4F0B-9CE5-CAB61D196F11
var IID_IAppExtensionCatalogStatics = syscall.GUID{0x3C36668A, 0x5F18, 0x4F0B,
	[8]byte{0x9C, 0xE5, 0xCA, 0xB6, 0x1D, 0x19, 0x6F, 0x11}}

type IAppExtensionCatalogStaticsInterface interface {
	win32.IInspectableInterface
	Open(appExtensionName string) *IAppExtensionCatalog
}

type IAppExtensionCatalogStaticsVtbl struct {
	win32.IInspectableVtbl
	Open uintptr
}

type IAppExtensionCatalogStatics struct {
	win32.IInspectable
}

func (this *IAppExtensionCatalogStatics) Vtbl() *IAppExtensionCatalogStaticsVtbl {
	return (*IAppExtensionCatalogStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppExtensionCatalogStatics) Open(appExtensionName string) *IAppExtensionCatalog {
	var _result *IAppExtensionCatalog
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Open, uintptr(unsafe.Pointer(this)), NewHStr(appExtensionName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 39E59234-3351-4A8D-9745-E7D3DD45BC48
var IID_IAppExtensionPackageInstalledEventArgs = syscall.GUID{0x39E59234, 0x3351, 0x4A8D,
	[8]byte{0x97, 0x45, 0xE7, 0xD3, 0xDD, 0x45, 0xBC, 0x48}}

type IAppExtensionPackageInstalledEventArgsInterface interface {
	win32.IInspectableInterface
	Get_AppExtensionName() string
	Get_Package() *IPackage
	Get_Extensions() *IVectorView[*IAppExtension]
}

type IAppExtensionPackageInstalledEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_AppExtensionName uintptr
	Get_Package          uintptr
	Get_Extensions       uintptr
}

type IAppExtensionPackageInstalledEventArgs struct {
	win32.IInspectable
}

func (this *IAppExtensionPackageInstalledEventArgs) Vtbl() *IAppExtensionPackageInstalledEventArgsVtbl {
	return (*IAppExtensionPackageInstalledEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppExtensionPackageInstalledEventArgs) Get_AppExtensionName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppExtensionName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppExtensionPackageInstalledEventArgs) Get_Package() *IPackage {
	var _result *IPackage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Package, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppExtensionPackageInstalledEventArgs) Get_Extensions() *IVectorView[*IAppExtension] {
	var _result *IVectorView[*IAppExtension]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Extensions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1CE17433-1153-44FD-87B1-8AE1050303DF
var IID_IAppExtensionPackageStatusChangedEventArgs = syscall.GUID{0x1CE17433, 0x1153, 0x44FD,
	[8]byte{0x87, 0xB1, 0x8A, 0xE1, 0x05, 0x03, 0x03, 0xDF}}

type IAppExtensionPackageStatusChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_AppExtensionName() string
	Get_Package() *IPackage
}

type IAppExtensionPackageStatusChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_AppExtensionName uintptr
	Get_Package          uintptr
}

type IAppExtensionPackageStatusChangedEventArgs struct {
	win32.IInspectable
}

func (this *IAppExtensionPackageStatusChangedEventArgs) Vtbl() *IAppExtensionPackageStatusChangedEventArgsVtbl {
	return (*IAppExtensionPackageStatusChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppExtensionPackageStatusChangedEventArgs) Get_AppExtensionName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppExtensionName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppExtensionPackageStatusChangedEventArgs) Get_Package() *IPackage {
	var _result *IPackage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Package, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 60F160C5-171E-40FF-AE98-AB2C20DD4D75
var IID_IAppExtensionPackageUninstallingEventArgs = syscall.GUID{0x60F160C5, 0x171E, 0x40FF,
	[8]byte{0xAE, 0x98, 0xAB, 0x2C, 0x20, 0xDD, 0x4D, 0x75}}

type IAppExtensionPackageUninstallingEventArgsInterface interface {
	win32.IInspectableInterface
	Get_AppExtensionName() string
	Get_Package() *IPackage
}

type IAppExtensionPackageUninstallingEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_AppExtensionName uintptr
	Get_Package          uintptr
}

type IAppExtensionPackageUninstallingEventArgs struct {
	win32.IInspectable
}

func (this *IAppExtensionPackageUninstallingEventArgs) Vtbl() *IAppExtensionPackageUninstallingEventArgsVtbl {
	return (*IAppExtensionPackageUninstallingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppExtensionPackageUninstallingEventArgs) Get_AppExtensionName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppExtensionName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppExtensionPackageUninstallingEventArgs) Get_Package() *IPackage {
	var _result *IPackage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Package, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3A83C43F-797E-44B5-BA24-A4C8B5A543D7
var IID_IAppExtensionPackageUpdatedEventArgs = syscall.GUID{0x3A83C43F, 0x797E, 0x44B5,
	[8]byte{0xBA, 0x24, 0xA4, 0xC8, 0xB5, 0xA5, 0x43, 0xD7}}

type IAppExtensionPackageUpdatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_AppExtensionName() string
	Get_Package() *IPackage
	Get_Extensions() *IVectorView[*IAppExtension]
}

type IAppExtensionPackageUpdatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_AppExtensionName uintptr
	Get_Package          uintptr
	Get_Extensions       uintptr
}

type IAppExtensionPackageUpdatedEventArgs struct {
	win32.IInspectable
}

func (this *IAppExtensionPackageUpdatedEventArgs) Vtbl() *IAppExtensionPackageUpdatedEventArgsVtbl {
	return (*IAppExtensionPackageUpdatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppExtensionPackageUpdatedEventArgs) Get_AppExtensionName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppExtensionName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppExtensionPackageUpdatedEventArgs) Get_Package() *IPackage {
	var _result *IPackage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Package, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppExtensionPackageUpdatedEventArgs) Get_Extensions() *IVectorView[*IAppExtension] {
	var _result *IVectorView[*IAppExtension]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Extensions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7ED59329-1A65-4800-A700-B321009E306A
var IID_IAppExtensionPackageUpdatingEventArgs = syscall.GUID{0x7ED59329, 0x1A65, 0x4800,
	[8]byte{0xA7, 0x00, 0xB3, 0x21, 0x00, 0x9E, 0x30, 0x6A}}

type IAppExtensionPackageUpdatingEventArgsInterface interface {
	win32.IInspectableInterface
	Get_AppExtensionName() string
	Get_Package() *IPackage
}

type IAppExtensionPackageUpdatingEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_AppExtensionName uintptr
	Get_Package          uintptr
}

type IAppExtensionPackageUpdatingEventArgs struct {
	win32.IInspectable
}

func (this *IAppExtensionPackageUpdatingEventArgs) Vtbl() *IAppExtensionPackageUpdatingEventArgsVtbl {
	return (*IAppExtensionPackageUpdatingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppExtensionPackageUpdatingEventArgs) Get_AppExtensionName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppExtensionName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppExtensionPackageUpdatingEventArgs) Get_Package() *IPackage {
	var _result *IPackage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Package, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type AppExtension struct {
	RtClass
	*IAppExtension
}

type AppExtensionPackageInstalledEventArgs struct {
	RtClass
	*IAppExtensionPackageInstalledEventArgs
}

type AppExtensionPackageStatusChangedEventArgs struct {
	RtClass
	*IAppExtensionPackageStatusChangedEventArgs
}

type AppExtensionPackageUninstallingEventArgs struct {
	RtClass
	*IAppExtensionPackageUninstallingEventArgs
}

type AppExtensionPackageUpdatedEventArgs struct {
	RtClass
	*IAppExtensionPackageUpdatedEventArgs
}

type AppExtensionPackageUpdatingEventArgs struct {
	RtClass
	*IAppExtensionPackageUpdatingEventArgs
}
