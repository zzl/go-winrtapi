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
type InkAnalysisDrawingKind int32

const (
	InkAnalysisDrawingKind_Drawing             InkAnalysisDrawingKind = 0
	InkAnalysisDrawingKind_Circle              InkAnalysisDrawingKind = 1
	InkAnalysisDrawingKind_Ellipse             InkAnalysisDrawingKind = 2
	InkAnalysisDrawingKind_Triangle            InkAnalysisDrawingKind = 3
	InkAnalysisDrawingKind_IsoscelesTriangle   InkAnalysisDrawingKind = 4
	InkAnalysisDrawingKind_EquilateralTriangle InkAnalysisDrawingKind = 5
	InkAnalysisDrawingKind_RightTriangle       InkAnalysisDrawingKind = 6
	InkAnalysisDrawingKind_Quadrilateral       InkAnalysisDrawingKind = 7
	InkAnalysisDrawingKind_Rectangle           InkAnalysisDrawingKind = 8
	InkAnalysisDrawingKind_Square              InkAnalysisDrawingKind = 9
	InkAnalysisDrawingKind_Diamond             InkAnalysisDrawingKind = 10
	InkAnalysisDrawingKind_Trapezoid           InkAnalysisDrawingKind = 11
	InkAnalysisDrawingKind_Parallelogram       InkAnalysisDrawingKind = 12
	InkAnalysisDrawingKind_Pentagon            InkAnalysisDrawingKind = 13
	InkAnalysisDrawingKind_Hexagon             InkAnalysisDrawingKind = 14
)

// enum
type InkAnalysisNodeKind int32

const (
	InkAnalysisNodeKind_UnclassifiedInk InkAnalysisNodeKind = 0
	InkAnalysisNodeKind_Root            InkAnalysisNodeKind = 1
	InkAnalysisNodeKind_WritingRegion   InkAnalysisNodeKind = 2
	InkAnalysisNodeKind_Paragraph       InkAnalysisNodeKind = 3
	InkAnalysisNodeKind_Line            InkAnalysisNodeKind = 4
	InkAnalysisNodeKind_InkWord         InkAnalysisNodeKind = 5
	InkAnalysisNodeKind_InkBullet       InkAnalysisNodeKind = 6
	InkAnalysisNodeKind_InkDrawing      InkAnalysisNodeKind = 7
	InkAnalysisNodeKind_ListItem        InkAnalysisNodeKind = 8
)

// enum
type InkAnalysisStatus int32

const (
	InkAnalysisStatus_Updated   InkAnalysisStatus = 0
	InkAnalysisStatus_Unchanged InkAnalysisStatus = 1
)

// enum
type InkAnalysisStrokeKind int32

const (
	InkAnalysisStrokeKind_Auto    InkAnalysisStrokeKind = 0
	InkAnalysisStrokeKind_Writing InkAnalysisStrokeKind = 1
	InkAnalysisStrokeKind_Drawing InkAnalysisStrokeKind = 2
)

// interfaces

// EE049368-6110-4136-95F9-EE809FC20030
var IID_IInkAnalysisInkBullet = syscall.GUID{0xEE049368, 0x6110, 0x4136,
	[8]byte{0x95, 0xF9, 0xEE, 0x80, 0x9F, 0xC2, 0x00, 0x30}}

type IInkAnalysisInkBulletInterface interface {
	win32.IInspectableInterface
	Get_RecognizedText() string
}

type IInkAnalysisInkBulletVtbl struct {
	win32.IInspectableVtbl
	Get_RecognizedText uintptr
}

type IInkAnalysisInkBullet struct {
	win32.IInspectable
}

