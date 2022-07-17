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
type WiFiDirectServiceAdvertisementStatus int32

const (
	WiFiDirectServiceAdvertisementStatus_Created WiFiDirectServiceAdvertisementStatus = 0
	WiFiDirectServiceAdvertisementStatus_Started WiFiDirectServiceAdvertisementStatus = 1
	WiFiDirectServiceAdvertisementStatus_Stopped WiFiDirectServiceAdvertisementStatus = 2
	WiFiDirectServiceAdvertisementStatus_Aborted WiFiDirectServiceAdvertisementStatus = 3
)

// enum
type WiFiDirectServiceConfigurationMethod int32

const (
	WiFiDirectServiceConfigurationMethod_Default    WiFiDirectServiceConfigurationMethod = 0
	WiFiDirectServiceConfigurationMethod_PinDisplay WiFiDirectServiceConfigurationMethod = 1
	WiFiDirectServiceConfigurationMethod_PinEntry   WiFiDirectServiceConfigurationMethod = 2
)

// enum
type WiFiDirectServiceError int32

const (
	WiFiDirectServiceError_Success             WiFiDirectServiceError = 0
	WiFiDirectServiceError_RadioNotAvailable   WiFiDirectServiceError = 1
	WiFiDirectServiceError_ResourceInUse       WiFiDirectServiceError = 2
	WiFiDirectServiceError_UnsupportedHardware WiFiDirectServiceError = 3
	WiFiDirectServiceError_NoHardware          WiFiDirectServiceError = 4
)

// enum
type WiFiDirectServiceIPProtocol int32

const (
	WiFiDirectServiceIPProtocol_Tcp WiFiDirectServiceIPProtocol = 6
	WiFiDirectServiceIPProtocol_Udp WiFiDirectServiceIPProtocol = 17
)

// enum
type WiFiDirectServiceSessionErrorStatus int32

const (
	WiFiDirectServiceSessionErrorStatus_Ok                   WiFiDirectServiceSessionErrorStatus = 0
	WiFiDirectServiceSessionErrorStatus_Disassociated        WiFiDirectServiceSessionErrorStatus = 1
	WiFiDirectServiceSessionErrorStatus_LocalClose           WiFiDirectServiceSessionErrorStatus = 2
	WiFiDirectServiceSessionErrorStatus_RemoteClose          WiFiDirectServiceSessionErrorStatus = 3
	WiFiDirectServiceSessionErrorStatus_SystemFailure        WiFiDirectServiceSessionErrorStatus = 4
	WiFiDirectServiceSessionErrorStatus_NoResponseFromRemote WiFiDirectServiceSessionErrorStatus = 5
)

// enum
type WiFiDirectServiceSessionStatus int32

const (
	WiFiDirectServiceSessionStatus_Closed    WiFiDirectServiceSessionStatus = 0
	WiFiDirectServiceSessionStatus_Initiated WiFiDirectServiceSessionStatus = 1
	WiFiDirectServiceSessionStatus_Requested WiFiDirectServiceSessionStatus = 2
	WiFiDirectServiceSessionStatus_Open      WiFiDirectServiceSessionStatus = 3
)

// enum
type WiFiDirectServiceStatus int32

const (
	WiFiDirectServiceStatus_Available WiFiDirectServiceStatus = 0
	WiFiDirectServiceStatus_Busy      WiFiDirectServiceStatus = 1
	WiFiDirectServiceStatus_Custom    WiFiDirectServiceStatus = 2
)

// interfaces

// 50AABBB8-5F71-45EC-84F1-A1E4FC7879A3
var IID_IWiFiDirectService = syscall.GUID{0x50AABBB8, 0x5F71, 0x45EC,
	[8]byte{0x84, 0xF1, 0xA1, 0xE4, 0xFC, 0x78, 0x79, 0xA3}}

