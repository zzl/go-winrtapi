package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type MediaFrameReaderAcquisitionMode int32

const (
	MediaFrameReaderAcquisitionMode_Realtime MediaFrameReaderAcquisitionMode = 0
	MediaFrameReaderAcquisitionMode_Buffered MediaFrameReaderAcquisitionMode = 1
)

// enum
type MediaFrameReaderStartStatus int32

const (
	MediaFrameReaderStartStatus_Success                      MediaFrameReaderStartStatus = 0
	MediaFrameReaderStartStatus_UnknownFailure               MediaFrameReaderStartStatus = 1
	MediaFrameReaderStartStatus_DeviceNotAvailable           MediaFrameReaderStartStatus = 2
	MediaFrameReaderStartStatus_OutputFormatNotSupported     MediaFrameReaderStartStatus = 3
	MediaFrameReaderStartStatus_ExclusiveControlNotAvailable MediaFrameReaderStartStatus = 4
)

// enum
type MediaFrameSourceGetPropertyStatus int32

const (
	MediaFrameSourceGetPropertyStatus_Success                      MediaFrameSourceGetPropertyStatus = 0
	MediaFrameSourceGetPropertyStatus_UnknownFailure               MediaFrameSourceGetPropertyStatus = 1
	MediaFrameSourceGetPropertyStatus_NotSupported                 MediaFrameSourceGetPropertyStatus = 2
	MediaFrameSourceGetPropertyStatus_DeviceNotAvailable           MediaFrameSourceGetPropertyStatus = 3
	MediaFrameSourceGetPropertyStatus_MaxPropertyValueSizeTooSmall MediaFrameSourceGetPropertyStatus = 4
	MediaFrameSourceGetPropertyStatus_MaxPropertyValueSizeRequired MediaFrameSourceGetPropertyStatus = 5
)

// enum
type MediaFrameSourceKind int32

const (
	MediaFrameSourceKind_Custom   MediaFrameSourceKind = 0
	MediaFrameSourceKind_Color    MediaFrameSourceKind = 1
	MediaFrameSourceKind_Infrared MediaFrameSourceKind = 2
	MediaFrameSourceKind_Depth    MediaFrameSourceKind = 3
	MediaFrameSourceKind_Audio    MediaFrameSourceKind = 4
	MediaFrameSourceKind_Image    MediaFrameSourceKind = 5
	MediaFrameSourceKind_Metadata MediaFrameSourceKind = 6
)

// enum
type MediaFrameSourceSetPropertyStatus int32

const (
	MediaFrameSourceSetPropertyStatus_Success            MediaFrameSourceSetPropertyStatus = 0
	MediaFrameSourceSetPropertyStatus_UnknownFailure     MediaFrameSourceSetPropertyStatus = 1
	MediaFrameSourceSetPropertyStatus_NotSupported       MediaFrameSourceSetPropertyStatus = 2
	MediaFrameSourceSetPropertyStatus_InvalidValue       MediaFrameSourceSetPropertyStatus = 3
	MediaFrameSourceSetPropertyStatus_DeviceNotAvailable MediaFrameSourceSetPropertyStatus = 4
	MediaFrameSourceSetPropertyStatus_NotInControl       MediaFrameSourceSetPropertyStatus = 5
)

// enum
type MultiSourceMediaFrameReaderStartStatus int32

const (
	MultiSourceMediaFrameReaderStartStatus_Success               MultiSourceMediaFrameReaderStartStatus = 0
	MultiSourceMediaFrameReaderStartStatus_NotSupported          MultiSourceMediaFrameReaderStartStatus = 1
	MultiSourceMediaFrameReaderStartStatus_InsufficientResources MultiSourceMediaFrameReaderStartStatus = 2
	MultiSourceMediaFrameReaderStartStatus_DeviceNotAvailable    MultiSourceMediaFrameReaderStartStatus = 3
	MultiSourceMediaFrameReaderStartStatus_UnknownFailure        MultiSourceMediaFrameReaderStartStatus = 4
)

// interfaces

// A3A9FEFF-8021-441B-9A46-E7F0137B7981
var IID_IAudioMediaFrame = syscall.GUID{0xA3A9FEFF, 0x8021, 0x441B,
	[8]byte{0x9A, 0x46, 0xE7, 0xF0, 0x13, 0x7B, 0x79, 0x81}}

type IAudioMediaFrameInterface interface {
	win32.IInspectableInterface
	Get_FrameReference() *IMediaFrameReference
	Get_AudioEncodingProperties() *IAudioEncodingProperties
	GetAudioFrame() *IAudioFrame
}

type IAudioMediaFrameVtbl struct {
	win32.IInspectableVtbl
	Get_FrameReference          uintptr
	Get_AudioEncodingProperties uintptr
	GetAudioFrame               uintptr
}

type IAudioMediaFrame struct {
	win32.IInspectable
}

