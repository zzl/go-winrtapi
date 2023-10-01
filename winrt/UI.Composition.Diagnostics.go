package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
// flags
type CompositionDebugOverdrawContentKinds uint32

const (
	CompositionDebugOverdrawContentKinds_None              CompositionDebugOverdrawContentKinds = 0
	CompositionDebugOverdrawContentKinds_OffscreenRendered CompositionDebugOverdrawContentKinds = 1
	CompositionDebugOverdrawContentKinds_Colors            CompositionDebugOverdrawContentKinds = 2
	CompositionDebugOverdrawContentKinds_Effects           CompositionDebugOverdrawContentKinds = 4
	CompositionDebugOverdrawContentKinds_Shadows           CompositionDebugOverdrawContentKinds = 8
	CompositionDebugOverdrawContentKinds_Lights            CompositionDebugOverdrawContentKinds = 16
	CompositionDebugOverdrawContentKinds_Surfaces          CompositionDebugOverdrawContentKinds = 32
	CompositionDebugOverdrawContentKinds_SwapChains        CompositionDebugOverdrawContentKinds = 64
	CompositionDebugOverdrawContentKinds_All               CompositionDebugOverdrawContentKinds = 4294967295
)

// interfaces

// E49C90AC-2FF3-5805-718C-B725EE07650F
var IID_ICompositionDebugHeatMaps = syscall.GUID{0xE49C90AC, 0x2FF3, 0x5805,
	[8]byte{0x71, 0x8C, 0xB7, 0x25, 0xEE, 0x07, 0x65, 0x0F}}

type ICompositionDebugHeatMapsInterface interface {
	win32.IInspectableInterface
	Hide(subtree *IVisual)
	ShowMemoryUsage(subtree *IVisual)
	ShowOverdraw(subtree *IVisual, contentKinds CompositionDebugOverdrawContentKinds)
	ShowRedraw(subtree *IVisual)
}

type ICompositionDebugHeatMapsVtbl struct {
	win32.IInspectableVtbl
	Hide            uintptr
	ShowMemoryUsage uintptr
	ShowOverdraw    uintptr
	ShowRedraw      uintptr
}

type ICompositionDebugHeatMaps struct {
	win32.IInspectable
}

func (this *ICompositionDebugHeatMaps) Vtbl() *ICompositionDebugHeatMapsVtbl {
	return (*ICompositionDebugHeatMapsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionDebugHeatMaps) Hide(subtree *IVisual) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Hide, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(subtree)))
	_ = _hr
}

func (this *ICompositionDebugHeatMaps) ShowMemoryUsage(subtree *IVisual) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowMemoryUsage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(subtree)))
	_ = _hr
}

func (this *ICompositionDebugHeatMaps) ShowOverdraw(subtree *IVisual, contentKinds CompositionDebugOverdrawContentKinds) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowOverdraw, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(subtree)), uintptr(contentKinds))
	_ = _hr
}

func (this *ICompositionDebugHeatMaps) ShowRedraw(subtree *IVisual) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowRedraw, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(subtree)))
	_ = _hr
}

// 2831987E-1D82-4D38-B7B7-EFD11C7BC3D1
var IID_ICompositionDebugSettings = syscall.GUID{0x2831987E, 0x1D82, 0x4D38,
	[8]byte{0xB7, 0xB7, 0xEF, 0xD1, 0x1C, 0x7B, 0xC3, 0xD1}}

type ICompositionDebugSettingsInterface interface {
	win32.IInspectableInterface
	Get_HeatMaps() *ICompositionDebugHeatMaps
}

type ICompositionDebugSettingsVtbl struct {
	win32.IInspectableVtbl
	Get_HeatMaps uintptr
}

type ICompositionDebugSettings struct {
	win32.IInspectable
}

func (this *ICompositionDebugSettings) Vtbl() *ICompositionDebugSettingsVtbl {
	return (*ICompositionDebugSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionDebugSettings) Get_HeatMaps() *ICompositionDebugHeatMaps {
	var _result *ICompositionDebugHeatMaps
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HeatMaps, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 64EC1F1E-6AF8-4AF8-B814-C870FD5A9505
var IID_ICompositionDebugSettingsStatics = syscall.GUID{0x64EC1F1E, 0x6AF8, 0x4AF8,
	[8]byte{0xB8, 0x14, 0xC8, 0x70, 0xFD, 0x5A, 0x95, 0x05}}

type ICompositionDebugSettingsStaticsInterface interface {
	win32.IInspectableInterface
	TryGetSettings(compositor *ICompositor) *ICompositionDebugSettings
}

type ICompositionDebugSettingsStaticsVtbl struct {
	win32.IInspectableVtbl
	TryGetSettings uintptr
}

type ICompositionDebugSettingsStatics struct {
	win32.IInspectable
}

func (this *ICompositionDebugSettingsStatics) Vtbl() *ICompositionDebugSettingsStaticsVtbl {
	return (*ICompositionDebugSettingsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionDebugSettingsStatics) TryGetSettings(compositor *ICompositor) *ICompositionDebugSettings {
	var _result *ICompositionDebugSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetSettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(compositor)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type CompositionDebugHeatMaps struct {
	RtClass
	*ICompositionDebugHeatMaps
}

type CompositionDebugSettings struct {
	RtClass
	*ICompositionDebugSettings
}

func NewICompositionDebugSettingsStatics() *ICompositionDebugSettingsStatics {
	var p *ICompositionDebugSettingsStatics
	hs := NewHStr("Windows.UI.Composition.Diagnostics.CompositionDebugSettings")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICompositionDebugSettingsStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
