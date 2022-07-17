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
type CompressAlgorithm int32

const (
	CompressAlgorithm_InvalidAlgorithm CompressAlgorithm = 0
	CompressAlgorithm_NullAlgorithm    CompressAlgorithm = 1
	CompressAlgorithm_Mszip            CompressAlgorithm = 2
	CompressAlgorithm_Xpress           CompressAlgorithm = 3
	CompressAlgorithm_XpressHuff       CompressAlgorithm = 4
	CompressAlgorithm_Lzms             CompressAlgorithm = 5
)

// interfaces

// 0AC3645A-57AC-4EE1-B702-84D39D5424E0
var IID_ICompressor = syscall.GUID{0x0AC3645A, 0x57AC, 0x4EE1,
	[8]byte{0xB7, 0x02, 0x84, 0xD3, 0x9D, 0x54, 0x24, 0xE0}}

type ICompressorInterface interface {
	win32.IInspectableInterface
	FinishAsync() *IAsyncOperation[bool]
	DetachStream() *IOutputStream
}

type ICompressorVtbl struct {
	win32.IInspectableVtbl
	FinishAsync  uintptr
	DetachStream uintptr
}

type ICompressor struct {
	win32.IInspectable
}

func (this *ICompressor) Vtbl() *ICompressorVtbl {
	return (*ICompressorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompressor) FinishAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FinishAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompressor) DetachStream() *IOutputStream {
	var _result *IOutputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DetachStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5F3D96A4-2CFB-442C-A8BA-D7D11B039DA0
var IID_ICompressorFactory = syscall.GUID{0x5F3D96A4, 0x2CFB, 0x442C,
	[8]byte{0xA8, 0xBA, 0xD7, 0xD1, 0x1B, 0x03, 0x9D, 0xA0}}

type ICompressorFactoryInterface interface {
	win32.IInspectableInterface
	CreateCompressor(underlyingStream *IOutputStream) *ICompressor
	CreateCompressorEx(underlyingStream *IOutputStream, algorithm CompressAlgorithm, blockSize uint32) *ICompressor
}

type ICompressorFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateCompressor   uintptr
	CreateCompressorEx uintptr
}

type ICompressorFactory struct {
	win32.IInspectable
}

func (this *ICompressorFactory) Vtbl() *ICompressorFactoryVtbl {
	return (*ICompressorFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompressorFactory) CreateCompressor(underlyingStream *IOutputStream) *ICompressor {
	var _result *ICompressor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateCompressor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(underlyingStream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompressorFactory) CreateCompressorEx(underlyingStream *IOutputStream, algorithm CompressAlgorithm, blockSize uint32) *ICompressor {
	var _result *ICompressor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateCompressorEx, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(underlyingStream)), uintptr(algorithm), uintptr(blockSize), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B883FE46-D68A-4C8B-ADA0-4EE813FC5283
var IID_IDecompressor = syscall.GUID{0xB883FE46, 0xD68A, 0x4C8B,
	[8]byte{0xAD, 0xA0, 0x4E, 0xE8, 0x13, 0xFC, 0x52, 0x83}}

type IDecompressorInterface interface {
	win32.IInspectableInterface
	DetachStream() *IInputStream
}

type IDecompressorVtbl struct {
	win32.IInspectableVtbl
	DetachStream uintptr
}

type IDecompressor struct {
	win32.IInspectable
}

func (this *IDecompressor) Vtbl() *IDecompressorVtbl {
	return (*IDecompressorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDecompressor) DetachStream() *IInputStream {
	var _result *IInputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DetachStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5337E252-1DA2-42E1-8834-0379D28D742F
var IID_IDecompressorFactory = syscall.GUID{0x5337E252, 0x1DA2, 0x42E1,
	[8]byte{0x88, 0x34, 0x03, 0x79, 0xD2, 0x8D, 0x74, 0x2F}}

type IDecompressorFactoryInterface interface {
	win32.IInspectableInterface
	CreateDecompressor(underlyingStream *IInputStream) *IDecompressor
}

type IDecompressorFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateDecompressor uintptr
}

type IDecompressorFactory struct {
	win32.IInspectable
}

func (this *IDecompressorFactory) Vtbl() *IDecompressorFactoryVtbl {
	return (*IDecompressorFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDecompressorFactory) CreateDecompressor(underlyingStream *IInputStream) *IDecompressor {
	var _result *IDecompressor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDecompressor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(underlyingStream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type Compressor struct {
	RtClass
	*ICompressor
}

func NewCompressor_CreateCompressor(underlyingStream *IOutputStream) *Compressor {
	hs := NewHStr("Windows.Storage.Compression.Compressor")
	var pFac *ICompressorFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICompressorFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ICompressor
	p = pFac.CreateCompressor(underlyingStream)
	result := &Compressor{
		RtClass:     RtClass{PInspect: &p.IInspectable},
		ICompressor: p,
	}
	com.AddToScope(result)
	return result
}

func NewCompressor_CreateCompressorEx(underlyingStream *IOutputStream, algorithm CompressAlgorithm, blockSize uint32) *Compressor {
	hs := NewHStr("Windows.Storage.Compression.Compressor")
	var pFac *ICompressorFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICompressorFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ICompressor
	p = pFac.CreateCompressorEx(underlyingStream, algorithm, blockSize)
	result := &Compressor{
		RtClass:     RtClass{PInspect: &p.IInspectable},
		ICompressor: p,
	}
	com.AddToScope(result)
	return result
}

type Decompressor struct {
	RtClass
	*IDecompressor
}

func NewDecompressor_CreateDecompressor(underlyingStream *IInputStream) *Decompressor {
	hs := NewHStr("Windows.Storage.Compression.Decompressor")
	var pFac *IDecompressorFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IDecompressorFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IDecompressor
	p = pFac.CreateDecompressor(underlyingStream)
	result := &Decompressor{
		RtClass:       RtClass{PInspect: &p.IInspectable},
		IDecompressor: p,
	}
	com.AddToScope(result)
	return result
}
