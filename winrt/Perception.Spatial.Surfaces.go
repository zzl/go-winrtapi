package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"log"
	"syscall"
	"unsafe"
)

// interfaces

// F8E9EBE7-39B7-3962-BB03-57F56E1FB0A1
var IID_ISpatialSurfaceInfo = syscall.GUID{0xF8E9EBE7, 0x39B7, 0x3962,
	[8]byte{0xBB, 0x03, 0x57, 0xF5, 0x6E, 0x1F, 0xB0, 0xA1}}

type ISpatialSurfaceInfoInterface interface {
	win32.IInspectableInterface
	Get_Id() syscall.GUID
	Get_UpdateTime() DateTime
	TryGetBounds(coordinateSystem *ISpatialCoordinateSystem) *IReference[SpatialBoundingOrientedBox]
	TryComputeLatestMeshAsync(maxTrianglesPerCubicMeter float64) *IAsyncOperation[*ISpatialSurfaceMesh]
	TryComputeLatestMeshWithOptionsAsync(maxTrianglesPerCubicMeter float64, options *ISpatialSurfaceMeshOptions) *IAsyncOperation[*ISpatialSurfaceMesh]
}

type ISpatialSurfaceInfoVtbl struct {
	win32.IInspectableVtbl
	Get_Id                               uintptr
	Get_UpdateTime                       uintptr
	TryGetBounds                         uintptr
	TryComputeLatestMeshAsync            uintptr
	TryComputeLatestMeshWithOptionsAsync uintptr
}

type ISpatialSurfaceInfo struct {
	win32.IInspectable
}

