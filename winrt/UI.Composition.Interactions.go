package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
// flags
type InteractionBindingAxisModes uint32

const (
	InteractionBindingAxisModes_None      InteractionBindingAxisModes = 0
	InteractionBindingAxisModes_PositionX InteractionBindingAxisModes = 1
	InteractionBindingAxisModes_PositionY InteractionBindingAxisModes = 2
	InteractionBindingAxisModes_Scale     InteractionBindingAxisModes = 4
)

// enum
type InteractionChainingMode int32

const (
	InteractionChainingMode_Auto   InteractionChainingMode = 0
	InteractionChainingMode_Always InteractionChainingMode = 1
	InteractionChainingMode_Never  InteractionChainingMode = 2
)

// enum
type InteractionSourceMode int32

const (
	InteractionSourceMode_Disabled              InteractionSourceMode = 0
	InteractionSourceMode_EnabledWithInertia    InteractionSourceMode = 1
	InteractionSourceMode_EnabledWithoutInertia InteractionSourceMode = 2
)

// enum
type InteractionSourceRedirectionMode int32

const (
	InteractionSourceRedirectionMode_Disabled InteractionSourceRedirectionMode = 0
	InteractionSourceRedirectionMode_Enabled  InteractionSourceRedirectionMode = 1
)

// enum
type InteractionTrackerClampingOption int32

const (
	InteractionTrackerClampingOption_Auto     InteractionTrackerClampingOption = 0
	InteractionTrackerClampingOption_Disabled InteractionTrackerClampingOption = 1
)

// enum
type InteractionTrackerPositionUpdateOption int32

const (
	InteractionTrackerPositionUpdateOption_Default                         InteractionTrackerPositionUpdateOption = 0
	InteractionTrackerPositionUpdateOption_AllowActiveCustomScaleAnimation InteractionTrackerPositionUpdateOption = 1
)

// enum
type VisualInteractionSourceRedirectionMode int32

const (
	VisualInteractionSourceRedirectionMode_Off                            VisualInteractionSourceRedirectionMode = 0
	VisualInteractionSourceRedirectionMode_CapableTouchpadOnly            VisualInteractionSourceRedirectionMode = 1
	VisualInteractionSourceRedirectionMode_PointerWheelOnly               VisualInteractionSourceRedirectionMode = 2
	VisualInteractionSourceRedirectionMode_CapableTouchpadAndPointerWheel VisualInteractionSourceRedirectionMode = 3
)

// interfaces

// 43250538-EB73-4561-A71D-1A43EAEB7A9B
var IID_ICompositionConditionalValue = syscall.GUID{0x43250538, 0xEB73, 0x4561,
	[8]byte{0xA7, 0x1D, 0x1A, 0x43, 0xEA, 0xEB, 0x7A, 0x9B}}

type ICompositionConditionalValueInterface interface {
	win32.IInspectableInterface
	Get_Condition() *IExpressionAnimation
	Put_Condition(value *IExpressionAnimation)
	Get_Value() *IExpressionAnimation
	Put_Value(value *IExpressionAnimation)
}

type ICompositionConditionalValueVtbl struct {
	win32.IInspectableVtbl
	Get_Condition uintptr
	Put_Condition uintptr
	Get_Value     uintptr
	Put_Value     uintptr
}

type ICompositionConditionalValue struct {
	win32.IInspectable
}

