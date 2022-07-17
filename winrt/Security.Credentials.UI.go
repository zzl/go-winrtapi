package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type AuthenticationProtocol int32

const (
	AuthenticationProtocol_Basic     AuthenticationProtocol = 0
	AuthenticationProtocol_Digest    AuthenticationProtocol = 1
	AuthenticationProtocol_Ntlm      AuthenticationProtocol = 2
	AuthenticationProtocol_Kerberos  AuthenticationProtocol = 3
	AuthenticationProtocol_Negotiate AuthenticationProtocol = 4
	AuthenticationProtocol_CredSsp   AuthenticationProtocol = 5
	AuthenticationProtocol_Custom    AuthenticationProtocol = 6
)

// enum
type CredentialSaveOption int32

const (
	CredentialSaveOption_Unselected CredentialSaveOption = 0
	CredentialSaveOption_Selected   CredentialSaveOption = 1
	CredentialSaveOption_Hidden     CredentialSaveOption = 2
)

// enum
type UserConsentVerificationResult int32

const (
	UserConsentVerificationResult_Verified             UserConsentVerificationResult = 0
	UserConsentVerificationResult_DeviceNotPresent     UserConsentVerificationResult = 1
	UserConsentVerificationResult_NotConfiguredForUser UserConsentVerificationResult = 2
	UserConsentVerificationResult_DisabledByPolicy     UserConsentVerificationResult = 3
	UserConsentVerificationResult_DeviceBusy           UserConsentVerificationResult = 4
	UserConsentVerificationResult_RetriesExhausted     UserConsentVerificationResult = 5
	UserConsentVerificationResult_Canceled             UserConsentVerificationResult = 6
)

// enum
type UserConsentVerifierAvailability int32

const (
	UserConsentVerifierAvailability_Available            UserConsentVerifierAvailability = 0
	UserConsentVerifierAvailability_DeviceNotPresent     UserConsentVerifierAvailability = 1
	UserConsentVerifierAvailability_NotConfiguredForUser UserConsentVerifierAvailability = 2
	UserConsentVerifierAvailability_DisabledByPolicy     UserConsentVerifierAvailability = 3
	UserConsentVerifierAvailability_DeviceBusy           UserConsentVerifierAvailability = 4
)

// interfaces

// 965A0B4C-95FA-467F-992B-0B22E5859BF6
var IID_ICredentialPickerOptions = syscall.GUID{0x965A0B4C, 0x95FA, 0x467F,
	[8]byte{0x99, 0x2B, 0x0B, 0x22, 0xE5, 0x85, 0x9B, 0xF6}}

type ICredentialPickerOptionsInterface interface {
	win32.IInspectableInterface
	Put_Caption(value string)
	Get_Caption() string
	Put_Message(value string)
	Get_Message() string
	Put_ErrorCode(value uint32)
	Get_ErrorCode() uint32
	Put_TargetName(value string)
	Get_TargetName() string
	Put_AuthenticationProtocol(value AuthenticationProtocol)
	Get_AuthenticationProtocol() AuthenticationProtocol
	Put_CustomAuthenticationProtocol(value string)
	Get_CustomAuthenticationProtocol() string
	Put_PreviousCredential(value *IBuffer)
	Get_PreviousCredential() *IBuffer
	Put_AlwaysDisplayDialog(value bool)
	Get_AlwaysDisplayDialog() bool
	Put_CallerSavesCredential(value bool)
	Get_CallerSavesCredential() bool
	Put_CredentialSaveOption(value CredentialSaveOption)
	Get_CredentialSaveOption() CredentialSaveOption
}

type ICredentialPickerOptionsVtbl struct {
	win32.IInspectableVtbl
	Put_Caption                      uintptr
	Get_Caption                      uintptr
	Put_Message                      uintptr
	Get_Message                      uintptr
	Put_ErrorCode                    uintptr
	Get_ErrorCode                    uintptr
	Put_TargetName                   uintptr
	Get_TargetName                   uintptr
	Put_AuthenticationProtocol       uintptr
	Get_AuthenticationProtocol       uintptr
	Put_CustomAuthenticationProtocol uintptr
	Get_CustomAuthenticationProtocol uintptr
	Put_PreviousCredential           uintptr
	Get_PreviousCredential           uintptr
	Put_AlwaysDisplayDialog          uintptr
	Get_AlwaysDisplayDialog          uintptr
	Put_CallerSavesCredential        uintptr
	Get_CallerSavesCredential        uintptr
	Put_CredentialSaveOption         uintptr
	Get_CredentialSaveOption         uintptr
}

