package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type Capi1KdfTargetAlgorithm int32

const (
	Capi1KdfTargetAlgorithm_NotAes Capi1KdfTargetAlgorithm = 0
	Capi1KdfTargetAlgorithm_Aes    Capi1KdfTargetAlgorithm = 1
)

// enum
type CryptographicPadding int32

const (
	CryptographicPadding_None        CryptographicPadding = 0
	CryptographicPadding_RsaOaep     CryptographicPadding = 1
	CryptographicPadding_RsaPkcs1V15 CryptographicPadding = 2
	CryptographicPadding_RsaPss      CryptographicPadding = 3
)

// enum
type CryptographicPrivateKeyBlobType int32

const (
	CryptographicPrivateKeyBlobType_Pkcs8RawPrivateKeyInfo  CryptographicPrivateKeyBlobType = 0
	CryptographicPrivateKeyBlobType_Pkcs1RsaPrivateKey      CryptographicPrivateKeyBlobType = 1
	CryptographicPrivateKeyBlobType_BCryptPrivateKey        CryptographicPrivateKeyBlobType = 2
	CryptographicPrivateKeyBlobType_Capi1PrivateKey         CryptographicPrivateKeyBlobType = 3
	CryptographicPrivateKeyBlobType_BCryptEccFullPrivateKey CryptographicPrivateKeyBlobType = 4
)

// enum
type CryptographicPublicKeyBlobType int32

const (
	CryptographicPublicKeyBlobType_X509SubjectPublicKeyInfo CryptographicPublicKeyBlobType = 0
	CryptographicPublicKeyBlobType_Pkcs1RsaPublicKey        CryptographicPublicKeyBlobType = 1
	CryptographicPublicKeyBlobType_BCryptPublicKey          CryptographicPublicKeyBlobType = 2
	CryptographicPublicKeyBlobType_Capi1PublicKey           CryptographicPublicKeyBlobType = 3
	CryptographicPublicKeyBlobType_BCryptEccFullPublicKey   CryptographicPublicKeyBlobType = 4
)

// interfaces

// CAF6FCE4-67C0-46AA-84F9-752E77449F9B
var IID_IAsymmetricAlgorithmNamesStatics = syscall.GUID{0xCAF6FCE4, 0x67C0, 0x46AA,
	[8]byte{0x84, 0xF9, 0x75, 0x2E, 0x77, 0x44, 0x9F, 0x9B}}

type IAsymmetricAlgorithmNamesStaticsInterface interface {
	win32.IInspectableInterface
	Get_RsaPkcs1() string
	Get_RsaOaepSha1() string
	Get_RsaOaepSha256() string
	Get_RsaOaepSha384() string
	Get_RsaOaepSha512() string
	Get_EcdsaP256Sha256() string
	Get_EcdsaP384Sha384() string
	Get_EcdsaP521Sha512() string
	Get_DsaSha1() string
	Get_DsaSha256() string
	Get_RsaSignPkcs1Sha1() string
	Get_RsaSignPkcs1Sha256() string
	Get_RsaSignPkcs1Sha384() string
	Get_RsaSignPkcs1Sha512() string
	Get_RsaSignPssSha1() string
	Get_RsaSignPssSha256() string
	Get_RsaSignPssSha384() string
	Get_RsaSignPssSha512() string
}

type IAsymmetricAlgorithmNamesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_RsaPkcs1           uintptr
	Get_RsaOaepSha1        uintptr
	Get_RsaOaepSha256      uintptr
	Get_RsaOaepSha384      uintptr
	Get_RsaOaepSha512      uintptr
	Get_EcdsaP256Sha256    uintptr
	Get_EcdsaP384Sha384    uintptr
	Get_EcdsaP521Sha512    uintptr
	Get_DsaSha1            uintptr
	Get_DsaSha256          uintptr
	Get_RsaSignPkcs1Sha1   uintptr
	Get_RsaSignPkcs1Sha256 uintptr
	Get_RsaSignPkcs1Sha384 uintptr
	Get_RsaSignPkcs1Sha512 uintptr
	Get_RsaSignPssSha1     uintptr
	Get_RsaSignPssSha256   uintptr
	Get_RsaSignPssSha384   uintptr
	Get_RsaSignPssSha512   uintptr
}

type IAsymmetricAlgorithmNamesStatics struct {
	win32.IInspectable
}

func (this *IAsymmetricAlgorithmNamesStatics) Vtbl() *IAsymmetricAlgorithmNamesStaticsVtbl {
	return (*IAsymmetricAlgorithmNamesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAsymmetricAlgorithmNamesStatics) Get_RsaPkcs1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RsaPkcs1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAsymmetricAlgorithmNamesStatics) Get_RsaOaepSha1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RsaOaepSha1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAsymmetricAlgorithmNamesStatics) Get_RsaOaepSha256() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RsaOaepSha256, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAsymmetricAlgorithmNamesStatics) Get_RsaOaepSha384() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RsaOaepSha384, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAsymmetricAlgorithmNamesStatics) Get_RsaOaepSha512() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RsaOaepSha512, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAsymmetricAlgorithmNamesStatics) Get_EcdsaP256Sha256() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EcdsaP256Sha256, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAsymmetricAlgorithmNamesStatics) Get_EcdsaP384Sha384() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EcdsaP384Sha384, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAsymmetricAlgorithmNamesStatics) Get_EcdsaP521Sha512() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EcdsaP521Sha512, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAsymmetricAlgorithmNamesStatics) Get_DsaSha1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DsaSha1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAsymmetricAlgorithmNamesStatics) Get_DsaSha256() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DsaSha256, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAsymmetricAlgorithmNamesStatics) Get_RsaSignPkcs1Sha1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RsaSignPkcs1Sha1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAsymmetricAlgorithmNamesStatics) Get_RsaSignPkcs1Sha256() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RsaSignPkcs1Sha256, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAsymmetricAlgorithmNamesStatics) Get_RsaSignPkcs1Sha384() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RsaSignPkcs1Sha384, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAsymmetricAlgorithmNamesStatics) Get_RsaSignPkcs1Sha512() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RsaSignPkcs1Sha512, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAsymmetricAlgorithmNamesStatics) Get_RsaSignPssSha1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RsaSignPssSha1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAsymmetricAlgorithmNamesStatics) Get_RsaSignPssSha256() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RsaSignPssSha256, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAsymmetricAlgorithmNamesStatics) Get_RsaSignPssSha384() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RsaSignPssSha384, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAsymmetricAlgorithmNamesStatics) Get_RsaSignPssSha512() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RsaSignPssSha512, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// F141C0D6-4BFF-4F23-BA66-6045B137D5DF
var IID_IAsymmetricAlgorithmNamesStatics2 = syscall.GUID{0xF141C0D6, 0x4BFF, 0x4F23,
	[8]byte{0xBA, 0x66, 0x60, 0x45, 0xB1, 0x37, 0xD5, 0xDF}}

type IAsymmetricAlgorithmNamesStatics2Interface interface {
	win32.IInspectableInterface
	Get_EcdsaSha256() string
	Get_EcdsaSha384() string
	Get_EcdsaSha512() string
}

type IAsymmetricAlgorithmNamesStatics2Vtbl struct {
	win32.IInspectableVtbl
	Get_EcdsaSha256 uintptr
	Get_EcdsaSha384 uintptr
	Get_EcdsaSha512 uintptr
}

type IAsymmetricAlgorithmNamesStatics2 struct {
	win32.IInspectable
}

func (this *IAsymmetricAlgorithmNamesStatics2) Vtbl() *IAsymmetricAlgorithmNamesStatics2Vtbl {
	return (*IAsymmetricAlgorithmNamesStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAsymmetricAlgorithmNamesStatics2) Get_EcdsaSha256() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EcdsaSha256, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAsymmetricAlgorithmNamesStatics2) Get_EcdsaSha384() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EcdsaSha384, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAsymmetricAlgorithmNamesStatics2) Get_EcdsaSha512() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EcdsaSha512, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// E8D2FF37-6259-4E88-B7E0-94191FDE699E
var IID_IAsymmetricKeyAlgorithmProvider = syscall.GUID{0xE8D2FF37, 0x6259, 0x4E88,
	[8]byte{0xB7, 0xE0, 0x94, 0x19, 0x1F, 0xDE, 0x69, 0x9E}}

type IAsymmetricKeyAlgorithmProviderInterface interface {
	win32.IInspectableInterface
	Get_AlgorithmName() string
	CreateKeyPair(keySize uint32) *ICryptographicKey
	ImportDefaultPrivateKeyBlob(keyBlob *IBuffer) *ICryptographicKey
	ImportKeyPairWithBlobType(keyBlob *IBuffer, BlobType CryptographicPrivateKeyBlobType) *ICryptographicKey
	ImportDefaultPublicKeyBlob(keyBlob *IBuffer) *ICryptographicKey
	ImportPublicKeyWithBlobType(keyBlob *IBuffer, BlobType CryptographicPublicKeyBlobType) *ICryptographicKey
}