func (this *ISpatialSurfaceInfo) Vtbl() *ISpatialSurfaceInfoVtbl {
	return (*ISpatialSurfaceInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialSurfaceInfo) Get_Id() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialSurfaceInfo) Get_UpdateTime() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UpdateTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialSurfaceInfo) TryGetBounds(coordinateSystem *ISpatialCoordinateSystem) *IReference[SpatialBoundingOrientedBox] {
	var _result *IReference[SpatialBoundingOrientedBox]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetBounds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(coordinateSystem)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialSurfaceInfo) TryComputeLatestMeshAsync(maxTrianglesPerCubicMeter float64) *IAsyncOperation[*ISpatialSurfaceMesh] {
	var _result *IAsyncOperation[*ISpatialSurfaceMesh]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryComputeLatestMeshAsync, uintptr(unsafe.Pointer(this)), uintptr(maxTrianglesPerCubicMeter), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialSurfaceInfo) TryComputeLatestMeshWithOptionsAsync(maxTrianglesPerCubicMeter float64, options *ISpatialSurfaceMeshOptions) *IAsyncOperation[*ISpatialSurfaceMesh] {
	var _result *IAsyncOperation[*ISpatialSurfaceMesh]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryComputeLatestMeshWithOptionsAsync, uintptr(unsafe.Pointer(this)), uintptr(maxTrianglesPerCubicMeter), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 108F57D9-DF0D-3950-A0FD-F972C77C27B4
var IID_ISpatialSurfaceMesh = syscall.GUID{0x108F57D9, 0xDF0D, 0x3950,
	[8]byte{0xA0, 0xFD, 0xF9, 0x72, 0xC7, 0x7C, 0x27, 0xB4}}

type ISpatialSurfaceMeshInterface interface {
	win32.IInspectableInterface
	Get_SurfaceInfo() *ISpatialSurfaceInfo
	Get_CoordinateSystem() *ISpatialCoordinateSystem
	Get_TriangleIndices() *ISpatialSurfaceMeshBuffer
	Get_VertexPositions() *ISpatialSurfaceMeshBuffer
	Get_VertexPositionScale() Vector3
	Get_VertexNormals() *ISpatialSurfaceMeshBuffer
}

type ISpatialSurfaceMeshVtbl struct {
	win32.IInspectableVtbl
	Get_SurfaceInfo         uintptr
	Get_CoordinateSystem    uintptr
	Get_TriangleIndices     uintptr
	Get_VertexPositions     uintptr
	Get_VertexPositionScale uintptr
	Get_VertexNormals       uintptr
}

type ISpatialSurfaceMesh struct {
	win32.IInspectable
}

func (this *ISpatialSurfaceMesh) Vtbl() *ISpatialSurfaceMeshVtbl {
	return (*ISpatialSurfaceMeshVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialSurfaceMesh) Get_SurfaceInfo() *ISpatialSurfaceInfo {
	var _result *ISpatialSurfaceInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SurfaceInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialSurfaceMesh) Get_CoordinateSystem() *ISpatialCoordinateSystem {
	var _result *ISpatialCoordinateSystem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CoordinateSystem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialSurfaceMesh) Get_TriangleIndices() *ISpatialSurfaceMeshBuffer {
	var _result *ISpatialSurfaceMeshBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TriangleIndices, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialSurfaceMesh) Get_VertexPositions() *ISpatialSurfaceMeshBuffer {
	var _result *ISpatialSurfaceMeshBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VertexPositions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialSurfaceMesh) Get_VertexPositionScale() Vector3 {
	var _result Vector3
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VertexPositionScale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialSurfaceMesh) Get_VertexNormals() *ISpatialSurfaceMeshBuffer {
	var _result *ISpatialSurfaceMeshBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VertexNormals, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 93CF59E0-871F-33F8-98B2-03D101458F6F
var IID_ISpatialSurfaceMeshBuffer = syscall.GUID{0x93CF59E0, 0x871F, 0x33F8,
	[8]byte{0x98, 0xB2, 0x03, 0xD1, 0x01, 0x45, 0x8F, 0x6F}}

type ISpatialSurfaceMeshBufferInterface interface {
	win32.IInspectableInterface
	Get_Format() unsafe.Pointer
	Get_Stride() uint32
	Get_ElementCount() uint32
	Get_Data() *IBuffer
}

type ISpatialSurfaceMeshBufferVtbl struct {
	win32.IInspectableVtbl
	Get_Format       uintptr
	Get_Stride       uintptr
	Get_ElementCount uintptr
	Get_Data         uintptr
}

type ISpatialSurfaceMeshBuffer struct {
	win32.IInspectable
}

func (this *ISpatialSurfaceMeshBuffer) Vtbl() *ISpatialSurfaceMeshBufferVtbl {
	return (*ISpatialSurfaceMeshBufferVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialSurfaceMeshBuffer) Get_Format() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Format, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialSurfaceMeshBuffer) Get_Stride() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Stride, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialSurfaceMeshBuffer) Get_ElementCount() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ElementCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialSurfaceMeshBuffer) Get_Data() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Data, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D2759F89-3572-3D2D-A10D-5FEE9394AA37
var IID_ISpatialSurfaceMeshOptions = syscall.GUID{0xD2759F89, 0x3572, 0x3D2D,
	[8]byte{0xA1, 0x0D, 0x5F, 0xEE, 0x93, 0x94, 0xAA, 0x37}}

type ISpatialSurfaceMeshOptionsInterface interface {
	win32.IInspectableInterface
	Get_VertexPositionFormat() unsafe.Pointer
	Put_VertexPositionFormat(value unsafe.Pointer)
	Get_TriangleIndexFormat() unsafe.Pointer
	Put_TriangleIndexFormat(value unsafe.Pointer)
	Get_VertexNormalFormat() unsafe.Pointer
	Put_VertexNormalFormat(value unsafe.Pointer)
	Get_IncludeVertexNormals() bool
	Put_IncludeVertexNormals(value bool)
}

type ISpatialSurfaceMeshOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_VertexPositionFormat uintptr
	Put_VertexPositionFormat uintptr
	Get_TriangleIndexFormat  uintptr
	Put_TriangleIndexFormat  uintptr
	Get_VertexNormalFormat   uintptr
	Put_VertexNormalFormat   uintptr
	Get_IncludeVertexNormals uintptr
	Put_IncludeVertexNormals uintptr
}

type ISpatialSurfaceMeshOptions struct {
	win32.IInspectable
}

