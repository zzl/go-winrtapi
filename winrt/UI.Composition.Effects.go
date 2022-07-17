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
type SceneLightingEffectReflectanceModel int32

const (
	SceneLightingEffectReflectanceModel_BlinnPhong                SceneLightingEffectReflectanceModel = 0
	SceneLightingEffectReflectanceModel_PhysicallyBasedBlinnPhong SceneLightingEffectReflectanceModel = 1
)

// interfaces

// 91BB5E52-95D1-4F8B-9A5A-6408B24B8C6A
var IID_ISceneLightingEffect = syscall.GUID{0x91BB5E52, 0x95D1, 0x4F8B,
	[8]byte{0x9A, 0x5A, 0x64, 0x08, 0xB2, 0x4B, 0x8C, 0x6A}}

type ISceneLightingEffectInterface interface {
	win32.IInspectableInterface
	Get_AmbientAmount() float32
	Put_AmbientAmount(value float32)
	Get_DiffuseAmount() float32
	Put_DiffuseAmount(value float32)
	Get_NormalMapSource() unsafe.Pointer
	Put_NormalMapSource(value unsafe.Pointer)
	Get_SpecularAmount() float32
	Put_SpecularAmount(value float32)
	Get_SpecularShine() float32
	Put_SpecularShine(value float32)
}

type ISceneLightingEffectVtbl struct {
	win32.IInspectableVtbl
	Get_AmbientAmount   uintptr
	Put_AmbientAmount   uintptr
	Get_DiffuseAmount   uintptr
	Put_DiffuseAmount   uintptr
	Get_NormalMapSource uintptr
	Put_NormalMapSource uintptr
	Get_SpecularAmount  uintptr
	Put_SpecularAmount  uintptr
	Get_SpecularShine   uintptr
	Put_SpecularShine   uintptr
}

type ISceneLightingEffect struct {
	win32.IInspectable
}

func (this *ISceneLightingEffect) Vtbl() *ISceneLightingEffectVtbl {
	return (*ISceneLightingEffectVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISceneLightingEffect) Get_AmbientAmount() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AmbientAmount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISceneLightingEffect) Put_AmbientAmount(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AmbientAmount, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISceneLightingEffect) Get_DiffuseAmount() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DiffuseAmount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISceneLightingEffect) Put_DiffuseAmount(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DiffuseAmount, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISceneLightingEffect) Get_NormalMapSource() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NormalMapSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISceneLightingEffect) Put_NormalMapSource(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_NormalMapSource, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISceneLightingEffect) Get_SpecularAmount() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SpecularAmount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISceneLightingEffect) Put_SpecularAmount(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SpecularAmount, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISceneLightingEffect) Get_SpecularShine() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SpecularShine, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISceneLightingEffect) Put_SpecularShine(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SpecularShine, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 9E270E81-72F0-4C5C-95F8-8A6E0024F409
var IID_ISceneLightingEffect2 = syscall.GUID{0x9E270E81, 0x72F0, 0x4C5C,
	[8]byte{0x95, 0xF8, 0x8A, 0x6E, 0x00, 0x24, 0xF4, 0x09}}

type ISceneLightingEffect2Interface interface {
	win32.IInspectableInterface
	Get_ReflectanceModel() SceneLightingEffectReflectanceModel
	Put_ReflectanceModel(value SceneLightingEffectReflectanceModel)
}

type ISceneLightingEffect2Vtbl struct {
	win32.IInspectableVtbl
	Get_ReflectanceModel uintptr
	Put_ReflectanceModel uintptr
}

type ISceneLightingEffect2 struct {
	win32.IInspectable
}

func (this *ISceneLightingEffect2) Vtbl() *ISceneLightingEffect2Vtbl {
	return (*ISceneLightingEffect2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISceneLightingEffect2) Get_ReflectanceModel() SceneLightingEffectReflectanceModel {
	var _result SceneLightingEffectReflectanceModel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReflectanceModel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISceneLightingEffect2) Put_ReflectanceModel(value SceneLightingEffectReflectanceModel) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReflectanceModel, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// classes

type SceneLightingEffect struct {
	RtClass
	*ISceneLightingEffect
}

func NewSceneLightingEffect() *SceneLightingEffect {
	hs := NewHStr("Windows.UI.Composition.Effects.SceneLightingEffect")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &SceneLightingEffect{
		RtClass:              RtClass{PInspect: p},
		ISceneLightingEffect: (*ISceneLightingEffect)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}
