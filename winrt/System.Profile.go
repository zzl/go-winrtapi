package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type PlatformDataCollectionLevel int32

const (
	PlatformDataCollectionLevel_Security PlatformDataCollectionLevel = 0
	PlatformDataCollectionLevel_Basic    PlatformDataCollectionLevel = 1
	PlatformDataCollectionLevel_Enhanced PlatformDataCollectionLevel = 2
	PlatformDataCollectionLevel_Full     PlatformDataCollectionLevel = 3
)

// enum
type SystemIdentificationSource int32

const (
	SystemIdentificationSource_None     SystemIdentificationSource = 0
	SystemIdentificationSource_Tpm      SystemIdentificationSource = 1
	SystemIdentificationSource_Uefi     SystemIdentificationSource = 2
	SystemIdentificationSource_Registry SystemIdentificationSource = 3
)

// enum
type SystemOutOfBoxExperienceState int32

const (
	SystemOutOfBoxExperienceState_NotStarted SystemOutOfBoxExperienceState = 0
	SystemOutOfBoxExperienceState_InProgress SystemOutOfBoxExperienceState = 1
	SystemOutOfBoxExperienceState_Completed  SystemOutOfBoxExperienceState = 2
)

// enum
// flags
type UnsupportedAppRequirementReasons uint32

const (
	UnsupportedAppRequirementReasons_Unknown        UnsupportedAppRequirementReasons = 0
	UnsupportedAppRequirementReasons_DeniedBySystem UnsupportedAppRequirementReasons = 1
)

// structs

type ProfileHardwareTokenContract struct {
}

type ProfileRetailInfoContract struct {
}

type ProfileSharedModeContract struct {
}

// interfaces

// 1D5EE066-188D-5BA9-4387-ACAEB0E7E305
var IID_IAnalyticsInfoStatics = syscall.GUID{0x1D5EE066, 0x188D, 0x5BA9,
	[8]byte{0x43, 0x87, 0xAC, 0xAE, 0xB0, 0xE7, 0xE3, 0x05}}

type IAnalyticsInfoStaticsInterface interface {
	win32.IInspectableInterface
	Get_VersionInfo() *IAnalyticsVersionInfo
	Get_DeviceForm() string
}

type IAnalyticsInfoStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_VersionInfo uintptr
	Get_DeviceForm  uintptr
}

type IAnalyticsInfoStatics struct {
	win32.IInspectable
}

func (this *IAnalyticsInfoStatics) Vtbl() *IAnalyticsInfoStaticsVtbl {
	return (*IAnalyticsInfoStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAnalyticsInfoStatics) Get_VersionInfo() *IAnalyticsVersionInfo {
	var _result *IAnalyticsVersionInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VersionInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAnalyticsInfoStatics) Get_DeviceForm() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceForm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 101704EA-A7F9-46D2-AB94-016865AFDB25
var IID_IAnalyticsInfoStatics2 = syscall.GUID{0x101704EA, 0xA7F9, 0x46D2,
	[8]byte{0xAB, 0x94, 0x01, 0x68, 0x65, 0xAF, 0xDB, 0x25}}

type IAnalyticsInfoStatics2Interface interface {
	win32.IInspectableInterface
	GetSystemPropertiesAsync(attributeNames *IIterable[string]) *IAsyncOperation[*IMapView[string, string]]
}

type IAnalyticsInfoStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetSystemPropertiesAsync uintptr
}

type IAnalyticsInfoStatics2 struct {
	win32.IInspectable
}

