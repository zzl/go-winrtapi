package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type SceneAlphaMode int32

const (
	SceneAlphaMode_Opaque    SceneAlphaMode = 0
	SceneAlphaMode_AlphaTest SceneAlphaMode = 1
	SceneAlphaMode_Blend     SceneAlphaMode = 2
)

// enum
type SceneAttributeSemantic int32

const (
	SceneAttributeSemantic_Index     SceneAttributeSemantic = 0
	SceneAttributeSemantic_Vertex    SceneAttributeSemantic = 1
	SceneAttributeSemantic_Normal    SceneAttributeSemantic = 2
	SceneAttributeSemantic_TexCoord0 SceneAttributeSemantic = 3
	SceneAttributeSemantic_TexCoord1 SceneAttributeSemantic = 4
	SceneAttributeSemantic_Color     SceneAttributeSemantic = 5
	SceneAttributeSemantic_Tangent   SceneAttributeSemantic = 6
)

// enum
type SceneComponentType int32

const (
	SceneComponentType_MeshRendererComponent SceneComponentType = 0
)

// enum
type SceneWrappingMode int32

const (
	SceneWrappingMode_ClampToEdge    SceneWrappingMode = 0
	SceneWrappingMode_MirroredRepeat SceneWrappingMode = 1
	SceneWrappingMode_Repeat         SceneWrappingMode = 2
)

// interfaces

// 5D8FFC70-C618-4083-8251-9962593114AA
var IID_ISceneBoundingBox = syscall.GUID{0x5D8FFC70, 0xC618, 0x4083,
	[8]byte{0x82, 0x51, 0x99, 0x62, 0x59, 0x31, 0x14, 0xAA}}

type ISceneBoundingBoxInterface interface {
	win32.IInspectableInterface
	Get_Center() Vector3
	Get_Extents() Vector3
	Get_Max() Vector3
	Get_Min() Vector3
	Get_Size() Vector3
}

type ISceneBoundingBoxVtbl struct {
	win32.IInspectableVtbl
	Get_Center  uintptr
	Get_Extents uintptr
	Get_Max     uintptr
	Get_Min     uintptr
	Get_Size    uintptr
}

type ISceneBoundingBox struct {
	win32.IInspectable
}

