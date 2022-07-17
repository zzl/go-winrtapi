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
type AudioEncodingQuality int32

const (
	AudioEncodingQuality_Auto   AudioEncodingQuality = 0
	AudioEncodingQuality_High   AudioEncodingQuality = 1
	AudioEncodingQuality_Medium AudioEncodingQuality = 2
	AudioEncodingQuality_Low    AudioEncodingQuality = 3
)

// enum
// flags
type MediaMirroringOptions uint32

const (
	MediaMirroringOptions_None       MediaMirroringOptions = 0
	MediaMirroringOptions_Horizontal MediaMirroringOptions = 1
	MediaMirroringOptions_Vertical   MediaMirroringOptions = 2
)

// enum
type MediaPixelFormat int32

const (
	MediaPixelFormat_Nv12  MediaPixelFormat = 0
	MediaPixelFormat_Bgra8 MediaPixelFormat = 1
	MediaPixelFormat_P010  MediaPixelFormat = 2
)

// enum
type MediaRotation int32

const (
	MediaRotation_None                MediaRotation = 0
	MediaRotation_Clockwise90Degrees  MediaRotation = 1
	MediaRotation_Clockwise180Degrees MediaRotation = 2
	MediaRotation_Clockwise270Degrees MediaRotation = 3
)

// enum
type MediaThumbnailFormat int32

const (
	MediaThumbnailFormat_Bmp   MediaThumbnailFormat = 0
	MediaThumbnailFormat_Bgra8 MediaThumbnailFormat = 1
)

// enum
type SphericalVideoFrameFormat int32

const (
	SphericalVideoFrameFormat_None            SphericalVideoFrameFormat = 0
	SphericalVideoFrameFormat_Unsupported     SphericalVideoFrameFormat = 1
	SphericalVideoFrameFormat_Equirectangular SphericalVideoFrameFormat = 2
)

// enum
type StereoscopicVideoPackingMode int32

const (
	StereoscopicVideoPackingMode_None       StereoscopicVideoPackingMode = 0
	StereoscopicVideoPackingMode_SideBySide StereoscopicVideoPackingMode = 1
	StereoscopicVideoPackingMode_TopBottom  StereoscopicVideoPackingMode = 2
)

// enum
type VideoEncodingQuality int32

const (
	VideoEncodingQuality_Auto     VideoEncodingQuality = 0
	VideoEncodingQuality_HD1080p  VideoEncodingQuality = 1
	VideoEncodingQuality_HD720p   VideoEncodingQuality = 2
	VideoEncodingQuality_Wvga     VideoEncodingQuality = 3
	VideoEncodingQuality_Ntsc     VideoEncodingQuality = 4
	VideoEncodingQuality_Pal      VideoEncodingQuality = 5
	VideoEncodingQuality_Vga      VideoEncodingQuality = 6
	VideoEncodingQuality_Qvga     VideoEncodingQuality = 7
	VideoEncodingQuality_Uhd2160p VideoEncodingQuality = 8
	VideoEncodingQuality_Uhd4320p VideoEncodingQuality = 9
)

// interfaces

// 62BC7A16-005C-4B3B-8A0B-0A090E9687F3
var IID_IAudioEncodingProperties = syscall.GUID{0x62BC7A16, 0x005C, 0x4B3B,
	[8]byte{0x8A, 0x0B, 0x0A, 0x09, 0x0E, 0x96, 0x87, 0xF3}}

type IAudioEncodingPropertiesInterface interface {
	win32.IInspectableInterface
	Put_Bitrate(value uint32)
	Get_Bitrate() uint32
	Put_ChannelCount(value uint32)
	Get_ChannelCount() uint32
	Put_SampleRate(value uint32)
	Get_SampleRate() uint32
	Put_BitsPerSample(value uint32)
	Get_BitsPerSample() uint32
}

type IAudioEncodingPropertiesVtbl struct {
	win32.IInspectableVtbl
	Put_Bitrate       uintptr
	Get_Bitrate       uintptr
	Put_ChannelCount  uintptr
	Get_ChannelCount  uintptr
	Put_SampleRate    uintptr
	Get_SampleRate    uintptr
	Put_BitsPerSample uintptr
	Get_BitsPerSample uintptr
}

type IAudioEncodingProperties struct {
	win32.IInspectable
}

func (this *IAudioEncodingProperties) Vtbl() *IAudioEncodingPropertiesVtbl {
	return (*IAudioEncodingPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioEncodingProperties) Put_Bitrate(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Bitrate, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAudioEncodingProperties) Get_Bitrate() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bitrate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioEncodingProperties) Put_ChannelCount(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ChannelCount, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAudioEncodingProperties) Get_ChannelCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ChannelCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioEncodingProperties) Put_SampleRate(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SampleRate, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAudioEncodingProperties) Get_SampleRate() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SampleRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioEncodingProperties) Put_BitsPerSample(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BitsPerSample, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAudioEncodingProperties) Get_BitsPerSample() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BitsPerSample, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C45D54DA-80BD-4C23-80D5-72D4A181E894
var IID_IAudioEncodingProperties2 = syscall.GUID{0xC45D54DA, 0x80BD, 0x4C23,
	[8]byte{0x80, 0xD5, 0x72, 0xD4, 0xA1, 0x81, 0xE8, 0x94}}

type IAudioEncodingProperties2Interface interface {
	win32.IInspectableInterface
	Get_IsSpatial() bool
}

type IAudioEncodingProperties2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsSpatial uintptr
}

type IAudioEncodingProperties2 struct {
	win32.IInspectable
}

