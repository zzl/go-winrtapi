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
type CertificateChainPolicy int32

const (
	CertificateChainPolicy_Base             CertificateChainPolicy = 0
	CertificateChainPolicy_Ssl              CertificateChainPolicy = 1
	CertificateChainPolicy_NTAuthentication CertificateChainPolicy = 2
	CertificateChainPolicy_MicrosoftRoot    CertificateChainPolicy = 3
)

// enum
type ChainValidationResult int32

const (
	ChainValidationResult_Success                           ChainValidationResult = 0
	ChainValidationResult_Untrusted                         ChainValidationResult = 1
	ChainValidationResult_Revoked                           ChainValidationResult = 2
	ChainValidationResult_Expired                           ChainValidationResult = 3
	ChainValidationResult_IncompleteChain                   ChainValidationResult = 4
	ChainValidationResult_InvalidSignature                  ChainValidationResult = 5
	ChainValidationResult_WrongUsage                        ChainValidationResult = 6
	ChainValidationResult_InvalidName                       ChainValidationResult = 7
	ChainValidationResult_InvalidCertificateAuthorityPolicy ChainValidationResult = 8
	ChainValidationResult_BasicConstraintsError             ChainValidationResult = 9
	ChainValidationResult_UnknownCriticalExtension          ChainValidationResult = 10
	ChainValidationResult_RevocationInformationMissing      ChainValidationResult = 11
	ChainValidationResult_RevocationFailure                 ChainValidationResult = 12
	ChainValidationResult_OtherErrors                       ChainValidationResult = 13
)

// enum
// flags
type EnrollKeyUsages uint32

const (
	EnrollKeyUsages_None         EnrollKeyUsages = 0
	EnrollKeyUsages_Decryption   EnrollKeyUsages = 1
	EnrollKeyUsages_Signing      EnrollKeyUsages = 2
	EnrollKeyUsages_KeyAgreement EnrollKeyUsages = 4
	EnrollKeyUsages_All          EnrollKeyUsages = 16777215
)

// enum
type ExportOption int32

const (
	ExportOption_NotExportable ExportOption = 0
	ExportOption_Exportable    ExportOption = 1
)

// enum
// flags
type InstallOptions uint32

const (
	InstallOptions_None          InstallOptions = 0
	InstallOptions_DeleteExpired InstallOptions = 1
)

// enum
type KeyProtectionLevel int32

const (
	KeyProtectionLevel_NoConsent              KeyProtectionLevel = 0
	KeyProtectionLevel_ConsentOnly            KeyProtectionLevel = 1
	KeyProtectionLevel_ConsentWithPassword    KeyProtectionLevel = 2
	KeyProtectionLevel_ConsentWithFingerprint KeyProtectionLevel = 3
)

// enum
type KeySize int32

const (
	KeySize_Invalid KeySize = 0
	KeySize_Rsa2048 KeySize = 2048
	KeySize_Rsa4096 KeySize = 4096
)

// enum
type SignatureValidationResult int32

const (
	SignatureValidationResult_Success          SignatureValidationResult = 0
	SignatureValidationResult_InvalidParameter SignatureValidationResult = 1
	SignatureValidationResult_BadMessage       SignatureValidationResult = 2
	SignatureValidationResult_InvalidSignature SignatureValidationResult = 3
	SignatureValidationResult_OtherErrors      SignatureValidationResult = 4
)

// interfaces

// 333F740C-04D8-43B3-B278-8C5FCC9BE5A0
var IID_ICertificate = syscall.GUID{0x333F740C, 0x04D8, 0x43B3,
	[8]byte{0xB2, 0x78, 0x8C, 0x5F, 0xCC, 0x9B, 0xE5, 0xA0}}

type ICertificateInterface interface {
	win32.IInspectableInterface
	BuildChainAsync(certificates *IIterable[*ICertificate]) *IAsyncOperation[*ICertificateChain]
	BuildChainWithParametersAsync(certificates *IIterable[*ICertificate], parameters *IChainBuildingParameters) *IAsyncOperation[*ICertificateChain]
	Get_SerialNumber() []byte
	GetHashValue() []byte
	GetHashValueWithAlgorithm(hashAlgorithmName string) []byte
	GetCertificateBlob() *IBuffer
	Get_Subject() string
	Get_Issuer() string
	Get_HasPrivateKey() bool
	Get_IsStronglyProtected() bool
	Get_ValidFrom() DateTime
	Get_ValidTo() DateTime
	Get_EnhancedKeyUsages() *IVectorView[string]
	Put_FriendlyName(value string)
	Get_FriendlyName() string
}

type ICertificateVtbl struct {
	win32.IInspectableVtbl
	BuildChainAsync               uintptr
	BuildChainWithParametersAsync uintptr
	Get_SerialNumber              uintptr
	GetHashValue                  uintptr
	GetHashValueWithAlgorithm     uintptr
	GetCertificateBlob            uintptr
	Get_Subject                   uintptr
	Get_Issuer                    uintptr
	Get_HasPrivateKey             uintptr
	Get_IsStronglyProtected       uintptr
	Get_ValidFrom                 uintptr
	Get_ValidTo                   uintptr
	Get_EnhancedKeyUsages         uintptr
	Put_FriendlyName              uintptr
	Get_FriendlyName              uintptr
}

type ICertificate struct {
	win32.IInspectable
}

func (this *ICertificate) Vtbl() *ICertificateVtbl {
	return (*ICertificateVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICertificate) BuildChainAsync(certificates *IIterable[*ICertificate]) *IAsyncOperation[*ICertificateChain] {
	var _result *IAsyncOperation[*ICertificateChain]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().BuildChainAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(certificates)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICertificate) BuildChainWithParametersAsync(certificates *IIterable[*ICertificate], parameters *IChainBuildingParameters) *IAsyncOperation[*ICertificateChain] {
	var _result *IAsyncOperation[*ICertificateChain]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().BuildChainWithParametersAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(certificates)), uintptr(unsafe.Pointer(parameters)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICertificate) Get_SerialNumber() []byte {
	var _result []byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SerialNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificate) GetHashValue() []byte {
	var _result []byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetHashValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificate) GetHashValueWithAlgorithm(hashAlgorithmName string) []byte {
	var _result []byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetHashValueWithAlgorithm, uintptr(unsafe.Pointer(this)), NewHStr(hashAlgorithmName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificate) GetCertificateBlob() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCertificateBlob, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICertificate) Get_Subject() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Subject, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICertificate) Get_Issuer() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Issuer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICertificate) Get_HasPrivateKey() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasPrivateKey, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificate) Get_IsStronglyProtected() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsStronglyProtected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificate) Get_ValidFrom() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ValidFrom, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificate) Get_ValidTo() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ValidTo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificate) Get_EnhancedKeyUsages() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EnhancedKeyUsages, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICertificate) Put_FriendlyName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FriendlyName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICertificate) Get_FriendlyName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FriendlyName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 17B8374C-8A25-4D96-A492-8FC29AC4FDA6
var IID_ICertificate2 = syscall.GUID{0x17B8374C, 0x8A25, 0x4D96,
	[8]byte{0xA4, 0x92, 0x8F, 0xC2, 0x9A, 0xC4, 0xFD, 0xA6}}

type ICertificate2Interface interface {
	win32.IInspectableInterface
	Get_IsSecurityDeviceBound() bool
	Get_KeyUsages() *ICertificateKeyUsages
	Get_KeyAlgorithmName() string
	Get_SignatureAlgorithmName() string
	Get_SignatureHashAlgorithmName() string
	Get_SubjectAlternativeName() *ISubjectAlternativeNameInfo
}

type ICertificate2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsSecurityDeviceBound      uintptr
	Get_KeyUsages                  uintptr
	Get_KeyAlgorithmName           uintptr
	Get_SignatureAlgorithmName     uintptr
	Get_SignatureHashAlgorithmName uintptr
	Get_SubjectAlternativeName     uintptr
}

type ICertificate2 struct {
	win32.IInspectable
}

func (this *ICertificate2) Vtbl() *ICertificate2Vtbl {
	return (*ICertificate2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICertificate2) Get_IsSecurityDeviceBound() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSecurityDeviceBound, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificate2) Get_KeyUsages() *ICertificateKeyUsages {
	var _result *ICertificateKeyUsages
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KeyUsages, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICertificate2) Get_KeyAlgorithmName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KeyAlgorithmName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICertificate2) Get_SignatureAlgorithmName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SignatureAlgorithmName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICertificate2) Get_SignatureHashAlgorithmName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SignatureHashAlgorithmName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICertificate2) Get_SubjectAlternativeName() *ISubjectAlternativeNameInfo {
	var _result *ISubjectAlternativeNameInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SubjectAlternativeName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BE51A966-AE5F-4652-ACE7-C6D7E7724CF3
var IID_ICertificate3 = syscall.GUID{0xBE51A966, 0xAE5F, 0x4652,
	[8]byte{0xAC, 0xE7, 0xC6, 0xD7, 0xE7, 0x72, 0x4C, 0xF3}}

type ICertificate3Interface interface {
	win32.IInspectableInterface
	Get_IsPerUser() bool
	Get_StoreName() string
	Get_KeyStorageProviderName() string
}

type ICertificate3Vtbl struct {
	win32.IInspectableVtbl
	Get_IsPerUser              uintptr
	Get_StoreName              uintptr
	Get_KeyStorageProviderName uintptr
}

type ICertificate3 struct {
	win32.IInspectable
}

func (this *ICertificate3) Vtbl() *ICertificate3Vtbl {
	return (*ICertificate3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICertificate3) Get_IsPerUser() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPerUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificate3) Get_StoreName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StoreName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICertificate3) Get_KeyStorageProviderName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KeyStorageProviderName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 20BF5385-3691-4501-A62C-FD97278B31EE
var IID_ICertificateChain = syscall.GUID{0x20BF5385, 0x3691, 0x4501,
	[8]byte{0xA6, 0x2C, 0xFD, 0x97, 0x27, 0x8B, 0x31, 0xEE}}

type ICertificateChainInterface interface {
	win32.IInspectableInterface
	Validate() ChainValidationResult
	ValidateWithParameters(parameter *IChainValidationParameters) ChainValidationResult
	GetCertificates(includeRoot bool) *IVectorView[*ICertificate]
}

type ICertificateChainVtbl struct {
	win32.IInspectableVtbl
	Validate               uintptr
	ValidateWithParameters uintptr
	GetCertificates        uintptr
}

type ICertificateChain struct {
	win32.IInspectable
}

func (this *ICertificateChain) Vtbl() *ICertificateChainVtbl {
	return (*ICertificateChainVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICertificateChain) Validate() ChainValidationResult {
	var _result ChainValidationResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Validate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificateChain) ValidateWithParameters(parameter *IChainValidationParameters) ChainValidationResult {
	var _result ChainValidationResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ValidateWithParameters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(parameter)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificateChain) GetCertificates(includeRoot bool) *IVectorView[*ICertificate] {
	var _result *IVectorView[*ICertificate]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCertificates, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&includeRoot))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8846EF3F-A986-48FB-9FD7-9AEC06935BF1
