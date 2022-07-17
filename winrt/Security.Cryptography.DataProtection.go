package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"log"
	"syscall"
	"unsafe"
)

// interfaces

// 09639948-ED22-4270-BD1C-6D72C00F8787
var IID_IDataProtectionProvider = syscall.GUID{0x09639948, 0xED22, 0x4270,
	[8]byte{0xBD, 0x1C, 0x6D, 0x72, 0xC0, 0x0F, 0x87, 0x87}}

type IDataProtectionProviderInterface interface {
	win32.IInspectableInterface
	ProtectAsync(data *IBuffer) *IAsyncOperation[*IBuffer]
	UnprotectAsync(data *IBuffer) *IAsyncOperation[*IBuffer]
	ProtectStreamAsync(src *IInputStream, dest *IOutputStream) *IAsyncAction
	UnprotectStreamAsync(src *IInputStream, dest *IOutputStream) *IAsyncAction
}

type IDataProtectionProviderVtbl struct {
	win32.IInspectableVtbl
	ProtectAsync         uintptr
	UnprotectAsync       uintptr
	ProtectStreamAsync   uintptr
	UnprotectStreamAsync uintptr
}

type IDataProtectionProvider struct {
	win32.IInspectable
}

func (this *IDataProtectionProvider) Vtbl() *IDataProtectionProviderVtbl {
	return (*IDataProtectionProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataProtectionProvider) ProtectAsync(data *IBuffer) *IAsyncOperation[*IBuffer] {
	var _result *IAsyncOperation[*IBuffer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ProtectAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataProtectionProvider) UnprotectAsync(data *IBuffer) *IAsyncOperation[*IBuffer] {
	var _result *IAsyncOperation[*IBuffer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UnprotectAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataProtectionProvider) ProtectStreamAsync(src *IInputStream, dest *IOutputStream) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ProtectStreamAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(dest)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataProtectionProvider) UnprotectStreamAsync(src *IInputStream, dest *IOutputStream) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UnprotectStreamAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(dest)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// ADF33DAC-4932-4CDF-AC41-7214333514CA
var IID_IDataProtectionProviderFactory = syscall.GUID{0xADF33DAC, 0x4932, 0x4CDF,
	[8]byte{0xAC, 0x41, 0x72, 0x14, 0x33, 0x35, 0x14, 0xCA}}

type IDataProtectionProviderFactoryInterface interface {
	win32.IInspectableInterface
	CreateOverloadExplicit(protectionDescriptor string) *IDataProtectionProvider
}

type IDataProtectionProviderFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateOverloadExplicit uintptr
}

type IDataProtectionProviderFactory struct {
	win32.IInspectable
}

func (this *IDataProtectionProviderFactory) Vtbl() *IDataProtectionProviderFactoryVtbl {
	return (*IDataProtectionProviderFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataProtectionProviderFactory) CreateOverloadExplicit(protectionDescriptor string) *IDataProtectionProvider {
	var _result *IDataProtectionProvider
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateOverloadExplicit, uintptr(unsafe.Pointer(this)), NewHStr(protectionDescriptor).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type DataProtectionProvider struct {
	RtClass
	*IDataProtectionProvider
}

func NewDataProtectionProvider() *DataProtectionProvider {
	hs := NewHStr("Windows.Security.Cryptography.DataProtection.DataProtectionProvider")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &DataProtectionProvider{
		RtClass:                 RtClass{PInspect: p},
		IDataProtectionProvider: (*IDataProtectionProvider)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewDataProtectionProvider_CreateOverloadExplicit(protectionDescriptor string) *DataProtectionProvider {
	hs := NewHStr("Windows.Security.Cryptography.DataProtection.DataProtectionProvider")
	var pFac *IDataProtectionProviderFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IDataProtectionProviderFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IDataProtectionProvider
	p = pFac.CreateOverloadExplicit(protectionDescriptor)
	result := &DataProtectionProvider{
		RtClass:                 RtClass{PInspect: &p.IInspectable},
		IDataProtectionProvider: p,
	}
	com.AddToScope(result)
	return result
}
