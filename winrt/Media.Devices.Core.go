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
type FrameFlashMode int32

const (
	FrameFlashMode_Disable FrameFlashMode = 0
	FrameFlashMode_Enable  FrameFlashMode = 1
	FrameFlashMode_Global  FrameFlashMode = 2
)

// interfaces

// 0AA6ED32-6589-49DA-AFDE-594270CA0AAC
var IID_ICameraIntrinsics = syscall.GUID{0x0AA6ED32, 0x6589, 0x49DA,
	[8]byte{0xAF, 0xDE, 0x59, 0x42, 0x70, 0xCA, 0x0A, 0xAC}}

type ICameraIntrinsicsInterface interface {
	win32.IInspectableInterface
	Get_FocalLength() unsafe.Pointer
	Get_PrincipalPoint() unsafe.Pointer
	Get_RadialDistortion() unsafe.Pointer
	Get_TangentialDistortion() unsafe.Pointer
	Get_ImageWidth() uint32
	Get_ImageHeight() uint32
	ProjectOntoFrame(coordinate unsafe.Pointer) Point
	UnprojectAtUnitDepth(pixelCoordinate Point) unsafe.Pointer
	ProjectManyOntoFrame(coordinates unsafe.Pointer, resultsLength uint32, results *Point)
	UnprojectPixelsAtUnitDepth(pixelCoordinatesLength uint32, pixelCoordinates *Point, results unsafe.Pointer)
}

type ICameraIntrinsicsVtbl struct {
	win32.IInspectableVtbl
	Get_FocalLength            uintptr
	Get_PrincipalPoint         uintptr
	Get_RadialDistortion       uintptr
	Get_TangentialDistortion   uintptr
	Get_ImageWidth             uintptr
	Get_ImageHeight            uintptr
	ProjectOntoFrame           uintptr
	UnprojectAtUnitDepth       uintptr
	ProjectManyOntoFrame       uintptr
	UnprojectPixelsAtUnitDepth uintptr
}

type ICameraIntrinsics struct {
	win32.IInspectable
}

func (this *ICameraIntrinsics) Vtbl() *ICameraIntrinsicsVtbl {
	return (*ICameraIntrinsicsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICameraIntrinsics) Get_FocalLength() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FocalLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICameraIntrinsics) Get_PrincipalPoint() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PrincipalPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICameraIntrinsics) Get_RadialDistortion() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RadialDistortion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICameraIntrinsics) Get_TangentialDistortion() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TangentialDistortion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICameraIntrinsics) Get_ImageWidth() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ImageWidth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICameraIntrinsics) Get_ImageHeight() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ImageHeight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICameraIntrinsics) ProjectOntoFrame(coordinate unsafe.Pointer) Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ProjectOntoFrame, uintptr(unsafe.Pointer(this)), uintptr(coordinate), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICameraIntrinsics) UnprojectAtUnitDepth(pixelCoordinate Point) unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UnprojectAtUnitDepth, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&pixelCoordinate)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICameraIntrinsics) ProjectManyOntoFrame(coordinates unsafe.Pointer, resultsLength uint32, results *Point) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ProjectManyOntoFrame, uintptr(unsafe.Pointer(this)), uintptr(coordinates), uintptr(resultsLength), uintptr(unsafe.Pointer(results)))
	_ = _hr
}

func (this *ICameraIntrinsics) UnprojectPixelsAtUnitDepth(pixelCoordinatesLength uint32, pixelCoordinates *Point, results unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UnprojectPixelsAtUnitDepth, uintptr(unsafe.Pointer(this)), uintptr(pixelCoordinatesLength), uintptr(unsafe.Pointer(pixelCoordinates)), uintptr(results))
	_ = _hr
}

// 0CDAA447-0798-4B4D-839F-C5EC414DB27A
var IID_ICameraIntrinsics2 = syscall.GUID{0x0CDAA447, 0x0798, 0x4B4D,
	[8]byte{0x83, 0x9F, 0xC5, 0xEC, 0x41, 0x4D, 0xB2, 0x7A}}

type ICameraIntrinsics2Interface interface {
	win32.IInspectableInterface
	Get_UndistortedProjectionTransform() unsafe.Pointer
	DistortPoint(input Point) Point
	DistortPoints(inputsLength uint32, inputs *Point, resultsLength uint32, results *Point)
	UndistortPoint(input Point) Point
	UndistortPoints(inputsLength uint32, inputs *Point, resultsLength uint32, results *Point)
}

