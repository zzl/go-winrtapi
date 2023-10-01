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
type AppServiceClosedStatus int32

const (
	AppServiceClosedStatus_Completed              AppServiceClosedStatus = 0
	AppServiceClosedStatus_Canceled               AppServiceClosedStatus = 1
	AppServiceClosedStatus_ResourceLimitsExceeded AppServiceClosedStatus = 2
	AppServiceClosedStatus_Unknown                AppServiceClosedStatus = 3
)

// enum
type AppServiceConnectionStatus int32

const (
	AppServiceConnectionStatus_Success                       AppServiceConnectionStatus = 0
	AppServiceConnectionStatus_AppNotInstalled               AppServiceConnectionStatus = 1
	AppServiceConnectionStatus_AppUnavailable                AppServiceConnectionStatus = 2
	AppServiceConnectionStatus_AppServiceUnavailable         AppServiceConnectionStatus = 3
	AppServiceConnectionStatus_Unknown                       AppServiceConnectionStatus = 4
	AppServiceConnectionStatus_RemoteSystemUnavailable       AppServiceConnectionStatus = 5
	AppServiceConnectionStatus_RemoteSystemNotSupportedByApp AppServiceConnectionStatus = 6
	AppServiceConnectionStatus_NotAuthorized                 AppServiceConnectionStatus = 7
	AppServiceConnectionStatus_AuthenticationError           AppServiceConnectionStatus = 8
	AppServiceConnectionStatus_NetworkNotAvailable           AppServiceConnectionStatus = 9
	AppServiceConnectionStatus_DisabledByPolicy              AppServiceConnectionStatus = 10
	AppServiceConnectionStatus_WebServiceUnavailable         AppServiceConnectionStatus = 11
)

// enum
type AppServiceResponseStatus int32

const (
	AppServiceResponseStatus_Success                 AppServiceResponseStatus = 0
	AppServiceResponseStatus_Failure                 AppServiceResponseStatus = 1
	AppServiceResponseStatus_ResourceLimitsExceeded  AppServiceResponseStatus = 2
	AppServiceResponseStatus_Unknown                 AppServiceResponseStatus = 3
	AppServiceResponseStatus_RemoteSystemUnavailable AppServiceResponseStatus = 4
	AppServiceResponseStatus_MessageSizeTooLarge     AppServiceResponseStatus = 5
	AppServiceResponseStatus_AppUnavailable          AppServiceResponseStatus = 6
	AppServiceResponseStatus_AuthenticationError     AppServiceResponseStatus = 7
	AppServiceResponseStatus_NetworkNotAvailable     AppServiceResponseStatus = 8
	AppServiceResponseStatus_DisabledByPolicy        AppServiceResponseStatus = 9
	AppServiceResponseStatus_WebServiceUnavailable   AppServiceResponseStatus = 10
)

// enum
type StatelessAppServiceResponseStatus int32

const (
	StatelessAppServiceResponseStatus_Success                       StatelessAppServiceResponseStatus = 0
	StatelessAppServiceResponseStatus_AppNotInstalled               StatelessAppServiceResponseStatus = 1
	StatelessAppServiceResponseStatus_AppUnavailable                StatelessAppServiceResponseStatus = 2
	StatelessAppServiceResponseStatus_AppServiceUnavailable         StatelessAppServiceResponseStatus = 3
	StatelessAppServiceResponseStatus_RemoteSystemUnavailable       StatelessAppServiceResponseStatus = 4
	StatelessAppServiceResponseStatus_RemoteSystemNotSupportedByApp StatelessAppServiceResponseStatus = 5
	StatelessAppServiceResponseStatus_NotAuthorized                 StatelessAppServiceResponseStatus = 6
	StatelessAppServiceResponseStatus_ResourceLimitsExceeded        StatelessAppServiceResponseStatus = 7
	StatelessAppServiceResponseStatus_MessageSizeTooLarge           StatelessAppServiceResponseStatus = 8
	StatelessAppServiceResponseStatus_Failure                       StatelessAppServiceResponseStatus = 9
	StatelessAppServiceResponseStatus_Unknown                       StatelessAppServiceResponseStatus = 10
	StatelessAppServiceResponseStatus_AuthenticationError           StatelessAppServiceResponseStatus = 11
	StatelessAppServiceResponseStatus_NetworkNotAvailable           StatelessAppServiceResponseStatus = 12
	StatelessAppServiceResponseStatus_DisabledByPolicy              StatelessAppServiceResponseStatus = 13
	StatelessAppServiceResponseStatus_WebServiceUnavailable         StatelessAppServiceResponseStatus = 14
)

// interfaces

