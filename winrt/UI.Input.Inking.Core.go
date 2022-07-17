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
type CoreWetStrokeDisposition int32

const (
	CoreWetStrokeDisposition_Inking    CoreWetStrokeDisposition = 0
	CoreWetStrokeDisposition_Completed CoreWetStrokeDisposition = 1
	CoreWetStrokeDisposition_Canceled  CoreWetStrokeDisposition = 2
)

// interfaces

// FDA015D3-9D66-4F7D-A57F-CC70B9CFAA76
var IID_ICoreIncrementalInkStroke = syscall.GUID{0xFDA015D3, 0x9D66, 0x4F7D,
	[8]byte{0xA5, 0x7F, 0xCC, 0x70, 0xB9, 0xCF, 0xAA, 0x76}}

type ICoreIncrementalInkStrokeInterface interface {
	win32.IInspectableInterface
	AppendInkPoints(inkPoints *IIterable[*IInkPoint]) Rect
	CreateInkStroke() *IInkStroke
	Get_DrawingAttributes() *IInkDrawingAttributes
	Get_PointTransform() unsafe.Pointer
	Get_BoundingRect() Rect
}

type ICoreIncrementalInkStrokeVtbl struct {
	win32.IInspectableVtbl
	AppendInkPoints       uintptr
	CreateInkStroke       uintptr
	Get_DrawingAttributes uintptr
	Get_PointTransform    uintptr
	Get_BoundingRect      uintptr
}

type ICoreIncrementalInkStroke struct {
	win32.IInspectable
}

