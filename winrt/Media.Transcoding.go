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
type MediaVideoProcessingAlgorithm int32

const (
	MediaVideoProcessingAlgorithm_Default   MediaVideoProcessingAlgorithm = 0
	MediaVideoProcessingAlgorithm_MrfCrf444 MediaVideoProcessingAlgorithm = 1
)

// enum
type TranscodeFailureReason int32

const (
	TranscodeFailureReason_None           TranscodeFailureReason = 0
	TranscodeFailureReason_Unknown        TranscodeFailureReason = 1
	TranscodeFailureReason_InvalidProfile TranscodeFailureReason = 2
	TranscodeFailureReason_CodecNotFound  TranscodeFailureReason = 3
)

// interfaces

// 190C99D2-A0AA-4D34-86BC-EED1B12C2F5B
var IID_IMediaTranscoder = syscall.GUID{0x190C99D2, 0xA0AA, 0x4D34,
	[8]byte{0x86, 0xBC, 0xEE, 0xD1, 0xB1, 0x2C, 0x2F, 0x5B}}

type IMediaTranscoderInterface interface {
	win32.IInspectableInterface
	Put_TrimStartTime(value TimeSpan)
	Get_TrimStartTime() TimeSpan
	Put_TrimStopTime(value TimeSpan)
	Get_TrimStopTime() TimeSpan
	Put_AlwaysReencode(value bool)
	Get_AlwaysReencode() bool
	Put_HardwareAccelerationEnabled(value bool)
	Get_HardwareAccelerationEnabled() bool
	AddAudioEffect(activatableClassId string)
	AddAudioEffectWithSettings(activatableClassId string, effectRequired bool, configuration *IPropertySet)
	AddVideoEffect(activatableClassId string)
	AddVideoEffectWithSettings(activatableClassId string, effectRequired bool, configuration *IPropertySet)
	ClearEffects()
	PrepareFileTranscodeAsync(source *IStorageFile, destination *IStorageFile, profile *IMediaEncodingProfile) *IAsyncOperation[*IPrepareTranscodeResult]
	PrepareStreamTranscodeAsync(source *IRandomAccessStream, destination *IRandomAccessStream, profile *IMediaEncodingProfile) *IAsyncOperation[*IPrepareTranscodeResult]
}

type IMediaTranscoderVtbl struct {
	win32.IInspectableVtbl
	Put_TrimStartTime               uintptr
	Get_TrimStartTime               uintptr
	Put_TrimStopTime                uintptr
	Get_TrimStopTime                uintptr
	Put_AlwaysReencode              uintptr
	Get_AlwaysReencode              uintptr
	Put_HardwareAccelerationEnabled uintptr
	Get_HardwareAccelerationEnabled uintptr
	AddAudioEffect                  uintptr
	AddAudioEffectWithSettings      uintptr
	AddVideoEffect                  uintptr
	AddVideoEffectWithSettings      uintptr
	ClearEffects                    uintptr
	PrepareFileTranscodeAsync       uintptr
	PrepareStreamTranscodeAsync     uintptr
}

type IMediaTranscoder struct {
	win32.IInspectable
}

func (this *IMediaTranscoder) Vtbl() *IMediaTranscoderVtbl {
	return (*IMediaTranscoderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaTranscoder) Put_TrimStartTime(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TrimStartTime, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IMediaTranscoder) Get_TrimStartTime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TrimStartTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaTranscoder) Put_TrimStopTime(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TrimStopTime, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IMediaTranscoder) Get_TrimStopTime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TrimStopTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaTranscoder) Put_AlwaysReencode(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AlwaysReencode, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaTranscoder) Get_AlwaysReencode() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlwaysReencode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaTranscoder) Put_HardwareAccelerationEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_HardwareAccelerationEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IMediaTranscoder) Get_HardwareAccelerationEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HardwareAccelerationEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaTranscoder) AddAudioEffect(activatableClassId string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddAudioEffect, uintptr(unsafe.Pointer(this)), NewHStr(activatableClassId).Ptr)
	_ = _hr
}

func (this *IMediaTranscoder) AddAudioEffectWithSettings(activatableClassId string, effectRequired bool, configuration *IPropertySet) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddAudioEffectWithSettings, uintptr(unsafe.Pointer(this)), NewHStr(activatableClassId).Ptr, uintptr(*(*byte)(unsafe.Pointer(&effectRequired))), uintptr(unsafe.Pointer(configuration)))
	_ = _hr
}

func (this *IMediaTranscoder) AddVideoEffect(activatableClassId string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddVideoEffect, uintptr(unsafe.Pointer(this)), NewHStr(activatableClassId).Ptr)
	_ = _hr
}

func (this *IMediaTranscoder) AddVideoEffectWithSettings(activatableClassId string, effectRequired bool, configuration *IPropertySet) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddVideoEffectWithSettings, uintptr(unsafe.Pointer(this)), NewHStr(activatableClassId).Ptr, uintptr(*(*byte)(unsafe.Pointer(&effectRequired))), uintptr(unsafe.Pointer(configuration)))
	_ = _hr
}

