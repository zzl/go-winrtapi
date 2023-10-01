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
type BitmapAlphaMode int32

const (
	BitmapAlphaMode_Premultiplied BitmapAlphaMode = 0
	BitmapAlphaMode_Straight      BitmapAlphaMode = 1
	BitmapAlphaMode_Ignore        BitmapAlphaMode = 2
)

// enum
type BitmapBufferAccessMode int32

const (
	BitmapBufferAccessMode_Read      BitmapBufferAccessMode = 0
	BitmapBufferAccessMode_ReadWrite BitmapBufferAccessMode = 1
	BitmapBufferAccessMode_Write     BitmapBufferAccessMode = 2
)

// enum
type BitmapFlip int32

const (
	BitmapFlip_None       BitmapFlip = 0
	BitmapFlip_Horizontal BitmapFlip = 1
	BitmapFlip_Vertical   BitmapFlip = 2
)

// enum
type BitmapInterpolationMode int32

const (
	BitmapInterpolationMode_NearestNeighbor BitmapInterpolationMode = 0
	BitmapInterpolationMode_Linear          BitmapInterpolationMode = 1
	BitmapInterpolationMode_Cubic           BitmapInterpolationMode = 2
	BitmapInterpolationMode_Fant            BitmapInterpolationMode = 3
)

// enum
type BitmapPixelFormat int32

const (
	BitmapPixelFormat_Unknown BitmapPixelFormat = 0
	BitmapPixelFormat_Rgba16  BitmapPixelFormat = 12
	BitmapPixelFormat_Rgba8   BitmapPixelFormat = 30
	BitmapPixelFormat_Gray16  BitmapPixelFormat = 57
	BitmapPixelFormat_Gray8   BitmapPixelFormat = 62
	BitmapPixelFormat_Bgra8   BitmapPixelFormat = 87
	BitmapPixelFormat_Nv12    BitmapPixelFormat = 103
	BitmapPixelFormat_P010    BitmapPixelFormat = 104
	BitmapPixelFormat_Yuy2    BitmapPixelFormat = 107
)

// enum
type BitmapRotation int32

const (
	BitmapRotation_None                BitmapRotation = 0
	BitmapRotation_Clockwise90Degrees  BitmapRotation = 1
	BitmapRotation_Clockwise180Degrees BitmapRotation = 2
	BitmapRotation_Clockwise270Degrees BitmapRotation = 3
)

// enum
type ColorManagementMode int32

const (
	ColorManagementMode_DoNotColorManage  ColorManagementMode = 0
	ColorManagementMode_ColorManageToSRgb ColorManagementMode = 1
)

// enum
type ExifOrientationMode int32

const (
	ExifOrientationMode_IgnoreExifOrientation  ExifOrientationMode = 0
	ExifOrientationMode_RespectExifOrientation ExifOrientationMode = 1
)

// enum
type JpegSubsamplingMode int32

const (
	JpegSubsamplingMode_Default  JpegSubsamplingMode = 0
	JpegSubsamplingMode_Y4Cb2Cr0 JpegSubsamplingMode = 1
	JpegSubsamplingMode_Y4Cb2Cr2 JpegSubsamplingMode = 2
	JpegSubsamplingMode_Y4Cb4Cr4 JpegSubsamplingMode = 3
)

// enum
type PngFilterMode int32

const (
	PngFilterMode_Automatic PngFilterMode = 0
	PngFilterMode_None      PngFilterMode = 1
	PngFilterMode_Sub       PngFilterMode = 2
	PngFilterMode_Up        PngFilterMode = 3
	PngFilterMode_Average   PngFilterMode = 4
	PngFilterMode_Paeth     PngFilterMode = 5
	PngFilterMode_Adaptive  PngFilterMode = 6
)

// enum
type TiffCompressionMode int32

const (
	TiffCompressionMode_Automatic        TiffCompressionMode = 0
	TiffCompressionMode_None             TiffCompressionMode = 1
	TiffCompressionMode_Ccitt3           TiffCompressionMode = 2
	TiffCompressionMode_Ccitt4           TiffCompressionMode = 3
	TiffCompressionMode_Lzw              TiffCompressionMode = 4
	TiffCompressionMode_Rle              TiffCompressionMode = 5
	TiffCompressionMode_Zip              TiffCompressionMode = 6
	TiffCompressionMode_LzwhDifferencing TiffCompressionMode = 7
)

// structs

type BitmapBounds struct {
	X      uint32
	Y      uint32
	Width  uint32
	Height uint32
}

type BitmapPlaneDescription struct {
	StartIndex int32
	Width      int32
	Height     int32
	Stride     int32
}

type BitmapSize struct {
	Width  uint32
	Height uint32
}

// interfaces

// A53E04C4-399C-438C-B28F-A63A6B83D1A1
var IID_IBitmapBuffer = syscall.GUID{0xA53E04C4, 0x399C, 0x438C,
	[8]byte{0xB2, 0x8F, 0xA6, 0x3A, 0x6B, 0x83, 0xD1, 0xA1}}

type IBitmapBufferInterface interface {
	win32.IInspectableInterface
	GetPlaneCount() int32
	GetPlaneDescription(index int32) BitmapPlaneDescription
}

type IBitmapBufferVtbl struct {
	win32.IInspectableVtbl
	GetPlaneCount       uintptr
	GetPlaneDescription uintptr
}

type IBitmapBuffer struct {
	win32.IInspectable
}