var IID_ICertificateEnrollmentManagerStatics = syscall.GUID{0x8846EF3F, 0xA986, 0x48FB,
	[8]byte{0x9F, 0xD7, 0x9A, 0xEC, 0x06, 0x93, 0x5B, 0xF1}}

type ICertificateEnrollmentManagerStaticsInterface interface {
	win32.IInspectableInterface
	CreateRequestAsync(request *ICertificateRequestProperties) *IAsyncOperation[string]
	InstallCertificateAsync(certificate string, installOption InstallOptions) *IAsyncAction
	ImportPfxDataAsync(pfxData string, password string, exportable ExportOption, keyProtectionLevel KeyProtectionLevel, installOption InstallOptions, friendlyName string) *IAsyncAction
}

type ICertificateEnrollmentManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateRequestAsync      uintptr
	InstallCertificateAsync uintptr
	ImportPfxDataAsync      uintptr
}

type ICertificateEnrollmentManagerStatics struct {
	win32.IInspectable
}

func (this *ICertificateEnrollmentManagerStatics) Vtbl() *ICertificateEnrollmentManagerStaticsVtbl {
	return (*ICertificateEnrollmentManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICertificateEnrollmentManagerStatics) CreateRequestAsync(request *ICertificateRequestProperties) *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateRequestAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(request)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICertificateEnrollmentManagerStatics) InstallCertificateAsync(certificate string, installOption InstallOptions) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InstallCertificateAsync, uintptr(unsafe.Pointer(this)), NewHStr(certificate).Ptr, uintptr(installOption), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICertificateEnrollmentManagerStatics) ImportPfxDataAsync(pfxData string, password string, exportable ExportOption, keyProtectionLevel KeyProtectionLevel, installOption InstallOptions, friendlyName string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ImportPfxDataAsync, uintptr(unsafe.Pointer(this)), NewHStr(pfxData).Ptr, NewHStr(password).Ptr, uintptr(exportable), uintptr(keyProtectionLevel), uintptr(installOption), NewHStr(friendlyName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DC5B1C33-6429-4014-999C-5D9735802D1D
var IID_ICertificateEnrollmentManagerStatics2 = syscall.GUID{0xDC5B1C33, 0x6429, 0x4014,
	[8]byte{0x99, 0x9C, 0x5D, 0x97, 0x35, 0x80, 0x2D, 0x1D}}

type ICertificateEnrollmentManagerStatics2Interface interface {
	win32.IInspectableInterface
	Get_UserCertificateEnrollmentManager() *IUserCertificateEnrollmentManager
	ImportPfxDataToKspAsync(pfxData string, password string, exportable ExportOption, keyProtectionLevel KeyProtectionLevel, installOption InstallOptions, friendlyName string, keyStorageProvider string) *IAsyncAction
}

type ICertificateEnrollmentManagerStatics2Vtbl struct {
	win32.IInspectableVtbl
	Get_UserCertificateEnrollmentManager uintptr
	ImportPfxDataToKspAsync              uintptr
}

type ICertificateEnrollmentManagerStatics2 struct {
	win32.IInspectable
}

func (this *ICertificateEnrollmentManagerStatics2) Vtbl() *ICertificateEnrollmentManagerStatics2Vtbl {
	return (*ICertificateEnrollmentManagerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICertificateEnrollmentManagerStatics2) Get_UserCertificateEnrollmentManager() *IUserCertificateEnrollmentManager {
	var _result *IUserCertificateEnrollmentManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UserCertificateEnrollmentManager, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICertificateEnrollmentManagerStatics2) ImportPfxDataToKspAsync(pfxData string, password string, exportable ExportOption, keyProtectionLevel KeyProtectionLevel, installOption InstallOptions, friendlyName string, keyStorageProvider string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ImportPfxDataToKspAsync, uintptr(unsafe.Pointer(this)), NewHStr(pfxData).Ptr, NewHStr(password).Ptr, uintptr(exportable), uintptr(keyProtectionLevel), uintptr(installOption), NewHStr(friendlyName).Ptr, NewHStr(keyStorageProvider).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FDEC82BE-617C-425A-B72D-398B26AC7264
var IID_ICertificateEnrollmentManagerStatics3 = syscall.GUID{0xFDEC82BE, 0x617C, 0x425A,
	[8]byte{0xB7, 0x2D, 0x39, 0x8B, 0x26, 0xAC, 0x72, 0x64}}

type ICertificateEnrollmentManagerStatics3Interface interface {
	win32.IInspectableInterface
	ImportPfxDataToKspWithParametersAsync(pfxData string, password string, pfxImportParameters *IPfxImportParameters) *IAsyncAction
}

type ICertificateEnrollmentManagerStatics3Vtbl struct {
	win32.IInspectableVtbl
	ImportPfxDataToKspWithParametersAsync uintptr
}

type ICertificateEnrollmentManagerStatics3 struct {
	win32.IInspectable
}

func (this *ICertificateEnrollmentManagerStatics3) Vtbl() *ICertificateEnrollmentManagerStatics3Vtbl {
	return (*ICertificateEnrollmentManagerStatics3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICertificateEnrollmentManagerStatics3) ImportPfxDataToKspWithParametersAsync(pfxData string, password string, pfxImportParameters *IPfxImportParameters) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ImportPfxDataToKspWithParametersAsync, uintptr(unsafe.Pointer(this)), NewHStr(pfxData).Ptr, NewHStr(password).Ptr, uintptr(unsafe.Pointer(pfxImportParameters)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 84CF0656-A9E6-454D-8E45-2EA7C4BCD53B
var IID_ICertificateExtension = syscall.GUID{0x84CF0656, 0xA9E6, 0x454D,
	[8]byte{0x8E, 0x45, 0x2E, 0xA7, 0xC4, 0xBC, 0xD5, 0x3B}}

type ICertificateExtensionInterface interface {
	win32.IInspectableInterface
	Get_ObjectId() string
	Put_ObjectId(value string)
	Get_IsCritical() bool
	Put_IsCritical(value bool)
	EncodeValue(value string)
	Get_Value() []byte
	Put_Value(valueLength uint32, value *byte)
}

type ICertificateExtensionVtbl struct {
	win32.IInspectableVtbl
	Get_ObjectId   uintptr
	Put_ObjectId   uintptr
	Get_IsCritical uintptr
	Put_IsCritical uintptr
	EncodeValue    uintptr
	Get_Value      uintptr
	Put_Value      uintptr
}

type ICertificateExtension struct {
	win32.IInspectable
}

func (this *ICertificateExtension) Vtbl() *ICertificateExtensionVtbl {
	return (*ICertificateExtensionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICertificateExtension) Get_ObjectId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ObjectId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICertificateExtension) Put_ObjectId(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ObjectId, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICertificateExtension) Get_IsCritical() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCritical, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificateExtension) Put_IsCritical(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsCritical, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICertificateExtension) EncodeValue(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EncodeValue, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICertificateExtension) Get_Value() []byte {
	var _result []byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificateExtension) Put_Value(valueLength uint32, value *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Value, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 17B4221C-4BAF-44A2-9608-04FB62B16942
var IID_ICertificateFactory = syscall.GUID{0x17B4221C, 0x4BAF, 0x44A2,
	[8]byte{0x96, 0x08, 0x04, 0xFB, 0x62, 0xB1, 0x69, 0x42}}

type ICertificateFactoryInterface interface {
	win32.IInspectableInterface
	CreateCertificate(certBlob *IBuffer) *ICertificate
}

type ICertificateFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateCertificate uintptr
}

type ICertificateFactory struct {
	win32.IInspectable
}

func (this *ICertificateFactory) Vtbl() *ICertificateFactoryVtbl {
	return (*ICertificateFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICertificateFactory) CreateCertificate(certBlob *IBuffer) *ICertificate {
	var _result *ICertificate
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateCertificate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(certBlob)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6AC6206F-E1CF-486A-B485-A69C83E46FD1
var IID_ICertificateKeyUsages = syscall.GUID{0x6AC6206F, 0xE1CF, 0x486A,
	[8]byte{0xB4, 0x85, 0xA6, 0x9C, 0x83, 0xE4, 0x6F, 0xD1}}

type ICertificateKeyUsagesInterface interface {
	win32.IInspectableInterface
	Get_EncipherOnly() bool
	Put_EncipherOnly(value bool)
	Get_CrlSign() bool
	Put_CrlSign(value bool)
	Get_KeyCertificateSign() bool
	Put_KeyCertificateSign(value bool)
	Get_KeyAgreement() bool
	Put_KeyAgreement(value bool)
	Get_DataEncipherment() bool
	Put_DataEncipherment(value bool)
	Get_KeyEncipherment() bool
	Put_KeyEncipherment(value bool)
	Get_NonRepudiation() bool
	Put_NonRepudiation(value bool)
	Get_DigitalSignature() bool
	Put_DigitalSignature(value bool)
}

type ICertificateKeyUsagesVtbl struct {
	win32.IInspectableVtbl
	Get_EncipherOnly       uintptr
	Put_EncipherOnly       uintptr
	Get_CrlSign            uintptr
	Put_CrlSign            uintptr
	Get_KeyCertificateSign uintptr
	Put_KeyCertificateSign uintptr
	Get_KeyAgreement       uintptr
	Put_KeyAgreement       uintptr
	Get_DataEncipherment   uintptr
	Put_DataEncipherment   uintptr
	Get_KeyEncipherment    uintptr
	Put_KeyEncipherment    uintptr
	Get_NonRepudiation     uintptr
	Put_NonRepudiation     uintptr
	Get_DigitalSignature   uintptr
	Put_DigitalSignature   uintptr
}

type ICertificateKeyUsages struct {
	win32.IInspectable
}

func (this *ICertificateKeyUsages) Vtbl() *ICertificateKeyUsagesVtbl {
	return (*ICertificateKeyUsagesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICertificateKeyUsages) Get_EncipherOnly() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EncipherOnly, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificateKeyUsages) Put_EncipherOnly(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_EncipherOnly, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICertificateKeyUsages) Get_CrlSign() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CrlSign, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificateKeyUsages) Put_CrlSign(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CrlSign, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICertificateKeyUsages) Get_KeyCertificateSign() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KeyCertificateSign, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificateKeyUsages) Put_KeyCertificateSign(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_KeyCertificateSign, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICertificateKeyUsages) Get_KeyAgreement() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KeyAgreement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificateKeyUsages) Put_KeyAgreement(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_KeyAgreement, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICertificateKeyUsages) Get_DataEncipherment() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DataEncipherment, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificateKeyUsages) Put_DataEncipherment(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DataEncipherment, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICertificateKeyUsages) Get_KeyEncipherment() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KeyEncipherment, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificateKeyUsages) Put_KeyEncipherment(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_KeyEncipherment, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICertificateKeyUsages) Get_NonRepudiation() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NonRepudiation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificateKeyUsages) Put_NonRepudiation(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_NonRepudiation, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICertificateKeyUsages) Get_DigitalSignature() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DigitalSignature, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificateKeyUsages) Put_DigitalSignature(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DigitalSignature, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 5B082A31-A728-4916-B5EE-FFCB8ACF2417
var IID_ICertificateQuery = syscall.GUID{0x5B082A31, 0xA728, 0x4916,
	[8]byte{0xB5, 0xEE, 0xFF, 0xCB, 0x8A, 0xCF, 0x24, 0x17}}

type ICertificateQueryInterface interface {
	win32.IInspectableInterface
	Get_EnhancedKeyUsages() *IVector[string]
	Get_IssuerName() string
	Put_IssuerName(value string)
	Get_FriendlyName() string
	Put_FriendlyName(value string)
	Get_Thumbprint() []byte
	Put_Thumbprint(valueLength uint32, value *byte)
	Get_HardwareOnly() bool
	Put_HardwareOnly(value bool)
}

type ICertificateQueryVtbl struct {
	win32.IInspectableVtbl
	Get_EnhancedKeyUsages uintptr
	Get_IssuerName        uintptr
	Put_IssuerName        uintptr
	Get_FriendlyName      uintptr
	Put_FriendlyName      uintptr
	Get_Thumbprint        uintptr
	Put_Thumbprint        uintptr
	Get_HardwareOnly      uintptr
	Put_HardwareOnly      uintptr
}

type ICertificateQuery struct {
	win32.IInspectable
}

func (this *ICertificateQuery) Vtbl() *ICertificateQueryVtbl {
	return (*ICertificateQueryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICertificateQuery) Get_EnhancedKeyUsages() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EnhancedKeyUsages, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICertificateQuery) Get_IssuerName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IssuerName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICertificateQuery) Put_IssuerName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IssuerName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICertificateQuery) Get_FriendlyName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FriendlyName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICertificateQuery) Put_FriendlyName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FriendlyName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICertificateQuery) Get_Thumbprint() []byte {
	var _result []byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Thumbprint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificateQuery) Put_Thumbprint(valueLength uint32, value *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Thumbprint, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ICertificateQuery) Get_HardwareOnly() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HardwareOnly, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificateQuery) Put_HardwareOnly(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_HardwareOnly, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 935A0AF7-0BD9-4F75-B8C2-E27A7F74EECD
var IID_ICertificateQuery2 = syscall.GUID{0x935A0AF7, 0x0BD9, 0x4F75,
	[8]byte{0xB8, 0xC2, 0xE2, 0x7A, 0x7F, 0x74, 0xEE, 0xCD}}

type ICertificateQuery2Interface interface {
	win32.IInspectableInterface
	Get_IncludeDuplicates() bool
	Put_IncludeDuplicates(value bool)
	Get_IncludeExpiredCertificates() bool
	Put_IncludeExpiredCertificates(value bool)
	Get_StoreName() string
	Put_StoreName(value string)
}

type ICertificateQuery2Vtbl struct {
	win32.IInspectableVtbl
	Get_IncludeDuplicates          uintptr
	Put_IncludeDuplicates          uintptr
	Get_IncludeExpiredCertificates uintptr
	Put_IncludeExpiredCertificates uintptr
	Get_StoreName                  uintptr
	Put_StoreName                  uintptr
}

type ICertificateQuery2 struct {
	win32.IInspectable
}

func (this *ICertificateQuery2) Vtbl() *ICertificateQuery2Vtbl {
	return (*ICertificateQuery2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICertificateQuery2) Get_IncludeDuplicates() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IncludeDuplicates, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificateQuery2) Put_IncludeDuplicates(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IncludeDuplicates, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICertificateQuery2) Get_IncludeExpiredCertificates() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IncludeExpiredCertificates, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificateQuery2) Put_IncludeExpiredCertificates(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IncludeExpiredCertificates, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICertificateQuery2) Get_StoreName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StoreName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICertificateQuery2) Put_StoreName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StoreName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// 487E84F6-94E2-4DCE-8833-1A700A37A29A
