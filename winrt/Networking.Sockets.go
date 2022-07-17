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
type ControlChannelTriggerResetReason int32

const (
	ControlChannelTriggerResetReason_FastUserSwitched   ControlChannelTriggerResetReason = 0
	ControlChannelTriggerResetReason_LowPowerExit       ControlChannelTriggerResetReason = 1
	ControlChannelTriggerResetReason_QuietHoursExit     ControlChannelTriggerResetReason = 2
	ControlChannelTriggerResetReason_ApplicationRestart ControlChannelTriggerResetReason = 3
)

// enum
type ControlChannelTriggerResourceType int32

const (
	ControlChannelTriggerResourceType_RequestSoftwareSlot ControlChannelTriggerResourceType = 0
	ControlChannelTriggerResourceType_RequestHardwareSlot ControlChannelTriggerResourceType = 1
)

// enum
type ControlChannelTriggerStatus int32

const (
	ControlChannelTriggerStatus_HardwareSlotRequested ControlChannelTriggerStatus = 0
	ControlChannelTriggerStatus_SoftwareSlotAllocated ControlChannelTriggerStatus = 1
	ControlChannelTriggerStatus_HardwareSlotAllocated ControlChannelTriggerStatus = 2
	ControlChannelTriggerStatus_PolicyError           ControlChannelTriggerStatus = 3
	ControlChannelTriggerStatus_SystemError           ControlChannelTriggerStatus = 4
	ControlChannelTriggerStatus_TransportDisconnected ControlChannelTriggerStatus = 5
	ControlChannelTriggerStatus_ServiceUnavailable    ControlChannelTriggerStatus = 6
)

// enum
type MessageWebSocketReceiveMode int32

const (
	MessageWebSocketReceiveMode_FullMessage    MessageWebSocketReceiveMode = 0
	MessageWebSocketReceiveMode_PartialMessage MessageWebSocketReceiveMode = 1
)

// enum
type SocketActivityConnectedStandbyAction int32

const (
	SocketActivityConnectedStandbyAction_DoNotWake SocketActivityConnectedStandbyAction = 0
	SocketActivityConnectedStandbyAction_Wake      SocketActivityConnectedStandbyAction = 1
)

// enum
type SocketActivityKind int32

const (
	SocketActivityKind_None                 SocketActivityKind = 0
	SocketActivityKind_StreamSocketListener SocketActivityKind = 1
	SocketActivityKind_DatagramSocket       SocketActivityKind = 2
	SocketActivityKind_StreamSocket         SocketActivityKind = 3
)

// enum
type SocketActivityTriggerReason int32

const (
	SocketActivityTriggerReason_None                  SocketActivityTriggerReason = 0
	SocketActivityTriggerReason_SocketActivity        SocketActivityTriggerReason = 1
	SocketActivityTriggerReason_ConnectionAccepted    SocketActivityTriggerReason = 2
	SocketActivityTriggerReason_KeepAliveTimerExpired SocketActivityTriggerReason = 3
	SocketActivityTriggerReason_SocketClosed          SocketActivityTriggerReason = 4
)

// enum
type SocketErrorStatus int32

const (
	SocketErrorStatus_Unknown                            SocketErrorStatus = 0
	SocketErrorStatus_OperationAborted                   SocketErrorStatus = 1
	SocketErrorStatus_HttpInvalidServerResponse          SocketErrorStatus = 2
	SocketErrorStatus_ConnectionTimedOut                 SocketErrorStatus = 3
	SocketErrorStatus_AddressFamilyNotSupported          SocketErrorStatus = 4
	SocketErrorStatus_SocketTypeNotSupported             SocketErrorStatus = 5
	SocketErrorStatus_HostNotFound                       SocketErrorStatus = 6
	SocketErrorStatus_NoDataRecordOfRequestedType        SocketErrorStatus = 7
	SocketErrorStatus_NonAuthoritativeHostNotFound       SocketErrorStatus = 8
	SocketErrorStatus_ClassTypeNotFound                  SocketErrorStatus = 9
	SocketErrorStatus_AddressAlreadyInUse                SocketErrorStatus = 10
	SocketErrorStatus_CannotAssignRequestedAddress       SocketErrorStatus = 11
	SocketErrorStatus_ConnectionRefused                  SocketErrorStatus = 12
	SocketErrorStatus_NetworkIsUnreachable               SocketErrorStatus = 13
	SocketErrorStatus_UnreachableHost                    SocketErrorStatus = 14
	SocketErrorStatus_NetworkIsDown                      SocketErrorStatus = 15
	SocketErrorStatus_NetworkDroppedConnectionOnReset    SocketErrorStatus = 16
	SocketErrorStatus_SoftwareCausedConnectionAbort      SocketErrorStatus = 17
	SocketErrorStatus_ConnectionResetByPeer              SocketErrorStatus = 18
	SocketErrorStatus_HostIsDown                         SocketErrorStatus = 19
	SocketErrorStatus_NoAddressesFound                   SocketErrorStatus = 20
	SocketErrorStatus_TooManyOpenFiles                   SocketErrorStatus = 21
	SocketErrorStatus_MessageTooLong                     SocketErrorStatus = 22
	SocketErrorStatus_CertificateExpired                 SocketErrorStatus = 23
	SocketErrorStatus_CertificateUntrustedRoot           SocketErrorStatus = 24
	SocketErrorStatus_CertificateCommonNameIsIncorrect   SocketErrorStatus = 25
	SocketErrorStatus_CertificateWrongUsage              SocketErrorStatus = 26
	SocketErrorStatus_CertificateRevoked                 SocketErrorStatus = 27
	SocketErrorStatus_CertificateNoRevocationCheck       SocketErrorStatus = 28
	SocketErrorStatus_CertificateRevocationServerOffline SocketErrorStatus = 29
	SocketErrorStatus_CertificateIsInvalid               SocketErrorStatus = 30
)

// enum
type SocketMessageType int32

const (
	SocketMessageType_Binary SocketMessageType = 0
	SocketMessageType_Utf8   SocketMessageType = 1
)

// enum
type SocketProtectionLevel int32

const (
	SocketProtectionLevel_PlainSocket                                SocketProtectionLevel = 0
	SocketProtectionLevel_Ssl                                        SocketProtectionLevel = 1
	SocketProtectionLevel_SslAllowNullEncryption                     SocketProtectionLevel = 2
	SocketProtectionLevel_BluetoothEncryptionAllowNullAuthentication SocketProtectionLevel = 3
	SocketProtectionLevel_BluetoothEncryptionWithAuthentication      SocketProtectionLevel = 4
	SocketProtectionLevel_Ssl3AllowWeakEncryption                    SocketProtectionLevel = 5
	SocketProtectionLevel_Tls10                                      SocketProtectionLevel = 6
	SocketProtectionLevel_Tls11                                      SocketProtectionLevel = 7
	SocketProtectionLevel_Tls12                                      SocketProtectionLevel = 8
	SocketProtectionLevel_Unspecified                                SocketProtectionLevel = 9
)

// enum
type SocketQualityOfService int32

const (
	SocketQualityOfService_Normal     SocketQualityOfService = 0
	SocketQualityOfService_LowLatency SocketQualityOfService = 1
)

// enum
type SocketSslErrorSeverity int32

const (
	SocketSslErrorSeverity_None      SocketSslErrorSeverity = 0
	SocketSslErrorSeverity_Ignorable SocketSslErrorSeverity = 1
	SocketSslErrorSeverity_Fatal     SocketSslErrorSeverity = 2
)

// structs

type BandwidthStatistics struct {
	OutboundBitsPerSecond            uint64
	InboundBitsPerSecond             uint64
	OutboundBitsPerSecondInstability uint64
	InboundBitsPerSecondInstability  uint64
	OutboundBandwidthPeaked          bool
	InboundBandwidthPeaked           bool
}

type ControlChannelTriggerContract struct {
}

type RoundTripTimeStatistics struct {
	Variance uint32
	Max      uint32
	Min      uint32
	Sum      uint32
}

// interfaces

// 7D1431A7-EE96-40E8-A199-8703CD969EC3
var IID_IControlChannelTrigger = syscall.GUID{0x7D1431A7, 0xEE96, 0x40E8,
	[8]byte{0xA1, 0x99, 0x87, 0x03, 0xCD, 0x96, 0x9E, 0xC3}}

type IControlChannelTriggerInterface interface {
	win32.IInspectableInterface
	Get_ControlChannelTriggerId() string
	Get_ServerKeepAliveIntervalInMinutes() uint32
	Put_ServerKeepAliveIntervalInMinutes(value uint32)
	Get_CurrentKeepAliveIntervalInMinutes() uint32
	Get_TransportObject() interface{}
	Get_KeepAliveTrigger() *IBackgroundTrigger
	Get_PushNotificationTrigger() *IBackgroundTrigger
	UsingTransport(transport interface{})
	WaitForPushEnabled() ControlChannelTriggerStatus
	DecreaseNetworkKeepAliveInterval()
	FlushTransport()
}

type IControlChannelTriggerVtbl struct {
	win32.IInspectableVtbl
	Get_ControlChannelTriggerId           uintptr
	Get_ServerKeepAliveIntervalInMinutes  uintptr
	Put_ServerKeepAliveIntervalInMinutes  uintptr
	Get_CurrentKeepAliveIntervalInMinutes uintptr
	Get_TransportObject                   uintptr
	Get_KeepAliveTrigger                  uintptr
	Get_PushNotificationTrigger           uintptr
	UsingTransport                        uintptr
	WaitForPushEnabled                    uintptr
	DecreaseNetworkKeepAliveInterval      uintptr
	FlushTransport                        uintptr
}

type IControlChannelTrigger struct {
	win32.IInspectable
}

func (this *IControlChannelTrigger) Vtbl() *IControlChannelTriggerVtbl {
	return (*IControlChannelTriggerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IControlChannelTrigger) Get_ControlChannelTriggerId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ControlChannelTriggerId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IControlChannelTrigger) Get_ServerKeepAliveIntervalInMinutes() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServerKeepAliveIntervalInMinutes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IControlChannelTrigger) Put_ServerKeepAliveIntervalInMinutes(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ServerKeepAliveIntervalInMinutes, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IControlChannelTrigger) Get_CurrentKeepAliveIntervalInMinutes() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentKeepAliveIntervalInMinutes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IControlChannelTrigger) Get_TransportObject() interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TransportObject, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IControlChannelTrigger) Get_KeepAliveTrigger() *IBackgroundTrigger {
	var _result *IBackgroundTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KeepAliveTrigger, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IControlChannelTrigger) Get_PushNotificationTrigger() *IBackgroundTrigger {
	var _result *IBackgroundTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PushNotificationTrigger, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IControlChannelTrigger) UsingTransport(transport interface{}) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UsingTransport, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&transport)))
	_ = _hr
}

func (this *IControlChannelTrigger) WaitForPushEnabled() ControlChannelTriggerStatus {
	var _result ControlChannelTriggerStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WaitForPushEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IControlChannelTrigger) DecreaseNetworkKeepAliveInterval() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DecreaseNetworkKeepAliveInterval, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IControlChannelTrigger) FlushTransport() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FlushTransport, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// AF00D237-51BE-4514-9725-3556E1879580
var IID_IControlChannelTrigger2 = syscall.GUID{0xAF00D237, 0x51BE, 0x4514,
	[8]byte{0x97, 0x25, 0x35, 0x56, 0xE1, 0x87, 0x95, 0x80}}

type IControlChannelTrigger2Interface interface {
	win32.IInspectableInterface
	Get_IsWakeFromLowPowerSupported() bool
}

type IControlChannelTrigger2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsWakeFromLowPowerSupported uintptr
}

type IControlChannelTrigger2 struct {
	win32.IInspectable
}

func (this *IControlChannelTrigger2) Vtbl() *IControlChannelTrigger2Vtbl {
	return (*IControlChannelTrigger2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IControlChannelTrigger2) Get_IsWakeFromLowPowerSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsWakeFromLowPowerSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 1B36E047-89BB-4236-96AC-71D012BB4869
var IID_IControlChannelTriggerEventDetails = syscall.GUID{0x1B36E047, 0x89BB, 0x4236,
	[8]byte{0x96, 0xAC, 0x71, 0xD0, 0x12, 0xBB, 0x48, 0x69}}

type IControlChannelTriggerEventDetailsInterface interface {
	win32.IInspectableInterface
	Get_ControlChannelTrigger() *IControlChannelTrigger
}

type IControlChannelTriggerEventDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_ControlChannelTrigger uintptr
}

type IControlChannelTriggerEventDetails struct {
	win32.IInspectable
}