func (this *IBitmapBuffer) Vtbl() *IBitmapBufferVtbl {
	return (*IBitmapBufferVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBitmapBuffer) GetPlaneCount() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPlaneCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapBuffer) GetPlaneDescription(index int32) BitmapPlaneDescription {
	var _result BitmapPlaneDescription
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPlaneDescription, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 400CAAF2-C4B0-4392-A3B0-6F6F9BA95CB4
var IID_IBitmapCodecInformation = syscall.GUID{0x400CAAF2, 0xC4B0, 0x4392,
	[8]byte{0xA3, 0xB0, 0x6F, 0x6F, 0x9B, 0xA9, 0x5C, 0xB4}}

type IBitmapCodecInformationInterface interface {
	win32.IInspectableInterface
	Get_CodecId() syscall.GUID
	Get_FileExtensions() *IVectorView[string]
	Get_FriendlyName() string
	Get_MimeTypes() *IVectorView[string]
}

type IBitmapCodecInformationVtbl struct {
	win32.IInspectableVtbl
	Get_CodecId        uintptr
	Get_FileExtensions uintptr
	Get_FriendlyName   uintptr
	Get_MimeTypes      uintptr
}

type IBitmapCodecInformation struct {
	win32.IInspectable
}

func (this *IBitmapCodecInformation) Vtbl() *IBitmapCodecInformationVtbl {
	return (*IBitmapCodecInformationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBitmapCodecInformation) Get_CodecId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CodecId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapCodecInformation) Get_FileExtensions() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FileExtensions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBitmapCodecInformation) Get_FriendlyName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FriendlyName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IBitmapCodecInformation) Get_MimeTypes() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MimeTypes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// ACEF22BA-1D74-4C91-9DFC-9620745233E6
var IID_IBitmapDecoder = syscall.GUID{0xACEF22BA, 0x1D74, 0x4C91,
	[8]byte{0x9D, 0xFC, 0x96, 0x20, 0x74, 0x52, 0x33, 0xE6}}

type IBitmapDecoderInterface interface {
	win32.IInspectableInterface
	Get_BitmapContainerProperties() *IBitmapPropertiesView
	Get_DecoderInformation() *IBitmapCodecInformation
	Get_FrameCount() uint32
	GetPreviewAsync() *IAsyncOperation[*IRandomAccessStreamWithContentType]
	GetFrameAsync(frameIndex uint32) *IAsyncOperation[*IBitmapFrame]
}

type IBitmapDecoderVtbl struct {
	win32.IInspectableVtbl
	Get_BitmapContainerProperties uintptr
	Get_DecoderInformation        uintptr
	Get_FrameCount                uintptr
	GetPreviewAsync               uintptr
	GetFrameAsync                 uintptr
}

type IBitmapDecoder struct {
	win32.IInspectable
}

func (this *IBitmapDecoder) Vtbl() *IBitmapDecoderVtbl {
	return (*IBitmapDecoderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBitmapDecoder) Get_BitmapContainerProperties() *IBitmapPropertiesView {
	var _result *IBitmapPropertiesView
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BitmapContainerProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBitmapDecoder) Get_DecoderInformation() *IBitmapCodecInformation {
	var _result *IBitmapCodecInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DecoderInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBitmapDecoder) Get_FrameCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapDecoder) GetPreviewAsync() *IAsyncOperation[*IRandomAccessStreamWithContentType] {
	var _result *IAsyncOperation[*IRandomAccessStreamWithContentType]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPreviewAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBitmapDecoder) GetFrameAsync(frameIndex uint32) *IAsyncOperation[*IBitmapFrame] {
	var _result *IAsyncOperation[*IBitmapFrame]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetFrameAsync, uintptr(unsafe.Pointer(this)), uintptr(frameIndex), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 438CCB26-BCEF-4E95-BAD6-23A822E58D01
var IID_IBitmapDecoderStatics = syscall.GUID{0x438CCB26, 0xBCEF, 0x4E95,
	[8]byte{0xBA, 0xD6, 0x23, 0xA8, 0x22, 0xE5, 0x8D, 0x01}}

type IBitmapDecoderStaticsInterface interface {
	win32.IInspectableInterface
	Get_BmpDecoderId() syscall.GUID
	Get_JpegDecoderId() syscall.GUID
	Get_PngDecoderId() syscall.GUID
	Get_TiffDecoderId() syscall.GUID
	Get_GifDecoderId() syscall.GUID
	Get_JpegXRDecoderId() syscall.GUID
	Get_IcoDecoderId() syscall.GUID
	GetDecoderInformationEnumerator() *IVectorView[*IBitmapCodecInformation]
	CreateAsync(stream *IRandomAccessStream) *IAsyncOperation[*IBitmapDecoder]
	CreateWithIdAsync(decoderId syscall.GUID, stream *IRandomAccessStream) *IAsyncOperation[*IBitmapDecoder]
}

type IBitmapDecoderStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_BmpDecoderId                uintptr
	Get_JpegDecoderId               uintptr
	Get_PngDecoderId                uintptr
	Get_TiffDecoderId               uintptr
	Get_GifDecoderId                uintptr
	Get_JpegXRDecoderId             uintptr
	Get_IcoDecoderId                uintptr
	GetDecoderInformationEnumerator uintptr
	CreateAsync                     uintptr
	CreateWithIdAsync               uintptr
}

type IBitmapDecoderStatics struct {
	win32.IInspectable
}

func (this *IBitmapDecoderStatics) Vtbl() *IBitmapDecoderStaticsVtbl {
	return (*IBitmapDecoderStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBitmapDecoderStatics) Get_BmpDecoderId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BmpDecoderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapDecoderStatics) Get_JpegDecoderId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_JpegDecoderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapDecoderStatics) Get_PngDecoderId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PngDecoderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapDecoderStatics) Get_TiffDecoderId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TiffDecoderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapDecoderStatics) Get_GifDecoderId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GifDecoderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapDecoderStatics) Get_JpegXRDecoderId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_JpegXRDecoderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapDecoderStatics) Get_IcoDecoderId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IcoDecoderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapDecoderStatics) GetDecoderInformationEnumerator() *IVectorView[*IBitmapCodecInformation] {
	var _result *IVectorView[*IBitmapCodecInformation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDecoderInformationEnumerator, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBitmapDecoderStatics) CreateAsync(stream *IRandomAccessStream) *IAsyncOperation[*IBitmapDecoder] {
	var _result *IAsyncOperation[*IBitmapDecoder]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBitmapDecoderStatics) CreateWithIdAsync(decoderId syscall.GUID, stream *IRandomAccessStream) *IAsyncOperation[*IBitmapDecoder] {
	var _result *IAsyncOperation[*IBitmapDecoder]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithIdAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&decoderId)), uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 50BA68EA-99A1-40C4-80D9-AEF0DAFA6C3F
