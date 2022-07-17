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
type GraphicsTrustStatus int32

const (
	GraphicsTrustStatus_TrustNotRequired        GraphicsTrustStatus = 0
	GraphicsTrustStatus_TrustEstablished        GraphicsTrustStatus = 1
	GraphicsTrustStatus_EnvironmentNotSupported GraphicsTrustStatus = 2
	GraphicsTrustStatus_DriverNotSupported      GraphicsTrustStatus = 3
	GraphicsTrustStatus_DriverSigningFailure    GraphicsTrustStatus = 4
	GraphicsTrustStatus_UnknownFailure          GraphicsTrustStatus = 5
)

// enum
type HdcpProtection int32

const (
	HdcpProtection_Off                   HdcpProtection = 0
	HdcpProtection_On                    HdcpProtection = 1
	HdcpProtection_OnWithTypeEnforcement HdcpProtection = 2
)

// enum
type HdcpSetProtectionResult int32

const (
	HdcpSetProtectionResult_Success        HdcpSetProtectionResult = 0
	HdcpSetProtectionResult_TimedOut       HdcpSetProtectionResult = 1
	HdcpSetProtectionResult_NotSupported   HdcpSetProtectionResult = 2
	HdcpSetProtectionResult_UnknownFailure HdcpSetProtectionResult = 3
)

// enum
type ProtectionCapabilityResult int32

const (
	ProtectionCapabilityResult_NotSupported ProtectionCapabilityResult = 0
	ProtectionCapabilityResult_Maybe        ProtectionCapabilityResult = 1
	ProtectionCapabilityResult_Probably     ProtectionCapabilityResult = 2
)

// enum
type RenewalStatus int32

const (
	RenewalStatus_NotStarted                   RenewalStatus = 0
	RenewalStatus_UpdatesInProgress            RenewalStatus = 1
	RenewalStatus_UserCancelled                RenewalStatus = 2
	RenewalStatus_AppComponentsMayNeedUpdating RenewalStatus = 3
	RenewalStatus_NoComponentsFound            RenewalStatus = 4
)

// enum
// flags
type RevocationAndRenewalReasons uint32

const (
	RevocationAndRenewalReasons_UserModeComponentLoad                     RevocationAndRenewalReasons = 1
	RevocationAndRenewalReasons_KernelModeComponentLoad                   RevocationAndRenewalReasons = 2
	RevocationAndRenewalReasons_AppComponent                              RevocationAndRenewalReasons = 4
	RevocationAndRenewalReasons_GlobalRevocationListLoadFailed            RevocationAndRenewalReasons = 16
	RevocationAndRenewalReasons_InvalidGlobalRevocationListSignature      RevocationAndRenewalReasons = 32
	RevocationAndRenewalReasons_GlobalRevocationListAbsent                RevocationAndRenewalReasons = 4096
	RevocationAndRenewalReasons_ComponentRevoked                          RevocationAndRenewalReasons = 8192
	RevocationAndRenewalReasons_InvalidComponentCertificateExtendedKeyUse RevocationAndRenewalReasons = 16384
	RevocationAndRenewalReasons_ComponentCertificateRevoked               RevocationAndRenewalReasons = 32768
	RevocationAndRenewalReasons_InvalidComponentCertificateRoot           RevocationAndRenewalReasons = 65536
	RevocationAndRenewalReasons_ComponentHighSecurityCertificateRevoked   RevocationAndRenewalReasons = 131072
	RevocationAndRenewalReasons_ComponentLowSecurityCertificateRevoked    RevocationAndRenewalReasons = 262144
	RevocationAndRenewalReasons_BootDriverVerificationFailed              RevocationAndRenewalReasons = 1048576
	RevocationAndRenewalReasons_ComponentSignedWithTestCertificate        RevocationAndRenewalReasons = 16777216
	RevocationAndRenewalReasons_EncryptionFailure                         RevocationAndRenewalReasons = 268435456
)

// structs

type ProtectionRenewalContract struct {
}

// func types

//95DA643C-6DB9-424B-86CA-091AF432081C
type ComponentLoadFailedEventHandler func(sender *IMediaProtectionManager, e *IComponentLoadFailedEventArgs) com.Error

//64E12A45-973B-4A3A-B260-91898A49A82C
type RebootNeededEventHandler func(sender *IMediaProtectionManager) com.Error

//D2D690BA-CAC9-48E1-95C0-D38495A84055
type ServiceRequestedEventHandler func(sender *IMediaProtectionManager, e *IServiceRequestedEventArgs) com.Error

