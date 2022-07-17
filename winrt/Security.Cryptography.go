package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type BinaryStringEncoding int32

const (
	BinaryStringEncoding_Utf8    BinaryStringEncoding = 0
	BinaryStringEncoding_Utf16LE BinaryStringEncoding = 1
	BinaryStringEncoding_Utf16BE BinaryStringEncoding = 2
)

// interfaces

// 320B7E22-3CB0-4CDF-8663-1D28910065EB
var IID_ICryptographicBufferStatics = syscall.GUID{0x320B7E22, 0x3CB0, 0x4CDF,
	[8]byte{0x86, 0x63, 0x1D, 0x28, 0x91, 0x00, 0x65, 0xEB}}

type ICryptographicBufferStaticsInterface interface {
	win32.IInspectableInterface
	Compare(object1 *IBuffer, object2 *IBuffer) bool
	GenerateRandom(length uint32) *IBuffer
	GenerateRandomNumber() uint32
	CreateFromByteArray(valueLength uint32, value *byte) *IBuffer
	CopyToByteArray(buffer *IBuffer, valueLength uint32, value *byte)
	DecodeFromHexString(value string) *IBuffer
	EncodeToHexString(buffer *IBuffer) string
	DecodeFromBase64String(value string) *IBuffer
	EncodeToBase64String(buffer *IBuffer) string
	ConvertStringToBinary(value string, encoding BinaryStringEncoding) *IBuffer
	ConvertBinaryToString(encoding BinaryStringEncoding, buffer *IBuffer) string
}

type ICryptographicBufferStaticsVtbl struct {
	win32.IInspectableVtbl
	Compare                uintptr
	GenerateRandom         uintptr
	GenerateRandomNumber   uintptr
	CreateFromByteArray    uintptr
	CopyToByteArray        uintptr
	DecodeFromHexString    uintptr
	EncodeToHexString      uintptr
	DecodeFromBase64String uintptr
	EncodeToBase64String   uintptr
	ConvertStringToBinary  uintptr
	ConvertBinaryToString  uintptr
}

type ICryptographicBufferStatics struct {
	win32.IInspectable
}

func (this *ICryptographicBufferStatics) Vtbl() *ICryptographicBufferStaticsVtbl {
	return (*ICryptographicBufferStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICryptographicBufferStatics) Compare(object1 *IBuffer, object2 *IBuffer) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Compare, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(object1)), uintptr(unsafe.Pointer(object2)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICryptographicBufferStatics) GenerateRandom(length uint32) *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GenerateRandom, uintptr(unsafe.Pointer(this)), uintptr(length), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICryptographicBufferStatics) GenerateRandomNumber() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GenerateRandomNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICryptographicBufferStatics) CreateFromByteArray(valueLength uint32, value *byte) *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromByteArray, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICryptographicBufferStatics) CopyToByteArray(buffer *IBuffer, valueLength uint32, value *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CopyToByteArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(buffer)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ICryptographicBufferStatics) DecodeFromHexString(value string) *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DecodeFromHexString, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICryptographicBufferStatics) EncodeToHexString(buffer *IBuffer) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EncodeToHexString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICryptographicBufferStatics) DecodeFromBase64String(value string) *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DecodeFromBase64String, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICryptographicBufferStatics) EncodeToBase64String(buffer *IBuffer) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EncodeToBase64String, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICryptographicBufferStatics) ConvertStringToBinary(value string, encoding BinaryStringEncoding) *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConvertStringToBinary, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr, uintptr(encoding), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICryptographicBufferStatics) ConvertBinaryToString(encoding BinaryStringEncoding, buffer *IBuffer) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConvertBinaryToString, uintptr(unsafe.Pointer(this)), uintptr(encoding), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// classes

type CryptographicBuffer struct {
	RtClass
}

func NewICryptographicBufferStatics() *ICryptographicBufferStatics {
	var p *ICryptographicBufferStatics
	hs := NewHStr("Windows.Security.Cryptography.CryptographicBuffer")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICryptographicBufferStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
