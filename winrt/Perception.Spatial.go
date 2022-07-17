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
type SpatialAnchorExportPurpose int32

const (
	SpatialAnchorExportPurpose_Relocalization SpatialAnchorExportPurpose = 0
	SpatialAnchorExportPurpose_Sharing        SpatialAnchorExportPurpose = 1
)

// enum
type SpatialEntityWatcherStatus int32

const (
	SpatialEntityWatcherStatus_Created              SpatialEntityWatcherStatus = 0
	SpatialEntityWatcherStatus_Started              SpatialEntityWatcherStatus = 1
	SpatialEntityWatcherStatus_EnumerationCompleted SpatialEntityWatcherStatus = 2
	SpatialEntityWatcherStatus_Stopping             SpatialEntityWatcherStatus = 3
	SpatialEntityWatcherStatus_Stopped              SpatialEntityWatcherStatus = 4
	SpatialEntityWatcherStatus_Aborted              SpatialEntityWatcherStatus = 5
)

// enum
type SpatialLocatability int32

const (
	SpatialLocatability_Unavailable                  SpatialLocatability = 0
	SpatialLocatability_OrientationOnly              SpatialLocatability = 1
	SpatialLocatability_PositionalTrackingActivating SpatialLocatability = 2
	SpatialLocatability_PositionalTrackingActive     SpatialLocatability = 3
	SpatialLocatability_PositionalTrackingInhibited  SpatialLocatability = 4
)

// enum
type SpatialLookDirectionRange int32

const (
	SpatialLookDirectionRange_ForwardOnly     SpatialLookDirectionRange = 0
	SpatialLookDirectionRange_Omnidirectional SpatialLookDirectionRange = 1
)

// enum
type SpatialMovementRange int32

const (
	SpatialMovementRange_NoMovement SpatialMovementRange = 0
	SpatialMovementRange_Bounded    SpatialMovementRange = 1
)

// enum
type SpatialPerceptionAccessStatus int32

const (
	SpatialPerceptionAccessStatus_Unspecified    SpatialPerceptionAccessStatus = 0
	SpatialPerceptionAccessStatus_Allowed        SpatialPerceptionAccessStatus = 1
	SpatialPerceptionAccessStatus_DeniedByUser   SpatialPerceptionAccessStatus = 2
	SpatialPerceptionAccessStatus_DeniedBySystem SpatialPerceptionAccessStatus = 3
)

// structs

type SpatialBoundingBox struct {
	Center  Vector3
	Extents Vector3
}

type SpatialBoundingFrustum struct {
	Near   Plane
	Far    Plane
	Right  Plane
	Left   Plane
	Top    Plane
	Bottom Plane
}

type SpatialBoundingOrientedBox struct {
	Center      Vector3
	Extents     Vector3
	Orientation Quaternion
}

type SpatialBoundingSphere struct {
	Center Vector3
	Radius float32
}

type SpatialRay struct {
	Origin    Vector3
	Direction Vector3
}

// interfaces

// 0529E5CE-1D34-3702-BCEC-EABFF578A869
var IID_ISpatialAnchor = syscall.GUID{0x0529E5CE, 0x1D34, 0x3702,
	[8]byte{0xBC, 0xEC, 0xEA, 0xBF, 0xF5, 0x78, 0xA8, 0x69}}

type ISpatialAnchorInterface interface {
	win32.IInspectableInterface
	Get_CoordinateSystem() *ISpatialCoordinateSystem
	Get_RawCoordinateSystem() *ISpatialCoordinateSystem
	Add_RawCoordinateSystemAdjusted(handler TypedEventHandler[*ISpatialAnchor, *ISpatialAnchorRawCoordinateSystemAdjustedEventArgs]) EventRegistrationToken
	Remove_RawCoordinateSystemAdjusted(cookie EventRegistrationToken)
}

type ISpatialAnchorVtbl struct {
	win32.IInspectableVtbl
	Get_CoordinateSystem               uintptr
	Get_RawCoordinateSystem            uintptr
	Add_RawCoordinateSystemAdjusted    uintptr
	Remove_RawCoordinateSystemAdjusted uintptr
}

type ISpatialAnchor struct {
	win32.IInspectable
}

func (this *ISpatialAnchor) Vtbl() *ISpatialAnchorVtbl {
	return (*ISpatialAnchorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialAnchor) Get_CoordinateSystem() *ISpatialCoordinateSystem {
	var _result *ISpatialCoordinateSystem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CoordinateSystem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialAnchor) Get_RawCoordinateSystem() *ISpatialCoordinateSystem {
	var _result *ISpatialCoordinateSystem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RawCoordinateSystem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialAnchor) Add_RawCoordinateSystemAdjusted(handler TypedEventHandler[*ISpatialAnchor, *ISpatialAnchorRawCoordinateSystemAdjustedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_RawCoordinateSystemAdjusted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialAnchor) Remove_RawCoordinateSystemAdjusted(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_RawCoordinateSystemAdjusted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

// ED17C908-A695-4CF6-92FD-97263BA71047
var IID_ISpatialAnchor2 = syscall.GUID{0xED17C908, 0xA695, 0x4CF6,
	[8]byte{0x92, 0xFD, 0x97, 0x26, 0x3B, 0xA7, 0x10, 0x47}}

type ISpatialAnchor2Interface interface {
	win32.IInspectableInterface
	Get_RemovedByUser() bool
}

type ISpatialAnchor2Vtbl struct {
	win32.IInspectableVtbl
	Get_RemovedByUser uintptr
}

type ISpatialAnchor2 struct {
	win32.IInspectable
}

func (this *ISpatialAnchor2) Vtbl() *ISpatialAnchor2Vtbl {
	return (*ISpatialAnchor2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialAnchor2) Get_RemovedByUser() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RemovedByUser, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 77C25B2B-3409-4088-B91B-FDFD05D1648F
var IID_ISpatialAnchorExportSufficiency = syscall.GUID{0x77C25B2B, 0x3409, 0x4088,
	[8]byte{0xB9, 0x1B, 0xFD, 0xFD, 0x05, 0xD1, 0x64, 0x8F}}

type ISpatialAnchorExportSufficiencyInterface interface {
	win32.IInspectableInterface
	Get_IsMinimallySufficient() bool
	Get_SufficiencyLevel() float64
	Get_RecommendedSufficiencyLevel() float64
}

type ISpatialAnchorExportSufficiencyVtbl struct {
	win32.IInspectableVtbl
	Get_IsMinimallySufficient       uintptr
	Get_SufficiencyLevel            uintptr
	Get_RecommendedSufficiencyLevel uintptr
}

type ISpatialAnchorExportSufficiency struct {
	win32.IInspectable
}

func (this *ISpatialAnchorExportSufficiency) Vtbl() *ISpatialAnchorExportSufficiencyVtbl {
	return (*ISpatialAnchorExportSufficiencyVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialAnchorExportSufficiency) Get_IsMinimallySufficient() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsMinimallySufficient, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialAnchorExportSufficiency) Get_SufficiencyLevel() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SufficiencyLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialAnchorExportSufficiency) Get_RecommendedSufficiencyLevel() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RecommendedSufficiencyLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 9A2A4338-24FB-4269-89C5-88304AEEF20F
var IID_ISpatialAnchorExporter = syscall.GUID{0x9A2A4338, 0x24FB, 0x4269,
	[8]byte{0x89, 0xC5, 0x88, 0x30, 0x4A, 0xEE, 0xF2, 0x0F}}

type ISpatialAnchorExporterInterface interface {
	win32.IInspectableInterface
	GetAnchorExportSufficiencyAsync(anchor *ISpatialAnchor, purpose SpatialAnchorExportPurpose) *IAsyncOperation[*ISpatialAnchorExportSufficiency]
	TryExportAnchorAsync(anchor *ISpatialAnchor, purpose SpatialAnchorExportPurpose, stream *IOutputStream) *IAsyncOperation[bool]
}

type ISpatialAnchorExporterVtbl struct {
	win32.IInspectableVtbl
	GetAnchorExportSufficiencyAsync uintptr
	TryExportAnchorAsync            uintptr
}

type ISpatialAnchorExporter struct {
	win32.IInspectable
}

func (this *ISpatialAnchorExporter) Vtbl() *ISpatialAnchorExporterVtbl {
	return (*ISpatialAnchorExporterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialAnchorExporter) GetAnchorExportSufficiencyAsync(anchor *ISpatialAnchor, purpose SpatialAnchorExportPurpose) *IAsyncOperation[*ISpatialAnchorExportSufficiency] {
	var _result *IAsyncOperation[*ISpatialAnchorExportSufficiency]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAnchorExportSufficiencyAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(anchor)), uintptr(purpose), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialAnchorExporter) TryExportAnchorAsync(anchor *ISpatialAnchor, purpose SpatialAnchorExportPurpose, stream *IOutputStream) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryExportAnchorAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(anchor)), uintptr(purpose), uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// ED2507B8-2475-439C-85FF-7FED341FDC88
var IID_ISpatialAnchorExporterStatics = syscall.GUID{0xED2507B8, 0x2475, 0x439C,
	[8]byte{0x85, 0xFF, 0x7F, 0xED, 0x34, 0x1F, 0xDC, 0x88}}

type ISpatialAnchorExporterStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *ISpatialAnchorExporter
	RequestAccessAsync() *IAsyncOperation[SpatialPerceptionAccessStatus]
}

type ISpatialAnchorExporterStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault         uintptr
	RequestAccessAsync uintptr
}

