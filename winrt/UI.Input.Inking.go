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
type HandwritingLineHeight int32

const (
	HandwritingLineHeight_Small  HandwritingLineHeight = 0
	HandwritingLineHeight_Medium HandwritingLineHeight = 1
	HandwritingLineHeight_Large  HandwritingLineHeight = 2
)

// enum
type InkDrawingAttributesKind int32

const (
	InkDrawingAttributesKind_Default InkDrawingAttributesKind = 0
	InkDrawingAttributesKind_Pencil  InkDrawingAttributesKind = 1
)

// enum
type InkHighContrastAdjustment int32

const (
	InkHighContrastAdjustment_UseSystemColorsWhenNecessary InkHighContrastAdjustment = 0
	InkHighContrastAdjustment_UseSystemColors              InkHighContrastAdjustment = 1
	InkHighContrastAdjustment_UseOriginalColors            InkHighContrastAdjustment = 2
)

// enum
type InkInputProcessingMode int32

const (
	InkInputProcessingMode_None    InkInputProcessingMode = 0
	InkInputProcessingMode_Inking  InkInputProcessingMode = 1
	InkInputProcessingMode_Erasing InkInputProcessingMode = 2
)

// enum
type InkInputRightDragAction int32

const (
	InkInputRightDragAction_LeaveUnprocessed InkInputRightDragAction = 0
	InkInputRightDragAction_AllowProcessing  InkInputRightDragAction = 1
)

// enum
type InkManipulationMode int32

const (
	InkManipulationMode_Inking    InkManipulationMode = 0
	InkManipulationMode_Erasing   InkManipulationMode = 1
	InkManipulationMode_Selecting InkManipulationMode = 2
)

// enum
type InkPersistenceFormat int32

const (
	InkPersistenceFormat_GifWithEmbeddedIsf InkPersistenceFormat = 0
	InkPersistenceFormat_Isf                InkPersistenceFormat = 1
)

// enum
type InkPresenterPredefinedConfiguration int32

const (
	InkPresenterPredefinedConfiguration_SimpleSinglePointer   InkPresenterPredefinedConfiguration = 0
	InkPresenterPredefinedConfiguration_SimpleMultiplePointer InkPresenterPredefinedConfiguration = 1
)

// enum
type InkPresenterStencilKind int32

const (
	InkPresenterStencilKind_Other      InkPresenterStencilKind = 0
	InkPresenterStencilKind_Ruler      InkPresenterStencilKind = 1
	InkPresenterStencilKind_Protractor InkPresenterStencilKind = 2
)

// enum
type InkRecognitionTarget int32

const (
	InkRecognitionTarget_All      InkRecognitionTarget = 0
	InkRecognitionTarget_Selected InkRecognitionTarget = 1
	InkRecognitionTarget_Recent   InkRecognitionTarget = 2
)

// enum
type PenHandedness int32

const (
	PenHandedness_Right PenHandedness = 0
	PenHandedness_Left  PenHandedness = 1
)

// enum
type PenTipShape int32

const (
	PenTipShape_Circle    PenTipShape = 0
	PenTipShape_Rectangle PenTipShape = 1
)

// interfaces

// 97A2176C-6774-48AD-84F0-48F5A9BE74F9
var IID_IInkDrawingAttributes = syscall.GUID{0x97A2176C, 0x6774, 0x48AD,
	[8]byte{0x84, 0xF0, 0x48, 0xF5, 0xA9, 0xBE, 0x74, 0xF9}}

type IInkDrawingAttributesInterface interface {
	win32.IInspectableInterface
	Get_Color() unsafe.Pointer
	Put_Color(value unsafe.Pointer)
	Get_PenTip() PenTipShape
	Put_PenTip(value PenTipShape)
	Get_Size() Size
	Put_Size(value Size)
	Get_IgnorePressure() bool
	Put_IgnorePressure(value bool)
	Get_FitToCurve() bool
	Put_FitToCurve(value bool)
}

type IInkDrawingAttributesVtbl struct {
	win32.IInspectableVtbl
	Get_Color          uintptr
	Put_Color          uintptr
	Get_PenTip         uintptr
	Put_PenTip         uintptr
	Get_Size           uintptr
	Put_Size           uintptr
	Get_IgnorePressure uintptr
	Put_IgnorePressure uintptr
	Get_FitToCurve     uintptr
	Put_FitToCurve     uintptr
}

type IInkDrawingAttributes struct {
	win32.IInspectable
}