var IID_ICertificateRequestProperties = syscall.GUID{0x487E84F6, 0x94E2, 0x4DCE,
	[8]byte{0x88, 0x33, 0x1A, 0x70, 0x0A, 0x37, 0xA2, 0x9A}}

type ICertificateRequestPropertiesInterface interface {
	win32.IInspectableInterface
	Get_Subject() string
	Put_Subject(value string)
	Get_KeyAlgorithmName() string
	Put_KeyAlgorithmName(value string)
	Get_KeySize() uint32
	Put_KeySize(value uint32)
	Get_FriendlyName() string
	Put_FriendlyName(value string)
	Get_HashAlgorithmName() string
	Put_HashAlgorithmName(value string)
	Get_Exportable() ExportOption
	Put_Exportable(value ExportOption)
	Get_KeyUsages() EnrollKeyUsages
	Put_KeyUsages(value EnrollKeyUsages)
	Get_KeyProtectionLevel() KeyProtectionLevel
	Put_KeyProtectionLevel(value KeyProtectionLevel)
	Get_KeyStorageProviderName() string
	Put_KeyStorageProviderName(value string)
}

type ICertificateRequestPropertiesVtbl struct {
	win32.IInspectableVtbl
	Get_Subject                uintptr
	Put_Subject                uintptr
	Get_KeyAlgorithmName       uintptr
	Put_KeyAlgorithmName       uintptr
	Get_KeySize                uintptr
	Put_KeySize                uintptr
	Get_FriendlyName           uintptr
	Put_FriendlyName           uintptr
	Get_HashAlgorithmName      uintptr
	Put_HashAlgorithmName      uintptr
	Get_Exportable             uintptr
	Put_Exportable             uintptr
	Get_KeyUsages              uintptr
	Put_KeyUsages              uintptr
	Get_KeyProtectionLevel     uintptr
	Put_KeyProtectionLevel     uintptr
	Get_KeyStorageProviderName uintptr
	Put_KeyStorageProviderName uintptr
}

type ICertificateRequestProperties struct {
	win32.IInspectable
}

func (this *ICertificateRequestProperties) Vtbl() *ICertificateRequestPropertiesVtbl {
	return (*ICertificateRequestPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICertificateRequestProperties) Get_Subject() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Subject, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICertificateRequestProperties) Put_Subject(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Subject, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICertificateRequestProperties) Get_KeyAlgorithmName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KeyAlgorithmName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICertificateRequestProperties) Put_KeyAlgorithmName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_KeyAlgorithmName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICertificateRequestProperties) Get_KeySize() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KeySize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificateRequestProperties) Put_KeySize(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_KeySize, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICertificateRequestProperties) Get_FriendlyName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FriendlyName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICertificateRequestProperties) Put_FriendlyName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FriendlyName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICertificateRequestProperties) Get_HashAlgorithmName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HashAlgorithmName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICertificateRequestProperties) Put_HashAlgorithmName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_HashAlgorithmName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICertificateRequestProperties) Get_Exportable() ExportOption {
	var _result ExportOption
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Exportable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificateRequestProperties) Put_Exportable(value ExportOption) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Exportable, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICertificateRequestProperties) Get_KeyUsages() EnrollKeyUsages {
	var _result EnrollKeyUsages
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KeyUsages, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificateRequestProperties) Put_KeyUsages(value EnrollKeyUsages) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_KeyUsages, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICertificateRequestProperties) Get_KeyProtectionLevel() KeyProtectionLevel {
	var _result KeyProtectionLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KeyProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificateRequestProperties) Put_KeyProtectionLevel(value KeyProtectionLevel) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_KeyProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICertificateRequestProperties) Get_KeyStorageProviderName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KeyStorageProviderName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICertificateRequestProperties) Put_KeyStorageProviderName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_KeyStorageProviderName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// 3DA0C954-D73F-4FF3-A0A6-0677C0ADA05B
var IID_ICertificateRequestProperties2 = syscall.GUID{0x3DA0C954, 0xD73F, 0x4FF3,
	[8]byte{0xA0, 0xA6, 0x06, 0x77, 0xC0, 0xAD, 0xA0, 0x5B}}

type ICertificateRequestProperties2Interface interface {
	win32.IInspectableInterface
	Get_SmartcardReaderName() string
	Put_SmartcardReaderName(value string)
	Get_SigningCertificate() *ICertificate
	Put_SigningCertificate(value *ICertificate)
	Get_AttestationCredentialCertificate() *ICertificate
	Put_AttestationCredentialCertificate(value *ICertificate)
}

type ICertificateRequestProperties2Vtbl struct {
	win32.IInspectableVtbl
	Get_SmartcardReaderName              uintptr
	Put_SmartcardReaderName              uintptr
	Get_SigningCertificate               uintptr
	Put_SigningCertificate               uintptr
	Get_AttestationCredentialCertificate uintptr
	Put_AttestationCredentialCertificate uintptr
}

type ICertificateRequestProperties2 struct {
	win32.IInspectable
}

func (this *ICertificateRequestProperties2) Vtbl() *ICertificateRequestProperties2Vtbl {
	return (*ICertificateRequestProperties2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICertificateRequestProperties2) Get_SmartcardReaderName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SmartcardReaderName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICertificateRequestProperties2) Put_SmartcardReaderName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SmartcardReaderName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICertificateRequestProperties2) Get_SigningCertificate() *ICertificate {
	var _result *ICertificate
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SigningCertificate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICertificateRequestProperties2) Put_SigningCertificate(value *ICertificate) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SigningCertificate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ICertificateRequestProperties2) Get_AttestationCredentialCertificate() *ICertificate {
	var _result *ICertificate
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AttestationCredentialCertificate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICertificateRequestProperties2) Put_AttestationCredentialCertificate(value *ICertificate) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AttestationCredentialCertificate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// E687F616-734D-46B1-9D4C-6EDFDBFC845B
var IID_ICertificateRequestProperties3 = syscall.GUID{0xE687F616, 0x734D, 0x46B1,
	[8]byte{0x9D, 0x4C, 0x6E, 0xDF, 0xDB, 0xFC, 0x84, 0x5B}}

type ICertificateRequestProperties3Interface interface {
	win32.IInspectableInterface
	Get_CurveName() string
	Put_CurveName(value string)
	Get_CurveParameters() []byte
	Put_CurveParameters(valueLength uint32, value *byte)
	Get_ContainerNamePrefix() string
	Put_ContainerNamePrefix(value string)
	Get_ContainerName() string
	Put_ContainerName(value string)
	Get_UseExistingKey() bool
	Put_UseExistingKey(value bool)
}

type ICertificateRequestProperties3Vtbl struct {
	win32.IInspectableVtbl
	Get_CurveName           uintptr
	Put_CurveName           uintptr
	Get_CurveParameters     uintptr
	Put_CurveParameters     uintptr
	Get_ContainerNamePrefix uintptr
	Put_ContainerNamePrefix uintptr
	Get_ContainerName       uintptr
	Put_ContainerName       uintptr
	Get_UseExistingKey      uintptr
	Put_UseExistingKey      uintptr
}

