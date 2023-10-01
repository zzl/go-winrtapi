package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type ImageScannerAutoCroppingMode int32

const (
	ImageScannerAutoCroppingMode_Disabled       ImageScannerAutoCroppingMode = 0
	ImageScannerAutoCroppingMode_SingleRegion   ImageScannerAutoCroppingMode = 1
	ImageScannerAutoCroppingMode_MultipleRegion ImageScannerAutoCroppingMode = 2
)

// enum
type ImageScannerColorMode int32

const (
	ImageScannerColorMode_Color      ImageScannerColorMode = 0
	ImageScannerColorMode_Grayscale  ImageScannerColorMode = 1
	ImageScannerColorMode_Monochrome ImageScannerColorMode = 2
	ImageScannerColorMode_AutoColor  ImageScannerColorMode = 3
)

// enum
type ImageScannerFormat int32

const (
	ImageScannerFormat_Jpeg                    ImageScannerFormat = 0
	ImageScannerFormat_Png                     ImageScannerFormat = 1
	ImageScannerFormat_DeviceIndependentBitmap ImageScannerFormat = 2
	ImageScannerFormat_Tiff                    ImageScannerFormat = 3
	ImageScannerFormat_Xps                     ImageScannerFormat = 4
	ImageScannerFormat_OpenXps                 ImageScannerFormat = 5
	ImageScannerFormat_Pdf                     ImageScannerFormat = 6
)

// enum
type ImageScannerScanSource int32

const (
	ImageScannerScanSource_Default        ImageScannerScanSource = 0
	ImageScannerScanSource_Flatbed        ImageScannerScanSource = 1
	ImageScannerScanSource_Feeder         ImageScannerScanSource = 2
	ImageScannerScanSource_AutoConfigured ImageScannerScanSource = 3
)

// structs

type ImageScannerResolution struct {
	DpiX float32
	DpiY float32
}

type ScannerDeviceContract struct {
}

// interfaces

// 53A88F78-5298-48A0-8DA3-8087519665E0
var IID_IImageScanner = syscall.GUID{0x53A88F78, 0x5298, 0x48A0,
	[8]byte{0x8D, 0xA3, 0x80, 0x87, 0x51, 0x96, 0x65, 0xE0}}

type IImageScannerInterface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
	Get_DefaultScanSource() ImageScannerScanSource
	IsScanSourceSupported(value ImageScannerScanSource) bool
	Get_FlatbedConfiguration() *IImageScannerFormatConfiguration
	Get_FeederConfiguration() *IImageScannerFormatConfiguration
	Get_AutoConfiguration() *IImageScannerFormatConfiguration
	IsPreviewSupported(scanSource ImageScannerScanSource) bool
	ScanPreviewToStreamAsync(scanSource ImageScannerScanSource, targetStream *IRandomAccessStream) *IAsyncOperation[*IImageScannerPreviewResult]
	ScanFilesToFolderAsync(scanSource ImageScannerScanSource, storageFolder *IStorageFolder) *IAsyncOperationWithProgress[*IImageScannerScanResult, uint32]
}

type IImageScannerVtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId             uintptr
	Get_DefaultScanSource    uintptr
	IsScanSourceSupported    uintptr
	Get_FlatbedConfiguration uintptr
	Get_FeederConfiguration  uintptr
	Get_AutoConfiguration    uintptr
	IsPreviewSupported       uintptr
	ScanPreviewToStreamAsync uintptr
	ScanFilesToFolderAsync   uintptr
}

type IImageScanner struct {
	win32.IInspectable
}