func (this *IInkDrawingAttributes) Vtbl() *IInkDrawingAttributesVtbl {
	return (*IInkDrawingAttributesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkDrawingAttributes) Get_Color() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Color, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkDrawingAttributes) Put_Color(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Color, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IInkDrawingAttributes) Get_PenTip() PenTipShape {
	var _result PenTipShape
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PenTip, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkDrawingAttributes) Put_PenTip(value PenTipShape) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PenTip, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IInkDrawingAttributes) Get_Size() Size {
	var _result Size
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Size, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkDrawingAttributes) Put_Size(value Size) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Size, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IInkDrawingAttributes) Get_IgnorePressure() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IgnorePressure, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkDrawingAttributes) Put_IgnorePressure(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IgnorePressure, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IInkDrawingAttributes) Get_FitToCurve() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FitToCurve, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkDrawingAttributes) Put_FitToCurve(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_FitToCurve, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 7CAB6508-8EC4-42FD-A5A5-E4B7D1D5316D
var IID_IInkDrawingAttributes2 = syscall.GUID{0x7CAB6508, 0x8EC4, 0x42FD,
	[8]byte{0xA5, 0xA5, 0xE4, 0xB7, 0xD1, 0xD5, 0x31, 0x6D}}

type IInkDrawingAttributes2Interface interface {
	win32.IInspectableInterface
	Get_PenTipTransform() unsafe.Pointer
	Put_PenTipTransform(value unsafe.Pointer)
	Get_DrawAsHighlighter() bool
	Put_DrawAsHighlighter(value bool)
}

type IInkDrawingAttributes2Vtbl struct {
	win32.IInspectableVtbl
	Get_PenTipTransform   uintptr
	Put_PenTipTransform   uintptr
	Get_DrawAsHighlighter uintptr
	Put_DrawAsHighlighter uintptr
}

type IInkDrawingAttributes2 struct {
	win32.IInspectable
}

func (this *IInkDrawingAttributes2) Vtbl() *IInkDrawingAttributes2Vtbl {
	return (*IInkDrawingAttributes2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkDrawingAttributes2) Get_PenTipTransform() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PenTipTransform, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkDrawingAttributes2) Put_PenTipTransform(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PenTipTransform, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IInkDrawingAttributes2) Get_DrawAsHighlighter() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DrawAsHighlighter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkDrawingAttributes2) Put_DrawAsHighlighter(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DrawAsHighlighter, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 72020002-7D5B-4690-8AF4-E664CBE2B74F
var IID_IInkDrawingAttributes3 = syscall.GUID{0x72020002, 0x7D5B, 0x4690,
	[8]byte{0x8A, 0xF4, 0xE6, 0x64, 0xCB, 0xE2, 0xB7, 0x4F}}

type IInkDrawingAttributes3Interface interface {
	win32.IInspectableInterface
	Get_Kind() InkDrawingAttributesKind
	Get_PencilProperties() *IInkDrawingAttributesPencilProperties
}

type IInkDrawingAttributes3Vtbl struct {
	win32.IInspectableVtbl
	Get_Kind             uintptr
	Get_PencilProperties uintptr
}

type IInkDrawingAttributes3 struct {
	win32.IInspectable
}

func (this *IInkDrawingAttributes3) Vtbl() *IInkDrawingAttributes3Vtbl {
	return (*IInkDrawingAttributes3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkDrawingAttributes3) Get_Kind() InkDrawingAttributesKind {
	var _result InkDrawingAttributesKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Kind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkDrawingAttributes3) Get_PencilProperties() *IInkDrawingAttributesPencilProperties {
	var _result *IInkDrawingAttributesPencilProperties
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PencilProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// EF65DC25-9F19-456D-91A3-BC3A3D91C5FB
var IID_IInkDrawingAttributes4 = syscall.GUID{0xEF65DC25, 0x9F19, 0x456D,
	[8]byte{0x91, 0xA3, 0xBC, 0x3A, 0x3D, 0x91, 0xC5, 0xFB}}

type IInkDrawingAttributes4Interface interface {
	win32.IInspectableInterface
	Get_IgnoreTilt() bool
	Put_IgnoreTilt(value bool)
}

type IInkDrawingAttributes4Vtbl struct {
	win32.IInspectableVtbl
	Get_IgnoreTilt uintptr
	Put_IgnoreTilt uintptr
}

type IInkDrawingAttributes4 struct {
	win32.IInspectable
}

func (this *IInkDrawingAttributes4) Vtbl() *IInkDrawingAttributes4Vtbl {
	return (*IInkDrawingAttributes4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkDrawingAttributes4) Get_IgnoreTilt() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IgnoreTilt, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkDrawingAttributes4) Put_IgnoreTilt(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IgnoreTilt, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// D11AA0BB-0775-4852-AE64-41143A7AE6C9
var IID_IInkDrawingAttributes5 = syscall.GUID{0xD11AA0BB, 0x0775, 0x4852,
	[8]byte{0xAE, 0x64, 0x41, 0x14, 0x3A, 0x7A, 0xE6, 0xC9}}

type IInkDrawingAttributes5Interface interface {
	win32.IInspectableInterface
	Get_ModelerAttributes() *IInkModelerAttributes
}

type IInkDrawingAttributes5Vtbl struct {
	win32.IInspectableVtbl
	Get_ModelerAttributes uintptr
}

type IInkDrawingAttributes5 struct {
	win32.IInspectable
}

func (this *IInkDrawingAttributes5) Vtbl() *IInkDrawingAttributes5Vtbl {
	return (*IInkDrawingAttributes5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkDrawingAttributes5) Get_ModelerAttributes() *IInkModelerAttributes {
	var _result *IInkModelerAttributes
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ModelerAttributes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4F2534CB-2D86-41BB-B0E8-E4C2A0253C52
var IID_IInkDrawingAttributesPencilProperties = syscall.GUID{0x4F2534CB, 0x2D86, 0x41BB,
	[8]byte{0xB0, 0xE8, 0xE4, 0xC2, 0xA0, 0x25, 0x3C, 0x52}}

type IInkDrawingAttributesPencilPropertiesInterface interface {
	win32.IInspectableInterface
	Get_Opacity() float64
	Put_Opacity(value float64)
}

type IInkDrawingAttributesPencilPropertiesVtbl struct {
	win32.IInspectableVtbl
	Get_Opacity uintptr
	Put_Opacity uintptr
}

type IInkDrawingAttributesPencilProperties struct {
	win32.IInspectable
}

func (this *IInkDrawingAttributesPencilProperties) Vtbl() *IInkDrawingAttributesPencilPropertiesVtbl {
	return (*IInkDrawingAttributesPencilPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkDrawingAttributesPencilProperties) Get_Opacity() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Opacity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkDrawingAttributesPencilProperties) Put_Opacity(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Opacity, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// F731E03F-1A65-4862-96CB-6E1665E17F6D
var IID_IInkDrawingAttributesStatics = syscall.GUID{0xF731E03F, 0x1A65, 0x4862,
	[8]byte{0x96, 0xCB, 0x6E, 0x16, 0x65, 0xE1, 0x7F, 0x6D}}

type IInkDrawingAttributesStaticsInterface interface {
	win32.IInspectableInterface
	CreateForPencil() *IInkDrawingAttributes
}

type IInkDrawingAttributesStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateForPencil uintptr
}

type IInkDrawingAttributesStatics struct {
	win32.IInspectable
}

func (this *IInkDrawingAttributesStatics) Vtbl() *IInkDrawingAttributesStaticsVtbl {
	return (*IInkDrawingAttributesStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkDrawingAttributesStatics) CreateForPencil() *IInkDrawingAttributes {
	var _result *IInkDrawingAttributes
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateForPencil, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 93A68DC4-0B7B-49D7-B34F-9901E524DCF2
var IID_IInkInputConfiguration = syscall.GUID{0x93A68DC4, 0x0B7B, 0x49D7,
	[8]byte{0xB3, 0x4F, 0x99, 0x01, 0xE5, 0x24, 0xDC, 0xF2}}

type IInkInputConfigurationInterface interface {
	win32.IInspectableInterface
	Get_IsPrimaryBarrelButtonInputEnabled() bool
	Put_IsPrimaryBarrelButtonInputEnabled(value bool)
	Get_IsEraserInputEnabled() bool
	Put_IsEraserInputEnabled(value bool)
}

type IInkInputConfigurationVtbl struct {
	win32.IInspectableVtbl
	Get_IsPrimaryBarrelButtonInputEnabled uintptr
	Put_IsPrimaryBarrelButtonInputEnabled uintptr
	Get_IsEraserInputEnabled              uintptr
	Put_IsEraserInputEnabled              uintptr
}

type IInkInputConfiguration struct {
	win32.IInspectable
}

func (this *IInkInputConfiguration) Vtbl() *IInkInputConfigurationVtbl {
	return (*IInkInputConfigurationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkInputConfiguration) Get_IsPrimaryBarrelButtonInputEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPrimaryBarrelButtonInputEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkInputConfiguration) Put_IsPrimaryBarrelButtonInputEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsPrimaryBarrelButtonInputEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IInkInputConfiguration) Get_IsEraserInputEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEraserInputEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkInputConfiguration) Put_IsEraserInputEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsEraserInputEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 6AC2272E-81B4-5CC4-A36D-D057C387DFDA
var IID_IInkInputConfiguration2 = syscall.GUID{0x6AC2272E, 0x81B4, 0x5CC4,
	[8]byte{0xA3, 0x6D, 0xD0, 0x57, 0xC3, 0x87, 0xDF, 0xDA}}

type IInkInputConfiguration2Interface interface {
	win32.IInspectableInterface
	Get_IsPenHapticFeedbackEnabled() bool
	Put_IsPenHapticFeedbackEnabled(value bool)
}

type IInkInputConfiguration2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsPenHapticFeedbackEnabled uintptr
	Put_IsPenHapticFeedbackEnabled uintptr
}

type IInkInputConfiguration2 struct {
	win32.IInspectable
}

func (this *IInkInputConfiguration2) Vtbl() *IInkInputConfiguration2Vtbl {
	return (*IInkInputConfiguration2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkInputConfiguration2) Get_IsPenHapticFeedbackEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsPenHapticFeedbackEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkInputConfiguration2) Put_IsPenHapticFeedbackEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsPenHapticFeedbackEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 2778D85E-33CA-4B06-A6D3-AC3945116D37
var IID_IInkInputProcessingConfiguration = syscall.GUID{0x2778D85E, 0x33CA, 0x4B06,
	[8]byte{0xA6, 0xD3, 0xAC, 0x39, 0x45, 0x11, 0x6D, 0x37}}

type IInkInputProcessingConfigurationInterface interface {
	win32.IInspectableInterface
	Get_Mode() InkInputProcessingMode
	Put_Mode(value InkInputProcessingMode)
	Get_RightDragAction() InkInputRightDragAction
	Put_RightDragAction(value InkInputRightDragAction)
}

type IInkInputProcessingConfigurationVtbl struct {
	win32.IInspectableVtbl
	Get_Mode            uintptr
	Put_Mode            uintptr
	Get_RightDragAction uintptr
	Put_RightDragAction uintptr
}

type IInkInputProcessingConfiguration struct {
	win32.IInspectable
}

func (this *IInkInputProcessingConfiguration) Vtbl() *IInkInputProcessingConfigurationVtbl {
	return (*IInkInputProcessingConfigurationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkInputProcessingConfiguration) Get_Mode() InkInputProcessingMode {
	var _result InkInputProcessingMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkInputProcessingConfiguration) Put_Mode(value InkInputProcessingMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Mode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IInkInputProcessingConfiguration) Get_RightDragAction() InkInputRightDragAction {
	var _result InkInputRightDragAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RightDragAction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkInputProcessingConfiguration) Put_RightDragAction(value InkInputRightDragAction) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RightDragAction, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 4744737D-671B-4163-9C95-4E8D7A035FE1
var IID_IInkManager = syscall.GUID{0x4744737D, 0x671B, 0x4163,
	[8]byte{0x9C, 0x95, 0x4E, 0x8D, 0x7A, 0x03, 0x5F, 0xE1}}

type IInkManagerInterface interface {
	win32.IInspectableInterface
	Get_Mode() InkManipulationMode
	Put_Mode(value InkManipulationMode)
	ProcessPointerDown(pointerPoint *IPointerPoint)
	ProcessPointerUpdate(pointerPoint *IPointerPoint) interface{}
	ProcessPointerUp(pointerPoint *IPointerPoint) Rect
	SetDefaultDrawingAttributes(drawingAttributes *IInkDrawingAttributes)
	RecognizeAsync2(recognitionTarget InkRecognitionTarget) *IAsyncOperation[*IVectorView[*IInkRecognitionResult]]
}

type IInkManagerVtbl struct {
	win32.IInspectableVtbl
	Get_Mode                    uintptr
	Put_Mode                    uintptr
	ProcessPointerDown          uintptr
	ProcessPointerUpdate        uintptr
	ProcessPointerUp            uintptr
	SetDefaultDrawingAttributes uintptr
	RecognizeAsync2             uintptr
}

type IInkManager struct {
	win32.IInspectable
}

func (this *IInkManager) Vtbl() *IInkManagerVtbl {
	return (*IInkManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkManager) Get_Mode() InkManipulationMode {
	var _result InkManipulationMode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Mode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkManager) Put_Mode(value InkManipulationMode) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Mode, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IInkManager) ProcessPointerDown(pointerPoint *IPointerPoint) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ProcessPointerDown, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pointerPoint)))
	_ = _hr
}

func (this *IInkManager) ProcessPointerUpdate(pointerPoint *IPointerPoint) interface{} {
	var _result interface{}
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ProcessPointerUpdate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pointerPoint)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkManager) ProcessPointerUp(pointerPoint *IPointerPoint) Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ProcessPointerUp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pointerPoint)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkManager) SetDefaultDrawingAttributes(drawingAttributes *IInkDrawingAttributes) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetDefaultDrawingAttributes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(drawingAttributes)))
	_ = _hr
}