type ICameraIntrinsics2Vtbl struct {
	win32.IInspectableVtbl
	Get_UndistortedProjectionTransform uintptr
	DistortPoint                       uintptr
	DistortPoints                      uintptr
	UndistortPoint                     uintptr
	UndistortPoints                    uintptr
}

type ICameraIntrinsics2 struct {
	win32.IInspectable
}

func (this *ICameraIntrinsics2) Vtbl() *ICameraIntrinsics2Vtbl {
	return (*ICameraIntrinsics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICameraIntrinsics2) Get_UndistortedProjectionTransform() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UndistortedProjectionTransform, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICameraIntrinsics2) DistortPoint(input Point) Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DistortPoint, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&input)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICameraIntrinsics2) DistortPoints(inputsLength uint32, inputs *Point, resultsLength uint32, results *Point) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DistortPoints, uintptr(unsafe.Pointer(this)), uintptr(inputsLength), uintptr(unsafe.Pointer(inputs)), uintptr(resultsLength), uintptr(unsafe.Pointer(results)))
	_ = _hr
}

func (this *ICameraIntrinsics2) UndistortPoint(input Point) Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UndistortPoint, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&input)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICameraIntrinsics2) UndistortPoints(inputsLength uint32, inputs *Point, resultsLength uint32, results *Point) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UndistortPoints, uintptr(unsafe.Pointer(this)), uintptr(inputsLength), uintptr(unsafe.Pointer(inputs)), uintptr(resultsLength), uintptr(unsafe.Pointer(results)))
	_ = _hr
}

// C0DDC486-2132-4A34-A659-9BFE2A055712
var IID_ICameraIntrinsicsFactory = syscall.GUID{0xC0DDC486, 0x2132, 0x4A34,
	[8]byte{0xA6, 0x59, 0x9B, 0xFE, 0x2A, 0x05, 0x57, 0x12}}

type ICameraIntrinsicsFactoryInterface interface {
	win32.IInspectableInterface
	Create(focalLength unsafe.Pointer, principalPoint unsafe.Pointer, radialDistortion unsafe.Pointer, tangentialDistortion unsafe.Pointer, imageWidth uint32, imageHeight uint32) *ICameraIntrinsics
}

type ICameraIntrinsicsFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ICameraIntrinsicsFactory struct {
	win32.IInspectable
}

func (this *ICameraIntrinsicsFactory) Vtbl() *ICameraIntrinsicsFactoryVtbl {
	return (*ICameraIntrinsicsFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICameraIntrinsicsFactory) Create(focalLength unsafe.Pointer, principalPoint unsafe.Pointer, radialDistortion unsafe.Pointer, tangentialDistortion unsafe.Pointer, imageWidth uint32, imageHeight uint32) *ICameraIntrinsics {
	var _result *ICameraIntrinsics
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(focalLength), uintptr(principalPoint), uintptr(radialDistortion), uintptr(tangentialDistortion), uintptr(imageWidth), uintptr(imageHeight), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F95D89FB-8AF0-4CB0-926D-696866E5046A
var IID_IDepthCorrelatedCoordinateMapper = syscall.GUID{0xF95D89FB, 0x8AF0, 0x4CB0,
	[8]byte{0x92, 0x6D, 0x69, 0x68, 0x66, 0xE5, 0x04, 0x6A}}

type IDepthCorrelatedCoordinateMapperInterface interface {
	win32.IInspectableInterface
	UnprojectPoint(sourcePoint Point, targetCoordinateSystem *ISpatialCoordinateSystem) unsafe.Pointer
	UnprojectPoints(sourcePointsLength uint32, sourcePoints *Point, targetCoordinateSystem *ISpatialCoordinateSystem, results unsafe.Pointer)
	MapPoint(sourcePoint Point, targetCoordinateSystem *ISpatialCoordinateSystem, targetCameraIntrinsics *ICameraIntrinsics) Point
	MapPoints(sourcePointsLength uint32, sourcePoints *Point, targetCoordinateSystem *ISpatialCoordinateSystem, targetCameraIntrinsics *ICameraIntrinsics, resultsLength uint32, results *Point)
}

type IDepthCorrelatedCoordinateMapperVtbl struct {
	win32.IInspectableVtbl
	UnprojectPoint  uintptr
	UnprojectPoints uintptr
	MapPoint        uintptr
	MapPoints       uintptr
}

type IDepthCorrelatedCoordinateMapper struct {
	win32.IInspectable
}

func (this *IDepthCorrelatedCoordinateMapper) Vtbl() *IDepthCorrelatedCoordinateMapperVtbl {
	return (*IDepthCorrelatedCoordinateMapperVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDepthCorrelatedCoordinateMapper) UnprojectPoint(sourcePoint Point, targetCoordinateSystem *ISpatialCoordinateSystem) unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UnprojectPoint, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&sourcePoint)), uintptr(unsafe.Pointer(targetCoordinateSystem)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDepthCorrelatedCoordinateMapper) UnprojectPoints(sourcePointsLength uint32, sourcePoints *Point, targetCoordinateSystem *ISpatialCoordinateSystem, results unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UnprojectPoints, uintptr(unsafe.Pointer(this)), uintptr(sourcePointsLength), uintptr(unsafe.Pointer(sourcePoints)), uintptr(unsafe.Pointer(targetCoordinateSystem)), uintptr(results))
	_ = _hr
}

