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
type AllJoynAuthenticationMechanism int32

const (
	AllJoynAuthenticationMechanism_None         AllJoynAuthenticationMechanism = 0
	AllJoynAuthenticationMechanism_SrpAnonymous AllJoynAuthenticationMechanism = 1
	AllJoynAuthenticationMechanism_SrpLogon     AllJoynAuthenticationMechanism = 2
	AllJoynAuthenticationMechanism_EcdheNull    AllJoynAuthenticationMechanism = 3
	AllJoynAuthenticationMechanism_EcdhePsk     AllJoynAuthenticationMechanism = 4
	AllJoynAuthenticationMechanism_EcdheEcdsa   AllJoynAuthenticationMechanism = 5
	AllJoynAuthenticationMechanism_EcdheSpeke   AllJoynAuthenticationMechanism = 6
)

// enum
type AllJoynBusAttachmentState int32

const (
	AllJoynBusAttachmentState_Disconnected  AllJoynBusAttachmentState = 0
	AllJoynBusAttachmentState_Connecting    AllJoynBusAttachmentState = 1
	AllJoynBusAttachmentState_Connected     AllJoynBusAttachmentState = 2
	AllJoynBusAttachmentState_Disconnecting AllJoynBusAttachmentState = 3
)

// enum
type AllJoynSessionLostReason int32

const (
	AllJoynSessionLostReason_None                   AllJoynSessionLostReason = 0
	AllJoynSessionLostReason_ProducerLeftSession    AllJoynSessionLostReason = 1
	AllJoynSessionLostReason_ProducerClosedAbruptly AllJoynSessionLostReason = 2
	AllJoynSessionLostReason_RemovedByProducer      AllJoynSessionLostReason = 3
	AllJoynSessionLostReason_LinkTimeout            AllJoynSessionLostReason = 4
	AllJoynSessionLostReason_Other                  AllJoynSessionLostReason = 5
)

// enum
type AllJoynTrafficType int32

const (
	AllJoynTrafficType_Unknown       AllJoynTrafficType = 0
	AllJoynTrafficType_Messages      AllJoynTrafficType = 1
	AllJoynTrafficType_RawUnreliable AllJoynTrafficType = 2
	AllJoynTrafficType_RawReliable   AllJoynTrafficType = 4
)

// interfaces

// E5A9BF00-1FA2-4839-93EF-F9DF404890F7
var IID_IAllJoynAboutData = syscall.GUID{0xE5A9BF00, 0x1FA2, 0x4839,
	[8]byte{0x93, 0xEF, 0xF9, 0xDF, 0x40, 0x48, 0x90, 0xF7}}

type IAllJoynAboutDataInterface interface {
	win32.IInspectableInterface
	Get_IsEnabled() bool
	Put_IsEnabled(value bool)
	Get_DefaultAppName() string
	Put_DefaultAppName(value string)
	Get_AppNames() *IMap[string, string]
	Get_DateOfManufacture() *IReference[DateTime]
	Put_DateOfManufacture(value *IReference[DateTime])
	Get_DefaultDescription() string
	Put_DefaultDescription(value string)
	Get_Descriptions() *IMap[string, string]
	Get_DefaultManufacturer() string
	Put_DefaultManufacturer(value string)
	Get_Manufacturers() *IMap[string, string]
	Get_ModelNumber() string
	Put_ModelNumber(value string)
	Get_SoftwareVersion() string
	Put_SoftwareVersion(value string)
	Get_SupportUrl() *IUriRuntimeClass
	Put_SupportUrl(value *IUriRuntimeClass)
	Get_AppId() syscall.GUID
	Put_AppId(value syscall.GUID)
}

type IAllJoynAboutDataVtbl struct {
	win32.IInspectableVtbl
	Get_IsEnabled           uintptr
	Put_IsEnabled           uintptr
	Get_DefaultAppName      uintptr
	Put_DefaultAppName      uintptr
	Get_AppNames            uintptr
	Get_DateOfManufacture   uintptr
	Put_DateOfManufacture   uintptr
	Get_DefaultDescription  uintptr
	Put_DefaultDescription  uintptr
	Get_Descriptions        uintptr
	Get_DefaultManufacturer uintptr
	Put_DefaultManufacturer uintptr
	Get_Manufacturers       uintptr
	Get_ModelNumber         uintptr
	Put_ModelNumber         uintptr
	Get_SoftwareVersion     uintptr
	Put_SoftwareVersion     uintptr
	Get_SupportUrl          uintptr
	Put_SupportUrl          uintptr
	Get_AppId               uintptr
	Put_AppId               uintptr
}

type IAllJoynAboutData struct {
	win32.IInspectable
}

func (this *IAllJoynAboutData) Vtbl() *IAllJoynAboutDataVtbl {
	return (*IAllJoynAboutDataVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynAboutData) Get_IsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynAboutData) Put_IsEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAllJoynAboutData) Get_DefaultAppName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DefaultAppName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAllJoynAboutData) Put_DefaultAppName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DefaultAppName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IAllJoynAboutData) Get_AppNames() *IMap[string, string] {
	var _result *IMap[string, string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppNames, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAllJoynAboutData) Get_DateOfManufacture() *IReference[DateTime] {
	var _result *IReference[DateTime]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DateOfManufacture, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAllJoynAboutData) Put_DateOfManufacture(value *IReference[DateTime]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DateOfManufacture, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAllJoynAboutData) Get_DefaultDescription() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DefaultDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAllJoynAboutData) Put_DefaultDescription(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DefaultDescription, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IAllJoynAboutData) Get_Descriptions() *IMap[string, string] {
	var _result *IMap[string, string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Descriptions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAllJoynAboutData) Get_DefaultManufacturer() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DefaultManufacturer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAllJoynAboutData) Put_DefaultManufacturer(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DefaultManufacturer, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IAllJoynAboutData) Get_Manufacturers() *IMap[string, string] {
	var _result *IMap[string, string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Manufacturers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAllJoynAboutData) Get_ModelNumber() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ModelNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAllJoynAboutData) Put_ModelNumber(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ModelNumber, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IAllJoynAboutData) Get_SoftwareVersion() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SoftwareVersion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAllJoynAboutData) Put_SoftwareVersion(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SoftwareVersion, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IAllJoynAboutData) Get_SupportUrl() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportUrl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAllJoynAboutData) Put_SupportUrl(value *IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SupportUrl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAllJoynAboutData) Get_AppId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynAboutData) Put_AppId(value syscall.GUID) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AppId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

// 6823111F-6212-4934-9C48-E19CA4984288
var IID_IAllJoynAboutDataView = syscall.GUID{0x6823111F, 0x6212, 0x4934,
	[8]byte{0x9C, 0x48, 0xE1, 0x9C, 0xA4, 0x98, 0x42, 0x88}}

type IAllJoynAboutDataViewInterface interface {
	win32.IInspectableInterface
	Get_Status() int32
	Get_Properties() *IMapView[string, interface{}]
	Get_AJSoftwareVersion() string
	Get_AppId() syscall.GUID
	Get_DateOfManufacture() *IReference[DateTime]
	Get_DefaultLanguage() *ILanguage
	Get_DeviceId() string
	Get_HardwareVersion() string
	Get_ModelNumber() string
	Get_SoftwareVersion() string
	Get_SupportedLanguages() *IVectorView[*ILanguage]
	Get_SupportUrl() *IUriRuntimeClass
	Get_AppName() string
	Get_Description() string
	Get_DeviceName() string
	Get_Manufacturer() string
}

type IAllJoynAboutDataViewVtbl struct {
	win32.IInspectableVtbl
	Get_Status             uintptr
	Get_Properties         uintptr
	Get_AJSoftwareVersion  uintptr
	Get_AppId              uintptr
	Get_DateOfManufacture  uintptr
	Get_DefaultLanguage    uintptr
	Get_DeviceId           uintptr
	Get_HardwareVersion    uintptr
	Get_ModelNumber        uintptr
	Get_SoftwareVersion    uintptr
	Get_SupportedLanguages uintptr
	Get_SupportUrl         uintptr
	Get_AppName            uintptr
	Get_Description        uintptr
	Get_DeviceName         uintptr
	Get_Manufacturer       uintptr
}

type IAllJoynAboutDataView struct {
	win32.IInspectable
}

func (this *IAllJoynAboutDataView) Vtbl() *IAllJoynAboutDataViewVtbl {
	return (*IAllJoynAboutDataViewVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynAboutDataView) Get_Status() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynAboutDataView) Get_Properties() *IMapView[string, interface{}] {
	var _result *IMapView[string, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAllJoynAboutDataView) Get_AJSoftwareVersion() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AJSoftwareVersion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAllJoynAboutDataView) Get_AppId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynAboutDataView) Get_DateOfManufacture() *IReference[DateTime] {
	var _result *IReference[DateTime]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DateOfManufacture, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAllJoynAboutDataView) Get_DefaultLanguage() *ILanguage {
	var _result *ILanguage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DefaultLanguage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAllJoynAboutDataView) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAllJoynAboutDataView) Get_HardwareVersion() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HardwareVersion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAllJoynAboutDataView) Get_ModelNumber() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ModelNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAllJoynAboutDataView) Get_SoftwareVersion() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SoftwareVersion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAllJoynAboutDataView) Get_SupportedLanguages() *IVectorView[*ILanguage] {
	var _result *IVectorView[*ILanguage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedLanguages, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAllJoynAboutDataView) Get_SupportUrl() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportUrl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAllJoynAboutDataView) Get_AppName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAllJoynAboutDataView) Get_Description() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Description, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAllJoynAboutDataView) Get_DeviceName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAllJoynAboutDataView) Get_Manufacturer() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Manufacturer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 57EDB688-0C5E-416E-88B5-39B32D25C47D
var IID_IAllJoynAboutDataViewStatics = syscall.GUID{0x57EDB688, 0x0C5E, 0x416E,
	[8]byte{0x88, 0xB5, 0x39, 0xB3, 0x2D, 0x25, 0xC4, 0x7D}}