var IID_IBitmapDecoderStatics2 = syscall.GUID{0x50BA68EA, 0x99A1, 0x40C4,
	[8]byte{0x80, 0xD9, 0xAE, 0xF0, 0xDA, 0xFA, 0x6C, 0x3F}}

type IBitmapDecoderStatics2Interface interface {
	win32.IInspectableInterface
	Get_HeifDecoderId() syscall.GUID
	Get_WebpDecoderId() syscall.GUID
}

type IBitmapDecoderStatics2Vtbl struct {
	win32.IInspectableVtbl
	Get_HeifDecoderId uintptr
	Get_WebpDecoderId uintptr
}

type IBitmapDecoderStatics2 struct {
	win32.IInspectable
}

func (this *IBitmapDecoderStatics2) Vtbl() *IBitmapDecoderStatics2Vtbl {
	return (*IBitmapDecoderStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBitmapDecoderStatics2) Get_HeifDecoderId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HeifDecoderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapDecoderStatics2) Get_WebpDecoderId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WebpDecoderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 2BC468E3-E1F8-4B54-95E8-32919551CE62
var IID_IBitmapEncoder = syscall.GUID{0x2BC468E3, 0xE1F8, 0x4B54,
	[8]byte{0x95, 0xE8, 0x32, 0x91, 0x95, 0x51, 0xCE, 0x62}}

type IBitmapEncoderInterface interface {
	win32.IInspectableInterface
	Get_EncoderInformation() *IBitmapCodecInformation
	Get_BitmapProperties() *IBitmapProperties
	Get_BitmapContainerProperties() *IBitmapProperties
	Get_IsThumbnailGenerated() bool
	Put_IsThumbnailGenerated(value bool)
	Get_GeneratedThumbnailWidth() uint32
	Put_GeneratedThumbnailWidth(value uint32)
	Get_GeneratedThumbnailHeight() uint32
	Put_GeneratedThumbnailHeight(value uint32)
	Get_BitmapTransform() *IBitmapTransform
	SetPixelData(pixelFormat BitmapPixelFormat, alphaMode BitmapAlphaMode, width uint32, height uint32, dpiX float64, dpiY float64, pixelsLength uint32, pixels *byte)
	GoToNextFrameAsync() *IAsyncAction
	GoToNextFrameWithEncodingOptionsAsync(encodingOptions *IIterable[*IKeyValuePair[string, *IBitmapTypedValue]]) *IAsyncAction
	FlushAsync() *IAsyncAction
}

type IBitmapEncoderVtbl struct {
	win32.IInspectableVtbl
	Get_EncoderInformation                uintptr
	Get_BitmapProperties                  uintptr
	Get_BitmapContainerProperties         uintptr
	Get_IsThumbnailGenerated              uintptr
	Put_IsThumbnailGenerated              uintptr
	Get_GeneratedThumbnailWidth           uintptr
	Put_GeneratedThumbnailWidth           uintptr
	Get_GeneratedThumbnailHeight          uintptr
	Put_GeneratedThumbnailHeight          uintptr
	Get_BitmapTransform                   uintptr
	SetPixelData                          uintptr
	GoToNextFrameAsync                    uintptr
	GoToNextFrameWithEncodingOptionsAsync uintptr
	FlushAsync                            uintptr
}

type IBitmapEncoder struct {
	win32.IInspectable
}

func (this *IBitmapEncoder) Vtbl() *IBitmapEncoderVtbl {
	return (*IBitmapEncoderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBitmapEncoder) Get_EncoderInformation() *IBitmapCodecInformation {
	var _result *IBitmapCodecInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EncoderInformation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBitmapEncoder) Get_BitmapProperties() *IBitmapProperties {
	var _result *IBitmapProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BitmapProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBitmapEncoder) Get_BitmapContainerProperties() *IBitmapProperties {
	var _result *IBitmapProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BitmapContainerProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBitmapEncoder) Get_IsThumbnailGenerated() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsThumbnailGenerated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapEncoder) Put_IsThumbnailGenerated(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsThumbnailGenerated, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IBitmapEncoder) Get_GeneratedThumbnailWidth() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GeneratedThumbnailWidth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapEncoder) Put_GeneratedThumbnailWidth(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_GeneratedThumbnailWidth, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IBitmapEncoder) Get_GeneratedThumbnailHeight() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GeneratedThumbnailHeight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapEncoder) Put_GeneratedThumbnailHeight(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_GeneratedThumbnailHeight, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IBitmapEncoder) Get_BitmapTransform() *IBitmapTransform {
	var _result *IBitmapTransform
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BitmapTransform, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBitmapEncoder) SetPixelData(pixelFormat BitmapPixelFormat, alphaMode BitmapAlphaMode, width uint32, height uint32, dpiX float64, dpiY float64, pixelsLength uint32, pixels *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetPixelData, uintptr(unsafe.Pointer(this)), uintptr(pixelFormat), uintptr(alphaMode), uintptr(width), uintptr(height), uintptr(dpiX), uintptr(dpiY), uintptr(pixelsLength), uintptr(unsafe.Pointer(pixels)))
	_ = _hr
}

func (this *IBitmapEncoder) GoToNextFrameAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GoToNextFrameAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBitmapEncoder) GoToNextFrameWithEncodingOptionsAsync(encodingOptions *IIterable[*IKeyValuePair[string, *IBitmapTypedValue]]) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GoToNextFrameWithEncodingOptionsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(encodingOptions)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBitmapEncoder) FlushAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FlushAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A74356A7-A4E4-4EB9-8E40-564DE7E1CCB2
var IID_IBitmapEncoderStatics = syscall.GUID{0xA74356A7, 0xA4E4, 0x4EB9,
	[8]byte{0x8E, 0x40, 0x56, 0x4D, 0xE7, 0xE1, 0xCC, 0xB2}}

