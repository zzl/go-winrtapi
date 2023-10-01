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
type AnimationEffect int32

const (
	AnimationEffect_Expand               AnimationEffect = 0
	AnimationEffect_Collapse             AnimationEffect = 1
	AnimationEffect_Reposition           AnimationEffect = 2
	AnimationEffect_FadeIn               AnimationEffect = 3
	AnimationEffect_FadeOut              AnimationEffect = 4
	AnimationEffect_AddToList            AnimationEffect = 5
	AnimationEffect_DeleteFromList       AnimationEffect = 6
	AnimationEffect_AddToGrid            AnimationEffect = 7
	AnimationEffect_DeleteFromGrid       AnimationEffect = 8
	AnimationEffect_AddToSearchGrid      AnimationEffect = 9
	AnimationEffect_DeleteFromSearchGrid AnimationEffect = 10
	AnimationEffect_AddToSearchList      AnimationEffect = 11
	AnimationEffect_DeleteFromSearchList AnimationEffect = 12
	AnimationEffect_ShowEdgeUI           AnimationEffect = 13
	AnimationEffect_ShowPanel            AnimationEffect = 14
	AnimationEffect_HideEdgeUI           AnimationEffect = 15
	AnimationEffect_HidePanel            AnimationEffect = 16
	AnimationEffect_ShowPopup            AnimationEffect = 17
	AnimationEffect_HidePopup            AnimationEffect = 18
	AnimationEffect_PointerDown          AnimationEffect = 19
	AnimationEffect_PointerUp            AnimationEffect = 20
	AnimationEffect_DragSourceStart      AnimationEffect = 21
	AnimationEffect_DragSourceEnd        AnimationEffect = 22
	AnimationEffect_TransitionContent    AnimationEffect = 23
	AnimationEffect_Reveal               AnimationEffect = 24
	AnimationEffect_Hide                 AnimationEffect = 25
	AnimationEffect_DragBetweenEnter     AnimationEffect = 26
	AnimationEffect_DragBetweenLeave     AnimationEffect = 27
	AnimationEffect_SwipeSelect          AnimationEffect = 28
	AnimationEffect_SwipeDeselect        AnimationEffect = 29
	AnimationEffect_SwipeReveal          AnimationEffect = 30
	AnimationEffect_EnterPage            AnimationEffect = 31
	AnimationEffect_TransitionPage       AnimationEffect = 32
	AnimationEffect_CrossFade            AnimationEffect = 33
	AnimationEffect_Peek                 AnimationEffect = 34
	AnimationEffect_UpdateBadge          AnimationEffect = 35
)

// enum
type AnimationEffectTarget int32

const (
	AnimationEffectTarget_Primary    AnimationEffectTarget = 0
	AnimationEffectTarget_Added      AnimationEffectTarget = 1
	AnimationEffectTarget_Affected   AnimationEffectTarget = 2
	AnimationEffectTarget_Background AnimationEffectTarget = 3
	AnimationEffectTarget_Content    AnimationEffectTarget = 4
	AnimationEffectTarget_Deleted    AnimationEffectTarget = 5
	AnimationEffectTarget_Deselected AnimationEffectTarget = 6
	AnimationEffectTarget_DragSource AnimationEffectTarget = 7
	AnimationEffectTarget_Hidden     AnimationEffectTarget = 8
	AnimationEffectTarget_Incoming   AnimationEffectTarget = 9
	AnimationEffectTarget_Outgoing   AnimationEffectTarget = 10
	AnimationEffectTarget_Outline    AnimationEffectTarget = 11
	AnimationEffectTarget_Remaining  AnimationEffectTarget = 12
	AnimationEffectTarget_Revealed   AnimationEffectTarget = 13
	AnimationEffectTarget_RowIn      AnimationEffectTarget = 14
	AnimationEffectTarget_RowOut     AnimationEffectTarget = 15
	AnimationEffectTarget_Selected   AnimationEffectTarget = 16
	AnimationEffectTarget_Selection  AnimationEffectTarget = 17
	AnimationEffectTarget_Shown      AnimationEffectTarget = 18
	AnimationEffectTarget_Tapped     AnimationEffectTarget = 19
)

