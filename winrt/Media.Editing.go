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
type MediaTrimmingPreference int32

const (
	MediaTrimmingPreference_Fast    MediaTrimmingPreference = 0
	MediaTrimmingPreference_Precise MediaTrimmingPreference = 1
)

// enum
type VideoFramePrecision int32

const (
	VideoFramePrecision_NearestFrame    VideoFramePrecision = 0
	VideoFramePrecision_NearestKeyFrame VideoFramePrecision = 1
)

// interfaces

// 4B91B3BD-9E21-4266-A9C2-67DD011A2357
var IID_IBackgroundAudioTrack = syscall.GUID{0x4B91B3BD, 0x9E21, 0x4266,
	[8]byte{0xA9, 0xC2, 0x67, 0xDD, 0x01, 0x1A, 0x23, 0x57}}

type IBackgroundAudioTrackInterface interface {
	win32.IInspectableInterface
	Get_TrimTimeFromStart() TimeSpan
	Put_TrimTimeFromStart(value TimeSpan)
	Get_TrimTimeFromEnd() TimeSpan
	Put_TrimTimeFromEnd(value TimeSpan)
	Get_OriginalDuration() TimeSpan
	Get_TrimmedDuration() TimeSpan
	Get_UserData() *IMap[string, string]
	Put_Delay(value TimeSpan)
	Get_Delay() TimeSpan
	Put_Volume(value float64)
	Get_Volume() float64
	Clone() *IBackgroundAudioTrack
	GetAudioEncodingProperties() *IAudioEncodingProperties
	Get_AudioEffectDefinitions() *IVector[*IAudioEffectDefinition]
}

type IBackgroundAudioTrackVtbl struct {
	win32.IInspectableVtbl
	Get_TrimTimeFromStart      uintptr
	Put_TrimTimeFromStart      uintptr
	Get_TrimTimeFromEnd        uintptr
	Put_TrimTimeFromEnd        uintptr
	Get_OriginalDuration       uintptr
	Get_TrimmedDuration        uintptr
	Get_UserData               uintptr
	Put_Delay                  uintptr
	Get_Delay                  uintptr
	Put_Volume                 uintptr
	Get_Volume                 uintptr
	Clone                      uintptr
	GetAudioEncodingProperties uintptr
	Get_AudioEffectDefinitions uintptr
}

type IBackgroundAudioTrack struct {
	win32.IInspectable
}

func (this *IBackgroundAudioTrack) Vtbl() *IBackgroundAudioTrackVtbl {
	return (*IBackgroundAudioTrackVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundAudioTrack) Get_TrimTimeFromStart() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TrimTimeFromStart, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundAudioTrack) Put_TrimTimeFromStart(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TrimTimeFromStart, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IBackgroundAudioTrack) Get_TrimTimeFromEnd() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TrimTimeFromEnd, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundAudioTrack) Put_TrimTimeFromEnd(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TrimTimeFromEnd, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IBackgroundAudioTrack) Get_OriginalDuration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OriginalDuration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundAudioTrack) Get_TrimmedDuration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TrimmedDuration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundAudioTrack) Get_UserData() *IMap[string, string] {
	var _result *IMap[string, string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UserData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundAudioTrack) Put_Delay(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Delay, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IBackgroundAudioTrack) Get_Delay() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Delay, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundAudioTrack) Put_Volume(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Volume, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IBackgroundAudioTrack) Get_Volume() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Volume, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBackgroundAudioTrack) Clone() *IBackgroundAudioTrack {
	var _result *IBackgroundAudioTrack
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundAudioTrack) GetAudioEncodingProperties() *IAudioEncodingProperties {
	var _result *IAudioEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAudioEncodingProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundAudioTrack) Get_AudioEffectDefinitions() *IVector[*IAudioEffectDefinition] {
	var _result *IVector[*IAudioEffectDefinition]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioEffectDefinitions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D9B1C0D7-D018-42A8-A559-CB4D9E97E664
var IID_IBackgroundAudioTrackStatics = syscall.GUID{0xD9B1C0D7, 0xD018, 0x42A8,
	[8]byte{0xA5, 0x59, 0xCB, 0x4D, 0x9E, 0x97, 0xE6, 0x64}}

type IBackgroundAudioTrackStaticsInterface interface {
	win32.IInspectableInterface
	CreateFromEmbeddedAudioTrack(embeddedAudioTrack *IEmbeddedAudioTrack) *IBackgroundAudioTrack
	CreateFromFileAsync(file *IStorageFile) *IAsyncOperation[*IBackgroundAudioTrack]
}

type IBackgroundAudioTrackStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateFromEmbeddedAudioTrack uintptr
	CreateFromFileAsync          uintptr
}

