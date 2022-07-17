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
type ConditionForceEffectKind int32

const (
	ConditionForceEffectKind_Spring   ConditionForceEffectKind = 0
	ConditionForceEffectKind_Damper   ConditionForceEffectKind = 1
	ConditionForceEffectKind_Inertia  ConditionForceEffectKind = 2
	ConditionForceEffectKind_Friction ConditionForceEffectKind = 3
)

// enum
// flags
type ForceFeedbackEffectAxes uint32

const (
	ForceFeedbackEffectAxes_None ForceFeedbackEffectAxes = 0
	ForceFeedbackEffectAxes_X    ForceFeedbackEffectAxes = 1
	ForceFeedbackEffectAxes_Y    ForceFeedbackEffectAxes = 2
	ForceFeedbackEffectAxes_Z    ForceFeedbackEffectAxes = 4
)

// enum
type ForceFeedbackEffectState int32

const (
	ForceFeedbackEffectState_Stopped ForceFeedbackEffectState = 0
	ForceFeedbackEffectState_Running ForceFeedbackEffectState = 1
	ForceFeedbackEffectState_Paused  ForceFeedbackEffectState = 2
	ForceFeedbackEffectState_Faulted ForceFeedbackEffectState = 3
)

// enum
type ForceFeedbackLoadEffectResult int32

const (
	ForceFeedbackLoadEffectResult_Succeeded          ForceFeedbackLoadEffectResult = 0
	ForceFeedbackLoadEffectResult_EffectStorageFull  ForceFeedbackLoadEffectResult = 1
	ForceFeedbackLoadEffectResult_EffectNotSupported ForceFeedbackLoadEffectResult = 2
)

// enum
type PeriodicForceEffectKind int32

const (
	PeriodicForceEffectKind_SquareWave       PeriodicForceEffectKind = 0
	PeriodicForceEffectKind_SineWave         PeriodicForceEffectKind = 1
	PeriodicForceEffectKind_TriangleWave     PeriodicForceEffectKind = 2
	PeriodicForceEffectKind_SawtoothWaveUp   PeriodicForceEffectKind = 3
	PeriodicForceEffectKind_SawtoothWaveDown PeriodicForceEffectKind = 4
)

// interfaces

// 32D1EA68-3695-4E69-85C0-CD1944189140
var IID_IConditionForceEffect = syscall.GUID{0x32D1EA68, 0x3695, 0x4E69,
	[8]byte{0x85, 0xC0, 0xCD, 0x19, 0x44, 0x18, 0x91, 0x40}}

type IConditionForceEffectInterface interface {
	win32.IInspectableInterface
	Get_Kind() ConditionForceEffectKind
	SetParameters(direction unsafe.Pointer, positiveCoefficient float32, negativeCoefficient float32, maxPositiveMagnitude float32, maxNegativeMagnitude float32, deadZone float32, bias float32)
}

type IConditionForceEffectVtbl struct {
	win32.IInspectableVtbl
	Get_Kind      uintptr
	SetParameters uintptr
}

type IConditionForceEffect struct {
	win32.IInspectable
}

func (this *IConditionForceEffect) Vtbl() *IConditionForceEffectVtbl {
	return (*IConditionForceEffectVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IConditionForceEffect) Get_Kind() ConditionForceEffectKind {
	var _result ConditionForceEffectKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Kind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IConditionForceEffect) SetParameters(direction unsafe.Pointer, positiveCoefficient float32, negativeCoefficient float32, maxPositiveMagnitude float32, maxNegativeMagnitude float32, deadZone float32, bias float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetParameters, uintptr(unsafe.Pointer(this)), uintptr(direction), uintptr(positiveCoefficient), uintptr(negativeCoefficient), uintptr(maxPositiveMagnitude), uintptr(maxNegativeMagnitude), uintptr(deadZone), uintptr(bias))
	_ = _hr
}

