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
type LearningModelDeviceKind int32

const (
	LearningModelDeviceKind_Default                LearningModelDeviceKind = 0
	LearningModelDeviceKind_Cpu                    LearningModelDeviceKind = 1
	LearningModelDeviceKind_DirectX                LearningModelDeviceKind = 2
	LearningModelDeviceKind_DirectXHighPerformance LearningModelDeviceKind = 3
	LearningModelDeviceKind_DirectXMinPower        LearningModelDeviceKind = 4
)

// enum
type LearningModelFeatureKind int32

const (
	LearningModelFeatureKind_Tensor   LearningModelFeatureKind = 0
	LearningModelFeatureKind_Sequence LearningModelFeatureKind = 1
	LearningModelFeatureKind_Map      LearningModelFeatureKind = 2
	LearningModelFeatureKind_Image    LearningModelFeatureKind = 3
)

// enum
type TensorKind int32

const (
	TensorKind_Undefined  TensorKind = 0
	TensorKind_Float      TensorKind = 1
	TensorKind_UInt8      TensorKind = 2
	TensorKind_Int8       TensorKind = 3
	TensorKind_UInt16     TensorKind = 4
	TensorKind_Int16      TensorKind = 5
	TensorKind_Int32      TensorKind = 6
	TensorKind_Int64      TensorKind = 7
	TensorKind_String     TensorKind = 8
	TensorKind_Boolean    TensorKind = 9
	TensorKind_Float16    TensorKind = 10
	TensorKind_Double     TensorKind = 11
	TensorKind_UInt32     TensorKind = 12
	TensorKind_UInt64     TensorKind = 13
	TensorKind_Complex64  TensorKind = 14
	TensorKind_Complex128 TensorKind = 15
)

// structs

type MachineLearningContract struct {
}

// interfaces

// 365585A5-171A-4A2A-985F-265159D3895A
var IID_IImageFeatureDescriptor = syscall.GUID{0x365585A5, 0x171A, 0x4A2A,
	[8]byte{0x98, 0x5F, 0x26, 0x51, 0x59, 0xD3, 0x89, 0x5A}}

type IImageFeatureDescriptorInterface interface {
	win32.IInspectableInterface
	Get_BitmapPixelFormat() BitmapPixelFormat
	Get_BitmapAlphaMode() BitmapAlphaMode
	Get_Width() uint32
	Get_Height() uint32
}

type IImageFeatureDescriptorVtbl struct {
	win32.IInspectableVtbl
	Get_BitmapPixelFormat uintptr
	Get_BitmapAlphaMode   uintptr
	Get_Width             uintptr
	Get_Height            uintptr
}

type IImageFeatureDescriptor struct {
	win32.IInspectable
}

func (this *IImageFeatureDescriptor) Vtbl() *IImageFeatureDescriptorVtbl {
	return (*IImageFeatureDescriptorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IImageFeatureDescriptor) Get_BitmapPixelFormat() BitmapPixelFormat {
	var _result BitmapPixelFormat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BitmapPixelFormat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageFeatureDescriptor) Get_BitmapAlphaMode() BitmapAlphaMode {
	var _result BitmapAlphaMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BitmapAlphaMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageFeatureDescriptor) Get_Width() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Width, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageFeatureDescriptor) Get_Height() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Height, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F0414FD9-C9AA-4405-B7FB-94F87C8A3037
var IID_IImageFeatureValue = syscall.GUID{0xF0414FD9, 0xC9AA, 0x4405,
	[8]byte{0xB7, 0xFB, 0x94, 0xF8, 0x7C, 0x8A, 0x30, 0x37}}

type IImageFeatureValueInterface interface {
	win32.IInspectableInterface
	Get_VideoFrame() *IVideoFrame
}

type IImageFeatureValueVtbl struct {
	win32.IInspectableVtbl
	Get_VideoFrame uintptr
}

type IImageFeatureValue struct {
	win32.IInspectable
}