func (this *IControlChannelTriggerEventDetails) Vtbl() *IControlChannelTriggerEventDetailsVtbl {
	return (*IControlChannelTriggerEventDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IControlChannelTriggerEventDetails) Get_ControlChannelTrigger() *IControlChannelTrigger {
	var _result *IControlChannelTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ControlChannelTrigger, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DA4B7CF0-8D71-446F-88C3-B95184A2D6CD
var IID_IControlChannelTriggerFactory = syscall.GUID{0xDA4B7CF0, 0x8D71, 0x446F,
	[8]byte{0x88, 0xC3, 0xB9, 0x51, 0x84, 0xA2, 0xD6, 0xCD}}

type IControlChannelTriggerFactoryInterface interface {
	win32.IInspectableInterface
	CreateControlChannelTrigger(channelId string, serverKeepAliveIntervalInMinutes uint32) *IControlChannelTrigger
	CreateControlChannelTriggerEx(channelId string, serverKeepAliveIntervalInMinutes uint32, resourceRequestType ControlChannelTriggerResourceType) *IControlChannelTrigger
}

type IControlChannelTriggerFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateControlChannelTrigger   uintptr
	CreateControlChannelTriggerEx uintptr
}

type IControlChannelTriggerFactory struct {
	win32.IInspectable
}

func (this *IControlChannelTriggerFactory) Vtbl() *IControlChannelTriggerFactoryVtbl {
	return (*IControlChannelTriggerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IControlChannelTriggerFactory) CreateControlChannelTrigger(channelId string, serverKeepAliveIntervalInMinutes uint32) *IControlChannelTrigger {
	var _result *IControlChannelTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateControlChannelTrigger, uintptr(unsafe.Pointer(this)), NewHStr(channelId).Ptr, uintptr(serverKeepAliveIntervalInMinutes), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IControlChannelTriggerFactory) CreateControlChannelTriggerEx(channelId string, serverKeepAliveIntervalInMinutes uint32, resourceRequestType ControlChannelTriggerResourceType) *IControlChannelTrigger {
	var _result *IControlChannelTrigger
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateControlChannelTriggerEx, uintptr(unsafe.Pointer(this)), NewHStr(channelId).Ptr, uintptr(serverKeepAliveIntervalInMinutes), uintptr(resourceRequestType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6851038E-8EC4-42FE-9BB2-21E91B7BFCB1
var IID_IControlChannelTriggerResetEventDetails = syscall.GUID{0x6851038E, 0x8EC4, 0x42FE,
	[8]byte{0x9B, 0xB2, 0x21, 0xE9, 0x1B, 0x7B, 0xFC, 0xB1}}

type IControlChannelTriggerResetEventDetailsInterface interface {
	win32.IInspectableInterface
	Get_ResetReason() ControlChannelTriggerResetReason
	Get_HardwareSlotReset() bool
	Get_SoftwareSlotReset() bool
}

type IControlChannelTriggerResetEventDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_ResetReason       uintptr
	Get_HardwareSlotReset uintptr
	Get_SoftwareSlotReset uintptr
}

type IControlChannelTriggerResetEventDetails struct {
	win32.IInspectable
}

func (this *IControlChannelTriggerResetEventDetails) Vtbl() *IControlChannelTriggerResetEventDetailsVtbl {
	return (*IControlChannelTriggerResetEventDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IControlChannelTriggerResetEventDetails) Get_ResetReason() ControlChannelTriggerResetReason {
	var _result ControlChannelTriggerResetReason
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResetReason, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IControlChannelTriggerResetEventDetails) Get_HardwareSlotReset() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HardwareSlotReset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IControlChannelTriggerResetEventDetails) Get_SoftwareSlotReset() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SoftwareSlotReset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7FE25BBB-C3BC-4677-8446-CA28A465A3AF
var IID_IDatagramSocket = syscall.GUID{0x7FE25BBB, 0xC3BC, 0x4677,
	[8]byte{0x84, 0x46, 0xCA, 0x28, 0xA4, 0x65, 0xA3, 0xAF}}

type IDatagramSocketInterface interface {
	win32.IInspectableInterface
	Get_Control() *IDatagramSocketControl
	Get_Information() *IDatagramSocketInformation
	Get_OutputStream() *IOutputStream
	ConnectAsync(remoteHostName *IHostName, remoteServiceName string) *IAsyncAction
	ConnectWithEndpointPairAsync(endpointPair *IEndpointPair) *IAsyncAction
	BindServiceNameAsync(localServiceName string) *IAsyncAction
	BindEndpointAsync(localHostName *IHostName, localServiceName string) *IAsyncAction
	JoinMulticastGroup(host *IHostName)
	GetOutputStreamAsync(remoteHostName *IHostName, remoteServiceName string) *IAsyncOperation[*IOutputStream]
	GetOutputStreamWithEndpointPairAsync(endpointPair *IEndpointPair) *IAsyncOperation[*IOutputStream]
	Add_MessageReceived(eventHandler TypedEventHandler[*IDatagramSocket, *IDatagramSocketMessageReceivedEventArgs]) EventRegistrationToken
	Remove_MessageReceived(eventCookie EventRegistrationToken)
}

type IDatagramSocketVtbl struct {
	win32.IInspectableVtbl
	Get_Control                          uintptr
	Get_Information                      uintptr
	Get_OutputStream                     uintptr
	ConnectAsync                         uintptr
	ConnectWithEndpointPairAsync         uintptr
	BindServiceNameAsync                 uintptr
	BindEndpointAsync                    uintptr
	JoinMulticastGroup                   uintptr
	GetOutputStreamAsync                 uintptr
	GetOutputStreamWithEndpointPairAsync uintptr
	Add_MessageReceived                  uintptr
	Remove_MessageReceived               uintptr
}

type IDatagramSocket struct {
	win32.IInspectable
}

func (this *IDatagramSocket) Vtbl() *IDatagramSocketVtbl {
	return (*IDatagramSocketVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDatagramSocket) Get_Control() *IDatagramSocketControl {
	var _result *IDatagramSocketControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Control, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDatagramSocket) Get_Information() *IDatagramSocketInformation {
	var _result *IDatagramSocketInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Information, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDatagramSocket) Get_OutputStream() *IOutputStream {
	var _result *IOutputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutputStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDatagramSocket) ConnectAsync(remoteHostName *IHostName, remoteServiceName string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConnectAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(remoteHostName)), NewHStr(remoteServiceName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDatagramSocket) ConnectWithEndpointPairAsync(endpointPair *IEndpointPair) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConnectWithEndpointPairAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(endpointPair)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDatagramSocket) BindServiceNameAsync(localServiceName string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().BindServiceNameAsync, uintptr(unsafe.Pointer(this)), NewHStr(localServiceName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDatagramSocket) BindEndpointAsync(localHostName *IHostName, localServiceName string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().BindEndpointAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(localHostName)), NewHStr(localServiceName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDatagramSocket) JoinMulticastGroup(host *IHostName) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().JoinMulticastGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(host)))
	_ = _hr
}

func (this *IDatagramSocket) GetOutputStreamAsync(remoteHostName *IHostName, remoteServiceName string) *IAsyncOperation[*IOutputStream] {
	var _result *IAsyncOperation[*IOutputStream]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetOutputStreamAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(remoteHostName)), NewHStr(remoteServiceName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDatagramSocket) GetOutputStreamWithEndpointPairAsync(endpointPair *IEndpointPair) *IAsyncOperation[*IOutputStream] {
	var _result *IAsyncOperation[*IOutputStream]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetOutputStreamWithEndpointPairAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(endpointPair)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDatagramSocket) Add_MessageReceived(eventHandler TypedEventHandler[*IDatagramSocket, *IDatagramSocketMessageReceivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_MessageReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(eventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDatagramSocket) Remove_MessageReceived(eventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_MessageReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&eventCookie)))
	_ = _hr
}

// D83BA354-9A9D-4185-A20A-1424C9C2A7CD
var IID_IDatagramSocket2 = syscall.GUID{0xD83BA354, 0x9A9D, 0x4185,
	[8]byte{0xA2, 0x0A, 0x14, 0x24, 0xC9, 0xC2, 0xA7, 0xCD}}

type IDatagramSocket2Interface interface {
	win32.IInspectableInterface
	BindServiceNameAndAdapterAsync(localServiceName string, adapter *INetworkAdapter) *IAsyncAction
}

type IDatagramSocket2Vtbl struct {
	win32.IInspectableVtbl
	BindServiceNameAndAdapterAsync uintptr
}

type IDatagramSocket2 struct {
	win32.IInspectable
}

func (this *IDatagramSocket2) Vtbl() *IDatagramSocket2Vtbl {
	return (*IDatagramSocket2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDatagramSocket2) BindServiceNameAndAdapterAsync(localServiceName string, adapter *INetworkAdapter) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().BindServiceNameAndAdapterAsync, uintptr(unsafe.Pointer(this)), NewHStr(localServiceName).Ptr, uintptr(unsafe.Pointer(adapter)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 37544F09-AB92-4306-9AC1-0C381283D9C6
var IID_IDatagramSocket3 = syscall.GUID{0x37544F09, 0xAB92, 0x4306,
	[8]byte{0x9A, 0xC1, 0x0C, 0x38, 0x12, 0x83, 0xD9, 0xC6}}

type IDatagramSocket3Interface interface {
	win32.IInspectableInterface
	CancelIOAsync() *IAsyncAction
	EnableTransferOwnership(taskId syscall.GUID)
	EnableTransferOwnershipWithConnectedStandbyAction(taskId syscall.GUID, connectedStandbyAction SocketActivityConnectedStandbyAction)
	TransferOwnership(socketId string)
	TransferOwnershipWithContext(socketId string, data *ISocketActivityContext)
	TransferOwnershipWithContextAndKeepAliveTime(socketId string, data *ISocketActivityContext, keepAliveTime TimeSpan)
}

type IDatagramSocket3Vtbl struct {
	win32.IInspectableVtbl
	CancelIOAsync                                     uintptr
	EnableTransferOwnership                           uintptr
	EnableTransferOwnershipWithConnectedStandbyAction uintptr
	TransferOwnership                                 uintptr
	TransferOwnershipWithContext                      uintptr
	TransferOwnershipWithContextAndKeepAliveTime      uintptr
}

type IDatagramSocket3 struct {
	win32.IInspectable
}

func (this *IDatagramSocket3) Vtbl() *IDatagramSocket3Vtbl {
	return (*IDatagramSocket3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDatagramSocket3) CancelIOAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CancelIOAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDatagramSocket3) EnableTransferOwnership(taskId syscall.GUID) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnableTransferOwnership, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&taskId)))
	_ = _hr
}

func (this *IDatagramSocket3) EnableTransferOwnershipWithConnectedStandbyAction(taskId syscall.GUID, connectedStandbyAction SocketActivityConnectedStandbyAction) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnableTransferOwnershipWithConnectedStandbyAction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&taskId)), uintptr(connectedStandbyAction))
	_ = _hr
}

func (this *IDatagramSocket3) TransferOwnership(socketId string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TransferOwnership, uintptr(unsafe.Pointer(this)), NewHStr(socketId).Ptr)
	_ = _hr
}

func (this *IDatagramSocket3) TransferOwnershipWithContext(socketId string, data *ISocketActivityContext) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TransferOwnershipWithContext, uintptr(unsafe.Pointer(this)), NewHStr(socketId).Ptr, uintptr(unsafe.Pointer(data)))
	_ = _hr
}

func (this *IDatagramSocket3) TransferOwnershipWithContextAndKeepAliveTime(socketId string, data *ISocketActivityContext, keepAliveTime TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TransferOwnershipWithContextAndKeepAliveTime, uintptr(unsafe.Pointer(this)), NewHStr(socketId).Ptr, uintptr(unsafe.Pointer(data)), *(*uintptr)(unsafe.Pointer(&keepAliveTime)))
	_ = _hr
}

// 52AC3F2E-349A-4135-BB58-B79B2647D390
var IID_IDatagramSocketControl = syscall.GUID{0x52AC3F2E, 0x349A, 0x4135,
	[8]byte{0xBB, 0x58, 0xB7, 0x9B, 0x26, 0x47, 0xD3, 0x90}}

type IDatagramSocketControlInterface interface {
	win32.IInspectableInterface
	Get_QualityOfService() SocketQualityOfService
	Put_QualityOfService(value SocketQualityOfService)
	Get_OutboundUnicastHopLimit() byte
	Put_OutboundUnicastHopLimit(value byte)
}

type IDatagramSocketControlVtbl struct {
	win32.IInspectableVtbl
	Get_QualityOfService        uintptr
	Put_QualityOfService        uintptr
	Get_OutboundUnicastHopLimit uintptr
	Put_OutboundUnicastHopLimit uintptr
}

type IDatagramSocketControl struct {
	win32.IInspectable
}

func (this *IDatagramSocketControl) Vtbl() *IDatagramSocketControlVtbl {
	return (*IDatagramSocketControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDatagramSocketControl) Get_QualityOfService() SocketQualityOfService {
	var _result SocketQualityOfService
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_QualityOfService, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDatagramSocketControl) Put_QualityOfService(value SocketQualityOfService) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_QualityOfService, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDatagramSocketControl) Get_OutboundUnicastHopLimit() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutboundUnicastHopLimit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDatagramSocketControl) Put_OutboundUnicastHopLimit(value byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_OutboundUnicastHopLimit, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 33EAD5C2-979C-4415-82A1-3CFAF646C192
var IID_IDatagramSocketControl2 = syscall.GUID{0x33EAD5C2, 0x979C, 0x4415,
	[8]byte{0x82, 0xA1, 0x3C, 0xFA, 0xF6, 0x46, 0xC1, 0x92}}

type IDatagramSocketControl2Interface interface {
	win32.IInspectableInterface
	Get_InboundBufferSizeInBytes() uint32
	Put_InboundBufferSizeInBytes(value uint32)
	Get_DontFragment() bool
	Put_DontFragment(value bool)
}

type IDatagramSocketControl2Vtbl struct {
	win32.IInspectableVtbl
	Get_InboundBufferSizeInBytes uintptr
	Put_InboundBufferSizeInBytes uintptr
	Get_DontFragment             uintptr
	Put_DontFragment             uintptr
}

type IDatagramSocketControl2 struct {
	win32.IInspectable
}

func (this *IDatagramSocketControl2) Vtbl() *IDatagramSocketControl2Vtbl {
	return (*IDatagramSocketControl2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDatagramSocketControl2) Get_InboundBufferSizeInBytes() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InboundBufferSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDatagramSocketControl2) Put_InboundBufferSizeInBytes(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InboundBufferSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDatagramSocketControl2) Get_DontFragment() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DontFragment, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDatagramSocketControl2) Put_DontFragment(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DontFragment, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// D4EB8256-1F6D-4598-9B57-D42A001DF349
var IID_IDatagramSocketControl3 = syscall.GUID{0xD4EB8256, 0x1F6D, 0x4598,
	[8]byte{0x9B, 0x57, 0xD4, 0x2A, 0x00, 0x1D, 0xF3, 0x49}}

type IDatagramSocketControl3Interface interface {
	win32.IInspectableInterface
	Get_MulticastOnly() bool
	Put_MulticastOnly(value bool)
}

type IDatagramSocketControl3Vtbl struct {
	win32.IInspectableVtbl
	Get_MulticastOnly uintptr
	Put_MulticastOnly uintptr
}

type IDatagramSocketControl3 struct {
	win32.IInspectable
}

func (this *IDatagramSocketControl3) Vtbl() *IDatagramSocketControl3Vtbl {
	return (*IDatagramSocketControl3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDatagramSocketControl3) Get_MulticastOnly() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MulticastOnly, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDatagramSocketControl3) Put_MulticastOnly(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MulticastOnly, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 5F1A569A-55FB-48CD-9706-7A974F7B1585
var IID_IDatagramSocketInformation = syscall.GUID{0x5F1A569A, 0x55FB, 0x48CD,
	[8]byte{0x97, 0x06, 0x7A, 0x97, 0x4F, 0x7B, 0x15, 0x85}}

type IDatagramSocketInformationInterface interface {
	win32.IInspectableInterface
	Get_LocalAddress() *IHostName
	Get_LocalPort() string
	Get_RemoteAddress() *IHostName
	Get_RemotePort() string
}

type IDatagramSocketInformationVtbl struct {
	win32.IInspectableVtbl
	Get_LocalAddress  uintptr
	Get_LocalPort     uintptr
	Get_RemoteAddress uintptr
	Get_RemotePort    uintptr
}

type IDatagramSocketInformation struct {
	win32.IInspectable
}

func (this *IDatagramSocketInformation) Vtbl() *IDatagramSocketInformationVtbl {
	return (*IDatagramSocketInformationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDatagramSocketInformation) Get_LocalAddress() *IHostName {
	var _result *IHostName
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocalAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDatagramSocketInformation) Get_LocalPort() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocalPort, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDatagramSocketInformation) Get_RemoteAddress() *IHostName {
	var _result *IHostName
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RemoteAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDatagramSocketInformation) Get_RemotePort() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RemotePort, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 9E2DDCA2-1712-4CE4-B179-8C652C6D107E
var IID_IDatagramSocketMessageReceivedEventArgs = syscall.GUID{0x9E2DDCA2, 0x1712, 0x4CE4,
	[8]byte{0xB1, 0x79, 0x8C, 0x65, 0x2C, 0x6D, 0x10, 0x7E}}

type IDatagramSocketMessageReceivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_RemoteAddress() *IHostName
	Get_RemotePort() string
	Get_LocalAddress() *IHostName
	GetDataReader() *IDataReader
	GetDataStream() *IInputStream
}

type IDatagramSocketMessageReceivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_RemoteAddress uintptr
	Get_RemotePort    uintptr
	Get_LocalAddress  uintptr
	GetDataReader     uintptr
	GetDataStream     uintptr
}

