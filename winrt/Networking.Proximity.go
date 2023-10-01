package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
// flags
type PeerDiscoveryTypes uint32

const (
	PeerDiscoveryTypes_None      PeerDiscoveryTypes = 0
	PeerDiscoveryTypes_Browse    PeerDiscoveryTypes = 1
	PeerDiscoveryTypes_Triggered PeerDiscoveryTypes = 2
)

// enum
type PeerRole int32

const (
	PeerRole_Peer   PeerRole = 0
	PeerRole_Host   PeerRole = 1
	PeerRole_Client PeerRole = 2
)

// enum
type PeerWatcherStatus int32

const (
	PeerWatcherStatus_Created              PeerWatcherStatus = 0
	PeerWatcherStatus_Started              PeerWatcherStatus = 1
	PeerWatcherStatus_EnumerationCompleted PeerWatcherStatus = 2
	PeerWatcherStatus_Stopping             PeerWatcherStatus = 3
	PeerWatcherStatus_Stopped              PeerWatcherStatus = 4
	PeerWatcherStatus_Aborted              PeerWatcherStatus = 5
)

// enum
type TriggeredConnectState int32

const (
	TriggeredConnectState_PeerFound  TriggeredConnectState = 0
	TriggeredConnectState_Listening  TriggeredConnectState = 1
	TriggeredConnectState_Connecting TriggeredConnectState = 2
	TriggeredConnectState_Completed  TriggeredConnectState = 3
	TriggeredConnectState_Canceled   TriggeredConnectState = 4
	TriggeredConnectState_Failed     TriggeredConnectState = 5
)

// func types

// EFA9DA69-F6E1-49C9-A49E-8E0FC58FB911
type DeviceArrivedEventHandler func(sender *IProximityDevice) com.Error

// EFA9DA69-F6E2-49C9-A49E-8E0FC58FB911
type DeviceDepartedEventHandler func(sender *IProximityDevice) com.Error

// EFAB0782-F6E2-4675-A045-D8E320C24808
type MessageReceivedHandler func(sender *IProximityDevice, message *IProximityMessage) com.Error

// EFAA0B4A-F6E2-4D7D-856C-78FC8EFC021E
type MessageTransmittedHandler func(sender *IProximityDevice, messageId int64) com.Error

// interfaces

// EB6891AE-4F1E-4C66-BD0D-46924A942E08
var IID_IConnectionRequestedEventArgs = syscall.GUID{0xEB6891AE, 0x4F1E, 0x4C66,
	[8]byte{0xBD, 0x0D, 0x46, 0x92, 0x4A, 0x94, 0x2E, 0x08}}

type IConnectionRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_PeerInformation() *IPeerInformation
}

type IConnectionRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_PeerInformation uintptr
}

type IConnectionRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IConnectionRequestedEventArgs) Vtbl() *IConnectionRequestedEventArgsVtbl {
	return (*IConnectionRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IConnectionRequestedEventArgs) Get_PeerInformation() *IPeerInformation {
	var _result *IPeerInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PeerInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 914B3B61-F6E1-47C4-A14C-148A1903D0C6
var IID_IPeerFinderStatics = syscall.GUID{0x914B3B61, 0xF6E1, 0x47C4,
	[8]byte{0xA1, 0x4C, 0x14, 0x8A, 0x19, 0x03, 0xD0, 0xC6}}

type IPeerFinderStaticsInterface interface {
	win32.IInspectableInterface
	Get_AllowBluetooth() bool
	Put_AllowBluetooth(value bool)
	Get_AllowInfrastructure() bool
	Put_AllowInfrastructure(value bool)
	Get_AllowWiFiDirect() bool
	Put_AllowWiFiDirect(value bool)
	Get_DisplayName() string
	Put_DisplayName(value string)
	Get_SupportedDiscoveryTypes() PeerDiscoveryTypes
	Get_AlternateIdentities() *IMap[string, string]
	Start()
	StartWithMessage(peerMessage string)
	Stop()
	Add_TriggeredConnectionStateChanged(handler TypedEventHandler[interface{}, *ITriggeredConnectionStateChangedEventArgs]) EventRegistrationToken
	Remove_TriggeredConnectionStateChanged(cookie EventRegistrationToken)
	Add_ConnectionRequested(handler TypedEventHandler[interface{}, *IConnectionRequestedEventArgs]) EventRegistrationToken
	Remove_ConnectionRequested(cookie EventRegistrationToken)
	FindAllPeersAsync() *IAsyncOperation[*IVectorView[*IPeerInformation]]
	ConnectAsync(peerInformation *IPeerInformation) *IAsyncOperation[*IStreamSocket]
}

type IPeerFinderStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_AllowBluetooth                     uintptr
	Put_AllowBluetooth                     uintptr
	Get_AllowInfrastructure                uintptr
	Put_AllowInfrastructure                uintptr
	Get_AllowWiFiDirect                    uintptr
	Put_AllowWiFiDirect                    uintptr
	Get_DisplayName                        uintptr
	Put_DisplayName                        uintptr
	Get_SupportedDiscoveryTypes            uintptr
	Get_AlternateIdentities                uintptr
	Start                                  uintptr
	StartWithMessage                       uintptr
	Stop                                   uintptr
	Add_TriggeredConnectionStateChanged    uintptr
	Remove_TriggeredConnectionStateChanged uintptr
	Add_ConnectionRequested                uintptr
	Remove_ConnectionRequested             uintptr
	FindAllPeersAsync                      uintptr
	ConnectAsync                           uintptr
}

type IPeerFinderStatics struct {
	win32.IInspectable
}

func (this *IPeerFinderStatics) Vtbl() *IPeerFinderStaticsVtbl {
	return (*IPeerFinderStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPeerFinderStatics) Get_AllowBluetooth() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllowBluetooth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPeerFinderStatics) Put_AllowBluetooth(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AllowBluetooth, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPeerFinderStatics) Get_AllowInfrastructure() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllowInfrastructure, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPeerFinderStatics) Put_AllowInfrastructure(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AllowInfrastructure, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPeerFinderStatics) Get_AllowWiFiDirect() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllowWiFiDirect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPeerFinderStatics) Put_AllowWiFiDirect(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AllowWiFiDirect, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPeerFinderStatics) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPeerFinderStatics) Put_DisplayName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DisplayName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPeerFinderStatics) Get_SupportedDiscoveryTypes() PeerDiscoveryTypes {
	var _result PeerDiscoveryTypes
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedDiscoveryTypes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPeerFinderStatics) Get_AlternateIdentities() *IMap[string, string] {
	var _result *IMap[string, string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlternateIdentities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPeerFinderStatics) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IPeerFinderStatics) StartWithMessage(peerMessage string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartWithMessage, uintptr(unsafe.Pointer(this)), NewHStr(peerMessage).Ptr)
	_ = _hr
}