type IBackgroundAudioTrackStatics struct {
	win32.IInspectable
}

func (this *IBackgroundAudioTrackStatics) Vtbl() *IBackgroundAudioTrackStaticsVtbl {
	return (*IBackgroundAudioTrackStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBackgroundAudioTrackStatics) CreateFromEmbeddedAudioTrack(embeddedAudioTrack *IEmbeddedAudioTrack) *IBackgroundAudioTrack {
	var _result *IBackgroundAudioTrack
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromEmbeddedAudioTrack, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(embeddedAudioTrack)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBackgroundAudioTrackStatics) CreateFromFileAsync(file *IStorageFile) *IAsyncOperation[*IBackgroundAudioTrack] {
	var _result *IAsyncOperation[*IBackgroundAudioTrack]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromFileAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 55EE5A7A-2D30-3FBA-A190-4F1A6454F88F
var IID_IEmbeddedAudioTrack = syscall.GUID{0x55EE5A7A, 0x2D30, 0x3FBA,
	[8]byte{0xA1, 0x90, 0x4F, 0x1A, 0x64, 0x54, 0xF8, 0x8F}}

type IEmbeddedAudioTrackInterface interface {
	win32.IInspectableInterface
	GetAudioEncodingProperties() *IAudioEncodingProperties
}

type IEmbeddedAudioTrackVtbl struct {
	win32.IInspectableVtbl
	GetAudioEncodingProperties uintptr
}

type IEmbeddedAudioTrack struct {
	win32.IInspectable
}

func (this *IEmbeddedAudioTrack) Vtbl() *IEmbeddedAudioTrackVtbl {
	return (*IEmbeddedAudioTrackVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IEmbeddedAudioTrack) GetAudioEncodingProperties() *IAudioEncodingProperties {
	var _result *IAudioEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAudioEncodingProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 53F25366-5FBA-3EA4-8693-24761811140A
var IID_IMediaClip = syscall.GUID{0x53F25366, 0x5FBA, 0x3EA4,
	[8]byte{0x86, 0x93, 0x24, 0x76, 0x18, 0x11, 0x14, 0x0A}}

type IMediaClipInterface interface {
	win32.IInspectableInterface
	Get_TrimTimeFromStart() TimeSpan
	Put_TrimTimeFromStart(value TimeSpan)
	Get_TrimTimeFromEnd() TimeSpan
	Put_TrimTimeFromEnd(value TimeSpan)
	Get_OriginalDuration() TimeSpan
	Get_TrimmedDuration() TimeSpan
	Get_UserData() *IMap[string, string]
	Clone() *IMediaClip
	Get_StartTimeInComposition() TimeSpan
	Get_EndTimeInComposition() TimeSpan
	Get_EmbeddedAudioTracks() *IVectorView[*IEmbeddedAudioTrack]
	Get_SelectedEmbeddedAudioTrackIndex() uint32
	Put_SelectedEmbeddedAudioTrackIndex(value uint32)
	Put_Volume(value float64)
	Get_Volume() float64
	GetVideoEncodingProperties() *IVideoEncodingProperties
	Get_AudioEffectDefinitions() *IVector[*IAudioEffectDefinition]
	Get_VideoEffectDefinitions() *IVector[*IVideoEffectDefinition]
}

type IMediaClipVtbl struct {
	win32.IInspectableVtbl
	Get_TrimTimeFromStart               uintptr
	Put_TrimTimeFromStart               uintptr
	Get_TrimTimeFromEnd                 uintptr
	Put_TrimTimeFromEnd                 uintptr
	Get_OriginalDuration                uintptr
	Get_TrimmedDuration                 uintptr
	Get_UserData                        uintptr
	Clone                               uintptr
	Get_StartTimeInComposition          uintptr
	Get_EndTimeInComposition            uintptr
	Get_EmbeddedAudioTracks             uintptr
	Get_SelectedEmbeddedAudioTrackIndex uintptr
	Put_SelectedEmbeddedAudioTrackIndex uintptr
	Put_Volume                          uintptr
	Get_Volume                          uintptr
	GetVideoEncodingProperties          uintptr
	Get_AudioEffectDefinitions          uintptr
	Get_VideoEffectDefinitions          uintptr
}

type IMediaClip struct {
	win32.IInspectable
}

func (this *IMediaClip) Vtbl() *IMediaClipVtbl {
	return (*IMediaClipVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaClip) Get_TrimTimeFromStart() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TrimTimeFromStart, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaClip) Put_TrimTimeFromStart(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TrimTimeFromStart, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IMediaClip) Get_TrimTimeFromEnd() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TrimTimeFromEnd, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaClip) Put_TrimTimeFromEnd(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TrimTimeFromEnd, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IMediaClip) Get_OriginalDuration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OriginalDuration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaClip) Get_TrimmedDuration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TrimmedDuration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaClip) Get_UserData() *IMap[string, string] {
	var _result *IMap[string, string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UserData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaClip) Clone() *IMediaClip {
	var _result *IMediaClip
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaClip) Get_StartTimeInComposition() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StartTimeInComposition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaClip) Get_EndTimeInComposition() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EndTimeInComposition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaClip) Get_EmbeddedAudioTracks() *IVectorView[*IEmbeddedAudioTrack] {
	var _result *IVectorView[*IEmbeddedAudioTrack]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EmbeddedAudioTracks, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaClip) Get_SelectedEmbeddedAudioTrackIndex() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SelectedEmbeddedAudioTrackIndex, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaClip) Put_SelectedEmbeddedAudioTrackIndex(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SelectedEmbeddedAudioTrackIndex, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMediaClip) Put_Volume(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Volume, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMediaClip) Get_Volume() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Volume, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaClip) GetVideoEncodingProperties() *IVideoEncodingProperties {
	var _result *IVideoEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetVideoEncodingProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaClip) Get_AudioEffectDefinitions() *IVector[*IAudioEffectDefinition] {
	var _result *IVector[*IAudioEffectDefinition]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioEffectDefinitions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaClip) Get_VideoEffectDefinitions() *IVector[*IVideoEffectDefinition] {
	var _result *IVector[*IVideoEffectDefinition]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoEffectDefinitions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FA402B68-928F-43C4-BC6E-783A1A359656
var IID_IMediaClipStatics = syscall.GUID{0xFA402B68, 0x928F, 0x43C4,
	[8]byte{0xBC, 0x6E, 0x78, 0x3A, 0x1A, 0x35, 0x96, 0x56}}

type IMediaClipStaticsInterface interface {
	win32.IInspectableInterface
	CreateFromColor(color unsafe.Pointer, originalDuration TimeSpan) *IMediaClip
	CreateFromFileAsync(file *IStorageFile) *IAsyncOperation[*IMediaClip]
	CreateFromImageFileAsync(file *IStorageFile, originalDuration TimeSpan) *IAsyncOperation[*IMediaClip]
}

type IMediaClipStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateFromColor          uintptr
	CreateFromFileAsync      uintptr
	CreateFromImageFileAsync uintptr
}