type IAllJoynAboutDataViewStaticsInterface interface {
	win32.IInspectableInterface
	GetDataBySessionPortAsync(uniqueName string, busAttachment *IAllJoynBusAttachment, sessionPort uint16) *IAsyncOperation[*IAllJoynAboutDataView]
	GetDataBySessionPortWithLanguageAsync(uniqueName string, busAttachment *IAllJoynBusAttachment, sessionPort uint16, language *ILanguage) *IAsyncOperation[*IAllJoynAboutDataView]
}

type IAllJoynAboutDataViewStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDataBySessionPortAsync             uintptr
	GetDataBySessionPortWithLanguageAsync uintptr
}

type IAllJoynAboutDataViewStatics struct {
	win32.IInspectable
}

func (this *IAllJoynAboutDataViewStatics) Vtbl() *IAllJoynAboutDataViewStaticsVtbl {
	return (*IAllJoynAboutDataViewStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynAboutDataViewStatics) GetDataBySessionPortAsync(uniqueName string, busAttachment *IAllJoynBusAttachment, sessionPort uint16) *IAsyncOperation[*IAllJoynAboutDataView] {
	var _result *IAsyncOperation[*IAllJoynAboutDataView]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDataBySessionPortAsync, uintptr(unsafe.Pointer(this)), NewHStr(uniqueName).Ptr, uintptr(unsafe.Pointer(busAttachment)), uintptr(sessionPort), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAllJoynAboutDataViewStatics) GetDataBySessionPortWithLanguageAsync(uniqueName string, busAttachment *IAllJoynBusAttachment, sessionPort uint16, language *ILanguage) *IAsyncOperation[*IAllJoynAboutDataView] {
	var _result *IAsyncOperation[*IAllJoynAboutDataView]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDataBySessionPortWithLanguageAsync, uintptr(unsafe.Pointer(this)), NewHStr(uniqueName).Ptr, uintptr(unsafe.Pointer(busAttachment)), uintptr(sessionPort), uintptr(unsafe.Pointer(language)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4DA817D2-CD1D-4023-A7C4-16DEF89C28DF
var IID_IAllJoynAcceptSessionJoiner = syscall.GUID{0x4DA817D2, 0xCD1D, 0x4023,
	[8]byte{0xA7, 0xC4, 0x16, 0xDE, 0xF8, 0x9C, 0x28, 0xDF}}

type IAllJoynAcceptSessionJoinerInterface interface {
	win32.IInspectableInterface
	Accept()
}

type IAllJoynAcceptSessionJoinerVtbl struct {
	win32.IInspectableVtbl
	Accept uintptr
}

type IAllJoynAcceptSessionJoiner struct {
	win32.IInspectable
}

func (this *IAllJoynAcceptSessionJoiner) Vtbl() *IAllJoynAcceptSessionJoinerVtbl {
	return (*IAllJoynAcceptSessionJoinerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynAcceptSessionJoiner) Accept() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Accept, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 4EFB5365-3E8A-4257-8F10-539CE0D56C0F
var IID_IAllJoynAcceptSessionJoinerEventArgs = syscall.GUID{0x4EFB5365, 0x3E8A, 0x4257,
	[8]byte{0x8F, 0x10, 0x53, 0x9C, 0xE0, 0xD5, 0x6C, 0x0F}}

type IAllJoynAcceptSessionJoinerEventArgsInterface interface {
	win32.IInspectableInterface
	Get_UniqueName() string
	Get_SessionPort() uint16
	Get_TrafficType() AllJoynTrafficType
	Get_SamePhysicalNode() bool
	Get_SameNetwork() bool
	Accept()
}

type IAllJoynAcceptSessionJoinerEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_UniqueName       uintptr
	Get_SessionPort      uintptr
	Get_TrafficType      uintptr
	Get_SamePhysicalNode uintptr
	Get_SameNetwork      uintptr
	Accept               uintptr
}

type IAllJoynAcceptSessionJoinerEventArgs struct {
	win32.IInspectable
}

func (this *IAllJoynAcceptSessionJoinerEventArgs) Vtbl() *IAllJoynAcceptSessionJoinerEventArgsVtbl {
	return (*IAllJoynAcceptSessionJoinerEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynAcceptSessionJoinerEventArgs) Get_UniqueName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UniqueName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAllJoynAcceptSessionJoinerEventArgs) Get_SessionPort() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SessionPort, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynAcceptSessionJoinerEventArgs) Get_TrafficType() AllJoynTrafficType {
	var _result AllJoynTrafficType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TrafficType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynAcceptSessionJoinerEventArgs) Get_SamePhysicalNode() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SamePhysicalNode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynAcceptSessionJoinerEventArgs) Get_SameNetwork() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SameNetwork, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynAcceptSessionJoinerEventArgs) Accept() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Accept, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// B4435BC0-6145-429E-84DB-D5BFE772B14F
var IID_IAllJoynAcceptSessionJoinerEventArgsFactory = syscall.GUID{0xB4435BC0, 0x6145, 0x429E,
	[8]byte{0x84, 0xDB, 0xD5, 0xBF, 0xE7, 0x72, 0xB1, 0x4F}}

type IAllJoynAcceptSessionJoinerEventArgsFactoryInterface interface {
	win32.IInspectableInterface
	Create(uniqueName string, sessionPort uint16, trafficType AllJoynTrafficType, proximity byte, acceptSessionJoiner *IAllJoynAcceptSessionJoiner) *IAllJoynAcceptSessionJoinerEventArgs
}

type IAllJoynAcceptSessionJoinerEventArgsFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IAllJoynAcceptSessionJoinerEventArgsFactory struct {
	win32.IInspectable
}