func (this *IMediaTranscoder) ClearEffects() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearEffects, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IMediaTranscoder) PrepareFileTranscodeAsync(source *IStorageFile, destination *IStorageFile, profile *IMediaEncodingProfile) *IAsyncOperation[*IPrepareTranscodeResult] {
	var _result *IAsyncOperation[*IPrepareTranscodeResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PrepareFileTranscodeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(source)), uintptr(unsafe.Pointer(destination)), uintptr(unsafe.Pointer(profile)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaTranscoder) PrepareStreamTranscodeAsync(source *IRandomAccessStream, destination *IRandomAccessStream, profile *IMediaEncodingProfile) *IAsyncOperation[*IPrepareTranscodeResult] {
	var _result *IAsyncOperation[*IPrepareTranscodeResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PrepareStreamTranscodeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(source)), uintptr(unsafe.Pointer(destination)), uintptr(unsafe.Pointer(profile)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 40531D74-35E0-4F04-8574-CA8BC4E5A082
var IID_IMediaTranscoder2 = syscall.GUID{0x40531D74, 0x35E0, 0x4F04,
	[8]byte{0x85, 0x74, 0xCA, 0x8B, 0xC4, 0xE5, 0xA0, 0x82}}

type IMediaTranscoder2Interface interface {
	win32.IInspectableInterface
	PrepareMediaStreamSourceTranscodeAsync(source *IMediaSource, destination *IRandomAccessStream, profile *IMediaEncodingProfile) *IAsyncOperation[*IPrepareTranscodeResult]
	Put_VideoProcessingAlgorithm(value MediaVideoProcessingAlgorithm)
	Get_VideoProcessingAlgorithm() MediaVideoProcessingAlgorithm
}

type IMediaTranscoder2Vtbl struct {
	win32.IInspectableVtbl
	PrepareMediaStreamSourceTranscodeAsync uintptr
	Put_VideoProcessingAlgorithm           uintptr
	Get_VideoProcessingAlgorithm           uintptr
}

type IMediaTranscoder2 struct {
	win32.IInspectable
}

func (this *IMediaTranscoder2) Vtbl() *IMediaTranscoder2Vtbl {
	return (*IMediaTranscoder2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaTranscoder2) PrepareMediaStreamSourceTranscodeAsync(source *IMediaSource, destination *IRandomAccessStream, profile *IMediaEncodingProfile) *IAsyncOperation[*IPrepareTranscodeResult] {
	var _result *IAsyncOperation[*IPrepareTranscodeResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PrepareMediaStreamSourceTranscodeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(source)), uintptr(unsafe.Pointer(destination)), uintptr(unsafe.Pointer(profile)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaTranscoder2) Put_VideoProcessingAlgorithm(value MediaVideoProcessingAlgorithm) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_VideoProcessingAlgorithm, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMediaTranscoder2) Get_VideoProcessingAlgorithm() MediaVideoProcessingAlgorithm {
	var _result MediaVideoProcessingAlgorithm
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoProcessingAlgorithm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 05F25DCE-994F-4A34-9D68-97CCCE1730D6
var IID_IPrepareTranscodeResult = syscall.GUID{0x05F25DCE, 0x994F, 0x4A34,
	[8]byte{0x9D, 0x68, 0x97, 0xCC, 0xCE, 0x17, 0x30, 0xD6}}

type IPrepareTranscodeResultInterface interface {
	win32.IInspectableInterface
	Get_CanTranscode() bool
	Get_FailureReason() TranscodeFailureReason
	TranscodeAsync() *IAsyncActionWithProgress[float64]
}

type IPrepareTranscodeResultVtbl struct {
	win32.IInspectableVtbl
	Get_CanTranscode  uintptr
	Get_FailureReason uintptr
	TranscodeAsync    uintptr
}

type IPrepareTranscodeResult struct {
	win32.IInspectable
}

func (this *IPrepareTranscodeResult) Vtbl() *IPrepareTranscodeResultVtbl {
	return (*IPrepareTranscodeResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPrepareTranscodeResult) Get_CanTranscode() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanTranscode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrepareTranscodeResult) Get_FailureReason() TranscodeFailureReason {
	var _result TranscodeFailureReason
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FailureReason, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPrepareTranscodeResult) TranscodeAsync() *IAsyncActionWithProgress[float64] {
	var _result *IAsyncActionWithProgress[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TranscodeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type MediaTranscoder struct {
	RtClass
	*IMediaTranscoder
}

func NewMediaTranscoder() *MediaTranscoder {
	hs := NewHStr("Windows.Media.Transcoding.MediaTranscoder")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &MediaTranscoder{
		RtClass:          RtClass{PInspect: p},
		IMediaTranscoder: (*IMediaTranscoder)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type PrepareTranscodeResult struct {
	RtClass
	*IPrepareTranscodeResult
}
