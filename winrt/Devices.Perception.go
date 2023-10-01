package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type PerceptionFrameSourceAccessStatus int32

const (
	PerceptionFrameSourceAccessStatus_Unspecified    PerceptionFrameSourceAccessStatus = 0
	PerceptionFrameSourceAccessStatus_Allowed        PerceptionFrameSourceAccessStatus = 1
	PerceptionFrameSourceAccessStatus_DeniedByUser   PerceptionFrameSourceAccessStatus = 2
	PerceptionFrameSourceAccessStatus_DeniedBySystem PerceptionFrameSourceAccessStatus = 3
)

// enum
type PerceptionFrameSourcePropertyChangeStatus int32

const (
	PerceptionFrameSourcePropertyChangeStatus_Unknown              PerceptionFrameSourcePropertyChangeStatus = 0
	PerceptionFrameSourcePropertyChangeStatus_Accepted             PerceptionFrameSourcePropertyChangeStatus = 1
	PerceptionFrameSourcePropertyChangeStatus_LostControl          PerceptionFrameSourcePropertyChangeStatus = 2
	PerceptionFrameSourcePropertyChangeStatus_PropertyNotSupported PerceptionFrameSourcePropertyChangeStatus = 3
	PerceptionFrameSourcePropertyChangeStatus_PropertyReadOnly     PerceptionFrameSourcePropertyChangeStatus = 4
	PerceptionFrameSourcePropertyChangeStatus_ValueOutOfRange      PerceptionFrameSourcePropertyChangeStatus = 5
)

// interfaces

// 08C03978-437A-4D97-A663-FD3195600249
var IID_IKnownCameraIntrinsicsPropertiesStatics = syscall.GUID{0x08C03978, 0x437A, 0x4D97,
	[8]byte{0xA6, 0x63, 0xFD, 0x31, 0x95, 0x60, 0x02, 0x49}}

type IKnownCameraIntrinsicsPropertiesStaticsInterface interface {
	win32.IInspectableInterface
	Get_FocalLength() string
	Get_PrincipalPoint() string
	Get_RadialDistortion() string
	Get_TangentialDistortion() string
}

type IKnownCameraIntrinsicsPropertiesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_FocalLength          uintptr
	Get_PrincipalPoint       uintptr
	Get_RadialDistortion     uintptr
	Get_TangentialDistortion uintptr
}

type IKnownCameraIntrinsicsPropertiesStatics struct {
	win32.IInspectable
}

func (this *IKnownCameraIntrinsicsPropertiesStatics) Vtbl() *IKnownCameraIntrinsicsPropertiesStaticsVtbl {
	return (*IKnownCameraIntrinsicsPropertiesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKnownCameraIntrinsicsPropertiesStatics) Get_FocalLength() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FocalLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownCameraIntrinsicsPropertiesStatics) Get_PrincipalPoint() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PrincipalPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownCameraIntrinsicsPropertiesStatics) Get_RadialDistortion() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RadialDistortion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownCameraIntrinsicsPropertiesStatics) Get_TangentialDistortion() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TangentialDistortion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 5DF1CCA2-01F8-4A87-B859-D5E5B7E1DE4B
var IID_IKnownPerceptionColorFrameSourcePropertiesStatics = syscall.GUID{0x5DF1CCA2, 0x01F8, 0x4A87,
	[8]byte{0xB8, 0x59, 0xD5, 0xE5, 0xB7, 0xE1, 0xDE, 0x4B}}

type IKnownPerceptionColorFrameSourcePropertiesStaticsInterface interface {
	win32.IInspectableInterface
	Get_Exposure() string
	Get_AutoExposureEnabled() string
	Get_ExposureCompensation() string
}

type IKnownPerceptionColorFrameSourcePropertiesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Exposure             uintptr
	Get_AutoExposureEnabled  uintptr
	Get_ExposureCompensation uintptr
}

type IKnownPerceptionColorFrameSourcePropertiesStatics struct {
	win32.IInspectable
}

func (this *IKnownPerceptionColorFrameSourcePropertiesStatics) Vtbl() *IKnownPerceptionColorFrameSourcePropertiesStaticsVtbl {
	return (*IKnownPerceptionColorFrameSourcePropertiesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKnownPerceptionColorFrameSourcePropertiesStatics) Get_Exposure() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Exposure, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownPerceptionColorFrameSourcePropertiesStatics) Get_AutoExposureEnabled() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AutoExposureEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownPerceptionColorFrameSourcePropertiesStatics) Get_ExposureCompensation() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExposureCompensation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 5DF1CCA2-01F8-4A87-B859-D5E5B7E1DE4A
var IID_IKnownPerceptionDepthFrameSourcePropertiesStatics = syscall.GUID{0x5DF1CCA2, 0x01F8, 0x4A87,
	[8]byte{0xB8, 0x59, 0xD5, 0xE5, 0xB7, 0xE1, 0xDE, 0x4A}}

type IKnownPerceptionDepthFrameSourcePropertiesStaticsInterface interface {
	win32.IInspectableInterface
	Get_MinDepth() string
	Get_MaxDepth() string
}

type IKnownPerceptionDepthFrameSourcePropertiesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_MinDepth uintptr
	Get_MaxDepth uintptr
}

type IKnownPerceptionDepthFrameSourcePropertiesStatics struct {
	win32.IInspectable
}

func (this *IKnownPerceptionDepthFrameSourcePropertiesStatics) Vtbl() *IKnownPerceptionDepthFrameSourcePropertiesStaticsVtbl {
	return (*IKnownPerceptionDepthFrameSourcePropertiesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKnownPerceptionDepthFrameSourcePropertiesStatics) Get_MinDepth() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MinDepth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownPerceptionDepthFrameSourcePropertiesStatics) Get_MaxDepth() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxDepth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 5DF1CCA2-01F8-4A87-B859-D5E5B7E1DE47
var IID_IKnownPerceptionFrameSourcePropertiesStatics = syscall.GUID{0x5DF1CCA2, 0x01F8, 0x4A87,
	[8]byte{0xB8, 0x59, 0xD5, 0xE5, 0xB7, 0xE1, 0xDE, 0x47}}

type IKnownPerceptionFrameSourcePropertiesStaticsInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	Get_PhysicalDeviceIds() string
	Get_FrameKind() string
	Get_DeviceModelVersion() string
	Get_EnclosureLocation() string
}

type IKnownPerceptionFrameSourcePropertiesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Id                 uintptr
	Get_PhysicalDeviceIds  uintptr
	Get_FrameKind          uintptr
	Get_DeviceModelVersion uintptr
	Get_EnclosureLocation  uintptr
}

type IKnownPerceptionFrameSourcePropertiesStatics struct {
	win32.IInspectable
}

func (this *IKnownPerceptionFrameSourcePropertiesStatics) Vtbl() *IKnownPerceptionFrameSourcePropertiesStaticsVtbl {
	return (*IKnownPerceptionFrameSourcePropertiesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKnownPerceptionFrameSourcePropertiesStatics) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownPerceptionFrameSourcePropertiesStatics) Get_PhysicalDeviceIds() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhysicalDeviceIds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownPerceptionFrameSourcePropertiesStatics) Get_FrameKind() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownPerceptionFrameSourcePropertiesStatics) Get_DeviceModelVersion() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceModelVersion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownPerceptionFrameSourcePropertiesStatics) Get_EnclosureLocation() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EnclosureLocation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// A9C86871-05DC-4A4D-8A5C-A4ECF26BBC46
var IID_IKnownPerceptionFrameSourcePropertiesStatics2 = syscall.GUID{0xA9C86871, 0x05DC, 0x4A4D,
	[8]byte{0x8A, 0x5C, 0xA4, 0xEC, 0xF2, 0x6B, 0xBC, 0x46}}

type IKnownPerceptionFrameSourcePropertiesStatics2Interface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
}

type IKnownPerceptionFrameSourcePropertiesStatics2Vtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId uintptr
}

type IKnownPerceptionFrameSourcePropertiesStatics2 struct {
	win32.IInspectable
}

func (this *IKnownPerceptionFrameSourcePropertiesStatics2) Vtbl() *IKnownPerceptionFrameSourcePropertiesStatics2Vtbl {
	return (*IKnownPerceptionFrameSourcePropertiesStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKnownPerceptionFrameSourcePropertiesStatics2) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 5DF1CCA2-01F8-4A87-B859-D5E5B7E1DE49
var IID_IKnownPerceptionInfraredFrameSourcePropertiesStatics = syscall.GUID{0x5DF1CCA2, 0x01F8, 0x4A87,
	[8]byte{0xB8, 0x59, 0xD5, 0xE5, 0xB7, 0xE1, 0xDE, 0x49}}

type IKnownPerceptionInfraredFrameSourcePropertiesStaticsInterface interface {
	win32.IInspectableInterface
	Get_Exposure() string
	Get_AutoExposureEnabled() string
	Get_ExposureCompensation() string
	Get_ActiveIlluminationEnabled() string
	Get_AmbientSubtractionEnabled() string
	Get_StructureLightPatternEnabled() string
	Get_InterleavedIlluminationEnabled() string
}

type IKnownPerceptionInfraredFrameSourcePropertiesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Exposure                       uintptr
	Get_AutoExposureEnabled            uintptr
	Get_ExposureCompensation           uintptr
	Get_ActiveIlluminationEnabled      uintptr
	Get_AmbientSubtractionEnabled      uintptr
	Get_StructureLightPatternEnabled   uintptr
	Get_InterleavedIlluminationEnabled uintptr
}

type IKnownPerceptionInfraredFrameSourcePropertiesStatics struct {
	win32.IInspectable
}

func (this *IKnownPerceptionInfraredFrameSourcePropertiesStatics) Vtbl() *IKnownPerceptionInfraredFrameSourcePropertiesStaticsVtbl {
	return (*IKnownPerceptionInfraredFrameSourcePropertiesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKnownPerceptionInfraredFrameSourcePropertiesStatics) Get_Exposure() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Exposure, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownPerceptionInfraredFrameSourcePropertiesStatics) Get_AutoExposureEnabled() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AutoExposureEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownPerceptionInfraredFrameSourcePropertiesStatics) Get_ExposureCompensation() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExposureCompensation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownPerceptionInfraredFrameSourcePropertiesStatics) Get_ActiveIlluminationEnabled() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ActiveIlluminationEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownPerceptionInfraredFrameSourcePropertiesStatics) Get_AmbientSubtractionEnabled() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AmbientSubtractionEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownPerceptionInfraredFrameSourcePropertiesStatics) Get_StructureLightPatternEnabled() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StructureLightPatternEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownPerceptionInfraredFrameSourcePropertiesStatics) Get_InterleavedIlluminationEnabled() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InterleavedIlluminationEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 5DF1CCA2-01F8-4A87-B859-D5E5B7E1DE48
var IID_IKnownPerceptionVideoFrameSourcePropertiesStatics = syscall.GUID{0x5DF1CCA2, 0x01F8, 0x4A87,
	[8]byte{0xB8, 0x59, 0xD5, 0xE5, 0xB7, 0xE1, 0xDE, 0x48}}

type IKnownPerceptionVideoFrameSourcePropertiesStaticsInterface interface {
	win32.IInspectableInterface
	Get_VideoProfile() string
	Get_SupportedVideoProfiles() string
	Get_AvailableVideoProfiles() string
	Get_IsMirrored() string
	Get_CameraIntrinsics() string
}

type IKnownPerceptionVideoFrameSourcePropertiesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_VideoProfile           uintptr
	Get_SupportedVideoProfiles uintptr
	Get_AvailableVideoProfiles uintptr
	Get_IsMirrored             uintptr
	Get_CameraIntrinsics       uintptr
}

type IKnownPerceptionVideoFrameSourcePropertiesStatics struct {
	win32.IInspectable
}

func (this *IKnownPerceptionVideoFrameSourcePropertiesStatics) Vtbl() *IKnownPerceptionVideoFrameSourcePropertiesStaticsVtbl {
	return (*IKnownPerceptionVideoFrameSourcePropertiesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKnownPerceptionVideoFrameSourcePropertiesStatics) Get_VideoProfile() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoProfile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownPerceptionVideoFrameSourcePropertiesStatics) Get_SupportedVideoProfiles() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedVideoProfiles, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownPerceptionVideoFrameSourcePropertiesStatics) Get_AvailableVideoProfiles() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AvailableVideoProfiles, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownPerceptionVideoFrameSourcePropertiesStatics) Get_IsMirrored() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsMirrored, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownPerceptionVideoFrameSourcePropertiesStatics) Get_CameraIntrinsics() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CameraIntrinsics, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 8F08E2E7-5A76-43E3-A13A-DA3D91A9EF98