type IAsymmetricKeyAlgorithmProviderVtbl struct {
	win32.IInspectableVtbl
	Get_AlgorithmName           uintptr
	CreateKeyPair               uintptr
	ImportDefaultPrivateKeyBlob uintptr
	ImportKeyPairWithBlobType   uintptr
	ImportDefaultPublicKeyBlob  uintptr
	ImportPublicKeyWithBlobType uintptr
}

type IAsymmetricKeyAlgorithmProvider struct {
	win32.IInspectable
}

func (this *IAsymmetricKeyAlgorithmProvider) Vtbl() *IAsymmetricKeyAlgorithmProviderVtbl {
	return (*IAsymmetricKeyAlgorithmProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAsymmetricKeyAlgorithmProvider) Get_AlgorithmName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlgorithmName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAsymmetricKeyAlgorithmProvider) CreateKeyPair(keySize uint32) *ICryptographicKey {
	var _result *ICryptographicKey
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateKeyPair, uintptr(unsafe.Pointer(this)), uintptr(keySize), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAsymmetricKeyAlgorithmProvider) ImportDefaultPrivateKeyBlob(keyBlob *IBuffer) *ICryptographicKey {
	var _result *ICryptographicKey
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ImportDefaultPrivateKeyBlob, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(keyBlob)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAsymmetricKeyAlgorithmProvider) ImportKeyPairWithBlobType(keyBlob *IBuffer, BlobType CryptographicPrivateKeyBlobType) *ICryptographicKey {
	var _result *ICryptographicKey
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ImportKeyPairWithBlobType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(keyBlob)), uintptr(BlobType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAsymmetricKeyAlgorithmProvider) ImportDefaultPublicKeyBlob(keyBlob *IBuffer) *ICryptographicKey {
	var _result *ICryptographicKey
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ImportDefaultPublicKeyBlob, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(keyBlob)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAsymmetricKeyAlgorithmProvider) ImportPublicKeyWithBlobType(keyBlob *IBuffer, BlobType CryptographicPublicKeyBlobType) *ICryptographicKey {
	var _result *ICryptographicKey
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ImportPublicKeyWithBlobType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(keyBlob)), uintptr(BlobType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4E322A7E-7C4D-4997-AC4F-1B848B36306E
var IID_IAsymmetricKeyAlgorithmProvider2 = syscall.GUID{0x4E322A7E, 0x7C4D, 0x4997,
	[8]byte{0xAC, 0x4F, 0x1B, 0x84, 0x8B, 0x36, 0x30, 0x6E}}

type IAsymmetricKeyAlgorithmProvider2Interface interface {
	win32.IInspectableInterface
	CreateKeyPairWithCurveName(curveName string) *ICryptographicKey
	CreateKeyPairWithCurveParameters(parametersLength uint32, parameters *byte) *ICryptographicKey
}

type IAsymmetricKeyAlgorithmProvider2Vtbl struct {
	win32.IInspectableVtbl
	CreateKeyPairWithCurveName       uintptr
	CreateKeyPairWithCurveParameters uintptr
}

type IAsymmetricKeyAlgorithmProvider2 struct {
	win32.IInspectable
}

func (this *IAsymmetricKeyAlgorithmProvider2) Vtbl() *IAsymmetricKeyAlgorithmProvider2Vtbl {
	return (*IAsymmetricKeyAlgorithmProvider2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAsymmetricKeyAlgorithmProvider2) CreateKeyPairWithCurveName(curveName string) *ICryptographicKey {
	var _result *ICryptographicKey
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateKeyPairWithCurveName, uintptr(unsafe.Pointer(this)), NewHStr(curveName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAsymmetricKeyAlgorithmProvider2) CreateKeyPairWithCurveParameters(parametersLength uint32, parameters *byte) *ICryptographicKey {
	var _result *ICryptographicKey
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateKeyPairWithCurveParameters, uintptr(unsafe.Pointer(this)), uintptr(parametersLength), uintptr(unsafe.Pointer(parameters)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 425BDE18-A7F3-47A6-A8D2-C48D6033A65C
var IID_IAsymmetricKeyAlgorithmProviderStatics = syscall.GUID{0x425BDE18, 0xA7F3, 0x47A6,
	[8]byte{0xA8, 0xD2, 0xC4, 0x8D, 0x60, 0x33, 0xA6, 0x5C}}

type IAsymmetricKeyAlgorithmProviderStaticsInterface interface {
	win32.IInspectableInterface
	OpenAlgorithm(algorithm string) *IAsymmetricKeyAlgorithmProvider
}

type IAsymmetricKeyAlgorithmProviderStaticsVtbl struct {
	win32.IInspectableVtbl
	OpenAlgorithm uintptr
}

type IAsymmetricKeyAlgorithmProviderStatics struct {
	win32.IInspectable
}

func (this *IAsymmetricKeyAlgorithmProviderStatics) Vtbl() *IAsymmetricKeyAlgorithmProviderStaticsVtbl {
	return (*IAsymmetricKeyAlgorithmProviderStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAsymmetricKeyAlgorithmProviderStatics) OpenAlgorithm(algorithm string) *IAsymmetricKeyAlgorithmProvider {
	var _result *IAsymmetricKeyAlgorithmProvider
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenAlgorithm, uintptr(unsafe.Pointer(this)), NewHStr(algorithm).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9FEA0639-6FF7-4C85-A095-95EB31715EB9
var IID_ICryptographicEngineStatics = syscall.GUID{0x9FEA0639, 0x6FF7, 0x4C85,
	[8]byte{0xA0, 0x95, 0x95, 0xEB, 0x31, 0x71, 0x5E, 0xB9}}

type ICryptographicEngineStaticsInterface interface {
	win32.IInspectableInterface
	Encrypt(key *ICryptographicKey, data *IBuffer, iv *IBuffer) *IBuffer
	Decrypt(key *ICryptographicKey, data *IBuffer, iv *IBuffer) *IBuffer
	EncryptAndAuthenticate(key *ICryptographicKey, data *IBuffer, nonce *IBuffer, authenticatedData *IBuffer) *IEncryptedAndAuthenticatedData
	DecryptAndAuthenticate(key *ICryptographicKey, data *IBuffer, nonce *IBuffer, authenticationTag *IBuffer, authenticatedData *IBuffer) *IBuffer
	Sign(key *ICryptographicKey, data *IBuffer) *IBuffer
	VerifySignature(key *ICryptographicKey, data *IBuffer, signature *IBuffer) bool
	DeriveKeyMaterial(key *ICryptographicKey, parameters *IKeyDerivationParameters, desiredKeySize uint32) *IBuffer
}

type ICryptographicEngineStaticsVtbl struct {
	win32.IInspectableVtbl
	Encrypt                uintptr
	Decrypt                uintptr
	EncryptAndAuthenticate uintptr
	DecryptAndAuthenticate uintptr
	Sign                   uintptr
	VerifySignature        uintptr
	DeriveKeyMaterial      uintptr
}

type ICryptographicEngineStatics struct {
	win32.IInspectable
}

func (this *ICryptographicEngineStatics) Vtbl() *ICryptographicEngineStaticsVtbl {
	return (*ICryptographicEngineStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICryptographicEngineStatics) Encrypt(key *ICryptographicKey, data *IBuffer, iv *IBuffer) *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Encrypt, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(iv)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICryptographicEngineStatics) Decrypt(key *ICryptographicKey, data *IBuffer, iv *IBuffer) *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Decrypt, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(iv)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICryptographicEngineStatics) EncryptAndAuthenticate(key *ICryptographicKey, data *IBuffer, nonce *IBuffer, authenticatedData *IBuffer) *IEncryptedAndAuthenticatedData {
	var _result *IEncryptedAndAuthenticatedData
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EncryptAndAuthenticate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(nonce)), uintptr(unsafe.Pointer(authenticatedData)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICryptographicEngineStatics) DecryptAndAuthenticate(key *ICryptographicKey, data *IBuffer, nonce *IBuffer, authenticationTag *IBuffer, authenticatedData *IBuffer) *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DecryptAndAuthenticate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(nonce)), uintptr(unsafe.Pointer(authenticationTag)), uintptr(unsafe.Pointer(authenticatedData)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICryptographicEngineStatics) Sign(key *ICryptographicKey, data *IBuffer) *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Sign, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICryptographicEngineStatics) VerifySignature(key *ICryptographicKey, data *IBuffer, signature *IBuffer) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().VerifySignature, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(signature)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICryptographicEngineStatics) DeriveKeyMaterial(key *ICryptographicKey, parameters *IKeyDerivationParameters, desiredKeySize uint32) *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DeriveKeyMaterial, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(parameters)), uintptr(desiredKeySize), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 675948FE-DF9F-4191-92C7-6CE6F58420E0
var IID_ICryptographicEngineStatics2 = syscall.GUID{0x675948FE, 0xDF9F, 0x4191,
	[8]byte{0x92, 0xC7, 0x6C, 0xE6, 0xF5, 0x84, 0x20, 0xE0}}

type ICryptographicEngineStatics2Interface interface {
	win32.IInspectableInterface
	SignHashedData(key *ICryptographicKey, data *IBuffer) *IBuffer
	VerifySignatureWithHashInput(key *ICryptographicKey, data *IBuffer, signature *IBuffer) bool
	DecryptAsync(key *ICryptographicKey, data *IBuffer, iv *IBuffer) *IAsyncOperation[*IBuffer]
	SignAsync(key *ICryptographicKey, data *IBuffer) *IAsyncOperation[*IBuffer]
	SignHashedDataAsync(key *ICryptographicKey, data *IBuffer) *IAsyncOperation[*IBuffer]
}

type ICryptographicEngineStatics2Vtbl struct {
	win32.IInspectableVtbl
	SignHashedData               uintptr
	VerifySignatureWithHashInput uintptr
	DecryptAsync                 uintptr
	SignAsync                    uintptr
	SignHashedDataAsync          uintptr
}

type ICryptographicEngineStatics2 struct {
	win32.IInspectable
}

func (this *ICryptographicEngineStatics2) Vtbl() *ICryptographicEngineStatics2Vtbl {
	return (*ICryptographicEngineStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICryptographicEngineStatics2) SignHashedData(key *ICryptographicKey, data *IBuffer) *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SignHashedData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICryptographicEngineStatics2) VerifySignatureWithHashInput(key *ICryptographicKey, data *IBuffer, signature *IBuffer) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().VerifySignatureWithHashInput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(signature)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICryptographicEngineStatics2) DecryptAsync(key *ICryptographicKey, data *IBuffer, iv *IBuffer) *IAsyncOperation[*IBuffer] {
	var _result *IAsyncOperation[*IBuffer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DecryptAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(iv)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICryptographicEngineStatics2) SignAsync(key *ICryptographicKey, data *IBuffer) *IAsyncOperation[*IBuffer] {
	var _result *IAsyncOperation[*IBuffer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SignAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICryptographicEngineStatics2) SignHashedDataAsync(key *ICryptographicKey, data *IBuffer) *IAsyncOperation[*IBuffer] {
	var _result *IAsyncOperation[*IBuffer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SignHashedDataAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// ED2A3B70-8E7B-4009-8401-FFD1A62EEB27
var IID_ICryptographicKey = syscall.GUID{0xED2A3B70, 0x8E7B, 0x4009,
	[8]byte{0x84, 0x01, 0xFF, 0xD1, 0xA6, 0x2E, 0xEB, 0x27}}

type ICryptographicKeyInterface interface {
	win32.IInspectableInterface
	Get_KeySize() uint32
	ExportDefaultPrivateKeyBlobType() *IBuffer
	ExportPrivateKeyWithBlobType(BlobType CryptographicPrivateKeyBlobType) *IBuffer
	ExportDefaultPublicKeyBlobType() *IBuffer
	ExportPublicKeyWithBlobType(BlobType CryptographicPublicKeyBlobType) *IBuffer
}

type ICryptographicKeyVtbl struct {
	win32.IInspectableVtbl
	Get_KeySize                     uintptr
	ExportDefaultPrivateKeyBlobType uintptr
	ExportPrivateKeyWithBlobType    uintptr
	ExportDefaultPublicKeyBlobType  uintptr
	ExportPublicKeyWithBlobType     uintptr
}

type ICryptographicKey struct {
	win32.IInspectable
}

func (this *ICryptographicKey) Vtbl() *ICryptographicKeyVtbl {
	return (*ICryptographicKeyVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICryptographicKey) Get_KeySize() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KeySize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICryptographicKey) ExportDefaultPrivateKeyBlobType() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ExportDefaultPrivateKeyBlobType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICryptographicKey) ExportPrivateKeyWithBlobType(BlobType CryptographicPrivateKeyBlobType) *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ExportPrivateKeyWithBlobType, uintptr(unsafe.Pointer(this)), uintptr(BlobType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICryptographicKey) ExportDefaultPublicKeyBlobType() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ExportDefaultPublicKeyBlobType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICryptographicKey) ExportPublicKeyWithBlobType(BlobType CryptographicPublicKeyBlobType) *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ExportPublicKeyWithBlobType, uintptr(unsafe.Pointer(this)), uintptr(BlobType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B3FF930C-AEEB-409E-B7D4-9B95295AAECF
var IID_IEccCurveNamesStatics = syscall.GUID{0xB3FF930C, 0xAEEB, 0x409E,
	[8]byte{0xB7, 0xD4, 0x9B, 0x95, 0x29, 0x5A, 0xAE, 0xCF}}

type IEccCurveNamesStaticsInterface interface {
	win32.IInspectableInterface
	Get_BrainpoolP160r1() string
	Get_BrainpoolP160t1() string
	Get_BrainpoolP192r1() string
	Get_BrainpoolP192t1() string
	Get_BrainpoolP224r1() string
	Get_BrainpoolP224t1() string
	Get_BrainpoolP256r1() string
	Get_BrainpoolP256t1() string
	Get_BrainpoolP320r1() string
	Get_BrainpoolP320t1() string
	Get_BrainpoolP384r1() string
	Get_BrainpoolP384t1() string
	Get_BrainpoolP512r1() string
	Get_BrainpoolP512t1() string
	Get_Curve25519() string
	Get_Ec192wapi() string
	Get_NistP192() string
	Get_NistP224() string
	Get_NistP256() string
	Get_NistP384() string
	Get_NistP521() string
	Get_NumsP256t1() string
	Get_NumsP384t1() string
	Get_NumsP512t1() string
	Get_SecP160k1() string
	Get_SecP160r1() string
	Get_SecP160r2() string
	Get_SecP192k1() string
	Get_SecP192r1() string
	Get_SecP224k1() string
	Get_SecP224r1() string
	Get_SecP256k1() string
	Get_SecP256r1() string
	Get_SecP384r1() string
	Get_SecP521r1() string
	Get_Wtls7() string
	Get_Wtls9() string
	Get_Wtls12() string
	Get_X962P192v1() string
	Get_X962P192v2() string
	Get_X962P192v3() string
	Get_X962P239v1() string
	Get_X962P239v2() string
	Get_X962P239v3() string
	Get_X962P256v1() string
	Get_AllEccCurveNames() *IVectorView[string]
}

type IEccCurveNamesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_BrainpoolP160r1  uintptr
	Get_BrainpoolP160t1  uintptr
	Get_BrainpoolP192r1  uintptr
	Get_BrainpoolP192t1  uintptr
	Get_BrainpoolP224r1  uintptr
	Get_BrainpoolP224t1  uintptr
	Get_BrainpoolP256r1  uintptr
	Get_BrainpoolP256t1  uintptr
	Get_BrainpoolP320r1  uintptr
	Get_BrainpoolP320t1  uintptr
	Get_BrainpoolP384r1  uintptr
	Get_BrainpoolP384t1  uintptr
	Get_BrainpoolP512r1  uintptr
	Get_BrainpoolP512t1  uintptr
	Get_Curve25519       uintptr
	Get_Ec192wapi        uintptr
	Get_NistP192         uintptr
	Get_NistP224         uintptr
	Get_NistP256         uintptr
	Get_NistP384         uintptr
	Get_NistP521         uintptr
	Get_NumsP256t1       uintptr
	Get_NumsP384t1       uintptr
	Get_NumsP512t1       uintptr
	Get_SecP160k1        uintptr
	Get_SecP160r1        uintptr
	Get_SecP160r2        uintptr
	Get_SecP192k1        uintptr
	Get_SecP192r1        uintptr
	Get_SecP224k1        uintptr
	Get_SecP224r1        uintptr
	Get_SecP256k1        uintptr
	Get_SecP256r1        uintptr
	Get_SecP384r1        uintptr
	Get_SecP521r1        uintptr
	Get_Wtls7            uintptr
	Get_Wtls9            uintptr
	Get_Wtls12           uintptr
	Get_X962P192v1       uintptr
	Get_X962P192v2       uintptr
	Get_X962P192v3       uintptr
	Get_X962P239v1       uintptr
	Get_X962P239v2       uintptr
	Get_X962P239v3       uintptr
	Get_X962P256v1       uintptr
	Get_AllEccCurveNames uintptr
}

type IEccCurveNamesStatics struct {
	win32.IInspectable
}

func (this *IEccCurveNamesStatics) Vtbl() *IEccCurveNamesStaticsVtbl {
	return (*IEccCurveNamesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IEccCurveNamesStatics) Get_BrainpoolP160r1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BrainpoolP160r1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_BrainpoolP160t1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BrainpoolP160t1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_BrainpoolP192r1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BrainpoolP192r1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_BrainpoolP192t1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BrainpoolP192t1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_BrainpoolP224r1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BrainpoolP224r1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_BrainpoolP224t1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BrainpoolP224t1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_BrainpoolP256r1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BrainpoolP256r1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_BrainpoolP256t1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BrainpoolP256t1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_BrainpoolP320r1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BrainpoolP320r1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_BrainpoolP320t1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BrainpoolP320t1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_BrainpoolP384r1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BrainpoolP384r1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_BrainpoolP384t1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BrainpoolP384t1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_BrainpoolP512r1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BrainpoolP512r1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_BrainpoolP512t1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BrainpoolP512t1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_Curve25519() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Curve25519, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_Ec192wapi() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ec192wapi, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_NistP192() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NistP192, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_NistP224() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NistP224, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_NistP256() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NistP256, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_NistP384() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NistP384, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_NistP521() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NistP521, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_NumsP256t1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NumsP256t1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_NumsP384t1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NumsP384t1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_NumsP512t1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NumsP512t1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_SecP160k1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SecP160k1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_SecP160r1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SecP160r1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_SecP160r2() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SecP160r2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_SecP192k1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SecP192k1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_SecP192r1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SecP192r1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_SecP224k1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SecP224k1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_SecP224r1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SecP224r1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_SecP256k1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SecP256k1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_SecP256r1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SecP256r1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_SecP384r1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SecP384r1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_SecP521r1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SecP521r1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_Wtls7() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Wtls7, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_Wtls9() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Wtls9, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_Wtls12() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Wtls12, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_X962P192v1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_X962P192v1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_X962P192v2() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_X962P192v2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_X962P192v3() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_X962P192v3, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_X962P239v1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_X962P239v1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_X962P239v2() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_X962P239v2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_X962P239v3() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_X962P239v3, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_X962P256v1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_X962P256v1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IEccCurveNamesStatics) Get_AllEccCurveNames() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllEccCurveNames, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6FA42FE7-1ECB-4B00-BEA5-60B83F862F17
var IID_IEncryptedAndAuthenticatedData = syscall.GUID{0x6FA42FE7, 0x1ECB, 0x4B00,
	[8]byte{0xBE, 0xA5, 0x60, 0xB8, 0x3F, 0x86, 0x2F, 0x17}}

type IEncryptedAndAuthenticatedDataInterface interface {
	win32.IInspectableInterface
	Get_EncryptedData() *IBuffer
	Get_AuthenticationTag() *IBuffer
}

type IEncryptedAndAuthenticatedDataVtbl struct {
	win32.IInspectableVtbl
	Get_EncryptedData     uintptr
	Get_AuthenticationTag uintptr
}

type IEncryptedAndAuthenticatedData struct {
	win32.IInspectable
}

func (this *IEncryptedAndAuthenticatedData) Vtbl() *IEncryptedAndAuthenticatedDataVtbl {
	return (*IEncryptedAndAuthenticatedDataVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IEncryptedAndAuthenticatedData) Get_EncryptedData() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EncryptedData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IEncryptedAndAuthenticatedData) Get_AuthenticationTag() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AuthenticationTag, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6B5E0516-DE96-4F0A-8D57-DCC9DAE36C76
var IID_IHashAlgorithmNamesStatics = syscall.GUID{0x6B5E0516, 0xDE96, 0x4F0A,
	[8]byte{0x8D, 0x57, 0xDC, 0xC9, 0xDA, 0xE3, 0x6C, 0x76}}

type IHashAlgorithmNamesStaticsInterface interface {
	win32.IInspectableInterface
	Get_Md5() string
	Get_Sha1() string
	Get_Sha256() string
	Get_Sha384() string
	Get_Sha512() string
}

type IHashAlgorithmNamesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Md5    uintptr
	Get_Sha1   uintptr
	Get_Sha256 uintptr
	Get_Sha384 uintptr
	Get_Sha512 uintptr
}

type IHashAlgorithmNamesStatics struct {
	win32.IInspectable
}

func (this *IHashAlgorithmNamesStatics) Vtbl() *IHashAlgorithmNamesStaticsVtbl {
	return (*IHashAlgorithmNamesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHashAlgorithmNamesStatics) Get_Md5() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Md5, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHashAlgorithmNamesStatics) Get_Sha1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Sha1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHashAlgorithmNamesStatics) Get_Sha256() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Sha256, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHashAlgorithmNamesStatics) Get_Sha384() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Sha384, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHashAlgorithmNamesStatics) Get_Sha512() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Sha512, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// BE9B3080-B2C3-422B-BCE1-EC90EFB5D7B5
