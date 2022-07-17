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
type AdvancedPhotoMode int32

const (
	AdvancedPhotoMode_Auto     AdvancedPhotoMode = 0
	AdvancedPhotoMode_Standard AdvancedPhotoMode = 1
	AdvancedPhotoMode_Hdr      AdvancedPhotoMode = 2
	AdvancedPhotoMode_LowLight AdvancedPhotoMode = 3
)

// enum
type AudioDeviceRole int32

const (
	AudioDeviceRole_Default        AudioDeviceRole = 0
	AudioDeviceRole_Communications AudioDeviceRole = 1
)

// enum
type AutoFocusRange int32

const (
	AutoFocusRange_FullRange AutoFocusRange = 0
	AutoFocusRange_Macro     AutoFocusRange = 1
	AutoFocusRange_Normal    AutoFocusRange = 2
)

// enum
type CameraOcclusionKind int32

const (
	CameraOcclusionKind_Lid            CameraOcclusionKind = 0
	CameraOcclusionKind_CameraHardware CameraOcclusionKind = 1
)

// enum
type CameraStreamState int32

const (
	CameraStreamState_NotStreaming      CameraStreamState = 0
	CameraStreamState_Streaming         CameraStreamState = 1
	CameraStreamState_BlockedForPrivacy CameraStreamState = 2
	CameraStreamState_Shutdown          CameraStreamState = 3
)

// enum
type CaptureSceneMode int32

const (
	CaptureSceneMode_Auto          CaptureSceneMode = 0
	CaptureSceneMode_Manual        CaptureSceneMode = 1
	CaptureSceneMode_Macro         CaptureSceneMode = 2
	CaptureSceneMode_Portrait      CaptureSceneMode = 3
	CaptureSceneMode_Sport         CaptureSceneMode = 4
	CaptureSceneMode_Snow          CaptureSceneMode = 5
	CaptureSceneMode_Night         CaptureSceneMode = 6
	CaptureSceneMode_Beach         CaptureSceneMode = 7
	CaptureSceneMode_Sunset        CaptureSceneMode = 8
	CaptureSceneMode_Candlelight   CaptureSceneMode = 9
	CaptureSceneMode_Landscape     CaptureSceneMode = 10
	CaptureSceneMode_NightPortrait CaptureSceneMode = 11
	CaptureSceneMode_Backlit       CaptureSceneMode = 12
)

// enum
type CaptureUse int32

const (
	CaptureUse_None  CaptureUse = 0
	CaptureUse_Photo CaptureUse = 1
	CaptureUse_Video CaptureUse = 2
)

// enum
type ColorTemperaturePreset int32

const (
	ColorTemperaturePreset_Auto        ColorTemperaturePreset = 0
	ColorTemperaturePreset_Manual      ColorTemperaturePreset = 1
	ColorTemperaturePreset_Cloudy      ColorTemperaturePreset = 2
	ColorTemperaturePreset_Daylight    ColorTemperaturePreset = 3
	ColorTemperaturePreset_Flash       ColorTemperaturePreset = 4
	ColorTemperaturePreset_Fluorescent ColorTemperaturePreset = 5
	ColorTemperaturePreset_Tungsten    ColorTemperaturePreset = 6
	ColorTemperaturePreset_Candlelight ColorTemperaturePreset = 7
)

// enum
type DigitalWindowMode int32

const (
	DigitalWindowMode_Off  DigitalWindowMode = 0
	DigitalWindowMode_On   DigitalWindowMode = 1
	DigitalWindowMode_Auto DigitalWindowMode = 2
)

// enum
type FocusMode int32

const (
	FocusMode_Auto       FocusMode = 0
	FocusMode_Single     FocusMode = 1
	FocusMode_Continuous FocusMode = 2
	FocusMode_Manual     FocusMode = 3
)

// enum
type FocusPreset int32

const (
	FocusPreset_Auto           FocusPreset = 0
	FocusPreset_Manual         FocusPreset = 1
	FocusPreset_AutoMacro      FocusPreset = 2
	FocusPreset_AutoNormal     FocusPreset = 3
	FocusPreset_AutoInfinity   FocusPreset = 4
	FocusPreset_AutoHyperfocal FocusPreset = 5
)

// enum
type HdrVideoMode int32

const (
	HdrVideoMode_Off  HdrVideoMode = 0
	HdrVideoMode_On   HdrVideoMode = 1
	HdrVideoMode_Auto HdrVideoMode = 2
)

// enum
type InfraredTorchMode int32

const (
	InfraredTorchMode_Off                          InfraredTorchMode = 0
	InfraredTorchMode_On                           InfraredTorchMode = 1
	InfraredTorchMode_AlternatingFrameIllumination InfraredTorchMode = 2
)

// enum
type IsoSpeedPreset int32

const (
	IsoSpeedPreset_Auto     IsoSpeedPreset = 0
	IsoSpeedPreset_Iso50    IsoSpeedPreset = 1
	IsoSpeedPreset_Iso80    IsoSpeedPreset = 2
	IsoSpeedPreset_Iso100   IsoSpeedPreset = 3
	IsoSpeedPreset_Iso200   IsoSpeedPreset = 4
	IsoSpeedPreset_Iso400   IsoSpeedPreset = 5
	IsoSpeedPreset_Iso800   IsoSpeedPreset = 6
	IsoSpeedPreset_Iso1600  IsoSpeedPreset = 7
	IsoSpeedPreset_Iso3200  IsoSpeedPreset = 8
	IsoSpeedPreset_Iso6400  IsoSpeedPreset = 9
	IsoSpeedPreset_Iso12800 IsoSpeedPreset = 10
	IsoSpeedPreset_Iso25600 IsoSpeedPreset = 11
)

// enum
type ManualFocusDistance int32

const (
	ManualFocusDistance_Infinity   ManualFocusDistance = 0
	ManualFocusDistance_Hyperfocal ManualFocusDistance = 1
	ManualFocusDistance_Nearest    ManualFocusDistance = 2
)

// enum
type MediaCaptureFocusState int32

const (
	MediaCaptureFocusState_Uninitialized MediaCaptureFocusState = 0
	MediaCaptureFocusState_Lost          MediaCaptureFocusState = 1
	MediaCaptureFocusState_Searching     MediaCaptureFocusState = 2
	MediaCaptureFocusState_Focused       MediaCaptureFocusState = 3
	MediaCaptureFocusState_Failed        MediaCaptureFocusState = 4
)

// enum
type MediaCaptureOptimization int32

const (
	MediaCaptureOptimization_Default            MediaCaptureOptimization = 0
	MediaCaptureOptimization_Quality            MediaCaptureOptimization = 1
	MediaCaptureOptimization_Latency            MediaCaptureOptimization = 2
	MediaCaptureOptimization_Power              MediaCaptureOptimization = 3
	MediaCaptureOptimization_LatencyThenQuality MediaCaptureOptimization = 4
	MediaCaptureOptimization_LatencyThenPower   MediaCaptureOptimization = 5
	MediaCaptureOptimization_PowerAndQuality    MediaCaptureOptimization = 6
)

// enum
type MediaCapturePauseBehavior int32

const (
	MediaCapturePauseBehavior_RetainHardwareResources  MediaCapturePauseBehavior = 0
	MediaCapturePauseBehavior_ReleaseHardwareResources MediaCapturePauseBehavior = 1
)

// enum
type OpticalImageStabilizationMode int32

const (
	OpticalImageStabilizationMode_Off  OpticalImageStabilizationMode = 0
	OpticalImageStabilizationMode_On   OpticalImageStabilizationMode = 1
	OpticalImageStabilizationMode_Auto OpticalImageStabilizationMode = 2
)

// enum
type RegionOfInterestType int32

const (
	RegionOfInterestType_Unknown RegionOfInterestType = 0
	RegionOfInterestType_Face    RegionOfInterestType = 1
)

// enum
type SendCommandStatus int32

const (
	SendCommandStatus_Success            SendCommandStatus = 0
	SendCommandStatus_DeviceNotAvailable SendCommandStatus = 1
)

// enum
type TelephonyKey int32

const (
	TelephonyKey_D0    TelephonyKey = 0
	TelephonyKey_D1    TelephonyKey = 1
	TelephonyKey_D2    TelephonyKey = 2
	TelephonyKey_D3    TelephonyKey = 3
	TelephonyKey_D4    TelephonyKey = 4
	TelephonyKey_D5    TelephonyKey = 5
	TelephonyKey_D6    TelephonyKey = 6
	TelephonyKey_D7    TelephonyKey = 7
	TelephonyKey_D8    TelephonyKey = 8
	TelephonyKey_D9    TelephonyKey = 9
	TelephonyKey_Star  TelephonyKey = 10
	TelephonyKey_Pound TelephonyKey = 11
	TelephonyKey_A     TelephonyKey = 12
	TelephonyKey_B     TelephonyKey = 13
	TelephonyKey_C     TelephonyKey = 14
	TelephonyKey_D     TelephonyKey = 15
)

// enum
type VideoDeviceControllerGetDevicePropertyStatus int32

const (
	VideoDeviceControllerGetDevicePropertyStatus_Success                      VideoDeviceControllerGetDevicePropertyStatus = 0
	VideoDeviceControllerGetDevicePropertyStatus_UnknownFailure               VideoDeviceControllerGetDevicePropertyStatus = 1
	VideoDeviceControllerGetDevicePropertyStatus_BufferTooSmall               VideoDeviceControllerGetDevicePropertyStatus = 2
	VideoDeviceControllerGetDevicePropertyStatus_NotSupported                 VideoDeviceControllerGetDevicePropertyStatus = 3
	VideoDeviceControllerGetDevicePropertyStatus_DeviceNotAvailable           VideoDeviceControllerGetDevicePropertyStatus = 4
	VideoDeviceControllerGetDevicePropertyStatus_MaxPropertyValueSizeTooSmall VideoDeviceControllerGetDevicePropertyStatus = 5
	VideoDeviceControllerGetDevicePropertyStatus_MaxPropertyValueSizeRequired VideoDeviceControllerGetDevicePropertyStatus = 6
)

// enum
type VideoDeviceControllerSetDevicePropertyStatus int32

const (
	VideoDeviceControllerSetDevicePropertyStatus_Success            VideoDeviceControllerSetDevicePropertyStatus = 0
	VideoDeviceControllerSetDevicePropertyStatus_UnknownFailure     VideoDeviceControllerSetDevicePropertyStatus = 1
	VideoDeviceControllerSetDevicePropertyStatus_NotSupported       VideoDeviceControllerSetDevicePropertyStatus = 2
	VideoDeviceControllerSetDevicePropertyStatus_InvalidValue       VideoDeviceControllerSetDevicePropertyStatus = 3
	VideoDeviceControllerSetDevicePropertyStatus_DeviceNotAvailable VideoDeviceControllerSetDevicePropertyStatus = 4
	VideoDeviceControllerSetDevicePropertyStatus_NotInControl       VideoDeviceControllerSetDevicePropertyStatus = 5
)

// enum
type VideoTemporalDenoisingMode int32

const (
	VideoTemporalDenoisingMode_Off  VideoTemporalDenoisingMode = 0
	VideoTemporalDenoisingMode_On   VideoTemporalDenoisingMode = 1
	VideoTemporalDenoisingMode_Auto VideoTemporalDenoisingMode = 2
)

// enum
type ZoomTransitionMode int32

const (
	ZoomTransitionMode_Auto   ZoomTransitionMode = 0
	ZoomTransitionMode_Direct ZoomTransitionMode = 1
	ZoomTransitionMode_Smooth ZoomTransitionMode = 2
)

// structs

type CallControlContract struct {
}

// func types

//596F759F-50DF-4454-BC63-4D3D01B61958
type CallControlEventHandler func(sender *ICallControl) com.Error

//5ABBFFDB-C21F-4BC4-891B-257E28C1B1A4
type DialRequestedEventHandler func(sender *ICallControl, e *IDialRequestedEventArgs) com.Error

//E637A454-C527-422C-8926-C9AF83B559A0
type KeypadPressedEventHandler func(sender *ICallControl, e *IKeypadPressedEventArgs) com.Error

//BAF257D1-4EBD-4B84-9F47-6EC43D75D8B1
type RedialRequestedEventHandler func(sender *ICallControl, e *IRedialRequestedEventArgs) com.Error

// interfaces

// 08F3863A-0018-445B-93D2-646D1C5ED05C
var IID_IAdvancedPhotoCaptureSettings = syscall.GUID{0x08F3863A, 0x0018, 0x445B,
	[8]byte{0x93, 0xD2, 0x64, 0x6D, 0x1C, 0x5E, 0xD0, 0x5C}}

type IAdvancedPhotoCaptureSettingsInterface interface {
	win32.IInspectableInterface
	Get_Mode() AdvancedPhotoMode
	Put_Mode(value AdvancedPhotoMode)
}

type IAdvancedPhotoCaptureSettingsVtbl struct {
	win32.IInspectableVtbl
	Get_Mode uintptr
	Put_Mode uintptr
}

type IAdvancedPhotoCaptureSettings struct {
	win32.IInspectable
}

func (this *IAdvancedPhotoCaptureSettings) Vtbl() *IAdvancedPhotoCaptureSettingsVtbl {
	return (*IAdvancedPhotoCaptureSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdvancedPhotoCaptureSettings) Get_Mode() AdvancedPhotoMode {
	var _result AdvancedPhotoMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdvancedPhotoCaptureSettings) Put_Mode(value AdvancedPhotoMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Mode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// C5B15486-9001-4682-9309-68EAE0080EEC
var IID_IAdvancedPhotoControl = syscall.GUID{0xC5B15486, 0x9001, 0x4682,
	[8]byte{0x93, 0x09, 0x68, 0xEA, 0xE0, 0x08, 0x0E, 0xEC}}

type IAdvancedPhotoControlInterface interface {
	win32.IInspectableInterface
	Get_Supported() bool
	Get_SupportedModes() *IVectorView[AdvancedPhotoMode]
	Get_Mode() AdvancedPhotoMode
	Configure(settings *IAdvancedPhotoCaptureSettings)
}

type IAdvancedPhotoControlVtbl struct {
	win32.IInspectableVtbl
	Get_Supported      uintptr
	Get_SupportedModes uintptr
	Get_Mode           uintptr
	Configure          uintptr
}

type IAdvancedPhotoControl struct {
	win32.IInspectable
}

func (this *IAdvancedPhotoControl) Vtbl() *IAdvancedPhotoControlVtbl {
	return (*IAdvancedPhotoControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdvancedPhotoControl) Get_Supported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Supported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdvancedPhotoControl) Get_SupportedModes() *IVectorView[AdvancedPhotoMode] {
	var _result *IVectorView[AdvancedPhotoMode]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedModes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdvancedPhotoControl) Get_Mode() AdvancedPhotoMode {
	var _result AdvancedPhotoMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdvancedPhotoControl) Configure(settings *IAdvancedPhotoCaptureSettings) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Configure, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(settings)))
	_ = _hr
}

// DE6FF4D3-2B96-4583-80AB-B5B01DC6A8D7
var IID_IAdvancedVideoCaptureDeviceController = syscall.GUID{0xDE6FF4D3, 0x2B96, 0x4583,
	[8]byte{0x80, 0xAB, 0xB5, 0xB0, 0x1D, 0xC6, 0xA8, 0xD7}}

type IAdvancedVideoCaptureDeviceControllerInterface interface {
	win32.IInspectableInterface
	SetDeviceProperty(propertyId string, propertyValue interface{})
	GetDeviceProperty(propertyId string) interface{}
}

type IAdvancedVideoCaptureDeviceControllerVtbl struct {
	win32.IInspectableVtbl
	SetDeviceProperty uintptr
	GetDeviceProperty uintptr
}

type IAdvancedVideoCaptureDeviceController struct {
	win32.IInspectable
}

func (this *IAdvancedVideoCaptureDeviceController) Vtbl() *IAdvancedVideoCaptureDeviceControllerVtbl {
	return (*IAdvancedVideoCaptureDeviceControllerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdvancedVideoCaptureDeviceController) SetDeviceProperty(propertyId string, propertyValue interface{}) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetDeviceProperty, uintptr(unsafe.Pointer(this)), NewHStr(propertyId).Ptr, *(*uintptr)(unsafe.Pointer(&propertyValue)))
	_ = _hr
}