func (this *ICompositionConditionalValue) Vtbl() *ICompositionConditionalValueVtbl {
	return (*ICompositionConditionalValueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionConditionalValue) Get_Condition() *IExpressionAnimation {
	var _result *IExpressionAnimation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Condition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionConditionalValue) Put_Condition(value *IExpressionAnimation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Condition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ICompositionConditionalValue) Get_Value() *IExpressionAnimation {
	var _result *IExpressionAnimation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositionConditionalValue) Put_Value(value *IExpressionAnimation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 090C4B72-8467-4D0A-9065-AC46B80A5522
var IID_ICompositionConditionalValueStatics = syscall.GUID{0x090C4B72, 0x8467, 0x4D0A,
	[8]byte{0x90, 0x65, 0xAC, 0x46, 0xB8, 0x0A, 0x55, 0x22}}

type ICompositionConditionalValueStaticsInterface interface {
	win32.IInspectableInterface
	Create(compositor *ICompositor) *ICompositionConditionalValue
}

type ICompositionConditionalValueStaticsVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ICompositionConditionalValueStatics struct {
	win32.IInspectable
}

func (this *ICompositionConditionalValueStatics) Vtbl() *ICompositionConditionalValueStaticsVtbl {
	return (*ICompositionConditionalValueStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionConditionalValueStatics) Create(compositor *ICompositor) *ICompositionConditionalValue {
	var _result *ICompositionConditionalValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(compositor)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 043B2431-06E3-495A-BA54-409F0017FAC0
var IID_ICompositionInteractionSource = syscall.GUID{0x043B2431, 0x06E3, 0x495A,
	[8]byte{0xBA, 0x54, 0x40, 0x9F, 0x00, 0x17, 0xFA, 0xC0}}

type ICompositionInteractionSourceInterface interface {
	win32.IInspectableInterface
}

type ICompositionInteractionSourceVtbl struct {
	win32.IInspectableVtbl
}

type ICompositionInteractionSource struct {
	win32.IInspectable
}

func (this *ICompositionInteractionSource) Vtbl() *ICompositionInteractionSourceVtbl {
	return (*ICompositionInteractionSourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 1B468E4B-A5BF-47D8-A547-3894155A158C
var IID_ICompositionInteractionSourceCollection = syscall.GUID{0x1B468E4B, 0xA5BF, 0x47D8,
	[8]byte{0xA5, 0x47, 0x38, 0x94, 0x15, 0x5A, 0x15, 0x8C}}

type ICompositionInteractionSourceCollectionInterface interface {
	win32.IInspectableInterface
	Get_Count() int32
	Add(value *ICompositionInteractionSource)
	Remove(value *ICompositionInteractionSource)
	RemoveAll()
}

type ICompositionInteractionSourceCollectionVtbl struct {
	win32.IInspectableVtbl
	Get_Count uintptr
	Add       uintptr
	Remove    uintptr
	RemoveAll uintptr
}

type ICompositionInteractionSourceCollection struct {
	win32.IInspectable
}

func (this *ICompositionInteractionSourceCollection) Vtbl() *ICompositionInteractionSourceCollectionVtbl {
	return (*ICompositionInteractionSourceCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositionInteractionSourceCollection) Get_Count() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Count, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICompositionInteractionSourceCollection) Add(value *ICompositionInteractionSource) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ICompositionInteractionSourceCollection) Remove(value *ICompositionInteractionSource) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ICompositionInteractionSourceCollection) RemoveAll() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveAll, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// A78347E5-A9D1-4D02-985E-B930CD0B9DA4
var IID_IInteractionSourceConfiguration = syscall.GUID{0xA78347E5, 0xA9D1, 0x4D02,
	[8]byte{0x98, 0x5E, 0xB9, 0x30, 0xCD, 0x0B, 0x9D, 0xA4}}

type IInteractionSourceConfigurationInterface interface {
	win32.IInspectableInterface
	Get_PositionXSourceMode() InteractionSourceRedirectionMode
	Put_PositionXSourceMode(value InteractionSourceRedirectionMode)
	Get_PositionYSourceMode() InteractionSourceRedirectionMode
	Put_PositionYSourceMode(value InteractionSourceRedirectionMode)
	Get_ScaleSourceMode() InteractionSourceRedirectionMode
	Put_ScaleSourceMode(value InteractionSourceRedirectionMode)
}

type IInteractionSourceConfigurationVtbl struct {
	win32.IInspectableVtbl
	Get_PositionXSourceMode uintptr
	Put_PositionXSourceMode uintptr
	Get_PositionYSourceMode uintptr
	Put_PositionYSourceMode uintptr
	Get_ScaleSourceMode     uintptr
	Put_ScaleSourceMode     uintptr
}

type IInteractionSourceConfiguration struct {
	win32.IInspectable
}

func (this *IInteractionSourceConfiguration) Vtbl() *IInteractionSourceConfigurationVtbl {
	return (*IInteractionSourceConfigurationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInteractionSourceConfiguration) Get_PositionXSourceMode() InteractionSourceRedirectionMode {
	var _result InteractionSourceRedirectionMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PositionXSourceMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInteractionSourceConfiguration) Put_PositionXSourceMode(value InteractionSourceRedirectionMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PositionXSourceMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IInteractionSourceConfiguration) Get_PositionYSourceMode() InteractionSourceRedirectionMode {
	var _result InteractionSourceRedirectionMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PositionYSourceMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInteractionSourceConfiguration) Put_PositionYSourceMode(value InteractionSourceRedirectionMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PositionYSourceMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IInteractionSourceConfiguration) Get_ScaleSourceMode() InteractionSourceRedirectionMode {
	var _result InteractionSourceRedirectionMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScaleSourceMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInteractionSourceConfiguration) Put_ScaleSourceMode(value InteractionSourceRedirectionMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ScaleSourceMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 2A8E8CB1-1000-4416-8363-CC27FB877308
var IID_IInteractionTracker = syscall.GUID{0x2A8E8CB1, 0x1000, 0x4416,
	[8]byte{0x83, 0x63, 0xCC, 0x27, 0xFB, 0x87, 0x73, 0x08}}

type IInteractionTrackerInterface interface {
	win32.IInspectableInterface
	Get_InteractionSources() *ICompositionInteractionSourceCollection
	Get_IsPositionRoundingSuggested() bool
	Get_MaxPosition() unsafe.Pointer
	Put_MaxPosition(value unsafe.Pointer)
	Get_MaxScale() float32
	Put_MaxScale(value float32)
	Get_MinPosition() unsafe.Pointer
	Put_MinPosition(value unsafe.Pointer)
	Get_MinScale() float32
	Put_MinScale(value float32)
	Get_NaturalRestingPosition() unsafe.Pointer
	Get_NaturalRestingScale() float32
	Get_Owner() *IInteractionTrackerOwner
	Get_Position() unsafe.Pointer
	Get_PositionInertiaDecayRate() *IReference[unsafe.Pointer]
	Put_PositionInertiaDecayRate(value *IReference[unsafe.Pointer])
	Get_PositionVelocityInPixelsPerSecond() unsafe.Pointer
	Get_Scale() float32
	Get_ScaleInertiaDecayRate() *IReference[float32]
	Put_ScaleInertiaDecayRate(value *IReference[float32])
	Get_ScaleVelocityInPercentPerSecond() float32
	AdjustPositionXIfGreaterThanThreshold(adjustment float32, positionThreshold float32)
	AdjustPositionYIfGreaterThanThreshold(adjustment float32, positionThreshold float32)
	ConfigurePositionXInertiaModifiers(modifiers *IIterable[*IInteractionTrackerInertiaModifier])
	ConfigurePositionYInertiaModifiers(modifiers *IIterable[*IInteractionTrackerInertiaModifier])
	ConfigureScaleInertiaModifiers(modifiers *IIterable[*IInteractionTrackerInertiaModifier])
	TryUpdatePosition(value unsafe.Pointer) int32
	TryUpdatePositionBy(amount unsafe.Pointer) int32
	TryUpdatePositionWithAnimation(animation *ICompositionAnimation) int32
	TryUpdatePositionWithAdditionalVelocity(velocityInPixelsPerSecond unsafe.Pointer) int32
	TryUpdateScale(value float32, centerPoint unsafe.Pointer) int32
	TryUpdateScaleWithAnimation(animation *ICompositionAnimation, centerPoint unsafe.Pointer) int32
	TryUpdateScaleWithAdditionalVelocity(velocityInPercentPerSecond float32, centerPoint unsafe.Pointer) int32
}

type IInteractionTrackerVtbl struct {
	win32.IInspectableVtbl
	Get_InteractionSources                  uintptr
	Get_IsPositionRoundingSuggested         uintptr
	Get_MaxPosition                         uintptr
	Put_MaxPosition                         uintptr
	Get_MaxScale                            uintptr
	Put_MaxScale                            uintptr
	Get_MinPosition                         uintptr
	Put_MinPosition                         uintptr
	Get_MinScale                            uintptr
	Put_MinScale                            uintptr
	Get_NaturalRestingPosition              uintptr
	Get_NaturalRestingScale                 uintptr
	Get_Owner                               uintptr
	Get_Position                            uintptr
	Get_PositionInertiaDecayRate            uintptr
	Put_PositionInertiaDecayRate            uintptr
	Get_PositionVelocityInPixelsPerSecond   uintptr
	Get_Scale                               uintptr
	Get_ScaleInertiaDecayRate               uintptr
	Put_ScaleInertiaDecayRate               uintptr
	Get_ScaleVelocityInPercentPerSecond     uintptr
	AdjustPositionXIfGreaterThanThreshold   uintptr
	AdjustPositionYIfGreaterThanThreshold   uintptr
	ConfigurePositionXInertiaModifiers      uintptr
	ConfigurePositionYInertiaModifiers      uintptr
	ConfigureScaleInertiaModifiers          uintptr
	TryUpdatePosition                       uintptr
	TryUpdatePositionBy                     uintptr
	TryUpdatePositionWithAnimation          uintptr
	TryUpdatePositionWithAdditionalVelocity uintptr
	TryUpdateScale                          uintptr
	TryUpdateScaleWithAnimation             uintptr
	TryUpdateScaleWithAdditionalVelocity    uintptr
}

type IInteractionTracker struct {
	win32.IInspectable
}

func (this *IInteractionTracker) Vtbl() *IInteractionTrackerVtbl {
	return (*IInteractionTrackerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInteractionTracker) Get_InteractionSources() *ICompositionInteractionSourceCollection {
	var _result *ICompositionInteractionSourceCollection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InteractionSources, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInteractionTracker) Get_IsPositionRoundingSuggested() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPositionRoundingSuggested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInteractionTracker) Get_MaxPosition() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxPosition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInteractionTracker) Put_MaxPosition(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxPosition, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IInteractionTracker) Get_MaxScale() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxScale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInteractionTracker) Put_MaxScale(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxScale, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IInteractionTracker) Get_MinPosition() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinPosition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInteractionTracker) Put_MinPosition(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MinPosition, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IInteractionTracker) Get_MinScale() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinScale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInteractionTracker) Put_MinScale(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MinScale, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IInteractionTracker) Get_NaturalRestingPosition() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NaturalRestingPosition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInteractionTracker) Get_NaturalRestingScale() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NaturalRestingScale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInteractionTracker) Get_Owner() *IInteractionTrackerOwner {
	var _result *IInteractionTrackerOwner
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Owner, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInteractionTracker) Get_Position() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInteractionTracker) Get_PositionInertiaDecayRate() *IReference[unsafe.Pointer] {
	var _result *IReference[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PositionInertiaDecayRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInteractionTracker) Put_PositionInertiaDecayRate(value *IReference[unsafe.Pointer]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PositionInertiaDecayRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IInteractionTracker) Get_PositionVelocityInPixelsPerSecond() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PositionVelocityInPixelsPerSecond, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInteractionTracker) Get_Scale() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Scale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInteractionTracker) Get_ScaleInertiaDecayRate() *IReference[float32] {
	var _result *IReference[float32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScaleInertiaDecayRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInteractionTracker) Put_ScaleInertiaDecayRate(value *IReference[float32]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ScaleInertiaDecayRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IInteractionTracker) Get_ScaleVelocityInPercentPerSecond() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScaleVelocityInPercentPerSecond, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInteractionTracker) AdjustPositionXIfGreaterThanThreshold(adjustment float32, positionThreshold float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AdjustPositionXIfGreaterThanThreshold, uintptr(unsafe.Pointer(this)), uintptr(adjustment), uintptr(positionThreshold))
	_ = _hr
}

func (this *IInteractionTracker) AdjustPositionYIfGreaterThanThreshold(adjustment float32, positionThreshold float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AdjustPositionYIfGreaterThanThreshold, uintptr(unsafe.Pointer(this)), uintptr(adjustment), uintptr(positionThreshold))
	_ = _hr
}

func (this *IInteractionTracker) ConfigurePositionXInertiaModifiers(modifiers *IIterable[*IInteractionTrackerInertiaModifier]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConfigurePositionXInertiaModifiers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(modifiers)))
	_ = _hr
}

func (this *IInteractionTracker) ConfigurePositionYInertiaModifiers(modifiers *IIterable[*IInteractionTrackerInertiaModifier]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConfigurePositionYInertiaModifiers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(modifiers)))
	_ = _hr
}

func (this *IInteractionTracker) ConfigureScaleInertiaModifiers(modifiers *IIterable[*IInteractionTrackerInertiaModifier]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConfigureScaleInertiaModifiers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(modifiers)))
	_ = _hr
}

func (this *IInteractionTracker) TryUpdatePosition(value unsafe.Pointer) int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryUpdatePosition, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInteractionTracker) TryUpdatePositionBy(amount unsafe.Pointer) int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryUpdatePositionBy, uintptr(unsafe.Pointer(this)), uintptr(amount), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInteractionTracker) TryUpdatePositionWithAnimation(animation *ICompositionAnimation) int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryUpdatePositionWithAnimation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(animation)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInteractionTracker) TryUpdatePositionWithAdditionalVelocity(velocityInPixelsPerSecond unsafe.Pointer) int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryUpdatePositionWithAdditionalVelocity, uintptr(unsafe.Pointer(this)), uintptr(velocityInPixelsPerSecond), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInteractionTracker) TryUpdateScale(value float32, centerPoint unsafe.Pointer) int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryUpdateScale, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(centerPoint), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInteractionTracker) TryUpdateScaleWithAnimation(animation *ICompositionAnimation, centerPoint unsafe.Pointer) int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryUpdateScaleWithAnimation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(animation)), uintptr(centerPoint), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInteractionTracker) TryUpdateScaleWithAdditionalVelocity(velocityInPercentPerSecond float32, centerPoint unsafe.Pointer) int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryUpdateScaleWithAdditionalVelocity, uintptr(unsafe.Pointer(this)), uintptr(velocityInPercentPerSecond), uintptr(centerPoint), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 25769A3E-CE6D-448C-8386-92620D240756
var IID_IInteractionTracker2 = syscall.GUID{0x25769A3E, 0xCE6D, 0x448C,
	[8]byte{0x83, 0x86, 0x92, 0x62, 0x0D, 0x24, 0x07, 0x56}}

type IInteractionTracker2Interface interface {
	win32.IInspectableInterface
	ConfigureCenterPointXInertiaModifiers(conditionalValues *IIterable[*ICompositionConditionalValue])
	ConfigureCenterPointYInertiaModifiers(conditionalValues *IIterable[*ICompositionConditionalValue])
}