func (this *IAnalyticsInfoStatics2) Vtbl() *IAnalyticsInfoStatics2Vtbl {
	return (*IAnalyticsInfoStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAnalyticsInfoStatics2) GetSystemPropertiesAsync(attributeNames *IIterable[string]) *IAsyncOperation[*IMapView[string, string]] {
	var _result *IAsyncOperation[*IMapView[string, string]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSystemPropertiesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(attributeNames)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 926130B8-9955-4C74-BDC1-7CD0DECF9B03
var IID_IAnalyticsVersionInfo = syscall.GUID{0x926130B8, 0x9955, 0x4C74,
	[8]byte{0xBD, 0xC1, 0x7C, 0xD0, 0xDE, 0xCF, 0x9B, 0x03}}

type IAnalyticsVersionInfoInterface interface {
	win32.IInspectableInterface
	Get_DeviceFamily() string
	Get_DeviceFamilyVersion() string
}

type IAnalyticsVersionInfoVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceFamily        uintptr
	Get_DeviceFamilyVersion uintptr
}

type IAnalyticsVersionInfo struct {
	win32.IInspectable
}

func (this *IAnalyticsVersionInfo) Vtbl() *IAnalyticsVersionInfoVtbl {
	return (*IAnalyticsVersionInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAnalyticsVersionInfo) Get_DeviceFamily() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceFamily, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAnalyticsVersionInfo) Get_DeviceFamilyVersion() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceFamilyVersion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 76E915B1-FF36-407C-9F57-160D3E540747
var IID_IAnalyticsVersionInfo2 = syscall.GUID{0x76E915B1, 0xFF36, 0x407C,
	[8]byte{0x9F, 0x57, 0x16, 0x0D, 0x3E, 0x54, 0x07, 0x47}}

type IAnalyticsVersionInfo2Interface interface {
	win32.IInspectableInterface
	Get_ProductName() string
}

type IAnalyticsVersionInfo2Vtbl struct {
	win32.IInspectableVtbl
	Get_ProductName uintptr
}

type IAnalyticsVersionInfo2 struct {
	win32.IInspectable
}

func (this *IAnalyticsVersionInfo2) Vtbl() *IAnalyticsVersionInfo2Vtbl {
	return (*IAnalyticsVersionInfo2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAnalyticsVersionInfo2) Get_ProductName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProductName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 1664A082-0F38-5C99-83E4-48995970861C
var IID_IAppApplicabilityStatics = syscall.GUID{0x1664A082, 0x0F38, 0x5C99,
	[8]byte{0x83, 0xE4, 0x48, 0x99, 0x59, 0x70, 0x86, 0x1C}}

type IAppApplicabilityStaticsInterface interface {
	win32.IInspectableInterface
	GetUnsupportedAppRequirements(capabilities *IIterable[string]) *IVectorView[*IUnsupportedAppRequirement]
}

type IAppApplicabilityStaticsVtbl struct {
	win32.IInspectableVtbl
	GetUnsupportedAppRequirements uintptr
}

type IAppApplicabilityStatics struct {
	win32.IInspectable
}

func (this *IAppApplicabilityStatics) Vtbl() *IAppApplicabilityStaticsVtbl {
	return (*IAppApplicabilityStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppApplicabilityStatics) GetUnsupportedAppRequirements(capabilities *IIterable[string]) *IVectorView[*IUnsupportedAppRequirement] {
	var _result *IVectorView[*IUnsupportedAppRequirement]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetUnsupportedAppRequirements, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(capabilities)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FC53F0EF-4D3E-4E13-9B23-505F4D091E92
var IID_IEducationSettingsStatics = syscall.GUID{0xFC53F0EF, 0x4D3E, 0x4E13,
	[8]byte{0x9B, 0x23, 0x50, 0x5F, 0x4D, 0x09, 0x1E, 0x92}}

type IEducationSettingsStaticsInterface interface {
	win32.IInspectableInterface
	Get_IsEducationEnvironment() bool
}

type IEducationSettingsStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_IsEducationEnvironment uintptr
}

type IEducationSettingsStatics struct {
	win32.IInspectable
}

func (this *IEducationSettingsStatics) Vtbl() *IEducationSettingsStaticsVtbl {
	return (*IEducationSettingsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IEducationSettingsStatics) Get_IsEducationEnvironment() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEducationEnvironment, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 971260E0-F170-4A42-BD55-A900B212DAE2
var IID_IHardwareIdentificationStatics = syscall.GUID{0x971260E0, 0xF170, 0x4A42,
	[8]byte{0xBD, 0x55, 0xA9, 0x00, 0xB2, 0x12, 0xDA, 0xE2}}

type IHardwareIdentificationStaticsInterface interface {
	win32.IInspectableInterface
	GetPackageSpecificToken(nonce *IBuffer) *IHardwareToken
}

type IHardwareIdentificationStaticsVtbl struct {
	win32.IInspectableVtbl
	GetPackageSpecificToken uintptr
}

type IHardwareIdentificationStatics struct {
	win32.IInspectable
}

func (this *IHardwareIdentificationStatics) Vtbl() *IHardwareIdentificationStaticsVtbl {
	return (*IHardwareIdentificationStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHardwareIdentificationStatics) GetPackageSpecificToken(nonce *IBuffer) *IHardwareToken {
	var _result *IHardwareToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPackageSpecificToken, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(nonce)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 28F6D4C0-FB12-40A4-8167-7F4E03D2724C
var IID_IHardwareToken = syscall.GUID{0x28F6D4C0, 0xFB12, 0x40A4,
	[8]byte{0x81, 0x67, 0x7F, 0x4E, 0x03, 0xD2, 0x72, 0x4C}}

type IHardwareTokenInterface interface {
	win32.IInspectableInterface
	Get_Id() *IBuffer
	Get_Signature() *IBuffer
	Get_Certificate() *IBuffer
}

type IHardwareTokenVtbl struct {
	win32.IInspectableVtbl
	Get_Id          uintptr
	Get_Signature   uintptr
	Get_Certificate uintptr
}

type IHardwareToken struct {
	win32.IInspectable
}

func (this *IHardwareToken) Vtbl() *IHardwareTokenVtbl {
	return (*IHardwareTokenVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHardwareToken) Get_Id() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHardwareToken) Get_Signature() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Signature, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHardwareToken) Get_Certificate() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Certificate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 99571178-500F-487E-8E75-29E551728712
var IID_IKnownRetailInfoPropertiesStatics = syscall.GUID{0x99571178, 0x500F, 0x487E,
	[8]byte{0x8E, 0x75, 0x29, 0xE5, 0x51, 0x72, 0x87, 0x12}}

type IKnownRetailInfoPropertiesStaticsInterface interface {
	win32.IInspectableInterface
	Get_RetailAccessCode() string
	Get_ManufacturerName() string
	Get_ModelName() string
	Get_DisplayModelName() string
	Get_Price() string
	Get_IsFeatured() string
	Get_FormFactor() string
	Get_ScreenSize() string
	Get_Weight() string
	Get_DisplayDescription() string
	Get_BatteryLifeDescription() string
	Get_ProcessorDescription() string
	Get_Memory() string
	Get_StorageDescription() string
	Get_GraphicsDescription() string
	Get_FrontCameraDescription() string
	Get_RearCameraDescription() string
	Get_HasNfc() string
	Get_HasSdSlot() string
	Get_HasOpticalDrive() string
	Get_IsOfficeInstalled() string
	Get_WindowsEdition() string
}

type IKnownRetailInfoPropertiesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_RetailAccessCode       uintptr
	Get_ManufacturerName       uintptr
	Get_ModelName              uintptr
	Get_DisplayModelName       uintptr
	Get_Price                  uintptr
	Get_IsFeatured             uintptr
	Get_FormFactor             uintptr
	Get_ScreenSize             uintptr
	Get_Weight                 uintptr
	Get_DisplayDescription     uintptr
	Get_BatteryLifeDescription uintptr
	Get_ProcessorDescription   uintptr
	Get_Memory                 uintptr
	Get_StorageDescription     uintptr
	Get_GraphicsDescription    uintptr
	Get_FrontCameraDescription uintptr
	Get_RearCameraDescription  uintptr
	Get_HasNfc                 uintptr
	Get_HasSdSlot              uintptr
	Get_HasOpticalDrive        uintptr
	Get_IsOfficeInstalled      uintptr
	Get_WindowsEdition         uintptr
}

type IKnownRetailInfoPropertiesStatics struct {
	win32.IInspectable
}

func (this *IKnownRetailInfoPropertiesStatics) Vtbl() *IKnownRetailInfoPropertiesStaticsVtbl {
	return (*IKnownRetailInfoPropertiesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKnownRetailInfoPropertiesStatics) Get_RetailAccessCode() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RetailAccessCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownRetailInfoPropertiesStatics) Get_ManufacturerName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ManufacturerName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownRetailInfoPropertiesStatics) Get_ModelName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ModelName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownRetailInfoPropertiesStatics) Get_DisplayModelName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayModelName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownRetailInfoPropertiesStatics) Get_Price() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Price, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownRetailInfoPropertiesStatics) Get_IsFeatured() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsFeatured, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownRetailInfoPropertiesStatics) Get_FormFactor() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FormFactor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownRetailInfoPropertiesStatics) Get_ScreenSize() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScreenSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownRetailInfoPropertiesStatics) Get_Weight() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Weight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownRetailInfoPropertiesStatics) Get_DisplayDescription() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownRetailInfoPropertiesStatics) Get_BatteryLifeDescription() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BatteryLifeDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownRetailInfoPropertiesStatics) Get_ProcessorDescription() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProcessorDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownRetailInfoPropertiesStatics) Get_Memory() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Memory, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownRetailInfoPropertiesStatics) Get_StorageDescription() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StorageDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownRetailInfoPropertiesStatics) Get_GraphicsDescription() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GraphicsDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownRetailInfoPropertiesStatics) Get_FrontCameraDescription() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrontCameraDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownRetailInfoPropertiesStatics) Get_RearCameraDescription() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RearCameraDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownRetailInfoPropertiesStatics) Get_HasNfc() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasNfc, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownRetailInfoPropertiesStatics) Get_HasSdSlot() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasSdSlot, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownRetailInfoPropertiesStatics) Get_HasOpticalDrive() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasOpticalDrive, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownRetailInfoPropertiesStatics) Get_IsOfficeInstalled() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsOfficeInstalled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownRetailInfoPropertiesStatics) Get_WindowsEdition() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WindowsEdition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// B6E24C1B-7B1C-4B32-8C62-A66597CE723A
