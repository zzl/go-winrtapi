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
type SpeechAppendedSilence int32

const (
	SpeechAppendedSilence_Default SpeechAppendedSilence = 0
	SpeechAppendedSilence_Min     SpeechAppendedSilence = 1
)

// enum
type SpeechPunctuationSilence int32

const (
	SpeechPunctuationSilence_Default SpeechPunctuationSilence = 0
	SpeechPunctuationSilence_Min     SpeechPunctuationSilence = 1
)

// enum
type VoiceGender int32

const (
	VoiceGender_Male   VoiceGender = 0
	VoiceGender_Female VoiceGender = 1
)

// interfaces

// 7D526ECC-7533-4C3F-85BE-888C2BAEEBDC
var IID_IInstalledVoicesStatic = syscall.GUID{0x7D526ECC, 0x7533, 0x4C3F,
	[8]byte{0x85, 0xBE, 0x88, 0x8C, 0x2B, 0xAE, 0xEB, 0xDC}}

type IInstalledVoicesStaticInterface interface {
	win32.IInspectableInterface
	Get_AllVoices() *IVectorView[*IVoiceInformation]
	Get_DefaultVoice() *IVoiceInformation
}

type IInstalledVoicesStaticVtbl struct {
	win32.IInspectableVtbl
	Get_AllVoices    uintptr
	Get_DefaultVoice uintptr
}

type IInstalledVoicesStatic struct {
	win32.IInspectable
}

