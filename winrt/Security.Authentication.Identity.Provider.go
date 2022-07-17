package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type SecondaryAuthenticationFactorAuthenticationMessage int32

const (
	SecondaryAuthenticationFactorAuthenticationMessage_Invalid                   SecondaryAuthenticationFactorAuthenticationMessage = 0
	SecondaryAuthenticationFactorAuthenticationMessage_SwipeUpWelcome            SecondaryAuthenticationFactorAuthenticationMessage = 1
	SecondaryAuthenticationFactorAuthenticationMessage_TapWelcome                SecondaryAuthenticationFactorAuthenticationMessage = 2
	SecondaryAuthenticationFactorAuthenticationMessage_DeviceNeedsAttention      SecondaryAuthenticationFactorAuthenticationMessage = 3
	SecondaryAuthenticationFactorAuthenticationMessage_LookingForDevice          SecondaryAuthenticationFactorAuthenticationMessage = 4
	SecondaryAuthenticationFactorAuthenticationMessage_LookingForDevicePluggedin SecondaryAuthenticationFactorAuthenticationMessage = 5
	SecondaryAuthenticationFactorAuthenticationMessage_BluetoothIsDisabled       SecondaryAuthenticationFactorAuthenticationMessage = 6
	SecondaryAuthenticationFactorAuthenticationMessage_NfcIsDisabled             SecondaryAuthenticationFactorAuthenticationMessage = 7
	SecondaryAuthenticationFactorAuthenticationMessage_WiFiIsDisabled            SecondaryAuthenticationFactorAuthenticationMessage = 8
	SecondaryAuthenticationFactorAuthenticationMessage_ExtraTapIsRequired        SecondaryAuthenticationFactorAuthenticationMessage = 9
	SecondaryAuthenticationFactorAuthenticationMessage_DisabledByPolicy          SecondaryAuthenticationFactorAuthenticationMessage = 10
	SecondaryAuthenticationFactorAuthenticationMessage_TapOnDeviceRequired       SecondaryAuthenticationFactorAuthenticationMessage = 11
	SecondaryAuthenticationFactorAuthenticationMessage_HoldFinger                SecondaryAuthenticationFactorAuthenticationMessage = 12
	SecondaryAuthenticationFactorAuthenticationMessage_ScanFinger                SecondaryAuthenticationFactorAuthenticationMessage = 13
	SecondaryAuthenticationFactorAuthenticationMessage_UnauthorizedUser          SecondaryAuthenticationFactorAuthenticationMessage = 14
	SecondaryAuthenticationFactorAuthenticationMessage_ReregisterRequired        SecondaryAuthenticationFactorAuthenticationMessage = 15
	SecondaryAuthenticationFactorAuthenticationMessage_TryAgain                  SecondaryAuthenticationFactorAuthenticationMessage = 16
	SecondaryAuthenticationFactorAuthenticationMessage_SayPassphrase             SecondaryAuthenticationFactorAuthenticationMessage = 17
	SecondaryAuthenticationFactorAuthenticationMessage_ReadyToSignIn             SecondaryAuthenticationFactorAuthenticationMessage = 18
	SecondaryAuthenticationFactorAuthenticationMessage_UseAnotherSignInOption    SecondaryAuthenticationFactorAuthenticationMessage = 19
	SecondaryAuthenticationFactorAuthenticationMessage_ConnectionRequired        SecondaryAuthenticationFactorAuthenticationMessage = 20
	SecondaryAuthenticationFactorAuthenticationMessage_TimeLimitExceeded         SecondaryAuthenticationFactorAuthenticationMessage = 21
	SecondaryAuthenticationFactorAuthenticationMessage_CanceledByUser            SecondaryAuthenticationFactorAuthenticationMessage = 22
	SecondaryAuthenticationFactorAuthenticationMessage_CenterHand                SecondaryAuthenticationFactorAuthenticationMessage = 23
	SecondaryAuthenticationFactorAuthenticationMessage_MoveHandCloser            SecondaryAuthenticationFactorAuthenticationMessage = 24
	SecondaryAuthenticationFactorAuthenticationMessage_MoveHandFarther           SecondaryAuthenticationFactorAuthenticationMessage = 25
	SecondaryAuthenticationFactorAuthenticationMessage_PlaceHandAbove            SecondaryAuthenticationFactorAuthenticationMessage = 26
	SecondaryAuthenticationFactorAuthenticationMessage_RecognitionFailed         SecondaryAuthenticationFactorAuthenticationMessage = 27
	SecondaryAuthenticationFactorAuthenticationMessage_DeviceUnavailable         SecondaryAuthenticationFactorAuthenticationMessage = 28
)

// enum
type SecondaryAuthenticationFactorAuthenticationScenario int32

const (
	SecondaryAuthenticationFactorAuthenticationScenario_SignIn           SecondaryAuthenticationFactorAuthenticationScenario = 0
	SecondaryAuthenticationFactorAuthenticationScenario_CredentialPrompt SecondaryAuthenticationFactorAuthenticationScenario = 1
)