type ICredentialPickerOptions struct {
	win32.IInspectable
}

func (this *ICredentialPickerOptions) Vtbl() *ICredentialPickerOptionsVtbl {
	return (*ICredentialPickerOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICredentialPickerOptions) Put_Caption(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Caption, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICredentialPickerOptions) Get_Caption() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Caption, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICredentialPickerOptions) Put_Message(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Message, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICredentialPickerOptions) Get_Message() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Message, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICredentialPickerOptions) Put_ErrorCode(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ErrorCode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICredentialPickerOptions) Get_ErrorCode() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ErrorCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICredentialPickerOptions) Put_TargetName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TargetName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICredentialPickerOptions) Get_TargetName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TargetName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICredentialPickerOptions) Put_AuthenticationProtocol(value AuthenticationProtocol) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AuthenticationProtocol, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICredentialPickerOptions) Get_AuthenticationProtocol() AuthenticationProtocol {
	var _result AuthenticationProtocol
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AuthenticationProtocol, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICredentialPickerOptions) Put_CustomAuthenticationProtocol(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CustomAuthenticationProtocol, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICredentialPickerOptions) Get_CustomAuthenticationProtocol() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CustomAuthenticationProtocol, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICredentialPickerOptions) Put_PreviousCredential(value *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PreviousCredential, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ICredentialPickerOptions) Get_PreviousCredential() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PreviousCredential, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICredentialPickerOptions) Put_AlwaysDisplayDialog(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AlwaysDisplayDialog, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICredentialPickerOptions) Get_AlwaysDisplayDialog() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlwaysDisplayDialog, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICredentialPickerOptions) Put_CallerSavesCredential(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CallerSavesCredential, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICredentialPickerOptions) Get_CallerSavesCredential() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CallerSavesCredential, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICredentialPickerOptions) Put_CredentialSaveOption(value CredentialSaveOption) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CredentialSaveOption, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICredentialPickerOptions) Get_CredentialSaveOption() CredentialSaveOption {
	var _result CredentialSaveOption
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CredentialSaveOption, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 1948F99A-CC30-410C-9C38-CC0884C5B3D7
var IID_ICredentialPickerResults = syscall.GUID{0x1948F99A, 0xCC30, 0x410C,
	[8]byte{0x9C, 0x38, 0xCC, 0x08, 0x84, 0xC5, 0xB3, 0xD7}}

type ICredentialPickerResultsInterface interface {
	win32.IInspectableInterface
	Get_ErrorCode() uint32
	Get_CredentialSaveOption() CredentialSaveOption
	Get_CredentialSaved() bool
	Get_Credential() *IBuffer
	Get_CredentialDomainName() string
	Get_CredentialUserName() string
	Get_CredentialPassword() string
}

type ICredentialPickerResultsVtbl struct {
	win32.IInspectableVtbl
	Get_ErrorCode            uintptr
	Get_CredentialSaveOption uintptr
	Get_CredentialSaved      uintptr
	Get_Credential           uintptr
	Get_CredentialDomainName uintptr
	Get_CredentialUserName   uintptr
	Get_CredentialPassword   uintptr
}

type ICredentialPickerResults struct {
	win32.IInspectable
}

func (this *ICredentialPickerResults) Vtbl() *ICredentialPickerResultsVtbl {
	return (*ICredentialPickerResultsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICredentialPickerResults) Get_ErrorCode() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ErrorCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICredentialPickerResults) Get_CredentialSaveOption() CredentialSaveOption {
	var _result CredentialSaveOption
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CredentialSaveOption, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICredentialPickerResults) Get_CredentialSaved() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CredentialSaved, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICredentialPickerResults) Get_Credential() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Credential, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICredentialPickerResults) Get_CredentialDomainName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CredentialDomainName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICredentialPickerResults) Get_CredentialUserName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CredentialUserName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICredentialPickerResults) Get_CredentialPassword() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CredentialPassword, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// AA3A5C73-C9EA-4782-99FB-E6D7E938E12D
var IID_ICredentialPickerStatics = syscall.GUID{0xAA3A5C73, 0xC9EA, 0x4782,
	[8]byte{0x99, 0xFB, 0xE6, 0xD7, 0xE9, 0x38, 0xE1, 0x2D}}

