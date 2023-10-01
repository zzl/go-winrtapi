package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// structs

type PerceptionAutomationCoreContract struct {
}

// interfaces

// 0BB04541-4CE2-4923-9A76-8187ECC59112
var IID_ICorePerceptionAutomationStatics = syscall.GUID{0x0BB04541, 0x4CE2, 0x4923,
	[8]byte{0x9A, 0x76, 0x81, 0x87, 0xEC, 0xC5, 0x91, 0x12}}

type ICorePerceptionAutomationStaticsInterface interface {
	win32.IInspectableInterface
	SetActivationFactoryProvider(provider *IGetActivationFactory)
}

type ICorePerceptionAutomationStaticsVtbl struct {
	win32.IInspectableVtbl
	SetActivationFactoryProvider uintptr
}

type ICorePerceptionAutomationStatics struct {
	win32.IInspectable
}

func (this *ICorePerceptionAutomationStatics) Vtbl() *ICorePerceptionAutomationStaticsVtbl {
	return (*ICorePerceptionAutomationStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICorePerceptionAutomationStatics) SetActivationFactoryProvider(provider *IGetActivationFactory) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetActivationFactoryProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(provider)))
	_ = _hr
}

// classes

type CorePerceptionAutomation struct {
	RtClass
}

func NewICorePerceptionAutomationStatics() *ICorePerceptionAutomationStatics {
	var p *ICorePerceptionAutomationStatics
	hs := NewHStr("Windows.Perception.Automation.Core.CorePerceptionAutomation")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICorePerceptionAutomationStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
