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
type DataClasses uint32

const (
	DataClasses_None                  DataClasses = 0
	DataClasses_Gprs                  DataClasses = 1
	DataClasses_Edge                  DataClasses = 2
	DataClasses_Umts                  DataClasses = 4
	DataClasses_Hsdpa                 DataClasses = 8
	DataClasses_Hsupa                 DataClasses = 16
	DataClasses_LteAdvanced           DataClasses = 32
	DataClasses_NewRadioNonStandalone DataClasses = 64
	DataClasses_NewRadioStandalone    DataClasses = 128
	DataClasses_Cdma1xRtt             DataClasses = 65536
	DataClasses_Cdma1xEvdo            DataClasses = 131072
	DataClasses_Cdma1xEvdoRevA        DataClasses = 262144
	DataClasses_Cdma1xEvdv            DataClasses = 524288
	DataClasses_Cdma3xRtt             DataClasses = 1048576
	DataClasses_Cdma1xEvdoRevB        DataClasses = 2097152
	DataClasses_CdmaUmb               DataClasses = 4194304
	DataClasses_Custom                DataClasses = 2147483648
)

// enum
type ESimAuthenticationPreference int32

const (
	ESimAuthenticationPreference_OnEntry  ESimAuthenticationPreference = 0
	ESimAuthenticationPreference_OnAction ESimAuthenticationPreference = 1
	ESimAuthenticationPreference_Never    ESimAuthenticationPreference = 2
)

// enum
type ESimDiscoverResultKind int32

const (
	ESimDiscoverResultKind_None            ESimDiscoverResultKind = 0
	ESimDiscoverResultKind_Events          ESimDiscoverResultKind = 1
	ESimDiscoverResultKind_ProfileMetadata ESimDiscoverResultKind = 2
)

// enum
type ESimOperationStatus int32

const (
	ESimOperationStatus_Success                            ESimOperationStatus = 0
	ESimOperationStatus_NotAuthorized                      ESimOperationStatus = 1
	ESimOperationStatus_NotFound                           ESimOperationStatus = 2
	ESimOperationStatus_PolicyViolation                    ESimOperationStatus = 3
	ESimOperationStatus_InsufficientSpaceOnCard            ESimOperationStatus = 4
	ESimOperationStatus_ServerFailure                      ESimOperationStatus = 5
	ESimOperationStatus_ServerNotReachable                 ESimOperationStatus = 6
	ESimOperationStatus_TimeoutWaitingForUserConsent       ESimOperationStatus = 7
	ESimOperationStatus_IncorrectConfirmationCode          ESimOperationStatus = 8
	ESimOperationStatus_ConfirmationCodeMaxRetriesExceeded ESimOperationStatus = 9
	ESimOperationStatus_CardRemoved                        ESimOperationStatus = 10
	ESimOperationStatus_CardBusy                           ESimOperationStatus = 11
	ESimOperationStatus_Other                              ESimOperationStatus = 12
	ESimOperationStatus_CardGeneralFailure                 ESimOperationStatus = 13
	ESimOperationStatus_ConfirmationCodeMissing            ESimOperationStatus = 14
	ESimOperationStatus_InvalidMatchingId                  ESimOperationStatus = 15
	ESimOperationStatus_NoEligibleProfileForThisDevice     ESimOperationStatus = 16
	ESimOperationStatus_OperationAborted                   ESimOperationStatus = 17
	ESimOperationStatus_EidMismatch                        ESimOperationStatus = 18
	ESimOperationStatus_ProfileNotAvailableForNewBinding   ESimOperationStatus = 19
	ESimOperationStatus_ProfileNotReleasedByOperator       ESimOperationStatus = 20
	ESimOperationStatus_OperationProhibitedByProfileClass  ESimOperationStatus = 21
	ESimOperationStatus_ProfileNotPresent                  ESimOperationStatus = 22
	ESimOperationStatus_NoCorrespondingRequest             ESimOperationStatus = 23
	ESimOperationStatus_TimeoutWaitingForResponse          ESimOperationStatus = 24
	ESimOperationStatus_IccidAlreadyExists                 ESimOperationStatus = 25
	ESimOperationStatus_ProfileProcessingError             ESimOperationStatus = 26
	ESimOperationStatus_ServerNotTrusted                   ESimOperationStatus = 27
	ESimOperationStatus_ProfileDownloadMaxRetriesExceeded  ESimOperationStatus = 28
)

// enum
type ESimProfileClass int32

const (
	ESimProfileClass_Operational  ESimProfileClass = 0
	ESimProfileClass_Test         ESimProfileClass = 1
	ESimProfileClass_Provisioning ESimProfileClass = 2
)

// enum
type ESimProfileMetadataState int32

const (
	ESimProfileMetadataState_Unknown           ESimProfileMetadataState = 0
	ESimProfileMetadataState_WaitingForInstall ESimProfileMetadataState = 1
	ESimProfileMetadataState_Downloading       ESimProfileMetadataState = 2
	ESimProfileMetadataState_Installing        ESimProfileMetadataState = 3
	ESimProfileMetadataState_Expired           ESimProfileMetadataState = 4
	ESimProfileMetadataState_RejectingDownload ESimProfileMetadataState = 5
	ESimProfileMetadataState_NoLongerAvailable ESimProfileMetadataState = 6
	ESimProfileMetadataState_DeniedByPolicy    ESimProfileMetadataState = 7
)

// enum
type ESimProfileState int32

const (
	ESimProfileState_Unknown  ESimProfileState = 0
	ESimProfileState_Disabled ESimProfileState = 1
	ESimProfileState_Enabled  ESimProfileState = 2
	ESimProfileState_Deleted  ESimProfileState = 3
)

// enum
type ESimState int32

const (
	ESimState_Unknown ESimState = 0
	ESimState_Idle    ESimState = 1
	ESimState_Removed ESimState = 2
	ESimState_Busy    ESimState = 3
)

// enum
type ESimWatcherStatus int32

const (
	ESimWatcherStatus_Created              ESimWatcherStatus = 0
	ESimWatcherStatus_Started              ESimWatcherStatus = 1
	ESimWatcherStatus_EnumerationCompleted ESimWatcherStatus = 2
	ESimWatcherStatus_Stopping             ESimWatcherStatus = 3
	ESimWatcherStatus_Stopped              ESimWatcherStatus = 4
)

// enum
type HotspotAuthenticationResponseCode int32

const (
	HotspotAuthenticationResponseCode_NoError                    HotspotAuthenticationResponseCode = 0
	HotspotAuthenticationResponseCode_LoginSucceeded             HotspotAuthenticationResponseCode = 50
	HotspotAuthenticationResponseCode_LoginFailed                HotspotAuthenticationResponseCode = 100
	HotspotAuthenticationResponseCode_RadiusServerError          HotspotAuthenticationResponseCode = 102
	HotspotAuthenticationResponseCode_NetworkAdministratorError  HotspotAuthenticationResponseCode = 105
	HotspotAuthenticationResponseCode_LoginAborted               HotspotAuthenticationResponseCode = 151
	HotspotAuthenticationResponseCode_AccessGatewayInternalError HotspotAuthenticationResponseCode = 255
)

// enum
type MobileBroadbandAccountWatcherStatus int32

const (
	MobileBroadbandAccountWatcherStatus_Created              MobileBroadbandAccountWatcherStatus = 0
	MobileBroadbandAccountWatcherStatus_Started              MobileBroadbandAccountWatcherStatus = 1
	MobileBroadbandAccountWatcherStatus_EnumerationCompleted MobileBroadbandAccountWatcherStatus = 2
	MobileBroadbandAccountWatcherStatus_Stopped              MobileBroadbandAccountWatcherStatus = 3
	MobileBroadbandAccountWatcherStatus_Aborted              MobileBroadbandAccountWatcherStatus = 4
)

// enum
type MobileBroadbandDeviceType int32

const (
	MobileBroadbandDeviceType_Unknown   MobileBroadbandDeviceType = 0
	MobileBroadbandDeviceType_Embedded  MobileBroadbandDeviceType = 1
	MobileBroadbandDeviceType_Removable MobileBroadbandDeviceType = 2
	MobileBroadbandDeviceType_Remote    MobileBroadbandDeviceType = 3
)

// enum
type MobileBroadbandModemStatus int32

const (
	MobileBroadbandModemStatus_Success         MobileBroadbandModemStatus = 0
	MobileBroadbandModemStatus_OtherFailure    MobileBroadbandModemStatus = 1
	MobileBroadbandModemStatus_Busy            MobileBroadbandModemStatus = 2
	MobileBroadbandModemStatus_NoDeviceSupport MobileBroadbandModemStatus = 3
)

// enum
type MobileBroadbandPinFormat int32

const (
	MobileBroadbandPinFormat_Unknown      MobileBroadbandPinFormat = 0
	MobileBroadbandPinFormat_Numeric      MobileBroadbandPinFormat = 1
	MobileBroadbandPinFormat_Alphanumeric MobileBroadbandPinFormat = 2
)

// enum
type MobileBroadbandPinLockState int32

const (
	MobileBroadbandPinLockState_Unknown               MobileBroadbandPinLockState = 0
	MobileBroadbandPinLockState_Unlocked              MobileBroadbandPinLockState = 1
	MobileBroadbandPinLockState_PinRequired           MobileBroadbandPinLockState = 2
	MobileBroadbandPinLockState_PinUnblockKeyRequired MobileBroadbandPinLockState = 3
)

// enum
type MobileBroadbandPinType int32

const (
	MobileBroadbandPinType_None               MobileBroadbandPinType = 0
	MobileBroadbandPinType_Custom             MobileBroadbandPinType = 1
	MobileBroadbandPinType_Pin1               MobileBroadbandPinType = 2
	MobileBroadbandPinType_Pin2               MobileBroadbandPinType = 3
	MobileBroadbandPinType_SimPin             MobileBroadbandPinType = 4
	MobileBroadbandPinType_FirstSimPin        MobileBroadbandPinType = 5
	MobileBroadbandPinType_NetworkPin         MobileBroadbandPinType = 6
	MobileBroadbandPinType_NetworkSubsetPin   MobileBroadbandPinType = 7
	MobileBroadbandPinType_ServiceProviderPin MobileBroadbandPinType = 8
	MobileBroadbandPinType_CorporatePin       MobileBroadbandPinType = 9
	MobileBroadbandPinType_SubsidyLock        MobileBroadbandPinType = 10
)

// enum
type MobileBroadbandRadioState int32

const (
	MobileBroadbandRadioState_Off MobileBroadbandRadioState = 0
	MobileBroadbandRadioState_On  MobileBroadbandRadioState = 1
)

// enum
type MobileBroadbandSlotState int32

const (
	MobileBroadbandSlotState_Unmanaged           MobileBroadbandSlotState = 0
	MobileBroadbandSlotState_Unknown             MobileBroadbandSlotState = 1
	MobileBroadbandSlotState_OffEmpty            MobileBroadbandSlotState = 2
	MobileBroadbandSlotState_Off                 MobileBroadbandSlotState = 3
	MobileBroadbandSlotState_Empty               MobileBroadbandSlotState = 4
	MobileBroadbandSlotState_NotReady            MobileBroadbandSlotState = 5
	MobileBroadbandSlotState_Active              MobileBroadbandSlotState = 6
	MobileBroadbandSlotState_Error               MobileBroadbandSlotState = 7
	MobileBroadbandSlotState_ActiveEsim          MobileBroadbandSlotState = 8
	MobileBroadbandSlotState_ActiveEsimNoProfile MobileBroadbandSlotState = 9
)

// enum
type MobileBroadbandUiccAppOperationStatus int32

const (
	MobileBroadbandUiccAppOperationStatus_Success                MobileBroadbandUiccAppOperationStatus = 0
	MobileBroadbandUiccAppOperationStatus_InvalidUiccFilePath    MobileBroadbandUiccAppOperationStatus = 1
	MobileBroadbandUiccAppOperationStatus_AccessConditionNotHeld MobileBroadbandUiccAppOperationStatus = 2
	MobileBroadbandUiccAppOperationStatus_UiccBusy               MobileBroadbandUiccAppOperationStatus = 3
)

// enum
type NetworkDeviceStatus int32

const (
	NetworkDeviceStatus_DeviceNotReady        NetworkDeviceStatus = 0
	NetworkDeviceStatus_DeviceReady           NetworkDeviceStatus = 1
	NetworkDeviceStatus_SimNotInserted        NetworkDeviceStatus = 2
	NetworkDeviceStatus_BadSim                NetworkDeviceStatus = 3
	NetworkDeviceStatus_DeviceHardwareFailure NetworkDeviceStatus = 4
	NetworkDeviceStatus_AccountNotActivated   NetworkDeviceStatus = 5
	NetworkDeviceStatus_DeviceLocked          NetworkDeviceStatus = 6
	NetworkDeviceStatus_DeviceBlocked         NetworkDeviceStatus = 7
)

// enum
type NetworkOperatorDataUsageNotificationKind int32

const (
	NetworkOperatorDataUsageNotificationKind_DataUsageProgress NetworkOperatorDataUsageNotificationKind = 0
)

// enum
type NetworkOperatorEventMessageType int32

const (
	NetworkOperatorEventMessageType_Gsm                              NetworkOperatorEventMessageType = 0
	NetworkOperatorEventMessageType_Cdma                             NetworkOperatorEventMessageType = 1
	NetworkOperatorEventMessageType_Ussd                             NetworkOperatorEventMessageType = 2
	NetworkOperatorEventMessageType_DataPlanThresholdReached         NetworkOperatorEventMessageType = 3
	NetworkOperatorEventMessageType_DataPlanReset                    NetworkOperatorEventMessageType = 4
	NetworkOperatorEventMessageType_DataPlanDeleted                  NetworkOperatorEventMessageType = 5
	NetworkOperatorEventMessageType_ProfileConnected                 NetworkOperatorEventMessageType = 6
	NetworkOperatorEventMessageType_ProfileDisconnected              NetworkOperatorEventMessageType = 7
	NetworkOperatorEventMessageType_RegisteredRoaming                NetworkOperatorEventMessageType = 8
	NetworkOperatorEventMessageType_RegisteredHome                   NetworkOperatorEventMessageType = 9
	NetworkOperatorEventMessageType_TetheringEntitlementCheck        NetworkOperatorEventMessageType = 10
	NetworkOperatorEventMessageType_TetheringOperationalStateChanged NetworkOperatorEventMessageType = 11
	NetworkOperatorEventMessageType_TetheringNumberOfClientsChanged  NetworkOperatorEventMessageType = 12
)

// enum
type NetworkRegistrationState int32

const (
	NetworkRegistrationState_None         NetworkRegistrationState = 0
	NetworkRegistrationState_Deregistered NetworkRegistrationState = 1
	NetworkRegistrationState_Searching    NetworkRegistrationState = 2
	NetworkRegistrationState_Home         NetworkRegistrationState = 3
	NetworkRegistrationState_Roaming      NetworkRegistrationState = 4
	NetworkRegistrationState_Partner      NetworkRegistrationState = 5
	NetworkRegistrationState_Denied       NetworkRegistrationState = 6
)

// enum
type ProfileMediaType int32

const (
	ProfileMediaType_Wlan ProfileMediaType = 0
	ProfileMediaType_Wwan ProfileMediaType = 1
)

// enum
type TetheringCapability int32

const (
	TetheringCapability_Enabled                           TetheringCapability = 0
	TetheringCapability_DisabledByGroupPolicy             TetheringCapability = 1
	TetheringCapability_DisabledByHardwareLimitation      TetheringCapability = 2
	TetheringCapability_DisabledByOperator                TetheringCapability = 3
	TetheringCapability_DisabledBySku                     TetheringCapability = 4
	TetheringCapability_DisabledByRequiredAppNotInstalled TetheringCapability = 5
	TetheringCapability_DisabledDueToUnknownCause         TetheringCapability = 6
	TetheringCapability_DisabledBySystemCapability        TetheringCapability = 7
)

// enum
type TetheringOperationStatus int32

const (
	TetheringOperationStatus_Success                    TetheringOperationStatus = 0
	TetheringOperationStatus_Unknown                    TetheringOperationStatus = 1
	TetheringOperationStatus_MobileBroadbandDeviceOff   TetheringOperationStatus = 2
	TetheringOperationStatus_WiFiDeviceOff              TetheringOperationStatus = 3
	TetheringOperationStatus_EntitlementCheckTimeout    TetheringOperationStatus = 4
	TetheringOperationStatus_EntitlementCheckFailure    TetheringOperationStatus = 5
	TetheringOperationStatus_OperationInProgress        TetheringOperationStatus = 6
	TetheringOperationStatus_BluetoothDeviceOff         TetheringOperationStatus = 7
	TetheringOperationStatus_NetworkLimitedConnectivity TetheringOperationStatus = 8
)

// enum
type TetheringOperationalState int32

const (
	TetheringOperationalState_Unknown      TetheringOperationalState = 0
	TetheringOperationalState_On           TetheringOperationalState = 1
	TetheringOperationalState_Off          TetheringOperationalState = 2
	TetheringOperationalState_InTransition TetheringOperationalState = 3
)

// enum
type TetheringWiFiBand int32

const (
	TetheringWiFiBand_Auto                  TetheringWiFiBand = 0
	TetheringWiFiBand_TwoPointFourGigahertz TetheringWiFiBand = 1
	TetheringWiFiBand_FiveGigahertz         TetheringWiFiBand = 2
)

// enum
type UiccAccessCondition int32

const (
	UiccAccessCondition_AlwaysAllowed   UiccAccessCondition = 0
	UiccAccessCondition_Pin1            UiccAccessCondition = 1
	UiccAccessCondition_Pin2            UiccAccessCondition = 2
	UiccAccessCondition_Pin3            UiccAccessCondition = 3
	UiccAccessCondition_Pin4            UiccAccessCondition = 4
	UiccAccessCondition_Administrative5 UiccAccessCondition = 5
	UiccAccessCondition_Administrative6 UiccAccessCondition = 6
	UiccAccessCondition_NeverAllowed    UiccAccessCondition = 7
)

// enum
type UiccAppKind int32

const (
	UiccAppKind_Unknown UiccAppKind = 0
	UiccAppKind_MF      UiccAppKind = 1
	UiccAppKind_MFSim   UiccAppKind = 2
	UiccAppKind_MFRuim  UiccAppKind = 3
	UiccAppKind_USim    UiccAppKind = 4
	UiccAppKind_CSim    UiccAppKind = 5
	UiccAppKind_ISim    UiccAppKind = 6
)

// enum
type UiccAppRecordKind int32

const (
	UiccAppRecordKind_Unknown        UiccAppRecordKind = 0
	UiccAppRecordKind_Transparent    UiccAppRecordKind = 1
	UiccAppRecordKind_RecordOriented UiccAppRecordKind = 2
)

// enum
type UssdResultCode int32

const (
	UssdResultCode_NoActionRequired      UssdResultCode = 0
	UssdResultCode_ActionRequired        UssdResultCode = 1
	UssdResultCode_Terminated            UssdResultCode = 2
	UssdResultCode_OtherLocalClient      UssdResultCode = 3
	UssdResultCode_OperationNotSupported UssdResultCode = 4
	UssdResultCode_NetworkTimeout        UssdResultCode = 5
)

// structs

type ESimProfileInstallProgress struct {
	TotalSizeInBytes     int32
	InstalledSizeInBytes int32
}

type LegacyNetworkOperatorsContract struct {
}

type NetworkOperatorsFdnContract struct {
}

type ProfileUsage struct {
	UsageInMegabytes uint32
	LastSyncTime     DateTime
}

// interfaces

// 6F6E6E26-F123-437D-8CED-DC1D2BC0C3A9
var IID_IESim = syscall.GUID{0x6F6E6E26, 0xF123, 0x437D,
	[8]byte{0x8C, 0xED, 0xDC, 0x1D, 0x2B, 0xC0, 0xC3, 0xA9}}

type IESimInterface interface {
	win32.IInspectableInterface
	Get_AvailableMemoryInBytes() *IReference[int32]
	Get_Eid() string
	Get_FirmwareVersion() string
	Get_MobileBroadbandModemDeviceId() string
	Get_Policy() *IESimPolicy
	Get_State() ESimState
	GetProfiles() *IVectorView[*IESimProfile]
	DeleteProfileAsync(profileId string) *IAsyncOperation[*IESimOperationResult]
	DownloadProfileMetadataAsync(activationCode string) *IAsyncOperation[*IESimDownloadProfileMetadataResult]
	ResetAsync() *IAsyncOperation[*IESimOperationResult]
	Add_ProfileChanged(handler TypedEventHandler[*IESim, interface{}]) EventRegistrationToken
	Remove_ProfileChanged(token EventRegistrationToken)
}

type IESimVtbl struct {
	win32.IInspectableVtbl
	Get_AvailableMemoryInBytes       uintptr
	Get_Eid                          uintptr
	Get_FirmwareVersion              uintptr
	Get_MobileBroadbandModemDeviceId uintptr
	Get_Policy                       uintptr
	Get_State                        uintptr
	GetProfiles                      uintptr
	DeleteProfileAsync               uintptr
	DownloadProfileMetadataAsync     uintptr
	ResetAsync                       uintptr
	Add_ProfileChanged               uintptr
	Remove_ProfileChanged            uintptr
}

type IESim struct {
	win32.IInspectable
}

