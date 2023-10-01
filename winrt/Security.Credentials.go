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
type KeyCredentialAttestationStatus int32

const (
	KeyCredentialAttestationStatus_Success          KeyCredentialAttestationStatus = 0
	KeyCredentialAttestationStatus_UnknownError     KeyCredentialAttestationStatus = 1
	KeyCredentialAttestationStatus_NotSupported     KeyCredentialAttestationStatus = 2
	KeyCredentialAttestationStatus_TemporaryFailure KeyCredentialAttestationStatus = 3
)

// enum
type KeyCredentialCreationOption int32

const (
	KeyCredentialCreationOption_ReplaceExisting KeyCredentialCreationOption = 0
	KeyCredentialCreationOption_FailIfExists    KeyCredentialCreationOption = 1
)

// enum
type KeyCredentialStatus int32

const (
	KeyCredentialStatus_Success                 KeyCredentialStatus = 0
	KeyCredentialStatus_UnknownError            KeyCredentialStatus = 1
	KeyCredentialStatus_NotFound                KeyCredentialStatus = 2
	KeyCredentialStatus_UserCanceled            KeyCredentialStatus = 3
	KeyCredentialStatus_UserPrefersPassword     KeyCredentialStatus = 4
	KeyCredentialStatus_CredentialAlreadyExists KeyCredentialStatus = 5
	KeyCredentialStatus_SecurityDeviceLocked    KeyCredentialStatus = 6
)

// enum
type WebAccountPictureSize int32

const (
	WebAccountPictureSize_Size64x64     WebAccountPictureSize = 64
	WebAccountPictureSize_Size208x208   WebAccountPictureSize = 208
	WebAccountPictureSize_Size424x424   WebAccountPictureSize = 424
	WebAccountPictureSize_Size1080x1080 WebAccountPictureSize = 1080
)

// enum
type WebAccountState int32

const (
	WebAccountState_None      WebAccountState = 0
	WebAccountState_Connected WebAccountState = 1
	WebAccountState_Error     WebAccountState = 2
)

// interfaces

// 54EF13A1-BF26-47B5-97DD-DE779B7CAD58
var IID_ICredentialFactory = syscall.GUID{0x54EF13A1, 0xBF26, 0x47B5,
	[8]byte{0x97, 0xDD, 0xDE, 0x77, 0x9B, 0x7C, 0xAD, 0x58}}

type ICredentialFactoryInterface interface {
	win32.IInspectableInterface
	CreatePasswordCredential(resource string, userName string, password string) *IPasswordCredential
}

type ICredentialFactoryVtbl struct {
	win32.IInspectableVtbl
	CreatePasswordCredential uintptr
}

type ICredentialFactory struct {
	win32.IInspectable
}