func (this *IAdvancedVideoCaptureDeviceController) GetDeviceProperty(propertyId string) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceProperty, uintptr(unsafe.Pointer(this)), NewHStr(propertyId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C621B82D-D6F0-5C1B-A388-A6E938407146
var IID_IAdvancedVideoCaptureDeviceController10 = syscall.GUID{0xC621B82D, 0xD6F0, 0x5C1B,
	[8]byte{0xA3, 0x88, 0xA6, 0xE9, 0x38, 0x40, 0x71, 0x46}}

type IAdvancedVideoCaptureDeviceController10Interface interface {
	win32.IInspectableInterface
	Get_CameraOcclusionInfo() *ICameraOcclusionInfo
}

type IAdvancedVideoCaptureDeviceController10Vtbl struct {
	win32.IInspectableVtbl
	Get_CameraOcclusionInfo uintptr
}

type IAdvancedVideoCaptureDeviceController10 struct {
	win32.IInspectable
}

func (this *IAdvancedVideoCaptureDeviceController10) Vtbl() *IAdvancedVideoCaptureDeviceController10Vtbl {
	return (*IAdvancedVideoCaptureDeviceController10Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdvancedVideoCaptureDeviceController10) Get_CameraOcclusionInfo() *ICameraOcclusionInfo {
	var _result *ICameraOcclusionInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CameraOcclusionInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D5B65AE2-3772-580C-A630-E75DE9106904
var IID_IAdvancedVideoCaptureDeviceController11 = syscall.GUID{0xD5B65AE2, 0x3772, 0x580C,
	[8]byte{0xA6, 0x30, 0xE7, 0x5D, 0xE9, 0x10, 0x69, 0x04}}

type IAdvancedVideoCaptureDeviceController11Interface interface {
	win32.IInspectableInterface
	TryAcquireExclusiveControl(deviceId string, mode MediaCaptureDeviceExclusiveControlReleaseMode) bool
}

type IAdvancedVideoCaptureDeviceController11Vtbl struct {
	win32.IInspectableVtbl
	TryAcquireExclusiveControl uintptr
}

type IAdvancedVideoCaptureDeviceController11 struct {
	win32.IInspectable
}

func (this *IAdvancedVideoCaptureDeviceController11) Vtbl() *IAdvancedVideoCaptureDeviceController11Vtbl {
	return (*IAdvancedVideoCaptureDeviceController11Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdvancedVideoCaptureDeviceController11) TryAcquireExclusiveControl(deviceId string, mode MediaCaptureDeviceExclusiveControlReleaseMode) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryAcquireExclusiveControl, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(mode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 8BB94F8F-F11A-43DB-B402-11930B80AE56
var IID_IAdvancedVideoCaptureDeviceController2 = syscall.GUID{0x8BB94F8F, 0xF11A, 0x43DB,
	[8]byte{0xB4, 0x02, 0x11, 0x93, 0x0B, 0x80, 0xAE, 0x56}}

type IAdvancedVideoCaptureDeviceController2Interface interface {
	win32.IInspectableInterface
	Get_LowLagPhotoSequence() *ILowLagPhotoSequenceControl
	Get_LowLagPhoto() *ILowLagPhotoControl
	Get_SceneModeControl() *ISceneModeControl
	Get_TorchControl() *ITorchControl
	Get_FlashControl() *IFlashControl
	Get_WhiteBalanceControl() *IWhiteBalanceControl
	Get_ExposureControl() *IExposureControl
	Get_FocusControl() *IFocusControl
	Get_ExposureCompensationControl() *IExposureCompensationControl
	Get_IsoSpeedControl() *IIsoSpeedControl
	Get_RegionsOfInterestControl() *IRegionsOfInterestControl
	Get_PrimaryUse() CaptureUse
	Put_PrimaryUse(value CaptureUse)
}

type IAdvancedVideoCaptureDeviceController2Vtbl struct {
	win32.IInspectableVtbl
	Get_LowLagPhotoSequence         uintptr
	Get_LowLagPhoto                 uintptr
	Get_SceneModeControl            uintptr
	Get_TorchControl                uintptr
	Get_FlashControl                uintptr
	Get_WhiteBalanceControl         uintptr
	Get_ExposureControl             uintptr
	Get_FocusControl                uintptr
	Get_ExposureCompensationControl uintptr
	Get_IsoSpeedControl             uintptr
	Get_RegionsOfInterestControl    uintptr
	Get_PrimaryUse                  uintptr
	Put_PrimaryUse                  uintptr
}

type IAdvancedVideoCaptureDeviceController2 struct {
	win32.IInspectable
}

func (this *IAdvancedVideoCaptureDeviceController2) Vtbl() *IAdvancedVideoCaptureDeviceController2Vtbl {
	return (*IAdvancedVideoCaptureDeviceController2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdvancedVideoCaptureDeviceController2) Get_LowLagPhotoSequence() *ILowLagPhotoSequenceControl {
	var _result *ILowLagPhotoSequenceControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LowLagPhotoSequence, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdvancedVideoCaptureDeviceController2) Get_LowLagPhoto() *ILowLagPhotoControl {
	var _result *ILowLagPhotoControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LowLagPhoto, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdvancedVideoCaptureDeviceController2) Get_SceneModeControl() *ISceneModeControl {
	var _result *ISceneModeControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SceneModeControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdvancedVideoCaptureDeviceController2) Get_TorchControl() *ITorchControl {
	var _result *ITorchControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TorchControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdvancedVideoCaptureDeviceController2) Get_FlashControl() *IFlashControl {
	var _result *IFlashControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FlashControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdvancedVideoCaptureDeviceController2) Get_WhiteBalanceControl() *IWhiteBalanceControl {
	var _result *IWhiteBalanceControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WhiteBalanceControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdvancedVideoCaptureDeviceController2) Get_ExposureControl() *IExposureControl {
	var _result *IExposureControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExposureControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdvancedVideoCaptureDeviceController2) Get_FocusControl() *IFocusControl {
	var _result *IFocusControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FocusControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdvancedVideoCaptureDeviceController2) Get_ExposureCompensationControl() *IExposureCompensationControl {
	var _result *IExposureCompensationControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExposureCompensationControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdvancedVideoCaptureDeviceController2) Get_IsoSpeedControl() *IIsoSpeedControl {
	var _result *IIsoSpeedControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsoSpeedControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdvancedVideoCaptureDeviceController2) Get_RegionsOfInterestControl() *IRegionsOfInterestControl {
	var _result *IRegionsOfInterestControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RegionsOfInterestControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdvancedVideoCaptureDeviceController2) Get_PrimaryUse() CaptureUse {
	var _result CaptureUse
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PrimaryUse, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdvancedVideoCaptureDeviceController2) Put_PrimaryUse(value CaptureUse) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PrimaryUse, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// A98B8F34-EE0D-470C-B9F0-4229C4BBD089
var IID_IAdvancedVideoCaptureDeviceController3 = syscall.GUID{0xA98B8F34, 0xEE0D, 0x470C,
	[8]byte{0xB9, 0xF0, 0x42, 0x29, 0xC4, 0xBB, 0xD0, 0x89}}

type IAdvancedVideoCaptureDeviceController3Interface interface {
	win32.IInspectableInterface
	Get_VariablePhotoSequenceController() *IVariablePhotoSequenceController
	Get_PhotoConfirmationControl() *IPhotoConfirmationControl
	Get_ZoomControl() *IZoomControl
}

type IAdvancedVideoCaptureDeviceController3Vtbl struct {
	win32.IInspectableVtbl
	Get_VariablePhotoSequenceController uintptr
	Get_PhotoConfirmationControl        uintptr
	Get_ZoomControl                     uintptr
}

type IAdvancedVideoCaptureDeviceController3 struct {
	win32.IInspectable
}

func (this *IAdvancedVideoCaptureDeviceController3) Vtbl() *IAdvancedVideoCaptureDeviceController3Vtbl {
	return (*IAdvancedVideoCaptureDeviceController3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdvancedVideoCaptureDeviceController3) Get_VariablePhotoSequenceController() *IVariablePhotoSequenceController {
	var _result *IVariablePhotoSequenceController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VariablePhotoSequenceController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdvancedVideoCaptureDeviceController3) Get_PhotoConfirmationControl() *IPhotoConfirmationControl {
	var _result *IPhotoConfirmationControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhotoConfirmationControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdvancedVideoCaptureDeviceController3) Get_ZoomControl() *IZoomControl {
	var _result *IZoomControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ZoomControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EA9FBFAF-D371-41C3-9A17-824A87EBDFD2
var IID_IAdvancedVideoCaptureDeviceController4 = syscall.GUID{0xEA9FBFAF, 0xD371, 0x41C3,
	[8]byte{0x9A, 0x17, 0x82, 0x4A, 0x87, 0xEB, 0xDF, 0xD2}}

type IAdvancedVideoCaptureDeviceController4Interface interface {
	win32.IInspectableInterface
	Get_ExposurePriorityVideoControl() *IExposurePriorityVideoControl
	Get_DesiredOptimization() MediaCaptureOptimization
	Put_DesiredOptimization(value MediaCaptureOptimization)
	Get_HdrVideoControl() *IHdrVideoControl
	Get_OpticalImageStabilizationControl() *IOpticalImageStabilizationControl
	Get_AdvancedPhotoControl() *IAdvancedPhotoControl
}

type IAdvancedVideoCaptureDeviceController4Vtbl struct {
	win32.IInspectableVtbl
	Get_ExposurePriorityVideoControl     uintptr
	Get_DesiredOptimization              uintptr
	Put_DesiredOptimization              uintptr
	Get_HdrVideoControl                  uintptr
	Get_OpticalImageStabilizationControl uintptr
	Get_AdvancedPhotoControl             uintptr
}

type IAdvancedVideoCaptureDeviceController4 struct {
	win32.IInspectable
}

func (this *IAdvancedVideoCaptureDeviceController4) Vtbl() *IAdvancedVideoCaptureDeviceController4Vtbl {
	return (*IAdvancedVideoCaptureDeviceController4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdvancedVideoCaptureDeviceController4) Get_ExposurePriorityVideoControl() *IExposurePriorityVideoControl {
	var _result *IExposurePriorityVideoControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExposurePriorityVideoControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdvancedVideoCaptureDeviceController4) Get_DesiredOptimization() MediaCaptureOptimization {
	var _result MediaCaptureOptimization
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesiredOptimization, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdvancedVideoCaptureDeviceController4) Put_DesiredOptimization(value MediaCaptureOptimization) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DesiredOptimization, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAdvancedVideoCaptureDeviceController4) Get_HdrVideoControl() *IHdrVideoControl {
	var _result *IHdrVideoControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HdrVideoControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdvancedVideoCaptureDeviceController4) Get_OpticalImageStabilizationControl() *IOpticalImageStabilizationControl {
	var _result *IOpticalImageStabilizationControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OpticalImageStabilizationControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdvancedVideoCaptureDeviceController4) Get_AdvancedPhotoControl() *IAdvancedPhotoControl {
	var _result *IAdvancedPhotoControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AdvancedPhotoControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 33512B17-B9CB-4A23-B875-F9EAAB535492
var IID_IAdvancedVideoCaptureDeviceController5 = syscall.GUID{0x33512B17, 0xB9CB, 0x4A23,
	[8]byte{0xB8, 0x75, 0xF9, 0xEA, 0xAB, 0x53, 0x54, 0x92}}

type IAdvancedVideoCaptureDeviceController5Interface interface {
	win32.IInspectableInterface
	Get_Id() string
	GetDevicePropertyById(propertyId string, maxPropertyValueSize *IReference[uint32]) *IVideoDeviceControllerGetDevicePropertyResult
	SetDevicePropertyById(propertyId string, propertyValue interface{}) VideoDeviceControllerSetDevicePropertyStatus
	GetDevicePropertyByExtendedId(extendedPropertyIdLength uint32, extendedPropertyId *byte, maxPropertyValueSize *IReference[uint32]) *IVideoDeviceControllerGetDevicePropertyResult
	SetDevicePropertyByExtendedId(extendedPropertyIdLength uint32, extendedPropertyId *byte, propertyValueLength uint32, propertyValue *byte) VideoDeviceControllerSetDevicePropertyStatus
}

type IAdvancedVideoCaptureDeviceController5Vtbl struct {
	win32.IInspectableVtbl
	Get_Id                        uintptr
	GetDevicePropertyById         uintptr
	SetDevicePropertyById         uintptr
	GetDevicePropertyByExtendedId uintptr
	SetDevicePropertyByExtendedId uintptr
}

type IAdvancedVideoCaptureDeviceController5 struct {
	win32.IInspectable
}

func (this *IAdvancedVideoCaptureDeviceController5) Vtbl() *IAdvancedVideoCaptureDeviceController5Vtbl {
	return (*IAdvancedVideoCaptureDeviceController5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdvancedVideoCaptureDeviceController5) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAdvancedVideoCaptureDeviceController5) GetDevicePropertyById(propertyId string, maxPropertyValueSize *IReference[uint32]) *IVideoDeviceControllerGetDevicePropertyResult {
	var _result *IVideoDeviceControllerGetDevicePropertyResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDevicePropertyById, uintptr(unsafe.Pointer(this)), NewHStr(propertyId).Ptr, uintptr(unsafe.Pointer(maxPropertyValueSize)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdvancedVideoCaptureDeviceController5) SetDevicePropertyById(propertyId string, propertyValue interface{}) VideoDeviceControllerSetDevicePropertyStatus {
	var _result VideoDeviceControllerSetDevicePropertyStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetDevicePropertyById, uintptr(unsafe.Pointer(this)), NewHStr(propertyId).Ptr, *(*uintptr)(unsafe.Pointer(&propertyValue)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAdvancedVideoCaptureDeviceController5) GetDevicePropertyByExtendedId(extendedPropertyIdLength uint32, extendedPropertyId *byte, maxPropertyValueSize *IReference[uint32]) *IVideoDeviceControllerGetDevicePropertyResult {
	var _result *IVideoDeviceControllerGetDevicePropertyResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDevicePropertyByExtendedId, uintptr(unsafe.Pointer(this)), uintptr(extendedPropertyIdLength), uintptr(unsafe.Pointer(extendedPropertyId)), uintptr(unsafe.Pointer(maxPropertyValueSize)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAdvancedVideoCaptureDeviceController5) SetDevicePropertyByExtendedId(extendedPropertyIdLength uint32, extendedPropertyId *byte, propertyValueLength uint32, propertyValue *byte) VideoDeviceControllerSetDevicePropertyStatus {
	var _result VideoDeviceControllerSetDevicePropertyStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetDevicePropertyByExtendedId, uintptr(unsafe.Pointer(this)), uintptr(extendedPropertyIdLength), uintptr(unsafe.Pointer(extendedPropertyId)), uintptr(propertyValueLength), uintptr(unsafe.Pointer(propertyValue)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// B6563A53-68A1-44B7-9F89-B5FA97AC0CBE
var IID_IAdvancedVideoCaptureDeviceController6 = syscall.GUID{0xB6563A53, 0x68A1, 0x44B7,
	[8]byte{0x9F, 0x89, 0xB5, 0xFA, 0x97, 0xAC, 0x0C, 0xBE}}

type IAdvancedVideoCaptureDeviceController6Interface interface {
	win32.IInspectableInterface
	Get_VideoTemporalDenoisingControl() *IVideoTemporalDenoisingControl
}

type IAdvancedVideoCaptureDeviceController6Vtbl struct {
	win32.IInspectableVtbl
	Get_VideoTemporalDenoisingControl uintptr
}

type IAdvancedVideoCaptureDeviceController6 struct {
	win32.IInspectable
}

func (this *IAdvancedVideoCaptureDeviceController6) Vtbl() *IAdvancedVideoCaptureDeviceController6Vtbl {
	return (*IAdvancedVideoCaptureDeviceController6Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdvancedVideoCaptureDeviceController6) Get_VideoTemporalDenoisingControl() *IVideoTemporalDenoisingControl {
	var _result *IVideoTemporalDenoisingControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoTemporalDenoisingControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8D2927F0-A054-50E7-B7DF-7C04234D10F0
var IID_IAdvancedVideoCaptureDeviceController7 = syscall.GUID{0x8D2927F0, 0xA054, 0x50E7,
	[8]byte{0xB7, 0xDF, 0x7C, 0x04, 0x23, 0x4D, 0x10, 0xF0}}

type IAdvancedVideoCaptureDeviceController7Interface interface {
	win32.IInspectableInterface
	Get_InfraredTorchControl() *IInfraredTorchControl
}

type IAdvancedVideoCaptureDeviceController7Vtbl struct {
	win32.IInspectableVtbl
	Get_InfraredTorchControl uintptr
}

type IAdvancedVideoCaptureDeviceController7 struct {
	win32.IInspectable
}

func (this *IAdvancedVideoCaptureDeviceController7) Vtbl() *IAdvancedVideoCaptureDeviceController7Vtbl {
	return (*IAdvancedVideoCaptureDeviceController7Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdvancedVideoCaptureDeviceController7) Get_InfraredTorchControl() *IInfraredTorchControl {
	var _result *IInfraredTorchControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InfraredTorchControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D843F010-E7FB-595B-9A78-0E54C4532B43
var IID_IAdvancedVideoCaptureDeviceController8 = syscall.GUID{0xD843F010, 0xE7FB, 0x595B,
	[8]byte{0x9A, 0x78, 0x0E, 0x54, 0xC4, 0x53, 0x2B, 0x43}}

type IAdvancedVideoCaptureDeviceController8Interface interface {
	win32.IInspectableInterface
	Get_PanelBasedOptimizationControl() *IPanelBasedOptimizationControl
}

type IAdvancedVideoCaptureDeviceController8Vtbl struct {
	win32.IInspectableVtbl
	Get_PanelBasedOptimizationControl uintptr
}

type IAdvancedVideoCaptureDeviceController8 struct {
	win32.IInspectable
}

func (this *IAdvancedVideoCaptureDeviceController8) Vtbl() *IAdvancedVideoCaptureDeviceController8Vtbl {
	return (*IAdvancedVideoCaptureDeviceController8Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdvancedVideoCaptureDeviceController8) Get_PanelBasedOptimizationControl() *IPanelBasedOptimizationControl {
	var _result *IPanelBasedOptimizationControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PanelBasedOptimizationControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8BDCA95D-0255-51BC-A10D-5A169EC1625A
var IID_IAdvancedVideoCaptureDeviceController9 = syscall.GUID{0x8BDCA95D, 0x0255, 0x51BC,
	[8]byte{0xA1, 0x0D, 0x5A, 0x16, 0x9E, 0xC1, 0x62, 0x5A}}

type IAdvancedVideoCaptureDeviceController9Interface interface {
	win32.IInspectableInterface
	Get_DigitalWindowControl() *IDigitalWindowControl
}

type IAdvancedVideoCaptureDeviceController9Vtbl struct {
	win32.IInspectableVtbl
	Get_DigitalWindowControl uintptr
}

type IAdvancedVideoCaptureDeviceController9 struct {
	win32.IInspectable
}

func (this *IAdvancedVideoCaptureDeviceController9) Vtbl() *IAdvancedVideoCaptureDeviceController9Vtbl {
	return (*IAdvancedVideoCaptureDeviceController9Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAdvancedVideoCaptureDeviceController9) Get_DigitalWindowControl() *IDigitalWindowControl {
	var _result *IDigitalWindowControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DigitalWindowControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EDD4A388-79C7-4F7C-90E8-EF934B21580A
var IID_IAudioDeviceController = syscall.GUID{0xEDD4A388, 0x79C7, 0x4F7C,
	[8]byte{0x90, 0xE8, 0xEF, 0x93, 0x4B, 0x21, 0x58, 0x0A}}

type IAudioDeviceControllerInterface interface {
	win32.IInspectableInterface
	Put_Muted(value bool)
	Get_Muted() bool
	Put_VolumePercent(value float32)
	Get_VolumePercent() float32
}

type IAudioDeviceControllerVtbl struct {
	win32.IInspectableVtbl
	Put_Muted         uintptr
	Get_Muted         uintptr
	Put_VolumePercent uintptr
	Get_VolumePercent uintptr
}

type IAudioDeviceController struct {
	win32.IInspectable
}

func (this *IAudioDeviceController) Vtbl() *IAudioDeviceControllerVtbl {
	return (*IAudioDeviceControllerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioDeviceController) Put_Muted(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Muted, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAudioDeviceController) Get_Muted() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Muted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioDeviceController) Put_VolumePercent(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_VolumePercent, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAudioDeviceController) Get_VolumePercent() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VolumePercent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 86CFAC36-47C1-4B33-9852-8773EC4BE123
var IID_IAudioDeviceModule = syscall.GUID{0x86CFAC36, 0x47C1, 0x4B33,
	[8]byte{0x98, 0x52, 0x87, 0x73, 0xEC, 0x4B, 0xE1, 0x23}}

type IAudioDeviceModuleInterface interface {
	win32.IInspectableInterface
	Get_ClassId() string
	Get_DisplayName() string
	Get_InstanceId() uint32
	Get_MajorVersion() uint32
	Get_MinorVersion() uint32
	SendCommandAsync(Command *IBuffer) *IAsyncOperation[*IModuleCommandResult]
}

type IAudioDeviceModuleVtbl struct {
	win32.IInspectableVtbl
	Get_ClassId      uintptr
	Get_DisplayName  uintptr
	Get_InstanceId   uintptr
	Get_MajorVersion uintptr
	Get_MinorVersion uintptr
	SendCommandAsync uintptr
}

type IAudioDeviceModule struct {
	win32.IInspectable
}

func (this *IAudioDeviceModule) Vtbl() *IAudioDeviceModuleVtbl {
	return (*IAudioDeviceModuleVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioDeviceModule) Get_ClassId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ClassId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAudioDeviceModule) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IAudioDeviceModule) Get_InstanceId() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InstanceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioDeviceModule) Get_MajorVersion() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MajorVersion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioDeviceModule) Get_MinorVersion() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinorVersion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioDeviceModule) SendCommandAsync(Command *IBuffer) *IAsyncOperation[*IModuleCommandResult] {
	var _result *IAsyncOperation[*IModuleCommandResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SendCommandAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(Command)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E3E3CCAF-224C-48BE-956B-9A13134E96E8
var IID_IAudioDeviceModuleNotificationEventArgs = syscall.GUID{0xE3E3CCAF, 0x224C, 0x48BE,
	[8]byte{0x95, 0x6B, 0x9A, 0x13, 0x13, 0x4E, 0x96, 0xE8}}

type IAudioDeviceModuleNotificationEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Module() *IAudioDeviceModule
	Get_NotificationData() *IBuffer
}

type IAudioDeviceModuleNotificationEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Module           uintptr
	Get_NotificationData uintptr
}

type IAudioDeviceModuleNotificationEventArgs struct {
	win32.IInspectable
}

func (this *IAudioDeviceModuleNotificationEventArgs) Vtbl() *IAudioDeviceModuleNotificationEventArgsVtbl {
	return (*IAudioDeviceModuleNotificationEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioDeviceModuleNotificationEventArgs) Get_Module() *IAudioDeviceModule {
	var _result *IAudioDeviceModule
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Module, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioDeviceModuleNotificationEventArgs) Get_NotificationData() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NotificationData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6AA40C4D-960A-4D1C-B318-0022604547ED
var IID_IAudioDeviceModulesManager = syscall.GUID{0x6AA40C4D, 0x960A, 0x4D1C,
	[8]byte{0xB3, 0x18, 0x00, 0x22, 0x60, 0x45, 0x47, 0xED}}

type IAudioDeviceModulesManagerInterface interface {
	win32.IInspectableInterface
	Add_ModuleNotificationReceived(handler TypedEventHandler[*IAudioDeviceModulesManager, *IAudioDeviceModuleNotificationEventArgs]) EventRegistrationToken
	Remove_ModuleNotificationReceived(token EventRegistrationToken)
	FindAllById(moduleId string) *IVectorView[*IAudioDeviceModule]
	FindAll() *IVectorView[*IAudioDeviceModule]
}

type IAudioDeviceModulesManagerVtbl struct {
	win32.IInspectableVtbl
	Add_ModuleNotificationReceived    uintptr
	Remove_ModuleNotificationReceived uintptr
	FindAllById                       uintptr
	FindAll                           uintptr
}

type IAudioDeviceModulesManager struct {
	win32.IInspectable
}

func (this *IAudioDeviceModulesManager) Vtbl() *IAudioDeviceModulesManagerVtbl {
	return (*IAudioDeviceModulesManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioDeviceModulesManager) Add_ModuleNotificationReceived(handler TypedEventHandler[*IAudioDeviceModulesManager, *IAudioDeviceModuleNotificationEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ModuleNotificationReceived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAudioDeviceModulesManager) Remove_ModuleNotificationReceived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ModuleNotificationReceived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IAudioDeviceModulesManager) FindAllById(moduleId string) *IVectorView[*IAudioDeviceModule] {
	var _result *IVectorView[*IAudioDeviceModule]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllById, uintptr(unsafe.Pointer(this)), NewHStr(moduleId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAudioDeviceModulesManager) FindAll() *IVectorView[*IAudioDeviceModule] {
	var _result *IVectorView[*IAudioDeviceModule]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAll, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8DB03670-E64D-4773-96C0-BC7EBF0E063F
var IID_IAudioDeviceModulesManagerFactory = syscall.GUID{0x8DB03670, 0xE64D, 0x4773,
	[8]byte{0x96, 0xC0, 0xBC, 0x7E, 0xBF, 0x0E, 0x06, 0x3F}}

type IAudioDeviceModulesManagerFactoryInterface interface {
	win32.IInspectableInterface
	Create(deviceId string) *IAudioDeviceModulesManager
}

type IAudioDeviceModulesManagerFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IAudioDeviceModulesManagerFactory struct {
	win32.IInspectable
}

func (this *IAudioDeviceModulesManagerFactory) Vtbl() *IAudioDeviceModulesManagerFactoryVtbl {
	return (*IAudioDeviceModulesManagerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAudioDeviceModulesManagerFactory) Create(deviceId string) *IAudioDeviceModulesManager {
	var _result *IAudioDeviceModulesManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A520D0D6-AE8D-45DB-8011-CA49D3B3E578
var IID_ICallControl = syscall.GUID{0xA520D0D6, 0xAE8D, 0x45DB,
	[8]byte{0x80, 0x11, 0xCA, 0x49, 0xD3, 0xB3, 0xE5, 0x78}}

type ICallControlInterface interface {
	win32.IInspectableInterface
	IndicateNewIncomingCall(enableRinger bool, callerId string) uint64
	IndicateNewOutgoingCall() uint64
	IndicateActiveCall(callToken uint64)
	EndCall(callToken uint64)
	Get_HasRinger() bool
	Add_AnswerRequested(handler CallControlEventHandler) EventRegistrationToken
	Remove_AnswerRequested(token EventRegistrationToken)
	Add_HangUpRequested(handler CallControlEventHandler) EventRegistrationToken
	Remove_HangUpRequested(token EventRegistrationToken)
	Add_DialRequested(handler DialRequestedEventHandler) EventRegistrationToken
	Remove_DialRequested(token EventRegistrationToken)
	Add_RedialRequested(handler RedialRequestedEventHandler) EventRegistrationToken
	Remove_RedialRequested(token EventRegistrationToken)
	Add_KeypadPressed(handler KeypadPressedEventHandler) EventRegistrationToken
	Remove_KeypadPressed(token EventRegistrationToken)
	Add_AudioTransferRequested(handler CallControlEventHandler) EventRegistrationToken
	Remove_AudioTransferRequested(token EventRegistrationToken)
}

type ICallControlVtbl struct {
	win32.IInspectableVtbl
	IndicateNewIncomingCall       uintptr
	IndicateNewOutgoingCall       uintptr
	IndicateActiveCall            uintptr
	EndCall                       uintptr
	Get_HasRinger                 uintptr
	Add_AnswerRequested           uintptr
	Remove_AnswerRequested        uintptr
	Add_HangUpRequested           uintptr
	Remove_HangUpRequested        uintptr
	Add_DialRequested             uintptr
	Remove_DialRequested          uintptr
	Add_RedialRequested           uintptr
	Remove_RedialRequested        uintptr
	Add_KeypadPressed             uintptr
	Remove_KeypadPressed          uintptr
	Add_AudioTransferRequested    uintptr
	Remove_AudioTransferRequested uintptr
}

type ICallControl struct {
	win32.IInspectable
}

func (this *ICallControl) Vtbl() *ICallControlVtbl {
	return (*ICallControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICallControl) IndicateNewIncomingCall(enableRinger bool, callerId string) uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IndicateNewIncomingCall, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&enableRinger))), NewHStr(callerId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICallControl) IndicateNewOutgoingCall() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IndicateNewOutgoingCall, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICallControl) IndicateActiveCall(callToken uint64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IndicateActiveCall, uintptr(unsafe.Pointer(this)), uintptr(callToken))
	_ = _hr
}

func (this *ICallControl) EndCall(callToken uint64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EndCall, uintptr(unsafe.Pointer(this)), uintptr(callToken))
	_ = _hr
}

func (this *ICallControl) Get_HasRinger() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasRinger, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICallControl) Add_AnswerRequested(handler CallControlEventHandler) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AnswerRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewOneArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICallControl) Remove_AnswerRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AnswerRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ICallControl) Add_HangUpRequested(handler CallControlEventHandler) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_HangUpRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewOneArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICallControl) Remove_HangUpRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_HangUpRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ICallControl) Add_DialRequested(handler DialRequestedEventHandler) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DialRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICallControl) Remove_DialRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DialRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ICallControl) Add_RedialRequested(handler RedialRequestedEventHandler) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_RedialRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICallControl) Remove_RedialRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_RedialRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ICallControl) Add_KeypadPressed(handler KeypadPressedEventHandler) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_KeypadPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICallControl) Remove_KeypadPressed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_KeypadPressed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ICallControl) Add_AudioTransferRequested(handler CallControlEventHandler) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AudioTransferRequested, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewOneArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICallControl) Remove_AudioTransferRequested(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AudioTransferRequested, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 03945AD5-85AB-40E1-AF19-56C94303B019
var IID_ICallControlStatics = syscall.GUID{0x03945AD5, 0x85AB, 0x40E1,
	[8]byte{0xAF, 0x19, 0x56, 0xC9, 0x43, 0x03, 0xB0, 0x19}}

type ICallControlStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *ICallControl
	FromId(deviceId string) *ICallControl
}

type ICallControlStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
	FromId     uintptr
}

type ICallControlStatics struct {
	win32.IInspectable
}

func (this *ICallControlStatics) Vtbl() *ICallControlStaticsVtbl {
	return (*ICallControlStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICallControlStatics) GetDefault() *ICallControl {
	var _result *ICallControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICallControlStatics) FromId(deviceId string) *ICallControl {
	var _result *ICallControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromId, uintptr(unsafe.Pointer(this)), NewHStr(deviceId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AF6C4AD0-A84D-5DB6-BE58-A5DA21CFE011
var IID_ICameraOcclusionInfo = syscall.GUID{0xAF6C4AD0, 0xA84D, 0x5DB6,
	[8]byte{0xBE, 0x58, 0xA5, 0xDA, 0x21, 0xCF, 0xE0, 0x11}}

type ICameraOcclusionInfoInterface interface {
	win32.IInspectableInterface
	GetState() *ICameraOcclusionState
	IsOcclusionKindSupported(occlusionKind CameraOcclusionKind) bool
	Add_StateChanged(handler TypedEventHandler[*ICameraOcclusionInfo, *ICameraOcclusionStateChangedEventArgs]) EventRegistrationToken
	Remove_StateChanged(token EventRegistrationToken)
}

type ICameraOcclusionInfoVtbl struct {
	win32.IInspectableVtbl
	GetState                 uintptr
	IsOcclusionKindSupported uintptr
	Add_StateChanged         uintptr
	Remove_StateChanged      uintptr
}

type ICameraOcclusionInfo struct {
	win32.IInspectable
}

func (this *ICameraOcclusionInfo) Vtbl() *ICameraOcclusionInfoVtbl {
	return (*ICameraOcclusionInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICameraOcclusionInfo) GetState() *ICameraOcclusionState {
	var _result *ICameraOcclusionState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICameraOcclusionInfo) IsOcclusionKindSupported(occlusionKind CameraOcclusionKind) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsOcclusionKindSupported, uintptr(unsafe.Pointer(this)), uintptr(occlusionKind), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICameraOcclusionInfo) Add_StateChanged(handler TypedEventHandler[*ICameraOcclusionInfo, *ICameraOcclusionStateChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICameraOcclusionInfo) Remove_StateChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 430ADEB8-6842-5E55-9BDE-04B4EF3A8A57
var IID_ICameraOcclusionState = syscall.GUID{0x430ADEB8, 0x6842, 0x5E55,
	[8]byte{0x9B, 0xDE, 0x04, 0xB4, 0xEF, 0x3A, 0x8A, 0x57}}

type ICameraOcclusionStateInterface interface {
	win32.IInspectableInterface
	Get_IsOccluded() bool
	IsOcclusionKind(occlusionKind CameraOcclusionKind) bool
}

type ICameraOcclusionStateVtbl struct {
	win32.IInspectableVtbl
	Get_IsOccluded  uintptr
	IsOcclusionKind uintptr
}

type ICameraOcclusionState struct {
	win32.IInspectable
}

func (this *ICameraOcclusionState) Vtbl() *ICameraOcclusionStateVtbl {
	return (*ICameraOcclusionStateVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICameraOcclusionState) Get_IsOccluded() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsOccluded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICameraOcclusionState) IsOcclusionKind(occlusionKind CameraOcclusionKind) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsOcclusionKind, uintptr(unsafe.Pointer(this)), uintptr(occlusionKind), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 8512D848-C0DE-57CA-A1CA-FB2C3D23DF55
var IID_ICameraOcclusionStateChangedEventArgs = syscall.GUID{0x8512D848, 0xC0DE, 0x57CA,
	[8]byte{0xA1, 0xCA, 0xFB, 0x2C, 0x3D, 0x23, 0xDF, 0x55}}

type ICameraOcclusionStateChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_State() *ICameraOcclusionState
}

type ICameraOcclusionStateChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_State uintptr
}

type ICameraOcclusionStateChangedEventArgs struct {
	win32.IInspectable
}

func (this *ICameraOcclusionStateChangedEventArgs) Vtbl() *ICameraOcclusionStateChangedEventArgsVtbl {
	return (*ICameraOcclusionStateChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICameraOcclusionStateChangedEventArgs) Get_State() *ICameraOcclusionState {
	var _result *ICameraOcclusionState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 110F882F-1C05-4657-A18E-47C9B69F07AB
var IID_IDefaultAudioDeviceChangedEventArgs = syscall.GUID{0x110F882F, 0x1C05, 0x4657,
	[8]byte{0xA1, 0x8E, 0x47, 0xC9, 0xB6, 0x9F, 0x07, 0xAB}}

type IDefaultAudioDeviceChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	Get_Role() AudioDeviceRole
}

type IDefaultAudioDeviceChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Id   uintptr
	Get_Role uintptr
}

type IDefaultAudioDeviceChangedEventArgs struct {
	win32.IInspectable
}

func (this *IDefaultAudioDeviceChangedEventArgs) Vtbl() *IDefaultAudioDeviceChangedEventArgsVtbl {
	return (*IDefaultAudioDeviceChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDefaultAudioDeviceChangedEventArgs) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDefaultAudioDeviceChangedEventArgs) Get_Role() AudioDeviceRole {
	var _result AudioDeviceRole
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Role, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 037B929E-953C-4286-8866-4F0F376C855A
var IID_IDialRequestedEventArgs = syscall.GUID{0x037B929E, 0x953C, 0x4286,
	[8]byte{0x88, 0x66, 0x4F, 0x0F, 0x37, 0x6C, 0x85, 0x5A}}

type IDialRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	Handled()
	Get_Contact() interface{}
}

type IDialRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Handled     uintptr
	Get_Contact uintptr
}

type IDialRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IDialRequestedEventArgs) Vtbl() *IDialRequestedEventArgsVtbl {
	return (*IDialRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDialRequestedEventArgs) Handled() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Handled, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IDialRequestedEventArgs) Get_Contact() interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Contact, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// DD4F21DD-D173-5C6B-8C25-BDD26D5122B1
var IID_IDigitalWindowBounds = syscall.GUID{0xDD4F21DD, 0xD173, 0x5C6B,
	[8]byte{0x8C, 0x25, 0xBD, 0xD2, 0x6D, 0x51, 0x22, 0xB1}}

type IDigitalWindowBoundsInterface interface {
	win32.IInspectableInterface
	Get_NormalizedOriginTop() float64
	Put_NormalizedOriginTop(value float64)
	Get_NormalizedOriginLeft() float64
	Put_NormalizedOriginLeft(value float64)
	Get_Scale() float64
	Put_Scale(value float64)
}

type IDigitalWindowBoundsVtbl struct {
	win32.IInspectableVtbl
	Get_NormalizedOriginTop  uintptr
	Put_NormalizedOriginTop  uintptr
	Get_NormalizedOriginLeft uintptr
	Put_NormalizedOriginLeft uintptr
	Get_Scale                uintptr
	Put_Scale                uintptr
}

type IDigitalWindowBounds struct {
	win32.IInspectable
}

func (this *IDigitalWindowBounds) Vtbl() *IDigitalWindowBoundsVtbl {
	return (*IDigitalWindowBoundsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDigitalWindowBounds) Get_NormalizedOriginTop() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NormalizedOriginTop, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDigitalWindowBounds) Put_NormalizedOriginTop(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_NormalizedOriginTop, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDigitalWindowBounds) Get_NormalizedOriginLeft() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NormalizedOriginLeft, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDigitalWindowBounds) Put_NormalizedOriginLeft(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_NormalizedOriginLeft, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDigitalWindowBounds) Get_Scale() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Scale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDigitalWindowBounds) Put_Scale(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Scale, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// D78BAD2C-F721-5244-A196-B56CCBEC606C
var IID_IDigitalWindowCapability = syscall.GUID{0xD78BAD2C, 0xF721, 0x5244,
	[8]byte{0xA1, 0x96, 0xB5, 0x6C, 0xCB, 0xEC, 0x60, 0x6C}}

type IDigitalWindowCapabilityInterface interface {
	win32.IInspectableInterface
	Get_Width() int32
	Get_Height() int32
	Get_MinScaleValue() float64
	Get_MaxScaleValue() float64
	Get_MinScaleValueWithoutUpsampling() float64
	Get_NormalizedFieldOfViewLimit() Rect
}

type IDigitalWindowCapabilityVtbl struct {
	win32.IInspectableVtbl
	Get_Width                          uintptr
	Get_Height                         uintptr
	Get_MinScaleValue                  uintptr
	Get_MaxScaleValue                  uintptr
	Get_MinScaleValueWithoutUpsampling uintptr
	Get_NormalizedFieldOfViewLimit     uintptr
}

type IDigitalWindowCapability struct {
	win32.IInspectable
}

func (this *IDigitalWindowCapability) Vtbl() *IDigitalWindowCapabilityVtbl {
	return (*IDigitalWindowCapabilityVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDigitalWindowCapability) Get_Width() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Width, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDigitalWindowCapability) Get_Height() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Height, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDigitalWindowCapability) Get_MinScaleValue() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinScaleValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDigitalWindowCapability) Get_MaxScaleValue() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxScaleValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDigitalWindowCapability) Get_MinScaleValueWithoutUpsampling() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinScaleValueWithoutUpsampling, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDigitalWindowCapability) Get_NormalizedFieldOfViewLimit() Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NormalizedFieldOfViewLimit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 23B69EFF-65D2-53EA-8780-DE582B48B544
var IID_IDigitalWindowControl = syscall.GUID{0x23B69EFF, 0x65D2, 0x53EA,
	[8]byte{0x87, 0x80, 0xDE, 0x58, 0x2B, 0x48, 0xB5, 0x44}}

type IDigitalWindowControlInterface interface {
	win32.IInspectableInterface
	Get_IsSupported() bool
	Get_SupportedModes() []DigitalWindowMode
	Get_CurrentMode() DigitalWindowMode
	GetBounds() *IDigitalWindowBounds
	Configure(digitalWindowMode DigitalWindowMode)
	ConfigureWithBounds(digitalWindowMode DigitalWindowMode, digitalWindowBounds *IDigitalWindowBounds)
	Get_SupportedCapabilities() *IVectorView[*IDigitalWindowCapability]
	GetCapabilityForSize(width int32, height int32) *IDigitalWindowCapability
}

type IDigitalWindowControlVtbl struct {
	win32.IInspectableVtbl
	Get_IsSupported           uintptr
	Get_SupportedModes        uintptr
	Get_CurrentMode           uintptr
	GetBounds                 uintptr
	Configure                 uintptr
	ConfigureWithBounds       uintptr
	Get_SupportedCapabilities uintptr
	GetCapabilityForSize      uintptr
}

type IDigitalWindowControl struct {
	win32.IInspectable
}

func (this *IDigitalWindowControl) Vtbl() *IDigitalWindowControlVtbl {
	return (*IDigitalWindowControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDigitalWindowControl) Get_IsSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDigitalWindowControl) Get_SupportedModes() []DigitalWindowMode {
	var _result []DigitalWindowMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedModes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDigitalWindowControl) Get_CurrentMode() DigitalWindowMode {
	var _result DigitalWindowMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDigitalWindowControl) GetBounds() *IDigitalWindowBounds {
	var _result *IDigitalWindowBounds
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetBounds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDigitalWindowControl) Configure(digitalWindowMode DigitalWindowMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Configure, uintptr(unsafe.Pointer(this)), uintptr(digitalWindowMode))
	_ = _hr
}