func (this *ISceneBoundingBox) Vtbl() *ISceneBoundingBoxVtbl {
	return (*ISceneBoundingBoxVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISceneBoundingBox) Get_Center() Vector3 {
	var _result Vector3
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Center, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISceneBoundingBox) Get_Extents() Vector3 {
	var _result Vector3
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Extents, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISceneBoundingBox) Get_Max() Vector3 {
	var _result Vector3
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Max, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISceneBoundingBox) Get_Min() Vector3 {
	var _result Vector3
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Min, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISceneBoundingBox) Get_Size() Vector3 {
	var _result Vector3
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Size, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// AE20FC96-226C-44BD-95CB-DD5ED9EBE9A5
var IID_ISceneComponent = syscall.GUID{0xAE20FC96, 0x226C, 0x44BD,
	[8]byte{0x95, 0xCB, 0xDD, 0x5E, 0xD9, 0xEB, 0xE9, 0xA5}}

type ISceneComponentInterface interface {
	win32.IInspectableInterface
	Get_ComponentType() SceneComponentType
}

type ISceneComponentVtbl struct {
	win32.IInspectableVtbl
	Get_ComponentType uintptr
}

type ISceneComponent struct {
	win32.IInspectable
}

func (this *ISceneComponent) Vtbl() *ISceneComponentVtbl {
	return (*ISceneComponentVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISceneComponent) Get_ComponentType() SceneComponentType {
	var _result SceneComponentType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ComponentType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C483791C-5F46-45E4-B666-A3D2259F9B2E
var IID_ISceneComponentCollection = syscall.GUID{0xC483791C, 0x5F46, 0x45E4,
	[8]byte{0xB6, 0x66, 0xA3, 0xD2, 0x25, 0x9F, 0x9B, 0x2E}}

type ISceneComponentCollectionInterface interface {
	win32.IInspectableInterface
}

type ISceneComponentCollectionVtbl struct {
	win32.IInspectableVtbl
}

type ISceneComponentCollection struct {
	win32.IInspectable
}

func (this *ISceneComponentCollection) Vtbl() *ISceneComponentCollectionVtbl {
	return (*ISceneComponentCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 5FBC5574-DDD8-5889-AB5B-D8FA716E7C9E
var IID_ISceneComponentFactory = syscall.GUID{0x5FBC5574, 0xDDD8, 0x5889,
	[8]byte{0xAB, 0x5B, 0xD8, 0xFA, 0x71, 0x6E, 0x7C, 0x9E}}

type ISceneComponentFactoryInterface interface {
	win32.IInspectableInterface
}

type ISceneComponentFactoryVtbl struct {
	win32.IInspectableVtbl
}

type ISceneComponentFactory struct {
	win32.IInspectable
}

func (this *ISceneComponentFactory) Vtbl() *ISceneComponentFactoryVtbl {
	return (*ISceneComponentFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 8CA74B7C-30DF-4E07-9490-37875AF1A123
var IID_ISceneMaterial = syscall.GUID{0x8CA74B7C, 0x30DF, 0x4E07,
	[8]byte{0x94, 0x90, 0x37, 0x87, 0x5A, 0xF1, 0xA1, 0x23}}

type ISceneMaterialInterface interface {
	win32.IInspectableInterface
}

type ISceneMaterialVtbl struct {
	win32.IInspectableVtbl
}

type ISceneMaterial struct {
	win32.IInspectable
}

func (this *ISceneMaterial) Vtbl() *ISceneMaterialVtbl {
	return (*ISceneMaterialVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 67536C19-A707-5254-A495-7FDC799893B9
var IID_ISceneMaterialFactory = syscall.GUID{0x67536C19, 0xA707, 0x5254,
	[8]byte{0xA4, 0x95, 0x7F, 0xDC, 0x79, 0x98, 0x93, 0xB9}}

type ISceneMaterialFactoryInterface interface {
	win32.IInspectableInterface
}

type ISceneMaterialFactoryVtbl struct {
	win32.IInspectableVtbl
}

type ISceneMaterialFactory struct {
	win32.IInspectable
}

func (this *ISceneMaterialFactory) Vtbl() *ISceneMaterialFactoryVtbl {
	return (*ISceneMaterialFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 422A1642-1EF1-485C-97E9-AE6F95AD812F
var IID_ISceneMaterialInput = syscall.GUID{0x422A1642, 0x1EF1, 0x485C,
	[8]byte{0x97, 0xE9, 0xAE, 0x6F, 0x95, 0xAD, 0x81, 0x2F}}

type ISceneMaterialInputInterface interface {
	win32.IInspectableInterface
}

type ISceneMaterialInputVtbl struct {
	win32.IInspectableVtbl
}

type ISceneMaterialInput struct {
	win32.IInspectable
}

func (this *ISceneMaterialInput) Vtbl() *ISceneMaterialInputVtbl {
	return (*ISceneMaterialInputVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// A88FEB74-7D0A-5E4C-A748-1015AF9CA74F
var IID_ISceneMaterialInputFactory = syscall.GUID{0xA88FEB74, 0x7D0A, 0x5E4C,
	[8]byte{0xA7, 0x48, 0x10, 0x15, 0xAF, 0x9C, 0xA7, 0x4F}}

type ISceneMaterialInputFactoryInterface interface {
	win32.IInspectableInterface
}

type ISceneMaterialInputFactoryVtbl struct {
	win32.IInspectableVtbl
}

type ISceneMaterialInputFactory struct {
	win32.IInspectable
}

func (this *ISceneMaterialInputFactory) Vtbl() *ISceneMaterialInputFactoryVtbl {
	return (*ISceneMaterialInputFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// EE9A1530-1155-4C0C-92BD-40020CF78347
var IID_ISceneMesh = syscall.GUID{0xEE9A1530, 0x1155, 0x4C0C,
	[8]byte{0x92, 0xBD, 0x40, 0x02, 0x0C, 0xF7, 0x83, 0x47}}

type ISceneMeshInterface interface {
	win32.IInspectableInterface
	Get_Bounds() *ISceneBoundingBox
	Get_PrimitiveTopology() unsafe.Pointer
	Put_PrimitiveTopology(value unsafe.Pointer)
	FillMeshAttribute(semantic SceneAttributeSemantic, format unsafe.Pointer, memory *IMemoryBuffer)
}

type ISceneMeshVtbl struct {
	win32.IInspectableVtbl
	Get_Bounds            uintptr
	Get_PrimitiveTopology uintptr
	Put_PrimitiveTopology uintptr
	FillMeshAttribute     uintptr
}

type ISceneMesh struct {
	win32.IInspectable
}

func (this *ISceneMesh) Vtbl() *ISceneMeshVtbl {
	return (*ISceneMeshVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISceneMesh) Get_Bounds() *ISceneBoundingBox {
	var _result *ISceneBoundingBox
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bounds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISceneMesh) Get_PrimitiveTopology() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PrimitiveTopology, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISceneMesh) Put_PrimitiveTopology(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PrimitiveTopology, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISceneMesh) FillMeshAttribute(semantic SceneAttributeSemantic, format unsafe.Pointer, memory *IMemoryBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FillMeshAttribute, uintptr(unsafe.Pointer(this)), uintptr(semantic), uintptr(format), uintptr(unsafe.Pointer(memory)))
	_ = _hr
}

// CE843171-3D43-4855-AA69-31FF988D049D
var IID_ISceneMeshMaterialAttributeMap = syscall.GUID{0xCE843171, 0x3D43, 0x4855,
	[8]byte{0xAA, 0x69, 0x31, 0xFF, 0x98, 0x8D, 0x04, 0x9D}}

type ISceneMeshMaterialAttributeMapInterface interface {
	win32.IInspectableInterface
}

type ISceneMeshMaterialAttributeMapVtbl struct {
	win32.IInspectableVtbl
}

type ISceneMeshMaterialAttributeMap struct {
	win32.IInspectable
}

func (this *ISceneMeshMaterialAttributeMap) Vtbl() *ISceneMeshMaterialAttributeMapVtbl {
	return (*ISceneMeshMaterialAttributeMapVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 9929F7E3-6364-477E-98FE-74ED9FD4C2DE
var IID_ISceneMeshRendererComponent = syscall.GUID{0x9929F7E3, 0x6364, 0x477E,
	[8]byte{0x98, 0xFE, 0x74, 0xED, 0x9F, 0xD4, 0xC2, 0xDE}}

type ISceneMeshRendererComponentInterface interface {
	win32.IInspectableInterface
	Get_Material() *ISceneMaterial
	Put_Material(value *ISceneMaterial)
	Get_Mesh() *ISceneMesh
	Put_Mesh(value *ISceneMesh)
	Get_UVMappings() *ISceneMeshMaterialAttributeMap
}

type ISceneMeshRendererComponentVtbl struct {
	win32.IInspectableVtbl
	Get_Material   uintptr
	Put_Material   uintptr
	Get_Mesh       uintptr
	Put_Mesh       uintptr
	Get_UVMappings uintptr
}

type ISceneMeshRendererComponent struct {
	win32.IInspectable
}

func (this *ISceneMeshRendererComponent) Vtbl() *ISceneMeshRendererComponentVtbl {
	return (*ISceneMeshRendererComponentVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISceneMeshRendererComponent) Get_Material() *ISceneMaterial {
	var _result *ISceneMaterial
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Material, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISceneMeshRendererComponent) Put_Material(value *ISceneMaterial) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Material, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ISceneMeshRendererComponent) Get_Mesh() *ISceneMesh {
	var _result *ISceneMesh
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mesh, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISceneMeshRendererComponent) Put_Mesh(value *ISceneMesh) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Mesh, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ISceneMeshRendererComponent) Get_UVMappings() *ISceneMeshMaterialAttributeMap {
	var _result *ISceneMeshMaterialAttributeMap
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UVMappings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4954F37A-4459-4521-BD6E-2B38B8D711EA
var IID_ISceneMeshRendererComponentStatics = syscall.GUID{0x4954F37A, 0x4459, 0x4521,
	[8]byte{0xBD, 0x6E, 0x2B, 0x38, 0xB8, 0xD7, 0x11, 0xEA}}

type ISceneMeshRendererComponentStaticsInterface interface {
	win32.IInspectableInterface
	Create(compositor *ICompositor) *ISceneMeshRendererComponent
}

type ISceneMeshRendererComponentStaticsVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ISceneMeshRendererComponentStatics struct {
	win32.IInspectable
}

func (this *ISceneMeshRendererComponentStatics) Vtbl() *ISceneMeshRendererComponentStaticsVtbl {
	return (*ISceneMeshRendererComponentStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISceneMeshRendererComponentStatics) Create(compositor *ICompositor) *ISceneMeshRendererComponent {
	var _result *ISceneMeshRendererComponent
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(compositor)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8412316C-7B57-473F-966B-81DC277B1751
var IID_ISceneMeshStatics = syscall.GUID{0x8412316C, 0x7B57, 0x473F,
	[8]byte{0x96, 0x6B, 0x81, 0xDC, 0x27, 0x7B, 0x17, 0x51}}

type ISceneMeshStaticsInterface interface {
	win32.IInspectableInterface
	Create(compositor *ICompositor) *ISceneMesh
}

type ISceneMeshStaticsVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ISceneMeshStatics struct {
	win32.IInspectable
}

func (this *ISceneMeshStatics) Vtbl() *ISceneMeshStaticsVtbl {
	return (*ISceneMeshStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISceneMeshStatics) Create(compositor *ICompositor) *ISceneMesh {
	var _result *ISceneMesh
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(compositor)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C1D91446-799C-429E-A4E4-5DA645F18E61
var IID_ISceneMetallicRoughnessMaterial = syscall.GUID{0xC1D91446, 0x799C, 0x429E,
	[8]byte{0xA4, 0xE4, 0x5D, 0xA6, 0x45, 0xF1, 0x8E, 0x61}}

type ISceneMetallicRoughnessMaterialInterface interface {
	win32.IInspectableInterface
	Get_BaseColorInput() *ISceneMaterialInput
	Put_BaseColorInput(value *ISceneMaterialInput)
	Get_BaseColorFactor() Vector4
	Put_BaseColorFactor(value Vector4)
	Get_MetallicFactor() float32
	Put_MetallicFactor(value float32)
	Get_MetallicRoughnessInput() *ISceneMaterialInput
	Put_MetallicRoughnessInput(value *ISceneMaterialInput)
	Get_RoughnessFactor() float32
	Put_RoughnessFactor(value float32)
}

type ISceneMetallicRoughnessMaterialVtbl struct {
	win32.IInspectableVtbl
	Get_BaseColorInput         uintptr
	Put_BaseColorInput         uintptr
	Get_BaseColorFactor        uintptr
	Put_BaseColorFactor        uintptr
	Get_MetallicFactor         uintptr
	Put_MetallicFactor         uintptr
	Get_MetallicRoughnessInput uintptr
	Put_MetallicRoughnessInput uintptr
	Get_RoughnessFactor        uintptr
	Put_RoughnessFactor        uintptr
}

type ISceneMetallicRoughnessMaterial struct {
	win32.IInspectable
}

func (this *ISceneMetallicRoughnessMaterial) Vtbl() *ISceneMetallicRoughnessMaterialVtbl {
	return (*ISceneMetallicRoughnessMaterialVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISceneMetallicRoughnessMaterial) Get_BaseColorInput() *ISceneMaterialInput {
	var _result *ISceneMaterialInput
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BaseColorInput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISceneMetallicRoughnessMaterial) Put_BaseColorInput(value *ISceneMaterialInput) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BaseColorInput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ISceneMetallicRoughnessMaterial) Get_BaseColorFactor() Vector4 {
	var _result Vector4
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BaseColorFactor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISceneMetallicRoughnessMaterial) Put_BaseColorFactor(value Vector4) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BaseColorFactor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ISceneMetallicRoughnessMaterial) Get_MetallicFactor() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MetallicFactor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISceneMetallicRoughnessMaterial) Put_MetallicFactor(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MetallicFactor, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISceneMetallicRoughnessMaterial) Get_MetallicRoughnessInput() *ISceneMaterialInput {
	var _result *ISceneMaterialInput
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MetallicRoughnessInput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISceneMetallicRoughnessMaterial) Put_MetallicRoughnessInput(value *ISceneMaterialInput) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_MetallicRoughnessInput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ISceneMetallicRoughnessMaterial) Get_RoughnessFactor() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RoughnessFactor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISceneMetallicRoughnessMaterial) Put_RoughnessFactor(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RoughnessFactor, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 3BDDCA50-6D9D-4531-8DC4-B27E3E49B7AB
var IID_ISceneMetallicRoughnessMaterialStatics = syscall.GUID{0x3BDDCA50, 0x6D9D, 0x4531,
	[8]byte{0x8D, 0xC4, 0xB2, 0x7E, 0x3E, 0x49, 0xB7, 0xAB}}

type ISceneMetallicRoughnessMaterialStaticsInterface interface {
	win32.IInspectableInterface
	Create(compositor *ICompositor) *ISceneMetallicRoughnessMaterial
}

type ISceneMetallicRoughnessMaterialStaticsVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ISceneMetallicRoughnessMaterialStatics struct {
	win32.IInspectable
}

func (this *ISceneMetallicRoughnessMaterialStatics) Vtbl() *ISceneMetallicRoughnessMaterialStaticsVtbl {
	return (*ISceneMetallicRoughnessMaterialStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISceneMetallicRoughnessMaterialStatics) Create(compositor *ICompositor) *ISceneMetallicRoughnessMaterial {
	var _result *ISceneMetallicRoughnessMaterial
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(compositor)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C05576C2-32B1-4269-980D-B98537100AE4
var IID_ISceneModelTransform = syscall.GUID{0xC05576C2, 0x32B1, 0x4269,
	[8]byte{0x98, 0x0D, 0xB9, 0x85, 0x37, 0x10, 0x0A, 0xE4}}

type ISceneModelTransformInterface interface {
	win32.IInspectableInterface
	Get_Orientation() Quaternion
	Put_Orientation(value Quaternion)
	Get_RotationAngle() float32
	Put_RotationAngle(value float32)
	Get_RotationAngleInDegrees() float32
	Put_RotationAngleInDegrees(value float32)
	Get_RotationAxis() Vector3
	Put_RotationAxis(value Vector3)
	Get_Scale() Vector3
	Put_Scale(value Vector3)
	Get_Translation() Vector3
	Put_Translation(value Vector3)
}

type ISceneModelTransformVtbl struct {
	win32.IInspectableVtbl
	Get_Orientation            uintptr
	Put_Orientation            uintptr
	Get_RotationAngle          uintptr
	Put_RotationAngle          uintptr
	Get_RotationAngleInDegrees uintptr
	Put_RotationAngleInDegrees uintptr
	Get_RotationAxis           uintptr
	Put_RotationAxis           uintptr
	Get_Scale                  uintptr
	Put_Scale                  uintptr
	Get_Translation            uintptr
	Put_Translation            uintptr
}

type ISceneModelTransform struct {
	win32.IInspectable
}

func (this *ISceneModelTransform) Vtbl() *ISceneModelTransformVtbl {
	return (*ISceneModelTransformVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISceneModelTransform) Get_Orientation() Quaternion {
	var _result Quaternion
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Orientation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISceneModelTransform) Put_Orientation(value Quaternion) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Orientation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ISceneModelTransform) Get_RotationAngle() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RotationAngle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISceneModelTransform) Put_RotationAngle(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RotationAngle, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISceneModelTransform) Get_RotationAngleInDegrees() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RotationAngleInDegrees, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISceneModelTransform) Put_RotationAngleInDegrees(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RotationAngleInDegrees, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISceneModelTransform) Get_RotationAxis() Vector3 {
	var _result Vector3
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RotationAxis, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISceneModelTransform) Put_RotationAxis(value Vector3) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RotationAxis, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ISceneModelTransform) Get_Scale() Vector3 {
	var _result Vector3
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Scale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISceneModelTransform) Put_Scale(value Vector3) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Scale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ISceneModelTransform) Get_Translation() Vector3 {
	var _result Vector3
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Translation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISceneModelTransform) Put_Translation(value Vector3) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Translation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

// ACF2C247-F307-4581-9C41-AF2E29C3B016
var IID_ISceneNode = syscall.GUID{0xACF2C247, 0xF307, 0x4581,
	[8]byte{0x9C, 0x41, 0xAF, 0x2E, 0x29, 0xC3, 0xB0, 0x16}}

type ISceneNodeInterface interface {
	win32.IInspectableInterface
	Get_Children() *IVector[*ISceneNode]
	Get_Components() *IVector[*ISceneComponent]
	Get_Parent() *ISceneNode
	Get_Transform() *ISceneModelTransform
	FindFirstComponentOfType(value SceneComponentType) *ISceneComponent
}

type ISceneNodeVtbl struct {
	win32.IInspectableVtbl
	Get_Children             uintptr
	Get_Components           uintptr
	Get_Parent               uintptr
	Get_Transform            uintptr
	FindFirstComponentOfType uintptr
}

type ISceneNode struct {
	win32.IInspectable
}

func (this *ISceneNode) Vtbl() *ISceneNodeVtbl {
	return (*ISceneNodeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISceneNode) Get_Children() *IVector[*ISceneNode] {
	var _result *IVector[*ISceneNode]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Children, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISceneNode) Get_Components() *IVector[*ISceneComponent] {
	var _result *IVector[*ISceneComponent]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Components, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISceneNode) Get_Parent() *ISceneNode {
	var _result *ISceneNode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Parent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISceneNode) Get_Transform() *ISceneModelTransform {
	var _result *ISceneModelTransform
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Transform, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISceneNode) FindFirstComponentOfType(value SceneComponentType) *ISceneComponent {
	var _result *ISceneComponent
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindFirstComponentOfType, uintptr(unsafe.Pointer(this)), uintptr(value), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 29ADA101-2DD9-4332-BE63-60D2CF4269F2
var IID_ISceneNodeCollection = syscall.GUID{0x29ADA101, 0x2DD9, 0x4332,
	[8]byte{0xBE, 0x63, 0x60, 0xD2, 0xCF, 0x42, 0x69, 0xF2}}

type ISceneNodeCollectionInterface interface {
	win32.IInspectableInterface
}

type ISceneNodeCollectionVtbl struct {
	win32.IInspectableVtbl
}

type ISceneNodeCollection struct {
	win32.IInspectable
}

func (this *ISceneNodeCollection) Vtbl() *ISceneNodeCollectionVtbl {
	return (*ISceneNodeCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 579A0FAA-BE9D-4210-908C-93D15FEED0B7
var IID_ISceneNodeStatics = syscall.GUID{0x579A0FAA, 0xBE9D, 0x4210,
	[8]byte{0x90, 0x8C, 0x93, 0xD1, 0x5F, 0xEE, 0xD0, 0xB7}}

type ISceneNodeStaticsInterface interface {
	win32.IInspectableInterface
	Create(compositor *ICompositor) *ISceneNode
}

type ISceneNodeStaticsVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ISceneNodeStatics struct {
	win32.IInspectable
}

func (this *ISceneNodeStatics) Vtbl() *ISceneNodeStaticsVtbl {
	return (*ISceneNodeStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISceneNodeStatics) Create(compositor *ICompositor) *ISceneNode {
	var _result *ISceneNode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(compositor)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1E94249B-0F1B-49EB-A819-877D8450005B
var IID_ISceneObject = syscall.GUID{0x1E94249B, 0x0F1B, 0x49EB,
	[8]byte{0xA8, 0x19, 0x87, 0x7D, 0x84, 0x50, 0x00, 0x5B}}

type ISceneObjectInterface interface {
	win32.IInspectableInterface
}

type ISceneObjectVtbl struct {
	win32.IInspectableVtbl
}

type ISceneObject struct {
	win32.IInspectable
}

func (this *ISceneObject) Vtbl() *ISceneObjectVtbl {
	return (*ISceneObjectVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 14FE799A-33E4-52EF-956C-44229D21F2C1
var IID_ISceneObjectFactory = syscall.GUID{0x14FE799A, 0x33E4, 0x52EF,
	[8]byte{0x95, 0x6C, 0x44, 0x22, 0x9D, 0x21, 0xF2, 0xC1}}

type ISceneObjectFactoryInterface interface {
	win32.IInspectableInterface
}

type ISceneObjectFactoryVtbl struct {
	win32.IInspectableVtbl
}

type ISceneObjectFactory struct {
	win32.IInspectable
}

func (this *ISceneObjectFactory) Vtbl() *ISceneObjectFactoryVtbl {
	return (*ISceneObjectFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// AAB6EBBE-D680-46DF-8294-B6800A9F95E7
var IID_IScenePbrMaterial = syscall.GUID{0xAAB6EBBE, 0xD680, 0x46DF,
	[8]byte{0x82, 0x94, 0xB6, 0x80, 0x0A, 0x9F, 0x95, 0xE7}}

type IScenePbrMaterialInterface interface {
	win32.IInspectableInterface
	Get_AlphaCutoff() float32
	Put_AlphaCutoff(value float32)
	Get_AlphaMode() SceneAlphaMode
	Put_AlphaMode(value SceneAlphaMode)
	Get_EmissiveInput() *ISceneMaterialInput
	Put_EmissiveInput(value *ISceneMaterialInput)
	Get_EmissiveFactor() Vector3
	Put_EmissiveFactor(value Vector3)
	Get_IsDoubleSided() bool
	Put_IsDoubleSided(value bool)
	Get_NormalInput() *ISceneMaterialInput
	Put_NormalInput(value *ISceneMaterialInput)
	Get_NormalScale() float32
	Put_NormalScale(value float32)
	Get_OcclusionInput() *ISceneMaterialInput
	Put_OcclusionInput(value *ISceneMaterialInput)
	Get_OcclusionStrength() float32
	Put_OcclusionStrength(value float32)
}

type IScenePbrMaterialVtbl struct {
	win32.IInspectableVtbl
	Get_AlphaCutoff       uintptr
	Put_AlphaCutoff       uintptr
	Get_AlphaMode         uintptr
	Put_AlphaMode         uintptr
	Get_EmissiveInput     uintptr
	Put_EmissiveInput     uintptr
	Get_EmissiveFactor    uintptr
	Put_EmissiveFactor    uintptr
	Get_IsDoubleSided     uintptr
	Put_IsDoubleSided     uintptr
	Get_NormalInput       uintptr
	Put_NormalInput       uintptr
	Get_NormalScale       uintptr
	Put_NormalScale       uintptr
	Get_OcclusionInput    uintptr
	Put_OcclusionInput    uintptr
	Get_OcclusionStrength uintptr
	Put_OcclusionStrength uintptr
}

type IScenePbrMaterial struct {
	win32.IInspectable
}

func (this *IScenePbrMaterial) Vtbl() *IScenePbrMaterialVtbl {
	return (*IScenePbrMaterialVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IScenePbrMaterial) Get_AlphaCutoff() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlphaCutoff, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IScenePbrMaterial) Put_AlphaCutoff(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AlphaCutoff, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IScenePbrMaterial) Get_AlphaMode() SceneAlphaMode {
	var _result SceneAlphaMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlphaMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IScenePbrMaterial) Put_AlphaMode(value SceneAlphaMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AlphaMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IScenePbrMaterial) Get_EmissiveInput() *ISceneMaterialInput {
	var _result *ISceneMaterialInput
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EmissiveInput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IScenePbrMaterial) Put_EmissiveInput(value *ISceneMaterialInput) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_EmissiveInput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IScenePbrMaterial) Get_EmissiveFactor() Vector3 {
	var _result Vector3
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EmissiveFactor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IScenePbrMaterial) Put_EmissiveFactor(value Vector3) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_EmissiveFactor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IScenePbrMaterial) Get_IsDoubleSided() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsDoubleSided, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IScenePbrMaterial) Put_IsDoubleSided(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsDoubleSided, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IScenePbrMaterial) Get_NormalInput() *ISceneMaterialInput {
	var _result *ISceneMaterialInput
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NormalInput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IScenePbrMaterial) Put_NormalInput(value *ISceneMaterialInput) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_NormalInput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IScenePbrMaterial) Get_NormalScale() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NormalScale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IScenePbrMaterial) Put_NormalScale(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_NormalScale, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IScenePbrMaterial) Get_OcclusionInput() *ISceneMaterialInput {
	var _result *ISceneMaterialInput
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OcclusionInput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IScenePbrMaterial) Put_OcclusionInput(value *ISceneMaterialInput) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_OcclusionInput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IScenePbrMaterial) Get_OcclusionStrength() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OcclusionStrength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IScenePbrMaterial) Put_OcclusionStrength(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_OcclusionStrength, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 2E3F3DFE-0B85-5727-B5BE-B7D3CBAC37FA
var IID_IScenePbrMaterialFactory = syscall.GUID{0x2E3F3DFE, 0x0B85, 0x5727,
	[8]byte{0xB5, 0xBE, 0xB7, 0xD3, 0xCB, 0xAC, 0x37, 0xFA}}

type IScenePbrMaterialFactoryInterface interface {
	win32.IInspectableInterface
}

type IScenePbrMaterialFactoryVtbl struct {
	win32.IInspectableVtbl
}

type IScenePbrMaterialFactory struct {
	win32.IInspectable
}

func (this *IScenePbrMaterialFactory) Vtbl() *IScenePbrMaterialFactoryVtbl {
	return (*IScenePbrMaterialFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// F1ACB857-CF4F-4025-9B25-A2D1944CF507
var IID_ISceneRendererComponent = syscall.GUID{0xF1ACB857, 0xCF4F, 0x4025,
	[8]byte{0x9B, 0x25, 0xA2, 0xD1, 0x94, 0x4C, 0xF5, 0x07}}

type ISceneRendererComponentInterface interface {
	win32.IInspectableInterface
}

type ISceneRendererComponentVtbl struct {
	win32.IInspectableVtbl
}

type ISceneRendererComponent struct {
	win32.IInspectable
}

func (this *ISceneRendererComponent) Vtbl() *ISceneRendererComponentVtbl {
	return (*ISceneRendererComponentVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 1DB6ED6C-AA2C-5967-9035-56352DC69658
var IID_ISceneRendererComponentFactory = syscall.GUID{0x1DB6ED6C, 0xAA2C, 0x5967,
	[8]byte{0x90, 0x35, 0x56, 0x35, 0x2D, 0xC6, 0x96, 0x58}}

type ISceneRendererComponentFactoryInterface interface {
	win32.IInspectableInterface
}

type ISceneRendererComponentFactoryVtbl struct {
	win32.IInspectableVtbl
}

type ISceneRendererComponentFactory struct {
	win32.IInspectable
}

func (this *ISceneRendererComponentFactory) Vtbl() *ISceneRendererComponentFactoryVtbl {
	return (*ISceneRendererComponentFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 9937DA5C-A9CA-4CFC-B3AA-088356518742
var IID_ISceneSurfaceMaterialInput = syscall.GUID{0x9937DA5C, 0xA9CA, 0x4CFC,
	[8]byte{0xB3, 0xAA, 0x08, 0x83, 0x56, 0x51, 0x87, 0x42}}

type ISceneSurfaceMaterialInputInterface interface {
	win32.IInspectableInterface
	Get_BitmapInterpolationMode() CompositionBitmapInterpolationMode
	Put_BitmapInterpolationMode(value CompositionBitmapInterpolationMode)
	Get_Surface() *ICompositionSurface
	Put_Surface(value *ICompositionSurface)
	Get_WrappingUMode() SceneWrappingMode
	Put_WrappingUMode(value SceneWrappingMode)
	Get_WrappingVMode() SceneWrappingMode
	Put_WrappingVMode(value SceneWrappingMode)
}

type ISceneSurfaceMaterialInputVtbl struct {
	win32.IInspectableVtbl
	Get_BitmapInterpolationMode uintptr
	Put_BitmapInterpolationMode uintptr
	Get_Surface                 uintptr
	Put_Surface                 uintptr
	Get_WrappingUMode           uintptr
	Put_WrappingUMode           uintptr
	Get_WrappingVMode           uintptr
	Put_WrappingVMode           uintptr
}

type ISceneSurfaceMaterialInput struct {
	win32.IInspectable
}

func (this *ISceneSurfaceMaterialInput) Vtbl() *ISceneSurfaceMaterialInputVtbl {
	return (*ISceneSurfaceMaterialInputVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISceneSurfaceMaterialInput) Get_BitmapInterpolationMode() CompositionBitmapInterpolationMode {
	var _result CompositionBitmapInterpolationMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BitmapInterpolationMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISceneSurfaceMaterialInput) Put_BitmapInterpolationMode(value CompositionBitmapInterpolationMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BitmapInterpolationMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISceneSurfaceMaterialInput) Get_Surface() *ICompositionSurface {
	var _result *ICompositionSurface
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Surface, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISceneSurfaceMaterialInput) Put_Surface(value *ICompositionSurface) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Surface, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *ISceneSurfaceMaterialInput) Get_WrappingUMode() SceneWrappingMode {
	var _result SceneWrappingMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WrappingUMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISceneSurfaceMaterialInput) Put_WrappingUMode(value SceneWrappingMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_WrappingUMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *ISceneSurfaceMaterialInput) Get_WrappingVMode() SceneWrappingMode {
	var _result SceneWrappingMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WrappingVMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISceneSurfaceMaterialInput) Put_WrappingVMode(value SceneWrappingMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_WrappingVMode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 5A2394D3-6429-4589-BBCF-B84F4F3CFBFE
var IID_ISceneSurfaceMaterialInputStatics = syscall.GUID{0x5A2394D3, 0x6429, 0x4589,
	[8]byte{0xBB, 0xCF, 0xB8, 0x4F, 0x4F, 0x3C, 0xFB, 0xFE}}

type ISceneSurfaceMaterialInputStaticsInterface interface {
	win32.IInspectableInterface
	Create(compositor *ICompositor) *ISceneSurfaceMaterialInput
}

type ISceneSurfaceMaterialInputStaticsVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ISceneSurfaceMaterialInputStatics struct {
	win32.IInspectable
}

func (this *ISceneSurfaceMaterialInputStatics) Vtbl() *ISceneSurfaceMaterialInputStaticsVtbl {
	return (*ISceneSurfaceMaterialInputStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISceneSurfaceMaterialInputStatics) Create(compositor *ICompositor) *ISceneSurfaceMaterialInput {
	var _result *ISceneSurfaceMaterialInput
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(compositor)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8E672C1E-D734-47B1-BE14-3D694FFA4301
var IID_ISceneVisual = syscall.GUID{0x8E672C1E, 0xD734, 0x47B1,
	[8]byte{0xBE, 0x14, 0x3D, 0x69, 0x4F, 0xFA, 0x43, 0x01}}

type ISceneVisualInterface interface {
	win32.IInspectableInterface
	Get_Root() *ISceneNode
	Put_Root(value *ISceneNode)
}

type ISceneVisualVtbl struct {
	win32.IInspectableVtbl
	Get_Root uintptr
	Put_Root uintptr
}

type ISceneVisual struct {
	win32.IInspectable
}

func (this *ISceneVisual) Vtbl() *ISceneVisualVtbl {
	return (*ISceneVisualVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISceneVisual) Get_Root() *ISceneNode {
	var _result *ISceneNode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Root, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISceneVisual) Put_Root(value *ISceneNode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Root, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// B8347E9A-50AA-4527-8D34-DE4CB8EA88B4
var IID_ISceneVisualStatics = syscall.GUID{0xB8347E9A, 0x50AA, 0x4527,
	[8]byte{0x8D, 0x34, 0xDE, 0x4C, 0xB8, 0xEA, 0x88, 0xB4}}

type ISceneVisualStaticsInterface interface {
	win32.IInspectableInterface
	Create(compositor *ICompositor) *ISceneVisual
}

type ISceneVisualStaticsVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ISceneVisualStatics struct {
	win32.IInspectable
}

func (this *ISceneVisualStatics) Vtbl() *ISceneVisualStaticsVtbl {
	return (*ISceneVisualStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISceneVisualStatics) Create(compositor *ICompositor) *ISceneVisual {
	var _result *ISceneVisual
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(compositor)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type SceneObject struct {
	RtClass
	*ISceneObject
}

type SceneBoundingBox struct {
	RtClass
	*ISceneBoundingBox
}

type SceneComponent struct {
	RtClass
	*ISceneComponent
}

type SceneComponentCollection struct {
	RtClass
	*IVector[*ISceneComponent]
}

type SceneMaterial struct {
	RtClass
	*ISceneMaterial
}

type SceneMaterialInput struct {
	RtClass
	*ISceneMaterialInput
}

type SceneMesh struct {
	RtClass
	*ISceneMesh
}

func NewISceneMeshStatics() *ISceneMeshStatics {
	var p *ISceneMeshStatics
	hs := NewHStr("Windows.UI.Composition.Scenes.SceneMesh")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISceneMeshStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type SceneMeshMaterialAttributeMap struct {
	RtClass
	*ISceneMeshMaterialAttributeMap
}

type SceneRendererComponent struct {
	RtClass
	*ISceneRendererComponent
}

type SceneMeshRendererComponent struct {
	RtClass
	*ISceneMeshRendererComponent
}

func NewISceneMeshRendererComponentStatics() *ISceneMeshRendererComponentStatics {
	var p *ISceneMeshRendererComponentStatics
	hs := NewHStr("Windows.UI.Composition.Scenes.SceneMeshRendererComponent")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISceneMeshRendererComponentStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type ScenePbrMaterial struct {
	RtClass
	*IScenePbrMaterial
}

type SceneMetallicRoughnessMaterial struct {
	RtClass
	*ISceneMetallicRoughnessMaterial
}

func NewISceneMetallicRoughnessMaterialStatics() *ISceneMetallicRoughnessMaterialStatics {
	var p *ISceneMetallicRoughnessMaterialStatics
	hs := NewHStr("Windows.UI.Composition.Scenes.SceneMetallicRoughnessMaterial")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISceneMetallicRoughnessMaterialStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type SceneModelTransform struct {
	RtClass
	*ISceneModelTransform
}

type SceneNode struct {
	RtClass
	*ISceneNode
}

func NewISceneNodeStatics() *ISceneNodeStatics {
	var p *ISceneNodeStatics
	hs := NewHStr("Windows.UI.Composition.Scenes.SceneNode")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISceneNodeStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type SceneNodeCollection struct {
	RtClass
	*IVector[*ISceneNode]
}

type SceneSurfaceMaterialInput struct {
	RtClass
	*ISceneSurfaceMaterialInput
}

func NewISceneSurfaceMaterialInputStatics() *ISceneSurfaceMaterialInputStatics {
	var p *ISceneSurfaceMaterialInputStatics
	hs := NewHStr("Windows.UI.Composition.Scenes.SceneSurfaceMaterialInput")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISceneSurfaceMaterialInputStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type SceneVisual struct {
	RtClass
	*ISceneVisual
}

func NewISceneVisualStatics() *ISceneVisualStatics {
	var p *ISceneVisualStatics
	hs := NewHStr("Windows.UI.Composition.Scenes.SceneVisual")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISceneVisualStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