type ICertificateRequestProperties3 struct {
	win32.IInspectable
}

func (this *ICertificateRequestProperties3) Vtbl() *ICertificateRequestProperties3Vtbl {
	return (*ICertificateRequestProperties3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICertificateRequestProperties3) Get_CurveName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurveName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICertificateRequestProperties3) Put_CurveName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CurveName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICertificateRequestProperties3) Get_CurveParameters() []byte {
	var _result []byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurveParameters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificateRequestProperties3) Put_CurveParameters(valueLength uint32, value *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CurveParameters, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ICertificateRequestProperties3) Get_ContainerNamePrefix() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContainerNamePrefix, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICertificateRequestProperties3) Put_ContainerNamePrefix(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ContainerNamePrefix, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICertificateRequestProperties3) Get_ContainerName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContainerName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICertificateRequestProperties3) Put_ContainerName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ContainerName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICertificateRequestProperties3) Get_UseExistingKey() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UseExistingKey, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICertificateRequestProperties3) Put_UseExistingKey(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_UseExistingKey, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 4E429AD2-1C61-4FEA-B8FE-135FB19CDCE4
var IID_ICertificateRequestProperties4 = syscall.GUID{0x4E429AD2, 0x1C61, 0x4FEA,
	[8]byte{0xB8, 0xFE, 0x13, 0x5F, 0xB1, 0x9C, 0xDC, 0xE4}}

type ICertificateRequestProperties4Interface interface {
	win32.IInspectableInterface
	Get_SuppressedDefaults() *IVector[string]
	Get_SubjectAlternativeName() *ISubjectAlternativeNameInfo
	Get_Extensions() *IVector[*ICertificateExtension]
}

type ICertificateRequestProperties4Vtbl struct {
	win32.IInspectableVtbl
	Get_SuppressedDefaults     uintptr
	Get_SubjectAlternativeName uintptr
	Get_Extensions             uintptr
}

type ICertificateRequestProperties4 struct {
	win32.IInspectable
}

func (this *ICertificateRequestProperties4) Vtbl() *ICertificateRequestProperties4Vtbl {
	return (*ICertificateRequestProperties4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICertificateRequestProperties4) Get_SuppressedDefaults() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SuppressedDefaults, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICertificateRequestProperties4) Get_SubjectAlternativeName() *ISubjectAlternativeNameInfo {
	var _result *ISubjectAlternativeNameInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SubjectAlternativeName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICertificateRequestProperties4) Get_Extensions() *IVector[*ICertificateExtension] {
	var _result *IVector[*ICertificateExtension]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Extensions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B0BFF720-344E-4331-AF14-A7F7A7EBC93A
var IID_ICertificateStore = syscall.GUID{0xB0BFF720, 0x344E, 0x4331,
	[8]byte{0xAF, 0x14, 0xA7, 0xF7, 0xA7, 0xEB, 0xC9, 0x3A}}

type ICertificateStoreInterface interface {
	win32.IInspectableInterface
	Add(certificate *ICertificate)
	Delete(certificate *ICertificate)
}

type ICertificateStoreVtbl struct {
	win32.IInspectableVtbl
	Add    uintptr
	Delete uintptr
}

type ICertificateStore struct {
	win32.IInspectable
}

func (this *ICertificateStore) Vtbl() *ICertificateStoreVtbl {
	return (*ICertificateStoreVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICertificateStore) Add(certificate *ICertificate) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(certificate)))
	_ = _hr
}

func (this *ICertificateStore) Delete(certificate *ICertificate) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Delete, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(certificate)))
	_ = _hr
}

// C7E68E4A-417D-4D1A-BABD-15687E549974
var IID_ICertificateStore2 = syscall.GUID{0xC7E68E4A, 0x417D, 0x4D1A,
	[8]byte{0xBA, 0xBD, 0x15, 0x68, 0x7E, 0x54, 0x99, 0x74}}

type ICertificateStore2Interface interface {
	win32.IInspectableInterface
	Get_Name() string
}

type ICertificateStore2Vtbl struct {
	win32.IInspectableVtbl
	Get_Name uintptr
}

type ICertificateStore2 struct {
	win32.IInspectable
}