type IDatagramSocketMessageReceivedEventArgs struct {
	win32.IInspectable
}

func (this *IDatagramSocketMessageReceivedEventArgs) Vtbl() *IDatagramSocketMessageReceivedEventArgsVtbl {
	return (*IDatagramSocketMessageReceivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDatagramSocketMessageReceivedEventArgs) Get_RemoteAddress() *IHostName {
	var _result *IHostName
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RemoteAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDatagramSocketMessageReceivedEventArgs) Get_RemotePort() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RemotePort, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDatagramSocketMessageReceivedEventArgs) Get_LocalAddress() *IHostName {
	var _result *IHostName
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocalAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDatagramSocketMessageReceivedEventArgs) GetDataReader() *IDataReader {
	var _result *IDataReader
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDataReader, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDatagramSocketMessageReceivedEventArgs) GetDataStream() *IInputStream {
	var _result *IInputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDataStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E9C62AEE-1494-4A21-BB7E-8589FC751D9D
var IID_IDatagramSocketStatics = syscall.GUID{0xE9C62AEE, 0x1494, 0x4A21,
	[8]byte{0xBB, 0x7E, 0x85, 0x89, 0xFC, 0x75, 0x1D, 0x9D}}

type IDatagramSocketStaticsInterface interface {
	win32.IInspectableInterface
	GetEndpointPairsAsync(remoteHostName *IHostName, remoteServiceName string) *IAsyncOperation[*IVectorView[*IEndpointPair]]
	GetEndpointPairsWithSortOptionsAsync(remoteHostName *IHostName, remoteServiceName string, sortOptions HostNameSortOptions) *IAsyncOperation[*IVectorView[*IEndpointPair]]
}

type IDatagramSocketStaticsVtbl struct {
	win32.IInspectableVtbl
	GetEndpointPairsAsync                uintptr
	GetEndpointPairsWithSortOptionsAsync uintptr
}

type IDatagramSocketStatics struct {
	win32.IInspectable
}

func (this *IDatagramSocketStatics) Vtbl() *IDatagramSocketStaticsVtbl {
	return (*IDatagramSocketStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDatagramSocketStatics) GetEndpointPairsAsync(remoteHostName *IHostName, remoteServiceName string) *IAsyncOperation[*IVectorView[*IEndpointPair]] {
	var _result *IAsyncOperation[*IVectorView[*IEndpointPair]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetEndpointPairsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(remoteHostName)), NewHStr(remoteServiceName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDatagramSocketStatics) GetEndpointPairsWithSortOptionsAsync(remoteHostName *IHostName, remoteServiceName string, sortOptions HostNameSortOptions) *IAsyncOperation[*IVectorView[*IEndpointPair]] {
	var _result *IAsyncOperation[*IVectorView[*IEndpointPair]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetEndpointPairsWithSortOptionsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(remoteHostName)), NewHStr(remoteServiceName).Ptr, uintptr(sortOptions), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 33727D08-34D5-4746-AD7B-8DDE5BC2EF88
var IID_IMessageWebSocket = syscall.GUID{0x33727D08, 0x34D5, 0x4746,
	[8]byte{0xAD, 0x7B, 0x8D, 0xDE, 0x5B, 0xC2, 0xEF, 0x88}}

type IMessageWebSocketInterface interface {
	win32.IInspectableInterface
	Get_Control() *IMessageWebSocketControl
	Get_Information() *IWebSocketInformation
	Add_MessageReceived(eventHandler TypedEventHandler[*IMessageWebSocket, *IMessageWebSocketMessageReceivedEventArgs]) EventRegistrationToken
	Remove_MessageReceived(eventCookie EventRegistrationToken)
}

type IMessageWebSocketVtbl struct {
	win32.IInspectableVtbl
	Get_Control            uintptr
	Get_Information        uintptr
	Add_MessageReceived    uintptr
	Remove_MessageReceived uintptr
}

type IMessageWebSocket struct {
	win32.IInspectable
}

func (this *IMessageWebSocket) Vtbl() *IMessageWebSocketVtbl {
	return (*IMessageWebSocketVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMessageWebSocket) Get_Control() *IMessageWebSocketControl {
	var _result *IMessageWebSocketControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Control, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMessageWebSocket) Get_Information() *IWebSocketInformation {
	var _result *IWebSocketInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Information, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMessageWebSocket) Add_MessageReceived(eventHandler TypedEventHandler[*IMessageWebSocket, *IMessageWebSocketMessageReceivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_MessageReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(eventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMessageWebSocket) Remove_MessageReceived(eventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_MessageReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&eventCookie)))
	_ = _hr
}

// BED0CEE7-F9C8-440A-9AD5-737281D9742E
var IID_IMessageWebSocket2 = syscall.GUID{0xBED0CEE7, 0xF9C8, 0x440A,
	[8]byte{0x9A, 0xD5, 0x73, 0x72, 0x81, 0xD9, 0x74, 0x2E}}

type IMessageWebSocket2Interface interface {
	win32.IInspectableInterface
	Add_ServerCustomValidationRequested(eventHandler TypedEventHandler[*IMessageWebSocket, *IWebSocketServerCustomValidationRequestedEventArgs]) EventRegistrationToken
	Remove_ServerCustomValidationRequested(eventCookie EventRegistrationToken)
}

type IMessageWebSocket2Vtbl struct {
	win32.IInspectableVtbl
	Add_ServerCustomValidationRequested    uintptr
	Remove_ServerCustomValidationRequested uintptr
}

type IMessageWebSocket2 struct {
	win32.IInspectable
}

func (this *IMessageWebSocket2) Vtbl() *IMessageWebSocket2Vtbl {
	return (*IMessageWebSocket2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMessageWebSocket2) Add_ServerCustomValidationRequested(eventHandler TypedEventHandler[*IMessageWebSocket, *IWebSocketServerCustomValidationRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ServerCustomValidationRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(eventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMessageWebSocket2) Remove_ServerCustomValidationRequested(eventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ServerCustomValidationRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&eventCookie)))
	_ = _hr
}

// 59D9DEFB-71AF-4349-8487-911FCF681597
var IID_IMessageWebSocket3 = syscall.GUID{0x59D9DEFB, 0x71AF, 0x4349,
	[8]byte{0x84, 0x87, 0x91, 0x1F, 0xCF, 0x68, 0x15, 0x97}}

type IMessageWebSocket3Interface interface {
	win32.IInspectableInterface
	SendNonfinalFrameAsync(data *IBuffer) *IAsyncOperationWithProgress[uint32, uint32]
	SendFinalFrameAsync(data *IBuffer) *IAsyncOperationWithProgress[uint32, uint32]
}

type IMessageWebSocket3Vtbl struct {
	win32.IInspectableVtbl
	SendNonfinalFrameAsync uintptr
	SendFinalFrameAsync    uintptr
}

type IMessageWebSocket3 struct {
	win32.IInspectable
}

func (this *IMessageWebSocket3) Vtbl() *IMessageWebSocket3Vtbl {
	return (*IMessageWebSocket3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMessageWebSocket3) SendNonfinalFrameAsync(data *IBuffer) *IAsyncOperationWithProgress[uint32, uint32] {
	var _result *IAsyncOperationWithProgress[uint32, uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendNonfinalFrameAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMessageWebSocket3) SendFinalFrameAsync(data *IBuffer) *IAsyncOperationWithProgress[uint32, uint32] {
	var _result *IAsyncOperationWithProgress[uint32, uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendFinalFrameAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8118388A-C629-4F0A-80FB-81FC05538862
var IID_IMessageWebSocketControl = syscall.GUID{0x8118388A, 0xC629, 0x4F0A,
	[8]byte{0x80, 0xFB, 0x81, 0xFC, 0x05, 0x53, 0x88, 0x62}}

type IMessageWebSocketControlInterface interface {
	win32.IInspectableInterface
	Get_MaxMessageSize() uint32
	Put_MaxMessageSize(value uint32)
	Get_MessageType() SocketMessageType
	Put_MessageType(value SocketMessageType)
}

type IMessageWebSocketControlVtbl struct {
	win32.IInspectableVtbl
	Get_MaxMessageSize uintptr
	Put_MaxMessageSize uintptr
	Get_MessageType    uintptr
	Put_MessageType    uintptr
}

type IMessageWebSocketControl struct {
	win32.IInspectable
}

func (this *IMessageWebSocketControl) Vtbl() *IMessageWebSocketControlVtbl {
	return (*IMessageWebSocketControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMessageWebSocketControl) Get_MaxMessageSize() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxMessageSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMessageWebSocketControl) Put_MaxMessageSize(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxMessageSize, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMessageWebSocketControl) Get_MessageType() SocketMessageType {
	var _result SocketMessageType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MessageType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMessageWebSocketControl) Put_MessageType(value SocketMessageType) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MessageType, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// E30FD791-080C-400A-A712-27DFA9E744D8
var IID_IMessageWebSocketControl2 = syscall.GUID{0xE30FD791, 0x080C, 0x400A,
	[8]byte{0xA7, 0x12, 0x27, 0xDF, 0xA9, 0xE7, 0x44, 0xD8}}

type IMessageWebSocketControl2Interface interface {
	win32.IInspectableInterface
	Get_DesiredUnsolicitedPongInterval() TimeSpan
	Put_DesiredUnsolicitedPongInterval(value TimeSpan)
	Get_ActualUnsolicitedPongInterval() TimeSpan
	Get_ReceiveMode() MessageWebSocketReceiveMode
	Put_ReceiveMode(value MessageWebSocketReceiveMode)
	Get_ClientCertificate() *ICertificate
	Put_ClientCertificate(value *ICertificate)
}

type IMessageWebSocketControl2Vtbl struct {
	win32.IInspectableVtbl
	Get_DesiredUnsolicitedPongInterval uintptr
	Put_DesiredUnsolicitedPongInterval uintptr
	Get_ActualUnsolicitedPongInterval  uintptr
	Get_ReceiveMode                    uintptr
	Put_ReceiveMode                    uintptr
	Get_ClientCertificate              uintptr
	Put_ClientCertificate              uintptr
}

type IMessageWebSocketControl2 struct {
	win32.IInspectable
}

func (this *IMessageWebSocketControl2) Vtbl() *IMessageWebSocketControl2Vtbl {
	return (*IMessageWebSocketControl2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMessageWebSocketControl2) Get_DesiredUnsolicitedPongInterval() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesiredUnsolicitedPongInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMessageWebSocketControl2) Put_DesiredUnsolicitedPongInterval(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DesiredUnsolicitedPongInterval, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IMessageWebSocketControl2) Get_ActualUnsolicitedPongInterval() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ActualUnsolicitedPongInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMessageWebSocketControl2) Get_ReceiveMode() MessageWebSocketReceiveMode {
	var _result MessageWebSocketReceiveMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReceiveMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMessageWebSocketControl2) Put_ReceiveMode(value MessageWebSocketReceiveMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReceiveMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMessageWebSocketControl2) Get_ClientCertificate() *ICertificate {
	var _result *ICertificate
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ClientCertificate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMessageWebSocketControl2) Put_ClientCertificate(value *ICertificate) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ClientCertificate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 478C22AC-4C4B-42ED-9ED7-1EF9F94FA3D5
var IID_IMessageWebSocketMessageReceivedEventArgs = syscall.GUID{0x478C22AC, 0x4C4B, 0x42ED,
	[8]byte{0x9E, 0xD7, 0x1E, 0xF9, 0xF9, 0x4F, 0xA3, 0xD5}}

type IMessageWebSocketMessageReceivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_MessageType() SocketMessageType
	GetDataReader() *IDataReader
	GetDataStream() *IInputStream
}

type IMessageWebSocketMessageReceivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_MessageType uintptr
	GetDataReader   uintptr
	GetDataStream   uintptr
}

type IMessageWebSocketMessageReceivedEventArgs struct {
	win32.IInspectable
}

func (this *IMessageWebSocketMessageReceivedEventArgs) Vtbl() *IMessageWebSocketMessageReceivedEventArgsVtbl {
	return (*IMessageWebSocketMessageReceivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMessageWebSocketMessageReceivedEventArgs) Get_MessageType() SocketMessageType {
	var _result SocketMessageType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MessageType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMessageWebSocketMessageReceivedEventArgs) GetDataReader() *IDataReader {
	var _result *IDataReader
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDataReader, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMessageWebSocketMessageReceivedEventArgs) GetDataStream() *IInputStream {
	var _result *IInputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDataStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 89CE06FD-DD6F-4A07-87F9-F9EB4D89D83D
var IID_IMessageWebSocketMessageReceivedEventArgs2 = syscall.GUID{0x89CE06FD, 0xDD6F, 0x4A07,
	[8]byte{0x87, 0xF9, 0xF9, 0xEB, 0x4D, 0x89, 0xD8, 0x3D}}

type IMessageWebSocketMessageReceivedEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_IsMessageComplete() bool
}

type IMessageWebSocketMessageReceivedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsMessageComplete uintptr
}

type IMessageWebSocketMessageReceivedEventArgs2 struct {
	win32.IInspectable
}

func (this *IMessageWebSocketMessageReceivedEventArgs2) Vtbl() *IMessageWebSocketMessageReceivedEventArgs2Vtbl {
	return (*IMessageWebSocketMessageReceivedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMessageWebSocketMessageReceivedEventArgs2) Get_IsMessageComplete() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsMessageComplete, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// E3AC9240-813B-5EFD-7E11-AE2305FC77F1
var IID_IServerMessageWebSocket = syscall.GUID{0xE3AC9240, 0x813B, 0x5EFD,
	[8]byte{0x7E, 0x11, 0xAE, 0x23, 0x05, 0xFC, 0x77, 0xF1}}

type IServerMessageWebSocketInterface interface {
	win32.IInspectableInterface
	Add_MessageReceived(value TypedEventHandler[*IServerMessageWebSocket, *IMessageWebSocketMessageReceivedEventArgs]) EventRegistrationToken
	Remove_MessageReceived(token EventRegistrationToken)
	Get_Control() *IServerMessageWebSocketControl
	Get_Information() *IServerMessageWebSocketInformation
	Get_OutputStream() *IOutputStream
	Add_Closed(value TypedEventHandler[*IServerMessageWebSocket, *IWebSocketClosedEventArgs]) EventRegistrationToken
	Remove_Closed(token EventRegistrationToken)
	CloseWithStatus(code uint16, reason string)
}

type IServerMessageWebSocketVtbl struct {
	win32.IInspectableVtbl
	Add_MessageReceived    uintptr
	Remove_MessageReceived uintptr
	Get_Control            uintptr
	Get_Information        uintptr
	Get_OutputStream       uintptr
	Add_Closed             uintptr
	Remove_Closed          uintptr
	CloseWithStatus        uintptr
}

type IServerMessageWebSocket struct {
	win32.IInspectable
}

func (this *IServerMessageWebSocket) Vtbl() *IServerMessageWebSocketVtbl {
	return (*IServerMessageWebSocketVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IServerMessageWebSocket) Add_MessageReceived(value TypedEventHandler[*IServerMessageWebSocket, *IMessageWebSocketMessageReceivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_MessageReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IServerMessageWebSocket) Remove_MessageReceived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_MessageReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IServerMessageWebSocket) Get_Control() *IServerMessageWebSocketControl {
	var _result *IServerMessageWebSocketControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Control, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IServerMessageWebSocket) Get_Information() *IServerMessageWebSocketInformation {
	var _result *IServerMessageWebSocketInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Information, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IServerMessageWebSocket) Get_OutputStream() *IOutputStream {
	var _result *IOutputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutputStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IServerMessageWebSocket) Add_Closed(value TypedEventHandler[*IServerMessageWebSocket, *IWebSocketClosedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Closed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IServerMessageWebSocket) Remove_Closed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Closed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IServerMessageWebSocket) CloseWithStatus(code uint16, reason string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CloseWithStatus, uintptr(unsafe.Pointer(this)), uintptr(code), NewHStr(reason).Ptr)
	_ = _hr
}

// 69C2F051-1C1F-587A-4519-2181610192B7
var IID_IServerMessageWebSocketControl = syscall.GUID{0x69C2F051, 0x1C1F, 0x587A,
	[8]byte{0x45, 0x19, 0x21, 0x81, 0x61, 0x01, 0x92, 0xB7}}

type IServerMessageWebSocketControlInterface interface {
	win32.IInspectableInterface
	Get_MessageType() SocketMessageType
	Put_MessageType(value SocketMessageType)
}

type IServerMessageWebSocketControlVtbl struct {
	win32.IInspectableVtbl
	Get_MessageType uintptr
	Put_MessageType uintptr
}

type IServerMessageWebSocketControl struct {
	win32.IInspectable
}

func (this *IServerMessageWebSocketControl) Vtbl() *IServerMessageWebSocketControlVtbl {
	return (*IServerMessageWebSocketControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IServerMessageWebSocketControl) Get_MessageType() SocketMessageType {
	var _result SocketMessageType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MessageType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IServerMessageWebSocketControl) Put_MessageType(value SocketMessageType) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MessageType, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// FC32B45F-4448-5505-6CC9-09AFA8915F5D
var IID_IServerMessageWebSocketInformation = syscall.GUID{0xFC32B45F, 0x4448, 0x5505,
	[8]byte{0x6C, 0xC9, 0x09, 0xAF, 0xA8, 0x91, 0x5F, 0x5D}}

type IServerMessageWebSocketInformationInterface interface {
	win32.IInspectableInterface
	Get_BandwidthStatistics() BandwidthStatistics
	Get_Protocol() string
	Get_LocalAddress() *IHostName
}

type IServerMessageWebSocketInformationVtbl struct {
	win32.IInspectableVtbl
	Get_BandwidthStatistics uintptr
	Get_Protocol            uintptr
	Get_LocalAddress        uintptr
}

type IServerMessageWebSocketInformation struct {
	win32.IInspectable
}

func (this *IServerMessageWebSocketInformation) Vtbl() *IServerMessageWebSocketInformationVtbl {
	return (*IServerMessageWebSocketInformationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IServerMessageWebSocketInformation) Get_BandwidthStatistics() BandwidthStatistics {
	var _result BandwidthStatistics
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BandwidthStatistics, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IServerMessageWebSocketInformation) Get_Protocol() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Protocol, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IServerMessageWebSocketInformation) Get_LocalAddress() *IHostName {
	var _result *IHostName
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocalAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2CED5BBF-74F6-55E4-79DF-9132680DFEE8
var IID_IServerStreamWebSocket = syscall.GUID{0x2CED5BBF, 0x74F6, 0x55E4,
	[8]byte{0x79, 0xDF, 0x91, 0x32, 0x68, 0x0D, 0xFE, 0xE8}}

type IServerStreamWebSocketInterface interface {
	win32.IInspectableInterface
	Get_Information() *IServerStreamWebSocketInformation
	Get_InputStream() *IInputStream
	Get_OutputStream() *IOutputStream
	Add_Closed(value TypedEventHandler[*IServerStreamWebSocket, *IWebSocketClosedEventArgs]) EventRegistrationToken
	Remove_Closed(token EventRegistrationToken)
	CloseWithStatus(code uint16, reason string)
}

type IServerStreamWebSocketVtbl struct {
	win32.IInspectableVtbl
	Get_Information  uintptr
	Get_InputStream  uintptr
	Get_OutputStream uintptr
	Add_Closed       uintptr
	Remove_Closed    uintptr
	CloseWithStatus  uintptr
}

type IServerStreamWebSocket struct {
	win32.IInspectable
}

func (this *IServerStreamWebSocket) Vtbl() *IServerStreamWebSocketVtbl {
	return (*IServerStreamWebSocketVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IServerStreamWebSocket) Get_Information() *IServerStreamWebSocketInformation {
	var _result *IServerStreamWebSocketInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Information, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IServerStreamWebSocket) Get_InputStream() *IInputStream {
	var _result *IInputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InputStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IServerStreamWebSocket) Get_OutputStream() *IOutputStream {
	var _result *IOutputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutputStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IServerStreamWebSocket) Add_Closed(value TypedEventHandler[*IServerStreamWebSocket, *IWebSocketClosedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Closed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IServerStreamWebSocket) Remove_Closed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Closed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IServerStreamWebSocket) CloseWithStatus(code uint16, reason string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CloseWithStatus, uintptr(unsafe.Pointer(this)), uintptr(code), NewHStr(reason).Ptr)
	_ = _hr
}

// FC32B45F-4448-5505-6CC9-09ABA8915F5D
var IID_IServerStreamWebSocketInformation = syscall.GUID{0xFC32B45F, 0x4448, 0x5505,
	[8]byte{0x6C, 0xC9, 0x09, 0xAB, 0xA8, 0x91, 0x5F, 0x5D}}

type IServerStreamWebSocketInformationInterface interface {
	win32.IInspectableInterface
	Get_BandwidthStatistics() BandwidthStatistics
	Get_Protocol() string
	Get_LocalAddress() *IHostName
}

type IServerStreamWebSocketInformationVtbl struct {
	win32.IInspectableVtbl
	Get_BandwidthStatistics uintptr
	Get_Protocol            uintptr
	Get_LocalAddress        uintptr
}

type IServerStreamWebSocketInformation struct {
	win32.IInspectable
}

func (this *IServerStreamWebSocketInformation) Vtbl() *IServerStreamWebSocketInformationVtbl {
	return (*IServerStreamWebSocketInformationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IServerStreamWebSocketInformation) Get_BandwidthStatistics() BandwidthStatistics {
	var _result BandwidthStatistics
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BandwidthStatistics, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IServerStreamWebSocketInformation) Get_Protocol() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Protocol, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IServerStreamWebSocketInformation) Get_LocalAddress() *IHostName {
	var _result *IHostName
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocalAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 43B04D64-4C85-4396-A637-1D973F6EBD49
var IID_ISocketActivityContext = syscall.GUID{0x43B04D64, 0x4C85, 0x4396,
	[8]byte{0xA6, 0x37, 0x1D, 0x97, 0x3F, 0x6E, 0xBD, 0x49}}

type ISocketActivityContextInterface interface {
	win32.IInspectableInterface
	Get_Data() *IBuffer
}

type ISocketActivityContextVtbl struct {
	win32.IInspectableVtbl
	Get_Data uintptr
}

type ISocketActivityContext struct {
	win32.IInspectable
}

func (this *ISocketActivityContext) Vtbl() *ISocketActivityContextVtbl {
	return (*ISocketActivityContextVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISocketActivityContext) Get_Data() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B99FC3C3-088C-4388-83AE-2525138E049A
var IID_ISocketActivityContextFactory = syscall.GUID{0xB99FC3C3, 0x088C, 0x4388,
	[8]byte{0x83, 0xAE, 0x25, 0x25, 0x13, 0x8E, 0x04, 0x9A}}

type ISocketActivityContextFactoryInterface interface {
	win32.IInspectableInterface
	Create(data *IBuffer) *ISocketActivityContext
}

type ISocketActivityContextFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ISocketActivityContextFactory struct {
	win32.IInspectable
}

func (this *ISocketActivityContextFactory) Vtbl() *ISocketActivityContextFactoryVtbl {
	return (*ISocketActivityContextFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISocketActivityContextFactory) Create(data *IBuffer) *ISocketActivityContext {
	var _result *ISocketActivityContext
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8D8A42E4-A87E-4B74-9968-185B2511DEFE
var IID_ISocketActivityInformation = syscall.GUID{0x8D8A42E4, 0xA87E, 0x4B74,
	[8]byte{0x99, 0x68, 0x18, 0x5B, 0x25, 0x11, 0xDE, 0xFE}}

type ISocketActivityInformationInterface interface {
	win32.IInspectableInterface
	Get_TaskId() syscall.GUID
	Get_Id() string
	Get_SocketKind() SocketActivityKind
	Get_Context() *ISocketActivityContext
	Get_DatagramSocket() *IDatagramSocket
	Get_StreamSocket() *IStreamSocket
	Get_StreamSocketListener() *IStreamSocketListener
}

type ISocketActivityInformationVtbl struct {
	win32.IInspectableVtbl
	Get_TaskId               uintptr
	Get_Id                   uintptr
	Get_SocketKind           uintptr
	Get_Context              uintptr
	Get_DatagramSocket       uintptr
	Get_StreamSocket         uintptr
	Get_StreamSocketListener uintptr
}

type ISocketActivityInformation struct {
	win32.IInspectable
}

func (this *ISocketActivityInformation) Vtbl() *ISocketActivityInformationVtbl {
	return (*ISocketActivityInformationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISocketActivityInformation) Get_TaskId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TaskId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISocketActivityInformation) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISocketActivityInformation) Get_SocketKind() SocketActivityKind {
	var _result SocketActivityKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SocketKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISocketActivityInformation) Get_Context() *ISocketActivityContext {
	var _result *ISocketActivityContext
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Context, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISocketActivityInformation) Get_DatagramSocket() *IDatagramSocket {
	var _result *IDatagramSocket
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DatagramSocket, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISocketActivityInformation) Get_StreamSocket() *IStreamSocket {
	var _result *IStreamSocket
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StreamSocket, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISocketActivityInformation) Get_StreamSocketListener() *IStreamSocketListener {
	var _result *IStreamSocketListener
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StreamSocketListener, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8570B47A-7E7D-4736-8041-1327A6543C56
var IID_ISocketActivityInformationStatics = syscall.GUID{0x8570B47A, 0x7E7D, 0x4736,
	[8]byte{0x80, 0x41, 0x13, 0x27, 0xA6, 0x54, 0x3C, 0x56}}

type ISocketActivityInformationStaticsInterface interface {
	win32.IInspectableInterface
	Get_AllSockets() *IMapView[string, *ISocketActivityInformation]
}

type ISocketActivityInformationStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_AllSockets uintptr
}

type ISocketActivityInformationStatics struct {
	win32.IInspectable
}

func (this *ISocketActivityInformationStatics) Vtbl() *ISocketActivityInformationStaticsVtbl {
	return (*ISocketActivityInformationStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISocketActivityInformationStatics) Get_AllSockets() *IMapView[string, *ISocketActivityInformation] {
	var _result *IMapView[string, *ISocketActivityInformation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllSockets, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 45F406A7-FC9F-4F81-ACAD-355FEF51E67B
var IID_ISocketActivityTriggerDetails = syscall.GUID{0x45F406A7, 0xFC9F, 0x4F81,
	[8]byte{0xAC, 0xAD, 0x35, 0x5F, 0xEF, 0x51, 0xE6, 0x7B}}

type ISocketActivityTriggerDetailsInterface interface {
	win32.IInspectableInterface
	Get_Reason() SocketActivityTriggerReason
	Get_SocketInformation() *ISocketActivityInformation
}

type ISocketActivityTriggerDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_Reason            uintptr
	Get_SocketInformation uintptr
}

type ISocketActivityTriggerDetails struct {
	win32.IInspectable
}

func (this *ISocketActivityTriggerDetails) Vtbl() *ISocketActivityTriggerDetailsVtbl {
	return (*ISocketActivityTriggerDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISocketActivityTriggerDetails) Get_Reason() SocketActivityTriggerReason {
	var _result SocketActivityTriggerReason
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Reason, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISocketActivityTriggerDetails) Get_SocketInformation() *ISocketActivityInformation {
	var _result *ISocketActivityInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SocketInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 828337F4-7D56-4D8E-B7B4-A07DD7C1BCA9
var IID_ISocketErrorStatics = syscall.GUID{0x828337F4, 0x7D56, 0x4D8E,
	[8]byte{0xB7, 0xB4, 0xA0, 0x7D, 0xD7, 0xC1, 0xBC, 0xA9}}

type ISocketErrorStaticsInterface interface {
	win32.IInspectableInterface
	GetStatus(hresult int32) SocketErrorStatus
}

type ISocketErrorStaticsVtbl struct {
	win32.IInspectableVtbl
	GetStatus uintptr
}

type ISocketErrorStatics struct {
	win32.IInspectable
}

func (this *ISocketErrorStatics) Vtbl() *ISocketErrorStaticsVtbl {
	return (*ISocketErrorStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISocketErrorStatics) GetStatus(hresult int32) SocketErrorStatus {
	var _result SocketErrorStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStatus, uintptr(unsafe.Pointer(this)), uintptr(hresult), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 69A22CF3-FC7B-4857-AF38-F6E7DE6A5B49
var IID_IStreamSocket = syscall.GUID{0x69A22CF3, 0xFC7B, 0x4857,
	[8]byte{0xAF, 0x38, 0xF6, 0xE7, 0xDE, 0x6A, 0x5B, 0x49}}

type IStreamSocketInterface interface {
	win32.IInspectableInterface
	Get_Control() *IStreamSocketControl
	Get_Information() *IStreamSocketInformation
	Get_InputStream() *IInputStream
	Get_OutputStream() *IOutputStream
	ConnectWithEndpointPairAsync(endpointPair *IEndpointPair) *IAsyncAction
	ConnectAsync(remoteHostName *IHostName, remoteServiceName string) *IAsyncAction
	ConnectWithEndpointPairAndProtectionLevelAsync(endpointPair *IEndpointPair, protectionLevel SocketProtectionLevel) *IAsyncAction
	ConnectWithProtectionLevelAsync(remoteHostName *IHostName, remoteServiceName string, protectionLevel SocketProtectionLevel) *IAsyncAction
	UpgradeToSslAsync(protectionLevel SocketProtectionLevel, validationHostName *IHostName) *IAsyncAction
}

type IStreamSocketVtbl struct {
	win32.IInspectableVtbl
	Get_Control                                    uintptr
	Get_Information                                uintptr
	Get_InputStream                                uintptr
	Get_OutputStream                               uintptr
	ConnectWithEndpointPairAsync                   uintptr
	ConnectAsync                                   uintptr
	ConnectWithEndpointPairAndProtectionLevelAsync uintptr
	ConnectWithProtectionLevelAsync                uintptr
	UpgradeToSslAsync                              uintptr
}

type IStreamSocket struct {
	win32.IInspectable
}

func (this *IStreamSocket) Vtbl() *IStreamSocketVtbl {
	return (*IStreamSocketVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStreamSocket) Get_Control() *IStreamSocketControl {
	var _result *IStreamSocketControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Control, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStreamSocket) Get_Information() *IStreamSocketInformation {
	var _result *IStreamSocketInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Information, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStreamSocket) Get_InputStream() *IInputStream {
	var _result *IInputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InputStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStreamSocket) Get_OutputStream() *IOutputStream {
	var _result *IOutputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutputStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStreamSocket) ConnectWithEndpointPairAsync(endpointPair *IEndpointPair) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConnectWithEndpointPairAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(endpointPair)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStreamSocket) ConnectAsync(remoteHostName *IHostName, remoteServiceName string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConnectAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(remoteHostName)), NewHStr(remoteServiceName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStreamSocket) ConnectWithEndpointPairAndProtectionLevelAsync(endpointPair *IEndpointPair, protectionLevel SocketProtectionLevel) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConnectWithEndpointPairAndProtectionLevelAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(endpointPair)), uintptr(protectionLevel), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStreamSocket) ConnectWithProtectionLevelAsync(remoteHostName *IHostName, remoteServiceName string, protectionLevel SocketProtectionLevel) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConnectWithProtectionLevelAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(remoteHostName)), NewHStr(remoteServiceName).Ptr, uintptr(protectionLevel), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStreamSocket) UpgradeToSslAsync(protectionLevel SocketProtectionLevel, validationHostName *IHostName) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UpgradeToSslAsync, uintptr(unsafe.Pointer(this)), uintptr(protectionLevel), uintptr(unsafe.Pointer(validationHostName)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 29D0E575-F314-4D09-ADF0-0FBD967FBD9F
var IID_IStreamSocket2 = syscall.GUID{0x29D0E575, 0xF314, 0x4D09,
	[8]byte{0xAD, 0xF0, 0x0F, 0xBD, 0x96, 0x7F, 0xBD, 0x9F}}

type IStreamSocket2Interface interface {
	win32.IInspectableInterface
	ConnectWithProtectionLevelAndAdapterAsync(remoteHostName *IHostName, remoteServiceName string, protectionLevel SocketProtectionLevel, adapter *INetworkAdapter) *IAsyncAction
}

type IStreamSocket2Vtbl struct {
	win32.IInspectableVtbl
	ConnectWithProtectionLevelAndAdapterAsync uintptr
}

type IStreamSocket2 struct {
	win32.IInspectable
}

func (this *IStreamSocket2) Vtbl() *IStreamSocket2Vtbl {
	return (*IStreamSocket2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStreamSocket2) ConnectWithProtectionLevelAndAdapterAsync(remoteHostName *IHostName, remoteServiceName string, protectionLevel SocketProtectionLevel, adapter *INetworkAdapter) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConnectWithProtectionLevelAndAdapterAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(remoteHostName)), NewHStr(remoteServiceName).Ptr, uintptr(protectionLevel), uintptr(unsafe.Pointer(adapter)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3F430B00-9D28-4854-BAC3-2301941EC223
var IID_IStreamSocket3 = syscall.GUID{0x3F430B00, 0x9D28, 0x4854,
	[8]byte{0xBA, 0xC3, 0x23, 0x01, 0x94, 0x1E, 0xC2, 0x23}}

type IStreamSocket3Interface interface {
	win32.IInspectableInterface
	CancelIOAsync() *IAsyncAction
	EnableTransferOwnership(taskId syscall.GUID)
	EnableTransferOwnershipWithConnectedStandbyAction(taskId syscall.GUID, connectedStandbyAction SocketActivityConnectedStandbyAction)
	TransferOwnership(socketId string)
	TransferOwnershipWithContext(socketId string, data *ISocketActivityContext)
	TransferOwnershipWithContextAndKeepAliveTime(socketId string, data *ISocketActivityContext, keepAliveTime TimeSpan)
}

type IStreamSocket3Vtbl struct {
	win32.IInspectableVtbl
	CancelIOAsync                                     uintptr
	EnableTransferOwnership                           uintptr
	EnableTransferOwnershipWithConnectedStandbyAction uintptr
	TransferOwnership                                 uintptr
	TransferOwnershipWithContext                      uintptr
	TransferOwnershipWithContextAndKeepAliveTime      uintptr
}

type IStreamSocket3 struct {
	win32.IInspectable
}

func (this *IStreamSocket3) Vtbl() *IStreamSocket3Vtbl {
	return (*IStreamSocket3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStreamSocket3) CancelIOAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CancelIOAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStreamSocket3) EnableTransferOwnership(taskId syscall.GUID) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnableTransferOwnership, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&taskId)))
	_ = _hr
}

func (this *IStreamSocket3) EnableTransferOwnershipWithConnectedStandbyAction(taskId syscall.GUID, connectedStandbyAction SocketActivityConnectedStandbyAction) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnableTransferOwnershipWithConnectedStandbyAction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&taskId)), uintptr(connectedStandbyAction))
	_ = _hr
}

func (this *IStreamSocket3) TransferOwnership(socketId string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TransferOwnership, uintptr(unsafe.Pointer(this)), NewHStr(socketId).Ptr)
	_ = _hr
}

func (this *IStreamSocket3) TransferOwnershipWithContext(socketId string, data *ISocketActivityContext) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TransferOwnershipWithContext, uintptr(unsafe.Pointer(this)), NewHStr(socketId).Ptr, uintptr(unsafe.Pointer(data)))
	_ = _hr
}

func (this *IStreamSocket3) TransferOwnershipWithContextAndKeepAliveTime(socketId string, data *ISocketActivityContext, keepAliveTime TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TransferOwnershipWithContextAndKeepAliveTime, uintptr(unsafe.Pointer(this)), NewHStr(socketId).Ptr, uintptr(unsafe.Pointer(data)), *(*uintptr)(unsafe.Pointer(&keepAliveTime)))
	_ = _hr
}

// FE25ADF1-92AB-4AF3-9992-0F4C85E36CC4
var IID_IStreamSocketControl = syscall.GUID{0xFE25ADF1, 0x92AB, 0x4AF3,
	[8]byte{0x99, 0x92, 0x0F, 0x4C, 0x85, 0xE3, 0x6C, 0xC4}}

type IStreamSocketControlInterface interface {
	win32.IInspectableInterface
	Get_NoDelay() bool
	Put_NoDelay(value bool)
	Get_KeepAlive() bool
	Put_KeepAlive(value bool)
	Get_OutboundBufferSizeInBytes() uint32
	Put_OutboundBufferSizeInBytes(value uint32)
	Get_QualityOfService() SocketQualityOfService
	Put_QualityOfService(value SocketQualityOfService)
	Get_OutboundUnicastHopLimit() byte
	Put_OutboundUnicastHopLimit(value byte)
}

type IStreamSocketControlVtbl struct {
	win32.IInspectableVtbl
	Get_NoDelay                   uintptr
	Put_NoDelay                   uintptr
	Get_KeepAlive                 uintptr
	Put_KeepAlive                 uintptr
	Get_OutboundBufferSizeInBytes uintptr
	Put_OutboundBufferSizeInBytes uintptr
	Get_QualityOfService          uintptr
	Put_QualityOfService          uintptr
	Get_OutboundUnicastHopLimit   uintptr
	Put_OutboundUnicastHopLimit   uintptr
}

type IStreamSocketControl struct {
	win32.IInspectable
}

func (this *IStreamSocketControl) Vtbl() *IStreamSocketControlVtbl {
	return (*IStreamSocketControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStreamSocketControl) Get_NoDelay() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NoDelay, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStreamSocketControl) Put_NoDelay(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_NoDelay, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IStreamSocketControl) Get_KeepAlive() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KeepAlive, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStreamSocketControl) Put_KeepAlive(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_KeepAlive, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IStreamSocketControl) Get_OutboundBufferSizeInBytes() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutboundBufferSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStreamSocketControl) Put_OutboundBufferSizeInBytes(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_OutboundBufferSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IStreamSocketControl) Get_QualityOfService() SocketQualityOfService {
	var _result SocketQualityOfService
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_QualityOfService, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStreamSocketControl) Put_QualityOfService(value SocketQualityOfService) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_QualityOfService, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IStreamSocketControl) Get_OutboundUnicastHopLimit() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutboundUnicastHopLimit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStreamSocketControl) Put_OutboundUnicastHopLimit(value byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_OutboundUnicastHopLimit, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// C2D09A56-060F-44C1-B8E2-1FBF60BD62C5
var IID_IStreamSocketControl2 = syscall.GUID{0xC2D09A56, 0x060F, 0x44C1,
	[8]byte{0xB8, 0xE2, 0x1F, 0xBF, 0x60, 0xBD, 0x62, 0xC5}}

type IStreamSocketControl2Interface interface {
	win32.IInspectableInterface
	Get_IgnorableServerCertificateErrors() *IVector[ChainValidationResult]
}

type IStreamSocketControl2Vtbl struct {
	win32.IInspectableVtbl
	Get_IgnorableServerCertificateErrors uintptr
}

type IStreamSocketControl2 struct {
	win32.IInspectable
}

func (this *IStreamSocketControl2) Vtbl() *IStreamSocketControl2Vtbl {
	return (*IStreamSocketControl2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStreamSocketControl2) Get_IgnorableServerCertificateErrors() *IVector[ChainValidationResult] {
	var _result *IVector[ChainValidationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IgnorableServerCertificateErrors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C56A444C-4E74-403E-894C-B31CAE5C7342
var IID_IStreamSocketControl3 = syscall.GUID{0xC56A444C, 0x4E74, 0x403E,
	[8]byte{0x89, 0x4C, 0xB3, 0x1C, 0xAE, 0x5C, 0x73, 0x42}}

type IStreamSocketControl3Interface interface {
	win32.IInspectableInterface
	Get_SerializeConnectionAttempts() bool
	Put_SerializeConnectionAttempts(value bool)
	Get_ClientCertificate() *ICertificate
	Put_ClientCertificate(value *ICertificate)
}

type IStreamSocketControl3Vtbl struct {
	win32.IInspectableVtbl
	Get_SerializeConnectionAttempts uintptr
	Put_SerializeConnectionAttempts uintptr
	Get_ClientCertificate           uintptr
	Put_ClientCertificate           uintptr
}

type IStreamSocketControl3 struct {
	win32.IInspectable
}

func (this *IStreamSocketControl3) Vtbl() *IStreamSocketControl3Vtbl {
	return (*IStreamSocketControl3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStreamSocketControl3) Get_SerializeConnectionAttempts() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SerializeConnectionAttempts, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStreamSocketControl3) Put_SerializeConnectionAttempts(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SerializeConnectionAttempts, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IStreamSocketControl3) Get_ClientCertificate() *ICertificate {
	var _result *ICertificate
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ClientCertificate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStreamSocketControl3) Put_ClientCertificate(value *ICertificate) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ClientCertificate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 964E2B3D-EC27-4888-B3CE-C74B418423AD
var IID_IStreamSocketControl4 = syscall.GUID{0x964E2B3D, 0xEC27, 0x4888,
	[8]byte{0xB3, 0xCE, 0xC7, 0x4B, 0x41, 0x84, 0x23, 0xAD}}

type IStreamSocketControl4Interface interface {
	win32.IInspectableInterface
	Get_MinProtectionLevel() SocketProtectionLevel
	Put_MinProtectionLevel(value SocketProtectionLevel)
}

type IStreamSocketControl4Vtbl struct {
	win32.IInspectableVtbl
	Get_MinProtectionLevel uintptr
	Put_MinProtectionLevel uintptr
}

type IStreamSocketControl4 struct {
	win32.IInspectable
}

func (this *IStreamSocketControl4) Vtbl() *IStreamSocketControl4Vtbl {
	return (*IStreamSocketControl4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStreamSocketControl4) Get_MinProtectionLevel() SocketProtectionLevel {
	var _result SocketProtectionLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStreamSocketControl4) Put_MinProtectionLevel(value SocketProtectionLevel) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MinProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 3B80AE30-5E68-4205-88F0-DC85D2E25DED
var IID_IStreamSocketInformation = syscall.GUID{0x3B80AE30, 0x5E68, 0x4205,
	[8]byte{0x88, 0xF0, 0xDC, 0x85, 0xD2, 0xE2, 0x5D, 0xED}}

type IStreamSocketInformationInterface interface {
	win32.IInspectableInterface
	Get_LocalAddress() *IHostName
	Get_LocalPort() string
	Get_RemoteHostName() *IHostName
	Get_RemoteAddress() *IHostName
	Get_RemoteServiceName() string
	Get_RemotePort() string
	Get_RoundTripTimeStatistics() RoundTripTimeStatistics
	Get_BandwidthStatistics() BandwidthStatistics
	Get_ProtectionLevel() SocketProtectionLevel
	Get_SessionKey() *IBuffer
}

type IStreamSocketInformationVtbl struct {
	win32.IInspectableVtbl
	Get_LocalAddress            uintptr
	Get_LocalPort               uintptr
	Get_RemoteHostName          uintptr
	Get_RemoteAddress           uintptr
	Get_RemoteServiceName       uintptr
	Get_RemotePort              uintptr
	Get_RoundTripTimeStatistics uintptr
	Get_BandwidthStatistics     uintptr
	Get_ProtectionLevel         uintptr
	Get_SessionKey              uintptr
}

type IStreamSocketInformation struct {
	win32.IInspectable
}

func (this *IStreamSocketInformation) Vtbl() *IStreamSocketInformationVtbl {
	return (*IStreamSocketInformationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStreamSocketInformation) Get_LocalAddress() *IHostName {
	var _result *IHostName
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocalAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStreamSocketInformation) Get_LocalPort() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocalPort, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStreamSocketInformation) Get_RemoteHostName() *IHostName {
	var _result *IHostName
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RemoteHostName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStreamSocketInformation) Get_RemoteAddress() *IHostName {
	var _result *IHostName
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RemoteAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStreamSocketInformation) Get_RemoteServiceName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RemoteServiceName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStreamSocketInformation) Get_RemotePort() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RemotePort, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStreamSocketInformation) Get_RoundTripTimeStatistics() RoundTripTimeStatistics {
	var _result RoundTripTimeStatistics
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RoundTripTimeStatistics, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStreamSocketInformation) Get_BandwidthStatistics() BandwidthStatistics {
	var _result BandwidthStatistics
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BandwidthStatistics, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStreamSocketInformation) Get_ProtectionLevel() SocketProtectionLevel {
	var _result SocketProtectionLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStreamSocketInformation) Get_SessionKey() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SessionKey, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 12C28452-4BDC-4EE4-976A-CF130E9D92E3
var IID_IStreamSocketInformation2 = syscall.GUID{0x12C28452, 0x4BDC, 0x4EE4,
	[8]byte{0x97, 0x6A, 0xCF, 0x13, 0x0E, 0x9D, 0x92, 0xE3}}

type IStreamSocketInformation2Interface interface {
	win32.IInspectableInterface
	Get_ServerCertificateErrorSeverity() SocketSslErrorSeverity
	Get_ServerCertificateErrors() *IVectorView[ChainValidationResult]
	Get_ServerCertificate() *ICertificate
	Get_ServerIntermediateCertificates() *IVectorView[*ICertificate]
}

type IStreamSocketInformation2Vtbl struct {
	win32.IInspectableVtbl
	Get_ServerCertificateErrorSeverity uintptr
	Get_ServerCertificateErrors        uintptr
	Get_ServerCertificate              uintptr
	Get_ServerIntermediateCertificates uintptr
}

type IStreamSocketInformation2 struct {
	win32.IInspectable
}

func (this *IStreamSocketInformation2) Vtbl() *IStreamSocketInformation2Vtbl {
	return (*IStreamSocketInformation2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStreamSocketInformation2) Get_ServerCertificateErrorSeverity() SocketSslErrorSeverity {
	var _result SocketSslErrorSeverity
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServerCertificateErrorSeverity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStreamSocketInformation2) Get_ServerCertificateErrors() *IVectorView[ChainValidationResult] {
	var _result *IVectorView[ChainValidationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServerCertificateErrors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStreamSocketInformation2) Get_ServerCertificate() *ICertificate {
	var _result *ICertificate
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServerCertificate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStreamSocketInformation2) Get_ServerIntermediateCertificates() *IVectorView[*ICertificate] {
	var _result *IVectorView[*ICertificate]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServerIntermediateCertificates, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FF513437-DF9F-4DF0-BF82-0EC5D7B35AAE
var IID_IStreamSocketListener = syscall.GUID{0xFF513437, 0xDF9F, 0x4DF0,
	[8]byte{0xBF, 0x82, 0x0E, 0xC5, 0xD7, 0xB3, 0x5A, 0xAE}}

type IStreamSocketListenerInterface interface {
	win32.IInspectableInterface
	Get_Control() *IStreamSocketListenerControl
	Get_Information() *IStreamSocketListenerInformation
	BindServiceNameAsync(localServiceName string) *IAsyncAction
	BindEndpointAsync(localHostName *IHostName, localServiceName string) *IAsyncAction
	Add_ConnectionReceived(eventHandler TypedEventHandler[*IStreamSocketListener, *IStreamSocketListenerConnectionReceivedEventArgs]) EventRegistrationToken
	Remove_ConnectionReceived(eventCookie EventRegistrationToken)
}

type IStreamSocketListenerVtbl struct {
	win32.IInspectableVtbl
	Get_Control               uintptr
	Get_Information           uintptr
	BindServiceNameAsync      uintptr
	BindEndpointAsync         uintptr
	Add_ConnectionReceived    uintptr
	Remove_ConnectionReceived uintptr
}

type IStreamSocketListener struct {
	win32.IInspectable
}

func (this *IStreamSocketListener) Vtbl() *IStreamSocketListenerVtbl {
	return (*IStreamSocketListenerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStreamSocketListener) Get_Control() *IStreamSocketListenerControl {
	var _result *IStreamSocketListenerControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Control, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStreamSocketListener) Get_Information() *IStreamSocketListenerInformation {
	var _result *IStreamSocketListenerInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Information, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStreamSocketListener) BindServiceNameAsync(localServiceName string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().BindServiceNameAsync, uintptr(unsafe.Pointer(this)), NewHStr(localServiceName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStreamSocketListener) BindEndpointAsync(localHostName *IHostName, localServiceName string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().BindEndpointAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(localHostName)), NewHStr(localServiceName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStreamSocketListener) Add_ConnectionReceived(eventHandler TypedEventHandler[*IStreamSocketListener, *IStreamSocketListenerConnectionReceivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ConnectionReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(eventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStreamSocketListener) Remove_ConnectionReceived(eventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ConnectionReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&eventCookie)))
	_ = _hr
}

// 658DC13E-BB3E-4458-B232-ED1088694B98
var IID_IStreamSocketListener2 = syscall.GUID{0x658DC13E, 0xBB3E, 0x4458,
	[8]byte{0xB2, 0x32, 0xED, 0x10, 0x88, 0x69, 0x4B, 0x98}}

type IStreamSocketListener2Interface interface {
	win32.IInspectableInterface
	BindServiceNameWithProtectionLevelAsync(localServiceName string, protectionLevel SocketProtectionLevel) *IAsyncAction
	BindServiceNameWithProtectionLevelAndAdapterAsync(localServiceName string, protectionLevel SocketProtectionLevel, adapter *INetworkAdapter) *IAsyncAction
}

type IStreamSocketListener2Vtbl struct {
	win32.IInspectableVtbl
	BindServiceNameWithProtectionLevelAsync           uintptr
	BindServiceNameWithProtectionLevelAndAdapterAsync uintptr
}

type IStreamSocketListener2 struct {
	win32.IInspectable
}

func (this *IStreamSocketListener2) Vtbl() *IStreamSocketListener2Vtbl {
	return (*IStreamSocketListener2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStreamSocketListener2) BindServiceNameWithProtectionLevelAsync(localServiceName string, protectionLevel SocketProtectionLevel) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().BindServiceNameWithProtectionLevelAsync, uintptr(unsafe.Pointer(this)), NewHStr(localServiceName).Ptr, uintptr(protectionLevel), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStreamSocketListener2) BindServiceNameWithProtectionLevelAndAdapterAsync(localServiceName string, protectionLevel SocketProtectionLevel, adapter *INetworkAdapter) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().BindServiceNameWithProtectionLevelAndAdapterAsync, uintptr(unsafe.Pointer(this)), NewHStr(localServiceName).Ptr, uintptr(protectionLevel), uintptr(unsafe.Pointer(adapter)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4798201C-BDF8-4919-8542-28D450E74507
var IID_IStreamSocketListener3 = syscall.GUID{0x4798201C, 0xBDF8, 0x4919,
	[8]byte{0x85, 0x42, 0x28, 0xD4, 0x50, 0xE7, 0x45, 0x07}}

type IStreamSocketListener3Interface interface {
	win32.IInspectableInterface
	CancelIOAsync() *IAsyncAction
	EnableTransferOwnership(taskId syscall.GUID)
	EnableTransferOwnershipWithConnectedStandbyAction(taskId syscall.GUID, connectedStandbyAction SocketActivityConnectedStandbyAction)
	TransferOwnership(socketId string)
	TransferOwnershipWithContext(socketId string, data *ISocketActivityContext)
}

type IStreamSocketListener3Vtbl struct {
	win32.IInspectableVtbl
	CancelIOAsync                                     uintptr
	EnableTransferOwnership                           uintptr
	EnableTransferOwnershipWithConnectedStandbyAction uintptr
	TransferOwnership                                 uintptr
	TransferOwnershipWithContext                      uintptr
}

type IStreamSocketListener3 struct {
	win32.IInspectable
}

func (this *IStreamSocketListener3) Vtbl() *IStreamSocketListener3Vtbl {
	return (*IStreamSocketListener3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStreamSocketListener3) CancelIOAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CancelIOAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStreamSocketListener3) EnableTransferOwnership(taskId syscall.GUID) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnableTransferOwnership, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&taskId)))
	_ = _hr
}

func (this *IStreamSocketListener3) EnableTransferOwnershipWithConnectedStandbyAction(taskId syscall.GUID, connectedStandbyAction SocketActivityConnectedStandbyAction) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EnableTransferOwnershipWithConnectedStandbyAction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&taskId)), uintptr(connectedStandbyAction))
	_ = _hr
}

