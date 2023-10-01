package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// interfaces

// 3D577EF6-4CEE-11E6-B535-001BDC06AB3B
var IID_IRadialControllerIndependentInputSource = syscall.GUID{0x3D577EF6, 0x4CEE, 0x11E6,
	[8]byte{0xB5, 0x35, 0x00, 0x1B, 0xDC, 0x06, 0xAB, 0x3B}}

type IRadialControllerIndependentInputSourceInterface interface {
	win32.IInspectableInterface
	Get_Controller() *IRadialController
	Get_Dispatcher() unsafe.Pointer
}

type IRadialControllerIndependentInputSourceVtbl struct {
	win32.IInspectableVtbl
	Get_Controller uintptr
	Get_Dispatcher uintptr
}

type IRadialControllerIndependentInputSource struct {
	win32.IInspectable
}

func (this *IRadialControllerIndependentInputSource) Vtbl() *IRadialControllerIndependentInputSourceVtbl {
	return (*IRadialControllerIndependentInputSourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialControllerIndependentInputSource) Get_Controller() *IRadialController {
	var _result *IRadialController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Controller, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRadialControllerIndependentInputSource) Get_Dispatcher() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Dispatcher, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7073AAD8-35F3-4EEB-8751-BE4D0A66FAF4
var IID_IRadialControllerIndependentInputSource2 = syscall.GUID{0x7073AAD8, 0x35F3, 0x4EEB,
	[8]byte{0x87, 0x51, 0xBE, 0x4D, 0x0A, 0x66, 0xFA, 0xF4}}

type IRadialControllerIndependentInputSource2Interface interface {
	win32.IInspectableInterface
	Get_DispatcherQueue() *IDispatcherQueue
}

type IRadialControllerIndependentInputSource2Vtbl struct {
	win32.IInspectableVtbl
	Get_DispatcherQueue uintptr
}

type IRadialControllerIndependentInputSource2 struct {
	win32.IInspectable
}

func (this *IRadialControllerIndependentInputSource2) Vtbl() *IRadialControllerIndependentInputSource2Vtbl {
	return (*IRadialControllerIndependentInputSource2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialControllerIndependentInputSource2) Get_DispatcherQueue() *IDispatcherQueue {
	var _result *IDispatcherQueue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DispatcherQueue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3D577EF5-4CEE-11E6-B535-001BDC06AB3B
var IID_IRadialControllerIndependentInputSourceStatics = syscall.GUID{0x3D577EF5, 0x4CEE, 0x11E6,
	[8]byte{0xB5, 0x35, 0x00, 0x1B, 0xDC, 0x06, 0xAB, 0x3B}}

type IRadialControllerIndependentInputSourceStaticsInterface interface {
	win32.IInspectableInterface
	CreateForView(view *ICoreApplicationView) *IRadialControllerIndependentInputSource
}

type IRadialControllerIndependentInputSourceStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateForView uintptr
}

type IRadialControllerIndependentInputSourceStatics struct {
	win32.IInspectable
}

func (this *IRadialControllerIndependentInputSourceStatics) Vtbl() *IRadialControllerIndependentInputSourceStaticsVtbl {
	return (*IRadialControllerIndependentInputSourceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRadialControllerIndependentInputSourceStatics) CreateForView(view *ICoreApplicationView) *IRadialControllerIndependentInputSource {
	var _result *IRadialControllerIndependentInputSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateForView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(view)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type RadialControllerIndependentInputSource struct {
	RtClass
	*IRadialControllerIndependentInputSource
}

func NewIRadialControllerIndependentInputSourceStatics() *IRadialControllerIndependentInputSourceStatics {
	var p *IRadialControllerIndependentInputSourceStatics
	hs := NewHStr("Windows.UI.Input.Core.RadialControllerIndependentInputSource")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IRadialControllerIndependentInputSourceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