func (this *IInkManager) RecognizeAsync2(recognitionTarget InkRecognitionTarget) *IAsyncOperation[*IVectorView[*IInkRecognitionResult]] {
	var _result *IAsyncOperation[*IVectorView[*IInkRecognitionResult]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RecognizeAsync2, uintptr(unsafe.Pointer(this)), uintptr(recognitionTarget), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BAD31F27-0CD9-4BFD-B6F3-9E03BA8D7454
var IID_IInkModelerAttributes = syscall.GUID{0xBAD31F27, 0x0CD9, 0x4BFD,
	[8]byte{0xB6, 0xF3, 0x9E, 0x03, 0xBA, 0x8D, 0x74, 0x54}}

type IInkModelerAttributesInterface interface {
	win32.IInspectableInterface
	Get_PredictionTime() TimeSpan
	Put_PredictionTime(value TimeSpan)
	Get_ScalingFactor() float32
	Put_ScalingFactor(value float32)
}

type IInkModelerAttributesVtbl struct {
	win32.IInspectableVtbl
	Get_PredictionTime uintptr
	Put_PredictionTime uintptr
	Get_ScalingFactor  uintptr
	Put_ScalingFactor  uintptr
}

type IInkModelerAttributes struct {
	win32.IInspectable
}

func (this *IInkModelerAttributes) Vtbl() *IInkModelerAttributesVtbl {
	return (*IInkModelerAttributesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkModelerAttributes) Get_PredictionTime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PredictionTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkModelerAttributes) Put_PredictionTime(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PredictionTime, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IInkModelerAttributes) Get_ScalingFactor() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ScalingFactor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkModelerAttributes) Put_ScalingFactor(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ScalingFactor, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 86D1D09A-4EF8-5E25-B7BC-B65424F16BB3
var IID_IInkModelerAttributes2 = syscall.GUID{0x86D1D09A, 0x4EF8, 0x5E25,
	[8]byte{0xB7, 0xBC, 0xB6, 0x54, 0x24, 0xF1, 0x6B, 0xB3}}

type IInkModelerAttributes2Interface interface {
	win32.IInspectableInterface
	Get_UseVelocityBasedPressure() bool
	Put_UseVelocityBasedPressure(value bool)
}

type IInkModelerAttributes2Vtbl struct {
	win32.IInspectableVtbl
	Get_UseVelocityBasedPressure uintptr
	Put_UseVelocityBasedPressure uintptr
}

type IInkModelerAttributes2 struct {
	win32.IInspectable
}

func (this *IInkModelerAttributes2) Vtbl() *IInkModelerAttributes2Vtbl {
	return (*IInkModelerAttributes2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkModelerAttributes2) Get_UseVelocityBasedPressure() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UseVelocityBasedPressure, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkModelerAttributes2) Put_UseVelocityBasedPressure(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_UseVelocityBasedPressure, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 9F87272B-858C-46A5-9B41-D195970459FD
var IID_IInkPoint = syscall.GUID{0x9F87272B, 0x858C, 0x46A5,
	[8]byte{0x9B, 0x41, 0xD1, 0x95, 0x97, 0x04, 0x59, 0xFD}}

type IInkPointInterface interface {
	win32.IInspectableInterface
	Get_Position() Point
	Get_Pressure() float32
}

type IInkPointVtbl struct {
	win32.IInspectableVtbl
	Get_Position uintptr
	Get_Pressure uintptr
}

type IInkPoint struct {
	win32.IInspectable
}

func (this *IInkPoint) Vtbl() *IInkPointVtbl {
	return (*IInkPointVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkPoint) Get_Position() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkPoint) Get_Pressure() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Pressure, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// FBA9C3F7-AE56-4D5C-BD2F-0AC45F5E4AF9
var IID_IInkPoint2 = syscall.GUID{0xFBA9C3F7, 0xAE56, 0x4D5C,
	[8]byte{0xBD, 0x2F, 0x0A, 0xC4, 0x5F, 0x5E, 0x4A, 0xF9}}

type IInkPoint2Interface interface {
	win32.IInspectableInterface
	Get_TiltX() float32
	Get_TiltY() float32
	Get_Timestamp() uint64
}

type IInkPoint2Vtbl struct {
	win32.IInspectableVtbl
	Get_TiltX     uintptr
	Get_TiltY     uintptr
	Get_Timestamp uintptr
}

type IInkPoint2 struct {
	win32.IInspectable
}

func (this *IInkPoint2) Vtbl() *IInkPoint2Vtbl {
	return (*IInkPoint2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkPoint2) Get_TiltX() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TiltX, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkPoint2) Get_TiltY() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TiltY, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkPoint2) Get_Timestamp() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timestamp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 29E5D51C-C98F-405D-9F3B-E53E31068D4D
var IID_IInkPointFactory = syscall.GUID{0x29E5D51C, 0xC98F, 0x405D,
	[8]byte{0x9F, 0x3B, 0xE5, 0x3E, 0x31, 0x06, 0x8D, 0x4D}}

type IInkPointFactoryInterface interface {
	win32.IInspectableInterface
	CreateInkPoint(position Point, pressure float32) *IInkPoint
}

type IInkPointFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateInkPoint uintptr
}

type IInkPointFactory struct {
	win32.IInspectable
}

func (this *IInkPointFactory) Vtbl() *IInkPointFactoryVtbl {
	return (*IInkPointFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkPointFactory) CreateInkPoint(position Point, pressure float32) *IInkPoint {
	var _result *IInkPoint
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateInkPoint, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&position)), uintptr(pressure), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E0145E85-DAFF-45F2-AD69-050D8256A209
var IID_IInkPointFactory2 = syscall.GUID{0xE0145E85, 0xDAFF, 0x45F2,
	[8]byte{0xAD, 0x69, 0x05, 0x0D, 0x82, 0x56, 0xA2, 0x09}}

type IInkPointFactory2Interface interface {
	win32.IInspectableInterface
	CreateInkPointWithTiltAndTimestamp(position Point, pressure float32, tiltX float32, tiltY float32, timestamp uint64) *IInkPoint
}

type IInkPointFactory2Vtbl struct {
	win32.IInspectableVtbl
	CreateInkPointWithTiltAndTimestamp uintptr
}

type IInkPointFactory2 struct {
	win32.IInspectable
}

func (this *IInkPointFactory2) Vtbl() *IInkPointFactory2Vtbl {
	return (*IInkPointFactory2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkPointFactory2) CreateInkPointWithTiltAndTimestamp(position Point, pressure float32, tiltX float32, tiltY float32, timestamp uint64) *IInkPoint {
	var _result *IInkPoint
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateInkPointWithTiltAndTimestamp, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&position)), uintptr(pressure), uintptr(tiltX), uintptr(tiltY), uintptr(timestamp), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A69B70E2-887B-458F-B173-4FE4438930A3
var IID_IInkPresenter = syscall.GUID{0xA69B70E2, 0x887B, 0x458F,
	[8]byte{0xB1, 0x73, 0x4F, 0xE4, 0x43, 0x89, 0x30, 0xA3}}

type IInkPresenterInterface interface {
	win32.IInspectableInterface
	Get_IsInputEnabled() bool
	Put_IsInputEnabled(value bool)
	Get_InputDeviceTypes() unsafe.Pointer
	Put_InputDeviceTypes(value unsafe.Pointer)
	Get_UnprocessedInput() *IInkUnprocessedInput
	Get_StrokeInput() *IInkStrokeInput
	Get_InputProcessingConfiguration() *IInkInputProcessingConfiguration
	Get_StrokeContainer() *IInkStrokeContainer
	Put_StrokeContainer(value *IInkStrokeContainer)
	CopyDefaultDrawingAttributes() *IInkDrawingAttributes
	UpdateDefaultDrawingAttributes(value *IInkDrawingAttributes)
	ActivateCustomDrying() *IInkSynchronizer
	SetPredefinedConfiguration(value InkPresenterPredefinedConfiguration)
	Add_StrokesCollected(handler TypedEventHandler[*IInkPresenter, *IInkStrokesCollectedEventArgs]) EventRegistrationToken
	Remove_StrokesCollected(cookie EventRegistrationToken)
	Add_StrokesErased(handler TypedEventHandler[*IInkPresenter, *IInkStrokesErasedEventArgs]) EventRegistrationToken
	Remove_StrokesErased(cookie EventRegistrationToken)
}

type IInkPresenterVtbl struct {
	win32.IInspectableVtbl
	Get_IsInputEnabled               uintptr
	Put_IsInputEnabled               uintptr
	Get_InputDeviceTypes             uintptr
	Put_InputDeviceTypes             uintptr
	Get_UnprocessedInput             uintptr
	Get_StrokeInput                  uintptr
	Get_InputProcessingConfiguration uintptr
	Get_StrokeContainer              uintptr
	Put_StrokeContainer              uintptr
	CopyDefaultDrawingAttributes     uintptr
	UpdateDefaultDrawingAttributes   uintptr
	ActivateCustomDrying             uintptr
	SetPredefinedConfiguration       uintptr
	Add_StrokesCollected             uintptr
	Remove_StrokesCollected          uintptr
	Add_StrokesErased                uintptr
	Remove_StrokesErased             uintptr
}

type IInkPresenter struct {
	win32.IInspectable
}

func (this *IInkPresenter) Vtbl() *IInkPresenterVtbl {
	return (*IInkPresenterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkPresenter) Get_IsInputEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsInputEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkPresenter) Put_IsInputEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsInputEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IInkPresenter) Get_InputDeviceTypes() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InputDeviceTypes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkPresenter) Put_InputDeviceTypes(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InputDeviceTypes, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IInkPresenter) Get_UnprocessedInput() *IInkUnprocessedInput {
	var _result *IInkUnprocessedInput
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UnprocessedInput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInkPresenter) Get_StrokeInput() *IInkStrokeInput {
	var _result *IInkStrokeInput
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StrokeInput, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInkPresenter) Get_InputProcessingConfiguration() *IInkInputProcessingConfiguration {
	var _result *IInkInputProcessingConfiguration
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InputProcessingConfiguration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInkPresenter) Get_StrokeContainer() *IInkStrokeContainer {
	var _result *IInkStrokeContainer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StrokeContainer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInkPresenter) Put_StrokeContainer(value *IInkStrokeContainer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StrokeContainer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IInkPresenter) CopyDefaultDrawingAttributes() *IInkDrawingAttributes {
	var _result *IInkDrawingAttributes
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CopyDefaultDrawingAttributes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInkPresenter) UpdateDefaultDrawingAttributes(value *IInkDrawingAttributes) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UpdateDefaultDrawingAttributes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IInkPresenter) ActivateCustomDrying() *IInkSynchronizer {
	var _result *IInkSynchronizer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ActivateCustomDrying, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInkPresenter) SetPredefinedConfiguration(value InkPresenterPredefinedConfiguration) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetPredefinedConfiguration, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IInkPresenter) Add_StrokesCollected(handler TypedEventHandler[*IInkPresenter, *IInkStrokesCollectedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StrokesCollected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkPresenter) Remove_StrokesCollected(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StrokesCollected, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IInkPresenter) Add_StrokesErased(handler TypedEventHandler[*IInkPresenter, *IInkStrokesErasedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StrokesErased, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkPresenter) Remove_StrokesErased(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StrokesErased, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

// CF53E612-9A34-11E6-9F33-A24FC0D9649C
var IID_IInkPresenter2 = syscall.GUID{0xCF53E612, 0x9A34, 0x11E6,
	[8]byte{0x9F, 0x33, 0xA2, 0x4F, 0xC0, 0xD9, 0x64, 0x9C}}

type IInkPresenter2Interface interface {
	win32.IInspectableInterface
	Get_HighContrastAdjustment() InkHighContrastAdjustment
	Put_HighContrastAdjustment(value InkHighContrastAdjustment)
}

type IInkPresenter2Vtbl struct {
	win32.IInspectableVtbl
	Get_HighContrastAdjustment uintptr
	Put_HighContrastAdjustment uintptr
}

type IInkPresenter2 struct {
	win32.IInspectable
}

func (this *IInkPresenter2) Vtbl() *IInkPresenter2Vtbl {
	return (*IInkPresenter2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkPresenter2) Get_HighContrastAdjustment() InkHighContrastAdjustment {
	var _result InkHighContrastAdjustment
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HighContrastAdjustment, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkPresenter2) Put_HighContrastAdjustment(value InkHighContrastAdjustment) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_HighContrastAdjustment, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 51E1CE89-D37D-4A90-83FC-7F5E9DFBF217
var IID_IInkPresenter3 = syscall.GUID{0x51E1CE89, 0xD37D, 0x4A90,
	[8]byte{0x83, 0xFC, 0x7F, 0x5E, 0x9D, 0xFB, 0xF2, 0x17}}

type IInkPresenter3Interface interface {
	win32.IInspectableInterface
	Get_InputConfiguration() *IInkInputConfiguration
}

type IInkPresenter3Vtbl struct {
	win32.IInspectableVtbl
	Get_InputConfiguration uintptr
}

type IInkPresenter3 struct {
	win32.IInspectable
}

func (this *IInkPresenter3) Vtbl() *IInkPresenter3Vtbl {
	return (*IInkPresenter3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkPresenter3) Get_InputConfiguration() *IInkInputConfiguration {
	var _result *IInkInputConfiguration
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InputConfiguration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7DE3F2AA-EF6C-4E91-A73B-5B70D56FBD17
var IID_IInkPresenterProtractor = syscall.GUID{0x7DE3F2AA, 0xEF6C, 0x4E91,
	[8]byte{0xA7, 0x3B, 0x5B, 0x70, 0xD5, 0x6F, 0xBD, 0x17}}

type IInkPresenterProtractorInterface interface {
	win32.IInspectableInterface
	Get_AreTickMarksVisible() bool
	Put_AreTickMarksVisible(value bool)
	Get_AreRaysVisible() bool
	Put_AreRaysVisible(value bool)
	Get_IsCenterMarkerVisible() bool
	Put_IsCenterMarkerVisible(value bool)
	Get_IsAngleReadoutVisible() bool
	Put_IsAngleReadoutVisible(value bool)
	Get_IsResizable() bool
	Put_IsResizable(value bool)
	Get_Radius() float64
	Put_Radius(value float64)
	Get_AccentColor() unsafe.Pointer
	Put_AccentColor(value unsafe.Pointer)
}

type IInkPresenterProtractorVtbl struct {
	win32.IInspectableVtbl
	Get_AreTickMarksVisible   uintptr
	Put_AreTickMarksVisible   uintptr
	Get_AreRaysVisible        uintptr
	Put_AreRaysVisible        uintptr
	Get_IsCenterMarkerVisible uintptr
	Put_IsCenterMarkerVisible uintptr
	Get_IsAngleReadoutVisible uintptr
	Put_IsAngleReadoutVisible uintptr
	Get_IsResizable           uintptr
	Put_IsResizable           uintptr
	Get_Radius                uintptr
	Put_Radius                uintptr
	Get_AccentColor           uintptr
	Put_AccentColor           uintptr
}

type IInkPresenterProtractor struct {
	win32.IInspectable
}

func (this *IInkPresenterProtractor) Vtbl() *IInkPresenterProtractorVtbl {
	return (*IInkPresenterProtractorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkPresenterProtractor) Get_AreTickMarksVisible() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AreTickMarksVisible, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkPresenterProtractor) Put_AreTickMarksVisible(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AreTickMarksVisible, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IInkPresenterProtractor) Get_AreRaysVisible() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AreRaysVisible, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkPresenterProtractor) Put_AreRaysVisible(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AreRaysVisible, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IInkPresenterProtractor) Get_IsCenterMarkerVisible() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCenterMarkerVisible, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkPresenterProtractor) Put_IsCenterMarkerVisible(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsCenterMarkerVisible, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IInkPresenterProtractor) Get_IsAngleReadoutVisible() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsAngleReadoutVisible, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkPresenterProtractor) Put_IsAngleReadoutVisible(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsAngleReadoutVisible, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IInkPresenterProtractor) Get_IsResizable() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsResizable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkPresenterProtractor) Put_IsResizable(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsResizable, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IInkPresenterProtractor) Get_Radius() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Radius, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkPresenterProtractor) Put_Radius(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Radius, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IInkPresenterProtractor) Get_AccentColor() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AccentColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkPresenterProtractor) Put_AccentColor(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AccentColor, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 320103C9-68FA-47E9-8127-8370711FC46C
var IID_IInkPresenterProtractorFactory = syscall.GUID{0x320103C9, 0x68FA, 0x47E9,
	[8]byte{0x81, 0x27, 0x83, 0x70, 0x71, 0x1F, 0xC4, 0x6C}}

type IInkPresenterProtractorFactoryInterface interface {
	win32.IInspectableInterface
	Create(inkPresenter *IInkPresenter) *IInkPresenterProtractor
}

type IInkPresenterProtractorFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IInkPresenterProtractorFactory struct {
	win32.IInspectable
}

func (this *IInkPresenterProtractorFactory) Vtbl() *IInkPresenterProtractorFactoryVtbl {
	return (*IInkPresenterProtractorFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkPresenterProtractorFactory) Create(inkPresenter *IInkPresenter) *IInkPresenterProtractor {
	var _result *IInkPresenterProtractor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(inkPresenter)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6CDA7D5A-DEC7-4DD7-877A-2133F183D48A
var IID_IInkPresenterRuler = syscall.GUID{0x6CDA7D5A, 0xDEC7, 0x4DD7,
	[8]byte{0x87, 0x7A, 0x21, 0x33, 0xF1, 0x83, 0xD4, 0x8A}}

type IInkPresenterRulerInterface interface {
	win32.IInspectableInterface
	Get_Length() float64
	Put_Length(value float64)
	Get_Width() float64
	Put_Width(value float64)
}

type IInkPresenterRulerVtbl struct {
	win32.IInspectableVtbl
	Get_Length uintptr
	Put_Length uintptr
	Get_Width  uintptr
	Put_Width  uintptr
}

type IInkPresenterRuler struct {
	win32.IInspectable
}

func (this *IInkPresenterRuler) Vtbl() *IInkPresenterRulerVtbl {
	return (*IInkPresenterRulerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkPresenterRuler) Get_Length() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Length, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkPresenterRuler) Put_Length(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Length, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IInkPresenterRuler) Get_Width() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Width, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkPresenterRuler) Put_Width(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Width, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 45130DC1-BC61-44D4-A423-54712AE671C4
var IID_IInkPresenterRuler2 = syscall.GUID{0x45130DC1, 0xBC61, 0x44D4,
	[8]byte{0xA4, 0x23, 0x54, 0x71, 0x2A, 0xE6, 0x71, 0xC4}}

type IInkPresenterRuler2Interface interface {
	win32.IInspectableInterface
	Get_AreTickMarksVisible() bool
	Put_AreTickMarksVisible(value bool)
	Get_IsCompassVisible() bool
	Put_IsCompassVisible(value bool)
}

type IInkPresenterRuler2Vtbl struct {
	win32.IInspectableVtbl
	Get_AreTickMarksVisible uintptr
	Put_AreTickMarksVisible uintptr
	Get_IsCompassVisible    uintptr
	Put_IsCompassVisible    uintptr
}

type IInkPresenterRuler2 struct {
	win32.IInspectable
}

func (this *IInkPresenterRuler2) Vtbl() *IInkPresenterRuler2Vtbl {
	return (*IInkPresenterRuler2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkPresenterRuler2) Get_AreTickMarksVisible() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AreTickMarksVisible, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkPresenterRuler2) Put_AreTickMarksVisible(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AreTickMarksVisible, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IInkPresenterRuler2) Get_IsCompassVisible() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsCompassVisible, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkPresenterRuler2) Put_IsCompassVisible(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsCompassVisible, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 34361BEB-9001-4A4B-A690-69DBAF63E501
var IID_IInkPresenterRulerFactory = syscall.GUID{0x34361BEB, 0x9001, 0x4A4B,
	[8]byte{0xA6, 0x90, 0x69, 0xDB, 0xAF, 0x63, 0xE5, 0x01}}

type IInkPresenterRulerFactoryInterface interface {
	win32.IInspectableInterface
	Create(inkPresenter *IInkPresenter) *IInkPresenterRuler
}

type IInkPresenterRulerFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IInkPresenterRulerFactory struct {
	win32.IInspectable
}

func (this *IInkPresenterRulerFactory) Vtbl() *IInkPresenterRulerFactoryVtbl {
	return (*IInkPresenterRulerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkPresenterRulerFactory) Create(inkPresenter *IInkPresenter) *IInkPresenterRuler {
	var _result *IInkPresenterRuler
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(inkPresenter)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 30D12D6D-3E06-4D02-B116-277FB5D8ADDC
var IID_IInkPresenterStencil = syscall.GUID{0x30D12D6D, 0x3E06, 0x4D02,
	[8]byte{0xB1, 0x16, 0x27, 0x7F, 0xB5, 0xD8, 0xAD, 0xDC}}

type IInkPresenterStencilInterface interface {
	win32.IInspectableInterface
	Get_Kind() InkPresenterStencilKind
	Get_IsVisible() bool
	Put_IsVisible(value bool)
	Get_BackgroundColor() unsafe.Pointer
	Put_BackgroundColor(value unsafe.Pointer)
	Get_ForegroundColor() unsafe.Pointer
	Put_ForegroundColor(value unsafe.Pointer)
	Get_Transform() unsafe.Pointer
	Put_Transform(value unsafe.Pointer)
}

type IInkPresenterStencilVtbl struct {
	win32.IInspectableVtbl
	Get_Kind            uintptr
	Get_IsVisible       uintptr
	Put_IsVisible       uintptr
	Get_BackgroundColor uintptr
	Put_BackgroundColor uintptr
	Get_ForegroundColor uintptr
	Put_ForegroundColor uintptr
	Get_Transform       uintptr
	Put_Transform       uintptr
}

type IInkPresenterStencil struct {
	win32.IInspectable
}

func (this *IInkPresenterStencil) Vtbl() *IInkPresenterStencilVtbl {
	return (*IInkPresenterStencilVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkPresenterStencil) Get_Kind() InkPresenterStencilKind {
	var _result InkPresenterStencilKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Kind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkPresenterStencil) Get_IsVisible() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsVisible, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkPresenterStencil) Put_IsVisible(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsVisible, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IInkPresenterStencil) Get_BackgroundColor() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BackgroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkPresenterStencil) Put_BackgroundColor(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BackgroundColor, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IInkPresenterStencil) Get_ForegroundColor() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ForegroundColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkPresenterStencil) Put_ForegroundColor(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ForegroundColor, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IInkPresenterStencil) Get_Transform() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Transform, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkPresenterStencil) Put_Transform(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Transform, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 36461A94-5068-40EF-8A05-2C2FB60908A2
var IID_IInkRecognitionResult = syscall.GUID{0x36461A94, 0x5068, 0x40EF,
	[8]byte{0x8A, 0x05, 0x2C, 0x2F, 0xB6, 0x09, 0x08, 0xA2}}

type IInkRecognitionResultInterface interface {
	win32.IInspectableInterface
	Get_BoundingRect() Rect
	GetTextCandidates() *IVectorView[string]
	GetStrokes() *IVectorView[*IInkStroke]
}

type IInkRecognitionResultVtbl struct {
	win32.IInspectableVtbl
	Get_BoundingRect  uintptr
	GetTextCandidates uintptr
	GetStrokes        uintptr
}

type IInkRecognitionResult struct {
	win32.IInspectable
}

func (this *IInkRecognitionResult) Vtbl() *IInkRecognitionResultVtbl {
	return (*IInkRecognitionResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkRecognitionResult) Get_BoundingRect() Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BoundingRect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkRecognitionResult) GetTextCandidates() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetTextCandidates, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInkRecognitionResult) GetStrokes() *IVectorView[*IInkStroke] {
	var _result *IVectorView[*IInkStroke]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStrokes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 077CCEA3-904D-442A-B151-AACA3631C43B
var IID_IInkRecognizer = syscall.GUID{0x077CCEA3, 0x904D, 0x442A,
	[8]byte{0xB1, 0x51, 0xAA, 0xCA, 0x36, 0x31, 0xC4, 0x3B}}

type IInkRecognizerInterface interface {
	win32.IInspectableInterface
	Get_Name() string
}

type IInkRecognizerVtbl struct {
	win32.IInspectableVtbl
	Get_Name uintptr
}

type IInkRecognizer struct {
	win32.IInspectable
}

func (this *IInkRecognizer) Vtbl() *IInkRecognizerVtbl {
	return (*IInkRecognizerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkRecognizer) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// A74D9A31-8047-4698-A912-F82A5085012F
var IID_IInkRecognizerContainer = syscall.GUID{0xA74D9A31, 0x8047, 0x4698,
	[8]byte{0xA9, 0x12, 0xF8, 0x2A, 0x50, 0x85, 0x01, 0x2F}}

type IInkRecognizerContainerInterface interface {
	win32.IInspectableInterface
	SetDefaultRecognizer(recognizer *IInkRecognizer)
	RecognizeAsync(strokeCollection *IInkStrokeContainer, recognitionTarget InkRecognitionTarget) *IAsyncOperation[*IVectorView[*IInkRecognitionResult]]
	GetRecognizers() *IVectorView[*IInkRecognizer]
}

type IInkRecognizerContainerVtbl struct {
	win32.IInspectableVtbl
	SetDefaultRecognizer uintptr
	RecognizeAsync       uintptr
	GetRecognizers       uintptr
}

type IInkRecognizerContainer struct {
	win32.IInspectable
}

func (this *IInkRecognizerContainer) Vtbl() *IInkRecognizerContainerVtbl {
	return (*IInkRecognizerContainerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkRecognizerContainer) SetDefaultRecognizer(recognizer *IInkRecognizer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetDefaultRecognizer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(recognizer)))
	_ = _hr
}

func (this *IInkRecognizerContainer) RecognizeAsync(strokeCollection *IInkStrokeContainer, recognitionTarget InkRecognitionTarget) *IAsyncOperation[*IVectorView[*IInkRecognitionResult]] {
	var _result *IAsyncOperation[*IVectorView[*IInkRecognitionResult]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RecognizeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(strokeCollection)), uintptr(recognitionTarget), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInkRecognizerContainer) GetRecognizers() *IVectorView[*IInkRecognizer] {
	var _result *IVectorView[*IInkRecognizer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetRecognizers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 15144D60-CCE3-4FCF-9D52-11518AB6AFD4
var IID_IInkStroke = syscall.GUID{0x15144D60, 0xCCE3, 0x4FCF,
	[8]byte{0x9D, 0x52, 0x11, 0x51, 0x8A, 0xB6, 0xAF, 0xD4}}

type IInkStrokeInterface interface {
	win32.IInspectableInterface
	Get_DrawingAttributes() *IInkDrawingAttributes
	Put_DrawingAttributes(value *IInkDrawingAttributes)
	Get_BoundingRect() Rect
	Get_Selected() bool
	Put_Selected(value bool)
	Get_Recognized() bool
	GetRenderingSegments() *IVectorView[*IInkStrokeRenderingSegment]
	Clone() *IInkStroke
}

type IInkStrokeVtbl struct {
	win32.IInspectableVtbl
	Get_DrawingAttributes uintptr
	Put_DrawingAttributes uintptr
	Get_BoundingRect      uintptr
	Get_Selected          uintptr
	Put_Selected          uintptr
	Get_Recognized        uintptr
	GetRenderingSegments  uintptr
	Clone                 uintptr
}

type IInkStroke struct {
	win32.IInspectable
}

func (this *IInkStroke) Vtbl() *IInkStrokeVtbl {
	return (*IInkStrokeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkStroke) Get_DrawingAttributes() *IInkDrawingAttributes {
	var _result *IInkDrawingAttributes
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DrawingAttributes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInkStroke) Put_DrawingAttributes(value *IInkDrawingAttributes) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DrawingAttributes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IInkStroke) Get_BoundingRect() Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BoundingRect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkStroke) Get_Selected() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Selected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkStroke) Put_Selected(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Selected, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IInkStroke) Get_Recognized() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Recognized, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkStroke) GetRenderingSegments() *IVectorView[*IInkStrokeRenderingSegment] {
	var _result *IVectorView[*IInkStrokeRenderingSegment]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetRenderingSegments, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInkStroke) Clone() *IInkStroke {
	var _result *IInkStroke
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 5DB9E4F4-BAFA-4DE1-89D3-201B1ED7D89B
var IID_IInkStroke2 = syscall.GUID{0x5DB9E4F4, 0xBAFA, 0x4DE1,
	[8]byte{0x89, 0xD3, 0x20, 0x1B, 0x1E, 0xD7, 0xD8, 0x9B}}

type IInkStroke2Interface interface {
	win32.IInspectableInterface
	Get_PointTransform() unsafe.Pointer
	Put_PointTransform(value unsafe.Pointer)
	GetInkPoints() *IVectorView[*IInkPoint]
}

type IInkStroke2Vtbl struct {
	win32.IInspectableVtbl
	Get_PointTransform uintptr
	Put_PointTransform uintptr
	GetInkPoints       uintptr
}

type IInkStroke2 struct {
	win32.IInspectable
}

func (this *IInkStroke2) Vtbl() *IInkStroke2Vtbl {
	return (*IInkStroke2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkStroke2) Get_PointTransform() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PointTransform, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkStroke2) Put_PointTransform(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PointTransform, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IInkStroke2) GetInkPoints() *IVectorView[*IInkPoint] {
	var _result *IVectorView[*IInkPoint]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetInkPoints, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4A807374-9499-411D-A1C4-68855D03D65F
var IID_IInkStroke3 = syscall.GUID{0x4A807374, 0x9499, 0x411D,
	[8]byte{0xA1, 0xC4, 0x68, 0x85, 0x5D, 0x03, 0xD6, 0x5F}}

type IInkStroke3Interface interface {
	win32.IInspectableInterface
	Get_Id() uint32
	Get_StrokeStartedTime() *IReference[DateTime]
	Put_StrokeStartedTime(value *IReference[DateTime])
	Get_StrokeDuration() *IReference[TimeSpan]
	Put_StrokeDuration(value *IReference[TimeSpan])
}

type IInkStroke3Vtbl struct {
	win32.IInspectableVtbl
	Get_Id                uintptr
	Get_StrokeStartedTime uintptr
	Put_StrokeStartedTime uintptr
	Get_StrokeDuration    uintptr
	Put_StrokeDuration    uintptr
}

type IInkStroke3 struct {
	win32.IInspectable
}

func (this *IInkStroke3) Vtbl() *IInkStroke3Vtbl {
	return (*IInkStroke3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkStroke3) Get_Id() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkStroke3) Get_StrokeStartedTime() *IReference[DateTime] {
	var _result *IReference[DateTime]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StrokeStartedTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInkStroke3) Put_StrokeStartedTime(value *IReference[DateTime]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StrokeStartedTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IInkStroke3) Get_StrokeDuration() *IReference[TimeSpan] {
	var _result *IReference[TimeSpan]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StrokeDuration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInkStroke3) Put_StrokeDuration(value *IReference[TimeSpan]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StrokeDuration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// CD5B62E5-B6E9-5B91-A577-1921D2348690
var IID_IInkStroke4 = syscall.GUID{0xCD5B62E5, 0xB6E9, 0x5B91,
	[8]byte{0xA5, 0x77, 0x19, 0x21, 0xD2, 0x34, 0x86, 0x90}}

type IInkStroke4Interface interface {
	win32.IInspectableInterface
	Get_PointerId() uint32
}

type IInkStroke4Vtbl struct {
	win32.IInspectableVtbl
	Get_PointerId uintptr
}

type IInkStroke4 struct {
	win32.IInspectable
}

func (this *IInkStroke4) Vtbl() *IInkStroke4Vtbl {
	return (*IInkStroke4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkStroke4) Get_PointerId() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PointerId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 82BBD1DC-1C63-41DC-9E07-4B4A70CED801
var IID_IInkStrokeBuilder = syscall.GUID{0x82BBD1DC, 0x1C63, 0x41DC,
	[8]byte{0x9E, 0x07, 0x4B, 0x4A, 0x70, 0xCE, 0xD8, 0x01}}

type IInkStrokeBuilderInterface interface {
	win32.IInspectableInterface
	BeginStroke(pointerPoint *IPointerPoint)
	AppendToStroke(pointerPoint *IPointerPoint) *IPointerPoint
	EndStroke(pointerPoint *IPointerPoint) *IInkStroke
	CreateStroke(points *IIterable[Point]) *IInkStroke
	SetDefaultDrawingAttributes(drawingAttributes *IInkDrawingAttributes)
}

type IInkStrokeBuilderVtbl struct {
	win32.IInspectableVtbl
	BeginStroke                 uintptr
	AppendToStroke              uintptr
	EndStroke                   uintptr
	CreateStroke                uintptr
	SetDefaultDrawingAttributes uintptr
}

type IInkStrokeBuilder struct {
	win32.IInspectable
}

func (this *IInkStrokeBuilder) Vtbl() *IInkStrokeBuilderVtbl {
	return (*IInkStrokeBuilderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkStrokeBuilder) BeginStroke(pointerPoint *IPointerPoint) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().BeginStroke, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pointerPoint)))
	_ = _hr
}

func (this *IInkStrokeBuilder) AppendToStroke(pointerPoint *IPointerPoint) *IPointerPoint {
	var _result *IPointerPoint
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AppendToStroke, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pointerPoint)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInkStrokeBuilder) EndStroke(pointerPoint *IPointerPoint) *IInkStroke {
	var _result *IInkStroke
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EndStroke, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pointerPoint)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInkStrokeBuilder) CreateStroke(points *IIterable[Point]) *IInkStroke {
	var _result *IInkStroke
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateStroke, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(points)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInkStrokeBuilder) SetDefaultDrawingAttributes(drawingAttributes *IInkDrawingAttributes) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetDefaultDrawingAttributes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(drawingAttributes)))
	_ = _hr
}

