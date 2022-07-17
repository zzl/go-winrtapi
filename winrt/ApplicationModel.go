package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
// flags
type AddResourcePackageOptions uint32

const (
	AddResourcePackageOptions_None                   AddResourcePackageOptions = 0
	AddResourcePackageOptions_ForceTargetAppShutdown AddResourcePackageOptions = 1
	AddResourcePackageOptions_ApplyUpdateIfAvailable AddResourcePackageOptions = 2
)

// enum
type AppExecutionContext int32

const (
	AppExecutionContext_Unknown AppExecutionContext = 0
	AppExecutionContext_Host    AppExecutionContext = 1
	AppExecutionContext_Guest   AppExecutionContext = 2
)

// enum
type AppInstallerPolicySource int32

const (
	AppInstallerPolicySource_Default AppInstallerPolicySource = 0
	AppInstallerPolicySource_System  AppInstallerPolicySource = 1
)

// enum
type FullTrustLaunchResult int32

const (
	FullTrustLaunchResult_Success      FullTrustLaunchResult = 0
	FullTrustLaunchResult_AccessDenied FullTrustLaunchResult = 1
	FullTrustLaunchResult_FileNotFound FullTrustLaunchResult = 2
	FullTrustLaunchResult_Unknown      FullTrustLaunchResult = 3
)

// enum
type LimitedAccessFeatureStatus int32

const (
	LimitedAccessFeatureStatus_Unavailable           LimitedAccessFeatureStatus = 0
	LimitedAccessFeatureStatus_Available             LimitedAccessFeatureStatus = 1
	LimitedAccessFeatureStatus_AvailableWithoutToken LimitedAccessFeatureStatus = 2
	LimitedAccessFeatureStatus_Unknown               LimitedAccessFeatureStatus = 3
)

// enum
type PackageContentGroupState int32

const (
	PackageContentGroupState_NotStaged PackageContentGroupState = 0
	PackageContentGroupState_Queued    PackageContentGroupState = 1
	PackageContentGroupState_Staging   PackageContentGroupState = 2
	PackageContentGroupState_Staged    PackageContentGroupState = 3
)

// enum
type PackageRelationship int32

const (
	PackageRelationship_Dependencies PackageRelationship = 0
	PackageRelationship_Dependents   PackageRelationship = 1
	PackageRelationship_All          PackageRelationship = 2
)

// enum
type PackageSignatureKind int32

const (
	PackageSignatureKind_None       PackageSignatureKind = 0
	PackageSignatureKind_Developer  PackageSignatureKind = 1
	PackageSignatureKind_Enterprise PackageSignatureKind = 2
	PackageSignatureKind_Store      PackageSignatureKind = 3
	PackageSignatureKind_System     PackageSignatureKind = 4
)

// enum
type PackageUpdateAvailability int32

const (
	PackageUpdateAvailability_Unknown   PackageUpdateAvailability = 0
	PackageUpdateAvailability_NoUpdates PackageUpdateAvailability = 1
	PackageUpdateAvailability_Available PackageUpdateAvailability = 2
	PackageUpdateAvailability_Required  PackageUpdateAvailability = 3
	PackageUpdateAvailability_Error     PackageUpdateAvailability = 4
)

// enum
type StartupTaskState int32

const (
	StartupTaskState_Disabled         StartupTaskState = 0
	StartupTaskState_DisabledByUser   StartupTaskState = 1
	StartupTaskState_Enabled          StartupTaskState = 2
	StartupTaskState_DisabledByPolicy StartupTaskState = 3
	StartupTaskState_EnabledByPolicy  StartupTaskState = 4
)

// structs

type FullTrustAppContract struct {
}

type PackageInstallProgress struct {
	PercentComplete uint32
}

type PackageVersion struct {
	Major    uint16
	Minor    uint16
	Build    uint16
	Revision uint16
}

type StartupTaskContract struct {
}

// interfaces

// 1AEB1103-E4D4-41AA-A4F6-C4A276E79EAC
var IID_IAppDisplayInfo = syscall.GUID{0x1AEB1103, 0xE4D4, 0x41AA,
	[8]byte{0xA4, 0xF6, 0xC4, 0xA2, 0x76, 0xE7, 0x9E, 0xAC}}

type IAppDisplayInfoInterface interface {
	win32.IInspectableInterface
	Get_DisplayName() string
	Get_Description() string
	GetLogo(size Size) *IRandomAccessStreamReference
}

type IAppDisplayInfoVtbl struct {
	win32.IInspectableVtbl
	Get_DisplayName uintptr
	Get_Description uintptr
	GetLogo         uintptr
}

type IAppDisplayInfo struct {
	win32.IInspectable
}

func (this *IAppDisplayInfo) Vtbl() *IAppDisplayInfoVtbl {
	return (*IAppDisplayInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppDisplayInfo) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppDisplayInfo) Get_Description() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Description, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppDisplayInfo) GetLogo(size Size) *IRandomAccessStreamReference {
	var _result *IRandomAccessStreamReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetLogo, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&size)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CF7F59B3-6A09-4DE8-A6C0-5792D56880D1
var IID_IAppInfo = syscall.GUID{0xCF7F59B3, 0x6A09, 0x4DE8,
	[8]byte{0xA6, 0xC0, 0x57, 0x92, 0xD5, 0x68, 0x80, 0xD1}}

type IAppInfoInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	Get_AppUserModelId() string
	Get_DisplayInfo() *IAppDisplayInfo
	Get_PackageFamilyName() string
}

type IAppInfoVtbl struct {
	win32.IInspectableVtbl
	Get_Id                uintptr
	Get_AppUserModelId    uintptr
	Get_DisplayInfo       uintptr
	Get_PackageFamilyName uintptr
}

type IAppInfo struct {
	win32.IInspectable
}

func (this *IAppInfo) Vtbl() *IAppInfoVtbl {
	return (*IAppInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppInfo) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppInfo) Get_AppUserModelId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppUserModelId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppInfo) Get_DisplayInfo() *IAppDisplayInfo {
	var _result *IAppDisplayInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppInfo) Get_PackageFamilyName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PackageFamilyName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// BE4B1F5A-2098-431B-BD25-B30878748D47
var IID_IAppInfo2 = syscall.GUID{0xBE4B1F5A, 0x2098, 0x431B,
	[8]byte{0xBD, 0x25, 0xB3, 0x08, 0x78, 0x74, 0x8D, 0x47}}

type IAppInfo2Interface interface {
	win32.IInspectableInterface
	Get_Package() *IPackage
}

type IAppInfo2Vtbl struct {
	win32.IInspectableVtbl
	Get_Package uintptr
}

type IAppInfo2 struct {
	win32.IInspectable
}