func (this *IInkAnalysisInkBullet) Vtbl() *IInkAnalysisInkBulletVtbl {
	return (*IInkAnalysisInkBulletVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkAnalysisInkBullet) Get_RecognizedText() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RecognizedText, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 6A85ED1F-1FE4-4E15-898C-8E112377E021
var IID_IInkAnalysisInkDrawing = syscall.GUID{0x6A85ED1F, 0x1FE4, 0x4E15,
	[8]byte{0x89, 0x8C, 0x8E, 0x11, 0x23, 0x77, 0xE0, 0x21}}

type IInkAnalysisInkDrawingInterface interface {
	win32.IInspectableInterface
	Get_DrawingKind() InkAnalysisDrawingKind
	Get_Center() Point
	Get_Points() *IVectorView[Point]
}

type IInkAnalysisInkDrawingVtbl struct {
	win32.IInspectableVtbl
	Get_DrawingKind uintptr
	Get_Center      uintptr
	Get_Points      uintptr
}

type IInkAnalysisInkDrawing struct {
	win32.IInspectable
}

func (this *IInkAnalysisInkDrawing) Vtbl() *IInkAnalysisInkDrawingVtbl {
	return (*IInkAnalysisInkDrawingVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkAnalysisInkDrawing) Get_DrawingKind() InkAnalysisDrawingKind {
	var _result InkAnalysisDrawingKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DrawingKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkAnalysisInkDrawing) Get_Center() Point {
	var _result Point
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Center, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkAnalysisInkDrawing) Get_Points() *IVectorView[Point] {
	var _result *IVectorView[Point]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Points, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4BD228AD-83AF-4034-8F3B-F8687DFFF436
var IID_IInkAnalysisInkWord = syscall.GUID{0x4BD228AD, 0x83AF, 0x4034,
	[8]byte{0x8F, 0x3B, 0xF8, 0x68, 0x7D, 0xFF, 0xF4, 0x36}}

type IInkAnalysisInkWordInterface interface {
	win32.IInspectableInterface
	Get_RecognizedText() string
	Get_TextAlternates() *IVectorView[string]
}

type IInkAnalysisInkWordVtbl struct {
	win32.IInspectableVtbl
	Get_RecognizedText uintptr
	Get_TextAlternates uintptr
}

type IInkAnalysisInkWord struct {
	win32.IInspectable
}

func (this *IInkAnalysisInkWord) Vtbl() *IInkAnalysisInkWordVtbl {
	return (*IInkAnalysisInkWordVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkAnalysisInkWord) Get_RecognizedText() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RecognizedText, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IInkAnalysisInkWord) Get_TextAlternates() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TextAlternates, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A06D048D-2B8D-4754-AD5A-D0871193A956
var IID_IInkAnalysisLine = syscall.GUID{0xA06D048D, 0x2B8D, 0x4754,
	[8]byte{0xAD, 0x5A, 0xD0, 0x87, 0x11, 0x93, 0xA9, 0x56}}

type IInkAnalysisLineInterface interface {
	win32.IInspectableInterface
	Get_RecognizedText() string
	Get_IndentLevel() int32
}

type IInkAnalysisLineVtbl struct {
	win32.IInspectableVtbl
	Get_RecognizedText uintptr
	Get_IndentLevel    uintptr
}

type IInkAnalysisLine struct {
	win32.IInspectable
}

func (this *IInkAnalysisLine) Vtbl() *IInkAnalysisLineVtbl {
	return (*IInkAnalysisLineVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkAnalysisLine) Get_RecognizedText() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RecognizedText, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IInkAnalysisLine) Get_IndentLevel() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IndentLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// B4E3C23F-C4C3-4C3A-A1A6-9D85547EE586
var IID_IInkAnalysisListItem = syscall.GUID{0xB4E3C23F, 0xC4C3, 0x4C3A,
	[8]byte{0xA1, 0xA6, 0x9D, 0x85, 0x54, 0x7E, 0xE5, 0x86}}

type IInkAnalysisListItemInterface interface {
	win32.IInspectableInterface
	Get_RecognizedText() string
}

type IInkAnalysisListItemVtbl struct {
	win32.IInspectableVtbl
	Get_RecognizedText uintptr
}

type IInkAnalysisListItem struct {
	win32.IInspectable
}

func (this *IInkAnalysisListItem) Vtbl() *IInkAnalysisListItemVtbl {
	return (*IInkAnalysisListItemVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkAnalysisListItem) Get_RecognizedText() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RecognizedText, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 30831F05-5F64-4A2C-BA37-4F4887879574
var IID_IInkAnalysisNode = syscall.GUID{0x30831F05, 0x5F64, 0x4A2C,
	[8]byte{0xBA, 0x37, 0x4F, 0x48, 0x87, 0x87, 0x95, 0x74}}

type IInkAnalysisNodeInterface interface {
	win32.IInspectableInterface
	Get_Id() uint32
	Get_Kind() InkAnalysisNodeKind
	Get_BoundingRect() Rect
	Get_RotatedBoundingRect() *IVectorView[Point]
	Get_Children() *IVectorView[*IInkAnalysisNode]
	Get_Parent() *IInkAnalysisNode
	GetStrokeIds() *IVectorView[uint32]
}

type IInkAnalysisNodeVtbl struct {
	win32.IInspectableVtbl
	Get_Id                  uintptr
	Get_Kind                uintptr
	Get_BoundingRect        uintptr
	Get_RotatedBoundingRect uintptr
	Get_Children            uintptr
	Get_Parent              uintptr
	GetStrokeIds            uintptr
}

type IInkAnalysisNode struct {
	win32.IInspectable
}

func (this *IInkAnalysisNode) Vtbl() *IInkAnalysisNodeVtbl {
	return (*IInkAnalysisNodeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkAnalysisNode) Get_Id() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkAnalysisNode) Get_Kind() InkAnalysisNodeKind {
	var _result InkAnalysisNodeKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Kind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkAnalysisNode) Get_BoundingRect() Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BoundingRect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkAnalysisNode) Get_RotatedBoundingRect() *IVectorView[Point] {
	var _result *IVectorView[Point]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RotatedBoundingRect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInkAnalysisNode) Get_Children() *IVectorView[*IInkAnalysisNode] {
	var _result *IVectorView[*IInkAnalysisNode]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Children, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInkAnalysisNode) Get_Parent() *IInkAnalysisNode {
	var _result *IInkAnalysisNode
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Parent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInkAnalysisNode) GetStrokeIds() *IVectorView[uint32] {
	var _result *IVectorView[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStrokeIds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D9AD045C-0CD1-4DD4-A68B-EB1F12B3D727
var IID_IInkAnalysisParagraph = syscall.GUID{0xD9AD045C, 0x0CD1, 0x4DD4,
	[8]byte{0xA6, 0x8B, 0xEB, 0x1F, 0x12, 0xB3, 0xD7, 0x27}}

type IInkAnalysisParagraphInterface interface {
	win32.IInspectableInterface
	Get_RecognizedText() string
}

type IInkAnalysisParagraphVtbl struct {
	win32.IInspectableVtbl
	Get_RecognizedText uintptr
}

type IInkAnalysisParagraph struct {
	win32.IInspectable
}

func (this *IInkAnalysisParagraph) Vtbl() *IInkAnalysisParagraphVtbl {
	return (*IInkAnalysisParagraphVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkAnalysisParagraph) Get_RecognizedText() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RecognizedText, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 8948BA79-A243-4AA3-A294-1F98BD0FF580
var IID_IInkAnalysisResult = syscall.GUID{0x8948BA79, 0xA243, 0x4AA3,
	[8]byte{0xA2, 0x94, 0x1F, 0x98, 0xBD, 0x0F, 0xF5, 0x80}}

type IInkAnalysisResultInterface interface {
	win32.IInspectableInterface
	Get_Status() InkAnalysisStatus
}

type IInkAnalysisResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
}

type IInkAnalysisResult struct {
	win32.IInspectable
}

func (this *IInkAnalysisResult) Vtbl() *IInkAnalysisResultVtbl {
	return (*IInkAnalysisResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkAnalysisResult) Get_Status() InkAnalysisStatus {
	var _result InkAnalysisStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 3FB6A3C4-2FDE-4061-8502-A90F32545B84
var IID_IInkAnalysisRoot = syscall.GUID{0x3FB6A3C4, 0x2FDE, 0x4061,
	[8]byte{0x85, 0x02, 0xA9, 0x0F, 0x32, 0x54, 0x5B, 0x84}}

type IInkAnalysisRootInterface interface {
	win32.IInspectableInterface
	Get_RecognizedText() string
	FindNodes(nodeKind InkAnalysisNodeKind) *IVectorView[*IInkAnalysisNode]
}

type IInkAnalysisRootVtbl struct {
	win32.IInspectableVtbl
	Get_RecognizedText uintptr
	FindNodes          uintptr
}

type IInkAnalysisRoot struct {
	win32.IInspectable
}

func (this *IInkAnalysisRoot) Vtbl() *IInkAnalysisRootVtbl {
	return (*IInkAnalysisRootVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkAnalysisRoot) Get_RecognizedText() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RecognizedText, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IInkAnalysisRoot) FindNodes(nodeKind InkAnalysisNodeKind) *IVectorView[*IInkAnalysisNode] {
	var _result *IVectorView[*IInkAnalysisNode]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindNodes, uintptr(unsafe.Pointer(this)), uintptr(nodeKind), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DD6D6231-BD16-4663-B5AE-941D3043EF5B
var IID_IInkAnalysisWritingRegion = syscall.GUID{0xDD6D6231, 0xBD16, 0x4663,
	[8]byte{0xB5, 0xAE, 0x94, 0x1D, 0x30, 0x43, 0xEF, 0x5B}}

type IInkAnalysisWritingRegionInterface interface {
	win32.IInspectableInterface
	Get_RecognizedText() string
}

type IInkAnalysisWritingRegionVtbl struct {
	win32.IInspectableVtbl
	Get_RecognizedText uintptr
}

type IInkAnalysisWritingRegion struct {
	win32.IInspectable
}

func (this *IInkAnalysisWritingRegion) Vtbl() *IInkAnalysisWritingRegionVtbl {
	return (*IInkAnalysisWritingRegionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkAnalysisWritingRegion) Get_RecognizedText() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RecognizedText, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// F12B8F95-0866-4DC5-8C77-F88614DFE38C
var IID_IInkAnalyzer = syscall.GUID{0xF12B8F95, 0x0866, 0x4DC5,
	[8]byte{0x8C, 0x77, 0xF8, 0x86, 0x14, 0xDF, 0xE3, 0x8C}}

type IInkAnalyzerInterface interface {
	win32.IInspectableInterface
	Get_AnalysisRoot() *IInkAnalysisRoot
	Get_IsAnalyzing() bool
	AddDataForStroke(stroke *IInkStroke)
	AddDataForStrokes(strokes *IIterable[*IInkStroke])
	ClearDataForAllStrokes()
	RemoveDataForStroke(strokeId uint32)
	RemoveDataForStrokes(strokeIds *IIterable[uint32])
	ReplaceDataForStroke(stroke *IInkStroke)
	SetStrokeDataKind(strokeId uint32, strokeKind InkAnalysisStrokeKind)
	AnalyzeAsync() *IAsyncOperation[*IInkAnalysisResult]
}

type IInkAnalyzerVtbl struct {
	win32.IInspectableVtbl
	Get_AnalysisRoot       uintptr
	Get_IsAnalyzing        uintptr
	AddDataForStroke       uintptr
	AddDataForStrokes      uintptr
	ClearDataForAllStrokes uintptr
	RemoveDataForStroke    uintptr
	RemoveDataForStrokes   uintptr
	ReplaceDataForStroke   uintptr
	SetStrokeDataKind      uintptr
	AnalyzeAsync           uintptr
}

type IInkAnalyzer struct {
	win32.IInspectable
}

func (this *IInkAnalyzer) Vtbl() *IInkAnalyzerVtbl {
	return (*IInkAnalyzerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkAnalyzer) Get_AnalysisRoot() *IInkAnalysisRoot {
	var _result *IInkAnalysisRoot
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AnalysisRoot, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IInkAnalyzer) Get_IsAnalyzing() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsAnalyzing, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IInkAnalyzer) AddDataForStroke(stroke *IInkStroke) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddDataForStroke, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stroke)))
	_ = _hr
}

func (this *IInkAnalyzer) AddDataForStrokes(strokes *IIterable[*IInkStroke]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddDataForStrokes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(strokes)))
	_ = _hr
}

func (this *IInkAnalyzer) ClearDataForAllStrokes() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearDataForAllStrokes, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IInkAnalyzer) RemoveDataForStroke(strokeId uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveDataForStroke, uintptr(unsafe.Pointer(this)), uintptr(strokeId))
	_ = _hr
}

func (this *IInkAnalyzer) RemoveDataForStrokes(strokeIds *IIterable[uint32]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveDataForStrokes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(strokeIds)))
	_ = _hr
}

