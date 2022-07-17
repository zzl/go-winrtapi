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
type AnimationControllerProgressBehavior int32

const (
	AnimationControllerProgressBehavior_Default           AnimationControllerProgressBehavior = 0
	AnimationControllerProgressBehavior_IncludesDelayTime AnimationControllerProgressBehavior = 1
)

// enum
type AnimationDelayBehavior int32

const (
	AnimationDelayBehavior_SetInitialValueAfterDelay  AnimationDelayBehavior = 0
	AnimationDelayBehavior_SetInitialValueBeforeDelay AnimationDelayBehavior = 1
)

// enum
type AnimationDirection int32

const (
	AnimationDirection_Normal           AnimationDirection = 0
	AnimationDirection_Reverse          AnimationDirection = 1
	AnimationDirection_Alternate        AnimationDirection = 2
	AnimationDirection_AlternateReverse AnimationDirection = 3
)

// enum
type AnimationIterationBehavior int32

const (
	AnimationIterationBehavior_Count   AnimationIterationBehavior = 0
	AnimationIterationBehavior_Forever AnimationIterationBehavior = 1
)

// enum
type AnimationPropertyAccessMode int32

const (
	AnimationPropertyAccessMode_None      AnimationPropertyAccessMode = 0
	AnimationPropertyAccessMode_ReadOnly  AnimationPropertyAccessMode = 1
	AnimationPropertyAccessMode_WriteOnly AnimationPropertyAccessMode = 2
	AnimationPropertyAccessMode_ReadWrite AnimationPropertyAccessMode = 3
)

// enum
type AnimationStopBehavior int32

const (
	AnimationStopBehavior_LeaveCurrentValue AnimationStopBehavior = 0
	AnimationStopBehavior_SetToInitialValue AnimationStopBehavior = 1
	AnimationStopBehavior_SetToFinalValue   AnimationStopBehavior = 2
)

// enum
type CompositionBackfaceVisibility int32

const (
	CompositionBackfaceVisibility_Inherit CompositionBackfaceVisibility = 0
	CompositionBackfaceVisibility_Visible CompositionBackfaceVisibility = 1
	CompositionBackfaceVisibility_Hidden  CompositionBackfaceVisibility = 2
)

// enum
// flags
type CompositionBatchTypes uint32

const (
	CompositionBatchTypes_None              CompositionBatchTypes = 0
	CompositionBatchTypes_Animation         CompositionBatchTypes = 1
	CompositionBatchTypes_Effect            CompositionBatchTypes = 2
	CompositionBatchTypes_InfiniteAnimation CompositionBatchTypes = 4
	CompositionBatchTypes_AllAnimations     CompositionBatchTypes = 5
)

// enum
type CompositionBitmapInterpolationMode int32

const (
	CompositionBitmapInterpolationMode_NearestNeighbor                CompositionBitmapInterpolationMode = 0
	CompositionBitmapInterpolationMode_Linear                         CompositionBitmapInterpolationMode = 1
	CompositionBitmapInterpolationMode_MagLinearMinLinearMipLinear    CompositionBitmapInterpolationMode = 2
	CompositionBitmapInterpolationMode_MagLinearMinLinearMipNearest   CompositionBitmapInterpolationMode = 3
	CompositionBitmapInterpolationMode_MagLinearMinNearestMipLinear   CompositionBitmapInterpolationMode = 4
	CompositionBitmapInterpolationMode_MagLinearMinNearestMipNearest  CompositionBitmapInterpolationMode = 5
	CompositionBitmapInterpolationMode_MagNearestMinLinearMipLinear   CompositionBitmapInterpolationMode = 6
	CompositionBitmapInterpolationMode_MagNearestMinLinearMipNearest  CompositionBitmapInterpolationMode = 7
	CompositionBitmapInterpolationMode_MagNearestMinNearestMipLinear  CompositionBitmapInterpolationMode = 8
	CompositionBitmapInterpolationMode_MagNearestMinNearestMipNearest CompositionBitmapInterpolationMode = 9
)

// enum
type CompositionBorderMode int32

const (
	CompositionBorderMode_Inherit CompositionBorderMode = 0
	CompositionBorderMode_Soft    CompositionBorderMode = 1
	CompositionBorderMode_Hard    CompositionBorderMode = 2
)

// enum
type CompositionColorSpace int32

const (
	CompositionColorSpace_Auto      CompositionColorSpace = 0
	CompositionColorSpace_Hsl       CompositionColorSpace = 1
	CompositionColorSpace_Rgb       CompositionColorSpace = 2
	CompositionColorSpace_HslLinear CompositionColorSpace = 3
	CompositionColorSpace_RgbLinear CompositionColorSpace = 4
)

// enum
type CompositionCompositeMode int32

const (
	CompositionCompositeMode_Inherit           CompositionCompositeMode = 0
	CompositionCompositeMode_SourceOver        CompositionCompositeMode = 1
	CompositionCompositeMode_DestinationInvert CompositionCompositeMode = 2
	CompositionCompositeMode_MinBlend          CompositionCompositeMode = 3
)

// enum
type CompositionDropShadowSourcePolicy int32

const (
	CompositionDropShadowSourcePolicy_Default                  CompositionDropShadowSourcePolicy = 0
	CompositionDropShadowSourcePolicy_InheritFromVisualContent CompositionDropShadowSourcePolicy = 1
)

// enum
type CompositionEasingFunctionMode int32

const (
	CompositionEasingFunctionMode_In    CompositionEasingFunctionMode = 0
	CompositionEasingFunctionMode_Out   CompositionEasingFunctionMode = 1
	CompositionEasingFunctionMode_InOut CompositionEasingFunctionMode = 2
)

// enum
type CompositionEffectFactoryLoadStatus int32

const (
	CompositionEffectFactoryLoadStatus_Success          CompositionEffectFactoryLoadStatus = 0
	CompositionEffectFactoryLoadStatus_EffectTooComplex CompositionEffectFactoryLoadStatus = 1
	CompositionEffectFactoryLoadStatus_Pending          CompositionEffectFactoryLoadStatus = 2
	CompositionEffectFactoryLoadStatus_Other            CompositionEffectFactoryLoadStatus = -1
)

// enum
type CompositionGetValueStatus int32

const (
	CompositionGetValueStatus_Succeeded    CompositionGetValueStatus = 0
	CompositionGetValueStatus_TypeMismatch CompositionGetValueStatus = 1
	CompositionGetValueStatus_NotFound     CompositionGetValueStatus = 2
)

// enum
type CompositionGradientExtendMode int32

const (
	CompositionGradientExtendMode_Clamp  CompositionGradientExtendMode = 0
	CompositionGradientExtendMode_Wrap   CompositionGradientExtendMode = 1
	CompositionGradientExtendMode_Mirror CompositionGradientExtendMode = 2
)

// enum
type CompositionMappingMode int32

const (
	CompositionMappingMode_Absolute CompositionMappingMode = 0
	CompositionMappingMode_Relative CompositionMappingMode = 1
)

// enum
type CompositionStretch int32

const (
	CompositionStretch_None          CompositionStretch = 0
	CompositionStretch_Fill          CompositionStretch = 1
	CompositionStretch_Uniform       CompositionStretch = 2
	CompositionStretch_UniformToFill CompositionStretch = 3
)

// enum
type CompositionStrokeCap int32

const (
	CompositionStrokeCap_Flat     CompositionStrokeCap = 0
	CompositionStrokeCap_Square   CompositionStrokeCap = 1
	CompositionStrokeCap_Round    CompositionStrokeCap = 2
	CompositionStrokeCap_Triangle CompositionStrokeCap = 3
)

// enum
type CompositionStrokeLineJoin int32

const (
	CompositionStrokeLineJoin_Miter        CompositionStrokeLineJoin = 0
	CompositionStrokeLineJoin_Bevel        CompositionStrokeLineJoin = 1
	CompositionStrokeLineJoin_Round        CompositionStrokeLineJoin = 2
	CompositionStrokeLineJoin_MiterOrBevel CompositionStrokeLineJoin = 3
)

// structs

type InkTrailPoint struct {
	Point  Point
	Radius float32
}

// interfaces

// A48130A1-B7C4-46F7-B9BF-DAF43A44E6EE
var IID_IAmbientLight = syscall.GUID{0xA48130A1, 0xB7C4, 0x46F7,
	[8]byte{0xB9, 0xBF, 0xDA, 0xF4, 0x3A, 0x44, 0xE6, 0xEE}}

type IAmbientLightInterface interface {
	win32.IInspectableInterface
	Get_Color() unsafe.Pointer
	Put_Color(value unsafe.Pointer)
}

type IAmbientLightVtbl struct {
	win32.IInspectableVtbl
	Get_Color uintptr
	Put_Color uintptr
}

type IAmbientLight struct {
	win32.IInspectable
}

func (this *IAmbientLight) Vtbl() *IAmbientLightVtbl {
	return (*IAmbientLightVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAmbientLight) Get_Color() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Color, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAmbientLight) Put_Color(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Color, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 3B64A6BF-5F97-4C94-86E5-042DD386B27D
var IID_IAmbientLight2 = syscall.GUID{0x3B64A6BF, 0x5F97, 0x4C94,
	[8]byte{0x86, 0xE5, 0x04, 0x2D, 0xD3, 0x86, 0xB2, 0x7D}}

type IAmbientLight2Interface interface {
	win32.IInspectableInterface
	Get_Intensity() float32
	Put_Intensity(value float32)
}

type IAmbientLight2Vtbl struct {
	win32.IInspectableVtbl
	Get_Intensity uintptr
	Put_Intensity uintptr
}

type IAmbientLight2 struct {
	win32.IInspectable
}

func (this *IAmbientLight2) Vtbl() *IAmbientLight2Vtbl {
	return (*IAmbientLight2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAmbientLight2) Get_Intensity() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Intensity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAmbientLight2) Put_Intensity(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Intensity, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// C934EFD2-0722-4F5F-A4E2-9510F3D43BF7
var IID_IAnimationController = syscall.GUID{0xC934EFD2, 0x0722, 0x4F5F,
	[8]byte{0xA4, 0xE2, 0x95, 0x10, 0xF3, 0xD4, 0x3B, 0xF7}}

type IAnimationControllerInterface interface {
	win32.IInspectableInterface
	Get_PlaybackRate() float32
	Put_PlaybackRate(value float32)
	Get_Progress() float32
	Put_Progress(value float32)
	Get_ProgressBehavior() AnimationControllerProgressBehavior
	Put_ProgressBehavior(value AnimationControllerProgressBehavior)
	Pause()
	Resume()
}

type IAnimationControllerVtbl struct {
	win32.IInspectableVtbl
	Get_PlaybackRate     uintptr
	Put_PlaybackRate     uintptr
	Get_Progress         uintptr
	Put_Progress         uintptr
	Get_ProgressBehavior uintptr
	Put_ProgressBehavior uintptr
	Pause                uintptr
	Resume               uintptr
}

type IAnimationController struct {
	win32.IInspectable
}

func (this *IAnimationController) Vtbl() *IAnimationControllerVtbl {
	return (*IAnimationControllerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAnimationController) Get_PlaybackRate() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PlaybackRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAnimationController) Put_PlaybackRate(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PlaybackRate, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAnimationController) Get_Progress() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Progress, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAnimationController) Put_Progress(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Progress, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAnimationController) Get_ProgressBehavior() AnimationControllerProgressBehavior {
	var _result AnimationControllerProgressBehavior
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProgressBehavior, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAnimationController) Put_ProgressBehavior(value AnimationControllerProgressBehavior) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ProgressBehavior, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAnimationController) Pause() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Pause, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IAnimationController) Resume() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Resume, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// E71164DF-651B-4800-B9E5-6A3BCFED3365
var IID_IAnimationControllerStatics = syscall.GUID{0xE71164DF, 0x651B, 0x4800,
	[8]byte{0xB9, 0xE5, 0x6A, 0x3B, 0xCF, 0xED, 0x33, 0x65}}

type IAnimationControllerStaticsInterface interface {
	win32.IInspectableInterface
	Get_MaxPlaybackRate() float32
	Get_MinPlaybackRate() float32
}

type IAnimationControllerStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_MaxPlaybackRate uintptr
	Get_MinPlaybackRate uintptr
}

type IAnimationControllerStatics struct {
	win32.IInspectable
}

func (this *IAnimationControllerStatics) Vtbl() *IAnimationControllerStaticsVtbl {
	return (*IAnimationControllerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAnimationControllerStatics) Get_MaxPlaybackRate() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxPlaybackRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAnimationControllerStatics) Get_MinPlaybackRate() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinPlaybackRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// E7141E0A-04B8-4FC5-A4DC-195392E57807
var IID_IAnimationObject = syscall.GUID{0xE7141E0A, 0x04B8, 0x4FC5,
	[8]byte{0xA4, 0xDC, 0x19, 0x53, 0x92, 0xE5, 0x78, 0x07}}

type IAnimationObjectInterface interface {
	win32.IInspectableInterface
	PopulatePropertyInfo(propertyName string, propertyInfo *IAnimationPropertyInfo)
}

type IAnimationObjectVtbl struct {
	win32.IInspectableVtbl
	PopulatePropertyInfo uintptr
}

type IAnimationObject struct {
	win32.IInspectable
}

func (this *IAnimationObject) Vtbl() *IAnimationObjectVtbl {
	return (*IAnimationObjectVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAnimationObject) PopulatePropertyInfo(propertyName string, propertyInfo *IAnimationPropertyInfo) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PopulatePropertyInfo, uintptr(unsafe.Pointer(this)), NewHStr(propertyName).Ptr, uintptr(unsafe.Pointer(propertyInfo)))
	_ = _hr
}

// F4716F05-ED77-4E3C-B328-5C3985B3738F
var IID_IAnimationPropertyInfo = syscall.GUID{0xF4716F05, 0xED77, 0x4E3C,
	[8]byte{0xB3, 0x28, 0x5C, 0x39, 0x85, 0xB3, 0x73, 0x8F}}

type IAnimationPropertyInfoInterface interface {
	win32.IInspectableInterface
	Get_AccessMode() AnimationPropertyAccessMode
	Put_AccessMode(value AnimationPropertyAccessMode)
}

type IAnimationPropertyInfoVtbl struct {
	win32.IInspectableVtbl
	Get_AccessMode uintptr
	Put_AccessMode uintptr
}

type IAnimationPropertyInfo struct {
	win32.IInspectable
}

func (this *IAnimationPropertyInfo) Vtbl() *IAnimationPropertyInfoVtbl {
	return (*IAnimationPropertyInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAnimationPropertyInfo) Get_AccessMode() AnimationPropertyAccessMode {
	var _result AnimationPropertyAccessMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AccessMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAnimationPropertyInfo) Put_AccessMode(value AnimationPropertyAccessMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AccessMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 591720B4-7472-5218-8B39-DFFE615AE6DA
var IID_IAnimationPropertyInfo2 = syscall.GUID{0x591720B4, 0x7472, 0x5218,
	[8]byte{0x8B, 0x39, 0xDF, 0xFE, 0x61, 0x5A, 0xE6, 0xDA}}

type IAnimationPropertyInfo2Interface interface {
	win32.IInspectableInterface
	GetResolvedCompositionObject() *ICompositionObject
	GetResolvedCompositionObjectProperty() string
}

type IAnimationPropertyInfo2Vtbl struct {
	win32.IInspectableVtbl
	GetResolvedCompositionObject         uintptr
	GetResolvedCompositionObjectProperty uintptr
}

type IAnimationPropertyInfo2 struct {
	win32.IInspectable
}

func (this *IAnimationPropertyInfo2) Vtbl() *IAnimationPropertyInfo2Vtbl {
	return (*IAnimationPropertyInfo2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAnimationPropertyInfo2) GetResolvedCompositionObject() *ICompositionObject {
	var _result *ICompositionObject
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetResolvedCompositionObject, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAnimationPropertyInfo2) GetResolvedCompositionObjectProperty() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetResolvedCompositionObjectProperty, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// B8560DA4-5E3C-545D-B263-7987A2BD27CB
var IID_IBackEasingFunction = syscall.GUID{0xB8560DA4, 0x5E3C, 0x545D,
	[8]byte{0xB2, 0x63, 0x79, 0x87, 0xA2, 0xBD, 0x27, 0xCB}}

type IBackEasingFunctionInterface interface {
	win32.IInspectableInterface
	Get_Mode() CompositionEasingFunctionMode
	Get_Amplitude() float32
}

type IBackEasingFunctionVtbl struct {
	win32.IInspectableVtbl
	Get_Mode      uintptr
	Get_Amplitude uintptr
}

type IBackEasingFunction struct {
	win32.IInspectable
}

func (this *IBackEasingFunction) Vtbl() *IBackEasingFunctionVtbl {
	return (*IBackEasingFunctionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackEasingFunction) Get_Mode() CompositionEasingFunctionMode {
	var _result CompositionEasingFunctionMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackEasingFunction) Get_Amplitude() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Amplitude, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 95E23A08-D1F4-4972-9770-3EFE68D82E14
var IID_IBooleanKeyFrameAnimation = syscall.GUID{0x95E23A08, 0xD1F4, 0x4972,
	[8]byte{0x97, 0x70, 0x3E, 0xFE, 0x68, 0xD8, 0x2E, 0x14}}

type IBooleanKeyFrameAnimationInterface interface {
	win32.IInspectableInterface
	InsertKeyFrame(normalizedProgressKey float32, value bool)
}

type IBooleanKeyFrameAnimationVtbl struct {
	win32.IInspectableVtbl
	InsertKeyFrame uintptr
}

type IBooleanKeyFrameAnimation struct {
	win32.IInspectable
}

func (this *IBooleanKeyFrameAnimation) Vtbl() *IBooleanKeyFrameAnimationVtbl {
	return (*IBooleanKeyFrameAnimationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBooleanKeyFrameAnimation) InsertKeyFrame(normalizedProgressKey float32, value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertKeyFrame, uintptr(unsafe.Pointer(this)), uintptr(normalizedProgressKey), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// E7FDB44B-AAD5-5174-9421-EEF8B75A6A43
var IID_IBounceEasingFunction = syscall.GUID{0xE7FDB44B, 0xAAD5, 0x5174,
	[8]byte{0x94, 0x21, 0xEE, 0xF8, 0xB7, 0x5A, 0x6A, 0x43}}

type IBounceEasingFunctionInterface interface {
	win32.IInspectableInterface
	Get_Mode() CompositionEasingFunctionMode
	Get_Bounces() int32
	Get_Bounciness() float32
}

type IBounceEasingFunctionVtbl struct {
	win32.IInspectableVtbl
	Get_Mode       uintptr
	Get_Bounces    uintptr
	Get_Bounciness uintptr
}

type IBounceEasingFunction struct {
	win32.IInspectable
}

func (this *IBounceEasingFunction) Vtbl() *IBounceEasingFunctionVtbl {
	return (*IBounceEasingFunctionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBounceEasingFunction) Get_Mode() CompositionEasingFunctionMode {
	var _result CompositionEasingFunctionMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBounceEasingFunction) Get_Bounces() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bounces, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBounceEasingFunction) Get_Bounciness() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bounciness, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// BAA30DCC-A633-4618-9B06-7F7C72C87CFF
var IID_IBounceScalarNaturalMotionAnimation = syscall.GUID{0xBAA30DCC, 0xA633, 0x4618,
	[8]byte{0x9B, 0x06, 0x7F, 0x7C, 0x72, 0xC8, 0x7C, 0xFF}}

type IBounceScalarNaturalMotionAnimationInterface interface {
	win32.IInspectableInterface
	Get_Acceleration() float32
	Put_Acceleration(value float32)
	Get_Restitution() float32
	Put_Restitution(value float32)
}

type IBounceScalarNaturalMotionAnimationVtbl struct {
	win32.IInspectableVtbl
	Get_Acceleration uintptr
	Put_Acceleration uintptr
	Get_Restitution  uintptr
	Put_Restitution  uintptr
}

type IBounceScalarNaturalMotionAnimation struct {
	win32.IInspectable
}

func (this *IBounceScalarNaturalMotionAnimation) Vtbl() *IBounceScalarNaturalMotionAnimationVtbl {
	return (*IBounceScalarNaturalMotionAnimationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBounceScalarNaturalMotionAnimation) Get_Acceleration() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Acceleration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBounceScalarNaturalMotionAnimation) Put_Acceleration(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Acceleration, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IBounceScalarNaturalMotionAnimation) Get_Restitution() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Restitution, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBounceScalarNaturalMotionAnimation) Put_Restitution(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Restitution, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// DA344196-2154-4B3C-88AA-47361204ECCD
var IID_IBounceVector2NaturalMotionAnimation = syscall.GUID{0xDA344196, 0x2154, 0x4B3C,
	[8]byte{0x88, 0xAA, 0x47, 0x36, 0x12, 0x04, 0xEC, 0xCD}}

type IBounceVector2NaturalMotionAnimationInterface interface {
	win32.IInspectableInterface
	Get_Acceleration() float32
	Put_Acceleration(value float32)
	Get_Restitution() float32
	Put_Restitution(value float32)
}

type IBounceVector2NaturalMotionAnimationVtbl struct {
	win32.IInspectableVtbl
	Get_Acceleration uintptr
	Put_Acceleration uintptr
	Get_Restitution  uintptr
	Put_Restitution  uintptr
}

type IBounceVector2NaturalMotionAnimation struct {
	win32.IInspectable
}

func (this *IBounceVector2NaturalMotionAnimation) Vtbl() *IBounceVector2NaturalMotionAnimationVtbl {
	return (*IBounceVector2NaturalMotionAnimationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBounceVector2NaturalMotionAnimation) Get_Acceleration() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Acceleration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBounceVector2NaturalMotionAnimation) Put_Acceleration(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Acceleration, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IBounceVector2NaturalMotionAnimation) Get_Restitution() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Restitution, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBounceVector2NaturalMotionAnimation) Put_Restitution(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Restitution, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 47DABC31-10D3-4518-86F1-09CAF742D113
var IID_IBounceVector3NaturalMotionAnimation = syscall.GUID{0x47DABC31, 0x10D3, 0x4518,
	[8]byte{0x86, 0xF1, 0x09, 0xCA, 0xF7, 0x42, 0xD1, 0x13}}

type IBounceVector3NaturalMotionAnimationInterface interface {
	win32.IInspectableInterface
	Get_Acceleration() float32
	Put_Acceleration(value float32)
	Get_Restitution() float32
	Put_Restitution(value float32)
}

type IBounceVector3NaturalMotionAnimationVtbl struct {
	win32.IInspectableVtbl
	Get_Acceleration uintptr
	Put_Acceleration uintptr
	Get_Restitution  uintptr
	Put_Restitution  uintptr
}

type IBounceVector3NaturalMotionAnimation struct {
	win32.IInspectable
}

func (this *IBounceVector3NaturalMotionAnimation) Vtbl() *IBounceVector3NaturalMotionAnimationVtbl {
	return (*IBounceVector3NaturalMotionAnimationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBounceVector3NaturalMotionAnimation) Get_Acceleration() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Acceleration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBounceVector3NaturalMotionAnimation) Put_Acceleration(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Acceleration, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IBounceVector3NaturalMotionAnimation) Get_Restitution() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Restitution, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBounceVector3NaturalMotionAnimation) Put_Restitution(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Restitution, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 1E07222A-6F82-5A28-8748-2E92FC46EE2B
var IID_ICircleEasingFunction = syscall.GUID{0x1E07222A, 0x6F82, 0x5A28,
	[8]byte{0x87, 0x48, 0x2E, 0x92, 0xFC, 0x46, 0xEE, 0x2B}}

type ICircleEasingFunctionInterface interface {
	win32.IInspectableInterface
	Get_Mode() CompositionEasingFunctionMode
}

type ICircleEasingFunctionVtbl struct {
	win32.IInspectableVtbl
	Get_Mode uintptr
}

type ICircleEasingFunction struct {
	win32.IInspectable
}

func (this *ICircleEasingFunction) Vtbl() *ICircleEasingFunctionVtbl {
	return (*ICircleEasingFunctionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICircleEasingFunction) Get_Mode() CompositionEasingFunctionMode {
	var _result CompositionEasingFunctionMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 93ADB5E9-8E05-4593-84A3-DCA152781E56
var IID_IColorKeyFrameAnimation = syscall.GUID{0x93ADB5E9, 0x8E05, 0x4593,
	[8]byte{0x84, 0xA3, 0xDC, 0xA1, 0x52, 0x78, 0x1E, 0x56}}

type IColorKeyFrameAnimationInterface interface {
	win32.IInspectableInterface
	Get_InterpolationColorSpace() CompositionColorSpace
	Put_InterpolationColorSpace(value CompositionColorSpace)
	InsertKeyFrame(normalizedProgressKey float32, value unsafe.Pointer)
	InsertKeyFrameWithEasingFunction(normalizedProgressKey float32, value unsafe.Pointer, easingFunction *ICompositionEasingFunction)
}

type IColorKeyFrameAnimationVtbl struct {
	win32.IInspectableVtbl
	Get_InterpolationColorSpace      uintptr
	Put_InterpolationColorSpace      uintptr
	InsertKeyFrame                   uintptr
	InsertKeyFrameWithEasingFunction uintptr
}

type IColorKeyFrameAnimation struct {
	win32.IInspectable
}

func (this *IColorKeyFrameAnimation) Vtbl() *IColorKeyFrameAnimationVtbl {
	return (*IColorKeyFrameAnimationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IColorKeyFrameAnimation) Get_InterpolationColorSpace() CompositionColorSpace {
	var _result CompositionColorSpace
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InterpolationColorSpace, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IColorKeyFrameAnimation) Put_InterpolationColorSpace(value CompositionColorSpace) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InterpolationColorSpace, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IColorKeyFrameAnimation) InsertKeyFrame(normalizedProgressKey float32, value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertKeyFrame, uintptr(unsafe.Pointer(this)), uintptr(normalizedProgressKey), uintptr(value))
	_ = _hr
}

func (this *IColorKeyFrameAnimation) InsertKeyFrameWithEasingFunction(normalizedProgressKey float32, value unsafe.Pointer, easingFunction *ICompositionEasingFunction) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertKeyFrameWithEasingFunction, uintptr(unsafe.Pointer(this)), uintptr(normalizedProgressKey), uintptr(value), uintptr(unsafe.Pointer(easingFunction)))
	_ = _hr
}

// 464C4C2C-1CAA-4061-9B40-E13FDE1503CA
var IID_ICompositionAnimation = syscall.GUID{0x464C4C2C, 0x1CAA, 0x4061,
	[8]byte{0x9B, 0x40, 0xE1, 0x3F, 0xDE, 0x15, 0x03, 0xCA}}

type ICompositionAnimationInterface interface {
	win32.IInspectableInterface
	ClearAllParameters()
	ClearParameter(key string)
	SetColorParameter(key string, value unsafe.Pointer)
	SetMatrix3x2Parameter(key string, value unsafe.Pointer)
	SetMatrix4x4Parameter(key string, value unsafe.Pointer)
	SetQuaternionParameter(key string, value unsafe.Pointer)
	SetReferenceParameter(key string, compositionObject *ICompositionObject)
	SetScalarParameter(key string, value float32)
	SetVector2Parameter(key string, value unsafe.Pointer)
	SetVector3Parameter(key string, value unsafe.Pointer)
	SetVector4Parameter(key string, value unsafe.Pointer)
}

type ICompositionAnimationVtbl struct {
	win32.IInspectableVtbl
	ClearAllParameters     uintptr
	ClearParameter         uintptr
	SetColorParameter      uintptr
	SetMatrix3x2Parameter  uintptr
	SetMatrix4x4Parameter  uintptr
	SetQuaternionParameter uintptr
	SetReferenceParameter  uintptr
	SetScalarParameter     uintptr
	SetVector2Parameter    uintptr
	SetVector3Parameter    uintptr
	SetVector4Parameter    uintptr
}

type ICompositionAnimation struct {
	win32.IInspectable
}

func (this *ICompositionAnimation) Vtbl() *ICompositionAnimationVtbl {
	return (*ICompositionAnimationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionAnimation) ClearAllParameters() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearAllParameters, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *ICompositionAnimation) ClearParameter(key string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearParameter, uintptr(unsafe.Pointer(this)), NewHStr(key).Ptr)
	_ = _hr
}

func (this *ICompositionAnimation) SetColorParameter(key string, value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetColorParameter, uintptr(unsafe.Pointer(this)), NewHStr(key).Ptr, uintptr(value))
	_ = _hr
}

func (this *ICompositionAnimation) SetMatrix3x2Parameter(key string, value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetMatrix3x2Parameter, uintptr(unsafe.Pointer(this)), NewHStr(key).Ptr, uintptr(value))
	_ = _hr
}

func (this *ICompositionAnimation) SetMatrix4x4Parameter(key string, value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetMatrix4x4Parameter, uintptr(unsafe.Pointer(this)), NewHStr(key).Ptr, uintptr(value))
	_ = _hr
}

func (this *ICompositionAnimation) SetQuaternionParameter(key string, value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetQuaternionParameter, uintptr(unsafe.Pointer(this)), NewHStr(key).Ptr, uintptr(value))
	_ = _hr
}

func (this *ICompositionAnimation) SetReferenceParameter(key string, compositionObject *ICompositionObject) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetReferenceParameter, uintptr(unsafe.Pointer(this)), NewHStr(key).Ptr, uintptr(unsafe.Pointer(compositionObject)))
	_ = _hr
}

func (this *ICompositionAnimation) SetScalarParameter(key string, value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetScalarParameter, uintptr(unsafe.Pointer(this)), NewHStr(key).Ptr, uintptr(value))
	_ = _hr
}

func (this *ICompositionAnimation) SetVector2Parameter(key string, value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetVector2Parameter, uintptr(unsafe.Pointer(this)), NewHStr(key).Ptr, uintptr(value))
	_ = _hr
}

func (this *ICompositionAnimation) SetVector3Parameter(key string, value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetVector3Parameter, uintptr(unsafe.Pointer(this)), NewHStr(key).Ptr, uintptr(value))
	_ = _hr
}

func (this *ICompositionAnimation) SetVector4Parameter(key string, value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetVector4Parameter, uintptr(unsafe.Pointer(this)), NewHStr(key).Ptr, uintptr(value))
	_ = _hr
}

// 369B603E-A80F-4948-93E3-ED23FB38C6CB
var IID_ICompositionAnimation2 = syscall.GUID{0x369B603E, 0xA80F, 0x4948,
	[8]byte{0x93, 0xE3, 0xED, 0x23, 0xFB, 0x38, 0xC6, 0xCB}}

type ICompositionAnimation2Interface interface {
	win32.IInspectableInterface
	SetBooleanParameter(key string, value bool)
	Get_Target() string
	Put_Target(value string)
}

type ICompositionAnimation2Vtbl struct {
	win32.IInspectableVtbl
	SetBooleanParameter uintptr
	Get_Target          uintptr
	Put_Target          uintptr
}

type ICompositionAnimation2 struct {
	win32.IInspectable
}