type IInteractionTracker2Vtbl struct {
	win32.IInspectableVtbl
	ConfigureCenterPointXInertiaModifiers uintptr
	ConfigureCenterPointYInertiaModifiers uintptr
}

type IInteractionTracker2 struct {
	win32.IInspectable
}

func (this *IInteractionTracker2) Vtbl() *IInteractionTracker2Vtbl {
	return (*IInteractionTracker2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInteractionTracker2) ConfigureCenterPointXInertiaModifiers(conditionalValues *IIterable[*ICompositionConditionalValue]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConfigureCenterPointXInertiaModifiers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(conditionalValues)))
	_ = _hr
}

func (this *IInteractionTracker2) ConfigureCenterPointYInertiaModifiers(conditionalValues *IIterable[*ICompositionConditionalValue]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConfigureCenterPointYInertiaModifiers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(conditionalValues)))
	_ = _hr
}

// E6C5D7A2-5C4B-42C6-84B7-F69441B18091
var IID_IInteractionTracker3 = syscall.GUID{0xE6C5D7A2, 0x5C4B, 0x42C6,
	[8]byte{0x84, 0xB7, 0xF6, 0x94, 0x41, 0xB1, 0x80, 0x91}}

type IInteractionTracker3Interface interface {
	win32.IInspectableInterface
	ConfigureVector2PositionInertiaModifiers(modifiers *IIterable[*IInteractionTrackerVector2InertiaModifier])
}

type IInteractionTracker3Vtbl struct {
	win32.IInspectableVtbl
	ConfigureVector2PositionInertiaModifiers uintptr
}

type IInteractionTracker3 struct {
	win32.IInspectable
}

func (this *IInteractionTracker3) Vtbl() *IInteractionTracker3Vtbl {
	return (*IInteractionTracker3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInteractionTracker3) ConfigureVector2PositionInertiaModifiers(modifiers *IIterable[*IInteractionTrackerVector2InertiaModifier]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConfigureVector2PositionInertiaModifiers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(modifiers)))
	_ = _hr
}

// EBD222BC-04AF-4AC7-847D-06EA36E80A16
var IID_IInteractionTracker4 = syscall.GUID{0xEBD222BC, 0x04AF, 0x4AC7,
	[8]byte{0x84, 0x7D, 0x06, 0xEA, 0x36, 0xE8, 0x0A, 0x16}}

type IInteractionTracker4Interface interface {
	win32.IInspectableInterface
	TryUpdatePositionWithOption(value unsafe.Pointer, option InteractionTrackerClampingOption) int32
	TryUpdatePositionByWithOption(amount unsafe.Pointer, option InteractionTrackerClampingOption) int32
	Get_IsInertiaFromImpulse() bool
}

type IInteractionTracker4Vtbl struct {
	win32.IInspectableVtbl
	TryUpdatePositionWithOption   uintptr
	TryUpdatePositionByWithOption uintptr
	Get_IsInertiaFromImpulse      uintptr
}

type IInteractionTracker4 struct {
	win32.IInspectable
}

func (this *IInteractionTracker4) Vtbl() *IInteractionTracker4Vtbl {
	return (*IInteractionTracker4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInteractionTracker4) TryUpdatePositionWithOption(value unsafe.Pointer, option InteractionTrackerClampingOption) int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryUpdatePositionWithOption, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(option), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInteractionTracker4) TryUpdatePositionByWithOption(amount unsafe.Pointer, option InteractionTrackerClampingOption) int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryUpdatePositionByWithOption, uintptr(unsafe.Pointer(this)), uintptr(amount), uintptr(option), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInteractionTracker4) Get_IsInertiaFromImpulse() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsInertiaFromImpulse, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D3EF5DA2-A254-40E4-88D5-44E4E16B5809
var IID_IInteractionTracker5 = syscall.GUID{0xD3EF5DA2, 0xA254, 0x40E4,
	[8]byte{0x88, 0xD5, 0x44, 0xE4, 0xE1, 0x6B, 0x58, 0x09}}

type IInteractionTracker5Interface interface {
	win32.IInspectableInterface
	TryUpdatePositionWithOption(value unsafe.Pointer, option InteractionTrackerClampingOption, posUpdateOption InteractionTrackerPositionUpdateOption) int32
}

type IInteractionTracker5Vtbl struct {
	win32.IInspectableVtbl
	TryUpdatePositionWithOption uintptr
}

type IInteractionTracker5 struct {
	win32.IInspectable
}

