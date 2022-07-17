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
type XboxLiveEndpointPairCreationBehaviors uint32

const (
	XboxLiveEndpointPairCreationBehaviors_None           XboxLiveEndpointPairCreationBehaviors = 0
	XboxLiveEndpointPairCreationBehaviors_ReevaluatePath XboxLiveEndpointPairCreationBehaviors = 1
)

// enum
type XboxLiveEndpointPairCreationStatus int32

const (
	XboxLiveEndpointPairCreationStatus_Succeeded                 XboxLiveEndpointPairCreationStatus = 0
	XboxLiveEndpointPairCreationStatus_NoLocalNetworks           XboxLiveEndpointPairCreationStatus = 1
	XboxLiveEndpointPairCreationStatus_NoCompatibleNetworkPaths  XboxLiveEndpointPairCreationStatus = 2
	XboxLiveEndpointPairCreationStatus_LocalSystemNotAuthorized  XboxLiveEndpointPairCreationStatus = 3
	XboxLiveEndpointPairCreationStatus_Canceled                  XboxLiveEndpointPairCreationStatus = 4
	XboxLiveEndpointPairCreationStatus_TimedOut                  XboxLiveEndpointPairCreationStatus = 5
	XboxLiveEndpointPairCreationStatus_RemoteSystemNotAuthorized XboxLiveEndpointPairCreationStatus = 6
	XboxLiveEndpointPairCreationStatus_RefusedDueToConfiguration XboxLiveEndpointPairCreationStatus = 7
	XboxLiveEndpointPairCreationStatus_UnexpectedInternalError   XboxLiveEndpointPairCreationStatus = 8
)

// enum
type XboxLiveEndpointPairState int32

const (
	XboxLiveEndpointPairState_Invalid                   XboxLiveEndpointPairState = 0
	XboxLiveEndpointPairState_CreatingOutbound          XboxLiveEndpointPairState = 1
	XboxLiveEndpointPairState_CreatingInbound           XboxLiveEndpointPairState = 2
	XboxLiveEndpointPairState_Ready                     XboxLiveEndpointPairState = 3
	XboxLiveEndpointPairState_DeletingLocally           XboxLiveEndpointPairState = 4
	XboxLiveEndpointPairState_RemoteEndpointTerminating XboxLiveEndpointPairState = 5
	XboxLiveEndpointPairState_Deleted                   XboxLiveEndpointPairState = 6
)

// enum
type XboxLiveNetworkAccessKind int32

const (
	XboxLiveNetworkAccessKind_Open     XboxLiveNetworkAccessKind = 0
	XboxLiveNetworkAccessKind_Moderate XboxLiveNetworkAccessKind = 1
	XboxLiveNetworkAccessKind_Strict   XboxLiveNetworkAccessKind = 2
)

// enum
type XboxLiveQualityOfServiceMeasurementStatus int32

const (
	XboxLiveQualityOfServiceMeasurementStatus_NotStarted                       XboxLiveQualityOfServiceMeasurementStatus = 0
	XboxLiveQualityOfServiceMeasurementStatus_InProgress                       XboxLiveQualityOfServiceMeasurementStatus = 1
	XboxLiveQualityOfServiceMeasurementStatus_InProgressWithProvisionalResults XboxLiveQualityOfServiceMeasurementStatus = 2
	XboxLiveQualityOfServiceMeasurementStatus_Succeeded                        XboxLiveQualityOfServiceMeasurementStatus = 3
	XboxLiveQualityOfServiceMeasurementStatus_NoLocalNetworks                  XboxLiveQualityOfServiceMeasurementStatus = 4
	XboxLiveQualityOfServiceMeasurementStatus_NoCompatibleNetworkPaths         XboxLiveQualityOfServiceMeasurementStatus = 5
	XboxLiveQualityOfServiceMeasurementStatus_LocalSystemNotAuthorized         XboxLiveQualityOfServiceMeasurementStatus = 6
	XboxLiveQualityOfServiceMeasurementStatus_Canceled                         XboxLiveQualityOfServiceMeasurementStatus = 7
	XboxLiveQualityOfServiceMeasurementStatus_TimedOut                         XboxLiveQualityOfServiceMeasurementStatus = 8
	XboxLiveQualityOfServiceMeasurementStatus_RemoteSystemNotAuthorized        XboxLiveQualityOfServiceMeasurementStatus = 9
	XboxLiveQualityOfServiceMeasurementStatus_RefusedDueToConfiguration        XboxLiveQualityOfServiceMeasurementStatus = 10
	XboxLiveQualityOfServiceMeasurementStatus_UnexpectedInternalError          XboxLiveQualityOfServiceMeasurementStatus = 11
)

// enum
type XboxLiveQualityOfServiceMetric int32

const (
	XboxLiveQualityOfServiceMetric_AverageLatencyInMilliseconds XboxLiveQualityOfServiceMetric = 0
	XboxLiveQualityOfServiceMetric_MinLatencyInMilliseconds     XboxLiveQualityOfServiceMetric = 1
	XboxLiveQualityOfServiceMetric_MaxLatencyInMilliseconds     XboxLiveQualityOfServiceMetric = 2
	XboxLiveQualityOfServiceMetric_AverageOutboundBitsPerSecond XboxLiveQualityOfServiceMetric = 3
	XboxLiveQualityOfServiceMetric_MinOutboundBitsPerSecond     XboxLiveQualityOfServiceMetric = 4
	XboxLiveQualityOfServiceMetric_MaxOutboundBitsPerSecond     XboxLiveQualityOfServiceMetric = 5
	XboxLiveQualityOfServiceMetric_AverageInboundBitsPerSecond  XboxLiveQualityOfServiceMetric = 6
	XboxLiveQualityOfServiceMetric_MinInboundBitsPerSecond      XboxLiveQualityOfServiceMetric = 7
	XboxLiveQualityOfServiceMetric_MaxInboundBitsPerSecond      XboxLiveQualityOfServiceMetric = 8
)