type ISpatialAnchorExporterStatics struct {
	win32.IInspectable
}

func (this *ISpatialAnchorExporterStatics) Vtbl() *ISpatialAnchorExporterStaticsVtbl {
	return (*ISpatialAnchorExporterStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialAnchorExporterStatics) GetDefault() *ISpatialAnchorExporter {
	var _result *ISpatialAnchorExporter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialAnchorExporterStatics) RequestAccessAsync() *IAsyncOperation[SpatialPerceptionAccessStatus] {
	var _result *IAsyncOperation[SpatialPerceptionAccessStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 88E30EAB-F3B7-420B-B086-8A80C07D910D
var IID_ISpatialAnchorManagerStatics = syscall.GUID{0x88E30EAB, 0xF3B7, 0x420B,
	[8]byte{0xB0, 0x86, 0x8A, 0x80, 0xC0, 0x7D, 0x91, 0x0D}}

type ISpatialAnchorManagerStaticsInterface interface {
	win32.IInspectableInterface
	RequestStoreAsync() *IAsyncOperation[*ISpatialAnchorStore]
}

type ISpatialAnchorManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	RequestStoreAsync uintptr
}

type ISpatialAnchorManagerStatics struct {
	win32.IInspectable
}

func (this *ISpatialAnchorManagerStatics) Vtbl() *ISpatialAnchorManagerStaticsVtbl {
	return (*ISpatialAnchorManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialAnchorManagerStatics) RequestStoreAsync() *IAsyncOperation[*ISpatialAnchorStore] {
	var _result *IAsyncOperation[*ISpatialAnchorStore]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestStoreAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A1E81EB8-56C7-3117-A2E4-81E0FCF28E00
var IID_ISpatialAnchorRawCoordinateSystemAdjustedEventArgs = syscall.GUID{0xA1E81EB8, 0x56C7, 0x3117,
	[8]byte{0xA2, 0xE4, 0x81, 0xE0, 0xFC, 0xF2, 0x8E, 0x00}}

type ISpatialAnchorRawCoordinateSystemAdjustedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_OldRawCoordinateSystemToNewRawCoordinateSystemTransform() unsafe.Pointer
}

type ISpatialAnchorRawCoordinateSystemAdjustedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_OldRawCoordinateSystemToNewRawCoordinateSystemTransform uintptr
}

type ISpatialAnchorRawCoordinateSystemAdjustedEventArgs struct {
	win32.IInspectable
}

func (this *ISpatialAnchorRawCoordinateSystemAdjustedEventArgs) Vtbl() *ISpatialAnchorRawCoordinateSystemAdjustedEventArgsVtbl {
	return (*ISpatialAnchorRawCoordinateSystemAdjustedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialAnchorRawCoordinateSystemAdjustedEventArgs) Get_OldRawCoordinateSystemToNewRawCoordinateSystemTransform() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OldRawCoordinateSystemToNewRawCoordinateSystemTransform, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// A9928642-0174-311C-AE79-0E5107669F16
var IID_ISpatialAnchorStatics = syscall.GUID{0xA9928642, 0x0174, 0x311C,
	[8]byte{0xAE, 0x79, 0x0E, 0x51, 0x07, 0x66, 0x9F, 0x16}}

type ISpatialAnchorStaticsInterface interface {
	win32.IInspectableInterface
	TryCreateRelativeTo(coordinateSystem *ISpatialCoordinateSystem) *ISpatialAnchor
	TryCreateWithPositionRelativeTo(coordinateSystem *ISpatialCoordinateSystem, position unsafe.Pointer) *ISpatialAnchor
	TryCreateWithPositionAndOrientationRelativeTo(coordinateSystem *ISpatialCoordinateSystem, position unsafe.Pointer, orientation unsafe.Pointer) *ISpatialAnchor
}

type ISpatialAnchorStaticsVtbl struct {
	win32.IInspectableVtbl
	TryCreateRelativeTo                           uintptr
	TryCreateWithPositionRelativeTo               uintptr
	TryCreateWithPositionAndOrientationRelativeTo uintptr
}

type ISpatialAnchorStatics struct {
	win32.IInspectable
}

func (this *ISpatialAnchorStatics) Vtbl() *ISpatialAnchorStaticsVtbl {
	return (*ISpatialAnchorStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialAnchorStatics) TryCreateRelativeTo(coordinateSystem *ISpatialCoordinateSystem) *ISpatialAnchor {
	var _result *ISpatialAnchor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryCreateRelativeTo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(coordinateSystem)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialAnchorStatics) TryCreateWithPositionRelativeTo(coordinateSystem *ISpatialCoordinateSystem, position unsafe.Pointer) *ISpatialAnchor {
	var _result *ISpatialAnchor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryCreateWithPositionRelativeTo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(coordinateSystem)), uintptr(position), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialAnchorStatics) TryCreateWithPositionAndOrientationRelativeTo(coordinateSystem *ISpatialCoordinateSystem, position unsafe.Pointer, orientation unsafe.Pointer) *ISpatialAnchor {
	var _result *ISpatialAnchor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryCreateWithPositionAndOrientationRelativeTo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(coordinateSystem)), uintptr(position), uintptr(orientation), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B0BC3636-486A-3CB0-9E6F-1245165C4DB6
var IID_ISpatialAnchorStore = syscall.GUID{0xB0BC3636, 0x486A, 0x3CB0,
	[8]byte{0x9E, 0x6F, 0x12, 0x45, 0x16, 0x5C, 0x4D, 0xB6}}

type ISpatialAnchorStoreInterface interface {
	win32.IInspectableInterface
	GetAllSavedAnchors() *IMapView[string, *ISpatialAnchor]
	TrySave(id string, anchor *ISpatialAnchor) bool
	Remove(id string)
	Clear()
}

type ISpatialAnchorStoreVtbl struct {
	win32.IInspectableVtbl
	GetAllSavedAnchors uintptr
	TrySave            uintptr
	Remove             uintptr
	Clear              uintptr
}

type ISpatialAnchorStore struct {
	win32.IInspectable
}

func (this *ISpatialAnchorStore) Vtbl() *ISpatialAnchorStoreVtbl {
	return (*ISpatialAnchorStoreVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialAnchorStore) GetAllSavedAnchors() *IMapView[string, *ISpatialAnchor] {
	var _result *IMapView[string, *ISpatialAnchor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAllSavedAnchors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialAnchorStore) TrySave(id string, anchor *ISpatialAnchor) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TrySave, uintptr(unsafe.Pointer(this)), NewHStr(id).Ptr, uintptr(unsafe.Pointer(anchor)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialAnchorStore) Remove(id string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove, uintptr(unsafe.Pointer(this)), NewHStr(id).Ptr)
	_ = _hr
}

func (this *ISpatialAnchorStore) Clear() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Clear, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 03BBF9B9-12D8-4BCE-8835-C5DF3AC0ADAB
var IID_ISpatialAnchorTransferManagerStatics = syscall.GUID{0x03BBF9B9, 0x12D8, 0x4BCE,
	[8]byte{0x88, 0x35, 0xC5, 0xDF, 0x3A, 0xC0, 0xAD, 0xAB}}

type ISpatialAnchorTransferManagerStaticsInterface interface {
	win32.IInspectableInterface
	TryImportAnchorsAsync(stream *IInputStream) *IAsyncOperation[*IMapView[string, *ISpatialAnchor]]
	TryExportAnchorsAsync(anchors *IIterable[*IKeyValuePair[string, *ISpatialAnchor]], stream *IOutputStream) *IAsyncOperation[bool]
	RequestAccessAsync() *IAsyncOperation[SpatialPerceptionAccessStatus]
}

type ISpatialAnchorTransferManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	TryImportAnchorsAsync uintptr
	TryExportAnchorsAsync uintptr
	RequestAccessAsync    uintptr
}

type ISpatialAnchorTransferManagerStatics struct {
	win32.IInspectable
}

func (this *ISpatialAnchorTransferManagerStatics) Vtbl() *ISpatialAnchorTransferManagerStaticsVtbl {
	return (*ISpatialAnchorTransferManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialAnchorTransferManagerStatics) TryImportAnchorsAsync(stream *IInputStream) *IAsyncOperation[*IMapView[string, *ISpatialAnchor]] {
	var _result *IAsyncOperation[*IMapView[string, *ISpatialAnchor]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryImportAnchorsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialAnchorTransferManagerStatics) TryExportAnchorsAsync(anchors *IIterable[*IKeyValuePair[string, *ISpatialAnchor]], stream *IOutputStream) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryExportAnchorsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(anchors)), uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialAnchorTransferManagerStatics) RequestAccessAsync() *IAsyncOperation[SpatialPerceptionAccessStatus] {
	var _result *IAsyncOperation[SpatialPerceptionAccessStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FB2065DA-68C3-33DF-B7AF-4C787207999C
var IID_ISpatialBoundingVolume = syscall.GUID{0xFB2065DA, 0x68C3, 0x33DF,
	[8]byte{0xB7, 0xAF, 0x4C, 0x78, 0x72, 0x07, 0x99, 0x9C}}

type ISpatialBoundingVolumeInterface interface {
	win32.IInspectableInterface
}

type ISpatialBoundingVolumeVtbl struct {
	win32.IInspectableVtbl
}

type ISpatialBoundingVolume struct {
	win32.IInspectable
}

func (this *ISpatialBoundingVolume) Vtbl() *ISpatialBoundingVolumeVtbl {
	return (*ISpatialBoundingVolumeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 05889117-B3E1-36D8-B017-566181A5B196
var IID_ISpatialBoundingVolumeStatics = syscall.GUID{0x05889117, 0xB3E1, 0x36D8,
	[8]byte{0xB0, 0x17, 0x56, 0x61, 0x81, 0xA5, 0xB1, 0x96}}

type ISpatialBoundingVolumeStaticsInterface interface {
	win32.IInspectableInterface
	FromBox(coordinateSystem *ISpatialCoordinateSystem, box SpatialBoundingBox) *ISpatialBoundingVolume
	FromOrientedBox(coordinateSystem *ISpatialCoordinateSystem, box SpatialBoundingOrientedBox) *ISpatialBoundingVolume
	FromSphere(coordinateSystem *ISpatialCoordinateSystem, sphere SpatialBoundingSphere) *ISpatialBoundingVolume
	FromFrustum(coordinateSystem *ISpatialCoordinateSystem, frustum SpatialBoundingFrustum) *ISpatialBoundingVolume
}

type ISpatialBoundingVolumeStaticsVtbl struct {
	win32.IInspectableVtbl
	FromBox         uintptr
	FromOrientedBox uintptr
	FromSphere      uintptr
	FromFrustum     uintptr
}

type ISpatialBoundingVolumeStatics struct {
	win32.IInspectable
}

func (this *ISpatialBoundingVolumeStatics) Vtbl() *ISpatialBoundingVolumeStaticsVtbl {
	return (*ISpatialBoundingVolumeStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialBoundingVolumeStatics) FromBox(coordinateSystem *ISpatialCoordinateSystem, box SpatialBoundingBox) *ISpatialBoundingVolume {
	var _result *ISpatialBoundingVolume
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromBox, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(coordinateSystem)), uintptr(unsafe.Pointer(&box)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialBoundingVolumeStatics) FromOrientedBox(coordinateSystem *ISpatialCoordinateSystem, box SpatialBoundingOrientedBox) *ISpatialBoundingVolume {
	var _result *ISpatialBoundingVolume
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromOrientedBox, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(coordinateSystem)), uintptr(unsafe.Pointer(&box)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialBoundingVolumeStatics) FromSphere(coordinateSystem *ISpatialCoordinateSystem, sphere SpatialBoundingSphere) *ISpatialBoundingVolume {
	var _result *ISpatialBoundingVolume
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromSphere, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(coordinateSystem)), uintptr(unsafe.Pointer(&sphere)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialBoundingVolumeStatics) FromFrustum(coordinateSystem *ISpatialCoordinateSystem, frustum SpatialBoundingFrustum) *ISpatialBoundingVolume {
	var _result *ISpatialBoundingVolume
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromFrustum, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(coordinateSystem)), uintptr(unsafe.Pointer(&frustum)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 69EBCA4B-60A3-3586-A653-59A7BD676D07
var IID_ISpatialCoordinateSystem = syscall.GUID{0x69EBCA4B, 0x60A3, 0x3586,
	[8]byte{0xA6, 0x53, 0x59, 0xA7, 0xBD, 0x67, 0x6D, 0x07}}

type ISpatialCoordinateSystemInterface interface {
	win32.IInspectableInterface
	TryGetTransformTo(target *ISpatialCoordinateSystem) *IReference[unsafe.Pointer]
}

type ISpatialCoordinateSystemVtbl struct {
	win32.IInspectableVtbl
	TryGetTransformTo uintptr
}

type ISpatialCoordinateSystem struct {
	win32.IInspectable
}

func (this *ISpatialCoordinateSystem) Vtbl() *ISpatialCoordinateSystemVtbl {
	return (*ISpatialCoordinateSystemVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialCoordinateSystem) TryGetTransformTo(target *ISpatialCoordinateSystem) *IReference[unsafe.Pointer] {
	var _result *IReference[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetTransformTo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(target)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 166DE955-E1EB-454C-BA08-E6C0668DDC65
var IID_ISpatialEntity = syscall.GUID{0x166DE955, 0xE1EB, 0x454C,
	[8]byte{0xBA, 0x08, 0xE6, 0xC0, 0x66, 0x8D, 0xDC, 0x65}}

type ISpatialEntityInterface interface {
	win32.IInspectableInterface
	Get_Id() string
	Get_Anchor() *ISpatialAnchor
	Get_Properties() *IPropertySet
}

type ISpatialEntityVtbl struct {
	win32.IInspectableVtbl
	Get_Id         uintptr
	Get_Anchor     uintptr
	Get_Properties uintptr
}

type ISpatialEntity struct {
	win32.IInspectable
}

func (this *ISpatialEntity) Vtbl() *ISpatialEntityVtbl {
	return (*ISpatialEntityVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialEntity) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISpatialEntity) Get_Anchor() *ISpatialAnchor {
	var _result *ISpatialAnchor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Anchor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialEntity) Get_Properties() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A397F49B-156A-4707-AC2C-D31D570ED399
var IID_ISpatialEntityAddedEventArgs = syscall.GUID{0xA397F49B, 0x156A, 0x4707,
	[8]byte{0xAC, 0x2C, 0xD3, 0x1D, 0x57, 0x0E, 0xD3, 0x99}}

type ISpatialEntityAddedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Entity() *ISpatialEntity
}

type ISpatialEntityAddedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Entity uintptr
}

type ISpatialEntityAddedEventArgs struct {
	win32.IInspectable
}

func (this *ISpatialEntityAddedEventArgs) Vtbl() *ISpatialEntityAddedEventArgsVtbl {
	return (*ISpatialEntityAddedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialEntityAddedEventArgs) Get_Entity() *ISpatialEntity {
	var _result *ISpatialEntity
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Entity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E1F1E325-349F-4225-A2F3-4B01C15FE056
var IID_ISpatialEntityFactory = syscall.GUID{0xE1F1E325, 0x349F, 0x4225,
	[8]byte{0xA2, 0xF3, 0x4B, 0x01, 0xC1, 0x5F, 0xE0, 0x56}}

type ISpatialEntityFactoryInterface interface {
	win32.IInspectableInterface
	CreateWithSpatialAnchor(spatialAnchor *ISpatialAnchor) *ISpatialEntity
	CreateWithSpatialAnchorAndProperties(spatialAnchor *ISpatialAnchor, propertySet *IPropertySet) *ISpatialEntity
}

type ISpatialEntityFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateWithSpatialAnchor              uintptr
	CreateWithSpatialAnchorAndProperties uintptr
}

type ISpatialEntityFactory struct {
	win32.IInspectable
}

func (this *ISpatialEntityFactory) Vtbl() *ISpatialEntityFactoryVtbl {
	return (*ISpatialEntityFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialEntityFactory) CreateWithSpatialAnchor(spatialAnchor *ISpatialAnchor) *ISpatialEntity {
	var _result *ISpatialEntity
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithSpatialAnchor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(spatialAnchor)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialEntityFactory) CreateWithSpatialAnchorAndProperties(spatialAnchor *ISpatialAnchor, propertySet *IPropertySet) *ISpatialEntity {
	var _result *ISpatialEntity
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithSpatialAnchorAndProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(spatialAnchor)), uintptr(unsafe.Pointer(propertySet)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 91741800-536D-4E9F-ABF6-415B5444D651
var IID_ISpatialEntityRemovedEventArgs = syscall.GUID{0x91741800, 0x536D, 0x4E9F,
	[8]byte{0xAB, 0xF6, 0x41, 0x5B, 0x54, 0x44, 0xD6, 0x51}}

type ISpatialEntityRemovedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Entity() *ISpatialEntity
}

type ISpatialEntityRemovedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Entity uintptr
}

type ISpatialEntityRemovedEventArgs struct {
	win32.IInspectable
}

func (this *ISpatialEntityRemovedEventArgs) Vtbl() *ISpatialEntityRemovedEventArgsVtbl {
	return (*ISpatialEntityRemovedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialEntityRemovedEventArgs) Get_Entity() *ISpatialEntity {
	var _result *ISpatialEntity
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Entity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 329788BA-E513-4F06-889D-1BE30ECF43E6
var IID_ISpatialEntityStore = syscall.GUID{0x329788BA, 0xE513, 0x4F06,
	[8]byte{0x88, 0x9D, 0x1B, 0xE3, 0x0E, 0xCF, 0x43, 0xE6}}

type ISpatialEntityStoreInterface interface {
	win32.IInspectableInterface
	SaveAsync(entity *ISpatialEntity) *IAsyncAction
	RemoveAsync(entity *ISpatialEntity) *IAsyncAction
	CreateEntityWatcher() *ISpatialEntityWatcher
}

type ISpatialEntityStoreVtbl struct {
	win32.IInspectableVtbl
	SaveAsync           uintptr
	RemoveAsync         uintptr
	CreateEntityWatcher uintptr
}

type ISpatialEntityStore struct {
	win32.IInspectable
}

func (this *ISpatialEntityStore) Vtbl() *ISpatialEntityStoreVtbl {
	return (*ISpatialEntityStoreVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialEntityStore) SaveAsync(entity *ISpatialEntity) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SaveAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(entity)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialEntityStore) RemoveAsync(entity *ISpatialEntity) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(entity)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialEntityStore) CreateEntityWatcher() *ISpatialEntityWatcher {
	var _result *ISpatialEntityWatcher
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateEntityWatcher, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6B4B389E-7C50-4E92-8A62-4D1D4B7CCD3E
var IID_ISpatialEntityStoreStatics = syscall.GUID{0x6B4B389E, 0x7C50, 0x4E92,
	[8]byte{0x8A, 0x62, 0x4D, 0x1D, 0x4B, 0x7C, 0xCD, 0x3E}}

type ISpatialEntityStoreStaticsInterface interface {
	win32.IInspectableInterface
	Get_IsSupported() bool
	TryGetForRemoteSystemSession(session unsafe.Pointer) *ISpatialEntityStore
}

type ISpatialEntityStoreStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_IsSupported              uintptr
	TryGetForRemoteSystemSession uintptr
}

type ISpatialEntityStoreStatics struct {
	win32.IInspectable
}

func (this *ISpatialEntityStoreStatics) Vtbl() *ISpatialEntityStoreStaticsVtbl {
	return (*ISpatialEntityStoreStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialEntityStoreStatics) Get_IsSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialEntityStoreStatics) TryGetForRemoteSystemSession(session unsafe.Pointer) *ISpatialEntityStore {
	var _result *ISpatialEntityStore
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetForRemoteSystemSession, uintptr(unsafe.Pointer(this)), uintptr(session), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E5671766-627B-43CB-A49F-B3BE6D47DEED
var IID_ISpatialEntityUpdatedEventArgs = syscall.GUID{0xE5671766, 0x627B, 0x43CB,
	[8]byte{0xA4, 0x9F, 0xB3, 0xBE, 0x6D, 0x47, 0xDE, 0xED}}

type ISpatialEntityUpdatedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Entity() *ISpatialEntity
}

type ISpatialEntityUpdatedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Entity uintptr
}

type ISpatialEntityUpdatedEventArgs struct {
	win32.IInspectable
}

func (this *ISpatialEntityUpdatedEventArgs) Vtbl() *ISpatialEntityUpdatedEventArgsVtbl {
	return (*ISpatialEntityUpdatedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialEntityUpdatedEventArgs) Get_Entity() *ISpatialEntity {
	var _result *ISpatialEntity
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Entity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B3B85FA0-6D5E-4BBC-805D-5FE5B9BA1959
var IID_ISpatialEntityWatcher = syscall.GUID{0xB3B85FA0, 0x6D5E, 0x4BBC,
	[8]byte{0x80, 0x5D, 0x5F, 0xE5, 0xB9, 0xBA, 0x19, 0x59}}

type ISpatialEntityWatcherInterface interface {
	win32.IInspectableInterface
	Get_Status() SpatialEntityWatcherStatus
	Add_Added(handler TypedEventHandler[*ISpatialEntityWatcher, *ISpatialEntityAddedEventArgs]) EventRegistrationToken
	Remove_Added(token EventRegistrationToken)
	Add_Updated(handler TypedEventHandler[*ISpatialEntityWatcher, *ISpatialEntityUpdatedEventArgs]) EventRegistrationToken
	Remove_Updated(token EventRegistrationToken)
	Add_Removed(handler TypedEventHandler[*ISpatialEntityWatcher, *ISpatialEntityRemovedEventArgs]) EventRegistrationToken
	Remove_Removed(token EventRegistrationToken)
	Add_EnumerationCompleted(handler TypedEventHandler[*ISpatialEntityWatcher, interface{}]) EventRegistrationToken
	Remove_EnumerationCompleted(token EventRegistrationToken)
	Start()
	Stop()
}

type ISpatialEntityWatcherVtbl struct {
	win32.IInspectableVtbl
	Get_Status                  uintptr
	Add_Added                   uintptr
	Remove_Added                uintptr
	Add_Updated                 uintptr
	Remove_Updated              uintptr
	Add_Removed                 uintptr
	Remove_Removed              uintptr
	Add_EnumerationCompleted    uintptr
	Remove_EnumerationCompleted uintptr
	Start                       uintptr
	Stop                        uintptr
}

type ISpatialEntityWatcher struct {
	win32.IInspectable
}

func (this *ISpatialEntityWatcher) Vtbl() *ISpatialEntityWatcherVtbl {
	return (*ISpatialEntityWatcherVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialEntityWatcher) Get_Status() SpatialEntityWatcherStatus {
	var _result SpatialEntityWatcherStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialEntityWatcher) Add_Added(handler TypedEventHandler[*ISpatialEntityWatcher, *ISpatialEntityAddedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Added, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialEntityWatcher) Remove_Added(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Added, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISpatialEntityWatcher) Add_Updated(handler TypedEventHandler[*ISpatialEntityWatcher, *ISpatialEntityUpdatedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Updated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialEntityWatcher) Remove_Updated(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Updated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISpatialEntityWatcher) Add_Removed(handler TypedEventHandler[*ISpatialEntityWatcher, *ISpatialEntityRemovedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Removed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialEntityWatcher) Remove_Removed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Removed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISpatialEntityWatcher) Add_EnumerationCompleted(handler TypedEventHandler[*ISpatialEntityWatcher, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_EnumerationCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialEntityWatcher) Remove_EnumerationCompleted(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_EnumerationCompleted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *ISpatialEntityWatcher) Start() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Start, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *ISpatialEntityWatcher) Stop() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Stop, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 1D81D29D-24A1-37D5-8FA1-39B4F9AD67E2
var IID_ISpatialLocation = syscall.GUID{0x1D81D29D, 0x24A1, 0x37D5,
	[8]byte{0x8F, 0xA1, 0x39, 0xB4, 0xF9, 0xAD, 0x67, 0xE2}}

type ISpatialLocationInterface interface {
	win32.IInspectableInterface
	Get_Position() unsafe.Pointer
	Get_Orientation() unsafe.Pointer
	Get_AbsoluteLinearVelocity() unsafe.Pointer
	Get_AbsoluteLinearAcceleration() unsafe.Pointer
	Get_AbsoluteAngularVelocity() unsafe.Pointer
	Get_AbsoluteAngularAcceleration() unsafe.Pointer
}

type ISpatialLocationVtbl struct {
	win32.IInspectableVtbl
	Get_Position                    uintptr
	Get_Orientation                 uintptr
	Get_AbsoluteLinearVelocity      uintptr
	Get_AbsoluteLinearAcceleration  uintptr
	Get_AbsoluteAngularVelocity     uintptr
	Get_AbsoluteAngularAcceleration uintptr
}

type ISpatialLocation struct {
	win32.IInspectable
}

func (this *ISpatialLocation) Vtbl() *ISpatialLocationVtbl {
	return (*ISpatialLocationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialLocation) Get_Position() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialLocation) Get_Orientation() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Orientation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialLocation) Get_AbsoluteLinearVelocity() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AbsoluteLinearVelocity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialLocation) Get_AbsoluteLinearAcceleration() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AbsoluteLinearAcceleration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialLocation) Get_AbsoluteAngularVelocity() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AbsoluteAngularVelocity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialLocation) Get_AbsoluteAngularAcceleration() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AbsoluteAngularAcceleration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 117F2416-38A7-4A18-B404-AB8FABE1D78B
var IID_ISpatialLocation2 = syscall.GUID{0x117F2416, 0x38A7, 0x4A18,
	[8]byte{0xB4, 0x04, 0xAB, 0x8F, 0xAB, 0xE1, 0xD7, 0x8B}}

type ISpatialLocation2Interface interface {
	win32.IInspectableInterface
	Get_AbsoluteAngularVelocityAxisAngle() unsafe.Pointer
	Get_AbsoluteAngularAccelerationAxisAngle() unsafe.Pointer
}

type ISpatialLocation2Vtbl struct {
	win32.IInspectableVtbl
	Get_AbsoluteAngularVelocityAxisAngle     uintptr
	Get_AbsoluteAngularAccelerationAxisAngle uintptr
}

type ISpatialLocation2 struct {
	win32.IInspectable
}

func (this *ISpatialLocation2) Vtbl() *ISpatialLocation2Vtbl {
	return (*ISpatialLocation2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialLocation2) Get_AbsoluteAngularVelocityAxisAngle() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AbsoluteAngularVelocityAxisAngle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialLocation2) Get_AbsoluteAngularAccelerationAxisAngle() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AbsoluteAngularAccelerationAxisAngle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F6478925-9E0C-3BB6-997E-B64ECCA24CF4
var IID_ISpatialLocator = syscall.GUID{0xF6478925, 0x9E0C, 0x3BB6,
	[8]byte{0x99, 0x7E, 0xB6, 0x4E, 0xCC, 0xA2, 0x4C, 0xF4}}

type ISpatialLocatorInterface interface {
	win32.IInspectableInterface
	Get_Locatability() SpatialLocatability
	Add_LocatabilityChanged(handler TypedEventHandler[*ISpatialLocator, interface{}]) EventRegistrationToken
	Remove_LocatabilityChanged(cookie EventRegistrationToken)
	Add_PositionalTrackingDeactivating(handler TypedEventHandler[*ISpatialLocator, *ISpatialLocatorPositionalTrackingDeactivatingEventArgs]) EventRegistrationToken
	Remove_PositionalTrackingDeactivating(cookie EventRegistrationToken)
	TryLocateAtTimestamp(timestamp *IPerceptionTimestamp, coordinateSystem *ISpatialCoordinateSystem) *ISpatialLocation
	CreateAttachedFrameOfReferenceAtCurrentHeading() *ISpatialLocatorAttachedFrameOfReference
	CreateAttachedFrameOfReferenceAtCurrentHeadingWithPosition(relativePosition unsafe.Pointer) *ISpatialLocatorAttachedFrameOfReference
	CreateAttachedFrameOfReferenceAtCurrentHeadingWithPositionAndOrientation(relativePosition unsafe.Pointer, relativeOrientation unsafe.Pointer) *ISpatialLocatorAttachedFrameOfReference
	CreateAttachedFrameOfReferenceAtCurrentHeadingWithPositionAndOrientationAndRelativeHeading(relativePosition unsafe.Pointer, relativeOrientation unsafe.Pointer, relativeHeadingInRadians float64) *ISpatialLocatorAttachedFrameOfReference
	CreateStationaryFrameOfReferenceAtCurrentLocation() *ISpatialStationaryFrameOfReference
	CreateStationaryFrameOfReferenceAtCurrentLocationWithPosition(relativePosition unsafe.Pointer) *ISpatialStationaryFrameOfReference
	CreateStationaryFrameOfReferenceAtCurrentLocationWithPositionAndOrientation(relativePosition unsafe.Pointer, relativeOrientation unsafe.Pointer) *ISpatialStationaryFrameOfReference
	CreateStationaryFrameOfReferenceAtCurrentLocationWithPositionAndOrientationAndRelativeHeading(relativePosition unsafe.Pointer, relativeOrientation unsafe.Pointer, relativeHeadingInRadians float64) *ISpatialStationaryFrameOfReference
}

type ISpatialLocatorVtbl struct {
	win32.IInspectableVtbl
	Get_Locatability                                                                              uintptr
	Add_LocatabilityChanged                                                                       uintptr
	Remove_LocatabilityChanged                                                                    uintptr
	Add_PositionalTrackingDeactivating                                                            uintptr
	Remove_PositionalTrackingDeactivating                                                         uintptr
	TryLocateAtTimestamp                                                                          uintptr
	CreateAttachedFrameOfReferenceAtCurrentHeading                                                uintptr
	CreateAttachedFrameOfReferenceAtCurrentHeadingWithPosition                                    uintptr
	CreateAttachedFrameOfReferenceAtCurrentHeadingWithPositionAndOrientation                      uintptr
	CreateAttachedFrameOfReferenceAtCurrentHeadingWithPositionAndOrientationAndRelativeHeading    uintptr
	CreateStationaryFrameOfReferenceAtCurrentLocation                                             uintptr
	CreateStationaryFrameOfReferenceAtCurrentLocationWithPosition                                 uintptr
	CreateStationaryFrameOfReferenceAtCurrentLocationWithPositionAndOrientation                   uintptr
	CreateStationaryFrameOfReferenceAtCurrentLocationWithPositionAndOrientationAndRelativeHeading uintptr
}

type ISpatialLocator struct {
	win32.IInspectable
}

func (this *ISpatialLocator) Vtbl() *ISpatialLocatorVtbl {
	return (*ISpatialLocatorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialLocator) Get_Locatability() SpatialLocatability {
	var _result SpatialLocatability
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Locatability, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialLocator) Add_LocatabilityChanged(handler TypedEventHandler[*ISpatialLocator, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_LocatabilityChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialLocator) Remove_LocatabilityChanged(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_LocatabilityChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *ISpatialLocator) Add_PositionalTrackingDeactivating(handler TypedEventHandler[*ISpatialLocator, *ISpatialLocatorPositionalTrackingDeactivatingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PositionalTrackingDeactivating, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialLocator) Remove_PositionalTrackingDeactivating(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PositionalTrackingDeactivating, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *ISpatialLocator) TryLocateAtTimestamp(timestamp *IPerceptionTimestamp, coordinateSystem *ISpatialCoordinateSystem) *ISpatialLocation {
	var _result *ISpatialLocation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryLocateAtTimestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(timestamp)), uintptr(unsafe.Pointer(coordinateSystem)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialLocator) CreateAttachedFrameOfReferenceAtCurrentHeading() *ISpatialLocatorAttachedFrameOfReference {
	var _result *ISpatialLocatorAttachedFrameOfReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAttachedFrameOfReferenceAtCurrentHeading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialLocator) CreateAttachedFrameOfReferenceAtCurrentHeadingWithPosition(relativePosition unsafe.Pointer) *ISpatialLocatorAttachedFrameOfReference {
	var _result *ISpatialLocatorAttachedFrameOfReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAttachedFrameOfReferenceAtCurrentHeadingWithPosition, uintptr(unsafe.Pointer(this)), uintptr(relativePosition), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialLocator) CreateAttachedFrameOfReferenceAtCurrentHeadingWithPositionAndOrientation(relativePosition unsafe.Pointer, relativeOrientation unsafe.Pointer) *ISpatialLocatorAttachedFrameOfReference {
	var _result *ISpatialLocatorAttachedFrameOfReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAttachedFrameOfReferenceAtCurrentHeadingWithPositionAndOrientation, uintptr(unsafe.Pointer(this)), uintptr(relativePosition), uintptr(relativeOrientation), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialLocator) CreateAttachedFrameOfReferenceAtCurrentHeadingWithPositionAndOrientationAndRelativeHeading(relativePosition unsafe.Pointer, relativeOrientation unsafe.Pointer, relativeHeadingInRadians float64) *ISpatialLocatorAttachedFrameOfReference {
	var _result *ISpatialLocatorAttachedFrameOfReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAttachedFrameOfReferenceAtCurrentHeadingWithPositionAndOrientationAndRelativeHeading, uintptr(unsafe.Pointer(this)), uintptr(relativePosition), uintptr(relativeOrientation), uintptr(relativeHeadingInRadians), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialLocator) CreateStationaryFrameOfReferenceAtCurrentLocation() *ISpatialStationaryFrameOfReference {
	var _result *ISpatialStationaryFrameOfReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateStationaryFrameOfReferenceAtCurrentLocation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialLocator) CreateStationaryFrameOfReferenceAtCurrentLocationWithPosition(relativePosition unsafe.Pointer) *ISpatialStationaryFrameOfReference {
	var _result *ISpatialStationaryFrameOfReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateStationaryFrameOfReferenceAtCurrentLocationWithPosition, uintptr(unsafe.Pointer(this)), uintptr(relativePosition), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialLocator) CreateStationaryFrameOfReferenceAtCurrentLocationWithPositionAndOrientation(relativePosition unsafe.Pointer, relativeOrientation unsafe.Pointer) *ISpatialStationaryFrameOfReference {
	var _result *ISpatialStationaryFrameOfReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateStationaryFrameOfReferenceAtCurrentLocationWithPositionAndOrientation, uintptr(unsafe.Pointer(this)), uintptr(relativePosition), uintptr(relativeOrientation), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialLocator) CreateStationaryFrameOfReferenceAtCurrentLocationWithPositionAndOrientationAndRelativeHeading(relativePosition unsafe.Pointer, relativeOrientation unsafe.Pointer, relativeHeadingInRadians float64) *ISpatialStationaryFrameOfReference {
	var _result *ISpatialStationaryFrameOfReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateStationaryFrameOfReferenceAtCurrentLocationWithPositionAndOrientationAndRelativeHeading, uintptr(unsafe.Pointer(this)), uintptr(relativePosition), uintptr(relativeOrientation), uintptr(relativeHeadingInRadians), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E1774EF6-1F4F-499C-9625-EF5E6ED7A048
var IID_ISpatialLocatorAttachedFrameOfReference = syscall.GUID{0xE1774EF6, 0x1F4F, 0x499C,
	[8]byte{0x96, 0x25, 0xEF, 0x5E, 0x6E, 0xD7, 0xA0, 0x48}}

type ISpatialLocatorAttachedFrameOfReferenceInterface interface {
	win32.IInspectableInterface
	Get_RelativePosition() unsafe.Pointer
	Put_RelativePosition(value unsafe.Pointer)
	Get_RelativeOrientation() unsafe.Pointer
	Put_RelativeOrientation(value unsafe.Pointer)
	AdjustHeading(headingOffsetInRadians float64)
	GetStationaryCoordinateSystemAtTimestamp(timestamp *IPerceptionTimestamp) *ISpatialCoordinateSystem
	TryGetRelativeHeadingAtTimestamp(timestamp *IPerceptionTimestamp) *IReference[float64]
}

type ISpatialLocatorAttachedFrameOfReferenceVtbl struct {
	win32.IInspectableVtbl
	Get_RelativePosition                     uintptr
	Put_RelativePosition                     uintptr
	Get_RelativeOrientation                  uintptr
	Put_RelativeOrientation                  uintptr
	AdjustHeading                            uintptr
	GetStationaryCoordinateSystemAtTimestamp uintptr
	TryGetRelativeHeadingAtTimestamp         uintptr
}

type ISpatialLocatorAttachedFrameOfReference struct {
	win32.IInspectable
}

func (this *ISpatialLocatorAttachedFrameOfReference) Vtbl() *ISpatialLocatorAttachedFrameOfReferenceVtbl {
	return (*ISpatialLocatorAttachedFrameOfReferenceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialLocatorAttachedFrameOfReference) Get_RelativePosition() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RelativePosition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialLocatorAttachedFrameOfReference) Put_RelativePosition(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RelativePosition, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISpatialLocatorAttachedFrameOfReference) Get_RelativeOrientation() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RelativeOrientation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialLocatorAttachedFrameOfReference) Put_RelativeOrientation(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RelativeOrientation, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISpatialLocatorAttachedFrameOfReference) AdjustHeading(headingOffsetInRadians float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AdjustHeading, uintptr(unsafe.Pointer(this)), uintptr(headingOffsetInRadians))
	_ = _hr
}

func (this *ISpatialLocatorAttachedFrameOfReference) GetStationaryCoordinateSystemAtTimestamp(timestamp *IPerceptionTimestamp) *ISpatialCoordinateSystem {
	var _result *ISpatialCoordinateSystem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStationaryCoordinateSystemAtTimestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(timestamp)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialLocatorAttachedFrameOfReference) TryGetRelativeHeadingAtTimestamp(timestamp *IPerceptionTimestamp) *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetRelativeHeadingAtTimestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(timestamp)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B8A84063-E3F4-368B-9061-9EA9D1D6CC16
var IID_ISpatialLocatorPositionalTrackingDeactivatingEventArgs = syscall.GUID{0xB8A84063, 0xE3F4, 0x368B,
	[8]byte{0x90, 0x61, 0x9E, 0xA9, 0xD1, 0xD6, 0xCC, 0x16}}

type ISpatialLocatorPositionalTrackingDeactivatingEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Canceled() bool
	Put_Canceled(value bool)
}

type ISpatialLocatorPositionalTrackingDeactivatingEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Canceled uintptr
	Put_Canceled uintptr
}

type ISpatialLocatorPositionalTrackingDeactivatingEventArgs struct {
	win32.IInspectable
}

func (this *ISpatialLocatorPositionalTrackingDeactivatingEventArgs) Vtbl() *ISpatialLocatorPositionalTrackingDeactivatingEventArgsVtbl {
	return (*ISpatialLocatorPositionalTrackingDeactivatingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialLocatorPositionalTrackingDeactivatingEventArgs) Get_Canceled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Canceled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialLocatorPositionalTrackingDeactivatingEventArgs) Put_Canceled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Canceled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// B76E3340-A7C2-361B-BB82-56E93B89B1BB
var IID_ISpatialLocatorStatics = syscall.GUID{0xB76E3340, 0xA7C2, 0x361B,
	[8]byte{0xBB, 0x82, 0x56, 0xE9, 0x3B, 0x89, 0xB1, 0xBB}}

type ISpatialLocatorStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *ISpatialLocator
}

type ISpatialLocatorStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
}

type ISpatialLocatorStatics struct {
	win32.IInspectable
}

func (this *ISpatialLocatorStatics) Vtbl() *ISpatialLocatorStaticsVtbl {
	return (*ISpatialLocatorStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialLocatorStatics) GetDefault() *ISpatialLocator {
	var _result *ISpatialLocator
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7A8A3464-AD0D-4590-AB86-33062B674926
var IID_ISpatialStageFrameOfReference = syscall.GUID{0x7A8A3464, 0xAD0D, 0x4590,
	[8]byte{0xAB, 0x86, 0x33, 0x06, 0x2B, 0x67, 0x49, 0x26}}

type ISpatialStageFrameOfReferenceInterface interface {
	win32.IInspectableInterface
	Get_CoordinateSystem() *ISpatialCoordinateSystem
	Get_MovementRange() SpatialMovementRange
	Get_LookDirectionRange() SpatialLookDirectionRange
	GetCoordinateSystemAtCurrentLocation(locator *ISpatialLocator) *ISpatialCoordinateSystem
	TryGetMovementBounds(coordinateSystem *ISpatialCoordinateSystem) unsafe.Pointer
}

type ISpatialStageFrameOfReferenceVtbl struct {
	win32.IInspectableVtbl
	Get_CoordinateSystem                 uintptr
	Get_MovementRange                    uintptr
	Get_LookDirectionRange               uintptr
	GetCoordinateSystemAtCurrentLocation uintptr
	TryGetMovementBounds                 uintptr
}

type ISpatialStageFrameOfReference struct {
	win32.IInspectable
}

func (this *ISpatialStageFrameOfReference) Vtbl() *ISpatialStageFrameOfReferenceVtbl {
	return (*ISpatialStageFrameOfReferenceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialStageFrameOfReference) Get_CoordinateSystem() *ISpatialCoordinateSystem {
	var _result *ISpatialCoordinateSystem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CoordinateSystem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialStageFrameOfReference) Get_MovementRange() SpatialMovementRange {
	var _result SpatialMovementRange
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MovementRange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialStageFrameOfReference) Get_LookDirectionRange() SpatialLookDirectionRange {
	var _result SpatialLookDirectionRange
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_LookDirectionRange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialStageFrameOfReference) GetCoordinateSystemAtCurrentLocation(locator *ISpatialLocator) *ISpatialCoordinateSystem {
	var _result *ISpatialCoordinateSystem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCoordinateSystemAtCurrentLocation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(locator)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialStageFrameOfReference) TryGetMovementBounds(coordinateSystem *ISpatialCoordinateSystem) unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetMovementBounds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(coordinateSystem)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F78D5C4D-A0A4-499C-8D91-A8C965D40654
var IID_ISpatialStageFrameOfReferenceStatics = syscall.GUID{0xF78D5C4D, 0xA0A4, 0x499C,
	[8]byte{0x8D, 0x91, 0xA8, 0xC9, 0x65, 0xD4, 0x06, 0x54}}

type ISpatialStageFrameOfReferenceStaticsInterface interface {
	win32.IInspectableInterface
	Get_Current() *ISpatialStageFrameOfReference
	Add_CurrentChanged(handler EventHandler[interface{}]) EventRegistrationToken
	Remove_CurrentChanged(cookie EventRegistrationToken)
	RequestNewStageAsync() *IAsyncOperation[*ISpatialStageFrameOfReference]
}

type ISpatialStageFrameOfReferenceStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_Current           uintptr
	Add_CurrentChanged    uintptr
	Remove_CurrentChanged uintptr
	RequestNewStageAsync  uintptr
}

type ISpatialStageFrameOfReferenceStatics struct {
	win32.IInspectable
}

func (this *ISpatialStageFrameOfReferenceStatics) Vtbl() *ISpatialStageFrameOfReferenceStaticsVtbl {
	return (*ISpatialStageFrameOfReferenceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialStageFrameOfReferenceStatics) Get_Current() *ISpatialStageFrameOfReference {
	var _result *ISpatialStageFrameOfReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Current, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialStageFrameOfReferenceStatics) Add_CurrentChanged(handler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_CurrentChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialStageFrameOfReferenceStatics) Remove_CurrentChanged(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_CurrentChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *ISpatialStageFrameOfReferenceStatics) RequestNewStageAsync() *IAsyncOperation[*ISpatialStageFrameOfReference] {
	var _result *IAsyncOperation[*ISpatialStageFrameOfReference]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestNewStageAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 09DBCCB9-BCF8-3E7F-BE7E-7EDCCBB178A8
var IID_ISpatialStationaryFrameOfReference = syscall.GUID{0x09DBCCB9, 0xBCF8, 0x3E7F,
	[8]byte{0xBE, 0x7E, 0x7E, 0xDC, 0xCB, 0xB1, 0x78, 0xA8}}

type ISpatialStationaryFrameOfReferenceInterface interface {
	win32.IInspectableInterface
	Get_CoordinateSystem() *ISpatialCoordinateSystem
}

type ISpatialStationaryFrameOfReferenceVtbl struct {
	win32.IInspectableVtbl
	Get_CoordinateSystem uintptr
}

type ISpatialStationaryFrameOfReference struct {
	win32.IInspectable
}

func (this *ISpatialStationaryFrameOfReference) Vtbl() *ISpatialStationaryFrameOfReferenceVtbl {
	return (*ISpatialStationaryFrameOfReferenceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialStationaryFrameOfReference) Get_CoordinateSystem() *ISpatialCoordinateSystem {
	var _result *ISpatialCoordinateSystem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CoordinateSystem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type SpatialAnchorManager struct {
	RtClass
}

func NewISpatialAnchorManagerStatics() *ISpatialAnchorManagerStatics {
	var p *ISpatialAnchorManagerStatics
	hs := NewHStr("Windows.Perception.Spatial.SpatialAnchorManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISpatialAnchorManagerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type SpatialEntity struct {
	RtClass
	*ISpatialEntity
}

func NewSpatialEntity_CreateWithSpatialAnchor(spatialAnchor *ISpatialAnchor) *SpatialEntity {
	hs := NewHStr("Windows.Perception.Spatial.SpatialEntity")
	var pFac *ISpatialEntityFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISpatialEntityFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISpatialEntity
	p = pFac.CreateWithSpatialAnchor(spatialAnchor)
	result := &SpatialEntity{
		RtClass:        RtClass{PInspect: &p.IInspectable},
		ISpatialEntity: p,
	}
	com.AddToScope(result)
	return result
}

func NewSpatialEntity_CreateWithSpatialAnchorAndProperties(spatialAnchor *ISpatialAnchor, propertySet *IPropertySet) *SpatialEntity {
	hs := NewHStr("Windows.Perception.Spatial.SpatialEntity")
	var pFac *ISpatialEntityFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISpatialEntityFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISpatialEntity
	p = pFac.CreateWithSpatialAnchorAndProperties(spatialAnchor, propertySet)
	result := &SpatialEntity{
		RtClass:        RtClass{PInspect: &p.IInspectable},
		ISpatialEntity: p,
	}
	com.AddToScope(result)
	return result
}