func (this *IStreamSocketListener3) TransferOwnership(socketId string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TransferOwnership, uintptr(unsafe.Pointer(this)), NewHStr(socketId).Ptr)
	_ = _hr
}

func (this *IStreamSocketListener3) TransferOwnershipWithContext(socketId string, data *ISocketActivityContext) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TransferOwnershipWithContext, uintptr(unsafe.Pointer(this)), NewHStr(socketId).Ptr, uintptr(unsafe.Pointer(data)))
	_ = _hr
}

// 0C472EA9-373F-447B-85B1-DDD4548803BA
var IID_IStreamSocketListenerConnectionReceivedEventArgs = syscall.GUID{0x0C472EA9, 0x373F, 0x447B,
	[8]byte{0x85, 0xB1, 0xDD, 0xD4, 0x54, 0x88, 0x03, 0xBA}}

type IStreamSocketListenerConnectionReceivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Socket() *IStreamSocket
}

type IStreamSocketListenerConnectionReceivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Socket uintptr
}

type IStreamSocketListenerConnectionReceivedEventArgs struct {
	win32.IInspectable
}

func (this *IStreamSocketListenerConnectionReceivedEventArgs) Vtbl() *IStreamSocketListenerConnectionReceivedEventArgsVtbl {
	return (*IStreamSocketListenerConnectionReceivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStreamSocketListenerConnectionReceivedEventArgs) Get_Socket() *IStreamSocket {
	var _result *IStreamSocket
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Socket, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 20D8C576-8D8A-4DBA-9722-A16C4D984980
var IID_IStreamSocketListenerControl = syscall.GUID{0x20D8C576, 0x8D8A, 0x4DBA,
	[8]byte{0x97, 0x22, 0xA1, 0x6C, 0x4D, 0x98, 0x49, 0x80}}

type IStreamSocketListenerControlInterface interface {
	win32.IInspectableInterface
	Get_QualityOfService() SocketQualityOfService
	Put_QualityOfService(value SocketQualityOfService)
}

type IStreamSocketListenerControlVtbl struct {
	win32.IInspectableVtbl
	Get_QualityOfService uintptr
	Put_QualityOfService uintptr
}

type IStreamSocketListenerControl struct {
	win32.IInspectable
}

func (this *IStreamSocketListenerControl) Vtbl() *IStreamSocketListenerControlVtbl {
	return (*IStreamSocketListenerControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStreamSocketListenerControl) Get_QualityOfService() SocketQualityOfService {
	var _result SocketQualityOfService
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_QualityOfService, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStreamSocketListenerControl) Put_QualityOfService(value SocketQualityOfService) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_QualityOfService, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 948BB665-2C3E-404B-B8B0-8EB249A2B0A1
var IID_IStreamSocketListenerControl2 = syscall.GUID{0x948BB665, 0x2C3E, 0x404B,
	[8]byte{0xB8, 0xB0, 0x8E, 0xB2, 0x49, 0xA2, 0xB0, 0xA1}}

type IStreamSocketListenerControl2Interface interface {
	win32.IInspectableInterface
	Get_NoDelay() bool
	Put_NoDelay(value bool)
	Get_KeepAlive() bool
	Put_KeepAlive(value bool)
	Get_OutboundBufferSizeInBytes() uint32
	Put_OutboundBufferSizeInBytes(value uint32)
	Get_OutboundUnicastHopLimit() byte
	Put_OutboundUnicastHopLimit(value byte)
}

type IStreamSocketListenerControl2Vtbl struct {
	win32.IInspectableVtbl
	Get_NoDelay                   uintptr
	Put_NoDelay                   uintptr
	Get_KeepAlive                 uintptr
	Put_KeepAlive                 uintptr
	Get_OutboundBufferSizeInBytes uintptr
	Put_OutboundBufferSizeInBytes uintptr
	Get_OutboundUnicastHopLimit   uintptr
	Put_OutboundUnicastHopLimit   uintptr
}

type IStreamSocketListenerControl2 struct {
	win32.IInspectable
}

func (this *IStreamSocketListenerControl2) Vtbl() *IStreamSocketListenerControl2Vtbl {
	return (*IStreamSocketListenerControl2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStreamSocketListenerControl2) Get_NoDelay() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NoDelay, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStreamSocketListenerControl2) Put_NoDelay(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_NoDelay, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IStreamSocketListenerControl2) Get_KeepAlive() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KeepAlive, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStreamSocketListenerControl2) Put_KeepAlive(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_KeepAlive, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IStreamSocketListenerControl2) Get_OutboundBufferSizeInBytes() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutboundBufferSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStreamSocketListenerControl2) Put_OutboundBufferSizeInBytes(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_OutboundBufferSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IStreamSocketListenerControl2) Get_OutboundUnicastHopLimit() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutboundUnicastHopLimit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStreamSocketListenerControl2) Put_OutboundUnicastHopLimit(value byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_OutboundUnicastHopLimit, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// E62BA82F-A63A-430B-BF62-29E93E5633B4
var IID_IStreamSocketListenerInformation = syscall.GUID{0xE62BA82F, 0xA63A, 0x430B,
	[8]byte{0xBF, 0x62, 0x29, 0xE9, 0x3E, 0x56, 0x33, 0xB4}}

type IStreamSocketListenerInformationInterface interface {
	win32.IInspectableInterface
	Get_LocalPort() string
}

type IStreamSocketListenerInformationVtbl struct {
	win32.IInspectableVtbl
	Get_LocalPort uintptr
}

type IStreamSocketListenerInformation struct {
	win32.IInspectable
}

func (this *IStreamSocketListenerInformation) Vtbl() *IStreamSocketListenerInformationVtbl {
	return (*IStreamSocketListenerInformationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStreamSocketListenerInformation) Get_LocalPort() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocalPort, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// A420BC4A-6E2E-4AF5-B556-355AE0CD4F29
var IID_IStreamSocketStatics = syscall.GUID{0xA420BC4A, 0x6E2E, 0x4AF5,
	[8]byte{0xB5, 0x56, 0x35, 0x5A, 0xE0, 0xCD, 0x4F, 0x29}}

type IStreamSocketStaticsInterface interface {
	win32.IInspectableInterface
	GetEndpointPairsAsync(remoteHostName *IHostName, remoteServiceName string) *IAsyncOperation[*IVectorView[*IEndpointPair]]
	GetEndpointPairsWithSortOptionsAsync(remoteHostName *IHostName, remoteServiceName string, sortOptions HostNameSortOptions) *IAsyncOperation[*IVectorView[*IEndpointPair]]
}

type IStreamSocketStaticsVtbl struct {
	win32.IInspectableVtbl
	GetEndpointPairsAsync                uintptr
	GetEndpointPairsWithSortOptionsAsync uintptr
}

type IStreamSocketStatics struct {
	win32.IInspectable
}

func (this *IStreamSocketStatics) Vtbl() *IStreamSocketStaticsVtbl {
	return (*IStreamSocketStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStreamSocketStatics) GetEndpointPairsAsync(remoteHostName *IHostName, remoteServiceName string) *IAsyncOperation[*IVectorView[*IEndpointPair]] {
	var _result *IAsyncOperation[*IVectorView[*IEndpointPair]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetEndpointPairsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(remoteHostName)), NewHStr(remoteServiceName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStreamSocketStatics) GetEndpointPairsWithSortOptionsAsync(remoteHostName *IHostName, remoteServiceName string, sortOptions HostNameSortOptions) *IAsyncOperation[*IVectorView[*IEndpointPair]] {
	var _result *IAsyncOperation[*IVectorView[*IEndpointPair]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetEndpointPairsWithSortOptionsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(remoteHostName)), NewHStr(remoteServiceName).Ptr, uintptr(sortOptions), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BD4A49D8-B289-45BB-97EB-C7525205A843
var IID_IStreamWebSocket = syscall.GUID{0xBD4A49D8, 0xB289, 0x45BB,
	[8]byte{0x97, 0xEB, 0xC7, 0x52, 0x52, 0x05, 0xA8, 0x43}}

type IStreamWebSocketInterface interface {
	win32.IInspectableInterface
	Get_Control() *IStreamWebSocketControl
	Get_Information() *IWebSocketInformation
	Get_InputStream() *IInputStream
}

type IStreamWebSocketVtbl struct {
	win32.IInspectableVtbl
	Get_Control     uintptr
	Get_Information uintptr
	Get_InputStream uintptr
}

type IStreamWebSocket struct {
	win32.IInspectable
}

func (this *IStreamWebSocket) Vtbl() *IStreamWebSocketVtbl {
	return (*IStreamWebSocketVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStreamWebSocket) Get_Control() *IStreamWebSocketControl {
	var _result *IStreamWebSocketControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Control, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStreamWebSocket) Get_Information() *IWebSocketInformation {
	var _result *IWebSocketInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Information, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStreamWebSocket) Get_InputStream() *IInputStream {
	var _result *IInputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InputStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AA4D08CB-93F5-4678-8236-57CCE5417ED5
var IID_IStreamWebSocket2 = syscall.GUID{0xAA4D08CB, 0x93F5, 0x4678,
	[8]byte{0x82, 0x36, 0x57, 0xCC, 0xE5, 0x41, 0x7E, 0xD5}}

type IStreamWebSocket2Interface interface {
	win32.IInspectableInterface
	Add_ServerCustomValidationRequested(eventHandler TypedEventHandler[*IStreamWebSocket, *IWebSocketServerCustomValidationRequestedEventArgs]) EventRegistrationToken
	Remove_ServerCustomValidationRequested(eventCookie EventRegistrationToken)
}

type IStreamWebSocket2Vtbl struct {
	win32.IInspectableVtbl
	Add_ServerCustomValidationRequested    uintptr
	Remove_ServerCustomValidationRequested uintptr
}

type IStreamWebSocket2 struct {
	win32.IInspectable
}

func (this *IStreamWebSocket2) Vtbl() *IStreamWebSocket2Vtbl {
	return (*IStreamWebSocket2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStreamWebSocket2) Add_ServerCustomValidationRequested(eventHandler TypedEventHandler[*IStreamWebSocket, *IWebSocketServerCustomValidationRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ServerCustomValidationRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(eventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStreamWebSocket2) Remove_ServerCustomValidationRequested(eventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ServerCustomValidationRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&eventCookie)))
	_ = _hr
}

// B4F478B1-A45A-48DB-953A-645B7D964C07
var IID_IStreamWebSocketControl = syscall.GUID{0xB4F478B1, 0xA45A, 0x48DB,
	[8]byte{0x95, 0x3A, 0x64, 0x5B, 0x7D, 0x96, 0x4C, 0x07}}

type IStreamWebSocketControlInterface interface {
	win32.IInspectableInterface
	Get_NoDelay() bool
	Put_NoDelay(value bool)
}

type IStreamWebSocketControlVtbl struct {
	win32.IInspectableVtbl
	Get_NoDelay uintptr
	Put_NoDelay uintptr
}

type IStreamWebSocketControl struct {
	win32.IInspectable
}

func (this *IStreamWebSocketControl) Vtbl() *IStreamWebSocketControlVtbl {
	return (*IStreamWebSocketControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStreamWebSocketControl) Get_NoDelay() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NoDelay, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStreamWebSocketControl) Put_NoDelay(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_NoDelay, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 215D9F7E-FA58-40DA-9F11-A48DAFE95037
var IID_IStreamWebSocketControl2 = syscall.GUID{0x215D9F7E, 0xFA58, 0x40DA,
	[8]byte{0x9F, 0x11, 0xA4, 0x8D, 0xAF, 0xE9, 0x50, 0x37}}

type IStreamWebSocketControl2Interface interface {
	win32.IInspectableInterface
	Get_DesiredUnsolicitedPongInterval() TimeSpan
	Put_DesiredUnsolicitedPongInterval(value TimeSpan)
	Get_ActualUnsolicitedPongInterval() TimeSpan
	Get_ClientCertificate() *ICertificate
	Put_ClientCertificate(value *ICertificate)
}

type IStreamWebSocketControl2Vtbl struct {
	win32.IInspectableVtbl
	Get_DesiredUnsolicitedPongInterval uintptr
	Put_DesiredUnsolicitedPongInterval uintptr
	Get_ActualUnsolicitedPongInterval  uintptr
	Get_ClientCertificate              uintptr
	Put_ClientCertificate              uintptr
}

type IStreamWebSocketControl2 struct {
	win32.IInspectable
}

func (this *IStreamWebSocketControl2) Vtbl() *IStreamWebSocketControl2Vtbl {
	return (*IStreamWebSocketControl2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStreamWebSocketControl2) Get_DesiredUnsolicitedPongInterval() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesiredUnsolicitedPongInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStreamWebSocketControl2) Put_DesiredUnsolicitedPongInterval(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DesiredUnsolicitedPongInterval, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IStreamWebSocketControl2) Get_ActualUnsolicitedPongInterval() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ActualUnsolicitedPongInterval, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStreamWebSocketControl2) Get_ClientCertificate() *ICertificate {
	var _result *ICertificate
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ClientCertificate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStreamWebSocketControl2) Put_ClientCertificate(value *ICertificate) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ClientCertificate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// F877396F-99B1-4E18-BC08-850C9ADF156E
var IID_IWebSocket = syscall.GUID{0xF877396F, 0x99B1, 0x4E18,
	[8]byte{0xBC, 0x08, 0x85, 0x0C, 0x9A, 0xDF, 0x15, 0x6E}}

type IWebSocketInterface interface {
	win32.IInspectableInterface
	Get_OutputStream() *IOutputStream
	ConnectAsync(uri *IUriRuntimeClass) *IAsyncAction
	SetRequestHeader(headerName string, headerValue string)
	Add_Closed(eventHandler TypedEventHandler[*IWebSocket, *IWebSocketClosedEventArgs]) EventRegistrationToken
	Remove_Closed(eventCookie EventRegistrationToken)
	CloseWithStatus(code uint16, reason string)
}

type IWebSocketVtbl struct {
	win32.IInspectableVtbl
	Get_OutputStream uintptr
	ConnectAsync     uintptr
	SetRequestHeader uintptr
	Add_Closed       uintptr
	Remove_Closed    uintptr
	CloseWithStatus  uintptr
}

type IWebSocket struct {
	win32.IInspectable
}

func (this *IWebSocket) Vtbl() *IWebSocketVtbl {
	return (*IWebSocketVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebSocket) Get_OutputStream() *IOutputStream {
	var _result *IOutputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutputStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebSocket) ConnectAsync(uri *IUriRuntimeClass) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConnectAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebSocket) SetRequestHeader(headerName string, headerValue string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetRequestHeader, uintptr(unsafe.Pointer(this)), NewHStr(headerName).Ptr, NewHStr(headerValue).Ptr)
	_ = _hr
}

func (this *IWebSocket) Add_Closed(eventHandler TypedEventHandler[*IWebSocket, *IWebSocketClosedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Closed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(eventHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWebSocket) Remove_Closed(eventCookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Closed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&eventCookie)))
	_ = _hr
}

func (this *IWebSocket) CloseWithStatus(code uint16, reason string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CloseWithStatus, uintptr(unsafe.Pointer(this)), uintptr(code), NewHStr(reason).Ptr)
	_ = _hr
}

// CEB78D07-D0A8-4703-A091-C8C2C0915BC3
var IID_IWebSocketClosedEventArgs = syscall.GUID{0xCEB78D07, 0xD0A8, 0x4703,
	[8]byte{0xA0, 0x91, 0xC8, 0xC2, 0xC0, 0x91, 0x5B, 0xC3}}

type IWebSocketClosedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Code() uint16
	Get_Reason() string
}

type IWebSocketClosedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Code   uintptr
	Get_Reason uintptr
}

type IWebSocketClosedEventArgs struct {
	win32.IInspectable
}

func (this *IWebSocketClosedEventArgs) Vtbl() *IWebSocketClosedEventArgsVtbl {
	return (*IWebSocketClosedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebSocketClosedEventArgs) Get_Code() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Code, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWebSocketClosedEventArgs) Get_Reason() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Reason, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 2EC4BDC3-D9A5-455A-9811-DE24D45337E9
var IID_IWebSocketControl = syscall.GUID{0x2EC4BDC3, 0xD9A5, 0x455A,
	[8]byte{0x98, 0x11, 0xDE, 0x24, 0xD4, 0x53, 0x37, 0xE9}}

type IWebSocketControlInterface interface {
	win32.IInspectableInterface
	Get_OutboundBufferSizeInBytes() uint32
	Put_OutboundBufferSizeInBytes(value uint32)
	Get_ServerCredential() *IPasswordCredential
	Put_ServerCredential(value *IPasswordCredential)
	Get_ProxyCredential() *IPasswordCredential
	Put_ProxyCredential(value *IPasswordCredential)
	Get_SupportedProtocols() *IVector[string]
}

type IWebSocketControlVtbl struct {
	win32.IInspectableVtbl
	Get_OutboundBufferSizeInBytes uintptr
	Put_OutboundBufferSizeInBytes uintptr
	Get_ServerCredential          uintptr
	Put_ServerCredential          uintptr
	Get_ProxyCredential           uintptr
	Put_ProxyCredential           uintptr
	Get_SupportedProtocols        uintptr
}

type IWebSocketControl struct {
	win32.IInspectable
}

func (this *IWebSocketControl) Vtbl() *IWebSocketControlVtbl {
	return (*IWebSocketControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebSocketControl) Get_OutboundBufferSizeInBytes() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutboundBufferSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWebSocketControl) Put_OutboundBufferSizeInBytes(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_OutboundBufferSizeInBytes, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IWebSocketControl) Get_ServerCredential() *IPasswordCredential {
	var _result *IPasswordCredential
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServerCredential, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebSocketControl) Put_ServerCredential(value *IPasswordCredential) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ServerCredential, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IWebSocketControl) Get_ProxyCredential() *IPasswordCredential {
	var _result *IPasswordCredential
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProxyCredential, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebSocketControl) Put_ProxyCredential(value *IPasswordCredential) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ProxyCredential, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IWebSocketControl) Get_SupportedProtocols() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedProtocols, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 79C3BE03-F2CA-461E-AF4E-9665BC2D0620
var IID_IWebSocketControl2 = syscall.GUID{0x79C3BE03, 0xF2CA, 0x461E,
	[8]byte{0xAF, 0x4E, 0x96, 0x65, 0xBC, 0x2D, 0x06, 0x20}}