func (this *IAppInfo2) Vtbl() *IAppInfo2Vtbl {
	return (*IAppInfo2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppInfo2) Get_Package() *IPackage {
	var _result *IPackage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Package, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 09A78E46-93A4-46DE-9397-0843B57115EA
var IID_IAppInfo3 = syscall.GUID{0x09A78E46, 0x93A4, 0x46DE,
	[8]byte{0x93, 0x97, 0x08, 0x43, 0xB5, 0x71, 0x15, 0xEA}}

type IAppInfo3Interface interface {
	win32.IInspectableInterface
	Get_ExecutionContext() AppExecutionContext
}

type IAppInfo3Vtbl struct {
	win32.IInspectableVtbl
	Get_ExecutionContext uintptr
}

type IAppInfo3 struct {
	win32.IInspectable
}

func (this *IAppInfo3) Vtbl() *IAppInfo3Vtbl {
	return (*IAppInfo3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppInfo3) Get_ExecutionContext() AppExecutionContext {
	var _result AppExecutionContext
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExecutionContext, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 2F34BDEB-1609-4554-9F33-12E1E803E0D4
var IID_IAppInfo4 = syscall.GUID{0x2F34BDEB, 0x1609, 0x4554,
	[8]byte{0x9F, 0x33, 0x12, 0xE1, 0xE8, 0x03, 0xE0, 0xD4}}

type IAppInfo4Interface interface {
	win32.IInspectableInterface
	Get_SupportedFileExtensions() []string
}

type IAppInfo4Vtbl struct {
	win32.IInspectableVtbl
	Get_SupportedFileExtensions uintptr
}

type IAppInfo4 struct {
	win32.IInspectable
}

func (this *IAppInfo4) Vtbl() *IAppInfo4Vtbl {
	return (*IAppInfo4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppInfo4) Get_SupportedFileExtensions() []string {
	var _result []string
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedFileExtensions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// CF1F782A-E48B-4F0C-9B0B-79C3F8957DD7
var IID_IAppInfoStatics = syscall.GUID{0xCF1F782A, 0xE48B, 0x4F0C,
	[8]byte{0x9B, 0x0B, 0x79, 0xC3, 0xF8, 0x95, 0x7D, 0xD7}}

type IAppInfoStaticsInterface interface {
	win32.IInspectableInterface
	Get_Current() *IAppInfo
	GetFromAppUserModelId(appUserModelId string) *IAppInfo
	GetFromAppUserModelIdForUser(user *IUser, appUserModelId string) *IAppInfo
}

type IAppInfoStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Current                  uintptr
	GetFromAppUserModelId        uintptr
	GetFromAppUserModelIdForUser uintptr
}

type IAppInfoStatics struct {
	win32.IInspectable
}

func (this *IAppInfoStatics) Vtbl() *IAppInfoStaticsVtbl {
	return (*IAppInfoStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppInfoStatics) Get_Current() *IAppInfo {
	var _result *IAppInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Current, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppInfoStatics) GetFromAppUserModelId(appUserModelId string) *IAppInfo {
	var _result *IAppInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFromAppUserModelId, uintptr(unsafe.Pointer(this)), NewHStr(appUserModelId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppInfoStatics) GetFromAppUserModelIdForUser(user *IUser, appUserModelId string) *IAppInfo {
	var _result *IAppInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFromAppUserModelIdForUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), NewHStr(appUserModelId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 29AB2AC0-D4F6-42A3-ADCD-D6583C659508
var IID_IAppInstallerInfo = syscall.GUID{0x29AB2AC0, 0xD4F6, 0x42A3,
	[8]byte{0xAD, 0xCD, 0xD6, 0x58, 0x3C, 0x65, 0x95, 0x08}}

type IAppInstallerInfoInterface interface {
	win32.IInspectableInterface
	Get_Uri() *IUriRuntimeClass
}

type IAppInstallerInfoVtbl struct {
	win32.IInspectableVtbl
	Get_Uri uintptr
}

type IAppInstallerInfo struct {
	win32.IInspectable
}

func (this *IAppInstallerInfo) Vtbl() *IAppInstallerInfoVtbl {
	return (*IAppInstallerInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppInstallerInfo) Get_Uri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Uri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D20F1388-8256-597C-8511-C84EC50D5E2B
var IID_IAppInstallerInfo2 = syscall.GUID{0xD20F1388, 0x8256, 0x597C,
	[8]byte{0x85, 0x11, 0xC8, 0x4E, 0xC5, 0x0D, 0x5E, 0x2B}}

type IAppInstallerInfo2Interface interface {
	win32.IInspectableInterface
	Get_OnLaunch() bool
	Get_HoursBetweenUpdateChecks() uint32
	Get_ShowPrompt() bool
	Get_UpdateBlocksActivation() bool
	Get_AutomaticBackgroundTask() bool
	Get_ForceUpdateFromAnyVersion() bool
	Get_IsAutoRepairEnabled() bool
	Get_Version() PackageVersion
	Get_LastChecked() DateTime
	Get_PausedUntil() *IReference[DateTime]
	Get_UpdateUris() *IVectorView[*IUriRuntimeClass]
	Get_RepairUris() *IVectorView[*IUriRuntimeClass]
	Get_DependencyPackageUris() *IVectorView[*IUriRuntimeClass]
	Get_OptionalPackageUris() *IVectorView[*IUriRuntimeClass]
	Get_PolicySource() AppInstallerPolicySource
}

type IAppInstallerInfo2Vtbl struct {
	win32.IInspectableVtbl
	Get_OnLaunch                  uintptr
	Get_HoursBetweenUpdateChecks  uintptr
	Get_ShowPrompt                uintptr
	Get_UpdateBlocksActivation    uintptr
	Get_AutomaticBackgroundTask   uintptr
	Get_ForceUpdateFromAnyVersion uintptr
	Get_IsAutoRepairEnabled       uintptr
	Get_Version                   uintptr
	Get_LastChecked               uintptr
	Get_PausedUntil               uintptr
	Get_UpdateUris                uintptr
	Get_RepairUris                uintptr
	Get_DependencyPackageUris     uintptr
	Get_OptionalPackageUris       uintptr
	Get_PolicySource              uintptr
}

type IAppInstallerInfo2 struct {
	win32.IInspectable
}

func (this *IAppInstallerInfo2) Vtbl() *IAppInstallerInfo2Vtbl {
	return (*IAppInstallerInfo2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppInstallerInfo2) Get_OnLaunch() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OnLaunch, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppInstallerInfo2) Get_HoursBetweenUpdateChecks() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HoursBetweenUpdateChecks, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppInstallerInfo2) Get_ShowPrompt() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ShowPrompt, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppInstallerInfo2) Get_UpdateBlocksActivation() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UpdateBlocksActivation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppInstallerInfo2) Get_AutomaticBackgroundTask() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AutomaticBackgroundTask, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppInstallerInfo2) Get_ForceUpdateFromAnyVersion() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ForceUpdateFromAnyVersion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppInstallerInfo2) Get_IsAutoRepairEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsAutoRepairEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppInstallerInfo2) Get_Version() PackageVersion {
	var _result PackageVersion
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Version, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppInstallerInfo2) Get_LastChecked() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LastChecked, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppInstallerInfo2) Get_PausedUntil() *IReference[DateTime] {
	var _result *IReference[DateTime]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PausedUntil, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppInstallerInfo2) Get_UpdateUris() *IVectorView[*IUriRuntimeClass] {
	var _result *IVectorView[*IUriRuntimeClass]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UpdateUris, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppInstallerInfo2) Get_RepairUris() *IVectorView[*IUriRuntimeClass] {
	var _result *IVectorView[*IUriRuntimeClass]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RepairUris, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppInstallerInfo2) Get_DependencyPackageUris() *IVectorView[*IUriRuntimeClass] {
	var _result *IVectorView[*IUriRuntimeClass]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DependencyPackageUris, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppInstallerInfo2) Get_OptionalPackageUris() *IVectorView[*IUriRuntimeClass] {
	var _result *IVectorView[*IUriRuntimeClass]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OptionalPackageUris, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppInstallerInfo2) Get_PolicySource() AppInstallerPolicySource {
	var _result AppInstallerPolicySource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PolicySource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 675F2B47-F25F-4532-9FD6-3633E0634D01
var IID_IAppInstance = syscall.GUID{0x675F2B47, 0xF25F, 0x4532,
	[8]byte{0x9F, 0xD6, 0x36, 0x33, 0xE0, 0x63, 0x4D, 0x01}}

type IAppInstanceInterface interface {
	win32.IInspectableInterface
	Get_Key() string
	Get_IsCurrentInstance() bool
	RedirectActivationTo()
}

type IAppInstanceVtbl struct {
	win32.IInspectableVtbl
	Get_Key               uintptr
	Get_IsCurrentInstance uintptr
	RedirectActivationTo  uintptr
}

type IAppInstance struct {
	win32.IInspectable
}

func (this *IAppInstance) Vtbl() *IAppInstanceVtbl {
	return (*IAppInstanceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppInstance) Get_Key() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Key, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppInstance) Get_IsCurrentInstance() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCurrentInstance, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppInstance) RedirectActivationTo() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RedirectActivationTo, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 9D11E77F-9EA6-47AF-A6EC-46784C5BA254
var IID_IAppInstanceStatics = syscall.GUID{0x9D11E77F, 0x9EA6, 0x47AF,
	[8]byte{0xA6, 0xEC, 0x46, 0x78, 0x4C, 0x5B, 0xA2, 0x54}}

type IAppInstanceStaticsInterface interface {
	win32.IInspectableInterface
	Get_RecommendedInstance() *IAppInstance
	GetActivatedEventArgs() *IActivatedEventArgs
	FindOrRegisterInstanceForKey(key string) *IAppInstance
	Unregister()
	GetInstances() *IVector[*IAppInstance]
}

type IAppInstanceStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_RecommendedInstance      uintptr
	GetActivatedEventArgs        uintptr
	FindOrRegisterInstanceForKey uintptr
	Unregister                   uintptr
	GetInstances                 uintptr
}

type IAppInstanceStatics struct {
	win32.IInspectable
}

func (this *IAppInstanceStatics) Vtbl() *IAppInstanceStaticsVtbl {
	return (*IAppInstanceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppInstanceStatics) Get_RecommendedInstance() *IAppInstance {
	var _result *IAppInstance
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RecommendedInstance, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppInstanceStatics) GetActivatedEventArgs() *IActivatedEventArgs {
	var _result *IActivatedEventArgs
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetActivatedEventArgs, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppInstanceStatics) FindOrRegisterInstanceForKey(key string) *IAppInstance {
	var _result *IAppInstance
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindOrRegisterInstanceForKey, uintptr(unsafe.Pointer(this)), NewHStr(key).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppInstanceStatics) Unregister() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Unregister, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IAppInstanceStatics) GetInstances() *IVector[*IAppInstance] {
	var _result *IVector[*IAppInstance]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetInstances, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9599DDCE-9BD3-435C-8054-C1ADD50028FE
var IID_ICameraApplicationManagerStatics = syscall.GUID{0x9599DDCE, 0x9BD3, 0x435C,
	[8]byte{0x80, 0x54, 0xC1, 0xAD, 0xD5, 0x00, 0x28, 0xFE}}

type ICameraApplicationManagerStaticsInterface interface {
	win32.IInspectableInterface
	ShowInstalledApplicationsUI()
}

type ICameraApplicationManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	ShowInstalledApplicationsUI uintptr
}

type ICameraApplicationManagerStatics struct {
	win32.IInspectable
}

func (this *ICameraApplicationManagerStatics) Vtbl() *ICameraApplicationManagerStaticsVtbl {
	return (*ICameraApplicationManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICameraApplicationManagerStatics) ShowInstalledApplicationsUI() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowInstalledApplicationsUI, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 2C3893CC-F81A-4E7A-B857-76A80887E185
var IID_IDesignModeStatics = syscall.GUID{0x2C3893CC, 0xF81A, 0x4E7A,
	[8]byte{0xB8, 0x57, 0x76, 0xA8, 0x08, 0x87, 0xE1, 0x85}}

type IDesignModeStaticsInterface interface {
	win32.IInspectableInterface
	Get_DesignModeEnabled() bool
}

type IDesignModeStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_DesignModeEnabled uintptr
}

type IDesignModeStatics struct {
	win32.IInspectable
}

func (this *IDesignModeStatics) Vtbl() *IDesignModeStaticsVtbl {
	return (*IDesignModeStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDesignModeStatics) Get_DesignModeEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesignModeEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 80CF8137-B064-4858-BEC8-3EBA22357535
var IID_IDesignModeStatics2 = syscall.GUID{0x80CF8137, 0xB064, 0x4858,
	[8]byte{0xBE, 0xC8, 0x3E, 0xBA, 0x22, 0x35, 0x75, 0x35}}

type IDesignModeStatics2Interface interface {
	win32.IInspectableInterface
	Get_DesignMode2Enabled() bool
}

type IDesignModeStatics2Vtbl struct {
	win32.IInspectableVtbl
	Get_DesignMode2Enabled uintptr
}

type IDesignModeStatics2 struct {
	win32.IInspectable
}

func (this *IDesignModeStatics2) Vtbl() *IDesignModeStatics2Vtbl {
	return (*IDesignModeStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDesignModeStatics2) Get_DesignMode2Enabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesignMode2Enabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F722DCC2-9827-403D-AAED-ECCA9AC17398
var IID_IEnteredBackgroundEventArgs = syscall.GUID{0xF722DCC2, 0x9827, 0x403D,
	[8]byte{0xAA, 0xED, 0xEC, 0xCA, 0x9A, 0xC1, 0x73, 0x98}}

type IEnteredBackgroundEventArgsInterface interface {
	win32.IInspectableInterface
	GetDeferral() *IDeferral
}

type IEnteredBackgroundEventArgsVtbl struct {
	win32.IInspectableVtbl
	GetDeferral uintptr
}

type IEnteredBackgroundEventArgs struct {
	win32.IInspectable
}

func (this *IEnteredBackgroundEventArgs) Vtbl() *IEnteredBackgroundEventArgsVtbl {
	return (*IEnteredBackgroundEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IEnteredBackgroundEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 41DD7EEA-B335-521F-B96C-5EA07F5B7329
var IID_IFindRelatedPackagesOptions = syscall.GUID{0x41DD7EEA, 0xB335, 0x521F,
	[8]byte{0xB9, 0x6C, 0x5E, 0xA0, 0x7F, 0x5B, 0x73, 0x29}}

type IFindRelatedPackagesOptionsInterface interface {
	win32.IInspectableInterface
	Get_Relationship() PackageRelationship
	Put_Relationship(value PackageRelationship)
	Get_IncludeFrameworks() bool
	Put_IncludeFrameworks(value bool)
	Get_IncludeHostRuntimes() bool
	Put_IncludeHostRuntimes(value bool)
	Get_IncludeOptionals() bool
	Put_IncludeOptionals(value bool)
	Get_IncludeResources() bool
	Put_IncludeResources(value bool)
}

type IFindRelatedPackagesOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_Relationship        uintptr
	Put_Relationship        uintptr
	Get_IncludeFrameworks   uintptr
	Put_IncludeFrameworks   uintptr
	Get_IncludeHostRuntimes uintptr
	Put_IncludeHostRuntimes uintptr
	Get_IncludeOptionals    uintptr
	Put_IncludeOptionals    uintptr
	Get_IncludeResources    uintptr
	Put_IncludeResources    uintptr
}

type IFindRelatedPackagesOptions struct {
	win32.IInspectable
}

func (this *IFindRelatedPackagesOptions) Vtbl() *IFindRelatedPackagesOptionsVtbl {
	return (*IFindRelatedPackagesOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFindRelatedPackagesOptions) Get_Relationship() PackageRelationship {
	var _result PackageRelationship
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Relationship, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFindRelatedPackagesOptions) Put_Relationship(value PackageRelationship) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Relationship, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IFindRelatedPackagesOptions) Get_IncludeFrameworks() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IncludeFrameworks, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFindRelatedPackagesOptions) Put_IncludeFrameworks(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IncludeFrameworks, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IFindRelatedPackagesOptions) Get_IncludeHostRuntimes() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IncludeHostRuntimes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFindRelatedPackagesOptions) Put_IncludeHostRuntimes(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IncludeHostRuntimes, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IFindRelatedPackagesOptions) Get_IncludeOptionals() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IncludeOptionals, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFindRelatedPackagesOptions) Put_IncludeOptionals(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IncludeOptionals, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IFindRelatedPackagesOptions) Get_IncludeResources() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IncludeResources, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFindRelatedPackagesOptions) Put_IncludeResources(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IncludeResources, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// D7D17254-A4FD-55C4-98CF-F2710B7D8BE2
var IID_IFindRelatedPackagesOptionsFactory = syscall.GUID{0xD7D17254, 0xA4FD, 0x55C4,
	[8]byte{0x98, 0xCF, 0xF2, 0x71, 0x0B, 0x7D, 0x8B, 0xE2}}

type IFindRelatedPackagesOptionsFactoryInterface interface {
	win32.IInspectableInterface
	CreateInstance(Relationship PackageRelationship) *IFindRelatedPackagesOptions
}

type IFindRelatedPackagesOptionsFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateInstance uintptr
}