func (this *IESim) Vtbl() *IESimVtbl {
	return (*IESimVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IESim) Get_AvailableMemoryInBytes() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AvailableMemoryInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IESim) Get_Eid() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Eid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IESim) Get_FirmwareVersion() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FirmwareVersion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IESim) Get_MobileBroadbandModemDeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MobileBroadbandModemDeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IESim) Get_Policy() *IESimPolicy {
	var _result *IESimPolicy
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Policy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IESim) Get_State() ESimState {
	var _result ESimState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IESim) GetProfiles() *IVectorView[*IESimProfile] {
	var _result *IVectorView[*IESimProfile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetProfiles, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IESim) DeleteProfileAsync(profileId string) *IAsyncOperation[*IESimOperationResult] {
	var _result *IAsyncOperation[*IESimOperationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DeleteProfileAsync, uintptr(unsafe.Pointer(this)), NewHStr(profileId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IESim) DownloadProfileMetadataAsync(activationCode string) *IAsyncOperation[*IESimDownloadProfileMetadataResult] {
	var _result *IAsyncOperation[*IESimDownloadProfileMetadataResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DownloadProfileMetadataAsync, uintptr(unsafe.Pointer(this)), NewHStr(activationCode).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IESim) ResetAsync() *IAsyncOperation[*IESimOperationResult] {
	var _result *IAsyncOperation[*IESimOperationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ResetAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IESim) Add_ProfileChanged(handler TypedEventHandler[*IESim, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ProfileChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IESim) Remove_ProfileChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ProfileChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// BD4FD0A0-C68F-56EB-B99B-8F34B8100299
var IID_IESim2 = syscall.GUID{0xBD4FD0A0, 0xC68F, 0x56EB,
	[8]byte{0xB9, 0x9B, 0x8F, 0x34, 0xB8, 0x10, 0x02, 0x99}}

type IESim2Interface interface {
	win32.IInspectableInterface
	Discover() *IESimDiscoverResult
	DiscoverWithServerAddressAndMatchingId(serverAddress string, matchingId string) *IESimDiscoverResult
	DiscoverAsync() *IAsyncOperation[*IESimDiscoverResult]
	DiscoverWithServerAddressAndMatchingIdAsync(serverAddress string, matchingId string) *IAsyncOperation[*IESimDiscoverResult]
}

type IESim2Vtbl struct {
	win32.IInspectableVtbl
	Discover                                    uintptr
	DiscoverWithServerAddressAndMatchingId      uintptr
	DiscoverAsync                               uintptr
	DiscoverWithServerAddressAndMatchingIdAsync uintptr
}

type IESim2 struct {
	win32.IInspectable
}

func (this *IESim2) Vtbl() *IESim2Vtbl {
	return (*IESim2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IESim2) Discover() *IESimDiscoverResult {
	var _result *IESimDiscoverResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Discover, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IESim2) DiscoverWithServerAddressAndMatchingId(serverAddress string, matchingId string) *IESimDiscoverResult {
	var _result *IESimDiscoverResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DiscoverWithServerAddressAndMatchingId, uintptr(unsafe.Pointer(this)), NewHStr(serverAddress).Ptr, NewHStr(matchingId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IESim2) DiscoverAsync() *IAsyncOperation[*IESimDiscoverResult] {
	var _result *IAsyncOperation[*IESimDiscoverResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DiscoverAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IESim2) DiscoverWithServerAddressAndMatchingIdAsync(serverAddress string, matchingId string) *IAsyncOperation[*IESimDiscoverResult] {
	var _result *IAsyncOperation[*IESimDiscoverResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DiscoverWithServerAddressAndMatchingIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(serverAddress).Ptr, NewHStr(matchingId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FE1EDF45-01B8-5D31-B8D3-D9CBEBB2B831
var IID_IESim3 = syscall.GUID{0xFE1EDF45, 0x01B8, 0x5D31,
	[8]byte{0xB8, 0xD3, 0xD9, 0xCB, 0xEB, 0xB2, 0xB8, 0x31}}

type IESim3Interface interface {
	win32.IInspectableInterface
	Get_SlotIndex() *IReference[int32]
}

type IESim3Vtbl struct {
	win32.IInspectableVtbl
	Get_SlotIndex uintptr
}

type IESim3 struct {
	win32.IInspectable
}

func (this *IESim3) Vtbl() *IESim3Vtbl {
	return (*IESim3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IESim3) Get_SlotIndex() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SlotIndex, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 38BD0A58-4D5A-4D08-8DA7-E73EFF369DDD
var IID_IESimAddedEventArgs = syscall.GUID{0x38BD0A58, 0x4D5A, 0x4D08,
	[8]byte{0x8D, 0xA7, 0xE7, 0x3E, 0xFF, 0x36, 0x9D, 0xDD}}

type IESimAddedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ESim() *IESim
}

type IESimAddedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ESim uintptr
}

type IESimAddedEventArgs struct {
	win32.IInspectable
}

func (this *IESimAddedEventArgs) Vtbl() *IESimAddedEventArgsVtbl {
	return (*IESimAddedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IESimAddedEventArgs) Get_ESim() *IESim {
	var _result *IESim
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ESim, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E59AC3E3-39BC-5F6F-9321-0D4A182D261B
var IID_IESimDiscoverEvent = syscall.GUID{0xE59AC3E3, 0x39BC, 0x5F6F,
	[8]byte{0x93, 0x21, 0x0D, 0x4A, 0x18, 0x2D, 0x26, 0x1B}}

type IESimDiscoverEventInterface interface {
	win32.IInspectableInterface
	Get_MatchingId() string
	Get_RspServerAddress() string
}

type IESimDiscoverEventVtbl struct {
	win32.IInspectableVtbl
	Get_MatchingId       uintptr
	Get_RspServerAddress uintptr
}

type IESimDiscoverEvent struct {
	win32.IInspectable
}

func (this *IESimDiscoverEvent) Vtbl() *IESimDiscoverEventVtbl {
	return (*IESimDiscoverEventVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IESimDiscoverEvent) Get_MatchingId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MatchingId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IESimDiscoverEvent) Get_RspServerAddress() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RspServerAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 56B4BB5E-AB2F-5AC6-B359-DD5A8E237926
var IID_IESimDiscoverResult = syscall.GUID{0x56B4BB5E, 0xAB2F, 0x5AC6,
	[8]byte{0xB3, 0x59, 0xDD, 0x5A, 0x8E, 0x23, 0x79, 0x26}}

type IESimDiscoverResultInterface interface {
	win32.IInspectableInterface
	Get_Events() *IVectorView[*IESimDiscoverEvent]
	Get_Kind() ESimDiscoverResultKind
	Get_ProfileMetadata() *IESimProfileMetadata
	Get_Result() *IESimOperationResult
}

type IESimDiscoverResultVtbl struct {
	win32.IInspectableVtbl
	Get_Events          uintptr
	Get_Kind            uintptr
	Get_ProfileMetadata uintptr
	Get_Result          uintptr
}

type IESimDiscoverResult struct {
	win32.IInspectable
}

func (this *IESimDiscoverResult) Vtbl() *IESimDiscoverResultVtbl {
	return (*IESimDiscoverResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IESimDiscoverResult) Get_Events() *IVectorView[*IESimDiscoverEvent] {
	var _result *IVectorView[*IESimDiscoverEvent]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Events, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IESimDiscoverResult) Get_Kind() ESimDiscoverResultKind {
	var _result ESimDiscoverResultKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Kind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IESimDiscoverResult) Get_ProfileMetadata() *IESimProfileMetadata {
	var _result *IESimProfileMetadata
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProfileMetadata, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IESimDiscoverResult) Get_Result() *IESimOperationResult {
	var _result *IESimOperationResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Result, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C4234D9E-5AD6-426D-8D00-4434F449AFEC
var IID_IESimDownloadProfileMetadataResult = syscall.GUID{0xC4234D9E, 0x5AD6, 0x426D,
	[8]byte{0x8D, 0x00, 0x44, 0x34, 0xF4, 0x49, 0xAF, 0xEC}}

type IESimDownloadProfileMetadataResultInterface interface {
	win32.IInspectableInterface
	Get_Result() *IESimOperationResult
	Get_ProfileMetadata() *IESimProfileMetadata
}

type IESimDownloadProfileMetadataResultVtbl struct {
	win32.IInspectableVtbl
	Get_Result          uintptr
	Get_ProfileMetadata uintptr
}

type IESimDownloadProfileMetadataResult struct {
	win32.IInspectable
}

func (this *IESimDownloadProfileMetadataResult) Vtbl() *IESimDownloadProfileMetadataResultVtbl {
	return (*IESimDownloadProfileMetadataResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IESimDownloadProfileMetadataResult) Get_Result() *IESimOperationResult {
	var _result *IESimOperationResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Result, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IESimDownloadProfileMetadataResult) Get_ProfileMetadata() *IESimProfileMetadata {
	var _result *IESimProfileMetadata
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProfileMetadata, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0BFA2C0C-DF88-4631-BF04-C12E281B3962
var IID_IESimManagerStatics = syscall.GUID{0x0BFA2C0C, 0xDF88, 0x4631,
	[8]byte{0xBF, 0x04, 0xC1, 0x2E, 0x28, 0x1B, 0x39, 0x62}}

type IESimManagerStaticsInterface interface {
	win32.IInspectableInterface
	Get_ServiceInfo() *IESimServiceInfo
	TryCreateESimWatcher() *IESimWatcher
	Add_ServiceInfoChanged(handler EventHandler[interface{}]) EventRegistrationToken
	Remove_ServiceInfoChanged(token EventRegistrationToken)
}

type IESimManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_ServiceInfo           uintptr
	TryCreateESimWatcher      uintptr
	Add_ServiceInfoChanged    uintptr
	Remove_ServiceInfoChanged uintptr
}

type IESimManagerStatics struct {
	win32.IInspectable
}

func (this *IESimManagerStatics) Vtbl() *IESimManagerStaticsVtbl {
	return (*IESimManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IESimManagerStatics) Get_ServiceInfo() *IESimServiceInfo {
	var _result *IESimServiceInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IESimManagerStatics) TryCreateESimWatcher() *IESimWatcher {
	var _result *IESimWatcher
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryCreateESimWatcher, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IESimManagerStatics) Add_ServiceInfoChanged(handler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ServiceInfoChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IESimManagerStatics) Remove_ServiceInfoChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ServiceInfoChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// A67B63B1-309B-4E77-9E7E-CD93F1DDC7B9
var IID_IESimOperationResult = syscall.GUID{0xA67B63B1, 0x309B, 0x4E77,
	[8]byte{0x9E, 0x7E, 0xCD, 0x93, 0xF1, 0xDD, 0xC7, 0xB9}}

type IESimOperationResultInterface interface {
	win32.IInspectableInterface
	Get_Status() ESimOperationStatus
}

type IESimOperationResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
}

type IESimOperationResult struct {
	win32.IInspectable
}

func (this *IESimOperationResult) Vtbl() *IESimOperationResultVtbl {
	return (*IESimOperationResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IESimOperationResult) Get_Status() ESimOperationStatus {
	var _result ESimOperationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 41E1B99D-CF7E-4315-882B-6F1E74B0D38F
var IID_IESimPolicy = syscall.GUID{0x41E1B99D, 0xCF7E, 0x4315,
	[8]byte{0x88, 0x2B, 0x6F, 0x1E, 0x74, 0xB0, 0xD3, 0x8F}}

type IESimPolicyInterface interface {
	win32.IInspectableInterface
	Get_ShouldEnableManagingUi() bool
}

type IESimPolicyVtbl struct {
	win32.IInspectableVtbl
	Get_ShouldEnableManagingUi uintptr
}

type IESimPolicy struct {
	win32.IInspectable
}

func (this *IESimPolicy) Vtbl() *IESimPolicyVtbl {
	return (*IESimPolicyVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IESimPolicy) Get_ShouldEnableManagingUi() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ShouldEnableManagingUi, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// EE1E7880-06A9-4027-B4F8-DDB23D7810E0
var IID_IESimProfile = syscall.GUID{0xEE1E7880, 0x06A9, 0x4027,
	[8]byte{0xB4, 0xF8, 0xDD, 0xB2, 0x3D, 0x78, 0x10, 0xE0}}

type IESimProfileInterface interface {
	win32.IInspectableInterface
	Get_Class() ESimProfileClass
	Get_Nickname() string
	Get_Policy() *IESimProfilePolicy
	Get_Id() string
	Get_ProviderIcon() *IRandomAccessStreamReference
	Get_ProviderId() string
	Get_ProviderName() string
	Get_State() ESimProfileState
	DisableAsync() *IAsyncOperation[*IESimOperationResult]
	EnableAsync() *IAsyncOperation[*IESimOperationResult]
	SetNicknameAsync(newNickname string) *IAsyncOperation[*IESimOperationResult]
}

type IESimProfileVtbl struct {
	win32.IInspectableVtbl
	Get_Class        uintptr
	Get_Nickname     uintptr
	Get_Policy       uintptr
	Get_Id           uintptr
	Get_ProviderIcon uintptr
	Get_ProviderId   uintptr
	Get_ProviderName uintptr
	Get_State        uintptr
	DisableAsync     uintptr
	EnableAsync      uintptr
	SetNicknameAsync uintptr
}

type IESimProfile struct {
	win32.IInspectable
}

func (this *IESimProfile) Vtbl() *IESimProfileVtbl {
	return (*IESimProfileVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IESimProfile) Get_Class() ESimProfileClass {
	var _result ESimProfileClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Class, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IESimProfile) Get_Nickname() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Nickname, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IESimProfile) Get_Policy() *IESimProfilePolicy {
	var _result *IESimProfilePolicy
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Policy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IESimProfile) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IESimProfile) Get_ProviderIcon() *IRandomAccessStreamReference {
	var _result *IRandomAccessStreamReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProviderIcon, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IESimProfile) Get_ProviderId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProviderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IESimProfile) Get_ProviderName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProviderName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IESimProfile) Get_State() ESimProfileState {
	var _result ESimProfileState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IESimProfile) DisableAsync() *IAsyncOperation[*IESimOperationResult] {
	var _result *IAsyncOperation[*IESimOperationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DisableAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IESimProfile) EnableAsync() *IAsyncOperation[*IESimOperationResult] {
	var _result *IAsyncOperation[*IESimOperationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnableAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IESimProfile) SetNicknameAsync(newNickname string) *IAsyncOperation[*IESimOperationResult] {
	var _result *IAsyncOperation[*IESimOperationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetNicknameAsync, uintptr(unsafe.Pointer(this)), NewHStr(newNickname).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// ED25831F-90DB-498D-A7B4-EBCE807D3C23
var IID_IESimProfileMetadata = syscall.GUID{0xED25831F, 0x90DB, 0x498D,
	[8]byte{0xA7, 0xB4, 0xEB, 0xCE, 0x80, 0x7D, 0x3C, 0x23}}

type IESimProfileMetadataInterface interface {
	win32.IInspectableInterface
	Get_IsConfirmationCodeRequired() bool
	Get_Policy() *IESimProfilePolicy
	Get_Id() string
	Get_ProviderIcon() *IRandomAccessStreamReference
	Get_ProviderId() string
	Get_ProviderName() string
	Get_State() ESimProfileMetadataState
	DenyInstallAsync() *IAsyncOperation[*IESimOperationResult]
	ConfirmInstallAsync() *IAsyncOperationWithProgress[*IESimOperationResult, ESimProfileInstallProgress]
	ConfirmInstallWithConfirmationCodeAsync(confirmationCode string) *IAsyncOperationWithProgress[*IESimOperationResult, ESimProfileInstallProgress]
	PostponeInstallAsync() *IAsyncOperation[*IESimOperationResult]
	Add_StateChanged(handler TypedEventHandler[*IESimProfileMetadata, interface{}]) EventRegistrationToken
	Remove_StateChanged(token EventRegistrationToken)
}

type IESimProfileMetadataVtbl struct {
	win32.IInspectableVtbl
	Get_IsConfirmationCodeRequired          uintptr
	Get_Policy                              uintptr
	Get_Id                                  uintptr
	Get_ProviderIcon                        uintptr
	Get_ProviderId                          uintptr
	Get_ProviderName                        uintptr
	Get_State                               uintptr
	DenyInstallAsync                        uintptr
	ConfirmInstallAsync                     uintptr
	ConfirmInstallWithConfirmationCodeAsync uintptr
	PostponeInstallAsync                    uintptr
	Add_StateChanged                        uintptr
	Remove_StateChanged                     uintptr
}

type IESimProfileMetadata struct {
	win32.IInspectable
}

func (this *IESimProfileMetadata) Vtbl() *IESimProfileMetadataVtbl {
	return (*IESimProfileMetadataVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IESimProfileMetadata) Get_IsConfirmationCodeRequired() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsConfirmationCodeRequired, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IESimProfileMetadata) Get_Policy() *IESimProfilePolicy {
	var _result *IESimProfilePolicy
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Policy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IESimProfileMetadata) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IESimProfileMetadata) Get_ProviderIcon() *IRandomAccessStreamReference {
	var _result *IRandomAccessStreamReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProviderIcon, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IESimProfileMetadata) Get_ProviderId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProviderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IESimProfileMetadata) Get_ProviderName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProviderName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IESimProfileMetadata) Get_State() ESimProfileMetadataState {
	var _result ESimProfileMetadataState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IESimProfileMetadata) DenyInstallAsync() *IAsyncOperation[*IESimOperationResult] {
	var _result *IAsyncOperation[*IESimOperationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DenyInstallAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IESimProfileMetadata) ConfirmInstallAsync() *IAsyncOperationWithProgress[*IESimOperationResult, ESimProfileInstallProgress] {
	var _result *IAsyncOperationWithProgress[*IESimOperationResult, ESimProfileInstallProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConfirmInstallAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IESimProfileMetadata) ConfirmInstallWithConfirmationCodeAsync(confirmationCode string) *IAsyncOperationWithProgress[*IESimOperationResult, ESimProfileInstallProgress] {
	var _result *IAsyncOperationWithProgress[*IESimOperationResult, ESimProfileInstallProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConfirmInstallWithConfirmationCodeAsync, uintptr(unsafe.Pointer(this)), NewHStr(confirmationCode).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IESimProfileMetadata) PostponeInstallAsync() *IAsyncOperation[*IESimOperationResult] {
	var _result *IAsyncOperation[*IESimOperationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PostponeInstallAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IESimProfileMetadata) Add_StateChanged(handler TypedEventHandler[*IESimProfileMetadata, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IESimProfileMetadata) Remove_StateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// E6DD0F1D-9C5C-46C5-A289-A948999BF062
var IID_IESimProfilePolicy = syscall.GUID{0xE6DD0F1D, 0x9C5C, 0x46C5,
	[8]byte{0xA2, 0x89, 0xA9, 0x48, 0x99, 0x9B, 0xF0, 0x62}}

type IESimProfilePolicyInterface interface {
	win32.IInspectableInterface
	Get_CanDelete() bool
	Get_CanDisable() bool
	Get_IsManagedByEnterprise() bool
}

type IESimProfilePolicyVtbl struct {
	win32.IInspectableVtbl
	Get_CanDelete             uintptr
	Get_CanDisable            uintptr
	Get_IsManagedByEnterprise uintptr
}

type IESimProfilePolicy struct {
	win32.IInspectable
}

func (this *IESimProfilePolicy) Vtbl() *IESimProfilePolicyVtbl {
	return (*IESimProfilePolicyVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IESimProfilePolicy) Get_CanDelete() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanDelete, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IESimProfilePolicy) Get_CanDisable() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanDisable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IESimProfilePolicy) Get_IsManagedByEnterprise() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsManagedByEnterprise, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// DEC5277B-2FD9-4ED9-8376-D9B5E41278A3
var IID_IESimRemovedEventArgs = syscall.GUID{0xDEC5277B, 0x2FD9, 0x4ED9,
	[8]byte{0x83, 0x76, 0xD9, 0xB5, 0xE4, 0x12, 0x78, 0xA3}}

type IESimRemovedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ESim() *IESim
}

type IESimRemovedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ESim uintptr
}

type IESimRemovedEventArgs struct {
	win32.IInspectable
}

func (this *IESimRemovedEventArgs) Vtbl() *IESimRemovedEventArgsVtbl {
	return (*IESimRemovedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IESimRemovedEventArgs) Get_ESim() *IESim {
	var _result *IESim
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ESim, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F16AABCF-7F59-4A51-8494-BD89D5FF50EE
var IID_IESimServiceInfo = syscall.GUID{0xF16AABCF, 0x7F59, 0x4A51,
	[8]byte{0x84, 0x94, 0xBD, 0x89, 0xD5, 0xFF, 0x50, 0xEE}}

type IESimServiceInfoInterface interface {
	win32.IInspectableInterface
	Get_AuthenticationPreference() ESimAuthenticationPreference
	Get_IsESimUiEnabled() bool
}

type IESimServiceInfoVtbl struct {
	win32.IInspectableVtbl
	Get_AuthenticationPreference uintptr
	Get_IsESimUiEnabled          uintptr
}

type IESimServiceInfo struct {
	win32.IInspectable
}

func (this *IESimServiceInfo) Vtbl() *IESimServiceInfoVtbl {
	return (*IESimServiceInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IESimServiceInfo) Get_AuthenticationPreference() ESimAuthenticationPreference {
	var _result ESimAuthenticationPreference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AuthenticationPreference, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IESimServiceInfo) Get_IsESimUiEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsESimUiEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4C125CEC-508D-4B88-83CB-68BEF8168D12
var IID_IESimUpdatedEventArgs = syscall.GUID{0x4C125CEC, 0x508D, 0x4B88,
	[8]byte{0x83, 0xCB, 0x68, 0xBE, 0xF8, 0x16, 0x8D, 0x12}}

type IESimUpdatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ESim() *IESim
}

type IESimUpdatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ESim uintptr
}

type IESimUpdatedEventArgs struct {
	win32.IInspectable
}

func (this *IESimUpdatedEventArgs) Vtbl() *IESimUpdatedEventArgsVtbl {
	return (*IESimUpdatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IESimUpdatedEventArgs) Get_ESim() *IESim {
	var _result *IESim
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ESim, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C1F84CEB-A28D-4FBF-9771-6E31B81CCF22
var IID_IESimWatcher = syscall.GUID{0xC1F84CEB, 0xA28D, 0x4FBF,
	[8]byte{0x97, 0x71, 0x6E, 0x31, 0xB8, 0x1C, 0xCF, 0x22}}

type IESimWatcherInterface interface {
	win32.IInspectableInterface
	Get_Status() ESimWatcherStatus
	Start()
	Stop()
	Add_Added(handler TypedEventHandler[*IESimWatcher, *IESimAddedEventArgs]) EventRegistrationToken
	Remove_Added(token EventRegistrationToken)
	Add_EnumerationCompleted(handler TypedEventHandler[*IESimWatcher, interface{}]) EventRegistrationToken
	Remove_EnumerationCompleted(token EventRegistrationToken)
	Add_Removed(handler TypedEventHandler[*IESimWatcher, *IESimRemovedEventArgs]) EventRegistrationToken
	Remove_Removed(token EventRegistrationToken)
	Add_Stopped(handler TypedEventHandler[*IESimWatcher, interface{}]) EventRegistrationToken
	Remove_Stopped(token EventRegistrationToken)
	Add_Updated(handler TypedEventHandler[*IESimWatcher, *IESimUpdatedEventArgs]) EventRegistrationToken
	Remove_Updated(token EventRegistrationToken)
}

type IESimWatcherVtbl struct {
	win32.IInspectableVtbl
	Get_Status                  uintptr
	Start                       uintptr
	Stop                        uintptr
	Add_Added                   uintptr
	Remove_Added                uintptr
	Add_EnumerationCompleted    uintptr
	Remove_EnumerationCompleted uintptr
	Add_Removed                 uintptr
	Remove_Removed              uintptr
	Add_Stopped                 uintptr
	Remove_Stopped              uintptr
	Add_Updated                 uintptr
	Remove_Updated              uintptr
}

type IESimWatcher struct {
	win32.IInspectable
}

func (this *IESimWatcher) Vtbl() *IESimWatcherVtbl {
	return (*IESimWatcherVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IESimWatcher) Get_Status() ESimWatcherStatus {
	var _result ESimWatcherStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IESimWatcher) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IESimWatcher) Stop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IESimWatcher) Add_Added(handler TypedEventHandler[*IESimWatcher, *IESimAddedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Added, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IESimWatcher) Remove_Added(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Added, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IESimWatcher) Add_EnumerationCompleted(handler TypedEventHandler[*IESimWatcher, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_EnumerationCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IESimWatcher) Remove_EnumerationCompleted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_EnumerationCompleted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IESimWatcher) Add_Removed(handler TypedEventHandler[*IESimWatcher, *IESimRemovedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Removed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IESimWatcher) Remove_Removed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Removed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IESimWatcher) Add_Stopped(handler TypedEventHandler[*IESimWatcher, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Stopped, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IESimWatcher) Remove_Stopped(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Stopped, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IESimWatcher) Add_Updated(handler TypedEventHandler[*IESimWatcher, *IESimUpdatedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Updated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IESimWatcher) Remove_Updated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Updated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// F2AA4395-F1E6-4319-AA3E-477CA64B2BDF
var IID_IFdnAccessManagerStatics = syscall.GUID{0xF2AA4395, 0xF1E6, 0x4319,
	[8]byte{0xAA, 0x3E, 0x47, 0x7C, 0xA6, 0x4B, 0x2B, 0xDF}}

type IFdnAccessManagerStaticsInterface interface {
	win32.IInspectableInterface
	RequestUnlockAsync(contactListId string) *IAsyncOperation[bool]
}

type IFdnAccessManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	RequestUnlockAsync uintptr
}

type IFdnAccessManagerStatics struct {
	win32.IInspectable
}

func (this *IFdnAccessManagerStatics) Vtbl() *IFdnAccessManagerStaticsVtbl {
	return (*IFdnAccessManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFdnAccessManagerStatics) RequestUnlockAsync(contactListId string) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestUnlockAsync, uintptr(unsafe.Pointer(this)), NewHStr(contactListId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E756C791-1003-4DE5-83C7-DE61D88831D0
var IID_IHotspotAuthenticationContext = syscall.GUID{0xE756C791, 0x1003, 0x4DE5,
	[8]byte{0x83, 0xC7, 0xDE, 0x61, 0xD8, 0x88, 0x31, 0xD0}}

type IHotspotAuthenticationContextInterface interface {
	win32.IInspectableInterface
	Get_WirelessNetworkId() []byte
	Get_NetworkAdapter() *INetworkAdapter
	Get_RedirectMessageUrl() *IUriRuntimeClass
	Get_RedirectMessageXml() unsafe.Pointer
	Get_AuthenticationUrl() *IUriRuntimeClass
	IssueCredentials(userName string, password string, extraParameters string, markAsManualConnectOnFailure bool)
	AbortAuthentication(markAsManual bool)
	SkipAuthentication()
	TriggerAttentionRequired(packageRelativeApplicationId string, applicationParameters string)
}

type IHotspotAuthenticationContextVtbl struct {
	win32.IInspectableVtbl
	Get_WirelessNetworkId    uintptr
	Get_NetworkAdapter       uintptr
	Get_RedirectMessageUrl   uintptr
	Get_RedirectMessageXml   uintptr
	Get_AuthenticationUrl    uintptr
	IssueCredentials         uintptr
	AbortAuthentication      uintptr
	SkipAuthentication       uintptr
	TriggerAttentionRequired uintptr
}

type IHotspotAuthenticationContext struct {
	win32.IInspectable
}

func (this *IHotspotAuthenticationContext) Vtbl() *IHotspotAuthenticationContextVtbl {
	return (*IHotspotAuthenticationContextVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHotspotAuthenticationContext) Get_WirelessNetworkId() []byte {
	var _result []byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WirelessNetworkId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHotspotAuthenticationContext) Get_NetworkAdapter() *INetworkAdapter {
	var _result *INetworkAdapter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkAdapter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHotspotAuthenticationContext) Get_RedirectMessageUrl() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RedirectMessageUrl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHotspotAuthenticationContext) Get_RedirectMessageXml() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RedirectMessageXml, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHotspotAuthenticationContext) Get_AuthenticationUrl() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AuthenticationUrl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHotspotAuthenticationContext) IssueCredentials(userName string, password string, extraParameters string, markAsManualConnectOnFailure bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IssueCredentials, uintptr(unsafe.Pointer(this)), NewHStr(userName).Ptr, NewHStr(password).Ptr, NewHStr(extraParameters).Ptr, uintptr(*(*byte)(unsafe.Pointer(&markAsManualConnectOnFailure))))
	_ = _hr
}

func (this *IHotspotAuthenticationContext) AbortAuthentication(markAsManual bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AbortAuthentication, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&markAsManual))))
	_ = _hr
}

func (this *IHotspotAuthenticationContext) SkipAuthentication() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SkipAuthentication, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IHotspotAuthenticationContext) TriggerAttentionRequired(packageRelativeApplicationId string, applicationParameters string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TriggerAttentionRequired, uintptr(unsafe.Pointer(this)), NewHStr(packageRelativeApplicationId).Ptr, NewHStr(applicationParameters).Ptr)
	_ = _hr
}

// E756C791-1004-4DE5-83C7-DE61D88831D0
var IID_IHotspotAuthenticationContext2 = syscall.GUID{0xE756C791, 0x1004, 0x4DE5,
	[8]byte{0x83, 0xC7, 0xDE, 0x61, 0xD8, 0x88, 0x31, 0xD0}}

type IHotspotAuthenticationContext2Interface interface {
	win32.IInspectableInterface
	IssueCredentialsAsync(userName string, password string, extraParameters string, markAsManualConnectOnFailure bool) *IAsyncOperation[*IHotspotCredentialsAuthenticationResult]
}

type IHotspotAuthenticationContext2Vtbl struct {
	win32.IInspectableVtbl
	IssueCredentialsAsync uintptr
}

type IHotspotAuthenticationContext2 struct {
	win32.IInspectable
}

func (this *IHotspotAuthenticationContext2) Vtbl() *IHotspotAuthenticationContext2Vtbl {
	return (*IHotspotAuthenticationContext2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHotspotAuthenticationContext2) IssueCredentialsAsync(userName string, password string, extraParameters string, markAsManualConnectOnFailure bool) *IAsyncOperation[*IHotspotCredentialsAuthenticationResult] {
	var _result *IAsyncOperation[*IHotspotCredentialsAuthenticationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IssueCredentialsAsync, uintptr(unsafe.Pointer(this)), NewHStr(userName).Ptr, NewHStr(password).Ptr, NewHStr(extraParameters).Ptr, uintptr(*(*byte)(unsafe.Pointer(&markAsManualConnectOnFailure))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E756C791-1002-4DE5-83C7-DE61D88831D0
var IID_IHotspotAuthenticationContextStatics = syscall.GUID{0xE756C791, 0x1002, 0x4DE5,
	[8]byte{0x83, 0xC7, 0xDE, 0x61, 0xD8, 0x88, 0x31, 0xD0}}

type IHotspotAuthenticationContextStaticsInterface interface {
	win32.IInspectableInterface
	TryGetAuthenticationContext(evenToken string, context *IHotspotAuthenticationContext) bool
}

type IHotspotAuthenticationContextStaticsVtbl struct {
	win32.IInspectableVtbl
	TryGetAuthenticationContext uintptr
}

type IHotspotAuthenticationContextStatics struct {
	win32.IInspectable
}

func (this *IHotspotAuthenticationContextStatics) Vtbl() *IHotspotAuthenticationContextStaticsVtbl {
	return (*IHotspotAuthenticationContextStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHotspotAuthenticationContextStatics) TryGetAuthenticationContext(evenToken string, context *IHotspotAuthenticationContext) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetAuthenticationContext, uintptr(unsafe.Pointer(this)), NewHStr(evenToken).Ptr, uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// E756C791-1001-4DE5-83C7-DE61D88831D0
var IID_IHotspotAuthenticationEventDetails = syscall.GUID{0xE756C791, 0x1001, 0x4DE5,
	[8]byte{0x83, 0xC7, 0xDE, 0x61, 0xD8, 0x88, 0x31, 0xD0}}

type IHotspotAuthenticationEventDetailsInterface interface {
	win32.IInspectableInterface
	Get_EventToken() string
}

type IHotspotAuthenticationEventDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_EventToken uintptr
}

type IHotspotAuthenticationEventDetails struct {
	win32.IInspectable
}

func (this *IHotspotAuthenticationEventDetails) Vtbl() *IHotspotAuthenticationEventDetailsVtbl {
	return (*IHotspotAuthenticationEventDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHotspotAuthenticationEventDetails) Get_EventToken() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EventToken, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// E756C791-1005-4DE5-83C7-DE61D88831D0
var IID_IHotspotCredentialsAuthenticationResult = syscall.GUID{0xE756C791, 0x1005, 0x4DE5,
	[8]byte{0x83, 0xC7, 0xDE, 0x61, 0xD8, 0x88, 0x31, 0xD0}}

type IHotspotCredentialsAuthenticationResultInterface interface {
	win32.IInspectableInterface
	Get_HasNetworkErrorOccurred() bool
	Get_ResponseCode() HotspotAuthenticationResponseCode
	Get_LogoffUrl() *IUriRuntimeClass
	Get_AuthenticationReplyXml() unsafe.Pointer
}

type IHotspotCredentialsAuthenticationResultVtbl struct {
	win32.IInspectableVtbl
	Get_HasNetworkErrorOccurred uintptr
	Get_ResponseCode            uintptr
	Get_LogoffUrl               uintptr
	Get_AuthenticationReplyXml  uintptr
}

type IHotspotCredentialsAuthenticationResult struct {
	win32.IInspectable
}

func (this *IHotspotCredentialsAuthenticationResult) Vtbl() *IHotspotCredentialsAuthenticationResultVtbl {
	return (*IHotspotCredentialsAuthenticationResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHotspotCredentialsAuthenticationResult) Get_HasNetworkErrorOccurred() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasNetworkErrorOccurred, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHotspotCredentialsAuthenticationResult) Get_ResponseCode() HotspotAuthenticationResponseCode {
	var _result HotspotAuthenticationResponseCode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResponseCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHotspotCredentialsAuthenticationResult) Get_LogoffUrl() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LogoffUrl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHotspotCredentialsAuthenticationResult) Get_AuthenticationReplyXml() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AuthenticationReplyXml, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// B458AEED-49F1-4C22-B073-96D511BF9C35
var IID_IKnownCSimFilePathsStatics = syscall.GUID{0xB458AEED, 0x49F1, 0x4C22,
	[8]byte{0xB0, 0x73, 0x96, 0xD5, 0x11, 0xBF, 0x9C, 0x35}}

type IKnownCSimFilePathsStaticsInterface interface {
	win32.IInspectableInterface
	Get_EFSpn() *IVectorView[uint32]
	Get_Gid1() *IVectorView[uint32]
	Get_Gid2() *IVectorView[uint32]
}

type IKnownCSimFilePathsStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_EFSpn uintptr
	Get_Gid1  uintptr
	Get_Gid2  uintptr
}

type IKnownCSimFilePathsStatics struct {
	win32.IInspectable
}

func (this *IKnownCSimFilePathsStatics) Vtbl() *IKnownCSimFilePathsStaticsVtbl {
	return (*IKnownCSimFilePathsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKnownCSimFilePathsStatics) Get_EFSpn() *IVectorView[uint32] {
	var _result *IVectorView[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EFSpn, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKnownCSimFilePathsStatics) Get_Gid1() *IVectorView[uint32] {
	var _result *IVectorView[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Gid1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKnownCSimFilePathsStatics) Get_Gid2() *IVectorView[uint32] {
	var _result *IVectorView[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Gid2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3883C8B9-FF24-4571-A867-09F960426E14
var IID_IKnownRuimFilePathsStatics = syscall.GUID{0x3883C8B9, 0xFF24, 0x4571,
	[8]byte{0xA8, 0x67, 0x09, 0xF9, 0x60, 0x42, 0x6E, 0x14}}

type IKnownRuimFilePathsStaticsInterface interface {
	win32.IInspectableInterface
	Get_EFSpn() *IVectorView[uint32]
	Get_Gid1() *IVectorView[uint32]
	Get_Gid2() *IVectorView[uint32]
}

type IKnownRuimFilePathsStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_EFSpn uintptr
	Get_Gid1  uintptr
	Get_Gid2  uintptr
}

type IKnownRuimFilePathsStatics struct {
	win32.IInspectable
}

func (this *IKnownRuimFilePathsStatics) Vtbl() *IKnownRuimFilePathsStaticsVtbl {
	return (*IKnownRuimFilePathsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKnownRuimFilePathsStatics) Get_EFSpn() *IVectorView[uint32] {
	var _result *IVectorView[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EFSpn, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKnownRuimFilePathsStatics) Get_Gid1() *IVectorView[uint32] {
	var _result *IVectorView[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Gid1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKnownRuimFilePathsStatics) Get_Gid2() *IVectorView[uint32] {
	var _result *IVectorView[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Gid2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 80CD1A63-37A5-43D3-80A3-CCD23E8FECEE
var IID_IKnownSimFilePathsStatics = syscall.GUID{0x80CD1A63, 0x37A5, 0x43D3,
	[8]byte{0x80, 0xA3, 0xCC, 0xD2, 0x3E, 0x8F, 0xEC, 0xEE}}

type IKnownSimFilePathsStaticsInterface interface {
	win32.IInspectableInterface
	Get_EFOns() *IVectorView[uint32]
	Get_EFSpn() *IVectorView[uint32]
	Get_Gid1() *IVectorView[uint32]
	Get_Gid2() *IVectorView[uint32]
}

type IKnownSimFilePathsStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_EFOns uintptr
	Get_EFSpn uintptr
	Get_Gid1  uintptr
	Get_Gid2  uintptr
}

type IKnownSimFilePathsStatics struct {
	win32.IInspectable
}

func (this *IKnownSimFilePathsStatics) Vtbl() *IKnownSimFilePathsStaticsVtbl {
	return (*IKnownSimFilePathsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKnownSimFilePathsStatics) Get_EFOns() *IVectorView[uint32] {
	var _result *IVectorView[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EFOns, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKnownSimFilePathsStatics) Get_EFSpn() *IVectorView[uint32] {
	var _result *IVectorView[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EFSpn, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKnownSimFilePathsStatics) Get_Gid1() *IVectorView[uint32] {
	var _result *IVectorView[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Gid1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKnownSimFilePathsStatics) Get_Gid2() *IVectorView[uint32] {
	var _result *IVectorView[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Gid2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7C34E581-1F1B-43F4-9530-8B092D32D71F
var IID_IKnownUSimFilePathsStatics = syscall.GUID{0x7C34E581, 0x1F1B, 0x43F4,
	[8]byte{0x95, 0x30, 0x8B, 0x09, 0x2D, 0x32, 0xD7, 0x1F}}

type IKnownUSimFilePathsStaticsInterface interface {
	win32.IInspectableInterface
	Get_EFSpn() *IVectorView[uint32]
	Get_EFOpl() *IVectorView[uint32]
	Get_EFPnn() *IVectorView[uint32]
	Get_Gid1() *IVectorView[uint32]
	Get_Gid2() *IVectorView[uint32]
}

type IKnownUSimFilePathsStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_EFSpn uintptr
	Get_EFOpl uintptr
	Get_EFPnn uintptr
	Get_Gid1  uintptr
	Get_Gid2  uintptr
}

type IKnownUSimFilePathsStatics struct {
	win32.IInspectable
}

func (this *IKnownUSimFilePathsStatics) Vtbl() *IKnownUSimFilePathsStaticsVtbl {
	return (*IKnownUSimFilePathsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKnownUSimFilePathsStatics) Get_EFSpn() *IVectorView[uint32] {
	var _result *IVectorView[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EFSpn, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKnownUSimFilePathsStatics) Get_EFOpl() *IVectorView[uint32] {
	var _result *IVectorView[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EFOpl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKnownUSimFilePathsStatics) Get_EFPnn() *IVectorView[uint32] {
	var _result *IVectorView[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EFPnn, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKnownUSimFilePathsStatics) Get_Gid1() *IVectorView[uint32] {
	var _result *IVectorView[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Gid1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKnownUSimFilePathsStatics) Get_Gid2() *IVectorView[uint32] {
	var _result *IVectorView[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Gid2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 36C24CCD-CEE2-43E0-A603-EE86A36D6570
var IID_IMobileBroadbandAccount = syscall.GUID{0x36C24CCD, 0xCEE2, 0x43E0,
	[8]byte{0xA6, 0x03, 0xEE, 0x86, 0xA3, 0x6D, 0x65, 0x70}}

type IMobileBroadbandAccountInterface interface {
	win32.IInspectableInterface
	Get_NetworkAccountId() string
	Get_ServiceProviderGuid() syscall.GUID
	Get_ServiceProviderName() string
	Get_CurrentNetwork() *IMobileBroadbandNetwork
	Get_CurrentDeviceInformation() *IMobileBroadbandDeviceInformation
}

type IMobileBroadbandAccountVtbl struct {
	win32.IInspectableVtbl
	Get_NetworkAccountId         uintptr
	Get_ServiceProviderGuid      uintptr
	Get_ServiceProviderName      uintptr
	Get_CurrentNetwork           uintptr
	Get_CurrentDeviceInformation uintptr
}

type IMobileBroadbandAccount struct {
	win32.IInspectable
}

func (this *IMobileBroadbandAccount) Vtbl() *IMobileBroadbandAccountVtbl {
	return (*IMobileBroadbandAccountVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandAccount) Get_NetworkAccountId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkAccountId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandAccount) Get_ServiceProviderGuid() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceProviderGuid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandAccount) Get_ServiceProviderName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceProviderName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandAccount) Get_CurrentNetwork() *IMobileBroadbandNetwork {
	var _result *IMobileBroadbandNetwork
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentNetwork, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandAccount) Get_CurrentDeviceInformation() *IMobileBroadbandDeviceInformation {
	var _result *IMobileBroadbandDeviceInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentDeviceInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 38F52F1C-1136-4257-959F-B658A352B6D4
var IID_IMobileBroadbandAccount2 = syscall.GUID{0x38F52F1C, 0x1136, 0x4257,
	[8]byte{0x95, 0x9F, 0xB6, 0x58, 0xA3, 0x52, 0xB6, 0xD4}}

type IMobileBroadbandAccount2Interface interface {
	win32.IInspectableInterface
	GetConnectionProfiles() *IVectorView[*IConnectionProfile]
}

type IMobileBroadbandAccount2Vtbl struct {
	win32.IInspectableVtbl
	GetConnectionProfiles uintptr
}

type IMobileBroadbandAccount2 struct {
	win32.IInspectable
}

func (this *IMobileBroadbandAccount2) Vtbl() *IMobileBroadbandAccount2Vtbl {
	return (*IMobileBroadbandAccount2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandAccount2) GetConnectionProfiles() *IVectorView[*IConnectionProfile] {
	var _result *IVectorView[*IConnectionProfile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetConnectionProfiles, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 092A1E21-9379-4B9B-AD31-D5FEE2F748C6
var IID_IMobileBroadbandAccount3 = syscall.GUID{0x092A1E21, 0x9379, 0x4B9B,
	[8]byte{0xAD, 0x31, 0xD5, 0xFE, 0xE2, 0xF7, 0x48, 0xC6}}

type IMobileBroadbandAccount3Interface interface {
	win32.IInspectableInterface
	Get_AccountExperienceUrl() *IUriRuntimeClass
}

type IMobileBroadbandAccount3Vtbl struct {
	win32.IInspectableVtbl
	Get_AccountExperienceUrl uintptr
}

type IMobileBroadbandAccount3 struct {
	win32.IInspectable
}

func (this *IMobileBroadbandAccount3) Vtbl() *IMobileBroadbandAccount3Vtbl {
	return (*IMobileBroadbandAccount3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandAccount3) Get_AccountExperienceUrl() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AccountExperienceUrl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3853C880-77DE-4C04-BEAD-A123B08C9F59
var IID_IMobileBroadbandAccountEventArgs = syscall.GUID{0x3853C880, 0x77DE, 0x4C04,
	[8]byte{0xBE, 0xAD, 0xA1, 0x23, 0xB0, 0x8C, 0x9F, 0x59}}

type IMobileBroadbandAccountEventArgsInterface interface {
	win32.IInspectableInterface
	Get_NetworkAccountId() string
}

type IMobileBroadbandAccountEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_NetworkAccountId uintptr
}

type IMobileBroadbandAccountEventArgs struct {
	win32.IInspectable
}

func (this *IMobileBroadbandAccountEventArgs) Vtbl() *IMobileBroadbandAccountEventArgsVtbl {
	return (*IMobileBroadbandAccountEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandAccountEventArgs) Get_NetworkAccountId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkAccountId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// AA7F4D24-AFC1-4FC8-AE9A-A9175310FAAD
var IID_IMobileBroadbandAccountStatics = syscall.GUID{0xAA7F4D24, 0xAFC1, 0x4FC8,
	[8]byte{0xAE, 0x9A, 0xA9, 0x17, 0x53, 0x10, 0xFA, 0xAD}}

type IMobileBroadbandAccountStaticsInterface interface {
	win32.IInspectableInterface
	Get_AvailableNetworkAccountIds() *IVectorView[string]
	CreateFromNetworkAccountId(networkAccountId string) *IMobileBroadbandAccount
}

type IMobileBroadbandAccountStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_AvailableNetworkAccountIds uintptr
	CreateFromNetworkAccountId     uintptr
}

type IMobileBroadbandAccountStatics struct {
	win32.IInspectable
}

func (this *IMobileBroadbandAccountStatics) Vtbl() *IMobileBroadbandAccountStaticsVtbl {
	return (*IMobileBroadbandAccountStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandAccountStatics) Get_AvailableNetworkAccountIds() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AvailableNetworkAccountIds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandAccountStatics) CreateFromNetworkAccountId(networkAccountId string) *IMobileBroadbandAccount {
	var _result *IMobileBroadbandAccount
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromNetworkAccountId, uintptr(unsafe.Pointer(this)), NewHStr(networkAccountId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7BC31D88-A6BD-49E1-80AB-6B91354A57D4
var IID_IMobileBroadbandAccountUpdatedEventArgs = syscall.GUID{0x7BC31D88, 0xA6BD, 0x49E1,
	[8]byte{0x80, 0xAB, 0x6B, 0x91, 0x35, 0x4A, 0x57, 0xD4}}

type IMobileBroadbandAccountUpdatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_NetworkAccountId() string
	Get_HasDeviceInformationChanged() bool
	Get_HasNetworkChanged() bool
}

type IMobileBroadbandAccountUpdatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_NetworkAccountId            uintptr
	Get_HasDeviceInformationChanged uintptr
	Get_HasNetworkChanged           uintptr
}

type IMobileBroadbandAccountUpdatedEventArgs struct {
	win32.IInspectable
}

func (this *IMobileBroadbandAccountUpdatedEventArgs) Vtbl() *IMobileBroadbandAccountUpdatedEventArgsVtbl {
	return (*IMobileBroadbandAccountUpdatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandAccountUpdatedEventArgs) Get_NetworkAccountId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkAccountId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandAccountUpdatedEventArgs) Get_HasDeviceInformationChanged() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasDeviceInformationChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandAccountUpdatedEventArgs) Get_HasNetworkChanged() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasNetworkChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 6BF3335E-23B5-449F-928D-5E0D3E04471D
var IID_IMobileBroadbandAccountWatcher = syscall.GUID{0x6BF3335E, 0x23B5, 0x449F,
	[8]byte{0x92, 0x8D, 0x5E, 0x0D, 0x3E, 0x04, 0x47, 0x1D}}

type IMobileBroadbandAccountWatcherInterface interface {
	win32.IInspectableInterface
	Add_AccountAdded(handler TypedEventHandler[*IMobileBroadbandAccountWatcher, *IMobileBroadbandAccountEventArgs]) EventRegistrationToken
	Remove_AccountAdded(cookie EventRegistrationToken)
	Add_AccountUpdated(handler TypedEventHandler[*IMobileBroadbandAccountWatcher, *IMobileBroadbandAccountUpdatedEventArgs]) EventRegistrationToken
	Remove_AccountUpdated(cookie EventRegistrationToken)
	Add_AccountRemoved(handler TypedEventHandler[*IMobileBroadbandAccountWatcher, *IMobileBroadbandAccountEventArgs]) EventRegistrationToken
	Remove_AccountRemoved(cookie EventRegistrationToken)
	Add_EnumerationCompleted(handler TypedEventHandler[*IMobileBroadbandAccountWatcher, interface{}]) EventRegistrationToken
	Remove_EnumerationCompleted(cookie EventRegistrationToken)
	Add_Stopped(handler TypedEventHandler[*IMobileBroadbandAccountWatcher, interface{}]) EventRegistrationToken
	Remove_Stopped(cookie EventRegistrationToken)
	Get_Status() MobileBroadbandAccountWatcherStatus
	Start()
	Stop()
}

type IMobileBroadbandAccountWatcherVtbl struct {
	win32.IInspectableVtbl
	Add_AccountAdded            uintptr
	Remove_AccountAdded         uintptr
	Add_AccountUpdated          uintptr
	Remove_AccountUpdated       uintptr
	Add_AccountRemoved          uintptr
	Remove_AccountRemoved       uintptr
	Add_EnumerationCompleted    uintptr
	Remove_EnumerationCompleted uintptr
	Add_Stopped                 uintptr
	Remove_Stopped              uintptr
	Get_Status                  uintptr
	Start                       uintptr
	Stop                        uintptr
}

type IMobileBroadbandAccountWatcher struct {
	win32.IInspectable
}

func (this *IMobileBroadbandAccountWatcher) Vtbl() *IMobileBroadbandAccountWatcherVtbl {
	return (*IMobileBroadbandAccountWatcherVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandAccountWatcher) Add_AccountAdded(handler TypedEventHandler[*IMobileBroadbandAccountWatcher, *IMobileBroadbandAccountEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AccountAdded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandAccountWatcher) Remove_AccountAdded(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AccountAdded, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IMobileBroadbandAccountWatcher) Add_AccountUpdated(handler TypedEventHandler[*IMobileBroadbandAccountWatcher, *IMobileBroadbandAccountUpdatedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AccountUpdated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandAccountWatcher) Remove_AccountUpdated(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AccountUpdated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IMobileBroadbandAccountWatcher) Add_AccountRemoved(handler TypedEventHandler[*IMobileBroadbandAccountWatcher, *IMobileBroadbandAccountEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AccountRemoved, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandAccountWatcher) Remove_AccountRemoved(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AccountRemoved, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IMobileBroadbandAccountWatcher) Add_EnumerationCompleted(handler TypedEventHandler[*IMobileBroadbandAccountWatcher, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_EnumerationCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandAccountWatcher) Remove_EnumerationCompleted(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_EnumerationCompleted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IMobileBroadbandAccountWatcher) Add_Stopped(handler TypedEventHandler[*IMobileBroadbandAccountWatcher, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Stopped, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandAccountWatcher) Remove_Stopped(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Stopped, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IMobileBroadbandAccountWatcher) Get_Status() MobileBroadbandAccountWatcherStatus {
	var _result MobileBroadbandAccountWatcherStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandAccountWatcher) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IMobileBroadbandAccountWatcher) Stop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// B9AF4B7E-CBF9-4109-90BE-5C06BFD513B6
var IID_IMobileBroadbandAntennaSar = syscall.GUID{0xB9AF4B7E, 0xCBF9, 0x4109,
	[8]byte{0x90, 0xBE, 0x5C, 0x06, 0xBF, 0xD5, 0x13, 0xB6}}

type IMobileBroadbandAntennaSarInterface interface {
	win32.IInspectableInterface
	Get_AntennaIndex() int32
	Get_SarBackoffIndex() int32
}

type IMobileBroadbandAntennaSarVtbl struct {
	win32.IInspectableVtbl
	Get_AntennaIndex    uintptr
	Get_SarBackoffIndex uintptr
}

type IMobileBroadbandAntennaSar struct {
	win32.IInspectable
}

func (this *IMobileBroadbandAntennaSar) Vtbl() *IMobileBroadbandAntennaSarVtbl {
	return (*IMobileBroadbandAntennaSarVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandAntennaSar) Get_AntennaIndex() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AntennaIndex, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandAntennaSar) Get_SarBackoffIndex() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SarBackoffIndex, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A91E1716-C04D-4A21-8698-1459DC672C6E
var IID_IMobileBroadbandAntennaSarFactory = syscall.GUID{0xA91E1716, 0xC04D, 0x4A21,
	[8]byte{0x86, 0x98, 0x14, 0x59, 0xDC, 0x67, 0x2C, 0x6E}}

type IMobileBroadbandAntennaSarFactoryInterface interface {
	win32.IInspectableInterface
	CreateWithIndex(antennaIndex int32, sarBackoffIndex int32) *IMobileBroadbandAntennaSar
}

type IMobileBroadbandAntennaSarFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateWithIndex uintptr
}

type IMobileBroadbandAntennaSarFactory struct {
	win32.IInspectable
}

func (this *IMobileBroadbandAntennaSarFactory) Vtbl() *IMobileBroadbandAntennaSarFactoryVtbl {
	return (*IMobileBroadbandAntennaSarFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandAntennaSarFactory) CreateWithIndex(antennaIndex int32, sarBackoffIndex int32) *IMobileBroadbandAntennaSar {
	var _result *IMobileBroadbandAntennaSar
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithIndex, uintptr(unsafe.Pointer(this)), uintptr(antennaIndex), uintptr(sarBackoffIndex), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0601B3B4-411A-4F2E-8287-76F5650C60CD
var IID_IMobileBroadbandCellCdma = syscall.GUID{0x0601B3B4, 0x411A, 0x4F2E,
	[8]byte{0x82, 0x87, 0x76, 0xF5, 0x65, 0x0C, 0x60, 0xCD}}

type IMobileBroadbandCellCdmaInterface interface {
	win32.IInspectableInterface
	Get_BaseStationId() *IReference[int32]
	Get_BaseStationPNCode() *IReference[int32]
	Get_BaseStationLatitude() *IReference[float64]
	Get_BaseStationLongitude() *IReference[float64]
	Get_BaseStationLastBroadcastGpsTime() *IReference[TimeSpan]
	Get_NetworkId() *IReference[int32]
	Get_PilotSignalStrengthInDB() *IReference[float64]
	Get_SystemId() *IReference[int32]
}

type IMobileBroadbandCellCdmaVtbl struct {
	win32.IInspectableVtbl
	Get_BaseStationId                   uintptr
	Get_BaseStationPNCode               uintptr
	Get_BaseStationLatitude             uintptr
	Get_BaseStationLongitude            uintptr
	Get_BaseStationLastBroadcastGpsTime uintptr
	Get_NetworkId                       uintptr
	Get_PilotSignalStrengthInDB         uintptr
	Get_SystemId                        uintptr
}

type IMobileBroadbandCellCdma struct {
	win32.IInspectable
}

func (this *IMobileBroadbandCellCdma) Vtbl() *IMobileBroadbandCellCdmaVtbl {
	return (*IMobileBroadbandCellCdmaVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandCellCdma) Get_BaseStationId() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BaseStationId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellCdma) Get_BaseStationPNCode() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BaseStationPNCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellCdma) Get_BaseStationLatitude() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BaseStationLatitude, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellCdma) Get_BaseStationLongitude() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BaseStationLongitude, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellCdma) Get_BaseStationLastBroadcastGpsTime() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BaseStationLastBroadcastGpsTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellCdma) Get_NetworkId() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellCdma) Get_PilotSignalStrengthInDB() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PilotSignalStrengthInDB, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellCdma) Get_SystemId() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SystemId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CC917F06-7EE0-47B8-9E1F-C3B48DF9DF5B
var IID_IMobileBroadbandCellGsm = syscall.GUID{0xCC917F06, 0x7EE0, 0x47B8,
	[8]byte{0x9E, 0x1F, 0xC3, 0xB4, 0x8D, 0xF9, 0xDF, 0x5B}}

type IMobileBroadbandCellGsmInterface interface {
	win32.IInspectableInterface
	Get_BaseStationId() *IReference[int32]
	Get_CellId() *IReference[int32]
	Get_ChannelNumber() *IReference[int32]
	Get_LocationAreaCode() *IReference[int32]
	Get_ProviderId() string
	Get_ReceivedSignalStrengthInDBm() *IReference[float64]
	Get_TimingAdvanceInBitPeriods() *IReference[int32]
}

type IMobileBroadbandCellGsmVtbl struct {
	win32.IInspectableVtbl
	Get_BaseStationId               uintptr
	Get_CellId                      uintptr
	Get_ChannelNumber               uintptr
	Get_LocationAreaCode            uintptr
	Get_ProviderId                  uintptr
	Get_ReceivedSignalStrengthInDBm uintptr
	Get_TimingAdvanceInBitPeriods   uintptr
}

type IMobileBroadbandCellGsm struct {
	win32.IInspectable
}

func (this *IMobileBroadbandCellGsm) Vtbl() *IMobileBroadbandCellGsmVtbl {
	return (*IMobileBroadbandCellGsmVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandCellGsm) Get_BaseStationId() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BaseStationId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellGsm) Get_CellId() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CellId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellGsm) Get_ChannelNumber() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ChannelNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellGsm) Get_LocationAreaCode() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocationAreaCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellGsm) Get_ProviderId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProviderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandCellGsm) Get_ReceivedSignalStrengthInDBm() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReceivedSignalStrengthInDBm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellGsm) Get_TimingAdvanceInBitPeriods() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TimingAdvanceInBitPeriods, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9197C87B-2B78-456D-8B53-AAA25D0AF741
var IID_IMobileBroadbandCellLte = syscall.GUID{0x9197C87B, 0x2B78, 0x456D,
	[8]byte{0x8B, 0x53, 0xAA, 0xA2, 0x5D, 0x0A, 0xF7, 0x41}}

type IMobileBroadbandCellLteInterface interface {
	win32.IInspectableInterface
	Get_CellId() *IReference[int32]
	Get_ChannelNumber() *IReference[int32]
	Get_PhysicalCellId() *IReference[int32]
	Get_ProviderId() string
	Get_ReferenceSignalReceivedPowerInDBm() *IReference[float64]
	Get_ReferenceSignalReceivedQualityInDBm() *IReference[float64]
	Get_TimingAdvanceInBitPeriods() *IReference[int32]
	Get_TrackingAreaCode() *IReference[int32]
}

type IMobileBroadbandCellLteVtbl struct {
	win32.IInspectableVtbl
	Get_CellId                              uintptr
	Get_ChannelNumber                       uintptr
	Get_PhysicalCellId                      uintptr
	Get_ProviderId                          uintptr
	Get_ReferenceSignalReceivedPowerInDBm   uintptr
	Get_ReferenceSignalReceivedQualityInDBm uintptr
	Get_TimingAdvanceInBitPeriods           uintptr
	Get_TrackingAreaCode                    uintptr
}

type IMobileBroadbandCellLte struct {
	win32.IInspectable
}

func (this *IMobileBroadbandCellLte) Vtbl() *IMobileBroadbandCellLteVtbl {
	return (*IMobileBroadbandCellLteVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandCellLte) Get_CellId() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CellId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellLte) Get_ChannelNumber() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ChannelNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellLte) Get_PhysicalCellId() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhysicalCellId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellLte) Get_ProviderId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProviderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandCellLte) Get_ReferenceSignalReceivedPowerInDBm() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReferenceSignalReceivedPowerInDBm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellLte) Get_ReferenceSignalReceivedQualityInDBm() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReferenceSignalReceivedQualityInDBm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellLte) Get_TimingAdvanceInBitPeriods() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TimingAdvanceInBitPeriods, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellLte) Get_TrackingAreaCode() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TrackingAreaCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A13F0DEB-66FC-4B4B-83A9-A487A3A5A0A6
var IID_IMobileBroadbandCellNR = syscall.GUID{0xA13F0DEB, 0x66FC, 0x4B4B,
	[8]byte{0x83, 0xA9, 0xA4, 0x87, 0xA3, 0xA5, 0xA0, 0xA6}}

type IMobileBroadbandCellNRInterface interface {
	win32.IInspectableInterface
	Get_CellId() *IReference[int64]
	Get_ChannelNumber() *IReference[int32]
	Get_PhysicalCellId() *IReference[int32]
	Get_ProviderId() string
	Get_ReferenceSignalReceivedPowerInDBm() *IReference[float64]
	Get_ReferenceSignalReceivedQualityInDBm() *IReference[float64]
	Get_TimingAdvanceInNanoseconds() *IReference[int32]
	Get_TrackingAreaCode() *IReference[int32]
	Get_SignalToNoiseRatioInDB() *IReference[float64]
}

type IMobileBroadbandCellNRVtbl struct {
	win32.IInspectableVtbl
	Get_CellId                              uintptr
	Get_ChannelNumber                       uintptr
	Get_PhysicalCellId                      uintptr
	Get_ProviderId                          uintptr
	Get_ReferenceSignalReceivedPowerInDBm   uintptr
	Get_ReferenceSignalReceivedQualityInDBm uintptr
	Get_TimingAdvanceInNanoseconds          uintptr
	Get_TrackingAreaCode                    uintptr
	Get_SignalToNoiseRatioInDB              uintptr
}

type IMobileBroadbandCellNR struct {
	win32.IInspectable
}

func (this *IMobileBroadbandCellNR) Vtbl() *IMobileBroadbandCellNRVtbl {
	return (*IMobileBroadbandCellNRVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandCellNR) Get_CellId() *IReference[int64] {
	var _result *IReference[int64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CellId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellNR) Get_ChannelNumber() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ChannelNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellNR) Get_PhysicalCellId() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhysicalCellId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellNR) Get_ProviderId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProviderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandCellNR) Get_ReferenceSignalReceivedPowerInDBm() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReferenceSignalReceivedPowerInDBm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellNR) Get_ReferenceSignalReceivedQualityInDBm() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReferenceSignalReceivedQualityInDBm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellNR) Get_TimingAdvanceInNanoseconds() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TimingAdvanceInNanoseconds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellNR) Get_TrackingAreaCode() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TrackingAreaCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellNR) Get_SignalToNoiseRatioInDB() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SignalToNoiseRatioInDB, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0EDA1655-DB0E-4182-8CDA-CC419A7BDE08
var IID_IMobileBroadbandCellTdscdma = syscall.GUID{0x0EDA1655, 0xDB0E, 0x4182,
	[8]byte{0x8C, 0xDA, 0xCC, 0x41, 0x9A, 0x7B, 0xDE, 0x08}}

type IMobileBroadbandCellTdscdmaInterface interface {
	win32.IInspectableInterface
	Get_CellId() *IReference[int32]
	Get_CellParameterId() *IReference[int32]
	Get_ChannelNumber() *IReference[int32]
	Get_LocationAreaCode() *IReference[int32]
	Get_PathLossInDB() *IReference[float64]
	Get_ProviderId() string
	Get_ReceivedSignalCodePowerInDBm() *IReference[float64]
	Get_TimingAdvanceInBitPeriods() *IReference[int32]
}

type IMobileBroadbandCellTdscdmaVtbl struct {
	win32.IInspectableVtbl
	Get_CellId                       uintptr
	Get_CellParameterId              uintptr
	Get_ChannelNumber                uintptr
	Get_LocationAreaCode             uintptr
	Get_PathLossInDB                 uintptr
	Get_ProviderId                   uintptr
	Get_ReceivedSignalCodePowerInDBm uintptr
	Get_TimingAdvanceInBitPeriods    uintptr
}

type IMobileBroadbandCellTdscdma struct {
	win32.IInspectable
}

func (this *IMobileBroadbandCellTdscdma) Vtbl() *IMobileBroadbandCellTdscdmaVtbl {
	return (*IMobileBroadbandCellTdscdmaVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandCellTdscdma) Get_CellId() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CellId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellTdscdma) Get_CellParameterId() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CellParameterId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellTdscdma) Get_ChannelNumber() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ChannelNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellTdscdma) Get_LocationAreaCode() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocationAreaCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellTdscdma) Get_PathLossInDB() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PathLossInDB, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellTdscdma) Get_ProviderId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProviderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandCellTdscdma) Get_ReceivedSignalCodePowerInDBm() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReceivedSignalCodePowerInDBm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellTdscdma) Get_TimingAdvanceInBitPeriods() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TimingAdvanceInBitPeriods, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 77B4B5AE-49C8-4F15-B285-4C26A7F67215
var IID_IMobileBroadbandCellUmts = syscall.GUID{0x77B4B5AE, 0x49C8, 0x4F15,
	[8]byte{0xB2, 0x85, 0x4C, 0x26, 0xA7, 0xF6, 0x72, 0x15}}

type IMobileBroadbandCellUmtsInterface interface {
	win32.IInspectableInterface
	Get_CellId() *IReference[int32]
	Get_ChannelNumber() *IReference[int32]
	Get_LocationAreaCode() *IReference[int32]
	Get_PathLossInDB() *IReference[float64]
	Get_PrimaryScramblingCode() *IReference[int32]
	Get_ProviderId() string
	Get_ReceivedSignalCodePowerInDBm() *IReference[float64]
	Get_SignalToNoiseRatioInDB() *IReference[float64]
}

type IMobileBroadbandCellUmtsVtbl struct {
	win32.IInspectableVtbl
	Get_CellId                       uintptr
	Get_ChannelNumber                uintptr
	Get_LocationAreaCode             uintptr
	Get_PathLossInDB                 uintptr
	Get_PrimaryScramblingCode        uintptr
	Get_ProviderId                   uintptr
	Get_ReceivedSignalCodePowerInDBm uintptr
	Get_SignalToNoiseRatioInDB       uintptr
}

type IMobileBroadbandCellUmts struct {
	win32.IInspectable
}

func (this *IMobileBroadbandCellUmts) Vtbl() *IMobileBroadbandCellUmtsVtbl {
	return (*IMobileBroadbandCellUmtsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandCellUmts) Get_CellId() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CellId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellUmts) Get_ChannelNumber() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ChannelNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellUmts) Get_LocationAreaCode() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocationAreaCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellUmts) Get_PathLossInDB() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PathLossInDB, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellUmts) Get_PrimaryScramblingCode() *IReference[int32] {
	var _result *IReference[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PrimaryScramblingCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellUmts) Get_ProviderId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProviderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandCellUmts) Get_ReceivedSignalCodePowerInDBm() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReceivedSignalCodePowerInDBm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellUmts) Get_SignalToNoiseRatioInDB() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SignalToNoiseRatioInDB, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 89A9562A-E472-4DA5-929C-DE61711DD261
var IID_IMobileBroadbandCellsInfo = syscall.GUID{0x89A9562A, 0xE472, 0x4DA5,
	[8]byte{0x92, 0x9C, 0xDE, 0x61, 0x71, 0x1D, 0xD2, 0x61}}

type IMobileBroadbandCellsInfoInterface interface {
	win32.IInspectableInterface
	Get_NeighboringCellsCdma() *IVectorView[*IMobileBroadbandCellCdma]
	Get_NeighboringCellsGsm() *IVectorView[*IMobileBroadbandCellGsm]
	Get_NeighboringCellsLte() *IVectorView[*IMobileBroadbandCellLte]
	Get_NeighboringCellsTdscdma() *IVectorView[*IMobileBroadbandCellTdscdma]
	Get_NeighboringCellsUmts() *IVectorView[*IMobileBroadbandCellUmts]
	Get_ServingCellsCdma() *IVectorView[*IMobileBroadbandCellCdma]
	Get_ServingCellsGsm() *IVectorView[*IMobileBroadbandCellGsm]
	Get_ServingCellsLte() *IVectorView[*IMobileBroadbandCellLte]
	Get_ServingCellsTdscdma() *IVectorView[*IMobileBroadbandCellTdscdma]
	Get_ServingCellsUmts() *IVectorView[*IMobileBroadbandCellUmts]
}

type IMobileBroadbandCellsInfoVtbl struct {
	win32.IInspectableVtbl
	Get_NeighboringCellsCdma    uintptr
	Get_NeighboringCellsGsm     uintptr
	Get_NeighboringCellsLte     uintptr
	Get_NeighboringCellsTdscdma uintptr
	Get_NeighboringCellsUmts    uintptr
	Get_ServingCellsCdma        uintptr
	Get_ServingCellsGsm         uintptr
	Get_ServingCellsLte         uintptr
	Get_ServingCellsTdscdma     uintptr
	Get_ServingCellsUmts        uintptr
}

type IMobileBroadbandCellsInfo struct {
	win32.IInspectable
}

func (this *IMobileBroadbandCellsInfo) Vtbl() *IMobileBroadbandCellsInfoVtbl {
	return (*IMobileBroadbandCellsInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandCellsInfo) Get_NeighboringCellsCdma() *IVectorView[*IMobileBroadbandCellCdma] {
	var _result *IVectorView[*IMobileBroadbandCellCdma]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NeighboringCellsCdma, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellsInfo) Get_NeighboringCellsGsm() *IVectorView[*IMobileBroadbandCellGsm] {
	var _result *IVectorView[*IMobileBroadbandCellGsm]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NeighboringCellsGsm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellsInfo) Get_NeighboringCellsLte() *IVectorView[*IMobileBroadbandCellLte] {
	var _result *IVectorView[*IMobileBroadbandCellLte]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NeighboringCellsLte, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellsInfo) Get_NeighboringCellsTdscdma() *IVectorView[*IMobileBroadbandCellTdscdma] {
	var _result *IVectorView[*IMobileBroadbandCellTdscdma]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NeighboringCellsTdscdma, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellsInfo) Get_NeighboringCellsUmts() *IVectorView[*IMobileBroadbandCellUmts] {
	var _result *IVectorView[*IMobileBroadbandCellUmts]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NeighboringCellsUmts, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellsInfo) Get_ServingCellsCdma() *IVectorView[*IMobileBroadbandCellCdma] {
	var _result *IVectorView[*IMobileBroadbandCellCdma]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServingCellsCdma, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellsInfo) Get_ServingCellsGsm() *IVectorView[*IMobileBroadbandCellGsm] {
	var _result *IVectorView[*IMobileBroadbandCellGsm]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServingCellsGsm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellsInfo) Get_ServingCellsLte() *IVectorView[*IMobileBroadbandCellLte] {
	var _result *IVectorView[*IMobileBroadbandCellLte]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServingCellsLte, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellsInfo) Get_ServingCellsTdscdma() *IVectorView[*IMobileBroadbandCellTdscdma] {
	var _result *IVectorView[*IMobileBroadbandCellTdscdma]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServingCellsTdscdma, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellsInfo) Get_ServingCellsUmts() *IVectorView[*IMobileBroadbandCellUmts] {
	var _result *IVectorView[*IMobileBroadbandCellUmts]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServingCellsUmts, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 66205912-B89F-4E12-BBB6-D5CF09A820CA
var IID_IMobileBroadbandCellsInfo2 = syscall.GUID{0x66205912, 0xB89F, 0x4E12,
	[8]byte{0xBB, 0xB6, 0xD5, 0xCF, 0x09, 0xA8, 0x20, 0xCA}}

type IMobileBroadbandCellsInfo2Interface interface {
	win32.IInspectableInterface
	Get_NeighboringCellsNR() *IVectorView[*IMobileBroadbandCellNR]
	Get_ServingCellsNR() *IVectorView[*IMobileBroadbandCellNR]
}

type IMobileBroadbandCellsInfo2Vtbl struct {
	win32.IInspectableVtbl
	Get_NeighboringCellsNR uintptr
	Get_ServingCellsNR     uintptr
}

type IMobileBroadbandCellsInfo2 struct {
	win32.IInspectable
}

func (this *IMobileBroadbandCellsInfo2) Vtbl() *IMobileBroadbandCellsInfo2Vtbl {
	return (*IMobileBroadbandCellsInfo2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandCellsInfo2) Get_NeighboringCellsNR() *IVectorView[*IMobileBroadbandCellNR] {
	var _result *IVectorView[*IMobileBroadbandCellNR]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NeighboringCellsNR, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandCellsInfo2) Get_ServingCellsNR() *IVectorView[*IMobileBroadbandCellNR] {
	var _result *IVectorView[*IMobileBroadbandCellNR]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServingCellsNR, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F718B184-C370-5FD4-A670-1846CB9BCE47
var IID_IMobileBroadbandCurrentSlotIndexChangedEventArgs = syscall.GUID{0xF718B184, 0xC370, 0x5FD4,
	[8]byte{0xA6, 0x70, 0x18, 0x46, 0xCB, 0x9B, 0xCE, 0x47}}

type IMobileBroadbandCurrentSlotIndexChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_CurrentSlotIndex() int32
}

type IMobileBroadbandCurrentSlotIndexChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_CurrentSlotIndex uintptr
}

type IMobileBroadbandCurrentSlotIndexChangedEventArgs struct {
	win32.IInspectable
}

func (this *IMobileBroadbandCurrentSlotIndexChangedEventArgs) Vtbl() *IMobileBroadbandCurrentSlotIndexChangedEventArgsVtbl {
	return (*IMobileBroadbandCurrentSlotIndexChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandCurrentSlotIndexChangedEventArgs) Get_CurrentSlotIndex() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentSlotIndex, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// E6D08168-E381-4C6E-9BE8-FE156969A446
var IID_IMobileBroadbandDeviceInformation = syscall.GUID{0xE6D08168, 0xE381, 0x4C6E,
	[8]byte{0x9B, 0xE8, 0xFE, 0x15, 0x69, 0x69, 0xA4, 0x46}}

type IMobileBroadbandDeviceInformationInterface interface {
	win32.IInspectableInterface
	Get_NetworkDeviceStatus() NetworkDeviceStatus
	Get_Manufacturer() string
	Get_Model() string
	Get_FirmwareInformation() string
	Get_CellularClass() CellularClass
	Get_DataClasses() DataClasses
	Get_CustomDataClass() string
	Get_MobileEquipmentId() string
	Get_TelephoneNumbers() *IVectorView[string]
	Get_SubscriberId() string
	Get_SimIccId() string
	Get_DeviceType() MobileBroadbandDeviceType
	Get_DeviceId() string
	Get_CurrentRadioState() MobileBroadbandRadioState
}

type IMobileBroadbandDeviceInformationVtbl struct {
	win32.IInspectableVtbl
	Get_NetworkDeviceStatus uintptr
	Get_Manufacturer        uintptr
	Get_Model               uintptr
	Get_FirmwareInformation uintptr
	Get_CellularClass       uintptr
	Get_DataClasses         uintptr
	Get_CustomDataClass     uintptr
	Get_MobileEquipmentId   uintptr
	Get_TelephoneNumbers    uintptr
	Get_SubscriberId        uintptr
	Get_SimIccId            uintptr
	Get_DeviceType          uintptr
	Get_DeviceId            uintptr
	Get_CurrentRadioState   uintptr
}

type IMobileBroadbandDeviceInformation struct {
	win32.IInspectable
}

func (this *IMobileBroadbandDeviceInformation) Vtbl() *IMobileBroadbandDeviceInformationVtbl {
	return (*IMobileBroadbandDeviceInformationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandDeviceInformation) Get_NetworkDeviceStatus() NetworkDeviceStatus {
	var _result NetworkDeviceStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkDeviceStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandDeviceInformation) Get_Manufacturer() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Manufacturer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandDeviceInformation) Get_Model() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Model, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandDeviceInformation) Get_FirmwareInformation() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FirmwareInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandDeviceInformation) Get_CellularClass() CellularClass {
	var _result CellularClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CellularClass, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandDeviceInformation) Get_DataClasses() DataClasses {
	var _result DataClasses
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DataClasses, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandDeviceInformation) Get_CustomDataClass() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CustomDataClass, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandDeviceInformation) Get_MobileEquipmentId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MobileEquipmentId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandDeviceInformation) Get_TelephoneNumbers() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TelephoneNumbers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandDeviceInformation) Get_SubscriberId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SubscriberId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandDeviceInformation) Get_SimIccId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SimIccId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandDeviceInformation) Get_DeviceType() MobileBroadbandDeviceType {
	var _result MobileBroadbandDeviceType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandDeviceInformation) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandDeviceInformation) Get_CurrentRadioState() MobileBroadbandRadioState {
	var _result MobileBroadbandRadioState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentRadioState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 2E467AF1-F932-4737-A722-03BA72370CB8
var IID_IMobileBroadbandDeviceInformation2 = syscall.GUID{0x2E467AF1, 0xF932, 0x4737,
	[8]byte{0xA7, 0x22, 0x03, 0xBA, 0x72, 0x37, 0x0C, 0xB8}}

type IMobileBroadbandDeviceInformation2Interface interface {
	win32.IInspectableInterface
	Get_PinManager() *IMobileBroadbandPinManager
	Get_Revision() string
	Get_SerialNumber() string
}

type IMobileBroadbandDeviceInformation2Vtbl struct {
	win32.IInspectableVtbl
	Get_PinManager   uintptr
	Get_Revision     uintptr
	Get_SerialNumber uintptr
}

type IMobileBroadbandDeviceInformation2 struct {
	win32.IInspectable
}

func (this *IMobileBroadbandDeviceInformation2) Vtbl() *IMobileBroadbandDeviceInformation2Vtbl {
	return (*IMobileBroadbandDeviceInformation2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandDeviceInformation2) Get_PinManager() *IMobileBroadbandPinManager {
	var _result *IMobileBroadbandPinManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PinManager, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandDeviceInformation2) Get_Revision() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Revision, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandDeviceInformation2) Get_SerialNumber() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SerialNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// E08BB4BD-5D30-4B5A-92CC-D54DF881D49E
var IID_IMobileBroadbandDeviceInformation3 = syscall.GUID{0xE08BB4BD, 0x5D30, 0x4B5A,
	[8]byte{0x92, 0xCC, 0xD5, 0x4D, 0xF8, 0x81, 0xD4, 0x9E}}

type IMobileBroadbandDeviceInformation3Interface interface {
	win32.IInspectableInterface
	Get_SimSpn() string
	Get_SimPnn() string
	Get_SimGid1() string
}

type IMobileBroadbandDeviceInformation3Vtbl struct {
	win32.IInspectableVtbl
	Get_SimSpn  uintptr
	Get_SimPnn  uintptr
	Get_SimGid1 uintptr
}

type IMobileBroadbandDeviceInformation3 struct {
	win32.IInspectable
}

func (this *IMobileBroadbandDeviceInformation3) Vtbl() *IMobileBroadbandDeviceInformation3Vtbl {
	return (*IMobileBroadbandDeviceInformation3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandDeviceInformation3) Get_SimSpn() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SimSpn, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandDeviceInformation3) Get_SimPnn() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SimPnn, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandDeviceInformation3) Get_SimGid1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SimGid1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 263F3152-7B9D-582C-B17C-F80A60B50031
var IID_IMobileBroadbandDeviceInformation4 = syscall.GUID{0x263F3152, 0x7B9D, 0x582C,
	[8]byte{0xB1, 0x7C, 0xF8, 0x0A, 0x60, 0xB5, 0x00, 0x31}}

type IMobileBroadbandDeviceInformation4Interface interface {
	win32.IInspectableInterface
	Get_SlotManager() *IMobileBroadbandSlotManager
}

type IMobileBroadbandDeviceInformation4Vtbl struct {
	win32.IInspectableVtbl
	Get_SlotManager uintptr
}

type IMobileBroadbandDeviceInformation4 struct {
	win32.IInspectable
}

func (this *IMobileBroadbandDeviceInformation4) Vtbl() *IMobileBroadbandDeviceInformation4Vtbl {
	return (*IMobileBroadbandDeviceInformation4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandDeviceInformation4) Get_SlotManager() *IMobileBroadbandSlotManager {
	var _result *IMobileBroadbandSlotManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SlotManager, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 22BE1A52-BD80-40AC-8E1F-2E07836A3DBD
var IID_IMobileBroadbandDeviceService = syscall.GUID{0x22BE1A52, 0xBD80, 0x40AC,
	[8]byte{0x8E, 0x1F, 0x2E, 0x07, 0x83, 0x6A, 0x3D, 0xBD}}

type IMobileBroadbandDeviceServiceInterface interface {
	win32.IInspectableInterface
	Get_DeviceServiceId() syscall.GUID
	Get_SupportedCommands() *IVectorView[uint32]
	OpenDataSession() *IMobileBroadbandDeviceServiceDataSession
	OpenCommandSession() *IMobileBroadbandDeviceServiceCommandSession
}

type IMobileBroadbandDeviceServiceVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceServiceId   uintptr
	Get_SupportedCommands uintptr
	OpenDataSession       uintptr
	OpenCommandSession    uintptr
}

type IMobileBroadbandDeviceService struct {
	win32.IInspectable
}

func (this *IMobileBroadbandDeviceService) Vtbl() *IMobileBroadbandDeviceServiceVtbl {
	return (*IMobileBroadbandDeviceServiceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandDeviceService) Get_DeviceServiceId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceServiceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandDeviceService) Get_SupportedCommands() *IVectorView[uint32] {
	var _result *IVectorView[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedCommands, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandDeviceService) OpenDataSession() *IMobileBroadbandDeviceServiceDataSession {
	var _result *IMobileBroadbandDeviceServiceDataSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenDataSession, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandDeviceService) OpenCommandSession() *IMobileBroadbandDeviceServiceCommandSession {
	var _result *IMobileBroadbandDeviceServiceCommandSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenCommandSession, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B0F46ABB-94D6-44B9-A538-F0810B645389
var IID_IMobileBroadbandDeviceServiceCommandResult = syscall.GUID{0xB0F46ABB, 0x94D6, 0x44B9,
	[8]byte{0xA5, 0x38, 0xF0, 0x81, 0x0B, 0x64, 0x53, 0x89}}

type IMobileBroadbandDeviceServiceCommandResultInterface interface {
	win32.IInspectableInterface
	Get_StatusCode() uint32
	Get_ResponseData() *IBuffer
}

type IMobileBroadbandDeviceServiceCommandResultVtbl struct {
	win32.IInspectableVtbl
	Get_StatusCode   uintptr
	Get_ResponseData uintptr
}

type IMobileBroadbandDeviceServiceCommandResult struct {
	win32.IInspectable
}

func (this *IMobileBroadbandDeviceServiceCommandResult) Vtbl() *IMobileBroadbandDeviceServiceCommandResultVtbl {
	return (*IMobileBroadbandDeviceServiceCommandResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandDeviceServiceCommandResult) Get_StatusCode() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StatusCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandDeviceServiceCommandResult) Get_ResponseData() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResponseData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FC098A45-913B-4914-B6C3-AE6304593E75
var IID_IMobileBroadbandDeviceServiceCommandSession = syscall.GUID{0xFC098A45, 0x913B, 0x4914,
	[8]byte{0xB6, 0xC3, 0xAE, 0x63, 0x04, 0x59, 0x3E, 0x75}}

type IMobileBroadbandDeviceServiceCommandSessionInterface interface {
	win32.IInspectableInterface
	SendQueryCommandAsync(commandId uint32, data *IBuffer) *IAsyncOperation[*IMobileBroadbandDeviceServiceCommandResult]
	SendSetCommandAsync(commandId uint32, data *IBuffer) *IAsyncOperation[*IMobileBroadbandDeviceServiceCommandResult]
	CloseSession()
}

type IMobileBroadbandDeviceServiceCommandSessionVtbl struct {
	win32.IInspectableVtbl
	SendQueryCommandAsync uintptr
	SendSetCommandAsync   uintptr
	CloseSession          uintptr
}

type IMobileBroadbandDeviceServiceCommandSession struct {
	win32.IInspectable
}

func (this *IMobileBroadbandDeviceServiceCommandSession) Vtbl() *IMobileBroadbandDeviceServiceCommandSessionVtbl {
	return (*IMobileBroadbandDeviceServiceCommandSessionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandDeviceServiceCommandSession) SendQueryCommandAsync(commandId uint32, data *IBuffer) *IAsyncOperation[*IMobileBroadbandDeviceServiceCommandResult] {
	var _result *IAsyncOperation[*IMobileBroadbandDeviceServiceCommandResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendQueryCommandAsync, uintptr(unsafe.Pointer(this)), uintptr(commandId), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandDeviceServiceCommandSession) SendSetCommandAsync(commandId uint32, data *IBuffer) *IAsyncOperation[*IMobileBroadbandDeviceServiceCommandResult] {
	var _result *IAsyncOperation[*IMobileBroadbandDeviceServiceCommandResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendSetCommandAsync, uintptr(unsafe.Pointer(this)), uintptr(commandId), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandDeviceServiceCommandSession) CloseSession() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CloseSession, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// B6AA13DE-1380-40E3-8618-73CBCA48138C
var IID_IMobileBroadbandDeviceServiceDataReceivedEventArgs = syscall.GUID{0xB6AA13DE, 0x1380, 0x40E3,
	[8]byte{0x86, 0x18, 0x73, 0xCB, 0xCA, 0x48, 0x13, 0x8C}}

type IMobileBroadbandDeviceServiceDataReceivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ReceivedData() *IBuffer
}

type IMobileBroadbandDeviceServiceDataReceivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ReceivedData uintptr
}

type IMobileBroadbandDeviceServiceDataReceivedEventArgs struct {
	win32.IInspectable
}

func (this *IMobileBroadbandDeviceServiceDataReceivedEventArgs) Vtbl() *IMobileBroadbandDeviceServiceDataReceivedEventArgsVtbl {
	return (*IMobileBroadbandDeviceServiceDataReceivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandDeviceServiceDataReceivedEventArgs) Get_ReceivedData() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReceivedData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DAD62333-8BCF-4289-8A37-045C2169486A
var IID_IMobileBroadbandDeviceServiceDataSession = syscall.GUID{0xDAD62333, 0x8BCF, 0x4289,
	[8]byte{0x8A, 0x37, 0x04, 0x5C, 0x21, 0x69, 0x48, 0x6A}}

type IMobileBroadbandDeviceServiceDataSessionInterface interface {
	win32.IInspectableInterface
	WriteDataAsync(value *IBuffer) *IAsyncAction
	CloseSession()
	Add_DataReceived(eventHandler TypedEventHandler[*IMobileBroadbandDeviceServiceDataSession, *IMobileBroadbandDeviceServiceDataReceivedEventArgs]) EventRegistrationToken
	Remove_DataReceived(eventCookie EventRegistrationToken)
}

type IMobileBroadbandDeviceServiceDataSessionVtbl struct {
	win32.IInspectableVtbl
	WriteDataAsync      uintptr
	CloseSession        uintptr
	Add_DataReceived    uintptr
	Remove_DataReceived uintptr
}

type IMobileBroadbandDeviceServiceDataSession struct {
	win32.IInspectable
}

func (this *IMobileBroadbandDeviceServiceDataSession) Vtbl() *IMobileBroadbandDeviceServiceDataSessionVtbl {
	return (*IMobileBroadbandDeviceServiceDataSessionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandDeviceServiceDataSession) WriteDataAsync(value *IBuffer) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteDataAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandDeviceServiceDataSession) CloseSession() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CloseSession, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IMobileBroadbandDeviceServiceDataSession) Add_DataReceived(eventHandler TypedEventHandler[*IMobileBroadbandDeviceServiceDataSession, *IMobileBroadbandDeviceServiceDataReceivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DataReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(eventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandDeviceServiceDataSession) Remove_DataReceived(eventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DataReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&eventCookie)))
	_ = _hr
}

// 53D69B5B-C4ED-45F0-803A-D9417A6D9846
var IID_IMobileBroadbandDeviceServiceInformation = syscall.GUID{0x53D69B5B, 0xC4ED, 0x45F0,
	[8]byte{0x80, 0x3A, 0xD9, 0x41, 0x7A, 0x6D, 0x98, 0x46}}

type IMobileBroadbandDeviceServiceInformationInterface interface {
	win32.IInspectableInterface
	Get_DeviceServiceId() syscall.GUID
	Get_IsDataReadSupported() bool
	Get_IsDataWriteSupported() bool
}

type IMobileBroadbandDeviceServiceInformationVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceServiceId      uintptr
	Get_IsDataReadSupported  uintptr
	Get_IsDataWriteSupported uintptr
}

type IMobileBroadbandDeviceServiceInformation struct {
	win32.IInspectable
}

func (this *IMobileBroadbandDeviceServiceInformation) Vtbl() *IMobileBroadbandDeviceServiceInformationVtbl {
	return (*IMobileBroadbandDeviceServiceInformationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandDeviceServiceInformation) Get_DeviceServiceId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceServiceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandDeviceServiceInformation) Get_IsDataReadSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDataReadSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandDeviceServiceInformation) Get_IsDataWriteSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDataWriteSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4A055B70-B9AE-4458-9241-A6A5FBF18A0C
var IID_IMobileBroadbandDeviceServiceTriggerDetails = syscall.GUID{0x4A055B70, 0xB9AE, 0x4458,
	[8]byte{0x92, 0x41, 0xA6, 0xA5, 0xFB, 0xF1, 0x8A, 0x0C}}

type IMobileBroadbandDeviceServiceTriggerDetailsInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_DeviceServiceId() syscall.GUID
	Get_ReceivedData() *IBuffer
}

type IMobileBroadbandDeviceServiceTriggerDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId        uintptr
	Get_DeviceServiceId uintptr
	Get_ReceivedData    uintptr
}

type IMobileBroadbandDeviceServiceTriggerDetails struct {
	win32.IInspectable
}

func (this *IMobileBroadbandDeviceServiceTriggerDetails) Vtbl() *IMobileBroadbandDeviceServiceTriggerDetailsVtbl {
	return (*IMobileBroadbandDeviceServiceTriggerDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandDeviceServiceTriggerDetails) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandDeviceServiceTriggerDetails) Get_DeviceServiceId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceServiceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandDeviceServiceTriggerDetails) Get_ReceivedData() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReceivedData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D83D5F16-336A-553F-94BB-0CD1A2FF0C81
var IID_IMobileBroadbandDeviceServiceTriggerDetails2 = syscall.GUID{0xD83D5F16, 0x336A, 0x553F,
	[8]byte{0x94, 0xBB, 0x0C, 0xD1, 0xA2, 0xFF, 0x0C, 0x81}}

type IMobileBroadbandDeviceServiceTriggerDetails2Interface interface {
	win32.IInspectableInterface
	Get_EventId() uint32
}

type IMobileBroadbandDeviceServiceTriggerDetails2Vtbl struct {
	win32.IInspectableVtbl
	Get_EventId uintptr
}

type IMobileBroadbandDeviceServiceTriggerDetails2 struct {
	win32.IInspectable
}

func (this *IMobileBroadbandDeviceServiceTriggerDetails2) Vtbl() *IMobileBroadbandDeviceServiceTriggerDetails2Vtbl {
	return (*IMobileBroadbandDeviceServiceTriggerDetails2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandDeviceServiceTriggerDetails2) Get_EventId() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EventId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D0356912-E9F9-4F67-A03D-43189A316BF1
var IID_IMobileBroadbandModem = syscall.GUID{0xD0356912, 0xE9F9, 0x4F67,
	[8]byte{0xA0, 0x3D, 0x43, 0x18, 0x9A, 0x31, 0x6B, 0xF1}}

type IMobileBroadbandModemInterface interface {
	win32.IInspectableInterface
	Get_CurrentAccount() *IMobileBroadbandAccount
	Get_DeviceInformation() *IMobileBroadbandDeviceInformation
	Get_MaxDeviceServiceCommandSizeInBytes() uint32
	Get_MaxDeviceServiceDataSizeInBytes() uint32
	Get_DeviceServices() *IVectorView[*IMobileBroadbandDeviceServiceInformation]
	GetDeviceService(deviceServiceId syscall.GUID) *IMobileBroadbandDeviceService
	Get_IsResetSupported() bool
	ResetAsync() *IAsyncAction
	GetCurrentConfigurationAsync() *IAsyncOperation[*IMobileBroadbandModemConfiguration]
	Get_CurrentNetwork() *IMobileBroadbandNetwork
}

type IMobileBroadbandModemVtbl struct {
	win32.IInspectableVtbl
	Get_CurrentAccount                     uintptr
	Get_DeviceInformation                  uintptr
	Get_MaxDeviceServiceCommandSizeInBytes uintptr
	Get_MaxDeviceServiceDataSizeInBytes    uintptr
	Get_DeviceServices                     uintptr
	GetDeviceService                       uintptr
	Get_IsResetSupported                   uintptr
	ResetAsync                             uintptr
	GetCurrentConfigurationAsync           uintptr
	Get_CurrentNetwork                     uintptr
}

type IMobileBroadbandModem struct {
	win32.IInspectable
}

func (this *IMobileBroadbandModem) Vtbl() *IMobileBroadbandModemVtbl {
	return (*IMobileBroadbandModemVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandModem) Get_CurrentAccount() *IMobileBroadbandAccount {
	var _result *IMobileBroadbandAccount
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentAccount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandModem) Get_DeviceInformation() *IMobileBroadbandDeviceInformation {
	var _result *IMobileBroadbandDeviceInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandModem) Get_MaxDeviceServiceCommandSizeInBytes() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxDeviceServiceCommandSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandModem) Get_MaxDeviceServiceDataSizeInBytes() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxDeviceServiceDataSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandModem) Get_DeviceServices() *IVectorView[*IMobileBroadbandDeviceServiceInformation] {
	var _result *IVectorView[*IMobileBroadbandDeviceServiceInformation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceServices, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandModem) GetDeviceService(deviceServiceId syscall.GUID) *IMobileBroadbandDeviceService {
	var _result *IMobileBroadbandDeviceService
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceService, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&deviceServiceId)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandModem) Get_IsResetSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsResetSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandModem) ResetAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ResetAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandModem) GetCurrentConfigurationAsync() *IAsyncOperation[*IMobileBroadbandModemConfiguration] {
	var _result *IAsyncOperation[*IMobileBroadbandModemConfiguration]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentConfigurationAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandModem) Get_CurrentNetwork() *IMobileBroadbandNetwork {
	var _result *IMobileBroadbandNetwork
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentNetwork, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 12862B28-B9EB-4EE2-BBE3-711F53EEA373
var IID_IMobileBroadbandModem2 = syscall.GUID{0x12862B28, 0xB9EB, 0x4EE2,
	[8]byte{0xBB, 0xE3, 0x71, 0x1F, 0x53, 0xEE, 0xA3, 0x73}}

type IMobileBroadbandModem2Interface interface {
	win32.IInspectableInterface
	GetIsPassthroughEnabledAsync() *IAsyncOperation[bool]
	SetIsPassthroughEnabledAsync(value bool) *IAsyncOperation[MobileBroadbandModemStatus]
}

type IMobileBroadbandModem2Vtbl struct {
	win32.IInspectableVtbl
	GetIsPassthroughEnabledAsync uintptr
	SetIsPassthroughEnabledAsync uintptr
}

type IMobileBroadbandModem2 struct {
	win32.IInspectable
}

func (this *IMobileBroadbandModem2) Vtbl() *IMobileBroadbandModem2Vtbl {
	return (*IMobileBroadbandModem2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandModem2) GetIsPassthroughEnabledAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetIsPassthroughEnabledAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandModem2) SetIsPassthroughEnabledAsync(value bool) *IAsyncOperation[MobileBroadbandModemStatus] {
	var _result *IAsyncOperation[MobileBroadbandModemStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetIsPassthroughEnabledAsync, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E9FEC6EA-2F34-4582-9102-C314D2A87EEC
var IID_IMobileBroadbandModem3 = syscall.GUID{0xE9FEC6EA, 0x2F34, 0x4582,
	[8]byte{0x91, 0x02, 0xC3, 0x14, 0xD2, 0xA8, 0x7E, 0xEC}}

type IMobileBroadbandModem3Interface interface {
	win32.IInspectableInterface
	TryGetPcoAsync() *IAsyncOperation[*IMobileBroadbandPco]
	Get_IsInEmergencyCallMode() bool
	Add_IsInEmergencyCallModeChanged(handler TypedEventHandler[*IMobileBroadbandModem, interface{}]) EventRegistrationToken
	Remove_IsInEmergencyCallModeChanged(token EventRegistrationToken)
}

type IMobileBroadbandModem3Vtbl struct {
	win32.IInspectableVtbl
	TryGetPcoAsync                      uintptr
	Get_IsInEmergencyCallMode           uintptr
	Add_IsInEmergencyCallModeChanged    uintptr
	Remove_IsInEmergencyCallModeChanged uintptr
}

type IMobileBroadbandModem3 struct {
	win32.IInspectable
}

func (this *IMobileBroadbandModem3) Vtbl() *IMobileBroadbandModem3Vtbl {
	return (*IMobileBroadbandModem3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandModem3) TryGetPcoAsync() *IAsyncOperation[*IMobileBroadbandPco] {
	var _result *IAsyncOperation[*IMobileBroadbandPco]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetPcoAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandModem3) Get_IsInEmergencyCallMode() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsInEmergencyCallMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandModem3) Add_IsInEmergencyCallModeChanged(handler TypedEventHandler[*IMobileBroadbandModem, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_IsInEmergencyCallModeChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandModem3) Remove_IsInEmergencyCallModeChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_IsInEmergencyCallModeChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 4A0398C2-91BE-412B-B569-586E9F0030D1
var IID_IMobileBroadbandModem4 = syscall.GUID{0x4A0398C2, 0x91BE, 0x412B,
	[8]byte{0xB5, 0x69, 0x58, 0x6E, 0x9F, 0x00, 0x30, 0xD1}}

type IMobileBroadbandModem4Interface interface {
	win32.IInspectableInterface
	SetIsPassthroughEnabledWithSlotIndexAsync(value bool, slotindex int32) *IAsyncOperation[MobileBroadbandModemStatus]
	GetIsPassthroughEnabledWithSlotIndexAsync(slotindex int32) *IAsyncOperation[bool]
	SetIsPassthroughEnabledWithSlotIndex(value bool, slotindex int32) MobileBroadbandModemStatus
	GetIsPassthroughEnabledWithSlotIndex(slotindex int32) bool
}

type IMobileBroadbandModem4Vtbl struct {
	win32.IInspectableVtbl
	SetIsPassthroughEnabledWithSlotIndexAsync uintptr
	GetIsPassthroughEnabledWithSlotIndexAsync uintptr
	SetIsPassthroughEnabledWithSlotIndex      uintptr
	GetIsPassthroughEnabledWithSlotIndex      uintptr
}

type IMobileBroadbandModem4 struct {
	win32.IInspectable
}

func (this *IMobileBroadbandModem4) Vtbl() *IMobileBroadbandModem4Vtbl {
	return (*IMobileBroadbandModem4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandModem4) SetIsPassthroughEnabledWithSlotIndexAsync(value bool, slotindex int32) *IAsyncOperation[MobileBroadbandModemStatus] {
	var _result *IAsyncOperation[MobileBroadbandModemStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetIsPassthroughEnabledWithSlotIndexAsync, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))), uintptr(slotindex), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandModem4) GetIsPassthroughEnabledWithSlotIndexAsync(slotindex int32) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetIsPassthroughEnabledWithSlotIndexAsync, uintptr(unsafe.Pointer(this)), uintptr(slotindex), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandModem4) SetIsPassthroughEnabledWithSlotIndex(value bool, slotindex int32) MobileBroadbandModemStatus {
	var _result MobileBroadbandModemStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetIsPassthroughEnabledWithSlotIndex, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))), uintptr(slotindex), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandModem4) GetIsPassthroughEnabledWithSlotIndex(slotindex int32) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetIsPassthroughEnabledWithSlotIndex, uintptr(unsafe.Pointer(this)), uintptr(slotindex), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// FCE035A3-D6CD-4320-B982-BE9D3EC7890F
var IID_IMobileBroadbandModemConfiguration = syscall.GUID{0xFCE035A3, 0xD6CD, 0x4320,
	[8]byte{0xB9, 0x82, 0xBE, 0x9D, 0x3E, 0xC7, 0x89, 0x0F}}

type IMobileBroadbandModemConfigurationInterface interface {
	win32.IInspectableInterface
	Get_Uicc() *IMobileBroadbandUicc
	Get_HomeProviderId() string
	Get_HomeProviderName() string
}

type IMobileBroadbandModemConfigurationVtbl struct {
	win32.IInspectableVtbl
	Get_Uicc             uintptr
	Get_HomeProviderId   uintptr
	Get_HomeProviderName uintptr
}

type IMobileBroadbandModemConfiguration struct {
	win32.IInspectable
}

func (this *IMobileBroadbandModemConfiguration) Vtbl() *IMobileBroadbandModemConfigurationVtbl {
	return (*IMobileBroadbandModemConfigurationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandModemConfiguration) Get_Uicc() *IMobileBroadbandUicc {
	var _result *IMobileBroadbandUicc
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Uicc, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandModemConfiguration) Get_HomeProviderId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HomeProviderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandModemConfiguration) Get_HomeProviderName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HomeProviderName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 320FF5C5-E460-42AE-AA51-69621E7A4477
var IID_IMobileBroadbandModemConfiguration2 = syscall.GUID{0x320FF5C5, 0xE460, 0x42AE,
	[8]byte{0xAA, 0x51, 0x69, 0x62, 0x1E, 0x7A, 0x44, 0x77}}

type IMobileBroadbandModemConfiguration2Interface interface {
	win32.IInspectableInterface
	Get_SarManager() *IMobileBroadbandSarManager
}

type IMobileBroadbandModemConfiguration2Vtbl struct {
	win32.IInspectableVtbl
	Get_SarManager uintptr
}

type IMobileBroadbandModemConfiguration2 struct {
	win32.IInspectable
}

func (this *IMobileBroadbandModemConfiguration2) Vtbl() *IMobileBroadbandModemConfiguration2Vtbl {
	return (*IMobileBroadbandModemConfiguration2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandModemConfiguration2) Get_SarManager() *IMobileBroadbandSarManager {
	var _result *IMobileBroadbandSarManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SarManager, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B5618FEC-E661-4330-9BB4-3480212EC354
var IID_IMobileBroadbandModemIsolation = syscall.GUID{0xB5618FEC, 0xE661, 0x4330,
	[8]byte{0x9B, 0xB4, 0x34, 0x80, 0x21, 0x2E, 0xC3, 0x54}}

type IMobileBroadbandModemIsolationInterface interface {
	win32.IInspectableInterface
	AddAllowedHost(host *IHostName)
	AddAllowedHostRange(first *IHostName, last *IHostName)
	ApplyConfigurationAsync() *IAsyncAction
	ClearConfigurationAsync() *IAsyncAction
}

type IMobileBroadbandModemIsolationVtbl struct {
	win32.IInspectableVtbl
	AddAllowedHost          uintptr
	AddAllowedHostRange     uintptr
	ApplyConfigurationAsync uintptr
	ClearConfigurationAsync uintptr
}

type IMobileBroadbandModemIsolation struct {
	win32.IInspectable
}

func (this *IMobileBroadbandModemIsolation) Vtbl() *IMobileBroadbandModemIsolationVtbl {
	return (*IMobileBroadbandModemIsolationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandModemIsolation) AddAllowedHost(host *IHostName) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddAllowedHost, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(host)))
	_ = _hr
}

func (this *IMobileBroadbandModemIsolation) AddAllowedHostRange(first *IHostName, last *IHostName) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddAllowedHostRange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(first)), uintptr(unsafe.Pointer(last)))
	_ = _hr
}

func (this *IMobileBroadbandModemIsolation) ApplyConfigurationAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ApplyConfigurationAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandModemIsolation) ClearConfigurationAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearConfigurationAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 21D7EC58-C2B1-4C2F-A030-72820A24ECD9
var IID_IMobileBroadbandModemIsolationFactory = syscall.GUID{0x21D7EC58, 0xC2B1, 0x4C2F,
	[8]byte{0xA0, 0x30, 0x72, 0x82, 0x0A, 0x24, 0xEC, 0xD9}}