func (this *IInteractionTracker5) Vtbl() *IInteractionTracker5Vtbl {
	return (*IInteractionTracker5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInteractionTracker5) TryUpdatePositionWithOption(value unsafe.Pointer, option InteractionTrackerClampingOption, posUpdateOption InteractionTrackerPositionUpdateOption) int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryUpdatePositionWithOption, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(option), uintptr(posUpdateOption), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 8D1C8CF1-D7B0-434C-A5D2-2D7611864834
var IID_IInteractionTrackerCustomAnimationStateEnteredArgs = syscall.GUID{0x8D1C8CF1, 0xD7B0, 0x434C,
	[8]byte{0xA5, 0xD2, 0x2D, 0x76, 0x11, 0x86, 0x48, 0x34}}

type IInteractionTrackerCustomAnimationStateEnteredArgsInterface interface {
	win32.IInspectableInterface
	Get_RequestId() int32
}

type IInteractionTrackerCustomAnimationStateEnteredArgsVtbl struct {
	win32.IInspectableVtbl
	Get_RequestId uintptr
}

type IInteractionTrackerCustomAnimationStateEnteredArgs struct {
	win32.IInspectable
}

func (this *IInteractionTrackerCustomAnimationStateEnteredArgs) Vtbl() *IInteractionTrackerCustomAnimationStateEnteredArgsVtbl {
	return (*IInteractionTrackerCustomAnimationStateEnteredArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInteractionTrackerCustomAnimationStateEnteredArgs) Get_RequestId() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 47D579B7-0985-5E99-B024-2F32C380C1A4
var IID_IInteractionTrackerCustomAnimationStateEnteredArgs2 = syscall.GUID{0x47D579B7, 0x0985, 0x5E99,
	[8]byte{0xB0, 0x24, 0x2F, 0x32, 0xC3, 0x80, 0xC1, 0xA4}}

type IInteractionTrackerCustomAnimationStateEnteredArgs2Interface interface {
	win32.IInspectableInterface
	Get_IsFromBinding() bool
}

type IInteractionTrackerCustomAnimationStateEnteredArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsFromBinding uintptr
}

type IInteractionTrackerCustomAnimationStateEnteredArgs2 struct {
	win32.IInspectable
}

func (this *IInteractionTrackerCustomAnimationStateEnteredArgs2) Vtbl() *IInteractionTrackerCustomAnimationStateEnteredArgs2Vtbl {
	return (*IInteractionTrackerCustomAnimationStateEnteredArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInteractionTrackerCustomAnimationStateEnteredArgs2) Get_IsFromBinding() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsFromBinding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 50012FAA-1510-4142-A1A5-019B09F8857B
var IID_IInteractionTrackerIdleStateEnteredArgs = syscall.GUID{0x50012FAA, 0x1510, 0x4142,
	[8]byte{0xA1, 0xA5, 0x01, 0x9B, 0x09, 0xF8, 0x85, 0x7B}}

type IInteractionTrackerIdleStateEnteredArgsInterface interface {
	win32.IInspectableInterface
	Get_RequestId() int32
}

type IInteractionTrackerIdleStateEnteredArgsVtbl struct {
	win32.IInspectableVtbl
	Get_RequestId uintptr
}

type IInteractionTrackerIdleStateEnteredArgs struct {
	win32.IInspectable
}

func (this *IInteractionTrackerIdleStateEnteredArgs) Vtbl() *IInteractionTrackerIdleStateEnteredArgsVtbl {
	return (*IInteractionTrackerIdleStateEnteredArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInteractionTrackerIdleStateEnteredArgs) Get_RequestId() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F2E771ED-B803-5137-9435-1C96E48721E9
var IID_IInteractionTrackerIdleStateEnteredArgs2 = syscall.GUID{0xF2E771ED, 0xB803, 0x5137,
	[8]byte{0x94, 0x35, 0x1C, 0x96, 0xE4, 0x87, 0x21, 0xE9}}

type IInteractionTrackerIdleStateEnteredArgs2Interface interface {
	win32.IInspectableInterface
	Get_IsFromBinding() bool
}

type IInteractionTrackerIdleStateEnteredArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsFromBinding uintptr
}

type IInteractionTrackerIdleStateEnteredArgs2 struct {
	win32.IInspectable
}

func (this *IInteractionTrackerIdleStateEnteredArgs2) Vtbl() *IInteractionTrackerIdleStateEnteredArgs2Vtbl {
	return (*IInteractionTrackerIdleStateEnteredArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInteractionTrackerIdleStateEnteredArgs2) Get_IsFromBinding() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsFromBinding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A0E2C920-26B4-4DA2-8B61-5E683979BBE2
var IID_IInteractionTrackerInertiaModifier = syscall.GUID{0xA0E2C920, 0x26B4, 0x4DA2,
	[8]byte{0x8B, 0x61, 0x5E, 0x68, 0x39, 0x79, 0xBB, 0xE2}}

type IInteractionTrackerInertiaModifierInterface interface {
	win32.IInspectableInterface
}

type IInteractionTrackerInertiaModifierVtbl struct {
	win32.IInspectableVtbl
}

type IInteractionTrackerInertiaModifier struct {
	win32.IInspectable
}

func (this *IInteractionTrackerInertiaModifier) Vtbl() *IInteractionTrackerInertiaModifierVtbl {
	return (*IInteractionTrackerInertiaModifierVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 993818FE-C94E-4B86-87F3-922665BA46B9
var IID_IInteractionTrackerInertiaModifierFactory = syscall.GUID{0x993818FE, 0xC94E, 0x4B86,
	[8]byte{0x87, 0xF3, 0x92, 0x26, 0x65, 0xBA, 0x46, 0xB9}}

type IInteractionTrackerInertiaModifierFactoryInterface interface {
	win32.IInspectableInterface
}

type IInteractionTrackerInertiaModifierFactoryVtbl struct {
	win32.IInspectableVtbl
}

type IInteractionTrackerInertiaModifierFactory struct {
	win32.IInspectable
}

func (this *IInteractionTrackerInertiaModifierFactory) Vtbl() *IInteractionTrackerInertiaModifierFactoryVtbl {
	return (*IInteractionTrackerInertiaModifierFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 04922FDC-F154-4CB8-BF33-CC1BA611E6DB
var IID_IInteractionTrackerInertiaMotion = syscall.GUID{0x04922FDC, 0xF154, 0x4CB8,
	[8]byte{0xBF, 0x33, 0xCC, 0x1B, 0xA6, 0x11, 0xE6, 0xDB}}

type IInteractionTrackerInertiaMotionInterface interface {
	win32.IInspectableInterface
	Get_Condition() *IExpressionAnimation
	Put_Condition(value *IExpressionAnimation)
	Get_Motion() *IExpressionAnimation
	Put_Motion(value *IExpressionAnimation)
}

type IInteractionTrackerInertiaMotionVtbl struct {
	win32.IInspectableVtbl
	Get_Condition uintptr
	Put_Condition uintptr
	Get_Motion    uintptr
	Put_Motion    uintptr
}

type IInteractionTrackerInertiaMotion struct {
	win32.IInspectable
}

func (this *IInteractionTrackerInertiaMotion) Vtbl() *IInteractionTrackerInertiaMotionVtbl {
	return (*IInteractionTrackerInertiaMotionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInteractionTrackerInertiaMotion) Get_Condition() *IExpressionAnimation {
	var _result *IExpressionAnimation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Condition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInteractionTrackerInertiaMotion) Put_Condition(value *IExpressionAnimation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Condition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IInteractionTrackerInertiaMotion) Get_Motion() *IExpressionAnimation {
	var _result *IExpressionAnimation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Motion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInteractionTrackerInertiaMotion) Put_Motion(value *IExpressionAnimation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Motion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 8CC83DD6-BA7B-431A-844B-6EAC9130F99A
var IID_IInteractionTrackerInertiaMotionStatics = syscall.GUID{0x8CC83DD6, 0xBA7B, 0x431A,
	[8]byte{0x84, 0x4B, 0x6E, 0xAC, 0x91, 0x30, 0xF9, 0x9A}}

type IInteractionTrackerInertiaMotionStaticsInterface interface {
	win32.IInspectableInterface
	Create(compositor *ICompositor) *IInteractionTrackerInertiaMotion
}

type IInteractionTrackerInertiaMotionStaticsVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IInteractionTrackerInertiaMotionStatics struct {
	win32.IInspectable
}

func (this *IInteractionTrackerInertiaMotionStatics) Vtbl() *IInteractionTrackerInertiaMotionStaticsVtbl {
	return (*IInteractionTrackerInertiaMotionStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInteractionTrackerInertiaMotionStatics) Create(compositor *ICompositor) *IInteractionTrackerInertiaMotion {
	var _result *IInteractionTrackerInertiaMotion
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(compositor)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 70ACDAAE-27DC-48ED-A3C3-6D61C9A029D2
var IID_IInteractionTrackerInertiaNaturalMotion = syscall.GUID{0x70ACDAAE, 0x27DC, 0x48ED,
	[8]byte{0xA3, 0xC3, 0x6D, 0x61, 0xC9, 0xA0, 0x29, 0xD2}}

type IInteractionTrackerInertiaNaturalMotionInterface interface {
	win32.IInspectableInterface
	Get_Condition() *IExpressionAnimation
	Put_Condition(value *IExpressionAnimation)
	Get_NaturalMotion() *IScalarNaturalMotionAnimation
	Put_NaturalMotion(value *IScalarNaturalMotionAnimation)
}

type IInteractionTrackerInertiaNaturalMotionVtbl struct {
	win32.IInspectableVtbl
	Get_Condition     uintptr
	Put_Condition     uintptr
	Get_NaturalMotion uintptr
	Put_NaturalMotion uintptr
}

type IInteractionTrackerInertiaNaturalMotion struct {
	win32.IInspectable
}

func (this *IInteractionTrackerInertiaNaturalMotion) Vtbl() *IInteractionTrackerInertiaNaturalMotionVtbl {
	return (*IInteractionTrackerInertiaNaturalMotionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInteractionTrackerInertiaNaturalMotion) Get_Condition() *IExpressionAnimation {
	var _result *IExpressionAnimation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Condition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInteractionTrackerInertiaNaturalMotion) Put_Condition(value *IExpressionAnimation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Condition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IInteractionTrackerInertiaNaturalMotion) Get_NaturalMotion() *IScalarNaturalMotionAnimation {
	var _result *IScalarNaturalMotionAnimation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NaturalMotion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInteractionTrackerInertiaNaturalMotion) Put_NaturalMotion(value *IScalarNaturalMotionAnimation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_NaturalMotion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// CFDA55B0-5E3E-4289-932D-EE5F50E74283
var IID_IInteractionTrackerInertiaNaturalMotionStatics = syscall.GUID{0xCFDA55B0, 0x5E3E, 0x4289,
	[8]byte{0x93, 0x2D, 0xEE, 0x5F, 0x50, 0xE7, 0x42, 0x83}}

type IInteractionTrackerInertiaNaturalMotionStaticsInterface interface {
	win32.IInspectableInterface
	Create(compositor *ICompositor) *IInteractionTrackerInertiaNaturalMotion
}

type IInteractionTrackerInertiaNaturalMotionStaticsVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IInteractionTrackerInertiaNaturalMotionStatics struct {
	win32.IInspectable
}

func (this *IInteractionTrackerInertiaNaturalMotionStatics) Vtbl() *IInteractionTrackerInertiaNaturalMotionStaticsVtbl {
	return (*IInteractionTrackerInertiaNaturalMotionStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInteractionTrackerInertiaNaturalMotionStatics) Create(compositor *ICompositor) *IInteractionTrackerInertiaNaturalMotion {
	var _result *IInteractionTrackerInertiaNaturalMotion
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(compositor)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 86F7EC09-5096-4170-9CC8-DF2FE101BB93
var IID_IInteractionTrackerInertiaRestingValue = syscall.GUID{0x86F7EC09, 0x5096, 0x4170,
	[8]byte{0x9C, 0xC8, 0xDF, 0x2F, 0xE1, 0x01, 0xBB, 0x93}}

type IInteractionTrackerInertiaRestingValueInterface interface {
	win32.IInspectableInterface
	Get_Condition() *IExpressionAnimation
	Put_Condition(value *IExpressionAnimation)
	Get_RestingValue() *IExpressionAnimation
	Put_RestingValue(value *IExpressionAnimation)
}

type IInteractionTrackerInertiaRestingValueVtbl struct {
	win32.IInspectableVtbl
	Get_Condition    uintptr
	Put_Condition    uintptr
	Get_RestingValue uintptr
	Put_RestingValue uintptr
}

type IInteractionTrackerInertiaRestingValue struct {
	win32.IInspectable
}

func (this *IInteractionTrackerInertiaRestingValue) Vtbl() *IInteractionTrackerInertiaRestingValueVtbl {
	return (*IInteractionTrackerInertiaRestingValueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInteractionTrackerInertiaRestingValue) Get_Condition() *IExpressionAnimation {
	var _result *IExpressionAnimation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Condition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInteractionTrackerInertiaRestingValue) Put_Condition(value *IExpressionAnimation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Condition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IInteractionTrackerInertiaRestingValue) Get_RestingValue() *IExpressionAnimation {
	var _result *IExpressionAnimation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RestingValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInteractionTrackerInertiaRestingValue) Put_RestingValue(value *IExpressionAnimation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RestingValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 18ED4699-0745-4096-BCAB-3A4E99569BCF
var IID_IInteractionTrackerInertiaRestingValueStatics = syscall.GUID{0x18ED4699, 0x0745, 0x4096,
	[8]byte{0xBC, 0xAB, 0x3A, 0x4E, 0x99, 0x56, 0x9B, 0xCF}}

type IInteractionTrackerInertiaRestingValueStaticsInterface interface {
	win32.IInspectableInterface
	Create(compositor *ICompositor) *IInteractionTrackerInertiaRestingValue
}

type IInteractionTrackerInertiaRestingValueStaticsVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IInteractionTrackerInertiaRestingValueStatics struct {
	win32.IInspectable
}

func (this *IInteractionTrackerInertiaRestingValueStatics) Vtbl() *IInteractionTrackerInertiaRestingValueStaticsVtbl {
	return (*IInteractionTrackerInertiaRestingValueStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInteractionTrackerInertiaRestingValueStatics) Create(compositor *ICompositor) *IInteractionTrackerInertiaRestingValue {
	var _result *IInteractionTrackerInertiaRestingValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(compositor)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 87108CF2-E7FF-4F7D-9FFD-D72F1E409B63
var IID_IInteractionTrackerInertiaStateEnteredArgs = syscall.GUID{0x87108CF2, 0xE7FF, 0x4F7D,
	[8]byte{0x9F, 0xFD, 0xD7, 0x2F, 0x1E, 0x40, 0x9B, 0x63}}

type IInteractionTrackerInertiaStateEnteredArgsInterface interface {
	win32.IInspectableInterface
	Get_ModifiedRestingPosition() *IReference[unsafe.Pointer]
	Get_ModifiedRestingScale() *IReference[float32]
	Get_NaturalRestingPosition() unsafe.Pointer
	Get_NaturalRestingScale() float32
	Get_PositionVelocityInPixelsPerSecond() unsafe.Pointer
	Get_RequestId() int32
	Get_ScaleVelocityInPercentPerSecond() float32
}

type IInteractionTrackerInertiaStateEnteredArgsVtbl struct {
	win32.IInspectableVtbl
	Get_ModifiedRestingPosition           uintptr
	Get_ModifiedRestingScale              uintptr
	Get_NaturalRestingPosition            uintptr
	Get_NaturalRestingScale               uintptr
	Get_PositionVelocityInPixelsPerSecond uintptr
	Get_RequestId                         uintptr
	Get_ScaleVelocityInPercentPerSecond   uintptr
}

type IInteractionTrackerInertiaStateEnteredArgs struct {
	win32.IInspectable
}

func (this *IInteractionTrackerInertiaStateEnteredArgs) Vtbl() *IInteractionTrackerInertiaStateEnteredArgsVtbl {
	return (*IInteractionTrackerInertiaStateEnteredArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInteractionTrackerInertiaStateEnteredArgs) Get_ModifiedRestingPosition() *IReference[unsafe.Pointer] {
	var _result *IReference[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ModifiedRestingPosition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInteractionTrackerInertiaStateEnteredArgs) Get_ModifiedRestingScale() *IReference[float32] {
	var _result *IReference[float32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ModifiedRestingScale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInteractionTrackerInertiaStateEnteredArgs) Get_NaturalRestingPosition() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NaturalRestingPosition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInteractionTrackerInertiaStateEnteredArgs) Get_NaturalRestingScale() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NaturalRestingScale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInteractionTrackerInertiaStateEnteredArgs) Get_PositionVelocityInPixelsPerSecond() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PositionVelocityInPixelsPerSecond, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInteractionTrackerInertiaStateEnteredArgs) Get_RequestId() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInteractionTrackerInertiaStateEnteredArgs) Get_ScaleVelocityInPercentPerSecond() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScaleVelocityInPercentPerSecond, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// B1EB32F6-C26C-41F6-A189-FABC22B323CC
var IID_IInteractionTrackerInertiaStateEnteredArgs2 = syscall.GUID{0xB1EB32F6, 0xC26C, 0x41F6,
	[8]byte{0xA1, 0x89, 0xFA, 0xBC, 0x22, 0xB3, 0x23, 0xCC}}

type IInteractionTrackerInertiaStateEnteredArgs2Interface interface {
	win32.IInspectableInterface
	Get_IsInertiaFromImpulse() bool
}

type IInteractionTrackerInertiaStateEnteredArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsInertiaFromImpulse uintptr
}

type IInteractionTrackerInertiaStateEnteredArgs2 struct {
	win32.IInspectable
}

func (this *IInteractionTrackerInertiaStateEnteredArgs2) Vtbl() *IInteractionTrackerInertiaStateEnteredArgs2Vtbl {
	return (*IInteractionTrackerInertiaStateEnteredArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInteractionTrackerInertiaStateEnteredArgs2) Get_IsInertiaFromImpulse() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsInertiaFromImpulse, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 48AC1C2F-47BD-59AF-A58C-79BD2EB9EF71
var IID_IInteractionTrackerInertiaStateEnteredArgs3 = syscall.GUID{0x48AC1C2F, 0x47BD, 0x59AF,
	[8]byte{0xA5, 0x8C, 0x79, 0xBD, 0x2E, 0xB9, 0xEF, 0x71}}

type IInteractionTrackerInertiaStateEnteredArgs3Interface interface {
	win32.IInspectableInterface
	Get_IsFromBinding() bool
}

type IInteractionTrackerInertiaStateEnteredArgs3Vtbl struct {
	win32.IInspectableVtbl
	Get_IsFromBinding uintptr
}

type IInteractionTrackerInertiaStateEnteredArgs3 struct {
	win32.IInspectable
}

func (this *IInteractionTrackerInertiaStateEnteredArgs3) Vtbl() *IInteractionTrackerInertiaStateEnteredArgs3Vtbl {
	return (*IInteractionTrackerInertiaStateEnteredArgs3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInteractionTrackerInertiaStateEnteredArgs3) Get_IsFromBinding() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsFromBinding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A7263939-A17B-4011-99FD-B5C24F143748
var IID_IInteractionTrackerInteractingStateEnteredArgs = syscall.GUID{0xA7263939, 0xA17B, 0x4011,
	[8]byte{0x99, 0xFD, 0xB5, 0xC2, 0x4F, 0x14, 0x37, 0x48}}

type IInteractionTrackerInteractingStateEnteredArgsInterface interface {
	win32.IInspectableInterface
	Get_RequestId() int32
}

type IInteractionTrackerInteractingStateEnteredArgsVtbl struct {
	win32.IInspectableVtbl
	Get_RequestId uintptr
}

type IInteractionTrackerInteractingStateEnteredArgs struct {
	win32.IInspectable
}

func (this *IInteractionTrackerInteractingStateEnteredArgs) Vtbl() *IInteractionTrackerInteractingStateEnteredArgsVtbl {
	return (*IInteractionTrackerInteractingStateEnteredArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInteractionTrackerInteractingStateEnteredArgs) Get_RequestId() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 509652D6-D488-59CD-819F-F52310295B11
var IID_IInteractionTrackerInteractingStateEnteredArgs2 = syscall.GUID{0x509652D6, 0xD488, 0x59CD,
	[8]byte{0x81, 0x9F, 0xF5, 0x23, 0x10, 0x29, 0x5B, 0x11}}

type IInteractionTrackerInteractingStateEnteredArgs2Interface interface {
	win32.IInspectableInterface
	Get_IsFromBinding() bool
}

type IInteractionTrackerInteractingStateEnteredArgs2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsFromBinding uintptr
}

type IInteractionTrackerInteractingStateEnteredArgs2 struct {
	win32.IInspectable
}

func (this *IInteractionTrackerInteractingStateEnteredArgs2) Vtbl() *IInteractionTrackerInteractingStateEnteredArgs2Vtbl {
	return (*IInteractionTrackerInteractingStateEnteredArgs2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInteractionTrackerInteractingStateEnteredArgs2) Get_IsFromBinding() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsFromBinding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// DB2E8AF3-4DEB-4E53-B29C-B06C9F96D651
var IID_IInteractionTrackerOwner = syscall.GUID{0xDB2E8AF3, 0x4DEB, 0x4E53,
	[8]byte{0xB2, 0x9C, 0xB0, 0x6C, 0x9F, 0x96, 0xD6, 0x51}}

type IInteractionTrackerOwnerInterface interface {
	win32.IInspectableInterface
	CustomAnimationStateEntered(sender *IInteractionTracker, args *IInteractionTrackerCustomAnimationStateEnteredArgs)
	IdleStateEntered(sender *IInteractionTracker, args *IInteractionTrackerIdleStateEnteredArgs)
	InertiaStateEntered(sender *IInteractionTracker, args *IInteractionTrackerInertiaStateEnteredArgs)
	InteractingStateEntered(sender *IInteractionTracker, args *IInteractionTrackerInteractingStateEnteredArgs)
	RequestIgnored(sender *IInteractionTracker, args *IInteractionTrackerRequestIgnoredArgs)
	ValuesChanged(sender *IInteractionTracker, args *IInteractionTrackerValuesChangedArgs)
}

type IInteractionTrackerOwnerVtbl struct {
	win32.IInspectableVtbl
	CustomAnimationStateEntered uintptr
	IdleStateEntered            uintptr
	InertiaStateEntered         uintptr
	InteractingStateEntered     uintptr
	RequestIgnored              uintptr
	ValuesChanged               uintptr
}

type IInteractionTrackerOwner struct {
	win32.IInspectable
}

func (this *IInteractionTrackerOwner) Vtbl() *IInteractionTrackerOwnerVtbl {
	return (*IInteractionTrackerOwnerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInteractionTrackerOwner) CustomAnimationStateEntered(sender *IInteractionTracker, args *IInteractionTrackerCustomAnimationStateEnteredArgs) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CustomAnimationStateEntered, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sender)), uintptr(unsafe.Pointer(args)))
	_ = _hr
}

func (this *IInteractionTrackerOwner) IdleStateEntered(sender *IInteractionTracker, args *IInteractionTrackerIdleStateEnteredArgs) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IdleStateEntered, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sender)), uintptr(unsafe.Pointer(args)))
	_ = _hr
}

func (this *IInteractionTrackerOwner) InertiaStateEntered(sender *IInteractionTracker, args *IInteractionTrackerInertiaStateEnteredArgs) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InertiaStateEntered, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sender)), uintptr(unsafe.Pointer(args)))
	_ = _hr
}

func (this *IInteractionTrackerOwner) InteractingStateEntered(sender *IInteractionTracker, args *IInteractionTrackerInteractingStateEnteredArgs) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InteractingStateEntered, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sender)), uintptr(unsafe.Pointer(args)))
	_ = _hr
}

func (this *IInteractionTrackerOwner) RequestIgnored(sender *IInteractionTracker, args *IInteractionTrackerRequestIgnoredArgs) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestIgnored, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sender)), uintptr(unsafe.Pointer(args)))
	_ = _hr
}

func (this *IInteractionTrackerOwner) ValuesChanged(sender *IInteractionTracker, args *IInteractionTrackerValuesChangedArgs) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ValuesChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sender)), uintptr(unsafe.Pointer(args)))
	_ = _hr
}