func (this *IAudioEncodingProperties2) Vtbl() *IAudioEncodingProperties2Vtbl {
	return (*IAudioEncodingProperties2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioEncodingProperties2) Get_IsSpatial() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSpatial, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 87600341-748C-4F8D-B0FD-10CAF08FF087
var IID_IAudioEncodingProperties3 = syscall.GUID{0x87600341, 0x748C, 0x4F8D,
	[8]byte{0xB0, 0xFD, 0x10, 0xCA, 0xF0, 0x8F, 0xF0, 0x87}}

type IAudioEncodingProperties3Interface interface {
	win32.IInspectableInterface
	Copy() *IAudioEncodingProperties
}

type IAudioEncodingProperties3Vtbl struct {
	win32.IInspectableVtbl
	Copy uintptr
}

type IAudioEncodingProperties3 struct {
	win32.IInspectable
}

func (this *IAudioEncodingProperties3) Vtbl() *IAudioEncodingProperties3Vtbl {
	return (*IAudioEncodingProperties3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioEncodingProperties3) Copy() *IAudioEncodingProperties {
	var _result *IAudioEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Copy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0CAD332C-EBE9-4527-B36D-E42A13CF38DB
var IID_IAudioEncodingPropertiesStatics = syscall.GUID{0x0CAD332C, 0xEBE9, 0x4527,
	[8]byte{0xB3, 0x6D, 0xE4, 0x2A, 0x13, 0xCF, 0x38, 0xDB}}

type IAudioEncodingPropertiesStaticsInterface interface {
	win32.IInspectableInterface
	CreateAac(sampleRate uint32, channelCount uint32, bitrate uint32) *IAudioEncodingProperties
	CreateAacAdts(sampleRate uint32, channelCount uint32, bitrate uint32) *IAudioEncodingProperties
	CreateMp3(sampleRate uint32, channelCount uint32, bitrate uint32) *IAudioEncodingProperties
	CreatePcm(sampleRate uint32, channelCount uint32, bitsPerSample uint32) *IAudioEncodingProperties
	CreateWma(sampleRate uint32, channelCount uint32, bitrate uint32) *IAudioEncodingProperties
}

type IAudioEncodingPropertiesStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateAac     uintptr
	CreateAacAdts uintptr
	CreateMp3     uintptr
	CreatePcm     uintptr
	CreateWma     uintptr
}

type IAudioEncodingPropertiesStatics struct {
	win32.IInspectable
}

func (this *IAudioEncodingPropertiesStatics) Vtbl() *IAudioEncodingPropertiesStaticsVtbl {
	return (*IAudioEncodingPropertiesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioEncodingPropertiesStatics) CreateAac(sampleRate uint32, channelCount uint32, bitrate uint32) *IAudioEncodingProperties {
	var _result *IAudioEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAac, uintptr(unsafe.Pointer(this)), uintptr(sampleRate), uintptr(channelCount), uintptr(bitrate), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioEncodingPropertiesStatics) CreateAacAdts(sampleRate uint32, channelCount uint32, bitrate uint32) *IAudioEncodingProperties {
	var _result *IAudioEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAacAdts, uintptr(unsafe.Pointer(this)), uintptr(sampleRate), uintptr(channelCount), uintptr(bitrate), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioEncodingPropertiesStatics) CreateMp3(sampleRate uint32, channelCount uint32, bitrate uint32) *IAudioEncodingProperties {
	var _result *IAudioEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateMp3, uintptr(unsafe.Pointer(this)), uintptr(sampleRate), uintptr(channelCount), uintptr(bitrate), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioEncodingPropertiesStatics) CreatePcm(sampleRate uint32, channelCount uint32, bitsPerSample uint32) *IAudioEncodingProperties {
	var _result *IAudioEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreatePcm, uintptr(unsafe.Pointer(this)), uintptr(sampleRate), uintptr(channelCount), uintptr(bitsPerSample), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioEncodingPropertiesStatics) CreateWma(sampleRate uint32, channelCount uint32, bitrate uint32) *IAudioEncodingProperties {
	var _result *IAudioEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWma, uintptr(unsafe.Pointer(this)), uintptr(sampleRate), uintptr(channelCount), uintptr(bitrate), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7489316F-77A0-433D-8ED5-4040280E8665
var IID_IAudioEncodingPropertiesStatics2 = syscall.GUID{0x7489316F, 0x77A0, 0x433D,
	[8]byte{0x8E, 0xD5, 0x40, 0x40, 0x28, 0x0E, 0x86, 0x65}}

type IAudioEncodingPropertiesStatics2Interface interface {
	win32.IInspectableInterface
	CreateAlac(sampleRate uint32, channelCount uint32, bitsPerSample uint32) *IAudioEncodingProperties
	CreateFlac(sampleRate uint32, channelCount uint32, bitsPerSample uint32) *IAudioEncodingProperties
}

type IAudioEncodingPropertiesStatics2Vtbl struct {
	win32.IInspectableVtbl
	CreateAlac uintptr
	CreateFlac uintptr
}

type IAudioEncodingPropertiesStatics2 struct {
	win32.IInspectable
}

func (this *IAudioEncodingPropertiesStatics2) Vtbl() *IAudioEncodingPropertiesStatics2Vtbl {
	return (*IAudioEncodingPropertiesStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioEncodingPropertiesStatics2) CreateAlac(sampleRate uint32, channelCount uint32, bitsPerSample uint32) *IAudioEncodingProperties {
	var _result *IAudioEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAlac, uintptr(unsafe.Pointer(this)), uintptr(sampleRate), uintptr(channelCount), uintptr(bitsPerSample), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioEncodingPropertiesStatics2) CreateFlac(sampleRate uint32, channelCount uint32, bitsPerSample uint32) *IAudioEncodingProperties {
	var _result *IAudioEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFlac, uintptr(unsafe.Pointer(this)), uintptr(sampleRate), uintptr(channelCount), uintptr(bitsPerSample), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 98F10D79-13EA-49FF-BE70-2673DB69702C
var IID_IAudioEncodingPropertiesWithFormatUserData = syscall.GUID{0x98F10D79, 0x13EA, 0x49FF,
	[8]byte{0xBE, 0x70, 0x26, 0x73, 0xDB, 0x69, 0x70, 0x2C}}

type IAudioEncodingPropertiesWithFormatUserDataInterface interface {
	win32.IInspectableInterface
	SetFormatUserData(valueLength uint32, value *byte)
	GetFormatUserData(valueLength uint32, value *byte)
}

type IAudioEncodingPropertiesWithFormatUserDataVtbl struct {
	win32.IInspectableVtbl
	SetFormatUserData uintptr
	GetFormatUserData uintptr
}

type IAudioEncodingPropertiesWithFormatUserData struct {
	win32.IInspectable
}

func (this *IAudioEncodingPropertiesWithFormatUserData) Vtbl() *IAudioEncodingPropertiesWithFormatUserDataVtbl {
	return (*IAudioEncodingPropertiesWithFormatUserDataVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioEncodingPropertiesWithFormatUserData) SetFormatUserData(valueLength uint32, value *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetFormatUserData, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAudioEncodingPropertiesWithFormatUserData) GetFormatUserData(valueLength uint32, value *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFormatUserData, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 59AC2A57-B32A-479E-8A61-4B7F2E9E7EA0
var IID_IContainerEncodingProperties = syscall.GUID{0x59AC2A57, 0xB32A, 0x479E,
	[8]byte{0x8A, 0x61, 0x4B, 0x7F, 0x2E, 0x9E, 0x7E, 0xA0}}

type IContainerEncodingPropertiesInterface interface {
	win32.IInspectableInterface
}

type IContainerEncodingPropertiesVtbl struct {
	win32.IInspectableVtbl
}

type IContainerEncodingProperties struct {
	win32.IInspectable
}

func (this *IContainerEncodingProperties) Vtbl() *IContainerEncodingPropertiesVtbl {
	return (*IContainerEncodingPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// B272C029-AE26-4819-BAAD-AD7A49B0A876
var IID_IContainerEncodingProperties2 = syscall.GUID{0xB272C029, 0xAE26, 0x4819,
	[8]byte{0xBA, 0xAD, 0xAD, 0x7A, 0x49, 0xB0, 0xA8, 0x76}}

type IContainerEncodingProperties2Interface interface {
	win32.IInspectableInterface
	Copy() *IContainerEncodingProperties
}

type IContainerEncodingProperties2Vtbl struct {
	win32.IInspectableVtbl
	Copy uintptr
}

type IContainerEncodingProperties2 struct {
	win32.IInspectable
}

func (this *IContainerEncodingProperties2) Vtbl() *IContainerEncodingProperties2Vtbl {
	return (*IContainerEncodingProperties2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IContainerEncodingProperties2) Copy() *IContainerEncodingProperties {
	var _result *IContainerEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Copy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 38654CA7-846A-4F97-A2E5-C3A15BBF70FD
var IID_IH264ProfileIdsStatics = syscall.GUID{0x38654CA7, 0x846A, 0x4F97,
	[8]byte{0xA2, 0xE5, 0xC3, 0xA1, 0x5B, 0xBF, 0x70, 0xFD}}

type IH264ProfileIdsStaticsInterface interface {
	win32.IInspectableInterface
	Get_ConstrainedBaseline() int32
	Get_Baseline() int32
	Get_Extended() int32
	Get_Main() int32
	Get_High() int32
	Get_High10() int32
	Get_High422() int32
	Get_High444() int32
	Get_StereoHigh() int32
	Get_MultiviewHigh() int32
}

type IH264ProfileIdsStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_ConstrainedBaseline uintptr
	Get_Baseline            uintptr
	Get_Extended            uintptr
	Get_Main                uintptr
	Get_High                uintptr
	Get_High10              uintptr
	Get_High422             uintptr
	Get_High444             uintptr
	Get_StereoHigh          uintptr
	Get_MultiviewHigh       uintptr
}

type IH264ProfileIdsStatics struct {
	win32.IInspectable
}

func (this *IH264ProfileIdsStatics) Vtbl() *IH264ProfileIdsStaticsVtbl {
	return (*IH264ProfileIdsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IH264ProfileIdsStatics) Get_ConstrainedBaseline() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ConstrainedBaseline, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IH264ProfileIdsStatics) Get_Baseline() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Baseline, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IH264ProfileIdsStatics) Get_Extended() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Extended, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IH264ProfileIdsStatics) Get_Main() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Main, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IH264ProfileIdsStatics) Get_High() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_High, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IH264ProfileIdsStatics) Get_High10() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_High10, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IH264ProfileIdsStatics) Get_High422() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_High422, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IH264ProfileIdsStatics) Get_High444() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_High444, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IH264ProfileIdsStatics) Get_StereoHigh() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StereoHigh, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IH264ProfileIdsStatics) Get_MultiviewHigh() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MultiviewHigh, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 78625635-F331-4189-B1C3-B48D5AE034F1
var IID_IImageEncodingProperties = syscall.GUID{0x78625635, 0xF331, 0x4189,
	[8]byte{0xB1, 0xC3, 0xB4, 0x8D, 0x5A, 0xE0, 0x34, 0xF1}}

type IImageEncodingPropertiesInterface interface {
	win32.IInspectableInterface
	Put_Width(value uint32)
	Get_Width() uint32
	Put_Height(value uint32)
	Get_Height() uint32
}

type IImageEncodingPropertiesVtbl struct {
	win32.IInspectableVtbl
	Put_Width  uintptr
	Get_Width  uintptr
	Put_Height uintptr
	Get_Height uintptr
}

type IImageEncodingProperties struct {
	win32.IInspectable
}

func (this *IImageEncodingProperties) Vtbl() *IImageEncodingPropertiesVtbl {
	return (*IImageEncodingPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IImageEncodingProperties) Put_Width(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Width, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IImageEncodingProperties) Get_Width() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Width, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageEncodingProperties) Put_Height(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Height, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IImageEncodingProperties) Get_Height() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Height, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C854A2DF-C923-469B-AC8E-6A9F3C1CD9E3
var IID_IImageEncodingProperties2 = syscall.GUID{0xC854A2DF, 0xC923, 0x469B,
	[8]byte{0xAC, 0x8E, 0x6A, 0x9F, 0x3C, 0x1C, 0xD9, 0xE3}}

type IImageEncodingProperties2Interface interface {
	win32.IInspectableInterface
	Copy() *IImageEncodingProperties
}

type IImageEncodingProperties2Vtbl struct {
	win32.IInspectableVtbl
	Copy uintptr
}

type IImageEncodingProperties2 struct {
	win32.IInspectable
}

func (this *IImageEncodingProperties2) Vtbl() *IImageEncodingProperties2Vtbl {
	return (*IImageEncodingProperties2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IImageEncodingProperties2) Copy() *IImageEncodingProperties {
	var _result *IImageEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Copy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 257C68DC-8B99-439E-AA59-913A36161297
var IID_IImageEncodingPropertiesStatics = syscall.GUID{0x257C68DC, 0x8B99, 0x439E,
	[8]byte{0xAA, 0x59, 0x91, 0x3A, 0x36, 0x16, 0x12, 0x97}}

type IImageEncodingPropertiesStaticsInterface interface {
	win32.IInspectableInterface
	CreateJpeg() *IImageEncodingProperties
	CreatePng() *IImageEncodingProperties
	CreateJpegXR() *IImageEncodingProperties
}

type IImageEncodingPropertiesStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateJpeg   uintptr
	CreatePng    uintptr
	CreateJpegXR uintptr
}

type IImageEncodingPropertiesStatics struct {
	win32.IInspectable
}

func (this *IImageEncodingPropertiesStatics) Vtbl() *IImageEncodingPropertiesStaticsVtbl {
	return (*IImageEncodingPropertiesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IImageEncodingPropertiesStatics) CreateJpeg() *IImageEncodingProperties {
	var _result *IImageEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateJpeg, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IImageEncodingPropertiesStatics) CreatePng() *IImageEncodingProperties {
	var _result *IImageEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreatePng, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IImageEncodingPropertiesStatics) CreateJpegXR() *IImageEncodingProperties {
	var _result *IImageEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateJpegXR, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F6C25B29-3824-46B0-956E-501329E1BE3C
var IID_IImageEncodingPropertiesStatics2 = syscall.GUID{0xF6C25B29, 0x3824, 0x46B0,
	[8]byte{0x95, 0x6E, 0x50, 0x13, 0x29, 0xE1, 0xBE, 0x3C}}

type IImageEncodingPropertiesStatics2Interface interface {
	win32.IInspectableInterface
	CreateUncompressed(format MediaPixelFormat) *IImageEncodingProperties
	CreateBmp() *IImageEncodingProperties
}

type IImageEncodingPropertiesStatics2Vtbl struct {
	win32.IInspectableVtbl
	CreateUncompressed uintptr
	CreateBmp          uintptr
}

type IImageEncodingPropertiesStatics2 struct {
	win32.IInspectable
}

func (this *IImageEncodingPropertiesStatics2) Vtbl() *IImageEncodingPropertiesStatics2Vtbl {
	return (*IImageEncodingPropertiesStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IImageEncodingPropertiesStatics2) CreateUncompressed(format MediaPixelFormat) *IImageEncodingProperties {
	var _result *IImageEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateUncompressed, uintptr(unsafe.Pointer(this)), uintptr(format), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IImageEncodingPropertiesStatics2) CreateBmp() *IImageEncodingProperties {
	var _result *IImageEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateBmp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 48F4814D-A2FF-48DC-8EA0-E90680663656
var IID_IImageEncodingPropertiesStatics3 = syscall.GUID{0x48F4814D, 0xA2FF, 0x48DC,
	[8]byte{0x8E, 0xA0, 0xE9, 0x06, 0x80, 0x66, 0x36, 0x56}}

type IImageEncodingPropertiesStatics3Interface interface {
	win32.IInspectableInterface
	CreateHeif() *IImageEncodingProperties
}

type IImageEncodingPropertiesStatics3Vtbl struct {
	win32.IInspectableVtbl
	CreateHeif uintptr
}

type IImageEncodingPropertiesStatics3 struct {
	win32.IInspectable
}

func (this *IImageEncodingPropertiesStatics3) Vtbl() *IImageEncodingPropertiesStatics3Vtbl {
	return (*IImageEncodingPropertiesStatics3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IImageEncodingPropertiesStatics3) CreateHeif() *IImageEncodingProperties {
	var _result *IImageEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateHeif, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E7DBF5A8-1DB9-4783-876B-3DFE12ACFDB3
var IID_IMediaEncodingProfile = syscall.GUID{0xE7DBF5A8, 0x1DB9, 0x4783,
	[8]byte{0x87, 0x6B, 0x3D, 0xFE, 0x12, 0xAC, 0xFD, 0xB3}}

type IMediaEncodingProfileInterface interface {
	win32.IInspectableInterface
	Put_Audio(value *IAudioEncodingProperties)
	Get_Audio() *IAudioEncodingProperties
	Put_Video(value *IVideoEncodingProperties)
	Get_Video() *IVideoEncodingProperties
	Put_Container(value *IContainerEncodingProperties)
	Get_Container() *IContainerEncodingProperties
}

type IMediaEncodingProfileVtbl struct {
	win32.IInspectableVtbl
	Put_Audio     uintptr
	Get_Audio     uintptr
	Put_Video     uintptr
	Get_Video     uintptr
	Put_Container uintptr
	Get_Container uintptr
}

type IMediaEncodingProfile struct {
	win32.IInspectable
}

func (this *IMediaEncodingProfile) Vtbl() *IMediaEncodingProfileVtbl {
	return (*IMediaEncodingProfileVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaEncodingProfile) Put_Audio(value *IAudioEncodingProperties) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Audio, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaEncodingProfile) Get_Audio() *IAudioEncodingProperties {
	var _result *IAudioEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Audio, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaEncodingProfile) Put_Video(value *IVideoEncodingProperties) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Video, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaEncodingProfile) Get_Video() *IVideoEncodingProperties {
	var _result *IVideoEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Video, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaEncodingProfile) Put_Container(value *IContainerEncodingProperties) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Container, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaEncodingProfile) Get_Container() *IContainerEncodingProperties {
	var _result *IContainerEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Container, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 349B3E0A-4035-488E-9877-85632865ED10
var IID_IMediaEncodingProfile2 = syscall.GUID{0x349B3E0A, 0x4035, 0x488E,
	[8]byte{0x98, 0x77, 0x85, 0x63, 0x28, 0x65, 0xED, 0x10}}

type IMediaEncodingProfile2Interface interface {
	win32.IInspectableInterface
	SetAudioTracks(value *IIterable[*IAudioStreamDescriptor])
	GetAudioTracks() *IVector[*IAudioStreamDescriptor]
	SetVideoTracks(value *IIterable[*IVideoStreamDescriptor])
	GetVideoTracks() *IVector[*IVideoStreamDescriptor]
}

type IMediaEncodingProfile2Vtbl struct {
	win32.IInspectableVtbl
	SetAudioTracks uintptr
	GetAudioTracks uintptr
	SetVideoTracks uintptr
	GetVideoTracks uintptr
}

type IMediaEncodingProfile2 struct {
	win32.IInspectable
}

func (this *IMediaEncodingProfile2) Vtbl() *IMediaEncodingProfile2Vtbl {
	return (*IMediaEncodingProfile2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaEncodingProfile2) SetAudioTracks(value *IIterable[*IAudioStreamDescriptor]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetAudioTracks, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaEncodingProfile2) GetAudioTracks() *IVector[*IAudioStreamDescriptor] {
	var _result *IVector[*IAudioStreamDescriptor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAudioTracks, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaEncodingProfile2) SetVideoTracks(value *IIterable[*IVideoStreamDescriptor]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetVideoTracks, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaEncodingProfile2) GetVideoTracks() *IVector[*IVideoStreamDescriptor] {
	var _result *IVector[*IVideoStreamDescriptor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetVideoTracks, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BA6EBE88-7570-4E69-ACCF-5611AD015F88
var IID_IMediaEncodingProfile3 = syscall.GUID{0xBA6EBE88, 0x7570, 0x4E69,
	[8]byte{0xAC, 0xCF, 0x56, 0x11, 0xAD, 0x01, 0x5F, 0x88}}

type IMediaEncodingProfile3Interface interface {
	win32.IInspectableInterface
	SetTimedMetadataTracks(value *IIterable[*IMediaStreamDescriptor])
	GetTimedMetadataTracks() *IVector[*IMediaStreamDescriptor]
}

type IMediaEncodingProfile3Vtbl struct {
	win32.IInspectableVtbl
	SetTimedMetadataTracks uintptr
	GetTimedMetadataTracks uintptr
}

type IMediaEncodingProfile3 struct {
	win32.IInspectable
}

func (this *IMediaEncodingProfile3) Vtbl() *IMediaEncodingProfile3Vtbl {
	return (*IMediaEncodingProfile3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaEncodingProfile3) SetTimedMetadataTracks(value *IIterable[*IMediaStreamDescriptor]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetTimedMetadataTracks, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IMediaEncodingProfile3) GetTimedMetadataTracks() *IVector[*IMediaStreamDescriptor] {
	var _result *IVector[*IMediaStreamDescriptor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetTimedMetadataTracks, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 197F352C-2EDE-4A45-A896-817A4854F8FE
var IID_IMediaEncodingProfileStatics = syscall.GUID{0x197F352C, 0x2EDE, 0x4A45,
	[8]byte{0xA8, 0x96, 0x81, 0x7A, 0x48, 0x54, 0xF8, 0xFE}}

type IMediaEncodingProfileStaticsInterface interface {
	win32.IInspectableInterface
	CreateM4a(quality AudioEncodingQuality) *IMediaEncodingProfile
	CreateMp3(quality AudioEncodingQuality) *IMediaEncodingProfile
	CreateWma(quality AudioEncodingQuality) *IMediaEncodingProfile
	CreateMp4(quality VideoEncodingQuality) *IMediaEncodingProfile
	CreateWmv(quality VideoEncodingQuality) *IMediaEncodingProfile
	CreateFromFileAsync(file *IStorageFile) *IAsyncOperation[*IMediaEncodingProfile]
	CreateFromStreamAsync(stream *IRandomAccessStream) *IAsyncOperation[*IMediaEncodingProfile]
}

type IMediaEncodingProfileStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateM4a             uintptr
	CreateMp3             uintptr
	CreateWma             uintptr
	CreateMp4             uintptr
	CreateWmv             uintptr
	CreateFromFileAsync   uintptr
	CreateFromStreamAsync uintptr
}

type IMediaEncodingProfileStatics struct {
	win32.IInspectable
}

func (this *IMediaEncodingProfileStatics) Vtbl() *IMediaEncodingProfileStaticsVtbl {
	return (*IMediaEncodingProfileStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaEncodingProfileStatics) CreateM4a(quality AudioEncodingQuality) *IMediaEncodingProfile {
	var _result *IMediaEncodingProfile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateM4a, uintptr(unsafe.Pointer(this)), uintptr(quality), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaEncodingProfileStatics) CreateMp3(quality AudioEncodingQuality) *IMediaEncodingProfile {
	var _result *IMediaEncodingProfile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateMp3, uintptr(unsafe.Pointer(this)), uintptr(quality), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaEncodingProfileStatics) CreateWma(quality AudioEncodingQuality) *IMediaEncodingProfile {
	var _result *IMediaEncodingProfile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWma, uintptr(unsafe.Pointer(this)), uintptr(quality), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaEncodingProfileStatics) CreateMp4(quality VideoEncodingQuality) *IMediaEncodingProfile {
	var _result *IMediaEncodingProfile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateMp4, uintptr(unsafe.Pointer(this)), uintptr(quality), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaEncodingProfileStatics) CreateWmv(quality VideoEncodingQuality) *IMediaEncodingProfile {
	var _result *IMediaEncodingProfile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWmv, uintptr(unsafe.Pointer(this)), uintptr(quality), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaEncodingProfileStatics) CreateFromFileAsync(file *IStorageFile) *IAsyncOperation[*IMediaEncodingProfile] {
	var _result *IAsyncOperation[*IMediaEncodingProfile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromFileAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaEncodingProfileStatics) CreateFromStreamAsync(stream *IRandomAccessStream) *IAsyncOperation[*IMediaEncodingProfile] {
	var _result *IAsyncOperation[*IMediaEncodingProfile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromStreamAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CE8DE74F-6AF4-4288-8FE2-79ADF1F79A43
var IID_IMediaEncodingProfileStatics2 = syscall.GUID{0xCE8DE74F, 0x6AF4, 0x4288,
	[8]byte{0x8F, 0xE2, 0x79, 0xAD, 0xF1, 0xF7, 0x9A, 0x43}}

type IMediaEncodingProfileStatics2Interface interface {
	win32.IInspectableInterface
	CreateWav(quality AudioEncodingQuality) *IMediaEncodingProfile
	CreateAvi(quality VideoEncodingQuality) *IMediaEncodingProfile
}

type IMediaEncodingProfileStatics2Vtbl struct {
	win32.IInspectableVtbl
	CreateWav uintptr
	CreateAvi uintptr
}

type IMediaEncodingProfileStatics2 struct {
	win32.IInspectable
}

func (this *IMediaEncodingProfileStatics2) Vtbl() *IMediaEncodingProfileStatics2Vtbl {
	return (*IMediaEncodingProfileStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaEncodingProfileStatics2) CreateWav(quality AudioEncodingQuality) *IMediaEncodingProfile {
	var _result *IMediaEncodingProfile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWav, uintptr(unsafe.Pointer(this)), uintptr(quality), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaEncodingProfileStatics2) CreateAvi(quality VideoEncodingQuality) *IMediaEncodingProfile {
	var _result *IMediaEncodingProfile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAvi, uintptr(unsafe.Pointer(this)), uintptr(quality), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 90DAC5AA-CF76-4294-A9ED-1A1420F51F6B
var IID_IMediaEncodingProfileStatics3 = syscall.GUID{0x90DAC5AA, 0xCF76, 0x4294,
	[8]byte{0xA9, 0xED, 0x1A, 0x14, 0x20, 0xF5, 0x1F, 0x6B}}

type IMediaEncodingProfileStatics3Interface interface {
	win32.IInspectableInterface
	CreateAlac(quality AudioEncodingQuality) *IMediaEncodingProfile
	CreateFlac(quality AudioEncodingQuality) *IMediaEncodingProfile
	CreateHevc(quality VideoEncodingQuality) *IMediaEncodingProfile
}

type IMediaEncodingProfileStatics3Vtbl struct {
	win32.IInspectableVtbl
	CreateAlac uintptr
	CreateFlac uintptr
	CreateHevc uintptr
}

type IMediaEncodingProfileStatics3 struct {
	win32.IInspectable
}

func (this *IMediaEncodingProfileStatics3) Vtbl() *IMediaEncodingProfileStatics3Vtbl {
	return (*IMediaEncodingProfileStatics3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaEncodingProfileStatics3) CreateAlac(quality AudioEncodingQuality) *IMediaEncodingProfile {
	var _result *IMediaEncodingProfile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAlac, uintptr(unsafe.Pointer(this)), uintptr(quality), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaEncodingProfileStatics3) CreateFlac(quality AudioEncodingQuality) *IMediaEncodingProfile {
	var _result *IMediaEncodingProfile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFlac, uintptr(unsafe.Pointer(this)), uintptr(quality), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaEncodingProfileStatics3) CreateHevc(quality VideoEncodingQuality) *IMediaEncodingProfile {
	var _result *IMediaEncodingProfile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateHevc, uintptr(unsafe.Pointer(this)), uintptr(quality), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B4002AF6-ACD4-4E5A-A24B-5D7498A8B8C4
var IID_IMediaEncodingProperties = syscall.GUID{0xB4002AF6, 0xACD4, 0x4E5A,
	[8]byte{0xA2, 0x4B, 0x5D, 0x74, 0x98, 0xA8, 0xB8, 0xC4}}

type IMediaEncodingPropertiesInterface interface {
	win32.IInspectableInterface
	Get_Properties() *IMap[syscall.GUID, interface{}]
	Get_Type() string
	Put_Subtype(value string)
	Get_Subtype() string
}

type IMediaEncodingPropertiesVtbl struct {
	win32.IInspectableVtbl
	Get_Properties uintptr
	Get_Type       uintptr
	Put_Subtype    uintptr
	Get_Subtype    uintptr
}

type IMediaEncodingProperties struct {
	win32.IInspectable
}

func (this *IMediaEncodingProperties) Vtbl() *IMediaEncodingPropertiesVtbl {
	return (*IMediaEncodingPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaEncodingProperties) Get_Properties() *IMap[syscall.GUID, interface{}] {
	var _result *IMap[syscall.GUID, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaEncodingProperties) Get_Type() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Type, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingProperties) Put_Subtype(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Subtype, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMediaEncodingProperties) Get_Subtype() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Subtype, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 37B6580E-A171-4464-BA5A-53189E48C1C8
var IID_IMediaEncodingSubtypesStatics = syscall.GUID{0x37B6580E, 0xA171, 0x4464,
	[8]byte{0xBA, 0x5A, 0x53, 0x18, 0x9E, 0x48, 0xC1, 0xC8}}

type IMediaEncodingSubtypesStaticsInterface interface {
	win32.IInspectableInterface
	Get_Aac() string
	Get_AacAdts() string
	Get_Ac3() string
	Get_AmrNb() string
	Get_AmrWb() string
	Get_Argb32() string
	Get_Asf() string
	Get_Avi() string
	Get_Bgra8() string
	Get_Bmp() string
	Get_Eac3() string
	Get_Float() string
	Get_Gif() string
	Get_H263() string
	Get_H264() string
	Get_H264Es() string
	Get_Hevc() string
	Get_HevcEs() string
	Get_Iyuv() string
	Get_Jpeg() string
	Get_JpegXr() string
	Get_Mjpg() string
	Get_Mpeg() string
	Get_Mpeg1() string
	Get_Mpeg2() string
	Get_Mp3() string
	Get_Mpeg4() string
	Get_Nv12() string
	Get_Pcm() string
	Get_Png() string
	Get_Rgb24() string
	Get_Rgb32() string
	Get_Tiff() string
	Get_Wave() string
	Get_Wma8() string
	Get_Wma9() string
	Get_Wmv3() string
	Get_Wvc1() string
	Get_Yuy2() string
	Get_Yv12() string
}

type IMediaEncodingSubtypesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Aac     uintptr
	Get_AacAdts uintptr
	Get_Ac3     uintptr
	Get_AmrNb   uintptr
	Get_AmrWb   uintptr
	Get_Argb32  uintptr
	Get_Asf     uintptr
	Get_Avi     uintptr
	Get_Bgra8   uintptr
	Get_Bmp     uintptr
	Get_Eac3    uintptr
	Get_Float   uintptr
	Get_Gif     uintptr
	Get_H263    uintptr
	Get_H264    uintptr
	Get_H264Es  uintptr
	Get_Hevc    uintptr
	Get_HevcEs  uintptr
	Get_Iyuv    uintptr
	Get_Jpeg    uintptr
	Get_JpegXr  uintptr
	Get_Mjpg    uintptr
	Get_Mpeg    uintptr
	Get_Mpeg1   uintptr
	Get_Mpeg2   uintptr
	Get_Mp3     uintptr
	Get_Mpeg4   uintptr
	Get_Nv12    uintptr
	Get_Pcm     uintptr
	Get_Png     uintptr
	Get_Rgb24   uintptr
	Get_Rgb32   uintptr
	Get_Tiff    uintptr
	Get_Wave    uintptr
	Get_Wma8    uintptr
	Get_Wma9    uintptr
	Get_Wmv3    uintptr
	Get_Wvc1    uintptr
	Get_Yuy2    uintptr
	Get_Yv12    uintptr
}

type IMediaEncodingSubtypesStatics struct {
	win32.IInspectable
}

func (this *IMediaEncodingSubtypesStatics) Vtbl() *IMediaEncodingSubtypesStaticsVtbl {
	return (*IMediaEncodingSubtypesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaEncodingSubtypesStatics) Get_Aac() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Aac, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_AacAdts() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AacAdts, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Ac3() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ac3, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_AmrNb() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AmrNb, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_AmrWb() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AmrWb, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Argb32() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Argb32, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Asf() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Asf, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Avi() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Avi, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Bgra8() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bgra8, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Bmp() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bmp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Eac3() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Eac3, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Float() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Float, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Gif() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Gif, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_H263() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_H263, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_H264() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_H264, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_H264Es() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_H264Es, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Hevc() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Hevc, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_HevcEs() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HevcEs, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Iyuv() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Iyuv, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Jpeg() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Jpeg, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_JpegXr() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_JpegXr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Mjpg() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mjpg, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Mpeg() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mpeg, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Mpeg1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mpeg1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Mpeg2() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mpeg2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Mp3() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mp3, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Mpeg4() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mpeg4, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Nv12() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Nv12, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Pcm() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Pcm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Png() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Png, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Rgb24() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Rgb24, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Rgb32() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Rgb32, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Tiff() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Tiff, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Wave() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Wave, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Wma8() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Wma8, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Wma9() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Wma9, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Wmv3() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Wmv3, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Wvc1() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Wvc1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Yuy2() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Yuy2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics) Get_Yv12() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Yv12, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 4B7CD23D-42FF-4D33-8531-0626BEE4B52D
var IID_IMediaEncodingSubtypesStatics2 = syscall.GUID{0x4B7CD23D, 0x42FF, 0x4D33,
	[8]byte{0x85, 0x31, 0x06, 0x26, 0xBE, 0xE4, 0xB5, 0x2D}}

type IMediaEncodingSubtypesStatics2Interface interface {
	win32.IInspectableInterface
	Get_Vp9() string
	Get_L8() string
	Get_L16() string
	Get_D16() string
}

type IMediaEncodingSubtypesStatics2Vtbl struct {
	win32.IInspectableVtbl
	Get_Vp9 uintptr
	Get_L8  uintptr
	Get_L16 uintptr
	Get_D16 uintptr
}

type IMediaEncodingSubtypesStatics2 struct {
	win32.IInspectable
}

func (this *IMediaEncodingSubtypesStatics2) Vtbl() *IMediaEncodingSubtypesStatics2Vtbl {
	return (*IMediaEncodingSubtypesStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaEncodingSubtypesStatics2) Get_Vp9() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Vp9, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics2) Get_L8() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_L8, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics2) Get_L16() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_L16, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics2) Get_D16() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_D16, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// BA2414E4-883D-464E-A44F-097DA08EF7FF
var IID_IMediaEncodingSubtypesStatics3 = syscall.GUID{0xBA2414E4, 0x883D, 0x464E,
	[8]byte{0xA4, 0x4F, 0x09, 0x7D, 0xA0, 0x8E, 0xF7, 0xFF}}

type IMediaEncodingSubtypesStatics3Interface interface {
	win32.IInspectableInterface
	Get_Alac() string
	Get_Flac() string
}

type IMediaEncodingSubtypesStatics3Vtbl struct {
	win32.IInspectableVtbl
	Get_Alac uintptr
	Get_Flac uintptr
}

type IMediaEncodingSubtypesStatics3 struct {
	win32.IInspectable
}

func (this *IMediaEncodingSubtypesStatics3) Vtbl() *IMediaEncodingSubtypesStatics3Vtbl {
	return (*IMediaEncodingSubtypesStatics3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaEncodingSubtypesStatics3) Get_Alac() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Alac, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics3) Get_Flac() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Flac, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// DDECE58A-3949-4644-8A2C-59EF02C642FA
var IID_IMediaEncodingSubtypesStatics4 = syscall.GUID{0xDDECE58A, 0x3949, 0x4644,
	[8]byte{0x8A, 0x2C, 0x59, 0xEF, 0x02, 0xC6, 0x42, 0xFA}}

type IMediaEncodingSubtypesStatics4Interface interface {
	win32.IInspectableInterface
	Get_P010() string
}

type IMediaEncodingSubtypesStatics4Vtbl struct {
	win32.IInspectableVtbl
	Get_P010 uintptr
}

type IMediaEncodingSubtypesStatics4 struct {
	win32.IInspectable
}

func (this *IMediaEncodingSubtypesStatics4) Vtbl() *IMediaEncodingSubtypesStatics4Vtbl {
	return (*IMediaEncodingSubtypesStatics4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaEncodingSubtypesStatics4) Get_P010() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_P010, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 5AD4A007-FFCE-4760-9828-5D0C99637E6A
var IID_IMediaEncodingSubtypesStatics5 = syscall.GUID{0x5AD4A007, 0xFFCE, 0x4760,
	[8]byte{0x98, 0x28, 0x5D, 0x0C, 0x99, 0x63, 0x7E, 0x6A}}

type IMediaEncodingSubtypesStatics5Interface interface {
	win32.IInspectableInterface
	Get_Heif() string
}

type IMediaEncodingSubtypesStatics5Vtbl struct {
	win32.IInspectableVtbl
	Get_Heif uintptr
}

type IMediaEncodingSubtypesStatics5 struct {
	win32.IInspectable
}

func (this *IMediaEncodingSubtypesStatics5) Vtbl() *IMediaEncodingSubtypesStatics5Vtbl {
	return (*IMediaEncodingSubtypesStatics5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaEncodingSubtypesStatics5) Get_Heif() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Heif, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// A1252973-A984-5912-93BB-54E7E569E053
var IID_IMediaEncodingSubtypesStatics6 = syscall.GUID{0xA1252973, 0xA984, 0x5912,
	[8]byte{0x93, 0xBB, 0x54, 0xE7, 0xE5, 0x69, 0xE0, 0x53}}

type IMediaEncodingSubtypesStatics6Interface interface {
	win32.IInspectableInterface
	Get_Pgs() string
	Get_Srt() string
	Get_Ssa() string
	Get_VobSub() string
}

type IMediaEncodingSubtypesStatics6Vtbl struct {
	win32.IInspectableVtbl
	Get_Pgs    uintptr
	Get_Srt    uintptr
	Get_Ssa    uintptr
	Get_VobSub uintptr
}

type IMediaEncodingSubtypesStatics6 struct {
	win32.IInspectable
}

func (this *IMediaEncodingSubtypesStatics6) Vtbl() *IMediaEncodingSubtypesStatics6Vtbl {
	return (*IMediaEncodingSubtypesStatics6Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaEncodingSubtypesStatics6) Get_Pgs() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Pgs, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics6) Get_Srt() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Srt, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics6) Get_Ssa() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Ssa, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaEncodingSubtypesStatics6) Get_VobSub() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VobSub, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// D2D0FEE5-8929-401D-AC78-7D357E378163
var IID_IMediaRatio = syscall.GUID{0xD2D0FEE5, 0x8929, 0x401D,
	[8]byte{0xAC, 0x78, 0x7D, 0x35, 0x7E, 0x37, 0x81, 0x63}}

type IMediaRatioInterface interface {
	win32.IInspectableInterface
	Put_Numerator(value uint32)
	Get_Numerator() uint32
	Put_Denominator(value uint32)
	Get_Denominator() uint32
}

type IMediaRatioVtbl struct {
	win32.IInspectableVtbl
	Put_Numerator   uintptr
	Get_Numerator   uintptr
	Put_Denominator uintptr
	Get_Denominator uintptr
}

type IMediaRatio struct {
	win32.IInspectable
}

func (this *IMediaRatio) Vtbl() *IMediaRatioVtbl {
	return (*IMediaRatioVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaRatio) Put_Numerator(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Numerator, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMediaRatio) Get_Numerator() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Numerator, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaRatio) Put_Denominator(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Denominator, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMediaRatio) Get_Denominator() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Denominator, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A461FF85-E57A-4128-9B21-D5331B04235C
var IID_IMpeg2ProfileIdsStatics = syscall.GUID{0xA461FF85, 0xE57A, 0x4128,
	[8]byte{0x9B, 0x21, 0xD5, 0x33, 0x1B, 0x04, 0x23, 0x5C}}

type IMpeg2ProfileIdsStaticsInterface interface {
	win32.IInspectableInterface
	Get_Simple() int32
	Get_Main() int32
	Get_SignalNoiseRatioScalable() int32
	Get_SpatiallyScalable() int32
	Get_High() int32
}

type IMpeg2ProfileIdsStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Simple                   uintptr
	Get_Main                     uintptr
	Get_SignalNoiseRatioScalable uintptr
	Get_SpatiallyScalable        uintptr
	Get_High                     uintptr
}

type IMpeg2ProfileIdsStatics struct {
	win32.IInspectable
}

func (this *IMpeg2ProfileIdsStatics) Vtbl() *IMpeg2ProfileIdsStaticsVtbl {
	return (*IMpeg2ProfileIdsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMpeg2ProfileIdsStatics) Get_Simple() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Simple, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMpeg2ProfileIdsStatics) Get_Main() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Main, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMpeg2ProfileIdsStatics) Get_SignalNoiseRatioScalable() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SignalNoiseRatioScalable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMpeg2ProfileIdsStatics) Get_SpatiallyScalable() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SpatiallyScalable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMpeg2ProfileIdsStatics) Get_High() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_High, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 51CD30D3-D690-4CFA-97F4-4A398E9DB420
var IID_ITimedMetadataEncodingProperties = syscall.GUID{0x51CD30D3, 0xD690, 0x4CFA,
	[8]byte{0x97, 0xF4, 0x4A, 0x39, 0x8E, 0x9D, 0xB4, 0x20}}

type ITimedMetadataEncodingPropertiesInterface interface {
	win32.IInspectableInterface
	SetFormatUserData(valueLength uint32, value *byte)
	GetFormatUserData(valueLength uint32, value *byte)
	Copy() *IMediaEncodingProperties
}

type ITimedMetadataEncodingPropertiesVtbl struct {
	win32.IInspectableVtbl
	SetFormatUserData uintptr
	GetFormatUserData uintptr
	Copy              uintptr
}

type ITimedMetadataEncodingProperties struct {
	win32.IInspectable
}

func (this *ITimedMetadataEncodingProperties) Vtbl() *ITimedMetadataEncodingPropertiesVtbl {
	return (*ITimedMetadataEncodingPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimedMetadataEncodingProperties) SetFormatUserData(valueLength uint32, value *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetFormatUserData, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ITimedMetadataEncodingProperties) GetFormatUserData(valueLength uint32, value *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFormatUserData, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ITimedMetadataEncodingProperties) Copy() *IMediaEncodingProperties {
	var _result *IMediaEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Copy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6629BB67-6E55-5643-89A0-7A7E8D85B52C
var IID_ITimedMetadataEncodingPropertiesStatics = syscall.GUID{0x6629BB67, 0x6E55, 0x5643,
	[8]byte{0x89, 0xA0, 0x7A, 0x7E, 0x8D, 0x85, 0xB5, 0x2C}}

type ITimedMetadataEncodingPropertiesStaticsInterface interface {
	win32.IInspectableInterface
	CreatePgs() *IMediaEncodingProperties
	CreateSrt() *IMediaEncodingProperties
	CreateSsa(formatUserDataLength uint32, formatUserData *byte) *IMediaEncodingProperties
	CreateVobSub(formatUserDataLength uint32, formatUserData *byte) *IMediaEncodingProperties
}

type ITimedMetadataEncodingPropertiesStaticsVtbl struct {
	win32.IInspectableVtbl
	CreatePgs    uintptr
	CreateSrt    uintptr
	CreateSsa    uintptr
	CreateVobSub uintptr
}

type ITimedMetadataEncodingPropertiesStatics struct {
	win32.IInspectable
}

func (this *ITimedMetadataEncodingPropertiesStatics) Vtbl() *ITimedMetadataEncodingPropertiesStaticsVtbl {
	return (*ITimedMetadataEncodingPropertiesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITimedMetadataEncodingPropertiesStatics) CreatePgs() *IMediaEncodingProperties {
	var _result *IMediaEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreatePgs, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITimedMetadataEncodingPropertiesStatics) CreateSrt() *IMediaEncodingProperties {
	var _result *IMediaEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSrt, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITimedMetadataEncodingPropertiesStatics) CreateSsa(formatUserDataLength uint32, formatUserData *byte) *IMediaEncodingProperties {
	var _result *IMediaEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateSsa, uintptr(unsafe.Pointer(this)), uintptr(formatUserDataLength), uintptr(unsafe.Pointer(formatUserData)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITimedMetadataEncodingPropertiesStatics) CreateVobSub(formatUserDataLength uint32, formatUserData *byte) *IMediaEncodingProperties {
	var _result *IMediaEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateVobSub, uintptr(unsafe.Pointer(this)), uintptr(formatUserDataLength), uintptr(unsafe.Pointer(formatUserData)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 76EE6C9A-37C2-4F2A-880A-1282BBB4373D
var IID_IVideoEncodingProperties = syscall.GUID{0x76EE6C9A, 0x37C2, 0x4F2A,
	[8]byte{0x88, 0x0A, 0x12, 0x82, 0xBB, 0xB4, 0x37, 0x3D}}

type IVideoEncodingPropertiesInterface interface {
	win32.IInspectableInterface
	Put_Bitrate(value uint32)
	Get_Bitrate() uint32
	Put_Width(value uint32)
	Get_Width() uint32
	Put_Height(value uint32)
	Get_Height() uint32
	Get_FrameRate() *IMediaRatio
	Get_PixelAspectRatio() *IMediaRatio
}

type IVideoEncodingPropertiesVtbl struct {
	win32.IInspectableVtbl
	Put_Bitrate          uintptr
	Get_Bitrate          uintptr
	Put_Width            uintptr
	Get_Width            uintptr
	Put_Height           uintptr
	Get_Height           uintptr
	Get_FrameRate        uintptr
	Get_PixelAspectRatio uintptr
}

type IVideoEncodingProperties struct {
	win32.IInspectable
}

func (this *IVideoEncodingProperties) Vtbl() *IVideoEncodingPropertiesVtbl {
	return (*IVideoEncodingPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoEncodingProperties) Put_Bitrate(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Bitrate, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVideoEncodingProperties) Get_Bitrate() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bitrate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoEncodingProperties) Put_Width(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Width, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVideoEncodingProperties) Get_Width() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Width, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoEncodingProperties) Put_Height(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Height, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVideoEncodingProperties) Get_Height() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Height, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoEncodingProperties) Get_FrameRate() *IMediaRatio {
	var _result *IMediaRatio
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoEncodingProperties) Get_PixelAspectRatio() *IMediaRatio {
	var _result *IMediaRatio
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PixelAspectRatio, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F743A1EF-D465-4290-A94B-EF0F1528F8E3
var IID_IVideoEncodingProperties2 = syscall.GUID{0xF743A1EF, 0xD465, 0x4290,
	[8]byte{0xA9, 0x4B, 0xEF, 0x0F, 0x15, 0x28, 0xF8, 0xE3}}

type IVideoEncodingProperties2Interface interface {
	win32.IInspectableInterface
	SetFormatUserData(valueLength uint32, value *byte)
	GetFormatUserData(valueLength uint32, value *byte)
	Put_ProfileId(value int32)
	Get_ProfileId() int32
}

type IVideoEncodingProperties2Vtbl struct {
	win32.IInspectableVtbl
	SetFormatUserData uintptr
	GetFormatUserData uintptr
	Put_ProfileId     uintptr
	Get_ProfileId     uintptr
}

type IVideoEncodingProperties2 struct {
	win32.IInspectable
}

func (this *IVideoEncodingProperties2) Vtbl() *IVideoEncodingProperties2Vtbl {
	return (*IVideoEncodingProperties2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoEncodingProperties2) SetFormatUserData(valueLength uint32, value *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetFormatUserData, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IVideoEncodingProperties2) GetFormatUserData(valueLength uint32, value *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFormatUserData, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IVideoEncodingProperties2) Put_ProfileId(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ProfileId, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVideoEncodingProperties2) Get_ProfileId() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProfileId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 386BCDC4-873A-479F-B3EB-56C1FCBEC6D7
var IID_IVideoEncodingProperties3 = syscall.GUID{0x386BCDC4, 0x873A, 0x479F,
	[8]byte{0xB3, 0xEB, 0x56, 0xC1, 0xFC, 0xBE, 0xC6, 0xD7}}

type IVideoEncodingProperties3Interface interface {
	win32.IInspectableInterface
	Get_StereoscopicVideoPackingMode() StereoscopicVideoPackingMode
}

type IVideoEncodingProperties3Vtbl struct {
	win32.IInspectableVtbl
	Get_StereoscopicVideoPackingMode uintptr
}

type IVideoEncodingProperties3 struct {
	win32.IInspectable
}

func (this *IVideoEncodingProperties3) Vtbl() *IVideoEncodingProperties3Vtbl {
	return (*IVideoEncodingProperties3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoEncodingProperties3) Get_StereoscopicVideoPackingMode() StereoscopicVideoPackingMode {
	var _result StereoscopicVideoPackingMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StereoscopicVideoPackingMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 724EF014-C10C-40F2-9D72-3EE13B45FA8E
var IID_IVideoEncodingProperties4 = syscall.GUID{0x724EF014, 0xC10C, 0x40F2,
	[8]byte{0x9D, 0x72, 0x3E, 0xE1, 0x3B, 0x45, 0xFA, 0x8E}}

type IVideoEncodingProperties4Interface interface {
	win32.IInspectableInterface
	Get_SphericalVideoFrameFormat() SphericalVideoFrameFormat
}

type IVideoEncodingProperties4Vtbl struct {
	win32.IInspectableVtbl
	Get_SphericalVideoFrameFormat uintptr
}

type IVideoEncodingProperties4 struct {
	win32.IInspectable
}

func (this *IVideoEncodingProperties4) Vtbl() *IVideoEncodingProperties4Vtbl {
	return (*IVideoEncodingProperties4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoEncodingProperties4) Get_SphericalVideoFrameFormat() SphericalVideoFrameFormat {
	var _result SphericalVideoFrameFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SphericalVideoFrameFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4959080F-272F-4ECE-A4DF-C0CCDB33D840
var IID_IVideoEncodingProperties5 = syscall.GUID{0x4959080F, 0x272F, 0x4ECE,
	[8]byte{0xA4, 0xDF, 0xC0, 0xCC, 0xDB, 0x33, 0xD8, 0x40}}

type IVideoEncodingProperties5Interface interface {
	win32.IInspectableInterface
	Copy() *IVideoEncodingProperties
}

type IVideoEncodingProperties5Vtbl struct {
	win32.IInspectableVtbl
	Copy uintptr
}

type IVideoEncodingProperties5 struct {
	win32.IInspectable
}

func (this *IVideoEncodingProperties5) Vtbl() *IVideoEncodingProperties5Vtbl {
	return (*IVideoEncodingProperties5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoEncodingProperties5) Copy() *IVideoEncodingProperties {
	var _result *IVideoEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Copy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3CE14D44-1DC5-43DB-9F38-EBEBF90152CB
var IID_IVideoEncodingPropertiesStatics = syscall.GUID{0x3CE14D44, 0x1DC5, 0x43DB,
	[8]byte{0x9F, 0x38, 0xEB, 0xEB, 0xF9, 0x01, 0x52, 0xCB}}

type IVideoEncodingPropertiesStaticsInterface interface {
	win32.IInspectableInterface
	CreateH264() *IVideoEncodingProperties
	CreateMpeg2() *IVideoEncodingProperties
	CreateUncompressed(subtype string, width uint32, height uint32) *IVideoEncodingProperties
}

type IVideoEncodingPropertiesStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateH264         uintptr
	CreateMpeg2        uintptr
	CreateUncompressed uintptr
}

type IVideoEncodingPropertiesStatics struct {
	win32.IInspectable
}

func (this *IVideoEncodingPropertiesStatics) Vtbl() *IVideoEncodingPropertiesStaticsVtbl {
	return (*IVideoEncodingPropertiesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoEncodingPropertiesStatics) CreateH264() *IVideoEncodingProperties {
	var _result *IVideoEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateH264, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoEncodingPropertiesStatics) CreateMpeg2() *IVideoEncodingProperties {
	var _result *IVideoEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateMpeg2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoEncodingPropertiesStatics) CreateUncompressed(subtype string, width uint32, height uint32) *IVideoEncodingProperties {
	var _result *IVideoEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateUncompressed, uintptr(unsafe.Pointer(this)), NewHStr(subtype).Ptr, uintptr(width), uintptr(height), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CF1EBD5D-49FE-4D00-B59A-CFA4DFC51944
var IID_IVideoEncodingPropertiesStatics2 = syscall.GUID{0xCF1EBD5D, 0x49FE, 0x4D00,
	[8]byte{0xB5, 0x9A, 0xCF, 0xA4, 0xDF, 0xC5, 0x19, 0x44}}

type IVideoEncodingPropertiesStatics2Interface interface {
	win32.IInspectableInterface
	CreateHevc() *IVideoEncodingProperties
}

type IVideoEncodingPropertiesStatics2Vtbl struct {
	win32.IInspectableVtbl
	CreateHevc uintptr
}

type IVideoEncodingPropertiesStatics2 struct {
	win32.IInspectable
}

func (this *IVideoEncodingPropertiesStatics2) Vtbl() *IVideoEncodingPropertiesStatics2Vtbl {
	return (*IVideoEncodingPropertiesStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoEncodingPropertiesStatics2) CreateHevc() *IVideoEncodingProperties {
	var _result *IVideoEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateHevc, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type AudioEncodingProperties struct {
	RtClass
	*IAudioEncodingProperties
}

func NewAudioEncodingProperties() *AudioEncodingProperties {
	hs := NewHStr("Windows.Media.MediaProperties.AudioEncodingProperties")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &AudioEncodingProperties{
		RtClass:                  RtClass{PInspect: p},
		IAudioEncodingProperties: (*IAudioEncodingProperties)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewIAudioEncodingPropertiesStatics() *IAudioEncodingPropertiesStatics {
	var p *IAudioEncodingPropertiesStatics
	hs := NewHStr("Windows.Media.MediaProperties.AudioEncodingProperties")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAudioEncodingPropertiesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIAudioEncodingPropertiesStatics2() *IAudioEncodingPropertiesStatics2 {
	var p *IAudioEncodingPropertiesStatics2
	hs := NewHStr("Windows.Media.MediaProperties.AudioEncodingProperties")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAudioEncodingPropertiesStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type ContainerEncodingProperties struct {
	RtClass
	*IContainerEncodingProperties
}

func NewContainerEncodingProperties() *ContainerEncodingProperties {
	hs := NewHStr("Windows.Media.MediaProperties.ContainerEncodingProperties")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &ContainerEncodingProperties{
		RtClass:                      RtClass{PInspect: p},
		IContainerEncodingProperties: (*IContainerEncodingProperties)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type ImageEncodingProperties struct {
	RtClass
	*IImageEncodingProperties
}

func NewImageEncodingProperties() *ImageEncodingProperties {
	hs := NewHStr("Windows.Media.MediaProperties.ImageEncodingProperties")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &ImageEncodingProperties{
		RtClass:                  RtClass{PInspect: p},
		IImageEncodingProperties: (*IImageEncodingProperties)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewIImageEncodingPropertiesStatics3() *IImageEncodingPropertiesStatics3 {
	var p *IImageEncodingPropertiesStatics3
	hs := NewHStr("Windows.Media.MediaProperties.ImageEncodingProperties")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IImageEncodingPropertiesStatics3, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIImageEncodingPropertiesStatics() *IImageEncodingPropertiesStatics {
	var p *IImageEncodingPropertiesStatics
	hs := NewHStr("Windows.Media.MediaProperties.ImageEncodingProperties")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IImageEncodingPropertiesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIImageEncodingPropertiesStatics2() *IImageEncodingPropertiesStatics2 {
	var p *IImageEncodingPropertiesStatics2
	hs := NewHStr("Windows.Media.MediaProperties.ImageEncodingProperties")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IImageEncodingPropertiesStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type MediaEncodingProfile struct {
	RtClass
	*IMediaEncodingProfile
}

func NewMediaEncodingProfile() *MediaEncodingProfile {
	hs := NewHStr("Windows.Media.MediaProperties.MediaEncodingProfile")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &MediaEncodingProfile{
		RtClass:               RtClass{PInspect: p},
		IMediaEncodingProfile: (*IMediaEncodingProfile)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewIMediaEncodingProfileStatics3() *IMediaEncodingProfileStatics3 {
	var p *IMediaEncodingProfileStatics3
	hs := NewHStr("Windows.Media.MediaProperties.MediaEncodingProfile")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMediaEncodingProfileStatics3, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIMediaEncodingProfileStatics() *IMediaEncodingProfileStatics {
	var p *IMediaEncodingProfileStatics
	hs := NewHStr("Windows.Media.MediaProperties.MediaEncodingProfile")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMediaEncodingProfileStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIMediaEncodingProfileStatics2() *IMediaEncodingProfileStatics2 {
	var p *IMediaEncodingProfileStatics2
	hs := NewHStr("Windows.Media.MediaProperties.MediaEncodingProfile")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMediaEncodingProfileStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type MediaPropertySet struct {
	RtClass
	*IMap[syscall.GUID, interface{}]
}

func NewMediaPropertySet() *MediaPropertySet {
	hs := NewHStr("Windows.Media.MediaProperties.MediaPropertySet")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &MediaPropertySet{
		RtClass: RtClass{PInspect: p},
		IMap:    (*IMap[syscall.GUID, interface{}])(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type MediaRatio struct {
	RtClass
	*IMediaRatio
}

type VideoEncodingProperties struct {
	RtClass
	*IVideoEncodingProperties
}

func NewVideoEncodingProperties() *VideoEncodingProperties {
	hs := NewHStr("Windows.Media.MediaProperties.VideoEncodingProperties")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &VideoEncodingProperties{
		RtClass:                  RtClass{PInspect: p},
		IVideoEncodingProperties: (*IVideoEncodingProperties)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewIVideoEncodingPropertiesStatics2() *IVideoEncodingPropertiesStatics2 {
	var p *IVideoEncodingPropertiesStatics2
	hs := NewHStr("Windows.Media.MediaProperties.VideoEncodingProperties")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IVideoEncodingPropertiesStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIVideoEncodingPropertiesStatics() *IVideoEncodingPropertiesStatics {
	var p *IVideoEncodingPropertiesStatics
	hs := NewHStr("Windows.Media.MediaProperties.VideoEncodingProperties")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IVideoEncodingPropertiesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