var IID_IHashAlgorithmProvider = syscall.GUID{0xBE9B3080, 0xB2C3, 0x422B,
	[8]byte{0xBC, 0xE1, 0xEC, 0x90, 0xEF, 0xB5, 0xD7, 0xB5}}

type IHashAlgorithmProviderInterface interface {
	win32.IInspectableInterface
	Get_AlgorithmName() string
	Get_HashLength() uint32
	HashData(data *IBuffer) *IBuffer
	CreateHash() *IHashComputation
}

type IHashAlgorithmProviderVtbl struct {
	win32.IInspectableVtbl
	Get_AlgorithmName uintptr
	Get_HashLength    uintptr
	HashData          uintptr
	CreateHash        uintptr
}

type IHashAlgorithmProvider struct {
	win32.IInspectable
}

func (this *IHashAlgorithmProvider) Vtbl() *IHashAlgorithmProviderVtbl {
	return (*IHashAlgorithmProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHashAlgorithmProvider) Get_AlgorithmName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlgorithmName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHashAlgorithmProvider) Get_HashLength() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HashLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHashAlgorithmProvider) HashData(data *IBuffer) *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().HashData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHashAlgorithmProvider) CreateHash() *IHashComputation {
	var _result *IHashComputation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateHash, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9FAC9741-5CC4-4336-AE38-6212B75A915A
var IID_IHashAlgorithmProviderStatics = syscall.GUID{0x9FAC9741, 0x5CC4, 0x4336,
	[8]byte{0xAE, 0x38, 0x62, 0x12, 0xB7, 0x5A, 0x91, 0x5A}}

type IHashAlgorithmProviderStaticsInterface interface {
	win32.IInspectableInterface
	OpenAlgorithm(algorithm string) *IHashAlgorithmProvider
}

type IHashAlgorithmProviderStaticsVtbl struct {
	win32.IInspectableVtbl
	OpenAlgorithm uintptr
}

type IHashAlgorithmProviderStatics struct {
	win32.IInspectable
}

func (this *IHashAlgorithmProviderStatics) Vtbl() *IHashAlgorithmProviderStaticsVtbl {
	return (*IHashAlgorithmProviderStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHashAlgorithmProviderStatics) OpenAlgorithm(algorithm string) *IHashAlgorithmProvider {
	var _result *IHashAlgorithmProvider
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenAlgorithm, uintptr(unsafe.Pointer(this)), NewHStr(algorithm).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5904D1B6-AD31-4603-A3A4-B1BDA98E2562
var IID_IHashComputation = syscall.GUID{0x5904D1B6, 0xAD31, 0x4603,
	[8]byte{0xA3, 0xA4, 0xB1, 0xBD, 0xA9, 0x8E, 0x25, 0x62}}

type IHashComputationInterface interface {
	win32.IInspectableInterface
	Append(data *IBuffer)
	GetValueAndReset() *IBuffer
}

type IHashComputationVtbl struct {
	win32.IInspectableVtbl
	Append           uintptr
	GetValueAndReset uintptr
}

type IHashComputation struct {
	win32.IInspectable
}

func (this *IHashComputation) Vtbl() *IHashComputationVtbl {
	return (*IHashComputationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHashComputation) Append(data *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Append, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(data)))
	_ = _hr
}

func (this *IHashComputation) GetValueAndReset() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetValueAndReset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7B6E363E-94D2-4739-A57B-022E0C3A402A
var IID_IKeyDerivationAlgorithmNamesStatics = syscall.GUID{0x7B6E363E, 0x94D2, 0x4739,
	[8]byte{0xA5, 0x7B, 0x02, 0x2E, 0x0C, 0x3A, 0x40, 0x2A}}

type IKeyDerivationAlgorithmNamesStaticsInterface interface {
	win32.IInspectableInterface
	Get_Pbkdf2Md5() string
	Get_Pbkdf2Sha1() string
	Get_Pbkdf2Sha256() string
	Get_Pbkdf2Sha384() string
	Get_Pbkdf2Sha512() string
	Get_Sp800108CtrHmacMd5() string
	Get_Sp800108CtrHmacSha1() string
	Get_Sp800108CtrHmacSha256() string
	Get_Sp800108CtrHmacSha384() string
	Get_Sp800108CtrHmacSha512() string
	Get_Sp80056aConcatMd5() string
	Get_Sp80056aConcatSha1() string
	Get_Sp80056aConcatSha256() string
	Get_Sp80056aConcatSha384() string
	Get_Sp80056aConcatSha512() string
}

type IKeyDerivationAlgorithmNamesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Pbkdf2Md5             uintptr
	Get_Pbkdf2Sha1            uintptr
	Get_Pbkdf2Sha256          uintptr
	Get_Pbkdf2Sha384          uintptr
	Get_Pbkdf2Sha512          uintptr
	Get_Sp800108CtrHmacMd5    uintptr
	Get_Sp800108CtrHmacSha1   uintptr
	Get_Sp800108CtrHmacSha256 uintptr
	Get_Sp800108CtrHmacSha384 uintptr
	Get_Sp800108CtrHmacSha512 uintptr
	Get_Sp80056aConcatMd5     uintptr
	Get_Sp80056aConcatSha1    uintptr
	Get_Sp80056aConcatSha256  uintptr
	Get_Sp80056aConcatSha384  uintptr
	Get_Sp80056aConcatSha512  uintptr
}

type IKeyDerivationAlgorithmNamesStatics struct {
	win32.IInspectable
}

func (this *IKeyDerivationAlgorithmNamesStatics) Vtbl() *IKeyDerivationAlgorithmNamesStaticsVtbl {
	return (*IKeyDerivationAlgorithmNamesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKeyDerivationAlgorithmNamesStatics) Get_Pbkdf2Md5() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Pbkdf2Md5, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyDerivationAlgorithmNamesStatics) Get_Pbkdf2Sha1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Pbkdf2Sha1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyDerivationAlgorithmNamesStatics) Get_Pbkdf2Sha256() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Pbkdf2Sha256, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyDerivationAlgorithmNamesStatics) Get_Pbkdf2Sha384() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Pbkdf2Sha384, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyDerivationAlgorithmNamesStatics) Get_Pbkdf2Sha512() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Pbkdf2Sha512, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyDerivationAlgorithmNamesStatics) Get_Sp800108CtrHmacMd5() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Sp800108CtrHmacMd5, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyDerivationAlgorithmNamesStatics) Get_Sp800108CtrHmacSha1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Sp800108CtrHmacSha1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyDerivationAlgorithmNamesStatics) Get_Sp800108CtrHmacSha256() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Sp800108CtrHmacSha256, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyDerivationAlgorithmNamesStatics) Get_Sp800108CtrHmacSha384() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Sp800108CtrHmacSha384, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyDerivationAlgorithmNamesStatics) Get_Sp800108CtrHmacSha512() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Sp800108CtrHmacSha512, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyDerivationAlgorithmNamesStatics) Get_Sp80056aConcatMd5() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Sp80056aConcatMd5, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyDerivationAlgorithmNamesStatics) Get_Sp80056aConcatSha1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Sp80056aConcatSha1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyDerivationAlgorithmNamesStatics) Get_Sp80056aConcatSha256() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Sp80056aConcatSha256, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyDerivationAlgorithmNamesStatics) Get_Sp80056aConcatSha384() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Sp80056aConcatSha384, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyDerivationAlgorithmNamesStatics) Get_Sp80056aConcatSha512() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Sp80056aConcatSha512, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 57953FAB-6044-466F-97F4-337B7808384D