func (this *IInkAnalyzer) ReplaceDataForStroke(stroke *IInkStroke) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReplaceDataForStroke, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stroke)))
	_ = _hr
}

func (this *IInkAnalyzer) SetStrokeDataKind(strokeId uint32, strokeKind InkAnalysisStrokeKind) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetStrokeDataKind, uintptr(unsafe.Pointer(this)), uintptr(strokeId), uintptr(strokeKind))
	_ = _hr
}

func (this *IInkAnalyzer) AnalyzeAsync() *IAsyncOperation[*IInkAnalysisResult] {
	var _result *IAsyncOperation[*IInkAnalysisResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AnalyzeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 29138686-1963-49D8-9589-E14384C769E3
var IID_IInkAnalyzerFactory = syscall.GUID{0x29138686, 0x1963, 0x49D8,
	[8]byte{0x95, 0x89, 0xE1, 0x43, 0x84, 0xC7, 0x69, 0xE3}}

type IInkAnalyzerFactoryInterface interface {
	win32.IInspectableInterface
	CreateAnalyzer() *IInkAnalyzer
}

type IInkAnalyzerFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateAnalyzer uintptr
}

type IInkAnalyzerFactory struct {
	win32.IInspectable
}

func (this *IInkAnalyzerFactory) Vtbl() *IInkAnalyzerFactoryVtbl {
	return (*IInkAnalyzerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInkAnalyzerFactory) CreateAnalyzer() *IInkAnalyzer {
	var _result *IInkAnalyzer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateAnalyzer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type InkAnalysisInkBullet struct {
	RtClass
	*IInkAnalysisInkBullet
}

type InkAnalysisInkDrawing struct {
	RtClass
	*IInkAnalysisInkDrawing
}

type InkAnalysisInkWord struct {
	RtClass
	*IInkAnalysisInkWord
}

type InkAnalysisLine struct {
	RtClass
	*IInkAnalysisLine
}

type InkAnalysisListItem struct {
	RtClass
	*IInkAnalysisListItem
}

type InkAnalysisNode struct {
	RtClass
	*IInkAnalysisNode
}

type InkAnalysisParagraph struct {
	RtClass
	*IInkAnalysisParagraph
}

type InkAnalysisResult struct {
	RtClass
	*IInkAnalysisResult
}

type InkAnalysisRoot struct {
	RtClass
	*IInkAnalysisRoot
}

type InkAnalysisWritingRegion struct {
	RtClass
	*IInkAnalysisWritingRegion
}

type InkAnalyzer struct {
	RtClass
	*IInkAnalyzer
}

func NewInkAnalyzer() *InkAnalyzer {
	hs := NewHStr("Windows.UI.Input.Inking.Analysis.InkAnalyzer")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &InkAnalyzer{
		RtClass:      RtClass{PInspect: p},
		IInkAnalyzer: (*IInkAnalyzer)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}