type IBitmapEncoderStaticsInterface interface {
	win32.IInspectableInterface
	Get_BmpEncoderId() syscall.GUID
	Get_JpegEncoderId() syscall.GUID
	Get_PngEncoderId() syscall.GUID
	Get_TiffEncoderId() syscall.GUID
	Get_GifEncoderId() syscall.GUID
	Get_JpegXREncoderId() syscall.GUID
	GetEncoderInformationEnumerator() *IVectorView[*IBitmapCodecInformation]
	CreateAsync(encoderId syscall.GUID, stream *IRandomAccessStream) *IAsyncOperation[*IBitmapEncoder]
	CreateWithEncodingOptionsAsync(encoderId syscall.GUID, stream *IRandomAccessStream, encodingOptions *IIterable[*IKeyValuePair[string, *IBitmapTypedValue]]) *IAsyncOperation[*IBitmapEncoder]
	CreateForTranscodingAsync(stream *IRandomAccessStream, bitmapDecoder *IBitmapDecoder) *IAsyncOperation[*IBitmapEncoder]
	CreateForInPlacePropertyEncodingAsync(bitmapDecoder *IBitmapDecoder) *IAsyncOperation[*IBitmapEncoder]
}

type IBitmapEncoderStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_BmpEncoderId                      uintptr
	Get_JpegEncoderId                     uintptr
	Get_PngEncoderId                      uintptr
	Get_TiffEncoderId                     uintptr
	Get_GifEncoderId                      uintptr
	Get_JpegXREncoderId                   uintptr
	GetEncoderInformationEnumerator       uintptr
	CreateAsync                           uintptr
	CreateWithEncodingOptionsAsync        uintptr
	CreateForTranscodingAsync             uintptr
	CreateForInPlacePropertyEncodingAsync uintptr
}

type IBitmapEncoderStatics struct {
	win32.IInspectable
}

func (this *IBitmapEncoderStatics) Vtbl() *IBitmapEncoderStaticsVtbl {
	return (*IBitmapEncoderStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBitmapEncoderStatics) Get_BmpEncoderId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BmpEncoderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapEncoderStatics) Get_JpegEncoderId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_JpegEncoderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapEncoderStatics) Get_PngEncoderId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PngEncoderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapEncoderStatics) Get_TiffEncoderId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TiffEncoderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapEncoderStatics) Get_GifEncoderId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GifEncoderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapEncoderStatics) Get_JpegXREncoderId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_JpegXREncoderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapEncoderStatics) GetEncoderInformationEnumerator() *IVectorView[*IBitmapCodecInformation] {
	var _result *IVectorView[*IBitmapCodecInformation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetEncoderInformationEnumerator, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBitmapEncoderStatics) CreateAsync(encoderId syscall.GUID, stream *IRandomAccessStream) *IAsyncOperation[*IBitmapEncoder] {
	var _result *IAsyncOperation[*IBitmapEncoder]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&encoderId)), uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBitmapEncoderStatics) CreateWithEncodingOptionsAsync(encoderId syscall.GUID, stream *IRandomAccessStream, encodingOptions *IIterable[*IKeyValuePair[string, *IBitmapTypedValue]]) *IAsyncOperation[*IBitmapEncoder] {
	var _result *IAsyncOperation[*IBitmapEncoder]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithEncodingOptionsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&encoderId)), uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(encodingOptions)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBitmapEncoderStatics) CreateForTranscodingAsync(stream *IRandomAccessStream, bitmapDecoder *IBitmapDecoder) *IAsyncOperation[*IBitmapEncoder] {
	var _result *IAsyncOperation[*IBitmapEncoder]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateForTranscodingAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(bitmapDecoder)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBitmapEncoderStatics) CreateForInPlacePropertyEncodingAsync(bitmapDecoder *IBitmapDecoder) *IAsyncOperation[*IBitmapEncoder] {
	var _result *IAsyncOperation[*IBitmapEncoder]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateForInPlacePropertyEncodingAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bitmapDecoder)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 33CBC259-FE31-41B1-B812-086D21E87E16
var IID_IBitmapEncoderStatics2 = syscall.GUID{0x33CBC259, 0xFE31, 0x41B1,
	[8]byte{0xB8, 0x12, 0x08, 0x6D, 0x21, 0xE8, 0x7E, 0x16}}

type IBitmapEncoderStatics2Interface interface {
	win32.IInspectableInterface
	Get_HeifEncoderId() syscall.GUID
}

type IBitmapEncoderStatics2Vtbl struct {
	win32.IInspectableVtbl
	Get_HeifEncoderId uintptr
}

type IBitmapEncoderStatics2 struct {
	win32.IInspectable
}

func (this *IBitmapEncoderStatics2) Vtbl() *IBitmapEncoderStatics2Vtbl {
	return (*IBitmapEncoderStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBitmapEncoderStatics2) Get_HeifEncoderId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HeifEncoderId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 686CD241-4330-4C77-ACE4-0334968B1768
var IID_IBitmapEncoderWithSoftwareBitmap = syscall.GUID{0x686CD241, 0x4330, 0x4C77,
	[8]byte{0xAC, 0xE4, 0x03, 0x34, 0x96, 0x8B, 0x17, 0x68}}

type IBitmapEncoderWithSoftwareBitmapInterface interface {
	win32.IInspectableInterface
	SetSoftwareBitmap(bitmap *ISoftwareBitmap)
}

type IBitmapEncoderWithSoftwareBitmapVtbl struct {
	win32.IInspectableVtbl
	SetSoftwareBitmap uintptr
}

type IBitmapEncoderWithSoftwareBitmap struct {
	win32.IInspectable
}

func (this *IBitmapEncoderWithSoftwareBitmap) Vtbl() *IBitmapEncoderWithSoftwareBitmapVtbl {
	return (*IBitmapEncoderWithSoftwareBitmapVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBitmapEncoderWithSoftwareBitmap) SetSoftwareBitmap(bitmap *ISoftwareBitmap) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetSoftwareBitmap, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bitmap)))
	_ = _hr
}

// 72A49A1C-8081-438D-91BC-94ECFC8185C6
var IID_IBitmapFrame = syscall.GUID{0x72A49A1C, 0x8081, 0x438D,
	[8]byte{0x91, 0xBC, 0x94, 0xEC, 0xFC, 0x81, 0x85, 0xC6}}

