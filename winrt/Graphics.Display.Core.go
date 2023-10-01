package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type HdmiDisplayColorSpace int32

const (
	HdmiDisplayColorSpace_RgbLimited HdmiDisplayColorSpace = 0
	HdmiDisplayColorSpace_RgbFull    HdmiDisplayColorSpace = 1
	HdmiDisplayColorSpace_BT2020     HdmiDisplayColorSpace = 2
	HdmiDisplayColorSpace_BT709      HdmiDisplayColorSpace = 3
)

// enum
type HdmiDisplayHdrOption int32

const (
	HdmiDisplayHdrOption_None                  HdmiDisplayHdrOption = 0
	HdmiDisplayHdrOption_EotfSdr               HdmiDisplayHdrOption = 1
	HdmiDisplayHdrOption_Eotf2084              HdmiDisplayHdrOption = 2
	HdmiDisplayHdrOption_DolbyVisionLowLatency HdmiDisplayHdrOption = 3
)

// enum
type HdmiDisplayPixelEncoding int32

const (
	HdmiDisplayPixelEncoding_Rgb444 HdmiDisplayPixelEncoding = 0
	HdmiDisplayPixelEncoding_Ycc444 HdmiDisplayPixelEncoding = 1
	HdmiDisplayPixelEncoding_Ycc422 HdmiDisplayPixelEncoding = 2
	HdmiDisplayPixelEncoding_Ycc420 HdmiDisplayPixelEncoding = 3
)

// structs

type HdmiDisplayHdr2086Metadata struct {
	RedPrimaryX               uint16
	RedPrimaryY               uint16
	GreenPrimaryX             uint16
	GreenPrimaryY             uint16
	BluePrimaryX              uint16
	BluePrimaryY              uint16
	WhitePointX               uint16
	WhitePointY               uint16
	MaxMasteringLuminance     uint16
	MinMasteringLuminance     uint16
	MaxContentLightLevel      uint16
	MaxFrameAverageLightLevel uint16
}

// interfaces

// 130B3C0A-F565-476E-ABD5-EA05AEE74C69
var IID_IHdmiDisplayInformation = syscall.GUID{0x130B3C0A, 0xF565, 0x476E,
	[8]byte{0xAB, 0xD5, 0xEA, 0x05, 0xAE, 0xE7, 0x4C, 0x69}}

type IHdmiDisplayInformationInterface interface {
	win32.IInspectableInterface
	GetSupportedDisplayModes() *IVectorView[*IHdmiDisplayMode]
	GetCurrentDisplayMode() *IHdmiDisplayMode
	SetDefaultDisplayModeAsync() *IAsyncAction
	RequestSetCurrentDisplayModeAsync(mode *IHdmiDisplayMode) *IAsyncOperation[bool]
	RequestSetCurrentDisplayModeWithHdrAsync(mode *IHdmiDisplayMode, hdrOption HdmiDisplayHdrOption) *IAsyncOperation[bool]
	RequestSetCurrentDisplayModeWithHdrAndMetadataAsync(mode *IHdmiDisplayMode, hdrOption HdmiDisplayHdrOption, hdrMetadata HdmiDisplayHdr2086Metadata) *IAsyncOperation[bool]
	Add_DisplayModesChanged(value TypedEventHandler[*IHdmiDisplayInformation, interface{}]) EventRegistrationToken
	Remove_DisplayModesChanged(token EventRegistrationToken)
}

type IHdmiDisplayInformationVtbl struct {
	win32.IInspectableVtbl
	GetSupportedDisplayModes                            uintptr
	GetCurrentDisplayMode                               uintptr
	SetDefaultDisplayModeAsync                          uintptr
	RequestSetCurrentDisplayModeAsync                   uintptr
	RequestSetCurrentDisplayModeWithHdrAsync            uintptr
	RequestSetCurrentDisplayModeWithHdrAndMetadataAsync uintptr
	Add_DisplayModesChanged                             uintptr
	Remove_DisplayModesChanged                          uintptr
}

type IHdmiDisplayInformation struct {
	win32.IInspectable
}