func (this *IDigitalWindowControl) ConfigureWithBounds(digitalWindowMode DigitalWindowMode, digitalWindowBounds *IDigitalWindowBounds) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ConfigureWithBounds, uintptr(unsafe.Pointer(this)), uintptr(digitalWindowMode), uintptr(unsafe.Pointer(digitalWindowBounds)))
	_ = _hr
}

func (this *IDigitalWindowControl) Get_SupportedCapabilities() *IVectorView[*IDigitalWindowCapability] {
	var _result *IVectorView[*IDigitalWindowCapability]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedCapabilities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDigitalWindowControl) GetCapabilityForSize(width int32, height int32) *IDigitalWindowCapability {
	var _result *IDigitalWindowCapability
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCapabilityForSize, uintptr(unsafe.Pointer(this)), uintptr(width), uintptr(height), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 81C8E834-DCEC-4011-A610-1F3847E64ACA
var IID_IExposureCompensationControl = syscall.GUID{0x81C8E834, 0xDCEC, 0x4011,
	[8]byte{0xA6, 0x10, 0x1F, 0x38, 0x47, 0xE6, 0x4A, 0xCA}}

type IExposureCompensationControlInterface interface {
	win32.IInspectableInterface
	Get_Supported() bool
	Get_Min() float32
	Get_Max() float32
	Get_Step() float32
	Get_Value() float32
	SetValueAsync(value float32) *IAsyncAction
}

type IExposureCompensationControlVtbl struct {
	win32.IInspectableVtbl
	Get_Supported uintptr
	Get_Min       uintptr
	Get_Max       uintptr
	Get_Step      uintptr
	Get_Value     uintptr
	SetValueAsync uintptr
}

type IExposureCompensationControl struct {
	win32.IInspectable
}

func (this *IExposureCompensationControl) Vtbl() *IExposureCompensationControlVtbl {
	return (*IExposureCompensationControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IExposureCompensationControl) Get_Supported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Supported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IExposureCompensationControl) Get_Min() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Min, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IExposureCompensationControl) Get_Max() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Max, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IExposureCompensationControl) Get_Step() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Step, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IExposureCompensationControl) Get_Value() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IExposureCompensationControl) SetValueAsync(value float32) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetValueAsync, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 09E8CBE2-AD96-4F28-A0E0-96ED7E1B5FD2