type IWiFiDirectServiceInterface interface {
	win32.IInspectableInterface
	Get_RemoteServiceInfo() *IBuffer
	Get_SupportedConfigurationMethods() *IVectorView[WiFiDirectServiceConfigurationMethod]
	Get_PreferGroupOwnerMode() bool
	Put_PreferGroupOwnerMode(value bool)
	Get_SessionInfo() *IBuffer
	Put_SessionInfo(value *IBuffer)
	Get_ServiceError() WiFiDirectServiceError
	Add_SessionDeferred(handler TypedEventHandler[*IWiFiDirectService, *IWiFiDirectServiceSessionDeferredEventArgs]) EventRegistrationToken
	Remove_SessionDeferred(token EventRegistrationToken)
	GetProvisioningInfoAsync(selectedConfigurationMethod WiFiDirectServiceConfigurationMethod) *IAsyncOperation[*IWiFiDirectServiceProvisioningInfo]
	ConnectAsync() *IAsyncOperation[*IWiFiDirectServiceSession]
	ConnectAsyncWithPin(pin string) *IAsyncOperation[*IWiFiDirectServiceSession]
}

type IWiFiDirectServiceVtbl struct {
	win32.IInspectableVtbl
	Get_RemoteServiceInfo             uintptr
	Get_SupportedConfigurationMethods uintptr
	Get_PreferGroupOwnerMode          uintptr
	Put_PreferGroupOwnerMode          uintptr
	Get_SessionInfo                   uintptr
	Put_SessionInfo                   uintptr
	Get_ServiceError                  uintptr
	Add_SessionDeferred               uintptr
	Remove_SessionDeferred            uintptr
	GetProvisioningInfoAsync          uintptr
	ConnectAsync                      uintptr
	ConnectAsyncWithPin               uintptr
}

type IWiFiDirectService struct {
	win32.IInspectable
}