func (this *ICredentialFactory) Vtbl() *ICredentialFactoryVtbl {
	return (*ICredentialFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICredentialFactory) CreatePasswordCredential(resource string, userName string, password string) *IPasswordCredential {
	var _result *IPasswordCredential
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreatePasswordCredential, uintptr(unsafe.Pointer(this)), NewHStr(resource).Ptr, NewHStr(userName).Ptr, NewHStr(password).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9585EF8D-457B-4847-B11A-FA960BBDB138
var IID_IKeyCredential = syscall.GUID{0x9585EF8D, 0x457B, 0x4847,
	[8]byte{0xB1, 0x1A, 0xFA, 0x96, 0x0B, 0xBD, 0xB1, 0x38}}

type IKeyCredentialInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	RetrievePublicKeyWithDefaultBlobType() *IBuffer
	RetrievePublicKeyWithBlobType(blobType CryptographicPublicKeyBlobType) *IBuffer
	RequestSignAsync(data *IBuffer) *IAsyncOperation[*IKeyCredentialOperationResult]
	GetAttestationAsync() *IAsyncOperation[*IKeyCredentialAttestationResult]
}

type IKeyCredentialVtbl struct {
	win32.IInspectableVtbl
	Get_Name                             uintptr
	RetrievePublicKeyWithDefaultBlobType uintptr
	RetrievePublicKeyWithBlobType        uintptr
	RequestSignAsync                     uintptr
	GetAttestationAsync                  uintptr
}

type IKeyCredential struct {
	win32.IInspectable
}

func (this *IKeyCredential) Vtbl() *IKeyCredentialVtbl {
	return (*IKeyCredentialVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKeyCredential) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyCredential) RetrievePublicKeyWithDefaultBlobType() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RetrievePublicKeyWithDefaultBlobType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKeyCredential) RetrievePublicKeyWithBlobType(blobType CryptographicPublicKeyBlobType) *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RetrievePublicKeyWithBlobType, uintptr(unsafe.Pointer(this)), uintptr(blobType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKeyCredential) RequestSignAsync(data *IBuffer) *IAsyncOperation[*IKeyCredentialOperationResult] {
	var _result *IAsyncOperation[*IKeyCredentialOperationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestSignAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKeyCredential) GetAttestationAsync() *IAsyncOperation[*IKeyCredentialAttestationResult] {
	var _result *IAsyncOperation[*IKeyCredentialAttestationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAttestationAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 78AAB3A1-A3C1-4103-B6CC-472C44171CBB
var IID_IKeyCredentialAttestationResult = syscall.GUID{0x78AAB3A1, 0xA3C1, 0x4103,
	[8]byte{0xB6, 0xCC, 0x47, 0x2C, 0x44, 0x17, 0x1C, 0xBB}}

type IKeyCredentialAttestationResultInterface interface {
	win32.IInspectableInterface
	Get_CertificateChainBuffer() *IBuffer
	Get_AttestationBuffer() *IBuffer
	Get_Status() KeyCredentialAttestationStatus
}

type IKeyCredentialAttestationResultVtbl struct {
	win32.IInspectableVtbl
	Get_CertificateChainBuffer uintptr
	Get_AttestationBuffer      uintptr
	Get_Status                 uintptr
}

type IKeyCredentialAttestationResult struct {
	win32.IInspectable
}

func (this *IKeyCredentialAttestationResult) Vtbl() *IKeyCredentialAttestationResultVtbl {
	return (*IKeyCredentialAttestationResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKeyCredentialAttestationResult) Get_CertificateChainBuffer() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CertificateChainBuffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKeyCredentialAttestationResult) Get_AttestationBuffer() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AttestationBuffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKeyCredentialAttestationResult) Get_Status() KeyCredentialAttestationStatus {
	var _result KeyCredentialAttestationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 6AAC468B-0EF1-4CE0-8290-4106DA6A63B5
var IID_IKeyCredentialManagerStatics = syscall.GUID{0x6AAC468B, 0x0EF1, 0x4CE0,
	[8]byte{0x82, 0x90, 0x41, 0x06, 0xDA, 0x6A, 0x63, 0xB5}}

type IKeyCredentialManagerStaticsInterface interface {
	win32.IInspectableInterface
	IsSupportedAsync() *IAsyncOperation[bool]
	RenewAttestationAsync() *IAsyncAction
	RequestCreateAsync(name string, option KeyCredentialCreationOption) *IAsyncOperation[*IKeyCredentialRetrievalResult]
	OpenAsync(name string) *IAsyncOperation[*IKeyCredentialRetrievalResult]
	DeleteAsync(name string) *IAsyncAction
}

type IKeyCredentialManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	IsSupportedAsync      uintptr
	RenewAttestationAsync uintptr
	RequestCreateAsync    uintptr
	OpenAsync             uintptr
	DeleteAsync           uintptr
}

type IKeyCredentialManagerStatics struct {
	win32.IInspectable
}

func (this *IKeyCredentialManagerStatics) Vtbl() *IKeyCredentialManagerStaticsVtbl {
	return (*IKeyCredentialManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKeyCredentialManagerStatics) IsSupportedAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsSupportedAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKeyCredentialManagerStatics) RenewAttestationAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RenewAttestationAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKeyCredentialManagerStatics) RequestCreateAsync(name string, option KeyCredentialCreationOption) *IAsyncOperation[*IKeyCredentialRetrievalResult] {
	var _result *IAsyncOperation[*IKeyCredentialRetrievalResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestCreateAsync, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(option), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKeyCredentialManagerStatics) OpenAsync(name string) *IAsyncOperation[*IKeyCredentialRetrievalResult] {
	var _result *IAsyncOperation[*IKeyCredentialRetrievalResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenAsync, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKeyCredentialManagerStatics) DeleteAsync(name string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DeleteAsync, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F53786C1-5261-4CDD-976D-CC909AC71620
var IID_IKeyCredentialOperationResult = syscall.GUID{0xF53786C1, 0x5261, 0x4CDD,
	[8]byte{0x97, 0x6D, 0xCC, 0x90, 0x9A, 0xC7, 0x16, 0x20}}

type IKeyCredentialOperationResultInterface interface {
	win32.IInspectableInterface
	Get_Result() *IBuffer
	Get_Status() KeyCredentialStatus
}

type IKeyCredentialOperationResultVtbl struct {
	win32.IInspectableVtbl
	Get_Result uintptr
	Get_Status uintptr
}

type IKeyCredentialOperationResult struct {
	win32.IInspectable
}

func (this *IKeyCredentialOperationResult) Vtbl() *IKeyCredentialOperationResultVtbl {
	return (*IKeyCredentialOperationResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKeyCredentialOperationResult) Get_Result() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Result, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKeyCredentialOperationResult) Get_Status() KeyCredentialStatus {
	var _result KeyCredentialStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 58CD7703-8D87-4249-9B58-F6598CC9644E
var IID_IKeyCredentialRetrievalResult = syscall.GUID{0x58CD7703, 0x8D87, 0x4249,
	[8]byte{0x9B, 0x58, 0xF6, 0x59, 0x8C, 0xC9, 0x64, 0x4E}}

type IKeyCredentialRetrievalResultInterface interface {
	win32.IInspectableInterface
	Get_Credential() *IKeyCredential
	Get_Status() KeyCredentialStatus
}

type IKeyCredentialRetrievalResultVtbl struct {
	win32.IInspectableVtbl
	Get_Credential uintptr
	Get_Status     uintptr
}

type IKeyCredentialRetrievalResult struct {
	win32.IInspectable
}

func (this *IKeyCredentialRetrievalResult) Vtbl() *IKeyCredentialRetrievalResultVtbl {
	return (*IKeyCredentialRetrievalResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKeyCredentialRetrievalResult) Get_Credential() *IKeyCredential {
	var _result *IKeyCredential
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Credential, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKeyCredentialRetrievalResult) Get_Status() KeyCredentialStatus {
	var _result KeyCredentialStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 6AB18989-C720-41A7-A6C1-FEADB36329A0
var IID_IPasswordCredential = syscall.GUID{0x6AB18989, 0xC720, 0x41A7,
	[8]byte{0xA6, 0xC1, 0xFE, 0xAD, 0xB3, 0x63, 0x29, 0xA0}}

type IPasswordCredentialInterface interface {
	win32.IInspectableInterface
	Get_Resource() string
	Put_Resource(resource string)
	Get_UserName() string
	Put_UserName(userName string)
	Get_Password() string
	Put_Password(password string)
	RetrievePassword()
	Get_Properties() *IPropertySet
}

type IPasswordCredentialVtbl struct {
	win32.IInspectableVtbl
	Get_Resource     uintptr
	Put_Resource     uintptr
	Get_UserName     uintptr
	Put_UserName     uintptr
	Get_Password     uintptr
	Put_Password     uintptr
	RetrievePassword uintptr
	Get_Properties   uintptr
}

type IPasswordCredential struct {
	win32.IInspectable
}

func (this *IPasswordCredential) Vtbl() *IPasswordCredentialVtbl {
	return (*IPasswordCredentialVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPasswordCredential) Get_Resource() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Resource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPasswordCredential) Put_Resource(resource string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Resource, uintptr(unsafe.Pointer(this)), NewHStr(resource).Ptr)
	_ = _hr
}

func (this *IPasswordCredential) Get_UserName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UserName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPasswordCredential) Put_UserName(userName string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_UserName, uintptr(unsafe.Pointer(this)), NewHStr(userName).Ptr)
	_ = _hr
}

func (this *IPasswordCredential) Get_Password() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Password, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPasswordCredential) Put_Password(password string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Password, uintptr(unsafe.Pointer(this)), NewHStr(password).Ptr)
	_ = _hr
}

func (this *IPasswordCredential) RetrievePassword() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RetrievePassword, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IPasswordCredential) Get_Properties() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 61FD2C0B-C8D4-48C1-A54F-BC5A64205AF2
var IID_IPasswordVault = syscall.GUID{0x61FD2C0B, 0xC8D4, 0x48C1,
	[8]byte{0xA5, 0x4F, 0xBC, 0x5A, 0x64, 0x20, 0x5A, 0xF2}}

type IPasswordVaultInterface interface {
	win32.IInspectableInterface
	Add(credential *IPasswordCredential)
	Remove(credential *IPasswordCredential)
	Retrieve(resource string, userName string) *IPasswordCredential
	FindAllByResource(resource string) *IVectorView[*IPasswordCredential]
	FindAllByUserName(userName string) *IVectorView[*IPasswordCredential]
	RetrieveAll() *IVectorView[*IPasswordCredential]
}

type IPasswordVaultVtbl struct {
	win32.IInspectableVtbl
	Add               uintptr
	Remove            uintptr
	Retrieve          uintptr
	FindAllByResource uintptr
	FindAllByUserName uintptr
	RetrieveAll       uintptr
}

type IPasswordVault struct {
	win32.IInspectable
}

func (this *IPasswordVault) Vtbl() *IPasswordVaultVtbl {
	return (*IPasswordVaultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPasswordVault) Add(credential *IPasswordCredential) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(credential)))
	_ = _hr
}