// BD82BC27-731F-4CBC-BBBF-6D468044F1E5
var IID_IInkStrokeBuilder2 = syscall.GUID{0xBD82BC27, 0x731F, 0x4CBC,
	[8]byte{0xBB, 0xBF, 0x6D, 0x46, 0x80, 0x44, 0xF1, 0xE5}}

type IInkStrokeBuilder2Interface interface {
	win32.IInspectableInterface
	CreateStrokeFromInkPoints(inkPoints *IIterable[*IInkPoint], transform unsafe.Pointer) *IInkStroke
}

type IInkStrokeBuilder2Vtbl struct {
	win32.IInspectableVtbl
	CreateStrokeFromInkPoints uintptr
}

type IInkStrokeBuilder2 struct {
	win32.IInspectable
}

func (this *IInkStrokeBuilder2) Vtbl() *IInkStrokeBuilder2Vtbl {
	return (*IInkStrokeBuilder2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkStrokeBuilder2) CreateStrokeFromInkPoints(inkPoints *IIterable[*IInkPoint], transform unsafe.Pointer) *IInkStroke {
	var _result *IInkStroke
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateStrokeFromInkPoints, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(inkPoints)), uintptr(transform), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B2C71FCD-5472-46B1-A81D-C37A3D169441
var IID_IInkStrokeBuilder3 = syscall.GUID{0xB2C71FCD, 0x5472, 0x46B1,
	[8]byte{0xA8, 0x1D, 0xC3, 0x7A, 0x3D, 0x16, 0x94, 0x41}}

type IInkStrokeBuilder3Interface interface {
	win32.IInspectableInterface
	CreateStrokeFromInkPoints(inkPoints *IIterable[*IInkPoint], transform unsafe.Pointer, strokeStartedTime *IReference[DateTime], strokeDuration *IReference[TimeSpan]) *IInkStroke
}

type IInkStrokeBuilder3Vtbl struct {
	win32.IInspectableVtbl
	CreateStrokeFromInkPoints uintptr
}

type IInkStrokeBuilder3 struct {
	win32.IInspectable
}

func (this *IInkStrokeBuilder3) Vtbl() *IInkStrokeBuilder3Vtbl {
	return (*IInkStrokeBuilder3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkStrokeBuilder3) CreateStrokeFromInkPoints(inkPoints *IIterable[*IInkPoint], transform unsafe.Pointer, strokeStartedTime *IReference[DateTime], strokeDuration *IReference[TimeSpan]) *IInkStroke {
	var _result *IInkStroke
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateStrokeFromInkPoints, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(inkPoints)), uintptr(transform), uintptr(unsafe.Pointer(strokeStartedTime)), uintptr(unsafe.Pointer(strokeDuration)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 22ACCBC6-FAA9-4F14-B68C-F6CEE670AE16
var IID_IInkStrokeContainer = syscall.GUID{0x22ACCBC6, 0xFAA9, 0x4F14,
	[8]byte{0xB6, 0x8C, 0xF6, 0xCE, 0xE6, 0x70, 0xAE, 0x16}}

type IInkStrokeContainerInterface interface {
	win32.IInspectableInterface
	Get_BoundingRect() Rect
	AddStroke(stroke *IInkStroke)
	DeleteSelected() Rect
	MoveSelected(translation Point) Rect
	SelectWithPolyLine(polyline *IIterable[Point]) Rect
	SelectWithLine(from Point, to Point) Rect
	CopySelectedToClipboard()
	PasteFromClipboard(position Point) Rect
	CanPasteFromClipboard() bool
	LoadAsync(inputStream *IInputStream) *IAsyncActionWithProgress[uint64]
	SaveAsync(outputStream *IOutputStream) *IAsyncOperationWithProgress[uint32, uint32]
	UpdateRecognitionResults(recognitionResults *IVectorView[*IInkRecognitionResult])
	GetStrokes() *IVectorView[*IInkStroke]
	GetRecognitionResults() *IVectorView[*IInkRecognitionResult]
}

type IInkStrokeContainerVtbl struct {
	win32.IInspectableVtbl
	Get_BoundingRect         uintptr
	AddStroke                uintptr
	DeleteSelected           uintptr
	MoveSelected             uintptr
	SelectWithPolyLine       uintptr
	SelectWithLine           uintptr
	CopySelectedToClipboard  uintptr
	PasteFromClipboard       uintptr
	CanPasteFromClipboard    uintptr
	LoadAsync                uintptr
	SaveAsync                uintptr
	UpdateRecognitionResults uintptr
	GetStrokes               uintptr
	GetRecognitionResults    uintptr
}

type IInkStrokeContainer struct {
	win32.IInspectable
}

func (this *IInkStrokeContainer) Vtbl() *IInkStrokeContainerVtbl {
	return (*IInkStrokeContainerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkStrokeContainer) Get_BoundingRect() Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BoundingRect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkStrokeContainer) AddStroke(stroke *IInkStroke) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddStroke, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stroke)))
	_ = _hr
}