func (this *IWiFiDirectService) Vtbl() *IWiFiDirectServiceVtbl {
	return (*IWiFiDirectServiceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiDirectService) Get_RemoteServiceInfo() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RemoteServiceInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiDirectService) Get_SupportedConfigurationMethods() *IVectorView[WiFiDirectServiceConfigurationMethod] {
	var _result *IVectorView[WiFiDirectServiceConfigurationMethod]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedConfigurationMethods, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiDirectService) Get_PreferGroupOwnerMode() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PreferGroupOwnerMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectService) Put_PreferGroupOwnerMode(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PreferGroupOwnerMode, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IWiFiDirectService) Get_SessionInfo() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SessionInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiDirectService) Put_SessionInfo(value *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SessionInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IWiFiDirectService) Get_ServiceError() WiFiDirectServiceError {
	var _result WiFiDirectServiceError
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectService) Add_SessionDeferred(handler TypedEventHandler[*IWiFiDirectService, *IWiFiDirectServiceSessionDeferredEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SessionDeferred, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectService) Remove_SessionDeferred(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SessionDeferred, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IWiFiDirectService) GetProvisioningInfoAsync(selectedConfigurationMethod WiFiDirectServiceConfigurationMethod) *IAsyncOperation[*IWiFiDirectServiceProvisioningInfo] {
	var _result *IAsyncOperation[*IWiFiDirectServiceProvisioningInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetProvisioningInfoAsync, uintptr(unsafe.Pointer(this)), uintptr(selectedConfigurationMethod), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiDirectService) ConnectAsync() *IAsyncOperation[*IWiFiDirectServiceSession] {
	var _result *IAsyncOperation[*IWiFiDirectServiceSession]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConnectAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiDirectService) ConnectAsyncWithPin(pin string) *IAsyncOperation[*IWiFiDirectServiceSession] {
	var _result *IAsyncOperation[*IWiFiDirectServiceSession]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConnectAsyncWithPin, uintptr(unsafe.Pointer(this)), NewHStr(pin).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A4AA1EE1-9D8F-4F4F-93EE-7DDEA2E37F46
var IID_IWiFiDirectServiceAdvertiser = syscall.GUID{0xA4AA1EE1, 0x9D8F, 0x4F4F,
	[8]byte{0x93, 0xEE, 0x7D, 0xDE, 0xA2, 0xE3, 0x7F, 0x46}}

type IWiFiDirectServiceAdvertiserInterface interface {
	win32.IInspectableInterface
	Get_ServiceName() string
	Get_ServiceNamePrefixes() *IVector[string]
	Get_ServiceInfo() *IBuffer
	Put_ServiceInfo(value *IBuffer)
	Get_AutoAcceptSession() bool
	Put_AutoAcceptSession(value bool)
	Get_PreferGroupOwnerMode() bool
	Put_PreferGroupOwnerMode(value bool)
	Get_PreferredConfigurationMethods() *IVector[WiFiDirectServiceConfigurationMethod]
	Get_ServiceStatus() WiFiDirectServiceStatus
	Put_ServiceStatus(value WiFiDirectServiceStatus)
	Get_CustomServiceStatusCode() uint32
	Put_CustomServiceStatusCode(value uint32)
	Get_DeferredSessionInfo() *IBuffer
	Put_DeferredSessionInfo(value *IBuffer)
	Get_AdvertisementStatus() WiFiDirectServiceAdvertisementStatus
	Get_ServiceError() WiFiDirectServiceError
	Add_SessionRequested(handler TypedEventHandler[*IWiFiDirectServiceAdvertiser, *IWiFiDirectServiceSessionRequestedEventArgs]) EventRegistrationToken
	Remove_SessionRequested(token EventRegistrationToken)
	Add_AutoAcceptSessionConnected(handler TypedEventHandler[*IWiFiDirectServiceAdvertiser, *IWiFiDirectServiceAutoAcceptSessionConnectedEventArgs]) EventRegistrationToken
	Remove_AutoAcceptSessionConnected(token EventRegistrationToken)
	Add_AdvertisementStatusChanged(handler TypedEventHandler[*IWiFiDirectServiceAdvertiser, interface{}]) EventRegistrationToken
	Remove_AdvertisementStatusChanged(token EventRegistrationToken)
	ConnectAsync(deviceInfo *IDeviceInformation) *IAsyncOperation[*IWiFiDirectServiceSession]
	ConnectAsyncWithPin(deviceInfo *IDeviceInformation, pin string) *IAsyncOperation[*IWiFiDirectServiceSession]
	Start()
	Stop()
}

type IWiFiDirectServiceAdvertiserVtbl struct {
	win32.IInspectableVtbl
	Get_ServiceName                   uintptr
	Get_ServiceNamePrefixes           uintptr
	Get_ServiceInfo                   uintptr
	Put_ServiceInfo                   uintptr
	Get_AutoAcceptSession             uintptr
	Put_AutoAcceptSession             uintptr
	Get_PreferGroupOwnerMode          uintptr
	Put_PreferGroupOwnerMode          uintptr
	Get_PreferredConfigurationMethods uintptr
	Get_ServiceStatus                 uintptr
	Put_ServiceStatus                 uintptr
	Get_CustomServiceStatusCode       uintptr
	Put_CustomServiceStatusCode       uintptr
	Get_DeferredSessionInfo           uintptr
	Put_DeferredSessionInfo           uintptr
	Get_AdvertisementStatus           uintptr
	Get_ServiceError                  uintptr
	Add_SessionRequested              uintptr
	Remove_SessionRequested           uintptr
	Add_AutoAcceptSessionConnected    uintptr
	Remove_AutoAcceptSessionConnected uintptr
	Add_AdvertisementStatusChanged    uintptr
	Remove_AdvertisementStatusChanged uintptr
	ConnectAsync                      uintptr
	ConnectAsyncWithPin               uintptr
	Start                             uintptr
	Stop                              uintptr
}

type IWiFiDirectServiceAdvertiser struct {
	win32.IInspectable
}

func (this *IWiFiDirectServiceAdvertiser) Vtbl() *IWiFiDirectServiceAdvertiserVtbl {
	return (*IWiFiDirectServiceAdvertiserVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiDirectServiceAdvertiser) Get_ServiceName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWiFiDirectServiceAdvertiser) Get_ServiceNamePrefixes() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceNamePrefixes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiDirectServiceAdvertiser) Get_ServiceInfo() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiDirectServiceAdvertiser) Put_ServiceInfo(value *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ServiceInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IWiFiDirectServiceAdvertiser) Get_AutoAcceptSession() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AutoAcceptSession, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectServiceAdvertiser) Put_AutoAcceptSession(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AutoAcceptSession, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IWiFiDirectServiceAdvertiser) Get_PreferGroupOwnerMode() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PreferGroupOwnerMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectServiceAdvertiser) Put_PreferGroupOwnerMode(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PreferGroupOwnerMode, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IWiFiDirectServiceAdvertiser) Get_PreferredConfigurationMethods() *IVector[WiFiDirectServiceConfigurationMethod] {
	var _result *IVector[WiFiDirectServiceConfigurationMethod]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PreferredConfigurationMethods, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiDirectServiceAdvertiser) Get_ServiceStatus() WiFiDirectServiceStatus {
	var _result WiFiDirectServiceStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectServiceAdvertiser) Put_ServiceStatus(value WiFiDirectServiceStatus) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ServiceStatus, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IWiFiDirectServiceAdvertiser) Get_CustomServiceStatusCode() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CustomServiceStatusCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectServiceAdvertiser) Put_CustomServiceStatusCode(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CustomServiceStatusCode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IWiFiDirectServiceAdvertiser) Get_DeferredSessionInfo() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeferredSessionInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiDirectServiceAdvertiser) Put_DeferredSessionInfo(value *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DeferredSessionInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IWiFiDirectServiceAdvertiser) Get_AdvertisementStatus() WiFiDirectServiceAdvertisementStatus {
	var _result WiFiDirectServiceAdvertisementStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AdvertisementStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectServiceAdvertiser) Get_ServiceError() WiFiDirectServiceError {
	var _result WiFiDirectServiceError
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectServiceAdvertiser) Add_SessionRequested(handler TypedEventHandler[*IWiFiDirectServiceAdvertiser, *IWiFiDirectServiceSessionRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SessionRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectServiceAdvertiser) Remove_SessionRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SessionRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IWiFiDirectServiceAdvertiser) Add_AutoAcceptSessionConnected(handler TypedEventHandler[*IWiFiDirectServiceAdvertiser, *IWiFiDirectServiceAutoAcceptSessionConnectedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AutoAcceptSessionConnected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectServiceAdvertiser) Remove_AutoAcceptSessionConnected(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AutoAcceptSessionConnected, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IWiFiDirectServiceAdvertiser) Add_AdvertisementStatusChanged(handler TypedEventHandler[*IWiFiDirectServiceAdvertiser, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AdvertisementStatusChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectServiceAdvertiser) Remove_AdvertisementStatusChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AdvertisementStatusChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IWiFiDirectServiceAdvertiser) ConnectAsync(deviceInfo *IDeviceInformation) *IAsyncOperation[*IWiFiDirectServiceSession] {
	var _result *IAsyncOperation[*IWiFiDirectServiceSession]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConnectAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(deviceInfo)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiDirectServiceAdvertiser) ConnectAsyncWithPin(deviceInfo *IDeviceInformation, pin string) *IAsyncOperation[*IWiFiDirectServiceSession] {
	var _result *IAsyncOperation[*IWiFiDirectServiceSession]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConnectAsyncWithPin, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(deviceInfo)), NewHStr(pin).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiDirectServiceAdvertiser) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IWiFiDirectServiceAdvertiser) Stop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 3106AC0D-B446-4F13-9F9A-8AE925FEBA2B