func (this *ICompositionAnimation2) Vtbl() *ICompositionAnimation2Vtbl {
	return (*ICompositionAnimation2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionAnimation2) SetBooleanParameter(key string, value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetBooleanParameter, uintptr(unsafe.Pointer(this)), NewHStr(key).Ptr, uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICompositionAnimation2) Get_Target() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Target, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICompositionAnimation2) Put_Target(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Target, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// D51E030D-7DA4-4BD7-BC2D-F4517529F43A
var IID_ICompositionAnimation3 = syscall.GUID{0xD51E030D, 0x7DA4, 0x4BD7,
	[8]byte{0xBC, 0x2D, 0xF4, 0x51, 0x75, 0x29, 0xF4, 0x3A}}

type ICompositionAnimation3Interface interface {
	win32.IInspectableInterface
	Get_InitialValueExpressions() *IMap[string, string]
}

type ICompositionAnimation3Vtbl struct {
	win32.IInspectableVtbl
	Get_InitialValueExpressions uintptr
}

type ICompositionAnimation3 struct {
	win32.IInspectable
}

func (this *ICompositionAnimation3) Vtbl() *ICompositionAnimation3Vtbl {
	return (*ICompositionAnimation3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionAnimation3) Get_InitialValueExpressions() *IMap[string, string] {
	var _result *IMap[string, string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InitialValueExpressions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 770137BE-76BC-4E23-BFED-FE9CC20F6EC9
var IID_ICompositionAnimation4 = syscall.GUID{0x770137BE, 0x76BC, 0x4E23,
	[8]byte{0xBF, 0xED, 0xFE, 0x9C, 0xC2, 0x0F, 0x6E, 0xC9}}

type ICompositionAnimation4Interface interface {
	win32.IInspectableInterface
	SetExpressionReferenceParameter(parameterName string, source *IAnimationObject)
}

type ICompositionAnimation4Vtbl struct {
	win32.IInspectableVtbl
	SetExpressionReferenceParameter uintptr
}

type ICompositionAnimation4 struct {
	win32.IInspectable
}

func (this *ICompositionAnimation4) Vtbl() *ICompositionAnimation4Vtbl {
	return (*ICompositionAnimation4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionAnimation4) SetExpressionReferenceParameter(parameterName string, source *IAnimationObject) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetExpressionReferenceParameter, uintptr(unsafe.Pointer(this)), NewHStr(parameterName).Ptr, uintptr(unsafe.Pointer(source)))
	_ = _hr
}

// 1C2C2999-E818-48D3-A6DD-D78C82F8ACE9
var IID_ICompositionAnimationBase = syscall.GUID{0x1C2C2999, 0xE818, 0x48D3,
	[8]byte{0xA6, 0xDD, 0xD7, 0x8C, 0x82, 0xF8, 0xAC, 0xE9}}

type ICompositionAnimationBaseInterface interface {
	win32.IInspectableInterface
}

type ICompositionAnimationBaseVtbl struct {
	win32.IInspectableVtbl
}

type ICompositionAnimationBase struct {
	win32.IInspectable
}

func (this *ICompositionAnimationBase) Vtbl() *ICompositionAnimationBaseVtbl {
	return (*ICompositionAnimationBaseVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 10F6C4FB-6E51-4C25-BBD3-586A9BEC3EF4
var IID_ICompositionAnimationFactory = syscall.GUID{0x10F6C4FB, 0x6E51, 0x4C25,
	[8]byte{0xBB, 0xD3, 0x58, 0x6A, 0x9B, 0xEC, 0x3E, 0xF4}}

type ICompositionAnimationFactoryInterface interface {
	win32.IInspectableInterface
}

type ICompositionAnimationFactoryVtbl struct {
	win32.IInspectableVtbl
}

type ICompositionAnimationFactory struct {
	win32.IInspectable
}

func (this *ICompositionAnimationFactory) Vtbl() *ICompositionAnimationFactoryVtbl {
	return (*ICompositionAnimationFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 5E7CC90C-CD14-4E07-8A55-C72527AABDAC
var IID_ICompositionAnimationGroup = syscall.GUID{0x5E7CC90C, 0xCD14, 0x4E07,
	[8]byte{0x8A, 0x55, 0xC7, 0x25, 0x27, 0xAA, 0xBD, 0xAC}}

type ICompositionAnimationGroupInterface interface {
	win32.IInspectableInterface
	Get_Count() int32
	Add(value *ICompositionAnimation)
	Remove(value *ICompositionAnimation)
	RemoveAll()
}

type ICompositionAnimationGroupVtbl struct {
	win32.IInspectableVtbl
	Get_Count uintptr
	Add       uintptr
	Remove    uintptr
	RemoveAll uintptr
}

type ICompositionAnimationGroup struct {
	win32.IInspectable
}

func (this *ICompositionAnimationGroup) Vtbl() *ICompositionAnimationGroupVtbl {
	return (*ICompositionAnimationGroupVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionAnimationGroup) Get_Count() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Count, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionAnimationGroup) Add(value *ICompositionAnimation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ICompositionAnimationGroup) Remove(value *ICompositionAnimation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ICompositionAnimationGroup) RemoveAll() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveAll, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// C5ACAE58-3898-499E-8D7F-224E91286A5D
var IID_ICompositionBackdropBrush = syscall.GUID{0xC5ACAE58, 0x3898, 0x499E,
	[8]byte{0x8D, 0x7F, 0x22, 0x4E, 0x91, 0x28, 0x6A, 0x5D}}

type ICompositionBackdropBrushInterface interface {
	win32.IInspectableInterface
}

type ICompositionBackdropBrushVtbl struct {
	win32.IInspectableVtbl
}

type ICompositionBackdropBrush struct {
	win32.IInspectable
}

func (this *ICompositionBackdropBrush) Vtbl() *ICompositionBackdropBrushVtbl {
	return (*ICompositionBackdropBrushVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 0D00DAD0-9464-450A-A562-2E2698B0A812
var IID_ICompositionBatchCompletedEventArgs = syscall.GUID{0x0D00DAD0, 0x9464, 0x450A,
	[8]byte{0xA5, 0x62, 0x2E, 0x26, 0x98, 0xB0, 0xA8, 0x12}}

type ICompositionBatchCompletedEventArgsInterface interface {
	win32.IInspectableInterface
}

type ICompositionBatchCompletedEventArgsVtbl struct {
	win32.IInspectableVtbl
}

type ICompositionBatchCompletedEventArgs struct {
	win32.IInspectable
}

func (this *ICompositionBatchCompletedEventArgs) Vtbl() *ICompositionBatchCompletedEventArgsVtbl {
	return (*ICompositionBatchCompletedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// AB0D7608-30C0-40E9-B568-B60A6BD1FB46
var IID_ICompositionBrush = syscall.GUID{0xAB0D7608, 0x30C0, 0x40E9,
	[8]byte{0xB5, 0x68, 0xB6, 0x0A, 0x6B, 0xD1, 0xFB, 0x46}}

type ICompositionBrushInterface interface {
	win32.IInspectableInterface
}

type ICompositionBrushVtbl struct {
	win32.IInspectableVtbl
}

type ICompositionBrush struct {
	win32.IInspectable
}

func (this *ICompositionBrush) Vtbl() *ICompositionBrushVtbl {
	return (*ICompositionBrushVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// DA53FB4C-4650-47C4-AD76-765379607ED6
var IID_ICompositionBrushFactory = syscall.GUID{0xDA53FB4C, 0x4650, 0x47C4,
	[8]byte{0xAD, 0x76, 0x76, 0x53, 0x79, 0x60, 0x7E, 0xD6}}

type ICompositionBrushFactoryInterface interface {
	win32.IInspectableInterface
}

type ICompositionBrushFactoryVtbl struct {
	win32.IInspectableVtbl
}

type ICompositionBrushFactory struct {
	win32.IInspectable
}

func (this *ICompositionBrushFactory) Vtbl() *ICompositionBrushFactoryVtbl {
	return (*ICompositionBrushFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 8253353E-B517-48BC-B1E8-4B3561A2E181
var IID_ICompositionCapabilities = syscall.GUID{0x8253353E, 0xB517, 0x48BC,
	[8]byte{0xB1, 0xE8, 0x4B, 0x35, 0x61, 0xA2, 0xE1, 0x81}}

type ICompositionCapabilitiesInterface interface {
	win32.IInspectableInterface
	AreEffectsSupported() bool
	AreEffectsFast() bool
	Add_Changed(handler TypedEventHandler[*ICompositionCapabilities, interface{}]) EventRegistrationToken
	Remove_Changed(token EventRegistrationToken)
}

type ICompositionCapabilitiesVtbl struct {
	win32.IInspectableVtbl
	AreEffectsSupported uintptr
	AreEffectsFast      uintptr
	Add_Changed         uintptr
	Remove_Changed      uintptr
}

type ICompositionCapabilities struct {
	win32.IInspectable
}

func (this *ICompositionCapabilities) Vtbl() *ICompositionCapabilitiesVtbl {
	return (*ICompositionCapabilitiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionCapabilities) AreEffectsSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AreEffectsSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionCapabilities) AreEffectsFast() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AreEffectsFast, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionCapabilities) Add_Changed(handler TypedEventHandler[*ICompositionCapabilities, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Changed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionCapabilities) Remove_Changed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Changed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// F7B7A86E-6416-49E5-8DDF-AFE949E20562
var IID_ICompositionCapabilitiesStatics = syscall.GUID{0xF7B7A86E, 0x6416, 0x49E5,
	[8]byte{0x8D, 0xDF, 0xAF, 0xE9, 0x49, 0xE2, 0x05, 0x62}}

type ICompositionCapabilitiesStaticsInterface interface {
	win32.IInspectableInterface
	GetForCurrentView() *ICompositionCapabilities
}

type ICompositionCapabilitiesStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForCurrentView uintptr
}

type ICompositionCapabilitiesStatics struct {
	win32.IInspectable
}

func (this *ICompositionCapabilitiesStatics) Vtbl() *ICompositionCapabilitiesStaticsVtbl {
	return (*ICompositionCapabilitiesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionCapabilitiesStatics) GetForCurrentView() *ICompositionCapabilities {
	var _result *ICompositionCapabilities
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1CCD2A52-CFC7-4ACE-9983-146BB8EB6A3C
var IID_ICompositionClip = syscall.GUID{0x1CCD2A52, 0xCFC7, 0x4ACE,
	[8]byte{0x99, 0x83, 0x14, 0x6B, 0xB8, 0xEB, 0x6A, 0x3C}}

type ICompositionClipInterface interface {
	win32.IInspectableInterface
}

type ICompositionClipVtbl struct {
	win32.IInspectableVtbl
}

type ICompositionClip struct {
	win32.IInspectable
}

func (this *ICompositionClip) Vtbl() *ICompositionClipVtbl {
	return (*ICompositionClipVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 5893E069-3516-40E1-89E0-5BA924927235
var IID_ICompositionClip2 = syscall.GUID{0x5893E069, 0x3516, 0x40E1,
	[8]byte{0x89, 0xE0, 0x5B, 0xA9, 0x24, 0x92, 0x72, 0x35}}

type ICompositionClip2Interface interface {
	win32.IInspectableInterface
	Get_AnchorPoint() unsafe.Pointer
	Put_AnchorPoint(value unsafe.Pointer)
	Get_CenterPoint() unsafe.Pointer
	Put_CenterPoint(value unsafe.Pointer)
	Get_Offset() unsafe.Pointer
	Put_Offset(value unsafe.Pointer)
	Get_RotationAngle() float32
	Put_RotationAngle(value float32)
	Get_RotationAngleInDegrees() float32
	Put_RotationAngleInDegrees(value float32)
	Get_Scale() unsafe.Pointer
	Put_Scale(value unsafe.Pointer)
	Get_TransformMatrix() unsafe.Pointer
	Put_TransformMatrix(value unsafe.Pointer)
}

type ICompositionClip2Vtbl struct {
	win32.IInspectableVtbl
	Get_AnchorPoint            uintptr
	Put_AnchorPoint            uintptr
	Get_CenterPoint            uintptr
	Put_CenterPoint            uintptr
	Get_Offset                 uintptr
	Put_Offset                 uintptr
	Get_RotationAngle          uintptr
	Put_RotationAngle          uintptr
	Get_RotationAngleInDegrees uintptr
	Put_RotationAngleInDegrees uintptr
	Get_Scale                  uintptr
	Put_Scale                  uintptr
	Get_TransformMatrix        uintptr
	Put_TransformMatrix        uintptr
}

type ICompositionClip2 struct {
	win32.IInspectable
}

func (this *ICompositionClip2) Vtbl() *ICompositionClip2Vtbl {
	return (*ICompositionClip2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionClip2) Get_AnchorPoint() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AnchorPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionClip2) Put_AnchorPoint(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AnchorPoint, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionClip2) Get_CenterPoint() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CenterPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionClip2) Put_CenterPoint(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CenterPoint, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionClip2) Get_Offset() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Offset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionClip2) Put_Offset(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Offset, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionClip2) Get_RotationAngle() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RotationAngle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionClip2) Put_RotationAngle(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RotationAngle, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionClip2) Get_RotationAngleInDegrees() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RotationAngleInDegrees, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionClip2) Put_RotationAngleInDegrees(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RotationAngleInDegrees, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionClip2) Get_Scale() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Scale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionClip2) Put_Scale(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Scale, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionClip2) Get_TransformMatrix() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TransformMatrix, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionClip2) Put_TransformMatrix(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TransformMatrix, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// B9484CAF-20C7-4AED-AC4A-9C78BA1302CF
var IID_ICompositionClipFactory = syscall.GUID{0xB9484CAF, 0x20C7, 0x4AED,
	[8]byte{0xAC, 0x4A, 0x9C, 0x78, 0xBA, 0x13, 0x02, 0xCF}}

type ICompositionClipFactoryInterface interface {
	win32.IInspectableInterface
}

type ICompositionClipFactoryVtbl struct {
	win32.IInspectableVtbl
}

type ICompositionClipFactory struct {
	win32.IInspectable
}

func (this *ICompositionClipFactory) Vtbl() *ICompositionClipFactoryVtbl {
	return (*ICompositionClipFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 2B264C5E-BF35-4831-8642-CF70C20FFF2F
var IID_ICompositionColorBrush = syscall.GUID{0x2B264C5E, 0xBF35, 0x4831,
	[8]byte{0x86, 0x42, 0xCF, 0x70, 0xC2, 0x0F, 0xFF, 0x2F}}

type ICompositionColorBrushInterface interface {
	win32.IInspectableInterface
	Get_Color() unsafe.Pointer
	Put_Color(value unsafe.Pointer)
}

type ICompositionColorBrushVtbl struct {
	win32.IInspectableVtbl
	Get_Color uintptr
	Put_Color uintptr
}

type ICompositionColorBrush struct {
	win32.IInspectable
}

func (this *ICompositionColorBrush) Vtbl() *ICompositionColorBrushVtbl {
	return (*ICompositionColorBrushVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionColorBrush) Get_Color() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Color, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionColorBrush) Put_Color(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Color, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 6F00CA92-C801-4E41-9A8F-A53E20F57778
var IID_ICompositionColorGradientStop = syscall.GUID{0x6F00CA92, 0xC801, 0x4E41,
	[8]byte{0x9A, 0x8F, 0xA5, 0x3E, 0x20, 0xF5, 0x77, 0x78}}

type ICompositionColorGradientStopInterface interface {
	win32.IInspectableInterface
	Get_Color() unsafe.Pointer
	Put_Color(value unsafe.Pointer)
	Get_Offset() float32
	Put_Offset(value float32)
}

type ICompositionColorGradientStopVtbl struct {
	win32.IInspectableVtbl
	Get_Color  uintptr
	Put_Color  uintptr
	Get_Offset uintptr
	Put_Offset uintptr
}

type ICompositionColorGradientStop struct {
	win32.IInspectable
}

func (this *ICompositionColorGradientStop) Vtbl() *ICompositionColorGradientStopVtbl {
	return (*ICompositionColorGradientStopVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionColorGradientStop) Get_Color() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Color, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionColorGradientStop) Put_Color(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Color, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionColorGradientStop) Get_Offset() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Offset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionColorGradientStop) Put_Offset(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Offset, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 9F1D20EC-7B04-4B1D-90BC-9FA32C0CFD26
var IID_ICompositionColorGradientStopCollection = syscall.GUID{0x9F1D20EC, 0x7B04, 0x4B1D,
	[8]byte{0x90, 0xBC, 0x9F, 0xA3, 0x2C, 0x0C, 0xFD, 0x26}}

type ICompositionColorGradientStopCollectionInterface interface {
	win32.IInspectableInterface
}

type ICompositionColorGradientStopCollectionVtbl struct {
	win32.IInspectableVtbl
}

type ICompositionColorGradientStopCollection struct {
	win32.IInspectable
}

func (this *ICompositionColorGradientStopCollection) Vtbl() *ICompositionColorGradientStopCollectionVtbl {
	return (*ICompositionColorGradientStopCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 0D00DAD0-CA07-4400-8C8E-CB5DB08559CC
var IID_ICompositionCommitBatch = syscall.GUID{0x0D00DAD0, 0xCA07, 0x4400,
	[8]byte{0x8C, 0x8E, 0xCB, 0x5D, 0xB0, 0x85, 0x59, 0xCC}}

type ICompositionCommitBatchInterface interface {
	win32.IInspectableInterface
	Get_IsActive() bool
	Get_IsEnded() bool
	Add_Completed(handler TypedEventHandler[interface{}, *ICompositionBatchCompletedEventArgs]) EventRegistrationToken
	Remove_Completed(token EventRegistrationToken)
}

type ICompositionCommitBatchVtbl struct {
	win32.IInspectableVtbl
	Get_IsActive     uintptr
	Get_IsEnded      uintptr
	Add_Completed    uintptr
	Remove_Completed uintptr
}

type ICompositionCommitBatch struct {
	win32.IInspectable
}

func (this *ICompositionCommitBatch) Vtbl() *ICompositionCommitBatchVtbl {
	return (*ICompositionCommitBatchVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionCommitBatch) Get_IsActive() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsActive, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionCommitBatch) Get_IsEnded() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEnded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionCommitBatch) Add_Completed(handler TypedEventHandler[interface{}, *ICompositionBatchCompletedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Completed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionCommitBatch) Remove_Completed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Completed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 4F5E859B-2E5B-44A8-982C-AA0F69C16059
var IID_ICompositionContainerShape = syscall.GUID{0x4F5E859B, 0x2E5B, 0x44A8,
	[8]byte{0x98, 0x2C, 0xAA, 0x0F, 0x69, 0xC1, 0x60, 0x59}}

type ICompositionContainerShapeInterface interface {
	win32.IInspectableInterface
	Get_Shapes() *IVector[*ICompositionShape]
}

type ICompositionContainerShapeVtbl struct {
	win32.IInspectableVtbl
	Get_Shapes uintptr
}

type ICompositionContainerShape struct {
	win32.IInspectable
}

func (this *ICompositionContainerShape) Vtbl() *ICompositionContainerShapeVtbl {
	return (*ICompositionContainerShapeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionContainerShape) Get_Shapes() *IVector[*ICompositionShape] {
	var _result *IVector[*ICompositionShape]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Shapes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A166C300-FAD0-4D11-9E67-E433162FF49E
var IID_ICompositionDrawingSurface = syscall.GUID{0xA166C300, 0xFAD0, 0x4D11,
	[8]byte{0x9E, 0x67, 0xE4, 0x33, 0x16, 0x2F, 0xF4, 0x9E}}

type ICompositionDrawingSurfaceInterface interface {
	win32.IInspectableInterface
	Get_AlphaMode() unsafe.Pointer
	Get_PixelFormat() unsafe.Pointer
	Get_Size() Size
}

type ICompositionDrawingSurfaceVtbl struct {
	win32.IInspectableVtbl
	Get_AlphaMode   uintptr
	Get_PixelFormat uintptr
	Get_Size        uintptr
}

type ICompositionDrawingSurface struct {
	win32.IInspectable
}

func (this *ICompositionDrawingSurface) Vtbl() *ICompositionDrawingSurfaceVtbl {
	return (*ICompositionDrawingSurfaceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionDrawingSurface) Get_AlphaMode() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlphaMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionDrawingSurface) Get_PixelFormat() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PixelFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionDrawingSurface) Get_Size() Size {
	var _result Size
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Size, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// FAD0E88B-E354-44E8-8E3D-C4880D5A213F
var IID_ICompositionDrawingSurface2 = syscall.GUID{0xFAD0E88B, 0xE354, 0x44E8,
	[8]byte{0x8E, 0x3D, 0xC4, 0x88, 0x0D, 0x5A, 0x21, 0x3F}}

type ICompositionDrawingSurface2Interface interface {
	win32.IInspectableInterface
	Get_SizeInt32() unsafe.Pointer
	Resize(sizePixels unsafe.Pointer)
	Scroll(offset unsafe.Pointer)
	ScrollRect(offset unsafe.Pointer, scrollRect unsafe.Pointer)
	ScrollWithClip(offset unsafe.Pointer, clipRect unsafe.Pointer)
	ScrollRectWithClip(offset unsafe.Pointer, clipRect unsafe.Pointer, scrollRect unsafe.Pointer)
}

type ICompositionDrawingSurface2Vtbl struct {
	win32.IInspectableVtbl
	Get_SizeInt32      uintptr
	Resize             uintptr
	Scroll             uintptr
	ScrollRect         uintptr
	ScrollWithClip     uintptr
	ScrollRectWithClip uintptr
}

type ICompositionDrawingSurface2 struct {
	win32.IInspectable
}

func (this *ICompositionDrawingSurface2) Vtbl() *ICompositionDrawingSurface2Vtbl {
	return (*ICompositionDrawingSurface2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionDrawingSurface2) Get_SizeInt32() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SizeInt32, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionDrawingSurface2) Resize(sizePixels unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Resize, uintptr(unsafe.Pointer(this)), uintptr(sizePixels))
	_ = _hr
}

func (this *ICompositionDrawingSurface2) Scroll(offset unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Scroll, uintptr(unsafe.Pointer(this)), uintptr(offset))
	_ = _hr
}

func (this *ICompositionDrawingSurface2) ScrollRect(offset unsafe.Pointer, scrollRect unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ScrollRect, uintptr(unsafe.Pointer(this)), uintptr(offset), uintptr(scrollRect))
	_ = _hr
}

func (this *ICompositionDrawingSurface2) ScrollWithClip(offset unsafe.Pointer, clipRect unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ScrollWithClip, uintptr(unsafe.Pointer(this)), uintptr(offset), uintptr(clipRect))
	_ = _hr
}

func (this *ICompositionDrawingSurface2) ScrollRectWithClip(offset unsafe.Pointer, clipRect unsafe.Pointer, scrollRect unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ScrollRectWithClip, uintptr(unsafe.Pointer(this)), uintptr(offset), uintptr(clipRect), uintptr(scrollRect))
	_ = _hr
}

// 9497B00A-312D-46B9-9DB3-412FD79464C8
var IID_ICompositionDrawingSurfaceFactory = syscall.GUID{0x9497B00A, 0x312D, 0x46B9,
	[8]byte{0x9D, 0xB3, 0x41, 0x2F, 0xD7, 0x94, 0x64, 0xC8}}

type ICompositionDrawingSurfaceFactoryInterface interface {
	win32.IInspectableInterface
}

type ICompositionDrawingSurfaceFactoryVtbl struct {
	win32.IInspectableVtbl
}

type ICompositionDrawingSurfaceFactory struct {
	win32.IInspectable
}

func (this *ICompositionDrawingSurfaceFactory) Vtbl() *ICompositionDrawingSurfaceFactoryVtbl {
	return (*ICompositionDrawingSurfaceFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 5145E356-BF79-4EA8-8CC2-6B5B472E6C9A
var IID_ICompositionEasingFunction = syscall.GUID{0x5145E356, 0xBF79, 0x4EA8,
	[8]byte{0x8C, 0xC2, 0x6B, 0x5B, 0x47, 0x2E, 0x6C, 0x9A}}

type ICompositionEasingFunctionInterface interface {
	win32.IInspectableInterface
}

type ICompositionEasingFunctionVtbl struct {
	win32.IInspectableVtbl
}

type ICompositionEasingFunction struct {
	win32.IInspectable
}

func (this *ICompositionEasingFunction) Vtbl() *ICompositionEasingFunctionVtbl {
	return (*ICompositionEasingFunctionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 60840774-3DA0-4949-8200-7206C00190A0
var IID_ICompositionEasingFunctionFactory = syscall.GUID{0x60840774, 0x3DA0, 0x4949,
	[8]byte{0x82, 0x00, 0x72, 0x06, 0xC0, 0x01, 0x90, 0xA0}}

type ICompositionEasingFunctionFactoryInterface interface {
	win32.IInspectableInterface
}

type ICompositionEasingFunctionFactoryVtbl struct {
	win32.IInspectableVtbl
}

type ICompositionEasingFunctionFactory struct {
	win32.IInspectable
}

func (this *ICompositionEasingFunctionFactory) Vtbl() *ICompositionEasingFunctionFactoryVtbl {
	return (*ICompositionEasingFunctionFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 17A766B6-2936-53EA-B5AF-C642F4A61083
var IID_ICompositionEasingFunctionStatics = syscall.GUID{0x17A766B6, 0x2936, 0x53EA,
	[8]byte{0xB5, 0xAF, 0xC6, 0x42, 0xF4, 0xA6, 0x10, 0x83}}

type ICompositionEasingFunctionStaticsInterface interface {
	win32.IInspectableInterface
	CreateCubicBezierEasingFunction(owner *ICompositor, controlPoint1 unsafe.Pointer, controlPoint2 unsafe.Pointer) *ICubicBezierEasingFunction
	CreateLinearEasingFunction(owner *ICompositor) *ILinearEasingFunction
	CreateStepEasingFunction(owner *ICompositor) *IStepEasingFunction
	CreateStepEasingFunctionWithStepCount(owner *ICompositor, stepCount int32) *IStepEasingFunction
	CreateBackEasingFunction(owner *ICompositor, mode CompositionEasingFunctionMode, amplitude float32) *IBackEasingFunction
	CreateBounceEasingFunction(owner *ICompositor, mode CompositionEasingFunctionMode, bounces int32, bounciness float32) *IBounceEasingFunction
	CreateCircleEasingFunction(owner *ICompositor, mode CompositionEasingFunctionMode) *ICircleEasingFunction
	CreateElasticEasingFunction(owner *ICompositor, mode CompositionEasingFunctionMode, oscillations int32, springiness float32) *IElasticEasingFunction
	CreateExponentialEasingFunction(owner *ICompositor, mode CompositionEasingFunctionMode, exponent float32) *IExponentialEasingFunction
	CreatePowerEasingFunction(owner *ICompositor, mode CompositionEasingFunctionMode, power float32) *IPowerEasingFunction
	CreateSineEasingFunction(owner *ICompositor, mode CompositionEasingFunctionMode) *ISineEasingFunction
}

type ICompositionEasingFunctionStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateCubicBezierEasingFunction       uintptr
	CreateLinearEasingFunction            uintptr
	CreateStepEasingFunction              uintptr
	CreateStepEasingFunctionWithStepCount uintptr
	CreateBackEasingFunction              uintptr
	CreateBounceEasingFunction            uintptr
	CreateCircleEasingFunction            uintptr
	CreateElasticEasingFunction           uintptr
	CreateExponentialEasingFunction       uintptr
	CreatePowerEasingFunction             uintptr
	CreateSineEasingFunction              uintptr
}

type ICompositionEasingFunctionStatics struct {
	win32.IInspectable
}

func (this *ICompositionEasingFunctionStatics) Vtbl() *ICompositionEasingFunctionStaticsVtbl {
	return (*ICompositionEasingFunctionStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionEasingFunctionStatics) CreateCubicBezierEasingFunction(owner *ICompositor, controlPoint1 unsafe.Pointer, controlPoint2 unsafe.Pointer) *ICubicBezierEasingFunction {
	var _result *ICubicBezierEasingFunction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateCubicBezierEasingFunction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(owner)), uintptr(controlPoint1), uintptr(controlPoint2), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionEasingFunctionStatics) CreateLinearEasingFunction(owner *ICompositor) *ILinearEasingFunction {
	var _result *ILinearEasingFunction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateLinearEasingFunction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(owner)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionEasingFunctionStatics) CreateStepEasingFunction(owner *ICompositor) *IStepEasingFunction {
	var _result *IStepEasingFunction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateStepEasingFunction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(owner)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionEasingFunctionStatics) CreateStepEasingFunctionWithStepCount(owner *ICompositor, stepCount int32) *IStepEasingFunction {
	var _result *IStepEasingFunction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateStepEasingFunctionWithStepCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(owner)), uintptr(stepCount), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionEasingFunctionStatics) CreateBackEasingFunction(owner *ICompositor, mode CompositionEasingFunctionMode, amplitude float32) *IBackEasingFunction {
	var _result *IBackEasingFunction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateBackEasingFunction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(owner)), uintptr(mode), uintptr(amplitude), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionEasingFunctionStatics) CreateBounceEasingFunction(owner *ICompositor, mode CompositionEasingFunctionMode, bounces int32, bounciness float32) *IBounceEasingFunction {
	var _result *IBounceEasingFunction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateBounceEasingFunction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(owner)), uintptr(mode), uintptr(bounces), uintptr(bounciness), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionEasingFunctionStatics) CreateCircleEasingFunction(owner *ICompositor, mode CompositionEasingFunctionMode) *ICircleEasingFunction {
	var _result *ICircleEasingFunction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateCircleEasingFunction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(owner)), uintptr(mode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionEasingFunctionStatics) CreateElasticEasingFunction(owner *ICompositor, mode CompositionEasingFunctionMode, oscillations int32, springiness float32) *IElasticEasingFunction {
	var _result *IElasticEasingFunction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateElasticEasingFunction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(owner)), uintptr(mode), uintptr(oscillations), uintptr(springiness), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionEasingFunctionStatics) CreateExponentialEasingFunction(owner *ICompositor, mode CompositionEasingFunctionMode, exponent float32) *IExponentialEasingFunction {
	var _result *IExponentialEasingFunction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateExponentialEasingFunction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(owner)), uintptr(mode), uintptr(exponent), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionEasingFunctionStatics) CreatePowerEasingFunction(owner *ICompositor, mode CompositionEasingFunctionMode, power float32) *IPowerEasingFunction {
	var _result *IPowerEasingFunction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreatePowerEasingFunction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(owner)), uintptr(mode), uintptr(power), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionEasingFunctionStatics) CreateSineEasingFunction(owner *ICompositor, mode CompositionEasingFunctionMode) *ISineEasingFunction {
	var _result *ISineEasingFunction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSineEasingFunction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(owner)), uintptr(mode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BF7F795E-83CC-44BF-A447-3E3C071789EC
var IID_ICompositionEffectBrush = syscall.GUID{0xBF7F795E, 0x83CC, 0x44BF,
	[8]byte{0xA4, 0x47, 0x3E, 0x3C, 0x07, 0x17, 0x89, 0xEC}}

type ICompositionEffectBrushInterface interface {
	win32.IInspectableInterface
	GetSourceParameter(name string) *ICompositionBrush
	SetSourceParameter(name string, source *ICompositionBrush)
}

type ICompositionEffectBrushVtbl struct {
	win32.IInspectableVtbl
	GetSourceParameter uintptr
	SetSourceParameter uintptr
}

type ICompositionEffectBrush struct {
	win32.IInspectable
}

func (this *ICompositionEffectBrush) Vtbl() *ICompositionEffectBrushVtbl {
	return (*ICompositionEffectBrushVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionEffectBrush) GetSourceParameter(name string) *ICompositionBrush {
	var _result *ICompositionBrush
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSourceParameter, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionEffectBrush) SetSourceParameter(name string, source *ICompositionBrush) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetSourceParameter, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(source)))
	_ = _hr
}

// BE5624AF-BA7E-4510-9850-41C0B4FF74DF
var IID_ICompositionEffectFactory = syscall.GUID{0xBE5624AF, 0xBA7E, 0x4510,
	[8]byte{0x98, 0x50, 0x41, 0xC0, 0xB4, 0xFF, 0x74, 0xDF}}

type ICompositionEffectFactoryInterface interface {
	win32.IInspectableInterface
	CreateBrush() *ICompositionEffectBrush
	Get_ExtendedError() HResult
	Get_LoadStatus() CompositionEffectFactoryLoadStatus
}

type ICompositionEffectFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateBrush       uintptr
	Get_ExtendedError uintptr
	Get_LoadStatus    uintptr
}

type ICompositionEffectFactory struct {
	win32.IInspectable
}

func (this *ICompositionEffectFactory) Vtbl() *ICompositionEffectFactoryVtbl {
	return (*ICompositionEffectFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionEffectFactory) CreateBrush() *ICompositionEffectBrush {
	var _result *ICompositionEffectBrush
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateBrush, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionEffectFactory) Get_ExtendedError() HResult {
	var _result HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionEffectFactory) Get_LoadStatus() CompositionEffectFactoryLoadStatus {
	var _result CompositionEffectFactoryLoadStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LoadStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 858AB13A-3292-4E4E-B3BB-2B6C6544A6EE
var IID_ICompositionEffectSourceParameter = syscall.GUID{0x858AB13A, 0x3292, 0x4E4E,
	[8]byte{0xB3, 0xBB, 0x2B, 0x6C, 0x65, 0x44, 0xA6, 0xEE}}

type ICompositionEffectSourceParameterInterface interface {
	win32.IInspectableInterface
	Get_Name() string
}

type ICompositionEffectSourceParameterVtbl struct {
	win32.IInspectableVtbl
	Get_Name uintptr
}

type ICompositionEffectSourceParameter struct {
	win32.IInspectable
}

func (this *ICompositionEffectSourceParameter) Vtbl() *ICompositionEffectSourceParameterVtbl {
	return (*ICompositionEffectSourceParameterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionEffectSourceParameter) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// B3D9F276-ABA3-4724-ACF3-D0397464DB1C
var IID_ICompositionEffectSourceParameterFactory = syscall.GUID{0xB3D9F276, 0xABA3, 0x4724,
	[8]byte{0xAC, 0xF3, 0xD0, 0x39, 0x74, 0x64, 0xDB, 0x1C}}

type ICompositionEffectSourceParameterFactoryInterface interface {
	win32.IInspectableInterface
	Create(name string) *ICompositionEffectSourceParameter
}

type ICompositionEffectSourceParameterFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ICompositionEffectSourceParameterFactory struct {
	win32.IInspectable
}

func (this *ICompositionEffectSourceParameterFactory) Vtbl() *ICompositionEffectSourceParameterFactoryVtbl {
	return (*ICompositionEffectSourceParameterFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionEffectSourceParameterFactory) Create(name string) *ICompositionEffectSourceParameter {
	var _result *ICompositionEffectSourceParameter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4801F884-F6AD-4B93-AFA9-897B64E57B1F
var IID_ICompositionEllipseGeometry = syscall.GUID{0x4801F884, 0xF6AD, 0x4B93,
	[8]byte{0xAF, 0xA9, 0x89, 0x7B, 0x64, 0xE5, 0x7B, 0x1F}}

type ICompositionEllipseGeometryInterface interface {
	win32.IInspectableInterface
	Get_Center() unsafe.Pointer
	Put_Center(value unsafe.Pointer)
	Get_Radius() unsafe.Pointer
	Put_Radius(value unsafe.Pointer)
}

type ICompositionEllipseGeometryVtbl struct {
	win32.IInspectableVtbl
	Get_Center uintptr
	Put_Center uintptr
	Get_Radius uintptr
	Put_Radius uintptr
}

type ICompositionEllipseGeometry struct {
	win32.IInspectable
}

func (this *ICompositionEllipseGeometry) Vtbl() *ICompositionEllipseGeometryVtbl {
	return (*ICompositionEllipseGeometryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionEllipseGeometry) Get_Center() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Center, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionEllipseGeometry) Put_Center(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Center, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionEllipseGeometry) Get_Radius() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Radius, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionEllipseGeometry) Put_Radius(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Radius, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// C840B581-81C9-4444-A2C1-CCAECE3A50E5
var IID_ICompositionGeometricClip = syscall.GUID{0xC840B581, 0x81C9, 0x4444,
	[8]byte{0xA2, 0xC1, 0xCC, 0xAE, 0xCE, 0x3A, 0x50, 0xE5}}

type ICompositionGeometricClipInterface interface {
	win32.IInspectableInterface
	Get_Geometry() *ICompositionGeometry
	Put_Geometry(value *ICompositionGeometry)
	Get_ViewBox() *ICompositionViewBox
	Put_ViewBox(value *ICompositionViewBox)
}

type ICompositionGeometricClipVtbl struct {
	win32.IInspectableVtbl
	Get_Geometry uintptr
	Put_Geometry uintptr
	Get_ViewBox  uintptr
	Put_ViewBox  uintptr
}

type ICompositionGeometricClip struct {
	win32.IInspectable
}

func (this *ICompositionGeometricClip) Vtbl() *ICompositionGeometricClipVtbl {
	return (*ICompositionGeometricClipVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionGeometricClip) Get_Geometry() *ICompositionGeometry {
	var _result *ICompositionGeometry
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Geometry, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionGeometricClip) Put_Geometry(value *ICompositionGeometry) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Geometry, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ICompositionGeometricClip) Get_ViewBox() *ICompositionViewBox {
	var _result *ICompositionViewBox
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ViewBox, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionGeometricClip) Put_ViewBox(value *ICompositionViewBox) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ViewBox, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// E985217C-6A17-4207-ABD8-5FD3DD612A9D
var IID_ICompositionGeometry = syscall.GUID{0xE985217C, 0x6A17, 0x4207,
	[8]byte{0xAB, 0xD8, 0x5F, 0xD3, 0xDD, 0x61, 0x2A, 0x9D}}

type ICompositionGeometryInterface interface {
	win32.IInspectableInterface
	Get_TrimEnd() float32
	Put_TrimEnd(value float32)
	Get_TrimOffset() float32
	Put_TrimOffset(value float32)
	Get_TrimStart() float32
	Put_TrimStart(value float32)
}

type ICompositionGeometryVtbl struct {
	win32.IInspectableVtbl
	Get_TrimEnd    uintptr
	Put_TrimEnd    uintptr
	Get_TrimOffset uintptr
	Put_TrimOffset uintptr
	Get_TrimStart  uintptr
	Put_TrimStart  uintptr
}

type ICompositionGeometry struct {
	win32.IInspectable
}

func (this *ICompositionGeometry) Vtbl() *ICompositionGeometryVtbl {
	return (*ICompositionGeometryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionGeometry) Get_TrimEnd() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TrimEnd, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionGeometry) Put_TrimEnd(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TrimEnd, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionGeometry) Get_TrimOffset() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TrimOffset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionGeometry) Put_TrimOffset(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TrimOffset, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionGeometry) Get_TrimStart() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TrimStart, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionGeometry) Put_TrimStart(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TrimStart, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// BFFEBFE1-8C25-480B-9F56-FED6B288055D
var IID_ICompositionGeometryFactory = syscall.GUID{0xBFFEBFE1, 0x8C25, 0x480B,
	[8]byte{0x9F, 0x56, 0xFE, 0xD6, 0xB2, 0x88, 0x05, 0x5D}}

type ICompositionGeometryFactoryInterface interface {
	win32.IInspectableInterface
}

type ICompositionGeometryFactoryVtbl struct {
	win32.IInspectableVtbl
}

type ICompositionGeometryFactory struct {
	win32.IInspectable
}

func (this *ICompositionGeometryFactory) Vtbl() *ICompositionGeometryFactoryVtbl {
	return (*ICompositionGeometryFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 1D9709E0-FFC6-4C0E-A9AB-34144D4C9098
var IID_ICompositionGradientBrush = syscall.GUID{0x1D9709E0, 0xFFC6, 0x4C0E,
	[8]byte{0xA9, 0xAB, 0x34, 0x14, 0x4D, 0x4C, 0x90, 0x98}}

type ICompositionGradientBrushInterface interface {
	win32.IInspectableInterface
	Get_AnchorPoint() unsafe.Pointer
	Put_AnchorPoint(value unsafe.Pointer)
	Get_CenterPoint() unsafe.Pointer
	Put_CenterPoint(value unsafe.Pointer)
	Get_ColorStops() *ICompositionColorGradientStopCollection
	Get_ExtendMode() CompositionGradientExtendMode
	Put_ExtendMode(value CompositionGradientExtendMode)
	Get_InterpolationSpace() CompositionColorSpace
	Put_InterpolationSpace(value CompositionColorSpace)
	Get_Offset() unsafe.Pointer
	Put_Offset(value unsafe.Pointer)
	Get_RotationAngle() float32
	Put_RotationAngle(value float32)
	Get_RotationAngleInDegrees() float32
	Put_RotationAngleInDegrees(value float32)
	Get_Scale() unsafe.Pointer
	Put_Scale(value unsafe.Pointer)
	Get_TransformMatrix() unsafe.Pointer
	Put_TransformMatrix(value unsafe.Pointer)
}

type ICompositionGradientBrushVtbl struct {
	win32.IInspectableVtbl
	Get_AnchorPoint            uintptr
	Put_AnchorPoint            uintptr
	Get_CenterPoint            uintptr
	Put_CenterPoint            uintptr
	Get_ColorStops             uintptr
	Get_ExtendMode             uintptr
	Put_ExtendMode             uintptr
	Get_InterpolationSpace     uintptr
	Put_InterpolationSpace     uintptr
	Get_Offset                 uintptr
	Put_Offset                 uintptr
	Get_RotationAngle          uintptr
	Put_RotationAngle          uintptr
	Get_RotationAngleInDegrees uintptr
	Put_RotationAngleInDegrees uintptr
	Get_Scale                  uintptr
	Put_Scale                  uintptr
	Get_TransformMatrix        uintptr
	Put_TransformMatrix        uintptr
}

type ICompositionGradientBrush struct {
	win32.IInspectable
}

func (this *ICompositionGradientBrush) Vtbl() *ICompositionGradientBrushVtbl {
	return (*ICompositionGradientBrushVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionGradientBrush) Get_AnchorPoint() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AnchorPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionGradientBrush) Put_AnchorPoint(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AnchorPoint, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionGradientBrush) Get_CenterPoint() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CenterPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionGradientBrush) Put_CenterPoint(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CenterPoint, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionGradientBrush) Get_ColorStops() *ICompositionColorGradientStopCollection {
	var _result *ICompositionColorGradientStopCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ColorStops, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionGradientBrush) Get_ExtendMode() CompositionGradientExtendMode {
	var _result CompositionGradientExtendMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionGradientBrush) Put_ExtendMode(value CompositionGradientExtendMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ExtendMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionGradientBrush) Get_InterpolationSpace() CompositionColorSpace {
	var _result CompositionColorSpace
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InterpolationSpace, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionGradientBrush) Put_InterpolationSpace(value CompositionColorSpace) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InterpolationSpace, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionGradientBrush) Get_Offset() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Offset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionGradientBrush) Put_Offset(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Offset, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionGradientBrush) Get_RotationAngle() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RotationAngle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionGradientBrush) Put_RotationAngle(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RotationAngle, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionGradientBrush) Get_RotationAngleInDegrees() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RotationAngleInDegrees, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionGradientBrush) Put_RotationAngleInDegrees(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RotationAngleInDegrees, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionGradientBrush) Get_Scale() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Scale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionGradientBrush) Put_Scale(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Scale, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionGradientBrush) Get_TransformMatrix() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TransformMatrix, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionGradientBrush) Put_TransformMatrix(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TransformMatrix, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 899DD5A1-B4C7-4B33-A1B6-264ADDC26D10
var IID_ICompositionGradientBrush2 = syscall.GUID{0x899DD5A1, 0xB4C7, 0x4B33,
	[8]byte{0xA1, 0xB6, 0x26, 0x4A, 0xDD, 0xC2, 0x6D, 0x10}}

type ICompositionGradientBrush2Interface interface {
	win32.IInspectableInterface
	Get_MappingMode() CompositionMappingMode
	Put_MappingMode(value CompositionMappingMode)
}

type ICompositionGradientBrush2Vtbl struct {
	win32.IInspectableVtbl
	Get_MappingMode uintptr
	Put_MappingMode uintptr
}

type ICompositionGradientBrush2 struct {
	win32.IInspectable
}

func (this *ICompositionGradientBrush2) Vtbl() *ICompositionGradientBrush2Vtbl {
	return (*ICompositionGradientBrush2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionGradientBrush2) Get_MappingMode() CompositionMappingMode {
	var _result CompositionMappingMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MappingMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionGradientBrush2) Put_MappingMode(value CompositionMappingMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MappingMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 56D765D7-F189-48C9-9C8D-94DAF1BEC010
var IID_ICompositionGradientBrushFactory = syscall.GUID{0x56D765D7, 0xF189, 0x48C9,
	[8]byte{0x9C, 0x8D, 0x94, 0xDA, 0xF1, 0xBE, 0xC0, 0x10}}

type ICompositionGradientBrushFactoryInterface interface {
	win32.IInspectableInterface
}

type ICompositionGradientBrushFactoryVtbl struct {
	win32.IInspectableVtbl
}

type ICompositionGradientBrushFactory struct {
	win32.IInspectable
}

func (this *ICompositionGradientBrushFactory) Vtbl() *ICompositionGradientBrushFactoryVtbl {
	return (*ICompositionGradientBrushFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// FB22C6E1-80A2-4667-9936-DBEAF6EEFE95
var IID_ICompositionGraphicsDevice = syscall.GUID{0xFB22C6E1, 0x80A2, 0x4667,
	[8]byte{0x99, 0x36, 0xDB, 0xEA, 0xF6, 0xEE, 0xFE, 0x95}}

type ICompositionGraphicsDeviceInterface interface {
	win32.IInspectableInterface
	CreateDrawingSurface(sizePixels Size, pixelFormat unsafe.Pointer, alphaMode unsafe.Pointer) *ICompositionDrawingSurface
	Add_RenderingDeviceReplaced(handler TypedEventHandler[*ICompositionGraphicsDevice, *IRenderingDeviceReplacedEventArgs]) EventRegistrationToken
	Remove_RenderingDeviceReplaced(token EventRegistrationToken)
}

type ICompositionGraphicsDeviceVtbl struct {
	win32.IInspectableVtbl
	CreateDrawingSurface           uintptr
	Add_RenderingDeviceReplaced    uintptr
	Remove_RenderingDeviceReplaced uintptr
}

type ICompositionGraphicsDevice struct {
	win32.IInspectable
}

func (this *ICompositionGraphicsDevice) Vtbl() *ICompositionGraphicsDeviceVtbl {
	return (*ICompositionGraphicsDeviceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionGraphicsDevice) CreateDrawingSurface(sizePixels Size, pixelFormat unsafe.Pointer, alphaMode unsafe.Pointer) *ICompositionDrawingSurface {
	var _result *ICompositionDrawingSurface
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDrawingSurface, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&sizePixels)), uintptr(pixelFormat), uintptr(alphaMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionGraphicsDevice) Add_RenderingDeviceReplaced(handler TypedEventHandler[*ICompositionGraphicsDevice, *IRenderingDeviceReplacedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_RenderingDeviceReplaced, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionGraphicsDevice) Remove_RenderingDeviceReplaced(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_RenderingDeviceReplaced, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 0FB8BDF6-C0F0-4BCC-9FB8-084982490D7D
var IID_ICompositionGraphicsDevice2 = syscall.GUID{0x0FB8BDF6, 0xC0F0, 0x4BCC,
	[8]byte{0x9F, 0xB8, 0x08, 0x49, 0x82, 0x49, 0x0D, 0x7D}}

type ICompositionGraphicsDevice2Interface interface {
	win32.IInspectableInterface
	CreateDrawingSurface2(sizePixels unsafe.Pointer, pixelFormat unsafe.Pointer, alphaMode unsafe.Pointer) *ICompositionDrawingSurface
	CreateVirtualDrawingSurface(sizePixels unsafe.Pointer, pixelFormat unsafe.Pointer, alphaMode unsafe.Pointer) *ICompositionVirtualDrawingSurface
}

type ICompositionGraphicsDevice2Vtbl struct {
	win32.IInspectableVtbl
	CreateDrawingSurface2       uintptr
	CreateVirtualDrawingSurface uintptr
}

type ICompositionGraphicsDevice2 struct {
	win32.IInspectable
}

func (this *ICompositionGraphicsDevice2) Vtbl() *ICompositionGraphicsDevice2Vtbl {
	return (*ICompositionGraphicsDevice2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionGraphicsDevice2) CreateDrawingSurface2(sizePixels unsafe.Pointer, pixelFormat unsafe.Pointer, alphaMode unsafe.Pointer) *ICompositionDrawingSurface {
	var _result *ICompositionDrawingSurface
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDrawingSurface2, uintptr(unsafe.Pointer(this)), uintptr(sizePixels), uintptr(pixelFormat), uintptr(alphaMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionGraphicsDevice2) CreateVirtualDrawingSurface(sizePixels unsafe.Pointer, pixelFormat unsafe.Pointer, alphaMode unsafe.Pointer) *ICompositionVirtualDrawingSurface {
	var _result *ICompositionVirtualDrawingSurface
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateVirtualDrawingSurface, uintptr(unsafe.Pointer(this)), uintptr(sizePixels), uintptr(pixelFormat), uintptr(alphaMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 37F67514-D3EF-49D1-B69D-0D8EABEB3626
var IID_ICompositionGraphicsDevice3 = syscall.GUID{0x37F67514, 0xD3EF, 0x49D1,
	[8]byte{0xB6, 0x9D, 0x0D, 0x8E, 0xAB, 0xEB, 0x36, 0x26}}

type ICompositionGraphicsDevice3Interface interface {
	win32.IInspectableInterface
	CreateMipmapSurface(sizePixels unsafe.Pointer, pixelFormat unsafe.Pointer, alphaMode unsafe.Pointer) *ICompositionMipmapSurface
	Trim()
}

type ICompositionGraphicsDevice3Vtbl struct {
	win32.IInspectableVtbl
	CreateMipmapSurface uintptr
	Trim                uintptr
}

type ICompositionGraphicsDevice3 struct {
	win32.IInspectable
}

func (this *ICompositionGraphicsDevice3) Vtbl() *ICompositionGraphicsDevice3Vtbl {
	return (*ICompositionGraphicsDevice3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionGraphicsDevice3) CreateMipmapSurface(sizePixels unsafe.Pointer, pixelFormat unsafe.Pointer, alphaMode unsafe.Pointer) *ICompositionMipmapSurface {
	var _result *ICompositionMipmapSurface
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateMipmapSurface, uintptr(unsafe.Pointer(this)), uintptr(sizePixels), uintptr(pixelFormat), uintptr(alphaMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionGraphicsDevice3) Trim() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Trim, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 5A73BFF9-A97F-4CF5-BA46-98EF358E71B1
var IID_ICompositionGraphicsDevice4 = syscall.GUID{0x5A73BFF9, 0xA97F, 0x4CF5,
	[8]byte{0xBA, 0x46, 0x98, 0xEF, 0x35, 0x8E, 0x71, 0xB1}}

type ICompositionGraphicsDevice4Interface interface {
	win32.IInspectableInterface
	CaptureAsync(captureVisual *IVisual, size unsafe.Pointer, pixelFormat unsafe.Pointer, alphaMode unsafe.Pointer, sdrBoost float32) *IAsyncOperation[*ICompositionSurface]
}

type ICompositionGraphicsDevice4Vtbl struct {
	win32.IInspectableVtbl
	CaptureAsync uintptr
}

type ICompositionGraphicsDevice4 struct {
	win32.IInspectable
}

func (this *ICompositionGraphicsDevice4) Vtbl() *ICompositionGraphicsDevice4Vtbl {
	return (*ICompositionGraphicsDevice4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionGraphicsDevice4) CaptureAsync(captureVisual *IVisual, size unsafe.Pointer, pixelFormat unsafe.Pointer, alphaMode unsafe.Pointer, sdrBoost float32) *IAsyncOperation[*ICompositionSurface] {
	var _result *IAsyncOperation[*ICompositionSurface]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CaptureAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(captureVisual)), uintptr(size), uintptr(pixelFormat), uintptr(alphaMode), uintptr(sdrBoost), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 41A6D7C2-2E5D-4BC1-B09E-8F0A03E3D8D3
var IID_ICompositionLight = syscall.GUID{0x41A6D7C2, 0x2E5D, 0x4BC1,
	[8]byte{0xB0, 0x9E, 0x8F, 0x0A, 0x03, 0xE3, 0xD8, 0xD3}}

type ICompositionLightInterface interface {
	win32.IInspectableInterface
	Get_Targets() *IVisualUnorderedCollection
}

type ICompositionLightVtbl struct {
	win32.IInspectableVtbl
	Get_Targets uintptr
}

type ICompositionLight struct {
	win32.IInspectable
}

func (this *ICompositionLight) Vtbl() *ICompositionLightVtbl {
	return (*ICompositionLightVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionLight) Get_Targets() *IVisualUnorderedCollection {
	var _result *IVisualUnorderedCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Targets, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A7BCDA72-F35D-425D-9B98-23F4205F6669
var IID_ICompositionLight2 = syscall.GUID{0xA7BCDA72, 0xF35D, 0x425D,
	[8]byte{0x9B, 0x98, 0x23, 0xF4, 0x20, 0x5F, 0x66, 0x69}}

type ICompositionLight2Interface interface {
	win32.IInspectableInterface
	Get_ExclusionsFromTargets() *IVisualUnorderedCollection
}

type ICompositionLight2Vtbl struct {
	win32.IInspectableVtbl
	Get_ExclusionsFromTargets uintptr
}

type ICompositionLight2 struct {
	win32.IInspectable
}

func (this *ICompositionLight2) Vtbl() *ICompositionLight2Vtbl {
	return (*ICompositionLight2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionLight2) Get_ExclusionsFromTargets() *IVisualUnorderedCollection {
	var _result *IVisualUnorderedCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExclusionsFromTargets, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4B0B00E4-DF07-4959-B7A4-4F7E4233F838
var IID_ICompositionLight3 = syscall.GUID{0x4B0B00E4, 0xDF07, 0x4959,
	[8]byte{0xB7, 0xA4, 0x4F, 0x7E, 0x42, 0x33, 0xF8, 0x38}}

type ICompositionLight3Interface interface {
	win32.IInspectableInterface
	Get_IsEnabled() bool
	Put_IsEnabled(value bool)
}

type ICompositionLight3Vtbl struct {
	win32.IInspectableVtbl
	Get_IsEnabled uintptr
	Put_IsEnabled uintptr
}

type ICompositionLight3 struct {
	win32.IInspectable
}

func (this *ICompositionLight3) Vtbl() *ICompositionLight3Vtbl {
	return (*ICompositionLight3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionLight3) Get_IsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionLight3) Put_IsEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 069CF306-DA3C-4B44-838A-5E03D51ACE55
var IID_ICompositionLightFactory = syscall.GUID{0x069CF306, 0xDA3C, 0x4B44,
	[8]byte{0x83, 0x8A, 0x5E, 0x03, 0xD5, 0x1A, 0xCE, 0x55}}

type ICompositionLightFactoryInterface interface {
	win32.IInspectableInterface
}

type ICompositionLightFactoryVtbl struct {
	win32.IInspectableVtbl
}

type ICompositionLightFactory struct {
	win32.IInspectable
}

func (this *ICompositionLightFactory) Vtbl() *ICompositionLightFactoryVtbl {
	return (*ICompositionLightFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// DD7615A4-0C9A-4B67-8DCE-440A5BF9CDEC
var IID_ICompositionLineGeometry = syscall.GUID{0xDD7615A4, 0x0C9A, 0x4B67,
	[8]byte{0x8D, 0xCE, 0x44, 0x0A, 0x5B, 0xF9, 0xCD, 0xEC}}

type ICompositionLineGeometryInterface interface {
	win32.IInspectableInterface
	Get_Start() unsafe.Pointer
	Put_Start(value unsafe.Pointer)
	Get_End() unsafe.Pointer
	Put_End(value unsafe.Pointer)
}

type ICompositionLineGeometryVtbl struct {
	win32.IInspectableVtbl
	Get_Start uintptr
	Put_Start uintptr
	Get_End   uintptr
	Put_End   uintptr
}

type ICompositionLineGeometry struct {
	win32.IInspectable
}

func (this *ICompositionLineGeometry) Vtbl() *ICompositionLineGeometryVtbl {
	return (*ICompositionLineGeometryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionLineGeometry) Get_Start() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Start, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionLineGeometry) Put_Start(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Start, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionLineGeometry) Get_End() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_End, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionLineGeometry) Put_End(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_End, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 983BC519-A9DB-413C-A2D8-2A9056FC525E
var IID_ICompositionLinearGradientBrush = syscall.GUID{0x983BC519, 0xA9DB, 0x413C,
	[8]byte{0xA2, 0xD8, 0x2A, 0x90, 0x56, 0xFC, 0x52, 0x5E}}

type ICompositionLinearGradientBrushInterface interface {
	win32.IInspectableInterface
	Get_EndPoint() unsafe.Pointer
	Put_EndPoint(value unsafe.Pointer)
	Get_StartPoint() unsafe.Pointer
	Put_StartPoint(value unsafe.Pointer)
}

type ICompositionLinearGradientBrushVtbl struct {
	win32.IInspectableVtbl
	Get_EndPoint   uintptr
	Put_EndPoint   uintptr
	Get_StartPoint uintptr
	Put_StartPoint uintptr
}

type ICompositionLinearGradientBrush struct {
	win32.IInspectable
}

func (this *ICompositionLinearGradientBrush) Vtbl() *ICompositionLinearGradientBrushVtbl {
	return (*ICompositionLinearGradientBrushVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionLinearGradientBrush) Get_EndPoint() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EndPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionLinearGradientBrush) Put_EndPoint(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_EndPoint, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionLinearGradientBrush) Get_StartPoint() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StartPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionLinearGradientBrush) Put_StartPoint(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StartPoint, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 522CF09E-BE6B-4F41-BE49-F9226D471B4A
var IID_ICompositionMaskBrush = syscall.GUID{0x522CF09E, 0xBE6B, 0x4F41,
	[8]byte{0xBE, 0x49, 0xF9, 0x22, 0x6D, 0x47, 0x1B, 0x4A}}

type ICompositionMaskBrushInterface interface {
	win32.IInspectableInterface
	Get_Mask() *ICompositionBrush
	Put_Mask(value *ICompositionBrush)
	Get_Source() *ICompositionBrush
	Put_Source(value *ICompositionBrush)
}

type ICompositionMaskBrushVtbl struct {
	win32.IInspectableVtbl
	Get_Mask   uintptr
	Put_Mask   uintptr
	Get_Source uintptr
	Put_Source uintptr
}

type ICompositionMaskBrush struct {
	win32.IInspectable
}

func (this *ICompositionMaskBrush) Vtbl() *ICompositionMaskBrushVtbl {
	return (*ICompositionMaskBrushVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionMaskBrush) Get_Mask() *ICompositionBrush {
	var _result *ICompositionBrush
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mask, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionMaskBrush) Put_Mask(value *ICompositionBrush) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Mask, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ICompositionMaskBrush) Get_Source() *ICompositionBrush {
	var _result *ICompositionBrush
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Source, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionMaskBrush) Put_Source(value *ICompositionBrush) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Source, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 4863675C-CF4A-4B1C-9ECE-C5EC0C2B2FE6
var IID_ICompositionMipmapSurface = syscall.GUID{0x4863675C, 0xCF4A, 0x4B1C,
	[8]byte{0x9E, 0xCE, 0xC5, 0xEC, 0x0C, 0x2B, 0x2F, 0xE6}}

type ICompositionMipmapSurfaceInterface interface {
	win32.IInspectableInterface
	Get_LevelCount() uint32
	Get_AlphaMode() unsafe.Pointer
	Get_PixelFormat() unsafe.Pointer
	Get_SizeInt32() unsafe.Pointer
	GetDrawingSurfaceForLevel(level uint32) *ICompositionDrawingSurface
}

type ICompositionMipmapSurfaceVtbl struct {
	win32.IInspectableVtbl
	Get_LevelCount            uintptr
	Get_AlphaMode             uintptr
	Get_PixelFormat           uintptr
	Get_SizeInt32             uintptr
	GetDrawingSurfaceForLevel uintptr
}

type ICompositionMipmapSurface struct {
	win32.IInspectable
}

func (this *ICompositionMipmapSurface) Vtbl() *ICompositionMipmapSurfaceVtbl {
	return (*ICompositionMipmapSurfaceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionMipmapSurface) Get_LevelCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LevelCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionMipmapSurface) Get_AlphaMode() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlphaMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionMipmapSurface) Get_PixelFormat() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PixelFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionMipmapSurface) Get_SizeInt32() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SizeInt32, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionMipmapSurface) GetDrawingSurfaceForLevel(level uint32) *ICompositionDrawingSurface {
	var _result *ICompositionDrawingSurface
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDrawingSurfaceForLevel, uintptr(unsafe.Pointer(this)), uintptr(level), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F25154E4-BC8C-4BE7-B80F-8685B83C0186
var IID_ICompositionNineGridBrush = syscall.GUID{0xF25154E4, 0xBC8C, 0x4BE7,
	[8]byte{0xB8, 0x0F, 0x86, 0x85, 0xB8, 0x3C, 0x01, 0x86}}

type ICompositionNineGridBrushInterface interface {
	win32.IInspectableInterface
	Get_BottomInset() float32
	Put_BottomInset(value float32)
	Get_BottomInsetScale() float32
	Put_BottomInsetScale(value float32)
	Get_IsCenterHollow() bool
	Put_IsCenterHollow(value bool)
	Get_LeftInset() float32
	Put_LeftInset(value float32)
	Get_LeftInsetScale() float32
	Put_LeftInsetScale(value float32)
	Get_RightInset() float32
	Put_RightInset(value float32)
	Get_RightInsetScale() float32
	Put_RightInsetScale(value float32)
	Get_Source() *ICompositionBrush
	Put_Source(value *ICompositionBrush)
	Get_TopInset() float32
	Put_TopInset(value float32)
	Get_TopInsetScale() float32
	Put_TopInsetScale(value float32)
	SetInsets(inset float32)
	SetInsetsWithValues(left float32, top float32, right float32, bottom float32)
	SetInsetScales(scale float32)
	SetInsetScalesWithValues(left float32, top float32, right float32, bottom float32)
}

type ICompositionNineGridBrushVtbl struct {
	win32.IInspectableVtbl
	Get_BottomInset          uintptr
	Put_BottomInset          uintptr
	Get_BottomInsetScale     uintptr
	Put_BottomInsetScale     uintptr
	Get_IsCenterHollow       uintptr
	Put_IsCenterHollow       uintptr
	Get_LeftInset            uintptr
	Put_LeftInset            uintptr
	Get_LeftInsetScale       uintptr
	Put_LeftInsetScale       uintptr
	Get_RightInset           uintptr
	Put_RightInset           uintptr
	Get_RightInsetScale      uintptr
	Put_RightInsetScale      uintptr
	Get_Source               uintptr
	Put_Source               uintptr
	Get_TopInset             uintptr
	Put_TopInset             uintptr
	Get_TopInsetScale        uintptr
	Put_TopInsetScale        uintptr
	SetInsets                uintptr
	SetInsetsWithValues      uintptr
	SetInsetScales           uintptr
	SetInsetScalesWithValues uintptr
}

type ICompositionNineGridBrush struct {
	win32.IInspectable
}

func (this *ICompositionNineGridBrush) Vtbl() *ICompositionNineGridBrushVtbl {
	return (*ICompositionNineGridBrushVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionNineGridBrush) Get_BottomInset() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BottomInset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionNineGridBrush) Put_BottomInset(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BottomInset, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionNineGridBrush) Get_BottomInsetScale() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BottomInsetScale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionNineGridBrush) Put_BottomInsetScale(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BottomInsetScale, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionNineGridBrush) Get_IsCenterHollow() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCenterHollow, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionNineGridBrush) Put_IsCenterHollow(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsCenterHollow, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICompositionNineGridBrush) Get_LeftInset() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LeftInset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionNineGridBrush) Put_LeftInset(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LeftInset, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionNineGridBrush) Get_LeftInsetScale() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LeftInsetScale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionNineGridBrush) Put_LeftInsetScale(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LeftInsetScale, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionNineGridBrush) Get_RightInset() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RightInset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionNineGridBrush) Put_RightInset(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RightInset, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionNineGridBrush) Get_RightInsetScale() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RightInsetScale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionNineGridBrush) Put_RightInsetScale(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RightInsetScale, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionNineGridBrush) Get_Source() *ICompositionBrush {
	var _result *ICompositionBrush
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Source, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionNineGridBrush) Put_Source(value *ICompositionBrush) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Source, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ICompositionNineGridBrush) Get_TopInset() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TopInset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionNineGridBrush) Put_TopInset(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TopInset, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionNineGridBrush) Get_TopInsetScale() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TopInsetScale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionNineGridBrush) Put_TopInsetScale(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TopInsetScale, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionNineGridBrush) SetInsets(inset float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetInsets, uintptr(unsafe.Pointer(this)), uintptr(inset))
	_ = _hr
}

func (this *ICompositionNineGridBrush) SetInsetsWithValues(left float32, top float32, right float32, bottom float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetInsetsWithValues, uintptr(unsafe.Pointer(this)), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	_ = _hr
}

func (this *ICompositionNineGridBrush) SetInsetScales(scale float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetInsetScales, uintptr(unsafe.Pointer(this)), uintptr(scale))
	_ = _hr
}

func (this *ICompositionNineGridBrush) SetInsetScalesWithValues(left float32, top float32, right float32, bottom float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetInsetScalesWithValues, uintptr(unsafe.Pointer(this)), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	_ = _hr
}

// BCB4AD45-7609-4550-934F-16002A68FDED
var IID_ICompositionObject = syscall.GUID{0xBCB4AD45, 0x7609, 0x4550,
	[8]byte{0x93, 0x4F, 0x16, 0x00, 0x2A, 0x68, 0xFD, 0xED}}

type ICompositionObjectInterface interface {
	win32.IInspectableInterface
	Get_Compositor() *ICompositor
	Get_Dispatcher() unsafe.Pointer
	Get_Properties() *ICompositionPropertySet
	StartAnimation(propertyName string, animation *ICompositionAnimation)
	StopAnimation(propertyName string)
}

type ICompositionObjectVtbl struct {
	win32.IInspectableVtbl
	Get_Compositor uintptr
	Get_Dispatcher uintptr
	Get_Properties uintptr
	StartAnimation uintptr
	StopAnimation  uintptr
}

type ICompositionObject struct {
	win32.IInspectable
}

func (this *ICompositionObject) Vtbl() *ICompositionObjectVtbl {
	return (*ICompositionObjectVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionObject) Get_Compositor() *ICompositor {
	var _result *ICompositor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Compositor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionObject) Get_Dispatcher() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Dispatcher, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionObject) Get_Properties() *ICompositionPropertySet {
	var _result *ICompositionPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionObject) StartAnimation(propertyName string, animation *ICompositionAnimation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartAnimation, uintptr(unsafe.Pointer(this)), NewHStr(propertyName).Ptr, uintptr(unsafe.Pointer(animation)))
	_ = _hr
}

func (this *ICompositionObject) StopAnimation(propertyName string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopAnimation, uintptr(unsafe.Pointer(this)), NewHStr(propertyName).Ptr)
	_ = _hr
}

// EF874EA1-5CFF-4B68-9E30-A1519D08BA03
var IID_ICompositionObject2 = syscall.GUID{0xEF874EA1, 0x5CFF, 0x4B68,
	[8]byte{0x9E, 0x30, 0xA1, 0x51, 0x9D, 0x08, 0xBA, 0x03}}

type ICompositionObject2Interface interface {
	win32.IInspectableInterface
	Get_Comment() string
	Put_Comment(value string)
	Get_ImplicitAnimations() *IImplicitAnimationCollection
	Put_ImplicitAnimations(value *IImplicitAnimationCollection)
	StartAnimationGroup(value *ICompositionAnimationBase)
	StopAnimationGroup(value *ICompositionAnimationBase)
}

type ICompositionObject2Vtbl struct {
	win32.IInspectableVtbl
	Get_Comment            uintptr
	Put_Comment            uintptr
	Get_ImplicitAnimations uintptr
	Put_ImplicitAnimations uintptr
	StartAnimationGroup    uintptr
	StopAnimationGroup     uintptr
}

type ICompositionObject2 struct {
	win32.IInspectable
}

func (this *ICompositionObject2) Vtbl() *ICompositionObject2Vtbl {
	return (*ICompositionObject2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionObject2) Get_Comment() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Comment, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICompositionObject2) Put_Comment(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Comment, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICompositionObject2) Get_ImplicitAnimations() *IImplicitAnimationCollection {
	var _result *IImplicitAnimationCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ImplicitAnimations, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionObject2) Put_ImplicitAnimations(value *IImplicitAnimationCollection) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ImplicitAnimations, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ICompositionObject2) StartAnimationGroup(value *ICompositionAnimationBase) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartAnimationGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ICompositionObject2) StopAnimationGroup(value *ICompositionAnimationBase) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopAnimationGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 4BC27925-DACD-4CF2-98B1-986B76E7EBE6
var IID_ICompositionObject3 = syscall.GUID{0x4BC27925, 0xDACD, 0x4CF2,
	[8]byte{0x98, 0xB1, 0x98, 0x6B, 0x76, 0xE7, 0xEB, 0xE6}}

type ICompositionObject3Interface interface {
	win32.IInspectableInterface
	Get_DispatcherQueue() *IDispatcherQueue
}

type ICompositionObject3Vtbl struct {
	win32.IInspectableVtbl
	Get_DispatcherQueue uintptr
}

type ICompositionObject3 struct {
	win32.IInspectable
}

func (this *ICompositionObject3) Vtbl() *ICompositionObject3Vtbl {
	return (*ICompositionObject3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionObject3) Get_DispatcherQueue() *IDispatcherQueue {
	var _result *IDispatcherQueue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DispatcherQueue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0BB3784C-346B-4A7C-966B-7310966553D5
var IID_ICompositionObject4 = syscall.GUID{0x0BB3784C, 0x346B, 0x4A7C,
	[8]byte{0x96, 0x6B, 0x73, 0x10, 0x96, 0x65, 0x53, 0xD5}}

type ICompositionObject4Interface interface {
	win32.IInspectableInterface
	TryGetAnimationController(propertyName string) *IAnimationController
}

type ICompositionObject4Vtbl struct {
	win32.IInspectableVtbl
	TryGetAnimationController uintptr
}

type ICompositionObject4 struct {
	win32.IInspectable
}

func (this *ICompositionObject4) Vtbl() *ICompositionObject4Vtbl {
	return (*ICompositionObject4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionObject4) TryGetAnimationController(propertyName string) *IAnimationController {
	var _result *IAnimationController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetAnimationController, uintptr(unsafe.Pointer(this)), NewHStr(propertyName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1D7F391B-A130-5265-A62B-60B8E668965A
var IID_ICompositionObject5 = syscall.GUID{0x1D7F391B, 0xA130, 0x5265,
	[8]byte{0xA6, 0x2B, 0x60, 0xB8, 0xE6, 0x68, 0x96, 0x5A}}

type ICompositionObject5Interface interface {
	win32.IInspectableInterface
	StartAnimationWithController(propertyName string, animation *ICompositionAnimation, animationController *IAnimationController)
}

type ICompositionObject5Vtbl struct {
	win32.IInspectableVtbl
	StartAnimationWithController uintptr
}

type ICompositionObject5 struct {
	win32.IInspectable
}

func (this *ICompositionObject5) Vtbl() *ICompositionObject5Vtbl {
	return (*ICompositionObject5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionObject5) StartAnimationWithController(propertyName string, animation *ICompositionAnimation, animationController *IAnimationController) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartAnimationWithController, uintptr(unsafe.Pointer(this)), NewHStr(propertyName).Ptr, uintptr(unsafe.Pointer(animation)), uintptr(unsafe.Pointer(animationController)))
	_ = _hr
}

// 51205C5E-558A-4F2A-8D39-37BFE1E20DDD
var IID_ICompositionObjectFactory = syscall.GUID{0x51205C5E, 0x558A, 0x4F2A,
	[8]byte{0x8D, 0x39, 0x37, 0xBF, 0xE1, 0xE2, 0x0D, 0xDD}}

type ICompositionObjectFactoryInterface interface {
	win32.IInspectableInterface
}

type ICompositionObjectFactoryVtbl struct {
	win32.IInspectableVtbl
}

type ICompositionObjectFactory struct {
	win32.IInspectable
}

func (this *ICompositionObjectFactory) Vtbl() *ICompositionObjectFactoryVtbl {
	return (*ICompositionObjectFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// C1ED052F-1BA2-44BA-A904-6A882A0A5ADB
var IID_ICompositionObjectStatics = syscall.GUID{0xC1ED052F, 0x1BA2, 0x44BA,
	[8]byte{0xA9, 0x04, 0x6A, 0x88, 0x2A, 0x0A, 0x5A, 0xDB}}

type ICompositionObjectStaticsInterface interface {
	win32.IInspectableInterface
	StartAnimationWithIAnimationObject(target *IAnimationObject, propertyName string, animation *ICompositionAnimation)
	StartAnimationGroupWithIAnimationObject(target *IAnimationObject, animation *ICompositionAnimationBase)
}

type ICompositionObjectStaticsVtbl struct {
	win32.IInspectableVtbl
	StartAnimationWithIAnimationObject      uintptr
	StartAnimationGroupWithIAnimationObject uintptr
}

type ICompositionObjectStatics struct {
	win32.IInspectable
}

func (this *ICompositionObjectStatics) Vtbl() *ICompositionObjectStaticsVtbl {
	return (*ICompositionObjectStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionObjectStatics) StartAnimationWithIAnimationObject(target *IAnimationObject, propertyName string, animation *ICompositionAnimation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartAnimationWithIAnimationObject, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(target)), NewHStr(propertyName).Ptr, uintptr(unsafe.Pointer(animation)))
	_ = _hr
}

func (this *ICompositionObjectStatics) StartAnimationGroupWithIAnimationObject(target *IAnimationObject, animation *ICompositionAnimationBase) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartAnimationGroupWithIAnimationObject, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(target)), uintptr(unsafe.Pointer(animation)))
	_ = _hr
}

// 66DA1D5F-2E10-4F22-8A06-0A8151919E60
var IID_ICompositionPath = syscall.GUID{0x66DA1D5F, 0x2E10, 0x4F22,
	[8]byte{0x8A, 0x06, 0x0A, 0x81, 0x51, 0x91, 0x9E, 0x60}}

type ICompositionPathInterface interface {
	win32.IInspectableInterface
}

type ICompositionPathVtbl struct {
	win32.IInspectableVtbl
}

type ICompositionPath struct {
	win32.IInspectable
}

func (this *ICompositionPath) Vtbl() *ICompositionPathVtbl {
	return (*ICompositionPathVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 9C1E8C6A-0F33-4751-9437-EB3FB9D3AB07
var IID_ICompositionPathFactory = syscall.GUID{0x9C1E8C6A, 0x0F33, 0x4751,
	[8]byte{0x94, 0x37, 0xEB, 0x3F, 0xB9, 0xD3, 0xAB, 0x07}}

type ICompositionPathFactoryInterface interface {
	win32.IInspectableInterface
	Create(source unsafe.Pointer) *ICompositionPath
}

type ICompositionPathFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ICompositionPathFactory struct {
	win32.IInspectable
}

func (this *ICompositionPathFactory) Vtbl() *ICompositionPathFactoryVtbl {
	return (*ICompositionPathFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionPathFactory) Create(source unsafe.Pointer) *ICompositionPath {
	var _result *ICompositionPath
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(source), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0B6A417E-2C77-4C23-AF5E-6304C147BB61
var IID_ICompositionPathGeometry = syscall.GUID{0x0B6A417E, 0x2C77, 0x4C23,
	[8]byte{0xAF, 0x5E, 0x63, 0x04, 0xC1, 0x47, 0xBB, 0x61}}

type ICompositionPathGeometryInterface interface {
	win32.IInspectableInterface
	Get_Path() *ICompositionPath
	Put_Path(value *ICompositionPath)
}

type ICompositionPathGeometryVtbl struct {
	win32.IInspectableVtbl
	Get_Path uintptr
	Put_Path uintptr
}

type ICompositionPathGeometry struct {
	win32.IInspectable
}

func (this *ICompositionPathGeometry) Vtbl() *ICompositionPathGeometryVtbl {
	return (*ICompositionPathGeometryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionPathGeometry) Get_Path() *ICompositionPath {
	var _result *ICompositionPath
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Path, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionPathGeometry) Put_Path(value *ICompositionPath) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Path, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 285B8E72-4328-523F-BCF2-5557C52C3B25
var IID_ICompositionProjectedShadow = syscall.GUID{0x285B8E72, 0x4328, 0x523F,
	[8]byte{0xBC, 0xF2, 0x55, 0x57, 0xC5, 0x2C, 0x3B, 0x25}}

type ICompositionProjectedShadowInterface interface {
	win32.IInspectableInterface
	Get_BlurRadiusMultiplier() float32
	Put_BlurRadiusMultiplier(value float32)
	Get_Casters() *ICompositionProjectedShadowCasterCollection
	Get_LightSource() *ICompositionLight
	Put_LightSource(value *ICompositionLight)
	Get_MaxBlurRadius() float32
	Put_MaxBlurRadius(value float32)
	Get_MinBlurRadius() float32
	Put_MinBlurRadius(value float32)
	Get_Receivers() *ICompositionProjectedShadowReceiverUnorderedCollection
}

type ICompositionProjectedShadowVtbl struct {
	win32.IInspectableVtbl
	Get_BlurRadiusMultiplier uintptr
	Put_BlurRadiusMultiplier uintptr
	Get_Casters              uintptr
	Get_LightSource          uintptr
	Put_LightSource          uintptr
	Get_MaxBlurRadius        uintptr
	Put_MaxBlurRadius        uintptr
	Get_MinBlurRadius        uintptr
	Put_MinBlurRadius        uintptr
	Get_Receivers            uintptr
}

type ICompositionProjectedShadow struct {
	win32.IInspectable
}

func (this *ICompositionProjectedShadow) Vtbl() *ICompositionProjectedShadowVtbl {
	return (*ICompositionProjectedShadowVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionProjectedShadow) Get_BlurRadiusMultiplier() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BlurRadiusMultiplier, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionProjectedShadow) Put_BlurRadiusMultiplier(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BlurRadiusMultiplier, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionProjectedShadow) Get_Casters() *ICompositionProjectedShadowCasterCollection {
	var _result *ICompositionProjectedShadowCasterCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Casters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionProjectedShadow) Get_LightSource() *ICompositionLight {
	var _result *ICompositionLight
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LightSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionProjectedShadow) Put_LightSource(value *ICompositionLight) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LightSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ICompositionProjectedShadow) Get_MaxBlurRadius() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxBlurRadius, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionProjectedShadow) Put_MaxBlurRadius(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxBlurRadius, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionProjectedShadow) Get_MinBlurRadius() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinBlurRadius, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionProjectedShadow) Put_MinBlurRadius(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MinBlurRadius, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionProjectedShadow) Get_Receivers() *ICompositionProjectedShadowReceiverUnorderedCollection {
	var _result *ICompositionProjectedShadowReceiverUnorderedCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Receivers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B1D7D426-1E36-5A62-BE56-A16112FDD148
var IID_ICompositionProjectedShadowCaster = syscall.GUID{0xB1D7D426, 0x1E36, 0x5A62,
	[8]byte{0xBE, 0x56, 0xA1, 0x61, 0x12, 0xFD, 0xD1, 0x48}}

type ICompositionProjectedShadowCasterInterface interface {
	win32.IInspectableInterface
	Get_Brush() *ICompositionBrush
	Put_Brush(value *ICompositionBrush)
	Get_CastingVisual() *IVisual
	Put_CastingVisual(value *IVisual)
}

type ICompositionProjectedShadowCasterVtbl struct {
	win32.IInspectableVtbl
	Get_Brush         uintptr
	Put_Brush         uintptr
	Get_CastingVisual uintptr
	Put_CastingVisual uintptr
}

type ICompositionProjectedShadowCaster struct {
	win32.IInspectable
}

func (this *ICompositionProjectedShadowCaster) Vtbl() *ICompositionProjectedShadowCasterVtbl {
	return (*ICompositionProjectedShadowCasterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionProjectedShadowCaster) Get_Brush() *ICompositionBrush {
	var _result *ICompositionBrush
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Brush, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionProjectedShadowCaster) Put_Brush(value *ICompositionBrush) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Brush, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ICompositionProjectedShadowCaster) Get_CastingVisual() *IVisual {
	var _result *IVisual
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CastingVisual, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionProjectedShadowCaster) Put_CastingVisual(value *IVisual) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CastingVisual, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// D2525C0C-E07F-58A3-AC91-37F73EE91740
var IID_ICompositionProjectedShadowCasterCollection = syscall.GUID{0xD2525C0C, 0xE07F, 0x58A3,
	[8]byte{0xAC, 0x91, 0x37, 0xF7, 0x3E, 0xE9, 0x17, 0x40}}

type ICompositionProjectedShadowCasterCollectionInterface interface {
	win32.IInspectableInterface
	Get_Count() int32
	InsertAbove(newCaster *ICompositionProjectedShadowCaster, reference *ICompositionProjectedShadowCaster)
	InsertAtBottom(newCaster *ICompositionProjectedShadowCaster)
	InsertAtTop(newCaster *ICompositionProjectedShadowCaster)
	InsertBelow(newCaster *ICompositionProjectedShadowCaster, reference *ICompositionProjectedShadowCaster)
	Remove(caster *ICompositionProjectedShadowCaster)
	RemoveAll()
}

type ICompositionProjectedShadowCasterCollectionVtbl struct {
	win32.IInspectableVtbl
	Get_Count      uintptr
	InsertAbove    uintptr
	InsertAtBottom uintptr
	InsertAtTop    uintptr
	InsertBelow    uintptr
	Remove         uintptr
	RemoveAll      uintptr
}

type ICompositionProjectedShadowCasterCollection struct {
	win32.IInspectable
}

func (this *ICompositionProjectedShadowCasterCollection) Vtbl() *ICompositionProjectedShadowCasterCollectionVtbl {
	return (*ICompositionProjectedShadowCasterCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionProjectedShadowCasterCollection) Get_Count() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Count, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionProjectedShadowCasterCollection) InsertAbove(newCaster *ICompositionProjectedShadowCaster, reference *ICompositionProjectedShadowCaster) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertAbove, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(newCaster)), uintptr(unsafe.Pointer(reference)))
	_ = _hr
}

func (this *ICompositionProjectedShadowCasterCollection) InsertAtBottom(newCaster *ICompositionProjectedShadowCaster) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertAtBottom, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(newCaster)))
	_ = _hr
}

func (this *ICompositionProjectedShadowCasterCollection) InsertAtTop(newCaster *ICompositionProjectedShadowCaster) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertAtTop, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(newCaster)))
	_ = _hr
}

func (this *ICompositionProjectedShadowCasterCollection) InsertBelow(newCaster *ICompositionProjectedShadowCaster, reference *ICompositionProjectedShadowCaster) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertBelow, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(newCaster)), uintptr(unsafe.Pointer(reference)))
	_ = _hr
}

func (this *ICompositionProjectedShadowCasterCollection) Remove(caster *ICompositionProjectedShadowCaster) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(caster)))
	_ = _hr
}

func (this *ICompositionProjectedShadowCasterCollection) RemoveAll() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveAll, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 56FBB136-E94F-5299-AB5B-6E15E38BD899
var IID_ICompositionProjectedShadowCasterCollectionStatics = syscall.GUID{0x56FBB136, 0xE94F, 0x5299,
	[8]byte{0xAB, 0x5B, 0x6E, 0x15, 0xE3, 0x8B, 0xD8, 0x99}}

type ICompositionProjectedShadowCasterCollectionStaticsInterface interface {
	win32.IInspectableInterface
	Get_MaxRespectedCasters() int32
}

type ICompositionProjectedShadowCasterCollectionStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_MaxRespectedCasters uintptr
}

type ICompositionProjectedShadowCasterCollectionStatics struct {
	win32.IInspectable
}

func (this *ICompositionProjectedShadowCasterCollectionStatics) Vtbl() *ICompositionProjectedShadowCasterCollectionStaticsVtbl {
	return (*ICompositionProjectedShadowCasterCollectionStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionProjectedShadowCasterCollectionStatics) Get_MaxRespectedCasters() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxRespectedCasters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 1377985A-6A49-536A-9BE4-A96A8E5298A9
var IID_ICompositionProjectedShadowReceiver = syscall.GUID{0x1377985A, 0x6A49, 0x536A,
	[8]byte{0x9B, 0xE4, 0xA9, 0x6A, 0x8E, 0x52, 0x98, 0xA9}}

type ICompositionProjectedShadowReceiverInterface interface {
	win32.IInspectableInterface
	Get_ReceivingVisual() *IVisual
	Put_ReceivingVisual(value *IVisual)
}

type ICompositionProjectedShadowReceiverVtbl struct {
	win32.IInspectableVtbl
	Get_ReceivingVisual uintptr
	Put_ReceivingVisual uintptr
}

type ICompositionProjectedShadowReceiver struct {
	win32.IInspectable
}

func (this *ICompositionProjectedShadowReceiver) Vtbl() *ICompositionProjectedShadowReceiverVtbl {
	return (*ICompositionProjectedShadowReceiverVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionProjectedShadowReceiver) Get_ReceivingVisual() *IVisual {
	var _result *IVisual
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReceivingVisual, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionProjectedShadowReceiver) Put_ReceivingVisual(value *IVisual) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ReceivingVisual, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 02B3E3B7-27D2-599F-AC4B-AB787CDDE6FD
var IID_ICompositionProjectedShadowReceiverUnorderedCollection = syscall.GUID{0x02B3E3B7, 0x27D2, 0x599F,
	[8]byte{0xAC, 0x4B, 0xAB, 0x78, 0x7C, 0xDD, 0xE6, 0xFD}}

type ICompositionProjectedShadowReceiverUnorderedCollectionInterface interface {
	win32.IInspectableInterface
	Add(value *ICompositionProjectedShadowReceiver)
	Get_Count() int32
	Remove(value *ICompositionProjectedShadowReceiver)
	RemoveAll()
}

type ICompositionProjectedShadowReceiverUnorderedCollectionVtbl struct {
	win32.IInspectableVtbl
	Add       uintptr
	Get_Count uintptr
	Remove    uintptr
	RemoveAll uintptr
}

type ICompositionProjectedShadowReceiverUnorderedCollection struct {
	win32.IInspectable
}

func (this *ICompositionProjectedShadowReceiverUnorderedCollection) Vtbl() *ICompositionProjectedShadowReceiverUnorderedCollectionVtbl {
	return (*ICompositionProjectedShadowReceiverUnorderedCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionProjectedShadowReceiverUnorderedCollection) Add(value *ICompositionProjectedShadowReceiver) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ICompositionProjectedShadowReceiverUnorderedCollection) Get_Count() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Count, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionProjectedShadowReceiverUnorderedCollection) Remove(value *ICompositionProjectedShadowReceiver) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ICompositionProjectedShadowReceiverUnorderedCollection) RemoveAll() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveAll, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// C9D6D202-5F67-4453-9117-9EADD430D3C2
var IID_ICompositionPropertySet = syscall.GUID{0xC9D6D202, 0x5F67, 0x4453,
	[8]byte{0x91, 0x17, 0x9E, 0xAD, 0xD4, 0x30, 0xD3, 0xC2}}

type ICompositionPropertySetInterface interface {
	win32.IInspectableInterface
	InsertColor(propertyName string, value unsafe.Pointer)
	InsertMatrix3x2(propertyName string, value unsafe.Pointer)
	InsertMatrix4x4(propertyName string, value unsafe.Pointer)
	InsertQuaternion(propertyName string, value unsafe.Pointer)
	InsertScalar(propertyName string, value float32)
	InsertVector2(propertyName string, value unsafe.Pointer)
	InsertVector3(propertyName string, value unsafe.Pointer)
	InsertVector4(propertyName string, value unsafe.Pointer)
	TryGetColor(propertyName string, value unsafe.Pointer) CompositionGetValueStatus
	TryGetMatrix3x2(propertyName string, value unsafe.Pointer) CompositionGetValueStatus
	TryGetMatrix4x4(propertyName string, value unsafe.Pointer) CompositionGetValueStatus
	TryGetQuaternion(propertyName string, value unsafe.Pointer) CompositionGetValueStatus
	TryGetScalar(propertyName string, value float32) CompositionGetValueStatus
	TryGetVector2(propertyName string, value unsafe.Pointer) CompositionGetValueStatus
	TryGetVector3(propertyName string, value unsafe.Pointer) CompositionGetValueStatus
	TryGetVector4(propertyName string, value unsafe.Pointer) CompositionGetValueStatus
}

type ICompositionPropertySetVtbl struct {
	win32.IInspectableVtbl
	InsertColor      uintptr
	InsertMatrix3x2  uintptr
	InsertMatrix4x4  uintptr
	InsertQuaternion uintptr
	InsertScalar     uintptr
	InsertVector2    uintptr
	InsertVector3    uintptr
	InsertVector4    uintptr
	TryGetColor      uintptr
	TryGetMatrix3x2  uintptr
	TryGetMatrix4x4  uintptr
	TryGetQuaternion uintptr
	TryGetScalar     uintptr
	TryGetVector2    uintptr
	TryGetVector3    uintptr
	TryGetVector4    uintptr
}

type ICompositionPropertySet struct {
	win32.IInspectable
}

func (this *ICompositionPropertySet) Vtbl() *ICompositionPropertySetVtbl {
	return (*ICompositionPropertySetVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionPropertySet) InsertColor(propertyName string, value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertColor, uintptr(unsafe.Pointer(this)), NewHStr(propertyName).Ptr, uintptr(value))
	_ = _hr
}

func (this *ICompositionPropertySet) InsertMatrix3x2(propertyName string, value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertMatrix3x2, uintptr(unsafe.Pointer(this)), NewHStr(propertyName).Ptr, uintptr(value))
	_ = _hr
}

func (this *ICompositionPropertySet) InsertMatrix4x4(propertyName string, value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertMatrix4x4, uintptr(unsafe.Pointer(this)), NewHStr(propertyName).Ptr, uintptr(value))
	_ = _hr
}

func (this *ICompositionPropertySet) InsertQuaternion(propertyName string, value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertQuaternion, uintptr(unsafe.Pointer(this)), NewHStr(propertyName).Ptr, uintptr(value))
	_ = _hr
}

