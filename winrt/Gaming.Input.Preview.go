package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// interfaces

// 0BE1E6C5-D9BD-44EE-8362-488B2E464BFB
var IID_IGameControllerProviderInfoStatics = syscall.GUID{0x0BE1E6C5, 0xD9BD, 0x44EE,
	[8]byte{0x83, 0x62, 0x48, 0x8B, 0x2E, 0x46, 0x4B, 0xFB}}

type IGameControllerProviderInfoStaticsInterface interface {
	win32.IInspectableInterface
	GetParentProviderId(provider *IGameControllerProvider) string
	GetProviderId(provider *IGameControllerProvider) string
}

type IGameControllerProviderInfoStaticsVtbl struct {
	win32.IInspectableVtbl
	GetParentProviderId uintptr
	GetProviderId       uintptr
}

type IGameControllerProviderInfoStatics struct {
	win32.IInspectable
}

func (this *IGameControllerProviderInfoStatics) Vtbl() *IGameControllerProviderInfoStaticsVtbl {
	return (*IGameControllerProviderInfoStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameControllerProviderInfoStatics) GetParentProviderId(provider *IGameControllerProvider) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetParentProviderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(provider)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IGameControllerProviderInfoStatics) GetProviderId(provider *IGameControllerProvider) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetProviderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(provider)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// classes

type GameControllerProviderInfo struct {
	RtClass
}

func NewIGameControllerProviderInfoStatics() *IGameControllerProviderInfoStatics {
	var p *IGameControllerProviderInfoStatics
	hs := NewHStr("Windows.Gaming.Input.Preview.GameControllerProviderInfo")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGameControllerProviderInfoStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
