package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type AudioEffectType int32

const (
	AudioEffectType_Other                    AudioEffectType = 0
	AudioEffectType_AcousticEchoCancellation AudioEffectType = 1
	AudioEffectType_NoiseSuppression         AudioEffectType = 2
	AudioEffectType_AutomaticGainControl     AudioEffectType = 3
	AudioEffectType_BeamForming              AudioEffectType = 4
	AudioEffectType_ConstantToneRemoval      AudioEffectType = 5
	AudioEffectType_Equalizer                AudioEffectType = 6
	AudioEffectType_LoudnessEqualizer        AudioEffectType = 7
	AudioEffectType_BassBoost                AudioEffectType = 8
	AudioEffectType_VirtualSurround          AudioEffectType = 9
	AudioEffectType_VirtualHeadphones        AudioEffectType = 10
	AudioEffectType_SpeakerFill              AudioEffectType = 11
	AudioEffectType_RoomCorrection           AudioEffectType = 12
	AudioEffectType_BassManagement           AudioEffectType = 13
	AudioEffectType_EnvironmentalEffects     AudioEffectType = 14
	AudioEffectType_SpeakerProtection        AudioEffectType = 15
	AudioEffectType_SpeakerCompensation      AudioEffectType = 16
	AudioEffectType_DynamicRangeCompression  AudioEffectType = 17
	AudioEffectType_FarFieldBeamForming      AudioEffectType = 18
	AudioEffectType_DeepNoiseSuppression     AudioEffectType = 19
)

// enum
type MediaEffectClosedReason int32

const (
	MediaEffectClosedReason_Done                      MediaEffectClosedReason = 0
	MediaEffectClosedReason_UnknownError              MediaEffectClosedReason = 1
	MediaEffectClosedReason_UnsupportedEncodingFormat MediaEffectClosedReason = 2
	MediaEffectClosedReason_EffectCurrentlyUnloaded   MediaEffectClosedReason = 3
)

// enum
type MediaMemoryTypes int32

const (
	MediaMemoryTypes_Gpu       MediaMemoryTypes = 0
	MediaMemoryTypes_Cpu       MediaMemoryTypes = 1
	MediaMemoryTypes_GpuAndCpu MediaMemoryTypes = 2
)

// interfaces

// 8F85C271-038D-4393-8298-540110608EEF
var IID_IAudioCaptureEffectsManager = syscall.GUID{0x8F85C271, 0x038D, 0x4393,
	[8]byte{0x82, 0x98, 0x54, 0x01, 0x10, 0x60, 0x8E, 0xEF}}

type IAudioCaptureEffectsManagerInterface interface {
	win32.IInspectableInterface
	Add_AudioCaptureEffectsChanged(handler TypedEventHandler[*IAudioCaptureEffectsManager, interface{}]) EventRegistrationToken
	Remove_AudioCaptureEffectsChanged(token EventRegistrationToken)
	GetAudioCaptureEffects() *IVectorView[*IAudioEffect]
}

type IAudioCaptureEffectsManagerVtbl struct {
	win32.IInspectableVtbl
	Add_AudioCaptureEffectsChanged    uintptr
	Remove_AudioCaptureEffectsChanged uintptr
	GetAudioCaptureEffects            uintptr
}

type IAudioCaptureEffectsManager struct {
	win32.IInspectable
}