// enum
type XboxLiveSocketKind int32

const (
	XboxLiveSocketKind_None     XboxLiveSocketKind = 0
	XboxLiveSocketKind_Datagram XboxLiveSocketKind = 1
	XboxLiveSocketKind_Stream   XboxLiveSocketKind = 2
)

// structs

type XboxLiveSecureSocketsContract struct {
}

// interfaces

// F5BBD279-3C86-4B57-A31A-B9462408FD01
var IID_IXboxLiveDeviceAddress = syscall.GUID{0xF5BBD279, 0x3C86, 0x4B57,
	[8]byte{0xA3, 0x1A, 0xB9, 0x46, 0x24, 0x08, 0xFD, 0x01}}

type IXboxLiveDeviceAddressInterface interface {
	win32.IInspectableInterface
	Add_SnapshotChanged(handler TypedEventHandler[*IXboxLiveDeviceAddress, interface{}]) EventRegistrationToken
	Remove_SnapshotChanged(token EventRegistrationToken)
	GetSnapshotAsBase64() string
	GetSnapshotAsBuffer() *IBuffer
	GetSnapshotAsBytes(bufferLength uint32, buffer *byte, bytesWritten uint32)
	Compare(otherDeviceAddress *IXboxLiveDeviceAddress) int32
	Get_IsValid() bool
	Get_IsLocal() bool
	Get_NetworkAccessKind() XboxLiveNetworkAccessKind
}

type IXboxLiveDeviceAddressVtbl struct {
	win32.IInspectableVtbl
	Add_SnapshotChanged    uintptr
	Remove_SnapshotChanged uintptr
	GetSnapshotAsBase64    uintptr
	GetSnapshotAsBuffer    uintptr
	GetSnapshotAsBytes     uintptr
	Compare                uintptr
	Get_IsValid            uintptr
	Get_IsLocal            uintptr
	Get_NetworkAccessKind  uintptr
}

type IXboxLiveDeviceAddress struct {
	win32.IInspectable
}

func (this *IXboxLiveDeviceAddress) Vtbl() *IXboxLiveDeviceAddressVtbl {
	return (*IXboxLiveDeviceAddressVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXboxLiveDeviceAddress) Add_SnapshotChanged(handler TypedEventHandler[*IXboxLiveDeviceAddress, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SnapshotChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IXboxLiveDeviceAddress) Remove_SnapshotChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SnapshotChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IXboxLiveDeviceAddress) GetSnapshotAsBase64() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSnapshotAsBase64, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IXboxLiveDeviceAddress) GetSnapshotAsBuffer() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSnapshotAsBuffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IXboxLiveDeviceAddress) GetSnapshotAsBytes(bufferLength uint32, buffer *byte, bytesWritten uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSnapshotAsBytes, uintptr(unsafe.Pointer(this)), uintptr(bufferLength), uintptr(unsafe.Pointer(buffer)), uintptr(bytesWritten))
	_ = _hr
}

func (this *IXboxLiveDeviceAddress) Compare(otherDeviceAddress *IXboxLiveDeviceAddress) int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Compare, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(otherDeviceAddress)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IXboxLiveDeviceAddress) Get_IsValid() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsValid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IXboxLiveDeviceAddress) Get_IsLocal() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsLocal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IXboxLiveDeviceAddress) Get_NetworkAccessKind() XboxLiveNetworkAccessKind {
	var _result XboxLiveNetworkAccessKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkAccessKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5954A819-4A79-4931-827C-7F503E963263
var IID_IXboxLiveDeviceAddressStatics = syscall.GUID{0x5954A819, 0x4A79, 0x4931,
	[8]byte{0x82, 0x7C, 0x7F, 0x50, 0x3E, 0x96, 0x32, 0x63}}

type IXboxLiveDeviceAddressStaticsInterface interface {
	win32.IInspectableInterface
	CreateFromSnapshotBase64(base64 string) *IXboxLiveDeviceAddress
	CreateFromSnapshotBuffer(buffer *IBuffer) *IXboxLiveDeviceAddress
	CreateFromSnapshotBytes(bufferLength uint32, buffer *byte) *IXboxLiveDeviceAddress
	GetLocal() *IXboxLiveDeviceAddress
	Get_MaxSnapshotBytesSize() uint32
}

type IXboxLiveDeviceAddressStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateFromSnapshotBase64 uintptr
	CreateFromSnapshotBuffer uintptr
	CreateFromSnapshotBytes  uintptr
	GetLocal                 uintptr
	Get_MaxSnapshotBytesSize uintptr
}

type IXboxLiveDeviceAddressStatics struct {
	win32.IInspectable
}