type IMobileBroadbandModemIsolationFactoryInterface interface {
	win32.IInspectableInterface
	Create(modemDeviceId string, ruleGroupId string) *IMobileBroadbandModemIsolation
}

type IMobileBroadbandModemIsolationFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IMobileBroadbandModemIsolationFactory struct {
	win32.IInspectable
}

func (this *IMobileBroadbandModemIsolationFactory) Vtbl() *IMobileBroadbandModemIsolationFactoryVtbl {
	return (*IMobileBroadbandModemIsolationFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandModemIsolationFactory) Create(modemDeviceId string, ruleGroupId string) *IMobileBroadbandModemIsolation {
	var _result *IMobileBroadbandModemIsolation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(modemDeviceId).Ptr, NewHStr(ruleGroupId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F99ED637-D6F1-4A78-8CBC-6421A65063C8
var IID_IMobileBroadbandModemStatics = syscall.GUID{0xF99ED637, 0xD6F1, 0x4A78,
	[8]byte{0x8C, 0xBC, 0x64, 0x21, 0xA6, 0x50, 0x63, 0xC8}}

type IMobileBroadbandModemStaticsInterface interface {
	win32.IInspectableInterface
	GetDeviceSelector() string
	FromId(deviceId string) *IMobileBroadbandModem
	GetDefault() *IMobileBroadbandModem
}

type IMobileBroadbandModemStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelector uintptr
	FromId            uintptr
	GetDefault        uintptr
}

type IMobileBroadbandModemStatics struct {
	win32.IInspectable
}

func (this *IMobileBroadbandModemStatics) Vtbl() *IMobileBroadbandModemStaticsVtbl {
	return (*IMobileBroadbandModemStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandModemStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandModemStatics) FromId(deviceId string) *IMobileBroadbandModem {
	var _result *IMobileBroadbandModem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromId, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandModemStatics) GetDefault() *IMobileBroadbandModem {
	var _result *IMobileBroadbandModem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CB63928C-0309-4CB6-A8C1-6A5A3C8E1FF6
var IID_IMobileBroadbandNetwork = syscall.GUID{0xCB63928C, 0x0309, 0x4CB6,
	[8]byte{0xA8, 0xC1, 0x6A, 0x5A, 0x3C, 0x8E, 0x1F, 0xF6}}

type IMobileBroadbandNetworkInterface interface {
	win32.IInspectableInterface
	Get_NetworkAdapter() *INetworkAdapter
	Get_NetworkRegistrationState() NetworkRegistrationState
	Get_RegistrationNetworkError() uint32
	Get_PacketAttachNetworkError() uint32
	Get_ActivationNetworkError() uint32
	Get_AccessPointName() string
	Get_RegisteredDataClass() DataClasses
	Get_RegisteredProviderId() string
	Get_RegisteredProviderName() string
	ShowConnectionUI()
}

type IMobileBroadbandNetworkVtbl struct {
	win32.IInspectableVtbl
	Get_NetworkAdapter           uintptr
	Get_NetworkRegistrationState uintptr
	Get_RegistrationNetworkError uintptr
	Get_PacketAttachNetworkError uintptr
	Get_ActivationNetworkError   uintptr
	Get_AccessPointName          uintptr
	Get_RegisteredDataClass      uintptr
	Get_RegisteredProviderId     uintptr
	Get_RegisteredProviderName   uintptr
	ShowConnectionUI             uintptr
}

type IMobileBroadbandNetwork struct {
	win32.IInspectable
}

func (this *IMobileBroadbandNetwork) Vtbl() *IMobileBroadbandNetworkVtbl {
	return (*IMobileBroadbandNetworkVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandNetwork) Get_NetworkAdapter() *INetworkAdapter {
	var _result *INetworkAdapter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkAdapter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandNetwork) Get_NetworkRegistrationState() NetworkRegistrationState {
	var _result NetworkRegistrationState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkRegistrationState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandNetwork) Get_RegistrationNetworkError() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RegistrationNetworkError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandNetwork) Get_PacketAttachNetworkError() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PacketAttachNetworkError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandNetwork) Get_ActivationNetworkError() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ActivationNetworkError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandNetwork) Get_AccessPointName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AccessPointName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandNetwork) Get_RegisteredDataClass() DataClasses {
	var _result DataClasses
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RegisteredDataClass, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandNetwork) Get_RegisteredProviderId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RegisteredProviderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandNetwork) Get_RegisteredProviderName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RegisteredProviderName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandNetwork) ShowConnectionUI() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowConnectionUI, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 5A55DB22-62F7-4BDD-BA1D-477441960BA0
var IID_IMobileBroadbandNetwork2 = syscall.GUID{0x5A55DB22, 0x62F7, 0x4BDD,
	[8]byte{0xBA, 0x1D, 0x47, 0x74, 0x41, 0x96, 0x0B, 0xA0}}