func (this *IAudioCaptureEffectsManager) Vtbl() *IAudioCaptureEffectsManagerVtbl {
	return (*IAudioCaptureEffectsManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioCaptureEffectsManager) Add_AudioCaptureEffectsChanged(handler TypedEventHandler[*IAudioCaptureEffectsManager, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AudioCaptureEffectsChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioCaptureEffectsManager) Remove_AudioCaptureEffectsChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AudioCaptureEffectsChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAudioCaptureEffectsManager) GetAudioCaptureEffects() *IVectorView[*IAudioEffect] {
	var _result *IVectorView[*IAudioEffect]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAudioCaptureEffects, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 34AAFA51-9207-4055-BE93-6E5734A86AE4
var IID_IAudioEffect = syscall.GUID{0x34AAFA51, 0x9207, 0x4055,
	[8]byte{0xBE, 0x93, 0x6E, 0x57, 0x34, 0xA8, 0x6A, 0xE4}}

type IAudioEffectInterface interface {
	win32.IInspectableInterface
	Get_AudioEffectType() AudioEffectType
}

type IAudioEffectVtbl struct {
	win32.IInspectableVtbl
	Get_AudioEffectType uintptr
}

type IAudioEffect struct {
	win32.IInspectable
}

func (this *IAudioEffect) Vtbl() *IAudioEffectVtbl {
	return (*IAudioEffectVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioEffect) Get_AudioEffectType() AudioEffectType {
	var _result AudioEffectType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioEffectType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// E4D7F974-7D80-4F73-9089-E31C9DB9C294
var IID_IAudioEffectDefinition = syscall.GUID{0xE4D7F974, 0x7D80, 0x4F73,
	[8]byte{0x90, 0x89, 0xE3, 0x1C, 0x9D, 0xB9, 0xC2, 0x94}}

type IAudioEffectDefinitionInterface interface {
	win32.IInspectableInterface
	Get_ActivatableClassId() string
	Get_Properties() *IPropertySet
}

type IAudioEffectDefinitionVtbl struct {
	win32.IInspectableVtbl
	Get_ActivatableClassId uintptr
	Get_Properties         uintptr
}

type IAudioEffectDefinition struct {
	win32.IInspectable
}

func (this *IAudioEffectDefinition) Vtbl() *IAudioEffectDefinitionVtbl {
	return (*IAudioEffectDefinitionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioEffectDefinition) Get_ActivatableClassId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ActivatableClassId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAudioEffectDefinition) Get_Properties() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8E1DA646-E705-45ED-8A2B-FC4E4F405A97
var IID_IAudioEffectDefinitionFactory = syscall.GUID{0x8E1DA646, 0xE705, 0x45ED,
	[8]byte{0x8A, 0x2B, 0xFC, 0x4E, 0x4F, 0x40, 0x5A, 0x97}}

type IAudioEffectDefinitionFactoryInterface interface {
	win32.IInspectableInterface
	Create(activatableClassId string) *IAudioEffectDefinition
	CreateWithProperties(activatableClassId string, props *IPropertySet) *IAudioEffectDefinition
}

type IAudioEffectDefinitionFactoryVtbl struct {
	win32.IInspectableVtbl
	Create               uintptr
	CreateWithProperties uintptr
}

type IAudioEffectDefinitionFactory struct {
	win32.IInspectable
}

func (this *IAudioEffectDefinitionFactory) Vtbl() *IAudioEffectDefinitionFactoryVtbl {
	return (*IAudioEffectDefinitionFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioEffectDefinitionFactory) Create(activatableClassId string) *IAudioEffectDefinition {
	var _result *IAudioEffectDefinition
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(activatableClassId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioEffectDefinitionFactory) CreateWithProperties(activatableClassId string, props *IPropertySet) *IAudioEffectDefinition {
	var _result *IAudioEffectDefinition
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithProperties, uintptr(unsafe.Pointer(this)), NewHStr(activatableClassId).Ptr, uintptr(unsafe.Pointer(props)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 66406C04-86FA-47CC-A315-F489D8C3FE10
var IID_IAudioEffectsManagerStatics = syscall.GUID{0x66406C04, 0x86FA, 0x47CC,
	[8]byte{0xA3, 0x15, 0xF4, 0x89, 0xD8, 0xC3, 0xFE, 0x10}}

type IAudioEffectsManagerStaticsInterface interface {
	win32.IInspectableInterface
	CreateAudioRenderEffectsManager(deviceId string, category unsafe.Pointer) *IAudioRenderEffectsManager
	CreateAudioRenderEffectsManagerWithMode(deviceId string, category unsafe.Pointer, mode AudioProcessing) *IAudioRenderEffectsManager
	CreateAudioCaptureEffectsManager(deviceId string, category MediaCategory) *IAudioCaptureEffectsManager
	CreateAudioCaptureEffectsManagerWithMode(deviceId string, category MediaCategory, mode AudioProcessing) *IAudioCaptureEffectsManager
}

type IAudioEffectsManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateAudioRenderEffectsManager          uintptr
	CreateAudioRenderEffectsManagerWithMode  uintptr
	CreateAudioCaptureEffectsManager         uintptr
	CreateAudioCaptureEffectsManagerWithMode uintptr
}

type IAudioEffectsManagerStatics struct {
	win32.IInspectable
}

func (this *IAudioEffectsManagerStatics) Vtbl() *IAudioEffectsManagerStaticsVtbl {
	return (*IAudioEffectsManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioEffectsManagerStatics) CreateAudioRenderEffectsManager(deviceId string, category unsafe.Pointer) *IAudioRenderEffectsManager {
	var _result *IAudioRenderEffectsManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAudioRenderEffectsManager, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(category), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioEffectsManagerStatics) CreateAudioRenderEffectsManagerWithMode(deviceId string, category unsafe.Pointer, mode AudioProcessing) *IAudioRenderEffectsManager {
	var _result *IAudioRenderEffectsManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAudioRenderEffectsManagerWithMode, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(category), uintptr(mode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioEffectsManagerStatics) CreateAudioCaptureEffectsManager(deviceId string, category MediaCategory) *IAudioCaptureEffectsManager {
	var _result *IAudioCaptureEffectsManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAudioCaptureEffectsManager, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(category), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioEffectsManagerStatics) CreateAudioCaptureEffectsManagerWithMode(deviceId string, category MediaCategory, mode AudioProcessing) *IAudioCaptureEffectsManager {
	var _result *IAudioCaptureEffectsManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAudioCaptureEffectsManagerWithMode, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(category), uintptr(mode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4DC98966-8751-42B2-BFCB-39CA7864BD47
var IID_IAudioRenderEffectsManager = syscall.GUID{0x4DC98966, 0x8751, 0x42B2,
	[8]byte{0xBF, 0xCB, 0x39, 0xCA, 0x78, 0x64, 0xBD, 0x47}}

type IAudioRenderEffectsManagerInterface interface {
	win32.IInspectableInterface
	Add_AudioRenderEffectsChanged(handler TypedEventHandler[*IAudioRenderEffectsManager, interface{}]) EventRegistrationToken
	Remove_AudioRenderEffectsChanged(token EventRegistrationToken)
	GetAudioRenderEffects() *IVectorView[*IAudioEffect]
}

type IAudioRenderEffectsManagerVtbl struct {
	win32.IInspectableVtbl
	Add_AudioRenderEffectsChanged    uintptr
	Remove_AudioRenderEffectsChanged uintptr
	GetAudioRenderEffects            uintptr
}

type IAudioRenderEffectsManager struct {
	win32.IInspectable
}

func (this *IAudioRenderEffectsManager) Vtbl() *IAudioRenderEffectsManagerVtbl {
	return (*IAudioRenderEffectsManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioRenderEffectsManager) Add_AudioRenderEffectsChanged(handler TypedEventHandler[*IAudioRenderEffectsManager, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AudioRenderEffectsChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioRenderEffectsManager) Remove_AudioRenderEffectsChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AudioRenderEffectsChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAudioRenderEffectsManager) GetAudioRenderEffects() *IVectorView[*IAudioEffect] {
	var _result *IVectorView[*IAudioEffect]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAudioRenderEffects, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A844CD09-5ECC-44B3-BB4E-1DB07287139C
var IID_IAudioRenderEffectsManager2 = syscall.GUID{0xA844CD09, 0x5ECC, 0x44B3,
	[8]byte{0xBB, 0x4E, 0x1D, 0xB0, 0x72, 0x87, 0x13, 0x9C}}

type IAudioRenderEffectsManager2Interface interface {
	win32.IInspectableInterface
	Get_EffectsProviderThumbnail() *IRandomAccessStreamWithContentType
	Get_EffectsProviderSettingsLabel() string
	ShowSettingsUI()
}

type IAudioRenderEffectsManager2Vtbl struct {
	win32.IInspectableVtbl
	Get_EffectsProviderThumbnail     uintptr
	Get_EffectsProviderSettingsLabel uintptr
	ShowSettingsUI                   uintptr
}

type IAudioRenderEffectsManager2 struct {
	win32.IInspectable
}

func (this *IAudioRenderEffectsManager2) Vtbl() *IAudioRenderEffectsManager2Vtbl {
	return (*IAudioRenderEffectsManager2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioRenderEffectsManager2) Get_EffectsProviderThumbnail() *IRandomAccessStreamWithContentType {
	var _result *IRandomAccessStreamWithContentType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EffectsProviderThumbnail, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioRenderEffectsManager2) Get_EffectsProviderSettingsLabel() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EffectsProviderSettingsLabel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAudioRenderEffectsManager2) ShowSettingsUI() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ShowSettingsUI, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 8C062C53-6BC0-48B8-A99A-4B41550F1359
var IID_IBasicAudioEffect = syscall.GUID{0x8C062C53, 0x6BC0, 0x48B8,
	[8]byte{0xA9, 0x9A, 0x4B, 0x41, 0x55, 0x0F, 0x13, 0x59}}

type IBasicAudioEffectInterface interface {
	win32.IInspectableInterface
	Get_UseInputFrameForOutput() bool
	Get_SupportedEncodingProperties() *IVectorView[*IAudioEncodingProperties]
	SetEncodingProperties(encodingProperties *IAudioEncodingProperties)
	ProcessFrame(context *IProcessAudioFrameContext)
	Close(reason MediaEffectClosedReason)
	DiscardQueuedFrames()
}

type IBasicAudioEffectVtbl struct {
	win32.IInspectableVtbl
	Get_UseInputFrameForOutput      uintptr
	Get_SupportedEncodingProperties uintptr
	SetEncodingProperties           uintptr
	ProcessFrame                    uintptr
	Close                           uintptr
	DiscardQueuedFrames             uintptr
}

type IBasicAudioEffect struct {
	win32.IInspectable
}

func (this *IBasicAudioEffect) Vtbl() *IBasicAudioEffectVtbl {
	return (*IBasicAudioEffectVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBasicAudioEffect) Get_UseInputFrameForOutput() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UseInputFrameForOutput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBasicAudioEffect) Get_SupportedEncodingProperties() *IVectorView[*IAudioEncodingProperties] {
	var _result *IVectorView[*IAudioEncodingProperties]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedEncodingProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBasicAudioEffect) SetEncodingProperties(encodingProperties *IAudioEncodingProperties) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetEncodingProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(encodingProperties)))
	_ = _hr
}

func (this *IBasicAudioEffect) ProcessFrame(context *IProcessAudioFrameContext) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ProcessFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	_ = _hr
}

func (this *IBasicAudioEffect) Close(reason MediaEffectClosedReason) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Close, uintptr(unsafe.Pointer(this)), uintptr(reason))
	_ = _hr
}

func (this *IBasicAudioEffect) DiscardQueuedFrames() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DiscardQueuedFrames, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 8262C7EF-B360-40BE-949B-2FF42FF35693
var IID_IBasicVideoEffect = syscall.GUID{0x8262C7EF, 0xB360, 0x40BE,
	[8]byte{0x94, 0x9B, 0x2F, 0xF4, 0x2F, 0xF3, 0x56, 0x93}}

type IBasicVideoEffectInterface interface {
	win32.IInspectableInterface
	Get_IsReadOnly() bool
	Get_SupportedMemoryTypes() MediaMemoryTypes
	Get_TimeIndependent() bool
	Get_SupportedEncodingProperties() *IVectorView[*IVideoEncodingProperties]
	SetEncodingProperties(encodingProperties *IVideoEncodingProperties, device unsafe.Pointer)
	ProcessFrame(context *IProcessVideoFrameContext)
	Close(reason MediaEffectClosedReason)
	DiscardQueuedFrames()
}

type IBasicVideoEffectVtbl struct {
	win32.IInspectableVtbl
	Get_IsReadOnly                  uintptr
	Get_SupportedMemoryTypes        uintptr
	Get_TimeIndependent             uintptr
	Get_SupportedEncodingProperties uintptr
	SetEncodingProperties           uintptr
	ProcessFrame                    uintptr
	Close                           uintptr
	DiscardQueuedFrames             uintptr
}

type IBasicVideoEffect struct {
	win32.IInspectable
}

func (this *IBasicVideoEffect) Vtbl() *IBasicVideoEffectVtbl {
	return (*IBasicVideoEffectVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBasicVideoEffect) Get_IsReadOnly() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsReadOnly, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBasicVideoEffect) Get_SupportedMemoryTypes() MediaMemoryTypes {
	var _result MediaMemoryTypes
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedMemoryTypes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBasicVideoEffect) Get_TimeIndependent() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TimeIndependent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBasicVideoEffect) Get_SupportedEncodingProperties() *IVectorView[*IVideoEncodingProperties] {
	var _result *IVectorView[*IVideoEncodingProperties]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedEncodingProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBasicVideoEffect) SetEncodingProperties(encodingProperties *IVideoEncodingProperties, device unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetEncodingProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(encodingProperties)), uintptr(device))
	_ = _hr
}

func (this *IBasicVideoEffect) ProcessFrame(context *IProcessVideoFrameContext) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ProcessFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	_ = _hr
}

func (this *IBasicVideoEffect) Close(reason MediaEffectClosedReason) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Close, uintptr(unsafe.Pointer(this)), uintptr(reason))
	_ = _hr
}

func (this *IBasicVideoEffect) DiscardQueuedFrames() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DiscardQueuedFrames, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 6C30024B-F514-4278-A5F7-B9188049D110
var IID_ICompositeVideoFrameContext = syscall.GUID{0x6C30024B, 0xF514, 0x4278,
	[8]byte{0xA5, 0xF7, 0xB9, 0x18, 0x80, 0x49, 0xD1, 0x10}}

type ICompositeVideoFrameContextInterface interface {
	win32.IInspectableInterface
	Get_SurfacesToOverlay() *IVectorView[unsafe.Pointer]
	Get_BackgroundFrame() *IVideoFrame
	Get_OutputFrame() *IVideoFrame
	GetOverlayForSurface(surfaceToOverlay unsafe.Pointer) *IMediaOverlay
}

type ICompositeVideoFrameContextVtbl struct {
	win32.IInspectableVtbl
	Get_SurfacesToOverlay uintptr
	Get_BackgroundFrame   uintptr
	Get_OutputFrame       uintptr
	GetOverlayForSurface  uintptr
}

type ICompositeVideoFrameContext struct {
	win32.IInspectable
}

func (this *ICompositeVideoFrameContext) Vtbl() *ICompositeVideoFrameContextVtbl {
	return (*ICompositeVideoFrameContextVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICompositeVideoFrameContext) Get_SurfacesToOverlay() *IVectorView[unsafe.Pointer] {
	var _result *IVectorView[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SurfacesToOverlay, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositeVideoFrameContext) Get_BackgroundFrame() *IVideoFrame {
	var _result *IVideoFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BackgroundFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositeVideoFrameContext) Get_OutputFrame() *IVideoFrame {
	var _result *IVideoFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutputFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICompositeVideoFrameContext) GetOverlayForSurface(surfaceToOverlay unsafe.Pointer) *IMediaOverlay {
	var _result *IMediaOverlay
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetOverlayForSurface, uintptr(unsafe.Pointer(this)), uintptr(surfaceToOverlay), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4CD92946-1222-4A27-A586-FB3E20273255
var IID_IProcessAudioFrameContext = syscall.GUID{0x4CD92946, 0x1222, 0x4A27,
	[8]byte{0xA5, 0x86, 0xFB, 0x3E, 0x20, 0x27, 0x32, 0x55}}

type IProcessAudioFrameContextInterface interface {
	win32.IInspectableInterface
	Get_InputFrame() *IAudioFrame
	Get_OutputFrame() *IAudioFrame
}

type IProcessAudioFrameContextVtbl struct {
	win32.IInspectableVtbl
	Get_InputFrame  uintptr
	Get_OutputFrame uintptr
}

type IProcessAudioFrameContext struct {
	win32.IInspectable
}

func (this *IProcessAudioFrameContext) Vtbl() *IProcessAudioFrameContextVtbl {
	return (*IProcessAudioFrameContextVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProcessAudioFrameContext) Get_InputFrame() *IAudioFrame {
	var _result *IAudioFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InputFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProcessAudioFrameContext) Get_OutputFrame() *IAudioFrame {
	var _result *IAudioFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutputFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 276F0E2B-6461-401E-BA78-0FDAD6114EEC
var IID_IProcessVideoFrameContext = syscall.GUID{0x276F0E2B, 0x6461, 0x401E,
	[8]byte{0xBA, 0x78, 0x0F, 0xDA, 0xD6, 0x11, 0x4E, 0xEC}}

type IProcessVideoFrameContextInterface interface {
	win32.IInspectableInterface
	Get_InputFrame() *IVideoFrame
	Get_OutputFrame() *IVideoFrame
}

type IProcessVideoFrameContextVtbl struct {
	win32.IInspectableVtbl
	Get_InputFrame  uintptr
	Get_OutputFrame uintptr
}

type IProcessVideoFrameContext struct {
	win32.IInspectable
}

func (this *IProcessVideoFrameContext) Vtbl() *IProcessVideoFrameContextVtbl {
	return (*IProcessVideoFrameContextVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProcessVideoFrameContext) Get_InputFrame() *IVideoFrame {
	var _result *IVideoFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InputFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProcessVideoFrameContext) Get_OutputFrame() *IVideoFrame {
	var _result *IVideoFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutputFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 35053CD0-176C-4763-82C4-1B02DBE31737
var IID_ISlowMotionEffectDefinition = syscall.GUID{0x35053CD0, 0x176C, 0x4763,
	[8]byte{0x82, 0xC4, 0x1B, 0x02, 0xDB, 0xE3, 0x17, 0x37}}

type ISlowMotionEffectDefinitionInterface interface {
	win32.IInspectableInterface
	Get_TimeStretchRate() float64
	Put_TimeStretchRate(value float64)
}

type ISlowMotionEffectDefinitionVtbl struct {
	win32.IInspectableVtbl
	Get_TimeStretchRate uintptr
	Put_TimeStretchRate uintptr
}

type ISlowMotionEffectDefinition struct {
	win32.IInspectable
}

func (this *ISlowMotionEffectDefinition) Vtbl() *ISlowMotionEffectDefinitionVtbl {
	return (*ISlowMotionEffectDefinitionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISlowMotionEffectDefinition) Get_TimeStretchRate() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TimeStretchRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISlowMotionEffectDefinition) Put_TimeStretchRate(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TimeStretchRate, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 8510B43E-420C-420F-96C7-7C98BBA1FC55
var IID_IVideoCompositor = syscall.GUID{0x8510B43E, 0x420C, 0x420F,
	[8]byte{0x96, 0xC7, 0x7C, 0x98, 0xBB, 0xA1, 0xFC, 0x55}}

type IVideoCompositorInterface interface {
	win32.IInspectableInterface
	Get_TimeIndependent() bool
	SetEncodingProperties(backgroundProperties *IVideoEncodingProperties, device unsafe.Pointer)
	CompositeFrame(context *ICompositeVideoFrameContext)
	Close(reason MediaEffectClosedReason)
	DiscardQueuedFrames()
}

type IVideoCompositorVtbl struct {
	win32.IInspectableVtbl
	Get_TimeIndependent   uintptr
	SetEncodingProperties uintptr
	CompositeFrame        uintptr
	Close                 uintptr
	DiscardQueuedFrames   uintptr
}

type IVideoCompositor struct {
	win32.IInspectable
}

func (this *IVideoCompositor) Vtbl() *IVideoCompositorVtbl {
	return (*IVideoCompositorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoCompositor) Get_TimeIndependent() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TimeIndependent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoCompositor) SetEncodingProperties(backgroundProperties *IVideoEncodingProperties, device unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetEncodingProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(backgroundProperties)), uintptr(device))
	_ = _hr
}

func (this *IVideoCompositor) CompositeFrame(context *ICompositeVideoFrameContext) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CompositeFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	_ = _hr
}

func (this *IVideoCompositor) Close(reason MediaEffectClosedReason) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Close, uintptr(unsafe.Pointer(this)), uintptr(reason))
	_ = _hr
}

func (this *IVideoCompositor) DiscardQueuedFrames() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DiscardQueuedFrames, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 7946B8D0-2010-4AE3-9AB2-2CEF42EDD4D2
var IID_IVideoCompositorDefinition = syscall.GUID{0x7946B8D0, 0x2010, 0x4AE3,
	[8]byte{0x9A, 0xB2, 0x2C, 0xEF, 0x42, 0xED, 0xD4, 0xD2}}

type IVideoCompositorDefinitionInterface interface {
	win32.IInspectableInterface
	Get_ActivatableClassId() string
	Get_Properties() *IPropertySet
}

type IVideoCompositorDefinitionVtbl struct {
	win32.IInspectableVtbl
	Get_ActivatableClassId uintptr
	Get_Properties         uintptr
}

type IVideoCompositorDefinition struct {
	win32.IInspectable
}

func (this *IVideoCompositorDefinition) Vtbl() *IVideoCompositorDefinitionVtbl {
	return (*IVideoCompositorDefinitionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoCompositorDefinition) Get_ActivatableClassId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ActivatableClassId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IVideoCompositorDefinition) Get_Properties() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4366FD10-68B8-4D52-89B6-02A968CCA899
var IID_IVideoCompositorDefinitionFactory = syscall.GUID{0x4366FD10, 0x68B8, 0x4D52,
	[8]byte{0x89, 0xB6, 0x02, 0xA9, 0x68, 0xCC, 0xA8, 0x99}}

type IVideoCompositorDefinitionFactoryInterface interface {
	win32.IInspectableInterface
	Create(activatableClassId string) *IVideoCompositorDefinition
	CreateWithProperties(activatableClassId string, props *IPropertySet) *IVideoCompositorDefinition
}

type IVideoCompositorDefinitionFactoryVtbl struct {
	win32.IInspectableVtbl
	Create               uintptr
	CreateWithProperties uintptr
}

type IVideoCompositorDefinitionFactory struct {
	win32.IInspectable
}

func (this *IVideoCompositorDefinitionFactory) Vtbl() *IVideoCompositorDefinitionFactoryVtbl {
	return (*IVideoCompositorDefinitionFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoCompositorDefinitionFactory) Create(activatableClassId string) *IVideoCompositorDefinition {
	var _result *IVideoCompositorDefinition
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(activatableClassId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoCompositorDefinitionFactory) CreateWithProperties(activatableClassId string, props *IPropertySet) *IVideoCompositorDefinition {
	var _result *IVideoCompositorDefinition
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithProperties, uintptr(unsafe.Pointer(this)), NewHStr(activatableClassId).Ptr, uintptr(unsafe.Pointer(props)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 39F38CF0-8D0F-4F3E-84FC-2D46A5297943
var IID_IVideoEffectDefinition = syscall.GUID{0x39F38CF0, 0x8D0F, 0x4F3E,
	[8]byte{0x84, 0xFC, 0x2D, 0x46, 0xA5, 0x29, 0x79, 0x43}}

type IVideoEffectDefinitionInterface interface {
	win32.IInspectableInterface
	Get_ActivatableClassId() string
	Get_Properties() *IPropertySet
}

type IVideoEffectDefinitionVtbl struct {
	win32.IInspectableVtbl
	Get_ActivatableClassId uintptr
	Get_Properties         uintptr
}

type IVideoEffectDefinition struct {
	win32.IInspectable
}

func (this *IVideoEffectDefinition) Vtbl() *IVideoEffectDefinitionVtbl {
	return (*IVideoEffectDefinitionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoEffectDefinition) Get_ActivatableClassId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ActivatableClassId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IVideoEffectDefinition) Get_Properties() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 81439B4E-6E33-428F-9D21-B5AAFEF7617C
var IID_IVideoEffectDefinitionFactory = syscall.GUID{0x81439B4E, 0x6E33, 0x428F,
	[8]byte{0x9D, 0x21, 0xB5, 0xAA, 0xFE, 0xF7, 0x61, 0x7C}}

type IVideoEffectDefinitionFactoryInterface interface {
	win32.IInspectableInterface
	Create(activatableClassId string) *IVideoEffectDefinition
	CreateWithProperties(activatableClassId string, props *IPropertySet) *IVideoEffectDefinition
}

type IVideoEffectDefinitionFactoryVtbl struct {
	win32.IInspectableVtbl
	Create               uintptr
	CreateWithProperties uintptr
}

type IVideoEffectDefinitionFactory struct {
	win32.IInspectable
}

func (this *IVideoEffectDefinitionFactory) Vtbl() *IVideoEffectDefinitionFactoryVtbl {
	return (*IVideoEffectDefinitionFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoEffectDefinitionFactory) Create(activatableClassId string) *IVideoEffectDefinition {
	var _result *IVideoEffectDefinition
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(activatableClassId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoEffectDefinitionFactory) CreateWithProperties(activatableClassId string, props *IPropertySet) *IVideoEffectDefinition {
	var _result *IVideoEffectDefinition
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithProperties, uintptr(unsafe.Pointer(this)), NewHStr(activatableClassId).Ptr, uintptr(unsafe.Pointer(props)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9664BB6A-1EA6-4AA6-8074-ABE8851ECAE2
var IID_IVideoTransformEffectDefinition = syscall.GUID{0x9664BB6A, 0x1EA6, 0x4AA6,
	[8]byte{0x80, 0x74, 0xAB, 0xE8, 0x85, 0x1E, 0xCA, 0xE2}}

type IVideoTransformEffectDefinitionInterface interface {
	win32.IInspectableInterface
	Get_PaddingColor() unsafe.Pointer
	Put_PaddingColor(value unsafe.Pointer)
	Get_OutputSize() Size
	Put_OutputSize(value Size)
	Get_CropRectangle() Rect
	Put_CropRectangle(value Rect)
	Get_Rotation() MediaRotation
	Put_Rotation(value MediaRotation)
	Get_Mirror() MediaMirroringOptions
	Put_Mirror(value MediaMirroringOptions)
	Put_ProcessingAlgorithm(value MediaVideoProcessingAlgorithm)
	Get_ProcessingAlgorithm() MediaVideoProcessingAlgorithm
}

type IVideoTransformEffectDefinitionVtbl struct {
	win32.IInspectableVtbl
	Get_PaddingColor        uintptr
	Put_PaddingColor        uintptr
	Get_OutputSize          uintptr
	Put_OutputSize          uintptr
	Get_CropRectangle       uintptr
	Put_CropRectangle       uintptr
	Get_Rotation            uintptr
	Put_Rotation            uintptr
	Get_Mirror              uintptr
	Put_Mirror              uintptr
	Put_ProcessingAlgorithm uintptr
	Get_ProcessingAlgorithm uintptr
}

type IVideoTransformEffectDefinition struct {
	win32.IInspectable
}

func (this *IVideoTransformEffectDefinition) Vtbl() *IVideoTransformEffectDefinitionVtbl {
	return (*IVideoTransformEffectDefinitionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoTransformEffectDefinition) Get_PaddingColor() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PaddingColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoTransformEffectDefinition) Put_PaddingColor(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PaddingColor, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVideoTransformEffectDefinition) Get_OutputSize() Size {
	var _result Size
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutputSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoTransformEffectDefinition) Put_OutputSize(value Size) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_OutputSize, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IVideoTransformEffectDefinition) Get_CropRectangle() Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CropRectangle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoTransformEffectDefinition) Put_CropRectangle(value Rect) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CropRectangle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IVideoTransformEffectDefinition) Get_Rotation() MediaRotation {
	var _result MediaRotation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Rotation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoTransformEffectDefinition) Put_Rotation(value MediaRotation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Rotation, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVideoTransformEffectDefinition) Get_Mirror() MediaMirroringOptions {
	var _result MediaMirroringOptions
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mirror, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoTransformEffectDefinition) Put_Mirror(value MediaMirroringOptions) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Mirror, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVideoTransformEffectDefinition) Put_ProcessingAlgorithm(value MediaVideoProcessingAlgorithm) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ProcessingAlgorithm, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVideoTransformEffectDefinition) Get_ProcessingAlgorithm() MediaVideoProcessingAlgorithm {
	var _result MediaVideoProcessingAlgorithm
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProcessingAlgorithm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F0A8089F-66C8-4694-9FD9-1136ABF7444A
var IID_IVideoTransformEffectDefinition2 = syscall.GUID{0xF0A8089F, 0x66C8, 0x4694,
	[8]byte{0x9F, 0xD9, 0x11, 0x36, 0xAB, 0xF7, 0x44, 0x4A}}

type IVideoTransformEffectDefinition2Interface interface {
	win32.IInspectableInterface
	Get_SphericalProjection() *IVideoTransformSphericalProjection
}

type IVideoTransformEffectDefinition2Vtbl struct {
	win32.IInspectableVtbl
	Get_SphericalProjection uintptr
}

type IVideoTransformEffectDefinition2 struct {
	win32.IInspectable
}

func (this *IVideoTransformEffectDefinition2) Vtbl() *IVideoTransformEffectDefinition2Vtbl {
	return (*IVideoTransformEffectDefinition2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoTransformEffectDefinition2) Get_SphericalProjection() *IVideoTransformSphericalProjection {
	var _result *IVideoTransformSphericalProjection
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SphericalProjection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CF4401F0-9BF2-4C39-9F41-E022514A8468
var IID_IVideoTransformSphericalProjection = syscall.GUID{0xCF4401F0, 0x9BF2, 0x4C39,
	[8]byte{0x9F, 0x41, 0xE0, 0x22, 0x51, 0x4A, 0x84, 0x68}}

type IVideoTransformSphericalProjectionInterface interface {
	win32.IInspectableInterface
	Get_IsEnabled() bool
	Put_IsEnabled(value bool)
	Get_FrameFormat() SphericalVideoFrameFormat
	Put_FrameFormat(value SphericalVideoFrameFormat)
	Get_ProjectionMode() SphericalVideoProjectionMode
	Put_ProjectionMode(value SphericalVideoProjectionMode)
	Get_HorizontalFieldOfViewInDegrees() float64
	Put_HorizontalFieldOfViewInDegrees(value float64)
	Get_ViewOrientation() unsafe.Pointer
	Put_ViewOrientation(value unsafe.Pointer)
}

type IVideoTransformSphericalProjectionVtbl struct {
	win32.IInspectableVtbl
	Get_IsEnabled                      uintptr
	Put_IsEnabled                      uintptr
	Get_FrameFormat                    uintptr
	Put_FrameFormat                    uintptr
	Get_ProjectionMode                 uintptr
	Put_ProjectionMode                 uintptr
	Get_HorizontalFieldOfViewInDegrees uintptr
	Put_HorizontalFieldOfViewInDegrees uintptr
	Get_ViewOrientation                uintptr
	Put_ViewOrientation                uintptr
}

type IVideoTransformSphericalProjection struct {
	win32.IInspectable
}

func (this *IVideoTransformSphericalProjection) Vtbl() *IVideoTransformSphericalProjectionVtbl {
	return (*IVideoTransformSphericalProjectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoTransformSphericalProjection) Get_IsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoTransformSphericalProjection) Put_IsEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IVideoTransformSphericalProjection) Get_FrameFormat() SphericalVideoFrameFormat {
	var _result SphericalVideoFrameFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoTransformSphericalProjection) Put_FrameFormat(value SphericalVideoFrameFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FrameFormat, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVideoTransformSphericalProjection) Get_ProjectionMode() SphericalVideoProjectionMode {
	var _result SphericalVideoProjectionMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProjectionMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoTransformSphericalProjection) Put_ProjectionMode(value SphericalVideoProjectionMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ProjectionMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVideoTransformSphericalProjection) Get_HorizontalFieldOfViewInDegrees() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HorizontalFieldOfViewInDegrees, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoTransformSphericalProjection) Put_HorizontalFieldOfViewInDegrees(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_HorizontalFieldOfViewInDegrees, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVideoTransformSphericalProjection) Get_ViewOrientation() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ViewOrientation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoTransformSphericalProjection) Put_ViewOrientation(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ViewOrientation, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// classes

type AudioCaptureEffectsManager struct {
	RtClass
	*IAudioCaptureEffectsManager
}

type AudioEffect struct {
	RtClass
	*IAudioEffect
}

type CompositeVideoFrameContext struct {
	RtClass
	*ICompositeVideoFrameContext
}

type ProcessVideoFrameContext struct {
	RtClass
	*IProcessVideoFrameContext
}