func (this *IDepthCorrelatedCoordinateMapper) MapPoint(sourcePoint Point, targetCoordinateSystem *ISpatialCoordinateSystem, targetCameraIntrinsics *ICameraIntrinsics) Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MapPoint, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&sourcePoint)), uintptr(unsafe.Pointer(targetCoordinateSystem)), uintptr(unsafe.Pointer(targetCameraIntrinsics)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDepthCorrelatedCoordinateMapper) MapPoints(sourcePointsLength uint32, sourcePoints *Point, targetCoordinateSystem *ISpatialCoordinateSystem, targetCameraIntrinsics *ICameraIntrinsics, resultsLength uint32, results *Point) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MapPoints, uintptr(unsafe.Pointer(this)), uintptr(sourcePointsLength), uintptr(unsafe.Pointer(sourcePoints)), uintptr(unsafe.Pointer(targetCoordinateSystem)), uintptr(unsafe.Pointer(targetCameraIntrinsics)), uintptr(resultsLength), uintptr(unsafe.Pointer(results)))
	_ = _hr
}

// A8FFAE60-4E9E-4377-A789-E24C4AE7E544
var IID_IFrameControlCapabilities = syscall.GUID{0xA8FFAE60, 0x4E9E, 0x4377,
	[8]byte{0xA7, 0x89, 0xE2, 0x4C, 0x4A, 0xE7, 0xE5, 0x44}}

type IFrameControlCapabilitiesInterface interface {
	win32.IInspectableInterface
	Get_Exposure() *IFrameExposureCapabilities
	Get_ExposureCompensation() *IFrameExposureCompensationCapabilities
	Get_IsoSpeed() *IFrameIsoSpeedCapabilities
	Get_Focus() *IFrameFocusCapabilities
	Get_PhotoConfirmationSupported() bool
}

type IFrameControlCapabilitiesVtbl struct {
	win32.IInspectableVtbl
	Get_Exposure                   uintptr
	Get_ExposureCompensation       uintptr
	Get_IsoSpeed                   uintptr
	Get_Focus                      uintptr
	Get_PhotoConfirmationSupported uintptr
}

type IFrameControlCapabilities struct {
	win32.IInspectable
}