// 91A99264-1810-4EB6-A773-BFD3B8CDDBAB
var IID_IConditionForceEffectFactory = syscall.GUID{0x91A99264, 0x1810, 0x4EB6,
	[8]byte{0xA7, 0x73, 0xBF, 0xD3, 0xB8, 0xCD, 0xDB, 0xAB}}

type IConditionForceEffectFactoryInterface interface {
	win32.IInspectableInterface
	CreateInstance(effectKind ConditionForceEffectKind) *IForceFeedbackEffect
}

type IConditionForceEffectFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateInstance uintptr
}

type IConditionForceEffectFactory struct {
	win32.IInspectable
}

func (this *IConditionForceEffectFactory) Vtbl() *IConditionForceEffectFactoryVtbl {
	return (*IConditionForceEffectFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IConditionForceEffectFactory) CreateInstance(effectKind ConditionForceEffectKind) *IForceFeedbackEffect {
	var _result *IForceFeedbackEffect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateInstance, uintptr(unsafe.Pointer(this)), uintptr(effectKind), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9BFA0140-F3C7-415C-B068-0F068734BCE0
var IID_IConstantForceEffect = syscall.GUID{0x9BFA0140, 0xF3C7, 0x415C,
	[8]byte{0xB0, 0x68, 0x0F, 0x06, 0x87, 0x34, 0xBC, 0xE0}}

type IConstantForceEffectInterface interface {
	win32.IInspectableInterface
	SetParameters(vector unsafe.Pointer, duration TimeSpan)
	SetParametersWithEnvelope(vector unsafe.Pointer, attackGain float32, sustainGain float32, releaseGain float32, startDelay TimeSpan, attackDuration TimeSpan, sustainDuration TimeSpan, releaseDuration TimeSpan, repeatCount uint32)
}

type IConstantForceEffectVtbl struct {
	win32.IInspectableVtbl
	SetParameters             uintptr
	SetParametersWithEnvelope uintptr
}

type IConstantForceEffect struct {
	win32.IInspectable
}

func (this *IConstantForceEffect) Vtbl() *IConstantForceEffectVtbl {
	return (*IConstantForceEffectVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IConstantForceEffect) SetParameters(vector unsafe.Pointer, duration TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetParameters, uintptr(unsafe.Pointer(this)), uintptr(vector), *(*uintptr)(unsafe.Pointer(&duration)))
	_ = _hr
}

func (this *IConstantForceEffect) SetParametersWithEnvelope(vector unsafe.Pointer, attackGain float32, sustainGain float32, releaseGain float32, startDelay TimeSpan, attackDuration TimeSpan, sustainDuration TimeSpan, releaseDuration TimeSpan, repeatCount uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetParametersWithEnvelope, uintptr(unsafe.Pointer(this)), uintptr(vector), uintptr(attackGain), uintptr(sustainGain), uintptr(releaseGain), *(*uintptr)(unsafe.Pointer(&startDelay)), *(*uintptr)(unsafe.Pointer(&attackDuration)), *(*uintptr)(unsafe.Pointer(&sustainDuration)), *(*uintptr)(unsafe.Pointer(&releaseDuration)), uintptr(repeatCount))
	_ = _hr
}

// A17FBA0C-2AE4-48C2-8063-EABD0777CB89
var IID_IForceFeedbackEffect = syscall.GUID{0xA17FBA0C, 0x2AE4, 0x48C2,
	[8]byte{0x80, 0x63, 0xEA, 0xBD, 0x07, 0x77, 0xCB, 0x89}}

type IForceFeedbackEffectInterface interface {
	win32.IInspectableInterface
	Get_Gain() float64
	Put_Gain(value float64)
	Get_State() ForceFeedbackEffectState
	Start()
	Stop()
}

type IForceFeedbackEffectVtbl struct {
	win32.IInspectableVtbl
	Get_Gain  uintptr
	Put_Gain  uintptr
	Get_State uintptr
	Start     uintptr
	Stop      uintptr
}

type IForceFeedbackEffect struct {
	win32.IInspectable
}

func (this *IForceFeedbackEffect) Vtbl() *IForceFeedbackEffectVtbl {
	return (*IForceFeedbackEffectVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IForceFeedbackEffect) Get_Gain() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Gain, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IForceFeedbackEffect) Put_Gain(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Gain, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IForceFeedbackEffect) Get_State() ForceFeedbackEffectState {
	var _result ForceFeedbackEffectState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IForceFeedbackEffect) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IForceFeedbackEffect) Stop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 8D3D417C-A5EA-4516-8026-2B00F74EF6E5