type IFindRelatedPackagesOptionsFactory struct {
	win32.IInspectable
}

func (this *IFindRelatedPackagesOptionsFactory) Vtbl() *IFindRelatedPackagesOptionsFactoryVtbl {
	return (*IFindRelatedPackagesOptionsFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFindRelatedPackagesOptionsFactory) CreateInstance(Relationship PackageRelationship) *IFindRelatedPackagesOptions {
	var _result *IFindRelatedPackagesOptions
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateInstance, uintptr(unsafe.Pointer(this)), uintptr(Relationship), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8917D888-EDFB-515F-8E22-5EBCEB69DFD9
var IID_IFullTrustProcessLaunchResult = syscall.GUID{0x8917D888, 0xEDFB, 0x515F,
	[8]byte{0x8E, 0x22, 0x5E, 0xBC, 0xEB, 0x69, 0xDF, 0xD9}}

type IFullTrustProcessLaunchResultInterface interface {
	win32.IInspectableInterface
	Get_LaunchResult() FullTrustLaunchResult
	Get_ExtendedError() HResult
}

type IFullTrustProcessLaunchResultVtbl struct {
	win32.IInspectableVtbl
	Get_LaunchResult  uintptr
	Get_ExtendedError uintptr
}

type IFullTrustProcessLaunchResult struct {
	win32.IInspectable
}

func (this *IFullTrustProcessLaunchResult) Vtbl() *IFullTrustProcessLaunchResultVtbl {
	return (*IFullTrustProcessLaunchResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFullTrustProcessLaunchResult) Get_LaunchResult() FullTrustLaunchResult {
	var _result FullTrustLaunchResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LaunchResult, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFullTrustProcessLaunchResult) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D784837F-1100-3C6B-A455-F6262CC331B6
var IID_IFullTrustProcessLauncherStatics = syscall.GUID{0xD784837F, 0x1100, 0x3C6B,
	[8]byte{0xA4, 0x55, 0xF6, 0x26, 0x2C, 0xC3, 0x31, 0xB6}}

type IFullTrustProcessLauncherStaticsInterface interface {
	win32.IInspectableInterface
	LaunchFullTrustProcessForCurrentAppAsync() *IAsyncAction
	LaunchFullTrustProcessForCurrentAppWithParametersAsync(parameterGroupId string) *IAsyncAction
	LaunchFullTrustProcessForAppAsync(fullTrustPackageRelativeAppId string) *IAsyncAction
	LaunchFullTrustProcessForAppWithParametersAsync(fullTrustPackageRelativeAppId string, parameterGroupId string) *IAsyncAction
}

type IFullTrustProcessLauncherStaticsVtbl struct {
	win32.IInspectableVtbl
	LaunchFullTrustProcessForCurrentAppAsync               uintptr
	LaunchFullTrustProcessForCurrentAppWithParametersAsync uintptr
	LaunchFullTrustProcessForAppAsync                      uintptr
	LaunchFullTrustProcessForAppWithParametersAsync        uintptr
}

type IFullTrustProcessLauncherStatics struct {
	win32.IInspectable
}

func (this *IFullTrustProcessLauncherStatics) Vtbl() *IFullTrustProcessLauncherStaticsVtbl {
	return (*IFullTrustProcessLauncherStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFullTrustProcessLauncherStatics) LaunchFullTrustProcessForCurrentAppAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchFullTrustProcessForCurrentAppAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFullTrustProcessLauncherStatics) LaunchFullTrustProcessForCurrentAppWithParametersAsync(parameterGroupId string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchFullTrustProcessForCurrentAppWithParametersAsync, uintptr(unsafe.Pointer(this)), NewHStr(parameterGroupId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFullTrustProcessLauncherStatics) LaunchFullTrustProcessForAppAsync(fullTrustPackageRelativeAppId string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchFullTrustProcessForAppAsync, uintptr(unsafe.Pointer(this)), NewHStr(fullTrustPackageRelativeAppId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFullTrustProcessLauncherStatics) LaunchFullTrustProcessForAppWithParametersAsync(fullTrustPackageRelativeAppId string, parameterGroupId string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchFullTrustProcessForAppWithParametersAsync, uintptr(unsafe.Pointer(this)), NewHStr(fullTrustPackageRelativeAppId).Ptr, NewHStr(parameterGroupId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8B8ED72F-B65C-56CF-A1A7-2BF77CBC6EA8
var IID_IFullTrustProcessLauncherStatics2 = syscall.GUID{0x8B8ED72F, 0xB65C, 0x56CF,
	[8]byte{0xA1, 0xA7, 0x2B, 0xF7, 0x7C, 0xBC, 0x6E, 0xA8}}

type IFullTrustProcessLauncherStatics2Interface interface {
	win32.IInspectableInterface
	LaunchFullTrustProcessForCurrentAppWithArgumentsAsync(commandLine string) *IAsyncOperation[*IFullTrustProcessLaunchResult]
	LaunchFullTrustProcessForAppWithArgumentsAsync(fullTrustPackageRelativeAppId string, commandLine string) *IAsyncOperation[*IFullTrustProcessLaunchResult]
}

type IFullTrustProcessLauncherStatics2Vtbl struct {
	win32.IInspectableVtbl
	LaunchFullTrustProcessForCurrentAppWithArgumentsAsync uintptr
	LaunchFullTrustProcessForAppWithArgumentsAsync        uintptr
}

type IFullTrustProcessLauncherStatics2 struct {
	win32.IInspectable
}

func (this *IFullTrustProcessLauncherStatics2) Vtbl() *IFullTrustProcessLauncherStatics2Vtbl {
	return (*IFullTrustProcessLauncherStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFullTrustProcessLauncherStatics2) LaunchFullTrustProcessForCurrentAppWithArgumentsAsync(commandLine string) *IAsyncOperation[*IFullTrustProcessLaunchResult] {
	var _result *IAsyncOperation[*IFullTrustProcessLaunchResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchFullTrustProcessForCurrentAppWithArgumentsAsync, uintptr(unsafe.Pointer(this)), NewHStr(commandLine).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFullTrustProcessLauncherStatics2) LaunchFullTrustProcessForAppWithArgumentsAsync(fullTrustPackageRelativeAppId string, commandLine string) *IAsyncOperation[*IFullTrustProcessLaunchResult] {
	var _result *IAsyncOperation[*IFullTrustProcessLaunchResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LaunchFullTrustProcessForAppWithArgumentsAsync, uintptr(unsafe.Pointer(this)), NewHStr(fullTrustPackageRelativeAppId).Ptr, NewHStr(commandLine).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 39C6EC9A-AE6E-46F9-A07A-CFC23F88733E
var IID_ILeavingBackgroundEventArgs = syscall.GUID{0x39C6EC9A, 0xAE6E, 0x46F9,
	[8]byte{0xA0, 0x7A, 0xCF, 0xC2, 0x3F, 0x88, 0x73, 0x3E}}

type ILeavingBackgroundEventArgsInterface interface {
	win32.IInspectableInterface
	GetDeferral() *IDeferral
}

type ILeavingBackgroundEventArgsVtbl struct {
	win32.IInspectableVtbl
	GetDeferral uintptr
}

type ILeavingBackgroundEventArgs struct {
	win32.IInspectable
}

func (this *ILeavingBackgroundEventArgs) Vtbl() *ILeavingBackgroundEventArgsVtbl {
	return (*ILeavingBackgroundEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILeavingBackgroundEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D45156A6-1E24-5DDD-ABB4-6188ABA4D5BF
var IID_ILimitedAccessFeatureRequestResult = syscall.GUID{0xD45156A6, 0x1E24, 0x5DDD,
	[8]byte{0xAB, 0xB4, 0x61, 0x88, 0xAB, 0xA4, 0xD5, 0xBF}}

type ILimitedAccessFeatureRequestResultInterface interface {
	win32.IInspectableInterface
	Get_FeatureId() string
	Get_Status() LimitedAccessFeatureStatus
	Get_EstimatedRemovalDate() *IReference[DateTime]
}

type ILimitedAccessFeatureRequestResultVtbl struct {
	win32.IInspectableVtbl
	Get_FeatureId            uintptr
	Get_Status               uintptr
	Get_EstimatedRemovalDate uintptr
}

type ILimitedAccessFeatureRequestResult struct {
	win32.IInspectable
}

func (this *ILimitedAccessFeatureRequestResult) Vtbl() *ILimitedAccessFeatureRequestResultVtbl {
	return (*ILimitedAccessFeatureRequestResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILimitedAccessFeatureRequestResult) Get_FeatureId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FeatureId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILimitedAccessFeatureRequestResult) Get_Status() LimitedAccessFeatureStatus {
	var _result LimitedAccessFeatureStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILimitedAccessFeatureRequestResult) Get_EstimatedRemovalDate() *IReference[DateTime] {
	var _result *IReference[DateTime]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EstimatedRemovalDate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8BE612D4-302B-5FBF-A632-1A99E43E8925
var IID_ILimitedAccessFeaturesStatics = syscall.GUID{0x8BE612D4, 0x302B, 0x5FBF,
	[8]byte{0xA6, 0x32, 0x1A, 0x99, 0xE4, 0x3E, 0x89, 0x25}}

type ILimitedAccessFeaturesStaticsInterface interface {
	win32.IInspectableInterface
	TryUnlockFeature(featureId string, token string, attestation string) *ILimitedAccessFeatureRequestResult
}

type ILimitedAccessFeaturesStaticsVtbl struct {
	win32.IInspectableVtbl
	TryUnlockFeature uintptr
}

type ILimitedAccessFeaturesStatics struct {
	win32.IInspectable
}

func (this *ILimitedAccessFeaturesStatics) Vtbl() *ILimitedAccessFeaturesStaticsVtbl {
	return (*ILimitedAccessFeaturesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILimitedAccessFeaturesStatics) TryUnlockFeature(featureId string, token string, attestation string) *ILimitedAccessFeatureRequestResult {
	var _result *ILimitedAccessFeatureRequestResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryUnlockFeature, uintptr(unsafe.Pointer(this)), NewHStr(featureId).Ptr, NewHStr(token).Ptr, NewHStr(attestation).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 163C792F-BD75-413C-BF23-B1FE7B95D825
var IID_IPackage = syscall.GUID{0x163C792F, 0xBD75, 0x413C,
	[8]byte{0xBF, 0x23, 0xB1, 0xFE, 0x7B, 0x95, 0xD8, 0x25}}

type IPackageInterface interface {
	win32.IInspectableInterface
	Get_Id() *IPackageId
	Get_InstalledLocation() *IStorageFolder
	Get_IsFramework() bool
	Get_Dependencies() *IVectorView[*IPackage]
}

type IPackageVtbl struct {
	win32.IInspectableVtbl
	Get_Id                uintptr
	Get_InstalledLocation uintptr
	Get_IsFramework       uintptr
	Get_Dependencies      uintptr
}

type IPackage struct {
	win32.IInspectable
}

func (this *IPackage) Vtbl() *IPackageVtbl {
	return (*IPackageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackage) Get_Id() *IPackageId {
	var _result *IPackageId
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackage) Get_InstalledLocation() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InstalledLocation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackage) Get_IsFramework() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsFramework, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackage) Get_Dependencies() *IVectorView[*IPackage] {
	var _result *IVectorView[*IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Dependencies, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A6612FB6-7688-4ACE-95FB-359538E7AA01
var IID_IPackage2 = syscall.GUID{0xA6612FB6, 0x7688, 0x4ACE,
	[8]byte{0x95, 0xFB, 0x35, 0x95, 0x38, 0xE7, 0xAA, 0x01}}

type IPackage2Interface interface {
	win32.IInspectableInterface
	Get_DisplayName() string
	Get_PublisherDisplayName() string
	Get_Description() string
	Get_Logo() *IUriRuntimeClass
	Get_IsResourcePackage() bool
	Get_IsBundle() bool
	Get_IsDevelopmentMode() bool
}

type IPackage2Vtbl struct {
	win32.IInspectableVtbl
	Get_DisplayName          uintptr
	Get_PublisherDisplayName uintptr
	Get_Description          uintptr
	Get_Logo                 uintptr
	Get_IsResourcePackage    uintptr
	Get_IsBundle             uintptr
	Get_IsDevelopmentMode    uintptr
}

type IPackage2 struct {
	win32.IInspectable
}

func (this *IPackage2) Vtbl() *IPackage2Vtbl {
	return (*IPackage2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackage2) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPackage2) Get_PublisherDisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PublisherDisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPackage2) Get_Description() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Description, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPackage2) Get_Logo() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Logo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackage2) Get_IsResourcePackage() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsResourcePackage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackage2) Get_IsBundle() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsBundle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackage2) Get_IsDevelopmentMode() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDevelopmentMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5F738B61-F86A-4917-93D1-F1EE9D3B35D9
var IID_IPackage3 = syscall.GUID{0x5F738B61, 0xF86A, 0x4917,
	[8]byte{0x93, 0xD1, 0xF1, 0xEE, 0x9D, 0x3B, 0x35, 0xD9}}

type IPackage3Interface interface {
	win32.IInspectableInterface
	Get_Status() *IPackageStatus
	Get_InstalledDate() DateTime
	GetAppListEntriesAsync() *IAsyncOperation[*IVectorView[*IAppListEntry]]
}

type IPackage3Vtbl struct {
	win32.IInspectableVtbl
	Get_Status             uintptr
	Get_InstalledDate      uintptr
	GetAppListEntriesAsync uintptr
}

type IPackage3 struct {
	win32.IInspectable
}

func (this *IPackage3) Vtbl() *IPackage3Vtbl {
	return (*IPackage3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackage3) Get_Status() *IPackageStatus {
	var _result *IPackageStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackage3) Get_InstalledDate() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InstalledDate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackage3) GetAppListEntriesAsync() *IAsyncOperation[*IVectorView[*IAppListEntry]] {
	var _result *IAsyncOperation[*IVectorView[*IAppListEntry]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAppListEntriesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 65AED1AE-B95B-450C-882B-6255187F397E
var IID_IPackage4 = syscall.GUID{0x65AED1AE, 0xB95B, 0x450C,
	[8]byte{0x88, 0x2B, 0x62, 0x55, 0x18, 0x7F, 0x39, 0x7E}}

type IPackage4Interface interface {
	win32.IInspectableInterface
	Get_SignatureKind() PackageSignatureKind
	Get_IsOptional() bool
	VerifyContentIntegrityAsync() *IAsyncOperation[bool]
}

type IPackage4Vtbl struct {
	win32.IInspectableVtbl
	Get_SignatureKind           uintptr
	Get_IsOptional              uintptr
	VerifyContentIntegrityAsync uintptr
}

type IPackage4 struct {
	win32.IInspectable
}

func (this *IPackage4) Vtbl() *IPackage4Vtbl {
	return (*IPackage4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackage4) Get_SignatureKind() PackageSignatureKind {
	var _result PackageSignatureKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SignatureKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackage4) Get_IsOptional() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsOptional, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackage4) VerifyContentIntegrityAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().VerifyContentIntegrityAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0E842DD4-D9AC-45ED-9A1E-74CE056B2635
var IID_IPackage5 = syscall.GUID{0x0E842DD4, 0xD9AC, 0x45ED,
	[8]byte{0x9A, 0x1E, 0x74, 0xCE, 0x05, 0x6B, 0x26, 0x35}}

type IPackage5Interface interface {
	win32.IInspectableInterface
	GetContentGroupsAsync() *IAsyncOperation[*IVector[*IPackageContentGroup]]
	GetContentGroupAsync(name string) *IAsyncOperation[*IPackageContentGroup]
	StageContentGroupsAsync(names *IIterable[string]) *IAsyncOperation[*IVector[*IPackageContentGroup]]
	StageContentGroupsWithPriorityAsync(names *IIterable[string], moveToHeadOfQueue bool) *IAsyncOperation[*IVector[*IPackageContentGroup]]
	SetInUseAsync(inUse bool) *IAsyncOperation[bool]
}

type IPackage5Vtbl struct {
	win32.IInspectableVtbl
	GetContentGroupsAsync               uintptr
	GetContentGroupAsync                uintptr
	StageContentGroupsAsync             uintptr
	StageContentGroupsWithPriorityAsync uintptr
	SetInUseAsync                       uintptr
}

type IPackage5 struct {
	win32.IInspectable
}

func (this *IPackage5) Vtbl() *IPackage5Vtbl {
	return (*IPackage5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackage5) GetContentGroupsAsync() *IAsyncOperation[*IVector[*IPackageContentGroup]] {
	var _result *IAsyncOperation[*IVector[*IPackageContentGroup]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetContentGroupsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackage5) GetContentGroupAsync(name string) *IAsyncOperation[*IPackageContentGroup] {
	var _result *IAsyncOperation[*IPackageContentGroup]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetContentGroupAsync, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackage5) StageContentGroupsAsync(names *IIterable[string]) *IAsyncOperation[*IVector[*IPackageContentGroup]] {
	var _result *IAsyncOperation[*IVector[*IPackageContentGroup]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StageContentGroupsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(names)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackage5) StageContentGroupsWithPriorityAsync(names *IIterable[string], moveToHeadOfQueue bool) *IAsyncOperation[*IVector[*IPackageContentGroup]] {
	var _result *IAsyncOperation[*IVector[*IPackageContentGroup]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StageContentGroupsWithPriorityAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(names)), uintptr(*(*byte)(unsafe.Pointer(&moveToHeadOfQueue))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackage5) SetInUseAsync(inUse bool) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetInUseAsync, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&inUse))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8B1AD942-12D7-4754-AE4E-638CBC0E3A2E
var IID_IPackage6 = syscall.GUID{0x8B1AD942, 0x12D7, 0x4754,
	[8]byte{0xAE, 0x4E, 0x63, 0x8C, 0xBC, 0x0E, 0x3A, 0x2E}}

type IPackage6Interface interface {
	win32.IInspectableInterface
	GetAppInstallerInfo() *IAppInstallerInfo
	CheckUpdateAvailabilityAsync() *IAsyncOperation[*IPackageUpdateAvailabilityResult]
}

type IPackage6Vtbl struct {
	win32.IInspectableVtbl
	GetAppInstallerInfo          uintptr
	CheckUpdateAvailabilityAsync uintptr
}

type IPackage6 struct {
	win32.IInspectable
}

func (this *IPackage6) Vtbl() *IPackage6Vtbl {
	return (*IPackage6Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackage6) GetAppInstallerInfo() *IAppInstallerInfo {
	var _result *IAppInstallerInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAppInstallerInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackage6) CheckUpdateAvailabilityAsync() *IAsyncOperation[*IPackageUpdateAvailabilityResult] {
	var _result *IAsyncOperation[*IPackageUpdateAvailabilityResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CheckUpdateAvailabilityAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 86FF8D31-A2E4-45E0-9732-283A6D88FDE1
var IID_IPackage7 = syscall.GUID{0x86FF8D31, 0xA2E4, 0x45E0,
	[8]byte{0x97, 0x32, 0x28, 0x3A, 0x6D, 0x88, 0xFD, 0xE1}}

type IPackage7Interface interface {
	win32.IInspectableInterface
	Get_MutableLocation() *IStorageFolder
	Get_EffectiveLocation() *IStorageFolder
}

type IPackage7Vtbl struct {
	win32.IInspectableVtbl
	Get_MutableLocation   uintptr
	Get_EffectiveLocation uintptr
}

type IPackage7 struct {
	win32.IInspectable
}

func (this *IPackage7) Vtbl() *IPackage7Vtbl {
	return (*IPackage7Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackage7) Get_MutableLocation() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MutableLocation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackage7) Get_EffectiveLocation() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EffectiveLocation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2C584F7B-CE2A-4BE6-A093-77CFBB2A7EA1
var IID_IPackage8 = syscall.GUID{0x2C584F7B, 0xCE2A, 0x4BE6,
	[8]byte{0xA0, 0x93, 0x77, 0xCF, 0xBB, 0x2A, 0x7E, 0xA1}}

type IPackage8Interface interface {
	win32.IInspectableInterface
	Get_EffectiveExternalLocation() *IStorageFolder
	Get_MachineExternalLocation() *IStorageFolder
	Get_UserExternalLocation() *IStorageFolder
	Get_InstalledPath() string
	Get_MutablePath() string
	Get_EffectivePath() string
	Get_EffectiveExternalPath() string
	Get_MachineExternalPath() string
	Get_UserExternalPath() string
	GetLogoAsRandomAccessStreamReference(size Size) *IRandomAccessStreamReference
	GetAppListEntries() *IVectorView[*IAppListEntry]
	Get_IsStub() bool
}

type IPackage8Vtbl struct {
	win32.IInspectableVtbl
	Get_EffectiveExternalLocation        uintptr
	Get_MachineExternalLocation          uintptr
	Get_UserExternalLocation             uintptr
	Get_InstalledPath                    uintptr
	Get_MutablePath                      uintptr
	Get_EffectivePath                    uintptr
	Get_EffectiveExternalPath            uintptr
	Get_MachineExternalPath              uintptr
	Get_UserExternalPath                 uintptr
	GetLogoAsRandomAccessStreamReference uintptr
	GetAppListEntries                    uintptr
	Get_IsStub                           uintptr
}

type IPackage8 struct {
	win32.IInspectable
}

func (this *IPackage8) Vtbl() *IPackage8Vtbl {
	return (*IPackage8Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackage8) Get_EffectiveExternalLocation() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EffectiveExternalLocation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackage8) Get_MachineExternalLocation() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MachineExternalLocation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackage8) Get_UserExternalLocation() *IStorageFolder {
	var _result *IStorageFolder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UserExternalLocation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackage8) Get_InstalledPath() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InstalledPath, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPackage8) Get_MutablePath() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MutablePath, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPackage8) Get_EffectivePath() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EffectivePath, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPackage8) Get_EffectiveExternalPath() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EffectiveExternalPath, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPackage8) Get_MachineExternalPath() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MachineExternalPath, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPackage8) Get_UserExternalPath() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UserExternalPath, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPackage8) GetLogoAsRandomAccessStreamReference(size Size) *IRandomAccessStreamReference {
	var _result *IRandomAccessStreamReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetLogoAsRandomAccessStreamReference, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&size)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackage8) GetAppListEntries() *IVectorView[*IAppListEntry] {
	var _result *IVectorView[*IAppListEntry]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAppListEntries, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackage8) Get_IsStub() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsStub, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D5AB224F-D7E1-49EC-90CE-720CDBD02E9C
var IID_IPackage9 = syscall.GUID{0xD5AB224F, 0xD7E1, 0x49EC,
	[8]byte{0x90, 0xCE, 0x72, 0x0C, 0xDB, 0xD0, 0x2E, 0x9C}}

type IPackage9Interface interface {
	win32.IInspectableInterface
	FindRelatedPackages(options *IFindRelatedPackagesOptions) *IVector[*IPackage]
	Get_SourceUriSchemeName() string
}

type IPackage9Vtbl struct {
	win32.IInspectableVtbl
	FindRelatedPackages     uintptr
	Get_SourceUriSchemeName uintptr
}

type IPackage9 struct {
	win32.IInspectable
}

func (this *IPackage9) Vtbl() *IPackage9Vtbl {
	return (*IPackage9Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackage9) FindRelatedPackages(options *IFindRelatedPackagesOptions) *IVector[*IPackage] {
	var _result *IVector[*IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindRelatedPackages, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackage9) Get_SourceUriSchemeName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SourceUriSchemeName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 230A3751-9DE3-4445-BE74-91FB325ABEFE
var IID_IPackageCatalog = syscall.GUID{0x230A3751, 0x9DE3, 0x4445,
	[8]byte{0xBE, 0x74, 0x91, 0xFB, 0x32, 0x5A, 0xBE, 0xFE}}

type IPackageCatalogInterface interface {
	win32.IInspectableInterface
	Add_PackageStaging(handler TypedEventHandler[*IPackageCatalog, *IPackageStagingEventArgs]) EventRegistrationToken
	Remove_PackageStaging(token EventRegistrationToken)
	Add_PackageInstalling(handler TypedEventHandler[*IPackageCatalog, *IPackageInstallingEventArgs]) EventRegistrationToken
	Remove_PackageInstalling(token EventRegistrationToken)
	Add_PackageUpdating(handler TypedEventHandler[*IPackageCatalog, *IPackageUpdatingEventArgs]) EventRegistrationToken
	Remove_PackageUpdating(token EventRegistrationToken)
	Add_PackageUninstalling(handler TypedEventHandler[*IPackageCatalog, *IPackageUninstallingEventArgs]) EventRegistrationToken
	Remove_PackageUninstalling(token EventRegistrationToken)
	Add_PackageStatusChanged(handler TypedEventHandler[*IPackageCatalog, *IPackageStatusChangedEventArgs]) EventRegistrationToken
	Remove_PackageStatusChanged(token EventRegistrationToken)
}

type IPackageCatalogVtbl struct {
	win32.IInspectableVtbl
	Add_PackageStaging          uintptr
	Remove_PackageStaging       uintptr
	Add_PackageInstalling       uintptr
	Remove_PackageInstalling    uintptr
	Add_PackageUpdating         uintptr
	Remove_PackageUpdating      uintptr
	Add_PackageUninstalling     uintptr
	Remove_PackageUninstalling  uintptr
	Add_PackageStatusChanged    uintptr
	Remove_PackageStatusChanged uintptr
}

type IPackageCatalog struct {
	win32.IInspectable
}

func (this *IPackageCatalog) Vtbl() *IPackageCatalogVtbl {
	return (*IPackageCatalogVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageCatalog) Add_PackageStaging(handler TypedEventHandler[*IPackageCatalog, *IPackageStagingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PackageStaging, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageCatalog) Remove_PackageStaging(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PackageStaging, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPackageCatalog) Add_PackageInstalling(handler TypedEventHandler[*IPackageCatalog, *IPackageInstallingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PackageInstalling, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageCatalog) Remove_PackageInstalling(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PackageInstalling, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPackageCatalog) Add_PackageUpdating(handler TypedEventHandler[*IPackageCatalog, *IPackageUpdatingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PackageUpdating, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageCatalog) Remove_PackageUpdating(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PackageUpdating, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPackageCatalog) Add_PackageUninstalling(handler TypedEventHandler[*IPackageCatalog, *IPackageUninstallingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PackageUninstalling, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageCatalog) Remove_PackageUninstalling(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PackageUninstalling, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPackageCatalog) Add_PackageStatusChanged(handler TypedEventHandler[*IPackageCatalog, *IPackageStatusChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PackageStatusChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageCatalog) Remove_PackageStatusChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PackageStatusChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 96A60C36-8FF7-4344-B6BF-EE64C2207ED2
var IID_IPackageCatalog2 = syscall.GUID{0x96A60C36, 0x8FF7, 0x4344,
	[8]byte{0xB6, 0xBF, 0xEE, 0x64, 0xC2, 0x20, 0x7E, 0xD2}}

type IPackageCatalog2Interface interface {
	win32.IInspectableInterface
	Add_PackageContentGroupStaging(handler TypedEventHandler[*IPackageCatalog, *IPackageContentGroupStagingEventArgs]) EventRegistrationToken
	Remove_PackageContentGroupStaging(token EventRegistrationToken)
	AddOptionalPackageAsync(optionalPackageFamilyName string) *IAsyncOperation[*IPackageCatalogAddOptionalPackageResult]
}

type IPackageCatalog2Vtbl struct {
	win32.IInspectableVtbl
	Add_PackageContentGroupStaging    uintptr
	Remove_PackageContentGroupStaging uintptr
	AddOptionalPackageAsync           uintptr
}

type IPackageCatalog2 struct {
	win32.IInspectable
}

func (this *IPackageCatalog2) Vtbl() *IPackageCatalog2Vtbl {
	return (*IPackageCatalog2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageCatalog2) Add_PackageContentGroupStaging(handler TypedEventHandler[*IPackageCatalog, *IPackageContentGroupStagingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PackageContentGroupStaging, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageCatalog2) Remove_PackageContentGroupStaging(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PackageContentGroupStaging, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPackageCatalog2) AddOptionalPackageAsync(optionalPackageFamilyName string) *IAsyncOperation[*IPackageCatalogAddOptionalPackageResult] {
	var _result *IAsyncOperation[*IPackageCatalogAddOptionalPackageResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddOptionalPackageAsync, uintptr(unsafe.Pointer(this)), NewHStr(optionalPackageFamilyName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 96DD5C88-8837-43F9-9015-033434BA14F3
var IID_IPackageCatalog3 = syscall.GUID{0x96DD5C88, 0x8837, 0x43F9,
	[8]byte{0x90, 0x15, 0x03, 0x34, 0x34, 0xBA, 0x14, 0xF3}}

type IPackageCatalog3Interface interface {
	win32.IInspectableInterface
	RemoveOptionalPackagesAsync(optionalPackageFamilyNames *IIterable[string]) *IAsyncOperation[*IPackageCatalogRemoveOptionalPackagesResult]
}

type IPackageCatalog3Vtbl struct {
	win32.IInspectableVtbl
	RemoveOptionalPackagesAsync uintptr
}

type IPackageCatalog3 struct {
	win32.IInspectable
}

func (this *IPackageCatalog3) Vtbl() *IPackageCatalog3Vtbl {
	return (*IPackageCatalog3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageCatalog3) RemoveOptionalPackagesAsync(optionalPackageFamilyNames *IIterable[string]) *IAsyncOperation[*IPackageCatalogRemoveOptionalPackagesResult] {
	var _result *IAsyncOperation[*IPackageCatalogRemoveOptionalPackagesResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveOptionalPackagesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(optionalPackageFamilyNames)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C37C399B-44CC-4B7B-8BAF-796C04EAD3B9
var IID_IPackageCatalog4 = syscall.GUID{0xC37C399B, 0x44CC, 0x4B7B,
	[8]byte{0x8B, 0xAF, 0x79, 0x6C, 0x04, 0xEA, 0xD3, 0xB9}}

type IPackageCatalog4Interface interface {
	win32.IInspectableInterface
	AddResourcePackageAsync(resourcePackageFamilyName string, resourceID string, options AddResourcePackageOptions) *IAsyncOperationWithProgress[*IPackageCatalogAddResourcePackageResult, PackageInstallProgress]
	RemoveResourcePackagesAsync(resourcePackages *IIterable[*IPackage]) *IAsyncOperation[*IPackageCatalogRemoveResourcePackagesResult]
}

type IPackageCatalog4Vtbl struct {
	win32.IInspectableVtbl
	AddResourcePackageAsync     uintptr
	RemoveResourcePackagesAsync uintptr
}

type IPackageCatalog4 struct {
	win32.IInspectable
}

func (this *IPackageCatalog4) Vtbl() *IPackageCatalog4Vtbl {
	return (*IPackageCatalog4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageCatalog4) AddResourcePackageAsync(resourcePackageFamilyName string, resourceID string, options AddResourcePackageOptions) *IAsyncOperationWithProgress[*IPackageCatalogAddResourcePackageResult, PackageInstallProgress] {
	var _result *IAsyncOperationWithProgress[*IPackageCatalogAddResourcePackageResult, PackageInstallProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddResourcePackageAsync, uintptr(unsafe.Pointer(this)), NewHStr(resourcePackageFamilyName).Ptr, NewHStr(resourceID).Ptr, uintptr(options), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageCatalog4) RemoveResourcePackagesAsync(resourcePackages *IIterable[*IPackage]) *IAsyncOperation[*IPackageCatalogRemoveResourcePackagesResult] {
	var _result *IAsyncOperation[*IPackageCatalogRemoveResourcePackagesResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveResourcePackagesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(resourcePackages)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3BF10CD4-B4DF-47B3-A963-E2FA832F7DD3
var IID_IPackageCatalogAddOptionalPackageResult = syscall.GUID{0x3BF10CD4, 0xB4DF, 0x47B3,
	[8]byte{0xA9, 0x63, 0xE2, 0xFA, 0x83, 0x2F, 0x7D, 0xD3}}

type IPackageCatalogAddOptionalPackageResultInterface interface {
	win32.IInspectableInterface
	Get_Package() *IPackage
	Get_ExtendedError() HResult
}

type IPackageCatalogAddOptionalPackageResultVtbl struct {
	win32.IInspectableVtbl
	Get_Package       uintptr
	Get_ExtendedError uintptr
}

type IPackageCatalogAddOptionalPackageResult struct {
	win32.IInspectable
}

func (this *IPackageCatalogAddOptionalPackageResult) Vtbl() *IPackageCatalogAddOptionalPackageResultVtbl {
	return (*IPackageCatalogAddOptionalPackageResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageCatalogAddOptionalPackageResult) Get_Package() *IPackage {
	var _result *IPackage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Package, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageCatalogAddOptionalPackageResult) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 9636CE0D-3E17-493F-AA08-CCEC6FDEF699
var IID_IPackageCatalogAddResourcePackageResult = syscall.GUID{0x9636CE0D, 0x3E17, 0x493F,
	[8]byte{0xAA, 0x08, 0xCC, 0xEC, 0x6F, 0xDE, 0xF6, 0x99}}

type IPackageCatalogAddResourcePackageResultInterface interface {
	win32.IInspectableInterface
	Get_Package() *IPackage
	Get_IsComplete() bool
	Get_ExtendedError() HResult
}

type IPackageCatalogAddResourcePackageResultVtbl struct {
	win32.IInspectableVtbl
	Get_Package       uintptr
	Get_IsComplete    uintptr
	Get_ExtendedError uintptr
}

type IPackageCatalogAddResourcePackageResult struct {
	win32.IInspectable
}

func (this *IPackageCatalogAddResourcePackageResult) Vtbl() *IPackageCatalogAddResourcePackageResultVtbl {
	return (*IPackageCatalogAddResourcePackageResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageCatalogAddResourcePackageResult) Get_Package() *IPackage {
	var _result *IPackage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Package, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageCatalogAddResourcePackageResult) Get_IsComplete() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsComplete, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageCatalogAddResourcePackageResult) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 29D2F97B-D974-4E64-9359-22CADFD79828
var IID_IPackageCatalogRemoveOptionalPackagesResult = syscall.GUID{0x29D2F97B, 0xD974, 0x4E64,
	[8]byte{0x93, 0x59, 0x22, 0xCA, 0xDF, 0xD7, 0x98, 0x28}}

type IPackageCatalogRemoveOptionalPackagesResultInterface interface {
	win32.IInspectableInterface
	Get_PackagesRemoved() *IVectorView[*IPackage]
	Get_ExtendedError() HResult
}

type IPackageCatalogRemoveOptionalPackagesResultVtbl struct {
	win32.IInspectableVtbl
	Get_PackagesRemoved uintptr
	Get_ExtendedError   uintptr
}

type IPackageCatalogRemoveOptionalPackagesResult struct {
	win32.IInspectable
}

func (this *IPackageCatalogRemoveOptionalPackagesResult) Vtbl() *IPackageCatalogRemoveOptionalPackagesResultVtbl {
	return (*IPackageCatalogRemoveOptionalPackagesResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageCatalogRemoveOptionalPackagesResult) Get_PackagesRemoved() *IVectorView[*IPackage] {
	var _result *IVectorView[*IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PackagesRemoved, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageCatalogRemoveOptionalPackagesResult) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// AE719709-1A52-4321-87B3-E5A1A17981A7
var IID_IPackageCatalogRemoveResourcePackagesResult = syscall.GUID{0xAE719709, 0x1A52, 0x4321,
	[8]byte{0x87, 0xB3, 0xE5, 0xA1, 0xA1, 0x79, 0x81, 0xA7}}

type IPackageCatalogRemoveResourcePackagesResultInterface interface {
	win32.IInspectableInterface
	Get_PackagesRemoved() *IVectorView[*IPackage]
	Get_ExtendedError() HResult
}

type IPackageCatalogRemoveResourcePackagesResultVtbl struct {
	win32.IInspectableVtbl
	Get_PackagesRemoved uintptr
	Get_ExtendedError   uintptr
}

type IPackageCatalogRemoveResourcePackagesResult struct {
	win32.IInspectable
}

func (this *IPackageCatalogRemoveResourcePackagesResult) Vtbl() *IPackageCatalogRemoveResourcePackagesResultVtbl {
	return (*IPackageCatalogRemoveResourcePackagesResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageCatalogRemoveResourcePackagesResult) Get_PackagesRemoved() *IVectorView[*IPackage] {
	var _result *IVectorView[*IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PackagesRemoved, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageCatalogRemoveResourcePackagesResult) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A18C9696-E65B-4634-BA21-5E63EB7244A7
var IID_IPackageCatalogStatics = syscall.GUID{0xA18C9696, 0xE65B, 0x4634,
	[8]byte{0xBA, 0x21, 0x5E, 0x63, 0xEB, 0x72, 0x44, 0xA7}}

type IPackageCatalogStaticsInterface interface {
	win32.IInspectableInterface
	OpenForCurrentPackage() *IPackageCatalog
	OpenForCurrentUser() *IPackageCatalog
}

type IPackageCatalogStaticsVtbl struct {
	win32.IInspectableVtbl
	OpenForCurrentPackage uintptr
	OpenForCurrentUser    uintptr
}

type IPackageCatalogStatics struct {
	win32.IInspectable
}

func (this *IPackageCatalogStatics) Vtbl() *IPackageCatalogStaticsVtbl {
	return (*IPackageCatalogStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageCatalogStatics) OpenForCurrentPackage() *IPackageCatalog {
	var _result *IPackageCatalog
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenForCurrentPackage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageCatalogStatics) OpenForCurrentUser() *IPackageCatalog {
	var _result *IPackageCatalog
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenForCurrentUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4C11C159-9A28-598C-B185-55E1899B2BE4
var IID_IPackageCatalogStatics2 = syscall.GUID{0x4C11C159, 0x9A28, 0x598C,
	[8]byte{0xB1, 0x85, 0x55, 0xE1, 0x89, 0x9B, 0x2B, 0xE4}}

type IPackageCatalogStatics2Interface interface {
	win32.IInspectableInterface
	OpenForPackage(package_ *IPackage) *IPackageCatalog
}

type IPackageCatalogStatics2Vtbl struct {
	win32.IInspectableVtbl
	OpenForPackage uintptr
}

type IPackageCatalogStatics2 struct {
	win32.IInspectable
}

func (this *IPackageCatalogStatics2) Vtbl() *IPackageCatalogStatics2Vtbl {
	return (*IPackageCatalogStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageCatalogStatics2) OpenForPackage(package_ *IPackage) *IPackageCatalog {
	var _result *IPackageCatalog
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenForPackage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(package_)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8F62695D-120A-4798-B5E1-5800DDA8F2E1
var IID_IPackageContentGroup = syscall.GUID{0x8F62695D, 0x120A, 0x4798,
	[8]byte{0xB5, 0xE1, 0x58, 0x00, 0xDD, 0xA8, 0xF2, 0xE1}}

type IPackageContentGroupInterface interface {
	win32.IInspectableInterface
	Get_Package() *IPackage
	Get_Name() string
	Get_State() PackageContentGroupState
	Get_IsRequired() bool
}

type IPackageContentGroupVtbl struct {
	win32.IInspectableVtbl
	Get_Package    uintptr
	Get_Name       uintptr
	Get_State      uintptr
	Get_IsRequired uintptr
}

type IPackageContentGroup struct {
	win32.IInspectable
}

func (this *IPackageContentGroup) Vtbl() *IPackageContentGroupVtbl {
	return (*IPackageContentGroupVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageContentGroup) Get_Package() *IPackage {
	var _result *IPackage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Package, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageContentGroup) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPackageContentGroup) Get_State() PackageContentGroupState {
	var _result PackageContentGroupState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageContentGroup) Get_IsRequired() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsRequired, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 3D7BC27E-6F27-446C-986E-D4733D4D9113
var IID_IPackageContentGroupStagingEventArgs = syscall.GUID{0x3D7BC27E, 0x6F27, 0x446C,
	[8]byte{0x98, 0x6E, 0xD4, 0x73, 0x3D, 0x4D, 0x91, 0x13}}

type IPackageContentGroupStagingEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ActivityId() syscall.GUID
	Get_Package() *IPackage
	Get_Progress() float64
	Get_IsComplete() bool
	Get_ErrorCode() HResult
	Get_ContentGroupName() string
	Get_IsContentGroupRequired() bool
}

type IPackageContentGroupStagingEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ActivityId             uintptr
	Get_Package                uintptr
	Get_Progress               uintptr
	Get_IsComplete             uintptr
	Get_ErrorCode              uintptr
	Get_ContentGroupName       uintptr
	Get_IsContentGroupRequired uintptr
}

type IPackageContentGroupStagingEventArgs struct {
	win32.IInspectable
}

func (this *IPackageContentGroupStagingEventArgs) Vtbl() *IPackageContentGroupStagingEventArgsVtbl {
	return (*IPackageContentGroupStagingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageContentGroupStagingEventArgs) Get_ActivityId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ActivityId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageContentGroupStagingEventArgs) Get_Package() *IPackage {
	var _result *IPackage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Package, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageContentGroupStagingEventArgs) Get_Progress() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Progress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageContentGroupStagingEventArgs) Get_IsComplete() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsComplete, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageContentGroupStagingEventArgs) Get_ErrorCode() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ErrorCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageContentGroupStagingEventArgs) Get_ContentGroupName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentGroupName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPackageContentGroupStagingEventArgs) Get_IsContentGroupRequired() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsContentGroupRequired, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 70EE7619-5F12-4B92-B9EA-6CCADA13BC75
var IID_IPackageContentGroupStatics = syscall.GUID{0x70EE7619, 0x5F12, 0x4B92,
	[8]byte{0xB9, 0xEA, 0x6C, 0xCA, 0xDA, 0x13, 0xBC, 0x75}}

type IPackageContentGroupStaticsInterface interface {
	win32.IInspectableInterface
	Get_RequiredGroupName() string
}

type IPackageContentGroupStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_RequiredGroupName uintptr
}

type IPackageContentGroupStatics struct {
	win32.IInspectable
}

func (this *IPackageContentGroupStatics) Vtbl() *IPackageContentGroupStaticsVtbl {
	return (*IPackageContentGroupStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageContentGroupStatics) Get_RequiredGroupName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequiredGroupName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 1ADB665E-37C7-4790-9980-DD7AE74E8BB2
var IID_IPackageId = syscall.GUID{0x1ADB665E, 0x37C7, 0x4790,
	[8]byte{0x99, 0x80, 0xDD, 0x7A, 0xE7, 0x4E, 0x8B, 0xB2}}

type IPackageIdInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Get_Version() PackageVersion
	Get_Architecture() ProcessorArchitecture
	Get_ResourceId() string
	Get_Publisher() string
	Get_PublisherId() string
	Get_FullName() string
	Get_FamilyName() string
}

type IPackageIdVtbl struct {
	win32.IInspectableVtbl
	Get_Name         uintptr
	Get_Version      uintptr
	Get_Architecture uintptr
	Get_ResourceId   uintptr
	Get_Publisher    uintptr
	Get_PublisherId  uintptr
	Get_FullName     uintptr
	Get_FamilyName   uintptr
}

type IPackageId struct {
	win32.IInspectable
}

func (this *IPackageId) Vtbl() *IPackageIdVtbl {
	return (*IPackageIdVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageId) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPackageId) Get_Version() PackageVersion {
	var _result PackageVersion
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Version, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageId) Get_Architecture() ProcessorArchitecture {
	var _result ProcessorArchitecture
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Architecture, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageId) Get_ResourceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResourceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPackageId) Get_Publisher() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Publisher, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPackageId) Get_PublisherId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PublisherId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPackageId) Get_FullName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FullName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPackageId) Get_FamilyName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FamilyName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 40577A7C-0C9E-443D-9074-855F5CE0A08D
var IID_IPackageIdWithMetadata = syscall.GUID{0x40577A7C, 0x0C9E, 0x443D,
	[8]byte{0x90, 0x74, 0x85, 0x5F, 0x5C, 0xE0, 0xA0, 0x8D}}

type IPackageIdWithMetadataInterface interface {
	win32.IInspectableInterface
	Get_ProductId() string
	Get_Author() string
}

type IPackageIdWithMetadataVtbl struct {
	win32.IInspectableVtbl
	Get_ProductId uintptr
	Get_Author    uintptr
}

type IPackageIdWithMetadata struct {
	win32.IInspectable
}

func (this *IPackageIdWithMetadata) Vtbl() *IPackageIdWithMetadataVtbl {
	return (*IPackageIdWithMetadataVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageIdWithMetadata) Get_ProductId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProductId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPackageIdWithMetadata) Get_Author() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Author, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 97741EB7-AB7A-401A-8B61-EB0E7FAFF237
var IID_IPackageInstallingEventArgs = syscall.GUID{0x97741EB7, 0xAB7A, 0x401A,
	[8]byte{0x8B, 0x61, 0xEB, 0x0E, 0x7F, 0xAF, 0xF2, 0x37}}

type IPackageInstallingEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ActivityId() syscall.GUID
	Get_Package() *IPackage
	Get_Progress() float64
	Get_IsComplete() bool
	Get_ErrorCode() HResult
}

type IPackageInstallingEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ActivityId uintptr
	Get_Package    uintptr
	Get_Progress   uintptr
	Get_IsComplete uintptr
	Get_ErrorCode  uintptr
}

type IPackageInstallingEventArgs struct {
	win32.IInspectable
}

func (this *IPackageInstallingEventArgs) Vtbl() *IPackageInstallingEventArgsVtbl {
	return (*IPackageInstallingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageInstallingEventArgs) Get_ActivityId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ActivityId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageInstallingEventArgs) Get_Package() *IPackage {
	var _result *IPackage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Package, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageInstallingEventArgs) Get_Progress() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Progress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageInstallingEventArgs) Get_IsComplete() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsComplete, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageInstallingEventArgs) Get_ErrorCode() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ErrorCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 1041682D-54E2-4F51-B828-9EF7046C210F
var IID_IPackageStagingEventArgs = syscall.GUID{0x1041682D, 0x54E2, 0x4F51,
	[8]byte{0xB8, 0x28, 0x9E, 0xF7, 0x04, 0x6C, 0x21, 0x0F}}

type IPackageStagingEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ActivityId() syscall.GUID
	Get_Package() *IPackage
	Get_Progress() float64
	Get_IsComplete() bool
	Get_ErrorCode() HResult
}

type IPackageStagingEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ActivityId uintptr
	Get_Package    uintptr
	Get_Progress   uintptr
	Get_IsComplete uintptr
	Get_ErrorCode  uintptr
}

type IPackageStagingEventArgs struct {
	win32.IInspectable
}

func (this *IPackageStagingEventArgs) Vtbl() *IPackageStagingEventArgsVtbl {
	return (*IPackageStagingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageStagingEventArgs) Get_ActivityId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ActivityId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageStagingEventArgs) Get_Package() *IPackage {
	var _result *IPackage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Package, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageStagingEventArgs) Get_Progress() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Progress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageStagingEventArgs) Get_IsComplete() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsComplete, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageStagingEventArgs) Get_ErrorCode() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ErrorCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4E534BDF-2960-4878-97A4-9624DEB72F2D
var IID_IPackageStatics = syscall.GUID{0x4E534BDF, 0x2960, 0x4878,
	[8]byte{0x97, 0xA4, 0x96, 0x24, 0xDE, 0xB7, 0x2F, 0x2D}}

type IPackageStaticsInterface interface {
	win32.IInspectableInterface
	Get_Current() *IPackage
}

type IPackageStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Current uintptr
}

type IPackageStatics struct {
	win32.IInspectable
}

func (this *IPackageStatics) Vtbl() *IPackageStaticsVtbl {
	return (*IPackageStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageStatics) Get_Current() *IPackage {
	var _result *IPackage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Current, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5FE74F71-A365-4C09-A02D-046D525EA1DA
var IID_IPackageStatus = syscall.GUID{0x5FE74F71, 0xA365, 0x4C09,
	[8]byte{0xA0, 0x2D, 0x04, 0x6D, 0x52, 0x5E, 0xA1, 0xDA}}

type IPackageStatusInterface interface {
	win32.IInspectableInterface
	VerifyIsOK() bool
	Get_NotAvailable() bool
	Get_PackageOffline() bool
	Get_DataOffline() bool
	Get_Disabled() bool
	Get_NeedsRemediation() bool
	Get_LicenseIssue() bool
	Get_Modified() bool
	Get_Tampered() bool
	Get_DependencyIssue() bool
	Get_Servicing() bool
	Get_DeploymentInProgress() bool
}

type IPackageStatusVtbl struct {
	win32.IInspectableVtbl
	VerifyIsOK               uintptr
	Get_NotAvailable         uintptr
	Get_PackageOffline       uintptr
	Get_DataOffline          uintptr
	Get_Disabled             uintptr
	Get_NeedsRemediation     uintptr
	Get_LicenseIssue         uintptr
	Get_Modified             uintptr
	Get_Tampered             uintptr
	Get_DependencyIssue      uintptr
	Get_Servicing            uintptr
	Get_DeploymentInProgress uintptr
}

type IPackageStatus struct {
	win32.IInspectable
}

func (this *IPackageStatus) Vtbl() *IPackageStatusVtbl {
	return (*IPackageStatusVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageStatus) VerifyIsOK() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().VerifyIsOK, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageStatus) Get_NotAvailable() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NotAvailable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageStatus) Get_PackageOffline() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PackageOffline, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageStatus) Get_DataOffline() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DataOffline, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageStatus) Get_Disabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Disabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageStatus) Get_NeedsRemediation() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NeedsRemediation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageStatus) Get_LicenseIssue() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LicenseIssue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageStatus) Get_Modified() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Modified, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageStatus) Get_Tampered() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Tampered, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageStatus) Get_DependencyIssue() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DependencyIssue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageStatus) Get_Servicing() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Servicing, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageStatus) Get_DeploymentInProgress() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeploymentInProgress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F428FA93-7C56-4862-ACFA-ABAEDCC0694D
var IID_IPackageStatus2 = syscall.GUID{0xF428FA93, 0x7C56, 0x4862,
	[8]byte{0xAC, 0xFA, 0xAB, 0xAE, 0xDC, 0xC0, 0x69, 0x4D}}

type IPackageStatus2Interface interface {
	win32.IInspectableInterface
	Get_IsPartiallyStaged() bool
}

type IPackageStatus2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsPartiallyStaged uintptr
}

type IPackageStatus2 struct {
	win32.IInspectable
}

func (this *IPackageStatus2) Vtbl() *IPackageStatus2Vtbl {
	return (*IPackageStatus2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageStatus2) Get_IsPartiallyStaged() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPartiallyStaged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 437D714D-BD80-4A70-BC50-F6E796509575
var IID_IPackageStatusChangedEventArgs = syscall.GUID{0x437D714D, 0xBD80, 0x4A70,
	[8]byte{0xBC, 0x50, 0xF6, 0xE7, 0x96, 0x50, 0x95, 0x75}}

type IPackageStatusChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Package() *IPackage
}

type IPackageStatusChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Package uintptr
}

type IPackageStatusChangedEventArgs struct {
	win32.IInspectable
}

func (this *IPackageStatusChangedEventArgs) Vtbl() *IPackageStatusChangedEventArgsVtbl {
	return (*IPackageStatusChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageStatusChangedEventArgs) Get_Package() *IPackage {
	var _result *IPackage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Package, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4443AA52-AB22-44CD-82BB-4EC9B827367A
var IID_IPackageUninstallingEventArgs = syscall.GUID{0x4443AA52, 0xAB22, 0x44CD,
	[8]byte{0x82, 0xBB, 0x4E, 0xC9, 0xB8, 0x27, 0x36, 0x7A}}

type IPackageUninstallingEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ActivityId() syscall.GUID
	Get_Package() *IPackage
	Get_Progress() float64
	Get_IsComplete() bool
	Get_ErrorCode() HResult
}

type IPackageUninstallingEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ActivityId uintptr
	Get_Package    uintptr
	Get_Progress   uintptr
	Get_IsComplete uintptr
	Get_ErrorCode  uintptr
}

type IPackageUninstallingEventArgs struct {
	win32.IInspectable
}

func (this *IPackageUninstallingEventArgs) Vtbl() *IPackageUninstallingEventArgsVtbl {
	return (*IPackageUninstallingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageUninstallingEventArgs) Get_ActivityId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ActivityId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageUninstallingEventArgs) Get_Package() *IPackage {
	var _result *IPackage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Package, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageUninstallingEventArgs) Get_Progress() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Progress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageUninstallingEventArgs) Get_IsComplete() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsComplete, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageUninstallingEventArgs) Get_ErrorCode() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ErrorCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 114E5009-199A-48A1-A079-313C45634A71
var IID_IPackageUpdateAvailabilityResult = syscall.GUID{0x114E5009, 0x199A, 0x48A1,
	[8]byte{0xA0, 0x79, 0x31, 0x3C, 0x45, 0x63, 0x4A, 0x71}}

type IPackageUpdateAvailabilityResultInterface interface {
	win32.IInspectableInterface
	Get_Availability() PackageUpdateAvailability
	Get_ExtendedError() HResult
}

type IPackageUpdateAvailabilityResultVtbl struct {
	win32.IInspectableVtbl
	Get_Availability  uintptr
	Get_ExtendedError uintptr
}

type IPackageUpdateAvailabilityResult struct {
	win32.IInspectable
}

func (this *IPackageUpdateAvailabilityResult) Vtbl() *IPackageUpdateAvailabilityResultVtbl {
	return (*IPackageUpdateAvailabilityResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageUpdateAvailabilityResult) Get_Availability() PackageUpdateAvailability {
	var _result PackageUpdateAvailability
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Availability, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageUpdateAvailabilityResult) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// CD7B4228-FD74-443E-B114-23E677B0E86F
var IID_IPackageUpdatingEventArgs = syscall.GUID{0xCD7B4228, 0xFD74, 0x443E,
	[8]byte{0xB1, 0x14, 0x23, 0xE6, 0x77, 0xB0, 0xE8, 0x6F}}

type IPackageUpdatingEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ActivityId() syscall.GUID
	Get_SourcePackage() *IPackage
	Get_TargetPackage() *IPackage
	Get_Progress() float64
	Get_IsComplete() bool
	Get_ErrorCode() HResult
}

type IPackageUpdatingEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ActivityId    uintptr
	Get_SourcePackage uintptr
	Get_TargetPackage uintptr
	Get_Progress      uintptr
	Get_IsComplete    uintptr
	Get_ErrorCode     uintptr
}

type IPackageUpdatingEventArgs struct {
	win32.IInspectable
}

func (this *IPackageUpdatingEventArgs) Vtbl() *IPackageUpdatingEventArgsVtbl {
	return (*IPackageUpdatingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageUpdatingEventArgs) Get_ActivityId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ActivityId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageUpdatingEventArgs) Get_SourcePackage() *IPackage {
	var _result *IPackage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SourcePackage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageUpdatingEventArgs) Get_TargetPackage() *IPackage {
	var _result *IPackage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TargetPackage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageUpdatingEventArgs) Get_Progress() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Progress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageUpdatingEventArgs) Get_IsComplete() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsComplete, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageUpdatingEventArgs) Get_ErrorCode() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ErrorCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 95949780-1DE9-40F2-B452-0DE9F1910012
var IID_IPackageWithMetadata = syscall.GUID{0x95949780, 0x1DE9, 0x40F2,
	[8]byte{0xB4, 0x52, 0x0D, 0xE9, 0xF1, 0x91, 0x00, 0x12}}

type IPackageWithMetadataInterface interface {
	win32.IInspectableInterface
	Get_InstallDate() DateTime
	GetThumbnailToken() string
	Launch(parameters string)
}

type IPackageWithMetadataVtbl struct {
	win32.IInspectableVtbl
	Get_InstallDate   uintptr
	GetThumbnailToken uintptr
	Launch            uintptr
}

type IPackageWithMetadata struct {
	win32.IInspectable
}

func (this *IPackageWithMetadata) Vtbl() *IPackageWithMetadataVtbl {
	return (*IPackageWithMetadataVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageWithMetadata) Get_InstallDate() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InstallDate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageWithMetadata) GetThumbnailToken() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetThumbnailToken, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPackageWithMetadata) Launch(parameters string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Launch, uintptr(unsafe.Pointer(this)), NewHStr(parameters).Ptr)
	_ = _hr
}

// F75C23C8-B5F2-4F6C-88DD-36CB1D599D17
var IID_IStartupTask = syscall.GUID{0xF75C23C8, 0xB5F2, 0x4F6C,
	[8]byte{0x88, 0xDD, 0x36, 0xCB, 0x1D, 0x59, 0x9D, 0x17}}

type IStartupTaskInterface interface {
	win32.IInspectableInterface
	RequestEnableAsync() *IAsyncOperation[StartupTaskState]
	Disable()
	Get_State() StartupTaskState
	Get_TaskId() string
}

type IStartupTaskVtbl struct {
	win32.IInspectableVtbl
	RequestEnableAsync uintptr
	Disable            uintptr
	Get_State          uintptr
	Get_TaskId         uintptr
}

type IStartupTask struct {
	win32.IInspectable
}

func (this *IStartupTask) Vtbl() *IStartupTaskVtbl {
	return (*IStartupTaskVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStartupTask) RequestEnableAsync() *IAsyncOperation[StartupTaskState] {
	var _result *IAsyncOperation[StartupTaskState]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestEnableAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStartupTask) Disable() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Disable, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IStartupTask) Get_State() StartupTaskState {
	var _result StartupTaskState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStartupTask) Get_TaskId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TaskId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// EE5B60BD-A148-41A7-B26E-E8B88A1E62F8
var IID_IStartupTaskStatics = syscall.GUID{0xEE5B60BD, 0xA148, 0x41A7,
	[8]byte{0xB2, 0x6E, 0xE8, 0xB8, 0x8A, 0x1E, 0x62, 0xF8}}

type IStartupTaskStaticsInterface interface {
	win32.IInspectableInterface
	GetForCurrentPackageAsync() *IAsyncOperation[*IVectorView[*IStartupTask]]
	GetAsync(taskId string) *IAsyncOperation[*IStartupTask]
}

type IStartupTaskStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForCurrentPackageAsync uintptr
	GetAsync                  uintptr
}

type IStartupTaskStatics struct {
	win32.IInspectable
}

func (this *IStartupTaskStatics) Vtbl() *IStartupTaskStaticsVtbl {
	return (*IStartupTaskStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStartupTaskStatics) GetForCurrentPackageAsync() *IAsyncOperation[*IVectorView[*IStartupTask]] {
	var _result *IAsyncOperation[*IVectorView[*IStartupTask]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentPackageAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStartupTaskStatics) GetAsync(taskId string) *IAsyncOperation[*IStartupTask] {
	var _result *IAsyncOperation[*IStartupTask]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAsync, uintptr(unsafe.Pointer(this)), NewHStr(taskId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 59140509-8BC9-4EB4-B636-DABDC4F46F66
var IID_ISuspendingDeferral = syscall.GUID{0x59140509, 0x8BC9, 0x4EB4,
	[8]byte{0xB6, 0x36, 0xDA, 0xBD, 0xC4, 0xF4, 0x6F, 0x66}}

type ISuspendingDeferralInterface interface {
	win32.IInspectableInterface
	Complete()
}

type ISuspendingDeferralVtbl struct {
	win32.IInspectableVtbl
	Complete uintptr
}

type ISuspendingDeferral struct {
	win32.IInspectable
}

func (this *ISuspendingDeferral) Vtbl() *ISuspendingDeferralVtbl {
	return (*ISuspendingDeferralVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISuspendingDeferral) Complete() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Complete, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 96061C05-2DBA-4D08-B0BD-2B30A131C6AA
var IID_ISuspendingEventArgs = syscall.GUID{0x96061C05, 0x2DBA, 0x4D08,
	[8]byte{0xB0, 0xBD, 0x2B, 0x30, 0xA1, 0x31, 0xC6, 0xAA}}

type ISuspendingEventArgsInterface interface {
	win32.IInspectableInterface
	Get_SuspendingOperation() *ISuspendingOperation
}

type ISuspendingEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_SuspendingOperation uintptr
}

type ISuspendingEventArgs struct {
	win32.IInspectable
}

func (this *ISuspendingEventArgs) Vtbl() *ISuspendingEventArgsVtbl {
	return (*ISuspendingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISuspendingEventArgs) Get_SuspendingOperation() *ISuspendingOperation {
	var _result *ISuspendingOperation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SuspendingOperation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9DA4CA41-20E1-4E9B-9F65-A9F435340C3A
var IID_ISuspendingOperation = syscall.GUID{0x9DA4CA41, 0x20E1, 0x4E9B,
	[8]byte{0x9F, 0x65, 0xA9, 0xF4, 0x35, 0x34, 0x0C, 0x3A}}

type ISuspendingOperationInterface interface {
	win32.IInspectableInterface
	GetDeferral() *ISuspendingDeferral
	Get_Deadline() DateTime
}

type ISuspendingOperationVtbl struct {
	win32.IInspectableVtbl
	GetDeferral  uintptr
	Get_Deadline uintptr
}

type ISuspendingOperation struct {
	win32.IInspectable
}

func (this *ISuspendingOperation) Vtbl() *ISuspendingOperationVtbl {
	return (*ISuspendingOperationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISuspendingOperation) GetDeferral() *ISuspendingDeferral {
	var _result *ISuspendingDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISuspendingOperation) Get_Deadline() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Deadline, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type AppDisplayInfo struct {
	RtClass
	*IAppDisplayInfo
}

type AppInfo struct {
	RtClass
	*IAppInfo
}

func NewIAppInfoStatics() *IAppInfoStatics {
	var p *IAppInfoStatics
	hs := NewHStr("Windows.ApplicationModel.AppInfo")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAppInfoStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type AppInstallerInfo struct {
	RtClass
	*IAppInstallerInfo
}

type Package struct {
	RtClass
	*IPackage
}

func NewIPackageStatics() *IPackageStatics {
	var p *IPackageStatics
	hs := NewHStr("Windows.ApplicationModel.Package")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPackageStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type PackageCatalogRemoveResourcePackagesResult struct {
	RtClass
	*IPackageCatalogRemoveResourcePackagesResult
}

type PackageContentGroupStagingEventArgs struct {
	RtClass
	*IPackageContentGroupStagingEventArgs
}

type PackageId struct {
	RtClass
	*IPackageId
}

type PackageInstallingEventArgs struct {
	RtClass
	*IPackageInstallingEventArgs
}

type PackageStagingEventArgs struct {
	RtClass
	*IPackageStagingEventArgs
}

type PackageStatus struct {
	RtClass
	*IPackageStatus
}

type PackageStatusChangedEventArgs struct {
	RtClass
	*IPackageStatusChangedEventArgs
}

type PackageUninstallingEventArgs struct {
	RtClass
	*IPackageUninstallingEventArgs
}

type PackageUpdatingEventArgs struct {
	RtClass
	*IPackageUpdatingEventArgs
}