// enum
type SecondaryAuthenticationFactorAuthenticationStage int32

const (
	SecondaryAuthenticationFactorAuthenticationStage_NotStarted                 SecondaryAuthenticationFactorAuthenticationStage = 0
	SecondaryAuthenticationFactorAuthenticationStage_WaitingForUserConfirmation SecondaryAuthenticationFactorAuthenticationStage = 1
	SecondaryAuthenticationFactorAuthenticationStage_CollectingCredential       SecondaryAuthenticationFactorAuthenticationStage = 2
	SecondaryAuthenticationFactorAuthenticationStage_SuspendingAuthentication   SecondaryAuthenticationFactorAuthenticationStage = 3
	SecondaryAuthenticationFactorAuthenticationStage_CredentialCollected        SecondaryAuthenticationFactorAuthenticationStage = 4
	SecondaryAuthenticationFactorAuthenticationStage_CredentialAuthenticated    SecondaryAuthenticationFactorAuthenticationStage = 5
	SecondaryAuthenticationFactorAuthenticationStage_StoppingAuthentication     SecondaryAuthenticationFactorAuthenticationStage = 6
	SecondaryAuthenticationFactorAuthenticationStage_ReadyForLock               SecondaryAuthenticationFactorAuthenticationStage = 7
	SecondaryAuthenticationFactorAuthenticationStage_CheckingDevicePresence     SecondaryAuthenticationFactorAuthenticationStage = 8
)

// enum
type SecondaryAuthenticationFactorAuthenticationStatus int32

const (
	SecondaryAuthenticationFactorAuthenticationStatus_Failed                     SecondaryAuthenticationFactorAuthenticationStatus = 0
	SecondaryAuthenticationFactorAuthenticationStatus_Started                    SecondaryAuthenticationFactorAuthenticationStatus = 1
	SecondaryAuthenticationFactorAuthenticationStatus_UnknownDevice              SecondaryAuthenticationFactorAuthenticationStatus = 2
	SecondaryAuthenticationFactorAuthenticationStatus_DisabledByPolicy           SecondaryAuthenticationFactorAuthenticationStatus = 3
	SecondaryAuthenticationFactorAuthenticationStatus_InvalidAuthenticationStage SecondaryAuthenticationFactorAuthenticationStatus = 4
)

// enum
// flags
type SecondaryAuthenticationFactorDeviceCapabilities uint32

const (
	SecondaryAuthenticationFactorDeviceCapabilities_None                            SecondaryAuthenticationFactorDeviceCapabilities = 0
	SecondaryAuthenticationFactorDeviceCapabilities_SecureStorage                   SecondaryAuthenticationFactorDeviceCapabilities = 1
	SecondaryAuthenticationFactorDeviceCapabilities_StoreKeys                       SecondaryAuthenticationFactorDeviceCapabilities = 2
	SecondaryAuthenticationFactorDeviceCapabilities_ConfirmUserIntentToAuthenticate SecondaryAuthenticationFactorDeviceCapabilities = 4
	SecondaryAuthenticationFactorDeviceCapabilities_SupportSecureUserPresenceCheck  SecondaryAuthenticationFactorDeviceCapabilities = 8
	SecondaryAuthenticationFactorDeviceCapabilities_TransmittedDataIsEncrypted      SecondaryAuthenticationFactorDeviceCapabilities = 16
	SecondaryAuthenticationFactorDeviceCapabilities_HMacSha256                      SecondaryAuthenticationFactorDeviceCapabilities = 32
	SecondaryAuthenticationFactorDeviceCapabilities_CloseRangeDataTransmission      SecondaryAuthenticationFactorDeviceCapabilities = 64
)

// enum
type SecondaryAuthenticationFactorDeviceFindScope int32

const (
	SecondaryAuthenticationFactorDeviceFindScope_User     SecondaryAuthenticationFactorDeviceFindScope = 0
	SecondaryAuthenticationFactorDeviceFindScope_AllUsers SecondaryAuthenticationFactorDeviceFindScope = 1
)

// enum
type SecondaryAuthenticationFactorDevicePresence int32

const (
	SecondaryAuthenticationFactorDevicePresence_Absent  SecondaryAuthenticationFactorDevicePresence = 0
	SecondaryAuthenticationFactorDevicePresence_Present SecondaryAuthenticationFactorDevicePresence = 1
)

// enum
type SecondaryAuthenticationFactorDevicePresenceMonitoringMode int32

const (
	SecondaryAuthenticationFactorDevicePresenceMonitoringMode_Unsupported   SecondaryAuthenticationFactorDevicePresenceMonitoringMode = 0
	SecondaryAuthenticationFactorDevicePresenceMonitoringMode_AppManaged    SecondaryAuthenticationFactorDevicePresenceMonitoringMode = 1
	SecondaryAuthenticationFactorDevicePresenceMonitoringMode_SystemManaged SecondaryAuthenticationFactorDevicePresenceMonitoringMode = 2
)

