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
type CellularApnAuthenticationType int32

const (
	CellularApnAuthenticationType_None     CellularApnAuthenticationType = 0
	CellularApnAuthenticationType_Pap      CellularApnAuthenticationType = 1
	CellularApnAuthenticationType_Chap     CellularApnAuthenticationType = 2
	CellularApnAuthenticationType_Mschapv2 CellularApnAuthenticationType = 3
)

// enum
type ConnectionProfileDeleteStatus int32

const (
	ConnectionProfileDeleteStatus_Success        ConnectionProfileDeleteStatus = 0
	ConnectionProfileDeleteStatus_DeniedByUser   ConnectionProfileDeleteStatus = 1
	ConnectionProfileDeleteStatus_DeniedBySystem ConnectionProfileDeleteStatus = 2
	ConnectionProfileDeleteStatus_UnknownError   ConnectionProfileDeleteStatus = 3
)

// enum
type DataUsageGranularity int32

const (
	DataUsageGranularity_PerMinute DataUsageGranularity = 0
	DataUsageGranularity_PerHour   DataUsageGranularity = 1
	DataUsageGranularity_PerDay    DataUsageGranularity = 2
	DataUsageGranularity_Total     DataUsageGranularity = 3
)

// enum
type DomainConnectivityLevel int32

const (
	DomainConnectivityLevel_None            DomainConnectivityLevel = 0
	DomainConnectivityLevel_Unauthenticated DomainConnectivityLevel = 1
	DomainConnectivityLevel_Authenticated   DomainConnectivityLevel = 2
)

// enum
type NetworkAuthenticationType int32

const (
	NetworkAuthenticationType_None                  NetworkAuthenticationType = 0
	NetworkAuthenticationType_Unknown               NetworkAuthenticationType = 1
	NetworkAuthenticationType_Open80211             NetworkAuthenticationType = 2
	NetworkAuthenticationType_SharedKey80211        NetworkAuthenticationType = 3
	NetworkAuthenticationType_Wpa                   NetworkAuthenticationType = 4
	NetworkAuthenticationType_WpaPsk                NetworkAuthenticationType = 5
	NetworkAuthenticationType_WpaNone               NetworkAuthenticationType = 6
	NetworkAuthenticationType_Rsna                  NetworkAuthenticationType = 7
	NetworkAuthenticationType_RsnaPsk               NetworkAuthenticationType = 8
	NetworkAuthenticationType_Ihv                   NetworkAuthenticationType = 9
	NetworkAuthenticationType_Wpa3                  NetworkAuthenticationType = 10
	NetworkAuthenticationType_Wpa3Enterprise192Bits NetworkAuthenticationType = 10
	NetworkAuthenticationType_Wpa3Sae               NetworkAuthenticationType = 11
	NetworkAuthenticationType_Owe                   NetworkAuthenticationType = 12
	NetworkAuthenticationType_Wpa3Enterprise        NetworkAuthenticationType = 13
)

// enum
type NetworkConnectivityLevel int32

const (
	NetworkConnectivityLevel_None                      NetworkConnectivityLevel = 0
	NetworkConnectivityLevel_LocalAccess               NetworkConnectivityLevel = 1
	NetworkConnectivityLevel_ConstrainedInternetAccess NetworkConnectivityLevel = 2
	NetworkConnectivityLevel_InternetAccess            NetworkConnectivityLevel = 3
)

// enum
type NetworkCostType int32

const (
	NetworkCostType_Unknown      NetworkCostType = 0
	NetworkCostType_Unrestricted NetworkCostType = 1
	NetworkCostType_Fixed        NetworkCostType = 2
	NetworkCostType_Variable     NetworkCostType = 3
)

// enum
type NetworkEncryptionType int32

const (
	NetworkEncryptionType_None        NetworkEncryptionType = 0
	NetworkEncryptionType_Unknown     NetworkEncryptionType = 1
	NetworkEncryptionType_Wep         NetworkEncryptionType = 2
	NetworkEncryptionType_Wep40       NetworkEncryptionType = 3
	NetworkEncryptionType_Wep104      NetworkEncryptionType = 4
	NetworkEncryptionType_Tkip        NetworkEncryptionType = 5
	NetworkEncryptionType_Ccmp        NetworkEncryptionType = 6
	NetworkEncryptionType_WpaUseGroup NetworkEncryptionType = 7
	NetworkEncryptionType_RsnUseGroup NetworkEncryptionType = 8
	NetworkEncryptionType_Ihv         NetworkEncryptionType = 9
	NetworkEncryptionType_Gcmp        NetworkEncryptionType = 10
	NetworkEncryptionType_Gcmp256     NetworkEncryptionType = 11
)

// enum
// flags
type NetworkTypes uint32

const (
	NetworkTypes_None           NetworkTypes = 0
	NetworkTypes_Internet       NetworkTypes = 1
	NetworkTypes_PrivateNetwork NetworkTypes = 2
)

// enum
// flags
type RoamingStates uint32

const (
	RoamingStates_None       RoamingStates = 0
	RoamingStates_NotRoaming RoamingStates = 1
	RoamingStates_Roaming    RoamingStates = 2
)

// enum
type TriStates int32

const (
	TriStates_DoNotCare TriStates = 0
	TriStates_No        TriStates = 1
	TriStates_Yes       TriStates = 2
)

// enum
// flags
type WwanDataClass uint32

const (
	WwanDataClass_None           WwanDataClass = 0
	WwanDataClass_Gprs           WwanDataClass = 1
	WwanDataClass_Edge           WwanDataClass = 2
	WwanDataClass_Umts           WwanDataClass = 4
	WwanDataClass_Hsdpa          WwanDataClass = 8
	WwanDataClass_Hsupa          WwanDataClass = 16
	WwanDataClass_LteAdvanced    WwanDataClass = 32
	WwanDataClass_Cdma1xRtt      WwanDataClass = 65536
	WwanDataClass_Cdma1xEvdo     WwanDataClass = 131072
	WwanDataClass_Cdma1xEvdoRevA WwanDataClass = 262144
	WwanDataClass_Cdma1xEvdv     WwanDataClass = 524288
	WwanDataClass_Cdma3xRtt      WwanDataClass = 1048576
	WwanDataClass_Cdma1xEvdoRevB WwanDataClass = 2097152
	WwanDataClass_CdmaUmb        WwanDataClass = 4194304
	WwanDataClass_Custom         WwanDataClass = 2147483648
)

// enum
type WwanNetworkIPKind int32

const (
	WwanNetworkIPKind_None         WwanNetworkIPKind = 0
	WwanNetworkIPKind_Ipv4         WwanNetworkIPKind = 1
	WwanNetworkIPKind_Ipv6         WwanNetworkIPKind = 2
	WwanNetworkIPKind_Ipv4v6       WwanNetworkIPKind = 3
	WwanNetworkIPKind_Ipv4v6v4Xlat WwanNetworkIPKind = 4
)

// enum
type WwanNetworkRegistrationState int32

const (
	WwanNetworkRegistrationState_None         WwanNetworkRegistrationState = 0
	WwanNetworkRegistrationState_Deregistered WwanNetworkRegistrationState = 1
	WwanNetworkRegistrationState_Searching    WwanNetworkRegistrationState = 2
	WwanNetworkRegistrationState_Home         WwanNetworkRegistrationState = 3
	WwanNetworkRegistrationState_Roaming      WwanNetworkRegistrationState = 4
	WwanNetworkRegistrationState_Partner      WwanNetworkRegistrationState = 5
	WwanNetworkRegistrationState_Denied       WwanNetworkRegistrationState = 6
)

// structs

type NetworkUsageStates struct {
	Roaming TriStates
	Shared  TriStates
}

type WwanContract struct {
}

// func types

// 71BA143F-598E-49D0-84EB-8FEBAEDCC195
type NetworkStatusChangedEventHandler func(sender interface{}) com.Error

// interfaces

// F769B039-ECA2-45EB-ADE1-B0368B756C49
var IID_IAttributedNetworkUsage = syscall.GUID{0xF769B039, 0xECA2, 0x45EB,
	[8]byte{0xAD, 0xE1, 0xB0, 0x36, 0x8B, 0x75, 0x6C, 0x49}}

type IAttributedNetworkUsageInterface interface {
	win32.IInspectableInterface
	Get_BytesSent() uint64
	Get_BytesReceived() uint64
	Get_AttributionId() string
	Get_AttributionName() string
	Get_AttributionThumbnail() *IRandomAccessStreamReference
}