// 80DD82F1-CE25-488F-91DD-CB6455CCFF2E
var IID_IInteractionTrackerRequestIgnoredArgs = syscall.GUID{0x80DD82F1, 0xCE25, 0x488F,
	[8]byte{0x91, 0xDD, 0xCB, 0x64, 0x55, 0xCC, 0xFF, 0x2E}}

type IInteractionTrackerRequestIgnoredArgsInterface interface {
	win32.IInspectableInterface
	Get_RequestId() int32
}

type IInteractionTrackerRequestIgnoredArgsVtbl struct {
	win32.IInspectableVtbl
	Get_RequestId uintptr
}

type IInteractionTrackerRequestIgnoredArgs struct {
	win32.IInspectable
}

func (this *IInteractionTrackerRequestIgnoredArgs) Vtbl() *IInteractionTrackerRequestIgnoredArgsVtbl {
	return (*IInteractionTrackerRequestIgnoredArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInteractionTrackerRequestIgnoredArgs) Get_RequestId() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// BBA5D7B7-6590-4498-8D6C-EB62B514C92A
var IID_IInteractionTrackerStatics = syscall.GUID{0xBBA5D7B7, 0x6590, 0x4498,
	[8]byte{0x8D, 0x6C, 0xEB, 0x62, 0xB5, 0x14, 0xC9, 0x2A}}

type IInteractionTrackerStaticsInterface interface {
	win32.IInspectableInterface
	Create(compositor *ICompositor) *IInteractionTracker
	CreateWithOwner(compositor *ICompositor, owner *IInteractionTrackerOwner) *IInteractionTracker
}

type IInteractionTrackerStaticsVtbl struct {
	win32.IInspectableVtbl
	Create          uintptr
	CreateWithOwner uintptr
}

type IInteractionTrackerStatics struct {
	win32.IInspectable
}

func (this *IInteractionTrackerStatics) Vtbl() *IInteractionTrackerStaticsVtbl {
	return (*IInteractionTrackerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInteractionTrackerStatics) Create(compositor *ICompositor) *IInteractionTracker {
	var _result *IInteractionTracker
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(compositor)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInteractionTrackerStatics) CreateWithOwner(compositor *ICompositor, owner *IInteractionTrackerOwner) *IInteractionTracker {
	var _result *IInteractionTracker
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithOwner, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(compositor)), uintptr(unsafe.Pointer(owner)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 35E53720-46B7-5CB0-B505-F3D6884A6163
var IID_IInteractionTrackerStatics2 = syscall.GUID{0x35E53720, 0x46B7, 0x5CB0,
	[8]byte{0xB5, 0x05, 0xF3, 0xD6, 0x88, 0x4A, 0x61, 0x63}}

type IInteractionTrackerStatics2Interface interface {
	win32.IInspectableInterface
	SetBindingMode(boundTracker1 *IInteractionTracker, boundTracker2 *IInteractionTracker, axisMode InteractionBindingAxisModes)
	GetBindingMode(boundTracker1 *IInteractionTracker, boundTracker2 *IInteractionTracker) InteractionBindingAxisModes
}

type IInteractionTrackerStatics2Vtbl struct {
	win32.IInspectableVtbl
	SetBindingMode uintptr
	GetBindingMode uintptr
}

type IInteractionTrackerStatics2 struct {
	win32.IInspectable
}

func (this *IInteractionTrackerStatics2) Vtbl() *IInteractionTrackerStatics2Vtbl {
	return (*IInteractionTrackerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInteractionTrackerStatics2) SetBindingMode(boundTracker1 *IInteractionTracker, boundTracker2 *IInteractionTracker, axisMode InteractionBindingAxisModes) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetBindingMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(boundTracker1)), uintptr(unsafe.Pointer(boundTracker2)), uintptr(axisMode))
	_ = _hr
}

func (this *IInteractionTrackerStatics2) GetBindingMode(boundTracker1 *IInteractionTracker, boundTracker2 *IInteractionTracker) InteractionBindingAxisModes {
	var _result InteractionBindingAxisModes
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetBindingMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(boundTracker1)), uintptr(unsafe.Pointer(boundTracker2)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// CF1578EF-D3DF-4501-B9E6-F02FB22F73D0
var IID_IInteractionTrackerValuesChangedArgs = syscall.GUID{0xCF1578EF, 0xD3DF, 0x4501,
	[8]byte{0xB9, 0xE6, 0xF0, 0x2F, 0xB2, 0x2F, 0x73, 0xD0}}

type IInteractionTrackerValuesChangedArgsInterface interface {
	win32.IInspectableInterface
	Get_Position() unsafe.Pointer
	Get_RequestId() int32
	Get_Scale() float32
}

type IInteractionTrackerValuesChangedArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Position  uintptr
	Get_RequestId uintptr
	Get_Scale     uintptr
}

type IInteractionTrackerValuesChangedArgs struct {
	win32.IInspectable
}

func (this *IInteractionTrackerValuesChangedArgs) Vtbl() *IInteractionTrackerValuesChangedArgsVtbl {
	return (*IInteractionTrackerValuesChangedArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInteractionTrackerValuesChangedArgs) Get_Position() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInteractionTrackerValuesChangedArgs) Get_RequestId() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequestId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInteractionTrackerValuesChangedArgs) Get_Scale() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Scale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 87E08AB0-3086-4853-A4B7-77882AD5D7E3