var IID_IKnownPerceptionVideoProfilePropertiesStatics = syscall.GUID{0x8F08E2E7, 0x5A76, 0x43E3,
	[8]byte{0xA1, 0x3A, 0xDA, 0x3D, 0x91, 0xA9, 0xEF, 0x98}}

type IKnownPerceptionVideoProfilePropertiesStaticsInterface interface {
	win32.IInspectableInterface
	Get_BitmapPixelFormat() string
	Get_BitmapAlphaMode() string
	Get_Width() string
	Get_Height() string
	Get_FrameDuration() string
}

type IKnownPerceptionVideoProfilePropertiesStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_BitmapPixelFormat uintptr
	Get_BitmapAlphaMode   uintptr
	Get_Width             uintptr
	Get_Height            uintptr
	Get_FrameDuration     uintptr
}

type IKnownPerceptionVideoProfilePropertiesStatics struct {
	win32.IInspectable
}

func (this *IKnownPerceptionVideoProfilePropertiesStatics) Vtbl() *IKnownPerceptionVideoProfilePropertiesStaticsVtbl {
	return (*IKnownPerceptionVideoProfilePropertiesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKnownPerceptionVideoProfilePropertiesStatics) Get_BitmapPixelFormat() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BitmapPixelFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownPerceptionVideoProfilePropertiesStatics) Get_BitmapAlphaMode() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BitmapAlphaMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownPerceptionVideoProfilePropertiesStatics) Get_Width() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Width, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownPerceptionVideoProfilePropertiesStatics) Get_Height() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Height, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IKnownPerceptionVideoProfilePropertiesStatics) Get_FrameDuration() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameDuration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// FE621549-2CBF-4F94-9861-F817EA317747
var IID_IPerceptionColorFrame = syscall.GUID{0xFE621549, 0x2CBF, 0x4F94,
	[8]byte{0x98, 0x61, 0xF8, 0x17, 0xEA, 0x31, 0x77, 0x47}}

type IPerceptionColorFrameInterface interface {
	win32.IInspectableInterface
	Get_VideoFrame() *IVideoFrame
}

type IPerceptionColorFrameVtbl struct {
	win32.IInspectableVtbl
	Get_VideoFrame uintptr
}

type IPerceptionColorFrame struct {
	win32.IInspectable
}

func (this *IPerceptionColorFrame) Vtbl() *IPerceptionColorFrameVtbl {
	return (*IPerceptionColorFrameVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionColorFrame) Get_VideoFrame() *IVideoFrame {
	var _result *IVideoFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8FAD02D5-86F7-4D8D-B966-5A3761BA9F59
var IID_IPerceptionColorFrameArrivedEventArgs = syscall.GUID{0x8FAD02D5, 0x86F7, 0x4D8D,
	[8]byte{0xB9, 0x66, 0x5A, 0x37, 0x61, 0xBA, 0x9F, 0x59}}

type IPerceptionColorFrameArrivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_RelativeTime() TimeSpan
	TryOpenFrame() *IPerceptionColorFrame
}

type IPerceptionColorFrameArrivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_RelativeTime uintptr
	TryOpenFrame     uintptr
}

type IPerceptionColorFrameArrivedEventArgs struct {
	win32.IInspectable
}