type IMobileBroadbandNetwork2Interface interface {
	win32.IInspectableInterface
	GetVoiceCallSupportAsync() *IAsyncOperation[bool]
	Get_RegistrationUiccApps() *IVectorView[*IMobileBroadbandUiccApp]
}

type IMobileBroadbandNetwork2Vtbl struct {
	win32.IInspectableVtbl
	GetVoiceCallSupportAsync uintptr
	Get_RegistrationUiccApps uintptr
}

type IMobileBroadbandNetwork2 struct {
	win32.IInspectable
}

func (this *IMobileBroadbandNetwork2) Vtbl() *IMobileBroadbandNetwork2Vtbl {
	return (*IMobileBroadbandNetwork2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandNetwork2) GetVoiceCallSupportAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetVoiceCallSupportAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandNetwork2) Get_RegistrationUiccApps() *IVectorView[*IMobileBroadbandUiccApp] {
	var _result *IVectorView[*IMobileBroadbandUiccApp]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RegistrationUiccApps, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 33670A8A-C7EF-444C-AB6C-DF7EF7A390FE
var IID_IMobileBroadbandNetwork3 = syscall.GUID{0x33670A8A, 0xC7EF, 0x444C,
	[8]byte{0xAB, 0x6C, 0xDF, 0x7E, 0xF7, 0xA3, 0x90, 0xFE}}

type IMobileBroadbandNetwork3Interface interface {
	win32.IInspectableInterface
	GetCellsInfoAsync() *IAsyncOperation[*IMobileBroadbandCellsInfo]
}

type IMobileBroadbandNetwork3Vtbl struct {
	win32.IInspectableVtbl
	GetCellsInfoAsync uintptr
}

type IMobileBroadbandNetwork3 struct {
	win32.IInspectable
}

func (this *IMobileBroadbandNetwork3) Vtbl() *IMobileBroadbandNetwork3Vtbl {
	return (*IMobileBroadbandNetwork3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandNetwork3) GetCellsInfoAsync() *IAsyncOperation[*IMobileBroadbandCellsInfo] {
	var _result *IAsyncOperation[*IMobileBroadbandCellsInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCellsInfoAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BEAF94E1-960F-49B4-A08D-7D85E968C7EC
var IID_IMobileBroadbandNetworkRegistrationStateChange = syscall.GUID{0xBEAF94E1, 0x960F, 0x49B4,
	[8]byte{0xA0, 0x8D, 0x7D, 0x85, 0xE9, 0x68, 0xC7, 0xEC}}

type IMobileBroadbandNetworkRegistrationStateChangeInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_Network() *IMobileBroadbandNetwork
}

type IMobileBroadbandNetworkRegistrationStateChangeVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId uintptr
	Get_Network  uintptr
}

type IMobileBroadbandNetworkRegistrationStateChange struct {
	win32.IInspectable
}

func (this *IMobileBroadbandNetworkRegistrationStateChange) Vtbl() *IMobileBroadbandNetworkRegistrationStateChangeVtbl {
	return (*IMobileBroadbandNetworkRegistrationStateChangeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandNetworkRegistrationStateChange) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandNetworkRegistrationStateChange) Get_Network() *IMobileBroadbandNetwork {
	var _result *IMobileBroadbandNetwork
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Network, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 89135CFF-28B8-46AA-B137-1C4B0F21EDFE
var IID_IMobileBroadbandNetworkRegistrationStateChangeTriggerDetails = syscall.GUID{0x89135CFF, 0x28B8, 0x46AA,
	[8]byte{0xB1, 0x37, 0x1C, 0x4B, 0x0F, 0x21, 0xED, 0xFE}}

type IMobileBroadbandNetworkRegistrationStateChangeTriggerDetailsInterface interface {
	win32.IInspectableInterface
	Get_NetworkRegistrationStateChanges() *IVectorView[*IMobileBroadbandNetworkRegistrationStateChange]
}

type IMobileBroadbandNetworkRegistrationStateChangeTriggerDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_NetworkRegistrationStateChanges uintptr
}

type IMobileBroadbandNetworkRegistrationStateChangeTriggerDetails struct {
	win32.IInspectable
}

func (this *IMobileBroadbandNetworkRegistrationStateChangeTriggerDetails) Vtbl() *IMobileBroadbandNetworkRegistrationStateChangeTriggerDetailsVtbl {
	return (*IMobileBroadbandNetworkRegistrationStateChangeTriggerDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandNetworkRegistrationStateChangeTriggerDetails) Get_NetworkRegistrationStateChanges() *IVectorView[*IMobileBroadbandNetworkRegistrationStateChange] {
	var _result *IVectorView[*IMobileBroadbandNetworkRegistrationStateChange]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkRegistrationStateChanges, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D4E4FCBE-E3A3-43C5-A87B-6C86D229D7FA
var IID_IMobileBroadbandPco = syscall.GUID{0xD4E4FCBE, 0xE3A3, 0x43C5,
	[8]byte{0xA8, 0x7B, 0x6C, 0x86, 0xD2, 0x29, 0xD7, 0xFA}}

type IMobileBroadbandPcoInterface interface {
	win32.IInspectableInterface
	Get_Data() *IBuffer
	Get_IsComplete() bool
	Get_DeviceId() string
}

type IMobileBroadbandPcoVtbl struct {
	win32.IInspectableVtbl
	Get_Data       uintptr
	Get_IsComplete uintptr
	Get_DeviceId   uintptr
}

type IMobileBroadbandPco struct {
	win32.IInspectable
}

func (this *IMobileBroadbandPco) Vtbl() *IMobileBroadbandPcoVtbl {
	return (*IMobileBroadbandPcoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandPco) Get_Data() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandPco) Get_IsComplete() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsComplete, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandPco) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 263F5114-64E0-4493-909B-2D14A01962B1
var IID_IMobileBroadbandPcoDataChangeTriggerDetails = syscall.GUID{0x263F5114, 0x64E0, 0x4493,
	[8]byte{0x90, 0x9B, 0x2D, 0x14, 0xA0, 0x19, 0x62, 0xB1}}

type IMobileBroadbandPcoDataChangeTriggerDetailsInterface interface {
	win32.IInspectableInterface
	Get_UpdatedData() *IMobileBroadbandPco
}

type IMobileBroadbandPcoDataChangeTriggerDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_UpdatedData uintptr
}

type IMobileBroadbandPcoDataChangeTriggerDetails struct {
	win32.IInspectable
}

func (this *IMobileBroadbandPcoDataChangeTriggerDetails) Vtbl() *IMobileBroadbandPcoDataChangeTriggerDetailsVtbl {
	return (*IMobileBroadbandPcoDataChangeTriggerDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandPcoDataChangeTriggerDetails) Get_UpdatedData() *IMobileBroadbandPco {
	var _result *IMobileBroadbandPco
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UpdatedData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E661D709-E779-45BF-8281-75323DF9E321
var IID_IMobileBroadbandPin = syscall.GUID{0xE661D709, 0xE779, 0x45BF,
	[8]byte{0x82, 0x81, 0x75, 0x32, 0x3D, 0xF9, 0xE3, 0x21}}

type IMobileBroadbandPinInterface interface {
	win32.IInspectableInterface
	Get_Type() MobileBroadbandPinType
	Get_LockState() MobileBroadbandPinLockState
	Get_Format() MobileBroadbandPinFormat
	Get_Enabled() bool
	Get_MaxLength() uint32
	Get_MinLength() uint32
	Get_AttemptsRemaining() uint32
	EnableAsync(currentPin string) *IAsyncOperation[*IMobileBroadbandPinOperationResult]
	DisableAsync(currentPin string) *IAsyncOperation[*IMobileBroadbandPinOperationResult]
	EnterAsync(currentPin string) *IAsyncOperation[*IMobileBroadbandPinOperationResult]
	ChangeAsync(currentPin string, newPin string) *IAsyncOperation[*IMobileBroadbandPinOperationResult]
	UnblockAsync(pinUnblockKey string, newPin string) *IAsyncOperation[*IMobileBroadbandPinOperationResult]
}

type IMobileBroadbandPinVtbl struct {
	win32.IInspectableVtbl
	Get_Type              uintptr
	Get_LockState         uintptr
	Get_Format            uintptr
	Get_Enabled           uintptr
	Get_MaxLength         uintptr
	Get_MinLength         uintptr
	Get_AttemptsRemaining uintptr
	EnableAsync           uintptr
	DisableAsync          uintptr
	EnterAsync            uintptr
	ChangeAsync           uintptr
	UnblockAsync          uintptr
}

type IMobileBroadbandPin struct {
	win32.IInspectable
}

func (this *IMobileBroadbandPin) Vtbl() *IMobileBroadbandPinVtbl {
	return (*IMobileBroadbandPinVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandPin) Get_Type() MobileBroadbandPinType {
	var _result MobileBroadbandPinType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Type, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandPin) Get_LockState() MobileBroadbandPinLockState {
	var _result MobileBroadbandPinLockState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LockState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandPin) Get_Format() MobileBroadbandPinFormat {
	var _result MobileBroadbandPinFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Format, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandPin) Get_Enabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Enabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandPin) Get_MaxLength() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandPin) Get_MinLength() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandPin) Get_AttemptsRemaining() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AttemptsRemaining, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandPin) EnableAsync(currentPin string) *IAsyncOperation[*IMobileBroadbandPinOperationResult] {
	var _result *IAsyncOperation[*IMobileBroadbandPinOperationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnableAsync, uintptr(unsafe.Pointer(this)), NewHStr(currentPin).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandPin) DisableAsync(currentPin string) *IAsyncOperation[*IMobileBroadbandPinOperationResult] {
	var _result *IAsyncOperation[*IMobileBroadbandPinOperationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DisableAsync, uintptr(unsafe.Pointer(this)), NewHStr(currentPin).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandPin) EnterAsync(currentPin string) *IAsyncOperation[*IMobileBroadbandPinOperationResult] {
	var _result *IAsyncOperation[*IMobileBroadbandPinOperationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnterAsync, uintptr(unsafe.Pointer(this)), NewHStr(currentPin).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandPin) ChangeAsync(currentPin string, newPin string) *IAsyncOperation[*IMobileBroadbandPinOperationResult] {
	var _result *IAsyncOperation[*IMobileBroadbandPinOperationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ChangeAsync, uintptr(unsafe.Pointer(this)), NewHStr(currentPin).Ptr, NewHStr(newPin).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandPin) UnblockAsync(pinUnblockKey string, newPin string) *IAsyncOperation[*IMobileBroadbandPinOperationResult] {
	var _result *IAsyncOperation[*IMobileBroadbandPinOperationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UnblockAsync, uintptr(unsafe.Pointer(this)), NewHStr(pinUnblockKey).Ptr, NewHStr(newPin).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BE16673E-1F04-4F95-8B90-E7F559DDE7E5
var IID_IMobileBroadbandPinLockStateChange = syscall.GUID{0xBE16673E, 0x1F04, 0x4F95,
	[8]byte{0x8B, 0x90, 0xE7, 0xF5, 0x59, 0xDD, 0xE7, 0xE5}}

type IMobileBroadbandPinLockStateChangeInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_PinType() MobileBroadbandPinType
	Get_PinLockState() MobileBroadbandPinLockState
}

type IMobileBroadbandPinLockStateChangeVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId     uintptr
	Get_PinType      uintptr
	Get_PinLockState uintptr
}

type IMobileBroadbandPinLockStateChange struct {
	win32.IInspectable
}

func (this *IMobileBroadbandPinLockStateChange) Vtbl() *IMobileBroadbandPinLockStateChangeVtbl {
	return (*IMobileBroadbandPinLockStateChangeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandPinLockStateChange) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandPinLockStateChange) Get_PinType() MobileBroadbandPinType {
	var _result MobileBroadbandPinType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PinType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandPinLockStateChange) Get_PinLockState() MobileBroadbandPinLockState {
	var _result MobileBroadbandPinLockState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PinLockState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D338C091-3E91-4D38-9036-AEE83A6E79AD
var IID_IMobileBroadbandPinLockStateChangeTriggerDetails = syscall.GUID{0xD338C091, 0x3E91, 0x4D38,
	[8]byte{0x90, 0x36, 0xAE, 0xE8, 0x3A, 0x6E, 0x79, 0xAD}}

type IMobileBroadbandPinLockStateChangeTriggerDetailsInterface interface {
	win32.IInspectableInterface
	Get_PinLockStateChanges() *IVectorView[*IMobileBroadbandPinLockStateChange]
}

type IMobileBroadbandPinLockStateChangeTriggerDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_PinLockStateChanges uintptr
}

type IMobileBroadbandPinLockStateChangeTriggerDetails struct {
	win32.IInspectable
}

func (this *IMobileBroadbandPinLockStateChangeTriggerDetails) Vtbl() *IMobileBroadbandPinLockStateChangeTriggerDetailsVtbl {
	return (*IMobileBroadbandPinLockStateChangeTriggerDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandPinLockStateChangeTriggerDetails) Get_PinLockStateChanges() *IVectorView[*IMobileBroadbandPinLockStateChange] {
	var _result *IVectorView[*IMobileBroadbandPinLockStateChange]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PinLockStateChanges, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 83567EDD-6E1F-4B9B-A413-2B1F50CC36DF
var IID_IMobileBroadbandPinManager = syscall.GUID{0x83567EDD, 0x6E1F, 0x4B9B,
	[8]byte{0xA4, 0x13, 0x2B, 0x1F, 0x50, 0xCC, 0x36, 0xDF}}

type IMobileBroadbandPinManagerInterface interface {
	win32.IInspectableInterface
	Get_SupportedPins() *IVectorView[MobileBroadbandPinType]
	GetPin(pinType MobileBroadbandPinType) *IMobileBroadbandPin
}

type IMobileBroadbandPinManagerVtbl struct {
	win32.IInspectableVtbl
	Get_SupportedPins uintptr
	GetPin            uintptr
}

type IMobileBroadbandPinManager struct {
	win32.IInspectable
}

func (this *IMobileBroadbandPinManager) Vtbl() *IMobileBroadbandPinManagerVtbl {
	return (*IMobileBroadbandPinManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandPinManager) Get_SupportedPins() *IVectorView[MobileBroadbandPinType] {
	var _result *IVectorView[MobileBroadbandPinType]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedPins, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandPinManager) GetPin(pinType MobileBroadbandPinType) *IMobileBroadbandPin {
	var _result *IMobileBroadbandPin
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPin, uintptr(unsafe.Pointer(this)), uintptr(pinType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 11DDDC32-31E7-49F5-B663-123D3BEF0362
var IID_IMobileBroadbandPinOperationResult = syscall.GUID{0x11DDDC32, 0x31E7, 0x49F5,
	[8]byte{0xB6, 0x63, 0x12, 0x3D, 0x3B, 0xEF, 0x03, 0x62}}

type IMobileBroadbandPinOperationResultInterface interface {
	win32.IInspectableInterface
	Get_IsSuccessful() bool
	Get_AttemptsRemaining() uint32
}

type IMobileBroadbandPinOperationResultVtbl struct {
	win32.IInspectableVtbl
	Get_IsSuccessful      uintptr
	Get_AttemptsRemaining uintptr
}

type IMobileBroadbandPinOperationResult struct {
	win32.IInspectable
}

func (this *IMobileBroadbandPinOperationResult) Vtbl() *IMobileBroadbandPinOperationResultVtbl {
	return (*IMobileBroadbandPinOperationResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandPinOperationResult) Get_IsSuccessful() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSuccessful, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandPinOperationResult) Get_AttemptsRemaining() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AttemptsRemaining, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// B054A561-9833-4AED-9717-4348B21A24B3
var IID_IMobileBroadbandRadioStateChange = syscall.GUID{0xB054A561, 0x9833, 0x4AED,
	[8]byte{0x97, 0x17, 0x43, 0x48, 0xB2, 0x1A, 0x24, 0xB3}}

type IMobileBroadbandRadioStateChangeInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_RadioState() MobileBroadbandRadioState
}

type IMobileBroadbandRadioStateChangeVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId   uintptr
	Get_RadioState uintptr
}

type IMobileBroadbandRadioStateChange struct {
	win32.IInspectable
}

func (this *IMobileBroadbandRadioStateChange) Vtbl() *IMobileBroadbandRadioStateChangeVtbl {
	return (*IMobileBroadbandRadioStateChangeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandRadioStateChange) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandRadioStateChange) Get_RadioState() MobileBroadbandRadioState {
	var _result MobileBroadbandRadioState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RadioState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 71301ACE-093C-42C6-B0DB-AD1F75A65445
var IID_IMobileBroadbandRadioStateChangeTriggerDetails = syscall.GUID{0x71301ACE, 0x093C, 0x42C6,
	[8]byte{0xB0, 0xDB, 0xAD, 0x1F, 0x75, 0xA6, 0x54, 0x45}}

type IMobileBroadbandRadioStateChangeTriggerDetailsInterface interface {
	win32.IInspectableInterface
	Get_RadioStateChanges() *IVectorView[*IMobileBroadbandRadioStateChange]
}

type IMobileBroadbandRadioStateChangeTriggerDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_RadioStateChanges uintptr
}

type IMobileBroadbandRadioStateChangeTriggerDetails struct {
	win32.IInspectable
}

func (this *IMobileBroadbandRadioStateChangeTriggerDetails) Vtbl() *IMobileBroadbandRadioStateChangeTriggerDetailsVtbl {
	return (*IMobileBroadbandRadioStateChangeTriggerDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandRadioStateChangeTriggerDetails) Get_RadioStateChanges() *IVectorView[*IMobileBroadbandRadioStateChange] {
	var _result *IVectorView[*IMobileBroadbandRadioStateChange]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RadioStateChanges, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E5B26833-967E-40C9-A485-19C0DD209E22
var IID_IMobileBroadbandSarManager = syscall.GUID{0xE5B26833, 0x967E, 0x40C9,
	[8]byte{0xA4, 0x85, 0x19, 0xC0, 0xDD, 0x20, 0x9E, 0x22}}

type IMobileBroadbandSarManagerInterface interface {
	win32.IInspectableInterface
	Get_IsBackoffEnabled() bool
	Get_IsWiFiHardwareIntegrated() bool
	Get_IsSarControlledByHardware() bool
	Get_Antennas() *IVectorView[*IMobileBroadbandAntennaSar]
	Get_HysteresisTimerPeriod() TimeSpan
	Add_TransmissionStateChanged(handler TypedEventHandler[*IMobileBroadbandSarManager, *IMobileBroadbandTransmissionStateChangedEventArgs]) EventRegistrationToken
	Remove_TransmissionStateChanged(token EventRegistrationToken)
	EnableBackoffAsync() *IAsyncAction
	DisableBackoffAsync() *IAsyncAction
	SetConfigurationAsync(antennas *IIterable[*IMobileBroadbandAntennaSar]) *IAsyncAction
	RevertSarToHardwareControlAsync() *IAsyncAction
	SetTransmissionStateChangedHysteresisAsync(timerPeriod TimeSpan) *IAsyncAction
	GetIsTransmittingAsync() *IAsyncOperation[bool]
	StartTransmissionStateMonitoring()
	StopTransmissionStateMonitoring()
}

type IMobileBroadbandSarManagerVtbl struct {
	win32.IInspectableVtbl
	Get_IsBackoffEnabled                       uintptr
	Get_IsWiFiHardwareIntegrated               uintptr
	Get_IsSarControlledByHardware              uintptr
	Get_Antennas                               uintptr
	Get_HysteresisTimerPeriod                  uintptr
	Add_TransmissionStateChanged               uintptr
	Remove_TransmissionStateChanged            uintptr
	EnableBackoffAsync                         uintptr
	DisableBackoffAsync                        uintptr
	SetConfigurationAsync                      uintptr
	RevertSarToHardwareControlAsync            uintptr
	SetTransmissionStateChangedHysteresisAsync uintptr
	GetIsTransmittingAsync                     uintptr
	StartTransmissionStateMonitoring           uintptr
	StopTransmissionStateMonitoring            uintptr
}

type IMobileBroadbandSarManager struct {
	win32.IInspectable
}

func (this *IMobileBroadbandSarManager) Vtbl() *IMobileBroadbandSarManagerVtbl {
	return (*IMobileBroadbandSarManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandSarManager) Get_IsBackoffEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsBackoffEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandSarManager) Get_IsWiFiHardwareIntegrated() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsWiFiHardwareIntegrated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandSarManager) Get_IsSarControlledByHardware() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSarControlledByHardware, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandSarManager) Get_Antennas() *IVectorView[*IMobileBroadbandAntennaSar] {
	var _result *IVectorView[*IMobileBroadbandAntennaSar]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Antennas, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandSarManager) Get_HysteresisTimerPeriod() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HysteresisTimerPeriod, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandSarManager) Add_TransmissionStateChanged(handler TypedEventHandler[*IMobileBroadbandSarManager, *IMobileBroadbandTransmissionStateChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_TransmissionStateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandSarManager) Remove_TransmissionStateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_TransmissionStateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMobileBroadbandSarManager) EnableBackoffAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnableBackoffAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandSarManager) DisableBackoffAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DisableBackoffAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandSarManager) SetConfigurationAsync(antennas *IIterable[*IMobileBroadbandAntennaSar]) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetConfigurationAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(antennas)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandSarManager) RevertSarToHardwareControlAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RevertSarToHardwareControlAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandSarManager) SetTransmissionStateChangedHysteresisAsync(timerPeriod TimeSpan) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetTransmissionStateChangedHysteresisAsync, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&timerPeriod)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandSarManager) GetIsTransmittingAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetIsTransmittingAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandSarManager) StartTransmissionStateMonitoring() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartTransmissionStateMonitoring, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IMobileBroadbandSarManager) StopTransmissionStateMonitoring() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopTransmissionStateMonitoring, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// BD350B32-882E-542A-B17D-0BB1B49BAE9E