func (this *IInkStrokeContainer) DeleteSelected() Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DeleteSelected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkStrokeContainer) MoveSelected(translation Point) Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MoveSelected, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&translation)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkStrokeContainer) SelectWithPolyLine(polyline *IIterable[Point]) Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SelectWithPolyLine, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(polyline)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkStrokeContainer) SelectWithLine(from Point, to Point) Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SelectWithLine, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&from)), *(*uintptr)(unsafe.Pointer(&to)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkStrokeContainer) CopySelectedToClipboard() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CopySelectedToClipboard, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IInkStrokeContainer) PasteFromClipboard(position Point) Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PasteFromClipboard, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&position)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkStrokeContainer) CanPasteFromClipboard() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CanPasteFromClipboard, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkStrokeContainer) LoadAsync(inputStream *IInputStream) *IAsyncActionWithProgress[uint64] {
	var _result *IAsyncActionWithProgress[uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LoadAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(inputStream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInkStrokeContainer) SaveAsync(outputStream *IOutputStream) *IAsyncOperationWithProgress[uint32, uint32] {
	var _result *IAsyncOperationWithProgress[uint32, uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SaveAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(outputStream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInkStrokeContainer) UpdateRecognitionResults(recognitionResults *IVectorView[*IInkRecognitionResult]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UpdateRecognitionResults, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(recognitionResults)))
	_ = _hr
}

func (this *IInkStrokeContainer) GetStrokes() *IVectorView[*IInkStroke] {
	var _result *IVectorView[*IInkStroke]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStrokes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInkStrokeContainer) GetRecognitionResults() *IVectorView[*IInkRecognitionResult] {
	var _result *IVectorView[*IInkRecognitionResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetRecognitionResults, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8901D364-DA36-4BCF-9E5C-D195825995B4
var IID_IInkStrokeContainer2 = syscall.GUID{0x8901D364, 0xDA36, 0x4BCF,
	[8]byte{0x9E, 0x5C, 0xD1, 0x95, 0x82, 0x59, 0x95, 0xB4}}

type IInkStrokeContainer2Interface interface {
	win32.IInspectableInterface
	AddStrokes(strokes *IIterable[*IInkStroke])
	Clear()
}

type IInkStrokeContainer2Vtbl struct {
	win32.IInspectableVtbl
	AddStrokes uintptr
	Clear      uintptr
}

type IInkStrokeContainer2 struct {
	win32.IInspectable
}

func (this *IInkStrokeContainer2) Vtbl() *IInkStrokeContainer2Vtbl {
	return (*IInkStrokeContainer2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkStrokeContainer2) AddStrokes(strokes *IIterable[*IInkStroke]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddStrokes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(strokes)))
	_ = _hr
}

func (this *IInkStrokeContainer2) Clear() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Clear, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 3D07BEA5-BAEA-4C82-A719-7B83DA1067D2
var IID_IInkStrokeContainer3 = syscall.GUID{0x3D07BEA5, 0xBAEA, 0x4C82,
	[8]byte{0xA7, 0x19, 0x7B, 0x83, 0xDA, 0x10, 0x67, 0xD2}}

type IInkStrokeContainer3Interface interface {
	win32.IInspectableInterface
	SaveWithFormatAsync(outputStream *IOutputStream, inkPersistenceFormat InkPersistenceFormat) *IAsyncOperationWithProgress[uint32, uint32]
	GetStrokeById(id uint32) *IInkStroke
}

type IInkStrokeContainer3Vtbl struct {
	win32.IInspectableVtbl
	SaveWithFormatAsync uintptr
	GetStrokeById       uintptr
}

type IInkStrokeContainer3 struct {
	win32.IInspectable
}

func (this *IInkStrokeContainer3) Vtbl() *IInkStrokeContainer3Vtbl {
	return (*IInkStrokeContainer3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkStrokeContainer3) SaveWithFormatAsync(outputStream *IOutputStream, inkPersistenceFormat InkPersistenceFormat) *IAsyncOperationWithProgress[uint32, uint32] {
	var _result *IAsyncOperationWithProgress[uint32, uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SaveWithFormatAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(outputStream)), uintptr(inkPersistenceFormat), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInkStrokeContainer3) GetStrokeById(id uint32) *IInkStroke {
	var _result *IInkStroke
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStrokeById, uintptr(unsafe.Pointer(this)), uintptr(id), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CF2FFE7B-5E10-43C6-A080-88F26E1DC67D
var IID_IInkStrokeInput = syscall.GUID{0xCF2FFE7B, 0x5E10, 0x43C6,
	[8]byte{0xA0, 0x80, 0x88, 0xF2, 0x6E, 0x1D, 0xC6, 0x7D}}

type IInkStrokeInputInterface interface {
	win32.IInspectableInterface
	Add_StrokeStarted(handler TypedEventHandler[*IInkStrokeInput, unsafe.Pointer]) EventRegistrationToken
	Remove_StrokeStarted(cookie EventRegistrationToken)
	Add_StrokeContinued(handler TypedEventHandler[*IInkStrokeInput, unsafe.Pointer]) EventRegistrationToken
	Remove_StrokeContinued(cookie EventRegistrationToken)
	Add_StrokeEnded(handler TypedEventHandler[*IInkStrokeInput, unsafe.Pointer]) EventRegistrationToken
	Remove_StrokeEnded(cookie EventRegistrationToken)
	Add_StrokeCanceled(handler TypedEventHandler[*IInkStrokeInput, unsafe.Pointer]) EventRegistrationToken
	Remove_StrokeCanceled(cookie EventRegistrationToken)
	Get_InkPresenter() *IInkPresenter
}

type IInkStrokeInputVtbl struct {
	win32.IInspectableVtbl
	Add_StrokeStarted      uintptr
	Remove_StrokeStarted   uintptr
	Add_StrokeContinued    uintptr
	Remove_StrokeContinued uintptr
	Add_StrokeEnded        uintptr
	Remove_StrokeEnded     uintptr
	Add_StrokeCanceled     uintptr
	Remove_StrokeCanceled  uintptr
	Get_InkPresenter       uintptr
}

type IInkStrokeInput struct {
	win32.IInspectable
}

func (this *IInkStrokeInput) Vtbl() *IInkStrokeInputVtbl {
	return (*IInkStrokeInputVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkStrokeInput) Add_StrokeStarted(handler TypedEventHandler[*IInkStrokeInput, unsafe.Pointer]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StrokeStarted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkStrokeInput) Remove_StrokeStarted(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StrokeStarted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IInkStrokeInput) Add_StrokeContinued(handler TypedEventHandler[*IInkStrokeInput, unsafe.Pointer]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StrokeContinued, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkStrokeInput) Remove_StrokeContinued(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StrokeContinued, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IInkStrokeInput) Add_StrokeEnded(handler TypedEventHandler[*IInkStrokeInput, unsafe.Pointer]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StrokeEnded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkStrokeInput) Remove_StrokeEnded(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StrokeEnded, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IInkStrokeInput) Add_StrokeCanceled(handler TypedEventHandler[*IInkStrokeInput, unsafe.Pointer]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StrokeCanceled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkStrokeInput) Remove_StrokeCanceled(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StrokeCanceled, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IInkStrokeInput) Get_InkPresenter() *IInkPresenter {
	var _result *IInkPresenter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InkPresenter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 68510F1F-88E3-477A-A2FA-569F5F1F9BD5
var IID_IInkStrokeRenderingSegment = syscall.GUID{0x68510F1F, 0x88E3, 0x477A,
	[8]byte{0xA2, 0xFA, 0x56, 0x9F, 0x5F, 0x1F, 0x9B, 0xD5}}

type IInkStrokeRenderingSegmentInterface interface {
	win32.IInspectableInterface
	Get_Position() Point
	Get_BezierControlPoint1() Point
	Get_BezierControlPoint2() Point
	Get_Pressure() float32
	Get_TiltX() float32
	Get_TiltY() float32
	Get_Twist() float32
}

type IInkStrokeRenderingSegmentVtbl struct {
	win32.IInspectableVtbl
	Get_Position            uintptr
	Get_BezierControlPoint1 uintptr
	Get_BezierControlPoint2 uintptr
	Get_Pressure            uintptr
	Get_TiltX               uintptr
	Get_TiltY               uintptr
	Get_Twist               uintptr
}

type IInkStrokeRenderingSegment struct {
	win32.IInspectable
}

func (this *IInkStrokeRenderingSegment) Vtbl() *IInkStrokeRenderingSegmentVtbl {
	return (*IInkStrokeRenderingSegmentVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkStrokeRenderingSegment) Get_Position() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkStrokeRenderingSegment) Get_BezierControlPoint1() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BezierControlPoint1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkStrokeRenderingSegment) Get_BezierControlPoint2() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BezierControlPoint2, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkStrokeRenderingSegment) Get_Pressure() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Pressure, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkStrokeRenderingSegment) Get_TiltX() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TiltX, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkStrokeRenderingSegment) Get_TiltY() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TiltY, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkStrokeRenderingSegment) Get_Twist() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Twist, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C4F3F229-1938-495C-B4D9-6DE4B08D4811
var IID_IInkStrokesCollectedEventArgs = syscall.GUID{0xC4F3F229, 0x1938, 0x495C,
	[8]byte{0xB4, 0xD9, 0x6D, 0xE4, 0xB0, 0x8D, 0x48, 0x11}}

type IInkStrokesCollectedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Strokes() *IVectorView[*IInkStroke]
}

type IInkStrokesCollectedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Strokes uintptr
}