func (this *ICompositionPropertySet) InsertScalar(propertyName string, value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertScalar, uintptr(unsafe.Pointer(this)), NewHStr(propertyName).Ptr, uintptr(value))
	_ = _hr
}

func (this *ICompositionPropertySet) InsertVector2(propertyName string, value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertVector2, uintptr(unsafe.Pointer(this)), NewHStr(propertyName).Ptr, uintptr(value))
	_ = _hr
}

func (this *ICompositionPropertySet) InsertVector3(propertyName string, value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertVector3, uintptr(unsafe.Pointer(this)), NewHStr(propertyName).Ptr, uintptr(value))
	_ = _hr
}

func (this *ICompositionPropertySet) InsertVector4(propertyName string, value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertVector4, uintptr(unsafe.Pointer(this)), NewHStr(propertyName).Ptr, uintptr(value))
	_ = _hr
}

func (this *ICompositionPropertySet) TryGetColor(propertyName string, value unsafe.Pointer) CompositionGetValueStatus {
	var _result CompositionGetValueStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetColor, uintptr(unsafe.Pointer(this)), NewHStr(propertyName).Ptr, uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionPropertySet) TryGetMatrix3x2(propertyName string, value unsafe.Pointer) CompositionGetValueStatus {
	var _result CompositionGetValueStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetMatrix3x2, uintptr(unsafe.Pointer(this)), NewHStr(propertyName).Ptr, uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionPropertySet) TryGetMatrix4x4(propertyName string, value unsafe.Pointer) CompositionGetValueStatus {
	var _result CompositionGetValueStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetMatrix4x4, uintptr(unsafe.Pointer(this)), NewHStr(propertyName).Ptr, uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionPropertySet) TryGetQuaternion(propertyName string, value unsafe.Pointer) CompositionGetValueStatus {
	var _result CompositionGetValueStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetQuaternion, uintptr(unsafe.Pointer(this)), NewHStr(propertyName).Ptr, uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionPropertySet) TryGetScalar(propertyName string, value float32) CompositionGetValueStatus {
	var _result CompositionGetValueStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetScalar, uintptr(unsafe.Pointer(this)), NewHStr(propertyName).Ptr, uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionPropertySet) TryGetVector2(propertyName string, value unsafe.Pointer) CompositionGetValueStatus {
	var _result CompositionGetValueStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetVector2, uintptr(unsafe.Pointer(this)), NewHStr(propertyName).Ptr, uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionPropertySet) TryGetVector3(propertyName string, value unsafe.Pointer) CompositionGetValueStatus {
	var _result CompositionGetValueStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetVector3, uintptr(unsafe.Pointer(this)), NewHStr(propertyName).Ptr, uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionPropertySet) TryGetVector4(propertyName string, value unsafe.Pointer) CompositionGetValueStatus {
	var _result CompositionGetValueStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetVector4, uintptr(unsafe.Pointer(this)), NewHStr(propertyName).Ptr, uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// DE80731E-A211-4455-8880-7D0F3F6A44FD
var IID_ICompositionPropertySet2 = syscall.GUID{0xDE80731E, 0xA211, 0x4455,
	[8]byte{0x88, 0x80, 0x7D, 0x0F, 0x3F, 0x6A, 0x44, 0xFD}}

type ICompositionPropertySet2Interface interface {
	win32.IInspectableInterface
	InsertBoolean(propertyName string, value bool)
	TryGetBoolean(propertyName string, value bool) CompositionGetValueStatus
}

type ICompositionPropertySet2Vtbl struct {
	win32.IInspectableVtbl
	InsertBoolean uintptr
	TryGetBoolean uintptr
}

type ICompositionPropertySet2 struct {
	win32.IInspectable
}

func (this *ICompositionPropertySet2) Vtbl() *ICompositionPropertySet2Vtbl {
	return (*ICompositionPropertySet2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionPropertySet2) InsertBoolean(propertyName string, value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertBoolean, uintptr(unsafe.Pointer(this)), NewHStr(propertyName).Ptr, uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICompositionPropertySet2) TryGetBoolean(propertyName string, value bool) CompositionGetValueStatus {
	var _result CompositionGetValueStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetBoolean, uintptr(unsafe.Pointer(this)), NewHStr(propertyName).Ptr, uintptr(*(*byte)(unsafe.Pointer(&value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 3D3B50C5-E3FA-4CE2-B9FC-3EE12561788F
var IID_ICompositionRadialGradientBrush = syscall.GUID{0x3D3B50C5, 0xE3FA, 0x4CE2,
	[8]byte{0xB9, 0xFC, 0x3E, 0xE1, 0x25, 0x61, 0x78, 0x8F}}

type ICompositionRadialGradientBrushInterface interface {
	win32.IInspectableInterface
	Get_EllipseCenter() unsafe.Pointer
	Put_EllipseCenter(value unsafe.Pointer)
	Get_EllipseRadius() unsafe.Pointer
	Put_EllipseRadius(value unsafe.Pointer)
	Get_GradientOriginOffset() unsafe.Pointer
	Put_GradientOriginOffset(value unsafe.Pointer)
}

type ICompositionRadialGradientBrushVtbl struct {
	win32.IInspectableVtbl
	Get_EllipseCenter        uintptr
	Put_EllipseCenter        uintptr
	Get_EllipseRadius        uintptr
	Put_EllipseRadius        uintptr
	Get_GradientOriginOffset uintptr
	Put_GradientOriginOffset uintptr
}

type ICompositionRadialGradientBrush struct {
	win32.IInspectable
}

func (this *ICompositionRadialGradientBrush) Vtbl() *ICompositionRadialGradientBrushVtbl {
	return (*ICompositionRadialGradientBrushVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionRadialGradientBrush) Get_EllipseCenter() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EllipseCenter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionRadialGradientBrush) Put_EllipseCenter(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_EllipseCenter, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionRadialGradientBrush) Get_EllipseRadius() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EllipseRadius, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionRadialGradientBrush) Put_EllipseRadius(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_EllipseRadius, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionRadialGradientBrush) Get_GradientOriginOffset() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GradientOriginOffset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionRadialGradientBrush) Put_GradientOriginOffset(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_GradientOriginOffset, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 0CD51428-5356-4246-AECF-7A0B76975400
var IID_ICompositionRectangleGeometry = syscall.GUID{0x0CD51428, 0x5356, 0x4246,
	[8]byte{0xAE, 0xCF, 0x7A, 0x0B, 0x76, 0x97, 0x54, 0x00}}

type ICompositionRectangleGeometryInterface interface {
	win32.IInspectableInterface
	Get_Offset() unsafe.Pointer
	Put_Offset(value unsafe.Pointer)
	Get_Size() unsafe.Pointer
	Put_Size(value unsafe.Pointer)
}

type ICompositionRectangleGeometryVtbl struct {
	win32.IInspectableVtbl
	Get_Offset uintptr
	Put_Offset uintptr
	Get_Size   uintptr
	Put_Size   uintptr
}

type ICompositionRectangleGeometry struct {
	win32.IInspectable
}

func (this *ICompositionRectangleGeometry) Vtbl() *ICompositionRectangleGeometryVtbl {
	return (*ICompositionRectangleGeometryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionRectangleGeometry) Get_Offset() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Offset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionRectangleGeometry) Put_Offset(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Offset, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionRectangleGeometry) Get_Size() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Size, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionRectangleGeometry) Put_Size(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Size, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 8770C822-1D50-4B8B-B013-7C9A0E46935F
var IID_ICompositionRoundedRectangleGeometry = syscall.GUID{0x8770C822, 0x1D50, 0x4B8B,
	[8]byte{0xB0, 0x13, 0x7C, 0x9A, 0x0E, 0x46, 0x93, 0x5F}}

type ICompositionRoundedRectangleGeometryInterface interface {
	win32.IInspectableInterface
	Get_CornerRadius() unsafe.Pointer
	Put_CornerRadius(value unsafe.Pointer)
	Get_Offset() unsafe.Pointer
	Put_Offset(value unsafe.Pointer)
	Get_Size() unsafe.Pointer
	Put_Size(value unsafe.Pointer)
}

type ICompositionRoundedRectangleGeometryVtbl struct {
	win32.IInspectableVtbl
	Get_CornerRadius uintptr
	Put_CornerRadius uintptr
	Get_Offset       uintptr
	Put_Offset       uintptr
	Get_Size         uintptr
	Put_Size         uintptr
}

type ICompositionRoundedRectangleGeometry struct {
	win32.IInspectable
}

func (this *ICompositionRoundedRectangleGeometry) Vtbl() *ICompositionRoundedRectangleGeometryVtbl {
	return (*ICompositionRoundedRectangleGeometryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionRoundedRectangleGeometry) Get_CornerRadius() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CornerRadius, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionRoundedRectangleGeometry) Put_CornerRadius(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CornerRadius, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionRoundedRectangleGeometry) Get_Offset() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Offset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionRoundedRectangleGeometry) Put_Offset(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Offset, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionRoundedRectangleGeometry) Get_Size() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Size, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionRoundedRectangleGeometry) Put_Size(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Size, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 0D00DAD0-FB07-46FD-8C72-6280D1A3D1DD
var IID_ICompositionScopedBatch = syscall.GUID{0x0D00DAD0, 0xFB07, 0x46FD,
	[8]byte{0x8C, 0x72, 0x62, 0x80, 0xD1, 0xA3, 0xD1, 0xDD}}

type ICompositionScopedBatchInterface interface {
	win32.IInspectableInterface
	Get_IsActive() bool
	Get_IsEnded() bool
	End()
	Resume()
	Suspend()
	Add_Completed(handler TypedEventHandler[interface{}, *ICompositionBatchCompletedEventArgs]) EventRegistrationToken
	Remove_Completed(token EventRegistrationToken)
}

type ICompositionScopedBatchVtbl struct {
	win32.IInspectableVtbl
	Get_IsActive     uintptr
	Get_IsEnded      uintptr
	End              uintptr
	Resume           uintptr
	Suspend          uintptr
	Add_Completed    uintptr
	Remove_Completed uintptr
}

type ICompositionScopedBatch struct {
	win32.IInspectable
}

func (this *ICompositionScopedBatch) Vtbl() *ICompositionScopedBatchVtbl {
	return (*ICompositionScopedBatchVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionScopedBatch) Get_IsActive() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsActive, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionScopedBatch) Get_IsEnded() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEnded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionScopedBatch) End() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().End, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *ICompositionScopedBatch) Resume() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Resume, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *ICompositionScopedBatch) Suspend() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Suspend, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *ICompositionScopedBatch) Add_Completed(handler TypedEventHandler[interface{}, *ICompositionBatchCompletedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Completed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionScopedBatch) Remove_Completed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Completed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 329E52E2-4335-49CC-B14A-37782D10F0C4
var IID_ICompositionShadow = syscall.GUID{0x329E52E2, 0x4335, 0x49CC,
	[8]byte{0xB1, 0x4A, 0x37, 0x78, 0x2D, 0x10, 0xF0, 0xC4}}

type ICompositionShadowInterface interface {
	win32.IInspectableInterface
}

type ICompositionShadowVtbl struct {
	win32.IInspectableVtbl
}

type ICompositionShadow struct {
	win32.IInspectable
}

func (this *ICompositionShadow) Vtbl() *ICompositionShadowVtbl {
	return (*ICompositionShadowVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 221F492F-DCBA-4B91-999E-1DC217A01530
var IID_ICompositionShadowFactory = syscall.GUID{0x221F492F, 0xDCBA, 0x4B91,
	[8]byte{0x99, 0x9E, 0x1D, 0xC2, 0x17, 0xA0, 0x15, 0x30}}

type ICompositionShadowFactoryInterface interface {
	win32.IInspectableInterface
}

type ICompositionShadowFactoryVtbl struct {
	win32.IInspectableVtbl
}

type ICompositionShadowFactory struct {
	win32.IInspectable
}

func (this *ICompositionShadowFactory) Vtbl() *ICompositionShadowFactoryVtbl {
	return (*ICompositionShadowFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// B47CE2F7-9A88-42C4-9E87-2E500CA8688C
var IID_ICompositionShape = syscall.GUID{0xB47CE2F7, 0x9A88, 0x42C4,
	[8]byte{0x9E, 0x87, 0x2E, 0x50, 0x0C, 0xA8, 0x68, 0x8C}}

type ICompositionShapeInterface interface {
	win32.IInspectableInterface
	Get_CenterPoint() unsafe.Pointer
	Put_CenterPoint(value unsafe.Pointer)
	Get_Offset() unsafe.Pointer
	Put_Offset(value unsafe.Pointer)
	Get_RotationAngle() float32
	Put_RotationAngle(value float32)
	Get_RotationAngleInDegrees() float32
	Put_RotationAngleInDegrees(value float32)
	Get_Scale() unsafe.Pointer
	Put_Scale(value unsafe.Pointer)
	Get_TransformMatrix() unsafe.Pointer
	Put_TransformMatrix(value unsafe.Pointer)
}

type ICompositionShapeVtbl struct {
	win32.IInspectableVtbl
	Get_CenterPoint            uintptr
	Put_CenterPoint            uintptr
	Get_Offset                 uintptr
	Put_Offset                 uintptr
	Get_RotationAngle          uintptr
	Put_RotationAngle          uintptr
	Get_RotationAngleInDegrees uintptr
	Put_RotationAngleInDegrees uintptr
	Get_Scale                  uintptr
	Put_Scale                  uintptr
	Get_TransformMatrix        uintptr
	Put_TransformMatrix        uintptr
}

type ICompositionShape struct {
	win32.IInspectable
}

func (this *ICompositionShape) Vtbl() *ICompositionShapeVtbl {
	return (*ICompositionShapeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionShape) Get_CenterPoint() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CenterPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionShape) Put_CenterPoint(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CenterPoint, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionShape) Get_Offset() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Offset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionShape) Put_Offset(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Offset, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionShape) Get_RotationAngle() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RotationAngle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionShape) Put_RotationAngle(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RotationAngle, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionShape) Get_RotationAngleInDegrees() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RotationAngleInDegrees, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionShape) Put_RotationAngleInDegrees(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RotationAngleInDegrees, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionShape) Get_Scale() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Scale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionShape) Put_Scale(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Scale, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionShape) Get_TransformMatrix() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TransformMatrix, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionShape) Put_TransformMatrix(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TransformMatrix, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 1DFC36D0-B05A-44EF-82B0-12118BCD4CD0
var IID_ICompositionShapeFactory = syscall.GUID{0x1DFC36D0, 0xB05A, 0x44EF,
	[8]byte{0x82, 0xB0, 0x12, 0x11, 0x8B, 0xCD, 0x4C, 0xD0}}

type ICompositionShapeFactoryInterface interface {
	win32.IInspectableInterface
}

type ICompositionShapeFactoryVtbl struct {
	win32.IInspectableVtbl
}

type ICompositionShapeFactory struct {
	win32.IInspectable
}

func (this *ICompositionShapeFactory) Vtbl() *ICompositionShapeFactoryVtbl {
	return (*ICompositionShapeFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 401B61BB-0007-4363-B1F3-6BCC003FB83E
var IID_ICompositionSpriteShape = syscall.GUID{0x401B61BB, 0x0007, 0x4363,
	[8]byte{0xB1, 0xF3, 0x6B, 0xCC, 0x00, 0x3F, 0xB8, 0x3E}}

type ICompositionSpriteShapeInterface interface {
	win32.IInspectableInterface
	Get_FillBrush() *ICompositionBrush
	Put_FillBrush(value *ICompositionBrush)
	Get_Geometry() *ICompositionGeometry
	Put_Geometry(value *ICompositionGeometry)
	Get_IsStrokeNonScaling() bool
	Put_IsStrokeNonScaling(value bool)
	Get_StrokeBrush() *ICompositionBrush
	Put_StrokeBrush(value *ICompositionBrush)
	Get_StrokeDashArray() *IVector[float32]
	Get_StrokeDashCap() CompositionStrokeCap
	Put_StrokeDashCap(value CompositionStrokeCap)
	Get_StrokeDashOffset() float32
	Put_StrokeDashOffset(value float32)
	Get_StrokeEndCap() CompositionStrokeCap
	Put_StrokeEndCap(value CompositionStrokeCap)
	Get_StrokeLineJoin() CompositionStrokeLineJoin
	Put_StrokeLineJoin(value CompositionStrokeLineJoin)
	Get_StrokeMiterLimit() float32
	Put_StrokeMiterLimit(value float32)
	Get_StrokeStartCap() CompositionStrokeCap
	Put_StrokeStartCap(value CompositionStrokeCap)
	Get_StrokeThickness() float32
	Put_StrokeThickness(value float32)
}

type ICompositionSpriteShapeVtbl struct {
	win32.IInspectableVtbl
	Get_FillBrush          uintptr
	Put_FillBrush          uintptr
	Get_Geometry           uintptr
	Put_Geometry           uintptr
	Get_IsStrokeNonScaling uintptr
	Put_IsStrokeNonScaling uintptr
	Get_StrokeBrush        uintptr
	Put_StrokeBrush        uintptr
	Get_StrokeDashArray    uintptr
	Get_StrokeDashCap      uintptr
	Put_StrokeDashCap      uintptr
	Get_StrokeDashOffset   uintptr
	Put_StrokeDashOffset   uintptr
	Get_StrokeEndCap       uintptr
	Put_StrokeEndCap       uintptr
	Get_StrokeLineJoin     uintptr
	Put_StrokeLineJoin     uintptr
	Get_StrokeMiterLimit   uintptr
	Put_StrokeMiterLimit   uintptr
	Get_StrokeStartCap     uintptr
	Put_StrokeStartCap     uintptr
	Get_StrokeThickness    uintptr
	Put_StrokeThickness    uintptr
}

type ICompositionSpriteShape struct {
	win32.IInspectable
}

func (this *ICompositionSpriteShape) Vtbl() *ICompositionSpriteShapeVtbl {
	return (*ICompositionSpriteShapeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionSpriteShape) Get_FillBrush() *ICompositionBrush {
	var _result *ICompositionBrush
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FillBrush, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionSpriteShape) Put_FillBrush(value *ICompositionBrush) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FillBrush, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ICompositionSpriteShape) Get_Geometry() *ICompositionGeometry {
	var _result *ICompositionGeometry
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Geometry, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionSpriteShape) Put_Geometry(value *ICompositionGeometry) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Geometry, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ICompositionSpriteShape) Get_IsStrokeNonScaling() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsStrokeNonScaling, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionSpriteShape) Put_IsStrokeNonScaling(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsStrokeNonScaling, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICompositionSpriteShape) Get_StrokeBrush() *ICompositionBrush {
	var _result *ICompositionBrush
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StrokeBrush, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionSpriteShape) Put_StrokeBrush(value *ICompositionBrush) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StrokeBrush, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ICompositionSpriteShape) Get_StrokeDashArray() *IVector[float32] {
	var _result *IVector[float32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StrokeDashArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionSpriteShape) Get_StrokeDashCap() CompositionStrokeCap {
	var _result CompositionStrokeCap
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StrokeDashCap, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionSpriteShape) Put_StrokeDashCap(value CompositionStrokeCap) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StrokeDashCap, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionSpriteShape) Get_StrokeDashOffset() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StrokeDashOffset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionSpriteShape) Put_StrokeDashOffset(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StrokeDashOffset, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionSpriteShape) Get_StrokeEndCap() CompositionStrokeCap {
	var _result CompositionStrokeCap
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StrokeEndCap, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionSpriteShape) Put_StrokeEndCap(value CompositionStrokeCap) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StrokeEndCap, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionSpriteShape) Get_StrokeLineJoin() CompositionStrokeLineJoin {
	var _result CompositionStrokeLineJoin
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StrokeLineJoin, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionSpriteShape) Put_StrokeLineJoin(value CompositionStrokeLineJoin) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StrokeLineJoin, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionSpriteShape) Get_StrokeMiterLimit() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StrokeMiterLimit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionSpriteShape) Put_StrokeMiterLimit(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StrokeMiterLimit, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionSpriteShape) Get_StrokeStartCap() CompositionStrokeCap {
	var _result CompositionStrokeCap
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StrokeStartCap, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionSpriteShape) Put_StrokeStartCap(value CompositionStrokeCap) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StrokeStartCap, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionSpriteShape) Get_StrokeThickness() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StrokeThickness, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionSpriteShape) Put_StrokeThickness(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StrokeThickness, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 397DAFE4-B6C2-5BB9-951D-F5707DE8B7BC
var IID_ICompositionSupportsSystemBackdrop = syscall.GUID{0x397DAFE4, 0xB6C2, 0x5BB9,
	[8]byte{0x95, 0x1D, 0xF5, 0x70, 0x7D, 0xE8, 0xB7, 0xBC}}

type ICompositionSupportsSystemBackdropInterface interface {
	win32.IInspectableInterface
	Get_SystemBackdrop() *ICompositionBrush
	Put_SystemBackdrop(value *ICompositionBrush)
}

type ICompositionSupportsSystemBackdropVtbl struct {
	win32.IInspectableVtbl
	Get_SystemBackdrop uintptr
	Put_SystemBackdrop uintptr
}

type ICompositionSupportsSystemBackdrop struct {
	win32.IInspectable
}

func (this *ICompositionSupportsSystemBackdrop) Vtbl() *ICompositionSupportsSystemBackdropVtbl {
	return (*ICompositionSupportsSystemBackdropVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionSupportsSystemBackdrop) Get_SystemBackdrop() *ICompositionBrush {
	var _result *ICompositionBrush
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SystemBackdrop, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionSupportsSystemBackdrop) Put_SystemBackdrop(value *ICompositionBrush) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SystemBackdrop, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 1527540D-42C7-47A6-A408-668F79A90DFB
var IID_ICompositionSurface = syscall.GUID{0x1527540D, 0x42C7, 0x47A6,
	[8]byte{0xA4, 0x08, 0x66, 0x8F, 0x79, 0xA9, 0x0D, 0xFB}}

type ICompositionSurfaceInterface interface {
	win32.IInspectableInterface
}

type ICompositionSurfaceVtbl struct {
	win32.IInspectableVtbl
}

type ICompositionSurface struct {
	win32.IInspectable
}

func (this *ICompositionSurface) Vtbl() *ICompositionSurfaceVtbl {
	return (*ICompositionSurfaceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// AD016D79-1E4C-4C0D-9C29-83338C87C162
var IID_ICompositionSurfaceBrush = syscall.GUID{0xAD016D79, 0x1E4C, 0x4C0D,
	[8]byte{0x9C, 0x29, 0x83, 0x33, 0x8C, 0x87, 0xC1, 0x62}}

type ICompositionSurfaceBrushInterface interface {
	win32.IInspectableInterface
	Get_BitmapInterpolationMode() CompositionBitmapInterpolationMode
	Put_BitmapInterpolationMode(value CompositionBitmapInterpolationMode)
	Get_HorizontalAlignmentRatio() float32
	Put_HorizontalAlignmentRatio(value float32)
	Get_Stretch() CompositionStretch
	Put_Stretch(value CompositionStretch)
	Get_Surface() *ICompositionSurface
	Put_Surface(value *ICompositionSurface)
	Get_VerticalAlignmentRatio() float32
	Put_VerticalAlignmentRatio(value float32)
}

type ICompositionSurfaceBrushVtbl struct {
	win32.IInspectableVtbl
	Get_BitmapInterpolationMode  uintptr
	Put_BitmapInterpolationMode  uintptr
	Get_HorizontalAlignmentRatio uintptr
	Put_HorizontalAlignmentRatio uintptr
	Get_Stretch                  uintptr
	Put_Stretch                  uintptr
	Get_Surface                  uintptr
	Put_Surface                  uintptr
	Get_VerticalAlignmentRatio   uintptr
	Put_VerticalAlignmentRatio   uintptr
}

type ICompositionSurfaceBrush struct {
	win32.IInspectable
}

func (this *ICompositionSurfaceBrush) Vtbl() *ICompositionSurfaceBrushVtbl {
	return (*ICompositionSurfaceBrushVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionSurfaceBrush) Get_BitmapInterpolationMode() CompositionBitmapInterpolationMode {
	var _result CompositionBitmapInterpolationMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BitmapInterpolationMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionSurfaceBrush) Put_BitmapInterpolationMode(value CompositionBitmapInterpolationMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BitmapInterpolationMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionSurfaceBrush) Get_HorizontalAlignmentRatio() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HorizontalAlignmentRatio, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionSurfaceBrush) Put_HorizontalAlignmentRatio(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_HorizontalAlignmentRatio, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionSurfaceBrush) Get_Stretch() CompositionStretch {
	var _result CompositionStretch
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Stretch, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionSurfaceBrush) Put_Stretch(value CompositionStretch) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Stretch, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionSurfaceBrush) Get_Surface() *ICompositionSurface {
	var _result *ICompositionSurface
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Surface, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionSurfaceBrush) Put_Surface(value *ICompositionSurface) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Surface, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ICompositionSurfaceBrush) Get_VerticalAlignmentRatio() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VerticalAlignmentRatio, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionSurfaceBrush) Put_VerticalAlignmentRatio(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_VerticalAlignmentRatio, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// D27174D5-64F5-4692-9DC7-71B61D7E5880
var IID_ICompositionSurfaceBrush2 = syscall.GUID{0xD27174D5, 0x64F5, 0x4692,
	[8]byte{0x9D, 0xC7, 0x71, 0xB6, 0x1D, 0x7E, 0x58, 0x80}}

type ICompositionSurfaceBrush2Interface interface {
	win32.IInspectableInterface
	Get_AnchorPoint() unsafe.Pointer
	Put_AnchorPoint(value unsafe.Pointer)
	Get_CenterPoint() unsafe.Pointer
	Put_CenterPoint(value unsafe.Pointer)
	Get_Offset() unsafe.Pointer
	Put_Offset(value unsafe.Pointer)
	Get_RotationAngle() float32
	Put_RotationAngle(value float32)
	Get_RotationAngleInDegrees() float32
	Put_RotationAngleInDegrees(value float32)
	Get_Scale() unsafe.Pointer
	Put_Scale(value unsafe.Pointer)
	Get_TransformMatrix() unsafe.Pointer
	Put_TransformMatrix(value unsafe.Pointer)
}

type ICompositionSurfaceBrush2Vtbl struct {
	win32.IInspectableVtbl
	Get_AnchorPoint            uintptr
	Put_AnchorPoint            uintptr
	Get_CenterPoint            uintptr
	Put_CenterPoint            uintptr
	Get_Offset                 uintptr
	Put_Offset                 uintptr
	Get_RotationAngle          uintptr
	Put_RotationAngle          uintptr
	Get_RotationAngleInDegrees uintptr
	Put_RotationAngleInDegrees uintptr
	Get_Scale                  uintptr
	Put_Scale                  uintptr
	Get_TransformMatrix        uintptr
	Put_TransformMatrix        uintptr
}

type ICompositionSurfaceBrush2 struct {
	win32.IInspectable
}

func (this *ICompositionSurfaceBrush2) Vtbl() *ICompositionSurfaceBrush2Vtbl {
	return (*ICompositionSurfaceBrush2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionSurfaceBrush2) Get_AnchorPoint() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AnchorPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionSurfaceBrush2) Put_AnchorPoint(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AnchorPoint, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionSurfaceBrush2) Get_CenterPoint() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CenterPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionSurfaceBrush2) Put_CenterPoint(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CenterPoint, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionSurfaceBrush2) Get_Offset() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Offset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionSurfaceBrush2) Put_Offset(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Offset, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionSurfaceBrush2) Get_RotationAngle() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RotationAngle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionSurfaceBrush2) Put_RotationAngle(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RotationAngle, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionSurfaceBrush2) Get_RotationAngleInDegrees() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RotationAngleInDegrees, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionSurfaceBrush2) Put_RotationAngleInDegrees(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RotationAngleInDegrees, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionSurfaceBrush2) Get_Scale() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Scale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionSurfaceBrush2) Put_Scale(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Scale, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionSurfaceBrush2) Get_TransformMatrix() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TransformMatrix, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionSurfaceBrush2) Put_TransformMatrix(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TransformMatrix, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 550BB289-1FE0-42E5-8195-1EEFA87FF08E
var IID_ICompositionSurfaceBrush3 = syscall.GUID{0x550BB289, 0x1FE0, 0x42E5,
	[8]byte{0x81, 0x95, 0x1E, 0xEF, 0xA8, 0x7F, 0xF0, 0x8E}}

type ICompositionSurfaceBrush3Interface interface {
	win32.IInspectableInterface
	Get_SnapToPixels() bool
	Put_SnapToPixels(value bool)
}

type ICompositionSurfaceBrush3Vtbl struct {
	win32.IInspectableVtbl
	Get_SnapToPixels uintptr
	Put_SnapToPixels uintptr
}

type ICompositionSurfaceBrush3 struct {
	win32.IInspectable
}

func (this *ICompositionSurfaceBrush3) Vtbl() *ICompositionSurfaceBrush3Vtbl {
	return (*ICompositionSurfaceBrush3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionSurfaceBrush3) Get_SnapToPixels() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SnapToPixels, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionSurfaceBrush3) Put_SnapToPixels(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SnapToPixels, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// E01622C8-2332-55C7-8868-A7312C5C229D
var IID_ICompositionSurfaceFacade = syscall.GUID{0xE01622C8, 0x2332, 0x55C7,
	[8]byte{0x88, 0x68, 0xA7, 0x31, 0x2C, 0x5C, 0x22, 0x9D}}

type ICompositionSurfaceFacadeInterface interface {
	win32.IInspectableInterface
	GetRealSurface() *ICompositionSurface
}

type ICompositionSurfaceFacadeVtbl struct {
	win32.IInspectableVtbl
	GetRealSurface uintptr
}

type ICompositionSurfaceFacade struct {
	win32.IInspectable
}

func (this *ICompositionSurfaceFacade) Vtbl() *ICompositionSurfaceFacadeVtbl {
	return (*ICompositionSurfaceFacadeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionSurfaceFacade) GetRealSurface() *ICompositionSurface {
	var _result *ICompositionSurface
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetRealSurface, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A1BEA8BA-D726-4663-8129-6B5E7927FFA6
var IID_ICompositionTarget = syscall.GUID{0xA1BEA8BA, 0xD726, 0x4663,
	[8]byte{0x81, 0x29, 0x6B, 0x5E, 0x79, 0x27, 0xFF, 0xA6}}

type ICompositionTargetInterface interface {
	win32.IInspectableInterface
	Get_Root() *IVisual
	Put_Root(value *IVisual)
}

type ICompositionTargetVtbl struct {
	win32.IInspectableVtbl
	Get_Root uintptr
	Put_Root uintptr
}

type ICompositionTarget struct {
	win32.IInspectable
}

func (this *ICompositionTarget) Vtbl() *ICompositionTargetVtbl {
	return (*ICompositionTargetVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionTarget) Get_Root() *IVisual {
	var _result *IVisual
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Root, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionTarget) Put_Root(value *IVisual) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Root, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 93CD9D2B-8516-4B14-A8CE-F49E2119EC42
var IID_ICompositionTargetFactory = syscall.GUID{0x93CD9D2B, 0x8516, 0x4B14,
	[8]byte{0xA8, 0xCE, 0xF4, 0x9E, 0x21, 0x19, 0xEC, 0x42}}

type ICompositionTargetFactoryInterface interface {
	win32.IInspectableInterface
}

type ICompositionTargetFactoryVtbl struct {
	win32.IInspectableVtbl
}

type ICompositionTargetFactory struct {
	win32.IInspectable
}

func (this *ICompositionTargetFactory) Vtbl() *ICompositionTargetFactoryVtbl {
	return (*ICompositionTargetFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 7CD54529-FBED-4112-ABC5-185906DD927C
var IID_ICompositionTransform = syscall.GUID{0x7CD54529, 0xFBED, 0x4112,
	[8]byte{0xAB, 0xC5, 0x18, 0x59, 0x06, 0xDD, 0x92, 0x7C}}

type ICompositionTransformInterface interface {
	win32.IInspectableInterface
}

type ICompositionTransformVtbl struct {
	win32.IInspectableVtbl
}

type ICompositionTransform struct {
	win32.IInspectable
}

func (this *ICompositionTransform) Vtbl() *ICompositionTransformVtbl {
	return (*ICompositionTransformVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// AAAECA26-C149-517A-8F72-6BFF7A65CE08
var IID_ICompositionTransformFactory = syscall.GUID{0xAAAECA26, 0xC149, 0x517A,
	[8]byte{0x8F, 0x72, 0x6B, 0xFF, 0x7A, 0x65, 0xCE, 0x08}}

type ICompositionTransformFactoryInterface interface {
	win32.IInspectableInterface
}

type ICompositionTransformFactoryVtbl struct {
	win32.IInspectableVtbl
}

type ICompositionTransformFactory struct {
	win32.IInspectable
}

func (this *ICompositionTransformFactory) Vtbl() *ICompositionTransformFactoryVtbl {
	return (*ICompositionTransformFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// B440BF07-068F-4537-84C6-4ECBE019E1F4
var IID_ICompositionViewBox = syscall.GUID{0xB440BF07, 0x068F, 0x4537,
	[8]byte{0x84, 0xC6, 0x4E, 0xCB, 0xE0, 0x19, 0xE1, 0xF4}}

type ICompositionViewBoxInterface interface {
	win32.IInspectableInterface
	Get_HorizontalAlignmentRatio() float32
	Put_HorizontalAlignmentRatio(value float32)
	Get_Offset() unsafe.Pointer
	Put_Offset(value unsafe.Pointer)
	Get_Size() unsafe.Pointer
	Put_Size(value unsafe.Pointer)
	Get_Stretch() CompositionStretch
	Put_Stretch(value CompositionStretch)
	Get_VerticalAlignmentRatio() float32
	Put_VerticalAlignmentRatio(value float32)
}

type ICompositionViewBoxVtbl struct {
	win32.IInspectableVtbl
	Get_HorizontalAlignmentRatio uintptr
	Put_HorizontalAlignmentRatio uintptr
	Get_Offset                   uintptr
	Put_Offset                   uintptr
	Get_Size                     uintptr
	Put_Size                     uintptr
	Get_Stretch                  uintptr
	Put_Stretch                  uintptr
	Get_VerticalAlignmentRatio   uintptr
	Put_VerticalAlignmentRatio   uintptr
}

type ICompositionViewBox struct {
	win32.IInspectable
}

func (this *ICompositionViewBox) Vtbl() *ICompositionViewBoxVtbl {
	return (*ICompositionViewBoxVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionViewBox) Get_HorizontalAlignmentRatio() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HorizontalAlignmentRatio, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionViewBox) Put_HorizontalAlignmentRatio(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_HorizontalAlignmentRatio, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionViewBox) Get_Offset() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Offset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionViewBox) Put_Offset(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Offset, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionViewBox) Get_Size() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Size, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionViewBox) Put_Size(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Size, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionViewBox) Get_Stretch() CompositionStretch {
	var _result CompositionStretch
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Stretch, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionViewBox) Put_Stretch(value CompositionStretch) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Stretch, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionViewBox) Get_VerticalAlignmentRatio() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VerticalAlignmentRatio, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionViewBox) Put_VerticalAlignmentRatio(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_VerticalAlignmentRatio, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// A9C384DB-8740-4F94-8B9D-B68521E7863D
var IID_ICompositionVirtualDrawingSurface = syscall.GUID{0xA9C384DB, 0x8740, 0x4F94,
	[8]byte{0x8B, 0x9D, 0xB6, 0x85, 0x21, 0xE7, 0x86, 0x3D}}

type ICompositionVirtualDrawingSurfaceInterface interface {
	win32.IInspectableInterface
	Trim(rects unsafe.Pointer)
}

type ICompositionVirtualDrawingSurfaceVtbl struct {
	win32.IInspectableVtbl
	Trim uintptr
}

type ICompositionVirtualDrawingSurface struct {
	win32.IInspectable
}

func (this *ICompositionVirtualDrawingSurface) Vtbl() *ICompositionVirtualDrawingSurfaceVtbl {
	return (*ICompositionVirtualDrawingSurfaceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionVirtualDrawingSurface) Trim(rects unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Trim, uintptr(unsafe.Pointer(this)), uintptr(rects))
	_ = _hr
}

// 6766106C-D56B-4A49-B1DF-5076A0620768
var IID_ICompositionVirtualDrawingSurfaceFactory = syscall.GUID{0x6766106C, 0xD56B, 0x4A49,
	[8]byte{0xB1, 0xDF, 0x50, 0x76, 0xA0, 0x62, 0x07, 0x68}}

type ICompositionVirtualDrawingSurfaceFactoryInterface interface {
	win32.IInspectableInterface
}

type ICompositionVirtualDrawingSurfaceFactoryVtbl struct {
	win32.IInspectableVtbl
}

type ICompositionVirtualDrawingSurfaceFactory struct {
	win32.IInspectable
}

func (this *ICompositionVirtualDrawingSurfaceFactory) Vtbl() *ICompositionVirtualDrawingSurfaceFactoryVtbl {
	return (*ICompositionVirtualDrawingSurfaceFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// B224D803-4F6E-4A3F-8CAE-3DC1CDA74FC6
var IID_ICompositionVisualSurface = syscall.GUID{0xB224D803, 0x4F6E, 0x4A3F,
	[8]byte{0x8C, 0xAE, 0x3D, 0xC1, 0xCD, 0xA7, 0x4F, 0xC6}}

type ICompositionVisualSurfaceInterface interface {
	win32.IInspectableInterface
	Get_SourceVisual() *IVisual
	Put_SourceVisual(value *IVisual)
	Get_SourceOffset() unsafe.Pointer
	Put_SourceOffset(value unsafe.Pointer)
	Get_SourceSize() unsafe.Pointer
	Put_SourceSize(value unsafe.Pointer)
}

type ICompositionVisualSurfaceVtbl struct {
	win32.IInspectableVtbl
	Get_SourceVisual uintptr
	Put_SourceVisual uintptr
	Get_SourceOffset uintptr
	Put_SourceOffset uintptr
	Get_SourceSize   uintptr
	Put_SourceSize   uintptr
}

type ICompositionVisualSurface struct {
	win32.IInspectable
}

func (this *ICompositionVisualSurface) Vtbl() *ICompositionVisualSurfaceVtbl {
	return (*ICompositionVisualSurfaceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionVisualSurface) Get_SourceVisual() *IVisual {
	var _result *IVisual
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SourceVisual, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionVisualSurface) Put_SourceVisual(value *IVisual) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SourceVisual, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ICompositionVisualSurface) Get_SourceOffset() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SourceOffset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionVisualSurface) Put_SourceOffset(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SourceOffset, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositionVisualSurface) Get_SourceSize() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SourceSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionVisualSurface) Put_SourceSize(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SourceSize, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// B403CA50-7F8C-4E83-985F-CC45060036D8
var IID_ICompositor = syscall.GUID{0xB403CA50, 0x7F8C, 0x4E83,
	[8]byte{0x98, 0x5F, 0xCC, 0x45, 0x06, 0x00, 0x36, 0xD8}}

type ICompositorInterface interface {
	win32.IInspectableInterface
	CreateColorKeyFrameAnimation() *IColorKeyFrameAnimation
	CreateColorBrush() *ICompositionColorBrush
	CreateColorBrushWithColor(color unsafe.Pointer) *ICompositionColorBrush
	CreateContainerVisual() *IContainerVisual
	CreateCubicBezierEasingFunction(controlPoint1 unsafe.Pointer, controlPoint2 unsafe.Pointer) *ICubicBezierEasingFunction
	CreateEffectFactory(graphicsEffect unsafe.Pointer) *ICompositionEffectFactory
	CreateEffectFactoryWithProperties(graphicsEffect unsafe.Pointer, animatableProperties *IIterable[string]) *ICompositionEffectFactory
	CreateExpressionAnimation() *IExpressionAnimation
	CreateExpressionAnimationWithExpression(expression string) *IExpressionAnimation
	CreateInsetClip() *IInsetClip
	CreateInsetClipWithInsets(leftInset float32, topInset float32, rightInset float32, bottomInset float32) *IInsetClip
	CreateLinearEasingFunction() *ILinearEasingFunction
	CreatePropertySet() *ICompositionPropertySet
	CreateQuaternionKeyFrameAnimation() *IQuaternionKeyFrameAnimation
	CreateScalarKeyFrameAnimation() *IScalarKeyFrameAnimation
	CreateScopedBatch(batchType CompositionBatchTypes) *ICompositionScopedBatch
	CreateSpriteVisual() *ISpriteVisual
	CreateSurfaceBrush() *ICompositionSurfaceBrush
	CreateSurfaceBrushWithSurface(surface *ICompositionSurface) *ICompositionSurfaceBrush
	CreateTargetForCurrentView() *ICompositionTarget
	CreateVector2KeyFrameAnimation() *IVector2KeyFrameAnimation
	CreateVector3KeyFrameAnimation() *IVector3KeyFrameAnimation
	CreateVector4KeyFrameAnimation() *IVector4KeyFrameAnimation
	GetCommitBatch(batchType CompositionBatchTypes) *ICompositionCommitBatch
}

type ICompositorVtbl struct {
	win32.IInspectableVtbl
	CreateColorKeyFrameAnimation            uintptr
	CreateColorBrush                        uintptr
	CreateColorBrushWithColor               uintptr
	CreateContainerVisual                   uintptr
	CreateCubicBezierEasingFunction         uintptr
	CreateEffectFactory                     uintptr
	CreateEffectFactoryWithProperties       uintptr
	CreateExpressionAnimation               uintptr
	CreateExpressionAnimationWithExpression uintptr
	CreateInsetClip                         uintptr
	CreateInsetClipWithInsets               uintptr
	CreateLinearEasingFunction              uintptr
	CreatePropertySet                       uintptr
	CreateQuaternionKeyFrameAnimation       uintptr
	CreateScalarKeyFrameAnimation           uintptr
	CreateScopedBatch                       uintptr
	CreateSpriteVisual                      uintptr
	CreateSurfaceBrush                      uintptr
	CreateSurfaceBrushWithSurface           uintptr
	CreateTargetForCurrentView              uintptr
	CreateVector2KeyFrameAnimation          uintptr
	CreateVector3KeyFrameAnimation          uintptr
	CreateVector4KeyFrameAnimation          uintptr
	GetCommitBatch                          uintptr
}

type ICompositor struct {
	win32.IInspectable
}

func (this *ICompositor) Vtbl() *ICompositorVtbl {
	return (*ICompositorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositor) CreateColorKeyFrameAnimation() *IColorKeyFrameAnimation {
	var _result *IColorKeyFrameAnimation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateColorKeyFrameAnimation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor) CreateColorBrush() *ICompositionColorBrush {
	var _result *ICompositionColorBrush
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateColorBrush, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor) CreateColorBrushWithColor(color unsafe.Pointer) *ICompositionColorBrush {
	var _result *ICompositionColorBrush
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateColorBrushWithColor, uintptr(unsafe.Pointer(this)), uintptr(color), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor) CreateContainerVisual() *IContainerVisual {
	var _result *IContainerVisual
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateContainerVisual, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor) CreateCubicBezierEasingFunction(controlPoint1 unsafe.Pointer, controlPoint2 unsafe.Pointer) *ICubicBezierEasingFunction {
	var _result *ICubicBezierEasingFunction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateCubicBezierEasingFunction, uintptr(unsafe.Pointer(this)), uintptr(controlPoint1), uintptr(controlPoint2), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor) CreateEffectFactory(graphicsEffect unsafe.Pointer) *ICompositionEffectFactory {
	var _result *ICompositionEffectFactory
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateEffectFactory, uintptr(unsafe.Pointer(this)), uintptr(graphicsEffect), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor) CreateEffectFactoryWithProperties(graphicsEffect unsafe.Pointer, animatableProperties *IIterable[string]) *ICompositionEffectFactory {
	var _result *ICompositionEffectFactory
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateEffectFactoryWithProperties, uintptr(unsafe.Pointer(this)), uintptr(graphicsEffect), uintptr(unsafe.Pointer(animatableProperties)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor) CreateExpressionAnimation() *IExpressionAnimation {
	var _result *IExpressionAnimation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateExpressionAnimation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor) CreateExpressionAnimationWithExpression(expression string) *IExpressionAnimation {
	var _result *IExpressionAnimation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateExpressionAnimationWithExpression, uintptr(unsafe.Pointer(this)), NewHStr(expression).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor) CreateInsetClip() *IInsetClip {
	var _result *IInsetClip
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateInsetClip, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor) CreateInsetClipWithInsets(leftInset float32, topInset float32, rightInset float32, bottomInset float32) *IInsetClip {
	var _result *IInsetClip
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateInsetClipWithInsets, uintptr(unsafe.Pointer(this)), uintptr(leftInset), uintptr(topInset), uintptr(rightInset), uintptr(bottomInset), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor) CreateLinearEasingFunction() *ILinearEasingFunction {
	var _result *ILinearEasingFunction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateLinearEasingFunction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor) CreatePropertySet() *ICompositionPropertySet {
	var _result *ICompositionPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreatePropertySet, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor) CreateQuaternionKeyFrameAnimation() *IQuaternionKeyFrameAnimation {
	var _result *IQuaternionKeyFrameAnimation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateQuaternionKeyFrameAnimation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor) CreateScalarKeyFrameAnimation() *IScalarKeyFrameAnimation {
	var _result *IScalarKeyFrameAnimation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateScalarKeyFrameAnimation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor) CreateScopedBatch(batchType CompositionBatchTypes) *ICompositionScopedBatch {
	var _result *ICompositionScopedBatch
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateScopedBatch, uintptr(unsafe.Pointer(this)), uintptr(batchType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor) CreateSpriteVisual() *ISpriteVisual {
	var _result *ISpriteVisual
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSpriteVisual, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor) CreateSurfaceBrush() *ICompositionSurfaceBrush {
	var _result *ICompositionSurfaceBrush
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSurfaceBrush, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor) CreateSurfaceBrushWithSurface(surface *ICompositionSurface) *ICompositionSurfaceBrush {
	var _result *ICompositionSurfaceBrush
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSurfaceBrushWithSurface, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(surface)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor) CreateTargetForCurrentView() *ICompositionTarget {
	var _result *ICompositionTarget
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateTargetForCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor) CreateVector2KeyFrameAnimation() *IVector2KeyFrameAnimation {
	var _result *IVector2KeyFrameAnimation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateVector2KeyFrameAnimation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor) CreateVector3KeyFrameAnimation() *IVector3KeyFrameAnimation {
	var _result *IVector3KeyFrameAnimation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateVector3KeyFrameAnimation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor) CreateVector4KeyFrameAnimation() *IVector4KeyFrameAnimation {
	var _result *IVector4KeyFrameAnimation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateVector4KeyFrameAnimation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor) GetCommitBatch(batchType CompositionBatchTypes) *ICompositionCommitBatch {
	var _result *ICompositionCommitBatch
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCommitBatch, uintptr(unsafe.Pointer(this)), uintptr(batchType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 735081DC-5E24-45DA-A38F-E32CC349A9A0
var IID_ICompositor2 = syscall.GUID{0x735081DC, 0x5E24, 0x45DA,
	[8]byte{0xA3, 0x8F, 0xE3, 0x2C, 0xC3, 0x49, 0xA9, 0xA0}}

type ICompositor2Interface interface {
	win32.IInspectableInterface
	CreateAmbientLight() *IAmbientLight
	CreateAnimationGroup() *ICompositionAnimationGroup
	CreateBackdropBrush() *ICompositionBackdropBrush
	CreateDistantLight() *IDistantLight
	CreateDropShadow() *IDropShadow
	CreateImplicitAnimationCollection() *IImplicitAnimationCollection
	CreateLayerVisual() *ILayerVisual
	CreateMaskBrush() *ICompositionMaskBrush
	CreateNineGridBrush() *ICompositionNineGridBrush
	CreatePointLight() *IPointLight
	CreateSpotLight() *ISpotLight
	CreateStepEasingFunction() *IStepEasingFunction
	CreateStepEasingFunctionWithStepCount(stepCount int32) *IStepEasingFunction
}

type ICompositor2Vtbl struct {
	win32.IInspectableVtbl
	CreateAmbientLight                    uintptr
	CreateAnimationGroup                  uintptr
	CreateBackdropBrush                   uintptr
	CreateDistantLight                    uintptr
	CreateDropShadow                      uintptr
	CreateImplicitAnimationCollection     uintptr
	CreateLayerVisual                     uintptr
	CreateMaskBrush                       uintptr
	CreateNineGridBrush                   uintptr
	CreatePointLight                      uintptr
	CreateSpotLight                       uintptr
	CreateStepEasingFunction              uintptr
	CreateStepEasingFunctionWithStepCount uintptr
}

type ICompositor2 struct {
	win32.IInspectable
}

func (this *ICompositor2) Vtbl() *ICompositor2Vtbl {
	return (*ICompositor2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositor2) CreateAmbientLight() *IAmbientLight {
	var _result *IAmbientLight
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAmbientLight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor2) CreateAnimationGroup() *ICompositionAnimationGroup {
	var _result *ICompositionAnimationGroup
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAnimationGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor2) CreateBackdropBrush() *ICompositionBackdropBrush {
	var _result *ICompositionBackdropBrush
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateBackdropBrush, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor2) CreateDistantLight() *IDistantLight {
	var _result *IDistantLight
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDistantLight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor2) CreateDropShadow() *IDropShadow {
	var _result *IDropShadow
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDropShadow, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor2) CreateImplicitAnimationCollection() *IImplicitAnimationCollection {
	var _result *IImplicitAnimationCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateImplicitAnimationCollection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor2) CreateLayerVisual() *ILayerVisual {
	var _result *ILayerVisual
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateLayerVisual, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor2) CreateMaskBrush() *ICompositionMaskBrush {
	var _result *ICompositionMaskBrush
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateMaskBrush, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor2) CreateNineGridBrush() *ICompositionNineGridBrush {
	var _result *ICompositionNineGridBrush
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateNineGridBrush, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor2) CreatePointLight() *IPointLight {
	var _result *IPointLight
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreatePointLight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor2) CreateSpotLight() *ISpotLight {
	var _result *ISpotLight
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSpotLight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor2) CreateStepEasingFunction() *IStepEasingFunction {
	var _result *IStepEasingFunction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateStepEasingFunction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor2) CreateStepEasingFunctionWithStepCount(stepCount int32) *IStepEasingFunction {
	var _result *IStepEasingFunction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateStepEasingFunctionWithStepCount, uintptr(unsafe.Pointer(this)), uintptr(stepCount), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C9DD8EF0-6EB1-4E3C-A658-675D9C64D4AB
var IID_ICompositor3 = syscall.GUID{0xC9DD8EF0, 0x6EB1, 0x4E3C,
	[8]byte{0xA6, 0x58, 0x67, 0x5D, 0x9C, 0x64, 0xD4, 0xAB}}

type ICompositor3Interface interface {
	win32.IInspectableInterface
	CreateHostBackdropBrush() *ICompositionBackdropBrush
}

type ICompositor3Vtbl struct {
	win32.IInspectableVtbl
	CreateHostBackdropBrush uintptr
}

type ICompositor3 struct {
	win32.IInspectable
}

func (this *ICompositor3) Vtbl() *ICompositor3Vtbl {
	return (*ICompositor3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositor3) CreateHostBackdropBrush() *ICompositionBackdropBrush {
	var _result *ICompositionBackdropBrush
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateHostBackdropBrush, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AE47E78A-7910-4425-A482-A05B758ADCE9
var IID_ICompositor4 = syscall.GUID{0xAE47E78A, 0x7910, 0x4425,
	[8]byte{0xA4, 0x82, 0xA0, 0x5B, 0x75, 0x8A, 0xDC, 0xE9}}

type ICompositor4Interface interface {
	win32.IInspectableInterface
	CreateColorGradientStop() *ICompositionColorGradientStop
	CreateColorGradientStopWithOffsetAndColor(offset float32, color unsafe.Pointer) *ICompositionColorGradientStop
	CreateLinearGradientBrush() *ICompositionLinearGradientBrush
	CreateSpringScalarAnimation() *ISpringScalarNaturalMotionAnimation
	CreateSpringVector2Animation() *ISpringVector2NaturalMotionAnimation
	CreateSpringVector3Animation() *ISpringVector3NaturalMotionAnimation
}

type ICompositor4Vtbl struct {
	win32.IInspectableVtbl
	CreateColorGradientStop                   uintptr
	CreateColorGradientStopWithOffsetAndColor uintptr
	CreateLinearGradientBrush                 uintptr
	CreateSpringScalarAnimation               uintptr
	CreateSpringVector2Animation              uintptr
	CreateSpringVector3Animation              uintptr
}

type ICompositor4 struct {
	win32.IInspectable
}

func (this *ICompositor4) Vtbl() *ICompositor4Vtbl {
	return (*ICompositor4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositor4) CreateColorGradientStop() *ICompositionColorGradientStop {
	var _result *ICompositionColorGradientStop
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateColorGradientStop, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor4) CreateColorGradientStopWithOffsetAndColor(offset float32, color unsafe.Pointer) *ICompositionColorGradientStop {
	var _result *ICompositionColorGradientStop
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateColorGradientStopWithOffsetAndColor, uintptr(unsafe.Pointer(this)), uintptr(offset), uintptr(color), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor4) CreateLinearGradientBrush() *ICompositionLinearGradientBrush {
	var _result *ICompositionLinearGradientBrush
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateLinearGradientBrush, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor4) CreateSpringScalarAnimation() *ISpringScalarNaturalMotionAnimation {
	var _result *ISpringScalarNaturalMotionAnimation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSpringScalarAnimation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor4) CreateSpringVector2Animation() *ISpringVector2NaturalMotionAnimation {
	var _result *ISpringVector2NaturalMotionAnimation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSpringVector2Animation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor4) CreateSpringVector3Animation() *ISpringVector3NaturalMotionAnimation {
	var _result *ISpringVector3NaturalMotionAnimation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSpringVector3Animation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 48EA31AD-7FCD-4076-A79C-90CC4B852C9B
var IID_ICompositor5 = syscall.GUID{0x48EA31AD, 0x7FCD, 0x4076,
	[8]byte{0xA7, 0x9C, 0x90, 0xCC, 0x4B, 0x85, 0x2C, 0x9B}}

type ICompositor5Interface interface {
	win32.IInspectableInterface
	Get_Comment() string
	Put_Comment(value string)
	Get_GlobalPlaybackRate() float32
	Put_GlobalPlaybackRate(value float32)
	CreateBounceScalarAnimation() *IBounceScalarNaturalMotionAnimation
	CreateBounceVector2Animation() *IBounceVector2NaturalMotionAnimation
	CreateBounceVector3Animation() *IBounceVector3NaturalMotionAnimation
	CreateContainerShape() *ICompositionContainerShape
	CreateEllipseGeometry() *ICompositionEllipseGeometry
	CreateLineGeometry() *ICompositionLineGeometry
	CreatePathGeometry() *ICompositionPathGeometry
	CreatePathGeometryWithPath(path *ICompositionPath) *ICompositionPathGeometry
	CreatePathKeyFrameAnimation() *IPathKeyFrameAnimation
	CreateRectangleGeometry() *ICompositionRectangleGeometry
	CreateRoundedRectangleGeometry() *ICompositionRoundedRectangleGeometry
	CreateShapeVisual() *IShapeVisual
	CreateSpriteShape() *ICompositionSpriteShape
	CreateSpriteShapeWithGeometry(geometry *ICompositionGeometry) *ICompositionSpriteShape
	CreateViewBox() *ICompositionViewBox
	RequestCommitAsync() *IAsyncAction
}

type ICompositor5Vtbl struct {
	win32.IInspectableVtbl
	Get_Comment                    uintptr
	Put_Comment                    uintptr
	Get_GlobalPlaybackRate         uintptr
	Put_GlobalPlaybackRate         uintptr
	CreateBounceScalarAnimation    uintptr
	CreateBounceVector2Animation   uintptr
	CreateBounceVector3Animation   uintptr
	CreateContainerShape           uintptr
	CreateEllipseGeometry          uintptr
	CreateLineGeometry             uintptr
	CreatePathGeometry             uintptr
	CreatePathGeometryWithPath     uintptr
	CreatePathKeyFrameAnimation    uintptr
	CreateRectangleGeometry        uintptr
	CreateRoundedRectangleGeometry uintptr
	CreateShapeVisual              uintptr
	CreateSpriteShape              uintptr
	CreateSpriteShapeWithGeometry  uintptr
	CreateViewBox                  uintptr
	RequestCommitAsync             uintptr
}

type ICompositor5 struct {
	win32.IInspectable
}

func (this *ICompositor5) Vtbl() *ICompositor5Vtbl {
	return (*ICompositor5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositor5) Get_Comment() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Comment, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ICompositor5) Put_Comment(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Comment, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ICompositor5) Get_GlobalPlaybackRate() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GlobalPlaybackRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositor5) Put_GlobalPlaybackRate(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_GlobalPlaybackRate, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ICompositor5) CreateBounceScalarAnimation() *IBounceScalarNaturalMotionAnimation {
	var _result *IBounceScalarNaturalMotionAnimation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateBounceScalarAnimation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor5) CreateBounceVector2Animation() *IBounceVector2NaturalMotionAnimation {
	var _result *IBounceVector2NaturalMotionAnimation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateBounceVector2Animation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor5) CreateBounceVector3Animation() *IBounceVector3NaturalMotionAnimation {
	var _result *IBounceVector3NaturalMotionAnimation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateBounceVector3Animation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor5) CreateContainerShape() *ICompositionContainerShape {
	var _result *ICompositionContainerShape
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateContainerShape, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor5) CreateEllipseGeometry() *ICompositionEllipseGeometry {
	var _result *ICompositionEllipseGeometry
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateEllipseGeometry, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor5) CreateLineGeometry() *ICompositionLineGeometry {
	var _result *ICompositionLineGeometry
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateLineGeometry, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor5) CreatePathGeometry() *ICompositionPathGeometry {
	var _result *ICompositionPathGeometry
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreatePathGeometry, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor5) CreatePathGeometryWithPath(path *ICompositionPath) *ICompositionPathGeometry {
	var _result *ICompositionPathGeometry
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreatePathGeometryWithPath, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(path)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor5) CreatePathKeyFrameAnimation() *IPathKeyFrameAnimation {
	var _result *IPathKeyFrameAnimation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreatePathKeyFrameAnimation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor5) CreateRectangleGeometry() *ICompositionRectangleGeometry {
	var _result *ICompositionRectangleGeometry
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateRectangleGeometry, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor5) CreateRoundedRectangleGeometry() *ICompositionRoundedRectangleGeometry {
	var _result *ICompositionRoundedRectangleGeometry
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateRoundedRectangleGeometry, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor5) CreateShapeVisual() *IShapeVisual {
	var _result *IShapeVisual
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateShapeVisual, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor5) CreateSpriteShape() *ICompositionSpriteShape {
	var _result *ICompositionSpriteShape
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSpriteShape, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor5) CreateSpriteShapeWithGeometry(geometry *ICompositionGeometry) *ICompositionSpriteShape {
	var _result *ICompositionSpriteShape
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSpriteShapeWithGeometry, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(geometry)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor5) CreateViewBox() *ICompositionViewBox {
	var _result *ICompositionViewBox
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateViewBox, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor5) RequestCommitAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestCommitAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7A38B2BD-CEC8-4EEB-830F-D8D07AEDEBC3
var IID_ICompositor6 = syscall.GUID{0x7A38B2BD, 0xCEC8, 0x4EEB,
	[8]byte{0x83, 0x0F, 0xD8, 0xD0, 0x7A, 0xED, 0xEB, 0xC3}}

type ICompositor6Interface interface {
	win32.IInspectableInterface
	CreateGeometricClip() *ICompositionGeometricClip
	CreateGeometricClipWithGeometry(geometry *ICompositionGeometry) *ICompositionGeometricClip
	CreateRedirectVisual() *IRedirectVisual
	CreateRedirectVisualWithSourceVisual(source *IVisual) *IRedirectVisual
	CreateBooleanKeyFrameAnimation() *IBooleanKeyFrameAnimation
}

type ICompositor6Vtbl struct {
	win32.IInspectableVtbl
	CreateGeometricClip                  uintptr
	CreateGeometricClipWithGeometry      uintptr
	CreateRedirectVisual                 uintptr
	CreateRedirectVisualWithSourceVisual uintptr
	CreateBooleanKeyFrameAnimation       uintptr
}

type ICompositor6 struct {
	win32.IInspectable
}

func (this *ICompositor6) Vtbl() *ICompositor6Vtbl {
	return (*ICompositor6Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositor6) CreateGeometricClip() *ICompositionGeometricClip {
	var _result *ICompositionGeometricClip
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateGeometricClip, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor6) CreateGeometricClipWithGeometry(geometry *ICompositionGeometry) *ICompositionGeometricClip {
	var _result *ICompositionGeometricClip
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateGeometricClipWithGeometry, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(geometry)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor6) CreateRedirectVisual() *IRedirectVisual {
	var _result *IRedirectVisual
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateRedirectVisual, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor6) CreateRedirectVisualWithSourceVisual(source *IVisual) *IRedirectVisual {
	var _result *IRedirectVisual
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateRedirectVisualWithSourceVisual, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(source)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor6) CreateBooleanKeyFrameAnimation() *IBooleanKeyFrameAnimation {
	var _result *IBooleanKeyFrameAnimation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateBooleanKeyFrameAnimation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D3483FAD-9A12-53BA-BFC8-88B7FF7977C6
var IID_ICompositor7 = syscall.GUID{0xD3483FAD, 0x9A12, 0x53BA,
	[8]byte{0xBF, 0xC8, 0x88, 0xB7, 0xFF, 0x79, 0x77, 0xC6}}

type ICompositor7Interface interface {
	win32.IInspectableInterface
	Get_DispatcherQueue() *IDispatcherQueue
	CreateAnimationPropertyInfo() *IAnimationPropertyInfo
	CreateRectangleClip() *IRectangleClip
	CreateRectangleClipWithSides(left float32, top float32, right float32, bottom float32) *IRectangleClip
	CreateRectangleClipWithSidesAndRadius(left float32, top float32, right float32, bottom float32, topLeftRadius unsafe.Pointer, topRightRadius unsafe.Pointer, bottomRightRadius unsafe.Pointer, bottomLeftRadius unsafe.Pointer) *IRectangleClip
}

type ICompositor7Vtbl struct {
	win32.IInspectableVtbl
	Get_DispatcherQueue                   uintptr
	CreateAnimationPropertyInfo           uintptr
	CreateRectangleClip                   uintptr
	CreateRectangleClipWithSides          uintptr
	CreateRectangleClipWithSidesAndRadius uintptr
}

type ICompositor7 struct {
	win32.IInspectable
}

func (this *ICompositor7) Vtbl() *ICompositor7Vtbl {
	return (*ICompositor7Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositor7) Get_DispatcherQueue() *IDispatcherQueue {
	var _result *IDispatcherQueue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DispatcherQueue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor7) CreateAnimationPropertyInfo() *IAnimationPropertyInfo {
	var _result *IAnimationPropertyInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAnimationPropertyInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor7) CreateRectangleClip() *IRectangleClip {
	var _result *IRectangleClip
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateRectangleClip, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor7) CreateRectangleClipWithSides(left float32, top float32, right float32, bottom float32) *IRectangleClip {
	var _result *IRectangleClip
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateRectangleClipWithSides, uintptr(unsafe.Pointer(this)), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositor7) CreateRectangleClipWithSidesAndRadius(left float32, top float32, right float32, bottom float32, topLeftRadius unsafe.Pointer, topRightRadius unsafe.Pointer, bottomRightRadius unsafe.Pointer, bottomLeftRadius unsafe.Pointer) *IRectangleClip {
	var _result *IRectangleClip
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateRectangleClipWithSidesAndRadius, uintptr(unsafe.Pointer(this)), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom), uintptr(topLeftRadius), uintptr(topRightRadius), uintptr(bottomRightRadius), uintptr(bottomLeftRadius), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9A0BDEE2-FE7B-5F62-A366-9CF8EFFE2112
var IID_ICompositor8 = syscall.GUID{0x9A0BDEE2, 0xFE7B, 0x5F62,
	[8]byte{0xA3, 0x66, 0x9C, 0xF8, 0xEF, 0xFE, 0x21, 0x12}}

type ICompositor8Interface interface {
	win32.IInspectableInterface
	CreateAnimationController() *IAnimationController
}

type ICompositor8Vtbl struct {
	win32.IInspectableVtbl
	CreateAnimationController uintptr
}

type ICompositor8 struct {
	win32.IInspectable
}

func (this *ICompositor8) Vtbl() *ICompositor8Vtbl {
	return (*ICompositor8Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositor8) CreateAnimationController() *IAnimationController {
	var _result *IAnimationController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAnimationController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 080DB93E-121E-4D97-8B74-1DFCF91987EA
var IID_ICompositorStatics = syscall.GUID{0x080DB93E, 0x121E, 0x4D97,
	[8]byte{0x8B, 0x74, 0x1D, 0xFC, 0xF9, 0x19, 0x87, 0xEA}}

type ICompositorStaticsInterface interface {
	win32.IInspectableInterface
	Get_MaxGlobalPlaybackRate() float32
	Get_MinGlobalPlaybackRate() float32
}

type ICompositorStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_MaxGlobalPlaybackRate uintptr
	Get_MinGlobalPlaybackRate uintptr
}

type ICompositorStatics struct {
	win32.IInspectable
}

func (this *ICompositorStatics) Vtbl() *ICompositorStaticsVtbl {
	return (*ICompositorStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositorStatics) Get_MaxGlobalPlaybackRate() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxGlobalPlaybackRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositorStatics) Get_MinGlobalPlaybackRate() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinGlobalPlaybackRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 0D8FB190-F122-5B8D-9FDD-543B0D8EB7F3
var IID_ICompositorWithBlurredWallpaperBackdropBrush = syscall.GUID{0x0D8FB190, 0xF122, 0x5B8D,
	[8]byte{0x9F, 0xDD, 0x54, 0x3B, 0x0D, 0x8E, 0xB7, 0xF3}}

type ICompositorWithBlurredWallpaperBackdropBrushInterface interface {
	win32.IInspectableInterface
	TryCreateBlurredWallpaperBackdropBrush() *ICompositionBackdropBrush
}

type ICompositorWithBlurredWallpaperBackdropBrushVtbl struct {
	win32.IInspectableVtbl
	TryCreateBlurredWallpaperBackdropBrush uintptr
}

type ICompositorWithBlurredWallpaperBackdropBrush struct {
	win32.IInspectable
}

func (this *ICompositorWithBlurredWallpaperBackdropBrush) Vtbl() *ICompositorWithBlurredWallpaperBackdropBrushVtbl {
	return (*ICompositorWithBlurredWallpaperBackdropBrushVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositorWithBlurredWallpaperBackdropBrush) TryCreateBlurredWallpaperBackdropBrush() *ICompositionBackdropBrush {
	var _result *ICompositionBackdropBrush
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryCreateBlurredWallpaperBackdropBrush, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A2E6330E-8A60-5A38-BB85-B44EA901677C
var IID_ICompositorWithProjectedShadow = syscall.GUID{0xA2E6330E, 0x8A60, 0x5A38,
	[8]byte{0xBB, 0x85, 0xB4, 0x4E, 0xA9, 0x01, 0x67, 0x7C}}

type ICompositorWithProjectedShadowInterface interface {
	win32.IInspectableInterface
	CreateProjectedShadowCaster() *ICompositionProjectedShadowCaster
	CreateProjectedShadow() *ICompositionProjectedShadow
	CreateProjectedShadowReceiver() *ICompositionProjectedShadowReceiver
}

type ICompositorWithProjectedShadowVtbl struct {
	win32.IInspectableVtbl
	CreateProjectedShadowCaster   uintptr
	CreateProjectedShadow         uintptr
	CreateProjectedShadowReceiver uintptr
}

type ICompositorWithProjectedShadow struct {
	win32.IInspectable
}

func (this *ICompositorWithProjectedShadow) Vtbl() *ICompositorWithProjectedShadowVtbl {
	return (*ICompositorWithProjectedShadowVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositorWithProjectedShadow) CreateProjectedShadowCaster() *ICompositionProjectedShadowCaster {
	var _result *ICompositionProjectedShadowCaster
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateProjectedShadowCaster, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositorWithProjectedShadow) CreateProjectedShadow() *ICompositionProjectedShadow {
	var _result *ICompositionProjectedShadow
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateProjectedShadow, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositorWithProjectedShadow) CreateProjectedShadowReceiver() *ICompositionProjectedShadowReceiver {
	var _result *ICompositionProjectedShadowReceiver
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateProjectedShadowReceiver, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 98B9C1A7-8E71-4B53-B4A8-69BA5D19DC5B
var IID_ICompositorWithRadialGradient = syscall.GUID{0x98B9C1A7, 0x8E71, 0x4B53,
	[8]byte{0xB4, 0xA8, 0x69, 0xBA, 0x5D, 0x19, 0xDC, 0x5B}}

type ICompositorWithRadialGradientInterface interface {
	win32.IInspectableInterface
	CreateRadialGradientBrush() *ICompositionRadialGradientBrush
}

type ICompositorWithRadialGradientVtbl struct {
	win32.IInspectableVtbl
	CreateRadialGradientBrush uintptr
}

type ICompositorWithRadialGradient struct {
	win32.IInspectable
}

func (this *ICompositorWithRadialGradient) Vtbl() *ICompositorWithRadialGradientVtbl {
	return (*ICompositorWithRadialGradientVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositorWithRadialGradient) CreateRadialGradientBrush() *ICompositionRadialGradientBrush {
	var _result *ICompositionRadialGradientBrush
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateRadialGradientBrush, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CFA1658B-0123-4551-8891-89BDCC40322B
var IID_ICompositorWithVisualSurface = syscall.GUID{0xCFA1658B, 0x0123, 0x4551,
	[8]byte{0x88, 0x91, 0x89, 0xBD, 0xCC, 0x40, 0x32, 0x2B}}

type ICompositorWithVisualSurfaceInterface interface {
	win32.IInspectableInterface
	CreateVisualSurface() *ICompositionVisualSurface
}

type ICompositorWithVisualSurfaceVtbl struct {
	win32.IInspectableVtbl
	CreateVisualSurface uintptr
}

type ICompositorWithVisualSurface struct {
	win32.IInspectable
}

func (this *ICompositorWithVisualSurface) Vtbl() *ICompositorWithVisualSurfaceVtbl {
	return (*ICompositorWithVisualSurfaceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositorWithVisualSurface) CreateVisualSurface() *ICompositionVisualSurface {
	var _result *ICompositionVisualSurface
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateVisualSurface, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 02F6BC74-ED20-4773-AFE6-D49B4A93DB32
var IID_IContainerVisual = syscall.GUID{0x02F6BC74, 0xED20, 0x4773,
	[8]byte{0xAF, 0xE6, 0xD4, 0x9B, 0x4A, 0x93, 0xDB, 0x32}}

type IContainerVisualInterface interface {
	win32.IInspectableInterface
	Get_Children() *IVisualCollection
}

type IContainerVisualVtbl struct {
	win32.IInspectableVtbl
	Get_Children uintptr
}

type IContainerVisual struct {
	win32.IInspectable
}

func (this *IContainerVisual) Vtbl() *IContainerVisualVtbl {
	return (*IContainerVisualVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IContainerVisual) Get_Children() *IVisualCollection {
	var _result *IVisualCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Children, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0363A65B-C7DA-4D9A-95F4-69B5C8DF670B
var IID_IContainerVisualFactory = syscall.GUID{0x0363A65B, 0xC7DA, 0x4D9A,
	[8]byte{0x95, 0xF4, 0x69, 0xB5, 0xC8, 0xDF, 0x67, 0x0B}}

type IContainerVisualFactoryInterface interface {
	win32.IInspectableInterface
}

type IContainerVisualFactoryVtbl struct {
	win32.IInspectableVtbl
}

type IContainerVisualFactory struct {
	win32.IInspectable
}

func (this *IContainerVisualFactory) Vtbl() *IContainerVisualFactoryVtbl {
	return (*IContainerVisualFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 32350666-C1E8-44F9-96B8-C98ACF0AE698
var IID_ICubicBezierEasingFunction = syscall.GUID{0x32350666, 0xC1E8, 0x44F9,
	[8]byte{0x96, 0xB8, 0xC9, 0x8A, 0xCF, 0x0A, 0xE6, 0x98}}

type ICubicBezierEasingFunctionInterface interface {
	win32.IInspectableInterface
	Get_ControlPoint1() unsafe.Pointer
	Get_ControlPoint2() unsafe.Pointer
}

type ICubicBezierEasingFunctionVtbl struct {
	win32.IInspectableVtbl
	Get_ControlPoint1 uintptr
	Get_ControlPoint2 uintptr
}

type ICubicBezierEasingFunction struct {
	win32.IInspectable
}

func (this *ICubicBezierEasingFunction) Vtbl() *ICubicBezierEasingFunctionVtbl {
	return (*ICubicBezierEasingFunctionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICubicBezierEasingFunction) Get_ControlPoint1() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ControlPoint1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICubicBezierEasingFunction) Get_ControlPoint2() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ControlPoint2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 856E60B1-E1AB-5B23-8E3D-D513F221C998
var IID_IDelegatedInkTrailVisual = syscall.GUID{0x856E60B1, 0xE1AB, 0x5B23,
	[8]byte{0x8E, 0x3D, 0xD5, 0x13, 0xF2, 0x21, 0xC9, 0x98}}

type IDelegatedInkTrailVisualInterface interface {
	win32.IInspectableInterface
	AddTrailPoints(inkPointsLength uint32, inkPoints *InkTrailPoint) uint32
	AddTrailPointsWithPrediction(inkPointsLength uint32, inkPoints *InkTrailPoint, predictedInkPointsLength uint32, predictedInkPoints *InkTrailPoint) uint32
	RemoveTrailPoints(generationId uint32)
	StartNewTrail(color unsafe.Pointer)
}

type IDelegatedInkTrailVisualVtbl struct {
	win32.IInspectableVtbl
	AddTrailPoints               uintptr
	AddTrailPointsWithPrediction uintptr
	RemoveTrailPoints            uintptr
	StartNewTrail                uintptr
}

type IDelegatedInkTrailVisual struct {
	win32.IInspectable
}

func (this *IDelegatedInkTrailVisual) Vtbl() *IDelegatedInkTrailVisualVtbl {
	return (*IDelegatedInkTrailVisualVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDelegatedInkTrailVisual) AddTrailPoints(inkPointsLength uint32, inkPoints *InkTrailPoint) uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddTrailPoints, uintptr(unsafe.Pointer(this)), uintptr(inkPointsLength), uintptr(unsafe.Pointer(inkPoints)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDelegatedInkTrailVisual) AddTrailPointsWithPrediction(inkPointsLength uint32, inkPoints *InkTrailPoint, predictedInkPointsLength uint32, predictedInkPoints *InkTrailPoint) uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddTrailPointsWithPrediction, uintptr(unsafe.Pointer(this)), uintptr(inkPointsLength), uintptr(unsafe.Pointer(inkPoints)), uintptr(predictedInkPointsLength), uintptr(unsafe.Pointer(predictedInkPoints)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDelegatedInkTrailVisual) RemoveTrailPoints(generationId uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveTrailPoints, uintptr(unsafe.Pointer(this)), uintptr(generationId))
	_ = _hr
}

func (this *IDelegatedInkTrailVisual) StartNewTrail(color unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartNewTrail, uintptr(unsafe.Pointer(this)), uintptr(color))
	_ = _hr
}

// 0DAF6BD5-42C6-555C-9267-E0AC663AF836
var IID_IDelegatedInkTrailVisualStatics = syscall.GUID{0x0DAF6BD5, 0x42C6, 0x555C,
	[8]byte{0x92, 0x67, 0xE0, 0xAC, 0x66, 0x3A, 0xF8, 0x36}}

type IDelegatedInkTrailVisualStaticsInterface interface {
	win32.IInspectableInterface
	Create(compositor *ICompositor) *IDelegatedInkTrailVisual
	CreateForSwapChain(compositor *ICompositor, swapChain *ICompositionSurface) *IDelegatedInkTrailVisual
}

type IDelegatedInkTrailVisualStaticsVtbl struct {
	win32.IInspectableVtbl
	Create             uintptr
	CreateForSwapChain uintptr
}

type IDelegatedInkTrailVisualStatics struct {
	win32.IInspectable
}

func (this *IDelegatedInkTrailVisualStatics) Vtbl() *IDelegatedInkTrailVisualStaticsVtbl {
	return (*IDelegatedInkTrailVisualStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDelegatedInkTrailVisualStatics) Create(compositor *ICompositor) *IDelegatedInkTrailVisual {
	var _result *IDelegatedInkTrailVisual
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(compositor)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDelegatedInkTrailVisualStatics) CreateForSwapChain(compositor *ICompositor, swapChain *ICompositionSurface) *IDelegatedInkTrailVisual {
	var _result *IDelegatedInkTrailVisual
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateForSwapChain, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(compositor)), uintptr(unsafe.Pointer(swapChain)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 318CFAFC-5CE3-4B55-AB5D-07A00353AC99
var IID_IDistantLight = syscall.GUID{0x318CFAFC, 0x5CE3, 0x4B55,
	[8]byte{0xAB, 0x5D, 0x07, 0xA0, 0x03, 0x53, 0xAC, 0x99}}

type IDistantLightInterface interface {
	win32.IInspectableInterface
	Get_Color() unsafe.Pointer
	Put_Color(value unsafe.Pointer)
	Get_CoordinateSpace() *IVisual
	Put_CoordinateSpace(value *IVisual)
	Get_Direction() unsafe.Pointer
	Put_Direction(value unsafe.Pointer)
}

type IDistantLightVtbl struct {
	win32.IInspectableVtbl
	Get_Color           uintptr
	Put_Color           uintptr
	Get_CoordinateSpace uintptr
	Put_CoordinateSpace uintptr
	Get_Direction       uintptr
	Put_Direction       uintptr
}

type IDistantLight struct {
	win32.IInspectable
}

func (this *IDistantLight) Vtbl() *IDistantLightVtbl {
	return (*IDistantLightVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDistantLight) Get_Color() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Color, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDistantLight) Put_Color(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Color, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDistantLight) Get_CoordinateSpace() *IVisual {
	var _result *IVisual
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CoordinateSpace, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDistantLight) Put_CoordinateSpace(value *IVisual) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CoordinateSpace, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IDistantLight) Get_Direction() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Direction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDistantLight) Put_Direction(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Direction, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// DBCDAA1C-294B-48D7-B60E-76DF64AA392B
var IID_IDistantLight2 = syscall.GUID{0xDBCDAA1C, 0x294B, 0x48D7,
	[8]byte{0xB6, 0x0E, 0x76, 0xDF, 0x64, 0xAA, 0x39, 0x2B}}

type IDistantLight2Interface interface {
	win32.IInspectableInterface
	Get_Intensity() float32
	Put_Intensity(value float32)
}

type IDistantLight2Vtbl struct {
	win32.IInspectableVtbl
	Get_Intensity uintptr
	Put_Intensity uintptr
}

type IDistantLight2 struct {
	win32.IInspectable
}

func (this *IDistantLight2) Vtbl() *IDistantLight2Vtbl {
	return (*IDistantLight2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDistantLight2) Get_Intensity() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Intensity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDistantLight2) Put_Intensity(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Intensity, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// CB977C07-A154-4851-85E7-A8924C84FAD8
var IID_IDropShadow = syscall.GUID{0xCB977C07, 0xA154, 0x4851,
	[8]byte{0x85, 0xE7, 0xA8, 0x92, 0x4C, 0x84, 0xFA, 0xD8}}

type IDropShadowInterface interface {
	win32.IInspectableInterface
	Get_BlurRadius() float32
	Put_BlurRadius(value float32)
	Get_Color() unsafe.Pointer
	Put_Color(value unsafe.Pointer)
	Get_Mask() *ICompositionBrush
	Put_Mask(value *ICompositionBrush)
	Get_Offset() unsafe.Pointer
	Put_Offset(value unsafe.Pointer)
	Get_Opacity() float32
	Put_Opacity(value float32)
}

type IDropShadowVtbl struct {
	win32.IInspectableVtbl
	Get_BlurRadius uintptr
	Put_BlurRadius uintptr
	Get_Color      uintptr
	Put_Color      uintptr
	Get_Mask       uintptr
	Put_Mask       uintptr
	Get_Offset     uintptr
	Put_Offset     uintptr
	Get_Opacity    uintptr
	Put_Opacity    uintptr
}

type IDropShadow struct {
	win32.IInspectable
}

func (this *IDropShadow) Vtbl() *IDropShadowVtbl {
	return (*IDropShadowVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDropShadow) Get_BlurRadius() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BlurRadius, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDropShadow) Put_BlurRadius(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BlurRadius, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDropShadow) Get_Color() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Color, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDropShadow) Put_Color(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Color, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDropShadow) Get_Mask() *ICompositionBrush {
	var _result *ICompositionBrush
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mask, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDropShadow) Put_Mask(value *ICompositionBrush) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Mask, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IDropShadow) Get_Offset() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Offset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDropShadow) Put_Offset(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Offset, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDropShadow) Get_Opacity() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Opacity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDropShadow) Put_Opacity(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Opacity, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 6C4218BC-15B9-4C2D-8D4A-0767DF11977A
var IID_IDropShadow2 = syscall.GUID{0x6C4218BC, 0x15B9, 0x4C2D,
	[8]byte{0x8D, 0x4A, 0x07, 0x67, 0xDF, 0x11, 0x97, 0x7A}}

type IDropShadow2Interface interface {
	win32.IInspectableInterface
	Get_SourcePolicy() CompositionDropShadowSourcePolicy
	Put_SourcePolicy(value CompositionDropShadowSourcePolicy)
}

type IDropShadow2Vtbl struct {
	win32.IInspectableVtbl
	Get_SourcePolicy uintptr
	Put_SourcePolicy uintptr
}

type IDropShadow2 struct {
	win32.IInspectable
}

func (this *IDropShadow2) Vtbl() *IDropShadow2Vtbl {
	return (*IDropShadow2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDropShadow2) Get_SourcePolicy() CompositionDropShadowSourcePolicy {
	var _result CompositionDropShadowSourcePolicy
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SourcePolicy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDropShadow2) Put_SourcePolicy(value CompositionDropShadowSourcePolicy) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SourcePolicy, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 66DE6285-054E-5594-8475-C22CB51F1BD5
var IID_IElasticEasingFunction = syscall.GUID{0x66DE6285, 0x054E, 0x5594,
	[8]byte{0x84, 0x75, 0xC2, 0x2C, 0xB5, 0x1F, 0x1B, 0xD5}}

type IElasticEasingFunctionInterface interface {
	win32.IInspectableInterface
	Get_Mode() CompositionEasingFunctionMode
	Get_Oscillations() int32
	Get_Springiness() float32
}

type IElasticEasingFunctionVtbl struct {
	win32.IInspectableVtbl
	Get_Mode         uintptr
	Get_Oscillations uintptr
	Get_Springiness  uintptr
}

type IElasticEasingFunction struct {
	win32.IInspectable
}

func (this *IElasticEasingFunction) Vtbl() *IElasticEasingFunctionVtbl {
	return (*IElasticEasingFunctionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IElasticEasingFunction) Get_Mode() CompositionEasingFunctionMode {
	var _result CompositionEasingFunctionMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IElasticEasingFunction) Get_Oscillations() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Oscillations, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IElasticEasingFunction) Get_Springiness() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Springiness, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 6F7D1A51-98D2-5638-A34A-00486554C750
var IID_IExponentialEasingFunction = syscall.GUID{0x6F7D1A51, 0x98D2, 0x5638,
	[8]byte{0xA3, 0x4A, 0x00, 0x48, 0x65, 0x54, 0xC7, 0x50}}

type IExponentialEasingFunctionInterface interface {
	win32.IInspectableInterface
	Get_Mode() CompositionEasingFunctionMode
	Get_Exponent() float32
}

type IExponentialEasingFunctionVtbl struct {
	win32.IInspectableVtbl
	Get_Mode     uintptr
	Get_Exponent uintptr
}

type IExponentialEasingFunction struct {
	win32.IInspectable
}

func (this *IExponentialEasingFunction) Vtbl() *IExponentialEasingFunctionVtbl {
	return (*IExponentialEasingFunctionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IExponentialEasingFunction) Get_Mode() CompositionEasingFunctionMode {
	var _result CompositionEasingFunctionMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IExponentialEasingFunction) Get_Exponent() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Exponent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 6ACC5431-7D3D-4BF3-ABB6-F44BDC4888C1
var IID_IExpressionAnimation = syscall.GUID{0x6ACC5431, 0x7D3D, 0x4BF3,
	[8]byte{0xAB, 0xB6, 0xF4, 0x4B, 0xDC, 0x48, 0x88, 0xC1}}

type IExpressionAnimationInterface interface {
	win32.IInspectableInterface
	Get_Expression() string
	Put_Expression(value string)
}

type IExpressionAnimationVtbl struct {
	win32.IInspectableVtbl
	Get_Expression uintptr
	Put_Expression uintptr
}

type IExpressionAnimation struct {
	win32.IInspectable
}

func (this *IExpressionAnimation) Vtbl() *IExpressionAnimationVtbl {
	return (*IExpressionAnimationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IExpressionAnimation) Get_Expression() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Expression, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IExpressionAnimation) Put_Expression(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Expression, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// 0598A3FF-0A92-4C9D-A427-B25519250DBF
var IID_IImplicitAnimationCollection = syscall.GUID{0x0598A3FF, 0x0A92, 0x4C9D,
	[8]byte{0xA4, 0x27, 0xB2, 0x55, 0x19, 0x25, 0x0D, 0xBF}}

type IImplicitAnimationCollectionInterface interface {
	win32.IInspectableInterface
}

type IImplicitAnimationCollectionVtbl struct {
	win32.IInspectableVtbl
}

type IImplicitAnimationCollection struct {
	win32.IInspectable
}

func (this *IImplicitAnimationCollection) Vtbl() *IImplicitAnimationCollectionVtbl {
	return (*IImplicitAnimationCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 1E73E647-84C7-477A-B474-5880E0442E15
var IID_IInsetClip = syscall.GUID{0x1E73E647, 0x84C7, 0x477A,
	[8]byte{0xB4, 0x74, 0x58, 0x80, 0xE0, 0x44, 0x2E, 0x15}}

type IInsetClipInterface interface {
	win32.IInspectableInterface
	Get_BottomInset() float32
	Put_BottomInset(value float32)
	Get_LeftInset() float32
	Put_LeftInset(value float32)
	Get_RightInset() float32
	Put_RightInset(value float32)
	Get_TopInset() float32
	Put_TopInset(value float32)
}

type IInsetClipVtbl struct {
	win32.IInspectableVtbl
	Get_BottomInset uintptr
	Put_BottomInset uintptr
	Get_LeftInset   uintptr
	Put_LeftInset   uintptr
	Get_RightInset  uintptr
	Put_RightInset  uintptr
	Get_TopInset    uintptr
	Put_TopInset    uintptr
}

type IInsetClip struct {
	win32.IInspectable
}

func (this *IInsetClip) Vtbl() *IInsetClipVtbl {
	return (*IInsetClipVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInsetClip) Get_BottomInset() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BottomInset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInsetClip) Put_BottomInset(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BottomInset, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IInsetClip) Get_LeftInset() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LeftInset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInsetClip) Put_LeftInset(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LeftInset, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IInsetClip) Get_RightInset() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RightInset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInsetClip) Put_RightInset(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RightInset, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IInsetClip) Get_TopInset() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TopInset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInsetClip) Put_TopInset(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TopInset, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 126E7F22-3AE9-4540-9A8A-DEAE8A4A4A84
var IID_IKeyFrameAnimation = syscall.GUID{0x126E7F22, 0x3AE9, 0x4540,
	[8]byte{0x9A, 0x8A, 0xDE, 0xAE, 0x8A, 0x4A, 0x4A, 0x84}}

type IKeyFrameAnimationInterface interface {
	win32.IInspectableInterface
	Get_DelayTime() TimeSpan
	Put_DelayTime(value TimeSpan)
	Get_Duration() TimeSpan
	Put_Duration(value TimeSpan)
	Get_IterationBehavior() AnimationIterationBehavior
	Put_IterationBehavior(value AnimationIterationBehavior)
	Get_IterationCount() int32
	Put_IterationCount(value int32)
	Get_KeyFrameCount() int32
	Get_StopBehavior() AnimationStopBehavior
	Put_StopBehavior(value AnimationStopBehavior)
	InsertExpressionKeyFrame(normalizedProgressKey float32, value string)
	InsertExpressionKeyFrameWithEasingFunction(normalizedProgressKey float32, value string, easingFunction *ICompositionEasingFunction)
}

type IKeyFrameAnimationVtbl struct {
	win32.IInspectableVtbl
	Get_DelayTime                              uintptr
	Put_DelayTime                              uintptr
	Get_Duration                               uintptr
	Put_Duration                               uintptr
	Get_IterationBehavior                      uintptr
	Put_IterationBehavior                      uintptr
	Get_IterationCount                         uintptr
	Put_IterationCount                         uintptr
	Get_KeyFrameCount                          uintptr
	Get_StopBehavior                           uintptr
	Put_StopBehavior                           uintptr
	InsertExpressionKeyFrame                   uintptr
	InsertExpressionKeyFrameWithEasingFunction uintptr
}

type IKeyFrameAnimation struct {
	win32.IInspectable
}

func (this *IKeyFrameAnimation) Vtbl() *IKeyFrameAnimationVtbl {
	return (*IKeyFrameAnimationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKeyFrameAnimation) Get_DelayTime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DelayTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IKeyFrameAnimation) Put_DelayTime(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DelayTime, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IKeyFrameAnimation) Get_Duration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Duration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IKeyFrameAnimation) Put_Duration(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Duration, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IKeyFrameAnimation) Get_IterationBehavior() AnimationIterationBehavior {
	var _result AnimationIterationBehavior
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IterationBehavior, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IKeyFrameAnimation) Put_IterationBehavior(value AnimationIterationBehavior) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IterationBehavior, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IKeyFrameAnimation) Get_IterationCount() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IterationCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IKeyFrameAnimation) Put_IterationCount(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IterationCount, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IKeyFrameAnimation) Get_KeyFrameCount() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KeyFrameCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IKeyFrameAnimation) Get_StopBehavior() AnimationStopBehavior {
	var _result AnimationStopBehavior
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StopBehavior, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IKeyFrameAnimation) Put_StopBehavior(value AnimationStopBehavior) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StopBehavior, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IKeyFrameAnimation) InsertExpressionKeyFrame(normalizedProgressKey float32, value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertExpressionKeyFrame, uintptr(unsafe.Pointer(this)), uintptr(normalizedProgressKey), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IKeyFrameAnimation) InsertExpressionKeyFrameWithEasingFunction(normalizedProgressKey float32, value string, easingFunction *ICompositionEasingFunction) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertExpressionKeyFrameWithEasingFunction, uintptr(unsafe.Pointer(this)), uintptr(normalizedProgressKey), NewHStr(value).Ptr, uintptr(unsafe.Pointer(easingFunction)))
	_ = _hr
}

// F4B488BB-2940-4EC0-A41A-EB6D801A2F18
var IID_IKeyFrameAnimation2 = syscall.GUID{0xF4B488BB, 0x2940, 0x4EC0,
	[8]byte{0xA4, 0x1A, 0xEB, 0x6D, 0x80, 0x1A, 0x2F, 0x18}}

type IKeyFrameAnimation2Interface interface {
	win32.IInspectableInterface
	Get_Direction() AnimationDirection
	Put_Direction(value AnimationDirection)
}

type IKeyFrameAnimation2Vtbl struct {
	win32.IInspectableVtbl
	Get_Direction uintptr
	Put_Direction uintptr
}

type IKeyFrameAnimation2 struct {
	win32.IInspectable
}

func (this *IKeyFrameAnimation2) Vtbl() *IKeyFrameAnimation2Vtbl {
	return (*IKeyFrameAnimation2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKeyFrameAnimation2) Get_Direction() AnimationDirection {
	var _result AnimationDirection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Direction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IKeyFrameAnimation2) Put_Direction(value AnimationDirection) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Direction, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 845BF0B4-D8DE-462F-8753-C80D43C6FF5A
var IID_IKeyFrameAnimation3 = syscall.GUID{0x845BF0B4, 0xD8DE, 0x462F,
	[8]byte{0x87, 0x53, 0xC8, 0x0D, 0x43, 0xC6, 0xFF, 0x5A}}

type IKeyFrameAnimation3Interface interface {
	win32.IInspectableInterface
	Get_DelayBehavior() AnimationDelayBehavior
	Put_DelayBehavior(value AnimationDelayBehavior)
}

type IKeyFrameAnimation3Vtbl struct {
	win32.IInspectableVtbl
	Get_DelayBehavior uintptr
	Put_DelayBehavior uintptr
}

type IKeyFrameAnimation3 struct {
	win32.IInspectable
}

func (this *IKeyFrameAnimation3) Vtbl() *IKeyFrameAnimation3Vtbl {
	return (*IKeyFrameAnimation3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKeyFrameAnimation3) Get_DelayBehavior() AnimationDelayBehavior {
	var _result AnimationDelayBehavior
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DelayBehavior, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IKeyFrameAnimation3) Put_DelayBehavior(value AnimationDelayBehavior) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DelayBehavior, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// BF0803F8-712A-4FC1-8C87-970859ED8D2E
var IID_IKeyFrameAnimationFactory = syscall.GUID{0xBF0803F8, 0x712A, 0x4FC1,
	[8]byte{0x8C, 0x87, 0x97, 0x08, 0x59, 0xED, 0x8D, 0x2E}}

type IKeyFrameAnimationFactoryInterface interface {
	win32.IInspectableInterface
}

type IKeyFrameAnimationFactoryVtbl struct {
	win32.IInspectableVtbl
}

type IKeyFrameAnimationFactory struct {
	win32.IInspectable
}

func (this *IKeyFrameAnimationFactory) Vtbl() *IKeyFrameAnimationFactoryVtbl {
	return (*IKeyFrameAnimationFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// AF843985-0444-4887-8E83-B40B253F822C
var IID_ILayerVisual = syscall.GUID{0xAF843985, 0x0444, 0x4887,
	[8]byte{0x8E, 0x83, 0xB4, 0x0B, 0x25, 0x3F, 0x82, 0x2C}}

type ILayerVisualInterface interface {
	win32.IInspectableInterface
	Get_Effect() *ICompositionEffectBrush
	Put_Effect(value *ICompositionEffectBrush)
}

type ILayerVisualVtbl struct {
	win32.IInspectableVtbl
	Get_Effect uintptr
	Put_Effect uintptr
}

type ILayerVisual struct {
	win32.IInspectable
}

func (this *ILayerVisual) Vtbl() *ILayerVisualVtbl {
	return (*ILayerVisualVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILayerVisual) Get_Effect() *ICompositionEffectBrush {
	var _result *ICompositionEffectBrush
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Effect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILayerVisual) Put_Effect(value *ICompositionEffectBrush) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Effect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 98F9AEEB-6F23-49F1-90B1-1F59A14FBCE3
var IID_ILayerVisual2 = syscall.GUID{0x98F9AEEB, 0x6F23, 0x49F1,
	[8]byte{0x90, 0xB1, 0x1F, 0x59, 0xA1, 0x4F, 0xBC, 0xE3}}

type ILayerVisual2Interface interface {
	win32.IInspectableInterface
	Get_Shadow() *ICompositionShadow
	Put_Shadow(value *ICompositionShadow)
}

type ILayerVisual2Vtbl struct {
	win32.IInspectableVtbl
	Get_Shadow uintptr
	Put_Shadow uintptr
}

type ILayerVisual2 struct {
	win32.IInspectable
}

func (this *ILayerVisual2) Vtbl() *ILayerVisual2Vtbl {
	return (*ILayerVisual2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILayerVisual2) Get_Shadow() *ICompositionShadow {
	var _result *ICompositionShadow
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Shadow, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILayerVisual2) Put_Shadow(value *ICompositionShadow) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Shadow, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 9400975A-C7A6-46B3-ACF7-1A268A0A117D
var IID_ILinearEasingFunction = syscall.GUID{0x9400975A, 0xC7A6, 0x46B3,
	[8]byte{0xAC, 0xF7, 0x1A, 0x26, 0x8A, 0x0A, 0x11, 0x7D}}

type ILinearEasingFunctionInterface interface {
	win32.IInspectableInterface
}

type ILinearEasingFunctionVtbl struct {
	win32.IInspectableVtbl
}

type ILinearEasingFunction struct {
	win32.IInspectable
}

func (this *ILinearEasingFunction) Vtbl() *ILinearEasingFunctionVtbl {
	return (*ILinearEasingFunctionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 438DE12D-769B-4821-A949-284A6547E873
var IID_INaturalMotionAnimation = syscall.GUID{0x438DE12D, 0x769B, 0x4821,
	[8]byte{0xA9, 0x49, 0x28, 0x4A, 0x65, 0x47, 0xE8, 0x73}}

type INaturalMotionAnimationInterface interface {
	win32.IInspectableInterface
	Get_DelayBehavior() AnimationDelayBehavior
	Put_DelayBehavior(value AnimationDelayBehavior)
	Get_DelayTime() TimeSpan
	Put_DelayTime(value TimeSpan)
	Get_StopBehavior() AnimationStopBehavior
	Put_StopBehavior(value AnimationStopBehavior)
}

type INaturalMotionAnimationVtbl struct {
	win32.IInspectableVtbl
	Get_DelayBehavior uintptr
	Put_DelayBehavior uintptr
	Get_DelayTime     uintptr
	Put_DelayTime     uintptr
	Get_StopBehavior  uintptr
	Put_StopBehavior  uintptr
}

type INaturalMotionAnimation struct {
	win32.IInspectable
}

func (this *INaturalMotionAnimation) Vtbl() *INaturalMotionAnimationVtbl {
	return (*INaturalMotionAnimationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INaturalMotionAnimation) Get_DelayBehavior() AnimationDelayBehavior {
	var _result AnimationDelayBehavior
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DelayBehavior, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INaturalMotionAnimation) Put_DelayBehavior(value AnimationDelayBehavior) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DelayBehavior, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *INaturalMotionAnimation) Get_DelayTime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DelayTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INaturalMotionAnimation) Put_DelayTime(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DelayTime, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *INaturalMotionAnimation) Get_StopBehavior() AnimationStopBehavior {
	var _result AnimationStopBehavior
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StopBehavior, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *INaturalMotionAnimation) Put_StopBehavior(value AnimationStopBehavior) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StopBehavior, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// F53ACB06-CF6A-4387-A3FE-5221F3E7E0E0
var IID_INaturalMotionAnimationFactory = syscall.GUID{0xF53ACB06, 0xCF6A, 0x4387,
	[8]byte{0xA3, 0xFE, 0x52, 0x21, 0xF3, 0xE7, 0xE0, 0xE0}}

type INaturalMotionAnimationFactoryInterface interface {
	win32.IInspectableInterface
}

type INaturalMotionAnimationFactoryVtbl struct {
	win32.IInspectableVtbl
}

type INaturalMotionAnimationFactory struct {
	win32.IInspectable
}

func (this *INaturalMotionAnimationFactory) Vtbl() *INaturalMotionAnimationFactoryVtbl {
	return (*INaturalMotionAnimationFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 9D0D18C9-1576-4B3F-BE60-1D5031F5E71B
var IID_IPathKeyFrameAnimation = syscall.GUID{0x9D0D18C9, 0x1576, 0x4B3F,
	[8]byte{0xBE, 0x60, 0x1D, 0x50, 0x31, 0xF5, 0xE7, 0x1B}}

type IPathKeyFrameAnimationInterface interface {
	win32.IInspectableInterface
	InsertKeyFrame(normalizedProgressKey float32, path *ICompositionPath)
	InsertKeyFrameWithEasingFunction(normalizedProgressKey float32, path *ICompositionPath, easingFunction *ICompositionEasingFunction)
}

type IPathKeyFrameAnimationVtbl struct {
	win32.IInspectableVtbl
	InsertKeyFrame                   uintptr
	InsertKeyFrameWithEasingFunction uintptr
}

type IPathKeyFrameAnimation struct {
	win32.IInspectable
}

func (this *IPathKeyFrameAnimation) Vtbl() *IPathKeyFrameAnimationVtbl {
	return (*IPathKeyFrameAnimationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPathKeyFrameAnimation) InsertKeyFrame(normalizedProgressKey float32, path *ICompositionPath) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertKeyFrame, uintptr(unsafe.Pointer(this)), uintptr(normalizedProgressKey), uintptr(unsafe.Pointer(path)))
	_ = _hr
}

func (this *IPathKeyFrameAnimation) InsertKeyFrameWithEasingFunction(normalizedProgressKey float32, path *ICompositionPath, easingFunction *ICompositionEasingFunction) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertKeyFrameWithEasingFunction, uintptr(unsafe.Pointer(this)), uintptr(normalizedProgressKey), uintptr(unsafe.Pointer(path)), uintptr(unsafe.Pointer(easingFunction)))
	_ = _hr
}

// B18545B3-0C5A-4AB0-BEDC-4F3546948272
var IID_IPointLight = syscall.GUID{0xB18545B3, 0x0C5A, 0x4AB0,
	[8]byte{0xBE, 0xDC, 0x4F, 0x35, 0x46, 0x94, 0x82, 0x72}}

type IPointLightInterface interface {
	win32.IInspectableInterface
	Get_Color() unsafe.Pointer
	Put_Color(value unsafe.Pointer)
	Get_ConstantAttenuation() float32
	Put_ConstantAttenuation(value float32)
	Get_CoordinateSpace() *IVisual
	Put_CoordinateSpace(value *IVisual)
	Get_LinearAttenuation() float32
	Put_LinearAttenuation(value float32)
	Get_Offset() unsafe.Pointer
	Put_Offset(value unsafe.Pointer)
	Get_QuadraticAttenuation() float32
	Put_QuadraticAttenuation(value float32)
}

type IPointLightVtbl struct {
	win32.IInspectableVtbl
	Get_Color                uintptr
	Put_Color                uintptr
	Get_ConstantAttenuation  uintptr
	Put_ConstantAttenuation  uintptr
	Get_CoordinateSpace      uintptr
	Put_CoordinateSpace      uintptr
	Get_LinearAttenuation    uintptr
	Put_LinearAttenuation    uintptr
	Get_Offset               uintptr
	Put_Offset               uintptr
	Get_QuadraticAttenuation uintptr
	Put_QuadraticAttenuation uintptr
}

type IPointLight struct {
	win32.IInspectable
}

func (this *IPointLight) Vtbl() *IPointLightVtbl {
	return (*IPointLightVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPointLight) Get_Color() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Color, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointLight) Put_Color(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Color, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPointLight) Get_ConstantAttenuation() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConstantAttenuation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointLight) Put_ConstantAttenuation(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ConstantAttenuation, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPointLight) Get_CoordinateSpace() *IVisual {
	var _result *IVisual
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CoordinateSpace, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPointLight) Put_CoordinateSpace(value *IVisual) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CoordinateSpace, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IPointLight) Get_LinearAttenuation() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LinearAttenuation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointLight) Put_LinearAttenuation(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LinearAttenuation, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPointLight) Get_Offset() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Offset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointLight) Put_Offset(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Offset, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPointLight) Get_QuadraticAttenuation() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_QuadraticAttenuation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointLight) Put_QuadraticAttenuation(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_QuadraticAttenuation, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// EFE98F2C-0678-4F69-B164-A810D995BCB7
var IID_IPointLight2 = syscall.GUID{0xEFE98F2C, 0x0678, 0x4F69,
	[8]byte{0xB1, 0x64, 0xA8, 0x10, 0xD9, 0x95, 0xBC, 0xB7}}

type IPointLight2Interface interface {
	win32.IInspectableInterface
	Get_Intensity() float32
	Put_Intensity(value float32)
}

type IPointLight2Vtbl struct {
	win32.IInspectableVtbl
	Get_Intensity uintptr
	Put_Intensity uintptr
}

type IPointLight2 struct {
	win32.IInspectable
}

func (this *IPointLight2) Vtbl() *IPointLight2Vtbl {
	return (*IPointLight2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPointLight2) Get_Intensity() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Intensity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointLight2) Put_Intensity(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Intensity, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 4C0A8367-D4E9-468A-87AE-7BA43AB29485
var IID_IPointLight3 = syscall.GUID{0x4C0A8367, 0xD4E9, 0x468A,
	[8]byte{0x87, 0xAE, 0x7B, 0xA4, 0x3A, 0xB2, 0x94, 0x85}}

type IPointLight3Interface interface {
	win32.IInspectableInterface
	Get_MinAttenuationCutoff() float32
	Put_MinAttenuationCutoff(value float32)
	Get_MaxAttenuationCutoff() float32
	Put_MaxAttenuationCutoff(value float32)
}

type IPointLight3Vtbl struct {
	win32.IInspectableVtbl
	Get_MinAttenuationCutoff uintptr
	Put_MinAttenuationCutoff uintptr
	Get_MaxAttenuationCutoff uintptr
	Put_MaxAttenuationCutoff uintptr
}

type IPointLight3 struct {
	win32.IInspectable
}

func (this *IPointLight3) Vtbl() *IPointLight3Vtbl {
	return (*IPointLight3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPointLight3) Get_MinAttenuationCutoff() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinAttenuationCutoff, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointLight3) Put_MinAttenuationCutoff(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MinAttenuationCutoff, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IPointLight3) Get_MaxAttenuationCutoff() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxAttenuationCutoff, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPointLight3) Put_MaxAttenuationCutoff(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxAttenuationCutoff, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// C3FF53D6-138B-5815-891A-B7F615CCC563
var IID_IPowerEasingFunction = syscall.GUID{0xC3FF53D6, 0x138B, 0x5815,
	[8]byte{0x89, 0x1A, 0xB7, 0xF6, 0x15, 0xCC, 0xC5, 0x63}}

type IPowerEasingFunctionInterface interface {
	win32.IInspectableInterface
	Get_Mode() CompositionEasingFunctionMode
	Get_Power() float32
}

type IPowerEasingFunctionVtbl struct {
	win32.IInspectableVtbl
	Get_Mode  uintptr
	Get_Power uintptr
}

type IPowerEasingFunction struct {
	win32.IInspectable
}

func (this *IPowerEasingFunction) Vtbl() *IPowerEasingFunctionVtbl {
	return (*IPowerEasingFunctionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPowerEasingFunction) Get_Mode() CompositionEasingFunctionMode {
	var _result CompositionEasingFunctionMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPowerEasingFunction) Get_Power() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Power, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 404E5835-ECF6-4240-8520-671279CF36BC
var IID_IQuaternionKeyFrameAnimation = syscall.GUID{0x404E5835, 0xECF6, 0x4240,
	[8]byte{0x85, 0x20, 0x67, 0x12, 0x79, 0xCF, 0x36, 0xBC}}

type IQuaternionKeyFrameAnimationInterface interface {
	win32.IInspectableInterface
	InsertKeyFrame(normalizedProgressKey float32, value unsafe.Pointer)
	InsertKeyFrameWithEasingFunction(normalizedProgressKey float32, value unsafe.Pointer, easingFunction *ICompositionEasingFunction)
}

type IQuaternionKeyFrameAnimationVtbl struct {
	win32.IInspectableVtbl
	InsertKeyFrame                   uintptr
	InsertKeyFrameWithEasingFunction uintptr
}

type IQuaternionKeyFrameAnimation struct {
	win32.IInspectable
}

func (this *IQuaternionKeyFrameAnimation) Vtbl() *IQuaternionKeyFrameAnimationVtbl {
	return (*IQuaternionKeyFrameAnimationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IQuaternionKeyFrameAnimation) InsertKeyFrame(normalizedProgressKey float32, value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertKeyFrame, uintptr(unsafe.Pointer(this)), uintptr(normalizedProgressKey), uintptr(value))
	_ = _hr
}

func (this *IQuaternionKeyFrameAnimation) InsertKeyFrameWithEasingFunction(normalizedProgressKey float32, value unsafe.Pointer, easingFunction *ICompositionEasingFunction) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertKeyFrameWithEasingFunction, uintptr(unsafe.Pointer(this)), uintptr(normalizedProgressKey), uintptr(value), uintptr(unsafe.Pointer(easingFunction)))
	_ = _hr
}

// B3E7549E-00B4-5B53-8BE8-353F6C433101
var IID_IRectangleClip = syscall.GUID{0xB3E7549E, 0x00B4, 0x5B53,
	[8]byte{0x8B, 0xE8, 0x35, 0x3F, 0x6C, 0x43, 0x31, 0x01}}

type IRectangleClipInterface interface {
	win32.IInspectableInterface
	Get_Bottom() float32
	Put_Bottom(value float32)
	Get_BottomLeftRadius() unsafe.Pointer
	Put_BottomLeftRadius(value unsafe.Pointer)
	Get_BottomRightRadius() unsafe.Pointer
	Put_BottomRightRadius(value unsafe.Pointer)
	Get_Left() float32
	Put_Left(value float32)
	Get_Right() float32
	Put_Right(value float32)
	Get_Top() float32
	Put_Top(value float32)
	Get_TopLeftRadius() unsafe.Pointer
	Put_TopLeftRadius(value unsafe.Pointer)
	Get_TopRightRadius() unsafe.Pointer
	Put_TopRightRadius(value unsafe.Pointer)
}

type IRectangleClipVtbl struct {
	win32.IInspectableVtbl
	Get_Bottom            uintptr
	Put_Bottom            uintptr
	Get_BottomLeftRadius  uintptr
	Put_BottomLeftRadius  uintptr
	Get_BottomRightRadius uintptr
	Put_BottomRightRadius uintptr
	Get_Left              uintptr
	Put_Left              uintptr
	Get_Right             uintptr
	Put_Right             uintptr
	Get_Top               uintptr
	Put_Top               uintptr
	Get_TopLeftRadius     uintptr
	Put_TopLeftRadius     uintptr
	Get_TopRightRadius    uintptr
	Put_TopRightRadius    uintptr
}

type IRectangleClip struct {
	win32.IInspectable
}

func (this *IRectangleClip) Vtbl() *IRectangleClipVtbl {
	return (*IRectangleClipVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRectangleClip) Get_Bottom() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bottom, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRectangleClip) Put_Bottom(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Bottom, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IRectangleClip) Get_BottomLeftRadius() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BottomLeftRadius, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRectangleClip) Put_BottomLeftRadius(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BottomLeftRadius, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IRectangleClip) Get_BottomRightRadius() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BottomRightRadius, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRectangleClip) Put_BottomRightRadius(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BottomRightRadius, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IRectangleClip) Get_Left() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Left, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRectangleClip) Put_Left(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Left, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IRectangleClip) Get_Right() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Right, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRectangleClip) Put_Right(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Right, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IRectangleClip) Get_Top() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Top, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRectangleClip) Put_Top(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Top, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IRectangleClip) Get_TopLeftRadius() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TopLeftRadius, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRectangleClip) Put_TopLeftRadius(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TopLeftRadius, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IRectangleClip) Get_TopRightRadius() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TopRightRadius, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRectangleClip) Put_TopRightRadius(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TopRightRadius, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 8CC6E340-8B75-5422-B06F-09FFE9F8617E
var IID_IRedirectVisual = syscall.GUID{0x8CC6E340, 0x8B75, 0x5422,
	[8]byte{0xB0, 0x6F, 0x09, 0xFF, 0xE9, 0xF8, 0x61, 0x7E}}

type IRedirectVisualInterface interface {
	win32.IInspectableInterface
	Get_Source() *IVisual
	Put_Source(value *IVisual)
}

type IRedirectVisualVtbl struct {
	win32.IInspectableVtbl
	Get_Source uintptr
	Put_Source uintptr
}

type IRedirectVisual struct {
	win32.IInspectable
}

func (this *IRedirectVisual) Vtbl() *IRedirectVisualVtbl {
	return (*IRedirectVisualVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRedirectVisual) Get_Source() *IVisual {
	var _result *IVisual
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Source, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRedirectVisual) Put_Source(value *IVisual) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Source, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 3A31AC7D-28BF-4E7A-8524-71679D480F38
var IID_IRenderingDeviceReplacedEventArgs = syscall.GUID{0x3A31AC7D, 0x28BF, 0x4E7A,
	[8]byte{0x85, 0x24, 0x71, 0x67, 0x9D, 0x48, 0x0F, 0x38}}

type IRenderingDeviceReplacedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_GraphicsDevice() *ICompositionGraphicsDevice
}

type IRenderingDeviceReplacedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_GraphicsDevice uintptr
}

type IRenderingDeviceReplacedEventArgs struct {
	win32.IInspectable
}

func (this *IRenderingDeviceReplacedEventArgs) Vtbl() *IRenderingDeviceReplacedEventArgsVtbl {
	return (*IRenderingDeviceReplacedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRenderingDeviceReplacedEventArgs) Get_GraphicsDevice() *ICompositionGraphicsDevice {
	var _result *ICompositionGraphicsDevice
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GraphicsDevice, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AE288FA9-252C-4B95-A725-BF85E38000A1
var IID_IScalarKeyFrameAnimation = syscall.GUID{0xAE288FA9, 0x252C, 0x4B95,
	[8]byte{0xA7, 0x25, 0xBF, 0x85, 0xE3, 0x80, 0x00, 0xA1}}

type IScalarKeyFrameAnimationInterface interface {
	win32.IInspectableInterface
	InsertKeyFrame(normalizedProgressKey float32, value float32)
	InsertKeyFrameWithEasingFunction(normalizedProgressKey float32, value float32, easingFunction *ICompositionEasingFunction)
}

type IScalarKeyFrameAnimationVtbl struct {
	win32.IInspectableVtbl
	InsertKeyFrame                   uintptr
	InsertKeyFrameWithEasingFunction uintptr
}

type IScalarKeyFrameAnimation struct {
	win32.IInspectable
}

func (this *IScalarKeyFrameAnimation) Vtbl() *IScalarKeyFrameAnimationVtbl {
	return (*IScalarKeyFrameAnimationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IScalarKeyFrameAnimation) InsertKeyFrame(normalizedProgressKey float32, value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertKeyFrame, uintptr(unsafe.Pointer(this)), uintptr(normalizedProgressKey), uintptr(value))
	_ = _hr
}

func (this *IScalarKeyFrameAnimation) InsertKeyFrameWithEasingFunction(normalizedProgressKey float32, value float32, easingFunction *ICompositionEasingFunction) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertKeyFrameWithEasingFunction, uintptr(unsafe.Pointer(this)), uintptr(normalizedProgressKey), uintptr(value), uintptr(unsafe.Pointer(easingFunction)))
	_ = _hr
}

// 94A94581-BF92-495B-B5BD-D2C659430737
var IID_IScalarNaturalMotionAnimation = syscall.GUID{0x94A94581, 0xBF92, 0x495B,
	[8]byte{0xB5, 0xBD, 0xD2, 0xC6, 0x59, 0x43, 0x07, 0x37}}

type IScalarNaturalMotionAnimationInterface interface {
	win32.IInspectableInterface
	Get_FinalValue() *IReference[float32]
	Put_FinalValue(value *IReference[float32])
	Get_InitialValue() *IReference[float32]
	Put_InitialValue(value *IReference[float32])
	Get_InitialVelocity() float32
	Put_InitialVelocity(value float32)
}

type IScalarNaturalMotionAnimationVtbl struct {
	win32.IInspectableVtbl
	Get_FinalValue      uintptr
	Put_FinalValue      uintptr
	Get_InitialValue    uintptr
	Put_InitialValue    uintptr
	Get_InitialVelocity uintptr
	Put_InitialVelocity uintptr
}

type IScalarNaturalMotionAnimation struct {
	win32.IInspectable
}

func (this *IScalarNaturalMotionAnimation) Vtbl() *IScalarNaturalMotionAnimationVtbl {
	return (*IScalarNaturalMotionAnimationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IScalarNaturalMotionAnimation) Get_FinalValue() *IReference[float32] {
	var _result *IReference[float32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FinalValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IScalarNaturalMotionAnimation) Put_FinalValue(value *IReference[float32]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FinalValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IScalarNaturalMotionAnimation) Get_InitialValue() *IReference[float32] {
	var _result *IReference[float32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InitialValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IScalarNaturalMotionAnimation) Put_InitialValue(value *IReference[float32]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InitialValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IScalarNaturalMotionAnimation) Get_InitialVelocity() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InitialVelocity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IScalarNaturalMotionAnimation) Put_InitialVelocity(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InitialVelocity, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 835AA4FC-671C-41DD-AF48-AE8DEF8B1529
var IID_IScalarNaturalMotionAnimationFactory = syscall.GUID{0x835AA4FC, 0x671C, 0x41DD,
	[8]byte{0xAF, 0x48, 0xAE, 0x8D, 0xEF, 0x8B, 0x15, 0x29}}

type IScalarNaturalMotionAnimationFactoryInterface interface {
	win32.IInspectableInterface
}

type IScalarNaturalMotionAnimationFactoryVtbl struct {
	win32.IInspectableVtbl
}

type IScalarNaturalMotionAnimationFactory struct {
	win32.IInspectable
}

func (this *IScalarNaturalMotionAnimationFactory) Vtbl() *IScalarNaturalMotionAnimationFactoryVtbl {
	return (*IScalarNaturalMotionAnimationFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// F2BD13C3-BA7E-4B0F-9126-FFB7536B8176
var IID_IShapeVisual = syscall.GUID{0xF2BD13C3, 0xBA7E, 0x4B0F,
	[8]byte{0x91, 0x26, 0xFF, 0xB7, 0x53, 0x6B, 0x81, 0x76}}

type IShapeVisualInterface interface {
	win32.IInspectableInterface
	Get_Shapes() *IVector[*ICompositionShape]
	Get_ViewBox() *ICompositionViewBox
	Put_ViewBox(value *ICompositionViewBox)
}

type IShapeVisualVtbl struct {
	win32.IInspectableVtbl
	Get_Shapes  uintptr
	Get_ViewBox uintptr
	Put_ViewBox uintptr
}

type IShapeVisual struct {
	win32.IInspectable
}

func (this *IShapeVisual) Vtbl() *IShapeVisualVtbl {
	return (*IShapeVisualVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IShapeVisual) Get_Shapes() *IVector[*ICompositionShape] {
	var _result *IVector[*ICompositionShape]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Shapes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IShapeVisual) Get_ViewBox() *ICompositionViewBox {
	var _result *ICompositionViewBox
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ViewBox, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IShapeVisual) Put_ViewBox(value *ICompositionViewBox) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ViewBox, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// F1B518BF-9563-5474-BD13-44B2DF4B1D58
var IID_ISineEasingFunction = syscall.GUID{0xF1B518BF, 0x9563, 0x5474,
	[8]byte{0xBD, 0x13, 0x44, 0xB2, 0xDF, 0x4B, 0x1D, 0x58}}

type ISineEasingFunctionInterface interface {
	win32.IInspectableInterface
	Get_Mode() CompositionEasingFunctionMode
}

type ISineEasingFunctionVtbl struct {
	win32.IInspectableVtbl
	Get_Mode uintptr
}

type ISineEasingFunction struct {
	win32.IInspectable
}

func (this *ISineEasingFunction) Vtbl() *ISineEasingFunctionVtbl {
	return (*ISineEasingFunctionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISineEasingFunction) Get_Mode() CompositionEasingFunctionMode {
	var _result CompositionEasingFunctionMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5A9FE273-44A1-4F95-A422-8FA5116BDB44
var IID_ISpotLight = syscall.GUID{0x5A9FE273, 0x44A1, 0x4F95,
	[8]byte{0xA4, 0x22, 0x8F, 0xA5, 0x11, 0x6B, 0xDB, 0x44}}

type ISpotLightInterface interface {
	win32.IInspectableInterface
	Get_ConstantAttenuation() float32
	Put_ConstantAttenuation(value float32)
	Get_CoordinateSpace() *IVisual
	Put_CoordinateSpace(value *IVisual)
	Get_Direction() unsafe.Pointer
	Put_Direction(value unsafe.Pointer)
	Get_InnerConeAngle() float32
	Put_InnerConeAngle(value float32)
	Get_InnerConeAngleInDegrees() float32
	Put_InnerConeAngleInDegrees(value float32)
	Get_InnerConeColor() unsafe.Pointer
	Put_InnerConeColor(value unsafe.Pointer)
	Get_LinearAttenuation() float32
	Put_LinearAttenuation(value float32)
	Get_Offset() unsafe.Pointer
	Put_Offset(value unsafe.Pointer)
	Get_OuterConeAngle() float32
	Put_OuterConeAngle(value float32)
	Get_OuterConeAngleInDegrees() float32
	Put_OuterConeAngleInDegrees(value float32)
	Get_OuterConeColor() unsafe.Pointer
	Put_OuterConeColor(value unsafe.Pointer)
	Get_QuadraticAttenuation() float32
	Put_QuadraticAttenuation(value float32)
}

type ISpotLightVtbl struct {
	win32.IInspectableVtbl
	Get_ConstantAttenuation     uintptr
	Put_ConstantAttenuation     uintptr
	Get_CoordinateSpace         uintptr
	Put_CoordinateSpace         uintptr
	Get_Direction               uintptr
	Put_Direction               uintptr
	Get_InnerConeAngle          uintptr
	Put_InnerConeAngle          uintptr
	Get_InnerConeAngleInDegrees uintptr
	Put_InnerConeAngleInDegrees uintptr
	Get_InnerConeColor          uintptr
	Put_InnerConeColor          uintptr
	Get_LinearAttenuation       uintptr
	Put_LinearAttenuation       uintptr
	Get_Offset                  uintptr
	Put_Offset                  uintptr
	Get_OuterConeAngle          uintptr
	Put_OuterConeAngle          uintptr
	Get_OuterConeAngleInDegrees uintptr
	Put_OuterConeAngleInDegrees uintptr
	Get_OuterConeColor          uintptr
	Put_OuterConeColor          uintptr
	Get_QuadraticAttenuation    uintptr
	Put_QuadraticAttenuation    uintptr
}

type ISpotLight struct {
	win32.IInspectable
}

func (this *ISpotLight) Vtbl() *ISpotLightVtbl {
	return (*ISpotLightVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpotLight) Get_ConstantAttenuation() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConstantAttenuation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpotLight) Put_ConstantAttenuation(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ConstantAttenuation, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISpotLight) Get_CoordinateSpace() *IVisual {
	var _result *IVisual
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CoordinateSpace, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpotLight) Put_CoordinateSpace(value *IVisual) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CoordinateSpace, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ISpotLight) Get_Direction() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Direction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpotLight) Put_Direction(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Direction, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISpotLight) Get_InnerConeAngle() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InnerConeAngle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpotLight) Put_InnerConeAngle(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InnerConeAngle, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISpotLight) Get_InnerConeAngleInDegrees() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InnerConeAngleInDegrees, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpotLight) Put_InnerConeAngleInDegrees(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InnerConeAngleInDegrees, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISpotLight) Get_InnerConeColor() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InnerConeColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpotLight) Put_InnerConeColor(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InnerConeColor, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISpotLight) Get_LinearAttenuation() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LinearAttenuation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpotLight) Put_LinearAttenuation(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_LinearAttenuation, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISpotLight) Get_Offset() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Offset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpotLight) Put_Offset(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Offset, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISpotLight) Get_OuterConeAngle() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OuterConeAngle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpotLight) Put_OuterConeAngle(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_OuterConeAngle, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISpotLight) Get_OuterConeAngleInDegrees() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OuterConeAngleInDegrees, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpotLight) Put_OuterConeAngleInDegrees(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_OuterConeAngleInDegrees, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISpotLight) Get_OuterConeColor() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OuterConeColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpotLight) Put_OuterConeColor(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_OuterConeColor, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISpotLight) Get_QuadraticAttenuation() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_QuadraticAttenuation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpotLight) Put_QuadraticAttenuation(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_QuadraticAttenuation, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 64EE615E-0686-4DEA-A9E8-BC3A8C701459
var IID_ISpotLight2 = syscall.GUID{0x64EE615E, 0x0686, 0x4DEA,
	[8]byte{0xA9, 0xE8, 0xBC, 0x3A, 0x8C, 0x70, 0x14, 0x59}}

type ISpotLight2Interface interface {
	win32.IInspectableInterface
	Get_InnerConeIntensity() float32
	Put_InnerConeIntensity(value float32)
	Get_OuterConeIntensity() float32
	Put_OuterConeIntensity(value float32)
}

type ISpotLight2Vtbl struct {
	win32.IInspectableVtbl
	Get_InnerConeIntensity uintptr
	Put_InnerConeIntensity uintptr
	Get_OuterConeIntensity uintptr
	Put_OuterConeIntensity uintptr
}

type ISpotLight2 struct {
	win32.IInspectable
}

func (this *ISpotLight2) Vtbl() *ISpotLight2Vtbl {
	return (*ISpotLight2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpotLight2) Get_InnerConeIntensity() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InnerConeIntensity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpotLight2) Put_InnerConeIntensity(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InnerConeIntensity, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISpotLight2) Get_OuterConeIntensity() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OuterConeIntensity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpotLight2) Put_OuterConeIntensity(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_OuterConeIntensity, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// E4D03EEA-131F-480E-859E-B82705B74360
var IID_ISpotLight3 = syscall.GUID{0xE4D03EEA, 0x131F, 0x480E,
	[8]byte{0x85, 0x9E, 0xB8, 0x27, 0x05, 0xB7, 0x43, 0x60}}

type ISpotLight3Interface interface {
	win32.IInspectableInterface
	Get_MinAttenuationCutoff() float32
	Put_MinAttenuationCutoff(value float32)
	Get_MaxAttenuationCutoff() float32
	Put_MaxAttenuationCutoff(value float32)
}

type ISpotLight3Vtbl struct {
	win32.IInspectableVtbl
	Get_MinAttenuationCutoff uintptr
	Put_MinAttenuationCutoff uintptr
	Get_MaxAttenuationCutoff uintptr
	Put_MaxAttenuationCutoff uintptr
}

type ISpotLight3 struct {
	win32.IInspectable
}

func (this *ISpotLight3) Vtbl() *ISpotLight3Vtbl {
	return (*ISpotLight3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpotLight3) Get_MinAttenuationCutoff() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinAttenuationCutoff, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpotLight3) Put_MinAttenuationCutoff(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MinAttenuationCutoff, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISpotLight3) Get_MaxAttenuationCutoff() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxAttenuationCutoff, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpotLight3) Put_MaxAttenuationCutoff(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxAttenuationCutoff, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 0572A95F-37F9-4FBE-B87B-5CD03A89501C
var IID_ISpringScalarNaturalMotionAnimation = syscall.GUID{0x0572A95F, 0x37F9, 0x4FBE,
	[8]byte{0xB8, 0x7B, 0x5C, 0xD0, 0x3A, 0x89, 0x50, 0x1C}}

type ISpringScalarNaturalMotionAnimationInterface interface {
	win32.IInspectableInterface
	Get_DampingRatio() float32
	Put_DampingRatio(value float32)
	Get_Period() TimeSpan
	Put_Period(value TimeSpan)
}

type ISpringScalarNaturalMotionAnimationVtbl struct {
	win32.IInspectableVtbl
	Get_DampingRatio uintptr
	Put_DampingRatio uintptr
	Get_Period       uintptr
	Put_Period       uintptr
}

type ISpringScalarNaturalMotionAnimation struct {
	win32.IInspectable
}

func (this *ISpringScalarNaturalMotionAnimation) Vtbl() *ISpringScalarNaturalMotionAnimationVtbl {
	return (*ISpringScalarNaturalMotionAnimationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpringScalarNaturalMotionAnimation) Get_DampingRatio() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DampingRatio, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpringScalarNaturalMotionAnimation) Put_DampingRatio(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DampingRatio, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISpringScalarNaturalMotionAnimation) Get_Period() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Period, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpringScalarNaturalMotionAnimation) Put_Period(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Period, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

// 23F494B5-EE73-4F0F-A423-402B946DF4B3
var IID_ISpringVector2NaturalMotionAnimation = syscall.GUID{0x23F494B5, 0xEE73, 0x4F0F,
	[8]byte{0xA4, 0x23, 0x40, 0x2B, 0x94, 0x6D, 0xF4, 0xB3}}

type ISpringVector2NaturalMotionAnimationInterface interface {
	win32.IInspectableInterface
	Get_DampingRatio() float32
	Put_DampingRatio(value float32)
	Get_Period() TimeSpan
	Put_Period(value TimeSpan)
}

type ISpringVector2NaturalMotionAnimationVtbl struct {
	win32.IInspectableVtbl
	Get_DampingRatio uintptr
	Put_DampingRatio uintptr
	Get_Period       uintptr
	Put_Period       uintptr
}

type ISpringVector2NaturalMotionAnimation struct {
	win32.IInspectable
}

func (this *ISpringVector2NaturalMotionAnimation) Vtbl() *ISpringVector2NaturalMotionAnimationVtbl {
	return (*ISpringVector2NaturalMotionAnimationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpringVector2NaturalMotionAnimation) Get_DampingRatio() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DampingRatio, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpringVector2NaturalMotionAnimation) Put_DampingRatio(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DampingRatio, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISpringVector2NaturalMotionAnimation) Get_Period() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Period, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpringVector2NaturalMotionAnimation) Put_Period(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Period, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

// 6C8749DF-D57B-4794-8E2D-CECB11E194E5
var IID_ISpringVector3NaturalMotionAnimation = syscall.GUID{0x6C8749DF, 0xD57B, 0x4794,
	[8]byte{0x8E, 0x2D, 0xCE, 0xCB, 0x11, 0xE1, 0x94, 0xE5}}

type ISpringVector3NaturalMotionAnimationInterface interface {
	win32.IInspectableInterface
	Get_DampingRatio() float32
	Put_DampingRatio(value float32)
	Get_Period() TimeSpan
	Put_Period(value TimeSpan)
}

type ISpringVector3NaturalMotionAnimationVtbl struct {
	win32.IInspectableVtbl
	Get_DampingRatio uintptr
	Put_DampingRatio uintptr
	Get_Period       uintptr
	Put_Period       uintptr
}

type ISpringVector3NaturalMotionAnimation struct {
	win32.IInspectable
}

func (this *ISpringVector3NaturalMotionAnimation) Vtbl() *ISpringVector3NaturalMotionAnimationVtbl {
	return (*ISpringVector3NaturalMotionAnimationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpringVector3NaturalMotionAnimation) Get_DampingRatio() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DampingRatio, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpringVector3NaturalMotionAnimation) Put_DampingRatio(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DampingRatio, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISpringVector3NaturalMotionAnimation) Get_Period() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Period, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpringVector3NaturalMotionAnimation) Put_Period(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Period, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

// 08E05581-1AD1-4F97-9757-402D76E4233B
var IID_ISpriteVisual = syscall.GUID{0x08E05581, 0x1AD1, 0x4F97,
	[8]byte{0x97, 0x57, 0x40, 0x2D, 0x76, 0xE4, 0x23, 0x3B}}

type ISpriteVisualInterface interface {
	win32.IInspectableInterface
	Get_Brush() *ICompositionBrush
	Put_Brush(value *ICompositionBrush)
}

type ISpriteVisualVtbl struct {
	win32.IInspectableVtbl
	Get_Brush uintptr
	Put_Brush uintptr
}

type ISpriteVisual struct {
	win32.IInspectable
}

func (this *ISpriteVisual) Vtbl() *ISpriteVisualVtbl {
	return (*ISpriteVisualVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpriteVisual) Get_Brush() *ICompositionBrush {
	var _result *ICompositionBrush
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Brush, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpriteVisual) Put_Brush(value *ICompositionBrush) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Brush, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 588C9664-997A-4850-91FE-53CB58F81CE9
var IID_ISpriteVisual2 = syscall.GUID{0x588C9664, 0x997A, 0x4850,
	[8]byte{0x91, 0xFE, 0x53, 0xCB, 0x58, 0xF8, 0x1C, 0xE9}}

type ISpriteVisual2Interface interface {
	win32.IInspectableInterface
	Get_Shadow() *ICompositionShadow
	Put_Shadow(value *ICompositionShadow)
}

type ISpriteVisual2Vtbl struct {
	win32.IInspectableVtbl
	Get_Shadow uintptr
	Put_Shadow uintptr
}

type ISpriteVisual2 struct {
	win32.IInspectable
}

func (this *ISpriteVisual2) Vtbl() *ISpriteVisual2Vtbl {
	return (*ISpriteVisual2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpriteVisual2) Get_Shadow() *ICompositionShadow {
	var _result *ICompositionShadow
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Shadow, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpriteVisual2) Put_Shadow(value *ICompositionShadow) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Shadow, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// D0CAA74B-560C-4A0B-A5F6-206CA8C3ECD6
var IID_IStepEasingFunction = syscall.GUID{0xD0CAA74B, 0x560C, 0x4A0B,
	[8]byte{0xA5, 0xF6, 0x20, 0x6C, 0xA8, 0xC3, 0xEC, 0xD6}}

type IStepEasingFunctionInterface interface {
	win32.IInspectableInterface
	Get_FinalStep() int32
	Put_FinalStep(value int32)
	Get_InitialStep() int32
	Put_InitialStep(value int32)
	Get_IsFinalStepSingleFrame() bool
	Put_IsFinalStepSingleFrame(value bool)
	Get_IsInitialStepSingleFrame() bool
	Put_IsInitialStepSingleFrame(value bool)
	Get_StepCount() int32
	Put_StepCount(value int32)
}

type IStepEasingFunctionVtbl struct {
	win32.IInspectableVtbl
	Get_FinalStep                uintptr
	Put_FinalStep                uintptr
	Get_InitialStep              uintptr
	Put_InitialStep              uintptr
	Get_IsFinalStepSingleFrame   uintptr
	Put_IsFinalStepSingleFrame   uintptr
	Get_IsInitialStepSingleFrame uintptr
	Put_IsInitialStepSingleFrame uintptr
	Get_StepCount                uintptr
	Put_StepCount                uintptr
}

type IStepEasingFunction struct {
	win32.IInspectable
}

func (this *IStepEasingFunction) Vtbl() *IStepEasingFunctionVtbl {
	return (*IStepEasingFunctionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStepEasingFunction) Get_FinalStep() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FinalStep, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStepEasingFunction) Put_FinalStep(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FinalStep, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IStepEasingFunction) Get_InitialStep() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InitialStep, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStepEasingFunction) Put_InitialStep(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InitialStep, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IStepEasingFunction) Get_IsFinalStepSingleFrame() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsFinalStepSingleFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStepEasingFunction) Put_IsFinalStepSingleFrame(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsFinalStepSingleFrame, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IStepEasingFunction) Get_IsInitialStepSingleFrame() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsInitialStepSingleFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStepEasingFunction) Put_IsInitialStepSingleFrame(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsInitialStepSingleFrame, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IStepEasingFunction) Get_StepCount() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StepCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStepEasingFunction) Put_StepCount(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StepCount, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// DF414515-4E29-4F11-B55E-BF2A6EB36294
var IID_IVector2KeyFrameAnimation = syscall.GUID{0xDF414515, 0x4E29, 0x4F11,
	[8]byte{0xB5, 0x5E, 0xBF, 0x2A, 0x6E, 0xB3, 0x62, 0x94}}

type IVector2KeyFrameAnimationInterface interface {
	win32.IInspectableInterface
	InsertKeyFrame(normalizedProgressKey float32, value unsafe.Pointer)
	InsertKeyFrameWithEasingFunction(normalizedProgressKey float32, value unsafe.Pointer, easingFunction *ICompositionEasingFunction)
}

type IVector2KeyFrameAnimationVtbl struct {
	win32.IInspectableVtbl
	InsertKeyFrame                   uintptr
	InsertKeyFrameWithEasingFunction uintptr
}

type IVector2KeyFrameAnimation struct {
	win32.IInspectable
}

func (this *IVector2KeyFrameAnimation) Vtbl() *IVector2KeyFrameAnimationVtbl {
	return (*IVector2KeyFrameAnimationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVector2KeyFrameAnimation) InsertKeyFrame(normalizedProgressKey float32, value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertKeyFrame, uintptr(unsafe.Pointer(this)), uintptr(normalizedProgressKey), uintptr(value))
	_ = _hr
}

func (this *IVector2KeyFrameAnimation) InsertKeyFrameWithEasingFunction(normalizedProgressKey float32, value unsafe.Pointer, easingFunction *ICompositionEasingFunction) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertKeyFrameWithEasingFunction, uintptr(unsafe.Pointer(this)), uintptr(normalizedProgressKey), uintptr(value), uintptr(unsafe.Pointer(easingFunction)))
	_ = _hr
}

// 0F3E0B7D-E512-479D-A00C-77C93A30A395
var IID_IVector2NaturalMotionAnimation = syscall.GUID{0x0F3E0B7D, 0xE512, 0x479D,
	[8]byte{0xA0, 0x0C, 0x77, 0xC9, 0x3A, 0x30, 0xA3, 0x95}}

type IVector2NaturalMotionAnimationInterface interface {
	win32.IInspectableInterface
	Get_FinalValue() *IReference[unsafe.Pointer]
	Put_FinalValue(value *IReference[unsafe.Pointer])
	Get_InitialValue() *IReference[unsafe.Pointer]
	Put_InitialValue(value *IReference[unsafe.Pointer])
	Get_InitialVelocity() unsafe.Pointer
	Put_InitialVelocity(value unsafe.Pointer)
}

type IVector2NaturalMotionAnimationVtbl struct {
	win32.IInspectableVtbl
	Get_FinalValue      uintptr
	Put_FinalValue      uintptr
	Get_InitialValue    uintptr
	Put_InitialValue    uintptr
	Get_InitialVelocity uintptr
	Put_InitialVelocity uintptr
}

type IVector2NaturalMotionAnimation struct {
	win32.IInspectable
}

func (this *IVector2NaturalMotionAnimation) Vtbl() *IVector2NaturalMotionAnimationVtbl {
	return (*IVector2NaturalMotionAnimationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVector2NaturalMotionAnimation) Get_FinalValue() *IReference[unsafe.Pointer] {
	var _result *IReference[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FinalValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVector2NaturalMotionAnimation) Put_FinalValue(value *IReference[unsafe.Pointer]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FinalValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IVector2NaturalMotionAnimation) Get_InitialValue() *IReference[unsafe.Pointer] {
	var _result *IReference[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InitialValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVector2NaturalMotionAnimation) Put_InitialValue(value *IReference[unsafe.Pointer]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InitialValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IVector2NaturalMotionAnimation) Get_InitialVelocity() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InitialVelocity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVector2NaturalMotionAnimation) Put_InitialVelocity(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InitialVelocity, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 8C74FF61-0761-48A2-BDDB-6AFCC52B89D8
var IID_IVector2NaturalMotionAnimationFactory = syscall.GUID{0x8C74FF61, 0x0761, 0x48A2,
	[8]byte{0xBD, 0xDB, 0x6A, 0xFC, 0xC5, 0x2B, 0x89, 0xD8}}

type IVector2NaturalMotionAnimationFactoryInterface interface {
	win32.IInspectableInterface
}

type IVector2NaturalMotionAnimationFactoryVtbl struct {
	win32.IInspectableVtbl
}

type IVector2NaturalMotionAnimationFactory struct {
	win32.IInspectable
}

func (this *IVector2NaturalMotionAnimationFactory) Vtbl() *IVector2NaturalMotionAnimationFactoryVtbl {
	return (*IVector2NaturalMotionAnimationFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// C8039DAA-A281-43C2-A73D-B68E3C533C40
var IID_IVector3KeyFrameAnimation = syscall.GUID{0xC8039DAA, 0xA281, 0x43C2,
	[8]byte{0xA7, 0x3D, 0xB6, 0x8E, 0x3C, 0x53, 0x3C, 0x40}}

type IVector3KeyFrameAnimationInterface interface {
	win32.IInspectableInterface
	InsertKeyFrame(normalizedProgressKey float32, value unsafe.Pointer)
	InsertKeyFrameWithEasingFunction(normalizedProgressKey float32, value unsafe.Pointer, easingFunction *ICompositionEasingFunction)
}

type IVector3KeyFrameAnimationVtbl struct {
	win32.IInspectableVtbl
	InsertKeyFrame                   uintptr
	InsertKeyFrameWithEasingFunction uintptr
}

type IVector3KeyFrameAnimation struct {
	win32.IInspectable
}

func (this *IVector3KeyFrameAnimation) Vtbl() *IVector3KeyFrameAnimationVtbl {
	return (*IVector3KeyFrameAnimationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVector3KeyFrameAnimation) InsertKeyFrame(normalizedProgressKey float32, value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertKeyFrame, uintptr(unsafe.Pointer(this)), uintptr(normalizedProgressKey), uintptr(value))
	_ = _hr
}

func (this *IVector3KeyFrameAnimation) InsertKeyFrameWithEasingFunction(normalizedProgressKey float32, value unsafe.Pointer, easingFunction *ICompositionEasingFunction) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertKeyFrameWithEasingFunction, uintptr(unsafe.Pointer(this)), uintptr(normalizedProgressKey), uintptr(value), uintptr(unsafe.Pointer(easingFunction)))
	_ = _hr
}

// 9C17042C-E2CA-45AD-969E-4E78B7B9AD41
var IID_IVector3NaturalMotionAnimation = syscall.GUID{0x9C17042C, 0xE2CA, 0x45AD,
	[8]byte{0x96, 0x9E, 0x4E, 0x78, 0xB7, 0xB9, 0xAD, 0x41}}

type IVector3NaturalMotionAnimationInterface interface {
	win32.IInspectableInterface
	Get_FinalValue() *IReference[unsafe.Pointer]
	Put_FinalValue(value *IReference[unsafe.Pointer])
	Get_InitialValue() *IReference[unsafe.Pointer]
	Put_InitialValue(value *IReference[unsafe.Pointer])
	Get_InitialVelocity() unsafe.Pointer
	Put_InitialVelocity(value unsafe.Pointer)
}

type IVector3NaturalMotionAnimationVtbl struct {
	win32.IInspectableVtbl
	Get_FinalValue      uintptr
	Put_FinalValue      uintptr
	Get_InitialValue    uintptr
	Put_InitialValue    uintptr
	Get_InitialVelocity uintptr
	Put_InitialVelocity uintptr
}

type IVector3NaturalMotionAnimation struct {
	win32.IInspectable
}

func (this *IVector3NaturalMotionAnimation) Vtbl() *IVector3NaturalMotionAnimationVtbl {
	return (*IVector3NaturalMotionAnimationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVector3NaturalMotionAnimation) Get_FinalValue() *IReference[unsafe.Pointer] {
	var _result *IReference[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FinalValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVector3NaturalMotionAnimation) Put_FinalValue(value *IReference[unsafe.Pointer]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FinalValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IVector3NaturalMotionAnimation) Get_InitialValue() *IReference[unsafe.Pointer] {
	var _result *IReference[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InitialValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVector3NaturalMotionAnimation) Put_InitialValue(value *IReference[unsafe.Pointer]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InitialValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IVector3NaturalMotionAnimation) Get_InitialVelocity() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InitialVelocity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVector3NaturalMotionAnimation) Put_InitialVelocity(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InitialVelocity, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 21A81D2F-0880-457B-AC87-B609018C876D
var IID_IVector3NaturalMotionAnimationFactory = syscall.GUID{0x21A81D2F, 0x0880, 0x457B,
	[8]byte{0xAC, 0x87, 0xB6, 0x09, 0x01, 0x8C, 0x87, 0x6D}}

type IVector3NaturalMotionAnimationFactoryInterface interface {
	win32.IInspectableInterface
}

type IVector3NaturalMotionAnimationFactoryVtbl struct {
	win32.IInspectableVtbl
}

type IVector3NaturalMotionAnimationFactory struct {
	win32.IInspectable
}

func (this *IVector3NaturalMotionAnimationFactory) Vtbl() *IVector3NaturalMotionAnimationFactoryVtbl {
	return (*IVector3NaturalMotionAnimationFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 2457945B-ADDD-4385-9606-B6A3D5E4E1B9
var IID_IVector4KeyFrameAnimation = syscall.GUID{0x2457945B, 0xADDD, 0x4385,
	[8]byte{0x96, 0x06, 0xB6, 0xA3, 0xD5, 0xE4, 0xE1, 0xB9}}

type IVector4KeyFrameAnimationInterface interface {
	win32.IInspectableInterface
	InsertKeyFrame(normalizedProgressKey float32, value unsafe.Pointer)
	InsertKeyFrameWithEasingFunction(normalizedProgressKey float32, value unsafe.Pointer, easingFunction *ICompositionEasingFunction)
}

type IVector4KeyFrameAnimationVtbl struct {
	win32.IInspectableVtbl
	InsertKeyFrame                   uintptr
	InsertKeyFrameWithEasingFunction uintptr
}

type IVector4KeyFrameAnimation struct {
	win32.IInspectable
}

func (this *IVector4KeyFrameAnimation) Vtbl() *IVector4KeyFrameAnimationVtbl {
	return (*IVector4KeyFrameAnimationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVector4KeyFrameAnimation) InsertKeyFrame(normalizedProgressKey float32, value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertKeyFrame, uintptr(unsafe.Pointer(this)), uintptr(normalizedProgressKey), uintptr(value))
	_ = _hr
}

func (this *IVector4KeyFrameAnimation) InsertKeyFrameWithEasingFunction(normalizedProgressKey float32, value unsafe.Pointer, easingFunction *ICompositionEasingFunction) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertKeyFrameWithEasingFunction, uintptr(unsafe.Pointer(this)), uintptr(normalizedProgressKey), uintptr(value), uintptr(unsafe.Pointer(easingFunction)))
	_ = _hr
}

// 117E202D-A859-4C89-873B-C2AA566788E3
var IID_IVisual = syscall.GUID{0x117E202D, 0xA859, 0x4C89,
	[8]byte{0x87, 0x3B, 0xC2, 0xAA, 0x56, 0x67, 0x88, 0xE3}}

type IVisualInterface interface {
	win32.IInspectableInterface
	Get_AnchorPoint() unsafe.Pointer
	Put_AnchorPoint(value unsafe.Pointer)
	Get_BackfaceVisibility() CompositionBackfaceVisibility
	Put_BackfaceVisibility(value CompositionBackfaceVisibility)
	Get_BorderMode() CompositionBorderMode
	Put_BorderMode(value CompositionBorderMode)
	Get_CenterPoint() unsafe.Pointer
	Put_CenterPoint(value unsafe.Pointer)
	Get_Clip() *ICompositionClip
	Put_Clip(value *ICompositionClip)
	Get_CompositeMode() CompositionCompositeMode
	Put_CompositeMode(value CompositionCompositeMode)
	Get_IsVisible() bool
	Put_IsVisible(value bool)
	Get_Offset() unsafe.Pointer
	Put_Offset(value unsafe.Pointer)
	Get_Opacity() float32
	Put_Opacity(value float32)
	Get_Orientation() unsafe.Pointer
	Put_Orientation(value unsafe.Pointer)
	Get_Parent() *IContainerVisual
	Get_RotationAngle() float32
	Put_RotationAngle(value float32)
	Get_RotationAngleInDegrees() float32
	Put_RotationAngleInDegrees(value float32)
	Get_RotationAxis() unsafe.Pointer
	Put_RotationAxis(value unsafe.Pointer)
	Get_Scale() unsafe.Pointer
	Put_Scale(value unsafe.Pointer)
	Get_Size() unsafe.Pointer
	Put_Size(value unsafe.Pointer)
	Get_TransformMatrix() unsafe.Pointer
	Put_TransformMatrix(value unsafe.Pointer)
}

type IVisualVtbl struct {
	win32.IInspectableVtbl
	Get_AnchorPoint            uintptr
	Put_AnchorPoint            uintptr
	Get_BackfaceVisibility     uintptr
	Put_BackfaceVisibility     uintptr
	Get_BorderMode             uintptr
	Put_BorderMode             uintptr
	Get_CenterPoint            uintptr
	Put_CenterPoint            uintptr
	Get_Clip                   uintptr
	Put_Clip                   uintptr
	Get_CompositeMode          uintptr
	Put_CompositeMode          uintptr
	Get_IsVisible              uintptr
	Put_IsVisible              uintptr
	Get_Offset                 uintptr
	Put_Offset                 uintptr
	Get_Opacity                uintptr
	Put_Opacity                uintptr
	Get_Orientation            uintptr
	Put_Orientation            uintptr
	Get_Parent                 uintptr
	Get_RotationAngle          uintptr
	Put_RotationAngle          uintptr
	Get_RotationAngleInDegrees uintptr
	Put_RotationAngleInDegrees uintptr
	Get_RotationAxis           uintptr
	Put_RotationAxis           uintptr
	Get_Scale                  uintptr
	Put_Scale                  uintptr
	Get_Size                   uintptr
	Put_Size                   uintptr
	Get_TransformMatrix        uintptr
	Put_TransformMatrix        uintptr
}

type IVisual struct {
	win32.IInspectable
}

func (this *IVisual) Vtbl() *IVisualVtbl {
	return (*IVisualVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVisual) Get_AnchorPoint() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AnchorPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisual) Put_AnchorPoint(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AnchorPoint, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVisual) Get_BackfaceVisibility() CompositionBackfaceVisibility {
	var _result CompositionBackfaceVisibility
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BackfaceVisibility, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisual) Put_BackfaceVisibility(value CompositionBackfaceVisibility) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BackfaceVisibility, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVisual) Get_BorderMode() CompositionBorderMode {
	var _result CompositionBorderMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BorderMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisual) Put_BorderMode(value CompositionBorderMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BorderMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVisual) Get_CenterPoint() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CenterPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisual) Put_CenterPoint(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CenterPoint, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVisual) Get_Clip() *ICompositionClip {
	var _result *ICompositionClip
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Clip, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVisual) Put_Clip(value *ICompositionClip) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Clip, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IVisual) Get_CompositeMode() CompositionCompositeMode {
	var _result CompositionCompositeMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CompositeMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisual) Put_CompositeMode(value CompositionCompositeMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CompositeMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVisual) Get_IsVisible() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsVisible, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisual) Put_IsVisible(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsVisible, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IVisual) Get_Offset() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Offset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisual) Put_Offset(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Offset, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVisual) Get_Opacity() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Opacity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisual) Put_Opacity(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Opacity, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVisual) Get_Orientation() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Orientation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisual) Put_Orientation(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Orientation, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVisual) Get_Parent() *IContainerVisual {
	var _result *IContainerVisual
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Parent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVisual) Get_RotationAngle() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RotationAngle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisual) Put_RotationAngle(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RotationAngle, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVisual) Get_RotationAngleInDegrees() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RotationAngleInDegrees, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisual) Put_RotationAngleInDegrees(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RotationAngleInDegrees, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVisual) Get_RotationAxis() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RotationAxis, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisual) Put_RotationAxis(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RotationAxis, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVisual) Get_Scale() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Scale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisual) Put_Scale(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Scale, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVisual) Get_Size() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Size, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisual) Put_Size(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Size, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVisual) Get_TransformMatrix() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TransformMatrix, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisual) Put_TransformMatrix(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TransformMatrix, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 3052B611-56C3-4C3E-8BF3-F6E1AD473F06
var IID_IVisual2 = syscall.GUID{0x3052B611, 0x56C3, 0x4C3E,
	[8]byte{0x8B, 0xF3, 0xF6, 0xE1, 0xAD, 0x47, 0x3F, 0x06}}

type IVisual2Interface interface {
	win32.IInspectableInterface
	Get_ParentForTransform() *IVisual
	Put_ParentForTransform(value *IVisual)
	Get_RelativeOffsetAdjustment() unsafe.Pointer
	Put_RelativeOffsetAdjustment(value unsafe.Pointer)
	Get_RelativeSizeAdjustment() unsafe.Pointer
	Put_RelativeSizeAdjustment(value unsafe.Pointer)
}

type IVisual2Vtbl struct {
	win32.IInspectableVtbl
	Get_ParentForTransform       uintptr
	Put_ParentForTransform       uintptr
	Get_RelativeOffsetAdjustment uintptr
	Put_RelativeOffsetAdjustment uintptr
	Get_RelativeSizeAdjustment   uintptr
	Put_RelativeSizeAdjustment   uintptr
}

type IVisual2 struct {
	win32.IInspectable
}

func (this *IVisual2) Vtbl() *IVisual2Vtbl {
	return (*IVisual2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVisual2) Get_ParentForTransform() *IVisual {
	var _result *IVisual
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ParentForTransform, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVisual2) Put_ParentForTransform(value *IVisual) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ParentForTransform, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IVisual2) Get_RelativeOffsetAdjustment() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RelativeOffsetAdjustment, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisual2) Put_RelativeOffsetAdjustment(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RelativeOffsetAdjustment, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVisual2) Get_RelativeSizeAdjustment() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RelativeSizeAdjustment, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisual2) Put_RelativeSizeAdjustment(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RelativeSizeAdjustment, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 30BE580D-F4B6-4AB7-80DD-3738CBAC9F2C
var IID_IVisual3 = syscall.GUID{0x30BE580D, 0xF4B6, 0x4AB7,
	[8]byte{0x80, 0xDD, 0x37, 0x38, 0xCB, 0xAC, 0x9F, 0x2C}}

type IVisual3Interface interface {
	win32.IInspectableInterface
	Get_IsHitTestVisible() bool
	Put_IsHitTestVisible(value bool)
}

type IVisual3Vtbl struct {
	win32.IInspectableVtbl
	Get_IsHitTestVisible uintptr
	Put_IsHitTestVisible uintptr
}

type IVisual3 struct {
	win32.IInspectable
}

func (this *IVisual3) Vtbl() *IVisual3Vtbl {
	return (*IVisual3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVisual3) Get_IsHitTestVisible() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsHitTestVisible, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisual3) Put_IsHitTestVisible(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsHitTestVisible, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 9476BF11-E24B-5BF9-9EBE-6274109B2711
var IID_IVisual4 = syscall.GUID{0x9476BF11, 0xE24B, 0x5BF9,
	[8]byte{0x9E, 0xBE, 0x62, 0x74, 0x10, 0x9B, 0x27, 0x11}}

type IVisual4Interface interface {
	win32.IInspectableInterface
	Get_IsPixelSnappingEnabled() bool
	Put_IsPixelSnappingEnabled(value bool)
}

type IVisual4Vtbl struct {
	win32.IInspectableVtbl
	Get_IsPixelSnappingEnabled uintptr
	Put_IsPixelSnappingEnabled uintptr
}

type IVisual4 struct {
	win32.IInspectable
}

func (this *IVisual4) Vtbl() *IVisual4Vtbl {
	return (*IVisual4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVisual4) Get_IsPixelSnappingEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPixelSnappingEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisual4) Put_IsPixelSnappingEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsPixelSnappingEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 8B745505-FD3E-4A98-84A8-E949468C6BCB
var IID_IVisualCollection = syscall.GUID{0x8B745505, 0xFD3E, 0x4A98,
	[8]byte{0x84, 0xA8, 0xE9, 0x49, 0x46, 0x8C, 0x6B, 0xCB}}

type IVisualCollectionInterface interface {
	win32.IInspectableInterface
	Get_Count() int32
	InsertAbove(newChild *IVisual, sibling *IVisual)
	InsertAtBottom(newChild *IVisual)
	InsertAtTop(newChild *IVisual)
	InsertBelow(newChild *IVisual, sibling *IVisual)
	Remove(child *IVisual)
	RemoveAll()
}

type IVisualCollectionVtbl struct {
	win32.IInspectableVtbl
	Get_Count      uintptr
	InsertAbove    uintptr
	InsertAtBottom uintptr
	InsertAtTop    uintptr
	InsertBelow    uintptr
	Remove         uintptr
	RemoveAll      uintptr
}

type IVisualCollection struct {
	win32.IInspectable
}

func (this *IVisualCollection) Vtbl() *IVisualCollectionVtbl {
	return (*IVisualCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVisualCollection) Get_Count() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Count, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisualCollection) InsertAbove(newChild *IVisual, sibling *IVisual) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertAbove, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(newChild)), uintptr(unsafe.Pointer(sibling)))
	_ = _hr
}

func (this *IVisualCollection) InsertAtBottom(newChild *IVisual) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertAtBottom, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(newChild)))
	_ = _hr
}

func (this *IVisualCollection) InsertAtTop(newChild *IVisual) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertAtTop, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(newChild)))
	_ = _hr
}

func (this *IVisualCollection) InsertBelow(newChild *IVisual, sibling *IVisual) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertBelow, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(newChild)), uintptr(unsafe.Pointer(sibling)))
	_ = _hr
}

func (this *IVisualCollection) Remove(child *IVisual) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(child)))
	_ = _hr
}