func (this *IXboxLiveDeviceAddressStatics) Vtbl() *IXboxLiveDeviceAddressStaticsVtbl {
	return (*IXboxLiveDeviceAddressStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXboxLiveDeviceAddressStatics) CreateFromSnapshotBase64(base64 string) *IXboxLiveDeviceAddress {
	var _result *IXboxLiveDeviceAddress
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromSnapshotBase64, uintptr(unsafe.Pointer(this)), NewHStr(base64).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IXboxLiveDeviceAddressStatics) CreateFromSnapshotBuffer(buffer *IBuffer) *IXboxLiveDeviceAddress {
	var _result *IXboxLiveDeviceAddress
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromSnapshotBuffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IXboxLiveDeviceAddressStatics) CreateFromSnapshotBytes(bufferLength uint32, buffer *byte) *IXboxLiveDeviceAddress {
	var _result *IXboxLiveDeviceAddress
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromSnapshotBytes, uintptr(unsafe.Pointer(this)), uintptr(bufferLength), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IXboxLiveDeviceAddressStatics) GetLocal() *IXboxLiveDeviceAddress {
	var _result *IXboxLiveDeviceAddress
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetLocal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IXboxLiveDeviceAddressStatics) Get_MaxSnapshotBytesSize() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxSnapshotBytesSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 1E9A839B-813E-44E0-B87F-C87A093475E4
var IID_IXboxLiveEndpointPair = syscall.GUID{0x1E9A839B, 0x813E, 0x44E0,
	[8]byte{0xB8, 0x7F, 0xC8, 0x7A, 0x09, 0x34, 0x75, 0xE4}}

type IXboxLiveEndpointPairInterface interface {
	win32.IInspectableInterface
	Add_StateChanged(handler TypedEventHandler[*IXboxLiveEndpointPair, *IXboxLiveEndpointPairStateChangedEventArgs]) EventRegistrationToken
	Remove_StateChanged(token EventRegistrationToken)
	DeleteAsync() *IAsyncAction
	GetRemoteSocketAddressBytes(socketAddressLength uint32, socketAddress *byte)
	GetLocalSocketAddressBytes(socketAddressLength uint32, socketAddress *byte)
	Get_State() XboxLiveEndpointPairState
	Get_Template() *IXboxLiveEndpointPairTemplate
	Get_RemoteDeviceAddress() *IXboxLiveDeviceAddress
	Get_RemoteHostName() *IHostName
	Get_RemotePort() string
	Get_LocalHostName() *IHostName
	Get_LocalPort() string
}

type IXboxLiveEndpointPairVtbl struct {
	win32.IInspectableVtbl
	Add_StateChanged            uintptr
	Remove_StateChanged         uintptr
	DeleteAsync                 uintptr
	GetRemoteSocketAddressBytes uintptr
	GetLocalSocketAddressBytes  uintptr
	Get_State                   uintptr
	Get_Template                uintptr
	Get_RemoteDeviceAddress     uintptr
	Get_RemoteHostName          uintptr
	Get_RemotePort              uintptr
	Get_LocalHostName           uintptr
	Get_LocalPort               uintptr
}

type IXboxLiveEndpointPair struct {
	win32.IInspectable
}

func (this *IXboxLiveEndpointPair) Vtbl() *IXboxLiveEndpointPairVtbl {
	return (*IXboxLiveEndpointPairVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXboxLiveEndpointPair) Add_StateChanged(handler TypedEventHandler[*IXboxLiveEndpointPair, *IXboxLiveEndpointPairStateChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IXboxLiveEndpointPair) Remove_StateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IXboxLiveEndpointPair) DeleteAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DeleteAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IXboxLiveEndpointPair) GetRemoteSocketAddressBytes(socketAddressLength uint32, socketAddress *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetRemoteSocketAddressBytes, uintptr(unsafe.Pointer(this)), uintptr(socketAddressLength), uintptr(unsafe.Pointer(socketAddress)))
	_ = _hr
}

func (this *IXboxLiveEndpointPair) GetLocalSocketAddressBytes(socketAddressLength uint32, socketAddress *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetLocalSocketAddressBytes, uintptr(unsafe.Pointer(this)), uintptr(socketAddressLength), uintptr(unsafe.Pointer(socketAddress)))
	_ = _hr
}

func (this *IXboxLiveEndpointPair) Get_State() XboxLiveEndpointPairState {
	var _result XboxLiveEndpointPairState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IXboxLiveEndpointPair) Get_Template() *IXboxLiveEndpointPairTemplate {
	var _result *IXboxLiveEndpointPairTemplate
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Template, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IXboxLiveEndpointPair) Get_RemoteDeviceAddress() *IXboxLiveDeviceAddress {
	var _result *IXboxLiveDeviceAddress
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RemoteDeviceAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IXboxLiveEndpointPair) Get_RemoteHostName() *IHostName {
	var _result *IHostName
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RemoteHostName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IXboxLiveEndpointPair) Get_RemotePort() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RemotePort, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IXboxLiveEndpointPair) Get_LocalHostName() *IHostName {
	var _result *IHostName
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocalHostName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IXboxLiveEndpointPair) Get_LocalPort() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LocalPort, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// D9A8BB95-2AAB-4D1E-9794-33ECC0DCF0FE
var IID_IXboxLiveEndpointPairCreationResult = syscall.GUID{0xD9A8BB95, 0x2AAB, 0x4D1E,
	[8]byte{0x97, 0x94, 0x33, 0xEC, 0xC0, 0xDC, 0xF0, 0xFE}}

type IXboxLiveEndpointPairCreationResultInterface interface {
	win32.IInspectableInterface
	Get_DeviceAddress() *IXboxLiveDeviceAddress
	Get_Status() XboxLiveEndpointPairCreationStatus
	Get_IsExistingPathEvaluation() bool
	Get_EndpointPair() *IXboxLiveEndpointPair
}

type IXboxLiveEndpointPairCreationResultVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceAddress            uintptr
	Get_Status                   uintptr
	Get_IsExistingPathEvaluation uintptr
	Get_EndpointPair             uintptr
}

type IXboxLiveEndpointPairCreationResult struct {
	win32.IInspectable
}

func (this *IXboxLiveEndpointPairCreationResult) Vtbl() *IXboxLiveEndpointPairCreationResultVtbl {
	return (*IXboxLiveEndpointPairCreationResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXboxLiveEndpointPairCreationResult) Get_DeviceAddress() *IXboxLiveDeviceAddress {
	var _result *IXboxLiveDeviceAddress
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IXboxLiveEndpointPairCreationResult) Get_Status() XboxLiveEndpointPairCreationStatus {
	var _result XboxLiveEndpointPairCreationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IXboxLiveEndpointPairCreationResult) Get_IsExistingPathEvaluation() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsExistingPathEvaluation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IXboxLiveEndpointPairCreationResult) Get_EndpointPair() *IXboxLiveEndpointPair {
	var _result *IXboxLiveEndpointPair
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EndpointPair, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 592E3B55-DE08-44E7-AC3B-B9B9A169583A
var IID_IXboxLiveEndpointPairStateChangedEventArgs = syscall.GUID{0x592E3B55, 0xDE08, 0x44E7,
	[8]byte{0xAC, 0x3B, 0xB9, 0xB9, 0xA1, 0x69, 0x58, 0x3A}}

type IXboxLiveEndpointPairStateChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_OldState() XboxLiveEndpointPairState
	Get_NewState() XboxLiveEndpointPairState
}

type IXboxLiveEndpointPairStateChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_OldState uintptr
	Get_NewState uintptr
}

type IXboxLiveEndpointPairStateChangedEventArgs struct {
	win32.IInspectable
}

func (this *IXboxLiveEndpointPairStateChangedEventArgs) Vtbl() *IXboxLiveEndpointPairStateChangedEventArgsVtbl {
	return (*IXboxLiveEndpointPairStateChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXboxLiveEndpointPairStateChangedEventArgs) Get_OldState() XboxLiveEndpointPairState {
	var _result XboxLiveEndpointPairState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OldState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IXboxLiveEndpointPairStateChangedEventArgs) Get_NewState() XboxLiveEndpointPairState {
	var _result XboxLiveEndpointPairState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NewState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 64316B30-217A-4243-8EE1-6729281D27DB
var IID_IXboxLiveEndpointPairStatics = syscall.GUID{0x64316B30, 0x217A, 0x4243,
	[8]byte{0x8E, 0xE1, 0x67, 0x29, 0x28, 0x1D, 0x27, 0xDB}}

type IXboxLiveEndpointPairStaticsInterface interface {
	win32.IInspectableInterface
	FindEndpointPairBySocketAddressBytes(localSocketAddressLength uint32, localSocketAddress *byte, remoteSocketAddressLength uint32, remoteSocketAddress *byte) *IXboxLiveEndpointPair
	FindEndpointPairByHostNamesAndPorts(localHostName *IHostName, localPort string, remoteHostName *IHostName, remotePort string) *IXboxLiveEndpointPair
}

type IXboxLiveEndpointPairStaticsVtbl struct {
	win32.IInspectableVtbl
	FindEndpointPairBySocketAddressBytes uintptr
	FindEndpointPairByHostNamesAndPorts  uintptr
}

type IXboxLiveEndpointPairStatics struct {
	win32.IInspectable
}

func (this *IXboxLiveEndpointPairStatics) Vtbl() *IXboxLiveEndpointPairStaticsVtbl {
	return (*IXboxLiveEndpointPairStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXboxLiveEndpointPairStatics) FindEndpointPairBySocketAddressBytes(localSocketAddressLength uint32, localSocketAddress *byte, remoteSocketAddressLength uint32, remoteSocketAddress *byte) *IXboxLiveEndpointPair {
	var _result *IXboxLiveEndpointPair
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindEndpointPairBySocketAddressBytes, uintptr(unsafe.Pointer(this)), uintptr(localSocketAddressLength), uintptr(unsafe.Pointer(localSocketAddress)), uintptr(remoteSocketAddressLength), uintptr(unsafe.Pointer(remoteSocketAddress)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IXboxLiveEndpointPairStatics) FindEndpointPairByHostNamesAndPorts(localHostName *IHostName, localPort string, remoteHostName *IHostName, remotePort string) *IXboxLiveEndpointPair {
	var _result *IXboxLiveEndpointPair
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindEndpointPairByHostNamesAndPorts, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(localHostName)), NewHStr(localPort).Ptr, uintptr(unsafe.Pointer(remoteHostName)), NewHStr(remotePort).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6B286ECF-3457-40CE-B9A1-C0CFE0213EA7
var IID_IXboxLiveEndpointPairTemplate = syscall.GUID{0x6B286ECF, 0x3457, 0x40CE,
	[8]byte{0xB9, 0xA1, 0xC0, 0xCF, 0xE0, 0x21, 0x3E, 0xA7}}

type IXboxLiveEndpointPairTemplateInterface interface {
	win32.IInspectableInterface
	Add_InboundEndpointPairCreated(handler TypedEventHandler[*IXboxLiveEndpointPairTemplate, *IXboxLiveInboundEndpointPairCreatedEventArgs]) EventRegistrationToken
	Remove_InboundEndpointPairCreated(token EventRegistrationToken)
	CreateEndpointPairDefaultAsync(deviceAddress *IXboxLiveDeviceAddress) *IAsyncOperation[*IXboxLiveEndpointPairCreationResult]
	CreateEndpointPairWithBehaviorsAsync(deviceAddress *IXboxLiveDeviceAddress, behaviors XboxLiveEndpointPairCreationBehaviors) *IAsyncOperation[*IXboxLiveEndpointPairCreationResult]
	CreateEndpointPairForPortsDefaultAsync(deviceAddress *IXboxLiveDeviceAddress, initiatorPort string, acceptorPort string) *IAsyncOperation[*IXboxLiveEndpointPairCreationResult]
	CreateEndpointPairForPortsWithBehaviorsAsync(deviceAddress *IXboxLiveDeviceAddress, initiatorPort string, acceptorPort string, behaviors XboxLiveEndpointPairCreationBehaviors) *IAsyncOperation[*IXboxLiveEndpointPairCreationResult]
	Get_Name() string
	Get_SocketKind() XboxLiveSocketKind
	Get_InitiatorBoundPortRangeLower() uint16
	Get_InitiatorBoundPortRangeUpper() uint16
	Get_AcceptorBoundPortRangeLower() uint16
	Get_AcceptorBoundPortRangeUpper() uint16
	Get_EndpointPairs() *IVectorView[*IXboxLiveEndpointPair]
}

type IXboxLiveEndpointPairTemplateVtbl struct {
	win32.IInspectableVtbl
	Add_InboundEndpointPairCreated               uintptr
	Remove_InboundEndpointPairCreated            uintptr
	CreateEndpointPairDefaultAsync               uintptr
	CreateEndpointPairWithBehaviorsAsync         uintptr
	CreateEndpointPairForPortsDefaultAsync       uintptr
	CreateEndpointPairForPortsWithBehaviorsAsync uintptr
	Get_Name                                     uintptr
	Get_SocketKind                               uintptr
	Get_InitiatorBoundPortRangeLower             uintptr
	Get_InitiatorBoundPortRangeUpper             uintptr
	Get_AcceptorBoundPortRangeLower              uintptr
	Get_AcceptorBoundPortRangeUpper              uintptr
	Get_EndpointPairs                            uintptr
}

type IXboxLiveEndpointPairTemplate struct {
	win32.IInspectable
}

func (this *IXboxLiveEndpointPairTemplate) Vtbl() *IXboxLiveEndpointPairTemplateVtbl {
	return (*IXboxLiveEndpointPairTemplateVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXboxLiveEndpointPairTemplate) Add_InboundEndpointPairCreated(handler TypedEventHandler[*IXboxLiveEndpointPairTemplate, *IXboxLiveInboundEndpointPairCreatedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_InboundEndpointPairCreated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IXboxLiveEndpointPairTemplate) Remove_InboundEndpointPairCreated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_InboundEndpointPairCreated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IXboxLiveEndpointPairTemplate) CreateEndpointPairDefaultAsync(deviceAddress *IXboxLiveDeviceAddress) *IAsyncOperation[*IXboxLiveEndpointPairCreationResult] {
	var _result *IAsyncOperation[*IXboxLiveEndpointPairCreationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateEndpointPairDefaultAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(deviceAddress)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IXboxLiveEndpointPairTemplate) CreateEndpointPairWithBehaviorsAsync(deviceAddress *IXboxLiveDeviceAddress, behaviors XboxLiveEndpointPairCreationBehaviors) *IAsyncOperation[*IXboxLiveEndpointPairCreationResult] {
	var _result *IAsyncOperation[*IXboxLiveEndpointPairCreationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateEndpointPairWithBehaviorsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(deviceAddress)), uintptr(behaviors), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IXboxLiveEndpointPairTemplate) CreateEndpointPairForPortsDefaultAsync(deviceAddress *IXboxLiveDeviceAddress, initiatorPort string, acceptorPort string) *IAsyncOperation[*IXboxLiveEndpointPairCreationResult] {
	var _result *IAsyncOperation[*IXboxLiveEndpointPairCreationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateEndpointPairForPortsDefaultAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(deviceAddress)), NewHStr(initiatorPort).Ptr, NewHStr(acceptorPort).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IXboxLiveEndpointPairTemplate) CreateEndpointPairForPortsWithBehaviorsAsync(deviceAddress *IXboxLiveDeviceAddress, initiatorPort string, acceptorPort string, behaviors XboxLiveEndpointPairCreationBehaviors) *IAsyncOperation[*IXboxLiveEndpointPairCreationResult] {
	var _result *IAsyncOperation[*IXboxLiveEndpointPairCreationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateEndpointPairForPortsWithBehaviorsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(deviceAddress)), NewHStr(initiatorPort).Ptr, NewHStr(acceptorPort).Ptr, uintptr(behaviors), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IXboxLiveEndpointPairTemplate) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IXboxLiveEndpointPairTemplate) Get_SocketKind() XboxLiveSocketKind {
	var _result XboxLiveSocketKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SocketKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IXboxLiveEndpointPairTemplate) Get_InitiatorBoundPortRangeLower() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InitiatorBoundPortRangeLower, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IXboxLiveEndpointPairTemplate) Get_InitiatorBoundPortRangeUpper() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InitiatorBoundPortRangeUpper, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IXboxLiveEndpointPairTemplate) Get_AcceptorBoundPortRangeLower() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AcceptorBoundPortRangeLower, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IXboxLiveEndpointPairTemplate) Get_AcceptorBoundPortRangeUpper() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AcceptorBoundPortRangeUpper, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IXboxLiveEndpointPairTemplate) Get_EndpointPairs() *IVectorView[*IXboxLiveEndpointPair] {
	var _result *IVectorView[*IXboxLiveEndpointPair]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EndpointPairs, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1E13137B-737B-4A23-BC64-0870F75655BA
var IID_IXboxLiveEndpointPairTemplateStatics = syscall.GUID{0x1E13137B, 0x737B, 0x4A23,
	[8]byte{0xBC, 0x64, 0x08, 0x70, 0xF7, 0x56, 0x55, 0xBA}}

type IXboxLiveEndpointPairTemplateStaticsInterface interface {
	win32.IInspectableInterface
	GetTemplateByName(name string) *IXboxLiveEndpointPairTemplate
	Get_Templates() *IVectorView[*IXboxLiveEndpointPairTemplate]
}

type IXboxLiveEndpointPairTemplateStaticsVtbl struct {
	win32.IInspectableVtbl
	GetTemplateByName uintptr
	Get_Templates     uintptr
}

type IXboxLiveEndpointPairTemplateStatics struct {
	win32.IInspectable
}

func (this *IXboxLiveEndpointPairTemplateStatics) Vtbl() *IXboxLiveEndpointPairTemplateStaticsVtbl {
	return (*IXboxLiveEndpointPairTemplateStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXboxLiveEndpointPairTemplateStatics) GetTemplateByName(name string) *IXboxLiveEndpointPairTemplate {
	var _result *IXboxLiveEndpointPairTemplate
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetTemplateByName, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IXboxLiveEndpointPairTemplateStatics) Get_Templates() *IVectorView[*IXboxLiveEndpointPairTemplate] {
	var _result *IVectorView[*IXboxLiveEndpointPairTemplate]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Templates, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DC183B62-22BA-48D2-80DE-C23968BD198B
var IID_IXboxLiveInboundEndpointPairCreatedEventArgs = syscall.GUID{0xDC183B62, 0x22BA, 0x48D2,
	[8]byte{0x80, 0xDE, 0xC2, 0x39, 0x68, 0xBD, 0x19, 0x8B}}

type IXboxLiveInboundEndpointPairCreatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_EndpointPair() *IXboxLiveEndpointPair
}

type IXboxLiveInboundEndpointPairCreatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_EndpointPair uintptr
}

type IXboxLiveInboundEndpointPairCreatedEventArgs struct {
	win32.IInspectable
}

func (this *IXboxLiveInboundEndpointPairCreatedEventArgs) Vtbl() *IXboxLiveInboundEndpointPairCreatedEventArgsVtbl {
	return (*IXboxLiveInboundEndpointPairCreatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXboxLiveInboundEndpointPairCreatedEventArgs) Get_EndpointPair() *IXboxLiveEndpointPair {
	var _result *IXboxLiveEndpointPair
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EndpointPair, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4D682BCE-A5D6-47E6-A236-CFDE5FBDF2ED
var IID_IXboxLiveQualityOfServiceMeasurement = syscall.GUID{0x4D682BCE, 0xA5D6, 0x47E6,
	[8]byte{0xA2, 0x36, 0xCF, 0xDE, 0x5F, 0xBD, 0xF2, 0xED}}

type IXboxLiveQualityOfServiceMeasurementInterface interface {
	win32.IInspectableInterface
	MeasureAsync() *IAsyncAction
	GetMetricResultsForDevice(deviceAddress *IXboxLiveDeviceAddress) *IVectorView[*IXboxLiveQualityOfServiceMetricResult]
	GetMetricResultsForMetric(metric XboxLiveQualityOfServiceMetric) *IVectorView[*IXboxLiveQualityOfServiceMetricResult]
	GetMetricResult(deviceAddress *IXboxLiveDeviceAddress, metric XboxLiveQualityOfServiceMetric) *IXboxLiveQualityOfServiceMetricResult
	GetPrivatePayloadResult(deviceAddress *IXboxLiveDeviceAddress) *IXboxLiveQualityOfServicePrivatePayloadResult
	Get_Metrics() *IVector[XboxLiveQualityOfServiceMetric]
	Get_DeviceAddresses() *IVector[*IXboxLiveDeviceAddress]
	Get_ShouldRequestPrivatePayloads() bool
	Put_ShouldRequestPrivatePayloads(value bool)
	Get_TimeoutInMilliseconds() uint32
	Put_TimeoutInMilliseconds(value uint32)
	Get_NumberOfProbesToAttempt() uint32
	Put_NumberOfProbesToAttempt(value uint32)
	Get_NumberOfResultsPending() uint32
	Get_MetricResults() *IVectorView[*IXboxLiveQualityOfServiceMetricResult]
	Get_PrivatePayloadResults() *IVectorView[*IXboxLiveQualityOfServicePrivatePayloadResult]
}

type IXboxLiveQualityOfServiceMeasurementVtbl struct {
	win32.IInspectableVtbl
	MeasureAsync                     uintptr
	GetMetricResultsForDevice        uintptr
	GetMetricResultsForMetric        uintptr
	GetMetricResult                  uintptr
	GetPrivatePayloadResult          uintptr
	Get_Metrics                      uintptr
	Get_DeviceAddresses              uintptr
	Get_ShouldRequestPrivatePayloads uintptr
	Put_ShouldRequestPrivatePayloads uintptr
	Get_TimeoutInMilliseconds        uintptr
	Put_TimeoutInMilliseconds        uintptr
	Get_NumberOfProbesToAttempt      uintptr
	Put_NumberOfProbesToAttempt      uintptr
	Get_NumberOfResultsPending       uintptr
	Get_MetricResults                uintptr
	Get_PrivatePayloadResults        uintptr
}

type IXboxLiveQualityOfServiceMeasurement struct {
	win32.IInspectable
}

func (this *IXboxLiveQualityOfServiceMeasurement) Vtbl() *IXboxLiveQualityOfServiceMeasurementVtbl {
	return (*IXboxLiveQualityOfServiceMeasurementVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXboxLiveQualityOfServiceMeasurement) MeasureAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MeasureAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IXboxLiveQualityOfServiceMeasurement) GetMetricResultsForDevice(deviceAddress *IXboxLiveDeviceAddress) *IVectorView[*IXboxLiveQualityOfServiceMetricResult] {
	var _result *IVectorView[*IXboxLiveQualityOfServiceMetricResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetMetricResultsForDevice, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(deviceAddress)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IXboxLiveQualityOfServiceMeasurement) GetMetricResultsForMetric(metric XboxLiveQualityOfServiceMetric) *IVectorView[*IXboxLiveQualityOfServiceMetricResult] {
	var _result *IVectorView[*IXboxLiveQualityOfServiceMetricResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetMetricResultsForMetric, uintptr(unsafe.Pointer(this)), uintptr(metric), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IXboxLiveQualityOfServiceMeasurement) GetMetricResult(deviceAddress *IXboxLiveDeviceAddress, metric XboxLiveQualityOfServiceMetric) *IXboxLiveQualityOfServiceMetricResult {
	var _result *IXboxLiveQualityOfServiceMetricResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetMetricResult, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(deviceAddress)), uintptr(metric), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IXboxLiveQualityOfServiceMeasurement) GetPrivatePayloadResult(deviceAddress *IXboxLiveDeviceAddress) *IXboxLiveQualityOfServicePrivatePayloadResult {
	var _result *IXboxLiveQualityOfServicePrivatePayloadResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPrivatePayloadResult, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(deviceAddress)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IXboxLiveQualityOfServiceMeasurement) Get_Metrics() *IVector[XboxLiveQualityOfServiceMetric] {
	var _result *IVector[XboxLiveQualityOfServiceMetric]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Metrics, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IXboxLiveQualityOfServiceMeasurement) Get_DeviceAddresses() *IVector[*IXboxLiveDeviceAddress] {
	var _result *IVector[*IXboxLiveDeviceAddress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceAddresses, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IXboxLiveQualityOfServiceMeasurement) Get_ShouldRequestPrivatePayloads() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ShouldRequestPrivatePayloads, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IXboxLiveQualityOfServiceMeasurement) Put_ShouldRequestPrivatePayloads(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ShouldRequestPrivatePayloads, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IXboxLiveQualityOfServiceMeasurement) Get_TimeoutInMilliseconds() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TimeoutInMilliseconds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IXboxLiveQualityOfServiceMeasurement) Put_TimeoutInMilliseconds(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TimeoutInMilliseconds, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IXboxLiveQualityOfServiceMeasurement) Get_NumberOfProbesToAttempt() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NumberOfProbesToAttempt, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IXboxLiveQualityOfServiceMeasurement) Put_NumberOfProbesToAttempt(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_NumberOfProbesToAttempt, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IXboxLiveQualityOfServiceMeasurement) Get_NumberOfResultsPending() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NumberOfResultsPending, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IXboxLiveQualityOfServiceMeasurement) Get_MetricResults() *IVectorView[*IXboxLiveQualityOfServiceMetricResult] {
	var _result *IVectorView[*IXboxLiveQualityOfServiceMetricResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MetricResults, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IXboxLiveQualityOfServiceMeasurement) Get_PrivatePayloadResults() *IVectorView[*IXboxLiveQualityOfServicePrivatePayloadResult] {
	var _result *IVectorView[*IXboxLiveQualityOfServicePrivatePayloadResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PrivatePayloadResults, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6E352DCA-23CF-440A-B077-5E30857A8234
var IID_IXboxLiveQualityOfServiceMeasurementStatics = syscall.GUID{0x6E352DCA, 0x23CF, 0x440A,
	[8]byte{0xB0, 0x77, 0x5E, 0x30, 0x85, 0x7A, 0x82, 0x34}}

type IXboxLiveQualityOfServiceMeasurementStaticsInterface interface {
	win32.IInspectableInterface
	PublishPrivatePayloadBytes(payloadLength uint32, payload *byte)
	ClearPrivatePayload()
	Get_MaxSimultaneousProbeConnections() uint32
	Put_MaxSimultaneousProbeConnections(value uint32)
	Get_IsSystemOutboundBandwidthConstrained() bool
	Put_IsSystemOutboundBandwidthConstrained(value bool)
	Get_IsSystemInboundBandwidthConstrained() bool
	Put_IsSystemInboundBandwidthConstrained(value bool)
	Get_PublishedPrivatePayload() *IBuffer
	Put_PublishedPrivatePayload(value *IBuffer)
	Get_MaxPrivatePayloadSize() uint32
}

type IXboxLiveQualityOfServiceMeasurementStaticsVtbl struct {
	win32.IInspectableVtbl
	PublishPrivatePayloadBytes               uintptr
	ClearPrivatePayload                      uintptr
	Get_MaxSimultaneousProbeConnections      uintptr
	Put_MaxSimultaneousProbeConnections      uintptr
	Get_IsSystemOutboundBandwidthConstrained uintptr
	Put_IsSystemOutboundBandwidthConstrained uintptr
	Get_IsSystemInboundBandwidthConstrained  uintptr
	Put_IsSystemInboundBandwidthConstrained  uintptr
	Get_PublishedPrivatePayload              uintptr
	Put_PublishedPrivatePayload              uintptr
	Get_MaxPrivatePayloadSize                uintptr
}

type IXboxLiveQualityOfServiceMeasurementStatics struct {
	win32.IInspectable
}

func (this *IXboxLiveQualityOfServiceMeasurementStatics) Vtbl() *IXboxLiveQualityOfServiceMeasurementStaticsVtbl {
	return (*IXboxLiveQualityOfServiceMeasurementStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXboxLiveQualityOfServiceMeasurementStatics) PublishPrivatePayloadBytes(payloadLength uint32, payload *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PublishPrivatePayloadBytes, uintptr(unsafe.Pointer(this)), uintptr(payloadLength), uintptr(unsafe.Pointer(payload)))
	_ = _hr
}

func (this *IXboxLiveQualityOfServiceMeasurementStatics) ClearPrivatePayload() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearPrivatePayload, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IXboxLiveQualityOfServiceMeasurementStatics) Get_MaxSimultaneousProbeConnections() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxSimultaneousProbeConnections, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IXboxLiveQualityOfServiceMeasurementStatics) Put_MaxSimultaneousProbeConnections(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxSimultaneousProbeConnections, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IXboxLiveQualityOfServiceMeasurementStatics) Get_IsSystemOutboundBandwidthConstrained() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSystemOutboundBandwidthConstrained, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IXboxLiveQualityOfServiceMeasurementStatics) Put_IsSystemOutboundBandwidthConstrained(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsSystemOutboundBandwidthConstrained, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IXboxLiveQualityOfServiceMeasurementStatics) Get_IsSystemInboundBandwidthConstrained() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSystemInboundBandwidthConstrained, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IXboxLiveQualityOfServiceMeasurementStatics) Put_IsSystemInboundBandwidthConstrained(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsSystemInboundBandwidthConstrained, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IXboxLiveQualityOfServiceMeasurementStatics) Get_PublishedPrivatePayload() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PublishedPrivatePayload, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IXboxLiveQualityOfServiceMeasurementStatics) Put_PublishedPrivatePayload(value *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PublishedPrivatePayload, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IXboxLiveQualityOfServiceMeasurementStatics) Get_MaxPrivatePayloadSize() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxPrivatePayloadSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// AEEC53D1-3561-4782-B0CF-D3AE29D9FA87
var IID_IXboxLiveQualityOfServiceMetricResult = syscall.GUID{0xAEEC53D1, 0x3561, 0x4782,
	[8]byte{0xB0, 0xCF, 0xD3, 0xAE, 0x29, 0xD9, 0xFA, 0x87}}