var IID_IExposureControl = syscall.GUID{0x09E8CBE2, 0xAD96, 0x4F28,
	[8]byte{0xA0, 0xE0, 0x96, 0xED, 0x7E, 0x1B, 0x5F, 0xD2}}

type IExposureControlInterface interface {
	win32.IInspectableInterface
	Get_Supported() bool
	Get_Auto() bool
	SetAutoAsync(value bool) *IAsyncAction
	Get_Min() TimeSpan
	Get_Max() TimeSpan
	Get_Step() TimeSpan
	Get_Value() TimeSpan
	SetValueAsync(shutterDuration TimeSpan) *IAsyncAction
}

type IExposureControlVtbl struct {
	win32.IInspectableVtbl
	Get_Supported uintptr
	Get_Auto      uintptr
	SetAutoAsync  uintptr
	Get_Min       uintptr
	Get_Max       uintptr
	Get_Step      uintptr
	Get_Value     uintptr
	SetValueAsync uintptr
}

type IExposureControl struct {
	win32.IInspectable
}

func (this *IExposureControl) Vtbl() *IExposureControlVtbl {
	return (*IExposureControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IExposureControl) Get_Supported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Supported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IExposureControl) Get_Auto() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Auto, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IExposureControl) SetAutoAsync(value bool) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetAutoAsync, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IExposureControl) Get_Min() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Min, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IExposureControl) Get_Max() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Max, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IExposureControl) Get_Step() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Step, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IExposureControl) Get_Value() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IExposureControl) SetValueAsync(shutterDuration TimeSpan) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetValueAsync, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&shutterDuration)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2CB240A3-5168-4271-9EA5-47621A98A352
var IID_IExposurePriorityVideoControl = syscall.GUID{0x2CB240A3, 0x5168, 0x4271,
	[8]byte{0x9E, 0xA5, 0x47, 0x62, 0x1A, 0x98, 0xA3, 0x52}}

type IExposurePriorityVideoControlInterface interface {
	win32.IInspectableInterface
	Get_Supported() bool
	Get_Enabled() bool
	Put_Enabled(value bool)
}

type IExposurePriorityVideoControlVtbl struct {
	win32.IInspectableVtbl
	Get_Supported uintptr
	Get_Enabled   uintptr
	Put_Enabled   uintptr
}

type IExposurePriorityVideoControl struct {
	win32.IInspectable
}

func (this *IExposurePriorityVideoControl) Vtbl() *IExposurePriorityVideoControlVtbl {
	return (*IExposurePriorityVideoControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IExposurePriorityVideoControl) Get_Supported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Supported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IExposurePriorityVideoControl) Get_Enabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Enabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IExposurePriorityVideoControl) Put_Enabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Enabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// DEF41DBE-7D68-45E3-8C0F-BE7BB32837D0
var IID_IFlashControl = syscall.GUID{0xDEF41DBE, 0x7D68, 0x45E3,
	[8]byte{0x8C, 0x0F, 0xBE, 0x7B, 0xB3, 0x28, 0x37, 0xD0}}

type IFlashControlInterface interface {
	win32.IInspectableInterface
	Get_Supported() bool
	Get_PowerSupported() bool
	Get_RedEyeReductionSupported() bool
	Get_Enabled() bool
	Put_Enabled(value bool)
	Get_Auto() bool
	Put_Auto(value bool)
	Get_RedEyeReduction() bool
	Put_RedEyeReduction(value bool)
	Get_PowerPercent() float32
	Put_PowerPercent(value float32)
}