func (this *IPeerFinderStatics) Stop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IPeerFinderStatics) Add_TriggeredConnectionStateChanged(handler TypedEventHandler[interface{}, *ITriggeredConnectionStateChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_TriggeredConnectionStateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPeerFinderStatics) Remove_TriggeredConnectionStateChanged(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_TriggeredConnectionStateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IPeerFinderStatics) Add_ConnectionRequested(handler TypedEventHandler[interface{}, *IConnectionRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ConnectionRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPeerFinderStatics) Remove_ConnectionRequested(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ConnectionRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IPeerFinderStatics) FindAllPeersAsync() *IAsyncOperation[*IVectorView[*IPeerInformation]] {
	var _result *IAsyncOperation[*IVectorView[*IPeerInformation]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllPeersAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPeerFinderStatics) ConnectAsync(peerInformation *IPeerInformation) *IAsyncOperation[*IStreamSocket] {
	var _result *IAsyncOperation[*IStreamSocket]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConnectAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(peerInformation)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D6E73C65-FDD0-4B0B-9312-866408935D82
var IID_IPeerFinderStatics2 = syscall.GUID{0xD6E73C65, 0xFDD0, 0x4B0B,
	[8]byte{0x93, 0x12, 0x86, 0x64, 0x08, 0x93, 0x5D, 0x82}}

type IPeerFinderStatics2Interface interface {
	win32.IInspectableInterface
	Get_Role() PeerRole
	Put_Role(value PeerRole)
	Get_DiscoveryData() *IBuffer
	Put_DiscoveryData(value *IBuffer)
	CreateWatcher() *IPeerWatcher
}

type IPeerFinderStatics2Vtbl struct {
	win32.IInspectableVtbl
	Get_Role          uintptr
	Put_Role          uintptr
	Get_DiscoveryData uintptr
	Put_DiscoveryData uintptr
	CreateWatcher     uintptr
}

type IPeerFinderStatics2 struct {
	win32.IInspectable
}

func (this *IPeerFinderStatics2) Vtbl() *IPeerFinderStatics2Vtbl {
	return (*IPeerFinderStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPeerFinderStatics2) Get_Role() PeerRole {
	var _result PeerRole
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Role, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPeerFinderStatics2) Put_Role(value PeerRole) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Role, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPeerFinderStatics2) Get_DiscoveryData() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DiscoveryData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPeerFinderStatics2) Put_DiscoveryData(value *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DiscoveryData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPeerFinderStatics2) CreateWatcher() *IPeerWatcher {
	var _result *IPeerWatcher
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWatcher, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 20024F08-9FFF-45F4-B6E9-408B2EBEF373
var IID_IPeerInformation = syscall.GUID{0x20024F08, 0x9FFF, 0x45F4,
	[8]byte{0xB6, 0xE9, 0x40, 0x8B, 0x2E, 0xBE, 0xF3, 0x73}}

type IPeerInformationInterface interface {
	win32.IInspectableInterface
	Get_DisplayName() string
}

type IPeerInformationVtbl struct {
	win32.IInspectableVtbl
	Get_DisplayName uintptr
}

type IPeerInformation struct {
	win32.IInspectable
}

func (this *IPeerInformation) Vtbl() *IPeerInformationVtbl {
	return (*IPeerInformationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPeerInformation) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// B20F612A-DBD0-40F8-95BD-2D4209C7836F
var IID_IPeerInformation3 = syscall.GUID{0xB20F612A, 0xDBD0, 0x40F8,
	[8]byte{0x95, 0xBD, 0x2D, 0x42, 0x09, 0xC7, 0x83, 0x6F}}

type IPeerInformation3Interface interface {
	win32.IInspectableInterface
	Get_Id() string
	Get_DiscoveryData() *IBuffer
}

type IPeerInformation3Vtbl struct {
	win32.IInspectableVtbl
	Get_Id            uintptr
	Get_DiscoveryData uintptr
}

type IPeerInformation3 struct {
	win32.IInspectable
}

func (this *IPeerInformation3) Vtbl() *IPeerInformation3Vtbl {
	return (*IPeerInformation3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPeerInformation3) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPeerInformation3) Get_DiscoveryData() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DiscoveryData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// ECC7CCAD-1B70-4E8B-92DB-BBE781419308
var IID_IPeerInformationWithHostAndService = syscall.GUID{0xECC7CCAD, 0x1B70, 0x4E8B,
	[8]byte{0x92, 0xDB, 0xBB, 0xE7, 0x81, 0x41, 0x93, 0x08}}

type IPeerInformationWithHostAndServiceInterface interface {
	win32.IInspectableInterface
	Get_HostName() *IHostName
	Get_ServiceName() string
}

type IPeerInformationWithHostAndServiceVtbl struct {
	win32.IInspectableVtbl
	Get_HostName    uintptr
	Get_ServiceName uintptr
}

type IPeerInformationWithHostAndService struct {
	win32.IInspectable
}

func (this *IPeerInformationWithHostAndService) Vtbl() *IPeerInformationWithHostAndServiceVtbl {
	return (*IPeerInformationWithHostAndServiceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPeerInformationWithHostAndService) Get_HostName() *IHostName {
	var _result *IHostName
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HostName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPeerInformationWithHostAndService) Get_ServiceName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 3CEE21F8-2FA6-4679-9691-03C94A420F34
var IID_IPeerWatcher = syscall.GUID{0x3CEE21F8, 0x2FA6, 0x4679,
	[8]byte{0x96, 0x91, 0x03, 0xC9, 0x4A, 0x42, 0x0F, 0x34}}

type IPeerWatcherInterface interface {
	win32.IInspectableInterface
	Add_Added(handler TypedEventHandler[*IPeerWatcher, *IPeerInformation]) EventRegistrationToken
	Remove_Added(token EventRegistrationToken)
	Add_Removed(handler TypedEventHandler[*IPeerWatcher, *IPeerInformation]) EventRegistrationToken
	Remove_Removed(token EventRegistrationToken)
	Add_Updated(handler TypedEventHandler[*IPeerWatcher, *IPeerInformation]) EventRegistrationToken
	Remove_Updated(token EventRegistrationToken)
	Add_EnumerationCompleted(handler TypedEventHandler[*IPeerWatcher, interface{}]) EventRegistrationToken
	Remove_EnumerationCompleted(token EventRegistrationToken)
	Add_Stopped(handler TypedEventHandler[*IPeerWatcher, interface{}]) EventRegistrationToken
	Remove_Stopped(token EventRegistrationToken)
	Get_Status() PeerWatcherStatus
	Start()
	Stop()
}

type IPeerWatcherVtbl struct {
	win32.IInspectableVtbl
	Add_Added                   uintptr
	Remove_Added                uintptr
	Add_Removed                 uintptr
	Remove_Removed              uintptr
	Add_Updated                 uintptr
	Remove_Updated              uintptr
	Add_EnumerationCompleted    uintptr
	Remove_EnumerationCompleted uintptr
	Add_Stopped                 uintptr
	Remove_Stopped              uintptr
	Get_Status                  uintptr
	Start                       uintptr
	Stop                        uintptr
}

type IPeerWatcher struct {
	win32.IInspectable
}

func (this *IPeerWatcher) Vtbl() *IPeerWatcherVtbl {
	return (*IPeerWatcherVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPeerWatcher) Add_Added(handler TypedEventHandler[*IPeerWatcher, *IPeerInformation]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Added, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPeerWatcher) Remove_Added(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Added, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPeerWatcher) Add_Removed(handler TypedEventHandler[*IPeerWatcher, *IPeerInformation]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Removed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPeerWatcher) Remove_Removed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Removed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPeerWatcher) Add_Updated(handler TypedEventHandler[*IPeerWatcher, *IPeerInformation]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Updated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPeerWatcher) Remove_Updated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Updated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPeerWatcher) Add_EnumerationCompleted(handler TypedEventHandler[*IPeerWatcher, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_EnumerationCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPeerWatcher) Remove_EnumerationCompleted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_EnumerationCompleted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPeerWatcher) Add_Stopped(handler TypedEventHandler[*IPeerWatcher, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Stopped, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPeerWatcher) Remove_Stopped(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Stopped, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPeerWatcher) Get_Status() PeerWatcherStatus {
	var _result PeerWatcherStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPeerWatcher) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IPeerWatcher) Stop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// EFA8A552-F6E1-4329-A0FC-AB6B0FD28262
var IID_IProximityDevice = syscall.GUID{0xEFA8A552, 0xF6E1, 0x4329,
	[8]byte{0xA0, 0xFC, 0xAB, 0x6B, 0x0F, 0xD2, 0x82, 0x62}}

type IProximityDeviceInterface interface {
	win32.IInspectableInterface
	SubscribeForMessage(messageType string, messageReceivedHandler MessageReceivedHandler) int64
	PublishMessage(messageType string, message string) int64
	PublishMessageWithCallback(messageType string, message string, messageTransmittedHandler MessageTransmittedHandler) int64
	PublishBinaryMessage(messageType string, message *IBuffer) int64
	PublishBinaryMessageWithCallback(messageType string, message *IBuffer, messageTransmittedHandler MessageTransmittedHandler) int64
	PublishUriMessage(message *IUriRuntimeClass) int64
	PublishUriMessageWithCallback(message *IUriRuntimeClass, messageTransmittedHandler MessageTransmittedHandler) int64
	StopSubscribingForMessage(subscriptionId int64)
	StopPublishingMessage(messageId int64)
	Add_DeviceArrived(arrivedHandler DeviceArrivedEventHandler) EventRegistrationToken
	Remove_DeviceArrived(cookie EventRegistrationToken)
	Add_DeviceDeparted(departedHandler DeviceDepartedEventHandler) EventRegistrationToken
	Remove_DeviceDeparted(cookie EventRegistrationToken)
	Get_MaxMessageBytes() uint32
	Get_BitsPerSecond() uint64
	Get_DeviceId() string
}

type IProximityDeviceVtbl struct {
	win32.IInspectableVtbl
	SubscribeForMessage              uintptr
	PublishMessage                   uintptr
	PublishMessageWithCallback       uintptr
	PublishBinaryMessage             uintptr
	PublishBinaryMessageWithCallback uintptr
	PublishUriMessage                uintptr
	PublishUriMessageWithCallback    uintptr
	StopSubscribingForMessage        uintptr
	StopPublishingMessage            uintptr
	Add_DeviceArrived                uintptr
	Remove_DeviceArrived             uintptr
	Add_DeviceDeparted               uintptr
	Remove_DeviceDeparted            uintptr
	Get_MaxMessageBytes              uintptr
	Get_BitsPerSecond                uintptr
	Get_DeviceId                     uintptr
}

type IProximityDevice struct {
	win32.IInspectable
}

func (this *IProximityDevice) Vtbl() *IProximityDeviceVtbl {
	return (*IProximityDeviceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProximityDevice) SubscribeForMessage(messageType string, messageReceivedHandler MessageReceivedHandler) int64 {
	var _result int64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SubscribeForMessage, uintptr(unsafe.Pointer(this)), NewHStr(messageType).Ptr, uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(messageReceivedHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProximityDevice) PublishMessage(messageType string, message string) int64 {
	var _result int64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PublishMessage, uintptr(unsafe.Pointer(this)), NewHStr(messageType).Ptr, NewHStr(message).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProximityDevice) PublishMessageWithCallback(messageType string, message string, messageTransmittedHandler MessageTransmittedHandler) int64 {
	var _result int64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PublishMessageWithCallback, uintptr(unsafe.Pointer(this)), NewHStr(messageType).Ptr, NewHStr(message).Ptr, uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(messageTransmittedHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProximityDevice) PublishBinaryMessage(messageType string, message *IBuffer) int64 {
	var _result int64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PublishBinaryMessage, uintptr(unsafe.Pointer(this)), NewHStr(messageType).Ptr, uintptr(unsafe.Pointer(message)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProximityDevice) PublishBinaryMessageWithCallback(messageType string, message *IBuffer, messageTransmittedHandler MessageTransmittedHandler) int64 {
	var _result int64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PublishBinaryMessageWithCallback, uintptr(unsafe.Pointer(this)), NewHStr(messageType).Ptr, uintptr(unsafe.Pointer(message)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(messageTransmittedHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProximityDevice) PublishUriMessage(message *IUriRuntimeClass) int64 {
	var _result int64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PublishUriMessage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(message)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProximityDevice) PublishUriMessageWithCallback(message *IUriRuntimeClass, messageTransmittedHandler MessageTransmittedHandler) int64 {
	var _result int64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PublishUriMessageWithCallback, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(message)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(messageTransmittedHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProximityDevice) StopSubscribingForMessage(subscriptionId int64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopSubscribingForMessage, uintptr(unsafe.Pointer(this)), uintptr(subscriptionId))
	_ = _hr
}

func (this *IProximityDevice) StopPublishingMessage(messageId int64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopPublishingMessage, uintptr(unsafe.Pointer(this)), uintptr(messageId))
	_ = _hr
}

func (this *IProximityDevice) Add_DeviceArrived(arrivedHandler DeviceArrivedEventHandler) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DeviceArrived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewOneArgFuncDelegate(arrivedHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProximityDevice) Remove_DeviceArrived(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DeviceArrived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IProximityDevice) Add_DeviceDeparted(departedHandler DeviceDepartedEventHandler) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DeviceDeparted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewOneArgFuncDelegate(departedHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProximityDevice) Remove_DeviceDeparted(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DeviceDeparted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IProximityDevice) Get_MaxMessageBytes() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxMessageBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProximityDevice) Get_BitsPerSecond() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BitsPerSecond, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProximityDevice) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 914BA01D-F6E1-47C4-A14C-148A1903D0C6
var IID_IProximityDeviceStatics = syscall.GUID{0x914BA01D, 0xF6E1, 0x47C4,
	[8]byte{0xA1, 0x4C, 0x14, 0x8A, 0x19, 0x03, 0xD0, 0xC6}}

type IProximityDeviceStaticsInterface interface {
	win32.IInspectableInterface
	GetDeviceSelector() string
	GetDefault() *IProximityDevice
	FromId(deviceId string) *IProximityDevice
}

type IProximityDeviceStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDeviceSelector uintptr
	GetDefault        uintptr
	FromId            uintptr
}

type IProximityDeviceStatics struct {
	win32.IInspectable
}

func (this *IProximityDeviceStatics) Vtbl() *IProximityDeviceStaticsVtbl {
	return (*IProximityDeviceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProximityDeviceStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IProximityDeviceStatics) GetDefault() *IProximityDevice {
	var _result *IProximityDevice
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProximityDeviceStatics) FromId(deviceId string) *IProximityDevice {
	var _result *IProximityDevice
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromId, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EFAB0782-F6E1-4675-A045-D8E320C24808
var IID_IProximityMessage = syscall.GUID{0xEFAB0782, 0xF6E1, 0x4675,
	[8]byte{0xA0, 0x45, 0xD8, 0xE3, 0x20, 0xC2, 0x48, 0x08}}

type IProximityMessageInterface interface {
	win32.IInspectableInterface
	Get_MessageType() string
	Get_SubscriptionId() int64
	Get_Data() *IBuffer
	Get_DataAsString() string
}

type IProximityMessageVtbl struct {
	win32.IInspectableVtbl
	Get_MessageType    uintptr
	Get_SubscriptionId uintptr
	Get_Data           uintptr
	Get_DataAsString   uintptr
}

type IProximityMessage struct {
	win32.IInspectable
}

func (this *IProximityMessage) Vtbl() *IProximityMessageVtbl {
	return (*IProximityMessageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProximityMessage) Get_MessageType() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MessageType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IProximityMessage) Get_SubscriptionId() int64 {
	var _result int64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SubscriptionId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProximityMessage) Get_Data() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProximityMessage) Get_DataAsString() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DataAsString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// C6A780AD-F6E1-4D54-96E2-33F620BCA88A
var IID_ITriggeredConnectionStateChangedEventArgs = syscall.GUID{0xC6A780AD, 0xF6E1, 0x4D54,
	[8]byte{0x96, 0xE2, 0x33, 0xF6, 0x20, 0xBC, 0xA8, 0x8A}}

type ITriggeredConnectionStateChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_State() TriggeredConnectState
	Get_Id() uint32
	Get_Socket() *IStreamSocket
}

type ITriggeredConnectionStateChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_State  uintptr
	Get_Id     uintptr
	Get_Socket uintptr
}

type ITriggeredConnectionStateChangedEventArgs struct {
	win32.IInspectable
}

func (this *ITriggeredConnectionStateChangedEventArgs) Vtbl() *ITriggeredConnectionStateChangedEventArgsVtbl {
	return (*ITriggeredConnectionStateChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITriggeredConnectionStateChangedEventArgs) Get_State() TriggeredConnectState {
	var _result TriggeredConnectState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITriggeredConnectionStateChangedEventArgs) Get_Id() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITriggeredConnectionStateChangedEventArgs) Get_Socket() *IStreamSocket {
	var _result *IStreamSocket
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Socket, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type ProximityDevice struct {
	RtClass
	*IProximityDevice
}

func NewIProximityDeviceStatics() *IProximityDeviceStatics {
	var p *IProximityDeviceStatics
	hs := NewHStr("Windows.Networking.Proximity.ProximityDevice")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IProximityDeviceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type ProximityMessage struct {
	RtClass
	*IProximityMessage
}
