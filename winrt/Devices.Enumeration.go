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
type DeviceAccessStatus int32

const (
	DeviceAccessStatus_Unspecified    DeviceAccessStatus = 0
	DeviceAccessStatus_Allowed        DeviceAccessStatus = 1
	DeviceAccessStatus_DeniedByUser   DeviceAccessStatus = 2
	DeviceAccessStatus_DeniedBySystem DeviceAccessStatus = 3
)

// enum
type DeviceClass int32

const (
	DeviceClass_All                   DeviceClass = 0
	DeviceClass_AudioCapture          DeviceClass = 1
	DeviceClass_AudioRender           DeviceClass = 2
	DeviceClass_PortableStorageDevice DeviceClass = 3
	DeviceClass_VideoCapture          DeviceClass = 4
	DeviceClass_ImageScanner          DeviceClass = 5
	DeviceClass_Location              DeviceClass = 6
)

// enum
type DeviceInformationKind int32

const (
	DeviceInformationKind_Unknown                      DeviceInformationKind = 0
	DeviceInformationKind_DeviceInterface              DeviceInformationKind = 1
	DeviceInformationKind_DeviceContainer              DeviceInformationKind = 2
	DeviceInformationKind_Device                       DeviceInformationKind = 3
	DeviceInformationKind_DeviceInterfaceClass         DeviceInformationKind = 4
	DeviceInformationKind_AssociationEndpoint          DeviceInformationKind = 5
	DeviceInformationKind_AssociationEndpointContainer DeviceInformationKind = 6
	DeviceInformationKind_AssociationEndpointService   DeviceInformationKind = 7
	DeviceInformationKind_DevicePanel                  DeviceInformationKind = 8
)

// enum
// flags
type DevicePairingKinds uint32

const (
	DevicePairingKinds_None                      DevicePairingKinds = 0
	DevicePairingKinds_ConfirmOnly               DevicePairingKinds = 1
	DevicePairingKinds_DisplayPin                DevicePairingKinds = 2
	DevicePairingKinds_ProvidePin                DevicePairingKinds = 4
	DevicePairingKinds_ConfirmPinMatch           DevicePairingKinds = 8
	DevicePairingKinds_ProvidePasswordCredential DevicePairingKinds = 16
)

// enum
type DevicePairingProtectionLevel int32

const (
	DevicePairingProtectionLevel_Default                     DevicePairingProtectionLevel = 0
	DevicePairingProtectionLevel_None                        DevicePairingProtectionLevel = 1
	DevicePairingProtectionLevel_Encryption                  DevicePairingProtectionLevel = 2
	DevicePairingProtectionLevel_EncryptionAndAuthentication DevicePairingProtectionLevel = 3
)

// enum
type DevicePairingResultStatus int32

const (
	DevicePairingResultStatus_Paired                       DevicePairingResultStatus = 0
	DevicePairingResultStatus_NotReadyToPair               DevicePairingResultStatus = 1
	DevicePairingResultStatus_NotPaired                    DevicePairingResultStatus = 2
	DevicePairingResultStatus_AlreadyPaired                DevicePairingResultStatus = 3
	DevicePairingResultStatus_ConnectionRejected           DevicePairingResultStatus = 4
	DevicePairingResultStatus_TooManyConnections           DevicePairingResultStatus = 5
	DevicePairingResultStatus_HardwareFailure              DevicePairingResultStatus = 6
	DevicePairingResultStatus_AuthenticationTimeout        DevicePairingResultStatus = 7
	DevicePairingResultStatus_AuthenticationNotAllowed     DevicePairingResultStatus = 8
	DevicePairingResultStatus_AuthenticationFailure        DevicePairingResultStatus = 9
	DevicePairingResultStatus_NoSupportedProfiles          DevicePairingResultStatus = 10
	DevicePairingResultStatus_ProtectionLevelCouldNotBeMet DevicePairingResultStatus = 11
	DevicePairingResultStatus_AccessDenied                 DevicePairingResultStatus = 12
	DevicePairingResultStatus_InvalidCeremonyData          DevicePairingResultStatus = 13
	DevicePairingResultStatus_PairingCanceled              DevicePairingResultStatus = 14
	DevicePairingResultStatus_OperationAlreadyInProgress   DevicePairingResultStatus = 15
	DevicePairingResultStatus_RequiredHandlerNotRegistered DevicePairingResultStatus = 16
	DevicePairingResultStatus_RejectedByHandler            DevicePairingResultStatus = 17
	DevicePairingResultStatus_RemoteDeviceHasAssociation   DevicePairingResultStatus = 18
	DevicePairingResultStatus_Failed                       DevicePairingResultStatus = 19
)

// enum
// flags
type DevicePickerDisplayStatusOptions uint32

const (
	DevicePickerDisplayStatusOptions_None                 DevicePickerDisplayStatusOptions = 0
	DevicePickerDisplayStatusOptions_ShowProgress         DevicePickerDisplayStatusOptions = 1
	DevicePickerDisplayStatusOptions_ShowDisconnectButton DevicePickerDisplayStatusOptions = 2
	DevicePickerDisplayStatusOptions_ShowRetryButton      DevicePickerDisplayStatusOptions = 4
)

// enum
type DeviceUnpairingResultStatus int32

const (
	DeviceUnpairingResultStatus_Unpaired                   DeviceUnpairingResultStatus = 0
	DeviceUnpairingResultStatus_AlreadyUnpaired            DeviceUnpairingResultStatus = 1
	DeviceUnpairingResultStatus_OperationAlreadyInProgress DeviceUnpairingResultStatus = 2
	DeviceUnpairingResultStatus_AccessDenied               DeviceUnpairingResultStatus = 3
	DeviceUnpairingResultStatus_Failed                     DeviceUnpairingResultStatus = 4
)

// enum
type DeviceWatcherEventKind int32

const (
	DeviceWatcherEventKind_Add    DeviceWatcherEventKind = 0
	DeviceWatcherEventKind_Update DeviceWatcherEventKind = 1
	DeviceWatcherEventKind_Remove DeviceWatcherEventKind = 2
)

// enum
type DeviceWatcherStatus int32

const (
	DeviceWatcherStatus_Created              DeviceWatcherStatus = 0
	DeviceWatcherStatus_Started              DeviceWatcherStatus = 1
	DeviceWatcherStatus_EnumerationCompleted DeviceWatcherStatus = 2
	DeviceWatcherStatus_Stopping             DeviceWatcherStatus = 3
	DeviceWatcherStatus_Stopped              DeviceWatcherStatus = 4
	DeviceWatcherStatus_Aborted              DeviceWatcherStatus = 5
)

// enum
type Panel int32

const (
	Panel_Unknown Panel = 0
	Panel_Front   Panel = 1
	Panel_Back    Panel = 2
	Panel_Top     Panel = 3
	Panel_Bottom  Panel = 4
	Panel_Left    Panel = 5
	Panel_Right   Panel = 6
)

// interfaces

// DEDA0BCC-4F9D-4F58-9DBA-A9BC800408D5
var IID_IDeviceAccessChangedEventArgs = syscall.GUID{0xDEDA0BCC, 0x4F9D, 0x4F58,
	[8]byte{0x9D, 0xBA, 0xA9, 0xBC, 0x80, 0x04, 0x08, 0xD5}}

type IDeviceAccessChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Status() DeviceAccessStatus
}

type IDeviceAccessChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
}

type IDeviceAccessChangedEventArgs struct {
	win32.IInspectable
}

func (this *IDeviceAccessChangedEventArgs) Vtbl() *IDeviceAccessChangedEventArgsVtbl {
	return (*IDeviceAccessChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceAccessChangedEventArgs) Get_Status() DeviceAccessStatus {
	var _result DeviceAccessStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 82523262-934B-4B30-A178-ADC39F2F2BE3
var IID_IDeviceAccessChangedEventArgs2 = syscall.GUID{0x82523262, 0x934B, 0x4B30,
	[8]byte{0xA1, 0x78, 0xAD, 0xC3, 0x9F, 0x2F, 0x2B, 0xE3}}

type IDeviceAccessChangedEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_Id() string
}

type IDeviceAccessChangedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_Id uintptr
}