type IWebSocketControl2Interface interface {
	win32.IInspectableInterface
	Get_IgnorableServerCertificateErrors() *IVector[ChainValidationResult]
}

type IWebSocketControl2Vtbl struct {
	win32.IInspectableVtbl
	Get_IgnorableServerCertificateErrors uintptr
}

type IWebSocketControl2 struct {
	win32.IInspectable
}

func (this *IWebSocketControl2) Vtbl() *IWebSocketControl2Vtbl {
	return (*IWebSocketControl2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebSocketControl2) Get_IgnorableServerCertificateErrors() *IVector[ChainValidationResult] {
	var _result *IVector[ChainValidationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IgnorableServerCertificateErrors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 27CDF35B-1F61-4709-8E02-61283ADA4E9D
var IID_IWebSocketErrorStatics = syscall.GUID{0x27CDF35B, 0x1F61, 0x4709,
	[8]byte{0x8E, 0x02, 0x61, 0x28, 0x3A, 0xDA, 0x4E, 0x9D}}

type IWebSocketErrorStaticsInterface interface {
	win32.IInspectableInterface
	GetStatus(hresult int32) WebErrorStatus
}

type IWebSocketErrorStaticsVtbl struct {
	win32.IInspectableVtbl
	GetStatus uintptr
}

type IWebSocketErrorStatics struct {
	win32.IInspectable
}

func (this *IWebSocketErrorStatics) Vtbl() *IWebSocketErrorStaticsVtbl {
	return (*IWebSocketErrorStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebSocketErrorStatics) GetStatus(hresult int32) WebErrorStatus {
	var _result WebErrorStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStatus, uintptr(unsafe.Pointer(this)), uintptr(hresult), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5E01E316-C92A-47A5-B25F-07847639D181
var IID_IWebSocketInformation = syscall.GUID{0x5E01E316, 0xC92A, 0x47A5,
	[8]byte{0xB2, 0x5F, 0x07, 0x84, 0x76, 0x39, 0xD1, 0x81}}

type IWebSocketInformationInterface interface {
	win32.IInspectableInterface
	Get_LocalAddress() *IHostName
	Get_BandwidthStatistics() BandwidthStatistics
	Get_Protocol() string
}

type IWebSocketInformationVtbl struct {
	win32.IInspectableVtbl
	Get_LocalAddress        uintptr
	Get_BandwidthStatistics uintptr
	Get_Protocol            uintptr
}

type IWebSocketInformation struct {
	win32.IInspectable
}

func (this *IWebSocketInformation) Vtbl() *IWebSocketInformationVtbl {
	return (*IWebSocketInformationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebSocketInformation) Get_LocalAddress() *IHostName {
	var _result *IHostName
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocalAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebSocketInformation) Get_BandwidthStatistics() BandwidthStatistics {
	var _result BandwidthStatistics
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BandwidthStatistics, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWebSocketInformation) Get_Protocol() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Protocol, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// CE1D39CE-A1B7-4D43-8269-8D5B981BD47A
var IID_IWebSocketInformation2 = syscall.GUID{0xCE1D39CE, 0xA1B7, 0x4D43,
	[8]byte{0x82, 0x69, 0x8D, 0x5B, 0x98, 0x1B, 0xD4, 0x7A}}

type IWebSocketInformation2Interface interface {
	win32.IInspectableInterface
	Get_ServerCertificate() *ICertificate
	Get_ServerCertificateErrorSeverity() SocketSslErrorSeverity
	Get_ServerCertificateErrors() *IVectorView[ChainValidationResult]
	Get_ServerIntermediateCertificates() *IVectorView[*ICertificate]
}

type IWebSocketInformation2Vtbl struct {
	win32.IInspectableVtbl
	Get_ServerCertificate              uintptr
	Get_ServerCertificateErrorSeverity uintptr
	Get_ServerCertificateErrors        uintptr
	Get_ServerIntermediateCertificates uintptr
}

type IWebSocketInformation2 struct {
	win32.IInspectable
}

func (this *IWebSocketInformation2) Vtbl() *IWebSocketInformation2Vtbl {
	return (*IWebSocketInformation2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebSocketInformation2) Get_ServerCertificate() *ICertificate {
	var _result *ICertificate
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServerCertificate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebSocketInformation2) Get_ServerCertificateErrorSeverity() SocketSslErrorSeverity {
	var _result SocketSslErrorSeverity
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServerCertificateErrorSeverity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWebSocketInformation2) Get_ServerCertificateErrors() *IVectorView[ChainValidationResult] {
	var _result *IVectorView[ChainValidationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServerCertificateErrors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebSocketInformation2) Get_ServerIntermediateCertificates() *IVectorView[*ICertificate] {
	var _result *IVectorView[*ICertificate]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServerIntermediateCertificates, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FFEFFE48-022A-4AB7-8B36-E10AF4640E6B
var IID_IWebSocketServerCustomValidationRequestedEventArgs = syscall.GUID{0xFFEFFE48, 0x022A, 0x4AB7,
	[8]byte{0x8B, 0x36, 0xE1, 0x0A, 0xF4, 0x64, 0x0E, 0x6B}}

type IWebSocketServerCustomValidationRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_ServerCertificate() *ICertificate
	Get_ServerCertificateErrorSeverity() SocketSslErrorSeverity
	Get_ServerCertificateErrors() *IVectorView[ChainValidationResult]
	Get_ServerIntermediateCertificates() *IVectorView[*ICertificate]
	Reject()
	GetDeferral() *IDeferral
}

type IWebSocketServerCustomValidationRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ServerCertificate              uintptr
	Get_ServerCertificateErrorSeverity uintptr
	Get_ServerCertificateErrors        uintptr
	Get_ServerIntermediateCertificates uintptr
	Reject                             uintptr
	GetDeferral                        uintptr
}

type IWebSocketServerCustomValidationRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IWebSocketServerCustomValidationRequestedEventArgs) Vtbl() *IWebSocketServerCustomValidationRequestedEventArgsVtbl {
	return (*IWebSocketServerCustomValidationRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebSocketServerCustomValidationRequestedEventArgs) Get_ServerCertificate() *ICertificate {
	var _result *ICertificate
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServerCertificate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebSocketServerCustomValidationRequestedEventArgs) Get_ServerCertificateErrorSeverity() SocketSslErrorSeverity {
	var _result SocketSslErrorSeverity
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServerCertificateErrorSeverity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWebSocketServerCustomValidationRequestedEventArgs) Get_ServerCertificateErrors() *IVectorView[ChainValidationResult] {
	var _result *IVectorView[ChainValidationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServerCertificateErrors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebSocketServerCustomValidationRequestedEventArgs) Get_ServerIntermediateCertificates() *IVectorView[*ICertificate] {
	var _result *IVectorView[*ICertificate]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServerIntermediateCertificates, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebSocketServerCustomValidationRequestedEventArgs) Reject() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Reject, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IWebSocketServerCustomValidationRequestedEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type DatagramSocket struct {
	RtClass
	*IDatagramSocket
}

func NewDatagramSocket() *DatagramSocket {
	hs := NewHStr("Windows.Networking.Sockets.DatagramSocket")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &DatagramSocket{
		RtClass:         RtClass{PInspect: p},
		IDatagramSocket: (*IDatagramSocket)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewIDatagramSocketStatics() *IDatagramSocketStatics {
	var p *IDatagramSocketStatics
	hs := NewHStr("Windows.Networking.Sockets.DatagramSocket")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IDatagramSocketStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type DatagramSocketControl struct {
	RtClass
	*IDatagramSocketControl
}

type DatagramSocketInformation struct {
	RtClass
	*IDatagramSocketInformation
}

type DatagramSocketMessageReceivedEventArgs struct {
	RtClass
	*IDatagramSocketMessageReceivedEventArgs
}

type MessageWebSocket struct {
	RtClass
	*IMessageWebSocket
}

func NewMessageWebSocket() *MessageWebSocket {
	hs := NewHStr("Windows.Networking.Sockets.MessageWebSocket")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &MessageWebSocket{
		RtClass:           RtClass{PInspect: p},
		IMessageWebSocket: (*IMessageWebSocket)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type MessageWebSocketControl struct {
	RtClass
	*IMessageWebSocketControl
}

type MessageWebSocketInformation struct {
	RtClass
	*IWebSocketInformation
}

type MessageWebSocketMessageReceivedEventArgs struct {
	RtClass
	*IMessageWebSocketMessageReceivedEventArgs
}

type SocketError struct {
	RtClass
}

func NewISocketErrorStatics() *ISocketErrorStatics {
	var p *ISocketErrorStatics
	hs := NewHStr("Windows.Networking.Sockets.SocketError")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISocketErrorStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type StreamSocket struct {
	RtClass
	*IStreamSocket
}

func NewStreamSocket() *StreamSocket {
	hs := NewHStr("Windows.Networking.Sockets.StreamSocket")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &StreamSocket{
		RtClass:       RtClass{PInspect: p},
		IStreamSocket: (*IStreamSocket)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewIStreamSocketStatics() *IStreamSocketStatics {
	var p *IStreamSocketStatics
	hs := NewHStr("Windows.Networking.Sockets.StreamSocket")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IStreamSocketStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type StreamSocketControl struct {
	RtClass
	*IStreamSocketControl
}

type StreamSocketInformation struct {
	RtClass
	*IStreamSocketInformation
}

type StreamSocketListener struct {
	RtClass
	*IStreamSocketListener
}

func NewStreamSocketListener() *StreamSocketListener {
	hs := NewHStr("Windows.Networking.Sockets.StreamSocketListener")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &StreamSocketListener{
		RtClass:               RtClass{PInspect: p},
		IStreamSocketListener: (*IStreamSocketListener)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type StreamSocketListenerConnectionReceivedEventArgs struct {
	RtClass
	*IStreamSocketListenerConnectionReceivedEventArgs
}

type StreamSocketListenerControl struct {
	RtClass
	*IStreamSocketListenerControl
}

type StreamSocketListenerInformation struct {
	RtClass
	*IStreamSocketListenerInformation
}

type StreamWebSocket struct {
	RtClass
	*IStreamWebSocket
}

func NewStreamWebSocket() *StreamWebSocket {
	hs := NewHStr("Windows.Networking.Sockets.StreamWebSocket")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &StreamWebSocket{
		RtClass:          RtClass{PInspect: p},
		IStreamWebSocket: (*IStreamWebSocket)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type StreamWebSocketControl struct {
	RtClass
	*IStreamWebSocketControl
}

type StreamWebSocketInformation struct {
	RtClass
	*IWebSocketInformation
}

type WebSocketClosedEventArgs struct {
	RtClass
	*IWebSocketClosedEventArgs
}

type WebSocketError struct {
	RtClass
}

func NewIWebSocketErrorStatics() *IWebSocketErrorStatics {
	var p *IWebSocketErrorStatics
	hs := NewHStr("Windows.Networking.Sockets.WebSocketError")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWebSocketErrorStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type WebSocketServerCustomValidationRequestedEventArgs struct {
	RtClass
	*IWebSocketServerCustomValidationRequestedEventArgs
}