var IID_IInteractionTrackerVector2InertiaModifier = syscall.GUID{0x87E08AB0, 0x3086, 0x4853,
	[8]byte{0xA4, 0xB7, 0x77, 0x88, 0x2A, 0xD5, 0xD7, 0xE3}}

type IInteractionTrackerVector2InertiaModifierInterface interface {
	win32.IInspectableInterface
}

type IInteractionTrackerVector2InertiaModifierVtbl struct {
	win32.IInspectableVtbl
}

type IInteractionTrackerVector2InertiaModifier struct {
	win32.IInspectable
}

func (this *IInteractionTrackerVector2InertiaModifier) Vtbl() *IInteractionTrackerVector2InertiaModifierVtbl {
	return (*IInteractionTrackerVector2InertiaModifierVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 7401D6C4-6C6D-48DF-BC3E-171E227E7D7F
var IID_IInteractionTrackerVector2InertiaModifierFactory = syscall.GUID{0x7401D6C4, 0x6C6D, 0x48DF,
	[8]byte{0xBC, 0x3E, 0x17, 0x1E, 0x22, 0x7E, 0x7D, 0x7F}}

type IInteractionTrackerVector2InertiaModifierFactoryInterface interface {
	win32.IInspectableInterface
}

type IInteractionTrackerVector2InertiaModifierFactoryVtbl struct {
	win32.IInspectableVtbl
}

type IInteractionTrackerVector2InertiaModifierFactory struct {
	win32.IInspectable
}

func (this *IInteractionTrackerVector2InertiaModifierFactory) Vtbl() *IInteractionTrackerVector2InertiaModifierFactoryVtbl {
	return (*IInteractionTrackerVector2InertiaModifierFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 5F17695C-162D-4C07-9400-C282B28276CA
var IID_IInteractionTrackerVector2InertiaNaturalMotion = syscall.GUID{0x5F17695C, 0x162D, 0x4C07,
	[8]byte{0x94, 0x00, 0xC2, 0x82, 0xB2, 0x82, 0x76, 0xCA}}

type IInteractionTrackerVector2InertiaNaturalMotionInterface interface {
	win32.IInspectableInterface
	Get_Condition() *IExpressionAnimation
	Put_Condition(value *IExpressionAnimation)
	Get_NaturalMotion() *IVector2NaturalMotionAnimation
	Put_NaturalMotion(value *IVector2NaturalMotionAnimation)
}

type IInteractionTrackerVector2InertiaNaturalMotionVtbl struct {
	win32.IInspectableVtbl
	Get_Condition     uintptr
	Put_Condition     uintptr
	Get_NaturalMotion uintptr
	Put_NaturalMotion uintptr
}

type IInteractionTrackerVector2InertiaNaturalMotion struct {
	win32.IInspectable
}

func (this *IInteractionTrackerVector2InertiaNaturalMotion) Vtbl() *IInteractionTrackerVector2InertiaNaturalMotionVtbl {
	return (*IInteractionTrackerVector2InertiaNaturalMotionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInteractionTrackerVector2InertiaNaturalMotion) Get_Condition() *IExpressionAnimation {
	var _result *IExpressionAnimation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Condition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInteractionTrackerVector2InertiaNaturalMotion) Put_Condition(value *IExpressionAnimation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Condition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IInteractionTrackerVector2InertiaNaturalMotion) Get_NaturalMotion() *IVector2NaturalMotionAnimation {
	var _result *IVector2NaturalMotionAnimation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NaturalMotion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInteractionTrackerVector2InertiaNaturalMotion) Put_NaturalMotion(value *IVector2NaturalMotionAnimation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_NaturalMotion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 82001A48-09C0-434F-8189-141C66DF362F
var IID_IInteractionTrackerVector2InertiaNaturalMotionStatics = syscall.GUID{0x82001A48, 0x09C0, 0x434F,
	[8]byte{0x81, 0x89, 0x14, 0x1C, 0x66, 0xDF, 0x36, 0x2F}}

type IInteractionTrackerVector2InertiaNaturalMotionStaticsInterface interface {
	win32.IInspectableInterface
	Create(compositor *ICompositor) *IInteractionTrackerVector2InertiaNaturalMotion
}

type IInteractionTrackerVector2InertiaNaturalMotionStaticsVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IInteractionTrackerVector2InertiaNaturalMotionStatics struct {
	win32.IInspectable
}

func (this *IInteractionTrackerVector2InertiaNaturalMotionStatics) Vtbl() *IInteractionTrackerVector2InertiaNaturalMotionStaticsVtbl {
	return (*IInteractionTrackerVector2InertiaNaturalMotionStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInteractionTrackerVector2InertiaNaturalMotionStatics) Create(compositor *ICompositor) *IInteractionTrackerVector2InertiaNaturalMotion {
	var _result *IInteractionTrackerVector2InertiaNaturalMotion
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(compositor)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CA0E8A86-D8D6-4111-B088-70347BD2B0ED
var IID_IVisualInteractionSource = syscall.GUID{0xCA0E8A86, 0xD8D6, 0x4111,
	[8]byte{0xB0, 0x88, 0x70, 0x34, 0x7B, 0xD2, 0xB0, 0xED}}

type IVisualInteractionSourceInterface interface {
	win32.IInspectableInterface
	Get_IsPositionXRailsEnabled() bool
	Put_IsPositionXRailsEnabled(value bool)
	Get_IsPositionYRailsEnabled() bool
	Put_IsPositionYRailsEnabled(value bool)
	Get_ManipulationRedirectionMode() VisualInteractionSourceRedirectionMode
	Put_ManipulationRedirectionMode(value VisualInteractionSourceRedirectionMode)
	Get_PositionXChainingMode() InteractionChainingMode
	Put_PositionXChainingMode(value InteractionChainingMode)
	Get_PositionXSourceMode() InteractionSourceMode
	Put_PositionXSourceMode(value InteractionSourceMode)
	Get_PositionYChainingMode() InteractionChainingMode
	Put_PositionYChainingMode(value InteractionChainingMode)
	Get_PositionYSourceMode() InteractionSourceMode
	Put_PositionYSourceMode(value InteractionSourceMode)
	Get_ScaleChainingMode() InteractionChainingMode
	Put_ScaleChainingMode(value InteractionChainingMode)
	Get_ScaleSourceMode() InteractionSourceMode
	Put_ScaleSourceMode(value InteractionSourceMode)
	Get_Source() *IVisual
	TryRedirectForManipulation(pointerPoint *IPointerPoint)
}

type IVisualInteractionSourceVtbl struct {
	win32.IInspectableVtbl
	Get_IsPositionXRailsEnabled     uintptr
	Put_IsPositionXRailsEnabled     uintptr
	Get_IsPositionYRailsEnabled     uintptr
	Put_IsPositionYRailsEnabled     uintptr
	Get_ManipulationRedirectionMode uintptr
	Put_ManipulationRedirectionMode uintptr
	Get_PositionXChainingMode       uintptr
	Put_PositionXChainingMode       uintptr
	Get_PositionXSourceMode         uintptr
	Put_PositionXSourceMode         uintptr
	Get_PositionYChainingMode       uintptr
	Put_PositionYChainingMode       uintptr
	Get_PositionYSourceMode         uintptr
	Put_PositionYSourceMode         uintptr
	Get_ScaleChainingMode           uintptr
	Put_ScaleChainingMode           uintptr
	Get_ScaleSourceMode             uintptr
	Put_ScaleSourceMode             uintptr
	Get_Source                      uintptr
	TryRedirectForManipulation      uintptr
}

type IVisualInteractionSource struct {
	win32.IInspectable
}

func (this *IVisualInteractionSource) Vtbl() *IVisualInteractionSourceVtbl {
	return (*IVisualInteractionSourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVisualInteractionSource) Get_IsPositionXRailsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPositionXRailsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisualInteractionSource) Put_IsPositionXRailsEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsPositionXRailsEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IVisualInteractionSource) Get_IsPositionYRailsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPositionYRailsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisualInteractionSource) Put_IsPositionYRailsEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsPositionYRailsEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IVisualInteractionSource) Get_ManipulationRedirectionMode() VisualInteractionSourceRedirectionMode {
	var _result VisualInteractionSourceRedirectionMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ManipulationRedirectionMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisualInteractionSource) Put_ManipulationRedirectionMode(value VisualInteractionSourceRedirectionMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ManipulationRedirectionMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVisualInteractionSource) Get_PositionXChainingMode() InteractionChainingMode {
	var _result InteractionChainingMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PositionXChainingMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisualInteractionSource) Put_PositionXChainingMode(value InteractionChainingMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PositionXChainingMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVisualInteractionSource) Get_PositionXSourceMode() InteractionSourceMode {
	var _result InteractionSourceMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PositionXSourceMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisualInteractionSource) Put_PositionXSourceMode(value InteractionSourceMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PositionXSourceMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVisualInteractionSource) Get_PositionYChainingMode() InteractionChainingMode {
	var _result InteractionChainingMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PositionYChainingMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisualInteractionSource) Put_PositionYChainingMode(value InteractionChainingMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PositionYChainingMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVisualInteractionSource) Get_PositionYSourceMode() InteractionSourceMode {
	var _result InteractionSourceMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PositionYSourceMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisualInteractionSource) Put_PositionYSourceMode(value InteractionSourceMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PositionYSourceMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVisualInteractionSource) Get_ScaleChainingMode() InteractionChainingMode {
	var _result InteractionChainingMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScaleChainingMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisualInteractionSource) Put_ScaleChainingMode(value InteractionChainingMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ScaleChainingMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVisualInteractionSource) Get_ScaleSourceMode() InteractionSourceMode {
	var _result InteractionSourceMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScaleSourceMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisualInteractionSource) Put_ScaleSourceMode(value InteractionSourceMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ScaleSourceMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVisualInteractionSource) Get_Source() *IVisual {
	var _result *IVisual
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Source, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVisualInteractionSource) TryRedirectForManipulation(pointerPoint *IPointerPoint) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryRedirectForManipulation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pointerPoint)))
	_ = _hr
}

// AA914893-A73C-414D-80D0-249BAD2FBD93
var IID_IVisualInteractionSource2 = syscall.GUID{0xAA914893, 0xA73C, 0x414D,
	[8]byte{0x80, 0xD0, 0x24, 0x9B, 0xAD, 0x2F, 0xBD, 0x93}}

type IVisualInteractionSource2Interface interface {
	win32.IInspectableInterface
	Get_DeltaPosition() unsafe.Pointer
	Get_DeltaScale() float32
	Get_Position() unsafe.Pointer
	Get_PositionVelocity() unsafe.Pointer
	Get_Scale() float32
	Get_ScaleVelocity() float32
	ConfigureCenterPointXModifiers(conditionalValues *IIterable[*ICompositionConditionalValue])
	ConfigureCenterPointYModifiers(conditionalValues *IIterable[*ICompositionConditionalValue])
	ConfigureDeltaPositionXModifiers(conditionalValues *IIterable[*ICompositionConditionalValue])
	ConfigureDeltaPositionYModifiers(conditionalValues *IIterable[*ICompositionConditionalValue])
	ConfigureDeltaScaleModifiers(conditionalValues *IIterable[*ICompositionConditionalValue])
}

type IVisualInteractionSource2Vtbl struct {
	win32.IInspectableVtbl
	Get_DeltaPosition                uintptr
	Get_DeltaScale                   uintptr
	Get_Position                     uintptr
	Get_PositionVelocity             uintptr
	Get_Scale                        uintptr
	Get_ScaleVelocity                uintptr
	ConfigureCenterPointXModifiers   uintptr
	ConfigureCenterPointYModifiers   uintptr
	ConfigureDeltaPositionXModifiers uintptr
	ConfigureDeltaPositionYModifiers uintptr
	ConfigureDeltaScaleModifiers     uintptr
}

type IVisualInteractionSource2 struct {
	win32.IInspectable
}

func (this *IVisualInteractionSource2) Vtbl() *IVisualInteractionSource2Vtbl {
	return (*IVisualInteractionSource2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVisualInteractionSource2) Get_DeltaPosition() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeltaPosition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisualInteractionSource2) Get_DeltaScale() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeltaScale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisualInteractionSource2) Get_Position() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisualInteractionSource2) Get_PositionVelocity() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PositionVelocity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisualInteractionSource2) Get_Scale() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Scale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisualInteractionSource2) Get_ScaleVelocity() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScaleVelocity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVisualInteractionSource2) ConfigureCenterPointXModifiers(conditionalValues *IIterable[*ICompositionConditionalValue]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConfigureCenterPointXModifiers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(conditionalValues)))
	_ = _hr
}

