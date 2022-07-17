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
type WiFiDirectAdvertisementListenStateDiscoverability int32

const (
	WiFiDirectAdvertisementListenStateDiscoverability_None      WiFiDirectAdvertisementListenStateDiscoverability = 0
	WiFiDirectAdvertisementListenStateDiscoverability_Normal    WiFiDirectAdvertisementListenStateDiscoverability = 1
	WiFiDirectAdvertisementListenStateDiscoverability_Intensive WiFiDirectAdvertisementListenStateDiscoverability = 2
)

// enum
type WiFiDirectAdvertisementPublisherStatus int32

const (
	WiFiDirectAdvertisementPublisherStatus_Created WiFiDirectAdvertisementPublisherStatus = 0
	WiFiDirectAdvertisementPublisherStatus_Started WiFiDirectAdvertisementPublisherStatus = 1
	WiFiDirectAdvertisementPublisherStatus_Stopped WiFiDirectAdvertisementPublisherStatus = 2
	WiFiDirectAdvertisementPublisherStatus_Aborted WiFiDirectAdvertisementPublisherStatus = 3
)

// enum
type WiFiDirectConfigurationMethod int32

const (
	WiFiDirectConfigurationMethod_ProvidePin WiFiDirectConfigurationMethod = 0
	WiFiDirectConfigurationMethod_DisplayPin WiFiDirectConfigurationMethod = 1
	WiFiDirectConfigurationMethod_PushButton WiFiDirectConfigurationMethod = 2
)

// enum
type WiFiDirectConnectionStatus int32

const (
	WiFiDirectConnectionStatus_Disconnected WiFiDirectConnectionStatus = 0
	WiFiDirectConnectionStatus_Connected    WiFiDirectConnectionStatus = 1
)

// enum
type WiFiDirectDeviceSelectorType int32

const (
	WiFiDirectDeviceSelectorType_DeviceInterface     WiFiDirectDeviceSelectorType = 0
	WiFiDirectDeviceSelectorType_AssociationEndpoint WiFiDirectDeviceSelectorType = 1
)

// enum
type WiFiDirectError int32

const (
	WiFiDirectError_Success           WiFiDirectError = 0
	WiFiDirectError_RadioNotAvailable WiFiDirectError = 1
	WiFiDirectError_ResourceInUse     WiFiDirectError = 2
)

// enum
type WiFiDirectPairingProcedure int32

const (
	WiFiDirectPairingProcedure_GroupOwnerNegotiation WiFiDirectPairingProcedure = 0
	WiFiDirectPairingProcedure_Invitation            WiFiDirectPairingProcedure = 1
)

// interfaces

// AB511A2D-2A06-49A1-A584-61435C7905A6
var IID_IWiFiDirectAdvertisement = syscall.GUID{0xAB511A2D, 0x2A06, 0x49A1,
	[8]byte{0xA5, 0x84, 0x61, 0x43, 0x5C, 0x79, 0x05, 0xA6}}

type IWiFiDirectAdvertisementInterface interface {
	win32.IInspectableInterface
	Get_InformationElements() *IVector[*IWiFiDirectInformationElement]
	Put_InformationElements(value *IVector[*IWiFiDirectInformationElement])
	Get_ListenStateDiscoverability() WiFiDirectAdvertisementListenStateDiscoverability
	Put_ListenStateDiscoverability(value WiFiDirectAdvertisementListenStateDiscoverability)
	Get_IsAutonomousGroupOwnerEnabled() bool
	Put_IsAutonomousGroupOwnerEnabled(value bool)
	Get_LegacySettings() *IWiFiDirectLegacySettings
}

type IWiFiDirectAdvertisementVtbl struct {
	win32.IInspectableVtbl
	Get_InformationElements           uintptr
	Put_InformationElements           uintptr
	Get_ListenStateDiscoverability    uintptr
	Put_ListenStateDiscoverability    uintptr
	Get_IsAutonomousGroupOwnerEnabled uintptr
	Put_IsAutonomousGroupOwnerEnabled uintptr
	Get_LegacySettings                uintptr
}

type IWiFiDirectAdvertisement struct {
	win32.IInspectable
}