type IFlashControlVtbl struct {
	win32.IInspectableVtbl
	Get_Supported                uintptr
	Get_PowerSupported           uintptr
	Get_RedEyeReductionSupported uintptr
	Get_Enabled                  uintptr
	Put_Enabled                  uintptr
	Get_Auto                     uintptr
	Put_Auto                     uintptr
	Get_RedEyeReduction          uintptr
	Put_RedEyeReduction          uintptr
	Get_PowerPercent             uintptr
	Put_PowerPercent             uintptr
}

type IFlashControl struct {
	win32.IInspectable
}

func (this *IFlashControl) Vtbl() *IFlashControlVtbl {
	return (*IFlashControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFlashControl) Get_Supported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Supported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFlashControl) Get_PowerSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PowerSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFlashControl) Get_RedEyeReductionSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RedEyeReductionSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFlashControl) Get_Enabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Enabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFlashControl) Put_Enabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Enabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IFlashControl) Get_Auto() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Auto, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFlashControl) Put_Auto(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Auto, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IFlashControl) Get_RedEyeReduction() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RedEyeReduction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFlashControl) Put_RedEyeReduction(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RedEyeReduction, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IFlashControl) Get_PowerPercent() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PowerPercent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFlashControl) Put_PowerPercent(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PowerPercent, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 7D29CC9E-75E1-4AF7-BD7D-4E38E1C06CD6
var IID_IFlashControl2 = syscall.GUID{0x7D29CC9E, 0x75E1, 0x4AF7,
	[8]byte{0xBD, 0x7D, 0x4E, 0x38, 0xE1, 0xC0, 0x6C, 0xD6}}

type IFlashControl2Interface interface {
	win32.IInspectableInterface
	Get_AssistantLightSupported() bool
	Get_AssistantLightEnabled() bool
	Put_AssistantLightEnabled(value bool)
}

type IFlashControl2Vtbl struct {
	win32.IInspectableVtbl
	Get_AssistantLightSupported uintptr
	Get_AssistantLightEnabled   uintptr
	Put_AssistantLightEnabled   uintptr
}

type IFlashControl2 struct {
	win32.IInspectable
}

func (this *IFlashControl2) Vtbl() *IFlashControl2Vtbl {
	return (*IFlashControl2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFlashControl2) Get_AssistantLightSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AssistantLightSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFlashControl2) Get_AssistantLightEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AssistantLightEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFlashControl2) Put_AssistantLightEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AssistantLightEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// C0D889F6-5228-4453-B153-85606592B238
var IID_IFocusControl = syscall.GUID{0xC0D889F6, 0x5228, 0x4453,
	[8]byte{0xB1, 0x53, 0x85, 0x60, 0x65, 0x92, 0xB2, 0x38}}

type IFocusControlInterface interface {
	win32.IInspectableInterface
	Get_Supported() bool
	Get_SupportedPresets() *IVectorView[FocusPreset]
	Get_Preset() FocusPreset
	SetPresetAsync(preset FocusPreset) *IAsyncAction
	SetPresetWithCompletionOptionAsync(preset FocusPreset, completeBeforeFocus bool) *IAsyncAction
	Get_Min() uint32
	Get_Max() uint32
	Get_Step() uint32
	Get_Value() uint32
	SetValueAsync(focus uint32) *IAsyncAction
	FocusAsync() *IAsyncAction
}

type IFocusControlVtbl struct {
	win32.IInspectableVtbl
	Get_Supported                      uintptr
	Get_SupportedPresets               uintptr
	Get_Preset                         uintptr
	SetPresetAsync                     uintptr
	SetPresetWithCompletionOptionAsync uintptr
	Get_Min                            uintptr
	Get_Max                            uintptr
	Get_Step                           uintptr
	Get_Value                          uintptr
	SetValueAsync                      uintptr
	FocusAsync                         uintptr
}

type IFocusControl struct {
	win32.IInspectable
}

func (this *IFocusControl) Vtbl() *IFocusControlVtbl {
	return (*IFocusControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFocusControl) Get_Supported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Supported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFocusControl) Get_SupportedPresets() *IVectorView[FocusPreset] {
	var _result *IVectorView[FocusPreset]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedPresets, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFocusControl) Get_Preset() FocusPreset {
	var _result FocusPreset
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Preset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFocusControl) SetPresetAsync(preset FocusPreset) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetPresetAsync, uintptr(unsafe.Pointer(this)), uintptr(preset), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFocusControl) SetPresetWithCompletionOptionAsync(preset FocusPreset, completeBeforeFocus bool) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetPresetWithCompletionOptionAsync, uintptr(unsafe.Pointer(this)), uintptr(preset), uintptr(*(*byte)(unsafe.Pointer(&completeBeforeFocus))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFocusControl) Get_Min() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Min, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFocusControl) Get_Max() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Max, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFocusControl) Get_Step() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Step, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFocusControl) Get_Value() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFocusControl) SetValueAsync(focus uint32) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetValueAsync, uintptr(unsafe.Pointer(this)), uintptr(focus), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFocusControl) FocusAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FocusAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3F7CFF48-C534-4E9E-94C3-52EF2AFD5D07
var IID_IFocusControl2 = syscall.GUID{0x3F7CFF48, 0xC534, 0x4E9E,
	[8]byte{0x94, 0xC3, 0x52, 0xEF, 0x2A, 0xFD, 0x5D, 0x07}}

type IFocusControl2Interface interface {
	win32.IInspectableInterface
	Get_FocusChangedSupported() bool
	Get_WaitForFocusSupported() bool
	Get_SupportedFocusModes() *IVectorView[FocusMode]
	Get_SupportedFocusDistances() *IVectorView[ManualFocusDistance]
	Get_SupportedFocusRanges() *IVectorView[AutoFocusRange]
	Get_Mode() FocusMode
	Get_FocusState() MediaCaptureFocusState
	UnlockAsync() *IAsyncAction
	LockAsync() *IAsyncAction
	Configure(settings *IFocusSettings)
}

type IFocusControl2Vtbl struct {
	win32.IInspectableVtbl
	Get_FocusChangedSupported   uintptr
	Get_WaitForFocusSupported   uintptr
	Get_SupportedFocusModes     uintptr
	Get_SupportedFocusDistances uintptr
	Get_SupportedFocusRanges    uintptr
	Get_Mode                    uintptr
	Get_FocusState              uintptr
	UnlockAsync                 uintptr
	LockAsync                   uintptr
	Configure                   uintptr
}

type IFocusControl2 struct {
	win32.IInspectable
}

func (this *IFocusControl2) Vtbl() *IFocusControl2Vtbl {
	return (*IFocusControl2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFocusControl2) Get_FocusChangedSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FocusChangedSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFocusControl2) Get_WaitForFocusSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WaitForFocusSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFocusControl2) Get_SupportedFocusModes() *IVectorView[FocusMode] {
	var _result *IVectorView[FocusMode]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedFocusModes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFocusControl2) Get_SupportedFocusDistances() *IVectorView[ManualFocusDistance] {
	var _result *IVectorView[ManualFocusDistance]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedFocusDistances, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFocusControl2) Get_SupportedFocusRanges() *IVectorView[AutoFocusRange] {
	var _result *IVectorView[AutoFocusRange]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedFocusRanges, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFocusControl2) Get_Mode() FocusMode {
	var _result FocusMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFocusControl2) Get_FocusState() MediaCaptureFocusState {
	var _result MediaCaptureFocusState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FocusState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFocusControl2) UnlockAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UnlockAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFocusControl2) LockAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LockAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFocusControl2) Configure(settings *IFocusSettings) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Configure, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(settings)))
	_ = _hr
}

// 79958F6B-3263-4275-85D6-AEAE891C96EE
var IID_IFocusSettings = syscall.GUID{0x79958F6B, 0x3263, 0x4275,
	[8]byte{0x85, 0xD6, 0xAE, 0xAE, 0x89, 0x1C, 0x96, 0xEE}}

type IFocusSettingsInterface interface {
	win32.IInspectableInterface
	Get_Mode() FocusMode
	Put_Mode(value FocusMode)
	Get_AutoFocusRange() AutoFocusRange
	Put_AutoFocusRange(value AutoFocusRange)
	Get_Value() *IReference[uint32]
	Put_Value(value *IReference[uint32])
	Get_Distance() *IReference[ManualFocusDistance]
	Put_Distance(value *IReference[ManualFocusDistance])
	Get_WaitForFocus() bool
	Put_WaitForFocus(value bool)
	Get_DisableDriverFallback() bool
	Put_DisableDriverFallback(value bool)
}

type IFocusSettingsVtbl struct {
	win32.IInspectableVtbl
	Get_Mode                  uintptr
	Put_Mode                  uintptr
	Get_AutoFocusRange        uintptr
	Put_AutoFocusRange        uintptr
	Get_Value                 uintptr
	Put_Value                 uintptr
	Get_Distance              uintptr
	Put_Distance              uintptr
	Get_WaitForFocus          uintptr
	Put_WaitForFocus          uintptr
	Get_DisableDriverFallback uintptr
	Put_DisableDriverFallback uintptr
}

type IFocusSettings struct {
	win32.IInspectable
}

func (this *IFocusSettings) Vtbl() *IFocusSettingsVtbl {
	return (*IFocusSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFocusSettings) Get_Mode() FocusMode {
	var _result FocusMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFocusSettings) Put_Mode(value FocusMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Mode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IFocusSettings) Get_AutoFocusRange() AutoFocusRange {
	var _result AutoFocusRange
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AutoFocusRange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFocusSettings) Put_AutoFocusRange(value AutoFocusRange) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AutoFocusRange, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IFocusSettings) Get_Value() *IReference[uint32] {
	var _result *IReference[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFocusSettings) Put_Value(value *IReference[uint32]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IFocusSettings) Get_Distance() *IReference[ManualFocusDistance] {
	var _result *IReference[ManualFocusDistance]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Distance, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFocusSettings) Put_Distance(value *IReference[ManualFocusDistance]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Distance, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IFocusSettings) Get_WaitForFocus() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WaitForFocus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFocusSettings) Put_WaitForFocus(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_WaitForFocus, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IFocusSettings) Get_DisableDriverFallback() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisableDriverFallback, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFocusSettings) Put_DisableDriverFallback(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DisableDriverFallback, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 55D8E2D0-30C0-43BF-9B9A-9799D70CED94
var IID_IHdrVideoControl = syscall.GUID{0x55D8E2D0, 0x30C0, 0x43BF,
	[8]byte{0x9B, 0x9A, 0x97, 0x99, 0xD7, 0x0C, 0xED, 0x94}}

type IHdrVideoControlInterface interface {
	win32.IInspectableInterface
	Get_Supported() bool
	Get_SupportedModes() *IVectorView[HdrVideoMode]
	Get_Mode() HdrVideoMode
	Put_Mode(value HdrVideoMode)
}

type IHdrVideoControlVtbl struct {
	win32.IInspectableVtbl
	Get_Supported      uintptr
	Get_SupportedModes uintptr
	Get_Mode           uintptr
	Put_Mode           uintptr
}

type IHdrVideoControl struct {
	win32.IInspectable
}

func (this *IHdrVideoControl) Vtbl() *IHdrVideoControlVtbl {
	return (*IHdrVideoControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHdrVideoControl) Get_Supported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Supported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHdrVideoControl) Get_SupportedModes() *IVectorView[HdrVideoMode] {
	var _result *IVectorView[HdrVideoMode]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedModes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IHdrVideoControl) Get_Mode() HdrVideoMode {
	var _result HdrVideoMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IHdrVideoControl) Put_Mode(value HdrVideoMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Mode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 1CBA2C83-6CB6-5A04-A6FC-3BE7B33FF056
var IID_IInfraredTorchControl = syscall.GUID{0x1CBA2C83, 0x6CB6, 0x5A04,
	[8]byte{0xA6, 0xFC, 0x3B, 0xE7, 0xB3, 0x3F, 0xF0, 0x56}}

type IInfraredTorchControlInterface interface {
	win32.IInspectableInterface
	Get_IsSupported() bool
	Get_SupportedModes() *IVectorView[InfraredTorchMode]
	Get_CurrentMode() InfraredTorchMode
	Put_CurrentMode(value InfraredTorchMode)
	Get_MinPower() int32
	Get_MaxPower() int32
	Get_PowerStep() int32
	Get_Power() int32
	Put_Power(value int32)
}

type IInfraredTorchControlVtbl struct {
	win32.IInspectableVtbl
	Get_IsSupported    uintptr
	Get_SupportedModes uintptr
	Get_CurrentMode    uintptr
	Put_CurrentMode    uintptr
	Get_MinPower       uintptr
	Get_MaxPower       uintptr
	Get_PowerStep      uintptr
	Get_Power          uintptr
	Put_Power          uintptr
}

type IInfraredTorchControl struct {
	win32.IInspectable
}

func (this *IInfraredTorchControl) Vtbl() *IInfraredTorchControlVtbl {
	return (*IInfraredTorchControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInfraredTorchControl) Get_IsSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInfraredTorchControl) Get_SupportedModes() *IVectorView[InfraredTorchMode] {
	var _result *IVectorView[InfraredTorchMode]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedModes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInfraredTorchControl) Get_CurrentMode() InfraredTorchMode {
	var _result InfraredTorchMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInfraredTorchControl) Put_CurrentMode(value InfraredTorchMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CurrentMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IInfraredTorchControl) Get_MinPower() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinPower, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInfraredTorchControl) Get_MaxPower() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxPower, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInfraredTorchControl) Get_PowerStep() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PowerStep, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInfraredTorchControl) Get_Power() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Power, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInfraredTorchControl) Put_Power(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Power, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 27B6C322-25AD-4F1B-AAAB-524AB376CA33
var IID_IIsoSpeedControl = syscall.GUID{0x27B6C322, 0x25AD, 0x4F1B,
	[8]byte{0xAA, 0xAB, 0x52, 0x4A, 0xB3, 0x76, 0xCA, 0x33}}

type IIsoSpeedControlInterface interface {
	win32.IInspectableInterface
	Get_Supported() bool
	Get_SupportedPresets() *IVectorView[IsoSpeedPreset]
	Get_Preset() IsoSpeedPreset
	SetPresetAsync(preset IsoSpeedPreset) *IAsyncAction
}

type IIsoSpeedControlVtbl struct {
	win32.IInspectableVtbl
	Get_Supported        uintptr
	Get_SupportedPresets uintptr
	Get_Preset           uintptr
	SetPresetAsync       uintptr
}

type IIsoSpeedControl struct {
	win32.IInspectable
}

func (this *IIsoSpeedControl) Vtbl() *IIsoSpeedControlVtbl {
	return (*IIsoSpeedControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IIsoSpeedControl) Get_Supported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Supported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IIsoSpeedControl) Get_SupportedPresets() *IVectorView[IsoSpeedPreset] {
	var _result *IVectorView[IsoSpeedPreset]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedPresets, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IIsoSpeedControl) Get_Preset() IsoSpeedPreset {
	var _result IsoSpeedPreset
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Preset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IIsoSpeedControl) SetPresetAsync(preset IsoSpeedPreset) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetPresetAsync, uintptr(unsafe.Pointer(this)), uintptr(preset), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6F1578F2-6D77-4F8A-8C2F-6130B6395053
var IID_IIsoSpeedControl2 = syscall.GUID{0x6F1578F2, 0x6D77, 0x4F8A,
	[8]byte{0x8C, 0x2F, 0x61, 0x30, 0xB6, 0x39, 0x50, 0x53}}

type IIsoSpeedControl2Interface interface {
	win32.IInspectableInterface
	Get_Min() uint32
	Get_Max() uint32
	Get_Step() uint32
	Get_Value() uint32
	SetValueAsync(isoSpeed uint32) *IAsyncAction
	Get_Auto() bool
	SetAutoAsync() *IAsyncAction
}

type IIsoSpeedControl2Vtbl struct {
	win32.IInspectableVtbl
	Get_Min       uintptr
	Get_Max       uintptr
	Get_Step      uintptr
	Get_Value     uintptr
	SetValueAsync uintptr
	Get_Auto      uintptr
	SetAutoAsync  uintptr
}

type IIsoSpeedControl2 struct {
	win32.IInspectable
}

func (this *IIsoSpeedControl2) Vtbl() *IIsoSpeedControl2Vtbl {
	return (*IIsoSpeedControl2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IIsoSpeedControl2) Get_Min() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Min, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IIsoSpeedControl2) Get_Max() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Max, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IIsoSpeedControl2) Get_Step() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Step, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IIsoSpeedControl2) Get_Value() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IIsoSpeedControl2) SetValueAsync(isoSpeed uint32) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetValueAsync, uintptr(unsafe.Pointer(this)), uintptr(isoSpeed), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IIsoSpeedControl2) Get_Auto() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Auto, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IIsoSpeedControl2) SetAutoAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetAutoAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D3A43900-B4FA-49CD-9442-89AF6568F601