func (this *ISpatialSurfaceMeshOptions) Vtbl() *ISpatialSurfaceMeshOptionsVtbl {
	return (*ISpatialSurfaceMeshOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialSurfaceMeshOptions) Get_VertexPositionFormat() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VertexPositionFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialSurfaceMeshOptions) Put_VertexPositionFormat(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_VertexPositionFormat, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISpatialSurfaceMeshOptions) Get_TriangleIndexFormat() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TriangleIndexFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialSurfaceMeshOptions) Put_TriangleIndexFormat(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TriangleIndexFormat, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISpatialSurfaceMeshOptions) Get_VertexNormalFormat() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VertexNormalFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialSurfaceMeshOptions) Put_VertexNormalFormat(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_VertexNormalFormat, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISpatialSurfaceMeshOptions) Get_IncludeVertexNormals() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IncludeVertexNormals, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialSurfaceMeshOptions) Put_IncludeVertexNormals(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IncludeVertexNormals, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 9B340ABF-9781-4505-8935-013575CAAE5E
var IID_ISpatialSurfaceMeshOptionsStatics = syscall.GUID{0x9B340ABF, 0x9781, 0x4505,
	[8]byte{0x89, 0x35, 0x01, 0x35, 0x75, 0xCA, 0xAE, 0x5E}}

type ISpatialSurfaceMeshOptionsStaticsInterface interface {
	win32.IInspectableInterface
	Get_SupportedVertexPositionFormats() *IVectorView[unsafe.Pointer]
	Get_SupportedTriangleIndexFormats() *IVectorView[unsafe.Pointer]
	Get_SupportedVertexNormalFormats() *IVectorView[unsafe.Pointer]
}

type ISpatialSurfaceMeshOptionsStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_SupportedVertexPositionFormats uintptr
	Get_SupportedTriangleIndexFormats  uintptr
	Get_SupportedVertexNormalFormats   uintptr
}

type ISpatialSurfaceMeshOptionsStatics struct {
	win32.IInspectable
}

func (this *ISpatialSurfaceMeshOptionsStatics) Vtbl() *ISpatialSurfaceMeshOptionsStaticsVtbl {
	return (*ISpatialSurfaceMeshOptionsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialSurfaceMeshOptionsStatics) Get_SupportedVertexPositionFormats() *IVectorView[unsafe.Pointer] {
	var _result *IVectorView[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedVertexPositionFormats, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialSurfaceMeshOptionsStatics) Get_SupportedTriangleIndexFormats() *IVectorView[unsafe.Pointer] {
	var _result *IVectorView[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedTriangleIndexFormats, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialSurfaceMeshOptionsStatics) Get_SupportedVertexNormalFormats() *IVectorView[unsafe.Pointer] {
	var _result *IVectorView[unsafe.Pointer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedVertexNormalFormats, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 10B69819-DDCA-3483-AC3A-748FE8C86DF5
var IID_ISpatialSurfaceObserver = syscall.GUID{0x10B69819, 0xDDCA, 0x3483,
	[8]byte{0xAC, 0x3A, 0x74, 0x8F, 0xE8, 0xC8, 0x6D, 0xF5}}

type ISpatialSurfaceObserverInterface interface {
	win32.IInspectableInterface
	GetObservedSurfaces() *IMapView[syscall.GUID, *ISpatialSurfaceInfo]
	SetBoundingVolume(bounds *ISpatialBoundingVolume)
	SetBoundingVolumes(bounds *IIterable[*ISpatialBoundingVolume])
	Add_ObservedSurfacesChanged(handler TypedEventHandler[*ISpatialSurfaceObserver, interface{}]) EventRegistrationToken
	Remove_ObservedSurfacesChanged(token EventRegistrationToken)
}

type ISpatialSurfaceObserverVtbl struct {
	win32.IInspectableVtbl
	GetObservedSurfaces            uintptr
	SetBoundingVolume              uintptr
	SetBoundingVolumes             uintptr
	Add_ObservedSurfacesChanged    uintptr
	Remove_ObservedSurfacesChanged uintptr
}

type ISpatialSurfaceObserver struct {
	win32.IInspectable
}

func (this *ISpatialSurfaceObserver) Vtbl() *ISpatialSurfaceObserverVtbl {
	return (*ISpatialSurfaceObserverVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialSurfaceObserver) GetObservedSurfaces() *IMapView[syscall.GUID, *ISpatialSurfaceInfo] {
	var _result *IMapView[syscall.GUID, *ISpatialSurfaceInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetObservedSurfaces, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpatialSurfaceObserver) SetBoundingVolume(bounds *ISpatialBoundingVolume) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetBoundingVolume, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bounds)))
	_ = _hr
}

func (this *ISpatialSurfaceObserver) SetBoundingVolumes(bounds *IIterable[*ISpatialBoundingVolume]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetBoundingVolumes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bounds)))
	_ = _hr
}