type IBitmapFrameInterface interface {
	win32.IInspectableInterface
	GetThumbnailAsync() *IAsyncOperation[*IRandomAccessStreamWithContentType]
	Get_BitmapProperties() *IBitmapPropertiesView
	Get_BitmapPixelFormat() BitmapPixelFormat
	Get_BitmapAlphaMode() BitmapAlphaMode
	Get_DpiX() float64
	Get_DpiY() float64
	Get_PixelWidth() uint32
	Get_PixelHeight() uint32
	Get_OrientedPixelWidth() uint32
	Get_OrientedPixelHeight() uint32
	GetPixelDataAsync() *IAsyncOperation[*IPixelDataProvider]
	GetPixelDataTransformedAsync(pixelFormat BitmapPixelFormat, alphaMode BitmapAlphaMode, transform *IBitmapTransform, exifOrientationMode ExifOrientationMode, colorManagementMode ColorManagementMode) *IAsyncOperation[*IPixelDataProvider]
}

type IBitmapFrameVtbl struct {
	win32.IInspectableVtbl
	GetThumbnailAsync            uintptr
	Get_BitmapProperties         uintptr
	Get_BitmapPixelFormat        uintptr
	Get_BitmapAlphaMode          uintptr
	Get_DpiX                     uintptr
	Get_DpiY                     uintptr
	Get_PixelWidth               uintptr
	Get_PixelHeight              uintptr
	Get_OrientedPixelWidth       uintptr
	Get_OrientedPixelHeight      uintptr
	GetPixelDataAsync            uintptr
	GetPixelDataTransformedAsync uintptr
}

type IBitmapFrame struct {
	win32.IInspectable
}