func (this *ICertificateStore2) Vtbl() *ICertificateStore2Vtbl {
	return (*ICertificateStore2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICertificateStore2) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// FBECC739-C6FE-4DE7-99CF-74C3E596E032
var IID_ICertificateStoresStatics = syscall.GUID{0xFBECC739, 0xC6FE, 0x4DE7,
	[8]byte{0x99, 0xCF, 0x74, 0xC3, 0xE5, 0x96, 0xE0, 0x32}}

type ICertificateStoresStaticsInterface interface {
	win32.IInspectableInterface
	FindAllAsync() *IAsyncOperation[*IVectorView[*ICertificate]]
	FindAllWithQueryAsync(query *ICertificateQuery) *IAsyncOperation[*IVectorView[*ICertificate]]
	Get_TrustedRootCertificationAuthorities() *ICertificateStore
	Get_IntermediateCertificationAuthorities() *ICertificateStore
	GetStoreByName(storeName string) *ICertificateStore
}

type ICertificateStoresStaticsVtbl struct {
	win32.IInspectableVtbl
	FindAllAsync                             uintptr
	FindAllWithQueryAsync                    uintptr
	Get_TrustedRootCertificationAuthorities  uintptr
	Get_IntermediateCertificationAuthorities uintptr
	GetStoreByName                           uintptr
}

type ICertificateStoresStatics struct {
	win32.IInspectable
}

func (this *ICertificateStoresStatics) Vtbl() *ICertificateStoresStaticsVtbl {
	return (*ICertificateStoresStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICertificateStoresStatics) FindAllAsync() *IAsyncOperation[*IVectorView[*ICertificate]] {
	var _result *IAsyncOperation[*IVectorView[*ICertificate]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICertificateStoresStatics) FindAllWithQueryAsync(query *ICertificateQuery) *IAsyncOperation[*IVectorView[*ICertificate]] {
	var _result *IAsyncOperation[*IVectorView[*ICertificate]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllWithQueryAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(query)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICertificateStoresStatics) Get_TrustedRootCertificationAuthorities() *ICertificateStore {
	var _result *ICertificateStore
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TrustedRootCertificationAuthorities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICertificateStoresStatics) Get_IntermediateCertificationAuthorities() *ICertificateStore {
	var _result *ICertificateStore
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IntermediateCertificationAuthorities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICertificateStoresStatics) GetStoreByName(storeName string) *ICertificateStore {
	var _result *ICertificateStore
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStoreByName, uintptr(unsafe.Pointer(this)), NewHStr(storeName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FA900B79-A0D4-4B8C-BC55-C0A37EB141ED
var IID_ICertificateStoresStatics2 = syscall.GUID{0xFA900B79, 0xA0D4, 0x4B8C,
	[8]byte{0xBC, 0x55, 0xC0, 0xA3, 0x7E, 0xB1, 0x41, 0xED}}

type ICertificateStoresStatics2Interface interface {
	win32.IInspectableInterface
	GetUserStoreByName(storeName string) *IUserCertificateStore
}

type ICertificateStoresStatics2Vtbl struct {
	win32.IInspectableVtbl
	GetUserStoreByName uintptr
}

type ICertificateStoresStatics2 struct {
	win32.IInspectable
}

func (this *ICertificateStoresStatics2) Vtbl() *ICertificateStoresStatics2Vtbl {
	return (*ICertificateStoresStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICertificateStoresStatics2) GetUserStoreByName(storeName string) *IUserCertificateStore {
	var _result *IUserCertificateStore
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetUserStoreByName, uintptr(unsafe.Pointer(this)), NewHStr(storeName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 422BA922-7C8D-47B7-B59B-B12703733AC3
var IID_IChainBuildingParameters = syscall.GUID{0x422BA922, 0x7C8D, 0x47B7,
	[8]byte{0xB5, 0x9B, 0xB1, 0x27, 0x03, 0x73, 0x3A, 0xC3}}

type IChainBuildingParametersInterface interface {
	win32.IInspectableInterface
	Get_EnhancedKeyUsages() *IVector[string]
	Get_ValidationTimestamp() DateTime
	Put_ValidationTimestamp(value DateTime)
	Get_RevocationCheckEnabled() bool
	Put_RevocationCheckEnabled(value bool)
	Get_NetworkRetrievalEnabled() bool
	Put_NetworkRetrievalEnabled(value bool)
	Get_AuthorityInformationAccessEnabled() bool
	Put_AuthorityInformationAccessEnabled(value bool)
	Get_CurrentTimeValidationEnabled() bool
	Put_CurrentTimeValidationEnabled(value bool)
	Get_ExclusiveTrustRoots() *IVector[*ICertificate]
}

type IChainBuildingParametersVtbl struct {
	win32.IInspectableVtbl
	Get_EnhancedKeyUsages                 uintptr
	Get_ValidationTimestamp               uintptr
	Put_ValidationTimestamp               uintptr
	Get_RevocationCheckEnabled            uintptr
	Put_RevocationCheckEnabled            uintptr
	Get_NetworkRetrievalEnabled           uintptr
	Put_NetworkRetrievalEnabled           uintptr
	Get_AuthorityInformationAccessEnabled uintptr
	Put_AuthorityInformationAccessEnabled uintptr
	Get_CurrentTimeValidationEnabled      uintptr
	Put_CurrentTimeValidationEnabled      uintptr
	Get_ExclusiveTrustRoots               uintptr
}

type IChainBuildingParameters struct {
	win32.IInspectable
}

func (this *IChainBuildingParameters) Vtbl() *IChainBuildingParametersVtbl {
	return (*IChainBuildingParametersVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IChainBuildingParameters) Get_EnhancedKeyUsages() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EnhancedKeyUsages, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IChainBuildingParameters) Get_ValidationTimestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ValidationTimestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IChainBuildingParameters) Put_ValidationTimestamp(value DateTime) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ValidationTimestamp, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IChainBuildingParameters) Get_RevocationCheckEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RevocationCheckEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IChainBuildingParameters) Put_RevocationCheckEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RevocationCheckEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IChainBuildingParameters) Get_NetworkRetrievalEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NetworkRetrievalEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IChainBuildingParameters) Put_NetworkRetrievalEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_NetworkRetrievalEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IChainBuildingParameters) Get_AuthorityInformationAccessEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AuthorityInformationAccessEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IChainBuildingParameters) Put_AuthorityInformationAccessEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AuthorityInformationAccessEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IChainBuildingParameters) Get_CurrentTimeValidationEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentTimeValidationEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IChainBuildingParameters) Put_CurrentTimeValidationEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CurrentTimeValidationEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IChainBuildingParameters) Get_ExclusiveTrustRoots() *IVector[*ICertificate] {
	var _result *IVector[*ICertificate]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExclusiveTrustRoots, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C4743B4A-7EB0-4B56-A040-B9C8E655DDF3
var IID_IChainValidationParameters = syscall.GUID{0xC4743B4A, 0x7EB0, 0x4B56,
	[8]byte{0xA0, 0x40, 0xB9, 0xC8, 0xE6, 0x55, 0xDD, 0xF3}}

type IChainValidationParametersInterface interface {
	win32.IInspectableInterface
	Get_CertificateChainPolicy() CertificateChainPolicy
	Put_CertificateChainPolicy(value CertificateChainPolicy)
	Get_ServerDnsName() *IHostName
	Put_ServerDnsName(value *IHostName)
}

type IChainValidationParametersVtbl struct {
	win32.IInspectableVtbl
	Get_CertificateChainPolicy uintptr
	Put_CertificateChainPolicy uintptr
	Get_ServerDnsName          uintptr
	Put_ServerDnsName          uintptr
}

type IChainValidationParameters struct {
	win32.IInspectable
}

func (this *IChainValidationParameters) Vtbl() *IChainValidationParametersVtbl {
	return (*IChainValidationParametersVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IChainValidationParameters) Get_CertificateChainPolicy() CertificateChainPolicy {
	var _result CertificateChainPolicy
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CertificateChainPolicy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IChainValidationParameters) Put_CertificateChainPolicy(value CertificateChainPolicy) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CertificateChainPolicy, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IChainValidationParameters) Get_ServerDnsName() *IHostName {
	var _result *IHostName
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ServerDnsName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IChainValidationParameters) Put_ServerDnsName(value *IHostName) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ServerDnsName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 61899D9D-3757-4ECB-BDDC-0CA357D7A936
var IID_ICmsAttachedSignature = syscall.GUID{0x61899D9D, 0x3757, 0x4ECB,
	[8]byte{0xBD, 0xDC, 0x0C, 0xA3, 0x57, 0xD7, 0xA9, 0x36}}

type ICmsAttachedSignatureInterface interface {
	win32.IInspectableInterface
	Get_Certificates() *IVectorView[*ICertificate]
	Get_Content() []byte
	Get_Signers() *IVectorView[*ICmsSignerInfo]
	VerifySignature() SignatureValidationResult
}

type ICmsAttachedSignatureVtbl struct {
	win32.IInspectableVtbl
	Get_Certificates uintptr
	Get_Content      uintptr
	Get_Signers      uintptr
	VerifySignature  uintptr
}

type ICmsAttachedSignature struct {
	win32.IInspectable
}

func (this *ICmsAttachedSignature) Vtbl() *ICmsAttachedSignatureVtbl {
	return (*ICmsAttachedSignatureVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICmsAttachedSignature) Get_Certificates() *IVectorView[*ICertificate] {
	var _result *IVectorView[*ICertificate]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Certificates, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICmsAttachedSignature) Get_Content() []byte {
	var _result []byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Content, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICmsAttachedSignature) Get_Signers() *IVectorView[*ICmsSignerInfo] {
	var _result *IVectorView[*ICmsSignerInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Signers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICmsAttachedSignature) VerifySignature() SignatureValidationResult {
	var _result SignatureValidationResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().VerifySignature, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D0C8FC15-F757-4C64-A362-52CC1C77CFFB
var IID_ICmsAttachedSignatureFactory = syscall.GUID{0xD0C8FC15, 0xF757, 0x4C64,
	[8]byte{0xA3, 0x62, 0x52, 0xCC, 0x1C, 0x77, 0xCF, 0xFB}}

type ICmsAttachedSignatureFactoryInterface interface {
	win32.IInspectableInterface
	CreateCmsAttachedSignature(inputBlob *IBuffer) *ICmsAttachedSignature
}

type ICmsAttachedSignatureFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateCmsAttachedSignature uintptr
}

type ICmsAttachedSignatureFactory struct {
	win32.IInspectable
}

func (this *ICmsAttachedSignatureFactory) Vtbl() *ICmsAttachedSignatureFactoryVtbl {
	return (*ICmsAttachedSignatureFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICmsAttachedSignatureFactory) CreateCmsAttachedSignature(inputBlob *IBuffer) *ICmsAttachedSignature {
	var _result *ICmsAttachedSignature
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateCmsAttachedSignature, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(inputBlob)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 87989C8E-B0AD-498D-A7F5-78B59BCE4B36
var IID_ICmsAttachedSignatureStatics = syscall.GUID{0x87989C8E, 0xB0AD, 0x498D,
	[8]byte{0xA7, 0xF5, 0x78, 0xB5, 0x9B, 0xCE, 0x4B, 0x36}}

type ICmsAttachedSignatureStaticsInterface interface {
	win32.IInspectableInterface
	GenerateSignatureAsync(data *IBuffer, signers *IIterable[*ICmsSignerInfo], certificates *IIterable[*ICertificate]) *IAsyncOperation[*IBuffer]
}

type ICmsAttachedSignatureStaticsVtbl struct {
	win32.IInspectableVtbl
	GenerateSignatureAsync uintptr
}

type ICmsAttachedSignatureStatics struct {
	win32.IInspectable
}

func (this *ICmsAttachedSignatureStatics) Vtbl() *ICmsAttachedSignatureStaticsVtbl {
	return (*ICmsAttachedSignatureStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICmsAttachedSignatureStatics) GenerateSignatureAsync(data *IBuffer, signers *IIterable[*ICmsSignerInfo], certificates *IIterable[*ICertificate]) *IAsyncOperation[*IBuffer] {
	var _result *IAsyncOperation[*IBuffer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GenerateSignatureAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(signers)), uintptr(unsafe.Pointer(certificates)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0F1EF154-F65E-4536-8339-5944081DB2CA
var IID_ICmsDetachedSignature = syscall.GUID{0x0F1EF154, 0xF65E, 0x4536,
	[8]byte{0x83, 0x39, 0x59, 0x44, 0x08, 0x1D, 0xB2, 0xCA}}

type ICmsDetachedSignatureInterface interface {
	win32.IInspectableInterface
	Get_Certificates() *IVectorView[*ICertificate]
	Get_Signers() *IVectorView[*ICmsSignerInfo]
	VerifySignatureAsync(data *IInputStream) *IAsyncOperation[SignatureValidationResult]
}

type ICmsDetachedSignatureVtbl struct {
	win32.IInspectableVtbl
	Get_Certificates     uintptr
	Get_Signers          uintptr
	VerifySignatureAsync uintptr
}

type ICmsDetachedSignature struct {
	win32.IInspectable
}

func (this *ICmsDetachedSignature) Vtbl() *ICmsDetachedSignatureVtbl {
	return (*ICmsDetachedSignatureVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICmsDetachedSignature) Get_Certificates() *IVectorView[*ICertificate] {
	var _result *IVectorView[*ICertificate]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Certificates, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICmsDetachedSignature) Get_Signers() *IVectorView[*ICmsSignerInfo] {
	var _result *IVectorView[*ICmsSignerInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Signers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICmsDetachedSignature) VerifySignatureAsync(data *IInputStream) *IAsyncOperation[SignatureValidationResult] {
	var _result *IAsyncOperation[SignatureValidationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().VerifySignatureAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C4AB3503-AE7F-4387-AD19-00F150E48EBB
var IID_ICmsDetachedSignatureFactory = syscall.GUID{0xC4AB3503, 0xAE7F, 0x4387,
	[8]byte{0xAD, 0x19, 0x00, 0xF1, 0x50, 0xE4, 0x8E, 0xBB}}

type ICmsDetachedSignatureFactoryInterface interface {
	win32.IInspectableInterface
	CreateCmsDetachedSignature(inputBlob *IBuffer) *ICmsDetachedSignature
}

type ICmsDetachedSignatureFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateCmsDetachedSignature uintptr
}

type ICmsDetachedSignatureFactory struct {
	win32.IInspectable
}

func (this *ICmsDetachedSignatureFactory) Vtbl() *ICmsDetachedSignatureFactoryVtbl {
	return (*ICmsDetachedSignatureFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICmsDetachedSignatureFactory) CreateCmsDetachedSignature(inputBlob *IBuffer) *ICmsDetachedSignature {
	var _result *ICmsDetachedSignature
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateCmsDetachedSignature, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(inputBlob)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3D114CFD-BF9B-4682-9BE6-91F57C053808
var IID_ICmsDetachedSignatureStatics = syscall.GUID{0x3D114CFD, 0xBF9B, 0x4682,
	[8]byte{0x9B, 0xE6, 0x91, 0xF5, 0x7C, 0x05, 0x38, 0x08}}

type ICmsDetachedSignatureStaticsInterface interface {
	win32.IInspectableInterface
	GenerateSignatureAsync(data *IInputStream, signers *IIterable[*ICmsSignerInfo], certificates *IIterable[*ICertificate]) *IAsyncOperation[*IBuffer]
}

type ICmsDetachedSignatureStaticsVtbl struct {
	win32.IInspectableVtbl
	GenerateSignatureAsync uintptr
}

type ICmsDetachedSignatureStatics struct {
	win32.IInspectable
}

func (this *ICmsDetachedSignatureStatics) Vtbl() *ICmsDetachedSignatureStaticsVtbl {
	return (*ICmsDetachedSignatureStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICmsDetachedSignatureStatics) GenerateSignatureAsync(data *IInputStream, signers *IIterable[*ICmsSignerInfo], certificates *IIterable[*ICertificate]) *IAsyncOperation[*IBuffer] {
	var _result *IAsyncOperation[*IBuffer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GenerateSignatureAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(signers)), uintptr(unsafe.Pointer(certificates)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 50D020DB-1D2F-4C1A-B5C5-D0188FF91F47
var IID_ICmsSignerInfo = syscall.GUID{0x50D020DB, 0x1D2F, 0x4C1A,
	[8]byte{0xB5, 0xC5, 0xD0, 0x18, 0x8F, 0xF9, 0x1F, 0x47}}

type ICmsSignerInfoInterface interface {
	win32.IInspectableInterface
	Get_Certificate() *ICertificate
	Put_Certificate(value *ICertificate)
	Get_HashAlgorithmName() string
	Put_HashAlgorithmName(value string)
	Get_TimestampInfo() *ICmsTimestampInfo
}

type ICmsSignerInfoVtbl struct {
	win32.IInspectableVtbl
	Get_Certificate       uintptr
	Put_Certificate       uintptr
	Get_HashAlgorithmName uintptr
	Put_HashAlgorithmName uintptr
	Get_TimestampInfo     uintptr
}

type ICmsSignerInfo struct {
	win32.IInspectable
}

func (this *ICmsSignerInfo) Vtbl() *ICmsSignerInfoVtbl {
	return (*ICmsSignerInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICmsSignerInfo) Get_Certificate() *ICertificate {
	var _result *ICertificate
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Certificate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICmsSignerInfo) Put_Certificate(value *ICertificate) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Certificate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ICmsSignerInfo) Get_HashAlgorithmName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HashAlgorithmName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICmsSignerInfo) Put_HashAlgorithmName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_HashAlgorithmName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICmsSignerInfo) Get_TimestampInfo() *ICmsTimestampInfo {
	var _result *ICmsTimestampInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TimestampInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2F5F00F2-2C18-4F88-8435-C534086076F5
var IID_ICmsTimestampInfo = syscall.GUID{0x2F5F00F2, 0x2C18, 0x4F88,
	[8]byte{0x84, 0x35, 0xC5, 0x34, 0x08, 0x60, 0x76, 0xF5}}

type ICmsTimestampInfoInterface interface {
	win32.IInspectableInterface
	Get_SigningCertificate() *ICertificate
	Get_Certificates() *IVectorView[*ICertificate]
	Get_Timestamp() DateTime
}

type ICmsTimestampInfoVtbl struct {
	win32.IInspectableVtbl
	Get_SigningCertificate uintptr
	Get_Certificates       uintptr
	Get_Timestamp          uintptr
}

type ICmsTimestampInfo struct {
	win32.IInspectable
}

func (this *ICmsTimestampInfo) Vtbl() *ICmsTimestampInfoVtbl {
	return (*ICmsTimestampInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICmsTimestampInfo) Get_SigningCertificate() *ICertificate {
	var _result *ICertificate
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SigningCertificate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICmsTimestampInfo) Get_Certificates() *IVectorView[*ICertificate] {
	var _result *IVectorView[*ICertificate]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Certificates, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICmsTimestampInfo) Get_Timestamp() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 479065D7-7AC7-4581-8C3B-D07027140448
var IID_IKeyAlgorithmNamesStatics = syscall.GUID{0x479065D7, 0x7AC7, 0x4581,
	[8]byte{0x8C, 0x3B, 0xD0, 0x70, 0x27, 0x14, 0x04, 0x48}}

type IKeyAlgorithmNamesStaticsInterface interface {
	win32.IInspectableInterface
	Get_Rsa() string
	Get_Dsa() string
	Get_Ecdh256() string
	Get_Ecdh384() string
	Get_Ecdh521() string
	Get_Ecdsa256() string
	Get_Ecdsa384() string
	Get_Ecdsa521() string
}

type IKeyAlgorithmNamesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Rsa      uintptr
	Get_Dsa      uintptr
	Get_Ecdh256  uintptr
	Get_Ecdh384  uintptr
	Get_Ecdh521  uintptr
	Get_Ecdsa256 uintptr
	Get_Ecdsa384 uintptr
	Get_Ecdsa521 uintptr
}

type IKeyAlgorithmNamesStatics struct {
	win32.IInspectable
}

func (this *IKeyAlgorithmNamesStatics) Vtbl() *IKeyAlgorithmNamesStaticsVtbl {
	return (*IKeyAlgorithmNamesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKeyAlgorithmNamesStatics) Get_Rsa() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Rsa, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyAlgorithmNamesStatics) Get_Dsa() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Dsa, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyAlgorithmNamesStatics) Get_Ecdh256() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ecdh256, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyAlgorithmNamesStatics) Get_Ecdh384() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ecdh384, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyAlgorithmNamesStatics) Get_Ecdh521() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ecdh521, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyAlgorithmNamesStatics) Get_Ecdsa256() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ecdsa256, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyAlgorithmNamesStatics) Get_Ecdsa384() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ecdsa384, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyAlgorithmNamesStatics) Get_Ecdsa521() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ecdsa521, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// C99B5686-E1FD-4A4A-893D-A26F33DD8BB4
var IID_IKeyAlgorithmNamesStatics2 = syscall.GUID{0xC99B5686, 0xE1FD, 0x4A4A,
	[8]byte{0x89, 0x3D, 0xA2, 0x6F, 0x33, 0xDD, 0x8B, 0xB4}}

type IKeyAlgorithmNamesStatics2Interface interface {
	win32.IInspectableInterface
	Get_Ecdsa() string
	Get_Ecdh() string
}

type IKeyAlgorithmNamesStatics2Vtbl struct {
	win32.IInspectableVtbl
	Get_Ecdsa uintptr
	Get_Ecdh  uintptr
}

type IKeyAlgorithmNamesStatics2 struct {
	win32.IInspectable
}

func (this *IKeyAlgorithmNamesStatics2) Vtbl() *IKeyAlgorithmNamesStatics2Vtbl {
	return (*IKeyAlgorithmNamesStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKeyAlgorithmNamesStatics2) Get_Ecdsa() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ecdsa, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyAlgorithmNamesStatics2) Get_Ecdh() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ecdh, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 1648E246-F644-4326-88BE-3AF102D30E0C
var IID_IKeyAttestationHelperStatics = syscall.GUID{0x1648E246, 0xF644, 0x4326,
	[8]byte{0x88, 0xBE, 0x3A, 0xF1, 0x02, 0xD3, 0x0E, 0x0C}}

type IKeyAttestationHelperStaticsInterface interface {
	win32.IInspectableInterface
	DecryptTpmAttestationCredentialAsync(credential string) *IAsyncOperation[string]
	GetTpmAttestationCredentialId(credential string) string
}

type IKeyAttestationHelperStaticsVtbl struct {
	win32.IInspectableVtbl
	DecryptTpmAttestationCredentialAsync uintptr
	GetTpmAttestationCredentialId        uintptr
}

type IKeyAttestationHelperStatics struct {
	win32.IInspectable
}

func (this *IKeyAttestationHelperStatics) Vtbl() *IKeyAttestationHelperStaticsVtbl {
	return (*IKeyAttestationHelperStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKeyAttestationHelperStatics) DecryptTpmAttestationCredentialAsync(credential string) *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DecryptTpmAttestationCredentialAsync, uintptr(unsafe.Pointer(this)), NewHStr(credential).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKeyAttestationHelperStatics) GetTpmAttestationCredentialId(credential string) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetTpmAttestationCredentialId, uintptr(unsafe.Pointer(this)), NewHStr(credential).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 9C590B2C-A6C6-4A5E-9E64-E85D5279DF97
var IID_IKeyAttestationHelperStatics2 = syscall.GUID{0x9C590B2C, 0xA6C6, 0x4A5E,
	[8]byte{0x9E, 0x64, 0xE8, 0x5D, 0x52, 0x79, 0xDF, 0x97}}

type IKeyAttestationHelperStatics2Interface interface {
	win32.IInspectableInterface
	DecryptTpmAttestationCredentialWithContainerNameAsync(credential string, containerName string) *IAsyncOperation[string]
}

type IKeyAttestationHelperStatics2Vtbl struct {
	win32.IInspectableVtbl
	DecryptTpmAttestationCredentialWithContainerNameAsync uintptr
}

type IKeyAttestationHelperStatics2 struct {
	win32.IInspectable
}

func (this *IKeyAttestationHelperStatics2) Vtbl() *IKeyAttestationHelperStatics2Vtbl {
	return (*IKeyAttestationHelperStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKeyAttestationHelperStatics2) DecryptTpmAttestationCredentialWithContainerNameAsync(credential string, containerName string) *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DecryptTpmAttestationCredentialWithContainerNameAsync, uintptr(unsafe.Pointer(this)), NewHStr(credential).Ptr, NewHStr(containerName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AF186AE0-5529-4602-BD94-0AAB91957B5C
var IID_IKeyStorageProviderNamesStatics = syscall.GUID{0xAF186AE0, 0x5529, 0x4602,
	[8]byte{0xBD, 0x94, 0x0A, 0xAB, 0x91, 0x95, 0x7B, 0x5C}}

type IKeyStorageProviderNamesStaticsInterface interface {
	win32.IInspectableInterface
	Get_SoftwareKeyStorageProvider() string
	Get_SmartcardKeyStorageProvider() string
	Get_PlatformKeyStorageProvider() string
}

type IKeyStorageProviderNamesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_SoftwareKeyStorageProvider  uintptr
	Get_SmartcardKeyStorageProvider uintptr
	Get_PlatformKeyStorageProvider  uintptr
}

type IKeyStorageProviderNamesStatics struct {
	win32.IInspectable
}

func (this *IKeyStorageProviderNamesStatics) Vtbl() *IKeyStorageProviderNamesStaticsVtbl {
	return (*IKeyStorageProviderNamesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKeyStorageProviderNamesStatics) Get_SoftwareKeyStorageProvider() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SoftwareKeyStorageProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyStorageProviderNamesStatics) Get_SmartcardKeyStorageProvider() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SmartcardKeyStorageProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyStorageProviderNamesStatics) Get_PlatformKeyStorageProvider() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlatformKeyStorageProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 262D743D-9C2E-41CC-8812-C4D971DD7C60
var IID_IKeyStorageProviderNamesStatics2 = syscall.GUID{0x262D743D, 0x9C2E, 0x41CC,
	[8]byte{0x88, 0x12, 0xC4, 0xD9, 0x71, 0xDD, 0x7C, 0x60}}

type IKeyStorageProviderNamesStatics2Interface interface {
	win32.IInspectableInterface
	Get_PassportKeyStorageProvider() string
}

type IKeyStorageProviderNamesStatics2Vtbl struct {
	win32.IInspectableVtbl
	Get_PassportKeyStorageProvider uintptr
}

type IKeyStorageProviderNamesStatics2 struct {
	win32.IInspectable
}

func (this *IKeyStorageProviderNamesStatics2) Vtbl() *IKeyStorageProviderNamesStatics2Vtbl {
	return (*IKeyStorageProviderNamesStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKeyStorageProviderNamesStatics2) Get_PassportKeyStorageProvider() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PassportKeyStorageProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 680D3511-9A08-47C8-864A-2EDD4D8EB46C
var IID_IPfxImportParameters = syscall.GUID{0x680D3511, 0x9A08, 0x47C8,
	[8]byte{0x86, 0x4A, 0x2E, 0xDD, 0x4D, 0x8E, 0xB4, 0x6C}}

type IPfxImportParametersInterface interface {
	win32.IInspectableInterface
	Get_Exportable() ExportOption
	Put_Exportable(value ExportOption)
	Get_KeyProtectionLevel() KeyProtectionLevel
	Put_KeyProtectionLevel(value KeyProtectionLevel)
	Get_InstallOptions() InstallOptions
	Put_InstallOptions(value InstallOptions)
	Get_FriendlyName() string
	Put_FriendlyName(value string)
	Get_KeyStorageProviderName() string
	Put_KeyStorageProviderName(value string)
	Get_ContainerNamePrefix() string
	Put_ContainerNamePrefix(value string)
	Get_ReaderName() string
	Put_ReaderName(value string)
}

type IPfxImportParametersVtbl struct {
	win32.IInspectableVtbl
	Get_Exportable             uintptr
	Put_Exportable             uintptr
	Get_KeyProtectionLevel     uintptr
	Put_KeyProtectionLevel     uintptr
	Get_InstallOptions         uintptr
	Put_InstallOptions         uintptr
	Get_FriendlyName           uintptr
	Put_FriendlyName           uintptr
	Get_KeyStorageProviderName uintptr
	Put_KeyStorageProviderName uintptr
	Get_ContainerNamePrefix    uintptr
	Put_ContainerNamePrefix    uintptr
	Get_ReaderName             uintptr
	Put_ReaderName             uintptr
}

type IPfxImportParameters struct {
	win32.IInspectable
}

func (this *IPfxImportParameters) Vtbl() *IPfxImportParametersVtbl {
	return (*IPfxImportParametersVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPfxImportParameters) Get_Exportable() ExportOption {
	var _result ExportOption
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Exportable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPfxImportParameters) Put_Exportable(value ExportOption) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Exportable, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPfxImportParameters) Get_KeyProtectionLevel() KeyProtectionLevel {
	var _result KeyProtectionLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KeyProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPfxImportParameters) Put_KeyProtectionLevel(value KeyProtectionLevel) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_KeyProtectionLevel, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPfxImportParameters) Get_InstallOptions() InstallOptions {
	var _result InstallOptions
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InstallOptions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPfxImportParameters) Put_InstallOptions(value InstallOptions) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InstallOptions, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPfxImportParameters) Get_FriendlyName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FriendlyName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPfxImportParameters) Put_FriendlyName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FriendlyName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPfxImportParameters) Get_KeyStorageProviderName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KeyStorageProviderName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPfxImportParameters) Put_KeyStorageProviderName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_KeyStorageProviderName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPfxImportParameters) Get_ContainerNamePrefix() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContainerNamePrefix, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPfxImportParameters) Put_ContainerNamePrefix(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ContainerNamePrefix, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IPfxImportParameters) Get_ReaderName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReaderName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPfxImportParameters) Put_ReaderName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReaderName, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// 0C154ADB-A496-41F8-8FE5-9E96F36EFBF8
var IID_IStandardCertificateStoreNamesStatics = syscall.GUID{0x0C154ADB, 0xA496, 0x41F8,
	[8]byte{0x8F, 0xE5, 0x9E, 0x96, 0xF3, 0x6E, 0xFB, 0xF8}}

type IStandardCertificateStoreNamesStaticsInterface interface {
	win32.IInspectableInterface
	Get_Personal() string
	Get_TrustedRootCertificationAuthorities() string
	Get_IntermediateCertificationAuthorities() string
}

type IStandardCertificateStoreNamesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Personal                             uintptr
	Get_TrustedRootCertificationAuthorities  uintptr
	Get_IntermediateCertificationAuthorities uintptr
}

type IStandardCertificateStoreNamesStatics struct {
	win32.IInspectable
}

func (this *IStandardCertificateStoreNamesStatics) Vtbl() *IStandardCertificateStoreNamesStaticsVtbl {
	return (*IStandardCertificateStoreNamesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStandardCertificateStoreNamesStatics) Get_Personal() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Personal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStandardCertificateStoreNamesStatics) Get_TrustedRootCertificationAuthorities() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TrustedRootCertificationAuthorities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IStandardCertificateStoreNamesStatics) Get_IntermediateCertificationAuthorities() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IntermediateCertificationAuthorities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 582859F1-569D-4C20-BE7B-4E1C9A0BC52B