// enum
type SecondaryAuthenticationFactorDevicePresenceMonitoringRegistrationStatus int32

const (
	SecondaryAuthenticationFactorDevicePresenceMonitoringRegistrationStatus_Unsupported      SecondaryAuthenticationFactorDevicePresenceMonitoringRegistrationStatus = 0
	SecondaryAuthenticationFactorDevicePresenceMonitoringRegistrationStatus_Succeeded        SecondaryAuthenticationFactorDevicePresenceMonitoringRegistrationStatus = 1
	SecondaryAuthenticationFactorDevicePresenceMonitoringRegistrationStatus_DisabledByPolicy SecondaryAuthenticationFactorDevicePresenceMonitoringRegistrationStatus = 2
)

// enum
type SecondaryAuthenticationFactorFinishAuthenticationStatus int32

const (
	SecondaryAuthenticationFactorFinishAuthenticationStatus_Failed       SecondaryAuthenticationFactorFinishAuthenticationStatus = 0
	SecondaryAuthenticationFactorFinishAuthenticationStatus_Completed    SecondaryAuthenticationFactorFinishAuthenticationStatus = 1
	SecondaryAuthenticationFactorFinishAuthenticationStatus_NonceExpired SecondaryAuthenticationFactorFinishAuthenticationStatus = 2
)

// enum
type SecondaryAuthenticationFactorRegistrationStatus int32

const (
	SecondaryAuthenticationFactorRegistrationStatus_Failed           SecondaryAuthenticationFactorRegistrationStatus = 0
	SecondaryAuthenticationFactorRegistrationStatus_Started          SecondaryAuthenticationFactorRegistrationStatus = 1
	SecondaryAuthenticationFactorRegistrationStatus_CanceledByUser   SecondaryAuthenticationFactorRegistrationStatus = 2
	SecondaryAuthenticationFactorRegistrationStatus_PinSetupRequired SecondaryAuthenticationFactorRegistrationStatus = 3
	SecondaryAuthenticationFactorRegistrationStatus_DisabledByPolicy SecondaryAuthenticationFactorRegistrationStatus = 4
)

// interfaces

// 020A16E5-6A25-40A3-8C00-50A023F619D1
var IID_ISecondaryAuthenticationFactorAuthentication = syscall.GUID{0x020A16E5, 0x6A25, 0x40A3,
	[8]byte{0x8C, 0x00, 0x50, 0xA0, 0x23, 0xF6, 0x19, 0xD1}}

type ISecondaryAuthenticationFactorAuthenticationInterface interface {
	win32.IInspectableInterface
	Get_ServiceAuthenticationHmac() *IBuffer
	Get_SessionNonce() *IBuffer
	Get_DeviceNonce() *IBuffer
	Get_DeviceConfigurationData() *IBuffer
	FinishAuthenticationAsync(deviceHmac *IBuffer, sessionHmac *IBuffer) *IAsyncOperation[SecondaryAuthenticationFactorFinishAuthenticationStatus]
	AbortAuthenticationAsync(errorLogMessage string) *IAsyncAction
}

type ISecondaryAuthenticationFactorAuthenticationVtbl struct {
	win32.IInspectableVtbl
	Get_ServiceAuthenticationHmac uintptr
	Get_SessionNonce              uintptr
	Get_DeviceNonce               uintptr
	Get_DeviceConfigurationData   uintptr
	FinishAuthenticationAsync     uintptr
	AbortAuthenticationAsync      uintptr
}

type ISecondaryAuthenticationFactorAuthentication struct {
	win32.IInspectable
}