type IAttributedNetworkUsageVtbl struct {
	win32.IInspectableVtbl
	Get_BytesSent            uintptr
	Get_BytesReceived        uintptr
	Get_AttributionId        uintptr
	Get_AttributionName      uintptr
	Get_AttributionThumbnail uintptr
}

type IAttributedNetworkUsage struct {
	win32.IInspectable
}

func (this *IAttributedNetworkUsage) Vtbl() *IAttributedNetworkUsageVtbl {
	return (*IAttributedNetworkUsageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAttributedNetworkUsage) Get_BytesSent() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BytesSent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAttributedNetworkUsage) Get_BytesReceived() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BytesReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAttributedNetworkUsage) Get_AttributionId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AttributionId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAttributedNetworkUsage) Get_AttributionName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AttributionName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAttributedNetworkUsage) Get_AttributionThumbnail() *IRandomAccessStreamReference {
	var _result *IRandomAccessStreamReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AttributionThumbnail, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6FA529F4-EFFD-4542-9AB2-705BBF94943A
var IID_ICellularApnContext = syscall.GUID{0x6FA529F4, 0xEFFD, 0x4542,
	[8]byte{0x9A, 0xB2, 0x70, 0x5B, 0xBF, 0x94, 0x94, 0x3A}}

type ICellularApnContextInterface interface {
	win32.IInspectableInterface
	Get_ProviderId() string
	Put_ProviderId(value string)
	Get_AccessPointName() string
	Put_AccessPointName(value string)
	Get_UserName() string
	Put_UserName(value string)
	Get_Password() string
	Put_Password(value string)
	Get_IsCompressionEnabled() bool
	Put_IsCompressionEnabled(value bool)
	Get_AuthenticationType() CellularApnAuthenticationType
	Put_AuthenticationType(value CellularApnAuthenticationType)
}

type ICellularApnContextVtbl struct {
	win32.IInspectableVtbl
	Get_ProviderId           uintptr
	Put_ProviderId           uintptr
	Get_AccessPointName      uintptr
	Put_AccessPointName      uintptr
	Get_UserName             uintptr
	Put_UserName             uintptr
	Get_Password             uintptr
	Put_Password             uintptr
	Get_IsCompressionEnabled uintptr
	Put_IsCompressionEnabled uintptr
	Get_AuthenticationType   uintptr
	Put_AuthenticationType   uintptr
}

type ICellularApnContext struct {
	win32.IInspectable
}

func (this *ICellularApnContext) Vtbl() *ICellularApnContextVtbl {
	return (*ICellularApnContextVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICellularApnContext) Get_ProviderId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProviderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICellularApnContext) Put_ProviderId(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ProviderId, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICellularApnContext) Get_AccessPointName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AccessPointName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICellularApnContext) Put_AccessPointName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AccessPointName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICellularApnContext) Get_UserName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UserName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICellularApnContext) Put_UserName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_UserName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICellularApnContext) Get_Password() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Password, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICellularApnContext) Put_Password(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Password, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICellularApnContext) Get_IsCompressionEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCompressionEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICellularApnContext) Put_IsCompressionEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsCompressionEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICellularApnContext) Get_AuthenticationType() CellularApnAuthenticationType {
	var _result CellularApnAuthenticationType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AuthenticationType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICellularApnContext) Put_AuthenticationType(value CellularApnAuthenticationType) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AuthenticationType, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 76B0EB1A-AC49-4350-B1E5-DC4763BC69C7
var IID_ICellularApnContext2 = syscall.GUID{0x76B0EB1A, 0xAC49, 0x4350,
	[8]byte{0xB1, 0xE5, 0xDC, 0x47, 0x63, 0xBC, 0x69, 0xC7}}

type ICellularApnContext2Interface interface {
	win32.IInspectableInterface
	Get_ProfileName() string
	Put_ProfileName(value string)
}

type ICellularApnContext2Vtbl struct {
	win32.IInspectableVtbl
	Get_ProfileName uintptr
	Put_ProfileName uintptr
}

type ICellularApnContext2 struct {
	win32.IInspectable
}

func (this *ICellularApnContext2) Vtbl() *ICellularApnContext2Vtbl {
	return (*ICellularApnContext2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICellularApnContext2) Get_ProfileName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProfileName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICellularApnContext2) Put_ProfileName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ProfileName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// BAD7D829-3416-4B10-A202-BAC0B075BDAE
var IID_IConnectionCost = syscall.GUID{0xBAD7D829, 0x3416, 0x4B10,
	[8]byte{0xA2, 0x02, 0xBA, 0xC0, 0xB0, 0x75, 0xBD, 0xAE}}

type IConnectionCostInterface interface {
	win32.IInspectableInterface
	Get_NetworkCostType() NetworkCostType
	Get_Roaming() bool
	Get_OverDataLimit() bool
	Get_ApproachingDataLimit() bool
}

type IConnectionCostVtbl struct {
	win32.IInspectableVtbl
	Get_NetworkCostType      uintptr
	Get_Roaming              uintptr
	Get_OverDataLimit        uintptr
	Get_ApproachingDataLimit uintptr
}

type IConnectionCost struct {
	win32.IInspectable
}

func (this *IConnectionCost) Vtbl() *IConnectionCostVtbl {
	return (*IConnectionCostVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IConnectionCost) Get_NetworkCostType() NetworkCostType {
	var _result NetworkCostType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkCostType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IConnectionCost) Get_Roaming() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Roaming, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IConnectionCost) Get_OverDataLimit() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OverDataLimit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IConnectionCost) Get_ApproachingDataLimit() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ApproachingDataLimit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 8E113A05-E209-4549-BB25-5E0DB691CB05
var IID_IConnectionCost2 = syscall.GUID{0x8E113A05, 0xE209, 0x4549,
	[8]byte{0xBB, 0x25, 0x5E, 0x0D, 0xB6, 0x91, 0xCB, 0x05}}

type IConnectionCost2Interface interface {
	win32.IInspectableInterface
	Get_BackgroundDataUsageRestricted() bool
}

type IConnectionCost2Vtbl struct {
	win32.IInspectableVtbl
	Get_BackgroundDataUsageRestricted uintptr
}

type IConnectionCost2 struct {
	win32.IInspectable
}