// EF0D2507-D132-4C85-8395-3C31D5A1E941
var IID_IAppServiceCatalogStatics = syscall.GUID{0xEF0D2507, 0xD132, 0x4C85,
	[8]byte{0x83, 0x95, 0x3C, 0x31, 0xD5, 0xA1, 0xE9, 0x41}}

type IAppServiceCatalogStaticsInterface interface {
	win32.IInspectableInterface
	FindAppServiceProvidersAsync(appServiceName string) *IAsyncOperation[*IVectorView[*IAppInfo]]
}

type IAppServiceCatalogStaticsVtbl struct {
	win32.IInspectableVtbl
	FindAppServiceProvidersAsync uintptr
}

type IAppServiceCatalogStatics struct {
	win32.IInspectable
}

func (this *IAppServiceCatalogStatics) Vtbl() *IAppServiceCatalogStaticsVtbl {
	return (*IAppServiceCatalogStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppServiceCatalogStatics) FindAppServiceProvidersAsync(appServiceName string) *IAsyncOperation[*IVectorView[*IAppInfo]] {
	var _result *IAsyncOperation[*IVectorView[*IAppInfo]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAppServiceProvidersAsync, uintptr(unsafe.Pointer(this)), NewHStr(appServiceName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DE6016F6-CB03-4D35-AC8D-CC6303239731
var IID_IAppServiceClosedEventArgs = syscall.GUID{0xDE6016F6, 0xCB03, 0x4D35,
	[8]byte{0xAC, 0x8D, 0xCC, 0x63, 0x03, 0x23, 0x97, 0x31}}

type IAppServiceClosedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Status() AppServiceClosedStatus
}

type IAppServiceClosedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
}

type IAppServiceClosedEventArgs struct {
	win32.IInspectable
}

func (this *IAppServiceClosedEventArgs) Vtbl() *IAppServiceClosedEventArgsVtbl {
	return (*IAppServiceClosedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppServiceClosedEventArgs) Get_Status() AppServiceClosedStatus {
	var _result AppServiceClosedStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 9DD474A2-871F-4D52-89A9-9E090531BD27
var IID_IAppServiceConnection = syscall.GUID{0x9DD474A2, 0x871F, 0x4D52,
	[8]byte{0x89, 0xA9, 0x9E, 0x09, 0x05, 0x31, 0xBD, 0x27}}

type IAppServiceConnectionInterface interface {
	win32.IInspectableInterface
	Get_AppServiceName() string
	Put_AppServiceName(value string)
	Get_PackageFamilyName() string
	Put_PackageFamilyName(value string)
	OpenAsync() *IAsyncOperation[AppServiceConnectionStatus]
	SendMessageAsync(message *IPropertySet) *IAsyncOperation[*IAppServiceResponse]
	Add_RequestReceived(handler TypedEventHandler[*IAppServiceConnection, *IAppServiceRequestReceivedEventArgs]) EventRegistrationToken
	Remove_RequestReceived(token EventRegistrationToken)
	Add_ServiceClosed(handler TypedEventHandler[*IAppServiceConnection, *IAppServiceClosedEventArgs]) EventRegistrationToken
	Remove_ServiceClosed(token EventRegistrationToken)
}

type IAppServiceConnectionVtbl struct {
	win32.IInspectableVtbl
	Get_AppServiceName     uintptr
	Put_AppServiceName     uintptr
	Get_PackageFamilyName  uintptr
	Put_PackageFamilyName  uintptr
	OpenAsync              uintptr
	SendMessageAsync       uintptr
	Add_RequestReceived    uintptr
	Remove_RequestReceived uintptr
	Add_ServiceClosed      uintptr
	Remove_ServiceClosed   uintptr
}

type IAppServiceConnection struct {
	win32.IInspectable
}

func (this *IAppServiceConnection) Vtbl() *IAppServiceConnectionVtbl {
	return (*IAppServiceConnectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppServiceConnection) Get_AppServiceName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppServiceName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppServiceConnection) Put_AppServiceName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AppServiceName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IAppServiceConnection) Get_PackageFamilyName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PackageFamilyName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppServiceConnection) Put_PackageFamilyName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PackageFamilyName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IAppServiceConnection) OpenAsync() *IAsyncOperation[AppServiceConnectionStatus] {
	var _result *IAsyncOperation[AppServiceConnectionStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppServiceConnection) SendMessageAsync(message *IPropertySet) *IAsyncOperation[*IAppServiceResponse] {
	var _result *IAsyncOperation[*IAppServiceResponse]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendMessageAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(message)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppServiceConnection) Add_RequestReceived(handler TypedEventHandler[*IAppServiceConnection, *IAppServiceRequestReceivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_RequestReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppServiceConnection) Remove_RequestReceived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_RequestReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAppServiceConnection) Add_ServiceClosed(handler TypedEventHandler[*IAppServiceConnection, *IAppServiceClosedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ServiceClosed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAppServiceConnection) Remove_ServiceClosed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ServiceClosed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 8BDFCD5F-2302-4FBD-8061-52511C2F8BF9
var IID_IAppServiceConnection2 = syscall.GUID{0x8BDFCD5F, 0x2302, 0x4FBD,
	[8]byte{0x80, 0x61, 0x52, 0x51, 0x1C, 0x2F, 0x8B, 0xF9}}

type IAppServiceConnection2Interface interface {
	win32.IInspectableInterface
	OpenRemoteAsync(remoteSystemConnectionRequest unsafe.Pointer) *IAsyncOperation[AppServiceConnectionStatus]
	Get_User() *IUser
	Put_User(value *IUser)
}

type IAppServiceConnection2Vtbl struct {
	win32.IInspectableVtbl
	OpenRemoteAsync uintptr
	Get_User        uintptr
	Put_User        uintptr
}

type IAppServiceConnection2 struct {
	win32.IInspectable
}

func (this *IAppServiceConnection2) Vtbl() *IAppServiceConnection2Vtbl {
	return (*IAppServiceConnection2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppServiceConnection2) OpenRemoteAsync(remoteSystemConnectionRequest unsafe.Pointer) *IAsyncOperation[AppServiceConnectionStatus] {
	var _result *IAsyncOperation[AppServiceConnectionStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenRemoteAsync, uintptr(unsafe.Pointer(this)), uintptr(remoteSystemConnectionRequest), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppServiceConnection2) Get_User() *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_User, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppServiceConnection2) Put_User(value *IUser) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_User, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// ADC56CE9-D408-5673-8637-827A4B274168
var IID_IAppServiceConnectionStatics = syscall.GUID{0xADC56CE9, 0xD408, 0x5673,
	[8]byte{0x86, 0x37, 0x82, 0x7A, 0x4B, 0x27, 0x41, 0x68}}

type IAppServiceConnectionStaticsInterface interface {
	win32.IInspectableInterface
	SendStatelessMessageAsync(connection *IAppServiceConnection, connectionRequest unsafe.Pointer, message *IPropertySet) *IAsyncOperation[*IStatelessAppServiceResponse]
}

type IAppServiceConnectionStaticsVtbl struct {
	win32.IInspectableVtbl
	SendStatelessMessageAsync uintptr
}

type IAppServiceConnectionStatics struct {
	win32.IInspectable
}

func (this *IAppServiceConnectionStatics) Vtbl() *IAppServiceConnectionStaticsVtbl {
	return (*IAppServiceConnectionStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppServiceConnectionStatics) SendStatelessMessageAsync(connection *IAppServiceConnection, connectionRequest unsafe.Pointer, message *IPropertySet) *IAsyncOperation[*IStatelessAppServiceResponse] {
	var _result *IAsyncOperation[*IStatelessAppServiceResponse]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendStatelessMessageAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(connection)), uintptr(connectionRequest), uintptr(unsafe.Pointer(message)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7E1B5322-EAB0-4248-AE04-FDF93838E472
var IID_IAppServiceDeferral = syscall.GUID{0x7E1B5322, 0xEAB0, 0x4248,
	[8]byte{0xAE, 0x04, 0xFD, 0xF9, 0x38, 0x38, 0xE4, 0x72}}

type IAppServiceDeferralInterface interface {
	win32.IInspectableInterface
	Complete()
}

type IAppServiceDeferralVtbl struct {
	win32.IInspectableVtbl
	Complete uintptr
}

type IAppServiceDeferral struct {
	win32.IInspectable
}

func (this *IAppServiceDeferral) Vtbl() *IAppServiceDeferralVtbl {
	return (*IAppServiceDeferralVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppServiceDeferral) Complete() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Complete, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 20E58D9D-18DE-4B01-80BA-90A76204E3C8
var IID_IAppServiceRequest = syscall.GUID{0x20E58D9D, 0x18DE, 0x4B01,
	[8]byte{0x80, 0xBA, 0x90, 0xA7, 0x62, 0x04, 0xE3, 0xC8}}

type IAppServiceRequestInterface interface {
	win32.IInspectableInterface
	Get_Message() *IPropertySet
	SendResponseAsync(message *IPropertySet) *IAsyncOperation[AppServiceResponseStatus]
}

type IAppServiceRequestVtbl struct {
	win32.IInspectableVtbl
	Get_Message       uintptr
	SendResponseAsync uintptr
}

type IAppServiceRequest struct {
	win32.IInspectable
}

func (this *IAppServiceRequest) Vtbl() *IAppServiceRequestVtbl {
	return (*IAppServiceRequestVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppServiceRequest) Get_Message() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Message, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppServiceRequest) SendResponseAsync(message *IPropertySet) *IAsyncOperation[AppServiceResponseStatus] {
	var _result *IAsyncOperation[AppServiceResponseStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendResponseAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(message)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6E122360-FF65-44AE-9E45-857FE4180681
var IID_IAppServiceRequestReceivedEventArgs = syscall.GUID{0x6E122360, 0xFF65, 0x44AE,
	[8]byte{0x9E, 0x45, 0x85, 0x7F, 0xE4, 0x18, 0x06, 0x81}}

type IAppServiceRequestReceivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Request() *IAppServiceRequest
	GetDeferral() *IAppServiceDeferral
}

type IAppServiceRequestReceivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Request uintptr
	GetDeferral uintptr
}

type IAppServiceRequestReceivedEventArgs struct {
	win32.IInspectable
}

func (this *IAppServiceRequestReceivedEventArgs) Vtbl() *IAppServiceRequestReceivedEventArgsVtbl {
	return (*IAppServiceRequestReceivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppServiceRequestReceivedEventArgs) Get_Request() *IAppServiceRequest {
	var _result *IAppServiceRequest
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Request, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppServiceRequestReceivedEventArgs) GetDeferral() *IAppServiceDeferral {
	var _result *IAppServiceDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8D503CEC-9AA3-4E68-9559-9DE63E372CE4
var IID_IAppServiceResponse = syscall.GUID{0x8D503CEC, 0x9AA3, 0x4E68,
	[8]byte{0x95, 0x59, 0x9D, 0xE6, 0x3E, 0x37, 0x2C, 0xE4}}

type IAppServiceResponseInterface interface {
	win32.IInspectableInterface
	Get_Message() *IPropertySet
	Get_Status() AppServiceResponseStatus
}

type IAppServiceResponseVtbl struct {
	win32.IInspectableVtbl
	Get_Message uintptr
	Get_Status  uintptr
}

type IAppServiceResponse struct {
	win32.IInspectable
}

func (this *IAppServiceResponse) Vtbl() *IAppServiceResponseVtbl {
	return (*IAppServiceResponseVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppServiceResponse) Get_Message() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Message, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppServiceResponse) Get_Status() AppServiceResponseStatus {
	var _result AppServiceResponseStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 88A2DCAC-AD28-41B8-80BB-BDF1B2169E19
var IID_IAppServiceTriggerDetails = syscall.GUID{0x88A2DCAC, 0xAD28, 0x41B8,
	[8]byte{0x80, 0xBB, 0xBD, 0xF1, 0xB2, 0x16, 0x9E, 0x19}}

type IAppServiceTriggerDetailsInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Get_CallerPackageFamilyName() string
	Get_AppServiceConnection() *IAppServiceConnection
}

type IAppServiceTriggerDetailsVtbl struct {
	win32.IInspectableVtbl
	Get_Name                    uintptr
	Get_CallerPackageFamilyName uintptr
	Get_AppServiceConnection    uintptr
}

type IAppServiceTriggerDetails struct {
	win32.IInspectable
}

func (this *IAppServiceTriggerDetails) Vtbl() *IAppServiceTriggerDetailsVtbl {
	return (*IAppServiceTriggerDetailsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppServiceTriggerDetails) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppServiceTriggerDetails) Get_CallerPackageFamilyName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CallerPackageFamilyName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAppServiceTriggerDetails) Get_AppServiceConnection() *IAppServiceConnection {
	var _result *IAppServiceConnection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppServiceConnection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E83D54B2-28CC-43F2-B465-C0482E59E2DC
var IID_IAppServiceTriggerDetails2 = syscall.GUID{0xE83D54B2, 0x28CC, 0x43F2,
	[8]byte{0xB4, 0x65, 0xC0, 0x48, 0x2E, 0x59, 0xE2, 0xDC}}

type IAppServiceTriggerDetails2Interface interface {
	win32.IInspectableInterface
	Get_IsRemoteSystemConnection() bool
}

type IAppServiceTriggerDetails2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsRemoteSystemConnection uintptr
}

type IAppServiceTriggerDetails2 struct {
	win32.IInspectable
}

func (this *IAppServiceTriggerDetails2) Vtbl() *IAppServiceTriggerDetails2Vtbl {
	return (*IAppServiceTriggerDetails2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppServiceTriggerDetails2) Get_IsRemoteSystemConnection() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsRemoteSystemConnection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// FBD71E21-7939-4E68-9E3C-7780147AABB6
var IID_IAppServiceTriggerDetails3 = syscall.GUID{0xFBD71E21, 0x7939, 0x4E68,
	[8]byte{0x9E, 0x3C, 0x77, 0x80, 0x14, 0x7A, 0xAB, 0xB6}}

type IAppServiceTriggerDetails3Interface interface {
	win32.IInspectableInterface
	CheckCallerForCapabilityAsync(capabilityName string) *IAsyncOperation[bool]
}

type IAppServiceTriggerDetails3Vtbl struct {
	win32.IInspectableVtbl
	CheckCallerForCapabilityAsync uintptr
}

type IAppServiceTriggerDetails3 struct {
	win32.IInspectable
}

func (this *IAppServiceTriggerDetails3) Vtbl() *IAppServiceTriggerDetails3Vtbl {
	return (*IAppServiceTriggerDetails3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppServiceTriggerDetails3) CheckCallerForCapabilityAsync(capabilityName string) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CheckCallerForCapabilityAsync, uintptr(unsafe.Pointer(this)), NewHStr(capabilityName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1185B180-8861-5E30-AB55-1CF4D08BBF6D
var IID_IAppServiceTriggerDetails4 = syscall.GUID{0x1185B180, 0x8861, 0x5E30,
	[8]byte{0xAB, 0x55, 0x1C, 0xF4, 0xD0, 0x8B, 0xBF, 0x6D}}

type IAppServiceTriggerDetails4Interface interface {
	win32.IInspectableInterface
	Get_CallerRemoteConnectionToken() string
}

type IAppServiceTriggerDetails4Vtbl struct {
	win32.IInspectableVtbl
	Get_CallerRemoteConnectionToken uintptr
}

type IAppServiceTriggerDetails4 struct {
	win32.IInspectable
}

func (this *IAppServiceTriggerDetails4) Vtbl() *IAppServiceTriggerDetails4Vtbl {
	return (*IAppServiceTriggerDetails4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppServiceTriggerDetails4) Get_CallerRemoteConnectionToken() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CallerRemoteConnectionToken, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 43754AF7-A9EC-52FE-82E7-939B68DC9388
var IID_IStatelessAppServiceResponse = syscall.GUID{0x43754AF7, 0xA9EC, 0x52FE,
	[8]byte{0x82, 0xE7, 0x93, 0x9B, 0x68, 0xDC, 0x93, 0x88}}

type IStatelessAppServiceResponseInterface interface {
	win32.IInspectableInterface
	Get_Message() *IPropertySet
	Get_Status() StatelessAppServiceResponseStatus
}

type IStatelessAppServiceResponseVtbl struct {
	win32.IInspectableVtbl
	Get_Message uintptr
	Get_Status  uintptr
}

type IStatelessAppServiceResponse struct {
	win32.IInspectable
}

func (this *IStatelessAppServiceResponse) Vtbl() *IStatelessAppServiceResponseVtbl {
	return (*IStatelessAppServiceResponseVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStatelessAppServiceResponse) Get_Message() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Message, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStatelessAppServiceResponse) Get_Status() StatelessAppServiceResponseStatus {
	var _result StatelessAppServiceResponseStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type AppServiceClosedEventArgs struct {
	RtClass
	*IAppServiceClosedEventArgs
}

type AppServiceConnection struct {
	RtClass
	*IAppServiceConnection
}

func NewAppServiceConnection() *AppServiceConnection {
	hs := NewHStr("Windows.ApplicationModel.AppService.AppServiceConnection")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &AppServiceConnection{
		RtClass:               RtClass{PInspect: p},
		IAppServiceConnection: (*IAppServiceConnection)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewIAppServiceConnectionStatics() *IAppServiceConnectionStatics {
	var p *IAppServiceConnectionStatics
	hs := NewHStr("Windows.ApplicationModel.AppService.AppServiceConnection")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAppServiceConnectionStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type AppServiceDeferral struct {
	RtClass
	*IAppServiceDeferral
}

type AppServiceRequest struct {
	RtClass
	*IAppServiceRequest
}

type AppServiceRequestReceivedEventArgs struct {
	RtClass
	*IAppServiceRequestReceivedEventArgs
}

type AppServiceResponse struct {
	RtClass
	*IAppServiceResponse
}

type AppServiceTriggerDetails struct {
	RtClass
	*IAppServiceTriggerDetails
}

type StatelessAppServiceResponse struct {
	RtClass
	*IStatelessAppServiceResponse
}