var IID_IKeyDerivationAlgorithmNamesStatics2 = syscall.GUID{0x57953FAB, 0x6044, 0x466F,
	[8]byte{0x97, 0xF4, 0x33, 0x7B, 0x78, 0x08, 0x38, 0x4D}}

type IKeyDerivationAlgorithmNamesStatics2Interface interface {
	win32.IInspectableInterface
	Get_CapiKdfMd5() string
	Get_CapiKdfSha1() string
	Get_CapiKdfSha256() string
	Get_CapiKdfSha384() string
	Get_CapiKdfSha512() string
}

type IKeyDerivationAlgorithmNamesStatics2Vtbl struct {
	win32.IInspectableVtbl
	Get_CapiKdfMd5    uintptr
	Get_CapiKdfSha1   uintptr
	Get_CapiKdfSha256 uintptr
	Get_CapiKdfSha384 uintptr
	Get_CapiKdfSha512 uintptr
}

type IKeyDerivationAlgorithmNamesStatics2 struct {
	win32.IInspectable
}

func (this *IKeyDerivationAlgorithmNamesStatics2) Vtbl() *IKeyDerivationAlgorithmNamesStatics2Vtbl {
	return (*IKeyDerivationAlgorithmNamesStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKeyDerivationAlgorithmNamesStatics2) Get_CapiKdfMd5() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CapiKdfMd5, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyDerivationAlgorithmNamesStatics2) Get_CapiKdfSha1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CapiKdfSha1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyDerivationAlgorithmNamesStatics2) Get_CapiKdfSha256() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CapiKdfSha256, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyDerivationAlgorithmNamesStatics2) Get_CapiKdfSha384() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CapiKdfSha384, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyDerivationAlgorithmNamesStatics2) Get_CapiKdfSha512() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CapiKdfSha512, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// E1FBA83B-4671-43B7-9158-763AAA98B6BF