var IID_IPlatformDiagnosticsAndUsageDataSettingsStatics = syscall.GUID{0xB6E24C1B, 0x7B1C, 0x4B32,
	[8]byte{0x8C, 0x62, 0xA6, 0x65, 0x97, 0xCE, 0x72, 0x3A}}

type IPlatformDiagnosticsAndUsageDataSettingsStaticsInterface interface {
	win32.IInspectableInterface
	Get_CollectionLevel() PlatformDataCollectionLevel
	Add_CollectionLevelChanged(handler EventHandler[interface{}]) EventRegistrationToken
	Remove_CollectionLevelChanged(token EventRegistrationToken)
	CanCollectDiagnostics(level PlatformDataCollectionLevel) bool
}

type IPlatformDiagnosticsAndUsageDataSettingsStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_CollectionLevel           uintptr
	Add_CollectionLevelChanged    uintptr
	Remove_CollectionLevelChanged uintptr
	CanCollectDiagnostics         uintptr
}

type IPlatformDiagnosticsAndUsageDataSettingsStatics struct {
	win32.IInspectable
}

func (this *IPlatformDiagnosticsAndUsageDataSettingsStatics) Vtbl() *IPlatformDiagnosticsAndUsageDataSettingsStaticsVtbl {
	return (*IPlatformDiagnosticsAndUsageDataSettingsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPlatformDiagnosticsAndUsageDataSettingsStatics) Get_CollectionLevel() PlatformDataCollectionLevel {
	var _result PlatformDataCollectionLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CollectionLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPlatformDiagnosticsAndUsageDataSettingsStatics) Add_CollectionLevelChanged(handler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_CollectionLevelChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPlatformDiagnosticsAndUsageDataSettingsStatics) Remove_CollectionLevelChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_CollectionLevelChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPlatformDiagnosticsAndUsageDataSettingsStatics) CanCollectDiagnostics(level PlatformDataCollectionLevel) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CanCollectDiagnostics, uintptr(unsafe.Pointer(this)), uintptr(level), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 0712C6B8-8B92-4F2A-8499-031F1798D6EF