func (this *IBitmapFrame) Vtbl() *IBitmapFrameVtbl {
	return (*IBitmapFrameVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBitmapFrame) GetThumbnailAsync() *IAsyncOperation[*IRandomAccessStreamWithContentType] {
	var _result *IAsyncOperation[*IRandomAccessStreamWithContentType]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetThumbnailAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBitmapFrame) Get_BitmapProperties() *IBitmapPropertiesView {
	var _result *IBitmapPropertiesView
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BitmapProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBitmapFrame) Get_BitmapPixelFormat() BitmapPixelFormat {
	var _result BitmapPixelFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BitmapPixelFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapFrame) Get_BitmapAlphaMode() BitmapAlphaMode {
	var _result BitmapAlphaMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BitmapAlphaMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapFrame) Get_DpiX() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DpiX, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapFrame) Get_DpiY() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DpiY, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapFrame) Get_PixelWidth() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PixelWidth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapFrame) Get_PixelHeight() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PixelHeight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapFrame) Get_OrientedPixelWidth() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OrientedPixelWidth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapFrame) Get_OrientedPixelHeight() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OrientedPixelHeight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapFrame) GetPixelDataAsync() *IAsyncOperation[*IPixelDataProvider] {
	var _result *IAsyncOperation[*IPixelDataProvider]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPixelDataAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBitmapFrame) GetPixelDataTransformedAsync(pixelFormat BitmapPixelFormat, alphaMode BitmapAlphaMode, transform *IBitmapTransform, exifOrientationMode ExifOrientationMode, colorManagementMode ColorManagementMode) *IAsyncOperation[*IPixelDataProvider] {
	var _result *IAsyncOperation[*IPixelDataProvider]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPixelDataTransformedAsync, uintptr(unsafe.Pointer(this)), uintptr(pixelFormat), uintptr(alphaMode), uintptr(unsafe.Pointer(transform)), uintptr(exifOrientationMode), uintptr(colorManagementMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FE287C9A-420C-4963-87AD-691436E08383
var IID_IBitmapFrameWithSoftwareBitmap = syscall.GUID{0xFE287C9A, 0x420C, 0x4963,
	[8]byte{0x87, 0xAD, 0x69, 0x14, 0x36, 0xE0, 0x83, 0x83}}

type IBitmapFrameWithSoftwareBitmapInterface interface {
	win32.IInspectableInterface
	GetSoftwareBitmapAsync() *IAsyncOperation[*ISoftwareBitmap]
	GetSoftwareBitmapConvertedAsync(pixelFormat BitmapPixelFormat, alphaMode BitmapAlphaMode) *IAsyncOperation[*ISoftwareBitmap]
	GetSoftwareBitmapTransformedAsync(pixelFormat BitmapPixelFormat, alphaMode BitmapAlphaMode, transform *IBitmapTransform, exifOrientationMode ExifOrientationMode, colorManagementMode ColorManagementMode) *IAsyncOperation[*ISoftwareBitmap]
}

type IBitmapFrameWithSoftwareBitmapVtbl struct {
	win32.IInspectableVtbl
	GetSoftwareBitmapAsync            uintptr
	GetSoftwareBitmapConvertedAsync   uintptr
	GetSoftwareBitmapTransformedAsync uintptr
}

type IBitmapFrameWithSoftwareBitmap struct {
	win32.IInspectable
}

func (this *IBitmapFrameWithSoftwareBitmap) Vtbl() *IBitmapFrameWithSoftwareBitmapVtbl {
	return (*IBitmapFrameWithSoftwareBitmapVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBitmapFrameWithSoftwareBitmap) GetSoftwareBitmapAsync() *IAsyncOperation[*ISoftwareBitmap] {
	var _result *IAsyncOperation[*ISoftwareBitmap]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSoftwareBitmapAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBitmapFrameWithSoftwareBitmap) GetSoftwareBitmapConvertedAsync(pixelFormat BitmapPixelFormat, alphaMode BitmapAlphaMode) *IAsyncOperation[*ISoftwareBitmap] {
	var _result *IAsyncOperation[*ISoftwareBitmap]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSoftwareBitmapConvertedAsync, uintptr(unsafe.Pointer(this)), uintptr(pixelFormat), uintptr(alphaMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBitmapFrameWithSoftwareBitmap) GetSoftwareBitmapTransformedAsync(pixelFormat BitmapPixelFormat, alphaMode BitmapAlphaMode, transform *IBitmapTransform, exifOrientationMode ExifOrientationMode, colorManagementMode ColorManagementMode) *IAsyncOperation[*ISoftwareBitmap] {
	var _result *IAsyncOperation[*ISoftwareBitmap]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSoftwareBitmapTransformedAsync, uintptr(unsafe.Pointer(this)), uintptr(pixelFormat), uintptr(alphaMode), uintptr(unsafe.Pointer(transform)), uintptr(exifOrientationMode), uintptr(colorManagementMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EA9F4F1B-B505-4450-A4D1-E8CA94529D8D
var IID_IBitmapProperties = syscall.GUID{0xEA9F4F1B, 0xB505, 0x4450,
	[8]byte{0xA4, 0xD1, 0xE8, 0xCA, 0x94, 0x52, 0x9D, 0x8D}}

type IBitmapPropertiesInterface interface {
	win32.IInspectableInterface
	SetPropertiesAsync(propertiesToSet *IIterable[*IKeyValuePair[string, *IBitmapTypedValue]]) *IAsyncAction
}

type IBitmapPropertiesVtbl struct {
	win32.IInspectableVtbl
	SetPropertiesAsync uintptr
}

type IBitmapProperties struct {
	win32.IInspectable
}

func (this *IBitmapProperties) Vtbl() *IBitmapPropertiesVtbl {
	return (*IBitmapPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBitmapProperties) SetPropertiesAsync(propertiesToSet *IIterable[*IKeyValuePair[string, *IBitmapTypedValue]]) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetPropertiesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(propertiesToSet)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7E0FE87A-3A70-48F8-9C55-196CF5A545F5
var IID_IBitmapPropertiesView = syscall.GUID{0x7E0FE87A, 0x3A70, 0x48F8,
	[8]byte{0x9C, 0x55, 0x19, 0x6C, 0xF5, 0xA5, 0x45, 0xF5}}

type IBitmapPropertiesViewInterface interface {
	win32.IInspectableInterface
	GetPropertiesAsync(propertiesToRetrieve *IIterable[string]) *IAsyncOperation[*IMap[string, *IBitmapTypedValue]]
}

type IBitmapPropertiesViewVtbl struct {
	win32.IInspectableVtbl
	GetPropertiesAsync uintptr
}

type IBitmapPropertiesView struct {
	win32.IInspectable
}

func (this *IBitmapPropertiesView) Vtbl() *IBitmapPropertiesViewVtbl {
	return (*IBitmapPropertiesViewVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBitmapPropertiesView) GetPropertiesAsync(propertiesToRetrieve *IIterable[string]) *IAsyncOperation[*IMap[string, *IBitmapTypedValue]] {
	var _result *IAsyncOperation[*IMap[string, *IBitmapTypedValue]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPropertiesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(propertiesToRetrieve)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AE755344-E268-4D35-ADCF-E995D31A8D34
var IID_IBitmapTransform = syscall.GUID{0xAE755344, 0xE268, 0x4D35,
	[8]byte{0xAD, 0xCF, 0xE9, 0x95, 0xD3, 0x1A, 0x8D, 0x34}}

type IBitmapTransformInterface interface {
	win32.IInspectableInterface
	Get_ScaledWidth() uint32
	Put_ScaledWidth(value uint32)
	Get_ScaledHeight() uint32
	Put_ScaledHeight(value uint32)
	Get_InterpolationMode() BitmapInterpolationMode
	Put_InterpolationMode(value BitmapInterpolationMode)
	Get_Flip() BitmapFlip
	Put_Flip(value BitmapFlip)
	Get_Rotation() BitmapRotation
	Put_Rotation(value BitmapRotation)
	Get_Bounds() BitmapBounds
	Put_Bounds(value BitmapBounds)
}

type IBitmapTransformVtbl struct {
	win32.IInspectableVtbl
	Get_ScaledWidth       uintptr
	Put_ScaledWidth       uintptr
	Get_ScaledHeight      uintptr
	Put_ScaledHeight      uintptr
	Get_InterpolationMode uintptr
	Put_InterpolationMode uintptr
	Get_Flip              uintptr
	Put_Flip              uintptr
	Get_Rotation          uintptr
	Put_Rotation          uintptr
	Get_Bounds            uintptr
	Put_Bounds            uintptr
}

type IBitmapTransform struct {
	win32.IInspectable
}

func (this *IBitmapTransform) Vtbl() *IBitmapTransformVtbl {
	return (*IBitmapTransformVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBitmapTransform) Get_ScaledWidth() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScaledWidth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapTransform) Put_ScaledWidth(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ScaledWidth, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IBitmapTransform) Get_ScaledHeight() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScaledHeight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapTransform) Put_ScaledHeight(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ScaledHeight, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IBitmapTransform) Get_InterpolationMode() BitmapInterpolationMode {
	var _result BitmapInterpolationMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InterpolationMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapTransform) Put_InterpolationMode(value BitmapInterpolationMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InterpolationMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IBitmapTransform) Get_Flip() BitmapFlip {
	var _result BitmapFlip
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Flip, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapTransform) Put_Flip(value BitmapFlip) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Flip, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IBitmapTransform) Get_Rotation() BitmapRotation {
	var _result BitmapRotation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Rotation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapTransform) Put_Rotation(value BitmapRotation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Rotation, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IBitmapTransform) Get_Bounds() BitmapBounds {
	var _result BitmapBounds
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bounds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapTransform) Put_Bounds(value BitmapBounds) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Bounds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

// CD8044A9-2443-4000-B0CD-79316C56F589
var IID_IBitmapTypedValue = syscall.GUID{0xCD8044A9, 0x2443, 0x4000,
	[8]byte{0xB0, 0xCD, 0x79, 0x31, 0x6C, 0x56, 0xF5, 0x89}}

type IBitmapTypedValueInterface interface {
	win32.IInspectableInterface
	Get_Value() interface{}
	Get_Type() PropertyType
}

type IBitmapTypedValueVtbl struct {
	win32.IInspectableVtbl
	Get_Value uintptr
	Get_Type  uintptr
}

type IBitmapTypedValue struct {
	win32.IInspectable
}

func (this *IBitmapTypedValue) Vtbl() *IBitmapTypedValueVtbl {
	return (*IBitmapTypedValueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBitmapTypedValue) Get_Value() interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBitmapTypedValue) Get_Type() PropertyType {
	var _result PropertyType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Type, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 92DBB599-CE13-46BB-9545-CB3A3F63EB8B
var IID_IBitmapTypedValueFactory = syscall.GUID{0x92DBB599, 0xCE13, 0x46BB,
	[8]byte{0x95, 0x45, 0xCB, 0x3A, 0x3F, 0x63, 0xEB, 0x8B}}

type IBitmapTypedValueFactoryInterface interface {
	win32.IInspectableInterface
	Create(value interface{}, type_ PropertyType) *IBitmapTypedValue
}

type IBitmapTypedValueFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IBitmapTypedValueFactory struct {
	win32.IInspectable
}

func (this *IBitmapTypedValueFactory) Vtbl() *IBitmapTypedValueFactoryVtbl {
	return (*IBitmapTypedValueFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBitmapTypedValueFactory) Create(value interface{}, type_ PropertyType) *IBitmapTypedValue {
	var _result *IBitmapTypedValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)), uintptr(type_), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DD831F25-185C-4595-9FB9-CCBE6EC18A6F
var IID_IPixelDataProvider = syscall.GUID{0xDD831F25, 0x185C, 0x4595,
	[8]byte{0x9F, 0xB9, 0xCC, 0xBE, 0x6E, 0xC1, 0x8A, 0x6F}}

type IPixelDataProviderInterface interface {
	win32.IInspectableInterface
	DetachPixelData() []byte
}

type IPixelDataProviderVtbl struct {
	win32.IInspectableVtbl
	DetachPixelData uintptr
}

type IPixelDataProvider struct {
	win32.IInspectable
}

func (this *IPixelDataProvider) Vtbl() *IPixelDataProviderVtbl {
	return (*IPixelDataProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPixelDataProvider) DetachPixelData() []byte {
	var _result []byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DetachPixelData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 689E0708-7EEF-483F-963F-DA938818E073
var IID_ISoftwareBitmap = syscall.GUID{0x689E0708, 0x7EEF, 0x483F,
	[8]byte{0x96, 0x3F, 0xDA, 0x93, 0x88, 0x18, 0xE0, 0x73}}

type ISoftwareBitmapInterface interface {
	win32.IInspectableInterface
	Get_BitmapPixelFormat() BitmapPixelFormat
	Get_BitmapAlphaMode() BitmapAlphaMode
	Get_PixelWidth() int32
	Get_PixelHeight() int32
	Get_IsReadOnly() bool
	Put_DpiX(value float64)
	Get_DpiX() float64
	Put_DpiY(value float64)
	Get_DpiY() float64
	LockBuffer(mode BitmapBufferAccessMode) *IBitmapBuffer
	CopyTo(bitmap *ISoftwareBitmap)
	CopyFromBuffer(buffer *IBuffer)
	CopyToBuffer(buffer *IBuffer)
	GetReadOnlyView() *ISoftwareBitmap
}

type ISoftwareBitmapVtbl struct {
	win32.IInspectableVtbl
	Get_BitmapPixelFormat uintptr
	Get_BitmapAlphaMode   uintptr
	Get_PixelWidth        uintptr
	Get_PixelHeight       uintptr
	Get_IsReadOnly        uintptr
	Put_DpiX              uintptr
	Get_DpiX              uintptr
	Put_DpiY              uintptr
	Get_DpiY              uintptr
	LockBuffer            uintptr
	CopyTo                uintptr
	CopyFromBuffer        uintptr
	CopyToBuffer          uintptr
	GetReadOnlyView       uintptr
}

type ISoftwareBitmap struct {
	win32.IInspectable
}

func (this *ISoftwareBitmap) Vtbl() *ISoftwareBitmapVtbl {
	return (*ISoftwareBitmapVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISoftwareBitmap) Get_BitmapPixelFormat() BitmapPixelFormat {
	var _result BitmapPixelFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BitmapPixelFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISoftwareBitmap) Get_BitmapAlphaMode() BitmapAlphaMode {
	var _result BitmapAlphaMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BitmapAlphaMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISoftwareBitmap) Get_PixelWidth() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PixelWidth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISoftwareBitmap) Get_PixelHeight() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PixelHeight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISoftwareBitmap) Get_IsReadOnly() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsReadOnly, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISoftwareBitmap) Put_DpiX(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DpiX, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISoftwareBitmap) Get_DpiX() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DpiX, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISoftwareBitmap) Put_DpiY(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DpiY, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISoftwareBitmap) Get_DpiY() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DpiY, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISoftwareBitmap) LockBuffer(mode BitmapBufferAccessMode) *IBitmapBuffer {
	var _result *IBitmapBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LockBuffer, uintptr(unsafe.Pointer(this)), uintptr(mode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISoftwareBitmap) CopyTo(bitmap *ISoftwareBitmap) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CopyTo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bitmap)))
	_ = _hr
}

func (this *ISoftwareBitmap) CopyFromBuffer(buffer *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CopyFromBuffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(buffer)))
	_ = _hr
}

func (this *ISoftwareBitmap) CopyToBuffer(buffer *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CopyToBuffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(buffer)))
	_ = _hr
}