var IID_IKeyDerivationAlgorithmProvider = syscall.GUID{0xE1FBA83B, 0x4671, 0x43B7,
	[8]byte{0x91, 0x58, 0x76, 0x3A, 0xAA, 0x98, 0xB6, 0xBF}}

type IKeyDerivationAlgorithmProviderInterface interface {
	win32.IInspectableInterface
	Get_AlgorithmName() string
	CreateKey(keyMaterial *IBuffer) *ICryptographicKey
}

type IKeyDerivationAlgorithmProviderVtbl struct {
	win32.IInspectableVtbl
	Get_AlgorithmName uintptr
	CreateKey         uintptr
}

type IKeyDerivationAlgorithmProvider struct {
	win32.IInspectable
}

func (this *IKeyDerivationAlgorithmProvider) Vtbl() *IKeyDerivationAlgorithmProviderVtbl {
	return (*IKeyDerivationAlgorithmProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKeyDerivationAlgorithmProvider) Get_AlgorithmName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlgorithmName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKeyDerivationAlgorithmProvider) CreateKey(keyMaterial *IBuffer) *ICryptographicKey {
	var _result *ICryptographicKey
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateKey, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(keyMaterial)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0A22097A-0A1C-443B-9418-B9498AEB1603
var IID_IKeyDerivationAlgorithmProviderStatics = syscall.GUID{0x0A22097A, 0x0A1C, 0x443B,
	[8]byte{0x94, 0x18, 0xB9, 0x49, 0x8A, 0xEB, 0x16, 0x03}}

type IKeyDerivationAlgorithmProviderStaticsInterface interface {
	win32.IInspectableInterface
	OpenAlgorithm(algorithm string) *IKeyDerivationAlgorithmProvider
}

type IKeyDerivationAlgorithmProviderStaticsVtbl struct {
	win32.IInspectableVtbl
	OpenAlgorithm uintptr
}

type IKeyDerivationAlgorithmProviderStatics struct {
	win32.IInspectable
}

func (this *IKeyDerivationAlgorithmProviderStatics) Vtbl() *IKeyDerivationAlgorithmProviderStaticsVtbl {
	return (*IKeyDerivationAlgorithmProviderStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKeyDerivationAlgorithmProviderStatics) OpenAlgorithm(algorithm string) *IKeyDerivationAlgorithmProvider {
	var _result *IKeyDerivationAlgorithmProvider
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenAlgorithm, uintptr(unsafe.Pointer(this)), NewHStr(algorithm).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7BF05967-047B-4A8C-964A-469FFD5522E2
var IID_IKeyDerivationParameters = syscall.GUID{0x7BF05967, 0x047B, 0x4A8C,
	[8]byte{0x96, 0x4A, 0x46, 0x9F, 0xFD, 0x55, 0x22, 0xE2}}

type IKeyDerivationParametersInterface interface {
	win32.IInspectableInterface
	Get_KdfGenericBinary() *IBuffer
	Put_KdfGenericBinary(value *IBuffer)
	Get_IterationCount() uint32
}

type IKeyDerivationParametersVtbl struct {
	win32.IInspectableVtbl
	Get_KdfGenericBinary uintptr
	Put_KdfGenericBinary uintptr
	Get_IterationCount   uintptr
}

type IKeyDerivationParameters struct {
	win32.IInspectable
}

func (this *IKeyDerivationParameters) Vtbl() *IKeyDerivationParametersVtbl {
	return (*IKeyDerivationParametersVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKeyDerivationParameters) Get_KdfGenericBinary() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KdfGenericBinary, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKeyDerivationParameters) Put_KdfGenericBinary(value *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_KdfGenericBinary, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IKeyDerivationParameters) Get_IterationCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IterationCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// CD4166D1-417E-4F4C-B666-C0D879F3F8E0
var IID_IKeyDerivationParameters2 = syscall.GUID{0xCD4166D1, 0x417E, 0x4F4C,
	[8]byte{0xB6, 0x66, 0xC0, 0xD8, 0x79, 0xF3, 0xF8, 0xE0}}

type IKeyDerivationParameters2Interface interface {
	win32.IInspectableInterface
	Get_Capi1KdfTargetAlgorithm() Capi1KdfTargetAlgorithm
	Put_Capi1KdfTargetAlgorithm(value Capi1KdfTargetAlgorithm)
}

type IKeyDerivationParameters2Vtbl struct {
	win32.IInspectableVtbl
	Get_Capi1KdfTargetAlgorithm uintptr
	Put_Capi1KdfTargetAlgorithm uintptr
}

type IKeyDerivationParameters2 struct {
	win32.IInspectable
}

func (this *IKeyDerivationParameters2) Vtbl() *IKeyDerivationParameters2Vtbl {
	return (*IKeyDerivationParameters2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKeyDerivationParameters2) Get_Capi1KdfTargetAlgorithm() Capi1KdfTargetAlgorithm {
	var _result Capi1KdfTargetAlgorithm
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Capi1KdfTargetAlgorithm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IKeyDerivationParameters2) Put_Capi1KdfTargetAlgorithm(value Capi1KdfTargetAlgorithm) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Capi1KdfTargetAlgorithm, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// EA961FBE-F37F-4146-9DFE-A456F1735F4B
var IID_IKeyDerivationParametersStatics = syscall.GUID{0xEA961FBE, 0xF37F, 0x4146,
	[8]byte{0x9D, 0xFE, 0xA4, 0x56, 0xF1, 0x73, 0x5F, 0x4B}}

type IKeyDerivationParametersStaticsInterface interface {
	win32.IInspectableInterface
	BuildForPbkdf2(pbkdf2Salt *IBuffer, iterationCount uint32) *IKeyDerivationParameters
	BuildForSP800108(label *IBuffer, context *IBuffer) *IKeyDerivationParameters
	BuildForSP80056a(algorithmId *IBuffer, partyUInfo *IBuffer, partyVInfo *IBuffer, suppPubInfo *IBuffer, suppPrivInfo *IBuffer) *IKeyDerivationParameters
}

type IKeyDerivationParametersStaticsVtbl struct {
	win32.IInspectableVtbl
	BuildForPbkdf2   uintptr
	BuildForSP800108 uintptr
	BuildForSP80056a uintptr
}

type IKeyDerivationParametersStatics struct {
	win32.IInspectable
}

func (this *IKeyDerivationParametersStatics) Vtbl() *IKeyDerivationParametersStaticsVtbl {
	return (*IKeyDerivationParametersStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKeyDerivationParametersStatics) BuildForPbkdf2(pbkdf2Salt *IBuffer, iterationCount uint32) *IKeyDerivationParameters {
	var _result *IKeyDerivationParameters
	_hr, _, _ := syscall.SyscallN(this.Vtbl().BuildForPbkdf2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pbkdf2Salt)), uintptr(iterationCount), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKeyDerivationParametersStatics) BuildForSP800108(label *IBuffer, context *IBuffer) *IKeyDerivationParameters {
	var _result *IKeyDerivationParameters
	_hr, _, _ := syscall.SyscallN(this.Vtbl().BuildForSP800108, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(label)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IKeyDerivationParametersStatics) BuildForSP80056a(algorithmId *IBuffer, partyUInfo *IBuffer, partyVInfo *IBuffer, suppPubInfo *IBuffer, suppPrivInfo *IBuffer) *IKeyDerivationParameters {
	var _result *IKeyDerivationParameters
	_hr, _, _ := syscall.SyscallN(this.Vtbl().BuildForSP80056a, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(algorithmId)), uintptr(unsafe.Pointer(partyUInfo)), uintptr(unsafe.Pointer(partyVInfo)), uintptr(unsafe.Pointer(suppPubInfo)), uintptr(unsafe.Pointer(suppPrivInfo)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A5783DD5-58E3-4EFB-B283-A1653126E1BE
var IID_IKeyDerivationParametersStatics2 = syscall.GUID{0xA5783DD5, 0x58E3, 0x4EFB,
	[8]byte{0xB2, 0x83, 0xA1, 0x65, 0x31, 0x26, 0xE1, 0xBE}}

type IKeyDerivationParametersStatics2Interface interface {
	win32.IInspectableInterface
	BuildForCapi1Kdf(capi1KdfTargetAlgorithm Capi1KdfTargetAlgorithm) *IKeyDerivationParameters
}

type IKeyDerivationParametersStatics2Vtbl struct {
	win32.IInspectableVtbl
	BuildForCapi1Kdf uintptr
}

type IKeyDerivationParametersStatics2 struct {
	win32.IInspectable
}

func (this *IKeyDerivationParametersStatics2) Vtbl() *IKeyDerivationParametersStatics2Vtbl {
	return (*IKeyDerivationParametersStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKeyDerivationParametersStatics2) BuildForCapi1Kdf(capi1KdfTargetAlgorithm Capi1KdfTargetAlgorithm) *IKeyDerivationParameters {
	var _result *IKeyDerivationParameters
	_hr, _, _ := syscall.SyscallN(this.Vtbl().BuildForCapi1Kdf, uintptr(unsafe.Pointer(this)), uintptr(capi1KdfTargetAlgorithm), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 41412678-FB1E-43A4-895E-A9026E4390A3
var IID_IMacAlgorithmNamesStatics = syscall.GUID{0x41412678, 0xFB1E, 0x43A4,
	[8]byte{0x89, 0x5E, 0xA9, 0x02, 0x6E, 0x43, 0x90, 0xA3}}

type IMacAlgorithmNamesStaticsInterface interface {
	win32.IInspectableInterface
	Get_HmacMd5() string
	Get_HmacSha1() string
	Get_HmacSha256() string
	Get_HmacSha384() string
	Get_HmacSha512() string
	Get_AesCmac() string
}

type IMacAlgorithmNamesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_HmacMd5    uintptr
	Get_HmacSha1   uintptr
	Get_HmacSha256 uintptr
	Get_HmacSha384 uintptr
	Get_HmacSha512 uintptr
	Get_AesCmac    uintptr
}

type IMacAlgorithmNamesStatics struct {
	win32.IInspectable
}

func (this *IMacAlgorithmNamesStatics) Vtbl() *IMacAlgorithmNamesStaticsVtbl {
	return (*IMacAlgorithmNamesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMacAlgorithmNamesStatics) Get_HmacMd5() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HmacMd5, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMacAlgorithmNamesStatics) Get_HmacSha1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HmacSha1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMacAlgorithmNamesStatics) Get_HmacSha256() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HmacSha256, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMacAlgorithmNamesStatics) Get_HmacSha384() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HmacSha384, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMacAlgorithmNamesStatics) Get_HmacSha512() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HmacSha512, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMacAlgorithmNamesStatics) Get_AesCmac() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AesCmac, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 4A3FC5C3-1CBD-41CE-A092-AA0BC5D2D2F5
var IID_IMacAlgorithmProvider = syscall.GUID{0x4A3FC5C3, 0x1CBD, 0x41CE,
	[8]byte{0xA0, 0x92, 0xAA, 0x0B, 0xC5, 0xD2, 0xD2, 0xF5}}

type IMacAlgorithmProviderInterface interface {
	win32.IInspectableInterface
	Get_AlgorithmName() string
	Get_MacLength() uint32
	CreateKey(keyMaterial *IBuffer) *ICryptographicKey
}

type IMacAlgorithmProviderVtbl struct {
	win32.IInspectableVtbl
	Get_AlgorithmName uintptr
	Get_MacLength     uintptr
	CreateKey         uintptr
}

type IMacAlgorithmProvider struct {
	win32.IInspectable
}

func (this *IMacAlgorithmProvider) Vtbl() *IMacAlgorithmProviderVtbl {
	return (*IMacAlgorithmProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMacAlgorithmProvider) Get_AlgorithmName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlgorithmName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMacAlgorithmProvider) Get_MacLength() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MacLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMacAlgorithmProvider) CreateKey(keyMaterial *IBuffer) *ICryptographicKey {
	var _result *ICryptographicKey
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateKey, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(keyMaterial)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6DA32A15-D931-42ED-8E7E-C301CAEE119C
var IID_IMacAlgorithmProvider2 = syscall.GUID{0x6DA32A15, 0xD931, 0x42ED,
	[8]byte{0x8E, 0x7E, 0xC3, 0x01, 0xCA, 0xEE, 0x11, 0x9C}}

type IMacAlgorithmProvider2Interface interface {
	win32.IInspectableInterface
	CreateHash(keyMaterial *IBuffer) *IHashComputation
}

type IMacAlgorithmProvider2Vtbl struct {
	win32.IInspectableVtbl
	CreateHash uintptr
}

type IMacAlgorithmProvider2 struct {
	win32.IInspectable
}

func (this *IMacAlgorithmProvider2) Vtbl() *IMacAlgorithmProvider2Vtbl {
	return (*IMacAlgorithmProvider2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMacAlgorithmProvider2) CreateHash(keyMaterial *IBuffer) *IHashComputation {
	var _result *IHashComputation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateHash, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(keyMaterial)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C9BDC147-CC77-4DF0-9E4E-B921E080644C
var IID_IMacAlgorithmProviderStatics = syscall.GUID{0xC9BDC147, 0xCC77, 0x4DF0,
	[8]byte{0x9E, 0x4E, 0xB9, 0x21, 0xE0, 0x80, 0x64, 0x4C}}

type IMacAlgorithmProviderStaticsInterface interface {
	win32.IInspectableInterface
	OpenAlgorithm(algorithm string) *IMacAlgorithmProvider
}

type IMacAlgorithmProviderStaticsVtbl struct {
	win32.IInspectableVtbl
	OpenAlgorithm uintptr
}

type IMacAlgorithmProviderStatics struct {
	win32.IInspectable
}

func (this *IMacAlgorithmProviderStatics) Vtbl() *IMacAlgorithmProviderStaticsVtbl {
	return (*IMacAlgorithmProviderStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMacAlgorithmProviderStatics) OpenAlgorithm(algorithm string) *IMacAlgorithmProvider {
	var _result *IMacAlgorithmProvider
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenAlgorithm, uintptr(unsafe.Pointer(this)), NewHStr(algorithm).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 77274814-D9D4-4CF5-B668-E0457DF30894
var IID_IPersistedKeyProviderStatics = syscall.GUID{0x77274814, 0xD9D4, 0x4CF5,
	[8]byte{0xB6, 0x68, 0xE0, 0x45, 0x7D, 0xF3, 0x08, 0x94}}

type IPersistedKeyProviderStaticsInterface interface {
	win32.IInspectableInterface
	OpenKeyPairFromCertificateAsync(certificate *ICertificate, hashAlgorithmName string, padding CryptographicPadding) *IAsyncOperation[*ICryptographicKey]
	OpenPublicKeyFromCertificate(certificate *ICertificate, hashAlgorithmName string, padding CryptographicPadding) *ICryptographicKey
}

type IPersistedKeyProviderStaticsVtbl struct {
	win32.IInspectableVtbl
	OpenKeyPairFromCertificateAsync uintptr
	OpenPublicKeyFromCertificate    uintptr
}

type IPersistedKeyProviderStatics struct {
	win32.IInspectable
}

func (this *IPersistedKeyProviderStatics) Vtbl() *IPersistedKeyProviderStaticsVtbl {
	return (*IPersistedKeyProviderStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPersistedKeyProviderStatics) OpenKeyPairFromCertificateAsync(certificate *ICertificate, hashAlgorithmName string, padding CryptographicPadding) *IAsyncOperation[*ICryptographicKey] {
	var _result *IAsyncOperation[*ICryptographicKey]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenKeyPairFromCertificateAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(certificate)), NewHStr(hashAlgorithmName).Ptr, uintptr(padding), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPersistedKeyProviderStatics) OpenPublicKeyFromCertificate(certificate *ICertificate, hashAlgorithmName string, padding CryptographicPadding) *ICryptographicKey {
	var _result *ICryptographicKey
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenPublicKeyFromCertificate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(certificate)), NewHStr(hashAlgorithmName).Ptr, uintptr(padding), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6870727B-C996-4EAE-84D7-79B2AEB73B9C
var IID_ISymmetricAlgorithmNamesStatics = syscall.GUID{0x6870727B, 0xC996, 0x4EAE,
	[8]byte{0x84, 0xD7, 0x79, 0xB2, 0xAE, 0xB7, 0x3B, 0x9C}}

type ISymmetricAlgorithmNamesStaticsInterface interface {
	win32.IInspectableInterface
	Get_DesCbc() string
	Get_DesEcb() string
	Get_TripleDesCbc() string
	Get_TripleDesEcb() string
	Get_Rc2Cbc() string
	Get_Rc2Ecb() string
	Get_AesCbc() string
	Get_AesEcb() string
	Get_AesGcm() string
	Get_AesCcm() string
	Get_AesCbcPkcs7() string
	Get_AesEcbPkcs7() string
	Get_DesCbcPkcs7() string
	Get_DesEcbPkcs7() string
	Get_TripleDesCbcPkcs7() string
	Get_TripleDesEcbPkcs7() string
	Get_Rc2CbcPkcs7() string
	Get_Rc2EcbPkcs7() string
	Get_Rc4() string
}

type ISymmetricAlgorithmNamesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_DesCbc            uintptr
	Get_DesEcb            uintptr
	Get_TripleDesCbc      uintptr
	Get_TripleDesEcb      uintptr
	Get_Rc2Cbc            uintptr
	Get_Rc2Ecb            uintptr
	Get_AesCbc            uintptr
	Get_AesEcb            uintptr
	Get_AesGcm            uintptr
	Get_AesCcm            uintptr
	Get_AesCbcPkcs7       uintptr
	Get_AesEcbPkcs7       uintptr
	Get_DesCbcPkcs7       uintptr
	Get_DesEcbPkcs7       uintptr
	Get_TripleDesCbcPkcs7 uintptr
	Get_TripleDesEcbPkcs7 uintptr
	Get_Rc2CbcPkcs7       uintptr
	Get_Rc2EcbPkcs7       uintptr
	Get_Rc4               uintptr
}

type ISymmetricAlgorithmNamesStatics struct {
	win32.IInspectable
}

func (this *ISymmetricAlgorithmNamesStatics) Vtbl() *ISymmetricAlgorithmNamesStaticsVtbl {
	return (*ISymmetricAlgorithmNamesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISymmetricAlgorithmNamesStatics) Get_DesCbc() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesCbc, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISymmetricAlgorithmNamesStatics) Get_DesEcb() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesEcb, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISymmetricAlgorithmNamesStatics) Get_TripleDesCbc() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TripleDesCbc, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISymmetricAlgorithmNamesStatics) Get_TripleDesEcb() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TripleDesEcb, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISymmetricAlgorithmNamesStatics) Get_Rc2Cbc() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Rc2Cbc, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISymmetricAlgorithmNamesStatics) Get_Rc2Ecb() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Rc2Ecb, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISymmetricAlgorithmNamesStatics) Get_AesCbc() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AesCbc, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISymmetricAlgorithmNamesStatics) Get_AesEcb() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AesEcb, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISymmetricAlgorithmNamesStatics) Get_AesGcm() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AesGcm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISymmetricAlgorithmNamesStatics) Get_AesCcm() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AesCcm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISymmetricAlgorithmNamesStatics) Get_AesCbcPkcs7() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AesCbcPkcs7, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISymmetricAlgorithmNamesStatics) Get_AesEcbPkcs7() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AesEcbPkcs7, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISymmetricAlgorithmNamesStatics) Get_DesCbcPkcs7() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesCbcPkcs7, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISymmetricAlgorithmNamesStatics) Get_DesEcbPkcs7() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesEcbPkcs7, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISymmetricAlgorithmNamesStatics) Get_TripleDesCbcPkcs7() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TripleDesCbcPkcs7, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISymmetricAlgorithmNamesStatics) Get_TripleDesEcbPkcs7() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TripleDesEcbPkcs7, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISymmetricAlgorithmNamesStatics) Get_Rc2CbcPkcs7() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Rc2CbcPkcs7, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISymmetricAlgorithmNamesStatics) Get_Rc2EcbPkcs7() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Rc2EcbPkcs7, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISymmetricAlgorithmNamesStatics) Get_Rc4() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Rc4, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 3D7E4A33-3BD0-4902-8AC8-470D50D21376