func (this *IFrameControlCapabilities) Vtbl() *IFrameControlCapabilitiesVtbl {
	return (*IFrameControlCapabilitiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFrameControlCapabilities) Get_Exposure() *IFrameExposureCapabilities {
	var _result *IFrameExposureCapabilities
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Exposure, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFrameControlCapabilities) Get_ExposureCompensation() *IFrameExposureCompensationCapabilities {
	var _result *IFrameExposureCompensationCapabilities
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExposureCompensation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFrameControlCapabilities) Get_IsoSpeed() *IFrameIsoSpeedCapabilities {
	var _result *IFrameIsoSpeedCapabilities
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsoSpeed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFrameControlCapabilities) Get_Focus() *IFrameFocusCapabilities {
	var _result *IFrameFocusCapabilities
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Focus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFrameControlCapabilities) Get_PhotoConfirmationSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhotoConfirmationSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// CE9B0464-4730-440F-BD3E-EFE8A8F230A8
var IID_IFrameControlCapabilities2 = syscall.GUID{0xCE9B0464, 0x4730, 0x440F,
	[8]byte{0xBD, 0x3E, 0xEF, 0xE8, 0xA8, 0xF2, 0x30, 0xA8}}

type IFrameControlCapabilities2Interface interface {
	win32.IInspectableInterface
	Get_Flash() *IFrameFlashCapabilities
}

type IFrameControlCapabilities2Vtbl struct {
	win32.IInspectableVtbl
	Get_Flash uintptr
}

type IFrameControlCapabilities2 struct {
	win32.IInspectable
}

func (this *IFrameControlCapabilities2) Vtbl() *IFrameControlCapabilities2Vtbl {
	return (*IFrameControlCapabilities2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFrameControlCapabilities2) Get_Flash() *IFrameFlashCapabilities {
	var _result *IFrameFlashCapabilities
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Flash, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C16459D9-BAEF-4052-9177-48AFF2AF7522
var IID_IFrameController = syscall.GUID{0xC16459D9, 0xBAEF, 0x4052,
	[8]byte{0x91, 0x77, 0x48, 0xAF, 0xF2, 0xAF, 0x75, 0x22}}

type IFrameControllerInterface interface {
	win32.IInspectableInterface
	Get_ExposureControl() *IFrameExposureControl
	Get_ExposureCompensationControl() *IFrameExposureCompensationControl
	Get_IsoSpeedControl() *IFrameIsoSpeedControl
	Get_FocusControl() *IFrameFocusControl
	Get_PhotoConfirmationEnabled() *IReference[bool]
	Put_PhotoConfirmationEnabled(value *IReference[bool])
}

type IFrameControllerVtbl struct {
	win32.IInspectableVtbl
	Get_ExposureControl             uintptr
	Get_ExposureCompensationControl uintptr
	Get_IsoSpeedControl             uintptr
	Get_FocusControl                uintptr
	Get_PhotoConfirmationEnabled    uintptr
	Put_PhotoConfirmationEnabled    uintptr
}

type IFrameController struct {
	win32.IInspectable
}

func (this *IFrameController) Vtbl() *IFrameControllerVtbl {
	return (*IFrameControllerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFrameController) Get_ExposureControl() *IFrameExposureControl {
	var _result *IFrameExposureControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExposureControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFrameController) Get_ExposureCompensationControl() *IFrameExposureCompensationControl {
	var _result *IFrameExposureCompensationControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExposureCompensationControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFrameController) Get_IsoSpeedControl() *IFrameIsoSpeedControl {
	var _result *IFrameIsoSpeedControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsoSpeedControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFrameController) Get_FocusControl() *IFrameFocusControl {
	var _result *IFrameFocusControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FocusControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFrameController) Get_PhotoConfirmationEnabled() *IReference[bool] {
	var _result *IReference[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhotoConfirmationEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFrameController) Put_PhotoConfirmationEnabled(value *IReference[bool]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PhotoConfirmationEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 00D3BC75-D87C-485B-8A09-5C358568B427
var IID_IFrameController2 = syscall.GUID{0x00D3BC75, 0xD87C, 0x485B,
	[8]byte{0x8A, 0x09, 0x5C, 0x35, 0x85, 0x68, 0xB4, 0x27}}

type IFrameController2Interface interface {
	win32.IInspectableInterface
	Get_FlashControl() *IFrameFlashControl
}

type IFrameController2Vtbl struct {
	win32.IInspectableVtbl
	Get_FlashControl uintptr
}

type IFrameController2 struct {
	win32.IInspectable
}

func (this *IFrameController2) Vtbl() *IFrameController2Vtbl {
	return (*IFrameController2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFrameController2) Get_FlashControl() *IFrameFlashControl {
	var _result *IFrameFlashControl
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FlashControl, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BDBE9CE3-3985-4E72-97C2-0590D61307A1
var IID_IFrameExposureCapabilities = syscall.GUID{0xBDBE9CE3, 0x3985, 0x4E72,
	[8]byte{0x97, 0xC2, 0x05, 0x90, 0xD6, 0x13, 0x07, 0xA1}}

type IFrameExposureCapabilitiesInterface interface {
	win32.IInspectableInterface
	Get_Supported() bool
	Get_Min() TimeSpan
	Get_Max() TimeSpan
	Get_Step() TimeSpan
}

type IFrameExposureCapabilitiesVtbl struct {
	win32.IInspectableVtbl
	Get_Supported uintptr
	Get_Min       uintptr
	Get_Max       uintptr
	Get_Step      uintptr
}

type IFrameExposureCapabilities struct {
	win32.IInspectable
}

func (this *IFrameExposureCapabilities) Vtbl() *IFrameExposureCapabilitiesVtbl {
	return (*IFrameExposureCapabilitiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFrameExposureCapabilities) Get_Supported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Supported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFrameExposureCapabilities) Get_Min() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Min, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFrameExposureCapabilities) Get_Max() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Max, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFrameExposureCapabilities) Get_Step() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Step, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// B988A823-8065-41EE-B04F-722265954500
var IID_IFrameExposureCompensationCapabilities = syscall.GUID{0xB988A823, 0x8065, 0x41EE,
	[8]byte{0xB0, 0x4F, 0x72, 0x22, 0x65, 0x95, 0x45, 0x00}}

type IFrameExposureCompensationCapabilitiesInterface interface {
	win32.IInspectableInterface
	Get_Supported() bool
	Get_Min() float32
	Get_Max() float32
	Get_Step() float32
}

type IFrameExposureCompensationCapabilitiesVtbl struct {
	win32.IInspectableVtbl
	Get_Supported uintptr
	Get_Min       uintptr
	Get_Max       uintptr
	Get_Step      uintptr
}

type IFrameExposureCompensationCapabilities struct {
	win32.IInspectable
}

func (this *IFrameExposureCompensationCapabilities) Vtbl() *IFrameExposureCompensationCapabilitiesVtbl {
	return (*IFrameExposureCompensationCapabilitiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFrameExposureCompensationCapabilities) Get_Supported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Supported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFrameExposureCompensationCapabilities) Get_Min() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Min, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFrameExposureCompensationCapabilities) Get_Max() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Max, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFrameExposureCompensationCapabilities) Get_Step() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Step, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// E95896C9-F7F9-48CA-8591-A26531CB1578
var IID_IFrameExposureCompensationControl = syscall.GUID{0xE95896C9, 0xF7F9, 0x48CA,
	[8]byte{0x85, 0x91, 0xA2, 0x65, 0x31, 0xCB, 0x15, 0x78}}

type IFrameExposureCompensationControlInterface interface {
	win32.IInspectableInterface
	Get_Value() *IReference[float32]
	Put_Value(value *IReference[float32])
}

type IFrameExposureCompensationControlVtbl struct {
	win32.IInspectableVtbl
	Get_Value uintptr
	Put_Value uintptr
}

type IFrameExposureCompensationControl struct {
	win32.IInspectable
}

func (this *IFrameExposureCompensationControl) Vtbl() *IFrameExposureCompensationControlVtbl {
	return (*IFrameExposureCompensationControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFrameExposureCompensationControl) Get_Value() *IReference[float32] {
	var _result *IReference[float32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFrameExposureCompensationControl) Put_Value(value *IReference[float32]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// B1605A61-FFAF-4752-B621-F5B6F117F432
var IID_IFrameExposureControl = syscall.GUID{0xB1605A61, 0xFFAF, 0x4752,
	[8]byte{0xB6, 0x21, 0xF5, 0xB6, 0xF1, 0x17, 0xF4, 0x32}}

type IFrameExposureControlInterface interface {
	win32.IInspectableInterface
	Get_Auto() bool
	Put_Auto(value bool)
	Get_Value() *IReference[TimeSpan]
	Put_Value(value *IReference[TimeSpan])
}

type IFrameExposureControlVtbl struct {
	win32.IInspectableVtbl
	Get_Auto  uintptr
	Put_Auto  uintptr
	Get_Value uintptr
	Put_Value uintptr
}

type IFrameExposureControl struct {
	win32.IInspectable
}

func (this *IFrameExposureControl) Vtbl() *IFrameExposureControlVtbl {
	return (*IFrameExposureControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFrameExposureControl) Get_Auto() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Auto, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFrameExposureControl) Put_Auto(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Auto, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IFrameExposureControl) Get_Value() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFrameExposureControl) Put_Value(value *IReference[TimeSpan]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// BB9341A2-5EBE-4F62-8223-0E2B05BFBBD0
var IID_IFrameFlashCapabilities = syscall.GUID{0xBB9341A2, 0x5EBE, 0x4F62,
	[8]byte{0x82, 0x23, 0x0E, 0x2B, 0x05, 0xBF, 0xBB, 0xD0}}

type IFrameFlashCapabilitiesInterface interface {
	win32.IInspectableInterface
	Get_Supported() bool
	Get_RedEyeReductionSupported() bool
	Get_PowerSupported() bool
}

type IFrameFlashCapabilitiesVtbl struct {
	win32.IInspectableVtbl
	Get_Supported                uintptr
	Get_RedEyeReductionSupported uintptr
	Get_PowerSupported           uintptr
}

type IFrameFlashCapabilities struct {
	win32.IInspectable
}

func (this *IFrameFlashCapabilities) Vtbl() *IFrameFlashCapabilitiesVtbl {
	return (*IFrameFlashCapabilitiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFrameFlashCapabilities) Get_Supported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Supported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFrameFlashCapabilities) Get_RedEyeReductionSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RedEyeReductionSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFrameFlashCapabilities) Get_PowerSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PowerSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 75D5F6C7-BD45-4FAB-9375-45AC04B332C2
var IID_IFrameFlashControl = syscall.GUID{0x75D5F6C7, 0xBD45, 0x4FAB,
	[8]byte{0x93, 0x75, 0x45, 0xAC, 0x04, 0xB3, 0x32, 0xC2}}

type IFrameFlashControlInterface interface {
	win32.IInspectableInterface
	Get_Mode() FrameFlashMode
	Put_Mode(value FrameFlashMode)
	Get_Auto() bool
	Put_Auto(value bool)
	Get_RedEyeReduction() bool
	Put_RedEyeReduction(value bool)
	Get_PowerPercent() float32
	Put_PowerPercent(value float32)
}

type IFrameFlashControlVtbl struct {
	win32.IInspectableVtbl
	Get_Mode            uintptr
	Put_Mode            uintptr
	Get_Auto            uintptr
	Put_Auto            uintptr
	Get_RedEyeReduction uintptr
	Put_RedEyeReduction uintptr
	Get_PowerPercent    uintptr
	Put_PowerPercent    uintptr
}

type IFrameFlashControl struct {
	win32.IInspectable
}

func (this *IFrameFlashControl) Vtbl() *IFrameFlashControlVtbl {
	return (*IFrameFlashControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFrameFlashControl) Get_Mode() FrameFlashMode {
	var _result FrameFlashMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFrameFlashControl) Put_Mode(value FrameFlashMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Mode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IFrameFlashControl) Get_Auto() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Auto, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFrameFlashControl) Put_Auto(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Auto, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IFrameFlashControl) Get_RedEyeReduction() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RedEyeReduction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFrameFlashControl) Put_RedEyeReduction(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RedEyeReduction, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IFrameFlashControl) Get_PowerPercent() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PowerPercent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFrameFlashControl) Put_PowerPercent(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PowerPercent, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 7B25CD58-01C0-4065-9C40-C1A721425C1A
var IID_IFrameFocusCapabilities = syscall.GUID{0x7B25CD58, 0x01C0, 0x4065,
	[8]byte{0x9C, 0x40, 0xC1, 0xA7, 0x21, 0x42, 0x5C, 0x1A}}

type IFrameFocusCapabilitiesInterface interface {
	win32.IInspectableInterface
	Get_Supported() bool
	Get_Min() uint32
	Get_Max() uint32
	Get_Step() uint32
}

type IFrameFocusCapabilitiesVtbl struct {
	win32.IInspectableVtbl
	Get_Supported uintptr
	Get_Min       uintptr
	Get_Max       uintptr
	Get_Step      uintptr
}

type IFrameFocusCapabilities struct {
	win32.IInspectable
}

func (this *IFrameFocusCapabilities) Vtbl() *IFrameFocusCapabilitiesVtbl {
	return (*IFrameFocusCapabilitiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFrameFocusCapabilities) Get_Supported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Supported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFrameFocusCapabilities) Get_Min() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Min, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFrameFocusCapabilities) Get_Max() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Max, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFrameFocusCapabilities) Get_Step() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Step, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 272DF1D0-D912-4214-A67B-E38A8D48D8C6
var IID_IFrameFocusControl = syscall.GUID{0x272DF1D0, 0xD912, 0x4214,
	[8]byte{0xA6, 0x7B, 0xE3, 0x8A, 0x8D, 0x48, 0xD8, 0xC6}}

type IFrameFocusControlInterface interface {
	win32.IInspectableInterface
	Get_Value() *IReference[uint32]
	Put_Value(value *IReference[uint32])
}

type IFrameFocusControlVtbl struct {
	win32.IInspectableVtbl
	Get_Value uintptr
	Put_Value uintptr
}

type IFrameFocusControl struct {
	win32.IInspectable
}

func (this *IFrameFocusControl) Vtbl() *IFrameFocusControlVtbl {
	return (*IFrameFocusControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFrameFocusControl) Get_Value() *IReference[uint32] {
	var _result *IReference[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFrameFocusControl) Put_Value(value *IReference[uint32]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 16BDFF61-6DF6-4AC9-B92A-9F6ECD1AD2FA
var IID_IFrameIsoSpeedCapabilities = syscall.GUID{0x16BDFF61, 0x6DF6, 0x4AC9,
	[8]byte{0xB9, 0x2A, 0x9F, 0x6E, 0xCD, 0x1A, 0xD2, 0xFA}}

type IFrameIsoSpeedCapabilitiesInterface interface {
	win32.IInspectableInterface
	Get_Supported() bool
	Get_Min() uint32
	Get_Max() uint32
	Get_Step() uint32
}

type IFrameIsoSpeedCapabilitiesVtbl struct {
	win32.IInspectableVtbl
	Get_Supported uintptr
	Get_Min       uintptr
	Get_Max       uintptr
	Get_Step      uintptr
}

type IFrameIsoSpeedCapabilities struct {
	win32.IInspectable
}

func (this *IFrameIsoSpeedCapabilities) Vtbl() *IFrameIsoSpeedCapabilitiesVtbl {
	return (*IFrameIsoSpeedCapabilitiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFrameIsoSpeedCapabilities) Get_Supported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Supported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFrameIsoSpeedCapabilities) Get_Min() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Min, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFrameIsoSpeedCapabilities) Get_Max() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Max, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFrameIsoSpeedCapabilities) Get_Step() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Step, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 1A03EFED-786A-4C75-A557-7AB9A85F588C
var IID_IFrameIsoSpeedControl = syscall.GUID{0x1A03EFED, 0x786A, 0x4C75,
	[8]byte{0xA5, 0x57, 0x7A, 0xB9, 0xA8, 0x5F, 0x58, 0x8C}}

type IFrameIsoSpeedControlInterface interface {
	win32.IInspectableInterface
	Get_Auto() bool
	Put_Auto(value bool)
	Get_Value() *IReference[uint32]
	Put_Value(value *IReference[uint32])
}

type IFrameIsoSpeedControlVtbl struct {
	win32.IInspectableVtbl
	Get_Auto  uintptr
	Put_Auto  uintptr
	Get_Value uintptr
	Put_Value uintptr
}

type IFrameIsoSpeedControl struct {
	win32.IInspectable
}

func (this *IFrameIsoSpeedControl) Vtbl() *IFrameIsoSpeedControlVtbl {
	return (*IFrameIsoSpeedControlVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFrameIsoSpeedControl) Get_Auto() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Auto, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFrameIsoSpeedControl) Put_Auto(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Auto, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IFrameIsoSpeedControl) Get_Value() *IReference[uint32] {
	var _result *IReference[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFrameIsoSpeedControl) Put_Value(value *IReference[uint32]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// 7FBFF880-ED8C-43FD-A7C3-B35809E4229A
var IID_IVariablePhotoSequenceController = syscall.GUID{0x7FBFF880, 0xED8C, 0x43FD,
	[8]byte{0xA7, 0xC3, 0xB3, 0x58, 0x09, 0xE4, 0x22, 0x9A}}

type IVariablePhotoSequenceControllerInterface interface {
	win32.IInspectableInterface
	Get_Supported() bool
	Get_MaxPhotosPerSecond() float32
	Get_PhotosPerSecondLimit() float32
	Put_PhotosPerSecondLimit(value float32)
	GetHighestConcurrentFrameRate(captureProperties *IMediaEncodingProperties) *IMediaRatio
	GetCurrentFrameRate() *IMediaRatio
	Get_FrameCapabilities() *IFrameControlCapabilities
	Get_DesiredFrameControllers() *IVector[*IFrameController]
}

type IVariablePhotoSequenceControllerVtbl struct {
	win32.IInspectableVtbl
	Get_Supported                 uintptr
	Get_MaxPhotosPerSecond        uintptr
	Get_PhotosPerSecondLimit      uintptr
	Put_PhotosPerSecondLimit      uintptr
	GetHighestConcurrentFrameRate uintptr
	GetCurrentFrameRate           uintptr
	Get_FrameCapabilities         uintptr
	Get_DesiredFrameControllers   uintptr
}

type IVariablePhotoSequenceController struct {
	win32.IInspectable
}

func (this *IVariablePhotoSequenceController) Vtbl() *IVariablePhotoSequenceControllerVtbl {
	return (*IVariablePhotoSequenceControllerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVariablePhotoSequenceController) Get_Supported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Supported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVariablePhotoSequenceController) Get_MaxPhotosPerSecond() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxPhotosPerSecond, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVariablePhotoSequenceController) Get_PhotosPerSecondLimit() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhotosPerSecondLimit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVariablePhotoSequenceController) Put_PhotosPerSecondLimit(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PhotosPerSecondLimit, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVariablePhotoSequenceController) GetHighestConcurrentFrameRate(captureProperties *IMediaEncodingProperties) *IMediaRatio {
	var _result *IMediaRatio
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetHighestConcurrentFrameRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(captureProperties)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVariablePhotoSequenceController) GetCurrentFrameRate() *IMediaRatio {
	var _result *IMediaRatio
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentFrameRate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVariablePhotoSequenceController) Get_FrameCapabilities() *IFrameControlCapabilities {
	var _result *IFrameControlCapabilities
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameCapabilities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVariablePhotoSequenceController) Get_DesiredFrameControllers() *IVector[*IFrameController] {
	var _result *IVector[*IFrameController]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DesiredFrameControllers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type CameraIntrinsics struct {
	RtClass
	*ICameraIntrinsics
}

func NewCameraIntrinsics_Create(focalLength unsafe.Pointer, principalPoint unsafe.Pointer, radialDistortion unsafe.Pointer, tangentialDistortion unsafe.Pointer, imageWidth uint32, imageHeight uint32) *CameraIntrinsics {
	hs := NewHStr("Windows.Media.Devices.Core.CameraIntrinsics")
	var pFac *ICameraIntrinsicsFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICameraIntrinsicsFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ICameraIntrinsics
	p = pFac.Create(focalLength, principalPoint, radialDistortion, tangentialDistortion, imageWidth, imageHeight)
	result := &CameraIntrinsics{
		RtClass:           RtClass{PInspect: &p.IInspectable},
		ICameraIntrinsics: p,
	}
	com.AddToScope(result)
	return result
}

type DepthCorrelatedCoordinateMapper struct {
	RtClass
	*IDepthCorrelatedCoordinateMapper
}

type FrameControlCapabilities struct {
	RtClass
	*IFrameControlCapabilities
}

type FrameController struct {
	RtClass
	*IFrameController
}

func NewFrameController() *FrameController {
	hs := NewHStr("Windows.Media.Devices.Core.FrameController")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &FrameController{
		RtClass:          RtClass{PInspect: p},
		IFrameController: (*IFrameController)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type FrameExposureCapabilities struct {
	RtClass
	*IFrameExposureCapabilities
}

type FrameExposureCompensationCapabilities struct {
	RtClass
	*IFrameExposureCompensationCapabilities
}

type FrameExposureCompensationControl struct {
	RtClass
	*IFrameExposureCompensationControl
}

type FrameExposureControl struct {
	RtClass
	*IFrameExposureControl
}

type FrameFlashCapabilities struct {
	RtClass
	*IFrameFlashCapabilities
}

type FrameFlashControl struct {
	RtClass
	*IFrameFlashControl
}

type FrameFocusCapabilities struct {
	RtClass
	*IFrameFocusCapabilities
}

type FrameFocusControl struct {
	RtClass
	*IFrameFocusControl
}

type FrameIsoSpeedCapabilities struct {
	RtClass
	*IFrameIsoSpeedCapabilities
}

type FrameIsoSpeedControl struct {
	RtClass
	*IFrameIsoSpeedControl
}

type VariablePhotoSequenceController struct {
	RtClass
	*IVariablePhotoSequenceController
}