type ICredentialPickerStaticsInterface interface {
	win32.IInspectableInterface
	PickWithOptionsAsync(options *ICredentialPickerOptions) *IAsyncOperation[*ICredentialPickerResults]
	PickWithMessageAsync(targetName string, message string) *IAsyncOperation[*ICredentialPickerResults]
	PickWithCaptionAsync(targetName string, message string, caption string) *IAsyncOperation[*ICredentialPickerResults]
}

type ICredentialPickerStaticsVtbl struct {
	win32.IInspectableVtbl
	PickWithOptionsAsync uintptr
	PickWithMessageAsync uintptr
	PickWithCaptionAsync uintptr
}

type ICredentialPickerStatics struct {
	win32.IInspectable
}

func (this *ICredentialPickerStatics) Vtbl() *ICredentialPickerStaticsVtbl {
	return (*ICredentialPickerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICredentialPickerStatics) PickWithOptionsAsync(options *ICredentialPickerOptions) *IAsyncOperation[*ICredentialPickerResults] {
	var _result *IAsyncOperation[*ICredentialPickerResults]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PickWithOptionsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICredentialPickerStatics) PickWithMessageAsync(targetName string, message string) *IAsyncOperation[*ICredentialPickerResults] {
	var _result *IAsyncOperation[*ICredentialPickerResults]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PickWithMessageAsync, uintptr(unsafe.Pointer(this)), NewHStr(targetName).Ptr, NewHStr(message).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICredentialPickerStatics) PickWithCaptionAsync(targetName string, message string, caption string) *IAsyncOperation[*ICredentialPickerResults] {
	var _result *IAsyncOperation[*ICredentialPickerResults]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PickWithCaptionAsync, uintptr(unsafe.Pointer(this)), NewHStr(targetName).Ptr, NewHStr(message).Ptr, NewHStr(caption).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AF4F3F91-564C-4DDC-B8B5-973447627C65
var IID_IUserConsentVerifierStatics = syscall.GUID{0xAF4F3F91, 0x564C, 0x4DDC,
	[8]byte{0xB8, 0xB5, 0x97, 0x34, 0x47, 0x62, 0x7C, 0x65}}

type IUserConsentVerifierStaticsInterface interface {
	win32.IInspectableInterface
	CheckAvailabilityAsync() *IAsyncOperation[UserConsentVerifierAvailability]
	RequestVerificationAsync(message string) *IAsyncOperation[UserConsentVerificationResult]
}

type IUserConsentVerifierStaticsVtbl struct {
	win32.IInspectableVtbl
	CheckAvailabilityAsync   uintptr
	RequestVerificationAsync uintptr
}

type IUserConsentVerifierStatics struct {
	win32.IInspectable
}

func (this *IUserConsentVerifierStatics) Vtbl() *IUserConsentVerifierStaticsVtbl {
	return (*IUserConsentVerifierStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUserConsentVerifierStatics) CheckAvailabilityAsync() *IAsyncOperation[UserConsentVerifierAvailability] {
	var _result *IAsyncOperation[UserConsentVerifierAvailability]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CheckAvailabilityAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserConsentVerifierStatics) RequestVerificationAsync(message string) *IAsyncOperation[UserConsentVerificationResult] {
	var _result *IAsyncOperation[UserConsentVerificationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestVerificationAsync, uintptr(unsafe.Pointer(this)), NewHStr(message).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type UserConsentVerifier struct {
	RtClass
}

func NewIUserConsentVerifierStatics() *IUserConsentVerifierStatics {
	var p *IUserConsentVerifierStatics
	hs := NewHStr("Windows.Security.Credentials.UI.UserConsentVerifier")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IUserConsentVerifierStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