var IID_ISymmetricKeyAlgorithmProvider = syscall.GUID{0x3D7E4A33, 0x3BD0, 0x4902,
	[8]byte{0x8A, 0xC8, 0x47, 0x0D, 0x50, 0xD2, 0x13, 0x76}}

type ISymmetricKeyAlgorithmProviderInterface interface {
	win32.IInspectableInterface
	Get_AlgorithmName() string
	Get_BlockLength() uint32
	CreateSymmetricKey(keyMaterial *IBuffer) *ICryptographicKey
}

type ISymmetricKeyAlgorithmProviderVtbl struct {
	win32.IInspectableVtbl
	Get_AlgorithmName  uintptr
	Get_BlockLength    uintptr
	CreateSymmetricKey uintptr
}

type ISymmetricKeyAlgorithmProvider struct {
	win32.IInspectable
}

func (this *ISymmetricKeyAlgorithmProvider) Vtbl() *ISymmetricKeyAlgorithmProviderVtbl {
	return (*ISymmetricKeyAlgorithmProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISymmetricKeyAlgorithmProvider) Get_AlgorithmName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlgorithmName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISymmetricKeyAlgorithmProvider) Get_BlockLength() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BlockLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISymmetricKeyAlgorithmProvider) CreateSymmetricKey(keyMaterial *IBuffer) *ICryptographicKey {
	var _result *ICryptographicKey
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSymmetricKey, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(keyMaterial)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8D3B2326-1F37-491F-B60E-F5431B26B483
var IID_ISymmetricKeyAlgorithmProviderStatics = syscall.GUID{0x8D3B2326, 0x1F37, 0x491F,
	[8]byte{0xB6, 0x0E, 0xF5, 0x43, 0x1B, 0x26, 0xB4, 0x83}}

type ISymmetricKeyAlgorithmProviderStaticsInterface interface {
	win32.IInspectableInterface
	OpenAlgorithm(algorithm string) *ISymmetricKeyAlgorithmProvider
}

type ISymmetricKeyAlgorithmProviderStaticsVtbl struct {
	win32.IInspectableVtbl
	OpenAlgorithm uintptr
}

type ISymmetricKeyAlgorithmProviderStatics struct {
	win32.IInspectable
}

func (this *ISymmetricKeyAlgorithmProviderStatics) Vtbl() *ISymmetricKeyAlgorithmProviderStaticsVtbl {
	return (*ISymmetricKeyAlgorithmProviderStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISymmetricKeyAlgorithmProviderStatics) OpenAlgorithm(algorithm string) *ISymmetricKeyAlgorithmProvider {
	var _result *ISymmetricKeyAlgorithmProvider
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenAlgorithm, uintptr(unsafe.Pointer(this)), NewHStr(algorithm).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type AsymmetricAlgorithmNames struct {
	RtClass
}

func NewIAsymmetricAlgorithmNamesStatics2() *IAsymmetricAlgorithmNamesStatics2 {
	var p *IAsymmetricAlgorithmNamesStatics2
	hs := NewHStr("Windows.Security.Cryptography.Core.AsymmetricAlgorithmNames")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAsymmetricAlgorithmNamesStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIAsymmetricAlgorithmNamesStatics() *IAsymmetricAlgorithmNamesStatics {
	var p *IAsymmetricAlgorithmNamesStatics
	hs := NewHStr("Windows.Security.Cryptography.Core.AsymmetricAlgorithmNames")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAsymmetricAlgorithmNamesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type AsymmetricKeyAlgorithmProvider struct {
	RtClass
	*IAsymmetricKeyAlgorithmProvider
}

func NewIAsymmetricKeyAlgorithmProviderStatics() *IAsymmetricKeyAlgorithmProviderStatics {
	var p *IAsymmetricKeyAlgorithmProviderStatics
	hs := NewHStr("Windows.Security.Cryptography.Core.AsymmetricKeyAlgorithmProvider")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAsymmetricKeyAlgorithmProviderStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type CryptographicEngine struct {
	RtClass
}

func NewICryptographicEngineStatics2() *ICryptographicEngineStatics2 {
	var p *ICryptographicEngineStatics2
	hs := NewHStr("Windows.Security.Cryptography.Core.CryptographicEngine")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICryptographicEngineStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewICryptographicEngineStatics() *ICryptographicEngineStatics {
	var p *ICryptographicEngineStatics
	hs := NewHStr("Windows.Security.Cryptography.Core.CryptographicEngine")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICryptographicEngineStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type CryptographicHash struct {
	RtClass
	*IHashComputation
}

type CryptographicKey struct {
	RtClass
	*ICryptographicKey
}

type EccCurveNames struct {
	RtClass
}

func NewIEccCurveNamesStatics() *IEccCurveNamesStatics {
	var p *IEccCurveNamesStatics
	hs := NewHStr("Windows.Security.Cryptography.Core.EccCurveNames")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IEccCurveNamesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type EncryptedAndAuthenticatedData struct {
	RtClass
	*IEncryptedAndAuthenticatedData
}

type HashAlgorithmNames struct {
	RtClass
}

func NewIHashAlgorithmNamesStatics() *IHashAlgorithmNamesStatics {
	var p *IHashAlgorithmNamesStatics
	hs := NewHStr("Windows.Security.Cryptography.Core.HashAlgorithmNames")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHashAlgorithmNamesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type HashAlgorithmProvider struct {
	RtClass
	*IHashAlgorithmProvider
}

func NewIHashAlgorithmProviderStatics() *IHashAlgorithmProviderStatics {
	var p *IHashAlgorithmProviderStatics
	hs := NewHStr("Windows.Security.Cryptography.Core.HashAlgorithmProvider")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHashAlgorithmProviderStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type KeyDerivationAlgorithmNames struct {
	RtClass
}

func NewIKeyDerivationAlgorithmNamesStatics2() *IKeyDerivationAlgorithmNamesStatics2 {
	var p *IKeyDerivationAlgorithmNamesStatics2
	hs := NewHStr("Windows.Security.Cryptography.Core.KeyDerivationAlgorithmNames")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IKeyDerivationAlgorithmNamesStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIKeyDerivationAlgorithmNamesStatics() *IKeyDerivationAlgorithmNamesStatics {
	var p *IKeyDerivationAlgorithmNamesStatics
	hs := NewHStr("Windows.Security.Cryptography.Core.KeyDerivationAlgorithmNames")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IKeyDerivationAlgorithmNamesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type KeyDerivationAlgorithmProvider struct {
	RtClass
	*IKeyDerivationAlgorithmProvider
}

func NewIKeyDerivationAlgorithmProviderStatics() *IKeyDerivationAlgorithmProviderStatics {
	var p *IKeyDerivationAlgorithmProviderStatics
	hs := NewHStr("Windows.Security.Cryptography.Core.KeyDerivationAlgorithmProvider")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IKeyDerivationAlgorithmProviderStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type KeyDerivationParameters struct {
	RtClass
	*IKeyDerivationParameters
}

func NewIKeyDerivationParametersStatics() *IKeyDerivationParametersStatics {
	var p *IKeyDerivationParametersStatics
	hs := NewHStr("Windows.Security.Cryptography.Core.KeyDerivationParameters")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IKeyDerivationParametersStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIKeyDerivationParametersStatics2() *IKeyDerivationParametersStatics2 {
	var p *IKeyDerivationParametersStatics2
	hs := NewHStr("Windows.Security.Cryptography.Core.KeyDerivationParameters")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IKeyDerivationParametersStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type MacAlgorithmNames struct {
	RtClass
}

func NewIMacAlgorithmNamesStatics() *IMacAlgorithmNamesStatics {
	var p *IMacAlgorithmNamesStatics
	hs := NewHStr("Windows.Security.Cryptography.Core.MacAlgorithmNames")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMacAlgorithmNamesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type MacAlgorithmProvider struct {
	RtClass
	*IMacAlgorithmProvider
}

func NewIMacAlgorithmProviderStatics() *IMacAlgorithmProviderStatics {
	var p *IMacAlgorithmProviderStatics
	hs := NewHStr("Windows.Security.Cryptography.Core.MacAlgorithmProvider")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMacAlgorithmProviderStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type PersistedKeyProvider struct {
	RtClass
}

func NewIPersistedKeyProviderStatics() *IPersistedKeyProviderStatics {
	var p *IPersistedKeyProviderStatics
	hs := NewHStr("Windows.Security.Cryptography.Core.PersistedKeyProvider")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPersistedKeyProviderStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type SymmetricAlgorithmNames struct {
	RtClass
}

func NewISymmetricAlgorithmNamesStatics() *ISymmetricAlgorithmNamesStatics {
	var p *ISymmetricAlgorithmNamesStatics
	hs := NewHStr("Windows.Security.Cryptography.Core.SymmetricAlgorithmNames")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISymmetricAlgorithmNamesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type SymmetricKeyAlgorithmProvider struct {
	RtClass
	*ISymmetricKeyAlgorithmProvider
}

func NewISymmetricKeyAlgorithmProviderStatics() *ISymmetricKeyAlgorithmProviderStatics {
	var p *ISymmetricKeyAlgorithmProviderStatics
	hs := NewHStr("Windows.Security.Cryptography.Core.SymmetricKeyAlgorithmProvider")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISymmetricKeyAlgorithmProviderStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