func (this *IPerceptionColorFrameArrivedEventArgs) Vtbl() *IPerceptionColorFrameArrivedEventArgsVtbl {
	return (*IPerceptionColorFrameArrivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionColorFrameArrivedEventArgs) Get_RelativeTime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RelativeTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionColorFrameArrivedEventArgs) TryOpenFrame() *IPerceptionColorFrame {
	var _result *IPerceptionColorFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryOpenFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7650F56E-B9F5-461B-83AD-F222AF2AAADC
var IID_IPerceptionColorFrameReader = syscall.GUID{0x7650F56E, 0xB9F5, 0x461B,
	[8]byte{0x83, 0xAD, 0xF2, 0x22, 0xAF, 0x2A, 0xAA, 0xDC}}

type IPerceptionColorFrameReaderInterface interface {
	win32.IInspectableInterface
	Add_FrameArrived(handler TypedEventHandler[*IPerceptionColorFrameReader, *IPerceptionColorFrameArrivedEventArgs]) EventRegistrationToken
	Remove_FrameArrived(token EventRegistrationToken)
	Get_Source() *IPerceptionColorFrameSource
	Get_IsPaused() bool
	Put_IsPaused(value bool)
	TryReadLatestFrame() *IPerceptionColorFrame
}

type IPerceptionColorFrameReaderVtbl struct {
	win32.IInspectableVtbl
	Add_FrameArrived    uintptr
	Remove_FrameArrived uintptr
	Get_Source          uintptr
	Get_IsPaused        uintptr
	Put_IsPaused        uintptr
	TryReadLatestFrame  uintptr
}

type IPerceptionColorFrameReader struct {
	win32.IInspectable
}

func (this *IPerceptionColorFrameReader) Vtbl() *IPerceptionColorFrameReaderVtbl {
	return (*IPerceptionColorFrameReaderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionColorFrameReader) Add_FrameArrived(handler TypedEventHandler[*IPerceptionColorFrameReader, *IPerceptionColorFrameArrivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_FrameArrived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionColorFrameReader) Remove_FrameArrived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_FrameArrived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionColorFrameReader) Get_Source() *IPerceptionColorFrameSource {
	var _result *IPerceptionColorFrameSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Source, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionColorFrameReader) Get_IsPaused() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPaused, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionColorFrameReader) Put_IsPaused(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsPaused, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPerceptionColorFrameReader) TryReadLatestFrame() *IPerceptionColorFrame {
	var _result *IPerceptionColorFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryReadLatestFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DC6DBA7C-0B58-468D-9CA1-6DB04CC0477C
var IID_IPerceptionColorFrameSource = syscall.GUID{0xDC6DBA7C, 0x0B58, 0x468D,
	[8]byte{0x9C, 0xA1, 0x6D, 0xB0, 0x4C, 0xC0, 0x47, 0x7C}}

type IPerceptionColorFrameSourceInterface interface {
	win32.IInspectableInterface
	Add_AvailableChanged(handler TypedEventHandler[*IPerceptionColorFrameSource, interface{}]) EventRegistrationToken
	Remove_AvailableChanged(token EventRegistrationToken)
	Add_ActiveChanged(handler TypedEventHandler[*IPerceptionColorFrameSource, interface{}]) EventRegistrationToken
	Remove_ActiveChanged(token EventRegistrationToken)
	Add_PropertiesChanged(handler TypedEventHandler[*IPerceptionColorFrameSource, *IPerceptionFrameSourcePropertiesChangedEventArgs]) EventRegistrationToken
	Remove_PropertiesChanged(token EventRegistrationToken)
	Add_VideoProfileChanged(handler TypedEventHandler[*IPerceptionColorFrameSource, interface{}]) EventRegistrationToken
	Remove_VideoProfileChanged(token EventRegistrationToken)
	Add_CameraIntrinsicsChanged(handler TypedEventHandler[*IPerceptionColorFrameSource, interface{}]) EventRegistrationToken
	Remove_CameraIntrinsicsChanged(token EventRegistrationToken)
	Get_Id() string
	Get_DisplayName() string
	Get_DeviceKind() string
	Get_Available() bool
	Get_Active() bool
	Get_IsControlled() bool
	Get_Properties() *IMapView[string, interface{}]
	Get_SupportedVideoProfiles() *IVectorView[*IPerceptionVideoProfile]
	Get_AvailableVideoProfiles() *IVectorView[*IPerceptionVideoProfile]
	Get_VideoProfile() *IPerceptionVideoProfile
	Get_CameraIntrinsics() *ICameraIntrinsics
	AcquireControlSession() *IPerceptionControlSession
	CanControlIndependentlyFrom(targetId string) bool
	IsCorrelatedWith(targetId string) bool
	TryGetTransformTo(targetId string, result Matrix4x4) bool
	TryGetDepthCorrelatedCameraIntrinsicsAsync(correlatedDepthFrameSource *IPerceptionDepthFrameSource) *IAsyncOperation[*IPerceptionDepthCorrelatedCameraIntrinsics]
	TryGetDepthCorrelatedCoordinateMapperAsync(targetSourceId string, correlatedDepthFrameSource *IPerceptionDepthFrameSource) *IAsyncOperation[*IPerceptionDepthCorrelatedCoordinateMapper]
	TrySetVideoProfileAsync(controlSession *IPerceptionControlSession, profile *IPerceptionVideoProfile) *IAsyncOperation[*IPerceptionFrameSourcePropertyChangeResult]
	OpenReader() *IPerceptionColorFrameReader
}

type IPerceptionColorFrameSourceVtbl struct {
	win32.IInspectableVtbl
	Add_AvailableChanged                       uintptr
	Remove_AvailableChanged                    uintptr
	Add_ActiveChanged                          uintptr
	Remove_ActiveChanged                       uintptr
	Add_PropertiesChanged                      uintptr
	Remove_PropertiesChanged                   uintptr
	Add_VideoProfileChanged                    uintptr
	Remove_VideoProfileChanged                 uintptr
	Add_CameraIntrinsicsChanged                uintptr
	Remove_CameraIntrinsicsChanged             uintptr
	Get_Id                                     uintptr
	Get_DisplayName                            uintptr
	Get_DeviceKind                             uintptr
	Get_Available                              uintptr
	Get_Active                                 uintptr
	Get_IsControlled                           uintptr
	Get_Properties                             uintptr
	Get_SupportedVideoProfiles                 uintptr
	Get_AvailableVideoProfiles                 uintptr
	Get_VideoProfile                           uintptr
	Get_CameraIntrinsics                       uintptr
	AcquireControlSession                      uintptr
	CanControlIndependentlyFrom                uintptr
	IsCorrelatedWith                           uintptr
	TryGetTransformTo                          uintptr
	TryGetDepthCorrelatedCameraIntrinsicsAsync uintptr
	TryGetDepthCorrelatedCoordinateMapperAsync uintptr
	TrySetVideoProfileAsync                    uintptr
	OpenReader                                 uintptr
}

type IPerceptionColorFrameSource struct {
	win32.IInspectable
}

func (this *IPerceptionColorFrameSource) Vtbl() *IPerceptionColorFrameSourceVtbl {
	return (*IPerceptionColorFrameSourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionColorFrameSource) Add_AvailableChanged(handler TypedEventHandler[*IPerceptionColorFrameSource, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AvailableChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionColorFrameSource) Remove_AvailableChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AvailableChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionColorFrameSource) Add_ActiveChanged(handler TypedEventHandler[*IPerceptionColorFrameSource, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ActiveChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionColorFrameSource) Remove_ActiveChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ActiveChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionColorFrameSource) Add_PropertiesChanged(handler TypedEventHandler[*IPerceptionColorFrameSource, *IPerceptionFrameSourcePropertiesChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PropertiesChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionColorFrameSource) Remove_PropertiesChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PropertiesChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionColorFrameSource) Add_VideoProfileChanged(handler TypedEventHandler[*IPerceptionColorFrameSource, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_VideoProfileChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionColorFrameSource) Remove_VideoProfileChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_VideoProfileChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionColorFrameSource) Add_CameraIntrinsicsChanged(handler TypedEventHandler[*IPerceptionColorFrameSource, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_CameraIntrinsicsChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionColorFrameSource) Remove_CameraIntrinsicsChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_CameraIntrinsicsChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionColorFrameSource) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPerceptionColorFrameSource) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPerceptionColorFrameSource) Get_DeviceKind() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPerceptionColorFrameSource) Get_Available() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Available, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionColorFrameSource) Get_Active() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Active, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionColorFrameSource) Get_IsControlled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsControlled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionColorFrameSource) Get_Properties() *IMapView[string, interface{}] {
	var _result *IMapView[string, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionColorFrameSource) Get_SupportedVideoProfiles() *IVectorView[*IPerceptionVideoProfile] {
	var _result *IVectorView[*IPerceptionVideoProfile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedVideoProfiles, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionColorFrameSource) Get_AvailableVideoProfiles() *IVectorView[*IPerceptionVideoProfile] {
	var _result *IVectorView[*IPerceptionVideoProfile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AvailableVideoProfiles, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionColorFrameSource) Get_VideoProfile() *IPerceptionVideoProfile {
	var _result *IPerceptionVideoProfile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoProfile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionColorFrameSource) Get_CameraIntrinsics() *ICameraIntrinsics {
	var _result *ICameraIntrinsics
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CameraIntrinsics, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionColorFrameSource) AcquireControlSession() *IPerceptionControlSession {
	var _result *IPerceptionControlSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AcquireControlSession, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionColorFrameSource) CanControlIndependentlyFrom(targetId string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CanControlIndependentlyFrom, uintptr(unsafe.Pointer(this)), NewHStr(targetId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionColorFrameSource) IsCorrelatedWith(targetId string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsCorrelatedWith, uintptr(unsafe.Pointer(this)), NewHStr(targetId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionColorFrameSource) TryGetTransformTo(targetId string, result Matrix4x4) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetTransformTo, uintptr(unsafe.Pointer(this)), NewHStr(targetId).Ptr, uintptr(unsafe.Pointer(&result)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionColorFrameSource) TryGetDepthCorrelatedCameraIntrinsicsAsync(correlatedDepthFrameSource *IPerceptionDepthFrameSource) *IAsyncOperation[*IPerceptionDepthCorrelatedCameraIntrinsics] {
	var _result *IAsyncOperation[*IPerceptionDepthCorrelatedCameraIntrinsics]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetDepthCorrelatedCameraIntrinsicsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(correlatedDepthFrameSource)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionColorFrameSource) TryGetDepthCorrelatedCoordinateMapperAsync(targetSourceId string, correlatedDepthFrameSource *IPerceptionDepthFrameSource) *IAsyncOperation[*IPerceptionDepthCorrelatedCoordinateMapper] {
	var _result *IAsyncOperation[*IPerceptionDepthCorrelatedCoordinateMapper]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetDepthCorrelatedCoordinateMapperAsync, uintptr(unsafe.Pointer(this)), NewHStr(targetSourceId).Ptr, uintptr(unsafe.Pointer(correlatedDepthFrameSource)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionColorFrameSource) TrySetVideoProfileAsync(controlSession *IPerceptionControlSession, profile *IPerceptionVideoProfile) *IAsyncOperation[*IPerceptionFrameSourcePropertyChangeResult] {
	var _result *IAsyncOperation[*IPerceptionFrameSourcePropertyChangeResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TrySetVideoProfileAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(controlSession)), uintptr(unsafe.Pointer(profile)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionColorFrameSource) OpenReader() *IPerceptionColorFrameReader {
	var _result *IPerceptionColorFrameReader
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenReader, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F88008E5-5631-45ED-AD98-8C6AA04CFB91
var IID_IPerceptionColorFrameSource2 = syscall.GUID{0xF88008E5, 0x5631, 0x45ED,
	[8]byte{0xAD, 0x98, 0x8C, 0x6A, 0xA0, 0x4C, 0xFB, 0x91}}

type IPerceptionColorFrameSource2Interface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
}

type IPerceptionColorFrameSource2Vtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId uintptr
}

type IPerceptionColorFrameSource2 struct {
	win32.IInspectable
}

func (this *IPerceptionColorFrameSource2) Vtbl() *IPerceptionColorFrameSource2Vtbl {
	return (*IPerceptionColorFrameSource2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionColorFrameSource2) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// D16BF4E6-DA24-442C-BBD5-55549B5B94F3
var IID_IPerceptionColorFrameSourceAddedEventArgs = syscall.GUID{0xD16BF4E6, 0xDA24, 0x442C,
	[8]byte{0xBB, 0xD5, 0x55, 0x54, 0x9B, 0x5B, 0x94, 0xF3}}

type IPerceptionColorFrameSourceAddedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_FrameSource() *IPerceptionColorFrameSource
}

type IPerceptionColorFrameSourceAddedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_FrameSource uintptr
}

type IPerceptionColorFrameSourceAddedEventArgs struct {
	win32.IInspectable
}

func (this *IPerceptionColorFrameSourceAddedEventArgs) Vtbl() *IPerceptionColorFrameSourceAddedEventArgsVtbl {
	return (*IPerceptionColorFrameSourceAddedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionColorFrameSourceAddedEventArgs) Get_FrameSource() *IPerceptionColorFrameSource {
	var _result *IPerceptionColorFrameSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D277FA69-EB4C-42EF-BA4F-288F615C93C1
var IID_IPerceptionColorFrameSourceRemovedEventArgs = syscall.GUID{0xD277FA69, 0xEB4C, 0x42EF,
	[8]byte{0xBA, 0x4F, 0x28, 0x8F, 0x61, 0x5C, 0x93, 0xC1}}

type IPerceptionColorFrameSourceRemovedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_FrameSource() *IPerceptionColorFrameSource
}

type IPerceptionColorFrameSourceRemovedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_FrameSource uintptr
}

type IPerceptionColorFrameSourceRemovedEventArgs struct {
	win32.IInspectable
}

func (this *IPerceptionColorFrameSourceRemovedEventArgs) Vtbl() *IPerceptionColorFrameSourceRemovedEventArgsVtbl {
	return (*IPerceptionColorFrameSourceRemovedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionColorFrameSourceRemovedEventArgs) Get_FrameSource() *IPerceptionColorFrameSource {
	var _result *IPerceptionColorFrameSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5DF3CCA2-01F8-4A87-B859-D5E5B7E1DE49
var IID_IPerceptionColorFrameSourceStatics = syscall.GUID{0x5DF3CCA2, 0x01F8, 0x4A87,
	[8]byte{0xB8, 0x59, 0xD5, 0xE5, 0xB7, 0xE1, 0xDE, 0x49}}

type IPerceptionColorFrameSourceStaticsInterface interface {
	win32.IInspectableInterface
	CreateWatcher() *IPerceptionColorFrameSourceWatcher
	FindAllAsync() *IAsyncOperation[*IVectorView[*IPerceptionColorFrameSource]]
	FromIdAsync(id string) *IAsyncOperation[*IPerceptionColorFrameSource]
	RequestAccessAsync() *IAsyncOperation[PerceptionFrameSourceAccessStatus]
}

type IPerceptionColorFrameSourceStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateWatcher      uintptr
	FindAllAsync       uintptr
	FromIdAsync        uintptr
	RequestAccessAsync uintptr
}

type IPerceptionColorFrameSourceStatics struct {
	win32.IInspectable
}

func (this *IPerceptionColorFrameSourceStatics) Vtbl() *IPerceptionColorFrameSourceStaticsVtbl {
	return (*IPerceptionColorFrameSourceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionColorFrameSourceStatics) CreateWatcher() *IPerceptionColorFrameSourceWatcher {
	var _result *IPerceptionColorFrameSourceWatcher
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWatcher, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionColorFrameSourceStatics) FindAllAsync() *IAsyncOperation[*IVectorView[*IPerceptionColorFrameSource]] {
	var _result *IAsyncOperation[*IVectorView[*IPerceptionColorFrameSource]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionColorFrameSourceStatics) FromIdAsync(id string) *IAsyncOperation[*IPerceptionColorFrameSource] {
	var _result *IAsyncOperation[*IPerceptionColorFrameSource]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(id).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionColorFrameSourceStatics) RequestAccessAsync() *IAsyncOperation[PerceptionFrameSourceAccessStatus] {
	var _result *IAsyncOperation[PerceptionFrameSourceAccessStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 96BD1392-E667-40C4-89F9-1462DEA6A9CC
var IID_IPerceptionColorFrameSourceWatcher = syscall.GUID{0x96BD1392, 0xE667, 0x40C4,
	[8]byte{0x89, 0xF9, 0x14, 0x62, 0xDE, 0xA6, 0xA9, 0xCC}}

type IPerceptionColorFrameSourceWatcherInterface interface {
	win32.IInspectableInterface
	Add_SourceAdded(handler TypedEventHandler[*IPerceptionColorFrameSourceWatcher, *IPerceptionColorFrameSourceAddedEventArgs]) EventRegistrationToken
	Remove_SourceAdded(token EventRegistrationToken)
	Add_SourceRemoved(handler TypedEventHandler[*IPerceptionColorFrameSourceWatcher, *IPerceptionColorFrameSourceRemovedEventArgs]) EventRegistrationToken
	Remove_SourceRemoved(token EventRegistrationToken)
	Add_Stopped(handler TypedEventHandler[*IPerceptionColorFrameSourceWatcher, interface{}]) EventRegistrationToken
	Remove_Stopped(token EventRegistrationToken)
	Add_EnumerationCompleted(handler TypedEventHandler[*IPerceptionColorFrameSourceWatcher, interface{}]) EventRegistrationToken
	Remove_EnumerationCompleted(token EventRegistrationToken)
	Get_Status() DeviceWatcherStatus
	Start()
	Stop()
}

type IPerceptionColorFrameSourceWatcherVtbl struct {
	win32.IInspectableVtbl
	Add_SourceAdded             uintptr
	Remove_SourceAdded          uintptr
	Add_SourceRemoved           uintptr
	Remove_SourceRemoved        uintptr
	Add_Stopped                 uintptr
	Remove_Stopped              uintptr
	Add_EnumerationCompleted    uintptr
	Remove_EnumerationCompleted uintptr
	Get_Status                  uintptr
	Start                       uintptr
	Stop                        uintptr
}

type IPerceptionColorFrameSourceWatcher struct {
	win32.IInspectable
}

func (this *IPerceptionColorFrameSourceWatcher) Vtbl() *IPerceptionColorFrameSourceWatcherVtbl {
	return (*IPerceptionColorFrameSourceWatcherVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionColorFrameSourceWatcher) Add_SourceAdded(handler TypedEventHandler[*IPerceptionColorFrameSourceWatcher, *IPerceptionColorFrameSourceAddedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SourceAdded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionColorFrameSourceWatcher) Remove_SourceAdded(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SourceAdded, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionColorFrameSourceWatcher) Add_SourceRemoved(handler TypedEventHandler[*IPerceptionColorFrameSourceWatcher, *IPerceptionColorFrameSourceRemovedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SourceRemoved, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionColorFrameSourceWatcher) Remove_SourceRemoved(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SourceRemoved, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionColorFrameSourceWatcher) Add_Stopped(handler TypedEventHandler[*IPerceptionColorFrameSourceWatcher, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Stopped, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionColorFrameSourceWatcher) Remove_Stopped(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Stopped, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionColorFrameSourceWatcher) Add_EnumerationCompleted(handler TypedEventHandler[*IPerceptionColorFrameSourceWatcher, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_EnumerationCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionColorFrameSourceWatcher) Remove_EnumerationCompleted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_EnumerationCompleted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionColorFrameSourceWatcher) Get_Status() DeviceWatcherStatus {
	var _result DeviceWatcherStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionColorFrameSourceWatcher) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IPerceptionColorFrameSourceWatcher) Stop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 99998653-5A3D-417F-9239-F1889E548B48
var IID_IPerceptionControlSession = syscall.GUID{0x99998653, 0x5A3D, 0x417F,
	[8]byte{0x92, 0x39, 0xF1, 0x88, 0x9E, 0x54, 0x8B, 0x48}}

type IPerceptionControlSessionInterface interface {
	win32.IInspectableInterface
	Add_ControlLost(handler TypedEventHandler[*IPerceptionControlSession, interface{}]) EventRegistrationToken
	Remove_ControlLost(token EventRegistrationToken)
	TrySetPropertyAsync(name string, value interface{}) *IAsyncOperation[*IPerceptionFrameSourcePropertyChangeResult]
}

type IPerceptionControlSessionVtbl struct {
	win32.IInspectableVtbl
	Add_ControlLost     uintptr
	Remove_ControlLost  uintptr
	TrySetPropertyAsync uintptr
}

type IPerceptionControlSession struct {
	win32.IInspectable
}

func (this *IPerceptionControlSession) Vtbl() *IPerceptionControlSessionVtbl {
	return (*IPerceptionControlSessionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionControlSession) Add_ControlLost(handler TypedEventHandler[*IPerceptionControlSession, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ControlLost, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionControlSession) Remove_ControlLost(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ControlLost, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionControlSession) TrySetPropertyAsync(name string, value interface{}) *IAsyncOperation[*IPerceptionFrameSourcePropertyChangeResult] {
	var _result *IAsyncOperation[*IPerceptionFrameSourcePropertyChangeResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TrySetPropertyAsync, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, *(*uintptr)(unsafe.Pointer(&value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6548CA01-86DE-5BE1-6582-807FCF4C95CF
var IID_IPerceptionDepthCorrelatedCameraIntrinsics = syscall.GUID{0x6548CA01, 0x86DE, 0x5BE1,
	[8]byte{0x65, 0x82, 0x80, 0x7F, 0xCF, 0x4C, 0x95, 0xCF}}

type IPerceptionDepthCorrelatedCameraIntrinsicsInterface interface {
	win32.IInspectableInterface
	UnprojectPixelAtCorrelatedDepth(pixelCoordinate Point, depthFrame *IPerceptionDepthFrame) Vector3
	UnprojectPixelsAtCorrelatedDepth(sourceCoordinatesLength uint32, sourceCoordinates *Point, depthFrame *IPerceptionDepthFrame, resultsLength uint32, results *Vector3)
	UnprojectRegionPixelsAtCorrelatedDepthAsync(region Rect, depthFrame *IPerceptionDepthFrame, resultsLength uint32, results *Vector3) *IAsyncAction
	UnprojectAllPixelsAtCorrelatedDepthAsync(depthFrame *IPerceptionDepthFrame, resultsLength uint32, results *Vector3) *IAsyncAction
}

type IPerceptionDepthCorrelatedCameraIntrinsicsVtbl struct {
	win32.IInspectableVtbl
	UnprojectPixelAtCorrelatedDepth             uintptr
	UnprojectPixelsAtCorrelatedDepth            uintptr
	UnprojectRegionPixelsAtCorrelatedDepthAsync uintptr
	UnprojectAllPixelsAtCorrelatedDepthAsync    uintptr
}

type IPerceptionDepthCorrelatedCameraIntrinsics struct {
	win32.IInspectable
}

func (this *IPerceptionDepthCorrelatedCameraIntrinsics) Vtbl() *IPerceptionDepthCorrelatedCameraIntrinsicsVtbl {
	return (*IPerceptionDepthCorrelatedCameraIntrinsicsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionDepthCorrelatedCameraIntrinsics) UnprojectPixelAtCorrelatedDepth(pixelCoordinate Point, depthFrame *IPerceptionDepthFrame) Vector3 {
	var _result Vector3
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UnprojectPixelAtCorrelatedDepth, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&pixelCoordinate)), uintptr(unsafe.Pointer(depthFrame)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionDepthCorrelatedCameraIntrinsics) UnprojectPixelsAtCorrelatedDepth(sourceCoordinatesLength uint32, sourceCoordinates *Point, depthFrame *IPerceptionDepthFrame, resultsLength uint32, results *Vector3) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UnprojectPixelsAtCorrelatedDepth, uintptr(unsafe.Pointer(this)), uintptr(sourceCoordinatesLength), uintptr(unsafe.Pointer(sourceCoordinates)), uintptr(unsafe.Pointer(depthFrame)), uintptr(resultsLength), uintptr(unsafe.Pointer(results)))
	_ = _hr
}

func (this *IPerceptionDepthCorrelatedCameraIntrinsics) UnprojectRegionPixelsAtCorrelatedDepthAsync(region Rect, depthFrame *IPerceptionDepthFrame, resultsLength uint32, results *Vector3) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UnprojectRegionPixelsAtCorrelatedDepthAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&region)), uintptr(unsafe.Pointer(depthFrame)), uintptr(resultsLength), uintptr(unsafe.Pointer(results)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionDepthCorrelatedCameraIntrinsics) UnprojectAllPixelsAtCorrelatedDepthAsync(depthFrame *IPerceptionDepthFrame, resultsLength uint32, results *Vector3) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UnprojectAllPixelsAtCorrelatedDepthAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(depthFrame)), uintptr(resultsLength), uintptr(unsafe.Pointer(results)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5B4D9D1D-B5F6-469C-B8C2-B97A45E6863B
var IID_IPerceptionDepthCorrelatedCoordinateMapper = syscall.GUID{0x5B4D9D1D, 0xB5F6, 0x469C,
	[8]byte{0xB8, 0xC2, 0xB9, 0x7A, 0x45, 0xE6, 0x86, 0x3B}}

type IPerceptionDepthCorrelatedCoordinateMapperInterface interface {
	win32.IInspectableInterface
	MapPixelToTarget(sourcePixelCoordinate Point, depthFrame *IPerceptionDepthFrame) Point
	MapPixelsToTarget(sourceCoordinatesLength uint32, sourceCoordinates *Point, depthFrame *IPerceptionDepthFrame, resultsLength uint32, results *Point)
	MapRegionOfPixelsToTargetAsync(region Rect, depthFrame *IPerceptionDepthFrame, targetCoordinatesLength uint32, targetCoordinates *Point) *IAsyncAction
	MapAllPixelsToTargetAsync(depthFrame *IPerceptionDepthFrame, targetCoordinatesLength uint32, targetCoordinates *Point) *IAsyncAction
}

type IPerceptionDepthCorrelatedCoordinateMapperVtbl struct {
	win32.IInspectableVtbl
	MapPixelToTarget               uintptr
	MapPixelsToTarget              uintptr
	MapRegionOfPixelsToTargetAsync uintptr
	MapAllPixelsToTargetAsync      uintptr
}

type IPerceptionDepthCorrelatedCoordinateMapper struct {
	win32.IInspectable
}

func (this *IPerceptionDepthCorrelatedCoordinateMapper) Vtbl() *IPerceptionDepthCorrelatedCoordinateMapperVtbl {
	return (*IPerceptionDepthCorrelatedCoordinateMapperVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionDepthCorrelatedCoordinateMapper) MapPixelToTarget(sourcePixelCoordinate Point, depthFrame *IPerceptionDepthFrame) Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MapPixelToTarget, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&sourcePixelCoordinate)), uintptr(unsafe.Pointer(depthFrame)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionDepthCorrelatedCoordinateMapper) MapPixelsToTarget(sourceCoordinatesLength uint32, sourceCoordinates *Point, depthFrame *IPerceptionDepthFrame, resultsLength uint32, results *Point) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MapPixelsToTarget, uintptr(unsafe.Pointer(this)), uintptr(sourceCoordinatesLength), uintptr(unsafe.Pointer(sourceCoordinates)), uintptr(unsafe.Pointer(depthFrame)), uintptr(resultsLength), uintptr(unsafe.Pointer(results)))
	_ = _hr
}

func (this *IPerceptionDepthCorrelatedCoordinateMapper) MapRegionOfPixelsToTargetAsync(region Rect, depthFrame *IPerceptionDepthFrame, targetCoordinatesLength uint32, targetCoordinates *Point) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MapRegionOfPixelsToTargetAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&region)), uintptr(unsafe.Pointer(depthFrame)), uintptr(targetCoordinatesLength), uintptr(unsafe.Pointer(targetCoordinates)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionDepthCorrelatedCoordinateMapper) MapAllPixelsToTargetAsync(depthFrame *IPerceptionDepthFrame, targetCoordinatesLength uint32, targetCoordinates *Point) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MapAllPixelsToTargetAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(depthFrame)), uintptr(targetCoordinatesLength), uintptr(unsafe.Pointer(targetCoordinates)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A37B81FC-9906-4FFD-9161-0024B360B657
var IID_IPerceptionDepthFrame = syscall.GUID{0xA37B81FC, 0x9906, 0x4FFD,
	[8]byte{0x91, 0x61, 0x00, 0x24, 0xB3, 0x60, 0xB6, 0x57}}

type IPerceptionDepthFrameInterface interface {
	win32.IInspectableInterface
	Get_VideoFrame() *IVideoFrame
}

type IPerceptionDepthFrameVtbl struct {
	win32.IInspectableVtbl
	Get_VideoFrame uintptr
}

type IPerceptionDepthFrame struct {
	win32.IInspectable
}

func (this *IPerceptionDepthFrame) Vtbl() *IPerceptionDepthFrameVtbl {
	return (*IPerceptionDepthFrameVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionDepthFrame) Get_VideoFrame() *IVideoFrame {
	var _result *IVideoFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 443D25B2-B282-4637-9173-AC978435C985
var IID_IPerceptionDepthFrameArrivedEventArgs = syscall.GUID{0x443D25B2, 0xB282, 0x4637,
	[8]byte{0x91, 0x73, 0xAC, 0x97, 0x84, 0x35, 0xC9, 0x85}}

type IPerceptionDepthFrameArrivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_RelativeTime() TimeSpan
	TryOpenFrame() *IPerceptionDepthFrame
}

type IPerceptionDepthFrameArrivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_RelativeTime uintptr
	TryOpenFrame     uintptr
}

type IPerceptionDepthFrameArrivedEventArgs struct {
	win32.IInspectable
}

func (this *IPerceptionDepthFrameArrivedEventArgs) Vtbl() *IPerceptionDepthFrameArrivedEventArgsVtbl {
	return (*IPerceptionDepthFrameArrivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionDepthFrameArrivedEventArgs) Get_RelativeTime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RelativeTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionDepthFrameArrivedEventArgs) TryOpenFrame() *IPerceptionDepthFrame {
	var _result *IPerceptionDepthFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryOpenFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B1A3C09F-299B-4612-A4F7-270F25A096EC
var IID_IPerceptionDepthFrameReader = syscall.GUID{0xB1A3C09F, 0x299B, 0x4612,
	[8]byte{0xA4, 0xF7, 0x27, 0x0F, 0x25, 0xA0, 0x96, 0xEC}}

type IPerceptionDepthFrameReaderInterface interface {
	win32.IInspectableInterface
	Add_FrameArrived(handler TypedEventHandler[*IPerceptionDepthFrameReader, *IPerceptionDepthFrameArrivedEventArgs]) EventRegistrationToken
	Remove_FrameArrived(token EventRegistrationToken)
	Get_Source() *IPerceptionDepthFrameSource
	Get_IsPaused() bool
	Put_IsPaused(value bool)
	TryReadLatestFrame() *IPerceptionDepthFrame
}

type IPerceptionDepthFrameReaderVtbl struct {
	win32.IInspectableVtbl
	Add_FrameArrived    uintptr
	Remove_FrameArrived uintptr
	Get_Source          uintptr
	Get_IsPaused        uintptr
	Put_IsPaused        uintptr
	TryReadLatestFrame  uintptr
}

type IPerceptionDepthFrameReader struct {
	win32.IInspectable
}

func (this *IPerceptionDepthFrameReader) Vtbl() *IPerceptionDepthFrameReaderVtbl {
	return (*IPerceptionDepthFrameReaderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionDepthFrameReader) Add_FrameArrived(handler TypedEventHandler[*IPerceptionDepthFrameReader, *IPerceptionDepthFrameArrivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_FrameArrived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionDepthFrameReader) Remove_FrameArrived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_FrameArrived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionDepthFrameReader) Get_Source() *IPerceptionDepthFrameSource {
	var _result *IPerceptionDepthFrameSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Source, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionDepthFrameReader) Get_IsPaused() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPaused, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionDepthFrameReader) Put_IsPaused(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsPaused, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPerceptionDepthFrameReader) TryReadLatestFrame() *IPerceptionDepthFrame {
	var _result *IPerceptionDepthFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryReadLatestFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 79D433D6-47FB-4DF1-BFC9-F01D40BD9942
var IID_IPerceptionDepthFrameSource = syscall.GUID{0x79D433D6, 0x47FB, 0x4DF1,
	[8]byte{0xBF, 0xC9, 0xF0, 0x1D, 0x40, 0xBD, 0x99, 0x42}}

type IPerceptionDepthFrameSourceInterface interface {
	win32.IInspectableInterface
	Add_AvailableChanged(handler TypedEventHandler[*IPerceptionDepthFrameSource, interface{}]) EventRegistrationToken
	Remove_AvailableChanged(token EventRegistrationToken)
	Add_ActiveChanged(handler TypedEventHandler[*IPerceptionDepthFrameSource, interface{}]) EventRegistrationToken
	Remove_ActiveChanged(token EventRegistrationToken)
	Add_PropertiesChanged(handler TypedEventHandler[*IPerceptionDepthFrameSource, *IPerceptionFrameSourcePropertiesChangedEventArgs]) EventRegistrationToken
	Remove_PropertiesChanged(token EventRegistrationToken)
	Add_VideoProfileChanged(handler TypedEventHandler[*IPerceptionDepthFrameSource, interface{}]) EventRegistrationToken
	Remove_VideoProfileChanged(token EventRegistrationToken)
	Add_CameraIntrinsicsChanged(handler TypedEventHandler[*IPerceptionDepthFrameSource, interface{}]) EventRegistrationToken
	Remove_CameraIntrinsicsChanged(token EventRegistrationToken)
	Get_Id() string
	Get_DisplayName() string
	Get_DeviceKind() string
	Get_Available() bool
	Get_Active() bool
	Get_IsControlled() bool
	Get_Properties() *IMapView[string, interface{}]
	Get_SupportedVideoProfiles() *IVectorView[*IPerceptionVideoProfile]
	Get_AvailableVideoProfiles() *IVectorView[*IPerceptionVideoProfile]
	Get_VideoProfile() *IPerceptionVideoProfile
	Get_CameraIntrinsics() *ICameraIntrinsics
	AcquireControlSession() *IPerceptionControlSession
	CanControlIndependentlyFrom(targetId string) bool
	IsCorrelatedWith(targetId string) bool
	TryGetTransformTo(targetId string, result Matrix4x4) bool
	TryGetDepthCorrelatedCameraIntrinsicsAsync(target *IPerceptionDepthFrameSource) *IAsyncOperation[*IPerceptionDepthCorrelatedCameraIntrinsics]
	TryGetDepthCorrelatedCoordinateMapperAsync(targetId string, depthFrameSourceToMapWith *IPerceptionDepthFrameSource) *IAsyncOperation[*IPerceptionDepthCorrelatedCoordinateMapper]
	TrySetVideoProfileAsync(controlSession *IPerceptionControlSession, profile *IPerceptionVideoProfile) *IAsyncOperation[*IPerceptionFrameSourcePropertyChangeResult]
	OpenReader() *IPerceptionDepthFrameReader
}

type IPerceptionDepthFrameSourceVtbl struct {
	win32.IInspectableVtbl
	Add_AvailableChanged                       uintptr
	Remove_AvailableChanged                    uintptr
	Add_ActiveChanged                          uintptr
	Remove_ActiveChanged                       uintptr
	Add_PropertiesChanged                      uintptr
	Remove_PropertiesChanged                   uintptr
	Add_VideoProfileChanged                    uintptr
	Remove_VideoProfileChanged                 uintptr
	Add_CameraIntrinsicsChanged                uintptr
	Remove_CameraIntrinsicsChanged             uintptr
	Get_Id                                     uintptr
	Get_DisplayName                            uintptr
	Get_DeviceKind                             uintptr
	Get_Available                              uintptr
	Get_Active                                 uintptr
	Get_IsControlled                           uintptr
	Get_Properties                             uintptr
	Get_SupportedVideoProfiles                 uintptr
	Get_AvailableVideoProfiles                 uintptr
	Get_VideoProfile                           uintptr
	Get_CameraIntrinsics                       uintptr
	AcquireControlSession                      uintptr
	CanControlIndependentlyFrom                uintptr
	IsCorrelatedWith                           uintptr
	TryGetTransformTo                          uintptr
	TryGetDepthCorrelatedCameraIntrinsicsAsync uintptr
	TryGetDepthCorrelatedCoordinateMapperAsync uintptr
	TrySetVideoProfileAsync                    uintptr
	OpenReader                                 uintptr
}

type IPerceptionDepthFrameSource struct {
	win32.IInspectable
}

func (this *IPerceptionDepthFrameSource) Vtbl() *IPerceptionDepthFrameSourceVtbl {
	return (*IPerceptionDepthFrameSourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionDepthFrameSource) Add_AvailableChanged(handler TypedEventHandler[*IPerceptionDepthFrameSource, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AvailableChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionDepthFrameSource) Remove_AvailableChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AvailableChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionDepthFrameSource) Add_ActiveChanged(handler TypedEventHandler[*IPerceptionDepthFrameSource, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ActiveChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionDepthFrameSource) Remove_ActiveChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ActiveChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionDepthFrameSource) Add_PropertiesChanged(handler TypedEventHandler[*IPerceptionDepthFrameSource, *IPerceptionFrameSourcePropertiesChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PropertiesChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionDepthFrameSource) Remove_PropertiesChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PropertiesChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionDepthFrameSource) Add_VideoProfileChanged(handler TypedEventHandler[*IPerceptionDepthFrameSource, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_VideoProfileChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionDepthFrameSource) Remove_VideoProfileChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_VideoProfileChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionDepthFrameSource) Add_CameraIntrinsicsChanged(handler TypedEventHandler[*IPerceptionDepthFrameSource, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_CameraIntrinsicsChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionDepthFrameSource) Remove_CameraIntrinsicsChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_CameraIntrinsicsChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionDepthFrameSource) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPerceptionDepthFrameSource) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPerceptionDepthFrameSource) Get_DeviceKind() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPerceptionDepthFrameSource) Get_Available() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Available, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionDepthFrameSource) Get_Active() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Active, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionDepthFrameSource) Get_IsControlled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsControlled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionDepthFrameSource) Get_Properties() *IMapView[string, interface{}] {
	var _result *IMapView[string, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionDepthFrameSource) Get_SupportedVideoProfiles() *IVectorView[*IPerceptionVideoProfile] {
	var _result *IVectorView[*IPerceptionVideoProfile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedVideoProfiles, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionDepthFrameSource) Get_AvailableVideoProfiles() *IVectorView[*IPerceptionVideoProfile] {
	var _result *IVectorView[*IPerceptionVideoProfile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AvailableVideoProfiles, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionDepthFrameSource) Get_VideoProfile() *IPerceptionVideoProfile {
	var _result *IPerceptionVideoProfile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoProfile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionDepthFrameSource) Get_CameraIntrinsics() *ICameraIntrinsics {
	var _result *ICameraIntrinsics
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CameraIntrinsics, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionDepthFrameSource) AcquireControlSession() *IPerceptionControlSession {
	var _result *IPerceptionControlSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AcquireControlSession, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionDepthFrameSource) CanControlIndependentlyFrom(targetId string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CanControlIndependentlyFrom, uintptr(unsafe.Pointer(this)), NewHStr(targetId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionDepthFrameSource) IsCorrelatedWith(targetId string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsCorrelatedWith, uintptr(unsafe.Pointer(this)), NewHStr(targetId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionDepthFrameSource) TryGetTransformTo(targetId string, result Matrix4x4) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetTransformTo, uintptr(unsafe.Pointer(this)), NewHStr(targetId).Ptr, uintptr(unsafe.Pointer(&result)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionDepthFrameSource) TryGetDepthCorrelatedCameraIntrinsicsAsync(target *IPerceptionDepthFrameSource) *IAsyncOperation[*IPerceptionDepthCorrelatedCameraIntrinsics] {
	var _result *IAsyncOperation[*IPerceptionDepthCorrelatedCameraIntrinsics]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetDepthCorrelatedCameraIntrinsicsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(target)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionDepthFrameSource) TryGetDepthCorrelatedCoordinateMapperAsync(targetId string, depthFrameSourceToMapWith *IPerceptionDepthFrameSource) *IAsyncOperation[*IPerceptionDepthCorrelatedCoordinateMapper] {
	var _result *IAsyncOperation[*IPerceptionDepthCorrelatedCoordinateMapper]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetDepthCorrelatedCoordinateMapperAsync, uintptr(unsafe.Pointer(this)), NewHStr(targetId).Ptr, uintptr(unsafe.Pointer(depthFrameSourceToMapWith)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionDepthFrameSource) TrySetVideoProfileAsync(controlSession *IPerceptionControlSession, profile *IPerceptionVideoProfile) *IAsyncOperation[*IPerceptionFrameSourcePropertyChangeResult] {
	var _result *IAsyncOperation[*IPerceptionFrameSourcePropertyChangeResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TrySetVideoProfileAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(controlSession)), uintptr(unsafe.Pointer(profile)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionDepthFrameSource) OpenReader() *IPerceptionDepthFrameReader {
	var _result *IPerceptionDepthFrameReader
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenReader, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E3D23D2E-6E2C-4E6D-91D9-704CD8DFF79D
var IID_IPerceptionDepthFrameSource2 = syscall.GUID{0xE3D23D2E, 0x6E2C, 0x4E6D,
	[8]byte{0x91, 0xD9, 0x70, 0x4C, 0xD8, 0xDF, 0xF7, 0x9D}}

type IPerceptionDepthFrameSource2Interface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
}

type IPerceptionDepthFrameSource2Vtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId uintptr
}

type IPerceptionDepthFrameSource2 struct {
	win32.IInspectable
}

func (this *IPerceptionDepthFrameSource2) Vtbl() *IPerceptionDepthFrameSource2Vtbl {
	return (*IPerceptionDepthFrameSource2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionDepthFrameSource2) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 93A48168-8BF8-45D2-A2F8-4AC0931CC7A6
var IID_IPerceptionDepthFrameSourceAddedEventArgs = syscall.GUID{0x93A48168, 0x8BF8, 0x45D2,
	[8]byte{0xA2, 0xF8, 0x4A, 0xC0, 0x93, 0x1C, 0xC7, 0xA6}}

type IPerceptionDepthFrameSourceAddedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_FrameSource() *IPerceptionDepthFrameSource
}

type IPerceptionDepthFrameSourceAddedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_FrameSource uintptr
}

type IPerceptionDepthFrameSourceAddedEventArgs struct {
	win32.IInspectable
}

func (this *IPerceptionDepthFrameSourceAddedEventArgs) Vtbl() *IPerceptionDepthFrameSourceAddedEventArgsVtbl {
	return (*IPerceptionDepthFrameSourceAddedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionDepthFrameSourceAddedEventArgs) Get_FrameSource() *IPerceptionDepthFrameSource {
	var _result *IPerceptionDepthFrameSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A0C0CC4D-E96C-4D81-86DD-38B95E49C6DF
var IID_IPerceptionDepthFrameSourceRemovedEventArgs = syscall.GUID{0xA0C0CC4D, 0xE96C, 0x4D81,
	[8]byte{0x86, 0xDD, 0x38, 0xB9, 0x5E, 0x49, 0xC6, 0xDF}}

type IPerceptionDepthFrameSourceRemovedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_FrameSource() *IPerceptionDepthFrameSource
}

type IPerceptionDepthFrameSourceRemovedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_FrameSource uintptr
}

type IPerceptionDepthFrameSourceRemovedEventArgs struct {
	win32.IInspectable
}

func (this *IPerceptionDepthFrameSourceRemovedEventArgs) Vtbl() *IPerceptionDepthFrameSourceRemovedEventArgsVtbl {
	return (*IPerceptionDepthFrameSourceRemovedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionDepthFrameSourceRemovedEventArgs) Get_FrameSource() *IPerceptionDepthFrameSource {
	var _result *IPerceptionDepthFrameSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5DF3CCA2-01F8-4A87-B859-D5E5B7E1DE48
var IID_IPerceptionDepthFrameSourceStatics = syscall.GUID{0x5DF3CCA2, 0x01F8, 0x4A87,
	[8]byte{0xB8, 0x59, 0xD5, 0xE5, 0xB7, 0xE1, 0xDE, 0x48}}

type IPerceptionDepthFrameSourceStaticsInterface interface {
	win32.IInspectableInterface
	CreateWatcher() *IPerceptionDepthFrameSourceWatcher
	FindAllAsync() *IAsyncOperation[*IVectorView[*IPerceptionDepthFrameSource]]
	FromIdAsync(id string) *IAsyncOperation[*IPerceptionDepthFrameSource]
	RequestAccessAsync() *IAsyncOperation[PerceptionFrameSourceAccessStatus]
}

type IPerceptionDepthFrameSourceStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateWatcher      uintptr
	FindAllAsync       uintptr
	FromIdAsync        uintptr
	RequestAccessAsync uintptr
}

type IPerceptionDepthFrameSourceStatics struct {
	win32.IInspectable
}

func (this *IPerceptionDepthFrameSourceStatics) Vtbl() *IPerceptionDepthFrameSourceStaticsVtbl {
	return (*IPerceptionDepthFrameSourceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionDepthFrameSourceStatics) CreateWatcher() *IPerceptionDepthFrameSourceWatcher {
	var _result *IPerceptionDepthFrameSourceWatcher
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWatcher, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionDepthFrameSourceStatics) FindAllAsync() *IAsyncOperation[*IVectorView[*IPerceptionDepthFrameSource]] {
	var _result *IAsyncOperation[*IVectorView[*IPerceptionDepthFrameSource]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionDepthFrameSourceStatics) FromIdAsync(id string) *IAsyncOperation[*IPerceptionDepthFrameSource] {
	var _result *IAsyncOperation[*IPerceptionDepthFrameSource]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(id).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionDepthFrameSourceStatics) RequestAccessAsync() *IAsyncOperation[PerceptionFrameSourceAccessStatus] {
	var _result *IAsyncOperation[PerceptionFrameSourceAccessStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 780E96D1-8D02-4D2B-ADA4-5BA624A0EB10
var IID_IPerceptionDepthFrameSourceWatcher = syscall.GUID{0x780E96D1, 0x8D02, 0x4D2B,
	[8]byte{0xAD, 0xA4, 0x5B, 0xA6, 0x24, 0xA0, 0xEB, 0x10}}

type IPerceptionDepthFrameSourceWatcherInterface interface {
	win32.IInspectableInterface
	Add_SourceAdded(handler TypedEventHandler[*IPerceptionDepthFrameSourceWatcher, *IPerceptionDepthFrameSourceAddedEventArgs]) EventRegistrationToken
	Remove_SourceAdded(token EventRegistrationToken)
	Add_SourceRemoved(handler TypedEventHandler[*IPerceptionDepthFrameSourceWatcher, *IPerceptionDepthFrameSourceRemovedEventArgs]) EventRegistrationToken
	Remove_SourceRemoved(token EventRegistrationToken)
	Add_Stopped(handler TypedEventHandler[*IPerceptionDepthFrameSourceWatcher, interface{}]) EventRegistrationToken
	Remove_Stopped(token EventRegistrationToken)
	Add_EnumerationCompleted(handler TypedEventHandler[*IPerceptionDepthFrameSourceWatcher, interface{}]) EventRegistrationToken
	Remove_EnumerationCompleted(token EventRegistrationToken)
	Get_Status() DeviceWatcherStatus
	Start()
	Stop()
}

type IPerceptionDepthFrameSourceWatcherVtbl struct {
	win32.IInspectableVtbl
	Add_SourceAdded             uintptr
	Remove_SourceAdded          uintptr
	Add_SourceRemoved           uintptr
	Remove_SourceRemoved        uintptr
	Add_Stopped                 uintptr
	Remove_Stopped              uintptr
	Add_EnumerationCompleted    uintptr
	Remove_EnumerationCompleted uintptr
	Get_Status                  uintptr
	Start                       uintptr
	Stop                        uintptr
}

type IPerceptionDepthFrameSourceWatcher struct {
	win32.IInspectable
}

func (this *IPerceptionDepthFrameSourceWatcher) Vtbl() *IPerceptionDepthFrameSourceWatcherVtbl {
	return (*IPerceptionDepthFrameSourceWatcherVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionDepthFrameSourceWatcher) Add_SourceAdded(handler TypedEventHandler[*IPerceptionDepthFrameSourceWatcher, *IPerceptionDepthFrameSourceAddedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SourceAdded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionDepthFrameSourceWatcher) Remove_SourceAdded(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SourceAdded, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionDepthFrameSourceWatcher) Add_SourceRemoved(handler TypedEventHandler[*IPerceptionDepthFrameSourceWatcher, *IPerceptionDepthFrameSourceRemovedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SourceRemoved, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionDepthFrameSourceWatcher) Remove_SourceRemoved(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SourceRemoved, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionDepthFrameSourceWatcher) Add_Stopped(handler TypedEventHandler[*IPerceptionDepthFrameSourceWatcher, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Stopped, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionDepthFrameSourceWatcher) Remove_Stopped(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Stopped, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionDepthFrameSourceWatcher) Add_EnumerationCompleted(handler TypedEventHandler[*IPerceptionDepthFrameSourceWatcher, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_EnumerationCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionDepthFrameSourceWatcher) Remove_EnumerationCompleted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_EnumerationCompleted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionDepthFrameSourceWatcher) Get_Status() DeviceWatcherStatus {
	var _result DeviceWatcherStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionDepthFrameSourceWatcher) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IPerceptionDepthFrameSourceWatcher) Stop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 6C68E068-BCF1-4ECC-B891-7625D1244B6B
var IID_IPerceptionFrameSourcePropertiesChangedEventArgs = syscall.GUID{0x6C68E068, 0xBCF1, 0x4ECC,
	[8]byte{0xB8, 0x91, 0x76, 0x25, 0xD1, 0x24, 0x4B, 0x6B}}

type IPerceptionFrameSourcePropertiesChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_CollectionChange() CollectionChange
	Get_Key() string
}

type IPerceptionFrameSourcePropertiesChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_CollectionChange uintptr
	Get_Key              uintptr
}

type IPerceptionFrameSourcePropertiesChangedEventArgs struct {
	win32.IInspectable
}

func (this *IPerceptionFrameSourcePropertiesChangedEventArgs) Vtbl() *IPerceptionFrameSourcePropertiesChangedEventArgsVtbl {
	return (*IPerceptionFrameSourcePropertiesChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionFrameSourcePropertiesChangedEventArgs) Get_CollectionChange() CollectionChange {
	var _result CollectionChange
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CollectionChange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionFrameSourcePropertiesChangedEventArgs) Get_Key() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Key, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 1E33390A-3C90-4D22-B898-F42BBA6418FF
var IID_IPerceptionFrameSourcePropertyChangeResult = syscall.GUID{0x1E33390A, 0x3C90, 0x4D22,
	[8]byte{0xB8, 0x98, 0xF4, 0x2B, 0xBA, 0x64, 0x18, 0xFF}}

type IPerceptionFrameSourcePropertyChangeResultInterface interface {
	win32.IInspectableInterface
	Get_Status() PerceptionFrameSourcePropertyChangeStatus
	Get_NewValue() interface{}
}

type IPerceptionFrameSourcePropertyChangeResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status   uintptr
	Get_NewValue uintptr
}

type IPerceptionFrameSourcePropertyChangeResult struct {
	win32.IInspectable
}

func (this *IPerceptionFrameSourcePropertyChangeResult) Vtbl() *IPerceptionFrameSourcePropertyChangeResultVtbl {
	return (*IPerceptionFrameSourcePropertyChangeResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionFrameSourcePropertyChangeResult) Get_Status() PerceptionFrameSourcePropertyChangeStatus {
	var _result PerceptionFrameSourcePropertyChangeStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionFrameSourcePropertyChangeResult) Get_NewValue() interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NewValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// B0886276-849E-4C7A-8AE6-B56064532153
var IID_IPerceptionInfraredFrame = syscall.GUID{0xB0886276, 0x849E, 0x4C7A,
	[8]byte{0x8A, 0xE6, 0xB5, 0x60, 0x64, 0x53, 0x21, 0x53}}

type IPerceptionInfraredFrameInterface interface {
	win32.IInspectableInterface
	Get_VideoFrame() *IVideoFrame
}

type IPerceptionInfraredFrameVtbl struct {
	win32.IInspectableVtbl
	Get_VideoFrame uintptr
}

type IPerceptionInfraredFrame struct {
	win32.IInspectable
}

func (this *IPerceptionInfraredFrame) Vtbl() *IPerceptionInfraredFrameVtbl {
	return (*IPerceptionInfraredFrameVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionInfraredFrame) Get_VideoFrame() *IVideoFrame {
	var _result *IVideoFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9F77FAC7-B4BD-4857-9D50-BE8EF075DAEF
var IID_IPerceptionInfraredFrameArrivedEventArgs = syscall.GUID{0x9F77FAC7, 0xB4BD, 0x4857,
	[8]byte{0x9D, 0x50, 0xBE, 0x8E, 0xF0, 0x75, 0xDA, 0xEF}}

type IPerceptionInfraredFrameArrivedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_RelativeTime() TimeSpan
	TryOpenFrame() *IPerceptionInfraredFrame
}

type IPerceptionInfraredFrameArrivedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_RelativeTime uintptr
	TryOpenFrame     uintptr
}

type IPerceptionInfraredFrameArrivedEventArgs struct {
	win32.IInspectable
}

func (this *IPerceptionInfraredFrameArrivedEventArgs) Vtbl() *IPerceptionInfraredFrameArrivedEventArgsVtbl {
	return (*IPerceptionInfraredFrameArrivedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionInfraredFrameArrivedEventArgs) Get_RelativeTime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RelativeTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionInfraredFrameArrivedEventArgs) TryOpenFrame() *IPerceptionInfraredFrame {
	var _result *IPerceptionInfraredFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryOpenFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7960CE18-D39B-4FC8-A04A-929734C6756C
var IID_IPerceptionInfraredFrameReader = syscall.GUID{0x7960CE18, 0xD39B, 0x4FC8,
	[8]byte{0xA0, 0x4A, 0x92, 0x97, 0x34, 0xC6, 0x75, 0x6C}}

type IPerceptionInfraredFrameReaderInterface interface {
	win32.IInspectableInterface
	Add_FrameArrived(handler TypedEventHandler[*IPerceptionInfraredFrameReader, *IPerceptionInfraredFrameArrivedEventArgs]) EventRegistrationToken
	Remove_FrameArrived(token EventRegistrationToken)
	Get_Source() *IPerceptionInfraredFrameSource
	Get_IsPaused() bool
	Put_IsPaused(value bool)
	TryReadLatestFrame() *IPerceptionInfraredFrame
}

type IPerceptionInfraredFrameReaderVtbl struct {
	win32.IInspectableVtbl
	Add_FrameArrived    uintptr
	Remove_FrameArrived uintptr
	Get_Source          uintptr
	Get_IsPaused        uintptr
	Put_IsPaused        uintptr
	TryReadLatestFrame  uintptr
}

type IPerceptionInfraredFrameReader struct {
	win32.IInspectable
}

func (this *IPerceptionInfraredFrameReader) Vtbl() *IPerceptionInfraredFrameReaderVtbl {
	return (*IPerceptionInfraredFrameReaderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionInfraredFrameReader) Add_FrameArrived(handler TypedEventHandler[*IPerceptionInfraredFrameReader, *IPerceptionInfraredFrameArrivedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_FrameArrived, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionInfraredFrameReader) Remove_FrameArrived(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_FrameArrived, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionInfraredFrameReader) Get_Source() *IPerceptionInfraredFrameSource {
	var _result *IPerceptionInfraredFrameSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Source, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionInfraredFrameReader) Get_IsPaused() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPaused, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionInfraredFrameReader) Put_IsPaused(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsPaused, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IPerceptionInfraredFrameReader) TryReadLatestFrame() *IPerceptionInfraredFrame {
	var _result *IPerceptionInfraredFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryReadLatestFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 55B08742-1808-494E-9E30-9D2A7BE8F700
var IID_IPerceptionInfraredFrameSource = syscall.GUID{0x55B08742, 0x1808, 0x494E,
	[8]byte{0x9E, 0x30, 0x9D, 0x2A, 0x7B, 0xE8, 0xF7, 0x00}}

type IPerceptionInfraredFrameSourceInterface interface {
	win32.IInspectableInterface
	Add_AvailableChanged(handler TypedEventHandler[*IPerceptionInfraredFrameSource, interface{}]) EventRegistrationToken
	Remove_AvailableChanged(token EventRegistrationToken)
	Add_ActiveChanged(handler TypedEventHandler[*IPerceptionInfraredFrameSource, interface{}]) EventRegistrationToken
	Remove_ActiveChanged(token EventRegistrationToken)
	Add_PropertiesChanged(handler TypedEventHandler[*IPerceptionInfraredFrameSource, *IPerceptionFrameSourcePropertiesChangedEventArgs]) EventRegistrationToken
	Remove_PropertiesChanged(token EventRegistrationToken)
	Add_VideoProfileChanged(handler TypedEventHandler[*IPerceptionInfraredFrameSource, interface{}]) EventRegistrationToken
	Remove_VideoProfileChanged(token EventRegistrationToken)
	Add_CameraIntrinsicsChanged(handler TypedEventHandler[*IPerceptionInfraredFrameSource, interface{}]) EventRegistrationToken
	Remove_CameraIntrinsicsChanged(token EventRegistrationToken)
	Get_Id() string
	Get_DisplayName() string
	Get_DeviceKind() string
	Get_Available() bool
	Get_Active() bool
	Get_IsControlled() bool
	Get_Properties() *IMapView[string, interface{}]
	Get_SupportedVideoProfiles() *IVectorView[*IPerceptionVideoProfile]
	Get_AvailableVideoProfiles() *IVectorView[*IPerceptionVideoProfile]
	Get_VideoProfile() *IPerceptionVideoProfile
	Get_CameraIntrinsics() *ICameraIntrinsics
	AcquireControlSession() *IPerceptionControlSession
	CanControlIndependentlyFrom(targetId string) bool
	IsCorrelatedWith(targetId string) bool
	TryGetTransformTo(targetId string, result Matrix4x4) bool
	TryGetDepthCorrelatedCameraIntrinsicsAsync(target *IPerceptionDepthFrameSource) *IAsyncOperation[*IPerceptionDepthCorrelatedCameraIntrinsics]
	TryGetDepthCorrelatedCoordinateMapperAsync(targetId string, depthFrameSourceToMapWith *IPerceptionDepthFrameSource) *IAsyncOperation[*IPerceptionDepthCorrelatedCoordinateMapper]
	TrySetVideoProfileAsync(controlSession *IPerceptionControlSession, profile *IPerceptionVideoProfile) *IAsyncOperation[*IPerceptionFrameSourcePropertyChangeResult]
	OpenReader() *IPerceptionInfraredFrameReader
}

type IPerceptionInfraredFrameSourceVtbl struct {
	win32.IInspectableVtbl
	Add_AvailableChanged                       uintptr
	Remove_AvailableChanged                    uintptr
	Add_ActiveChanged                          uintptr
	Remove_ActiveChanged                       uintptr
	Add_PropertiesChanged                      uintptr
	Remove_PropertiesChanged                   uintptr
	Add_VideoProfileChanged                    uintptr
	Remove_VideoProfileChanged                 uintptr
	Add_CameraIntrinsicsChanged                uintptr
	Remove_CameraIntrinsicsChanged             uintptr
	Get_Id                                     uintptr
	Get_DisplayName                            uintptr
	Get_DeviceKind                             uintptr
	Get_Available                              uintptr
	Get_Active                                 uintptr
	Get_IsControlled                           uintptr
	Get_Properties                             uintptr
	Get_SupportedVideoProfiles                 uintptr
	Get_AvailableVideoProfiles                 uintptr
	Get_VideoProfile                           uintptr
	Get_CameraIntrinsics                       uintptr
	AcquireControlSession                      uintptr
	CanControlIndependentlyFrom                uintptr
	IsCorrelatedWith                           uintptr
	TryGetTransformTo                          uintptr
	TryGetDepthCorrelatedCameraIntrinsicsAsync uintptr
	TryGetDepthCorrelatedCoordinateMapperAsync uintptr
	TrySetVideoProfileAsync                    uintptr
	OpenReader                                 uintptr
}

type IPerceptionInfraredFrameSource struct {
	win32.IInspectable
}

func (this *IPerceptionInfraredFrameSource) Vtbl() *IPerceptionInfraredFrameSourceVtbl {
	return (*IPerceptionInfraredFrameSourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionInfraredFrameSource) Add_AvailableChanged(handler TypedEventHandler[*IPerceptionInfraredFrameSource, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_AvailableChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionInfraredFrameSource) Remove_AvailableChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_AvailableChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionInfraredFrameSource) Add_ActiveChanged(handler TypedEventHandler[*IPerceptionInfraredFrameSource, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ActiveChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionInfraredFrameSource) Remove_ActiveChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ActiveChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionInfraredFrameSource) Add_PropertiesChanged(handler TypedEventHandler[*IPerceptionInfraredFrameSource, *IPerceptionFrameSourcePropertiesChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PropertiesChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionInfraredFrameSource) Remove_PropertiesChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PropertiesChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionInfraredFrameSource) Add_VideoProfileChanged(handler TypedEventHandler[*IPerceptionInfraredFrameSource, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_VideoProfileChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionInfraredFrameSource) Remove_VideoProfileChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_VideoProfileChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionInfraredFrameSource) Add_CameraIntrinsicsChanged(handler TypedEventHandler[*IPerceptionInfraredFrameSource, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_CameraIntrinsicsChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionInfraredFrameSource) Remove_CameraIntrinsicsChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_CameraIntrinsicsChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionInfraredFrameSource) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPerceptionInfraredFrameSource) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPerceptionInfraredFrameSource) Get_DeviceKind() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPerceptionInfraredFrameSource) Get_Available() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Available, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionInfraredFrameSource) Get_Active() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Active, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionInfraredFrameSource) Get_IsControlled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsControlled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionInfraredFrameSource) Get_Properties() *IMapView[string, interface{}] {
	var _result *IMapView[string, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionInfraredFrameSource) Get_SupportedVideoProfiles() *IVectorView[*IPerceptionVideoProfile] {
	var _result *IVectorView[*IPerceptionVideoProfile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedVideoProfiles, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionInfraredFrameSource) Get_AvailableVideoProfiles() *IVectorView[*IPerceptionVideoProfile] {
	var _result *IVectorView[*IPerceptionVideoProfile]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AvailableVideoProfiles, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionInfraredFrameSource) Get_VideoProfile() *IPerceptionVideoProfile {
	var _result *IPerceptionVideoProfile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoProfile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionInfraredFrameSource) Get_CameraIntrinsics() *ICameraIntrinsics {
	var _result *ICameraIntrinsics
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CameraIntrinsics, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionInfraredFrameSource) AcquireControlSession() *IPerceptionControlSession {
	var _result *IPerceptionControlSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AcquireControlSession, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionInfraredFrameSource) CanControlIndependentlyFrom(targetId string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CanControlIndependentlyFrom, uintptr(unsafe.Pointer(this)), NewHStr(targetId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionInfraredFrameSource) IsCorrelatedWith(targetId string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsCorrelatedWith, uintptr(unsafe.Pointer(this)), NewHStr(targetId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionInfraredFrameSource) TryGetTransformTo(targetId string, result Matrix4x4) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetTransformTo, uintptr(unsafe.Pointer(this)), NewHStr(targetId).Ptr, uintptr(unsafe.Pointer(&result)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionInfraredFrameSource) TryGetDepthCorrelatedCameraIntrinsicsAsync(target *IPerceptionDepthFrameSource) *IAsyncOperation[*IPerceptionDepthCorrelatedCameraIntrinsics] {
	var _result *IAsyncOperation[*IPerceptionDepthCorrelatedCameraIntrinsics]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetDepthCorrelatedCameraIntrinsicsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(target)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionInfraredFrameSource) TryGetDepthCorrelatedCoordinateMapperAsync(targetId string, depthFrameSourceToMapWith *IPerceptionDepthFrameSource) *IAsyncOperation[*IPerceptionDepthCorrelatedCoordinateMapper] {
	var _result *IAsyncOperation[*IPerceptionDepthCorrelatedCoordinateMapper]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetDepthCorrelatedCoordinateMapperAsync, uintptr(unsafe.Pointer(this)), NewHStr(targetId).Ptr, uintptr(unsafe.Pointer(depthFrameSourceToMapWith)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionInfraredFrameSource) TrySetVideoProfileAsync(controlSession *IPerceptionControlSession, profile *IPerceptionVideoProfile) *IAsyncOperation[*IPerceptionFrameSourcePropertyChangeResult] {
	var _result *IAsyncOperation[*IPerceptionFrameSourcePropertyChangeResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TrySetVideoProfileAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(controlSession)), uintptr(unsafe.Pointer(profile)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionInfraredFrameSource) OpenReader() *IPerceptionInfraredFrameReader {
	var _result *IPerceptionInfraredFrameReader
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenReader, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DCD4D798-4B0B-4300-8D85-410817FAA032
var IID_IPerceptionInfraredFrameSource2 = syscall.GUID{0xDCD4D798, 0x4B0B, 0x4300,
	[8]byte{0x8D, 0x85, 0x41, 0x08, 0x17, 0xFA, 0xA0, 0x32}}

type IPerceptionInfraredFrameSource2Interface interface {
	win32.IInspectableInterface
	Get_DeviceId() string
}

type IPerceptionInfraredFrameSource2Vtbl struct {
	win32.IInspectableVtbl
	Get_DeviceId uintptr
}

type IPerceptionInfraredFrameSource2 struct {
	win32.IInspectable
}

func (this *IPerceptionInfraredFrameSource2) Vtbl() *IPerceptionInfraredFrameSource2Vtbl {
	return (*IPerceptionInfraredFrameSource2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionInfraredFrameSource2) Get_DeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 6D334120-95CE-4660-907A-D98035AA2B7C
var IID_IPerceptionInfraredFrameSourceAddedEventArgs = syscall.GUID{0x6D334120, 0x95CE, 0x4660,
	[8]byte{0x90, 0x7A, 0xD9, 0x80, 0x35, 0xAA, 0x2B, 0x7C}}

type IPerceptionInfraredFrameSourceAddedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_FrameSource() *IPerceptionInfraredFrameSource
}

type IPerceptionInfraredFrameSourceAddedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_FrameSource uintptr
}

type IPerceptionInfraredFrameSourceAddedEventArgs struct {
	win32.IInspectable
}

func (this *IPerceptionInfraredFrameSourceAddedEventArgs) Vtbl() *IPerceptionInfraredFrameSourceAddedEventArgsVtbl {
	return (*IPerceptionInfraredFrameSourceAddedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionInfraredFrameSourceAddedEventArgs) Get_FrameSource() *IPerceptionInfraredFrameSource {
	var _result *IPerceptionInfraredFrameSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EA1A8071-7A70-4A61-AF94-07303853F695
var IID_IPerceptionInfraredFrameSourceRemovedEventArgs = syscall.GUID{0xEA1A8071, 0x7A70, 0x4A61,
	[8]byte{0xAF, 0x94, 0x07, 0x30, 0x38, 0x53, 0xF6, 0x95}}

type IPerceptionInfraredFrameSourceRemovedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_FrameSource() *IPerceptionInfraredFrameSource
}

type IPerceptionInfraredFrameSourceRemovedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_FrameSource uintptr
}

type IPerceptionInfraredFrameSourceRemovedEventArgs struct {
	win32.IInspectable
}

func (this *IPerceptionInfraredFrameSourceRemovedEventArgs) Vtbl() *IPerceptionInfraredFrameSourceRemovedEventArgsVtbl {
	return (*IPerceptionInfraredFrameSourceRemovedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionInfraredFrameSourceRemovedEventArgs) Get_FrameSource() *IPerceptionInfraredFrameSource {
	var _result *IPerceptionInfraredFrameSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameSource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5DF3CCA2-01F8-4A87-B859-D5E5B7E1DE47
var IID_IPerceptionInfraredFrameSourceStatics = syscall.GUID{0x5DF3CCA2, 0x01F8, 0x4A87,
	[8]byte{0xB8, 0x59, 0xD5, 0xE5, 0xB7, 0xE1, 0xDE, 0x47}}

type IPerceptionInfraredFrameSourceStaticsInterface interface {
	win32.IInspectableInterface
	CreateWatcher() *IPerceptionInfraredFrameSourceWatcher
	FindAllAsync() *IAsyncOperation[*IVectorView[*IPerceptionInfraredFrameSource]]
	FromIdAsync(id string) *IAsyncOperation[*IPerceptionInfraredFrameSource]
	RequestAccessAsync() *IAsyncOperation[PerceptionFrameSourceAccessStatus]
}

type IPerceptionInfraredFrameSourceStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateWatcher      uintptr
	FindAllAsync       uintptr
	FromIdAsync        uintptr
	RequestAccessAsync uintptr
}

type IPerceptionInfraredFrameSourceStatics struct {
	win32.IInspectable
}

func (this *IPerceptionInfraredFrameSourceStatics) Vtbl() *IPerceptionInfraredFrameSourceStaticsVtbl {
	return (*IPerceptionInfraredFrameSourceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionInfraredFrameSourceStatics) CreateWatcher() *IPerceptionInfraredFrameSourceWatcher {
	var _result *IPerceptionInfraredFrameSourceWatcher
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWatcher, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionInfraredFrameSourceStatics) FindAllAsync() *IAsyncOperation[*IVectorView[*IPerceptionInfraredFrameSource]] {
	var _result *IAsyncOperation[*IVectorView[*IPerceptionInfraredFrameSource]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindAllAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionInfraredFrameSourceStatics) FromIdAsync(id string) *IAsyncOperation[*IPerceptionInfraredFrameSource] {
	var _result *IAsyncOperation[*IPerceptionInfraredFrameSource]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromIdAsync, uintptr(unsafe.Pointer(this)), NewHStr(id).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPerceptionInfraredFrameSourceStatics) RequestAccessAsync() *IAsyncOperation[PerceptionFrameSourceAccessStatus] {
	var _result *IAsyncOperation[PerceptionFrameSourceAccessStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 383CFF99-D70C-444D-A8B0-720C2E66FE3B
var IID_IPerceptionInfraredFrameSourceWatcher = syscall.GUID{0x383CFF99, 0xD70C, 0x444D,
	[8]byte{0xA8, 0xB0, 0x72, 0x0C, 0x2E, 0x66, 0xFE, 0x3B}}

type IPerceptionInfraredFrameSourceWatcherInterface interface {
	win32.IInspectableInterface
	Add_SourceAdded(handler TypedEventHandler[*IPerceptionInfraredFrameSourceWatcher, *IPerceptionInfraredFrameSourceAddedEventArgs]) EventRegistrationToken
	Remove_SourceAdded(token EventRegistrationToken)
	Add_SourceRemoved(handler TypedEventHandler[*IPerceptionInfraredFrameSourceWatcher, *IPerceptionInfraredFrameSourceRemovedEventArgs]) EventRegistrationToken
	Remove_SourceRemoved(token EventRegistrationToken)
	Add_Stopped(handler TypedEventHandler[*IPerceptionInfraredFrameSourceWatcher, interface{}]) EventRegistrationToken
	Remove_Stopped(token EventRegistrationToken)
	Add_EnumerationCompleted(handler TypedEventHandler[*IPerceptionInfraredFrameSourceWatcher, interface{}]) EventRegistrationToken
	Remove_EnumerationCompleted(token EventRegistrationToken)
	Get_Status() DeviceWatcherStatus
	Start()
	Stop()
}

type IPerceptionInfraredFrameSourceWatcherVtbl struct {
	win32.IInspectableVtbl
	Add_SourceAdded             uintptr
	Remove_SourceAdded          uintptr
	Add_SourceRemoved           uintptr
	Remove_SourceRemoved        uintptr
	Add_Stopped                 uintptr
	Remove_Stopped              uintptr
	Add_EnumerationCompleted    uintptr
	Remove_EnumerationCompleted uintptr
	Get_Status                  uintptr
	Start                       uintptr
	Stop                        uintptr
}

type IPerceptionInfraredFrameSourceWatcher struct {
	win32.IInspectable
}

func (this *IPerceptionInfraredFrameSourceWatcher) Vtbl() *IPerceptionInfraredFrameSourceWatcherVtbl {
	return (*IPerceptionInfraredFrameSourceWatcherVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionInfraredFrameSourceWatcher) Add_SourceAdded(handler TypedEventHandler[*IPerceptionInfraredFrameSourceWatcher, *IPerceptionInfraredFrameSourceAddedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SourceAdded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionInfraredFrameSourceWatcher) Remove_SourceAdded(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SourceAdded, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionInfraredFrameSourceWatcher) Add_SourceRemoved(handler TypedEventHandler[*IPerceptionInfraredFrameSourceWatcher, *IPerceptionInfraredFrameSourceRemovedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_SourceRemoved, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionInfraredFrameSourceWatcher) Remove_SourceRemoved(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_SourceRemoved, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionInfraredFrameSourceWatcher) Add_Stopped(handler TypedEventHandler[*IPerceptionInfraredFrameSourceWatcher, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Stopped, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionInfraredFrameSourceWatcher) Remove_Stopped(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Stopped, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionInfraredFrameSourceWatcher) Add_EnumerationCompleted(handler TypedEventHandler[*IPerceptionInfraredFrameSourceWatcher, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_EnumerationCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionInfraredFrameSourceWatcher) Remove_EnumerationCompleted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_EnumerationCompleted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IPerceptionInfraredFrameSourceWatcher) Get_Status() DeviceWatcherStatus {
	var _result DeviceWatcherStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionInfraredFrameSourceWatcher) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IPerceptionInfraredFrameSourceWatcher) Stop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 75763EA3-011A-470E-8225-6F05ADE25648
var IID_IPerceptionVideoProfile = syscall.GUID{0x75763EA3, 0x011A, 0x470E,
	[8]byte{0x82, 0x25, 0x6F, 0x05, 0xAD, 0xE2, 0x56, 0x48}}

type IPerceptionVideoProfileInterface interface {
	win32.IInspectableInterface
	Get_BitmapPixelFormat() BitmapPixelFormat
	Get_BitmapAlphaMode() BitmapAlphaMode
	Get_Width() int32
	Get_Height() int32
	Get_FrameDuration() TimeSpan
	IsEqual(other *IPerceptionVideoProfile) bool
}

type IPerceptionVideoProfileVtbl struct {
	win32.IInspectableVtbl
	Get_BitmapPixelFormat uintptr
	Get_BitmapAlphaMode   uintptr
	Get_Width             uintptr
	Get_Height            uintptr
	Get_FrameDuration     uintptr
	IsEqual               uintptr
}

type IPerceptionVideoProfile struct {
	win32.IInspectable
}

func (this *IPerceptionVideoProfile) Vtbl() *IPerceptionVideoProfileVtbl {
	return (*IPerceptionVideoProfileVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerceptionVideoProfile) Get_BitmapPixelFormat() BitmapPixelFormat {
	var _result BitmapPixelFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BitmapPixelFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionVideoProfile) Get_BitmapAlphaMode() BitmapAlphaMode {
	var _result BitmapAlphaMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BitmapAlphaMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionVideoProfile) Get_Width() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Width, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionVideoProfile) Get_Height() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Height, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionVideoProfile) Get_FrameDuration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FrameDuration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPerceptionVideoProfile) IsEqual(other *IPerceptionVideoProfile) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsEqual, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(other)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type KnownCameraIntrinsicsProperties struct {
	RtClass
}

func NewIKnownCameraIntrinsicsPropertiesStatics() *IKnownCameraIntrinsicsPropertiesStatics {
	var p *IKnownCameraIntrinsicsPropertiesStatics
	hs := NewHStr("Windows.Devices.Perception.KnownCameraIntrinsicsProperties")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IKnownCameraIntrinsicsPropertiesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type KnownPerceptionColorFrameSourceProperties struct {
	RtClass
}

func NewIKnownPerceptionColorFrameSourcePropertiesStatics() *IKnownPerceptionColorFrameSourcePropertiesStatics {
	var p *IKnownPerceptionColorFrameSourcePropertiesStatics
	hs := NewHStr("Windows.Devices.Perception.KnownPerceptionColorFrameSourceProperties")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IKnownPerceptionColorFrameSourcePropertiesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type KnownPerceptionDepthFrameSourceProperties struct {
	RtClass
}

func NewIKnownPerceptionDepthFrameSourcePropertiesStatics() *IKnownPerceptionDepthFrameSourcePropertiesStatics {
	var p *IKnownPerceptionDepthFrameSourcePropertiesStatics
	hs := NewHStr("Windows.Devices.Perception.KnownPerceptionDepthFrameSourceProperties")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IKnownPerceptionDepthFrameSourcePropertiesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type KnownPerceptionFrameSourceProperties struct {
	RtClass
}

func NewIKnownPerceptionFrameSourcePropertiesStatics2() *IKnownPerceptionFrameSourcePropertiesStatics2 {
	var p *IKnownPerceptionFrameSourcePropertiesStatics2
	hs := NewHStr("Windows.Devices.Perception.KnownPerceptionFrameSourceProperties")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IKnownPerceptionFrameSourcePropertiesStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIKnownPerceptionFrameSourcePropertiesStatics() *IKnownPerceptionFrameSourcePropertiesStatics {
	var p *IKnownPerceptionFrameSourcePropertiesStatics
	hs := NewHStr("Windows.Devices.Perception.KnownPerceptionFrameSourceProperties")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IKnownPerceptionFrameSourcePropertiesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type KnownPerceptionInfraredFrameSourceProperties struct {
	RtClass
}

func NewIKnownPerceptionInfraredFrameSourcePropertiesStatics() *IKnownPerceptionInfraredFrameSourcePropertiesStatics {
	var p *IKnownPerceptionInfraredFrameSourcePropertiesStatics
	hs := NewHStr("Windows.Devices.Perception.KnownPerceptionInfraredFrameSourceProperties")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IKnownPerceptionInfraredFrameSourcePropertiesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type KnownPerceptionVideoFrameSourceProperties struct {
	RtClass
}

func NewIKnownPerceptionVideoFrameSourcePropertiesStatics() *IKnownPerceptionVideoFrameSourcePropertiesStatics {
	var p *IKnownPerceptionVideoFrameSourcePropertiesStatics
	hs := NewHStr("Windows.Devices.Perception.KnownPerceptionVideoFrameSourceProperties")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IKnownPerceptionVideoFrameSourcePropertiesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type KnownPerceptionVideoProfileProperties struct {
	RtClass
}

func NewIKnownPerceptionVideoProfilePropertiesStatics() *IKnownPerceptionVideoProfilePropertiesStatics {
	var p *IKnownPerceptionVideoProfilePropertiesStatics
	hs := NewHStr("Windows.Devices.Perception.KnownPerceptionVideoProfileProperties")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IKnownPerceptionVideoProfilePropertiesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type PerceptionColorFrame struct {
	RtClass
	*IPerceptionColorFrame
}

type PerceptionColorFrameArrivedEventArgs struct {
	RtClass
	*IPerceptionColorFrameArrivedEventArgs
}

type PerceptionColorFrameReader struct {
	RtClass
	*IPerceptionColorFrameReader
}

type PerceptionColorFrameSource struct {
	RtClass
	*IPerceptionColorFrameSource
}

func NewIPerceptionColorFrameSourceStatics() *IPerceptionColorFrameSourceStatics {
	var p *IPerceptionColorFrameSourceStatics
	hs := NewHStr("Windows.Devices.Perception.PerceptionColorFrameSource")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPerceptionColorFrameSourceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type PerceptionColorFrameSourceAddedEventArgs struct {
	RtClass
	*IPerceptionColorFrameSourceAddedEventArgs
}

type PerceptionColorFrameSourceRemovedEventArgs struct {
	RtClass
	*IPerceptionColorFrameSourceRemovedEventArgs
}

type PerceptionColorFrameSourceWatcher struct {
	RtClass
	*IPerceptionColorFrameSourceWatcher
}

type PerceptionControlSession struct {
	RtClass
	*IPerceptionControlSession
}

type PerceptionDepthCorrelatedCameraIntrinsics struct {
	RtClass
	*IPerceptionDepthCorrelatedCameraIntrinsics
}

type PerceptionDepthCorrelatedCoordinateMapper struct {
	RtClass
	*IPerceptionDepthCorrelatedCoordinateMapper
}

type PerceptionDepthFrame struct {
	RtClass
	*IPerceptionDepthFrame
}

type PerceptionDepthFrameArrivedEventArgs struct {
	RtClass
	*IPerceptionDepthFrameArrivedEventArgs
}

type PerceptionDepthFrameReader struct {
	RtClass
	*IPerceptionDepthFrameReader
}

type PerceptionDepthFrameSource struct {
	RtClass
	*IPerceptionDepthFrameSource
}

func NewIPerceptionDepthFrameSourceStatics() *IPerceptionDepthFrameSourceStatics {
	var p *IPerceptionDepthFrameSourceStatics
	hs := NewHStr("Windows.Devices.Perception.PerceptionDepthFrameSource")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPerceptionDepthFrameSourceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type PerceptionDepthFrameSourceAddedEventArgs struct {
	RtClass
	*IPerceptionDepthFrameSourceAddedEventArgs
}

type PerceptionDepthFrameSourceRemovedEventArgs struct {
	RtClass
	*IPerceptionDepthFrameSourceRemovedEventArgs
}

type PerceptionDepthFrameSourceWatcher struct {
	RtClass
	*IPerceptionDepthFrameSourceWatcher
}

type PerceptionFrameSourcePropertiesChangedEventArgs struct {
	RtClass
	*IPerceptionFrameSourcePropertiesChangedEventArgs
}

type PerceptionFrameSourcePropertyChangeResult struct {
	RtClass
	*IPerceptionFrameSourcePropertyChangeResult
}

type PerceptionInfraredFrame struct {
	RtClass
	*IPerceptionInfraredFrame
}

type PerceptionInfraredFrameArrivedEventArgs struct {
	RtClass
	*IPerceptionInfraredFrameArrivedEventArgs
}

type PerceptionInfraredFrameReader struct {
	RtClass
	*IPerceptionInfraredFrameReader
}

type PerceptionInfraredFrameSource struct {
	RtClass
	*IPerceptionInfraredFrameSource
}

func NewIPerceptionInfraredFrameSourceStatics() *IPerceptionInfraredFrameSourceStatics {
	var p *IPerceptionInfraredFrameSourceStatics
	hs := NewHStr("Windows.Devices.Perception.PerceptionInfraredFrameSource")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IPerceptionInfraredFrameSourceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type PerceptionInfraredFrameSourceAddedEventArgs struct {
	RtClass
	*IPerceptionInfraredFrameSourceAddedEventArgs
}

type PerceptionInfraredFrameSourceRemovedEventArgs struct {
	RtClass
	*IPerceptionInfraredFrameSourceRemovedEventArgs
}

type PerceptionInfraredFrameSourceWatcher struct {
	RtClass
	*IPerceptionInfraredFrameSourceWatcher
}

type PerceptionVideoProfile struct {
	RtClass
	*IPerceptionVideoProfile
}