var IID_IWiFiDirectServiceAdvertiserFactory = syscall.GUID{0x3106AC0D, 0xB446, 0x4F13,
	[8]byte{0x9F, 0x9A, 0x8A, 0xE9, 0x25, 0xFE, 0xBA, 0x2B}}

type IWiFiDirectServiceAdvertiserFactoryInterface interface {
	win32.IInspectableInterface
	CreateWiFiDirectServiceAdvertiser(serviceName string) *IWiFiDirectServiceAdvertiser
}

type IWiFiDirectServiceAdvertiserFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateWiFiDirectServiceAdvertiser uintptr
}

type IWiFiDirectServiceAdvertiserFactory struct {
	win32.IInspectable
}

func (this *IWiFiDirectServiceAdvertiserFactory) Vtbl() *IWiFiDirectServiceAdvertiserFactoryVtbl {
	return (*IWiFiDirectServiceAdvertiserFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiDirectServiceAdvertiserFactory) CreateWiFiDirectServiceAdvertiser(serviceName string) *IWiFiDirectServiceAdvertiser {
	var _result *IWiFiDirectServiceAdvertiser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWiFiDirectServiceAdvertiser, uintptr(unsafe.Pointer(this)), NewHStr(serviceName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DCD9E01E-83DF-43E5-8F43-CBE8479E84EB
var IID_IWiFiDirectServiceAutoAcceptSessionConnectedEventArgs = syscall.GUID{0xDCD9E01E, 0x83DF, 0x43E5,
	[8]byte{0x8F, 0x43, 0xCB, 0xE8, 0x47, 0x9E, 0x84, 0xEB}}

type IWiFiDirectServiceAutoAcceptSessionConnectedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Session() *IWiFiDirectServiceSession
	Get_SessionInfo() *IBuffer
}

type IWiFiDirectServiceAutoAcceptSessionConnectedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Session     uintptr
	Get_SessionInfo uintptr
}

type IWiFiDirectServiceAutoAcceptSessionConnectedEventArgs struct {
	win32.IInspectable
}

func (this *IWiFiDirectServiceAutoAcceptSessionConnectedEventArgs) Vtbl() *IWiFiDirectServiceAutoAcceptSessionConnectedEventArgsVtbl {
	return (*IWiFiDirectServiceAutoAcceptSessionConnectedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiDirectServiceAutoAcceptSessionConnectedEventArgs) Get_Session() *IWiFiDirectServiceSession {
	var _result *IWiFiDirectServiceSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Session, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiDirectServiceAutoAcceptSessionConnectedEventArgs) Get_SessionInfo() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SessionInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8BDB7CFE-97D9-45A2-8E99-DB50910FB6A6
var IID_IWiFiDirectServiceProvisioningInfo = syscall.GUID{0x8BDB7CFE, 0x97D9, 0x45A2,
	[8]byte{0x8E, 0x99, 0xDB, 0x50, 0x91, 0x0F, 0xB6, 0xA6}}

type IWiFiDirectServiceProvisioningInfoInterface interface {
	win32.IInspectableInterface
	Get_SelectedConfigurationMethod() WiFiDirectServiceConfigurationMethod
	Get_IsGroupFormationNeeded() bool
}

type IWiFiDirectServiceProvisioningInfoVtbl struct {
	win32.IInspectableVtbl
	Get_SelectedConfigurationMethod uintptr
	Get_IsGroupFormationNeeded      uintptr
}

type IWiFiDirectServiceProvisioningInfo struct {
	win32.IInspectable
}

func (this *IWiFiDirectServiceProvisioningInfo) Vtbl() *IWiFiDirectServiceProvisioningInfoVtbl {
	return (*IWiFiDirectServiceProvisioningInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiDirectServiceProvisioningInfo) Get_SelectedConfigurationMethod() WiFiDirectServiceConfigurationMethod {
	var _result WiFiDirectServiceConfigurationMethod
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SelectedConfigurationMethod, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectServiceProvisioningInfo) Get_IsGroupFormationNeeded() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsGroupFormationNeeded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D4CEBAC1-3FD3-4F0E-B7BD-782906F44411
var IID_IWiFiDirectServiceRemotePortAddedEventArgs = syscall.GUID{0xD4CEBAC1, 0x3FD3, 0x4F0E,
	[8]byte{0xB7, 0xBD, 0x78, 0x29, 0x06, 0xF4, 0x44, 0x11}}

type IWiFiDirectServiceRemotePortAddedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_EndpointPairs() *IVectorView[*IEndpointPair]
	Get_Protocol() WiFiDirectServiceIPProtocol
}

type IWiFiDirectServiceRemotePortAddedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_EndpointPairs uintptr
	Get_Protocol      uintptr
}

type IWiFiDirectServiceRemotePortAddedEventArgs struct {
	win32.IInspectable
}

func (this *IWiFiDirectServiceRemotePortAddedEventArgs) Vtbl() *IWiFiDirectServiceRemotePortAddedEventArgsVtbl {
	return (*IWiFiDirectServiceRemotePortAddedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiDirectServiceRemotePortAddedEventArgs) Get_EndpointPairs() *IVectorView[*IEndpointPair] {
	var _result *IVectorView[*IEndpointPair]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EndpointPairs, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiDirectServiceRemotePortAddedEventArgs) Get_Protocol() WiFiDirectServiceIPProtocol {
	var _result WiFiDirectServiceIPProtocol
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Protocol, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 81142163-E426-47CB-8640-E1B3588BF26F
var IID_IWiFiDirectServiceSession = syscall.GUID{0x81142163, 0xE426, 0x47CB,
	[8]byte{0x86, 0x40, 0xE1, 0xB3, 0x58, 0x8B, 0xF2, 0x6F}}

type IWiFiDirectServiceSessionInterface interface {
	win32.IInspectableInterface
	Get_ServiceName() string
	Get_Status() WiFiDirectServiceSessionStatus
	Get_ErrorStatus() WiFiDirectServiceSessionErrorStatus
	Get_SessionId() uint32
	Get_AdvertisementId() uint32
	Get_ServiceAddress() string
	Get_SessionAddress() string
	GetConnectionEndpointPairs() *IVectorView[*IEndpointPair]
	Add_SessionStatusChanged(handler TypedEventHandler[*IWiFiDirectServiceSession, interface{}]) EventRegistrationToken
	Remove_SessionStatusChanged(token EventRegistrationToken)
	AddStreamSocketListenerAsync(value *IStreamSocketListener) *IAsyncAction
	AddDatagramSocketAsync(value *IDatagramSocket) *IAsyncAction
	Add_RemotePortAdded(handler TypedEventHandler[*IWiFiDirectServiceSession, *IWiFiDirectServiceRemotePortAddedEventArgs]) EventRegistrationToken
	Remove_RemotePortAdded(token EventRegistrationToken)
}

type IWiFiDirectServiceSessionVtbl struct {
	win32.IInspectableVtbl
	Get_ServiceName              uintptr
	Get_Status                   uintptr
	Get_ErrorStatus              uintptr
	Get_SessionId                uintptr
	Get_AdvertisementId          uintptr
	Get_ServiceAddress           uintptr
	Get_SessionAddress           uintptr
	GetConnectionEndpointPairs   uintptr
	Add_SessionStatusChanged     uintptr
	Remove_SessionStatusChanged  uintptr
	AddStreamSocketListenerAsync uintptr
	AddDatagramSocketAsync       uintptr
	Add_RemotePortAdded          uintptr
	Remove_RemotePortAdded       uintptr
}

type IWiFiDirectServiceSession struct {
	win32.IInspectable
}

func (this *IWiFiDirectServiceSession) Vtbl() *IWiFiDirectServiceSessionVtbl {
	return (*IWiFiDirectServiceSessionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiDirectServiceSession) Get_ServiceName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWiFiDirectServiceSession) Get_Status() WiFiDirectServiceSessionStatus {
	var _result WiFiDirectServiceSessionStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectServiceSession) Get_ErrorStatus() WiFiDirectServiceSessionErrorStatus {
	var _result WiFiDirectServiceSessionErrorStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ErrorStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectServiceSession) Get_SessionId() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SessionId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectServiceSession) Get_AdvertisementId() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AdvertisementId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectServiceSession) Get_ServiceAddress() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWiFiDirectServiceSession) Get_SessionAddress() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SessionAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWiFiDirectServiceSession) GetConnectionEndpointPairs() *IVectorView[*IEndpointPair] {
	var _result *IVectorView[*IEndpointPair]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetConnectionEndpointPairs, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiDirectServiceSession) Add_SessionStatusChanged(handler TypedEventHandler[*IWiFiDirectServiceSession, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SessionStatusChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectServiceSession) Remove_SessionStatusChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SessionStatusChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IWiFiDirectServiceSession) AddStreamSocketListenerAsync(value *IStreamSocketListener) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddStreamSocketListenerAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiDirectServiceSession) AddDatagramSocketAsync(value *IDatagramSocket) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddDatagramSocketAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiDirectServiceSession) Add_RemotePortAdded(handler TypedEventHandler[*IWiFiDirectServiceSession, *IWiFiDirectServiceRemotePortAddedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_RemotePortAdded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWiFiDirectServiceSession) Remove_RemotePortAdded(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_RemotePortAdded, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 8DFC197F-1201-4F1F-B6F4-5DF1B7B9FB2E
var IID_IWiFiDirectServiceSessionDeferredEventArgs = syscall.GUID{0x8DFC197F, 0x1201, 0x4F1F,
	[8]byte{0xB6, 0xF4, 0x5D, 0xF1, 0xB7, 0xB9, 0xFB, 0x2E}}

type IWiFiDirectServiceSessionDeferredEventArgsInterface interface {
	win32.IInspectableInterface
	Get_DeferredSessionInfo() *IBuffer
}

type IWiFiDirectServiceSessionDeferredEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_DeferredSessionInfo uintptr
}

type IWiFiDirectServiceSessionDeferredEventArgs struct {
	win32.IInspectable
}

func (this *IWiFiDirectServiceSessionDeferredEventArgs) Vtbl() *IWiFiDirectServiceSessionDeferredEventArgsVtbl {
	return (*IWiFiDirectServiceSessionDeferredEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiDirectServiceSessionDeferredEventArgs) Get_DeferredSessionInfo() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeferredSessionInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A0E27C8B-50CB-4A58-9BCF-E472B99FBA04
var IID_IWiFiDirectServiceSessionRequest = syscall.GUID{0xA0E27C8B, 0x50CB, 0x4A58,
	[8]byte{0x9B, 0xCF, 0xE4, 0x72, 0xB9, 0x9F, 0xBA, 0x04}}

type IWiFiDirectServiceSessionRequestInterface interface {
	win32.IInspectableInterface
	Get_DeviceInformation() *IDeviceInformation
	Get_ProvisioningInfo() *IWiFiDirectServiceProvisioningInfo
	Get_SessionInfo() *IBuffer
}

type IWiFiDirectServiceSessionRequestVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceInformation uintptr
	Get_ProvisioningInfo  uintptr
	Get_SessionInfo       uintptr
}

type IWiFiDirectServiceSessionRequest struct {
	win32.IInspectable
}

func (this *IWiFiDirectServiceSessionRequest) Vtbl() *IWiFiDirectServiceSessionRequestVtbl {
	return (*IWiFiDirectServiceSessionRequestVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiDirectServiceSessionRequest) Get_DeviceInformation() *IDeviceInformation {
	var _result *IDeviceInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiDirectServiceSessionRequest) Get_ProvisioningInfo() *IWiFiDirectServiceProvisioningInfo {
	var _result *IWiFiDirectServiceProvisioningInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProvisioningInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWiFiDirectServiceSessionRequest) Get_SessionInfo() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SessionInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 74BDCC11-53D6-4999-B4F8-6C8ECC1771E7
var IID_IWiFiDirectServiceSessionRequestedEventArgs = syscall.GUID{0x74BDCC11, 0x53D6, 0x4999,
	[8]byte{0xB4, 0xF8, 0x6C, 0x8E, 0xCC, 0x17, 0x71, 0xE7}}

type IWiFiDirectServiceSessionRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	GetSessionRequest() *IWiFiDirectServiceSessionRequest
}

type IWiFiDirectServiceSessionRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	GetSessionRequest uintptr
}

type IWiFiDirectServiceSessionRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IWiFiDirectServiceSessionRequestedEventArgs) Vtbl() *IWiFiDirectServiceSessionRequestedEventArgsVtbl {
	return (*IWiFiDirectServiceSessionRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiDirectServiceSessionRequestedEventArgs) GetSessionRequest() *IWiFiDirectServiceSessionRequest {
	var _result *IWiFiDirectServiceSessionRequest
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSessionRequest, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7DB40045-FD74-4688-B725-5DCE86ACF233
var IID_IWiFiDirectServiceStatics = syscall.GUID{0x7DB40045, 0xFD74, 0x4688,
	[8]byte{0xB7, 0x25, 0x5D, 0xCE, 0x86, 0xAC, 0xF2, 0x33}}

type IWiFiDirectServiceStaticsInterface interface {
	win32.IInspectableInterface
	GetSelector(serviceName string) string
	GetSelectorWithFilter(serviceName string, serviceInfoFilter *IBuffer) string
	FromIdAsync(deviceId string) *IAsyncOperation[*IWiFiDirectService]
}

type IWiFiDirectServiceStaticsVtbl struct {
	win32.IInspectableVtbl
	GetSelector           uintptr
	GetSelectorWithFilter uintptr
	FromIdAsync           uintptr
}

type IWiFiDirectServiceStatics struct {
	win32.IInspectable
}

func (this *IWiFiDirectServiceStatics) Vtbl() *IWiFiDirectServiceStaticsVtbl {
	return (*IWiFiDirectServiceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWiFiDirectServiceStatics) GetSelector(serviceName string) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSelector, uintptr(unsafe.Pointer(this)), NewHStr(serviceName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWiFiDirectServiceStatics) GetSelectorWithFilter(serviceName string, serviceInfoFilter *IBuffer) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSelectorWithFilter, uintptr(unsafe.Pointer(this)), NewHStr(serviceName).Ptr, uintptr(unsafe.Pointer(serviceInfoFilter)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWiFiDirectServiceStatics) FromIdAsync(deviceId string) *IAsyncOperation[*IWiFiDirectService] {
	var _result *IAsyncOperation[*IWiFiDirectService]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type WiFiDirectService struct {
	RtClass
	*IWiFiDirectService
}

func NewIWiFiDirectServiceStatics() *IWiFiDirectServiceStatics {
	var p *IWiFiDirectServiceStatics
	hs := NewHStr("Windows.Devices.WiFiDirect.Services.WiFiDirectService")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWiFiDirectServiceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type WiFiDirectServiceAdvertiser struct {
	RtClass
	*IWiFiDirectServiceAdvertiser
}

func NewWiFiDirectServiceAdvertiser_CreateWiFiDirectServiceAdvertiser(serviceName string) *WiFiDirectServiceAdvertiser {
	hs := NewHStr("Windows.Devices.WiFiDirect.Services.WiFiDirectServiceAdvertiser")
	var pFac *IWiFiDirectServiceAdvertiserFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWiFiDirectServiceAdvertiserFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IWiFiDirectServiceAdvertiser
	p = pFac.CreateWiFiDirectServiceAdvertiser(serviceName)
	result := &WiFiDirectServiceAdvertiser{
		RtClass:                      RtClass{PInspect: &p.IInspectable},
		IWiFiDirectServiceAdvertiser: p,
	}
	com.AddToScope(result)
	return result
}

type WiFiDirectServiceAutoAcceptSessionConnectedEventArgs struct {
	RtClass
	*IWiFiDirectServiceAutoAcceptSessionConnectedEventArgs
}

type WiFiDirectServiceProvisioningInfo struct {
	RtClass
	*IWiFiDirectServiceProvisioningInfo
}

type WiFiDirectServiceRemotePortAddedEventArgs struct {
	RtClass
	*IWiFiDirectServiceRemotePortAddedEventArgs
}

type WiFiDirectServiceSession struct {
	RtClass
	*IWiFiDirectServiceSession
}

type WiFiDirectServiceSessionDeferredEventArgs struct {
	RtClass
	*IWiFiDirectServiceSessionDeferredEventArgs
}

type WiFiDirectServiceSessionRequest struct {
	RtClass
	*IWiFiDirectServiceSessionRequest
}

type WiFiDirectServiceSessionRequestedEventArgs struct {
	RtClass
	*IWiFiDirectServiceSessionRequestedEventArgs
}