func (this *IHdmiDisplayInformation) Vtbl() *IHdmiDisplayInformationVtbl {
	return (*IHdmiDisplayInformationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHdmiDisplayInformation) GetSupportedDisplayModes() *IVectorView[*IHdmiDisplayMode] {
	var _result *IVectorView[*IHdmiDisplayMode]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSupportedDisplayModes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHdmiDisplayInformation) GetCurrentDisplayMode() *IHdmiDisplayMode {
	var _result *IHdmiDisplayMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentDisplayMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHdmiDisplayInformation) SetDefaultDisplayModeAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetDefaultDisplayModeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHdmiDisplayInformation) RequestSetCurrentDisplayModeAsync(mode *IHdmiDisplayMode) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestSetCurrentDisplayModeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(mode)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHdmiDisplayInformation) RequestSetCurrentDisplayModeWithHdrAsync(mode *IHdmiDisplayMode, hdrOption HdmiDisplayHdrOption) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestSetCurrentDisplayModeWithHdrAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(mode)), uintptr(hdrOption), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHdmiDisplayInformation) RequestSetCurrentDisplayModeWithHdrAndMetadataAsync(mode *IHdmiDisplayMode, hdrOption HdmiDisplayHdrOption, hdrMetadata HdmiDisplayHdr2086Metadata) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestSetCurrentDisplayModeWithHdrAndMetadataAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(mode)), uintptr(hdrOption), uintptr(unsafe.Pointer(&hdrMetadata)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHdmiDisplayInformation) Add_DisplayModesChanged(value TypedEventHandler[*IHdmiDisplayInformation, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DisplayModesChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHdmiDisplayInformation) Remove_DisplayModesChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DisplayModesChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 6CE6B260-F42A-4A15-914C-7B8E2A5A65DF
var IID_IHdmiDisplayInformationStatics = syscall.GUID{0x6CE6B260, 0xF42A, 0x4A15,
	[8]byte{0x91, 0x4C, 0x7B, 0x8E, 0x2A, 0x5A, 0x65, 0xDF}}

type IHdmiDisplayInformationStaticsInterface interface {
	win32.IInspectableInterface
	GetForCurrentView() *IHdmiDisplayInformation
}

type IHdmiDisplayInformationStaticsVtbl struct {
	win32.IInspectableVtbl
	GetForCurrentView uintptr
}

type IHdmiDisplayInformationStatics struct {
	win32.IInspectable
}

func (this *IHdmiDisplayInformationStatics) Vtbl() *IHdmiDisplayInformationStaticsVtbl {
	return (*IHdmiDisplayInformationStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHdmiDisplayInformationStatics) GetForCurrentView() *IHdmiDisplayInformation {
	var _result *IHdmiDisplayInformation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0C06D5AD-1B90-4F51-9981-EF5A1C0DDF66
var IID_IHdmiDisplayMode = syscall.GUID{0x0C06D5AD, 0x1B90, 0x4F51,
	[8]byte{0x99, 0x81, 0xEF, 0x5A, 0x1C, 0x0D, 0xDF, 0x66}}

type IHdmiDisplayModeInterface interface {
	win32.IInspectableInterface
	Get_ResolutionWidthInRawPixels() uint32
	Get_ResolutionHeightInRawPixels() uint32
	Get_RefreshRate() float64
	Get_StereoEnabled() bool
	Get_BitsPerPixel() uint16
	IsEqual(mode *IHdmiDisplayMode) bool
	Get_ColorSpace() HdmiDisplayColorSpace
	Get_PixelEncoding() HdmiDisplayPixelEncoding
	Get_IsSdrLuminanceSupported() bool
	Get_IsSmpte2084Supported() bool
	Get_Is2086MetadataSupported() bool
}

type IHdmiDisplayModeVtbl struct {
	win32.IInspectableVtbl
	Get_ResolutionWidthInRawPixels  uintptr
	Get_ResolutionHeightInRawPixels uintptr
	Get_RefreshRate                 uintptr
	Get_StereoEnabled               uintptr
	Get_BitsPerPixel                uintptr
	IsEqual                         uintptr
	Get_ColorSpace                  uintptr
	Get_PixelEncoding               uintptr
	Get_IsSdrLuminanceSupported     uintptr
	Get_IsSmpte2084Supported        uintptr
	Get_Is2086MetadataSupported     uintptr
}

type IHdmiDisplayMode struct {
	win32.IInspectable
}

func (this *IHdmiDisplayMode) Vtbl() *IHdmiDisplayModeVtbl {
	return (*IHdmiDisplayModeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHdmiDisplayMode) Get_ResolutionWidthInRawPixels() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResolutionWidthInRawPixels, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHdmiDisplayMode) Get_ResolutionHeightInRawPixels() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ResolutionHeightInRawPixels, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHdmiDisplayMode) Get_RefreshRate() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RefreshRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHdmiDisplayMode) Get_StereoEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StereoEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHdmiDisplayMode) Get_BitsPerPixel() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BitsPerPixel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHdmiDisplayMode) IsEqual(mode *IHdmiDisplayMode) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsEqual, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(mode)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHdmiDisplayMode) Get_ColorSpace() HdmiDisplayColorSpace {
	var _result HdmiDisplayColorSpace
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ColorSpace, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHdmiDisplayMode) Get_PixelEncoding() HdmiDisplayPixelEncoding {
	var _result HdmiDisplayPixelEncoding
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PixelEncoding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHdmiDisplayMode) Get_IsSdrLuminanceSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSdrLuminanceSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHdmiDisplayMode) Get_IsSmpte2084Supported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSmpte2084Supported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHdmiDisplayMode) Get_Is2086MetadataSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Is2086MetadataSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 07CD4E9F-4B3C-42B8-84E7-895368718AF2
var IID_IHdmiDisplayMode2 = syscall.GUID{0x07CD4E9F, 0x4B3C, 0x42B8,
	[8]byte{0x84, 0xE7, 0x89, 0x53, 0x68, 0x71, 0x8A, 0xF2}}

type IHdmiDisplayMode2Interface interface {
	win32.IInspectableInterface
	Get_IsDolbyVisionLowLatencySupported() bool
}

type IHdmiDisplayMode2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsDolbyVisionLowLatencySupported uintptr
}

type IHdmiDisplayMode2 struct {
	win32.IInspectable
}

func (this *IHdmiDisplayMode2) Vtbl() *IHdmiDisplayMode2Vtbl {
	return (*IHdmiDisplayMode2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHdmiDisplayMode2) Get_IsDolbyVisionLowLatencySupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDolbyVisionLowLatencySupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type HdmiDisplayInformation struct {
	RtClass
	*IHdmiDisplayInformation
}

func NewIHdmiDisplayInformationStatics() *IHdmiDisplayInformationStatics {
	var p *IHdmiDisplayInformationStatics
	hs := NewHStr("Windows.Graphics.Display.Core.HdmiDisplayInformation")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IHdmiDisplayInformationStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type HdmiDisplayMode struct {
	RtClass
	*IHdmiDisplayMode
}