func (this *IAllJoynAcceptSessionJoinerEventArgsFactory) Vtbl() *IAllJoynAcceptSessionJoinerEventArgsFactoryVtbl {
	return (*IAllJoynAcceptSessionJoinerEventArgsFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynAcceptSessionJoinerEventArgsFactory) Create(uniqueName string, sessionPort uint16, trafficType AllJoynTrafficType, proximity byte, acceptSessionJoiner *IAllJoynAcceptSessionJoiner) *IAllJoynAcceptSessionJoinerEventArgs {
	var _result *IAllJoynAcceptSessionJoinerEventArgs
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(uniqueName).Ptr, uintptr(sessionPort), uintptr(trafficType), uintptr(proximity), uintptr(unsafe.Pointer(acceptSessionJoiner)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 97B4701C-15DC-4B53-B6A4-7D134300D7BF
var IID_IAllJoynAuthenticationCompleteEventArgs = syscall.GUID{0x97B4701C, 0x15DC, 0x4B53,
	[8]byte{0xB6, 0xA4, 0x7D, 0x13, 0x43, 0x00, 0xD7, 0xBF}}

type IAllJoynAuthenticationCompleteEventArgsInterface interface {
	win32.IInspectableInterface
	Get_AuthenticationMechanism() AllJoynAuthenticationMechanism
	Get_PeerUniqueName() string
	Get_Succeeded() bool
}

type IAllJoynAuthenticationCompleteEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_AuthenticationMechanism uintptr
	Get_PeerUniqueName          uintptr
	Get_Succeeded               uintptr
}

type IAllJoynAuthenticationCompleteEventArgs struct {
	win32.IInspectable
}

func (this *IAllJoynAuthenticationCompleteEventArgs) Vtbl() *IAllJoynAuthenticationCompleteEventArgsVtbl {
	return (*IAllJoynAuthenticationCompleteEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynAuthenticationCompleteEventArgs) Get_AuthenticationMechanism() AllJoynAuthenticationMechanism {
	var _result AllJoynAuthenticationMechanism
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AuthenticationMechanism, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynAuthenticationCompleteEventArgs) Get_PeerUniqueName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PeerUniqueName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAllJoynAuthenticationCompleteEventArgs) Get_Succeeded() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Succeeded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F309F153-1EED-42C3-A20E-436D41FE62F6
var IID_IAllJoynBusAttachment = syscall.GUID{0xF309F153, 0x1EED, 0x42C3,
	[8]byte{0xA2, 0x0E, 0x43, 0x6D, 0x41, 0xFE, 0x62, 0xF6}}

type IAllJoynBusAttachmentInterface interface {
	win32.IInspectableInterface
	Get_AboutData() *IAllJoynAboutData
	Get_ConnectionSpecification() string
	Get_State() AllJoynBusAttachmentState
	Get_UniqueName() string
	PingAsync(uniqueName string) *IAsyncOperation[int32]
	Connect()
	Disconnect()
	Add_StateChanged(handler TypedEventHandler[*IAllJoynBusAttachment, *IAllJoynBusAttachmentStateChangedEventArgs]) EventRegistrationToken
	Remove_StateChanged(token EventRegistrationToken)
	Get_AuthenticationMechanisms() *IVector[AllJoynAuthenticationMechanism]
	Add_CredentialsRequested(handler TypedEventHandler[*IAllJoynBusAttachment, *IAllJoynCredentialsRequestedEventArgs]) EventRegistrationToken
	Remove_CredentialsRequested(token EventRegistrationToken)
	Add_CredentialsVerificationRequested(handler TypedEventHandler[*IAllJoynBusAttachment, *IAllJoynCredentialsVerificationRequestedEventArgs]) EventRegistrationToken
	Remove_CredentialsVerificationRequested(token EventRegistrationToken)
	Add_AuthenticationComplete(handler TypedEventHandler[*IAllJoynBusAttachment, *IAllJoynAuthenticationCompleteEventArgs]) EventRegistrationToken
	Remove_AuthenticationComplete(token EventRegistrationToken)
}

type IAllJoynBusAttachmentVtbl struct {
	win32.IInspectableVtbl
	Get_AboutData                           uintptr
	Get_ConnectionSpecification             uintptr
	Get_State                               uintptr
	Get_UniqueName                          uintptr
	PingAsync                               uintptr
	Connect                                 uintptr
	Disconnect                              uintptr
	Add_StateChanged                        uintptr
	Remove_StateChanged                     uintptr
	Get_AuthenticationMechanisms            uintptr
	Add_CredentialsRequested                uintptr
	Remove_CredentialsRequested             uintptr
	Add_CredentialsVerificationRequested    uintptr
	Remove_CredentialsVerificationRequested uintptr
	Add_AuthenticationComplete              uintptr
	Remove_AuthenticationComplete           uintptr
}

type IAllJoynBusAttachment struct {
	win32.IInspectable
}

func (this *IAllJoynBusAttachment) Vtbl() *IAllJoynBusAttachmentVtbl {
	return (*IAllJoynBusAttachmentVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynBusAttachment) Get_AboutData() *IAllJoynAboutData {
	var _result *IAllJoynAboutData
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AboutData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAllJoynBusAttachment) Get_ConnectionSpecification() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConnectionSpecification, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAllJoynBusAttachment) Get_State() AllJoynBusAttachmentState {
	var _result AllJoynBusAttachmentState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynBusAttachment) Get_UniqueName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UniqueName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAllJoynBusAttachment) PingAsync(uniqueName string) *IAsyncOperation[int32] {
	var _result *IAsyncOperation[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PingAsync, uintptr(unsafe.Pointer(this)), NewHStr(uniqueName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAllJoynBusAttachment) Connect() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Connect, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IAllJoynBusAttachment) Disconnect() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Disconnect, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IAllJoynBusAttachment) Add_StateChanged(handler TypedEventHandler[*IAllJoynBusAttachment, *IAllJoynBusAttachmentStateChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynBusAttachment) Remove_StateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAllJoynBusAttachment) Get_AuthenticationMechanisms() *IVector[AllJoynAuthenticationMechanism] {
	var _result *IVector[AllJoynAuthenticationMechanism]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AuthenticationMechanisms, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAllJoynBusAttachment) Add_CredentialsRequested(handler TypedEventHandler[*IAllJoynBusAttachment, *IAllJoynCredentialsRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_CredentialsRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynBusAttachment) Remove_CredentialsRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_CredentialsRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAllJoynBusAttachment) Add_CredentialsVerificationRequested(handler TypedEventHandler[*IAllJoynBusAttachment, *IAllJoynCredentialsVerificationRequestedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_CredentialsVerificationRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynBusAttachment) Remove_CredentialsVerificationRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_CredentialsVerificationRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAllJoynBusAttachment) Add_AuthenticationComplete(handler TypedEventHandler[*IAllJoynBusAttachment, *IAllJoynAuthenticationCompleteEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AuthenticationComplete, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynBusAttachment) Remove_AuthenticationComplete(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AuthenticationComplete, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 3474CB1E-2368-43B2-B43E-6A3AC1278D98
var IID_IAllJoynBusAttachment2 = syscall.GUID{0x3474CB1E, 0x2368, 0x43B2,
	[8]byte{0xB4, 0x3E, 0x6A, 0x3A, 0xC1, 0x27, 0x8D, 0x98}}

type IAllJoynBusAttachment2Interface interface {
	win32.IInspectableInterface
	GetAboutDataAsync(serviceInfo *IAllJoynServiceInfo) *IAsyncOperation[*IAllJoynAboutDataView]
	GetAboutDataWithLanguageAsync(serviceInfo *IAllJoynServiceInfo, language *ILanguage) *IAsyncOperation[*IAllJoynAboutDataView]
	Add_AcceptSessionJoinerRequested(handler TypedEventHandler[*IAllJoynBusAttachment, *IAllJoynAcceptSessionJoinerEventArgs]) EventRegistrationToken
	Remove_AcceptSessionJoinerRequested(token EventRegistrationToken)
	Add_SessionJoined(handler TypedEventHandler[*IAllJoynBusAttachment, *IAllJoynSessionJoinedEventArgs]) EventRegistrationToken
	Remove_SessionJoined(token EventRegistrationToken)
}

type IAllJoynBusAttachment2Vtbl struct {
	win32.IInspectableVtbl
	GetAboutDataAsync                   uintptr
	GetAboutDataWithLanguageAsync       uintptr
	Add_AcceptSessionJoinerRequested    uintptr
	Remove_AcceptSessionJoinerRequested uintptr
	Add_SessionJoined                   uintptr
	Remove_SessionJoined                uintptr
}

type IAllJoynBusAttachment2 struct {
	win32.IInspectable
}

func (this *IAllJoynBusAttachment2) Vtbl() *IAllJoynBusAttachment2Vtbl {
	return (*IAllJoynBusAttachment2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynBusAttachment2) GetAboutDataAsync(serviceInfo *IAllJoynServiceInfo) *IAsyncOperation[*IAllJoynAboutDataView] {
	var _result *IAsyncOperation[*IAllJoynAboutDataView]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAboutDataAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(serviceInfo)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAllJoynBusAttachment2) GetAboutDataWithLanguageAsync(serviceInfo *IAllJoynServiceInfo, language *ILanguage) *IAsyncOperation[*IAllJoynAboutDataView] {
	var _result *IAsyncOperation[*IAllJoynAboutDataView]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAboutDataWithLanguageAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(serviceInfo)), uintptr(unsafe.Pointer(language)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAllJoynBusAttachment2) Add_AcceptSessionJoinerRequested(handler TypedEventHandler[*IAllJoynBusAttachment, *IAllJoynAcceptSessionJoinerEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AcceptSessionJoinerRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynBusAttachment2) Remove_AcceptSessionJoinerRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AcceptSessionJoinerRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAllJoynBusAttachment2) Add_SessionJoined(handler TypedEventHandler[*IAllJoynBusAttachment, *IAllJoynSessionJoinedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SessionJoined, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynBusAttachment2) Remove_SessionJoined(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SessionJoined, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 642EF1A4-AD85-4DDF-90AE-604452B22288
var IID_IAllJoynBusAttachmentFactory = syscall.GUID{0x642EF1A4, 0xAD85, 0x4DDF,
	[8]byte{0x90, 0xAE, 0x60, 0x44, 0x52, 0xB2, 0x22, 0x88}}

type IAllJoynBusAttachmentFactoryInterface interface {
	win32.IInspectableInterface
	Create(connectionSpecification string) *IAllJoynBusAttachment
}

type IAllJoynBusAttachmentFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IAllJoynBusAttachmentFactory struct {
	win32.IInspectable
}

func (this *IAllJoynBusAttachmentFactory) Vtbl() *IAllJoynBusAttachmentFactoryVtbl {
	return (*IAllJoynBusAttachmentFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynBusAttachmentFactory) Create(connectionSpecification string) *IAllJoynBusAttachment {
	var _result *IAllJoynBusAttachment
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(connectionSpecification).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D82E75F4-C02A-41EC-A8D5-EAB1558953AA
var IID_IAllJoynBusAttachmentStateChangedEventArgs = syscall.GUID{0xD82E75F4, 0xC02A, 0x41EC,
	[8]byte{0xA8, 0xD5, 0xEA, 0xB1, 0x55, 0x89, 0x53, 0xAA}}

type IAllJoynBusAttachmentStateChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_State() AllJoynBusAttachmentState
	Get_Status() int32
}

type IAllJoynBusAttachmentStateChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_State  uintptr
	Get_Status uintptr
}

type IAllJoynBusAttachmentStateChangedEventArgs struct {
	win32.IInspectable
}

func (this *IAllJoynBusAttachmentStateChangedEventArgs) Vtbl() *IAllJoynBusAttachmentStateChangedEventArgsVtbl {
	return (*IAllJoynBusAttachmentStateChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynBusAttachmentStateChangedEventArgs) Get_State() AllJoynBusAttachmentState {
	var _result AllJoynBusAttachmentState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynBusAttachmentStateChangedEventArgs) Get_Status() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 839D4D3D-1051-40D7-872A-8D0141115B1F
var IID_IAllJoynBusAttachmentStatics = syscall.GUID{0x839D4D3D, 0x1051, 0x40D7,
	[8]byte{0x87, 0x2A, 0x8D, 0x01, 0x41, 0x11, 0x5B, 0x1F}}

type IAllJoynBusAttachmentStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *IAllJoynBusAttachment
	GetWatcher(requiredInterfaces *IIterable[string]) *IDeviceWatcher
}

type IAllJoynBusAttachmentStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
	GetWatcher uintptr
}

type IAllJoynBusAttachmentStatics struct {
	win32.IInspectable
}

func (this *IAllJoynBusAttachmentStatics) Vtbl() *IAllJoynBusAttachmentStaticsVtbl {
	return (*IAllJoynBusAttachmentStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynBusAttachmentStatics) GetDefault() *IAllJoynBusAttachment {
	var _result *IAllJoynBusAttachment
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAllJoynBusAttachmentStatics) GetWatcher(requiredInterfaces *IIterable[string]) *IDeviceWatcher {
	var _result *IDeviceWatcher
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetWatcher, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(requiredInterfaces)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E8FD825E-F73A-490C-8804-04E026643047
var IID_IAllJoynBusObject = syscall.GUID{0xE8FD825E, 0xF73A, 0x490C,
	[8]byte{0x88, 0x04, 0x04, 0xE0, 0x26, 0x64, 0x30, 0x47}}

type IAllJoynBusObjectInterface interface {
	win32.IInspectableInterface
	Start()
	Stop()
	AddProducer(producer *IAllJoynProducer)
	Get_BusAttachment() *IAllJoynBusAttachment
	Get_Session() *IAllJoynSession
	Add_Stopped(handler TypedEventHandler[*IAllJoynBusObject, *IAllJoynBusObjectStoppedEventArgs]) EventRegistrationToken
	Remove_Stopped(token EventRegistrationToken)
}

type IAllJoynBusObjectVtbl struct {
	win32.IInspectableVtbl
	Start             uintptr
	Stop              uintptr
	AddProducer       uintptr
	Get_BusAttachment uintptr
	Get_Session       uintptr
	Add_Stopped       uintptr
	Remove_Stopped    uintptr
}

type IAllJoynBusObject struct {
	win32.IInspectable
}

func (this *IAllJoynBusObject) Vtbl() *IAllJoynBusObjectVtbl {
	return (*IAllJoynBusObjectVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynBusObject) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IAllJoynBusObject) Stop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IAllJoynBusObject) AddProducer(producer *IAllJoynProducer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddProducer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(producer)))
	_ = _hr
}

func (this *IAllJoynBusObject) Get_BusAttachment() *IAllJoynBusAttachment {
	var _result *IAllJoynBusAttachment
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BusAttachment, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAllJoynBusObject) Get_Session() *IAllJoynSession {
	var _result *IAllJoynSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Session, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAllJoynBusObject) Add_Stopped(handler TypedEventHandler[*IAllJoynBusObject, *IAllJoynBusObjectStoppedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Stopped, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynBusObject) Remove_Stopped(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Stopped, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 2C2F9F0B-8E02-4F9C-AC27-EA6DAD5D3B50
var IID_IAllJoynBusObjectFactory = syscall.GUID{0x2C2F9F0B, 0x8E02, 0x4F9C,
	[8]byte{0xAC, 0x27, 0xEA, 0x6D, 0xAD, 0x5D, 0x3B, 0x50}}

type IAllJoynBusObjectFactoryInterface interface {
	win32.IInspectableInterface
	Create(objectPath string) *IAllJoynBusObject
	CreateWithBusAttachment(objectPath string, busAttachment *IAllJoynBusAttachment) *IAllJoynBusObject
}

type IAllJoynBusObjectFactoryVtbl struct {
	win32.IInspectableVtbl
	Create                  uintptr
	CreateWithBusAttachment uintptr
}

type IAllJoynBusObjectFactory struct {
	win32.IInspectable
}

func (this *IAllJoynBusObjectFactory) Vtbl() *IAllJoynBusObjectFactoryVtbl {
	return (*IAllJoynBusObjectFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynBusObjectFactory) Create(objectPath string) *IAllJoynBusObject {
	var _result *IAllJoynBusObject
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(objectPath).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAllJoynBusObjectFactory) CreateWithBusAttachment(objectPath string, busAttachment *IAllJoynBusAttachment) *IAllJoynBusObject {
	var _result *IAllJoynBusObject
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithBusAttachment, uintptr(unsafe.Pointer(this)), NewHStr(objectPath).Ptr, uintptr(unsafe.Pointer(busAttachment)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DE102115-EF8E-4D42-B93B-A2AE74519766
var IID_IAllJoynBusObjectStoppedEventArgs = syscall.GUID{0xDE102115, 0xEF8E, 0x4D42,
	[8]byte{0xB9, 0x3B, 0xA2, 0xAE, 0x74, 0x51, 0x97, 0x66}}

type IAllJoynBusObjectStoppedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Status() int32
}

type IAllJoynBusObjectStoppedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
}

type IAllJoynBusObjectStoppedEventArgs struct {
	win32.IInspectable
}

func (this *IAllJoynBusObjectStoppedEventArgs) Vtbl() *IAllJoynBusObjectStoppedEventArgsVtbl {
	return (*IAllJoynBusObjectStoppedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynBusObjectStoppedEventArgs) Get_Status() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 6B22FD48-D0A3-4255-953A-4772B4028073
var IID_IAllJoynBusObjectStoppedEventArgsFactory = syscall.GUID{0x6B22FD48, 0xD0A3, 0x4255,
	[8]byte{0x95, 0x3A, 0x47, 0x72, 0xB4, 0x02, 0x80, 0x73}}

type IAllJoynBusObjectStoppedEventArgsFactoryInterface interface {
	win32.IInspectableInterface
	Create(status int32) *IAllJoynBusObjectStoppedEventArgs
}

type IAllJoynBusObjectStoppedEventArgsFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IAllJoynBusObjectStoppedEventArgsFactory struct {
	win32.IInspectable
}

func (this *IAllJoynBusObjectStoppedEventArgsFactory) Vtbl() *IAllJoynBusObjectStoppedEventArgsFactoryVtbl {
	return (*IAllJoynBusObjectStoppedEventArgsFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynBusObjectStoppedEventArgsFactory) Create(status int32) *IAllJoynBusObjectStoppedEventArgs {
	var _result *IAllJoynBusObjectStoppedEventArgs
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(status), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 824650F2-A190-40B1-ABAB-349EC244DFAA
var IID_IAllJoynCredentials = syscall.GUID{0x824650F2, 0xA190, 0x40B1,
	[8]byte{0xAB, 0xAB, 0x34, 0x9E, 0xC2, 0x44, 0xDF, 0xAA}}

type IAllJoynCredentialsInterface interface {
	win32.IInspectableInterface
	Get_AuthenticationMechanism() AllJoynAuthenticationMechanism
	Get_Certificate() *ICertificate
	Put_Certificate(value *ICertificate)
	Get_PasswordCredential() *IPasswordCredential
	Put_PasswordCredential(value *IPasswordCredential)
	Get_Timeout() TimeSpan
	Put_Timeout(value TimeSpan)
}

type IAllJoynCredentialsVtbl struct {
	win32.IInspectableVtbl
	Get_AuthenticationMechanism uintptr
	Get_Certificate             uintptr
	Put_Certificate             uintptr
	Get_PasswordCredential      uintptr
	Put_PasswordCredential      uintptr
	Get_Timeout                 uintptr
	Put_Timeout                 uintptr
}

type IAllJoynCredentials struct {
	win32.IInspectable
}

func (this *IAllJoynCredentials) Vtbl() *IAllJoynCredentialsVtbl {
	return (*IAllJoynCredentialsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynCredentials) Get_AuthenticationMechanism() AllJoynAuthenticationMechanism {
	var _result AllJoynAuthenticationMechanism
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AuthenticationMechanism, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynCredentials) Get_Certificate() *ICertificate {
	var _result *ICertificate
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Certificate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAllJoynCredentials) Put_Certificate(value *ICertificate) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Certificate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAllJoynCredentials) Get_PasswordCredential() *IPasswordCredential {
	var _result *IPasswordCredential
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PasswordCredential, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAllJoynCredentials) Put_PasswordCredential(value *IPasswordCredential) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PasswordCredential, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAllJoynCredentials) Get_Timeout() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timeout, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynCredentials) Put_Timeout(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Timeout, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

// 6A87E34E-B069-4B80-9E1A-41BC837C65D2
var IID_IAllJoynCredentialsRequestedEventArgs = syscall.GUID{0x6A87E34E, 0xB069, 0x4B80,
	[8]byte{0x9E, 0x1A, 0x41, 0xBC, 0x83, 0x7C, 0x65, 0xD2}}

type IAllJoynCredentialsRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_AttemptCount() uint16
	Get_Credentials() *IAllJoynCredentials
	Get_PeerUniqueName() string
	Get_RequestedUserName() string
	GetDeferral() *IDeferral
}

type IAllJoynCredentialsRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_AttemptCount      uintptr
	Get_Credentials       uintptr
	Get_PeerUniqueName    uintptr
	Get_RequestedUserName uintptr
	GetDeferral           uintptr
}

type IAllJoynCredentialsRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IAllJoynCredentialsRequestedEventArgs) Vtbl() *IAllJoynCredentialsRequestedEventArgsVtbl {
	return (*IAllJoynCredentialsRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynCredentialsRequestedEventArgs) Get_AttemptCount() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AttemptCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynCredentialsRequestedEventArgs) Get_Credentials() *IAllJoynCredentials {
	var _result *IAllJoynCredentials
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Credentials, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAllJoynCredentialsRequestedEventArgs) Get_PeerUniqueName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PeerUniqueName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAllJoynCredentialsRequestedEventArgs) Get_RequestedUserName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestedUserName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAllJoynCredentialsRequestedEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 800A7612-B805-44AF-A2E1-792AB655A2D0
var IID_IAllJoynCredentialsVerificationRequestedEventArgs = syscall.GUID{0x800A7612, 0xB805, 0x44AF,
	[8]byte{0xA2, 0xE1, 0x79, 0x2A, 0xB6, 0x55, 0xA2, 0xD0}}

type IAllJoynCredentialsVerificationRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_AuthenticationMechanism() AllJoynAuthenticationMechanism
	Get_PeerUniqueName() string
	Get_PeerCertificate() *ICertificate
	Get_PeerCertificateErrorSeverity() SocketSslErrorSeverity
	Get_PeerCertificateErrors() *IVectorView[ChainValidationResult]
	Get_PeerIntermediateCertificates() *IVectorView[*ICertificate]
	Accept()
	GetDeferral() *IDeferral
}

type IAllJoynCredentialsVerificationRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_AuthenticationMechanism      uintptr
	Get_PeerUniqueName               uintptr
	Get_PeerCertificate              uintptr
	Get_PeerCertificateErrorSeverity uintptr
	Get_PeerCertificateErrors        uintptr
	Get_PeerIntermediateCertificates uintptr
	Accept                           uintptr
	GetDeferral                      uintptr
}

type IAllJoynCredentialsVerificationRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IAllJoynCredentialsVerificationRequestedEventArgs) Vtbl() *IAllJoynCredentialsVerificationRequestedEventArgsVtbl {
	return (*IAllJoynCredentialsVerificationRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynCredentialsVerificationRequestedEventArgs) Get_AuthenticationMechanism() AllJoynAuthenticationMechanism {
	var _result AllJoynAuthenticationMechanism
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AuthenticationMechanism, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynCredentialsVerificationRequestedEventArgs) Get_PeerUniqueName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PeerUniqueName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAllJoynCredentialsVerificationRequestedEventArgs) Get_PeerCertificate() *ICertificate {
	var _result *ICertificate
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PeerCertificate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAllJoynCredentialsVerificationRequestedEventArgs) Get_PeerCertificateErrorSeverity() SocketSslErrorSeverity {
	var _result SocketSslErrorSeverity
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PeerCertificateErrorSeverity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynCredentialsVerificationRequestedEventArgs) Get_PeerCertificateErrors() *IVectorView[ChainValidationResult] {
	var _result *IVectorView[ChainValidationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PeerCertificateErrors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAllJoynCredentialsVerificationRequestedEventArgs) Get_PeerIntermediateCertificates() *IVectorView[*ICertificate] {
	var _result *IVectorView[*ICertificate]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PeerIntermediateCertificates, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAllJoynCredentialsVerificationRequestedEventArgs) Accept() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Accept, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IAllJoynCredentialsVerificationRequestedEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FF2B0127-2C12-4859-AA3A-C74461EE814C
var IID_IAllJoynMessageInfo = syscall.GUID{0xFF2B0127, 0x2C12, 0x4859,
	[8]byte{0xAA, 0x3A, 0xC7, 0x44, 0x61, 0xEE, 0x81, 0x4C}}

type IAllJoynMessageInfoInterface interface {
	win32.IInspectableInterface
	Get_SenderUniqueName() string
}

type IAllJoynMessageInfoVtbl struct {
	win32.IInspectableVtbl
	Get_SenderUniqueName uintptr
}

type IAllJoynMessageInfo struct {
	win32.IInspectable
}

func (this *IAllJoynMessageInfo) Vtbl() *IAllJoynMessageInfoVtbl {
	return (*IAllJoynMessageInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynMessageInfo) Get_SenderUniqueName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SenderUniqueName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 34664C2A-8289-43D4-B4A8-3F4DE359F043
var IID_IAllJoynMessageInfoFactory = syscall.GUID{0x34664C2A, 0x8289, 0x43D4,
	[8]byte{0xB4, 0xA8, 0x3F, 0x4D, 0xE3, 0x59, 0xF0, 0x43}}

type IAllJoynMessageInfoFactoryInterface interface {
	win32.IInspectableInterface
	Create(senderUniqueName string) *IAllJoynMessageInfo
}

type IAllJoynMessageInfoFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IAllJoynMessageInfoFactory struct {
	win32.IInspectable
}

func (this *IAllJoynMessageInfoFactory) Vtbl() *IAllJoynMessageInfoFactoryVtbl {
	return (*IAllJoynMessageInfoFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynMessageInfoFactory) Create(senderUniqueName string) *IAllJoynMessageInfo {
	var _result *IAllJoynMessageInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(senderUniqueName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9D084679-469B-495A-A710-AC50F123069F
var IID_IAllJoynProducer = syscall.GUID{0x9D084679, 0x469B, 0x495A,
	[8]byte{0xA7, 0x10, 0xAC, 0x50, 0xF1, 0x23, 0x06, 0x9F}}

type IAllJoynProducerInterface interface {
	win32.IInspectableInterface
	SetBusObject(busObject *IAllJoynBusObject)
}

type IAllJoynProducerVtbl struct {
	win32.IInspectableVtbl
	SetBusObject uintptr
}

type IAllJoynProducer struct {
	win32.IInspectable
}

func (this *IAllJoynProducer) Vtbl() *IAllJoynProducerVtbl {
	return (*IAllJoynProducerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynProducer) SetBusObject(busObject *IAllJoynBusObject) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetBusObject, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(busObject)))
	_ = _hr
}

// 51309770-4937-492D-8080-236439987CEB
var IID_IAllJoynProducerStoppedEventArgs = syscall.GUID{0x51309770, 0x4937, 0x492D,
	[8]byte{0x80, 0x80, 0x23, 0x64, 0x39, 0x98, 0x7C, 0xEB}}

type IAllJoynProducerStoppedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Status() int32
}

type IAllJoynProducerStoppedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
}

type IAllJoynProducerStoppedEventArgs struct {
	win32.IInspectable
}

func (this *IAllJoynProducerStoppedEventArgs) Vtbl() *IAllJoynProducerStoppedEventArgsVtbl {
	return (*IAllJoynProducerStoppedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynProducerStoppedEventArgs) Get_Status() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 56529961-B219-4D6E-9F78-FA3F99FA8FE5
var IID_IAllJoynProducerStoppedEventArgsFactory = syscall.GUID{0x56529961, 0xB219, 0x4D6E,
	[8]byte{0x9F, 0x78, 0xFA, 0x3F, 0x99, 0xFA, 0x8F, 0xE5}}

type IAllJoynProducerStoppedEventArgsFactoryInterface interface {
	win32.IInspectableInterface
	Create(status int32) *IAllJoynProducerStoppedEventArgs
}

type IAllJoynProducerStoppedEventArgsFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IAllJoynProducerStoppedEventArgsFactory struct {
	win32.IInspectable
}

func (this *IAllJoynProducerStoppedEventArgsFactory) Vtbl() *IAllJoynProducerStoppedEventArgsFactoryVtbl {
	return (*IAllJoynProducerStoppedEventArgsFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynProducerStoppedEventArgsFactory) Create(status int32) *IAllJoynProducerStoppedEventArgs {
	var _result *IAllJoynProducerStoppedEventArgs
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(status), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4CBE8209-B93E-4182-999B-DDD000F9C575
var IID_IAllJoynServiceInfo = syscall.GUID{0x4CBE8209, 0xB93E, 0x4182,
	[8]byte{0x99, 0x9B, 0xDD, 0xD0, 0x00, 0xF9, 0xC5, 0x75}}

type IAllJoynServiceInfoInterface interface {
	win32.IInspectableInterface
	Get_UniqueName() string
	Get_ObjectPath() string
	Get_SessionPort() uint16
}

type IAllJoynServiceInfoVtbl struct {
	win32.IInspectableVtbl
	Get_UniqueName  uintptr
	Get_ObjectPath  uintptr
	Get_SessionPort uintptr
}

type IAllJoynServiceInfo struct {
	win32.IInspectable
}

func (this *IAllJoynServiceInfo) Vtbl() *IAllJoynServiceInfoVtbl {
	return (*IAllJoynServiceInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynServiceInfo) Get_UniqueName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UniqueName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAllJoynServiceInfo) Get_ObjectPath() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ObjectPath, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAllJoynServiceInfo) Get_SessionPort() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SessionPort, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7581DABD-FE03-4F4B-94A4-F02FDCBD11B8
var IID_IAllJoynServiceInfoFactory = syscall.GUID{0x7581DABD, 0xFE03, 0x4F4B,
	[8]byte{0x94, 0xA4, 0xF0, 0x2F, 0xDC, 0xBD, 0x11, 0xB8}}

type IAllJoynServiceInfoFactoryInterface interface {
	win32.IInspectableInterface
	Create(uniqueName string, objectPath string, sessionPort uint16) *IAllJoynServiceInfo
}

type IAllJoynServiceInfoFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IAllJoynServiceInfoFactory struct {
	win32.IInspectable
}

func (this *IAllJoynServiceInfoFactory) Vtbl() *IAllJoynServiceInfoFactoryVtbl {
	return (*IAllJoynServiceInfoFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynServiceInfoFactory) Create(uniqueName string, objectPath string, sessionPort uint16) *IAllJoynServiceInfo {
	var _result *IAllJoynServiceInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(uniqueName).Ptr, NewHStr(objectPath).Ptr, uintptr(sessionPort), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3057A95F-1D3F-41F3-8969-E32792627396
var IID_IAllJoynServiceInfoRemovedEventArgs = syscall.GUID{0x3057A95F, 0x1D3F, 0x41F3,
	[8]byte{0x89, 0x69, 0xE3, 0x27, 0x92, 0x62, 0x73, 0x96}}

type IAllJoynServiceInfoRemovedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_UniqueName() string
}

type IAllJoynServiceInfoRemovedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_UniqueName uintptr
}

type IAllJoynServiceInfoRemovedEventArgs struct {
	win32.IInspectable
}

func (this *IAllJoynServiceInfoRemovedEventArgs) Vtbl() *IAllJoynServiceInfoRemovedEventArgsVtbl {
	return (*IAllJoynServiceInfoRemovedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynServiceInfoRemovedEventArgs) Get_UniqueName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UniqueName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 0DBF8627-9AFF-4955-9227-6953BAF41569
var IID_IAllJoynServiceInfoRemovedEventArgsFactory = syscall.GUID{0x0DBF8627, 0x9AFF, 0x4955,
	[8]byte{0x92, 0x27, 0x69, 0x53, 0xBA, 0xF4, 0x15, 0x69}}

type IAllJoynServiceInfoRemovedEventArgsFactoryInterface interface {
	win32.IInspectableInterface
	Create(uniqueName string) *IAllJoynServiceInfoRemovedEventArgs
}

type IAllJoynServiceInfoRemovedEventArgsFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IAllJoynServiceInfoRemovedEventArgsFactory struct {
	win32.IInspectable
}

func (this *IAllJoynServiceInfoRemovedEventArgsFactory) Vtbl() *IAllJoynServiceInfoRemovedEventArgsFactoryVtbl {
	return (*IAllJoynServiceInfoRemovedEventArgsFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynServiceInfoRemovedEventArgsFactory) Create(uniqueName string) *IAllJoynServiceInfoRemovedEventArgs {
	var _result *IAllJoynServiceInfoRemovedEventArgs
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(uniqueName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5678570A-603A-49FC-B750-0EF13609213C
var IID_IAllJoynServiceInfoStatics = syscall.GUID{0x5678570A, 0x603A, 0x49FC,
	[8]byte{0xB7, 0x50, 0x0E, 0xF1, 0x36, 0x09, 0x21, 0x3C}}

type IAllJoynServiceInfoStaticsInterface interface {
	win32.IInspectableInterface
	FromIdAsync(deviceId string) *IAsyncOperation[*IAllJoynServiceInfo]
}

type IAllJoynServiceInfoStaticsVtbl struct {
	win32.IInspectableVtbl
	FromIdAsync uintptr
}

type IAllJoynServiceInfoStatics struct {
	win32.IInspectable
}

func (this *IAllJoynServiceInfoStatics) Vtbl() *IAllJoynServiceInfoStaticsVtbl {
	return (*IAllJoynServiceInfoStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynServiceInfoStatics) FromIdAsync(deviceId string) *IAsyncOperation[*IAllJoynServiceInfo] {
	var _result *IAsyncOperation[*IAllJoynServiceInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E8D11B0C-C0D4-406C-88A9-A93EFA85D4B1
var IID_IAllJoynSession = syscall.GUID{0xE8D11B0C, 0xC0D4, 0x406C,
	[8]byte{0x88, 0xA9, 0xA9, 0x3E, 0xFA, 0x85, 0xD4, 0xB1}}

type IAllJoynSessionInterface interface {
	win32.IInspectableInterface
	Get_Id() int32
	Get_Status() int32
	RemoveMemberAsync(uniqueName string) *IAsyncOperation[int32]
	Add_MemberAdded(handler TypedEventHandler[*IAllJoynSession, *IAllJoynSessionMemberAddedEventArgs]) EventRegistrationToken
	Remove_MemberAdded(token EventRegistrationToken)
	Add_MemberRemoved(handler TypedEventHandler[*IAllJoynSession, *IAllJoynSessionMemberRemovedEventArgs]) EventRegistrationToken
	Remove_MemberRemoved(token EventRegistrationToken)
	Add_Lost(handler TypedEventHandler[*IAllJoynSession, *IAllJoynSessionLostEventArgs]) EventRegistrationToken
	Remove_Lost(token EventRegistrationToken)
}

type IAllJoynSessionVtbl struct {
	win32.IInspectableVtbl
	Get_Id               uintptr
	Get_Status           uintptr
	RemoveMemberAsync    uintptr
	Add_MemberAdded      uintptr
	Remove_MemberAdded   uintptr
	Add_MemberRemoved    uintptr
	Remove_MemberRemoved uintptr
	Add_Lost             uintptr
	Remove_Lost          uintptr
}

type IAllJoynSession struct {
	win32.IInspectable
}

func (this *IAllJoynSession) Vtbl() *IAllJoynSessionVtbl {
	return (*IAllJoynSessionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynSession) Get_Id() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynSession) Get_Status() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynSession) RemoveMemberAsync(uniqueName string) *IAsyncOperation[int32] {
	var _result *IAsyncOperation[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveMemberAsync, uintptr(unsafe.Pointer(this)), NewHStr(uniqueName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAllJoynSession) Add_MemberAdded(handler TypedEventHandler[*IAllJoynSession, *IAllJoynSessionMemberAddedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_MemberAdded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynSession) Remove_MemberAdded(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_MemberAdded, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAllJoynSession) Add_MemberRemoved(handler TypedEventHandler[*IAllJoynSession, *IAllJoynSessionMemberRemovedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_MemberRemoved, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynSession) Remove_MemberRemoved(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_MemberRemoved, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAllJoynSession) Add_Lost(handler TypedEventHandler[*IAllJoynSession, *IAllJoynSessionLostEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Lost, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynSession) Remove_Lost(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Lost, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 9E9F5BD0-B5D7-47C5-8DAB-B040CC192871
var IID_IAllJoynSessionJoinedEventArgs = syscall.GUID{0x9E9F5BD0, 0xB5D7, 0x47C5,
	[8]byte{0x8D, 0xAB, 0xB0, 0x40, 0xCC, 0x19, 0x28, 0x71}}

type IAllJoynSessionJoinedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Session() *IAllJoynSession
}

type IAllJoynSessionJoinedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Session uintptr
}

type IAllJoynSessionJoinedEventArgs struct {
	win32.IInspectable
}

func (this *IAllJoynSessionJoinedEventArgs) Vtbl() *IAllJoynSessionJoinedEventArgsVtbl {
	return (*IAllJoynSessionJoinedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynSessionJoinedEventArgs) Get_Session() *IAllJoynSession {
	var _result *IAllJoynSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Session, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6824D689-D6CB-4D9E-A09E-35806870B17F
var IID_IAllJoynSessionJoinedEventArgsFactory = syscall.GUID{0x6824D689, 0xD6CB, 0x4D9E,
	[8]byte{0xA0, 0x9E, 0x35, 0x80, 0x68, 0x70, 0xB1, 0x7F}}

type IAllJoynSessionJoinedEventArgsFactoryInterface interface {
	win32.IInspectableInterface
	Create(session *IAllJoynSession) *IAllJoynSessionJoinedEventArgs
}

type IAllJoynSessionJoinedEventArgsFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IAllJoynSessionJoinedEventArgsFactory struct {
	win32.IInspectable
}

func (this *IAllJoynSessionJoinedEventArgsFactory) Vtbl() *IAllJoynSessionJoinedEventArgsFactoryVtbl {
	return (*IAllJoynSessionJoinedEventArgsFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynSessionJoinedEventArgsFactory) Create(session *IAllJoynSession) *IAllJoynSessionJoinedEventArgs {
	var _result *IAllJoynSessionJoinedEventArgs
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(session)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E766A48A-8BB8-4954-AE67-D2FA43D1F96B
var IID_IAllJoynSessionLostEventArgs = syscall.GUID{0xE766A48A, 0x8BB8, 0x4954,
	[8]byte{0xAE, 0x67, 0xD2, 0xFA, 0x43, 0xD1, 0xF9, 0x6B}}

type IAllJoynSessionLostEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Reason() AllJoynSessionLostReason
}

type IAllJoynSessionLostEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Reason uintptr
}

type IAllJoynSessionLostEventArgs struct {
	win32.IInspectable
}

func (this *IAllJoynSessionLostEventArgs) Vtbl() *IAllJoynSessionLostEventArgsVtbl {
	return (*IAllJoynSessionLostEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynSessionLostEventArgs) Get_Reason() AllJoynSessionLostReason {
	var _result AllJoynSessionLostReason
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Reason, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 13BBFD32-D2F4-49C9-980E-2805E13586B1
var IID_IAllJoynSessionLostEventArgsFactory = syscall.GUID{0x13BBFD32, 0xD2F4, 0x49C9,
	[8]byte{0x98, 0x0E, 0x28, 0x05, 0xE1, 0x35, 0x86, 0xB1}}

type IAllJoynSessionLostEventArgsFactoryInterface interface {
	win32.IInspectableInterface
	Create(reason AllJoynSessionLostReason) *IAllJoynSessionLostEventArgs
}

type IAllJoynSessionLostEventArgsFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IAllJoynSessionLostEventArgsFactory struct {
	win32.IInspectable
}

func (this *IAllJoynSessionLostEventArgsFactory) Vtbl() *IAllJoynSessionLostEventArgsFactoryVtbl {
	return (*IAllJoynSessionLostEventArgsFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynSessionLostEventArgsFactory) Create(reason AllJoynSessionLostReason) *IAllJoynSessionLostEventArgs {
	var _result *IAllJoynSessionLostEventArgs
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(reason), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 49A2798A-0DD1-46C1-9CD6-27190E503A5E
var IID_IAllJoynSessionMemberAddedEventArgs = syscall.GUID{0x49A2798A, 0x0DD1, 0x46C1,
	[8]byte{0x9C, 0xD6, 0x27, 0x19, 0x0E, 0x50, 0x3A, 0x5E}}

type IAllJoynSessionMemberAddedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_UniqueName() string
}

type IAllJoynSessionMemberAddedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_UniqueName uintptr
}

type IAllJoynSessionMemberAddedEventArgs struct {
	win32.IInspectable
}

func (this *IAllJoynSessionMemberAddedEventArgs) Vtbl() *IAllJoynSessionMemberAddedEventArgsVtbl {
	return (*IAllJoynSessionMemberAddedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynSessionMemberAddedEventArgs) Get_UniqueName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UniqueName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 341DE352-1D33-40A1-A1D3-E5777020E1F1
var IID_IAllJoynSessionMemberAddedEventArgsFactory = syscall.GUID{0x341DE352, 0x1D33, 0x40A1,
	[8]byte{0xA1, 0xD3, 0xE5, 0x77, 0x70, 0x20, 0xE1, 0xF1}}

type IAllJoynSessionMemberAddedEventArgsFactoryInterface interface {
	win32.IInspectableInterface
	Create(uniqueName string) *IAllJoynSessionMemberAddedEventArgs
}

type IAllJoynSessionMemberAddedEventArgsFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IAllJoynSessionMemberAddedEventArgsFactory struct {
	win32.IInspectable
}

func (this *IAllJoynSessionMemberAddedEventArgsFactory) Vtbl() *IAllJoynSessionMemberAddedEventArgsFactoryVtbl {
	return (*IAllJoynSessionMemberAddedEventArgsFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynSessionMemberAddedEventArgsFactory) Create(uniqueName string) *IAllJoynSessionMemberAddedEventArgs {
	var _result *IAllJoynSessionMemberAddedEventArgs
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(uniqueName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 409A219F-AA4A-4893-B430-BAA1B63C6219
var IID_IAllJoynSessionMemberRemovedEventArgs = syscall.GUID{0x409A219F, 0xAA4A, 0x4893,
	[8]byte{0xB4, 0x30, 0xBA, 0xA1, 0xB6, 0x3C, 0x62, 0x19}}

type IAllJoynSessionMemberRemovedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_UniqueName() string
}

type IAllJoynSessionMemberRemovedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_UniqueName uintptr
}

type IAllJoynSessionMemberRemovedEventArgs struct {
	win32.IInspectable
}

func (this *IAllJoynSessionMemberRemovedEventArgs) Vtbl() *IAllJoynSessionMemberRemovedEventArgsVtbl {
	return (*IAllJoynSessionMemberRemovedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynSessionMemberRemovedEventArgs) Get_UniqueName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UniqueName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// C4D355E8-42B8-4B67-B757-D0CFCAD59280
var IID_IAllJoynSessionMemberRemovedEventArgsFactory = syscall.GUID{0xC4D355E8, 0x42B8, 0x4B67,
	[8]byte{0xB7, 0x57, 0xD0, 0xCF, 0xCA, 0xD5, 0x92, 0x80}}

type IAllJoynSessionMemberRemovedEventArgsFactoryInterface interface {
	win32.IInspectableInterface
	Create(uniqueName string) *IAllJoynSessionMemberRemovedEventArgs
}

type IAllJoynSessionMemberRemovedEventArgsFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IAllJoynSessionMemberRemovedEventArgsFactory struct {
	win32.IInspectable
}

func (this *IAllJoynSessionMemberRemovedEventArgsFactory) Vtbl() *IAllJoynSessionMemberRemovedEventArgsFactoryVtbl {
	return (*IAllJoynSessionMemberRemovedEventArgsFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynSessionMemberRemovedEventArgsFactory) Create(uniqueName string) *IAllJoynSessionMemberRemovedEventArgs {
	var _result *IAllJoynSessionMemberRemovedEventArgs
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(uniqueName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9E05D604-A06C-46D4-B46C-0B0B54105B44
var IID_IAllJoynSessionStatics = syscall.GUID{0x9E05D604, 0xA06C, 0x46D4,
	[8]byte{0xB4, 0x6C, 0x0B, 0x0B, 0x54, 0x10, 0x5B, 0x44}}

type IAllJoynSessionStaticsInterface interface {
	win32.IInspectableInterface
	GetFromServiceInfoAsync(serviceInfo *IAllJoynServiceInfo) *IAsyncOperation[*IAllJoynSession]
	GetFromServiceInfoAndBusAttachmentAsync(serviceInfo *IAllJoynServiceInfo, busAttachment *IAllJoynBusAttachment) *IAsyncOperation[*IAllJoynSession]
}

type IAllJoynSessionStaticsVtbl struct {
	win32.IInspectableVtbl
	GetFromServiceInfoAsync                 uintptr
	GetFromServiceInfoAndBusAttachmentAsync uintptr
}

type IAllJoynSessionStatics struct {
	win32.IInspectable
}

func (this *IAllJoynSessionStatics) Vtbl() *IAllJoynSessionStaticsVtbl {
	return (*IAllJoynSessionStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynSessionStatics) GetFromServiceInfoAsync(serviceInfo *IAllJoynServiceInfo) *IAsyncOperation[*IAllJoynSession] {
	var _result *IAsyncOperation[*IAllJoynSession]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFromServiceInfoAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(serviceInfo)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAllJoynSessionStatics) GetFromServiceInfoAndBusAttachmentAsync(serviceInfo *IAllJoynServiceInfo, busAttachment *IAllJoynBusAttachment) *IAsyncOperation[*IAllJoynSession] {
	var _result *IAsyncOperation[*IAllJoynSession]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFromServiceInfoAndBusAttachmentAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(serviceInfo)), uintptr(unsafe.Pointer(busAttachment)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D0B7A17E-0D29-4DA9-8AC6-54C554BEDBC5
var IID_IAllJoynStatusStatics = syscall.GUID{0xD0B7A17E, 0x0D29, 0x4DA9,
	[8]byte{0x8A, 0xC6, 0x54, 0xC5, 0x54, 0xBE, 0xDB, 0xC5}}

type IAllJoynStatusStaticsInterface interface {
	win32.IInspectableInterface
	Get_Ok() int32
	Get_Fail() int32
	Get_OperationTimedOut() int32
	Get_OtherEndClosed() int32
	Get_ConnectionRefused() int32
	Get_AuthenticationFailed() int32
	Get_AuthenticationRejectedByUser() int32
	Get_SslConnectFailed() int32
	Get_SslIdentityVerificationFailed() int32
	Get_InsufficientSecurity() int32
	Get_InvalidArgument1() int32
	Get_InvalidArgument2() int32
	Get_InvalidArgument3() int32
	Get_InvalidArgument4() int32
	Get_InvalidArgument5() int32
	Get_InvalidArgument6() int32
	Get_InvalidArgument7() int32
	Get_InvalidArgument8() int32
}

type IAllJoynStatusStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Ok                            uintptr
	Get_Fail                          uintptr
	Get_OperationTimedOut             uintptr
	Get_OtherEndClosed                uintptr
	Get_ConnectionRefused             uintptr
	Get_AuthenticationFailed          uintptr
	Get_AuthenticationRejectedByUser  uintptr
	Get_SslConnectFailed              uintptr
	Get_SslIdentityVerificationFailed uintptr
	Get_InsufficientSecurity          uintptr
	Get_InvalidArgument1              uintptr
	Get_InvalidArgument2              uintptr
	Get_InvalidArgument3              uintptr
	Get_InvalidArgument4              uintptr
	Get_InvalidArgument5              uintptr
	Get_InvalidArgument6              uintptr
	Get_InvalidArgument7              uintptr
	Get_InvalidArgument8              uintptr
}

type IAllJoynStatusStatics struct {
	win32.IInspectable
}

func (this *IAllJoynStatusStatics) Vtbl() *IAllJoynStatusStaticsVtbl {
	return (*IAllJoynStatusStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynStatusStatics) Get_Ok() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ok, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynStatusStatics) Get_Fail() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Fail, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynStatusStatics) Get_OperationTimedOut() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OperationTimedOut, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynStatusStatics) Get_OtherEndClosed() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OtherEndClosed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynStatusStatics) Get_ConnectionRefused() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConnectionRefused, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynStatusStatics) Get_AuthenticationFailed() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AuthenticationFailed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynStatusStatics) Get_AuthenticationRejectedByUser() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AuthenticationRejectedByUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynStatusStatics) Get_SslConnectFailed() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SslConnectFailed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynStatusStatics) Get_SslIdentityVerificationFailed() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SslIdentityVerificationFailed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynStatusStatics) Get_InsufficientSecurity() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InsufficientSecurity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynStatusStatics) Get_InvalidArgument1() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InvalidArgument1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynStatusStatics) Get_InvalidArgument2() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InvalidArgument2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynStatusStatics) Get_InvalidArgument3() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InvalidArgument3, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynStatusStatics) Get_InvalidArgument4() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InvalidArgument4, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynStatusStatics) Get_InvalidArgument5() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InvalidArgument5, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynStatusStatics) Get_InvalidArgument6() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InvalidArgument6, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynStatusStatics) Get_InvalidArgument7() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InvalidArgument7, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAllJoynStatusStatics) Get_InvalidArgument8() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InvalidArgument8, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C9FCA03B-701D-4AA8-97DD-A2BB0A8F5FA3
var IID_IAllJoynWatcherStoppedEventArgs = syscall.GUID{0xC9FCA03B, 0x701D, 0x4AA8,
	[8]byte{0x97, 0xDD, 0xA2, 0xBB, 0x0A, 0x8F, 0x5F, 0xA3}}

type IAllJoynWatcherStoppedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Status() int32
}

type IAllJoynWatcherStoppedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
}

type IAllJoynWatcherStoppedEventArgs struct {
	win32.IInspectable
}

func (this *IAllJoynWatcherStoppedEventArgs) Vtbl() *IAllJoynWatcherStoppedEventArgsVtbl {
	return (*IAllJoynWatcherStoppedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynWatcherStoppedEventArgs) Get_Status() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 878FA5A8-2D50-47E1-904A-20BF0D48C782
var IID_IAllJoynWatcherStoppedEventArgsFactory = syscall.GUID{0x878FA5A8, 0x2D50, 0x47E1,
	[8]byte{0x90, 0x4A, 0x20, 0xBF, 0x0D, 0x48, 0xC7, 0x82}}

type IAllJoynWatcherStoppedEventArgsFactoryInterface interface {
	win32.IInspectableInterface
	Create(status int32) *IAllJoynWatcherStoppedEventArgs
}

type IAllJoynWatcherStoppedEventArgsFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IAllJoynWatcherStoppedEventArgsFactory struct {
	win32.IInspectable
}

func (this *IAllJoynWatcherStoppedEventArgsFactory) Vtbl() *IAllJoynWatcherStoppedEventArgsFactoryVtbl {
	return (*IAllJoynWatcherStoppedEventArgsFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAllJoynWatcherStoppedEventArgsFactory) Create(status int32) *IAllJoynWatcherStoppedEventArgs {
	var _result *IAllJoynWatcherStoppedEventArgs
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(status), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type AllJoynAboutData struct {
	RtClass
	*IAllJoynAboutData
}

type AllJoynAboutDataView struct {
	RtClass
	*IAllJoynAboutDataView
}

func NewIAllJoynAboutDataViewStatics() *IAllJoynAboutDataViewStatics {
	var p *IAllJoynAboutDataViewStatics
	hs := NewHStr("Windows.Devices.AllJoyn.AllJoynAboutDataView")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAllJoynAboutDataViewStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type AllJoynAcceptSessionJoinerEventArgs struct {
	RtClass
	*IAllJoynAcceptSessionJoinerEventArgs
}

func NewAllJoynAcceptSessionJoinerEventArgs_Create(uniqueName string, sessionPort uint16, trafficType AllJoynTrafficType, proximity byte, acceptSessionJoiner *IAllJoynAcceptSessionJoiner) *AllJoynAcceptSessionJoinerEventArgs {
	hs := NewHStr("Windows.Devices.AllJoyn.AllJoynAcceptSessionJoinerEventArgs")
	var pFac *IAllJoynAcceptSessionJoinerEventArgsFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAllJoynAcceptSessionJoinerEventArgsFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IAllJoynAcceptSessionJoinerEventArgs
	p = pFac.Create(uniqueName, sessionPort, trafficType, proximity, acceptSessionJoiner)
	result := &AllJoynAcceptSessionJoinerEventArgs{
		RtClass:                              RtClass{PInspect: &p.IInspectable},
		IAllJoynAcceptSessionJoinerEventArgs: p,
	}
	com.AddToScope(result)
	return result
}

type AllJoynAuthenticationCompleteEventArgs struct {
	RtClass
	*IAllJoynAuthenticationCompleteEventArgs
}

type AllJoynBusAttachment struct {
	RtClass
	*IAllJoynBusAttachment
}

func NewAllJoynBusAttachment() *AllJoynBusAttachment {
	hs := NewHStr("Windows.Devices.AllJoyn.AllJoynBusAttachment")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &AllJoynBusAttachment{
		RtClass:               RtClass{PInspect: p},
		IAllJoynBusAttachment: (*IAllJoynBusAttachment)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewAllJoynBusAttachment_Create(connectionSpecification string) *AllJoynBusAttachment {
	hs := NewHStr("Windows.Devices.AllJoyn.AllJoynBusAttachment")
	var pFac *IAllJoynBusAttachmentFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAllJoynBusAttachmentFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IAllJoynBusAttachment
	p = pFac.Create(connectionSpecification)
	result := &AllJoynBusAttachment{
		RtClass:               RtClass{PInspect: &p.IInspectable},
		IAllJoynBusAttachment: p,
	}
	com.AddToScope(result)
	return result
}

func NewIAllJoynBusAttachmentStatics() *IAllJoynBusAttachmentStatics {
	var p *IAllJoynBusAttachmentStatics
	hs := NewHStr("Windows.Devices.AllJoyn.AllJoynBusAttachment")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAllJoynBusAttachmentStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type AllJoynBusAttachmentStateChangedEventArgs struct {
	RtClass
	*IAllJoynBusAttachmentStateChangedEventArgs
}

type AllJoynBusObject struct {
	RtClass
	*IAllJoynBusObject
}

func NewAllJoynBusObject() *AllJoynBusObject {
	hs := NewHStr("Windows.Devices.AllJoyn.AllJoynBusObject")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &AllJoynBusObject{
		RtClass:           RtClass{PInspect: p},
		IAllJoynBusObject: (*IAllJoynBusObject)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewAllJoynBusObject_Create(objectPath string) *AllJoynBusObject {
	hs := NewHStr("Windows.Devices.AllJoyn.AllJoynBusObject")
	var pFac *IAllJoynBusObjectFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAllJoynBusObjectFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IAllJoynBusObject
	p = pFac.Create(objectPath)
	result := &AllJoynBusObject{
		RtClass:           RtClass{PInspect: &p.IInspectable},
		IAllJoynBusObject: p,
	}
	com.AddToScope(result)
	return result
}

func NewAllJoynBusObject_CreateWithBusAttachment(objectPath string, busAttachment *IAllJoynBusAttachment) *AllJoynBusObject {
	hs := NewHStr("Windows.Devices.AllJoyn.AllJoynBusObject")
	var pFac *IAllJoynBusObjectFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAllJoynBusObjectFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IAllJoynBusObject
	p = pFac.CreateWithBusAttachment(objectPath, busAttachment)
	result := &AllJoynBusObject{
		RtClass:           RtClass{PInspect: &p.IInspectable},
		IAllJoynBusObject: p,
	}
	com.AddToScope(result)
	return result
}

type AllJoynBusObjectStoppedEventArgs struct {
	RtClass
	*IAllJoynBusObjectStoppedEventArgs
}

func NewAllJoynBusObjectStoppedEventArgs_Create(status int32) *AllJoynBusObjectStoppedEventArgs {
	hs := NewHStr("Windows.Devices.AllJoyn.AllJoynBusObjectStoppedEventArgs")
	var pFac *IAllJoynBusObjectStoppedEventArgsFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAllJoynBusObjectStoppedEventArgsFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IAllJoynBusObjectStoppedEventArgs
	p = pFac.Create(status)
	result := &AllJoynBusObjectStoppedEventArgs{
		RtClass:                           RtClass{PInspect: &p.IInspectable},
		IAllJoynBusObjectStoppedEventArgs: p,
	}
	com.AddToScope(result)
	return result
}

type AllJoynCredentials struct {
	RtClass
	*IAllJoynCredentials
}

type AllJoynCredentialsRequestedEventArgs struct {
	RtClass
	*IAllJoynCredentialsRequestedEventArgs
}

type AllJoynCredentialsVerificationRequestedEventArgs struct {
	RtClass
	*IAllJoynCredentialsVerificationRequestedEventArgs
}

type AllJoynMessageInfo struct {
	RtClass
	*IAllJoynMessageInfo
}

func NewAllJoynMessageInfo_Create(senderUniqueName string) *AllJoynMessageInfo {
	hs := NewHStr("Windows.Devices.AllJoyn.AllJoynMessageInfo")
	var pFac *IAllJoynMessageInfoFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAllJoynMessageInfoFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IAllJoynMessageInfo
	p = pFac.Create(senderUniqueName)
	result := &AllJoynMessageInfo{
		RtClass:             RtClass{PInspect: &p.IInspectable},
		IAllJoynMessageInfo: p,
	}
	com.AddToScope(result)
	return result
}

type AllJoynProducerStoppedEventArgs struct {
	RtClass
	*IAllJoynProducerStoppedEventArgs
}

func NewAllJoynProducerStoppedEventArgs_Create(status int32) *AllJoynProducerStoppedEventArgs {
	hs := NewHStr("Windows.Devices.AllJoyn.AllJoynProducerStoppedEventArgs")
	var pFac *IAllJoynProducerStoppedEventArgsFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAllJoynProducerStoppedEventArgsFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IAllJoynProducerStoppedEventArgs
	p = pFac.Create(status)
	result := &AllJoynProducerStoppedEventArgs{
		RtClass:                          RtClass{PInspect: &p.IInspectable},
		IAllJoynProducerStoppedEventArgs: p,
	}
	com.AddToScope(result)
	return result
}

type AllJoynServiceInfo struct {
	RtClass
	*IAllJoynServiceInfo
}

func NewAllJoynServiceInfo_Create(uniqueName string, objectPath string, sessionPort uint16) *AllJoynServiceInfo {
	hs := NewHStr("Windows.Devices.AllJoyn.AllJoynServiceInfo")
	var pFac *IAllJoynServiceInfoFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAllJoynServiceInfoFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IAllJoynServiceInfo
	p = pFac.Create(uniqueName, objectPath, sessionPort)
	result := &AllJoynServiceInfo{
		RtClass:             RtClass{PInspect: &p.IInspectable},
		IAllJoynServiceInfo: p,
	}
	com.AddToScope(result)
	return result
}

func NewIAllJoynServiceInfoStatics() *IAllJoynServiceInfoStatics {
	var p *IAllJoynServiceInfoStatics
	hs := NewHStr("Windows.Devices.AllJoyn.AllJoynServiceInfo")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAllJoynServiceInfoStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type AllJoynServiceInfoRemovedEventArgs struct {
	RtClass
	*IAllJoynServiceInfoRemovedEventArgs
}

func NewAllJoynServiceInfoRemovedEventArgs_Create(uniqueName string) *AllJoynServiceInfoRemovedEventArgs {
	hs := NewHStr("Windows.Devices.AllJoyn.AllJoynServiceInfoRemovedEventArgs")
	var pFac *IAllJoynServiceInfoRemovedEventArgsFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAllJoynServiceInfoRemovedEventArgsFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IAllJoynServiceInfoRemovedEventArgs
	p = pFac.Create(uniqueName)
	result := &AllJoynServiceInfoRemovedEventArgs{
		RtClass:                             RtClass{PInspect: &p.IInspectable},
		IAllJoynServiceInfoRemovedEventArgs: p,
	}
	com.AddToScope(result)
	return result
}

type AllJoynSession struct {
	RtClass
	*IAllJoynSession
}

func NewIAllJoynSessionStatics() *IAllJoynSessionStatics {
	var p *IAllJoynSessionStatics
	hs := NewHStr("Windows.Devices.AllJoyn.AllJoynSession")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAllJoynSessionStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type AllJoynSessionJoinedEventArgs struct {
	RtClass
	*IAllJoynSessionJoinedEventArgs
}

func NewAllJoynSessionJoinedEventArgs_Create(session *IAllJoynSession) *AllJoynSessionJoinedEventArgs {
	hs := NewHStr("Windows.Devices.AllJoyn.AllJoynSessionJoinedEventArgs")
	var pFac *IAllJoynSessionJoinedEventArgsFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAllJoynSessionJoinedEventArgsFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IAllJoynSessionJoinedEventArgs
	p = pFac.Create(session)
	result := &AllJoynSessionJoinedEventArgs{
		RtClass:                        RtClass{PInspect: &p.IInspectable},
		IAllJoynSessionJoinedEventArgs: p,
	}
	com.AddToScope(result)
	return result
}

type AllJoynSessionLostEventArgs struct {
	RtClass
	*IAllJoynSessionLostEventArgs
}

func NewAllJoynSessionLostEventArgs_Create(reason AllJoynSessionLostReason) *AllJoynSessionLostEventArgs {
	hs := NewHStr("Windows.Devices.AllJoyn.AllJoynSessionLostEventArgs")
	var pFac *IAllJoynSessionLostEventArgsFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAllJoynSessionLostEventArgsFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IAllJoynSessionLostEventArgs
	p = pFac.Create(reason)
	result := &AllJoynSessionLostEventArgs{
		RtClass:                      RtClass{PInspect: &p.IInspectable},
		IAllJoynSessionLostEventArgs: p,
	}
	com.AddToScope(result)
	return result
}

type AllJoynSessionMemberAddedEventArgs struct {
	RtClass
	*IAllJoynSessionMemberAddedEventArgs
}

func NewAllJoynSessionMemberAddedEventArgs_Create(uniqueName string) *AllJoynSessionMemberAddedEventArgs {
	hs := NewHStr("Windows.Devices.AllJoyn.AllJoynSessionMemberAddedEventArgs")
	var pFac *IAllJoynSessionMemberAddedEventArgsFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAllJoynSessionMemberAddedEventArgsFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IAllJoynSessionMemberAddedEventArgs
	p = pFac.Create(uniqueName)
	result := &AllJoynSessionMemberAddedEventArgs{
		RtClass:                             RtClass{PInspect: &p.IInspectable},
		IAllJoynSessionMemberAddedEventArgs: p,
	}
	com.AddToScope(result)
	return result
}

type AllJoynSessionMemberRemovedEventArgs struct {
	RtClass
	*IAllJoynSessionMemberRemovedEventArgs
}

func NewAllJoynSessionMemberRemovedEventArgs_Create(uniqueName string) *AllJoynSessionMemberRemovedEventArgs {
	hs := NewHStr("Windows.Devices.AllJoyn.AllJoynSessionMemberRemovedEventArgs")
	var pFac *IAllJoynSessionMemberRemovedEventArgsFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAllJoynSessionMemberRemovedEventArgsFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IAllJoynSessionMemberRemovedEventArgs
	p = pFac.Create(uniqueName)
	result := &AllJoynSessionMemberRemovedEventArgs{
		RtClass:                               RtClass{PInspect: &p.IInspectable},
		IAllJoynSessionMemberRemovedEventArgs: p,
	}
	com.AddToScope(result)
	return result
}

type AllJoynStatus struct {
	RtClass
}

func NewIAllJoynStatusStatics() *IAllJoynStatusStatics {
	var p *IAllJoynStatusStatics
	hs := NewHStr("Windows.Devices.AllJoyn.AllJoynStatus")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAllJoynStatusStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type AllJoynWatcherStoppedEventArgs struct {
	RtClass
	*IAllJoynWatcherStoppedEventArgs
}

func NewAllJoynWatcherStoppedEventArgs_Create(status int32) *AllJoynWatcherStoppedEventArgs {
	hs := NewHStr("Windows.Devices.AllJoyn.AllJoynWatcherStoppedEventArgs")
	var pFac *IAllJoynWatcherStoppedEventArgsFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAllJoynWatcherStoppedEventArgsFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IAllJoynWatcherStoppedEventArgs
	p = pFac.Create(status)
	result := &AllJoynWatcherStoppedEventArgs{
		RtClass:                         RtClass{PInspect: &p.IInspectable},
		IAllJoynWatcherStoppedEventArgs: p,
	}
	com.AddToScope(result)
	return result
}