var IID_IMobileBroadbandSlotInfo = syscall.GUID{0xBD350B32, 0x882E, 0x542A,
	[8]byte{0xB1, 0x7D, 0x0B, 0xB1, 0xB4, 0x9B, 0xAE, 0x9E}}

type IMobileBroadbandSlotInfoInterface interface {
	win32.IInspectableInterface
	Get_Index() int32
	Get_State() MobileBroadbandSlotState
}

type IMobileBroadbandSlotInfoVtbl struct {
	win32.IInspectableVtbl
	Get_Index uintptr
	Get_State uintptr
}

type IMobileBroadbandSlotInfo struct {
	win32.IInspectable
}

func (this *IMobileBroadbandSlotInfo) Vtbl() *IMobileBroadbandSlotInfoVtbl {
	return (*IMobileBroadbandSlotInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandSlotInfo) Get_Index() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Index, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandSlotInfo) Get_State() MobileBroadbandSlotState {
	var _result MobileBroadbandSlotState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 393CB039-CA44-524C-822D-83A3620F0EFC
var IID_IMobileBroadbandSlotInfo2 = syscall.GUID{0x393CB039, 0xCA44, 0x524C,
	[8]byte{0x82, 0x2D, 0x83, 0xA3, 0x62, 0x0F, 0x0E, 0xFC}}

type IMobileBroadbandSlotInfo2Interface interface {
	win32.IInspectableInterface
	Get_IccId() string
}

type IMobileBroadbandSlotInfo2Vtbl struct {
	win32.IInspectableVtbl
	Get_IccId uintptr
}

type IMobileBroadbandSlotInfo2 struct {
	win32.IInspectable
}

func (this *IMobileBroadbandSlotInfo2) Vtbl() *IMobileBroadbandSlotInfo2Vtbl {
	return (*IMobileBroadbandSlotInfo2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandSlotInfo2) Get_IccId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IccId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 3158839F-950C-54CE-A48D-BA4529B48F0F
var IID_IMobileBroadbandSlotInfoChangedEventArgs = syscall.GUID{0x3158839F, 0x950C, 0x54CE,
	[8]byte{0xA4, 0x8D, 0xBA, 0x45, 0x29, 0xB4, 0x8F, 0x0F}}

type IMobileBroadbandSlotInfoChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_SlotInfo() *IMobileBroadbandSlotInfo
}

type IMobileBroadbandSlotInfoChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_SlotInfo uintptr
}