// enum
type PropertyAnimationType int32

const (
	PropertyAnimationType_Scale       PropertyAnimationType = 0
	PropertyAnimationType_Translation PropertyAnimationType = 1
	PropertyAnimationType_Opacity     PropertyAnimationType = 2
)

// structs

type AnimationMetricsContract struct {
}

// interfaces

// 7D11A549-BE3D-41DE-B081-05C149962F9B
var IID_IAnimationDescription = syscall.GUID{0x7D11A549, 0xBE3D, 0x41DE,
	[8]byte{0xB0, 0x81, 0x05, 0xC1, 0x49, 0x96, 0x2F, 0x9B}}

type IAnimationDescriptionInterface interface {
	win32.IInspectableInterface
	Get_Animations() *IVectorView[*IPropertyAnimation]
	Get_StaggerDelay() TimeSpan
	Get_StaggerDelayFactor() float32
	Get_DelayLimit() TimeSpan
	Get_ZOrder() int32
}

type IAnimationDescriptionVtbl struct {
	win32.IInspectableVtbl
	Get_Animations         uintptr
	Get_StaggerDelay       uintptr
	Get_StaggerDelayFactor uintptr
	Get_DelayLimit         uintptr
	Get_ZOrder             uintptr
}

type IAnimationDescription struct {
	win32.IInspectable
}

func (this *IAnimationDescription) Vtbl() *IAnimationDescriptionVtbl {
	return (*IAnimationDescriptionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAnimationDescription) Get_Animations() *IVectorView[*IPropertyAnimation] {
	var _result *IVectorView[*IPropertyAnimation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Animations, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAnimationDescription) Get_StaggerDelay() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StaggerDelay, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAnimationDescription) Get_StaggerDelayFactor() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StaggerDelayFactor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAnimationDescription) Get_DelayLimit() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DelayLimit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAnimationDescription) Get_ZOrder() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ZOrder, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C6E27ABE-C1FB-48B5-9271-ECC70AC86EF0
var IID_IAnimationDescriptionFactory = syscall.GUID{0xC6E27ABE, 0xC1FB, 0x48B5,
	[8]byte{0x92, 0x71, 0xEC, 0xC7, 0x0A, 0xC8, 0x6E, 0xF0}}

type IAnimationDescriptionFactoryInterface interface {
	win32.IInspectableInterface
	CreateInstance(effect AnimationEffect, target AnimationEffectTarget) *IAnimationDescription
}

type IAnimationDescriptionFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateInstance uintptr
}

type IAnimationDescriptionFactory struct {
	win32.IInspectable
}