func (this *IVisualInteractionSource2) ConfigureCenterPointYModifiers(conditionalValues *IIterable[*ICompositionConditionalValue]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConfigureCenterPointYModifiers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(conditionalValues)))
	_ = _hr
}

func (this *IVisualInteractionSource2) ConfigureDeltaPositionXModifiers(conditionalValues *IIterable[*ICompositionConditionalValue]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConfigureDeltaPositionXModifiers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(conditionalValues)))
	_ = _hr
}

func (this *IVisualInteractionSource2) ConfigureDeltaPositionYModifiers(conditionalValues *IIterable[*ICompositionConditionalValue]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConfigureDeltaPositionYModifiers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(conditionalValues)))
	_ = _hr
}

func (this *IVisualInteractionSource2) ConfigureDeltaScaleModifiers(conditionalValues *IIterable[*ICompositionConditionalValue]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConfigureDeltaScaleModifiers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(conditionalValues)))
	_ = _hr
}

// D941EF2A-0D5C-4057-92D7-C9711533204F
var IID_IVisualInteractionSource3 = syscall.GUID{0xD941EF2A, 0x0D5C, 0x4057,
	[8]byte{0x92, 0xD7, 0xC9, 0x71, 0x15, 0x33, 0x20, 0x4F}}