var IID_IKeypadPressedEventArgs = syscall.GUID{0xD3A43900, 0xB4FA, 0x49CD,
	[8]byte{0x94, 0x42, 0x89, 0xAF, 0x65, 0x68, 0xF6, 0x01}}

type IKeypadPressedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_TelephonyKey() TelephonyKey
}

type IKeypadPressedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_TelephonyKey uintptr
}

type IKeypadPressedEventArgs struct {
	win32.IInspectable
}

func (this *IKeypadPressedEventArgs) Vtbl() *IKeypadPressedEventArgsVtbl {
	return (*IKeypadPressedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKeypadPressedEventArgs) Get_TelephonyKey() TelephonyKey {
	var _result TelephonyKey
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TelephonyKey, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 6D5C4DD0-FADF-415D-AEE6-3BAA529300C9
var IID_ILowLagPhotoControl = syscall.GUID{0x6D5C4DD0, 0xFADF, 0x415D,
	[8]byte{0xAE, 0xE6, 0x3B, 0xAA, 0x52, 0x93, 0x00, 0xC9}}

type ILowLagPhotoControlInterface interface {
	win32.IInspectableInterface
	GetHighestConcurrentFrameRate(captureProperties *IMediaEncodingProperties) *IMediaRatio
	GetCurrentFrameRate() *IMediaRatio
	Get_ThumbnailEnabled() bool
	Put_ThumbnailEnabled(value bool)
	Get_ThumbnailFormat() MediaThumbnailFormat
	Put_ThumbnailFormat(value MediaThumbnailFormat)
	Get_DesiredThumbnailSize() uint32
	Put_DesiredThumbnailSize(value uint32)
	Get_HardwareAcceleratedThumbnailSupported() uint32
}

type ILowLagPhotoControlVtbl struct {
	win32.IInspectableVtbl
	GetHighestConcurrentFrameRate             uintptr
	GetCurrentFrameRate                       uintptr
	Get_ThumbnailEnabled                      uintptr
	Put_ThumbnailEnabled                      uintptr
	Get_ThumbnailFormat                       uintptr
	Put_ThumbnailFormat                       uintptr
	Get_DesiredThumbnailSize                  uintptr
	Put_DesiredThumbnailSize                  uintptr
	Get_HardwareAcceleratedThumbnailSupported uintptr
}

type ILowLagPhotoControl struct {
	win32.IInspectable
}

func (this *ILowLagPhotoControl) Vtbl() *ILowLagPhotoControlVtbl {
	return (*ILowLagPhotoControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILowLagPhotoControl) GetHighestConcurrentFrameRate(captureProperties *IMediaEncodingProperties) *IMediaRatio {
	var _result *IMediaRatio
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetHighestConcurrentFrameRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(captureProperties)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILowLagPhotoControl) GetCurrentFrameRate() *IMediaRatio {
	var _result *IMediaRatio
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentFrameRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILowLagPhotoControl) Get_ThumbnailEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ThumbnailEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILowLagPhotoControl) Put_ThumbnailEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ThumbnailEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ILowLagPhotoControl) Get_ThumbnailFormat() MediaThumbnailFormat {
	var _result MediaThumbnailFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ThumbnailFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILowLagPhotoControl) Put_ThumbnailFormat(value MediaThumbnailFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ThumbnailFormat, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ILowLagPhotoControl) Get_DesiredThumbnailSize() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesiredThumbnailSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILowLagPhotoControl) Put_DesiredThumbnailSize(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DesiredThumbnailSize, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ILowLagPhotoControl) Get_HardwareAcceleratedThumbnailSupported() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HardwareAcceleratedThumbnailSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 3DCF909D-6D16-409C-BAFE-B9A594C6FDE6
var IID_ILowLagPhotoSequenceControl = syscall.GUID{0x3DCF909D, 0x6D16, 0x409C,
	[8]byte{0xBA, 0xFE, 0xB9, 0xA5, 0x94, 0xC6, 0xFD, 0xE6}}

type ILowLagPhotoSequenceControlInterface interface {
	win32.IInspectableInterface
	Get_Supported() bool
	Get_MaxPastPhotos() uint32
	Get_MaxPhotosPerSecond() float32
	Get_PastPhotoLimit() uint32
	Put_PastPhotoLimit(value uint32)
	Get_PhotosPerSecondLimit() float32
	Put_PhotosPerSecondLimit(value float32)
	GetHighestConcurrentFrameRate(captureProperties *IMediaEncodingProperties) *IMediaRatio
	GetCurrentFrameRate() *IMediaRatio
	Get_ThumbnailEnabled() bool
	Put_ThumbnailEnabled(value bool)
	Get_ThumbnailFormat() MediaThumbnailFormat
	Put_ThumbnailFormat(value MediaThumbnailFormat)
	Get_DesiredThumbnailSize() uint32
	Put_DesiredThumbnailSize(value uint32)
	Get_HardwareAcceleratedThumbnailSupported() uint32
}

type ILowLagPhotoSequenceControlVtbl struct {
	win32.IInspectableVtbl
	Get_Supported                             uintptr
	Get_MaxPastPhotos                         uintptr
	Get_MaxPhotosPerSecond                    uintptr
	Get_PastPhotoLimit                        uintptr
	Put_PastPhotoLimit                        uintptr
	Get_PhotosPerSecondLimit                  uintptr
	Put_PhotosPerSecondLimit                  uintptr
	GetHighestConcurrentFrameRate             uintptr
	GetCurrentFrameRate                       uintptr
	Get_ThumbnailEnabled                      uintptr
	Put_ThumbnailEnabled                      uintptr
	Get_ThumbnailFormat                       uintptr
	Put_ThumbnailFormat                       uintptr
	Get_DesiredThumbnailSize                  uintptr
	Put_DesiredThumbnailSize                  uintptr
	Get_HardwareAcceleratedThumbnailSupported uintptr
}

type ILowLagPhotoSequenceControl struct {
	win32.IInspectable
}

func (this *ILowLagPhotoSequenceControl) Vtbl() *ILowLagPhotoSequenceControlVtbl {
	return (*ILowLagPhotoSequenceControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILowLagPhotoSequenceControl) Get_Supported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Supported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILowLagPhotoSequenceControl) Get_MaxPastPhotos() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxPastPhotos, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILowLagPhotoSequenceControl) Get_MaxPhotosPerSecond() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxPhotosPerSecond, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILowLagPhotoSequenceControl) Get_PastPhotoLimit() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PastPhotoLimit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILowLagPhotoSequenceControl) Put_PastPhotoLimit(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PastPhotoLimit, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ILowLagPhotoSequenceControl) Get_PhotosPerSecondLimit() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhotosPerSecondLimit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILowLagPhotoSequenceControl) Put_PhotosPerSecondLimit(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PhotosPerSecondLimit, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ILowLagPhotoSequenceControl) GetHighestConcurrentFrameRate(captureProperties *IMediaEncodingProperties) *IMediaRatio {
	var _result *IMediaRatio
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetHighestConcurrentFrameRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(captureProperties)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILowLagPhotoSequenceControl) GetCurrentFrameRate() *IMediaRatio {
	var _result *IMediaRatio
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentFrameRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILowLagPhotoSequenceControl) Get_ThumbnailEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ThumbnailEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILowLagPhotoSequenceControl) Put_ThumbnailEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ThumbnailEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ILowLagPhotoSequenceControl) Get_ThumbnailFormat() MediaThumbnailFormat {
	var _result MediaThumbnailFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ThumbnailFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILowLagPhotoSequenceControl) Put_ThumbnailFormat(value MediaThumbnailFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ThumbnailFormat, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ILowLagPhotoSequenceControl) Get_DesiredThumbnailSize() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesiredThumbnailSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILowLagPhotoSequenceControl) Put_DesiredThumbnailSize(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DesiredThumbnailSize, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ILowLagPhotoSequenceControl) Get_HardwareAcceleratedThumbnailSupported() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HardwareAcceleratedThumbnailSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// EFA8DFA9-6F75-4863-BA0B-583F3036B4DE
var IID_IMediaDeviceControl = syscall.GUID{0xEFA8DFA9, 0x6F75, 0x4863,
	[8]byte{0xBA, 0x0B, 0x58, 0x3F, 0x30, 0x36, 0xB4, 0xDE}}

type IMediaDeviceControlInterface interface {
	win32.IInspectableInterface
	Get_Capabilities() *IMediaDeviceControlCapabilities
	TryGetValue(value float64) bool
	TrySetValue(value float64) bool
	TryGetAuto(value bool) bool
	TrySetAuto(value bool) bool
}

type IMediaDeviceControlVtbl struct {
	win32.IInspectableVtbl
	Get_Capabilities uintptr
	TryGetValue      uintptr
	TrySetValue      uintptr
	TryGetAuto       uintptr
	TrySetAuto       uintptr
}

type IMediaDeviceControl struct {
	win32.IInspectable
}

func (this *IMediaDeviceControl) Vtbl() *IMediaDeviceControlVtbl {
	return (*IMediaDeviceControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaDeviceControl) Get_Capabilities() *IMediaDeviceControlCapabilities {
	var _result *IMediaDeviceControlCapabilities
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Capabilities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaDeviceControl) TryGetValue(value float64) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetValue, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaDeviceControl) TrySetValue(value float64) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TrySetValue, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaDeviceControl) TryGetAuto(value bool) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetAuto, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaDeviceControl) TrySetAuto(value bool) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TrySetAuto, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 23005816-EB85-43E2-B92B-8240D5EE70EC
var IID_IMediaDeviceControlCapabilities = syscall.GUID{0x23005816, 0xEB85, 0x43E2,
	[8]byte{0xB9, 0x2B, 0x82, 0x40, 0xD5, 0xEE, 0x70, 0xEC}}

type IMediaDeviceControlCapabilitiesInterface interface {
	win32.IInspectableInterface
	Get_Supported() bool
	Get_Min() float64
	Get_Max() float64
	Get_Step() float64
	Get_Default() float64
	Get_AutoModeSupported() bool
}

type IMediaDeviceControlCapabilitiesVtbl struct {
	win32.IInspectableVtbl
	Get_Supported         uintptr
	Get_Min               uintptr
	Get_Max               uintptr
	Get_Step              uintptr
	Get_Default           uintptr
	Get_AutoModeSupported uintptr
}

type IMediaDeviceControlCapabilities struct {
	win32.IInspectable
}

func (this *IMediaDeviceControlCapabilities) Vtbl() *IMediaDeviceControlCapabilitiesVtbl {
	return (*IMediaDeviceControlCapabilitiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaDeviceControlCapabilities) Get_Supported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Supported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaDeviceControlCapabilities) Get_Min() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Min, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaDeviceControlCapabilities) Get_Max() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Max, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaDeviceControlCapabilities) Get_Step() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Step, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaDeviceControlCapabilities) Get_Default() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Default, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaDeviceControlCapabilities) Get_AutoModeSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AutoModeSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F6F8F5CE-209A-48FB-86FC-D44578F317E6
var IID_IMediaDeviceController = syscall.GUID{0xF6F8F5CE, 0x209A, 0x48FB,
	[8]byte{0x86, 0xFC, 0xD4, 0x45, 0x78, 0xF3, 0x17, 0xE6}}

type IMediaDeviceControllerInterface interface {
	win32.IInspectableInterface
	GetAvailableMediaStreamProperties(mediaStreamType MediaStreamType) *IVectorView[*IMediaEncodingProperties]
	GetMediaStreamProperties(mediaStreamType MediaStreamType) *IMediaEncodingProperties
	SetMediaStreamPropertiesAsync(mediaStreamType MediaStreamType, mediaEncodingProperties *IMediaEncodingProperties) *IAsyncAction
}

type IMediaDeviceControllerVtbl struct {
	win32.IInspectableVtbl
	GetAvailableMediaStreamProperties uintptr
	GetMediaStreamProperties          uintptr
	SetMediaStreamPropertiesAsync     uintptr
}

type IMediaDeviceController struct {
	win32.IInspectable
}