type IXboxLiveQualityOfServiceMetricResultInterface interface {
	win32.IInspectableInterface
	Get_Status() XboxLiveQualityOfServiceMeasurementStatus
	Get_DeviceAddress() *IXboxLiveDeviceAddress
	Get_Metric() XboxLiveQualityOfServiceMetric
	Get_Value() uint64
}

type IXboxLiveQualityOfServiceMetricResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status        uintptr
	Get_DeviceAddress uintptr
	Get_Metric        uintptr
	Get_Value         uintptr
}

type IXboxLiveQualityOfServiceMetricResult struct {
	win32.IInspectable
}

func (this *IXboxLiveQualityOfServiceMetricResult) Vtbl() *IXboxLiveQualityOfServiceMetricResultVtbl {
	return (*IXboxLiveQualityOfServiceMetricResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXboxLiveQualityOfServiceMetricResult) Get_Status() XboxLiveQualityOfServiceMeasurementStatus {
	var _result XboxLiveQualityOfServiceMeasurementStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IXboxLiveQualityOfServiceMetricResult) Get_DeviceAddress() *IXboxLiveDeviceAddress {
	var _result *IXboxLiveDeviceAddress
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IXboxLiveQualityOfServiceMetricResult) Get_Metric() XboxLiveQualityOfServiceMetric {
	var _result XboxLiveQualityOfServiceMetric
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Metric, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IXboxLiveQualityOfServiceMetricResult) Get_Value() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5A6302AE-6F38-41C0-9FCC-EA6CB978CAFC
var IID_IXboxLiveQualityOfServicePrivatePayloadResult = syscall.GUID{0x5A6302AE, 0x6F38, 0x41C0,
	[8]byte{0x9F, 0xCC, 0xEA, 0x6C, 0xB9, 0x78, 0xCA, 0xFC}}

type IXboxLiveQualityOfServicePrivatePayloadResultInterface interface {
	win32.IInspectableInterface
	Get_Status() XboxLiveQualityOfServiceMeasurementStatus
	Get_DeviceAddress() *IXboxLiveDeviceAddress
	Get_Value() *IBuffer
}

type IXboxLiveQualityOfServicePrivatePayloadResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status        uintptr
	Get_DeviceAddress uintptr
	Get_Value         uintptr
}