type IDeviceAccessChangedEventArgs2 struct {
	win32.IInspectable
}

func (this *IDeviceAccessChangedEventArgs2) Vtbl() *IDeviceAccessChangedEventArgs2Vtbl {
	return (*IDeviceAccessChangedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceAccessChangedEventArgs2) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 0BAA9A73-6DE5-4915-8DDD-9A0554A6F545
var IID_IDeviceAccessInformation = syscall.GUID{0x0BAA9A73, 0x6DE5, 0x4915,
	[8]byte{0x8D, 0xDD, 0x9A, 0x05, 0x54, 0xA6, 0xF5, 0x45}}

type IDeviceAccessInformationInterface interface {
	win32.IInspectableInterface
	Add_AccessChanged(handler TypedEventHandler[*IDeviceAccessInformation, *IDeviceAccessChangedEventArgs]) EventRegistrationToken
	Remove_AccessChanged(cookie EventRegistrationToken)
	Get_CurrentStatus() DeviceAccessStatus
}

type IDeviceAccessInformationVtbl struct {
	win32.IInspectableVtbl
	Add_AccessChanged    uintptr
	Remove_AccessChanged uintptr
	Get_CurrentStatus    uintptr
}

type IDeviceAccessInformation struct {
	win32.IInspectable
}

func (this *IDeviceAccessInformation) Vtbl() *IDeviceAccessInformationVtbl {
	return (*IDeviceAccessInformationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceAccessInformation) Add_AccessChanged(handler TypedEventHandler[*IDeviceAccessInformation, *IDeviceAccessChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AccessChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDeviceAccessInformation) Remove_AccessChanged(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AccessChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IDeviceAccessInformation) Get_CurrentStatus() DeviceAccessStatus {
	var _result DeviceAccessStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 574BD3D3-5F30-45CD-8A94-724FE5973084
var IID_IDeviceAccessInformationStatics = syscall.GUID{0x574BD3D3, 0x5F30, 0x45CD,
	[8]byte{0x8A, 0x94, 0x72, 0x4F, 0xE5, 0x97, 0x30, 0x84}}

type IDeviceAccessInformationStaticsInterface interface {
	win32.IInspectableInterface
	CreateFromId(deviceId string) *IDeviceAccessInformation
	CreateFromDeviceClassId(deviceClassId syscall.GUID) *IDeviceAccessInformation
	CreateFromDeviceClass(deviceClass DeviceClass) *IDeviceAccessInformation
}

type IDeviceAccessInformationStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateFromId            uintptr
	CreateFromDeviceClassId uintptr
	CreateFromDeviceClass   uintptr
}

type IDeviceAccessInformationStatics struct {
	win32.IInspectable
}

func (this *IDeviceAccessInformationStatics) Vtbl() *IDeviceAccessInformationStaticsVtbl {
	return (*IDeviceAccessInformationStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceAccessInformationStatics) CreateFromId(deviceId string) *IDeviceAccessInformation {
	var _result *IDeviceAccessInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromId, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDeviceAccessInformationStatics) CreateFromDeviceClassId(deviceClassId syscall.GUID) *IDeviceAccessInformation {
	var _result *IDeviceAccessInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromDeviceClassId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&deviceClassId)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDeviceAccessInformationStatics) CreateFromDeviceClass(deviceClass DeviceClass) *IDeviceAccessInformation {
	var _result *IDeviceAccessInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromDeviceClass, uintptr(unsafe.Pointer(this)), uintptr(deviceClass), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B8578C0C-BBC1-484B-BFFA-7B31DCC200B2
var IID_IDeviceConnectionChangeTriggerDetails = syscall.GUID{0xB8578C0C, 0xBBC1, 0x484B,
	[8]byte{0xBF, 0xFA, 0x7B, 0x31, 0xDC, 0xC2, 0x00, 0xB2}}

type IDeviceConnectionChangeTriggerDetailsInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
}

type IDeviceConnectionChangeTriggerDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId uintptr
}

type IDeviceConnectionChangeTriggerDetails struct {
	win32.IInspectable
}

func (this *IDeviceConnectionChangeTriggerDetails) Vtbl() *IDeviceConnectionChangeTriggerDetailsVtbl {
	return (*IDeviceConnectionChangeTriggerDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceConnectionChangeTriggerDetails) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 8E44B56D-F902-4A00-B536-F37992E6A2A7
var IID_IDeviceDisconnectButtonClickedEventArgs = syscall.GUID{0x8E44B56D, 0xF902, 0x4A00,
	[8]byte{0xB5, 0x36, 0xF3, 0x79, 0x92, 0xE6, 0xA2, 0xA7}}

type IDeviceDisconnectButtonClickedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Device() *IDeviceInformation
}

type IDeviceDisconnectButtonClickedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Device uintptr
}

type IDeviceDisconnectButtonClickedEventArgs struct {
	win32.IInspectable
}

func (this *IDeviceDisconnectButtonClickedEventArgs) Vtbl() *IDeviceDisconnectButtonClickedEventArgsVtbl {
	return (*IDeviceDisconnectButtonClickedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceDisconnectButtonClickedEventArgs) Get_Device() *IDeviceInformation {
	var _result *IDeviceInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Device, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// ABA0FB95-4398-489D-8E44-E6130927011F
var IID_IDeviceInformation = syscall.GUID{0xABA0FB95, 0x4398, 0x489D,
	[8]byte{0x8E, 0x44, 0xE6, 0x13, 0x09, 0x27, 0x01, 0x1F}}

type IDeviceInformationInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	Get_Name() string
	Get_IsEnabled() bool
	Get_IsDefault() bool
	Get_EnclosureLocation() *IEnclosureLocation
	Get_Properties() *IMapView[string, interface{}]
	Update(updateInfo *IDeviceInformationUpdate)
	GetThumbnailAsync() *IAsyncOperation[*IRandomAccessStreamWithContentType]
	GetGlyphThumbnailAsync() *IAsyncOperation[*IRandomAccessStreamWithContentType]
}

type IDeviceInformationVtbl struct {
	win32.IInspectableVtbl
	Get_Id                 uintptr
	Get_Name               uintptr
	Get_IsEnabled          uintptr
	Get_IsDefault          uintptr
	Get_EnclosureLocation  uintptr
	Get_Properties         uintptr
	Update                 uintptr
	GetThumbnailAsync      uintptr
	GetGlyphThumbnailAsync uintptr
}

type IDeviceInformation struct {
	win32.IInspectable
}

func (this *IDeviceInformation) Vtbl() *IDeviceInformationVtbl {
	return (*IDeviceInformationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceInformation) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDeviceInformation) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDeviceInformation) Get_IsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDeviceInformation) Get_IsDefault() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDeviceInformation) Get_EnclosureLocation() *IEnclosureLocation {
	var _result *IEnclosureLocation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EnclosureLocation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDeviceInformation) Get_Properties() *IMapView[string, interface{}] {
	var _result *IMapView[string, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDeviceInformation) Update(updateInfo *IDeviceInformationUpdate) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Update, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(updateInfo)))
	_ = _hr
}