func (this *IMediaDeviceController) Vtbl() *IMediaDeviceControllerVtbl {
	return (*IMediaDeviceControllerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaDeviceController) GetAvailableMediaStreamProperties(mediaStreamType MediaStreamType) *IVectorView[*IMediaEncodingProperties] {
	var _result *IVectorView[*IMediaEncodingProperties]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAvailableMediaStreamProperties, uintptr(unsafe.Pointer(this)), uintptr(mediaStreamType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaDeviceController) GetMediaStreamProperties(mediaStreamType MediaStreamType) *IMediaEncodingProperties {
	var _result *IMediaEncodingProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetMediaStreamProperties, uintptr(unsafe.Pointer(this)), uintptr(mediaStreamType), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMediaDeviceController) SetMediaStreamPropertiesAsync(mediaStreamType MediaStreamType, mediaEncodingProperties *IMediaEncodingProperties) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetMediaStreamPropertiesAsync, uintptr(unsafe.Pointer(this)), uintptr(mediaStreamType), uintptr(unsafe.Pointer(mediaEncodingProperties)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AA2D9A40-909F-4BBA-BF8B-0C0D296F14F0
var IID_IMediaDeviceStatics = syscall.GUID{0xAA2D9A40, 0x909F, 0x4BBA,
	[8]byte{0xBF, 0x8B, 0x0C, 0x0D, 0x29, 0x6F, 0x14, 0xF0}}

type IMediaDeviceStaticsInterface interface {
	win32.IInspectableInterface
	GetAudioCaptureSelector() string
	GetAudioRenderSelector() string
	GetVideoCaptureSelector() string
	GetDefaultAudioCaptureId(role AudioDeviceRole) string
	GetDefaultAudioRenderId(role AudioDeviceRole) string
	Add_DefaultAudioCaptureDeviceChanged(handler TypedEventHandler[interface{}, *IDefaultAudioDeviceChangedEventArgs]) EventRegistrationToken
	Remove_DefaultAudioCaptureDeviceChanged(cookie EventRegistrationToken)
	Add_DefaultAudioRenderDeviceChanged(handler TypedEventHandler[interface{}, *IDefaultAudioDeviceChangedEventArgs]) EventRegistrationToken
	Remove_DefaultAudioRenderDeviceChanged(cookie EventRegistrationToken)
}

type IMediaDeviceStaticsVtbl struct {
	win32.IInspectableVtbl
	GetAudioCaptureSelector                 uintptr
	GetAudioRenderSelector                  uintptr
	GetVideoCaptureSelector                 uintptr
	GetDefaultAudioCaptureId                uintptr
	GetDefaultAudioRenderId                 uintptr
	Add_DefaultAudioCaptureDeviceChanged    uintptr
	Remove_DefaultAudioCaptureDeviceChanged uintptr
	Add_DefaultAudioRenderDeviceChanged     uintptr
	Remove_DefaultAudioRenderDeviceChanged  uintptr
}

type IMediaDeviceStatics struct {
	win32.IInspectable
}

func (this *IMediaDeviceStatics) Vtbl() *IMediaDeviceStaticsVtbl {
	return (*IMediaDeviceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMediaDeviceStatics) GetAudioCaptureSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAudioCaptureSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaDeviceStatics) GetAudioRenderSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAudioRenderSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaDeviceStatics) GetVideoCaptureSelector() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetVideoCaptureSelector, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaDeviceStatics) GetDefaultAudioCaptureId(role AudioDeviceRole) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefaultAudioCaptureId, uintptr(unsafe.Pointer(this)), uintptr(role), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaDeviceStatics) GetDefaultAudioRenderId(role AudioDeviceRole) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefaultAudioRenderId, uintptr(unsafe.Pointer(this)), uintptr(role), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMediaDeviceStatics) Add_DefaultAudioCaptureDeviceChanged(handler TypedEventHandler[interface{}, *IDefaultAudioDeviceChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DefaultAudioCaptureDeviceChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaDeviceStatics) Remove_DefaultAudioCaptureDeviceChanged(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DefaultAudioCaptureDeviceChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IMediaDeviceStatics) Add_DefaultAudioRenderDeviceChanged(handler TypedEventHandler[interface{}, *IDefaultAudioDeviceChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_DefaultAudioRenderDeviceChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMediaDeviceStatics) Remove_DefaultAudioRenderDeviceChanged(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_DefaultAudioRenderDeviceChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

// 520D1EB4-1374-4C7D-B1E4-39DCDF3EAE4E
var IID_IModuleCommandResult = syscall.GUID{0x520D1EB4, 0x1374, 0x4C7D,
	[8]byte{0xB1, 0xE4, 0x39, 0xDC, 0xDF, 0x3E, 0xAE, 0x4E}}

type IModuleCommandResultInterface interface {
	win32.IInspectableInterface
	Get_Status() SendCommandStatus
	Get_Result() *IBuffer
}

type IModuleCommandResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
	Get_Result uintptr
}

type IModuleCommandResult struct {
	win32.IInspectable
}

func (this *IModuleCommandResult) Vtbl() *IModuleCommandResultVtbl {
	return (*IModuleCommandResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IModuleCommandResult) Get_Status() SendCommandStatus {
	var _result SendCommandStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IModuleCommandResult) Get_Result() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Result, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BFAD9C1D-00BC-423B-8EB2-A0178CA94247
var IID_IOpticalImageStabilizationControl = syscall.GUID{0xBFAD9C1D, 0x00BC, 0x423B,
	[8]byte{0x8E, 0xB2, 0xA0, 0x17, 0x8C, 0xA9, 0x42, 0x47}}

type IOpticalImageStabilizationControlInterface interface {
	win32.IInspectableInterface
	Get_Supported() bool
	Get_SupportedModes() *IVectorView[OpticalImageStabilizationMode]
	Get_Mode() OpticalImageStabilizationMode
	Put_Mode(value OpticalImageStabilizationMode)
}

type IOpticalImageStabilizationControlVtbl struct {
	win32.IInspectableVtbl
	Get_Supported      uintptr
	Get_SupportedModes uintptr
	Get_Mode           uintptr
	Put_Mode           uintptr
}

type IOpticalImageStabilizationControl struct {
	win32.IInspectable
}

func (this *IOpticalImageStabilizationControl) Vtbl() *IOpticalImageStabilizationControlVtbl {
	return (*IOpticalImageStabilizationControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IOpticalImageStabilizationControl) Get_Supported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Supported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IOpticalImageStabilizationControl) Get_SupportedModes() *IVectorView[OpticalImageStabilizationMode] {
	var _result *IVectorView[OpticalImageStabilizationMode]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedModes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IOpticalImageStabilizationControl) Get_Mode() OpticalImageStabilizationMode {
	var _result OpticalImageStabilizationMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IOpticalImageStabilizationControl) Put_Mode(value OpticalImageStabilizationMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Mode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 33323223-6247-5419-A5A4-3D808645D917
var IID_IPanelBasedOptimizationControl = syscall.GUID{0x33323223, 0x6247, 0x5419,
	[8]byte{0xA5, 0xA4, 0x3D, 0x80, 0x86, 0x45, 0xD9, 0x17}}

type IPanelBasedOptimizationControlInterface interface {
	win32.IInspectableInterface
	Get_IsSupported() bool
	Get_Panel() Panel
	Put_Panel(value Panel)
}

type IPanelBasedOptimizationControlVtbl struct {
	win32.IInspectableVtbl
	Get_IsSupported uintptr
	Get_Panel       uintptr
	Put_Panel       uintptr
}

type IPanelBasedOptimizationControl struct {
	win32.IInspectable
}

func (this *IPanelBasedOptimizationControl) Vtbl() *IPanelBasedOptimizationControlVtbl {
	return (*IPanelBasedOptimizationControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPanelBasedOptimizationControl) Get_IsSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPanelBasedOptimizationControl) Get_Panel() Panel {
	var _result Panel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Panel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPanelBasedOptimizationControl) Put_Panel(value Panel) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Panel, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// C8F3F363-FF5E-4582-A9A8-0550F85A4A76
var IID_IPhotoConfirmationControl = syscall.GUID{0xC8F3F363, 0xFF5E, 0x4582,
	[8]byte{0xA9, 0xA8, 0x05, 0x50, 0xF8, 0x5A, 0x4A, 0x76}}

type IPhotoConfirmationControlInterface interface {
	win32.IInspectableInterface
	Get_Supported() bool
	Get_Enabled() bool
	Put_Enabled(value bool)
	Get_PixelFormat() MediaPixelFormat
	Put_PixelFormat(format MediaPixelFormat)
}

type IPhotoConfirmationControlVtbl struct {
	win32.IInspectableVtbl
	Get_Supported   uintptr
	Get_Enabled     uintptr
	Put_Enabled     uintptr
	Get_PixelFormat uintptr
	Put_PixelFormat uintptr
}

type IPhotoConfirmationControl struct {
	win32.IInspectable
}

func (this *IPhotoConfirmationControl) Vtbl() *IPhotoConfirmationControlVtbl {
	return (*IPhotoConfirmationControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPhotoConfirmationControl) Get_Supported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Supported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoConfirmationControl) Get_Enabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Enabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoConfirmationControl) Put_Enabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Enabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPhotoConfirmationControl) Get_PixelFormat() MediaPixelFormat {
	var _result MediaPixelFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PixelFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPhotoConfirmationControl) Put_PixelFormat(format MediaPixelFormat) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PixelFormat, uintptr(unsafe.Pointer(this)), uintptr(format))
	_ = _hr
}

// 7EB55209-76AB-4C31-B40E-4B58379D580C
var IID_IRedialRequestedEventArgs = syscall.GUID{0x7EB55209, 0x76AB, 0x4C31,
	[8]byte{0xB4, 0x0E, 0x4B, 0x58, 0x37, 0x9D, 0x58, 0x0C}}

type IRedialRequestedEventArgsInterface interface {
	win32.IInspectableInterface
	Handled()
}

type IRedialRequestedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Handled uintptr
}

type IRedialRequestedEventArgs struct {
	win32.IInspectable
}

func (this *IRedialRequestedEventArgs) Vtbl() *IRedialRequestedEventArgsVtbl {
	return (*IRedialRequestedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRedialRequestedEventArgs) Handled() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Handled, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// E5ECC834-CE66-4E05-A78F-CF391A5EC2D1
var IID_IRegionOfInterest = syscall.GUID{0xE5ECC834, 0xCE66, 0x4E05,
	[8]byte{0xA7, 0x8F, 0xCF, 0x39, 0x1A, 0x5E, 0xC2, 0xD1}}

type IRegionOfInterestInterface interface {
	win32.IInspectableInterface
	Get_AutoFocusEnabled() bool
	Put_AutoFocusEnabled(value bool)
	Get_AutoWhiteBalanceEnabled() bool
	Put_AutoWhiteBalanceEnabled(value bool)
	Get_AutoExposureEnabled() bool
	Put_AutoExposureEnabled(value bool)
	Get_Bounds() Rect
	Put_Bounds(value Rect)
}

type IRegionOfInterestVtbl struct {
	win32.IInspectableVtbl
	Get_AutoFocusEnabled        uintptr
	Put_AutoFocusEnabled        uintptr
	Get_AutoWhiteBalanceEnabled uintptr
	Put_AutoWhiteBalanceEnabled uintptr
	Get_AutoExposureEnabled     uintptr
	Put_AutoExposureEnabled     uintptr
	Get_Bounds                  uintptr
	Put_Bounds                  uintptr
}

type IRegionOfInterest struct {
	win32.IInspectable
}

func (this *IRegionOfInterest) Vtbl() *IRegionOfInterestVtbl {
	return (*IRegionOfInterestVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRegionOfInterest) Get_AutoFocusEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AutoFocusEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRegionOfInterest) Put_AutoFocusEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AutoFocusEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IRegionOfInterest) Get_AutoWhiteBalanceEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AutoWhiteBalanceEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRegionOfInterest) Put_AutoWhiteBalanceEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AutoWhiteBalanceEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IRegionOfInterest) Get_AutoExposureEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AutoExposureEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRegionOfInterest) Put_AutoExposureEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AutoExposureEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IRegionOfInterest) Get_Bounds() Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bounds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRegionOfInterest) Put_Bounds(value Rect) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Bounds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

// 19FE2A91-73AA-4D51-8A9D-56CCF7DB7F54
var IID_IRegionOfInterest2 = syscall.GUID{0x19FE2A91, 0x73AA, 0x4D51,
	[8]byte{0x8A, 0x9D, 0x56, 0xCC, 0xF7, 0xDB, 0x7F, 0x54}}

type IRegionOfInterest2Interface interface {
	win32.IInspectableInterface
	Get_Type() RegionOfInterestType
	Put_Type(value RegionOfInterestType)
	Get_BoundsNormalized() bool
	Put_BoundsNormalized(value bool)
	Get_Weight() uint32
	Put_Weight(value uint32)
}

type IRegionOfInterest2Vtbl struct {
	win32.IInspectableVtbl
	Get_Type             uintptr
	Put_Type             uintptr
	Get_BoundsNormalized uintptr
	Put_BoundsNormalized uintptr
	Get_Weight           uintptr
	Put_Weight           uintptr
}

type IRegionOfInterest2 struct {
	win32.IInspectable
}

func (this *IRegionOfInterest2) Vtbl() *IRegionOfInterest2Vtbl {
	return (*IRegionOfInterest2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRegionOfInterest2) Get_Type() RegionOfInterestType {
	var _result RegionOfInterestType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Type, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRegionOfInterest2) Put_Type(value RegionOfInterestType) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Type, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IRegionOfInterest2) Get_BoundsNormalized() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BoundsNormalized, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRegionOfInterest2) Put_BoundsNormalized(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BoundsNormalized, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IRegionOfInterest2) Get_Weight() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Weight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRegionOfInterest2) Put_Weight(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Weight, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// C323F527-AB0B-4558-8B5B-DF5693DB0378
var IID_IRegionsOfInterestControl = syscall.GUID{0xC323F527, 0xAB0B, 0x4558,
	[8]byte{0x8B, 0x5B, 0xDF, 0x56, 0x93, 0xDB, 0x03, 0x78}}

type IRegionsOfInterestControlInterface interface {
	win32.IInspectableInterface
	Get_MaxRegions() uint32
	SetRegionsAsync(regions *IIterable[*IRegionOfInterest]) *IAsyncAction
	SetRegionsWithLockAsync(regions *IIterable[*IRegionOfInterest], lockValues bool) *IAsyncAction
	ClearRegionsAsync() *IAsyncAction
	Get_AutoFocusSupported() bool
	Get_AutoWhiteBalanceSupported() bool
	Get_AutoExposureSupported() bool
}

type IRegionsOfInterestControlVtbl struct {
	win32.IInspectableVtbl
	Get_MaxRegions                uintptr
	SetRegionsAsync               uintptr
	SetRegionsWithLockAsync       uintptr
	ClearRegionsAsync             uintptr
	Get_AutoFocusSupported        uintptr
	Get_AutoWhiteBalanceSupported uintptr
	Get_AutoExposureSupported     uintptr
}

type IRegionsOfInterestControl struct {
	win32.IInspectable
}

func (this *IRegionsOfInterestControl) Vtbl() *IRegionsOfInterestControlVtbl {
	return (*IRegionsOfInterestControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRegionsOfInterestControl) Get_MaxRegions() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxRegions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRegionsOfInterestControl) SetRegionsAsync(regions *IIterable[*IRegionOfInterest]) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetRegionsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(regions)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRegionsOfInterestControl) SetRegionsWithLockAsync(regions *IIterable[*IRegionOfInterest], lockValues bool) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetRegionsWithLockAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(regions)), uintptr(*(*byte)(unsafe.Pointer(&lockValues))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRegionsOfInterestControl) ClearRegionsAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearRegionsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRegionsOfInterestControl) Get_AutoFocusSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AutoFocusSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRegionsOfInterestControl) Get_AutoWhiteBalanceSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AutoWhiteBalanceSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRegionsOfInterestControl) Get_AutoExposureSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AutoExposureSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D48E5AF7-8D59-4854-8C62-12C70BA89B7C
var IID_ISceneModeControl = syscall.GUID{0xD48E5AF7, 0x8D59, 0x4854,
	[8]byte{0x8C, 0x62, 0x12, 0xC7, 0x0B, 0xA8, 0x9B, 0x7C}}

type ISceneModeControlInterface interface {
	win32.IInspectableInterface
	Get_SupportedModes() *IVectorView[CaptureSceneMode]
	Get_Value() CaptureSceneMode
	SetValueAsync(sceneMode CaptureSceneMode) *IAsyncAction
}

type ISceneModeControlVtbl struct {
	win32.IInspectableVtbl
	Get_SupportedModes uintptr
	Get_Value          uintptr
	SetValueAsync      uintptr
}

type ISceneModeControl struct {
	win32.IInspectable
}