var IID_IForceFeedbackMotor = syscall.GUID{0x8D3D417C, 0xA5EA, 0x4516,
	[8]byte{0x80, 0x26, 0x2B, 0x00, 0xF7, 0x4E, 0xF6, 0xE5}}

type IForceFeedbackMotorInterface interface {
	win32.IInspectableInterface
	Get_AreEffectsPaused() bool
	Get_MasterGain() float64
	Put_MasterGain(value float64)
	Get_IsEnabled() bool
	Get_SupportedAxes() ForceFeedbackEffectAxes
	LoadEffectAsync(effect *IForceFeedbackEffect) *IAsyncOperation[ForceFeedbackLoadEffectResult]
	PauseAllEffects()
	ResumeAllEffects()
	StopAllEffects()
	TryDisableAsync() *IAsyncOperation[bool]
	TryEnableAsync() *IAsyncOperation[bool]
	TryResetAsync() *IAsyncOperation[bool]
	TryUnloadEffectAsync(effect *IForceFeedbackEffect) *IAsyncOperation[bool]
}

type IForceFeedbackMotorVtbl struct {
	win32.IInspectableVtbl
	Get_AreEffectsPaused uintptr
	Get_MasterGain       uintptr
	Put_MasterGain       uintptr
	Get_IsEnabled        uintptr
	Get_SupportedAxes    uintptr
	LoadEffectAsync      uintptr
	PauseAllEffects      uintptr
	ResumeAllEffects     uintptr
	StopAllEffects       uintptr
	TryDisableAsync      uintptr
	TryEnableAsync       uintptr
	TryResetAsync        uintptr
	TryUnloadEffectAsync uintptr
}

type IForceFeedbackMotor struct {
	win32.IInspectable
}

func (this *IForceFeedbackMotor) Vtbl() *IForceFeedbackMotorVtbl {
	return (*IForceFeedbackMotorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IForceFeedbackMotor) Get_AreEffectsPaused() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AreEffectsPaused, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IForceFeedbackMotor) Get_MasterGain() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MasterGain, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IForceFeedbackMotor) Put_MasterGain(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MasterGain, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IForceFeedbackMotor) Get_IsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IForceFeedbackMotor) Get_SupportedAxes() ForceFeedbackEffectAxes {
	var _result ForceFeedbackEffectAxes
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedAxes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IForceFeedbackMotor) LoadEffectAsync(effect *IForceFeedbackEffect) *IAsyncOperation[ForceFeedbackLoadEffectResult] {
	var _result *IAsyncOperation[ForceFeedbackLoadEffectResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LoadEffectAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(effect)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IForceFeedbackMotor) PauseAllEffects() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PauseAllEffects, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IForceFeedbackMotor) ResumeAllEffects() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ResumeAllEffects, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IForceFeedbackMotor) StopAllEffects() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopAllEffects, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IForceFeedbackMotor) TryDisableAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryDisableAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IForceFeedbackMotor) TryEnableAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryEnableAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IForceFeedbackMotor) TryResetAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryResetAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IForceFeedbackMotor) TryUnloadEffectAsync(effect *IForceFeedbackEffect) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryUnloadEffectAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(effect)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5C5138D7-FC75-4D52-9A0A-EFE4CAB5FE64
var IID_IPeriodicForceEffect = syscall.GUID{0x5C5138D7, 0xFC75, 0x4D52,
	[8]byte{0x9A, 0x0A, 0xEF, 0xE4, 0xCA, 0xB5, 0xFE, 0x64}}