func (this *IAnimationDescriptionFactory) Vtbl() *IAnimationDescriptionFactoryVtbl {
	return (*IAnimationDescriptionFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAnimationDescriptionFactory) CreateInstance(effect AnimationEffect, target AnimationEffectTarget) *IAnimationDescription {
	var _result *IAnimationDescription
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateInstance, uintptr(unsafe.Pointer(this)), uintptr(effect), uintptr(target), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 803AABE5-EE7E-455F-84E9-2506AFB8D2B4
var IID_IOpacityAnimation = syscall.GUID{0x803AABE5, 0xEE7E, 0x455F,
	[8]byte{0x84, 0xE9, 0x25, 0x06, 0xAF, 0xB8, 0xD2, 0xB4}}

type IOpacityAnimationInterface interface {
	win32.IInspectableInterface
	Get_InitialOpacity() *IReference[float32]
	Get_FinalOpacity() float32
}

type IOpacityAnimationVtbl struct {
	win32.IInspectableVtbl
	Get_InitialOpacity uintptr
	Get_FinalOpacity   uintptr
}

type IOpacityAnimation struct {
	win32.IInspectable
}

func (this *IOpacityAnimation) Vtbl() *IOpacityAnimationVtbl {
	return (*IOpacityAnimationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IOpacityAnimation) Get_InitialOpacity() *IReference[float32] {
	var _result *IReference[float32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InitialOpacity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IOpacityAnimation) Get_FinalOpacity() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FinalOpacity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 3A01B4DA-4D8C-411E-B615-1ADE683A9903
var IID_IPropertyAnimation = syscall.GUID{0x3A01B4DA, 0x4D8C, 0x411E,
	[8]byte{0xB6, 0x15, 0x1A, 0xDE, 0x68, 0x3A, 0x99, 0x03}}

type IPropertyAnimationInterface interface {
	win32.IInspectableInterface
	Get_Type() PropertyAnimationType
	Get_Delay() TimeSpan
	Get_Duration() TimeSpan
	Get_Control1() Point
	Get_Control2() Point
}

type IPropertyAnimationVtbl struct {
	win32.IInspectableVtbl
	Get_Type     uintptr
	Get_Delay    uintptr
	Get_Duration uintptr
	Get_Control1 uintptr
	Get_Control2 uintptr
}

type IPropertyAnimation struct {
	win32.IInspectable
}

func (this *IPropertyAnimation) Vtbl() *IPropertyAnimationVtbl {
	return (*IPropertyAnimationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertyAnimation) Get_Type() PropertyAnimationType {
	var _result PropertyAnimationType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Type, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyAnimation) Get_Delay() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Delay, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyAnimation) Get_Duration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Duration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyAnimation) Get_Control1() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Control1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPropertyAnimation) Get_Control2() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Control2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 023552C7-71AB-428C-9C9F-D31780964995
var IID_IScaleAnimation = syscall.GUID{0x023552C7, 0x71AB, 0x428C,
	[8]byte{0x9C, 0x9F, 0xD3, 0x17, 0x80, 0x96, 0x49, 0x95}}

type IScaleAnimationInterface interface {
	win32.IInspectableInterface
	Get_InitialScaleX() *IReference[float32]
	Get_InitialScaleY() *IReference[float32]
	Get_FinalScaleX() float32
	Get_FinalScaleY() float32
	Get_NormalizedOrigin() Point
}

type IScaleAnimationVtbl struct {
	win32.IInspectableVtbl
	Get_InitialScaleX    uintptr
	Get_InitialScaleY    uintptr
	Get_FinalScaleX      uintptr
	Get_FinalScaleY      uintptr
	Get_NormalizedOrigin uintptr
}

type IScaleAnimation struct {
	win32.IInspectable
}

func (this *IScaleAnimation) Vtbl() *IScaleAnimationVtbl {
	return (*IScaleAnimationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IScaleAnimation) Get_InitialScaleX() *IReference[float32] {
	var _result *IReference[float32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InitialScaleX, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IScaleAnimation) Get_InitialScaleY() *IReference[float32] {
	var _result *IReference[float32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InitialScaleY, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IScaleAnimation) Get_FinalScaleX() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FinalScaleX, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IScaleAnimation) Get_FinalScaleY() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FinalScaleY, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IScaleAnimation) Get_NormalizedOrigin() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NormalizedOrigin, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type AnimationDescription struct {
	RtClass
	*IAnimationDescription
}

func NewAnimationDescription_CreateInstance(effect AnimationEffect, target AnimationEffectTarget) *AnimationDescription {
	hs := NewHStr("Windows.UI.Core.AnimationMetrics.AnimationDescription")
	var pFac *IAnimationDescriptionFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAnimationDescriptionFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IAnimationDescription
	p = pFac.CreateInstance(effect, target)
	result := &AnimationDescription{
		RtClass:               RtClass{PInspect: &p.IInspectable},
		IAnimationDescription: p,
	}
	com.AddToScope(result)
	return result
}

type OpacityAnimation struct {
	RtClass
	*IOpacityAnimation
}

type PropertyAnimation struct {
	RtClass
	*IPropertyAnimation
}

type ScaleAnimation struct {
	RtClass
	*IScaleAnimation
}

type TranslationAnimation struct {
	RtClass
	*IPropertyAnimation
}