var IID_ISubjectAlternativeNameInfo = syscall.GUID{0x582859F1, 0x569D, 0x4C20,
	[8]byte{0xBE, 0x7B, 0x4E, 0x1C, 0x9A, 0x0B, 0xC5, 0x2B}}

type ISubjectAlternativeNameInfoInterface interface {
	win32.IInspectableInterface
	Get_EmailName() *IVectorView[string]
	Get_IPAddress() *IVectorView[string]
	Get_Url() *IVectorView[string]
	Get_DnsName() *IVectorView[string]
	Get_DistinguishedName() *IVectorView[string]
	Get_PrincipalName() *IVectorView[string]
}

type ISubjectAlternativeNameInfoVtbl struct {
	win32.IInspectableVtbl
	Get_EmailName         uintptr
	Get_IPAddress         uintptr
	Get_Url               uintptr
	Get_DnsName           uintptr
	Get_DistinguishedName uintptr
	Get_PrincipalName     uintptr
}

type ISubjectAlternativeNameInfo struct {
	win32.IInspectable
}

func (this *ISubjectAlternativeNameInfo) Vtbl() *ISubjectAlternativeNameInfoVtbl {
	return (*ISubjectAlternativeNameInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISubjectAlternativeNameInfo) Get_EmailName() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EmailName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISubjectAlternativeNameInfo) Get_IPAddress() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IPAddress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISubjectAlternativeNameInfo) Get_Url() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Url, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISubjectAlternativeNameInfo) Get_DnsName() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DnsName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISubjectAlternativeNameInfo) Get_DistinguishedName() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DistinguishedName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISubjectAlternativeNameInfo) Get_PrincipalName() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PrincipalName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 437A78C6-1C51-41EA-B34A-3D654398A370