var IID_IRetailInfoStatics = syscall.GUID{0x0712C6B8, 0x8B92, 0x4F2A,
	[8]byte{0x84, 0x99, 0x03, 0x1F, 0x17, 0x98, 0xD6, 0xEF}}

type IRetailInfoStaticsInterface interface {
	win32.IInspectableInterface
	Get_IsDemoModeEnabled() bool
	Get_Properties() *IMapView[string, interface{}]
}

type IRetailInfoStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_IsDemoModeEnabled uintptr
	Get_Properties        uintptr
}

type IRetailInfoStatics struct {
	win32.IInspectable
}

func (this *IRetailInfoStatics) Vtbl() *IRetailInfoStaticsVtbl {
	return (*IRetailInfoStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRetailInfoStatics) Get_IsDemoModeEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDemoModeEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRetailInfoStatics) Get_Properties() *IMapView[string, interface{}] {
	var _result *IMapView[string, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 893DF40E-CAD6-4D50-8C49-6FCFC03EDB29
var IID_ISharedModeSettingsStatics = syscall.GUID{0x893DF40E, 0xCAD6, 0x4D50,
	[8]byte{0x8C, 0x49, 0x6F, 0xCF, 0xC0, 0x3E, 0xDB, 0x29}}

type ISharedModeSettingsStaticsInterface interface {
	win32.IInspectableInterface
	Get_IsEnabled() bool
}

type ISharedModeSettingsStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_IsEnabled uintptr
}

type ISharedModeSettingsStatics struct {
	win32.IInspectable
}

func (this *ISharedModeSettingsStatics) Vtbl() *ISharedModeSettingsStaticsVtbl {
	return (*ISharedModeSettingsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISharedModeSettingsStatics) Get_IsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 608988A4-CCF1-4EE8-A5E2-FD6A1D0CFAC8
var IID_ISharedModeSettingsStatics2 = syscall.GUID{0x608988A4, 0xCCF1, 0x4EE8,
	[8]byte{0xA5, 0xE2, 0xFD, 0x6A, 0x1D, 0x0C, 0xFA, 0xC8}}

type ISharedModeSettingsStatics2Interface interface {
	win32.IInspectableInterface
	Get_ShouldAvoidLocalStorage() bool
}

type ISharedModeSettingsStatics2Vtbl struct {
	win32.IInspectableVtbl
	Get_ShouldAvoidLocalStorage uintptr
}

type ISharedModeSettingsStatics2 struct {
	win32.IInspectable
}

func (this *ISharedModeSettingsStatics2) Vtbl() *ISharedModeSettingsStatics2Vtbl {
	return (*ISharedModeSettingsStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISharedModeSettingsStatics2) Get_ShouldAvoidLocalStorage() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ShouldAvoidLocalStorage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 0C659E7D-C3C2-4D33-A2DF-21BC41916EB3
var IID_ISystemIdentificationInfo = syscall.GUID{0x0C659E7D, 0xC3C2, 0x4D33,
	[8]byte{0xA2, 0xDF, 0x21, 0xBC, 0x41, 0x91, 0x6E, 0xB3}}

type ISystemIdentificationInfoInterface interface {
	win32.IInspectableInterface
	Get_Id() *IBuffer
	Get_Source() SystemIdentificationSource
}

type ISystemIdentificationInfoVtbl struct {
	win32.IInspectableVtbl
	Get_Id     uintptr
	Get_Source uintptr
}

type ISystemIdentificationInfo struct {
	win32.IInspectable
}

func (this *ISystemIdentificationInfo) Vtbl() *ISystemIdentificationInfoVtbl {
	return (*ISystemIdentificationInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemIdentificationInfo) Get_Id() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISystemIdentificationInfo) Get_Source() SystemIdentificationSource {
	var _result SystemIdentificationSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Source, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5581F42A-D3DF-4D93-A37D-C41A616C6D01
var IID_ISystemIdentificationStatics = syscall.GUID{0x5581F42A, 0xD3DF, 0x4D93,
	[8]byte{0xA3, 0x7D, 0xC4, 0x1A, 0x61, 0x6C, 0x6D, 0x01}}

type ISystemIdentificationStaticsInterface interface {
	win32.IInspectableInterface
	GetSystemIdForPublisher() *ISystemIdentificationInfo
	GetSystemIdForUser(user *IUser) *ISystemIdentificationInfo
}

type ISystemIdentificationStaticsVtbl struct {
	win32.IInspectableVtbl
	GetSystemIdForPublisher uintptr
	GetSystemIdForUser      uintptr
}

type ISystemIdentificationStatics struct {
	win32.IInspectable
}

func (this *ISystemIdentificationStatics) Vtbl() *ISystemIdentificationStaticsVtbl {
	return (*ISystemIdentificationStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemIdentificationStatics) GetSystemIdForPublisher() *ISystemIdentificationInfo {
	var _result *ISystemIdentificationInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSystemIdForPublisher, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISystemIdentificationStatics) GetSystemIdForUser(user *IUser) *ISystemIdentificationInfo {
	var _result *ISystemIdentificationInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSystemIdForUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2C9620A8-1D88-5E2D-A324-A543AF4247EE
var IID_ISystemSetupInfoStatics = syscall.GUID{0x2C9620A8, 0x1D88, 0x5E2D,
	[8]byte{0xA3, 0x24, 0xA5, 0x43, 0xAF, 0x42, 0x47, 0xEE}}

type ISystemSetupInfoStaticsInterface interface {
	win32.IInspectableInterface
	Get_OutOfBoxExperienceState() SystemOutOfBoxExperienceState
	Add_OutOfBoxExperienceStateChanged(handler EventHandler[interface{}]) EventRegistrationToken
	Remove_OutOfBoxExperienceStateChanged(token EventRegistrationToken)
}

type ISystemSetupInfoStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_OutOfBoxExperienceState           uintptr
	Add_OutOfBoxExperienceStateChanged    uintptr
	Remove_OutOfBoxExperienceStateChanged uintptr
}

type ISystemSetupInfoStatics struct {
	win32.IInspectable
}

func (this *ISystemSetupInfoStatics) Vtbl() *ISystemSetupInfoStaticsVtbl {
	return (*ISystemSetupInfoStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISystemSetupInfoStatics) Get_OutOfBoxExperienceState() SystemOutOfBoxExperienceState {
	var _result SystemOutOfBoxExperienceState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutOfBoxExperienceState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemSetupInfoStatics) Add_OutOfBoxExperienceStateChanged(handler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_OutOfBoxExperienceStateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISystemSetupInfoStatics) Remove_OutOfBoxExperienceStateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_OutOfBoxExperienceStateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 6182445C-894B-5CBC-8976-A98E0A9B998D
var IID_IUnsupportedAppRequirement = syscall.GUID{0x6182445C, 0x894B, 0x5CBC,
	[8]byte{0x89, 0x76, 0xA9, 0x8E, 0x0A, 0x9B, 0x99, 0x8D}}

type IUnsupportedAppRequirementInterface interface {
	win32.IInspectableInterface
	Get_Requirement() string
	Get_Reasons() UnsupportedAppRequirementReasons
}

type IUnsupportedAppRequirementVtbl struct {
	win32.IInspectableVtbl
	Get_Requirement uintptr
	Get_Reasons     uintptr
}

type IUnsupportedAppRequirement struct {
	win32.IInspectable
}

func (this *IUnsupportedAppRequirement) Vtbl() *IUnsupportedAppRequirementVtbl {
	return (*IUnsupportedAppRequirementVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUnsupportedAppRequirement) Get_Requirement() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Requirement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUnsupportedAppRequirement) Get_Reasons() UnsupportedAppRequirementReasons {
	var _result UnsupportedAppRequirementReasons
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Reasons, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7D1D81DB-8D63-4789-9EA5-DDCF65A94F3C
var IID_IWindowsIntegrityPolicyStatics = syscall.GUID{0x7D1D81DB, 0x8D63, 0x4789,
	[8]byte{0x9E, 0xA5, 0xDD, 0xCF, 0x65, 0xA9, 0x4F, 0x3C}}

type IWindowsIntegrityPolicyStaticsInterface interface {
	win32.IInspectableInterface
	Get_IsEnabled() bool
	Get_IsEnabledForTrial() bool
	Get_CanDisable() bool
	Get_IsDisableSupported() bool
	Add_PolicyChanged(handler EventHandler[interface{}]) EventRegistrationToken
	Remove_PolicyChanged(token EventRegistrationToken)
}

type IWindowsIntegrityPolicyStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_IsEnabled          uintptr
	Get_IsEnabledForTrial  uintptr
	Get_CanDisable         uintptr
	Get_IsDisableSupported uintptr
	Add_PolicyChanged      uintptr
	Remove_PolicyChanged   uintptr
}

type IWindowsIntegrityPolicyStatics struct {
	win32.IInspectable
}

func (this *IWindowsIntegrityPolicyStatics) Vtbl() *IWindowsIntegrityPolicyStaticsVtbl {
	return (*IWindowsIntegrityPolicyStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWindowsIntegrityPolicyStatics) Get_IsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWindowsIntegrityPolicyStatics) Get_IsEnabledForTrial() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEnabledForTrial, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWindowsIntegrityPolicyStatics) Get_CanDisable() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanDisable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWindowsIntegrityPolicyStatics) Get_IsDisableSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDisableSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWindowsIntegrityPolicyStatics) Add_PolicyChanged(handler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PolicyChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWindowsIntegrityPolicyStatics) Remove_PolicyChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PolicyChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// classes

type PlatformDiagnosticsAndUsageDataSettings struct {
	RtClass
}

func NewIPlatformDiagnosticsAndUsageDataSettingsStatics() *IPlatformDiagnosticsAndUsageDataSettingsStatics {
	var p *IPlatformDiagnosticsAndUsageDataSettingsStatics
	hs := NewHStr("Windows.System.Profile.PlatformDiagnosticsAndUsageDataSettings")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPlatformDiagnosticsAndUsageDataSettingsStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type SharedModeSettings struct {
	RtClass
}

func NewISharedModeSettingsStatics2() *ISharedModeSettingsStatics2 {
	var p *ISharedModeSettingsStatics2
	hs := NewHStr("Windows.System.Profile.SharedModeSettings")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISharedModeSettingsStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewISharedModeSettingsStatics() *ISharedModeSettingsStatics {
	var p *ISharedModeSettingsStatics
	hs := NewHStr("Windows.System.Profile.SharedModeSettings")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISharedModeSettingsStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type SystemSetupInfo struct {
	RtClass
}

func NewISystemSetupInfoStatics() *ISystemSetupInfoStatics {
	var p *ISystemSetupInfoStatics
	hs := NewHStr("Windows.System.Profile.SystemSetupInfo")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISystemSetupInfoStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