func (this *IImageScanner) Vtbl() *IImageScannerVtbl {
	return (*IImageScannerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IImageScanner) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IImageScanner) Get_DefaultScanSource() ImageScannerScanSource {
	var _result ImageScannerScanSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DefaultScanSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScanner) IsScanSourceSupported(value ImageScannerScanSource) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsScanSourceSupported, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScanner) Get_FlatbedConfiguration() *IImageScannerFormatConfiguration {
	var _result *IImageScannerFormatConfiguration
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FlatbedConfiguration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IImageScanner) Get_FeederConfiguration() *IImageScannerFormatConfiguration {
	var _result *IImageScannerFormatConfiguration
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FeederConfiguration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IImageScanner) Get_AutoConfiguration() *IImageScannerFormatConfiguration {
	var _result *IImageScannerFormatConfiguration
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AutoConfiguration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IImageScanner) IsPreviewSupported(scanSource ImageScannerScanSource) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsPreviewSupported, uintptr(unsafe.Pointer(this)), uintptr(scanSource), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScanner) ScanPreviewToStreamAsync(scanSource ImageScannerScanSource, targetStream *IRandomAccessStream) *IAsyncOperation[*IImageScannerPreviewResult] {
	var _result *IAsyncOperation[*IImageScannerPreviewResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ScanPreviewToStreamAsync, uintptr(unsafe.Pointer(this)), uintptr(scanSource), uintptr(unsafe.Pointer(targetStream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IImageScanner) ScanFilesToFolderAsync(scanSource ImageScannerScanSource, storageFolder *IStorageFolder) *IAsyncOperationWithProgress[*IImageScannerScanResult, uint32] {
	var _result *IAsyncOperationWithProgress[*IImageScannerScanResult, uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ScanFilesToFolderAsync, uintptr(unsafe.Pointer(this)), uintptr(scanSource), uintptr(unsafe.Pointer(storageFolder)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 74BDACEE-FA97-4C17-8280-40E39C6DCC67
var IID_IImageScannerFeederConfiguration = syscall.GUID{0x74BDACEE, 0xFA97, 0x4C17,
	[8]byte{0x82, 0x80, 0x40, 0xE3, 0x9C, 0x6D, 0xCC, 0x67}}

type IImageScannerFeederConfigurationInterface interface {
	win32.IInspectableInterface
	Get_CanAutoDetectPageSize() bool
	Get_AutoDetectPageSize() bool
	Put_AutoDetectPageSize(value bool)
	Get_PageSize() PrintMediaSize
	Put_PageSize(value PrintMediaSize)
	Get_PageOrientation() PrintOrientation
	Put_PageOrientation(value PrintOrientation)
	Get_PageSizeDimensions() Size
	IsPageSizeSupported(pageSize PrintMediaSize, pageOrientation PrintOrientation) bool
	Get_MaxNumberOfPages() uint32
	Put_MaxNumberOfPages(value uint32)
	Get_CanScanDuplex() bool
	Get_Duplex() bool
	Put_Duplex(value bool)
	Get_CanScanAhead() bool
	Get_ScanAhead() bool
	Put_ScanAhead(value bool)
}

type IImageScannerFeederConfigurationVtbl struct {
	win32.IInspectableVtbl
	Get_CanAutoDetectPageSize uintptr
	Get_AutoDetectPageSize    uintptr
	Put_AutoDetectPageSize    uintptr
	Get_PageSize              uintptr
	Put_PageSize              uintptr
	Get_PageOrientation       uintptr
	Put_PageOrientation       uintptr
	Get_PageSizeDimensions    uintptr
	IsPageSizeSupported       uintptr
	Get_MaxNumberOfPages      uintptr
	Put_MaxNumberOfPages      uintptr
	Get_CanScanDuplex         uintptr
	Get_Duplex                uintptr
	Put_Duplex                uintptr
	Get_CanScanAhead          uintptr
	Get_ScanAhead             uintptr
	Put_ScanAhead             uintptr
}

type IImageScannerFeederConfiguration struct {
	win32.IInspectable
}

func (this *IImageScannerFeederConfiguration) Vtbl() *IImageScannerFeederConfigurationVtbl {
	return (*IImageScannerFeederConfigurationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IImageScannerFeederConfiguration) Get_CanAutoDetectPageSize() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanAutoDetectPageSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerFeederConfiguration) Get_AutoDetectPageSize() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AutoDetectPageSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerFeederConfiguration) Put_AutoDetectPageSize(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AutoDetectPageSize, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IImageScannerFeederConfiguration) Get_PageSize() PrintMediaSize {
	var _result PrintMediaSize
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PageSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerFeederConfiguration) Put_PageSize(value PrintMediaSize) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PageSize, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IImageScannerFeederConfiguration) Get_PageOrientation() PrintOrientation {
	var _result PrintOrientation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PageOrientation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerFeederConfiguration) Put_PageOrientation(value PrintOrientation) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PageOrientation, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IImageScannerFeederConfiguration) Get_PageSizeDimensions() Size {
	var _result Size
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PageSizeDimensions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerFeederConfiguration) IsPageSizeSupported(pageSize PrintMediaSize, pageOrientation PrintOrientation) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsPageSizeSupported, uintptr(unsafe.Pointer(this)), uintptr(pageSize), uintptr(pageOrientation), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerFeederConfiguration) Get_MaxNumberOfPages() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxNumberOfPages, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerFeederConfiguration) Put_MaxNumberOfPages(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MaxNumberOfPages, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IImageScannerFeederConfiguration) Get_CanScanDuplex() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanScanDuplex, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerFeederConfiguration) Get_Duplex() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Duplex, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerFeederConfiguration) Put_Duplex(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Duplex, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IImageScannerFeederConfiguration) Get_CanScanAhead() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanScanAhead, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerFeederConfiguration) Get_ScanAhead() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScanAhead, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerFeederConfiguration) Put_ScanAhead(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ScanAhead, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// AE275D11-DADF-4010-BF10-CCA5C83DCBB0
var IID_IImageScannerFormatConfiguration = syscall.GUID{0xAE275D11, 0xDADF, 0x4010,
	[8]byte{0xBF, 0x10, 0xCC, 0xA5, 0xC8, 0x3D, 0xCB, 0xB0}}

type IImageScannerFormatConfigurationInterface interface {
	win32.IInspectableInterface
	Get_DefaultFormat() ImageScannerFormat
	Get_Format() ImageScannerFormat
	Put_Format(value ImageScannerFormat)
	IsFormatSupported(value ImageScannerFormat) bool
}

type IImageScannerFormatConfigurationVtbl struct {
	win32.IInspectableVtbl
	Get_DefaultFormat uintptr
	Get_Format        uintptr
	Put_Format        uintptr
	IsFormatSupported uintptr
}

type IImageScannerFormatConfiguration struct {
	win32.IInspectable
}

func (this *IImageScannerFormatConfiguration) Vtbl() *IImageScannerFormatConfigurationVtbl {
	return (*IImageScannerFormatConfigurationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IImageScannerFormatConfiguration) Get_DefaultFormat() ImageScannerFormat {
	var _result ImageScannerFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DefaultFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerFormatConfiguration) Get_Format() ImageScannerFormat {
	var _result ImageScannerFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Format, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerFormatConfiguration) Put_Format(value ImageScannerFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Format, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IImageScannerFormatConfiguration) IsFormatSupported(value ImageScannerFormat) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsFormatSupported, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 08B7FE8E-8891-441D-BE9C-176FA109C8BB
var IID_IImageScannerPreviewResult = syscall.GUID{0x08B7FE8E, 0x8891, 0x441D,
	[8]byte{0xBE, 0x9C, 0x17, 0x6F, 0xA1, 0x09, 0xC8, 0xBB}}

type IImageScannerPreviewResultInterface interface {
	win32.IInspectableInterface
	Get_Succeeded() bool
	Get_Format() ImageScannerFormat
}

type IImageScannerPreviewResultVtbl struct {
	win32.IInspectableVtbl
	Get_Succeeded uintptr
	Get_Format    uintptr
}

type IImageScannerPreviewResult struct {
	win32.IInspectable
}

func (this *IImageScannerPreviewResult) Vtbl() *IImageScannerPreviewResultVtbl {
	return (*IImageScannerPreviewResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IImageScannerPreviewResult) Get_Succeeded() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Succeeded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerPreviewResult) Get_Format() ImageScannerFormat {
	var _result ImageScannerFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Format, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C91624CD-9037-4E48-84C1-AC0975076BC5
var IID_IImageScannerScanResult = syscall.GUID{0xC91624CD, 0x9037, 0x4E48,
	[8]byte{0x84, 0xC1, 0xAC, 0x09, 0x75, 0x07, 0x6B, 0xC5}}

type IImageScannerScanResultInterface interface {
	win32.IInspectableInterface
	Get_ScannedFiles() *IVectorView[*IStorageFile]
}

type IImageScannerScanResultVtbl struct {
	win32.IInspectableVtbl
	Get_ScannedFiles uintptr
}

type IImageScannerScanResult struct {
	win32.IInspectable
}

func (this *IImageScannerScanResult) Vtbl() *IImageScannerScanResultVtbl {
	return (*IImageScannerScanResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IImageScannerScanResult) Get_ScannedFiles() *IVectorView[*IStorageFile] {
	var _result *IVectorView[*IStorageFile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScannedFiles, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BFB50055-0B44-4C82-9E89-205F9C234E59
var IID_IImageScannerSourceConfiguration = syscall.GUID{0xBFB50055, 0x0B44, 0x4C82,
	[8]byte{0x9E, 0x89, 0x20, 0x5F, 0x9C, 0x23, 0x4E, 0x59}}

type IImageScannerSourceConfigurationInterface interface {
	win32.IInspectableInterface
	Get_MinScanArea() Size
	Get_MaxScanArea() Size
	Get_SelectedScanRegion() Rect
	Put_SelectedScanRegion(value Rect)
	Get_AutoCroppingMode() ImageScannerAutoCroppingMode
	Put_AutoCroppingMode(value ImageScannerAutoCroppingMode)
	IsAutoCroppingModeSupported(value ImageScannerAutoCroppingMode) bool
	Get_MinResolution() ImageScannerResolution
	Get_MaxResolution() ImageScannerResolution
	Get_OpticalResolution() ImageScannerResolution
	Get_DesiredResolution() ImageScannerResolution
	Put_DesiredResolution(value ImageScannerResolution)
	Get_ActualResolution() ImageScannerResolution
	Get_DefaultColorMode() ImageScannerColorMode
	Get_ColorMode() ImageScannerColorMode
	Put_ColorMode(value ImageScannerColorMode)
	IsColorModeSupported(value ImageScannerColorMode) bool
	Get_MinBrightness() int32
	Get_MaxBrightness() int32
	Get_BrightnessStep() uint32
	Get_DefaultBrightness() int32
	Get_Brightness() int32
	Put_Brightness(value int32)
	Get_MinContrast() int32
	Get_MaxContrast() int32
	Get_ContrastStep() uint32
	Get_DefaultContrast() int32
	Get_Contrast() int32
	Put_Contrast(value int32)
}

type IImageScannerSourceConfigurationVtbl struct {
	win32.IInspectableVtbl
	Get_MinScanArea             uintptr
	Get_MaxScanArea             uintptr
	Get_SelectedScanRegion      uintptr
	Put_SelectedScanRegion      uintptr
	Get_AutoCroppingMode        uintptr
	Put_AutoCroppingMode        uintptr
	IsAutoCroppingModeSupported uintptr
	Get_MinResolution           uintptr
	Get_MaxResolution           uintptr
	Get_OpticalResolution       uintptr
	Get_DesiredResolution       uintptr
	Put_DesiredResolution       uintptr
	Get_ActualResolution        uintptr
	Get_DefaultColorMode        uintptr
	Get_ColorMode               uintptr
	Put_ColorMode               uintptr
	IsColorModeSupported        uintptr
	Get_MinBrightness           uintptr
	Get_MaxBrightness           uintptr
	Get_BrightnessStep          uintptr
	Get_DefaultBrightness       uintptr
	Get_Brightness              uintptr
	Put_Brightness              uintptr
	Get_MinContrast             uintptr
	Get_MaxContrast             uintptr
	Get_ContrastStep            uintptr
	Get_DefaultContrast         uintptr
	Get_Contrast                uintptr
	Put_Contrast                uintptr
}

type IImageScannerSourceConfiguration struct {
	win32.IInspectable
}

func (this *IImageScannerSourceConfiguration) Vtbl() *IImageScannerSourceConfigurationVtbl {
	return (*IImageScannerSourceConfigurationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IImageScannerSourceConfiguration) Get_MinScanArea() Size {
	var _result Size
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinScanArea, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerSourceConfiguration) Get_MaxScanArea() Size {
	var _result Size
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxScanArea, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerSourceConfiguration) Get_SelectedScanRegion() Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SelectedScanRegion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerSourceConfiguration) Put_SelectedScanRegion(value Rect) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SelectedScanRegion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IImageScannerSourceConfiguration) Get_AutoCroppingMode() ImageScannerAutoCroppingMode {
	var _result ImageScannerAutoCroppingMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AutoCroppingMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerSourceConfiguration) Put_AutoCroppingMode(value ImageScannerAutoCroppingMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AutoCroppingMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IImageScannerSourceConfiguration) IsAutoCroppingModeSupported(value ImageScannerAutoCroppingMode) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsAutoCroppingModeSupported, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerSourceConfiguration) Get_MinResolution() ImageScannerResolution {
	var _result ImageScannerResolution
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinResolution, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerSourceConfiguration) Get_MaxResolution() ImageScannerResolution {
	var _result ImageScannerResolution
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxResolution, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerSourceConfiguration) Get_OpticalResolution() ImageScannerResolution {
	var _result ImageScannerResolution
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OpticalResolution, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerSourceConfiguration) Get_DesiredResolution() ImageScannerResolution {
	var _result ImageScannerResolution
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesiredResolution, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerSourceConfiguration) Put_DesiredResolution(value ImageScannerResolution) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DesiredResolution, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IImageScannerSourceConfiguration) Get_ActualResolution() ImageScannerResolution {
	var _result ImageScannerResolution
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ActualResolution, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerSourceConfiguration) Get_DefaultColorMode() ImageScannerColorMode {
	var _result ImageScannerColorMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DefaultColorMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerSourceConfiguration) Get_ColorMode() ImageScannerColorMode {
	var _result ImageScannerColorMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ColorMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerSourceConfiguration) Put_ColorMode(value ImageScannerColorMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ColorMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IImageScannerSourceConfiguration) IsColorModeSupported(value ImageScannerColorMode) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsColorModeSupported, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerSourceConfiguration) Get_MinBrightness() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinBrightness, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerSourceConfiguration) Get_MaxBrightness() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxBrightness, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerSourceConfiguration) Get_BrightnessStep() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BrightnessStep, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerSourceConfiguration) Get_DefaultBrightness() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DefaultBrightness, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerSourceConfiguration) Get_Brightness() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Brightness, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerSourceConfiguration) Put_Brightness(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Brightness, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IImageScannerSourceConfiguration) Get_MinContrast() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinContrast, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerSourceConfiguration) Get_MaxContrast() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxContrast, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerSourceConfiguration) Get_ContrastStep() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContrastStep, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerSourceConfiguration) Get_DefaultContrast() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DefaultContrast, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerSourceConfiguration) Get_Contrast() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Contrast, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageScannerSourceConfiguration) Put_Contrast(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Contrast, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// BC57E70E-D804-4477-9FB5-B911B5473897
var IID_IImageScannerStatics = syscall.GUID{0xBC57E70E, 0xD804, 0x4477,
	[8]byte{0x9F, 0xB5, 0xB9, 0x11, 0xB5, 0x47, 0x38, 0x97}}

type IImageScannerStaticsInterface interface {
	win32.IInspectableInterface
	FromIdAsync(deviceId string) *IAsyncOperation[*IImageScanner]
	GetDeviceSelector() string
}

type IImageScannerStaticsVtbl struct {
	win32.IInspectableVtbl
	FromIdAsync       uintptr
	GetDeviceSelector uintptr
}

type IImageScannerStatics struct {
	win32.IInspectable
}

func (this *IImageScannerStatics) Vtbl() *IImageScannerStaticsVtbl {
	return (*IImageScannerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IImageScannerStatics) FromIdAsync(deviceId string) *IAsyncOperation[*IImageScanner] {
	var _result *IAsyncOperation[*IImageScanner]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IImageScannerStatics) GetDeviceSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// classes

type ImageScanner struct {
	RtClass
	*IImageScanner
}

func NewIImageScannerStatics() *IImageScannerStatics {
	var p *IImageScannerStatics
	hs := NewHStr("Windows.Devices.Scanners.ImageScanner")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IImageScannerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type ImageScannerAutoConfiguration struct {
	RtClass
	*IImageScannerFormatConfiguration
}

type ImageScannerFeederConfiguration struct {
	RtClass
	*IImageScannerFormatConfiguration
}

type ImageScannerFlatbedConfiguration struct {
	RtClass
	*IImageScannerFormatConfiguration
}

type ImageScannerPreviewResult struct {
	RtClass
	*IImageScannerPreviewResult
}

type ImageScannerScanResult struct {
	RtClass
	*IImageScannerScanResult
}