func (this *ISceneModeControl) Vtbl() *ISceneModeControlVtbl {
	return (*ISceneModeControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISceneModeControl) Get_SupportedModes() *IVectorView[CaptureSceneMode] {
	var _result *IVectorView[CaptureSceneMode]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedModes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISceneModeControl) Get_Value() CaptureSceneMode {
	var _result CaptureSceneMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISceneModeControl) SetValueAsync(sceneMode CaptureSceneMode) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetValueAsync, uintptr(unsafe.Pointer(this)), uintptr(sceneMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A6053665-8250-416C-919A-724296AFA306
var IID_ITorchControl = syscall.GUID{0xA6053665, 0x8250, 0x416C,
	[8]byte{0x91, 0x9A, 0x72, 0x42, 0x96, 0xAF, 0xA3, 0x06}}

type ITorchControlInterface interface {
	win32.IInspectableInterface
	Get_Supported() bool
	Get_PowerSupported() bool
	Get_Enabled() bool
	Put_Enabled(value bool)
	Get_PowerPercent() float32
	Put_PowerPercent(value float32)
}

type ITorchControlVtbl struct {
	win32.IInspectableVtbl
	Get_Supported      uintptr
	Get_PowerSupported uintptr
	Get_Enabled        uintptr
	Put_Enabled        uintptr
	Get_PowerPercent   uintptr
	Put_PowerPercent   uintptr
}

type ITorchControl struct {
	win32.IInspectable
}

func (this *ITorchControl) Vtbl() *ITorchControlVtbl {
	return (*ITorchControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITorchControl) Get_Supported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Supported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITorchControl) Get_PowerSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PowerSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITorchControl) Get_Enabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Enabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITorchControl) Put_Enabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Enabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ITorchControl) Get_PowerPercent() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PowerPercent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITorchControl) Put_PowerPercent(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PowerPercent, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 99555575-2E2E-40B8-B6C7-F82D10013210
var IID_IVideoDeviceController = syscall.GUID{0x99555575, 0x2E2E, 0x40B8,
	[8]byte{0xB6, 0xC7, 0xF8, 0x2D, 0x10, 0x01, 0x32, 0x10}}

type IVideoDeviceControllerInterface interface {
	win32.IInspectableInterface
	Get_Brightness() *IMediaDeviceControl
	Get_Contrast() *IMediaDeviceControl
	Get_Hue() *IMediaDeviceControl
	Get_WhiteBalance() *IMediaDeviceControl
	Get_BacklightCompensation() *IMediaDeviceControl
	Get_Pan() *IMediaDeviceControl
	Get_Tilt() *IMediaDeviceControl
	Get_Zoom() *IMediaDeviceControl
	Get_Roll() *IMediaDeviceControl
	Get_Exposure() *IMediaDeviceControl
	Get_Focus() *IMediaDeviceControl
	TrySetPowerlineFrequency(value PowerlineFrequency) bool
	TryGetPowerlineFrequency(value PowerlineFrequency) bool
}

type IVideoDeviceControllerVtbl struct {
	win32.IInspectableVtbl
	Get_Brightness            uintptr
	Get_Contrast              uintptr
	Get_Hue                   uintptr
	Get_WhiteBalance          uintptr
	Get_BacklightCompensation uintptr
	Get_Pan                   uintptr
	Get_Tilt                  uintptr
	Get_Zoom                  uintptr
	Get_Roll                  uintptr
	Get_Exposure              uintptr
	Get_Focus                 uintptr
	TrySetPowerlineFrequency  uintptr
	TryGetPowerlineFrequency  uintptr
}

type IVideoDeviceController struct {
	win32.IInspectable
}

func (this *IVideoDeviceController) Vtbl() *IVideoDeviceControllerVtbl {
	return (*IVideoDeviceControllerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoDeviceController) Get_Brightness() *IMediaDeviceControl {
	var _result *IMediaDeviceControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Brightness, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoDeviceController) Get_Contrast() *IMediaDeviceControl {
	var _result *IMediaDeviceControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Contrast, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoDeviceController) Get_Hue() *IMediaDeviceControl {
	var _result *IMediaDeviceControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Hue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoDeviceController) Get_WhiteBalance() *IMediaDeviceControl {
	var _result *IMediaDeviceControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WhiteBalance, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoDeviceController) Get_BacklightCompensation() *IMediaDeviceControl {
	var _result *IMediaDeviceControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BacklightCompensation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoDeviceController) Get_Pan() *IMediaDeviceControl {
	var _result *IMediaDeviceControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Pan, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoDeviceController) Get_Tilt() *IMediaDeviceControl {
	var _result *IMediaDeviceControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Tilt, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoDeviceController) Get_Zoom() *IMediaDeviceControl {
	var _result *IMediaDeviceControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Zoom, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoDeviceController) Get_Roll() *IMediaDeviceControl {
	var _result *IMediaDeviceControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Roll, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoDeviceController) Get_Exposure() *IMediaDeviceControl {
	var _result *IMediaDeviceControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Exposure, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoDeviceController) Get_Focus() *IMediaDeviceControl {
	var _result *IMediaDeviceControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Focus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoDeviceController) TrySetPowerlineFrequency(value PowerlineFrequency) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TrySetPowerlineFrequency, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoDeviceController) TryGetPowerlineFrequency(value PowerlineFrequency) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetPowerlineFrequency, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C5D88395-6ED5-4790-8B5D-0EF13935D0F8
var IID_IVideoDeviceControllerGetDevicePropertyResult = syscall.GUID{0xC5D88395, 0x6ED5, 0x4790,
	[8]byte{0x8B, 0x5D, 0x0E, 0xF1, 0x39, 0x35, 0xD0, 0xF8}}

type IVideoDeviceControllerGetDevicePropertyResultInterface interface {
	win32.IInspectableInterface
	Get_Status() VideoDeviceControllerGetDevicePropertyStatus
	Get_Value() interface{}
}

type IVideoDeviceControllerGetDevicePropertyResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
	Get_Value  uintptr
}

type IVideoDeviceControllerGetDevicePropertyResult struct {
	win32.IInspectable
}

func (this *IVideoDeviceControllerGetDevicePropertyResult) Vtbl() *IVideoDeviceControllerGetDevicePropertyResultVtbl {
	return (*IVideoDeviceControllerGetDevicePropertyResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoDeviceControllerGetDevicePropertyResult) Get_Status() VideoDeviceControllerGetDevicePropertyStatus {
	var _result VideoDeviceControllerGetDevicePropertyStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoDeviceControllerGetDevicePropertyResult) Get_Value() interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7AB34735-3E2A-4A32-BAFF-4358C4FBDD57
var IID_IVideoTemporalDenoisingControl = syscall.GUID{0x7AB34735, 0x3E2A, 0x4A32,
	[8]byte{0xBA, 0xFF, 0x43, 0x58, 0xC4, 0xFB, 0xDD, 0x57}}

type IVideoTemporalDenoisingControlInterface interface {
	win32.IInspectableInterface
	Get_Supported() bool
	Get_SupportedModes() *IVectorView[VideoTemporalDenoisingMode]
	Get_Mode() VideoTemporalDenoisingMode
	Put_Mode(value VideoTemporalDenoisingMode)
}

type IVideoTemporalDenoisingControlVtbl struct {
	win32.IInspectableVtbl
	Get_Supported      uintptr
	Get_SupportedModes uintptr
	Get_Mode           uintptr
	Put_Mode           uintptr
}

type IVideoTemporalDenoisingControl struct {
	win32.IInspectable
}

func (this *IVideoTemporalDenoisingControl) Vtbl() *IVideoTemporalDenoisingControlVtbl {
	return (*IVideoTemporalDenoisingControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoTemporalDenoisingControl) Get_Supported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Supported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoTemporalDenoisingControl) Get_SupportedModes() *IVectorView[VideoTemporalDenoisingMode] {
	var _result *IVectorView[VideoTemporalDenoisingMode]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedModes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoTemporalDenoisingControl) Get_Mode() VideoTemporalDenoisingMode {
	var _result VideoTemporalDenoisingMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoTemporalDenoisingControl) Put_Mode(value VideoTemporalDenoisingMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Mode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 781F047E-7162-49C8-A8F9-9481C565363E
var IID_IWhiteBalanceControl = syscall.GUID{0x781F047E, 0x7162, 0x49C8,
	[8]byte{0xA8, 0xF9, 0x94, 0x81, 0xC5, 0x65, 0x36, 0x3E}}

type IWhiteBalanceControlInterface interface {
	win32.IInspectableInterface
	Get_Supported() bool
	Get_Preset() ColorTemperaturePreset
	SetPresetAsync(preset ColorTemperaturePreset) *IAsyncAction
	Get_Min() uint32
	Get_Max() uint32
	Get_Step() uint32
	Get_Value() uint32
	SetValueAsync(temperature uint32) *IAsyncAction
}

type IWhiteBalanceControlVtbl struct {
	win32.IInspectableVtbl
	Get_Supported  uintptr
	Get_Preset     uintptr
	SetPresetAsync uintptr
	Get_Min        uintptr
	Get_Max        uintptr
	Get_Step       uintptr
	Get_Value      uintptr
	SetValueAsync  uintptr
}

type IWhiteBalanceControl struct {
	win32.IInspectable
}

func (this *IWhiteBalanceControl) Vtbl() *IWhiteBalanceControlVtbl {
	return (*IWhiteBalanceControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWhiteBalanceControl) Get_Supported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Supported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWhiteBalanceControl) Get_Preset() ColorTemperaturePreset {
	var _result ColorTemperaturePreset
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Preset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWhiteBalanceControl) SetPresetAsync(preset ColorTemperaturePreset) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetPresetAsync, uintptr(unsafe.Pointer(this)), uintptr(preset), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IWhiteBalanceControl) Get_Min() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Min, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWhiteBalanceControl) Get_Max() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Max, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWhiteBalanceControl) Get_Step() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Step, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWhiteBalanceControl) Get_Value() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IWhiteBalanceControl) SetValueAsync(temperature uint32) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetValueAsync, uintptr(unsafe.Pointer(this)), uintptr(temperature), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3A1E0B12-32DA-4C17-BFD7-8D0C73C8F5A5
var IID_IZoomControl = syscall.GUID{0x3A1E0B12, 0x32DA, 0x4C17,
	[8]byte{0xBF, 0xD7, 0x8D, 0x0C, 0x73, 0xC8, 0xF5, 0xA5}}

type IZoomControlInterface interface {
	win32.IInspectableInterface
	Get_Supported() bool
	Get_Min() float32
	Get_Max() float32
	Get_Step() float32
	Get_Value() float32
	Put_Value(value float32)
}

type IZoomControlVtbl struct {
	win32.IInspectableVtbl
	Get_Supported uintptr
	Get_Min       uintptr
	Get_Max       uintptr
	Get_Step      uintptr
	Get_Value     uintptr
	Put_Value     uintptr
}

type IZoomControl struct {
	win32.IInspectable
}

func (this *IZoomControl) Vtbl() *IZoomControlVtbl {
	return (*IZoomControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IZoomControl) Get_Supported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Supported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IZoomControl) Get_Min() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Min, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IZoomControl) Get_Max() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Max, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IZoomControl) Get_Step() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Step, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IZoomControl) Get_Value() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IZoomControl) Put_Value(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Value, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 69843DB0-2E99-4641-8529-184F319D1671
var IID_IZoomControl2 = syscall.GUID{0x69843DB0, 0x2E99, 0x4641,
	[8]byte{0x85, 0x29, 0x18, 0x4F, 0x31, 0x9D, 0x16, 0x71}}

type IZoomControl2Interface interface {
	win32.IInspectableInterface
	Get_SupportedModes() *IVectorView[ZoomTransitionMode]
	Get_Mode() ZoomTransitionMode
	Configure(settings *IZoomSettings)
}

type IZoomControl2Vtbl struct {
	win32.IInspectableVtbl
	Get_SupportedModes uintptr
	Get_Mode           uintptr
	Configure          uintptr
}

type IZoomControl2 struct {
	win32.IInspectable
}

func (this *IZoomControl2) Vtbl() *IZoomControl2Vtbl {
	return (*IZoomControl2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IZoomControl2) Get_SupportedModes() *IVectorView[ZoomTransitionMode] {
	var _result *IVectorView[ZoomTransitionMode]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedModes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IZoomControl2) Get_Mode() ZoomTransitionMode {
	var _result ZoomTransitionMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IZoomControl2) Configure(settings *IZoomSettings) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Configure, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(settings)))
	_ = _hr
}

// 6AD66B24-14B4-4BFD-B18F-88FE24463B52
var IID_IZoomSettings = syscall.GUID{0x6AD66B24, 0x14B4, 0x4BFD,
	[8]byte{0xB1, 0x8F, 0x88, 0xFE, 0x24, 0x46, 0x3B, 0x52}}

type IZoomSettingsInterface interface {
	win32.IInspectableInterface
	Get_Mode() ZoomTransitionMode
	Put_Mode(value ZoomTransitionMode)
	Get_Value() float32
	Put_Value(value float32)
}

type IZoomSettingsVtbl struct {
	win32.IInspectableVtbl
	Get_Mode  uintptr
	Put_Mode  uintptr
	Get_Value uintptr
	Put_Value uintptr
}

type IZoomSettings struct {
	win32.IInspectable
}

func (this *IZoomSettings) Vtbl() *IZoomSettingsVtbl {
	return (*IZoomSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IZoomSettings) Get_Mode() ZoomTransitionMode {
	var _result ZoomTransitionMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IZoomSettings) Put_Mode(value ZoomTransitionMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Mode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IZoomSettings) Get_Value() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IZoomSettings) Put_Value(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Value, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// classes

type AdvancedPhotoCaptureSettings struct {
	RtClass
	*IAdvancedPhotoCaptureSettings
}

func NewAdvancedPhotoCaptureSettings() *AdvancedPhotoCaptureSettings {
	hs := NewHStr("Windows.Media.Devices.AdvancedPhotoCaptureSettings")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &AdvancedPhotoCaptureSettings{
		RtClass:                       RtClass{PInspect: p},
		IAdvancedPhotoCaptureSettings: (*IAdvancedPhotoCaptureSettings)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type AdvancedPhotoControl struct {
	RtClass
	*IAdvancedPhotoControl
}

type AudioDeviceController struct {
	RtClass
	*IAudioDeviceController
}

type AudioDeviceModuleNotificationEventArgs struct {
	RtClass
	*IAudioDeviceModuleNotificationEventArgs
}

type AudioDeviceModulesManager struct {
	RtClass
	*IAudioDeviceModulesManager
}

func NewAudioDeviceModulesManager_Create(deviceId string) *AudioDeviceModulesManager {
	hs := NewHStr("Windows.Media.Devices.AudioDeviceModulesManager")
	var pFac *IAudioDeviceModulesManagerFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IAudioDeviceModulesManagerFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IAudioDeviceModulesManager
	p = pFac.Create(deviceId)
	result := &AudioDeviceModulesManager{
		RtClass:                    RtClass{PInspect: &p.IInspectable},
		IAudioDeviceModulesManager: p,
	}
	com.AddToScope(result)
	return result
}

type CallControl struct {
	RtClass
	*ICallControl
}

func NewICallControlStatics() *ICallControlStatics {
	var p *ICallControlStatics
	hs := NewHStr("Windows.Media.Devices.CallControl")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICallControlStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type CameraOcclusionInfo struct {
	RtClass
	*ICameraOcclusionInfo
}

type CameraOcclusionState struct {
	RtClass
	*ICameraOcclusionState
}

type CameraOcclusionStateChangedEventArgs struct {
	RtClass
	*ICameraOcclusionStateChangedEventArgs
}

type ExposureCompensationControl struct {
	RtClass
	*IExposureCompensationControl
}

type ExposureControl struct {
	RtClass
	*IExposureControl
}

type ExposurePriorityVideoControl struct {
	RtClass
	*IExposurePriorityVideoControl
}

type FlashControl struct {
	RtClass
	*IFlashControl
}

type FocusControl struct {
	RtClass
	*IFocusControl
}

type FocusSettings struct {
	RtClass
	*IFocusSettings
}

func NewFocusSettings() *FocusSettings {
	hs := NewHStr("Windows.Media.Devices.FocusSettings")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &FocusSettings{
		RtClass:        RtClass{PInspect: p},
		IFocusSettings: (*IFocusSettings)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type HdrVideoControl struct {
	RtClass
	*IHdrVideoControl
}

type IsoSpeedControl struct {
	RtClass
	*IIsoSpeedControl
}

type LowLagPhotoControl struct {
	RtClass
	*ILowLagPhotoControl
}

type LowLagPhotoSequenceControl struct {
	RtClass
	*ILowLagPhotoSequenceControl
}

type MediaDevice struct {
	RtClass
}

func NewIMediaDeviceStatics() *IMediaDeviceStatics {
	var p *IMediaDeviceStatics
	hs := NewHStr("Windows.Media.Devices.MediaDevice")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IMediaDeviceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type MediaDeviceControl struct {
	RtClass
	*IMediaDeviceControl
}

type MediaDeviceControlCapabilities struct {
	RtClass
	*IMediaDeviceControlCapabilities
}

type ModuleCommandResult struct {
	RtClass
	*IModuleCommandResult
}

type OpticalImageStabilizationControl struct {
	RtClass
	*IOpticalImageStabilizationControl
}

type PanelBasedOptimizationControl struct {
	RtClass
	*IPanelBasedOptimizationControl
}

type PhotoConfirmationControl struct {
	RtClass
	*IPhotoConfirmationControl
}

type RegionOfInterest struct {
	RtClass
	*IRegionOfInterest
}

func NewRegionOfInterest() *RegionOfInterest {
	hs := NewHStr("Windows.Media.Devices.RegionOfInterest")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &RegionOfInterest{
		RtClass:           RtClass{PInspect: p},
		IRegionOfInterest: (*IRegionOfInterest)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type RegionsOfInterestControl struct {
	RtClass
	*IRegionsOfInterestControl
}

type SceneModeControl struct {
	RtClass
	*ISceneModeControl
}

type TorchControl struct {
	RtClass
	*ITorchControl
}

type VideoDeviceController struct {
	RtClass
	*IVideoDeviceController
}

type VideoDeviceControllerGetDevicePropertyResult struct {
	RtClass
	*IVideoDeviceControllerGetDevicePropertyResult
}

type WhiteBalanceControl struct {
	RtClass
	*IWhiteBalanceControl
}

type ZoomControl struct {
	RtClass
	*IZoomControl
}

type ZoomSettings struct {
	RtClass
	*IZoomSettings
}

func NewZoomSettings() *ZoomSettings {
	hs := NewHStr("Windows.Media.Devices.ZoomSettings")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &ZoomSettings{
		RtClass:       RtClass{PInspect: p},
		IZoomSettings: (*IZoomSettings)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}