type IMediaClipStatics struct {
	win32.IInspectable
}

func (this *IMediaClipStatics) Vtbl() *IMediaClipStaticsVtbl {
	return (*IMediaClipStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaClipStatics) CreateFromColor(color unsafe.Pointer, originalDuration TimeSpan) *IMediaClip {
	var _result *IMediaClip
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromColor, uintptr(unsafe.Pointer(this)), uintptr(color), *(*uintptr)(unsafe.Pointer(&originalDuration)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaClipStatics) CreateFromFileAsync(file *IStorageFile) *IAsyncOperation[*IMediaClip] {
	var _result *IAsyncOperation[*IMediaClip]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromFileAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaClipStatics) CreateFromImageFileAsync(file *IStorageFile, originalDuration TimeSpan) *IAsyncOperation[*IMediaClip] {
	var _result *IAsyncOperation[*IMediaClip]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromImageFileAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), *(*uintptr)(unsafe.Pointer(&originalDuration)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5B1DD7B3-854E-4D9B-877D-4774A556CD12
var IID_IMediaClipStatics2 = syscall.GUID{0x5B1DD7B3, 0x854E, 0x4D9B,
	[8]byte{0x87, 0x7D, 0x47, 0x74, 0xA5, 0x56, 0xCD, 0x12}}

type IMediaClipStatics2Interface interface {
	win32.IInspectableInterface
	CreateFromSurface(surface unsafe.Pointer, originalDuration TimeSpan) *IMediaClip
}

type IMediaClipStatics2Vtbl struct {
	win32.IInspectableVtbl
	CreateFromSurface uintptr
}

type IMediaClipStatics2 struct {
	win32.IInspectable
}

func (this *IMediaClipStatics2) Vtbl() *IMediaClipStatics2Vtbl {
	return (*IMediaClipStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaClipStatics2) CreateFromSurface(surface unsafe.Pointer, originalDuration TimeSpan) *IMediaClip {
	var _result *IMediaClip
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromSurface, uintptr(unsafe.Pointer(this)), uintptr(surface), *(*uintptr)(unsafe.Pointer(&originalDuration)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2E06E605-DC71-41D6-B837-2D2BC14A2947
var IID_IMediaComposition = syscall.GUID{0x2E06E605, 0xDC71, 0x41D6,
	[8]byte{0xB8, 0x37, 0x2D, 0x2B, 0xC1, 0x4A, 0x29, 0x47}}

type IMediaCompositionInterface interface {
	win32.IInspectableInterface
	Get_Duration() TimeSpan
	Get_Clips() *IVector[*IMediaClip]
	Get_BackgroundAudioTracks() *IVector[*IBackgroundAudioTrack]
	Get_UserData() *IMap[string, string]
	Clone() *IMediaComposition
	SaveAsync(file *IStorageFile) *IAsyncAction
	GetThumbnailAsync(timeFromStart TimeSpan, scaledWidth int32, scaledHeight int32, framePrecision VideoFramePrecision) *IAsyncOperation[*IRandomAccessStreamWithContentType]
	GetThumbnailsAsync(timesFromStart *IIterable[TimeSpan], scaledWidth int32, scaledHeight int32, framePrecision VideoFramePrecision) *IAsyncOperation[*IVectorView[*IRandomAccessStreamWithContentType]]
	RenderToFileAsync(destination *IStorageFile) *IAsyncOperationWithProgress[TranscodeFailureReason, float64]
	RenderToFileWithTrimmingPreferenceAsync(destination *IStorageFile, trimmingPreference MediaTrimmingPreference) *IAsyncOperationWithProgress[TranscodeFailureReason, float64]
	RenderToFileWithProfileAsync(destination *IStorageFile, trimmingPreference MediaTrimmingPreference, encodingProfile *IMediaEncodingProfile) *IAsyncOperationWithProgress[TranscodeFailureReason, float64]
	CreateDefaultEncodingProfile() *IMediaEncodingProfile
	GenerateMediaStreamSource() *IMediaStreamSource
	GenerateMediaStreamSourceWithProfile(encodingProfile *IMediaEncodingProfile) *IMediaStreamSource
	GeneratePreviewMediaStreamSource(scaledWidth int32, scaledHeight int32) *IMediaStreamSource
}

type IMediaCompositionVtbl struct {
	win32.IInspectableVtbl
	Get_Duration                            uintptr
	Get_Clips                               uintptr
	Get_BackgroundAudioTracks               uintptr
	Get_UserData                            uintptr
	Clone                                   uintptr
	SaveAsync                               uintptr
	GetThumbnailAsync                       uintptr
	GetThumbnailsAsync                      uintptr
	RenderToFileAsync                       uintptr
	RenderToFileWithTrimmingPreferenceAsync uintptr
	RenderToFileWithProfileAsync            uintptr
	CreateDefaultEncodingProfile            uintptr
	GenerateMediaStreamSource               uintptr
	GenerateMediaStreamSourceWithProfile    uintptr
	GeneratePreviewMediaStreamSource        uintptr
}

type IMediaComposition struct {
	win32.IInspectable
}

func (this *IMediaComposition) Vtbl() *IMediaCompositionVtbl {
	return (*IMediaCompositionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaComposition) Get_Duration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Duration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaComposition) Get_Clips() *IVector[*IMediaClip] {
	var _result *IVector[*IMediaClip]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Clips, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaComposition) Get_BackgroundAudioTracks() *IVector[*IBackgroundAudioTrack] {
	var _result *IVector[*IBackgroundAudioTrack]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BackgroundAudioTracks, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaComposition) Get_UserData() *IMap[string, string] {
	var _result *IMap[string, string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UserData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaComposition) Clone() *IMediaComposition {
	var _result *IMediaComposition
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaComposition) SaveAsync(file *IStorageFile) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SaveAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaComposition) GetThumbnailAsync(timeFromStart TimeSpan, scaledWidth int32, scaledHeight int32, framePrecision VideoFramePrecision) *IAsyncOperation[*IRandomAccessStreamWithContentType] {
	var _result *IAsyncOperation[*IRandomAccessStreamWithContentType]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetThumbnailAsync, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&timeFromStart)), uintptr(scaledWidth), uintptr(scaledHeight), uintptr(framePrecision), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaComposition) GetThumbnailsAsync(timesFromStart *IIterable[TimeSpan], scaledWidth int32, scaledHeight int32, framePrecision VideoFramePrecision) *IAsyncOperation[*IVectorView[*IRandomAccessStreamWithContentType]] {
	var _result *IAsyncOperation[*IVectorView[*IRandomAccessStreamWithContentType]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetThumbnailsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(timesFromStart)), uintptr(scaledWidth), uintptr(scaledHeight), uintptr(framePrecision), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaComposition) RenderToFileAsync(destination *IStorageFile) *IAsyncOperationWithProgress[TranscodeFailureReason, float64] {
	var _result *IAsyncOperationWithProgress[TranscodeFailureReason, float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RenderToFileAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(destination)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaComposition) RenderToFileWithTrimmingPreferenceAsync(destination *IStorageFile, trimmingPreference MediaTrimmingPreference) *IAsyncOperationWithProgress[TranscodeFailureReason, float64] {
	var _result *IAsyncOperationWithProgress[TranscodeFailureReason, float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RenderToFileWithTrimmingPreferenceAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(destination)), uintptr(trimmingPreference), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaComposition) RenderToFileWithProfileAsync(destination *IStorageFile, trimmingPreference MediaTrimmingPreference, encodingProfile *IMediaEncodingProfile) *IAsyncOperationWithProgress[TranscodeFailureReason, float64] {
	var _result *IAsyncOperationWithProgress[TranscodeFailureReason, float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RenderToFileWithProfileAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(destination)), uintptr(trimmingPreference), uintptr(unsafe.Pointer(encodingProfile)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaComposition) CreateDefaultEncodingProfile() *IMediaEncodingProfile {
	var _result *IMediaEncodingProfile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDefaultEncodingProfile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaComposition) GenerateMediaStreamSource() *IMediaStreamSource {
	var _result *IMediaStreamSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GenerateMediaStreamSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaComposition) GenerateMediaStreamSourceWithProfile(encodingProfile *IMediaEncodingProfile) *IMediaStreamSource {
	var _result *IMediaStreamSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GenerateMediaStreamSourceWithProfile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(encodingProfile)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaComposition) GeneratePreviewMediaStreamSource(scaledWidth int32, scaledHeight int32) *IMediaStreamSource {
	var _result *IMediaStreamSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GeneratePreviewMediaStreamSource, uintptr(unsafe.Pointer(this)), uintptr(scaledWidth), uintptr(scaledHeight), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A59E5372-2366-492C-BEC8-E6DFBA6D0281
var IID_IMediaComposition2 = syscall.GUID{0xA59E5372, 0x2366, 0x492C,
	[8]byte{0xBE, 0xC8, 0xE6, 0xDF, 0xBA, 0x6D, 0x02, 0x81}}

type IMediaComposition2Interface interface {
	win32.IInspectableInterface
	Get_OverlayLayers() *IVector[*IMediaOverlayLayer]
}

type IMediaComposition2Vtbl struct {
	win32.IInspectableVtbl
	Get_OverlayLayers uintptr
}

type IMediaComposition2 struct {
	win32.IInspectable
}

func (this *IMediaComposition2) Vtbl() *IMediaComposition2Vtbl {
	return (*IMediaComposition2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaComposition2) Get_OverlayLayers() *IVector[*IMediaOverlayLayer] {
	var _result *IVector[*IMediaOverlayLayer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OverlayLayers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 87A08F04-E32A-45CE-8F66-A30DF0766224
var IID_IMediaCompositionStatics = syscall.GUID{0x87A08F04, 0xE32A, 0x45CE,
	[8]byte{0x8F, 0x66, 0xA3, 0x0D, 0xF0, 0x76, 0x62, 0x24}}

type IMediaCompositionStaticsInterface interface {
	win32.IInspectableInterface
	LoadAsync(file *IStorageFile) *IAsyncOperation[*IMediaComposition]
}

type IMediaCompositionStaticsVtbl struct {
	win32.IInspectableVtbl
	LoadAsync uintptr
}

type IMediaCompositionStatics struct {
	win32.IInspectable
}

func (this *IMediaCompositionStatics) Vtbl() *IMediaCompositionStaticsVtbl {
	return (*IMediaCompositionStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaCompositionStatics) LoadAsync(file *IStorageFile) *IAsyncOperation[*IMediaComposition] {
	var _result *IAsyncOperation[*IMediaComposition]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LoadAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A902AE5D-7869-4830-8AB1-94DC01C05FA4
var IID_IMediaOverlay = syscall.GUID{0xA902AE5D, 0x7869, 0x4830,
	[8]byte{0x8A, 0xB1, 0x94, 0xDC, 0x01, 0xC0, 0x5F, 0xA4}}

type IMediaOverlayInterface interface {
	win32.IInspectableInterface
	Get_Position() Rect
	Put_Position(value Rect)
	Put_Delay(value TimeSpan)
	Get_Delay() TimeSpan
	Get_Opacity() float64
	Put_Opacity(value float64)
	Clone() *IMediaOverlay
	Get_Clip() *IMediaClip
	Get_AudioEnabled() bool
	Put_AudioEnabled(value bool)
}

type IMediaOverlayVtbl struct {
	win32.IInspectableVtbl
	Get_Position     uintptr
	Put_Position     uintptr
	Put_Delay        uintptr
	Get_Delay        uintptr
	Get_Opacity      uintptr
	Put_Opacity      uintptr
	Clone            uintptr
	Get_Clip         uintptr
	Get_AudioEnabled uintptr
	Put_AudioEnabled uintptr
}

type IMediaOverlay struct {
	win32.IInspectable
}

func (this *IMediaOverlay) Vtbl() *IMediaOverlayVtbl {
	return (*IMediaOverlayVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaOverlay) Get_Position() Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaOverlay) Put_Position(value Rect) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IMediaOverlay) Put_Delay(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Delay, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IMediaOverlay) Get_Delay() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Delay, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaOverlay) Get_Opacity() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Opacity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaOverlay) Put_Opacity(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Opacity, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMediaOverlay) Clone() *IMediaOverlay {
	var _result *IMediaOverlay
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaOverlay) Get_Clip() *IMediaClip {
	var _result *IMediaClip
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Clip, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaOverlay) Get_AudioEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaOverlay) Put_AudioEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AudioEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// B584828A-6188-4F8F-A2E0-AA552D598E18
var IID_IMediaOverlayFactory = syscall.GUID{0xB584828A, 0x6188, 0x4F8F,
	[8]byte{0xA2, 0xE0, 0xAA, 0x55, 0x2D, 0x59, 0x8E, 0x18}}

type IMediaOverlayFactoryInterface interface {
	win32.IInspectableInterface
	Create(clip *IMediaClip) *IMediaOverlay
	CreateWithPositionAndOpacity(clip *IMediaClip, position Rect, opacity float64) *IMediaOverlay
}

type IMediaOverlayFactoryVtbl struct {
	win32.IInspectableVtbl
	Create                       uintptr
	CreateWithPositionAndOpacity uintptr
}

type IMediaOverlayFactory struct {
	win32.IInspectable
}

func (this *IMediaOverlayFactory) Vtbl() *IMediaOverlayFactoryVtbl {
	return (*IMediaOverlayFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaOverlayFactory) Create(clip *IMediaClip) *IMediaOverlay {
	var _result *IMediaOverlay
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(clip)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaOverlayFactory) CreateWithPositionAndOpacity(clip *IMediaClip, position Rect, opacity float64) *IMediaOverlay {
	var _result *IMediaOverlay
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithPositionAndOpacity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(clip)), uintptr(unsafe.Pointer(&position)), uintptr(opacity), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A6D9BA57-EEDA-46C6-BBE5-E398C84168AC
var IID_IMediaOverlayLayer = syscall.GUID{0xA6D9BA57, 0xEEDA, 0x46C6,
	[8]byte{0xBB, 0xE5, 0xE3, 0x98, 0xC8, 0x41, 0x68, 0xAC}}

type IMediaOverlayLayerInterface interface {
	win32.IInspectableInterface
	Clone() *IMediaOverlayLayer
	Get_Overlays() *IVector[*IMediaOverlay]
	Get_CustomCompositorDefinition() *IVideoCompositorDefinition
}

type IMediaOverlayLayerVtbl struct {
	win32.IInspectableVtbl
	Clone                          uintptr
	Get_Overlays                   uintptr
	Get_CustomCompositorDefinition uintptr
}

type IMediaOverlayLayer struct {
	win32.IInspectable
}

func (this *IMediaOverlayLayer) Vtbl() *IMediaOverlayLayerVtbl {
	return (*IMediaOverlayLayerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaOverlayLayer) Clone() *IMediaOverlayLayer {
	var _result *IMediaOverlayLayer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaOverlayLayer) Get_Overlays() *IVector[*IMediaOverlay] {
	var _result *IVector[*IMediaOverlay]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Overlays, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaOverlayLayer) Get_CustomCompositorDefinition() *IVideoCompositorDefinition {
	var _result *IVideoCompositorDefinition
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CustomCompositorDefinition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 947CB473-A39E-4362-ABBF-9F8B5070A062
var IID_IMediaOverlayLayerFactory = syscall.GUID{0x947CB473, 0xA39E, 0x4362,
	[8]byte{0xAB, 0xBF, 0x9F, 0x8B, 0x50, 0x70, 0xA0, 0x62}}

type IMediaOverlayLayerFactoryInterface interface {
	win32.IInspectableInterface
	CreateWithCompositorDefinition(compositorDefinition *IVideoCompositorDefinition) *IMediaOverlayLayer
}

type IMediaOverlayLayerFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateWithCompositorDefinition uintptr
}

type IMediaOverlayLayerFactory struct {
	win32.IInspectable
}

func (this *IMediaOverlayLayerFactory) Vtbl() *IMediaOverlayLayerFactoryVtbl {
	return (*IMediaOverlayLayerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaOverlayLayerFactory) CreateWithCompositorDefinition(compositorDefinition *IVideoCompositorDefinition) *IMediaOverlayLayer {
	var _result *IMediaOverlayLayer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithCompositorDefinition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(compositorDefinition)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type BackgroundAudioTrack struct {
	RtClass
	*IBackgroundAudioTrack
}

func NewIBackgroundAudioTrackStatics() *IBackgroundAudioTrackStatics {
	var p *IBackgroundAudioTrackStatics
	hs := NewHStr("Windows.Media.Editing.BackgroundAudioTrack")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IBackgroundAudioTrackStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type EmbeddedAudioTrack struct {
	RtClass
	*IEmbeddedAudioTrack
}

type MediaClip struct {
	RtClass
	*IMediaClip
}

func NewIMediaClipStatics2() *IMediaClipStatics2 {
	var p *IMediaClipStatics2
	hs := NewHStr("Windows.Media.Editing.MediaClip")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMediaClipStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIMediaClipStatics() *IMediaClipStatics {
	var p *IMediaClipStatics
	hs := NewHStr("Windows.Media.Editing.MediaClip")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMediaClipStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type MediaComposition struct {
	RtClass
	*IMediaComposition
}

func NewMediaComposition() *MediaComposition {
	hs := NewHStr("Windows.Media.Editing.MediaComposition")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &MediaComposition{
		RtClass:           RtClass{PInspect: p},
		IMediaComposition: (*IMediaComposition)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewIMediaCompositionStatics() *IMediaCompositionStatics {
	var p *IMediaCompositionStatics
	hs := NewHStr("Windows.Media.Editing.MediaComposition")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMediaCompositionStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type MediaOverlay struct {
	RtClass
	*IMediaOverlay
}

func NewMediaOverlay_Create(clip *IMediaClip) *MediaOverlay {
	hs := NewHStr("Windows.Media.Editing.MediaOverlay")
	var pFac *IMediaOverlayFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMediaOverlayFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IMediaOverlay
	p = pFac.Create(clip)
	result := &MediaOverlay{
		RtClass:       RtClass{PInspect: &p.IInspectable},
		IMediaOverlay: p,
	}
	com.AddToScope(result)
	return result
}

func NewMediaOverlay_CreateWithPositionAndOpacity(clip *IMediaClip, position Rect, opacity float64) *MediaOverlay {
	hs := NewHStr("Windows.Media.Editing.MediaOverlay")
	var pFac *IMediaOverlayFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMediaOverlayFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IMediaOverlay
	p = pFac.CreateWithPositionAndOpacity(clip, position, opacity)
	result := &MediaOverlay{
		RtClass:       RtClass{PInspect: &p.IInspectable},
		IMediaOverlay: p,
	}
	com.AddToScope(result)
	return result
}

type MediaOverlayLayer struct {
	RtClass
	*IMediaOverlayLayer
}

func NewMediaOverlayLayer() *MediaOverlayLayer {
	hs := NewHStr("Windows.Media.Editing.MediaOverlayLayer")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &MediaOverlayLayer{
		RtClass:            RtClass{PInspect: p},
		IMediaOverlayLayer: (*IMediaOverlayLayer)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewMediaOverlayLayer_CreateWithCompositorDefinition(compositorDefinition *IVideoCompositorDefinition) *MediaOverlayLayer {
	hs := NewHStr("Windows.Media.Editing.MediaOverlayLayer")
	var pFac *IMediaOverlayLayerFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMediaOverlayLayerFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IMediaOverlayLayer
	p = pFac.CreateWithCompositorDefinition(compositorDefinition)
	result := &MediaOverlayLayer{
		RtClass:            RtClass{PInspect: &p.IInspectable},
		IMediaOverlayLayer: p,
	}
	com.AddToScope(result)
	return result
}