type IPeriodicForceEffectInterface interface {
	win32.IInspectableInterface
	Get_Kind() PeriodicForceEffectKind
	SetParameters(vector unsafe.Pointer, frequency float32, phase float32, bias float32, duration TimeSpan)
	SetParametersWithEnvelope(vector unsafe.Pointer, frequency float32, phase float32, bias float32, attackGain float32, sustainGain float32, releaseGain float32, startDelay TimeSpan, attackDuration TimeSpan, sustainDuration TimeSpan, releaseDuration TimeSpan, repeatCount uint32)
}

type IPeriodicForceEffectVtbl struct {
	win32.IInspectableVtbl
	Get_Kind                  uintptr
	SetParameters             uintptr
	SetParametersWithEnvelope uintptr
}

type IPeriodicForceEffect struct {
	win32.IInspectable
}

func (this *IPeriodicForceEffect) Vtbl() *IPeriodicForceEffectVtbl {
	return (*IPeriodicForceEffectVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPeriodicForceEffect) Get_Kind() PeriodicForceEffectKind {
	var _result PeriodicForceEffectKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Kind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPeriodicForceEffect) SetParameters(vector unsafe.Pointer, frequency float32, phase float32, bias float32, duration TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetParameters, uintptr(unsafe.Pointer(this)), uintptr(vector), uintptr(frequency), uintptr(phase), uintptr(bias), *(*uintptr)(unsafe.Pointer(&duration)))
	_ = _hr
}

func (this *IPeriodicForceEffect) SetParametersWithEnvelope(vector unsafe.Pointer, frequency float32, phase float32, bias float32, attackGain float32, sustainGain float32, releaseGain float32, startDelay TimeSpan, attackDuration TimeSpan, sustainDuration TimeSpan, releaseDuration TimeSpan, repeatCount uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetParametersWithEnvelope, uintptr(unsafe.Pointer(this)), uintptr(vector), uintptr(frequency), uintptr(phase), uintptr(bias), uintptr(attackGain), uintptr(sustainGain), uintptr(releaseGain), *(*uintptr)(unsafe.Pointer(&startDelay)), *(*uintptr)(unsafe.Pointer(&attackDuration)), *(*uintptr)(unsafe.Pointer(&sustainDuration)), *(*uintptr)(unsafe.Pointer(&releaseDuration)), uintptr(repeatCount))
	_ = _hr
}

// 6F62EB1A-9851-477B-B318-35ECAA15070F
var IID_IPeriodicForceEffectFactory = syscall.GUID{0x6F62EB1A, 0x9851, 0x477B,
	[8]byte{0xB3, 0x18, 0x35, 0xEC, 0xAA, 0x15, 0x07, 0x0F}}

type IPeriodicForceEffectFactoryInterface interface {
	win32.IInspectableInterface
	CreateInstance(effectKind PeriodicForceEffectKind) *IForceFeedbackEffect
}

type IPeriodicForceEffectFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateInstance uintptr
}

type IPeriodicForceEffectFactory struct {
	win32.IInspectable
}