func (this *ISecondaryAuthenticationFactorAuthentication) Vtbl() *ISecondaryAuthenticationFactorAuthenticationVtbl {
	return (*ISecondaryAuthenticationFactorAuthenticationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISecondaryAuthenticationFactorAuthentication) Get_ServiceAuthenticationHmac() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServiceAuthenticationHmac, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISecondaryAuthenticationFactorAuthentication) Get_SessionNonce() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SessionNonce, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISecondaryAuthenticationFactorAuthentication) Get_DeviceNonce() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceNonce, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISecondaryAuthenticationFactorAuthentication) Get_DeviceConfigurationData() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceConfigurationData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISecondaryAuthenticationFactorAuthentication) FinishAuthenticationAsync(deviceHmac *IBuffer, sessionHmac *IBuffer) *IAsyncOperation[SecondaryAuthenticationFactorFinishAuthenticationStatus] {
	var _result *IAsyncOperation[SecondaryAuthenticationFactorFinishAuthenticationStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FinishAuthenticationAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(deviceHmac)), uintptr(unsafe.Pointer(sessionHmac)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISecondaryAuthenticationFactorAuthentication) AbortAuthenticationAsync(errorLogMessage string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AbortAuthenticationAsync, uintptr(unsafe.Pointer(this)), NewHStr(errorLogMessage).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9CBB5987-EF6D-4BC2-BF49-4617515A0F9A
var IID_ISecondaryAuthenticationFactorAuthenticationResult = syscall.GUID{0x9CBB5987, 0xEF6D, 0x4BC2,
	[8]byte{0xBF, 0x49, 0x46, 0x17, 0x51, 0x5A, 0x0F, 0x9A}}

type ISecondaryAuthenticationFactorAuthenticationResultInterface interface {
	win32.IInspectableInterface
	Get_Status() SecondaryAuthenticationFactorAuthenticationStatus
	Get_Authentication() *ISecondaryAuthenticationFactorAuthentication
}

type ISecondaryAuthenticationFactorAuthenticationResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status         uintptr
	Get_Authentication uintptr
}

type ISecondaryAuthenticationFactorAuthenticationResult struct {
	win32.IInspectable
}

func (this *ISecondaryAuthenticationFactorAuthenticationResult) Vtbl() *ISecondaryAuthenticationFactorAuthenticationResultVtbl {
	return (*ISecondaryAuthenticationFactorAuthenticationResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISecondaryAuthenticationFactorAuthenticationResult) Get_Status() SecondaryAuthenticationFactorAuthenticationStatus {
	var _result SecondaryAuthenticationFactorAuthenticationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISecondaryAuthenticationFactorAuthenticationResult) Get_Authentication() *ISecondaryAuthenticationFactorAuthentication {
	var _result *ISecondaryAuthenticationFactorAuthentication
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Authentication, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D4A5EE56-7291-4073-BC1F-CCB8F5AFDF96
var IID_ISecondaryAuthenticationFactorAuthenticationStageChangedEventArgs = syscall.GUID{0xD4A5EE56, 0x7291, 0x4073,
	[8]byte{0xBC, 0x1F, 0xCC, 0xB8, 0xF5, 0xAF, 0xDF, 0x96}}

type ISecondaryAuthenticationFactorAuthenticationStageChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_StageInfo() *ISecondaryAuthenticationFactorAuthenticationStageInfo
}

type ISecondaryAuthenticationFactorAuthenticationStageChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_StageInfo uintptr
}

type ISecondaryAuthenticationFactorAuthenticationStageChangedEventArgs struct {
	win32.IInspectable
}

func (this *ISecondaryAuthenticationFactorAuthenticationStageChangedEventArgs) Vtbl() *ISecondaryAuthenticationFactorAuthenticationStageChangedEventArgsVtbl {
	return (*ISecondaryAuthenticationFactorAuthenticationStageChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISecondaryAuthenticationFactorAuthenticationStageChangedEventArgs) Get_StageInfo() *ISecondaryAuthenticationFactorAuthenticationStageInfo {
	var _result *ISecondaryAuthenticationFactorAuthenticationStageInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StageInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 56FEC28B-E8AA-4C0F-8E4C-A559E73ADD88
var IID_ISecondaryAuthenticationFactorAuthenticationStageInfo = syscall.GUID{0x56FEC28B, 0xE8AA, 0x4C0F,
	[8]byte{0x8E, 0x4C, 0xA5, 0x59, 0xE7, 0x3A, 0xDD, 0x88}}

type ISecondaryAuthenticationFactorAuthenticationStageInfoInterface interface {
	win32.IInspectableInterface
	Get_Stage() SecondaryAuthenticationFactorAuthenticationStage
	Get_Scenario() SecondaryAuthenticationFactorAuthenticationScenario
	Get_DeviceId() string
}

type ISecondaryAuthenticationFactorAuthenticationStageInfoVtbl struct {
	win32.IInspectableVtbl
	Get_Stage    uintptr
	Get_Scenario uintptr
	Get_DeviceId uintptr
}

type ISecondaryAuthenticationFactorAuthenticationStageInfo struct {
	win32.IInspectable
}

func (this *ISecondaryAuthenticationFactorAuthenticationStageInfo) Vtbl() *ISecondaryAuthenticationFactorAuthenticationStageInfoVtbl {
	return (*ISecondaryAuthenticationFactorAuthenticationStageInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISecondaryAuthenticationFactorAuthenticationStageInfo) Get_Stage() SecondaryAuthenticationFactorAuthenticationStage {
	var _result SecondaryAuthenticationFactorAuthenticationStage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Stage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISecondaryAuthenticationFactorAuthenticationStageInfo) Get_Scenario() SecondaryAuthenticationFactorAuthenticationScenario {
	var _result SecondaryAuthenticationFactorAuthenticationScenario
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Scenario, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISecondaryAuthenticationFactorAuthenticationStageInfo) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 3F582656-28F8-4E0F-AE8C-5898B9AE2469
var IID_ISecondaryAuthenticationFactorAuthenticationStatics = syscall.GUID{0x3F582656, 0x28F8, 0x4E0F,
	[8]byte{0xAE, 0x8C, 0x58, 0x98, 0xB9, 0xAE, 0x24, 0x69}}

type ISecondaryAuthenticationFactorAuthenticationStaticsInterface interface {
	win32.IInspectableInterface
	ShowNotificationMessageAsync(deviceName string, message SecondaryAuthenticationFactorAuthenticationMessage) *IAsyncAction
	StartAuthenticationAsync(deviceId string, serviceAuthenticationNonce *IBuffer) *IAsyncOperation[*ISecondaryAuthenticationFactorAuthenticationResult]
	Add_AuthenticationStageChanged(handler EventHandler[*ISecondaryAuthenticationFactorAuthenticationStageChangedEventArgs]) EventRegistrationToken
	Remove_AuthenticationStageChanged(token EventRegistrationToken)
	GetAuthenticationStageInfoAsync() *IAsyncOperation[*ISecondaryAuthenticationFactorAuthenticationStageInfo]
}

type ISecondaryAuthenticationFactorAuthenticationStaticsVtbl struct {
	win32.IInspectableVtbl
	ShowNotificationMessageAsync      uintptr
	StartAuthenticationAsync          uintptr
	Add_AuthenticationStageChanged    uintptr
	Remove_AuthenticationStageChanged uintptr
	GetAuthenticationStageInfoAsync   uintptr
}

type ISecondaryAuthenticationFactorAuthenticationStatics struct {
	win32.IInspectable
}

func (this *ISecondaryAuthenticationFactorAuthenticationStatics) Vtbl() *ISecondaryAuthenticationFactorAuthenticationStaticsVtbl {
	return (*ISecondaryAuthenticationFactorAuthenticationStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISecondaryAuthenticationFactorAuthenticationStatics) ShowNotificationMessageAsync(deviceName string, message SecondaryAuthenticationFactorAuthenticationMessage) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowNotificationMessageAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceName).Ptr, uintptr(message), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISecondaryAuthenticationFactorAuthenticationStatics) StartAuthenticationAsync(deviceId string, serviceAuthenticationNonce *IBuffer) *IAsyncOperation[*ISecondaryAuthenticationFactorAuthenticationResult] {
	var _result *IAsyncOperation[*ISecondaryAuthenticationFactorAuthenticationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartAuthenticationAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(serviceAuthenticationNonce)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISecondaryAuthenticationFactorAuthenticationStatics) Add_AuthenticationStageChanged(handler EventHandler[*ISecondaryAuthenticationFactorAuthenticationStageChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AuthenticationStageChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISecondaryAuthenticationFactorAuthenticationStatics) Remove_AuthenticationStageChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AuthenticationStageChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISecondaryAuthenticationFactorAuthenticationStatics) GetAuthenticationStageInfoAsync() *IAsyncOperation[*ISecondaryAuthenticationFactorAuthenticationStageInfo] {
	var _result *IAsyncOperation[*ISecondaryAuthenticationFactorAuthenticationStageInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAuthenticationStageInfoAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 90499A19-7EF2-4523-951C-A417A24ACF93
var IID_ISecondaryAuthenticationFactorDevicePresenceMonitoringRegistrationStatics = syscall.GUID{0x90499A19, 0x7EF2, 0x4523,
	[8]byte{0x95, 0x1C, 0xA4, 0x17, 0xA2, 0x4A, 0xCF, 0x93}}

type ISecondaryAuthenticationFactorDevicePresenceMonitoringRegistrationStaticsInterface interface {
	win32.IInspectableInterface
	RegisterDevicePresenceMonitoringAsync(deviceId string, deviceInstancePath string, monitoringMode SecondaryAuthenticationFactorDevicePresenceMonitoringMode) *IAsyncOperation[SecondaryAuthenticationFactorDevicePresenceMonitoringRegistrationStatus]
	RegisterDevicePresenceMonitoringWithNewDeviceAsync(deviceId string, deviceInstancePath string, monitoringMode SecondaryAuthenticationFactorDevicePresenceMonitoringMode, deviceFriendlyName string, deviceModelNumber string, deviceConfigurationData *IBuffer) *IAsyncOperation[SecondaryAuthenticationFactorDevicePresenceMonitoringRegistrationStatus]
	UnregisterDevicePresenceMonitoringAsync(deviceId string) *IAsyncAction
	IsDevicePresenceMonitoringSupported() bool
}

type ISecondaryAuthenticationFactorDevicePresenceMonitoringRegistrationStaticsVtbl struct {
	win32.IInspectableVtbl
	RegisterDevicePresenceMonitoringAsync              uintptr
	RegisterDevicePresenceMonitoringWithNewDeviceAsync uintptr
	UnregisterDevicePresenceMonitoringAsync            uintptr
	IsDevicePresenceMonitoringSupported                uintptr
}

type ISecondaryAuthenticationFactorDevicePresenceMonitoringRegistrationStatics struct {
	win32.IInspectable
}

func (this *ISecondaryAuthenticationFactorDevicePresenceMonitoringRegistrationStatics) Vtbl() *ISecondaryAuthenticationFactorDevicePresenceMonitoringRegistrationStaticsVtbl {
	return (*ISecondaryAuthenticationFactorDevicePresenceMonitoringRegistrationStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISecondaryAuthenticationFactorDevicePresenceMonitoringRegistrationStatics) RegisterDevicePresenceMonitoringAsync(deviceId string, deviceInstancePath string, monitoringMode SecondaryAuthenticationFactorDevicePresenceMonitoringMode) *IAsyncOperation[SecondaryAuthenticationFactorDevicePresenceMonitoringRegistrationStatus] {
	var _result *IAsyncOperation[SecondaryAuthenticationFactorDevicePresenceMonitoringRegistrationStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterDevicePresenceMonitoringAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, NewHStr(deviceInstancePath).Ptr, uintptr(monitoringMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISecondaryAuthenticationFactorDevicePresenceMonitoringRegistrationStatics) RegisterDevicePresenceMonitoringWithNewDeviceAsync(deviceId string, deviceInstancePath string, monitoringMode SecondaryAuthenticationFactorDevicePresenceMonitoringMode, deviceFriendlyName string, deviceModelNumber string, deviceConfigurationData *IBuffer) *IAsyncOperation[SecondaryAuthenticationFactorDevicePresenceMonitoringRegistrationStatus] {
	var _result *IAsyncOperation[SecondaryAuthenticationFactorDevicePresenceMonitoringRegistrationStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterDevicePresenceMonitoringWithNewDeviceAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, NewHStr(deviceInstancePath).Ptr, uintptr(monitoringMode), NewHStr(deviceFriendlyName).Ptr, NewHStr(deviceModelNumber).Ptr, uintptr(unsafe.Pointer(deviceConfigurationData)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISecondaryAuthenticationFactorDevicePresenceMonitoringRegistrationStatics) UnregisterDevicePresenceMonitoringAsync(deviceId string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UnregisterDevicePresenceMonitoringAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISecondaryAuthenticationFactorDevicePresenceMonitoringRegistrationStatics) IsDevicePresenceMonitoringSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsDevicePresenceMonitoringSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 1E2BA861-8533-4FCE-839B-ECB72410AC14
var IID_ISecondaryAuthenticationFactorInfo = syscall.GUID{0x1E2BA861, 0x8533, 0x4FCE,
	[8]byte{0x83, 0x9B, 0xEC, 0xB7, 0x24, 0x10, 0xAC, 0x14}}

type ISecondaryAuthenticationFactorInfoInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_DeviceFriendlyName() string
	Get_DeviceModelNumber() string
	Get_DeviceConfigurationData() *IBuffer
}

type ISecondaryAuthenticationFactorInfoVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId                uintptr
	Get_DeviceFriendlyName      uintptr
	Get_DeviceModelNumber       uintptr
	Get_DeviceConfigurationData uintptr
}

type ISecondaryAuthenticationFactorInfo struct {
	win32.IInspectable
}

func (this *ISecondaryAuthenticationFactorInfo) Vtbl() *ISecondaryAuthenticationFactorInfoVtbl {
	return (*ISecondaryAuthenticationFactorInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISecondaryAuthenticationFactorInfo) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISecondaryAuthenticationFactorInfo) Get_DeviceFriendlyName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceFriendlyName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISecondaryAuthenticationFactorInfo) Get_DeviceModelNumber() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceModelNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISecondaryAuthenticationFactorInfo) Get_DeviceConfigurationData() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceConfigurationData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 14D981A3-FC26-4FF7-ABC3-48E82A512A0A
var IID_ISecondaryAuthenticationFactorInfo2 = syscall.GUID{0x14D981A3, 0xFC26, 0x4FF7,
	[8]byte{0xAB, 0xC3, 0x48, 0xE8, 0x2A, 0x51, 0x2A, 0x0A}}

type ISecondaryAuthenticationFactorInfo2Interface interface {
	win32.IInspectableInterface
	Get_PresenceMonitoringMode() SecondaryAuthenticationFactorDevicePresenceMonitoringMode
	UpdateDevicePresenceAsync(presenceState SecondaryAuthenticationFactorDevicePresence) *IAsyncAction
	Get_IsAuthenticationSupported() bool
}

type ISecondaryAuthenticationFactorInfo2Vtbl struct {
	win32.IInspectableVtbl
	Get_PresenceMonitoringMode    uintptr
	UpdateDevicePresenceAsync     uintptr
	Get_IsAuthenticationSupported uintptr
}

type ISecondaryAuthenticationFactorInfo2 struct {
	win32.IInspectable
}

func (this *ISecondaryAuthenticationFactorInfo2) Vtbl() *ISecondaryAuthenticationFactorInfo2Vtbl {
	return (*ISecondaryAuthenticationFactorInfo2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISecondaryAuthenticationFactorInfo2) Get_PresenceMonitoringMode() SecondaryAuthenticationFactorDevicePresenceMonitoringMode {
	var _result SecondaryAuthenticationFactorDevicePresenceMonitoringMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PresenceMonitoringMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISecondaryAuthenticationFactorInfo2) UpdateDevicePresenceAsync(presenceState SecondaryAuthenticationFactorDevicePresence) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UpdateDevicePresenceAsync, uintptr(unsafe.Pointer(this)), uintptr(presenceState), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISecondaryAuthenticationFactorInfo2) Get_IsAuthenticationSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsAuthenticationSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 9F4CBBB4-8CBA-48B0-840D-DBB22A54C678
var IID_ISecondaryAuthenticationFactorRegistration = syscall.GUID{0x9F4CBBB4, 0x8CBA, 0x48B0,
	[8]byte{0x84, 0x0D, 0xDB, 0xB2, 0x2A, 0x54, 0xC6, 0x78}}

type ISecondaryAuthenticationFactorRegistrationInterface interface {
	win32.IInspectableInterface
	FinishRegisteringDeviceAsync(deviceConfigurationData *IBuffer) *IAsyncAction
	AbortRegisteringDeviceAsync(errorLogMessage string) *IAsyncAction
}

type ISecondaryAuthenticationFactorRegistrationVtbl struct {
	win32.IInspectableVtbl
	FinishRegisteringDeviceAsync uintptr
	AbortRegisteringDeviceAsync  uintptr
}

type ISecondaryAuthenticationFactorRegistration struct {
	win32.IInspectable
}

func (this *ISecondaryAuthenticationFactorRegistration) Vtbl() *ISecondaryAuthenticationFactorRegistrationVtbl {
	return (*ISecondaryAuthenticationFactorRegistrationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISecondaryAuthenticationFactorRegistration) FinishRegisteringDeviceAsync(deviceConfigurationData *IBuffer) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FinishRegisteringDeviceAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(deviceConfigurationData)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISecondaryAuthenticationFactorRegistration) AbortRegisteringDeviceAsync(errorLogMessage string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AbortRegisteringDeviceAsync, uintptr(unsafe.Pointer(this)), NewHStr(errorLogMessage).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A4FE35F0-ADE3-4981-AF6B-EC195921682A
var IID_ISecondaryAuthenticationFactorRegistrationResult = syscall.GUID{0xA4FE35F0, 0xADE3, 0x4981,
	[8]byte{0xAF, 0x6B, 0xEC, 0x19, 0x59, 0x21, 0x68, 0x2A}}

type ISecondaryAuthenticationFactorRegistrationResultInterface interface {
	win32.IInspectableInterface
	Get_Status() SecondaryAuthenticationFactorRegistrationStatus
	Get_Registration() *ISecondaryAuthenticationFactorRegistration
}

type ISecondaryAuthenticationFactorRegistrationResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status       uintptr
	Get_Registration uintptr
}

type ISecondaryAuthenticationFactorRegistrationResult struct {
	win32.IInspectable
}

func (this *ISecondaryAuthenticationFactorRegistrationResult) Vtbl() *ISecondaryAuthenticationFactorRegistrationResultVtbl {
	return (*ISecondaryAuthenticationFactorRegistrationResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISecondaryAuthenticationFactorRegistrationResult) Get_Status() SecondaryAuthenticationFactorRegistrationStatus {
	var _result SecondaryAuthenticationFactorRegistrationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISecondaryAuthenticationFactorRegistrationResult) Get_Registration() *ISecondaryAuthenticationFactorRegistration {
	var _result *ISecondaryAuthenticationFactorRegistration
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Registration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1ADF0F65-E3B7-4155-997F-B756EF65BEBA
var IID_ISecondaryAuthenticationFactorRegistrationStatics = syscall.GUID{0x1ADF0F65, 0xE3B7, 0x4155,
	[8]byte{0x99, 0x7F, 0xB7, 0x56, 0xEF, 0x65, 0xBE, 0xBA}}

type ISecondaryAuthenticationFactorRegistrationStaticsInterface interface {
	win32.IInspectableInterface
	RequestStartRegisteringDeviceAsync(deviceId string, capabilities SecondaryAuthenticationFactorDeviceCapabilities, deviceFriendlyName string, deviceModelNumber string, deviceKey *IBuffer, mutualAuthenticationKey *IBuffer) *IAsyncOperation[*ISecondaryAuthenticationFactorRegistrationResult]
	FindAllRegisteredDeviceInfoAsync(queryType SecondaryAuthenticationFactorDeviceFindScope) *IAsyncOperation[*IVectorView[*ISecondaryAuthenticationFactorInfo]]
	UnregisterDeviceAsync(deviceId string) *IAsyncAction
	UpdateDeviceConfigurationDataAsync(deviceId string, deviceConfigurationData *IBuffer) *IAsyncAction
}

type ISecondaryAuthenticationFactorRegistrationStaticsVtbl struct {
	win32.IInspectableVtbl
	RequestStartRegisteringDeviceAsync uintptr
	FindAllRegisteredDeviceInfoAsync   uintptr
	UnregisterDeviceAsync              uintptr
	UpdateDeviceConfigurationDataAsync uintptr
}

type ISecondaryAuthenticationFactorRegistrationStatics struct {
	win32.IInspectable
}

func (this *ISecondaryAuthenticationFactorRegistrationStatics) Vtbl() *ISecondaryAuthenticationFactorRegistrationStaticsVtbl {
	return (*ISecondaryAuthenticationFactorRegistrationStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISecondaryAuthenticationFactorRegistrationStatics) RequestStartRegisteringDeviceAsync(deviceId string, capabilities SecondaryAuthenticationFactorDeviceCapabilities, deviceFriendlyName string, deviceModelNumber string, deviceKey *IBuffer, mutualAuthenticationKey *IBuffer) *IAsyncOperation[*ISecondaryAuthenticationFactorRegistrationResult] {
	var _result *IAsyncOperation[*ISecondaryAuthenticationFactorRegistrationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestStartRegisteringDeviceAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(capabilities), NewHStr(deviceFriendlyName).Ptr, NewHStr(deviceModelNumber).Ptr, uintptr(unsafe.Pointer(deviceKey)), uintptr(unsafe.Pointer(mutualAuthenticationKey)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISecondaryAuthenticationFactorRegistrationStatics) FindAllRegisteredDeviceInfoAsync(queryType SecondaryAuthenticationFactorDeviceFindScope) *IAsyncOperation[*IVectorView[*ISecondaryAuthenticationFactorInfo]] {
	var _result *IAsyncOperation[*IVectorView[*ISecondaryAuthenticationFactorInfo]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllRegisteredDeviceInfoAsync, uintptr(unsafe.Pointer(this)), uintptr(queryType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISecondaryAuthenticationFactorRegistrationStatics) UnregisterDeviceAsync(deviceId string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UnregisterDeviceAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISecondaryAuthenticationFactorRegistrationStatics) UpdateDeviceConfigurationDataAsync(deviceId string, deviceConfigurationData *IBuffer) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UpdateDeviceConfigurationDataAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(deviceConfigurationData)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type SecondaryAuthenticationFactorAuthentication struct {
	RtClass
	*ISecondaryAuthenticationFactorAuthentication
}

func NewISecondaryAuthenticationFactorAuthenticationStatics() *ISecondaryAuthenticationFactorAuthenticationStatics {
	var p *ISecondaryAuthenticationFactorAuthenticationStatics
	hs := NewHStr("Windows.Security.Authentication.Identity.Provider.SecondaryAuthenticationFactorAuthentication")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISecondaryAuthenticationFactorAuthenticationStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type SecondaryAuthenticationFactorAuthenticationResult struct {
	RtClass
	*ISecondaryAuthenticationFactorAuthenticationResult
}

type SecondaryAuthenticationFactorAuthenticationStageChangedEventArgs struct {
	RtClass
	*ISecondaryAuthenticationFactorAuthenticationStageChangedEventArgs
}

type SecondaryAuthenticationFactorAuthenticationStageInfo struct {
	RtClass
	*ISecondaryAuthenticationFactorAuthenticationStageInfo
}

type SecondaryAuthenticationFactorInfo struct {
	RtClass
	*ISecondaryAuthenticationFactorInfo
}

type SecondaryAuthenticationFactorRegistration struct {
	RtClass
	*ISecondaryAuthenticationFactorRegistration
}

func NewISecondaryAuthenticationFactorDevicePresenceMonitoringRegistrationStatics() *ISecondaryAuthenticationFactorDevicePresenceMonitoringRegistrationStatics {
	var p *ISecondaryAuthenticationFactorDevicePresenceMonitoringRegistrationStatics
	hs := NewHStr("Windows.Security.Authentication.Identity.Provider.SecondaryAuthenticationFactorRegistration")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISecondaryAuthenticationFactorDevicePresenceMonitoringRegistrationStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewISecondaryAuthenticationFactorRegistrationStatics() *ISecondaryAuthenticationFactorRegistrationStatics {
	var p *ISecondaryAuthenticationFactorRegistrationStatics
	hs := NewHStr("Windows.Security.Authentication.Identity.Provider.SecondaryAuthenticationFactorRegistration")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISecondaryAuthenticationFactorRegistrationStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type SecondaryAuthenticationFactorRegistrationResult struct {
	RtClass
	*ISecondaryAuthenticationFactorRegistrationResult
}