func (this *IConnectionCost2) Vtbl() *IConnectionCost2Vtbl {
	return (*IConnectionCost2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IConnectionCost2) Get_BackgroundDataUsageRestricted() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BackgroundDataUsageRestricted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 71BA143C-598E-49D0-84EB-8FEBAEDCC195
var IID_IConnectionProfile = syscall.GUID{0x71BA143C, 0x598E, 0x49D0,
	[8]byte{0x84, 0xEB, 0x8F, 0xEB, 0xAE, 0xDC, 0xC1, 0x95}}

type IConnectionProfileInterface interface {
	win32.IInspectableInterface
	Get_ProfileName() string
	GetNetworkConnectivityLevel() NetworkConnectivityLevel
	GetNetworkNames() *IVectorView[string]
	GetConnectionCost() *IConnectionCost
	GetDataPlanStatus() *IDataPlanStatus
	Get_NetworkAdapter() *INetworkAdapter
	GetLocalUsage(StartTime DateTime, EndTime DateTime) *IDataUsage
	GetLocalUsagePerRoamingStates(StartTime DateTime, EndTime DateTime, States RoamingStates) *IDataUsage
	Get_NetworkSecuritySettings() *INetworkSecuritySettings
}

type IConnectionProfileVtbl struct {
	win32.IInspectableVtbl
	Get_ProfileName               uintptr
	GetNetworkConnectivityLevel   uintptr
	GetNetworkNames               uintptr
	GetConnectionCost             uintptr
	GetDataPlanStatus             uintptr
	Get_NetworkAdapter            uintptr
	GetLocalUsage                 uintptr
	GetLocalUsagePerRoamingStates uintptr
	Get_NetworkSecuritySettings   uintptr
}

type IConnectionProfile struct {
	win32.IInspectable
}

func (this *IConnectionProfile) Vtbl() *IConnectionProfileVtbl {
	return (*IConnectionProfileVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IConnectionProfile) Get_ProfileName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProfileName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IConnectionProfile) GetNetworkConnectivityLevel() NetworkConnectivityLevel {
	var _result NetworkConnectivityLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNetworkConnectivityLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IConnectionProfile) GetNetworkNames() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNetworkNames, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IConnectionProfile) GetConnectionCost() *IConnectionCost {
	var _result *IConnectionCost
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetConnectionCost, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IConnectionProfile) GetDataPlanStatus() *IDataPlanStatus {
	var _result *IDataPlanStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDataPlanStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IConnectionProfile) Get_NetworkAdapter() *INetworkAdapter {
	var _result *INetworkAdapter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkAdapter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IConnectionProfile) GetLocalUsage(StartTime DateTime, EndTime DateTime) *IDataUsage {
	var _result *IDataUsage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetLocalUsage, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&StartTime)), *(*uintptr)(unsafe.Pointer(&EndTime)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IConnectionProfile) GetLocalUsagePerRoamingStates(StartTime DateTime, EndTime DateTime, States RoamingStates) *IDataUsage {
	var _result *IDataUsage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetLocalUsagePerRoamingStates, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&StartTime)), *(*uintptr)(unsafe.Pointer(&EndTime)), uintptr(States), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IConnectionProfile) Get_NetworkSecuritySettings() *INetworkSecuritySettings {
	var _result *INetworkSecuritySettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkSecuritySettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E2045145-4C9F-400C-9150-7EC7D6E2888A
var IID_IConnectionProfile2 = syscall.GUID{0xE2045145, 0x4C9F, 0x400C,
	[8]byte{0x91, 0x50, 0x7E, 0xC7, 0xD6, 0xE2, 0x88, 0x8A}}

type IConnectionProfile2Interface interface {
	win32.IInspectableInterface
	Get_IsWwanConnectionProfile() bool
	Get_IsWlanConnectionProfile() bool
	Get_WwanConnectionProfileDetails() *IWwanConnectionProfileDetails
	Get_WlanConnectionProfileDetails() *IWlanConnectionProfileDetails
	Get_ServiceProviderGuid() *IReference[syscall.GUID]
	GetSignalBars() *IReference[byte]
	GetDomainConnectivityLevel() DomainConnectivityLevel
	GetNetworkUsageAsync(startTime DateTime, endTime DateTime, granularity DataUsageGranularity, states NetworkUsageStates) *IAsyncOperation[*IVectorView[*INetworkUsage]]
	GetConnectivityIntervalsAsync(startTime DateTime, endTime DateTime, states NetworkUsageStates) *IAsyncOperation[*IVectorView[*IConnectivityInterval]]
}

type IConnectionProfile2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsWwanConnectionProfile      uintptr
	Get_IsWlanConnectionProfile      uintptr
	Get_WwanConnectionProfileDetails uintptr
	Get_WlanConnectionProfileDetails uintptr
	Get_ServiceProviderGuid          uintptr
	GetSignalBars                    uintptr
	GetDomainConnectivityLevel       uintptr
	GetNetworkUsageAsync             uintptr
	GetConnectivityIntervalsAsync    uintptr
}

type IConnectionProfile2 struct {
	win32.IInspectable
}