func (this *IPeriodicForceEffectFactory) Vtbl() *IPeriodicForceEffectFactoryVtbl {
	return (*IPeriodicForceEffectFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPeriodicForceEffectFactory) CreateInstance(effectKind PeriodicForceEffectKind) *IForceFeedbackEffect {
	var _result *IForceFeedbackEffect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateInstance, uintptr(unsafe.Pointer(this)), uintptr(effectKind), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F1F81259-1CA6-4080-B56D-B43F3354D052
var IID_IRampForceEffect = syscall.GUID{0xF1F81259, 0x1CA6, 0x4080,
	[8]byte{0xB5, 0x6D, 0xB4, 0x3F, 0x33, 0x54, 0xD0, 0x52}}

type IRampForceEffectInterface interface {
	win32.IInspectableInterface
	SetParameters(startVector unsafe.Pointer, endVector unsafe.Pointer, duration TimeSpan)
	SetParametersWithEnvelope(startVector unsafe.Pointer, endVector unsafe.Pointer, attackGain float32, sustainGain float32, releaseGain float32, startDelay TimeSpan, attackDuration TimeSpan, sustainDuration TimeSpan, releaseDuration TimeSpan, repeatCount uint32)
}

type IRampForceEffectVtbl struct {
	win32.IInspectableVtbl
	SetParameters             uintptr
	SetParametersWithEnvelope uintptr
}

type IRampForceEffect struct {
	win32.IInspectable
}

func (this *IRampForceEffect) Vtbl() *IRampForceEffectVtbl {
	return (*IRampForceEffectVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRampForceEffect) SetParameters(startVector unsafe.Pointer, endVector unsafe.Pointer, duration TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetParameters, uintptr(unsafe.Pointer(this)), uintptr(startVector), uintptr(endVector), *(*uintptr)(unsafe.Pointer(&duration)))
	_ = _hr
}

func (this *IRampForceEffect) SetParametersWithEnvelope(startVector unsafe.Pointer, endVector unsafe.Pointer, attackGain float32, sustainGain float32, releaseGain float32, startDelay TimeSpan, attackDuration TimeSpan, sustainDuration TimeSpan, releaseDuration TimeSpan, repeatCount uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetParametersWithEnvelope, uintptr(unsafe.Pointer(this)), uintptr(startVector), uintptr(endVector), uintptr(attackGain), uintptr(sustainGain), uintptr(releaseGain), *(*uintptr)(unsafe.Pointer(&startDelay)), *(*uintptr)(unsafe.Pointer(&attackDuration)), *(*uintptr)(unsafe.Pointer(&sustainDuration)), *(*uintptr)(unsafe.Pointer(&releaseDuration)), uintptr(repeatCount))
	_ = _hr
}

// classes

type ConditionForceEffect struct {
	RtClass
	*IForceFeedbackEffect
}

func NewConditionForceEffect_CreateInstance(effectKind ConditionForceEffectKind) *ConditionForceEffect {
	hs := NewHStr("Windows.Gaming.Input.ForceFeedback.ConditionForceEffect")
	var pFac *IConditionForceEffectFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IConditionForceEffectFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IForceFeedbackEffect
	p = pFac.CreateInstance(effectKind)
	result := &ConditionForceEffect{
		RtClass:              RtClass{PInspect: &p.IInspectable},
		IForceFeedbackEffect: p,
	}
	com.AddToScope(result)
	return result
}

type ConstantForceEffect struct {
	RtClass
	*IForceFeedbackEffect
}

func NewConstantForceEffect() *ConstantForceEffect {
	hs := NewHStr("Windows.Gaming.Input.ForceFeedback.ConstantForceEffect")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &ConstantForceEffect{
		RtClass:              RtClass{PInspect: p},
		IForceFeedbackEffect: (*IForceFeedbackEffect)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type ForceFeedbackMotor struct {
	RtClass
	*IForceFeedbackMotor
}

type PeriodicForceEffect struct {
	RtClass
	*IForceFeedbackEffect
}

func NewPeriodicForceEffect_CreateInstance(effectKind PeriodicForceEffectKind) *PeriodicForceEffect {
	hs := NewHStr("Windows.Gaming.Input.ForceFeedback.PeriodicForceEffect")
	var pFac *IPeriodicForceEffectFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPeriodicForceEffectFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IForceFeedbackEffect
	p = pFac.CreateInstance(effectKind)
	result := &PeriodicForceEffect{
		RtClass:              RtClass{PInspect: &p.IInspectable},
		IForceFeedbackEffect: p,
	}
	com.AddToScope(result)
	return result
}

type RampForceEffect struct {
	RtClass
	*IForceFeedbackEffect
}

func NewRampForceEffect() *RampForceEffect {
	hs := NewHStr("Windows.Gaming.Input.ForceFeedback.RampForceEffect")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &RampForceEffect{
		RtClass:              RtClass{PInspect: p},
		IForceFeedbackEffect: (*IForceFeedbackEffect)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}