// interfaces

// 95972E93-7746-417E-8495-F031BBC5862C
var IID_IComponentLoadFailedEventArgs = syscall.GUID{0x95972E93, 0x7746, 0x417E,
	[8]byte{0x84, 0x95, 0xF0, 0x31, 0xBB, 0xC5, 0x86, 0x2C}}

type IComponentLoadFailedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Information() *IRevocationAndRenewalInformation
	Get_Completion() *IMediaProtectionServiceCompletion
}

type IComponentLoadFailedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Information uintptr
	Get_Completion  uintptr
}

type IComponentLoadFailedEventArgs struct {
	win32.IInspectable
}

func (this *IComponentLoadFailedEventArgs) Vtbl() *IComponentLoadFailedEventArgsVtbl {
	return (*IComponentLoadFailedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IComponentLoadFailedEventArgs) Get_Information() *IRevocationAndRenewalInformation {
	var _result *IRevocationAndRenewalInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Information, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IComponentLoadFailedEventArgs) Get_Completion() *IMediaProtectionServiceCompletion {
	var _result *IMediaProtectionServiceCompletion
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Completion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6FFBCD67-B795-48C5-8B7B-A7C4EFE202E3
var IID_IComponentRenewalStatics = syscall.GUID{0x6FFBCD67, 0xB795, 0x48C5,
	[8]byte{0x8B, 0x7B, 0xA7, 0xC4, 0xEF, 0xE2, 0x02, 0xE3}}

type IComponentRenewalStaticsInterface interface {
	win32.IInspectableInterface
	RenewSystemComponentsAsync(information *IRevocationAndRenewalInformation) *IAsyncOperationWithProgress[RenewalStatus, uint32]
}

type IComponentRenewalStaticsVtbl struct {
	win32.IInspectableVtbl
	RenewSystemComponentsAsync uintptr
}

type IComponentRenewalStatics struct {
	win32.IInspectable
}

func (this *IComponentRenewalStatics) Vtbl() *IComponentRenewalStaticsVtbl {
	return (*IComponentRenewalStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IComponentRenewalStatics) RenewSystemComponentsAsync(information *IRevocationAndRenewalInformation) *IAsyncOperationWithProgress[RenewalStatus, uint32] {
	var _result *IAsyncOperationWithProgress[RenewalStatus, uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RenewSystemComponentsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(information)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 718845E9-64D7-426D-809B-1BE461941A2A
var IID_IHdcpSession = syscall.GUID{0x718845E9, 0x64D7, 0x426D,
	[8]byte{0x80, 0x9B, 0x1B, 0xE4, 0x61, 0x94, 0x1A, 0x2A}}

type IHdcpSessionInterface interface {
	win32.IInspectableInterface
	IsEffectiveProtectionAtLeast(protection HdcpProtection) bool
	GetEffectiveProtection() *IReference[HdcpProtection]
	SetDesiredMinProtectionAsync(protection HdcpProtection) *IAsyncOperation[HdcpSetProtectionResult]
	Add_ProtectionChanged(handler TypedEventHandler[*IHdcpSession, interface{}]) EventRegistrationToken
	Remove_ProtectionChanged(token EventRegistrationToken)
}

type IHdcpSessionVtbl struct {
	win32.IInspectableVtbl
	IsEffectiveProtectionAtLeast uintptr
	GetEffectiveProtection       uintptr
	SetDesiredMinProtectionAsync uintptr
	Add_ProtectionChanged        uintptr
	Remove_ProtectionChanged     uintptr
}

type IHdcpSession struct {
	win32.IInspectable
}

func (this *IHdcpSession) Vtbl() *IHdcpSessionVtbl {
	return (*IHdcpSessionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHdcpSession) IsEffectiveProtectionAtLeast(protection HdcpProtection) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsEffectiveProtectionAtLeast, uintptr(unsafe.Pointer(this)), uintptr(protection), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHdcpSession) GetEffectiveProtection() *IReference[HdcpProtection] {
	var _result *IReference[HdcpProtection]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetEffectiveProtection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHdcpSession) SetDesiredMinProtectionAsync(protection HdcpProtection) *IAsyncOperation[HdcpSetProtectionResult] {
	var _result *IAsyncOperation[HdcpSetProtectionResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetDesiredMinProtectionAsync, uintptr(unsafe.Pointer(this)), uintptr(protection), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHdcpSession) Add_ProtectionChanged(handler TypedEventHandler[*IHdcpSession, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ProtectionChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHdcpSession) Remove_ProtectionChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ProtectionChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 45694947-C741-434B-A79E-474C12D93D2F
var IID_IMediaProtectionManager = syscall.GUID{0x45694947, 0xC741, 0x434B,
	[8]byte{0xA7, 0x9E, 0x47, 0x4C, 0x12, 0xD9, 0x3D, 0x2F}}

type IMediaProtectionManagerInterface interface {
	win32.IInspectableInterface
	Add_ServiceRequested(handler ServiceRequestedEventHandler) EventRegistrationToken
	Remove_ServiceRequested(cookie EventRegistrationToken)
	Add_RebootNeeded(handler RebootNeededEventHandler) EventRegistrationToken
	Remove_RebootNeeded(cookie EventRegistrationToken)
	Add_ComponentLoadFailed(handler ComponentLoadFailedEventHandler) EventRegistrationToken
	Remove_ComponentLoadFailed(cookie EventRegistrationToken)
	Get_Properties() *IPropertySet
}

type IMediaProtectionManagerVtbl struct {
	win32.IInspectableVtbl
	Add_ServiceRequested       uintptr
	Remove_ServiceRequested    uintptr
	Add_RebootNeeded           uintptr
	Remove_RebootNeeded        uintptr
	Add_ComponentLoadFailed    uintptr
	Remove_ComponentLoadFailed uintptr
	Get_Properties             uintptr
}

type IMediaProtectionManager struct {
	win32.IInspectable
}

func (this *IMediaProtectionManager) Vtbl() *IMediaProtectionManagerVtbl {
	return (*IMediaProtectionManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaProtectionManager) Add_ServiceRequested(handler ServiceRequestedEventHandler) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ServiceRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaProtectionManager) Remove_ServiceRequested(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ServiceRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IMediaProtectionManager) Add_RebootNeeded(handler RebootNeededEventHandler) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_RebootNeeded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewOneArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaProtectionManager) Remove_RebootNeeded(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_RebootNeeded, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IMediaProtectionManager) Add_ComponentLoadFailed(handler ComponentLoadFailedEventHandler) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ComponentLoadFailed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaProtectionManager) Remove_ComponentLoadFailed(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ComponentLoadFailed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IMediaProtectionManager) Get_Properties() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0C111226-7B26-4D31-95BB-9C1B08EF7FC0
var IID_IMediaProtectionPMPServer = syscall.GUID{0x0C111226, 0x7B26, 0x4D31,
	[8]byte{0x95, 0xBB, 0x9C, 0x1B, 0x08, 0xEF, 0x7F, 0xC0}}

type IMediaProtectionPMPServerInterface interface {
	win32.IInspectableInterface
	Get_Properties() *IPropertySet
}

type IMediaProtectionPMPServerVtbl struct {
	win32.IInspectableVtbl
	Get_Properties uintptr
}

type IMediaProtectionPMPServer struct {
	win32.IInspectable
}

func (this *IMediaProtectionPMPServer) Vtbl() *IMediaProtectionPMPServerVtbl {
	return (*IMediaProtectionPMPServerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaProtectionPMPServer) Get_Properties() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 602C8E5E-F7D2-487E-AF91-DBC4252B2182
var IID_IMediaProtectionPMPServerFactory = syscall.GUID{0x602C8E5E, 0xF7D2, 0x487E,
	[8]byte{0xAF, 0x91, 0xDB, 0xC4, 0x25, 0x2B, 0x21, 0x82}}

type IMediaProtectionPMPServerFactoryInterface interface {
	win32.IInspectableInterface
	CreatePMPServer(pProperties *IPropertySet) *IMediaProtectionPMPServer
}

type IMediaProtectionPMPServerFactoryVtbl struct {
	win32.IInspectableVtbl
	CreatePMPServer uintptr
}

type IMediaProtectionPMPServerFactory struct {
	win32.IInspectable
}

func (this *IMediaProtectionPMPServerFactory) Vtbl() *IMediaProtectionPMPServerFactoryVtbl {
	return (*IMediaProtectionPMPServerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaProtectionPMPServerFactory) CreatePMPServer(pProperties *IPropertySet) *IMediaProtectionPMPServer {
	var _result *IMediaProtectionPMPServer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreatePMPServer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pProperties)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8B5CCA18-CFD5-44EE-A2ED-DF76010C14B5
var IID_IMediaProtectionServiceCompletion = syscall.GUID{0x8B5CCA18, 0xCFD5, 0x44EE,
	[8]byte{0xA2, 0xED, 0xDF, 0x76, 0x01, 0x0C, 0x14, 0xB5}}

type IMediaProtectionServiceCompletionInterface interface {
	win32.IInspectableInterface
	Complete(success bool)
}

type IMediaProtectionServiceCompletionVtbl struct {
	win32.IInspectableVtbl
	Complete uintptr
}

type IMediaProtectionServiceCompletion struct {
	win32.IInspectable
}

func (this *IMediaProtectionServiceCompletion) Vtbl() *IMediaProtectionServiceCompletionVtbl {
	return (*IMediaProtectionServiceCompletionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaProtectionServiceCompletion) Complete(success bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Complete, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&success))))
	_ = _hr
}

// B1DE0EA6-2094-478D-87A4-8B95200F85C6
var IID_IMediaProtectionServiceRequest = syscall.GUID{0xB1DE0EA6, 0x2094, 0x478D,
	[8]byte{0x87, 0xA4, 0x8B, 0x95, 0x20, 0x0F, 0x85, 0xC6}}

type IMediaProtectionServiceRequestInterface interface {
	win32.IInspectableInterface
	Get_ProtectionSystem() syscall.GUID
	Get_Type() syscall.GUID
}

type IMediaProtectionServiceRequestVtbl struct {
	win32.IInspectableVtbl
	Get_ProtectionSystem uintptr
	Get_Type             uintptr
}

type IMediaProtectionServiceRequest struct {
	win32.IInspectable
}

func (this *IMediaProtectionServiceRequest) Vtbl() *IMediaProtectionServiceRequestVtbl {
	return (*IMediaProtectionServiceRequestVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaProtectionServiceRequest) Get_ProtectionSystem() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProtectionSystem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaProtectionServiceRequest) Get_Type() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Type, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C7AC5D7E-7480-4D29-A464-7BCD913DD8E4
var IID_IProtectionCapabilities = syscall.GUID{0xC7AC5D7E, 0x7480, 0x4D29,
	[8]byte{0xA4, 0x64, 0x7B, 0xCD, 0x91, 0x3D, 0xD8, 0xE4}}

type IProtectionCapabilitiesInterface interface {
	win32.IInspectableInterface
	IsTypeSupported(type_ string, keySystem string) ProtectionCapabilityResult
}

type IProtectionCapabilitiesVtbl struct {
	win32.IInspectableVtbl
	IsTypeSupported uintptr
}

type IProtectionCapabilities struct {
	win32.IInspectable
}

func (this *IProtectionCapabilities) Vtbl() *IProtectionCapabilitiesVtbl {
	return (*IProtectionCapabilitiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProtectionCapabilities) IsTypeSupported(type_ string, keySystem string) ProtectionCapabilityResult {
	var _result ProtectionCapabilityResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsTypeSupported, uintptr(unsafe.Pointer(this)), NewHStr(type_).Ptr, NewHStr(keySystem).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F3A1937B-2501-439E-A6E7-6FC95E175FCF
var IID_IRevocationAndRenewalInformation = syscall.GUID{0xF3A1937B, 0x2501, 0x439E,
	[8]byte{0xA6, 0xE7, 0x6F, 0xC9, 0x5E, 0x17, 0x5F, 0xCF}}

type IRevocationAndRenewalInformationInterface interface {
	win32.IInspectableInterface
	Get_Items() *IVector[*IRevocationAndRenewalItem]
}

type IRevocationAndRenewalInformationVtbl struct {
	win32.IInspectableVtbl
	Get_Items uintptr
}

type IRevocationAndRenewalInformation struct {
	win32.IInspectable
}

func (this *IRevocationAndRenewalInformation) Vtbl() *IRevocationAndRenewalInformationVtbl {
	return (*IRevocationAndRenewalInformationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRevocationAndRenewalInformation) Get_Items() *IVector[*IRevocationAndRenewalItem] {
	var _result *IVector[*IRevocationAndRenewalItem]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Items, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3099C20C-3CF0-49EA-902D-CAF32D2DDE2C
var IID_IRevocationAndRenewalItem = syscall.GUID{0x3099C20C, 0x3CF0, 0x49EA,
	[8]byte{0x90, 0x2D, 0xCA, 0xF3, 0x2D, 0x2D, 0xDE, 0x2C}}

type IRevocationAndRenewalItemInterface interface {
	win32.IInspectableInterface
	Get_Reasons() RevocationAndRenewalReasons
	Get_HeaderHash() string
	Get_PublicKeyHash() string
	Get_Name() string
	Get_RenewalId() string
}

type IRevocationAndRenewalItemVtbl struct {
	win32.IInspectableVtbl
	Get_Reasons       uintptr
	Get_HeaderHash    uintptr
	Get_PublicKeyHash uintptr
	Get_Name          uintptr
	Get_RenewalId     uintptr
}

type IRevocationAndRenewalItem struct {
	win32.IInspectable
}

func (this *IRevocationAndRenewalItem) Vtbl() *IRevocationAndRenewalItemVtbl {
	return (*IRevocationAndRenewalItemVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRevocationAndRenewalItem) Get_Reasons() RevocationAndRenewalReasons {
	var _result RevocationAndRenewalReasons
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Reasons, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRevocationAndRenewalItem) Get_HeaderHash() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HeaderHash, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IRevocationAndRenewalItem) Get_PublicKeyHash() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PublicKeyHash, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IRevocationAndRenewalItem) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IRevocationAndRenewalItem) Get_RenewalId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RenewalId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 34283BAF-ABB4-4FC1-BD89-93F106573A49
var IID_IServiceRequestedEventArgs = syscall.GUID{0x34283BAF, 0xABB4, 0x4FC1,
	[8]byte{0xBD, 0x89, 0x93, 0xF1, 0x06, 0x57, 0x3A, 0x49}}

type IServiceRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Request() *IMediaProtectionServiceRequest
	Get_Completion() *IMediaProtectionServiceCompletion
}

type IServiceRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Request    uintptr
	Get_Completion uintptr
}

type IServiceRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IServiceRequestedEventArgs) Vtbl() *IServiceRequestedEventArgsVtbl {
	return (*IServiceRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IServiceRequestedEventArgs) Get_Request() *IMediaProtectionServiceRequest {
	var _result *IMediaProtectionServiceRequest
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Request, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IServiceRequestedEventArgs) Get_Completion() *IMediaProtectionServiceCompletion {
	var _result *IMediaProtectionServiceCompletion
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Completion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 553C69D6-FAFE-4128-8DFA-130E398A13A7
var IID_IServiceRequestedEventArgs2 = syscall.GUID{0x553C69D6, 0xFAFE, 0x4128,
	[8]byte{0x8D, 0xFA, 0x13, 0x0E, 0x39, 0x8A, 0x13, 0xA7}}

type IServiceRequestedEventArgs2Interface interface {
	win32.IInspectableInterface
	Get_MediaPlaybackItem() *IMediaPlaybackItem
}

type IServiceRequestedEventArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_MediaPlaybackItem uintptr
}

type IServiceRequestedEventArgs2 struct {
	win32.IInspectable
}

func (this *IServiceRequestedEventArgs2) Vtbl() *IServiceRequestedEventArgs2Vtbl {
	return (*IServiceRequestedEventArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IServiceRequestedEventArgs2) Get_MediaPlaybackItem() *IMediaPlaybackItem {
	var _result *IMediaPlaybackItem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaPlaybackItem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type HdcpSession struct {
	RtClass
	*IHdcpSession
}

func NewHdcpSession() *HdcpSession {
	hs := NewHStr("Windows.Media.Protection.HdcpSession")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &HdcpSession{
		RtClass:      RtClass{PInspect: p},
		IHdcpSession: (*IHdcpSession)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type MediaProtectionPMPServer struct {
	RtClass
	*IMediaProtectionPMPServer
}

func NewMediaProtectionPMPServer_CreatePMPServer(pProperties *IPropertySet) *MediaProtectionPMPServer {
	hs := NewHStr("Windows.Media.Protection.MediaProtectionPMPServer")
	var pFac *IMediaProtectionPMPServerFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMediaProtectionPMPServerFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IMediaProtectionPMPServer
	p = pFac.CreatePMPServer(pProperties)
	result := &MediaProtectionPMPServer{
		RtClass:                   RtClass{PInspect: &p.IInspectable},
		IMediaProtectionPMPServer: p,
	}
	com.AddToScope(result)
	return result
}

type ProtectionCapabilities struct {
	RtClass
	*IProtectionCapabilities
}

func NewProtectionCapabilities() *ProtectionCapabilities {
	hs := NewHStr("Windows.Media.Protection.ProtectionCapabilities")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &ProtectionCapabilities{
		RtClass:                 RtClass{PInspect: p},
		IProtectionCapabilities: (*IProtectionCapabilities)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}