func (this *IVisualCollection) RemoveAll() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveAll, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 01E64612-1D82-42F4-8E3F-A722DED33FC7
var IID_IVisualElement = syscall.GUID{0x01E64612, 0x1D82, 0x42F4,
	[8]byte{0x8E, 0x3F, 0xA7, 0x22, 0xDE, 0xD3, 0x3F, 0xC7}}

type IVisualElementInterface interface {
	win32.IInspectableInterface
}

type IVisualElementVtbl struct {
	win32.IInspectableVtbl
}

type IVisualElement struct {
	win32.IInspectable
}

func (this *IVisualElement) Vtbl() *IVisualElementVtbl {
	return (*IVisualElementVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 993AE8A0-6057-5E40-918C-E06E0B7E7C64
var IID_IVisualElement2 = syscall.GUID{0x993AE8A0, 0x6057, 0x5E40,
	[8]byte{0x91, 0x8C, 0xE0, 0x6E, 0x0B, 0x7E, 0x7C, 0x64}}

type IVisualElement2Interface interface {
	win32.IInspectableInterface
	GetVisualInternal() *IVisual
}

type IVisualElement2Vtbl struct {
	win32.IInspectableVtbl
	GetVisualInternal uintptr
}

type IVisualElement2 struct {
	win32.IInspectable
}

func (this *IVisualElement2) Vtbl() *IVisualElement2Vtbl {
	return (*IVisualElement2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVisualElement2) GetVisualInternal() *IVisual {
	var _result *IVisual
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetVisualInternal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AD0FF93E-B502-4EB5-87B4-9A38A71D0137
var IID_IVisualFactory = syscall.GUID{0xAD0FF93E, 0xB502, 0x4EB5,
	[8]byte{0x87, 0xB4, 0x9A, 0x38, 0xA7, 0x1D, 0x01, 0x37}}

type IVisualFactoryInterface interface {
	win32.IInspectableInterface
}

type IVisualFactoryVtbl struct {
	win32.IInspectableVtbl
}

type IVisualFactory struct {
	win32.IInspectable
}

func (this *IVisualFactory) Vtbl() *IVisualFactoryVtbl {
	return (*IVisualFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 338FAA70-54C8-40A7-8029-C9CEEB0AA250
var IID_IVisualUnorderedCollection = syscall.GUID{0x338FAA70, 0x54C8, 0x40A7,
	[8]byte{0x80, 0x29, 0xC9, 0xCE, 0xEB, 0x0A, 0xA2, 0x50}}

type IVisualUnorderedCollectionInterface interface {
	win32.IInspectableInterface
	Get_Count() int32
	Add(newVisual *IVisual)
	Remove(visual *IVisual)
	RemoveAll()
}

type IVisualUnorderedCollectionVtbl struct {
	win32.IInspectableVtbl
	Get_Count uintptr
	Add       uintptr
	Remove    uintptr
	RemoveAll uintptr
}

type IVisualUnorderedCollection struct {
	win32.IInspectable
}

func (this *IVisualUnorderedCollection) Vtbl() *IVisualUnorderedCollectionVtbl {
	return (*IVisualUnorderedCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVisualUnorderedCollection) Get_Count() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Count, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisualUnorderedCollection) Add(newVisual *IVisual) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(newVisual)))
	_ = _hr
}

func (this *IVisualUnorderedCollection) Remove(visual *IVisual) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(visual)))
	_ = _hr
}