func (this *IAudioMediaFrame) Vtbl() *IAudioMediaFrameVtbl {
	return (*IAudioMediaFrameVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioMediaFrame) Get_FrameReference() *IMediaFrameReference {
	var _result *IMediaFrameReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameReference, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioMediaFrame) Get_AudioEncodingProperties() *IAudioEncodingProperties {
	var _result *IAudioEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioEncodingProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioMediaFrame) GetAudioFrame() *IAudioFrame {
	var _result *IAudioFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAudioFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B5B153C7-9B84-4062-B79C-A365B2596854
var IID_IBufferMediaFrame = syscall.GUID{0xB5B153C7, 0x9B84, 0x4062,
	[8]byte{0xB7, 0x9C, 0xA3, 0x65, 0xB2, 0x59, 0x68, 0x54}}

type IBufferMediaFrameInterface interface {
	win32.IInspectableInterface
	Get_FrameReference() *IMediaFrameReference
	Get_Buffer() *IBuffer
}

type IBufferMediaFrameVtbl struct {
	win32.IInspectableVtbl
	Get_FrameReference uintptr
	Get_Buffer         uintptr
}

type IBufferMediaFrame struct {
	win32.IInspectable
}

func (this *IBufferMediaFrame) Vtbl() *IBufferMediaFrameVtbl {
	return (*IBufferMediaFrameVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBufferMediaFrame) Get_FrameReference() *IMediaFrameReference {
	var _result *IMediaFrameReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameReference, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBufferMediaFrame) Get_Buffer() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Buffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 47135E4F-8549-45C0-925B-80D35EFDB10A
var IID_IDepthMediaFrame = syscall.GUID{0x47135E4F, 0x8549, 0x45C0,
	[8]byte{0x92, 0x5B, 0x80, 0xD3, 0x5E, 0xFD, 0xB1, 0x0A}}

type IDepthMediaFrameInterface interface {
	win32.IInspectableInterface
	Get_FrameReference() *IMediaFrameReference
	Get_VideoMediaFrame() *IVideoMediaFrame
	Get_DepthFormat() *IDepthMediaFrameFormat
	TryCreateCoordinateMapper(cameraIntrinsics *ICameraIntrinsics, coordinateSystem *ISpatialCoordinateSystem) *IDepthCorrelatedCoordinateMapper
}

type IDepthMediaFrameVtbl struct {
	win32.IInspectableVtbl
	Get_FrameReference        uintptr
	Get_VideoMediaFrame       uintptr
	Get_DepthFormat           uintptr
	TryCreateCoordinateMapper uintptr
}

type IDepthMediaFrame struct {
	win32.IInspectable
}

func (this *IDepthMediaFrame) Vtbl() *IDepthMediaFrameVtbl {
	return (*IDepthMediaFrameVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDepthMediaFrame) Get_FrameReference() *IMediaFrameReference {
	var _result *IMediaFrameReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameReference, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDepthMediaFrame) Get_VideoMediaFrame() *IVideoMediaFrame {
	var _result *IVideoMediaFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoMediaFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDepthMediaFrame) Get_DepthFormat() *IDepthMediaFrameFormat {
	var _result *IDepthMediaFrameFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DepthFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDepthMediaFrame) TryCreateCoordinateMapper(cameraIntrinsics *ICameraIntrinsics, coordinateSystem *ISpatialCoordinateSystem) *IDepthCorrelatedCoordinateMapper {
	var _result *IDepthCorrelatedCoordinateMapper
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryCreateCoordinateMapper, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(cameraIntrinsics)), uintptr(unsafe.Pointer(coordinateSystem)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6CCA473D-C4A4-4176-B0CD-33EAE3B35AA3
var IID_IDepthMediaFrame2 = syscall.GUID{0x6CCA473D, 0xC4A4, 0x4176,
	[8]byte{0xB0, 0xCD, 0x33, 0xEA, 0xE3, 0xB3, 0x5A, 0xA3}}

type IDepthMediaFrame2Interface interface {
	win32.IInspectableInterface
	Get_MaxReliableDepth() uint32
	Get_MinReliableDepth() uint32
}

type IDepthMediaFrame2Vtbl struct {
	win32.IInspectableVtbl
	Get_MaxReliableDepth uintptr
	Get_MinReliableDepth uintptr
}

type IDepthMediaFrame2 struct {
	win32.IInspectable
}

func (this *IDepthMediaFrame2) Vtbl() *IDepthMediaFrame2Vtbl {
	return (*IDepthMediaFrame2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDepthMediaFrame2) Get_MaxReliableDepth() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxReliableDepth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDepthMediaFrame2) Get_MinReliableDepth() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinReliableDepth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C312CF40-D729-453E-8780-2E04F140D28E
var IID_IDepthMediaFrameFormat = syscall.GUID{0xC312CF40, 0xD729, 0x453E,
	[8]byte{0x87, 0x80, 0x2E, 0x04, 0xF1, 0x40, 0xD2, 0x8E}}

type IDepthMediaFrameFormatInterface interface {
	win32.IInspectableInterface
	Get_VideoFormat() *IVideoMediaFrameFormat
	Get_DepthScaleInMeters() float64
}

type IDepthMediaFrameFormatVtbl struct {
	win32.IInspectableVtbl
	Get_VideoFormat        uintptr
	Get_DepthScaleInMeters uintptr
}

type IDepthMediaFrameFormat struct {
	win32.IInspectable
}

func (this *IDepthMediaFrameFormat) Vtbl() *IDepthMediaFrameFormatVtbl {
	return (*IDepthMediaFrameFormatVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDepthMediaFrameFormat) Get_VideoFormat() *IVideoMediaFrameFormat {
	var _result *IVideoMediaFrameFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDepthMediaFrameFormat) Get_DepthScaleInMeters() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DepthScaleInMeters, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 3FD13503-004B-4F0E-91AC-465299B41658
var IID_IInfraredMediaFrame = syscall.GUID{0x3FD13503, 0x004B, 0x4F0E,
	[8]byte{0x91, 0xAC, 0x46, 0x52, 0x99, 0xB4, 0x16, 0x58}}

type IInfraredMediaFrameInterface interface {
	win32.IInspectableInterface
	Get_FrameReference() *IMediaFrameReference
	Get_VideoMediaFrame() *IVideoMediaFrame
	Get_IsIlluminated() bool
}

type IInfraredMediaFrameVtbl struct {
	win32.IInspectableVtbl
	Get_FrameReference  uintptr
	Get_VideoMediaFrame uintptr
	Get_IsIlluminated   uintptr
}

type IInfraredMediaFrame struct {
	win32.IInspectable
}

func (this *IInfraredMediaFrame) Vtbl() *IInfraredMediaFrameVtbl {
	return (*IInfraredMediaFrameVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInfraredMediaFrame) Get_FrameReference() *IMediaFrameReference {
	var _result *IMediaFrameReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameReference, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInfraredMediaFrame) Get_VideoMediaFrame() *IVideoMediaFrame {
	var _result *IVideoMediaFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoMediaFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInfraredMediaFrame) Get_IsIlluminated() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsIlluminated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 0B430ADD-A490-4435-ADA1-9AFFD55239F7
var IID_IMediaFrameArrivedEventArgs = syscall.GUID{0x0B430ADD, 0xA490, 0x4435,
	[8]byte{0xAD, 0xA1, 0x9A, 0xFF, 0xD5, 0x52, 0x39, 0xF7}}

type IMediaFrameArrivedEventArgsInterface interface {
	win32.IInspectableInterface
}

type IMediaFrameArrivedEventArgsVtbl struct {
	win32.IInspectableVtbl
}

type IMediaFrameArrivedEventArgs struct {
	win32.IInspectable
}

func (this *IMediaFrameArrivedEventArgs) Vtbl() *IMediaFrameArrivedEventArgsVtbl {
	return (*IMediaFrameArrivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 71902B4E-B279-4A97-A9DB-BD5A2FB78F39
var IID_IMediaFrameFormat = syscall.GUID{0x71902B4E, 0xB279, 0x4A97,
	[8]byte{0xA9, 0xDB, 0xBD, 0x5A, 0x2F, 0xB7, 0x8F, 0x39}}

type IMediaFrameFormatInterface interface {
	win32.IInspectableInterface
	Get_MajorType() string
	Get_Subtype() string
	Get_FrameRate() *IMediaRatio
	Get_Properties() *IMapView[syscall.GUID, interface{}]
	Get_VideoFormat() *IVideoMediaFrameFormat
}

type IMediaFrameFormatVtbl struct {
	win32.IInspectableVtbl
	Get_MajorType   uintptr
	Get_Subtype     uintptr
	Get_FrameRate   uintptr
	Get_Properties  uintptr
	Get_VideoFormat uintptr
}

type IMediaFrameFormat struct {
	win32.IInspectable
}

func (this *IMediaFrameFormat) Vtbl() *IMediaFrameFormatVtbl {
	return (*IMediaFrameFormatVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaFrameFormat) Get_MajorType() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MajorType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaFrameFormat) Get_Subtype() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Subtype, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaFrameFormat) Get_FrameRate() *IMediaRatio {
	var _result *IMediaRatio
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaFrameFormat) Get_Properties() *IMapView[syscall.GUID, interface{}] {
	var _result *IMapView[syscall.GUID, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaFrameFormat) Get_VideoFormat() *IVideoMediaFrameFormat {
	var _result *IVideoMediaFrameFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 63856340-5E87-4C10-86D1-6DF097A6C6A8
var IID_IMediaFrameFormat2 = syscall.GUID{0x63856340, 0x5E87, 0x4C10,
	[8]byte{0x86, 0xD1, 0x6D, 0xF0, 0x97, 0xA6, 0xC6, 0xA8}}

type IMediaFrameFormat2Interface interface {
	win32.IInspectableInterface
	Get_AudioEncodingProperties() *IAudioEncodingProperties
}

type IMediaFrameFormat2Vtbl struct {
	win32.IInspectableVtbl
	Get_AudioEncodingProperties uintptr
}

type IMediaFrameFormat2 struct {
	win32.IInspectable
}

func (this *IMediaFrameFormat2) Vtbl() *IMediaFrameFormat2Vtbl {
	return (*IMediaFrameFormat2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaFrameFormat2) Get_AudioEncodingProperties() *IAudioEncodingProperties {
	var _result *IAudioEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioEncodingProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E4C94395-2028-48ED-90B0-D1C1B162E24C
var IID_IMediaFrameReader = syscall.GUID{0xE4C94395, 0x2028, 0x48ED,
	[8]byte{0x90, 0xB0, 0xD1, 0xC1, 0xB1, 0x62, 0xE2, 0x4C}}

type IMediaFrameReaderInterface interface {
	win32.IInspectableInterface
	Add_FrameArrived(handler TypedEventHandler[*IMediaFrameReader, *IMediaFrameArrivedEventArgs]) EventRegistrationToken
	Remove_FrameArrived(token EventRegistrationToken)
	TryAcquireLatestFrame() *IMediaFrameReference
	StartAsync() *IAsyncOperation[MediaFrameReaderStartStatus]
	StopAsync() *IAsyncAction
}

type IMediaFrameReaderVtbl struct {
	win32.IInspectableVtbl
	Add_FrameArrived      uintptr
	Remove_FrameArrived   uintptr
	TryAcquireLatestFrame uintptr
	StartAsync            uintptr
	StopAsync             uintptr
}

type IMediaFrameReader struct {
	win32.IInspectable
}

func (this *IMediaFrameReader) Vtbl() *IMediaFrameReaderVtbl {
	return (*IMediaFrameReaderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaFrameReader) Add_FrameArrived(handler TypedEventHandler[*IMediaFrameReader, *IMediaFrameArrivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_FrameArrived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaFrameReader) Remove_FrameArrived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_FrameArrived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaFrameReader) TryAcquireLatestFrame() *IMediaFrameReference {
	var _result *IMediaFrameReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryAcquireLatestFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaFrameReader) StartAsync() *IAsyncOperation[MediaFrameReaderStartStatus] {
	var _result *IAsyncOperation[MediaFrameReaderStartStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaFrameReader) StopAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 871127B3-8531-4050-87CC-A13733CF3E9B
var IID_IMediaFrameReader2 = syscall.GUID{0x871127B3, 0x8531, 0x4050,
	[8]byte{0x87, 0xCC, 0xA1, 0x37, 0x33, 0xCF, 0x3E, 0x9B}}

type IMediaFrameReader2Interface interface {
	win32.IInspectableInterface
	Put_AcquisitionMode(value MediaFrameReaderAcquisitionMode)
	Get_AcquisitionMode() MediaFrameReaderAcquisitionMode
}

type IMediaFrameReader2Vtbl struct {
	win32.IInspectableVtbl
	Put_AcquisitionMode uintptr
	Get_AcquisitionMode uintptr
}

type IMediaFrameReader2 struct {
	win32.IInspectable
}

func (this *IMediaFrameReader2) Vtbl() *IMediaFrameReader2Vtbl {
	return (*IMediaFrameReader2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaFrameReader2) Put_AcquisitionMode(value MediaFrameReaderAcquisitionMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AcquisitionMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMediaFrameReader2) Get_AcquisitionMode() MediaFrameReaderAcquisitionMode {
	var _result MediaFrameReaderAcquisitionMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AcquisitionMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F6B88641-F0DC-4044-8DC9-961CEDD05BAD
var IID_IMediaFrameReference = syscall.GUID{0xF6B88641, 0xF0DC, 0x4044,
	[8]byte{0x8D, 0xC9, 0x96, 0x1C, 0xED, 0xD0, 0x5B, 0xAD}}

type IMediaFrameReferenceInterface interface {
	win32.IInspectableInterface
	Get_SourceKind() MediaFrameSourceKind
	Get_Format() *IMediaFrameFormat
	Get_SystemRelativeTime() *IReference[TimeSpan]
	Get_Duration() TimeSpan
	Get_Properties() *IMapView[syscall.GUID, interface{}]
	Get_BufferMediaFrame() *IBufferMediaFrame
	Get_VideoMediaFrame() *IVideoMediaFrame
	Get_CoordinateSystem() *ISpatialCoordinateSystem
}

type IMediaFrameReferenceVtbl struct {
	win32.IInspectableVtbl
	Get_SourceKind         uintptr
	Get_Format             uintptr
	Get_SystemRelativeTime uintptr
	Get_Duration           uintptr
	Get_Properties         uintptr
	Get_BufferMediaFrame   uintptr
	Get_VideoMediaFrame    uintptr
	Get_CoordinateSystem   uintptr
}

type IMediaFrameReference struct {
	win32.IInspectable
}

func (this *IMediaFrameReference) Vtbl() *IMediaFrameReferenceVtbl {
	return (*IMediaFrameReferenceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaFrameReference) Get_SourceKind() MediaFrameSourceKind {
	var _result MediaFrameSourceKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SourceKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaFrameReference) Get_Format() *IMediaFrameFormat {
	var _result *IMediaFrameFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Format, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaFrameReference) Get_SystemRelativeTime() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SystemRelativeTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaFrameReference) Get_Duration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Duration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaFrameReference) Get_Properties() *IMapView[syscall.GUID, interface{}] {
	var _result *IMapView[syscall.GUID, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaFrameReference) Get_BufferMediaFrame() *IBufferMediaFrame {
	var _result *IBufferMediaFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BufferMediaFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaFrameReference) Get_VideoMediaFrame() *IVideoMediaFrame {
	var _result *IVideoMediaFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoMediaFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaFrameReference) Get_CoordinateSystem() *ISpatialCoordinateSystem {
	var _result *ISpatialCoordinateSystem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CoordinateSystem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DDBC3ECC-D5B2-49EF-836A-947D989B80C1
var IID_IMediaFrameReference2 = syscall.GUID{0xDDBC3ECC, 0xD5B2, 0x49EF,
	[8]byte{0x83, 0x6A, 0x94, 0x7D, 0x98, 0x9B, 0x80, 0xC1}}

type IMediaFrameReference2Interface interface {
	win32.IInspectableInterface
	Get_AudioMediaFrame() *IAudioMediaFrame
}

type IMediaFrameReference2Vtbl struct {
	win32.IInspectableVtbl
	Get_AudioMediaFrame uintptr
}

type IMediaFrameReference2 struct {
	win32.IInspectable
}

func (this *IMediaFrameReference2) Vtbl() *IMediaFrameReference2Vtbl {
	return (*IMediaFrameReference2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaFrameReference2) Get_AudioMediaFrame() *IAudioMediaFrame {
	var _result *IAudioMediaFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioMediaFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D6782953-90DB-46A8-8ADD-2AA884A8D253
var IID_IMediaFrameSource = syscall.GUID{0xD6782953, 0x90DB, 0x46A8,
	[8]byte{0x8A, 0xDD, 0x2A, 0xA8, 0x84, 0xA8, 0xD2, 0x53}}

type IMediaFrameSourceInterface interface {
	win32.IInspectableInterface
	Get_Info() *IMediaFrameSourceInfo
	Get_Controller() *IMediaFrameSourceController
	Get_SupportedFormats() *IVectorView[*IMediaFrameFormat]
	Get_CurrentFormat() *IMediaFrameFormat
	SetFormatAsync(format *IMediaFrameFormat) *IAsyncAction
	Add_FormatChanged(handler TypedEventHandler[*IMediaFrameSource, interface{}]) EventRegistrationToken
	Remove_FormatChanged(token EventRegistrationToken)
	TryGetCameraIntrinsics(format *IMediaFrameFormat) *ICameraIntrinsics
}

type IMediaFrameSourceVtbl struct {
	win32.IInspectableVtbl
	Get_Info               uintptr
	Get_Controller         uintptr
	Get_SupportedFormats   uintptr
	Get_CurrentFormat      uintptr
	SetFormatAsync         uintptr
	Add_FormatChanged      uintptr
	Remove_FormatChanged   uintptr
	TryGetCameraIntrinsics uintptr
}

type IMediaFrameSource struct {
	win32.IInspectable
}

func (this *IMediaFrameSource) Vtbl() *IMediaFrameSourceVtbl {
	return (*IMediaFrameSourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaFrameSource) Get_Info() *IMediaFrameSourceInfo {
	var _result *IMediaFrameSourceInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Info, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaFrameSource) Get_Controller() *IMediaFrameSourceController {
	var _result *IMediaFrameSourceController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Controller, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaFrameSource) Get_SupportedFormats() *IVectorView[*IMediaFrameFormat] {
	var _result *IVectorView[*IMediaFrameFormat]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedFormats, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaFrameSource) Get_CurrentFormat() *IMediaFrameFormat {
	var _result *IMediaFrameFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaFrameSource) SetFormatAsync(format *IMediaFrameFormat) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetFormatAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(format)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaFrameSource) Add_FormatChanged(handler TypedEventHandler[*IMediaFrameSource, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_FormatChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaFrameSource) Remove_FormatChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_FormatChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMediaFrameSource) TryGetCameraIntrinsics(format *IMediaFrameFormat) *ICameraIntrinsics {
	var _result *ICameraIntrinsics
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetCameraIntrinsics, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(format)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6D076635-316D-4B8F-B7B6-EEB04A8C6525
var IID_IMediaFrameSourceController = syscall.GUID{0x6D076635, 0x316D, 0x4B8F,
	[8]byte{0xB7, 0xB6, 0xEE, 0xB0, 0x4A, 0x8C, 0x65, 0x25}}

type IMediaFrameSourceControllerInterface interface {
	win32.IInspectableInterface
	GetPropertyAsync(propertyId string) *IAsyncOperation[*IMediaFrameSourceGetPropertyResult]
	SetPropertyAsync(propertyId string, propertyValue interface{}) *IAsyncOperation[MediaFrameSourceSetPropertyStatus]
	Get_VideoDeviceController() *IVideoDeviceController
}

type IMediaFrameSourceControllerVtbl struct {
	win32.IInspectableVtbl
	GetPropertyAsync          uintptr
	SetPropertyAsync          uintptr
	Get_VideoDeviceController uintptr
}

type IMediaFrameSourceController struct {
	win32.IInspectable
}

func (this *IMediaFrameSourceController) Vtbl() *IMediaFrameSourceControllerVtbl {
	return (*IMediaFrameSourceControllerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaFrameSourceController) GetPropertyAsync(propertyId string) *IAsyncOperation[*IMediaFrameSourceGetPropertyResult] {
	var _result *IAsyncOperation[*IMediaFrameSourceGetPropertyResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPropertyAsync, uintptr(unsafe.Pointer(this)), NewHStr(propertyId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaFrameSourceController) SetPropertyAsync(propertyId string, propertyValue interface{}) *IAsyncOperation[MediaFrameSourceSetPropertyStatus] {
	var _result *IAsyncOperation[MediaFrameSourceSetPropertyStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetPropertyAsync, uintptr(unsafe.Pointer(this)), NewHStr(propertyId).Ptr, *(*uintptr)(unsafe.Pointer(&propertyValue)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaFrameSourceController) Get_VideoDeviceController() *IVideoDeviceController {
	var _result *IVideoDeviceController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoDeviceController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EFC49FD4-FCF2-4A03-B4E4-AC9628739BEE
var IID_IMediaFrameSourceController2 = syscall.GUID{0xEFC49FD4, 0xFCF2, 0x4A03,
	[8]byte{0xB4, 0xE4, 0xAC, 0x96, 0x28, 0x73, 0x9B, 0xEE}}

type IMediaFrameSourceController2Interface interface {
	win32.IInspectableInterface
	GetPropertyByExtendedIdAsync(extendedPropertyIdLength uint32, extendedPropertyId *byte, maxPropertyValueSize *IReference[uint32]) *IAsyncOperation[*IMediaFrameSourceGetPropertyResult]
	SetPropertyByExtendedIdAsync(extendedPropertyIdLength uint32, extendedPropertyId *byte, propertyValueLength uint32, propertyValue *byte) *IAsyncOperation[MediaFrameSourceSetPropertyStatus]
}

type IMediaFrameSourceController2Vtbl struct {
	win32.IInspectableVtbl
	GetPropertyByExtendedIdAsync uintptr
	SetPropertyByExtendedIdAsync uintptr
}

type IMediaFrameSourceController2 struct {
	win32.IInspectable
}

func (this *IMediaFrameSourceController2) Vtbl() *IMediaFrameSourceController2Vtbl {
	return (*IMediaFrameSourceController2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaFrameSourceController2) GetPropertyByExtendedIdAsync(extendedPropertyIdLength uint32, extendedPropertyId *byte, maxPropertyValueSize *IReference[uint32]) *IAsyncOperation[*IMediaFrameSourceGetPropertyResult] {
	var _result *IAsyncOperation[*IMediaFrameSourceGetPropertyResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPropertyByExtendedIdAsync, uintptr(unsafe.Pointer(this)), uintptr(extendedPropertyIdLength), uintptr(unsafe.Pointer(extendedPropertyId)), uintptr(unsafe.Pointer(maxPropertyValueSize)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaFrameSourceController2) SetPropertyByExtendedIdAsync(extendedPropertyIdLength uint32, extendedPropertyId *byte, propertyValueLength uint32, propertyValue *byte) *IAsyncOperation[MediaFrameSourceSetPropertyStatus] {
	var _result *IAsyncOperation[MediaFrameSourceSetPropertyStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetPropertyByExtendedIdAsync, uintptr(unsafe.Pointer(this)), uintptr(extendedPropertyIdLength), uintptr(unsafe.Pointer(extendedPropertyId)), uintptr(propertyValueLength), uintptr(unsafe.Pointer(propertyValue)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1F0CF815-2464-4651-B1E8-4A82DBDB54DE
var IID_IMediaFrameSourceController3 = syscall.GUID{0x1F0CF815, 0x2464, 0x4651,
	[8]byte{0xB1, 0xE8, 0x4A, 0x82, 0xDB, 0xDB, 0x54, 0xDE}}

type IMediaFrameSourceController3Interface interface {
	win32.IInspectableInterface
	Get_AudioDeviceController() *IAudioDeviceController
}

type IMediaFrameSourceController3Vtbl struct {
	win32.IInspectableVtbl
	Get_AudioDeviceController uintptr
}

type IMediaFrameSourceController3 struct {
	win32.IInspectable
}

func (this *IMediaFrameSourceController3) Vtbl() *IMediaFrameSourceController3Vtbl {
	return (*IMediaFrameSourceController3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaFrameSourceController3) Get_AudioDeviceController() *IAudioDeviceController {
	var _result *IAudioDeviceController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudioDeviceController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 088616C2-3A64-4BD5-BD2B-E7C898D2F37A
var IID_IMediaFrameSourceGetPropertyResult = syscall.GUID{0x088616C2, 0x3A64, 0x4BD5,
	[8]byte{0xBD, 0x2B, 0xE7, 0xC8, 0x98, 0xD2, 0xF3, 0x7A}}

type IMediaFrameSourceGetPropertyResultInterface interface {
	win32.IInspectableInterface
	Get_Status() MediaFrameSourceGetPropertyStatus
	Get_Value() interface{}
}

type IMediaFrameSourceGetPropertyResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
	Get_Value  uintptr
}

type IMediaFrameSourceGetPropertyResult struct {
	win32.IInspectable
}

func (this *IMediaFrameSourceGetPropertyResult) Vtbl() *IMediaFrameSourceGetPropertyResultVtbl {
	return (*IMediaFrameSourceGetPropertyResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaFrameSourceGetPropertyResult) Get_Status() MediaFrameSourceGetPropertyStatus {
	var _result MediaFrameSourceGetPropertyStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaFrameSourceGetPropertyResult) Get_Value() interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7F605B87-4832-4B5F-AE3D-412FAAB37D34
var IID_IMediaFrameSourceGroup = syscall.GUID{0x7F605B87, 0x4832, 0x4B5F,
	[8]byte{0xAE, 0x3D, 0x41, 0x2F, 0xAA, 0xB3, 0x7D, 0x34}}

type IMediaFrameSourceGroupInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	Get_DisplayName() string
	Get_SourceInfos() *IVectorView[*IMediaFrameSourceInfo]
}

type IMediaFrameSourceGroupVtbl struct {
	win32.IInspectableVtbl
	Get_Id          uintptr
	Get_DisplayName uintptr
	Get_SourceInfos uintptr
}

type IMediaFrameSourceGroup struct {
	win32.IInspectable
}

func (this *IMediaFrameSourceGroup) Vtbl() *IMediaFrameSourceGroupVtbl {
	return (*IMediaFrameSourceGroupVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaFrameSourceGroup) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaFrameSourceGroup) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaFrameSourceGroup) Get_SourceInfos() *IVectorView[*IMediaFrameSourceInfo] {
	var _result *IVectorView[*IMediaFrameSourceInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SourceInfos, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1C48BFC5-436F-4508-94CF-D5D8B7326445
var IID_IMediaFrameSourceGroupStatics = syscall.GUID{0x1C48BFC5, 0x436F, 0x4508,
	[8]byte{0x94, 0xCF, 0xD5, 0xD8, 0xB7, 0x32, 0x64, 0x45}}

type IMediaFrameSourceGroupStaticsInterface interface {
	win32.IInspectableInterface
	FindAllAsync() *IAsyncOperation[*IVectorView[*IMediaFrameSourceGroup]]
	FromIdAsync(id string) *IAsyncOperation[*IMediaFrameSourceGroup]
	GetDeviceSelector() string
}

type IMediaFrameSourceGroupStaticsVtbl struct {
	win32.IInspectableVtbl
	FindAllAsync      uintptr
	FromIdAsync       uintptr
	GetDeviceSelector uintptr
}

type IMediaFrameSourceGroupStatics struct {
	win32.IInspectable
}

func (this *IMediaFrameSourceGroupStatics) Vtbl() *IMediaFrameSourceGroupStaticsVtbl {
	return (*IMediaFrameSourceGroupStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaFrameSourceGroupStatics) FindAllAsync() *IAsyncOperation[*IVectorView[*IMediaFrameSourceGroup]] {
	var _result *IAsyncOperation[*IVectorView[*IMediaFrameSourceGroup]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaFrameSourceGroupStatics) FromIdAsync(id string) *IAsyncOperation[*IMediaFrameSourceGroup] {
	var _result *IAsyncOperation[*IMediaFrameSourceGroup]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(id).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaFrameSourceGroupStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 87BDC9CD-4601-408F-91CF-038318CD0AF3
var IID_IMediaFrameSourceInfo = syscall.GUID{0x87BDC9CD, 0x4601, 0x408F,
	[8]byte{0x91, 0xCF, 0x03, 0x83, 0x18, 0xCD, 0x0A, 0xF3}}

type IMediaFrameSourceInfoInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	Get_MediaStreamType() MediaStreamType
	Get_SourceKind() MediaFrameSourceKind
	Get_SourceGroup() *IMediaFrameSourceGroup
	Get_DeviceInformation() *IDeviceInformation
	Get_Properties() *IMapView[syscall.GUID, interface{}]
	Get_CoordinateSystem() *ISpatialCoordinateSystem
}

type IMediaFrameSourceInfoVtbl struct {
	win32.IInspectableVtbl
	Get_Id                uintptr
	Get_MediaStreamType   uintptr
	Get_SourceKind        uintptr
	Get_SourceGroup       uintptr
	Get_DeviceInformation uintptr
	Get_Properties        uintptr
	Get_CoordinateSystem  uintptr
}

type IMediaFrameSourceInfo struct {
	win32.IInspectable
}

func (this *IMediaFrameSourceInfo) Vtbl() *IMediaFrameSourceInfoVtbl {
	return (*IMediaFrameSourceInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaFrameSourceInfo) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaFrameSourceInfo) Get_MediaStreamType() MediaStreamType {
	var _result MediaStreamType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaStreamType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaFrameSourceInfo) Get_SourceKind() MediaFrameSourceKind {
	var _result MediaFrameSourceKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SourceKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaFrameSourceInfo) Get_SourceGroup() *IMediaFrameSourceGroup {
	var _result *IMediaFrameSourceGroup
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SourceGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaFrameSourceInfo) Get_DeviceInformation() *IDeviceInformation {
	var _result *IDeviceInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaFrameSourceInfo) Get_Properties() *IMapView[syscall.GUID, interface{}] {
	var _result *IMapView[syscall.GUID, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaFrameSourceInfo) Get_CoordinateSystem() *ISpatialCoordinateSystem {
	var _result *ISpatialCoordinateSystem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CoordinateSystem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 195A7855-6457-42C6-A769-19B65BD32E6E
var IID_IMediaFrameSourceInfo2 = syscall.GUID{0x195A7855, 0x6457, 0x42C6,
	[8]byte{0xA7, 0x69, 0x19, 0xB6, 0x5B, 0xD3, 0x2E, 0x6E}}

type IMediaFrameSourceInfo2Interface interface {
	win32.IInspectableInterface
	Get_ProfileId() string
	Get_VideoProfileMediaDescription() *IVectorView[*IMediaCaptureVideoProfileMediaDescription]
}

type IMediaFrameSourceInfo2Vtbl struct {
	win32.IInspectableVtbl
	Get_ProfileId                    uintptr
	Get_VideoProfileMediaDescription uintptr
}

type IMediaFrameSourceInfo2 struct {
	win32.IInspectable
}

func (this *IMediaFrameSourceInfo2) Vtbl() *IMediaFrameSourceInfo2Vtbl {
	return (*IMediaFrameSourceInfo2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaFrameSourceInfo2) Get_ProfileId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProfileId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaFrameSourceInfo2) Get_VideoProfileMediaDescription() *IVectorView[*IMediaCaptureVideoProfileMediaDescription] {
	var _result *IVectorView[*IMediaCaptureVideoProfileMediaDescription]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoProfileMediaDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CA824AB6-66EA-5885-A2B6-26C0EEEC3C7B
var IID_IMediaFrameSourceInfo3 = syscall.GUID{0xCA824AB6, 0x66EA, 0x5885,
	[8]byte{0xA2, 0xB6, 0x26, 0xC0, 0xEE, 0xEC, 0x3C, 0x7B}}

type IMediaFrameSourceInfo3Interface interface {
	win32.IInspectableInterface
	GetRelativePanel(displayRegion unsafe.Pointer) Panel
}

type IMediaFrameSourceInfo3Vtbl struct {
	win32.IInspectableVtbl
	GetRelativePanel uintptr
}

type IMediaFrameSourceInfo3 struct {
	win32.IInspectable
}

func (this *IMediaFrameSourceInfo3) Vtbl() *IMediaFrameSourceInfo3Vtbl {
	return (*IMediaFrameSourceInfo3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaFrameSourceInfo3) GetRelativePanel(displayRegion unsafe.Pointer) Panel {
	var _result Panel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetRelativePanel, uintptr(unsafe.Pointer(this)), uintptr(displayRegion), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4817D721-85EB-470C-8F37-43CA5498E41D
var IID_IMediaFrameSourceInfo4 = syscall.GUID{0x4817D721, 0x85EB, 0x470C,
	[8]byte{0x8F, 0x37, 0x43, 0xCA, 0x54, 0x98, 0xE4, 0x1D}}

type IMediaFrameSourceInfo4Interface interface {
	win32.IInspectableInterface
	Get_IsShareable() bool
}

type IMediaFrameSourceInfo4Vtbl struct {
	win32.IInspectableVtbl
	Get_IsShareable uintptr
}

type IMediaFrameSourceInfo4 struct {
	win32.IInspectable
}

func (this *IMediaFrameSourceInfo4) Vtbl() *IMediaFrameSourceInfo4Vtbl {
	return (*IMediaFrameSourceInfo4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaFrameSourceInfo4) Get_IsShareable() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsShareable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 63115E01-CF51-48FD-AAB0-6D693EB48127
var IID_IMultiSourceMediaFrameArrivedEventArgs = syscall.GUID{0x63115E01, 0xCF51, 0x48FD,
	[8]byte{0xAA, 0xB0, 0x6D, 0x69, 0x3E, 0xB4, 0x81, 0x27}}

type IMultiSourceMediaFrameArrivedEventArgsInterface interface {
	win32.IInspectableInterface
}

type IMultiSourceMediaFrameArrivedEventArgsVtbl struct {
	win32.IInspectableVtbl
}

type IMultiSourceMediaFrameArrivedEventArgs struct {
	win32.IInspectable
}

func (this *IMultiSourceMediaFrameArrivedEventArgs) Vtbl() *IMultiSourceMediaFrameArrivedEventArgsVtbl {
	return (*IMultiSourceMediaFrameArrivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 8D144402-F763-488D-98F2-B437BCF075E7
var IID_IMultiSourceMediaFrameReader = syscall.GUID{0x8D144402, 0xF763, 0x488D,
	[8]byte{0x98, 0xF2, 0xB4, 0x37, 0xBC, 0xF0, 0x75, 0xE7}}

type IMultiSourceMediaFrameReaderInterface interface {
	win32.IInspectableInterface
	Add_FrameArrived(handler TypedEventHandler[*IMultiSourceMediaFrameReader, *IMultiSourceMediaFrameArrivedEventArgs]) EventRegistrationToken
	Remove_FrameArrived(token EventRegistrationToken)
	TryAcquireLatestFrame() *IMultiSourceMediaFrameReference
	StartAsync() *IAsyncOperation[MultiSourceMediaFrameReaderStartStatus]
	StopAsync() *IAsyncAction
}

type IMultiSourceMediaFrameReaderVtbl struct {
	win32.IInspectableVtbl
	Add_FrameArrived      uintptr
	Remove_FrameArrived   uintptr
	TryAcquireLatestFrame uintptr
	StartAsync            uintptr
	StopAsync             uintptr
}

type IMultiSourceMediaFrameReader struct {
	win32.IInspectable
}

func (this *IMultiSourceMediaFrameReader) Vtbl() *IMultiSourceMediaFrameReaderVtbl {
	return (*IMultiSourceMediaFrameReaderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMultiSourceMediaFrameReader) Add_FrameArrived(handler TypedEventHandler[*IMultiSourceMediaFrameReader, *IMultiSourceMediaFrameArrivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_FrameArrived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMultiSourceMediaFrameReader) Remove_FrameArrived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_FrameArrived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IMultiSourceMediaFrameReader) TryAcquireLatestFrame() *IMultiSourceMediaFrameReference {
	var _result *IMultiSourceMediaFrameReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryAcquireLatestFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMultiSourceMediaFrameReader) StartAsync() *IAsyncOperation[MultiSourceMediaFrameReaderStartStatus] {
	var _result *IAsyncOperation[MultiSourceMediaFrameReaderStartStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMultiSourceMediaFrameReader) StopAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EF5C8ABD-FC5C-4C6B-9D81-3CB9CC637C26
var IID_IMultiSourceMediaFrameReader2 = syscall.GUID{0xEF5C8ABD, 0xFC5C, 0x4C6B,
	[8]byte{0x9D, 0x81, 0x3C, 0xB9, 0xCC, 0x63, 0x7C, 0x26}}

type IMultiSourceMediaFrameReader2Interface interface {
	win32.IInspectableInterface
	Put_AcquisitionMode(value MediaFrameReaderAcquisitionMode)
	Get_AcquisitionMode() MediaFrameReaderAcquisitionMode
}

type IMultiSourceMediaFrameReader2Vtbl struct {
	win32.IInspectableVtbl
	Put_AcquisitionMode uintptr
	Get_AcquisitionMode uintptr
}

type IMultiSourceMediaFrameReader2 struct {
	win32.IInspectable
}

func (this *IMultiSourceMediaFrameReader2) Vtbl() *IMultiSourceMediaFrameReader2Vtbl {
	return (*IMultiSourceMediaFrameReader2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMultiSourceMediaFrameReader2) Put_AcquisitionMode(value MediaFrameReaderAcquisitionMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AcquisitionMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMultiSourceMediaFrameReader2) Get_AcquisitionMode() MediaFrameReaderAcquisitionMode {
	var _result MediaFrameReaderAcquisitionMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AcquisitionMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 21964B1A-7FE2-44D6-92E5-298E6D2810E9
var IID_IMultiSourceMediaFrameReference = syscall.GUID{0x21964B1A, 0x7FE2, 0x44D6,
	[8]byte{0x92, 0xE5, 0x29, 0x8E, 0x6D, 0x28, 0x10, 0xE9}}

type IMultiSourceMediaFrameReferenceInterface interface {
	win32.IInspectableInterface
	TryGetFrameReferenceBySourceId(sourceId string) *IMediaFrameReference
}

type IMultiSourceMediaFrameReferenceVtbl struct {
	win32.IInspectableVtbl
	TryGetFrameReferenceBySourceId uintptr
}

type IMultiSourceMediaFrameReference struct {
	win32.IInspectable
}

func (this *IMultiSourceMediaFrameReference) Vtbl() *IMultiSourceMediaFrameReferenceVtbl {
	return (*IMultiSourceMediaFrameReferenceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMultiSourceMediaFrameReference) TryGetFrameReferenceBySourceId(sourceId string) *IMediaFrameReference {
	var _result *IMediaFrameReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetFrameReferenceBySourceId, uintptr(unsafe.Pointer(this)), NewHStr(sourceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 00DD4CCB-32BD-4FE1-A013-7CC13CF5DBCF
var IID_IVideoMediaFrame = syscall.GUID{0x00DD4CCB, 0x32BD, 0x4FE1,
	[8]byte{0xA0, 0x13, 0x7C, 0xC1, 0x3C, 0xF5, 0xDB, 0xCF}}

type IVideoMediaFrameInterface interface {
	win32.IInspectableInterface
	Get_FrameReference() *IMediaFrameReference
	Get_VideoFormat() *IVideoMediaFrameFormat
	Get_SoftwareBitmap() *ISoftwareBitmap
	Get_Direct3DSurface() unsafe.Pointer
	Get_CameraIntrinsics() *ICameraIntrinsics
	Get_InfraredMediaFrame() *IInfraredMediaFrame
	Get_DepthMediaFrame() *IDepthMediaFrame
	GetVideoFrame() *IVideoFrame
}

type IVideoMediaFrameVtbl struct {
	win32.IInspectableVtbl
	Get_FrameReference     uintptr
	Get_VideoFormat        uintptr
	Get_SoftwareBitmap     uintptr
	Get_Direct3DSurface    uintptr
	Get_CameraIntrinsics   uintptr
	Get_InfraredMediaFrame uintptr
	Get_DepthMediaFrame    uintptr
	GetVideoFrame          uintptr
}

type IVideoMediaFrame struct {
	win32.IInspectable
}

func (this *IVideoMediaFrame) Vtbl() *IVideoMediaFrameVtbl {
	return (*IVideoMediaFrameVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoMediaFrame) Get_FrameReference() *IMediaFrameReference {
	var _result *IMediaFrameReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameReference, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoMediaFrame) Get_VideoFormat() *IVideoMediaFrameFormat {
	var _result *IVideoMediaFrameFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoMediaFrame) Get_SoftwareBitmap() *ISoftwareBitmap {
	var _result *ISoftwareBitmap
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SoftwareBitmap, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoMediaFrame) Get_Direct3DSurface() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Direct3DSurface, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoMediaFrame) Get_CameraIntrinsics() *ICameraIntrinsics {
	var _result *ICameraIntrinsics
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CameraIntrinsics, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoMediaFrame) Get_InfraredMediaFrame() *IInfraredMediaFrame {
	var _result *IInfraredMediaFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InfraredMediaFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoMediaFrame) Get_DepthMediaFrame() *IDepthMediaFrame {
	var _result *IDepthMediaFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DepthMediaFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoMediaFrame) GetVideoFrame() *IVideoFrame {
	var _result *IVideoFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetVideoFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 46027FC0-D71B-45C7-8F14-6D9A0AE604E4
var IID_IVideoMediaFrameFormat = syscall.GUID{0x46027FC0, 0xD71B, 0x45C7,
	[8]byte{0x8F, 0x14, 0x6D, 0x9A, 0x0A, 0xE6, 0x04, 0xE4}}

type IVideoMediaFrameFormatInterface interface {
	win32.IInspectableInterface
	Get_MediaFrameFormat() *IMediaFrameFormat
	Get_DepthFormat() *IDepthMediaFrameFormat
	Get_Width() uint32
	Get_Height() uint32
}

type IVideoMediaFrameFormatVtbl struct {
	win32.IInspectableVtbl
	Get_MediaFrameFormat uintptr
	Get_DepthFormat      uintptr
	Get_Width            uintptr
	Get_Height           uintptr
}

type IVideoMediaFrameFormat struct {
	win32.IInspectable
}

func (this *IVideoMediaFrameFormat) Vtbl() *IVideoMediaFrameFormatVtbl {
	return (*IVideoMediaFrameFormatVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoMediaFrameFormat) Get_MediaFrameFormat() *IMediaFrameFormat {
	var _result *IMediaFrameFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MediaFrameFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoMediaFrameFormat) Get_DepthFormat() *IDepthMediaFrameFormat {
	var _result *IDepthMediaFrameFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DepthFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoMediaFrameFormat) Get_Width() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Width, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoMediaFrameFormat) Get_Height() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Height, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type BufferMediaFrame struct {
	RtClass
	*IBufferMediaFrame
}

type DepthMediaFrame struct {
	RtClass
	*IDepthMediaFrame
}

type DepthMediaFrameFormat struct {
	RtClass
	*IDepthMediaFrameFormat
}

type InfraredMediaFrame struct {
	RtClass
	*IInfraredMediaFrame
}

type MediaFrameArrivedEventArgs struct {
	RtClass
	*IMediaFrameArrivedEventArgs
}

type MediaFrameFormat struct {
	RtClass
	*IMediaFrameFormat
}

type MediaFrameReader struct {
	RtClass
	*IMediaFrameReader
}

type MediaFrameReference struct {
	RtClass
	*IMediaFrameReference
}

type MediaFrameSource struct {
	RtClass
	*IMediaFrameSource
}

type MediaFrameSourceController struct {
	RtClass
	*IMediaFrameSourceController
}

type MediaFrameSourceGetPropertyResult struct {
	RtClass
	*IMediaFrameSourceGetPropertyResult
}

type MediaFrameSourceGroup struct {
	RtClass
	*IMediaFrameSourceGroup
}

func NewIMediaFrameSourceGroupStatics() *IMediaFrameSourceGroupStatics {
	var p *IMediaFrameSourceGroupStatics
	hs := NewHStr("Windows.Media.Capture.Frames.MediaFrameSourceGroup")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMediaFrameSourceGroupStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type MediaFrameSourceInfo struct {
	RtClass
	*IMediaFrameSourceInfo
}

type MultiSourceMediaFrameArrivedEventArgs struct {
	RtClass
	*IMultiSourceMediaFrameArrivedEventArgs
}

type MultiSourceMediaFrameReader struct {
	RtClass
	*IMultiSourceMediaFrameReader
}

type MultiSourceMediaFrameReference struct {
	RtClass
	*IMultiSourceMediaFrameReference
}

type VideoMediaFrame struct {
	RtClass
	*IVideoMediaFrame
}

type VideoMediaFrameFormat struct {
	RtClass
	*IVideoMediaFrameFormat
}