func (this *IWiFiDirectAdvertisement) Vtbl() *IWiFiDirectAdvertisementVtbl {
	return (*IWiFiDirectAdvertisementVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiDirectAdvertisement) Get_InformationElements() *IVector[*IWiFiDirectInformationElement] {
	var _result *IVector[*IWiFiDirectInformationElement]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InformationElements, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiDirectAdvertisement) Put_InformationElements(value *IVector[*IWiFiDirectInformationElement]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InformationElements, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IWiFiDirectAdvertisement) Get_ListenStateDiscoverability() WiFiDirectAdvertisementListenStateDiscoverability {
	var _result WiFiDirectAdvertisementListenStateDiscoverability
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ListenStateDiscoverability, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectAdvertisement) Put_ListenStateDiscoverability(value WiFiDirectAdvertisementListenStateDiscoverability) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ListenStateDiscoverability, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IWiFiDirectAdvertisement) Get_IsAutonomousGroupOwnerEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsAutonomousGroupOwnerEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectAdvertisement) Put_IsAutonomousGroupOwnerEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsAutonomousGroupOwnerEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IWiFiDirectAdvertisement) Get_LegacySettings() *IWiFiDirectLegacySettings {
	var _result *IWiFiDirectLegacySettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LegacySettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B759AA46-D816-491B-917A-B40D7DC403A2
var IID_IWiFiDirectAdvertisement2 = syscall.GUID{0xB759AA46, 0xD816, 0x491B,
	[8]byte{0x91, 0x7A, 0xB4, 0x0D, 0x7D, 0xC4, 0x03, 0xA2}}

type IWiFiDirectAdvertisement2Interface interface {
	win32.IInspectableInterface
	Get_SupportedConfigurationMethods() *IVector[WiFiDirectConfigurationMethod]
}

type IWiFiDirectAdvertisement2Vtbl struct {
	win32.IInspectableVtbl
	Get_SupportedConfigurationMethods uintptr
}

type IWiFiDirectAdvertisement2 struct {
	win32.IInspectable
}

func (this *IWiFiDirectAdvertisement2) Vtbl() *IWiFiDirectAdvertisement2Vtbl {
	return (*IWiFiDirectAdvertisement2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiDirectAdvertisement2) Get_SupportedConfigurationMethods() *IVector[WiFiDirectConfigurationMethod] {
	var _result *IVector[WiFiDirectConfigurationMethod]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedConfigurationMethods, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B35A2D1A-9B1F-45D9-925A-694D66DF68EF
var IID_IWiFiDirectAdvertisementPublisher = syscall.GUID{0xB35A2D1A, 0x9B1F, 0x45D9,
	[8]byte{0x92, 0x5A, 0x69, 0x4D, 0x66, 0xDF, 0x68, 0xEF}}

type IWiFiDirectAdvertisementPublisherInterface interface {
	win32.IInspectableInterface
	Get_Advertisement() *IWiFiDirectAdvertisement
	Get_Status() WiFiDirectAdvertisementPublisherStatus
	Add_StatusChanged(handler TypedEventHandler[*IWiFiDirectAdvertisementPublisher, *IWiFiDirectAdvertisementPublisherStatusChangedEventArgs]) EventRegistrationToken
	Remove_StatusChanged(token EventRegistrationToken)
	Start()
	Stop()
}

type IWiFiDirectAdvertisementPublisherVtbl struct {
	win32.IInspectableVtbl
	Get_Advertisement    uintptr
	Get_Status           uintptr
	Add_StatusChanged    uintptr
	Remove_StatusChanged uintptr
	Start                uintptr
	Stop                 uintptr
}

type IWiFiDirectAdvertisementPublisher struct {
	win32.IInspectable
}

func (this *IWiFiDirectAdvertisementPublisher) Vtbl() *IWiFiDirectAdvertisementPublisherVtbl {
	return (*IWiFiDirectAdvertisementPublisherVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiDirectAdvertisementPublisher) Get_Advertisement() *IWiFiDirectAdvertisement {
	var _result *IWiFiDirectAdvertisement
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Advertisement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiDirectAdvertisementPublisher) Get_Status() WiFiDirectAdvertisementPublisherStatus {
	var _result WiFiDirectAdvertisementPublisherStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectAdvertisementPublisher) Add_StatusChanged(handler TypedEventHandler[*IWiFiDirectAdvertisementPublisher, *IWiFiDirectAdvertisementPublisherStatusChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StatusChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectAdvertisementPublisher) Remove_StatusChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StatusChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IWiFiDirectAdvertisementPublisher) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IWiFiDirectAdvertisementPublisher) Stop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// AAFDE53C-5481-46E6-90DD-32116518F192
var IID_IWiFiDirectAdvertisementPublisherStatusChangedEventArgs = syscall.GUID{0xAAFDE53C, 0x5481, 0x46E6,
	[8]byte{0x90, 0xDD, 0x32, 0x11, 0x65, 0x18, 0xF1, 0x92}}

type IWiFiDirectAdvertisementPublisherStatusChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Status() WiFiDirectAdvertisementPublisherStatus
	Get_Error() WiFiDirectError
}

type IWiFiDirectAdvertisementPublisherStatusChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
	Get_Error  uintptr
}

type IWiFiDirectAdvertisementPublisherStatusChangedEventArgs struct {
	win32.IInspectable
}

func (this *IWiFiDirectAdvertisementPublisherStatusChangedEventArgs) Vtbl() *IWiFiDirectAdvertisementPublisherStatusChangedEventArgsVtbl {
	return (*IWiFiDirectAdvertisementPublisherStatusChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiDirectAdvertisementPublisherStatusChangedEventArgs) Get_Status() WiFiDirectAdvertisementPublisherStatus {
	var _result WiFiDirectAdvertisementPublisherStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectAdvertisementPublisherStatusChangedEventArgs) Get_Error() WiFiDirectError {
	var _result WiFiDirectError
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Error, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 699C1B0D-8D13-4EE9-B9EC-9C72F8251F7D
var IID_IWiFiDirectConnectionListener = syscall.GUID{0x699C1B0D, 0x8D13, 0x4EE9,
	[8]byte{0xB9, 0xEC, 0x9C, 0x72, 0xF8, 0x25, 0x1F, 0x7D}}

type IWiFiDirectConnectionListenerInterface interface {
	win32.IInspectableInterface
	Add_ConnectionRequested(handler TypedEventHandler[*IWiFiDirectConnectionListener, *IWiFiDirectConnectionRequestedEventArgs]) EventRegistrationToken
	Remove_ConnectionRequested(token EventRegistrationToken)
}

type IWiFiDirectConnectionListenerVtbl struct {
	win32.IInspectableVtbl
	Add_ConnectionRequested    uintptr
	Remove_ConnectionRequested uintptr
}

type IWiFiDirectConnectionListener struct {
	win32.IInspectable
}

func (this *IWiFiDirectConnectionListener) Vtbl() *IWiFiDirectConnectionListenerVtbl {
	return (*IWiFiDirectConnectionListenerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiDirectConnectionListener) Add_ConnectionRequested(handler TypedEventHandler[*IWiFiDirectConnectionListener, *IWiFiDirectConnectionRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ConnectionRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectConnectionListener) Remove_ConnectionRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ConnectionRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// B2E55405-5702-4B16-A02C-BBCD21EF6098
var IID_IWiFiDirectConnectionParameters = syscall.GUID{0xB2E55405, 0x5702, 0x4B16,
	[8]byte{0xA0, 0x2C, 0xBB, 0xCD, 0x21, 0xEF, 0x60, 0x98}}

type IWiFiDirectConnectionParametersInterface interface {
	win32.IInspectableInterface
	Get_GroupOwnerIntent() int16
	Put_GroupOwnerIntent(value int16)
}

type IWiFiDirectConnectionParametersVtbl struct {
	win32.IInspectableVtbl
	Get_GroupOwnerIntent uintptr
	Put_GroupOwnerIntent uintptr
}

type IWiFiDirectConnectionParameters struct {
	win32.IInspectable
}

func (this *IWiFiDirectConnectionParameters) Vtbl() *IWiFiDirectConnectionParametersVtbl {
	return (*IWiFiDirectConnectionParametersVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiDirectConnectionParameters) Get_GroupOwnerIntent() int16 {
	var _result int16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GroupOwnerIntent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectConnectionParameters) Put_GroupOwnerIntent(value int16) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_GroupOwnerIntent, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// AB3B0FBE-AA82-44B4-88C8-E3056B89801D
var IID_IWiFiDirectConnectionParameters2 = syscall.GUID{0xAB3B0FBE, 0xAA82, 0x44B4,
	[8]byte{0x88, 0xC8, 0xE3, 0x05, 0x6B, 0x89, 0x80, 0x1D}}

type IWiFiDirectConnectionParameters2Interface interface {
	win32.IInspectableInterface
	Get_PreferenceOrderedConfigurationMethods() *IVector[WiFiDirectConfigurationMethod]
	Get_PreferredPairingProcedure() WiFiDirectPairingProcedure
	Put_PreferredPairingProcedure(value WiFiDirectPairingProcedure)
}

type IWiFiDirectConnectionParameters2Vtbl struct {
	win32.IInspectableVtbl
	Get_PreferenceOrderedConfigurationMethods uintptr
	Get_PreferredPairingProcedure             uintptr
	Put_PreferredPairingProcedure             uintptr
}

type IWiFiDirectConnectionParameters2 struct {
	win32.IInspectable
}

func (this *IWiFiDirectConnectionParameters2) Vtbl() *IWiFiDirectConnectionParameters2Vtbl {
	return (*IWiFiDirectConnectionParameters2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiDirectConnectionParameters2) Get_PreferenceOrderedConfigurationMethods() *IVector[WiFiDirectConfigurationMethod] {
	var _result *IVector[WiFiDirectConfigurationMethod]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PreferenceOrderedConfigurationMethods, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiDirectConnectionParameters2) Get_PreferredPairingProcedure() WiFiDirectPairingProcedure {
	var _result WiFiDirectPairingProcedure
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PreferredPairingProcedure, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectConnectionParameters2) Put_PreferredPairingProcedure(value WiFiDirectPairingProcedure) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PreferredPairingProcedure, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 598AF493-7642-456F-B9D8-E8A9EB1F401A
var IID_IWiFiDirectConnectionParametersStatics = syscall.GUID{0x598AF493, 0x7642, 0x456F,
	[8]byte{0xB9, 0xD8, 0xE8, 0xA9, 0xEB, 0x1F, 0x40, 0x1A}}

type IWiFiDirectConnectionParametersStaticsInterface interface {
	win32.IInspectableInterface
	GetDevicePairingKinds(configurationMethod WiFiDirectConfigurationMethod) DevicePairingKinds
}

type IWiFiDirectConnectionParametersStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDevicePairingKinds uintptr
}

type IWiFiDirectConnectionParametersStatics struct {
	win32.IInspectable
}

func (this *IWiFiDirectConnectionParametersStatics) Vtbl() *IWiFiDirectConnectionParametersStaticsVtbl {
	return (*IWiFiDirectConnectionParametersStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiDirectConnectionParametersStatics) GetDevicePairingKinds(configurationMethod WiFiDirectConfigurationMethod) DevicePairingKinds {
	var _result DevicePairingKinds
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDevicePairingKinds, uintptr(unsafe.Pointer(this)), uintptr(configurationMethod), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 8EB99605-914F-49C3-A614-D18DC5B19B43
var IID_IWiFiDirectConnectionRequest = syscall.GUID{0x8EB99605, 0x914F, 0x49C3,
	[8]byte{0xA6, 0x14, 0xD1, 0x8D, 0xC5, 0xB1, 0x9B, 0x43}}

type IWiFiDirectConnectionRequestInterface interface {
	win32.IInspectableInterface
	Get_DeviceInformation() *IDeviceInformation
}

type IWiFiDirectConnectionRequestVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceInformation uintptr
}

type IWiFiDirectConnectionRequest struct {
	win32.IInspectable
}

func (this *IWiFiDirectConnectionRequest) Vtbl() *IWiFiDirectConnectionRequestVtbl {
	return (*IWiFiDirectConnectionRequestVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiDirectConnectionRequest) Get_DeviceInformation() *IDeviceInformation {
	var _result *IDeviceInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F99D20BE-D38D-484F-8215-E7B65ABF244C
var IID_IWiFiDirectConnectionRequestedEventArgs = syscall.GUID{0xF99D20BE, 0xD38D, 0x484F,
	[8]byte{0x82, 0x15, 0xE7, 0xB6, 0x5A, 0xBF, 0x24, 0x4C}}

type IWiFiDirectConnectionRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	GetConnectionRequest() *IWiFiDirectConnectionRequest
}

type IWiFiDirectConnectionRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	GetConnectionRequest uintptr
}

type IWiFiDirectConnectionRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IWiFiDirectConnectionRequestedEventArgs) Vtbl() *IWiFiDirectConnectionRequestedEventArgsVtbl {
	return (*IWiFiDirectConnectionRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiDirectConnectionRequestedEventArgs) GetConnectionRequest() *IWiFiDirectConnectionRequest {
	var _result *IWiFiDirectConnectionRequest
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetConnectionRequest, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 72DEAAA8-72EB-4DAE-8A28-8513355D2777
var IID_IWiFiDirectDevice = syscall.GUID{0x72DEAAA8, 0x72EB, 0x4DAE,
	[8]byte{0x8A, 0x28, 0x85, 0x13, 0x35, 0x5D, 0x27, 0x77}}

type IWiFiDirectDeviceInterface interface {
	win32.IInspectableInterface
	Get_ConnectionStatus() WiFiDirectConnectionStatus
	Get_DeviceId() string
	Add_ConnectionStatusChanged(handler TypedEventHandler[*IWiFiDirectDevice, interface{}]) EventRegistrationToken
	Remove_ConnectionStatusChanged(token EventRegistrationToken)
	GetConnectionEndpointPairs() *IVectorView[*IEndpointPair]
}

type IWiFiDirectDeviceVtbl struct {
	win32.IInspectableVtbl
	Get_ConnectionStatus           uintptr
	Get_DeviceId                   uintptr
	Add_ConnectionStatusChanged    uintptr
	Remove_ConnectionStatusChanged uintptr
	GetConnectionEndpointPairs     uintptr
}

type IWiFiDirectDevice struct {
	win32.IInspectable
}

func (this *IWiFiDirectDevice) Vtbl() *IWiFiDirectDeviceVtbl {
	return (*IWiFiDirectDeviceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiDirectDevice) Get_ConnectionStatus() WiFiDirectConnectionStatus {
	var _result WiFiDirectConnectionStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConnectionStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectDevice) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWiFiDirectDevice) Add_ConnectionStatusChanged(handler TypedEventHandler[*IWiFiDirectDevice, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ConnectionStatusChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectDevice) Remove_ConnectionStatusChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ConnectionStatusChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IWiFiDirectDevice) GetConnectionEndpointPairs() *IVectorView[*IEndpointPair] {
	var _result *IVectorView[*IEndpointPair]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetConnectionEndpointPairs, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E86CB57C-3AAC-4851-A792-482AAF931B04
var IID_IWiFiDirectDeviceStatics = syscall.GUID{0xE86CB57C, 0x3AAC, 0x4851,
	[8]byte{0xA7, 0x92, 0x48, 0x2A, 0xAF, 0x93, 0x1B, 0x04}}

type IWiFiDirectDeviceStaticsInterface interface {
	win32.IInspectableInterface
	GetDeviceSelector() string
	FromIdAsync(deviceId string) *IAsyncOperation[*IWiFiDirectDevice]
}

type IWiFiDirectDeviceStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelector uintptr
	FromIdAsync       uintptr
}

type IWiFiDirectDeviceStatics struct {
	win32.IInspectable
}

func (this *IWiFiDirectDeviceStatics) Vtbl() *IWiFiDirectDeviceStaticsVtbl {
	return (*IWiFiDirectDeviceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiDirectDeviceStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWiFiDirectDeviceStatics) FromIdAsync(deviceId string) *IAsyncOperation[*IWiFiDirectDevice] {
	var _result *IAsyncOperation[*IWiFiDirectDevice]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1A953E49-B103-437E-9226-AB67971342F9
var IID_IWiFiDirectDeviceStatics2 = syscall.GUID{0x1A953E49, 0xB103, 0x437E,
	[8]byte{0x92, 0x26, 0xAB, 0x67, 0x97, 0x13, 0x42, 0xF9}}

type IWiFiDirectDeviceStatics2Interface interface {
	win32.IInspectableInterface
	GetDeviceSelector(type_ WiFiDirectDeviceSelectorType) string
	FromIdAsync(deviceId string, connectionParameters *IWiFiDirectConnectionParameters) *IAsyncOperation[*IWiFiDirectDevice]
}

type IWiFiDirectDeviceStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelector uintptr
	FromIdAsync       uintptr
}

type IWiFiDirectDeviceStatics2 struct {
	win32.IInspectable
}

func (this *IWiFiDirectDeviceStatics2) Vtbl() *IWiFiDirectDeviceStatics2Vtbl {
	return (*IWiFiDirectDeviceStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiDirectDeviceStatics2) GetDeviceSelector(type_ WiFiDirectDeviceSelectorType) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(type_), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWiFiDirectDeviceStatics2) FromIdAsync(deviceId string, connectionParameters *IWiFiDirectConnectionParameters) *IAsyncOperation[*IWiFiDirectDevice] {
	var _result *IAsyncOperation[*IWiFiDirectDevice]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(connectionParameters)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AFFB72D6-76BB-497E-AC8B-DC72838BC309
var IID_IWiFiDirectInformationElement = syscall.GUID{0xAFFB72D6, 0x76BB, 0x497E,
	[8]byte{0xAC, 0x8B, 0xDC, 0x72, 0x83, 0x8B, 0xC3, 0x09}}

type IWiFiDirectInformationElementInterface interface {
	win32.IInspectableInterface
	Get_Oui() *IBuffer
	Put_Oui(value *IBuffer)
	Get_OuiType() byte
	Put_OuiType(value byte)
	Get_Value() *IBuffer
	Put_Value(value *IBuffer)
}

type IWiFiDirectInformationElementVtbl struct {
	win32.IInspectableVtbl
	Get_Oui     uintptr
	Put_Oui     uintptr
	Get_OuiType uintptr
	Put_OuiType uintptr
	Get_Value   uintptr
	Put_Value   uintptr
}

type IWiFiDirectInformationElement struct {
	win32.IInspectable
}

func (this *IWiFiDirectInformationElement) Vtbl() *IWiFiDirectInformationElementVtbl {
	return (*IWiFiDirectInformationElementVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiDirectInformationElement) Get_Oui() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Oui, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiDirectInformationElement) Put_Oui(value *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Oui, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IWiFiDirectInformationElement) Get_OuiType() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OuiType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectInformationElement) Put_OuiType(value byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_OuiType, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IWiFiDirectInformationElement) Get_Value() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiDirectInformationElement) Put_Value(value *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// DBD02F16-11A5-4E60-8CAA-34772148378A
var IID_IWiFiDirectInformationElementStatics = syscall.GUID{0xDBD02F16, 0x11A5, 0x4E60,
	[8]byte{0x8C, 0xAA, 0x34, 0x77, 0x21, 0x48, 0x37, 0x8A}}

type IWiFiDirectInformationElementStaticsInterface interface {
	win32.IInspectableInterface
	CreateFromBuffer(buffer *IBuffer) *IVector[*IWiFiDirectInformationElement]
	CreateFromDeviceInformation(deviceInformation *IDeviceInformation) *IVector[*IWiFiDirectInformationElement]
}

type IWiFiDirectInformationElementStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateFromBuffer            uintptr
	CreateFromDeviceInformation uintptr
}

type IWiFiDirectInformationElementStatics struct {
	win32.IInspectable
}

func (this *IWiFiDirectInformationElementStatics) Vtbl() *IWiFiDirectInformationElementStaticsVtbl {
	return (*IWiFiDirectInformationElementStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiDirectInformationElementStatics) CreateFromBuffer(buffer *IBuffer) *IVector[*IWiFiDirectInformationElement] {
	var _result *IVector[*IWiFiDirectInformationElement]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromBuffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiDirectInformationElementStatics) CreateFromDeviceInformation(deviceInformation *IDeviceInformation) *IVector[*IWiFiDirectInformationElement] {
	var _result *IVector[*IWiFiDirectInformationElement]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromDeviceInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(deviceInformation)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A64FDBBA-F2FD-4567-A91B-F5C2F5321057
var IID_IWiFiDirectLegacySettings = syscall.GUID{0xA64FDBBA, 0xF2FD, 0x4567,
	[8]byte{0xA9, 0x1B, 0xF5, 0xC2, 0xF5, 0x32, 0x10, 0x57}}

type IWiFiDirectLegacySettingsInterface interface {
	win32.IInspectableInterface
	Get_IsEnabled() bool
	Put_IsEnabled(value bool)
	Get_Ssid() string
	Put_Ssid(value string)
	Get_Passphrase() *IPasswordCredential
	Put_Passphrase(value *IPasswordCredential)
}

type IWiFiDirectLegacySettingsVtbl struct {
	win32.IInspectableVtbl
	Get_IsEnabled  uintptr
	Put_IsEnabled  uintptr
	Get_Ssid       uintptr
	Put_Ssid       uintptr
	Get_Passphrase uintptr
	Put_Passphrase uintptr
}

type IWiFiDirectLegacySettings struct {
	win32.IInspectable
}

func (this *IWiFiDirectLegacySettings) Vtbl() *IWiFiDirectLegacySettingsVtbl {
	return (*IWiFiDirectLegacySettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiDirectLegacySettings) Get_IsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectLegacySettings) Put_IsEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IWiFiDirectLegacySettings) Get_Ssid() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ssid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWiFiDirectLegacySettings) Put_Ssid(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Ssid, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IWiFiDirectLegacySettings) Get_Passphrase() *IPasswordCredential {
	var _result *IPasswordCredential
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Passphrase, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiDirectLegacySettings) Put_Passphrase(value *IPasswordCredential) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Passphrase, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// classes

type WiFiDirectAdvertisement struct {
	RtClass
	*IWiFiDirectAdvertisement
}

type WiFiDirectAdvertisementPublisher struct {
	RtClass
	*IWiFiDirectAdvertisementPublisher
}

func NewWiFiDirectAdvertisementPublisher() *WiFiDirectAdvertisementPublisher {
	hs := NewHStr("Windows.Devices.WiFiDirect.WiFiDirectAdvertisementPublisher")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &WiFiDirectAdvertisementPublisher{
		RtClass:                           RtClass{PInspect: p},
		IWiFiDirectAdvertisementPublisher: (*IWiFiDirectAdvertisementPublisher)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type WiFiDirectAdvertisementPublisherStatusChangedEventArgs struct {
	RtClass
	*IWiFiDirectAdvertisementPublisherStatusChangedEventArgs
}

type WiFiDirectConnectionListener struct {
	RtClass
	*IWiFiDirectConnectionListener
}

func NewWiFiDirectConnectionListener() *WiFiDirectConnectionListener {
	hs := NewHStr("Windows.Devices.WiFiDirect.WiFiDirectConnectionListener")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &WiFiDirectConnectionListener{
		RtClass:                       RtClass{PInspect: p},
		IWiFiDirectConnectionListener: (*IWiFiDirectConnectionListener)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type WiFiDirectConnectionParameters struct {
	RtClass
	*IWiFiDirectConnectionParameters
}

func NewWiFiDirectConnectionParameters() *WiFiDirectConnectionParameters {
	hs := NewHStr("Windows.Devices.WiFiDirect.WiFiDirectConnectionParameters")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &WiFiDirectConnectionParameters{
		RtClass:                         RtClass{PInspect: p},
		IWiFiDirectConnectionParameters: (*IWiFiDirectConnectionParameters)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewIWiFiDirectConnectionParametersStatics() *IWiFiDirectConnectionParametersStatics {
	var p *IWiFiDirectConnectionParametersStatics
	hs := NewHStr("Windows.Devices.WiFiDirect.WiFiDirectConnectionParameters")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWiFiDirectConnectionParametersStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type WiFiDirectConnectionRequest struct {
	RtClass
	*IWiFiDirectConnectionRequest
}

type WiFiDirectConnectionRequestedEventArgs struct {
	RtClass
	*IWiFiDirectConnectionRequestedEventArgs
}

type WiFiDirectDevice struct {
	RtClass
	*IWiFiDirectDevice
}

func NewIWiFiDirectDeviceStatics() *IWiFiDirectDeviceStatics {
	var p *IWiFiDirectDeviceStatics
	hs := NewHStr("Windows.Devices.WiFiDirect.WiFiDirectDevice")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWiFiDirectDeviceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIWiFiDirectDeviceStatics2() *IWiFiDirectDeviceStatics2 {
	var p *IWiFiDirectDeviceStatics2
	hs := NewHStr("Windows.Devices.WiFiDirect.WiFiDirectDevice")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWiFiDirectDeviceStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type WiFiDirectInformationElement struct {
	RtClass
	*IWiFiDirectInformationElement
}

func NewWiFiDirectInformationElement() *WiFiDirectInformationElement {
	hs := NewHStr("Windows.Devices.WiFiDirect.WiFiDirectInformationElement")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &WiFiDirectInformationElement{
		RtClass:                       RtClass{PInspect: p},
		IWiFiDirectInformationElement: (*IWiFiDirectInformationElement)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewIWiFiDirectInformationElementStatics() *IWiFiDirectInformationElementStatics {
	var p *IWiFiDirectInformationElementStatics
	hs := NewHStr("Windows.Devices.WiFiDirect.WiFiDirectInformationElement")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWiFiDirectInformationElementStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type WiFiDirectLegacySettings struct {
	RtClass
	*IWiFiDirectLegacySettings
}