func (this *IPasswordVault) Remove(credential *IPasswordCredential) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(credential)))
	_ = _hr
}

func (this *IPasswordVault) Retrieve(resource string, userName string) *IPasswordCredential {
	var _result *IPasswordCredential
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Retrieve, uintptr(unsafe.Pointer(this)), NewHStr(resource).Ptr, NewHStr(userName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPasswordVault) FindAllByResource(resource string) *IVectorView[*IPasswordCredential] {
	var _result *IVectorView[*IPasswordCredential]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllByResource, uintptr(unsafe.Pointer(this)), NewHStr(resource).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPasswordVault) FindAllByUserName(userName string) *IVectorView[*IPasswordCredential] {
	var _result *IVectorView[*IPasswordCredential]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllByUserName, uintptr(unsafe.Pointer(this)), NewHStr(userName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPasswordVault) RetrieveAll() *IVectorView[*IPasswordCredential] {
	var _result *IVectorView[*IPasswordCredential]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RetrieveAll, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 69473EB2-8031-49BE-80BB-96CB46D99ABA
var IID_IWebAccount = syscall.GUID{0x69473EB2, 0x8031, 0x49BE,
	[8]byte{0x80, 0xBB, 0x96, 0xCB, 0x46, 0xD9, 0x9A, 0xBA}}

type IWebAccountInterface interface {
	win32.IInspectableInterface
	Get_WebAccountProvider() *IWebAccountProvider
	Get_UserName() string
	Get_State() WebAccountState
}

type IWebAccountVtbl struct {
	win32.IInspectableVtbl
	Get_WebAccountProvider uintptr
	Get_UserName           uintptr
	Get_State              uintptr
}

type IWebAccount struct {
	win32.IInspectable
}

func (this *IWebAccount) Vtbl() *IWebAccountVtbl {
	return (*IWebAccountVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccount) Get_WebAccountProvider() *IWebAccountProvider {
	var _result *IWebAccountProvider
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WebAccountProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccount) Get_UserName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UserName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWebAccount) Get_State() WebAccountState {
	var _result WebAccountState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7B56D6F8-990B-4EB5-94A7-5621F3A8B824
var IID_IWebAccount2 = syscall.GUID{0x7B56D6F8, 0x990B, 0x4EB5,
	[8]byte{0x94, 0xA7, 0x56, 0x21, 0xF3, 0xA8, 0xB8, 0x24}}

type IWebAccount2Interface interface {
	win32.IInspectableInterface
	Get_Id() string
	Get_Properties() *IMapView[string, string]
	GetPictureAsync(desizedSize WebAccountPictureSize) *IAsyncOperation[*IRandomAccessStream]
	SignOutAsync() *IAsyncAction
	SignOutWithClientIdAsync(clientId string) *IAsyncAction
}

type IWebAccount2Vtbl struct {
	win32.IInspectableVtbl
	Get_Id                   uintptr
	Get_Properties           uintptr
	GetPictureAsync          uintptr
	SignOutAsync             uintptr
	SignOutWithClientIdAsync uintptr
}

type IWebAccount2 struct {
	win32.IInspectable
}

func (this *IWebAccount2) Vtbl() *IWebAccount2Vtbl {
	return (*IWebAccount2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccount2) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWebAccount2) Get_Properties() *IMapView[string, string] {
	var _result *IMapView[string, string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccount2) GetPictureAsync(desizedSize WebAccountPictureSize) *IAsyncOperation[*IRandomAccessStream] {
	var _result *IAsyncOperation[*IRandomAccessStream]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPictureAsync, uintptr(unsafe.Pointer(this)), uintptr(desizedSize), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccount2) SignOutAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SignOutAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWebAccount2) SignOutWithClientIdAsync(clientId string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SignOutWithClientIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(clientId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AC9AFB39-1DE9-4E92-B78F-0581A87F6E5C
var IID_IWebAccountFactory = syscall.GUID{0xAC9AFB39, 0x1DE9, 0x4E92,
	[8]byte{0xB7, 0x8F, 0x05, 0x81, 0xA8, 0x7F, 0x6E, 0x5C}}

type IWebAccountFactoryInterface interface {
	win32.IInspectableInterface
	CreateWebAccount(webAccountProvider *IWebAccountProvider, userName string, state WebAccountState) *IWebAccount
}

type IWebAccountFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateWebAccount uintptr
}

type IWebAccountFactory struct {
	win32.IInspectable
}

func (this *IWebAccountFactory) Vtbl() *IWebAccountFactoryVtbl {
	return (*IWebAccountFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountFactory) CreateWebAccount(webAccountProvider *IWebAccountProvider, userName string, state WebAccountState) *IWebAccount {
	var _result *IWebAccount
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWebAccount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(webAccountProvider)), NewHStr(userName).Ptr, uintptr(state), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 29DCC8C3-7AB9-4A7C-A336-B942F9DBF7C7
var IID_IWebAccountProvider = syscall.GUID{0x29DCC8C3, 0x7AB9, 0x4A7C,
	[8]byte{0xA3, 0x36, 0xB9, 0x42, 0xF9, 0xDB, 0xF7, 0xC7}}

type IWebAccountProviderInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	Get_DisplayName() string
	Get_IconUri() *IUriRuntimeClass
}

type IWebAccountProviderVtbl struct {
	win32.IInspectableVtbl
	Get_Id          uintptr
	Get_DisplayName uintptr
	Get_IconUri     uintptr
}

type IWebAccountProvider struct {
	win32.IInspectable
}

func (this *IWebAccountProvider) Vtbl() *IWebAccountProviderVtbl {
	return (*IWebAccountProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountProvider) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWebAccountProvider) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWebAccountProvider) Get_IconUri() *IUriRuntimeClass {
	var _result *IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IconUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4A01EB05-4E42-41D4-B518-E008A5163614
var IID_IWebAccountProvider2 = syscall.GUID{0x4A01EB05, 0x4E42, 0x41D4,
	[8]byte{0xB5, 0x18, 0xE0, 0x08, 0xA5, 0x16, 0x36, 0x14}}

type IWebAccountProvider2Interface interface {
	win32.IInspectableInterface
	Get_DisplayPurpose() string
	Get_Authority() string
}

type IWebAccountProvider2Vtbl struct {
	win32.IInspectableVtbl
	Get_DisplayPurpose uintptr
	Get_Authority      uintptr
}

type IWebAccountProvider2 struct {
	win32.IInspectable
}

func (this *IWebAccountProvider2) Vtbl() *IWebAccountProvider2Vtbl {
	return (*IWebAccountProvider2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountProvider2) Get_DisplayPurpose() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayPurpose, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IWebAccountProvider2) Get_Authority() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Authority, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// DA1C518B-970D-4D49-825C-F2706F8CA7FE
var IID_IWebAccountProvider3 = syscall.GUID{0xDA1C518B, 0x970D, 0x4D49,
	[8]byte{0x82, 0x5C, 0xF2, 0x70, 0x6F, 0x8C, 0xA7, 0xFE}}

type IWebAccountProvider3Interface interface {
	win32.IInspectableInterface
	Get_User() *IUser
}

type IWebAccountProvider3Vtbl struct {
	win32.IInspectableVtbl
	Get_User uintptr
}

type IWebAccountProvider3 struct {
	win32.IInspectable
}

func (this *IWebAccountProvider3) Vtbl() *IWebAccountProvider3Vtbl {
	return (*IWebAccountProvider3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountProvider3) Get_User() *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_User, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 718FD8DB-E796-4210-B74E-84D29894B080
var IID_IWebAccountProvider4 = syscall.GUID{0x718FD8DB, 0xE796, 0x4210,
	[8]byte{0xB7, 0x4E, 0x84, 0xD2, 0x98, 0x94, 0xB0, 0x80}}

type IWebAccountProvider4Interface interface {
	win32.IInspectableInterface
	Get_IsSystemProvider() bool
}

type IWebAccountProvider4Vtbl struct {
	win32.IInspectableVtbl
	Get_IsSystemProvider uintptr
}

type IWebAccountProvider4 struct {
	win32.IInspectable
}

func (this *IWebAccountProvider4) Vtbl() *IWebAccountProvider4Vtbl {
	return (*IWebAccountProvider4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountProvider4) Get_IsSystemProvider() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSystemProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 1D767DF1-E1E1-4B9A-A774-5C7C7E3BF371
var IID_IWebAccountProviderFactory = syscall.GUID{0x1D767DF1, 0xE1E1, 0x4B9A,
	[8]byte{0xA7, 0x74, 0x5C, 0x7C, 0x7E, 0x3B, 0xF3, 0x71}}

type IWebAccountProviderFactoryInterface interface {
	win32.IInspectableInterface
	CreateWebAccountProvider(id string, displayName string, iconUri *IUriRuntimeClass) *IWebAccountProvider
}

type IWebAccountProviderFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateWebAccountProvider uintptr
}

type IWebAccountProviderFactory struct {
	win32.IInspectable
}

func (this *IWebAccountProviderFactory) Vtbl() *IWebAccountProviderFactoryVtbl {
	return (*IWebAccountProviderFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWebAccountProviderFactory) CreateWebAccountProvider(id string, displayName string, iconUri *IUriRuntimeClass) *IWebAccountProvider {
	var _result *IWebAccountProvider
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWebAccountProvider, uintptr(unsafe.Pointer(this)), NewHStr(id).Ptr, NewHStr(displayName).Ptr, uintptr(unsafe.Pointer(iconUri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type KeyCredential struct {
	RtClass
	*IKeyCredential
}

type KeyCredentialAttestationResult struct {
	RtClass
	*IKeyCredentialAttestationResult
}

type KeyCredentialManager struct {
	RtClass
}

func NewIKeyCredentialManagerStatics() *IKeyCredentialManagerStatics {
	var p *IKeyCredentialManagerStatics
	hs := NewHStr("Windows.Security.Credentials.KeyCredentialManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IKeyCredentialManagerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type KeyCredentialOperationResult struct {
	RtClass
	*IKeyCredentialOperationResult
}

type KeyCredentialRetrievalResult struct {
	RtClass
	*IKeyCredentialRetrievalResult
}

type PasswordCredential struct {
	RtClass
	*IPasswordCredential
}

func NewPasswordCredential() *PasswordCredential {
	hs := NewHStr("Windows.Security.Credentials.PasswordCredential")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &PasswordCredential{
		RtClass:             RtClass{PInspect: p},
		IPasswordCredential: (*IPasswordCredential)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewPasswordCredential_CreatePasswordCredential(resource string, userName string, password string) *PasswordCredential {
	hs := NewHStr("Windows.Security.Credentials.PasswordCredential")
	var pFac *ICredentialFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICredentialFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IPasswordCredential
	p = pFac.CreatePasswordCredential(resource, userName, password)
	result := &PasswordCredential{
		RtClass:             RtClass{PInspect: &p.IInspectable},
		IPasswordCredential: p,
	}
	com.AddToScope(result)
	return result
}

type PasswordCredentialPropertyStore struct {
	RtClass
	*IPropertySet
}

func NewPasswordCredentialPropertyStore() *PasswordCredentialPropertyStore {
	hs := NewHStr("Windows.Security.Credentials.PasswordCredentialPropertyStore")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &PasswordCredentialPropertyStore{
		RtClass:      RtClass{PInspect: p},
		IPropertySet: (*IPropertySet)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type PasswordVault struct {
	RtClass
	*IPasswordVault
}

func NewPasswordVault() *PasswordVault {
	hs := NewHStr("Windows.Security.Credentials.PasswordVault")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &PasswordVault{
		RtClass:        RtClass{PInspect: p},
		IPasswordVault: (*IPasswordVault)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type WebAccount struct {
	RtClass
	*IWebAccount
}

func NewWebAccount_CreateWebAccount(webAccountProvider *IWebAccountProvider, userName string, state WebAccountState) *WebAccount {
	hs := NewHStr("Windows.Security.Credentials.WebAccount")
	var pFac *IWebAccountFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWebAccountFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IWebAccount
	p = pFac.CreateWebAccount(webAccountProvider, userName, state)
	result := &WebAccount{
		RtClass:     RtClass{PInspect: &p.IInspectable},
		IWebAccount: p,
	}
	com.AddToScope(result)
	return result
}

type WebAccountProvider struct {
	RtClass
	*IWebAccountProvider
}

func NewWebAccountProvider_CreateWebAccountProvider(id string, displayName string, iconUri *IUriRuntimeClass) *WebAccountProvider {
	hs := NewHStr("Windows.Security.Credentials.WebAccountProvider")
	var pFac *IWebAccountProviderFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IWebAccountProviderFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IWebAccountProvider
	p = pFac.CreateWebAccountProvider(id, displayName, iconUri)
	result := &WebAccountProvider{
		RtClass:             RtClass{PInspect: &p.IInspectable},
		IWebAccountProvider: p,
	}
	com.AddToScope(result)
	return result
}