func (this *IImageFeatureValue) Vtbl() *IImageFeatureValueVtbl {
	return (*IImageFeatureValueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IImageFeatureValue) Get_VideoFrame() *IVideoFrame {
	var _result *IVideoFrame
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_VideoFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1BC317FD-23CB-4610-B085-C8E1C87EBAA0
var IID_IImageFeatureValueStatics = syscall.GUID{0x1BC317FD, 0x23CB, 0x4610,
	[8]byte{0xB0, 0x85, 0xC8, 0xE1, 0xC8, 0x7E, 0xBA, 0xA0}}

type IImageFeatureValueStaticsInterface interface {
	win32.IInspectableInterface
	CreateFromVideoFrame(image *IVideoFrame) *IImageFeatureValue
}

type IImageFeatureValueStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateFromVideoFrame uintptr
}

type IImageFeatureValueStatics struct {
	win32.IInspectable
}

func (this *IImageFeatureValueStatics) Vtbl() *IImageFeatureValueStaticsVtbl {
	return (*IImageFeatureValueStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IImageFeatureValueStatics) CreateFromVideoFrame(image *IVideoFrame) *IImageFeatureValue {
	var _result *IImageFeatureValue
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromVideoFrame, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(image)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5B8E4920-489F-4E86-9128-265A327B78FA
var IID_ILearningModel = syscall.GUID{0x5B8E4920, 0x489F, 0x4E86,
	[8]byte{0x91, 0x28, 0x26, 0x5A, 0x32, 0x7B, 0x78, 0xFA}}

type ILearningModelInterface interface {
	win32.IInspectableInterface
	Get_Author() string
	Get_Name() string
	Get_Domain() string
	Get_Description() string
	Get_Version() int64
	Get_Metadata() *IMapView[string, string]
	Get_InputFeatures() *IVectorView[*ILearningModelFeatureDescriptor]
	Get_OutputFeatures() *IVectorView[*ILearningModelFeatureDescriptor]
}

type ILearningModelVtbl struct {
	win32.IInspectableVtbl
	Get_Author         uintptr
	Get_Name           uintptr
	Get_Domain         uintptr
	Get_Description    uintptr
	Get_Version        uintptr
	Get_Metadata       uintptr
	Get_InputFeatures  uintptr
	Get_OutputFeatures uintptr
}

type ILearningModel struct {
	win32.IInspectable
}

func (this *ILearningModel) Vtbl() *ILearningModelVtbl {
	return (*ILearningModelVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILearningModel) Get_Author() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Author, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILearningModel) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILearningModel) Get_Domain() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Domain, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILearningModel) Get_Description() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Description, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILearningModel) Get_Version() int64 {
	var _result int64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Version, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILearningModel) Get_Metadata() *IMapView[string, string] {
	var _result *IMapView[string, string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Metadata, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILearningModel) Get_InputFeatures() *IVectorView[*ILearningModelFeatureDescriptor] {
	var _result *IVectorView[*ILearningModelFeatureDescriptor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InputFeatures, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILearningModel) Get_OutputFeatures() *IVectorView[*ILearningModelFeatureDescriptor] {
	var _result *IVectorView[*ILearningModelFeatureDescriptor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OutputFeatures, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EA312F20-168F-4F8C-94FE-2E7AC31B4AA8
var IID_ILearningModelBinding = syscall.GUID{0xEA312F20, 0x168F, 0x4F8C,
	[8]byte{0x94, 0xFE, 0x2E, 0x7A, 0xC3, 0x1B, 0x4A, 0xA8}}

type ILearningModelBindingInterface interface {
	win32.IInspectableInterface
	Bind(name string, value interface{})
	BindWithProperties(name string, value interface{}, props *IPropertySet)
	Clear()
}

type ILearningModelBindingVtbl struct {
	win32.IInspectableVtbl
	Bind               uintptr
	BindWithProperties uintptr
	Clear              uintptr
}

type ILearningModelBinding struct {
	win32.IInspectable
}

func (this *ILearningModelBinding) Vtbl() *ILearningModelBindingVtbl {
	return (*ILearningModelBindingVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILearningModelBinding) Bind(name string, value interface{}) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Bind, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ILearningModelBinding) BindWithProperties(name string, value interface{}, props *IPropertySet) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().BindWithProperties, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, *(*uintptr)(unsafe.Pointer(&value)), uintptr(unsafe.Pointer(props)))
	_ = _hr
}

func (this *ILearningModelBinding) Clear() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Clear, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// C95F7A7A-E788-475E-8917-23AA381FAF0B
var IID_ILearningModelBindingFactory = syscall.GUID{0xC95F7A7A, 0xE788, 0x475E,
	[8]byte{0x89, 0x17, 0x23, 0xAA, 0x38, 0x1F, 0xAF, 0x0B}}

type ILearningModelBindingFactoryInterface interface {
	win32.IInspectableInterface
	CreateFromSession(session *ILearningModelSession) *ILearningModelBinding
}

type ILearningModelBindingFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateFromSession uintptr
}

type ILearningModelBindingFactory struct {
	win32.IInspectable
}

func (this *ILearningModelBindingFactory) Vtbl() *ILearningModelBindingFactoryVtbl {
	return (*ILearningModelBindingFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILearningModelBindingFactory) CreateFromSession(session *ILearningModelSession) *ILearningModelBinding {
	var _result *ILearningModelBinding
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromSession, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(session)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F5C2C8FE-3F56-4A8C-AC5F-FDB92D8B8252
var IID_ILearningModelDevice = syscall.GUID{0xF5C2C8FE, 0x3F56, 0x4A8C,
	[8]byte{0xAC, 0x5F, 0xFD, 0xB9, 0x2D, 0x8B, 0x82, 0x52}}

type ILearningModelDeviceInterface interface {
	win32.IInspectableInterface
	Get_AdapterId() unsafe.Pointer
	Get_Direct3D11Device() unsafe.Pointer
}

type ILearningModelDeviceVtbl struct {
	win32.IInspectableVtbl
	Get_AdapterId        uintptr
	Get_Direct3D11Device uintptr
}

type ILearningModelDevice struct {
	win32.IInspectable
}

func (this *ILearningModelDevice) Vtbl() *ILearningModelDeviceVtbl {
	return (*ILearningModelDeviceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILearningModelDevice) Get_AdapterId() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AdapterId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILearningModelDevice) Get_Direct3D11Device() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Direct3D11Device, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 9CFFD74D-B1E5-4F20-80AD-0A56690DB06B
var IID_ILearningModelDeviceFactory = syscall.GUID{0x9CFFD74D, 0xB1E5, 0x4F20,
	[8]byte{0x80, 0xAD, 0x0A, 0x56, 0x69, 0x0D, 0xB0, 0x6B}}

type ILearningModelDeviceFactoryInterface interface {
	win32.IInspectableInterface
	Create(deviceKind LearningModelDeviceKind) *ILearningModelDevice
}

type ILearningModelDeviceFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ILearningModelDeviceFactory struct {
	win32.IInspectable
}

func (this *ILearningModelDeviceFactory) Vtbl() *ILearningModelDeviceFactoryVtbl {
	return (*ILearningModelDeviceFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILearningModelDeviceFactory) Create(deviceKind LearningModelDeviceKind) *ILearningModelDevice {
	var _result *ILearningModelDevice
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(deviceKind), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 49F32107-A8BF-42BB-92C7-10B12DC5D21F
var IID_ILearningModelDeviceStatics = syscall.GUID{0x49F32107, 0xA8BF, 0x42BB,
	[8]byte{0x92, 0xC7, 0x10, 0xB1, 0x2D, 0xC5, 0xD2, 0x1F}}

type ILearningModelDeviceStaticsInterface interface {
	win32.IInspectableInterface
	CreateFromDirect3D11Device(device unsafe.Pointer) *ILearningModelDevice
}

type ILearningModelDeviceStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateFromDirect3D11Device uintptr
}

type ILearningModelDeviceStatics struct {
	win32.IInspectable
}

func (this *ILearningModelDeviceStatics) Vtbl() *ILearningModelDeviceStaticsVtbl {
	return (*ILearningModelDeviceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILearningModelDeviceStatics) CreateFromDirect3D11Device(device unsafe.Pointer) *ILearningModelDevice {
	var _result *ILearningModelDevice
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromDirect3D11Device, uintptr(unsafe.Pointer(this)), uintptr(device), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B2F9BFCD-960E-49C0-8593-EB190AE3EEE2
var IID_ILearningModelEvaluationResult = syscall.GUID{0xB2F9BFCD, 0x960E, 0x49C0,
	[8]byte{0x85, 0x93, 0xEB, 0x19, 0x0A, 0xE3, 0xEE, 0xE2}}

type ILearningModelEvaluationResultInterface interface {
	win32.IInspectableInterface
	Get_CorrelationId() string
	Get_ErrorStatus() int32
	Get_Succeeded() bool
	Get_Outputs() *IMapView[string, interface{}]
}

type ILearningModelEvaluationResultVtbl struct {
	win32.IInspectableVtbl
	Get_CorrelationId uintptr
	Get_ErrorStatus   uintptr
	Get_Succeeded     uintptr
	Get_Outputs       uintptr
}

type ILearningModelEvaluationResult struct {
	win32.IInspectable
}

func (this *ILearningModelEvaluationResult) Vtbl() *ILearningModelEvaluationResultVtbl {
	return (*ILearningModelEvaluationResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILearningModelEvaluationResult) Get_CorrelationId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CorrelationId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILearningModelEvaluationResult) Get_ErrorStatus() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ErrorStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILearningModelEvaluationResult) Get_Succeeded() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Succeeded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILearningModelEvaluationResult) Get_Outputs() *IMapView[string, interface{}] {
	var _result *IMapView[string, interface{}]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Outputs, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BC08CF7C-6ED0-4004-97BA-B9A2EECD2B4F
var IID_ILearningModelFeatureDescriptor = syscall.GUID{0xBC08CF7C, 0x6ED0, 0x4004,
	[8]byte{0x97, 0xBA, 0xB9, 0xA2, 0xEE, 0xCD, 0x2B, 0x4F}}

type ILearningModelFeatureDescriptorInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Get_Description() string
	Get_Kind() LearningModelFeatureKind
	Get_IsRequired() bool
}

type ILearningModelFeatureDescriptorVtbl struct {
	win32.IInspectableVtbl
	Get_Name        uintptr
	Get_Description uintptr
	Get_Kind        uintptr
	Get_IsRequired  uintptr
}

type ILearningModelFeatureDescriptor struct {
	win32.IInspectable
}

func (this *ILearningModelFeatureDescriptor) Vtbl() *ILearningModelFeatureDescriptorVtbl {
	return (*ILearningModelFeatureDescriptorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILearningModelFeatureDescriptor) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILearningModelFeatureDescriptor) Get_Description() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Description, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ILearningModelFeatureDescriptor) Get_Kind() LearningModelFeatureKind {
	var _result LearningModelFeatureKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Kind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILearningModelFeatureDescriptor) Get_IsRequired() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsRequired, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// F51005DB-4085-4DFE-9FED-95EB0C0CF75C
var IID_ILearningModelFeatureValue = syscall.GUID{0xF51005DB, 0x4085, 0x4DFE,
	[8]byte{0x9F, 0xED, 0x95, 0xEB, 0x0C, 0x0C, 0xF7, 0x5C}}

type ILearningModelFeatureValueInterface interface {
	win32.IInspectableInterface
	Get_Kind() LearningModelFeatureKind
}

type ILearningModelFeatureValueVtbl struct {
	win32.IInspectableVtbl
	Get_Kind uintptr
}

type ILearningModelFeatureValue struct {
	win32.IInspectable
}

func (this *ILearningModelFeatureValue) Vtbl() *ILearningModelFeatureValueVtbl {
	return (*ILearningModelFeatureValueVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILearningModelFeatureValue) Get_Kind() LearningModelFeatureKind {
	var _result LearningModelFeatureKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Kind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 2A222E5D-AFB1-47ED-BFAD-B5B3A459EC04
var IID_ILearningModelOperatorProvider = syscall.GUID{0x2A222E5D, 0xAFB1, 0x47ED,
	[8]byte{0xBF, 0xAD, 0xB5, 0xB3, 0xA4, 0x59, 0xEC, 0x04}}

type ILearningModelOperatorProviderInterface interface {
	win32.IInspectableInterface
}

type ILearningModelOperatorProviderVtbl struct {
	win32.IInspectableVtbl
}

type ILearningModelOperatorProvider struct {
	win32.IInspectable
}

func (this *ILearningModelOperatorProvider) Vtbl() *ILearningModelOperatorProviderVtbl {
	return (*ILearningModelOperatorProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 8E58F8F6-B787-4C11-90F0-7129AECA74A9
var IID_ILearningModelSession = syscall.GUID{0x8E58F8F6, 0xB787, 0x4C11,
	[8]byte{0x90, 0xF0, 0x71, 0x29, 0xAE, 0xCA, 0x74, 0xA9}}

type ILearningModelSessionInterface interface {
	win32.IInspectableInterface
	Get_Model() *ILearningModel
	Get_Device() *ILearningModelDevice
	Get_EvaluationProperties() *IPropertySet
	EvaluateAsync(bindings *ILearningModelBinding, correlationId string) *IAsyncOperation[*ILearningModelEvaluationResult]
	EvaluateFeaturesAsync(features *IMap[string, interface{}], correlationId string) *IAsyncOperation[*ILearningModelEvaluationResult]
	Evaluate(bindings *ILearningModelBinding, correlationId string) *ILearningModelEvaluationResult
	EvaluateFeatures(features *IMap[string, interface{}], correlationId string) *ILearningModelEvaluationResult
}

type ILearningModelSessionVtbl struct {
	win32.IInspectableVtbl
	Get_Model                uintptr
	Get_Device               uintptr
	Get_EvaluationProperties uintptr
	EvaluateAsync            uintptr
	EvaluateFeaturesAsync    uintptr
	Evaluate                 uintptr
	EvaluateFeatures         uintptr
}

type ILearningModelSession struct {
	win32.IInspectable
}

func (this *ILearningModelSession) Vtbl() *ILearningModelSessionVtbl {
	return (*ILearningModelSessionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILearningModelSession) Get_Model() *ILearningModel {
	var _result *ILearningModel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Model, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILearningModelSession) Get_Device() *ILearningModelDevice {
	var _result *ILearningModelDevice
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Device, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILearningModelSession) Get_EvaluationProperties() *IPropertySet {
	var _result *IPropertySet
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EvaluationProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILearningModelSession) EvaluateAsync(bindings *ILearningModelBinding, correlationId string) *IAsyncOperation[*ILearningModelEvaluationResult] {
	var _result *IAsyncOperation[*ILearningModelEvaluationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EvaluateAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bindings)), NewHStr(correlationId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILearningModelSession) EvaluateFeaturesAsync(features *IMap[string, interface{}], correlationId string) *IAsyncOperation[*ILearningModelEvaluationResult] {
	var _result *IAsyncOperation[*ILearningModelEvaluationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EvaluateFeaturesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(features)), NewHStr(correlationId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILearningModelSession) Evaluate(bindings *ILearningModelBinding, correlationId string) *ILearningModelEvaluationResult {
	var _result *ILearningModelEvaluationResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Evaluate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bindings)), NewHStr(correlationId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILearningModelSession) EvaluateFeatures(features *IMap[string, interface{}], correlationId string) *ILearningModelEvaluationResult {
	var _result *ILearningModelEvaluationResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EvaluateFeatures, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(features)), NewHStr(correlationId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0F6B881D-1C9B-47B6-BFE0-F1CF62A67579
var IID_ILearningModelSessionFactory = syscall.GUID{0x0F6B881D, 0x1C9B, 0x47B6,
	[8]byte{0xBF, 0xE0, 0xF1, 0xCF, 0x62, 0xA6, 0x75, 0x79}}

type ILearningModelSessionFactoryInterface interface {
	win32.IInspectableInterface
	CreateFromModel(model *ILearningModel) *ILearningModelSession
	CreateFromModelOnDevice(model *ILearningModel, deviceToRunOn *ILearningModelDevice) *ILearningModelSession
}

type ILearningModelSessionFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateFromModel         uintptr
	CreateFromModelOnDevice uintptr
}

type ILearningModelSessionFactory struct {
	win32.IInspectable
}

func (this *ILearningModelSessionFactory) Vtbl() *ILearningModelSessionFactoryVtbl {
	return (*ILearningModelSessionFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILearningModelSessionFactory) CreateFromModel(model *ILearningModel) *ILearningModelSession {
	var _result *ILearningModelSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromModel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(model)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILearningModelSessionFactory) CreateFromModelOnDevice(model *ILearningModel, deviceToRunOn *ILearningModelDevice) *ILearningModelSession {
	var _result *ILearningModelSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromModelOnDevice, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(model)), uintptr(unsafe.Pointer(deviceToRunOn)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4E5C88BF-0A1F-5FEC-ADE0-2FD91E4EF29B
var IID_ILearningModelSessionFactory2 = syscall.GUID{0x4E5C88BF, 0x0A1F, 0x5FEC,
	[8]byte{0xAD, 0xE0, 0x2F, 0xD9, 0x1E, 0x4E, 0xF2, 0x9B}}

type ILearningModelSessionFactory2Interface interface {
	win32.IInspectableInterface
	CreateFromModelOnDeviceWithSessionOptions(model *ILearningModel, deviceToRunOn *ILearningModelDevice, learningModelSessionOptions *ILearningModelSessionOptions) *ILearningModelSession
}

type ILearningModelSessionFactory2Vtbl struct {
	win32.IInspectableVtbl
	CreateFromModelOnDeviceWithSessionOptions uintptr
}

type ILearningModelSessionFactory2 struct {
	win32.IInspectable
}

func (this *ILearningModelSessionFactory2) Vtbl() *ILearningModelSessionFactory2Vtbl {
	return (*ILearningModelSessionFactory2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILearningModelSessionFactory2) CreateFromModelOnDeviceWithSessionOptions(model *ILearningModel, deviceToRunOn *ILearningModelDevice, learningModelSessionOptions *ILearningModelSessionOptions) *ILearningModelSession {
	var _result *ILearningModelSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromModelOnDeviceWithSessionOptions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(model)), uintptr(unsafe.Pointer(deviceToRunOn)), uintptr(unsafe.Pointer(learningModelSessionOptions)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B8F63FA1-134D-5133-8CFF-3A5C3C263BEB
var IID_ILearningModelSessionOptions = syscall.GUID{0xB8F63FA1, 0x134D, 0x5133,
	[8]byte{0x8C, 0xFF, 0x3A, 0x5C, 0x3C, 0x26, 0x3B, 0xEB}}

type ILearningModelSessionOptionsInterface interface {
	win32.IInspectableInterface
	Get_BatchSizeOverride() uint32
	Put_BatchSizeOverride(value uint32)
}

type ILearningModelSessionOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_BatchSizeOverride uintptr
	Put_BatchSizeOverride uintptr
}

type ILearningModelSessionOptions struct {
	win32.IInspectable
}

func (this *ILearningModelSessionOptions) Vtbl() *ILearningModelSessionOptionsVtbl {
	return (*ILearningModelSessionOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILearningModelSessionOptions) Get_BatchSizeOverride() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BatchSizeOverride, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILearningModelSessionOptions) Put_BatchSizeOverride(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BatchSizeOverride, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 6FCD1DC4-175F-5BD2-8DE5-2F2006A25ADF
var IID_ILearningModelSessionOptions2 = syscall.GUID{0x6FCD1DC4, 0x175F, 0x5BD2,
	[8]byte{0x8D, 0xE5, 0x2F, 0x20, 0x06, 0xA2, 0x5A, 0xDF}}

type ILearningModelSessionOptions2Interface interface {
	win32.IInspectableInterface
	Get_CloseModelOnSessionCreation() bool
	Put_CloseModelOnSessionCreation(value bool)
}

type ILearningModelSessionOptions2Vtbl struct {
	win32.IInspectableVtbl
	Get_CloseModelOnSessionCreation uintptr
	Put_CloseModelOnSessionCreation uintptr
}

type ILearningModelSessionOptions2 struct {
	win32.IInspectable
}

func (this *ILearningModelSessionOptions2) Vtbl() *ILearningModelSessionOptions2Vtbl {
	return (*ILearningModelSessionOptions2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILearningModelSessionOptions2) Get_CloseModelOnSessionCreation() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CloseModelOnSessionCreation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ILearningModelSessionOptions2) Put_CloseModelOnSessionCreation(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CloseModelOnSessionCreation, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 58E15CEE-D8C2-56FC-92E8-76D751081086
var IID_ILearningModelSessionOptions3 = syscall.GUID{0x58E15CEE, 0xD8C2, 0x56FC,
	[8]byte{0x92, 0xE8, 0x76, 0xD7, 0x51, 0x08, 0x10, 0x86}}

type ILearningModelSessionOptions3Interface interface {
	win32.IInspectableInterface
	OverrideNamedDimension(name string, dimension uint32)
}

type ILearningModelSessionOptions3Vtbl struct {
	win32.IInspectableVtbl
	OverrideNamedDimension uintptr
}

type ILearningModelSessionOptions3 struct {
	win32.IInspectable
}

func (this *ILearningModelSessionOptions3) Vtbl() *ILearningModelSessionOptions3Vtbl {
	return (*ILearningModelSessionOptions3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILearningModelSessionOptions3) OverrideNamedDimension(name string, dimension uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OverrideNamedDimension, uintptr(unsafe.Pointer(this)), NewHStr(name).Ptr, uintptr(dimension))
	_ = _hr
}

// E3B977E8-6952-4E47-8EF4-1F7F07897C6D
var IID_ILearningModelStatics = syscall.GUID{0xE3B977E8, 0x6952, 0x4E47,
	[8]byte{0x8E, 0xF4, 0x1F, 0x7F, 0x07, 0x89, 0x7C, 0x6D}}

type ILearningModelStaticsInterface interface {
	win32.IInspectableInterface
	LoadFromStorageFileAsync(modelFile *IStorageFile) *IAsyncOperation[*ILearningModel]
	LoadFromStreamAsync(modelStream *IRandomAccessStreamReference) *IAsyncOperation[*ILearningModel]
	LoadFromFilePath(filePath string) *ILearningModel
	LoadFromStream(modelStream *IRandomAccessStreamReference) *ILearningModel
	LoadFromStorageFileWithOperatorProviderAsync(modelFile *IStorageFile, operatorProvider *ILearningModelOperatorProvider) *IAsyncOperation[*ILearningModel]
	LoadFromStreamWithOperatorProviderAsync(modelStream *IRandomAccessStreamReference, operatorProvider *ILearningModelOperatorProvider) *IAsyncOperation[*ILearningModel]
	LoadFromFilePathWithOperatorProvider(filePath string, operatorProvider *ILearningModelOperatorProvider) *ILearningModel
	LoadFromStreamWithOperatorProvider(modelStream *IRandomAccessStreamReference, operatorProvider *ILearningModelOperatorProvider) *ILearningModel
}

type ILearningModelStaticsVtbl struct {
	win32.IInspectableVtbl
	LoadFromStorageFileAsync                     uintptr
	LoadFromStreamAsync                          uintptr
	LoadFromFilePath                             uintptr
	LoadFromStream                               uintptr
	LoadFromStorageFileWithOperatorProviderAsync uintptr
	LoadFromStreamWithOperatorProviderAsync      uintptr
	LoadFromFilePathWithOperatorProvider         uintptr
	LoadFromStreamWithOperatorProvider           uintptr
}

type ILearningModelStatics struct {
	win32.IInspectable
}

func (this *ILearningModelStatics) Vtbl() *ILearningModelStaticsVtbl {
	return (*ILearningModelStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILearningModelStatics) LoadFromStorageFileAsync(modelFile *IStorageFile) *IAsyncOperation[*ILearningModel] {
	var _result *IAsyncOperation[*ILearningModel]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LoadFromStorageFileAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(modelFile)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILearningModelStatics) LoadFromStreamAsync(modelStream *IRandomAccessStreamReference) *IAsyncOperation[*ILearningModel] {
	var _result *IAsyncOperation[*ILearningModel]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LoadFromStreamAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(modelStream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILearningModelStatics) LoadFromFilePath(filePath string) *ILearningModel {
	var _result *ILearningModel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LoadFromFilePath, uintptr(unsafe.Pointer(this)), NewHStr(filePath).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILearningModelStatics) LoadFromStream(modelStream *IRandomAccessStreamReference) *ILearningModel {
	var _result *ILearningModel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LoadFromStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(modelStream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILearningModelStatics) LoadFromStorageFileWithOperatorProviderAsync(modelFile *IStorageFile, operatorProvider *ILearningModelOperatorProvider) *IAsyncOperation[*ILearningModel] {
	var _result *IAsyncOperation[*ILearningModel]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LoadFromStorageFileWithOperatorProviderAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(modelFile)), uintptr(unsafe.Pointer(operatorProvider)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILearningModelStatics) LoadFromStreamWithOperatorProviderAsync(modelStream *IRandomAccessStreamReference, operatorProvider *ILearningModelOperatorProvider) *IAsyncOperation[*ILearningModel] {
	var _result *IAsyncOperation[*ILearningModel]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LoadFromStreamWithOperatorProviderAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(modelStream)), uintptr(unsafe.Pointer(operatorProvider)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILearningModelStatics) LoadFromFilePathWithOperatorProvider(filePath string, operatorProvider *ILearningModelOperatorProvider) *ILearningModel {
	var _result *ILearningModel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LoadFromFilePathWithOperatorProvider, uintptr(unsafe.Pointer(this)), NewHStr(filePath).Ptr, uintptr(unsafe.Pointer(operatorProvider)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ILearningModelStatics) LoadFromStreamWithOperatorProvider(modelStream *IRandomAccessStreamReference, operatorProvider *ILearningModelOperatorProvider) *ILearningModel {
	var _result *ILearningModel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LoadFromStreamWithOperatorProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(modelStream)), uintptr(unsafe.Pointer(operatorProvider)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 530424BD-A257-436D-9E60-C2981F7CC5C4
var IID_IMapFeatureDescriptor = syscall.GUID{0x530424BD, 0xA257, 0x436D,
	[8]byte{0x9E, 0x60, 0xC2, 0x98, 0x1F, 0x7C, 0xC5, 0xC4}}

type IMapFeatureDescriptorInterface interface {
	win32.IInspectableInterface
	Get_KeyKind() TensorKind
	Get_ValueDescriptor() *ILearningModelFeatureDescriptor
}

type IMapFeatureDescriptorVtbl struct {
	win32.IInspectableVtbl
	Get_KeyKind         uintptr
	Get_ValueDescriptor uintptr
}

type IMapFeatureDescriptor struct {
	win32.IInspectable
}

func (this *IMapFeatureDescriptor) Vtbl() *IMapFeatureDescriptorVtbl {
	return (*IMapFeatureDescriptorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMapFeatureDescriptor) Get_KeyKind() TensorKind {
	var _result TensorKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_KeyKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMapFeatureDescriptor) Get_ValueDescriptor() *ILearningModelFeatureDescriptor {
	var _result *ILearningModelFeatureDescriptor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ValueDescriptor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 84F6945A-562B-4D62-A851-739ACED96668
var IID_ISequenceFeatureDescriptor = syscall.GUID{0x84F6945A, 0x562B, 0x4D62,
	[8]byte{0xA8, 0x51, 0x73, 0x9A, 0xCE, 0xD9, 0x66, 0x68}}

type ISequenceFeatureDescriptorInterface interface {
	win32.IInspectableInterface
	Get_ElementDescriptor() *ILearningModelFeatureDescriptor
}

type ISequenceFeatureDescriptorVtbl struct {
	win32.IInspectableVtbl
	Get_ElementDescriptor uintptr
}

type ISequenceFeatureDescriptor struct {
	win32.IInspectable
}

func (this *ISequenceFeatureDescriptor) Vtbl() *ISequenceFeatureDescriptorVtbl {
	return (*ISequenceFeatureDescriptorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISequenceFeatureDescriptor) Get_ElementDescriptor() *ILearningModelFeatureDescriptor {
	var _result *ILearningModelFeatureDescriptor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ElementDescriptor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 05489593-A305-4A25-AD09-440119B4B7F6
var IID_ITensor = syscall.GUID{0x05489593, 0xA305, 0x4A25,
	[8]byte{0xAD, 0x09, 0x44, 0x01, 0x19, 0xB4, 0xB7, 0xF6}}

type ITensorInterface interface {
	win32.IInspectableInterface
	Get_TensorKind() TensorKind
	Get_Shape() *IVectorView[int64]
}

type ITensorVtbl struct {
	win32.IInspectableVtbl
	Get_TensorKind uintptr
	Get_Shape      uintptr
}

type ITensor struct {
	win32.IInspectable
}

func (this *ITensor) Vtbl() *ITensorVtbl {
	return (*ITensorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensor) Get_TensorKind() TensorKind {
	var _result TensorKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TensorKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITensor) Get_Shape() *IVectorView[int64] {
	var _result *IVectorView[int64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Shape, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 50F311ED-29E9-4A5C-A44D-8FC512584EED
var IID_ITensorBoolean = syscall.GUID{0x50F311ED, 0x29E9, 0x4A5C,
	[8]byte{0xA4, 0x4D, 0x8F, 0xC5, 0x12, 0x58, 0x4E, 0xED}}

type ITensorBooleanInterface interface {
	win32.IInspectableInterface
	GetAsVectorView() *IVectorView[bool]
}

type ITensorBooleanVtbl struct {
	win32.IInspectableVtbl
	GetAsVectorView uintptr
}

type ITensorBoolean struct {
	win32.IInspectable
}

func (this *ITensorBoolean) Vtbl() *ITensorBooleanVtbl {
	return (*ITensorBooleanVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorBoolean) GetAsVectorView() *IVectorView[bool] {
	var _result *IVectorView[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAsVectorView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2796862C-2357-49A7-B476-D0AA3DFE6866
var IID_ITensorBooleanStatics = syscall.GUID{0x2796862C, 0x2357, 0x49A7,
	[8]byte{0xB4, 0x76, 0xD0, 0xAA, 0x3D, 0xFE, 0x68, 0x66}}

type ITensorBooleanStaticsInterface interface {
	win32.IInspectableInterface
	Create() *ITensorBoolean
	Create2(shape *IIterable[int64]) *ITensorBoolean
	CreateFromArray(shape *IIterable[int64], dataLength uint32, data *bool) *ITensorBoolean
	CreateFromIterable(shape *IIterable[int64], data *IIterable[bool]) *ITensorBoolean
}

type ITensorBooleanStaticsVtbl struct {
	win32.IInspectableVtbl
	Create             uintptr
	Create2            uintptr
	CreateFromArray    uintptr
	CreateFromIterable uintptr
}

type ITensorBooleanStatics struct {
	win32.IInspectable
}

func (this *ITensorBooleanStatics) Vtbl() *ITensorBooleanStaticsVtbl {
	return (*ITensorBooleanStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorBooleanStatics) Create() *ITensorBoolean {
	var _result *ITensorBoolean
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorBooleanStatics) Create2(shape *IIterable[int64]) *ITensorBoolean {
	var _result *ITensorBoolean
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorBooleanStatics) CreateFromArray(shape *IIterable[int64], dataLength uint32, data *bool) *ITensorBoolean {
	var _result *ITensorBoolean
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(dataLength), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorBooleanStatics) CreateFromIterable(shape *IIterable[int64], data *IIterable[bool]) *ITensorBoolean {
	var _result *ITensorBoolean
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromIterable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A3A4A501-6A2D-52D7-B04B-C435BAEE0115
var IID_ITensorBooleanStatics2 = syscall.GUID{0xA3A4A501, 0x6A2D, 0x52D7,
	[8]byte{0xB0, 0x4B, 0xC4, 0x35, 0xBA, 0xEE, 0x01, 0x15}}

type ITensorBooleanStatics2Interface interface {
	win32.IInspectableInterface
	CreateFromShapeArrayAndDataArray(shapeLength uint32, shape *int64, dataLength uint32, data *bool) *ITensorBoolean
	CreateFromBuffer(shapeLength uint32, shape *int64, buffer *IBuffer) *ITensorBoolean
}

type ITensorBooleanStatics2Vtbl struct {
	win32.IInspectableVtbl
	CreateFromShapeArrayAndDataArray uintptr
	CreateFromBuffer                 uintptr
}

type ITensorBooleanStatics2 struct {
	win32.IInspectable
}

func (this *ITensorBooleanStatics2) Vtbl() *ITensorBooleanStatics2Vtbl {
	return (*ITensorBooleanStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorBooleanStatics2) CreateFromShapeArrayAndDataArray(shapeLength uint32, shape *int64, dataLength uint32, data *bool) *ITensorBoolean {
	var _result *ITensorBoolean
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromShapeArrayAndDataArray, uintptr(unsafe.Pointer(this)), uintptr(shapeLength), uintptr(unsafe.Pointer(shape)), uintptr(dataLength), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorBooleanStatics2) CreateFromBuffer(shapeLength uint32, shape *int64, buffer *IBuffer) *ITensorBoolean {
	var _result *ITensorBoolean
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromBuffer, uintptr(unsafe.Pointer(this)), uintptr(shapeLength), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 91E41252-7A8F-4F0E-A28F-9637FFC8A3D0
var IID_ITensorDouble = syscall.GUID{0x91E41252, 0x7A8F, 0x4F0E,
	[8]byte{0xA2, 0x8F, 0x96, 0x37, 0xFF, 0xC8, 0xA3, 0xD0}}

type ITensorDoubleInterface interface {
	win32.IInspectableInterface
	GetAsVectorView() *IVectorView[float64]
}

type ITensorDoubleVtbl struct {
	win32.IInspectableVtbl
	GetAsVectorView uintptr
}

type ITensorDouble struct {
	win32.IInspectable
}

func (this *ITensorDouble) Vtbl() *ITensorDoubleVtbl {
	return (*ITensorDoubleVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorDouble) GetAsVectorView() *IVectorView[float64] {
	var _result *IVectorView[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAsVectorView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A86693C5-9538-44E7-A3CA-5DF374A5A70C
var IID_ITensorDoubleStatics = syscall.GUID{0xA86693C5, 0x9538, 0x44E7,
	[8]byte{0xA3, 0xCA, 0x5D, 0xF3, 0x74, 0xA5, 0xA7, 0x0C}}

type ITensorDoubleStaticsInterface interface {
	win32.IInspectableInterface
	Create() *ITensorDouble
	Create2(shape *IIterable[int64]) *ITensorDouble
	CreateFromArray(shape *IIterable[int64], dataLength uint32, data *float64) *ITensorDouble
	CreateFromIterable(shape *IIterable[int64], data *IIterable[float64]) *ITensorDouble
}

type ITensorDoubleStaticsVtbl struct {
	win32.IInspectableVtbl
	Create             uintptr
	Create2            uintptr
	CreateFromArray    uintptr
	CreateFromIterable uintptr
}

type ITensorDoubleStatics struct {
	win32.IInspectable
}

func (this *ITensorDoubleStatics) Vtbl() *ITensorDoubleStaticsVtbl {
	return (*ITensorDoubleStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorDoubleStatics) Create() *ITensorDouble {
	var _result *ITensorDouble
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorDoubleStatics) Create2(shape *IIterable[int64]) *ITensorDouble {
	var _result *ITensorDouble
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorDoubleStatics) CreateFromArray(shape *IIterable[int64], dataLength uint32, data *float64) *ITensorDouble {
	var _result *ITensorDouble
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(dataLength), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorDoubleStatics) CreateFromIterable(shape *IIterable[int64], data *IIterable[float64]) *ITensorDouble {
	var _result *ITensorDouble
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromIterable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 93A570DE-5E9A-5094-85C8-592C655E68AC
var IID_ITensorDoubleStatics2 = syscall.GUID{0x93A570DE, 0x5E9A, 0x5094,
	[8]byte{0x85, 0xC8, 0x59, 0x2C, 0x65, 0x5E, 0x68, 0xAC}}

type ITensorDoubleStatics2Interface interface {
	win32.IInspectableInterface
	CreateFromShapeArrayAndDataArray(shapeLength uint32, shape *int64, dataLength uint32, data *float64) *ITensorDouble
	CreateFromBuffer(shapeLength uint32, shape *int64, buffer *IBuffer) *ITensorDouble
}

type ITensorDoubleStatics2Vtbl struct {
	win32.IInspectableVtbl
	CreateFromShapeArrayAndDataArray uintptr
	CreateFromBuffer                 uintptr
}

type ITensorDoubleStatics2 struct {
	win32.IInspectable
}

func (this *ITensorDoubleStatics2) Vtbl() *ITensorDoubleStatics2Vtbl {
	return (*ITensorDoubleStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorDoubleStatics2) CreateFromShapeArrayAndDataArray(shapeLength uint32, shape *int64, dataLength uint32, data *float64) *ITensorDouble {
	var _result *ITensorDouble
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromShapeArrayAndDataArray, uintptr(unsafe.Pointer(this)), uintptr(shapeLength), uintptr(unsafe.Pointer(shape)), uintptr(dataLength), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorDoubleStatics2) CreateFromBuffer(shapeLength uint32, shape *int64, buffer *IBuffer) *ITensorDouble {
	var _result *ITensorDouble
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromBuffer, uintptr(unsafe.Pointer(this)), uintptr(shapeLength), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 74455C80-946A-4310-A19C-EE0AF028FCE4
var IID_ITensorFeatureDescriptor = syscall.GUID{0x74455C80, 0x946A, 0x4310,
	[8]byte{0xA1, 0x9C, 0xEE, 0x0A, 0xF0, 0x28, 0xFC, 0xE4}}

type ITensorFeatureDescriptorInterface interface {
	win32.IInspectableInterface
	Get_TensorKind() TensorKind
	Get_Shape() *IVectorView[int64]
}

type ITensorFeatureDescriptorVtbl struct {
	win32.IInspectableVtbl
	Get_TensorKind uintptr
	Get_Shape      uintptr
}

type ITensorFeatureDescriptor struct {
	win32.IInspectable
}

func (this *ITensorFeatureDescriptor) Vtbl() *ITensorFeatureDescriptorVtbl {
	return (*ITensorFeatureDescriptorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorFeatureDescriptor) Get_TensorKind() TensorKind {
	var _result TensorKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TensorKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ITensorFeatureDescriptor) Get_Shape() *IVectorView[int64] {
	var _result *IVectorView[int64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Shape, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F2282D82-AA02-42C8-A0C8-DF1EFC9676E1
var IID_ITensorFloat = syscall.GUID{0xF2282D82, 0xAA02, 0x42C8,
	[8]byte{0xA0, 0xC8, 0xDF, 0x1E, 0xFC, 0x96, 0x76, 0xE1}}

type ITensorFloatInterface interface {
	win32.IInspectableInterface
	GetAsVectorView() *IVectorView[float32]
}

type ITensorFloatVtbl struct {
	win32.IInspectableVtbl
	GetAsVectorView uintptr
}

type ITensorFloat struct {
	win32.IInspectable
}

func (this *ITensorFloat) Vtbl() *ITensorFloatVtbl {
	return (*ITensorFloatVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorFloat) GetAsVectorView() *IVectorView[float32] {
	var _result *IVectorView[float32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAsVectorView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0AB994FC-5B89-4C3C-B5E4-5282A5316C0A
var IID_ITensorFloat16Bit = syscall.GUID{0x0AB994FC, 0x5B89, 0x4C3C,
	[8]byte{0xB5, 0xE4, 0x52, 0x82, 0xA5, 0x31, 0x6C, 0x0A}}

type ITensorFloat16BitInterface interface {
	win32.IInspectableInterface
	GetAsVectorView() *IVectorView[float32]
}

type ITensorFloat16BitVtbl struct {
	win32.IInspectableVtbl
	GetAsVectorView uintptr
}

type ITensorFloat16Bit struct {
	win32.IInspectable
}

func (this *ITensorFloat16Bit) Vtbl() *ITensorFloat16BitVtbl {
	return (*ITensorFloat16BitVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorFloat16Bit) GetAsVectorView() *IVectorView[float32] {
	var _result *IVectorView[float32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAsVectorView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A52DB6F5-318A-44D4-820B-0CDC7054A84A
var IID_ITensorFloat16BitStatics = syscall.GUID{0xA52DB6F5, 0x318A, 0x44D4,
	[8]byte{0x82, 0x0B, 0x0C, 0xDC, 0x70, 0x54, 0xA8, 0x4A}}

type ITensorFloat16BitStaticsInterface interface {
	win32.IInspectableInterface
	Create() *ITensorFloat16Bit
	Create2(shape *IIterable[int64]) *ITensorFloat16Bit
	CreateFromArray(shape *IIterable[int64], dataLength uint32, data *float32) *ITensorFloat16Bit
	CreateFromIterable(shape *IIterable[int64], data *IIterable[float32]) *ITensorFloat16Bit
}

type ITensorFloat16BitStaticsVtbl struct {
	win32.IInspectableVtbl
	Create             uintptr
	Create2            uintptr
	CreateFromArray    uintptr
	CreateFromIterable uintptr
}

type ITensorFloat16BitStatics struct {
	win32.IInspectable
}

func (this *ITensorFloat16BitStatics) Vtbl() *ITensorFloat16BitStaticsVtbl {
	return (*ITensorFloat16BitStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorFloat16BitStatics) Create() *ITensorFloat16Bit {
	var _result *ITensorFloat16Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorFloat16BitStatics) Create2(shape *IIterable[int64]) *ITensorFloat16Bit {
	var _result *ITensorFloat16Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorFloat16BitStatics) CreateFromArray(shape *IIterable[int64], dataLength uint32, data *float32) *ITensorFloat16Bit {
	var _result *ITensorFloat16Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(dataLength), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorFloat16BitStatics) CreateFromIterable(shape *IIterable[int64], data *IIterable[float32]) *ITensorFloat16Bit {
	var _result *ITensorFloat16Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromIterable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 68545726-2DC7-51BF-B470-0B344CC2A1BC
var IID_ITensorFloat16BitStatics2 = syscall.GUID{0x68545726, 0x2DC7, 0x51BF,
	[8]byte{0xB4, 0x70, 0x0B, 0x34, 0x4C, 0xC2, 0xA1, 0xBC}}

type ITensorFloat16BitStatics2Interface interface {
	win32.IInspectableInterface
	CreateFromShapeArrayAndDataArray(shapeLength uint32, shape *int64, dataLength uint32, data *float32) *ITensorFloat16Bit
	CreateFromBuffer(shapeLength uint32, shape *int64, buffer *IBuffer) *ITensorFloat16Bit
}

type ITensorFloat16BitStatics2Vtbl struct {
	win32.IInspectableVtbl
	CreateFromShapeArrayAndDataArray uintptr
	CreateFromBuffer                 uintptr
}

type ITensorFloat16BitStatics2 struct {
	win32.IInspectable
}

func (this *ITensorFloat16BitStatics2) Vtbl() *ITensorFloat16BitStatics2Vtbl {
	return (*ITensorFloat16BitStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorFloat16BitStatics2) CreateFromShapeArrayAndDataArray(shapeLength uint32, shape *int64, dataLength uint32, data *float32) *ITensorFloat16Bit {
	var _result *ITensorFloat16Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromShapeArrayAndDataArray, uintptr(unsafe.Pointer(this)), uintptr(shapeLength), uintptr(unsafe.Pointer(shape)), uintptr(dataLength), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorFloat16BitStatics2) CreateFromBuffer(shapeLength uint32, shape *int64, buffer *IBuffer) *ITensorFloat16Bit {
	var _result *ITensorFloat16Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromBuffer, uintptr(unsafe.Pointer(this)), uintptr(shapeLength), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DBCD395B-3BA3-452F-B10D-3C135E573FA9
var IID_ITensorFloatStatics = syscall.GUID{0xDBCD395B, 0x3BA3, 0x452F,
	[8]byte{0xB1, 0x0D, 0x3C, 0x13, 0x5E, 0x57, 0x3F, 0xA9}}

type ITensorFloatStaticsInterface interface {
	win32.IInspectableInterface
	Create() *ITensorFloat
	Create2(shape *IIterable[int64]) *ITensorFloat
	CreateFromArray(shape *IIterable[int64], dataLength uint32, data *float32) *ITensorFloat
	CreateFromIterable(shape *IIterable[int64], data *IIterable[float32]) *ITensorFloat
}

type ITensorFloatStaticsVtbl struct {
	win32.IInspectableVtbl
	Create             uintptr
	Create2            uintptr
	CreateFromArray    uintptr
	CreateFromIterable uintptr
}

type ITensorFloatStatics struct {
	win32.IInspectable
}

func (this *ITensorFloatStatics) Vtbl() *ITensorFloatStaticsVtbl {
	return (*ITensorFloatStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorFloatStatics) Create() *ITensorFloat {
	var _result *ITensorFloat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorFloatStatics) Create2(shape *IIterable[int64]) *ITensorFloat {
	var _result *ITensorFloat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorFloatStatics) CreateFromArray(shape *IIterable[int64], dataLength uint32, data *float32) *ITensorFloat {
	var _result *ITensorFloat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(dataLength), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorFloatStatics) CreateFromIterable(shape *IIterable[int64], data *IIterable[float32]) *ITensorFloat {
	var _result *ITensorFloat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromIterable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 24610BC1-5E44-5713-B281-8F4AD4D555E8
var IID_ITensorFloatStatics2 = syscall.GUID{0x24610BC1, 0x5E44, 0x5713,
	[8]byte{0xB2, 0x81, 0x8F, 0x4A, 0xD4, 0xD5, 0x55, 0xE8}}

type ITensorFloatStatics2Interface interface {
	win32.IInspectableInterface
	CreateFromShapeArrayAndDataArray(shapeLength uint32, shape *int64, dataLength uint32, data *float32) *ITensorFloat
	CreateFromBuffer(shapeLength uint32, shape *int64, buffer *IBuffer) *ITensorFloat
}

type ITensorFloatStatics2Vtbl struct {
	win32.IInspectableVtbl
	CreateFromShapeArrayAndDataArray uintptr
	CreateFromBuffer                 uintptr
}

type ITensorFloatStatics2 struct {
	win32.IInspectable
}

func (this *ITensorFloatStatics2) Vtbl() *ITensorFloatStatics2Vtbl {
	return (*ITensorFloatStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorFloatStatics2) CreateFromShapeArrayAndDataArray(shapeLength uint32, shape *int64, dataLength uint32, data *float32) *ITensorFloat {
	var _result *ITensorFloat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromShapeArrayAndDataArray, uintptr(unsafe.Pointer(this)), uintptr(shapeLength), uintptr(unsafe.Pointer(shape)), uintptr(dataLength), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorFloatStatics2) CreateFromBuffer(shapeLength uint32, shape *int64, buffer *IBuffer) *ITensorFloat {
	var _result *ITensorFloat
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromBuffer, uintptr(unsafe.Pointer(this)), uintptr(shapeLength), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 98A32D39-E6D6-44AF-8AFA-BAEBC44DC020
var IID_ITensorInt16Bit = syscall.GUID{0x98A32D39, 0xE6D6, 0x44AF,
	[8]byte{0x8A, 0xFA, 0xBA, 0xEB, 0xC4, 0x4D, 0xC0, 0x20}}

type ITensorInt16BitInterface interface {
	win32.IInspectableInterface
	GetAsVectorView() *IVectorView[int16]
}

type ITensorInt16BitVtbl struct {
	win32.IInspectableVtbl
	GetAsVectorView uintptr
}

type ITensorInt16Bit struct {
	win32.IInspectable
}

func (this *ITensorInt16Bit) Vtbl() *ITensorInt16BitVtbl {
	return (*ITensorInt16BitVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorInt16Bit) GetAsVectorView() *IVectorView[int16] {
	var _result *IVectorView[int16]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAsVectorView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 98646293-266E-4B1A-821F-E60D70898B91
var IID_ITensorInt16BitStatics = syscall.GUID{0x98646293, 0x266E, 0x4B1A,
	[8]byte{0x82, 0x1F, 0xE6, 0x0D, 0x70, 0x89, 0x8B, 0x91}}

type ITensorInt16BitStaticsInterface interface {
	win32.IInspectableInterface
	Create() *ITensorInt16Bit
	Create2(shape *IIterable[int64]) *ITensorInt16Bit
	CreateFromArray(shape *IIterable[int64], dataLength uint32, data *int16) *ITensorInt16Bit
	CreateFromIterable(shape *IIterable[int64], data *IIterable[int16]) *ITensorInt16Bit
}

type ITensorInt16BitStaticsVtbl struct {
	win32.IInspectableVtbl
	Create             uintptr
	Create2            uintptr
	CreateFromArray    uintptr
	CreateFromIterable uintptr
}

type ITensorInt16BitStatics struct {
	win32.IInspectable
}

func (this *ITensorInt16BitStatics) Vtbl() *ITensorInt16BitStaticsVtbl {
	return (*ITensorInt16BitStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorInt16BitStatics) Create() *ITensorInt16Bit {
	var _result *ITensorInt16Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorInt16BitStatics) Create2(shape *IIterable[int64]) *ITensorInt16Bit {
	var _result *ITensorInt16Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorInt16BitStatics) CreateFromArray(shape *IIterable[int64], dataLength uint32, data *int16) *ITensorInt16Bit {
	var _result *ITensorInt16Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(dataLength), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorInt16BitStatics) CreateFromIterable(shape *IIterable[int64], data *IIterable[int16]) *ITensorInt16Bit {
	var _result *ITensorInt16Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromIterable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0CD70CF4-696C-5E5F-95D8-5EBF9670148B
var IID_ITensorInt16BitStatics2 = syscall.GUID{0x0CD70CF4, 0x696C, 0x5E5F,
	[8]byte{0x95, 0xD8, 0x5E, 0xBF, 0x96, 0x70, 0x14, 0x8B}}

type ITensorInt16BitStatics2Interface interface {
	win32.IInspectableInterface
	CreateFromShapeArrayAndDataArray(shapeLength uint32, shape *int64, dataLength uint32, data *int16) *ITensorInt16Bit
	CreateFromBuffer(shapeLength uint32, shape *int64, buffer *IBuffer) *ITensorInt16Bit
}

type ITensorInt16BitStatics2Vtbl struct {
	win32.IInspectableVtbl
	CreateFromShapeArrayAndDataArray uintptr
	CreateFromBuffer                 uintptr
}

type ITensorInt16BitStatics2 struct {
	win32.IInspectable
}

func (this *ITensorInt16BitStatics2) Vtbl() *ITensorInt16BitStatics2Vtbl {
	return (*ITensorInt16BitStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorInt16BitStatics2) CreateFromShapeArrayAndDataArray(shapeLength uint32, shape *int64, dataLength uint32, data *int16) *ITensorInt16Bit {
	var _result *ITensorInt16Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromShapeArrayAndDataArray, uintptr(unsafe.Pointer(this)), uintptr(shapeLength), uintptr(unsafe.Pointer(shape)), uintptr(dataLength), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorInt16BitStatics2) CreateFromBuffer(shapeLength uint32, shape *int64, buffer *IBuffer) *ITensorInt16Bit {
	var _result *ITensorInt16Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromBuffer, uintptr(unsafe.Pointer(this)), uintptr(shapeLength), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2C0C28D3-207C-4486-A7D2-884522C5E589
var IID_ITensorInt32Bit = syscall.GUID{0x2C0C28D3, 0x207C, 0x4486,
	[8]byte{0xA7, 0xD2, 0x88, 0x45, 0x22, 0xC5, 0xE5, 0x89}}

type ITensorInt32BitInterface interface {
	win32.IInspectableInterface
	GetAsVectorView() *IVectorView[int32]
}

type ITensorInt32BitVtbl struct {
	win32.IInspectableVtbl
	GetAsVectorView uintptr
}

type ITensorInt32Bit struct {
	win32.IInspectable
}

func (this *ITensorInt32Bit) Vtbl() *ITensorInt32BitVtbl {
	return (*ITensorInt32BitVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorInt32Bit) GetAsVectorView() *IVectorView[int32] {
	var _result *IVectorView[int32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAsVectorView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6539864B-52FA-4E35-907C-834CAC417B50
var IID_ITensorInt32BitStatics = syscall.GUID{0x6539864B, 0x52FA, 0x4E35,
	[8]byte{0x90, 0x7C, 0x83, 0x4C, 0xAC, 0x41, 0x7B, 0x50}}

type ITensorInt32BitStaticsInterface interface {
	win32.IInspectableInterface
	Create() *ITensorInt32Bit
	Create2(shape *IIterable[int64]) *ITensorInt32Bit
	CreateFromArray(shape *IIterable[int64], dataLength uint32, data *int32) *ITensorInt32Bit
	CreateFromIterable(shape *IIterable[int64], data *IIterable[int32]) *ITensorInt32Bit
}

type ITensorInt32BitStaticsVtbl struct {
	win32.IInspectableVtbl
	Create             uintptr
	Create2            uintptr
	CreateFromArray    uintptr
	CreateFromIterable uintptr
}

type ITensorInt32BitStatics struct {
	win32.IInspectable
}

func (this *ITensorInt32BitStatics) Vtbl() *ITensorInt32BitStaticsVtbl {
	return (*ITensorInt32BitStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorInt32BitStatics) Create() *ITensorInt32Bit {
	var _result *ITensorInt32Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorInt32BitStatics) Create2(shape *IIterable[int64]) *ITensorInt32Bit {
	var _result *ITensorInt32Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorInt32BitStatics) CreateFromArray(shape *IIterable[int64], dataLength uint32, data *int32) *ITensorInt32Bit {
	var _result *ITensorInt32Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(dataLength), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorInt32BitStatics) CreateFromIterable(shape *IIterable[int64], data *IIterable[int32]) *ITensorInt32Bit {
	var _result *ITensorInt32Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromIterable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7C4B079A-E956-5CE0-A3BD-157D9D79B5EC
var IID_ITensorInt32BitStatics2 = syscall.GUID{0x7C4B079A, 0xE956, 0x5CE0,
	[8]byte{0xA3, 0xBD, 0x15, 0x7D, 0x9D, 0x79, 0xB5, 0xEC}}

type ITensorInt32BitStatics2Interface interface {
	win32.IInspectableInterface
	CreateFromShapeArrayAndDataArray(shapeLength uint32, shape *int64, dataLength uint32, data *int32) *ITensorInt32Bit
	CreateFromBuffer(shapeLength uint32, shape *int64, buffer *IBuffer) *ITensorInt32Bit
}

type ITensorInt32BitStatics2Vtbl struct {
	win32.IInspectableVtbl
	CreateFromShapeArrayAndDataArray uintptr
	CreateFromBuffer                 uintptr
}

type ITensorInt32BitStatics2 struct {
	win32.IInspectable
}

func (this *ITensorInt32BitStatics2) Vtbl() *ITensorInt32BitStatics2Vtbl {
	return (*ITensorInt32BitStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorInt32BitStatics2) CreateFromShapeArrayAndDataArray(shapeLength uint32, shape *int64, dataLength uint32, data *int32) *ITensorInt32Bit {
	var _result *ITensorInt32Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromShapeArrayAndDataArray, uintptr(unsafe.Pointer(this)), uintptr(shapeLength), uintptr(unsafe.Pointer(shape)), uintptr(dataLength), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorInt32BitStatics2) CreateFromBuffer(shapeLength uint32, shape *int64, buffer *IBuffer) *ITensorInt32Bit {
	var _result *ITensorInt32Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromBuffer, uintptr(unsafe.Pointer(this)), uintptr(shapeLength), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 499665BA-1FA2-45AD-AF25-A0BD9BDA4C87
var IID_ITensorInt64Bit = syscall.GUID{0x499665BA, 0x1FA2, 0x45AD,
	[8]byte{0xAF, 0x25, 0xA0, 0xBD, 0x9B, 0xDA, 0x4C, 0x87}}

type ITensorInt64BitInterface interface {
	win32.IInspectableInterface
	GetAsVectorView() *IVectorView[int64]
}

type ITensorInt64BitVtbl struct {
	win32.IInspectableVtbl
	GetAsVectorView uintptr
}

type ITensorInt64Bit struct {
	win32.IInspectable
}

func (this *ITensorInt64Bit) Vtbl() *ITensorInt64BitVtbl {
	return (*ITensorInt64BitVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorInt64Bit) GetAsVectorView() *IVectorView[int64] {
	var _result *IVectorView[int64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAsVectorView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9648AD9D-1198-4D74-9517-783AB62B9CC2
var IID_ITensorInt64BitStatics = syscall.GUID{0x9648AD9D, 0x1198, 0x4D74,
	[8]byte{0x95, 0x17, 0x78, 0x3A, 0xB6, 0x2B, 0x9C, 0xC2}}

type ITensorInt64BitStaticsInterface interface {
	win32.IInspectableInterface
	Create() *ITensorInt64Bit
	Create2(shape *IIterable[int64]) *ITensorInt64Bit
	CreateFromArray(shape *IIterable[int64], dataLength uint32, data *int64) *ITensorInt64Bit
	CreateFromIterable(shape *IIterable[int64], data *IIterable[int64]) *ITensorInt64Bit
}

type ITensorInt64BitStaticsVtbl struct {
	win32.IInspectableVtbl
	Create             uintptr
	Create2            uintptr
	CreateFromArray    uintptr
	CreateFromIterable uintptr
}

type ITensorInt64BitStatics struct {
	win32.IInspectable
}

func (this *ITensorInt64BitStatics) Vtbl() *ITensorInt64BitStaticsVtbl {
	return (*ITensorInt64BitStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorInt64BitStatics) Create() *ITensorInt64Bit {
	var _result *ITensorInt64Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorInt64BitStatics) Create2(shape *IIterable[int64]) *ITensorInt64Bit {
	var _result *ITensorInt64Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorInt64BitStatics) CreateFromArray(shape *IIterable[int64], dataLength uint32, data *int64) *ITensorInt64Bit {
	var _result *ITensorInt64Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(dataLength), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorInt64BitStatics) CreateFromIterable(shape *IIterable[int64], data *IIterable[int64]) *ITensorInt64Bit {
	var _result *ITensorInt64Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromIterable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6D3D9DCB-FF40-5EC2-89FE-084E2B6BC6DB
var IID_ITensorInt64BitStatics2 = syscall.GUID{0x6D3D9DCB, 0xFF40, 0x5EC2,
	[8]byte{0x89, 0xFE, 0x08, 0x4E, 0x2B, 0x6B, 0xC6, 0xDB}}

type ITensorInt64BitStatics2Interface interface {
	win32.IInspectableInterface
	CreateFromShapeArrayAndDataArray(shapeLength uint32, shape *int64, dataLength uint32, data *int64) *ITensorInt64Bit
	CreateFromBuffer(shapeLength uint32, shape *int64, buffer *IBuffer) *ITensorInt64Bit
}

type ITensorInt64BitStatics2Vtbl struct {
	win32.IInspectableVtbl
	CreateFromShapeArrayAndDataArray uintptr
	CreateFromBuffer                 uintptr
}

type ITensorInt64BitStatics2 struct {
	win32.IInspectable
}

func (this *ITensorInt64BitStatics2) Vtbl() *ITensorInt64BitStatics2Vtbl {
	return (*ITensorInt64BitStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorInt64BitStatics2) CreateFromShapeArrayAndDataArray(shapeLength uint32, shape *int64, dataLength uint32, data *int64) *ITensorInt64Bit {
	var _result *ITensorInt64Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromShapeArrayAndDataArray, uintptr(unsafe.Pointer(this)), uintptr(shapeLength), uintptr(unsafe.Pointer(shape)), uintptr(dataLength), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorInt64BitStatics2) CreateFromBuffer(shapeLength uint32, shape *int64, buffer *IBuffer) *ITensorInt64Bit {
	var _result *ITensorInt64Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromBuffer, uintptr(unsafe.Pointer(this)), uintptr(shapeLength), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CDDD97C5-FFD8-4FEF-AEFB-30E1A485B2EE
var IID_ITensorInt8Bit = syscall.GUID{0xCDDD97C5, 0xFFD8, 0x4FEF,
	[8]byte{0xAE, 0xFB, 0x30, 0xE1, 0xA4, 0x85, 0xB2, 0xEE}}

type ITensorInt8BitInterface interface {
	win32.IInspectableInterface
	GetAsVectorView() *IVectorView[byte]
}

type ITensorInt8BitVtbl struct {
	win32.IInspectableVtbl
	GetAsVectorView uintptr
}

type ITensorInt8Bit struct {
	win32.IInspectable
}

func (this *ITensorInt8Bit) Vtbl() *ITensorInt8BitVtbl {
	return (*ITensorInt8BitVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorInt8Bit) GetAsVectorView() *IVectorView[byte] {
	var _result *IVectorView[byte]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAsVectorView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B1A12284-095C-4C76-A661-AC4CEE1F3E8B
var IID_ITensorInt8BitStatics = syscall.GUID{0xB1A12284, 0x095C, 0x4C76,
	[8]byte{0xA6, 0x61, 0xAC, 0x4C, 0xEE, 0x1F, 0x3E, 0x8B}}

type ITensorInt8BitStaticsInterface interface {
	win32.IInspectableInterface
	Create() *ITensorInt8Bit
	Create2(shape *IIterable[int64]) *ITensorInt8Bit
	CreateFromArray(shape *IIterable[int64], dataLength uint32, data *byte) *ITensorInt8Bit
	CreateFromIterable(shape *IIterable[int64], data *IIterable[byte]) *ITensorInt8Bit
}

type ITensorInt8BitStaticsVtbl struct {
	win32.IInspectableVtbl
	Create             uintptr
	Create2            uintptr
	CreateFromArray    uintptr
	CreateFromIterable uintptr
}

type ITensorInt8BitStatics struct {
	win32.IInspectable
}

func (this *ITensorInt8BitStatics) Vtbl() *ITensorInt8BitStaticsVtbl {
	return (*ITensorInt8BitStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorInt8BitStatics) Create() *ITensorInt8Bit {
	var _result *ITensorInt8Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorInt8BitStatics) Create2(shape *IIterable[int64]) *ITensorInt8Bit {
	var _result *ITensorInt8Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorInt8BitStatics) CreateFromArray(shape *IIterable[int64], dataLength uint32, data *byte) *ITensorInt8Bit {
	var _result *ITensorInt8Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(dataLength), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorInt8BitStatics) CreateFromIterable(shape *IIterable[int64], data *IIterable[byte]) *ITensorInt8Bit {
	var _result *ITensorInt8Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromIterable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C0D59637-C468-56FB-9535-C052BDB93DC0
var IID_ITensorInt8BitStatics2 = syscall.GUID{0xC0D59637, 0xC468, 0x56FB,
	[8]byte{0x95, 0x35, 0xC0, 0x52, 0xBD, 0xB9, 0x3D, 0xC0}}

type ITensorInt8BitStatics2Interface interface {
	win32.IInspectableInterface
	CreateFromShapeArrayAndDataArray(shapeLength uint32, shape *int64, dataLength uint32, data *byte) *ITensorInt8Bit
	CreateFromBuffer(shapeLength uint32, shape *int64, buffer *IBuffer) *ITensorInt8Bit
}

type ITensorInt8BitStatics2Vtbl struct {
	win32.IInspectableVtbl
	CreateFromShapeArrayAndDataArray uintptr
	CreateFromBuffer                 uintptr
}

type ITensorInt8BitStatics2 struct {
	win32.IInspectable
}

func (this *ITensorInt8BitStatics2) Vtbl() *ITensorInt8BitStatics2Vtbl {
	return (*ITensorInt8BitStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorInt8BitStatics2) CreateFromShapeArrayAndDataArray(shapeLength uint32, shape *int64, dataLength uint32, data *byte) *ITensorInt8Bit {
	var _result *ITensorInt8Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromShapeArrayAndDataArray, uintptr(unsafe.Pointer(this)), uintptr(shapeLength), uintptr(unsafe.Pointer(shape)), uintptr(dataLength), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorInt8BitStatics2) CreateFromBuffer(shapeLength uint32, shape *int64, buffer *IBuffer) *ITensorInt8Bit {
	var _result *ITensorInt8Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromBuffer, uintptr(unsafe.Pointer(this)), uintptr(shapeLength), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 582335C8-BDB1-4610-BC75-35E9CBF009B7
var IID_ITensorString = syscall.GUID{0x582335C8, 0xBDB1, 0x4610,
	[8]byte{0xBC, 0x75, 0x35, 0xE9, 0xCB, 0xF0, 0x09, 0xB7}}

type ITensorStringInterface interface {
	win32.IInspectableInterface
	GetAsVectorView() *IVectorView[string]
}

type ITensorStringVtbl struct {
	win32.IInspectableVtbl
	GetAsVectorView uintptr
}

type ITensorString struct {
	win32.IInspectable
}

func (this *ITensorString) Vtbl() *ITensorStringVtbl {
	return (*ITensorStringVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorString) GetAsVectorView() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAsVectorView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 83623324-CF26-4F17-A2D4-20EF8D097D53
var IID_ITensorStringStatics = syscall.GUID{0x83623324, 0xCF26, 0x4F17,
	[8]byte{0xA2, 0xD4, 0x20, 0xEF, 0x8D, 0x09, 0x7D, 0x53}}

type ITensorStringStaticsInterface interface {
	win32.IInspectableInterface
	Create() *ITensorString
	Create2(shape *IIterable[int64]) *ITensorString
	CreateFromArray(shape *IIterable[int64], dataLength uint32, data *string) *ITensorString
	CreateFromIterable(shape *IIterable[int64], data *IIterable[string]) *ITensorString
}

type ITensorStringStaticsVtbl struct {
	win32.IInspectableVtbl
	Create             uintptr
	Create2            uintptr
	CreateFromArray    uintptr
	CreateFromIterable uintptr
}

type ITensorStringStatics struct {
	win32.IInspectable
}

func (this *ITensorStringStatics) Vtbl() *ITensorStringStaticsVtbl {
	return (*ITensorStringStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorStringStatics) Create() *ITensorString {
	var _result *ITensorString
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorStringStatics) Create2(shape *IIterable[int64]) *ITensorString {
	var _result *ITensorString
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorStringStatics) CreateFromArray(shape *IIterable[int64], dataLength uint32, data *string) *ITensorString {
	var _result *ITensorString
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(dataLength), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorStringStatics) CreateFromIterable(shape *IIterable[int64], data *IIterable[string]) *ITensorString {
	var _result *ITensorString
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromIterable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9E355ED0-C8E2-5254-9137-0193A3668FD8
var IID_ITensorStringStatics2 = syscall.GUID{0x9E355ED0, 0xC8E2, 0x5254,
	[8]byte{0x91, 0x37, 0x01, 0x93, 0xA3, 0x66, 0x8F, 0xD8}}

type ITensorStringStatics2Interface interface {
	win32.IInspectableInterface
	CreateFromShapeArrayAndDataArray(shapeLength uint32, shape *int64, dataLength uint32, data *string) *ITensorString
}

type ITensorStringStatics2Vtbl struct {
	win32.IInspectableVtbl
	CreateFromShapeArrayAndDataArray uintptr
}

type ITensorStringStatics2 struct {
	win32.IInspectable
}

func (this *ITensorStringStatics2) Vtbl() *ITensorStringStatics2Vtbl {
	return (*ITensorStringStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorStringStatics2) CreateFromShapeArrayAndDataArray(shapeLength uint32, shape *int64, dataLength uint32, data *string) *ITensorString {
	var _result *ITensorString
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromShapeArrayAndDataArray, uintptr(unsafe.Pointer(this)), uintptr(shapeLength), uintptr(unsafe.Pointer(shape)), uintptr(dataLength), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 68140F4B-23C0-42F3-81F6-A891C011BC3F
var IID_ITensorUInt16Bit = syscall.GUID{0x68140F4B, 0x23C0, 0x42F3,
	[8]byte{0x81, 0xF6, 0xA8, 0x91, 0xC0, 0x11, 0xBC, 0x3F}}

type ITensorUInt16BitInterface interface {
	win32.IInspectableInterface
	GetAsVectorView() *IVectorView[uint16]
}

type ITensorUInt16BitVtbl struct {
	win32.IInspectableVtbl
	GetAsVectorView uintptr
}

type ITensorUInt16Bit struct {
	win32.IInspectable
}

func (this *ITensorUInt16Bit) Vtbl() *ITensorUInt16BitVtbl {
	return (*ITensorUInt16BitVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorUInt16Bit) GetAsVectorView() *IVectorView[uint16] {
	var _result *IVectorView[uint16]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAsVectorView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5DF745DD-028A-481A-A27C-C7E6435E52DD
var IID_ITensorUInt16BitStatics = syscall.GUID{0x5DF745DD, 0x028A, 0x481A,
	[8]byte{0xA2, 0x7C, 0xC7, 0xE6, 0x43, 0x5E, 0x52, 0xDD}}

type ITensorUInt16BitStaticsInterface interface {
	win32.IInspectableInterface
	Create() *ITensorUInt16Bit
	Create2(shape *IIterable[int64]) *ITensorUInt16Bit
	CreateFromArray(shape *IIterable[int64], dataLength uint32, data *uint16) *ITensorUInt16Bit
	CreateFromIterable(shape *IIterable[int64], data *IIterable[uint16]) *ITensorUInt16Bit
}

type ITensorUInt16BitStaticsVtbl struct {
	win32.IInspectableVtbl
	Create             uintptr
	Create2            uintptr
	CreateFromArray    uintptr
	CreateFromIterable uintptr
}

type ITensorUInt16BitStatics struct {
	win32.IInspectable
}

func (this *ITensorUInt16BitStatics) Vtbl() *ITensorUInt16BitStaticsVtbl {
	return (*ITensorUInt16BitStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorUInt16BitStatics) Create() *ITensorUInt16Bit {
	var _result *ITensorUInt16Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorUInt16BitStatics) Create2(shape *IIterable[int64]) *ITensorUInt16Bit {
	var _result *ITensorUInt16Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorUInt16BitStatics) CreateFromArray(shape *IIterable[int64], dataLength uint32, data *uint16) *ITensorUInt16Bit {
	var _result *ITensorUInt16Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(dataLength), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorUInt16BitStatics) CreateFromIterable(shape *IIterable[int64], data *IIterable[uint16]) *ITensorUInt16Bit {
	var _result *ITensorUInt16Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromIterable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8AF40C64-D69F-5315-9348-490877BBD642
var IID_ITensorUInt16BitStatics2 = syscall.GUID{0x8AF40C64, 0xD69F, 0x5315,
	[8]byte{0x93, 0x48, 0x49, 0x08, 0x77, 0xBB, 0xD6, 0x42}}

type ITensorUInt16BitStatics2Interface interface {
	win32.IInspectableInterface
	CreateFromShapeArrayAndDataArray(shapeLength uint32, shape *int64, dataLength uint32, data *uint16) *ITensorUInt16Bit
	CreateFromBuffer(shapeLength uint32, shape *int64, buffer *IBuffer) *ITensorUInt16Bit
}

type ITensorUInt16BitStatics2Vtbl struct {
	win32.IInspectableVtbl
	CreateFromShapeArrayAndDataArray uintptr
	CreateFromBuffer                 uintptr
}

type ITensorUInt16BitStatics2 struct {
	win32.IInspectable
}

func (this *ITensorUInt16BitStatics2) Vtbl() *ITensorUInt16BitStatics2Vtbl {
	return (*ITensorUInt16BitStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorUInt16BitStatics2) CreateFromShapeArrayAndDataArray(shapeLength uint32, shape *int64, dataLength uint32, data *uint16) *ITensorUInt16Bit {
	var _result *ITensorUInt16Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromShapeArrayAndDataArray, uintptr(unsafe.Pointer(this)), uintptr(shapeLength), uintptr(unsafe.Pointer(shape)), uintptr(dataLength), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorUInt16BitStatics2) CreateFromBuffer(shapeLength uint32, shape *int64, buffer *IBuffer) *ITensorUInt16Bit {
	var _result *ITensorUInt16Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromBuffer, uintptr(unsafe.Pointer(this)), uintptr(shapeLength), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D8C9C2FF-7511-45A3-BFAC-C38F370D2237
var IID_ITensorUInt32Bit = syscall.GUID{0xD8C9C2FF, 0x7511, 0x45A3,
	[8]byte{0xBF, 0xAC, 0xC3, 0x8F, 0x37, 0x0D, 0x22, 0x37}}

type ITensorUInt32BitInterface interface {
	win32.IInspectableInterface
	GetAsVectorView() *IVectorView[uint32]
}

type ITensorUInt32BitVtbl struct {
	win32.IInspectableVtbl
	GetAsVectorView uintptr
}

type ITensorUInt32Bit struct {
	win32.IInspectable
}

func (this *ITensorUInt32Bit) Vtbl() *ITensorUInt32BitVtbl {
	return (*ITensorUInt32BitVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorUInt32Bit) GetAsVectorView() *IVectorView[uint32] {
	var _result *IVectorView[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAsVectorView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 417C3837-E773-4378-8E7F-0CC33DBEA697
var IID_ITensorUInt32BitStatics = syscall.GUID{0x417C3837, 0xE773, 0x4378,
	[8]byte{0x8E, 0x7F, 0x0C, 0xC3, 0x3D, 0xBE, 0xA6, 0x97}}

type ITensorUInt32BitStaticsInterface interface {
	win32.IInspectableInterface
	Create() *ITensorUInt32Bit
	Create2(shape *IIterable[int64]) *ITensorUInt32Bit
	CreateFromArray(shape *IIterable[int64], dataLength uint32, data *uint32) *ITensorUInt32Bit
	CreateFromIterable(shape *IIterable[int64], data *IIterable[uint32]) *ITensorUInt32Bit
}

type ITensorUInt32BitStaticsVtbl struct {
	win32.IInspectableVtbl
	Create             uintptr
	Create2            uintptr
	CreateFromArray    uintptr
	CreateFromIterable uintptr
}

type ITensorUInt32BitStatics struct {
	win32.IInspectable
}

func (this *ITensorUInt32BitStatics) Vtbl() *ITensorUInt32BitStaticsVtbl {
	return (*ITensorUInt32BitStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorUInt32BitStatics) Create() *ITensorUInt32Bit {
	var _result *ITensorUInt32Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorUInt32BitStatics) Create2(shape *IIterable[int64]) *ITensorUInt32Bit {
	var _result *ITensorUInt32Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorUInt32BitStatics) CreateFromArray(shape *IIterable[int64], dataLength uint32, data *uint32) *ITensorUInt32Bit {
	var _result *ITensorUInt32Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(dataLength), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorUInt32BitStatics) CreateFromIterable(shape *IIterable[int64], data *IIterable[uint32]) *ITensorUInt32Bit {
	var _result *ITensorUInt32Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromIterable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EF1A1F1C-314E-569D-B496-5C8447D20CD2
var IID_ITensorUInt32BitStatics2 = syscall.GUID{0xEF1A1F1C, 0x314E, 0x569D,
	[8]byte{0xB4, 0x96, 0x5C, 0x84, 0x47, 0xD2, 0x0C, 0xD2}}

type ITensorUInt32BitStatics2Interface interface {
	win32.IInspectableInterface
	CreateFromShapeArrayAndDataArray(shapeLength uint32, shape *int64, dataLength uint32, data *uint32) *ITensorUInt32Bit
	CreateFromBuffer(shapeLength uint32, shape *int64, buffer *IBuffer) *ITensorUInt32Bit
}

type ITensorUInt32BitStatics2Vtbl struct {
	win32.IInspectableVtbl
	CreateFromShapeArrayAndDataArray uintptr
	CreateFromBuffer                 uintptr
}

type ITensorUInt32BitStatics2 struct {
	win32.IInspectable
}

func (this *ITensorUInt32BitStatics2) Vtbl() *ITensorUInt32BitStatics2Vtbl {
	return (*ITensorUInt32BitStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorUInt32BitStatics2) CreateFromShapeArrayAndDataArray(shapeLength uint32, shape *int64, dataLength uint32, data *uint32) *ITensorUInt32Bit {
	var _result *ITensorUInt32Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromShapeArrayAndDataArray, uintptr(unsafe.Pointer(this)), uintptr(shapeLength), uintptr(unsafe.Pointer(shape)), uintptr(dataLength), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorUInt32BitStatics2) CreateFromBuffer(shapeLength uint32, shape *int64, buffer *IBuffer) *ITensorUInt32Bit {
	var _result *ITensorUInt32Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromBuffer, uintptr(unsafe.Pointer(this)), uintptr(shapeLength), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2E70FFAD-04BF-4825-839A-82BAEF8C7886
var IID_ITensorUInt64Bit = syscall.GUID{0x2E70FFAD, 0x04BF, 0x4825,
	[8]byte{0x83, 0x9A, 0x82, 0xBA, 0xEF, 0x8C, 0x78, 0x86}}

type ITensorUInt64BitInterface interface {
	win32.IInspectableInterface
	GetAsVectorView() *IVectorView[uint64]
}

type ITensorUInt64BitVtbl struct {
	win32.IInspectableVtbl
	GetAsVectorView uintptr
}

type ITensorUInt64Bit struct {
	win32.IInspectable
}

func (this *ITensorUInt64Bit) Vtbl() *ITensorUInt64BitVtbl {
	return (*ITensorUInt64BitVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorUInt64Bit) GetAsVectorView() *IVectorView[uint64] {
	var _result *IVectorView[uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAsVectorView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7A7E20EB-242F-47CB-A9C6-F602ECFBFEE4
var IID_ITensorUInt64BitStatics = syscall.GUID{0x7A7E20EB, 0x242F, 0x47CB,
	[8]byte{0xA9, 0xC6, 0xF6, 0x02, 0xEC, 0xFB, 0xFE, 0xE4}}

type ITensorUInt64BitStaticsInterface interface {
	win32.IInspectableInterface
	Create() *ITensorUInt64Bit
	Create2(shape *IIterable[int64]) *ITensorUInt64Bit
	CreateFromArray(shape *IIterable[int64], dataLength uint32, data *uint64) *ITensorUInt64Bit
	CreateFromIterable(shape *IIterable[int64], data *IIterable[uint64]) *ITensorUInt64Bit
}

type ITensorUInt64BitStaticsVtbl struct {
	win32.IInspectableVtbl
	Create             uintptr
	Create2            uintptr
	CreateFromArray    uintptr
	CreateFromIterable uintptr
}

type ITensorUInt64BitStatics struct {
	win32.IInspectable
}

func (this *ITensorUInt64BitStatics) Vtbl() *ITensorUInt64BitStaticsVtbl {
	return (*ITensorUInt64BitStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorUInt64BitStatics) Create() *ITensorUInt64Bit {
	var _result *ITensorUInt64Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorUInt64BitStatics) Create2(shape *IIterable[int64]) *ITensorUInt64Bit {
	var _result *ITensorUInt64Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorUInt64BitStatics) CreateFromArray(shape *IIterable[int64], dataLength uint32, data *uint64) *ITensorUInt64Bit {
	var _result *ITensorUInt64Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(dataLength), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorUInt64BitStatics) CreateFromIterable(shape *IIterable[int64], data *IIterable[uint64]) *ITensorUInt64Bit {
	var _result *ITensorUInt64Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromIterable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 085A687D-67E1-5B1E-B232-4FABE9CA20B3
var IID_ITensorUInt64BitStatics2 = syscall.GUID{0x085A687D, 0x67E1, 0x5B1E,
	[8]byte{0xB2, 0x32, 0x4F, 0xAB, 0xE9, 0xCA, 0x20, 0xB3}}

type ITensorUInt64BitStatics2Interface interface {
	win32.IInspectableInterface
	CreateFromShapeArrayAndDataArray(shapeLength uint32, shape *int64, dataLength uint32, data *uint64) *ITensorUInt64Bit
	CreateFromBuffer(shapeLength uint32, shape *int64, buffer *IBuffer) *ITensorUInt64Bit
}

type ITensorUInt64BitStatics2Vtbl struct {
	win32.IInspectableVtbl
	CreateFromShapeArrayAndDataArray uintptr
	CreateFromBuffer                 uintptr
}

type ITensorUInt64BitStatics2 struct {
	win32.IInspectable
}

func (this *ITensorUInt64BitStatics2) Vtbl() *ITensorUInt64BitStatics2Vtbl {
	return (*ITensorUInt64BitStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorUInt64BitStatics2) CreateFromShapeArrayAndDataArray(shapeLength uint32, shape *int64, dataLength uint32, data *uint64) *ITensorUInt64Bit {
	var _result *ITensorUInt64Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromShapeArrayAndDataArray, uintptr(unsafe.Pointer(this)), uintptr(shapeLength), uintptr(unsafe.Pointer(shape)), uintptr(dataLength), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorUInt64BitStatics2) CreateFromBuffer(shapeLength uint32, shape *int64, buffer *IBuffer) *ITensorUInt64Bit {
	var _result *ITensorUInt64Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromBuffer, uintptr(unsafe.Pointer(this)), uintptr(shapeLength), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 58E1AE27-622B-48E3-BE22-D867AED1DAAC
var IID_ITensorUInt8Bit = syscall.GUID{0x58E1AE27, 0x622B, 0x48E3,
	[8]byte{0xBE, 0x22, 0xD8, 0x67, 0xAE, 0xD1, 0xDA, 0xAC}}

type ITensorUInt8BitInterface interface {
	win32.IInspectableInterface
	GetAsVectorView() *IVectorView[byte]
}

type ITensorUInt8BitVtbl struct {
	win32.IInspectableVtbl
	GetAsVectorView uintptr
}

type ITensorUInt8Bit struct {
	win32.IInspectable
}

func (this *ITensorUInt8Bit) Vtbl() *ITensorUInt8BitVtbl {
	return (*ITensorUInt8BitVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorUInt8Bit) GetAsVectorView() *IVectorView[byte] {
	var _result *IVectorView[byte]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAsVectorView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 05F67583-BC24-4220-8A41-2DCD8C5ED33C
var IID_ITensorUInt8BitStatics = syscall.GUID{0x05F67583, 0xBC24, 0x4220,
	[8]byte{0x8A, 0x41, 0x2D, 0xCD, 0x8C, 0x5E, 0xD3, 0x3C}}

type ITensorUInt8BitStaticsInterface interface {
	win32.IInspectableInterface
	Create() *ITensorUInt8Bit
	Create2(shape *IIterable[int64]) *ITensorUInt8Bit
	CreateFromArray(shape *IIterable[int64], dataLength uint32, data *byte) *ITensorUInt8Bit
	CreateFromIterable(shape *IIterable[int64], data *IIterable[byte]) *ITensorUInt8Bit
}

type ITensorUInt8BitStaticsVtbl struct {
	win32.IInspectableVtbl
	Create             uintptr
	Create2            uintptr
	CreateFromArray    uintptr
	CreateFromIterable uintptr
}

type ITensorUInt8BitStatics struct {
	win32.IInspectable
}

func (this *ITensorUInt8BitStatics) Vtbl() *ITensorUInt8BitStaticsVtbl {
	return (*ITensorUInt8BitStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorUInt8BitStatics) Create() *ITensorUInt8Bit {
	var _result *ITensorUInt8Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorUInt8BitStatics) Create2(shape *IIterable[int64]) *ITensorUInt8Bit {
	var _result *ITensorUInt8Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorUInt8BitStatics) CreateFromArray(shape *IIterable[int64], dataLength uint32, data *byte) *ITensorUInt8Bit {
	var _result *ITensorUInt8Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(dataLength), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorUInt8BitStatics) CreateFromIterable(shape *IIterable[int64], data *IIterable[byte]) *ITensorUInt8Bit {
	var _result *ITensorUInt8Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromIterable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2BA042D6-373E-5A3A-A2FC-A6C41BD52789
var IID_ITensorUInt8BitStatics2 = syscall.GUID{0x2BA042D6, 0x373E, 0x5A3A,
	[8]byte{0xA2, 0xFC, 0xA6, 0xC4, 0x1B, 0xD5, 0x27, 0x89}}

type ITensorUInt8BitStatics2Interface interface {
	win32.IInspectableInterface
	CreateFromShapeArrayAndDataArray(shapeLength uint32, shape *int64, dataLength uint32, data *byte) *ITensorUInt8Bit
	CreateFromBuffer(shapeLength uint32, shape *int64, buffer *IBuffer) *ITensorUInt8Bit
}

type ITensorUInt8BitStatics2Vtbl struct {
	win32.IInspectableVtbl
	CreateFromShapeArrayAndDataArray uintptr
	CreateFromBuffer                 uintptr
}

type ITensorUInt8BitStatics2 struct {
	win32.IInspectable
}

func (this *ITensorUInt8BitStatics2) Vtbl() *ITensorUInt8BitStatics2Vtbl {
	return (*ITensorUInt8BitStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITensorUInt8BitStatics2) CreateFromShapeArrayAndDataArray(shapeLength uint32, shape *int64, dataLength uint32, data *byte) *ITensorUInt8Bit {
	var _result *ITensorUInt8Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromShapeArrayAndDataArray, uintptr(unsafe.Pointer(this)), uintptr(shapeLength), uintptr(unsafe.Pointer(shape)), uintptr(dataLength), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ITensorUInt8BitStatics2) CreateFromBuffer(shapeLength uint32, shape *int64, buffer *IBuffer) *ITensorUInt8Bit {
	var _result *ITensorUInt8Bit
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromBuffer, uintptr(unsafe.Pointer(this)), uintptr(shapeLength), uintptr(unsafe.Pointer(shape)), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type ImageFeatureDescriptor struct {
	RtClass
	*IImageFeatureDescriptor
}

type ImageFeatureValue struct {
	RtClass
	*IImageFeatureValue
}

func NewIImageFeatureValueStatics() *IImageFeatureValueStatics {
	var p *IImageFeatureValueStatics
	hs := NewHStr("Windows.AI.MachineLearning.ImageFeatureValue")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IImageFeatureValueStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type LearningModel struct {
	RtClass
	*ILearningModel
}

func NewILearningModelStatics() *ILearningModelStatics {
	var p *ILearningModelStatics
	hs := NewHStr("Windows.AI.MachineLearning.LearningModel")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ILearningModelStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type LearningModelBinding struct {
	RtClass
	*ILearningModelBinding
}

func NewLearningModelBinding_CreateFromSession(session *ILearningModelSession) *LearningModelBinding {
	hs := NewHStr("Windows.AI.MachineLearning.LearningModelBinding")
	var pFac *ILearningModelBindingFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ILearningModelBindingFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ILearningModelBinding
	p = pFac.CreateFromSession(session)
	result := &LearningModelBinding{
		RtClass:               RtClass{PInspect: &p.IInspectable},
		ILearningModelBinding: p,
	}
	com.AddToScope(result)
	return result
}

type LearningModelDevice struct {
	RtClass
	*ILearningModelDevice
}

func NewLearningModelDevice_Create(deviceKind LearningModelDeviceKind) *LearningModelDevice {
	hs := NewHStr("Windows.AI.MachineLearning.LearningModelDevice")
	var pFac *ILearningModelDeviceFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ILearningModelDeviceFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ILearningModelDevice
	p = pFac.Create(deviceKind)
	result := &LearningModelDevice{
		RtClass:              RtClass{PInspect: &p.IInspectable},
		ILearningModelDevice: p,
	}
	com.AddToScope(result)
	return result
}

func NewILearningModelDeviceStatics() *ILearningModelDeviceStatics {
	var p *ILearningModelDeviceStatics
	hs := NewHStr("Windows.AI.MachineLearning.LearningModelDevice")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ILearningModelDeviceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type LearningModelEvaluationResult struct {
	RtClass
	*ILearningModelEvaluationResult
}

type LearningModelSession struct {
	RtClass
	*ILearningModelSession
}

func NewLearningModelSession_CreateFromModelOnDeviceWithSessionOptions(model *ILearningModel, deviceToRunOn *ILearningModelDevice, learningModelSessionOptions *ILearningModelSessionOptions) *LearningModelSession {
	hs := NewHStr("Windows.AI.MachineLearning.LearningModelSession")
	var pFac *ILearningModelSessionFactory2
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ILearningModelSessionFactory2, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ILearningModelSession
	p = pFac.CreateFromModelOnDeviceWithSessionOptions(model, deviceToRunOn, learningModelSessionOptions)
	result := &LearningModelSession{
		RtClass:               RtClass{PInspect: &p.IInspectable},
		ILearningModelSession: p,
	}
	com.AddToScope(result)
	return result
}

type LearningModelSessionOptions struct {
	RtClass
	*ILearningModelSessionOptions
}

func NewLearningModelSessionOptions() *LearningModelSessionOptions {
	hs := NewHStr("Windows.AI.MachineLearning.LearningModelSessionOptions")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &LearningModelSessionOptions{
		RtClass:                      RtClass{PInspect: p},
		ILearningModelSessionOptions: (*ILearningModelSessionOptions)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type MapFeatureDescriptor struct {
	RtClass
	*IMapFeatureDescriptor
}

type SequenceFeatureDescriptor struct {
	RtClass
	*ISequenceFeatureDescriptor
}

type TensorBoolean struct {
	RtClass
	*ITensorBoolean
}

func NewITensorBooleanStatics2() *ITensorBooleanStatics2 {
	var p *ITensorBooleanStatics2
	hs := NewHStr("Windows.AI.MachineLearning.TensorBoolean")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITensorBooleanStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewITensorBooleanStatics() *ITensorBooleanStatics {
	var p *ITensorBooleanStatics
	hs := NewHStr("Windows.AI.MachineLearning.TensorBoolean")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITensorBooleanStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type TensorDouble struct {
	RtClass
	*ITensorDouble
}

func NewITensorDoubleStatics() *ITensorDoubleStatics {
	var p *ITensorDoubleStatics
	hs := NewHStr("Windows.AI.MachineLearning.TensorDouble")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITensorDoubleStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewITensorDoubleStatics2() *ITensorDoubleStatics2 {
	var p *ITensorDoubleStatics2
	hs := NewHStr("Windows.AI.MachineLearning.TensorDouble")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITensorDoubleStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type TensorFeatureDescriptor struct {
	RtClass
	*ITensorFeatureDescriptor
}

type TensorFloat struct {
	RtClass
	*ITensorFloat
}

func NewITensorFloatStatics2() *ITensorFloatStatics2 {
	var p *ITensorFloatStatics2
	hs := NewHStr("Windows.AI.MachineLearning.TensorFloat")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITensorFloatStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewITensorFloatStatics() *ITensorFloatStatics {
	var p *ITensorFloatStatics
	hs := NewHStr("Windows.AI.MachineLearning.TensorFloat")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITensorFloatStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type TensorFloat16Bit struct {
	RtClass
	*ITensorFloat16Bit
}

func NewITensorFloat16BitStatics() *ITensorFloat16BitStatics {
	var p *ITensorFloat16BitStatics
	hs := NewHStr("Windows.AI.MachineLearning.TensorFloat16Bit")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITensorFloat16BitStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewITensorFloat16BitStatics2() *ITensorFloat16BitStatics2 {
	var p *ITensorFloat16BitStatics2
	hs := NewHStr("Windows.AI.MachineLearning.TensorFloat16Bit")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITensorFloat16BitStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type TensorInt16Bit struct {
	RtClass
	*ITensorInt16Bit
}

func NewITensorInt16BitStatics() *ITensorInt16BitStatics {
	var p *ITensorInt16BitStatics
	hs := NewHStr("Windows.AI.MachineLearning.TensorInt16Bit")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITensorInt16BitStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewITensorInt16BitStatics2() *ITensorInt16BitStatics2 {
	var p *ITensorInt16BitStatics2
	hs := NewHStr("Windows.AI.MachineLearning.TensorInt16Bit")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITensorInt16BitStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type TensorInt32Bit struct {
	RtClass
	*ITensorInt32Bit
}

func NewITensorInt32BitStatics() *ITensorInt32BitStatics {
	var p *ITensorInt32BitStatics
	hs := NewHStr("Windows.AI.MachineLearning.TensorInt32Bit")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITensorInt32BitStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewITensorInt32BitStatics2() *ITensorInt32BitStatics2 {
	var p *ITensorInt32BitStatics2
	hs := NewHStr("Windows.AI.MachineLearning.TensorInt32Bit")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITensorInt32BitStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type TensorInt64Bit struct {
	RtClass
	*ITensorInt64Bit
}

func NewITensorInt64BitStatics2() *ITensorInt64BitStatics2 {
	var p *ITensorInt64BitStatics2
	hs := NewHStr("Windows.AI.MachineLearning.TensorInt64Bit")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITensorInt64BitStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewITensorInt64BitStatics() *ITensorInt64BitStatics {
	var p *ITensorInt64BitStatics
	hs := NewHStr("Windows.AI.MachineLearning.TensorInt64Bit")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITensorInt64BitStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type TensorInt8Bit struct {
	RtClass
	*ITensorInt8Bit
}

func NewITensorInt8BitStatics2() *ITensorInt8BitStatics2 {
	var p *ITensorInt8BitStatics2
	hs := NewHStr("Windows.AI.MachineLearning.TensorInt8Bit")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITensorInt8BitStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewITensorInt8BitStatics() *ITensorInt8BitStatics {
	var p *ITensorInt8BitStatics
	hs := NewHStr("Windows.AI.MachineLearning.TensorInt8Bit")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITensorInt8BitStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type TensorString struct {
	RtClass
	*ITensorString
}

func NewITensorStringStatics2() *ITensorStringStatics2 {
	var p *ITensorStringStatics2
	hs := NewHStr("Windows.AI.MachineLearning.TensorString")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITensorStringStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewITensorStringStatics() *ITensorStringStatics {
	var p *ITensorStringStatics
	hs := NewHStr("Windows.AI.MachineLearning.TensorString")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITensorStringStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type TensorUInt16Bit struct {
	RtClass
	*ITensorUInt16Bit
}

func NewITensorUInt16BitStatics2() *ITensorUInt16BitStatics2 {
	var p *ITensorUInt16BitStatics2
	hs := NewHStr("Windows.AI.MachineLearning.TensorUInt16Bit")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITensorUInt16BitStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewITensorUInt16BitStatics() *ITensorUInt16BitStatics {
	var p *ITensorUInt16BitStatics
	hs := NewHStr("Windows.AI.MachineLearning.TensorUInt16Bit")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITensorUInt16BitStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type TensorUInt32Bit struct {
	RtClass
	*ITensorUInt32Bit
}

func NewITensorUInt32BitStatics2() *ITensorUInt32BitStatics2 {
	var p *ITensorUInt32BitStatics2
	hs := NewHStr("Windows.AI.MachineLearning.TensorUInt32Bit")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITensorUInt32BitStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewITensorUInt32BitStatics() *ITensorUInt32BitStatics {
	var p *ITensorUInt32BitStatics
	hs := NewHStr("Windows.AI.MachineLearning.TensorUInt32Bit")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITensorUInt32BitStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type TensorUInt64Bit struct {
	RtClass
	*ITensorUInt64Bit
}

func NewITensorUInt64BitStatics2() *ITensorUInt64BitStatics2 {
	var p *ITensorUInt64BitStatics2
	hs := NewHStr("Windows.AI.MachineLearning.TensorUInt64Bit")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITensorUInt64BitStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewITensorUInt64BitStatics() *ITensorUInt64BitStatics {
	var p *ITensorUInt64BitStatics
	hs := NewHStr("Windows.AI.MachineLearning.TensorUInt64Bit")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITensorUInt64BitStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type TensorUInt8Bit struct {
	RtClass
	*ITensorUInt8Bit
}

func NewITensorUInt8BitStatics() *ITensorUInt8BitStatics {
	var p *ITensorUInt8BitStatics
	hs := NewHStr("Windows.AI.MachineLearning.TensorUInt8Bit")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITensorUInt8BitStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewITensorUInt8BitStatics2() *ITensorUInt8BitStatics2 {
	var p *ITensorUInt8BitStatics2
	hs := NewHStr("Windows.AI.MachineLearning.TensorUInt8Bit")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ITensorUInt8BitStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