func (this *IVisualUnorderedCollection) RemoveAll() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveAll, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// classes

type CompositionObject struct {
	RtClass
	*ICompositionObject
}

func NewICompositionObjectStatics() *ICompositionObjectStatics {
	var p *ICompositionObjectStatics
	hs := NewHStr("Windows.UI.Composition.CompositionObject")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICompositionObjectStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type CompositionLight struct {
	RtClass
	*ICompositionLight
}

type AmbientLight struct {
	RtClass
	*IAmbientLight
}

type AnimationController struct {
	RtClass
	*IAnimationController
}

func NewIAnimationControllerStatics() *IAnimationControllerStatics {
	var p *IAnimationControllerStatics
	hs := NewHStr("Windows.UI.Composition.AnimationController")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAnimationControllerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type AnimationPropertyInfo struct {
	RtClass
	*IAnimationPropertyInfo
}

type CompositionEasingFunction struct {
	RtClass
	*ICompositionEasingFunction
}

func NewICompositionEasingFunctionStatics() *ICompositionEasingFunctionStatics {
	var p *ICompositionEasingFunctionStatics
	hs := NewHStr("Windows.UI.Composition.CompositionEasingFunction")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICompositionEasingFunctionStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type CompositionAnimation struct {
	RtClass
	*ICompositionAnimation
}

type KeyFrameAnimation struct {
	RtClass
	*IKeyFrameAnimation
}

type BooleanKeyFrameAnimation struct {
	RtClass
	*IBooleanKeyFrameAnimation
}

type NaturalMotionAnimation struct {
	RtClass
	*INaturalMotionAnimation
}

type ScalarNaturalMotionAnimation struct {
	RtClass
	*IScalarNaturalMotionAnimation
}

type BounceScalarNaturalMotionAnimation struct {
	RtClass
	*IBounceScalarNaturalMotionAnimation
}

type Vector2NaturalMotionAnimation struct {
	RtClass
	*IVector2NaturalMotionAnimation
}

type BounceVector2NaturalMotionAnimation struct {
	RtClass
	*IBounceVector2NaturalMotionAnimation
}

type Vector3NaturalMotionAnimation struct {
	RtClass
	*IVector3NaturalMotionAnimation
}

type BounceVector3NaturalMotionAnimation struct {
	RtClass
	*IBounceVector3NaturalMotionAnimation
}

type ColorKeyFrameAnimation struct {
	RtClass
	*IColorKeyFrameAnimation
}

type CompositionAnimationGroup struct {
	RtClass
	*ICompositionAnimationGroup
}

type CompositionBrush struct {
	RtClass
	*ICompositionBrush
}

type CompositionBackdropBrush struct {
	RtClass
	*ICompositionBackdropBrush
}

type CompositionBatchCompletedEventArgs struct {
	RtClass
	*ICompositionBatchCompletedEventArgs
}

type CompositionCapabilities struct {
	RtClass
	*ICompositionCapabilities
}

func NewICompositionCapabilitiesStatics() *ICompositionCapabilitiesStatics {
	var p *ICompositionCapabilitiesStatics
	hs := NewHStr("Windows.UI.Composition.CompositionCapabilities")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICompositionCapabilitiesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type CompositionClip struct {
	RtClass
	*ICompositionClip
}

type CompositionColorBrush struct {
	RtClass
	*ICompositionColorBrush
}

type CompositionCommitBatch struct {
	RtClass
	*ICompositionCommitBatch
}

type CompositionShape struct {
	RtClass
	*ICompositionShape
}

type CompositionContainerShape struct {
	RtClass
	*ICompositionContainerShape
}

type CompositionDrawingSurface struct {
	RtClass
	*ICompositionDrawingSurface
}

type CompositionEffectBrush struct {
	RtClass
	*ICompositionEffectBrush
}

type CompositionEffectFactory struct {
	RtClass
	*ICompositionEffectFactory
}

type CompositionEffectSourceParameter struct {
	RtClass
	*ICompositionEffectSourceParameter
}

func NewCompositionEffectSourceParameter_Create(name string) *CompositionEffectSourceParameter {
	hs := NewHStr("Windows.UI.Composition.CompositionEffectSourceParameter")
	var pFac *ICompositionEffectSourceParameterFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICompositionEffectSourceParameterFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ICompositionEffectSourceParameter
	p = pFac.Create(name)
	result := &CompositionEffectSourceParameter{
		RtClass:                           RtClass{PInspect: &p.IInspectable},
		ICompositionEffectSourceParameter: p,
	}
	com.AddToScope(result)
	return result
}

type CompositionGeometry struct {
	RtClass
	*ICompositionGeometry
}

type CompositionEllipseGeometry struct {
	RtClass
	*ICompositionEllipseGeometry
}

type CompositionGeometricClip struct {
	RtClass
	*ICompositionGeometricClip
}

type CompositionGraphicsDevice struct {
	RtClass
	*ICompositionGraphicsDevice
}

type CompositionLineGeometry struct {
	RtClass
	*ICompositionLineGeometry
}

type CompositionMaskBrush struct {
	RtClass
	*ICompositionMaskBrush
}

type CompositionMipmapSurface struct {
	RtClass
	*ICompositionMipmapSurface
}

type CompositionNineGridBrush struct {
	RtClass
	*ICompositionNineGridBrush
}

type CompositionPath struct {
	RtClass
	*ICompositionPath
}

func NewCompositionPath_Create(source unsafe.Pointer) *CompositionPath {
	hs := NewHStr("Windows.UI.Composition.CompositionPath")
	var pFac *ICompositionPathFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICompositionPathFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ICompositionPath
	p = pFac.Create(source)
	result := &CompositionPath{
		RtClass:          RtClass{PInspect: &p.IInspectable},
		ICompositionPath: p,
	}
	com.AddToScope(result)
	return result
}

type CompositionPathGeometry struct {
	RtClass
	*ICompositionPathGeometry
}

type CompositionPropertySet struct {
	RtClass
	*ICompositionPropertySet
}

type CompositionRectangleGeometry struct {
	RtClass
	*ICompositionRectangleGeometry
}

type CompositionRoundedRectangleGeometry struct {
	RtClass
	*ICompositionRoundedRectangleGeometry
}

type CompositionScopedBatch struct {
	RtClass
	*ICompositionScopedBatch
}

type CompositionShadow struct {
	RtClass
	*ICompositionShadow
}

type CompositionShapeCollection struct {
	RtClass
	*IVector[*ICompositionShape]
}

type CompositionSpriteShape struct {
	RtClass
	*ICompositionSpriteShape
}

type CompositionStrokeDashArray struct {
	RtClass
	*IVector[float32]
}

type CompositionSurfaceBrush struct {
	RtClass
	*ICompositionSurfaceBrush
}

type CompositionTarget struct {
	RtClass
	*ICompositionTarget
}

type CompositionTransform struct {
	RtClass
	*ICompositionTransform
}

type CompositionViewBox struct {
	RtClass
	*ICompositionViewBox
}

type CompositionVirtualDrawingSurface struct {
	RtClass
	*ICompositionVirtualDrawingSurface
}

type Compositor struct {
	RtClass
	*ICompositor
}

func NewCompositor() *Compositor {
	hs := NewHStr("Windows.UI.Composition.Compositor")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &Compositor{
		RtClass:     RtClass{PInspect: p},
		ICompositor: (*ICompositor)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewICompositorStatics() *ICompositorStatics {
	var p *ICompositorStatics
	hs := NewHStr("Windows.UI.Composition.Compositor")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICompositorStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type Visual struct {
	RtClass
	*IVisual
}

type ContainerVisual struct {
	RtClass
	*IContainerVisual
}

type CubicBezierEasingFunction struct {
	RtClass
	*ICubicBezierEasingFunction
}

type DistantLight struct {
	RtClass
	*IDistantLight
}

type DropShadow struct {
	RtClass
	*IDropShadow
}

type ExpressionAnimation struct {
	RtClass
	*IExpressionAnimation
}

type ImplicitAnimationCollection struct {
	RtClass
	*IImplicitAnimationCollection
}

type InitialValueExpressionCollection struct {
	RtClass
	*IMap[string, string]
}

type InsetClip struct {
	RtClass
	*IInsetClip
}

type LayerVisual struct {
	RtClass
	*ILayerVisual
}

type LinearEasingFunction struct {
	RtClass
	*ILinearEasingFunction
}

type PathKeyFrameAnimation struct {
	RtClass
	*IPathKeyFrameAnimation
}

type PointLight struct {
	RtClass
	*IPointLight
}

type QuaternionKeyFrameAnimation struct {
	RtClass
	*IQuaternionKeyFrameAnimation
}

type RedirectVisual struct {
	RtClass
	*IRedirectVisual
}

type RenderingDeviceReplacedEventArgs struct {
	RtClass
	*IRenderingDeviceReplacedEventArgs
}

type ScalarKeyFrameAnimation struct {
	RtClass
	*IScalarKeyFrameAnimation
}

type ShapeVisual struct {
	RtClass
	*IShapeVisual
}

type SpotLight struct {
	RtClass
	*ISpotLight
}

type SpringScalarNaturalMotionAnimation struct {
	RtClass
	*ISpringScalarNaturalMotionAnimation
}

type SpringVector2NaturalMotionAnimation struct {
	RtClass
	*ISpringVector2NaturalMotionAnimation
}

type SpringVector3NaturalMotionAnimation struct {
	RtClass
	*ISpringVector3NaturalMotionAnimation
}

type SpriteVisual struct {
	RtClass
	*ISpriteVisual
}

type StepEasingFunction struct {
	RtClass
	*IStepEasingFunction
}

type Vector2KeyFrameAnimation struct {
	RtClass
	*IVector2KeyFrameAnimation
}

type Vector3KeyFrameAnimation struct {
	RtClass
	*IVector3KeyFrameAnimation
}

type Vector4KeyFrameAnimation struct {
	RtClass
	*IVector4KeyFrameAnimation
}

type VisualCollection struct {
	RtClass
	*IVisualCollection
}

type VisualUnorderedCollection struct {
	RtClass
	*IVisualUnorderedCollection
}