func (this *ISpatialSurfaceObserver) Add_ObservedSurfacesChanged(handler TypedEventHandler[*ISpatialSurfaceObserver, interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ObservedSurfacesChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpatialSurfaceObserver) Remove_ObservedSurfacesChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ObservedSurfacesChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 165951ED-2108-4168-9175-87E027BC9285
var IID_ISpatialSurfaceObserverStatics = syscall.GUID{0x165951ED, 0x2108, 0x4168,
	[8]byte{0x91, 0x75, 0x87, 0xE0, 0x27, 0xBC, 0x92, 0x85}}

type ISpatialSurfaceObserverStaticsInterface interface {
	win32.IInspectableInterface
	RequestAccessAsync() *IAsyncOperation[SpatialPerceptionAccessStatus]
}

type ISpatialSurfaceObserverStaticsVtbl struct {
	win32.IInspectableVtbl
	RequestAccessAsync uintptr
}

type ISpatialSurfaceObserverStatics struct {
	win32.IInspectable
}

func (this *ISpatialSurfaceObserverStatics) Vtbl() *ISpatialSurfaceObserverStaticsVtbl {
	return (*ISpatialSurfaceObserverStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialSurfaceObserverStatics) RequestAccessAsync() *IAsyncOperation[SpatialPerceptionAccessStatus] {
	var _result *IAsyncOperation[SpatialPerceptionAccessStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0F534261-C55D-4E6B-A895-A19DE69A42E3
var IID_ISpatialSurfaceObserverStatics2 = syscall.GUID{0x0F534261, 0xC55D, 0x4E6B,
	[8]byte{0xA8, 0x95, 0xA1, 0x9D, 0xE6, 0x9A, 0x42, 0xE3}}

type ISpatialSurfaceObserverStatics2Interface interface {
	win32.IInspectableInterface
	IsSupported() bool
}

type ISpatialSurfaceObserverStatics2Vtbl struct {
	win32.IInspectableVtbl
	IsSupported uintptr
}

type ISpatialSurfaceObserverStatics2 struct {
	win32.IInspectable
}

func (this *ISpatialSurfaceObserverStatics2) Vtbl() *ISpatialSurfaceObserverStatics2Vtbl {
	return (*ISpatialSurfaceObserverStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpatialSurfaceObserverStatics2) IsSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type SpatialSurfaceMeshOptions struct {
	RtClass
	*ISpatialSurfaceMeshOptions
}

func NewSpatialSurfaceMeshOptions() *SpatialSurfaceMeshOptions {
	hs := NewHStr("Windows.Perception.Spatial.Surfaces.SpatialSurfaceMeshOptions")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &SpatialSurfaceMeshOptions{
		RtClass:                    RtClass{PInspect: p},
		ISpatialSurfaceMeshOptions: (*ISpatialSurfaceMeshOptions)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewISpatialSurfaceMeshOptionsStatics() *ISpatialSurfaceMeshOptionsStatics {
	var p *ISpatialSurfaceMeshOptionsStatics
	hs := NewHStr("Windows.Perception.Spatial.Surfaces.SpatialSurfaceMeshOptions")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISpatialSurfaceMeshOptionsStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type SpatialSurfaceObserver struct {
	RtClass
	*ISpatialSurfaceObserver
}

func NewSpatialSurfaceObserver() *SpatialSurfaceObserver {
	hs := NewHStr("Windows.Perception.Spatial.Surfaces.SpatialSurfaceObserver")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &SpatialSurfaceObserver{
		RtClass:                 RtClass{PInspect: p},
		ISpatialSurfaceObserver: (*ISpatialSurfaceObserver)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewISpatialSurfaceObserverStatics() *ISpatialSurfaceObserverStatics {
	var p *ISpatialSurfaceObserverStatics
	hs := NewHStr("Windows.Perception.Spatial.Surfaces.SpatialSurfaceObserver")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISpatialSurfaceObserverStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewISpatialSurfaceObserverStatics2() *ISpatialSurfaceObserverStatics2 {
	var p *ISpatialSurfaceObserverStatics2
	hs := NewHStr("Windows.Perception.Spatial.Surfaces.SpatialSurfaceObserver")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISpatialSurfaceObserverStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