func (this *IInstalledVoicesStatic) Vtbl() *IInstalledVoicesStaticVtbl {
	return (*IInstalledVoicesStaticVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInstalledVoicesStatic) Get_AllVoices() *IVectorView[*IVoiceInformation] {
	var _result *IVectorView[*IVoiceInformation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllVoices, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInstalledVoicesStatic) Get_DefaultVoice() *IVoiceInformation {
	var _result *IVoiceInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DefaultVoice, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 64255F2E-358D-4058-BE9A-FD3FCB423530
var IID_IInstalledVoicesStatic2 = syscall.GUID{0x64255F2E, 0x358D, 0x4058,
	[8]byte{0xBE, 0x9A, 0xFD, 0x3F, 0xCB, 0x42, 0x35, 0x30}}

type IInstalledVoicesStatic2Interface interface {
	win32.IInspectableInterface
	TrySetDefaultVoiceAsync(voice *IVoiceInformation) *IAsyncOperation[bool]
}

type IInstalledVoicesStatic2Vtbl struct {
	win32.IInspectableVtbl
	TrySetDefaultVoiceAsync uintptr
}

type IInstalledVoicesStatic2 struct {
	win32.IInspectable
}

func (this *IInstalledVoicesStatic2) Vtbl() *IInstalledVoicesStatic2Vtbl {
	return (*IInstalledVoicesStatic2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInstalledVoicesStatic2) TrySetDefaultVoiceAsync(voice *IVoiceInformation) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TrySetDefaultVoiceAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(voice)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 83E46E93-244C-4622-BA0B-6229C4D0D65D
var IID_ISpeechSynthesisStream = syscall.GUID{0x83E46E93, 0x244C, 0x4622,
	[8]byte{0xBA, 0x0B, 0x62, 0x29, 0xC4, 0xD0, 0xD6, 0x5D}}

type ISpeechSynthesisStreamInterface interface {
	win32.IInspectableInterface
	Get_Markers() *IVectorView[*IMediaMarker]
}

type ISpeechSynthesisStreamVtbl struct {
	win32.IInspectableVtbl
	Get_Markers uintptr
}

type ISpeechSynthesisStream struct {
	win32.IInspectable
}

func (this *ISpeechSynthesisStream) Vtbl() *ISpeechSynthesisStreamVtbl {
	return (*ISpeechSynthesisStreamVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechSynthesisStream) Get_Markers() *IVectorView[*IMediaMarker] {
	var _result *IVectorView[*IMediaMarker]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Markers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CE9F7C76-97F4-4CED-AD68-D51C458E45C6
var IID_ISpeechSynthesizer = syscall.GUID{0xCE9F7C76, 0x97F4, 0x4CED,
	[8]byte{0xAD, 0x68, 0xD5, 0x1C, 0x45, 0x8E, 0x45, 0xC6}}

type ISpeechSynthesizerInterface interface {
	win32.IInspectableInterface
	SynthesizeTextToStreamAsync(text string) *IAsyncOperation[*ISpeechSynthesisStream]
	SynthesizeSsmlToStreamAsync(Ssml string) *IAsyncOperation[*ISpeechSynthesisStream]
	Put_Voice(value *IVoiceInformation)
	Get_Voice() *IVoiceInformation
}

type ISpeechSynthesizerVtbl struct {
	win32.IInspectableVtbl
	SynthesizeTextToStreamAsync uintptr
	SynthesizeSsmlToStreamAsync uintptr
	Put_Voice                   uintptr
	Get_Voice                   uintptr
}

type ISpeechSynthesizer struct {
	win32.IInspectable
}

func (this *ISpeechSynthesizer) Vtbl() *ISpeechSynthesizerVtbl {
	return (*ISpeechSynthesizerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechSynthesizer) SynthesizeTextToStreamAsync(text string) *IAsyncOperation[*ISpeechSynthesisStream] {
	var _result *IAsyncOperation[*ISpeechSynthesisStream]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SynthesizeTextToStreamAsync, uintptr(unsafe.Pointer(this)), NewHStr(text).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpeechSynthesizer) SynthesizeSsmlToStreamAsync(Ssml string) *IAsyncOperation[*ISpeechSynthesisStream] {
	var _result *IAsyncOperation[*ISpeechSynthesisStream]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SynthesizeSsmlToStreamAsync, uintptr(unsafe.Pointer(this)), NewHStr(Ssml).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpeechSynthesizer) Put_Voice(value *IVoiceInformation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Voice, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ISpeechSynthesizer) Get_Voice() *IVoiceInformation {
	var _result *IVoiceInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Voice, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A7C5ECB2-4339-4D6A-BBF8-C7A4F1544C2E
var IID_ISpeechSynthesizer2 = syscall.GUID{0xA7C5ECB2, 0x4339, 0x4D6A,
	[8]byte{0xBB, 0xF8, 0xC7, 0xA4, 0xF1, 0x54, 0x4C, 0x2E}}

type ISpeechSynthesizer2Interface interface {
	win32.IInspectableInterface
	Get_Options() *ISpeechSynthesizerOptions
}

type ISpeechSynthesizer2Vtbl struct {
	win32.IInspectableVtbl
	Get_Options uintptr
}

type ISpeechSynthesizer2 struct {
	win32.IInspectable
}

func (this *ISpeechSynthesizer2) Vtbl() *ISpeechSynthesizer2Vtbl {
	return (*ISpeechSynthesizer2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechSynthesizer2) Get_Options() *ISpeechSynthesizerOptions {
	var _result *ISpeechSynthesizerOptions
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Options, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A0E23871-CC3D-43C9-91B1-EE185324D83D
var IID_ISpeechSynthesizerOptions = syscall.GUID{0xA0E23871, 0xCC3D, 0x43C9,
	[8]byte{0x91, 0xB1, 0xEE, 0x18, 0x53, 0x24, 0xD8, 0x3D}}

type ISpeechSynthesizerOptionsInterface interface {
	win32.IInspectableInterface
	Get_IncludeWordBoundaryMetadata() bool
	Put_IncludeWordBoundaryMetadata(value bool)
	Get_IncludeSentenceBoundaryMetadata() bool
	Put_IncludeSentenceBoundaryMetadata(value bool)
}

type ISpeechSynthesizerOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_IncludeWordBoundaryMetadata     uintptr
	Put_IncludeWordBoundaryMetadata     uintptr
	Get_IncludeSentenceBoundaryMetadata uintptr
	Put_IncludeSentenceBoundaryMetadata uintptr
}

type ISpeechSynthesizerOptions struct {
	win32.IInspectable
}

func (this *ISpeechSynthesizerOptions) Vtbl() *ISpeechSynthesizerOptionsVtbl {
	return (*ISpeechSynthesizerOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechSynthesizerOptions) Get_IncludeWordBoundaryMetadata() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IncludeWordBoundaryMetadata, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpeechSynthesizerOptions) Put_IncludeWordBoundaryMetadata(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IncludeWordBoundaryMetadata, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ISpeechSynthesizerOptions) Get_IncludeSentenceBoundaryMetadata() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IncludeSentenceBoundaryMetadata, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpeechSynthesizerOptions) Put_IncludeSentenceBoundaryMetadata(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IncludeSentenceBoundaryMetadata, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 1CBEF60E-119C-4BED-B118-D250C3A25793
var IID_ISpeechSynthesizerOptions2 = syscall.GUID{0x1CBEF60E, 0x119C, 0x4BED,
	[8]byte{0xB1, 0x18, 0xD2, 0x50, 0xC3, 0xA2, 0x57, 0x93}}

type ISpeechSynthesizerOptions2Interface interface {
	win32.IInspectableInterface
	Get_AudioVolume() float64
	Put_AudioVolume(value float64)
	Get_SpeakingRate() float64
	Put_SpeakingRate(value float64)
	Get_AudioPitch() float64
	Put_AudioPitch(value float64)
}

type ISpeechSynthesizerOptions2Vtbl struct {
	win32.IInspectableVtbl
	Get_AudioVolume  uintptr
	Put_AudioVolume  uintptr
	Get_SpeakingRate uintptr
	Put_SpeakingRate uintptr
	Get_AudioPitch   uintptr
	Put_AudioPitch   uintptr
}

type ISpeechSynthesizerOptions2 struct {
	win32.IInspectable
}

func (this *ISpeechSynthesizerOptions2) Vtbl() *ISpeechSynthesizerOptions2Vtbl {
	return (*ISpeechSynthesizerOptions2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechSynthesizerOptions2) Get_AudioVolume() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioVolume, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpeechSynthesizerOptions2) Put_AudioVolume(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AudioVolume, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISpeechSynthesizerOptions2) Get_SpeakingRate() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SpeakingRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpeechSynthesizerOptions2) Put_SpeakingRate(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SpeakingRate, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISpeechSynthesizerOptions2) Get_AudioPitch() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioPitch, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpeechSynthesizerOptions2) Put_AudioPitch(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AudioPitch, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 401ED877-902C-4814-A582-A5D0C0769FA8
var IID_ISpeechSynthesizerOptions3 = syscall.GUID{0x401ED877, 0x902C, 0x4814,
	[8]byte{0xA5, 0x82, 0xA5, 0xD0, 0xC0, 0x76, 0x9F, 0xA8}}

type ISpeechSynthesizerOptions3Interface interface {
	win32.IInspectableInterface
	Get_AppendedSilence() SpeechAppendedSilence
	Put_AppendedSilence(value SpeechAppendedSilence)
	Get_PunctuationSilence() SpeechPunctuationSilence
	Put_PunctuationSilence(value SpeechPunctuationSilence)
}

type ISpeechSynthesizerOptions3Vtbl struct {
	win32.IInspectableVtbl
	Get_AppendedSilence    uintptr
	Put_AppendedSilence    uintptr
	Get_PunctuationSilence uintptr
	Put_PunctuationSilence uintptr
}

type ISpeechSynthesizerOptions3 struct {
	win32.IInspectable
}

func (this *ISpeechSynthesizerOptions3) Vtbl() *ISpeechSynthesizerOptions3Vtbl {
	return (*ISpeechSynthesizerOptions3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechSynthesizerOptions3) Get_AppendedSilence() SpeechAppendedSilence {
	var _result SpeechAppendedSilence
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppendedSilence, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpeechSynthesizerOptions3) Put_AppendedSilence(value SpeechAppendedSilence) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AppendedSilence, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISpeechSynthesizerOptions3) Get_PunctuationSilence() SpeechPunctuationSilence {
	var _result SpeechPunctuationSilence
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PunctuationSilence, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpeechSynthesizerOptions3) Put_PunctuationSilence(value SpeechPunctuationSilence) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PunctuationSilence, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// B127D6A4-1291-4604-AA9C-83134083352C
var IID_IVoiceInformation = syscall.GUID{0xB127D6A4, 0x1291, 0x4604,
	[8]byte{0xAA, 0x9C, 0x83, 0x13, 0x40, 0x83, 0x35, 0x2C}}

type IVoiceInformationInterface interface {
	win32.IInspectableInterface
	Get_DisplayName() string
	Get_Id() string
	Get_Language() string
	Get_Description() string
	Get_Gender() VoiceGender
}

type IVoiceInformationVtbl struct {
	win32.IInspectableVtbl
	Get_DisplayName uintptr
	Get_Id          uintptr
	Get_Language    uintptr
	Get_Description uintptr
	Get_Gender      uintptr
}

type IVoiceInformation struct {
	win32.IInspectable
}

func (this *IVoiceInformation) Vtbl() *IVoiceInformationVtbl {
	return (*IVoiceInformationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVoiceInformation) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IVoiceInformation) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IVoiceInformation) Get_Language() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Language, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IVoiceInformation) Get_Description() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Description, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IVoiceInformation) Get_Gender() VoiceGender {
	var _result VoiceGender
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Gender, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type SpeechSynthesisStream struct {
	RtClass
	*ISpeechSynthesisStream
}

type SpeechSynthesizer struct {
	RtClass
	*ISpeechSynthesizer
}

func NewSpeechSynthesizer() *SpeechSynthesizer {
	hs := NewHStr("Windows.Media.SpeechSynthesis.SpeechSynthesizer")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &SpeechSynthesizer{
		RtClass:            RtClass{PInspect: p},
		ISpeechSynthesizer: (*ISpeechSynthesizer)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewIInstalledVoicesStatic() *IInstalledVoicesStatic {
	var p *IInstalledVoicesStatic
	hs := NewHStr("Windows.Media.SpeechSynthesis.SpeechSynthesizer")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IInstalledVoicesStatic, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIInstalledVoicesStatic2() *IInstalledVoicesStatic2 {
	var p *IInstalledVoicesStatic2
	hs := NewHStr("Windows.Media.SpeechSynthesis.SpeechSynthesizer")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IInstalledVoicesStatic2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type SpeechSynthesizerOptions struct {
	RtClass
	*ISpeechSynthesizerOptions
}

type VoiceInformation struct {
	RtClass
	*IVoiceInformation
}