type IXboxLiveQualityOfServicePrivatePayloadResult struct {
	win32.IInspectable
}

func (this *IXboxLiveQualityOfServicePrivatePayloadResult) Vtbl() *IXboxLiveQualityOfServicePrivatePayloadResultVtbl {
	return (*IXboxLiveQualityOfServicePrivatePayloadResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXboxLiveQualityOfServicePrivatePayloadResult) Get_Status() XboxLiveQualityOfServiceMeasurementStatus {
	var _result XboxLiveQualityOfServiceMeasurementStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IXboxLiveQualityOfServicePrivatePayloadResult) Get_DeviceAddress() *IXboxLiveDeviceAddress {
	var _result *IXboxLiveDeviceAddress
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IXboxLiveQualityOfServicePrivatePayloadResult) Get_Value() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type XboxLiveDeviceAddress struct {
	RtClass
	*IXboxLiveDeviceAddress
}

func NewIXboxLiveDeviceAddressStatics() *IXboxLiveDeviceAddressStatics {
	var p *IXboxLiveDeviceAddressStatics
	hs := NewHStr("Windows.Networking.XboxLive.XboxLiveDeviceAddress")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IXboxLiveDeviceAddressStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type XboxLiveEndpointPair struct {
	RtClass
	*IXboxLiveEndpointPair
}

func NewIXboxLiveEndpointPairStatics() *IXboxLiveEndpointPairStatics {
	var p *IXboxLiveEndpointPairStatics
	hs := NewHStr("Windows.Networking.XboxLive.XboxLiveEndpointPair")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IXboxLiveEndpointPairStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type XboxLiveEndpointPairCreationResult struct {
	RtClass
	*IXboxLiveEndpointPairCreationResult
}

type XboxLiveEndpointPairStateChangedEventArgs struct {
	RtClass
	*IXboxLiveEndpointPairStateChangedEventArgs
}

type XboxLiveEndpointPairTemplate struct {
	RtClass
	*IXboxLiveEndpointPairTemplate
}

func NewIXboxLiveEndpointPairTemplateStatics() *IXboxLiveEndpointPairTemplateStatics {
	var p *IXboxLiveEndpointPairTemplateStatics
	hs := NewHStr("Windows.Networking.XboxLive.XboxLiveEndpointPairTemplate")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IXboxLiveEndpointPairTemplateStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type XboxLiveInboundEndpointPairCreatedEventArgs struct {
	RtClass
	*IXboxLiveInboundEndpointPairCreatedEventArgs
}

type XboxLiveQualityOfServiceMeasurement struct {
	RtClass
	*IXboxLiveQualityOfServiceMeasurement
}

func NewXboxLiveQualityOfServiceMeasurement() *XboxLiveQualityOfServiceMeasurement {
	hs := NewHStr("Windows.Networking.XboxLive.XboxLiveQualityOfServiceMeasurement")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &XboxLiveQualityOfServiceMeasurement{
		RtClass:                              RtClass{PInspect: p},
		IXboxLiveQualityOfServiceMeasurement: (*IXboxLiveQualityOfServiceMeasurement)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewIXboxLiveQualityOfServiceMeasurementStatics() *IXboxLiveQualityOfServiceMeasurementStatics {
	var p *IXboxLiveQualityOfServiceMeasurementStatics
	hs := NewHStr("Windows.Networking.XboxLive.XboxLiveQualityOfServiceMeasurement")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IXboxLiveQualityOfServiceMeasurementStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type XboxLiveQualityOfServiceMetricResult struct {
	RtClass
	*IXboxLiveQualityOfServiceMetricResult
}

type XboxLiveQualityOfServicePrivatePayloadResult struct {
	RtClass
	*IXboxLiveQualityOfServicePrivatePayloadResult
}