type IVisualInteractionSource3Interface interface {
	win32.IInspectableInterface
	Get_PointerWheelConfig() *IInteractionSourceConfiguration
}

type IVisualInteractionSource3Vtbl struct {
	win32.IInspectableVtbl
	Get_PointerWheelConfig uintptr
}

type IVisualInteractionSource3 struct {
	win32.IInspectable
}

func (this *IVisualInteractionSource3) Vtbl() *IVisualInteractionSource3Vtbl {
	return (*IVisualInteractionSource3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVisualInteractionSource3) Get_PointerWheelConfig() *IInteractionSourceConfiguration {
	var _result *IInteractionSourceConfiguration
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PointerWheelConfig, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B2CA917C-E98A-41F2-B3C9-891C9266C8F6
var IID_IVisualInteractionSourceObjectFactory = syscall.GUID{0xB2CA917C, 0xE98A, 0x41F2,
	[8]byte{0xB3, 0xC9, 0x89, 0x1C, 0x92, 0x66, 0xC8, 0xF6}}

type IVisualInteractionSourceObjectFactoryInterface interface {
	win32.IInspectableInterface
}

type IVisualInteractionSourceObjectFactoryVtbl struct {
	win32.IInspectableVtbl
}

type IVisualInteractionSourceObjectFactory struct {
	win32.IInspectable
}

func (this *IVisualInteractionSourceObjectFactory) Vtbl() *IVisualInteractionSourceObjectFactoryVtbl {
	return (*IVisualInteractionSourceObjectFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 369965E1-8645-4F75-BA00-6479CD10C8E6
var IID_IVisualInteractionSourceStatics = syscall.GUID{0x369965E1, 0x8645, 0x4F75,
	[8]byte{0xBA, 0x00, 0x64, 0x79, 0xCD, 0x10, 0xC8, 0xE6}}

type IVisualInteractionSourceStaticsInterface interface {
	win32.IInspectableInterface
	Create(source *IVisual) *IVisualInteractionSource
}

type IVisualInteractionSourceStaticsVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IVisualInteractionSourceStatics struct {
	win32.IInspectable
}

func (this *IVisualInteractionSourceStatics) Vtbl() *IVisualInteractionSourceStaticsVtbl {
	return (*IVisualInteractionSourceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVisualInteractionSourceStatics) Create(source *IVisual) *IVisualInteractionSource {
	var _result *IVisualInteractionSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(source)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A979C032-5764-55E0-BC1F-0778786DCFDE
var IID_IVisualInteractionSourceStatics2 = syscall.GUID{0xA979C032, 0x5764, 0x55E0,
	[8]byte{0xBC, 0x1F, 0x07, 0x78, 0x78, 0x6D, 0xCF, 0xDE}}

type IVisualInteractionSourceStatics2Interface interface {
	win32.IInspectableInterface
	CreateFromIVisualElement(source *IVisualElement) *IVisualInteractionSource
}

type IVisualInteractionSourceStatics2Vtbl struct {
	win32.IInspectableVtbl
	CreateFromIVisualElement uintptr
}

type IVisualInteractionSourceStatics2 struct {
	win32.IInspectable
}

func (this *IVisualInteractionSourceStatics2) Vtbl() *IVisualInteractionSourceStatics2Vtbl {
	return (*IVisualInteractionSourceStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVisualInteractionSourceStatics2) CreateFromIVisualElement(source *IVisualElement) *IVisualInteractionSource {
	var _result *IVisualInteractionSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromIVisualElement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(source)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type CompositionConditionalValue struct {
	RtClass
	*ICompositionConditionalValue
}

func NewICompositionConditionalValueStatics() *ICompositionConditionalValueStatics {
	var p *ICompositionConditionalValueStatics
	hs := NewHStr("Windows.UI.Composition.Interactions.CompositionConditionalValue")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICompositionConditionalValueStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type CompositionInteractionSourceCollection struct {
	RtClass
	*ICompositionInteractionSourceCollection
}

type InteractionSourceConfiguration struct {
	RtClass
	*IInteractionSourceConfiguration
}

type InteractionTracker struct {
	RtClass
	*IInteractionTracker
}

func NewIInteractionTrackerStatics() *IInteractionTrackerStatics {
	var p *IInteractionTrackerStatics
	hs := NewHStr("Windows.UI.Composition.Interactions.InteractionTracker")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IInteractionTrackerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIInteractionTrackerStatics2() *IInteractionTrackerStatics2 {
	var p *IInteractionTrackerStatics2
	hs := NewHStr("Windows.UI.Composition.Interactions.InteractionTracker")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IInteractionTrackerStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type InteractionTrackerCustomAnimationStateEnteredArgs struct {
	RtClass
	*IInteractionTrackerCustomAnimationStateEnteredArgs
}

type InteractionTrackerIdleStateEnteredArgs struct {
	RtClass
	*IInteractionTrackerIdleStateEnteredArgs
}

type InteractionTrackerInertiaModifier struct {
	RtClass
	*IInteractionTrackerInertiaModifier
}

type InteractionTrackerInertiaMotion struct {
	RtClass
	*IInteractionTrackerInertiaMotion
}

func NewIInteractionTrackerInertiaMotionStatics() *IInteractionTrackerInertiaMotionStatics {
	var p *IInteractionTrackerInertiaMotionStatics
	hs := NewHStr("Windows.UI.Composition.Interactions.InteractionTrackerInertiaMotion")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IInteractionTrackerInertiaMotionStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type InteractionTrackerInertiaNaturalMotion struct {
	RtClass
	*IInteractionTrackerInertiaNaturalMotion
}

func NewIInteractionTrackerInertiaNaturalMotionStatics() *IInteractionTrackerInertiaNaturalMotionStatics {
	var p *IInteractionTrackerInertiaNaturalMotionStatics
	hs := NewHStr("Windows.UI.Composition.Interactions.InteractionTrackerInertiaNaturalMotion")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IInteractionTrackerInertiaNaturalMotionStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type InteractionTrackerInertiaRestingValue struct {
	RtClass
	*IInteractionTrackerInertiaRestingValue
}

func NewIInteractionTrackerInertiaRestingValueStatics() *IInteractionTrackerInertiaRestingValueStatics {
	var p *IInteractionTrackerInertiaRestingValueStatics
	hs := NewHStr("Windows.UI.Composition.Interactions.InteractionTrackerInertiaRestingValue")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IInteractionTrackerInertiaRestingValueStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type InteractionTrackerInertiaStateEnteredArgs struct {
	RtClass
	*IInteractionTrackerInertiaStateEnteredArgs
}

type InteractionTrackerInteractingStateEnteredArgs struct {
	RtClass
	*IInteractionTrackerInteractingStateEnteredArgs
}

type InteractionTrackerRequestIgnoredArgs struct {
	RtClass
	*IInteractionTrackerRequestIgnoredArgs
}

type InteractionTrackerValuesChangedArgs struct {
	RtClass
	*IInteractionTrackerValuesChangedArgs
}

type InteractionTrackerVector2InertiaModifier struct {
	RtClass
	*IInteractionTrackerVector2InertiaModifier
}

type InteractionTrackerVector2InertiaNaturalMotion struct {
	RtClass
	*IInteractionTrackerVector2InertiaNaturalMotion
}

func NewIInteractionTrackerVector2InertiaNaturalMotionStatics() *IInteractionTrackerVector2InertiaNaturalMotionStatics {
	var p *IInteractionTrackerVector2InertiaNaturalMotionStatics
	hs := NewHStr("Windows.UI.Composition.Interactions.InteractionTrackerVector2InertiaNaturalMotion")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IInteractionTrackerVector2InertiaNaturalMotionStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type VisualInteractionSource struct {
	RtClass
	*IVisualInteractionSource
}

func NewIVisualInteractionSourceStatics() *IVisualInteractionSourceStatics {
	var p *IVisualInteractionSourceStatics
	hs := NewHStr("Windows.UI.Composition.Interactions.VisualInteractionSource")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IVisualInteractionSourceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIVisualInteractionSourceStatics2() *IVisualInteractionSourceStatics2 {
	var p *IVisualInteractionSourceStatics2
	hs := NewHStr("Windows.UI.Composition.Interactions.VisualInteractionSource")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IVisualInteractionSourceStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