func (this *ICoreIncrementalInkStroke) Vtbl() *ICoreIncrementalInkStrokeVtbl {
	return (*ICoreIncrementalInkStrokeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreIncrementalInkStroke) AppendInkPoints(inkPoints *IIterable[*IInkPoint]) Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AppendInkPoints, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(inkPoints)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreIncrementalInkStroke) CreateInkStroke() *IInkStroke {
	var _result *IInkStroke
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateInkStroke, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICoreIncrementalInkStroke) Get_DrawingAttributes() *IInkDrawingAttributes {
	var _result *IInkDrawingAttributes
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DrawingAttributes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICoreIncrementalInkStroke) Get_PointTransform() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PointTransform, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreIncrementalInkStroke) Get_BoundingRect() Rect {
	var _result Rect
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BoundingRect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// D7C59F46-8DA8-4F70-9751-E53BB6DF4596
var IID_ICoreIncrementalInkStrokeFactory = syscall.GUID{0xD7C59F46, 0x8DA8, 0x4F70,
	[8]byte{0x97, 0x51, 0xE5, 0x3B, 0xB6, 0xDF, 0x45, 0x96}}

type ICoreIncrementalInkStrokeFactoryInterface interface {
	win32.IInspectableInterface
	Create(drawingAttributes *IInkDrawingAttributes, pointTransform unsafe.Pointer) *ICoreIncrementalInkStroke
}

type ICoreIncrementalInkStrokeFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ICoreIncrementalInkStrokeFactory struct {
	win32.IInspectable
}

func (this *ICoreIncrementalInkStrokeFactory) Vtbl() *ICoreIncrementalInkStrokeFactoryVtbl {
	return (*ICoreIncrementalInkStrokeFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreIncrementalInkStrokeFactory) Create(drawingAttributes *IInkDrawingAttributes, pointTransform unsafe.Pointer) *ICoreIncrementalInkStroke {
	var _result *ICoreIncrementalInkStroke
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(drawingAttributes)), uintptr(pointTransform), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 39B38DA9-7639-4499-A5B5-191D00E35B16
var IID_ICoreInkIndependentInputSource = syscall.GUID{0x39B38DA9, 0x7639, 0x4499,
	[8]byte{0xA5, 0xB5, 0x19, 0x1D, 0x00, 0xE3, 0x5B, 0x16}}

type ICoreInkIndependentInputSourceInterface interface {
	win32.IInspectableInterface
	Add_PointerEntering(handler TypedEventHandler[*ICoreInkIndependentInputSource, unsafe.Pointer]) EventRegistrationToken
	Remove_PointerEntering(cookie EventRegistrationToken)
	Add_PointerHovering(handler TypedEventHandler[*ICoreInkIndependentInputSource, unsafe.Pointer]) EventRegistrationToken
	Remove_PointerHovering(cookie EventRegistrationToken)
	Add_PointerExiting(handler TypedEventHandler[*ICoreInkIndependentInputSource, unsafe.Pointer]) EventRegistrationToken
	Remove_PointerExiting(cookie EventRegistrationToken)
	Add_PointerPressing(handler TypedEventHandler[*ICoreInkIndependentInputSource, unsafe.Pointer]) EventRegistrationToken
	Remove_PointerPressing(cookie EventRegistrationToken)
	Add_PointerMoving(handler TypedEventHandler[*ICoreInkIndependentInputSource, unsafe.Pointer]) EventRegistrationToken
	Remove_PointerMoving(cookie EventRegistrationToken)
	Add_PointerReleasing(handler TypedEventHandler[*ICoreInkIndependentInputSource, unsafe.Pointer]) EventRegistrationToken
	Remove_PointerReleasing(cookie EventRegistrationToken)
	Add_PointerLost(handler TypedEventHandler[*ICoreInkIndependentInputSource, unsafe.Pointer]) EventRegistrationToken
	Remove_PointerLost(cookie EventRegistrationToken)
	Get_InkPresenter() *IInkPresenter
}

type ICoreInkIndependentInputSourceVtbl struct {
	win32.IInspectableVtbl
	Add_PointerEntering     uintptr
	Remove_PointerEntering  uintptr
	Add_PointerHovering     uintptr
	Remove_PointerHovering  uintptr
	Add_PointerExiting      uintptr
	Remove_PointerExiting   uintptr
	Add_PointerPressing     uintptr
	Remove_PointerPressing  uintptr
	Add_PointerMoving       uintptr
	Remove_PointerMoving    uintptr
	Add_PointerReleasing    uintptr
	Remove_PointerReleasing uintptr
	Add_PointerLost         uintptr
	Remove_PointerLost      uintptr
	Get_InkPresenter        uintptr
}

type ICoreInkIndependentInputSource struct {
	win32.IInspectable
}

func (this *ICoreInkIndependentInputSource) Vtbl() *ICoreInkIndependentInputSourceVtbl {
	return (*ICoreInkIndependentInputSourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreInkIndependentInputSource) Add_PointerEntering(handler TypedEventHandler[*ICoreInkIndependentInputSource, unsafe.Pointer]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PointerEntering, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreInkIndependentInputSource) Remove_PointerEntering(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PointerEntering, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *ICoreInkIndependentInputSource) Add_PointerHovering(handler TypedEventHandler[*ICoreInkIndependentInputSource, unsafe.Pointer]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PointerHovering, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreInkIndependentInputSource) Remove_PointerHovering(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PointerHovering, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *ICoreInkIndependentInputSource) Add_PointerExiting(handler TypedEventHandler[*ICoreInkIndependentInputSource, unsafe.Pointer]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PointerExiting, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreInkIndependentInputSource) Remove_PointerExiting(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PointerExiting, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *ICoreInkIndependentInputSource) Add_PointerPressing(handler TypedEventHandler[*ICoreInkIndependentInputSource, unsafe.Pointer]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PointerPressing, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreInkIndependentInputSource) Remove_PointerPressing(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PointerPressing, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *ICoreInkIndependentInputSource) Add_PointerMoving(handler TypedEventHandler[*ICoreInkIndependentInputSource, unsafe.Pointer]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PointerMoving, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreInkIndependentInputSource) Remove_PointerMoving(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PointerMoving, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *ICoreInkIndependentInputSource) Add_PointerReleasing(handler TypedEventHandler[*ICoreInkIndependentInputSource, unsafe.Pointer]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PointerReleasing, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreInkIndependentInputSource) Remove_PointerReleasing(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PointerReleasing, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *ICoreInkIndependentInputSource) Add_PointerLost(handler TypedEventHandler[*ICoreInkIndependentInputSource, unsafe.Pointer]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PointerLost, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreInkIndependentInputSource) Remove_PointerLost(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PointerLost, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *ICoreInkIndependentInputSource) Get_InkPresenter() *IInkPresenter {
	var _result *IInkPresenter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InkPresenter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2846B012-0B59-5BB9-A3C5-BECB7CF03A33
var IID_ICoreInkIndependentInputSource2 = syscall.GUID{0x2846B012, 0x0B59, 0x5BB9,
	[8]byte{0xA3, 0xC5, 0xBE, 0xCB, 0x7C, 0xF0, 0x3A, 0x33}}

type ICoreInkIndependentInputSource2Interface interface {
	win32.IInspectableInterface
	Get_PointerCursor() unsafe.Pointer
	Put_PointerCursor(value unsafe.Pointer)
}

type ICoreInkIndependentInputSource2Vtbl struct {
	win32.IInspectableVtbl
	Get_PointerCursor uintptr
	Put_PointerCursor uintptr
}

type ICoreInkIndependentInputSource2 struct {
	win32.IInspectable
}

func (this *ICoreInkIndependentInputSource2) Vtbl() *ICoreInkIndependentInputSource2Vtbl {
	return (*ICoreInkIndependentInputSource2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreInkIndependentInputSource2) Get_PointerCursor() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PointerCursor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreInkIndependentInputSource2) Put_PointerCursor(value unsafe.Pointer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PointerCursor, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 73E6011B-80C0-4DFB-9B66-10BA7F3F9C84
var IID_ICoreInkIndependentInputSourceStatics = syscall.GUID{0x73E6011B, 0x80C0, 0x4DFB,
	[8]byte{0x9B, 0x66, 0x10, 0xBA, 0x7F, 0x3F, 0x9C, 0x84}}

type ICoreInkIndependentInputSourceStaticsInterface interface {
	win32.IInspectableInterface
	Create(inkPresenter *IInkPresenter) *ICoreInkIndependentInputSource
}

type ICoreInkIndependentInputSourceStaticsVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ICoreInkIndependentInputSourceStatics struct {
	win32.IInspectable
}

func (this *ICoreInkIndependentInputSourceStatics) Vtbl() *ICoreInkIndependentInputSourceStaticsVtbl {
	return (*ICoreInkIndependentInputSourceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreInkIndependentInputSourceStatics) Create(inkPresenter *IInkPresenter) *ICoreInkIndependentInputSource {
	var _result *ICoreInkIndependentInputSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(inkPresenter)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 396E89E6-7D55-4617-9E58-68C70C9169B9
var IID_ICoreInkPresenterHost = syscall.GUID{0x396E89E6, 0x7D55, 0x4617,
	[8]byte{0x9E, 0x58, 0x68, 0xC7, 0x0C, 0x91, 0x69, 0xB9}}

type ICoreInkPresenterHostInterface interface {
	win32.IInspectableInterface
	Get_InkPresenter() *IInkPresenter
	Get_RootVisual() *IContainerVisual
	Put_RootVisual(value *IContainerVisual)
}

type ICoreInkPresenterHostVtbl struct {
	win32.IInspectableVtbl
	Get_InkPresenter uintptr
	Get_RootVisual   uintptr
	Put_RootVisual   uintptr
}

type ICoreInkPresenterHost struct {
	win32.IInspectable
}

func (this *ICoreInkPresenterHost) Vtbl() *ICoreInkPresenterHostVtbl {
	return (*ICoreInkPresenterHostVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreInkPresenterHost) Get_InkPresenter() *IInkPresenter {
	var _result *IInkPresenter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InkPresenter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICoreInkPresenterHost) Get_RootVisual() *IContainerVisual {
	var _result *IContainerVisual
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RootVisual, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICoreInkPresenterHost) Put_RootVisual(value *IContainerVisual) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RootVisual, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

// FB07D14C-3380-457A-A987-991357896C1B
var IID_ICoreWetStrokeUpdateEventArgs = syscall.GUID{0xFB07D14C, 0x3380, 0x457A,
	[8]byte{0xA9, 0x87, 0x99, 0x13, 0x57, 0x89, 0x6C, 0x1B}}

type ICoreWetStrokeUpdateEventArgsInterface interface {
	win32.IInspectableInterface
	Get_NewInkPoints() *IVector[*IInkPoint]
	Get_PointerId() uint32
	Get_Disposition() CoreWetStrokeDisposition
	Put_Disposition(value CoreWetStrokeDisposition)
}

type ICoreWetStrokeUpdateEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_NewInkPoints uintptr
	Get_PointerId    uintptr
	Get_Disposition  uintptr
	Put_Disposition  uintptr
}

type ICoreWetStrokeUpdateEventArgs struct {
	win32.IInspectable
}

func (this *ICoreWetStrokeUpdateEventArgs) Vtbl() *ICoreWetStrokeUpdateEventArgsVtbl {
	return (*ICoreWetStrokeUpdateEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreWetStrokeUpdateEventArgs) Get_NewInkPoints() *IVector[*IInkPoint] {
	var _result *IVector[*IInkPoint]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NewInkPoints, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICoreWetStrokeUpdateEventArgs) Get_PointerId() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PointerId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreWetStrokeUpdateEventArgs) Get_Disposition() CoreWetStrokeDisposition {
	var _result CoreWetStrokeDisposition
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Disposition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreWetStrokeUpdateEventArgs) Put_Disposition(value CoreWetStrokeDisposition) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Disposition, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 1F718E22-EE52-4E00-8209-4C3E5B21A3CC
var IID_ICoreWetStrokeUpdateSource = syscall.GUID{0x1F718E22, 0xEE52, 0x4E00,
	[8]byte{0x82, 0x09, 0x4C, 0x3E, 0x5B, 0x21, 0xA3, 0xCC}}

type ICoreWetStrokeUpdateSourceInterface interface {
	win32.IInspectableInterface
	Add_WetStrokeStarting(handler TypedEventHandler[*ICoreWetStrokeUpdateSource, *ICoreWetStrokeUpdateEventArgs]) EventRegistrationToken
	Remove_WetStrokeStarting(cookie EventRegistrationToken)
	Add_WetStrokeContinuing(handler TypedEventHandler[*ICoreWetStrokeUpdateSource, *ICoreWetStrokeUpdateEventArgs]) EventRegistrationToken
	Remove_WetStrokeContinuing(cookie EventRegistrationToken)
	Add_WetStrokeStopping(handler TypedEventHandler[*ICoreWetStrokeUpdateSource, *ICoreWetStrokeUpdateEventArgs]) EventRegistrationToken
	Remove_WetStrokeStopping(cookie EventRegistrationToken)
	Add_WetStrokeCompleted(handler TypedEventHandler[*ICoreWetStrokeUpdateSource, *ICoreWetStrokeUpdateEventArgs]) EventRegistrationToken
	Remove_WetStrokeCompleted(cookie EventRegistrationToken)
	Add_WetStrokeCanceled(handler TypedEventHandler[*ICoreWetStrokeUpdateSource, *ICoreWetStrokeUpdateEventArgs]) EventRegistrationToken
	Remove_WetStrokeCanceled(cookie EventRegistrationToken)
	Get_InkPresenter() *IInkPresenter
}

type ICoreWetStrokeUpdateSourceVtbl struct {
	win32.IInspectableVtbl
	Add_WetStrokeStarting      uintptr
	Remove_WetStrokeStarting   uintptr
	Add_WetStrokeContinuing    uintptr
	Remove_WetStrokeContinuing uintptr
	Add_WetStrokeStopping      uintptr
	Remove_WetStrokeStopping   uintptr
	Add_WetStrokeCompleted     uintptr
	Remove_WetStrokeCompleted  uintptr
	Add_WetStrokeCanceled      uintptr
	Remove_WetStrokeCanceled   uintptr
	Get_InkPresenter           uintptr
}

type ICoreWetStrokeUpdateSource struct {
	win32.IInspectable
}

func (this *ICoreWetStrokeUpdateSource) Vtbl() *ICoreWetStrokeUpdateSourceVtbl {
	return (*ICoreWetStrokeUpdateSourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreWetStrokeUpdateSource) Add_WetStrokeStarting(handler TypedEventHandler[*ICoreWetStrokeUpdateSource, *ICoreWetStrokeUpdateEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_WetStrokeStarting, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreWetStrokeUpdateSource) Remove_WetStrokeStarting(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_WetStrokeStarting, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *ICoreWetStrokeUpdateSource) Add_WetStrokeContinuing(handler TypedEventHandler[*ICoreWetStrokeUpdateSource, *ICoreWetStrokeUpdateEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_WetStrokeContinuing, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreWetStrokeUpdateSource) Remove_WetStrokeContinuing(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_WetStrokeContinuing, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *ICoreWetStrokeUpdateSource) Add_WetStrokeStopping(handler TypedEventHandler[*ICoreWetStrokeUpdateSource, *ICoreWetStrokeUpdateEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_WetStrokeStopping, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreWetStrokeUpdateSource) Remove_WetStrokeStopping(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_WetStrokeStopping, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *ICoreWetStrokeUpdateSource) Add_WetStrokeCompleted(handler TypedEventHandler[*ICoreWetStrokeUpdateSource, *ICoreWetStrokeUpdateEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_WetStrokeCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreWetStrokeUpdateSource) Remove_WetStrokeCompleted(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_WetStrokeCompleted, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *ICoreWetStrokeUpdateSource) Add_WetStrokeCanceled(handler TypedEventHandler[*ICoreWetStrokeUpdateSource, *ICoreWetStrokeUpdateEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_WetStrokeCanceled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICoreWetStrokeUpdateSource) Remove_WetStrokeCanceled(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_WetStrokeCanceled, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *ICoreWetStrokeUpdateSource) Get_InkPresenter() *IInkPresenter {
	var _result *IInkPresenter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InkPresenter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3DAD9CBA-1D3D-46AE-AB9D-8647486C6F90
var IID_ICoreWetStrokeUpdateSourceStatics = syscall.GUID{0x3DAD9CBA, 0x1D3D, 0x46AE,
	[8]byte{0xAB, 0x9D, 0x86, 0x47, 0x48, 0x6C, 0x6F, 0x90}}

type ICoreWetStrokeUpdateSourceStaticsInterface interface {
	win32.IInspectableInterface
	Create(inkPresenter *IInkPresenter) *ICoreWetStrokeUpdateSource
}

type ICoreWetStrokeUpdateSourceStaticsVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ICoreWetStrokeUpdateSourceStatics struct {
	win32.IInspectable
}

func (this *ICoreWetStrokeUpdateSourceStatics) Vtbl() *ICoreWetStrokeUpdateSourceStaticsVtbl {
	return (*ICoreWetStrokeUpdateSourceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICoreWetStrokeUpdateSourceStatics) Create(inkPresenter *IInkPresenter) *ICoreWetStrokeUpdateSource {
	var _result *ICoreWetStrokeUpdateSource
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(inkPresenter)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type CoreIncrementalInkStroke struct {
	RtClass
	*ICoreIncrementalInkStroke
}

func NewCoreIncrementalInkStroke_Create(drawingAttributes *IInkDrawingAttributes, pointTransform unsafe.Pointer) *CoreIncrementalInkStroke {
	hs := NewHStr("Windows.UI.Input.Inking.Core.CoreIncrementalInkStroke")
	var pFac *ICoreIncrementalInkStrokeFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICoreIncrementalInkStrokeFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ICoreIncrementalInkStroke
	p = pFac.Create(drawingAttributes, pointTransform)
	result := &CoreIncrementalInkStroke{
		RtClass:                   RtClass{PInspect: &p.IInspectable},
		ICoreIncrementalInkStroke: p,
	}
	com.AddToScope(result)
	return result
}

type CoreInkIndependentInputSource struct {
	RtClass
	*ICoreInkIndependentInputSource
}

func NewICoreInkIndependentInputSourceStatics() *ICoreInkIndependentInputSourceStatics {
	var p *ICoreInkIndependentInputSourceStatics
	hs := NewHStr("Windows.UI.Input.Inking.Core.CoreInkIndependentInputSource")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICoreInkIndependentInputSourceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type CoreInkPresenterHost struct {
	RtClass
	*ICoreInkPresenterHost
}

func NewCoreInkPresenterHost() *CoreInkPresenterHost {
	hs := NewHStr("Windows.UI.Input.Inking.Core.CoreInkPresenterHost")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &CoreInkPresenterHost{
		RtClass:               RtClass{PInspect: p},
		ICoreInkPresenterHost: (*ICoreInkPresenterHost)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type CoreWetStrokeUpdateEventArgs struct {
	RtClass
	*ICoreWetStrokeUpdateEventArgs
}

type CoreWetStrokeUpdateSource struct {
	RtClass
	*ICoreWetStrokeUpdateSource
}

func NewICoreWetStrokeUpdateSourceStatics() *ICoreWetStrokeUpdateSourceStatics {
	var p *ICoreWetStrokeUpdateSourceStatics
	hs := NewHStr("Windows.UI.Input.Inking.Core.CoreWetStrokeUpdateSource")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ICoreWetStrokeUpdateSourceStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