type IInkStrokesCollectedEventArgs struct {
	win32.IInspectable
}

func (this *IInkStrokesCollectedEventArgs) Vtbl() *IInkStrokesCollectedEventArgsVtbl {
	return (*IInkStrokesCollectedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkStrokesCollectedEventArgs) Get_Strokes() *IVectorView[*IInkStroke] {
	var _result *IVectorView[*IInkStroke]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Strokes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A4216A22-1503-4EBF-8FF5-2DE84584A8AA
var IID_IInkStrokesErasedEventArgs = syscall.GUID{0xA4216A22, 0x1503, 0x4EBF,
	[8]byte{0x8F, 0xF5, 0x2D, 0xE8, 0x45, 0x84, 0xA8, 0xAA}}

type IInkStrokesErasedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Strokes() *IVectorView[*IInkStroke]
}

type IInkStrokesErasedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Strokes uintptr
}

type IInkStrokesErasedEventArgs struct {
	win32.IInspectable
}

func (this *IInkStrokesErasedEventArgs) Vtbl() *IInkStrokesErasedEventArgsVtbl {
	return (*IInkStrokesErasedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkStrokesErasedEventArgs) Get_Strokes() *IVectorView[*IInkStroke] {
	var _result *IVectorView[*IInkStroke]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Strokes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9B9EA160-AE9B-45F9-8407-4B493B163661
var IID_IInkSynchronizer = syscall.GUID{0x9B9EA160, 0xAE9B, 0x45F9,
	[8]byte{0x84, 0x07, 0x4B, 0x49, 0x3B, 0x16, 0x36, 0x61}}

type IInkSynchronizerInterface interface {
	win32.IInspectableInterface
	BeginDry() *IVectorView[*IInkStroke]
	EndDry()
}

type IInkSynchronizerVtbl struct {
	win32.IInspectableVtbl
	BeginDry uintptr
	EndDry   uintptr
}

type IInkSynchronizer struct {
	win32.IInspectable
}

func (this *IInkSynchronizer) Vtbl() *IInkSynchronizerVtbl {
	return (*IInkSynchronizerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkSynchronizer) BeginDry() *IVectorView[*IInkStroke] {
	var _result *IVectorView[*IInkStroke]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().BeginDry, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInkSynchronizer) EndDry() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().EndDry, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// DB4445E0-8398-4921-AC3B-AB978C5BA256
var IID_IInkUnprocessedInput = syscall.GUID{0xDB4445E0, 0x8398, 0x4921,
	[8]byte{0xAC, 0x3B, 0xAB, 0x97, 0x8C, 0x5B, 0xA2, 0x56}}

type IInkUnprocessedInputInterface interface {
	win32.IInspectableInterface
	Add_PointerEntered(handler TypedEventHandler[*IInkUnprocessedInput, unsafe.Pointer]) EventRegistrationToken
	Remove_PointerEntered(cookie EventRegistrationToken)
	Add_PointerHovered(handler TypedEventHandler[*IInkUnprocessedInput, unsafe.Pointer]) EventRegistrationToken
	Remove_PointerHovered(cookie EventRegistrationToken)
	Add_PointerExited(handler TypedEventHandler[*IInkUnprocessedInput, unsafe.Pointer]) EventRegistrationToken
	Remove_PointerExited(cookie EventRegistrationToken)
	Add_PointerPressed(handler TypedEventHandler[*IInkUnprocessedInput, unsafe.Pointer]) EventRegistrationToken
	Remove_PointerPressed(cookie EventRegistrationToken)
	Add_PointerMoved(handler TypedEventHandler[*IInkUnprocessedInput, unsafe.Pointer]) EventRegistrationToken
	Remove_PointerMoved(cookie EventRegistrationToken)
	Add_PointerReleased(handler TypedEventHandler[*IInkUnprocessedInput, unsafe.Pointer]) EventRegistrationToken
	Remove_PointerReleased(cookie EventRegistrationToken)
	Add_PointerLost(handler TypedEventHandler[*IInkUnprocessedInput, unsafe.Pointer]) EventRegistrationToken
	Remove_PointerLost(cookie EventRegistrationToken)
	Get_InkPresenter() *IInkPresenter
}

type IInkUnprocessedInputVtbl struct {
	win32.IInspectableVtbl
	Add_PointerEntered     uintptr
	Remove_PointerEntered  uintptr
	Add_PointerHovered     uintptr
	Remove_PointerHovered  uintptr
	Add_PointerExited      uintptr
	Remove_PointerExited   uintptr
	Add_PointerPressed     uintptr
	Remove_PointerPressed  uintptr
	Add_PointerMoved       uintptr
	Remove_PointerMoved    uintptr
	Add_PointerReleased    uintptr
	Remove_PointerReleased uintptr
	Add_PointerLost        uintptr
	Remove_PointerLost     uintptr
	Get_InkPresenter       uintptr
}

type IInkUnprocessedInput struct {
	win32.IInspectable
}

func (this *IInkUnprocessedInput) Vtbl() *IInkUnprocessedInputVtbl {
	return (*IInkUnprocessedInputVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkUnprocessedInput) Add_PointerEntered(handler TypedEventHandler[*IInkUnprocessedInput, unsafe.Pointer]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PointerEntered, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkUnprocessedInput) Remove_PointerEntered(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PointerEntered, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IInkUnprocessedInput) Add_PointerHovered(handler TypedEventHandler[*IInkUnprocessedInput, unsafe.Pointer]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PointerHovered, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkUnprocessedInput) Remove_PointerHovered(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PointerHovered, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IInkUnprocessedInput) Add_PointerExited(handler TypedEventHandler[*IInkUnprocessedInput, unsafe.Pointer]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PointerExited, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkUnprocessedInput) Remove_PointerExited(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PointerExited, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IInkUnprocessedInput) Add_PointerPressed(handler TypedEventHandler[*IInkUnprocessedInput, unsafe.Pointer]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PointerPressed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkUnprocessedInput) Remove_PointerPressed(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PointerPressed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IInkUnprocessedInput) Add_PointerMoved(handler TypedEventHandler[*IInkUnprocessedInput, unsafe.Pointer]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PointerMoved, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkUnprocessedInput) Remove_PointerMoved(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PointerMoved, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IInkUnprocessedInput) Add_PointerReleased(handler TypedEventHandler[*IInkUnprocessedInput, unsafe.Pointer]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PointerReleased, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkUnprocessedInput) Remove_PointerReleased(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PointerReleased, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IInkUnprocessedInput) Add_PointerLost(handler TypedEventHandler[*IInkUnprocessedInput, unsafe.Pointer]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PointerLost, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkUnprocessedInput) Remove_PointerLost(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PointerLost, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *IInkUnprocessedInput) Get_InkPresenter() *IInkPresenter {
	var _result *IInkPresenter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InkPresenter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BC2CEB8F-0066-44A8-BB7A-B839B3DEB8F5
var IID_IPenAndInkSettings = syscall.GUID{0xBC2CEB8F, 0x0066, 0x44A8,
	[8]byte{0xBB, 0x7A, 0xB8, 0x39, 0xB3, 0xDE, 0xB8, 0xF5}}

type IPenAndInkSettingsInterface interface {
	win32.IInspectableInterface
	Get_IsHandwritingDirectlyIntoTextFieldEnabled() bool
	Get_PenHandedness() PenHandedness
	Get_HandwritingLineHeight() HandwritingLineHeight
	Get_FontFamilyName() string
	Get_UserConsentsToHandwritingTelemetryCollection() bool
	Get_IsTouchHandwritingEnabled() bool
}

type IPenAndInkSettingsVtbl struct {
	win32.IInspectableVtbl
	Get_IsHandwritingDirectlyIntoTextFieldEnabled    uintptr
	Get_PenHandedness                                uintptr
	Get_HandwritingLineHeight                        uintptr
	Get_FontFamilyName                               uintptr
	Get_UserConsentsToHandwritingTelemetryCollection uintptr
	Get_IsTouchHandwritingEnabled                    uintptr
}

type IPenAndInkSettings struct {
	win32.IInspectable
}

func (this *IPenAndInkSettings) Vtbl() *IPenAndInkSettingsVtbl {
	return (*IPenAndInkSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPenAndInkSettings) Get_IsHandwritingDirectlyIntoTextFieldEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsHandwritingDirectlyIntoTextFieldEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPenAndInkSettings) Get_PenHandedness() PenHandedness {
	var _result PenHandedness
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PenHandedness, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPenAndInkSettings) Get_HandwritingLineHeight() HandwritingLineHeight {
	var _result HandwritingLineHeight
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HandwritingLineHeight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPenAndInkSettings) Get_FontFamilyName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FontFamilyName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IPenAndInkSettings) Get_UserConsentsToHandwritingTelemetryCollection() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UserConsentsToHandwritingTelemetryCollection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPenAndInkSettings) Get_IsTouchHandwritingEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsTouchHandwritingEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 3262DA53-1F44-55E2-9929-EBF77E5481B8
var IID_IPenAndInkSettings2 = syscall.GUID{0x3262DA53, 0x1F44, 0x55E2,
	[8]byte{0x99, 0x29, 0xEB, 0xF7, 0x7E, 0x54, 0x81, 0xB8}}

type IPenAndInkSettings2Interface interface {
	win32.IInspectableInterface
	SetPenHandedness(value PenHandedness)
}

type IPenAndInkSettings2Vtbl struct {
	win32.IInspectableVtbl
	SetPenHandedness uintptr
}

type IPenAndInkSettings2 struct {
	win32.IInspectable
}

func (this *IPenAndInkSettings2) Vtbl() *IPenAndInkSettings2Vtbl {
	return (*IPenAndInkSettings2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPenAndInkSettings2) SetPenHandedness(value PenHandedness) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetPenHandedness, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// ED6DD036-5708-5C3C-96DB-F2F552EAB641
var IID_IPenAndInkSettingsStatics = syscall.GUID{0xED6DD036, 0x5708, 0x5C3C,
	[8]byte{0x96, 0xDB, 0xF2, 0xF5, 0x52, 0xEA, 0xB6, 0x41}}

type IPenAndInkSettingsStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *IPenAndInkSettings
}

type IPenAndInkSettingsStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault uintptr
}

type IPenAndInkSettingsStatics struct {
	win32.IInspectable
}

func (this *IPenAndInkSettingsStatics) Vtbl() *IPenAndInkSettingsStaticsVtbl {
	return (*IPenAndInkSettingsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPenAndInkSettingsStatics) GetDefault() *IPenAndInkSettings {
	var _result *IPenAndInkSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type InkDrawingAttributes struct {
	RtClass
	*IInkDrawingAttributes
}

func NewInkDrawingAttributes() *InkDrawingAttributes {
	hs := NewHStr("Windows.UI.Input.Inking.InkDrawingAttributes")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &InkDrawingAttributes{
		RtClass:               RtClass{PInspect: p},
		IInkDrawingAttributes: (*IInkDrawingAttributes)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewIInkDrawingAttributesStatics() *IInkDrawingAttributesStatics {
	var p *IInkDrawingAttributesStatics
	hs := NewHStr("Windows.UI.Input.Inking.InkDrawingAttributes")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IInkDrawingAttributesStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type InkDrawingAttributesPencilProperties struct {
	RtClass
	*IInkDrawingAttributesPencilProperties
}

type InkInputProcessingConfiguration struct {
	RtClass
	*IInkInputProcessingConfiguration
}

type InkManager struct {
	RtClass
	*IInkManager
}

func NewInkManager() *InkManager {
	hs := NewHStr("Windows.UI.Input.Inking.InkManager")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &InkManager{
		RtClass:     RtClass{PInspect: p},
		IInkManager: (*IInkManager)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type InkModelerAttributes struct {
	RtClass
	*IInkModelerAttributes
}

type InkPoint struct {
	RtClass
	*IInkPoint
}

func NewInkPoint_CreateInkPoint(position Point, pressure float32) *InkPoint {
	hs := NewHStr("Windows.UI.Input.Inking.InkPoint")
	var pFac *IInkPointFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IInkPointFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IInkPoint
	p = pFac.CreateInkPoint(position, pressure)
	result := &InkPoint{
		RtClass:   RtClass{PInspect: &p.IInspectable},
		IInkPoint: p,
	}
	com.AddToScope(result)
	return result
}

type InkPresenter struct {
	RtClass
	*IInkPresenter
}

type InkPresenterProtractor struct {
	RtClass
	*IInkPresenterProtractor
}

func NewInkPresenterProtractor_Create(inkPresenter *IInkPresenter) *InkPresenterProtractor {
	hs := NewHStr("Windows.UI.Input.Inking.InkPresenterProtractor")
	var pFac *IInkPresenterProtractorFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IInkPresenterProtractorFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IInkPresenterProtractor
	p = pFac.Create(inkPresenter)
	result := &InkPresenterProtractor{
		RtClass:                 RtClass{PInspect: &p.IInspectable},
		IInkPresenterProtractor: p,
	}
	com.AddToScope(result)
	return result
}

type InkPresenterRuler struct {
	RtClass
	*IInkPresenterRuler
}

func NewInkPresenterRuler_Create(inkPresenter *IInkPresenter) *InkPresenterRuler {
	hs := NewHStr("Windows.UI.Input.Inking.InkPresenterRuler")
	var pFac *IInkPresenterRulerFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IInkPresenterRulerFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IInkPresenterRuler
	p = pFac.Create(inkPresenter)
	result := &InkPresenterRuler{
		RtClass:            RtClass{PInspect: &p.IInspectable},
		IInkPresenterRuler: p,
	}
	com.AddToScope(result)
	return result
}

type InkRecognitionResult struct {
	RtClass
	*IInkRecognitionResult
}

type InkRecognizer struct {
	RtClass
	*IInkRecognizer
}

type InkRecognizerContainer struct {
	RtClass
	*IInkRecognizerContainer
}

func NewInkRecognizerContainer() *InkRecognizerContainer {
	hs := NewHStr("Windows.UI.Input.Inking.InkRecognizerContainer")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &InkRecognizerContainer{
		RtClass:                 RtClass{PInspect: p},
		IInkRecognizerContainer: (*IInkRecognizerContainer)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type InkStroke struct {
	RtClass
	*IInkStroke
}

type InkStrokeBuilder struct {
	RtClass
	*IInkStrokeBuilder
}

func NewInkStrokeBuilder() *InkStrokeBuilder {
	hs := NewHStr("Windows.UI.Input.Inking.InkStrokeBuilder")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &InkStrokeBuilder{
		RtClass:           RtClass{PInspect: p},
		IInkStrokeBuilder: (*IInkStrokeBuilder)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type InkStrokeContainer struct {
	RtClass
	*IInkStrokeContainer
}

func NewInkStrokeContainer() *InkStrokeContainer {
	hs := NewHStr("Windows.UI.Input.Inking.InkStrokeContainer")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &InkStrokeContainer{
		RtClass:             RtClass{PInspect: p},
		IInkStrokeContainer: (*IInkStrokeContainer)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type InkStrokeInput struct {
	RtClass
	*IInkStrokeInput
}

type InkStrokeRenderingSegment struct {
	RtClass
	*IInkStrokeRenderingSegment
}

type InkStrokesCollectedEventArgs struct {
	RtClass
	*IInkStrokesCollectedEventArgs
}

type InkStrokesErasedEventArgs struct {
	RtClass
	*IInkStrokesErasedEventArgs
}

type InkSynchronizer struct {
	RtClass
	*IInkSynchronizer
}

type InkUnprocessedInput struct {
	RtClass
	*IInkUnprocessedInput
}