func (this *ISoftwareBitmap) GetReadOnlyView() *ISoftwareBitmap {
	var _result *ISoftwareBitmap
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetReadOnlyView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C99FEB69-2D62-4D47-A6B3-4FDB6A07FDF8
var IID_ISoftwareBitmapFactory = syscall.GUID{0xC99FEB69, 0x2D62, 0x4D47,
	[8]byte{0xA6, 0xB3, 0x4F, 0xDB, 0x6A, 0x07, 0xFD, 0xF8}}

type ISoftwareBitmapFactoryInterface interface {
	win32.IInspectableInterface
	Create(format BitmapPixelFormat, width int32, height int32) *ISoftwareBitmap
	CreateWithAlpha(format BitmapPixelFormat, width int32, height int32, alpha BitmapAlphaMode) *ISoftwareBitmap
}

type ISoftwareBitmapFactoryVtbl struct {
	win32.IInspectableVtbl
	Create          uintptr
	CreateWithAlpha uintptr
}

type ISoftwareBitmapFactory struct {
	win32.IInspectable
}

func (this *ISoftwareBitmapFactory) Vtbl() *ISoftwareBitmapFactoryVtbl {
	return (*ISoftwareBitmapFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISoftwareBitmapFactory) Create(format BitmapPixelFormat, width int32, height int32) *ISoftwareBitmap {
	var _result *ISoftwareBitmap
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(format), uintptr(width), uintptr(height), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISoftwareBitmapFactory) CreateWithAlpha(format BitmapPixelFormat, width int32, height int32, alpha BitmapAlphaMode) *ISoftwareBitmap {
	var _result *ISoftwareBitmap
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithAlpha, uintptr(unsafe.Pointer(this)), uintptr(format), uintptr(width), uintptr(height), uintptr(alpha), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DF0385DB-672F-4A9D-806E-C2442F343E86
var IID_ISoftwareBitmapStatics = syscall.GUID{0xDF0385DB, 0x672F, 0x4A9D,
	[8]byte{0x80, 0x6E, 0xC2, 0x44, 0x2F, 0x34, 0x3E, 0x86}}

type ISoftwareBitmapStaticsInterface interface {
	win32.IInspectableInterface
	Copy(source *ISoftwareBitmap) *ISoftwareBitmap
	Convert(source *ISoftwareBitmap, format BitmapPixelFormat) *ISoftwareBitmap
	ConvertWithAlpha(source *ISoftwareBitmap, format BitmapPixelFormat, alpha BitmapAlphaMode) *ISoftwareBitmap
	CreateCopyFromBuffer(source *IBuffer, format BitmapPixelFormat, width int32, height int32) *ISoftwareBitmap
	CreateCopyWithAlphaFromBuffer(source *IBuffer, format BitmapPixelFormat, width int32, height int32, alpha BitmapAlphaMode) *ISoftwareBitmap
	CreateCopyFromSurfaceAsync(surface unsafe.Pointer) *IAsyncOperation[*ISoftwareBitmap]
	CreateCopyWithAlphaFromSurfaceAsync(surface unsafe.Pointer, alpha BitmapAlphaMode) *IAsyncOperation[*ISoftwareBitmap]
}

type ISoftwareBitmapStaticsVtbl struct {
	win32.IInspectableVtbl
	Copy                                uintptr
	Convert                             uintptr
	ConvertWithAlpha                    uintptr
	CreateCopyFromBuffer                uintptr
	CreateCopyWithAlphaFromBuffer       uintptr
	CreateCopyFromSurfaceAsync          uintptr
	CreateCopyWithAlphaFromSurfaceAsync uintptr
}

type ISoftwareBitmapStatics struct {
	win32.IInspectable
}

func (this *ISoftwareBitmapStatics) Vtbl() *ISoftwareBitmapStaticsVtbl {
	return (*ISoftwareBitmapStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISoftwareBitmapStatics) Copy(source *ISoftwareBitmap) *ISoftwareBitmap {
	var _result *ISoftwareBitmap
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Copy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(source)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISoftwareBitmapStatics) Convert(source *ISoftwareBitmap, format BitmapPixelFormat) *ISoftwareBitmap {
	var _result *ISoftwareBitmap
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Convert, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(source)), uintptr(format), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISoftwareBitmapStatics) ConvertWithAlpha(source *ISoftwareBitmap, format BitmapPixelFormat, alpha BitmapAlphaMode) *ISoftwareBitmap {
	var _result *ISoftwareBitmap
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConvertWithAlpha, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(source)), uintptr(format), uintptr(alpha), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISoftwareBitmapStatics) CreateCopyFromBuffer(source *IBuffer, format BitmapPixelFormat, width int32, height int32) *ISoftwareBitmap {
	var _result *ISoftwareBitmap
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateCopyFromBuffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(source)), uintptr(format), uintptr(width), uintptr(height), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISoftwareBitmapStatics) CreateCopyWithAlphaFromBuffer(source *IBuffer, format BitmapPixelFormat, width int32, height int32, alpha BitmapAlphaMode) *ISoftwareBitmap {
	var _result *ISoftwareBitmap
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateCopyWithAlphaFromBuffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(source)), uintptr(format), uintptr(width), uintptr(height), uintptr(alpha), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISoftwareBitmapStatics) CreateCopyFromSurfaceAsync(surface unsafe.Pointer) *IAsyncOperation[*ISoftwareBitmap] {
	var _result *IAsyncOperation[*ISoftwareBitmap]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateCopyFromSurfaceAsync, uintptr(unsafe.Pointer(this)), uintptr(surface), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISoftwareBitmapStatics) CreateCopyWithAlphaFromSurfaceAsync(surface unsafe.Pointer, alpha BitmapAlphaMode) *IAsyncOperation[*ISoftwareBitmap] {
	var _result *IAsyncOperation[*ISoftwareBitmap]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateCopyWithAlphaFromSurfaceAsync, uintptr(unsafe.Pointer(this)), uintptr(surface), uintptr(alpha), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type BitmapBuffer struct {
	RtClass
	*IBitmapBuffer
}

type SoftwareBitmap struct {
	RtClass
	*ISoftwareBitmap
}

func NewSoftwareBitmap_Create(format BitmapPixelFormat, width int32, height int32) *SoftwareBitmap {
	hs := NewHStr("Windows.Graphics.Imaging.SoftwareBitmap")
	var pFac *ISoftwareBitmapFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISoftwareBitmapFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISoftwareBitmap
	p = pFac.Create(format, width, height)
	result := &SoftwareBitmap{
		RtClass:         RtClass{PInspect: &p.IInspectable},
		ISoftwareBitmap: p,
	}
	com.AddToScope(result)
	return result
}

func NewSoftwareBitmap_CreateWithAlpha(format BitmapPixelFormat, width int32, height int32, alpha BitmapAlphaMode) *SoftwareBitmap {
	hs := NewHStr("Windows.Graphics.Imaging.SoftwareBitmap")
	var pFac *ISoftwareBitmapFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISoftwareBitmapFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISoftwareBitmap
	p = pFac.CreateWithAlpha(format, width, height, alpha)
	result := &SoftwareBitmap{
		RtClass:         RtClass{PInspect: &p.IInspectable},
		ISoftwareBitmap: p,
	}
	com.AddToScope(result)
	return result
}

func NewISoftwareBitmapStatics() *ISoftwareBitmapStatics {
	var p *ISoftwareBitmapStatics
	hs := NewHStr("Windows.Graphics.Imaging.SoftwareBitmap")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISoftwareBitmapStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