func (this *IConnectionProfile2) Vtbl() *IConnectionProfile2Vtbl {
	return (*IConnectionProfile2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IConnectionProfile2) Get_IsWwanConnectionProfile() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsWwanConnectionProfile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IConnectionProfile2) Get_IsWlanConnectionProfile() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsWlanConnectionProfile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IConnectionProfile2) Get_WwanConnectionProfileDetails() *IWwanConnectionProfileDetails {
	var _result *IWwanConnectionProfileDetails
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WwanConnectionProfileDetails, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IConnectionProfile2) Get_WlanConnectionProfileDetails() *IWlanConnectionProfileDetails {
	var _result *IWlanConnectionProfileDetails
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WlanConnectionProfileDetails, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IConnectionProfile2) Get_ServiceProviderGuid() *IReference[syscall.GUID] {
	var _result *IReference[syscall.GUID]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceProviderGuid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IConnectionProfile2) GetSignalBars() *IReference[byte] {
	var _result *IReference[byte]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSignalBars, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IConnectionProfile2) GetDomainConnectivityLevel() DomainConnectivityLevel {
	var _result DomainConnectivityLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDomainConnectivityLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IConnectionProfile2) GetNetworkUsageAsync(startTime DateTime, endTime DateTime, granularity DataUsageGranularity, states NetworkUsageStates) *IAsyncOperation[*IVectorView[*INetworkUsage]] {
	var _result *IAsyncOperation[*IVectorView[*INetworkUsage]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNetworkUsageAsync, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&startTime)), *(*uintptr)(unsafe.Pointer(&endTime)), uintptr(granularity), *(*uintptr)(unsafe.Pointer(&states)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IConnectionProfile2) GetConnectivityIntervalsAsync(startTime DateTime, endTime DateTime, states NetworkUsageStates) *IAsyncOperation[*IVectorView[*IConnectivityInterval]] {
	var _result *IAsyncOperation[*IVectorView[*IConnectivityInterval]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetConnectivityIntervalsAsync, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&startTime)), *(*uintptr)(unsafe.Pointer(&endTime)), *(*uintptr)(unsafe.Pointer(&states)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 578C2528-4CD9-4161-8045-201CFD5B115C
var IID_IConnectionProfile3 = syscall.GUID{0x578C2528, 0x4CD9, 0x4161,
	[8]byte{0x80, 0x45, 0x20, 0x1C, 0xFD, 0x5B, 0x11, 0x5C}}

type IConnectionProfile3Interface interface {
	win32.IInspectableInterface
	GetAttributedNetworkUsageAsync(startTime DateTime, endTime DateTime, states NetworkUsageStates) *IAsyncOperation[*IVectorView[*IAttributedNetworkUsage]]
}

type IConnectionProfile3Vtbl struct {
	win32.IInspectableVtbl
	GetAttributedNetworkUsageAsync uintptr
}

type IConnectionProfile3 struct {
	win32.IInspectable
}

func (this *IConnectionProfile3) Vtbl() *IConnectionProfile3Vtbl {
	return (*IConnectionProfile3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IConnectionProfile3) GetAttributedNetworkUsageAsync(startTime DateTime, endTime DateTime, states NetworkUsageStates) *IAsyncOperation[*IVectorView[*IAttributedNetworkUsage]] {
	var _result *IAsyncOperation[*IVectorView[*IAttributedNetworkUsage]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAttributedNetworkUsageAsync, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&startTime)), *(*uintptr)(unsafe.Pointer(&endTime)), *(*uintptr)(unsafe.Pointer(&states)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7A2D42CD-81E0-4AE6-ABED-AB9CA13EB714
var IID_IConnectionProfile4 = syscall.GUID{0x7A2D42CD, 0x81E0, 0x4AE6,
	[8]byte{0xAB, 0xED, 0xAB, 0x9C, 0xA1, 0x3E, 0xB7, 0x14}}

type IConnectionProfile4Interface interface {
	win32.IInspectableInterface
	GetProviderNetworkUsageAsync(startTime DateTime, endTime DateTime, states NetworkUsageStates) *IAsyncOperation[*IVectorView[*IProviderNetworkUsage]]
}

type IConnectionProfile4Vtbl struct {
	win32.IInspectableVtbl
	GetProviderNetworkUsageAsync uintptr
}

type IConnectionProfile4 struct {
	win32.IInspectable
}

func (this *IConnectionProfile4) Vtbl() *IConnectionProfile4Vtbl {
	return (*IConnectionProfile4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IConnectionProfile4) GetProviderNetworkUsageAsync(startTime DateTime, endTime DateTime, states NetworkUsageStates) *IAsyncOperation[*IVectorView[*IProviderNetworkUsage]] {
	var _result *IAsyncOperation[*IVectorView[*IProviderNetworkUsage]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetProviderNetworkUsageAsync, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&startTime)), *(*uintptr)(unsafe.Pointer(&endTime)), *(*uintptr)(unsafe.Pointer(&states)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 85361EC7-9C73-4BE0-8F14-578EEC71EE0E
var IID_IConnectionProfile5 = syscall.GUID{0x85361EC7, 0x9C73, 0x4BE0,
	[8]byte{0x8F, 0x14, 0x57, 0x8E, 0xEC, 0x71, 0xEE, 0x0E}}

type IConnectionProfile5Interface interface {
	win32.IInspectableInterface
	Get_CanDelete() bool
	TryDeleteAsync() *IAsyncOperation[ConnectionProfileDeleteStatus]
}

type IConnectionProfile5Vtbl struct {
	win32.IInspectableVtbl
	Get_CanDelete  uintptr
	TryDeleteAsync uintptr
}

type IConnectionProfile5 struct {
	win32.IInspectable
}

func (this *IConnectionProfile5) Vtbl() *IConnectionProfile5Vtbl {
	return (*IConnectionProfile5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IConnectionProfile5) Get_CanDelete() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanDelete, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IConnectionProfile5) TryDeleteAsync() *IAsyncOperation[ConnectionProfileDeleteStatus] {
	var _result *IAsyncOperation[ConnectionProfileDeleteStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryDeleteAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 204C7CC8-BD2D-4E8D-A4B3-455EC337388A
var IID_IConnectionProfileFilter = syscall.GUID{0x204C7CC8, 0xBD2D, 0x4E8D,
	[8]byte{0xA4, 0xB3, 0x45, 0x5E, 0xC3, 0x37, 0x38, 0x8A}}

type IConnectionProfileFilterInterface interface {
	win32.IInspectableInterface
	Put_IsConnected(value bool)
	Get_IsConnected() bool
	Put_IsWwanConnectionProfile(value bool)
	Get_IsWwanConnectionProfile() bool
	Put_IsWlanConnectionProfile(value bool)
	Get_IsWlanConnectionProfile() bool
	Put_NetworkCostType(value NetworkCostType)
	Get_NetworkCostType() NetworkCostType
	Put_ServiceProviderGuid(value *IReference[syscall.GUID])
	Get_ServiceProviderGuid() *IReference[syscall.GUID]
}

type IConnectionProfileFilterVtbl struct {
	win32.IInspectableVtbl
	Put_IsConnected             uintptr
	Get_IsConnected             uintptr
	Put_IsWwanConnectionProfile uintptr
	Get_IsWwanConnectionProfile uintptr
	Put_IsWlanConnectionProfile uintptr
	Get_IsWlanConnectionProfile uintptr
	Put_NetworkCostType         uintptr
	Get_NetworkCostType         uintptr
	Put_ServiceProviderGuid     uintptr
	Get_ServiceProviderGuid     uintptr
}

type IConnectionProfileFilter struct {
	win32.IInspectable
}

func (this *IConnectionProfileFilter) Vtbl() *IConnectionProfileFilterVtbl {
	return (*IConnectionProfileFilterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IConnectionProfileFilter) Put_IsConnected(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsConnected, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IConnectionProfileFilter) Get_IsConnected() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsConnected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IConnectionProfileFilter) Put_IsWwanConnectionProfile(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsWwanConnectionProfile, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IConnectionProfileFilter) Get_IsWwanConnectionProfile() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsWwanConnectionProfile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IConnectionProfileFilter) Put_IsWlanConnectionProfile(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsWlanConnectionProfile, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IConnectionProfileFilter) Get_IsWlanConnectionProfile() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsWlanConnectionProfile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IConnectionProfileFilter) Put_NetworkCostType(value NetworkCostType) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_NetworkCostType, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IConnectionProfileFilter) Get_NetworkCostType() NetworkCostType {
	var _result NetworkCostType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkCostType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IConnectionProfileFilter) Put_ServiceProviderGuid(value *IReference[syscall.GUID]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ServiceProviderGuid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IConnectionProfileFilter) Get_ServiceProviderGuid() *IReference[syscall.GUID] {
	var _result *IReference[syscall.GUID]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceProviderGuid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CD068EE1-C3FC-4FAD-9DDC-593FAA4B7885
var IID_IConnectionProfileFilter2 = syscall.GUID{0xCD068EE1, 0xC3FC, 0x4FAD,
	[8]byte{0x9D, 0xDC, 0x59, 0x3F, 0xAA, 0x4B, 0x78, 0x85}}

type IConnectionProfileFilter2Interface interface {
	win32.IInspectableInterface
	Put_IsRoaming(value *IReference[bool])
	Get_IsRoaming() *IReference[bool]
	Put_IsOverDataLimit(value *IReference[bool])
	Get_IsOverDataLimit() *IReference[bool]
	Put_IsBackgroundDataUsageRestricted(value *IReference[bool])
	Get_IsBackgroundDataUsageRestricted() *IReference[bool]
	Get_RawData() *IBuffer
}

type IConnectionProfileFilter2Vtbl struct {
	win32.IInspectableVtbl
	Put_IsRoaming                       uintptr
	Get_IsRoaming                       uintptr
	Put_IsOverDataLimit                 uintptr
	Get_IsOverDataLimit                 uintptr
	Put_IsBackgroundDataUsageRestricted uintptr
	Get_IsBackgroundDataUsageRestricted uintptr
	Get_RawData                         uintptr
}

type IConnectionProfileFilter2 struct {
	win32.IInspectable
}

func (this *IConnectionProfileFilter2) Vtbl() *IConnectionProfileFilter2Vtbl {
	return (*IConnectionProfileFilter2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IConnectionProfileFilter2) Put_IsRoaming(value *IReference[bool]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsRoaming, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IConnectionProfileFilter2) Get_IsRoaming() *IReference[bool] {
	var _result *IReference[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsRoaming, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IConnectionProfileFilter2) Put_IsOverDataLimit(value *IReference[bool]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsOverDataLimit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IConnectionProfileFilter2) Get_IsOverDataLimit() *IReference[bool] {
	var _result *IReference[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsOverDataLimit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IConnectionProfileFilter2) Put_IsBackgroundDataUsageRestricted(value *IReference[bool]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsBackgroundDataUsageRestricted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IConnectionProfileFilter2) Get_IsBackgroundDataUsageRestricted() *IReference[bool] {
	var _result *IReference[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsBackgroundDataUsageRestricted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IConnectionProfileFilter2) Get_RawData() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RawData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0AAA09C0-5014-447C-8809-AEE4CB0AF94A
var IID_IConnectionProfileFilter3 = syscall.GUID{0x0AAA09C0, 0x5014, 0x447C,
	[8]byte{0x88, 0x09, 0xAE, 0xE4, 0xCB, 0x0A, 0xF9, 0x4A}}

type IConnectionProfileFilter3Interface interface {
	win32.IInspectableInterface
	Put_PurposeGuid(value *IReference[syscall.GUID])
	Get_PurposeGuid() *IReference[syscall.GUID]
}

type IConnectionProfileFilter3Vtbl struct {
	win32.IInspectableVtbl
	Put_PurposeGuid uintptr
	Get_PurposeGuid uintptr
}

type IConnectionProfileFilter3 struct {
	win32.IInspectable
}

func (this *IConnectionProfileFilter3) Vtbl() *IConnectionProfileFilter3Vtbl {
	return (*IConnectionProfileFilter3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IConnectionProfileFilter3) Put_PurposeGuid(value *IReference[syscall.GUID]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PurposeGuid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IConnectionProfileFilter3) Get_PurposeGuid() *IReference[syscall.GUID] {
	var _result *IReference[syscall.GUID]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PurposeGuid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FF905D4C-F83B-41B0-8A0C-1462D9C56B73
var IID_IConnectionSession = syscall.GUID{0xFF905D4C, 0xF83B, 0x41B0,
	[8]byte{0x8A, 0x0C, 0x14, 0x62, 0xD9, 0xC5, 0x6B, 0x73}}

type IConnectionSessionInterface interface {
	win32.IInspectableInterface
	Get_ConnectionProfile() *IConnectionProfile
}

type IConnectionSessionVtbl struct {
	win32.IInspectableVtbl
	Get_ConnectionProfile uintptr
}

type IConnectionSession struct {
	win32.IInspectable
}

func (this *IConnectionSession) Vtbl() *IConnectionSessionVtbl {
	return (*IConnectionSessionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IConnectionSession) Get_ConnectionProfile() *IConnectionProfile {
	var _result *IConnectionProfile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConnectionProfile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4FAA3FFF-6746-4824-A964-EED8E87F8709
var IID_IConnectivityInterval = syscall.GUID{0x4FAA3FFF, 0x6746, 0x4824,
	[8]byte{0xA9, 0x64, 0xEE, 0xD8, 0xE8, 0x7F, 0x87, 0x09}}

type IConnectivityIntervalInterface interface {
	win32.IInspectableInterface
	Get_StartTime() DateTime
	Get_ConnectionDuration() TimeSpan
}

type IConnectivityIntervalVtbl struct {
	win32.IInspectableVtbl
	Get_StartTime          uintptr
	Get_ConnectionDuration uintptr
}

type IConnectivityInterval struct {
	win32.IInspectable
}

func (this *IConnectivityInterval) Vtbl() *IConnectivityIntervalVtbl {
	return (*IConnectivityIntervalVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IConnectivityInterval) Get_StartTime() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StartTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IConnectivityInterval) Get_ConnectionDuration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConnectionDuration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5120D4B1-4FB1-48B0-AFC9-42E0092A8164
var IID_IConnectivityManagerStatics = syscall.GUID{0x5120D4B1, 0x4FB1, 0x48B0,
	[8]byte{0xAF, 0xC9, 0x42, 0xE0, 0x09, 0x2A, 0x81, 0x64}}

type IConnectivityManagerStaticsInterface interface {
	win32.IInspectableInterface
	AcquireConnectionAsync(cellularApnContext *ICellularApnContext) *IAsyncOperation[*IConnectionSession]
	AddHttpRoutePolicy(routePolicy *IRoutePolicy)
	RemoveHttpRoutePolicy(routePolicy *IRoutePolicy)
}

type IConnectivityManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	AcquireConnectionAsync uintptr
	AddHttpRoutePolicy     uintptr
	RemoveHttpRoutePolicy  uintptr
}

type IConnectivityManagerStatics struct {
	win32.IInspectable
}

func (this *IConnectivityManagerStatics) Vtbl() *IConnectivityManagerStaticsVtbl {
	return (*IConnectivityManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IConnectivityManagerStatics) AcquireConnectionAsync(cellularApnContext *ICellularApnContext) *IAsyncOperation[*IConnectionSession] {
	var _result *IAsyncOperation[*IConnectionSession]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AcquireConnectionAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(cellularApnContext)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IConnectivityManagerStatics) AddHttpRoutePolicy(routePolicy *IRoutePolicy) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddHttpRoutePolicy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(routePolicy)))
	_ = _hr
}

func (this *IConnectivityManagerStatics) RemoveHttpRoutePolicy(routePolicy *IRoutePolicy) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveHttpRoutePolicy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(routePolicy)))
	_ = _hr
}

// 977A8B8C-3885-40F3-8851-42CD2BD568BB
var IID_IDataPlanStatus = syscall.GUID{0x977A8B8C, 0x3885, 0x40F3,
	[8]byte{0x88, 0x51, 0x42, 0xCD, 0x2B, 0xD5, 0x68, 0xBB}}

type IDataPlanStatusInterface interface {
	win32.IInspectableInterface
	Get_DataPlanUsage() *IDataPlanUsage
	Get_DataLimitInMegabytes() *IReference[uint32]
	Get_InboundBitsPerSecond() *IReference[uint64]
	Get_OutboundBitsPerSecond() *IReference[uint64]
	Get_NextBillingCycle() *IReference[DateTime]
	Get_MaxTransferSizeInMegabytes() *IReference[uint32]
}

type IDataPlanStatusVtbl struct {
	win32.IInspectableVtbl
	Get_DataPlanUsage              uintptr
	Get_DataLimitInMegabytes       uintptr
	Get_InboundBitsPerSecond       uintptr
	Get_OutboundBitsPerSecond      uintptr
	Get_NextBillingCycle           uintptr
	Get_MaxTransferSizeInMegabytes uintptr
}

type IDataPlanStatus struct {
	win32.IInspectable
}

func (this *IDataPlanStatus) Vtbl() *IDataPlanStatusVtbl {
	return (*IDataPlanStatusVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataPlanStatus) Get_DataPlanUsage() *IDataPlanUsage {
	var _result *IDataPlanUsage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DataPlanUsage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPlanStatus) Get_DataLimitInMegabytes() *IReference[uint32] {
	var _result *IReference[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DataLimitInMegabytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPlanStatus) Get_InboundBitsPerSecond() *IReference[uint64] {
	var _result *IReference[uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InboundBitsPerSecond, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPlanStatus) Get_OutboundBitsPerSecond() *IReference[uint64] {
	var _result *IReference[uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutboundBitsPerSecond, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPlanStatus) Get_NextBillingCycle() *IReference[DateTime] {
	var _result *IReference[DateTime]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NextBillingCycle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataPlanStatus) Get_MaxTransferSizeInMegabytes() *IReference[uint32] {
	var _result *IReference[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxTransferSizeInMegabytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B921492D-3B44-47FF-B361-BE59E69ED1B0
var IID_IDataPlanUsage = syscall.GUID{0xB921492D, 0x3B44, 0x47FF,
	[8]byte{0xB3, 0x61, 0xBE, 0x59, 0xE6, 0x9E, 0xD1, 0xB0}}

type IDataPlanUsageInterface interface {
	win32.IInspectableInterface
	Get_MegabytesUsed() uint32
	Get_LastSyncTime() DateTime
}

type IDataPlanUsageVtbl struct {
	win32.IInspectableVtbl
	Get_MegabytesUsed uintptr
	Get_LastSyncTime  uintptr
}

type IDataPlanUsage struct {
	win32.IInspectable
}

func (this *IDataPlanUsage) Vtbl() *IDataPlanUsageVtbl {
	return (*IDataPlanUsageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataPlanUsage) Get_MegabytesUsed() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MegabytesUsed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataPlanUsage) Get_LastSyncTime() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LastSyncTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C1431DD3-B146-4D39-B959-0C69B096C512
var IID_IDataUsage = syscall.GUID{0xC1431DD3, 0xB146, 0x4D39,
	[8]byte{0xB9, 0x59, 0x0C, 0x69, 0xB0, 0x96, 0xC5, 0x12}}

type IDataUsageInterface interface {
	win32.IInspectableInterface
	Get_BytesSent() uint64
	Get_BytesReceived() uint64
}

type IDataUsageVtbl struct {
	win32.IInspectableVtbl
	Get_BytesSent     uintptr
	Get_BytesReceived uintptr
}

type IDataUsage struct {
	win32.IInspectable
}

func (this *IDataUsage) Vtbl() *IDataUsageVtbl {
	return (*IDataUsageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataUsage) Get_BytesSent() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BytesSent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataUsage) Get_BytesReceived() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BytesReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D85145E0-138F-47D7-9B3A-36BB488CEF33
var IID_IIPInformation = syscall.GUID{0xD85145E0, 0x138F, 0x47D7,
	[8]byte{0x9B, 0x3A, 0x36, 0xBB, 0x48, 0x8C, 0xEF, 0x33}}

type IIPInformationInterface interface {
	win32.IInspectableInterface
	Get_NetworkAdapter() *INetworkAdapter
	Get_PrefixLength() *IReference[byte]
}

type IIPInformationVtbl struct {
	win32.IInspectableVtbl
	Get_NetworkAdapter uintptr
	Get_PrefixLength   uintptr
}

type IIPInformation struct {
	win32.IInspectable
}

func (this *IIPInformation) Vtbl() *IIPInformationVtbl {
	return (*IIPInformationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IIPInformation) Get_NetworkAdapter() *INetworkAdapter {
	var _result *INetworkAdapter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkAdapter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IIPInformation) Get_PrefixLength() *IReference[byte] {
	var _result *IReference[byte]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PrefixLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 48AA53AA-1108-4546-A6CB-9A74DA4B7BA0
var IID_ILanIdentifier = syscall.GUID{0x48AA53AA, 0x1108, 0x4546,
	[8]byte{0xA6, 0xCB, 0x9A, 0x74, 0xDA, 0x4B, 0x7B, 0xA0}}

type ILanIdentifierInterface interface {
	win32.IInspectableInterface
	Get_InfrastructureId() *ILanIdentifierData
	Get_PortId() *ILanIdentifierData
	Get_NetworkAdapterId() syscall.GUID
}

type ILanIdentifierVtbl struct {
	win32.IInspectableVtbl
	Get_InfrastructureId uintptr
	Get_PortId           uintptr
	Get_NetworkAdapterId uintptr
}

type ILanIdentifier struct {
	win32.IInspectable
}

func (this *ILanIdentifier) Vtbl() *ILanIdentifierVtbl {
	return (*ILanIdentifierVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILanIdentifier) Get_InfrastructureId() *ILanIdentifierData {
	var _result *ILanIdentifierData
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InfrastructureId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILanIdentifier) Get_PortId() *ILanIdentifierData {
	var _result *ILanIdentifierData
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PortId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILanIdentifier) Get_NetworkAdapterId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkAdapterId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A74E83C3-D639-45BE-A36A-C4E4AEAF6D9B
var IID_ILanIdentifierData = syscall.GUID{0xA74E83C3, 0xD639, 0x45BE,
	[8]byte{0xA3, 0x6A, 0xC4, 0xE4, 0xAE, 0xAF, 0x6D, 0x9B}}

type ILanIdentifierDataInterface interface {
	win32.IInspectableInterface
	Get_Type() uint32
	Get_Value() *IVectorView[byte]
}

type ILanIdentifierDataVtbl struct {
	win32.IInspectableVtbl
	Get_Type  uintptr
	Get_Value uintptr
}

type ILanIdentifierData struct {
	win32.IInspectable
}

func (this *ILanIdentifierData) Vtbl() *ILanIdentifierDataVtbl {
	return (*ILanIdentifierDataVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILanIdentifierData) Get_Type() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Type, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILanIdentifierData) Get_Value() *IVectorView[byte] {
	var _result *IVectorView[byte]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3B542E03-5388-496C-A8A3-AFFD39AEC2E6
var IID_INetworkAdapter = syscall.GUID{0x3B542E03, 0x5388, 0x496C,
	[8]byte{0xA8, 0xA3, 0xAF, 0xFD, 0x39, 0xAE, 0xC2, 0xE6}}

type INetworkAdapterInterface interface {
	win32.IInspectableInterface
	Get_OutboundMaxBitsPerSecond() uint64
	Get_InboundMaxBitsPerSecond() uint64
	Get_IanaInterfaceType() uint32
	Get_NetworkItem() *INetworkItem
	Get_NetworkAdapterId() syscall.GUID
	GetConnectedProfileAsync() *IAsyncOperation[*IConnectionProfile]
}

type INetworkAdapterVtbl struct {
	win32.IInspectableVtbl
	Get_OutboundMaxBitsPerSecond uintptr
	Get_InboundMaxBitsPerSecond  uintptr
	Get_IanaInterfaceType        uintptr
	Get_NetworkItem              uintptr
	Get_NetworkAdapterId         uintptr
	GetConnectedProfileAsync     uintptr
}

type INetworkAdapter struct {
	win32.IInspectable
}

func (this *INetworkAdapter) Vtbl() *INetworkAdapterVtbl {
	return (*INetworkAdapterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INetworkAdapter) Get_OutboundMaxBitsPerSecond() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutboundMaxBitsPerSecond, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INetworkAdapter) Get_InboundMaxBitsPerSecond() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InboundMaxBitsPerSecond, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INetworkAdapter) Get_IanaInterfaceType() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IanaInterfaceType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INetworkAdapter) Get_NetworkItem() *INetworkItem {
	var _result *INetworkItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkItem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *INetworkAdapter) Get_NetworkAdapterId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkAdapterId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INetworkAdapter) GetConnectedProfileAsync() *IAsyncOperation[*IConnectionProfile] {
	var _result *IAsyncOperation[*IConnectionProfile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetConnectedProfileAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5074F851-950D-4165-9C15-365619481EEA
var IID_INetworkInformationStatics = syscall.GUID{0x5074F851, 0x950D, 0x4165,
	[8]byte{0x9C, 0x15, 0x36, 0x56, 0x19, 0x48, 0x1E, 0xEA}}

type INetworkInformationStaticsInterface interface {
	win32.IInspectableInterface
	GetConnectionProfiles() *IVectorView[*IConnectionProfile]
	GetInternetConnectionProfile() *IConnectionProfile
	GetLanIdentifiers() *IVectorView[*ILanIdentifier]
	GetHostNames() *IVectorView[*IHostName]
	GetProxyConfigurationAsync(uri *IUriRuntimeClass) *IAsyncOperation[*IProxyConfiguration]
	GetSortedEndpointPairs(destinationList *IIterable[*IEndpointPair], sortOptions HostNameSortOptions) *IVectorView[*IEndpointPair]
	Add_NetworkStatusChanged(networkStatusHandler NetworkStatusChangedEventHandler) EventRegistrationToken
	Remove_NetworkStatusChanged(eventCookie EventRegistrationToken)
}

type INetworkInformationStaticsVtbl struct {
	win32.IInspectableVtbl
	GetConnectionProfiles        uintptr
	GetInternetConnectionProfile uintptr
	GetLanIdentifiers            uintptr
	GetHostNames                 uintptr
	GetProxyConfigurationAsync   uintptr
	GetSortedEndpointPairs       uintptr
	Add_NetworkStatusChanged     uintptr
	Remove_NetworkStatusChanged  uintptr
}

type INetworkInformationStatics struct {
	win32.IInspectable
}

func (this *INetworkInformationStatics) Vtbl() *INetworkInformationStaticsVtbl {
	return (*INetworkInformationStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INetworkInformationStatics) GetConnectionProfiles() *IVectorView[*IConnectionProfile] {
	var _result *IVectorView[*IConnectionProfile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetConnectionProfiles, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *INetworkInformationStatics) GetInternetConnectionProfile() *IConnectionProfile {
	var _result *IConnectionProfile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetInternetConnectionProfile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *INetworkInformationStatics) GetLanIdentifiers() *IVectorView[*ILanIdentifier] {
	var _result *IVectorView[*ILanIdentifier]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetLanIdentifiers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *INetworkInformationStatics) GetHostNames() *IVectorView[*IHostName] {
	var _result *IVectorView[*IHostName]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetHostNames, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *INetworkInformationStatics) GetProxyConfigurationAsync(uri *IUriRuntimeClass) *IAsyncOperation[*IProxyConfiguration] {
	var _result *IAsyncOperation[*IProxyConfiguration]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetProxyConfigurationAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *INetworkInformationStatics) GetSortedEndpointPairs(destinationList *IIterable[*IEndpointPair], sortOptions HostNameSortOptions) *IVectorView[*IEndpointPair] {
	var _result *IVectorView[*IEndpointPair]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSortedEndpointPairs, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(destinationList)), uintptr(sortOptions), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *INetworkInformationStatics) Add_NetworkStatusChanged(networkStatusHandler NetworkStatusChangedEventHandler) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_NetworkStatusChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewOneArgFuncDelegate(networkStatusHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INetworkInformationStatics) Remove_NetworkStatusChanged(eventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_NetworkStatusChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&eventCookie)))
	_ = _hr
}

// 459CED14-2832-49B6-BA6E-E265F04786A8
var IID_INetworkInformationStatics2 = syscall.GUID{0x459CED14, 0x2832, 0x49B6,
	[8]byte{0xBA, 0x6E, 0xE2, 0x65, 0xF0, 0x47, 0x86, 0xA8}}

type INetworkInformationStatics2Interface interface {
	win32.IInspectableInterface
	FindConnectionProfilesAsync(pProfileFilter *IConnectionProfileFilter) *IAsyncOperation[*IVectorView[*IConnectionProfile]]
}

type INetworkInformationStatics2Vtbl struct {
	win32.IInspectableVtbl
	FindConnectionProfilesAsync uintptr
}

type INetworkInformationStatics2 struct {
	win32.IInspectable
}

func (this *INetworkInformationStatics2) Vtbl() *INetworkInformationStatics2Vtbl {
	return (*INetworkInformationStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INetworkInformationStatics2) FindConnectionProfilesAsync(pProfileFilter *IConnectionProfileFilter) *IAsyncOperation[*IVectorView[*IConnectionProfile]] {
	var _result *IAsyncOperation[*IVectorView[*IConnectionProfile]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindConnectionProfilesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pProfileFilter)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 01BC4D39-F5E0-4567-A28C-42080C831B2B
var IID_INetworkItem = syscall.GUID{0x01BC4D39, 0xF5E0, 0x4567,
	[8]byte{0xA2, 0x8C, 0x42, 0x08, 0x0C, 0x83, 0x1B, 0x2B}}

type INetworkItemInterface interface {
	win32.IInspectableInterface
	Get_NetworkId() syscall.GUID
	GetNetworkTypes() NetworkTypes
}

type INetworkItemVtbl struct {
	win32.IInspectableVtbl
	Get_NetworkId   uintptr
	GetNetworkTypes uintptr
}

type INetworkItem struct {
	win32.IInspectable
}

func (this *INetworkItem) Vtbl() *INetworkItemVtbl {
	return (*INetworkItemVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INetworkItem) Get_NetworkId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INetworkItem) GetNetworkTypes() NetworkTypes {
	var _result NetworkTypes
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNetworkTypes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7CA07E8D-917B-4B5F-B84D-28F7A5AC5402
var IID_INetworkSecuritySettings = syscall.GUID{0x7CA07E8D, 0x917B, 0x4B5F,
	[8]byte{0xB8, 0x4D, 0x28, 0xF7, 0xA5, 0xAC, 0x54, 0x02}}

type INetworkSecuritySettingsInterface interface {
	win32.IInspectableInterface
	Get_NetworkAuthenticationType() NetworkAuthenticationType
	Get_NetworkEncryptionType() NetworkEncryptionType
}

type INetworkSecuritySettingsVtbl struct {
	win32.IInspectableVtbl
	Get_NetworkAuthenticationType uintptr
	Get_NetworkEncryptionType     uintptr
}

type INetworkSecuritySettings struct {
	win32.IInspectable
}

func (this *INetworkSecuritySettings) Vtbl() *INetworkSecuritySettingsVtbl {
	return (*INetworkSecuritySettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INetworkSecuritySettings) Get_NetworkAuthenticationType() NetworkAuthenticationType {
	var _result NetworkAuthenticationType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkAuthenticationType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INetworkSecuritySettings) Get_NetworkEncryptionType() NetworkEncryptionType {
	var _result NetworkEncryptionType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkEncryptionType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 1F0CF333-D7A6-44DD-A4E9-687C476B903D
var IID_INetworkStateChangeEventDetails = syscall.GUID{0x1F0CF333, 0xD7A6, 0x44DD,
	[8]byte{0xA4, 0xE9, 0x68, 0x7C, 0x47, 0x6B, 0x90, 0x3D}}

type INetworkStateChangeEventDetailsInterface interface {
	win32.IInspectableInterface
	Get_HasNewInternetConnectionProfile() bool
	Get_HasNewConnectionCost() bool
	Get_HasNewNetworkConnectivityLevel() bool
	Get_HasNewDomainConnectivityLevel() bool
	Get_HasNewHostNameList() bool
	Get_HasNewWwanRegistrationState() bool
}

type INetworkStateChangeEventDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_HasNewInternetConnectionProfile uintptr
	Get_HasNewConnectionCost            uintptr
	Get_HasNewNetworkConnectivityLevel  uintptr
	Get_HasNewDomainConnectivityLevel   uintptr
	Get_HasNewHostNameList              uintptr
	Get_HasNewWwanRegistrationState     uintptr
}

type INetworkStateChangeEventDetails struct {
	win32.IInspectable
}

func (this *INetworkStateChangeEventDetails) Vtbl() *INetworkStateChangeEventDetailsVtbl {
	return (*INetworkStateChangeEventDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INetworkStateChangeEventDetails) Get_HasNewInternetConnectionProfile() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasNewInternetConnectionProfile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INetworkStateChangeEventDetails) Get_HasNewConnectionCost() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasNewConnectionCost, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INetworkStateChangeEventDetails) Get_HasNewNetworkConnectivityLevel() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasNewNetworkConnectivityLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INetworkStateChangeEventDetails) Get_HasNewDomainConnectivityLevel() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasNewDomainConnectivityLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INetworkStateChangeEventDetails) Get_HasNewHostNameList() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasNewHostNameList, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INetworkStateChangeEventDetails) Get_HasNewWwanRegistrationState() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasNewWwanRegistrationState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D643C0E8-30D3-4F6A-AD47-6A1873CEB3C1
var IID_INetworkStateChangeEventDetails2 = syscall.GUID{0xD643C0E8, 0x30D3, 0x4F6A,
	[8]byte{0xAD, 0x47, 0x6A, 0x18, 0x73, 0xCE, 0xB3, 0xC1}}

type INetworkStateChangeEventDetails2Interface interface {
	win32.IInspectableInterface
	Get_HasNewTetheringOperationalState() bool
	Get_HasNewTetheringClientCount() bool
}

type INetworkStateChangeEventDetails2Vtbl struct {
	win32.IInspectableVtbl
	Get_HasNewTetheringOperationalState uintptr
	Get_HasNewTetheringClientCount      uintptr
}

type INetworkStateChangeEventDetails2 struct {
	win32.IInspectable
}

func (this *INetworkStateChangeEventDetails2) Vtbl() *INetworkStateChangeEventDetails2Vtbl {
	return (*INetworkStateChangeEventDetails2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INetworkStateChangeEventDetails2) Get_HasNewTetheringOperationalState() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasNewTetheringOperationalState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INetworkStateChangeEventDetails2) Get_HasNewTetheringClientCount() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasNewTetheringClientCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 49DA8FCE-9985-4927-BF5B-072B5C65F8D9
var IID_INetworkUsage = syscall.GUID{0x49DA8FCE, 0x9985, 0x4927,
	[8]byte{0xBF, 0x5B, 0x07, 0x2B, 0x5C, 0x65, 0xF8, 0xD9}}

type INetworkUsageInterface interface {
	win32.IInspectableInterface
	Get_BytesSent() uint64
	Get_BytesReceived() uint64
	Get_ConnectionDuration() TimeSpan
}

type INetworkUsageVtbl struct {
	win32.IInspectableVtbl
	Get_BytesSent          uintptr
	Get_BytesReceived      uintptr
	Get_ConnectionDuration uintptr
}

type INetworkUsage struct {
	win32.IInspectable
}

func (this *INetworkUsage) Vtbl() *INetworkUsageVtbl {
	return (*INetworkUsageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INetworkUsage) Get_BytesSent() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BytesSent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INetworkUsage) Get_BytesReceived() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BytesReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INetworkUsage) Get_ConnectionDuration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConnectionDuration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5EC69E04-7931-48C8-B8F3-46300FA42728
var IID_IProviderNetworkUsage = syscall.GUID{0x5EC69E04, 0x7931, 0x48C8,
	[8]byte{0xB8, 0xF3, 0x46, 0x30, 0x0F, 0xA4, 0x27, 0x28}}

type IProviderNetworkUsageInterface interface {
	win32.IInspectableInterface
	Get_BytesSent() uint64
	Get_BytesReceived() uint64
	Get_ProviderId() string
}

type IProviderNetworkUsageVtbl struct {
	win32.IInspectableVtbl
	Get_BytesSent     uintptr
	Get_BytesReceived uintptr
	Get_ProviderId    uintptr
}

type IProviderNetworkUsage struct {
	win32.IInspectable
}

func (this *IProviderNetworkUsage) Vtbl() *IProviderNetworkUsageVtbl {
	return (*IProviderNetworkUsageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProviderNetworkUsage) Get_BytesSent() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BytesSent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProviderNetworkUsage) Get_BytesReceived() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BytesReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProviderNetworkUsage) Get_ProviderId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProviderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// EF3A60B4-9004-4DD6-B7D8-B3E502F4AAD0
var IID_IProxyConfiguration = syscall.GUID{0xEF3A60B4, 0x9004, 0x4DD6,
	[8]byte{0xB7, 0xD8, 0xB3, 0xE5, 0x02, 0xF4, 0xAA, 0xD0}}

type IProxyConfigurationInterface interface {
	win32.IInspectableInterface
	Get_ProxyUris() *IVectorView[*IUriRuntimeClass]
	Get_CanConnectDirectly() bool
}

type IProxyConfigurationVtbl struct {
	win32.IInspectableVtbl
	Get_ProxyUris          uintptr
	Get_CanConnectDirectly uintptr
}

type IProxyConfiguration struct {
	win32.IInspectable
}

func (this *IProxyConfiguration) Vtbl() *IProxyConfigurationVtbl {
	return (*IProxyConfigurationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProxyConfiguration) Get_ProxyUris() *IVectorView[*IUriRuntimeClass] {
	var _result *IVectorView[*IUriRuntimeClass]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProxyUris, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProxyConfiguration) Get_CanConnectDirectly() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanConnectDirectly, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 11ABC4AC-0FC7-42E4-8742-569923B1CA11
var IID_IRoutePolicy = syscall.GUID{0x11ABC4AC, 0x0FC7, 0x42E4,
	[8]byte{0x87, 0x42, 0x56, 0x99, 0x23, 0xB1, 0xCA, 0x11}}

type IRoutePolicyInterface interface {
	win32.IInspectableInterface
	Get_ConnectionProfile() *IConnectionProfile
	Get_HostName() *IHostName
	Get_HostNameType() DomainNameType
}

type IRoutePolicyVtbl struct {
	win32.IInspectableVtbl
	Get_ConnectionProfile uintptr
	Get_HostName          uintptr
	Get_HostNameType      uintptr
}

type IRoutePolicy struct {
	win32.IInspectable
}

func (this *IRoutePolicy) Vtbl() *IRoutePolicyVtbl {
	return (*IRoutePolicyVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRoutePolicy) Get_ConnectionProfile() *IConnectionProfile {
	var _result *IConnectionProfile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConnectionProfile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRoutePolicy) Get_HostName() *IHostName {
	var _result *IHostName
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HostName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRoutePolicy) Get_HostNameType() DomainNameType {
	var _result DomainNameType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HostNameType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 36027933-A18E-4DB5-A697-F58FA7364E44
var IID_IRoutePolicyFactory = syscall.GUID{0x36027933, 0xA18E, 0x4DB5,
	[8]byte{0xA6, 0x97, 0xF5, 0x8F, 0xA7, 0x36, 0x4E, 0x44}}

type IRoutePolicyFactoryInterface interface {
	win32.IInspectableInterface
	CreateRoutePolicy(connectionProfile *IConnectionProfile, hostName *IHostName, type_ DomainNameType) *IRoutePolicy
}

type IRoutePolicyFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateRoutePolicy uintptr
}

type IRoutePolicyFactory struct {
	win32.IInspectable
}

func (this *IRoutePolicyFactory) Vtbl() *IRoutePolicyFactoryVtbl {
	return (*IRoutePolicyFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRoutePolicyFactory) CreateRoutePolicy(connectionProfile *IConnectionProfile, hostName *IHostName, type_ DomainNameType) *IRoutePolicy {
	var _result *IRoutePolicy
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateRoutePolicy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(connectionProfile)), uintptr(unsafe.Pointer(hostName)), uintptr(type_), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 562098CB-B35A-4BF1-A884-B7557E88FF86
var IID_IWlanConnectionProfileDetails = syscall.GUID{0x562098CB, 0xB35A, 0x4BF1,
	[8]byte{0xA8, 0x84, 0xB7, 0x55, 0x7E, 0x88, 0xFF, 0x86}}

type IWlanConnectionProfileDetailsInterface interface {
	win32.IInspectableInterface
	GetConnectedSsid() string
}

type IWlanConnectionProfileDetailsVtbl struct {
	win32.IInspectableVtbl
	GetConnectedSsid uintptr
}

type IWlanConnectionProfileDetails struct {
	win32.IInspectable
}

func (this *IWlanConnectionProfileDetails) Vtbl() *IWlanConnectionProfileDetailsVtbl {
	return (*IWlanConnectionProfileDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWlanConnectionProfileDetails) GetConnectedSsid() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetConnectedSsid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 0E4DA8FE-835F-4DF3-82FD-DF556EBC09EF
var IID_IWwanConnectionProfileDetails = syscall.GUID{0x0E4DA8FE, 0x835F, 0x4DF3,
	[8]byte{0x82, 0xFD, 0xDF, 0x55, 0x6E, 0xBC, 0x09, 0xEF}}

type IWwanConnectionProfileDetailsInterface interface {
	win32.IInspectableInterface
	Get_HomeProviderId() string
	Get_AccessPointName() string
	GetNetworkRegistrationState() WwanNetworkRegistrationState
	GetCurrentDataClass() WwanDataClass
}

type IWwanConnectionProfileDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_HomeProviderId          uintptr
	Get_AccessPointName         uintptr
	GetNetworkRegistrationState uintptr
	GetCurrentDataClass         uintptr
}

type IWwanConnectionProfileDetails struct {
	win32.IInspectable
}

func (this *IWwanConnectionProfileDetails) Vtbl() *IWwanConnectionProfileDetailsVtbl {
	return (*IWwanConnectionProfileDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWwanConnectionProfileDetails) Get_HomeProviderId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HomeProviderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWwanConnectionProfileDetails) Get_AccessPointName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AccessPointName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWwanConnectionProfileDetails) GetNetworkRegistrationState() WwanNetworkRegistrationState {
	var _result WwanNetworkRegistrationState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetNetworkRegistrationState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWwanConnectionProfileDetails) GetCurrentDataClass() WwanDataClass {
	var _result WwanDataClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentDataClass, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7A754EDE-A1ED-48B2-8E92-B460033D52E2
var IID_IWwanConnectionProfileDetails2 = syscall.GUID{0x7A754EDE, 0xA1ED, 0x48B2,
	[8]byte{0x8E, 0x92, 0xB4, 0x60, 0x03, 0x3D, 0x52, 0xE2}}

type IWwanConnectionProfileDetails2Interface interface {
	win32.IInspectableInterface
	Get_IPKind() WwanNetworkIPKind
	Get_PurposeGuids() *IVectorView[syscall.GUID]
}

type IWwanConnectionProfileDetails2Vtbl struct {
	win32.IInspectableVtbl
	Get_IPKind       uintptr
	Get_PurposeGuids uintptr
}

type IWwanConnectionProfileDetails2 struct {
	win32.IInspectable
}

func (this *IWwanConnectionProfileDetails2) Vtbl() *IWwanConnectionProfileDetails2Vtbl {
	return (*IWwanConnectionProfileDetails2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWwanConnectionProfileDetails2) Get_IPKind() WwanNetworkIPKind {
	var _result WwanNetworkIPKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IPKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWwanConnectionProfileDetails2) Get_PurposeGuids() *IVectorView[syscall.GUID] {
	var _result *IVectorView[syscall.GUID]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PurposeGuids, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type CellularApnContext struct {
	RtClass
	*ICellularApnContext
}

func NewCellularApnContext() *CellularApnContext {
	hs := NewHStr("Windows.Networking.Connectivity.CellularApnContext")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &CellularApnContext{
		RtClass:             RtClass{PInspect: p},
		ICellularApnContext: (*ICellularApnContext)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type ConnectionCost struct {
	RtClass
	*IConnectionCost
}

type ConnectionProfile struct {
	RtClass
	*IConnectionProfile
}

type ConnectionProfileFilter struct {
	RtClass
	*IConnectionProfileFilter
}

func NewConnectionProfileFilter() *ConnectionProfileFilter {
	hs := NewHStr("Windows.Networking.Connectivity.ConnectionProfileFilter")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &ConnectionProfileFilter{
		RtClass:                  RtClass{PInspect: p},
		IConnectionProfileFilter: (*IConnectionProfileFilter)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type ConnectionSession struct {
	RtClass
	*IConnectionSession
}

type ConnectivityInterval struct {
	RtClass
	*IConnectivityInterval
}

type ConnectivityManager struct {
	RtClass
}

func NewIConnectivityManagerStatics() *IConnectivityManagerStatics {
	var p *IConnectivityManagerStatics
	hs := NewHStr("Windows.Networking.Connectivity.ConnectivityManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IConnectivityManagerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type DataPlanStatus struct {
	RtClass
	*IDataPlanStatus
}

type DataPlanUsage struct {
	RtClass
	*IDataPlanUsage
}

type DataUsage struct {
	RtClass
	*IDataUsage
}

type IPInformation struct {
	RtClass
	*IIPInformation
}

type LanIdentifier struct {
	RtClass
	*ILanIdentifier
}

type LanIdentifierData struct {
	RtClass
	*ILanIdentifierData
}

type NetworkAdapter struct {
	RtClass
	*INetworkAdapter
}

type NetworkInformation struct {
	RtClass
}

func NewINetworkInformationStatics() *INetworkInformationStatics {
	var p *INetworkInformationStatics
	hs := NewHStr("Windows.Networking.Connectivity.NetworkInformation")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_INetworkInformationStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewINetworkInformationStatics2() *INetworkInformationStatics2 {
	var p *INetworkInformationStatics2
	hs := NewHStr("Windows.Networking.Connectivity.NetworkInformation")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_INetworkInformationStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type NetworkItem struct {
	RtClass
	*INetworkItem
}

type NetworkSecuritySettings struct {
	RtClass
	*INetworkSecuritySettings
}

type NetworkStateChangeEventDetails struct {
	RtClass
	*INetworkStateChangeEventDetails
}

type NetworkUsage struct {
	RtClass
	*INetworkUsage
}

type ProxyConfiguration struct {
	RtClass
	*IProxyConfiguration
}

type RoutePolicy struct {
	RtClass
	*IRoutePolicy
}

func NewRoutePolicy_CreateRoutePolicy(connectionProfile *IConnectionProfile, hostName *IHostName, type_ DomainNameType) *RoutePolicy {
	hs := NewHStr("Windows.Networking.Connectivity.RoutePolicy")
	var pFac *IRoutePolicyFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IRoutePolicyFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IRoutePolicy
	p = pFac.CreateRoutePolicy(connectionProfile, hostName, type_)
	result := &RoutePolicy{
		RtClass:      RtClass{PInspect: &p.IInspectable},
		IRoutePolicy: p,
	}
	com.AddToScope(result)
	return result
}

type WlanConnectionProfileDetails struct {
	RtClass
	*IWlanConnectionProfileDetails
}

type WwanConnectionProfileDetails struct {
	RtClass
	*IWwanConnectionProfileDetails
}