type IMobileBroadbandSlotInfoChangedEventArgs struct {
	win32.IInspectable
}

func (this *IMobileBroadbandSlotInfoChangedEventArgs) Vtbl() *IMobileBroadbandSlotInfoChangedEventArgsVtbl {
	return (*IMobileBroadbandSlotInfoChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandSlotInfoChangedEventArgs) Get_SlotInfo() *IMobileBroadbandSlotInfo {
	var _result *IMobileBroadbandSlotInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SlotInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EBA07CD6-2019-5F81-A294-CC364A11D0B2
var IID_IMobileBroadbandSlotManager = syscall.GUID{0xEBA07CD6, 0x2019, 0x5F81,
	[8]byte{0xA2, 0x94, 0xCC, 0x36, 0x4A, 0x11, 0xD0, 0xB2}}

type IMobileBroadbandSlotManagerInterface interface {
	win32.IInspectableInterface
	Get_SlotInfos() *IVectorView[*IMobileBroadbandSlotInfo]
	Get_CurrentSlotIndex() int32
	SetCurrentSlot(slotIndex int32) MobileBroadbandModemStatus
	SetCurrentSlotAsync(slotIndex int32) *IAsyncOperation[MobileBroadbandModemStatus]
	Add_SlotInfoChanged(handler TypedEventHandler[*IMobileBroadbandSlotManager, *IMobileBroadbandSlotInfoChangedEventArgs]) EventRegistrationToken
	Remove_SlotInfoChanged(token EventRegistrationToken)
	Add_CurrentSlotIndexChanged(handler TypedEventHandler[*IMobileBroadbandSlotManager, *IMobileBroadbandCurrentSlotIndexChangedEventArgs]) EventRegistrationToken
	Remove_CurrentSlotIndexChanged(token EventRegistrationToken)
}

type IMobileBroadbandSlotManagerVtbl struct {
	win32.IInspectableVtbl
	Get_SlotInfos                  uintptr
	Get_CurrentSlotIndex           uintptr
	SetCurrentSlot                 uintptr
	SetCurrentSlotAsync            uintptr
	Add_SlotInfoChanged            uintptr
	Remove_SlotInfoChanged         uintptr
	Add_CurrentSlotIndexChanged    uintptr
	Remove_CurrentSlotIndexChanged uintptr
}

type IMobileBroadbandSlotManager struct {
	win32.IInspectable
}

func (this *IMobileBroadbandSlotManager) Vtbl() *IMobileBroadbandSlotManagerVtbl {
	return (*IMobileBroadbandSlotManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandSlotManager) Get_SlotInfos() *IVectorView[*IMobileBroadbandSlotInfo] {
	var _result *IVectorView[*IMobileBroadbandSlotInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SlotInfos, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandSlotManager) Get_CurrentSlotIndex() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentSlotIndex, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandSlotManager) SetCurrentSlot(slotIndex int32) MobileBroadbandModemStatus {
	var _result MobileBroadbandModemStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetCurrentSlot, uintptr(unsafe.Pointer(this)), uintptr(slotIndex), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandSlotManager) SetCurrentSlotAsync(slotIndex int32) *IAsyncOperation[MobileBroadbandModemStatus] {
	var _result *IAsyncOperation[MobileBroadbandModemStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetCurrentSlotAsync, uintptr(unsafe.Pointer(this)), uintptr(slotIndex), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandSlotManager) Add_SlotInfoChanged(handler TypedEventHandler[*IMobileBroadbandSlotManager, *IMobileBroadbandSlotInfoChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SlotInfoChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandSlotManager) Remove_SlotInfoChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SlotInfoChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMobileBroadbandSlotManager) Add_CurrentSlotIndexChanged(handler TypedEventHandler[*IMobileBroadbandSlotManager, *IMobileBroadbandCurrentSlotIndexChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_CurrentSlotIndexChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandSlotManager) Remove_CurrentSlotIndexChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_CurrentSlotIndexChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 612E3875-040A-4F99-A4F9-61D7C32DA129
var IID_IMobileBroadbandTransmissionStateChangedEventArgs = syscall.GUID{0x612E3875, 0x040A, 0x4F99,
	[8]byte{0xA4, 0xF9, 0x61, 0xD7, 0xC3, 0x2D, 0xA1, 0x29}}

type IMobileBroadbandTransmissionStateChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_IsTransmitting() bool
}

type IMobileBroadbandTransmissionStateChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_IsTransmitting uintptr
}

type IMobileBroadbandTransmissionStateChangedEventArgs struct {
	win32.IInspectable
}

func (this *IMobileBroadbandTransmissionStateChangedEventArgs) Vtbl() *IMobileBroadbandTransmissionStateChangedEventArgsVtbl {
	return (*IMobileBroadbandTransmissionStateChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandTransmissionStateChangedEventArgs) Get_IsTransmitting() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsTransmitting, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// E634F691-525A-4CE2-8FCE-AA4162579154
var IID_IMobileBroadbandUicc = syscall.GUID{0xE634F691, 0x525A, 0x4CE2,
	[8]byte{0x8F, 0xCE, 0xAA, 0x41, 0x62, 0x57, 0x91, 0x54}}

type IMobileBroadbandUiccInterface interface {
	win32.IInspectableInterface
	Get_SimIccId() string
	GetUiccAppsAsync() *IAsyncOperation[*IMobileBroadbandUiccAppsResult]
}

type IMobileBroadbandUiccVtbl struct {
	win32.IInspectableVtbl
	Get_SimIccId     uintptr
	GetUiccAppsAsync uintptr
}

type IMobileBroadbandUicc struct {
	win32.IInspectable
}

func (this *IMobileBroadbandUicc) Vtbl() *IMobileBroadbandUiccVtbl {
	return (*IMobileBroadbandUiccVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandUicc) Get_SimIccId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SimIccId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMobileBroadbandUicc) GetUiccAppsAsync() *IAsyncOperation[*IMobileBroadbandUiccAppsResult] {
	var _result *IAsyncOperation[*IMobileBroadbandUiccAppsResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetUiccAppsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4D170556-98A1-43DD-B2EC-50C90CF248DF
var IID_IMobileBroadbandUiccApp = syscall.GUID{0x4D170556, 0x98A1, 0x43DD,
	[8]byte{0xB2, 0xEC, 0x50, 0xC9, 0x0C, 0xF2, 0x48, 0xDF}}

type IMobileBroadbandUiccAppInterface interface {
	win32.IInspectableInterface
	Get_Id() *IBuffer
	Get_Kind() UiccAppKind
	GetRecordDetailsAsync(uiccFilePath *IIterable[uint32]) *IAsyncOperation[*IMobileBroadbandUiccAppRecordDetailsResult]
	ReadRecordAsync(uiccFilePath *IIterable[uint32], recordIndex int32) *IAsyncOperation[*IMobileBroadbandUiccAppReadRecordResult]
}

type IMobileBroadbandUiccAppVtbl struct {
	win32.IInspectableVtbl
	Get_Id                uintptr
	Get_Kind              uintptr
	GetRecordDetailsAsync uintptr
	ReadRecordAsync       uintptr
}

type IMobileBroadbandUiccApp struct {
	win32.IInspectable
}

func (this *IMobileBroadbandUiccApp) Vtbl() *IMobileBroadbandUiccAppVtbl {
	return (*IMobileBroadbandUiccAppVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandUiccApp) Get_Id() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandUiccApp) Get_Kind() UiccAppKind {
	var _result UiccAppKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Kind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandUiccApp) GetRecordDetailsAsync(uiccFilePath *IIterable[uint32]) *IAsyncOperation[*IMobileBroadbandUiccAppRecordDetailsResult] {
	var _result *IAsyncOperation[*IMobileBroadbandUiccAppRecordDetailsResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetRecordDetailsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uiccFilePath)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMobileBroadbandUiccApp) ReadRecordAsync(uiccFilePath *IIterable[uint32], recordIndex int32) *IAsyncOperation[*IMobileBroadbandUiccAppReadRecordResult] {
	var _result *IAsyncOperation[*IMobileBroadbandUiccAppReadRecordResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadRecordAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uiccFilePath)), uintptr(recordIndex), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 64C95285-358E-47C5-8249-695F383B2BDB
var IID_IMobileBroadbandUiccAppReadRecordResult = syscall.GUID{0x64C95285, 0x358E, 0x47C5,
	[8]byte{0x82, 0x49, 0x69, 0x5F, 0x38, 0x3B, 0x2B, 0xDB}}

type IMobileBroadbandUiccAppReadRecordResultInterface interface {
	win32.IInspectableInterface
	Get_Status() MobileBroadbandUiccAppOperationStatus
	Get_Data() *IBuffer
}

type IMobileBroadbandUiccAppReadRecordResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
	Get_Data   uintptr
}

type IMobileBroadbandUiccAppReadRecordResult struct {
	win32.IInspectable
}

func (this *IMobileBroadbandUiccAppReadRecordResult) Vtbl() *IMobileBroadbandUiccAppReadRecordResultVtbl {
	return (*IMobileBroadbandUiccAppReadRecordResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandUiccAppReadRecordResult) Get_Status() MobileBroadbandUiccAppOperationStatus {
	var _result MobileBroadbandUiccAppOperationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandUiccAppReadRecordResult) Get_Data() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D919682F-BE14-4934-981D-2F57B9ED83E6
var IID_IMobileBroadbandUiccAppRecordDetailsResult = syscall.GUID{0xD919682F, 0xBE14, 0x4934,
	[8]byte{0x98, 0x1D, 0x2F, 0x57, 0xB9, 0xED, 0x83, 0xE6}}

type IMobileBroadbandUiccAppRecordDetailsResultInterface interface {
	win32.IInspectableInterface
	Get_Status() MobileBroadbandUiccAppOperationStatus
	Get_Kind() UiccAppRecordKind
	Get_RecordCount() int32
	Get_RecordSize() int32
	Get_ReadAccessCondition() UiccAccessCondition
	Get_WriteAccessCondition() UiccAccessCondition
}

type IMobileBroadbandUiccAppRecordDetailsResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status               uintptr
	Get_Kind                 uintptr
	Get_RecordCount          uintptr
	Get_RecordSize           uintptr
	Get_ReadAccessCondition  uintptr
	Get_WriteAccessCondition uintptr
}

type IMobileBroadbandUiccAppRecordDetailsResult struct {
	win32.IInspectable
}

func (this *IMobileBroadbandUiccAppRecordDetailsResult) Vtbl() *IMobileBroadbandUiccAppRecordDetailsResultVtbl {
	return (*IMobileBroadbandUiccAppRecordDetailsResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandUiccAppRecordDetailsResult) Get_Status() MobileBroadbandUiccAppOperationStatus {
	var _result MobileBroadbandUiccAppOperationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandUiccAppRecordDetailsResult) Get_Kind() UiccAppRecordKind {
	var _result UiccAppRecordKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Kind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandUiccAppRecordDetailsResult) Get_RecordCount() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RecordCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandUiccAppRecordDetailsResult) Get_RecordSize() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RecordSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandUiccAppRecordDetailsResult) Get_ReadAccessCondition() UiccAccessCondition {
	var _result UiccAccessCondition
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReadAccessCondition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandUiccAppRecordDetailsResult) Get_WriteAccessCondition() UiccAccessCondition {
	var _result UiccAccessCondition
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WriteAccessCondition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 744930EB-8157-4A41-8494-6BF54C9B1D2B
var IID_IMobileBroadbandUiccAppsResult = syscall.GUID{0x744930EB, 0x8157, 0x4A41,
	[8]byte{0x84, 0x94, 0x6B, 0xF5, 0x4C, 0x9B, 0x1D, 0x2B}}

type IMobileBroadbandUiccAppsResultInterface interface {
	win32.IInspectableInterface
	Get_Status() MobileBroadbandUiccAppOperationStatus
	Get_UiccApps() *IVectorView[*IMobileBroadbandUiccApp]
}

type IMobileBroadbandUiccAppsResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status   uintptr
	Get_UiccApps uintptr
}

type IMobileBroadbandUiccAppsResult struct {
	win32.IInspectable
}

func (this *IMobileBroadbandUiccAppsResult) Vtbl() *IMobileBroadbandUiccAppsResultVtbl {
	return (*IMobileBroadbandUiccAppsResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMobileBroadbandUiccAppsResult) Get_Status() MobileBroadbandUiccAppOperationStatus {
	var _result MobileBroadbandUiccAppOperationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMobileBroadbandUiccAppsResult) Get_UiccApps() *IVectorView[*IMobileBroadbandUiccApp] {
	var _result *IVectorView[*IMobileBroadbandUiccApp]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UiccApps, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 50E3126D-A465-4EEB-9317-28A167630CEA
var IID_INetworkOperatorDataUsageTriggerDetails = syscall.GUID{0x50E3126D, 0xA465, 0x4EEB,
	[8]byte{0x93, 0x17, 0x28, 0xA1, 0x67, 0x63, 0x0C, 0xEA}}

type INetworkOperatorDataUsageTriggerDetailsInterface interface {
	win32.IInspectableInterface
	Get_NotificationKind() NetworkOperatorDataUsageNotificationKind
}

type INetworkOperatorDataUsageTriggerDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_NotificationKind uintptr
}

type INetworkOperatorDataUsageTriggerDetails struct {
	win32.IInspectable
}

func (this *INetworkOperatorDataUsageTriggerDetails) Vtbl() *INetworkOperatorDataUsageTriggerDetailsVtbl {
	return (*INetworkOperatorDataUsageTriggerDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INetworkOperatorDataUsageTriggerDetails) Get_NotificationKind() NetworkOperatorDataUsageNotificationKind {
	var _result NetworkOperatorDataUsageNotificationKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NotificationKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// BC68A9D1-82E1-4488-9F2C-1276C2468FAC
var IID_INetworkOperatorNotificationEventDetails = syscall.GUID{0xBC68A9D1, 0x82E1, 0x4488,
	[8]byte{0x9F, 0x2C, 0x12, 0x76, 0xC2, 0x46, 0x8F, 0xAC}}

type INetworkOperatorNotificationEventDetailsInterface interface {
	win32.IInspectableInterface
	Get_NotificationType() NetworkOperatorEventMessageType
	Get_NetworkAccountId() string
	Get_EncodingType() byte
	Get_Message() string
	Get_RuleId() string
	Get_SmsMessage() *ISmsMessage
}

type INetworkOperatorNotificationEventDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_NotificationType uintptr
	Get_NetworkAccountId uintptr
	Get_EncodingType     uintptr
	Get_Message          uintptr
	Get_RuleId           uintptr
	Get_SmsMessage       uintptr
}

type INetworkOperatorNotificationEventDetails struct {
	win32.IInspectable
}

func (this *INetworkOperatorNotificationEventDetails) Vtbl() *INetworkOperatorNotificationEventDetailsVtbl {
	return (*INetworkOperatorNotificationEventDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INetworkOperatorNotificationEventDetails) Get_NotificationType() NetworkOperatorEventMessageType {
	var _result NetworkOperatorEventMessageType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NotificationType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INetworkOperatorNotificationEventDetails) Get_NetworkAccountId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkAccountId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INetworkOperatorNotificationEventDetails) Get_EncodingType() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EncodingType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INetworkOperatorNotificationEventDetails) Get_Message() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Message, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INetworkOperatorNotificationEventDetails) Get_RuleId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RuleId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INetworkOperatorNotificationEventDetails) Get_SmsMessage() *ISmsMessage {
	var _result *ISmsMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SmsMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0BCC0284-412E-403D-ACC6-B757E34774A4
var IID_INetworkOperatorTetheringAccessPointConfiguration = syscall.GUID{0x0BCC0284, 0x412E, 0x403D,
	[8]byte{0xAC, 0xC6, 0xB7, 0x57, 0xE3, 0x47, 0x74, 0xA4}}

type INetworkOperatorTetheringAccessPointConfigurationInterface interface {
	win32.IInspectableInterface
	Get_Ssid() string
	Put_Ssid(value string)
	Get_Passphrase() string
	Put_Passphrase(value string)
}

type INetworkOperatorTetheringAccessPointConfigurationVtbl struct {
	win32.IInspectableVtbl
	Get_Ssid       uintptr
	Put_Ssid       uintptr
	Get_Passphrase uintptr
	Put_Passphrase uintptr
}

type INetworkOperatorTetheringAccessPointConfiguration struct {
	win32.IInspectable
}

func (this *INetworkOperatorTetheringAccessPointConfiguration) Vtbl() *INetworkOperatorTetheringAccessPointConfigurationVtbl {
	return (*INetworkOperatorTetheringAccessPointConfigurationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INetworkOperatorTetheringAccessPointConfiguration) Get_Ssid() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ssid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INetworkOperatorTetheringAccessPointConfiguration) Put_Ssid(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Ssid, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *INetworkOperatorTetheringAccessPointConfiguration) Get_Passphrase() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Passphrase, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INetworkOperatorTetheringAccessPointConfiguration) Put_Passphrase(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Passphrase, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// B1809142-7238-59A0-928B-74AB46FD64B6
var IID_INetworkOperatorTetheringAccessPointConfiguration2 = syscall.GUID{0xB1809142, 0x7238, 0x59A0,
	[8]byte{0x92, 0x8B, 0x74, 0xAB, 0x46, 0xFD, 0x64, 0xB6}}

type INetworkOperatorTetheringAccessPointConfiguration2Interface interface {
	win32.IInspectableInterface
	IsBandSupported(band TetheringWiFiBand) bool
	IsBandSupportedAsync(band TetheringWiFiBand) *IAsyncOperation[bool]
	Get_Band() TetheringWiFiBand
	Put_Band(value TetheringWiFiBand)
}

type INetworkOperatorTetheringAccessPointConfiguration2Vtbl struct {
	win32.IInspectableVtbl
	IsBandSupported      uintptr
	IsBandSupportedAsync uintptr
	Get_Band             uintptr
	Put_Band             uintptr
}

type INetworkOperatorTetheringAccessPointConfiguration2 struct {
	win32.IInspectable
}

func (this *INetworkOperatorTetheringAccessPointConfiguration2) Vtbl() *INetworkOperatorTetheringAccessPointConfiguration2Vtbl {
	return (*INetworkOperatorTetheringAccessPointConfiguration2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INetworkOperatorTetheringAccessPointConfiguration2) IsBandSupported(band TetheringWiFiBand) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsBandSupported, uintptr(unsafe.Pointer(this)), uintptr(band), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INetworkOperatorTetheringAccessPointConfiguration2) IsBandSupportedAsync(band TetheringWiFiBand) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsBandSupportedAsync, uintptr(unsafe.Pointer(this)), uintptr(band), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *INetworkOperatorTetheringAccessPointConfiguration2) Get_Band() TetheringWiFiBand {
	var _result TetheringWiFiBand
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Band, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INetworkOperatorTetheringAccessPointConfiguration2) Put_Band(value TetheringWiFiBand) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Band, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 709D254C-595F-4847-BB30-646935542918
var IID_INetworkOperatorTetheringClient = syscall.GUID{0x709D254C, 0x595F, 0x4847,
	[8]byte{0xBB, 0x30, 0x64, 0x69, 0x35, 0x54, 0x29, 0x18}}

type INetworkOperatorTetheringClientInterface interface {
	win32.IInspectableInterface
	Get_MacAddress() string
	Get_HostNames() *IVectorView[*IHostName]
}

type INetworkOperatorTetheringClientVtbl struct {
	win32.IInspectableVtbl
	Get_MacAddress uintptr
	Get_HostNames  uintptr
}

type INetworkOperatorTetheringClient struct {
	win32.IInspectable
}

func (this *INetworkOperatorTetheringClient) Vtbl() *INetworkOperatorTetheringClientVtbl {
	return (*INetworkOperatorTetheringClientVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INetworkOperatorTetheringClient) Get_MacAddress() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MacAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *INetworkOperatorTetheringClient) Get_HostNames() *IVectorView[*IHostName] {
	var _result *IVectorView[*IHostName]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HostNames, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 91B14016-8DCA-4225-BBED-EEF8B8D718D7
var IID_INetworkOperatorTetheringClientManager = syscall.GUID{0x91B14016, 0x8DCA, 0x4225,
	[8]byte{0xBB, 0xED, 0xEE, 0xF8, 0xB8, 0xD7, 0x18, 0xD7}}

type INetworkOperatorTetheringClientManagerInterface interface {
	win32.IInspectableInterface
	GetTetheringClients() *IVectorView[*INetworkOperatorTetheringClient]
}

type INetworkOperatorTetheringClientManagerVtbl struct {
	win32.IInspectableVtbl
	GetTetheringClients uintptr
}

type INetworkOperatorTetheringClientManager struct {
	win32.IInspectable
}

func (this *INetworkOperatorTetheringClientManager) Vtbl() *INetworkOperatorTetheringClientManagerVtbl {
	return (*INetworkOperatorTetheringClientManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INetworkOperatorTetheringClientManager) GetTetheringClients() *IVectorView[*INetworkOperatorTetheringClient] {
	var _result *IVectorView[*INetworkOperatorTetheringClient]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetTetheringClients, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0108916D-9E9A-4AF6-8DA3-60493B19C204
var IID_INetworkOperatorTetheringEntitlementCheck = syscall.GUID{0x0108916D, 0x9E9A, 0x4AF6,
	[8]byte{0x8D, 0xA3, 0x60, 0x49, 0x3B, 0x19, 0xC2, 0x04}}

type INetworkOperatorTetheringEntitlementCheckInterface interface {
	win32.IInspectableInterface
	AuthorizeTethering(allow bool, entitlementFailureReason string)
}

type INetworkOperatorTetheringEntitlementCheckVtbl struct {
	win32.IInspectableVtbl
	AuthorizeTethering uintptr
}

type INetworkOperatorTetheringEntitlementCheck struct {
	win32.IInspectable
}

func (this *INetworkOperatorTetheringEntitlementCheck) Vtbl() *INetworkOperatorTetheringEntitlementCheckVtbl {
	return (*INetworkOperatorTetheringEntitlementCheckVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INetworkOperatorTetheringEntitlementCheck) AuthorizeTethering(allow bool, entitlementFailureReason string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AuthorizeTethering, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&allow))), NewHStr(entitlementFailureReason).Ptr)
	_ = _hr
}

// D45A8DA0-0E86-4D98-8BA4-DD70D4B764D3
var IID_INetworkOperatorTetheringManager = syscall.GUID{0xD45A8DA0, 0x0E86, 0x4D98,
	[8]byte{0x8B, 0xA4, 0xDD, 0x70, 0xD4, 0xB7, 0x64, 0xD3}}

type INetworkOperatorTetheringManagerInterface interface {
	win32.IInspectableInterface
	Get_MaxClientCount() uint32
	Get_ClientCount() uint32
	Get_TetheringOperationalState() TetheringOperationalState
	GetCurrentAccessPointConfiguration() *INetworkOperatorTetheringAccessPointConfiguration
	ConfigureAccessPointAsync(configuration *INetworkOperatorTetheringAccessPointConfiguration) *IAsyncAction
	StartTetheringAsync() *IAsyncOperation[*INetworkOperatorTetheringOperationResult]
	StopTetheringAsync() *IAsyncOperation[*INetworkOperatorTetheringOperationResult]
}

type INetworkOperatorTetheringManagerVtbl struct {
	win32.IInspectableVtbl
	Get_MaxClientCount                 uintptr
	Get_ClientCount                    uintptr
	Get_TetheringOperationalState      uintptr
	GetCurrentAccessPointConfiguration uintptr
	ConfigureAccessPointAsync          uintptr
	StartTetheringAsync                uintptr
	StopTetheringAsync                 uintptr
}

type INetworkOperatorTetheringManager struct {
	win32.IInspectable
}

func (this *INetworkOperatorTetheringManager) Vtbl() *INetworkOperatorTetheringManagerVtbl {
	return (*INetworkOperatorTetheringManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INetworkOperatorTetheringManager) Get_MaxClientCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxClientCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INetworkOperatorTetheringManager) Get_ClientCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ClientCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INetworkOperatorTetheringManager) Get_TetheringOperationalState() TetheringOperationalState {
	var _result TetheringOperationalState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TetheringOperationalState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INetworkOperatorTetheringManager) GetCurrentAccessPointConfiguration() *INetworkOperatorTetheringAccessPointConfiguration {
	var _result *INetworkOperatorTetheringAccessPointConfiguration
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentAccessPointConfiguration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *INetworkOperatorTetheringManager) ConfigureAccessPointAsync(configuration *INetworkOperatorTetheringAccessPointConfiguration) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConfigureAccessPointAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(configuration)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *INetworkOperatorTetheringManager) StartTetheringAsync() *IAsyncOperation[*INetworkOperatorTetheringOperationResult] {
	var _result *IAsyncOperation[*INetworkOperatorTetheringOperationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartTetheringAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *INetworkOperatorTetheringManager) StopTetheringAsync() *IAsyncOperation[*INetworkOperatorTetheringOperationResult] {
	var _result *IAsyncOperation[*INetworkOperatorTetheringOperationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopTetheringAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3EBCBACC-F8C3-405C-9964-70A1EEABE194
var IID_INetworkOperatorTetheringManagerStatics = syscall.GUID{0x3EBCBACC, 0xF8C3, 0x405C,
	[8]byte{0x99, 0x64, 0x70, 0xA1, 0xEE, 0xAB, 0xE1, 0x94}}

type INetworkOperatorTetheringManagerStaticsInterface interface {
	win32.IInspectableInterface
	GetTetheringCapability(networkAccountId string) TetheringCapability
	CreateFromNetworkAccountId(networkAccountId string) *INetworkOperatorTetheringManager
}

type INetworkOperatorTetheringManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetTetheringCapability     uintptr
	CreateFromNetworkAccountId uintptr
}

type INetworkOperatorTetheringManagerStatics struct {
	win32.IInspectable
}

func (this *INetworkOperatorTetheringManagerStatics) Vtbl() *INetworkOperatorTetheringManagerStaticsVtbl {
	return (*INetworkOperatorTetheringManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INetworkOperatorTetheringManagerStatics) GetTetheringCapability(networkAccountId string) TetheringCapability {
	var _result TetheringCapability
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetTetheringCapability, uintptr(unsafe.Pointer(this)), NewHStr(networkAccountId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INetworkOperatorTetheringManagerStatics) CreateFromNetworkAccountId(networkAccountId string) *INetworkOperatorTetheringManager {
	var _result *INetworkOperatorTetheringManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromNetworkAccountId, uintptr(unsafe.Pointer(this)), NewHStr(networkAccountId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5B235412-35F0-49E7-9B08-16D278FBAA42
var IID_INetworkOperatorTetheringManagerStatics2 = syscall.GUID{0x5B235412, 0x35F0, 0x49E7,
	[8]byte{0x9B, 0x08, 0x16, 0xD2, 0x78, 0xFB, 0xAA, 0x42}}

type INetworkOperatorTetheringManagerStatics2Interface interface {
	win32.IInspectableInterface
	GetTetheringCapabilityFromConnectionProfile(profile *IConnectionProfile) TetheringCapability
	CreateFromConnectionProfile(profile *IConnectionProfile) *INetworkOperatorTetheringManager
}

type INetworkOperatorTetheringManagerStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetTetheringCapabilityFromConnectionProfile uintptr
	CreateFromConnectionProfile                 uintptr
}

type INetworkOperatorTetheringManagerStatics2 struct {
	win32.IInspectable
}

func (this *INetworkOperatorTetheringManagerStatics2) Vtbl() *INetworkOperatorTetheringManagerStatics2Vtbl {
	return (*INetworkOperatorTetheringManagerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INetworkOperatorTetheringManagerStatics2) GetTetheringCapabilityFromConnectionProfile(profile *IConnectionProfile) TetheringCapability {
	var _result TetheringCapability
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetTetheringCapabilityFromConnectionProfile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(profile)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INetworkOperatorTetheringManagerStatics2) CreateFromConnectionProfile(profile *IConnectionProfile) *INetworkOperatorTetheringManager {
	var _result *INetworkOperatorTetheringManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromConnectionProfile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(profile)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8FDAADB6-4AF9-4F21-9B58-D53E9F24231E
var IID_INetworkOperatorTetheringManagerStatics3 = syscall.GUID{0x8FDAADB6, 0x4AF9, 0x4F21,
	[8]byte{0x9B, 0x58, 0xD5, 0x3E, 0x9F, 0x24, 0x23, 0x1E}}

type INetworkOperatorTetheringManagerStatics3Interface interface {
	win32.IInspectableInterface
	CreateFromConnectionProfileWithTargetAdapter(profile *IConnectionProfile, adapter *INetworkAdapter) *INetworkOperatorTetheringManager
}

type INetworkOperatorTetheringManagerStatics3Vtbl struct {
	win32.IInspectableVtbl
	CreateFromConnectionProfileWithTargetAdapter uintptr
}

type INetworkOperatorTetheringManagerStatics3 struct {
	win32.IInspectable
}

func (this *INetworkOperatorTetheringManagerStatics3) Vtbl() *INetworkOperatorTetheringManagerStatics3Vtbl {
	return (*INetworkOperatorTetheringManagerStatics3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INetworkOperatorTetheringManagerStatics3) CreateFromConnectionProfileWithTargetAdapter(profile *IConnectionProfile, adapter *INetworkAdapter) *INetworkOperatorTetheringManager {
	var _result *INetworkOperatorTetheringManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromConnectionProfileWithTargetAdapter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(profile)), uintptr(unsafe.Pointer(adapter)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B3B9F9D0-EBFF-46A4-A847-D663D8B0977E
var IID_INetworkOperatorTetheringManagerStatics4 = syscall.GUID{0xB3B9F9D0, 0xEBFF, 0x46A4,
	[8]byte{0xA8, 0x47, 0xD6, 0x63, 0xD8, 0xB0, 0x97, 0x7E}}

type INetworkOperatorTetheringManagerStatics4Interface interface {
	win32.IInspectableInterface
	IsNoConnectionsTimeoutEnabled() bool
	EnableNoConnectionsTimeout()
	EnableNoConnectionsTimeoutAsync() *IAsyncAction
	DisableNoConnectionsTimeout()
	DisableNoConnectionsTimeoutAsync() *IAsyncAction
}

type INetworkOperatorTetheringManagerStatics4Vtbl struct {
	win32.IInspectableVtbl
	IsNoConnectionsTimeoutEnabled    uintptr
	EnableNoConnectionsTimeout       uintptr
	EnableNoConnectionsTimeoutAsync  uintptr
	DisableNoConnectionsTimeout      uintptr
	DisableNoConnectionsTimeoutAsync uintptr
}

type INetworkOperatorTetheringManagerStatics4 struct {
	win32.IInspectable
}

func (this *INetworkOperatorTetheringManagerStatics4) Vtbl() *INetworkOperatorTetheringManagerStatics4Vtbl {
	return (*INetworkOperatorTetheringManagerStatics4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INetworkOperatorTetheringManagerStatics4) IsNoConnectionsTimeoutEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsNoConnectionsTimeoutEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INetworkOperatorTetheringManagerStatics4) EnableNoConnectionsTimeout() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnableNoConnectionsTimeout, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *INetworkOperatorTetheringManagerStatics4) EnableNoConnectionsTimeoutAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnableNoConnectionsTimeoutAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *INetworkOperatorTetheringManagerStatics4) DisableNoConnectionsTimeout() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DisableNoConnectionsTimeout, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *INetworkOperatorTetheringManagerStatics4) DisableNoConnectionsTimeoutAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DisableNoConnectionsTimeoutAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EBD203A1-01BA-476D-B4B3-BF3D12C8F80C
var IID_INetworkOperatorTetheringOperationResult = syscall.GUID{0xEBD203A1, 0x01BA, 0x476D,
	[8]byte{0xB4, 0xB3, 0xBF, 0x3D, 0x12, 0xC8, 0xF8, 0x0C}}

type INetworkOperatorTetheringOperationResultInterface interface {
	win32.IInspectableInterface
	Get_Status() TetheringOperationStatus
	Get_AdditionalErrorMessage() string
}

type INetworkOperatorTetheringOperationResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status                 uintptr
	Get_AdditionalErrorMessage uintptr
}

type INetworkOperatorTetheringOperationResult struct {
	win32.IInspectable
}

func (this *INetworkOperatorTetheringOperationResult) Vtbl() *INetworkOperatorTetheringOperationResultVtbl {
	return (*INetworkOperatorTetheringOperationResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INetworkOperatorTetheringOperationResult) Get_Status() TetheringOperationStatus {
	var _result TetheringOperationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INetworkOperatorTetheringOperationResult) Get_AdditionalErrorMessage() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AdditionalErrorMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 217700E0-8203-11DF-ADB9-F4CE462D9137
var IID_IProvisionFromXmlDocumentResults = syscall.GUID{0x217700E0, 0x8203, 0x11DF,
	[8]byte{0xAD, 0xB9, 0xF4, 0xCE, 0x46, 0x2D, 0x91, 0x37}}

type IProvisionFromXmlDocumentResultsInterface interface {
	win32.IInspectableInterface
	Get_AllElementsProvisioned() bool
	Get_ProvisionResultsXml() string
}

type IProvisionFromXmlDocumentResultsVtbl struct {
	win32.IInspectableVtbl
	Get_AllElementsProvisioned uintptr
	Get_ProvisionResultsXml    uintptr
}

type IProvisionFromXmlDocumentResults struct {
	win32.IInspectable
}

func (this *IProvisionFromXmlDocumentResults) Vtbl() *IProvisionFromXmlDocumentResultsVtbl {
	return (*IProvisionFromXmlDocumentResultsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProvisionFromXmlDocumentResults) Get_AllElementsProvisioned() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllElementsProvisioned, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProvisionFromXmlDocumentResults) Get_ProvisionResultsXml() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProvisionResultsXml, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 217700E0-8202-11DF-ADB9-F4CE462D9137
var IID_IProvisionedProfile = syscall.GUID{0x217700E0, 0x8202, 0x11DF,
	[8]byte{0xAD, 0xB9, 0xF4, 0xCE, 0x46, 0x2D, 0x91, 0x37}}

type IProvisionedProfileInterface interface {
	win32.IInspectableInterface
	UpdateCost(value NetworkCostType)
	UpdateUsage(value ProfileUsage)
}

type IProvisionedProfileVtbl struct {
	win32.IInspectableVtbl
	UpdateCost  uintptr
	UpdateUsage uintptr
}

type IProvisionedProfile struct {
	win32.IInspectable
}

func (this *IProvisionedProfile) Vtbl() *IProvisionedProfileVtbl {
	return (*IProvisionedProfileVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProvisionedProfile) UpdateCost(value NetworkCostType) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UpdateCost, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IProvisionedProfile) UpdateUsage(value ProfileUsage) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UpdateUsage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

// 217700E0-8201-11DF-ADB9-F4CE462D9137
var IID_IProvisioningAgent = syscall.GUID{0x217700E0, 0x8201, 0x11DF,
	[8]byte{0xAD, 0xB9, 0xF4, 0xCE, 0x46, 0x2D, 0x91, 0x37}}

type IProvisioningAgentInterface interface {
	win32.IInspectableInterface
	ProvisionFromXmlDocumentAsync(provisioningXmlDocument string) *IAsyncOperation[*IProvisionFromXmlDocumentResults]
	GetProvisionedProfile(mediaType ProfileMediaType, profileName string) *IProvisionedProfile
}

type IProvisioningAgentVtbl struct {
	win32.IInspectableVtbl
	ProvisionFromXmlDocumentAsync uintptr
	GetProvisionedProfile         uintptr
}

type IProvisioningAgent struct {
	win32.IInspectable
}

func (this *IProvisioningAgent) Vtbl() *IProvisioningAgentVtbl {
	return (*IProvisioningAgentVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProvisioningAgent) ProvisionFromXmlDocumentAsync(provisioningXmlDocument string) *IAsyncOperation[*IProvisionFromXmlDocumentResults] {
	var _result *IAsyncOperation[*IProvisionFromXmlDocumentResults]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ProvisionFromXmlDocumentAsync, uintptr(unsafe.Pointer(this)), NewHStr(provisioningXmlDocument).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProvisioningAgent) GetProvisionedProfile(mediaType ProfileMediaType, profileName string) *IProvisionedProfile {
	var _result *IProvisionedProfile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetProvisionedProfile, uintptr(unsafe.Pointer(this)), uintptr(mediaType), NewHStr(profileName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 217700E0-8101-11DF-ADB9-F4CE462D9137
var IID_IProvisioningAgentStaticMethods = syscall.GUID{0x217700E0, 0x8101, 0x11DF,
	[8]byte{0xAD, 0xB9, 0xF4, 0xCE, 0x46, 0x2D, 0x91, 0x37}}

type IProvisioningAgentStaticMethodsInterface interface {
	win32.IInspectableInterface
	CreateFromNetworkAccountId(networkAccountId string) *IProvisioningAgent
}

type IProvisioningAgentStaticMethodsVtbl struct {
	win32.IInspectableVtbl
	CreateFromNetworkAccountId uintptr
}

type IProvisioningAgentStaticMethods struct {
	win32.IInspectable
}

func (this *IProvisioningAgentStaticMethods) Vtbl() *IProvisioningAgentStaticMethodsVtbl {
	return (*IProvisioningAgentStaticMethodsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProvisioningAgentStaticMethods) CreateFromNetworkAccountId(networkAccountId string) *IProvisioningAgent {
	var _result *IProvisioningAgent
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromNetworkAccountId, uintptr(unsafe.Pointer(this)), NewHStr(networkAccountId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 03C65E9D-5926-41F3-A94E-B50926FC421B
var IID_ITetheringEntitlementCheckTriggerDetails = syscall.GUID{0x03C65E9D, 0x5926, 0x41F3,
	[8]byte{0xA9, 0x4E, 0xB5, 0x09, 0x26, 0xFC, 0x42, 0x1B}}

type ITetheringEntitlementCheckTriggerDetailsInterface interface {
	win32.IInspectableInterface
	Get_NetworkAccountId() string
	AllowTethering()
	DenyTethering(entitlementFailureReason string)
}

type ITetheringEntitlementCheckTriggerDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_NetworkAccountId uintptr
	AllowTethering       uintptr
	DenyTethering        uintptr
}

type ITetheringEntitlementCheckTriggerDetails struct {
	win32.IInspectable
}

func (this *ITetheringEntitlementCheckTriggerDetails) Vtbl() *ITetheringEntitlementCheckTriggerDetailsVtbl {
	return (*ITetheringEntitlementCheckTriggerDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITetheringEntitlementCheckTriggerDetails) Get_NetworkAccountId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkAccountId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ITetheringEntitlementCheckTriggerDetails) AllowTethering() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AllowTethering, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *ITetheringEntitlementCheckTriggerDetails) DenyTethering(entitlementFailureReason string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DenyTethering, uintptr(unsafe.Pointer(this)), NewHStr(entitlementFailureReason).Ptr)
	_ = _hr
}

// 2F9ACF82-2004-4D5D-BF81-2ABA1B4BE4A8
var IID_IUssdMessage = syscall.GUID{0x2F9ACF82, 0x2004, 0x4D5D,
	[8]byte{0xBF, 0x81, 0x2A, 0xBA, 0x1B, 0x4B, 0xE4, 0xA8}}

type IUssdMessageInterface interface {
	win32.IInspectableInterface
	Get_DataCodingScheme() byte
	Put_DataCodingScheme(value byte)
	GetPayload() []byte
	SetPayload(valueLength uint32, value *byte)
	Get_PayloadAsText() string
	Put_PayloadAsText(value string)
}

type IUssdMessageVtbl struct {
	win32.IInspectableVtbl
	Get_DataCodingScheme uintptr
	Put_DataCodingScheme uintptr
	GetPayload           uintptr
	SetPayload           uintptr
	Get_PayloadAsText    uintptr
	Put_PayloadAsText    uintptr
}

type IUssdMessage struct {
	win32.IInspectable
}

func (this *IUssdMessage) Vtbl() *IUssdMessageVtbl {
	return (*IUssdMessageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUssdMessage) Get_DataCodingScheme() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DataCodingScheme, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUssdMessage) Put_DataCodingScheme(value byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DataCodingScheme, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IUssdMessage) GetPayload() []byte {
	var _result []byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPayload, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUssdMessage) SetPayload(valueLength uint32, value *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetPayload, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IUssdMessage) Get_PayloadAsText() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PayloadAsText, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IUssdMessage) Put_PayloadAsText(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PayloadAsText, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// 2F9ACF82-1003-4D5D-BF81-2ABA1B4BE4A8
var IID_IUssdMessageFactory = syscall.GUID{0x2F9ACF82, 0x1003, 0x4D5D,
	[8]byte{0xBF, 0x81, 0x2A, 0xBA, 0x1B, 0x4B, 0xE4, 0xA8}}

type IUssdMessageFactoryInterface interface {
	win32.IInspectableInterface
	CreateMessage(messageText string) *IUssdMessage
}

type IUssdMessageFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateMessage uintptr
}

type IUssdMessageFactory struct {
	win32.IInspectable
}

func (this *IUssdMessageFactory) Vtbl() *IUssdMessageFactoryVtbl {
	return (*IUssdMessageFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUssdMessageFactory) CreateMessage(messageText string) *IUssdMessage {
	var _result *IUssdMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateMessage, uintptr(unsafe.Pointer(this)), NewHStr(messageText).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2F9ACF82-2005-4D5D-BF81-2ABA1B4BE4A8
var IID_IUssdReply = syscall.GUID{0x2F9ACF82, 0x2005, 0x4D5D,
	[8]byte{0xBF, 0x81, 0x2A, 0xBA, 0x1B, 0x4B, 0xE4, 0xA8}}

type IUssdReplyInterface interface {
	win32.IInspectableInterface
	Get_ResultCode() UssdResultCode
	Get_Message() *IUssdMessage
}

type IUssdReplyVtbl struct {
	win32.IInspectableVtbl
	Get_ResultCode uintptr
	Get_Message    uintptr
}

type IUssdReply struct {
	win32.IInspectable
}

func (this *IUssdReply) Vtbl() *IUssdReplyVtbl {
	return (*IUssdReplyVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUssdReply) Get_ResultCode() UssdResultCode {
	var _result UssdResultCode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResultCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUssdReply) Get_Message() *IUssdMessage {
	var _result *IUssdMessage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Message, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2F9ACF82-2002-4D5D-BF81-2ABA1B4BE4A8
var IID_IUssdSession = syscall.GUID{0x2F9ACF82, 0x2002, 0x4D5D,
	[8]byte{0xBF, 0x81, 0x2A, 0xBA, 0x1B, 0x4B, 0xE4, 0xA8}}

type IUssdSessionInterface interface {
	win32.IInspectableInterface
	SendMessageAndGetReplyAsync(message *IUssdMessage) *IAsyncOperation[*IUssdReply]
	Close()
}

type IUssdSessionVtbl struct {
	win32.IInspectableVtbl
	SendMessageAndGetReplyAsync uintptr
	Close                       uintptr
}

type IUssdSession struct {
	win32.IInspectable
}

func (this *IUssdSession) Vtbl() *IUssdSessionVtbl {
	return (*IUssdSessionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUssdSession) SendMessageAndGetReplyAsync(message *IUssdMessage) *IAsyncOperation[*IUssdReply] {
	var _result *IAsyncOperation[*IUssdReply]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendMessageAndGetReplyAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(message)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUssdSession) Close() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Close, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 2F9ACF82-1001-4D5D-BF81-2ABA1B4BE4A8
var IID_IUssdSessionStatics = syscall.GUID{0x2F9ACF82, 0x1001, 0x4D5D,
	[8]byte{0xBF, 0x81, 0x2A, 0xBA, 0x1B, 0x4B, 0xE4, 0xA8}}

type IUssdSessionStaticsInterface interface {
	win32.IInspectableInterface
	CreateFromNetworkAccountId(networkAccountId string) *IUssdSession
	CreateFromNetworkInterfaceId(networkInterfaceId string) *IUssdSession
}

type IUssdSessionStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateFromNetworkAccountId   uintptr
	CreateFromNetworkInterfaceId uintptr
}

type IUssdSessionStatics struct {
	win32.IInspectable
}

func (this *IUssdSessionStatics) Vtbl() *IUssdSessionStaticsVtbl {
	return (*IUssdSessionStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUssdSessionStatics) CreateFromNetworkAccountId(networkAccountId string) *IUssdSession {
	var _result *IUssdSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromNetworkAccountId, uintptr(unsafe.Pointer(this)), NewHStr(networkAccountId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUssdSessionStatics) CreateFromNetworkInterfaceId(networkInterfaceId string) *IUssdSession {
	var _result *IUssdSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromNetworkInterfaceId, uintptr(unsafe.Pointer(this)), NewHStr(networkInterfaceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type HotspotCredentialsAuthenticationResult struct {
	RtClass
	*IHotspotCredentialsAuthenticationResult
}

type KnownCSimFilePaths struct {
	RtClass
}

func NewIKnownCSimFilePathsStatics() *IKnownCSimFilePathsStatics {
	var p *IKnownCSimFilePathsStatics
	hs := NewHStr("Windows.Networking.NetworkOperators.KnownCSimFilePaths")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IKnownCSimFilePathsStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type KnownRuimFilePaths struct {
	RtClass
}

func NewIKnownRuimFilePathsStatics() *IKnownRuimFilePathsStatics {
	var p *IKnownRuimFilePathsStatics
	hs := NewHStr("Windows.Networking.NetworkOperators.KnownRuimFilePaths")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IKnownRuimFilePathsStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type KnownSimFilePaths struct {
	RtClass
}

func NewIKnownSimFilePathsStatics() *IKnownSimFilePathsStatics {
	var p *IKnownSimFilePathsStatics
	hs := NewHStr("Windows.Networking.NetworkOperators.KnownSimFilePaths")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IKnownSimFilePathsStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type KnownUSimFilePaths struct {
	RtClass
}

func NewIKnownUSimFilePathsStatics() *IKnownUSimFilePathsStatics {
	var p *IKnownUSimFilePathsStatics
	hs := NewHStr("Windows.Networking.NetworkOperators.KnownUSimFilePaths")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IKnownUSimFilePathsStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type MobileBroadbandAccount struct {
	RtClass
	*IMobileBroadbandAccount
}

func NewIMobileBroadbandAccountStatics() *IMobileBroadbandAccountStatics {
	var p *IMobileBroadbandAccountStatics
	hs := NewHStr("Windows.Networking.NetworkOperators.MobileBroadbandAccount")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMobileBroadbandAccountStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type MobileBroadbandAccountEventArgs struct {
	RtClass
	*IMobileBroadbandAccountEventArgs
}

type MobileBroadbandAccountUpdatedEventArgs struct {
	RtClass
	*IMobileBroadbandAccountUpdatedEventArgs
}

type MobileBroadbandAccountWatcher struct {
	RtClass
	*IMobileBroadbandAccountWatcher
}

func NewMobileBroadbandAccountWatcher() *MobileBroadbandAccountWatcher {
	hs := NewHStr("Windows.Networking.NetworkOperators.MobileBroadbandAccountWatcher")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &MobileBroadbandAccountWatcher{
		RtClass:                        RtClass{PInspect: p},
		IMobileBroadbandAccountWatcher: (*IMobileBroadbandAccountWatcher)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type MobileBroadbandDeviceInformation struct {
	RtClass
	*IMobileBroadbandDeviceInformation
}

type MobileBroadbandDeviceService struct {
	RtClass
	*IMobileBroadbandDeviceService
}

type MobileBroadbandDeviceServiceCommandResult struct {
	RtClass
	*IMobileBroadbandDeviceServiceCommandResult
}

type MobileBroadbandDeviceServiceCommandSession struct {
	RtClass
	*IMobileBroadbandDeviceServiceCommandSession
}

type MobileBroadbandDeviceServiceDataSession struct {
	RtClass
	*IMobileBroadbandDeviceServiceDataSession
}

type MobileBroadbandDeviceServiceInformation struct {
	RtClass
	*IMobileBroadbandDeviceServiceInformation
}

type MobileBroadbandDeviceServiceTriggerDetails struct {
	RtClass
	*IMobileBroadbandDeviceServiceTriggerDetails
}

type MobileBroadbandModem struct {
	RtClass
	*IMobileBroadbandModem
}

func NewIMobileBroadbandModemStatics() *IMobileBroadbandModemStatics {
	var p *IMobileBroadbandModemStatics
	hs := NewHStr("Windows.Networking.NetworkOperators.MobileBroadbandModem")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMobileBroadbandModemStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type MobileBroadbandModemConfiguration struct {
	RtClass
	*IMobileBroadbandModemConfiguration
}

type MobileBroadbandNetwork struct {
	RtClass
	*IMobileBroadbandNetwork
}

type MobileBroadbandNetworkRegistrationStateChange struct {
	RtClass
	*IMobileBroadbandNetworkRegistrationStateChange
}

type MobileBroadbandNetworkRegistrationStateChangeTriggerDetails struct {
	RtClass
	*IMobileBroadbandNetworkRegistrationStateChangeTriggerDetails
}

type MobileBroadbandPin struct {
	RtClass
	*IMobileBroadbandPin
}

type MobileBroadbandPinLockStateChange struct {
	RtClass
	*IMobileBroadbandPinLockStateChange
}

type MobileBroadbandPinLockStateChangeTriggerDetails struct {
	RtClass
	*IMobileBroadbandPinLockStateChangeTriggerDetails
}

type MobileBroadbandPinManager struct {
	RtClass
	*IMobileBroadbandPinManager
}

type MobileBroadbandPinOperationResult struct {
	RtClass
	*IMobileBroadbandPinOperationResult
}

type MobileBroadbandRadioStateChange struct {
	RtClass
	*IMobileBroadbandRadioStateChange
}

type MobileBroadbandRadioStateChangeTriggerDetails struct {
	RtClass
	*IMobileBroadbandRadioStateChangeTriggerDetails
}

type MobileBroadbandUicc struct {
	RtClass
	*IMobileBroadbandUicc
}

type MobileBroadbandUiccApp struct {
	RtClass
	*IMobileBroadbandUiccApp
}

type MobileBroadbandUiccAppReadRecordResult struct {
	RtClass
	*IMobileBroadbandUiccAppReadRecordResult
}

type MobileBroadbandUiccAppRecordDetailsResult struct {
	RtClass
	*IMobileBroadbandUiccAppRecordDetailsResult
}

type MobileBroadbandUiccAppsResult struct {
	RtClass
	*IMobileBroadbandUiccAppsResult
}

type NetworkOperatorNotificationEventDetails struct {
	RtClass
	*INetworkOperatorNotificationEventDetails
}

type NetworkOperatorTetheringAccessPointConfiguration struct {
	RtClass
	*INetworkOperatorTetheringAccessPointConfiguration
}

func NewNetworkOperatorTetheringAccessPointConfiguration() *NetworkOperatorTetheringAccessPointConfiguration {
	hs := NewHStr("Windows.Networking.NetworkOperators.NetworkOperatorTetheringAccessPointConfiguration")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &NetworkOperatorTetheringAccessPointConfiguration{
		RtClass: RtClass{PInspect: p},
		INetworkOperatorTetheringAccessPointConfiguration: (*INetworkOperatorTetheringAccessPointConfiguration)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type NetworkOperatorTetheringClient struct {
	RtClass
	*INetworkOperatorTetheringClient
}

type NetworkOperatorTetheringManager struct {
	RtClass
	*INetworkOperatorTetheringManager
}

func NewINetworkOperatorTetheringManagerStatics4() *INetworkOperatorTetheringManagerStatics4 {
	var p *INetworkOperatorTetheringManagerStatics4
	hs := NewHStr("Windows.Networking.NetworkOperators.NetworkOperatorTetheringManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_INetworkOperatorTetheringManagerStatics4, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewINetworkOperatorTetheringManagerStatics() *INetworkOperatorTetheringManagerStatics {
	var p *INetworkOperatorTetheringManagerStatics
	hs := NewHStr("Windows.Networking.NetworkOperators.NetworkOperatorTetheringManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_INetworkOperatorTetheringManagerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewINetworkOperatorTetheringManagerStatics3() *INetworkOperatorTetheringManagerStatics3 {
	var p *INetworkOperatorTetheringManagerStatics3
	hs := NewHStr("Windows.Networking.NetworkOperators.NetworkOperatorTetheringManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_INetworkOperatorTetheringManagerStatics3, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewINetworkOperatorTetheringManagerStatics2() *INetworkOperatorTetheringManagerStatics2 {
	var p *INetworkOperatorTetheringManagerStatics2
	hs := NewHStr("Windows.Networking.NetworkOperators.NetworkOperatorTetheringManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_INetworkOperatorTetheringManagerStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type NetworkOperatorTetheringOperationResult struct {
	RtClass
	*INetworkOperatorTetheringOperationResult
}

type ProvisionFromXmlDocumentResults struct {
	RtClass
	*IProvisionFromXmlDocumentResults
}

type ProvisionedProfile struct {
	RtClass
	*IProvisionedProfile
}

type ProvisioningAgent struct {
	RtClass
	*IProvisioningAgent
}

func NewProvisioningAgent() *ProvisioningAgent {
	hs := NewHStr("Windows.Networking.NetworkOperators.ProvisioningAgent")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &ProvisioningAgent{
		RtClass:            RtClass{PInspect: p},
		IProvisioningAgent: (*IProvisioningAgent)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewIProvisioningAgentStaticMethods() *IProvisioningAgentStaticMethods {
	var p *IProvisioningAgentStaticMethods
	hs := NewHStr("Windows.Networking.NetworkOperators.ProvisioningAgent")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IProvisioningAgentStaticMethods, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type UssdMessage struct {
	RtClass
	*IUssdMessage
}

func NewUssdMessage_CreateMessage(messageText string) *UssdMessage {
	hs := NewHStr("Windows.Networking.NetworkOperators.UssdMessage")
	var pFac *IUssdMessageFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IUssdMessageFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IUssdMessage
	p = pFac.CreateMessage(messageText)
	result := &UssdMessage{
		RtClass:      RtClass{PInspect: &p.IInspectable},
		IUssdMessage: p,
	}
	com.AddToScope(result)
	return result
}

type UssdReply struct {
	RtClass
	*IUssdReply
}

type UssdSession struct {
	RtClass
	*IUssdSession
}

func NewIUssdSessionStatics() *IUssdSessionStatics {
	var p *IUssdSessionStatics
	hs := NewHStr("Windows.Networking.NetworkOperators.UssdSession")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IUssdSessionStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