func (this *IDeviceInformation) GetThumbnailAsync() *IAsyncOperation[*IRandomAccessStreamWithContentType] {
	var _result *IAsyncOperation[*IRandomAccessStreamWithContentType]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetThumbnailAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDeviceInformation) GetGlyphThumbnailAsync() *IAsyncOperation[*IRandomAccessStreamWithContentType] {
	var _result *IAsyncOperation[*IRandomAccessStreamWithContentType]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetGlyphThumbnailAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F156A638-7997-48D9-A10C-269D46533F48
var IID_IDeviceInformation2 = syscall.GUID{0xF156A638, 0x7997, 0x48D9,
	[8]byte{0xA1, 0x0C, 0x26, 0x9D, 0x46, 0x53, 0x3F, 0x48}}

type IDeviceInformation2Interface interface {
	win32.IInspectableInterface
	Get_Kind() DeviceInformationKind
	Get_Pairing() *IDeviceInformationPairing
}

type IDeviceInformation2Vtbl struct {
	win32.IInspectableVtbl
	Get_Kind    uintptr
	Get_Pairing uintptr
}

type IDeviceInformation2 struct {
	win32.IInspectable
}

func (this *IDeviceInformation2) Vtbl() *IDeviceInformation2Vtbl {
	return (*IDeviceInformation2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceInformation2) Get_Kind() DeviceInformationKind {
	var _result DeviceInformationKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Kind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDeviceInformation2) Get_Pairing() *IDeviceInformationPairing {
	var _result *IDeviceInformationPairing
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Pairing, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 85138C02-4EE6-4914-8370-107A39144C0E
var IID_IDeviceInformationCustomPairing = syscall.GUID{0x85138C02, 0x4EE6, 0x4914,
	[8]byte{0x83, 0x70, 0x10, 0x7A, 0x39, 0x14, 0x4C, 0x0E}}

type IDeviceInformationCustomPairingInterface interface {
	win32.IInspectableInterface
	PairAsync(pairingKindsSupported DevicePairingKinds) *IAsyncOperation[*IDevicePairingResult]
	PairWithProtectionLevelAsync(pairingKindsSupported DevicePairingKinds, minProtectionLevel DevicePairingProtectionLevel) *IAsyncOperation[*IDevicePairingResult]
	PairWithProtectionLevelAndSettingsAsync(pairingKindsSupported DevicePairingKinds, minProtectionLevel DevicePairingProtectionLevel, devicePairingSettings *IDevicePairingSettings) *IAsyncOperation[*IDevicePairingResult]
	Add_PairingRequested(handler TypedEventHandler[*IDeviceInformationCustomPairing, *IDevicePairingRequestedEventArgs]) EventRegistrationToken
	Remove_PairingRequested(token EventRegistrationToken)
}

type IDeviceInformationCustomPairingVtbl struct {
	win32.IInspectableVtbl
	PairAsync                               uintptr
	PairWithProtectionLevelAsync            uintptr
	PairWithProtectionLevelAndSettingsAsync uintptr
	Add_PairingRequested                    uintptr
	Remove_PairingRequested                 uintptr
}

type IDeviceInformationCustomPairing struct {
	win32.IInspectable
}

func (this *IDeviceInformationCustomPairing) Vtbl() *IDeviceInformationCustomPairingVtbl {
	return (*IDeviceInformationCustomPairingVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceInformationCustomPairing) PairAsync(pairingKindsSupported DevicePairingKinds) *IAsyncOperation[*IDevicePairingResult] {
	var _result *IAsyncOperation[*IDevicePairingResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PairAsync, uintptr(unsafe.Pointer(this)), uintptr(pairingKindsSupported), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDeviceInformationCustomPairing) PairWithProtectionLevelAsync(pairingKindsSupported DevicePairingKinds, minProtectionLevel DevicePairingProtectionLevel) *IAsyncOperation[*IDevicePairingResult] {
	var _result *IAsyncOperation[*IDevicePairingResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PairWithProtectionLevelAsync, uintptr(unsafe.Pointer(this)), uintptr(pairingKindsSupported), uintptr(minProtectionLevel), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDeviceInformationCustomPairing) PairWithProtectionLevelAndSettingsAsync(pairingKindsSupported DevicePairingKinds, minProtectionLevel DevicePairingProtectionLevel, devicePairingSettings *IDevicePairingSettings) *IAsyncOperation[*IDevicePairingResult] {
	var _result *IAsyncOperation[*IDevicePairingResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PairWithProtectionLevelAndSettingsAsync, uintptr(unsafe.Pointer(this)), uintptr(pairingKindsSupported), uintptr(minProtectionLevel), uintptr(unsafe.Pointer(devicePairingSettings)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDeviceInformationCustomPairing) Add_PairingRequested(handler TypedEventHandler[*IDeviceInformationCustomPairing, *IDevicePairingRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PairingRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDeviceInformationCustomPairing) Remove_PairingRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PairingRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 2C4769F5-F684-40D5-8469-E8DBAAB70485
var IID_IDeviceInformationPairing = syscall.GUID{0x2C4769F5, 0xF684, 0x40D5,
	[8]byte{0x84, 0x69, 0xE8, 0xDB, 0xAA, 0xB7, 0x04, 0x85}}

type IDeviceInformationPairingInterface interface {
	win32.IInspectableInterface
	Get_IsPaired() bool
	Get_CanPair() bool
	PairAsync() *IAsyncOperation[*IDevicePairingResult]
	PairWithProtectionLevelAsync(minProtectionLevel DevicePairingProtectionLevel) *IAsyncOperation[*IDevicePairingResult]
}

type IDeviceInformationPairingVtbl struct {
	win32.IInspectableVtbl
	Get_IsPaired                 uintptr
	Get_CanPair                  uintptr
	PairAsync                    uintptr
	PairWithProtectionLevelAsync uintptr
}

type IDeviceInformationPairing struct {
	win32.IInspectable
}

func (this *IDeviceInformationPairing) Vtbl() *IDeviceInformationPairingVtbl {
	return (*IDeviceInformationPairingVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceInformationPairing) Get_IsPaired() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPaired, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDeviceInformationPairing) Get_CanPair() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanPair, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDeviceInformationPairing) PairAsync() *IAsyncOperation[*IDevicePairingResult] {
	var _result *IAsyncOperation[*IDevicePairingResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PairAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDeviceInformationPairing) PairWithProtectionLevelAsync(minProtectionLevel DevicePairingProtectionLevel) *IAsyncOperation[*IDevicePairingResult] {
	var _result *IAsyncOperation[*IDevicePairingResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PairWithProtectionLevelAsync, uintptr(unsafe.Pointer(this)), uintptr(minProtectionLevel), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F68612FD-0AEE-4328-85CC-1C742BB1790D
var IID_IDeviceInformationPairing2 = syscall.GUID{0xF68612FD, 0x0AEE, 0x4328,
	[8]byte{0x85, 0xCC, 0x1C, 0x74, 0x2B, 0xB1, 0x79, 0x0D}}

type IDeviceInformationPairing2Interface interface {
	win32.IInspectableInterface
	Get_ProtectionLevel() DevicePairingProtectionLevel
	Get_Custom() *IDeviceInformationCustomPairing
	PairWithProtectionLevelAndSettingsAsync(minProtectionLevel DevicePairingProtectionLevel, devicePairingSettings *IDevicePairingSettings) *IAsyncOperation[*IDevicePairingResult]
	UnpairAsync() *IAsyncOperation[*IDeviceUnpairingResult]
}

type IDeviceInformationPairing2Vtbl struct {
	win32.IInspectableVtbl
	Get_ProtectionLevel                     uintptr
	Get_Custom                              uintptr
	PairWithProtectionLevelAndSettingsAsync uintptr
	UnpairAsync                             uintptr
}

type IDeviceInformationPairing2 struct {
	win32.IInspectable
}

func (this *IDeviceInformationPairing2) Vtbl() *IDeviceInformationPairing2Vtbl {
	return (*IDeviceInformationPairing2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceInformationPairing2) Get_ProtectionLevel() DevicePairingProtectionLevel {
	var _result DevicePairingProtectionLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDeviceInformationPairing2) Get_Custom() *IDeviceInformationCustomPairing {
	var _result *IDeviceInformationCustomPairing
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Custom, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDeviceInformationPairing2) PairWithProtectionLevelAndSettingsAsync(minProtectionLevel DevicePairingProtectionLevel, devicePairingSettings *IDevicePairingSettings) *IAsyncOperation[*IDevicePairingResult] {
	var _result *IAsyncOperation[*IDevicePairingResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PairWithProtectionLevelAndSettingsAsync, uintptr(unsafe.Pointer(this)), uintptr(minProtectionLevel), uintptr(unsafe.Pointer(devicePairingSettings)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDeviceInformationPairing2) UnpairAsync() *IAsyncOperation[*IDeviceUnpairingResult] {
	var _result *IAsyncOperation[*IDeviceUnpairingResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UnpairAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E915C408-36D4-49A1-BF13-514173799B6B
var IID_IDeviceInformationPairingStatics = syscall.GUID{0xE915C408, 0x36D4, 0x49A1,
	[8]byte{0xBF, 0x13, 0x51, 0x41, 0x73, 0x79, 0x9B, 0x6B}}

type IDeviceInformationPairingStaticsInterface interface {
	win32.IInspectableInterface
	TryRegisterForAllInboundPairingRequests(pairingKindsSupported DevicePairingKinds) bool
}

type IDeviceInformationPairingStaticsVtbl struct {
	win32.IInspectableVtbl
	TryRegisterForAllInboundPairingRequests uintptr
}

type IDeviceInformationPairingStatics struct {
	win32.IInspectable
}

func (this *IDeviceInformationPairingStatics) Vtbl() *IDeviceInformationPairingStaticsVtbl {
	return (*IDeviceInformationPairingStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceInformationPairingStatics) TryRegisterForAllInboundPairingRequests(pairingKindsSupported DevicePairingKinds) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryRegisterForAllInboundPairingRequests, uintptr(unsafe.Pointer(this)), uintptr(pairingKindsSupported), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 04DE5372-B7B7-476B-A74F-C5836A704D98
var IID_IDeviceInformationPairingStatics2 = syscall.GUID{0x04DE5372, 0xB7B7, 0x476B,
	[8]byte{0xA7, 0x4F, 0xC5, 0x83, 0x6A, 0x70, 0x4D, 0x98}}

type IDeviceInformationPairingStatics2Interface interface {
	win32.IInspectableInterface
	TryRegisterForAllInboundPairingRequestsWithProtectionLevel(pairingKindsSupported DevicePairingKinds, minProtectionLevel DevicePairingProtectionLevel) bool
}

type IDeviceInformationPairingStatics2Vtbl struct {
	win32.IInspectableVtbl
	TryRegisterForAllInboundPairingRequestsWithProtectionLevel uintptr
}

type IDeviceInformationPairingStatics2 struct {
	win32.IInspectable
}

func (this *IDeviceInformationPairingStatics2) Vtbl() *IDeviceInformationPairingStatics2Vtbl {
	return (*IDeviceInformationPairingStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceInformationPairingStatics2) TryRegisterForAllInboundPairingRequestsWithProtectionLevel(pairingKindsSupported DevicePairingKinds, minProtectionLevel DevicePairingProtectionLevel) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryRegisterForAllInboundPairingRequestsWithProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(pairingKindsSupported), uintptr(minProtectionLevel), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C17F100E-3A46-4A78-8013-769DC9B97390
var IID_IDeviceInformationStatics = syscall.GUID{0xC17F100E, 0x3A46, 0x4A78,
	[8]byte{0x80, 0x13, 0x76, 0x9D, 0xC9, 0xB9, 0x73, 0x90}}

type IDeviceInformationStaticsInterface interface {
	win32.IInspectableInterface
	CreateFromIdAsync(deviceId string) *IAsyncOperation[*IDeviceInformation]
	CreateFromIdAsyncAdditionalProperties(deviceId string, additionalProperties *IIterable[string]) *IAsyncOperation[*IDeviceInformation]
	FindAllAsync() *IAsyncOperation[*IVectorView[*IDeviceInformation]]
	FindAllAsyncDeviceClass(deviceClass DeviceClass) *IAsyncOperation[*IVectorView[*IDeviceInformation]]
	FindAllAsyncAqsFilter(aqsFilter string) *IAsyncOperation[*IVectorView[*IDeviceInformation]]
	FindAllAsyncAqsFilterAndAdditionalProperties(aqsFilter string, additionalProperties *IIterable[string]) *IAsyncOperation[*IVectorView[*IDeviceInformation]]
	CreateWatcher() *IDeviceWatcher
	CreateWatcherDeviceClass(deviceClass DeviceClass) *IDeviceWatcher
	CreateWatcherAqsFilter(aqsFilter string) *IDeviceWatcher
	CreateWatcherAqsFilterAndAdditionalProperties(aqsFilter string, additionalProperties *IIterable[string]) *IDeviceWatcher
}

type IDeviceInformationStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateFromIdAsync                             uintptr
	CreateFromIdAsyncAdditionalProperties         uintptr
	FindAllAsync                                  uintptr
	FindAllAsyncDeviceClass                       uintptr
	FindAllAsyncAqsFilter                         uintptr
	FindAllAsyncAqsFilterAndAdditionalProperties  uintptr
	CreateWatcher                                 uintptr
	CreateWatcherDeviceClass                      uintptr
	CreateWatcherAqsFilter                        uintptr
	CreateWatcherAqsFilterAndAdditionalProperties uintptr
}

type IDeviceInformationStatics struct {
	win32.IInspectable
}

func (this *IDeviceInformationStatics) Vtbl() *IDeviceInformationStaticsVtbl {
	return (*IDeviceInformationStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceInformationStatics) CreateFromIdAsync(deviceId string) *IAsyncOperation[*IDeviceInformation] {
	var _result *IAsyncOperation[*IDeviceInformation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDeviceInformationStatics) CreateFromIdAsyncAdditionalProperties(deviceId string, additionalProperties *IIterable[string]) *IAsyncOperation[*IDeviceInformation] {
	var _result *IAsyncOperation[*IDeviceInformation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromIdAsyncAdditionalProperties, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(additionalProperties)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDeviceInformationStatics) FindAllAsync() *IAsyncOperation[*IVectorView[*IDeviceInformation]] {
	var _result *IAsyncOperation[*IVectorView[*IDeviceInformation]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDeviceInformationStatics) FindAllAsyncDeviceClass(deviceClass DeviceClass) *IAsyncOperation[*IVectorView[*IDeviceInformation]] {
	var _result *IAsyncOperation[*IVectorView[*IDeviceInformation]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllAsyncDeviceClass, uintptr(unsafe.Pointer(this)), uintptr(deviceClass), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDeviceInformationStatics) FindAllAsyncAqsFilter(aqsFilter string) *IAsyncOperation[*IVectorView[*IDeviceInformation]] {
	var _result *IAsyncOperation[*IVectorView[*IDeviceInformation]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllAsyncAqsFilter, uintptr(unsafe.Pointer(this)), NewHStr(aqsFilter).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDeviceInformationStatics) FindAllAsyncAqsFilterAndAdditionalProperties(aqsFilter string, additionalProperties *IIterable[string]) *IAsyncOperation[*IVectorView[*IDeviceInformation]] {
	var _result *IAsyncOperation[*IVectorView[*IDeviceInformation]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllAsyncAqsFilterAndAdditionalProperties, uintptr(unsafe.Pointer(this)), NewHStr(aqsFilter).Ptr, uintptr(unsafe.Pointer(additionalProperties)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDeviceInformationStatics) CreateWatcher() *IDeviceWatcher {
	var _result *IDeviceWatcher
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWatcher, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDeviceInformationStatics) CreateWatcherDeviceClass(deviceClass DeviceClass) *IDeviceWatcher {
	var _result *IDeviceWatcher
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWatcherDeviceClass, uintptr(unsafe.Pointer(this)), uintptr(deviceClass), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDeviceInformationStatics) CreateWatcherAqsFilter(aqsFilter string) *IDeviceWatcher {
	var _result *IDeviceWatcher
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWatcherAqsFilter, uintptr(unsafe.Pointer(this)), NewHStr(aqsFilter).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDeviceInformationStatics) CreateWatcherAqsFilterAndAdditionalProperties(aqsFilter string, additionalProperties *IIterable[string]) *IDeviceWatcher {
	var _result *IDeviceWatcher
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWatcherAqsFilterAndAdditionalProperties, uintptr(unsafe.Pointer(this)), NewHStr(aqsFilter).Ptr, uintptr(unsafe.Pointer(additionalProperties)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 493B4F34-A84F-45FD-9167-15D1CB1BD1F9
var IID_IDeviceInformationStatics2 = syscall.GUID{0x493B4F34, 0xA84F, 0x45FD,
	[8]byte{0x91, 0x67, 0x15, 0xD1, 0xCB, 0x1B, 0xD1, 0xF9}}

type IDeviceInformationStatics2Interface interface {
	win32.IInspectableInterface
	GetAqsFilterFromDeviceClass(deviceClass DeviceClass) string
	CreateFromIdAsyncWithKindAndAdditionalProperties(deviceId string, additionalProperties *IIterable[string], kind DeviceInformationKind) *IAsyncOperation[*IDeviceInformation]
	FindAllAsyncWithKindAqsFilterAndAdditionalProperties(aqsFilter string, additionalProperties *IIterable[string], kind DeviceInformationKind) *IAsyncOperation[*IVectorView[*IDeviceInformation]]
	CreateWatcherWithKindAqsFilterAndAdditionalProperties(aqsFilter string, additionalProperties *IIterable[string], kind DeviceInformationKind) *IDeviceWatcher
}

type IDeviceInformationStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetAqsFilterFromDeviceClass                           uintptr
	CreateFromIdAsyncWithKindAndAdditionalProperties      uintptr
	FindAllAsyncWithKindAqsFilterAndAdditionalProperties  uintptr
	CreateWatcherWithKindAqsFilterAndAdditionalProperties uintptr
}

type IDeviceInformationStatics2 struct {
	win32.IInspectable
}

func (this *IDeviceInformationStatics2) Vtbl() *IDeviceInformationStatics2Vtbl {
	return (*IDeviceInformationStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceInformationStatics2) GetAqsFilterFromDeviceClass(deviceClass DeviceClass) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAqsFilterFromDeviceClass, uintptr(unsafe.Pointer(this)), uintptr(deviceClass), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDeviceInformationStatics2) CreateFromIdAsyncWithKindAndAdditionalProperties(deviceId string, additionalProperties *IIterable[string], kind DeviceInformationKind) *IAsyncOperation[*IDeviceInformation] {
	var _result *IAsyncOperation[*IDeviceInformation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromIdAsyncWithKindAndAdditionalProperties, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(additionalProperties)), uintptr(kind), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDeviceInformationStatics2) FindAllAsyncWithKindAqsFilterAndAdditionalProperties(aqsFilter string, additionalProperties *IIterable[string], kind DeviceInformationKind) *IAsyncOperation[*IVectorView[*IDeviceInformation]] {
	var _result *IAsyncOperation[*IVectorView[*IDeviceInformation]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllAsyncWithKindAqsFilterAndAdditionalProperties, uintptr(unsafe.Pointer(this)), NewHStr(aqsFilter).Ptr, uintptr(unsafe.Pointer(additionalProperties)), uintptr(kind), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDeviceInformationStatics2) CreateWatcherWithKindAqsFilterAndAdditionalProperties(aqsFilter string, additionalProperties *IIterable[string], kind DeviceInformationKind) *IDeviceWatcher {
	var _result *IDeviceWatcher
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWatcherWithKindAqsFilterAndAdditionalProperties, uintptr(unsafe.Pointer(this)), NewHStr(aqsFilter).Ptr, uintptr(unsafe.Pointer(additionalProperties)), uintptr(kind), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8F315305-D972-44B7-A37E-9E822C78213B
var IID_IDeviceInformationUpdate = syscall.GUID{0x8F315305, 0xD972, 0x44B7,
	[8]byte{0xA3, 0x7E, 0x9E, 0x82, 0x2C, 0x78, 0x21, 0x3B}}

type IDeviceInformationUpdateInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	Get_Properties() *IMapView[string, interface{}]
}

type IDeviceInformationUpdateVtbl struct {
	win32.IInspectableVtbl
	Get_Id         uintptr
	Get_Properties uintptr
}

type IDeviceInformationUpdate struct {
	win32.IInspectable
}

func (this *IDeviceInformationUpdate) Vtbl() *IDeviceInformationUpdateVtbl {
	return (*IDeviceInformationUpdateVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceInformationUpdate) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDeviceInformationUpdate) Get_Properties() *IMapView[string, interface{}] {
	var _result *IMapView[string, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5D9D148C-A873-485E-BAA6-AA620788E3CC
var IID_IDeviceInformationUpdate2 = syscall.GUID{0x5D9D148C, 0xA873, 0x485E,
	[8]byte{0xBA, 0xA6, 0xAA, 0x62, 0x07, 0x88, 0xE3, 0xCC}}

type IDeviceInformationUpdate2Interface interface {
	win32.IInspectableInterface
	Get_Kind() DeviceInformationKind
}

type IDeviceInformationUpdate2Vtbl struct {
	win32.IInspectableVtbl
	Get_Kind uintptr
}

type IDeviceInformationUpdate2 struct {
	win32.IInspectable
}

func (this *IDeviceInformationUpdate2) Vtbl() *IDeviceInformationUpdate2Vtbl {
	return (*IDeviceInformationUpdate2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceInformationUpdate2) Get_Kind() DeviceInformationKind {
	var _result DeviceInformationKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Kind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F717FC56-DE6B-487F-8376-0180ACA69963
var IID_IDevicePairingRequestedEventArgs = syscall.GUID{0xF717FC56, 0xDE6B, 0x487F,
	[8]byte{0x83, 0x76, 0x01, 0x80, 0xAC, 0xA6, 0x99, 0x63}}

type IDevicePairingRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_DeviceInformation() *IDeviceInformation
	Get_PairingKind() DevicePairingKinds
	Get_Pin() string
	Accept()
	AcceptWithPin(pin string)
	GetDeferral() *IDeferral
}

type IDevicePairingRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceInformation uintptr
	Get_PairingKind       uintptr
	Get_Pin               uintptr
	Accept                uintptr
	AcceptWithPin         uintptr
	GetDeferral           uintptr
}

type IDevicePairingRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IDevicePairingRequestedEventArgs) Vtbl() *IDevicePairingRequestedEventArgsVtbl {
	return (*IDevicePairingRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDevicePairingRequestedEventArgs) Get_DeviceInformation() *IDeviceInformation {
	var _result *IDeviceInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDevicePairingRequestedEventArgs) Get_PairingKind() DevicePairingKinds {
	var _result DevicePairingKinds
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PairingKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDevicePairingRequestedEventArgs) Get_Pin() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Pin, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDevicePairingRequestedEventArgs) Accept() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Accept, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IDevicePairingRequestedEventArgs) AcceptWithPin(pin string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AcceptWithPin, uintptr(unsafe.Pointer(this)), NewHStr(pin).Ptr)
	_ = _hr
}

func (this *IDevicePairingRequestedEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C83752D9-E4D3-4DB0-A360-A105E437DBDC
var IID_IDevicePairingRequestedEventArgs2 = syscall.GUID{0xC83752D9, 0xE4D3, 0x4DB0,
	[8]byte{0xA3, 0x60, 0xA1, 0x05, 0xE4, 0x37, 0xDB, 0xDC}}

type IDevicePairingRequestedEventArgs2Interface interface {
	win32.IInspectableInterface
	AcceptWithPasswordCredential(passwordCredential *IPasswordCredential)
}

type IDevicePairingRequestedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	AcceptWithPasswordCredential uintptr
}

type IDevicePairingRequestedEventArgs2 struct {
	win32.IInspectable
}

func (this *IDevicePairingRequestedEventArgs2) Vtbl() *IDevicePairingRequestedEventArgs2Vtbl {
	return (*IDevicePairingRequestedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDevicePairingRequestedEventArgs2) AcceptWithPasswordCredential(passwordCredential *IPasswordCredential) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AcceptWithPasswordCredential, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(passwordCredential)))
	_ = _hr
}

// 072B02BF-DD95-4025-9B37-DE51ADBA37B7
var IID_IDevicePairingResult = syscall.GUID{0x072B02BF, 0xDD95, 0x4025,
	[8]byte{0x9B, 0x37, 0xDE, 0x51, 0xAD, 0xBA, 0x37, 0xB7}}

type IDevicePairingResultInterface interface {
	win32.IInspectableInterface
	Get_Status() DevicePairingResultStatus
	Get_ProtectionLevelUsed() DevicePairingProtectionLevel
}

type IDevicePairingResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status              uintptr
	Get_ProtectionLevelUsed uintptr
}

type IDevicePairingResult struct {
	win32.IInspectable
}

func (this *IDevicePairingResult) Vtbl() *IDevicePairingResultVtbl {
	return (*IDevicePairingResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDevicePairingResult) Get_Status() DevicePairingResultStatus {
	var _result DevicePairingResultStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDevicePairingResult) Get_ProtectionLevelUsed() DevicePairingProtectionLevel {
	var _result DevicePairingProtectionLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProtectionLevelUsed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 482CB27C-83BB-420E-BE51-6602B222DE54
var IID_IDevicePairingSettings = syscall.GUID{0x482CB27C, 0x83BB, 0x420E,
	[8]byte{0xBE, 0x51, 0x66, 0x02, 0xB2, 0x22, 0xDE, 0x54}}

type IDevicePairingSettingsInterface interface {
	win32.IInspectableInterface
}

type IDevicePairingSettingsVtbl struct {
	win32.IInspectableVtbl
}

type IDevicePairingSettings struct {
	win32.IInspectable
}

func (this *IDevicePairingSettings) Vtbl() *IDevicePairingSettingsVtbl {
	return (*IDevicePairingSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 84997AA2-034A-4440-8813-7D0BD479BF5A
var IID_IDevicePicker = syscall.GUID{0x84997AA2, 0x034A, 0x4440,
	[8]byte{0x88, 0x13, 0x7D, 0x0B, 0xD4, 0x79, 0xBF, 0x5A}}

type IDevicePickerInterface interface {
	win32.IInspectableInterface
	Get_Filter() *IDevicePickerFilter
	Get_Appearance() *IDevicePickerAppearance
	Get_RequestedProperties() *IVector[string]
	Add_DeviceSelected(handler TypedEventHandler[*IDevicePicker, *IDeviceSelectedEventArgs]) EventRegistrationToken
	Remove_DeviceSelected(token EventRegistrationToken)
	Add_DisconnectButtonClicked(handler TypedEventHandler[*IDevicePicker, *IDeviceDisconnectButtonClickedEventArgs]) EventRegistrationToken
	Remove_DisconnectButtonClicked(token EventRegistrationToken)
	Add_DevicePickerDismissed(handler TypedEventHandler[*IDevicePicker, interface{}]) EventRegistrationToken
	Remove_DevicePickerDismissed(token EventRegistrationToken)
	Show(selection Rect)
	ShowWithPlacement(selection Rect, placement Placement)
	PickSingleDeviceAsync(selection Rect) *IAsyncOperation[*IDeviceInformation]
	PickSingleDeviceAsyncWithPlacement(selection Rect, placement Placement) *IAsyncOperation[*IDeviceInformation]
	Hide()
	SetDisplayStatus(device *IDeviceInformation, status string, options DevicePickerDisplayStatusOptions)
}

type IDevicePickerVtbl struct {
	win32.IInspectableVtbl
	Get_Filter                         uintptr
	Get_Appearance                     uintptr
	Get_RequestedProperties            uintptr
	Add_DeviceSelected                 uintptr
	Remove_DeviceSelected              uintptr
	Add_DisconnectButtonClicked        uintptr
	Remove_DisconnectButtonClicked     uintptr
	Add_DevicePickerDismissed          uintptr
	Remove_DevicePickerDismissed       uintptr
	Show                               uintptr
	ShowWithPlacement                  uintptr
	PickSingleDeviceAsync              uintptr
	PickSingleDeviceAsyncWithPlacement uintptr
	Hide                               uintptr
	SetDisplayStatus                   uintptr
}

type IDevicePicker struct {
	win32.IInspectable
}

func (this *IDevicePicker) Vtbl() *IDevicePickerVtbl {
	return (*IDevicePickerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDevicePicker) Get_Filter() *IDevicePickerFilter {
	var _result *IDevicePickerFilter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Filter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDevicePicker) Get_Appearance() *IDevicePickerAppearance {
	var _result *IDevicePickerAppearance
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Appearance, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDevicePicker) Get_RequestedProperties() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestedProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDevicePicker) Add_DeviceSelected(handler TypedEventHandler[*IDevicePicker, *IDeviceSelectedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DeviceSelected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDevicePicker) Remove_DeviceSelected(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DeviceSelected, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IDevicePicker) Add_DisconnectButtonClicked(handler TypedEventHandler[*IDevicePicker, *IDeviceDisconnectButtonClickedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DisconnectButtonClicked, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDevicePicker) Remove_DisconnectButtonClicked(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DisconnectButtonClicked, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IDevicePicker) Add_DevicePickerDismissed(handler TypedEventHandler[*IDevicePicker, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DevicePickerDismissed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDevicePicker) Remove_DevicePickerDismissed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DevicePickerDismissed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IDevicePicker) Show(selection Rect) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Show, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&selection)))
	_ = _hr
}

func (this *IDevicePicker) ShowWithPlacement(selection Rect, placement Placement) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowWithPlacement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&selection)), uintptr(placement))
	_ = _hr
}

func (this *IDevicePicker) PickSingleDeviceAsync(selection Rect) *IAsyncOperation[*IDeviceInformation] {
	var _result *IAsyncOperation[*IDeviceInformation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PickSingleDeviceAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&selection)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDevicePicker) PickSingleDeviceAsyncWithPlacement(selection Rect, placement Placement) *IAsyncOperation[*IDeviceInformation] {
	var _result *IAsyncOperation[*IDeviceInformation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PickSingleDeviceAsyncWithPlacement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&selection)), uintptr(placement), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDevicePicker) Hide() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Hide, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IDevicePicker) SetDisplayStatus(device *IDeviceInformation, status string, options DevicePickerDisplayStatusOptions) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetDisplayStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(device)), NewHStr(status).Ptr, uintptr(options))
	_ = _hr
}

// E69A12C6-E627-4ED8-9B6C-460AF445E56D
var IID_IDevicePickerAppearance = syscall.GUID{0xE69A12C6, 0xE627, 0x4ED8,
	[8]byte{0x9B, 0x6C, 0x46, 0x0A, 0xF4, 0x45, 0xE5, 0x6D}}

type IDevicePickerAppearanceInterface interface {
	win32.IInspectableInterface
	Get_Title() string
	Put_Title(value string)
	Get_ForegroundColor() unsafe.Pointer
	Put_ForegroundColor(value unsafe.Pointer)
	Get_BackgroundColor() unsafe.Pointer
	Put_BackgroundColor(value unsafe.Pointer)
	Get_AccentColor() unsafe.Pointer
	Put_AccentColor(value unsafe.Pointer)
	Get_SelectedForegroundColor() unsafe.Pointer
	Put_SelectedForegroundColor(value unsafe.Pointer)
	Get_SelectedBackgroundColor() unsafe.Pointer
	Put_SelectedBackgroundColor(value unsafe.Pointer)
	Get_SelectedAccentColor() unsafe.Pointer
	Put_SelectedAccentColor(value unsafe.Pointer)
}

type IDevicePickerAppearanceVtbl struct {
	win32.IInspectableVtbl
	Get_Title                   uintptr
	Put_Title                   uintptr
	Get_ForegroundColor         uintptr
	Put_ForegroundColor         uintptr
	Get_BackgroundColor         uintptr
	Put_BackgroundColor         uintptr
	Get_AccentColor             uintptr
	Put_AccentColor             uintptr
	Get_SelectedForegroundColor uintptr
	Put_SelectedForegroundColor uintptr
	Get_SelectedBackgroundColor uintptr
	Put_SelectedBackgroundColor uintptr
	Get_SelectedAccentColor     uintptr
	Put_SelectedAccentColor     uintptr
}

type IDevicePickerAppearance struct {
	win32.IInspectable
}

func (this *IDevicePickerAppearance) Vtbl() *IDevicePickerAppearanceVtbl {
	return (*IDevicePickerAppearanceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDevicePickerAppearance) Get_Title() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Title, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDevicePickerAppearance) Put_Title(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Title, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IDevicePickerAppearance) Get_ForegroundColor() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ForegroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDevicePickerAppearance) Put_ForegroundColor(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ForegroundColor, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDevicePickerAppearance) Get_BackgroundColor() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BackgroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDevicePickerAppearance) Put_BackgroundColor(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BackgroundColor, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDevicePickerAppearance) Get_AccentColor() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AccentColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDevicePickerAppearance) Put_AccentColor(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AccentColor, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDevicePickerAppearance) Get_SelectedForegroundColor() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SelectedForegroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDevicePickerAppearance) Put_SelectedForegroundColor(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SelectedForegroundColor, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDevicePickerAppearance) Get_SelectedBackgroundColor() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SelectedBackgroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDevicePickerAppearance) Put_SelectedBackgroundColor(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SelectedBackgroundColor, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDevicePickerAppearance) Get_SelectedAccentColor() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SelectedAccentColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDevicePickerAppearance) Put_SelectedAccentColor(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SelectedAccentColor, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 91DB92A2-57CB-48F1-9B59-A59B7A1F02A2
var IID_IDevicePickerFilter = syscall.GUID{0x91DB92A2, 0x57CB, 0x48F1,
	[8]byte{0x9B, 0x59, 0xA5, 0x9B, 0x7A, 0x1F, 0x02, 0xA2}}

type IDevicePickerFilterInterface interface {
	win32.IInspectableInterface
	Get_SupportedDeviceClasses() *IVector[DeviceClass]
	Get_SupportedDeviceSelectors() *IVector[string]
}

type IDevicePickerFilterVtbl struct {
	win32.IInspectableVtbl
	Get_SupportedDeviceClasses   uintptr
	Get_SupportedDeviceSelectors uintptr
}

type IDevicePickerFilter struct {
	win32.IInspectable
}

func (this *IDevicePickerFilter) Vtbl() *IDevicePickerFilterVtbl {
	return (*IDevicePickerFilterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDevicePickerFilter) Get_SupportedDeviceClasses() *IVector[DeviceClass] {
	var _result *IVector[DeviceClass]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedDeviceClasses, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDevicePickerFilter) Get_SupportedDeviceSelectors() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedDeviceSelectors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 269EDADE-1D2F-4940-8402-4156B81D3C77
var IID_IDeviceSelectedEventArgs = syscall.GUID{0x269EDADE, 0x1D2F, 0x4940,
	[8]byte{0x84, 0x02, 0x41, 0x56, 0xB8, 0x1D, 0x3C, 0x77}}

type IDeviceSelectedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_SelectedDevice() *IDeviceInformation
}

type IDeviceSelectedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_SelectedDevice uintptr
}

type IDeviceSelectedEventArgs struct {
	win32.IInspectable
}

func (this *IDeviceSelectedEventArgs) Vtbl() *IDeviceSelectedEventArgsVtbl {
	return (*IDeviceSelectedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceSelectedEventArgs) Get_SelectedDevice() *IDeviceInformation {
	var _result *IDeviceInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SelectedDevice, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 66F44AD3-79D9-444B-92CF-A92EF72571C7
var IID_IDeviceUnpairingResult = syscall.GUID{0x66F44AD3, 0x79D9, 0x444B,
	[8]byte{0x92, 0xCF, 0xA9, 0x2E, 0xF7, 0x25, 0x71, 0xC7}}

type IDeviceUnpairingResultInterface interface {
	win32.IInspectableInterface
	Get_Status() DeviceUnpairingResultStatus
}

type IDeviceUnpairingResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
}

type IDeviceUnpairingResult struct {
	win32.IInspectable
}

func (this *IDeviceUnpairingResult) Vtbl() *IDeviceUnpairingResultVtbl {
	return (*IDeviceUnpairingResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceUnpairingResult) Get_Status() DeviceUnpairingResultStatus {
	var _result DeviceUnpairingResultStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C9EAB97D-8F6B-4F96-A9F4-ABC814E22271
var IID_IDeviceWatcher = syscall.GUID{0xC9EAB97D, 0x8F6B, 0x4F96,
	[8]byte{0xA9, 0xF4, 0xAB, 0xC8, 0x14, 0xE2, 0x22, 0x71}}

type IDeviceWatcherInterface interface {
	win32.IInspectableInterface
	Add_Added(handler TypedEventHandler[*IDeviceWatcher, *IDeviceInformation]) EventRegistrationToken
	Remove_Added(token EventRegistrationToken)
	Add_Updated(handler TypedEventHandler[*IDeviceWatcher, *IDeviceInformationUpdate]) EventRegistrationToken
	Remove_Updated(token EventRegistrationToken)
	Add_Removed(handler TypedEventHandler[*IDeviceWatcher, *IDeviceInformationUpdate]) EventRegistrationToken
	Remove_Removed(token EventRegistrationToken)
	Add_EnumerationCompleted(handler TypedEventHandler[*IDeviceWatcher, interface{}]) EventRegistrationToken
	Remove_EnumerationCompleted(token EventRegistrationToken)
	Add_Stopped(handler TypedEventHandler[*IDeviceWatcher, interface{}]) EventRegistrationToken
	Remove_Stopped(token EventRegistrationToken)
	Get_Status() DeviceWatcherStatus
	Start()
	Stop()
}

type IDeviceWatcherVtbl struct {
	win32.IInspectableVtbl
	Add_Added                   uintptr
	Remove_Added                uintptr
	Add_Updated                 uintptr
	Remove_Updated              uintptr
	Add_Removed                 uintptr
	Remove_Removed              uintptr
	Add_EnumerationCompleted    uintptr
	Remove_EnumerationCompleted uintptr
	Add_Stopped                 uintptr
	Remove_Stopped              uintptr
	Get_Status                  uintptr
	Start                       uintptr
	Stop                        uintptr
}

type IDeviceWatcher struct {
	win32.IInspectable
}

func (this *IDeviceWatcher) Vtbl() *IDeviceWatcherVtbl {
	return (*IDeviceWatcherVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceWatcher) Add_Added(handler TypedEventHandler[*IDeviceWatcher, *IDeviceInformation]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Added, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDeviceWatcher) Remove_Added(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Added, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IDeviceWatcher) Add_Updated(handler TypedEventHandler[*IDeviceWatcher, *IDeviceInformationUpdate]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Updated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDeviceWatcher) Remove_Updated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Updated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IDeviceWatcher) Add_Removed(handler TypedEventHandler[*IDeviceWatcher, *IDeviceInformationUpdate]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Removed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDeviceWatcher) Remove_Removed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Removed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IDeviceWatcher) Add_EnumerationCompleted(handler TypedEventHandler[*IDeviceWatcher, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_EnumerationCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDeviceWatcher) Remove_EnumerationCompleted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_EnumerationCompleted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IDeviceWatcher) Add_Stopped(handler TypedEventHandler[*IDeviceWatcher, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Stopped, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDeviceWatcher) Remove_Stopped(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Stopped, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IDeviceWatcher) Get_Status() DeviceWatcherStatus {
	var _result DeviceWatcherStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDeviceWatcher) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IDeviceWatcher) Stop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// FF08456E-ED14-49E9-9A69-8117C54AE971
var IID_IDeviceWatcher2 = syscall.GUID{0xFF08456E, 0xED14, 0x49E9,
	[8]byte{0x9A, 0x69, 0x81, 0x17, 0xC5, 0x4A, 0xE9, 0x71}}

type IDeviceWatcher2Interface interface {
	win32.IInspectableInterface
	GetBackgroundTrigger(requestedEventKinds *IIterable[DeviceWatcherEventKind]) *IDeviceWatcherTrigger
}

type IDeviceWatcher2Vtbl struct {
	win32.IInspectableVtbl
	GetBackgroundTrigger uintptr
}

type IDeviceWatcher2 struct {
	win32.IInspectable
}

func (this *IDeviceWatcher2) Vtbl() *IDeviceWatcher2Vtbl {
	return (*IDeviceWatcher2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceWatcher2) GetBackgroundTrigger(requestedEventKinds *IIterable[DeviceWatcherEventKind]) *IDeviceWatcherTrigger {
	var _result *IDeviceWatcherTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetBackgroundTrigger, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(requestedEventKinds)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 74AA9C0B-1DBD-47FD-B635-3CC556D0FF8B
var IID_IDeviceWatcherEvent = syscall.GUID{0x74AA9C0B, 0x1DBD, 0x47FD,
	[8]byte{0xB6, 0x35, 0x3C, 0xC5, 0x56, 0xD0, 0xFF, 0x8B}}

type IDeviceWatcherEventInterface interface {
	win32.IInspectableInterface
	Get_Kind() DeviceWatcherEventKind
	Get_DeviceInformation() *IDeviceInformation
	Get_DeviceInformationUpdate() *IDeviceInformationUpdate
}

type IDeviceWatcherEventVtbl struct {
	win32.IInspectableVtbl
	Get_Kind                    uintptr
	Get_DeviceInformation       uintptr
	Get_DeviceInformationUpdate uintptr
}

type IDeviceWatcherEvent struct {
	win32.IInspectable
}

func (this *IDeviceWatcherEvent) Vtbl() *IDeviceWatcherEventVtbl {
	return (*IDeviceWatcherEventVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceWatcherEvent) Get_Kind() DeviceWatcherEventKind {
	var _result DeviceWatcherEventKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Kind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDeviceWatcherEvent) Get_DeviceInformation() *IDeviceInformation {
	var _result *IDeviceInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDeviceWatcherEvent) Get_DeviceInformationUpdate() *IDeviceInformationUpdate {
	var _result *IDeviceInformationUpdate
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceInformationUpdate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 38808119-4CB7-4E57-A56D-776D07CBFEF9
var IID_IDeviceWatcherTriggerDetails = syscall.GUID{0x38808119, 0x4CB7, 0x4E57,
	[8]byte{0xA5, 0x6D, 0x77, 0x6D, 0x07, 0xCB, 0xFE, 0xF9}}

type IDeviceWatcherTriggerDetailsInterface interface {
	win32.IInspectableInterface
	Get_DeviceWatcherEvents() *IVectorView[*IDeviceWatcherEvent]
}

type IDeviceWatcherTriggerDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceWatcherEvents uintptr
}

type IDeviceWatcherTriggerDetails struct {
	win32.IInspectable
}

func (this *IDeviceWatcherTriggerDetails) Vtbl() *IDeviceWatcherTriggerDetailsVtbl {
	return (*IDeviceWatcherTriggerDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeviceWatcherTriggerDetails) Get_DeviceWatcherEvents() *IVectorView[*IDeviceWatcherEvent] {
	var _result *IVectorView[*IDeviceWatcherEvent]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceWatcherEvents, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 42340A27-5810-459C-AABB-C65E1F813ECF
var IID_IEnclosureLocation = syscall.GUID{0x42340A27, 0x5810, 0x459C,
	[8]byte{0xAA, 0xBB, 0xC6, 0x5E, 0x1F, 0x81, 0x3E, 0xCF}}

type IEnclosureLocationInterface interface {
	win32.IInspectableInterface
	Get_InDock() bool
	Get_InLid() bool
	Get_Panel() Panel
}

type IEnclosureLocationVtbl struct {
	win32.IInspectableVtbl
	Get_InDock uintptr
	Get_InLid  uintptr
	Get_Panel  uintptr
}

type IEnclosureLocation struct {
	win32.IInspectable
}

func (this *IEnclosureLocation) Vtbl() *IEnclosureLocationVtbl {
	return (*IEnclosureLocationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IEnclosureLocation) Get_InDock() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InDock, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IEnclosureLocation) Get_InLid() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InLid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IEnclosureLocation) Get_Panel() Panel {
	var _result Panel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Panel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 2885995B-E07D-485D-8A9E-BDF29AEF4F66
var IID_IEnclosureLocation2 = syscall.GUID{0x2885995B, 0xE07D, 0x485D,
	[8]byte{0x8A, 0x9E, 0xBD, 0xF2, 0x9A, 0xEF, 0x4F, 0x66}}

type IEnclosureLocation2Interface interface {
	win32.IInspectableInterface
	Get_RotationAngleInDegreesClockwise() uint32
}

type IEnclosureLocation2Vtbl struct {
	win32.IInspectableVtbl
	Get_RotationAngleInDegreesClockwise uintptr
}

type IEnclosureLocation2 struct {
	win32.IInspectable
}

func (this *IEnclosureLocation2) Vtbl() *IEnclosureLocation2Vtbl {
	return (*IEnclosureLocation2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IEnclosureLocation2) Get_RotationAngleInDegreesClockwise() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RotationAngleInDegreesClockwise, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type DeviceAccessChangedEventArgs struct {
	RtClass
	*IDeviceAccessChangedEventArgs
}

type DeviceDisconnectButtonClickedEventArgs struct {
	RtClass
	*IDeviceDisconnectButtonClickedEventArgs
}

type DeviceInformation struct {
	RtClass
	*IDeviceInformation
}

func NewIDeviceInformationStatics() *IDeviceInformationStatics {
	var p *IDeviceInformationStatics
	hs := NewHStr("Windows.Devices.Enumeration.DeviceInformation")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IDeviceInformationStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIDeviceInformationStatics2() *IDeviceInformationStatics2 {
	var p *IDeviceInformationStatics2
	hs := NewHStr("Windows.Devices.Enumeration.DeviceInformation")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IDeviceInformationStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type DeviceInformationCollection struct {
	RtClass
	*IVectorView[*IDeviceInformation]
}

type DeviceInformationUpdate struct {
	RtClass
	*IDeviceInformationUpdate
}

type DevicePicker struct {
	RtClass
	*IDevicePicker
}

func NewDevicePicker() *DevicePicker {
	hs := NewHStr("Windows.Devices.Enumeration.DevicePicker")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &DevicePicker{
		RtClass:       RtClass{PInspect: p},
		IDevicePicker: (*IDevicePicker)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type DevicePickerAppearance struct {
	RtClass
	*IDevicePickerAppearance
}

type DevicePickerFilter struct {
	RtClass
	*IDevicePickerFilter
}

type DeviceSelectedEventArgs struct {
	RtClass
	*IDeviceSelectedEventArgs
}

type DeviceThumbnail struct {
	RtClass
	*IRandomAccessStreamWithContentType
}

type DeviceWatcher struct {
	RtClass
	*IDeviceWatcher
}

type EnclosureLocation struct {
	RtClass
	*IEnclosureLocation
}