var IID_ISubjectAlternativeNameInfo2 = syscall.GUID{0x437A78C6, 0x1C51, 0x41EA,
	[8]byte{0xB3, 0x4A, 0x3D, 0x65, 0x43, 0x98, 0xA3, 0x70}}

type ISubjectAlternativeNameInfo2Interface interface {
	win32.IInspectableInterface
	Get_EmailNames() *IVector[string]
	Get_IPAddresses() *IVector[string]
	Get_Urls() *IVector[string]
	Get_DnsNames() *IVector[string]
	Get_DistinguishedNames() *IVector[string]
	Get_PrincipalNames() *IVector[string]
	Get_Extension() *ICertificateExtension
}

type ISubjectAlternativeNameInfo2Vtbl struct {
	win32.IInspectableVtbl
	Get_EmailNames         uintptr
	Get_IPAddresses        uintptr
	Get_Urls               uintptr
	Get_DnsNames           uintptr
	Get_DistinguishedNames uintptr
	Get_PrincipalNames     uintptr
	Get_Extension          uintptr
}

type ISubjectAlternativeNameInfo2 struct {
	win32.IInspectable
}

func (this *ISubjectAlternativeNameInfo2) Vtbl() *ISubjectAlternativeNameInfo2Vtbl {
	return (*ISubjectAlternativeNameInfo2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISubjectAlternativeNameInfo2) Get_EmailNames() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EmailNames, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISubjectAlternativeNameInfo2) Get_IPAddresses() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IPAddresses, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISubjectAlternativeNameInfo2) Get_Urls() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Urls, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISubjectAlternativeNameInfo2) Get_DnsNames() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DnsNames, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISubjectAlternativeNameInfo2) Get_DistinguishedNames() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DistinguishedNames, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISubjectAlternativeNameInfo2) Get_PrincipalNames() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PrincipalNames, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISubjectAlternativeNameInfo2) Get_Extension() *ICertificateExtension {
	var _result *ICertificateExtension
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Extension, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 96313718-22E1-4819-B20B-AB46A6ECA06E
var IID_IUserCertificateEnrollmentManager = syscall.GUID{0x96313718, 0x22E1, 0x4819,
	[8]byte{0xB2, 0x0B, 0xAB, 0x46, 0xA6, 0xEC, 0xA0, 0x6E}}

type IUserCertificateEnrollmentManagerInterface interface {
	win32.IInspectableInterface
	CreateRequestAsync(request *ICertificateRequestProperties) *IAsyncOperation[string]
	InstallCertificateAsync(certificate string, installOption InstallOptions) *IAsyncAction
	ImportPfxDataAsync(pfxData string, password string, exportable ExportOption, keyProtectionLevel KeyProtectionLevel, installOption InstallOptions, friendlyName string) *IAsyncAction
	ImportPfxDataToKspAsync(pfxData string, password string, exportable ExportOption, keyProtectionLevel KeyProtectionLevel, installOption InstallOptions, friendlyName string, keyStorageProvider string) *IAsyncAction
}

type IUserCertificateEnrollmentManagerVtbl struct {
	win32.IInspectableVtbl
	CreateRequestAsync      uintptr
	InstallCertificateAsync uintptr
	ImportPfxDataAsync      uintptr
	ImportPfxDataToKspAsync uintptr
}

type IUserCertificateEnrollmentManager struct {
	win32.IInspectable
}

func (this *IUserCertificateEnrollmentManager) Vtbl() *IUserCertificateEnrollmentManagerVtbl {
	return (*IUserCertificateEnrollmentManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUserCertificateEnrollmentManager) CreateRequestAsync(request *ICertificateRequestProperties) *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateRequestAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(request)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserCertificateEnrollmentManager) InstallCertificateAsync(certificate string, installOption InstallOptions) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InstallCertificateAsync, uintptr(unsafe.Pointer(this)), NewHStr(certificate).Ptr, uintptr(installOption), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserCertificateEnrollmentManager) ImportPfxDataAsync(pfxData string, password string, exportable ExportOption, keyProtectionLevel KeyProtectionLevel, installOption InstallOptions, friendlyName string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ImportPfxDataAsync, uintptr(unsafe.Pointer(this)), NewHStr(pfxData).Ptr, NewHStr(password).Ptr, uintptr(exportable), uintptr(keyProtectionLevel), uintptr(installOption), NewHStr(friendlyName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserCertificateEnrollmentManager) ImportPfxDataToKspAsync(pfxData string, password string, exportable ExportOption, keyProtectionLevel KeyProtectionLevel, installOption InstallOptions, friendlyName string, keyStorageProvider string) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ImportPfxDataToKspAsync, uintptr(unsafe.Pointer(this)), NewHStr(pfxData).Ptr, NewHStr(password).Ptr, uintptr(exportable), uintptr(keyProtectionLevel), uintptr(installOption), NewHStr(friendlyName).Ptr, NewHStr(keyStorageProvider).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0DAD9CB1-65DE-492A-B86D-FC5C482C3747
var IID_IUserCertificateEnrollmentManager2 = syscall.GUID{0x0DAD9CB1, 0x65DE, 0x492A,
	[8]byte{0xB8, 0x6D, 0xFC, 0x5C, 0x48, 0x2C, 0x37, 0x47}}

type IUserCertificateEnrollmentManager2Interface interface {
	win32.IInspectableInterface
	ImportPfxDataToKspWithParametersAsync(pfxData string, password string, pfxImportParameters *IPfxImportParameters) *IAsyncAction
}

type IUserCertificateEnrollmentManager2Vtbl struct {
	win32.IInspectableVtbl
	ImportPfxDataToKspWithParametersAsync uintptr
}

type IUserCertificateEnrollmentManager2 struct {
	win32.IInspectable
}

func (this *IUserCertificateEnrollmentManager2) Vtbl() *IUserCertificateEnrollmentManager2Vtbl {
	return (*IUserCertificateEnrollmentManager2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUserCertificateEnrollmentManager2) ImportPfxDataToKspWithParametersAsync(pfxData string, password string, pfxImportParameters *IPfxImportParameters) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ImportPfxDataToKspWithParametersAsync, uintptr(unsafe.Pointer(this)), NewHStr(pfxData).Ptr, NewHStr(password).Ptr, uintptr(unsafe.Pointer(pfxImportParameters)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C9FB1D83-789F-4B4E-9180-045A757AAC6D
var IID_IUserCertificateStore = syscall.GUID{0xC9FB1D83, 0x789F, 0x4B4E,
	[8]byte{0x91, 0x80, 0x04, 0x5A, 0x75, 0x7A, 0xAC, 0x6D}}

type IUserCertificateStoreInterface interface {
	win32.IInspectableInterface
	RequestAddAsync(certificate *ICertificate) *IAsyncOperation[bool]
	RequestDeleteAsync(certificate *ICertificate) *IAsyncOperation[bool]
	Get_Name() string
}

type IUserCertificateStoreVtbl struct {
	win32.IInspectableVtbl
	RequestAddAsync    uintptr
	RequestDeleteAsync uintptr
	Get_Name           uintptr
}

type IUserCertificateStore struct {
	win32.IInspectable
}

func (this *IUserCertificateStore) Vtbl() *IUserCertificateStoreVtbl {
	return (*IUserCertificateStoreVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUserCertificateStore) RequestAddAsync(certificate *ICertificate) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAddAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(certificate)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserCertificateStore) RequestDeleteAsync(certificate *ICertificate) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestDeleteAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(certificate)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IUserCertificateStore) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// classes

type Certificate struct {
	RtClass
	*ICertificate
}

func NewCertificate_CreateCertificate(certBlob *IBuffer) *Certificate {
	hs := NewHStr("Windows.Security.Cryptography.Certificates.Certificate")
	var pFac *ICertificateFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICertificateFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ICertificate
	p = pFac.CreateCertificate(certBlob)
	result := &Certificate{
		RtClass:      RtClass{PInspect: &p.IInspectable},
		ICertificate: p,
	}
	com.AddToScope(result)
	return result
}

type CertificateChain struct {
	RtClass
	*ICertificateChain
}

type CertificateEnrollmentManager struct {
	RtClass
}

func NewICertificateEnrollmentManagerStatics2() *ICertificateEnrollmentManagerStatics2 {
	var p *ICertificateEnrollmentManagerStatics2
	hs := NewHStr("Windows.Security.Cryptography.Certificates.CertificateEnrollmentManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICertificateEnrollmentManagerStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewICertificateEnrollmentManagerStatics3() *ICertificateEnrollmentManagerStatics3 {
	var p *ICertificateEnrollmentManagerStatics3
	hs := NewHStr("Windows.Security.Cryptography.Certificates.CertificateEnrollmentManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICertificateEnrollmentManagerStatics3, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewICertificateEnrollmentManagerStatics() *ICertificateEnrollmentManagerStatics {
	var p *ICertificateEnrollmentManagerStatics
	hs := NewHStr("Windows.Security.Cryptography.Certificates.CertificateEnrollmentManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICertificateEnrollmentManagerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type CertificateExtension struct {
	RtClass
	*ICertificateExtension
}

func NewCertificateExtension() *CertificateExtension {
	hs := NewHStr("Windows.Security.Cryptography.Certificates.CertificateExtension")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &CertificateExtension{
		RtClass:               RtClass{PInspect: p},
		ICertificateExtension: (*ICertificateExtension)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type CertificateKeyUsages struct {
	RtClass
	*ICertificateKeyUsages
}

func NewCertificateKeyUsages() *CertificateKeyUsages {
	hs := NewHStr("Windows.Security.Cryptography.Certificates.CertificateKeyUsages")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &CertificateKeyUsages{
		RtClass:               RtClass{PInspect: p},
		ICertificateKeyUsages: (*ICertificateKeyUsages)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type CertificateQuery struct {
	RtClass
	*ICertificateQuery
}

func NewCertificateQuery() *CertificateQuery {
	hs := NewHStr("Windows.Security.Cryptography.Certificates.CertificateQuery")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &CertificateQuery{
		RtClass:           RtClass{PInspect: p},
		ICertificateQuery: (*ICertificateQuery)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type CertificateRequestProperties struct {
	RtClass
	*ICertificateRequestProperties
}

func NewCertificateRequestProperties() *CertificateRequestProperties {
	hs := NewHStr("Windows.Security.Cryptography.Certificates.CertificateRequestProperties")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &CertificateRequestProperties{
		RtClass:                       RtClass{PInspect: p},
		ICertificateRequestProperties: (*ICertificateRequestProperties)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type CertificateStore struct {
	RtClass
	*ICertificateStore
}

type CertificateStores struct {
	RtClass
}

func NewICertificateStoresStatics2() *ICertificateStoresStatics2 {
	var p *ICertificateStoresStatics2
	hs := NewHStr("Windows.Security.Cryptography.Certificates.CertificateStores")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICertificateStoresStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewICertificateStoresStatics() *ICertificateStoresStatics {
	var p *ICertificateStoresStatics
	hs := NewHStr("Windows.Security.Cryptography.Certificates.CertificateStores")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICertificateStoresStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type ChainBuildingParameters struct {
	RtClass
	*IChainBuildingParameters
}

func NewChainBuildingParameters() *ChainBuildingParameters {
	hs := NewHStr("Windows.Security.Cryptography.Certificates.ChainBuildingParameters")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &ChainBuildingParameters{
		RtClass:                  RtClass{PInspect: p},
		IChainBuildingParameters: (*IChainBuildingParameters)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type ChainValidationParameters struct {
	RtClass
	*IChainValidationParameters
}

func NewChainValidationParameters() *ChainValidationParameters {
	hs := NewHStr("Windows.Security.Cryptography.Certificates.ChainValidationParameters")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &ChainValidationParameters{
		RtClass:                    RtClass{PInspect: p},
		IChainValidationParameters: (*IChainValidationParameters)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type CmsAttachedSignature struct {
	RtClass
	*ICmsAttachedSignature
}

func NewCmsAttachedSignature_CreateCmsAttachedSignature(inputBlob *IBuffer) *CmsAttachedSignature {
	hs := NewHStr("Windows.Security.Cryptography.Certificates.CmsAttachedSignature")
	var pFac *ICmsAttachedSignatureFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICmsAttachedSignatureFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ICmsAttachedSignature
	p = pFac.CreateCmsAttachedSignature(inputBlob)
	result := &CmsAttachedSignature{
		RtClass:               RtClass{PInspect: &p.IInspectable},
		ICmsAttachedSignature: p,
	}
	com.AddToScope(result)
	return result
}

func NewICmsAttachedSignatureStatics() *ICmsAttachedSignatureStatics {
	var p *ICmsAttachedSignatureStatics
	hs := NewHStr("Windows.Security.Cryptography.Certificates.CmsAttachedSignature")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICmsAttachedSignatureStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type CmsDetachedSignature struct {
	RtClass
	*ICmsDetachedSignature
}

func NewCmsDetachedSignature_CreateCmsDetachedSignature(inputBlob *IBuffer) *CmsDetachedSignature {
	hs := NewHStr("Windows.Security.Cryptography.Certificates.CmsDetachedSignature")
	var pFac *ICmsDetachedSignatureFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICmsDetachedSignatureFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ICmsDetachedSignature
	p = pFac.CreateCmsDetachedSignature(inputBlob)
	result := &CmsDetachedSignature{
		RtClass:               RtClass{PInspect: &p.IInspectable},
		ICmsDetachedSignature: p,
	}
	com.AddToScope(result)
	return result
}

func NewICmsDetachedSignatureStatics() *ICmsDetachedSignatureStatics {
	var p *ICmsDetachedSignatureStatics
	hs := NewHStr("Windows.Security.Cryptography.Certificates.CmsDetachedSignature")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICmsDetachedSignatureStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type CmsSignerInfo struct {
	RtClass
	*ICmsSignerInfo
}

func NewCmsSignerInfo() *CmsSignerInfo {
	hs := NewHStr("Windows.Security.Cryptography.Certificates.CmsSignerInfo")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &CmsSignerInfo{
		RtClass:        RtClass{PInspect: p},
		ICmsSignerInfo: (*ICmsSignerInfo)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type CmsTimestampInfo struct {
	RtClass
	*ICmsTimestampInfo
}

type KeyAlgorithmNames struct {
	RtClass
}

func NewIKeyAlgorithmNamesStatics() *IKeyAlgorithmNamesStatics {
	var p *IKeyAlgorithmNamesStatics
	hs := NewHStr("Windows.Security.Cryptography.Certificates.KeyAlgorithmNames")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IKeyAlgorithmNamesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIKeyAlgorithmNamesStatics2() *IKeyAlgorithmNamesStatics2 {
	var p *IKeyAlgorithmNamesStatics2
	hs := NewHStr("Windows.Security.Cryptography.Certificates.KeyAlgorithmNames")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IKeyAlgorithmNamesStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type KeyAttestationHelper struct {
	RtClass
}

func NewIKeyAttestationHelperStatics2() *IKeyAttestationHelperStatics2 {
	var p *IKeyAttestationHelperStatics2
	hs := NewHStr("Windows.Security.Cryptography.Certificates.KeyAttestationHelper")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IKeyAttestationHelperStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIKeyAttestationHelperStatics() *IKeyAttestationHelperStatics {
	var p *IKeyAttestationHelperStatics
	hs := NewHStr("Windows.Security.Cryptography.Certificates.KeyAttestationHelper")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IKeyAttestationHelperStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type KeyStorageProviderNames struct {
	RtClass
}

func NewIKeyStorageProviderNamesStatics2() *IKeyStorageProviderNamesStatics2 {
	var p *IKeyStorageProviderNamesStatics2
	hs := NewHStr("Windows.Security.Cryptography.Certificates.KeyStorageProviderNames")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IKeyStorageProviderNamesStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIKeyStorageProviderNamesStatics() *IKeyStorageProviderNamesStatics {
	var p *IKeyStorageProviderNamesStatics
	hs := NewHStr("Windows.Security.Cryptography.Certificates.KeyStorageProviderNames")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IKeyStorageProviderNamesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type PfxImportParameters struct {
	RtClass
	*IPfxImportParameters
}

func NewPfxImportParameters() *PfxImportParameters {
	hs := NewHStr("Windows.Security.Cryptography.Certificates.PfxImportParameters")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &PfxImportParameters{
		RtClass:              RtClass{PInspect: p},
		IPfxImportParameters: (*IPfxImportParameters)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type StandardCertificateStoreNames struct {
	RtClass
}

func NewIStandardCertificateStoreNamesStatics() *IStandardCertificateStoreNamesStatics {
	var p *IStandardCertificateStoreNamesStatics
	hs := NewHStr("Windows.Security.Cryptography.Certificates.StandardCertificateStoreNames")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IStandardCertificateStoreNamesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type SubjectAlternativeNameInfo struct {
	RtClass
	*ISubjectAlternativeNameInfo
}

func NewSubjectAlternativeNameInfo() *SubjectAlternativeNameInfo {
	hs := NewHStr("Windows.Security.Cryptography.Certificates.SubjectAlternativeNameInfo")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &SubjectAlternativeNameInfo{
		RtClass:                     RtClass{PInspect: p},
		ISubjectAlternativeNameInfo: (*ISubjectAlternativeNameInfo)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type UserCertificateEnrollmentManager struct {
	RtClass
	*IUserCertificateEnrollmentManager
}

type UserCertificateStore struct {
	RtClass
	*IUserCertificateStore
}
